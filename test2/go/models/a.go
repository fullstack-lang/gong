package models

type A struct {
	Name string

	Bs []*B
}

func (a *A) Foo() {
	// stage := gong.NewStage(path)
}
