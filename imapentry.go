package types

import "github.com/bsless/immutable/map"

type IMapEntry interface {
	maps.Entry
	Key() Any
	Val() Any
}
