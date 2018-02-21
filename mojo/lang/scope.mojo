///
type Scope {
    attribute_decls: {String: AttributeDecl} @1 @linked
    const_decls:     {String: ConstDecl} @2 @linked
    var_decls:       {String: VarDecl} @3

    data_decls:      {String: DataDecl} @4 @linked
    struct_decls:    {String: StructDecl} @5 @linked
    enum_decls:      {String: EnumDecl} @6 @linked
    interface_decls: {String: InterfaceDecl} @7 @linked
    type_alias_decls:{String: TypeAliasDecl} @8 @linked

    func_decls:      {String: FuncDecl} @9 @linked
}
