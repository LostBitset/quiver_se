import scala.collection.mutable.{
    ListBuffer => MutList, HashMap => MutMap, HashSet => MutSet
}

class SeirEnv(
    var stack: MutList[MutSet[String]], // which vars are defined in a given frame
    var vars: MutMap[String, MutList[SeirVal]] // the definitions of a var
):
    def enterScope: Unit =
        stack.addOne(MutSet.empty)
    
    def leaveScope: Unit =
        stack.dropRightInPlace(1)(0).foreach({
            vars(_).dropRightInPlace(1)
        })
