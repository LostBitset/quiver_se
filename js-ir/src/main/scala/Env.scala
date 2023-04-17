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
        stack.dropRightInPlace(1).last.foreach({
            vars(_).dropRightInPlace(1)
        })

    def declare(key: String): Unit =
        stack.last.add(key)
    
    // Linear time for now, but this shouldn't really be a problem
    def define(key: String, value: SeirVal): Unit =
        val lastIndex = vars.size - 1
        vars(key).patchInPlace(lastIndex, List(value), 1)

    def apply(key: String): SeirVal =
        vars(key).last
