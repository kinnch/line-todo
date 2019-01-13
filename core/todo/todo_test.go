package todo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_decodeMessage_invalid_format(t *testing.T) {
	todo,err := decodeMessage("Buy milk")
	assert.Zero(t,todo)
	assert.NotNil(t,err)
}