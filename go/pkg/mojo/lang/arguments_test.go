package lang

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewArguments(t *testing.T) {
	arguments := NewArguments(NewStringLiteralExpressionFrom("foo"), NewStringLiteralExpressionFrom("bar"))
	assert.Equal(t, 2, len(arguments))
}
