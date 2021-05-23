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

func NewStructDecl(decl *StructDecl) *Declaration {
	return &Declaration{
		Declaration: &Declaration_StructDecl{
			StructDecl: decl,
		},
	}
}

func NewInterfaceDecl(decl *InterfaceDecl) *Declaration {
	return &Declaration{
		Declaration: &Declaration_InterfaceDecl{
			InterfaceDecl: decl,
		},
	}
}

func NewTypeAliasDecl(decl *TypeAliasDecl) *Declaration {
	return &Declaration{
		Declaration: &Declaration_TypeAliasDecl{
			TypeAliasDecl: decl,
		},
	}
}

func NewDeclarationFromTypeDeclaration(decl *TypeDeclaration) *Declaration {
	switch decl.TypeDeclaration.(type) {
	case *TypeDeclaration_StructDecl:
		return &Declaration{
			Declaration: &Declaration_StructDecl{
				StructDecl: decl.GetStructDecl(),
			},
		}
	case *TypeDeclaration_InterfaceDecl:
		return &Declaration{
			Declaration: &Declaration_InterfaceDecl{
				InterfaceDecl: decl.GetInterfaceDecl(),
			},
		}
	case *TypeDeclaration_EnumDecl:
		return &Declaration{
			Declaration: &Declaration_EnumDecl{
				EnumDecl: decl.GetEnumDecl(),
			},
		}
	case *TypeDeclaration_TypeAliasDecl:
		return &Declaration{
			Declaration: &Declaration_TypeAliasDecl{
				TypeAliasDecl: decl.GetTypeAliasDecl(),
			},
		}
	case *TypeDeclaration_GenericParameter:
		return &Declaration{
			Declaration: &Declaration_GenericParameter{
				GenericParameter: decl.GetGenericParameter(),
			},
		}
	default:
		return nil
	}
}

func NewGenericParameterDeclaration(parameter *GenericParameter) *Declaration {
	return &Declaration{
		Declaration: &Declaration_GenericParameter{
			GenericParameter: parameter,
		},
	}
}

func (m *Declaration) GetName() string {
	switch m.Declaration.(type) {
	case *Declaration_StructDecl:
		return m.GetStructDecl().Name
	case *Declaration_InterfaceDecl:
		return m.GetInterfaceDecl().Name
	case *Declaration_EnumDecl:
		return m.GetEnumDecl().Name
	case *Declaration_TypeAliasDecl:
		return m.GetTypeAliasDecl().Name
	case *Declaration_GenericParameter:
		return m.GetGenericParameter().Name
	default:
		return ""
	}
}

func (m *Declaration) GetPackageName() string {
	switch m.Declaration.(type) {
	case *Declaration_StructDecl:
		return m.GetStructDecl().PackageName
	case *Declaration_InterfaceDecl:
		return m.GetInterfaceDecl().PackageName
	case *Declaration_EnumDecl:
		return m.GetEnumDecl().PackageName
	case *Declaration_TypeAliasDecl:
		return m.GetTypeAliasDecl().PackageName
	case *Declaration_GenericParameter:
		return m.GetGenericParameter().PackageName
	default:
		return ""
	}
}

func (m *Declaration) SetPackageName(name string) {
	switch m.Declaration.(type) {
	case *Declaration_StructDecl:
		m.GetStructDecl().PackageName = name
	case *Declaration_InterfaceDecl:
		m.GetInterfaceDecl().PackageName = name
	case *Declaration_EnumDecl:
		m.GetEnumDecl().PackageName = name
	case *Declaration_TypeAliasDecl:
		m.GetTypeAliasDecl().PackageName = name
	case *Declaration_GenericParameter:
		m.GetGenericParameter().PackageName = name
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

func GetBoolAttribute(attributes []*Attribute, name string) (bool, error) {
	for _, attribute := range attributes {
		if attribute.Name == name {
			if len(attribute.Arguments) > 0 {
				if value := attribute.Arguments[0].GetBoolLiteralExpr(); value != nil {
					return value.Value, nil
				}
			}
		}
	}
	return false, errors.New("NotFound")
}

func SetBoolAttribute(attributes []*Attribute, name string, value bool) []*Attribute {
	for _, attribute := range attributes {
		if attribute.Name == name {
			if len(attribute.Arguments) > 0 {
				attribute.Arguments[0] = NewBoolLiteralExpr(&BoolLiteralExpr{
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
		Arguments: []*Expression{NewBoolLiteralExpr(&BoolLiteralExpr{
			Kind:     0,
			Implicit: false,
			Value:    value,
		})},
	})

	return attributes
}

func GetIntegerAttribute(attributes []*Attribute, name string) (int64, error) {
	for _, attribute := range attributes {
		if attribute.Name == name {
			if len(attribute.Arguments) > 0 {
				if integer := attribute.Arguments[0].GetIntegerLiteralExpr(); integer != nil {
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
			if len(attribute.Arguments) > 0 {
				attribute.Arguments[0] = NewIntegerLiteralExpr(&IntegerLiteralExpr{
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
		Arguments: []*Expression{NewIntegerLiteralExpr(&IntegerLiteralExpr{
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
			if len(attribute.Arguments) > 0 {
				if value := attribute.Arguments[0].GetStringLiteralExpr(); value != nil {
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
			if len(attribute.Arguments) > 0 {
				attribute.Arguments[0] = NewStringLiteralExpr(&StringLiteralExpr{
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
		Arguments: []*Expression{NewStringLiteralExpr(&StringLiteralExpr{
			Kind:     0,
			Implicit: false,
			Value:    value,
		})},
	})

	return attributes
}

func RemoveAttribute(attributes []*Attribute, name string) []*Attribute {
	var newAttributes []*Attribute
	for _, attribute := range attributes {
		if attribute.Name == name {
			continue
		}
		newAttributes = append(newAttributes, attribute)
	}

	return newAttributes
}
