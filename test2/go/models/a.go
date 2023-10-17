package models

type A struct {
	Name string

	B *B

	Bs []*B
}

func (a *A) Foo() {
	// stage := gong.NewStage(path)
}
