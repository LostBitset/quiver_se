import scala.collection.mutable.{ListBuffer => MutList}

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
        segmentPC(pc, SeirCallbackRef.Top)
    )

def segmentPC(pc: List[String], entryCallback: SeirCallbackRef): List[PathCondSegment] =
    
