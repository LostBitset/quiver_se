(set-option :produce-unsat-cores true)
(set-option :smt.core.minimize true) 








(assert (not false))





(check-sat)
(get-model)

