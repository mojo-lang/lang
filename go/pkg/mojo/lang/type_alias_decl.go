package lang

import "errors"

const OriginalTypeAliasName = "original_type_alias_name"

func (x *TypeAliasDecl) SetStartPosition(position *Position) {
    if x != nil {
        x.StartPosition = PatchPosition(x.StartPosition, position)
    }
}

func (x *TypeAliasDecl) SetEndPosition(position *Position) {
    if x != nil {
        x.EndPosition = PatchPosition(x.EndPosition, position)
    }
}

func (x *TypeAliasDecl) MergeComment(comment *Comment) (bool, error) {
    if x != nil {
        return MergeCommentToTerms(comment, x.GetTerms())
    }

    return false, errors.New("nil NominalType")
}

func (x *TypeAliasDecl) GetTerms() []interface{} {
    if x != nil {
        var terms []interface{}

        for _, attribute := range x.Attributes {
            terms = append(terms, attribute)
        }

        terms = append(terms,
            NewSymbolTerm(x.KeywordPosition, TermTypeKeyword, KeywordType),
            NewSymbolTerm(x.NamePosition, TermTypeName, x.Name))

        for _, parameter := range x.GenericParameters {
            terms = append(terms, parameter)
        }

        terms = append(terms, x.Type)

        return terms
    }
    return nil
}

func (x *TypeAliasDecl) GetFullName() string {
    if x != nil {
        return GetFullName(x.PackageName, GetEnclosingNames(x.EnclosingType), x.Name)
    }
    return ""
}

func (x *TypeAliasDecl) SetScope(scope *Scope) {
    if x != nil {
        x.Scope = scope
    }
}

func (x *TypeAliasDecl) IsGeneric() bool {
    return x != nil && len(x.GenericParameters) > 0
}
