package types

type Cons struct {
	ASeq
	first Any
	more  ISeq
}

func MakeCons(o Any, more ISeq) *Cons {
	cell := new(Cons)
	cell.first = o
	cell.more = more
	return cell
}

func (c Cons) First() Any {
	return c.first
}

func (c Cons) More() ISeq {
	return c.more
}

func (c Cons) Next() ISeq {
	return c.More().Seq()
}

func (c Cons) Count() int {
	// TODO implement me
	return 1
}

func (c Cons) WithMeta(meta IPersistentMap) *Cons {
	// TODO implement me
	return MakeCons(c.first, c.more)
}
