package lang

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"reflect"
	"sort"
)

type PositionGetter interface {
	StartPositionGetter
	EndPositionGetter
}

type PositionSetter interface {
	StartPositionSetter
	EndPositionSetter
}

type PositionGetters []PositionGetter

func (p PositionGetters) Len() int {
	return len(p)
}
func (p PositionGetters) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p PositionGetters) Less(i, j int) bool {
	return p[i].GetStartPosition().Compare(p[j].GetStartPosition()) < 0
}

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

var fieldMap = map[int]string{
	0: "StartPosition", 1: "EndPosition",
}

func GetUnionPosition(union interface{}, index int) *Position {
	if union != nil {
		if pos, ok := core.GetUnionField(union, fieldMap[index]).Interface().(*Position); ok {
			return pos
		}
	}
	return nil
}

func SetUnionPosition(union interface{}, index int, position *Position) {
	if union != nil {
		value := reflect.ValueOf(union)
		for {
			if _, ok := value.Interface().(core.IsUnion); ok {
				value = reflect.Indirect(value).Field(0).Elem()
				value = reflect.Indirect(value).Field(0)
			} else {
				break
			}
		}

		value = reflect.Indirect(value).Field(index)
		if value.CanSet() {
			if value.IsNil() {
				value.Set(reflect.ValueOf(position))
			} else {
				if pos, ok := value.Interface().(*Position); ok {
					value.Set(reflect.ValueOf(PatchPosition(pos, position)))
				}
			}
		}
	}
}

func (m *Position) Compare(position *Position) int {
	if m == nil {
		if position == nil {
			return 0
		} else {
			return -1
		}
	} else {
		if position == nil {
			return 1
		} else {
			if m.Line == position.Line {
				if m.Column == position.Column {
					return 0
				} else if m.Column < position.Column {
					return -1
				} else {
					return 1
				}
			} else if m.Line < position.Line {
				return -1
			} else {
				return 1
			}
		}
	}
}

func (m *Position) AppendLeadingComment(comment interface{}) {
	if m != nil {
		switch c := comment.(type) {
		case *Comment:
			m.LeadingComments = append(m.LeadingComments, c)
		case *BlockComment:
			m.LeadingComments = append(m.LeadingComments, NewBlockCommentComment(c))
		case *MultiLineComment:
			m.LeadingComments = append(m.LeadingComments, NewMultiLineCommentComment(c))
		case *Document:
			m.LeadingComments = append(m.LeadingComments, NewDocumentComment(c))
		}
		m.LeadingComments = SortComments(m.LeadingComments)
	}
}

func (m *Position) AppendTailingComment(comment interface{}) {
	if m != nil {
		switch c := comment.(type) {
		case *Comment:
			m.TailingComments = append(m.TailingComments, c)
		case *BlockComment:
			m.TailingComments = append(m.TailingComments, NewBlockCommentComment(c))
		case *MultiLineComment:
			m.TailingComments = append(m.TailingComments, NewMultiLineCommentComment(c))
		case *Document:
			m.TailingComments = append(m.TailingComments, NewDocumentComment(c))
		}
		m.TailingComments = SortComments(m.TailingComments)
	}
}

func SortComments(comments []*Comment) []*Comment {
	var gs PositionGetters
	for _, comment := range comments {
		gs = append(gs, comment)
	}
	sort.Sort(&gs)

	var sortedComments []*Comment
	for _, g := range gs {
		sortedComments = append(sortedComments, g.(*Comment))
	}
	return sortedComments
}
