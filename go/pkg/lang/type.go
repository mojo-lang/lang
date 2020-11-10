package lang

import "errors"

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
