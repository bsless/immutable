package types

type EmptyList struct {
	IPersistentList
	ISeq
	Counted
	IHashEq
	Obj
	// hasheq int // TODO: implement
}

func (l EmptyList) HashCode() int {
	return 1
}

func (l EmptyList) HashEq() int {
	// TODO implement
	return 1
}

func (l EmptyList) String() string {
	return "()"
}

func (l EmptyList) First() Any {
	return nil
}

func (l EmptyList) next() ISeq {
	return nil
}

func (l EmptyList) More() ISeq {
	return l
}

func (l EmptyList) Cons(o Any) *PersistentList {
	return MakePersistentList(l.Meta(), o, nil, 1)
}

func MakeEmptyList() *PersistentList {
	return &PersistentList{}
}

// TODO change meta type to map
func MakeEmptyListWithMeta(meta IPersistentMap) *PersistentList {
	return &PersistentList{meta: meta}
}

type PersistentList struct {
	IPersistentList
	IReduce
	ASeq
	first Any
	rest IPersistentList
	count int
}

func MakePersistentList(meta IPersistentMap, first Any, rest IPersistentList, count int) *PersistentList {
	return &PersistentList{
		first: first,
		rest: rest,
		meta: meta,
		count: count,
	}
}

func (l PersistentList) Cons(o Any) *PersistentList {
	var ilist IPersistentList
	ilist = &l
	return MakePersistentList(l.Meta(), o, ilist, l.count + 1)
}

func Create(objs ...Any) *PersistentList {
	ret := MakeEmptyList()
	for obj := range(objs) {
		ret = ret.Cons(obj)
	}
	return ret
}
