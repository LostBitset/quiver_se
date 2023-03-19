; Force minimal core generation
(set-option :produce-unsat-cores true)
(set-option :smt.core.minimize true)

; Declare some variables
(declare-fun a () Int)
(declare-fun b () Int)
(declare-fun c () Bool)
(declare-fun d () Bool)

; Constraints
;          Formula Part
;          ---------------
(assert (! (= c (<= a b))  :named ga_0))
(assert (! (= c (> a b))   :named ga_1))
(assert (! (= c d)         :named ga_2))

; Invoke the solver
(check-sat)
(get-unsat-core)

