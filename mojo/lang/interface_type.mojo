///
///
///
type InterfaceType {
    ///
    package: String @4 | Package @5 @reference

    ///
    methods: [FunctionDecl] @10

    ///
    inherits: [NominalType] @11
}