given SeirPrelude = SeirPrelude(List(
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
))
