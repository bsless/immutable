package types

type IPersistentCollection interface {
	Seqable
	Counted
	Cons(o Any) IPersistentCollection
	Empty() IPersistentCollection
	Equiv(o Any) bool
}
