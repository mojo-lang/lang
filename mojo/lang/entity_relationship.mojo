
type EntityRelationship {
    nodes: {String: EntityNode} @1
    edges: {String: EntityEdge} @2

    node_from_edges: {String: EntityEdges} @10
    node_to_edges: {String: EntityEdges} @11
}
