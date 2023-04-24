package qse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPHashMap(t *testing.T) {
	pm := NewPHashMap[uint32_H, int]()
	assert.Equal(t, false, pm.HasKey(uint32_H{7}))
	pm2 := pm.Assoc(uint32_H{7}, 440)
	assert.Equal(t, false, pm.HasKey(uint32_H{7}))
	assert.Equal(t, true, pm2.HasKey(uint32_H{7}))
	val, _ := pm2.Index(uint32_H{7})
	assert.Equal(t, val, 440)
	pm3 := pm2.Dissoc(uint32_H{7})
	assert.Equal(t, false, pm.HasKey(uint32_H{7}))
	assert.Equal(t, true, pm2.HasKey(uint32_H{7}))
	assert.Equal(t, false, pm3.HasKey(uint32_H{7}))
}

