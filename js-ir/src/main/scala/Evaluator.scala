case class SeirPrelude(exprs: List[SeirExpr]):
    def transform(src: SeirExpr): SeirExpr =
        SeirExpr.Scope(
            exprs ++ List(src)
        )

type SeirFnRepr = PartialFunction[List[SeirVal], SeirVal]

case class ShadowOpSpec(shadow: String, op: String)

case class ShadowHandles(handles: Map[ShadowOpSpec, List[SeirVal] => Any])

case class SeirEvaluator(
    vars: Map[String, SeirVal] = Map(),
    shadowHandles: ShadowHandles = summon[ShadowHandles]
):
    def eval(expr: SeirExpr): SeirVal =
        ???
    
    def apply(call: SeirExpr.Call): SeirVal =
        call match
            case SeirExpr.Call(f, args) =>
                (
                    eval(f)
                        .repr
                        .asInstanceOf[SeirFnRepr]
                )(args.map(eval))

def evalSeir(expr: SeirExpr): SeirVal =
    SeirEvaluator().eval(
        summon[SeirPrelude].transform(expr)
    )
