	;; GENERATED SMTLib2VA code, targeting z3 @@ <qse.SMTLib2VAStringSystem>.Prologue
	;; Force MUC generation @@ <qse.SMTLib2VAStringSystem>.Prologue
	(set-option :produce-unsat-cores true)
	(set-option :smt.core.minimize true) ;; *z3 specific* @@ <qse.SMTLib2VAStringSystem>.Prologue
	;; Declarations, generated from free_funs @@ <qse.SMTLib2VAStringSystem>.GenDecls
	(declare-fun X () Real)
	;; Clauses of the conjunction, as assertions @@ <qse.SMTLib2VAStringSystem>.CheckSat
	;; All should be named, but only at the top level @@ ...
	;; This allows MUC generation to produce an MUS @@ ...
(assert (! (= (< X 42.0) false) :named ga_10))
	;; Get the results from the SMT solver @@ <SMTLib2VAStringSystem>.Epilogue
	;; This mostly s-exprs, but special lists are delimited with "[|" and "|]" @@ ...
	;; This allows it to be matched quickly with regexes @@ ...
	(echo "
;; Solver done, response below @@ ::smtlib2VA-invocation @@ <SMTLib2VAStringSystem>.Epilogue
[|resp
	[|resp.sat ")
(check-sat)
(echo "|]
	[|resp.mus ")
(get-unsat-core)
(echo "|]
	[|resp.mdl ")
(get-model)
(echo "|]|]
")
	;; MARK EOF
