package lang

import (
	"github.com/iancoleman/strcase"
	"strings"
)

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

// convert full type name to file name
func TypeNameToFileName(fullName string) string {
	segments := strings.Split(fullName, ".")
	for i, segment := range segments {
		if len(segment) > 0 && segment[0] >= 'A' && segment[0] <= 'Z' {
			segments[i] = strcase.ToSnake(segment)
		}
	}
	return strings.Join(segments, "/")
}
