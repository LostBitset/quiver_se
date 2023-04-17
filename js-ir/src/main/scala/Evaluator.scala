case class SeirPrelude(exprs: List[SeirExpr]):
    def transform(src: SeirExpr): SeirExpr =
        SeirExpr.Scope(
            exprs ++ List(src)
        )

type SeirFnRepr = PartialFunction[List[SeirVal], SeirVal]

case class ShadowOpSpec(shadow: String, op: String)

case class ShadowHandles(handles: Map[ShadowOpSpec, List[SeirVal] => Any])

case class QuotedCapture(expr: SeirExpr)

case class SeirEvaluator(
    vars: Map[String, SeirVal] = Map(),
    shadowHandles: ShadowHandles = summon[ShadowHandles]
):
    def eval(expr: SeirExpr, arguments: List[SeirVal] = List()): SeirVal =
        ???
    
    def apply(call: SeirExpr.Call): SeirVal =
        call match
            case SeirExpr.Call(f, argValues) =>
                val args = argValues.map(eval(_))
                eval(f).repr match
                    case QuotedCapture(expr) =>
                        eval(expr, args)
                    case other =>
                        (
                            other.asInstanceOf[SeirFnRepr]
                        )(args)

def evalSeir(expr: SeirExpr): SeirVal =
    SeirEvaluator().eval(
        summon[SeirPrelude].transform(expr)
    )
