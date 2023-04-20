import scala.collection.mutable.{ListBuffer => MutList}
import scala.annotation.tailrec

case class SegmentedPathCond(segments: List[PathCondSegment])

case class PathCondSegment(callback: SeirCallbackRef, constraint: List[PathCondItem])

case class PathCondItem(constraint: String, followed_value: Boolean)

enum SeirCallbackRef:
    case ForEvent(name: String)
    case Top

def extractSPC(evaluator: SeirEvaluator): SegmentedPathCond =
    val pc =
        evaluator
            .shadowCtx
            .map("path-cond")
            .asInstanceOf[MutList[String]]
            .toList
    ???

def splitGroups[A, B](list: List[A], f: A => Option[B], start: B): Map[B, List[A]] =
  @tailrec
  def loop(acc: Map[B, List[A]], rest: List[A], group: List[A], key: B): Map[B, List[A]] =
    rest match
      case Nil => acc + (key -> group.reverse)
      case x :: xs =>
        f(x) match
            case None =>
                loop(acc, xs, x :: group, key)
            case Some(key) =>
                loop(acc + (key -> group.reverse), xs, List.empty, key)
  loop(Map(start -> List.empty), list, List.empty, start)
