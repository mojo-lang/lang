package lang

import (
	"bytes"
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/document/go/pkg/markdown"
)

func (m *Document) GetContent() string {
	if m != nil {
		content := bytes.Buffer{}
		for i, line := range m.Lines {
			if content.Len() == 0 && len(line.Line) == 0 {
				continue
			}
			content.WriteString(line.Line)

			if i < len(m.Lines) - 1 {
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
