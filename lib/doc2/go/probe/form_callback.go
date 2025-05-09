// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/doc2/go/models"
)

// code to avoid error when generated code does not need to import packages
const __dummmy__time = time.Nanosecond

var _ = __dummmy__time

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)

var _ = __dummmy__letters

const __dummy__log = log.Ldate

var _ = __dummy__log

// end of code to avoid error when generated code does not need to import packages

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
			// WARNING : this form deals with the N-N association "GongStructShape.AttributeShapes []*AttributeShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of AttributeShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.GongStructShape
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "GongStructShape"
				rf.Fieldname = "AttributeShapes"
				formerAssociationSource := models.GetReverseFieldOwner(
					attributeshapeFormCallback.probe.stageOfInterest,
					attributeshape_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.GongStructShape)
					if !ok {
						log.Fatalln("Source of GongStructShape.AttributeShapes []*AttributeShape, is not an GongStructShape instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.AttributeShapes, attributeshape_)
					formerSource.AttributeShapes = slices.Delete(formerSource.AttributeShapes, idx, idx+1)
				}
				// That could mean we clear the assocation for all source instances
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.GongStructShape
			for _gongstructshape := range *models.GetGongstructInstancesSet[models.GongStructShape](attributeshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _gongstructshape.GetName() == newSourceName.GetName() {
					newSource = _gongstructshape // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of GongStructShape.AttributeShapes []*AttributeShape, with name", newSourceName, ", does not exist")
				break
			}

			// append the value to the new source field
			newSource.AttributeShapes = append(newSource.AttributeShapes, attributeshape_)
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
		case "Description":
			FormDivBasicFieldToField(&(classdiagram_.Description), formDiv)
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
			// WARNING : this form deals with the N-N association "DiagramPackage.Classdiagrams []*Classdiagram" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Classdiagram). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.DiagramPackage
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "DiagramPackage"
				rf.Fieldname = "Classdiagrams"
				formerAssociationSource := models.GetReverseFieldOwner(
					classdiagramFormCallback.probe.stageOfInterest,
					classdiagram_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.DiagramPackage)
					if !ok {
						log.Fatalln("Source of DiagramPackage.Classdiagrams []*Classdiagram, is not an DiagramPackage instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.Classdiagrams, classdiagram_)
					formerSource.Classdiagrams = slices.Delete(formerSource.Classdiagrams, idx, idx+1)
				}
				// That could mean we clear the assocation for all source instances
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.DiagramPackage
			for _diagrampackage := range *models.GetGongstructInstancesSet[models.DiagramPackage](classdiagramFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagrampackage.GetName() == newSourceName.GetName() {
					newSource = _diagrampackage // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of DiagramPackage.Classdiagrams []*Classdiagram, with name", newSourceName, ", does not exist")
				break
			}

			// append the value to the new source field
			newSource.Classdiagrams = append(newSource.Classdiagrams, classdiagram_)
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
			// WARNING : this form deals with the N-N association "Classdiagram.GongEnumShapes []*GongEnumShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of GongEnumShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Classdiagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Classdiagram"
				rf.Fieldname = "GongEnumShapes"
				formerAssociationSource := models.GetReverseFieldOwner(
					gongenumshapeFormCallback.probe.stageOfInterest,
					gongenumshape_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Classdiagram)
					if !ok {
						log.Fatalln("Source of Classdiagram.GongEnumShapes []*GongEnumShape, is not an Classdiagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.GongEnumShapes, gongenumshape_)
					formerSource.GongEnumShapes = slices.Delete(formerSource.GongEnumShapes, idx, idx+1)
				}
				// That could mean we clear the assocation for all source instances
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Classdiagram
			for _classdiagram := range *models.GetGongstructInstancesSet[models.Classdiagram](gongenumshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _classdiagram.GetName() == newSourceName.GetName() {
					newSource = _classdiagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Classdiagram.GongEnumShapes []*GongEnumShape, with name", newSourceName, ", does not exist")
				break
			}

			// append the value to the new source field
			newSource.GongEnumShapes = append(newSource.GongEnumShapes, gongenumshape_)
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
			// WARNING : this form deals with the N-N association "GongEnumShape.GongEnumValueShapes []*GongEnumValueShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of GongEnumValueShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.GongEnumShape
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "GongEnumShape"
				rf.Fieldname = "GongEnumValueShapes"
				formerAssociationSource := models.GetReverseFieldOwner(
					gongenumvalueshapeFormCallback.probe.stageOfInterest,
					gongenumvalueshape_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.GongEnumShape)
					if !ok {
						log.Fatalln("Source of GongEnumShape.GongEnumValueShapes []*GongEnumValueShape, is not an GongEnumShape instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.GongEnumValueShapes, gongenumvalueshape_)
					formerSource.GongEnumValueShapes = slices.Delete(formerSource.GongEnumValueShapes, idx, idx+1)
				}
				// That could mean we clear the assocation for all source instances
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.GongEnumShape
			for _gongenumshape := range *models.GetGongstructInstancesSet[models.GongEnumShape](gongenumvalueshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _gongenumshape.GetName() == newSourceName.GetName() {
					newSource = _gongenumshape // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of GongEnumShape.GongEnumValueShapes []*GongEnumValueShape, with name", newSourceName, ", does not exist")
				break
			}

			// append the value to the new source field
			newSource.GongEnumValueShapes = append(newSource.GongEnumValueShapes, gongenumvalueshape_)
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
			// WARNING : this form deals with the N-N association "GongNoteShape.GongNoteLinkShapes []*GongNoteLinkShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of GongNoteLinkShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.GongNoteShape
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "GongNoteShape"
				rf.Fieldname = "GongNoteLinkShapes"
				formerAssociationSource := models.GetReverseFieldOwner(
					gongnotelinkshapeFormCallback.probe.stageOfInterest,
					gongnotelinkshape_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.GongNoteShape)
					if !ok {
						log.Fatalln("Source of GongNoteShape.GongNoteLinkShapes []*GongNoteLinkShape, is not an GongNoteShape instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.GongNoteLinkShapes, gongnotelinkshape_)
					formerSource.GongNoteLinkShapes = slices.Delete(formerSource.GongNoteLinkShapes, idx, idx+1)
				}
				// That could mean we clear the assocation for all source instances
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.GongNoteShape
			for _gongnoteshape := range *models.GetGongstructInstancesSet[models.GongNoteShape](gongnotelinkshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _gongnoteshape.GetName() == newSourceName.GetName() {
					newSource = _gongnoteshape // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of GongNoteShape.GongNoteLinkShapes []*GongNoteLinkShape, with name", newSourceName, ", does not exist")
				break
			}

			// append the value to the new source field
			newSource.GongNoteLinkShapes = append(newSource.GongNoteLinkShapes, gongnotelinkshape_)
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
			// WARNING : this form deals with the N-N association "Classdiagram.GongNoteShapes []*GongNoteShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of GongNoteShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Classdiagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Classdiagram"
				rf.Fieldname = "GongNoteShapes"
				formerAssociationSource := models.GetReverseFieldOwner(
					gongnoteshapeFormCallback.probe.stageOfInterest,
					gongnoteshape_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Classdiagram)
					if !ok {
						log.Fatalln("Source of Classdiagram.GongNoteShapes []*GongNoteShape, is not an Classdiagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.GongNoteShapes, gongnoteshape_)
					formerSource.GongNoteShapes = slices.Delete(formerSource.GongNoteShapes, idx, idx+1)
				}
				// That could mean we clear the assocation for all source instances
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Classdiagram
			for _classdiagram := range *models.GetGongstructInstancesSet[models.Classdiagram](gongnoteshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _classdiagram.GetName() == newSourceName.GetName() {
					newSource = _classdiagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Classdiagram.GongNoteShapes []*GongNoteShape, with name", newSourceName, ", does not exist")
				break
			}

			// append the value to the new source field
			newSource.GongNoteShapes = append(newSource.GongNoteShapes, gongnoteshape_)
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
			// WARNING : this form deals with the N-N association "Classdiagram.GongStructShapes []*GongStructShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of GongStructShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Classdiagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Classdiagram"
				rf.Fieldname = "GongStructShapes"
				formerAssociationSource := models.GetReverseFieldOwner(
					gongstructshapeFormCallback.probe.stageOfInterest,
					gongstructshape_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Classdiagram)
					if !ok {
						log.Fatalln("Source of Classdiagram.GongStructShapes []*GongStructShape, is not an Classdiagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.GongStructShapes, gongstructshape_)
					formerSource.GongStructShapes = slices.Delete(formerSource.GongStructShapes, idx, idx+1)
				}
				// That could mean we clear the assocation for all source instances
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Classdiagram
			for _classdiagram := range *models.GetGongstructInstancesSet[models.Classdiagram](gongstructshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _classdiagram.GetName() == newSourceName.GetName() {
					newSource = _classdiagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Classdiagram.GongStructShapes []*GongStructShape, with name", newSourceName, ", does not exist")
				break
			}

			// append the value to the new source field
			newSource.GongStructShapes = append(newSource.GongStructShapes, gongstructshape_)
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
			// WARNING : this form deals with the N-N association "GongStructShape.LinkShapes []*LinkShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of LinkShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.GongStructShape
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "GongStructShape"
				rf.Fieldname = "LinkShapes"
				formerAssociationSource := models.GetReverseFieldOwner(
					linkshapeFormCallback.probe.stageOfInterest,
					linkshape_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.GongStructShape)
					if !ok {
						log.Fatalln("Source of GongStructShape.LinkShapes []*LinkShape, is not an GongStructShape instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.LinkShapes, linkshape_)
					formerSource.LinkShapes = slices.Delete(formerSource.LinkShapes, idx, idx+1)
				}
				// That could mean we clear the assocation for all source instances
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.GongStructShape
			for _gongstructshape := range *models.GetGongstructInstancesSet[models.GongStructShape](linkshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _gongstructshape.GetName() == newSourceName.GetName() {
					newSource = _gongstructshape // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of GongStructShape.LinkShapes []*LinkShape, with name", newSourceName, ", does not exist")
				break
			}

			// append the value to the new source field
			newSource.LinkShapes = append(newSource.LinkShapes, linkshape_)
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
