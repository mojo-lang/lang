package lang

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"reflect"
)

func (m *ObjectLiteralExpr) SetStartPosition(position *Position) {
	if m != nil {
		m.StartPosition = PatchPosition(m.StartPosition, position)
	}
}

func (m *ObjectLiteralExpr) SetEndPosition(position *Position) {
	if m != nil {
		m.EndPosition = PatchPosition(m.EndPosition, position)
	}
}

func (m *ObjectLiteralExpr) To(object interface{}) error {
	if m != nil && object != nil {
		ot := reflect.TypeOf(object)
		ov := reflect.ValueOf(object).Elem()
		for _, field := range m.Fields {
			fieldName := strcase.ToCamel(field.Name)
			if f, found := ot.FieldByName(fieldName); found {
				v := ov.FieldByIndex(f.Index)
				switch v.Type().Kind() {
				case reflect.String:
					v.SetString(field.Value.GetStringLiteralExpr().GetValue())
				case reflect.Struct:

				}
			}
		}
	}

	return nil
}

func (m *ObjectLiteralExpr) From(object interface{}) error {
	return nil
}
