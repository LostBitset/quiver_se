// import scala.util._

case class SeirPrelude(exprs: List[SeirExpr]):
    def transform(src: SeirExpr): SeirExpr =
        SeirExpr.Scope(
            exprs ++ List(src)
        )

case class SeirEvaluator(src: SeirExpr, vars: Map[String, SeirVal] = Map())

given Conversion[SeirExpr, SeirEvaluator] =
    expr => SeirEvaluator(
        summon[SeirPrelude].transform(expr)
    )

given SeirPrelude = SeirPrelude(List(
    SeirExpr.Decl("+"),
    SeirExpr.Def("+", SeirExpr.Re(SeirVal(
        (a: Int, b: Int) => a + b
    )))
))
