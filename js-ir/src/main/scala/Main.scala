@main def hello: Unit =
  println("Nothing here yet. ")

class IrValue[T <: IrType](c: IrConstruct[T])

enum IrConstruct[+T <: IrType]:
  case Const(value: T)                extends IrConstruct[T]
  case Array(elems: List[IrValue[?]]) extends IrConstruct[IrArray]

enum IrExpr[+R <: IrType]:
  case Value[R <: IrType](repr: IrValue[R]) extends IrExpr[R]
  case FCall[R <: IrType](func: IrExpr[IrFunc[R]], args: List[IrValue[?]]) extends IrExpr[R]

sealed trait IrType

class IrFunc[+R] extends IrType
class IrArray    extends IrType
