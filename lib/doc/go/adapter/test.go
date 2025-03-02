package adapter

type Foo struct {
}

type Bar struct {
}

func (*Bar) Get() *Foo {
	return nil
}

type I_Venus1 interface {
}

type I_Venus2 interface {
	Get() I_Venus1
}

type I_Mars1 interface {
}

type I_Mars2 interface {
	Get() I_Mars1
}

var _ I_Venus1 = &(Foo{}) // Foo meets I_Venus1 interface
var _ I_Mars1 = &(Foo{})  // Foo meets I_Mars1 interface

// var _ I_Venus2 = &(Bar{}) // *Bar does not implement I_Venus2 (wrong type for method Get)
// var _ I_Mars2 = &(Bar{})  // *Bar does not implement I_Mars2 (wrong type for method Get)
