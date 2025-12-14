// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/dsme/cld/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__Category1FormCallback(
	category1 *models.Category1,
	probe *Probe,
	formGroup *table.FormGroup,
) (category1FormCallback *Category1FormCallback) {
	category1FormCallback = new(Category1FormCallback)
	category1FormCallback.probe = probe
	category1FormCallback.category1 = category1
	category1FormCallback.formGroup = formGroup

	category1FormCallback.CreationMode = (category1 == nil)

	return
}

type Category1FormCallback struct {
	category1 *models.Category1

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (category1FormCallback *Category1FormCallback) OnSave() {

	// log.Println("Category1FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	category1FormCallback.probe.formStage.Checkout()

	if category1FormCallback.category1 == nil {
		category1FormCallback.category1 = new(models.Category1).Stage(category1FormCallback.probe.stageOfInterest)
	}
	category1_ := category1FormCallback.category1
	_ = category1_

	for _, formDiv := range category1FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(category1_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if category1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		category1_.Unstage(category1FormCallback.probe.stageOfInterest)
	}

	category1FormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Category1](
		category1FormCallback.probe,
	)

	// display a new form by reset the form stage
	if category1FormCallback.CreationMode || category1FormCallback.formGroup.HasSuppressButtonBeenPressed {
		category1FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(category1FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Category1FormCallback(
			nil,
			category1FormCallback.probe,
			newFormGroup,
		)
		category1 := new(models.Category1)
		FillUpForm(category1, newFormGroup, category1FormCallback.probe)
		category1FormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(category1FormCallback.probe)
}
func __gong__New__Category1ShapeFormCallback(
	category1shape *models.Category1Shape,
	probe *Probe,
	formGroup *table.FormGroup,
) (category1shapeFormCallback *Category1ShapeFormCallback) {
	category1shapeFormCallback = new(Category1ShapeFormCallback)
	category1shapeFormCallback.probe = probe
	category1shapeFormCallback.category1shape = category1shape
	category1shapeFormCallback.formGroup = formGroup

	category1shapeFormCallback.CreationMode = (category1shape == nil)

	return
}

type Category1ShapeFormCallback struct {
	category1shape *models.Category1Shape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (category1shapeFormCallback *Category1ShapeFormCallback) OnSave() {

	// log.Println("Category1ShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	category1shapeFormCallback.probe.formStage.Checkout()

	if category1shapeFormCallback.category1shape == nil {
		category1shapeFormCallback.category1shape = new(models.Category1Shape).Stage(category1shapeFormCallback.probe.stageOfInterest)
	}
	category1shape_ := category1shapeFormCallback.category1shape
	_ = category1shape_

	for _, formDiv := range category1shapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(category1shape_.Name), formDiv)
		case "Category1":
			FormDivSelectFieldToField(&(category1shape_.Category1), category1shapeFormCallback.probe.stageOfInterest, formDiv)
		case "X":
			FormDivBasicFieldToField(&(category1shape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(category1shape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(category1shape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(category1shape_.Height), formDiv)
		case "Diagram:Category1Shapes":
			// WARNING : this form deals with the N-N association "Diagram.Category1Shapes []*Category1Shape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Category1Shape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "Category1Shapes"
				formerAssociationSource := category1shape_.GongGetReverseFieldOwner(
					category1shapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.Category1Shapes []*Category1Shape, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Category1Shapes, category1shape_)
					formerSource.Category1Shapes = slices.Delete(formerSource.Category1Shapes, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Diagram
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](category1shapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.Category1Shapes []*Category1Shape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Category1Shapes = append(newSource.Category1Shapes, category1shape_)
		}
	}

	// manage the suppress operation
	if category1shapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		category1shape_.Unstage(category1shapeFormCallback.probe.stageOfInterest)
	}

	category1shapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Category1Shape](
		category1shapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if category1shapeFormCallback.CreationMode || category1shapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		category1shapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(category1shapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Category1ShapeFormCallback(
			nil,
			category1shapeFormCallback.probe,
			newFormGroup,
		)
		category1shape := new(models.Category1Shape)
		FillUpForm(category1shape, newFormGroup, category1shapeFormCallback.probe)
		category1shapeFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(category1shapeFormCallback.probe)
}
func __gong__New__Category2FormCallback(
	category2 *models.Category2,
	probe *Probe,
	formGroup *table.FormGroup,
) (category2FormCallback *Category2FormCallback) {
	category2FormCallback = new(Category2FormCallback)
	category2FormCallback.probe = probe
	category2FormCallback.category2 = category2
	category2FormCallback.formGroup = formGroup

	category2FormCallback.CreationMode = (category2 == nil)

	return
}

type Category2FormCallback struct {
	category2 *models.Category2

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (category2FormCallback *Category2FormCallback) OnSave() {

	// log.Println("Category2FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	category2FormCallback.probe.formStage.Checkout()

	if category2FormCallback.category2 == nil {
		category2FormCallback.category2 = new(models.Category2).Stage(category2FormCallback.probe.stageOfInterest)
	}
	category2_ := category2FormCallback.category2
	_ = category2_

	for _, formDiv := range category2FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(category2_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if category2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		category2_.Unstage(category2FormCallback.probe.stageOfInterest)
	}

	category2FormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Category2](
		category2FormCallback.probe,
	)

	// display a new form by reset the form stage
	if category2FormCallback.CreationMode || category2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		category2FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(category2FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Category2FormCallback(
			nil,
			category2FormCallback.probe,
			newFormGroup,
		)
		category2 := new(models.Category2)
		FillUpForm(category2, newFormGroup, category2FormCallback.probe)
		category2FormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(category2FormCallback.probe)
}
func __gong__New__Category2ShapeFormCallback(
	category2shape *models.Category2Shape,
	probe *Probe,
	formGroup *table.FormGroup,
) (category2shapeFormCallback *Category2ShapeFormCallback) {
	category2shapeFormCallback = new(Category2ShapeFormCallback)
	category2shapeFormCallback.probe = probe
	category2shapeFormCallback.category2shape = category2shape
	category2shapeFormCallback.formGroup = formGroup

	category2shapeFormCallback.CreationMode = (category2shape == nil)

	return
}

type Category2ShapeFormCallback struct {
	category2shape *models.Category2Shape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (category2shapeFormCallback *Category2ShapeFormCallback) OnSave() {

	// log.Println("Category2ShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	category2shapeFormCallback.probe.formStage.Checkout()

	if category2shapeFormCallback.category2shape == nil {
		category2shapeFormCallback.category2shape = new(models.Category2Shape).Stage(category2shapeFormCallback.probe.stageOfInterest)
	}
	category2shape_ := category2shapeFormCallback.category2shape
	_ = category2shape_

	for _, formDiv := range category2shapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(category2shape_.Name), formDiv)
		case "Category2":
			FormDivSelectFieldToField(&(category2shape_.Category2), category2shapeFormCallback.probe.stageOfInterest, formDiv)
		case "X":
			FormDivBasicFieldToField(&(category2shape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(category2shape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(category2shape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(category2shape_.Height), formDiv)
		case "Diagram:Category2Shapes":
			// WARNING : this form deals with the N-N association "Diagram.Category2Shapes []*Category2Shape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Category2Shape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "Category2Shapes"
				formerAssociationSource := category2shape_.GongGetReverseFieldOwner(
					category2shapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.Category2Shapes []*Category2Shape, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Category2Shapes, category2shape_)
					formerSource.Category2Shapes = slices.Delete(formerSource.Category2Shapes, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Diagram
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](category2shapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.Category2Shapes []*Category2Shape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Category2Shapes = append(newSource.Category2Shapes, category2shape_)
		}
	}

	// manage the suppress operation
	if category2shapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		category2shape_.Unstage(category2shapeFormCallback.probe.stageOfInterest)
	}

	category2shapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Category2Shape](
		category2shapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if category2shapeFormCallback.CreationMode || category2shapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		category2shapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(category2shapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Category2ShapeFormCallback(
			nil,
			category2shapeFormCallback.probe,
			newFormGroup,
		)
		category2shape := new(models.Category2Shape)
		FillUpForm(category2shape, newFormGroup, category2shapeFormCallback.probe)
		category2shapeFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(category2shapeFormCallback.probe)
}
func __gong__New__Category3FormCallback(
	category3 *models.Category3,
	probe *Probe,
	formGroup *table.FormGroup,
) (category3FormCallback *Category3FormCallback) {
	category3FormCallback = new(Category3FormCallback)
	category3FormCallback.probe = probe
	category3FormCallback.category3 = category3
	category3FormCallback.formGroup = formGroup

	category3FormCallback.CreationMode = (category3 == nil)

	return
}

type Category3FormCallback struct {
	category3 *models.Category3

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (category3FormCallback *Category3FormCallback) OnSave() {

	// log.Println("Category3FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	category3FormCallback.probe.formStage.Checkout()

	if category3FormCallback.category3 == nil {
		category3FormCallback.category3 = new(models.Category3).Stage(category3FormCallback.probe.stageOfInterest)
	}
	category3_ := category3FormCallback.category3
	_ = category3_

	for _, formDiv := range category3FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(category3_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if category3FormCallback.formGroup.HasSuppressButtonBeenPressed {
		category3_.Unstage(category3FormCallback.probe.stageOfInterest)
	}

	category3FormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Category3](
		category3FormCallback.probe,
	)

	// display a new form by reset the form stage
	if category3FormCallback.CreationMode || category3FormCallback.formGroup.HasSuppressButtonBeenPressed {
		category3FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(category3FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Category3FormCallback(
			nil,
			category3FormCallback.probe,
			newFormGroup,
		)
		category3 := new(models.Category3)
		FillUpForm(category3, newFormGroup, category3FormCallback.probe)
		category3FormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(category3FormCallback.probe)
}
func __gong__New__Category3ShapeFormCallback(
	category3shape *models.Category3Shape,
	probe *Probe,
	formGroup *table.FormGroup,
) (category3shapeFormCallback *Category3ShapeFormCallback) {
	category3shapeFormCallback = new(Category3ShapeFormCallback)
	category3shapeFormCallback.probe = probe
	category3shapeFormCallback.category3shape = category3shape
	category3shapeFormCallback.formGroup = formGroup

	category3shapeFormCallback.CreationMode = (category3shape == nil)

	return
}

type Category3ShapeFormCallback struct {
	category3shape *models.Category3Shape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (category3shapeFormCallback *Category3ShapeFormCallback) OnSave() {

	// log.Println("Category3ShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	category3shapeFormCallback.probe.formStage.Checkout()

	if category3shapeFormCallback.category3shape == nil {
		category3shapeFormCallback.category3shape = new(models.Category3Shape).Stage(category3shapeFormCallback.probe.stageOfInterest)
	}
	category3shape_ := category3shapeFormCallback.category3shape
	_ = category3shape_

	for _, formDiv := range category3shapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(category3shape_.Name), formDiv)
		case "Category3":
			FormDivSelectFieldToField(&(category3shape_.Category3), category3shapeFormCallback.probe.stageOfInterest, formDiv)
		case "X":
			FormDivBasicFieldToField(&(category3shape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(category3shape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(category3shape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(category3shape_.Height), formDiv)
		case "Diagram:Category3Shapes":
			// WARNING : this form deals with the N-N association "Diagram.Category3Shapes []*Category3Shape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Category3Shape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "Category3Shapes"
				formerAssociationSource := category3shape_.GongGetReverseFieldOwner(
					category3shapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.Category3Shapes []*Category3Shape, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Category3Shapes, category3shape_)
					formerSource.Category3Shapes = slices.Delete(formerSource.Category3Shapes, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Diagram
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](category3shapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.Category3Shapes []*Category3Shape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Category3Shapes = append(newSource.Category3Shapes, category3shape_)
		}
	}

	// manage the suppress operation
	if category3shapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		category3shape_.Unstage(category3shapeFormCallback.probe.stageOfInterest)
	}

	category3shapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Category3Shape](
		category3shapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if category3shapeFormCallback.CreationMode || category3shapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		category3shapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(category3shapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Category3ShapeFormCallback(
			nil,
			category3shapeFormCallback.probe,
			newFormGroup,
		)
		category3shape := new(models.Category3Shape)
		FillUpForm(category3shape, newFormGroup, category3shapeFormCallback.probe)
		category3shapeFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(category3shapeFormCallback.probe)
}
func __gong__New__ControlPointShapeFormCallback(
	controlpointshape *models.ControlPointShape,
	probe *Probe,
	formGroup *table.FormGroup,
) (controlpointshapeFormCallback *ControlPointShapeFormCallback) {
	controlpointshapeFormCallback = new(ControlPointShapeFormCallback)
	controlpointshapeFormCallback.probe = probe
	controlpointshapeFormCallback.controlpointshape = controlpointshape
	controlpointshapeFormCallback.formGroup = formGroup

	controlpointshapeFormCallback.CreationMode = (controlpointshape == nil)

	return
}

type ControlPointShapeFormCallback struct {
	controlpointshape *models.ControlPointShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (controlpointshapeFormCallback *ControlPointShapeFormCallback) OnSave() {

	// log.Println("ControlPointShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	controlpointshapeFormCallback.probe.formStage.Checkout()

	if controlpointshapeFormCallback.controlpointshape == nil {
		controlpointshapeFormCallback.controlpointshape = new(models.ControlPointShape).Stage(controlpointshapeFormCallback.probe.stageOfInterest)
	}
	controlpointshape_ := controlpointshapeFormCallback.controlpointshape
	_ = controlpointshape_

	for _, formDiv := range controlpointshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(controlpointshape_.Name), formDiv)
		case "X_Relative":
			FormDivBasicFieldToField(&(controlpointshape_.X_Relative), formDiv)
		case "Y_Relative":
			FormDivBasicFieldToField(&(controlpointshape_.Y_Relative), formDiv)
		case "IsStartShapeTheClosestShape":
			FormDivBasicFieldToField(&(controlpointshape_.IsStartShapeTheClosestShape), formDiv)
		case "InfluenceShape:ControlPointShapes":
			// WARNING : this form deals with the N-N association "InfluenceShape.ControlPointShapes []*ControlPointShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ControlPointShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.InfluenceShape
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "InfluenceShape"
				rf.Fieldname = "ControlPointShapes"
				formerAssociationSource := controlpointshape_.GongGetReverseFieldOwner(
					controlpointshapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.InfluenceShape)
					if !ok {
						log.Fatalln("Source of InfluenceShape.ControlPointShapes []*ControlPointShape, is not an InfluenceShape instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ControlPointShapes, controlpointshape_)
					formerSource.ControlPointShapes = slices.Delete(formerSource.ControlPointShapes, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.InfluenceShape
			for _influenceshape := range *models.GetGongstructInstancesSet[models.InfluenceShape](controlpointshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _influenceshape.GetName() == newSourceName.GetName() {
					newSource = _influenceshape // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of InfluenceShape.ControlPointShapes []*ControlPointShape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ControlPointShapes = append(newSource.ControlPointShapes, controlpointshape_)
		}
	}

	// manage the suppress operation
	if controlpointshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		controlpointshape_.Unstage(controlpointshapeFormCallback.probe.stageOfInterest)
	}

	controlpointshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ControlPointShape](
		controlpointshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if controlpointshapeFormCallback.CreationMode || controlpointshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		controlpointshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(controlpointshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ControlPointShapeFormCallback(
			nil,
			controlpointshapeFormCallback.probe,
			newFormGroup,
		)
		controlpointshape := new(models.ControlPointShape)
		FillUpForm(controlpointshape, newFormGroup, controlpointshapeFormCallback.probe)
		controlpointshapeFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(controlpointshapeFormCallback.probe)
}
func __gong__New__DeskFormCallback(
	desk *models.Desk,
	probe *Probe,
	formGroup *table.FormGroup,
) (deskFormCallback *DeskFormCallback) {
	deskFormCallback = new(DeskFormCallback)
	deskFormCallback.probe = probe
	deskFormCallback.desk = desk
	deskFormCallback.formGroup = formGroup

	deskFormCallback.CreationMode = (desk == nil)

	return
}

type DeskFormCallback struct {
	desk *models.Desk

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (deskFormCallback *DeskFormCallback) OnSave() {

	// log.Println("DeskFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	deskFormCallback.probe.formStage.Checkout()

	if deskFormCallback.desk == nil {
		deskFormCallback.desk = new(models.Desk).Stage(deskFormCallback.probe.stageOfInterest)
	}
	desk_ := deskFormCallback.desk
	_ = desk_

	for _, formDiv := range deskFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(desk_.Name), formDiv)
		case "SelectedDiagram":
			FormDivSelectFieldToField(&(desk_.SelectedDiagram), deskFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if deskFormCallback.formGroup.HasSuppressButtonBeenPressed {
		desk_.Unstage(deskFormCallback.probe.stageOfInterest)
	}

	deskFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Desk](
		deskFormCallback.probe,
	)

	// display a new form by reset the form stage
	if deskFormCallback.CreationMode || deskFormCallback.formGroup.HasSuppressButtonBeenPressed {
		deskFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(deskFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DeskFormCallback(
			nil,
			deskFormCallback.probe,
			newFormGroup,
		)
		desk := new(models.Desk)
		FillUpForm(desk, newFormGroup, deskFormCallback.probe)
		deskFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(deskFormCallback.probe)
}
func __gong__New__DiagramFormCallback(
	diagram *models.Diagram,
	probe *Probe,
	formGroup *table.FormGroup,
) (diagramFormCallback *DiagramFormCallback) {
	diagramFormCallback = new(DiagramFormCallback)
	diagramFormCallback.probe = probe
	diagramFormCallback.diagram = diagram
	diagramFormCallback.formGroup = formGroup

	diagramFormCallback.CreationMode = (diagram == nil)

	return
}

type DiagramFormCallback struct {
	diagram *models.Diagram

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (diagramFormCallback *DiagramFormCallback) OnSave() {

	// log.Println("DiagramFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	diagramFormCallback.probe.formStage.Checkout()

	if diagramFormCallback.diagram == nil {
		diagramFormCallback.diagram = new(models.Diagram).Stage(diagramFormCallback.probe.stageOfInterest)
	}
	diagram_ := diagramFormCallback.diagram
	_ = diagram_

	for _, formDiv := range diagramFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(diagram_.Name), formDiv)
		case "Category1Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Category1Shape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Category1Shape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Category1Shape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			diagram_.Category1Shapes = instanceSlice

		case "Category2Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Category2Shape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Category2Shape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Category2Shape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			diagram_.Category2Shapes = instanceSlice

		case "Category3Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Category3Shape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Category3Shape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Category3Shape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			diagram_.Category3Shapes = instanceSlice

		case "InfluenceShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.InfluenceShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.InfluenceShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.InfluenceShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			diagram_.InfluenceShapes = instanceSlice

		case "IsEditable":
			FormDivBasicFieldToField(&(diagram_.IsEditable), formDiv)
		case "IsNodeExpanded":
			FormDivBasicFieldToField(&(diagram_.IsNodeExpanded), formDiv)
		case "IsCategory1NodeExpanded":
			FormDivBasicFieldToField(&(diagram_.IsCategory1NodeExpanded), formDiv)
		case "IsCategory2NodeExpanded":
			FormDivBasicFieldToField(&(diagram_.IsCategory2NodeExpanded), formDiv)
		case "IsCategory3NodeExpanded":
			FormDivBasicFieldToField(&(diagram_.IsCategory3NodeExpanded), formDiv)
		case "IsInfluenceCategoryNodeExpanded":
			FormDivBasicFieldToField(&(diagram_.IsInfluenceCategoryNodeExpanded), formDiv)
		case "IsCategory1Shown":
			FormDivBasicFieldToField(&(diagram_.IsCategory1Shown), formDiv)
		case "IsCategory2Shown":
			FormDivBasicFieldToField(&(diagram_.IsCategory2Shown), formDiv)
		case "IsCategory3Shown":
			FormDivBasicFieldToField(&(diagram_.IsCategory3Shown), formDiv)
		case "IsInfluenceCategoryShown":
			FormDivBasicFieldToField(&(diagram_.IsInfluenceCategoryShown), formDiv)
		case "XMargin":
			FormDivBasicFieldToField(&(diagram_.XMargin), formDiv)
		case "YMargin":
			FormDivBasicFieldToField(&(diagram_.YMargin), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(diagram_.Height), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(diagram_.Width), formDiv)
		case "RedColorCode":
			FormDivBasicFieldToField(&(diagram_.RedColorCode), formDiv)
		case "BackgroundGreyColorCode":
			FormDivBasicFieldToField(&(diagram_.BackgroundGreyColorCode), formDiv)
		case "GrayColorCode":
			FormDivBasicFieldToField(&(diagram_.GrayColorCode), formDiv)
		case "Category1RectAnchorType":
			FormDivEnumStringFieldToField(&(diagram_.Category1RectAnchorType), formDiv)
		case "Category1TextAnchorType":
			FormDivEnumStringFieldToField(&(diagram_.Category1TextAnchorType), formDiv)
		case "Category1DominantBaselineType":
			FormDivEnumStringFieldToField(&(diagram_.Category1DominantBaselineType), formDiv)
		case "Category1FontSize":
			FormDivBasicFieldToField(&(diagram_.Category1FontSize), formDiv)
		case "Category1FontWeigth":
			FormDivBasicFieldToField(&(diagram_.Category1FontWeigth), formDiv)
		case "Category1FontFamily":
			FormDivBasicFieldToField(&(diagram_.Category1FontFamily), formDiv)
		case "Category1LetterSpacing":
			FormDivBasicFieldToField(&(diagram_.Category1LetterSpacing), formDiv)
		case "Category2TypeFontSize":
			FormDivBasicFieldToField(&(diagram_.Category2TypeFontSize), formDiv)
		case "Category2TypeFontWeigth":
			FormDivBasicFieldToField(&(diagram_.Category2TypeFontWeigth), formDiv)
		case "Category2TypeFontFamily":
			FormDivBasicFieldToField(&(diagram_.Category2TypeFontFamily), formDiv)
		case "Category2TypeLetterSpacing":
			FormDivBasicFieldToField(&(diagram_.Category2TypeLetterSpacing), formDiv)
		case "Category2TypeRectAnchorType":
			FormDivEnumStringFieldToField(&(diagram_.Category2TypeRectAnchorType), formDiv)
		case "Category2DominantBaselineType":
			FormDivEnumStringFieldToField(&(diagram_.Category2DominantBaselineType), formDiv)
		case "Category2StrokeWidth":
			FormDivBasicFieldToField(&(diagram_.Category2StrokeWidth), formDiv)
		case "Category3RectAnchorType":
			FormDivEnumStringFieldToField(&(diagram_.Category3RectAnchorType), formDiv)
		case "Category3TextAnchorType":
			FormDivEnumStringFieldToField(&(diagram_.Category3TextAnchorType), formDiv)
		case "Category3DominantBaselineType":
			FormDivEnumStringFieldToField(&(diagram_.Category3DominantBaselineType), formDiv)
		case "Category3FontSize":
			FormDivBasicFieldToField(&(diagram_.Category3FontSize), formDiv)
		case "Category3FontWeigth":
			FormDivBasicFieldToField(&(diagram_.Category3FontWeigth), formDiv)
		case "Category3FontFamily":
			FormDivBasicFieldToField(&(diagram_.Category3FontFamily), formDiv)
		case "Category3LetterSpacing":
			FormDivBasicFieldToField(&(diagram_.Category3LetterSpacing), formDiv)
		case "InfluenceArrowSize":
			FormDivBasicFieldToField(&(diagram_.InfluenceArrowSize), formDiv)
		case "InfluenceArrowStartOffset":
			FormDivBasicFieldToField(&(diagram_.InfluenceArrowStartOffset), formDiv)
		case "InfluenceArrowEndOffset":
			FormDivBasicFieldToField(&(diagram_.InfluenceArrowEndOffset), formDiv)
		case "InfluenceCornerRadius":
			FormDivBasicFieldToField(&(diagram_.InfluenceCornerRadius), formDiv)
		case "InfluenceDashedLinePattern":
			FormDivBasicFieldToField(&(diagram_.InfluenceDashedLinePattern), formDiv)
		}
	}

	// manage the suppress operation
	if diagramFormCallback.formGroup.HasSuppressButtonBeenPressed {
		diagram_.Unstage(diagramFormCallback.probe.stageOfInterest)
	}

	diagramFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Diagram](
		diagramFormCallback.probe,
	)

	// display a new form by reset the form stage
	if diagramFormCallback.CreationMode || diagramFormCallback.formGroup.HasSuppressButtonBeenPressed {
		diagramFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(diagramFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DiagramFormCallback(
			nil,
			diagramFormCallback.probe,
			newFormGroup,
		)
		diagram := new(models.Diagram)
		FillUpForm(diagram, newFormGroup, diagramFormCallback.probe)
		diagramFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(diagramFormCallback.probe)
}
func __gong__New__InfluenceFormCallback(
	influence *models.Influence,
	probe *Probe,
	formGroup *table.FormGroup,
) (influenceFormCallback *InfluenceFormCallback) {
	influenceFormCallback = new(InfluenceFormCallback)
	influenceFormCallback.probe = probe
	influenceFormCallback.influence = influence
	influenceFormCallback.formGroup = formGroup

	influenceFormCallback.CreationMode = (influence == nil)

	return
}

type InfluenceFormCallback struct {
	influence *models.Influence

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (influenceFormCallback *InfluenceFormCallback) OnSave() {

	// log.Println("InfluenceFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	influenceFormCallback.probe.formStage.Checkout()

	if influenceFormCallback.influence == nil {
		influenceFormCallback.influence = new(models.Influence).Stage(influenceFormCallback.probe.stageOfInterest)
	}
	influence_ := influenceFormCallback.influence
	_ = influence_

	for _, formDiv := range influenceFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(influence_.Name), formDiv)
		case "SourceCategory1":
			FormDivSelectFieldToField(&(influence_.SourceCategory1), influenceFormCallback.probe.stageOfInterest, formDiv)
		case "SourceCategory2":
			FormDivSelectFieldToField(&(influence_.SourceCategory2), influenceFormCallback.probe.stageOfInterest, formDiv)
		case "SourceCategory3":
			FormDivSelectFieldToField(&(influence_.SourceCategory3), influenceFormCallback.probe.stageOfInterest, formDiv)
		case "TargetCategory1":
			FormDivSelectFieldToField(&(influence_.TargetCategory1), influenceFormCallback.probe.stageOfInterest, formDiv)
		case "TargetCategory2":
			FormDivSelectFieldToField(&(influence_.TargetCategory2), influenceFormCallback.probe.stageOfInterest, formDiv)
		case "TargetCategory3":
			FormDivSelectFieldToField(&(influence_.TargetCategory3), influenceFormCallback.probe.stageOfInterest, formDiv)
		case "IsHypothtical":
			FormDivBasicFieldToField(&(influence_.IsHypothtical), formDiv)
		}
	}

	// manage the suppress operation
	if influenceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		influence_.Unstage(influenceFormCallback.probe.stageOfInterest)
	}

	influenceFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Influence](
		influenceFormCallback.probe,
	)

	// display a new form by reset the form stage
	if influenceFormCallback.CreationMode || influenceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		influenceFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(influenceFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__InfluenceFormCallback(
			nil,
			influenceFormCallback.probe,
			newFormGroup,
		)
		influence := new(models.Influence)
		FillUpForm(influence, newFormGroup, influenceFormCallback.probe)
		influenceFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(influenceFormCallback.probe)
}
func __gong__New__InfluenceShapeFormCallback(
	influenceshape *models.InfluenceShape,
	probe *Probe,
	formGroup *table.FormGroup,
) (influenceshapeFormCallback *InfluenceShapeFormCallback) {
	influenceshapeFormCallback = new(InfluenceShapeFormCallback)
	influenceshapeFormCallback.probe = probe
	influenceshapeFormCallback.influenceshape = influenceshape
	influenceshapeFormCallback.formGroup = formGroup

	influenceshapeFormCallback.CreationMode = (influenceshape == nil)

	return
}

type InfluenceShapeFormCallback struct {
	influenceshape *models.InfluenceShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (influenceshapeFormCallback *InfluenceShapeFormCallback) OnSave() {

	// log.Println("InfluenceShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	influenceshapeFormCallback.probe.formStage.Checkout()

	if influenceshapeFormCallback.influenceshape == nil {
		influenceshapeFormCallback.influenceshape = new(models.InfluenceShape).Stage(influenceshapeFormCallback.probe.stageOfInterest)
	}
	influenceshape_ := influenceshapeFormCallback.influenceshape
	_ = influenceshape_

	for _, formDiv := range influenceshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(influenceshape_.Name), formDiv)
		case "Influence":
			FormDivSelectFieldToField(&(influenceshape_.Influence), influenceshapeFormCallback.probe.stageOfInterest, formDiv)
		case "ControlPointShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ControlPointShape](influenceshapeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ControlPointShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ControlPointShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					influenceshapeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			ids, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			for _, id := range ids {
				instanceSlice = append(instanceSlice, map_id_instances[id])
			}
			influenceshape_.ControlPointShapes = instanceSlice

		case "Diagram:InfluenceShapes":
			// WARNING : this form deals with the N-N association "Diagram.InfluenceShapes []*InfluenceShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of InfluenceShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Diagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Diagram"
				rf.Fieldname = "InfluenceShapes"
				formerAssociationSource := influenceshape_.GongGetReverseFieldOwner(
					influenceshapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Diagram)
					if !ok {
						log.Fatalln("Source of Diagram.InfluenceShapes []*InfluenceShape, is not an Diagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.InfluenceShapes, influenceshape_)
					formerSource.InfluenceShapes = slices.Delete(formerSource.InfluenceShapes, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Diagram
			for _diagram := range *models.GetGongstructInstancesSet[models.Diagram](influenceshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagram.GetName() == newSourceName.GetName() {
					newSource = _diagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Diagram.InfluenceShapes []*InfluenceShape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.InfluenceShapes = append(newSource.InfluenceShapes, influenceshape_)
		}
	}

	// manage the suppress operation
	if influenceshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		influenceshape_.Unstage(influenceshapeFormCallback.probe.stageOfInterest)
	}

	influenceshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.InfluenceShape](
		influenceshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if influenceshapeFormCallback.CreationMode || influenceshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		influenceshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(influenceshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__InfluenceShapeFormCallback(
			nil,
			influenceshapeFormCallback.probe,
			newFormGroup,
		)
		influenceshape := new(models.InfluenceShape)
		FillUpForm(influenceshape, newFormGroup, influenceshapeFormCallback.probe)
		influenceshapeFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(influenceshapeFormCallback.probe)
}
