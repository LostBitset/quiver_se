import grapple.json.{_, given}

import scala.language.implicitConversions

given JsonInput[ConcolicQuery] with
    def read(json: JsonValue) =
        ConcolicQuery(json("languages"), json("vars"), json("source"))

given JsonInput[ConcolicQueryLanguages] with
    def read(json: JsonValue) =
        ConcolicQueryLanguages(json("smt"), json("source"))

given JsonInput[ConcolicVarDesc] with
    def read(json: JsonValue) =
        ConcolicVarDesc(
            json("smt_name"),
            json("assigned_value"),
            json("sort"),
            json("source_name")
        )
