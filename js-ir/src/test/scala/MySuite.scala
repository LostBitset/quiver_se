// For more information on writing tests, see
// https://scalameta.org/munit/docs/getting-started.html
class TestSuite extends munit.FunSuite:

  test("sanity check") {
    assertEquals(1 + 1, 2)
  } // */

  test("parses one token") {
    assertEquals(
      SeirParser("hello").takeToken,
      SeirTok.IdentLike("hello")
    )
    assertEquals(
      SeirParser(".").takeToken,
      SeirTok.CallHead
    )
    assertEquals(
      SeirParser(" \n  hello").takeToken,
      SeirTok.IdentLike("hello")
    )
    assertEquals(
      SeirParser(" \n   .").takeToken,
      SeirTok.CallHead
    )
  } // */

  test("parses captures") {
    assertEquals(
      SeirParser("   <something() with stuff >").takeToken,
      SeirTok.Hidden("something() with stuff ")
    )
  } // */

  test("parses multiple tokens") {
    val parser = SeirParser(" (.hi ) ")
    assertEquals(parser.takeToken, SeirTok.LParen)
    assertEquals(parser.takeToken, SeirTok.CallHead)
    assertEquals(parser.takeToken, SeirTok.IdentLike("hi"))
    assertEquals(parser.takeToken, SeirTok.RParen)
    assertEquals(parser.takeToken, SeirTok.EOF)
  } // */

  test("parses nested expressions at the token level") {
    val parser =  SeirParser("(.- (.+ a b) (.inc a))")
    assertEquals(parser.takeToken, SeirTok.LParen)
    assertEquals(parser.takeToken, SeirTok.CallHead)
    assertEquals(parser.takeToken, SeirTok.IdentLike("-"))
    assertEquals(parser.takeToken, SeirTok.LParen)
    assertEquals(parser.takeToken, SeirTok.CallHead)
    assertEquals(parser.takeToken, SeirTok.IdentLike("+"))
    assertEquals(parser.takeToken, SeirTok.IdentLike("a"))
    assertEquals(parser.takeToken, SeirTok.IdentLike("b"))
    assertEquals(parser.takeToken, SeirTok.RParen)
    assertEquals(parser.takeToken, SeirTok.LParen)
    assertEquals(parser.takeToken, SeirTok.CallHead)
    assertEquals(parser.takeToken, SeirTok.IdentLike("inc"))
    assertEquals(parser.takeToken, SeirTok.IdentLike("a"))
    assertEquals(parser.takeToken, SeirTok.RParen)
    assertEquals(parser.takeToken, SeirTok.RParen)
    assertEquals(parser.takeToken, SeirTok.EOF)
  } // */

  test ("parses empty string at the token level") {
    assertEquals(
      SeirParser("").takeToken,
      SeirTok.EOF
    )
  }

  test("parses expressions") {
    val text = """
    |(scope
    |  (decl x)
    |  (decl y)
    |  (def inc ~(.+ ~1 {int 1}))
    |  (def x {int 1})
    |  (def y (.inc x))
    |  (.if (.= y {int 3})
    |    (hidden <console.log("true")>)
    |    (hidden <console.log("false")>)))
    """.stripMargin
    val parser = SeirParser(text)
    val parsed = parser.takeExpr.get
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
                  SeirExpr.Re(SeirVal(3))
                )
              ),
              SeirExpr.Hidden("console.log(\"true\")"),
              SeirExpr.Hidden("console.log(\"false\")"),
            )
          )
        )
      )
    )
  } // */
