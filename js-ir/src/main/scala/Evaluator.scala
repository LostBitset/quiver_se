case class SeirPrelude(exprs: List[SeirExpr]):
    def transform(src: SeirExpr): SeirExpr =
        SeirExpr.Scope(
            exprs ++ List(src)
        )

type SeirFnRepr = PartialFunction[List[SeirVal], SeirVal]

case class ShadowOpSpec(shadow: String, op: String)

case class ShadowHandles(handles: Map[ShadowOpSpec, List[SeirVal] => Any])

case class HiddenProc(proc: String => SeirVal):
    def apply(text: String): SeirVal =
        proc(text)

case class QuotedCapture(expr: SeirExpr)

case class SeirEvaluator(
    var env: SeirEnv = SeirEnv(),
    shadowHandles: ShadowHandles = summon[ShadowHandles]
):
    def eval(expr: SeirExpr, arguments: List[SeirVal] = List()): SeirVal =
        expr match
            case SeirExpr.Re(value) => value
            case SeirExpr.Scope(of) => ???
            case SeirExpr.Decl(name) =>
                env.declare(name)
                SeirVal(())
            case SeirExpr.Def(name, to) =>
                env.define(name, eval(to))
                SeirVal(())
            case SeirExpr.Var(name) => env(name)
            case call: SeirExpr.Call => apply(call)
            case SeirExpr.Hidden(str) =>
                summon[HiddenProc](str)
            case SeirExpr.Capture(expr) => SeirVal(QuotedCapture(expr))
            case SeirExpr.ArgRef(pos) => arguments(pos)
    
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
