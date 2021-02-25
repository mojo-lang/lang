package lang

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"strconv"
	"strings"
)

func NewPackageRequirementVersion(str string) *Package_Requirement_Version {
	v := &Package_Requirement_Version{}
	v.Type = Package_Requirement_Version_TYPE_CARET

	if strings.Contains(str, "*") {
		v.Type = Package_Requirement_Version_TYPE_WILDCARD
		v.Range = parseWildcard(str)
	} else {
		switch str[0] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			v.Type = Package_Requirement_Version_TYPE_CARET
			v.Range = parseCaret(str)
		case '^':
			v.Type = Package_Requirement_Version_TYPE_CARET
			v.Range = parseCaret(str[1:])
		case '~':
			v.Type = Package_Requirement_Version_TYPE_TILDE
			v.Range = parseTilde(str[1:])
		case '>', '<', '=':
			v.Type = Package_Requirement_Version_TYPE_COMPARISON
			v.Range = parseComparison(str)
		}
	}

	return v
}

type version []string

func newVersion(v string) version {
	v = strings.TrimSpace(v)
	return strings.Split(v, ".")
}

func (v version) firstNonZero() int {
	for i, segment := range v {
		if segment > "0" {
			return i
		}
	}
	return len(v) - 1
}

func (v version) toVersion() *core.Version {
	version := &core.Version{}
	values := []*uint64{&version.Major, &version.Minor, &version.Patch}
	for i, segment := range v {
		value, err := strconv.Atoi(segment)
		if err != nil {
		}
		*values[i] = uint64(value)
	}
	return version
}

func parseWildcard(v string) *core.VersionRange {
	return nil
}

func parseCaret(v string) *core.VersionRange {
	vr := &core.VersionRange{}
	version := newVersion(v)
	vr.Min = version.toVersion()
	vr.Max = &core.Version{
		Major: vr.Min.Major,
		Minor: vr.Min.Minor,
		Patch: vr.Min.Patch,
	}

	if i := version.firstNonZero(); i >= 0 && i < 3 {
		values := []*uint64{&vr.Max.Major, &vr.Max.Minor, &vr.Max.Patch}
		*values[i]++
		for j := i + 1; j < 3; j++ {
			*values[j] = 0
		}
		vr.MaxExcluded = true
	}
	return vr
}

func parseTilde(v string) *core.VersionRange {
	return nil
}

func parseComparison(version string) *core.VersionRange {
	segments := strings.Split(version, ",")
	for _, segment := range segments {
		segment = strings.TrimSpace(segment)
		if strings.HasPrefix(segment, ">=") {
		} else if strings.HasPrefix(segment, ">") {
		} else if strings.HasPrefix(segment, "<=") {
		} else if strings.HasPrefix(segment, "<") {
		}
	}

	return nil
}
