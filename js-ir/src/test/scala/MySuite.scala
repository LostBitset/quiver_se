import scala.collection.mutable.{ListBuffer => MutList}

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
    |  (decl inc)
    |  (def inc ~(.+ ~#1 {int 1}))
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
          SeirExpr.Decl("inc"),
          SeirExpr.Def("inc", SeirExpr.Capture(
            SeirExpr.Call(
              SeirExpr.Var("+"),
              List(
                SeirExpr.ArgRef(1),
                SeirExpr.Re(SeirVal(1))
              )
            )
          )),
          SeirExpr.Def("x", SeirExpr.Re(SeirVal(1))),
          SeirExpr.Def(
            "y",
            SeirExpr.Call(
              SeirExpr.Var("inc"),
              List(SeirExpr.Var("x"))
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

  test("prelude transform works") {
    val prelude = SeirPrelude(List(
      SeirExpr.Decl("a"),
      SeirExpr.Decl("b")
    ))
    val transformed = prelude.transform(
      SeirExpr.Scope(List(
        SeirExpr.Decl("w")
      ))
    )
    assertEquals(
      transformed,
      SeirExpr.Scope(List(
        SeirExpr.Decl("a"),
        SeirExpr.Decl("b"),
        SeirExpr.Scope(List(
          SeirExpr.Decl("w")
        ))
      ))
    )
  } // */

  test("env works properly 1/2") {
    var env = SeirEnv()
    env.enterScope
    env.declare("test")
    env.enterScope
    env.enterScope
    env.define("test", SeirVal("hello"))
    env.leaveScope
    assert(env.isDefined("test"))
    assertEquals(
      env("test"),
      SeirVal("hello")
    )
    env.leaveScope
    env.leaveScope
    env.enterScope
    assert(!env.isDefined("test"))
    env.enterScope
  } // */

  test("env works properly 2/2") {
    var env = SeirEnv()
    env.enterScope
    env.declare("f")
    env.enterScope
    env.define("f", SeirVal(false))
    env.leaveScope
    env.enterScope
    env.declare("m")
    env.define("m", SeirVal(42))
    env.declare("y")
    env.define("y", SeirVal("something"))
    assert(env.isDefined("f"))
    assert(env.isDefined("m"))
    assert(env.isDefined("y"))
    assertEquals(
      env("f"),
      SeirVal(false)
    )
    assertEquals(
      env("m"),
      SeirVal(42)
    )
    assertEquals(
      env("y"),
      SeirVal("something")
    )
    env.leaveScope
    env.leaveScope
    assert(!env.isDefined("f"))
    assert(!env.isDefined("m"))
    assert(!env.isDefined("y"))
  } // */

  test("evaluation") {
    val text = """
    |(scope
    |  (decl x)
    |  (decl y)
    |  (decl inc)
    |  (def inc ~(.+ ~#0 {int 1}))
    |  (def x {int 5})
    |  (def y (.inc x))
    |  (scope x y))
    """.stripMargin
    val parser = SeirParser(text)
    val expr = parser.takeExpr.get
    assertEquals(
      evalSeir(expr),
      SeirVal(6, Map("@@varname" -> "y"))
    )
  } // */

  test("evaluation with nested functions") {
    val text = """
    |(scope
    |  (decl ap2)
    |  (decl one)
    |  (decl inc)
    |  (def ap2 ~(.~#0 ~#1 ~#2))
    |  (def one {int 1})
    |  (def inc
    |    ~(.ap2
    |      ~(.+ ~#0 ~#1)
    |      ~#0
    |      one))
    |  (.inc (.inc one)))
    """.stripMargin
    val parser = SeirParser(text)
    val expr = parser.takeExpr.get
    assertEquals(
      evalSeir(expr).repr,
      3
    )
  } // */

  test("simple shadow extraction") {
    val extracted = summon[ShadowHandles].extract("smt")(SeirVal(7))
    assertEquals(extracted, Some("7"))
  } // */

  test("evaluation with shadows directly") {
    val text = """
    |(scope
    |  (.+ X {int 1}))
    """.stripMargin
    val parser = SeirParser(text)
    val exprNoContext = parser.takeExpr.get
    val customPrelude = SeirPrelude(List(
      SeirExpr.Decl("X"),
      SeirExpr.Def("X", SeirExpr.Re(SeirVal(
        5,
        Map("smt" -> "X")
      )))
    ))
    val expr = customPrelude.transform(exprNoContext)
    assertEquals(
      evalSeir(expr),
      SeirVal(
        6,
        Map("smt" -> "(+ X 1)")
      )
    )
  } // */

  test("evaluation with shadows indirectly") {
    val text = """
    |(scope
    |  (decl inc)
    |  (def inc ~(.+ ~#0 {int 1}))
    |  (def x (scope X))
    |  (scope x (.inc x)))
    """.stripMargin
    val parser = SeirParser(text)
    val exprNoContext = parser.takeExpr.get
    val customPrelude = SeirPrelude(List(
      SeirExpr.Decl("X"),
      SeirExpr.Def("X", SeirExpr.Re(SeirVal(
        5,
        Map("smt" -> "X")
      )))
    ))
    val expr = customPrelude.transform(exprNoContext)
    assertEquals(
      evalSeir(expr),
      SeirVal(
        6,
        Map("smt" -> "(+ X 1)")
      )
    )
  } // */

  test("evaluation with if") {
    val exprT =
      SeirParser("(.if true {int 1} {int 4})")
        .takeExpr
        .get
    val exprF =
      SeirParser("(.if false {int 1} {int 4})")
        .takeExpr
        .get
    assertEquals(evalSeir(exprT).repr, 1)
    assertEquals(evalSeir(exprF).repr, 4)
  } // */

  test("path conditions") {
    val text = """
    |(scope
    |  (.if X
    |    (.if Y
    |      {int 2}
    |      {int 3})
    |    {int 2}))
    """.stripMargin
    val parser = SeirParser(text)
    val exprNoContext = parser.takeExpr.get
    val customPrelude = SeirPrelude(List(
      SeirExpr.Decl("X"),
      SeirExpr.Def("X", SeirExpr.Re(SeirVal(
        true,
        Map("smt" -> "X")
      ))),
      SeirExpr.Decl("Y"),
      SeirExpr.Def("Y", SeirExpr.Re(SeirVal(
        false,
        Map("smt" -> "Y")
      )))
    ))
    val expr = customPrelude.transform(exprNoContext)
    val evaluator = SeirEvaluator()
    evaluator.evalSeir(expr)
    assertEquals(
      evaluator
        .shadowCtx
        .map
        .get("path-cond")
        .get
        .asInstanceOf[MutList[String]]
        .toList,
      List("(not Y)", "X")
    )
  } // */
