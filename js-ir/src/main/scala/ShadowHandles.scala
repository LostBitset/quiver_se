given ShadowHandles = ShadowHandles(
    TiedMap[ShadowOpSpec]()
        + (ShadowOpSpec("smt", ShadowOp.Promote), repr => {
            repr match
                case int: Int => Some(int.toString)
                case _ => None
        })
        + (ShadowOpSpec("smt", ShadowOp.Named("+")), (args, ctx) => {
            val spaceSep = args.map(_.shadow).asInstanceOf[List[String]].mkString(" ")
            s"(+ ${spaceSep})"
        })
)
