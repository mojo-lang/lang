package lang

import "errors"

func (m *FunctionDecl) SetStartPosition(position *Position) {
	if m != nil {
		m.StartPosition = PatchPosition(m.StartPosition, position)
	}
}

func (m *FunctionDecl) SetEndPosition(position *Position) {
	if m != nil {
		m.EndPosition = PatchPosition(m.EndPosition, position)
	}
}

func (m *FunctionDecl) GetAttributeArguments(name string) ([]*Argument, error) {
	if m != nil {
		return GetAttributeArguments(m.Attributes, name)
	}
	return nil, errors.New("FunctionDecl is nil")
}

func (m *FunctionDecl) GetAttributeArgument(name string) (*Argument, error) {
	if m != nil {
		return GetAttributeArgument(m.Attributes, name)
	}
	return nil, errors.New("FunctionDecl is nil")
}

func (m *FunctionDecl) HasAttribute(name string) bool {
	if m != nil {
		return HasAttribute(m.Attributes, name)
	}
	return false
}

func (m *FunctionDecl) GetAttribute(name string) *Attribute {
	if m != nil {
		return GetAttribute(m.Attributes, name)
	}
	return nil
}

func (m *FunctionDecl) GetBoolAttribute(name string) (bool, error) {
	argument, err := m.GetAttributeArgument(name)
	if err != nil {
		return false, err
	}

	return argument.GetBool()
}

func (m *FunctionDecl) SetBoolAttribute(name string, value bool) *Attribute {
	if m != nil {
		m.Attributes = SetBoolAttribute(m.Attributes, name, value)
		return m.Attributes[len(m.Attributes)-1]
	}
	return nil
}

func (m *FunctionDecl) SetImplicitBoolAttribute(name string, value bool) *Attribute {
	if m != nil {
		return m.SetBoolAttribute(name, value).SetImplicit(true)
	}
	return nil
}

func (m *FunctionDecl) GetIntegerAttribute(name string) (int64, error) {
	if argument, err := m.GetAttributeArgument(name); err != nil {
		return 0, err
	} else {
		return argument.GetInteger()
	}
}

func (m *FunctionDecl) SetIntegerAttribute(name string, value int64) *Attribute {
	if m != nil {
		m.Attributes = SetIntegerAttribute(m.Attributes, name, value)
		return m.Attributes[len(m.Attributes)-1]
	}
	return nil
}

func (m *FunctionDecl) SetImplicitIntegerAttribute(name string, value int64) *Attribute {
	if m != nil {
		return m.SetIntegerAttribute(name, value).SetImplicit(true)
	}
	return nil
}

func (m *FunctionDecl) GetStringAttribute(name string) (string, error) {
	if argument, err := m.GetAttributeArgument(name); err != nil {
		return "", err
	} else {
		return argument.GetString()
	}
}

func (m *FunctionDecl) SetStringAttribute(name string, value string) *Attribute {
	if m != nil {
		m.Attributes = SetStringAttribute(m.Attributes, name, value)
		return m.Attributes[len(m.Attributes)-1]
	}
	return nil
}

func (m *FunctionDecl) SetImplicitStringAttribute(name string, value string) *Attribute {
	if m != nil {
		return m.SetStringAttribute(name, value).SetImplicit(true)
	}
	return nil
}

func (m *FunctionDecl) RemoveAttribute(name string) {
	if m != nil {
		m.Attributes = RemoveAttribute(m.Attributes, name)
	}
}

func (m *FunctionDecl) GetResource() string {
	if m != nil {
		//if m.GetStringAttribute()
	}
	return ""
}
