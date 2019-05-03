package types

import "reflect"
import "fmt"

type EmptyList struct {
	IPersistentList
	ISeq
	Counted
	IHashEq
	Obj
	// hasheq int // TODO: implement
}

func (l EmptyList) HashCode() int {
	return 1
}

func (l EmptyList) HashEq() int {
	// TODO implement
	return 1
}

func (l EmptyList) String() string {
	return "()"
}

func (l EmptyList) First() Any {
	return nil
}

func (l EmptyList) Next() ISeq {
	return nil
}

func (l EmptyList) More() ISeq {
	return l.ISeq.More()
}

func (l EmptyList) Cons(o Any) *PersistentList {
	return MakePersistentList(l.Meta(), o, nil, 1)
}

func MakeEmptyList() *PersistentList {
	return &PersistentList{}
}

func MakeEmptyListWithMeta(meta IPersistentMap) *PersistentList {
	pl := &PersistentList{}
	pl.meta = meta
	return pl
}

type PersistentList struct {
	IPersistentList
	IReduce
	// ASeq
	ISeq
	Sequential
	IHashEq
	Obj
	hash int
	first Any
	rest *PersistentList
	count int
}

func MakePersistentList(meta IPersistentMap, first Any, rest *PersistentList, count int) *PersistentList {
	pl := &PersistentList{
		first: first,
		rest:  rest,
		count: count,
	}
	pl.meta = meta
	return pl
}

func Create(objs ...Any) *PersistentList {
	ret := MakeEmptyList()
	for i := len(objs)-1; i >=0; i-- {
		fmt.Println("Consing", objs[i])
		ret = ret.Cons(objs[i])
	}
	// for _, obj := range objs {
	// 	fmt.Println("Consing", obj)
	// 	ret = ret.Cons(obj)
	// }
	return ret
}

func (l PersistentList) Cons(o Any) *PersistentList {
	// ilist := l.IPersistentList
	return MakePersistentList(l.Meta(), o, &l, l.count + 1)
}

func (a PersistentList) HashCode() int {
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

func (a PersistentList) Empty() IPersistentCollection {
	// TODO implement me
	// return PersistentList.Empty
	return nil
}

func (a PersistentList) Equiv(obj Any) bool {
	// TODO implement me
	rt := reflect.TypeOf(obj)
	_, isSeq := obj.(Sequential)
	eq := false
	if !(rt.Kind() == reflect.Array || rt.Kind() == reflect.Slice || isSeq) {
		eq = false
	}
	return eq
}

func (a PersistentList) Equals(obj Any) bool {
	// TODO implement me
	return false
}

func (a PersistentList) Count() int {
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

func (a PersistentList) Seq() *PersistentList {
	return &a
}

func (a PersistentList) First() Any {
	return a.first
}

func (a PersistentList) Next() *PersistentList {
	return a.rest
}

func (a PersistentList) More() *PersistentList {
	s := a.Next()
	if s == nil {
		return nil
	}
	return s
}
