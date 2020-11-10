///
///
///
type InterfaceType {
    ///
    name: String @1

    ///
    package: String @4 | Package @5 @reference

    ///
    methods: [FuncDecl] @10

    ///
    inherits: [NominalType] @11
}