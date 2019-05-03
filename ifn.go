package types

type IFn interface {
	Call() Any // TODO: figure this out. Golang func interface?
	Run()      // TODO: figure this out. Golang func interface?
	Invoke(args ...Any) Any
}
