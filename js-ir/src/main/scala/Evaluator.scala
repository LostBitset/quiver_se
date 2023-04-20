import scala.collection.mutable.{ListBuffer => MutList}

case class SeirPrelude(exprs: List[SeirExpr]):
    def transform(src: SeirExpr): SeirExpr =
        SeirExpr.Scope(
            exprs ++ List(src)
        )

type SeirFnRepr = PartialFunction[List[SeirVal], SeirVal]

case class ShadowCtx(
    var map: Map[String, Any] = Map(
        "@@evtrxns" -> MutList[String]()
    )
)

case class ShadowOpSpec[+A](shadow: String, op: ShadowOp[A])

enum ShadowOp[+A]:
    case Named(name: String) extends ShadowOp[(List[ShadowPersp], ShadowCtx) => Any]
    case Promote extends ShadowOp[Any => Option[Any]]

case class ShadowHandles(handles: TiedMap[ShadowOpSpec]):
    def extract(shadow: String)(value: SeirVal): Option[Any] =
        value.shadows.get(shadow) match
            case Some(value) =>
                Some(value)
            case None =>
                handles
                    .get(ShadowOpSpec(shadow, ShadowOp.Promote))
                    .flatMap(_(value.repr))

case class ShadowPersp(repr: Any, shadow: Any)

case class HiddenProc(proc: String => SeirVal):
    def apply(text: String): SeirVal =
        proc(text)

case class QuotedCapture(expr: SeirExpr)

case class SeirBoundEvent(cap: QuotedCapture, name: String)

case class SeirEvaluator(
    var env: SeirEnv = SeirEnv(),
    var shadowCtx: ShadowCtx = ShadowCtx(),
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
            case SeirExpr.DefEvent(name, callback) =>
                val cap = rec(callback)
                val toVal = SeirVal(
                    SeirBoundEvent(
                        cap.repr.asInstanceOf[QuotedCapture],
                        name
                    ),
                    cap.shadows
                )
                env.define(name, toVal)
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
                ap(fValue, argValues)
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
            case SeirBoundEvent(cap, name) =>
                shadowCtx
                    .map
                    .get("@@evtrxns")
                    .get
                    .asInstanceOf[MutList[String]]
                    .addOne(name)
                applyDropShadows(SeirVal(cap), args)
            case other =>
                other.asInstanceOf[SeirFnRepr](args)
    
    def ap(f: SeirVal, args: List[SeirVal]): SeirVal =
        val base = applyDropShadows(f, args)
        f.shadows.get("@@varname") match
            case Some(name: String) =>
                val rewrittenShadows = args
                    .map(_.shadows.keys)
                    .reduce(_ ++ _)
                    .filterNot(_.startsWith("@@"))
                    .flatMap(
                        shadow => shadowHandles.handles.get(
                            ShadowOpSpec(shadow, ShadowOp.Named(name))
                        )
                            .map(shadow -> _(
                                args
                                    .zipWithIndex
                                    .flatMap(
                                        (x, i) =>
                                            shadowHandles.extract(shadow)(x)
                                                .map(shadowVal =>
                                                    ShadowPersp(
                                                        args(i).repr,
                                                        shadowVal
                                                    )
                                                )
                                    ),
                                shadowCtx,
                            ))
                    )
                    .toMap
                SeirVal(base.repr, base.shadows ++ rewrittenShadows)
            case Some(bad) =>
                throw IllegalArgumentException(s"expected @@varname -> <String>, got \"$bad\"")
            case None =>
                base
    
    def evalSeir(expr: SeirExpr): SeirVal =
        eval(
            summon[SeirPrelude].transform(expr)
        )
    
    def eventTransitions: List[String] =
        shadowCtx
            .map
            .get("@@evtrxns")
            .get
            .asInstanceOf[MutList[String]]
            .toList


def evalSeir(expr: SeirExpr): SeirVal =
    SeirEvaluator().evalSeir(expr)
