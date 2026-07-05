// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/structure/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
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
		case "Part":
			FormDivSelectFieldToField(&(allocatedresourceshape_.Part), allocatedresourceshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Resource":
			FormDivSelectFieldToField(&(allocatedresourceshape_.Resource), allocatedresourceshapeFormCallback.probe.stageOfInterest, formDiv)
		case "DiagramStructure:AllocatedResourceShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramStructure instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramStructure instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramStructure](allocatedresourceshapeFormCallback.probe.stageOfInterest)
			targetDiagramStructureIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramStructureIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramStructure instances and update their AllocatedResourceShapes slice
			for _diagramstructure := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramStructure](allocatedresourceshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(allocatedresourceshapeFormCallback.probe.stageOfInterest, _diagramstructure)
				
				// if DiagramStructure is selected
				if targetDiagramStructureIDs[id] {
					// ensure allocatedresourceshape_ is in _diagramstructure.AllocatedResourceShapes
					found := false
					for _, _b := range _diagramstructure.AllocatedResourceShapes {
						if _b == allocatedresourceshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramstructure.AllocatedResourceShapes = append(_diagramstructure.AllocatedResourceShapes, allocatedresourceshape_)
						allocatedresourceshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "AllocatedResourceShapes", &_diagramstructure.AllocatedResourceShapes)
					}
				} else {
					// ensure allocatedresourceshape_ is NOT in _diagramstructure.AllocatedResourceShapes
					idx := slices.Index(_diagramstructure.AllocatedResourceShapes, allocatedresourceshape_)
					if idx != -1 {
						_diagramstructure.AllocatedResourceShapes = slices.Delete(_diagramstructure.AllocatedResourceShapes, idx, idx+1)
						allocatedresourceshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "AllocatedResourceShapes", &_diagramstructure.AllocatedResourceShapes)
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
func __gong__New__AllocatedSystemShapeFormCallback(
	allocatedsystemshape *models.AllocatedSystemShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (allocatedsystemshapeFormCallback *AllocatedSystemShapeFormCallback) {
	allocatedsystemshapeFormCallback = new(AllocatedSystemShapeFormCallback)
	allocatedsystemshapeFormCallback.probe = probe
	allocatedsystemshapeFormCallback.allocatedsystemshape = allocatedsystemshape
	allocatedsystemshapeFormCallback.formGroup = formGroup

	allocatedsystemshapeFormCallback.CreationMode = (allocatedsystemshape == nil)

	return
}

type AllocatedSystemShapeFormCallback struct {
	allocatedsystemshape *models.AllocatedSystemShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (allocatedsystemshapeFormCallback *AllocatedSystemShapeFormCallback) OnSave() {
	allocatedsystemshapeFormCallback.probe.stageOfInterest.Lock()
	defer allocatedsystemshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("AllocatedSystemShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	allocatedsystemshapeFormCallback.probe.formStage.Checkout()

	if allocatedsystemshapeFormCallback.allocatedsystemshape == nil {
		allocatedsystemshapeFormCallback.allocatedsystemshape = new(models.AllocatedSystemShape).Stage(allocatedsystemshapeFormCallback.probe.stageOfInterest)
	}
	allocatedsystemshape_ := allocatedsystemshapeFormCallback.allocatedsystemshape
	_ = allocatedsystemshape_

	for _, formDiv := range allocatedsystemshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(allocatedsystemshape_.Name), formDiv)
		case "Part":
			FormDivSelectFieldToField(&(allocatedsystemshape_.Part), allocatedsystemshapeFormCallback.probe.stageOfInterest, formDiv)
		case "System":
			FormDivSelectFieldToField(&(allocatedsystemshape_.System), allocatedsystemshapeFormCallback.probe.stageOfInterest, formDiv)
		case "DiagramStructure:AllocatedSystemShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramStructure instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramStructure instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramStructure](allocatedsystemshapeFormCallback.probe.stageOfInterest)
			targetDiagramStructureIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramStructureIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramStructure instances and update their AllocatedSystemShapes slice
			for _diagramstructure := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramStructure](allocatedsystemshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(allocatedsystemshapeFormCallback.probe.stageOfInterest, _diagramstructure)
				
				// if DiagramStructure is selected
				if targetDiagramStructureIDs[id] {
					// ensure allocatedsystemshape_ is in _diagramstructure.AllocatedSystemShapes
					found := false
					for _, _b := range _diagramstructure.AllocatedSystemShapes {
						if _b == allocatedsystemshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramstructure.AllocatedSystemShapes = append(_diagramstructure.AllocatedSystemShapes, allocatedsystemshape_)
						allocatedsystemshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "AllocatedSystemShapes", &_diagramstructure.AllocatedSystemShapes)
					}
				} else {
					// ensure allocatedsystemshape_ is NOT in _diagramstructure.AllocatedSystemShapes
					idx := slices.Index(_diagramstructure.AllocatedSystemShapes, allocatedsystemshape_)
					if idx != -1 {
						_diagramstructure.AllocatedSystemShapes = slices.Delete(_diagramstructure.AllocatedSystemShapes, idx, idx+1)
						allocatedsystemshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "AllocatedSystemShapes", &_diagramstructure.AllocatedSystemShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if allocatedsystemshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		allocatedsystemshape_.Unstage(allocatedsystemshapeFormCallback.probe.stageOfInterest)
	}

	allocatedsystemshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.AllocatedSystemShape](
		allocatedsystemshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if allocatedsystemshapeFormCallback.CreationMode || allocatedsystemshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		allocatedsystemshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(allocatedsystemshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AllocatedSystemShapeFormCallback(
			nil,
			allocatedsystemshapeFormCallback.probe,
			newFormGroup,
		)
		allocatedsystemshape := new(models.AllocatedSystemShape)
		FillUpForm(allocatedsystemshape, newFormGroup, allocatedsystemshapeFormCallback.probe)
		allocatedsystemshapeFormCallback.probe.formStage.Commit()
	}

	allocatedsystemshapeFormCallback.probe.ux_tree()
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
		case "IsExpanded":
			FormDivBasicFieldToField(&(controlflow_.IsExpanded), formDiv)
		case "Start":
			FormDivSelectFieldToField(&(controlflow_.Start), controlflowFormCallback.probe.stageOfInterest, formDiv)
		case "End":
			FormDivSelectFieldToField(&(controlflow_.End), controlflowFormCallback.probe.stageOfInterest, formDiv)
		case "DiagramStructure:ControlFlowsWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramStructure instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramStructure instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramStructure](controlflowFormCallback.probe.stageOfInterest)
			targetDiagramStructureIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramStructureIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramStructure instances and update their ControlFlowsWhoseNodeIsExpanded slice
			for _diagramstructure := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramStructure](controlflowFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(controlflowFormCallback.probe.stageOfInterest, _diagramstructure)
				
				// if DiagramStructure is selected
				if targetDiagramStructureIDs[id] {
					// ensure controlflow_ is in _diagramstructure.ControlFlowsWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramstructure.ControlFlowsWhoseNodeIsExpanded {
						if _b == controlflow_ {
							found = true
							break
						}
					}
					if !found {
						_diagramstructure.ControlFlowsWhoseNodeIsExpanded = append(_diagramstructure.ControlFlowsWhoseNodeIsExpanded, controlflow_)
						controlflowFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "ControlFlowsWhoseNodeIsExpanded", &_diagramstructure.ControlFlowsWhoseNodeIsExpanded)
					}
				} else {
					// ensure controlflow_ is NOT in _diagramstructure.ControlFlowsWhoseNodeIsExpanded
					idx := slices.Index(_diagramstructure.ControlFlowsWhoseNodeIsExpanded, controlflow_)
					if idx != -1 {
						_diagramstructure.ControlFlowsWhoseNodeIsExpanded = slices.Delete(_diagramstructure.ControlFlowsWhoseNodeIsExpanded, idx, idx+1)
						controlflowFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "ControlFlowsWhoseNodeIsExpanded", &_diagramstructure.ControlFlowsWhoseNodeIsExpanded)
					}
				}
			}
		case "Part:ControlFlows":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Part instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Part instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Part](controlflowFormCallback.probe.stageOfInterest)
			targetPartIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetPartIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Part instances and update their ControlFlows slice
			for _part := range *models.GetGongstructInstancesSetFromPointerType[*models.Part](controlflowFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(controlflowFormCallback.probe.stageOfInterest, _part)
				
				// if Part is selected
				if targetPartIDs[id] {
					// ensure controlflow_ is in _part.ControlFlows
					found := false
					for _, _b := range _part.ControlFlows {
						if _b == controlflow_ {
							found = true
							break
						}
					}
					if !found {
						_part.ControlFlows = append(_part.ControlFlows, controlflow_)
						controlflowFormCallback.probe.UpdateSliceOfPointersCallback(_part, "ControlFlows", &_part.ControlFlows)
					}
				} else {
					// ensure controlflow_ is NOT in _part.ControlFlows
					idx := slices.Index(_part.ControlFlows, controlflow_)
					if idx != -1 {
						_part.ControlFlows = slices.Delete(_part.ControlFlows, idx, idx+1)
						controlflowFormCallback.probe.UpdateSliceOfPointersCallback(_part, "ControlFlows", &_part.ControlFlows)
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
		case "DiagramStructure:ControlFlow_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramStructure instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramStructure instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramStructure](controlflowshapeFormCallback.probe.stageOfInterest)
			targetDiagramStructureIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramStructureIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramStructure instances and update their ControlFlow_Shapes slice
			for _diagramstructure := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramStructure](controlflowshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(controlflowshapeFormCallback.probe.stageOfInterest, _diagramstructure)
				
				// if DiagramStructure is selected
				if targetDiagramStructureIDs[id] {
					// ensure controlflowshape_ is in _diagramstructure.ControlFlow_Shapes
					found := false
					for _, _b := range _diagramstructure.ControlFlow_Shapes {
						if _b == controlflowshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramstructure.ControlFlow_Shapes = append(_diagramstructure.ControlFlow_Shapes, controlflowshape_)
						controlflowshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "ControlFlow_Shapes", &_diagramstructure.ControlFlow_Shapes)
					}
				} else {
					// ensure controlflowshape_ is NOT in _diagramstructure.ControlFlow_Shapes
					idx := slices.Index(_diagramstructure.ControlFlow_Shapes, controlflowshape_)
					if idx != -1 {
						_diagramstructure.ControlFlow_Shapes = slices.Delete(_diagramstructure.ControlFlow_Shapes, idx, idx+1)
						controlflowshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "ControlFlow_Shapes", &_diagramstructure.ControlFlow_Shapes)
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
		case "IsExpanded":
			FormDivBasicFieldToField(&(data_.IsExpanded), formDiv)
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
		case "DiagramStructure:DatasWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramStructure instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramStructure instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramStructure](dataFormCallback.probe.stageOfInterest)
			targetDiagramStructureIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramStructureIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramStructure instances and update their DatasWhoseNodeIsExpanded slice
			for _diagramstructure := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramStructure](dataFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(dataFormCallback.probe.stageOfInterest, _diagramstructure)
				
				// if DiagramStructure is selected
				if targetDiagramStructureIDs[id] {
					// ensure data_ is in _diagramstructure.DatasWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramstructure.DatasWhoseNodeIsExpanded {
						if _b == data_ {
							found = true
							break
						}
					}
					if !found {
						_diagramstructure.DatasWhoseNodeIsExpanded = append(_diagramstructure.DatasWhoseNodeIsExpanded, data_)
						dataFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "DatasWhoseNodeIsExpanded", &_diagramstructure.DatasWhoseNodeIsExpanded)
					}
				} else {
					// ensure data_ is NOT in _diagramstructure.DatasWhoseNodeIsExpanded
					idx := slices.Index(_diagramstructure.DatasWhoseNodeIsExpanded, data_)
					if idx != -1 {
						_diagramstructure.DatasWhoseNodeIsExpanded = slices.Delete(_diagramstructure.DatasWhoseNodeIsExpanded, idx, idx+1)
						dataFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "DatasWhoseNodeIsExpanded", &_diagramstructure.DatasWhoseNodeIsExpanded)
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
		case "StartPort":
			FormDivSelectFieldToField(&(dataflow_.StartPort), dataflowFormCallback.probe.stageOfInterest, formDiv)
		case "EndPort":
			FormDivSelectFieldToField(&(dataflow_.EndPort), dataflowFormCallback.probe.stageOfInterest, formDiv)
		case "StartExternalPart":
			FormDivSelectFieldToField(&(dataflow_.StartExternalPart), dataflowFormCallback.probe.stageOfInterest, formDiv)
		case "EndExternalPart":
			FormDivSelectFieldToField(&(dataflow_.EndExternalPart), dataflowFormCallback.probe.stageOfInterest, formDiv)
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
		case "Direction":
			FormDivEnumStringFieldToField(&(dataflow_.Direction), formDiv)
		case "IsDatasNodeExpanded":
			FormDivBasicFieldToField(&(dataflow_.IsDatasNodeExpanded), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(dataflow_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(dataflow_.IsExpanded), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(dataflow_.Type), formDiv)
		case "DiagramStructure:DataFlowsWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramStructure instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramStructure instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramStructure](dataflowFormCallback.probe.stageOfInterest)
			targetDiagramStructureIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramStructureIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramStructure instances and update their DataFlowsWhoseNodeIsExpanded slice
			for _diagramstructure := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramStructure](dataflowFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(dataflowFormCallback.probe.stageOfInterest, _diagramstructure)
				
				// if DiagramStructure is selected
				if targetDiagramStructureIDs[id] {
					// ensure dataflow_ is in _diagramstructure.DataFlowsWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramstructure.DataFlowsWhoseNodeIsExpanded {
						if _b == dataflow_ {
							found = true
							break
						}
					}
					if !found {
						_diagramstructure.DataFlowsWhoseNodeIsExpanded = append(_diagramstructure.DataFlowsWhoseNodeIsExpanded, dataflow_)
						dataflowFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "DataFlowsWhoseNodeIsExpanded", &_diagramstructure.DataFlowsWhoseNodeIsExpanded)
					}
				} else {
					// ensure dataflow_ is NOT in _diagramstructure.DataFlowsWhoseNodeIsExpanded
					idx := slices.Index(_diagramstructure.DataFlowsWhoseNodeIsExpanded, dataflow_)
					if idx != -1 {
						_diagramstructure.DataFlowsWhoseNodeIsExpanded = slices.Delete(_diagramstructure.DataFlowsWhoseNodeIsExpanded, idx, idx+1)
						dataflowFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "DataFlowsWhoseNodeIsExpanded", &_diagramstructure.DataFlowsWhoseNodeIsExpanded)
					}
				}
			}
		case "DiagramStructure:DataFlowsWhoseDataNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramStructure instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramStructure instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramStructure](dataflowFormCallback.probe.stageOfInterest)
			targetDiagramStructureIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramStructureIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramStructure instances and update their DataFlowsWhoseDataNodeIsExpanded slice
			for _diagramstructure := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramStructure](dataflowFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(dataflowFormCallback.probe.stageOfInterest, _diagramstructure)
				
				// if DiagramStructure is selected
				if targetDiagramStructureIDs[id] {
					// ensure dataflow_ is in _diagramstructure.DataFlowsWhoseDataNodeIsExpanded
					found := false
					for _, _b := range _diagramstructure.DataFlowsWhoseDataNodeIsExpanded {
						if _b == dataflow_ {
							found = true
							break
						}
					}
					if !found {
						_diagramstructure.DataFlowsWhoseDataNodeIsExpanded = append(_diagramstructure.DataFlowsWhoseDataNodeIsExpanded, dataflow_)
						dataflowFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "DataFlowsWhoseDataNodeIsExpanded", &_diagramstructure.DataFlowsWhoseDataNodeIsExpanded)
					}
				} else {
					// ensure dataflow_ is NOT in _diagramstructure.DataFlowsWhoseDataNodeIsExpanded
					idx := slices.Index(_diagramstructure.DataFlowsWhoseDataNodeIsExpanded, dataflow_)
					if idx != -1 {
						_diagramstructure.DataFlowsWhoseDataNodeIsExpanded = slices.Delete(_diagramstructure.DataFlowsWhoseDataNodeIsExpanded, idx, idx+1)
						dataflowFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "DataFlowsWhoseDataNodeIsExpanded", &_diagramstructure.DataFlowsWhoseDataNodeIsExpanded)
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
		case "System:DataFlows":
			// 1. Decode the AssociationStorage which contains the rowIDs of the System instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target System instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.System](dataflowFormCallback.probe.stageOfInterest)
			targetSystemIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSystemIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all System instances and update their DataFlows slice
			for _system := range *models.GetGongstructInstancesSetFromPointerType[*models.System](dataflowFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(dataflowFormCallback.probe.stageOfInterest, _system)
				
				// if System is selected
				if targetSystemIDs[id] {
					// ensure dataflow_ is in _system.DataFlows
					found := false
					for _, _b := range _system.DataFlows {
						if _b == dataflow_ {
							found = true
							break
						}
					}
					if !found {
						_system.DataFlows = append(_system.DataFlows, dataflow_)
						dataflowFormCallback.probe.UpdateSliceOfPointersCallback(_system, "DataFlows", &_system.DataFlows)
					}
				} else {
					// ensure dataflow_ is NOT in _system.DataFlows
					idx := slices.Index(_system.DataFlows, dataflow_)
					if idx != -1 {
						_system.DataFlows = slices.Delete(_system.DataFlows, idx, idx+1)
						dataflowFormCallback.probe.UpdateSliceOfPointersCallback(_system, "DataFlows", &_system.DataFlows)
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
		case "DiagramStructure:DataFlow_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramStructure instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramStructure instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramStructure](dataflowshapeFormCallback.probe.stageOfInterest)
			targetDiagramStructureIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramStructureIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramStructure instances and update their DataFlow_Shapes slice
			for _diagramstructure := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramStructure](dataflowshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(dataflowshapeFormCallback.probe.stageOfInterest, _diagramstructure)
				
				// if DiagramStructure is selected
				if targetDiagramStructureIDs[id] {
					// ensure dataflowshape_ is in _diagramstructure.DataFlow_Shapes
					found := false
					for _, _b := range _diagramstructure.DataFlow_Shapes {
						if _b == dataflowshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramstructure.DataFlow_Shapes = append(_diagramstructure.DataFlow_Shapes, dataflowshape_)
						dataflowshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "DataFlow_Shapes", &_diagramstructure.DataFlow_Shapes)
					}
				} else {
					// ensure dataflowshape_ is NOT in _diagramstructure.DataFlow_Shapes
					idx := slices.Index(_diagramstructure.DataFlow_Shapes, dataflowshape_)
					if idx != -1 {
						_diagramstructure.DataFlow_Shapes = slices.Delete(_diagramstructure.DataFlow_Shapes, idx, idx+1)
						dataflowshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "DataFlow_Shapes", &_diagramstructure.DataFlow_Shapes)
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
		case "DiagramStructure:Data_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramStructure instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramStructure instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramStructure](datashapeFormCallback.probe.stageOfInterest)
			targetDiagramStructureIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramStructureIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramStructure instances and update their Data_Shapes slice
			for _diagramstructure := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramStructure](datashapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(datashapeFormCallback.probe.stageOfInterest, _diagramstructure)
				
				// if DiagramStructure is selected
				if targetDiagramStructureIDs[id] {
					// ensure datashape_ is in _diagramstructure.Data_Shapes
					found := false
					for _, _b := range _diagramstructure.Data_Shapes {
						if _b == datashape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramstructure.Data_Shapes = append(_diagramstructure.Data_Shapes, datashape_)
						datashapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "Data_Shapes", &_diagramstructure.Data_Shapes)
					}
				} else {
					// ensure datashape_ is NOT in _diagramstructure.Data_Shapes
					idx := slices.Index(_diagramstructure.Data_Shapes, datashape_)
					if idx != -1 {
						_diagramstructure.Data_Shapes = slices.Delete(_diagramstructure.Data_Shapes, idx, idx+1)
						datashapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "Data_Shapes", &_diagramstructure.Data_Shapes)
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
func __gong__New__DiagramStructureFormCallback(
	diagramstructure *models.DiagramStructure,
	probe *Probe,
	formGroup *form.FormGroup,
) (diagramstructureFormCallback *DiagramStructureFormCallback) {
	diagramstructureFormCallback = new(DiagramStructureFormCallback)
	diagramstructureFormCallback.probe = probe
	diagramstructureFormCallback.diagramstructure = diagramstructure
	diagramstructureFormCallback.formGroup = formGroup

	diagramstructureFormCallback.CreationMode = (diagramstructure == nil)

	return
}

type DiagramStructureFormCallback struct {
	diagramstructure *models.DiagramStructure

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (diagramstructureFormCallback *DiagramStructureFormCallback) OnSave() {
	diagramstructureFormCallback.probe.stageOfInterest.Lock()
	defer diagramstructureFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DiagramStructureFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	diagramstructureFormCallback.probe.formStage.Checkout()

	if diagramstructureFormCallback.diagramstructure == nil {
		diagramstructureFormCallback.diagramstructure = new(models.DiagramStructure).Stage(diagramstructureFormCallback.probe.stageOfInterest)
	}
	diagramstructure_ := diagramstructureFormCallback.diagramstructure
	_ = diagramstructure_

	for _, formDiv := range diagramstructureFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(diagramstructure_.Name), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(diagramstructure_.Description), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(diagramstructure_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(diagramstructure_.IsExpanded), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(diagramstructure_.IsChecked), formDiv)
		case "IsEditable_":
			FormDivBasicFieldToField(&(diagramstructure_.IsEditable_), formDiv)
		case "IsShowPrefix":
			FormDivBasicFieldToField(&(diagramstructure_.IsShowPrefix), formDiv)
		case "DefaultBoxWidth":
			FormDivBasicFieldToField(&(diagramstructure_.DefaultBoxWidth), formDiv)
		case "DefaultBoxHeigth":
			FormDivBasicFieldToField(&(diagramstructure_.DefaultBoxHeigth), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(diagramstructure_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(diagramstructure_.Height), formDiv)
		case "System_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.SystemShape](diagramstructureFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.SystemShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.SystemShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramstructureFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.SystemShape](diagramstructureFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramstructure_.System_Shapes = instanceSlice
			diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(diagramstructure_, "System_Shapes", &diagramstructure_.System_Shapes)

		case "IsSystemsNodeExpanded":
			FormDivBasicFieldToField(&(diagramstructure_.IsSystemsNodeExpanded), formDiv)
		case "SystemsWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.System](diagramstructureFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.System, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.System)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramstructureFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.System](diagramstructureFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramstructure_.SystemsWhoseNodeIsExpanded = instanceSlice
			diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(diagramstructure_, "SystemsWhoseNodeIsExpanded", &diagramstructure_.SystemsWhoseNodeIsExpanded)

		case "Part_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.PartShape](diagramstructureFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.PartShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.PartShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramstructureFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.PartShape](diagramstructureFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramstructure_.Part_Shapes = instanceSlice
			diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(diagramstructure_, "Part_Shapes", &diagramstructure_.Part_Shapes)

		case "IsPartsNodeExpanded":
			FormDivBasicFieldToField(&(diagramstructure_.IsPartsNodeExpanded), formDiv)
		case "PartWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Part](diagramstructureFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Part, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Part)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramstructureFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Part](diagramstructureFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramstructure_.PartWhoseNodeIsExpanded = instanceSlice
			diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(diagramstructure_, "PartWhoseNodeIsExpanded", &diagramstructure_.PartWhoseNodeIsExpanded)

		case "ExternalPart_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ExternalPartShape](diagramstructureFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ExternalPartShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ExternalPartShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramstructureFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ExternalPartShape](diagramstructureFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramstructure_.ExternalPart_Shapes = instanceSlice
			diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(diagramstructure_, "ExternalPart_Shapes", &diagramstructure_.ExternalPart_Shapes)

		case "IsExternalPartsNodeExpanded":
			FormDivBasicFieldToField(&(diagramstructure_.IsExternalPartsNodeExpanded), formDiv)
		case "ExternalPartWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Part](diagramstructureFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Part, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Part)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramstructureFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Part](diagramstructureFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramstructure_.ExternalPartWhoseNodeIsExpanded = instanceSlice
			diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(diagramstructure_, "ExternalPartWhoseNodeIsExpanded", &diagramstructure_.ExternalPartWhoseNodeIsExpanded)

		case "ExternalPartsWhoseOutDataFlowsNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Part](diagramstructureFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Part, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Part)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramstructureFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Part](diagramstructureFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramstructure_.ExternalPartsWhoseOutDataFlowsNodeIsExpanded = instanceSlice
			diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(diagramstructure_, "ExternalPartsWhoseOutDataFlowsNodeIsExpanded", &diagramstructure_.ExternalPartsWhoseOutDataFlowsNodeIsExpanded)

		case "ExternalPartsWhoseInDataFlowsNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Part](diagramstructureFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Part, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Part)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramstructureFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Part](diagramstructureFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramstructure_.ExternalPartsWhoseInDataFlowsNodeIsExpanded = instanceSlice
			diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(diagramstructure_, "ExternalPartsWhoseInDataFlowsNodeIsExpanded", &diagramstructure_.ExternalPartsWhoseInDataFlowsNodeIsExpanded)

		case "PortsWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Port](diagramstructureFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Port, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Port)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramstructureFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Port](diagramstructureFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramstructure_.PortsWhoseNodeIsExpanded = instanceSlice
			diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(diagramstructure_, "PortsWhoseNodeIsExpanded", &diagramstructure_.PortsWhoseNodeIsExpanded)

		case "Port_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.PortShape](diagramstructureFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.PortShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.PortShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramstructureFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.PortShape](diagramstructureFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramstructure_.Port_Shapes = instanceSlice
			diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(diagramstructure_, "Port_Shapes", &diagramstructure_.Port_Shapes)

		case "ControlFlowsWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ControlFlow](diagramstructureFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ControlFlow, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ControlFlow)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramstructureFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ControlFlow](diagramstructureFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramstructure_.ControlFlowsWhoseNodeIsExpanded = instanceSlice
			diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(diagramstructure_, "ControlFlowsWhoseNodeIsExpanded", &diagramstructure_.ControlFlowsWhoseNodeIsExpanded)

		case "ControlFlow_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ControlFlowShape](diagramstructureFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ControlFlowShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ControlFlowShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramstructureFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ControlFlowShape](diagramstructureFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramstructure_.ControlFlow_Shapes = instanceSlice
			diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(diagramstructure_, "ControlFlow_Shapes", &diagramstructure_.ControlFlow_Shapes)

		case "DataFlowsWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DataFlow](diagramstructureFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DataFlow, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DataFlow)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramstructureFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.DataFlow](diagramstructureFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramstructure_.DataFlowsWhoseNodeIsExpanded = instanceSlice
			diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(diagramstructure_, "DataFlowsWhoseNodeIsExpanded", &diagramstructure_.DataFlowsWhoseNodeIsExpanded)

		case "DataFlow_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DataFlowShape](diagramstructureFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DataFlowShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DataFlowShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramstructureFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.DataFlowShape](diagramstructureFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramstructure_.DataFlow_Shapes = instanceSlice
			diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(diagramstructure_, "DataFlow_Shapes", &diagramstructure_.DataFlow_Shapes)

		case "DatasWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Data](diagramstructureFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Data, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Data)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramstructureFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Data](diagramstructureFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramstructure_.DatasWhoseNodeIsExpanded = instanceSlice
			diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(diagramstructure_, "DatasWhoseNodeIsExpanded", &diagramstructure_.DatasWhoseNodeIsExpanded)

		case "Data_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DataShape](diagramstructureFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DataShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DataShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramstructureFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.DataShape](diagramstructureFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramstructure_.Data_Shapes = instanceSlice
			diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(diagramstructure_, "Data_Shapes", &diagramstructure_.Data_Shapes)

		case "DataFlowsWhoseDataNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DataFlow](diagramstructureFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DataFlow, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DataFlow)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramstructureFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.DataFlow](diagramstructureFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramstructure_.DataFlowsWhoseDataNodeIsExpanded = instanceSlice
			diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(diagramstructure_, "DataFlowsWhoseDataNodeIsExpanded", &diagramstructure_.DataFlowsWhoseDataNodeIsExpanded)

		case "AllocatedResourcesWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Resource](diagramstructureFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Resource, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Resource)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramstructureFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Resource](diagramstructureFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramstructure_.AllocatedResourcesWhoseNodeIsExpanded = instanceSlice
			diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(diagramstructure_, "AllocatedResourcesWhoseNodeIsExpanded", &diagramstructure_.AllocatedResourcesWhoseNodeIsExpanded)

		case "AllocatedResourceShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.AllocatedResourceShape](diagramstructureFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.AllocatedResourceShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.AllocatedResourceShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramstructureFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.AllocatedResourceShape](diagramstructureFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramstructure_.AllocatedResourceShapes = instanceSlice
			diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(diagramstructure_, "AllocatedResourceShapes", &diagramstructure_.AllocatedResourceShapes)

		case "AllocatedSystemesWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.System](diagramstructureFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.System, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.System)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramstructureFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.System](diagramstructureFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramstructure_.AllocatedSystemesWhoseNodeIsExpanded = instanceSlice
			diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(diagramstructure_, "AllocatedSystemesWhoseNodeIsExpanded", &diagramstructure_.AllocatedSystemesWhoseNodeIsExpanded)

		case "AllocatedSystemShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.AllocatedSystemShape](diagramstructureFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.AllocatedSystemShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.AllocatedSystemShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramstructureFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.AllocatedSystemShape](diagramstructureFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramstructure_.AllocatedSystemShapes = instanceSlice
			diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(diagramstructure_, "AllocatedSystemShapes", &diagramstructure_.AllocatedSystemShapes)

		case "Note_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.NoteShape](diagramstructureFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.NoteShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.NoteShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramstructureFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.NoteShape](diagramstructureFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramstructure_.Note_Shapes = instanceSlice
			diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(diagramstructure_, "Note_Shapes", &diagramstructure_.Note_Shapes)

		case "NotesWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Note](diagramstructureFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Note, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Note)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramstructureFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Note](diagramstructureFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramstructure_.NotesWhoseNodeIsExpanded = instanceSlice
			diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(diagramstructure_, "NotesWhoseNodeIsExpanded", &diagramstructure_.NotesWhoseNodeIsExpanded)

		case "IsNotesNodeExpanded":
			FormDivBasicFieldToField(&(diagramstructure_.IsNotesNodeExpanded), formDiv)
		case "NotePortShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.NotePortShape](diagramstructureFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.NotePortShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.NotePortShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramstructureFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.NotePortShape](diagramstructureFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramstructure_.NotePortShapes = instanceSlice
			diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(diagramstructure_, "NotePortShapes", &diagramstructure_.NotePortShapes)

		case "NotePartShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.NotePartShape](diagramstructureFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.NotePartShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.NotePartShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramstructureFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.NotePartShape](diagramstructureFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramstructure_.NotePartShapes = instanceSlice
			diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(diagramstructure_, "NotePartShapes", &diagramstructure_.NotePartShapes)

		case "System:DiagramStructures":
			// 1. Decode the AssociationStorage which contains the rowIDs of the System instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target System instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.System](diagramstructureFormCallback.probe.stageOfInterest)
			targetSystemIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSystemIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all System instances and update their DiagramStructures slice
			for _system := range *models.GetGongstructInstancesSetFromPointerType[*models.System](diagramstructureFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(diagramstructureFormCallback.probe.stageOfInterest, _system)
				
				// if System is selected
				if targetSystemIDs[id] {
					// ensure diagramstructure_ is in _system.DiagramStructures
					found := false
					for _, _b := range _system.DiagramStructures {
						if _b == diagramstructure_ {
							found = true
							break
						}
					}
					if !found {
						_system.DiagramStructures = append(_system.DiagramStructures, diagramstructure_)
						diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(_system, "DiagramStructures", &_system.DiagramStructures)
					}
				} else {
					// ensure diagramstructure_ is NOT in _system.DiagramStructures
					idx := slices.Index(_system.DiagramStructures, diagramstructure_)
					if idx != -1 {
						_system.DiagramStructures = slices.Delete(_system.DiagramStructures, idx, idx+1)
						diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(_system, "DiagramStructures", &_system.DiagramStructures)
					}
				}
			}
		case "System:DiagramStructureWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the System instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target System instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.System](diagramstructureFormCallback.probe.stageOfInterest)
			targetSystemIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSystemIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all System instances and update their DiagramStructureWhoseNodeIsExpanded slice
			for _system := range *models.GetGongstructInstancesSetFromPointerType[*models.System](diagramstructureFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(diagramstructureFormCallback.probe.stageOfInterest, _system)
				
				// if System is selected
				if targetSystemIDs[id] {
					// ensure diagramstructure_ is in _system.DiagramStructureWhoseNodeIsExpanded
					found := false
					for _, _b := range _system.DiagramStructureWhoseNodeIsExpanded {
						if _b == diagramstructure_ {
							found = true
							break
						}
					}
					if !found {
						_system.DiagramStructureWhoseNodeIsExpanded = append(_system.DiagramStructureWhoseNodeIsExpanded, diagramstructure_)
						diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(_system, "DiagramStructureWhoseNodeIsExpanded", &_system.DiagramStructureWhoseNodeIsExpanded)
					}
				} else {
					// ensure diagramstructure_ is NOT in _system.DiagramStructureWhoseNodeIsExpanded
					idx := slices.Index(_system.DiagramStructureWhoseNodeIsExpanded, diagramstructure_)
					if idx != -1 {
						_system.DiagramStructureWhoseNodeIsExpanded = slices.Delete(_system.DiagramStructureWhoseNodeIsExpanded, idx, idx+1)
						diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(_system, "DiagramStructureWhoseNodeIsExpanded", &_system.DiagramStructureWhoseNodeIsExpanded)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if diagramstructureFormCallback.formGroup.HasSuppressButtonBeenPressed {
		diagramstructure_.Unstage(diagramstructureFormCallback.probe.stageOfInterest)
	}

	diagramstructureFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.DiagramStructure](
		diagramstructureFormCallback.probe,
	)

	// display a new form by reset the form stage
	if diagramstructureFormCallback.CreationMode || diagramstructureFormCallback.formGroup.HasSuppressButtonBeenPressed {
		diagramstructureFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(diagramstructureFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DiagramStructureFormCallback(
			nil,
			diagramstructureFormCallback.probe,
			newFormGroup,
		)
		diagramstructure := new(models.DiagramStructure)
		FillUpForm(diagramstructure, newFormGroup, diagramstructureFormCallback.probe)
		diagramstructureFormCallback.probe.formStage.Commit()
	}

	diagramstructureFormCallback.probe.ux_tree()
}
func __gong__New__ExternalPartShapeFormCallback(
	externalpartshape *models.ExternalPartShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (externalpartshapeFormCallback *ExternalPartShapeFormCallback) {
	externalpartshapeFormCallback = new(ExternalPartShapeFormCallback)
	externalpartshapeFormCallback.probe = probe
	externalpartshapeFormCallback.externalpartshape = externalpartshape
	externalpartshapeFormCallback.formGroup = formGroup

	externalpartshapeFormCallback.CreationMode = (externalpartshape == nil)

	return
}

type ExternalPartShapeFormCallback struct {
	externalpartshape *models.ExternalPartShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (externalpartshapeFormCallback *ExternalPartShapeFormCallback) OnSave() {
	externalpartshapeFormCallback.probe.stageOfInterest.Lock()
	defer externalpartshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ExternalPartShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	externalpartshapeFormCallback.probe.formStage.Checkout()

	if externalpartshapeFormCallback.externalpartshape == nil {
		externalpartshapeFormCallback.externalpartshape = new(models.ExternalPartShape).Stage(externalpartshapeFormCallback.probe.stageOfInterest)
	}
	externalpartshape_ := externalpartshapeFormCallback.externalpartshape
	_ = externalpartshape_

	for _, formDiv := range externalpartshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(externalpartshape_.Name), formDiv)
		case "Part":
			FormDivSelectFieldToField(&(externalpartshape_.Part), externalpartshapeFormCallback.probe.stageOfInterest, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(externalpartshape_.IsExpanded), formDiv)
		case "X":
			FormDivBasicFieldToField(&(externalpartshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(externalpartshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(externalpartshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(externalpartshape_.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(externalpartshape_.IsHidden), formDiv)
		case "TailHeigth":
			FormDivBasicFieldToField(&(externalpartshape_.TailHeigth), formDiv)
		case "DiagramStructure:ExternalPart_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramStructure instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramStructure instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramStructure](externalpartshapeFormCallback.probe.stageOfInterest)
			targetDiagramStructureIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramStructureIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramStructure instances and update their ExternalPart_Shapes slice
			for _diagramstructure := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramStructure](externalpartshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(externalpartshapeFormCallback.probe.stageOfInterest, _diagramstructure)
				
				// if DiagramStructure is selected
				if targetDiagramStructureIDs[id] {
					// ensure externalpartshape_ is in _diagramstructure.ExternalPart_Shapes
					found := false
					for _, _b := range _diagramstructure.ExternalPart_Shapes {
						if _b == externalpartshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramstructure.ExternalPart_Shapes = append(_diagramstructure.ExternalPart_Shapes, externalpartshape_)
						externalpartshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "ExternalPart_Shapes", &_diagramstructure.ExternalPart_Shapes)
					}
				} else {
					// ensure externalpartshape_ is NOT in _diagramstructure.ExternalPart_Shapes
					idx := slices.Index(_diagramstructure.ExternalPart_Shapes, externalpartshape_)
					if idx != -1 {
						_diagramstructure.ExternalPart_Shapes = slices.Delete(_diagramstructure.ExternalPart_Shapes, idx, idx+1)
						externalpartshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "ExternalPart_Shapes", &_diagramstructure.ExternalPart_Shapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if externalpartshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		externalpartshape_.Unstage(externalpartshapeFormCallback.probe.stageOfInterest)
	}

	externalpartshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ExternalPartShape](
		externalpartshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if externalpartshapeFormCallback.CreationMode || externalpartshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		externalpartshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(externalpartshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ExternalPartShapeFormCallback(
			nil,
			externalpartshapeFormCallback.probe,
			newFormGroup,
		)
		externalpartshape := new(models.ExternalPartShape)
		FillUpForm(externalpartshape, newFormGroup, externalpartshapeFormCallback.probe)
		externalpartshapeFormCallback.probe.formStage.Commit()
	}

	externalpartshapeFormCallback.probe.ux_tree()
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
		case "IsExpanded":
			FormDivBasicFieldToField(&(library_.IsExpanded), formDiv)
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
		case "RootSystemes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.System](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.System, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.System)

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
			map_RowID_ID := GetMap_RowID_ID[*models.System](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.RootSystemes = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "RootSystemes", &library_.RootSystemes)

		case "IsSystemesNodeExpanded":
			FormDivBasicFieldToField(&(library_.IsSystemesNodeExpanded), formDiv)
		case "SystemsWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.System](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.System, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.System)

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
			map_RowID_ID := GetMap_RowID_ID[*models.System](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.SystemsWhoseNodeIsExpanded = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "SystemsWhoseNodeIsExpanded", &library_.SystemsWhoseNodeIsExpanded)

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

		case "PartsWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Part](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Part, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Part)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Part](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.PartsWhoseNodeIsExpanded = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "PartsWhoseNodeIsExpanded", &library_.PartsWhoseNodeIsExpanded)

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
		case "IsExpanded":
			FormDivBasicFieldToField(&(note_.IsExpanded), formDiv)
		case "IsPartsNodeExpanded":
			FormDivBasicFieldToField(&(note_.IsPartsNodeExpanded), formDiv)
		case "Parts":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Part](noteFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Part, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Part)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Part](noteFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			note_.Parts = instanceSlice
			noteFormCallback.probe.UpdateSliceOfPointersCallback(note_, "Parts", &note_.Parts)

		case "IsPortsNodeExpanded":
			FormDivBasicFieldToField(&(note_.IsPortsNodeExpanded), formDiv)
		case "Ports":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Port](noteFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Port, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Port)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Port](noteFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			note_.Ports = instanceSlice
			noteFormCallback.probe.UpdateSliceOfPointersCallback(note_, "Ports", &note_.Ports)

		case "DiagramStructure:NotesWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramStructure instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramStructure instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramStructure](noteFormCallback.probe.stageOfInterest)
			targetDiagramStructureIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramStructureIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramStructure instances and update their NotesWhoseNodeIsExpanded slice
			for _diagramstructure := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramStructure](noteFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(noteFormCallback.probe.stageOfInterest, _diagramstructure)
				
				// if DiagramStructure is selected
				if targetDiagramStructureIDs[id] {
					// ensure note_ is in _diagramstructure.NotesWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramstructure.NotesWhoseNodeIsExpanded {
						if _b == note_ {
							found = true
							break
						}
					}
					if !found {
						_diagramstructure.NotesWhoseNodeIsExpanded = append(_diagramstructure.NotesWhoseNodeIsExpanded, note_)
						noteFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "NotesWhoseNodeIsExpanded", &_diagramstructure.NotesWhoseNodeIsExpanded)
					}
				} else {
					// ensure note_ is NOT in _diagramstructure.NotesWhoseNodeIsExpanded
					idx := slices.Index(_diagramstructure.NotesWhoseNodeIsExpanded, note_)
					if idx != -1 {
						_diagramstructure.NotesWhoseNodeIsExpanded = slices.Delete(_diagramstructure.NotesWhoseNodeIsExpanded, idx, idx+1)
						noteFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "NotesWhoseNodeIsExpanded", &_diagramstructure.NotesWhoseNodeIsExpanded)
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
func __gong__New__NotePartShapeFormCallback(
	notepartshape *models.NotePartShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (notepartshapeFormCallback *NotePartShapeFormCallback) {
	notepartshapeFormCallback = new(NotePartShapeFormCallback)
	notepartshapeFormCallback.probe = probe
	notepartshapeFormCallback.notepartshape = notepartshape
	notepartshapeFormCallback.formGroup = formGroup

	notepartshapeFormCallback.CreationMode = (notepartshape == nil)

	return
}

type NotePartShapeFormCallback struct {
	notepartshape *models.NotePartShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (notepartshapeFormCallback *NotePartShapeFormCallback) OnSave() {
	notepartshapeFormCallback.probe.stageOfInterest.Lock()
	defer notepartshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("NotePartShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	notepartshapeFormCallback.probe.formStage.Checkout()

	if notepartshapeFormCallback.notepartshape == nil {
		notepartshapeFormCallback.notepartshape = new(models.NotePartShape).Stage(notepartshapeFormCallback.probe.stageOfInterest)
	}
	notepartshape_ := notepartshapeFormCallback.notepartshape
	_ = notepartshape_

	for _, formDiv := range notepartshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(notepartshape_.Name), formDiv)
		case "Note":
			FormDivSelectFieldToField(&(notepartshape_.Note), notepartshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Part":
			FormDivSelectFieldToField(&(notepartshape_.Part), notepartshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(notepartshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(notepartshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(notepartshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(notepartshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(notepartshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(notepartshape_.IsHidden), formDiv)
		case "DiagramStructure:NotePartShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramStructure instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramStructure instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramStructure](notepartshapeFormCallback.probe.stageOfInterest)
			targetDiagramStructureIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramStructureIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramStructure instances and update their NotePartShapes slice
			for _diagramstructure := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramStructure](notepartshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(notepartshapeFormCallback.probe.stageOfInterest, _diagramstructure)
				
				// if DiagramStructure is selected
				if targetDiagramStructureIDs[id] {
					// ensure notepartshape_ is in _diagramstructure.NotePartShapes
					found := false
					for _, _b := range _diagramstructure.NotePartShapes {
						if _b == notepartshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramstructure.NotePartShapes = append(_diagramstructure.NotePartShapes, notepartshape_)
						notepartshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "NotePartShapes", &_diagramstructure.NotePartShapes)
					}
				} else {
					// ensure notepartshape_ is NOT in _diagramstructure.NotePartShapes
					idx := slices.Index(_diagramstructure.NotePartShapes, notepartshape_)
					if idx != -1 {
						_diagramstructure.NotePartShapes = slices.Delete(_diagramstructure.NotePartShapes, idx, idx+1)
						notepartshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "NotePartShapes", &_diagramstructure.NotePartShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if notepartshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		notepartshape_.Unstage(notepartshapeFormCallback.probe.stageOfInterest)
	}

	notepartshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.NotePartShape](
		notepartshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if notepartshapeFormCallback.CreationMode || notepartshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		notepartshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(notepartshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__NotePartShapeFormCallback(
			nil,
			notepartshapeFormCallback.probe,
			newFormGroup,
		)
		notepartshape := new(models.NotePartShape)
		FillUpForm(notepartshape, newFormGroup, notepartshapeFormCallback.probe)
		notepartshapeFormCallback.probe.formStage.Commit()
	}

	notepartshapeFormCallback.probe.ux_tree()
}
func __gong__New__NotePortShapeFormCallback(
	noteportshape *models.NotePortShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (noteportshapeFormCallback *NotePortShapeFormCallback) {
	noteportshapeFormCallback = new(NotePortShapeFormCallback)
	noteportshapeFormCallback.probe = probe
	noteportshapeFormCallback.noteportshape = noteportshape
	noteportshapeFormCallback.formGroup = formGroup

	noteportshapeFormCallback.CreationMode = (noteportshape == nil)

	return
}

type NotePortShapeFormCallback struct {
	noteportshape *models.NotePortShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (noteportshapeFormCallback *NotePortShapeFormCallback) OnSave() {
	noteportshapeFormCallback.probe.stageOfInterest.Lock()
	defer noteportshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("NotePortShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	noteportshapeFormCallback.probe.formStage.Checkout()

	if noteportshapeFormCallback.noteportshape == nil {
		noteportshapeFormCallback.noteportshape = new(models.NotePortShape).Stage(noteportshapeFormCallback.probe.stageOfInterest)
	}
	noteportshape_ := noteportshapeFormCallback.noteportshape
	_ = noteportshape_

	for _, formDiv := range noteportshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(noteportshape_.Name), formDiv)
		case "Note":
			FormDivSelectFieldToField(&(noteportshape_.Note), noteportshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Port":
			FormDivSelectFieldToField(&(noteportshape_.Port), noteportshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(noteportshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(noteportshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(noteportshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(noteportshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(noteportshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(noteportshape_.IsHidden), formDiv)
		case "DiagramStructure:NotePortShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramStructure instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramStructure instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramStructure](noteportshapeFormCallback.probe.stageOfInterest)
			targetDiagramStructureIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramStructureIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramStructure instances and update their NotePortShapes slice
			for _diagramstructure := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramStructure](noteportshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(noteportshapeFormCallback.probe.stageOfInterest, _diagramstructure)
				
				// if DiagramStructure is selected
				if targetDiagramStructureIDs[id] {
					// ensure noteportshape_ is in _diagramstructure.NotePortShapes
					found := false
					for _, _b := range _diagramstructure.NotePortShapes {
						if _b == noteportshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramstructure.NotePortShapes = append(_diagramstructure.NotePortShapes, noteportshape_)
						noteportshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "NotePortShapes", &_diagramstructure.NotePortShapes)
					}
				} else {
					// ensure noteportshape_ is NOT in _diagramstructure.NotePortShapes
					idx := slices.Index(_diagramstructure.NotePortShapes, noteportshape_)
					if idx != -1 {
						_diagramstructure.NotePortShapes = slices.Delete(_diagramstructure.NotePortShapes, idx, idx+1)
						noteportshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "NotePortShapes", &_diagramstructure.NotePortShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if noteportshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteportshape_.Unstage(noteportshapeFormCallback.probe.stageOfInterest)
	}

	noteportshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.NotePortShape](
		noteportshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if noteportshapeFormCallback.CreationMode || noteportshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteportshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(noteportshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__NotePortShapeFormCallback(
			nil,
			noteportshapeFormCallback.probe,
			newFormGroup,
		)
		noteportshape := new(models.NotePortShape)
		FillUpForm(noteportshape, newFormGroup, noteportshapeFormCallback.probe)
		noteportshapeFormCallback.probe.formStage.Commit()
	}

	noteportshapeFormCallback.probe.ux_tree()
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
		case "DiagramStructure:Note_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramStructure instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramStructure instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramStructure](noteshapeFormCallback.probe.stageOfInterest)
			targetDiagramStructureIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramStructureIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramStructure instances and update their Note_Shapes slice
			for _diagramstructure := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramStructure](noteshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(noteshapeFormCallback.probe.stageOfInterest, _diagramstructure)
				
				// if DiagramStructure is selected
				if targetDiagramStructureIDs[id] {
					// ensure noteshape_ is in _diagramstructure.Note_Shapes
					found := false
					for _, _b := range _diagramstructure.Note_Shapes {
						if _b == noteshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramstructure.Note_Shapes = append(_diagramstructure.Note_Shapes, noteshape_)
						noteshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "Note_Shapes", &_diagramstructure.Note_Shapes)
					}
				} else {
					// ensure noteshape_ is NOT in _diagramstructure.Note_Shapes
					idx := slices.Index(_diagramstructure.Note_Shapes, noteshape_)
					if idx != -1 {
						_diagramstructure.Note_Shapes = slices.Delete(_diagramstructure.Note_Shapes, idx, idx+1)
						noteshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "Note_Shapes", &_diagramstructure.Note_Shapes)
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
func __gong__New__PartFormCallback(
	part *models.Part,
	probe *Probe,
	formGroup *form.FormGroup,
) (partFormCallback *PartFormCallback) {
	partFormCallback = new(PartFormCallback)
	partFormCallback.probe = probe
	partFormCallback.part = part
	partFormCallback.formGroup = formGroup

	partFormCallback.CreationMode = (part == nil)

	return
}

type PartFormCallback struct {
	part *models.Part

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (partFormCallback *PartFormCallback) OnSave() {
	partFormCallback.probe.stageOfInterest.Lock()
	defer partFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PartFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	partFormCallback.probe.formStage.Checkout()

	if partFormCallback.part == nil {
		partFormCallback.part = new(models.Part).Stage(partFormCallback.probe.stageOfInterest)
	}
	part_ := partFormCallback.part
	_ = part_

	for _, formDiv := range partFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(part_.Name), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(part_.Description), formDiv)
		case "Ports":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Port](partFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Port, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Port)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					partFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Port](partFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			part_.Ports = instanceSlice
			partFormCallback.probe.UpdateSliceOfPointersCallback(part_, "Ports", &part_.Ports)

		case "TypeOfPart":
			FormDivSelectFieldToField(&(part_.TypeOfPart), partFormCallback.probe.stageOfInterest, formDiv)
		case "IsPartNameNotSystemName":
			FormDivBasicFieldToField(&(part_.IsPartNameNotSystemName), formDiv)
		case "IsControlFlowsNodeExpanded":
			FormDivBasicFieldToField(&(part_.IsControlFlowsNodeExpanded), formDiv)
		case "ControlFlows":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ControlFlow](partFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ControlFlow, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ControlFlow)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					partFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ControlFlow](partFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			part_.ControlFlows = instanceSlice
			partFormCallback.probe.UpdateSliceOfPointersCallback(part_, "ControlFlows", &part_.ControlFlows)

		case "PortWhoseOutControlFlowsNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Port](partFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Port, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Port)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					partFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Port](partFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			part_.PortWhoseOutControlFlowsNodeIsExpanded = instanceSlice
			partFormCallback.probe.UpdateSliceOfPointersCallback(part_, "PortWhoseOutControlFlowsNodeIsExpanded", &part_.PortWhoseOutControlFlowsNodeIsExpanded)

		case "PortWhoseInControlFlowsNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Port](partFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Port, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Port)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					partFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Port](partFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			part_.PortWhoseInControlFlowsNodeIsExpanded = instanceSlice
			partFormCallback.probe.UpdateSliceOfPointersCallback(part_, "PortWhoseInControlFlowsNodeIsExpanded", &part_.PortWhoseInControlFlowsNodeIsExpanded)

		case "IsDataFlowsNodeExpanded":
			FormDivBasicFieldToField(&(part_.IsDataFlowsNodeExpanded), formDiv)
		case "PortWhoseOutDataFlowsNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Port](partFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Port, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Port)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					partFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Port](partFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			part_.PortWhoseOutDataFlowsNodeIsExpanded = instanceSlice
			partFormCallback.probe.UpdateSliceOfPointersCallback(part_, "PortWhoseOutDataFlowsNodeIsExpanded", &part_.PortWhoseOutDataFlowsNodeIsExpanded)

		case "PortWhoseInDataFlowsNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Port](partFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Port, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Port)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					partFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Port](partFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			part_.PortWhoseInDataFlowsNodeIsExpanded = instanceSlice
			partFormCallback.probe.UpdateSliceOfPointersCallback(part_, "PortWhoseInDataFlowsNodeIsExpanded", &part_.PortWhoseInDataFlowsNodeIsExpanded)

		case "PartAnchoredPath":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.PartAnchoredPath](partFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.PartAnchoredPath, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.PartAnchoredPath)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					partFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.PartAnchoredPath](partFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			part_.PartAnchoredPath = instanceSlice
			partFormCallback.probe.UpdateSliceOfPointersCallback(part_, "PartAnchoredPath", &part_.PartAnchoredPath)

		case "ComputedPrefix":
			FormDivBasicFieldToField(&(part_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(part_.IsExpanded), formDiv)
		case "IsPortsNodeExpanded":
			FormDivBasicFieldToField(&(part_.IsPortsNodeExpanded), formDiv)
		case "DiagramStructure:PartWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramStructure instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramStructure instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramStructure](partFormCallback.probe.stageOfInterest)
			targetDiagramStructureIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramStructureIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramStructure instances and update their PartWhoseNodeIsExpanded slice
			for _diagramstructure := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramStructure](partFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(partFormCallback.probe.stageOfInterest, _diagramstructure)
				
				// if DiagramStructure is selected
				if targetDiagramStructureIDs[id] {
					// ensure part_ is in _diagramstructure.PartWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramstructure.PartWhoseNodeIsExpanded {
						if _b == part_ {
							found = true
							break
						}
					}
					if !found {
						_diagramstructure.PartWhoseNodeIsExpanded = append(_diagramstructure.PartWhoseNodeIsExpanded, part_)
						partFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "PartWhoseNodeIsExpanded", &_diagramstructure.PartWhoseNodeIsExpanded)
					}
				} else {
					// ensure part_ is NOT in _diagramstructure.PartWhoseNodeIsExpanded
					idx := slices.Index(_diagramstructure.PartWhoseNodeIsExpanded, part_)
					if idx != -1 {
						_diagramstructure.PartWhoseNodeIsExpanded = slices.Delete(_diagramstructure.PartWhoseNodeIsExpanded, idx, idx+1)
						partFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "PartWhoseNodeIsExpanded", &_diagramstructure.PartWhoseNodeIsExpanded)
					}
				}
			}
		case "DiagramStructure:ExternalPartWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramStructure instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramStructure instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramStructure](partFormCallback.probe.stageOfInterest)
			targetDiagramStructureIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramStructureIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramStructure instances and update their ExternalPartWhoseNodeIsExpanded slice
			for _diagramstructure := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramStructure](partFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(partFormCallback.probe.stageOfInterest, _diagramstructure)
				
				// if DiagramStructure is selected
				if targetDiagramStructureIDs[id] {
					// ensure part_ is in _diagramstructure.ExternalPartWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramstructure.ExternalPartWhoseNodeIsExpanded {
						if _b == part_ {
							found = true
							break
						}
					}
					if !found {
						_diagramstructure.ExternalPartWhoseNodeIsExpanded = append(_diagramstructure.ExternalPartWhoseNodeIsExpanded, part_)
						partFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "ExternalPartWhoseNodeIsExpanded", &_diagramstructure.ExternalPartWhoseNodeIsExpanded)
					}
				} else {
					// ensure part_ is NOT in _diagramstructure.ExternalPartWhoseNodeIsExpanded
					idx := slices.Index(_diagramstructure.ExternalPartWhoseNodeIsExpanded, part_)
					if idx != -1 {
						_diagramstructure.ExternalPartWhoseNodeIsExpanded = slices.Delete(_diagramstructure.ExternalPartWhoseNodeIsExpanded, idx, idx+1)
						partFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "ExternalPartWhoseNodeIsExpanded", &_diagramstructure.ExternalPartWhoseNodeIsExpanded)
					}
				}
			}
		case "DiagramStructure:ExternalPartsWhoseOutDataFlowsNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramStructure instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramStructure instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramStructure](partFormCallback.probe.stageOfInterest)
			targetDiagramStructureIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramStructureIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramStructure instances and update their ExternalPartsWhoseOutDataFlowsNodeIsExpanded slice
			for _diagramstructure := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramStructure](partFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(partFormCallback.probe.stageOfInterest, _diagramstructure)
				
				// if DiagramStructure is selected
				if targetDiagramStructureIDs[id] {
					// ensure part_ is in _diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded
					found := false
					for _, _b := range _diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded {
						if _b == part_ {
							found = true
							break
						}
					}
					if !found {
						_diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded = append(_diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded, part_)
						partFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "ExternalPartsWhoseOutDataFlowsNodeIsExpanded", &_diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded)
					}
				} else {
					// ensure part_ is NOT in _diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded
					idx := slices.Index(_diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded, part_)
					if idx != -1 {
						_diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded = slices.Delete(_diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded, idx, idx+1)
						partFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "ExternalPartsWhoseOutDataFlowsNodeIsExpanded", &_diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded)
					}
				}
			}
		case "DiagramStructure:ExternalPartsWhoseInDataFlowsNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramStructure instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramStructure instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramStructure](partFormCallback.probe.stageOfInterest)
			targetDiagramStructureIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramStructureIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramStructure instances and update their ExternalPartsWhoseInDataFlowsNodeIsExpanded slice
			for _diagramstructure := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramStructure](partFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(partFormCallback.probe.stageOfInterest, _diagramstructure)
				
				// if DiagramStructure is selected
				if targetDiagramStructureIDs[id] {
					// ensure part_ is in _diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded
					found := false
					for _, _b := range _diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded {
						if _b == part_ {
							found = true
							break
						}
					}
					if !found {
						_diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded = append(_diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded, part_)
						partFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "ExternalPartsWhoseInDataFlowsNodeIsExpanded", &_diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded)
					}
				} else {
					// ensure part_ is NOT in _diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded
					idx := slices.Index(_diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded, part_)
					if idx != -1 {
						_diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded = slices.Delete(_diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded, idx, idx+1)
						partFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "ExternalPartsWhoseInDataFlowsNodeIsExpanded", &_diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded)
					}
				}
			}
		case "Library:PartsWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](partFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their PartsWhoseNodeIsExpanded slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](partFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(partFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure part_ is in _library.PartsWhoseNodeIsExpanded
					found := false
					for _, _b := range _library.PartsWhoseNodeIsExpanded {
						if _b == part_ {
							found = true
							break
						}
					}
					if !found {
						_library.PartsWhoseNodeIsExpanded = append(_library.PartsWhoseNodeIsExpanded, part_)
						partFormCallback.probe.UpdateSliceOfPointersCallback(_library, "PartsWhoseNodeIsExpanded", &_library.PartsWhoseNodeIsExpanded)
					}
				} else {
					// ensure part_ is NOT in _library.PartsWhoseNodeIsExpanded
					idx := slices.Index(_library.PartsWhoseNodeIsExpanded, part_)
					if idx != -1 {
						_library.PartsWhoseNodeIsExpanded = slices.Delete(_library.PartsWhoseNodeIsExpanded, idx, idx+1)
						partFormCallback.probe.UpdateSliceOfPointersCallback(_library, "PartsWhoseNodeIsExpanded", &_library.PartsWhoseNodeIsExpanded)
					}
				}
			}
		case "Note:Parts":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Note instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Note instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Note](partFormCallback.probe.stageOfInterest)
			targetNoteIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetNoteIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Note instances and update their Parts slice
			for _note := range *models.GetGongstructInstancesSetFromPointerType[*models.Note](partFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(partFormCallback.probe.stageOfInterest, _note)
				
				// if Note is selected
				if targetNoteIDs[id] {
					// ensure part_ is in _note.Parts
					found := false
					for _, _b := range _note.Parts {
						if _b == part_ {
							found = true
							break
						}
					}
					if !found {
						_note.Parts = append(_note.Parts, part_)
						partFormCallback.probe.UpdateSliceOfPointersCallback(_note, "Parts", &_note.Parts)
					}
				} else {
					// ensure part_ is NOT in _note.Parts
					idx := slices.Index(_note.Parts, part_)
					if idx != -1 {
						_note.Parts = slices.Delete(_note.Parts, idx, idx+1)
						partFormCallback.probe.UpdateSliceOfPointersCallback(_note, "Parts", &_note.Parts)
					}
				}
			}
		case "System:Parts":
			// 1. Decode the AssociationStorage which contains the rowIDs of the System instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target System instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.System](partFormCallback.probe.stageOfInterest)
			targetSystemIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSystemIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all System instances and update their Parts slice
			for _system := range *models.GetGongstructInstancesSetFromPointerType[*models.System](partFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(partFormCallback.probe.stageOfInterest, _system)
				
				// if System is selected
				if targetSystemIDs[id] {
					// ensure part_ is in _system.Parts
					found := false
					for _, _b := range _system.Parts {
						if _b == part_ {
							found = true
							break
						}
					}
					if !found {
						_system.Parts = append(_system.Parts, part_)
						partFormCallback.probe.UpdateSliceOfPointersCallback(_system, "Parts", &_system.Parts)
					}
				} else {
					// ensure part_ is NOT in _system.Parts
					idx := slices.Index(_system.Parts, part_)
					if idx != -1 {
						_system.Parts = slices.Delete(_system.Parts, idx, idx+1)
						partFormCallback.probe.UpdateSliceOfPointersCallback(_system, "Parts", &_system.Parts)
					}
				}
			}
		case "System:PartWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the System instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target System instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.System](partFormCallback.probe.stageOfInterest)
			targetSystemIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSystemIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all System instances and update their PartWhoseNodeIsExpanded slice
			for _system := range *models.GetGongstructInstancesSetFromPointerType[*models.System](partFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(partFormCallback.probe.stageOfInterest, _system)
				
				// if System is selected
				if targetSystemIDs[id] {
					// ensure part_ is in _system.PartWhoseNodeIsExpanded
					found := false
					for _, _b := range _system.PartWhoseNodeIsExpanded {
						if _b == part_ {
							found = true
							break
						}
					}
					if !found {
						_system.PartWhoseNodeIsExpanded = append(_system.PartWhoseNodeIsExpanded, part_)
						partFormCallback.probe.UpdateSliceOfPointersCallback(_system, "PartWhoseNodeIsExpanded", &_system.PartWhoseNodeIsExpanded)
					}
				} else {
					// ensure part_ is NOT in _system.PartWhoseNodeIsExpanded
					idx := slices.Index(_system.PartWhoseNodeIsExpanded, part_)
					if idx != -1 {
						_system.PartWhoseNodeIsExpanded = slices.Delete(_system.PartWhoseNodeIsExpanded, idx, idx+1)
						partFormCallback.probe.UpdateSliceOfPointersCallback(_system, "PartWhoseNodeIsExpanded", &_system.PartWhoseNodeIsExpanded)
					}
				}
			}
		case "System:ExternalParts":
			// 1. Decode the AssociationStorage which contains the rowIDs of the System instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target System instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.System](partFormCallback.probe.stageOfInterest)
			targetSystemIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSystemIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all System instances and update their ExternalParts slice
			for _system := range *models.GetGongstructInstancesSetFromPointerType[*models.System](partFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(partFormCallback.probe.stageOfInterest, _system)
				
				// if System is selected
				if targetSystemIDs[id] {
					// ensure part_ is in _system.ExternalParts
					found := false
					for _, _b := range _system.ExternalParts {
						if _b == part_ {
							found = true
							break
						}
					}
					if !found {
						_system.ExternalParts = append(_system.ExternalParts, part_)
						partFormCallback.probe.UpdateSliceOfPointersCallback(_system, "ExternalParts", &_system.ExternalParts)
					}
				} else {
					// ensure part_ is NOT in _system.ExternalParts
					idx := slices.Index(_system.ExternalParts, part_)
					if idx != -1 {
						_system.ExternalParts = slices.Delete(_system.ExternalParts, idx, idx+1)
						partFormCallback.probe.UpdateSliceOfPointersCallback(_system, "ExternalParts", &_system.ExternalParts)
					}
				}
			}
		case "System:ExternalPartWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the System instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target System instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.System](partFormCallback.probe.stageOfInterest)
			targetSystemIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSystemIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all System instances and update their ExternalPartWhoseNodeIsExpanded slice
			for _system := range *models.GetGongstructInstancesSetFromPointerType[*models.System](partFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(partFormCallback.probe.stageOfInterest, _system)
				
				// if System is selected
				if targetSystemIDs[id] {
					// ensure part_ is in _system.ExternalPartWhoseNodeIsExpanded
					found := false
					for _, _b := range _system.ExternalPartWhoseNodeIsExpanded {
						if _b == part_ {
							found = true
							break
						}
					}
					if !found {
						_system.ExternalPartWhoseNodeIsExpanded = append(_system.ExternalPartWhoseNodeIsExpanded, part_)
						partFormCallback.probe.UpdateSliceOfPointersCallback(_system, "ExternalPartWhoseNodeIsExpanded", &_system.ExternalPartWhoseNodeIsExpanded)
					}
				} else {
					// ensure part_ is NOT in _system.ExternalPartWhoseNodeIsExpanded
					idx := slices.Index(_system.ExternalPartWhoseNodeIsExpanded, part_)
					if idx != -1 {
						_system.ExternalPartWhoseNodeIsExpanded = slices.Delete(_system.ExternalPartWhoseNodeIsExpanded, idx, idx+1)
						partFormCallback.probe.UpdateSliceOfPointersCallback(_system, "ExternalPartWhoseNodeIsExpanded", &_system.ExternalPartWhoseNodeIsExpanded)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if partFormCallback.formGroup.HasSuppressButtonBeenPressed {
		part_.Unstage(partFormCallback.probe.stageOfInterest)
	}

	partFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Part](
		partFormCallback.probe,
	)

	// display a new form by reset the form stage
	if partFormCallback.CreationMode || partFormCallback.formGroup.HasSuppressButtonBeenPressed {
		partFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(partFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PartFormCallback(
			nil,
			partFormCallback.probe,
			newFormGroup,
		)
		part := new(models.Part)
		FillUpForm(part, newFormGroup, partFormCallback.probe)
		partFormCallback.probe.formStage.Commit()
	}

	partFormCallback.probe.ux_tree()
}
func __gong__New__PartAnchoredPathFormCallback(
	partanchoredpath *models.PartAnchoredPath,
	probe *Probe,
	formGroup *form.FormGroup,
) (partanchoredpathFormCallback *PartAnchoredPathFormCallback) {
	partanchoredpathFormCallback = new(PartAnchoredPathFormCallback)
	partanchoredpathFormCallback.probe = probe
	partanchoredpathFormCallback.partanchoredpath = partanchoredpath
	partanchoredpathFormCallback.formGroup = formGroup

	partanchoredpathFormCallback.CreationMode = (partanchoredpath == nil)

	return
}

type PartAnchoredPathFormCallback struct {
	partanchoredpath *models.PartAnchoredPath

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (partanchoredpathFormCallback *PartAnchoredPathFormCallback) OnSave() {
	partanchoredpathFormCallback.probe.stageOfInterest.Lock()
	defer partanchoredpathFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PartAnchoredPathFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	partanchoredpathFormCallback.probe.formStage.Checkout()

	if partanchoredpathFormCallback.partanchoredpath == nil {
		partanchoredpathFormCallback.partanchoredpath = new(models.PartAnchoredPath).Stage(partanchoredpathFormCallback.probe.stageOfInterest)
	}
	partanchoredpath_ := partanchoredpathFormCallback.partanchoredpath
	_ = partanchoredpath_

	for _, formDiv := range partanchoredpathFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(partanchoredpath_.Name), formDiv)
		case "Definition":
			FormDivBasicFieldToField(&(partanchoredpath_.Definition), formDiv)
		case "X_Offset":
			FormDivBasicFieldToField(&(partanchoredpath_.X_Offset), formDiv)
		case "Y_Offset":
			FormDivBasicFieldToField(&(partanchoredpath_.Y_Offset), formDiv)
		case "RectAnchorType":
			FormDivEnumStringFieldToField(&(partanchoredpath_.RectAnchorType), formDiv)
		case "ScalePropotionnally":
			FormDivBasicFieldToField(&(partanchoredpath_.ScalePropotionnally), formDiv)
		case "AppliedScaling":
			FormDivBasicFieldToField(&(partanchoredpath_.AppliedScaling), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(partanchoredpath_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(partanchoredpath_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(partanchoredpath_.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(partanchoredpath_.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(partanchoredpath_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(partanchoredpath_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(partanchoredpath_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(partanchoredpath_.Transform), formDiv)
		case "Part:PartAnchoredPath":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Part instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Part instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Part](partanchoredpathFormCallback.probe.stageOfInterest)
			targetPartIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetPartIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Part instances and update their PartAnchoredPath slice
			for _part := range *models.GetGongstructInstancesSetFromPointerType[*models.Part](partanchoredpathFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(partanchoredpathFormCallback.probe.stageOfInterest, _part)
				
				// if Part is selected
				if targetPartIDs[id] {
					// ensure partanchoredpath_ is in _part.PartAnchoredPath
					found := false
					for _, _b := range _part.PartAnchoredPath {
						if _b == partanchoredpath_ {
							found = true
							break
						}
					}
					if !found {
						_part.PartAnchoredPath = append(_part.PartAnchoredPath, partanchoredpath_)
						partanchoredpathFormCallback.probe.UpdateSliceOfPointersCallback(_part, "PartAnchoredPath", &_part.PartAnchoredPath)
					}
				} else {
					// ensure partanchoredpath_ is NOT in _part.PartAnchoredPath
					idx := slices.Index(_part.PartAnchoredPath, partanchoredpath_)
					if idx != -1 {
						_part.PartAnchoredPath = slices.Delete(_part.PartAnchoredPath, idx, idx+1)
						partanchoredpathFormCallback.probe.UpdateSliceOfPointersCallback(_part, "PartAnchoredPath", &_part.PartAnchoredPath)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if partanchoredpathFormCallback.formGroup.HasSuppressButtonBeenPressed {
		partanchoredpath_.Unstage(partanchoredpathFormCallback.probe.stageOfInterest)
	}

	partanchoredpathFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.PartAnchoredPath](
		partanchoredpathFormCallback.probe,
	)

	// display a new form by reset the form stage
	if partanchoredpathFormCallback.CreationMode || partanchoredpathFormCallback.formGroup.HasSuppressButtonBeenPressed {
		partanchoredpathFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(partanchoredpathFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PartAnchoredPathFormCallback(
			nil,
			partanchoredpathFormCallback.probe,
			newFormGroup,
		)
		partanchoredpath := new(models.PartAnchoredPath)
		FillUpForm(partanchoredpath, newFormGroup, partanchoredpathFormCallback.probe)
		partanchoredpathFormCallback.probe.formStage.Commit()
	}

	partanchoredpathFormCallback.probe.ux_tree()
}
func __gong__New__PartShapeFormCallback(
	partshape *models.PartShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (partshapeFormCallback *PartShapeFormCallback) {
	partshapeFormCallback = new(PartShapeFormCallback)
	partshapeFormCallback.probe = probe
	partshapeFormCallback.partshape = partshape
	partshapeFormCallback.formGroup = formGroup

	partshapeFormCallback.CreationMode = (partshape == nil)

	return
}

type PartShapeFormCallback struct {
	partshape *models.PartShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (partshapeFormCallback *PartShapeFormCallback) OnSave() {
	partshapeFormCallback.probe.stageOfInterest.Lock()
	defer partshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PartShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	partshapeFormCallback.probe.formStage.Checkout()

	if partshapeFormCallback.partshape == nil {
		partshapeFormCallback.partshape = new(models.PartShape).Stage(partshapeFormCallback.probe.stageOfInterest)
	}
	partshape_ := partshapeFormCallback.partshape
	_ = partshape_

	for _, formDiv := range partshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(partshape_.Name), formDiv)
		case "Part":
			FormDivSelectFieldToField(&(partshape_.Part), partshapeFormCallback.probe.stageOfInterest, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(partshape_.IsExpanded), formDiv)
		case "X":
			FormDivBasicFieldToField(&(partshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(partshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(partshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(partshape_.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(partshape_.IsHidden), formDiv)
		case "DiagramStructure:Part_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramStructure instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramStructure instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramStructure](partshapeFormCallback.probe.stageOfInterest)
			targetDiagramStructureIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramStructureIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramStructure instances and update their Part_Shapes slice
			for _diagramstructure := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramStructure](partshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(partshapeFormCallback.probe.stageOfInterest, _diagramstructure)
				
				// if DiagramStructure is selected
				if targetDiagramStructureIDs[id] {
					// ensure partshape_ is in _diagramstructure.Part_Shapes
					found := false
					for _, _b := range _diagramstructure.Part_Shapes {
						if _b == partshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramstructure.Part_Shapes = append(_diagramstructure.Part_Shapes, partshape_)
						partshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "Part_Shapes", &_diagramstructure.Part_Shapes)
					}
				} else {
					// ensure partshape_ is NOT in _diagramstructure.Part_Shapes
					idx := slices.Index(_diagramstructure.Part_Shapes, partshape_)
					if idx != -1 {
						_diagramstructure.Part_Shapes = slices.Delete(_diagramstructure.Part_Shapes, idx, idx+1)
						partshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "Part_Shapes", &_diagramstructure.Part_Shapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if partshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		partshape_.Unstage(partshapeFormCallback.probe.stageOfInterest)
	}

	partshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.PartShape](
		partshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if partshapeFormCallback.CreationMode || partshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		partshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(partshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PartShapeFormCallback(
			nil,
			partshapeFormCallback.probe,
			newFormGroup,
		)
		partshape := new(models.PartShape)
		FillUpForm(partshape, newFormGroup, partshapeFormCallback.probe)
		partshapeFormCallback.probe.formStage.Commit()
	}

	partshapeFormCallback.probe.ux_tree()
}
func __gong__New__PortFormCallback(
	port *models.Port,
	probe *Probe,
	formGroup *form.FormGroup,
) (portFormCallback *PortFormCallback) {
	portFormCallback = new(PortFormCallback)
	portFormCallback.probe = probe
	portFormCallback.port = port
	portFormCallback.formGroup = formGroup

	portFormCallback.CreationMode = (port == nil)

	return
}

type PortFormCallback struct {
	port *models.Port

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (portFormCallback *PortFormCallback) OnSave() {
	portFormCallback.probe.stageOfInterest.Lock()
	defer portFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PortFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	portFormCallback.probe.formStage.Checkout()

	if portFormCallback.port == nil {
		portFormCallback.port = new(models.Port).Stage(portFormCallback.probe.stageOfInterest)
	}
	port_ := portFormCallback.port
	_ = port_

	for _, formDiv := range portFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(port_.Name), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(port_.Description), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(port_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(port_.IsExpanded), formDiv)
		case "DiagramStructure:PortsWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramStructure instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramStructure instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramStructure](portFormCallback.probe.stageOfInterest)
			targetDiagramStructureIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramStructureIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramStructure instances and update their PortsWhoseNodeIsExpanded slice
			for _diagramstructure := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramStructure](portFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(portFormCallback.probe.stageOfInterest, _diagramstructure)
				
				// if DiagramStructure is selected
				if targetDiagramStructureIDs[id] {
					// ensure port_ is in _diagramstructure.PortsWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramstructure.PortsWhoseNodeIsExpanded {
						if _b == port_ {
							found = true
							break
						}
					}
					if !found {
						_diagramstructure.PortsWhoseNodeIsExpanded = append(_diagramstructure.PortsWhoseNodeIsExpanded, port_)
						portFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "PortsWhoseNodeIsExpanded", &_diagramstructure.PortsWhoseNodeIsExpanded)
					}
				} else {
					// ensure port_ is NOT in _diagramstructure.PortsWhoseNodeIsExpanded
					idx := slices.Index(_diagramstructure.PortsWhoseNodeIsExpanded, port_)
					if idx != -1 {
						_diagramstructure.PortsWhoseNodeIsExpanded = slices.Delete(_diagramstructure.PortsWhoseNodeIsExpanded, idx, idx+1)
						portFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "PortsWhoseNodeIsExpanded", &_diagramstructure.PortsWhoseNodeIsExpanded)
					}
				}
			}
		case "Note:Ports":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Note instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Note instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Note](portFormCallback.probe.stageOfInterest)
			targetNoteIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetNoteIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Note instances and update their Ports slice
			for _note := range *models.GetGongstructInstancesSetFromPointerType[*models.Note](portFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(portFormCallback.probe.stageOfInterest, _note)
				
				// if Note is selected
				if targetNoteIDs[id] {
					// ensure port_ is in _note.Ports
					found := false
					for _, _b := range _note.Ports {
						if _b == port_ {
							found = true
							break
						}
					}
					if !found {
						_note.Ports = append(_note.Ports, port_)
						portFormCallback.probe.UpdateSliceOfPointersCallback(_note, "Ports", &_note.Ports)
					}
				} else {
					// ensure port_ is NOT in _note.Ports
					idx := slices.Index(_note.Ports, port_)
					if idx != -1 {
						_note.Ports = slices.Delete(_note.Ports, idx, idx+1)
						portFormCallback.probe.UpdateSliceOfPointersCallback(_note, "Ports", &_note.Ports)
					}
				}
			}
		case "Part:Ports":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Part instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Part instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Part](portFormCallback.probe.stageOfInterest)
			targetPartIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetPartIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Part instances and update their Ports slice
			for _part := range *models.GetGongstructInstancesSetFromPointerType[*models.Part](portFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(portFormCallback.probe.stageOfInterest, _part)
				
				// if Part is selected
				if targetPartIDs[id] {
					// ensure port_ is in _part.Ports
					found := false
					for _, _b := range _part.Ports {
						if _b == port_ {
							found = true
							break
						}
					}
					if !found {
						_part.Ports = append(_part.Ports, port_)
						portFormCallback.probe.UpdateSliceOfPointersCallback(_part, "Ports", &_part.Ports)
					}
				} else {
					// ensure port_ is NOT in _part.Ports
					idx := slices.Index(_part.Ports, port_)
					if idx != -1 {
						_part.Ports = slices.Delete(_part.Ports, idx, idx+1)
						portFormCallback.probe.UpdateSliceOfPointersCallback(_part, "Ports", &_part.Ports)
					}
				}
			}
		case "Part:PortWhoseOutControlFlowsNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Part instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Part instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Part](portFormCallback.probe.stageOfInterest)
			targetPartIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetPartIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Part instances and update their PortWhoseOutControlFlowsNodeIsExpanded slice
			for _part := range *models.GetGongstructInstancesSetFromPointerType[*models.Part](portFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(portFormCallback.probe.stageOfInterest, _part)
				
				// if Part is selected
				if targetPartIDs[id] {
					// ensure port_ is in _part.PortWhoseOutControlFlowsNodeIsExpanded
					found := false
					for _, _b := range _part.PortWhoseOutControlFlowsNodeIsExpanded {
						if _b == port_ {
							found = true
							break
						}
					}
					if !found {
						_part.PortWhoseOutControlFlowsNodeIsExpanded = append(_part.PortWhoseOutControlFlowsNodeIsExpanded, port_)
						portFormCallback.probe.UpdateSliceOfPointersCallback(_part, "PortWhoseOutControlFlowsNodeIsExpanded", &_part.PortWhoseOutControlFlowsNodeIsExpanded)
					}
				} else {
					// ensure port_ is NOT in _part.PortWhoseOutControlFlowsNodeIsExpanded
					idx := slices.Index(_part.PortWhoseOutControlFlowsNodeIsExpanded, port_)
					if idx != -1 {
						_part.PortWhoseOutControlFlowsNodeIsExpanded = slices.Delete(_part.PortWhoseOutControlFlowsNodeIsExpanded, idx, idx+1)
						portFormCallback.probe.UpdateSliceOfPointersCallback(_part, "PortWhoseOutControlFlowsNodeIsExpanded", &_part.PortWhoseOutControlFlowsNodeIsExpanded)
					}
				}
			}
		case "Part:PortWhoseInControlFlowsNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Part instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Part instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Part](portFormCallback.probe.stageOfInterest)
			targetPartIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetPartIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Part instances and update their PortWhoseInControlFlowsNodeIsExpanded slice
			for _part := range *models.GetGongstructInstancesSetFromPointerType[*models.Part](portFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(portFormCallback.probe.stageOfInterest, _part)
				
				// if Part is selected
				if targetPartIDs[id] {
					// ensure port_ is in _part.PortWhoseInControlFlowsNodeIsExpanded
					found := false
					for _, _b := range _part.PortWhoseInControlFlowsNodeIsExpanded {
						if _b == port_ {
							found = true
							break
						}
					}
					if !found {
						_part.PortWhoseInControlFlowsNodeIsExpanded = append(_part.PortWhoseInControlFlowsNodeIsExpanded, port_)
						portFormCallback.probe.UpdateSliceOfPointersCallback(_part, "PortWhoseInControlFlowsNodeIsExpanded", &_part.PortWhoseInControlFlowsNodeIsExpanded)
					}
				} else {
					// ensure port_ is NOT in _part.PortWhoseInControlFlowsNodeIsExpanded
					idx := slices.Index(_part.PortWhoseInControlFlowsNodeIsExpanded, port_)
					if idx != -1 {
						_part.PortWhoseInControlFlowsNodeIsExpanded = slices.Delete(_part.PortWhoseInControlFlowsNodeIsExpanded, idx, idx+1)
						portFormCallback.probe.UpdateSliceOfPointersCallback(_part, "PortWhoseInControlFlowsNodeIsExpanded", &_part.PortWhoseInControlFlowsNodeIsExpanded)
					}
				}
			}
		case "Part:PortWhoseOutDataFlowsNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Part instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Part instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Part](portFormCallback.probe.stageOfInterest)
			targetPartIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetPartIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Part instances and update their PortWhoseOutDataFlowsNodeIsExpanded slice
			for _part := range *models.GetGongstructInstancesSetFromPointerType[*models.Part](portFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(portFormCallback.probe.stageOfInterest, _part)
				
				// if Part is selected
				if targetPartIDs[id] {
					// ensure port_ is in _part.PortWhoseOutDataFlowsNodeIsExpanded
					found := false
					for _, _b := range _part.PortWhoseOutDataFlowsNodeIsExpanded {
						if _b == port_ {
							found = true
							break
						}
					}
					if !found {
						_part.PortWhoseOutDataFlowsNodeIsExpanded = append(_part.PortWhoseOutDataFlowsNodeIsExpanded, port_)
						portFormCallback.probe.UpdateSliceOfPointersCallback(_part, "PortWhoseOutDataFlowsNodeIsExpanded", &_part.PortWhoseOutDataFlowsNodeIsExpanded)
					}
				} else {
					// ensure port_ is NOT in _part.PortWhoseOutDataFlowsNodeIsExpanded
					idx := slices.Index(_part.PortWhoseOutDataFlowsNodeIsExpanded, port_)
					if idx != -1 {
						_part.PortWhoseOutDataFlowsNodeIsExpanded = slices.Delete(_part.PortWhoseOutDataFlowsNodeIsExpanded, idx, idx+1)
						portFormCallback.probe.UpdateSliceOfPointersCallback(_part, "PortWhoseOutDataFlowsNodeIsExpanded", &_part.PortWhoseOutDataFlowsNodeIsExpanded)
					}
				}
			}
		case "Part:PortWhoseInDataFlowsNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Part instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Part instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Part](portFormCallback.probe.stageOfInterest)
			targetPartIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetPartIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Part instances and update their PortWhoseInDataFlowsNodeIsExpanded slice
			for _part := range *models.GetGongstructInstancesSetFromPointerType[*models.Part](portFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(portFormCallback.probe.stageOfInterest, _part)
				
				// if Part is selected
				if targetPartIDs[id] {
					// ensure port_ is in _part.PortWhoseInDataFlowsNodeIsExpanded
					found := false
					for _, _b := range _part.PortWhoseInDataFlowsNodeIsExpanded {
						if _b == port_ {
							found = true
							break
						}
					}
					if !found {
						_part.PortWhoseInDataFlowsNodeIsExpanded = append(_part.PortWhoseInDataFlowsNodeIsExpanded, port_)
						portFormCallback.probe.UpdateSliceOfPointersCallback(_part, "PortWhoseInDataFlowsNodeIsExpanded", &_part.PortWhoseInDataFlowsNodeIsExpanded)
					}
				} else {
					// ensure port_ is NOT in _part.PortWhoseInDataFlowsNodeIsExpanded
					idx := slices.Index(_part.PortWhoseInDataFlowsNodeIsExpanded, port_)
					if idx != -1 {
						_part.PortWhoseInDataFlowsNodeIsExpanded = slices.Delete(_part.PortWhoseInDataFlowsNodeIsExpanded, idx, idx+1)
						portFormCallback.probe.UpdateSliceOfPointersCallback(_part, "PortWhoseInDataFlowsNodeIsExpanded", &_part.PortWhoseInDataFlowsNodeIsExpanded)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if portFormCallback.formGroup.HasSuppressButtonBeenPressed {
		port_.Unstage(portFormCallback.probe.stageOfInterest)
	}

	portFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Port](
		portFormCallback.probe,
	)

	// display a new form by reset the form stage
	if portFormCallback.CreationMode || portFormCallback.formGroup.HasSuppressButtonBeenPressed {
		portFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(portFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PortFormCallback(
			nil,
			portFormCallback.probe,
			newFormGroup,
		)
		port := new(models.Port)
		FillUpForm(port, newFormGroup, portFormCallback.probe)
		portFormCallback.probe.formStage.Commit()
	}

	portFormCallback.probe.ux_tree()
}
func __gong__New__PortShapeFormCallback(
	portshape *models.PortShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (portshapeFormCallback *PortShapeFormCallback) {
	portshapeFormCallback = new(PortShapeFormCallback)
	portshapeFormCallback.probe = probe
	portshapeFormCallback.portshape = portshape
	portshapeFormCallback.formGroup = formGroup

	portshapeFormCallback.CreationMode = (portshape == nil)

	return
}

type PortShapeFormCallback struct {
	portshape *models.PortShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (portshapeFormCallback *PortShapeFormCallback) OnSave() {
	portshapeFormCallback.probe.stageOfInterest.Lock()
	defer portshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PortShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	portshapeFormCallback.probe.formStage.Checkout()

	if portshapeFormCallback.portshape == nil {
		portshapeFormCallback.portshape = new(models.PortShape).Stage(portshapeFormCallback.probe.stageOfInterest)
	}
	portshape_ := portshapeFormCallback.portshape
	_ = portshape_

	for _, formDiv := range portshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(portshape_.Name), formDiv)
		case "Port":
			FormDivSelectFieldToField(&(portshape_.Port), portshapeFormCallback.probe.stageOfInterest, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(portshape_.IsExpanded), formDiv)
		case "X":
			FormDivBasicFieldToField(&(portshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(portshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(portshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(portshape_.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(portshape_.IsHidden), formDiv)
		case "DiagramStructure:Port_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramStructure instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramStructure instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramStructure](portshapeFormCallback.probe.stageOfInterest)
			targetDiagramStructureIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramStructureIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramStructure instances and update their Port_Shapes slice
			for _diagramstructure := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramStructure](portshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(portshapeFormCallback.probe.stageOfInterest, _diagramstructure)
				
				// if DiagramStructure is selected
				if targetDiagramStructureIDs[id] {
					// ensure portshape_ is in _diagramstructure.Port_Shapes
					found := false
					for _, _b := range _diagramstructure.Port_Shapes {
						if _b == portshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramstructure.Port_Shapes = append(_diagramstructure.Port_Shapes, portshape_)
						portshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "Port_Shapes", &_diagramstructure.Port_Shapes)
					}
				} else {
					// ensure portshape_ is NOT in _diagramstructure.Port_Shapes
					idx := slices.Index(_diagramstructure.Port_Shapes, portshape_)
					if idx != -1 {
						_diagramstructure.Port_Shapes = slices.Delete(_diagramstructure.Port_Shapes, idx, idx+1)
						portshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "Port_Shapes", &_diagramstructure.Port_Shapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if portshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		portshape_.Unstage(portshapeFormCallback.probe.stageOfInterest)
	}

	portshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.PortShape](
		portshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if portshapeFormCallback.CreationMode || portshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		portshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(portshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PortShapeFormCallback(
			nil,
			portshapeFormCallback.probe,
			newFormGroup,
		)
		portshape := new(models.PortShape)
		FillUpForm(portshape, newFormGroup, portshapeFormCallback.probe)
		portshapeFormCallback.probe.formStage.Commit()
	}

	portshapeFormCallback.probe.ux_tree()
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
		case "IsExpanded":
			FormDivBasicFieldToField(&(resource_.IsExpanded), formDiv)
		case "SVG_Path":
			FormDivBasicFieldToField(&(resource_.SVG_Path), formDiv)
		case "InverseAppliedScaling":
			FormDivBasicFieldToField(&(resource_.InverseAppliedScaling), formDiv)
		case "DiagramStructure:AllocatedResourcesWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramStructure instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramStructure instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramStructure](resourceFormCallback.probe.stageOfInterest)
			targetDiagramStructureIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramStructureIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramStructure instances and update their AllocatedResourcesWhoseNodeIsExpanded slice
			for _diagramstructure := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramStructure](resourceFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(resourceFormCallback.probe.stageOfInterest, _diagramstructure)
				
				// if DiagramStructure is selected
				if targetDiagramStructureIDs[id] {
					// ensure resource_ is in _diagramstructure.AllocatedResourcesWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramstructure.AllocatedResourcesWhoseNodeIsExpanded {
						if _b == resource_ {
							found = true
							break
						}
					}
					if !found {
						_diagramstructure.AllocatedResourcesWhoseNodeIsExpanded = append(_diagramstructure.AllocatedResourcesWhoseNodeIsExpanded, resource_)
						resourceFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "AllocatedResourcesWhoseNodeIsExpanded", &_diagramstructure.AllocatedResourcesWhoseNodeIsExpanded)
					}
				} else {
					// ensure resource_ is NOT in _diagramstructure.AllocatedResourcesWhoseNodeIsExpanded
					idx := slices.Index(_diagramstructure.AllocatedResourcesWhoseNodeIsExpanded, resource_)
					if idx != -1 {
						_diagramstructure.AllocatedResourcesWhoseNodeIsExpanded = slices.Delete(_diagramstructure.AllocatedResourcesWhoseNodeIsExpanded, idx, idx+1)
						resourceFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "AllocatedResourcesWhoseNodeIsExpanded", &_diagramstructure.AllocatedResourcesWhoseNodeIsExpanded)
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
func __gong__New__SystemFormCallback(
	system *models.System,
	probe *Probe,
	formGroup *form.FormGroup,
) (systemFormCallback *SystemFormCallback) {
	systemFormCallback = new(SystemFormCallback)
	systemFormCallback.probe = probe
	systemFormCallback.system = system
	systemFormCallback.formGroup = formGroup

	systemFormCallback.CreationMode = (system == nil)

	return
}

type SystemFormCallback struct {
	system *models.System

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (systemFormCallback *SystemFormCallback) OnSave() {
	systemFormCallback.probe.stageOfInterest.Lock()
	defer systemFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("SystemFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	systemFormCallback.probe.formStage.Checkout()

	if systemFormCallback.system == nil {
		systemFormCallback.system = new(models.System).Stage(systemFormCallback.probe.stageOfInterest)
	}
	system_ := systemFormCallback.system
	_ = system_

	for _, formDiv := range systemFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(system_.Name), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(system_.Description), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(system_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(system_.IsExpanded), formDiv)
		case "SVG_Path":
			FormDivBasicFieldToField(&(system_.SVG_Path), formDiv)
		case "InverseAppliedScaling":
			FormDivBasicFieldToField(&(system_.InverseAppliedScaling), formDiv)
		case "DiagramStructures":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DiagramStructure](systemFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DiagramStructure, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DiagramStructure)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					systemFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramStructure](systemFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			system_.DiagramStructures = instanceSlice
			systemFormCallback.probe.UpdateSliceOfPointersCallback(system_, "DiagramStructures", &system_.DiagramStructures)

		case "DiagramStructureWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DiagramStructure](systemFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DiagramStructure, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DiagramStructure)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					systemFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramStructure](systemFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			system_.DiagramStructureWhoseNodeIsExpanded = instanceSlice
			systemFormCallback.probe.UpdateSliceOfPointersCallback(system_, "DiagramStructureWhoseNodeIsExpanded", &system_.DiagramStructureWhoseNodeIsExpanded)

		case "IsSubSystemNodeExpanded":
			FormDivBasicFieldToField(&(system_.IsSubSystemNodeExpanded), formDiv)
		case "SubSystemes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.System](systemFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.System, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.System)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					systemFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.System](systemFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			system_.SubSystemes = instanceSlice
			systemFormCallback.probe.UpdateSliceOfPointersCallback(system_, "SubSystemes", &system_.SubSystemes)

		case "Parts":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Part](systemFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Part, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Part)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					systemFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Part](systemFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			system_.Parts = instanceSlice
			systemFormCallback.probe.UpdateSliceOfPointersCallback(system_, "Parts", &system_.Parts)

		case "PartWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Part](systemFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Part, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Part)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					systemFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Part](systemFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			system_.PartWhoseNodeIsExpanded = instanceSlice
			systemFormCallback.probe.UpdateSliceOfPointersCallback(system_, "PartWhoseNodeIsExpanded", &system_.PartWhoseNodeIsExpanded)

		case "DataFlows":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DataFlow](systemFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DataFlow, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DataFlow)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					systemFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.DataFlow](systemFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			system_.DataFlows = instanceSlice
			systemFormCallback.probe.UpdateSliceOfPointersCallback(system_, "DataFlows", &system_.DataFlows)

		case "IsDataFlowsNodeExpanded":
			FormDivBasicFieldToField(&(system_.IsDataFlowsNodeExpanded), formDiv)
		case "ExternalParts":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Part](systemFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Part, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Part)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					systemFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Part](systemFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			system_.ExternalParts = instanceSlice
			systemFormCallback.probe.UpdateSliceOfPointersCallback(system_, "ExternalParts", &system_.ExternalParts)

		case "ExternalPartWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Part](systemFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Part, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Part)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					systemFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Part](systemFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			system_.ExternalPartWhoseNodeIsExpanded = instanceSlice
			systemFormCallback.probe.UpdateSliceOfPointersCallback(system_, "ExternalPartWhoseNodeIsExpanded", &system_.ExternalPartWhoseNodeIsExpanded)

		case "DiagramStructure:SystemsWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramStructure instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramStructure instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramStructure](systemFormCallback.probe.stageOfInterest)
			targetDiagramStructureIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramStructureIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramStructure instances and update their SystemsWhoseNodeIsExpanded slice
			for _diagramstructure := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramStructure](systemFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(systemFormCallback.probe.stageOfInterest, _diagramstructure)
				
				// if DiagramStructure is selected
				if targetDiagramStructureIDs[id] {
					// ensure system_ is in _diagramstructure.SystemsWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramstructure.SystemsWhoseNodeIsExpanded {
						if _b == system_ {
							found = true
							break
						}
					}
					if !found {
						_diagramstructure.SystemsWhoseNodeIsExpanded = append(_diagramstructure.SystemsWhoseNodeIsExpanded, system_)
						systemFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "SystemsWhoseNodeIsExpanded", &_diagramstructure.SystemsWhoseNodeIsExpanded)
					}
				} else {
					// ensure system_ is NOT in _diagramstructure.SystemsWhoseNodeIsExpanded
					idx := slices.Index(_diagramstructure.SystemsWhoseNodeIsExpanded, system_)
					if idx != -1 {
						_diagramstructure.SystemsWhoseNodeIsExpanded = slices.Delete(_diagramstructure.SystemsWhoseNodeIsExpanded, idx, idx+1)
						systemFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "SystemsWhoseNodeIsExpanded", &_diagramstructure.SystemsWhoseNodeIsExpanded)
					}
				}
			}
		case "DiagramStructure:AllocatedSystemesWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramStructure instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramStructure instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramStructure](systemFormCallback.probe.stageOfInterest)
			targetDiagramStructureIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramStructureIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramStructure instances and update their AllocatedSystemesWhoseNodeIsExpanded slice
			for _diagramstructure := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramStructure](systemFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(systemFormCallback.probe.stageOfInterest, _diagramstructure)
				
				// if DiagramStructure is selected
				if targetDiagramStructureIDs[id] {
					// ensure system_ is in _diagramstructure.AllocatedSystemesWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramstructure.AllocatedSystemesWhoseNodeIsExpanded {
						if _b == system_ {
							found = true
							break
						}
					}
					if !found {
						_diagramstructure.AllocatedSystemesWhoseNodeIsExpanded = append(_diagramstructure.AllocatedSystemesWhoseNodeIsExpanded, system_)
						systemFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "AllocatedSystemesWhoseNodeIsExpanded", &_diagramstructure.AllocatedSystemesWhoseNodeIsExpanded)
					}
				} else {
					// ensure system_ is NOT in _diagramstructure.AllocatedSystemesWhoseNodeIsExpanded
					idx := slices.Index(_diagramstructure.AllocatedSystemesWhoseNodeIsExpanded, system_)
					if idx != -1 {
						_diagramstructure.AllocatedSystemesWhoseNodeIsExpanded = slices.Delete(_diagramstructure.AllocatedSystemesWhoseNodeIsExpanded, idx, idx+1)
						systemFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "AllocatedSystemesWhoseNodeIsExpanded", &_diagramstructure.AllocatedSystemesWhoseNodeIsExpanded)
					}
				}
			}
		case "Library:RootSystemes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](systemFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their RootSystemes slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](systemFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(systemFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure system_ is in _library.RootSystemes
					found := false
					for _, _b := range _library.RootSystemes {
						if _b == system_ {
							found = true
							break
						}
					}
					if !found {
						_library.RootSystemes = append(_library.RootSystemes, system_)
						systemFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootSystemes", &_library.RootSystemes)
					}
				} else {
					// ensure system_ is NOT in _library.RootSystemes
					idx := slices.Index(_library.RootSystemes, system_)
					if idx != -1 {
						_library.RootSystemes = slices.Delete(_library.RootSystemes, idx, idx+1)
						systemFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootSystemes", &_library.RootSystemes)
					}
				}
			}
		case "Library:SystemsWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](systemFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their SystemsWhoseNodeIsExpanded slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](systemFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(systemFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure system_ is in _library.SystemsWhoseNodeIsExpanded
					found := false
					for _, _b := range _library.SystemsWhoseNodeIsExpanded {
						if _b == system_ {
							found = true
							break
						}
					}
					if !found {
						_library.SystemsWhoseNodeIsExpanded = append(_library.SystemsWhoseNodeIsExpanded, system_)
						systemFormCallback.probe.UpdateSliceOfPointersCallback(_library, "SystemsWhoseNodeIsExpanded", &_library.SystemsWhoseNodeIsExpanded)
					}
				} else {
					// ensure system_ is NOT in _library.SystemsWhoseNodeIsExpanded
					idx := slices.Index(_library.SystemsWhoseNodeIsExpanded, system_)
					if idx != -1 {
						_library.SystemsWhoseNodeIsExpanded = slices.Delete(_library.SystemsWhoseNodeIsExpanded, idx, idx+1)
						systemFormCallback.probe.UpdateSliceOfPointersCallback(_library, "SystemsWhoseNodeIsExpanded", &_library.SystemsWhoseNodeIsExpanded)
					}
				}
			}
		case "System:SubSystemes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the System instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target System instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.System](systemFormCallback.probe.stageOfInterest)
			targetSystemIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSystemIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all System instances and update their SubSystemes slice
			for _system := range *models.GetGongstructInstancesSetFromPointerType[*models.System](systemFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(systemFormCallback.probe.stageOfInterest, _system)
				
				// if System is selected
				if targetSystemIDs[id] {
					// ensure system_ is in _system.SubSystemes
					found := false
					for _, _b := range _system.SubSystemes {
						if _b == system_ {
							found = true
							break
						}
					}
					if !found {
						_system.SubSystemes = append(_system.SubSystemes, system_)
						systemFormCallback.probe.UpdateSliceOfPointersCallback(_system, "SubSystemes", &_system.SubSystemes)
					}
				} else {
					// ensure system_ is NOT in _system.SubSystemes
					idx := slices.Index(_system.SubSystemes, system_)
					if idx != -1 {
						_system.SubSystemes = slices.Delete(_system.SubSystemes, idx, idx+1)
						systemFormCallback.probe.UpdateSliceOfPointersCallback(_system, "SubSystemes", &_system.SubSystemes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if systemFormCallback.formGroup.HasSuppressButtonBeenPressed {
		system_.Unstage(systemFormCallback.probe.stageOfInterest)
	}

	systemFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.System](
		systemFormCallback.probe,
	)

	// display a new form by reset the form stage
	if systemFormCallback.CreationMode || systemFormCallback.formGroup.HasSuppressButtonBeenPressed {
		systemFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(systemFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SystemFormCallback(
			nil,
			systemFormCallback.probe,
			newFormGroup,
		)
		system := new(models.System)
		FillUpForm(system, newFormGroup, systemFormCallback.probe)
		systemFormCallback.probe.formStage.Commit()
	}

	systemFormCallback.probe.ux_tree()
}
func __gong__New__SystemShapeFormCallback(
	systemshape *models.SystemShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (systemshapeFormCallback *SystemShapeFormCallback) {
	systemshapeFormCallback = new(SystemShapeFormCallback)
	systemshapeFormCallback.probe = probe
	systemshapeFormCallback.systemshape = systemshape
	systemshapeFormCallback.formGroup = formGroup

	systemshapeFormCallback.CreationMode = (systemshape == nil)

	return
}

type SystemShapeFormCallback struct {
	systemshape *models.SystemShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (systemshapeFormCallback *SystemShapeFormCallback) OnSave() {
	systemshapeFormCallback.probe.stageOfInterest.Lock()
	defer systemshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("SystemShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	systemshapeFormCallback.probe.formStage.Checkout()

	if systemshapeFormCallback.systemshape == nil {
		systemshapeFormCallback.systemshape = new(models.SystemShape).Stage(systemshapeFormCallback.probe.stageOfInterest)
	}
	systemshape_ := systemshapeFormCallback.systemshape
	_ = systemshape_

	for _, formDiv := range systemshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(systemshape_.Name), formDiv)
		case "System":
			FormDivSelectFieldToField(&(systemshape_.System), systemshapeFormCallback.probe.stageOfInterest, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(systemshape_.IsExpanded), formDiv)
		case "X":
			FormDivBasicFieldToField(&(systemshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(systemshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(systemshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(systemshape_.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(systemshape_.IsHidden), formDiv)
		case "DiagramStructure:System_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramStructure instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramStructure instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramStructure](systemshapeFormCallback.probe.stageOfInterest)
			targetDiagramStructureIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramStructureIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramStructure instances and update their System_Shapes slice
			for _diagramstructure := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramStructure](systemshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(systemshapeFormCallback.probe.stageOfInterest, _diagramstructure)
				
				// if DiagramStructure is selected
				if targetDiagramStructureIDs[id] {
					// ensure systemshape_ is in _diagramstructure.System_Shapes
					found := false
					for _, _b := range _diagramstructure.System_Shapes {
						if _b == systemshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramstructure.System_Shapes = append(_diagramstructure.System_Shapes, systemshape_)
						systemshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "System_Shapes", &_diagramstructure.System_Shapes)
					}
				} else {
					// ensure systemshape_ is NOT in _diagramstructure.System_Shapes
					idx := slices.Index(_diagramstructure.System_Shapes, systemshape_)
					if idx != -1 {
						_diagramstructure.System_Shapes = slices.Delete(_diagramstructure.System_Shapes, idx, idx+1)
						systemshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "System_Shapes", &_diagramstructure.System_Shapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if systemshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		systemshape_.Unstage(systemshapeFormCallback.probe.stageOfInterest)
	}

	systemshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.SystemShape](
		systemshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if systemshapeFormCallback.CreationMode || systemshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		systemshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(systemshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SystemShapeFormCallback(
			nil,
			systemshapeFormCallback.probe,
			newFormGroup,
		)
		systemshape := new(models.SystemShape)
		FillUpForm(systemshape, newFormGroup, systemshapeFormCallback.probe)
		systemshapeFormCallback.probe.formStage.Commit()
	}

	systemshapeFormCallback.probe.ux_tree()
}
