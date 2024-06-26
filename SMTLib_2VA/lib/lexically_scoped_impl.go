package smtlib2va

import (
	"github.com/sirupsen/logrus"
)

// lvbls = lo vanbi be lo snicne = the environment of the variables

func NewLexicallyScoped() (lvbls LexicallyScoped) {
	lvbls = LexicallyScoped{
		NewSliceStack[*[]Var](),
		make(map[string]*SliceStack[LexicallyScopedIndex]),
	}
	return
}

func (lvbls *LexicallyScoped) EnterScope() {
	backing := make([]Var, 0)
	lvbls.stack.Push(&backing)
}

func (lvbls *LexicallyScoped) LeaveScope() {
	vars := lvbls.stack.Peek()
	for _, v := range *vars {
		if lvbls.names[v.name].Length() < 2 {
			delete(lvbls.names, v.name)
		} else {
			lvbls.names[v.name].SilentPop()
		}
	}
	lvbls.stack.SilentPop()
}

func (lvbls *LexicallyScoped) WriteIndexTrusting(index LexicallyScopedIndex, val string) {
	slot_ref := &(*lvbls.stack.Index(index.stack_index))[index.frame_index].slot
	slot_ref.Write(val)
}

func (lvbls *LexicallyScoped) SetVar(name string, val *string) {
	scope_ref := lvbls.stack.Peek()
	stack_index := lvbls.stack.Length() - 1
	frame_index := len(*scope_ref)
	if stack_ref, ok := lvbls.names[name]; ok && val != nil {
		index := stack_ref.Peek()
		lvbls.WriteIndexTrusting(index, *val)
	} else {
		if val == nil {
			*scope_ref = append(*scope_ref, Var{name, NewVarSlot()})
		} else {
			slot := NewVarSlot()
			slot.Write(*val)
			*scope_ref = append(*scope_ref, Var{name, slot})
		}
		backing := NewSliceStack[LexicallyScopedIndex]()
		backing.Push(
			LexicallyScopedIndex{
				stack_index,
				frame_index,
			},
		)
		lvbls.names[name] = &backing
	}
}

func (lvbls *LexicallyScoped) DeclVar(name string) {
	lvbls.SetVar(name, nil)
	logrus.Infof(
		"[LexicallyScoped] Declared variable \"%s\".",
		name,
	)
}

func (lvbls *LexicallyScoped) WriteVar(name string, val string) {
	lvbls.SetVar(name, &val)
	logrus.Infof(
		"[LexicallyScoped] Wrote variable \"%s\".",
		name,
	)
}

func (lvbls LexicallyScoped) IsDefined(name string) (defined bool) {
	_, ok_name := lvbls.names[name]
	defined = ok_name
	return
}

func (lvbls LexicallyScoped) ReadVarSafe(name string) (val string, ok bool) {
	ok = lvbls.IsDefined(name)
	if !ok {
		return
	}
	index := lvbls.names[name].Peek()
	val = lvbls.IndexReadTrusting(index)
	return
}

func (lvbls LexicallyScoped) IndexReadTrusting(index LexicallyScopedIndex) (val string) {
	v := (*lvbls.stack.Index(index.stack_index))[index.frame_index]
	val = v.slot.Read()
	return
}

func (lvbls LexicallyScoped) ReadVar(name string) (val string) {
	val_maybe, ok := lvbls.ReadVarSafe(name)
	if !ok {
		val = "@+__2VA_UNKNOWN__+"
	}
	val = val_maybe
	return
}
