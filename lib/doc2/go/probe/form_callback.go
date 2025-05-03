// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/doc2/go/models"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)

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

	log.Println("AttributeShapeFormCallback, OnSave")

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
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGongStructShapeOwner != nil {
					idx := slices.Index(pastGongStructShapeOwner.AttributeShapes, attributeshape_)
					pastGongStructShapeOwner.AttributeShapes = slices.Delete(pastGongStructShapeOwner.AttributeShapes, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _gongstructshape := range *models.GetGongstructInstancesSet[models.GongStructShape](attributeshapeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _gongstructshape.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGongStructShapeOwner := _gongstructshape // we have a match
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

	// manage the suppress operation
	if attributeshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		attributeshape_.Unstage(attributeshapeFormCallback.probe.stageOfInterest)
	}

	attributeshapeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.AttributeShape](
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

	fillUpTree(attributeshapeFormCallback.probe)
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

	log.Println("ClassdiagramFormCallback, OnSave")

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
		case "IsInDrawMode":
			FormDivBasicFieldToField(&(classdiagram_.IsInDrawMode), formDiv)
		case "IsInRenameMode":
			FormDivBasicFieldToField(&(classdiagram_.IsInRenameMode), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(classdiagram_.IsExpanded), formDiv)
		case "NodeGongStructsIsExpanded":
			FormDivBasicFieldToField(&(classdiagram_.NodeGongStructsIsExpanded), formDiv)
		case "NodeGongStructsBinaryEncoding":
			FormDivBasicFieldToField(&(classdiagram_.NodeGongStructsBinaryEncoding), formDiv)
		case "NodeGongEnumsIsExpanded":
			FormDivBasicFieldToField(&(classdiagram_.NodeGongEnumsIsExpanded), formDiv)
		case "NodeGongNotesIsExpanded":
			FormDivBasicFieldToField(&(classdiagram_.NodeGongNotesIsExpanded), formDiv)
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
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastDiagramPackageOwner != nil {
					idx := slices.Index(pastDiagramPackageOwner.Classdiagrams, classdiagram_)
					pastDiagramPackageOwner.Classdiagrams = slices.Delete(pastDiagramPackageOwner.Classdiagrams, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _diagrampackage := range *models.GetGongstructInstancesSet[models.DiagramPackage](classdiagramFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _diagrampackage.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newDiagramPackageOwner := _diagrampackage // we have a match
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

	// manage the suppress operation
	if classdiagramFormCallback.formGroup.HasSuppressButtonBeenPressed {
		classdiagram_.Unstage(classdiagramFormCallback.probe.stageOfInterest)
	}

	classdiagramFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Classdiagram](
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

	fillUpTree(classdiagramFormCallback.probe)
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

	log.Println("DiagramPackageFormCallback, OnSave")

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
		case "IsEditable":
			FormDivBasicFieldToField(&(diagrampackage_.IsEditable), formDiv)
		case "IsReloaded":
			FormDivBasicFieldToField(&(diagrampackage_.IsReloaded), formDiv)
		case "AbsolutePathToDiagramPackage":
			FormDivBasicFieldToField(&(diagrampackage_.AbsolutePathToDiagramPackage), formDiv)
		}
	}

	// manage the suppress operation
	if diagrampackageFormCallback.formGroup.HasSuppressButtonBeenPressed {
		diagrampackage_.Unstage(diagrampackageFormCallback.probe.stageOfInterest)
	}

	diagrampackageFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.DiagramPackage](
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

	fillUpTree(diagrampackageFormCallback.probe)
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

	log.Println("GongEnumShapeFormCallback, OnSave")

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
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastClassdiagramOwner != nil {
					idx := slices.Index(pastClassdiagramOwner.GongEnumShapes, gongenumshape_)
					pastClassdiagramOwner.GongEnumShapes = slices.Delete(pastClassdiagramOwner.GongEnumShapes, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _classdiagram := range *models.GetGongstructInstancesSet[models.Classdiagram](gongenumshapeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _classdiagram.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newClassdiagramOwner := _classdiagram // we have a match
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

	// manage the suppress operation
	if gongenumshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongenumshape_.Unstage(gongenumshapeFormCallback.probe.stageOfInterest)
	}

	gongenumshapeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.GongEnumShape](
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

	fillUpTree(gongenumshapeFormCallback.probe)
}
func __gong__New__GongEnumValueEntryFormCallback(
	gongenumvalueentry *models.GongEnumValueEntry,
	probe *Probe,
	formGroup *table.FormGroup,
) (gongenumvalueentryFormCallback *GongEnumValueEntryFormCallback) {
	gongenumvalueentryFormCallback = new(GongEnumValueEntryFormCallback)
	gongenumvalueentryFormCallback.probe = probe
	gongenumvalueentryFormCallback.gongenumvalueentry = gongenumvalueentry
	gongenumvalueentryFormCallback.formGroup = formGroup

	gongenumvalueentryFormCallback.CreationMode = (gongenumvalueentry == nil)

	return
}

type GongEnumValueEntryFormCallback struct {
	gongenumvalueentry *models.GongEnumValueEntry

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (gongenumvalueentryFormCallback *GongEnumValueEntryFormCallback) OnSave() {

	log.Println("GongEnumValueEntryFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gongenumvalueentryFormCallback.probe.formStage.Checkout()

	if gongenumvalueentryFormCallback.gongenumvalueentry == nil {
		gongenumvalueentryFormCallback.gongenumvalueentry = new(models.GongEnumValueEntry).Stage(gongenumvalueentryFormCallback.probe.stageOfInterest)
	}
	gongenumvalueentry_ := gongenumvalueentryFormCallback.gongenumvalueentry
	_ = gongenumvalueentry_

	for _, formDiv := range gongenumvalueentryFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gongenumvalueentry_.Name), formDiv)
		case "Identifier":
			FormDivBasicFieldToField(&(gongenumvalueentry_.Identifier), formDiv)
		case "GongEnumShape:GongEnumValueEntrys":
			// we need to retrieve the field owner before the change
			var pastGongEnumShapeOwner *models.GongEnumShape
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongEnumShape"
			rf.Fieldname = "GongEnumValueEntrys"
			reverseFieldOwner := models.GetReverseFieldOwner(
				gongenumvalueentryFormCallback.probe.stageOfInterest,
				gongenumvalueentry_,
				&rf)

			if reverseFieldOwner != nil {
				pastGongEnumShapeOwner = reverseFieldOwner.(*models.GongEnumShape)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGongEnumShapeOwner != nil {
					idx := slices.Index(pastGongEnumShapeOwner.GongEnumValueEntrys, gongenumvalueentry_)
					pastGongEnumShapeOwner.GongEnumValueEntrys = slices.Delete(pastGongEnumShapeOwner.GongEnumValueEntrys, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _gongenumshape := range *models.GetGongstructInstancesSet[models.GongEnumShape](gongenumvalueentryFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _gongenumshape.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGongEnumShapeOwner := _gongenumshape // we have a match
						if pastGongEnumShapeOwner != nil {
							if newGongEnumShapeOwner != pastGongEnumShapeOwner {
								idx := slices.Index(pastGongEnumShapeOwner.GongEnumValueEntrys, gongenumvalueentry_)
								pastGongEnumShapeOwner.GongEnumValueEntrys = slices.Delete(pastGongEnumShapeOwner.GongEnumValueEntrys, idx, idx+1)
								newGongEnumShapeOwner.GongEnumValueEntrys = append(newGongEnumShapeOwner.GongEnumValueEntrys, gongenumvalueentry_)
							}
						} else {
							newGongEnumShapeOwner.GongEnumValueEntrys = append(newGongEnumShapeOwner.GongEnumValueEntrys, gongenumvalueentry_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if gongenumvalueentryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongenumvalueentry_.Unstage(gongenumvalueentryFormCallback.probe.stageOfInterest)
	}

	gongenumvalueentryFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.GongEnumValueEntry](
		gongenumvalueentryFormCallback.probe,
	)
	gongenumvalueentryFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if gongenumvalueentryFormCallback.CreationMode || gongenumvalueentryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongenumvalueentryFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(gongenumvalueentryFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GongEnumValueEntryFormCallback(
			nil,
			gongenumvalueentryFormCallback.probe,
			newFormGroup,
		)
		gongenumvalueentry := new(models.GongEnumValueEntry)
		FillUpForm(gongenumvalueentry, newFormGroup, gongenumvalueentryFormCallback.probe)
		gongenumvalueentryFormCallback.probe.formStage.Commit()
	}

	fillUpTree(gongenumvalueentryFormCallback.probe)
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

	log.Println("GongStructShapeFormCallback, OnSave")

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
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastClassdiagramOwner != nil {
					idx := slices.Index(pastClassdiagramOwner.GongStructShapes, gongstructshape_)
					pastClassdiagramOwner.GongStructShapes = slices.Delete(pastClassdiagramOwner.GongStructShapes, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _classdiagram := range *models.GetGongstructInstancesSet[models.Classdiagram](gongstructshapeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _classdiagram.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newClassdiagramOwner := _classdiagram // we have a match
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

	// manage the suppress operation
	if gongstructshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongstructshape_.Unstage(gongstructshapeFormCallback.probe.stageOfInterest)
	}

	gongstructshapeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.GongStructShape](
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

	fillUpTree(gongstructshapeFormCallback.probe)
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

	log.Println("LinkShapeFormCallback, OnSave")

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
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastGongStructShapeOwner != nil {
					idx := slices.Index(pastGongStructShapeOwner.LinkShapes, linkshape_)
					pastGongStructShapeOwner.LinkShapes = slices.Delete(pastGongStructShapeOwner.LinkShapes, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _gongstructshape := range *models.GetGongstructInstancesSet[models.GongStructShape](linkshapeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _gongstructshape.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newGongStructShapeOwner := _gongstructshape // we have a match
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

	// manage the suppress operation
	if linkshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		linkshape_.Unstage(linkshapeFormCallback.probe.stageOfInterest)
	}

	linkshapeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.LinkShape](
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

	fillUpTree(linkshapeFormCallback.probe)
}
func __gong__New__NoteShapeFormCallback(
	noteshape *models.NoteShape,
	probe *Probe,
	formGroup *table.FormGroup,
) (noteshapeFormCallback *NoteShapeFormCallback) {
	noteshapeFormCallback = new(NoteShapeFormCallback)
	noteshapeFormCallback.probe = probe
	noteshapeFormCallback.noteshape = noteshape
	noteshapeFormCallback.formGroup = formGroup

	noteshapeFormCallback.CreationMode = (noteshape == nil)

	return
}

type NoteShapeFormCallback struct {
	noteshape *models.NoteShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (noteshapeFormCallback *NoteShapeFormCallback) OnSave() {

	log.Println("NoteShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	noteshapeFormCallback.probe.formStage.Checkout()

	if noteshapeFormCallback.noteshape == nil {
		noteshapeFormCallback.noteshape = new(models.NoteShape).Stage(noteshapeFormCallback.probe.stageOfInterest)
	}
	noteshape_ := noteshapeFormCallback.noteshape
	_ = noteshape_

	for _, formDiv := range noteshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(noteshape_.Name), formDiv)
		case "Identifier":
			FormDivBasicFieldToField(&(noteshape_.Identifier), formDiv)
		case "Body":
			FormDivBasicFieldToField(&(noteshape_.Body), formDiv)
		case "BodyHTML":
			FormDivBasicFieldToField(&(noteshape_.BodyHTML), formDiv)
		case "X":
			FormDivBasicFieldToField(&(noteshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(noteshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(noteshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(noteshape_.Height), formDiv)
		case "Matched":
			FormDivBasicFieldToField(&(noteshape_.Matched), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(noteshape_.IsExpanded), formDiv)
		case "Classdiagram:NoteShapes":
			// we need to retrieve the field owner before the change
			var pastClassdiagramOwner *models.Classdiagram
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Classdiagram"
			rf.Fieldname = "NoteShapes"
			reverseFieldOwner := models.GetReverseFieldOwner(
				noteshapeFormCallback.probe.stageOfInterest,
				noteshape_,
				&rf)

			if reverseFieldOwner != nil {
				pastClassdiagramOwner = reverseFieldOwner.(*models.Classdiagram)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastClassdiagramOwner != nil {
					idx := slices.Index(pastClassdiagramOwner.NoteShapes, noteshape_)
					pastClassdiagramOwner.NoteShapes = slices.Delete(pastClassdiagramOwner.NoteShapes, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _classdiagram := range *models.GetGongstructInstancesSet[models.Classdiagram](noteshapeFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _classdiagram.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newClassdiagramOwner := _classdiagram // we have a match
						if pastClassdiagramOwner != nil {
							if newClassdiagramOwner != pastClassdiagramOwner {
								idx := slices.Index(pastClassdiagramOwner.NoteShapes, noteshape_)
								pastClassdiagramOwner.NoteShapes = slices.Delete(pastClassdiagramOwner.NoteShapes, idx, idx+1)
								newClassdiagramOwner.NoteShapes = append(newClassdiagramOwner.NoteShapes, noteshape_)
							}
						} else {
							newClassdiagramOwner.NoteShapes = append(newClassdiagramOwner.NoteShapes, noteshape_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if noteshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteshape_.Unstage(noteshapeFormCallback.probe.stageOfInterest)
	}

	noteshapeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.NoteShape](
		noteshapeFormCallback.probe,
	)
	noteshapeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if noteshapeFormCallback.CreationMode || noteshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(noteshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__NoteShapeFormCallback(
			nil,
			noteshapeFormCallback.probe,
			newFormGroup,
		)
		noteshape := new(models.NoteShape)
		FillUpForm(noteshape, newFormGroup, noteshapeFormCallback.probe)
		noteshapeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(noteshapeFormCallback.probe)
}
func __gong__New__NoteShapeLinkFormCallback(
	noteshapelink *models.NoteShapeLink,
	probe *Probe,
	formGroup *table.FormGroup,
) (noteshapelinkFormCallback *NoteShapeLinkFormCallback) {
	noteshapelinkFormCallback = new(NoteShapeLinkFormCallback)
	noteshapelinkFormCallback.probe = probe
	noteshapelinkFormCallback.noteshapelink = noteshapelink
	noteshapelinkFormCallback.formGroup = formGroup

	noteshapelinkFormCallback.CreationMode = (noteshapelink == nil)

	return
}

type NoteShapeLinkFormCallback struct {
	noteshapelink *models.NoteShapeLink

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (noteshapelinkFormCallback *NoteShapeLinkFormCallback) OnSave() {

	log.Println("NoteShapeLinkFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	noteshapelinkFormCallback.probe.formStage.Checkout()

	if noteshapelinkFormCallback.noteshapelink == nil {
		noteshapelinkFormCallback.noteshapelink = new(models.NoteShapeLink).Stage(noteshapelinkFormCallback.probe.stageOfInterest)
	}
	noteshapelink_ := noteshapelinkFormCallback.noteshapelink
	_ = noteshapelink_

	for _, formDiv := range noteshapelinkFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(noteshapelink_.Name), formDiv)
		case "Identifier":
			FormDivBasicFieldToField(&(noteshapelink_.Identifier), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(noteshapelink_.Type), formDiv)
		case "NoteShape:NoteShapeLinks":
			// we need to retrieve the field owner before the change
			var pastNoteShapeOwner *models.NoteShape
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "NoteShape"
			rf.Fieldname = "NoteShapeLinks"
			reverseFieldOwner := models.GetReverseFieldOwner(
				noteshapelinkFormCallback.probe.stageOfInterest,
				noteshapelink_,
				&rf)

			if reverseFieldOwner != nil {
				pastNoteShapeOwner = reverseFieldOwner.(*models.NoteShape)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastNoteShapeOwner != nil {
					idx := slices.Index(pastNoteShapeOwner.NoteShapeLinks, noteshapelink_)
					pastNoteShapeOwner.NoteShapeLinks = slices.Delete(pastNoteShapeOwner.NoteShapeLinks, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _noteshape := range *models.GetGongstructInstancesSet[models.NoteShape](noteshapelinkFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _noteshape.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newNoteShapeOwner := _noteshape // we have a match
						if pastNoteShapeOwner != nil {
							if newNoteShapeOwner != pastNoteShapeOwner {
								idx := slices.Index(pastNoteShapeOwner.NoteShapeLinks, noteshapelink_)
								pastNoteShapeOwner.NoteShapeLinks = slices.Delete(pastNoteShapeOwner.NoteShapeLinks, idx, idx+1)
								newNoteShapeOwner.NoteShapeLinks = append(newNoteShapeOwner.NoteShapeLinks, noteshapelink_)
							}
						} else {
							newNoteShapeOwner.NoteShapeLinks = append(newNoteShapeOwner.NoteShapeLinks, noteshapelink_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if noteshapelinkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteshapelink_.Unstage(noteshapelinkFormCallback.probe.stageOfInterest)
	}

	noteshapelinkFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.NoteShapeLink](
		noteshapelinkFormCallback.probe,
	)
	noteshapelinkFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if noteshapelinkFormCallback.CreationMode || noteshapelinkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteshapelinkFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(noteshapelinkFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__NoteShapeLinkFormCallback(
			nil,
			noteshapelinkFormCallback.probe,
			newFormGroup,
		)
		noteshapelink := new(models.NoteShapeLink)
		FillUpForm(noteshapelink, newFormGroup, noteshapelinkFormCallback.probe)
		noteshapelinkFormCallback.probe.formStage.Commit()
	}

	fillUpTree(noteshapelinkFormCallback.probe)
}
