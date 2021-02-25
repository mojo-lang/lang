package lang

import "errors"

func sameName(attribute *Attribute, fullName string) bool {
	pkg, name := ParseIdentifierName(fullName)
	if len(pkg) > 0 {
		return attribute.Package == pkg && attribute.Name == name
	} else {
		return attribute.Name == name
	}
}

func GetBoolAttribute(attributes []*Attribute, name string) (bool, error) {
	for _, attribute := range attributes {
		if sameName(attribute, name) {
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
		if sameName(attribute, name) {
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

	attributes = append(attributes, &Attribute{
		Name: name,
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
	for _, attribute := range attributes {
		if sameName(attribute, name) {
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
		if sameName(attribute, name) {
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

	attributes = append(attributes, &Attribute{
		Name: name,
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
	for _, attribute := range attributes {
		if sameName(attribute, name) {
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
		if sameName(attribute, name) {
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

	attributes = append(attributes, &Attribute{
		Name: name,
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
		if sameName(attribute, name) {
			continue
		}
		newAttributes = append(newAttributes, attribute)
	}

	return newAttributes
}
