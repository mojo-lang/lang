//
//
//

/// There are 3 main classes of nodes: Expressions and type nodes,
/// statement nodes, and declaration nodes. The node names usually
/// match the corresponding Go spec production names to which they
/// correspond. The node fields correspond to the individual parts
/// of the respective productions.
///
/// All nodes contain position information marking the beginning of
/// the corresponding source text segment; it is accessible via the
/// Pos accessor method. Nodes may contain additional position info
/// for language constructs where comments may be found between parts
/// of the construct (typically any larger, parenthesized subpart).
/// That position information is needed to properly position comments
/// when printing the construct.
///
type Node = Expression | Declaration | Statement
