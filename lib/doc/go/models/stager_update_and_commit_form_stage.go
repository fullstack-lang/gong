package models

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"
)

func (stager *Stager) UpdateAndCommitFormStage() {

	stager.formStage.Reset()

	var selectedDiagram *Classdiagram

	var diagramPackage *DiagramPackage
	for diagramPackage = range *GetGongstructInstancesSet[DiagramPackage](stager.stage) {

		selectedDiagram = diagramPackage.SelectedClassdiagram

		// if no class diagram is selected generate a blank diagram
		if selectedDiagram == nil {
			selectedDiagram = new(Classdiagram)
		}
	}
	if diagramPackage == nil {
		return
	}

	formGroup := &form.FormGroup{
		Name: stager.formStage.GetName(),
		FormDivs: []*form.FormDiv{
			{
				Name: "Edit Classdiagram",
				FormFields: []*form.FormField{
					{
						FormFieldString: &form.FormFieldString{
							Name:       "Classdiagram edit",
							Value:      selectedDiagram.Description,
							IsTextArea: true,
						},
						Placeholder:      "Description of the package",
						HasBespokeHeight: true,
						BespokeHeightPx:  200,
						HasBespokeWidth:  true,
						BespokeWidthPx:   800,
					},
				},
			},
		},
	}

	formGroup.OnSave = &ClassDiagramFormCallback{
		stager:       stager,
		formGroup:    formGroup,
		classdiagram: selectedDiagram,
	}

	form.StageBranch(stager.formStage, formGroup)

	stager.formStage.Commit()
}
