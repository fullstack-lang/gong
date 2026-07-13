// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__ArcNormalVectorShapeFormCallback(
	arcnormalvectorshape *models.ArcNormalVectorShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (arcnormalvectorshapeFormCallback *ArcNormalVectorShapeFormCallback) {
	arcnormalvectorshapeFormCallback = new(ArcNormalVectorShapeFormCallback)
	arcnormalvectorshapeFormCallback.probe = probe
	arcnormalvectorshapeFormCallback.arcnormalvectorshape = arcnormalvectorshape
	arcnormalvectorshapeFormCallback.formGroup = formGroup

	arcnormalvectorshapeFormCallback.CreationMode = (arcnormalvectorshape == nil)

	return
}

type ArcNormalVectorShapeFormCallback struct {
	arcnormalvectorshape *models.ArcNormalVectorShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (arcnormalvectorshapeFormCallback *ArcNormalVectorShapeFormCallback) OnSave() {
	arcnormalvectorshapeFormCallback.probe.stageOfInterest.Lock()
	defer arcnormalvectorshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ArcNormalVectorShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	arcnormalvectorshapeFormCallback.probe.formStage.Checkout()

	if arcnormalvectorshapeFormCallback.arcnormalvectorshape == nil {
		arcnormalvectorshapeFormCallback.arcnormalvectorshape = new(models.ArcNormalVectorShape).Stage(arcnormalvectorshapeFormCallback.probe.stageOfInterest)
	}
	arcnormalvectorshape_ := arcnormalvectorshapeFormCallback.arcnormalvectorshape
	_ = arcnormalvectorshape_

	for _, formDiv := range arcnormalvectorshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(arcnormalvectorshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(arcnormalvectorshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(arcnormalvectorshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(arcnormalvectorshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(arcnormalvectorshape_.EndY), formDiv)
		case "ArcNormalVectorShapeGrid:ArcNormalVectorShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the ArcNormalVectorShapeGrid instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target ArcNormalVectorShapeGrid instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.ArcNormalVectorShapeGrid](arcnormalvectorshapeFormCallback.probe.stageOfInterest)
			targetArcNormalVectorShapeGridIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetArcNormalVectorShapeGridIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all ArcNormalVectorShapeGrid instances and update their ArcNormalVectorShapes slice
			for _arcnormalvectorshapegrid := range *models.GetGongstructInstancesSetFromPointerType[*models.ArcNormalVectorShapeGrid](arcnormalvectorshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(arcnormalvectorshapeFormCallback.probe.stageOfInterest, _arcnormalvectorshapegrid)
				
				// if ArcNormalVectorShapeGrid is selected
				if targetArcNormalVectorShapeGridIDs[id] {
					// ensure arcnormalvectorshape_ is in _arcnormalvectorshapegrid.ArcNormalVectorShapes
					found := false
					for _, _b := range _arcnormalvectorshapegrid.ArcNormalVectorShapes {
						if _b == arcnormalvectorshape_ {
							found = true
							break
						}
					}
					if !found {
						_arcnormalvectorshapegrid.ArcNormalVectorShapes = append(_arcnormalvectorshapegrid.ArcNormalVectorShapes, arcnormalvectorshape_)
						arcnormalvectorshapeFormCallback.probe.UpdateSliceOfPointersCallback(_arcnormalvectorshapegrid, "ArcNormalVectorShapes", &_arcnormalvectorshapegrid.ArcNormalVectorShapes)
					}
				} else {
					// ensure arcnormalvectorshape_ is NOT in _arcnormalvectorshapegrid.ArcNormalVectorShapes
					idx := slices.Index(_arcnormalvectorshapegrid.ArcNormalVectorShapes, arcnormalvectorshape_)
					if idx != -1 {
						_arcnormalvectorshapegrid.ArcNormalVectorShapes = slices.Delete(_arcnormalvectorshapegrid.ArcNormalVectorShapes, idx, idx+1)
						arcnormalvectorshapeFormCallback.probe.UpdateSliceOfPointersCallback(_arcnormalvectorshapegrid, "ArcNormalVectorShapes", &_arcnormalvectorshapegrid.ArcNormalVectorShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if arcnormalvectorshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		arcnormalvectorshape_.Unstage(arcnormalvectorshapeFormCallback.probe.stageOfInterest)
	}

	arcnormalvectorshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ArcNormalVectorShape](
		arcnormalvectorshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if arcnormalvectorshapeFormCallback.CreationMode || arcnormalvectorshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		arcnormalvectorshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(arcnormalvectorshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ArcNormalVectorShapeFormCallback(
			nil,
			arcnormalvectorshapeFormCallback.probe,
			newFormGroup,
		)
		arcnormalvectorshape := new(models.ArcNormalVectorShape)
		FillUpForm(arcnormalvectorshape, newFormGroup, arcnormalvectorshapeFormCallback.probe)
		arcnormalvectorshapeFormCallback.probe.formStage.Commit()
	}

	arcnormalvectorshapeFormCallback.probe.ux_tree()
}
func __gong__New__ArcNormalVectorShapeGridFormCallback(
	arcnormalvectorshapegrid *models.ArcNormalVectorShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (arcnormalvectorshapegridFormCallback *ArcNormalVectorShapeGridFormCallback) {
	arcnormalvectorshapegridFormCallback = new(ArcNormalVectorShapeGridFormCallback)
	arcnormalvectorshapegridFormCallback.probe = probe
	arcnormalvectorshapegridFormCallback.arcnormalvectorshapegrid = arcnormalvectorshapegrid
	arcnormalvectorshapegridFormCallback.formGroup = formGroup

	arcnormalvectorshapegridFormCallback.CreationMode = (arcnormalvectorshapegrid == nil)

	return
}

type ArcNormalVectorShapeGridFormCallback struct {
	arcnormalvectorshapegrid *models.ArcNormalVectorShapeGrid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (arcnormalvectorshapegridFormCallback *ArcNormalVectorShapeGridFormCallback) OnSave() {
	arcnormalvectorshapegridFormCallback.probe.stageOfInterest.Lock()
	defer arcnormalvectorshapegridFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ArcNormalVectorShapeGridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	arcnormalvectorshapegridFormCallback.probe.formStage.Checkout()

	if arcnormalvectorshapegridFormCallback.arcnormalvectorshapegrid == nil {
		arcnormalvectorshapegridFormCallback.arcnormalvectorshapegrid = new(models.ArcNormalVectorShapeGrid).Stage(arcnormalvectorshapegridFormCallback.probe.stageOfInterest)
	}
	arcnormalvectorshapegrid_ := arcnormalvectorshapegridFormCallback.arcnormalvectorshapegrid
	_ = arcnormalvectorshapegrid_

	for _, formDiv := range arcnormalvectorshapegridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(arcnormalvectorshapegrid_.Name), formDiv)
		case "ArcNormalVectorShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ArcNormalVectorShape](arcnormalvectorshapegridFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ArcNormalVectorShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ArcNormalVectorShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					arcnormalvectorshapegridFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ArcNormalVectorShape](arcnormalvectorshapegridFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			arcnormalvectorshapegrid_.ArcNormalVectorShapes = instanceSlice
			arcnormalvectorshapegridFormCallback.probe.UpdateSliceOfPointersCallback(arcnormalvectorshapegrid_, "ArcNormalVectorShapes", &arcnormalvectorshapegrid_.ArcNormalVectorShapes)

		}
	}

	// manage the suppress operation
	if arcnormalvectorshapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		arcnormalvectorshapegrid_.Unstage(arcnormalvectorshapegridFormCallback.probe.stageOfInterest)
	}

	arcnormalvectorshapegridFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ArcNormalVectorShapeGrid](
		arcnormalvectorshapegridFormCallback.probe,
	)

	// display a new form by reset the form stage
	if arcnormalvectorshapegridFormCallback.CreationMode || arcnormalvectorshapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		arcnormalvectorshapegridFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(arcnormalvectorshapegridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ArcNormalVectorShapeGridFormCallback(
			nil,
			arcnormalvectorshapegridFormCallback.probe,
			newFormGroup,
		)
		arcnormalvectorshapegrid := new(models.ArcNormalVectorShapeGrid)
		FillUpForm(arcnormalvectorshapegrid, newFormGroup, arcnormalvectorshapegridFormCallback.probe)
		arcnormalvectorshapegridFormCallback.probe.formStage.Commit()
	}

	arcnormalvectorshapegridFormCallback.probe.ux_tree()
}
func __gong__New__AxesShapeFormCallback(
	axesshape *models.AxesShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (axesshapeFormCallback *AxesShapeFormCallback) {
	axesshapeFormCallback = new(AxesShapeFormCallback)
	axesshapeFormCallback.probe = probe
	axesshapeFormCallback.axesshape = axesshape
	axesshapeFormCallback.formGroup = formGroup

	axesshapeFormCallback.CreationMode = (axesshape == nil)

	return
}

type AxesShapeFormCallback struct {
	axesshape *models.AxesShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (axesshapeFormCallback *AxesShapeFormCallback) OnSave() {
	axesshapeFormCallback.probe.stageOfInterest.Lock()
	defer axesshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("AxesShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	axesshapeFormCallback.probe.formStage.Checkout()

	if axesshapeFormCallback.axesshape == nil {
		axesshapeFormCallback.axesshape = new(models.AxesShape).Stage(axesshapeFormCallback.probe.stageOfInterest)
	}
	axesshape_ := axesshapeFormCallback.axesshape
	_ = axesshape_

	for _, formDiv := range axesshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(axesshape_.Name), formDiv)
		case "LengthX":
			FormDivBasicFieldToField(&(axesshape_.LengthX), formDiv)
		case "LengthY":
			FormDivBasicFieldToField(&(axesshape_.LengthY), formDiv)
		case "IsWithHiddenHandle":
			FormDivBasicFieldToField(&(axesshape_.IsWithHiddenHandle), formDiv)
		}
	}

	// manage the suppress operation
	if axesshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		axesshape_.Unstage(axesshapeFormCallback.probe.stageOfInterest)
	}

	axesshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.AxesShape](
		axesshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if axesshapeFormCallback.CreationMode || axesshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		axesshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(axesshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AxesShapeFormCallback(
			nil,
			axesshapeFormCallback.probe,
			newFormGroup,
		)
		axesshape := new(models.AxesShape)
		FillUpForm(axesshape, newFormGroup, axesshapeFormCallback.probe)
		axesshapeFormCallback.probe.formStage.Commit()
	}

	axesshapeFormCallback.probe.ux_tree()
}
func __gong__New__BaseVectorShapeFormCallback(
	basevectorshape *models.BaseVectorShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (basevectorshapeFormCallback *BaseVectorShapeFormCallback) {
	basevectorshapeFormCallback = new(BaseVectorShapeFormCallback)
	basevectorshapeFormCallback.probe = probe
	basevectorshapeFormCallback.basevectorshape = basevectorshape
	basevectorshapeFormCallback.formGroup = formGroup

	basevectorshapeFormCallback.CreationMode = (basevectorshape == nil)

	return
}

type BaseVectorShapeFormCallback struct {
	basevectorshape *models.BaseVectorShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (basevectorshapeFormCallback *BaseVectorShapeFormCallback) OnSave() {
	basevectorshapeFormCallback.probe.stageOfInterest.Lock()
	defer basevectorshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("BaseVectorShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	basevectorshapeFormCallback.probe.formStage.Checkout()

	if basevectorshapeFormCallback.basevectorshape == nil {
		basevectorshapeFormCallback.basevectorshape = new(models.BaseVectorShape).Stage(basevectorshapeFormCallback.probe.stageOfInterest)
	}
	basevectorshape_ := basevectorshapeFormCallback.basevectorshape
	_ = basevectorshape_

	for _, formDiv := range basevectorshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(basevectorshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(basevectorshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(basevectorshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(basevectorshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(basevectorshape_.EndY), formDiv)
		case "BaseVectorShapeGrid:BaseVectorShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the BaseVectorShapeGrid instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target BaseVectorShapeGrid instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.BaseVectorShapeGrid](basevectorshapeFormCallback.probe.stageOfInterest)
			targetBaseVectorShapeGridIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetBaseVectorShapeGridIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all BaseVectorShapeGrid instances and update their BaseVectorShapes slice
			for _basevectorshapegrid := range *models.GetGongstructInstancesSetFromPointerType[*models.BaseVectorShapeGrid](basevectorshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(basevectorshapeFormCallback.probe.stageOfInterest, _basevectorshapegrid)
				
				// if BaseVectorShapeGrid is selected
				if targetBaseVectorShapeGridIDs[id] {
					// ensure basevectorshape_ is in _basevectorshapegrid.BaseVectorShapes
					found := false
					for _, _b := range _basevectorshapegrid.BaseVectorShapes {
						if _b == basevectorshape_ {
							found = true
							break
						}
					}
					if !found {
						_basevectorshapegrid.BaseVectorShapes = append(_basevectorshapegrid.BaseVectorShapes, basevectorshape_)
						basevectorshapeFormCallback.probe.UpdateSliceOfPointersCallback(_basevectorshapegrid, "BaseVectorShapes", &_basevectorshapegrid.BaseVectorShapes)
					}
				} else {
					// ensure basevectorshape_ is NOT in _basevectorshapegrid.BaseVectorShapes
					idx := slices.Index(_basevectorshapegrid.BaseVectorShapes, basevectorshape_)
					if idx != -1 {
						_basevectorshapegrid.BaseVectorShapes = slices.Delete(_basevectorshapegrid.BaseVectorShapes, idx, idx+1)
						basevectorshapeFormCallback.probe.UpdateSliceOfPointersCallback(_basevectorshapegrid, "BaseVectorShapes", &_basevectorshapegrid.BaseVectorShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if basevectorshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		basevectorshape_.Unstage(basevectorshapeFormCallback.probe.stageOfInterest)
	}

	basevectorshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.BaseVectorShape](
		basevectorshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if basevectorshapeFormCallback.CreationMode || basevectorshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		basevectorshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(basevectorshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BaseVectorShapeFormCallback(
			nil,
			basevectorshapeFormCallback.probe,
			newFormGroup,
		)
		basevectorshape := new(models.BaseVectorShape)
		FillUpForm(basevectorshape, newFormGroup, basevectorshapeFormCallback.probe)
		basevectorshapeFormCallback.probe.formStage.Commit()
	}

	basevectorshapeFormCallback.probe.ux_tree()
}
func __gong__New__BaseVectorShapeGridFormCallback(
	basevectorshapegrid *models.BaseVectorShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (basevectorshapegridFormCallback *BaseVectorShapeGridFormCallback) {
	basevectorshapegridFormCallback = new(BaseVectorShapeGridFormCallback)
	basevectorshapegridFormCallback.probe = probe
	basevectorshapegridFormCallback.basevectorshapegrid = basevectorshapegrid
	basevectorshapegridFormCallback.formGroup = formGroup

	basevectorshapegridFormCallback.CreationMode = (basevectorshapegrid == nil)

	return
}

type BaseVectorShapeGridFormCallback struct {
	basevectorshapegrid *models.BaseVectorShapeGrid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (basevectorshapegridFormCallback *BaseVectorShapeGridFormCallback) OnSave() {
	basevectorshapegridFormCallback.probe.stageOfInterest.Lock()
	defer basevectorshapegridFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("BaseVectorShapeGridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	basevectorshapegridFormCallback.probe.formStage.Checkout()

	if basevectorshapegridFormCallback.basevectorshapegrid == nil {
		basevectorshapegridFormCallback.basevectorshapegrid = new(models.BaseVectorShapeGrid).Stage(basevectorshapegridFormCallback.probe.stageOfInterest)
	}
	basevectorshapegrid_ := basevectorshapegridFormCallback.basevectorshapegrid
	_ = basevectorshapegrid_

	for _, formDiv := range basevectorshapegridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(basevectorshapegrid_.Name), formDiv)
		case "BaseVectorShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.BaseVectorShape](basevectorshapegridFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.BaseVectorShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.BaseVectorShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					basevectorshapegridFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.BaseVectorShape](basevectorshapegridFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			basevectorshapegrid_.BaseVectorShapes = instanceSlice
			basevectorshapegridFormCallback.probe.UpdateSliceOfPointersCallback(basevectorshapegrid_, "BaseVectorShapes", &basevectorshapegrid_.BaseVectorShapes)

		}
	}

	// manage the suppress operation
	if basevectorshapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		basevectorshapegrid_.Unstage(basevectorshapegridFormCallback.probe.stageOfInterest)
	}

	basevectorshapegridFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.BaseVectorShapeGrid](
		basevectorshapegridFormCallback.probe,
	)

	// display a new form by reset the form stage
	if basevectorshapegridFormCallback.CreationMode || basevectorshapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		basevectorshapegridFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(basevectorshapegridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BaseVectorShapeGridFormCallback(
			nil,
			basevectorshapegridFormCallback.probe,
			newFormGroup,
		)
		basevectorshapegrid := new(models.BaseVectorShapeGrid)
		FillUpForm(basevectorshapegrid, newFormGroup, basevectorshapegridFormCallback.probe)
		basevectorshapegridFormCallback.probe.formStage.Commit()
	}

	basevectorshapegridFormCallback.probe.ux_tree()
}
func __gong__New__BottomEndArcShapeFormCallback(
	bottomendarcshape *models.BottomEndArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (bottomendarcshapeFormCallback *BottomEndArcShapeFormCallback) {
	bottomendarcshapeFormCallback = new(BottomEndArcShapeFormCallback)
	bottomendarcshapeFormCallback.probe = probe
	bottomendarcshapeFormCallback.bottomendarcshape = bottomendarcshape
	bottomendarcshapeFormCallback.formGroup = formGroup

	bottomendarcshapeFormCallback.CreationMode = (bottomendarcshape == nil)

	return
}

type BottomEndArcShapeFormCallback struct {
	bottomendarcshape *models.BottomEndArcShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (bottomendarcshapeFormCallback *BottomEndArcShapeFormCallback) OnSave() {
	bottomendarcshapeFormCallback.probe.stageOfInterest.Lock()
	defer bottomendarcshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("BottomEndArcShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bottomendarcshapeFormCallback.probe.formStage.Checkout()

	if bottomendarcshapeFormCallback.bottomendarcshape == nil {
		bottomendarcshapeFormCallback.bottomendarcshape = new(models.BottomEndArcShape).Stage(bottomendarcshapeFormCallback.probe.stageOfInterest)
	}
	bottomendarcshape_ := bottomendarcshapeFormCallback.bottomendarcshape
	_ = bottomendarcshape_

	for _, formDiv := range bottomendarcshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(bottomendarcshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(bottomendarcshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(bottomendarcshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(bottomendarcshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(bottomendarcshape_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(bottomendarcshape_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(bottomendarcshape_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(bottomendarcshape_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(bottomendarcshape_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(bottomendarcshape_.RadiusY), formDiv)
		case "BottomEndArcShapeGrid:BottomEndArcShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the BottomEndArcShapeGrid instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target BottomEndArcShapeGrid instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.BottomEndArcShapeGrid](bottomendarcshapeFormCallback.probe.stageOfInterest)
			targetBottomEndArcShapeGridIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetBottomEndArcShapeGridIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all BottomEndArcShapeGrid instances and update their BottomEndArcShapes slice
			for _bottomendarcshapegrid := range *models.GetGongstructInstancesSetFromPointerType[*models.BottomEndArcShapeGrid](bottomendarcshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(bottomendarcshapeFormCallback.probe.stageOfInterest, _bottomendarcshapegrid)
				
				// if BottomEndArcShapeGrid is selected
				if targetBottomEndArcShapeGridIDs[id] {
					// ensure bottomendarcshape_ is in _bottomendarcshapegrid.BottomEndArcShapes
					found := false
					for _, _b := range _bottomendarcshapegrid.BottomEndArcShapes {
						if _b == bottomendarcshape_ {
							found = true
							break
						}
					}
					if !found {
						_bottomendarcshapegrid.BottomEndArcShapes = append(_bottomendarcshapegrid.BottomEndArcShapes, bottomendarcshape_)
						bottomendarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_bottomendarcshapegrid, "BottomEndArcShapes", &_bottomendarcshapegrid.BottomEndArcShapes)
					}
				} else {
					// ensure bottomendarcshape_ is NOT in _bottomendarcshapegrid.BottomEndArcShapes
					idx := slices.Index(_bottomendarcshapegrid.BottomEndArcShapes, bottomendarcshape_)
					if idx != -1 {
						_bottomendarcshapegrid.BottomEndArcShapes = slices.Delete(_bottomendarcshapegrid.BottomEndArcShapes, idx, idx+1)
						bottomendarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_bottomendarcshapegrid, "BottomEndArcShapes", &_bottomendarcshapegrid.BottomEndArcShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if bottomendarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bottomendarcshape_.Unstage(bottomendarcshapeFormCallback.probe.stageOfInterest)
	}

	bottomendarcshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.BottomEndArcShape](
		bottomendarcshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if bottomendarcshapeFormCallback.CreationMode || bottomendarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bottomendarcshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(bottomendarcshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BottomEndArcShapeFormCallback(
			nil,
			bottomendarcshapeFormCallback.probe,
			newFormGroup,
		)
		bottomendarcshape := new(models.BottomEndArcShape)
		FillUpForm(bottomendarcshape, newFormGroup, bottomendarcshapeFormCallback.probe)
		bottomendarcshapeFormCallback.probe.formStage.Commit()
	}

	bottomendarcshapeFormCallback.probe.ux_tree()
}
func __gong__New__BottomEndArcShapeGridFormCallback(
	bottomendarcshapegrid *models.BottomEndArcShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (bottomendarcshapegridFormCallback *BottomEndArcShapeGridFormCallback) {
	bottomendarcshapegridFormCallback = new(BottomEndArcShapeGridFormCallback)
	bottomendarcshapegridFormCallback.probe = probe
	bottomendarcshapegridFormCallback.bottomendarcshapegrid = bottomendarcshapegrid
	bottomendarcshapegridFormCallback.formGroup = formGroup

	bottomendarcshapegridFormCallback.CreationMode = (bottomendarcshapegrid == nil)

	return
}

type BottomEndArcShapeGridFormCallback struct {
	bottomendarcshapegrid *models.BottomEndArcShapeGrid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (bottomendarcshapegridFormCallback *BottomEndArcShapeGridFormCallback) OnSave() {
	bottomendarcshapegridFormCallback.probe.stageOfInterest.Lock()
	defer bottomendarcshapegridFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("BottomEndArcShapeGridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bottomendarcshapegridFormCallback.probe.formStage.Checkout()

	if bottomendarcshapegridFormCallback.bottomendarcshapegrid == nil {
		bottomendarcshapegridFormCallback.bottomendarcshapegrid = new(models.BottomEndArcShapeGrid).Stage(bottomendarcshapegridFormCallback.probe.stageOfInterest)
	}
	bottomendarcshapegrid_ := bottomendarcshapegridFormCallback.bottomendarcshapegrid
	_ = bottomendarcshapegrid_

	for _, formDiv := range bottomendarcshapegridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(bottomendarcshapegrid_.Name), formDiv)
		case "BottomEndArcShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.BottomEndArcShape](bottomendarcshapegridFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.BottomEndArcShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.BottomEndArcShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					bottomendarcshapegridFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.BottomEndArcShape](bottomendarcshapegridFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			bottomendarcshapegrid_.BottomEndArcShapes = instanceSlice
			bottomendarcshapegridFormCallback.probe.UpdateSliceOfPointersCallback(bottomendarcshapegrid_, "BottomEndArcShapes", &bottomendarcshapegrid_.BottomEndArcShapes)

		}
	}

	// manage the suppress operation
	if bottomendarcshapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bottomendarcshapegrid_.Unstage(bottomendarcshapegridFormCallback.probe.stageOfInterest)
	}

	bottomendarcshapegridFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.BottomEndArcShapeGrid](
		bottomendarcshapegridFormCallback.probe,
	)

	// display a new form by reset the form stage
	if bottomendarcshapegridFormCallback.CreationMode || bottomendarcshapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bottomendarcshapegridFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(bottomendarcshapegridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BottomEndArcShapeGridFormCallback(
			nil,
			bottomendarcshapegridFormCallback.probe,
			newFormGroup,
		)
		bottomendarcshapegrid := new(models.BottomEndArcShapeGrid)
		FillUpForm(bottomendarcshapegrid, newFormGroup, bottomendarcshapegridFormCallback.probe)
		bottomendarcshapegridFormCallback.probe.formStage.Commit()
	}

	bottomendarcshapegridFormCallback.probe.ux_tree()
}
func __gong__New__BottomStackGrowthCurveEndArcShapeFormCallback(
	bottomstackgrowthcurveendarcshape *models.BottomStackGrowthCurveEndArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (bottomstackgrowthcurveendarcshapeFormCallback *BottomStackGrowthCurveEndArcShapeFormCallback) {
	bottomstackgrowthcurveendarcshapeFormCallback = new(BottomStackGrowthCurveEndArcShapeFormCallback)
	bottomstackgrowthcurveendarcshapeFormCallback.probe = probe
	bottomstackgrowthcurveendarcshapeFormCallback.bottomstackgrowthcurveendarcshape = bottomstackgrowthcurveendarcshape
	bottomstackgrowthcurveendarcshapeFormCallback.formGroup = formGroup

	bottomstackgrowthcurveendarcshapeFormCallback.CreationMode = (bottomstackgrowthcurveendarcshape == nil)

	return
}

type BottomStackGrowthCurveEndArcShapeFormCallback struct {
	bottomstackgrowthcurveendarcshape *models.BottomStackGrowthCurveEndArcShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (bottomstackgrowthcurveendarcshapeFormCallback *BottomStackGrowthCurveEndArcShapeFormCallback) OnSave() {
	bottomstackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest.Lock()
	defer bottomstackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("BottomStackGrowthCurveEndArcShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bottomstackgrowthcurveendarcshapeFormCallback.probe.formStage.Checkout()

	if bottomstackgrowthcurveendarcshapeFormCallback.bottomstackgrowthcurveendarcshape == nil {
		bottomstackgrowthcurveendarcshapeFormCallback.bottomstackgrowthcurveendarcshape = new(models.BottomStackGrowthCurveEndArcShape).Stage(bottomstackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest)
	}
	bottomstackgrowthcurveendarcshape_ := bottomstackgrowthcurveendarcshapeFormCallback.bottomstackgrowthcurveendarcshape
	_ = bottomstackgrowthcurveendarcshape_

	for _, formDiv := range bottomstackgrowthcurveendarcshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(bottomstackgrowthcurveendarcshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(bottomstackgrowthcurveendarcshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(bottomstackgrowthcurveendarcshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(bottomstackgrowthcurveendarcshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(bottomstackgrowthcurveendarcshape_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(bottomstackgrowthcurveendarcshape_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(bottomstackgrowthcurveendarcshape_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(bottomstackgrowthcurveendarcshape_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(bottomstackgrowthcurveendarcshape_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(bottomstackgrowthcurveendarcshape_.RadiusY), formDiv)
		case "BottomStackOfGrowthCurve:BottomStackGrowthCurveEndArcShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the BottomStackOfGrowthCurve instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target BottomStackOfGrowthCurve instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.BottomStackOfGrowthCurve](bottomstackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest)
			targetBottomStackOfGrowthCurveIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetBottomStackOfGrowthCurveIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all BottomStackOfGrowthCurve instances and update their BottomStackGrowthCurveEndArcShapes slice
			for _bottomstackofgrowthcurve := range *models.GetGongstructInstancesSetFromPointerType[*models.BottomStackOfGrowthCurve](bottomstackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(bottomstackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest, _bottomstackofgrowthcurve)
				
				// if BottomStackOfGrowthCurve is selected
				if targetBottomStackOfGrowthCurveIDs[id] {
					// ensure bottomstackgrowthcurveendarcshape_ is in _bottomstackofgrowthcurve.BottomStackGrowthCurveEndArcShapes
					found := false
					for _, _b := range _bottomstackofgrowthcurve.BottomStackGrowthCurveEndArcShapes {
						if _b == bottomstackgrowthcurveendarcshape_ {
							found = true
							break
						}
					}
					if !found {
						_bottomstackofgrowthcurve.BottomStackGrowthCurveEndArcShapes = append(_bottomstackofgrowthcurve.BottomStackGrowthCurveEndArcShapes, bottomstackgrowthcurveendarcshape_)
						bottomstackgrowthcurveendarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_bottomstackofgrowthcurve, "BottomStackGrowthCurveEndArcShapes", &_bottomstackofgrowthcurve.BottomStackGrowthCurveEndArcShapes)
					}
				} else {
					// ensure bottomstackgrowthcurveendarcshape_ is NOT in _bottomstackofgrowthcurve.BottomStackGrowthCurveEndArcShapes
					idx := slices.Index(_bottomstackofgrowthcurve.BottomStackGrowthCurveEndArcShapes, bottomstackgrowthcurveendarcshape_)
					if idx != -1 {
						_bottomstackofgrowthcurve.BottomStackGrowthCurveEndArcShapes = slices.Delete(_bottomstackofgrowthcurve.BottomStackGrowthCurveEndArcShapes, idx, idx+1)
						bottomstackgrowthcurveendarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_bottomstackofgrowthcurve, "BottomStackGrowthCurveEndArcShapes", &_bottomstackofgrowthcurve.BottomStackGrowthCurveEndArcShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if bottomstackgrowthcurveendarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bottomstackgrowthcurveendarcshape_.Unstage(bottomstackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest)
	}

	bottomstackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.BottomStackGrowthCurveEndArcShape](
		bottomstackgrowthcurveendarcshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if bottomstackgrowthcurveendarcshapeFormCallback.CreationMode || bottomstackgrowthcurveendarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bottomstackgrowthcurveendarcshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(bottomstackgrowthcurveendarcshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BottomStackGrowthCurveEndArcShapeFormCallback(
			nil,
			bottomstackgrowthcurveendarcshapeFormCallback.probe,
			newFormGroup,
		)
		bottomstackgrowthcurveendarcshape := new(models.BottomStackGrowthCurveEndArcShape)
		FillUpForm(bottomstackgrowthcurveendarcshape, newFormGroup, bottomstackgrowthcurveendarcshapeFormCallback.probe)
		bottomstackgrowthcurveendarcshapeFormCallback.probe.formStage.Commit()
	}

	bottomstackgrowthcurveendarcshapeFormCallback.probe.ux_tree()
}
func __gong__New__BottomStackGrowthCurveStartArcShapeFormCallback(
	bottomstackgrowthcurvestartarcshape *models.BottomStackGrowthCurveStartArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (bottomstackgrowthcurvestartarcshapeFormCallback *BottomStackGrowthCurveStartArcShapeFormCallback) {
	bottomstackgrowthcurvestartarcshapeFormCallback = new(BottomStackGrowthCurveStartArcShapeFormCallback)
	bottomstackgrowthcurvestartarcshapeFormCallback.probe = probe
	bottomstackgrowthcurvestartarcshapeFormCallback.bottomstackgrowthcurvestartarcshape = bottomstackgrowthcurvestartarcshape
	bottomstackgrowthcurvestartarcshapeFormCallback.formGroup = formGroup

	bottomstackgrowthcurvestartarcshapeFormCallback.CreationMode = (bottomstackgrowthcurvestartarcshape == nil)

	return
}

type BottomStackGrowthCurveStartArcShapeFormCallback struct {
	bottomstackgrowthcurvestartarcshape *models.BottomStackGrowthCurveStartArcShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (bottomstackgrowthcurvestartarcshapeFormCallback *BottomStackGrowthCurveStartArcShapeFormCallback) OnSave() {
	bottomstackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest.Lock()
	defer bottomstackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("BottomStackGrowthCurveStartArcShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bottomstackgrowthcurvestartarcshapeFormCallback.probe.formStage.Checkout()

	if bottomstackgrowthcurvestartarcshapeFormCallback.bottomstackgrowthcurvestartarcshape == nil {
		bottomstackgrowthcurvestartarcshapeFormCallback.bottomstackgrowthcurvestartarcshape = new(models.BottomStackGrowthCurveStartArcShape).Stage(bottomstackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest)
	}
	bottomstackgrowthcurvestartarcshape_ := bottomstackgrowthcurvestartarcshapeFormCallback.bottomstackgrowthcurvestartarcshape
	_ = bottomstackgrowthcurvestartarcshape_

	for _, formDiv := range bottomstackgrowthcurvestartarcshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(bottomstackgrowthcurvestartarcshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(bottomstackgrowthcurvestartarcshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(bottomstackgrowthcurvestartarcshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(bottomstackgrowthcurvestartarcshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(bottomstackgrowthcurvestartarcshape_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(bottomstackgrowthcurvestartarcshape_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(bottomstackgrowthcurvestartarcshape_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(bottomstackgrowthcurvestartarcshape_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(bottomstackgrowthcurvestartarcshape_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(bottomstackgrowthcurvestartarcshape_.RadiusY), formDiv)
		case "BottomStackOfGrowthCurve:BottomStackGrowthCurveStartArcShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the BottomStackOfGrowthCurve instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target BottomStackOfGrowthCurve instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.BottomStackOfGrowthCurve](bottomstackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest)
			targetBottomStackOfGrowthCurveIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetBottomStackOfGrowthCurveIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all BottomStackOfGrowthCurve instances and update their BottomStackGrowthCurveStartArcShapes slice
			for _bottomstackofgrowthcurve := range *models.GetGongstructInstancesSetFromPointerType[*models.BottomStackOfGrowthCurve](bottomstackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(bottomstackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest, _bottomstackofgrowthcurve)
				
				// if BottomStackOfGrowthCurve is selected
				if targetBottomStackOfGrowthCurveIDs[id] {
					// ensure bottomstackgrowthcurvestartarcshape_ is in _bottomstackofgrowthcurve.BottomStackGrowthCurveStartArcShapes
					found := false
					for _, _b := range _bottomstackofgrowthcurve.BottomStackGrowthCurveStartArcShapes {
						if _b == bottomstackgrowthcurvestartarcshape_ {
							found = true
							break
						}
					}
					if !found {
						_bottomstackofgrowthcurve.BottomStackGrowthCurveStartArcShapes = append(_bottomstackofgrowthcurve.BottomStackGrowthCurveStartArcShapes, bottomstackgrowthcurvestartarcshape_)
						bottomstackgrowthcurvestartarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_bottomstackofgrowthcurve, "BottomStackGrowthCurveStartArcShapes", &_bottomstackofgrowthcurve.BottomStackGrowthCurveStartArcShapes)
					}
				} else {
					// ensure bottomstackgrowthcurvestartarcshape_ is NOT in _bottomstackofgrowthcurve.BottomStackGrowthCurveStartArcShapes
					idx := slices.Index(_bottomstackofgrowthcurve.BottomStackGrowthCurveStartArcShapes, bottomstackgrowthcurvestartarcshape_)
					if idx != -1 {
						_bottomstackofgrowthcurve.BottomStackGrowthCurveStartArcShapes = slices.Delete(_bottomstackofgrowthcurve.BottomStackGrowthCurveStartArcShapes, idx, idx+1)
						bottomstackgrowthcurvestartarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_bottomstackofgrowthcurve, "BottomStackGrowthCurveStartArcShapes", &_bottomstackofgrowthcurve.BottomStackGrowthCurveStartArcShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if bottomstackgrowthcurvestartarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bottomstackgrowthcurvestartarcshape_.Unstage(bottomstackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest)
	}

	bottomstackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.BottomStackGrowthCurveStartArcShape](
		bottomstackgrowthcurvestartarcshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if bottomstackgrowthcurvestartarcshapeFormCallback.CreationMode || bottomstackgrowthcurvestartarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bottomstackgrowthcurvestartarcshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(bottomstackgrowthcurvestartarcshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BottomStackGrowthCurveStartArcShapeFormCallback(
			nil,
			bottomstackgrowthcurvestartarcshapeFormCallback.probe,
			newFormGroup,
		)
		bottomstackgrowthcurvestartarcshape := new(models.BottomStackGrowthCurveStartArcShape)
		FillUpForm(bottomstackgrowthcurvestartarcshape, newFormGroup, bottomstackgrowthcurvestartarcshapeFormCallback.probe)
		bottomstackgrowthcurvestartarcshapeFormCallback.probe.formStage.Commit()
	}

	bottomstackgrowthcurvestartarcshapeFormCallback.probe.ux_tree()
}
func __gong__New__BottomStackOfGrowthCurveFormCallback(
	bottomstackofgrowthcurve *models.BottomStackOfGrowthCurve,
	probe *Probe,
	formGroup *form.FormGroup,
) (bottomstackofgrowthcurveFormCallback *BottomStackOfGrowthCurveFormCallback) {
	bottomstackofgrowthcurveFormCallback = new(BottomStackOfGrowthCurveFormCallback)
	bottomstackofgrowthcurveFormCallback.probe = probe
	bottomstackofgrowthcurveFormCallback.bottomstackofgrowthcurve = bottomstackofgrowthcurve
	bottomstackofgrowthcurveFormCallback.formGroup = formGroup

	bottomstackofgrowthcurveFormCallback.CreationMode = (bottomstackofgrowthcurve == nil)

	return
}

type BottomStackOfGrowthCurveFormCallback struct {
	bottomstackofgrowthcurve *models.BottomStackOfGrowthCurve

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (bottomstackofgrowthcurveFormCallback *BottomStackOfGrowthCurveFormCallback) OnSave() {
	bottomstackofgrowthcurveFormCallback.probe.stageOfInterest.Lock()
	defer bottomstackofgrowthcurveFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("BottomStackOfGrowthCurveFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bottomstackofgrowthcurveFormCallback.probe.formStage.Checkout()

	if bottomstackofgrowthcurveFormCallback.bottomstackofgrowthcurve == nil {
		bottomstackofgrowthcurveFormCallback.bottomstackofgrowthcurve = new(models.BottomStackOfGrowthCurve).Stage(bottomstackofgrowthcurveFormCallback.probe.stageOfInterest)
	}
	bottomstackofgrowthcurve_ := bottomstackofgrowthcurveFormCallback.bottomstackofgrowthcurve
	_ = bottomstackofgrowthcurve_

	for _, formDiv := range bottomstackofgrowthcurveFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(bottomstackofgrowthcurve_.Name), formDiv)
		case "BottomStackGrowthCurveStartArcShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.BottomStackGrowthCurveStartArcShape](bottomstackofgrowthcurveFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.BottomStackGrowthCurveStartArcShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.BottomStackGrowthCurveStartArcShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					bottomstackofgrowthcurveFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.BottomStackGrowthCurveStartArcShape](bottomstackofgrowthcurveFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			bottomstackofgrowthcurve_.BottomStackGrowthCurveStartArcShapes = instanceSlice
			bottomstackofgrowthcurveFormCallback.probe.UpdateSliceOfPointersCallback(bottomstackofgrowthcurve_, "BottomStackGrowthCurveStartArcShapes", &bottomstackofgrowthcurve_.BottomStackGrowthCurveStartArcShapes)

		case "BottomStackGrowthCurveEndArcShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.BottomStackGrowthCurveEndArcShape](bottomstackofgrowthcurveFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.BottomStackGrowthCurveEndArcShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.BottomStackGrowthCurveEndArcShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					bottomstackofgrowthcurveFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.BottomStackGrowthCurveEndArcShape](bottomstackofgrowthcurveFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			bottomstackofgrowthcurve_.BottomStackGrowthCurveEndArcShapes = instanceSlice
			bottomstackofgrowthcurveFormCallback.probe.UpdateSliceOfPointersCallback(bottomstackofgrowthcurve_, "BottomStackGrowthCurveEndArcShapes", &bottomstackofgrowthcurve_.BottomStackGrowthCurveEndArcShapes)

		}
	}

	// manage the suppress operation
	if bottomstackofgrowthcurveFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bottomstackofgrowthcurve_.Unstage(bottomstackofgrowthcurveFormCallback.probe.stageOfInterest)
	}

	bottomstackofgrowthcurveFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.BottomStackOfGrowthCurve](
		bottomstackofgrowthcurveFormCallback.probe,
	)

	// display a new form by reset the form stage
	if bottomstackofgrowthcurveFormCallback.CreationMode || bottomstackofgrowthcurveFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bottomstackofgrowthcurveFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(bottomstackofgrowthcurveFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BottomStackOfGrowthCurveFormCallback(
			nil,
			bottomstackofgrowthcurveFormCallback.probe,
			newFormGroup,
		)
		bottomstackofgrowthcurve := new(models.BottomStackOfGrowthCurve)
		FillUpForm(bottomstackofgrowthcurve, newFormGroup, bottomstackofgrowthcurveFormCallback.probe)
		bottomstackofgrowthcurveFormCallback.probe.formStage.Commit()
	}

	bottomstackofgrowthcurveFormCallback.probe.ux_tree()
}
func __gong__New__BottomStartArcShapeFormCallback(
	bottomstartarcshape *models.BottomStartArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (bottomstartarcshapeFormCallback *BottomStartArcShapeFormCallback) {
	bottomstartarcshapeFormCallback = new(BottomStartArcShapeFormCallback)
	bottomstartarcshapeFormCallback.probe = probe
	bottomstartarcshapeFormCallback.bottomstartarcshape = bottomstartarcshape
	bottomstartarcshapeFormCallback.formGroup = formGroup

	bottomstartarcshapeFormCallback.CreationMode = (bottomstartarcshape == nil)

	return
}

type BottomStartArcShapeFormCallback struct {
	bottomstartarcshape *models.BottomStartArcShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (bottomstartarcshapeFormCallback *BottomStartArcShapeFormCallback) OnSave() {
	bottomstartarcshapeFormCallback.probe.stageOfInterest.Lock()
	defer bottomstartarcshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("BottomStartArcShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bottomstartarcshapeFormCallback.probe.formStage.Checkout()

	if bottomstartarcshapeFormCallback.bottomstartarcshape == nil {
		bottomstartarcshapeFormCallback.bottomstartarcshape = new(models.BottomStartArcShape).Stage(bottomstartarcshapeFormCallback.probe.stageOfInterest)
	}
	bottomstartarcshape_ := bottomstartarcshapeFormCallback.bottomstartarcshape
	_ = bottomstartarcshape_

	for _, formDiv := range bottomstartarcshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(bottomstartarcshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(bottomstartarcshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(bottomstartarcshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(bottomstartarcshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(bottomstartarcshape_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(bottomstartarcshape_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(bottomstartarcshape_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(bottomstartarcshape_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(bottomstartarcshape_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(bottomstartarcshape_.RadiusY), formDiv)
		case "BottomStartArcShapeGrid:BottomStartArcShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the BottomStartArcShapeGrid instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target BottomStartArcShapeGrid instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.BottomStartArcShapeGrid](bottomstartarcshapeFormCallback.probe.stageOfInterest)
			targetBottomStartArcShapeGridIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetBottomStartArcShapeGridIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all BottomStartArcShapeGrid instances and update their BottomStartArcShapes slice
			for _bottomstartarcshapegrid := range *models.GetGongstructInstancesSetFromPointerType[*models.BottomStartArcShapeGrid](bottomstartarcshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(bottomstartarcshapeFormCallback.probe.stageOfInterest, _bottomstartarcshapegrid)
				
				// if BottomStartArcShapeGrid is selected
				if targetBottomStartArcShapeGridIDs[id] {
					// ensure bottomstartarcshape_ is in _bottomstartarcshapegrid.BottomStartArcShapes
					found := false
					for _, _b := range _bottomstartarcshapegrid.BottomStartArcShapes {
						if _b == bottomstartarcshape_ {
							found = true
							break
						}
					}
					if !found {
						_bottomstartarcshapegrid.BottomStartArcShapes = append(_bottomstartarcshapegrid.BottomStartArcShapes, bottomstartarcshape_)
						bottomstartarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_bottomstartarcshapegrid, "BottomStartArcShapes", &_bottomstartarcshapegrid.BottomStartArcShapes)
					}
				} else {
					// ensure bottomstartarcshape_ is NOT in _bottomstartarcshapegrid.BottomStartArcShapes
					idx := slices.Index(_bottomstartarcshapegrid.BottomStartArcShapes, bottomstartarcshape_)
					if idx != -1 {
						_bottomstartarcshapegrid.BottomStartArcShapes = slices.Delete(_bottomstartarcshapegrid.BottomStartArcShapes, idx, idx+1)
						bottomstartarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_bottomstartarcshapegrid, "BottomStartArcShapes", &_bottomstartarcshapegrid.BottomStartArcShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if bottomstartarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bottomstartarcshape_.Unstage(bottomstartarcshapeFormCallback.probe.stageOfInterest)
	}

	bottomstartarcshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.BottomStartArcShape](
		bottomstartarcshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if bottomstartarcshapeFormCallback.CreationMode || bottomstartarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bottomstartarcshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(bottomstartarcshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BottomStartArcShapeFormCallback(
			nil,
			bottomstartarcshapeFormCallback.probe,
			newFormGroup,
		)
		bottomstartarcshape := new(models.BottomStartArcShape)
		FillUpForm(bottomstartarcshape, newFormGroup, bottomstartarcshapeFormCallback.probe)
		bottomstartarcshapeFormCallback.probe.formStage.Commit()
	}

	bottomstartarcshapeFormCallback.probe.ux_tree()
}
func __gong__New__BottomStartArcShapeGridFormCallback(
	bottomstartarcshapegrid *models.BottomStartArcShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (bottomstartarcshapegridFormCallback *BottomStartArcShapeGridFormCallback) {
	bottomstartarcshapegridFormCallback = new(BottomStartArcShapeGridFormCallback)
	bottomstartarcshapegridFormCallback.probe = probe
	bottomstartarcshapegridFormCallback.bottomstartarcshapegrid = bottomstartarcshapegrid
	bottomstartarcshapegridFormCallback.formGroup = formGroup

	bottomstartarcshapegridFormCallback.CreationMode = (bottomstartarcshapegrid == nil)

	return
}

type BottomStartArcShapeGridFormCallback struct {
	bottomstartarcshapegrid *models.BottomStartArcShapeGrid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (bottomstartarcshapegridFormCallback *BottomStartArcShapeGridFormCallback) OnSave() {
	bottomstartarcshapegridFormCallback.probe.stageOfInterest.Lock()
	defer bottomstartarcshapegridFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("BottomStartArcShapeGridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bottomstartarcshapegridFormCallback.probe.formStage.Checkout()

	if bottomstartarcshapegridFormCallback.bottomstartarcshapegrid == nil {
		bottomstartarcshapegridFormCallback.bottomstartarcshapegrid = new(models.BottomStartArcShapeGrid).Stage(bottomstartarcshapegridFormCallback.probe.stageOfInterest)
	}
	bottomstartarcshapegrid_ := bottomstartarcshapegridFormCallback.bottomstartarcshapegrid
	_ = bottomstartarcshapegrid_

	for _, formDiv := range bottomstartarcshapegridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(bottomstartarcshapegrid_.Name), formDiv)
		case "BottomStartArcShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.BottomStartArcShape](bottomstartarcshapegridFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.BottomStartArcShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.BottomStartArcShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					bottomstartarcshapegridFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.BottomStartArcShape](bottomstartarcshapegridFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			bottomstartarcshapegrid_.BottomStartArcShapes = instanceSlice
			bottomstartarcshapegridFormCallback.probe.UpdateSliceOfPointersCallback(bottomstartarcshapegrid_, "BottomStartArcShapes", &bottomstartarcshapegrid_.BottomStartArcShapes)

		}
	}

	// manage the suppress operation
	if bottomstartarcshapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bottomstartarcshapegrid_.Unstage(bottomstartarcshapegridFormCallback.probe.stageOfInterest)
	}

	bottomstartarcshapegridFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.BottomStartArcShapeGrid](
		bottomstartarcshapegridFormCallback.probe,
	)

	// display a new form by reset the form stage
	if bottomstartarcshapegridFormCallback.CreationMode || bottomstartarcshapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bottomstartarcshapegridFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(bottomstartarcshapegridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BottomStartArcShapeGridFormCallback(
			nil,
			bottomstartarcshapegridFormCallback.probe,
			newFormGroup,
		)
		bottomstartarcshapegrid := new(models.BottomStartArcShapeGrid)
		FillUpForm(bottomstartarcshapegrid, newFormGroup, bottomstartarcshapegridFormCallback.probe)
		bottomstartarcshapegridFormCallback.probe.formStage.Commit()
	}

	bottomstartarcshapegridFormCallback.probe.ux_tree()
}
func __gong__New__CircleGridShapeFormCallback(
	circlegridshape *models.CircleGridShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (circlegridshapeFormCallback *CircleGridShapeFormCallback) {
	circlegridshapeFormCallback = new(CircleGridShapeFormCallback)
	circlegridshapeFormCallback.probe = probe
	circlegridshapeFormCallback.circlegridshape = circlegridshape
	circlegridshapeFormCallback.formGroup = formGroup

	circlegridshapeFormCallback.CreationMode = (circlegridshape == nil)

	return
}

type CircleGridShapeFormCallback struct {
	circlegridshape *models.CircleGridShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (circlegridshapeFormCallback *CircleGridShapeFormCallback) OnSave() {
	circlegridshapeFormCallback.probe.stageOfInterest.Lock()
	defer circlegridshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("CircleGridShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	circlegridshapeFormCallback.probe.formStage.Checkout()

	if circlegridshapeFormCallback.circlegridshape == nil {
		circlegridshapeFormCallback.circlegridshape = new(models.CircleGridShape).Stage(circlegridshapeFormCallback.probe.stageOfInterest)
	}
	circlegridshape_ := circlegridshapeFormCallback.circlegridshape
	_ = circlegridshape_

	for _, formDiv := range circlegridshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(circlegridshape_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if circlegridshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		circlegridshape_.Unstage(circlegridshapeFormCallback.probe.stageOfInterest)
	}

	circlegridshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.CircleGridShape](
		circlegridshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if circlegridshapeFormCallback.CreationMode || circlegridshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		circlegridshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(circlegridshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CircleGridShapeFormCallback(
			nil,
			circlegridshapeFormCallback.probe,
			newFormGroup,
		)
		circlegridshape := new(models.CircleGridShape)
		FillUpForm(circlegridshape, newFormGroup, circlegridshapeFormCallback.probe)
		circlegridshapeFormCallback.probe.formStage.Commit()
	}

	circlegridshapeFormCallback.probe.ux_tree()
}
func __gong__New__EndArcShapeFormCallback(
	endarcshape *models.EndArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (endarcshapeFormCallback *EndArcShapeFormCallback) {
	endarcshapeFormCallback = new(EndArcShapeFormCallback)
	endarcshapeFormCallback.probe = probe
	endarcshapeFormCallback.endarcshape = endarcshape
	endarcshapeFormCallback.formGroup = formGroup

	endarcshapeFormCallback.CreationMode = (endarcshape == nil)

	return
}

type EndArcShapeFormCallback struct {
	endarcshape *models.EndArcShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (endarcshapeFormCallback *EndArcShapeFormCallback) OnSave() {
	endarcshapeFormCallback.probe.stageOfInterest.Lock()
	defer endarcshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("EndArcShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	endarcshapeFormCallback.probe.formStage.Checkout()

	if endarcshapeFormCallback.endarcshape == nil {
		endarcshapeFormCallback.endarcshape = new(models.EndArcShape).Stage(endarcshapeFormCallback.probe.stageOfInterest)
	}
	endarcshape_ := endarcshapeFormCallback.endarcshape
	_ = endarcshape_

	for _, formDiv := range endarcshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(endarcshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(endarcshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(endarcshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(endarcshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(endarcshape_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(endarcshape_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(endarcshape_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(endarcshape_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(endarcshape_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(endarcshape_.RadiusY), formDiv)
		case "EndArcShapeGrid:EndArcShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the EndArcShapeGrid instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target EndArcShapeGrid instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.EndArcShapeGrid](endarcshapeFormCallback.probe.stageOfInterest)
			targetEndArcShapeGridIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetEndArcShapeGridIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all EndArcShapeGrid instances and update their EndArcShapes slice
			for _endarcshapegrid := range *models.GetGongstructInstancesSetFromPointerType[*models.EndArcShapeGrid](endarcshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(endarcshapeFormCallback.probe.stageOfInterest, _endarcshapegrid)
				
				// if EndArcShapeGrid is selected
				if targetEndArcShapeGridIDs[id] {
					// ensure endarcshape_ is in _endarcshapegrid.EndArcShapes
					found := false
					for _, _b := range _endarcshapegrid.EndArcShapes {
						if _b == endarcshape_ {
							found = true
							break
						}
					}
					if !found {
						_endarcshapegrid.EndArcShapes = append(_endarcshapegrid.EndArcShapes, endarcshape_)
						endarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_endarcshapegrid, "EndArcShapes", &_endarcshapegrid.EndArcShapes)
					}
				} else {
					// ensure endarcshape_ is NOT in _endarcshapegrid.EndArcShapes
					idx := slices.Index(_endarcshapegrid.EndArcShapes, endarcshape_)
					if idx != -1 {
						_endarcshapegrid.EndArcShapes = slices.Delete(_endarcshapegrid.EndArcShapes, idx, idx+1)
						endarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_endarcshapegrid, "EndArcShapes", &_endarcshapegrid.EndArcShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if endarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		endarcshape_.Unstage(endarcshapeFormCallback.probe.stageOfInterest)
	}

	endarcshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.EndArcShape](
		endarcshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if endarcshapeFormCallback.CreationMode || endarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		endarcshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(endarcshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__EndArcShapeFormCallback(
			nil,
			endarcshapeFormCallback.probe,
			newFormGroup,
		)
		endarcshape := new(models.EndArcShape)
		FillUpForm(endarcshape, newFormGroup, endarcshapeFormCallback.probe)
		endarcshapeFormCallback.probe.formStage.Commit()
	}

	endarcshapeFormCallback.probe.ux_tree()
}
func __gong__New__EndArcShapeGridFormCallback(
	endarcshapegrid *models.EndArcShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (endarcshapegridFormCallback *EndArcShapeGridFormCallback) {
	endarcshapegridFormCallback = new(EndArcShapeGridFormCallback)
	endarcshapegridFormCallback.probe = probe
	endarcshapegridFormCallback.endarcshapegrid = endarcshapegrid
	endarcshapegridFormCallback.formGroup = formGroup

	endarcshapegridFormCallback.CreationMode = (endarcshapegrid == nil)

	return
}

type EndArcShapeGridFormCallback struct {
	endarcshapegrid *models.EndArcShapeGrid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (endarcshapegridFormCallback *EndArcShapeGridFormCallback) OnSave() {
	endarcshapegridFormCallback.probe.stageOfInterest.Lock()
	defer endarcshapegridFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("EndArcShapeGridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	endarcshapegridFormCallback.probe.formStage.Checkout()

	if endarcshapegridFormCallback.endarcshapegrid == nil {
		endarcshapegridFormCallback.endarcshapegrid = new(models.EndArcShapeGrid).Stage(endarcshapegridFormCallback.probe.stageOfInterest)
	}
	endarcshapegrid_ := endarcshapegridFormCallback.endarcshapegrid
	_ = endarcshapegrid_

	for _, formDiv := range endarcshapegridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(endarcshapegrid_.Name), formDiv)
		case "EndArcShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.EndArcShape](endarcshapegridFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.EndArcShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.EndArcShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					endarcshapegridFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.EndArcShape](endarcshapegridFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			endarcshapegrid_.EndArcShapes = instanceSlice
			endarcshapegridFormCallback.probe.UpdateSliceOfPointersCallback(endarcshapegrid_, "EndArcShapes", &endarcshapegrid_.EndArcShapes)

		}
	}

	// manage the suppress operation
	if endarcshapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		endarcshapegrid_.Unstage(endarcshapegridFormCallback.probe.stageOfInterest)
	}

	endarcshapegridFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.EndArcShapeGrid](
		endarcshapegridFormCallback.probe,
	)

	// display a new form by reset the form stage
	if endarcshapegridFormCallback.CreationMode || endarcshapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		endarcshapegridFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(endarcshapegridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__EndArcShapeGridFormCallback(
			nil,
			endarcshapegridFormCallback.probe,
			newFormGroup,
		)
		endarcshapegrid := new(models.EndArcShapeGrid)
		FillUpForm(endarcshapegrid, newFormGroup, endarcshapegridFormCallback.probe)
		endarcshapegridFormCallback.probe.formStage.Commit()
	}

	endarcshapegridFormCallback.probe.ux_tree()
}
func __gong__New__ExplanationTextShapeFormCallback(
	explanationtextshape *models.ExplanationTextShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (explanationtextshapeFormCallback *ExplanationTextShapeFormCallback) {
	explanationtextshapeFormCallback = new(ExplanationTextShapeFormCallback)
	explanationtextshapeFormCallback.probe = probe
	explanationtextshapeFormCallback.explanationtextshape = explanationtextshape
	explanationtextshapeFormCallback.formGroup = formGroup

	explanationtextshapeFormCallback.CreationMode = (explanationtextshape == nil)

	return
}

type ExplanationTextShapeFormCallback struct {
	explanationtextshape *models.ExplanationTextShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (explanationtextshapeFormCallback *ExplanationTextShapeFormCallback) OnSave() {
	explanationtextshapeFormCallback.probe.stageOfInterest.Lock()
	defer explanationtextshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ExplanationTextShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	explanationtextshapeFormCallback.probe.formStage.Checkout()

	if explanationtextshapeFormCallback.explanationtextshape == nil {
		explanationtextshapeFormCallback.explanationtextshape = new(models.ExplanationTextShape).Stage(explanationtextshapeFormCallback.probe.stageOfInterest)
	}
	explanationtextshape_ := explanationtextshapeFormCallback.explanationtextshape
	_ = explanationtextshape_

	for _, formDiv := range explanationtextshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(explanationtextshape_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if explanationtextshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		explanationtextshape_.Unstage(explanationtextshapeFormCallback.probe.stageOfInterest)
	}

	explanationtextshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ExplanationTextShape](
		explanationtextshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if explanationtextshapeFormCallback.CreationMode || explanationtextshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		explanationtextshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(explanationtextshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ExplanationTextShapeFormCallback(
			nil,
			explanationtextshapeFormCallback.probe,
			newFormGroup,
		)
		explanationtextshape := new(models.ExplanationTextShape)
		FillUpForm(explanationtextshape, newFormGroup, explanationtextshapeFormCallback.probe)
		explanationtextshapeFormCallback.probe.formStage.Commit()
	}

	explanationtextshapeFormCallback.probe.ux_tree()
}
func __gong__New__GridPathShapeFormCallback(
	gridpathshape *models.GridPathShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (gridpathshapeFormCallback *GridPathShapeFormCallback) {
	gridpathshapeFormCallback = new(GridPathShapeFormCallback)
	gridpathshapeFormCallback.probe = probe
	gridpathshapeFormCallback.gridpathshape = gridpathshape
	gridpathshapeFormCallback.formGroup = formGroup

	gridpathshapeFormCallback.CreationMode = (gridpathshape == nil)

	return
}

type GridPathShapeFormCallback struct {
	gridpathshape *models.GridPathShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (gridpathshapeFormCallback *GridPathShapeFormCallback) OnSave() {
	gridpathshapeFormCallback.probe.stageOfInterest.Lock()
	defer gridpathshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("GridPathShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	gridpathshapeFormCallback.probe.formStage.Checkout()

	if gridpathshapeFormCallback.gridpathshape == nil {
		gridpathshapeFormCallback.gridpathshape = new(models.GridPathShape).Stage(gridpathshapeFormCallback.probe.stageOfInterest)
	}
	gridpathshape_ := gridpathshapeFormCallback.gridpathshape
	_ = gridpathshape_

	for _, formDiv := range gridpathshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(gridpathshape_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if gridpathshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gridpathshape_.Unstage(gridpathshapeFormCallback.probe.stageOfInterest)
	}

	gridpathshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.GridPathShape](
		gridpathshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if gridpathshapeFormCallback.CreationMode || gridpathshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		gridpathshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(gridpathshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GridPathShapeFormCallback(
			nil,
			gridpathshapeFormCallback.probe,
			newFormGroup,
		)
		gridpathshape := new(models.GridPathShape)
		FillUpForm(gridpathshape, newFormGroup, gridpathshapeFormCallback.probe)
		gridpathshapeFormCallback.probe.formStage.Commit()
	}

	gridpathshapeFormCallback.probe.ux_tree()
}
func __gong__New__GrowthCurve2DFormCallback(
	growthcurve2d *models.GrowthCurve2D,
	probe *Probe,
	formGroup *form.FormGroup,
) (growthcurve2dFormCallback *GrowthCurve2DFormCallback) {
	growthcurve2dFormCallback = new(GrowthCurve2DFormCallback)
	growthcurve2dFormCallback.probe = probe
	growthcurve2dFormCallback.growthcurve2d = growthcurve2d
	growthcurve2dFormCallback.formGroup = formGroup

	growthcurve2dFormCallback.CreationMode = (growthcurve2d == nil)

	return
}

type GrowthCurve2DFormCallback struct {
	growthcurve2d *models.GrowthCurve2D

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (growthcurve2dFormCallback *GrowthCurve2DFormCallback) OnSave() {
	growthcurve2dFormCallback.probe.stageOfInterest.Lock()
	defer growthcurve2dFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("GrowthCurve2DFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	growthcurve2dFormCallback.probe.formStage.Checkout()

	if growthcurve2dFormCallback.growthcurve2d == nil {
		growthcurve2dFormCallback.growthcurve2d = new(models.GrowthCurve2D).Stage(growthcurve2dFormCallback.probe.stageOfInterest)
	}
	growthcurve2d_ := growthcurve2dFormCallback.growthcurve2d
	_ = growthcurve2d_

	for _, formDiv := range growthcurve2dFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(growthcurve2d_.Name), formDiv)
		case "StartArcShapeGrid":
			FormDivSelectFieldToField(&(growthcurve2d_.StartArcShapeGrid), growthcurve2dFormCallback.probe.stageOfInterest, formDiv)
		case "EndArcShapeGrid":
			FormDivSelectFieldToField(&(growthcurve2d_.EndArcShapeGrid), growthcurve2dFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if growthcurve2dFormCallback.formGroup.HasSuppressButtonBeenPressed {
		growthcurve2d_.Unstage(growthcurve2dFormCallback.probe.stageOfInterest)
	}

	growthcurve2dFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.GrowthCurve2D](
		growthcurve2dFormCallback.probe,
	)

	// display a new form by reset the form stage
	if growthcurve2dFormCallback.CreationMode || growthcurve2dFormCallback.formGroup.HasSuppressButtonBeenPressed {
		growthcurve2dFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(growthcurve2dFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GrowthCurve2DFormCallback(
			nil,
			growthcurve2dFormCallback.probe,
			newFormGroup,
		)
		growthcurve2d := new(models.GrowthCurve2D)
		FillUpForm(growthcurve2d, newFormGroup, growthcurve2dFormCallback.probe)
		growthcurve2dFormCallback.probe.formStage.Commit()
	}

	growthcurve2dFormCallback.probe.ux_tree()
}
func __gong__New__GrowthCurveBezierShapeFormCallback(
	growthcurvebeziershape *models.GrowthCurveBezierShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (growthcurvebeziershapeFormCallback *GrowthCurveBezierShapeFormCallback) {
	growthcurvebeziershapeFormCallback = new(GrowthCurveBezierShapeFormCallback)
	growthcurvebeziershapeFormCallback.probe = probe
	growthcurvebeziershapeFormCallback.growthcurvebeziershape = growthcurvebeziershape
	growthcurvebeziershapeFormCallback.formGroup = formGroup

	growthcurvebeziershapeFormCallback.CreationMode = (growthcurvebeziershape == nil)

	return
}

type GrowthCurveBezierShapeFormCallback struct {
	growthcurvebeziershape *models.GrowthCurveBezierShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (growthcurvebeziershapeFormCallback *GrowthCurveBezierShapeFormCallback) OnSave() {
	growthcurvebeziershapeFormCallback.probe.stageOfInterest.Lock()
	defer growthcurvebeziershapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("GrowthCurveBezierShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	growthcurvebeziershapeFormCallback.probe.formStage.Checkout()

	if growthcurvebeziershapeFormCallback.growthcurvebeziershape == nil {
		growthcurvebeziershapeFormCallback.growthcurvebeziershape = new(models.GrowthCurveBezierShape).Stage(growthcurvebeziershapeFormCallback.probe.stageOfInterest)
	}
	growthcurvebeziershape_ := growthcurvebeziershapeFormCallback.growthcurvebeziershape
	_ = growthcurvebeziershape_

	for _, formDiv := range growthcurvebeziershapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(growthcurvebeziershape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(growthcurvebeziershape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(growthcurvebeziershape_.StartY), formDiv)
		case "ControlPointStartX":
			FormDivBasicFieldToField(&(growthcurvebeziershape_.ControlPointStartX), formDiv)
		case "ControlPointStartY":
			FormDivBasicFieldToField(&(growthcurvebeziershape_.ControlPointStartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(growthcurvebeziershape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(growthcurvebeziershape_.EndY), formDiv)
		case "ControlPointEndX":
			FormDivBasicFieldToField(&(growthcurvebeziershape_.ControlPointEndX), formDiv)
		case "ControlPointEndY":
			FormDivBasicFieldToField(&(growthcurvebeziershape_.ControlPointEndY), formDiv)
		case "GrowthCurveBezierShapeGrid:GrowthCurveBezierShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the GrowthCurveBezierShapeGrid instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target GrowthCurveBezierShapeGrid instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.GrowthCurveBezierShapeGrid](growthcurvebeziershapeFormCallback.probe.stageOfInterest)
			targetGrowthCurveBezierShapeGridIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetGrowthCurveBezierShapeGridIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all GrowthCurveBezierShapeGrid instances and update their GrowthCurveBezierShapes slice
			for _growthcurvebeziershapegrid := range *models.GetGongstructInstancesSetFromPointerType[*models.GrowthCurveBezierShapeGrid](growthcurvebeziershapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(growthcurvebeziershapeFormCallback.probe.stageOfInterest, _growthcurvebeziershapegrid)
				
				// if GrowthCurveBezierShapeGrid is selected
				if targetGrowthCurveBezierShapeGridIDs[id] {
					// ensure growthcurvebeziershape_ is in _growthcurvebeziershapegrid.GrowthCurveBezierShapes
					found := false
					for _, _b := range _growthcurvebeziershapegrid.GrowthCurveBezierShapes {
						if _b == growthcurvebeziershape_ {
							found = true
							break
						}
					}
					if !found {
						_growthcurvebeziershapegrid.GrowthCurveBezierShapes = append(_growthcurvebeziershapegrid.GrowthCurveBezierShapes, growthcurvebeziershape_)
						growthcurvebeziershapeFormCallback.probe.UpdateSliceOfPointersCallback(_growthcurvebeziershapegrid, "GrowthCurveBezierShapes", &_growthcurvebeziershapegrid.GrowthCurveBezierShapes)
					}
				} else {
					// ensure growthcurvebeziershape_ is NOT in _growthcurvebeziershapegrid.GrowthCurveBezierShapes
					idx := slices.Index(_growthcurvebeziershapegrid.GrowthCurveBezierShapes, growthcurvebeziershape_)
					if idx != -1 {
						_growthcurvebeziershapegrid.GrowthCurveBezierShapes = slices.Delete(_growthcurvebeziershapegrid.GrowthCurveBezierShapes, idx, idx+1)
						growthcurvebeziershapeFormCallback.probe.UpdateSliceOfPointersCallback(_growthcurvebeziershapegrid, "GrowthCurveBezierShapes", &_growthcurvebeziershapegrid.GrowthCurveBezierShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if growthcurvebeziershapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		growthcurvebeziershape_.Unstage(growthcurvebeziershapeFormCallback.probe.stageOfInterest)
	}

	growthcurvebeziershapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.GrowthCurveBezierShape](
		growthcurvebeziershapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if growthcurvebeziershapeFormCallback.CreationMode || growthcurvebeziershapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		growthcurvebeziershapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(growthcurvebeziershapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GrowthCurveBezierShapeFormCallback(
			nil,
			growthcurvebeziershapeFormCallback.probe,
			newFormGroup,
		)
		growthcurvebeziershape := new(models.GrowthCurveBezierShape)
		FillUpForm(growthcurvebeziershape, newFormGroup, growthcurvebeziershapeFormCallback.probe)
		growthcurvebeziershapeFormCallback.probe.formStage.Commit()
	}

	growthcurvebeziershapeFormCallback.probe.ux_tree()
}
func __gong__New__GrowthCurveBezierShapeGridFormCallback(
	growthcurvebeziershapegrid *models.GrowthCurveBezierShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (growthcurvebeziershapegridFormCallback *GrowthCurveBezierShapeGridFormCallback) {
	growthcurvebeziershapegridFormCallback = new(GrowthCurveBezierShapeGridFormCallback)
	growthcurvebeziershapegridFormCallback.probe = probe
	growthcurvebeziershapegridFormCallback.growthcurvebeziershapegrid = growthcurvebeziershapegrid
	growthcurvebeziershapegridFormCallback.formGroup = formGroup

	growthcurvebeziershapegridFormCallback.CreationMode = (growthcurvebeziershapegrid == nil)

	return
}

type GrowthCurveBezierShapeGridFormCallback struct {
	growthcurvebeziershapegrid *models.GrowthCurveBezierShapeGrid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (growthcurvebeziershapegridFormCallback *GrowthCurveBezierShapeGridFormCallback) OnSave() {
	growthcurvebeziershapegridFormCallback.probe.stageOfInterest.Lock()
	defer growthcurvebeziershapegridFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("GrowthCurveBezierShapeGridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	growthcurvebeziershapegridFormCallback.probe.formStage.Checkout()

	if growthcurvebeziershapegridFormCallback.growthcurvebeziershapegrid == nil {
		growthcurvebeziershapegridFormCallback.growthcurvebeziershapegrid = new(models.GrowthCurveBezierShapeGrid).Stage(growthcurvebeziershapegridFormCallback.probe.stageOfInterest)
	}
	growthcurvebeziershapegrid_ := growthcurvebeziershapegridFormCallback.growthcurvebeziershapegrid
	_ = growthcurvebeziershapegrid_

	for _, formDiv := range growthcurvebeziershapegridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(growthcurvebeziershapegrid_.Name), formDiv)
		case "GrowthCurveBezierShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.GrowthCurveBezierShape](growthcurvebeziershapegridFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.GrowthCurveBezierShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.GrowthCurveBezierShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					growthcurvebeziershapegridFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.GrowthCurveBezierShape](growthcurvebeziershapegridFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			growthcurvebeziershapegrid_.GrowthCurveBezierShapes = instanceSlice
			growthcurvebeziershapegridFormCallback.probe.UpdateSliceOfPointersCallback(growthcurvebeziershapegrid_, "GrowthCurveBezierShapes", &growthcurvebeziershapegrid_.GrowthCurveBezierShapes)

		}
	}

	// manage the suppress operation
	if growthcurvebeziershapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		growthcurvebeziershapegrid_.Unstage(growthcurvebeziershapegridFormCallback.probe.stageOfInterest)
	}

	growthcurvebeziershapegridFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.GrowthCurveBezierShapeGrid](
		growthcurvebeziershapegridFormCallback.probe,
	)

	// display a new form by reset the form stage
	if growthcurvebeziershapegridFormCallback.CreationMode || growthcurvebeziershapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		growthcurvebeziershapegridFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(growthcurvebeziershapegridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GrowthCurveBezierShapeGridFormCallback(
			nil,
			growthcurvebeziershapegridFormCallback.probe,
			newFormGroup,
		)
		growthcurvebeziershapegrid := new(models.GrowthCurveBezierShapeGrid)
		FillUpForm(growthcurvebeziershapegrid, newFormGroup, growthcurvebeziershapegridFormCallback.probe)
		growthcurvebeziershapegridFormCallback.probe.formStage.Commit()
	}

	growthcurvebeziershapegridFormCallback.probe.ux_tree()
}
func __gong__New__GrowthCurveRhombusGridShapeFormCallback(
	growthcurverhombusgridshape *models.GrowthCurveRhombusGridShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (growthcurverhombusgridshapeFormCallback *GrowthCurveRhombusGridShapeFormCallback) {
	growthcurverhombusgridshapeFormCallback = new(GrowthCurveRhombusGridShapeFormCallback)
	growthcurverhombusgridshapeFormCallback.probe = probe
	growthcurverhombusgridshapeFormCallback.growthcurverhombusgridshape = growthcurverhombusgridshape
	growthcurverhombusgridshapeFormCallback.formGroup = formGroup

	growthcurverhombusgridshapeFormCallback.CreationMode = (growthcurverhombusgridshape == nil)

	return
}

type GrowthCurveRhombusGridShapeFormCallback struct {
	growthcurverhombusgridshape *models.GrowthCurveRhombusGridShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (growthcurverhombusgridshapeFormCallback *GrowthCurveRhombusGridShapeFormCallback) OnSave() {
	growthcurverhombusgridshapeFormCallback.probe.stageOfInterest.Lock()
	defer growthcurverhombusgridshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("GrowthCurveRhombusGridShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	growthcurverhombusgridshapeFormCallback.probe.formStage.Checkout()

	if growthcurverhombusgridshapeFormCallback.growthcurverhombusgridshape == nil {
		growthcurverhombusgridshapeFormCallback.growthcurverhombusgridshape = new(models.GrowthCurveRhombusGridShape).Stage(growthcurverhombusgridshapeFormCallback.probe.stageOfInterest)
	}
	growthcurverhombusgridshape_ := growthcurverhombusgridshapeFormCallback.growthcurverhombusgridshape
	_ = growthcurverhombusgridshape_

	for _, formDiv := range growthcurverhombusgridshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(growthcurverhombusgridshape_.Name), formDiv)
		case "GrowthCurveRhombusShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.GrowthCurveRhombusShape](growthcurverhombusgridshapeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.GrowthCurveRhombusShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.GrowthCurveRhombusShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					growthcurverhombusgridshapeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.GrowthCurveRhombusShape](growthcurverhombusgridshapeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			growthcurverhombusgridshape_.GrowthCurveRhombusShapes = instanceSlice
			growthcurverhombusgridshapeFormCallback.probe.UpdateSliceOfPointersCallback(growthcurverhombusgridshape_, "GrowthCurveRhombusShapes", &growthcurverhombusgridshape_.GrowthCurveRhombusShapes)

		}
	}

	// manage the suppress operation
	if growthcurverhombusgridshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		growthcurverhombusgridshape_.Unstage(growthcurverhombusgridshapeFormCallback.probe.stageOfInterest)
	}

	growthcurverhombusgridshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.GrowthCurveRhombusGridShape](
		growthcurverhombusgridshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if growthcurverhombusgridshapeFormCallback.CreationMode || growthcurverhombusgridshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		growthcurverhombusgridshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(growthcurverhombusgridshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GrowthCurveRhombusGridShapeFormCallback(
			nil,
			growthcurverhombusgridshapeFormCallback.probe,
			newFormGroup,
		)
		growthcurverhombusgridshape := new(models.GrowthCurveRhombusGridShape)
		FillUpForm(growthcurverhombusgridshape, newFormGroup, growthcurverhombusgridshapeFormCallback.probe)
		growthcurverhombusgridshapeFormCallback.probe.formStage.Commit()
	}

	growthcurverhombusgridshapeFormCallback.probe.ux_tree()
}
func __gong__New__GrowthCurveRhombusShapeFormCallback(
	growthcurverhombusshape *models.GrowthCurveRhombusShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (growthcurverhombusshapeFormCallback *GrowthCurveRhombusShapeFormCallback) {
	growthcurverhombusshapeFormCallback = new(GrowthCurveRhombusShapeFormCallback)
	growthcurverhombusshapeFormCallback.probe = probe
	growthcurverhombusshapeFormCallback.growthcurverhombusshape = growthcurverhombusshape
	growthcurverhombusshapeFormCallback.formGroup = formGroup

	growthcurverhombusshapeFormCallback.CreationMode = (growthcurverhombusshape == nil)

	return
}

type GrowthCurveRhombusShapeFormCallback struct {
	growthcurverhombusshape *models.GrowthCurveRhombusShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (growthcurverhombusshapeFormCallback *GrowthCurveRhombusShapeFormCallback) OnSave() {
	growthcurverhombusshapeFormCallback.probe.stageOfInterest.Lock()
	defer growthcurverhombusshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("GrowthCurveRhombusShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	growthcurverhombusshapeFormCallback.probe.formStage.Checkout()

	if growthcurverhombusshapeFormCallback.growthcurverhombusshape == nil {
		growthcurverhombusshapeFormCallback.growthcurverhombusshape = new(models.GrowthCurveRhombusShape).Stage(growthcurverhombusshapeFormCallback.probe.stageOfInterest)
	}
	growthcurverhombusshape_ := growthcurverhombusshapeFormCallback.growthcurverhombusshape
	_ = growthcurverhombusshape_

	for _, formDiv := range growthcurverhombusshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(growthcurverhombusshape_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(growthcurverhombusshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(growthcurverhombusshape_.Y), formDiv)
		case "GrowthCurveRhombusGridShape:GrowthCurveRhombusShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the GrowthCurveRhombusGridShape instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target GrowthCurveRhombusGridShape instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.GrowthCurveRhombusGridShape](growthcurverhombusshapeFormCallback.probe.stageOfInterest)
			targetGrowthCurveRhombusGridShapeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetGrowthCurveRhombusGridShapeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all GrowthCurveRhombusGridShape instances and update their GrowthCurveRhombusShapes slice
			for _growthcurverhombusgridshape := range *models.GetGongstructInstancesSetFromPointerType[*models.GrowthCurveRhombusGridShape](growthcurverhombusshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(growthcurverhombusshapeFormCallback.probe.stageOfInterest, _growthcurverhombusgridshape)
				
				// if GrowthCurveRhombusGridShape is selected
				if targetGrowthCurveRhombusGridShapeIDs[id] {
					// ensure growthcurverhombusshape_ is in _growthcurverhombusgridshape.GrowthCurveRhombusShapes
					found := false
					for _, _b := range _growthcurverhombusgridshape.GrowthCurveRhombusShapes {
						if _b == growthcurverhombusshape_ {
							found = true
							break
						}
					}
					if !found {
						_growthcurverhombusgridshape.GrowthCurveRhombusShapes = append(_growthcurverhombusgridshape.GrowthCurveRhombusShapes, growthcurverhombusshape_)
						growthcurverhombusshapeFormCallback.probe.UpdateSliceOfPointersCallback(_growthcurverhombusgridshape, "GrowthCurveRhombusShapes", &_growthcurverhombusgridshape.GrowthCurveRhombusShapes)
					}
				} else {
					// ensure growthcurverhombusshape_ is NOT in _growthcurverhombusgridshape.GrowthCurveRhombusShapes
					idx := slices.Index(_growthcurverhombusgridshape.GrowthCurveRhombusShapes, growthcurverhombusshape_)
					if idx != -1 {
						_growthcurverhombusgridshape.GrowthCurveRhombusShapes = slices.Delete(_growthcurverhombusgridshape.GrowthCurveRhombusShapes, idx, idx+1)
						growthcurverhombusshapeFormCallback.probe.UpdateSliceOfPointersCallback(_growthcurverhombusgridshape, "GrowthCurveRhombusShapes", &_growthcurverhombusgridshape.GrowthCurveRhombusShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if growthcurverhombusshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		growthcurverhombusshape_.Unstage(growthcurverhombusshapeFormCallback.probe.stageOfInterest)
	}

	growthcurverhombusshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.GrowthCurveRhombusShape](
		growthcurverhombusshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if growthcurverhombusshapeFormCallback.CreationMode || growthcurverhombusshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		growthcurverhombusshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(growthcurverhombusshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GrowthCurveRhombusShapeFormCallback(
			nil,
			growthcurverhombusshapeFormCallback.probe,
			newFormGroup,
		)
		growthcurverhombusshape := new(models.GrowthCurveRhombusShape)
		FillUpForm(growthcurverhombusshape, newFormGroup, growthcurverhombusshapeFormCallback.probe)
		growthcurverhombusshapeFormCallback.probe.formStage.Commit()
	}

	growthcurverhombusshapeFormCallback.probe.ux_tree()
}
func __gong__New__GrowthVectorShapeFormCallback(
	growthvectorshape *models.GrowthVectorShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (growthvectorshapeFormCallback *GrowthVectorShapeFormCallback) {
	growthvectorshapeFormCallback = new(GrowthVectorShapeFormCallback)
	growthvectorshapeFormCallback.probe = probe
	growthvectorshapeFormCallback.growthvectorshape = growthvectorshape
	growthvectorshapeFormCallback.formGroup = formGroup

	growthvectorshapeFormCallback.CreationMode = (growthvectorshape == nil)

	return
}

type GrowthVectorShapeFormCallback struct {
	growthvectorshape *models.GrowthVectorShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (growthvectorshapeFormCallback *GrowthVectorShapeFormCallback) OnSave() {
	growthvectorshapeFormCallback.probe.stageOfInterest.Lock()
	defer growthvectorshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("GrowthVectorShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	growthvectorshapeFormCallback.probe.formStage.Checkout()

	if growthvectorshapeFormCallback.growthvectorshape == nil {
		growthvectorshapeFormCallback.growthvectorshape = new(models.GrowthVectorShape).Stage(growthvectorshapeFormCallback.probe.stageOfInterest)
	}
	growthvectorshape_ := growthvectorshapeFormCallback.growthvectorshape
	_ = growthvectorshape_

	for _, formDiv := range growthvectorshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(growthvectorshape_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(growthvectorshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(growthvectorshape_.Y), formDiv)
		}
	}

	// manage the suppress operation
	if growthvectorshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		growthvectorshape_.Unstage(growthvectorshapeFormCallback.probe.stageOfInterest)
	}

	growthvectorshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.GrowthVectorShape](
		growthvectorshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if growthvectorshapeFormCallback.CreationMode || growthvectorshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		growthvectorshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(growthvectorshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GrowthVectorShapeFormCallback(
			nil,
			growthvectorshapeFormCallback.probe,
			newFormGroup,
		)
		growthvectorshape := new(models.GrowthVectorShape)
		FillUpForm(growthvectorshape, newFormGroup, growthvectorshapeFormCallback.probe)
		growthvectorshapeFormCallback.probe.formStage.Commit()
	}

	growthvectorshapeFormCallback.probe.ux_tree()
}
func __gong__New__InitialRhombusGridShapeFormCallback(
	initialrhombusgridshape *models.InitialRhombusGridShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (initialrhombusgridshapeFormCallback *InitialRhombusGridShapeFormCallback) {
	initialrhombusgridshapeFormCallback = new(InitialRhombusGridShapeFormCallback)
	initialrhombusgridshapeFormCallback.probe = probe
	initialrhombusgridshapeFormCallback.initialrhombusgridshape = initialrhombusgridshape
	initialrhombusgridshapeFormCallback.formGroup = formGroup

	initialrhombusgridshapeFormCallback.CreationMode = (initialrhombusgridshape == nil)

	return
}

type InitialRhombusGridShapeFormCallback struct {
	initialrhombusgridshape *models.InitialRhombusGridShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (initialrhombusgridshapeFormCallback *InitialRhombusGridShapeFormCallback) OnSave() {
	initialrhombusgridshapeFormCallback.probe.stageOfInterest.Lock()
	defer initialrhombusgridshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("InitialRhombusGridShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	initialrhombusgridshapeFormCallback.probe.formStage.Checkout()

	if initialrhombusgridshapeFormCallback.initialrhombusgridshape == nil {
		initialrhombusgridshapeFormCallback.initialrhombusgridshape = new(models.InitialRhombusGridShape).Stage(initialrhombusgridshapeFormCallback.probe.stageOfInterest)
	}
	initialrhombusgridshape_ := initialrhombusgridshapeFormCallback.initialrhombusgridshape
	_ = initialrhombusgridshape_

	for _, formDiv := range initialrhombusgridshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(initialrhombusgridshape_.Name), formDiv)
		case "InitialRhombusShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.InitialRhombusShape](initialrhombusgridshapeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.InitialRhombusShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.InitialRhombusShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					initialrhombusgridshapeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.InitialRhombusShape](initialrhombusgridshapeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			initialrhombusgridshape_.InitialRhombusShapes = instanceSlice
			initialrhombusgridshapeFormCallback.probe.UpdateSliceOfPointersCallback(initialrhombusgridshape_, "InitialRhombusShapes", &initialrhombusgridshape_.InitialRhombusShapes)

		}
	}

	// manage the suppress operation
	if initialrhombusgridshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		initialrhombusgridshape_.Unstage(initialrhombusgridshapeFormCallback.probe.stageOfInterest)
	}

	initialrhombusgridshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.InitialRhombusGridShape](
		initialrhombusgridshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if initialrhombusgridshapeFormCallback.CreationMode || initialrhombusgridshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		initialrhombusgridshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(initialrhombusgridshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__InitialRhombusGridShapeFormCallback(
			nil,
			initialrhombusgridshapeFormCallback.probe,
			newFormGroup,
		)
		initialrhombusgridshape := new(models.InitialRhombusGridShape)
		FillUpForm(initialrhombusgridshape, newFormGroup, initialrhombusgridshapeFormCallback.probe)
		initialrhombusgridshapeFormCallback.probe.formStage.Commit()
	}

	initialrhombusgridshapeFormCallback.probe.ux_tree()
}
func __gong__New__InitialRhombusShapeFormCallback(
	initialrhombusshape *models.InitialRhombusShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (initialrhombusshapeFormCallback *InitialRhombusShapeFormCallback) {
	initialrhombusshapeFormCallback = new(InitialRhombusShapeFormCallback)
	initialrhombusshapeFormCallback.probe = probe
	initialrhombusshapeFormCallback.initialrhombusshape = initialrhombusshape
	initialrhombusshapeFormCallback.formGroup = formGroup

	initialrhombusshapeFormCallback.CreationMode = (initialrhombusshape == nil)

	return
}

type InitialRhombusShapeFormCallback struct {
	initialrhombusshape *models.InitialRhombusShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (initialrhombusshapeFormCallback *InitialRhombusShapeFormCallback) OnSave() {
	initialrhombusshapeFormCallback.probe.stageOfInterest.Lock()
	defer initialrhombusshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("InitialRhombusShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	initialrhombusshapeFormCallback.probe.formStage.Checkout()

	if initialrhombusshapeFormCallback.initialrhombusshape == nil {
		initialrhombusshapeFormCallback.initialrhombusshape = new(models.InitialRhombusShape).Stage(initialrhombusshapeFormCallback.probe.stageOfInterest)
	}
	initialrhombusshape_ := initialrhombusshapeFormCallback.initialrhombusshape
	_ = initialrhombusshape_

	for _, formDiv := range initialrhombusshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(initialrhombusshape_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(initialrhombusshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(initialrhombusshape_.Y), formDiv)
		case "InitialRhombusGridShape:InitialRhombusShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the InitialRhombusGridShape instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target InitialRhombusGridShape instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.InitialRhombusGridShape](initialrhombusshapeFormCallback.probe.stageOfInterest)
			targetInitialRhombusGridShapeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetInitialRhombusGridShapeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all InitialRhombusGridShape instances and update their InitialRhombusShapes slice
			for _initialrhombusgridshape := range *models.GetGongstructInstancesSetFromPointerType[*models.InitialRhombusGridShape](initialrhombusshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(initialrhombusshapeFormCallback.probe.stageOfInterest, _initialrhombusgridshape)
				
				// if InitialRhombusGridShape is selected
				if targetInitialRhombusGridShapeIDs[id] {
					// ensure initialrhombusshape_ is in _initialrhombusgridshape.InitialRhombusShapes
					found := false
					for _, _b := range _initialrhombusgridshape.InitialRhombusShapes {
						if _b == initialrhombusshape_ {
							found = true
							break
						}
					}
					if !found {
						_initialrhombusgridshape.InitialRhombusShapes = append(_initialrhombusgridshape.InitialRhombusShapes, initialrhombusshape_)
						initialrhombusshapeFormCallback.probe.UpdateSliceOfPointersCallback(_initialrhombusgridshape, "InitialRhombusShapes", &_initialrhombusgridshape.InitialRhombusShapes)
					}
				} else {
					// ensure initialrhombusshape_ is NOT in _initialrhombusgridshape.InitialRhombusShapes
					idx := slices.Index(_initialrhombusgridshape.InitialRhombusShapes, initialrhombusshape_)
					if idx != -1 {
						_initialrhombusgridshape.InitialRhombusShapes = slices.Delete(_initialrhombusgridshape.InitialRhombusShapes, idx, idx+1)
						initialrhombusshapeFormCallback.probe.UpdateSliceOfPointersCallback(_initialrhombusgridshape, "InitialRhombusShapes", &_initialrhombusgridshape.InitialRhombusShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if initialrhombusshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		initialrhombusshape_.Unstage(initialrhombusshapeFormCallback.probe.stageOfInterest)
	}

	initialrhombusshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.InitialRhombusShape](
		initialrhombusshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if initialrhombusshapeFormCallback.CreationMode || initialrhombusshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		initialrhombusshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(initialrhombusshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__InitialRhombusShapeFormCallback(
			nil,
			initialrhombusshapeFormCallback.probe,
			newFormGroup,
		)
		initialrhombusshape := new(models.InitialRhombusShape)
		FillUpForm(initialrhombusshape, newFormGroup, initialrhombusshapeFormCallback.probe)
		initialrhombusshapeFormCallback.probe.formStage.Commit()
	}

	initialrhombusshapeFormCallback.probe.ux_tree()
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

		case "NbPixPerCharacter":
			FormDivBasicFieldToField(&(library_.NbPixPerCharacter), formDiv)
		case "LogoSVGFile":
			FormDivBasicFieldToField(&(library_.LogoSVGFile), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(library_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(library_.IsExpanded), formDiv)
		case "IsRootLibrary":
			FormDivBasicFieldToField(&(library_.IsRootLibrary), formDiv)
		case "Plants":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Plant](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Plant, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Plant)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Plant](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.Plants = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "Plants", &library_.Plants)

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
func __gong__New__NextCircleShapeFormCallback(
	nextcircleshape *models.NextCircleShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (nextcircleshapeFormCallback *NextCircleShapeFormCallback) {
	nextcircleshapeFormCallback = new(NextCircleShapeFormCallback)
	nextcircleshapeFormCallback.probe = probe
	nextcircleshapeFormCallback.nextcircleshape = nextcircleshape
	nextcircleshapeFormCallback.formGroup = formGroup

	nextcircleshapeFormCallback.CreationMode = (nextcircleshape == nil)

	return
}

type NextCircleShapeFormCallback struct {
	nextcircleshape *models.NextCircleShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (nextcircleshapeFormCallback *NextCircleShapeFormCallback) OnSave() {
	nextcircleshapeFormCallback.probe.stageOfInterest.Lock()
	defer nextcircleshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("NextCircleShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	nextcircleshapeFormCallback.probe.formStage.Checkout()

	if nextcircleshapeFormCallback.nextcircleshape == nil {
		nextcircleshapeFormCallback.nextcircleshape = new(models.NextCircleShape).Stage(nextcircleshapeFormCallback.probe.stageOfInterest)
	}
	nextcircleshape_ := nextcircleshapeFormCallback.nextcircleshape
	_ = nextcircleshape_

	for _, formDiv := range nextcircleshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(nextcircleshape_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if nextcircleshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		nextcircleshape_.Unstage(nextcircleshapeFormCallback.probe.stageOfInterest)
	}

	nextcircleshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.NextCircleShape](
		nextcircleshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if nextcircleshapeFormCallback.CreationMode || nextcircleshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		nextcircleshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(nextcircleshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__NextCircleShapeFormCallback(
			nil,
			nextcircleshapeFormCallback.probe,
			newFormGroup,
		)
		nextcircleshape := new(models.NextCircleShape)
		FillUpForm(nextcircleshape, newFormGroup, nextcircleshapeFormCallback.probe)
		nextcircleshapeFormCallback.probe.formStage.Commit()
	}

	nextcircleshapeFormCallback.probe.ux_tree()
}
func __gong__New__PerpendicularVectorFormCallback(
	perpendicularvector *models.PerpendicularVector,
	probe *Probe,
	formGroup *form.FormGroup,
) (perpendicularvectorFormCallback *PerpendicularVectorFormCallback) {
	perpendicularvectorFormCallback = new(PerpendicularVectorFormCallback)
	perpendicularvectorFormCallback.probe = probe
	perpendicularvectorFormCallback.perpendicularvector = perpendicularvector
	perpendicularvectorFormCallback.formGroup = formGroup

	perpendicularvectorFormCallback.CreationMode = (perpendicularvector == nil)

	return
}

type PerpendicularVectorFormCallback struct {
	perpendicularvector *models.PerpendicularVector

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (perpendicularvectorFormCallback *PerpendicularVectorFormCallback) OnSave() {
	perpendicularvectorFormCallback.probe.stageOfInterest.Lock()
	defer perpendicularvectorFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PerpendicularVectorFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	perpendicularvectorFormCallback.probe.formStage.Checkout()

	if perpendicularvectorFormCallback.perpendicularvector == nil {
		perpendicularvectorFormCallback.perpendicularvector = new(models.PerpendicularVector).Stage(perpendicularvectorFormCallback.probe.stageOfInterest)
	}
	perpendicularvector_ := perpendicularvectorFormCallback.perpendicularvector
	_ = perpendicularvector_

	for _, formDiv := range perpendicularvectorFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(perpendicularvector_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(perpendicularvector_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(perpendicularvector_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(perpendicularvector_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(perpendicularvector_.EndY), formDiv)
		case "PerpendicularVectorGrid:PerpendicularVectors":
			// 1. Decode the AssociationStorage which contains the rowIDs of the PerpendicularVectorGrid instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target PerpendicularVectorGrid instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.PerpendicularVectorGrid](perpendicularvectorFormCallback.probe.stageOfInterest)
			targetPerpendicularVectorGridIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetPerpendicularVectorGridIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all PerpendicularVectorGrid instances and update their PerpendicularVectors slice
			for _perpendicularvectorgrid := range *models.GetGongstructInstancesSetFromPointerType[*models.PerpendicularVectorGrid](perpendicularvectorFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(perpendicularvectorFormCallback.probe.stageOfInterest, _perpendicularvectorgrid)
				
				// if PerpendicularVectorGrid is selected
				if targetPerpendicularVectorGridIDs[id] {
					// ensure perpendicularvector_ is in _perpendicularvectorgrid.PerpendicularVectors
					found := false
					for _, _b := range _perpendicularvectorgrid.PerpendicularVectors {
						if _b == perpendicularvector_ {
							found = true
							break
						}
					}
					if !found {
						_perpendicularvectorgrid.PerpendicularVectors = append(_perpendicularvectorgrid.PerpendicularVectors, perpendicularvector_)
						perpendicularvectorFormCallback.probe.UpdateSliceOfPointersCallback(_perpendicularvectorgrid, "PerpendicularVectors", &_perpendicularvectorgrid.PerpendicularVectors)
					}
				} else {
					// ensure perpendicularvector_ is NOT in _perpendicularvectorgrid.PerpendicularVectors
					idx := slices.Index(_perpendicularvectorgrid.PerpendicularVectors, perpendicularvector_)
					if idx != -1 {
						_perpendicularvectorgrid.PerpendicularVectors = slices.Delete(_perpendicularvectorgrid.PerpendicularVectors, idx, idx+1)
						perpendicularvectorFormCallback.probe.UpdateSliceOfPointersCallback(_perpendicularvectorgrid, "PerpendicularVectors", &_perpendicularvectorgrid.PerpendicularVectors)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if perpendicularvectorFormCallback.formGroup.HasSuppressButtonBeenPressed {
		perpendicularvector_.Unstage(perpendicularvectorFormCallback.probe.stageOfInterest)
	}

	perpendicularvectorFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.PerpendicularVector](
		perpendicularvectorFormCallback.probe,
	)

	// display a new form by reset the form stage
	if perpendicularvectorFormCallback.CreationMode || perpendicularvectorFormCallback.formGroup.HasSuppressButtonBeenPressed {
		perpendicularvectorFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(perpendicularvectorFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PerpendicularVectorFormCallback(
			nil,
			perpendicularvectorFormCallback.probe,
			newFormGroup,
		)
		perpendicularvector := new(models.PerpendicularVector)
		FillUpForm(perpendicularvector, newFormGroup, perpendicularvectorFormCallback.probe)
		perpendicularvectorFormCallback.probe.formStage.Commit()
	}

	perpendicularvectorFormCallback.probe.ux_tree()
}
func __gong__New__PerpendicularVectorGridFormCallback(
	perpendicularvectorgrid *models.PerpendicularVectorGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (perpendicularvectorgridFormCallback *PerpendicularVectorGridFormCallback) {
	perpendicularvectorgridFormCallback = new(PerpendicularVectorGridFormCallback)
	perpendicularvectorgridFormCallback.probe = probe
	perpendicularvectorgridFormCallback.perpendicularvectorgrid = perpendicularvectorgrid
	perpendicularvectorgridFormCallback.formGroup = formGroup

	perpendicularvectorgridFormCallback.CreationMode = (perpendicularvectorgrid == nil)

	return
}

type PerpendicularVectorGridFormCallback struct {
	perpendicularvectorgrid *models.PerpendicularVectorGrid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (perpendicularvectorgridFormCallback *PerpendicularVectorGridFormCallback) OnSave() {
	perpendicularvectorgridFormCallback.probe.stageOfInterest.Lock()
	defer perpendicularvectorgridFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PerpendicularVectorGridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	perpendicularvectorgridFormCallback.probe.formStage.Checkout()

	if perpendicularvectorgridFormCallback.perpendicularvectorgrid == nil {
		perpendicularvectorgridFormCallback.perpendicularvectorgrid = new(models.PerpendicularVectorGrid).Stage(perpendicularvectorgridFormCallback.probe.stageOfInterest)
	}
	perpendicularvectorgrid_ := perpendicularvectorgridFormCallback.perpendicularvectorgrid
	_ = perpendicularvectorgrid_

	for _, formDiv := range perpendicularvectorgridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(perpendicularvectorgrid_.Name), formDiv)
		case "PerpendicularVectors":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.PerpendicularVector](perpendicularvectorgridFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.PerpendicularVector, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.PerpendicularVector)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					perpendicularvectorgridFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.PerpendicularVector](perpendicularvectorgridFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			perpendicularvectorgrid_.PerpendicularVectors = instanceSlice
			perpendicularvectorgridFormCallback.probe.UpdateSliceOfPointersCallback(perpendicularvectorgrid_, "PerpendicularVectors", &perpendicularvectorgrid_.PerpendicularVectors)

		}
	}

	// manage the suppress operation
	if perpendicularvectorgridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		perpendicularvectorgrid_.Unstage(perpendicularvectorgridFormCallback.probe.stageOfInterest)
	}

	perpendicularvectorgridFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.PerpendicularVectorGrid](
		perpendicularvectorgridFormCallback.probe,
	)

	// display a new form by reset the form stage
	if perpendicularvectorgridFormCallback.CreationMode || perpendicularvectorgridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		perpendicularvectorgridFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(perpendicularvectorgridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PerpendicularVectorGridFormCallback(
			nil,
			perpendicularvectorgridFormCallback.probe,
			newFormGroup,
		)
		perpendicularvectorgrid := new(models.PerpendicularVectorGrid)
		FillUpForm(perpendicularvectorgrid, newFormGroup, perpendicularvectorgridFormCallback.probe)
		perpendicularvectorgridFormCallback.probe.formStage.Commit()
	}

	perpendicularvectorgridFormCallback.probe.ux_tree()
}
func __gong__New__PerpendicularVectorGridHalfwayFormCallback(
	perpendicularvectorgridhalfway *models.PerpendicularVectorGridHalfway,
	probe *Probe,
	formGroup *form.FormGroup,
) (perpendicularvectorgridhalfwayFormCallback *PerpendicularVectorGridHalfwayFormCallback) {
	perpendicularvectorgridhalfwayFormCallback = new(PerpendicularVectorGridHalfwayFormCallback)
	perpendicularvectorgridhalfwayFormCallback.probe = probe
	perpendicularvectorgridhalfwayFormCallback.perpendicularvectorgridhalfway = perpendicularvectorgridhalfway
	perpendicularvectorgridhalfwayFormCallback.formGroup = formGroup

	perpendicularvectorgridhalfwayFormCallback.CreationMode = (perpendicularvectorgridhalfway == nil)

	return
}

type PerpendicularVectorGridHalfwayFormCallback struct {
	perpendicularvectorgridhalfway *models.PerpendicularVectorGridHalfway

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (perpendicularvectorgridhalfwayFormCallback *PerpendicularVectorGridHalfwayFormCallback) OnSave() {
	perpendicularvectorgridhalfwayFormCallback.probe.stageOfInterest.Lock()
	defer perpendicularvectorgridhalfwayFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PerpendicularVectorGridHalfwayFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	perpendicularvectorgridhalfwayFormCallback.probe.formStage.Checkout()

	if perpendicularvectorgridhalfwayFormCallback.perpendicularvectorgridhalfway == nil {
		perpendicularvectorgridhalfwayFormCallback.perpendicularvectorgridhalfway = new(models.PerpendicularVectorGridHalfway).Stage(perpendicularvectorgridhalfwayFormCallback.probe.stageOfInterest)
	}
	perpendicularvectorgridhalfway_ := perpendicularvectorgridhalfwayFormCallback.perpendicularvectorgridhalfway
	_ = perpendicularvectorgridhalfway_

	for _, formDiv := range perpendicularvectorgridhalfwayFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(perpendicularvectorgridhalfway_.Name), formDiv)
		case "PerpendicularVectorHalfways":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.PerpendicularVectorHalfway](perpendicularvectorgridhalfwayFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.PerpendicularVectorHalfway, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.PerpendicularVectorHalfway)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					perpendicularvectorgridhalfwayFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.PerpendicularVectorHalfway](perpendicularvectorgridhalfwayFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			perpendicularvectorgridhalfway_.PerpendicularVectorHalfways = instanceSlice
			perpendicularvectorgridhalfwayFormCallback.probe.UpdateSliceOfPointersCallback(perpendicularvectorgridhalfway_, "PerpendicularVectorHalfways", &perpendicularvectorgridhalfway_.PerpendicularVectorHalfways)

		}
	}

	// manage the suppress operation
	if perpendicularvectorgridhalfwayFormCallback.formGroup.HasSuppressButtonBeenPressed {
		perpendicularvectorgridhalfway_.Unstage(perpendicularvectorgridhalfwayFormCallback.probe.stageOfInterest)
	}

	perpendicularvectorgridhalfwayFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.PerpendicularVectorGridHalfway](
		perpendicularvectorgridhalfwayFormCallback.probe,
	)

	// display a new form by reset the form stage
	if perpendicularvectorgridhalfwayFormCallback.CreationMode || perpendicularvectorgridhalfwayFormCallback.formGroup.HasSuppressButtonBeenPressed {
		perpendicularvectorgridhalfwayFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(perpendicularvectorgridhalfwayFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PerpendicularVectorGridHalfwayFormCallback(
			nil,
			perpendicularvectorgridhalfwayFormCallback.probe,
			newFormGroup,
		)
		perpendicularvectorgridhalfway := new(models.PerpendicularVectorGridHalfway)
		FillUpForm(perpendicularvectorgridhalfway, newFormGroup, perpendicularvectorgridhalfwayFormCallback.probe)
		perpendicularvectorgridhalfwayFormCallback.probe.formStage.Commit()
	}

	perpendicularvectorgridhalfwayFormCallback.probe.ux_tree()
}
func __gong__New__PerpendicularVectorHalfwayFormCallback(
	perpendicularvectorhalfway *models.PerpendicularVectorHalfway,
	probe *Probe,
	formGroup *form.FormGroup,
) (perpendicularvectorhalfwayFormCallback *PerpendicularVectorHalfwayFormCallback) {
	perpendicularvectorhalfwayFormCallback = new(PerpendicularVectorHalfwayFormCallback)
	perpendicularvectorhalfwayFormCallback.probe = probe
	perpendicularvectorhalfwayFormCallback.perpendicularvectorhalfway = perpendicularvectorhalfway
	perpendicularvectorhalfwayFormCallback.formGroup = formGroup

	perpendicularvectorhalfwayFormCallback.CreationMode = (perpendicularvectorhalfway == nil)

	return
}

type PerpendicularVectorHalfwayFormCallback struct {
	perpendicularvectorhalfway *models.PerpendicularVectorHalfway

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (perpendicularvectorhalfwayFormCallback *PerpendicularVectorHalfwayFormCallback) OnSave() {
	perpendicularvectorhalfwayFormCallback.probe.stageOfInterest.Lock()
	defer perpendicularvectorhalfwayFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PerpendicularVectorHalfwayFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	perpendicularvectorhalfwayFormCallback.probe.formStage.Checkout()

	if perpendicularvectorhalfwayFormCallback.perpendicularvectorhalfway == nil {
		perpendicularvectorhalfwayFormCallback.perpendicularvectorhalfway = new(models.PerpendicularVectorHalfway).Stage(perpendicularvectorhalfwayFormCallback.probe.stageOfInterest)
	}
	perpendicularvectorhalfway_ := perpendicularvectorhalfwayFormCallback.perpendicularvectorhalfway
	_ = perpendicularvectorhalfway_

	for _, formDiv := range perpendicularvectorhalfwayFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(perpendicularvectorhalfway_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(perpendicularvectorhalfway_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(perpendicularvectorhalfway_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(perpendicularvectorhalfway_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(perpendicularvectorhalfway_.EndY), formDiv)
		case "PerpendicularVectorGridHalfway:PerpendicularVectorHalfways":
			// 1. Decode the AssociationStorage which contains the rowIDs of the PerpendicularVectorGridHalfway instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target PerpendicularVectorGridHalfway instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.PerpendicularVectorGridHalfway](perpendicularvectorhalfwayFormCallback.probe.stageOfInterest)
			targetPerpendicularVectorGridHalfwayIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetPerpendicularVectorGridHalfwayIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all PerpendicularVectorGridHalfway instances and update their PerpendicularVectorHalfways slice
			for _perpendicularvectorgridhalfway := range *models.GetGongstructInstancesSetFromPointerType[*models.PerpendicularVectorGridHalfway](perpendicularvectorhalfwayFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(perpendicularvectorhalfwayFormCallback.probe.stageOfInterest, _perpendicularvectorgridhalfway)
				
				// if PerpendicularVectorGridHalfway is selected
				if targetPerpendicularVectorGridHalfwayIDs[id] {
					// ensure perpendicularvectorhalfway_ is in _perpendicularvectorgridhalfway.PerpendicularVectorHalfways
					found := false
					for _, _b := range _perpendicularvectorgridhalfway.PerpendicularVectorHalfways {
						if _b == perpendicularvectorhalfway_ {
							found = true
							break
						}
					}
					if !found {
						_perpendicularvectorgridhalfway.PerpendicularVectorHalfways = append(_perpendicularvectorgridhalfway.PerpendicularVectorHalfways, perpendicularvectorhalfway_)
						perpendicularvectorhalfwayFormCallback.probe.UpdateSliceOfPointersCallback(_perpendicularvectorgridhalfway, "PerpendicularVectorHalfways", &_perpendicularvectorgridhalfway.PerpendicularVectorHalfways)
					}
				} else {
					// ensure perpendicularvectorhalfway_ is NOT in _perpendicularvectorgridhalfway.PerpendicularVectorHalfways
					idx := slices.Index(_perpendicularvectorgridhalfway.PerpendicularVectorHalfways, perpendicularvectorhalfway_)
					if idx != -1 {
						_perpendicularvectorgridhalfway.PerpendicularVectorHalfways = slices.Delete(_perpendicularvectorgridhalfway.PerpendicularVectorHalfways, idx, idx+1)
						perpendicularvectorhalfwayFormCallback.probe.UpdateSliceOfPointersCallback(_perpendicularvectorgridhalfway, "PerpendicularVectorHalfways", &_perpendicularvectorgridhalfway.PerpendicularVectorHalfways)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if perpendicularvectorhalfwayFormCallback.formGroup.HasSuppressButtonBeenPressed {
		perpendicularvectorhalfway_.Unstage(perpendicularvectorhalfwayFormCallback.probe.stageOfInterest)
	}

	perpendicularvectorhalfwayFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.PerpendicularVectorHalfway](
		perpendicularvectorhalfwayFormCallback.probe,
	)

	// display a new form by reset the form stage
	if perpendicularvectorhalfwayFormCallback.CreationMode || perpendicularvectorhalfwayFormCallback.formGroup.HasSuppressButtonBeenPressed {
		perpendicularvectorhalfwayFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(perpendicularvectorhalfwayFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PerpendicularVectorHalfwayFormCallback(
			nil,
			perpendicularvectorhalfwayFormCallback.probe,
			newFormGroup,
		)
		perpendicularvectorhalfway := new(models.PerpendicularVectorHalfway)
		FillUpForm(perpendicularvectorhalfway, newFormGroup, perpendicularvectorhalfwayFormCallback.probe)
		perpendicularvectorhalfwayFormCallback.probe.formStage.Commit()
	}

	perpendicularvectorhalfwayFormCallback.probe.ux_tree()
}
func __gong__New__PlantFormCallback(
	plant *models.Plant,
	probe *Probe,
	formGroup *form.FormGroup,
) (plantFormCallback *PlantFormCallback) {
	plantFormCallback = new(PlantFormCallback)
	plantFormCallback.probe = probe
	plantFormCallback.plant = plant
	plantFormCallback.formGroup = formGroup

	plantFormCallback.CreationMode = (plant == nil)

	return
}

type PlantFormCallback struct {
	plant *models.Plant

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (plantFormCallback *PlantFormCallback) OnSave() {
	plantFormCallback.probe.stageOfInterest.Lock()
	defer plantFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PlantFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	plantFormCallback.probe.formStage.Checkout()

	if plantFormCallback.plant == nil {
		plantFormCallback.plant = new(models.Plant).Stage(plantFormCallback.probe.stageOfInterest)
	}
	plant_ := plantFormCallback.plant
	_ = plant_

	for _, formDiv := range plantFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(plant_.Name), formDiv)
		case "N":
			FormDivBasicFieldToField(&(plant_.N), formDiv)
		case "M":
			FormDivBasicFieldToField(&(plant_.M), formDiv)
		case "StackHeight":
			FormDivBasicFieldToField(&(plant_.StackHeight), formDiv)
		case "RhombusInsideAngle":
			FormDivBasicFieldToField(&(plant_.RhombusInsideAngle), formDiv)
		case "Thickness":
			FormDivBasicFieldToField(&(plant_.Thickness), formDiv)
		case "RhombusSideLength":
			FormDivBasicFieldToField(&(plant_.RhombusSideLength), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(plant_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(plant_.IsExpanded), formDiv)
		case "IsSelected":
			FormDivBasicFieldToField(&(plant_.IsSelected), formDiv)
		case "IsPlantDiagramsNodeExpanded":
			FormDivBasicFieldToField(&(plant_.IsPlantDiagramsNodeExpanded), formDiv)
		case "PlantDiagrams":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.PlantDiagram](plantFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.PlantDiagram, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.PlantDiagram)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					plantFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.PlantDiagram](plantFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			plant_.PlantDiagrams = instanceSlice
			plantFormCallback.probe.UpdateSliceOfPointersCallback(plant_, "PlantDiagrams", &plant_.PlantDiagrams)

		case "AxesShape":
			FormDivSelectFieldToField(&(plant_.AxesShape), plantFormCallback.probe.stageOfInterest, formDiv)
		case "ReferenceRhombus":
			FormDivSelectFieldToField(&(plant_.ReferenceRhombus), plantFormCallback.probe.stageOfInterest, formDiv)
		case "PlantCircumferenceShape":
			FormDivSelectFieldToField(&(plant_.PlantCircumferenceShape), plantFormCallback.probe.stageOfInterest, formDiv)
		case "GridPathShape":
			FormDivSelectFieldToField(&(plant_.GridPathShape), plantFormCallback.probe.stageOfInterest, formDiv)
		case "InitialRhombusGridShape":
			FormDivSelectFieldToField(&(plant_.InitialRhombusGridShape), plantFormCallback.probe.stageOfInterest, formDiv)
		case "ExplanationTextShape":
			FormDivSelectFieldToField(&(plant_.ExplanationTextShape), plantFormCallback.probe.stageOfInterest, formDiv)
		case "RotatedReferenceRhombus":
			FormDivSelectFieldToField(&(plant_.RotatedReferenceRhombus), plantFormCallback.probe.stageOfInterest, formDiv)
		case "RotatedPlantCircumferenceShape":
			FormDivSelectFieldToField(&(plant_.RotatedPlantCircumferenceShape), plantFormCallback.probe.stageOfInterest, formDiv)
		case "RotatedGridPathShape":
			FormDivSelectFieldToField(&(plant_.RotatedGridPathShape), plantFormCallback.probe.stageOfInterest, formDiv)
		case "RotatedRhombusGridShape2":
			FormDivSelectFieldToField(&(plant_.RotatedRhombusGridShape2), plantFormCallback.probe.stageOfInterest, formDiv)
		case "GrowthCurveRhombusGridShape":
			FormDivSelectFieldToField(&(plant_.GrowthCurveRhombusGridShape), plantFormCallback.probe.stageOfInterest, formDiv)
		case "GrowthVectorShape":
			FormDivSelectFieldToField(&(plant_.GrowthVectorShape), plantFormCallback.probe.stageOfInterest, formDiv)
		case "PerpendicularVectorGrid":
			FormDivSelectFieldToField(&(plant_.PerpendicularVectorGrid), plantFormCallback.probe.stageOfInterest, formDiv)
		case "PerpendicularVectorGridHalfway":
			FormDivSelectFieldToField(&(plant_.PerpendicularVectorGridHalfway), plantFormCallback.probe.stageOfInterest, formDiv)
		case "BaseVectorShapeGrid":
			FormDivSelectFieldToField(&(plant_.BaseVectorShapeGrid), plantFormCallback.probe.stageOfInterest, formDiv)
		case "ArcNormalVectorShapeGrid":
			FormDivSelectFieldToField(&(plant_.ArcNormalVectorShapeGrid), plantFormCallback.probe.stageOfInterest, formDiv)
		case "StartArcShapeGrid":
			FormDivSelectFieldToField(&(plant_.StartArcShapeGrid), plantFormCallback.probe.stageOfInterest, formDiv)
		case "TopStartArcShapeGrid":
			FormDivSelectFieldToField(&(plant_.TopStartArcShapeGrid), plantFormCallback.probe.stageOfInterest, formDiv)
		case "EndArcShapeGrid":
			FormDivSelectFieldToField(&(plant_.EndArcShapeGrid), plantFormCallback.probe.stageOfInterest, formDiv)
		case "TopEndArcShapeGrid":
			FormDivSelectFieldToField(&(plant_.TopEndArcShapeGrid), plantFormCallback.probe.stageOfInterest, formDiv)
		case "BottomStartArcShapeGrid":
			FormDivSelectFieldToField(&(plant_.BottomStartArcShapeGrid), plantFormCallback.probe.stageOfInterest, formDiv)
		case "BottomEndArcShapeGrid":
			FormDivSelectFieldToField(&(plant_.BottomEndArcShapeGrid), plantFormCallback.probe.stageOfInterest, formDiv)
		case "GrowthCurveBezierShapeGrid":
			FormDivSelectFieldToField(&(plant_.GrowthCurveBezierShapeGrid), plantFormCallback.probe.stageOfInterest, formDiv)
		case "StackOfGrowthCurve":
			FormDivSelectFieldToField(&(plant_.StackOfGrowthCurve), plantFormCallback.probe.stageOfInterest, formDiv)
		case "TopStackOfGrowthCurve":
			FormDivSelectFieldToField(&(plant_.TopStackOfGrowthCurve), plantFormCallback.probe.stageOfInterest, formDiv)
		case "BottomStackOfGrowthCurve":
			FormDivSelectFieldToField(&(plant_.BottomStackOfGrowthCurve), plantFormCallback.probe.stageOfInterest, formDiv)
		case "GrowthCurve2D":
			FormDivSelectFieldToField(&(plant_.GrowthCurve2D), plantFormCallback.probe.stageOfInterest, formDiv)
		case "TopGrowthCurve2D":
			FormDivSelectFieldToField(&(plant_.TopGrowthCurve2D), plantFormCallback.probe.stageOfInterest, formDiv)
		case "Library:Plants":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](plantFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their Plants slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](plantFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(plantFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure plant_ is in _library.Plants
					found := false
					for _, _b := range _library.Plants {
						if _b == plant_ {
							found = true
							break
						}
					}
					if !found {
						_library.Plants = append(_library.Plants, plant_)
						plantFormCallback.probe.UpdateSliceOfPointersCallback(_library, "Plants", &_library.Plants)
					}
				} else {
					// ensure plant_ is NOT in _library.Plants
					idx := slices.Index(_library.Plants, plant_)
					if idx != -1 {
						_library.Plants = slices.Delete(_library.Plants, idx, idx+1)
						plantFormCallback.probe.UpdateSliceOfPointersCallback(_library, "Plants", &_library.Plants)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if plantFormCallback.formGroup.HasSuppressButtonBeenPressed {
		plant_.Unstage(plantFormCallback.probe.stageOfInterest)
	}

	plantFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Plant](
		plantFormCallback.probe,
	)

	// display a new form by reset the form stage
	if plantFormCallback.CreationMode || plantFormCallback.formGroup.HasSuppressButtonBeenPressed {
		plantFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(plantFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PlantFormCallback(
			nil,
			plantFormCallback.probe,
			newFormGroup,
		)
		plant := new(models.Plant)
		FillUpForm(plant, newFormGroup, plantFormCallback.probe)
		plantFormCallback.probe.formStage.Commit()
	}

	plantFormCallback.probe.ux_tree()
}
func __gong__New__PlantCircumferenceShapeFormCallback(
	plantcircumferenceshape *models.PlantCircumferenceShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (plantcircumferenceshapeFormCallback *PlantCircumferenceShapeFormCallback) {
	plantcircumferenceshapeFormCallback = new(PlantCircumferenceShapeFormCallback)
	plantcircumferenceshapeFormCallback.probe = probe
	plantcircumferenceshapeFormCallback.plantcircumferenceshape = plantcircumferenceshape
	plantcircumferenceshapeFormCallback.formGroup = formGroup

	plantcircumferenceshapeFormCallback.CreationMode = (plantcircumferenceshape == nil)

	return
}

type PlantCircumferenceShapeFormCallback struct {
	plantcircumferenceshape *models.PlantCircumferenceShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (plantcircumferenceshapeFormCallback *PlantCircumferenceShapeFormCallback) OnSave() {
	plantcircumferenceshapeFormCallback.probe.stageOfInterest.Lock()
	defer plantcircumferenceshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PlantCircumferenceShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	plantcircumferenceshapeFormCallback.probe.formStage.Checkout()

	if plantcircumferenceshapeFormCallback.plantcircumferenceshape == nil {
		plantcircumferenceshapeFormCallback.plantcircumferenceshape = new(models.PlantCircumferenceShape).Stage(plantcircumferenceshapeFormCallback.probe.stageOfInterest)
	}
	plantcircumferenceshape_ := plantcircumferenceshapeFormCallback.plantcircumferenceshape
	_ = plantcircumferenceshape_

	for _, formDiv := range plantcircumferenceshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(plantcircumferenceshape_.Name), formDiv)
		case "AngleDegree":
			FormDivBasicFieldToField(&(plantcircumferenceshape_.AngleDegree), formDiv)
		case "Length":
			FormDivBasicFieldToField(&(plantcircumferenceshape_.Length), formDiv)
		}
	}

	// manage the suppress operation
	if plantcircumferenceshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		plantcircumferenceshape_.Unstage(plantcircumferenceshapeFormCallback.probe.stageOfInterest)
	}

	plantcircumferenceshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.PlantCircumferenceShape](
		plantcircumferenceshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if plantcircumferenceshapeFormCallback.CreationMode || plantcircumferenceshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		plantcircumferenceshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(plantcircumferenceshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PlantCircumferenceShapeFormCallback(
			nil,
			plantcircumferenceshapeFormCallback.probe,
			newFormGroup,
		)
		plantcircumferenceshape := new(models.PlantCircumferenceShape)
		FillUpForm(plantcircumferenceshape, newFormGroup, plantcircumferenceshapeFormCallback.probe)
		plantcircumferenceshapeFormCallback.probe.formStage.Commit()
	}

	plantcircumferenceshapeFormCallback.probe.ux_tree()
}
func __gong__New__PlantDiagramFormCallback(
	plantdiagram *models.PlantDiagram,
	probe *Probe,
	formGroup *form.FormGroup,
) (plantdiagramFormCallback *PlantDiagramFormCallback) {
	plantdiagramFormCallback = new(PlantDiagramFormCallback)
	plantdiagramFormCallback.probe = probe
	plantdiagramFormCallback.plantdiagram = plantdiagram
	plantdiagramFormCallback.formGroup = formGroup

	plantdiagramFormCallback.CreationMode = (plantdiagram == nil)

	return
}

type PlantDiagramFormCallback struct {
	plantdiagram *models.PlantDiagram

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (plantdiagramFormCallback *PlantDiagramFormCallback) OnSave() {
	plantdiagramFormCallback.probe.stageOfInterest.Lock()
	defer plantdiagramFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PlantDiagramFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	plantdiagramFormCallback.probe.formStage.Checkout()

	if plantdiagramFormCallback.plantdiagram == nil {
		plantdiagramFormCallback.plantdiagram = new(models.PlantDiagram).Stage(plantdiagramFormCallback.probe.stageOfInterest)
	}
	plantdiagram_ := plantdiagramFormCallback.plantdiagram
	_ = plantdiagram_

	for _, formDiv := range plantdiagramFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(plantdiagram_.Name), formDiv)
		case "OriginX":
			FormDivBasicFieldToField(&(plantdiagram_.OriginX), formDiv)
		case "OriginY":
			FormDivBasicFieldToField(&(plantdiagram_.OriginY), formDiv)
		case "IsRhombusNodesExpanded":
			FormDivBasicFieldToField(&(plantdiagram_.IsRhombusNodesExpanded), formDiv)
		case "IsHiddenAxesShape":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenAxesShape), formDiv)
		case "IsHiddenReferenceRhombus":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenReferenceRhombus), formDiv)
		case "IsHiddenPlantCircumferenceShape":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenPlantCircumferenceShape), formDiv)
		case "IsHiddenGridPathShape":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenGridPathShape), formDiv)
		case "IsHiddenRhombusGridShape":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenRhombusGridShape), formDiv)
		case "IsHiddenExplanationTextShape":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenExplanationTextShape), formDiv)
		case "IsHiddenRotatedReferenceRhombus":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenRotatedReferenceRhombus), formDiv)
		case "IsHiddenRotatedPlantCircumferenceShape":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenRotatedPlantCircumferenceShape), formDiv)
		case "IsHiddenRotatedGridPathShape":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenRotatedGridPathShape), formDiv)
		case "IsHiddenRotatedRhombusGridShape":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenRotatedRhombusGridShape), formDiv)
		case "IsHiddenGrowthPathRhombusGridShape":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenGrowthPathRhombusGridShape), formDiv)
		case "IsHiddenGrowthVectorShape":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenGrowthVectorShape), formDiv)
		case "IsHiddenPerpendicularVectorGrid":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenPerpendicularVectorGrid), formDiv)
		case "IsHiddenPerpendicularVectorGridHalfway":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenPerpendicularVectorGridHalfway), formDiv)
		case "IsHiddenBaseVectorShapeGrid":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenBaseVectorShapeGrid), formDiv)
		case "IsHiddenArcNormalVectorShapeGrid":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenArcNormalVectorShapeGrid), formDiv)
		case "IsHiddenStartArcShapeGrid":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenStartArcShapeGrid), formDiv)
		case "IsHiddenTopStartArcShapeGrid":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenTopStartArcShapeGrid), formDiv)
		case "IsHiddenEndArcShapeGrid":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenEndArcShapeGrid), formDiv)
		case "IsHiddenTopEndArcShapeGrid":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenTopEndArcShapeGrid), formDiv)
		case "IsHiddenBottomStartArcShapeGrid":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenBottomStartArcShapeGrid), formDiv)
		case "IsHiddenBottomEndArcShapeGrid":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenBottomEndArcShapeGrid), formDiv)
		case "IsHiddenGrowthCurveBezierShapeGrid":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenGrowthCurveBezierShapeGrid), formDiv)
		case "IsHiddenStackOfGrowthCurve":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenStackOfGrowthCurve), formDiv)
		case "IsHiddenTopStackOfGrowthCurve":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenTopStackOfGrowthCurve), formDiv)
		case "IsHiddenBottomStackOfGrowthCurve":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenBottomStackOfGrowthCurve), formDiv)
		case "IsHiddenGrowthCurve2D":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenGrowthCurve2D), formDiv)
		case "IsHiddenTopGrowthCurve2D":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenTopGrowthCurve2D), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(plantdiagram_.IsChecked), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(plantdiagram_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(plantdiagram_.IsExpanded), formDiv)
		case "Rendered3DShape":
			FormDivSelectFieldToField(&(plantdiagram_.Rendered3DShape), plantdiagramFormCallback.probe.stageOfInterest, formDiv)
		case "Plant:PlantDiagrams":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Plant instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Plant instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Plant](plantdiagramFormCallback.probe.stageOfInterest)
			targetPlantIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetPlantIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Plant instances and update their PlantDiagrams slice
			for _plant := range *models.GetGongstructInstancesSetFromPointerType[*models.Plant](plantdiagramFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(plantdiagramFormCallback.probe.stageOfInterest, _plant)
				
				// if Plant is selected
				if targetPlantIDs[id] {
					// ensure plantdiagram_ is in _plant.PlantDiagrams
					found := false
					for _, _b := range _plant.PlantDiagrams {
						if _b == plantdiagram_ {
							found = true
							break
						}
					}
					if !found {
						_plant.PlantDiagrams = append(_plant.PlantDiagrams, plantdiagram_)
						plantdiagramFormCallback.probe.UpdateSliceOfPointersCallback(_plant, "PlantDiagrams", &_plant.PlantDiagrams)
					}
				} else {
					// ensure plantdiagram_ is NOT in _plant.PlantDiagrams
					idx := slices.Index(_plant.PlantDiagrams, plantdiagram_)
					if idx != -1 {
						_plant.PlantDiagrams = slices.Delete(_plant.PlantDiagrams, idx, idx+1)
						plantdiagramFormCallback.probe.UpdateSliceOfPointersCallback(_plant, "PlantDiagrams", &_plant.PlantDiagrams)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if plantdiagramFormCallback.formGroup.HasSuppressButtonBeenPressed {
		plantdiagram_.Unstage(plantdiagramFormCallback.probe.stageOfInterest)
	}

	plantdiagramFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.PlantDiagram](
		plantdiagramFormCallback.probe,
	)

	// display a new form by reset the form stage
	if plantdiagramFormCallback.CreationMode || plantdiagramFormCallback.formGroup.HasSuppressButtonBeenPressed {
		plantdiagramFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(plantdiagramFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PlantDiagramFormCallback(
			nil,
			plantdiagramFormCallback.probe,
			newFormGroup,
		)
		plantdiagram := new(models.PlantDiagram)
		FillUpForm(plantdiagram, newFormGroup, plantdiagramFormCallback.probe)
		plantdiagramFormCallback.probe.formStage.Commit()
	}

	plantdiagramFormCallback.probe.ux_tree()
}
func __gong__New__Rendered3DShapeFormCallback(
	rendered3dshape *models.Rendered3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (rendered3dshapeFormCallback *Rendered3DShapeFormCallback) {
	rendered3dshapeFormCallback = new(Rendered3DShapeFormCallback)
	rendered3dshapeFormCallback.probe = probe
	rendered3dshapeFormCallback.rendered3dshape = rendered3dshape
	rendered3dshapeFormCallback.formGroup = formGroup

	rendered3dshapeFormCallback.CreationMode = (rendered3dshape == nil)

	return
}

type Rendered3DShapeFormCallback struct {
	rendered3dshape *models.Rendered3DShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (rendered3dshapeFormCallback *Rendered3DShapeFormCallback) OnSave() {
	rendered3dshapeFormCallback.probe.stageOfInterest.Lock()
	defer rendered3dshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("Rendered3DShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rendered3dshapeFormCallback.probe.formStage.Checkout()

	if rendered3dshapeFormCallback.rendered3dshape == nil {
		rendered3dshapeFormCallback.rendered3dshape = new(models.Rendered3DShape).Stage(rendered3dshapeFormCallback.probe.stageOfInterest)
	}
	rendered3dshape_ := rendered3dshapeFormCallback.rendered3dshape
	_ = rendered3dshape_

	for _, formDiv := range rendered3dshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rendered3dshape_.Name), formDiv)
		case "ViewX":
			FormDivBasicFieldToField(&(rendered3dshape_.ViewX), formDiv)
		case "ViewY":
			FormDivBasicFieldToField(&(rendered3dshape_.ViewY), formDiv)
		case "ViewZ":
			FormDivBasicFieldToField(&(rendered3dshape_.ViewZ), formDiv)
		case "TargetX":
			FormDivBasicFieldToField(&(rendered3dshape_.TargetX), formDiv)
		case "TargetY":
			FormDivBasicFieldToField(&(rendered3dshape_.TargetY), formDiv)
		case "TargetZ":
			FormDivBasicFieldToField(&(rendered3dshape_.TargetZ), formDiv)
		case "Fov":
			FormDivBasicFieldToField(&(rendered3dshape_.Fov), formDiv)
		}
	}

	// manage the suppress operation
	if rendered3dshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rendered3dshape_.Unstage(rendered3dshapeFormCallback.probe.stageOfInterest)
	}

	rendered3dshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Rendered3DShape](
		rendered3dshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if rendered3dshapeFormCallback.CreationMode || rendered3dshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rendered3dshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(rendered3dshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__Rendered3DShapeFormCallback(
			nil,
			rendered3dshapeFormCallback.probe,
			newFormGroup,
		)
		rendered3dshape := new(models.Rendered3DShape)
		FillUpForm(rendered3dshape, newFormGroup, rendered3dshapeFormCallback.probe)
		rendered3dshapeFormCallback.probe.formStage.Commit()
	}

	rendered3dshapeFormCallback.probe.ux_tree()
}
func __gong__New__RhombusShapeFormCallback(
	rhombusshape *models.RhombusShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (rhombusshapeFormCallback *RhombusShapeFormCallback) {
	rhombusshapeFormCallback = new(RhombusShapeFormCallback)
	rhombusshapeFormCallback.probe = probe
	rhombusshapeFormCallback.rhombusshape = rhombusshape
	rhombusshapeFormCallback.formGroup = formGroup

	rhombusshapeFormCallback.CreationMode = (rhombusshape == nil)

	return
}

type RhombusShapeFormCallback struct {
	rhombusshape *models.RhombusShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (rhombusshapeFormCallback *RhombusShapeFormCallback) OnSave() {
	rhombusshapeFormCallback.probe.stageOfInterest.Lock()
	defer rhombusshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("RhombusShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rhombusshapeFormCallback.probe.formStage.Checkout()

	if rhombusshapeFormCallback.rhombusshape == nil {
		rhombusshapeFormCallback.rhombusshape = new(models.RhombusShape).Stage(rhombusshapeFormCallback.probe.stageOfInterest)
	}
	rhombusshape_ := rhombusshapeFormCallback.rhombusshape
	_ = rhombusshape_

	for _, formDiv := range rhombusshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rhombusshape_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(rhombusshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(rhombusshape_.Y), formDiv)
		}
	}

	// manage the suppress operation
	if rhombusshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rhombusshape_.Unstage(rhombusshapeFormCallback.probe.stageOfInterest)
	}

	rhombusshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.RhombusShape](
		rhombusshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if rhombusshapeFormCallback.CreationMode || rhombusshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rhombusshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(rhombusshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RhombusShapeFormCallback(
			nil,
			rhombusshapeFormCallback.probe,
			newFormGroup,
		)
		rhombusshape := new(models.RhombusShape)
		FillUpForm(rhombusshape, newFormGroup, rhombusshapeFormCallback.probe)
		rhombusshapeFormCallback.probe.formStage.Commit()
	}

	rhombusshapeFormCallback.probe.ux_tree()
}
func __gong__New__RotatedRhombusGridShapeFormCallback(
	rotatedrhombusgridshape *models.RotatedRhombusGridShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (rotatedrhombusgridshapeFormCallback *RotatedRhombusGridShapeFormCallback) {
	rotatedrhombusgridshapeFormCallback = new(RotatedRhombusGridShapeFormCallback)
	rotatedrhombusgridshapeFormCallback.probe = probe
	rotatedrhombusgridshapeFormCallback.rotatedrhombusgridshape = rotatedrhombusgridshape
	rotatedrhombusgridshapeFormCallback.formGroup = formGroup

	rotatedrhombusgridshapeFormCallback.CreationMode = (rotatedrhombusgridshape == nil)

	return
}

type RotatedRhombusGridShapeFormCallback struct {
	rotatedrhombusgridshape *models.RotatedRhombusGridShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (rotatedrhombusgridshapeFormCallback *RotatedRhombusGridShapeFormCallback) OnSave() {
	rotatedrhombusgridshapeFormCallback.probe.stageOfInterest.Lock()
	defer rotatedrhombusgridshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("RotatedRhombusGridShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rotatedrhombusgridshapeFormCallback.probe.formStage.Checkout()

	if rotatedrhombusgridshapeFormCallback.rotatedrhombusgridshape == nil {
		rotatedrhombusgridshapeFormCallback.rotatedrhombusgridshape = new(models.RotatedRhombusGridShape).Stage(rotatedrhombusgridshapeFormCallback.probe.stageOfInterest)
	}
	rotatedrhombusgridshape_ := rotatedrhombusgridshapeFormCallback.rotatedrhombusgridshape
	_ = rotatedrhombusgridshape_

	for _, formDiv := range rotatedrhombusgridshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rotatedrhombusgridshape_.Name), formDiv)
		case "RotatedRhombusShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.RotatedRhombusShape](rotatedrhombusgridshapeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.RotatedRhombusShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.RotatedRhombusShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					rotatedrhombusgridshapeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.RotatedRhombusShape](rotatedrhombusgridshapeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			rotatedrhombusgridshape_.RotatedRhombusShapes = instanceSlice
			rotatedrhombusgridshapeFormCallback.probe.UpdateSliceOfPointersCallback(rotatedrhombusgridshape_, "RotatedRhombusShapes", &rotatedrhombusgridshape_.RotatedRhombusShapes)

		}
	}

	// manage the suppress operation
	if rotatedrhombusgridshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rotatedrhombusgridshape_.Unstage(rotatedrhombusgridshapeFormCallback.probe.stageOfInterest)
	}

	rotatedrhombusgridshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.RotatedRhombusGridShape](
		rotatedrhombusgridshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if rotatedrhombusgridshapeFormCallback.CreationMode || rotatedrhombusgridshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rotatedrhombusgridshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(rotatedrhombusgridshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RotatedRhombusGridShapeFormCallback(
			nil,
			rotatedrhombusgridshapeFormCallback.probe,
			newFormGroup,
		)
		rotatedrhombusgridshape := new(models.RotatedRhombusGridShape)
		FillUpForm(rotatedrhombusgridshape, newFormGroup, rotatedrhombusgridshapeFormCallback.probe)
		rotatedrhombusgridshapeFormCallback.probe.formStage.Commit()
	}

	rotatedrhombusgridshapeFormCallback.probe.ux_tree()
}
func __gong__New__RotatedRhombusShapeFormCallback(
	rotatedrhombusshape *models.RotatedRhombusShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (rotatedrhombusshapeFormCallback *RotatedRhombusShapeFormCallback) {
	rotatedrhombusshapeFormCallback = new(RotatedRhombusShapeFormCallback)
	rotatedrhombusshapeFormCallback.probe = probe
	rotatedrhombusshapeFormCallback.rotatedrhombusshape = rotatedrhombusshape
	rotatedrhombusshapeFormCallback.formGroup = formGroup

	rotatedrhombusshapeFormCallback.CreationMode = (rotatedrhombusshape == nil)

	return
}

type RotatedRhombusShapeFormCallback struct {
	rotatedrhombusshape *models.RotatedRhombusShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (rotatedrhombusshapeFormCallback *RotatedRhombusShapeFormCallback) OnSave() {
	rotatedrhombusshapeFormCallback.probe.stageOfInterest.Lock()
	defer rotatedrhombusshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("RotatedRhombusShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rotatedrhombusshapeFormCallback.probe.formStage.Checkout()

	if rotatedrhombusshapeFormCallback.rotatedrhombusshape == nil {
		rotatedrhombusshapeFormCallback.rotatedrhombusshape = new(models.RotatedRhombusShape).Stage(rotatedrhombusshapeFormCallback.probe.stageOfInterest)
	}
	rotatedrhombusshape_ := rotatedrhombusshapeFormCallback.rotatedrhombusshape
	_ = rotatedrhombusshape_

	for _, formDiv := range rotatedrhombusshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rotatedrhombusshape_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(rotatedrhombusshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(rotatedrhombusshape_.Y), formDiv)
		case "RotatedRhombusGridShape:RotatedRhombusShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the RotatedRhombusGridShape instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target RotatedRhombusGridShape instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.RotatedRhombusGridShape](rotatedrhombusshapeFormCallback.probe.stageOfInterest)
			targetRotatedRhombusGridShapeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetRotatedRhombusGridShapeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all RotatedRhombusGridShape instances and update their RotatedRhombusShapes slice
			for _rotatedrhombusgridshape := range *models.GetGongstructInstancesSetFromPointerType[*models.RotatedRhombusGridShape](rotatedrhombusshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(rotatedrhombusshapeFormCallback.probe.stageOfInterest, _rotatedrhombusgridshape)
				
				// if RotatedRhombusGridShape is selected
				if targetRotatedRhombusGridShapeIDs[id] {
					// ensure rotatedrhombusshape_ is in _rotatedrhombusgridshape.RotatedRhombusShapes
					found := false
					for _, _b := range _rotatedrhombusgridshape.RotatedRhombusShapes {
						if _b == rotatedrhombusshape_ {
							found = true
							break
						}
					}
					if !found {
						_rotatedrhombusgridshape.RotatedRhombusShapes = append(_rotatedrhombusgridshape.RotatedRhombusShapes, rotatedrhombusshape_)
						rotatedrhombusshapeFormCallback.probe.UpdateSliceOfPointersCallback(_rotatedrhombusgridshape, "RotatedRhombusShapes", &_rotatedrhombusgridshape.RotatedRhombusShapes)
					}
				} else {
					// ensure rotatedrhombusshape_ is NOT in _rotatedrhombusgridshape.RotatedRhombusShapes
					idx := slices.Index(_rotatedrhombusgridshape.RotatedRhombusShapes, rotatedrhombusshape_)
					if idx != -1 {
						_rotatedrhombusgridshape.RotatedRhombusShapes = slices.Delete(_rotatedrhombusgridshape.RotatedRhombusShapes, idx, idx+1)
						rotatedrhombusshapeFormCallback.probe.UpdateSliceOfPointersCallback(_rotatedrhombusgridshape, "RotatedRhombusShapes", &_rotatedrhombusgridshape.RotatedRhombusShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if rotatedrhombusshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rotatedrhombusshape_.Unstage(rotatedrhombusshapeFormCallback.probe.stageOfInterest)
	}

	rotatedrhombusshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.RotatedRhombusShape](
		rotatedrhombusshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if rotatedrhombusshapeFormCallback.CreationMode || rotatedrhombusshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rotatedrhombusshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(rotatedrhombusshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RotatedRhombusShapeFormCallback(
			nil,
			rotatedrhombusshapeFormCallback.probe,
			newFormGroup,
		)
		rotatedrhombusshape := new(models.RotatedRhombusShape)
		FillUpForm(rotatedrhombusshape, newFormGroup, rotatedrhombusshapeFormCallback.probe)
		rotatedrhombusshapeFormCallback.probe.formStage.Commit()
	}

	rotatedrhombusshapeFormCallback.probe.ux_tree()
}
func __gong__New__StackGrowthCurveEndArcShapeFormCallback(
	stackgrowthcurveendarcshape *models.StackGrowthCurveEndArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackgrowthcurveendarcshapeFormCallback *StackGrowthCurveEndArcShapeFormCallback) {
	stackgrowthcurveendarcshapeFormCallback = new(StackGrowthCurveEndArcShapeFormCallback)
	stackgrowthcurveendarcshapeFormCallback.probe = probe
	stackgrowthcurveendarcshapeFormCallback.stackgrowthcurveendarcshape = stackgrowthcurveendarcshape
	stackgrowthcurveendarcshapeFormCallback.formGroup = formGroup

	stackgrowthcurveendarcshapeFormCallback.CreationMode = (stackgrowthcurveendarcshape == nil)

	return
}

type StackGrowthCurveEndArcShapeFormCallback struct {
	stackgrowthcurveendarcshape *models.StackGrowthCurveEndArcShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (stackgrowthcurveendarcshapeFormCallback *StackGrowthCurveEndArcShapeFormCallback) OnSave() {
	stackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest.Lock()
	defer stackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StackGrowthCurveEndArcShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	stackgrowthcurveendarcshapeFormCallback.probe.formStage.Checkout()

	if stackgrowthcurveendarcshapeFormCallback.stackgrowthcurveendarcshape == nil {
		stackgrowthcurveendarcshapeFormCallback.stackgrowthcurveendarcshape = new(models.StackGrowthCurveEndArcShape).Stage(stackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest)
	}
	stackgrowthcurveendarcshape_ := stackgrowthcurveendarcshapeFormCallback.stackgrowthcurveendarcshape
	_ = stackgrowthcurveendarcshape_

	for _, formDiv := range stackgrowthcurveendarcshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(stackgrowthcurveendarcshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(stackgrowthcurveendarcshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(stackgrowthcurveendarcshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(stackgrowthcurveendarcshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(stackgrowthcurveendarcshape_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(stackgrowthcurveendarcshape_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(stackgrowthcurveendarcshape_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(stackgrowthcurveendarcshape_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(stackgrowthcurveendarcshape_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(stackgrowthcurveendarcshape_.RadiusY), formDiv)
		case "StackOfGrowthCurve:StackGrowthCurveEndArcShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the StackOfGrowthCurve instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target StackOfGrowthCurve instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.StackOfGrowthCurve](stackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest)
			targetStackOfGrowthCurveIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetStackOfGrowthCurveIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all StackOfGrowthCurve instances and update their StackGrowthCurveEndArcShapes slice
			for _stackofgrowthcurve := range *models.GetGongstructInstancesSetFromPointerType[*models.StackOfGrowthCurve](stackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(stackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest, _stackofgrowthcurve)
				
				// if StackOfGrowthCurve is selected
				if targetStackOfGrowthCurveIDs[id] {
					// ensure stackgrowthcurveendarcshape_ is in _stackofgrowthcurve.StackGrowthCurveEndArcShapes
					found := false
					for _, _b := range _stackofgrowthcurve.StackGrowthCurveEndArcShapes {
						if _b == stackgrowthcurveendarcshape_ {
							found = true
							break
						}
					}
					if !found {
						_stackofgrowthcurve.StackGrowthCurveEndArcShapes = append(_stackofgrowthcurve.StackGrowthCurveEndArcShapes, stackgrowthcurveendarcshape_)
						stackgrowthcurveendarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_stackofgrowthcurve, "StackGrowthCurveEndArcShapes", &_stackofgrowthcurve.StackGrowthCurveEndArcShapes)
					}
				} else {
					// ensure stackgrowthcurveendarcshape_ is NOT in _stackofgrowthcurve.StackGrowthCurveEndArcShapes
					idx := slices.Index(_stackofgrowthcurve.StackGrowthCurveEndArcShapes, stackgrowthcurveendarcshape_)
					if idx != -1 {
						_stackofgrowthcurve.StackGrowthCurveEndArcShapes = slices.Delete(_stackofgrowthcurve.StackGrowthCurveEndArcShapes, idx, idx+1)
						stackgrowthcurveendarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_stackofgrowthcurve, "StackGrowthCurveEndArcShapes", &_stackofgrowthcurve.StackGrowthCurveEndArcShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if stackgrowthcurveendarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackgrowthcurveendarcshape_.Unstage(stackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest)
	}

	stackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StackGrowthCurveEndArcShape](
		stackgrowthcurveendarcshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if stackgrowthcurveendarcshapeFormCallback.CreationMode || stackgrowthcurveendarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackgrowthcurveendarcshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(stackgrowthcurveendarcshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StackGrowthCurveEndArcShapeFormCallback(
			nil,
			stackgrowthcurveendarcshapeFormCallback.probe,
			newFormGroup,
		)
		stackgrowthcurveendarcshape := new(models.StackGrowthCurveEndArcShape)
		FillUpForm(stackgrowthcurveendarcshape, newFormGroup, stackgrowthcurveendarcshapeFormCallback.probe)
		stackgrowthcurveendarcshapeFormCallback.probe.formStage.Commit()
	}

	stackgrowthcurveendarcshapeFormCallback.probe.ux_tree()
}
func __gong__New__StackGrowthCurveStartArcShapeFormCallback(
	stackgrowthcurvestartarcshape *models.StackGrowthCurveStartArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackgrowthcurvestartarcshapeFormCallback *StackGrowthCurveStartArcShapeFormCallback) {
	stackgrowthcurvestartarcshapeFormCallback = new(StackGrowthCurveStartArcShapeFormCallback)
	stackgrowthcurvestartarcshapeFormCallback.probe = probe
	stackgrowthcurvestartarcshapeFormCallback.stackgrowthcurvestartarcshape = stackgrowthcurvestartarcshape
	stackgrowthcurvestartarcshapeFormCallback.formGroup = formGroup

	stackgrowthcurvestartarcshapeFormCallback.CreationMode = (stackgrowthcurvestartarcshape == nil)

	return
}

type StackGrowthCurveStartArcShapeFormCallback struct {
	stackgrowthcurvestartarcshape *models.StackGrowthCurveStartArcShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (stackgrowthcurvestartarcshapeFormCallback *StackGrowthCurveStartArcShapeFormCallback) OnSave() {
	stackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest.Lock()
	defer stackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StackGrowthCurveStartArcShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	stackgrowthcurvestartarcshapeFormCallback.probe.formStage.Checkout()

	if stackgrowthcurvestartarcshapeFormCallback.stackgrowthcurvestartarcshape == nil {
		stackgrowthcurvestartarcshapeFormCallback.stackgrowthcurvestartarcshape = new(models.StackGrowthCurveStartArcShape).Stage(stackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest)
	}
	stackgrowthcurvestartarcshape_ := stackgrowthcurvestartarcshapeFormCallback.stackgrowthcurvestartarcshape
	_ = stackgrowthcurvestartarcshape_

	for _, formDiv := range stackgrowthcurvestartarcshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(stackgrowthcurvestartarcshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(stackgrowthcurvestartarcshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(stackgrowthcurvestartarcshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(stackgrowthcurvestartarcshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(stackgrowthcurvestartarcshape_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(stackgrowthcurvestartarcshape_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(stackgrowthcurvestartarcshape_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(stackgrowthcurvestartarcshape_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(stackgrowthcurvestartarcshape_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(stackgrowthcurvestartarcshape_.RadiusY), formDiv)
		case "StackOfGrowthCurve:StackGrowthCurveStartArcShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the StackOfGrowthCurve instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target StackOfGrowthCurve instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.StackOfGrowthCurve](stackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest)
			targetStackOfGrowthCurveIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetStackOfGrowthCurveIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all StackOfGrowthCurve instances and update their StackGrowthCurveStartArcShapes slice
			for _stackofgrowthcurve := range *models.GetGongstructInstancesSetFromPointerType[*models.StackOfGrowthCurve](stackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(stackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest, _stackofgrowthcurve)
				
				// if StackOfGrowthCurve is selected
				if targetStackOfGrowthCurveIDs[id] {
					// ensure stackgrowthcurvestartarcshape_ is in _stackofgrowthcurve.StackGrowthCurveStartArcShapes
					found := false
					for _, _b := range _stackofgrowthcurve.StackGrowthCurveStartArcShapes {
						if _b == stackgrowthcurvestartarcshape_ {
							found = true
							break
						}
					}
					if !found {
						_stackofgrowthcurve.StackGrowthCurveStartArcShapes = append(_stackofgrowthcurve.StackGrowthCurveStartArcShapes, stackgrowthcurvestartarcshape_)
						stackgrowthcurvestartarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_stackofgrowthcurve, "StackGrowthCurveStartArcShapes", &_stackofgrowthcurve.StackGrowthCurveStartArcShapes)
					}
				} else {
					// ensure stackgrowthcurvestartarcshape_ is NOT in _stackofgrowthcurve.StackGrowthCurveStartArcShapes
					idx := slices.Index(_stackofgrowthcurve.StackGrowthCurveStartArcShapes, stackgrowthcurvestartarcshape_)
					if idx != -1 {
						_stackofgrowthcurve.StackGrowthCurveStartArcShapes = slices.Delete(_stackofgrowthcurve.StackGrowthCurveStartArcShapes, idx, idx+1)
						stackgrowthcurvestartarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_stackofgrowthcurve, "StackGrowthCurveStartArcShapes", &_stackofgrowthcurve.StackGrowthCurveStartArcShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if stackgrowthcurvestartarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackgrowthcurvestartarcshape_.Unstage(stackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest)
	}

	stackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StackGrowthCurveStartArcShape](
		stackgrowthcurvestartarcshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if stackgrowthcurvestartarcshapeFormCallback.CreationMode || stackgrowthcurvestartarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackgrowthcurvestartarcshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(stackgrowthcurvestartarcshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StackGrowthCurveStartArcShapeFormCallback(
			nil,
			stackgrowthcurvestartarcshapeFormCallback.probe,
			newFormGroup,
		)
		stackgrowthcurvestartarcshape := new(models.StackGrowthCurveStartArcShape)
		FillUpForm(stackgrowthcurvestartarcshape, newFormGroup, stackgrowthcurvestartarcshapeFormCallback.probe)
		stackgrowthcurvestartarcshapeFormCallback.probe.formStage.Commit()
	}

	stackgrowthcurvestartarcshapeFormCallback.probe.ux_tree()
}
func __gong__New__StackOfGrowthCurveFormCallback(
	stackofgrowthcurve *models.StackOfGrowthCurve,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackofgrowthcurveFormCallback *StackOfGrowthCurveFormCallback) {
	stackofgrowthcurveFormCallback = new(StackOfGrowthCurveFormCallback)
	stackofgrowthcurveFormCallback.probe = probe
	stackofgrowthcurveFormCallback.stackofgrowthcurve = stackofgrowthcurve
	stackofgrowthcurveFormCallback.formGroup = formGroup

	stackofgrowthcurveFormCallback.CreationMode = (stackofgrowthcurve == nil)

	return
}

type StackOfGrowthCurveFormCallback struct {
	stackofgrowthcurve *models.StackOfGrowthCurve

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (stackofgrowthcurveFormCallback *StackOfGrowthCurveFormCallback) OnSave() {
	stackofgrowthcurveFormCallback.probe.stageOfInterest.Lock()
	defer stackofgrowthcurveFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StackOfGrowthCurveFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	stackofgrowthcurveFormCallback.probe.formStage.Checkout()

	if stackofgrowthcurveFormCallback.stackofgrowthcurve == nil {
		stackofgrowthcurveFormCallback.stackofgrowthcurve = new(models.StackOfGrowthCurve).Stage(stackofgrowthcurveFormCallback.probe.stageOfInterest)
	}
	stackofgrowthcurve_ := stackofgrowthcurveFormCallback.stackofgrowthcurve
	_ = stackofgrowthcurve_

	for _, formDiv := range stackofgrowthcurveFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(stackofgrowthcurve_.Name), formDiv)
		case "StackGrowthCurveStartArcShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.StackGrowthCurveStartArcShape](stackofgrowthcurveFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.StackGrowthCurveStartArcShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.StackGrowthCurveStartArcShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					stackofgrowthcurveFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.StackGrowthCurveStartArcShape](stackofgrowthcurveFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			stackofgrowthcurve_.StackGrowthCurveStartArcShapes = instanceSlice
			stackofgrowthcurveFormCallback.probe.UpdateSliceOfPointersCallback(stackofgrowthcurve_, "StackGrowthCurveStartArcShapes", &stackofgrowthcurve_.StackGrowthCurveStartArcShapes)

		case "StackGrowthCurveEndArcShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.StackGrowthCurveEndArcShape](stackofgrowthcurveFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.StackGrowthCurveEndArcShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.StackGrowthCurveEndArcShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					stackofgrowthcurveFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.StackGrowthCurveEndArcShape](stackofgrowthcurveFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			stackofgrowthcurve_.StackGrowthCurveEndArcShapes = instanceSlice
			stackofgrowthcurveFormCallback.probe.UpdateSliceOfPointersCallback(stackofgrowthcurve_, "StackGrowthCurveEndArcShapes", &stackofgrowthcurve_.StackGrowthCurveEndArcShapes)

		}
	}

	// manage the suppress operation
	if stackofgrowthcurveFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackofgrowthcurve_.Unstage(stackofgrowthcurveFormCallback.probe.stageOfInterest)
	}

	stackofgrowthcurveFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StackOfGrowthCurve](
		stackofgrowthcurveFormCallback.probe,
	)

	// display a new form by reset the form stage
	if stackofgrowthcurveFormCallback.CreationMode || stackofgrowthcurveFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackofgrowthcurveFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(stackofgrowthcurveFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StackOfGrowthCurveFormCallback(
			nil,
			stackofgrowthcurveFormCallback.probe,
			newFormGroup,
		)
		stackofgrowthcurve := new(models.StackOfGrowthCurve)
		FillUpForm(stackofgrowthcurve, newFormGroup, stackofgrowthcurveFormCallback.probe)
		stackofgrowthcurveFormCallback.probe.formStage.Commit()
	}

	stackofgrowthcurveFormCallback.probe.ux_tree()
}
func __gong__New__StartArcShapeFormCallback(
	startarcshape *models.StartArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (startarcshapeFormCallback *StartArcShapeFormCallback) {
	startarcshapeFormCallback = new(StartArcShapeFormCallback)
	startarcshapeFormCallback.probe = probe
	startarcshapeFormCallback.startarcshape = startarcshape
	startarcshapeFormCallback.formGroup = formGroup

	startarcshapeFormCallback.CreationMode = (startarcshape == nil)

	return
}

type StartArcShapeFormCallback struct {
	startarcshape *models.StartArcShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (startarcshapeFormCallback *StartArcShapeFormCallback) OnSave() {
	startarcshapeFormCallback.probe.stageOfInterest.Lock()
	defer startarcshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StartArcShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	startarcshapeFormCallback.probe.formStage.Checkout()

	if startarcshapeFormCallback.startarcshape == nil {
		startarcshapeFormCallback.startarcshape = new(models.StartArcShape).Stage(startarcshapeFormCallback.probe.stageOfInterest)
	}
	startarcshape_ := startarcshapeFormCallback.startarcshape
	_ = startarcshape_

	for _, formDiv := range startarcshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(startarcshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(startarcshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(startarcshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(startarcshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(startarcshape_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(startarcshape_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(startarcshape_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(startarcshape_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(startarcshape_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(startarcshape_.RadiusY), formDiv)
		case "StartArcShapeGrid:StartArcShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the StartArcShapeGrid instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target StartArcShapeGrid instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.StartArcShapeGrid](startarcshapeFormCallback.probe.stageOfInterest)
			targetStartArcShapeGridIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetStartArcShapeGridIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all StartArcShapeGrid instances and update their StartArcShapes slice
			for _startarcshapegrid := range *models.GetGongstructInstancesSetFromPointerType[*models.StartArcShapeGrid](startarcshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(startarcshapeFormCallback.probe.stageOfInterest, _startarcshapegrid)
				
				// if StartArcShapeGrid is selected
				if targetStartArcShapeGridIDs[id] {
					// ensure startarcshape_ is in _startarcshapegrid.StartArcShapes
					found := false
					for _, _b := range _startarcshapegrid.StartArcShapes {
						if _b == startarcshape_ {
							found = true
							break
						}
					}
					if !found {
						_startarcshapegrid.StartArcShapes = append(_startarcshapegrid.StartArcShapes, startarcshape_)
						startarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_startarcshapegrid, "StartArcShapes", &_startarcshapegrid.StartArcShapes)
					}
				} else {
					// ensure startarcshape_ is NOT in _startarcshapegrid.StartArcShapes
					idx := slices.Index(_startarcshapegrid.StartArcShapes, startarcshape_)
					if idx != -1 {
						_startarcshapegrid.StartArcShapes = slices.Delete(_startarcshapegrid.StartArcShapes, idx, idx+1)
						startarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_startarcshapegrid, "StartArcShapes", &_startarcshapegrid.StartArcShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if startarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		startarcshape_.Unstage(startarcshapeFormCallback.probe.stageOfInterest)
	}

	startarcshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StartArcShape](
		startarcshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if startarcshapeFormCallback.CreationMode || startarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		startarcshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(startarcshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StartArcShapeFormCallback(
			nil,
			startarcshapeFormCallback.probe,
			newFormGroup,
		)
		startarcshape := new(models.StartArcShape)
		FillUpForm(startarcshape, newFormGroup, startarcshapeFormCallback.probe)
		startarcshapeFormCallback.probe.formStage.Commit()
	}

	startarcshapeFormCallback.probe.ux_tree()
}
func __gong__New__StartArcShapeGridFormCallback(
	startarcshapegrid *models.StartArcShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (startarcshapegridFormCallback *StartArcShapeGridFormCallback) {
	startarcshapegridFormCallback = new(StartArcShapeGridFormCallback)
	startarcshapegridFormCallback.probe = probe
	startarcshapegridFormCallback.startarcshapegrid = startarcshapegrid
	startarcshapegridFormCallback.formGroup = formGroup

	startarcshapegridFormCallback.CreationMode = (startarcshapegrid == nil)

	return
}

type StartArcShapeGridFormCallback struct {
	startarcshapegrid *models.StartArcShapeGrid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (startarcshapegridFormCallback *StartArcShapeGridFormCallback) OnSave() {
	startarcshapegridFormCallback.probe.stageOfInterest.Lock()
	defer startarcshapegridFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StartArcShapeGridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	startarcshapegridFormCallback.probe.formStage.Checkout()

	if startarcshapegridFormCallback.startarcshapegrid == nil {
		startarcshapegridFormCallback.startarcshapegrid = new(models.StartArcShapeGrid).Stage(startarcshapegridFormCallback.probe.stageOfInterest)
	}
	startarcshapegrid_ := startarcshapegridFormCallback.startarcshapegrid
	_ = startarcshapegrid_

	for _, formDiv := range startarcshapegridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(startarcshapegrid_.Name), formDiv)
		case "StartArcShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.StartArcShape](startarcshapegridFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.StartArcShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.StartArcShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					startarcshapegridFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.StartArcShape](startarcshapegridFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			startarcshapegrid_.StartArcShapes = instanceSlice
			startarcshapegridFormCallback.probe.UpdateSliceOfPointersCallback(startarcshapegrid_, "StartArcShapes", &startarcshapegrid_.StartArcShapes)

		}
	}

	// manage the suppress operation
	if startarcshapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		startarcshapegrid_.Unstage(startarcshapegridFormCallback.probe.stageOfInterest)
	}

	startarcshapegridFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StartArcShapeGrid](
		startarcshapegridFormCallback.probe,
	)

	// display a new form by reset the form stage
	if startarcshapegridFormCallback.CreationMode || startarcshapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		startarcshapegridFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(startarcshapegridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StartArcShapeGridFormCallback(
			nil,
			startarcshapegridFormCallback.probe,
			newFormGroup,
		)
		startarcshapegrid := new(models.StartArcShapeGrid)
		FillUpForm(startarcshapegrid, newFormGroup, startarcshapegridFormCallback.probe)
		startarcshapegridFormCallback.probe.formStage.Commit()
	}

	startarcshapegridFormCallback.probe.ux_tree()
}
func __gong__New__TopEndArcShapeFormCallback(
	topendarcshape *models.TopEndArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (topendarcshapeFormCallback *TopEndArcShapeFormCallback) {
	topendarcshapeFormCallback = new(TopEndArcShapeFormCallback)
	topendarcshapeFormCallback.probe = probe
	topendarcshapeFormCallback.topendarcshape = topendarcshape
	topendarcshapeFormCallback.formGroup = formGroup

	topendarcshapeFormCallback.CreationMode = (topendarcshape == nil)

	return
}

type TopEndArcShapeFormCallback struct {
	topendarcshape *models.TopEndArcShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (topendarcshapeFormCallback *TopEndArcShapeFormCallback) OnSave() {
	topendarcshapeFormCallback.probe.stageOfInterest.Lock()
	defer topendarcshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TopEndArcShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	topendarcshapeFormCallback.probe.formStage.Checkout()

	if topendarcshapeFormCallback.topendarcshape == nil {
		topendarcshapeFormCallback.topendarcshape = new(models.TopEndArcShape).Stage(topendarcshapeFormCallback.probe.stageOfInterest)
	}
	topendarcshape_ := topendarcshapeFormCallback.topendarcshape
	_ = topendarcshape_

	for _, formDiv := range topendarcshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(topendarcshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(topendarcshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(topendarcshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(topendarcshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(topendarcshape_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(topendarcshape_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(topendarcshape_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(topendarcshape_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(topendarcshape_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(topendarcshape_.RadiusY), formDiv)
		case "TopEndArcShapeGrid:TopEndArcShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the TopEndArcShapeGrid instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target TopEndArcShapeGrid instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.TopEndArcShapeGrid](topendarcshapeFormCallback.probe.stageOfInterest)
			targetTopEndArcShapeGridIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetTopEndArcShapeGridIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all TopEndArcShapeGrid instances and update their TopEndArcShapes slice
			for _topendarcshapegrid := range *models.GetGongstructInstancesSetFromPointerType[*models.TopEndArcShapeGrid](topendarcshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(topendarcshapeFormCallback.probe.stageOfInterest, _topendarcshapegrid)
				
				// if TopEndArcShapeGrid is selected
				if targetTopEndArcShapeGridIDs[id] {
					// ensure topendarcshape_ is in _topendarcshapegrid.TopEndArcShapes
					found := false
					for _, _b := range _topendarcshapegrid.TopEndArcShapes {
						if _b == topendarcshape_ {
							found = true
							break
						}
					}
					if !found {
						_topendarcshapegrid.TopEndArcShapes = append(_topendarcshapegrid.TopEndArcShapes, topendarcshape_)
						topendarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_topendarcshapegrid, "TopEndArcShapes", &_topendarcshapegrid.TopEndArcShapes)
					}
				} else {
					// ensure topendarcshape_ is NOT in _topendarcshapegrid.TopEndArcShapes
					idx := slices.Index(_topendarcshapegrid.TopEndArcShapes, topendarcshape_)
					if idx != -1 {
						_topendarcshapegrid.TopEndArcShapes = slices.Delete(_topendarcshapegrid.TopEndArcShapes, idx, idx+1)
						topendarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_topendarcshapegrid, "TopEndArcShapes", &_topendarcshapegrid.TopEndArcShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if topendarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topendarcshape_.Unstage(topendarcshapeFormCallback.probe.stageOfInterest)
	}

	topendarcshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TopEndArcShape](
		topendarcshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if topendarcshapeFormCallback.CreationMode || topendarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topendarcshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(topendarcshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TopEndArcShapeFormCallback(
			nil,
			topendarcshapeFormCallback.probe,
			newFormGroup,
		)
		topendarcshape := new(models.TopEndArcShape)
		FillUpForm(topendarcshape, newFormGroup, topendarcshapeFormCallback.probe)
		topendarcshapeFormCallback.probe.formStage.Commit()
	}

	topendarcshapeFormCallback.probe.ux_tree()
}
func __gong__New__TopEndArcShapeGridFormCallback(
	topendarcshapegrid *models.TopEndArcShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (topendarcshapegridFormCallback *TopEndArcShapeGridFormCallback) {
	topendarcshapegridFormCallback = new(TopEndArcShapeGridFormCallback)
	topendarcshapegridFormCallback.probe = probe
	topendarcshapegridFormCallback.topendarcshapegrid = topendarcshapegrid
	topendarcshapegridFormCallback.formGroup = formGroup

	topendarcshapegridFormCallback.CreationMode = (topendarcshapegrid == nil)

	return
}

type TopEndArcShapeGridFormCallback struct {
	topendarcshapegrid *models.TopEndArcShapeGrid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (topendarcshapegridFormCallback *TopEndArcShapeGridFormCallback) OnSave() {
	topendarcshapegridFormCallback.probe.stageOfInterest.Lock()
	defer topendarcshapegridFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TopEndArcShapeGridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	topendarcshapegridFormCallback.probe.formStage.Checkout()

	if topendarcshapegridFormCallback.topendarcshapegrid == nil {
		topendarcshapegridFormCallback.topendarcshapegrid = new(models.TopEndArcShapeGrid).Stage(topendarcshapegridFormCallback.probe.stageOfInterest)
	}
	topendarcshapegrid_ := topendarcshapegridFormCallback.topendarcshapegrid
	_ = topendarcshapegrid_

	for _, formDiv := range topendarcshapegridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(topendarcshapegrid_.Name), formDiv)
		case "TopEndArcShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.TopEndArcShape](topendarcshapegridFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.TopEndArcShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.TopEndArcShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					topendarcshapegridFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.TopEndArcShape](topendarcshapegridFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			topendarcshapegrid_.TopEndArcShapes = instanceSlice
			topendarcshapegridFormCallback.probe.UpdateSliceOfPointersCallback(topendarcshapegrid_, "TopEndArcShapes", &topendarcshapegrid_.TopEndArcShapes)

		}
	}

	// manage the suppress operation
	if topendarcshapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topendarcshapegrid_.Unstage(topendarcshapegridFormCallback.probe.stageOfInterest)
	}

	topendarcshapegridFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TopEndArcShapeGrid](
		topendarcshapegridFormCallback.probe,
	)

	// display a new form by reset the form stage
	if topendarcshapegridFormCallback.CreationMode || topendarcshapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topendarcshapegridFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(topendarcshapegridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TopEndArcShapeGridFormCallback(
			nil,
			topendarcshapegridFormCallback.probe,
			newFormGroup,
		)
		topendarcshapegrid := new(models.TopEndArcShapeGrid)
		FillUpForm(topendarcshapegrid, newFormGroup, topendarcshapegridFormCallback.probe)
		topendarcshapegridFormCallback.probe.formStage.Commit()
	}

	topendarcshapegridFormCallback.probe.ux_tree()
}
func __gong__New__TopGrowthCurve2DFormCallback(
	topgrowthcurve2d *models.TopGrowthCurve2D,
	probe *Probe,
	formGroup *form.FormGroup,
) (topgrowthcurve2dFormCallback *TopGrowthCurve2DFormCallback) {
	topgrowthcurve2dFormCallback = new(TopGrowthCurve2DFormCallback)
	topgrowthcurve2dFormCallback.probe = probe
	topgrowthcurve2dFormCallback.topgrowthcurve2d = topgrowthcurve2d
	topgrowthcurve2dFormCallback.formGroup = formGroup

	topgrowthcurve2dFormCallback.CreationMode = (topgrowthcurve2d == nil)

	return
}

type TopGrowthCurve2DFormCallback struct {
	topgrowthcurve2d *models.TopGrowthCurve2D

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (topgrowthcurve2dFormCallback *TopGrowthCurve2DFormCallback) OnSave() {
	topgrowthcurve2dFormCallback.probe.stageOfInterest.Lock()
	defer topgrowthcurve2dFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TopGrowthCurve2DFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	topgrowthcurve2dFormCallback.probe.formStage.Checkout()

	if topgrowthcurve2dFormCallback.topgrowthcurve2d == nil {
		topgrowthcurve2dFormCallback.topgrowthcurve2d = new(models.TopGrowthCurve2D).Stage(topgrowthcurve2dFormCallback.probe.stageOfInterest)
	}
	topgrowthcurve2d_ := topgrowthcurve2dFormCallback.topgrowthcurve2d
	_ = topgrowthcurve2d_

	for _, formDiv := range topgrowthcurve2dFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(topgrowthcurve2d_.Name), formDiv)
		case "TopStartArcShapeGrid":
			FormDivSelectFieldToField(&(topgrowthcurve2d_.TopStartArcShapeGrid), topgrowthcurve2dFormCallback.probe.stageOfInterest, formDiv)
		case "TopEndArcShapeGrid":
			FormDivSelectFieldToField(&(topgrowthcurve2d_.TopEndArcShapeGrid), topgrowthcurve2dFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if topgrowthcurve2dFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topgrowthcurve2d_.Unstage(topgrowthcurve2dFormCallback.probe.stageOfInterest)
	}

	topgrowthcurve2dFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TopGrowthCurve2D](
		topgrowthcurve2dFormCallback.probe,
	)

	// display a new form by reset the form stage
	if topgrowthcurve2dFormCallback.CreationMode || topgrowthcurve2dFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topgrowthcurve2dFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(topgrowthcurve2dFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TopGrowthCurve2DFormCallback(
			nil,
			topgrowthcurve2dFormCallback.probe,
			newFormGroup,
		)
		topgrowthcurve2d := new(models.TopGrowthCurve2D)
		FillUpForm(topgrowthcurve2d, newFormGroup, topgrowthcurve2dFormCallback.probe)
		topgrowthcurve2dFormCallback.probe.formStage.Commit()
	}

	topgrowthcurve2dFormCallback.probe.ux_tree()
}
func __gong__New__TopStackGrowthCurveEndArcShapeFormCallback(
	topstackgrowthcurveendarcshape *models.TopStackGrowthCurveEndArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (topstackgrowthcurveendarcshapeFormCallback *TopStackGrowthCurveEndArcShapeFormCallback) {
	topstackgrowthcurveendarcshapeFormCallback = new(TopStackGrowthCurveEndArcShapeFormCallback)
	topstackgrowthcurveendarcshapeFormCallback.probe = probe
	topstackgrowthcurveendarcshapeFormCallback.topstackgrowthcurveendarcshape = topstackgrowthcurveendarcshape
	topstackgrowthcurveendarcshapeFormCallback.formGroup = formGroup

	topstackgrowthcurveendarcshapeFormCallback.CreationMode = (topstackgrowthcurveendarcshape == nil)

	return
}

type TopStackGrowthCurveEndArcShapeFormCallback struct {
	topstackgrowthcurveendarcshape *models.TopStackGrowthCurveEndArcShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (topstackgrowthcurveendarcshapeFormCallback *TopStackGrowthCurveEndArcShapeFormCallback) OnSave() {
	topstackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest.Lock()
	defer topstackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TopStackGrowthCurveEndArcShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	topstackgrowthcurveendarcshapeFormCallback.probe.formStage.Checkout()

	if topstackgrowthcurveendarcshapeFormCallback.topstackgrowthcurveendarcshape == nil {
		topstackgrowthcurveendarcshapeFormCallback.topstackgrowthcurveendarcshape = new(models.TopStackGrowthCurveEndArcShape).Stage(topstackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest)
	}
	topstackgrowthcurveendarcshape_ := topstackgrowthcurveendarcshapeFormCallback.topstackgrowthcurveendarcshape
	_ = topstackgrowthcurveendarcshape_

	for _, formDiv := range topstackgrowthcurveendarcshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(topstackgrowthcurveendarcshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(topstackgrowthcurveendarcshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(topstackgrowthcurveendarcshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(topstackgrowthcurveendarcshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(topstackgrowthcurveendarcshape_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(topstackgrowthcurveendarcshape_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(topstackgrowthcurveendarcshape_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(topstackgrowthcurveendarcshape_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(topstackgrowthcurveendarcshape_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(topstackgrowthcurveendarcshape_.RadiusY), formDiv)
		case "TopStackOfGrowthCurve:TopStackGrowthCurveEndArcShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the TopStackOfGrowthCurve instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target TopStackOfGrowthCurve instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.TopStackOfGrowthCurve](topstackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest)
			targetTopStackOfGrowthCurveIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetTopStackOfGrowthCurveIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all TopStackOfGrowthCurve instances and update their TopStackGrowthCurveEndArcShapes slice
			for _topstackofgrowthcurve := range *models.GetGongstructInstancesSetFromPointerType[*models.TopStackOfGrowthCurve](topstackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(topstackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest, _topstackofgrowthcurve)
				
				// if TopStackOfGrowthCurve is selected
				if targetTopStackOfGrowthCurveIDs[id] {
					// ensure topstackgrowthcurveendarcshape_ is in _topstackofgrowthcurve.TopStackGrowthCurveEndArcShapes
					found := false
					for _, _b := range _topstackofgrowthcurve.TopStackGrowthCurveEndArcShapes {
						if _b == topstackgrowthcurveendarcshape_ {
							found = true
							break
						}
					}
					if !found {
						_topstackofgrowthcurve.TopStackGrowthCurveEndArcShapes = append(_topstackofgrowthcurve.TopStackGrowthCurveEndArcShapes, topstackgrowthcurveendarcshape_)
						topstackgrowthcurveendarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_topstackofgrowthcurve, "TopStackGrowthCurveEndArcShapes", &_topstackofgrowthcurve.TopStackGrowthCurveEndArcShapes)
					}
				} else {
					// ensure topstackgrowthcurveendarcshape_ is NOT in _topstackofgrowthcurve.TopStackGrowthCurveEndArcShapes
					idx := slices.Index(_topstackofgrowthcurve.TopStackGrowthCurveEndArcShapes, topstackgrowthcurveendarcshape_)
					if idx != -1 {
						_topstackofgrowthcurve.TopStackGrowthCurveEndArcShapes = slices.Delete(_topstackofgrowthcurve.TopStackGrowthCurveEndArcShapes, idx, idx+1)
						topstackgrowthcurveendarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_topstackofgrowthcurve, "TopStackGrowthCurveEndArcShapes", &_topstackofgrowthcurve.TopStackGrowthCurveEndArcShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if topstackgrowthcurveendarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstackgrowthcurveendarcshape_.Unstage(topstackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest)
	}

	topstackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TopStackGrowthCurveEndArcShape](
		topstackgrowthcurveendarcshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if topstackgrowthcurveendarcshapeFormCallback.CreationMode || topstackgrowthcurveendarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstackgrowthcurveendarcshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(topstackgrowthcurveendarcshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TopStackGrowthCurveEndArcShapeFormCallback(
			nil,
			topstackgrowthcurveendarcshapeFormCallback.probe,
			newFormGroup,
		)
		topstackgrowthcurveendarcshape := new(models.TopStackGrowthCurveEndArcShape)
		FillUpForm(topstackgrowthcurveendarcshape, newFormGroup, topstackgrowthcurveendarcshapeFormCallback.probe)
		topstackgrowthcurveendarcshapeFormCallback.probe.formStage.Commit()
	}

	topstackgrowthcurveendarcshapeFormCallback.probe.ux_tree()
}
func __gong__New__TopStackGrowthCurveStartArcShapeFormCallback(
	topstackgrowthcurvestartarcshape *models.TopStackGrowthCurveStartArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (topstackgrowthcurvestartarcshapeFormCallback *TopStackGrowthCurveStartArcShapeFormCallback) {
	topstackgrowthcurvestartarcshapeFormCallback = new(TopStackGrowthCurveStartArcShapeFormCallback)
	topstackgrowthcurvestartarcshapeFormCallback.probe = probe
	topstackgrowthcurvestartarcshapeFormCallback.topstackgrowthcurvestartarcshape = topstackgrowthcurvestartarcshape
	topstackgrowthcurvestartarcshapeFormCallback.formGroup = formGroup

	topstackgrowthcurvestartarcshapeFormCallback.CreationMode = (topstackgrowthcurvestartarcshape == nil)

	return
}

type TopStackGrowthCurveStartArcShapeFormCallback struct {
	topstackgrowthcurvestartarcshape *models.TopStackGrowthCurveStartArcShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (topstackgrowthcurvestartarcshapeFormCallback *TopStackGrowthCurveStartArcShapeFormCallback) OnSave() {
	topstackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest.Lock()
	defer topstackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TopStackGrowthCurveStartArcShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	topstackgrowthcurvestartarcshapeFormCallback.probe.formStage.Checkout()

	if topstackgrowthcurvestartarcshapeFormCallback.topstackgrowthcurvestartarcshape == nil {
		topstackgrowthcurvestartarcshapeFormCallback.topstackgrowthcurvestartarcshape = new(models.TopStackGrowthCurveStartArcShape).Stage(topstackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest)
	}
	topstackgrowthcurvestartarcshape_ := topstackgrowthcurvestartarcshapeFormCallback.topstackgrowthcurvestartarcshape
	_ = topstackgrowthcurvestartarcshape_

	for _, formDiv := range topstackgrowthcurvestartarcshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(topstackgrowthcurvestartarcshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(topstackgrowthcurvestartarcshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(topstackgrowthcurvestartarcshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(topstackgrowthcurvestartarcshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(topstackgrowthcurvestartarcshape_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(topstackgrowthcurvestartarcshape_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(topstackgrowthcurvestartarcshape_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(topstackgrowthcurvestartarcshape_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(topstackgrowthcurvestartarcshape_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(topstackgrowthcurvestartarcshape_.RadiusY), formDiv)
		case "TopStackOfGrowthCurve:TopStackGrowthCurveStartArcShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the TopStackOfGrowthCurve instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target TopStackOfGrowthCurve instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.TopStackOfGrowthCurve](topstackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest)
			targetTopStackOfGrowthCurveIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetTopStackOfGrowthCurveIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all TopStackOfGrowthCurve instances and update their TopStackGrowthCurveStartArcShapes slice
			for _topstackofgrowthcurve := range *models.GetGongstructInstancesSetFromPointerType[*models.TopStackOfGrowthCurve](topstackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(topstackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest, _topstackofgrowthcurve)
				
				// if TopStackOfGrowthCurve is selected
				if targetTopStackOfGrowthCurveIDs[id] {
					// ensure topstackgrowthcurvestartarcshape_ is in _topstackofgrowthcurve.TopStackGrowthCurveStartArcShapes
					found := false
					for _, _b := range _topstackofgrowthcurve.TopStackGrowthCurveStartArcShapes {
						if _b == topstackgrowthcurvestartarcshape_ {
							found = true
							break
						}
					}
					if !found {
						_topstackofgrowthcurve.TopStackGrowthCurveStartArcShapes = append(_topstackofgrowthcurve.TopStackGrowthCurveStartArcShapes, topstackgrowthcurvestartarcshape_)
						topstackgrowthcurvestartarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_topstackofgrowthcurve, "TopStackGrowthCurveStartArcShapes", &_topstackofgrowthcurve.TopStackGrowthCurveStartArcShapes)
					}
				} else {
					// ensure topstackgrowthcurvestartarcshape_ is NOT in _topstackofgrowthcurve.TopStackGrowthCurveStartArcShapes
					idx := slices.Index(_topstackofgrowthcurve.TopStackGrowthCurveStartArcShapes, topstackgrowthcurvestartarcshape_)
					if idx != -1 {
						_topstackofgrowthcurve.TopStackGrowthCurveStartArcShapes = slices.Delete(_topstackofgrowthcurve.TopStackGrowthCurveStartArcShapes, idx, idx+1)
						topstackgrowthcurvestartarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_topstackofgrowthcurve, "TopStackGrowthCurveStartArcShapes", &_topstackofgrowthcurve.TopStackGrowthCurveStartArcShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if topstackgrowthcurvestartarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstackgrowthcurvestartarcshape_.Unstage(topstackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest)
	}

	topstackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TopStackGrowthCurveStartArcShape](
		topstackgrowthcurvestartarcshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if topstackgrowthcurvestartarcshapeFormCallback.CreationMode || topstackgrowthcurvestartarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstackgrowthcurvestartarcshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(topstackgrowthcurvestartarcshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TopStackGrowthCurveStartArcShapeFormCallback(
			nil,
			topstackgrowthcurvestartarcshapeFormCallback.probe,
			newFormGroup,
		)
		topstackgrowthcurvestartarcshape := new(models.TopStackGrowthCurveStartArcShape)
		FillUpForm(topstackgrowthcurvestartarcshape, newFormGroup, topstackgrowthcurvestartarcshapeFormCallback.probe)
		topstackgrowthcurvestartarcshapeFormCallback.probe.formStage.Commit()
	}

	topstackgrowthcurvestartarcshapeFormCallback.probe.ux_tree()
}
func __gong__New__TopStackOfGrowthCurveFormCallback(
	topstackofgrowthcurve *models.TopStackOfGrowthCurve,
	probe *Probe,
	formGroup *form.FormGroup,
) (topstackofgrowthcurveFormCallback *TopStackOfGrowthCurveFormCallback) {
	topstackofgrowthcurveFormCallback = new(TopStackOfGrowthCurveFormCallback)
	topstackofgrowthcurveFormCallback.probe = probe
	topstackofgrowthcurveFormCallback.topstackofgrowthcurve = topstackofgrowthcurve
	topstackofgrowthcurveFormCallback.formGroup = formGroup

	topstackofgrowthcurveFormCallback.CreationMode = (topstackofgrowthcurve == nil)

	return
}

type TopStackOfGrowthCurveFormCallback struct {
	topstackofgrowthcurve *models.TopStackOfGrowthCurve

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (topstackofgrowthcurveFormCallback *TopStackOfGrowthCurveFormCallback) OnSave() {
	topstackofgrowthcurveFormCallback.probe.stageOfInterest.Lock()
	defer topstackofgrowthcurveFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TopStackOfGrowthCurveFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	topstackofgrowthcurveFormCallback.probe.formStage.Checkout()

	if topstackofgrowthcurveFormCallback.topstackofgrowthcurve == nil {
		topstackofgrowthcurveFormCallback.topstackofgrowthcurve = new(models.TopStackOfGrowthCurve).Stage(topstackofgrowthcurveFormCallback.probe.stageOfInterest)
	}
	topstackofgrowthcurve_ := topstackofgrowthcurveFormCallback.topstackofgrowthcurve
	_ = topstackofgrowthcurve_

	for _, formDiv := range topstackofgrowthcurveFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(topstackofgrowthcurve_.Name), formDiv)
		case "TopStackGrowthCurveStartArcShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.TopStackGrowthCurveStartArcShape](topstackofgrowthcurveFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.TopStackGrowthCurveStartArcShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.TopStackGrowthCurveStartArcShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					topstackofgrowthcurveFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.TopStackGrowthCurveStartArcShape](topstackofgrowthcurveFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			topstackofgrowthcurve_.TopStackGrowthCurveStartArcShapes = instanceSlice
			topstackofgrowthcurveFormCallback.probe.UpdateSliceOfPointersCallback(topstackofgrowthcurve_, "TopStackGrowthCurveStartArcShapes", &topstackofgrowthcurve_.TopStackGrowthCurveStartArcShapes)

		case "TopStackGrowthCurveEndArcShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.TopStackGrowthCurveEndArcShape](topstackofgrowthcurveFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.TopStackGrowthCurveEndArcShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.TopStackGrowthCurveEndArcShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					topstackofgrowthcurveFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.TopStackGrowthCurveEndArcShape](topstackofgrowthcurveFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			topstackofgrowthcurve_.TopStackGrowthCurveEndArcShapes = instanceSlice
			topstackofgrowthcurveFormCallback.probe.UpdateSliceOfPointersCallback(topstackofgrowthcurve_, "TopStackGrowthCurveEndArcShapes", &topstackofgrowthcurve_.TopStackGrowthCurveEndArcShapes)

		}
	}

	// manage the suppress operation
	if topstackofgrowthcurveFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstackofgrowthcurve_.Unstage(topstackofgrowthcurveFormCallback.probe.stageOfInterest)
	}

	topstackofgrowthcurveFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TopStackOfGrowthCurve](
		topstackofgrowthcurveFormCallback.probe,
	)

	// display a new form by reset the form stage
	if topstackofgrowthcurveFormCallback.CreationMode || topstackofgrowthcurveFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstackofgrowthcurveFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(topstackofgrowthcurveFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TopStackOfGrowthCurveFormCallback(
			nil,
			topstackofgrowthcurveFormCallback.probe,
			newFormGroup,
		)
		topstackofgrowthcurve := new(models.TopStackOfGrowthCurve)
		FillUpForm(topstackofgrowthcurve, newFormGroup, topstackofgrowthcurveFormCallback.probe)
		topstackofgrowthcurveFormCallback.probe.formStage.Commit()
	}

	topstackofgrowthcurveFormCallback.probe.ux_tree()
}
func __gong__New__TopStartArcShapeFormCallback(
	topstartarcshape *models.TopStartArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (topstartarcshapeFormCallback *TopStartArcShapeFormCallback) {
	topstartarcshapeFormCallback = new(TopStartArcShapeFormCallback)
	topstartarcshapeFormCallback.probe = probe
	topstartarcshapeFormCallback.topstartarcshape = topstartarcshape
	topstartarcshapeFormCallback.formGroup = formGroup

	topstartarcshapeFormCallback.CreationMode = (topstartarcshape == nil)

	return
}

type TopStartArcShapeFormCallback struct {
	topstartarcshape *models.TopStartArcShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (topstartarcshapeFormCallback *TopStartArcShapeFormCallback) OnSave() {
	topstartarcshapeFormCallback.probe.stageOfInterest.Lock()
	defer topstartarcshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TopStartArcShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	topstartarcshapeFormCallback.probe.formStage.Checkout()

	if topstartarcshapeFormCallback.topstartarcshape == nil {
		topstartarcshapeFormCallback.topstartarcshape = new(models.TopStartArcShape).Stage(topstartarcshapeFormCallback.probe.stageOfInterest)
	}
	topstartarcshape_ := topstartarcshapeFormCallback.topstartarcshape
	_ = topstartarcshape_

	for _, formDiv := range topstartarcshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(topstartarcshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(topstartarcshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(topstartarcshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(topstartarcshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(topstartarcshape_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(topstartarcshape_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(topstartarcshape_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(topstartarcshape_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(topstartarcshape_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(topstartarcshape_.RadiusY), formDiv)
		case "TopStartArcShapeGrid:TopStartArcShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the TopStartArcShapeGrid instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target TopStartArcShapeGrid instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.TopStartArcShapeGrid](topstartarcshapeFormCallback.probe.stageOfInterest)
			targetTopStartArcShapeGridIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetTopStartArcShapeGridIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all TopStartArcShapeGrid instances and update their TopStartArcShapes slice
			for _topstartarcshapegrid := range *models.GetGongstructInstancesSetFromPointerType[*models.TopStartArcShapeGrid](topstartarcshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(topstartarcshapeFormCallback.probe.stageOfInterest, _topstartarcshapegrid)
				
				// if TopStartArcShapeGrid is selected
				if targetTopStartArcShapeGridIDs[id] {
					// ensure topstartarcshape_ is in _topstartarcshapegrid.TopStartArcShapes
					found := false
					for _, _b := range _topstartarcshapegrid.TopStartArcShapes {
						if _b == topstartarcshape_ {
							found = true
							break
						}
					}
					if !found {
						_topstartarcshapegrid.TopStartArcShapes = append(_topstartarcshapegrid.TopStartArcShapes, topstartarcshape_)
						topstartarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_topstartarcshapegrid, "TopStartArcShapes", &_topstartarcshapegrid.TopStartArcShapes)
					}
				} else {
					// ensure topstartarcshape_ is NOT in _topstartarcshapegrid.TopStartArcShapes
					idx := slices.Index(_topstartarcshapegrid.TopStartArcShapes, topstartarcshape_)
					if idx != -1 {
						_topstartarcshapegrid.TopStartArcShapes = slices.Delete(_topstartarcshapegrid.TopStartArcShapes, idx, idx+1)
						topstartarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_topstartarcshapegrid, "TopStartArcShapes", &_topstartarcshapegrid.TopStartArcShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if topstartarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstartarcshape_.Unstage(topstartarcshapeFormCallback.probe.stageOfInterest)
	}

	topstartarcshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TopStartArcShape](
		topstartarcshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if topstartarcshapeFormCallback.CreationMode || topstartarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstartarcshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(topstartarcshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TopStartArcShapeFormCallback(
			nil,
			topstartarcshapeFormCallback.probe,
			newFormGroup,
		)
		topstartarcshape := new(models.TopStartArcShape)
		FillUpForm(topstartarcshape, newFormGroup, topstartarcshapeFormCallback.probe)
		topstartarcshapeFormCallback.probe.formStage.Commit()
	}

	topstartarcshapeFormCallback.probe.ux_tree()
}
func __gong__New__TopStartArcShapeGridFormCallback(
	topstartarcshapegrid *models.TopStartArcShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (topstartarcshapegridFormCallback *TopStartArcShapeGridFormCallback) {
	topstartarcshapegridFormCallback = new(TopStartArcShapeGridFormCallback)
	topstartarcshapegridFormCallback.probe = probe
	topstartarcshapegridFormCallback.topstartarcshapegrid = topstartarcshapegrid
	topstartarcshapegridFormCallback.formGroup = formGroup

	topstartarcshapegridFormCallback.CreationMode = (topstartarcshapegrid == nil)

	return
}

type TopStartArcShapeGridFormCallback struct {
	topstartarcshapegrid *models.TopStartArcShapeGrid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (topstartarcshapegridFormCallback *TopStartArcShapeGridFormCallback) OnSave() {
	topstartarcshapegridFormCallback.probe.stageOfInterest.Lock()
	defer topstartarcshapegridFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TopStartArcShapeGridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	topstartarcshapegridFormCallback.probe.formStage.Checkout()

	if topstartarcshapegridFormCallback.topstartarcshapegrid == nil {
		topstartarcshapegridFormCallback.topstartarcshapegrid = new(models.TopStartArcShapeGrid).Stage(topstartarcshapegridFormCallback.probe.stageOfInterest)
	}
	topstartarcshapegrid_ := topstartarcshapegridFormCallback.topstartarcshapegrid
	_ = topstartarcshapegrid_

	for _, formDiv := range topstartarcshapegridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(topstartarcshapegrid_.Name), formDiv)
		case "TopStartArcShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.TopStartArcShape](topstartarcshapegridFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.TopStartArcShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.TopStartArcShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					topstartarcshapegridFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.TopStartArcShape](topstartarcshapegridFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			topstartarcshapegrid_.TopStartArcShapes = instanceSlice
			topstartarcshapegridFormCallback.probe.UpdateSliceOfPointersCallback(topstartarcshapegrid_, "TopStartArcShapes", &topstartarcshapegrid_.TopStartArcShapes)

		}
	}

	// manage the suppress operation
	if topstartarcshapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstartarcshapegrid_.Unstage(topstartarcshapegridFormCallback.probe.stageOfInterest)
	}

	topstartarcshapegridFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TopStartArcShapeGrid](
		topstartarcshapegridFormCallback.probe,
	)

	// display a new form by reset the form stage
	if topstartarcshapegridFormCallback.CreationMode || topstartarcshapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstartarcshapegridFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(topstartarcshapegridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TopStartArcShapeGridFormCallback(
			nil,
			topstartarcshapegridFormCallback.probe,
			newFormGroup,
		)
		topstartarcshapegrid := new(models.TopStartArcShapeGrid)
		FillUpForm(topstartarcshapegrid, newFormGroup, topstartarcshapegridFormCallback.probe)
		topstartarcshapegridFormCallback.probe.formStage.Commit()
	}

	topstartarcshapegridFormCallback.probe.ux_tree()
}
