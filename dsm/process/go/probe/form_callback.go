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
func __gong__New__AllocatedProcessShapeFormCallback(
	allocatedprocessshape *models.AllocatedProcessShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (allocatedprocessshapeFormCallback *AllocatedProcessShapeFormCallback) {
	allocatedprocessshapeFormCallback = new(AllocatedProcessShapeFormCallback)
	allocatedprocessshapeFormCallback.probe = probe
	allocatedprocessshapeFormCallback.allocatedprocessshape = allocatedprocessshape
	allocatedprocessshapeFormCallback.formGroup = formGroup

	allocatedprocessshapeFormCallback.CreationMode = (allocatedprocessshape == nil)

	return
}

type AllocatedProcessShapeFormCallback struct {
	allocatedprocessshape *models.AllocatedProcessShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (allocatedprocessshapeFormCallback *AllocatedProcessShapeFormCallback) OnSave() {
	allocatedprocessshapeFormCallback.probe.stageOfInterest.Lock()
	defer allocatedprocessshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("AllocatedProcessShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	allocatedprocessshapeFormCallback.probe.formStage.Checkout()

	if allocatedprocessshapeFormCallback.allocatedprocessshape == nil {
		allocatedprocessshapeFormCallback.allocatedprocessshape = new(models.AllocatedProcessShape).Stage(allocatedprocessshapeFormCallback.probe.stageOfInterest)
	}
	allocatedprocessshape_ := allocatedprocessshapeFormCallback.allocatedprocessshape
	_ = allocatedprocessshape_

	for _, formDiv := range allocatedprocessshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(allocatedprocessshape_.Name), formDiv)
		case "Participant":
			FormDivSelectFieldToField(&(allocatedprocessshape_.Participant), allocatedprocessshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Process":
			FormDivSelectFieldToField(&(allocatedprocessshape_.Process), allocatedprocessshapeFormCallback.probe.stageOfInterest, formDiv)
		case "DiagramProcess:AllocatedProcessShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramProcess instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramProcess instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramProcess](allocatedprocessshapeFormCallback.probe.stageOfInterest)
			targetDiagramProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramProcess instances and update their AllocatedProcessShapes slice
			for _diagramprocess := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramProcess](allocatedprocessshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(allocatedprocessshapeFormCallback.probe.stageOfInterest, _diagramprocess)
				
				// if DiagramProcess is selected
				if targetDiagramProcessIDs[id] {
					// ensure allocatedprocessshape_ is in _diagramprocess.AllocatedProcessShapes
					found := false
					for _, _b := range _diagramprocess.AllocatedProcessShapes {
						if _b == allocatedprocessshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramprocess.AllocatedProcessShapes = append(_diagramprocess.AllocatedProcessShapes, allocatedprocessshape_)
						allocatedprocessshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "AllocatedProcessShapes", &_diagramprocess.AllocatedProcessShapes)
					}
				} else {
					// ensure allocatedprocessshape_ is NOT in _diagramprocess.AllocatedProcessShapes
					idx := slices.Index(_diagramprocess.AllocatedProcessShapes, allocatedprocessshape_)
					if idx != -1 {
						_diagramprocess.AllocatedProcessShapes = slices.Delete(_diagramprocess.AllocatedProcessShapes, idx, idx+1)
						allocatedprocessshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "AllocatedProcessShapes", &_diagramprocess.AllocatedProcessShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if allocatedprocessshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		allocatedprocessshape_.Unstage(allocatedprocessshapeFormCallback.probe.stageOfInterest)
	}

	allocatedprocessshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.AllocatedProcessShape](
		allocatedprocessshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if allocatedprocessshapeFormCallback.CreationMode || allocatedprocessshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		allocatedprocessshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(allocatedprocessshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AllocatedProcessShapeFormCallback(
			nil,
			allocatedprocessshapeFormCallback.probe,
			newFormGroup,
		)
		allocatedprocessshape := new(models.AllocatedProcessShape)
		FillUpForm(allocatedprocessshape, newFormGroup, allocatedprocessshapeFormCallback.probe)
		allocatedprocessshapeFormCallback.probe.formStage.Commit()
	}

	allocatedprocessshapeFormCallback.probe.ux_tree()
}
func __gong__New__AllocatedResourceShapeFormCallback(
	allocatedresourceshape *models.AllocatedResourceShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (allocatedresourceshapeFormCallback *AllocatedResourceShapeFormCallback) {
	allocatedresourceshapeFormCallback = new(AllocatedResourceShapeFormCallback)
	allocatedresourceshapeFormCallback.probe = probe
	allocatedresourceshapeFormCallback.allocatedresourceshape = allocatedresourceshape
	allocatedresourceshapeFormCallback.formGroup = formGroup

	allocatedresourceshapeFormCallback.CreationMode = (allocatedresourceshape == nil)

	return
}

type AllocatedResourceShapeFormCallback struct {
	allocatedresourceshape *models.AllocatedResourceShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (allocatedresourceshapeFormCallback *AllocatedResourceShapeFormCallback) OnSave() {
	allocatedresourceshapeFormCallback.probe.stageOfInterest.Lock()
	defer allocatedresourceshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("AllocatedResourceShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	allocatedresourceshapeFormCallback.probe.formStage.Checkout()

	if allocatedresourceshapeFormCallback.allocatedresourceshape == nil {
		allocatedresourceshapeFormCallback.allocatedresourceshape = new(models.AllocatedResourceShape).Stage(allocatedresourceshapeFormCallback.probe.stageOfInterest)
	}
	allocatedresourceshape_ := allocatedresourceshapeFormCallback.allocatedresourceshape
	_ = allocatedresourceshape_

	for _, formDiv := range allocatedresourceshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(allocatedresourceshape_.Name), formDiv)
		case "Participant":
			FormDivSelectFieldToField(&(allocatedresourceshape_.Participant), allocatedresourceshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Resource":
			FormDivSelectFieldToField(&(allocatedresourceshape_.Resource), allocatedresourceshapeFormCallback.probe.stageOfInterest, formDiv)
		case "DiagramProcess:AllocatedResourceShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramProcess instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramProcess instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramProcess](allocatedresourceshapeFormCallback.probe.stageOfInterest)
			targetDiagramProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramProcess instances and update their AllocatedResourceShapes slice
			for _diagramprocess := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramProcess](allocatedresourceshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(allocatedresourceshapeFormCallback.probe.stageOfInterest, _diagramprocess)
				
				// if DiagramProcess is selected
				if targetDiagramProcessIDs[id] {
					// ensure allocatedresourceshape_ is in _diagramprocess.AllocatedResourceShapes
					found := false
					for _, _b := range _diagramprocess.AllocatedResourceShapes {
						if _b == allocatedresourceshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramprocess.AllocatedResourceShapes = append(_diagramprocess.AllocatedResourceShapes, allocatedresourceshape_)
						allocatedresourceshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "AllocatedResourceShapes", &_diagramprocess.AllocatedResourceShapes)
					}
				} else {
					// ensure allocatedresourceshape_ is NOT in _diagramprocess.AllocatedResourceShapes
					idx := slices.Index(_diagramprocess.AllocatedResourceShapes, allocatedresourceshape_)
					if idx != -1 {
						_diagramprocess.AllocatedResourceShapes = slices.Delete(_diagramprocess.AllocatedResourceShapes, idx, idx+1)
						allocatedresourceshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "AllocatedResourceShapes", &_diagramprocess.AllocatedResourceShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if allocatedresourceshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		allocatedresourceshape_.Unstage(allocatedresourceshapeFormCallback.probe.stageOfInterest)
	}

	allocatedresourceshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.AllocatedResourceShape](
		allocatedresourceshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if allocatedresourceshapeFormCallback.CreationMode || allocatedresourceshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		allocatedresourceshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(allocatedresourceshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AllocatedResourceShapeFormCallback(
			nil,
			allocatedresourceshapeFormCallback.probe,
			newFormGroup,
		)
		allocatedresourceshape := new(models.AllocatedResourceShape)
		FillUpForm(allocatedresourceshape, newFormGroup, allocatedresourceshapeFormCallback.probe)
		allocatedresourceshapeFormCallback.probe.formStage.Commit()
	}

	allocatedresourceshapeFormCallback.probe.ux_tree()
}
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
		case "Description":
			FormDivBasicFieldToField(&(controlflow_.Description), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(controlflow_.ComputedPrefix), formDiv)
		case "Start":
			FormDivSelectFieldToField(&(controlflow_.Start), controlflowFormCallback.probe.stageOfInterest, formDiv)
		case "End":
			FormDivSelectFieldToField(&(controlflow_.End), controlflowFormCallback.probe.stageOfInterest, formDiv)
		case "DiagramProcess:ControlFlowsWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramProcess instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramProcess instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramProcess](controlflowFormCallback.probe.stageOfInterest)
			targetDiagramProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramProcess instances and update their ControlFlowsWhoseNodeIsExpanded slice
			for _diagramprocess := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramProcess](controlflowFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(controlflowFormCallback.probe.stageOfInterest, _diagramprocess)
				
				// if DiagramProcess is selected
				if targetDiagramProcessIDs[id] {
					// ensure controlflow_ is in _diagramprocess.ControlFlowsWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramprocess.ControlFlowsWhoseNodeIsExpanded {
						if _b == controlflow_ {
							found = true
							break
						}
					}
					if !found {
						_diagramprocess.ControlFlowsWhoseNodeIsExpanded = append(_diagramprocess.ControlFlowsWhoseNodeIsExpanded, controlflow_)
						controlflowFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "ControlFlowsWhoseNodeIsExpanded", &_diagramprocess.ControlFlowsWhoseNodeIsExpanded)
					}
				} else {
					// ensure controlflow_ is NOT in _diagramprocess.ControlFlowsWhoseNodeIsExpanded
					idx := slices.Index(_diagramprocess.ControlFlowsWhoseNodeIsExpanded, controlflow_)
					if idx != -1 {
						_diagramprocess.ControlFlowsWhoseNodeIsExpanded = slices.Delete(_diagramprocess.ControlFlowsWhoseNodeIsExpanded, idx, idx+1)
						controlflowFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "ControlFlowsWhoseNodeIsExpanded", &_diagramprocess.ControlFlowsWhoseNodeIsExpanded)
					}
				}
			}
		case "Participant:ControlFlows":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Participant instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Participant instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Participant](controlflowFormCallback.probe.stageOfInterest)
			targetParticipantIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetParticipantIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Participant instances and update their ControlFlows slice
			for _participant := range *models.GetGongstructInstancesSetFromPointerType[*models.Participant](controlflowFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(controlflowFormCallback.probe.stageOfInterest, _participant)
				
				// if Participant is selected
				if targetParticipantIDs[id] {
					// ensure controlflow_ is in _participant.ControlFlows
					found := false
					for _, _b := range _participant.ControlFlows {
						if _b == controlflow_ {
							found = true
							break
						}
					}
					if !found {
						_participant.ControlFlows = append(_participant.ControlFlows, controlflow_)
						controlflowFormCallback.probe.UpdateSliceOfPointersCallback(_participant, "ControlFlows", &_participant.ControlFlows)
					}
				} else {
					// ensure controlflow_ is NOT in _participant.ControlFlows
					idx := slices.Index(_participant.ControlFlows, controlflow_)
					if idx != -1 {
						_participant.ControlFlows = slices.Delete(_participant.ControlFlows, idx, idx+1)
						controlflowFormCallback.probe.UpdateSliceOfPointersCallback(_participant, "ControlFlows", &_participant.ControlFlows)
					}
				}
			}
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
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramProcess instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramProcess instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramProcess](controlflowshapeFormCallback.probe.stageOfInterest)
			targetDiagramProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramProcess instances and update their ControlFlow_Shapes slice
			for _diagramprocess := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramProcess](controlflowshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(controlflowshapeFormCallback.probe.stageOfInterest, _diagramprocess)
				
				// if DiagramProcess is selected
				if targetDiagramProcessIDs[id] {
					// ensure controlflowshape_ is in _diagramprocess.ControlFlow_Shapes
					found := false
					for _, _b := range _diagramprocess.ControlFlow_Shapes {
						if _b == controlflowshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramprocess.ControlFlow_Shapes = append(_diagramprocess.ControlFlow_Shapes, controlflowshape_)
						controlflowshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "ControlFlow_Shapes", &_diagramprocess.ControlFlow_Shapes)
					}
				} else {
					// ensure controlflowshape_ is NOT in _diagramprocess.ControlFlow_Shapes
					idx := slices.Index(_diagramprocess.ControlFlow_Shapes, controlflowshape_)
					if idx != -1 {
						_diagramprocess.ControlFlow_Shapes = slices.Delete(_diagramprocess.ControlFlow_Shapes, idx, idx+1)
						controlflowshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "ControlFlow_Shapes", &_diagramprocess.ControlFlow_Shapes)
					}
				}
			}
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
func __gong__New__DataFormCallback(
	data *models.Data,
	probe *Probe,
	formGroup *form.FormGroup,
) (dataFormCallback *DataFormCallback) {
	dataFormCallback = new(DataFormCallback)
	dataFormCallback.probe = probe
	dataFormCallback.data = data
	dataFormCallback.formGroup = formGroup

	dataFormCallback.CreationMode = (data == nil)

	return
}

type DataFormCallback struct {
	data *models.Data

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (dataFormCallback *DataFormCallback) OnSave() {
	dataFormCallback.probe.stageOfInterest.Lock()
	defer dataFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DataFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	dataFormCallback.probe.formStage.Checkout()

	if dataFormCallback.data == nil {
		dataFormCallback.data = new(models.Data).Stage(dataFormCallback.probe.stageOfInterest)
	}
	data_ := dataFormCallback.data
	_ = data_

	for _, formDiv := range dataFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(data_.Name), formDiv)
		case "Acronym":
			FormDivBasicFieldToField(&(data_.Acronym), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(data_.Description), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(data_.ComputedPrefix), formDiv)
		case "SVG_Path":
			FormDivBasicFieldToField(&(data_.SVG_Path), formDiv)
		case "InverseAppliedScaling":
			FormDivBasicFieldToField(&(data_.InverseAppliedScaling), formDiv)
		case "DataFlow:Datas":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DataFlow instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DataFlow instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DataFlow](dataFormCallback.probe.stageOfInterest)
			targetDataFlowIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDataFlowIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DataFlow instances and update their Datas slice
			for _dataflow := range *models.GetGongstructInstancesSetFromPointerType[*models.DataFlow](dataFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(dataFormCallback.probe.stageOfInterest, _dataflow)
				
				// if DataFlow is selected
				if targetDataFlowIDs[id] {
					// ensure data_ is in _dataflow.Datas
					found := false
					for _, _b := range _dataflow.Datas {
						if _b == data_ {
							found = true
							break
						}
					}
					if !found {
						_dataflow.Datas = append(_dataflow.Datas, data_)
						dataFormCallback.probe.UpdateSliceOfPointersCallback(_dataflow, "Datas", &_dataflow.Datas)
					}
				} else {
					// ensure data_ is NOT in _dataflow.Datas
					idx := slices.Index(_dataflow.Datas, data_)
					if idx != -1 {
						_dataflow.Datas = slices.Delete(_dataflow.Datas, idx, idx+1)
						dataFormCallback.probe.UpdateSliceOfPointersCallback(_dataflow, "Datas", &_dataflow.Datas)
					}
				}
			}
		case "DiagramProcess:DatasWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramProcess instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramProcess instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramProcess](dataFormCallback.probe.stageOfInterest)
			targetDiagramProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramProcess instances and update their DatasWhoseNodeIsExpanded slice
			for _diagramprocess := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramProcess](dataFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(dataFormCallback.probe.stageOfInterest, _diagramprocess)
				
				// if DiagramProcess is selected
				if targetDiagramProcessIDs[id] {
					// ensure data_ is in _diagramprocess.DatasWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramprocess.DatasWhoseNodeIsExpanded {
						if _b == data_ {
							found = true
							break
						}
					}
					if !found {
						_diagramprocess.DatasWhoseNodeIsExpanded = append(_diagramprocess.DatasWhoseNodeIsExpanded, data_)
						dataFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "DatasWhoseNodeIsExpanded", &_diagramprocess.DatasWhoseNodeIsExpanded)
					}
				} else {
					// ensure data_ is NOT in _diagramprocess.DatasWhoseNodeIsExpanded
					idx := slices.Index(_diagramprocess.DatasWhoseNodeIsExpanded, data_)
					if idx != -1 {
						_diagramprocess.DatasWhoseNodeIsExpanded = slices.Delete(_diagramprocess.DatasWhoseNodeIsExpanded, idx, idx+1)
						dataFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "DatasWhoseNodeIsExpanded", &_diagramprocess.DatasWhoseNodeIsExpanded)
					}
				}
			}
		case "Library:RootDatas":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](dataFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their RootDatas slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](dataFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(dataFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure data_ is in _library.RootDatas
					found := false
					for _, _b := range _library.RootDatas {
						if _b == data_ {
							found = true
							break
						}
					}
					if !found {
						_library.RootDatas = append(_library.RootDatas, data_)
						dataFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootDatas", &_library.RootDatas)
					}
				} else {
					// ensure data_ is NOT in _library.RootDatas
					idx := slices.Index(_library.RootDatas, data_)
					if idx != -1 {
						_library.RootDatas = slices.Delete(_library.RootDatas, idx, idx+1)
						dataFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootDatas", &_library.RootDatas)
					}
				}
			}
		case "Library:DatasWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](dataFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their DatasWhoseNodeIsExpanded slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](dataFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(dataFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure data_ is in _library.DatasWhoseNodeIsExpanded
					found := false
					for _, _b := range _library.DatasWhoseNodeIsExpanded {
						if _b == data_ {
							found = true
							break
						}
					}
					if !found {
						_library.DatasWhoseNodeIsExpanded = append(_library.DatasWhoseNodeIsExpanded, data_)
						dataFormCallback.probe.UpdateSliceOfPointersCallback(_library, "DatasWhoseNodeIsExpanded", &_library.DatasWhoseNodeIsExpanded)
					}
				} else {
					// ensure data_ is NOT in _library.DatasWhoseNodeIsExpanded
					idx := slices.Index(_library.DatasWhoseNodeIsExpanded, data_)
					if idx != -1 {
						_library.DatasWhoseNodeIsExpanded = slices.Delete(_library.DatasWhoseNodeIsExpanded, idx, idx+1)
						dataFormCallback.probe.UpdateSliceOfPointersCallback(_library, "DatasWhoseNodeIsExpanded", &_library.DatasWhoseNodeIsExpanded)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if dataFormCallback.formGroup.HasSuppressButtonBeenPressed {
		data_.Unstage(dataFormCallback.probe.stageOfInterest)
	}

	dataFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Data](
		dataFormCallback.probe,
	)

	// display a new form by reset the form stage
	if dataFormCallback.CreationMode || dataFormCallback.formGroup.HasSuppressButtonBeenPressed {
		dataFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(dataFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DataFormCallback(
			nil,
			dataFormCallback.probe,
			newFormGroup,
		)
		data := new(models.Data)
		FillUpForm(data, newFormGroup, dataFormCallback.probe)
		dataFormCallback.probe.formStage.Commit()
	}

	dataFormCallback.probe.ux_tree()
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
		case "Datas":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Data](dataflowFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Data, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Data)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					dataflowFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Data](dataflowFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			dataflow_.Datas = instanceSlice
			dataflowFormCallback.probe.UpdateSliceOfPointersCallback(dataflow_, "Datas", &dataflow_.Datas)

		case "Description":
			FormDivBasicFieldToField(&(dataflow_.Description), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(dataflow_.ComputedPrefix), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(dataflow_.Type), formDiv)
		case "StartTask":
			FormDivSelectFieldToField(&(dataflow_.StartTask), dataflowFormCallback.probe.stageOfInterest, formDiv)
		case "EndTask":
			FormDivSelectFieldToField(&(dataflow_.EndTask), dataflowFormCallback.probe.stageOfInterest, formDiv)
		case "StartExternalParticipant":
			FormDivSelectFieldToField(&(dataflow_.StartExternalParticipant), dataflowFormCallback.probe.stageOfInterest, formDiv)
		case "EndExternalParticipant":
			FormDivSelectFieldToField(&(dataflow_.EndExternalParticipant), dataflowFormCallback.probe.stageOfInterest, formDiv)
		case "IsDatasNodeExpanded":
			FormDivBasicFieldToField(&(dataflow_.IsDatasNodeExpanded), formDiv)
		case "DiagramProcess:DataFlowsWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramProcess instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramProcess instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramProcess](dataflowFormCallback.probe.stageOfInterest)
			targetDiagramProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramProcess instances and update their DataFlowsWhoseNodeIsExpanded slice
			for _diagramprocess := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramProcess](dataflowFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(dataflowFormCallback.probe.stageOfInterest, _diagramprocess)
				
				// if DiagramProcess is selected
				if targetDiagramProcessIDs[id] {
					// ensure dataflow_ is in _diagramprocess.DataFlowsWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramprocess.DataFlowsWhoseNodeIsExpanded {
						if _b == dataflow_ {
							found = true
							break
						}
					}
					if !found {
						_diagramprocess.DataFlowsWhoseNodeIsExpanded = append(_diagramprocess.DataFlowsWhoseNodeIsExpanded, dataflow_)
						dataflowFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "DataFlowsWhoseNodeIsExpanded", &_diagramprocess.DataFlowsWhoseNodeIsExpanded)
					}
				} else {
					// ensure dataflow_ is NOT in _diagramprocess.DataFlowsWhoseNodeIsExpanded
					idx := slices.Index(_diagramprocess.DataFlowsWhoseNodeIsExpanded, dataflow_)
					if idx != -1 {
						_diagramprocess.DataFlowsWhoseNodeIsExpanded = slices.Delete(_diagramprocess.DataFlowsWhoseNodeIsExpanded, idx, idx+1)
						dataflowFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "DataFlowsWhoseNodeIsExpanded", &_diagramprocess.DataFlowsWhoseNodeIsExpanded)
					}
				}
			}
		case "DiagramProcess:DataFlowsWhoseDataNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramProcess instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramProcess instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramProcess](dataflowFormCallback.probe.stageOfInterest)
			targetDiagramProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramProcess instances and update their DataFlowsWhoseDataNodeIsExpanded slice
			for _diagramprocess := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramProcess](dataflowFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(dataflowFormCallback.probe.stageOfInterest, _diagramprocess)
				
				// if DiagramProcess is selected
				if targetDiagramProcessIDs[id] {
					// ensure dataflow_ is in _diagramprocess.DataFlowsWhoseDataNodeIsExpanded
					found := false
					for _, _b := range _diagramprocess.DataFlowsWhoseDataNodeIsExpanded {
						if _b == dataflow_ {
							found = true
							break
						}
					}
					if !found {
						_diagramprocess.DataFlowsWhoseDataNodeIsExpanded = append(_diagramprocess.DataFlowsWhoseDataNodeIsExpanded, dataflow_)
						dataflowFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "DataFlowsWhoseDataNodeIsExpanded", &_diagramprocess.DataFlowsWhoseDataNodeIsExpanded)
					}
				} else {
					// ensure dataflow_ is NOT in _diagramprocess.DataFlowsWhoseDataNodeIsExpanded
					idx := slices.Index(_diagramprocess.DataFlowsWhoseDataNodeIsExpanded, dataflow_)
					if idx != -1 {
						_diagramprocess.DataFlowsWhoseDataNodeIsExpanded = slices.Delete(_diagramprocess.DataFlowsWhoseDataNodeIsExpanded, idx, idx+1)
						dataflowFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "DataFlowsWhoseDataNodeIsExpanded", &_diagramprocess.DataFlowsWhoseDataNodeIsExpanded)
					}
				}
			}
		case "Library:RootDataFlows":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](dataflowFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their RootDataFlows slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](dataflowFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(dataflowFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure dataflow_ is in _library.RootDataFlows
					found := false
					for _, _b := range _library.RootDataFlows {
						if _b == dataflow_ {
							found = true
							break
						}
					}
					if !found {
						_library.RootDataFlows = append(_library.RootDataFlows, dataflow_)
						dataflowFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootDataFlows", &_library.RootDataFlows)
					}
				} else {
					// ensure dataflow_ is NOT in _library.RootDataFlows
					idx := slices.Index(_library.RootDataFlows, dataflow_)
					if idx != -1 {
						_library.RootDataFlows = slices.Delete(_library.RootDataFlows, idx, idx+1)
						dataflowFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootDataFlows", &_library.RootDataFlows)
					}
				}
			}
		case "Library:DataFlowsWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](dataflowFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their DataFlowsWhoseNodeIsExpanded slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](dataflowFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(dataflowFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure dataflow_ is in _library.DataFlowsWhoseNodeIsExpanded
					found := false
					for _, _b := range _library.DataFlowsWhoseNodeIsExpanded {
						if _b == dataflow_ {
							found = true
							break
						}
					}
					if !found {
						_library.DataFlowsWhoseNodeIsExpanded = append(_library.DataFlowsWhoseNodeIsExpanded, dataflow_)
						dataflowFormCallback.probe.UpdateSliceOfPointersCallback(_library, "DataFlowsWhoseNodeIsExpanded", &_library.DataFlowsWhoseNodeIsExpanded)
					}
				} else {
					// ensure dataflow_ is NOT in _library.DataFlowsWhoseNodeIsExpanded
					idx := slices.Index(_library.DataFlowsWhoseNodeIsExpanded, dataflow_)
					if idx != -1 {
						_library.DataFlowsWhoseNodeIsExpanded = slices.Delete(_library.DataFlowsWhoseNodeIsExpanded, idx, idx+1)
						dataflowFormCallback.probe.UpdateSliceOfPointersCallback(_library, "DataFlowsWhoseNodeIsExpanded", &_library.DataFlowsWhoseNodeIsExpanded)
					}
				}
			}
		case "Process:DataFlows":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Process instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Process instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Process](dataflowFormCallback.probe.stageOfInterest)
			targetProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Process instances and update their DataFlows slice
			for _process := range *models.GetGongstructInstancesSetFromPointerType[*models.Process](dataflowFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(dataflowFormCallback.probe.stageOfInterest, _process)
				
				// if Process is selected
				if targetProcessIDs[id] {
					// ensure dataflow_ is in _process.DataFlows
					found := false
					for _, _b := range _process.DataFlows {
						if _b == dataflow_ {
							found = true
							break
						}
					}
					if !found {
						_process.DataFlows = append(_process.DataFlows, dataflow_)
						dataflowFormCallback.probe.UpdateSliceOfPointersCallback(_process, "DataFlows", &_process.DataFlows)
					}
				} else {
					// ensure dataflow_ is NOT in _process.DataFlows
					idx := slices.Index(_process.DataFlows, dataflow_)
					if idx != -1 {
						_process.DataFlows = slices.Delete(_process.DataFlows, idx, idx+1)
						dataflowFormCallback.probe.UpdateSliceOfPointersCallback(_process, "DataFlows", &_process.DataFlows)
					}
				}
			}
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
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramProcess instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramProcess instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramProcess](dataflowshapeFormCallback.probe.stageOfInterest)
			targetDiagramProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramProcess instances and update their DataFlow_Shapes slice
			for _diagramprocess := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramProcess](dataflowshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(dataflowshapeFormCallback.probe.stageOfInterest, _diagramprocess)
				
				// if DiagramProcess is selected
				if targetDiagramProcessIDs[id] {
					// ensure dataflowshape_ is in _diagramprocess.DataFlow_Shapes
					found := false
					for _, _b := range _diagramprocess.DataFlow_Shapes {
						if _b == dataflowshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramprocess.DataFlow_Shapes = append(_diagramprocess.DataFlow_Shapes, dataflowshape_)
						dataflowshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "DataFlow_Shapes", &_diagramprocess.DataFlow_Shapes)
					}
				} else {
					// ensure dataflowshape_ is NOT in _diagramprocess.DataFlow_Shapes
					idx := slices.Index(_diagramprocess.DataFlow_Shapes, dataflowshape_)
					if idx != -1 {
						_diagramprocess.DataFlow_Shapes = slices.Delete(_diagramprocess.DataFlow_Shapes, idx, idx+1)
						dataflowshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "DataFlow_Shapes", &_diagramprocess.DataFlow_Shapes)
					}
				}
			}
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
func __gong__New__DataShapeFormCallback(
	datashape *models.DataShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (datashapeFormCallback *DataShapeFormCallback) {
	datashapeFormCallback = new(DataShapeFormCallback)
	datashapeFormCallback.probe = probe
	datashapeFormCallback.datashape = datashape
	datashapeFormCallback.formGroup = formGroup

	datashapeFormCallback.CreationMode = (datashape == nil)

	return
}

type DataShapeFormCallback struct {
	datashape *models.DataShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (datashapeFormCallback *DataShapeFormCallback) OnSave() {
	datashapeFormCallback.probe.stageOfInterest.Lock()
	defer datashapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DataShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	datashapeFormCallback.probe.formStage.Checkout()

	if datashapeFormCallback.datashape == nil {
		datashapeFormCallback.datashape = new(models.DataShape).Stage(datashapeFormCallback.probe.stageOfInterest)
	}
	datashape_ := datashapeFormCallback.datashape
	_ = datashape_

	for _, formDiv := range datashapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(datashape_.Name), formDiv)
		case "Data":
			FormDivSelectFieldToField(&(datashape_.Data), datashapeFormCallback.probe.stageOfInterest, formDiv)
		case "DataFlow":
			FormDivSelectFieldToField(&(datashape_.DataFlow), datashapeFormCallback.probe.stageOfInterest, formDiv)
		case "DiagramProcess:Data_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramProcess instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramProcess instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramProcess](datashapeFormCallback.probe.stageOfInterest)
			targetDiagramProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramProcess instances and update their Data_Shapes slice
			for _diagramprocess := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramProcess](datashapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(datashapeFormCallback.probe.stageOfInterest, _diagramprocess)
				
				// if DiagramProcess is selected
				if targetDiagramProcessIDs[id] {
					// ensure datashape_ is in _diagramprocess.Data_Shapes
					found := false
					for _, _b := range _diagramprocess.Data_Shapes {
						if _b == datashape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramprocess.Data_Shapes = append(_diagramprocess.Data_Shapes, datashape_)
						datashapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "Data_Shapes", &_diagramprocess.Data_Shapes)
					}
				} else {
					// ensure datashape_ is NOT in _diagramprocess.Data_Shapes
					idx := slices.Index(_diagramprocess.Data_Shapes, datashape_)
					if idx != -1 {
						_diagramprocess.Data_Shapes = slices.Delete(_diagramprocess.Data_Shapes, idx, idx+1)
						datashapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "Data_Shapes", &_diagramprocess.Data_Shapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if datashapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datashape_.Unstage(datashapeFormCallback.probe.stageOfInterest)
	}

	datashapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.DataShape](
		datashapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if datashapeFormCallback.CreationMode || datashapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		datashapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(datashapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DataShapeFormCallback(
			nil,
			datashapeFormCallback.probe,
			newFormGroup,
		)
		datashape := new(models.DataShape)
		FillUpForm(datashape, newFormGroup, datashapeFormCallback.probe)
		datashapeFormCallback.probe.formStage.Commit()
	}

	datashapeFormCallback.probe.ux_tree()
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
		case "Description":
			FormDivBasicFieldToField(&(diagramprocess_.Description), formDiv)
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
			diagramprocessFormCallback.probe.UpdateSliceOfPointersCallback(diagramprocess_, "Process_Shapes", &diagramprocess_.Process_Shapes)

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
			diagramprocessFormCallback.probe.UpdateSliceOfPointersCallback(diagramprocess_, "ProcesssWhoseNodeIsExpanded", &diagramprocess_.ProcesssWhoseNodeIsExpanded)

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
			diagramprocessFormCallback.probe.UpdateSliceOfPointersCallback(diagramprocess_, "Participant_Shapes", &diagramprocess_.Participant_Shapes)

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
			diagramprocessFormCallback.probe.UpdateSliceOfPointersCallback(diagramprocess_, "ParticipantWhoseNodeIsExpanded", &diagramprocess_.ParticipantWhoseNodeIsExpanded)

		case "ExternalParticipant_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ExternalParticipantShape](diagramprocessFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ExternalParticipantShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ExternalParticipantShape)

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
			map_RowID_ID := GetMap_RowID_ID[*models.ExternalParticipantShape](diagramprocessFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramprocess_.ExternalParticipant_Shapes = instanceSlice
			diagramprocessFormCallback.probe.UpdateSliceOfPointersCallback(diagramprocess_, "ExternalParticipant_Shapes", &diagramprocess_.ExternalParticipant_Shapes)

		case "IsExternalParticipantsNodeExpanded":
			FormDivBasicFieldToField(&(diagramprocess_.IsExternalParticipantsNodeExpanded), formDiv)
		case "ExternalParticipantWhoseNodeIsExpanded":
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
			diagramprocess_.ExternalParticipantWhoseNodeIsExpanded = instanceSlice
			diagramprocessFormCallback.probe.UpdateSliceOfPointersCallback(diagramprocess_, "ExternalParticipantWhoseNodeIsExpanded", &diagramprocess_.ExternalParticipantWhoseNodeIsExpanded)

		case "ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded":
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
			diagramprocess_.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded = instanceSlice
			diagramprocessFormCallback.probe.UpdateSliceOfPointersCallback(diagramprocess_, "ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded", &diagramprocess_.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded)

		case "ExternalParticipantsWhoseInDataFlowsNodeIsExpanded":
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
			diagramprocess_.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded = instanceSlice
			diagramprocessFormCallback.probe.UpdateSliceOfPointersCallback(diagramprocess_, "ExternalParticipantsWhoseInDataFlowsNodeIsExpanded", &diagramprocess_.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded)

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
			diagramprocessFormCallback.probe.UpdateSliceOfPointersCallback(diagramprocess_, "TasksWhoseNodeIsExpanded", &diagramprocess_.TasksWhoseNodeIsExpanded)

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
			diagramprocessFormCallback.probe.UpdateSliceOfPointersCallback(diagramprocess_, "Task_Shapes", &diagramprocess_.Task_Shapes)

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
			diagramprocessFormCallback.probe.UpdateSliceOfPointersCallback(diagramprocess_, "ControlFlowsWhoseNodeIsExpanded", &diagramprocess_.ControlFlowsWhoseNodeIsExpanded)

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
			diagramprocessFormCallback.probe.UpdateSliceOfPointersCallback(diagramprocess_, "ControlFlow_Shapes", &diagramprocess_.ControlFlow_Shapes)

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
			diagramprocessFormCallback.probe.UpdateSliceOfPointersCallback(diagramprocess_, "DataFlowsWhoseNodeIsExpanded", &diagramprocess_.DataFlowsWhoseNodeIsExpanded)

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
			diagramprocessFormCallback.probe.UpdateSliceOfPointersCallback(diagramprocess_, "DataFlow_Shapes", &diagramprocess_.DataFlow_Shapes)

		case "DatasWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Data](diagramprocessFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Data, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Data)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Data](diagramprocessFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramprocess_.DatasWhoseNodeIsExpanded = instanceSlice
			diagramprocessFormCallback.probe.UpdateSliceOfPointersCallback(diagramprocess_, "DatasWhoseNodeIsExpanded", &diagramprocess_.DatasWhoseNodeIsExpanded)

		case "Data_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DataShape](diagramprocessFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DataShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DataShape)

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
			map_RowID_ID := GetMap_RowID_ID[*models.DataShape](diagramprocessFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramprocess_.Data_Shapes = instanceSlice
			diagramprocessFormCallback.probe.UpdateSliceOfPointersCallback(diagramprocess_, "Data_Shapes", &diagramprocess_.Data_Shapes)

		case "DataFlowsWhoseDataNodeIsExpanded":
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
			diagramprocess_.DataFlowsWhoseDataNodeIsExpanded = instanceSlice
			diagramprocessFormCallback.probe.UpdateSliceOfPointersCallback(diagramprocess_, "DataFlowsWhoseDataNodeIsExpanded", &diagramprocess_.DataFlowsWhoseDataNodeIsExpanded)

		case "AllocatedResourcesWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Resource](diagramprocessFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Resource, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Resource)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Resource](diagramprocessFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramprocess_.AllocatedResourcesWhoseNodeIsExpanded = instanceSlice
			diagramprocessFormCallback.probe.UpdateSliceOfPointersCallback(diagramprocess_, "AllocatedResourcesWhoseNodeIsExpanded", &diagramprocess_.AllocatedResourcesWhoseNodeIsExpanded)

		case "AllocatedResourceShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.AllocatedResourceShape](diagramprocessFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.AllocatedResourceShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.AllocatedResourceShape)

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
			map_RowID_ID := GetMap_RowID_ID[*models.AllocatedResourceShape](diagramprocessFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramprocess_.AllocatedResourceShapes = instanceSlice
			diagramprocessFormCallback.probe.UpdateSliceOfPointersCallback(diagramprocess_, "AllocatedResourceShapes", &diagramprocess_.AllocatedResourceShapes)

		case "AllocatedProcessesWhoseNodeIsExpanded":
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
			diagramprocess_.AllocatedProcessesWhoseNodeIsExpanded = instanceSlice
			diagramprocessFormCallback.probe.UpdateSliceOfPointersCallback(diagramprocess_, "AllocatedProcessesWhoseNodeIsExpanded", &diagramprocess_.AllocatedProcessesWhoseNodeIsExpanded)

		case "AllocatedProcessShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.AllocatedProcessShape](diagramprocessFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.AllocatedProcessShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.AllocatedProcessShape)

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
			map_RowID_ID := GetMap_RowID_ID[*models.AllocatedProcessShape](diagramprocessFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramprocess_.AllocatedProcessShapes = instanceSlice
			diagramprocessFormCallback.probe.UpdateSliceOfPointersCallback(diagramprocess_, "AllocatedProcessShapes", &diagramprocess_.AllocatedProcessShapes)

		case "Note_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.NoteShape](diagramprocessFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.NoteShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.NoteShape)

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
			map_RowID_ID := GetMap_RowID_ID[*models.NoteShape](diagramprocessFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramprocess_.Note_Shapes = instanceSlice
			diagramprocessFormCallback.probe.UpdateSliceOfPointersCallback(diagramprocess_, "Note_Shapes", &diagramprocess_.Note_Shapes)

		case "NotesWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Note](diagramprocessFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Note, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Note)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Note](diagramprocessFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramprocess_.NotesWhoseNodeIsExpanded = instanceSlice
			diagramprocessFormCallback.probe.UpdateSliceOfPointersCallback(diagramprocess_, "NotesWhoseNodeIsExpanded", &diagramprocess_.NotesWhoseNodeIsExpanded)

		case "IsNotesNodeExpanded":
			FormDivBasicFieldToField(&(diagramprocess_.IsNotesNodeExpanded), formDiv)
		case "NoteTaskShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.NoteTaskShape](diagramprocessFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.NoteTaskShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.NoteTaskShape)

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
			map_RowID_ID := GetMap_RowID_ID[*models.NoteTaskShape](diagramprocessFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramprocess_.NoteTaskShapes = instanceSlice
			diagramprocessFormCallback.probe.UpdateSliceOfPointersCallback(diagramprocess_, "NoteTaskShapes", &diagramprocess_.NoteTaskShapes)

		case "Process:DiagramProcesss":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Process instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Process instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Process](diagramprocessFormCallback.probe.stageOfInterest)
			targetProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Process instances and update their DiagramProcesss slice
			for _process := range *models.GetGongstructInstancesSetFromPointerType[*models.Process](diagramprocessFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(diagramprocessFormCallback.probe.stageOfInterest, _process)
				
				// if Process is selected
				if targetProcessIDs[id] {
					// ensure diagramprocess_ is in _process.DiagramProcesss
					found := false
					for _, _b := range _process.DiagramProcesss {
						if _b == diagramprocess_ {
							found = true
							break
						}
					}
					if !found {
						_process.DiagramProcesss = append(_process.DiagramProcesss, diagramprocess_)
						diagramprocessFormCallback.probe.UpdateSliceOfPointersCallback(_process, "DiagramProcesss", &_process.DiagramProcesss)
					}
				} else {
					// ensure diagramprocess_ is NOT in _process.DiagramProcesss
					idx := slices.Index(_process.DiagramProcesss, diagramprocess_)
					if idx != -1 {
						_process.DiagramProcesss = slices.Delete(_process.DiagramProcesss, idx, idx+1)
						diagramprocessFormCallback.probe.UpdateSliceOfPointersCallback(_process, "DiagramProcesss", &_process.DiagramProcesss)
					}
				}
			}
		case "Process:DiagramProcessWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Process instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Process instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Process](diagramprocessFormCallback.probe.stageOfInterest)
			targetProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Process instances and update their DiagramProcessWhoseNodeIsExpanded slice
			for _process := range *models.GetGongstructInstancesSetFromPointerType[*models.Process](diagramprocessFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(diagramprocessFormCallback.probe.stageOfInterest, _process)
				
				// if Process is selected
				if targetProcessIDs[id] {
					// ensure diagramprocess_ is in _process.DiagramProcessWhoseNodeIsExpanded
					found := false
					for _, _b := range _process.DiagramProcessWhoseNodeIsExpanded {
						if _b == diagramprocess_ {
							found = true
							break
						}
					}
					if !found {
						_process.DiagramProcessWhoseNodeIsExpanded = append(_process.DiagramProcessWhoseNodeIsExpanded, diagramprocess_)
						diagramprocessFormCallback.probe.UpdateSliceOfPointersCallback(_process, "DiagramProcessWhoseNodeIsExpanded", &_process.DiagramProcessWhoseNodeIsExpanded)
					}
				} else {
					// ensure diagramprocess_ is NOT in _process.DiagramProcessWhoseNodeIsExpanded
					idx := slices.Index(_process.DiagramProcessWhoseNodeIsExpanded, diagramprocess_)
					if idx != -1 {
						_process.DiagramProcessWhoseNodeIsExpanded = slices.Delete(_process.DiagramProcessWhoseNodeIsExpanded, idx, idx+1)
						diagramprocessFormCallback.probe.UpdateSliceOfPointersCallback(_process, "DiagramProcessWhoseNodeIsExpanded", &_process.DiagramProcessWhoseNodeIsExpanded)
					}
				}
			}
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
func __gong__New__ExternalParticipantShapeFormCallback(
	externalparticipantshape *models.ExternalParticipantShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (externalparticipantshapeFormCallback *ExternalParticipantShapeFormCallback) {
	externalparticipantshapeFormCallback = new(ExternalParticipantShapeFormCallback)
	externalparticipantshapeFormCallback.probe = probe
	externalparticipantshapeFormCallback.externalparticipantshape = externalparticipantshape
	externalparticipantshapeFormCallback.formGroup = formGroup

	externalparticipantshapeFormCallback.CreationMode = (externalparticipantshape == nil)

	return
}

type ExternalParticipantShapeFormCallback struct {
	externalparticipantshape *models.ExternalParticipantShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (externalparticipantshapeFormCallback *ExternalParticipantShapeFormCallback) OnSave() {
	externalparticipantshapeFormCallback.probe.stageOfInterest.Lock()
	defer externalparticipantshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ExternalParticipantShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	externalparticipantshapeFormCallback.probe.formStage.Checkout()

	if externalparticipantshapeFormCallback.externalparticipantshape == nil {
		externalparticipantshapeFormCallback.externalparticipantshape = new(models.ExternalParticipantShape).Stage(externalparticipantshapeFormCallback.probe.stageOfInterest)
	}
	externalparticipantshape_ := externalparticipantshapeFormCallback.externalparticipantshape
	_ = externalparticipantshape_

	for _, formDiv := range externalparticipantshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(externalparticipantshape_.Name), formDiv)
		case "Participant":
			FormDivSelectFieldToField(&(externalparticipantshape_.Participant), externalparticipantshapeFormCallback.probe.stageOfInterest, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(externalparticipantshape_.IsExpanded), formDiv)
		case "X":
			FormDivBasicFieldToField(&(externalparticipantshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(externalparticipantshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(externalparticipantshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(externalparticipantshape_.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(externalparticipantshape_.IsHidden), formDiv)
		case "TailHeigth":
			FormDivBasicFieldToField(&(externalparticipantshape_.TailHeigth), formDiv)
		case "DiagramProcess:ExternalParticipant_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramProcess instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramProcess instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramProcess](externalparticipantshapeFormCallback.probe.stageOfInterest)
			targetDiagramProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramProcess instances and update their ExternalParticipant_Shapes slice
			for _diagramprocess := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramProcess](externalparticipantshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(externalparticipantshapeFormCallback.probe.stageOfInterest, _diagramprocess)
				
				// if DiagramProcess is selected
				if targetDiagramProcessIDs[id] {
					// ensure externalparticipantshape_ is in _diagramprocess.ExternalParticipant_Shapes
					found := false
					for _, _b := range _diagramprocess.ExternalParticipant_Shapes {
						if _b == externalparticipantshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramprocess.ExternalParticipant_Shapes = append(_diagramprocess.ExternalParticipant_Shapes, externalparticipantshape_)
						externalparticipantshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "ExternalParticipant_Shapes", &_diagramprocess.ExternalParticipant_Shapes)
					}
				} else {
					// ensure externalparticipantshape_ is NOT in _diagramprocess.ExternalParticipant_Shapes
					idx := slices.Index(_diagramprocess.ExternalParticipant_Shapes, externalparticipantshape_)
					if idx != -1 {
						_diagramprocess.ExternalParticipant_Shapes = slices.Delete(_diagramprocess.ExternalParticipant_Shapes, idx, idx+1)
						externalparticipantshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "ExternalParticipant_Shapes", &_diagramprocess.ExternalParticipant_Shapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if externalparticipantshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		externalparticipantshape_.Unstage(externalparticipantshapeFormCallback.probe.stageOfInterest)
	}

	externalparticipantshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ExternalParticipantShape](
		externalparticipantshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if externalparticipantshapeFormCallback.CreationMode || externalparticipantshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		externalparticipantshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(externalparticipantshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ExternalParticipantShapeFormCallback(
			nil,
			externalparticipantshapeFormCallback.probe,
			newFormGroup,
		)
		externalparticipantshape := new(models.ExternalParticipantShape)
		FillUpForm(externalparticipantshape, newFormGroup, externalparticipantshapeFormCallback.probe)
		externalparticipantshapeFormCallback.probe.formStage.Commit()
	}

	externalparticipantshapeFormCallback.probe.ux_tree()
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
		case "Description":
			FormDivBasicFieldToField(&(library_.Description), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(library_.ComputedPrefix), formDiv)
		case "IsRootLibrary":
			FormDivBasicFieldToField(&(library_.IsRootLibrary), formDiv)
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
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "SubLibraries", &library_.SubLibraries)

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
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "SubLibrariesWhoseNodeIsExpanded", &library_.SubLibrariesWhoseNodeIsExpanded)

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
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "RootProcesses", &library_.RootProcesses)

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
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "ProcesssWhoseNodeIsExpanded", &library_.ProcesssWhoseNodeIsExpanded)

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
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "RootDataFlows", &library_.RootDataFlows)

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
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "DataFlowsWhoseNodeIsExpanded", &library_.DataFlowsWhoseNodeIsExpanded)

		case "RootDatas":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Data](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Data, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Data)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Data](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.RootDatas = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "RootDatas", &library_.RootDatas)

		case "IsDatasNodeExpanded":
			FormDivBasicFieldToField(&(library_.IsDatasNodeExpanded), formDiv)
		case "DatasWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Data](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Data, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Data)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Data](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.DatasWhoseNodeIsExpanded = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "DatasWhoseNodeIsExpanded", &library_.DatasWhoseNodeIsExpanded)

		case "RootResources":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Resource](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Resource, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Resource)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Resource](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.RootResources = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "RootResources", &library_.RootResources)

		case "IsResourcesNodeExpanded":
			FormDivBasicFieldToField(&(library_.IsResourcesNodeExpanded), formDiv)
		case "ResourcesWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Resource](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Resource, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Resource)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Resource](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.ResourcesWhoseNodeIsExpanded = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "ResourcesWhoseNodeIsExpanded", &library_.ResourcesWhoseNodeIsExpanded)

		case "ParticipantsWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Participant](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Participant, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Participant)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Participant](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.ParticipantsWhoseNodeIsExpanded = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "ParticipantsWhoseNodeIsExpanded", &library_.ParticipantsWhoseNodeIsExpanded)

		case "RootNotes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Note](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Note, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Note)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Note](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.RootNotes = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "RootNotes", &library_.RootNotes)

		case "IsNotesNodeExpanded":
			FormDivBasicFieldToField(&(library_.IsNotesNodeExpanded), formDiv)
		case "NotesWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Note](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Note, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Note)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Note](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.NotesWhoseNodeIsExpanded = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "NotesWhoseNodeIsExpanded", &library_.NotesWhoseNodeIsExpanded)

		case "IsExpandedTmp":
			FormDivBasicFieldToField(&(library_.IsExpandedTmp), formDiv)
		case "Library:SubLibraries":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](libraryFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their SubLibraries slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](libraryFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(libraryFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure library_ is in _library.SubLibraries
					found := false
					for _, _b := range _library.SubLibraries {
						if _b == library_ {
							found = true
							break
						}
					}
					if !found {
						_library.SubLibraries = append(_library.SubLibraries, library_)
						libraryFormCallback.probe.UpdateSliceOfPointersCallback(_library, "SubLibraries", &_library.SubLibraries)
					}
				} else {
					// ensure library_ is NOT in _library.SubLibraries
					idx := slices.Index(_library.SubLibraries, library_)
					if idx != -1 {
						_library.SubLibraries = slices.Delete(_library.SubLibraries, idx, idx+1)
						libraryFormCallback.probe.UpdateSliceOfPointersCallback(_library, "SubLibraries", &_library.SubLibraries)
					}
				}
			}
		case "Library:SubLibrariesWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](libraryFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their SubLibrariesWhoseNodeIsExpanded slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](libraryFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(libraryFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure library_ is in _library.SubLibrariesWhoseNodeIsExpanded
					found := false
					for _, _b := range _library.SubLibrariesWhoseNodeIsExpanded {
						if _b == library_ {
							found = true
							break
						}
					}
					if !found {
						_library.SubLibrariesWhoseNodeIsExpanded = append(_library.SubLibrariesWhoseNodeIsExpanded, library_)
						libraryFormCallback.probe.UpdateSliceOfPointersCallback(_library, "SubLibrariesWhoseNodeIsExpanded", &_library.SubLibrariesWhoseNodeIsExpanded)
					}
				} else {
					// ensure library_ is NOT in _library.SubLibrariesWhoseNodeIsExpanded
					idx := slices.Index(_library.SubLibrariesWhoseNodeIsExpanded, library_)
					if idx != -1 {
						_library.SubLibrariesWhoseNodeIsExpanded = slices.Delete(_library.SubLibrariesWhoseNodeIsExpanded, idx, idx+1)
						libraryFormCallback.probe.UpdateSliceOfPointersCallback(_library, "SubLibrariesWhoseNodeIsExpanded", &_library.SubLibrariesWhoseNodeIsExpanded)
					}
				}
			}
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
func __gong__New__NoteFormCallback(
	note *models.Note,
	probe *Probe,
	formGroup *form.FormGroup,
) (noteFormCallback *NoteFormCallback) {
	noteFormCallback = new(NoteFormCallback)
	noteFormCallback.probe = probe
	noteFormCallback.note = note
	noteFormCallback.formGroup = formGroup

	noteFormCallback.CreationMode = (note == nil)

	return
}

type NoteFormCallback struct {
	note *models.Note

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (noteFormCallback *NoteFormCallback) OnSave() {
	noteFormCallback.probe.stageOfInterest.Lock()
	defer noteFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("NoteFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	noteFormCallback.probe.formStage.Checkout()

	if noteFormCallback.note == nil {
		noteFormCallback.note = new(models.Note).Stage(noteFormCallback.probe.stageOfInterest)
	}
	note_ := noteFormCallback.note
	_ = note_

	for _, formDiv := range noteFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(note_.Name), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(note_.Description), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(note_.ComputedPrefix), formDiv)
		case "IsTasksNodeExpanded":
			FormDivBasicFieldToField(&(note_.IsTasksNodeExpanded), formDiv)
		case "Tasks":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Task](noteFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Task, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Task)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					noteFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Task](noteFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			note_.Tasks = instanceSlice
			noteFormCallback.probe.UpdateSliceOfPointersCallback(note_, "Tasks", &note_.Tasks)

		case "DiagramProcess:NotesWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramProcess instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramProcess instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramProcess](noteFormCallback.probe.stageOfInterest)
			targetDiagramProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramProcess instances and update their NotesWhoseNodeIsExpanded slice
			for _diagramprocess := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramProcess](noteFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(noteFormCallback.probe.stageOfInterest, _diagramprocess)
				
				// if DiagramProcess is selected
				if targetDiagramProcessIDs[id] {
					// ensure note_ is in _diagramprocess.NotesWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramprocess.NotesWhoseNodeIsExpanded {
						if _b == note_ {
							found = true
							break
						}
					}
					if !found {
						_diagramprocess.NotesWhoseNodeIsExpanded = append(_diagramprocess.NotesWhoseNodeIsExpanded, note_)
						noteFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "NotesWhoseNodeIsExpanded", &_diagramprocess.NotesWhoseNodeIsExpanded)
					}
				} else {
					// ensure note_ is NOT in _diagramprocess.NotesWhoseNodeIsExpanded
					idx := slices.Index(_diagramprocess.NotesWhoseNodeIsExpanded, note_)
					if idx != -1 {
						_diagramprocess.NotesWhoseNodeIsExpanded = slices.Delete(_diagramprocess.NotesWhoseNodeIsExpanded, idx, idx+1)
						noteFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "NotesWhoseNodeIsExpanded", &_diagramprocess.NotesWhoseNodeIsExpanded)
					}
				}
			}
		case "Library:RootNotes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](noteFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their RootNotes slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](noteFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(noteFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure note_ is in _library.RootNotes
					found := false
					for _, _b := range _library.RootNotes {
						if _b == note_ {
							found = true
							break
						}
					}
					if !found {
						_library.RootNotes = append(_library.RootNotes, note_)
						noteFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootNotes", &_library.RootNotes)
					}
				} else {
					// ensure note_ is NOT in _library.RootNotes
					idx := slices.Index(_library.RootNotes, note_)
					if idx != -1 {
						_library.RootNotes = slices.Delete(_library.RootNotes, idx, idx+1)
						noteFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootNotes", &_library.RootNotes)
					}
				}
			}
		case "Library:NotesWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](noteFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their NotesWhoseNodeIsExpanded slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](noteFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(noteFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure note_ is in _library.NotesWhoseNodeIsExpanded
					found := false
					for _, _b := range _library.NotesWhoseNodeIsExpanded {
						if _b == note_ {
							found = true
							break
						}
					}
					if !found {
						_library.NotesWhoseNodeIsExpanded = append(_library.NotesWhoseNodeIsExpanded, note_)
						noteFormCallback.probe.UpdateSliceOfPointersCallback(_library, "NotesWhoseNodeIsExpanded", &_library.NotesWhoseNodeIsExpanded)
					}
				} else {
					// ensure note_ is NOT in _library.NotesWhoseNodeIsExpanded
					idx := slices.Index(_library.NotesWhoseNodeIsExpanded, note_)
					if idx != -1 {
						_library.NotesWhoseNodeIsExpanded = slices.Delete(_library.NotesWhoseNodeIsExpanded, idx, idx+1)
						noteFormCallback.probe.UpdateSliceOfPointersCallback(_library, "NotesWhoseNodeIsExpanded", &_library.NotesWhoseNodeIsExpanded)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if noteFormCallback.formGroup.HasSuppressButtonBeenPressed {
		note_.Unstage(noteFormCallback.probe.stageOfInterest)
	}

	noteFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Note](
		noteFormCallback.probe,
	)

	// display a new form by reset the form stage
	if noteFormCallback.CreationMode || noteFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(noteFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__NoteFormCallback(
			nil,
			noteFormCallback.probe,
			newFormGroup,
		)
		note := new(models.Note)
		FillUpForm(note, newFormGroup, noteFormCallback.probe)
		noteFormCallback.probe.formStage.Commit()
	}

	noteFormCallback.probe.ux_tree()
}
func __gong__New__NoteShapeFormCallback(
	noteshape *models.NoteShape,
	probe *Probe,
	formGroup *form.FormGroup,
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

	formGroup *form.FormGroup
}

func (noteshapeFormCallback *NoteShapeFormCallback) OnSave() {
	noteshapeFormCallback.probe.stageOfInterest.Lock()
	defer noteshapeFormCallback.probe.stageOfInterest.Unlock()

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
		case "Note":
			FormDivSelectFieldToField(&(noteshape_.Note), noteshapeFormCallback.probe.stageOfInterest, formDiv)
		case "X":
			FormDivBasicFieldToField(&(noteshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(noteshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(noteshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(noteshape_.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(noteshape_.IsHidden), formDiv)
		case "DiagramProcess:Note_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramProcess instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramProcess instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramProcess](noteshapeFormCallback.probe.stageOfInterest)
			targetDiagramProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramProcess instances and update their Note_Shapes slice
			for _diagramprocess := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramProcess](noteshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(noteshapeFormCallback.probe.stageOfInterest, _diagramprocess)
				
				// if DiagramProcess is selected
				if targetDiagramProcessIDs[id] {
					// ensure noteshape_ is in _diagramprocess.Note_Shapes
					found := false
					for _, _b := range _diagramprocess.Note_Shapes {
						if _b == noteshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramprocess.Note_Shapes = append(_diagramprocess.Note_Shapes, noteshape_)
						noteshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "Note_Shapes", &_diagramprocess.Note_Shapes)
					}
				} else {
					// ensure noteshape_ is NOT in _diagramprocess.Note_Shapes
					idx := slices.Index(_diagramprocess.Note_Shapes, noteshape_)
					if idx != -1 {
						_diagramprocess.Note_Shapes = slices.Delete(_diagramprocess.Note_Shapes, idx, idx+1)
						noteshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "Note_Shapes", &_diagramprocess.Note_Shapes)
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
	updateProbeTable[*models.NoteShape](
		noteshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if noteshapeFormCallback.CreationMode || noteshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
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

	noteshapeFormCallback.probe.ux_tree()
}
func __gong__New__NoteTaskShapeFormCallback(
	notetaskshape *models.NoteTaskShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (notetaskshapeFormCallback *NoteTaskShapeFormCallback) {
	notetaskshapeFormCallback = new(NoteTaskShapeFormCallback)
	notetaskshapeFormCallback.probe = probe
	notetaskshapeFormCallback.notetaskshape = notetaskshape
	notetaskshapeFormCallback.formGroup = formGroup

	notetaskshapeFormCallback.CreationMode = (notetaskshape == nil)

	return
}

type NoteTaskShapeFormCallback struct {
	notetaskshape *models.NoteTaskShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (notetaskshapeFormCallback *NoteTaskShapeFormCallback) OnSave() {
	notetaskshapeFormCallback.probe.stageOfInterest.Lock()
	defer notetaskshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("NoteTaskShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	notetaskshapeFormCallback.probe.formStage.Checkout()

	if notetaskshapeFormCallback.notetaskshape == nil {
		notetaskshapeFormCallback.notetaskshape = new(models.NoteTaskShape).Stage(notetaskshapeFormCallback.probe.stageOfInterest)
	}
	notetaskshape_ := notetaskshapeFormCallback.notetaskshape
	_ = notetaskshape_

	for _, formDiv := range notetaskshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(notetaskshape_.Name), formDiv)
		case "Note":
			FormDivSelectFieldToField(&(notetaskshape_.Note), notetaskshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Task":
			FormDivSelectFieldToField(&(notetaskshape_.Task), notetaskshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(notetaskshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(notetaskshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(notetaskshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(notetaskshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(notetaskshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(notetaskshape_.IsHidden), formDiv)
		case "DiagramProcess:NoteTaskShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramProcess instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramProcess instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramProcess](notetaskshapeFormCallback.probe.stageOfInterest)
			targetDiagramProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramProcess instances and update their NoteTaskShapes slice
			for _diagramprocess := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramProcess](notetaskshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(notetaskshapeFormCallback.probe.stageOfInterest, _diagramprocess)
				
				// if DiagramProcess is selected
				if targetDiagramProcessIDs[id] {
					// ensure notetaskshape_ is in _diagramprocess.NoteTaskShapes
					found := false
					for _, _b := range _diagramprocess.NoteTaskShapes {
						if _b == notetaskshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramprocess.NoteTaskShapes = append(_diagramprocess.NoteTaskShapes, notetaskshape_)
						notetaskshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "NoteTaskShapes", &_diagramprocess.NoteTaskShapes)
					}
				} else {
					// ensure notetaskshape_ is NOT in _diagramprocess.NoteTaskShapes
					idx := slices.Index(_diagramprocess.NoteTaskShapes, notetaskshape_)
					if idx != -1 {
						_diagramprocess.NoteTaskShapes = slices.Delete(_diagramprocess.NoteTaskShapes, idx, idx+1)
						notetaskshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "NoteTaskShapes", &_diagramprocess.NoteTaskShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if notetaskshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		notetaskshape_.Unstage(notetaskshapeFormCallback.probe.stageOfInterest)
	}

	notetaskshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.NoteTaskShape](
		notetaskshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if notetaskshapeFormCallback.CreationMode || notetaskshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		notetaskshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(notetaskshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__NoteTaskShapeFormCallback(
			nil,
			notetaskshapeFormCallback.probe,
			newFormGroup,
		)
		notetaskshape := new(models.NoteTaskShape)
		FillUpForm(notetaskshape, newFormGroup, notetaskshapeFormCallback.probe)
		notetaskshapeFormCallback.probe.formStage.Commit()
	}

	notetaskshapeFormCallback.probe.ux_tree()
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
		case "IsProcessResource":
			FormDivBasicFieldToField(&(participant_.IsProcessResource), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(participant_.Description), formDiv)
		case "Resources":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Resource](participantFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Resource, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Resource)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Resource](participantFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			participant_.Resources = instanceSlice
			participantFormCallback.probe.UpdateSliceOfPointersCallback(participant_, "Resources", &participant_.Resources)

		case "IsResourcesNodeExpanded":
			FormDivBasicFieldToField(&(participant_.IsResourcesNodeExpanded), formDiv)
		case "Processes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Process](participantFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Process, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Process)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Process](participantFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			participant_.Processes = instanceSlice
			participantFormCallback.probe.UpdateSliceOfPointersCallback(participant_, "Processes", &participant_.Processes)

		case "IsProcessesNodeExpanded":
			FormDivBasicFieldToField(&(participant_.IsProcessesNodeExpanded), formDiv)
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
			participantFormCallback.probe.UpdateSliceOfPointersCallback(participant_, "Tasks", &participant_.Tasks)

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
			participantFormCallback.probe.UpdateSliceOfPointersCallback(participant_, "ControlFlows", &participant_.ControlFlows)

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
			participantFormCallback.probe.UpdateSliceOfPointersCallback(participant_, "TaskWhoseOutControlFlowsNodeIsExpanded", &participant_.TaskWhoseOutControlFlowsNodeIsExpanded)

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
			participantFormCallback.probe.UpdateSliceOfPointersCallback(participant_, "TaskWhoseInControlFlowsNodeIsExpanded", &participant_.TaskWhoseInControlFlowsNodeIsExpanded)

		case "IsDataFlowsNodeExpanded":
			FormDivBasicFieldToField(&(participant_.IsDataFlowsNodeExpanded), formDiv)
		case "TaskWhoseOutDataFlowsNodeIsExpanded":
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
			participant_.TaskWhoseOutDataFlowsNodeIsExpanded = instanceSlice
			participantFormCallback.probe.UpdateSliceOfPointersCallback(participant_, "TaskWhoseOutDataFlowsNodeIsExpanded", &participant_.TaskWhoseOutDataFlowsNodeIsExpanded)

		case "TaskWhoseInDataFlowsNodeIsExpanded":
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
			participant_.TaskWhoseInDataFlowsNodeIsExpanded = instanceSlice
			participantFormCallback.probe.UpdateSliceOfPointersCallback(participant_, "TaskWhoseInDataFlowsNodeIsExpanded", &participant_.TaskWhoseInDataFlowsNodeIsExpanded)

		case "DiagramProcess:ParticipantWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramProcess instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramProcess instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramProcess](participantFormCallback.probe.stageOfInterest)
			targetDiagramProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramProcess instances and update their ParticipantWhoseNodeIsExpanded slice
			for _diagramprocess := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramProcess](participantFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(participantFormCallback.probe.stageOfInterest, _diagramprocess)
				
				// if DiagramProcess is selected
				if targetDiagramProcessIDs[id] {
					// ensure participant_ is in _diagramprocess.ParticipantWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramprocess.ParticipantWhoseNodeIsExpanded {
						if _b == participant_ {
							found = true
							break
						}
					}
					if !found {
						_diagramprocess.ParticipantWhoseNodeIsExpanded = append(_diagramprocess.ParticipantWhoseNodeIsExpanded, participant_)
						participantFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "ParticipantWhoseNodeIsExpanded", &_diagramprocess.ParticipantWhoseNodeIsExpanded)
					}
				} else {
					// ensure participant_ is NOT in _diagramprocess.ParticipantWhoseNodeIsExpanded
					idx := slices.Index(_diagramprocess.ParticipantWhoseNodeIsExpanded, participant_)
					if idx != -1 {
						_diagramprocess.ParticipantWhoseNodeIsExpanded = slices.Delete(_diagramprocess.ParticipantWhoseNodeIsExpanded, idx, idx+1)
						participantFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "ParticipantWhoseNodeIsExpanded", &_diagramprocess.ParticipantWhoseNodeIsExpanded)
					}
				}
			}
		case "DiagramProcess:ExternalParticipantWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramProcess instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramProcess instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramProcess](participantFormCallback.probe.stageOfInterest)
			targetDiagramProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramProcess instances and update their ExternalParticipantWhoseNodeIsExpanded slice
			for _diagramprocess := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramProcess](participantFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(participantFormCallback.probe.stageOfInterest, _diagramprocess)
				
				// if DiagramProcess is selected
				if targetDiagramProcessIDs[id] {
					// ensure participant_ is in _diagramprocess.ExternalParticipantWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramprocess.ExternalParticipantWhoseNodeIsExpanded {
						if _b == participant_ {
							found = true
							break
						}
					}
					if !found {
						_diagramprocess.ExternalParticipantWhoseNodeIsExpanded = append(_diagramprocess.ExternalParticipantWhoseNodeIsExpanded, participant_)
						participantFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "ExternalParticipantWhoseNodeIsExpanded", &_diagramprocess.ExternalParticipantWhoseNodeIsExpanded)
					}
				} else {
					// ensure participant_ is NOT in _diagramprocess.ExternalParticipantWhoseNodeIsExpanded
					idx := slices.Index(_diagramprocess.ExternalParticipantWhoseNodeIsExpanded, participant_)
					if idx != -1 {
						_diagramprocess.ExternalParticipantWhoseNodeIsExpanded = slices.Delete(_diagramprocess.ExternalParticipantWhoseNodeIsExpanded, idx, idx+1)
						participantFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "ExternalParticipantWhoseNodeIsExpanded", &_diagramprocess.ExternalParticipantWhoseNodeIsExpanded)
					}
				}
			}
		case "DiagramProcess:ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramProcess instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramProcess instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramProcess](participantFormCallback.probe.stageOfInterest)
			targetDiagramProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramProcess instances and update their ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded slice
			for _diagramprocess := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramProcess](participantFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(participantFormCallback.probe.stageOfInterest, _diagramprocess)
				
				// if DiagramProcess is selected
				if targetDiagramProcessIDs[id] {
					// ensure participant_ is in _diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded
					found := false
					for _, _b := range _diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded {
						if _b == participant_ {
							found = true
							break
						}
					}
					if !found {
						_diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded = append(_diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded, participant_)
						participantFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded", &_diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded)
					}
				} else {
					// ensure participant_ is NOT in _diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded
					idx := slices.Index(_diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded, participant_)
					if idx != -1 {
						_diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded = slices.Delete(_diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded, idx, idx+1)
						participantFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded", &_diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded)
					}
				}
			}
		case "DiagramProcess:ExternalParticipantsWhoseInDataFlowsNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramProcess instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramProcess instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramProcess](participantFormCallback.probe.stageOfInterest)
			targetDiagramProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramProcess instances and update their ExternalParticipantsWhoseInDataFlowsNodeIsExpanded slice
			for _diagramprocess := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramProcess](participantFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(participantFormCallback.probe.stageOfInterest, _diagramprocess)
				
				// if DiagramProcess is selected
				if targetDiagramProcessIDs[id] {
					// ensure participant_ is in _diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded
					found := false
					for _, _b := range _diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded {
						if _b == participant_ {
							found = true
							break
						}
					}
					if !found {
						_diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded = append(_diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded, participant_)
						participantFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "ExternalParticipantsWhoseInDataFlowsNodeIsExpanded", &_diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded)
					}
				} else {
					// ensure participant_ is NOT in _diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded
					idx := slices.Index(_diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded, participant_)
					if idx != -1 {
						_diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded = slices.Delete(_diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded, idx, idx+1)
						participantFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "ExternalParticipantsWhoseInDataFlowsNodeIsExpanded", &_diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded)
					}
				}
			}
		case "Library:ParticipantsWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](participantFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their ParticipantsWhoseNodeIsExpanded slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](participantFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(participantFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure participant_ is in _library.ParticipantsWhoseNodeIsExpanded
					found := false
					for _, _b := range _library.ParticipantsWhoseNodeIsExpanded {
						if _b == participant_ {
							found = true
							break
						}
					}
					if !found {
						_library.ParticipantsWhoseNodeIsExpanded = append(_library.ParticipantsWhoseNodeIsExpanded, participant_)
						participantFormCallback.probe.UpdateSliceOfPointersCallback(_library, "ParticipantsWhoseNodeIsExpanded", &_library.ParticipantsWhoseNodeIsExpanded)
					}
				} else {
					// ensure participant_ is NOT in _library.ParticipantsWhoseNodeIsExpanded
					idx := slices.Index(_library.ParticipantsWhoseNodeIsExpanded, participant_)
					if idx != -1 {
						_library.ParticipantsWhoseNodeIsExpanded = slices.Delete(_library.ParticipantsWhoseNodeIsExpanded, idx, idx+1)
						participantFormCallback.probe.UpdateSliceOfPointersCallback(_library, "ParticipantsWhoseNodeIsExpanded", &_library.ParticipantsWhoseNodeIsExpanded)
					}
				}
			}
		case "Process:Participants":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Process instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Process instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Process](participantFormCallback.probe.stageOfInterest)
			targetProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Process instances and update their Participants slice
			for _process := range *models.GetGongstructInstancesSetFromPointerType[*models.Process](participantFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(participantFormCallback.probe.stageOfInterest, _process)
				
				// if Process is selected
				if targetProcessIDs[id] {
					// ensure participant_ is in _process.Participants
					found := false
					for _, _b := range _process.Participants {
						if _b == participant_ {
							found = true
							break
						}
					}
					if !found {
						_process.Participants = append(_process.Participants, participant_)
						participantFormCallback.probe.UpdateSliceOfPointersCallback(_process, "Participants", &_process.Participants)
					}
				} else {
					// ensure participant_ is NOT in _process.Participants
					idx := slices.Index(_process.Participants, participant_)
					if idx != -1 {
						_process.Participants = slices.Delete(_process.Participants, idx, idx+1)
						participantFormCallback.probe.UpdateSliceOfPointersCallback(_process, "Participants", &_process.Participants)
					}
				}
			}
		case "Process:ParticipantWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Process instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Process instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Process](participantFormCallback.probe.stageOfInterest)
			targetProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Process instances and update their ParticipantWhoseNodeIsExpanded slice
			for _process := range *models.GetGongstructInstancesSetFromPointerType[*models.Process](participantFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(participantFormCallback.probe.stageOfInterest, _process)
				
				// if Process is selected
				if targetProcessIDs[id] {
					// ensure participant_ is in _process.ParticipantWhoseNodeIsExpanded
					found := false
					for _, _b := range _process.ParticipantWhoseNodeIsExpanded {
						if _b == participant_ {
							found = true
							break
						}
					}
					if !found {
						_process.ParticipantWhoseNodeIsExpanded = append(_process.ParticipantWhoseNodeIsExpanded, participant_)
						participantFormCallback.probe.UpdateSliceOfPointersCallback(_process, "ParticipantWhoseNodeIsExpanded", &_process.ParticipantWhoseNodeIsExpanded)
					}
				} else {
					// ensure participant_ is NOT in _process.ParticipantWhoseNodeIsExpanded
					idx := slices.Index(_process.ParticipantWhoseNodeIsExpanded, participant_)
					if idx != -1 {
						_process.ParticipantWhoseNodeIsExpanded = slices.Delete(_process.ParticipantWhoseNodeIsExpanded, idx, idx+1)
						participantFormCallback.probe.UpdateSliceOfPointersCallback(_process, "ParticipantWhoseNodeIsExpanded", &_process.ParticipantWhoseNodeIsExpanded)
					}
				}
			}
		case "Process:ExternalParticipants":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Process instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Process instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Process](participantFormCallback.probe.stageOfInterest)
			targetProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Process instances and update their ExternalParticipants slice
			for _process := range *models.GetGongstructInstancesSetFromPointerType[*models.Process](participantFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(participantFormCallback.probe.stageOfInterest, _process)
				
				// if Process is selected
				if targetProcessIDs[id] {
					// ensure participant_ is in _process.ExternalParticipants
					found := false
					for _, _b := range _process.ExternalParticipants {
						if _b == participant_ {
							found = true
							break
						}
					}
					if !found {
						_process.ExternalParticipants = append(_process.ExternalParticipants, participant_)
						participantFormCallback.probe.UpdateSliceOfPointersCallback(_process, "ExternalParticipants", &_process.ExternalParticipants)
					}
				} else {
					// ensure participant_ is NOT in _process.ExternalParticipants
					idx := slices.Index(_process.ExternalParticipants, participant_)
					if idx != -1 {
						_process.ExternalParticipants = slices.Delete(_process.ExternalParticipants, idx, idx+1)
						participantFormCallback.probe.UpdateSliceOfPointersCallback(_process, "ExternalParticipants", &_process.ExternalParticipants)
					}
				}
			}
		case "Process:ExternalParticipantWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Process instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Process instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Process](participantFormCallback.probe.stageOfInterest)
			targetProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Process instances and update their ExternalParticipantWhoseNodeIsExpanded slice
			for _process := range *models.GetGongstructInstancesSetFromPointerType[*models.Process](participantFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(participantFormCallback.probe.stageOfInterest, _process)
				
				// if Process is selected
				if targetProcessIDs[id] {
					// ensure participant_ is in _process.ExternalParticipantWhoseNodeIsExpanded
					found := false
					for _, _b := range _process.ExternalParticipantWhoseNodeIsExpanded {
						if _b == participant_ {
							found = true
							break
						}
					}
					if !found {
						_process.ExternalParticipantWhoseNodeIsExpanded = append(_process.ExternalParticipantWhoseNodeIsExpanded, participant_)
						participantFormCallback.probe.UpdateSliceOfPointersCallback(_process, "ExternalParticipantWhoseNodeIsExpanded", &_process.ExternalParticipantWhoseNodeIsExpanded)
					}
				} else {
					// ensure participant_ is NOT in _process.ExternalParticipantWhoseNodeIsExpanded
					idx := slices.Index(_process.ExternalParticipantWhoseNodeIsExpanded, participant_)
					if idx != -1 {
						_process.ExternalParticipantWhoseNodeIsExpanded = slices.Delete(_process.ExternalParticipantWhoseNodeIsExpanded, idx, idx+1)
						participantFormCallback.probe.UpdateSliceOfPointersCallback(_process, "ExternalParticipantWhoseNodeIsExpanded", &_process.ExternalParticipantWhoseNodeIsExpanded)
					}
				}
			}
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
		case "WidthWeight":
			FormDivBasicFieldToField(&(participantshape_.WidthWeight), formDiv)
		case "DiagramProcess:Participant_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramProcess instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramProcess instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramProcess](participantshapeFormCallback.probe.stageOfInterest)
			targetDiagramProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramProcess instances and update their Participant_Shapes slice
			for _diagramprocess := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramProcess](participantshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(participantshapeFormCallback.probe.stageOfInterest, _diagramprocess)
				
				// if DiagramProcess is selected
				if targetDiagramProcessIDs[id] {
					// ensure participantshape_ is in _diagramprocess.Participant_Shapes
					found := false
					for _, _b := range _diagramprocess.Participant_Shapes {
						if _b == participantshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramprocess.Participant_Shapes = append(_diagramprocess.Participant_Shapes, participantshape_)
						participantshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "Participant_Shapes", &_diagramprocess.Participant_Shapes)
					}
				} else {
					// ensure participantshape_ is NOT in _diagramprocess.Participant_Shapes
					idx := slices.Index(_diagramprocess.Participant_Shapes, participantshape_)
					if idx != -1 {
						_diagramprocess.Participant_Shapes = slices.Delete(_diagramprocess.Participant_Shapes, idx, idx+1)
						participantshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "Participant_Shapes", &_diagramprocess.Participant_Shapes)
					}
				}
			}
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
		case "Description":
			FormDivBasicFieldToField(&(process_.Description), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(process_.ComputedPrefix), formDiv)
		case "SVG_Path":
			FormDivBasicFieldToField(&(process_.SVG_Path), formDiv)
		case "InverseAppliedScaling":
			FormDivBasicFieldToField(&(process_.InverseAppliedScaling), formDiv)
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
			processFormCallback.probe.UpdateSliceOfPointersCallback(process_, "DiagramProcesss", &process_.DiagramProcesss)

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
			processFormCallback.probe.UpdateSliceOfPointersCallback(process_, "DiagramProcessWhoseNodeIsExpanded", &process_.DiagramProcessWhoseNodeIsExpanded)

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
			processFormCallback.probe.UpdateSliceOfPointersCallback(process_, "SubProcesses", &process_.SubProcesses)

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
			processFormCallback.probe.UpdateSliceOfPointersCallback(process_, "Participants", &process_.Participants)

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
			processFormCallback.probe.UpdateSliceOfPointersCallback(process_, "ParticipantWhoseNodeIsExpanded", &process_.ParticipantWhoseNodeIsExpanded)

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
			processFormCallback.probe.UpdateSliceOfPointersCallback(process_, "DataFlows", &process_.DataFlows)

		case "IsDataFlowsNodeExpanded":
			FormDivBasicFieldToField(&(process_.IsDataFlowsNodeExpanded), formDiv)
		case "ExternalParticipants":
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
			process_.ExternalParticipants = instanceSlice
			processFormCallback.probe.UpdateSliceOfPointersCallback(process_, "ExternalParticipants", &process_.ExternalParticipants)

		case "ExternalParticipantWhoseNodeIsExpanded":
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
			process_.ExternalParticipantWhoseNodeIsExpanded = instanceSlice
			processFormCallback.probe.UpdateSliceOfPointersCallback(process_, "ExternalParticipantWhoseNodeIsExpanded", &process_.ExternalParticipantWhoseNodeIsExpanded)

		case "DiagramProcess:ProcesssWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramProcess instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramProcess instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramProcess](processFormCallback.probe.stageOfInterest)
			targetDiagramProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramProcess instances and update their ProcesssWhoseNodeIsExpanded slice
			for _diagramprocess := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramProcess](processFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(processFormCallback.probe.stageOfInterest, _diagramprocess)
				
				// if DiagramProcess is selected
				if targetDiagramProcessIDs[id] {
					// ensure process_ is in _diagramprocess.ProcesssWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramprocess.ProcesssWhoseNodeIsExpanded {
						if _b == process_ {
							found = true
							break
						}
					}
					if !found {
						_diagramprocess.ProcesssWhoseNodeIsExpanded = append(_diagramprocess.ProcesssWhoseNodeIsExpanded, process_)
						processFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "ProcesssWhoseNodeIsExpanded", &_diagramprocess.ProcesssWhoseNodeIsExpanded)
					}
				} else {
					// ensure process_ is NOT in _diagramprocess.ProcesssWhoseNodeIsExpanded
					idx := slices.Index(_diagramprocess.ProcesssWhoseNodeIsExpanded, process_)
					if idx != -1 {
						_diagramprocess.ProcesssWhoseNodeIsExpanded = slices.Delete(_diagramprocess.ProcesssWhoseNodeIsExpanded, idx, idx+1)
						processFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "ProcesssWhoseNodeIsExpanded", &_diagramprocess.ProcesssWhoseNodeIsExpanded)
					}
				}
			}
		case "DiagramProcess:AllocatedProcessesWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramProcess instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramProcess instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramProcess](processFormCallback.probe.stageOfInterest)
			targetDiagramProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramProcess instances and update their AllocatedProcessesWhoseNodeIsExpanded slice
			for _diagramprocess := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramProcess](processFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(processFormCallback.probe.stageOfInterest, _diagramprocess)
				
				// if DiagramProcess is selected
				if targetDiagramProcessIDs[id] {
					// ensure process_ is in _diagramprocess.AllocatedProcessesWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramprocess.AllocatedProcessesWhoseNodeIsExpanded {
						if _b == process_ {
							found = true
							break
						}
					}
					if !found {
						_diagramprocess.AllocatedProcessesWhoseNodeIsExpanded = append(_diagramprocess.AllocatedProcessesWhoseNodeIsExpanded, process_)
						processFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "AllocatedProcessesWhoseNodeIsExpanded", &_diagramprocess.AllocatedProcessesWhoseNodeIsExpanded)
					}
				} else {
					// ensure process_ is NOT in _diagramprocess.AllocatedProcessesWhoseNodeIsExpanded
					idx := slices.Index(_diagramprocess.AllocatedProcessesWhoseNodeIsExpanded, process_)
					if idx != -1 {
						_diagramprocess.AllocatedProcessesWhoseNodeIsExpanded = slices.Delete(_diagramprocess.AllocatedProcessesWhoseNodeIsExpanded, idx, idx+1)
						processFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "AllocatedProcessesWhoseNodeIsExpanded", &_diagramprocess.AllocatedProcessesWhoseNodeIsExpanded)
					}
				}
			}
		case "Library:RootProcesses":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](processFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their RootProcesses slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](processFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(processFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure process_ is in _library.RootProcesses
					found := false
					for _, _b := range _library.RootProcesses {
						if _b == process_ {
							found = true
							break
						}
					}
					if !found {
						_library.RootProcesses = append(_library.RootProcesses, process_)
						processFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootProcesses", &_library.RootProcesses)
					}
				} else {
					// ensure process_ is NOT in _library.RootProcesses
					idx := slices.Index(_library.RootProcesses, process_)
					if idx != -1 {
						_library.RootProcesses = slices.Delete(_library.RootProcesses, idx, idx+1)
						processFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootProcesses", &_library.RootProcesses)
					}
				}
			}
		case "Library:ProcesssWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](processFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their ProcesssWhoseNodeIsExpanded slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](processFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(processFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure process_ is in _library.ProcesssWhoseNodeIsExpanded
					found := false
					for _, _b := range _library.ProcesssWhoseNodeIsExpanded {
						if _b == process_ {
							found = true
							break
						}
					}
					if !found {
						_library.ProcesssWhoseNodeIsExpanded = append(_library.ProcesssWhoseNodeIsExpanded, process_)
						processFormCallback.probe.UpdateSliceOfPointersCallback(_library, "ProcesssWhoseNodeIsExpanded", &_library.ProcesssWhoseNodeIsExpanded)
					}
				} else {
					// ensure process_ is NOT in _library.ProcesssWhoseNodeIsExpanded
					idx := slices.Index(_library.ProcesssWhoseNodeIsExpanded, process_)
					if idx != -1 {
						_library.ProcesssWhoseNodeIsExpanded = slices.Delete(_library.ProcesssWhoseNodeIsExpanded, idx, idx+1)
						processFormCallback.probe.UpdateSliceOfPointersCallback(_library, "ProcesssWhoseNodeIsExpanded", &_library.ProcesssWhoseNodeIsExpanded)
					}
				}
			}
		case "Participant:Processes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Participant instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Participant instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Participant](processFormCallback.probe.stageOfInterest)
			targetParticipantIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetParticipantIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Participant instances and update their Processes slice
			for _participant := range *models.GetGongstructInstancesSetFromPointerType[*models.Participant](processFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(processFormCallback.probe.stageOfInterest, _participant)
				
				// if Participant is selected
				if targetParticipantIDs[id] {
					// ensure process_ is in _participant.Processes
					found := false
					for _, _b := range _participant.Processes {
						if _b == process_ {
							found = true
							break
						}
					}
					if !found {
						_participant.Processes = append(_participant.Processes, process_)
						processFormCallback.probe.UpdateSliceOfPointersCallback(_participant, "Processes", &_participant.Processes)
					}
				} else {
					// ensure process_ is NOT in _participant.Processes
					idx := slices.Index(_participant.Processes, process_)
					if idx != -1 {
						_participant.Processes = slices.Delete(_participant.Processes, idx, idx+1)
						processFormCallback.probe.UpdateSliceOfPointersCallback(_participant, "Processes", &_participant.Processes)
					}
				}
			}
		case "Process:SubProcesses":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Process instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Process instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Process](processFormCallback.probe.stageOfInterest)
			targetProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Process instances and update their SubProcesses slice
			for _process := range *models.GetGongstructInstancesSetFromPointerType[*models.Process](processFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(processFormCallback.probe.stageOfInterest, _process)
				
				// if Process is selected
				if targetProcessIDs[id] {
					// ensure process_ is in _process.SubProcesses
					found := false
					for _, _b := range _process.SubProcesses {
						if _b == process_ {
							found = true
							break
						}
					}
					if !found {
						_process.SubProcesses = append(_process.SubProcesses, process_)
						processFormCallback.probe.UpdateSliceOfPointersCallback(_process, "SubProcesses", &_process.SubProcesses)
					}
				} else {
					// ensure process_ is NOT in _process.SubProcesses
					idx := slices.Index(_process.SubProcesses, process_)
					if idx != -1 {
						_process.SubProcesses = slices.Delete(_process.SubProcesses, idx, idx+1)
						processFormCallback.probe.UpdateSliceOfPointersCallback(_process, "SubProcesses", &_process.SubProcesses)
					}
				}
			}
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
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramProcess instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramProcess instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramProcess](processshapeFormCallback.probe.stageOfInterest)
			targetDiagramProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramProcess instances and update their Process_Shapes slice
			for _diagramprocess := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramProcess](processshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(processshapeFormCallback.probe.stageOfInterest, _diagramprocess)
				
				// if DiagramProcess is selected
				if targetDiagramProcessIDs[id] {
					// ensure processshape_ is in _diagramprocess.Process_Shapes
					found := false
					for _, _b := range _diagramprocess.Process_Shapes {
						if _b == processshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramprocess.Process_Shapes = append(_diagramprocess.Process_Shapes, processshape_)
						processshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "Process_Shapes", &_diagramprocess.Process_Shapes)
					}
				} else {
					// ensure processshape_ is NOT in _diagramprocess.Process_Shapes
					idx := slices.Index(_diagramprocess.Process_Shapes, processshape_)
					if idx != -1 {
						_diagramprocess.Process_Shapes = slices.Delete(_diagramprocess.Process_Shapes, idx, idx+1)
						processshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "Process_Shapes", &_diagramprocess.Process_Shapes)
					}
				}
			}
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
func __gong__New__ResourceFormCallback(
	resource *models.Resource,
	probe *Probe,
	formGroup *form.FormGroup,
) (resourceFormCallback *ResourceFormCallback) {
	resourceFormCallback = new(ResourceFormCallback)
	resourceFormCallback.probe = probe
	resourceFormCallback.resource = resource
	resourceFormCallback.formGroup = formGroup

	resourceFormCallback.CreationMode = (resource == nil)

	return
}

type ResourceFormCallback struct {
	resource *models.Resource

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (resourceFormCallback *ResourceFormCallback) OnSave() {
	resourceFormCallback.probe.stageOfInterest.Lock()
	defer resourceFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ResourceFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	resourceFormCallback.probe.formStage.Checkout()

	if resourceFormCallback.resource == nil {
		resourceFormCallback.resource = new(models.Resource).Stage(resourceFormCallback.probe.stageOfInterest)
	}
	resource_ := resourceFormCallback.resource
	_ = resource_

	for _, formDiv := range resourceFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(resource_.Name), formDiv)
		case "Acronym":
			FormDivBasicFieldToField(&(resource_.Acronym), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(resource_.Description), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(resource_.ComputedPrefix), formDiv)
		case "SVG_Path":
			FormDivBasicFieldToField(&(resource_.SVG_Path), formDiv)
		case "InverseAppliedScaling":
			FormDivBasicFieldToField(&(resource_.InverseAppliedScaling), formDiv)
		case "DiagramProcess:AllocatedResourcesWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramProcess instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramProcess instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramProcess](resourceFormCallback.probe.stageOfInterest)
			targetDiagramProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramProcess instances and update their AllocatedResourcesWhoseNodeIsExpanded slice
			for _diagramprocess := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramProcess](resourceFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(resourceFormCallback.probe.stageOfInterest, _diagramprocess)
				
				// if DiagramProcess is selected
				if targetDiagramProcessIDs[id] {
					// ensure resource_ is in _diagramprocess.AllocatedResourcesWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramprocess.AllocatedResourcesWhoseNodeIsExpanded {
						if _b == resource_ {
							found = true
							break
						}
					}
					if !found {
						_diagramprocess.AllocatedResourcesWhoseNodeIsExpanded = append(_diagramprocess.AllocatedResourcesWhoseNodeIsExpanded, resource_)
						resourceFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "AllocatedResourcesWhoseNodeIsExpanded", &_diagramprocess.AllocatedResourcesWhoseNodeIsExpanded)
					}
				} else {
					// ensure resource_ is NOT in _diagramprocess.AllocatedResourcesWhoseNodeIsExpanded
					idx := slices.Index(_diagramprocess.AllocatedResourcesWhoseNodeIsExpanded, resource_)
					if idx != -1 {
						_diagramprocess.AllocatedResourcesWhoseNodeIsExpanded = slices.Delete(_diagramprocess.AllocatedResourcesWhoseNodeIsExpanded, idx, idx+1)
						resourceFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "AllocatedResourcesWhoseNodeIsExpanded", &_diagramprocess.AllocatedResourcesWhoseNodeIsExpanded)
					}
				}
			}
		case "Library:RootResources":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](resourceFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their RootResources slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](resourceFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(resourceFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure resource_ is in _library.RootResources
					found := false
					for _, _b := range _library.RootResources {
						if _b == resource_ {
							found = true
							break
						}
					}
					if !found {
						_library.RootResources = append(_library.RootResources, resource_)
						resourceFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootResources", &_library.RootResources)
					}
				} else {
					// ensure resource_ is NOT in _library.RootResources
					idx := slices.Index(_library.RootResources, resource_)
					if idx != -1 {
						_library.RootResources = slices.Delete(_library.RootResources, idx, idx+1)
						resourceFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootResources", &_library.RootResources)
					}
				}
			}
		case "Library:ResourcesWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](resourceFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their ResourcesWhoseNodeIsExpanded slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](resourceFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(resourceFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure resource_ is in _library.ResourcesWhoseNodeIsExpanded
					found := false
					for _, _b := range _library.ResourcesWhoseNodeIsExpanded {
						if _b == resource_ {
							found = true
							break
						}
					}
					if !found {
						_library.ResourcesWhoseNodeIsExpanded = append(_library.ResourcesWhoseNodeIsExpanded, resource_)
						resourceFormCallback.probe.UpdateSliceOfPointersCallback(_library, "ResourcesWhoseNodeIsExpanded", &_library.ResourcesWhoseNodeIsExpanded)
					}
				} else {
					// ensure resource_ is NOT in _library.ResourcesWhoseNodeIsExpanded
					idx := slices.Index(_library.ResourcesWhoseNodeIsExpanded, resource_)
					if idx != -1 {
						_library.ResourcesWhoseNodeIsExpanded = slices.Delete(_library.ResourcesWhoseNodeIsExpanded, idx, idx+1)
						resourceFormCallback.probe.UpdateSliceOfPointersCallback(_library, "ResourcesWhoseNodeIsExpanded", &_library.ResourcesWhoseNodeIsExpanded)
					}
				}
			}
		case "Participant:Resources":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Participant instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Participant instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Participant](resourceFormCallback.probe.stageOfInterest)
			targetParticipantIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetParticipantIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Participant instances and update their Resources slice
			for _participant := range *models.GetGongstructInstancesSetFromPointerType[*models.Participant](resourceFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(resourceFormCallback.probe.stageOfInterest, _participant)
				
				// if Participant is selected
				if targetParticipantIDs[id] {
					// ensure resource_ is in _participant.Resources
					found := false
					for _, _b := range _participant.Resources {
						if _b == resource_ {
							found = true
							break
						}
					}
					if !found {
						_participant.Resources = append(_participant.Resources, resource_)
						resourceFormCallback.probe.UpdateSliceOfPointersCallback(_participant, "Resources", &_participant.Resources)
					}
				} else {
					// ensure resource_ is NOT in _participant.Resources
					idx := slices.Index(_participant.Resources, resource_)
					if idx != -1 {
						_participant.Resources = slices.Delete(_participant.Resources, idx, idx+1)
						resourceFormCallback.probe.UpdateSliceOfPointersCallback(_participant, "Resources", &_participant.Resources)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if resourceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		resource_.Unstage(resourceFormCallback.probe.stageOfInterest)
	}

	resourceFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Resource](
		resourceFormCallback.probe,
	)

	// display a new form by reset the form stage
	if resourceFormCallback.CreationMode || resourceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		resourceFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(resourceFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ResourceFormCallback(
			nil,
			resourceFormCallback.probe,
			newFormGroup,
		)
		resource := new(models.Resource)
		FillUpForm(resource, newFormGroup, resourceFormCallback.probe)
		resourceFormCallback.probe.formStage.Commit()
	}

	resourceFormCallback.probe.ux_tree()
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
		case "Description":
			FormDivBasicFieldToField(&(task_.Description), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(task_.ComputedPrefix), formDiv)
		case "IsStartTask":
			FormDivBasicFieldToField(&(task_.IsStartTask), formDiv)
		case "IsEndTask":
			FormDivBasicFieldToField(&(task_.IsEndTask), formDiv)
		case "Type":
			FormDivSelectFieldToField(&(task_.Type), taskFormCallback.probe.stageOfInterest, formDiv)
		case "IsTaskNameNotProcessName":
			FormDivBasicFieldToField(&(task_.IsTaskNameNotProcessName), formDiv)
		case "DiagramProcess:TasksWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramProcess instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramProcess instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramProcess](taskFormCallback.probe.stageOfInterest)
			targetDiagramProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramProcess instances and update their TasksWhoseNodeIsExpanded slice
			for _diagramprocess := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramProcess](taskFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(taskFormCallback.probe.stageOfInterest, _diagramprocess)
				
				// if DiagramProcess is selected
				if targetDiagramProcessIDs[id] {
					// ensure task_ is in _diagramprocess.TasksWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramprocess.TasksWhoseNodeIsExpanded {
						if _b == task_ {
							found = true
							break
						}
					}
					if !found {
						_diagramprocess.TasksWhoseNodeIsExpanded = append(_diagramprocess.TasksWhoseNodeIsExpanded, task_)
						taskFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "TasksWhoseNodeIsExpanded", &_diagramprocess.TasksWhoseNodeIsExpanded)
					}
				} else {
					// ensure task_ is NOT in _diagramprocess.TasksWhoseNodeIsExpanded
					idx := slices.Index(_diagramprocess.TasksWhoseNodeIsExpanded, task_)
					if idx != -1 {
						_diagramprocess.TasksWhoseNodeIsExpanded = slices.Delete(_diagramprocess.TasksWhoseNodeIsExpanded, idx, idx+1)
						taskFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "TasksWhoseNodeIsExpanded", &_diagramprocess.TasksWhoseNodeIsExpanded)
					}
				}
			}
		case "Note:Tasks":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Note instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Note instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Note](taskFormCallback.probe.stageOfInterest)
			targetNoteIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetNoteIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Note instances and update their Tasks slice
			for _note := range *models.GetGongstructInstancesSetFromPointerType[*models.Note](taskFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(taskFormCallback.probe.stageOfInterest, _note)
				
				// if Note is selected
				if targetNoteIDs[id] {
					// ensure task_ is in _note.Tasks
					found := false
					for _, _b := range _note.Tasks {
						if _b == task_ {
							found = true
							break
						}
					}
					if !found {
						_note.Tasks = append(_note.Tasks, task_)
						taskFormCallback.probe.UpdateSliceOfPointersCallback(_note, "Tasks", &_note.Tasks)
					}
				} else {
					// ensure task_ is NOT in _note.Tasks
					idx := slices.Index(_note.Tasks, task_)
					if idx != -1 {
						_note.Tasks = slices.Delete(_note.Tasks, idx, idx+1)
						taskFormCallback.probe.UpdateSliceOfPointersCallback(_note, "Tasks", &_note.Tasks)
					}
				}
			}
		case "Participant:Tasks":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Participant instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Participant instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Participant](taskFormCallback.probe.stageOfInterest)
			targetParticipantIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetParticipantIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Participant instances and update their Tasks slice
			for _participant := range *models.GetGongstructInstancesSetFromPointerType[*models.Participant](taskFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(taskFormCallback.probe.stageOfInterest, _participant)
				
				// if Participant is selected
				if targetParticipantIDs[id] {
					// ensure task_ is in _participant.Tasks
					found := false
					for _, _b := range _participant.Tasks {
						if _b == task_ {
							found = true
							break
						}
					}
					if !found {
						_participant.Tasks = append(_participant.Tasks, task_)
						taskFormCallback.probe.UpdateSliceOfPointersCallback(_participant, "Tasks", &_participant.Tasks)
					}
				} else {
					// ensure task_ is NOT in _participant.Tasks
					idx := slices.Index(_participant.Tasks, task_)
					if idx != -1 {
						_participant.Tasks = slices.Delete(_participant.Tasks, idx, idx+1)
						taskFormCallback.probe.UpdateSliceOfPointersCallback(_participant, "Tasks", &_participant.Tasks)
					}
				}
			}
		case "Participant:TaskWhoseOutControlFlowsNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Participant instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Participant instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Participant](taskFormCallback.probe.stageOfInterest)
			targetParticipantIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetParticipantIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Participant instances and update their TaskWhoseOutControlFlowsNodeIsExpanded slice
			for _participant := range *models.GetGongstructInstancesSetFromPointerType[*models.Participant](taskFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(taskFormCallback.probe.stageOfInterest, _participant)
				
				// if Participant is selected
				if targetParticipantIDs[id] {
					// ensure task_ is in _participant.TaskWhoseOutControlFlowsNodeIsExpanded
					found := false
					for _, _b := range _participant.TaskWhoseOutControlFlowsNodeIsExpanded {
						if _b == task_ {
							found = true
							break
						}
					}
					if !found {
						_participant.TaskWhoseOutControlFlowsNodeIsExpanded = append(_participant.TaskWhoseOutControlFlowsNodeIsExpanded, task_)
						taskFormCallback.probe.UpdateSliceOfPointersCallback(_participant, "TaskWhoseOutControlFlowsNodeIsExpanded", &_participant.TaskWhoseOutControlFlowsNodeIsExpanded)
					}
				} else {
					// ensure task_ is NOT in _participant.TaskWhoseOutControlFlowsNodeIsExpanded
					idx := slices.Index(_participant.TaskWhoseOutControlFlowsNodeIsExpanded, task_)
					if idx != -1 {
						_participant.TaskWhoseOutControlFlowsNodeIsExpanded = slices.Delete(_participant.TaskWhoseOutControlFlowsNodeIsExpanded, idx, idx+1)
						taskFormCallback.probe.UpdateSliceOfPointersCallback(_participant, "TaskWhoseOutControlFlowsNodeIsExpanded", &_participant.TaskWhoseOutControlFlowsNodeIsExpanded)
					}
				}
			}
		case "Participant:TaskWhoseInControlFlowsNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Participant instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Participant instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Participant](taskFormCallback.probe.stageOfInterest)
			targetParticipantIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetParticipantIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Participant instances and update their TaskWhoseInControlFlowsNodeIsExpanded slice
			for _participant := range *models.GetGongstructInstancesSetFromPointerType[*models.Participant](taskFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(taskFormCallback.probe.stageOfInterest, _participant)
				
				// if Participant is selected
				if targetParticipantIDs[id] {
					// ensure task_ is in _participant.TaskWhoseInControlFlowsNodeIsExpanded
					found := false
					for _, _b := range _participant.TaskWhoseInControlFlowsNodeIsExpanded {
						if _b == task_ {
							found = true
							break
						}
					}
					if !found {
						_participant.TaskWhoseInControlFlowsNodeIsExpanded = append(_participant.TaskWhoseInControlFlowsNodeIsExpanded, task_)
						taskFormCallback.probe.UpdateSliceOfPointersCallback(_participant, "TaskWhoseInControlFlowsNodeIsExpanded", &_participant.TaskWhoseInControlFlowsNodeIsExpanded)
					}
				} else {
					// ensure task_ is NOT in _participant.TaskWhoseInControlFlowsNodeIsExpanded
					idx := slices.Index(_participant.TaskWhoseInControlFlowsNodeIsExpanded, task_)
					if idx != -1 {
						_participant.TaskWhoseInControlFlowsNodeIsExpanded = slices.Delete(_participant.TaskWhoseInControlFlowsNodeIsExpanded, idx, idx+1)
						taskFormCallback.probe.UpdateSliceOfPointersCallback(_participant, "TaskWhoseInControlFlowsNodeIsExpanded", &_participant.TaskWhoseInControlFlowsNodeIsExpanded)
					}
				}
			}
		case "Participant:TaskWhoseOutDataFlowsNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Participant instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Participant instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Participant](taskFormCallback.probe.stageOfInterest)
			targetParticipantIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetParticipantIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Participant instances and update their TaskWhoseOutDataFlowsNodeIsExpanded slice
			for _participant := range *models.GetGongstructInstancesSetFromPointerType[*models.Participant](taskFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(taskFormCallback.probe.stageOfInterest, _participant)
				
				// if Participant is selected
				if targetParticipantIDs[id] {
					// ensure task_ is in _participant.TaskWhoseOutDataFlowsNodeIsExpanded
					found := false
					for _, _b := range _participant.TaskWhoseOutDataFlowsNodeIsExpanded {
						if _b == task_ {
							found = true
							break
						}
					}
					if !found {
						_participant.TaskWhoseOutDataFlowsNodeIsExpanded = append(_participant.TaskWhoseOutDataFlowsNodeIsExpanded, task_)
						taskFormCallback.probe.UpdateSliceOfPointersCallback(_participant, "TaskWhoseOutDataFlowsNodeIsExpanded", &_participant.TaskWhoseOutDataFlowsNodeIsExpanded)
					}
				} else {
					// ensure task_ is NOT in _participant.TaskWhoseOutDataFlowsNodeIsExpanded
					idx := slices.Index(_participant.TaskWhoseOutDataFlowsNodeIsExpanded, task_)
					if idx != -1 {
						_participant.TaskWhoseOutDataFlowsNodeIsExpanded = slices.Delete(_participant.TaskWhoseOutDataFlowsNodeIsExpanded, idx, idx+1)
						taskFormCallback.probe.UpdateSliceOfPointersCallback(_participant, "TaskWhoseOutDataFlowsNodeIsExpanded", &_participant.TaskWhoseOutDataFlowsNodeIsExpanded)
					}
				}
			}
		case "Participant:TaskWhoseInDataFlowsNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Participant instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Participant instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Participant](taskFormCallback.probe.stageOfInterest)
			targetParticipantIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetParticipantIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Participant instances and update their TaskWhoseInDataFlowsNodeIsExpanded slice
			for _participant := range *models.GetGongstructInstancesSetFromPointerType[*models.Participant](taskFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(taskFormCallback.probe.stageOfInterest, _participant)
				
				// if Participant is selected
				if targetParticipantIDs[id] {
					// ensure task_ is in _participant.TaskWhoseInDataFlowsNodeIsExpanded
					found := false
					for _, _b := range _participant.TaskWhoseInDataFlowsNodeIsExpanded {
						if _b == task_ {
							found = true
							break
						}
					}
					if !found {
						_participant.TaskWhoseInDataFlowsNodeIsExpanded = append(_participant.TaskWhoseInDataFlowsNodeIsExpanded, task_)
						taskFormCallback.probe.UpdateSliceOfPointersCallback(_participant, "TaskWhoseInDataFlowsNodeIsExpanded", &_participant.TaskWhoseInDataFlowsNodeIsExpanded)
					}
				} else {
					// ensure task_ is NOT in _participant.TaskWhoseInDataFlowsNodeIsExpanded
					idx := slices.Index(_participant.TaskWhoseInDataFlowsNodeIsExpanded, task_)
					if idx != -1 {
						_participant.TaskWhoseInDataFlowsNodeIsExpanded = slices.Delete(_participant.TaskWhoseInDataFlowsNodeIsExpanded, idx, idx+1)
						taskFormCallback.probe.UpdateSliceOfPointersCallback(_participant, "TaskWhoseInDataFlowsNodeIsExpanded", &_participant.TaskWhoseInDataFlowsNodeIsExpanded)
					}
				}
			}
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
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramProcess instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramProcess instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramProcess](taskshapeFormCallback.probe.stageOfInterest)
			targetDiagramProcessIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramProcessIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramProcess instances and update their Task_Shapes slice
			for _diagramprocess := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramProcess](taskshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(taskshapeFormCallback.probe.stageOfInterest, _diagramprocess)
				
				// if DiagramProcess is selected
				if targetDiagramProcessIDs[id] {
					// ensure taskshape_ is in _diagramprocess.Task_Shapes
					found := false
					for _, _b := range _diagramprocess.Task_Shapes {
						if _b == taskshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramprocess.Task_Shapes = append(_diagramprocess.Task_Shapes, taskshape_)
						taskshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "Task_Shapes", &_diagramprocess.Task_Shapes)
					}
				} else {
					// ensure taskshape_ is NOT in _diagramprocess.Task_Shapes
					idx := slices.Index(_diagramprocess.Task_Shapes, taskshape_)
					if idx != -1 {
						_diagramprocess.Task_Shapes = slices.Delete(_diagramprocess.Task_Shapes, idx, idx+1)
						taskshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramprocess, "Task_Shapes", &_diagramprocess.Task_Shapes)
					}
				}
			}
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
