///
///
///
type SourceFile {
    name: String @1
    full_name: String @2
    package: Package @3

    imports: [Import]

    // declarations
    //variables: [Variable]
    constants: [Constant]
    attributes: [Attribute]

    functions: [Function]

    data_types: [DataType]
    struct_types: [StructType]
    enum_types: [EnumType]
    service_types: [ServiceTyp]

    // statements except declarations
    statements: [Term]

    source_code_info: SourceCodeInfo
}

/// Encapsulates information about the original source file from which a
/// SourceFile was generated.
type SourceCodeInfo {
    type Location {
        /// Identifies which part of the FileDescriptorProto was defined at this
        /// location.
        ///
        /// Each element is a field number or an index.  They form a path from
        /// the root FileDescriptorProto to the place where the definition.  For
        /// example, this path:
        ///   [ 4, 3, 2, 7, 1 ]
        /// refers to:
        ///   file.message_type(3)  // 4, 3
        ///       .field(7)         // 2, 7
        ///       .name()           // 1
        /// This is because FileDescriptorProto.message_type has field number 4:
        ///   repeated DescriptorProto message_type = 4;
        /// and DescriptorProto.field has field number 2:
        ///   repeated FieldDescriptorProto field = 2;
        /// and FieldDescriptorProto.name has field number 1:
        ///   optional string name = 1;
        ///
        /// Thus, the above path gives the location of a field name.  If we removed
        /// the last element:
        ///   [ 4, 3, 2, 7 ]
        /// this path refers to the whole field declaration (from the beginning
        /// of the label to the terminating semicolon).
        path: [Int32] @1

        /// Always has exactly three or four elements: start line, start column,
        /// end line (optional, otherwise assumed same as start line), end column.
        /// These are packed into a single field for efficiency.  Note that line
        /// and column numbers are zero-based -- typically you will want to add
        /// 1 to each before displaying to a user.
        span: PosSpan
    }

    /// A Location identifies a piece of source code in a .mojo file which
    /// corresponds to a particular definition.  This information is intended
    /// to be useful to IDEs, code indexers, documentation generators, and similar
    /// tools.
    ///
    /// For example, say we have a file like:
    ///   message Foo {
    ///     optional string foo = 1;
    ///   }
    /// Let's look at just the field definition:
    ///   optional string foo = 1;
    ///   ^       ^^     ^^  ^  ^^^
    ///   a       bc     de  f  ghi
    /// We have the following locations:
    ///   span   path               represents
    ///   [a,i)  [ 4, 0, 2, 0 ]     The whole field definition.
    ///   [a,b)  [ 4, 0, 2, 0, 4 ]  The label (optional).
    ///   [c,d)  [ 4, 0, 2, 0, 5 ]  The type (string).
    ///   [e,f)  [ 4, 0, 2, 0, 1 ]  The name (foo).
    ///   [g,h)  [ 4, 0, 2, 0, 3 ]  The number (1).
    ///
    /// Notes:
    /// - A location may refer to a repeated field itself (i.e. not to any
    ///   particular index within it).  This is used whenever a set of elements are
    ///   logically enclosed in a single code segment.  For example, an entire
    ///   extend block (possibly containing multiple extension definitions) will
    ///   have an outer location whose path refers to the "extensions" repeated
    ///   field without an index.
    /// - Multiple locations may have the same path.  This happens when a single
    ///   logical declaration is spread out across multiple places.  The most
    ///   obvious example is the "extend" block again -- there may be multiple
    ///   extend blocks in the same scope, each of which will have the same path.
    /// - A location's span is not always a subset of its parent's span.  For
    ///   example, the "extendee" of an extension declaration appears at the
    ///   beginning of the "extend" block and is shared by all extensions within
    ///   the block.
    /// - Just because a location's span is a subset of some other location's span
    ///   does not mean that it is a descendent.  For example, a "group" defines
    ///   both a type and a field in a single declaration.  Thus, the locations
    ///   corresponding to the type and field and their components will overlap.
    /// - Code which tries to interpret locations should probably be designed to
    ///   ignore those that it doesn't understand, as more types of locations could
    ///   be recorded in the future.
    locations: [Location]
}