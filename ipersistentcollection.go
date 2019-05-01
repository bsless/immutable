package types

type IPersistentCollection interface {
	Seqable
	Count() int
	Cons(o Any) IPersistentCollection
	Empty() IPersistentCollection
	Equiv(o Any) bool
}
