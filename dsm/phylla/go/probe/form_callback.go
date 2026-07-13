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
func __gong__New__BottomEndArcShapeV2FormCallback(
	bottomendarcshapev2 *models.BottomEndArcShapeV2,
	probe *Probe,
	formGroup *form.FormGroup,
) (bottomendarcshapev2FormCallback *BottomEndArcShapeV2FormCallback) {
	bottomendarcshapev2FormCallback = new(BottomEndArcShapeV2FormCallback)
	bottomendarcshapev2FormCallback.probe = probe
	bottomendarcshapev2FormCallback.bottomendarcshapev2 = bottomendarcshapev2
	bottomendarcshapev2FormCallback.formGroup = formGroup

	bottomendarcshapev2FormCallback.CreationMode = (bottomendarcshapev2 == nil)

	return
}

type BottomEndArcShapeV2FormCallback struct {
	bottomendarcshapev2 *models.BottomEndArcShapeV2

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (bottomendarcshapev2FormCallback *BottomEndArcShapeV2FormCallback) OnSave() {
	bottomendarcshapev2FormCallback.probe.stageOfInterest.Lock()
	defer bottomendarcshapev2FormCallback.probe.stageOfInterest.Unlock()

	// log.Println("BottomEndArcShapeV2FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bottomendarcshapev2FormCallback.probe.formStage.Checkout()

	if bottomendarcshapev2FormCallback.bottomendarcshapev2 == nil {
		bottomendarcshapev2FormCallback.bottomendarcshapev2 = new(models.BottomEndArcShapeV2).Stage(bottomendarcshapev2FormCallback.probe.stageOfInterest)
	}
	bottomendarcshapev2_ := bottomendarcshapev2FormCallback.bottomendarcshapev2
	_ = bottomendarcshapev2_

	for _, formDiv := range bottomendarcshapev2FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(bottomendarcshapev2_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(bottomendarcshapev2_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(bottomendarcshapev2_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(bottomendarcshapev2_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(bottomendarcshapev2_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(bottomendarcshapev2_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(bottomendarcshapev2_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(bottomendarcshapev2_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(bottomendarcshapev2_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(bottomendarcshapev2_.RadiusY), formDiv)
		case "BottomEndArcShapeV2Grid:BottomEndArcShapesV2":
			// 1. Decode the AssociationStorage which contains the rowIDs of the BottomEndArcShapeV2Grid instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target BottomEndArcShapeV2Grid instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.BottomEndArcShapeV2Grid](bottomendarcshapev2FormCallback.probe.stageOfInterest)
			targetBottomEndArcShapeV2GridIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetBottomEndArcShapeV2GridIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all BottomEndArcShapeV2Grid instances and update their BottomEndArcShapesV2 slice
			for _bottomendarcshapev2grid := range *models.GetGongstructInstancesSetFromPointerType[*models.BottomEndArcShapeV2Grid](bottomendarcshapev2FormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(bottomendarcshapev2FormCallback.probe.stageOfInterest, _bottomendarcshapev2grid)
				
				// if BottomEndArcShapeV2Grid is selected
				if targetBottomEndArcShapeV2GridIDs[id] {
					// ensure bottomendarcshapev2_ is in _bottomendarcshapev2grid.BottomEndArcShapesV2
					found := false
					for _, _b := range _bottomendarcshapev2grid.BottomEndArcShapesV2 {
						if _b == bottomendarcshapev2_ {
							found = true
							break
						}
					}
					if !found {
						_bottomendarcshapev2grid.BottomEndArcShapesV2 = append(_bottomendarcshapev2grid.BottomEndArcShapesV2, bottomendarcshapev2_)
						bottomendarcshapev2FormCallback.probe.UpdateSliceOfPointersCallback(_bottomendarcshapev2grid, "BottomEndArcShapesV2", &_bottomendarcshapev2grid.BottomEndArcShapesV2)
					}
				} else {
					// ensure bottomendarcshapev2_ is NOT in _bottomendarcshapev2grid.BottomEndArcShapesV2
					idx := slices.Index(_bottomendarcshapev2grid.BottomEndArcShapesV2, bottomendarcshapev2_)
					if idx != -1 {
						_bottomendarcshapev2grid.BottomEndArcShapesV2 = slices.Delete(_bottomendarcshapev2grid.BottomEndArcShapesV2, idx, idx+1)
						bottomendarcshapev2FormCallback.probe.UpdateSliceOfPointersCallback(_bottomendarcshapev2grid, "BottomEndArcShapesV2", &_bottomendarcshapev2grid.BottomEndArcShapesV2)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if bottomendarcshapev2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		bottomendarcshapev2_.Unstage(bottomendarcshapev2FormCallback.probe.stageOfInterest)
	}

	bottomendarcshapev2FormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.BottomEndArcShapeV2](
		bottomendarcshapev2FormCallback.probe,
	)

	// display a new form by reset the form stage
	if bottomendarcshapev2FormCallback.CreationMode || bottomendarcshapev2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		bottomendarcshapev2FormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(bottomendarcshapev2FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BottomEndArcShapeV2FormCallback(
			nil,
			bottomendarcshapev2FormCallback.probe,
			newFormGroup,
		)
		bottomendarcshapev2 := new(models.BottomEndArcShapeV2)
		FillUpForm(bottomendarcshapev2, newFormGroup, bottomendarcshapev2FormCallback.probe)
		bottomendarcshapev2FormCallback.probe.formStage.Commit()
	}

	bottomendarcshapev2FormCallback.probe.ux_tree()
}
func __gong__New__BottomEndArcShapeV2GridFormCallback(
	bottomendarcshapev2grid *models.BottomEndArcShapeV2Grid,
	probe *Probe,
	formGroup *form.FormGroup,
) (bottomendarcshapev2gridFormCallback *BottomEndArcShapeV2GridFormCallback) {
	bottomendarcshapev2gridFormCallback = new(BottomEndArcShapeV2GridFormCallback)
	bottomendarcshapev2gridFormCallback.probe = probe
	bottomendarcshapev2gridFormCallback.bottomendarcshapev2grid = bottomendarcshapev2grid
	bottomendarcshapev2gridFormCallback.formGroup = formGroup

	bottomendarcshapev2gridFormCallback.CreationMode = (bottomendarcshapev2grid == nil)

	return
}

type BottomEndArcShapeV2GridFormCallback struct {
	bottomendarcshapev2grid *models.BottomEndArcShapeV2Grid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (bottomendarcshapev2gridFormCallback *BottomEndArcShapeV2GridFormCallback) OnSave() {
	bottomendarcshapev2gridFormCallback.probe.stageOfInterest.Lock()
	defer bottomendarcshapev2gridFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("BottomEndArcShapeV2GridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bottomendarcshapev2gridFormCallback.probe.formStage.Checkout()

	if bottomendarcshapev2gridFormCallback.bottomendarcshapev2grid == nil {
		bottomendarcshapev2gridFormCallback.bottomendarcshapev2grid = new(models.BottomEndArcShapeV2Grid).Stage(bottomendarcshapev2gridFormCallback.probe.stageOfInterest)
	}
	bottomendarcshapev2grid_ := bottomendarcshapev2gridFormCallback.bottomendarcshapev2grid
	_ = bottomendarcshapev2grid_

	for _, formDiv := range bottomendarcshapev2gridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(bottomendarcshapev2grid_.Name), formDiv)
		case "BottomEndArcShapesV2":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.BottomEndArcShapeV2](bottomendarcshapev2gridFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.BottomEndArcShapeV2, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.BottomEndArcShapeV2)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					bottomendarcshapev2gridFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.BottomEndArcShapeV2](bottomendarcshapev2gridFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			bottomendarcshapev2grid_.BottomEndArcShapesV2 = instanceSlice
			bottomendarcshapev2gridFormCallback.probe.UpdateSliceOfPointersCallback(bottomendarcshapev2grid_, "BottomEndArcShapesV2", &bottomendarcshapev2grid_.BottomEndArcShapesV2)

		}
	}

	// manage the suppress operation
	if bottomendarcshapev2gridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bottomendarcshapev2grid_.Unstage(bottomendarcshapev2gridFormCallback.probe.stageOfInterest)
	}

	bottomendarcshapev2gridFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.BottomEndArcShapeV2Grid](
		bottomendarcshapev2gridFormCallback.probe,
	)

	// display a new form by reset the form stage
	if bottomendarcshapev2gridFormCallback.CreationMode || bottomendarcshapev2gridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bottomendarcshapev2gridFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(bottomendarcshapev2gridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BottomEndArcShapeV2GridFormCallback(
			nil,
			bottomendarcshapev2gridFormCallback.probe,
			newFormGroup,
		)
		bottomendarcshapev2grid := new(models.BottomEndArcShapeV2Grid)
		FillUpForm(bottomendarcshapev2grid, newFormGroup, bottomendarcshapev2gridFormCallback.probe)
		bottomendarcshapev2gridFormCallback.probe.formStage.Commit()
	}

	bottomendarcshapev2gridFormCallback.probe.ux_tree()
}
func __gong__New__BottomStackGrowthCurveEndArcShapeV2FormCallback(
	bottomstackgrowthcurveendarcshapev2 *models.BottomStackGrowthCurveEndArcShapeV2,
	probe *Probe,
	formGroup *form.FormGroup,
) (bottomstackgrowthcurveendarcshapev2FormCallback *BottomStackGrowthCurveEndArcShapeV2FormCallback) {
	bottomstackgrowthcurveendarcshapev2FormCallback = new(BottomStackGrowthCurveEndArcShapeV2FormCallback)
	bottomstackgrowthcurveendarcshapev2FormCallback.probe = probe
	bottomstackgrowthcurveendarcshapev2FormCallback.bottomstackgrowthcurveendarcshapev2 = bottomstackgrowthcurveendarcshapev2
	bottomstackgrowthcurveendarcshapev2FormCallback.formGroup = formGroup

	bottomstackgrowthcurveendarcshapev2FormCallback.CreationMode = (bottomstackgrowthcurveendarcshapev2 == nil)

	return
}

type BottomStackGrowthCurveEndArcShapeV2FormCallback struct {
	bottomstackgrowthcurveendarcshapev2 *models.BottomStackGrowthCurveEndArcShapeV2

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (bottomstackgrowthcurveendarcshapev2FormCallback *BottomStackGrowthCurveEndArcShapeV2FormCallback) OnSave() {
	bottomstackgrowthcurveendarcshapev2FormCallback.probe.stageOfInterest.Lock()
	defer bottomstackgrowthcurveendarcshapev2FormCallback.probe.stageOfInterest.Unlock()

	// log.Println("BottomStackGrowthCurveEndArcShapeV2FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bottomstackgrowthcurveendarcshapev2FormCallback.probe.formStage.Checkout()

	if bottomstackgrowthcurveendarcshapev2FormCallback.bottomstackgrowthcurveendarcshapev2 == nil {
		bottomstackgrowthcurveendarcshapev2FormCallback.bottomstackgrowthcurveendarcshapev2 = new(models.BottomStackGrowthCurveEndArcShapeV2).Stage(bottomstackgrowthcurveendarcshapev2FormCallback.probe.stageOfInterest)
	}
	bottomstackgrowthcurveendarcshapev2_ := bottomstackgrowthcurveendarcshapev2FormCallback.bottomstackgrowthcurveendarcshapev2
	_ = bottomstackgrowthcurveendarcshapev2_

	for _, formDiv := range bottomstackgrowthcurveendarcshapev2FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(bottomstackgrowthcurveendarcshapev2_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(bottomstackgrowthcurveendarcshapev2_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(bottomstackgrowthcurveendarcshapev2_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(bottomstackgrowthcurveendarcshapev2_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(bottomstackgrowthcurveendarcshapev2_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(bottomstackgrowthcurveendarcshapev2_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(bottomstackgrowthcurveendarcshapev2_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(bottomstackgrowthcurveendarcshapev2_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(bottomstackgrowthcurveendarcshapev2_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(bottomstackgrowthcurveendarcshapev2_.RadiusY), formDiv)
		case "BottomStackOfGrowthCurveV2:BottomStackGrowthCurveEndArcShapeV2s":
			// 1. Decode the AssociationStorage which contains the rowIDs of the BottomStackOfGrowthCurveV2 instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target BottomStackOfGrowthCurveV2 instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.BottomStackOfGrowthCurveV2](bottomstackgrowthcurveendarcshapev2FormCallback.probe.stageOfInterest)
			targetBottomStackOfGrowthCurveV2IDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetBottomStackOfGrowthCurveV2IDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all BottomStackOfGrowthCurveV2 instances and update their BottomStackGrowthCurveEndArcShapeV2s slice
			for _bottomstackofgrowthcurvev2 := range *models.GetGongstructInstancesSetFromPointerType[*models.BottomStackOfGrowthCurveV2](bottomstackgrowthcurveendarcshapev2FormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(bottomstackgrowthcurveendarcshapev2FormCallback.probe.stageOfInterest, _bottomstackofgrowthcurvev2)
				
				// if BottomStackOfGrowthCurveV2 is selected
				if targetBottomStackOfGrowthCurveV2IDs[id] {
					// ensure bottomstackgrowthcurveendarcshapev2_ is in _bottomstackofgrowthcurvev2.BottomStackGrowthCurveEndArcShapeV2s
					found := false
					for _, _b := range _bottomstackofgrowthcurvev2.BottomStackGrowthCurveEndArcShapeV2s {
						if _b == bottomstackgrowthcurveendarcshapev2_ {
							found = true
							break
						}
					}
					if !found {
						_bottomstackofgrowthcurvev2.BottomStackGrowthCurveEndArcShapeV2s = append(_bottomstackofgrowthcurvev2.BottomStackGrowthCurveEndArcShapeV2s, bottomstackgrowthcurveendarcshapev2_)
						bottomstackgrowthcurveendarcshapev2FormCallback.probe.UpdateSliceOfPointersCallback(_bottomstackofgrowthcurvev2, "BottomStackGrowthCurveEndArcShapeV2s", &_bottomstackofgrowthcurvev2.BottomStackGrowthCurveEndArcShapeV2s)
					}
				} else {
					// ensure bottomstackgrowthcurveendarcshapev2_ is NOT in _bottomstackofgrowthcurvev2.BottomStackGrowthCurveEndArcShapeV2s
					idx := slices.Index(_bottomstackofgrowthcurvev2.BottomStackGrowthCurveEndArcShapeV2s, bottomstackgrowthcurveendarcshapev2_)
					if idx != -1 {
						_bottomstackofgrowthcurvev2.BottomStackGrowthCurveEndArcShapeV2s = slices.Delete(_bottomstackofgrowthcurvev2.BottomStackGrowthCurveEndArcShapeV2s, idx, idx+1)
						bottomstackgrowthcurveendarcshapev2FormCallback.probe.UpdateSliceOfPointersCallback(_bottomstackofgrowthcurvev2, "BottomStackGrowthCurveEndArcShapeV2s", &_bottomstackofgrowthcurvev2.BottomStackGrowthCurveEndArcShapeV2s)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if bottomstackgrowthcurveendarcshapev2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		bottomstackgrowthcurveendarcshapev2_.Unstage(bottomstackgrowthcurveendarcshapev2FormCallback.probe.stageOfInterest)
	}

	bottomstackgrowthcurveendarcshapev2FormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.BottomStackGrowthCurveEndArcShapeV2](
		bottomstackgrowthcurveendarcshapev2FormCallback.probe,
	)

	// display a new form by reset the form stage
	if bottomstackgrowthcurveendarcshapev2FormCallback.CreationMode || bottomstackgrowthcurveendarcshapev2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		bottomstackgrowthcurveendarcshapev2FormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(bottomstackgrowthcurveendarcshapev2FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BottomStackGrowthCurveEndArcShapeV2FormCallback(
			nil,
			bottomstackgrowthcurveendarcshapev2FormCallback.probe,
			newFormGroup,
		)
		bottomstackgrowthcurveendarcshapev2 := new(models.BottomStackGrowthCurveEndArcShapeV2)
		FillUpForm(bottomstackgrowthcurveendarcshapev2, newFormGroup, bottomstackgrowthcurveendarcshapev2FormCallback.probe)
		bottomstackgrowthcurveendarcshapev2FormCallback.probe.formStage.Commit()
	}

	bottomstackgrowthcurveendarcshapev2FormCallback.probe.ux_tree()
}
func __gong__New__BottomStackGrowthCurveStartArcShapeV2FormCallback(
	bottomstackgrowthcurvestartarcshapev2 *models.BottomStackGrowthCurveStartArcShapeV2,
	probe *Probe,
	formGroup *form.FormGroup,
) (bottomstackgrowthcurvestartarcshapev2FormCallback *BottomStackGrowthCurveStartArcShapeV2FormCallback) {
	bottomstackgrowthcurvestartarcshapev2FormCallback = new(BottomStackGrowthCurveStartArcShapeV2FormCallback)
	bottomstackgrowthcurvestartarcshapev2FormCallback.probe = probe
	bottomstackgrowthcurvestartarcshapev2FormCallback.bottomstackgrowthcurvestartarcshapev2 = bottomstackgrowthcurvestartarcshapev2
	bottomstackgrowthcurvestartarcshapev2FormCallback.formGroup = formGroup

	bottomstackgrowthcurvestartarcshapev2FormCallback.CreationMode = (bottomstackgrowthcurvestartarcshapev2 == nil)

	return
}

type BottomStackGrowthCurveStartArcShapeV2FormCallback struct {
	bottomstackgrowthcurvestartarcshapev2 *models.BottomStackGrowthCurveStartArcShapeV2

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (bottomstackgrowthcurvestartarcshapev2FormCallback *BottomStackGrowthCurveStartArcShapeV2FormCallback) OnSave() {
	bottomstackgrowthcurvestartarcshapev2FormCallback.probe.stageOfInterest.Lock()
	defer bottomstackgrowthcurvestartarcshapev2FormCallback.probe.stageOfInterest.Unlock()

	// log.Println("BottomStackGrowthCurveStartArcShapeV2FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bottomstackgrowthcurvestartarcshapev2FormCallback.probe.formStage.Checkout()

	if bottomstackgrowthcurvestartarcshapev2FormCallback.bottomstackgrowthcurvestartarcshapev2 == nil {
		bottomstackgrowthcurvestartarcshapev2FormCallback.bottomstackgrowthcurvestartarcshapev2 = new(models.BottomStackGrowthCurveStartArcShapeV2).Stage(bottomstackgrowthcurvestartarcshapev2FormCallback.probe.stageOfInterest)
	}
	bottomstackgrowthcurvestartarcshapev2_ := bottomstackgrowthcurvestartarcshapev2FormCallback.bottomstackgrowthcurvestartarcshapev2
	_ = bottomstackgrowthcurvestartarcshapev2_

	for _, formDiv := range bottomstackgrowthcurvestartarcshapev2FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(bottomstackgrowthcurvestartarcshapev2_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(bottomstackgrowthcurvestartarcshapev2_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(bottomstackgrowthcurvestartarcshapev2_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(bottomstackgrowthcurvestartarcshapev2_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(bottomstackgrowthcurvestartarcshapev2_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(bottomstackgrowthcurvestartarcshapev2_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(bottomstackgrowthcurvestartarcshapev2_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(bottomstackgrowthcurvestartarcshapev2_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(bottomstackgrowthcurvestartarcshapev2_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(bottomstackgrowthcurvestartarcshapev2_.RadiusY), formDiv)
		case "BottomStackOfGrowthCurveV2:BottomStackGrowthCurveStartArcShapeV2s":
			// 1. Decode the AssociationStorage which contains the rowIDs of the BottomStackOfGrowthCurveV2 instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target BottomStackOfGrowthCurveV2 instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.BottomStackOfGrowthCurveV2](bottomstackgrowthcurvestartarcshapev2FormCallback.probe.stageOfInterest)
			targetBottomStackOfGrowthCurveV2IDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetBottomStackOfGrowthCurveV2IDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all BottomStackOfGrowthCurveV2 instances and update their BottomStackGrowthCurveStartArcShapeV2s slice
			for _bottomstackofgrowthcurvev2 := range *models.GetGongstructInstancesSetFromPointerType[*models.BottomStackOfGrowthCurveV2](bottomstackgrowthcurvestartarcshapev2FormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(bottomstackgrowthcurvestartarcshapev2FormCallback.probe.stageOfInterest, _bottomstackofgrowthcurvev2)
				
				// if BottomStackOfGrowthCurveV2 is selected
				if targetBottomStackOfGrowthCurveV2IDs[id] {
					// ensure bottomstackgrowthcurvestartarcshapev2_ is in _bottomstackofgrowthcurvev2.BottomStackGrowthCurveStartArcShapeV2s
					found := false
					for _, _b := range _bottomstackofgrowthcurvev2.BottomStackGrowthCurveStartArcShapeV2s {
						if _b == bottomstackgrowthcurvestartarcshapev2_ {
							found = true
							break
						}
					}
					if !found {
						_bottomstackofgrowthcurvev2.BottomStackGrowthCurveStartArcShapeV2s = append(_bottomstackofgrowthcurvev2.BottomStackGrowthCurveStartArcShapeV2s, bottomstackgrowthcurvestartarcshapev2_)
						bottomstackgrowthcurvestartarcshapev2FormCallback.probe.UpdateSliceOfPointersCallback(_bottomstackofgrowthcurvev2, "BottomStackGrowthCurveStartArcShapeV2s", &_bottomstackofgrowthcurvev2.BottomStackGrowthCurveStartArcShapeV2s)
					}
				} else {
					// ensure bottomstackgrowthcurvestartarcshapev2_ is NOT in _bottomstackofgrowthcurvev2.BottomStackGrowthCurveStartArcShapeV2s
					idx := slices.Index(_bottomstackofgrowthcurvev2.BottomStackGrowthCurveStartArcShapeV2s, bottomstackgrowthcurvestartarcshapev2_)
					if idx != -1 {
						_bottomstackofgrowthcurvev2.BottomStackGrowthCurveStartArcShapeV2s = slices.Delete(_bottomstackofgrowthcurvev2.BottomStackGrowthCurveStartArcShapeV2s, idx, idx+1)
						bottomstackgrowthcurvestartarcshapev2FormCallback.probe.UpdateSliceOfPointersCallback(_bottomstackofgrowthcurvev2, "BottomStackGrowthCurveStartArcShapeV2s", &_bottomstackofgrowthcurvev2.BottomStackGrowthCurveStartArcShapeV2s)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if bottomstackgrowthcurvestartarcshapev2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		bottomstackgrowthcurvestartarcshapev2_.Unstage(bottomstackgrowthcurvestartarcshapev2FormCallback.probe.stageOfInterest)
	}

	bottomstackgrowthcurvestartarcshapev2FormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.BottomStackGrowthCurveStartArcShapeV2](
		bottomstackgrowthcurvestartarcshapev2FormCallback.probe,
	)

	// display a new form by reset the form stage
	if bottomstackgrowthcurvestartarcshapev2FormCallback.CreationMode || bottomstackgrowthcurvestartarcshapev2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		bottomstackgrowthcurvestartarcshapev2FormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(bottomstackgrowthcurvestartarcshapev2FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BottomStackGrowthCurveStartArcShapeV2FormCallback(
			nil,
			bottomstackgrowthcurvestartarcshapev2FormCallback.probe,
			newFormGroup,
		)
		bottomstackgrowthcurvestartarcshapev2 := new(models.BottomStackGrowthCurveStartArcShapeV2)
		FillUpForm(bottomstackgrowthcurvestartarcshapev2, newFormGroup, bottomstackgrowthcurvestartarcshapev2FormCallback.probe)
		bottomstackgrowthcurvestartarcshapev2FormCallback.probe.formStage.Commit()
	}

	bottomstackgrowthcurvestartarcshapev2FormCallback.probe.ux_tree()
}
func __gong__New__BottomStackOfGrowthCurveV2FormCallback(
	bottomstackofgrowthcurvev2 *models.BottomStackOfGrowthCurveV2,
	probe *Probe,
	formGroup *form.FormGroup,
) (bottomstackofgrowthcurvev2FormCallback *BottomStackOfGrowthCurveV2FormCallback) {
	bottomstackofgrowthcurvev2FormCallback = new(BottomStackOfGrowthCurveV2FormCallback)
	bottomstackofgrowthcurvev2FormCallback.probe = probe
	bottomstackofgrowthcurvev2FormCallback.bottomstackofgrowthcurvev2 = bottomstackofgrowthcurvev2
	bottomstackofgrowthcurvev2FormCallback.formGroup = formGroup

	bottomstackofgrowthcurvev2FormCallback.CreationMode = (bottomstackofgrowthcurvev2 == nil)

	return
}

type BottomStackOfGrowthCurveV2FormCallback struct {
	bottomstackofgrowthcurvev2 *models.BottomStackOfGrowthCurveV2

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (bottomstackofgrowthcurvev2FormCallback *BottomStackOfGrowthCurveV2FormCallback) OnSave() {
	bottomstackofgrowthcurvev2FormCallback.probe.stageOfInterest.Lock()
	defer bottomstackofgrowthcurvev2FormCallback.probe.stageOfInterest.Unlock()

	// log.Println("BottomStackOfGrowthCurveV2FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bottomstackofgrowthcurvev2FormCallback.probe.formStage.Checkout()

	if bottomstackofgrowthcurvev2FormCallback.bottomstackofgrowthcurvev2 == nil {
		bottomstackofgrowthcurvev2FormCallback.bottomstackofgrowthcurvev2 = new(models.BottomStackOfGrowthCurveV2).Stage(bottomstackofgrowthcurvev2FormCallback.probe.stageOfInterest)
	}
	bottomstackofgrowthcurvev2_ := bottomstackofgrowthcurvev2FormCallback.bottomstackofgrowthcurvev2
	_ = bottomstackofgrowthcurvev2_

	for _, formDiv := range bottomstackofgrowthcurvev2FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(bottomstackofgrowthcurvev2_.Name), formDiv)
		case "BottomStackGrowthCurveStartArcShapeV2s":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.BottomStackGrowthCurveStartArcShapeV2](bottomstackofgrowthcurvev2FormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.BottomStackGrowthCurveStartArcShapeV2, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.BottomStackGrowthCurveStartArcShapeV2)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					bottomstackofgrowthcurvev2FormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.BottomStackGrowthCurveStartArcShapeV2](bottomstackofgrowthcurvev2FormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			bottomstackofgrowthcurvev2_.BottomStackGrowthCurveStartArcShapeV2s = instanceSlice
			bottomstackofgrowthcurvev2FormCallback.probe.UpdateSliceOfPointersCallback(bottomstackofgrowthcurvev2_, "BottomStackGrowthCurveStartArcShapeV2s", &bottomstackofgrowthcurvev2_.BottomStackGrowthCurveStartArcShapeV2s)

		case "BottomStackGrowthCurveEndArcShapeV2s":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.BottomStackGrowthCurveEndArcShapeV2](bottomstackofgrowthcurvev2FormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.BottomStackGrowthCurveEndArcShapeV2, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.BottomStackGrowthCurveEndArcShapeV2)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					bottomstackofgrowthcurvev2FormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.BottomStackGrowthCurveEndArcShapeV2](bottomstackofgrowthcurvev2FormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			bottomstackofgrowthcurvev2_.BottomStackGrowthCurveEndArcShapeV2s = instanceSlice
			bottomstackofgrowthcurvev2FormCallback.probe.UpdateSliceOfPointersCallback(bottomstackofgrowthcurvev2_, "BottomStackGrowthCurveEndArcShapeV2s", &bottomstackofgrowthcurvev2_.BottomStackGrowthCurveEndArcShapeV2s)

		}
	}

	// manage the suppress operation
	if bottomstackofgrowthcurvev2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		bottomstackofgrowthcurvev2_.Unstage(bottomstackofgrowthcurvev2FormCallback.probe.stageOfInterest)
	}

	bottomstackofgrowthcurvev2FormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.BottomStackOfGrowthCurveV2](
		bottomstackofgrowthcurvev2FormCallback.probe,
	)

	// display a new form by reset the form stage
	if bottomstackofgrowthcurvev2FormCallback.CreationMode || bottomstackofgrowthcurvev2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		bottomstackofgrowthcurvev2FormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(bottomstackofgrowthcurvev2FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BottomStackOfGrowthCurveV2FormCallback(
			nil,
			bottomstackofgrowthcurvev2FormCallback.probe,
			newFormGroup,
		)
		bottomstackofgrowthcurvev2 := new(models.BottomStackOfGrowthCurveV2)
		FillUpForm(bottomstackofgrowthcurvev2, newFormGroup, bottomstackofgrowthcurvev2FormCallback.probe)
		bottomstackofgrowthcurvev2FormCallback.probe.formStage.Commit()
	}

	bottomstackofgrowthcurvev2FormCallback.probe.ux_tree()
}
func __gong__New__BottomStartArcShapeV2FormCallback(
	bottomstartarcshapev2 *models.BottomStartArcShapeV2,
	probe *Probe,
	formGroup *form.FormGroup,
) (bottomstartarcshapev2FormCallback *BottomStartArcShapeV2FormCallback) {
	bottomstartarcshapev2FormCallback = new(BottomStartArcShapeV2FormCallback)
	bottomstartarcshapev2FormCallback.probe = probe
	bottomstartarcshapev2FormCallback.bottomstartarcshapev2 = bottomstartarcshapev2
	bottomstartarcshapev2FormCallback.formGroup = formGroup

	bottomstartarcshapev2FormCallback.CreationMode = (bottomstartarcshapev2 == nil)

	return
}

type BottomStartArcShapeV2FormCallback struct {
	bottomstartarcshapev2 *models.BottomStartArcShapeV2

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (bottomstartarcshapev2FormCallback *BottomStartArcShapeV2FormCallback) OnSave() {
	bottomstartarcshapev2FormCallback.probe.stageOfInterest.Lock()
	defer bottomstartarcshapev2FormCallback.probe.stageOfInterest.Unlock()

	// log.Println("BottomStartArcShapeV2FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bottomstartarcshapev2FormCallback.probe.formStage.Checkout()

	if bottomstartarcshapev2FormCallback.bottomstartarcshapev2 == nil {
		bottomstartarcshapev2FormCallback.bottomstartarcshapev2 = new(models.BottomStartArcShapeV2).Stage(bottomstartarcshapev2FormCallback.probe.stageOfInterest)
	}
	bottomstartarcshapev2_ := bottomstartarcshapev2FormCallback.bottomstartarcshapev2
	_ = bottomstartarcshapev2_

	for _, formDiv := range bottomstartarcshapev2FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(bottomstartarcshapev2_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(bottomstartarcshapev2_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(bottomstartarcshapev2_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(bottomstartarcshapev2_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(bottomstartarcshapev2_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(bottomstartarcshapev2_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(bottomstartarcshapev2_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(bottomstartarcshapev2_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(bottomstartarcshapev2_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(bottomstartarcshapev2_.RadiusY), formDiv)
		case "BottomStartArcShapeV2Grid:BottomStartArcShapesV2":
			// 1. Decode the AssociationStorage which contains the rowIDs of the BottomStartArcShapeV2Grid instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target BottomStartArcShapeV2Grid instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.BottomStartArcShapeV2Grid](bottomstartarcshapev2FormCallback.probe.stageOfInterest)
			targetBottomStartArcShapeV2GridIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetBottomStartArcShapeV2GridIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all BottomStartArcShapeV2Grid instances and update their BottomStartArcShapesV2 slice
			for _bottomstartarcshapev2grid := range *models.GetGongstructInstancesSetFromPointerType[*models.BottomStartArcShapeV2Grid](bottomstartarcshapev2FormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(bottomstartarcshapev2FormCallback.probe.stageOfInterest, _bottomstartarcshapev2grid)
				
				// if BottomStartArcShapeV2Grid is selected
				if targetBottomStartArcShapeV2GridIDs[id] {
					// ensure bottomstartarcshapev2_ is in _bottomstartarcshapev2grid.BottomStartArcShapesV2
					found := false
					for _, _b := range _bottomstartarcshapev2grid.BottomStartArcShapesV2 {
						if _b == bottomstartarcshapev2_ {
							found = true
							break
						}
					}
					if !found {
						_bottomstartarcshapev2grid.BottomStartArcShapesV2 = append(_bottomstartarcshapev2grid.BottomStartArcShapesV2, bottomstartarcshapev2_)
						bottomstartarcshapev2FormCallback.probe.UpdateSliceOfPointersCallback(_bottomstartarcshapev2grid, "BottomStartArcShapesV2", &_bottomstartarcshapev2grid.BottomStartArcShapesV2)
					}
				} else {
					// ensure bottomstartarcshapev2_ is NOT in _bottomstartarcshapev2grid.BottomStartArcShapesV2
					idx := slices.Index(_bottomstartarcshapev2grid.BottomStartArcShapesV2, bottomstartarcshapev2_)
					if idx != -1 {
						_bottomstartarcshapev2grid.BottomStartArcShapesV2 = slices.Delete(_bottomstartarcshapev2grid.BottomStartArcShapesV2, idx, idx+1)
						bottomstartarcshapev2FormCallback.probe.UpdateSliceOfPointersCallback(_bottomstartarcshapev2grid, "BottomStartArcShapesV2", &_bottomstartarcshapev2grid.BottomStartArcShapesV2)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if bottomstartarcshapev2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		bottomstartarcshapev2_.Unstage(bottomstartarcshapev2FormCallback.probe.stageOfInterest)
	}

	bottomstartarcshapev2FormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.BottomStartArcShapeV2](
		bottomstartarcshapev2FormCallback.probe,
	)

	// display a new form by reset the form stage
	if bottomstartarcshapev2FormCallback.CreationMode || bottomstartarcshapev2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		bottomstartarcshapev2FormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(bottomstartarcshapev2FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BottomStartArcShapeV2FormCallback(
			nil,
			bottomstartarcshapev2FormCallback.probe,
			newFormGroup,
		)
		bottomstartarcshapev2 := new(models.BottomStartArcShapeV2)
		FillUpForm(bottomstartarcshapev2, newFormGroup, bottomstartarcshapev2FormCallback.probe)
		bottomstartarcshapev2FormCallback.probe.formStage.Commit()
	}

	bottomstartarcshapev2FormCallback.probe.ux_tree()
}
func __gong__New__BottomStartArcShapeV2GridFormCallback(
	bottomstartarcshapev2grid *models.BottomStartArcShapeV2Grid,
	probe *Probe,
	formGroup *form.FormGroup,
) (bottomstartarcshapev2gridFormCallback *BottomStartArcShapeV2GridFormCallback) {
	bottomstartarcshapev2gridFormCallback = new(BottomStartArcShapeV2GridFormCallback)
	bottomstartarcshapev2gridFormCallback.probe = probe
	bottomstartarcshapev2gridFormCallback.bottomstartarcshapev2grid = bottomstartarcshapev2grid
	bottomstartarcshapev2gridFormCallback.formGroup = formGroup

	bottomstartarcshapev2gridFormCallback.CreationMode = (bottomstartarcshapev2grid == nil)

	return
}

type BottomStartArcShapeV2GridFormCallback struct {
	bottomstartarcshapev2grid *models.BottomStartArcShapeV2Grid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (bottomstartarcshapev2gridFormCallback *BottomStartArcShapeV2GridFormCallback) OnSave() {
	bottomstartarcshapev2gridFormCallback.probe.stageOfInterest.Lock()
	defer bottomstartarcshapev2gridFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("BottomStartArcShapeV2GridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	bottomstartarcshapev2gridFormCallback.probe.formStage.Checkout()

	if bottomstartarcshapev2gridFormCallback.bottomstartarcshapev2grid == nil {
		bottomstartarcshapev2gridFormCallback.bottomstartarcshapev2grid = new(models.BottomStartArcShapeV2Grid).Stage(bottomstartarcshapev2gridFormCallback.probe.stageOfInterest)
	}
	bottomstartarcshapev2grid_ := bottomstartarcshapev2gridFormCallback.bottomstartarcshapev2grid
	_ = bottomstartarcshapev2grid_

	for _, formDiv := range bottomstartarcshapev2gridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(bottomstartarcshapev2grid_.Name), formDiv)
		case "BottomStartArcShapesV2":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.BottomStartArcShapeV2](bottomstartarcshapev2gridFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.BottomStartArcShapeV2, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.BottomStartArcShapeV2)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					bottomstartarcshapev2gridFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.BottomStartArcShapeV2](bottomstartarcshapev2gridFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			bottomstartarcshapev2grid_.BottomStartArcShapesV2 = instanceSlice
			bottomstartarcshapev2gridFormCallback.probe.UpdateSliceOfPointersCallback(bottomstartarcshapev2grid_, "BottomStartArcShapesV2", &bottomstartarcshapev2grid_.BottomStartArcShapesV2)

		}
	}

	// manage the suppress operation
	if bottomstartarcshapev2gridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bottomstartarcshapev2grid_.Unstage(bottomstartarcshapev2gridFormCallback.probe.stageOfInterest)
	}

	bottomstartarcshapev2gridFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.BottomStartArcShapeV2Grid](
		bottomstartarcshapev2gridFormCallback.probe,
	)

	// display a new form by reset the form stage
	if bottomstartarcshapev2gridFormCallback.CreationMode || bottomstartarcshapev2gridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bottomstartarcshapev2gridFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(bottomstartarcshapev2gridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BottomStartArcShapeV2GridFormCallback(
			nil,
			bottomstartarcshapev2gridFormCallback.probe,
			newFormGroup,
		)
		bottomstartarcshapev2grid := new(models.BottomStartArcShapeV2Grid)
		FillUpForm(bottomstartarcshapev2grid, newFormGroup, bottomstartarcshapev2gridFormCallback.probe)
		bottomstartarcshapev2gridFormCallback.probe.formStage.Commit()
	}

	bottomstartarcshapev2gridFormCallback.probe.ux_tree()
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
func __gong__New__EndArcShapeV2FormCallback(
	endarcshapev2 *models.EndArcShapeV2,
	probe *Probe,
	formGroup *form.FormGroup,
) (endarcshapev2FormCallback *EndArcShapeV2FormCallback) {
	endarcshapev2FormCallback = new(EndArcShapeV2FormCallback)
	endarcshapev2FormCallback.probe = probe
	endarcshapev2FormCallback.endarcshapev2 = endarcshapev2
	endarcshapev2FormCallback.formGroup = formGroup

	endarcshapev2FormCallback.CreationMode = (endarcshapev2 == nil)

	return
}

type EndArcShapeV2FormCallback struct {
	endarcshapev2 *models.EndArcShapeV2

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (endarcshapev2FormCallback *EndArcShapeV2FormCallback) OnSave() {
	endarcshapev2FormCallback.probe.stageOfInterest.Lock()
	defer endarcshapev2FormCallback.probe.stageOfInterest.Unlock()

	// log.Println("EndArcShapeV2FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	endarcshapev2FormCallback.probe.formStage.Checkout()

	if endarcshapev2FormCallback.endarcshapev2 == nil {
		endarcshapev2FormCallback.endarcshapev2 = new(models.EndArcShapeV2).Stage(endarcshapev2FormCallback.probe.stageOfInterest)
	}
	endarcshapev2_ := endarcshapev2FormCallback.endarcshapev2
	_ = endarcshapev2_

	for _, formDiv := range endarcshapev2FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(endarcshapev2_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(endarcshapev2_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(endarcshapev2_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(endarcshapev2_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(endarcshapev2_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(endarcshapev2_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(endarcshapev2_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(endarcshapev2_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(endarcshapev2_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(endarcshapev2_.RadiusY), formDiv)
		case "EndArcShapeV2Grid:EndArcShapesV2":
			// 1. Decode the AssociationStorage which contains the rowIDs of the EndArcShapeV2Grid instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target EndArcShapeV2Grid instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.EndArcShapeV2Grid](endarcshapev2FormCallback.probe.stageOfInterest)
			targetEndArcShapeV2GridIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetEndArcShapeV2GridIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all EndArcShapeV2Grid instances and update their EndArcShapesV2 slice
			for _endarcshapev2grid := range *models.GetGongstructInstancesSetFromPointerType[*models.EndArcShapeV2Grid](endarcshapev2FormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(endarcshapev2FormCallback.probe.stageOfInterest, _endarcshapev2grid)
				
				// if EndArcShapeV2Grid is selected
				if targetEndArcShapeV2GridIDs[id] {
					// ensure endarcshapev2_ is in _endarcshapev2grid.EndArcShapesV2
					found := false
					for _, _b := range _endarcshapev2grid.EndArcShapesV2 {
						if _b == endarcshapev2_ {
							found = true
							break
						}
					}
					if !found {
						_endarcshapev2grid.EndArcShapesV2 = append(_endarcshapev2grid.EndArcShapesV2, endarcshapev2_)
						endarcshapev2FormCallback.probe.UpdateSliceOfPointersCallback(_endarcshapev2grid, "EndArcShapesV2", &_endarcshapev2grid.EndArcShapesV2)
					}
				} else {
					// ensure endarcshapev2_ is NOT in _endarcshapev2grid.EndArcShapesV2
					idx := slices.Index(_endarcshapev2grid.EndArcShapesV2, endarcshapev2_)
					if idx != -1 {
						_endarcshapev2grid.EndArcShapesV2 = slices.Delete(_endarcshapev2grid.EndArcShapesV2, idx, idx+1)
						endarcshapev2FormCallback.probe.UpdateSliceOfPointersCallback(_endarcshapev2grid, "EndArcShapesV2", &_endarcshapev2grid.EndArcShapesV2)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if endarcshapev2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		endarcshapev2_.Unstage(endarcshapev2FormCallback.probe.stageOfInterest)
	}

	endarcshapev2FormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.EndArcShapeV2](
		endarcshapev2FormCallback.probe,
	)

	// display a new form by reset the form stage
	if endarcshapev2FormCallback.CreationMode || endarcshapev2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		endarcshapev2FormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(endarcshapev2FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__EndArcShapeV2FormCallback(
			nil,
			endarcshapev2FormCallback.probe,
			newFormGroup,
		)
		endarcshapev2 := new(models.EndArcShapeV2)
		FillUpForm(endarcshapev2, newFormGroup, endarcshapev2FormCallback.probe)
		endarcshapev2FormCallback.probe.formStage.Commit()
	}

	endarcshapev2FormCallback.probe.ux_tree()
}
func __gong__New__EndArcShapeV2GridFormCallback(
	endarcshapev2grid *models.EndArcShapeV2Grid,
	probe *Probe,
	formGroup *form.FormGroup,
) (endarcshapev2gridFormCallback *EndArcShapeV2GridFormCallback) {
	endarcshapev2gridFormCallback = new(EndArcShapeV2GridFormCallback)
	endarcshapev2gridFormCallback.probe = probe
	endarcshapev2gridFormCallback.endarcshapev2grid = endarcshapev2grid
	endarcshapev2gridFormCallback.formGroup = formGroup

	endarcshapev2gridFormCallback.CreationMode = (endarcshapev2grid == nil)

	return
}

type EndArcShapeV2GridFormCallback struct {
	endarcshapev2grid *models.EndArcShapeV2Grid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (endarcshapev2gridFormCallback *EndArcShapeV2GridFormCallback) OnSave() {
	endarcshapev2gridFormCallback.probe.stageOfInterest.Lock()
	defer endarcshapev2gridFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("EndArcShapeV2GridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	endarcshapev2gridFormCallback.probe.formStage.Checkout()

	if endarcshapev2gridFormCallback.endarcshapev2grid == nil {
		endarcshapev2gridFormCallback.endarcshapev2grid = new(models.EndArcShapeV2Grid).Stage(endarcshapev2gridFormCallback.probe.stageOfInterest)
	}
	endarcshapev2grid_ := endarcshapev2gridFormCallback.endarcshapev2grid
	_ = endarcshapev2grid_

	for _, formDiv := range endarcshapev2gridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(endarcshapev2grid_.Name), formDiv)
		case "EndArcShapesV2":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.EndArcShapeV2](endarcshapev2gridFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.EndArcShapeV2, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.EndArcShapeV2)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					endarcshapev2gridFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.EndArcShapeV2](endarcshapev2gridFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			endarcshapev2grid_.EndArcShapesV2 = instanceSlice
			endarcshapev2gridFormCallback.probe.UpdateSliceOfPointersCallback(endarcshapev2grid_, "EndArcShapesV2", &endarcshapev2grid_.EndArcShapesV2)

		}
	}

	// manage the suppress operation
	if endarcshapev2gridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		endarcshapev2grid_.Unstage(endarcshapev2gridFormCallback.probe.stageOfInterest)
	}

	endarcshapev2gridFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.EndArcShapeV2Grid](
		endarcshapev2gridFormCallback.probe,
	)

	// display a new form by reset the form stage
	if endarcshapev2gridFormCallback.CreationMode || endarcshapev2gridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		endarcshapev2gridFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(endarcshapev2gridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__EndArcShapeV2GridFormCallback(
			nil,
			endarcshapev2gridFormCallback.probe,
			newFormGroup,
		)
		endarcshapev2grid := new(models.EndArcShapeV2Grid)
		FillUpForm(endarcshapev2grid, newFormGroup, endarcshapev2gridFormCallback.probe)
		endarcshapev2gridFormCallback.probe.formStage.Commit()
	}

	endarcshapev2gridFormCallback.probe.ux_tree()
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
		case "StartArcShapeV2Grid":
			FormDivSelectFieldToField(&(plant_.StartArcShapeV2Grid), plantFormCallback.probe.stageOfInterest, formDiv)
		case "TopStartArcShapeV2Grid":
			FormDivSelectFieldToField(&(plant_.TopStartArcShapeV2Grid), plantFormCallback.probe.stageOfInterest, formDiv)
		case "EndArcShapeGrid":
			FormDivSelectFieldToField(&(plant_.EndArcShapeGrid), plantFormCallback.probe.stageOfInterest, formDiv)
		case "EndArcShapeV2Grid":
			FormDivSelectFieldToField(&(plant_.EndArcShapeV2Grid), plantFormCallback.probe.stageOfInterest, formDiv)
		case "TopEndArcShapeV2Grid":
			FormDivSelectFieldToField(&(plant_.TopEndArcShapeV2Grid), plantFormCallback.probe.stageOfInterest, formDiv)
		case "BottomStartArcShapeV2Grid":
			FormDivSelectFieldToField(&(plant_.BottomStartArcShapeV2Grid), plantFormCallback.probe.stageOfInterest, formDiv)
		case "BottomEndArcShapeV2Grid":
			FormDivSelectFieldToField(&(plant_.BottomEndArcShapeV2Grid), plantFormCallback.probe.stageOfInterest, formDiv)
		case "GrowthCurveBezierShapeGrid":
			FormDivSelectFieldToField(&(plant_.GrowthCurveBezierShapeGrid), plantFormCallback.probe.stageOfInterest, formDiv)
		case "StackOfGrowthCurve":
			FormDivSelectFieldToField(&(plant_.StackOfGrowthCurve), plantFormCallback.probe.stageOfInterest, formDiv)
		case "StackOfGrowthCurveV2":
			FormDivSelectFieldToField(&(plant_.StackOfGrowthCurveV2), plantFormCallback.probe.stageOfInterest, formDiv)
		case "TopStackOfGrowthCurveV2":
			FormDivSelectFieldToField(&(plant_.TopStackOfGrowthCurveV2), plantFormCallback.probe.stageOfInterest, formDiv)
		case "BottomStackOfGrowthCurveV2":
			FormDivSelectFieldToField(&(plant_.BottomStackOfGrowthCurveV2), plantFormCallback.probe.stageOfInterest, formDiv)
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
		case "IsHiddenStartArcShapeV2Grid":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenStartArcShapeV2Grid), formDiv)
		case "IsHiddenTopStartArcShapeV2Grid":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenTopStartArcShapeV2Grid), formDiv)
		case "IsHiddenEndArcShapeGrid":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenEndArcShapeGrid), formDiv)
		case "IsHiddenEndArcShapeV2Grid":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenEndArcShapeV2Grid), formDiv)
		case "IsHiddenTopEndArcShapeV2Grid":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenTopEndArcShapeV2Grid), formDiv)
		case "IsHiddenBottomStartArcShapeV2Grid":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenBottomStartArcShapeV2Grid), formDiv)
		case "IsHiddenBottomEndArcShapeV2Grid":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenBottomEndArcShapeV2Grid), formDiv)
		case "IsHiddenGrowthCurveBezierShapeGrid":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenGrowthCurveBezierShapeGrid), formDiv)
		case "IsHiddenStackOfGrowthCurve":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenStackOfGrowthCurve), formDiv)
		case "IsHiddenStackOfGrowthCurveV2":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenStackOfGrowthCurveV2), formDiv)
		case "IsHiddenTopStackOfGrowthCurveV2":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenTopStackOfGrowthCurveV2), formDiv)
		case "IsHiddenBottomStackOfGrowthCurveV2":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenBottomStackOfGrowthCurveV2), formDiv)
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
func __gong__New__StackGrowthCurveBezierShapeFormCallback(
	stackgrowthcurvebeziershape *models.StackGrowthCurveBezierShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackgrowthcurvebeziershapeFormCallback *StackGrowthCurveBezierShapeFormCallback) {
	stackgrowthcurvebeziershapeFormCallback = new(StackGrowthCurveBezierShapeFormCallback)
	stackgrowthcurvebeziershapeFormCallback.probe = probe
	stackgrowthcurvebeziershapeFormCallback.stackgrowthcurvebeziershape = stackgrowthcurvebeziershape
	stackgrowthcurvebeziershapeFormCallback.formGroup = formGroup

	stackgrowthcurvebeziershapeFormCallback.CreationMode = (stackgrowthcurvebeziershape == nil)

	return
}

type StackGrowthCurveBezierShapeFormCallback struct {
	stackgrowthcurvebeziershape *models.StackGrowthCurveBezierShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (stackgrowthcurvebeziershapeFormCallback *StackGrowthCurveBezierShapeFormCallback) OnSave() {
	stackgrowthcurvebeziershapeFormCallback.probe.stageOfInterest.Lock()
	defer stackgrowthcurvebeziershapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StackGrowthCurveBezierShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	stackgrowthcurvebeziershapeFormCallback.probe.formStage.Checkout()

	if stackgrowthcurvebeziershapeFormCallback.stackgrowthcurvebeziershape == nil {
		stackgrowthcurvebeziershapeFormCallback.stackgrowthcurvebeziershape = new(models.StackGrowthCurveBezierShape).Stage(stackgrowthcurvebeziershapeFormCallback.probe.stageOfInterest)
	}
	stackgrowthcurvebeziershape_ := stackgrowthcurvebeziershapeFormCallback.stackgrowthcurvebeziershape
	_ = stackgrowthcurvebeziershape_

	for _, formDiv := range stackgrowthcurvebeziershapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(stackgrowthcurvebeziershape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(stackgrowthcurvebeziershape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(stackgrowthcurvebeziershape_.StartY), formDiv)
		case "ControlPointStartX":
			FormDivBasicFieldToField(&(stackgrowthcurvebeziershape_.ControlPointStartX), formDiv)
		case "ControlPointStartY":
			FormDivBasicFieldToField(&(stackgrowthcurvebeziershape_.ControlPointStartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(stackgrowthcurvebeziershape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(stackgrowthcurvebeziershape_.EndY), formDiv)
		case "ControlPointEndX":
			FormDivBasicFieldToField(&(stackgrowthcurvebeziershape_.ControlPointEndX), formDiv)
		case "ControlPointEndY":
			FormDivBasicFieldToField(&(stackgrowthcurvebeziershape_.ControlPointEndY), formDiv)
		case "StackOfGrowthCurve:StackGrowthCurveBezierShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the StackOfGrowthCurve instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target StackOfGrowthCurve instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.StackOfGrowthCurve](stackgrowthcurvebeziershapeFormCallback.probe.stageOfInterest)
			targetStackOfGrowthCurveIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetStackOfGrowthCurveIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all StackOfGrowthCurve instances and update their StackGrowthCurveBezierShapes slice
			for _stackofgrowthcurve := range *models.GetGongstructInstancesSetFromPointerType[*models.StackOfGrowthCurve](stackgrowthcurvebeziershapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(stackgrowthcurvebeziershapeFormCallback.probe.stageOfInterest, _stackofgrowthcurve)
				
				// if StackOfGrowthCurve is selected
				if targetStackOfGrowthCurveIDs[id] {
					// ensure stackgrowthcurvebeziershape_ is in _stackofgrowthcurve.StackGrowthCurveBezierShapes
					found := false
					for _, _b := range _stackofgrowthcurve.StackGrowthCurveBezierShapes {
						if _b == stackgrowthcurvebeziershape_ {
							found = true
							break
						}
					}
					if !found {
						_stackofgrowthcurve.StackGrowthCurveBezierShapes = append(_stackofgrowthcurve.StackGrowthCurveBezierShapes, stackgrowthcurvebeziershape_)
						stackgrowthcurvebeziershapeFormCallback.probe.UpdateSliceOfPointersCallback(_stackofgrowthcurve, "StackGrowthCurveBezierShapes", &_stackofgrowthcurve.StackGrowthCurveBezierShapes)
					}
				} else {
					// ensure stackgrowthcurvebeziershape_ is NOT in _stackofgrowthcurve.StackGrowthCurveBezierShapes
					idx := slices.Index(_stackofgrowthcurve.StackGrowthCurveBezierShapes, stackgrowthcurvebeziershape_)
					if idx != -1 {
						_stackofgrowthcurve.StackGrowthCurveBezierShapes = slices.Delete(_stackofgrowthcurve.StackGrowthCurveBezierShapes, idx, idx+1)
						stackgrowthcurvebeziershapeFormCallback.probe.UpdateSliceOfPointersCallback(_stackofgrowthcurve, "StackGrowthCurveBezierShapes", &_stackofgrowthcurve.StackGrowthCurveBezierShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if stackgrowthcurvebeziershapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackgrowthcurvebeziershape_.Unstage(stackgrowthcurvebeziershapeFormCallback.probe.stageOfInterest)
	}

	stackgrowthcurvebeziershapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StackGrowthCurveBezierShape](
		stackgrowthcurvebeziershapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if stackgrowthcurvebeziershapeFormCallback.CreationMode || stackgrowthcurvebeziershapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackgrowthcurvebeziershapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(stackgrowthcurvebeziershapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StackGrowthCurveBezierShapeFormCallback(
			nil,
			stackgrowthcurvebeziershapeFormCallback.probe,
			newFormGroup,
		)
		stackgrowthcurvebeziershape := new(models.StackGrowthCurveBezierShape)
		FillUpForm(stackgrowthcurvebeziershape, newFormGroup, stackgrowthcurvebeziershapeFormCallback.probe)
		stackgrowthcurvebeziershapeFormCallback.probe.formStage.Commit()
	}

	stackgrowthcurvebeziershapeFormCallback.probe.ux_tree()
}
func __gong__New__StackGrowthCurveEndArcShapeV2FormCallback(
	stackgrowthcurveendarcshapev2 *models.StackGrowthCurveEndArcShapeV2,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackgrowthcurveendarcshapev2FormCallback *StackGrowthCurveEndArcShapeV2FormCallback) {
	stackgrowthcurveendarcshapev2FormCallback = new(StackGrowthCurveEndArcShapeV2FormCallback)
	stackgrowthcurveendarcshapev2FormCallback.probe = probe
	stackgrowthcurveendarcshapev2FormCallback.stackgrowthcurveendarcshapev2 = stackgrowthcurveendarcshapev2
	stackgrowthcurveendarcshapev2FormCallback.formGroup = formGroup

	stackgrowthcurveendarcshapev2FormCallback.CreationMode = (stackgrowthcurveendarcshapev2 == nil)

	return
}

type StackGrowthCurveEndArcShapeV2FormCallback struct {
	stackgrowthcurveendarcshapev2 *models.StackGrowthCurveEndArcShapeV2

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (stackgrowthcurveendarcshapev2FormCallback *StackGrowthCurveEndArcShapeV2FormCallback) OnSave() {
	stackgrowthcurveendarcshapev2FormCallback.probe.stageOfInterest.Lock()
	defer stackgrowthcurveendarcshapev2FormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StackGrowthCurveEndArcShapeV2FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	stackgrowthcurveendarcshapev2FormCallback.probe.formStage.Checkout()

	if stackgrowthcurveendarcshapev2FormCallback.stackgrowthcurveendarcshapev2 == nil {
		stackgrowthcurveendarcshapev2FormCallback.stackgrowthcurveendarcshapev2 = new(models.StackGrowthCurveEndArcShapeV2).Stage(stackgrowthcurveendarcshapev2FormCallback.probe.stageOfInterest)
	}
	stackgrowthcurveendarcshapev2_ := stackgrowthcurveendarcshapev2FormCallback.stackgrowthcurveendarcshapev2
	_ = stackgrowthcurveendarcshapev2_

	for _, formDiv := range stackgrowthcurveendarcshapev2FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(stackgrowthcurveendarcshapev2_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(stackgrowthcurveendarcshapev2_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(stackgrowthcurveendarcshapev2_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(stackgrowthcurveendarcshapev2_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(stackgrowthcurveendarcshapev2_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(stackgrowthcurveendarcshapev2_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(stackgrowthcurveendarcshapev2_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(stackgrowthcurveendarcshapev2_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(stackgrowthcurveendarcshapev2_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(stackgrowthcurveendarcshapev2_.RadiusY), formDiv)
		case "StackOfGrowthCurveV2:StackGrowthCurveEndArcShapeV2s":
			// 1. Decode the AssociationStorage which contains the rowIDs of the StackOfGrowthCurveV2 instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target StackOfGrowthCurveV2 instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.StackOfGrowthCurveV2](stackgrowthcurveendarcshapev2FormCallback.probe.stageOfInterest)
			targetStackOfGrowthCurveV2IDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetStackOfGrowthCurveV2IDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all StackOfGrowthCurveV2 instances and update their StackGrowthCurveEndArcShapeV2s slice
			for _stackofgrowthcurvev2 := range *models.GetGongstructInstancesSetFromPointerType[*models.StackOfGrowthCurveV2](stackgrowthcurveendarcshapev2FormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(stackgrowthcurveendarcshapev2FormCallback.probe.stageOfInterest, _stackofgrowthcurvev2)
				
				// if StackOfGrowthCurveV2 is selected
				if targetStackOfGrowthCurveV2IDs[id] {
					// ensure stackgrowthcurveendarcshapev2_ is in _stackofgrowthcurvev2.StackGrowthCurveEndArcShapeV2s
					found := false
					for _, _b := range _stackofgrowthcurvev2.StackGrowthCurveEndArcShapeV2s {
						if _b == stackgrowthcurveendarcshapev2_ {
							found = true
							break
						}
					}
					if !found {
						_stackofgrowthcurvev2.StackGrowthCurveEndArcShapeV2s = append(_stackofgrowthcurvev2.StackGrowthCurveEndArcShapeV2s, stackgrowthcurveendarcshapev2_)
						stackgrowthcurveendarcshapev2FormCallback.probe.UpdateSliceOfPointersCallback(_stackofgrowthcurvev2, "StackGrowthCurveEndArcShapeV2s", &_stackofgrowthcurvev2.StackGrowthCurveEndArcShapeV2s)
					}
				} else {
					// ensure stackgrowthcurveendarcshapev2_ is NOT in _stackofgrowthcurvev2.StackGrowthCurveEndArcShapeV2s
					idx := slices.Index(_stackofgrowthcurvev2.StackGrowthCurveEndArcShapeV2s, stackgrowthcurveendarcshapev2_)
					if idx != -1 {
						_stackofgrowthcurvev2.StackGrowthCurveEndArcShapeV2s = slices.Delete(_stackofgrowthcurvev2.StackGrowthCurveEndArcShapeV2s, idx, idx+1)
						stackgrowthcurveendarcshapev2FormCallback.probe.UpdateSliceOfPointersCallback(_stackofgrowthcurvev2, "StackGrowthCurveEndArcShapeV2s", &_stackofgrowthcurvev2.StackGrowthCurveEndArcShapeV2s)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if stackgrowthcurveendarcshapev2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackgrowthcurveendarcshapev2_.Unstage(stackgrowthcurveendarcshapev2FormCallback.probe.stageOfInterest)
	}

	stackgrowthcurveendarcshapev2FormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StackGrowthCurveEndArcShapeV2](
		stackgrowthcurveendarcshapev2FormCallback.probe,
	)

	// display a new form by reset the form stage
	if stackgrowthcurveendarcshapev2FormCallback.CreationMode || stackgrowthcurveendarcshapev2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackgrowthcurveendarcshapev2FormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(stackgrowthcurveendarcshapev2FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StackGrowthCurveEndArcShapeV2FormCallback(
			nil,
			stackgrowthcurveendarcshapev2FormCallback.probe,
			newFormGroup,
		)
		stackgrowthcurveendarcshapev2 := new(models.StackGrowthCurveEndArcShapeV2)
		FillUpForm(stackgrowthcurveendarcshapev2, newFormGroup, stackgrowthcurveendarcshapev2FormCallback.probe)
		stackgrowthcurveendarcshapev2FormCallback.probe.formStage.Commit()
	}

	stackgrowthcurveendarcshapev2FormCallback.probe.ux_tree()
}
func __gong__New__StackGrowthCurveStartArcShapeV2FormCallback(
	stackgrowthcurvestartarcshapev2 *models.StackGrowthCurveStartArcShapeV2,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackgrowthcurvestartarcshapev2FormCallback *StackGrowthCurveStartArcShapeV2FormCallback) {
	stackgrowthcurvestartarcshapev2FormCallback = new(StackGrowthCurveStartArcShapeV2FormCallback)
	stackgrowthcurvestartarcshapev2FormCallback.probe = probe
	stackgrowthcurvestartarcshapev2FormCallback.stackgrowthcurvestartarcshapev2 = stackgrowthcurvestartarcshapev2
	stackgrowthcurvestartarcshapev2FormCallback.formGroup = formGroup

	stackgrowthcurvestartarcshapev2FormCallback.CreationMode = (stackgrowthcurvestartarcshapev2 == nil)

	return
}

type StackGrowthCurveStartArcShapeV2FormCallback struct {
	stackgrowthcurvestartarcshapev2 *models.StackGrowthCurveStartArcShapeV2

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (stackgrowthcurvestartarcshapev2FormCallback *StackGrowthCurveStartArcShapeV2FormCallback) OnSave() {
	stackgrowthcurvestartarcshapev2FormCallback.probe.stageOfInterest.Lock()
	defer stackgrowthcurvestartarcshapev2FormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StackGrowthCurveStartArcShapeV2FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	stackgrowthcurvestartarcshapev2FormCallback.probe.formStage.Checkout()

	if stackgrowthcurvestartarcshapev2FormCallback.stackgrowthcurvestartarcshapev2 == nil {
		stackgrowthcurvestartarcshapev2FormCallback.stackgrowthcurvestartarcshapev2 = new(models.StackGrowthCurveStartArcShapeV2).Stage(stackgrowthcurvestartarcshapev2FormCallback.probe.stageOfInterest)
	}
	stackgrowthcurvestartarcshapev2_ := stackgrowthcurvestartarcshapev2FormCallback.stackgrowthcurvestartarcshapev2
	_ = stackgrowthcurvestartarcshapev2_

	for _, formDiv := range stackgrowthcurvestartarcshapev2FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(stackgrowthcurvestartarcshapev2_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(stackgrowthcurvestartarcshapev2_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(stackgrowthcurvestartarcshapev2_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(stackgrowthcurvestartarcshapev2_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(stackgrowthcurvestartarcshapev2_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(stackgrowthcurvestartarcshapev2_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(stackgrowthcurvestartarcshapev2_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(stackgrowthcurvestartarcshapev2_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(stackgrowthcurvestartarcshapev2_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(stackgrowthcurvestartarcshapev2_.RadiusY), formDiv)
		case "StackOfGrowthCurveV2:StackGrowthCurveStartArcShapeV2s":
			// 1. Decode the AssociationStorage which contains the rowIDs of the StackOfGrowthCurveV2 instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target StackOfGrowthCurveV2 instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.StackOfGrowthCurveV2](stackgrowthcurvestartarcshapev2FormCallback.probe.stageOfInterest)
			targetStackOfGrowthCurveV2IDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetStackOfGrowthCurveV2IDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all StackOfGrowthCurveV2 instances and update their StackGrowthCurveStartArcShapeV2s slice
			for _stackofgrowthcurvev2 := range *models.GetGongstructInstancesSetFromPointerType[*models.StackOfGrowthCurveV2](stackgrowthcurvestartarcshapev2FormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(stackgrowthcurvestartarcshapev2FormCallback.probe.stageOfInterest, _stackofgrowthcurvev2)
				
				// if StackOfGrowthCurveV2 is selected
				if targetStackOfGrowthCurveV2IDs[id] {
					// ensure stackgrowthcurvestartarcshapev2_ is in _stackofgrowthcurvev2.StackGrowthCurveStartArcShapeV2s
					found := false
					for _, _b := range _stackofgrowthcurvev2.StackGrowthCurveStartArcShapeV2s {
						if _b == stackgrowthcurvestartarcshapev2_ {
							found = true
							break
						}
					}
					if !found {
						_stackofgrowthcurvev2.StackGrowthCurveStartArcShapeV2s = append(_stackofgrowthcurvev2.StackGrowthCurveStartArcShapeV2s, stackgrowthcurvestartarcshapev2_)
						stackgrowthcurvestartarcshapev2FormCallback.probe.UpdateSliceOfPointersCallback(_stackofgrowthcurvev2, "StackGrowthCurveStartArcShapeV2s", &_stackofgrowthcurvev2.StackGrowthCurveStartArcShapeV2s)
					}
				} else {
					// ensure stackgrowthcurvestartarcshapev2_ is NOT in _stackofgrowthcurvev2.StackGrowthCurveStartArcShapeV2s
					idx := slices.Index(_stackofgrowthcurvev2.StackGrowthCurveStartArcShapeV2s, stackgrowthcurvestartarcshapev2_)
					if idx != -1 {
						_stackofgrowthcurvev2.StackGrowthCurveStartArcShapeV2s = slices.Delete(_stackofgrowthcurvev2.StackGrowthCurveStartArcShapeV2s, idx, idx+1)
						stackgrowthcurvestartarcshapev2FormCallback.probe.UpdateSliceOfPointersCallback(_stackofgrowthcurvev2, "StackGrowthCurveStartArcShapeV2s", &_stackofgrowthcurvev2.StackGrowthCurveStartArcShapeV2s)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if stackgrowthcurvestartarcshapev2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackgrowthcurvestartarcshapev2_.Unstage(stackgrowthcurvestartarcshapev2FormCallback.probe.stageOfInterest)
	}

	stackgrowthcurvestartarcshapev2FormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StackGrowthCurveStartArcShapeV2](
		stackgrowthcurvestartarcshapev2FormCallback.probe,
	)

	// display a new form by reset the form stage
	if stackgrowthcurvestartarcshapev2FormCallback.CreationMode || stackgrowthcurvestartarcshapev2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackgrowthcurvestartarcshapev2FormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(stackgrowthcurvestartarcshapev2FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StackGrowthCurveStartArcShapeV2FormCallback(
			nil,
			stackgrowthcurvestartarcshapev2FormCallback.probe,
			newFormGroup,
		)
		stackgrowthcurvestartarcshapev2 := new(models.StackGrowthCurveStartArcShapeV2)
		FillUpForm(stackgrowthcurvestartarcshapev2, newFormGroup, stackgrowthcurvestartarcshapev2FormCallback.probe)
		stackgrowthcurvestartarcshapev2FormCallback.probe.formStage.Commit()
	}

	stackgrowthcurvestartarcshapev2FormCallback.probe.ux_tree()
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
		case "StackGrowthCurveBezierShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.StackGrowthCurveBezierShape](stackofgrowthcurveFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.StackGrowthCurveBezierShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.StackGrowthCurveBezierShape)

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
			map_RowID_ID := GetMap_RowID_ID[*models.StackGrowthCurveBezierShape](stackofgrowthcurveFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			stackofgrowthcurve_.StackGrowthCurveBezierShapes = instanceSlice
			stackofgrowthcurveFormCallback.probe.UpdateSliceOfPointersCallback(stackofgrowthcurve_, "StackGrowthCurveBezierShapes", &stackofgrowthcurve_.StackGrowthCurveBezierShapes)

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
func __gong__New__StackOfGrowthCurveV2FormCallback(
	stackofgrowthcurvev2 *models.StackOfGrowthCurveV2,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackofgrowthcurvev2FormCallback *StackOfGrowthCurveV2FormCallback) {
	stackofgrowthcurvev2FormCallback = new(StackOfGrowthCurveV2FormCallback)
	stackofgrowthcurvev2FormCallback.probe = probe
	stackofgrowthcurvev2FormCallback.stackofgrowthcurvev2 = stackofgrowthcurvev2
	stackofgrowthcurvev2FormCallback.formGroup = formGroup

	stackofgrowthcurvev2FormCallback.CreationMode = (stackofgrowthcurvev2 == nil)

	return
}

type StackOfGrowthCurveV2FormCallback struct {
	stackofgrowthcurvev2 *models.StackOfGrowthCurveV2

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (stackofgrowthcurvev2FormCallback *StackOfGrowthCurveV2FormCallback) OnSave() {
	stackofgrowthcurvev2FormCallback.probe.stageOfInterest.Lock()
	defer stackofgrowthcurvev2FormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StackOfGrowthCurveV2FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	stackofgrowthcurvev2FormCallback.probe.formStage.Checkout()

	if stackofgrowthcurvev2FormCallback.stackofgrowthcurvev2 == nil {
		stackofgrowthcurvev2FormCallback.stackofgrowthcurvev2 = new(models.StackOfGrowthCurveV2).Stage(stackofgrowthcurvev2FormCallback.probe.stageOfInterest)
	}
	stackofgrowthcurvev2_ := stackofgrowthcurvev2FormCallback.stackofgrowthcurvev2
	_ = stackofgrowthcurvev2_

	for _, formDiv := range stackofgrowthcurvev2FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(stackofgrowthcurvev2_.Name), formDiv)
		case "StackGrowthCurveStartArcShapeV2s":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.StackGrowthCurveStartArcShapeV2](stackofgrowthcurvev2FormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.StackGrowthCurveStartArcShapeV2, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.StackGrowthCurveStartArcShapeV2)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					stackofgrowthcurvev2FormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.StackGrowthCurveStartArcShapeV2](stackofgrowthcurvev2FormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			stackofgrowthcurvev2_.StackGrowthCurveStartArcShapeV2s = instanceSlice
			stackofgrowthcurvev2FormCallback.probe.UpdateSliceOfPointersCallback(stackofgrowthcurvev2_, "StackGrowthCurveStartArcShapeV2s", &stackofgrowthcurvev2_.StackGrowthCurveStartArcShapeV2s)

		case "StackGrowthCurveEndArcShapeV2s":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.StackGrowthCurveEndArcShapeV2](stackofgrowthcurvev2FormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.StackGrowthCurveEndArcShapeV2, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.StackGrowthCurveEndArcShapeV2)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					stackofgrowthcurvev2FormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.StackGrowthCurveEndArcShapeV2](stackofgrowthcurvev2FormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			stackofgrowthcurvev2_.StackGrowthCurveEndArcShapeV2s = instanceSlice
			stackofgrowthcurvev2FormCallback.probe.UpdateSliceOfPointersCallback(stackofgrowthcurvev2_, "StackGrowthCurveEndArcShapeV2s", &stackofgrowthcurvev2_.StackGrowthCurveEndArcShapeV2s)

		}
	}

	// manage the suppress operation
	if stackofgrowthcurvev2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackofgrowthcurvev2_.Unstage(stackofgrowthcurvev2FormCallback.probe.stageOfInterest)
	}

	stackofgrowthcurvev2FormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StackOfGrowthCurveV2](
		stackofgrowthcurvev2FormCallback.probe,
	)

	// display a new form by reset the form stage
	if stackofgrowthcurvev2FormCallback.CreationMode || stackofgrowthcurvev2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackofgrowthcurvev2FormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(stackofgrowthcurvev2FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StackOfGrowthCurveV2FormCallback(
			nil,
			stackofgrowthcurvev2FormCallback.probe,
			newFormGroup,
		)
		stackofgrowthcurvev2 := new(models.StackOfGrowthCurveV2)
		FillUpForm(stackofgrowthcurvev2, newFormGroup, stackofgrowthcurvev2FormCallback.probe)
		stackofgrowthcurvev2FormCallback.probe.formStage.Commit()
	}

	stackofgrowthcurvev2FormCallback.probe.ux_tree()
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
func __gong__New__StartArcShapeV2FormCallback(
	startarcshapev2 *models.StartArcShapeV2,
	probe *Probe,
	formGroup *form.FormGroup,
) (startarcshapev2FormCallback *StartArcShapeV2FormCallback) {
	startarcshapev2FormCallback = new(StartArcShapeV2FormCallback)
	startarcshapev2FormCallback.probe = probe
	startarcshapev2FormCallback.startarcshapev2 = startarcshapev2
	startarcshapev2FormCallback.formGroup = formGroup

	startarcshapev2FormCallback.CreationMode = (startarcshapev2 == nil)

	return
}

type StartArcShapeV2FormCallback struct {
	startarcshapev2 *models.StartArcShapeV2

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (startarcshapev2FormCallback *StartArcShapeV2FormCallback) OnSave() {
	startarcshapev2FormCallback.probe.stageOfInterest.Lock()
	defer startarcshapev2FormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StartArcShapeV2FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	startarcshapev2FormCallback.probe.formStage.Checkout()

	if startarcshapev2FormCallback.startarcshapev2 == nil {
		startarcshapev2FormCallback.startarcshapev2 = new(models.StartArcShapeV2).Stage(startarcshapev2FormCallback.probe.stageOfInterest)
	}
	startarcshapev2_ := startarcshapev2FormCallback.startarcshapev2
	_ = startarcshapev2_

	for _, formDiv := range startarcshapev2FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(startarcshapev2_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(startarcshapev2_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(startarcshapev2_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(startarcshapev2_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(startarcshapev2_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(startarcshapev2_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(startarcshapev2_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(startarcshapev2_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(startarcshapev2_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(startarcshapev2_.RadiusY), formDiv)
		case "StartArcShapeV2Grid:StartArcShapesV2":
			// 1. Decode the AssociationStorage which contains the rowIDs of the StartArcShapeV2Grid instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target StartArcShapeV2Grid instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.StartArcShapeV2Grid](startarcshapev2FormCallback.probe.stageOfInterest)
			targetStartArcShapeV2GridIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetStartArcShapeV2GridIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all StartArcShapeV2Grid instances and update their StartArcShapesV2 slice
			for _startarcshapev2grid := range *models.GetGongstructInstancesSetFromPointerType[*models.StartArcShapeV2Grid](startarcshapev2FormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(startarcshapev2FormCallback.probe.stageOfInterest, _startarcshapev2grid)
				
				// if StartArcShapeV2Grid is selected
				if targetStartArcShapeV2GridIDs[id] {
					// ensure startarcshapev2_ is in _startarcshapev2grid.StartArcShapesV2
					found := false
					for _, _b := range _startarcshapev2grid.StartArcShapesV2 {
						if _b == startarcshapev2_ {
							found = true
							break
						}
					}
					if !found {
						_startarcshapev2grid.StartArcShapesV2 = append(_startarcshapev2grid.StartArcShapesV2, startarcshapev2_)
						startarcshapev2FormCallback.probe.UpdateSliceOfPointersCallback(_startarcshapev2grid, "StartArcShapesV2", &_startarcshapev2grid.StartArcShapesV2)
					}
				} else {
					// ensure startarcshapev2_ is NOT in _startarcshapev2grid.StartArcShapesV2
					idx := slices.Index(_startarcshapev2grid.StartArcShapesV2, startarcshapev2_)
					if idx != -1 {
						_startarcshapev2grid.StartArcShapesV2 = slices.Delete(_startarcshapev2grid.StartArcShapesV2, idx, idx+1)
						startarcshapev2FormCallback.probe.UpdateSliceOfPointersCallback(_startarcshapev2grid, "StartArcShapesV2", &_startarcshapev2grid.StartArcShapesV2)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if startarcshapev2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		startarcshapev2_.Unstage(startarcshapev2FormCallback.probe.stageOfInterest)
	}

	startarcshapev2FormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StartArcShapeV2](
		startarcshapev2FormCallback.probe,
	)

	// display a new form by reset the form stage
	if startarcshapev2FormCallback.CreationMode || startarcshapev2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		startarcshapev2FormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(startarcshapev2FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StartArcShapeV2FormCallback(
			nil,
			startarcshapev2FormCallback.probe,
			newFormGroup,
		)
		startarcshapev2 := new(models.StartArcShapeV2)
		FillUpForm(startarcshapev2, newFormGroup, startarcshapev2FormCallback.probe)
		startarcshapev2FormCallback.probe.formStage.Commit()
	}

	startarcshapev2FormCallback.probe.ux_tree()
}
func __gong__New__StartArcShapeV2GridFormCallback(
	startarcshapev2grid *models.StartArcShapeV2Grid,
	probe *Probe,
	formGroup *form.FormGroup,
) (startarcshapev2gridFormCallback *StartArcShapeV2GridFormCallback) {
	startarcshapev2gridFormCallback = new(StartArcShapeV2GridFormCallback)
	startarcshapev2gridFormCallback.probe = probe
	startarcshapev2gridFormCallback.startarcshapev2grid = startarcshapev2grid
	startarcshapev2gridFormCallback.formGroup = formGroup

	startarcshapev2gridFormCallback.CreationMode = (startarcshapev2grid == nil)

	return
}

type StartArcShapeV2GridFormCallback struct {
	startarcshapev2grid *models.StartArcShapeV2Grid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (startarcshapev2gridFormCallback *StartArcShapeV2GridFormCallback) OnSave() {
	startarcshapev2gridFormCallback.probe.stageOfInterest.Lock()
	defer startarcshapev2gridFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StartArcShapeV2GridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	startarcshapev2gridFormCallback.probe.formStage.Checkout()

	if startarcshapev2gridFormCallback.startarcshapev2grid == nil {
		startarcshapev2gridFormCallback.startarcshapev2grid = new(models.StartArcShapeV2Grid).Stage(startarcshapev2gridFormCallback.probe.stageOfInterest)
	}
	startarcshapev2grid_ := startarcshapev2gridFormCallback.startarcshapev2grid
	_ = startarcshapev2grid_

	for _, formDiv := range startarcshapev2gridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(startarcshapev2grid_.Name), formDiv)
		case "StartArcShapesV2":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.StartArcShapeV2](startarcshapev2gridFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.StartArcShapeV2, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.StartArcShapeV2)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					startarcshapev2gridFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.StartArcShapeV2](startarcshapev2gridFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			startarcshapev2grid_.StartArcShapesV2 = instanceSlice
			startarcshapev2gridFormCallback.probe.UpdateSliceOfPointersCallback(startarcshapev2grid_, "StartArcShapesV2", &startarcshapev2grid_.StartArcShapesV2)

		}
	}

	// manage the suppress operation
	if startarcshapev2gridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		startarcshapev2grid_.Unstage(startarcshapev2gridFormCallback.probe.stageOfInterest)
	}

	startarcshapev2gridFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StartArcShapeV2Grid](
		startarcshapev2gridFormCallback.probe,
	)

	// display a new form by reset the form stage
	if startarcshapev2gridFormCallback.CreationMode || startarcshapev2gridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		startarcshapev2gridFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(startarcshapev2gridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StartArcShapeV2GridFormCallback(
			nil,
			startarcshapev2gridFormCallback.probe,
			newFormGroup,
		)
		startarcshapev2grid := new(models.StartArcShapeV2Grid)
		FillUpForm(startarcshapev2grid, newFormGroup, startarcshapev2gridFormCallback.probe)
		startarcshapev2gridFormCallback.probe.formStage.Commit()
	}

	startarcshapev2gridFormCallback.probe.ux_tree()
}
func __gong__New__TopEndArcShapeV2FormCallback(
	topendarcshapev2 *models.TopEndArcShapeV2,
	probe *Probe,
	formGroup *form.FormGroup,
) (topendarcshapev2FormCallback *TopEndArcShapeV2FormCallback) {
	topendarcshapev2FormCallback = new(TopEndArcShapeV2FormCallback)
	topendarcshapev2FormCallback.probe = probe
	topendarcshapev2FormCallback.topendarcshapev2 = topendarcshapev2
	topendarcshapev2FormCallback.formGroup = formGroup

	topendarcshapev2FormCallback.CreationMode = (topendarcshapev2 == nil)

	return
}

type TopEndArcShapeV2FormCallback struct {
	topendarcshapev2 *models.TopEndArcShapeV2

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (topendarcshapev2FormCallback *TopEndArcShapeV2FormCallback) OnSave() {
	topendarcshapev2FormCallback.probe.stageOfInterest.Lock()
	defer topendarcshapev2FormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TopEndArcShapeV2FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	topendarcshapev2FormCallback.probe.formStage.Checkout()

	if topendarcshapev2FormCallback.topendarcshapev2 == nil {
		topendarcshapev2FormCallback.topendarcshapev2 = new(models.TopEndArcShapeV2).Stage(topendarcshapev2FormCallback.probe.stageOfInterest)
	}
	topendarcshapev2_ := topendarcshapev2FormCallback.topendarcshapev2
	_ = topendarcshapev2_

	for _, formDiv := range topendarcshapev2FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(topendarcshapev2_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(topendarcshapev2_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(topendarcshapev2_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(topendarcshapev2_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(topendarcshapev2_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(topendarcshapev2_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(topendarcshapev2_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(topendarcshapev2_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(topendarcshapev2_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(topendarcshapev2_.RadiusY), formDiv)
		case "TopEndArcShapeV2Grid:TopEndArcShapesV2":
			// 1. Decode the AssociationStorage which contains the rowIDs of the TopEndArcShapeV2Grid instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target TopEndArcShapeV2Grid instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.TopEndArcShapeV2Grid](topendarcshapev2FormCallback.probe.stageOfInterest)
			targetTopEndArcShapeV2GridIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetTopEndArcShapeV2GridIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all TopEndArcShapeV2Grid instances and update their TopEndArcShapesV2 slice
			for _topendarcshapev2grid := range *models.GetGongstructInstancesSetFromPointerType[*models.TopEndArcShapeV2Grid](topendarcshapev2FormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(topendarcshapev2FormCallback.probe.stageOfInterest, _topendarcshapev2grid)
				
				// if TopEndArcShapeV2Grid is selected
				if targetTopEndArcShapeV2GridIDs[id] {
					// ensure topendarcshapev2_ is in _topendarcshapev2grid.TopEndArcShapesV2
					found := false
					for _, _b := range _topendarcshapev2grid.TopEndArcShapesV2 {
						if _b == topendarcshapev2_ {
							found = true
							break
						}
					}
					if !found {
						_topendarcshapev2grid.TopEndArcShapesV2 = append(_topendarcshapev2grid.TopEndArcShapesV2, topendarcshapev2_)
						topendarcshapev2FormCallback.probe.UpdateSliceOfPointersCallback(_topendarcshapev2grid, "TopEndArcShapesV2", &_topendarcshapev2grid.TopEndArcShapesV2)
					}
				} else {
					// ensure topendarcshapev2_ is NOT in _topendarcshapev2grid.TopEndArcShapesV2
					idx := slices.Index(_topendarcshapev2grid.TopEndArcShapesV2, topendarcshapev2_)
					if idx != -1 {
						_topendarcshapev2grid.TopEndArcShapesV2 = slices.Delete(_topendarcshapev2grid.TopEndArcShapesV2, idx, idx+1)
						topendarcshapev2FormCallback.probe.UpdateSliceOfPointersCallback(_topendarcshapev2grid, "TopEndArcShapesV2", &_topendarcshapev2grid.TopEndArcShapesV2)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if topendarcshapev2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		topendarcshapev2_.Unstage(topendarcshapev2FormCallback.probe.stageOfInterest)
	}

	topendarcshapev2FormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TopEndArcShapeV2](
		topendarcshapev2FormCallback.probe,
	)

	// display a new form by reset the form stage
	if topendarcshapev2FormCallback.CreationMode || topendarcshapev2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		topendarcshapev2FormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(topendarcshapev2FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TopEndArcShapeV2FormCallback(
			nil,
			topendarcshapev2FormCallback.probe,
			newFormGroup,
		)
		topendarcshapev2 := new(models.TopEndArcShapeV2)
		FillUpForm(topendarcshapev2, newFormGroup, topendarcshapev2FormCallback.probe)
		topendarcshapev2FormCallback.probe.formStage.Commit()
	}

	topendarcshapev2FormCallback.probe.ux_tree()
}
func __gong__New__TopEndArcShapeV2GridFormCallback(
	topendarcshapev2grid *models.TopEndArcShapeV2Grid,
	probe *Probe,
	formGroup *form.FormGroup,
) (topendarcshapev2gridFormCallback *TopEndArcShapeV2GridFormCallback) {
	topendarcshapev2gridFormCallback = new(TopEndArcShapeV2GridFormCallback)
	topendarcshapev2gridFormCallback.probe = probe
	topendarcshapev2gridFormCallback.topendarcshapev2grid = topendarcshapev2grid
	topendarcshapev2gridFormCallback.formGroup = formGroup

	topendarcshapev2gridFormCallback.CreationMode = (topendarcshapev2grid == nil)

	return
}

type TopEndArcShapeV2GridFormCallback struct {
	topendarcshapev2grid *models.TopEndArcShapeV2Grid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (topendarcshapev2gridFormCallback *TopEndArcShapeV2GridFormCallback) OnSave() {
	topendarcshapev2gridFormCallback.probe.stageOfInterest.Lock()
	defer topendarcshapev2gridFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TopEndArcShapeV2GridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	topendarcshapev2gridFormCallback.probe.formStage.Checkout()

	if topendarcshapev2gridFormCallback.topendarcshapev2grid == nil {
		topendarcshapev2gridFormCallback.topendarcshapev2grid = new(models.TopEndArcShapeV2Grid).Stage(topendarcshapev2gridFormCallback.probe.stageOfInterest)
	}
	topendarcshapev2grid_ := topendarcshapev2gridFormCallback.topendarcshapev2grid
	_ = topendarcshapev2grid_

	for _, formDiv := range topendarcshapev2gridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(topendarcshapev2grid_.Name), formDiv)
		case "TopEndArcShapesV2":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.TopEndArcShapeV2](topendarcshapev2gridFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.TopEndArcShapeV2, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.TopEndArcShapeV2)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					topendarcshapev2gridFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.TopEndArcShapeV2](topendarcshapev2gridFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			topendarcshapev2grid_.TopEndArcShapesV2 = instanceSlice
			topendarcshapev2gridFormCallback.probe.UpdateSliceOfPointersCallback(topendarcshapev2grid_, "TopEndArcShapesV2", &topendarcshapev2grid_.TopEndArcShapesV2)

		}
	}

	// manage the suppress operation
	if topendarcshapev2gridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topendarcshapev2grid_.Unstage(topendarcshapev2gridFormCallback.probe.stageOfInterest)
	}

	topendarcshapev2gridFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TopEndArcShapeV2Grid](
		topendarcshapev2gridFormCallback.probe,
	)

	// display a new form by reset the form stage
	if topendarcshapev2gridFormCallback.CreationMode || topendarcshapev2gridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topendarcshapev2gridFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(topendarcshapev2gridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TopEndArcShapeV2GridFormCallback(
			nil,
			topendarcshapev2gridFormCallback.probe,
			newFormGroup,
		)
		topendarcshapev2grid := new(models.TopEndArcShapeV2Grid)
		FillUpForm(topendarcshapev2grid, newFormGroup, topendarcshapev2gridFormCallback.probe)
		topendarcshapev2gridFormCallback.probe.formStage.Commit()
	}

	topendarcshapev2gridFormCallback.probe.ux_tree()
}
func __gong__New__TopStackGrowthCurveEndArcShapeV2FormCallback(
	topstackgrowthcurveendarcshapev2 *models.TopStackGrowthCurveEndArcShapeV2,
	probe *Probe,
	formGroup *form.FormGroup,
) (topstackgrowthcurveendarcshapev2FormCallback *TopStackGrowthCurveEndArcShapeV2FormCallback) {
	topstackgrowthcurveendarcshapev2FormCallback = new(TopStackGrowthCurveEndArcShapeV2FormCallback)
	topstackgrowthcurveendarcshapev2FormCallback.probe = probe
	topstackgrowthcurveendarcshapev2FormCallback.topstackgrowthcurveendarcshapev2 = topstackgrowthcurveendarcshapev2
	topstackgrowthcurveendarcshapev2FormCallback.formGroup = formGroup

	topstackgrowthcurveendarcshapev2FormCallback.CreationMode = (topstackgrowthcurveendarcshapev2 == nil)

	return
}

type TopStackGrowthCurveEndArcShapeV2FormCallback struct {
	topstackgrowthcurveendarcshapev2 *models.TopStackGrowthCurveEndArcShapeV2

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (topstackgrowthcurveendarcshapev2FormCallback *TopStackGrowthCurveEndArcShapeV2FormCallback) OnSave() {
	topstackgrowthcurveendarcshapev2FormCallback.probe.stageOfInterest.Lock()
	defer topstackgrowthcurveendarcshapev2FormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TopStackGrowthCurveEndArcShapeV2FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	topstackgrowthcurveendarcshapev2FormCallback.probe.formStage.Checkout()

	if topstackgrowthcurveendarcshapev2FormCallback.topstackgrowthcurveendarcshapev2 == nil {
		topstackgrowthcurveendarcshapev2FormCallback.topstackgrowthcurveendarcshapev2 = new(models.TopStackGrowthCurveEndArcShapeV2).Stage(topstackgrowthcurveendarcshapev2FormCallback.probe.stageOfInterest)
	}
	topstackgrowthcurveendarcshapev2_ := topstackgrowthcurveendarcshapev2FormCallback.topstackgrowthcurveendarcshapev2
	_ = topstackgrowthcurveendarcshapev2_

	for _, formDiv := range topstackgrowthcurveendarcshapev2FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(topstackgrowthcurveendarcshapev2_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(topstackgrowthcurveendarcshapev2_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(topstackgrowthcurveendarcshapev2_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(topstackgrowthcurveendarcshapev2_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(topstackgrowthcurveendarcshapev2_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(topstackgrowthcurveendarcshapev2_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(topstackgrowthcurveendarcshapev2_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(topstackgrowthcurveendarcshapev2_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(topstackgrowthcurveendarcshapev2_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(topstackgrowthcurveendarcshapev2_.RadiusY), formDiv)
		case "TopStackOfGrowthCurveV2:TopStackGrowthCurveEndArcShapeV2s":
			// 1. Decode the AssociationStorage which contains the rowIDs of the TopStackOfGrowthCurveV2 instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target TopStackOfGrowthCurveV2 instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.TopStackOfGrowthCurveV2](topstackgrowthcurveendarcshapev2FormCallback.probe.stageOfInterest)
			targetTopStackOfGrowthCurveV2IDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetTopStackOfGrowthCurveV2IDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all TopStackOfGrowthCurveV2 instances and update their TopStackGrowthCurveEndArcShapeV2s slice
			for _topstackofgrowthcurvev2 := range *models.GetGongstructInstancesSetFromPointerType[*models.TopStackOfGrowthCurveV2](topstackgrowthcurveendarcshapev2FormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(topstackgrowthcurveendarcshapev2FormCallback.probe.stageOfInterest, _topstackofgrowthcurvev2)
				
				// if TopStackOfGrowthCurveV2 is selected
				if targetTopStackOfGrowthCurveV2IDs[id] {
					// ensure topstackgrowthcurveendarcshapev2_ is in _topstackofgrowthcurvev2.TopStackGrowthCurveEndArcShapeV2s
					found := false
					for _, _b := range _topstackofgrowthcurvev2.TopStackGrowthCurveEndArcShapeV2s {
						if _b == topstackgrowthcurveendarcshapev2_ {
							found = true
							break
						}
					}
					if !found {
						_topstackofgrowthcurvev2.TopStackGrowthCurveEndArcShapeV2s = append(_topstackofgrowthcurvev2.TopStackGrowthCurveEndArcShapeV2s, topstackgrowthcurveendarcshapev2_)
						topstackgrowthcurveendarcshapev2FormCallback.probe.UpdateSliceOfPointersCallback(_topstackofgrowthcurvev2, "TopStackGrowthCurveEndArcShapeV2s", &_topstackofgrowthcurvev2.TopStackGrowthCurveEndArcShapeV2s)
					}
				} else {
					// ensure topstackgrowthcurveendarcshapev2_ is NOT in _topstackofgrowthcurvev2.TopStackGrowthCurveEndArcShapeV2s
					idx := slices.Index(_topstackofgrowthcurvev2.TopStackGrowthCurveEndArcShapeV2s, topstackgrowthcurveendarcshapev2_)
					if idx != -1 {
						_topstackofgrowthcurvev2.TopStackGrowthCurveEndArcShapeV2s = slices.Delete(_topstackofgrowthcurvev2.TopStackGrowthCurveEndArcShapeV2s, idx, idx+1)
						topstackgrowthcurveendarcshapev2FormCallback.probe.UpdateSliceOfPointersCallback(_topstackofgrowthcurvev2, "TopStackGrowthCurveEndArcShapeV2s", &_topstackofgrowthcurvev2.TopStackGrowthCurveEndArcShapeV2s)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if topstackgrowthcurveendarcshapev2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstackgrowthcurveendarcshapev2_.Unstage(topstackgrowthcurveendarcshapev2FormCallback.probe.stageOfInterest)
	}

	topstackgrowthcurveendarcshapev2FormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TopStackGrowthCurveEndArcShapeV2](
		topstackgrowthcurveendarcshapev2FormCallback.probe,
	)

	// display a new form by reset the form stage
	if topstackgrowthcurveendarcshapev2FormCallback.CreationMode || topstackgrowthcurveendarcshapev2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstackgrowthcurveendarcshapev2FormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(topstackgrowthcurveendarcshapev2FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TopStackGrowthCurveEndArcShapeV2FormCallback(
			nil,
			topstackgrowthcurveendarcshapev2FormCallback.probe,
			newFormGroup,
		)
		topstackgrowthcurveendarcshapev2 := new(models.TopStackGrowthCurveEndArcShapeV2)
		FillUpForm(topstackgrowthcurveendarcshapev2, newFormGroup, topstackgrowthcurveendarcshapev2FormCallback.probe)
		topstackgrowthcurveendarcshapev2FormCallback.probe.formStage.Commit()
	}

	topstackgrowthcurveendarcshapev2FormCallback.probe.ux_tree()
}
func __gong__New__TopStackGrowthCurveStartArcShapeV2FormCallback(
	topstackgrowthcurvestartarcshapev2 *models.TopStackGrowthCurveStartArcShapeV2,
	probe *Probe,
	formGroup *form.FormGroup,
) (topstackgrowthcurvestartarcshapev2FormCallback *TopStackGrowthCurveStartArcShapeV2FormCallback) {
	topstackgrowthcurvestartarcshapev2FormCallback = new(TopStackGrowthCurveStartArcShapeV2FormCallback)
	topstackgrowthcurvestartarcshapev2FormCallback.probe = probe
	topstackgrowthcurvestartarcshapev2FormCallback.topstackgrowthcurvestartarcshapev2 = topstackgrowthcurvestartarcshapev2
	topstackgrowthcurvestartarcshapev2FormCallback.formGroup = formGroup

	topstackgrowthcurvestartarcshapev2FormCallback.CreationMode = (topstackgrowthcurvestartarcshapev2 == nil)

	return
}

type TopStackGrowthCurveStartArcShapeV2FormCallback struct {
	topstackgrowthcurvestartarcshapev2 *models.TopStackGrowthCurveStartArcShapeV2

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (topstackgrowthcurvestartarcshapev2FormCallback *TopStackGrowthCurveStartArcShapeV2FormCallback) OnSave() {
	topstackgrowthcurvestartarcshapev2FormCallback.probe.stageOfInterest.Lock()
	defer topstackgrowthcurvestartarcshapev2FormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TopStackGrowthCurveStartArcShapeV2FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	topstackgrowthcurvestartarcshapev2FormCallback.probe.formStage.Checkout()

	if topstackgrowthcurvestartarcshapev2FormCallback.topstackgrowthcurvestartarcshapev2 == nil {
		topstackgrowthcurvestartarcshapev2FormCallback.topstackgrowthcurvestartarcshapev2 = new(models.TopStackGrowthCurveStartArcShapeV2).Stage(topstackgrowthcurvestartarcshapev2FormCallback.probe.stageOfInterest)
	}
	topstackgrowthcurvestartarcshapev2_ := topstackgrowthcurvestartarcshapev2FormCallback.topstackgrowthcurvestartarcshapev2
	_ = topstackgrowthcurvestartarcshapev2_

	for _, formDiv := range topstackgrowthcurvestartarcshapev2FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(topstackgrowthcurvestartarcshapev2_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(topstackgrowthcurvestartarcshapev2_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(topstackgrowthcurvestartarcshapev2_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(topstackgrowthcurvestartarcshapev2_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(topstackgrowthcurvestartarcshapev2_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(topstackgrowthcurvestartarcshapev2_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(topstackgrowthcurvestartarcshapev2_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(topstackgrowthcurvestartarcshapev2_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(topstackgrowthcurvestartarcshapev2_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(topstackgrowthcurvestartarcshapev2_.RadiusY), formDiv)
		case "TopStackOfGrowthCurveV2:TopStackGrowthCurveStartArcShapeV2s":
			// 1. Decode the AssociationStorage which contains the rowIDs of the TopStackOfGrowthCurveV2 instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target TopStackOfGrowthCurveV2 instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.TopStackOfGrowthCurveV2](topstackgrowthcurvestartarcshapev2FormCallback.probe.stageOfInterest)
			targetTopStackOfGrowthCurveV2IDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetTopStackOfGrowthCurveV2IDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all TopStackOfGrowthCurveV2 instances and update their TopStackGrowthCurveStartArcShapeV2s slice
			for _topstackofgrowthcurvev2 := range *models.GetGongstructInstancesSetFromPointerType[*models.TopStackOfGrowthCurveV2](topstackgrowthcurvestartarcshapev2FormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(topstackgrowthcurvestartarcshapev2FormCallback.probe.stageOfInterest, _topstackofgrowthcurvev2)
				
				// if TopStackOfGrowthCurveV2 is selected
				if targetTopStackOfGrowthCurveV2IDs[id] {
					// ensure topstackgrowthcurvestartarcshapev2_ is in _topstackofgrowthcurvev2.TopStackGrowthCurveStartArcShapeV2s
					found := false
					for _, _b := range _topstackofgrowthcurvev2.TopStackGrowthCurveStartArcShapeV2s {
						if _b == topstackgrowthcurvestartarcshapev2_ {
							found = true
							break
						}
					}
					if !found {
						_topstackofgrowthcurvev2.TopStackGrowthCurveStartArcShapeV2s = append(_topstackofgrowthcurvev2.TopStackGrowthCurveStartArcShapeV2s, topstackgrowthcurvestartarcshapev2_)
						topstackgrowthcurvestartarcshapev2FormCallback.probe.UpdateSliceOfPointersCallback(_topstackofgrowthcurvev2, "TopStackGrowthCurveStartArcShapeV2s", &_topstackofgrowthcurvev2.TopStackGrowthCurveStartArcShapeV2s)
					}
				} else {
					// ensure topstackgrowthcurvestartarcshapev2_ is NOT in _topstackofgrowthcurvev2.TopStackGrowthCurveStartArcShapeV2s
					idx := slices.Index(_topstackofgrowthcurvev2.TopStackGrowthCurveStartArcShapeV2s, topstackgrowthcurvestartarcshapev2_)
					if idx != -1 {
						_topstackofgrowthcurvev2.TopStackGrowthCurveStartArcShapeV2s = slices.Delete(_topstackofgrowthcurvev2.TopStackGrowthCurveStartArcShapeV2s, idx, idx+1)
						topstackgrowthcurvestartarcshapev2FormCallback.probe.UpdateSliceOfPointersCallback(_topstackofgrowthcurvev2, "TopStackGrowthCurveStartArcShapeV2s", &_topstackofgrowthcurvev2.TopStackGrowthCurveStartArcShapeV2s)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if topstackgrowthcurvestartarcshapev2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstackgrowthcurvestartarcshapev2_.Unstage(topstackgrowthcurvestartarcshapev2FormCallback.probe.stageOfInterest)
	}

	topstackgrowthcurvestartarcshapev2FormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TopStackGrowthCurveStartArcShapeV2](
		topstackgrowthcurvestartarcshapev2FormCallback.probe,
	)

	// display a new form by reset the form stage
	if topstackgrowthcurvestartarcshapev2FormCallback.CreationMode || topstackgrowthcurvestartarcshapev2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstackgrowthcurvestartarcshapev2FormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(topstackgrowthcurvestartarcshapev2FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TopStackGrowthCurveStartArcShapeV2FormCallback(
			nil,
			topstackgrowthcurvestartarcshapev2FormCallback.probe,
			newFormGroup,
		)
		topstackgrowthcurvestartarcshapev2 := new(models.TopStackGrowthCurveStartArcShapeV2)
		FillUpForm(topstackgrowthcurvestartarcshapev2, newFormGroup, topstackgrowthcurvestartarcshapev2FormCallback.probe)
		topstackgrowthcurvestartarcshapev2FormCallback.probe.formStage.Commit()
	}

	topstackgrowthcurvestartarcshapev2FormCallback.probe.ux_tree()
}
func __gong__New__TopStackOfGrowthCurveV2FormCallback(
	topstackofgrowthcurvev2 *models.TopStackOfGrowthCurveV2,
	probe *Probe,
	formGroup *form.FormGroup,
) (topstackofgrowthcurvev2FormCallback *TopStackOfGrowthCurveV2FormCallback) {
	topstackofgrowthcurvev2FormCallback = new(TopStackOfGrowthCurveV2FormCallback)
	topstackofgrowthcurvev2FormCallback.probe = probe
	topstackofgrowthcurvev2FormCallback.topstackofgrowthcurvev2 = topstackofgrowthcurvev2
	topstackofgrowthcurvev2FormCallback.formGroup = formGroup

	topstackofgrowthcurvev2FormCallback.CreationMode = (topstackofgrowthcurvev2 == nil)

	return
}

type TopStackOfGrowthCurveV2FormCallback struct {
	topstackofgrowthcurvev2 *models.TopStackOfGrowthCurveV2

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (topstackofgrowthcurvev2FormCallback *TopStackOfGrowthCurveV2FormCallback) OnSave() {
	topstackofgrowthcurvev2FormCallback.probe.stageOfInterest.Lock()
	defer topstackofgrowthcurvev2FormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TopStackOfGrowthCurveV2FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	topstackofgrowthcurvev2FormCallback.probe.formStage.Checkout()

	if topstackofgrowthcurvev2FormCallback.topstackofgrowthcurvev2 == nil {
		topstackofgrowthcurvev2FormCallback.topstackofgrowthcurvev2 = new(models.TopStackOfGrowthCurveV2).Stage(topstackofgrowthcurvev2FormCallback.probe.stageOfInterest)
	}
	topstackofgrowthcurvev2_ := topstackofgrowthcurvev2FormCallback.topstackofgrowthcurvev2
	_ = topstackofgrowthcurvev2_

	for _, formDiv := range topstackofgrowthcurvev2FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(topstackofgrowthcurvev2_.Name), formDiv)
		case "TopStackGrowthCurveStartArcShapeV2s":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.TopStackGrowthCurveStartArcShapeV2](topstackofgrowthcurvev2FormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.TopStackGrowthCurveStartArcShapeV2, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.TopStackGrowthCurveStartArcShapeV2)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					topstackofgrowthcurvev2FormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.TopStackGrowthCurveStartArcShapeV2](topstackofgrowthcurvev2FormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			topstackofgrowthcurvev2_.TopStackGrowthCurveStartArcShapeV2s = instanceSlice
			topstackofgrowthcurvev2FormCallback.probe.UpdateSliceOfPointersCallback(topstackofgrowthcurvev2_, "TopStackGrowthCurveStartArcShapeV2s", &topstackofgrowthcurvev2_.TopStackGrowthCurveStartArcShapeV2s)

		case "TopStackGrowthCurveEndArcShapeV2s":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.TopStackGrowthCurveEndArcShapeV2](topstackofgrowthcurvev2FormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.TopStackGrowthCurveEndArcShapeV2, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.TopStackGrowthCurveEndArcShapeV2)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					topstackofgrowthcurvev2FormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.TopStackGrowthCurveEndArcShapeV2](topstackofgrowthcurvev2FormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			topstackofgrowthcurvev2_.TopStackGrowthCurveEndArcShapeV2s = instanceSlice
			topstackofgrowthcurvev2FormCallback.probe.UpdateSliceOfPointersCallback(topstackofgrowthcurvev2_, "TopStackGrowthCurveEndArcShapeV2s", &topstackofgrowthcurvev2_.TopStackGrowthCurveEndArcShapeV2s)

		}
	}

	// manage the suppress operation
	if topstackofgrowthcurvev2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstackofgrowthcurvev2_.Unstage(topstackofgrowthcurvev2FormCallback.probe.stageOfInterest)
	}

	topstackofgrowthcurvev2FormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TopStackOfGrowthCurveV2](
		topstackofgrowthcurvev2FormCallback.probe,
	)

	// display a new form by reset the form stage
	if topstackofgrowthcurvev2FormCallback.CreationMode || topstackofgrowthcurvev2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstackofgrowthcurvev2FormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(topstackofgrowthcurvev2FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TopStackOfGrowthCurveV2FormCallback(
			nil,
			topstackofgrowthcurvev2FormCallback.probe,
			newFormGroup,
		)
		topstackofgrowthcurvev2 := new(models.TopStackOfGrowthCurveV2)
		FillUpForm(topstackofgrowthcurvev2, newFormGroup, topstackofgrowthcurvev2FormCallback.probe)
		topstackofgrowthcurvev2FormCallback.probe.formStage.Commit()
	}

	topstackofgrowthcurvev2FormCallback.probe.ux_tree()
}
func __gong__New__TopStartArcShapeV2FormCallback(
	topstartarcshapev2 *models.TopStartArcShapeV2,
	probe *Probe,
	formGroup *form.FormGroup,
) (topstartarcshapev2FormCallback *TopStartArcShapeV2FormCallback) {
	topstartarcshapev2FormCallback = new(TopStartArcShapeV2FormCallback)
	topstartarcshapev2FormCallback.probe = probe
	topstartarcshapev2FormCallback.topstartarcshapev2 = topstartarcshapev2
	topstartarcshapev2FormCallback.formGroup = formGroup

	topstartarcshapev2FormCallback.CreationMode = (topstartarcshapev2 == nil)

	return
}

type TopStartArcShapeV2FormCallback struct {
	topstartarcshapev2 *models.TopStartArcShapeV2

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (topstartarcshapev2FormCallback *TopStartArcShapeV2FormCallback) OnSave() {
	topstartarcshapev2FormCallback.probe.stageOfInterest.Lock()
	defer topstartarcshapev2FormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TopStartArcShapeV2FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	topstartarcshapev2FormCallback.probe.formStage.Checkout()

	if topstartarcshapev2FormCallback.topstartarcshapev2 == nil {
		topstartarcshapev2FormCallback.topstartarcshapev2 = new(models.TopStartArcShapeV2).Stage(topstartarcshapev2FormCallback.probe.stageOfInterest)
	}
	topstartarcshapev2_ := topstartarcshapev2FormCallback.topstartarcshapev2
	_ = topstartarcshapev2_

	for _, formDiv := range topstartarcshapev2FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(topstartarcshapev2_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(topstartarcshapev2_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(topstartarcshapev2_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(topstartarcshapev2_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(topstartarcshapev2_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(topstartarcshapev2_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(topstartarcshapev2_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(topstartarcshapev2_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(topstartarcshapev2_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(topstartarcshapev2_.RadiusY), formDiv)
		case "TopStartArcShapeV2Grid:TopStartArcShapesV2":
			// 1. Decode the AssociationStorage which contains the rowIDs of the TopStartArcShapeV2Grid instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target TopStartArcShapeV2Grid instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.TopStartArcShapeV2Grid](topstartarcshapev2FormCallback.probe.stageOfInterest)
			targetTopStartArcShapeV2GridIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetTopStartArcShapeV2GridIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all TopStartArcShapeV2Grid instances and update their TopStartArcShapesV2 slice
			for _topstartarcshapev2grid := range *models.GetGongstructInstancesSetFromPointerType[*models.TopStartArcShapeV2Grid](topstartarcshapev2FormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(topstartarcshapev2FormCallback.probe.stageOfInterest, _topstartarcshapev2grid)
				
				// if TopStartArcShapeV2Grid is selected
				if targetTopStartArcShapeV2GridIDs[id] {
					// ensure topstartarcshapev2_ is in _topstartarcshapev2grid.TopStartArcShapesV2
					found := false
					for _, _b := range _topstartarcshapev2grid.TopStartArcShapesV2 {
						if _b == topstartarcshapev2_ {
							found = true
							break
						}
					}
					if !found {
						_topstartarcshapev2grid.TopStartArcShapesV2 = append(_topstartarcshapev2grid.TopStartArcShapesV2, topstartarcshapev2_)
						topstartarcshapev2FormCallback.probe.UpdateSliceOfPointersCallback(_topstartarcshapev2grid, "TopStartArcShapesV2", &_topstartarcshapev2grid.TopStartArcShapesV2)
					}
				} else {
					// ensure topstartarcshapev2_ is NOT in _topstartarcshapev2grid.TopStartArcShapesV2
					idx := slices.Index(_topstartarcshapev2grid.TopStartArcShapesV2, topstartarcshapev2_)
					if idx != -1 {
						_topstartarcshapev2grid.TopStartArcShapesV2 = slices.Delete(_topstartarcshapev2grid.TopStartArcShapesV2, idx, idx+1)
						topstartarcshapev2FormCallback.probe.UpdateSliceOfPointersCallback(_topstartarcshapev2grid, "TopStartArcShapesV2", &_topstartarcshapev2grid.TopStartArcShapesV2)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if topstartarcshapev2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstartarcshapev2_.Unstage(topstartarcshapev2FormCallback.probe.stageOfInterest)
	}

	topstartarcshapev2FormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TopStartArcShapeV2](
		topstartarcshapev2FormCallback.probe,
	)

	// display a new form by reset the form stage
	if topstartarcshapev2FormCallback.CreationMode || topstartarcshapev2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstartarcshapev2FormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(topstartarcshapev2FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TopStartArcShapeV2FormCallback(
			nil,
			topstartarcshapev2FormCallback.probe,
			newFormGroup,
		)
		topstartarcshapev2 := new(models.TopStartArcShapeV2)
		FillUpForm(topstartarcshapev2, newFormGroup, topstartarcshapev2FormCallback.probe)
		topstartarcshapev2FormCallback.probe.formStage.Commit()
	}

	topstartarcshapev2FormCallback.probe.ux_tree()
}
func __gong__New__TopStartArcShapeV2GridFormCallback(
	topstartarcshapev2grid *models.TopStartArcShapeV2Grid,
	probe *Probe,
	formGroup *form.FormGroup,
) (topstartarcshapev2gridFormCallback *TopStartArcShapeV2GridFormCallback) {
	topstartarcshapev2gridFormCallback = new(TopStartArcShapeV2GridFormCallback)
	topstartarcshapev2gridFormCallback.probe = probe
	topstartarcshapev2gridFormCallback.topstartarcshapev2grid = topstartarcshapev2grid
	topstartarcshapev2gridFormCallback.formGroup = formGroup

	topstartarcshapev2gridFormCallback.CreationMode = (topstartarcshapev2grid == nil)

	return
}

type TopStartArcShapeV2GridFormCallback struct {
	topstartarcshapev2grid *models.TopStartArcShapeV2Grid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (topstartarcshapev2gridFormCallback *TopStartArcShapeV2GridFormCallback) OnSave() {
	topstartarcshapev2gridFormCallback.probe.stageOfInterest.Lock()
	defer topstartarcshapev2gridFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TopStartArcShapeV2GridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	topstartarcshapev2gridFormCallback.probe.formStage.Checkout()

	if topstartarcshapev2gridFormCallback.topstartarcshapev2grid == nil {
		topstartarcshapev2gridFormCallback.topstartarcshapev2grid = new(models.TopStartArcShapeV2Grid).Stage(topstartarcshapev2gridFormCallback.probe.stageOfInterest)
	}
	topstartarcshapev2grid_ := topstartarcshapev2gridFormCallback.topstartarcshapev2grid
	_ = topstartarcshapev2grid_

	for _, formDiv := range topstartarcshapev2gridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(topstartarcshapev2grid_.Name), formDiv)
		case "TopStartArcShapesV2":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.TopStartArcShapeV2](topstartarcshapev2gridFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.TopStartArcShapeV2, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.TopStartArcShapeV2)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					topstartarcshapev2gridFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.TopStartArcShapeV2](topstartarcshapev2gridFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			topstartarcshapev2grid_.TopStartArcShapesV2 = instanceSlice
			topstartarcshapev2gridFormCallback.probe.UpdateSliceOfPointersCallback(topstartarcshapev2grid_, "TopStartArcShapesV2", &topstartarcshapev2grid_.TopStartArcShapesV2)

		}
	}

	// manage the suppress operation
	if topstartarcshapev2gridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstartarcshapev2grid_.Unstage(topstartarcshapev2gridFormCallback.probe.stageOfInterest)
	}

	topstartarcshapev2gridFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TopStartArcShapeV2Grid](
		topstartarcshapev2gridFormCallback.probe,
	)

	// display a new form by reset the form stage
	if topstartarcshapev2gridFormCallback.CreationMode || topstartarcshapev2gridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstartarcshapev2gridFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(topstartarcshapev2gridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TopStartArcShapeV2GridFormCallback(
			nil,
			topstartarcshapev2gridFormCallback.probe,
			newFormGroup,
		)
		topstartarcshapev2grid := new(models.TopStartArcShapeV2Grid)
		FillUpForm(topstartarcshapev2grid, newFormGroup, topstartarcshapev2gridFormCallback.probe)
		topstartarcshapev2gridFormCallback.probe.formStage.Commit()
	}

	topstartarcshapev2gridFormCallback.probe.ux_tree()
}
