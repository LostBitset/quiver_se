package edit

import (
	"testing"

	"src.elv.sh/pkg/cli/term"
	"src.elv.sh/pkg/tt"
)

func TestBufToHTML(t *testing.T) {
	tt.Test(t, tt.Fn("bufToHTML", bufToHTML), tt.Table{
		// Just plain text.
		tt.Args(
			bb().Write("abc").Buffer(),
		).Rets(
			`abc` + "\n",
		),
		// Just styled text.
		tt.Args(
			bb().WriteStringSGR("abc", "31").Buffer(),
		).Rets(
			`<span class="sgr-31">abc</span>` + "\n",
		),
		// Mixing plain and styled texts.
		tt.Args(
			bb().
				WriteStringSGR("abc", "31").
				Write(" def ").
				WriteStringSGR("xyz", "1").
				Buffer(),
		).Rets(
			`<span class="sgr-31">abc</span> def <span class="sgr-1">xyz</span>` + "\n",
		),
		// Multiple lines.
		tt.Args(
			bb().
				WriteStringSGR("abc", "31").
				Newline().Write("def").
				Newline().WriteStringSGR("xyz", "1").
				Buffer(),
		).Rets(
			`<span class="sgr-31">abc</span>` + "\n" +
				`def` + "\n" +
				`<span class="sgr-1">xyz</span>` + "\n",
		),
	})
}

func bb() *term.BufferBuilder { return term.NewBufferBuilder(20) }
