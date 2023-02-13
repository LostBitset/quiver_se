(set-option :produce-unsat-cores true)
(set-option :smt.core.minimize true) ;; *z3 specific*

(assert true)

(check-sat)
(get-model)

