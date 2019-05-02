package types

type IObj interface {
	IMeta
	WithMeta(meta IPersistentMap) IObj
}
