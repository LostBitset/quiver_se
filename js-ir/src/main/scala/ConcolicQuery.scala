case class ConcolicQuery(
    languages: ConcolicQueryLanguages,
    vars: List[ConcolicVarDesc],
    source: String
):
    assert(languages.source == "seir")
    assert(languages.smt == "smtlib_2va")

    def runExtractSPC: SegmentedPathCond =
        val exprNoContext = SeirParser(source).takeExpr.get
        val symPrelude = SeirPrelude(
            vars.flatMap(desc => List(
                desc.toExprDeclare,
                desc.toExprDefine
            ))
        )
        val expr = symPrelude.transform(exprNoContext)
        val evaluator = SeirEvaluator()
        evaluator.evalSeir(expr)
        extractSPC(evaluator)
    
    def run: ConcolicResult =
        ConcolicResult(languages.toResultLanguages, runExtractSPC)

case class ConcolicQueryLanguages(smt: String, source: String):
    def toResultLanguages: ConcolicResultLanguages =
        ConcolicResultLanguages(smt)

class UnrecognizedSmtSort(sort: String)
    extends Exception(s"Unrecognized sort $sort")

case class ConcolicVarDesc(smt_name: String, value: String, sort: String, source_name: String):
    def getRepr: Any =
        sort match
            case "Int" => value.toInt
            case "Bool" =>
                value match
                    case "true" => true
                    case "false" => false
                    case _ => throw IllegalArgumentException(value)
            case _ => throw UnrecognizedSmtSort(sort)

    def toExprDeclare: SeirExpr.Decl =
        SeirExpr.Decl(source_name)

    def toExprDefine: SeirExpr.Def =
        SeirExpr.Def(
            source_name,
            SeirExpr.Re(SeirVal(
                getRepr,
                Map("smt" -> smt_name)
            ))
        )

case class ConcolicResult(languages: ConcolicResultLanguages, spc: SegmentedPathCond)

case class ConcolicResultLanguages(smt: String)
