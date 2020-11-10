type ImportDecl : Decl {
    type Identifier {
         name : String @1
         alias : String @2
    }

    import_package : String @10
    import_package_alias : String @11

    identifiers : [Identifier] @12
}