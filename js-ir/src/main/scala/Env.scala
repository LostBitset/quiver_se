import scala.collection.mutable.{
    ListBuffer => MutList, HashMap => MutMap, HashSet => MutSet
}

def formatCondRaw(constraint: String): String =
    s"${SMT_COND_MAGIC}t;@__RAW__$constraint"

object SeirEnv:
    def apply(): SeirEnv =
        SeirEnv(ShadowCtx(), ShadowHandles(TiedMap()))

case class SeirEnv(
    shadowCtx: ShadowCtx,
    shadowHandles: ShadowHandles,
    var stack: MutList[MutSet[String]] = MutList(), // which vars are defined in a given frame
    var vars: MutMap[String, MutList[SeirVal]] = MutMap() // the definitions of a var
):
    def denoteRaw2VA(s: String): Unit =
        val constraint = formatCondRaw(s)
        if !(shadowCtx.map contains "path-cond") then
            shadowCtx.map += ("path-cond", MutList.empty[String])
        shadowCtx.map("path-cond")
            .asInstanceOf[MutList[String]]
            .addOne(constraint)

    def enterScope: Unit =
        stack.addOne(MutSet.empty)
        denoteRaw2VA("(*/enter-scope/*)")
    
    def leaveScope: Unit =
        stack.last.foreach(dropped => {
            vars(dropped).dropRightInPlace(1)
            if vars(dropped).isEmpty then
                vars.remove(dropped)
        })
        stack.dropRightInPlace(1)
        denoteRaw2VA("(*/leave-scope/*)")

    def declare(key: String, collapseSMT: Boolean = true): Unit =
        println(s"$key -> $collapseSMT")
        stack.last.add(key)
        if collapseSMT then
            denoteRaw2VA(s"(*/decl-var/* **seirVar_$key)")
    
    // Linear time for now, but this shouldn't really be a problem
    def define(key: String, value: SeirVal, collapseSMT: Boolean = true): Unit =
        val toWrite =
            shadowHandles.extract("smt")(value) match
                case Some(extractedSMT) if collapseSMT =>
                    val smtReferenceEncoding = s"(*/read-var/* **seirVar_$key)";
                    denoteRaw2VA(s"(*/write-var/* **seirVar_$key *{{$extractedSMT}}*)")
                    SeirVal(
                        value.repr,
                        value.shadows + ("smt" -> smtReferenceEncoding)
                    )
                case _ =>
                    value
        val lastIndex = vars.size - 1
        if !(vars contains key) then
            vars
                .put(key, MutList(toWrite))
        else
            vars(key)
                .patchInPlace(lastIndex, List(toWrite), 1)

    def apply(key: String): SeirVal =
        val assignments = vars(key).toList
        assignments.last

    def isDefined(key: String): Boolean =
        vars contains key
    
    override def toString: String =
        s"""
        |SeirEnv@${hashCode}
        |├─ Stack
        |${
            stack.zipWithIndex.map((x, i) =>
                (
                    if (i + 1) == stack.length then
                        s"   └─ (Frame $i)\n"
                    else
                        s"   ├─ (Frame $i)\n"
                )
                + x.zipWithIndex.map((item, j) =>
                    if (j + 1) == stack.length then
                        s"      └─ $item"
                    else
                        s"      ├─ $item"
                ).mkString("\n")
            ).mkString("\n")
        }
        |└─ Vars (bindings)
        |${
            vars.keys.toList.zipWithIndex.map((k, i) =>
                (
                    if (i + 1) == vars.size then
                        s"   └─ List for $k\n"
                    else
                        s"   ├─ List for $k\n"
                )
                + vars(k).zipWithIndex.map((item, j) =>
                    if (j + 1) == vars.size then
                        s"      └─ $item"
                    else
                        s"      ├─ $item"
                ).mkString("\n")
            ).mkString("\n")
        }
        """.stripMargin
    