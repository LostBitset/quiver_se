@main def hello: Unit =
  println("Nothing here yet. ")

enum IrExpr[+R : IrType]:
  case Value[R : IrType](repr: R)                                    extends IrExpr[R]
  case FCall[R : IrType](func: IrExpr[IrFunc[R]], args: List[IrAny]) extends IrExpr[R]
  case Var[R : IrType](name: String)                                 extends IrExpr[R]

type IrAny = IrWrapped[?]
case class IrWrapped[T : IrType](value: T)

enum IrStmt[+R : IrType]:
  case Expr[V](expr: IrExpr[V])               extends IrStmt[Nothing]
  case Decl(name: String)                     extends IrStmt[Nothing]
  case Bind[V](expr: IrExpr[V], name: String) extends IrStmt[Nothing]
  case Retn[R : IrType](expr: IrExpr[R])      extends IrStmt[R]

type IrBody[+R] = List[IrStmt[R]]

trait IrType[A]

given BottomIsType: IrType[Nothing]                                       with {}

enum IrFunc[+R : IrType]:
  case Constr[R : IrType](body: IrBody[R], args: List[String]) extends IrFunc[R]
  case Lambda[R : IrType](body: IrBody[R], args: List[String]) extends IrFunc[R]
  case Builtin(name: String)                                   extends IrFunc[Nothing]

given IrFuncIsType[R]: IrType[IrFunc[R]]                                  with {}

case class IrArray(backing: List[IrAny])
given IrArrayIsType: IrType[IrArray]                                      with {}

case class IrNumber(n: Int | Double)
given IrNumberIsType: IrType[IrNumber]                                    with {}

case class IrBool(b: Boolean)
given IrBoolIsType: IrType[IrBool]                                        with {}

case class IrString(s: String)
given IrStringIsType: IrType[IrString]                                    with {}

case class Shadow[+R : IrType, +S : ShadowVBound[R]](value: R, shadow: S)
given ShadowIsType[R : IrType, S : ShadowVBound[R]]: IrType[Shadow[R, S]] with {}

type ShadowVBound = [R] =>> [S] =>> ShadowV[R, S]
trait ShadowV[-R : IrType, +S]:
  def basicForm(value: R): S

given IrToShadow[R : IrType, S : ShadowVBound[R]]: Conversion[R, Shadow[R, S]]   =
  (from: R) =>
    Shadow(from, summon[ShadowV[R, S]].basicForm(from))

given IrFromShadow[R : IrType, S : ShadowVBound[R]]: Conversion[Shadow[R, S], R] =
  _.value
