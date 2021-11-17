package lang

import (
	"fmt"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"strings"
)

func NewGlobalPackage() *Package {
	return &Package{
		Summary: "Global",
		Scope:   NewScope(),
	}
}

func GetPackageName(fullPkgName string) string {
	if len(fullPkgName) == 0 {
		return ""
	}

	segments := strings.Split(fullPkgName, ".")
	return segments[len(segments)-1]
}

func GetPackageParentName(fullPkgName string) string {
	if len(fullPkgName) == 0 {
		return ""
	}

	segments := strings.Split(fullPkgName, ".")
	if len(segments) > 1 {
		return strings.Join(segments[:len(segments)-1], ".")
	}
	return ""
}

func GetPackageParentNames(fullPkgName string) []string {
	segments := strings.Split(fullPkgName, ".")
	var parents []string
	for i := 0; i < len(segments)-1; i++ {
		parents = append(parents, strings.Join(segments[:i+1], "."))
	}
	return parents
}

func GetGoPackageName(fullName string) string {
	segments := strings.Split(fullName, ".")
	if len(segments) > 0 {
		name := segments[len(segments)-1]
		if core.IsVersionTag(name) {
			if len(segments) > 1 {
				name = segments[len(segments)-2]
			}
		}
		return name
	}

	return ""
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

func (m *Package) HasChild(name string) bool {
	for _, p := range m.Children {
		if p.FullName == name {
			return true
		}
	}
	return false
}

func (m *Package) GoModName() string {
	if m.Repository != nil {
		return fmt.Sprintf("%s%s/go", m.Repository.GetAuthority().GetHost(), m.Repository.GetPath())
	}
	return ""
}

func (m *Package) GoFullPackageName() string {
	if m.Repository != nil {
		fullName := strings.ReplaceAll(m.FullName, ".", "/")
		return fmt.Sprintf("%s/pkg/%s", m.GoModName(), fullName)
	}
	return ""
}

func (m *Package) GoPackageName() string {
	if m.Repository != nil {
		name := m.Name
		if core.IsVersionTag(name) {
			segments := strings.Split(m.FullName, ".")
			if len(segments) > 1 {
				name = segments[len(segments)-2]
			}
		}
		return name
	}
	return ""
}

func (m *Package) GetIdentifier(name string) *Identifier {
	if m != nil {
		id := m.Scope.GetIdentifier(name)
		if id != nil {
			return id
		}

		for _, pkg := range m.ResolvedDependencies {
			id = pkg.GetIdentifier(name)
			if id != nil {
				return id
			}
		}
	}
	return nil
}

//func (m *Package) Merge(pkg *Package)  {
//	//m.ResolvedDependencies = pkg.ResolvedDependencies
//	m.Children = pkg.Children
//	m.Scope = pkg.Scope
//	m.SourceFiles = pkg.SourceFiles
//}
