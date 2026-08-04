package models

// OnAfterUpdate, notice that camera == stagedCamera
func (camera *Camera) OnAfterUpdate(stage *Stage, _, frontCamera *Camera) {

	camera.X = frontCamera.X
	camera.Y = frontCamera.Y
	camera.Z = frontCamera.Z
	camera.TargetX = frontCamera.TargetX
	camera.TargetY = frontCamera.TargetY
	camera.TargetZ = frontCamera.TargetZ
	camera.Fov = frontCamera.Fov

	if camera.OnUpdate != nil {
		camera.OnUpdate(frontCamera)
	}

	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
}
