package tt

import (
	"fmt"
	"strings"
	"testing"
)

// testT implements the T interface and is used to verify the Test function's
// interaction with T.
type testT []string

func (t *testT) Helper() {}

func (t *testT) Errorf(format string, args ...interface{}) {
	*t = append(*t, fmt.Sprintf(format, args...))
}

// Simple functions to test.

func add(x, y int) int {
	return x + y
}

func addsub(x int, y int) (int, int) {
	return x + y, x - y
}

func TestTTPass(t *testing.T) {
	var testT testT
	Test(&testT, Fn("addsub", addsub), Table{
		Args(1, 10).Rets(11, -9),
	})
	if len(testT) > 0 {
		t.Errorf("Test errors when test should pass")
	}
}

func TestTTFailDefaultFmtOneReturn(t *testing.T) {
	var testT testT
	Test(&testT,
		Fn("add", add),
		Table{Args(1, 10).Rets(12)},
	)
	assertOneError(t, testT, "add(1, 10) returns (-want +got):\n")
}

func TestTTFailDefaultFmtMultiReturn(t *testing.T) {
	var testT testT
	Test(&testT,
		Fn("addsub", addsub),
		Table{Args(1, 10).Rets(11, -90)},
	)
	assertOneError(t, testT, "addsub(1, 10) returns (-want +got):\n")
}

func TestTTFailCustomFmt(t *testing.T) {
	var testT testT
	Test(&testT,
		Fn("addsub", addsub).ArgsFmt("x = %d, y = %d").RetsFmt("(a = %d, b = %d)"),
		Table{Args(1, 10).Rets(11, -90)},
	)
	assertOneError(t, testT,
		"addsub(x = 1, y = 10) returns (-want +got):\n")
}

func assertOneError(t *testing.T, testT testT, wantPrefix string) {
	t.Helper()
	switch len(testT) {
	case 0:
		t.Errorf("Test didn't error when it should")
	case 1:
		if !strings.HasPrefix(testT[0], wantPrefix) {
			t.Errorf("Test wrote message:\ngot:  %q\nwant: %q...", testT[0], wantPrefix)
		}
	default:
		t.Errorf("Test wrote too many error messages")
	}
}
