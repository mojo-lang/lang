package lang

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"strings"
)

func NewGlobalPackage() *Package {
	return &Package{
		Summary: "Global",
		Scope:   NewScope(),
	}
}

func GetPackageName(fullName string) string {
	if len(fullName) == 0 {
		return ""
	}

	segments := strings.Split(fullName, ".")
	return segments[len(segments)-1]
}

func GetPackageParentName(fullName string) string {
	if len(fullName) == 0 {
		return ""
	}

	segments := strings.Split(fullName, ".")
	if len(segments) > 1 {
		return strings.Join(segments[:len(segments)-1], ".")
	}
	return ""
}

func GetPackageParentNames(fullName string) []string {
	segments := strings.Split(fullName, ".")
	var parents []string
	for i := 0; i < len(segments)-1; i++ {
		parents = append(parents, strings.Join(segments[:i+1], "."))
	}
	return parents
}

func (m *Package) ParentName() string {
	return GetPackageParentName(m.FullName)
}

func (m *Package) ParentNames() []string {
	return GetPackageParentNames(m.FullName)
}

func (m *Package) IsGlobal() bool {
	return len(m.Name) == 0 && m.Summary == "Global" && m.Scope != nil
}

func (m *Package) SetScope(scope *Scope) {
	if m != nil {
		m.Scope = scope
	}
}

func (m *Package) GetAllPackage() map[string]*Package {
	packages := make(map[string]*Package)
	packages[m.FullName] = m

	for _, pkg := range m.Children {
		packages[pkg.FullName] = pkg
	}
	return packages
}

func (m *Package) GetAllPackageCount() int {
	total := 1
	for _, pkg := range m.Children {
		total += pkg.GetAllPackageCount()
	}
	return total
}

func (m *Package) SetExtraBool(key string, value bool) {
	if m.ExtraInfo == nil {
		m.ExtraInfo = &core.Object{}
	}
	m.ExtraInfo.SetBool(key, value)
}

func (m *Package) SetExtraString(key string, value string) {
	if m.ExtraInfo == nil {
		m.ExtraInfo = &core.Object{}
	}
	m.ExtraInfo.SetString(key, value)
}

func (m *Package) GetExtraBool(key string) bool {
	if m.ExtraInfo == nil {
		return false
	}
	return m.ExtraInfo.GetBool(key)
}

func (m *Package) GetExtraString(key string) string {
	if m.ExtraInfo == nil {
		return ""
	}
	return m.ExtraInfo.GetString(key)
}

func (m *Package) GetSourceFile(name string) *SourceFile {
	if m != nil {
		for _, file := range m.SourceFiles {
			if file.FullName == name {
				return file
			}
		}
	}
	return nil
}

func (m *Package) GetOrganization() string {
	if m == nil {
		return ""
	}

	// first find only has organization
	for _, author := range m.Authors {
		if len(author.Author) == 0 && len(author.Organization) > 0 {
			return author.Organization
		}
	}

	for _, author := range m.Authors {
		if len(author.Organization) > 0 {
			return author.Organization
		}
	}

	return ""
}

//func (m *Package) Merge(pkg *Package)  {
//	//m.ResolvedDependencies = pkg.ResolvedDependencies
//	m.Children = pkg.Children
//	m.Scope = pkg.Scope
//	m.SourceFiles = pkg.SourceFiles
//}