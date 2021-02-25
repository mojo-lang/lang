package lang

import "strings"

func ParseIdentifierName(fullName string) (string, string) {
	index := strings.LastIndex(fullName, ".")
	if index > 0 {
		return fullName[0:index], fullName[index+1:]
	} else {
		return "", fullName
	}
}
