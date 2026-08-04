package models

import "log"

func (canvas *Canvas) OnAfterUpdate(stage *Stage, _, frontCanvas *Canvas) {

	canvas.LastRendering = frontCanvas.LastRendering
	canvas.IsWithLastRenderingUpdate = frontCanvas.IsWithLastRenderingUpdate

	log.Printf("[Backend TRACE] Canvas OnAfterUpdate called! LastRendering: %s (IsWithLastRenderingUpdate: %v)", canvas.LastRendering.Format("15:04:05.000000"), canvas.IsWithLastRenderingUpdate)

	if canvas.OnUpdate != nil {
		canvas.OnUpdate(canvas)
	}
}
