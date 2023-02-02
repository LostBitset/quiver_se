package qse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdLiteral(t *testing.T) {
	var idsrc IdSource
	var lit1 IdLiteral[uint]
	var lit2 IdLiteral[uint]
	lit1 = IdLiteral[uint]{
		WithId_H[uint]{42, idsrc.Gen()},
		false,
	}
	lit2 = IdLiteral[uint]{
		WithId_H[uint]{42, idsrc.Gen()},
		false,
	}
	assert.True(t, lit1.value.Hash32() != lit2.value.Hash32())
}
