package probe

const ButtonImplGongstructFileTemplate = `// generated code - do not edit
package probe

import (

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	gongtree_models "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type ButtonImplGongstruct struct {
	gongStruct *gong_models.GongStruct
	Icon       gongtree_buttons.ButtonType
	probe *Probe
}

func NewButtonImplGongstruct(
	gongStruct *gong_models.GongStruct,
	icon gongtree_buttons.ButtonType,
	probe *Probe,
) (buttonImplGongstruct *ButtonImplGongstruct) {

	buttonImplGongstruct = new(ButtonImplGongstruct)
	buttonImplGongstruct.Icon = icon
	buttonImplGongstruct.gongStruct = gongStruct
	buttonImplGongstruct.probe = probe

	return
}

func (buttonImpl *ButtonImplGongstruct) ButtonUpdated(
	gongtreeStage *gongtree_models.Stage,
	stageButton, front *gongtree_models.Button) {

	// log.Println("ButtonImplGongstruct: ButtonUpdated")

	FillUpFormFromGongstructName(
		buttonImpl.probe,
		buttonImpl.gongStruct.Name,
		true,
	)
}
`
