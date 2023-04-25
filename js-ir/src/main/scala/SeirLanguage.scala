import scala.util._
import scala.reflect.Typeable

case class SeirVal(repr: Any, shadows: Map[String, Any] = Map())

type SeirIdent = String

enum SeirExpr:
  case Re(value: SeirVal)
  case Scope(of: List[SeirExpr])
  case Decl(name: SeirIdent)
  case DeclNoTransform(name: SeirIdent)
  case Def(name: SeirIdent, to: SeirExpr)
  case DefNoTransform(name: SeirIdent, to: SeirExpr)
  case DefEvent(name: SeirIdent, callback: SeirExpr)
  case DefEventNoTransform(name: SeirIdent, callback: SeirExpr)
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
  case Hidden(text: String)
  case Sigil
  case SigilArgRef(text: String)
  case EOF

class SeirParseError(private val msg: String)
  extends Exception(msg)

class SeirParseUnmatchedParenError()
  extends SeirParseError("unmatched right paren")

class SeirParseEOFError()
  extends SeirParseError("unexpected eof")

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
  
  // Treats any whitespace as spaces
  def takeUntil(search: String): String =
    stealthTake match
      case Some((ch, take)) =>
        if search contains ch then
          ""
        else if (search contains ' ') && ch.isWhitespace then
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
        case '<' =>
          val text = takeUntil(">")
          take match
            case None => SeirTok.EOF
            case _ => ()
          SeirTok.Hidden(text)
        case '~' =>
          stealthTake match
            case Some(('#', take)) =>
              take()
              SeirTok.SigilArgRef(
                takeUntil(" )}").strip
              )
            case Some(_) => SeirTok.Sigil
            case None => SeirTok.EOF
        case ch =>
          if ch.isWhitespace then
            takeToken
          else
            SeirTok.IdentLike((
              ch.toString() ++ takeUntil(" )}")
            ).strip)
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
        mkFailure(s"required token \"$tok\", got \"$bad\"")

  def takeExprsUntil[E <: SeirParseError](using Typeable[E]): Try[List[SeirExpr]] =
    takeExpr match
      case Success(expr) =>
        takeExprsUntil[E]
          .map(expr :: _)
      case Failure(exn) =>
        exn match
          case _: E => Success(List())
          case _ => Failure(exn)
  
  def takeRemainingExprs = takeExprsUntil[SeirParseUnmatchedParenError]

  def takeAllExprs = takeExprsUntil[SeirParseEOFError]
  
  def takeExpr: Try[SeirExpr] =
    takeToken match
      case SeirTok.IdentLike(text) =>
        Success(SeirExpr.Var(text))
      case SeirTok.LBrace =>
        takeToken match
          case SeirTok.IdentLike(ctorName) =>
            SeirParser.ctors.get(ctorName) match
              case Some(ctor) => ctor(this)
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
              case "declnt" =>
                takeToken match
                  case SeirTok.IdentLike(name) => takeTokenRequire(
                    SeirTok.RParen,
                    SeirExpr.DeclNoTransform(name)
                  )
                  case bad =>
                    mkFailure("unexpected name for declnt \"$bad\"")
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
              case "defnt" =>
                takeToken match
                  case SeirTok.IdentLike(name) =>
                    takeExpr
                      .flatMap(expr => takeTokenRequire(
                        SeirTok.RParen,
                        SeirExpr.DefNoTransform(name, expr)
                      ))
                  case bad =>
                    mkFailure(s"unexpected name for defnt \"$bad\"")
              case "defev" =>
                takeToken match
                  case SeirTok.IdentLike(name) =>
                    takeExpr
                      .flatMap(expr => takeTokenRequire(
                        SeirTok.RParen,
                        SeirExpr.DefEvent(name, expr)
                      ))
                  case bad =>
                    mkFailure(s"unexpected name for defev \"$bad\"")
              case "defevnt" =>
                takeToken match
                  case SeirTok.IdentLike(name) =>
                    takeExpr
                      .flatMap(expr => takeTokenRequire(
                        SeirTok.RParen,
                        SeirExpr.DefEventNoTransform(name, expr)
                      ))
                  case bad =>
                    mkFailure(s"unexpected name for defev \"$bad\"")
              case "hidden" =>
                takeToken match
                  case SeirTok.Hidden(text) => takeTokenRequire(
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
      case SeirTok.Sigil =>
        takeExpr
          .map(SeirExpr.Capture.apply)
      case SeirTok.SigilArgRef(text) =>
        text.toIntOption match
          case Some(int) => Success(
            SeirExpr.ArgRef(int)
          )
          case None => mkFailure(s"arg ref must be integer, not \"$text\"")
      case SeirTok.RParen =>
        Failure(SeirParseUnmatchedParenError())
      case SeirTok.EOF =>
        Failure(SeirParseEOFError())
      case bad => mkFailure(s"unexpected start of expr \"$bad\"")

object SeirParser:
  private var ctors: Map[String, SeirParser => Try[SeirExpr]] = Map()

  def register(ctorName: String, ctor: SeirParser => Try[SeirExpr]): Unit =
    ctors += (ctorName, ctor)

  SeirParser.register("int", p =>
    p.takeToken match
      case SeirTok.IdentLike(text) =>
        text.toIntOption match
          case Some(int) => p.takeTokenRequire(
            SeirTok.RBrace,
            SeirExpr.Re(SeirVal(int))
          )
          case None =>
            p.mkFailure(s"int ctor expected integer, got \"$text\"")
      case bad =>
        p.mkFailure(s"int ctor got unexpected \"$bad\"")
  )
