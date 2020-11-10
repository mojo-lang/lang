///
type Scope {
    attribute_decls: {String: AttributeDecl @reference} @1 
    const_decls:     {String: ConstDecl @reference} @2
    var_decls:       {String: VarDecl reference} @3

    data_decls:      {String: DataDecl @reference} @4
    struct_decls:    {String: StructDecl @reference} @5
    enum_decls:      {String: EnumDecl @reference} @6
    interface_decls: {String: InterfaceDecl @reference} @7
    type_alias_decls:{String: TypeAliasDecl @reference} @8

    func_decls:      {String: FuncDecl @reference} @10
}
