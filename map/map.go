package maps

type Entry interface {
	GetKey() interface{}
	GetValue() interface{}
	SetValue(v interface{}) interface{}
	Equals(o interface{}) bool
	HashCode() int
}
