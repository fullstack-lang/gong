// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
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
		case "IsInDrawMode":
			FormDivBasicFieldToField(&(classdiagram_.IsInDrawMode), formDiv)
		case "DiagramPackage:Classdiagrams":
			// WARNING : this form deals with the N-N association "DiagramPackage.Classdiagrams []*Classdiagram" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Classdiagram). Setting up a value
			// will discard the former value is there is one.
			//
			// the algorithm is
			// 1/ get the former source of the association
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
				break // nothing else to do for this field
			}

			// we need to deal with the 2 cases:
			// 1 the field source is unchanged
			// 2 the field source is changed

			// 1 field source is unchanged
			if formerSource != nil && formerSource.GetName() == newSourceName.GetName() {
				break // nothing else to do for this field
			}

			// 2 field source is changed -->
			// (1) clear the source slice field if it exist
			// (2) find the new source
			// (3) append the new value to the new source field

			// (1) clear the source slice field if it exist
			if formerSource != nil {
				idx := slices.Index(formerSource.Classdiagrams, classdiagram_)
				formerSource.Classdiagrams = slices.Delete(formerSource.Classdiagrams, idx, idx+1)
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

			// (3) append the new value to the new source field
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
func __gong__New__FieldFormCallback(
	field *models.Field,
	probe *Probe,
	formGroup *table.FormGroup,
) (fieldFormCallback *FieldFormCallback) {
	fieldFormCallback = new(FieldFormCallback)
	fieldFormCallback.probe = probe
	fieldFormCallback.field = field
	fieldFormCallback.formGroup = formGroup

	fieldFormCallback.CreationMode = (field == nil)

	return
}

type FieldFormCallback struct {
	field *models.Field

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (fieldFormCallback *FieldFormCallback) OnSave() {

	// log.Println("FieldFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	fieldFormCallback.probe.formStage.Checkout()

	if fieldFormCallback.field == nil {
		fieldFormCallback.field = new(models.Field).Stage(fieldFormCallback.probe.stageOfInterest)
	}
	field_ := fieldFormCallback.field
	_ = field_

	for _, formDiv := range fieldFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(field_.Name), formDiv)
		case "Identifier":
			FormDivBasicFieldToField(&(field_.Identifier), formDiv)
		case "FieldTypeAsString":
			FormDivBasicFieldToField(&(field_.FieldTypeAsString), formDiv)
		case "Structname":
			FormDivBasicFieldToField(&(field_.Structname), formDiv)
		case "Fieldtypename":
			FormDivBasicFieldToField(&(field_.Fieldtypename), formDiv)
		case "GongStructShape:Fields":
			// WARNING : this form deals with the N-N association "GongStructShape.Fields []*Field" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Field). Setting up a value
			// will discard the former value is there is one.
			//
			// the algorithm is
			// 1/ get the former source of the association
			var formerSource *models.GongStructShape
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "GongStructShape"
				rf.Fieldname = "Fields"
				formerAssociationSource := models.GetReverseFieldOwner(
					fieldFormCallback.probe.stageOfInterest,
					field_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.GongStructShape)
					if !ok {
						log.Fatalln("Source of GongStructShape.Fields []*Field, is not an GongStructShape instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.Fields, field_)
					formerSource.Fields = slices.Delete(formerSource.Fields, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// we need to deal with the 2 cases:
			// 1 the field source is unchanged
			// 2 the field source is changed

			// 1 field source is unchanged
			if formerSource != nil && formerSource.GetName() == newSourceName.GetName() {
				break // nothing else to do for this field
			}

			// 2 field source is changed -->
			// (1) clear the source slice field if it exist
			// (2) find the new source
			// (3) append the new value to the new source field

			// (1) clear the source slice field if it exist
			if formerSource != nil {
				idx := slices.Index(formerSource.Fields, field_)
				formerSource.Fields = slices.Delete(formerSource.Fields, idx, idx+1)
			}

			// (2) find the source
			var newSource *models.GongStructShape
			for _gongstructshape := range *models.GetGongstructInstancesSet[models.GongStructShape](fieldFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _gongstructshape.GetName() == newSourceName.GetName() {
					newSource = _gongstructshape // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of GongStructShape.Fields []*Field, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Fields = append(newSource.Fields, field_)
		}
	}

	// manage the suppress operation
	if fieldFormCallback.formGroup.HasSuppressButtonBeenPressed {
		field_.Unstage(fieldFormCallback.probe.stageOfInterest)
	}

	fieldFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Field](
		fieldFormCallback.probe,
	)
	fieldFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if fieldFormCallback.CreationMode || fieldFormCallback.formGroup.HasSuppressButtonBeenPressed {
		fieldFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(fieldFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FieldFormCallback(
			nil,
			fieldFormCallback.probe,
			newFormGroup,
		)
		field := new(models.Field)
		FillUpForm(field, newFormGroup, fieldFormCallback.probe)
		fieldFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(fieldFormCallback.probe)
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
		case "Position":
			FormDivSelectFieldToField(&(gongenumshape_.Position), gongenumshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Identifier":
			FormDivBasicFieldToField(&(gongenumshape_.Identifier), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(gongenumshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(gongenumshape_.Height), formDiv)
		case "Classdiagram:GongEnumShapes":
			// WARNING : this form deals with the N-N association "Classdiagram.GongEnumShapes []*GongEnumShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of GongEnumShape). Setting up a value
			// will discard the former value is there is one.
			//
			// the algorithm is
			// 1/ get the former source of the association
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
				break // nothing else to do for this field
			}

			// we need to deal with the 2 cases:
			// 1 the field source is unchanged
			// 2 the field source is changed

			// 1 field source is unchanged
			if formerSource != nil && formerSource.GetName() == newSourceName.GetName() {
				break // nothing else to do for this field
			}

			// 2 field source is changed -->
			// (1) clear the source slice field if it exist
			// (2) find the new source
			// (3) append the new value to the new source field

			// (1) clear the source slice field if it exist
			if formerSource != nil {
				idx := slices.Index(formerSource.GongEnumShapes, gongenumshape_)
				formerSource.GongEnumShapes = slices.Delete(formerSource.GongEnumShapes, idx, idx+1)
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

			// (3) append the new value to the new source field
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

	// log.Println("GongEnumValueEntryFormCallback, OnSave")

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
			// WARNING : this form deals with the N-N association "GongEnumShape.GongEnumValueEntrys []*GongEnumValueEntry" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of GongEnumValueEntry). Setting up a value
			// will discard the former value is there is one.
			//
			// the algorithm is
			// 1/ get the former source of the association
			var formerSource *models.GongEnumShape
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "GongEnumShape"
				rf.Fieldname = "GongEnumValueEntrys"
				formerAssociationSource := models.GetReverseFieldOwner(
					gongenumvalueentryFormCallback.probe.stageOfInterest,
					gongenumvalueentry_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.GongEnumShape)
					if !ok {
						log.Fatalln("Source of GongEnumShape.GongEnumValueEntrys []*GongEnumValueEntry, is not an GongEnumShape instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.GongEnumValueEntrys, gongenumvalueentry_)
					formerSource.GongEnumValueEntrys = slices.Delete(formerSource.GongEnumValueEntrys, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// we need to deal with the 2 cases:
			// 1 the field source is unchanged
			// 2 the field source is changed

			// 1 field source is unchanged
			if formerSource != nil && formerSource.GetName() == newSourceName.GetName() {
				break // nothing else to do for this field
			}

			// 2 field source is changed -->
			// (1) clear the source slice field if it exist
			// (2) find the new source
			// (3) append the new value to the new source field

			// (1) clear the source slice field if it exist
			if formerSource != nil {
				idx := slices.Index(formerSource.GongEnumValueEntrys, gongenumvalueentry_)
				formerSource.GongEnumValueEntrys = slices.Delete(formerSource.GongEnumValueEntrys, idx, idx+1)
			}

			// (2) find the source
			var newSource *models.GongEnumShape
			for _gongenumshape := range *models.GetGongstructInstancesSet[models.GongEnumShape](gongenumvalueentryFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _gongenumshape.GetName() == newSourceName.GetName() {
					newSource = _gongenumshape // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of GongEnumShape.GongEnumValueEntrys []*GongEnumValueEntry, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.GongEnumValueEntrys = append(newSource.GongEnumValueEntrys, gongenumvalueentry_)
		}
	}

	// manage the suppress operation
	if gongenumvalueentryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gongenumvalueentry_.Unstage(gongenumvalueentryFormCallback.probe.stageOfInterest)
	}

	gongenumvalueentryFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.GongEnumValueEntry](
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

	updateAndCommitTree(gongenumvalueentryFormCallback.probe)
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
		case "Position":
			FormDivSelectFieldToField(&(gongstructshape_.Position), gongstructshapeFormCallback.probe.stageOfInterest, formDiv)
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
			// the algorithm is
			// 1/ get the former source of the association
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
				break // nothing else to do for this field
			}

			// we need to deal with the 2 cases:
			// 1 the field source is unchanged
			// 2 the field source is changed

			// 1 field source is unchanged
			if formerSource != nil && formerSource.GetName() == newSourceName.GetName() {
				break // nothing else to do for this field
			}

			// 2 field source is changed -->
			// (1) clear the source slice field if it exist
			// (2) find the new source
			// (3) append the new value to the new source field

			// (1) clear the source slice field if it exist
			if formerSource != nil {
				idx := slices.Index(formerSource.GongStructShapes, gongstructshape_)
				formerSource.GongStructShapes = slices.Delete(formerSource.GongStructShapes, idx, idx+1)
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

			// (3) append the new value to the new source field
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
func __gong__New__LinkFormCallback(
	link *models.Link,
	probe *Probe,
	formGroup *table.FormGroup,
) (linkFormCallback *LinkFormCallback) {
	linkFormCallback = new(LinkFormCallback)
	linkFormCallback.probe = probe
	linkFormCallback.link = link
	linkFormCallback.formGroup = formGroup

	linkFormCallback.CreationMode = (link == nil)

	return
}

type LinkFormCallback struct {
	link *models.Link

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (linkFormCallback *LinkFormCallback) OnSave() {

	// log.Println("LinkFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	linkFormCallback.probe.formStage.Checkout()

	if linkFormCallback.link == nil {
		linkFormCallback.link = new(models.Link).Stage(linkFormCallback.probe.stageOfInterest)
	}
	link_ := linkFormCallback.link
	_ = link_

	for _, formDiv := range linkFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(link_.Name), formDiv)
		case "Identifier":
			FormDivBasicFieldToField(&(link_.Identifier), formDiv)
		case "Fieldtypename":
			FormDivBasicFieldToField(&(link_.Fieldtypename), formDiv)
		case "FieldOffsetX":
			FormDivBasicFieldToField(&(link_.FieldOffsetX), formDiv)
		case "FieldOffsetY":
			FormDivBasicFieldToField(&(link_.FieldOffsetY), formDiv)
		case "TargetMultiplicity":
			FormDivEnumStringFieldToField(&(link_.TargetMultiplicity), formDiv)
		case "TargetMultiplicityOffsetX":
			FormDivBasicFieldToField(&(link_.TargetMultiplicityOffsetX), formDiv)
		case "TargetMultiplicityOffsetY":
			FormDivBasicFieldToField(&(link_.TargetMultiplicityOffsetY), formDiv)
		case "SourceMultiplicity":
			FormDivEnumStringFieldToField(&(link_.SourceMultiplicity), formDiv)
		case "SourceMultiplicityOffsetX":
			FormDivBasicFieldToField(&(link_.SourceMultiplicityOffsetX), formDiv)
		case "SourceMultiplicityOffsetY":
			FormDivBasicFieldToField(&(link_.SourceMultiplicityOffsetY), formDiv)
		case "Middlevertice":
			FormDivSelectFieldToField(&(link_.Middlevertice), linkFormCallback.probe.stageOfInterest, formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(link_.StartOrientation), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(link_.StartRatio), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(link_.EndOrientation), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(link_.EndRatio), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(link_.CornerOffsetRatio), formDiv)
		case "GongStructShape:Links":
			// WARNING : this form deals with the N-N association "GongStructShape.Links []*Link" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Link). Setting up a value
			// will discard the former value is there is one.
			//
			// the algorithm is
			// 1/ get the former source of the association
			var formerSource *models.GongStructShape
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "GongStructShape"
				rf.Fieldname = "Links"
				formerAssociationSource := models.GetReverseFieldOwner(
					linkFormCallback.probe.stageOfInterest,
					link_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.GongStructShape)
					if !ok {
						log.Fatalln("Source of GongStructShape.Links []*Link, is not an GongStructShape instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.Links, link_)
					formerSource.Links = slices.Delete(formerSource.Links, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// we need to deal with the 2 cases:
			// 1 the field source is unchanged
			// 2 the field source is changed

			// 1 field source is unchanged
			if formerSource != nil && formerSource.GetName() == newSourceName.GetName() {
				break // nothing else to do for this field
			}

			// 2 field source is changed -->
			// (1) clear the source slice field if it exist
			// (2) find the new source
			// (3) append the new value to the new source field

			// (1) clear the source slice field if it exist
			if formerSource != nil {
				idx := slices.Index(formerSource.Links, link_)
				formerSource.Links = slices.Delete(formerSource.Links, idx, idx+1)
			}

			// (2) find the source
			var newSource *models.GongStructShape
			for _gongstructshape := range *models.GetGongstructInstancesSet[models.GongStructShape](linkFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _gongstructshape.GetName() == newSourceName.GetName() {
					newSource = _gongstructshape // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of GongStructShape.Links []*Link, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Links = append(newSource.Links, link_)
		}
	}

	// manage the suppress operation
	if linkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		link_.Unstage(linkFormCallback.probe.stageOfInterest)
	}

	linkFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Link](
		linkFormCallback.probe,
	)
	linkFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if linkFormCallback.CreationMode || linkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		linkFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(linkFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LinkFormCallback(
			nil,
			linkFormCallback.probe,
			newFormGroup,
		)
		link := new(models.Link)
		FillUpForm(link, newFormGroup, linkFormCallback.probe)
		linkFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(linkFormCallback.probe)
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

	// log.Println("NoteShapeFormCallback, OnSave")

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
		case "Classdiagram:NoteShapes":
			// WARNING : this form deals with the N-N association "Classdiagram.NoteShapes []*NoteShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of NoteShape). Setting up a value
			// will discard the former value is there is one.
			//
			// the algorithm is
			// 1/ get the former source of the association
			var formerSource *models.Classdiagram
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Classdiagram"
				rf.Fieldname = "NoteShapes"
				formerAssociationSource := models.GetReverseFieldOwner(
					noteshapeFormCallback.probe.stageOfInterest,
					noteshape_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Classdiagram)
					if !ok {
						log.Fatalln("Source of Classdiagram.NoteShapes []*NoteShape, is not an Classdiagram instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.NoteShapes, noteshape_)
					formerSource.NoteShapes = slices.Delete(formerSource.NoteShapes, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// we need to deal with the 2 cases:
			// 1 the field source is unchanged
			// 2 the field source is changed

			// 1 field source is unchanged
			if formerSource != nil && formerSource.GetName() == newSourceName.GetName() {
				break // nothing else to do for this field
			}

			// 2 field source is changed -->
			// (1) clear the source slice field if it exist
			// (2) find the new source
			// (3) append the new value to the new source field

			// (1) clear the source slice field if it exist
			if formerSource != nil {
				idx := slices.Index(formerSource.NoteShapes, noteshape_)
				formerSource.NoteShapes = slices.Delete(formerSource.NoteShapes, idx, idx+1)
			}

			// (2) find the source
			var newSource *models.Classdiagram
			for _classdiagram := range *models.GetGongstructInstancesSet[models.Classdiagram](noteshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _classdiagram.GetName() == newSourceName.GetName() {
					newSource = _classdiagram // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Classdiagram.NoteShapes []*NoteShape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.NoteShapes = append(newSource.NoteShapes, noteshape_)
		}
	}

	// manage the suppress operation
	if noteshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteshape_.Unstage(noteshapeFormCallback.probe.stageOfInterest)
	}

	noteshapeFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.NoteShape](
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

	updateAndCommitTree(noteshapeFormCallback.probe)
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

	// log.Println("NoteShapeLinkFormCallback, OnSave")

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
			// WARNING : this form deals with the N-N association "NoteShape.NoteShapeLinks []*NoteShapeLink" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of NoteShapeLink). Setting up a value
			// will discard the former value is there is one.
			//
			// the algorithm is
			// 1/ get the former source of the association
			var formerSource *models.NoteShape
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "NoteShape"
				rf.Fieldname = "NoteShapeLinks"
				formerAssociationSource := models.GetReverseFieldOwner(
					noteshapelinkFormCallback.probe.stageOfInterest,
					noteshapelink_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.NoteShape)
					if !ok {
						log.Fatalln("Source of NoteShape.NoteShapeLinks []*NoteShapeLink, is not an NoteShape instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.NoteShapeLinks, noteshapelink_)
					formerSource.NoteShapeLinks = slices.Delete(formerSource.NoteShapeLinks, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// we need to deal with the 2 cases:
			// 1 the field source is unchanged
			// 2 the field source is changed

			// 1 field source is unchanged
			if formerSource != nil && formerSource.GetName() == newSourceName.GetName() {
				break // nothing else to do for this field
			}

			// 2 field source is changed -->
			// (1) clear the source slice field if it exist
			// (2) find the new source
			// (3) append the new value to the new source field

			// (1) clear the source slice field if it exist
			if formerSource != nil {
				idx := slices.Index(formerSource.NoteShapeLinks, noteshapelink_)
				formerSource.NoteShapeLinks = slices.Delete(formerSource.NoteShapeLinks, idx, idx+1)
			}

			// (2) find the source
			var newSource *models.NoteShape
			for _noteshape := range *models.GetGongstructInstancesSet[models.NoteShape](noteshapelinkFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _noteshape.GetName() == newSourceName.GetName() {
					newSource = _noteshape // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of NoteShape.NoteShapeLinks []*NoteShapeLink, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.NoteShapeLinks = append(newSource.NoteShapeLinks, noteshapelink_)
		}
	}

	// manage the suppress operation
	if noteshapelinkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteshapelink_.Unstage(noteshapelinkFormCallback.probe.stageOfInterest)
	}

	noteshapelinkFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.NoteShapeLink](
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

	updateAndCommitTree(noteshapelinkFormCallback.probe)
}
func __gong__New__PositionFormCallback(
	position *models.Position,
	probe *Probe,
	formGroup *table.FormGroup,
) (positionFormCallback *PositionFormCallback) {
	positionFormCallback = new(PositionFormCallback)
	positionFormCallback.probe = probe
	positionFormCallback.position = position
	positionFormCallback.formGroup = formGroup

	positionFormCallback.CreationMode = (position == nil)

	return
}

type PositionFormCallback struct {
	position *models.Position

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (positionFormCallback *PositionFormCallback) OnSave() {

	// log.Println("PositionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	positionFormCallback.probe.formStage.Checkout()

	if positionFormCallback.position == nil {
		positionFormCallback.position = new(models.Position).Stage(positionFormCallback.probe.stageOfInterest)
	}
	position_ := positionFormCallback.position
	_ = position_

	for _, formDiv := range positionFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "X":
			FormDivBasicFieldToField(&(position_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(position_.Y), formDiv)
		case "Name":
			FormDivBasicFieldToField(&(position_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if positionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		position_.Unstage(positionFormCallback.probe.stageOfInterest)
	}

	positionFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Position](
		positionFormCallback.probe,
	)
	positionFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if positionFormCallback.CreationMode || positionFormCallback.formGroup.HasSuppressButtonBeenPressed {
		positionFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(positionFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PositionFormCallback(
			nil,
			positionFormCallback.probe,
			newFormGroup,
		)
		position := new(models.Position)
		FillUpForm(position, newFormGroup, positionFormCallback.probe)
		positionFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(positionFormCallback.probe)
}
func __gong__New__UmlStateFormCallback(
	umlstate *models.UmlState,
	probe *Probe,
	formGroup *table.FormGroup,
) (umlstateFormCallback *UmlStateFormCallback) {
	umlstateFormCallback = new(UmlStateFormCallback)
	umlstateFormCallback.probe = probe
	umlstateFormCallback.umlstate = umlstate
	umlstateFormCallback.formGroup = formGroup

	umlstateFormCallback.CreationMode = (umlstate == nil)

	return
}

type UmlStateFormCallback struct {
	umlstate *models.UmlState

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (umlstateFormCallback *UmlStateFormCallback) OnSave() {

	// log.Println("UmlStateFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	umlstateFormCallback.probe.formStage.Checkout()

	if umlstateFormCallback.umlstate == nil {
		umlstateFormCallback.umlstate = new(models.UmlState).Stage(umlstateFormCallback.probe.stageOfInterest)
	}
	umlstate_ := umlstateFormCallback.umlstate
	_ = umlstate_

	for _, formDiv := range umlstateFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(umlstate_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(umlstate_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(umlstate_.Y), formDiv)
		case "Umlsc:States":
			// WARNING : this form deals with the N-N association "Umlsc.States []*UmlState" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of UmlState). Setting up a value
			// will discard the former value is there is one.
			//
			// the algorithm is
			// 1/ get the former source of the association
			var formerSource *models.Umlsc
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Umlsc"
				rf.Fieldname = "States"
				formerAssociationSource := models.GetReverseFieldOwner(
					umlstateFormCallback.probe.stageOfInterest,
					umlstate_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Umlsc)
					if !ok {
						log.Fatalln("Source of Umlsc.States []*UmlState, is not an Umlsc instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.States, umlstate_)
					formerSource.States = slices.Delete(formerSource.States, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// we need to deal with the 2 cases:
			// 1 the field source is unchanged
			// 2 the field source is changed

			// 1 field source is unchanged
			if formerSource != nil && formerSource.GetName() == newSourceName.GetName() {
				break // nothing else to do for this field
			}

			// 2 field source is changed -->
			// (1) clear the source slice field if it exist
			// (2) find the new source
			// (3) append the new value to the new source field

			// (1) clear the source slice field if it exist
			if formerSource != nil {
				idx := slices.Index(formerSource.States, umlstate_)
				formerSource.States = slices.Delete(formerSource.States, idx, idx+1)
			}

			// (2) find the source
			var newSource *models.Umlsc
			for _umlsc := range *models.GetGongstructInstancesSet[models.Umlsc](umlstateFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _umlsc.GetName() == newSourceName.GetName() {
					newSource = _umlsc // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Umlsc.States []*UmlState, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.States = append(newSource.States, umlstate_)
		}
	}

	// manage the suppress operation
	if umlstateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		umlstate_.Unstage(umlstateFormCallback.probe.stageOfInterest)
	}

	umlstateFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.UmlState](
		umlstateFormCallback.probe,
	)
	umlstateFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if umlstateFormCallback.CreationMode || umlstateFormCallback.formGroup.HasSuppressButtonBeenPressed {
		umlstateFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(umlstateFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__UmlStateFormCallback(
			nil,
			umlstateFormCallback.probe,
			newFormGroup,
		)
		umlstate := new(models.UmlState)
		FillUpForm(umlstate, newFormGroup, umlstateFormCallback.probe)
		umlstateFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(umlstateFormCallback.probe)
}
func __gong__New__UmlscFormCallback(
	umlsc *models.Umlsc,
	probe *Probe,
	formGroup *table.FormGroup,
) (umlscFormCallback *UmlscFormCallback) {
	umlscFormCallback = new(UmlscFormCallback)
	umlscFormCallback.probe = probe
	umlscFormCallback.umlsc = umlsc
	umlscFormCallback.formGroup = formGroup

	umlscFormCallback.CreationMode = (umlsc == nil)

	return
}

type UmlscFormCallback struct {
	umlsc *models.Umlsc

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (umlscFormCallback *UmlscFormCallback) OnSave() {

	// log.Println("UmlscFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	umlscFormCallback.probe.formStage.Checkout()

	if umlscFormCallback.umlsc == nil {
		umlscFormCallback.umlsc = new(models.Umlsc).Stage(umlscFormCallback.probe.stageOfInterest)
	}
	umlsc_ := umlscFormCallback.umlsc
	_ = umlsc_

	for _, formDiv := range umlscFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(umlsc_.Name), formDiv)
		case "Activestate":
			FormDivBasicFieldToField(&(umlsc_.Activestate), formDiv)
		case "IsInDrawMode":
			FormDivBasicFieldToField(&(umlsc_.IsInDrawMode), formDiv)
		case "DiagramPackage:Umlscs":
			// WARNING : this form deals with the N-N association "DiagramPackage.Umlscs []*Umlsc" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Umlsc). Setting up a value
			// will discard the former value is there is one.
			//
			// the algorithm is
			// 1/ get the former source of the association
			var formerSource *models.DiagramPackage
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "DiagramPackage"
				rf.Fieldname = "Umlscs"
				formerAssociationSource := models.GetReverseFieldOwner(
					umlscFormCallback.probe.stageOfInterest,
					umlsc_,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.DiagramPackage)
					if !ok {
						log.Fatalln("Source of DiagramPackage.Umlscs []*Umlsc, is not an DiagramPackage instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				if formerSource != nil {
					idx := slices.Index(formerSource.Umlscs, umlsc_)
					formerSource.Umlscs = slices.Delete(formerSource.Umlscs, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// we need to deal with the 2 cases:
			// 1 the field source is unchanged
			// 2 the field source is changed

			// 1 field source is unchanged
			if formerSource != nil && formerSource.GetName() == newSourceName.GetName() {
				break // nothing else to do for this field
			}

			// 2 field source is changed -->
			// (1) clear the source slice field if it exist
			// (2) find the new source
			// (3) append the new value to the new source field

			// (1) clear the source slice field if it exist
			if formerSource != nil {
				idx := slices.Index(formerSource.Umlscs, umlsc_)
				formerSource.Umlscs = slices.Delete(formerSource.Umlscs, idx, idx+1)
			}

			// (2) find the source
			var newSource *models.DiagramPackage
			for _diagrampackage := range *models.GetGongstructInstancesSet[models.DiagramPackage](umlscFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagrampackage.GetName() == newSourceName.GetName() {
					newSource = _diagrampackage // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of DiagramPackage.Umlscs []*Umlsc, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Umlscs = append(newSource.Umlscs, umlsc_)
		}
	}

	// manage the suppress operation
	if umlscFormCallback.formGroup.HasSuppressButtonBeenPressed {
		umlsc_.Unstage(umlscFormCallback.probe.stageOfInterest)
	}

	umlscFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Umlsc](
		umlscFormCallback.probe,
	)
	umlscFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if umlscFormCallback.CreationMode || umlscFormCallback.formGroup.HasSuppressButtonBeenPressed {
		umlscFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(umlscFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__UmlscFormCallback(
			nil,
			umlscFormCallback.probe,
			newFormGroup,
		)
		umlsc := new(models.Umlsc)
		FillUpForm(umlsc, newFormGroup, umlscFormCallback.probe)
		umlscFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(umlscFormCallback.probe)
}
func __gong__New__VerticeFormCallback(
	vertice *models.Vertice,
	probe *Probe,
	formGroup *table.FormGroup,
) (verticeFormCallback *VerticeFormCallback) {
	verticeFormCallback = new(VerticeFormCallback)
	verticeFormCallback.probe = probe
	verticeFormCallback.vertice = vertice
	verticeFormCallback.formGroup = formGroup

	verticeFormCallback.CreationMode = (vertice == nil)

	return
}

type VerticeFormCallback struct {
	vertice *models.Vertice

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (verticeFormCallback *VerticeFormCallback) OnSave() {

	// log.Println("VerticeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	verticeFormCallback.probe.formStage.Checkout()

	if verticeFormCallback.vertice == nil {
		verticeFormCallback.vertice = new(models.Vertice).Stage(verticeFormCallback.probe.stageOfInterest)
	}
	vertice_ := verticeFormCallback.vertice
	_ = vertice_

	for _, formDiv := range verticeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "X":
			FormDivBasicFieldToField(&(vertice_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(vertice_.Y), formDiv)
		case "Name":
			FormDivBasicFieldToField(&(vertice_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if verticeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		vertice_.Unstage(verticeFormCallback.probe.stageOfInterest)
	}

	verticeFormCallback.probe.stageOfInterest.Commit()
	updateAndCommitTable[models.Vertice](
		verticeFormCallback.probe,
	)
	verticeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if verticeFormCallback.CreationMode || verticeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		verticeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(verticeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__VerticeFormCallback(
			nil,
			verticeFormCallback.probe,
			newFormGroup,
		)
		vertice := new(models.Vertice)
		FillUpForm(vertice, newFormGroup, verticeFormCallback.probe)
		verticeFormCallback.probe.formStage.Commit()
	}

	updateAndCommitTree(verticeFormCallback.probe)
}
