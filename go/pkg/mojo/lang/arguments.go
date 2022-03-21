package lang

import "errors"

type Arguments []*Argument

func (a Arguments) ToObjectLiteralExpr() (*ObjectLiteralExpr, error) {
    object := &ObjectLiteralExpr{}
    for _, argument := range a {
        if len(argument.Label) == 0 {
            return nil, errors.New("failed to convert to object because of lack of label in argument")
        }

        object.Fields = append(object.Fields, &ObjectLiteralExpr_Field{
            StartPosition: argument.StartPosition,
            EndPosition:   argument.EndPosition,
            Name:          argument.Label,
            Value:         argument.Value,
        })
    }
    object.Implicit = true
    return object, nil
}
