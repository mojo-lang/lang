// A File node represents a Mojo source file.
//
// The Comments list contains all comments in the source file in order of
// appearance, including the comments that are pointed to from other nodes
// via Doc and Comment fields.
//
// For correct printing of source code containing comments (using packages
// go/format and go/printer), special care must be taken to update comments
// when a File's syntax tree is modified: For printing, comments are interspersed
// between tokens based on their position. If syntax tree nodes are
// removed or moved, relevant comments in their vicinity must also be removed
// (from the File.Comments list) or moved accordingly (by updating their
// positions). A CommentMap may be used to facilitate some of these operations.
//
// Whether and how a comment is associated with a node depends on the
// interpretation of the syntax tree by the manipulating program: Except for Doc
// and Comment comments directly associated with nodes, the remaining comments
// are "free-floating" (see also issues #18593, #20744).
//
type SourceFile {
    name:       String @1

    full_name:  String @2

    url: Url @3 // file://xx

    /// 循环依赖文件名称 'circle_dependency'
    attributes: [Attribute]  @4

    package_name:  String  @10

    imports:    [ImportDecl] @11 //< imports in this file

    statements: [Statement]  @12

    // resolved identifiers not in this source file
    resolved_identifiers: [Identifier] @16

    unresolved_identifiers:  [Identifier] @17 //< unresolved identifiers in this file

    dependencies: [SourceFile @reference] @21 //<

    scope:       Scope @30 //< package scope (this file only)
}