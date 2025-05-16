// generated code - do not edit
package probe

import (
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/doc2/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

// insertion point
func __gong__New__AttributeShapeFormCallback(
	attributeshape *models.AttributeShape,
	probe *Probe,
	formGroup *table.FormGroup,
) (attributeshapeFormCallback *AttributeShapeFormCallback) {
	attributeshapeFormCallback = new(AttributeShapeFormCallback)
	attributeshapeFormCallback.probe = probe
	attributeshapeFormCallback.attributeshape = attributeshape
	attributeshapeFormCallback.formGroup = formGroup

	attributeshapeFormCallback.CreationMode = (attributeshape == nil)

	return
}

type AttributeShapeFormCallback struct {
	attributeshape *models.AttributeShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (attributeshapeFormCallback *AttributeShapeFormCallback) OnSave() {

	// log.Println("AttributeShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	attributeshapeFormCallback.probe.formStage.Checkout()

	if attributeshapeFormCallback.attributeshape == nil {
		attributeshapeFormCallback.attributeshape = new(models.AttributeShape).Stage(attributeshapeFormCallback.probe.stageOfInterest)
	}
	attributeshape_ := attributeshapeFormCallback.attributeshape
	_ = attributeshape_

	for _, formDiv := range attributeshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(attributeshape_.Name), formDiv)
		case "Identifier":
			FormDivBasicFieldToField(&(attributeshape_.Identifier), formDiv)
		case "FieldTypeAsString":
			FormDivBasicFieldToField(&(attributeshape_.FieldTypeAsString), formDiv)
		case "Structname":
			FormDivBasicFieldToField(&(attributeshape_.Structname), formDiv)
		case "Fieldtypename":
			FormDivBasicFieldToField(&(attributeshape_.Fieldtypename), formDiv)
		case "GongStructShape:AttributeShapes":
			// we need to retrieve the field owner before the change
			var pastGongStructShapeOwner *models.GongStructShape
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongStructShape"
			rf.Fieldname = "AttributeShapes"
			reverseFieldOwner := models.GetReverseFieldOwner(
				attributeshapeFormCallback.probe.stageOfInterest,
				attributeshape_,
				&rf)

			if reverseFieldOwner != nil {
				pastGongStructShapeOwner = reverseFieldOwner.(*models.GongStructShape)
			}
			fieldValue := formDiv.FormFields[0].FormFieldSelect.Value
			if fieldValue == nil {
				if pastGongStructShapeOwner != nil {
					idx := slices.Index(pastGongStructShapeOwner.AttributeShapes, attributeshape_)
					pastGongStructShapeOwner.AttributeShapes = slices.Delete(pastGongStructShapeOwner.AttributeShapes, idx, idx+1)
				}
			} else {

				// if the name of the field value is the same as of the past owner
				// it is assumed the owner has not changed
				// therefore, the owner must be eventualy changed if the name is different
				if pastGongStructShapeOwner.GetName() != fieldValue.GetName() {

					// we need to retrieve the field owner after the change
					// parse all astrcut and get the one with the name in the
					// div
					for _gongstructshape := range *models.GetGongstructInstancesSet[models.GongStructShape](attributeshapeFormCallback.probe.stageOfInterest) {

						// the match is base on the name
						if _gongstructshape.GetName() == fieldValue.GetName() {
							newGongStructShapeOwner := _gongstructshape // we have a match

							// we remove the attributeshape_ instance from the pastGongStructShapeOwner field
							if pastGongStructShapeOwner != nil {
								if newGongStructShapeOwner != pastGongStructShapeOwner {
									idx := slices.Index(pastGongStructShapeOwner.AttributeShapes, attributeshape_)
									pastGongStructShapeOwner.AttributeShapes = slices.Delete(pastGongStructShapeOwner.AttributeShapes, idx, idx+1)
									newGongStructShapeOwner.AttributeShapes = append(newGongStructShapeOwner.AttributeShapes, attributeshape_)
								}
							} else {
								newGongStructShapeOwner.AttributeShapes = append(newGongStructShapeOwner.AttributeShapes, attributeshape_)
							}
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if attributeshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributeshape_.Unstage(attributeshapeFormCallback.probe.stageOfInterest)
	}

	attributeshapeFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.AttributeShape](
		attributeshapeFormCallback.probe,
	)
	attributeshapeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if attributeshapeFormCallback.CreationMode || attributeshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributeshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(attributeshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AttributeShapeFormCallback(
			nil,
			attributeshapeFormCallback.probe,
			newFormGroup,
		)
		attributeshape := new(models.AttributeShape)
		FillUpForm(attributeshape, newFormGroup, attributeshapeFormCallback.probe)
		attributeshapeFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(attributeshapeFormCallback.probe)
}
func __gong__New__ClassdiagramFormCallback(
	classdiagram *models.Classdiagram,
	probe *Probe,
	formGroup *table.FormGroup,
) (classdiagramFormCallback *ClassdiagramFormCallback) {
	classdiagramFormCallback = new(ClassdiagramFormCallback)
	classdiagramFormCallback.probe = probe
	classdiagramFormCallback.classdiagram = classdiagram
	classdiagramFormCallback.formGroup = formGroup

	classdiagramFormCallback.CreationMode = (classdiagram == nil)

	return
}

type ClassdiagramFormCallback struct {
	classdiagram *models.Classdiagram

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (classdiagramFormCallback *ClassdiagramFormCallback) OnSave() {

	// log.Println("ClassdiagramFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	classdiagramFormCallback.probe.formStage.Checkout()

	if classdiagramFormCallback.classdiagram == nil {
		classdiagramFormCallback.classdiagram = new(models.Classdiagram).Stage(classdiagramFormCallback.probe.stageOfInterest)
	}
	classdiagram_ := classdiagramFormCallback.classdiagram
	_ = classdiagram_

	for _, formDiv := range classdiagramFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(classdiagram_.Name), formDiv)
		case "IsInRenameMode":
			FormDivBasicFieldToField(&(classdiagram_.IsInRenameMode), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(classdiagram_.IsExpanded), formDiv)
		case "NodeGongStructsIsExpanded":
			FormDivBasicFieldToField(&(classdiagram_.NodeGongStructsIsExpanded), formDiv)
		case "NodeGongStructNodeExpansionBinaryEncoding":
			FormDivBasicFieldToField(&(classdiagram_.NodeGongStructNodeExpansionBinaryEncoding), formDiv)
		case "NodeGongEnumsIsExpanded":
			FormDivBasicFieldToField(&(classdiagram_.NodeGongEnumsIsExpanded), formDiv)
		case "NodeGongEnumNodeExpansionBinaryEncoding":
			FormDivBasicFieldToField(&(classdiagram_.NodeGongEnumNodeExpansionBinaryEncoding), formDiv)
		case "NodeGongNotesIsExpanded":
			FormDivBasicFieldToField(&(classdiagram_.NodeGongNotesIsExpanded), formDiv)
		case "NodeGongNoteNodeExpansionBinaryEncoding":
			FormDivBasicFieldToField(&(classdiagram_.NodeGongNoteNodeExpansionBinaryEncoding), formDiv)
		case "DiagramPackage:Classdiagrams":
			// we need to retrieve the field owner before the change
			var pastDiagramPackageOwner *models.DiagramPackage
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DiagramPackage"
			rf.Fieldname = "Classdiagrams"
			reverseFieldOwner := models.GetReverseFieldOwner(
				classdiagramFormCallback.probe.stageOfInterest,
				classdiagram_,
				&rf)

			if reverseFieldOwner != nil {
				pastDiagramPackageOwner = reverseFieldOwner.(*models.DiagramPackage)
			}
			fieldValue := formDiv.FormFields[0].FormFieldSelect.Value
			if fieldValue == nil {
				if pastDiagramPackageOwner != nil {
					idx := slices.Index(pastDiagramPackageOwner.Classdiagrams, classdiagram_)
					pastDiagramPackageOwner.Classdiagrams = slices.Delete(pastDiagramPackageOwner.Classdiagrams, idx, idx+1)
				}
			} else {

				// if the name of the field value is the same as of the past owner
				// it is assumed the owner has not changed
				// therefore, the owner must be eventualy changed if the name is different
				if pastDiagramPackageOwner.GetName() != fieldValue.GetName() {

					// we need to retrieve the field owner after the change
					// parse all astrcut and get the one with the name in the
					// div
					for _diagrampackage := range *models.GetGongstructInstancesSet[models.DiagramPackage](classdiagramFormCallback.probe.stageOfInterest) {

						// the match is base on the name
						if _diagrampackage.GetName() == fieldValue.GetName() {
							newDiagramPackageOwner := _diagrampackage // we have a match

							// we remove the classdiagram_ instance from the pastDiagramPackageOwner field
							if pastDiagramPackageOwner != nil {
								if newDiagramPackageOwner != pastDiagramPackageOwner {
									idx := slices.Index(pastDiagramPackageOwner.Classdiagrams, classdiagram_)
									pastDiagramPackageOwner.Classdiagrams = slices.Delete(pastDiagramPackageOwner.Classdiagrams, idx, idx+1)
									newDiagramPackageOwner.Classdiagrams = append(newDiagramPackageOwner.Classdiagrams, classdiagram_)
								}
							} else {
								newDiagramPackageOwner.Classdiagrams = append(newDiagramPackageOwner.Classdiagrams, classdiagram_)
							}
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if classdiagramFormCallback.formGroup.HasSuppressButtonBeenPressed {
		classdiagram_.Unstage(classdiagramFormCallback.probe.stageOfInterest)
	}

	classdiagramFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Classdiagram](
		classdiagramFormCallback.probe,
	)
	classdiagramFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if classdiagramFormCallback.CreationMode || classdiagramFormCallback.formGroup.HasSuppressButtonBeenPressed {
		classdiagramFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(classdiagramFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ClassdiagramFormCallback(
			nil,
			classdiagramFormCallback.probe,
			newFormGroup,
		)
		classdiagram := new(models.Classdiagram)
		FillUpForm(classdiagram, newFormGroup, classdiagramFormCallback.probe)
		classdiagramFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(classdiagramFormCallback.probe)
}
func __gong__New__DiagramPackageFormCallback(
	diagrampackage *models.DiagramPackage,
	probe *Probe,
	formGroup *table.FormGroup,
) (diagrampackageFormCallback *DiagramPackageFormCallback) {
	diagrampackageFormCallback = new(DiagramPackageFormCallback)
	diagrampackageFormCallback.probe = probe
	diagrampackageFormCallback.diagrampackage = diagrampackage
	diagrampackageFormCallback.formGroup = formGroup

	diagrampackageFormCallback.CreationMode = (diagrampackage == nil)

	return
}

type DiagramPackageFormCallback struct {
	diagrampackage *models.DiagramPackage

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (diagrampackageFormCallback *DiagramPackageFormCallback) OnSave() {

	// log.Println("DiagramPackageFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	diagrampackageFormCallback.probe.formStage.Checkout()

	if diagrampackageFormCallback.diagrampackage == nil {
		diagrampackageFormCallback.diagrampackage = new(models.DiagramPackage).Stage(diagrampackageFormCallback.probe.stageOfInterest)
	}
	diagrampackage_ := diagrampackageFormCallback.diagrampackage
	_ = diagrampackage_

	for _, formDiv := range diagrampackageFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(diagrampackage_.Name), formDiv)
		case "Path":
			FormDivBasicFieldToField(&(diagrampackage_.Path), formDiv)
		case "GongModelPath":
			FormDivBasicFieldToField(&(diagrampackage_.GongModelPath), formDiv)
		case "SelectedClassdiagram":
			FormDivSelectFieldToField(&(diagrampackage_.SelectedClassdiagram), diagrampackageFormCallback.probe.stageOfInterest, formDiv)
		case "AbsolutePathToDiagramPackage":
			FormDivBasicFieldToField(&(diagrampackage_.AbsolutePathToDiagramPackage), formDiv)
		}
	}

	// manage the suppress operation
	if diagrampackageFormCallback.formGroup.HasSuppressButtonBeenPressed {
		diagrampackage_.Unstage(diagrampackageFormCallback.probe.stageOfInterest)
	}

	diagrampackageFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.DiagramPackage](
		diagrampackageFormCallback.probe,
	)
	diagrampackageFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if diagrampackageFormCallback.CreationMode || diagrampackageFormCallback.formGroup.HasSuppressButtonBeenPressed {
		diagrampackageFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(diagrampackageFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DiagramPackageFormCallback(
			nil,
			diagrampackageFormCallback.probe,
			newFormGroup,
		)
		diagrampackage := new(models.DiagramPackage)
		FillUpForm(diagrampackage, newFormGroup, diagrampackageFormCallback.probe)
		diagrampackageFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(diagrampackageFormCallback.probe)
}
func __gong__New__GongEnumShapeFormCallback(
	gongenumshape *models.GongEnumShape,
	probe *Probe,
	formGroup *table.FormGroup,
) (gongenumshapeFormCallback *GongEnumShapeFormCallback) {
	gongenumshapeFormCallback = new(GongEnumShapeFormCallback)
	gongenumshapeFormCallback.probe = probe
	gongenumshapeFormCallback.gongenumshape = gongenumshape
	gongenumshapeFormCallback.formGroup = formGroup

	gongenumshapeFormCallback.CreationMode = (gongenumshape == nil)

	return
}

type GongEnumShapeFormCallback struct {
	gongenumshape *models.GongEnumShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (gongenumshapeFormCallback *GongEnumShapeFormCallback) OnSave() {

	// log.Println("GongEnumShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongenumshapeFormCallback.probe.formStage.Checkout()

	if gongenumshapeFormCallback.gongenumshape == nil {
		gongenumshapeFormCallback.gongenumshape = new(models.GongEnumShape).Stage(gongenumshapeFormCallback.probe.stageOfInterest)
	}
	gongenumshape_ := gongenumshapeFormCallback.gongenumshape
	_ = gongenumshape_

	for _, formDiv := range gongenumshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongenumshape_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(gongenumshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(gongenumshape_.Y), formDiv)
		case "Identifier":
			FormDivBasicFieldToField(&(gongenumshape_.Identifier), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(gongenumshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(gongenumshape_.Height), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(gongenumshape_.IsExpanded), formDiv)
		case "Classdiagram:GongEnumShapes":
			// we need to retrieve the field owner before the change
			var pastClassdiagramOwner *models.Classdiagram
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Classdiagram"
			rf.Fieldname = "GongEnumShapes"
			reverseFieldOwner := models.GetReverseFieldOwner(
				gongenumshapeFormCallback.probe.stageOfInterest,
				gongenumshape_,
				&rf)

			if reverseFieldOwner != nil {
				pastClassdiagramOwner = reverseFieldOwner.(*models.Classdiagram)
			}
			fieldValue := formDiv.FormFields[0].FormFieldSelect.Value
			if fieldValue == nil {
				if pastClassdiagramOwner != nil {
					idx := slices.Index(pastClassdiagramOwner.GongEnumShapes, gongenumshape_)
					pastClassdiagramOwner.GongEnumShapes = slices.Delete(pastClassdiagramOwner.GongEnumShapes, idx, idx+1)
				}
			} else {

				// if the name of the field value is the same as of the past owner
				// it is assumed the owner has not changed
				// therefore, the owner must be eventualy changed if the name is different
				if pastClassdiagramOwner.GetName() != fieldValue.GetName() {

					// we need to retrieve the field owner after the change
					// parse all astrcut and get the one with the name in the
					// div
					for _classdiagram := range *models.GetGongstructInstancesSet[models.Classdiagram](gongenumshapeFormCallback.probe.stageOfInterest) {

						// the match is base on the name
						if _classdiagram.GetName() == fieldValue.GetName() {
							newClassdiagramOwner := _classdiagram // we have a match

							// we remove the gongenumshape_ instance from the pastClassdiagramOwner field
							if pastClassdiagramOwner != nil {
								if newClassdiagramOwner != pastClassdiagramOwner {
									idx := slices.Index(pastClassdiagramOwner.GongEnumShapes, gongenumshape_)
									pastClassdiagramOwner.GongEnumShapes = slices.Delete(pastClassdiagramOwner.GongEnumShapes, idx, idx+1)
									newClassdiagramOwner.GongEnumShapes = append(newClassdiagramOwner.GongEnumShapes, gongenumshape_)
								}
							} else {
								newClassdiagramOwner.GongEnumShapes = append(newClassdiagramOwner.GongEnumShapes, gongenumshape_)
							}
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if gongenumshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongenumshape_.Unstage(gongenumshapeFormCallback.probe.stageOfInterest)
	}

	gongenumshapeFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.GongEnumShape](
		gongenumshapeFormCallback.probe,
	)
	gongenumshapeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if gongenumshapeFormCallback.CreationMode || gongenumshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongenumshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(gongenumshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GongEnumShapeFormCallback(
			nil,
			gongenumshapeFormCallback.probe,
			newFormGroup,
		)
		gongenumshape := new(models.GongEnumShape)
		FillUpForm(gongenumshape, newFormGroup, gongenumshapeFormCallback.probe)
		gongenumshapeFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(gongenumshapeFormCallback.probe)
}
func __gong__New__GongEnumValueShapeFormCallback(
	gongenumvalueshape *models.GongEnumValueShape,
	probe *Probe,
	formGroup *table.FormGroup,
) (gongenumvalueshapeFormCallback *GongEnumValueShapeFormCallback) {
	gongenumvalueshapeFormCallback = new(GongEnumValueShapeFormCallback)
	gongenumvalueshapeFormCallback.probe = probe
	gongenumvalueshapeFormCallback.gongenumvalueshape = gongenumvalueshape
	gongenumvalueshapeFormCallback.formGroup = formGroup

	gongenumvalueshapeFormCallback.CreationMode = (gongenumvalueshape == nil)

	return
}

type GongEnumValueShapeFormCallback struct {
	gongenumvalueshape *models.GongEnumValueShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (gongenumvalueshapeFormCallback *GongEnumValueShapeFormCallback) OnSave() {

	// log.Println("GongEnumValueShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongenumvalueshapeFormCallback.probe.formStage.Checkout()

	if gongenumvalueshapeFormCallback.gongenumvalueshape == nil {
		gongenumvalueshapeFormCallback.gongenumvalueshape = new(models.GongEnumValueShape).Stage(gongenumvalueshapeFormCallback.probe.stageOfInterest)
	}
	gongenumvalueshape_ := gongenumvalueshapeFormCallback.gongenumvalueshape
	_ = gongenumvalueshape_

	for _, formDiv := range gongenumvalueshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongenumvalueshape_.Name), formDiv)
		case "Identifier":
			FormDivBasicFieldToField(&(gongenumvalueshape_.Identifier), formDiv)
		case "GongEnumShape:GongEnumValueShapes":
			// we need to retrieve the field owner before the change
			var pastGongEnumShapeOwner *models.GongEnumShape
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongEnumShape"
			rf.Fieldname = "GongEnumValueShapes"
			reverseFieldOwner := models.GetReverseFieldOwner(
				gongenumvalueshapeFormCallback.probe.stageOfInterest,
				gongenumvalueshape_,
				&rf)

			if reverseFieldOwner != nil {
				pastGongEnumShapeOwner = reverseFieldOwner.(*models.GongEnumShape)
			}
			fieldValue := formDiv.FormFields[0].FormFieldSelect.Value
			if fieldValue == nil {
				if pastGongEnumShapeOwner != nil {
					idx := slices.Index(pastGongEnumShapeOwner.GongEnumValueShapes, gongenumvalueshape_)
					pastGongEnumShapeOwner.GongEnumValueShapes = slices.Delete(pastGongEnumShapeOwner.GongEnumValueShapes, idx, idx+1)
				}
			} else {

				// if the name of the field value is the same as of the past owner
				// it is assumed the owner has not changed
				// therefore, the owner must be eventualy changed if the name is different
				if pastGongEnumShapeOwner.GetName() != fieldValue.GetName() {

					// we need to retrieve the field owner after the change
					// parse all astrcut and get the one with the name in the
					// div
					for _gongenumshape := range *models.GetGongstructInstancesSet[models.GongEnumShape](gongenumvalueshapeFormCallback.probe.stageOfInterest) {

						// the match is base on the name
						if _gongenumshape.GetName() == fieldValue.GetName() {
							newGongEnumShapeOwner := _gongenumshape // we have a match

							// we remove the gongenumvalueshape_ instance from the pastGongEnumShapeOwner field
							if pastGongEnumShapeOwner != nil {
								if newGongEnumShapeOwner != pastGongEnumShapeOwner {
									idx := slices.Index(pastGongEnumShapeOwner.GongEnumValueShapes, gongenumvalueshape_)
									pastGongEnumShapeOwner.GongEnumValueShapes = slices.Delete(pastGongEnumShapeOwner.GongEnumValueShapes, idx, idx+1)
									newGongEnumShapeOwner.GongEnumValueShapes = append(newGongEnumShapeOwner.GongEnumValueShapes, gongenumvalueshape_)
								}
							} else {
								newGongEnumShapeOwner.GongEnumValueShapes = append(newGongEnumShapeOwner.GongEnumValueShapes, gongenumvalueshape_)
							}
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if gongenumvalueshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongenumvalueshape_.Unstage(gongenumvalueshapeFormCallback.probe.stageOfInterest)
	}

	gongenumvalueshapeFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.GongEnumValueShape](
		gongenumvalueshapeFormCallback.probe,
	)
	gongenumvalueshapeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if gongenumvalueshapeFormCallback.CreationMode || gongenumvalueshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongenumvalueshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(gongenumvalueshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GongEnumValueShapeFormCallback(
			nil,
			gongenumvalueshapeFormCallback.probe,
			newFormGroup,
		)
		gongenumvalueshape := new(models.GongEnumValueShape)
		FillUpForm(gongenumvalueshape, newFormGroup, gongenumvalueshapeFormCallback.probe)
		gongenumvalueshapeFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(gongenumvalueshapeFormCallback.probe)
}
func __gong__New__GongNoteLinkShapeFormCallback(
	gongnotelinkshape *models.GongNoteLinkShape,
	probe *Probe,
	formGroup *table.FormGroup,
) (gongnotelinkshapeFormCallback *GongNoteLinkShapeFormCallback) {
	gongnotelinkshapeFormCallback = new(GongNoteLinkShapeFormCallback)
	gongnotelinkshapeFormCallback.probe = probe
	gongnotelinkshapeFormCallback.gongnotelinkshape = gongnotelinkshape
	gongnotelinkshapeFormCallback.formGroup = formGroup

	gongnotelinkshapeFormCallback.CreationMode = (gongnotelinkshape == nil)

	return
}

type GongNoteLinkShapeFormCallback struct {
	gongnotelinkshape *models.GongNoteLinkShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (gongnotelinkshapeFormCallback *GongNoteLinkShapeFormCallback) OnSave() {

	// log.Println("GongNoteLinkShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongnotelinkshapeFormCallback.probe.formStage.Checkout()

	if gongnotelinkshapeFormCallback.gongnotelinkshape == nil {
		gongnotelinkshapeFormCallback.gongnotelinkshape = new(models.GongNoteLinkShape).Stage(gongnotelinkshapeFormCallback.probe.stageOfInterest)
	}
	gongnotelinkshape_ := gongnotelinkshapeFormCallback.gongnotelinkshape
	_ = gongnotelinkshape_

	for _, formDiv := range gongnotelinkshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongnotelinkshape_.Name), formDiv)
		case "Identifier":
			FormDivBasicFieldToField(&(gongnotelinkshape_.Identifier), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(gongnotelinkshape_.Type), formDiv)
		case "GongNoteShape:GongNoteLinkShapes":
			// we need to retrieve the field owner before the change
			var pastGongNoteShapeOwner *models.GongNoteShape
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongNoteShape"
			rf.Fieldname = "GongNoteLinkShapes"
			reverseFieldOwner := models.GetReverseFieldOwner(
				gongnotelinkshapeFormCallback.probe.stageOfInterest,
				gongnotelinkshape_,
				&rf)

			if reverseFieldOwner != nil {
				pastGongNoteShapeOwner = reverseFieldOwner.(*models.GongNoteShape)
			}
			fieldValue := formDiv.FormFields[0].FormFieldSelect.Value
			if fieldValue == nil {
				if pastGongNoteShapeOwner != nil {
					idx := slices.Index(pastGongNoteShapeOwner.GongNoteLinkShapes, gongnotelinkshape_)
					pastGongNoteShapeOwner.GongNoteLinkShapes = slices.Delete(pastGongNoteShapeOwner.GongNoteLinkShapes, idx, idx+1)
				}
			} else {

				// if the name of the field value is the same as of the past owner
				// it is assumed the owner has not changed
				// therefore, the owner must be eventualy changed if the name is different
				if pastGongNoteShapeOwner.GetName() != fieldValue.GetName() {

					// we need to retrieve the field owner after the change
					// parse all astrcut and get the one with the name in the
					// div
					for _gongnoteshape := range *models.GetGongstructInstancesSet[models.GongNoteShape](gongnotelinkshapeFormCallback.probe.stageOfInterest) {

						// the match is base on the name
						if _gongnoteshape.GetName() == fieldValue.GetName() {
							newGongNoteShapeOwner := _gongnoteshape // we have a match

							// we remove the gongnotelinkshape_ instance from the pastGongNoteShapeOwner field
							if pastGongNoteShapeOwner != nil {
								if newGongNoteShapeOwner != pastGongNoteShapeOwner {
									idx := slices.Index(pastGongNoteShapeOwner.GongNoteLinkShapes, gongnotelinkshape_)
									pastGongNoteShapeOwner.GongNoteLinkShapes = slices.Delete(pastGongNoteShapeOwner.GongNoteLinkShapes, idx, idx+1)
									newGongNoteShapeOwner.GongNoteLinkShapes = append(newGongNoteShapeOwner.GongNoteLinkShapes, gongnotelinkshape_)
								}
							} else {
								newGongNoteShapeOwner.GongNoteLinkShapes = append(newGongNoteShapeOwner.GongNoteLinkShapes, gongnotelinkshape_)
							}
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if gongnotelinkshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongnotelinkshape_.Unstage(gongnotelinkshapeFormCallback.probe.stageOfInterest)
	}

	gongnotelinkshapeFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.GongNoteLinkShape](
		gongnotelinkshapeFormCallback.probe,
	)
	gongnotelinkshapeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if gongnotelinkshapeFormCallback.CreationMode || gongnotelinkshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongnotelinkshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(gongnotelinkshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GongNoteLinkShapeFormCallback(
			nil,
			gongnotelinkshapeFormCallback.probe,
			newFormGroup,
		)
		gongnotelinkshape := new(models.GongNoteLinkShape)
		FillUpForm(gongnotelinkshape, newFormGroup, gongnotelinkshapeFormCallback.probe)
		gongnotelinkshapeFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(gongnotelinkshapeFormCallback.probe)
}
func __gong__New__GongNoteShapeFormCallback(
	gongnoteshape *models.GongNoteShape,
	probe *Probe,
	formGroup *table.FormGroup,
) (gongnoteshapeFormCallback *GongNoteShapeFormCallback) {
	gongnoteshapeFormCallback = new(GongNoteShapeFormCallback)
	gongnoteshapeFormCallback.probe = probe
	gongnoteshapeFormCallback.gongnoteshape = gongnoteshape
	gongnoteshapeFormCallback.formGroup = formGroup

	gongnoteshapeFormCallback.CreationMode = (gongnoteshape == nil)

	return
}

type GongNoteShapeFormCallback struct {
	gongnoteshape *models.GongNoteShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (gongnoteshapeFormCallback *GongNoteShapeFormCallback) OnSave() {

	// log.Println("GongNoteShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongnoteshapeFormCallback.probe.formStage.Checkout()

	if gongnoteshapeFormCallback.gongnoteshape == nil {
		gongnoteshapeFormCallback.gongnoteshape = new(models.GongNoteShape).Stage(gongnoteshapeFormCallback.probe.stageOfInterest)
	}
	gongnoteshape_ := gongnoteshapeFormCallback.gongnoteshape
	_ = gongnoteshape_

	for _, formDiv := range gongnoteshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongnoteshape_.Name), formDiv)
		case "Identifier":
			FormDivBasicFieldToField(&(gongnoteshape_.Identifier), formDiv)
		case "Body":
			FormDivBasicFieldToField(&(gongnoteshape_.Body), formDiv)
		case "BodyHTML":
			FormDivBasicFieldToField(&(gongnoteshape_.BodyHTML), formDiv)
		case "X":
			FormDivBasicFieldToField(&(gongnoteshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(gongnoteshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(gongnoteshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(gongnoteshape_.Height), formDiv)
		case "Matched":
			FormDivBasicFieldToField(&(gongnoteshape_.Matched), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(gongnoteshape_.IsExpanded), formDiv)
		case "Classdiagram:GongNoteShapes":
			// we need to retrieve the field owner before the change
			var pastClassdiagramOwner *models.Classdiagram
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Classdiagram"
			rf.Fieldname = "GongNoteShapes"
			reverseFieldOwner := models.GetReverseFieldOwner(
				gongnoteshapeFormCallback.probe.stageOfInterest,
				gongnoteshape_,
				&rf)

			if reverseFieldOwner != nil {
				pastClassdiagramOwner = reverseFieldOwner.(*models.Classdiagram)
			}
			fieldValue := formDiv.FormFields[0].FormFieldSelect.Value
			if fieldValue == nil {
				if pastClassdiagramOwner != nil {
					idx := slices.Index(pastClassdiagramOwner.GongNoteShapes, gongnoteshape_)
					pastClassdiagramOwner.GongNoteShapes = slices.Delete(pastClassdiagramOwner.GongNoteShapes, idx, idx+1)
				}
			} else {

				// if the name of the field value is the same as of the past owner
				// it is assumed the owner has not changed
				// therefore, the owner must be eventualy changed if the name is different
				if pastClassdiagramOwner.GetName() != fieldValue.GetName() {

					// we need to retrieve the field owner after the change
					// parse all astrcut and get the one with the name in the
					// div
					for _classdiagram := range *models.GetGongstructInstancesSet[models.Classdiagram](gongnoteshapeFormCallback.probe.stageOfInterest) {

						// the match is base on the name
						if _classdiagram.GetName() == fieldValue.GetName() {
							newClassdiagramOwner := _classdiagram // we have a match

							// we remove the gongnoteshape_ instance from the pastClassdiagramOwner field
							if pastClassdiagramOwner != nil {
								if newClassdiagramOwner != pastClassdiagramOwner {
									idx := slices.Index(pastClassdiagramOwner.GongNoteShapes, gongnoteshape_)
									pastClassdiagramOwner.GongNoteShapes = slices.Delete(pastClassdiagramOwner.GongNoteShapes, idx, idx+1)
									newClassdiagramOwner.GongNoteShapes = append(newClassdiagramOwner.GongNoteShapes, gongnoteshape_)
								}
							} else {
								newClassdiagramOwner.GongNoteShapes = append(newClassdiagramOwner.GongNoteShapes, gongnoteshape_)
							}
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if gongnoteshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongnoteshape_.Unstage(gongnoteshapeFormCallback.probe.stageOfInterest)
	}

	gongnoteshapeFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.GongNoteShape](
		gongnoteshapeFormCallback.probe,
	)
	gongnoteshapeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if gongnoteshapeFormCallback.CreationMode || gongnoteshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongnoteshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(gongnoteshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GongNoteShapeFormCallback(
			nil,
			gongnoteshapeFormCallback.probe,
			newFormGroup,
		)
		gongnoteshape := new(models.GongNoteShape)
		FillUpForm(gongnoteshape, newFormGroup, gongnoteshapeFormCallback.probe)
		gongnoteshapeFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(gongnoteshapeFormCallback.probe)
}
func __gong__New__GongStructShapeFormCallback(
	gongstructshape *models.GongStructShape,
	probe *Probe,
	formGroup *table.FormGroup,
) (gongstructshapeFormCallback *GongStructShapeFormCallback) {
	gongstructshapeFormCallback = new(GongStructShapeFormCallback)
	gongstructshapeFormCallback.probe = probe
	gongstructshapeFormCallback.gongstructshape = gongstructshape
	gongstructshapeFormCallback.formGroup = formGroup

	gongstructshapeFormCallback.CreationMode = (gongstructshape == nil)

	return
}

type GongStructShapeFormCallback struct {
	gongstructshape *models.GongStructShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (gongstructshapeFormCallback *GongStructShapeFormCallback) OnSave() {

	// log.Println("GongStructShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongstructshapeFormCallback.probe.formStage.Checkout()

	if gongstructshapeFormCallback.gongstructshape == nil {
		gongstructshapeFormCallback.gongstructshape = new(models.GongStructShape).Stage(gongstructshapeFormCallback.probe.stageOfInterest)
	}
	gongstructshape_ := gongstructshapeFormCallback.gongstructshape
	_ = gongstructshape_

	for _, formDiv := range gongstructshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongstructshape_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(gongstructshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(gongstructshape_.Y), formDiv)
		case "Identifier":
			FormDivBasicFieldToField(&(gongstructshape_.Identifier), formDiv)
		case "ShowNbInstances":
			FormDivBasicFieldToField(&(gongstructshape_.ShowNbInstances), formDiv)
		case "NbInstances":
			FormDivBasicFieldToField(&(gongstructshape_.NbInstances), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(gongstructshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(gongstructshape_.Height), formDiv)
		case "IsSelected":
			FormDivBasicFieldToField(&(gongstructshape_.IsSelected), formDiv)
		case "Classdiagram:GongStructShapes":
			// we need to retrieve the field owner before the change
			var pastClassdiagramOwner *models.Classdiagram
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Classdiagram"
			rf.Fieldname = "GongStructShapes"
			reverseFieldOwner := models.GetReverseFieldOwner(
				gongstructshapeFormCallback.probe.stageOfInterest,
				gongstructshape_,
				&rf)

			if reverseFieldOwner != nil {
				pastClassdiagramOwner = reverseFieldOwner.(*models.Classdiagram)
			}
			fieldValue := formDiv.FormFields[0].FormFieldSelect.Value
			if fieldValue == nil {
				if pastClassdiagramOwner != nil {
					idx := slices.Index(pastClassdiagramOwner.GongStructShapes, gongstructshape_)
					pastClassdiagramOwner.GongStructShapes = slices.Delete(pastClassdiagramOwner.GongStructShapes, idx, idx+1)
				}
			} else {

				// if the name of the field value is the same as of the past owner
				// it is assumed the owner has not changed
				// therefore, the owner must be eventualy changed if the name is different
				if pastClassdiagramOwner.GetName() != fieldValue.GetName() {

					// we need to retrieve the field owner after the change
					// parse all astrcut and get the one with the name in the
					// div
					for _classdiagram := range *models.GetGongstructInstancesSet[models.Classdiagram](gongstructshapeFormCallback.probe.stageOfInterest) {

						// the match is base on the name
						if _classdiagram.GetName() == fieldValue.GetName() {
							newClassdiagramOwner := _classdiagram // we have a match

							// we remove the gongstructshape_ instance from the pastClassdiagramOwner field
							if pastClassdiagramOwner != nil {
								if newClassdiagramOwner != pastClassdiagramOwner {
									idx := slices.Index(pastClassdiagramOwner.GongStructShapes, gongstructshape_)
									pastClassdiagramOwner.GongStructShapes = slices.Delete(pastClassdiagramOwner.GongStructShapes, idx, idx+1)
									newClassdiagramOwner.GongStructShapes = append(newClassdiagramOwner.GongStructShapes, gongstructshape_)
								}
							} else {
								newClassdiagramOwner.GongStructShapes = append(newClassdiagramOwner.GongStructShapes, gongstructshape_)
							}
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if gongstructshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongstructshape_.Unstage(gongstructshapeFormCallback.probe.stageOfInterest)
	}

	gongstructshapeFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.GongStructShape](
		gongstructshapeFormCallback.probe,
	)
	gongstructshapeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if gongstructshapeFormCallback.CreationMode || gongstructshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongstructshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(gongstructshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GongStructShapeFormCallback(
			nil,
			gongstructshapeFormCallback.probe,
			newFormGroup,
		)
		gongstructshape := new(models.GongStructShape)
		FillUpForm(gongstructshape, newFormGroup, gongstructshapeFormCallback.probe)
		gongstructshapeFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(gongstructshapeFormCallback.probe)
}
func __gong__New__LinkShapeFormCallback(
	linkshape *models.LinkShape,
	probe *Probe,
	formGroup *table.FormGroup,
) (linkshapeFormCallback *LinkShapeFormCallback) {
	linkshapeFormCallback = new(LinkShapeFormCallback)
	linkshapeFormCallback.probe = probe
	linkshapeFormCallback.linkshape = linkshape
	linkshapeFormCallback.formGroup = formGroup

	linkshapeFormCallback.CreationMode = (linkshape == nil)

	return
}

type LinkShapeFormCallback struct {
	linkshape *models.LinkShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (linkshapeFormCallback *LinkShapeFormCallback) OnSave() {

	// log.Println("LinkShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	linkshapeFormCallback.probe.formStage.Checkout()

	if linkshapeFormCallback.linkshape == nil {
		linkshapeFormCallback.linkshape = new(models.LinkShape).Stage(linkshapeFormCallback.probe.stageOfInterest)
	}
	linkshape_ := linkshapeFormCallback.linkshape
	_ = linkshape_

	for _, formDiv := range linkshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(linkshape_.Name), formDiv)
		case "Identifier":
			FormDivBasicFieldToField(&(linkshape_.Identifier), formDiv)
		case "Fieldtypename":
			FormDivBasicFieldToField(&(linkshape_.Fieldtypename), formDiv)
		case "FieldOffsetX":
			FormDivBasicFieldToField(&(linkshape_.FieldOffsetX), formDiv)
		case "FieldOffsetY":
			FormDivBasicFieldToField(&(linkshape_.FieldOffsetY), formDiv)
		case "TargetMultiplicity":
			FormDivEnumStringFieldToField(&(linkshape_.TargetMultiplicity), formDiv)
		case "TargetMultiplicityOffsetX":
			FormDivBasicFieldToField(&(linkshape_.TargetMultiplicityOffsetX), formDiv)
		case "TargetMultiplicityOffsetY":
			FormDivBasicFieldToField(&(linkshape_.TargetMultiplicityOffsetY), formDiv)
		case "SourceMultiplicity":
			FormDivEnumStringFieldToField(&(linkshape_.SourceMultiplicity), formDiv)
		case "SourceMultiplicityOffsetX":
			FormDivBasicFieldToField(&(linkshape_.SourceMultiplicityOffsetX), formDiv)
		case "SourceMultiplicityOffsetY":
			FormDivBasicFieldToField(&(linkshape_.SourceMultiplicityOffsetY), formDiv)
		case "X":
			FormDivBasicFieldToField(&(linkshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(linkshape_.Y), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(linkshape_.StartOrientation), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(linkshape_.StartRatio), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(linkshape_.EndOrientation), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(linkshape_.EndRatio), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(linkshape_.CornerOffsetRatio), formDiv)
		case "GongStructShape:LinkShapes":
			// we need to retrieve the field owner before the change
			var pastGongStructShapeOwner *models.GongStructShape
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongStructShape"
			rf.Fieldname = "LinkShapes"
			reverseFieldOwner := models.GetReverseFieldOwner(
				linkshapeFormCallback.probe.stageOfInterest,
				linkshape_,
				&rf)

			if reverseFieldOwner != nil {
				pastGongStructShapeOwner = reverseFieldOwner.(*models.GongStructShape)
			}
			fieldValue := formDiv.FormFields[0].FormFieldSelect.Value
			if fieldValue == nil {
				if pastGongStructShapeOwner != nil {
					idx := slices.Index(pastGongStructShapeOwner.LinkShapes, linkshape_)
					pastGongStructShapeOwner.LinkShapes = slices.Delete(pastGongStructShapeOwner.LinkShapes, idx, idx+1)
				}
			} else {

				// if the name of the field value is the same as of the past owner
				// it is assumed the owner has not changed
				// therefore, the owner must be eventualy changed if the name is different
				if pastGongStructShapeOwner.GetName() != fieldValue.GetName() {

					// we need to retrieve the field owner after the change
					// parse all astrcut and get the one with the name in the
					// div
					for _gongstructshape := range *models.GetGongstructInstancesSet[models.GongStructShape](linkshapeFormCallback.probe.stageOfInterest) {

						// the match is base on the name
						if _gongstructshape.GetName() == fieldValue.GetName() {
							newGongStructShapeOwner := _gongstructshape // we have a match

							// we remove the linkshape_ instance from the pastGongStructShapeOwner field
							if pastGongStructShapeOwner != nil {
								if newGongStructShapeOwner != pastGongStructShapeOwner {
									idx := slices.Index(pastGongStructShapeOwner.LinkShapes, linkshape_)
									pastGongStructShapeOwner.LinkShapes = slices.Delete(pastGongStructShapeOwner.LinkShapes, idx, idx+1)
									newGongStructShapeOwner.LinkShapes = append(newGongStructShapeOwner.LinkShapes, linkshape_)
								}
							} else {
								newGongStructShapeOwner.LinkShapes = append(newGongStructShapeOwner.LinkShapes, linkshape_)
							}
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if linkshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		linkshape_.Unstage(linkshapeFormCallback.probe.stageOfInterest)
	}

	linkshapeFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.LinkShape](
		linkshapeFormCallback.probe,
	)
	linkshapeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if linkshapeFormCallback.CreationMode || linkshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		linkshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(linkshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LinkShapeFormCallback(
			nil,
			linkshapeFormCallback.probe,
			newFormGroup,
		)
		linkshape := new(models.LinkShape)
		FillUpForm(linkshape, newFormGroup, linkshapeFormCallback.probe)
		linkshapeFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(linkshapeFormCallback.probe)
}
