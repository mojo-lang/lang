package lang

import (
	"errors"
	"io"
	"reflect"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

const (
	EnumDeclName      = "EnumDecl"
	StructDeclName    = "StructDecl"
	InterfaceDeclName = "InterfaceDecl"
)

var declInfos map[reflect.Type]StructJsonInfo

func init() {
	jsoniter.RegisterTypeDecoder("lang.Declaration", &DeclarationCodec{})
	jsoniter.RegisterTypeEncoder("lang.Declaration", &DeclarationCodec{})

	declInfos = GetStructJsonInfos(jsoniter.Config{},
		PackageDecl{},
		ImportDecl{},
		ConstantDecl{},
		EnumDecl{},
		StructDecl{},
		InterfaceDecl{},
		FunctionDecl{},
	)
}

func NewEnumDecl(decl *EnumDecl) *Declaration {
	return &Declaration{
		Declaration: &Declaration_EnumDecl{
			EnumDecl: decl,
		},
	}
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
			//stream.Error = fmt.Errorf("%v.%s", encoder.typ, stream.Error.Error())
		}
	}
}

func (codec *DeclarationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return ((*Declaration)(ptr)).Declaration == nil
}

func GetIntegerAttribute(attributes []*Attribute, name string) (int64, error) {
	for _, attribute := range attributes {
		if attribute.Name == name {
			if len(attribute.Expressions) > 0 {
				integer := attribute.Expressions[0].GetIntegerLiteralExpr()
				if integer != nil {
					return integer.EvalValue(), nil
				}
			}
		}
	}
	return 0, errors.New("NotFound")
}

func SetIntegerAttribute(attributes []*Attribute, name string, value int64) []*Attribute {
	for _, attribute := range attributes {
		if attribute.Name == name {
			if len(attribute.Expressions) > 0 {
				attribute.Expressions[0] = NewIntegerLiteralExpr(&IntegerLiteralExpr{
					Kind:     0,
					Implicit: false,
					Value:    value,
				})
				return attributes
			}
		}
	}

	attributes = append(attributes, &Attribute{
		Name: name,
		Expressions: []*Expression{NewIntegerLiteralExpr(&IntegerLiteralExpr{
			Kind:     0,
			Implicit: false,
			Value:    value,
		})},
	})

	return attributes
}

func GetStringAttribute(attributes []*Attribute, name string) (string, error) {
	for _, attribute := range attributes {
		if attribute.Name == name {
			if len(attribute.Expressions) > 0 {
				value := attribute.Expressions[0].GetStringLiteralExpr()
				if value != nil {
					return value.Value, nil
				}
			}
		}
	}
	return "", errors.New("NotFound")
}

func SetStringAttribute(attributes []*Attribute, name string, value string) []*Attribute {
	for _, attribute := range attributes {
		if attribute.Name == name {
			if len(attribute.Expressions) > 0 {
				attribute.Expressions[0] = NewStringLiteralExpr(&StringLiteralExpr{
					Kind:     0,
					Implicit: false,
					Value:    value,
				})
				return attributes
			}
		}
	}

	attributes = append(attributes, &Attribute{
		Name: name,
		Expressions: []*Expression{NewStringLiteralExpr(&StringLiteralExpr{
			Kind:     0,
			Implicit: false,
			Value:    value,
		})},
	})

	return attributes
}
