
type Document {
    /// position of first character belonging to the Document
    start_position: Position @1

    /// position of first character immediately after the Document
    end_position: Position @2

    lines: [LineDocument] @15
}

type LineDocument {
    /// position of first character belonging to the LineDocument
    start_position: Position @1

    /// position of first character immediately after the LineDocument
    end_position: Position @2

    line: String @15
}
