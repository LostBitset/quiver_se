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
    SegmentedPathCond(
        splitGroups(pc, SeirCallbackRef.Top)(item =>
            if item.startsWith(SMT_EVTRXN_MAGIC) then
                Some(item.drop(SMT_EVTRXN_MAGIC.length))
            else
                None
        )
            .map({
                case (callback: SeirCallbackRef, constraintStrings) =>
                    PathCondSegment(
                        callback,
                        constraintStrings.map(decodePathCondItem)
                    )
            })
    )

def decodePathCondItem(s: String): PathCondItem =
    

// -- begin nh
def splitGroups[A, B](list: List[A], start: B)(f: A => Option[B]): List[(B, List[A])] = {
  // A helper function to accumulate groups into a list of pairs
  @annotation.tailrec
  def loop(acc: List[(B, List[A])], rest: List[A], currentGroup: List[A], currentKey: B): List[(B, List[A])] = {
    rest match {
      case Nil => acc :+ (currentKey, currentGroup.reverse) // Reverse the last group and add it to the accumulator
      case x :: xs => f(x) match {
        case None => loop(acc, xs, x :: currentGroup, currentKey) // Keep accumulating the current group
        case Some(key) => loop(acc :+ (currentKey, currentGroup.reverse), xs, List.empty, key) // Start a new group with the new key
      }
    }
  }

  loop(List(), list, List.empty, start)
}
// -- end nh
