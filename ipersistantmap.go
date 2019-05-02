package types

type IPersistentMap interface {
	// Iterable // TODO Add

	// Associative
	IPersistentCollection // Also Counted!
	ILookup
	CountainsKey(key Any) bool
	EntryAt(key Any) IMapEntry
	Assoc(key Any, val Any) IPersistentMap

	AssocEx(key Any, val Any) IPersistentMap
	Without(key Any) IPersistentMap
}
