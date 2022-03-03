package lang

import (
	"errors"
)

func (m *ValueDecl) SetStartPosition(position *Position) {
	if m != nil {
		m.StartPosition = PatchPosition(m.StartPosition, position)
	}
}

func (m *ValueDecl) SetEndPosition(position *Position) {
	if m != nil {
		m.EndPosition = PatchPosition(m.EndPosition, position)
	}
}

func (m *ValueDecl) GetAttributeArguments(name string) ([]*Argument, error) {
	if m != nil {
		arguments, err := GetAttributeArguments(m.Attributes, name)
		if err != nil && m.Type != nil {
			return GetAttributeArguments(m.Type.Attributes, name)
		}
		return arguments, err
	}

	return nil, errors.New("ValueDecl is nil")
}

func (m *ValueDecl) GetAttributeArgument(name string) (*Argument, error) {
	if m != nil {
		argument, err := GetAttributeArgument(m.Attributes, name)
		if err != nil && m.Type != nil {
			return GetAttributeArgument(m.Type.Attributes, name)
		}
		return argument, err
	}

	return nil, errors.New("ValueDecl is nil")
}

func (m *ValueDecl) GetAttribute(name string) *Attribute {
	if m != nil {
		attr := GetAttribute(m.Attributes, name)
		if attr == nil && m.Type != nil {
			attr = GetAttribute(m.Type.Attributes, name)
		}
		return attr
	}
	return nil
}

func (m *ValueDecl) HasAttribute(name string) bool {
	if m != nil {
		attr := HasAttribute(m.Attributes, name)
		if !attr && m.Type != nil {
			attr = HasAttribute(m.Type.Attributes, name)
		}
		return attr
	}

	return false
}

func (m *ValueDecl) GetBoolAttribute(name string) (bool, error) {
	argument, err := m.GetAttributeArgument(name)
	if err != nil {
		return false, err
	}

	return argument.GetBool()
}

func (m *ValueDecl) SetBoolAttribute(name string, value bool) *Attribute {
	if m != nil && m.Type != nil {
		m.Type.Attributes = SetBoolAttribute(m.Type.Attributes, name, value)
		return m.Type.Attributes[len(m.Type.Attributes)-1]
	}
	return nil
}

func (m *ValueDecl) SetImplicitBoolAttribute(name string, value bool) *Attribute {
	if m != nil && m.Type != nil {
		return m.SetBoolAttribute(name, value).SetImplicit(true)
	}
	return nil
}

func (m *ValueDecl) GetIntegerAttribute(name string) (int64, error) {
	if argument, err := m.GetAttributeArgument(name); err != nil {
		return 0, err
	} else {
		return argument.GetInteger()
	}
}

func (m *ValueDecl) SetIntegerAttribute(name string, value int64) *Attribute {
	if m != nil && m.Type != nil {
		m.Type.Attributes = SetIntegerAttribute(m.Type.Attributes, name, value)
		return m.Type.Attributes[len(m.Type.Attributes)-1]
	}
	return nil
}

func (m *ValueDecl) SetImplicitIntegerAttribute(name string, value int64) *Attribute {
	if m != nil && m.Type != nil {
		return m.SetIntegerAttribute(name, value).SetImplicit(true)
	}
	return nil
}

func (m *ValueDecl) GetStringAttribute(name string) (string, error) {
	argument, err := m.GetAttributeArgument(name)
	if err != nil {
		return "", err
	}

	return argument.GetString()
}

func (m *ValueDecl) SetStringAttribute(name string, value string) *Attribute {
	if m != nil && m.Type != nil {
		m.Type.Attributes = SetStringAttribute(m.Type.Attributes, name, value)
		return m.Type.Attributes[len(m.Type.Attributes)-1]
	}
	return nil
}

func (m *ValueDecl) SetImplicitStringAttribute(name string, value string) *Attribute {
	if m != nil && m.Type != nil {
		return m.SetStringAttribute(name, value).SetImplicit(true)
	}
	return nil
}

func (m *ValueDecl) RemoveAttribute(name string) {
	if m != nil {
		m.Attributes = RemoveAttribute(m.Attributes, name)
		if m.Type != nil {
			m.Type.Attributes = RemoveAttribute(m.Type.Attributes, name)
		}
	}
}

func (m *ValueDecl) IsArrayType() bool {
	return m.GetType().IsArrayType()
}

func (m *ValueDecl) IsMapType() bool {
	return m.GetType().IsMapType()
}

func (m *ValueDecl) IsUnionType() bool {
	return m.GetType().IsUnionType()
}

func (m *ValueDecl) IsScalarType() bool {
	return m.GetType().IsScalar()
}

func (m *ValueDecl) IsEnum() bool {
	return m.GetType().IsEnum()
}
