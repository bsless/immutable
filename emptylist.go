package types

type EmptyList struct {
	// interfaces
	IPersistentList
	ISeq
	Counted
	IHashEq
	// inheritence
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

func (l EmptyList) Next() ISeq {
	return nil
}

func (l EmptyList) More() ISeq {
	return l.ISeq.More()
}

func (l EmptyList) Cons(o Any) *PersistentList {
	return NewPersistentList(l.Meta(), o, nil, 1)
}

func (l EmptyList) Count() int {
	return 0
}

func (l EmptyList) Empty() IPersistentCollection {
	return l
}

// TODO Implement
func (l EmptyList) Equals(o Any) bool {
	return false
}

func (l EmptyList) Equiv(o Any) bool {
	return l.Equals(o)
}

func (l EmptyList) Seq() ISeq {
	return nil
}

func (l EmptyList) WithMeta(meta IPersistentMap) *EmptyList {
	if meta != l.Meta() {
		return NewEmptyListWithMeta(meta)
	}
	return &l
}

func NewEmptyList() *EmptyList {
	pl := &EmptyList{}
	pl.Obj.IObj = pl
	pl.ISeq = pl
	pl.IPersistentList = pl
	pl.IHashEq = pl
	return pl
}

func NewEmptyListWithMeta(meta IPersistentMap) *EmptyList {
	pl := NewEmptyList()
	pl.meta = meta
	return pl
}

var EMPTY = NewEmptyList()
