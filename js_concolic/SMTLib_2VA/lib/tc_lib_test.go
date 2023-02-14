package smtlib2va

import (
	"testing"

	tc "github.com/kandu/go_tailcall/tailcall"
	"github.com/stretchr/testify/assert"
)

func TcExample(a int) tc.TailRec[int] {
	if a <= 0 {
		return tc.TailRet(a)
	} else {
		return tc.TailCall(TcExample, a-1)
	}
}

func TestLibTailCall(t *testing.T) {
	assert.Equal(t, 0, tc.TailStart(TcExample(4_000_000)))
}
