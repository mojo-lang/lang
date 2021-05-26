package lang

import "strings"

func GetFullName(pkg string, enclosingNames []string, decl string) string {
	fullName := decl
	if len(enclosingNames) > 0 {
		fullName = strings.Join([]string{strings.Join(enclosingNames, "."), fullName}, ".")
	}

	if len(pkg) > 0 {
		fullName = strings.Join([]string{pkg, fullName}, ".")
	}
	return fullName
}

/// get enclosing names from enclosing type in `NominalType` or `Declaration`
func GetEnclosingNames(t *NominalType) []string {
	var enclosingNames []string
	enclosing := t
	for enclosing != nil {
		enclosingNames = append(enclosingNames, enclosing.Name)
		enclosing = enclosing.EnclosingType
	}
	return enclosingNames
}
