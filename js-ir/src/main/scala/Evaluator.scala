case class SeirPrelude(exprs: List[SeirExpr]):
    def transform(src: SeirExpr): SeirExpr =
        SeirExpr.Scope(
            exprs ++ List(src)
        )

case class ShadowOpSpec(shadow: String, op: String)

case class ShadowHandles(handles: Map[ShadowOpSpec, List[SeirVal] => Any])

case class SeirEvaluator(
    src: SeirExpr,
    vars: Map[String, SeirVal] = Map(),
    shadowHandles: ShadowHandles = ShadowHandles(Map())
):
    def eval: SeirVal = ???

given Conversion[SeirExpr, SeirEvaluator] =
    expr => SeirEvaluator(
        summon[SeirPrelude].transform(expr),
        shadowHandles = summon[ShadowHandles]
    )
