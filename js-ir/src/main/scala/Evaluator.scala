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
        lazy val rec = { eval(_, arguments) }
        expr match
            case SeirExpr.Re(value) =>
                value
            case SeirExpr.Scope(of) =>
                env.enterScope
                var lastSlot = SeirVal(())
                of.foreach(expr => {
                    lastSlot = rec(expr)
                })
                env.leaveScope
                lastSlot
            case SeirExpr.Decl(name) =>
                env.declare(name)
                SeirVal(())
            case SeirExpr.Def(name, to) =>
                env.define(name, rec(to))
                SeirVal(())
            case SeirExpr.Var(name) =>
                env(name)
            case SeirExpr.Call(f, args) =>
                val argValues = args.map(rec)
                val fValue = rec(f)
                apply(fValue, argValues)
            case SeirExpr.Hidden(str) =>
                summon[HiddenProc](str)
            case SeirExpr.Capture(expr) =>
                SeirVal(QuotedCapture(expr))
            case SeirExpr.ArgRef(pos) =>
                arguments(pos)
            
    def apply(f: SeirVal, args: List[SeirVal]): SeirVal =
        f.repr match
            case QuotedCapture(expr) =>
                eval(expr, args)
            case other =>
                other.asInstanceOf[SeirFnRepr](args)

def evalSeir(expr: SeirExpr): SeirVal =
    SeirEvaluator().eval(
        summon[SeirPrelude].transform(expr)
    )
