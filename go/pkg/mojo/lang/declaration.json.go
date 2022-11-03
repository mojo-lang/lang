package lang

import (
	"io"
	"reflect"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	jsoniter.RegisterTypeDecoder("lang.Declaration", &DeclarationCodec{})
	jsoniter.RegisterTypeEncoder("lang.Declaration", &DeclarationCodec{})
}

type DeclarationCodec struct {
}

func (codec *DeclarationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	any := iter.ReadAny()
	declaration := (*Declaration)(ptr)
	if any.ValueType() == jsoniter.ObjectValue {
		t := any.Get("@type").ToString()
		switch t {
		case EnumDeclName:
			decl := &EnumDecl{}
			any.ToVal(decl)
			declaration.Declaration = &Declaration_EnumDecl{EnumDecl: decl}
		case StructDeclName:
			decl := &StructDecl{}
			any.ToVal(decl)
			declaration.Declaration = &Declaration_StructDecl{StructDecl: decl}
		}
	}
}

func (codec *DeclarationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	declaration := (*Declaration)(ptr)
	if declaration != nil && declaration.Declaration != nil {
		stream.WriteObjectStart()
		stream.WriteObjectField("@type")

		var v reflect.Value
		switch decl := declaration.Declaration.(type) {
		case *Declaration_EnumDecl:
			v = reflect.ValueOf(decl.EnumDecl).Elem()
			stream.WriteVal(EnumDeclName)
		case *Declaration_StructDecl:
			v = reflect.ValueOf(decl.StructDecl).Elem()
			stream.WriteVal(StructDeclName)
		}

		t := declInfos[v.Type()]

		// iterator the fields
		count := v.NumField()
		for i := 0; i < count; i++ {
			info := t.Fields[i]
			field := v.Field(i)

			if info.Ignored {
				continue
			}

			if info.OmitEmpty {
				if field.IsZero() {
					continue
				}

				k := field.Kind()
				switch k {
				case reflect.Map, reflect.Ptr, reflect.Interface, reflect.Struct, reflect.Slice:
					if info.Encoder.IsEmpty(unsafe.Pointer(field.Pointer())) {
						continue
					}
				}
			}

			stream.WriteMore()
			stream.WriteObjectField(t.Fields[i].ToName)
			stream.WriteVal(v.Field(i).Interface())
		}

		stream.WriteObjectEnd()

		if stream.Error != nil && stream.Error != io.EOF {
			// stream.Error = fmt.Errorf("%v.%s", encoder.typ, stream.Error.Error())
		}
	}
}

func (codec *DeclarationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return ((*Declaration)(ptr)).Declaration == nil
}
