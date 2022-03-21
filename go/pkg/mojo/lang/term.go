package lang

const (
    TermTypeKeyword = "keyword"
    TermTypeName    = "name"
    TermTypeStart   = "start"
    TermTypeEnd     = "end"
    TermTypeSymbol  = "symbol"
)

func (m *Term) SetStartPosition(position *Position) {
    if m != nil {
        m.StartPosition = PatchPosition(m.StartPosition, position)
    }
}

func (m *Term) SetEndPosition(position *Position) {
    if m != nil {
        m.EndPosition = PatchPosition(m.EndPosition, position)
    }
}

func NewSymbolTerm(position *Position, termType string, value string) *Term {
    if position != nil {
        endPosition := position
        if len(value) > 0 {
            endPosition = &Position{
                Line:   position.Line,
                Column: position.Column + int64(len(value)),
            }
        }
        return &Term{StartPosition: position,
            EndPosition: endPosition,
            Type:        termType,
            Value:       value,
        }
    } else if len(termType) > 0 && len(value) > 0 {
        return &Term{Type: termType, Value: value}
    }

    return nil
}
