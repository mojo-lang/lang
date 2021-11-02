package lang

import "errors"

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

func (m *ValueDecl) SetBoolAttribute(name string, value bool) {
	if m != nil && m.Type != nil {
		m.Type.Attributes = SetBoolAttribute(m.Type.Attributes, name, value)
	}
}

func (m *ValueDecl) GetIntegerAttribute(name string) (int64, error) {
	argument, err := m.GetAttributeArgument(name)
	if err != nil {
		return 0, err
	}

	return argument.GetInteger()
}

func (m *ValueDecl) SetIntegerAttribute(name string, value int64) {
	if m != nil && m.Type != nil {
		m.Type.Attributes = SetIntegerAttribute(m.Type.Attributes, name, value)
	}
}

func (m *ValueDecl) GetStringAttribute(name string) (string, error) {
	argument, err := m.GetAttributeArgument(name)
	if err != nil {
		return "", err
	}

	return argument.GetString()
}

func (m *ValueDecl) SetStringAttribute(name string, value string) {
	if m != nil && m.Type != nil {
		m.Type.Attributes = SetStringAttribute(m.Type.Attributes, name, value)
	}
}

func (m *ValueDecl) RemoveAttribute(name string) {
	if m != nil {
		m.Attributes = RemoveAttribute(m.Attributes, name)
		if m.Type != nil {
			m.Type.Attributes = RemoveAttribute(m.Type.Attributes, name)
		}
	}
}
