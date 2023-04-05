@main def hello: Unit =
  println("Nothing here yet. ")

class IrValue[T : IrType](c: IrConstruct[T])

enum IrConstruct[+T : IrType]:
  case Const[T: IrType](value: T)     extends IrConstruct[T]
  case Array(elems: List[IrValue[?]]) extends IrConstruct[IrArray]

enum IrExpr[+R : IrType]:
  case Value[R : IrType](repr: IrValue[R]) extends IrExpr[R]
  case FCall[R : IrType](func: IrExpr[IrFunc[R]], args: List[IrValue[?]]) extends IrExpr[R]

trait IrType[A]

class IrFunc[+R]
given IrFuncIsType[R]: IrType[IrFunc[R]] with {}

class IrArray
given IrArrayIsType: IrType[IrArray] with {}

class Shadow[+R : IrType, +S](value: R, shadow: S)
given ShadowIsType[R : IrType, S]: IrType[Shadow[R, S]] with {}
