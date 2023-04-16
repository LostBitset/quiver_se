import scala.util._

@main def main: Unit =
  println("Nothing here yet. ")

case class SeirVal(repr: Any)

type SeirIdent = String

enum SeirExpr:
  case Re(value: SeirVal)
  case Scope(of: List[SeirExpr])
  case Decl(name: SeirIdent)
  case Def(name: SeirIdent, to: SeirExpr)
  case Var(name: SeirIdent)
  case Call(f: SeirExpr, args: List[SeirExpr])
  case Hidden(str: String)
  case Capture(expr: SeirExpr)
  case ArgRef(pos: Int)

enum SeirTok:
  case LParen
  case RParen
  case LBrace
  case RBrace
  case IdentLike(text: String)
  case CallHead
  case Capture(text: String)
  case EOF

final case class SeirParseError(private val msg: String)
  extends Exception(msg)

// Extractor object to allow matching strings
// by their head and tail.
object ~~:: {
  def unapply(str: String): Option[(Char, String)] =
    str.headOption.map { (_, str.tail) }
}

class SeirParser(var text: String):
  
  def stealthTake: Option[(Char, () => Char)] =
    text match
      case h ~~:: t => Some(
        (h, () => {
          text = t
          h
        })
      )
      case _ => None

  def take: Option[Char] =
    stealthTake.map({
      case (_, take) => take()
    })
  
  def takeUntil(search: String): String =
    stealthTake match
      case Some((ch, take)) =>
        if search contains ch then
          ""
        else
          take().toString ++ takeUntil(search)
      case None => ""

  def takeToken: SeirTok =
    take match
      case Some(ch) => ch match
        case '(' => SeirTok.LParen
        case ')' => SeirTok.RParen
        case '{' => SeirTok.LBrace
        case '}' => SeirTok.RBrace
        case '.' => SeirTok.CallHead
        case '<' => SeirTok.Capture(takeUntil(">"))
        case ch =>
          if ch.isWhitespace then
            takeToken
          else
            SeirTok.IdentLike(
              ch.toString() ++ takeUntil(" )")
            )
      case None =>
        SeirTok.EOF
  
  def mkFailure[T](text: String): Failure[T] =
    Failure[T](
      new SeirParseError(text)
    )
  
  def takeTokenRequire[A](tok: SeirTok, x: A): Try[A] =
    takeToken match
      case `tok` => Success(x)
      case bad =>
        mkFailure(s"required token \"$tok\", got \"bad\"")

  def takeRemainingExprs: Try[List[SeirExpr]] = ???
  
  def takeExpr: Try[SeirExpr] =
    takeToken match
      case SeirTok.IdentLike(text) =>
        Success(SeirExpr.Var(text))
      case SeirTok.LBrace =>
        takeToken match
          case SeirTok.IdentLike(ctorName) =>
            SeirParser.ctors.get(ctorName) match
              case Some(ctor) =>
                  ctor(this)
                    .flatMap(x => takeTokenRequire(SeirTok.RBrace, x))
              case None => mkFailure(s"unknown ctor \"$ctorName\"")
          case _ => mkFailure("construction must start with identifier")
      case SeirTok.LParen =>
        takeToken match
          case SeirTok.IdentLike(head) =>
            head match
              case "scope" =>
                takeRemainingExprs
                  .flatMap(seq => Success(
                    SeirExpr.Scope(seq)
                  ))
              case "decl" =>
                takeToken match
                  case SeirTok.IdentLike(name) => takeTokenRequire(
                    SeirTok.RParen,
                    SeirExpr.Decl(name)
                  )
                  case bad =>
                    mkFailure("unexpected name for decl \"$bad\"")
              case "def" =>
                takeToken match
                  case SeirTok.IdentLike(name) =>
                    takeExpr
                      .flatMap(expr => takeTokenRequire(
                        SeirTok.RParen,
                        SeirExpr.Def(name, expr)
                      ))
                  case bad =>
                    mkFailure(s"unexpected name for def \"$bad\"")
              case "hidden" =>
                takeToken match
                  case SeirTok.Capture(text) => takeTokenRequire(
                    SeirTok.RParen,
                    SeirExpr.Hidden(text)
                  )
                  case bad =>
                    mkFailure(s"unexpected token after head hidden \"$bad\"")
              case bad =>
                mkFailure(s"unexpected ident-like head \"$bad\"")
          case SeirTok.CallHead =>
            takeRemainingExprs
              .flatMap(
                seq => seq match
                  case f :: args => Success(
                    SeirExpr.Call(f, args)
                  )
                  case _ => mkFailure(s"call requires function")
              )
          case bad => mkFailure(s"unexpected head \"$bad\"")
      case bad => mkFailure(s"unexpected start of expr \"$bad\"")

object SeirParser:
  private var ctors: Map[String, SeirParser => Try[SeirExpr]] = Map()
