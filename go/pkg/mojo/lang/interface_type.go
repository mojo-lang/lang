package lang

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
	"strings"
)

import "github.com/gertd/go-pluralize"

// type name pluralize
// generic type name
// union label

var pluralizing *pluralize.Client

func init() {
	pluralizing = pluralize.NewClient()
}

type MethodName struct {
	Method   string
	Resource string
}

const (
	GetMethod    = "get"
	ListMethod   = "list"
	CreateMethod = "create"
	UpdateMethod = "update"
	DeleteMethod = "delete"

	BatchMethod       = "batch"
	BatchGetMethod    = "batch_get"
	BatchCreateMethod = "batch_create"
	BatchUpdateMethod = "batch_update"
	BatchDeleteMethod = "batch_delete"
)

func (m *InterfaceType) GetMethodGroups() map[string][]*FunctionDecl {
	groups := make(map[string][]*FunctionDecl)
	if m == nil {
		return groups
	}

	var unprocessedFuncs []*FunctionDecl
	for _, method := range m.Methods {
		if name := method.parseStandardMethod(); name != nil {
			// set attribution
			method.SetImplicitStringAttribute(core.MethodAttributeName, name.Method)
			method.SetImplicitStringAttribute(core.ResourceAttributeName, name.Resource)

			groups[name.Resource] = append(groups[name.Resource], method)
		} else {
			unprocessedFuncs = append(unprocessedFuncs, method)
		}
	}

	for _, function := range unprocessedFuncs {
		name := pluralizing.Singular(function.Name)
		for group := range groups {
			gn := "_" + group
			if strings.HasSuffix(name, gn) {
				function.SetImplicitStringAttribute(core.MethodAttributeName, strings.TrimSuffix(function.Name, gn))
				function.SetImplicitStringAttribute(core.ResourceAttributeName, group)
				groups[group] = append(groups[group], function)
			}
		}
	}

	return groups
}

func (m *FunctionDecl) parseStandardMethod() *MethodName {
	if m != nil {
		getMethodName := func(name string) *MethodName {
			return &MethodName{Method: name, Resource: strcase.ToCamel(pluralizing.Singular(m.Name[len(name)+1:]))}
		}

		segments := strings.Split(m.Name, "_")
		if len(segments) > 0 {
			switch segments[0] {
			case GetMethod:
				return getMethodName(GetMethod)
			case ListMethod:
				return getMethodName(ListMethod)
			case CreateMethod:
				return getMethodName(CreateMethod)
			case UpdateMethod:
				return getMethodName(UpdateMethod)
			case DeleteMethod:
				return getMethodName(DeleteMethod)
			case BatchMethod:
				if len(segments) > 1 {
					switch segments[1] {
					case GetMethod:
						return getMethodName(BatchGetMethod)
					case CreateMethod:
						return getMethodName(BatchCreateMethod)
					case UpdateMethod:
						return getMethodName(BatchUpdateMethod)
					case DeleteMethod:
						return getMethodName(BatchDeleteMethod)
					}
				}
			}
		}
	}
	return nil
}
