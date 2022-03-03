package lang

type StartPositionGetter interface {
	GetStartPosition() *Position
}

type StartPositionSetter interface {
	SetStartPosition(position *Position)
}

func SetStartPosition(object interface{}, position *Position) {
	if setter, ok := object.(StartPositionSetter); ok {
		setter.SetStartPosition(position)
	}
}

type EndPositionGetter interface {
	GetEndPosition() *Position
}

type EndPositionSetter interface {
	SetEndPosition(position *Position)
}

func SetEndPosition(object interface{}, position *Position) {
	if setter, ok := object.(EndPositionSetter); ok {
		setter.SetEndPosition(position)
	}
}

// PatchPosition patch the set value to target
func PatchPosition(target *Position, patch *Position) *Position {
	if target == nil {
		return patch
	}
	if patch == nil {
		return target
	}

	// patch position has set
	if len(patch.Filename) > 0 || patch.Line > 0 || patch.Column > 0 || patch.Offset > 0 {
		target.Filename = patch.Filename
		target.Line = patch.Line
		target.Column = patch.Column
		target.Offset = patch.Offset
	}

	if len(patch.LeadingComments) > 0 {
		target.LeadingComments = append(target.LeadingComments, patch.LeadingComments...)
	}
	if len(patch.TailingComments) > 0 {
		target.TailingComments = append(target.TailingComments, patch.TailingComments...)
	}

	return target
}
