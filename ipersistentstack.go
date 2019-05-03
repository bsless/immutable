package types

type IPersistentStack interface {
	IPersistentCollection
	Peek() Any
	Pop() IPersistentStack
}
