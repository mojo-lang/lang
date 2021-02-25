///
///
///
type InterfaceType {
    start_position: Position @1 // identifier position

    end_position: Position @2 // identifier position

    ///
    methods: [FunctionDecl] @10

    ///
    inherits: [NominalType] @11
}