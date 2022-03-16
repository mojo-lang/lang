package lang

import (
	"errors"
	"reflect"
)

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

func NewTypeAliasDeclStatement(decl *TypeAliasDecl) *Statement {
	return &Statement{
		Statement: &Statement_Declaration{
			Declaration: &Declaration{
				Declaration: &Declaration_TypeAliasDecl{
					TypeAliasDecl: decl,
				},
			},
		},
	}
}

func NewFunctionDeclStatement(decl *FunctionDecl) *Statement {
	return &Statement{
		Statement: &Statement_Declaration{
			Declaration: &Declaration{
				Declaration: &Declaration_FunctionDecl{
					FunctionDecl: decl,
				},
			},
		},
	}
}

func NewVariableDeclStatement(decl *VariableDecl) *Statement {
	return &Statement{
		Statement: &Statement_Declaration{
			Declaration: &Declaration{
				Declaration: &Declaration_VariableDecl{
					VariableDecl: decl,
				},
			},
		},
	}
}

func NewAttributeDeclStatement(decl *AttributeDecl) *Statement {
	return &Statement{
		Statement: &Statement_Declaration{
			Declaration: &Declaration{
				Declaration: &Declaration_AttributeDecl{
					AttributeDecl: decl,
				},
			},
		},
	}
}

func (m *Statement) IsUnion() {
}

func (m *Statement) GetStartPosition() *Position {
	return GetUnionPosition(m, 0)
}

func (m *Statement) GetEndPosition() *Position {
	return GetUnionPosition(m, 1)
}

func (m *Statement) SetStartPosition(position *Position) {
	SetUnionPosition(m, 0, position)
}

func (m *Statement) SetEndPosition(position *Position) {
	SetUnionPosition(m, 1, position)
}

func (m *Statement) MergeComment(comment *Comment) (bool, error) {
	if m != nil {
		value := reflect.ValueOf(m.Statement)
		value = reflect.Indirect(value).Field(0)
		if merger, ok := value.Interface().(CommentMerger); ok {
			return merger.MergeComment(comment)
		} else {
			// error
		}
	}

	return false, errors.New("nil Declaration")
}
