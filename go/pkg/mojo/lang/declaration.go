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

func NewAttributeDeclaration(decl *AttributeDecl) *Declaration {
    return &Declaration{
        Declaration: &Declaration_AttributeDecl{
            AttributeDecl: decl,
        },
    }
}

func NewAttributeAliasDeclaration(decl *AttributeAliasDecl) *Declaration {
    return &Declaration{
        Declaration: &Declaration_AttributeAliasDecl{
            AttributeAliasDecl: decl,
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

func (x *Declaration) IsUnion() {
}

func (x *Declaration) GetStartPosition() *Position {
    return GetUnionPosition(x, StartPositionFieldName)
}

func (x *Declaration) GetEndPosition() *Position {
    return GetUnionPosition(x, EndPositionFieldName)
}

func (x *Declaration) SetStartPosition(position *Position) {
    SetUnionPosition(x, StartPositionFieldName, position)
}

func (x *Declaration) SetEndPosition(position *Position) {
    SetUnionPosition(x, EndPositionFieldName, position)
}

func (x *Declaration) GetName() string {
    if x != nil {
        return core.GetUnionField(x, "Name").String()
    }
    return ""
}

func (x *Declaration) GetEnclosingType() *NominalType {
    if x != nil {
        switch x.Declaration.(type) {
        case *Declaration_StructDecl:
            return x.GetStructDecl().EnclosingType
        case *Declaration_InterfaceDecl:
            return x.GetInterfaceDecl().EnclosingType
        case *Declaration_EnumDecl:
            return x.GetEnumDecl().EnclosingType
        case *Declaration_TypeAliasDecl:
            return x.GetTypeAliasDecl().EnclosingType
        case *Declaration_GenericParameter:
            return x.GetGenericParameter().EnclosingType
        default:
            return nil
        }
    }
    return nil
}

func (x *Declaration) GetPackageName() string {
    if x != nil {
        switch x.Declaration.(type) {
        case *Declaration_StructDecl:
            return x.GetStructDecl().PackageName
        case *Declaration_InterfaceDecl:
            return x.GetInterfaceDecl().PackageName
        case *Declaration_EnumDecl:
            return x.GetEnumDecl().PackageName
        case *Declaration_TypeAliasDecl:
            return x.GetTypeAliasDecl().PackageName
        case *Declaration_GenericParameter:
            return x.GetGenericParameter().PackageName
        default:
            return ""
        }
    }
    return ""
}

func (x *Declaration) SetPackageName(name string) {
    if x != nil {
        switch x.Declaration.(type) {
        case *Declaration_StructDecl:
            x.GetStructDecl().PackageName = name
        case *Declaration_InterfaceDecl:
            x.GetInterfaceDecl().PackageName = name
        case *Declaration_EnumDecl:
            x.GetEnumDecl().PackageName = name
        case *Declaration_TypeAliasDecl:
            x.GetTypeAliasDecl().PackageName = name
        case *Declaration_GenericParameter:
            x.GetGenericParameter().PackageName = name
        }
    }
}

func (x *Declaration) IsGeneric() bool {
    switch decl := x.Declaration.(type) {
    case *Declaration_TypeAliasDecl:
        return decl.TypeAliasDecl.IsGeneric()
    case *Declaration_StructDecl:
        return decl.StructDecl.IsGeneric()
    }
    return false
}

func (x *Declaration) IsInstantiatedGeneric() bool {
    return strings.ContainsAny(x.GetStructDecl().GetName(), "<>")
}

func (x *Declaration) MergeComment(comment *Comment) (bool, error) {
    if x != nil {
        value := reflect.ValueOf(x.Declaration)
        value = reflect.Indirect(value).Field(0)
        if merger, ok := value.Interface().(CommentMerger); ok {
            return merger.MergeComment(comment)
        } else {
            // error
        }
    }

    return false, errors.New("nil Declaration")
}
