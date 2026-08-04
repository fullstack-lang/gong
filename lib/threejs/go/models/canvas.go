package models

import "log"

func (canvas *Canvas) OnAfterUpdate(stage *Stage, _, frontCanvas *Canvas) {

	canvas.LastRendering = frontCanvas.LastRendering
	canvas.IsWithLastRenderingUpdate = frontCanvas.IsWithLastRenderingUpdate
	canvas.Frame64BitsEncoded = frontCanvas.Frame64BitsEncoded

	log.Printf("[Backend TRACE] Canvas OnAfterUpdate called! LastRendering: %s (IsWithLastRenderingUpdate: %v, Frame64BitsEncoded length: %d)",
		canvas.LastRendering.Format("15:04:05.000000"), canvas.IsWithLastRenderingUpdate, len(canvas.Frame64BitsEncoded))

	if canvas.OnUpdate != nil {
		canvas.OnUpdate(canvas)
	}
}
