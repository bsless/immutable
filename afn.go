package types

type AFn struct{}

func (f AFn) Invoke(args ...Any) Any {
	panic("ArityException")
}

func (f AFn) Run() {
	f.Invoke()
}

func (f AFn) Call() Any {
	return f.Invoke()
}

func ApplyToHelper(ifn IFn, arglist ISeq) Any {
	// TODO implement me
	return nil
}

func (f AFn) ApplyTo(arglist ISeq) Any {
	return ApplyToHelper(f, arglist)
}
