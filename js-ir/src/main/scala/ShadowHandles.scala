import scala.collection.mutable.{ListBuffer => MutList}

given ShadowHandles = ShadowHandles(
    TiedMap[ShadowOpSpec]()
        + (ShadowOpSpec("smt", ShadowOp.Promote), repr => {
            repr match
                case int: Int => Some(int.toString)
                case _ => None
        })
        + (ShadowOpSpec("smt", ShadowOp.OnEventTransition), (event, ctx) => {
            if !(ctx.map contains "path-cond") then
                ctx.map += ("path-cond", MutList.empty[String])
            ctx.map("path-cond")
                .asInstanceOf[MutList[String]]
                .addOne(s"@@MAGIC:event-transition=$event")
        })
        + (ShadowOpSpec("smt", ShadowOp.Named("+")), (args, ctx) => {
            val spaceSep = args.map(_.shadow).asInstanceOf[List[String]].mkString(" ")
            s"(+ ${spaceSep})"
        })
        + (ShadowOpSpec("smt", ShadowOp.Named("int=")), (args, ctx) => {
            val spaceSep = args.map(_.shadow).asInstanceOf[List[Int]].mkString(" ")
            s"(= ${spaceSep})"
        })
        + (ShadowOpSpec("smt", ShadowOp.NamedRelaxed("if")), (args, ctx) => {
            args(0) match
                case Some(cond) =>
                    val cVal = cond.repr.asInstanceOf[Boolean]
                    val cSym = cond.shadow.asInstanceOf[String]
                    val pathCondition =
                        if cVal then
                            cSym
                        else
                            s"(not ${cSym})"
                    if !(ctx.map contains "path-cond") then
                        ctx.map += ("path-cond", MutList.empty[String])
                    ctx.map("path-cond")
                        .asInstanceOf[MutList[String]]
                        .addOne(pathCondition)
                case None => ()
            None
        })
        + (ShadowOpSpec("smt", ShadowOp.Named("if")), (args, ctx) => {
            val cond = args(0).repr.asInstanceOf[Boolean]
            if cond then args(1).shadow else args(2).shadow
        })
)
