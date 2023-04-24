given SeirPrelude = SeirPrelude(
    prelude_DefNoTransformine_failure ++ List(
        SeirExpr.DeclNoTransform("+"),
        SeirExpr.DefNoTransform("+", SeirExpr.Re(SeirVal({
            case List(SeirVal(a : Int, _), SeirVal(b : Int, _)) =>
                SeirVal(a + b)
        } : SeirFnRepr))),
        SeirExpr.DeclNoTransform("if"),
        SeirExpr.DefNoTransform("if", SeirExpr.Re(SeirVal({
            case List(SeirVal(c : Boolean, _), SeirVal(t, _), SeirVal(f, _)) =>
                SeirVal(if c then t else f)
        } : SeirFnRepr))),
        SeirExpr.DeclNoTransform("true"),
        SeirExpr.DefNoTransform("true", SeirExpr.Re(SeirVal(true))),
        SeirExpr.DeclNoTransform("false"),
        SeirExpr.DefNoTransform("false", SeirExpr.Re(SeirVal(false))),
        SeirExpr.DeclNoTransform("int="),
        SeirExpr.DefNoTransform("int=", SeirExpr.Re(SeirVal({
            case List(SeirVal(a : Int, _), SeirVal(b: Int, _)) =>
                SeirVal(a == b)
        } : SeirFnRepr)))
    )
)

class SeirProgramReportsFailure()
    extends Exception("UNHANDLED SEIR PROGRAM FAILURE")

val prelude_DefNoTransformine_failure = List(
    SeirExpr.DeclNoTransform("_crash"),
    SeirExpr.DefNoTransform("_crash", SeirExpr.Capture(
        SeirExpr.Call(
            SeirExpr.Var("__seirevr_FAIL"),
            List()
        )
    )),
    SeirExpr.DeclNoTransform("__seirevr_FAIL"),
    SeirExpr.DefEventNoTransform("__seirevr_FAIL", SeirExpr.Capture(
        SeirExpr.Call(
            SeirExpr.Var("__seirvrr_FAILNOTRACE"),
            List()
        )
    )),
    SeirExpr.DeclNoTransform("__seirvrr_FAILNOTRACE"),
    SeirExpr.DefNoTransform("__seirvrr_FAILNOTRACE", SeirExpr.Re(SeirVal({
        case List() =>
            throw SeirProgramReportsFailure()
    } : SeirFnRepr)))
)
