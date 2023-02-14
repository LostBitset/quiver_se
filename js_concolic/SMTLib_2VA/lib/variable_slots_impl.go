package smtlib2va

func NewVarSlot() (slot VarSlot) {
	slot = VarSlot{nil}
	return
}

func (slot *VarSlot) Write(val string) {
	if slot.value == nil {
		slot.value = &val
	} else {
		*slot.value = val
	}
}

func (slot VarSlot) Read() (val string) {
	val = *slot.value
	return
}
