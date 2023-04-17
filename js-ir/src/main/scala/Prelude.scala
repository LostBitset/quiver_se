given SeirPrelude = SeirPrelude(List(
    SeirExpr.Decl("+"),
    SeirExpr.Def("+", SeirExpr.Re(SeirVal({
        case List(SeirVal(a : Int, _), SeirVal(b : Int, _)) =>
            SeirVal(a + b)
    } : SeirFnRepr)))
))
