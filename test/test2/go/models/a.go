package models

type A struct {
	Name string

	NumberField int

	B *B

	Bs []*B
}

func (a *A) Foo() {
	// stage := gong.NewStage(path)
}
