package golang

const ButtonImplGongstructFileTemplate = `package data

import (
	"log"
	"strings"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongrouter_models "github.com/fullstack-lang/gongrouter/go/models"
	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type ButtonImplGongstruct struct {
	gongStruct      *gong_models.GongStruct
	Icon            gongtree_buttons.ButtonType
	gongrouterStage *gongrouter_models.StageStruct
	editorRouter    *gongrouter_models.Outlet
}

func NewButtonImplGongstruct(
	gongStruct *gong_models.GongStruct,
	icon gongtree_buttons.ButtonType,
	gongrouterStage *gongrouter_models.StageStruct,
	editorRouter *gongrouter_models.Outlet,
) (buttonImplGongstruct *ButtonImplGongstruct) {

	buttonImplGongstruct = new(ButtonImplGongstruct)
	buttonImplGongstruct.Icon = icon
	buttonImplGongstruct.gongStruct = gongStruct
	buttonImplGongstruct.editorRouter = editorRouter
	buttonImplGongstruct.gongrouterStage = gongrouterStage

	return
}

func (buttonImpl *ButtonImplGongstruct) ButtonUpdated(
	gongtreeStage *gongtree_models.StageStruct,
	stageButton, front *gongtree_models.Button) {

	log.Println("ButtonImplGongstruct: ButtonUpdated")

	buttonImpl.editorRouter.Path =
		"{{PkgPathRootWithoutSlashes}}-" +
			strings.ToLower(buttonImpl.gongStruct.Name) + "-adder"

	buttonImpl.gongrouterStage.Commit()
}
`
