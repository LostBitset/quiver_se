package qse

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZ3SMTLib2VAQueryUnsat(t *testing.T) {
	query_str := strings.TrimSpace(`
	;; Force MUC generation
	(set-option :produce-unsat-cores true)
	(set-option :smt.core.minimize true) ;; *z3 specific*
	;; Declarations
	(declare-fun a () Int)
	(declare-fun b () Int)
	(declare-fun gt () Bool)
	;; Assertions (named for MUC/MUS generation)
	(assert (! (= gt (> a b))   :named ta_0))
	(assert (! (= gt (<= a b))  :named ta_1))
	(assert (! (< a b)          :named ta_2))
	;; Invoke the solver
	(echo ".sat")
	(check-sat)
	(echo ".model")
	(get-model)
	(echo ".muc")
	(get-unsat-core)
	`)
	query := NewZ3SMTLib2VAQuery(query_str)
	expected := strings.TrimSpace(`
	.sat
	unsat
	.model
	(error "line 16 column 11: model is not available")
	.muc
	(ta_1 ta_0)
	`)
	var sb_expected strings.Builder
	for _, line := range strings.Split(expected, "\n") {
		if len(line) > 0 && line[0] == '\t' {
			sb_expected.WriteString(line[1:])
		} else {
			sb_expected.WriteString(line)
		}
		sb_expected.WriteRune('\n')
	}
	expected = strings.TrimSpace(sb_expected.String())
	actual := strings.TrimSpace(
		query.Run(),
	)
	assert.Equal(t, expected, actual)
}

func TestZ3SMTLib2VAQuerySat(t *testing.T) {
	query_str := strings.TrimSpace(`
	;; Force MUC generation
	(set-option :produce-unsat-cores true)
	(set-option :smt.core.minimize true) ;; *z3 specific*
	;; Declarations
	(declare-fun a () Int)
	(declare-fun b () Int)
	(declare-fun gt () Bool)
	;; Assertions (named for MUC/MUS generation)
	(assert (! (= gt (> a b))   :named ta_0))
	(assert (! (< a b)          :named ta_1))
	;; Invoke the solver
	(echo ".sat")
	(check-sat)
	(echo ".model")
	(get-model)
	(echo ".muc")
	(get-unsat-core)
	`)
	query := NewZ3SMTLib2VAQuery(query_str)
	expected := strings.TrimSpace(`
	.sat
	sat
	.model
	(
	  (define-fun a () Int
	    0)
	  (define-fun ta_1 () Bool
	    (< a b))
	  (define-fun gt () Bool
	    false)
	  (define-fun b () Int
	    1)
	  (define-fun ta_0 () Bool
	    (= gt (> a b)))
	)
	.muc
	(error "line 17 column 16: unsat core is not available")
	`)
	var sb_expected strings.Builder
	for _, line := range strings.Split(expected, "\n") {
		if len(line) > 0 && line[0] == '\t' {
			sb_expected.WriteString(line[1:])
		} else {
			sb_expected.WriteString(line)
		}
		sb_expected.WriteRune('\n')
	}
	expected = strings.TrimSpace(sb_expected.String())
	actual := strings.TrimSpace(
		query.Run(),
	)
	assert.Equal(t, expected, actual)
}
