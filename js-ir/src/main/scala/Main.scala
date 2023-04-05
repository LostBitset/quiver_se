@main def hello: Unit =
  println("Nothing here yet. ")

enum IrExpr[+R : IrType]:
  case Value[R : IrType](repr: R)                                    extends IrExpr[R]
  case FCall[R : IrType](func: IrExpr[IrFunc[R]], args: List[IrAny]) extends IrExpr[R]
  case Var[R : IrType](name: String)                                 extends IrExpr[R]

type IrAny = IrWrapped[?]
class IrWrapped[T : IrType](value: T)

enum IrStmt[+R : IrType]:
  case Expr[V](expr: IrExpr[V])          extends IrStmt[Nothing]
  case Retn[R : IrType](expr: IrExpr[R]) extends IrStmt[R]

type IrBody[+R] = List[IrStmt[R]]

trait IrType[A]

given BottomIsType: IrType[Nothing]                     with {}

enum IrFunc[+R : IrType]:
  case Constr[R : IrType](body: IrBody[R], args: List[String]) extends IrFunc[R]
  case Lambda[R : IrType](body: IrBody[R], args: List[String]) extends IrFunc[R]
  case Builtin(name: String)                                   extends IrFunc[Nothing]

given IrFuncIsType[R]: IrType[IrFunc[R]]                with {}

class IrArray(backing: List[IrAny])
given IrArrayIsType: IrType[IrArray]                    with {}

class IrNumber(n: Int | Double)
given IrNumberIsType: IrType[IrNumber]                  with {}

class IrBool(b: Boolean)
given IrBoolIsType: IrType[IrBool]                      with {}

class IrString(s: String)
given IrStringIsType: IrType[IrString]                  with {}

class Shadow[+R : IrType, +S](value: R, shadow: S)
given ShadowIsType[R : IrType, S]: IrType[Shadow[R, S]] with {}
