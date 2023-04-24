given SeirPrelude = SeirPrelude(
    prelude_define_failure ++ List(
        SeirExpr.Decl("+"),
        SeirExpr.Def("+", SeirExpr.Re(SeirVal({
            case List(SeirVal(a : Int, _), SeirVal(b : Int, _)) =>
                SeirVal(a + b)
        } : SeirFnRepr))),
        SeirExpr.Decl("if"),
        SeirExpr.Def("if", SeirExpr.Re(SeirVal({
            case List(SeirVal(c : Boolean, _), SeirVal(t, _), SeirVal(f, _)) =>
                SeirVal(if c then t else f)
        } : SeirFnRepr))),
        SeirExpr.Decl("true"),
        SeirExpr.Def("true", SeirExpr.Re(SeirVal(true))),
        SeirExpr.Decl("false"),
        SeirExpr.Def("false", SeirExpr.Re(SeirVal(false))),
        SeirExpr.Decl("int="),
        SeirExpr.Def("int=", SeirExpr.Re(SeirVal({
            case List(SeirVal(a : Int, _), SeirVal(b: Int, _)) =>
                SeirVal(a == b)
        } : SeirFnRepr)))
    )
)

class SeirProgramReportsFailure()
    extends Exception("UNHANDLED SEIR PROGRAM FAILURE")

val prelude_define_failure = List(
    SeirExpr.Decl("_crash"),
    SeirExpr.Def("_crash", SeirExpr.Capture(
        SeirExpr.Call(
            SeirExpr.Var("__seirevr_FAIL"),
            List()
        )
    )),
    SeirExpr.Decl("__seirevr_FAIL"),
    SeirExpr.DefEvent("__seirevr_FAIL", SeirExpr.Capture(
        SeirExpr.Call(
            SeirExpr.Var("__seirvrr_FAILNOTRACE"),
            List()
        )
    )),
    SeirExpr.Decl("__seirvrr_FAILNOTRACE"),
    SeirExpr.Def("__seirvrr_FAILNOTRACE", SeirExpr.Re(SeirVal({
        case List() =>
            throw SeirProgramReportsFailure()
    } : SeirFnRepr)))
)
