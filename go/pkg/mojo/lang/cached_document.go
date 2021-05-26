package lang

func (m *CachedDocument) SetDocument(document *Document) {
	if m != nil {
		m.Value = document
	}
}
