package data

import (
	"log"
	"strings"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongrouter_models "github.com/fullstack-lang/gongrouter/go/models"
	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

type ButtonImplGongstructLegacy struct {
	gongStruct      *gong_models.GongStruct
	Icon            gongtree_buttons.ButtonType
	gongrouterStage *gongrouter_models.StageStruct
	editorRouter    *gongrouter_models.Outlet
}

func NewButtonImplGongstructLegacy(
	gongStruct *gong_models.GongStruct,
	icon gongtree_buttons.ButtonType,
	gongrouterStage *gongrouter_models.StageStruct,
	editorRouter *gongrouter_models.Outlet,
) (buttonImplGongstructLegacy *ButtonImplGongstructLegacy) {

	buttonImplGongstructLegacy = new(ButtonImplGongstructLegacy)
	buttonImplGongstructLegacy.Icon = icon
	buttonImplGongstructLegacy.gongStruct = gongStruct
	buttonImplGongstructLegacy.editorRouter = editorRouter
	buttonImplGongstructLegacy.gongrouterStage = gongrouterStage

	return
}

func (buttonImplLegacy *ButtonImplGongstructLegacy) ButtonUpdated(
	gongtreeStage *gongtree_models.StageStruct,
	stageButton, front *gongtree_models.Button) {

	log.Println("ButtonImplGongstruct: ButtonUpdated")

	buttonImplLegacy.editorRouter.Path =
		"github_com_fullstack_lang_gong_test_go-" +
			strings.ToLower(buttonImplLegacy.gongStruct.Name) + "-adder"

	buttonImplLegacy.gongrouterStage.Commit()
}
