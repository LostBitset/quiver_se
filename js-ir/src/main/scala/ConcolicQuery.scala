case class ConcolicQuery(
    languages: ConcolicQueryLanguages,
    vars: List[ConcolicVarDesc],
    source: String
):
    def run: SegmentedPathCond =
        val exprNoContext = SeirParser(source).takeExpr.get
        val symPrelude = SeirPrelude(vars.map(_.toExpr))
        val expr = symPrelude.transform(exprNoContext)
        val evaluator = SeirEvaluator()
        evaluator.evalSeir(expr)
        extractSPC(evaluator)

case class ConcolicQueryLanguages(smt: String, source: String)

class UnrecognizedSmtSort(sort: String)
    extends Exception(s"Unrecognized sort $sort")

case class ConcolicVarDesc(smt_name: String, value: String, sort: String, source_name: String):
    def getAsSeirValue: SeirVal = SeirVal(
        sort match
            case "Int" => value.toInt
            case _ => throw UnrecognizedSmtSort(sort)
    )
