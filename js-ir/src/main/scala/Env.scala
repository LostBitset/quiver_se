import scala.collection.mutable.{
    ListBuffer => MutList, HashMap => MutMap, HashSet => MutSet
}

class SeirEnv(
    var stack: MutList[MutSet[String]] = MutList(), // which vars are defined in a given frame
    var vars: MutMap[String, MutList[SeirVal]] = MutMap() // the definitions of a var
):
    def enterScope: Unit =
        stack.addOne(MutSet.empty)
    
    def leaveScope: Unit =
        stack.last.foreach(dropped => {
            vars(dropped).dropRightInPlace(1)
            if vars(dropped).isEmpty then
                vars.remove(dropped)
        })
        stack.dropRightInPlace(1)

    def declare(key: String): Unit =
        stack.last.add(key)
    
    // Linear time for now, but this shouldn't really be a problem
    def define(key: String, value: SeirVal): Unit =
        val lastIndex = vars.size - 1
        if !(vars contains key) then
            vars
                .put(key, MutList(value))
        else
            vars(key)
                .patchInPlace(lastIndex, List(value), 1)

    def apply(key: String): SeirVal =
        vars(key).last

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
                        s"   └─ Key: $k\n"
                    else
                        s"   ├─ Key: $k\n"
                )
                + vars(k).zipWithIndex.map((item, j) =>
                    if (j + 1) == stack.length then
                        s"      └─ $item"
                    else
                        s"      ├─ $item"
                ).mkString("\n")
            ).mkString("\n")
        }
        """.stripMargin
    