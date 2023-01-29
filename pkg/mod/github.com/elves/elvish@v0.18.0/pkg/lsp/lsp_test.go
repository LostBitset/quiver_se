package lsp

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	lsp "github.com/sourcegraph/go-lsp"
	"github.com/sourcegraph/jsonrpc2"
	"src.elv.sh/pkg/prog"
	"src.elv.sh/pkg/prog/progtest"
	"src.elv.sh/pkg/testutil"
)

var bgCtx = context.Background()

var diagTests = []struct {
	name      string
	text      string
	wantDiags []lsp.Diagnostic
}{
	{"empty", "", []lsp.Diagnostic{}},
	{"no error", "echo", []lsp.Diagnostic{}},
	{"single error", "$!", []lsp.Diagnostic{
		{
			Range: lsp.Range{
				Start: lsp.Position{Line: 0, Character: 1},
				End:   lsp.Position{Line: 0, Character: 2}},
			Severity: lsp.Error, Source: "parse", Message: "should be variable name",
		},
	}},
	{"multi line with NL", "\n$!", []lsp.Diagnostic{
		{
			Range: lsp.Range{
				Start: lsp.Position{Line: 1, Character: 1},
				End:   lsp.Position{Line: 1, Character: 2}},
			Severity: lsp.Error, Source: "parse", Message: "should be variable name",
		},
	}},
	{"multi line with CR", "\r$!", []lsp.Diagnostic{
		{
			Range: lsp.Range{
				Start: lsp.Position{Line: 1, Character: 1},
				End:   lsp.Position{Line: 1, Character: 2}},
			Severity: lsp.Error, Source: "parse", Message: "should be variable name",
		},
	}},
	{"multi line with CRNL", "\r\n$!", []lsp.Diagnostic{
		{
			Range: lsp.Range{
				Start: lsp.Position{Line: 1, Character: 1},
				End:   lsp.Position{Line: 1, Character: 2}},
			Severity: lsp.Error, Source: "parse", Message: "should be variable name",
		},
	}},
	{"text with code point beyond FFFF", "\U00010000 $!", []lsp.Diagnostic{
		{
			Range: lsp.Range{
				Start: lsp.Position{Line: 0, Character: 4},
				End:   lsp.Position{Line: 0, Character: 5}},
			Severity: lsp.Error, Source: "parse", Message: "should be variable name",
		},
	}},
}

func TestDidOpenDiagnostics(t *testing.T) {
	f := setup(t)
	for _, test := range diagTests {
		t.Run(test.name, func(t *testing.T) {
			f.conn.Notify(bgCtx, "textDocument/didOpen", didOpenParams(test.text))
			checkDiag(t, f, diagParam(test.wantDiags))
		})
	}
}

func TestDidChangeDiagnostics(t *testing.T) {
	f := setup(t)
	f.conn.Notify(bgCtx, "textDocument/didOpen", didOpenParams(""))
	checkDiag(t, f, diagParam([]lsp.Diagnostic{}))

	for _, test := range diagTests {
		t.Run(test.name, func(t *testing.T) {
			f.conn.Notify(bgCtx, "textDocument/didChange", didChangeParams(test.text))
			checkDiag(t, f, diagParam(test.wantDiags))
		})
	}
}

func TestHover(t *testing.T) {
	f := setup(t)

	f.conn.Notify(bgCtx, "textDocument/didOpen", didOpenParams(""))
	// Hover is a no-op now; just check that it doesn't error.
	var hover lsp.Hover
	err := f.conn.Call(bgCtx, "textDocument/hover", struct{}{}, &hover)
	if err != nil {
		t.Errorf("got error %v", err)
	}
}

var completionTests = []struct {
	name     string
	text     string
	params   lsp.CompletionParams
	wantKind lsp.CompletionItemKind
}{
	{"command", "", completionParams(0, 0), lsp.CIKFunction},
	{"variable", "put $", completionParams(0, 5), lsp.CIKVariable},
	{"bad", "put [", completionParams(0, 5), 0},
}

func TestCompletion(t *testing.T) {
	f := setup(t)
	testutil.Setenv(t, "PATH", "")

	for _, test := range completionTests {
		t.Run(test.name, func(t *testing.T) {
			var items []lsp.CompletionItem
			f.conn.Notify(bgCtx, "textDocument/didOpen", didOpenParams(test.text))
			err := f.conn.Call(bgCtx, "textDocument/completion", test.params, &items)
			if err != nil {
				t.Errorf("got error %v", err)
			}
			if test.wantKind == 0 {
				if len(items) > 0 {
					t.Errorf("got %v items, want 0", len(items))
				}
			} else {
				if len(items) == 0 {
					t.Fatalf("got 0 items, want non-zero")
				}
				if items[0].Kind != test.wantKind {
					t.Errorf("got kind %v, want %v", items[0].Kind, test.wantKind)
				}
			}
		})
	}
}

var jsonrpcErrorTests = []struct {
	method  string
	params  interface{}
	wantErr error
}{
	{"unknown/method", struct{}{}, errMethodNotFound},
	{"textDocument/didOpen", []int{}, errInvalidParams},
	{"textDocument/didChange", []int{}, errInvalidParams},
	{"textDocument/completion", []int{}, errInvalidParams},
}

func TestJSONRPCErrors(t *testing.T) {
	f := setup(t)
	for _, test := range jsonrpcErrorTests {
		t.Run(test.method, func(t *testing.T) {
			err := f.conn.Call(context.Background(), test.method, test.params, &struct{}{})
			if err.Error() != test.wantErr.Error() {
				t.Errorf("got error %v, want %v", err, errMethodNotFound)
			}
		})
	}
}

func TestProgramErrors(t *testing.T) {
	progtest.Test(t, &Program{},
		progtest.ThatElvish("").
			ExitsWith(2).
			WritesStderr("internal error: no suitable subprogram\n"))
}

const testURI = "file:///foo"

func didOpenParams(text string) lsp.DidOpenTextDocumentParams {
	return lsp.DidOpenTextDocumentParams{
		TextDocument: lsp.TextDocumentItem{URI: testURI, Text: text}}
}

func didChangeParams(text string) lsp.DidChangeTextDocumentParams {
	return lsp.DidChangeTextDocumentParams{
		TextDocument: lsp.VersionedTextDocumentIdentifier{
			TextDocumentIdentifier: lsp.TextDocumentIdentifier{URI: testURI},
		},
		ContentChanges: []lsp.TextDocumentContentChangeEvent{
			{Text: text},
		}}
}

func diagParam(diags []lsp.Diagnostic) lsp.PublishDiagnosticsParams {
	return lsp.PublishDiagnosticsParams{URI: testURI, Diagnostics: diags}
}

func checkDiag(t *testing.T, f *clientFixture, want lsp.PublishDiagnosticsParams) {
	t.Helper()
	select {
	case got := <-f.diags:
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	case <-time.After(testutil.Scaled(time.Second)):
		t.Errorf("time out")
	}
}

func completionParams(line, char int) lsp.CompletionParams {
	return lsp.CompletionParams{
		TextDocumentPositionParams: lsp.TextDocumentPositionParams{
			TextDocument: lsp.TextDocumentIdentifier{URI: testURI},
			Position:     lsp.Position{Line: line, Character: char},
		},
	}
}

type clientFixture struct {
	conn  *jsonrpc2.Conn
	diags <-chan lsp.PublishDiagnosticsParams
}

func setup(t *testing.T) *clientFixture {
	r0, w0 := testutil.MustPipe()
	r1, w1 := testutil.MustPipe()

	// Run server
	done := make(chan struct{})
	go func() {
		prog.Run([3]*os.File{r0, w1, nil}, []string{"elvish", "-lsp"}, &Program{})
		close(done)
	}()
	t.Cleanup(func() { <-done })

	// Run client
	diags := make(chan lsp.PublishDiagnosticsParams, 100)
	client := client{diags}
	conn := jsonrpc2.NewConn(context.Background(),
		jsonrpc2.NewBufferedStream(transport{r1, w0}, jsonrpc2.VSCodeObjectCodec{}),
		client.handler())
	t.Cleanup(func() { conn.Close() })

	// LSP handshake
	err := conn.Call(context.Background(),
		"initialize", lsp.InitializeParams{}, &lsp.InitializeResult{})
	if err != nil {
		t.Errorf("got error %v, want nil", err)
	}
	err = conn.Notify(context.Background(), "initialized", struct{}{})
	if err != nil {
		t.Errorf("got error %v, want nil", err)
	}

	return &clientFixture{conn, diags}
}

type client struct {
	diags chan<- lsp.PublishDiagnosticsParams
}

func (c *client) handler() jsonrpc2.Handler {
	return routingHandler(map[string]method{
		"textDocument/publishDiagnostics": c.publishDiagnostics,
	})
}

func (c *client) publishDiagnostics(_ context.Context, _ jsonrpc2.JSONRPC2, rawParams json.RawMessage) (interface{}, error) {
	var params lsp.PublishDiagnosticsParams
	err := json.Unmarshal(rawParams, &params)
	if err != nil {
		panic(fmt.Sprintf("parse PublishDiagnosticsParams: %v", err))
	}
	c.diags <- params
	return nil, nil
}
