package types

type Associative interface {
	IPersistentCollection
	ILookup
	CountainsKey(key Any) bool
	EntryAt(key Any) IMapEntry
	Assoc(key Any, val Any) Associative
}
