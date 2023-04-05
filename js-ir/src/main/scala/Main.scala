@main def hello: Unit =
  println("Nothing here yet. ")

trait IrType[V]:
  def resolveMethod[R](method: String): IrFunc[R]

enum IrExpr[+V]:
  case Const[V : IrType](value: V)                            extends IrExpr[V]
  case Var[V: IrType](value: String)                          extends IrExpr[V]
  case Call[R : IrType](fn: IrFunc[R], args: List[IrExpr[?]]) extends IrExpr[R]

type IrFunc[+R] = IrExpr[IrLambda[R]]

class IrLambda[+R : IrType]
