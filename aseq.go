package types

import "reflect"

type ASeq struct {
	ISeq
	Sequential
	IHashEq
	Obj
	hash int
	// hasheq int
}

// TODO implement
func (a ASeq) String() {}

func (a ASeq) HashCode() int {
	// TODO make use of golang's hashing, hashCode is implemented by any Java object
	if a.hash == 0 {
		hash := 1
		for s := a.Seq(); s != nil; s = s.Next() {
			rest := 0
			// if s.First() != nil {
			// 	rest = s.First().HashCode()
			// }
			hash = 31*hash + rest
		}
		a.hash = hash
	}
	return a.hash
}

func (a ASeq) Empty() IPersistentCollection {
	// TODO implement me
	// return PersistentList.Empty
	return nil
}

func (a ASeq) Equiv(obj Any) bool {
	// TODO implement me
	rt := reflect.TypeOf(obj)
	_, isSeq := obj.(Sequential)
	eq := false
	if !(rt.Kind() == reflect.Array || rt.Kind() == reflect.Slice || isSeq) {
		eq = false
	}
	return eq
}

func (a ASeq) Equals(obj Any) bool {
	// TODO implement me
	return false
}

func (a ASeq) Count() int {
	i := 1
	for s := a.Next(); s != nil; s, i = s.Next(), i+1 {
		var u interface{} = s
		_, ok := u.(Counted)
		if ok {
			return i + s.Count()
		}
	}
	return i
}

func (a ASeq) Seq() ISeq {
	return a
}

func (a ASeq) Cons(o Any) ISeq {
	cons := NewCons(o, a)
	return cons
}

func (a ASeq) More() ISeq {
	s := a.Next()
	if s == nil {
		return nil // TODO empty list
	}
	return s
}

// func (a ASeq) First() Any {
// 	return a.Seq().First()
// }

// func (a ASeq) Next() ISeq {
// 	return a.Seq().Next()
// }
