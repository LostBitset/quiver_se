(set-option :produce-unsat-cores true)
(set-option :smt.core.minimize true) ;; *z3 specific*
(assert (not false))
(check-sat)
(get-model)
