package lang

import (
    "github.com/mojo-lang/core/go/pkg/mojo/core"
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

func firstNonZero(v *core.Version) int {
    if v.Major > 0 || v.Level == 1 {
        return 0
    } else if v.Minor > 0 || v.Level == 2 {
        return 1
    } else if v.Major > 0 {
        return 2
    }
    return 2
}

func parseWildcard(v string) *core.VersionRange {
    return nil
}

func parseCaret(v string) *core.VersionRange {
    vr := &core.VersionRange{}

    version, _ := core.ParseVersion(v)

    vr.Min = version
    vr.Max = &core.Version{
        Major: vr.Min.Major,
        Minor: vr.Min.Minor,
        Patch: vr.Min.Patch,
    }

    if i := firstNonZero(version); i >= 0 && i < 3 {
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
