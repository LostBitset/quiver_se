package smtlib2va

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexicallyScoped(t *testing.T) {
	lvbls := NewLexicallyScoped()
	lvbls.EnterScope()
	lvbls.DeclVar("test")
	lvbls.EnterScope()
	lvbls.EnterScope()
	lvbls.WriteVar("test", "hello")
	lvbls.LeaveScope()
	assert.True(t, lvbls.IsDefined("test"))
	assert.Equal(t, "hello", lvbls.ReadVar("test"))
	lvbls.LeaveScope()
	lvbls.LeaveScope()
	lvbls.EnterScope()
	assert.False(t, lvbls.IsDefined("test"))
	lvbls.EnterScope()
}

func TestLexicallyScoped2(t *testing.T) {
	lvbls := NewLexicallyScoped()
	lvbls.EnterScope()
	lvbls.DeclVar("f")
	lvbls.EnterScope()
	lvbls.WriteVar("f", "false")
	lvbls.LeaveScope()
	lvbls.EnterScope()
	lvbls.DeclVar("m")
	lvbls.WriteVar("m", "42.0")
	lvbls.DeclVar("y")
	lvbls.WriteVar("y", "something")
	assert.True(t, lvbls.IsDefined("m"))
	assert.True(t, lvbls.IsDefined("y"))
	assert.True(t, lvbls.IsDefined("f"))
	lvbls.LeaveScope()
	lvbls.LeaveScope()
}
