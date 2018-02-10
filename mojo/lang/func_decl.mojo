type FuncDecl : Decl {
    ///
    name: String @10
    
    ///
    generic_params: [GenericParam] @11
    
    ///
    type: FuncType @12

    ///
    attributes: [Attribute] @13

    ///
    body: BlockStmt @14
}

/// A function parameter declaration.
class ParamDecl : public VarDecl {
}

/// \brief Base class for function-like declarations.
class AbstractFunctionDecl : public ValueDecl, public GenericContext

/// FuncDecl - 'func' declaration.
class FuncDecl final : public AbstractFunctionDecl,


