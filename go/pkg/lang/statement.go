package lang

func NewExpressionStatement(expression *Expression) *Statement {
	return &Statement{
		Statement: &Statement_Expression{Expression: expression},
	}
}

func NewDeclarationStatement(decl *Declaration) *Statement {
	return &Statement{
		Statement: &Statement_Declaration{
			Declaration: decl,
		},
	}
}

func NewPackageDeclStatement(decl *PackageDecl) *Statement {
	return &Statement{
		Statement: &Statement_Declaration{
			Declaration: &Declaration{
				Declaration: &Declaration_PackageDecl{
					PackageDecl: decl,
				},
			},
		},
	}
}

func NewImportDeclStatement(decl *ImportDecl) *Statement {
	return &Statement{
		Statement: &Statement_Declaration{
			Declaration: &Declaration{
				Declaration: &Declaration_ImportDecl{
					ImportDecl: decl,
				},
			},
		},
	}
}

func NewEnumDeclStatement(decl *EnumDecl) *Statement {
	return &Statement{
		Statement: &Statement_Declaration{
			Declaration: &Declaration{
				Declaration: &Declaration_EnumDecl{
					EnumDecl: decl,
				},
			},
		},
	}
}

func NewStructDeclStatement(decl *StructDecl) *Statement {
	return &Statement{
		Statement: &Statement_Declaration{
			Declaration: &Declaration{
				Declaration: &Declaration_StructDecl{
					StructDecl: decl,
				},
			},
		},
	}
}

func NewInterfaceDeclStatement(decl *InterfaceDecl) *Statement {
	return &Statement{
		Statement: &Statement_Declaration{
			Declaration: &Declaration{
				Declaration: &Declaration_InterfaceDecl{
					InterfaceDecl: decl,
				},
			},
		},
	}
}

func NewTypeAliasDeclStatement(decl *NominalDecl) *Statement {
	return &Statement{
		Statement: &Statement_Declaration{
			Declaration: &Declaration{
				Declaration: &Declaration_NominalDecl{
					NominalDecl: decl,
				},
			},
		},
	}
}

func NewFunctionDeclStatement(decl *FuncDecl) *Statement {
	return &Statement{
		Statement: &Statement_Declaration{
			Declaration: &Declaration{
				Declaration: &Declaration_FuncDecl{
					FuncDecl: decl,
				},
			},
		},
	}
}
