package lang

import (
	"strings"

	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
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

func GetTypePackageName(fullName string) string {
	if len(fullName) == 0 {
		return ""
	}

	segments := strings.Split(fullName, ".")
	for i, segment := range segments {
		if IsTypeName(segment) {
			if i == 0 {
				return ""
			}
			return strings.Join(segments[:i], ".")
		}
	}
	return fullName
}

func GetTypeTypeName(fullName string) string {
	if len(fullName) == 0 {
		return ""
	}

	segments := strings.Split(fullName, ".")
	for i, segment := range segments {
		if IsTypeName(segment) {
			return strings.Join(segments[i:], ".")
		}
	}
	return ""
}

///
func IsTypeName(name string) bool {
	if len(name) == 0 {
		return false
	}

	if !strings.Contains(name, ".") {
		return name[0] >= 'A' && name[0] <= 'Z'
	} else {
		segments := strings.Split(name, ".")
		n := segments[len(segments)-1]
		if len(n) > 0 {
			return n[0] >= 'A' && n[0] <= 'Z'
		}
	}
	return false
}

/// get enclosing names from enclosing type in `NominalType` or `Declaration`
func GetEnclosingNames(t *NominalType) []string {
	var enclosingNames []string
	enclosing := t
	for enclosing != nil {
		enclosingNames = append(enclosingNames, enclosing.Name)
		enclosing = enclosing.EnclosingType
	}
	if len(enclosingNames) > 1 {
		last := len(enclosingNames) - 1
		for i := 0; i < len(enclosingNames)/2; i++ {
			enclosingNames[i], enclosingNames[last-i] = enclosingNames[last-i], enclosingNames[i]
		}
	}

	return enclosingNames
}

// convert full type name to file name
func TypeNameToFileName(fullName string) string {
	segments := strings.Split(fullName, ".")
	for i, segment := range segments {
		if i < len(segments) - 1 {
			segments[i] = strcase.ToKebab(segment)
		} else {
			if len(segment) > 0 && segment[0] >= 'A' && segment[0] <= 'Z' {
				segments[i] = strcase.ToSnake(segment)
			}
		}
	}
	return strings.Join(segments, "/")
}
