package types

type Obj struct {
	IObj
	// TODO add Serializable
	meta IPersistentMap
}

func MakeObj() *Obj {
	return &Obj{meta: nil}
}

func MakeObjWithMeta(meta IPersistentMap) *Obj {
	return &Obj{meta: meta}
}

func (o Obj) Meta() IPersistentMap {
	return o.meta
}
