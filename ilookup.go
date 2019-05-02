package types

type ILookup interface {
	ValAt(key Any) Any
	ValAtDefault(key Any, notFound Any) Any
}
