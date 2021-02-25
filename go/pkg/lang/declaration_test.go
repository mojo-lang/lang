package lang

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

const EnumDeclStr = `{"@type":"EnumDecl","name":"Foo","type":{"package":"mojo.lang.test"}}`

func TestDeclarationCodec_Encode(t *testing.T) {
	decl := NewEnumDecl(&EnumDecl{
		Name: "Foo",
		Type: &EnumType{
			Package: "mojo.lang.test",
		},
	})

	str, err := jsoniter.ConfigDefault.Marshal(decl)
	assert.NoError(t, err)
	assert.Equal(t, EnumDeclStr, string(str))
}

func TestDeclarationCodec_Decode(t *testing.T) {
	decl := &Declaration{}
	err := jsoniter.ConfigDefault.Unmarshal([]byte(EnumDeclStr), decl)
	assert.NoError(t, err)
	assert.Equal(t, "Foo", decl.GetEnumDecl().Name)
}
