package lang

func (m *StructType) FieldNames() []string {
	if m != nil {
		var fieldNames []string
		for _, inherit := range m.Inherits {
			names := inherit.GetTypeDeclaration().GetStructDecl().FieldNames()
			fieldNames = append(fieldNames, names...)
		}

		for _, field := range m.Fields {
			fieldNames = append(fieldNames, field.Name)
		}
		return fieldNames
	}
	return nil
}
