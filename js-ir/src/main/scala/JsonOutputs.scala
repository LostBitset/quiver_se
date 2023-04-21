import grapple.json.{_, given}

import scala.language.implicitConversions

given JsonOutput[ConcolicResult] with
    def write(u: ConcolicResult) =
        Json.obj(
            "languages" -> u.languages,
            "segmented_path_condition" -> u.spc
        )

given JsonOutput[ConcolicResultLanguages] with
    def write(u: ConcolicResultLanguages) =
        Json.obj(
            "smt" -> u.smt
        )

given JsonOutput[SegmentedPathCond] with
    def write(u: SegmentedPathCond) = u.segments

given JsonOutput[PathCondSegment] with
    def write(u: PathCondSegment) =
        Json.obj(
            "callback" -> u.callback,
            "path_cond_segment" -> u.constraint
        )

given JsonOutput[PathCondItem] with
    def write(u: PathCondItem) =
        Json.obj(
            "constraint" -> u.constraint,
            "followed_value" -> u.followed_value
        )

given JsonOutput[SeirCallbackRef] with
    def write(u: SeirCallbackRef) =
        u match
            case SeirCallbackRef.ForEvent(name) =>
                name
            case SeirCallbackRef.Top =>
                JsonNull
