package types

type IReduceInit interface {
	ReduceInit(f IFn, start Any) Any
}

type IReduce interface {
	IReduceInit
	Reduce(f IFn) Any
}

