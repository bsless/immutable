package types

type IObj interface {
	IMeta
	WithMeta(meta IPersistentMap) IObj // TODO change meta type to IPersistentMap
}
