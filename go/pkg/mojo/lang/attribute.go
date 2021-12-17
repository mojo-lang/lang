package lang

import "errors"

func NewBoolAttribute(pkg string, name string) *Attribute {
	return newBoolAttribute(pkg, name, false)
}

func NewStringAttribute(pkg string, name string, value string) *Attribute {
	return newStringAttribute(pkg, name, value, false)
}

func NewImplicitStringAttribute(pkg string, name string, value string) *Attribute {
	return newStringAttribute(pkg, name, value, false)
}

func NewImplicitBoolAttribute(pkg string, name string) *Attribute {
	return newBoolAttribute(pkg, name, true)
}

func newBoolAttribute(pkg string, name string, implicit bool) *Attribute {
	return &Attribute{
		PackageName: pkg,
		Name:        name,
		Arguments: []*Argument{{
			Value: NewBoolLiteralExpression(&BoolLiteralExpr{
				Kind:     0,
				Implicit: implicit,
				Value:    true,
			}),
		},
		},
	}
}

func newStringAttribute(pkg string, name string, value string, implicit bool) *Attribute {
	return &Attribute{
		PackageName: pkg,
		Name:        name,
		Arguments: []*Argument{{
			Value: NewStringLiteralExpression(&StringLiteralExpr{
				Kind:     0,
				Implicit: implicit,
				Value:    value,
			}),
		},
		},
	}
}

func (m *Attribute) SameName(fullName string) bool {
	if m == nil {
		return false
	}

	pkg, name := ParseIdentifierName(fullName)
	if len(pkg) > 0 {
		return m.PackageName == pkg && m.Name == name
	} else {
		return m.Name == name
	}
}

func (m *Attribute) GetFullName() string {
	fullName := ""
	if m != nil {
		if len(m.PackageName) > 0 {
			fullName = m.PackageName + "."
		}
		fullName = fullName + m.Name
	}
	return fullName
}

func (m *Attribute) Repeatable() bool {
	if m != nil && m.Declaration != nil {
		if v, err := GetBoolAttribute(m.Declaration.Attributes, "repeatable"); err == nil {
			return v
		}
	}
	return false
}

func (m *Attribute) GetArrayLiteralArgument() *ArrayLiteralExpr {
	array := &ArrayLiteralExpr{
		Implicit: true,
	}

	if len(m.Arguments) == 1 {
		if arrayLiteral := m.Arguments[0].Value.GetArrayLiteralExpr(); arrayLiteral != nil {
			return arrayLiteral
		}
	}

	for _, argument := range m.Arguments {
		if len(argument.Label) > 0 {
			return nil
		}

		array.Elements = append(array.Elements, argument.Value)
	}
	return array
}

func (m *Attribute) GetObjectLiteralArgument() *ObjectLiteralExpr {
	object := &ObjectLiteralExpr{
		Implicit: true,
	}
	if len(m.Arguments) == 1 {
		if objectLiteral := m.Arguments[0].Value.GetObjectLiteralExpr(); objectLiteral != nil {
			return objectLiteral
		}
	}

	for _, argument := range m.Arguments {
		if len(argument.Label) == 0 {
			return nil
		}

		object.Fields = append(object.Fields, &ObjectLiteralExpr_Field{
			Name:  argument.Label,
			Value: argument.Value,
		})
	}
	return object
}

func (m *Attribute) GetMapLiteralArgument() *MapLiteralExpr {
	if len(m.Arguments) == 1 {
		if mapLiteral := m.Arguments[0].Value.GetMapLiteralExpr(); mapLiteral != nil {
			return mapLiteral
		}
	}

	object := m.GetObjectLiteralArgument()
	if object != nil {
		mapLiteral := &MapLiteralExpr{
			Implicit: true,
		}
		for _, field := range object.Fields {
			mapLiteral.Entries = append(mapLiteral.Entries, &MapLiteralExpr_Entry{
				Key:   field.Name,
				Value: field.Value,
			})
		}
		return mapLiteral
	}
	return nil
}

func GetAttribute(attributes []*Attribute, name string) *Attribute {
	for _, attribute := range attributes {
		if attribute.SameName(name) {
			return attribute
		}
	}
	return nil
}

// make sure returned Argument list which length greater than 0
func GetAttributeArguments(attributes []*Attribute, name string) ([]*Argument, error) {
	for _, attribute := range attributes {
		if attribute.SameName(name) {
			if len(attribute.Arguments) > 0 {
				return attribute.Arguments, nil
			} else {
				break
			}
		}
	}
	return nil, errors.New("NotFound")
}

func GetAttributeArgument(attributes []*Attribute, name string) (*Argument, error) {
	for _, attribute := range attributes {
		if attribute.SameName(name) {
			if len(attribute.Arguments) > 0 {
				return attribute.Arguments[0], nil
			} else {
				break
			}
		}
	}
	return nil, errors.New("NotFound")
}

func HasAttribute(attributes []*Attribute, name string) bool {
	for _, attribute := range attributes {
		if attribute.SameName(name) {
			return true
		}
	}
	return false
}

func GetBoolAttribute(attributes []*Attribute, name string) (bool, error) {
	for _, attribute := range attributes {
		if attribute.SameName(name) {
			if len(attribute.Arguments) > 0 {
				if value := attribute.Arguments[0].Value.GetBoolLiteralExpr(); value != nil {
					return value.Value, nil
				}
			} else { // default value is true
				return true, nil
			}
		}
	}
	return false, errors.New("NotFound")
}

func SetBoolAttribute(attributes []*Attribute, name string, value bool) []*Attribute {
	for _, attribute := range attributes {
		if attribute.SameName(name) {
			if len(attribute.Arguments) > 0 {
				attribute.Arguments[0].Value = NewBoolLiteralExpression(&BoolLiteralExpr{
					Kind:     0,
					Implicit: false,
					Value:    value,
				})
				return attributes
			}
		}
	}

	pkg, n := ParseIdentifierName(name)
	attributes = append(attributes, &Attribute{
		PackageName: pkg,
		Name:        n,
		Arguments: []*Argument{
			{
				Value: NewBoolLiteralExpression(&BoolLiteralExpr{
					Kind:     0,
					Implicit: false,
					Value:    value,
				}),
			},
		},
	})

	return attributes
}

func GetIntegerAttribute(attributes []*Attribute, name string) (int64, error) {
	argument, err := GetAttributeArgument(attributes, name)
	if err != nil {
		return 0, err
	}
	return argument.GetInteger()
}

func SetIntegerAttribute(attributes []*Attribute, name string, value int64) []*Attribute {
	for _, attribute := range attributes {
		if attribute.SameName(name) {
			if len(attribute.Arguments) > 0 {
				attribute.Arguments[0] = NewIntegerLiteralArgument(&IntegerLiteralExpr{
					Kind:     0,
					Implicit: false,
					Value:    value,
				})
				return attributes
			}
		}
	}

	pkg, n := ParseIdentifierName(name)
	attributes = append(attributes, &Attribute{
		PackageName: pkg,
		Name:        n,
		Arguments: []*Argument{NewIntegerLiteralArgument(&IntegerLiteralExpr{
			Kind:     0,
			Implicit: false,
			Value:    value,
		}),
		},
	})

	return attributes
}

func GetStringAttribute(attributes []*Attribute, name string) (string, error) {
	argument, err := GetAttributeArgument(attributes, name)
	if err != nil {
		return "", err
	}
	return argument.GetString()
}

func GetStringValuesAttribute(attributes []*Attribute, name string) ([]string, error) {
	var values []string
	arguments, err := GetAttributeArguments(attributes, name)
	if err != nil {
		return nil, err
	}

	if len(arguments) == 1 {
		if strArray, err := arguments[0].GetStringArray(); err == nil {
			return strArray, nil
		}
	}

	for _, argument := range arguments {
		if val, err := argument.GetString(); err != nil {
			return nil, err
		} else {
			values = append(values, val)
		}
	}
	return values, nil
}

func SetStringAttribute(attributes []*Attribute, name string, value string) []*Attribute {
	for _, attribute := range attributes {
		if attribute.SameName(name) {
			if len(attribute.Arguments) > 0 {
				attribute.Arguments[0] = NewStringLiteralArgument(&StringLiteralExpr{
					Kind:     0,
					Implicit: false,
					Value:    value,
				})
				return attributes
			}
		}
	}

	pkg, n := ParseIdentifierName(name)
	attributes = append(attributes, &Attribute{
		PackageName: pkg,
		Name:        n,
		Arguments: []*Argument{NewStringLiteralArgument(&StringLiteralExpr{
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
		if attribute.SameName(name) {
			continue
		}
		newAttributes = append(newAttributes, attribute)
	}

	return newAttributes
}

func MergeAttributes(dst []*Attribute, src []*Attribute) []*Attribute {
	index := make(map[string][]*Attribute)
	setAttributes := func(attributes []*Attribute) {
		for _, attr := range attributes {
			fullName := attr.GetFullName()
			if attr.Repeatable() || len(index[fullName]) == 0 {
				index[fullName] = append(index[fullName], attr)
			}
		}
	}
	setAttributes(dst)
	setAttributes(src)

	var attributes []*Attribute
	for _, attrs := range index {
		attributes = append(attributes, attrs...)
	}
	return attributes
}
