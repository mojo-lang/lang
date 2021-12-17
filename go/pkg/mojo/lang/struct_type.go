package lang

import "github.com/mojo-lang/core/go/pkg/mojo/core"

type FieldNamOption int32

const (
	FieldNamOptionDefault FieldNamOption = iota
	FieldNamOptionUseAlias
)

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
