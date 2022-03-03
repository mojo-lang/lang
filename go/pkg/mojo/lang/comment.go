package lang

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
