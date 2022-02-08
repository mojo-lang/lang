package lang

func NewEntityRelationSet() *EntityRelationSet {
	return &EntityRelationSet{
		Nodes:         make(map[string]*EntityNode),
		Edges:         make(map[string]*EntityEdge),
		NodeEdges: make(map[string]*EntityEdges),
		NodeRelations:   make(map[string]*EntityRelations),
	}
}

func (m *EntityRelationSet) GetNode(name string) *EntityNode {
	if m != nil {
		return m.Nodes[name]
	}
	return nil
}
