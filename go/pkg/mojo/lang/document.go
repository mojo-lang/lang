package lang

import (
    "bytes"
    "github.com/mojo-lang/core/go/pkg/logs"
    "github.com/mojo-lang/document/go/pkg/markdown"
)

type DocumentGetter interface {
    GetDocument() *Document
}

func (m *Document) SetStartPosition(position *Position) {
    if m != nil {
        m.StartPosition = PatchPosition(m.StartPosition, position)
    }
}

func (m *Document) SetEndPosition(position *Position) {
    if m != nil {
        m.EndPosition = PatchPosition(m.EndPosition, position)
    }
}

func (m *Document) GetContent() string {
    if m != nil {
        content := bytes.Buffer{}
        for i, line := range m.Lines {
            if content.Len() == 0 && len(line.Text) == 0 {
                continue
            }
            content.WriteString(line.Text)

            if i < len(m.Lines)-1 {
                content.WriteString("\n")
            }
        }
        return content.String()
    }
    return ""
}

func (m *Document) Parse() *Document {
    if m != nil {
        mk := markdown.New()
        doc, err := mk.Parse(m.GetContent())
        if err != nil {
            logs.Warnw("failed to parse the document to markdown")
        } else {
            m.Structured = doc
        }
    }
    return m
}
