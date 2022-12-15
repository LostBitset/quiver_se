package quiver_se

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSanityCheck(t *testing.T) {
	assert.Equal(t, Greet("World"), "Hello, World!")
}

