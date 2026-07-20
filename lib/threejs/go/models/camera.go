package models

// OnAfterUpdate, notice that camera == stagedCamera
func (camera *Camera) OnAfterUpdate(stage *Stage, _, frontCamera *Camera) {

	camera.Position = frontCamera.Position
	camera.TargetX = frontCamera.TargetX
	camera.TargetY = frontCamera.TargetY
	camera.TargetZ = frontCamera.TargetZ
	camera.Fov = frontCamera.Fov

	if camera.OnUpdate != nil {
		camera.OnUpdate(frontCamera)
	}
}
