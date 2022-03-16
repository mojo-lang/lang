package lang

import (
	"errors"
)

type CommentMerger interface {
	MergeComment(comment *Comment) (bool, error)
}

func NewComment(comment interface{}) *Comment {
	switch c := comment.(type) {
	case *Document:
		return NewDocumentComment(c)
	case *BlockComment:
		return NewBlockCommentComment(c)
	case *MultiLineComment:
		return NewMultiLineCommentComment(c)
	}
	return nil
}

func NewComments(comments ...interface{}) []*Comment {
	var cs []*Comment
	for _, comment := range comments {
		nc := NewComment(comment)
		if nc != nil {
			cs = append(cs, nc)
		}
	}
	return cs
}

func NewDocumentComment(document *Document) *Comment {
	return &Comment{
		Comment: &Comment_Document{
			Document: document,
		},
	}
}

func NewBlockCommentComment(comment *BlockComment) *Comment {
	return &Comment{
		Comment: &Comment_BlockComment{BlockComment: comment},
	}
}

func NewMultiLineCommentComment(comment *MultiLineComment) *Comment {
	return &Comment{Comment: &Comment_MultiLineComment{MultiLineComment: comment}}
}

func (m *Comment) IsUnion() {
}

func (m *Comment) GetStartPosition() *Position {
	if m != nil {
		switch comment := m.Comment.(type) {
		case *Comment_Document:
			return comment.Document.GetStartPosition()
		case *Comment_BlockComment:
			return comment.BlockComment.GetStartPosition()
		case *Comment_MultiLineComment:
			return comment.MultiLineComment.GetStartPosition()
		}
	}
	return nil
}

func (m *Comment) SetStartPosition(position *Position) {
	if m != nil {
		switch comment := m.Comment.(type) {
		case *Comment_Document:
			comment.Document.SetStartPosition(position)
		case *Comment_BlockComment:
			comment.BlockComment.SetStartPosition(position)
		case *Comment_MultiLineComment:
			comment.MultiLineComment.SetStartPosition(position)
		}
	}
}

func (m *Comment) GetEndPosition() *Position {
	if m != nil {
		switch comment := m.Comment.(type) {
		case *Comment_Document:
			return comment.Document.GetEndPosition()
		case *Comment_BlockComment:
			return comment.BlockComment.GetEndPosition()
		case *Comment_MultiLineComment:
			return comment.MultiLineComment.GetEndPosition()
		}
	}
	return nil
}

func (m *Comment) SetEndPosition(position *Position) {
	if m != nil {
		switch comment := m.Comment.(type) {
		case *Comment_Document:
			comment.Document.SetEndPosition(position)
		case *Comment_BlockComment:
			comment.BlockComment.SetEndPosition(position)
		case *Comment_MultiLineComment:
			comment.MultiLineComment.SetEndPosition(position)
		}
	}
}

func (m *Comment) IsFollowing() bool {
	return m != nil && m.GetMultiLineComment().IsFollowing()
}

type RelativeType int

const (
	RelativeLeading = iota
	RelativeInternal
	RelativeTailing
)

func (m *Comment) ComparePosition(start *Position, end *Position) RelativeType {
	if m.GetEndPosition().Compare(start) < 0 {
		return RelativeLeading
	} else if m.GetStartPosition().Compare(end) > 0 {
		return RelativeTailing
	} else {
		return RelativeInternal
	}
}

func (m *BlockComment) SetStartPosition(position *Position) {
	if m != nil {
		m.StartPosition = PatchPosition(m.StartPosition, position)
	}
}

func (m *BlockComment) SetEndPosition(position *Position) {
	if m != nil {
		m.EndPosition = PatchPosition(m.EndPosition, position)
	}
}

func (m *BlockComment) IsInLine() bool {
	return m != nil && m.StartPosition.GetLine() == m.EndPosition.GetLine()
}

func (m *LineComment) SetStartPosition(position *Position) {
	if m != nil {
		m.StartPosition = PatchPosition(m.StartPosition, position)
	}
}

func (m *LineComment) SetEndPosition(position *Position) {
	if m != nil {
		m.EndPosition = PatchPosition(m.EndPosition, position)
	}
}

func (m *MultiLineComment) SetStartPosition(position *Position) {
	if m != nil {
		m.StartPosition = PatchPosition(m.StartPosition, position)
	}
}

func (m *MultiLineComment) SetEndPosition(position *Position) {
	if m != nil {
		m.EndPosition = PatchPosition(m.EndPosition, position)
	}
}

func (m *MultiLineComment) IsFollowing() bool {
	return m != nil && len(m.Lines) > 0 && m.Lines[0].Following
}

func MergeCommentToTerms(comment *Comment, terms ...interface{}) (bool, error) {
	for _, term := range terms {
		if positioner, ok := term.(PositionGetter); ok {
			relative := comment.ComparePosition(positioner.GetStartPosition(), positioner.GetEndPosition())
			if relative == RelativeLeading {
				positioner.GetStartPosition().AppendLeadingComment(comment)
				return true, nil
			} else if relative == RelativeInternal {
				if merger, ok := term.(CommentMerger); ok {
					return merger.MergeComment(comment)
				}
				return false, errors.New("")
			} else {
				if comment.IsFollowing() && comment.GetStartPosition().GetLine() == positioner.GetEndPosition().GetLine() {
					positioner.GetEndPosition().AppendTailingComment(comment)
					return true, nil
				}
			}
		} else {
			return false, errors.New("")
		}
	}
	return false, errors.New("")
}
