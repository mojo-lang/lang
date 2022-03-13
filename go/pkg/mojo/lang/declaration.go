package lang

import (
    "errors"
    jsoniter "github.com/json-iterator/go"
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "reflect"
    "strings"
)

const (
    EnumDeclName      = "EnumDecl"
    StructDeclName    = "StructDecl"
    InterfaceDeclName = "InterfaceDecl"
)

var declInfos map[reflect.Type]StructJsonInfo

func init() {
    declInfos = GetStructJsonInfos(jsoniter.Config{},
        PackageDecl{},
        ImportDecl{},
        ConstantDecl{},
        EnumDecl{},
        StructDecl{},
        InterfaceDecl{},
        FunctionDecl{},
    )
}

func NewEnumDeclaration(decl *EnumDecl) *Declaration {
    return &Declaration{
        Declaration: &Declaration_EnumDecl{
            EnumDecl: decl,
        },
    }
}

func NewStructDeclaration(decl *StructDecl) *Declaration {
    return &Declaration{
        Declaration: &Declaration_StructDecl{
            StructDecl: decl,
        },
    }
}

func NewInterfaceDeclaration(decl *InterfaceDecl) *Declaration {
    return &Declaration{
        Declaration: &Declaration_InterfaceDecl{
            InterfaceDecl: decl,
        },
    }
}

func NewTypeAliasDeclaration(decl *TypeAliasDecl) *Declaration {
    return &Declaration{
        Declaration: &Declaration_TypeAliasDecl{
            TypeAliasDecl: decl,
        },
    }
}

func NewDeclarationFromTypeDeclaration(decl *TypeDeclaration) *Declaration {
    if decl != nil {
        switch decl.TypeDeclaration.(type) {
        case *TypeDeclaration_StructDecl:
            return &Declaration{
                Declaration: &Declaration_StructDecl{
                    StructDecl: decl.GetStructDecl(),
                },
            }
        case *TypeDeclaration_InterfaceDecl:
            return &Declaration{
                Declaration: &Declaration_InterfaceDecl{
                    InterfaceDecl: decl.GetInterfaceDecl(),
                },
            }
        case *TypeDeclaration_EnumDecl:
            return &Declaration{
                Declaration: &Declaration_EnumDecl{
                    EnumDecl: decl.GetEnumDecl(),
                },
            }
        case *TypeDeclaration_TypeAliasDecl:
            return &Declaration{
                Declaration: &Declaration_TypeAliasDecl{
                    TypeAliasDecl: decl.GetTypeAliasDecl(),
                },
            }
        case *TypeDeclaration_GenericParameter:
            return &Declaration{
                Declaration: &Declaration_GenericParameter{
                    GenericParameter: decl.GetGenericParameter(),
                },
            }
        default:
            return nil
        }
    }
    return nil
}

func NewGenericParameterDeclaration(parameter *GenericParameter) *Declaration {
    return &Declaration{
        Declaration: &Declaration_GenericParameter{
            GenericParameter: parameter,
        },
    }
}

func (m *Declaration) IsUnion() {
}

func (m *Declaration) GetStartPosition() *Position {
    return GetUnionPosition(m, 0)
}

func (m *Declaration) GetEndPosition() *Position {
    return GetUnionPosition(m, 1)
}

func (m *Declaration) SetStartPosition(position *Position) {
    SetUnionPosition(m, 0, position)
}

func (m *Declaration) SetEndPosition(position *Position) {
    SetUnionPosition(m, 1, position)
}

func (m *Declaration) GetName() string {
    if m != nil {
        return core.GetUnionField(m, "Name").String()
    }
    return ""
}

func (m *Declaration) GetEnclosingType() *NominalType {
    if m != nil {
        switch m.Declaration.(type) {
        case *Declaration_StructDecl:
            return m.GetStructDecl().EnclosingType
        case *Declaration_InterfaceDecl:
            return m.GetInterfaceDecl().EnclosingType
        case *Declaration_EnumDecl:
            return m.GetEnumDecl().EnclosingType
        case *Declaration_TypeAliasDecl:
            return m.GetTypeAliasDecl().EnclosingType
        case *Declaration_GenericParameter:
            return m.GetGenericParameter().EnclosingType
        default:
            return nil
        }
    }
    return nil
}

func (m *Declaration) GetPackageName() string {
    if m != nil {
        switch m.Declaration.(type) {
        case *Declaration_StructDecl:
            return m.GetStructDecl().PackageName
        case *Declaration_InterfaceDecl:
            return m.GetInterfaceDecl().PackageName
        case *Declaration_EnumDecl:
            return m.GetEnumDecl().PackageName
        case *Declaration_TypeAliasDecl:
            return m.GetTypeAliasDecl().PackageName
        case *Declaration_GenericParameter:
            return m.GetGenericParameter().PackageName
        default:
            return ""
        }
    }
    return ""
}

func (m *Declaration) SetPackageName(name string) {
    if m != nil {
        switch m.Declaration.(type) {
        case *Declaration_StructDecl:
            m.GetStructDecl().PackageName = name
        case *Declaration_InterfaceDecl:
            m.GetInterfaceDecl().PackageName = name
        case *Declaration_EnumDecl:
            m.GetEnumDecl().PackageName = name
        case *Declaration_TypeAliasDecl:
            m.GetTypeAliasDecl().PackageName = name
        case *Declaration_GenericParameter:
            m.GetGenericParameter().PackageName = name
        }
    }
}

func (m *Declaration) IsGeneric() bool {
    switch decl := m.Declaration.(type) {
    case *Declaration_TypeAliasDecl:
        return decl.TypeAliasDecl.IsGeneric()
    case *Declaration_StructDecl:
        return decl.StructDecl.IsGeneric()
    }
    return false
}

func (m *Declaration) IsInstantiatedGeneric() bool {
    return strings.ContainsAny(m.GetStructDecl().GetName(), "<>")
}

func (m *Declaration) MergeComment(comment *Comment) (bool, error) {
    if m != nil {
        value := reflect.ValueOf(m.Declaration)
        value = reflect.Indirect(value).Field(0)
        if merger, ok := value.Interface().(CommentMerger); ok {
            return merger.MergeComment(comment)
        } else {
            // error
        }
    }

    return false, errors.New("nil Declaration")
}
