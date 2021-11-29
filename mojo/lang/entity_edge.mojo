
type EntityEdge {
    /// the id of the edge, will be generated from the node names
    id: String @1
    

    name: String @2

    /// Whether the Decl represents something directly written in source or
    /// it was implicitly generated by the type-checker.
    implicit: Bool @3

    from: EntityNode @5
    to: EntityNode @6

    inverse: Bool @10
    multiple: Bool @11
}

type EntityEdges = [EntityEdge]
