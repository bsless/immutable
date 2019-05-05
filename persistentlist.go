package types

type PersistentList struct {
	// interfaces
	IPersistentList
	IReduce
	// inheritence
	ASeq
	// members
	first Any
	rest IPersistentList
	count int
}

func (l PersistentList) First() Any {
	return l.first
}

func (l PersistentList) Next() ISeq {
	if count == 1 {
		return nil
	}
	return rest
}

func (l PersistentList) Peek() Any {
	return first
}

func (l PersistentList) Pop() IPersistentList {
	if rest == nil {
		return EMPTY.WithMeta(l.meta)
	}
	return l.rest
}

func (l PersistentList) Count() int {
	return l.count
}

func (l PersistentList) Empty() IPersistentCollection {
	return EMPTY.WithMeta(l.meta)
}

func NewPersistentList(meta IPersistentMap, first Any, rest IPersistentList, count int) *PersistentList {
	pl := &PersistentList{
		first: first,
		rest:  rest,
		count: count,
	}
	pl.meta = meta
	pl.ASeq.ISeq = pl
	pl.ASeq.IObj = pl
	pl.ASeq.IHashEq = pl
	return pl
}

func (l PersistentList) Cons(o Any) *PersistentList {
	return NewPersistentList(l.Meta(), o, l, l.count + 1)
}

// func Create(objs ...Any) IPersistentList {
// 	ret := EMPTY
// 	for obj := range objs {
// 		ret = ret.Cons(obj)
// 	}
// 	return ret
// }
