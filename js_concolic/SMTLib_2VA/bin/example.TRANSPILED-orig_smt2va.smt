(set-option :produce-unsat-cores true)
(set-option :smt.core.minimize true) 

(*/enter-scope/*)
(*/decl-var/* **x)
(*/enter-scope/*)
(*/enter-scope/*)
(*/write-var/* **x "just \"s and (*/enter-scope/*)")
(*/write-var/* **x true)
(*/leave-scope/*)
(assert (*/read-var/* **x))
(*/leave-scope/*)
(*/enter-scope/*)
(*/leave-scope/*)
(*/leave-scope/*)

(check-sat)
(get-model)

