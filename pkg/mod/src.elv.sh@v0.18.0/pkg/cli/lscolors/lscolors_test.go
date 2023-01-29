package lscolors

import (
	"os"
	"testing"

	"src.elv.sh/pkg/testutil"
)

func TestLsColors(t *testing.T) {
	testutil.InTempDir(t)
	SetTestLsColors(t)

	// Test both feature-based and extension-based coloring.

	colorist := GetColorist()

	os.Mkdir("dir", 0755)
	create("a.png")

	wantDirStyle := "34"
	if style := colorist.GetStyle("dir"); style != wantDirStyle {
		t.Errorf("Got dir style %q, want %q", style, wantDirStyle)
	}
	wantPngStyle := "31"
	if style := colorist.GetStyle("a.png"); style != wantPngStyle {
		t.Errorf("Got dir style %q, want %q", style, wantPngStyle)
	}
}
