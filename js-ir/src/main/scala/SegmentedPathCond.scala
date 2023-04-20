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
    PathCondSegment(
        segmentPC(pc, SeirCallbackRef.Top, List())
    )

@tailrec
def segmentPC(
    pc: List[String], entryCallback: SeirCallbackRef, pfx: List[PathCondSegment]
): List[PathCondSegment] =
    pc match
        case h :: t =>
            if h.startsWith(SMT_EVTRXN_MAGIC) then
                segmentPC(
                    t,
                    SeirCallbackRef.ForEvent(
                        h.drop(SMT_EVTRXN_MAGIC.length)
                    ),
                    pfx
                )
            else
                segmentPC(
                    t,
                    entryCallback,
                    (
                        PathCondSegment
                    ) :: pfx,
                )
        case Nil => pfx
    
    
