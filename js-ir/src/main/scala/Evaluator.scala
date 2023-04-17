case class SeirPrelude(exprs: List[SeirExpr]):
    def transform(src: SeirExpr): SeirExpr =
        SeirExpr.Scope(
            exprs ++ List(src)
        )

type SeirFnRepr = PartialFunction[List[SeirVal], SeirVal]

case class ShadowOpSpec[+A](shadow: String, op: ShadowOp[A])

enum ShadowOp[+A]:
    case Named(name: String) extends ShadowOp[List[Any] => Any]
    case Promote() extends ShadowOp[Any => Any]
    case Fallback() extends ShadowOp[Any]

case class ShadowHandles(handles: TiedMap[ShadowOpSpec]):
    def extract(shadow: String)(value: Any): Any =
        handles.get(shadow)

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
                val base = env(name)
                SeirVal(
                    base.repr,
                    base.shadows + ("@@varname" -> name)
                )
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
            
    def applyDropShadows(f: SeirVal, args: List[SeirVal]): SeirVal =
        f.repr match
            case QuotedCapture(expr) =>
                eval(expr, args)
            case other =>
                other.asInstanceOf[SeirFnRepr](args)
    
    def apply(f: SeirVal, args: List[SeirVal]): SeirVal =
        val base = applyDropShadows(f, args)
        base.shadows.get("@@varname") match
            case Some(name: String) =>
                val rewrittenShadows = base
                    .shadows
                    .keys
                    .filter(_.startsWith("@@"))
                    .flatMap(
                        shadow => shadowHandles.handles.get(
                            ShadowOpSpec(shadow, ShadowOp.Named(name))
                        )
                            .map(shadow -> _(
                                args.map(
                                    shadowHandles.extract(shadow)
                                )
                            ))
                    )
                    .toMap
                SeirVal(base.repr, base.shadows ++ rewrittenShadows)
            case Some(bad) =>
                throw IllegalArgumentException(s"expected @@varname -> <String>, got \"$bad\"")
            case None =>
                base


def evalSeir(expr: SeirExpr): SeirVal =
    SeirEvaluator().eval(
        summon[SeirPrelude].transform(expr)
    )
