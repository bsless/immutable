package types

type ISeq interface {
	Seqable
	Count() int
	Empty() IPersistentCollection
	Equiv(o Any) bool
	First() Any
	Next() ISeq
	More() ISeq
	Cons(o Any) ISeq
}

