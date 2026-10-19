// generated code - do not edit
package probe

func (probe *Probe) ux_form() {
}

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	_ = isNewInstance
	switch gongstructName {
	case "A", "models.A":
		StageSetNewInstance_A_Stage(probe)
	case "B", "models.B":
		StageSetNewInstance_B_Stage(probe)
	case "x.X":
		StageSetNewInstance_X_XStage(probe)
	case "y.Y":
		StageSetNewInstance_Y_YStage(probe)
	case "model.SubModel":
		StageSetNewInstance_SubModel_ModelStage(probe)
	}
}
