package models

func (canvas *Canvas) OnAfterUpdate(stage *Stage, _, frontCanvas *Canvas) {

	if canvas.OnUpdate != nil {
		canvas.OnUpdate(frontCanvas)
	}
}
