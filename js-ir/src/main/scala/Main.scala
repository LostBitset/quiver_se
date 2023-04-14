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
  case Prop(obj: SeirExpr, property: SeirIdent)
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

class SeirParser(var text: String):

  def take: Option[Char] = text match:
    case h :: t => {
      text = t
      h
    }
    case _ => None

  def takeToken: SeirTok =
    take match:
      case Some(ch) => ch match:
        case '(' => SeirTok.LParen
        case ')' => SeirTok.RParen
        case '{' => SeirTok.LBrace
        case '}' => SeirTok.RBrace
        case '.' => SeirTok.CallHead(takeUntil(" )"))
        case '<' => SeirTok.Capture(takeUntil(">"))
        case ch =>
          if ch.isWhitespace then
            takeToken
          else
            SeirTok.IdentLike(takeUntil(" )"))
      case None =>
        SeirTok.EOF
