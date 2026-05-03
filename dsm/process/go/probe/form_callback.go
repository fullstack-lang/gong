// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/process/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__ControlFlowFormCallback(
	controlflow *models.ControlFlow,
	probe *Probe,
	formGroup *form.FormGroup,
) (controlflowFormCallback *ControlFlowFormCallback) {
	controlflowFormCallback = new(ControlFlowFormCallback)
	controlflowFormCallback.probe = probe
	controlflowFormCallback.controlflow = controlflow
	controlflowFormCallback.formGroup = formGroup

	controlflowFormCallback.CreationMode = (controlflow == nil)

	return
}

type ControlFlowFormCallback struct {
	controlflow *models.ControlFlow

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (controlflowFormCallback *ControlFlowFormCallback) OnSave() {
	controlflowFormCallback.probe.stageOfInterest.Lock()
	defer controlflowFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ControlFlowFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	controlflowFormCallback.probe.formStage.Checkout()

	if controlflowFormCallback.controlflow == nil {
		controlflowFormCallback.controlflow = new(models.ControlFlow).Stage(controlflowFormCallback.probe.stageOfInterest)
	}
	controlflow_ := controlflowFormCallback.controlflow
	_ = controlflow_

	for _, formDiv := range controlflowFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(controlflow_.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(controlflow_.ComputedPrefix), formDiv)
		case "Start":
			FormDivSelectFieldToField(&(controlflow_.Start), controlflowFormCallback.probe.stageOfInterest, formDiv)
		case "End":
			FormDivSelectFieldToField(&(controlflow_.End), controlflowFormCallback.probe.stageOfInterest, formDiv)
		case "DiagramProcess:ControlFlowsWhoseNodeIsExpanded":
			// WARNING : this form deals with the N-N association "DiagramProcess.ControlFlowsWhoseNodeIsExpanded []*ControlFlow" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ControlFlow). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.DiagramProcess
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "DiagramProcess"
				rf.Fieldname = "ControlFlowsWhoseNodeIsExpanded"
				formerAssociationSource := controlflow_.GongGetReverseFieldOwner(
					controlflowFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.DiagramProcess)
					if !ok {
						log.Fatalln("Source of DiagramProcess.ControlFlowsWhoseNodeIsExpanded []*ControlFlow, is not an DiagramProcess instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ControlFlowsWhoseNodeIsExpanded, controlflow_)
					formerSource.ControlFlowsWhoseNodeIsExpanded = slices.Delete(formerSource.ControlFlowsWhoseNodeIsExpanded, idx, idx+1)
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
			var newSource *models.DiagramProcess
			for _diagramprocess := range *models.GetGongstructInstancesSet[models.DiagramProcess](controlflowFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagramprocess.GetName() == newSourceName.GetName() {
					newSource = _diagramprocess // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of DiagramProcess.ControlFlowsWhoseNodeIsExpanded []*ControlFlow, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ControlFlowsWhoseNodeIsExpanded = append(newSource.ControlFlowsWhoseNodeIsExpanded, controlflow_)
		case "Participant:ControlFlows":
			// WARNING : this form deals with the N-N association "Participant.ControlFlows []*ControlFlow" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ControlFlow). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Participant
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Participant"
				rf.Fieldname = "ControlFlows"
				formerAssociationSource := controlflow_.GongGetReverseFieldOwner(
					controlflowFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Participant)
					if !ok {
						log.Fatalln("Source of Participant.ControlFlows []*ControlFlow, is not an Participant instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ControlFlows, controlflow_)
					formerSource.ControlFlows = slices.Delete(formerSource.ControlFlows, idx, idx+1)
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
			var newSource *models.Participant
			for _participant := range *models.GetGongstructInstancesSet[models.Participant](controlflowFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _participant.GetName() == newSourceName.GetName() {
					newSource = _participant // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Participant.ControlFlows []*ControlFlow, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ControlFlows = append(newSource.ControlFlows, controlflow_)
		}
	}

	// manage the suppress operation
	if controlflowFormCallback.formGroup.HasSuppressButtonBeenPressed {
		controlflow_.Unstage(controlflowFormCallback.probe.stageOfInterest)
	}

	controlflowFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ControlFlow](
		controlflowFormCallback.probe,
	)

	// display a new form by reset the form stage
	if controlflowFormCallback.CreationMode || controlflowFormCallback.formGroup.HasSuppressButtonBeenPressed {
		controlflowFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(controlflowFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ControlFlowFormCallback(
			nil,
			controlflowFormCallback.probe,
			newFormGroup,
		)
		controlflow := new(models.ControlFlow)
		FillUpForm(controlflow, newFormGroup, controlflowFormCallback.probe)
		controlflowFormCallback.probe.formStage.Commit()
	}

	controlflowFormCallback.probe.ux_tree()
}
func __gong__New__ControlFlowShapeFormCallback(
	controlflowshape *models.ControlFlowShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (controlflowshapeFormCallback *ControlFlowShapeFormCallback) {
	controlflowshapeFormCallback = new(ControlFlowShapeFormCallback)
	controlflowshapeFormCallback.probe = probe
	controlflowshapeFormCallback.controlflowshape = controlflowshape
	controlflowshapeFormCallback.formGroup = formGroup

	controlflowshapeFormCallback.CreationMode = (controlflowshape == nil)

	return
}

type ControlFlowShapeFormCallback struct {
	controlflowshape *models.ControlFlowShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (controlflowshapeFormCallback *ControlFlowShapeFormCallback) OnSave() {
	controlflowshapeFormCallback.probe.stageOfInterest.Lock()
	defer controlflowshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ControlFlowShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	controlflowshapeFormCallback.probe.formStage.Checkout()

	if controlflowshapeFormCallback.controlflowshape == nil {
		controlflowshapeFormCallback.controlflowshape = new(models.ControlFlowShape).Stage(controlflowshapeFormCallback.probe.stageOfInterest)
	}
	controlflowshape_ := controlflowshapeFormCallback.controlflowshape
	_ = controlflowshape_

	for _, formDiv := range controlflowshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(controlflowshape_.Name), formDiv)
		case "ControlFlow":
			FormDivSelectFieldToField(&(controlflowshape_.ControlFlow), controlflowshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(controlflowshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(controlflowshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(controlflowshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(controlflowshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(controlflowshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(controlflowshape_.IsHidden), formDiv)
		case "DiagramProcess:ControlFlow_Shapes":
			// WARNING : this form deals with the N-N association "DiagramProcess.ControlFlow_Shapes []*ControlFlowShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ControlFlowShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.DiagramProcess
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "DiagramProcess"
				rf.Fieldname = "ControlFlow_Shapes"
				formerAssociationSource := controlflowshape_.GongGetReverseFieldOwner(
					controlflowshapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.DiagramProcess)
					if !ok {
						log.Fatalln("Source of DiagramProcess.ControlFlow_Shapes []*ControlFlowShape, is not an DiagramProcess instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ControlFlow_Shapes, controlflowshape_)
					formerSource.ControlFlow_Shapes = slices.Delete(formerSource.ControlFlow_Shapes, idx, idx+1)
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
			var newSource *models.DiagramProcess
			for _diagramprocess := range *models.GetGongstructInstancesSet[models.DiagramProcess](controlflowshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagramprocess.GetName() == newSourceName.GetName() {
					newSource = _diagramprocess // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of DiagramProcess.ControlFlow_Shapes []*ControlFlowShape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ControlFlow_Shapes = append(newSource.ControlFlow_Shapes, controlflowshape_)
		}
	}

	// manage the suppress operation
	if controlflowshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		controlflowshape_.Unstage(controlflowshapeFormCallback.probe.stageOfInterest)
	}

	controlflowshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ControlFlowShape](
		controlflowshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if controlflowshapeFormCallback.CreationMode || controlflowshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		controlflowshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(controlflowshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ControlFlowShapeFormCallback(
			nil,
			controlflowshapeFormCallback.probe,
			newFormGroup,
		)
		controlflowshape := new(models.ControlFlowShape)
		FillUpForm(controlflowshape, newFormGroup, controlflowshapeFormCallback.probe)
		controlflowshapeFormCallback.probe.formStage.Commit()
	}

	controlflowshapeFormCallback.probe.ux_tree()
}
func __gong__New__DataFlowFormCallback(
	dataflow *models.DataFlow,
	probe *Probe,
	formGroup *form.FormGroup,
) (dataflowFormCallback *DataFlowFormCallback) {
	dataflowFormCallback = new(DataFlowFormCallback)
	dataflowFormCallback.probe = probe
	dataflowFormCallback.dataflow = dataflow
	dataflowFormCallback.formGroup = formGroup

	dataflowFormCallback.CreationMode = (dataflow == nil)

	return
}

type DataFlowFormCallback struct {
	dataflow *models.DataFlow

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (dataflowFormCallback *DataFlowFormCallback) OnSave() {
	dataflowFormCallback.probe.stageOfInterest.Lock()
	defer dataflowFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DataFlowFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	dataflowFormCallback.probe.formStage.Checkout()

	if dataflowFormCallback.dataflow == nil {
		dataflowFormCallback.dataflow = new(models.DataFlow).Stage(dataflowFormCallback.probe.stageOfInterest)
	}
	dataflow_ := dataflowFormCallback.dataflow
	_ = dataflow_

	for _, formDiv := range dataflowFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(dataflow_.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(dataflow_.ComputedPrefix), formDiv)
		case "Start":
			FormDivSelectFieldToField(&(dataflow_.Start), dataflowFormCallback.probe.stageOfInterest, formDiv)
		case "End":
			FormDivSelectFieldToField(&(dataflow_.End), dataflowFormCallback.probe.stageOfInterest, formDiv)
		case "DiagramProcess:DataFlowsWhoseNodeIsExpanded":
			// WARNING : this form deals with the N-N association "DiagramProcess.DataFlowsWhoseNodeIsExpanded []*DataFlow" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of DataFlow). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.DiagramProcess
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "DiagramProcess"
				rf.Fieldname = "DataFlowsWhoseNodeIsExpanded"
				formerAssociationSource := dataflow_.GongGetReverseFieldOwner(
					dataflowFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.DiagramProcess)
					if !ok {
						log.Fatalln("Source of DiagramProcess.DataFlowsWhoseNodeIsExpanded []*DataFlow, is not an DiagramProcess instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.DataFlowsWhoseNodeIsExpanded, dataflow_)
					formerSource.DataFlowsWhoseNodeIsExpanded = slices.Delete(formerSource.DataFlowsWhoseNodeIsExpanded, idx, idx+1)
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
			var newSource *models.DiagramProcess
			for _diagramprocess := range *models.GetGongstructInstancesSet[models.DiagramProcess](dataflowFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagramprocess.GetName() == newSourceName.GetName() {
					newSource = _diagramprocess // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of DiagramProcess.DataFlowsWhoseNodeIsExpanded []*DataFlow, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.DataFlowsWhoseNodeIsExpanded = append(newSource.DataFlowsWhoseNodeIsExpanded, dataflow_)
		case "Library:RootDataFlows":
			// WARNING : this form deals with the N-N association "Library.RootDataFlows []*DataFlow" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of DataFlow). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Library
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Library"
				rf.Fieldname = "RootDataFlows"
				formerAssociationSource := dataflow_.GongGetReverseFieldOwner(
					dataflowFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Library)
					if !ok {
						log.Fatalln("Source of Library.RootDataFlows []*DataFlow, is not an Library instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.RootDataFlows, dataflow_)
					formerSource.RootDataFlows = slices.Delete(formerSource.RootDataFlows, idx, idx+1)
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
			var newSource *models.Library
			for _library := range *models.GetGongstructInstancesSet[models.Library](dataflowFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _library.GetName() == newSourceName.GetName() {
					newSource = _library // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Library.RootDataFlows []*DataFlow, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.RootDataFlows = append(newSource.RootDataFlows, dataflow_)
		case "Library:DataFlowsWhoseNodeIsExpanded":
			// WARNING : this form deals with the N-N association "Library.DataFlowsWhoseNodeIsExpanded []*DataFlow" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of DataFlow). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Library
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Library"
				rf.Fieldname = "DataFlowsWhoseNodeIsExpanded"
				formerAssociationSource := dataflow_.GongGetReverseFieldOwner(
					dataflowFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Library)
					if !ok {
						log.Fatalln("Source of Library.DataFlowsWhoseNodeIsExpanded []*DataFlow, is not an Library instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.DataFlowsWhoseNodeIsExpanded, dataflow_)
					formerSource.DataFlowsWhoseNodeIsExpanded = slices.Delete(formerSource.DataFlowsWhoseNodeIsExpanded, idx, idx+1)
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
			var newSource *models.Library
			for _library := range *models.GetGongstructInstancesSet[models.Library](dataflowFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _library.GetName() == newSourceName.GetName() {
					newSource = _library // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Library.DataFlowsWhoseNodeIsExpanded []*DataFlow, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.DataFlowsWhoseNodeIsExpanded = append(newSource.DataFlowsWhoseNodeIsExpanded, dataflow_)
		case "Process:DataFlows":
			// WARNING : this form deals with the N-N association "Process.DataFlows []*DataFlow" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of DataFlow). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Process
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Process"
				rf.Fieldname = "DataFlows"
				formerAssociationSource := dataflow_.GongGetReverseFieldOwner(
					dataflowFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Process)
					if !ok {
						log.Fatalln("Source of Process.DataFlows []*DataFlow, is not an Process instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.DataFlows, dataflow_)
					formerSource.DataFlows = slices.Delete(formerSource.DataFlows, idx, idx+1)
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
			var newSource *models.Process
			for _process := range *models.GetGongstructInstancesSet[models.Process](dataflowFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _process.GetName() == newSourceName.GetName() {
					newSource = _process // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Process.DataFlows []*DataFlow, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.DataFlows = append(newSource.DataFlows, dataflow_)
		}
	}

	// manage the suppress operation
	if dataflowFormCallback.formGroup.HasSuppressButtonBeenPressed {
		dataflow_.Unstage(dataflowFormCallback.probe.stageOfInterest)
	}

	dataflowFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.DataFlow](
		dataflowFormCallback.probe,
	)

	// display a new form by reset the form stage
	if dataflowFormCallback.CreationMode || dataflowFormCallback.formGroup.HasSuppressButtonBeenPressed {
		dataflowFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(dataflowFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DataFlowFormCallback(
			nil,
			dataflowFormCallback.probe,
			newFormGroup,
		)
		dataflow := new(models.DataFlow)
		FillUpForm(dataflow, newFormGroup, dataflowFormCallback.probe)
		dataflowFormCallback.probe.formStage.Commit()
	}

	dataflowFormCallback.probe.ux_tree()
}
func __gong__New__DataFlowShapeFormCallback(
	dataflowshape *models.DataFlowShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (dataflowshapeFormCallback *DataFlowShapeFormCallback) {
	dataflowshapeFormCallback = new(DataFlowShapeFormCallback)
	dataflowshapeFormCallback.probe = probe
	dataflowshapeFormCallback.dataflowshape = dataflowshape
	dataflowshapeFormCallback.formGroup = formGroup

	dataflowshapeFormCallback.CreationMode = (dataflowshape == nil)

	return
}

type DataFlowShapeFormCallback struct {
	dataflowshape *models.DataFlowShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (dataflowshapeFormCallback *DataFlowShapeFormCallback) OnSave() {
	dataflowshapeFormCallback.probe.stageOfInterest.Lock()
	defer dataflowshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DataFlowShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	dataflowshapeFormCallback.probe.formStage.Checkout()

	if dataflowshapeFormCallback.dataflowshape == nil {
		dataflowshapeFormCallback.dataflowshape = new(models.DataFlowShape).Stage(dataflowshapeFormCallback.probe.stageOfInterest)
	}
	dataflowshape_ := dataflowshapeFormCallback.dataflowshape
	_ = dataflowshape_

	for _, formDiv := range dataflowshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(dataflowshape_.Name), formDiv)
		case "DataFlow":
			FormDivSelectFieldToField(&(dataflowshape_.DataFlow), dataflowshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(dataflowshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(dataflowshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(dataflowshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(dataflowshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(dataflowshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(dataflowshape_.IsHidden), formDiv)
		case "DiagramProcess:DataFlow_Shapes":
			// WARNING : this form deals with the N-N association "DiagramProcess.DataFlow_Shapes []*DataFlowShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of DataFlowShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.DiagramProcess
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "DiagramProcess"
				rf.Fieldname = "DataFlow_Shapes"
				formerAssociationSource := dataflowshape_.GongGetReverseFieldOwner(
					dataflowshapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.DiagramProcess)
					if !ok {
						log.Fatalln("Source of DiagramProcess.DataFlow_Shapes []*DataFlowShape, is not an DiagramProcess instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.DataFlow_Shapes, dataflowshape_)
					formerSource.DataFlow_Shapes = slices.Delete(formerSource.DataFlow_Shapes, idx, idx+1)
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
			var newSource *models.DiagramProcess
			for _diagramprocess := range *models.GetGongstructInstancesSet[models.DiagramProcess](dataflowshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagramprocess.GetName() == newSourceName.GetName() {
					newSource = _diagramprocess // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of DiagramProcess.DataFlow_Shapes []*DataFlowShape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.DataFlow_Shapes = append(newSource.DataFlow_Shapes, dataflowshape_)
		}
	}

	// manage the suppress operation
	if dataflowshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		dataflowshape_.Unstage(dataflowshapeFormCallback.probe.stageOfInterest)
	}

	dataflowshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.DataFlowShape](
		dataflowshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if dataflowshapeFormCallback.CreationMode || dataflowshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		dataflowshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(dataflowshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DataFlowShapeFormCallback(
			nil,
			dataflowshapeFormCallback.probe,
			newFormGroup,
		)
		dataflowshape := new(models.DataFlowShape)
		FillUpForm(dataflowshape, newFormGroup, dataflowshapeFormCallback.probe)
		dataflowshapeFormCallback.probe.formStage.Commit()
	}

	dataflowshapeFormCallback.probe.ux_tree()
}
func __gong__New__DiagramProcessFormCallback(
	diagramprocess *models.DiagramProcess,
	probe *Probe,
	formGroup *form.FormGroup,
) (diagramprocessFormCallback *DiagramProcessFormCallback) {
	diagramprocessFormCallback = new(DiagramProcessFormCallback)
	diagramprocessFormCallback.probe = probe
	diagramprocessFormCallback.diagramprocess = diagramprocess
	diagramprocessFormCallback.formGroup = formGroup

	diagramprocessFormCallback.CreationMode = (diagramprocess == nil)

	return
}

type DiagramProcessFormCallback struct {
	diagramprocess *models.DiagramProcess

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (diagramprocessFormCallback *DiagramProcessFormCallback) OnSave() {
	diagramprocessFormCallback.probe.stageOfInterest.Lock()
	defer diagramprocessFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DiagramProcessFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	diagramprocessFormCallback.probe.formStage.Checkout()

	if diagramprocessFormCallback.diagramprocess == nil {
		diagramprocessFormCallback.diagramprocess = new(models.DiagramProcess).Stage(diagramprocessFormCallback.probe.stageOfInterest)
	}
	diagramprocess_ := diagramprocessFormCallback.diagramprocess
	_ = diagramprocess_

	for _, formDiv := range diagramprocessFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(diagramprocess_.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(diagramprocess_.ComputedPrefix), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(diagramprocess_.IsChecked), formDiv)
		case "IsEditable_":
			FormDivBasicFieldToField(&(diagramprocess_.IsEditable_), formDiv)
		case "IsShowPrefix":
			FormDivBasicFieldToField(&(diagramprocess_.IsShowPrefix), formDiv)
		case "DefaultBoxWidth":
			FormDivBasicFieldToField(&(diagramprocess_.DefaultBoxWidth), formDiv)
		case "DefaultBoxHeigth":
			FormDivBasicFieldToField(&(diagramprocess_.DefaultBoxHeigth), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(diagramprocess_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(diagramprocess_.Height), formDiv)
		case "Process_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ProcessShape](diagramprocessFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ProcessShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ProcessShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramprocessFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ProcessShape](diagramprocessFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramprocess_.Process_Shapes = instanceSlice

		case "IsProcesssNodeExpanded":
			FormDivBasicFieldToField(&(diagramprocess_.IsProcesssNodeExpanded), formDiv)
		case "ProcesssWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Process](diagramprocessFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Process, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Process)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramprocessFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Process](diagramprocessFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramprocess_.ProcesssWhoseNodeIsExpanded = instanceSlice

		case "Participant_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ParticipantShape](diagramprocessFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ParticipantShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ParticipantShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramprocessFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ParticipantShape](diagramprocessFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramprocess_.Participant_Shapes = instanceSlice

		case "IsParticipantsNodeExpanded":
			FormDivBasicFieldToField(&(diagramprocess_.IsParticipantsNodeExpanded), formDiv)
		case "ParticipantWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Participant](diagramprocessFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Participant, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Participant)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramprocessFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Participant](diagramprocessFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramprocess_.ParticipantWhoseNodeIsExpanded = instanceSlice

		case "TasksWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Task](diagramprocessFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Task, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Task)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramprocessFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Task](diagramprocessFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramprocess_.TasksWhoseNodeIsExpanded = instanceSlice

		case "Task_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.TaskShape](diagramprocessFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.TaskShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.TaskShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramprocessFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.TaskShape](diagramprocessFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramprocess_.Task_Shapes = instanceSlice

		case "ControlFlowsWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ControlFlow](diagramprocessFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ControlFlow, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ControlFlow)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramprocessFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ControlFlow](diagramprocessFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramprocess_.ControlFlowsWhoseNodeIsExpanded = instanceSlice

		case "ControlFlow_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ControlFlowShape](diagramprocessFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ControlFlowShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ControlFlowShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramprocessFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ControlFlowShape](diagramprocessFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramprocess_.ControlFlow_Shapes = instanceSlice

		case "DataFlowsWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DataFlow](diagramprocessFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DataFlow, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DataFlow)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramprocessFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.DataFlow](diagramprocessFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramprocess_.DataFlowsWhoseNodeIsExpanded = instanceSlice

		case "DataFlow_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DataFlowShape](diagramprocessFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DataFlowShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DataFlowShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramprocessFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.DataFlowShape](diagramprocessFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramprocess_.DataFlow_Shapes = instanceSlice

		case "Process:DiagramProcesss":
			// WARNING : this form deals with the N-N association "Process.DiagramProcesss []*DiagramProcess" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of DiagramProcess). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Process
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Process"
				rf.Fieldname = "DiagramProcesss"
				formerAssociationSource := diagramprocess_.GongGetReverseFieldOwner(
					diagramprocessFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Process)
					if !ok {
						log.Fatalln("Source of Process.DiagramProcesss []*DiagramProcess, is not an Process instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.DiagramProcesss, diagramprocess_)
					formerSource.DiagramProcesss = slices.Delete(formerSource.DiagramProcesss, idx, idx+1)
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
			var newSource *models.Process
			for _process := range *models.GetGongstructInstancesSet[models.Process](diagramprocessFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _process.GetName() == newSourceName.GetName() {
					newSource = _process // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Process.DiagramProcesss []*DiagramProcess, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.DiagramProcesss = append(newSource.DiagramProcesss, diagramprocess_)
		case "Process:DiagramProcessWhoseNodeIsExpanded":
			// WARNING : this form deals with the N-N association "Process.DiagramProcessWhoseNodeIsExpanded []*DiagramProcess" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of DiagramProcess). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Process
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Process"
				rf.Fieldname = "DiagramProcessWhoseNodeIsExpanded"
				formerAssociationSource := diagramprocess_.GongGetReverseFieldOwner(
					diagramprocessFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Process)
					if !ok {
						log.Fatalln("Source of Process.DiagramProcessWhoseNodeIsExpanded []*DiagramProcess, is not an Process instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.DiagramProcessWhoseNodeIsExpanded, diagramprocess_)
					formerSource.DiagramProcessWhoseNodeIsExpanded = slices.Delete(formerSource.DiagramProcessWhoseNodeIsExpanded, idx, idx+1)
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
			var newSource *models.Process
			for _process := range *models.GetGongstructInstancesSet[models.Process](diagramprocessFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _process.GetName() == newSourceName.GetName() {
					newSource = _process // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Process.DiagramProcessWhoseNodeIsExpanded []*DiagramProcess, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.DiagramProcessWhoseNodeIsExpanded = append(newSource.DiagramProcessWhoseNodeIsExpanded, diagramprocess_)
		}
	}

	// manage the suppress operation
	if diagramprocessFormCallback.formGroup.HasSuppressButtonBeenPressed {
		diagramprocess_.Unstage(diagramprocessFormCallback.probe.stageOfInterest)
	}

	diagramprocessFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.DiagramProcess](
		diagramprocessFormCallback.probe,
	)

	// display a new form by reset the form stage
	if diagramprocessFormCallback.CreationMode || diagramprocessFormCallback.formGroup.HasSuppressButtonBeenPressed {
		diagramprocessFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(diagramprocessFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DiagramProcessFormCallback(
			nil,
			diagramprocessFormCallback.probe,
			newFormGroup,
		)
		diagramprocess := new(models.DiagramProcess)
		FillUpForm(diagramprocess, newFormGroup, diagramprocessFormCallback.probe)
		diagramprocessFormCallback.probe.formStage.Commit()
	}

	diagramprocessFormCallback.probe.ux_tree()
}
func __gong__New__LibraryFormCallback(
	library *models.Library,
	probe *Probe,
	formGroup *form.FormGroup,
) (libraryFormCallback *LibraryFormCallback) {
	libraryFormCallback = new(LibraryFormCallback)
	libraryFormCallback.probe = probe
	libraryFormCallback.library = library
	libraryFormCallback.formGroup = formGroup

	libraryFormCallback.CreationMode = (library == nil)

	return
}

type LibraryFormCallback struct {
	library *models.Library

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (libraryFormCallback *LibraryFormCallback) OnSave() {
	libraryFormCallback.probe.stageOfInterest.Lock()
	defer libraryFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("LibraryFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	libraryFormCallback.probe.formStage.Checkout()

	if libraryFormCallback.library == nil {
		libraryFormCallback.library = new(models.Library).Stage(libraryFormCallback.probe.stageOfInterest)
	}
	library_ := libraryFormCallback.library
	_ = library_

	for _, formDiv := range libraryFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(library_.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(library_.ComputedPrefix), formDiv)
		case "SubLibraries":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Library](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Library, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Library)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					libraryFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Library](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.SubLibraries = instanceSlice

		case "IsSubLibrariesNodeExpanded":
			FormDivBasicFieldToField(&(library_.IsSubLibrariesNodeExpanded), formDiv)
		case "SubLibrariesWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Library](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Library, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Library)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					libraryFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Library](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.SubLibrariesWhoseNodeIsExpanded = instanceSlice

		case "NbPixPerCharacter":
			FormDivBasicFieldToField(&(library_.NbPixPerCharacter), formDiv)
		case "LogoSVGFile":
			FormDivBasicFieldToField(&(library_.LogoSVGFile), formDiv)
		case "RootProcesses":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Process](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Process, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Process)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					libraryFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Process](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.RootProcesses = instanceSlice

		case "IsProcessesNodeExpanded":
			FormDivBasicFieldToField(&(library_.IsProcessesNodeExpanded), formDiv)
		case "ProcesssWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Process](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Process, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Process)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					libraryFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Process](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.ProcesssWhoseNodeIsExpanded = instanceSlice

		case "RootDataFlows":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DataFlow](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DataFlow, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DataFlow)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					libraryFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.DataFlow](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.RootDataFlows = instanceSlice

		case "IsDataFlowsNodeExpanded":
			FormDivBasicFieldToField(&(library_.IsDataFlowsNodeExpanded), formDiv)
		case "DataFlowsWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DataFlow](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DataFlow, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DataFlow)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					libraryFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.DataFlow](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.DataFlowsWhoseNodeIsExpanded = instanceSlice

		case "IsExpandedTmp":
			FormDivBasicFieldToField(&(library_.IsExpandedTmp), formDiv)
		case "Library:SubLibraries":
			// WARNING : this form deals with the N-N association "Library.SubLibraries []*Library" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Library). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Library
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Library"
				rf.Fieldname = "SubLibraries"
				formerAssociationSource := library_.GongGetReverseFieldOwner(
					libraryFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Library)
					if !ok {
						log.Fatalln("Source of Library.SubLibraries []*Library, is not an Library instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.SubLibraries, library_)
					formerSource.SubLibraries = slices.Delete(formerSource.SubLibraries, idx, idx+1)
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
			var newSource *models.Library
			for _library := range *models.GetGongstructInstancesSet[models.Library](libraryFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _library.GetName() == newSourceName.GetName() {
					newSource = _library // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Library.SubLibraries []*Library, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.SubLibraries = append(newSource.SubLibraries, library_)
		case "Library:SubLibrariesWhoseNodeIsExpanded":
			// WARNING : this form deals with the N-N association "Library.SubLibrariesWhoseNodeIsExpanded []*Library" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Library). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Library
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Library"
				rf.Fieldname = "SubLibrariesWhoseNodeIsExpanded"
				formerAssociationSource := library_.GongGetReverseFieldOwner(
					libraryFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Library)
					if !ok {
						log.Fatalln("Source of Library.SubLibrariesWhoseNodeIsExpanded []*Library, is not an Library instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.SubLibrariesWhoseNodeIsExpanded, library_)
					formerSource.SubLibrariesWhoseNodeIsExpanded = slices.Delete(formerSource.SubLibrariesWhoseNodeIsExpanded, idx, idx+1)
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
			var newSource *models.Library
			for _library := range *models.GetGongstructInstancesSet[models.Library](libraryFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _library.GetName() == newSourceName.GetName() {
					newSource = _library // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Library.SubLibrariesWhoseNodeIsExpanded []*Library, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.SubLibrariesWhoseNodeIsExpanded = append(newSource.SubLibrariesWhoseNodeIsExpanded, library_)
		}
	}

	// manage the suppress operation
	if libraryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		library_.Unstage(libraryFormCallback.probe.stageOfInterest)
	}

	libraryFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Library](
		libraryFormCallback.probe,
	)

	// display a new form by reset the form stage
	if libraryFormCallback.CreationMode || libraryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		libraryFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(libraryFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LibraryFormCallback(
			nil,
			libraryFormCallback.probe,
			newFormGroup,
		)
		library := new(models.Library)
		FillUpForm(library, newFormGroup, libraryFormCallback.probe)
		libraryFormCallback.probe.formStage.Commit()
	}

	libraryFormCallback.probe.ux_tree()
}
func __gong__New__ParticipantFormCallback(
	participant *models.Participant,
	probe *Probe,
	formGroup *form.FormGroup,
) (participantFormCallback *ParticipantFormCallback) {
	participantFormCallback = new(ParticipantFormCallback)
	participantFormCallback.probe = probe
	participantFormCallback.participant = participant
	participantFormCallback.formGroup = formGroup

	participantFormCallback.CreationMode = (participant == nil)

	return
}

type ParticipantFormCallback struct {
	participant *models.Participant

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (participantFormCallback *ParticipantFormCallback) OnSave() {
	participantFormCallback.probe.stageOfInterest.Lock()
	defer participantFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ParticipantFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	participantFormCallback.probe.formStage.Checkout()

	if participantFormCallback.participant == nil {
		participantFormCallback.participant = new(models.Participant).Stage(participantFormCallback.probe.stageOfInterest)
	}
	participant_ := participantFormCallback.participant
	_ = participant_

	for _, formDiv := range participantFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(participant_.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(participant_.ComputedPrefix), formDiv)
		case "IsTasksNodeExpanded":
			FormDivBasicFieldToField(&(participant_.IsTasksNodeExpanded), formDiv)
		case "Tasks":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Task](participantFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Task, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Task)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					participantFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Task](participantFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			participant_.Tasks = instanceSlice

		case "IsControlFlowsNodeExpanded":
			FormDivBasicFieldToField(&(participant_.IsControlFlowsNodeExpanded), formDiv)
		case "ControlFlows":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ControlFlow](participantFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ControlFlow, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ControlFlow)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					participantFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ControlFlow](participantFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			participant_.ControlFlows = instanceSlice

		case "TaskWhoseOutControlFlowsNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Task](participantFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Task, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Task)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					participantFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Task](participantFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			participant_.TaskWhoseOutControlFlowsNodeIsExpanded = instanceSlice

		case "TaskWhoseInControlFlowsNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Task](participantFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Task, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Task)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					participantFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Task](participantFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			participant_.TaskWhoseInControlFlowsNodeIsExpanded = instanceSlice

		case "DiagramProcess:ParticipantWhoseNodeIsExpanded":
			// WARNING : this form deals with the N-N association "DiagramProcess.ParticipantWhoseNodeIsExpanded []*Participant" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Participant). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.DiagramProcess
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "DiagramProcess"
				rf.Fieldname = "ParticipantWhoseNodeIsExpanded"
				formerAssociationSource := participant_.GongGetReverseFieldOwner(
					participantFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.DiagramProcess)
					if !ok {
						log.Fatalln("Source of DiagramProcess.ParticipantWhoseNodeIsExpanded []*Participant, is not an DiagramProcess instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ParticipantWhoseNodeIsExpanded, participant_)
					formerSource.ParticipantWhoseNodeIsExpanded = slices.Delete(formerSource.ParticipantWhoseNodeIsExpanded, idx, idx+1)
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
			var newSource *models.DiagramProcess
			for _diagramprocess := range *models.GetGongstructInstancesSet[models.DiagramProcess](participantFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagramprocess.GetName() == newSourceName.GetName() {
					newSource = _diagramprocess // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of DiagramProcess.ParticipantWhoseNodeIsExpanded []*Participant, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ParticipantWhoseNodeIsExpanded = append(newSource.ParticipantWhoseNodeIsExpanded, participant_)
		case "Process:Participants":
			// WARNING : this form deals with the N-N association "Process.Participants []*Participant" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Participant). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Process
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Process"
				rf.Fieldname = "Participants"
				formerAssociationSource := participant_.GongGetReverseFieldOwner(
					participantFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Process)
					if !ok {
						log.Fatalln("Source of Process.Participants []*Participant, is not an Process instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Participants, participant_)
					formerSource.Participants = slices.Delete(formerSource.Participants, idx, idx+1)
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
			var newSource *models.Process
			for _process := range *models.GetGongstructInstancesSet[models.Process](participantFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _process.GetName() == newSourceName.GetName() {
					newSource = _process // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Process.Participants []*Participant, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Participants = append(newSource.Participants, participant_)
		case "Process:ParticipantWhoseNodeIsExpanded":
			// WARNING : this form deals with the N-N association "Process.ParticipantWhoseNodeIsExpanded []*Participant" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Participant). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Process
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Process"
				rf.Fieldname = "ParticipantWhoseNodeIsExpanded"
				formerAssociationSource := participant_.GongGetReverseFieldOwner(
					participantFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Process)
					if !ok {
						log.Fatalln("Source of Process.ParticipantWhoseNodeIsExpanded []*Participant, is not an Process instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ParticipantWhoseNodeIsExpanded, participant_)
					formerSource.ParticipantWhoseNodeIsExpanded = slices.Delete(formerSource.ParticipantWhoseNodeIsExpanded, idx, idx+1)
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
			var newSource *models.Process
			for _process := range *models.GetGongstructInstancesSet[models.Process](participantFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _process.GetName() == newSourceName.GetName() {
					newSource = _process // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Process.ParticipantWhoseNodeIsExpanded []*Participant, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ParticipantWhoseNodeIsExpanded = append(newSource.ParticipantWhoseNodeIsExpanded, participant_)
		}
	}

	// manage the suppress operation
	if participantFormCallback.formGroup.HasSuppressButtonBeenPressed {
		participant_.Unstage(participantFormCallback.probe.stageOfInterest)
	}

	participantFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Participant](
		participantFormCallback.probe,
	)

	// display a new form by reset the form stage
	if participantFormCallback.CreationMode || participantFormCallback.formGroup.HasSuppressButtonBeenPressed {
		participantFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(participantFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ParticipantFormCallback(
			nil,
			participantFormCallback.probe,
			newFormGroup,
		)
		participant := new(models.Participant)
		FillUpForm(participant, newFormGroup, participantFormCallback.probe)
		participantFormCallback.probe.formStage.Commit()
	}

	participantFormCallback.probe.ux_tree()
}
func __gong__New__ParticipantShapeFormCallback(
	participantshape *models.ParticipantShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (participantshapeFormCallback *ParticipantShapeFormCallback) {
	participantshapeFormCallback = new(ParticipantShapeFormCallback)
	participantshapeFormCallback.probe = probe
	participantshapeFormCallback.participantshape = participantshape
	participantshapeFormCallback.formGroup = formGroup

	participantshapeFormCallback.CreationMode = (participantshape == nil)

	return
}

type ParticipantShapeFormCallback struct {
	participantshape *models.ParticipantShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (participantshapeFormCallback *ParticipantShapeFormCallback) OnSave() {
	participantshapeFormCallback.probe.stageOfInterest.Lock()
	defer participantshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ParticipantShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	participantshapeFormCallback.probe.formStage.Checkout()

	if participantshapeFormCallback.participantshape == nil {
		participantshapeFormCallback.participantshape = new(models.ParticipantShape).Stage(participantshapeFormCallback.probe.stageOfInterest)
	}
	participantshape_ := participantshapeFormCallback.participantshape
	_ = participantshape_

	for _, formDiv := range participantshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(participantshape_.Name), formDiv)
		case "Participant":
			FormDivSelectFieldToField(&(participantshape_.Participant), participantshapeFormCallback.probe.stageOfInterest, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(participantshape_.IsExpanded), formDiv)
		case "X":
			FormDivBasicFieldToField(&(participantshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(participantshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(participantshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(participantshape_.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(participantshape_.IsHidden), formDiv)
		case "DiagramProcess:Participant_Shapes":
			// WARNING : this form deals with the N-N association "DiagramProcess.Participant_Shapes []*ParticipantShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ParticipantShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.DiagramProcess
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "DiagramProcess"
				rf.Fieldname = "Participant_Shapes"
				formerAssociationSource := participantshape_.GongGetReverseFieldOwner(
					participantshapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.DiagramProcess)
					if !ok {
						log.Fatalln("Source of DiagramProcess.Participant_Shapes []*ParticipantShape, is not an DiagramProcess instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Participant_Shapes, participantshape_)
					formerSource.Participant_Shapes = slices.Delete(formerSource.Participant_Shapes, idx, idx+1)
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
			var newSource *models.DiagramProcess
			for _diagramprocess := range *models.GetGongstructInstancesSet[models.DiagramProcess](participantshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagramprocess.GetName() == newSourceName.GetName() {
					newSource = _diagramprocess // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of DiagramProcess.Participant_Shapes []*ParticipantShape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Participant_Shapes = append(newSource.Participant_Shapes, participantshape_)
		}
	}

	// manage the suppress operation
	if participantshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		participantshape_.Unstage(participantshapeFormCallback.probe.stageOfInterest)
	}

	participantshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ParticipantShape](
		participantshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if participantshapeFormCallback.CreationMode || participantshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		participantshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(participantshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ParticipantShapeFormCallback(
			nil,
			participantshapeFormCallback.probe,
			newFormGroup,
		)
		participantshape := new(models.ParticipantShape)
		FillUpForm(participantshape, newFormGroup, participantshapeFormCallback.probe)
		participantshapeFormCallback.probe.formStage.Commit()
	}

	participantshapeFormCallback.probe.ux_tree()
}
func __gong__New__ProcessFormCallback(
	process *models.Process,
	probe *Probe,
	formGroup *form.FormGroup,
) (processFormCallback *ProcessFormCallback) {
	processFormCallback = new(ProcessFormCallback)
	processFormCallback.probe = probe
	processFormCallback.process = process
	processFormCallback.formGroup = formGroup

	processFormCallback.CreationMode = (process == nil)

	return
}

type ProcessFormCallback struct {
	process *models.Process

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (processFormCallback *ProcessFormCallback) OnSave() {
	processFormCallback.probe.stageOfInterest.Lock()
	defer processFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ProcessFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	processFormCallback.probe.formStage.Checkout()

	if processFormCallback.process == nil {
		processFormCallback.process = new(models.Process).Stage(processFormCallback.probe.stageOfInterest)
	}
	process_ := processFormCallback.process
	_ = process_

	for _, formDiv := range processFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(process_.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(process_.ComputedPrefix), formDiv)
		case "DiagramProcesss":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DiagramProcess](processFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DiagramProcess, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DiagramProcess)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					processFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramProcess](processFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			process_.DiagramProcesss = instanceSlice

		case "DiagramProcessWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DiagramProcess](processFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DiagramProcess, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DiagramProcess)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					processFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramProcess](processFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			process_.DiagramProcessWhoseNodeIsExpanded = instanceSlice

		case "IsSubProcessNodeExpanded":
			FormDivBasicFieldToField(&(process_.IsSubProcessNodeExpanded), formDiv)
		case "SubProcesses":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Process](processFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Process, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Process)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					processFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Process](processFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			process_.SubProcesses = instanceSlice

		case "Participants":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Participant](processFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Participant, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Participant)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					processFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Participant](processFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			process_.Participants = instanceSlice

		case "ParticipantWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Participant](processFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Participant, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Participant)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					processFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Participant](processFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			process_.ParticipantWhoseNodeIsExpanded = instanceSlice

		case "DataFlows":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DataFlow](processFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DataFlow, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DataFlow)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					processFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.DataFlow](processFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			process_.DataFlows = instanceSlice

		case "IsDataFlowsNodeExpanded":
			FormDivBasicFieldToField(&(process_.IsDataFlowsNodeExpanded), formDiv)
		case "DiagramProcess:ProcesssWhoseNodeIsExpanded":
			// WARNING : this form deals with the N-N association "DiagramProcess.ProcesssWhoseNodeIsExpanded []*Process" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Process). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.DiagramProcess
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "DiagramProcess"
				rf.Fieldname = "ProcesssWhoseNodeIsExpanded"
				formerAssociationSource := process_.GongGetReverseFieldOwner(
					processFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.DiagramProcess)
					if !ok {
						log.Fatalln("Source of DiagramProcess.ProcesssWhoseNodeIsExpanded []*Process, is not an DiagramProcess instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ProcesssWhoseNodeIsExpanded, process_)
					formerSource.ProcesssWhoseNodeIsExpanded = slices.Delete(formerSource.ProcesssWhoseNodeIsExpanded, idx, idx+1)
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
			var newSource *models.DiagramProcess
			for _diagramprocess := range *models.GetGongstructInstancesSet[models.DiagramProcess](processFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagramprocess.GetName() == newSourceName.GetName() {
					newSource = _diagramprocess // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of DiagramProcess.ProcesssWhoseNodeIsExpanded []*Process, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ProcesssWhoseNodeIsExpanded = append(newSource.ProcesssWhoseNodeIsExpanded, process_)
		case "Library:RootProcesses":
			// WARNING : this form deals with the N-N association "Library.RootProcesses []*Process" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Process). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Library
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Library"
				rf.Fieldname = "RootProcesses"
				formerAssociationSource := process_.GongGetReverseFieldOwner(
					processFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Library)
					if !ok {
						log.Fatalln("Source of Library.RootProcesses []*Process, is not an Library instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.RootProcesses, process_)
					formerSource.RootProcesses = slices.Delete(formerSource.RootProcesses, idx, idx+1)
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
			var newSource *models.Library
			for _library := range *models.GetGongstructInstancesSet[models.Library](processFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _library.GetName() == newSourceName.GetName() {
					newSource = _library // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Library.RootProcesses []*Process, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.RootProcesses = append(newSource.RootProcesses, process_)
		case "Library:ProcesssWhoseNodeIsExpanded":
			// WARNING : this form deals with the N-N association "Library.ProcesssWhoseNodeIsExpanded []*Process" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Process). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Library
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Library"
				rf.Fieldname = "ProcesssWhoseNodeIsExpanded"
				formerAssociationSource := process_.GongGetReverseFieldOwner(
					processFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Library)
					if !ok {
						log.Fatalln("Source of Library.ProcesssWhoseNodeIsExpanded []*Process, is not an Library instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.ProcesssWhoseNodeIsExpanded, process_)
					formerSource.ProcesssWhoseNodeIsExpanded = slices.Delete(formerSource.ProcesssWhoseNodeIsExpanded, idx, idx+1)
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
			var newSource *models.Library
			for _library := range *models.GetGongstructInstancesSet[models.Library](processFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _library.GetName() == newSourceName.GetName() {
					newSource = _library // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Library.ProcesssWhoseNodeIsExpanded []*Process, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.ProcesssWhoseNodeIsExpanded = append(newSource.ProcesssWhoseNodeIsExpanded, process_)
		case "Process:SubProcesses":
			// WARNING : this form deals with the N-N association "Process.SubProcesses []*Process" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Process). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Process
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Process"
				rf.Fieldname = "SubProcesses"
				formerAssociationSource := process_.GongGetReverseFieldOwner(
					processFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Process)
					if !ok {
						log.Fatalln("Source of Process.SubProcesses []*Process, is not an Process instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.SubProcesses, process_)
					formerSource.SubProcesses = slices.Delete(formerSource.SubProcesses, idx, idx+1)
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
			var newSource *models.Process
			for _process := range *models.GetGongstructInstancesSet[models.Process](processFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _process.GetName() == newSourceName.GetName() {
					newSource = _process // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Process.SubProcesses []*Process, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.SubProcesses = append(newSource.SubProcesses, process_)
		}
	}

	// manage the suppress operation
	if processFormCallback.formGroup.HasSuppressButtonBeenPressed {
		process_.Unstage(processFormCallback.probe.stageOfInterest)
	}

	processFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Process](
		processFormCallback.probe,
	)

	// display a new form by reset the form stage
	if processFormCallback.CreationMode || processFormCallback.formGroup.HasSuppressButtonBeenPressed {
		processFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(processFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ProcessFormCallback(
			nil,
			processFormCallback.probe,
			newFormGroup,
		)
		process := new(models.Process)
		FillUpForm(process, newFormGroup, processFormCallback.probe)
		processFormCallback.probe.formStage.Commit()
	}

	processFormCallback.probe.ux_tree()
}
func __gong__New__ProcessShapeFormCallback(
	processshape *models.ProcessShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (processshapeFormCallback *ProcessShapeFormCallback) {
	processshapeFormCallback = new(ProcessShapeFormCallback)
	processshapeFormCallback.probe = probe
	processshapeFormCallback.processshape = processshape
	processshapeFormCallback.formGroup = formGroup

	processshapeFormCallback.CreationMode = (processshape == nil)

	return
}

type ProcessShapeFormCallback struct {
	processshape *models.ProcessShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (processshapeFormCallback *ProcessShapeFormCallback) OnSave() {
	processshapeFormCallback.probe.stageOfInterest.Lock()
	defer processshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ProcessShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	processshapeFormCallback.probe.formStage.Checkout()

	if processshapeFormCallback.processshape == nil {
		processshapeFormCallback.processshape = new(models.ProcessShape).Stage(processshapeFormCallback.probe.stageOfInterest)
	}
	processshape_ := processshapeFormCallback.processshape
	_ = processshape_

	for _, formDiv := range processshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(processshape_.Name), formDiv)
		case "Process":
			FormDivSelectFieldToField(&(processshape_.Process), processshapeFormCallback.probe.stageOfInterest, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(processshape_.IsExpanded), formDiv)
		case "X":
			FormDivBasicFieldToField(&(processshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(processshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(processshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(processshape_.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(processshape_.IsHidden), formDiv)
		case "DiagramProcess:Process_Shapes":
			// WARNING : this form deals with the N-N association "DiagramProcess.Process_Shapes []*ProcessShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of ProcessShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.DiagramProcess
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "DiagramProcess"
				rf.Fieldname = "Process_Shapes"
				formerAssociationSource := processshape_.GongGetReverseFieldOwner(
					processshapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.DiagramProcess)
					if !ok {
						log.Fatalln("Source of DiagramProcess.Process_Shapes []*ProcessShape, is not an DiagramProcess instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Process_Shapes, processshape_)
					formerSource.Process_Shapes = slices.Delete(formerSource.Process_Shapes, idx, idx+1)
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
			var newSource *models.DiagramProcess
			for _diagramprocess := range *models.GetGongstructInstancesSet[models.DiagramProcess](processshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagramprocess.GetName() == newSourceName.GetName() {
					newSource = _diagramprocess // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of DiagramProcess.Process_Shapes []*ProcessShape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Process_Shapes = append(newSource.Process_Shapes, processshape_)
		}
	}

	// manage the suppress operation
	if processshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		processshape_.Unstage(processshapeFormCallback.probe.stageOfInterest)
	}

	processshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ProcessShape](
		processshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if processshapeFormCallback.CreationMode || processshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		processshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(processshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ProcessShapeFormCallback(
			nil,
			processshapeFormCallback.probe,
			newFormGroup,
		)
		processshape := new(models.ProcessShape)
		FillUpForm(processshape, newFormGroup, processshapeFormCallback.probe)
		processshapeFormCallback.probe.formStage.Commit()
	}

	processshapeFormCallback.probe.ux_tree()
}
func __gong__New__TaskFormCallback(
	task *models.Task,
	probe *Probe,
	formGroup *form.FormGroup,
) (taskFormCallback *TaskFormCallback) {
	taskFormCallback = new(TaskFormCallback)
	taskFormCallback.probe = probe
	taskFormCallback.task = task
	taskFormCallback.formGroup = formGroup

	taskFormCallback.CreationMode = (task == nil)

	return
}

type TaskFormCallback struct {
	task *models.Task

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (taskFormCallback *TaskFormCallback) OnSave() {
	taskFormCallback.probe.stageOfInterest.Lock()
	defer taskFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TaskFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	taskFormCallback.probe.formStage.Checkout()

	if taskFormCallback.task == nil {
		taskFormCallback.task = new(models.Task).Stage(taskFormCallback.probe.stageOfInterest)
	}
	task_ := taskFormCallback.task
	_ = task_

	for _, formDiv := range taskFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(task_.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(task_.ComputedPrefix), formDiv)
		case "IsStartTask":
			FormDivBasicFieldToField(&(task_.IsStartTask), formDiv)
		case "IsEndTask":
			FormDivBasicFieldToField(&(task_.IsEndTask), formDiv)
		case "DiagramProcess:TasksWhoseNodeIsExpanded":
			// WARNING : this form deals with the N-N association "DiagramProcess.TasksWhoseNodeIsExpanded []*Task" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Task). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.DiagramProcess
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "DiagramProcess"
				rf.Fieldname = "TasksWhoseNodeIsExpanded"
				formerAssociationSource := task_.GongGetReverseFieldOwner(
					taskFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.DiagramProcess)
					if !ok {
						log.Fatalln("Source of DiagramProcess.TasksWhoseNodeIsExpanded []*Task, is not an DiagramProcess instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.TasksWhoseNodeIsExpanded, task_)
					formerSource.TasksWhoseNodeIsExpanded = slices.Delete(formerSource.TasksWhoseNodeIsExpanded, idx, idx+1)
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
			var newSource *models.DiagramProcess
			for _diagramprocess := range *models.GetGongstructInstancesSet[models.DiagramProcess](taskFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagramprocess.GetName() == newSourceName.GetName() {
					newSource = _diagramprocess // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of DiagramProcess.TasksWhoseNodeIsExpanded []*Task, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.TasksWhoseNodeIsExpanded = append(newSource.TasksWhoseNodeIsExpanded, task_)
		case "Participant:Tasks":
			// WARNING : this form deals with the N-N association "Participant.Tasks []*Task" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Task). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Participant
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Participant"
				rf.Fieldname = "Tasks"
				formerAssociationSource := task_.GongGetReverseFieldOwner(
					taskFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Participant)
					if !ok {
						log.Fatalln("Source of Participant.Tasks []*Task, is not an Participant instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Tasks, task_)
					formerSource.Tasks = slices.Delete(formerSource.Tasks, idx, idx+1)
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
			var newSource *models.Participant
			for _participant := range *models.GetGongstructInstancesSet[models.Participant](taskFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _participant.GetName() == newSourceName.GetName() {
					newSource = _participant // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Participant.Tasks []*Task, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Tasks = append(newSource.Tasks, task_)
		case "Participant:TaskWhoseOutControlFlowsNodeIsExpanded":
			// WARNING : this form deals with the N-N association "Participant.TaskWhoseOutControlFlowsNodeIsExpanded []*Task" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Task). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Participant
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Participant"
				rf.Fieldname = "TaskWhoseOutControlFlowsNodeIsExpanded"
				formerAssociationSource := task_.GongGetReverseFieldOwner(
					taskFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Participant)
					if !ok {
						log.Fatalln("Source of Participant.TaskWhoseOutControlFlowsNodeIsExpanded []*Task, is not an Participant instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.TaskWhoseOutControlFlowsNodeIsExpanded, task_)
					formerSource.TaskWhoseOutControlFlowsNodeIsExpanded = slices.Delete(formerSource.TaskWhoseOutControlFlowsNodeIsExpanded, idx, idx+1)
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
			var newSource *models.Participant
			for _participant := range *models.GetGongstructInstancesSet[models.Participant](taskFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _participant.GetName() == newSourceName.GetName() {
					newSource = _participant // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Participant.TaskWhoseOutControlFlowsNodeIsExpanded []*Task, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.TaskWhoseOutControlFlowsNodeIsExpanded = append(newSource.TaskWhoseOutControlFlowsNodeIsExpanded, task_)
		case "Participant:TaskWhoseInControlFlowsNodeIsExpanded":
			// WARNING : this form deals with the N-N association "Participant.TaskWhoseInControlFlowsNodeIsExpanded []*Task" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Task). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Participant
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Participant"
				rf.Fieldname = "TaskWhoseInControlFlowsNodeIsExpanded"
				formerAssociationSource := task_.GongGetReverseFieldOwner(
					taskFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Participant)
					if !ok {
						log.Fatalln("Source of Participant.TaskWhoseInControlFlowsNodeIsExpanded []*Task, is not an Participant instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.TaskWhoseInControlFlowsNodeIsExpanded, task_)
					formerSource.TaskWhoseInControlFlowsNodeIsExpanded = slices.Delete(formerSource.TaskWhoseInControlFlowsNodeIsExpanded, idx, idx+1)
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
			var newSource *models.Participant
			for _participant := range *models.GetGongstructInstancesSet[models.Participant](taskFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _participant.GetName() == newSourceName.GetName() {
					newSource = _participant // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Participant.TaskWhoseInControlFlowsNodeIsExpanded []*Task, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.TaskWhoseInControlFlowsNodeIsExpanded = append(newSource.TaskWhoseInControlFlowsNodeIsExpanded, task_)
		}
	}

	// manage the suppress operation
	if taskFormCallback.formGroup.HasSuppressButtonBeenPressed {
		task_.Unstage(taskFormCallback.probe.stageOfInterest)
	}

	taskFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Task](
		taskFormCallback.probe,
	)

	// display a new form by reset the form stage
	if taskFormCallback.CreationMode || taskFormCallback.formGroup.HasSuppressButtonBeenPressed {
		taskFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(taskFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TaskFormCallback(
			nil,
			taskFormCallback.probe,
			newFormGroup,
		)
		task := new(models.Task)
		FillUpForm(task, newFormGroup, taskFormCallback.probe)
		taskFormCallback.probe.formStage.Commit()
	}

	taskFormCallback.probe.ux_tree()
}
func __gong__New__TaskShapeFormCallback(
	taskshape *models.TaskShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (taskshapeFormCallback *TaskShapeFormCallback) {
	taskshapeFormCallback = new(TaskShapeFormCallback)
	taskshapeFormCallback.probe = probe
	taskshapeFormCallback.taskshape = taskshape
	taskshapeFormCallback.formGroup = formGroup

	taskshapeFormCallback.CreationMode = (taskshape == nil)

	return
}

type TaskShapeFormCallback struct {
	taskshape *models.TaskShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (taskshapeFormCallback *TaskShapeFormCallback) OnSave() {
	taskshapeFormCallback.probe.stageOfInterest.Lock()
	defer taskshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TaskShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	taskshapeFormCallback.probe.formStage.Checkout()

	if taskshapeFormCallback.taskshape == nil {
		taskshapeFormCallback.taskshape = new(models.TaskShape).Stage(taskshapeFormCallback.probe.stageOfInterest)
	}
	taskshape_ := taskshapeFormCallback.taskshape
	_ = taskshape_

	for _, formDiv := range taskshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(taskshape_.Name), formDiv)
		case "Task":
			FormDivSelectFieldToField(&(taskshape_.Task), taskshapeFormCallback.probe.stageOfInterest, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(taskshape_.IsExpanded), formDiv)
		case "X":
			FormDivBasicFieldToField(&(taskshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(taskshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(taskshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(taskshape_.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(taskshape_.IsHidden), formDiv)
		case "DiagramProcess:Task_Shapes":
			// WARNING : this form deals with the N-N association "DiagramProcess.Task_Shapes []*TaskShape" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of TaskShape). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.DiagramProcess
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "DiagramProcess"
				rf.Fieldname = "Task_Shapes"
				formerAssociationSource := taskshape_.GongGetReverseFieldOwner(
					taskshapeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.DiagramProcess)
					if !ok {
						log.Fatalln("Source of DiagramProcess.Task_Shapes []*TaskShape, is not an DiagramProcess instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Task_Shapes, taskshape_)
					formerSource.Task_Shapes = slices.Delete(formerSource.Task_Shapes, idx, idx+1)
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
			var newSource *models.DiagramProcess
			for _diagramprocess := range *models.GetGongstructInstancesSet[models.DiagramProcess](taskshapeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _diagramprocess.GetName() == newSourceName.GetName() {
					newSource = _diagramprocess // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of DiagramProcess.Task_Shapes []*TaskShape, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Task_Shapes = append(newSource.Task_Shapes, taskshape_)
		}
	}

	// manage the suppress operation
	if taskshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		taskshape_.Unstage(taskshapeFormCallback.probe.stageOfInterest)
	}

	taskshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TaskShape](
		taskshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if taskshapeFormCallback.CreationMode || taskshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		taskshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(taskshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TaskShapeFormCallback(
			nil,
			taskshapeFormCallback.probe,
			newFormGroup,
		)
		taskshape := new(models.TaskShape)
		FillUpForm(taskshape, newFormGroup, taskshapeFormCallback.probe)
		taskshapeFormCallback.probe.formStage.Commit()
	}

	taskshapeFormCallback.probe.ux_tree()
}
