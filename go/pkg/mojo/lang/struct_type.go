package lang

import "github.com/mojo-lang/core/go/pkg/mojo/core"

type FieldNamOption int32

const (
	FieldNamOptionDefault FieldNamOption = iota
	FieldNamOptionUseAlias
)

func (m *StructType) SetStartPosition(position *Position) {
	if m != nil {
		m.StartPosition = PatchPosition(m.StartPosition, position)
	}
}

func (m *StructType) SetEndPosition(position *Position) {
	if m != nil {
		m.EndPosition = PatchPosition(m.EndPosition, position)
	}
}

func (m *StructType) FieldNames(option FieldNamOption) []string {
	if m != nil {
		var fieldNames []string
		for _, inherit := range m.Inherits {
			names := inherit.GetTypeDeclaration().GetStructDecl().FieldNames(option)
			fieldNames = append(fieldNames, names...)
		}

		for _, field := range m.Fields {
			if option == FieldNamOptionUseAlias {
				if alias, _ := field.GetStringAttribute(core.AliasAttributeName); len(alias) > 0 {
					fieldNames = append(fieldNames, alias)
					continue
				}
			}

			fieldNames = append(fieldNames, field.Name)
		}
		return fieldNames
	}
	return nil
}

func (m *StructType) GetField(name string) *ValueDecl {
	if m != nil {
		for _, inherit := range m.Inherits {
			if field := inherit.GetTypeDeclaration().GetStructDecl().GetField(name); field != nil {
				return field
			}
		}

		for _, field := range m.Fields {
			if field.Name == name {
				return field
			}
		}
	}
	return nil
}

func (m *StructType) EachField(iter func(decl *ValueDecl) error) error {
	if m != nil {
		for _, inherit := range m.Inherits {
			if err := inherit.GetTypeDeclaration().GetStructDecl().EachField(iter); err != nil {
				return err
			}
		}

		for _, field := range m.Fields {
			if err := iter(field); err != nil {
				return err
			}
		}
	}
	return nil
}
