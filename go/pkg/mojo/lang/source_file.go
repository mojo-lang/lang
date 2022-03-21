package lang

import "strings"

func (m *SourceFile) IsGenericInstantiated() bool {
    return strings.Contains(m.GetName(), "<")
}

func (m *SourceFile) SetScope(scope *Scope) {
    if m != nil {
        m.Scope = scope
    }
}

func (m *SourceFile) GetTypeDeclarations() []*TypeDeclaration {
    var decls []*TypeDeclaration
    for _, statement := range m.GetStatements() {
        if decl := statement.GetDeclaration(); decl != nil {
            switch decl.Declaration.(type) {
            case *Declaration_StructDecl, *Declaration_EnumDecl, *Declaration_InterfaceDecl, *Declaration_TypeAliasDecl:
                decls = append(decls, NewTypeDeclarationFromDeclaration(decl))
            }
        }
    }
    return decls
}

func (m *SourceFile) GetStructDecls() []*StructDecl {
    var decls []*StructDecl
    for _, statement := range m.GetStatements() {
        if decl := statement.GetDeclaration().GetStructDecl(); decl != nil {
            decls = append(decls, decl)
        }
    }
    return decls
}

func (m *SourceFile) GetTypeAliasDecls() []*TypeAliasDecl {
    var decls []*TypeAliasDecl
    for _, statement := range m.GetStatements() {
        if decl := statement.GetDeclaration().GetTypeAliasDecl(); decl != nil {
            decls = append(decls, decl)
        }
    }
    return decls
}
