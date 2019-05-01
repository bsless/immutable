package types

type ASeq struct {
	ISeq
	Sequential
	IHashEq
	// hash int
	// hasheq int
}

func (a ASeq) String() {}

func (a ASeq) Empty() IPersistentCollection {
	// return PersistentList.Empty
	return nil
}

func (a ASeq) Equiv(obj Any) bool {
	// TODO implement me
	return false
}

func (a ASeq) Equals(obj Any) bool {
	// TODO implement me
	return false
}

func (a ASeq) Count() int {
	// TODO implement me
	i := 1
	return i
}

func (a ASeq) Seq() ISeq {
	return a
}

func (a ASeq) First() Any {
	return a.Seq().First()
}

func (a ASeq) Next() ISeq {
	return a.Seq().Next()
}

func (a ASeq) More() ISeq {
	s := a.Next()
	if s == nil {
		return nil
	}
	return s
}

func (a ASeq) Cons(o Any) ISeq {
	cons := MakeConsCell(o, a)
	return cons
}

type ConsCell struct {
	ASeq
	first Any
	more ISeq
}

func MakeConsCell(o Any, more ISeq) *ConsCell {
	cell := new(ConsCell)
	cell.first = o
	cell.more = more
	return cell
}
