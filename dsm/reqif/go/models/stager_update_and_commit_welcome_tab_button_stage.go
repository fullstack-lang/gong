package models

import (
	button "github.com/fullstack-lang/gong/lib/button/go/models"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
)

func (stager *Stager) UpdateAndCommitWelcomeTabButtonStage() {
	stage := stager.welcomeTabButtonStage
	stage.Reset()

	layout := new(button.Layout).Stage(stage)

	group1 := new(button.Group).Stage(stage)
	group1.Percentage = 100
	layout.Groups = append(layout.Groups, group1)

	buttonGeneratesModel := button.NewButton(
		// stager is the target of the button. stager implements interface method OnAfterUpdateButton()
		&GenerateModelButtonProxy{
			stager: stager,
		},
		"Generates go model",
		string(buttons.BUTTON_add_comment),
		"Generates go model",
	)

	group1.Buttons = append(group1.Buttons, buttonGeneratesModel)

	buttonExportStaticSite := button.NewButton(
		// stager is the target of the button. stager implements interface method OnAfterUpdateButton()
		&ExportStaticSiteButtonProxy{
			stager: stager,
		},
		"Export Static Web Site",
		string(buttons.BUTTON_web),
		"Export Static Web Site",
	)

	group1.Buttons = append(group1.Buttons, buttonExportStaticSite)

	if stager.pathToOutputReqifFile != "" {
		buttonExportModifiedReqif := button.NewButton(
			// stager is the target of the button. stager implements interface method OnAfterUpdateButton()
			&ExportModifiedReqifButtonProxy{
				stager: stager,
			},
			"Export Modified ReqIF",
			string(buttons.BUTTON_fact_check),
			"Export Modified ReqIF",
		)

		group1.Buttons = append(group1.Buttons, buttonExportModifiedReqif)
	}

	{
		buttonExportModifiedReqif := button.NewButton(
			// stager is the target of the button. stager implements interface method OnAfterUpdateButton()
			&ResetReqifButtonProxy{
				stager: stager,
			},
			"Reset ReqIF file",
			string(buttons.BUTTON_reset_tv),
			"Reset ReqIF file",
		)

		group1.Buttons = append(group1.Buttons, buttonExportModifiedReqif)
	}

	buttonExportRenderingCong := button.NewButton(
		&ExportRenderingConfButtonProxy{
			stager: stager,
		},
		"Export Rendering Configuration",
		string(buttons.BUTTON_fact_check),
		"Export Rendering Configuration",
	)

	group1.Buttons = append(group1.Buttons, buttonExportRenderingCong)

	stager.welcomeTabButtonStage.Commit()
}

type GenerateModelButtonProxy struct {
	stager *Stager
}

// GetButtonsStage implements models.Target.
func (e *GenerateModelButtonProxy) GetButtonsStage() *button.Stage {
	return e.stager.welcomeTabButtonStage
}

// OnAfterUpdateButton implements models.Target.
func (e *GenerateModelButtonProxy) OnAfterUpdateButton() {
	e.stager.modelGenerator.GenerateModels(e.stager)
}

type ExportStaticSiteButtonProxy struct {
	stager *Stager
}

// GetButtonsStage implements models.Target.
func (e *ExportStaticSiteButtonProxy) GetButtonsStage() *button.Stage {
	return e.stager.welcomeTabButtonStage
}

// OnAfterUpdateButton implements models.Target.
func (e *ExportStaticSiteButtonProxy) OnAfterUpdateButton() {

	e.stager.UpdateAndCommitSsgStage()
	e.stager.generatesSiteFromSSGStage()
}

type ExportModifiedReqifButtonProxy struct {
	stager *Stager
}

// GetButtonsStage implements models.Target.
func (e *ExportModifiedReqifButtonProxy) GetButtonsStage() *button.Stage {
	return e.stager.welcomeTabButtonStage
}

// OnAfterUpdateButton implements models.Target.
func (e *ExportModifiedReqifButtonProxy) OnAfterUpdateButton() {

	e.stager.reqifExporter.ExportReqif(e.stager)
}

type ResetReqifButtonProxy struct {
	stager *Stager
}

// GetButtonsStage implements models.Target.
func (proxy *ResetReqifButtonProxy) GetButtonsStage() *button.Stage {
	return proxy.stager.welcomeTabButtonStage
}

// OnAfterUpdateButton implements models.Target.
func (proxy *ResetReqifButtonProxy) OnAfterUpdateButton() {

	proxy.stager.stage.Reset()

	var req_if REQ_IF
	var coreContent A_CORE_CONTENT
	req_if.CORE_CONTENT = &coreContent
	var REQ_IF_CONTENT REQ_IF_CONTENT
	req_if.CORE_CONTENT.REQ_IF_CONTENT = &REQ_IF_CONTENT
	var A_DATATYPES A_DATATYPES
	req_if.CORE_CONTENT.REQ_IF_CONTENT.DATATYPES = &A_DATATYPES
	var A_SPEC_TYPES A_SPEC_TYPES
	req_if.CORE_CONTENT.REQ_IF_CONTENT.SPEC_TYPES = &A_SPEC_TYPES
	var A_SPEC_OBJECTS A_SPEC_OBJECTS
	req_if.CORE_CONTENT.REQ_IF_CONTENT.SPEC_OBJECTS = &A_SPEC_OBJECTS
	var A_SPEC_RELATIONS A_SPEC_RELATIONS
	req_if.CORE_CONTENT.REQ_IF_CONTENT.SPEC_RELATIONS = &A_SPEC_RELATIONS
	var A_SPECIFICATIONS A_SPECIFICATIONS
	req_if.CORE_CONTENT.REQ_IF_CONTENT.SPECIFICATIONS = &A_SPECIFICATIONS

	StageBranch(proxy.stager.stage, &req_if)
	proxy.stager.rootReqif = &req_if

	proxy.stager.stage.Commit()

	proxy.stager.GetSpecificationsTreeUpdater().UpdateAttributeDefinitionNb(proxy.stager)
	proxy.stager.GetSpecificationsTreeUpdater().UpdateAndCommitSpecificationsMarkdownStage(proxy.stager)
	proxy.stager.GetSpecificationsTreeUpdater().UpdateAndCommitSpecificationsTreeStage(proxy.stager)
	proxy.stager.GetSpecTypesTreeUpdater().UpdateAndCommitSpecTypesTreeStage(proxy.stager)
	proxy.stager.specObjectsUX.UpdateAndCommitSpecObjectsTreeStage(proxy.stager)
	proxy.stager.dataTypesTreeUpdater.UpdateAndCommitDataTypeTreeStage(proxy.stager)
	proxy.stager.specRelationsUX.UpdateAndCommitSpecRelationsTreeStage(proxy.stager)

}
