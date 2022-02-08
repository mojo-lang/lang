package lang

import (
	"fmt"
	"strings"

	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
)

func NewGlobalPackage() *Package {
	return &Package{
		Summary:  "Global",
		Scope:    NewScope(),
		Implicit: true,
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

// the segment of the package name will to be a folder, using kebab-case style
func PackageNameToPath(fullName string) string {
	return strings.ReplaceAll(strcase.ToKebabWithIgnore(fullName, "."), ".", "/")
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

// auto genereated package, used to padding package tree
// Global Package is also a padding package
func (m *Package) IsPadding() bool {
	return m.Implicit && len(m.SourceFiles) == 0
}

func (m *Package) SetScope(scope *Scope) {
	if m != nil {
		m.Scope = scope
	}
}

func (m *Package) GetAllPackageArray() []*Package {
	packages := []*Package{m}

	for _, pkg := range m.Children {
		ps := pkg.GetAllPackageArray()
		packages = append(packages, ps...)
	}
	return packages
}

func (m *Package) GetAllPackages() map[string]*Package {
	packages := make(map[string]*Package)
	packages[m.FullName] = m

	for _, pkg := range m.Children {
		ps := pkg.GetAllPackages()
		for k, v := range ps {
			packages[k] = v
		}
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
	if len(m.FullName) > 0 {
		segments := strings.Split(m.FullName, ".")
		for i := 0; i < len(segments); i++ {
			segments[i] = strcase.ToKebab(segments[i])
		}
		fullName := strings.Join(segments, "/")
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

		for _, pkg := range m.Children {
			id = pkg.GetIdentifier(name)
			if id != nil {
				return id
			}
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

// resolved identifier may be defined in self package or defined in the dependent package
func (m *Package) GetResolvedIdentifier(fullName string) *Identifier {
	if m != nil {
		for _, file := range m.SourceFiles {
			if id := FindIdentifier(file.ResolvedIdentifiers, fullName); id != nil {
				return id
			}
		}

		for _, pkg := range m.ResolvedDependencies {
			if id := pkg.GetResolvedIdentifier(fullName); id != nil {
				return id
			}
		}

		return m.GetIdentifier(fullName)
	}
	return nil
}

func (m *Package) GetEntityNode(name string) *EntityNode {
	if m != nil && m.EntityRelationSet != nil {
		typeName := GetTypeTypeName(name)
		if node := m.EntityRelationSet.GetNode(typeName); node != nil {
			return node
		}
		if node := m.EntityRelationSet.GetNode(name); node != nil {
			return node
		}

		for _, pkg := range m.Children {
			if node := pkg.GetEntityNode(name); node != nil {
				return node
			}
		}

		for _, pkg := range m.ResolvedDependencies {
			if node := pkg.GetEntityNode(name); node != nil {
				return node
			}
		}
	}
	return nil
}

func (m *Package) SetEntityNode(name string, node *EntityNode) *Package {
	if m != nil {
		if m.EntityRelationSet == nil {
			m.EntityRelationSet = NewEntityRelationSet()
		}
		m.EntityRelationSet.Nodes[name] = node
	}
	return m
}

func (m *Package) GetEntityEdge(name string) *EntityEdge {
	if m != nil && m.EntityRelationSet != nil {
		return m.EntityRelationSet.Edges[name]
	}
	return nil
}

func (m *Package) SetEntityEdge(name string, edge *EntityEdge) *Package {
	if m != nil {
		if m.EntityRelationSet == nil {
			m.EntityRelationSet = NewEntityRelationSet()
		}
		m.EntityRelationSet.Edges[name] = edge
	}
	return m
}

func (m *Package) DeleteEntityEdge(name string) *Package {
	if m != nil && m.EntityRelationSet != nil {
		delete(m.EntityRelationSet.Edges, name)
	}
	return m
}
