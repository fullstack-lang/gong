package x

type StageStruct struct {
	X_As           map[*X_A]any
	X_As_mapString map[string]*X_A
}

func NewStage(path string) (stage *StageStruct) {

	return new(StageStruct)
}
