package models

type Cursor struct {
	Name string

	// where it starts and end
	StartX, EndX float64
	Y1, Y2       float64

	DurationSeconds float64

	Presentation

	IsPlaying bool
}

func (cursor *Cursor) PlayCursor(cursorStage *StageStruct, isPlaying bool) {
	cursor.IsPlaying = isPlaying
	cursorStage.Commit()
}
