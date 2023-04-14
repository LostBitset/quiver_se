// For more information on writing tests, see
// https://scalameta.org/munit/docs/getting-started.html
class TestSuite extends munit.FunSuite:

  test("sanity check") {
    assertEquals(1 + 1, 2)
  }

  test("parses tokens") {
    val text = """
    |(scope
    |  (decl x)
    |  (decl y)
    |  (def x {int 1})
    |  (def y (.+ x x))
    |  (.if (.= y {int 3})
    |    (hidden <console.log("true")>)
    |    (hidden <console.log("false")>)))
    """.stripMargin
    val parser = SeirParser(text)
    val parsed = parser.takeExpr
    assertEquals(
      parsed,
      SeirExpr.Scope(
        List(
          SeirExpr.Decl("x"),
          SeirExpr.Decl("y"),
          SeirExpr.Def("x", SeirExpr.Re(SeirVal(1))),
          SeirExpr.Def(
            "y",
            SeirExpr.Call(
              SeirExpr.Var("+"),
              List(SeirExpr.Var("x"), SeirExpr.Var("x"))
            )
          ),
          SeirExpr.Call(
            SeirExpr.Var("if"),
            List(
              SeirExpr.Call(
                SeirExpr.Var("="),
                List(
                  SeirExpr.Var("y"),
                  SeirExpr.Re(SeirVal(1))
                )
              ),
              SeirExpr.Hidden("console.log(\"true\")"),
              SeirExpr.Hidden("console.log(\"false\")"),
            )
          )
        )
      )
    )
  }
