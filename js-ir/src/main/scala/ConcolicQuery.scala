case class ConcolicQuery(
    languages: ConcolicQueryLanguages,
    vars: Map[String, ConcolicVarDesc],
    source: String
)

case class ConcolicQueryLanguages(smt: String, source: String)

class UnrecognizedSmtSort(sort: String)
    extends Exception(s"Unrecognized sort $sort")

case class ConcolicVarDesc(value: String, sort: String, source_name: String):
    def asSeirValue: SeirVal = SeirVal(
        sort match
            case "Int" => value.toInt
            case _ => throw UnrecognizedSmtSort(sort)
    )
