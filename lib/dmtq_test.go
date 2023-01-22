package qse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDMTQ(t *testing.T) {
	var q Quiver[uint32_H, map[Literal[uint32_H]]struct{}, *DMT[uint32_H, QuiverIndex]]
}
