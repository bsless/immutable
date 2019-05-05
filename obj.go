package types

type Obj struct {
	IObj
	// TODO add Serializable
	meta IPersistentMap
}

func NewObj() *Obj {
	return &Obj{meta: nil}
}

func NewObjWithMeta(meta IPersistentMap) *Obj {
	return &Obj{meta: meta}
}

func (o Obj) Meta() IPersistentMap {
	return o.meta
}
