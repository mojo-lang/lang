package lang

func (m *ReferencedPackage) SetKey(key string) {
	if x, ok := m.GetReferencedPackage().(*ReferencedPackage_Key); ok {
		x.Key = key
	} else {
		m.ReferencedPackage = &ReferencedPackage_Key{
			Key: key,
		}
	}
}

func (m *ReferencedPackage) SetPackage(pkg *Package) {
	if x, ok := m.GetReferencedPackage().(*ReferencedPackage_Package); ok {
		x.Package = pkg
	} else {
		m.ReferencedPackage = &ReferencedPackage_Package{
			Package: pkg,
		}
	}
}
