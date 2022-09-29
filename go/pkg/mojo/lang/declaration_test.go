package lang

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewVariableDeclaration(t *testing.T) {
	decl := NewVariableDeclaration(nil)
	assert.Nil(t, decl.GetVariableDecl())

	decl = NewVariableDeclaration(&VariableDecl{Name: "foo"})
	assert.NotNil(t, decl.GetVariableDecl())
	assert.Equal(t, "foo", decl.GetName())
}
