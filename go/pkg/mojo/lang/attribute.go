package lang

import "errors"

func (m *Attribute) SameName(fullName string) bool {
	if m == nil {
		return false
	}

	pkg, name := ParseIdentifierName(fullName)
	if len(pkg) > 0 {
		return m.Package == pkg && m.Name == name
	} else {
		return m.Name == name
	}
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
				attribute.Arguments[0].Value = NewBoolLiteralExpr(&BoolLiteralExpr{
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
		Package: pkg,
		Name:    n,
		Arguments: []*Argument{
			{
				Value: NewBoolLiteralExpr(&BoolLiteralExpr{
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
		Package: pkg,
		Name:    n,
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
		Package: pkg,
		Name:    n,
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
