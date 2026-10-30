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
	case "Catalog", "models.Catalog":
		StageSetNewInstance_Catalog_Stage(probe)
	case "DerivedItemType", "models.DerivedItemType":
		StageSetNewInstance_DerivedItemType_Stage(probe)
	case "refs.RefTargetType":
		StageSetNewInstance_RefTargetType_RefsStage(probe)
	case "common.BaseItemType":
		StageSetNewInstance_BaseItemType_CommonStage(probe)
	}
}
