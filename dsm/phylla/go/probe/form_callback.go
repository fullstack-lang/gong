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
func __gong__New__EndHalfwayArcShapeFormCallback(
	endhalfwayarcshape *models.EndHalfwayArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (endhalfwayarcshapeFormCallback *EndHalfwayArcShapeFormCallback) {
	endhalfwayarcshapeFormCallback = new(EndHalfwayArcShapeFormCallback)
	endhalfwayarcshapeFormCallback.probe = probe
	endhalfwayarcshapeFormCallback.endhalfwayarcshape = endhalfwayarcshape
	endhalfwayarcshapeFormCallback.formGroup = formGroup

	endhalfwayarcshapeFormCallback.CreationMode = (endhalfwayarcshape == nil)

	return
}

type EndHalfwayArcShapeFormCallback struct {
	endhalfwayarcshape *models.EndHalfwayArcShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (endhalfwayarcshapeFormCallback *EndHalfwayArcShapeFormCallback) OnSave() {
	endhalfwayarcshapeFormCallback.probe.stageOfInterest.Lock()
	defer endhalfwayarcshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("EndHalfwayArcShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	endhalfwayarcshapeFormCallback.probe.formStage.Checkout()

	if endhalfwayarcshapeFormCallback.endhalfwayarcshape == nil {
		endhalfwayarcshapeFormCallback.endhalfwayarcshape = new(models.EndHalfwayArcShape).Stage(endhalfwayarcshapeFormCallback.probe.stageOfInterest)
	}
	endhalfwayarcshape_ := endhalfwayarcshapeFormCallback.endhalfwayarcshape
	_ = endhalfwayarcshape_

	for _, formDiv := range endhalfwayarcshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(endhalfwayarcshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(endhalfwayarcshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(endhalfwayarcshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(endhalfwayarcshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(endhalfwayarcshape_.EndY), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(endhalfwayarcshape_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(endhalfwayarcshape_.RadiusY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(endhalfwayarcshape_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(endhalfwayarcshape_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(endhalfwayarcshape_.SweepFlag), formDiv)
		case "EndHalfwayArcShapeGrid:EndHalfwayArcShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the EndHalfwayArcShapeGrid instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target EndHalfwayArcShapeGrid instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.EndHalfwayArcShapeGrid](endhalfwayarcshapeFormCallback.probe.stageOfInterest)
			targetEndHalfwayArcShapeGridIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetEndHalfwayArcShapeGridIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all EndHalfwayArcShapeGrid instances and update their EndHalfwayArcShapes slice
			for _endhalfwayarcshapegrid := range *models.GetGongstructInstancesSetFromPointerType[*models.EndHalfwayArcShapeGrid](endhalfwayarcshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(endhalfwayarcshapeFormCallback.probe.stageOfInterest, _endhalfwayarcshapegrid)
				
				// if EndHalfwayArcShapeGrid is selected
				if targetEndHalfwayArcShapeGridIDs[id] {
					// ensure endhalfwayarcshape_ is in _endhalfwayarcshapegrid.EndHalfwayArcShapes
					found := false
					for _, _b := range _endhalfwayarcshapegrid.EndHalfwayArcShapes {
						if _b == endhalfwayarcshape_ {
							found = true
							break
						}
					}
					if !found {
						_endhalfwayarcshapegrid.EndHalfwayArcShapes = append(_endhalfwayarcshapegrid.EndHalfwayArcShapes, endhalfwayarcshape_)
						endhalfwayarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_endhalfwayarcshapegrid, "EndHalfwayArcShapes", &_endhalfwayarcshapegrid.EndHalfwayArcShapes)
					}
				} else {
					// ensure endhalfwayarcshape_ is NOT in _endhalfwayarcshapegrid.EndHalfwayArcShapes
					idx := slices.Index(_endhalfwayarcshapegrid.EndHalfwayArcShapes, endhalfwayarcshape_)
					if idx != -1 {
						_endhalfwayarcshapegrid.EndHalfwayArcShapes = slices.Delete(_endhalfwayarcshapegrid.EndHalfwayArcShapes, idx, idx+1)
						endhalfwayarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_endhalfwayarcshapegrid, "EndHalfwayArcShapes", &_endhalfwayarcshapegrid.EndHalfwayArcShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if endhalfwayarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		endhalfwayarcshape_.Unstage(endhalfwayarcshapeFormCallback.probe.stageOfInterest)
	}

	endhalfwayarcshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.EndHalfwayArcShape](
		endhalfwayarcshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if endhalfwayarcshapeFormCallback.CreationMode || endhalfwayarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		endhalfwayarcshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(endhalfwayarcshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__EndHalfwayArcShapeFormCallback(
			nil,
			endhalfwayarcshapeFormCallback.probe,
			newFormGroup,
		)
		endhalfwayarcshape := new(models.EndHalfwayArcShape)
		FillUpForm(endhalfwayarcshape, newFormGroup, endhalfwayarcshapeFormCallback.probe)
		endhalfwayarcshapeFormCallback.probe.formStage.Commit()
	}

	endhalfwayarcshapeFormCallback.probe.ux_tree()
}
func __gong__New__EndHalfwayArcShapeGridFormCallback(
	endhalfwayarcshapegrid *models.EndHalfwayArcShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (endhalfwayarcshapegridFormCallback *EndHalfwayArcShapeGridFormCallback) {
	endhalfwayarcshapegridFormCallback = new(EndHalfwayArcShapeGridFormCallback)
	endhalfwayarcshapegridFormCallback.probe = probe
	endhalfwayarcshapegridFormCallback.endhalfwayarcshapegrid = endhalfwayarcshapegrid
	endhalfwayarcshapegridFormCallback.formGroup = formGroup

	endhalfwayarcshapegridFormCallback.CreationMode = (endhalfwayarcshapegrid == nil)

	return
}

type EndHalfwayArcShapeGridFormCallback struct {
	endhalfwayarcshapegrid *models.EndHalfwayArcShapeGrid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (endhalfwayarcshapegridFormCallback *EndHalfwayArcShapeGridFormCallback) OnSave() {
	endhalfwayarcshapegridFormCallback.probe.stageOfInterest.Lock()
	defer endhalfwayarcshapegridFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("EndHalfwayArcShapeGridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	endhalfwayarcshapegridFormCallback.probe.formStage.Checkout()

	if endhalfwayarcshapegridFormCallback.endhalfwayarcshapegrid == nil {
		endhalfwayarcshapegridFormCallback.endhalfwayarcshapegrid = new(models.EndHalfwayArcShapeGrid).Stage(endhalfwayarcshapegridFormCallback.probe.stageOfInterest)
	}
	endhalfwayarcshapegrid_ := endhalfwayarcshapegridFormCallback.endhalfwayarcshapegrid
	_ = endhalfwayarcshapegrid_

	for _, formDiv := range endhalfwayarcshapegridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(endhalfwayarcshapegrid_.Name), formDiv)
		case "EndHalfwayArcShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.EndHalfwayArcShape](endhalfwayarcshapegridFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.EndHalfwayArcShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.EndHalfwayArcShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					endhalfwayarcshapegridFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.EndHalfwayArcShape](endhalfwayarcshapegridFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			endhalfwayarcshapegrid_.EndHalfwayArcShapes = instanceSlice
			endhalfwayarcshapegridFormCallback.probe.UpdateSliceOfPointersCallback(endhalfwayarcshapegrid_, "EndHalfwayArcShapes", &endhalfwayarcshapegrid_.EndHalfwayArcShapes)

		}
	}

	// manage the suppress operation
	if endhalfwayarcshapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		endhalfwayarcshapegrid_.Unstage(endhalfwayarcshapegridFormCallback.probe.stageOfInterest)
	}

	endhalfwayarcshapegridFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.EndHalfwayArcShapeGrid](
		endhalfwayarcshapegridFormCallback.probe,
	)

	// display a new form by reset the form stage
	if endhalfwayarcshapegridFormCallback.CreationMode || endhalfwayarcshapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		endhalfwayarcshapegridFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(endhalfwayarcshapegridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__EndHalfwayArcShapeGridFormCallback(
			nil,
			endhalfwayarcshapegridFormCallback.probe,
			newFormGroup,
		)
		endhalfwayarcshapegrid := new(models.EndHalfwayArcShapeGrid)
		FillUpForm(endhalfwayarcshapegrid, newFormGroup, endhalfwayarcshapegridFormCallback.probe)
		endhalfwayarcshapegridFormCallback.probe.formStage.Commit()
	}

	endhalfwayarcshapegridFormCallback.probe.ux_tree()
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
		case "StartHalfwayArcShapeGrid":
			FormDivSelectFieldToField(&(growthcurve2d_.StartHalfwayArcShapeGrid), growthcurve2dFormCallback.probe.stageOfInterest, formDiv)
		case "EndHalfwayArcShapeGrid":
			FormDivSelectFieldToField(&(growthcurve2d_.EndHalfwayArcShapeGrid), growthcurve2dFormCallback.probe.stageOfInterest, formDiv)
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
func __gong__New__GrowthCurve2DRibbonFormCallback(
	growthcurve2dribbon *models.GrowthCurve2DRibbon,
	probe *Probe,
	formGroup *form.FormGroup,
) (growthcurve2dribbonFormCallback *GrowthCurve2DRibbonFormCallback) {
	growthcurve2dribbonFormCallback = new(GrowthCurve2DRibbonFormCallback)
	growthcurve2dribbonFormCallback.probe = probe
	growthcurve2dribbonFormCallback.growthcurve2dribbon = growthcurve2dribbon
	growthcurve2dribbonFormCallback.formGroup = formGroup

	growthcurve2dribbonFormCallback.CreationMode = (growthcurve2dribbon == nil)

	return
}

type GrowthCurve2DRibbonFormCallback struct {
	growthcurve2dribbon *models.GrowthCurve2DRibbon

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (growthcurve2dribbonFormCallback *GrowthCurve2DRibbonFormCallback) OnSave() {
	growthcurve2dribbonFormCallback.probe.stageOfInterest.Lock()
	defer growthcurve2dribbonFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("GrowthCurve2DRibbonFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	growthcurve2dribbonFormCallback.probe.formStage.Checkout()

	if growthcurve2dribbonFormCallback.growthcurve2dribbon == nil {
		growthcurve2dribbonFormCallback.growthcurve2dribbon = new(models.GrowthCurve2DRibbon).Stage(growthcurve2dribbonFormCallback.probe.stageOfInterest)
	}
	growthcurve2dribbon_ := growthcurve2dribbonFormCallback.growthcurve2dribbon
	_ = growthcurve2dribbon_

	for _, formDiv := range growthcurve2dribbonFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(growthcurve2dribbon_.Name), formDiv)
		case "GrowthCurve2DRibbonStartShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.GrowthCurve2DRibbonStartShape](growthcurve2dribbonFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.GrowthCurve2DRibbonStartShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.GrowthCurve2DRibbonStartShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					growthcurve2dribbonFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.GrowthCurve2DRibbonStartShape](growthcurve2dribbonFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			growthcurve2dribbon_.GrowthCurve2DRibbonStartShapes = instanceSlice
			growthcurve2dribbonFormCallback.probe.UpdateSliceOfPointersCallback(growthcurve2dribbon_, "GrowthCurve2DRibbonStartShapes", &growthcurve2dribbon_.GrowthCurve2DRibbonStartShapes)

		case "GrowthCurve2DRibbonEndShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.GrowthCurve2DRibbonEndShape](growthcurve2dribbonFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.GrowthCurve2DRibbonEndShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.GrowthCurve2DRibbonEndShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					growthcurve2dribbonFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.GrowthCurve2DRibbonEndShape](growthcurve2dribbonFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			growthcurve2dribbon_.GrowthCurve2DRibbonEndShapes = instanceSlice
			growthcurve2dribbonFormCallback.probe.UpdateSliceOfPointersCallback(growthcurve2dribbon_, "GrowthCurve2DRibbonEndShapes", &growthcurve2dribbon_.GrowthCurve2DRibbonEndShapes)

		}
	}

	// manage the suppress operation
	if growthcurve2dribbonFormCallback.formGroup.HasSuppressButtonBeenPressed {
		growthcurve2dribbon_.Unstage(growthcurve2dribbonFormCallback.probe.stageOfInterest)
	}

	growthcurve2dribbonFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.GrowthCurve2DRibbon](
		growthcurve2dribbonFormCallback.probe,
	)

	// display a new form by reset the form stage
	if growthcurve2dribbonFormCallback.CreationMode || growthcurve2dribbonFormCallback.formGroup.HasSuppressButtonBeenPressed {
		growthcurve2dribbonFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(growthcurve2dribbonFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GrowthCurve2DRibbonFormCallback(
			nil,
			growthcurve2dribbonFormCallback.probe,
			newFormGroup,
		)
		growthcurve2dribbon := new(models.GrowthCurve2DRibbon)
		FillUpForm(growthcurve2dribbon, newFormGroup, growthcurve2dribbonFormCallback.probe)
		growthcurve2dribbonFormCallback.probe.formStage.Commit()
	}

	growthcurve2dribbonFormCallback.probe.ux_tree()
}
func __gong__New__GrowthCurve2DRibbonEndShapeFormCallback(
	growthcurve2dribbonendshape *models.GrowthCurve2DRibbonEndShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (growthcurve2dribbonendshapeFormCallback *GrowthCurve2DRibbonEndShapeFormCallback) {
	growthcurve2dribbonendshapeFormCallback = new(GrowthCurve2DRibbonEndShapeFormCallback)
	growthcurve2dribbonendshapeFormCallback.probe = probe
	growthcurve2dribbonendshapeFormCallback.growthcurve2dribbonendshape = growthcurve2dribbonendshape
	growthcurve2dribbonendshapeFormCallback.formGroup = formGroup

	growthcurve2dribbonendshapeFormCallback.CreationMode = (growthcurve2dribbonendshape == nil)

	return
}

type GrowthCurve2DRibbonEndShapeFormCallback struct {
	growthcurve2dribbonendshape *models.GrowthCurve2DRibbonEndShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (growthcurve2dribbonendshapeFormCallback *GrowthCurve2DRibbonEndShapeFormCallback) OnSave() {
	growthcurve2dribbonendshapeFormCallback.probe.stageOfInterest.Lock()
	defer growthcurve2dribbonendshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("GrowthCurve2DRibbonEndShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	growthcurve2dribbonendshapeFormCallback.probe.formStage.Checkout()

	if growthcurve2dribbonendshapeFormCallback.growthcurve2dribbonendshape == nil {
		growthcurve2dribbonendshapeFormCallback.growthcurve2dribbonendshape = new(models.GrowthCurve2DRibbonEndShape).Stage(growthcurve2dribbonendshapeFormCallback.probe.stageOfInterest)
	}
	growthcurve2dribbonendshape_ := growthcurve2dribbonendshapeFormCallback.growthcurve2dribbonendshape
	_ = growthcurve2dribbonendshape_

	for _, formDiv := range growthcurve2dribbonendshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(growthcurve2dribbonendshape_.Name), formDiv)
		case "BottomStartX":
			FormDivBasicFieldToField(&(growthcurve2dribbonendshape_.BottomStartX), formDiv)
		case "BottomStartY":
			FormDivBasicFieldToField(&(growthcurve2dribbonendshape_.BottomStartY), formDiv)
		case "BottomEndX":
			FormDivBasicFieldToField(&(growthcurve2dribbonendshape_.BottomEndX), formDiv)
		case "BottomEndY":
			FormDivBasicFieldToField(&(growthcurve2dribbonendshape_.BottomEndY), formDiv)
		case "BottomRadiusX":
			FormDivBasicFieldToField(&(growthcurve2dribbonendshape_.BottomRadiusX), formDiv)
		case "BottomRadiusY":
			FormDivBasicFieldToField(&(growthcurve2dribbonendshape_.BottomRadiusY), formDiv)
		case "BottomXAxisRotation":
			FormDivBasicFieldToField(&(growthcurve2dribbonendshape_.BottomXAxisRotation), formDiv)
		case "BottomLargeArcFlag":
			FormDivBasicFieldToField(&(growthcurve2dribbonendshape_.BottomLargeArcFlag), formDiv)
		case "BottomSweepFlag":
			FormDivBasicFieldToField(&(growthcurve2dribbonendshape_.BottomSweepFlag), formDiv)
		case "TopStartX":
			FormDivBasicFieldToField(&(growthcurve2dribbonendshape_.TopStartX), formDiv)
		case "TopStartY":
			FormDivBasicFieldToField(&(growthcurve2dribbonendshape_.TopStartY), formDiv)
		case "TopEndX":
			FormDivBasicFieldToField(&(growthcurve2dribbonendshape_.TopEndX), formDiv)
		case "TopEndY":
			FormDivBasicFieldToField(&(growthcurve2dribbonendshape_.TopEndY), formDiv)
		case "TopRadiusX":
			FormDivBasicFieldToField(&(growthcurve2dribbonendshape_.TopRadiusX), formDiv)
		case "TopRadiusY":
			FormDivBasicFieldToField(&(growthcurve2dribbonendshape_.TopRadiusY), formDiv)
		case "TopXAxisRotation":
			FormDivBasicFieldToField(&(growthcurve2dribbonendshape_.TopXAxisRotation), formDiv)
		case "TopLargeArcFlag":
			FormDivBasicFieldToField(&(growthcurve2dribbonendshape_.TopLargeArcFlag), formDiv)
		case "TopSweepFlag":
			FormDivBasicFieldToField(&(growthcurve2dribbonendshape_.TopSweepFlag), formDiv)
		case "GrowthCurve2DRibbon:GrowthCurve2DRibbonEndShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the GrowthCurve2DRibbon instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target GrowthCurve2DRibbon instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.GrowthCurve2DRibbon](growthcurve2dribbonendshapeFormCallback.probe.stageOfInterest)
			targetGrowthCurve2DRibbonIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetGrowthCurve2DRibbonIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all GrowthCurve2DRibbon instances and update their GrowthCurve2DRibbonEndShapes slice
			for _growthcurve2dribbon := range *models.GetGongstructInstancesSetFromPointerType[*models.GrowthCurve2DRibbon](growthcurve2dribbonendshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(growthcurve2dribbonendshapeFormCallback.probe.stageOfInterest, _growthcurve2dribbon)
				
				// if GrowthCurve2DRibbon is selected
				if targetGrowthCurve2DRibbonIDs[id] {
					// ensure growthcurve2dribbonendshape_ is in _growthcurve2dribbon.GrowthCurve2DRibbonEndShapes
					found := false
					for _, _b := range _growthcurve2dribbon.GrowthCurve2DRibbonEndShapes {
						if _b == growthcurve2dribbonendshape_ {
							found = true
							break
						}
					}
					if !found {
						_growthcurve2dribbon.GrowthCurve2DRibbonEndShapes = append(_growthcurve2dribbon.GrowthCurve2DRibbonEndShapes, growthcurve2dribbonendshape_)
						growthcurve2dribbonendshapeFormCallback.probe.UpdateSliceOfPointersCallback(_growthcurve2dribbon, "GrowthCurve2DRibbonEndShapes", &_growthcurve2dribbon.GrowthCurve2DRibbonEndShapes)
					}
				} else {
					// ensure growthcurve2dribbonendshape_ is NOT in _growthcurve2dribbon.GrowthCurve2DRibbonEndShapes
					idx := slices.Index(_growthcurve2dribbon.GrowthCurve2DRibbonEndShapes, growthcurve2dribbonendshape_)
					if idx != -1 {
						_growthcurve2dribbon.GrowthCurve2DRibbonEndShapes = slices.Delete(_growthcurve2dribbon.GrowthCurve2DRibbonEndShapes, idx, idx+1)
						growthcurve2dribbonendshapeFormCallback.probe.UpdateSliceOfPointersCallback(_growthcurve2dribbon, "GrowthCurve2DRibbonEndShapes", &_growthcurve2dribbon.GrowthCurve2DRibbonEndShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if growthcurve2dribbonendshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		growthcurve2dribbonendshape_.Unstage(growthcurve2dribbonendshapeFormCallback.probe.stageOfInterest)
	}

	growthcurve2dribbonendshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.GrowthCurve2DRibbonEndShape](
		growthcurve2dribbonendshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if growthcurve2dribbonendshapeFormCallback.CreationMode || growthcurve2dribbonendshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		growthcurve2dribbonendshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(growthcurve2dribbonendshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GrowthCurve2DRibbonEndShapeFormCallback(
			nil,
			growthcurve2dribbonendshapeFormCallback.probe,
			newFormGroup,
		)
		growthcurve2dribbonendshape := new(models.GrowthCurve2DRibbonEndShape)
		FillUpForm(growthcurve2dribbonendshape, newFormGroup, growthcurve2dribbonendshapeFormCallback.probe)
		growthcurve2dribbonendshapeFormCallback.probe.formStage.Commit()
	}

	growthcurve2dribbonendshapeFormCallback.probe.ux_tree()
}
func __gong__New__GrowthCurve2DRibbonStartShapeFormCallback(
	growthcurve2dribbonstartshape *models.GrowthCurve2DRibbonStartShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (growthcurve2dribbonstartshapeFormCallback *GrowthCurve2DRibbonStartShapeFormCallback) {
	growthcurve2dribbonstartshapeFormCallback = new(GrowthCurve2DRibbonStartShapeFormCallback)
	growthcurve2dribbonstartshapeFormCallback.probe = probe
	growthcurve2dribbonstartshapeFormCallback.growthcurve2dribbonstartshape = growthcurve2dribbonstartshape
	growthcurve2dribbonstartshapeFormCallback.formGroup = formGroup

	growthcurve2dribbonstartshapeFormCallback.CreationMode = (growthcurve2dribbonstartshape == nil)

	return
}

type GrowthCurve2DRibbonStartShapeFormCallback struct {
	growthcurve2dribbonstartshape *models.GrowthCurve2DRibbonStartShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (growthcurve2dribbonstartshapeFormCallback *GrowthCurve2DRibbonStartShapeFormCallback) OnSave() {
	growthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest.Lock()
	defer growthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("GrowthCurve2DRibbonStartShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	growthcurve2dribbonstartshapeFormCallback.probe.formStage.Checkout()

	if growthcurve2dribbonstartshapeFormCallback.growthcurve2dribbonstartshape == nil {
		growthcurve2dribbonstartshapeFormCallback.growthcurve2dribbonstartshape = new(models.GrowthCurve2DRibbonStartShape).Stage(growthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest)
	}
	growthcurve2dribbonstartshape_ := growthcurve2dribbonstartshapeFormCallback.growthcurve2dribbonstartshape
	_ = growthcurve2dribbonstartshape_

	for _, formDiv := range growthcurve2dribbonstartshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(growthcurve2dribbonstartshape_.Name), formDiv)
		case "BottomStartX":
			FormDivBasicFieldToField(&(growthcurve2dribbonstartshape_.BottomStartX), formDiv)
		case "BottomStartY":
			FormDivBasicFieldToField(&(growthcurve2dribbonstartshape_.BottomStartY), formDiv)
		case "BottomEndX":
			FormDivBasicFieldToField(&(growthcurve2dribbonstartshape_.BottomEndX), formDiv)
		case "BottomEndY":
			FormDivBasicFieldToField(&(growthcurve2dribbonstartshape_.BottomEndY), formDiv)
		case "BottomRadiusX":
			FormDivBasicFieldToField(&(growthcurve2dribbonstartshape_.BottomRadiusX), formDiv)
		case "BottomRadiusY":
			FormDivBasicFieldToField(&(growthcurve2dribbonstartshape_.BottomRadiusY), formDiv)
		case "BottomXAxisRotation":
			FormDivBasicFieldToField(&(growthcurve2dribbonstartshape_.BottomXAxisRotation), formDiv)
		case "BottomLargeArcFlag":
			FormDivBasicFieldToField(&(growthcurve2dribbonstartshape_.BottomLargeArcFlag), formDiv)
		case "BottomSweepFlag":
			FormDivBasicFieldToField(&(growthcurve2dribbonstartshape_.BottomSweepFlag), formDiv)
		case "TopStartX":
			FormDivBasicFieldToField(&(growthcurve2dribbonstartshape_.TopStartX), formDiv)
		case "TopStartY":
			FormDivBasicFieldToField(&(growthcurve2dribbonstartshape_.TopStartY), formDiv)
		case "TopEndX":
			FormDivBasicFieldToField(&(growthcurve2dribbonstartshape_.TopEndX), formDiv)
		case "TopEndY":
			FormDivBasicFieldToField(&(growthcurve2dribbonstartshape_.TopEndY), formDiv)
		case "TopRadiusX":
			FormDivBasicFieldToField(&(growthcurve2dribbonstartshape_.TopRadiusX), formDiv)
		case "TopRadiusY":
			FormDivBasicFieldToField(&(growthcurve2dribbonstartshape_.TopRadiusY), formDiv)
		case "TopXAxisRotation":
			FormDivBasicFieldToField(&(growthcurve2dribbonstartshape_.TopXAxisRotation), formDiv)
		case "TopLargeArcFlag":
			FormDivBasicFieldToField(&(growthcurve2dribbonstartshape_.TopLargeArcFlag), formDiv)
		case "TopSweepFlag":
			FormDivBasicFieldToField(&(growthcurve2dribbonstartshape_.TopSweepFlag), formDiv)
		case "GrowthCurve2DRibbon:GrowthCurve2DRibbonStartShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the GrowthCurve2DRibbon instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target GrowthCurve2DRibbon instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.GrowthCurve2DRibbon](growthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest)
			targetGrowthCurve2DRibbonIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetGrowthCurve2DRibbonIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all GrowthCurve2DRibbon instances and update their GrowthCurve2DRibbonStartShapes slice
			for _growthcurve2dribbon := range *models.GetGongstructInstancesSetFromPointerType[*models.GrowthCurve2DRibbon](growthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(growthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest, _growthcurve2dribbon)
				
				// if GrowthCurve2DRibbon is selected
				if targetGrowthCurve2DRibbonIDs[id] {
					// ensure growthcurve2dribbonstartshape_ is in _growthcurve2dribbon.GrowthCurve2DRibbonStartShapes
					found := false
					for _, _b := range _growthcurve2dribbon.GrowthCurve2DRibbonStartShapes {
						if _b == growthcurve2dribbonstartshape_ {
							found = true
							break
						}
					}
					if !found {
						_growthcurve2dribbon.GrowthCurve2DRibbonStartShapes = append(_growthcurve2dribbon.GrowthCurve2DRibbonStartShapes, growthcurve2dribbonstartshape_)
						growthcurve2dribbonstartshapeFormCallback.probe.UpdateSliceOfPointersCallback(_growthcurve2dribbon, "GrowthCurve2DRibbonStartShapes", &_growthcurve2dribbon.GrowthCurve2DRibbonStartShapes)
					}
				} else {
					// ensure growthcurve2dribbonstartshape_ is NOT in _growthcurve2dribbon.GrowthCurve2DRibbonStartShapes
					idx := slices.Index(_growthcurve2dribbon.GrowthCurve2DRibbonStartShapes, growthcurve2dribbonstartshape_)
					if idx != -1 {
						_growthcurve2dribbon.GrowthCurve2DRibbonStartShapes = slices.Delete(_growthcurve2dribbon.GrowthCurve2DRibbonStartShapes, idx, idx+1)
						growthcurve2dribbonstartshapeFormCallback.probe.UpdateSliceOfPointersCallback(_growthcurve2dribbon, "GrowthCurve2DRibbonStartShapes", &_growthcurve2dribbon.GrowthCurve2DRibbonStartShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if growthcurve2dribbonstartshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		growthcurve2dribbonstartshape_.Unstage(growthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest)
	}

	growthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.GrowthCurve2DRibbonStartShape](
		growthcurve2dribbonstartshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if growthcurve2dribbonstartshapeFormCallback.CreationMode || growthcurve2dribbonstartshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		growthcurve2dribbonstartshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(growthcurve2dribbonstartshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__GrowthCurve2DRibbonStartShapeFormCallback(
			nil,
			growthcurve2dribbonstartshapeFormCallback.probe,
			newFormGroup,
		)
		growthcurve2dribbonstartshape := new(models.GrowthCurve2DRibbonStartShape)
		FillUpForm(growthcurve2dribbonstartshape, newFormGroup, growthcurve2dribbonstartshapeFormCallback.probe)
		growthcurve2dribbonstartshapeFormCallback.probe.formStage.Commit()
	}

	growthcurve2dribbonstartshapeFormCallback.probe.ux_tree()
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
func __gong__New__MidArcVectorShapeFormCallback(
	midarcvectorshape *models.MidArcVectorShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (midarcvectorshapeFormCallback *MidArcVectorShapeFormCallback) {
	midarcvectorshapeFormCallback = new(MidArcVectorShapeFormCallback)
	midarcvectorshapeFormCallback.probe = probe
	midarcvectorshapeFormCallback.midarcvectorshape = midarcvectorshape
	midarcvectorshapeFormCallback.formGroup = formGroup

	midarcvectorshapeFormCallback.CreationMode = (midarcvectorshape == nil)

	return
}

type MidArcVectorShapeFormCallback struct {
	midarcvectorshape *models.MidArcVectorShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (midarcvectorshapeFormCallback *MidArcVectorShapeFormCallback) OnSave() {
	midarcvectorshapeFormCallback.probe.stageOfInterest.Lock()
	defer midarcvectorshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("MidArcVectorShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	midarcvectorshapeFormCallback.probe.formStage.Checkout()

	if midarcvectorshapeFormCallback.midarcvectorshape == nil {
		midarcvectorshapeFormCallback.midarcvectorshape = new(models.MidArcVectorShape).Stage(midarcvectorshapeFormCallback.probe.stageOfInterest)
	}
	midarcvectorshape_ := midarcvectorshapeFormCallback.midarcvectorshape
	_ = midarcvectorshape_

	for _, formDiv := range midarcvectorshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(midarcvectorshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(midarcvectorshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(midarcvectorshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(midarcvectorshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(midarcvectorshape_.EndY), formDiv)
		case "MidArcVectorShapeGrid:MidArcVectorShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the MidArcVectorShapeGrid instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target MidArcVectorShapeGrid instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.MidArcVectorShapeGrid](midarcvectorshapeFormCallback.probe.stageOfInterest)
			targetMidArcVectorShapeGridIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetMidArcVectorShapeGridIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all MidArcVectorShapeGrid instances and update their MidArcVectorShapes slice
			for _midarcvectorshapegrid := range *models.GetGongstructInstancesSetFromPointerType[*models.MidArcVectorShapeGrid](midarcvectorshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(midarcvectorshapeFormCallback.probe.stageOfInterest, _midarcvectorshapegrid)
				
				// if MidArcVectorShapeGrid is selected
				if targetMidArcVectorShapeGridIDs[id] {
					// ensure midarcvectorshape_ is in _midarcvectorshapegrid.MidArcVectorShapes
					found := false
					for _, _b := range _midarcvectorshapegrid.MidArcVectorShapes {
						if _b == midarcvectorshape_ {
							found = true
							break
						}
					}
					if !found {
						_midarcvectorshapegrid.MidArcVectorShapes = append(_midarcvectorshapegrid.MidArcVectorShapes, midarcvectorshape_)
						midarcvectorshapeFormCallback.probe.UpdateSliceOfPointersCallback(_midarcvectorshapegrid, "MidArcVectorShapes", &_midarcvectorshapegrid.MidArcVectorShapes)
					}
				} else {
					// ensure midarcvectorshape_ is NOT in _midarcvectorshapegrid.MidArcVectorShapes
					idx := slices.Index(_midarcvectorshapegrid.MidArcVectorShapes, midarcvectorshape_)
					if idx != -1 {
						_midarcvectorshapegrid.MidArcVectorShapes = slices.Delete(_midarcvectorshapegrid.MidArcVectorShapes, idx, idx+1)
						midarcvectorshapeFormCallback.probe.UpdateSliceOfPointersCallback(_midarcvectorshapegrid, "MidArcVectorShapes", &_midarcvectorshapegrid.MidArcVectorShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if midarcvectorshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		midarcvectorshape_.Unstage(midarcvectorshapeFormCallback.probe.stageOfInterest)
	}

	midarcvectorshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.MidArcVectorShape](
		midarcvectorshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if midarcvectorshapeFormCallback.CreationMode || midarcvectorshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		midarcvectorshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(midarcvectorshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MidArcVectorShapeFormCallback(
			nil,
			midarcvectorshapeFormCallback.probe,
			newFormGroup,
		)
		midarcvectorshape := new(models.MidArcVectorShape)
		FillUpForm(midarcvectorshape, newFormGroup, midarcvectorshapeFormCallback.probe)
		midarcvectorshapeFormCallback.probe.formStage.Commit()
	}

	midarcvectorshapeFormCallback.probe.ux_tree()
}
func __gong__New__MidArcVectorShapeGridFormCallback(
	midarcvectorshapegrid *models.MidArcVectorShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (midarcvectorshapegridFormCallback *MidArcVectorShapeGridFormCallback) {
	midarcvectorshapegridFormCallback = new(MidArcVectorShapeGridFormCallback)
	midarcvectorshapegridFormCallback.probe = probe
	midarcvectorshapegridFormCallback.midarcvectorshapegrid = midarcvectorshapegrid
	midarcvectorshapegridFormCallback.formGroup = formGroup

	midarcvectorshapegridFormCallback.CreationMode = (midarcvectorshapegrid == nil)

	return
}

type MidArcVectorShapeGridFormCallback struct {
	midarcvectorshapegrid *models.MidArcVectorShapeGrid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (midarcvectorshapegridFormCallback *MidArcVectorShapeGridFormCallback) OnSave() {
	midarcvectorshapegridFormCallback.probe.stageOfInterest.Lock()
	defer midarcvectorshapegridFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("MidArcVectorShapeGridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	midarcvectorshapegridFormCallback.probe.formStage.Checkout()

	if midarcvectorshapegridFormCallback.midarcvectorshapegrid == nil {
		midarcvectorshapegridFormCallback.midarcvectorshapegrid = new(models.MidArcVectorShapeGrid).Stage(midarcvectorshapegridFormCallback.probe.stageOfInterest)
	}
	midarcvectorshapegrid_ := midarcvectorshapegridFormCallback.midarcvectorshapegrid
	_ = midarcvectorshapegrid_

	for _, formDiv := range midarcvectorshapegridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(midarcvectorshapegrid_.Name), formDiv)
		case "MidArcVectorShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.MidArcVectorShape](midarcvectorshapegridFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.MidArcVectorShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.MidArcVectorShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					midarcvectorshapegridFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.MidArcVectorShape](midarcvectorshapegridFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			midarcvectorshapegrid_.MidArcVectorShapes = instanceSlice
			midarcvectorshapegridFormCallback.probe.UpdateSliceOfPointersCallback(midarcvectorshapegrid_, "MidArcVectorShapes", &midarcvectorshapegrid_.MidArcVectorShapes)

		}
	}

	// manage the suppress operation
	if midarcvectorshapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		midarcvectorshapegrid_.Unstage(midarcvectorshapegridFormCallback.probe.stageOfInterest)
	}

	midarcvectorshapegridFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.MidArcVectorShapeGrid](
		midarcvectorshapegridFormCallback.probe,
	)

	// display a new form by reset the form stage
	if midarcvectorshapegridFormCallback.CreationMode || midarcvectorshapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		midarcvectorshapegridFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(midarcvectorshapegridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__MidArcVectorShapeGridFormCallback(
			nil,
			midarcvectorshapegridFormCallback.probe,
			newFormGroup,
		)
		midarcvectorshapegrid := new(models.MidArcVectorShapeGrid)
		FillUpForm(midarcvectorshapegrid, newFormGroup, midarcvectorshapegridFormCallback.probe)
		midarcvectorshapegridFormCallback.probe.formStage.Commit()
	}

	midarcvectorshapegridFormCallback.probe.ux_tree()
}
func __gong__New__PartiallyGrowthCurve2DRibbonFormCallback(
	partiallygrowthcurve2dribbon *models.PartiallyGrowthCurve2DRibbon,
	probe *Probe,
	formGroup *form.FormGroup,
) (partiallygrowthcurve2dribbonFormCallback *PartiallyGrowthCurve2DRibbonFormCallback) {
	partiallygrowthcurve2dribbonFormCallback = new(PartiallyGrowthCurve2DRibbonFormCallback)
	partiallygrowthcurve2dribbonFormCallback.probe = probe
	partiallygrowthcurve2dribbonFormCallback.partiallygrowthcurve2dribbon = partiallygrowthcurve2dribbon
	partiallygrowthcurve2dribbonFormCallback.formGroup = formGroup

	partiallygrowthcurve2dribbonFormCallback.CreationMode = (partiallygrowthcurve2dribbon == nil)

	return
}

type PartiallyGrowthCurve2DRibbonFormCallback struct {
	partiallygrowthcurve2dribbon *models.PartiallyGrowthCurve2DRibbon

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (partiallygrowthcurve2dribbonFormCallback *PartiallyGrowthCurve2DRibbonFormCallback) OnSave() {
	partiallygrowthcurve2dribbonFormCallback.probe.stageOfInterest.Lock()
	defer partiallygrowthcurve2dribbonFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PartiallyGrowthCurve2DRibbonFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	partiallygrowthcurve2dribbonFormCallback.probe.formStage.Checkout()

	if partiallygrowthcurve2dribbonFormCallback.partiallygrowthcurve2dribbon == nil {
		partiallygrowthcurve2dribbonFormCallback.partiallygrowthcurve2dribbon = new(models.PartiallyGrowthCurve2DRibbon).Stage(partiallygrowthcurve2dribbonFormCallback.probe.stageOfInterest)
	}
	partiallygrowthcurve2dribbon_ := partiallygrowthcurve2dribbonFormCallback.partiallygrowthcurve2dribbon
	_ = partiallygrowthcurve2dribbon_

	for _, formDiv := range partiallygrowthcurve2dribbonFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbon_.Name), formDiv)
		case "PartiallyGrowthCurve2DRibbonStartShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.PartiallyGrowthCurve2DRibbonStartShape](partiallygrowthcurve2dribbonFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.PartiallyGrowthCurve2DRibbonStartShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.PartiallyGrowthCurve2DRibbonStartShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					partiallygrowthcurve2dribbonFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.PartiallyGrowthCurve2DRibbonStartShape](partiallygrowthcurve2dribbonFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			partiallygrowthcurve2dribbon_.PartiallyGrowthCurve2DRibbonStartShapes = instanceSlice
			partiallygrowthcurve2dribbonFormCallback.probe.UpdateSliceOfPointersCallback(partiallygrowthcurve2dribbon_, "PartiallyGrowthCurve2DRibbonStartShapes", &partiallygrowthcurve2dribbon_.PartiallyGrowthCurve2DRibbonStartShapes)

		case "PartiallyGrowthCurve2DRibbonEndShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.PartiallyGrowthCurve2DRibbonEndShape](partiallygrowthcurve2dribbonFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.PartiallyGrowthCurve2DRibbonEndShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.PartiallyGrowthCurve2DRibbonEndShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					partiallygrowthcurve2dribbonFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.PartiallyGrowthCurve2DRibbonEndShape](partiallygrowthcurve2dribbonFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			partiallygrowthcurve2dribbon_.PartiallyGrowthCurve2DRibbonEndShapes = instanceSlice
			partiallygrowthcurve2dribbonFormCallback.probe.UpdateSliceOfPointersCallback(partiallygrowthcurve2dribbon_, "PartiallyGrowthCurve2DRibbonEndShapes", &partiallygrowthcurve2dribbon_.PartiallyGrowthCurve2DRibbonEndShapes)

		}
	}

	// manage the suppress operation
	if partiallygrowthcurve2dribbonFormCallback.formGroup.HasSuppressButtonBeenPressed {
		partiallygrowthcurve2dribbon_.Unstage(partiallygrowthcurve2dribbonFormCallback.probe.stageOfInterest)
	}

	partiallygrowthcurve2dribbonFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.PartiallyGrowthCurve2DRibbon](
		partiallygrowthcurve2dribbonFormCallback.probe,
	)

	// display a new form by reset the form stage
	if partiallygrowthcurve2dribbonFormCallback.CreationMode || partiallygrowthcurve2dribbonFormCallback.formGroup.HasSuppressButtonBeenPressed {
		partiallygrowthcurve2dribbonFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(partiallygrowthcurve2dribbonFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PartiallyGrowthCurve2DRibbonFormCallback(
			nil,
			partiallygrowthcurve2dribbonFormCallback.probe,
			newFormGroup,
		)
		partiallygrowthcurve2dribbon := new(models.PartiallyGrowthCurve2DRibbon)
		FillUpForm(partiallygrowthcurve2dribbon, newFormGroup, partiallygrowthcurve2dribbonFormCallback.probe)
		partiallygrowthcurve2dribbonFormCallback.probe.formStage.Commit()
	}

	partiallygrowthcurve2dribbonFormCallback.probe.ux_tree()
}
func __gong__New__PartiallyGrowthCurve2DRibbonEndShapeFormCallback(
	partiallygrowthcurve2dribbonendshape *models.PartiallyGrowthCurve2DRibbonEndShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (partiallygrowthcurve2dribbonendshapeFormCallback *PartiallyGrowthCurve2DRibbonEndShapeFormCallback) {
	partiallygrowthcurve2dribbonendshapeFormCallback = new(PartiallyGrowthCurve2DRibbonEndShapeFormCallback)
	partiallygrowthcurve2dribbonendshapeFormCallback.probe = probe
	partiallygrowthcurve2dribbonendshapeFormCallback.partiallygrowthcurve2dribbonendshape = partiallygrowthcurve2dribbonendshape
	partiallygrowthcurve2dribbonendshapeFormCallback.formGroup = formGroup

	partiallygrowthcurve2dribbonendshapeFormCallback.CreationMode = (partiallygrowthcurve2dribbonendshape == nil)

	return
}

type PartiallyGrowthCurve2DRibbonEndShapeFormCallback struct {
	partiallygrowthcurve2dribbonendshape *models.PartiallyGrowthCurve2DRibbonEndShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (partiallygrowthcurve2dribbonendshapeFormCallback *PartiallyGrowthCurve2DRibbonEndShapeFormCallback) OnSave() {
	partiallygrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest.Lock()
	defer partiallygrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PartiallyGrowthCurve2DRibbonEndShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	partiallygrowthcurve2dribbonendshapeFormCallback.probe.formStage.Checkout()

	if partiallygrowthcurve2dribbonendshapeFormCallback.partiallygrowthcurve2dribbonendshape == nil {
		partiallygrowthcurve2dribbonendshapeFormCallback.partiallygrowthcurve2dribbonendshape = new(models.PartiallyGrowthCurve2DRibbonEndShape).Stage(partiallygrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest)
	}
	partiallygrowthcurve2dribbonendshape_ := partiallygrowthcurve2dribbonendshapeFormCallback.partiallygrowthcurve2dribbonendshape
	_ = partiallygrowthcurve2dribbonendshape_

	for _, formDiv := range partiallygrowthcurve2dribbonendshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonendshape_.Name), formDiv)
		case "BottomStartX":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonendshape_.BottomStartX), formDiv)
		case "BottomStartY":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonendshape_.BottomStartY), formDiv)
		case "BottomEndX":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonendshape_.BottomEndX), formDiv)
		case "BottomEndY":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonendshape_.BottomEndY), formDiv)
		case "BottomRadiusX":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonendshape_.BottomRadiusX), formDiv)
		case "BottomRadiusY":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonendshape_.BottomRadiusY), formDiv)
		case "BottomXAxisRotation":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonendshape_.BottomXAxisRotation), formDiv)
		case "BottomLargeArcFlag":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonendshape_.BottomLargeArcFlag), formDiv)
		case "BottomSweepFlag":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonendshape_.BottomSweepFlag), formDiv)
		case "TopStartX":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonendshape_.TopStartX), formDiv)
		case "TopStartY":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonendshape_.TopStartY), formDiv)
		case "TopEndX":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonendshape_.TopEndX), formDiv)
		case "TopEndY":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonendshape_.TopEndY), formDiv)
		case "TopRadiusX":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonendshape_.TopRadiusX), formDiv)
		case "TopRadiusY":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonendshape_.TopRadiusY), formDiv)
		case "TopXAxisRotation":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonendshape_.TopXAxisRotation), formDiv)
		case "TopLargeArcFlag":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonendshape_.TopLargeArcFlag), formDiv)
		case "TopSweepFlag":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonendshape_.TopSweepFlag), formDiv)
		case "PartiallyGrowthCurve2DRibbon:PartiallyGrowthCurve2DRibbonEndShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the PartiallyGrowthCurve2DRibbon instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target PartiallyGrowthCurve2DRibbon instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.PartiallyGrowthCurve2DRibbon](partiallygrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest)
			targetPartiallyGrowthCurve2DRibbonIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetPartiallyGrowthCurve2DRibbonIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all PartiallyGrowthCurve2DRibbon instances and update their PartiallyGrowthCurve2DRibbonEndShapes slice
			for _partiallygrowthcurve2dribbon := range *models.GetGongstructInstancesSetFromPointerType[*models.PartiallyGrowthCurve2DRibbon](partiallygrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(partiallygrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest, _partiallygrowthcurve2dribbon)
				
				// if PartiallyGrowthCurve2DRibbon is selected
				if targetPartiallyGrowthCurve2DRibbonIDs[id] {
					// ensure partiallygrowthcurve2dribbonendshape_ is in _partiallygrowthcurve2dribbon.PartiallyGrowthCurve2DRibbonEndShapes
					found := false
					for _, _b := range _partiallygrowthcurve2dribbon.PartiallyGrowthCurve2DRibbonEndShapes {
						if _b == partiallygrowthcurve2dribbonendshape_ {
							found = true
							break
						}
					}
					if !found {
						_partiallygrowthcurve2dribbon.PartiallyGrowthCurve2DRibbonEndShapes = append(_partiallygrowthcurve2dribbon.PartiallyGrowthCurve2DRibbonEndShapes, partiallygrowthcurve2dribbonendshape_)
						partiallygrowthcurve2dribbonendshapeFormCallback.probe.UpdateSliceOfPointersCallback(_partiallygrowthcurve2dribbon, "PartiallyGrowthCurve2DRibbonEndShapes", &_partiallygrowthcurve2dribbon.PartiallyGrowthCurve2DRibbonEndShapes)
					}
				} else {
					// ensure partiallygrowthcurve2dribbonendshape_ is NOT in _partiallygrowthcurve2dribbon.PartiallyGrowthCurve2DRibbonEndShapes
					idx := slices.Index(_partiallygrowthcurve2dribbon.PartiallyGrowthCurve2DRibbonEndShapes, partiallygrowthcurve2dribbonendshape_)
					if idx != -1 {
						_partiallygrowthcurve2dribbon.PartiallyGrowthCurve2DRibbonEndShapes = slices.Delete(_partiallygrowthcurve2dribbon.PartiallyGrowthCurve2DRibbonEndShapes, idx, idx+1)
						partiallygrowthcurve2dribbonendshapeFormCallback.probe.UpdateSliceOfPointersCallback(_partiallygrowthcurve2dribbon, "PartiallyGrowthCurve2DRibbonEndShapes", &_partiallygrowthcurve2dribbon.PartiallyGrowthCurve2DRibbonEndShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if partiallygrowthcurve2dribbonendshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		partiallygrowthcurve2dribbonendshape_.Unstage(partiallygrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest)
	}

	partiallygrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.PartiallyGrowthCurve2DRibbonEndShape](
		partiallygrowthcurve2dribbonendshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if partiallygrowthcurve2dribbonendshapeFormCallback.CreationMode || partiallygrowthcurve2dribbonendshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		partiallygrowthcurve2dribbonendshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(partiallygrowthcurve2dribbonendshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PartiallyGrowthCurve2DRibbonEndShapeFormCallback(
			nil,
			partiallygrowthcurve2dribbonendshapeFormCallback.probe,
			newFormGroup,
		)
		partiallygrowthcurve2dribbonendshape := new(models.PartiallyGrowthCurve2DRibbonEndShape)
		FillUpForm(partiallygrowthcurve2dribbonendshape, newFormGroup, partiallygrowthcurve2dribbonendshapeFormCallback.probe)
		partiallygrowthcurve2dribbonendshapeFormCallback.probe.formStage.Commit()
	}

	partiallygrowthcurve2dribbonendshapeFormCallback.probe.ux_tree()
}
func __gong__New__PartiallyGrowthCurve2DRibbonStartShapeFormCallback(
	partiallygrowthcurve2dribbonstartshape *models.PartiallyGrowthCurve2DRibbonStartShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (partiallygrowthcurve2dribbonstartshapeFormCallback *PartiallyGrowthCurve2DRibbonStartShapeFormCallback) {
	partiallygrowthcurve2dribbonstartshapeFormCallback = new(PartiallyGrowthCurve2DRibbonStartShapeFormCallback)
	partiallygrowthcurve2dribbonstartshapeFormCallback.probe = probe
	partiallygrowthcurve2dribbonstartshapeFormCallback.partiallygrowthcurve2dribbonstartshape = partiallygrowthcurve2dribbonstartshape
	partiallygrowthcurve2dribbonstartshapeFormCallback.formGroup = formGroup

	partiallygrowthcurve2dribbonstartshapeFormCallback.CreationMode = (partiallygrowthcurve2dribbonstartshape == nil)

	return
}

type PartiallyGrowthCurve2DRibbonStartShapeFormCallback struct {
	partiallygrowthcurve2dribbonstartshape *models.PartiallyGrowthCurve2DRibbonStartShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (partiallygrowthcurve2dribbonstartshapeFormCallback *PartiallyGrowthCurve2DRibbonStartShapeFormCallback) OnSave() {
	partiallygrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest.Lock()
	defer partiallygrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PartiallyGrowthCurve2DRibbonStartShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	partiallygrowthcurve2dribbonstartshapeFormCallback.probe.formStage.Checkout()

	if partiallygrowthcurve2dribbonstartshapeFormCallback.partiallygrowthcurve2dribbonstartshape == nil {
		partiallygrowthcurve2dribbonstartshapeFormCallback.partiallygrowthcurve2dribbonstartshape = new(models.PartiallyGrowthCurve2DRibbonStartShape).Stage(partiallygrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest)
	}
	partiallygrowthcurve2dribbonstartshape_ := partiallygrowthcurve2dribbonstartshapeFormCallback.partiallygrowthcurve2dribbonstartshape
	_ = partiallygrowthcurve2dribbonstartshape_

	for _, formDiv := range partiallygrowthcurve2dribbonstartshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonstartshape_.Name), formDiv)
		case "BottomStartX":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonstartshape_.BottomStartX), formDiv)
		case "BottomStartY":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonstartshape_.BottomStartY), formDiv)
		case "BottomEndX":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonstartshape_.BottomEndX), formDiv)
		case "BottomEndY":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonstartshape_.BottomEndY), formDiv)
		case "BottomRadiusX":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonstartshape_.BottomRadiusX), formDiv)
		case "BottomRadiusY":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonstartshape_.BottomRadiusY), formDiv)
		case "BottomXAxisRotation":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonstartshape_.BottomXAxisRotation), formDiv)
		case "BottomLargeArcFlag":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonstartshape_.BottomLargeArcFlag), formDiv)
		case "BottomSweepFlag":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonstartshape_.BottomSweepFlag), formDiv)
		case "TopStartX":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonstartshape_.TopStartX), formDiv)
		case "TopStartY":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonstartshape_.TopStartY), formDiv)
		case "TopEndX":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonstartshape_.TopEndX), formDiv)
		case "TopEndY":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonstartshape_.TopEndY), formDiv)
		case "TopRadiusX":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonstartshape_.TopRadiusX), formDiv)
		case "TopRadiusY":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonstartshape_.TopRadiusY), formDiv)
		case "TopXAxisRotation":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonstartshape_.TopXAxisRotation), formDiv)
		case "TopLargeArcFlag":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonstartshape_.TopLargeArcFlag), formDiv)
		case "TopSweepFlag":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dribbonstartshape_.TopSweepFlag), formDiv)
		case "PartiallyGrowthCurve2DRibbon:PartiallyGrowthCurve2DRibbonStartShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the PartiallyGrowthCurve2DRibbon instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target PartiallyGrowthCurve2DRibbon instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.PartiallyGrowthCurve2DRibbon](partiallygrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest)
			targetPartiallyGrowthCurve2DRibbonIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetPartiallyGrowthCurve2DRibbonIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all PartiallyGrowthCurve2DRibbon instances and update their PartiallyGrowthCurve2DRibbonStartShapes slice
			for _partiallygrowthcurve2dribbon := range *models.GetGongstructInstancesSetFromPointerType[*models.PartiallyGrowthCurve2DRibbon](partiallygrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(partiallygrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest, _partiallygrowthcurve2dribbon)
				
				// if PartiallyGrowthCurve2DRibbon is selected
				if targetPartiallyGrowthCurve2DRibbonIDs[id] {
					// ensure partiallygrowthcurve2dribbonstartshape_ is in _partiallygrowthcurve2dribbon.PartiallyGrowthCurve2DRibbonStartShapes
					found := false
					for _, _b := range _partiallygrowthcurve2dribbon.PartiallyGrowthCurve2DRibbonStartShapes {
						if _b == partiallygrowthcurve2dribbonstartshape_ {
							found = true
							break
						}
					}
					if !found {
						_partiallygrowthcurve2dribbon.PartiallyGrowthCurve2DRibbonStartShapes = append(_partiallygrowthcurve2dribbon.PartiallyGrowthCurve2DRibbonStartShapes, partiallygrowthcurve2dribbonstartshape_)
						partiallygrowthcurve2dribbonstartshapeFormCallback.probe.UpdateSliceOfPointersCallback(_partiallygrowthcurve2dribbon, "PartiallyGrowthCurve2DRibbonStartShapes", &_partiallygrowthcurve2dribbon.PartiallyGrowthCurve2DRibbonStartShapes)
					}
				} else {
					// ensure partiallygrowthcurve2dribbonstartshape_ is NOT in _partiallygrowthcurve2dribbon.PartiallyGrowthCurve2DRibbonStartShapes
					idx := slices.Index(_partiallygrowthcurve2dribbon.PartiallyGrowthCurve2DRibbonStartShapes, partiallygrowthcurve2dribbonstartshape_)
					if idx != -1 {
						_partiallygrowthcurve2dribbon.PartiallyGrowthCurve2DRibbonStartShapes = slices.Delete(_partiallygrowthcurve2dribbon.PartiallyGrowthCurve2DRibbonStartShapes, idx, idx+1)
						partiallygrowthcurve2dribbonstartshapeFormCallback.probe.UpdateSliceOfPointersCallback(_partiallygrowthcurve2dribbon, "PartiallyGrowthCurve2DRibbonStartShapes", &_partiallygrowthcurve2dribbon.PartiallyGrowthCurve2DRibbonStartShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if partiallygrowthcurve2dribbonstartshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		partiallygrowthcurve2dribbonstartshape_.Unstage(partiallygrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest)
	}

	partiallygrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.PartiallyGrowthCurve2DRibbonStartShape](
		partiallygrowthcurve2dribbonstartshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if partiallygrowthcurve2dribbonstartshapeFormCallback.CreationMode || partiallygrowthcurve2dribbonstartshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		partiallygrowthcurve2dribbonstartshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(partiallygrowthcurve2dribbonstartshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PartiallyGrowthCurve2DRibbonStartShapeFormCallback(
			nil,
			partiallygrowthcurve2dribbonstartshapeFormCallback.probe,
			newFormGroup,
		)
		partiallygrowthcurve2dribbonstartshape := new(models.PartiallyGrowthCurve2DRibbonStartShape)
		FillUpForm(partiallygrowthcurve2dribbonstartshape, newFormGroup, partiallygrowthcurve2dribbonstartshapeFormCallback.probe)
		partiallygrowthcurve2dribbonstartshapeFormCallback.probe.formStage.Commit()
	}

	partiallygrowthcurve2dribbonstartshapeFormCallback.probe.ux_tree()
}
func __gong__New__PartiallyGrowthCurve2DTrajectoryFormCallback(
	partiallygrowthcurve2dtrajectory *models.PartiallyGrowthCurve2DTrajectory,
	probe *Probe,
	formGroup *form.FormGroup,
) (partiallygrowthcurve2dtrajectoryFormCallback *PartiallyGrowthCurve2DTrajectoryFormCallback) {
	partiallygrowthcurve2dtrajectoryFormCallback = new(PartiallyGrowthCurve2DTrajectoryFormCallback)
	partiallygrowthcurve2dtrajectoryFormCallback.probe = probe
	partiallygrowthcurve2dtrajectoryFormCallback.partiallygrowthcurve2dtrajectory = partiallygrowthcurve2dtrajectory
	partiallygrowthcurve2dtrajectoryFormCallback.formGroup = formGroup

	partiallygrowthcurve2dtrajectoryFormCallback.CreationMode = (partiallygrowthcurve2dtrajectory == nil)

	return
}

type PartiallyGrowthCurve2DTrajectoryFormCallback struct {
	partiallygrowthcurve2dtrajectory *models.PartiallyGrowthCurve2DTrajectory

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (partiallygrowthcurve2dtrajectoryFormCallback *PartiallyGrowthCurve2DTrajectoryFormCallback) OnSave() {
	partiallygrowthcurve2dtrajectoryFormCallback.probe.stageOfInterest.Lock()
	defer partiallygrowthcurve2dtrajectoryFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PartiallyGrowthCurve2DTrajectoryFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	partiallygrowthcurve2dtrajectoryFormCallback.probe.formStage.Checkout()

	if partiallygrowthcurve2dtrajectoryFormCallback.partiallygrowthcurve2dtrajectory == nil {
		partiallygrowthcurve2dtrajectoryFormCallback.partiallygrowthcurve2dtrajectory = new(models.PartiallyGrowthCurve2DTrajectory).Stage(partiallygrowthcurve2dtrajectoryFormCallback.probe.stageOfInterest)
	}
	partiallygrowthcurve2dtrajectory_ := partiallygrowthcurve2dtrajectoryFormCallback.partiallygrowthcurve2dtrajectory
	_ = partiallygrowthcurve2dtrajectory_

	for _, formDiv := range partiallygrowthcurve2dtrajectoryFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dtrajectory_.Name), formDiv)
		case "PartiallyGrowthCurve2DTrajectoryShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.PartiallyGrowthCurve2DTrajectoryShape](partiallygrowthcurve2dtrajectoryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.PartiallyGrowthCurve2DTrajectoryShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.PartiallyGrowthCurve2DTrajectoryShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					partiallygrowthcurve2dtrajectoryFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.PartiallyGrowthCurve2DTrajectoryShape](partiallygrowthcurve2dtrajectoryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			partiallygrowthcurve2dtrajectory_.PartiallyGrowthCurve2DTrajectoryShapes = instanceSlice
			partiallygrowthcurve2dtrajectoryFormCallback.probe.UpdateSliceOfPointersCallback(partiallygrowthcurve2dtrajectory_, "PartiallyGrowthCurve2DTrajectoryShapes", &partiallygrowthcurve2dtrajectory_.PartiallyGrowthCurve2DTrajectoryShapes)

		}
	}

	// manage the suppress operation
	if partiallygrowthcurve2dtrajectoryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		partiallygrowthcurve2dtrajectory_.Unstage(partiallygrowthcurve2dtrajectoryFormCallback.probe.stageOfInterest)
	}

	partiallygrowthcurve2dtrajectoryFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.PartiallyGrowthCurve2DTrajectory](
		partiallygrowthcurve2dtrajectoryFormCallback.probe,
	)

	// display a new form by reset the form stage
	if partiallygrowthcurve2dtrajectoryFormCallback.CreationMode || partiallygrowthcurve2dtrajectoryFormCallback.formGroup.HasSuppressButtonBeenPressed {
		partiallygrowthcurve2dtrajectoryFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(partiallygrowthcurve2dtrajectoryFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PartiallyGrowthCurve2DTrajectoryFormCallback(
			nil,
			partiallygrowthcurve2dtrajectoryFormCallback.probe,
			newFormGroup,
		)
		partiallygrowthcurve2dtrajectory := new(models.PartiallyGrowthCurve2DTrajectory)
		FillUpForm(partiallygrowthcurve2dtrajectory, newFormGroup, partiallygrowthcurve2dtrajectoryFormCallback.probe)
		partiallygrowthcurve2dtrajectoryFormCallback.probe.formStage.Commit()
	}

	partiallygrowthcurve2dtrajectoryFormCallback.probe.ux_tree()
}
func __gong__New__PartiallyGrowthCurve2DTrajectoryP1CurveShapeFormCallback(
	partiallygrowthcurve2dtrajectoryp1curveshape *models.PartiallyGrowthCurve2DTrajectoryP1CurveShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback *PartiallyGrowthCurve2DTrajectoryP1CurveShapeFormCallback) {
	partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback = new(PartiallyGrowthCurve2DTrajectoryP1CurveShapeFormCallback)
	partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback.probe = probe
	partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback.partiallygrowthcurve2dtrajectoryp1curveshape = partiallygrowthcurve2dtrajectoryp1curveshape
	partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback.formGroup = formGroup

	partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback.CreationMode = (partiallygrowthcurve2dtrajectoryp1curveshape == nil)

	return
}

type PartiallyGrowthCurve2DTrajectoryP1CurveShapeFormCallback struct {
	partiallygrowthcurve2dtrajectoryp1curveshape *models.PartiallyGrowthCurve2DTrajectoryP1CurveShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback *PartiallyGrowthCurve2DTrajectoryP1CurveShapeFormCallback) OnSave() {
	partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback.probe.stageOfInterest.Lock()
	defer partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PartiallyGrowthCurve2DTrajectoryP1CurveShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback.probe.formStage.Checkout()

	if partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback.partiallygrowthcurve2dtrajectoryp1curveshape == nil {
		partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback.partiallygrowthcurve2dtrajectoryp1curveshape = new(models.PartiallyGrowthCurve2DTrajectoryP1CurveShape).Stage(partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback.probe.stageOfInterest)
	}
	partiallygrowthcurve2dtrajectoryp1curveshape_ := partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback.partiallygrowthcurve2dtrajectoryp1curveshape
	_ = partiallygrowthcurve2dtrajectoryp1curveshape_

	for _, formDiv := range partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dtrajectoryp1curveshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dtrajectoryp1curveshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dtrajectoryp1curveshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dtrajectoryp1curveshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dtrajectoryp1curveshape_.EndY), formDiv)
		case "PartiallyGrowthCurve2DTrajectoryP1P2:P1CurveShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the PartiallyGrowthCurve2DTrajectoryP1P2 instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target PartiallyGrowthCurve2DTrajectoryP1P2 instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.PartiallyGrowthCurve2DTrajectoryP1P2](partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback.probe.stageOfInterest)
			targetPartiallyGrowthCurve2DTrajectoryP1P2IDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetPartiallyGrowthCurve2DTrajectoryP1P2IDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all PartiallyGrowthCurve2DTrajectoryP1P2 instances and update their P1CurveShapes slice
			for _partiallygrowthcurve2dtrajectoryp1p2 := range *models.GetGongstructInstancesSetFromPointerType[*models.PartiallyGrowthCurve2DTrajectoryP1P2](partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback.probe.stageOfInterest, _partiallygrowthcurve2dtrajectoryp1p2)
				
				// if PartiallyGrowthCurve2DTrajectoryP1P2 is selected
				if targetPartiallyGrowthCurve2DTrajectoryP1P2IDs[id] {
					// ensure partiallygrowthcurve2dtrajectoryp1curveshape_ is in _partiallygrowthcurve2dtrajectoryp1p2.P1CurveShapes
					found := false
					for _, _b := range _partiallygrowthcurve2dtrajectoryp1p2.P1CurveShapes {
						if _b == partiallygrowthcurve2dtrajectoryp1curveshape_ {
							found = true
							break
						}
					}
					if !found {
						_partiallygrowthcurve2dtrajectoryp1p2.P1CurveShapes = append(_partiallygrowthcurve2dtrajectoryp1p2.P1CurveShapes, partiallygrowthcurve2dtrajectoryp1curveshape_)
						partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback.probe.UpdateSliceOfPointersCallback(_partiallygrowthcurve2dtrajectoryp1p2, "P1CurveShapes", &_partiallygrowthcurve2dtrajectoryp1p2.P1CurveShapes)
					}
				} else {
					// ensure partiallygrowthcurve2dtrajectoryp1curveshape_ is NOT in _partiallygrowthcurve2dtrajectoryp1p2.P1CurveShapes
					idx := slices.Index(_partiallygrowthcurve2dtrajectoryp1p2.P1CurveShapes, partiallygrowthcurve2dtrajectoryp1curveshape_)
					if idx != -1 {
						_partiallygrowthcurve2dtrajectoryp1p2.P1CurveShapes = slices.Delete(_partiallygrowthcurve2dtrajectoryp1p2.P1CurveShapes, idx, idx+1)
						partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback.probe.UpdateSliceOfPointersCallback(_partiallygrowthcurve2dtrajectoryp1p2, "P1CurveShapes", &_partiallygrowthcurve2dtrajectoryp1p2.P1CurveShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		partiallygrowthcurve2dtrajectoryp1curveshape_.Unstage(partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback.probe.stageOfInterest)
	}

	partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.PartiallyGrowthCurve2DTrajectoryP1CurveShape](
		partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback.CreationMode || partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PartiallyGrowthCurve2DTrajectoryP1CurveShapeFormCallback(
			nil,
			partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback.probe,
			newFormGroup,
		)
		partiallygrowthcurve2dtrajectoryp1curveshape := new(models.PartiallyGrowthCurve2DTrajectoryP1CurveShape)
		FillUpForm(partiallygrowthcurve2dtrajectoryp1curveshape, newFormGroup, partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback.probe)
		partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback.probe.formStage.Commit()
	}

	partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback.probe.ux_tree()
}
func __gong__New__PartiallyGrowthCurve2DTrajectoryP1P2FormCallback(
	partiallygrowthcurve2dtrajectoryp1p2 *models.PartiallyGrowthCurve2DTrajectoryP1P2,
	probe *Probe,
	formGroup *form.FormGroup,
) (partiallygrowthcurve2dtrajectoryp1p2FormCallback *PartiallyGrowthCurve2DTrajectoryP1P2FormCallback) {
	partiallygrowthcurve2dtrajectoryp1p2FormCallback = new(PartiallyGrowthCurve2DTrajectoryP1P2FormCallback)
	partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe = probe
	partiallygrowthcurve2dtrajectoryp1p2FormCallback.partiallygrowthcurve2dtrajectoryp1p2 = partiallygrowthcurve2dtrajectoryp1p2
	partiallygrowthcurve2dtrajectoryp1p2FormCallback.formGroup = formGroup

	partiallygrowthcurve2dtrajectoryp1p2FormCallback.CreationMode = (partiallygrowthcurve2dtrajectoryp1p2 == nil)

	return
}

type PartiallyGrowthCurve2DTrajectoryP1P2FormCallback struct {
	partiallygrowthcurve2dtrajectoryp1p2 *models.PartiallyGrowthCurve2DTrajectoryP1P2

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (partiallygrowthcurve2dtrajectoryp1p2FormCallback *PartiallyGrowthCurve2DTrajectoryP1P2FormCallback) OnSave() {
	partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe.stageOfInterest.Lock()
	defer partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PartiallyGrowthCurve2DTrajectoryP1P2FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe.formStage.Checkout()

	if partiallygrowthcurve2dtrajectoryp1p2FormCallback.partiallygrowthcurve2dtrajectoryp1p2 == nil {
		partiallygrowthcurve2dtrajectoryp1p2FormCallback.partiallygrowthcurve2dtrajectoryp1p2 = new(models.PartiallyGrowthCurve2DTrajectoryP1P2).Stage(partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe.stageOfInterest)
	}
	partiallygrowthcurve2dtrajectoryp1p2_ := partiallygrowthcurve2dtrajectoryp1p2FormCallback.partiallygrowthcurve2dtrajectoryp1p2
	_ = partiallygrowthcurve2dtrajectoryp1p2_

	for _, formDiv := range partiallygrowthcurve2dtrajectoryp1p2FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dtrajectoryp1p2_.Name), formDiv)
		case "P1PointShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.PartiallyGrowthCurve2DTrajectoryP1PointShape](partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.PartiallyGrowthCurve2DTrajectoryP1PointShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.PartiallyGrowthCurve2DTrajectoryP1PointShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.PartiallyGrowthCurve2DTrajectoryP1PointShape](partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			partiallygrowthcurve2dtrajectoryp1p2_.P1PointShapes = instanceSlice
			partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe.UpdateSliceOfPointersCallback(partiallygrowthcurve2dtrajectoryp1p2_, "P1PointShapes", &partiallygrowthcurve2dtrajectoryp1p2_.P1PointShapes)

		case "P2PointShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.PartiallyGrowthCurve2DTrajectoryP2PointShape](partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.PartiallyGrowthCurve2DTrajectoryP2PointShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.PartiallyGrowthCurve2DTrajectoryP2PointShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.PartiallyGrowthCurve2DTrajectoryP2PointShape](partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			partiallygrowthcurve2dtrajectoryp1p2_.P2PointShapes = instanceSlice
			partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe.UpdateSliceOfPointersCallback(partiallygrowthcurve2dtrajectoryp1p2_, "P2PointShapes", &partiallygrowthcurve2dtrajectoryp1p2_.P2PointShapes)

		case "P1CurveShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.PartiallyGrowthCurve2DTrajectoryP1CurveShape](partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.PartiallyGrowthCurve2DTrajectoryP1CurveShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.PartiallyGrowthCurve2DTrajectoryP1CurveShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.PartiallyGrowthCurve2DTrajectoryP1CurveShape](partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			partiallygrowthcurve2dtrajectoryp1p2_.P1CurveShapes = instanceSlice
			partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe.UpdateSliceOfPointersCallback(partiallygrowthcurve2dtrajectoryp1p2_, "P1CurveShapes", &partiallygrowthcurve2dtrajectoryp1p2_.P1CurveShapes)

		case "P2CurveShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.PartiallyGrowthCurve2DTrajectoryP2CurveShape](partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.PartiallyGrowthCurve2DTrajectoryP2CurveShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.PartiallyGrowthCurve2DTrajectoryP2CurveShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.PartiallyGrowthCurve2DTrajectoryP2CurveShape](partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			partiallygrowthcurve2dtrajectoryp1p2_.P2CurveShapes = instanceSlice
			partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe.UpdateSliceOfPointersCallback(partiallygrowthcurve2dtrajectoryp1p2_, "P2CurveShapes", &partiallygrowthcurve2dtrajectoryp1p2_.P2CurveShapes)

		case "P1P2PairLineShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape](partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape](partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			partiallygrowthcurve2dtrajectoryp1p2_.P1P2PairLineShapes = instanceSlice
			partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe.UpdateSliceOfPointersCallback(partiallygrowthcurve2dtrajectoryp1p2_, "P1P2PairLineShapes", &partiallygrowthcurve2dtrajectoryp1p2_.P1P2PairLineShapes)

		}
	}

	// manage the suppress operation
	if partiallygrowthcurve2dtrajectoryp1p2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		partiallygrowthcurve2dtrajectoryp1p2_.Unstage(partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe.stageOfInterest)
	}

	partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.PartiallyGrowthCurve2DTrajectoryP1P2](
		partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe,
	)

	// display a new form by reset the form stage
	if partiallygrowthcurve2dtrajectoryp1p2FormCallback.CreationMode || partiallygrowthcurve2dtrajectoryp1p2FormCallback.formGroup.HasSuppressButtonBeenPressed {
		partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PartiallyGrowthCurve2DTrajectoryP1P2FormCallback(
			nil,
			partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe,
			newFormGroup,
		)
		partiallygrowthcurve2dtrajectoryp1p2 := new(models.PartiallyGrowthCurve2DTrajectoryP1P2)
		FillUpForm(partiallygrowthcurve2dtrajectoryp1p2, newFormGroup, partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe)
		partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe.formStage.Commit()
	}

	partiallygrowthcurve2dtrajectoryp1p2FormCallback.probe.ux_tree()
}
func __gong__New__PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeFormCallback(
	partiallygrowthcurve2dtrajectoryp1p2pairlineshape *models.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeFormCallback) {
	partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback = new(PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeFormCallback)
	partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback.probe = probe
	partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback.partiallygrowthcurve2dtrajectoryp1p2pairlineshape = partiallygrowthcurve2dtrajectoryp1p2pairlineshape
	partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback.formGroup = formGroup

	partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback.CreationMode = (partiallygrowthcurve2dtrajectoryp1p2pairlineshape == nil)

	return
}

type PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeFormCallback struct {
	partiallygrowthcurve2dtrajectoryp1p2pairlineshape *models.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeFormCallback) OnSave() {
	partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback.probe.stageOfInterest.Lock()
	defer partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback.probe.formStage.Checkout()

	if partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback.partiallygrowthcurve2dtrajectoryp1p2pairlineshape == nil {
		partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback.partiallygrowthcurve2dtrajectoryp1p2pairlineshape = new(models.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape).Stage(partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback.probe.stageOfInterest)
	}
	partiallygrowthcurve2dtrajectoryp1p2pairlineshape_ := partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback.partiallygrowthcurve2dtrajectoryp1p2pairlineshape
	_ = partiallygrowthcurve2dtrajectoryp1p2pairlineshape_

	for _, formDiv := range partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dtrajectoryp1p2pairlineshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dtrajectoryp1p2pairlineshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dtrajectoryp1p2pairlineshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dtrajectoryp1p2pairlineshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dtrajectoryp1p2pairlineshape_.EndY), formDiv)
		case "PartiallyGrowthCurve2DTrajectoryP1P2:P1P2PairLineShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the PartiallyGrowthCurve2DTrajectoryP1P2 instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target PartiallyGrowthCurve2DTrajectoryP1P2 instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.PartiallyGrowthCurve2DTrajectoryP1P2](partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback.probe.stageOfInterest)
			targetPartiallyGrowthCurve2DTrajectoryP1P2IDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetPartiallyGrowthCurve2DTrajectoryP1P2IDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all PartiallyGrowthCurve2DTrajectoryP1P2 instances and update their P1P2PairLineShapes slice
			for _partiallygrowthcurve2dtrajectoryp1p2 := range *models.GetGongstructInstancesSetFromPointerType[*models.PartiallyGrowthCurve2DTrajectoryP1P2](partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback.probe.stageOfInterest, _partiallygrowthcurve2dtrajectoryp1p2)
				
				// if PartiallyGrowthCurve2DTrajectoryP1P2 is selected
				if targetPartiallyGrowthCurve2DTrajectoryP1P2IDs[id] {
					// ensure partiallygrowthcurve2dtrajectoryp1p2pairlineshape_ is in _partiallygrowthcurve2dtrajectoryp1p2.P1P2PairLineShapes
					found := false
					for _, _b := range _partiallygrowthcurve2dtrajectoryp1p2.P1P2PairLineShapes {
						if _b == partiallygrowthcurve2dtrajectoryp1p2pairlineshape_ {
							found = true
							break
						}
					}
					if !found {
						_partiallygrowthcurve2dtrajectoryp1p2.P1P2PairLineShapes = append(_partiallygrowthcurve2dtrajectoryp1p2.P1P2PairLineShapes, partiallygrowthcurve2dtrajectoryp1p2pairlineshape_)
						partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback.probe.UpdateSliceOfPointersCallback(_partiallygrowthcurve2dtrajectoryp1p2, "P1P2PairLineShapes", &_partiallygrowthcurve2dtrajectoryp1p2.P1P2PairLineShapes)
					}
				} else {
					// ensure partiallygrowthcurve2dtrajectoryp1p2pairlineshape_ is NOT in _partiallygrowthcurve2dtrajectoryp1p2.P1P2PairLineShapes
					idx := slices.Index(_partiallygrowthcurve2dtrajectoryp1p2.P1P2PairLineShapes, partiallygrowthcurve2dtrajectoryp1p2pairlineshape_)
					if idx != -1 {
						_partiallygrowthcurve2dtrajectoryp1p2.P1P2PairLineShapes = slices.Delete(_partiallygrowthcurve2dtrajectoryp1p2.P1P2PairLineShapes, idx, idx+1)
						partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback.probe.UpdateSliceOfPointersCallback(_partiallygrowthcurve2dtrajectoryp1p2, "P1P2PairLineShapes", &_partiallygrowthcurve2dtrajectoryp1p2.P1P2PairLineShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		partiallygrowthcurve2dtrajectoryp1p2pairlineshape_.Unstage(partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback.probe.stageOfInterest)
	}

	partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape](
		partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback.CreationMode || partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeFormCallback(
			nil,
			partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback.probe,
			newFormGroup,
		)
		partiallygrowthcurve2dtrajectoryp1p2pairlineshape := new(models.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape)
		FillUpForm(partiallygrowthcurve2dtrajectoryp1p2pairlineshape, newFormGroup, partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback.probe)
		partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback.probe.formStage.Commit()
	}

	partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback.probe.ux_tree()
}
func __gong__New__PartiallyGrowthCurve2DTrajectoryP1PointShapeFormCallback(
	partiallygrowthcurve2dtrajectoryp1pointshape *models.PartiallyGrowthCurve2DTrajectoryP1PointShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback *PartiallyGrowthCurve2DTrajectoryP1PointShapeFormCallback) {
	partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback = new(PartiallyGrowthCurve2DTrajectoryP1PointShapeFormCallback)
	partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback.probe = probe
	partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback.partiallygrowthcurve2dtrajectoryp1pointshape = partiallygrowthcurve2dtrajectoryp1pointshape
	partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback.formGroup = formGroup

	partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback.CreationMode = (partiallygrowthcurve2dtrajectoryp1pointshape == nil)

	return
}

type PartiallyGrowthCurve2DTrajectoryP1PointShapeFormCallback struct {
	partiallygrowthcurve2dtrajectoryp1pointshape *models.PartiallyGrowthCurve2DTrajectoryP1PointShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback *PartiallyGrowthCurve2DTrajectoryP1PointShapeFormCallback) OnSave() {
	partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback.probe.stageOfInterest.Lock()
	defer partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PartiallyGrowthCurve2DTrajectoryP1PointShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback.probe.formStage.Checkout()

	if partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback.partiallygrowthcurve2dtrajectoryp1pointshape == nil {
		partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback.partiallygrowthcurve2dtrajectoryp1pointshape = new(models.PartiallyGrowthCurve2DTrajectoryP1PointShape).Stage(partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback.probe.stageOfInterest)
	}
	partiallygrowthcurve2dtrajectoryp1pointshape_ := partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback.partiallygrowthcurve2dtrajectoryp1pointshape
	_ = partiallygrowthcurve2dtrajectoryp1pointshape_

	for _, formDiv := range partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dtrajectoryp1pointshape_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dtrajectoryp1pointshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dtrajectoryp1pointshape_.Y), formDiv)
		case "PartiallyGrowthCurve2DTrajectoryP1P2:P1PointShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the PartiallyGrowthCurve2DTrajectoryP1P2 instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target PartiallyGrowthCurve2DTrajectoryP1P2 instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.PartiallyGrowthCurve2DTrajectoryP1P2](partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback.probe.stageOfInterest)
			targetPartiallyGrowthCurve2DTrajectoryP1P2IDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetPartiallyGrowthCurve2DTrajectoryP1P2IDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all PartiallyGrowthCurve2DTrajectoryP1P2 instances and update their P1PointShapes slice
			for _partiallygrowthcurve2dtrajectoryp1p2 := range *models.GetGongstructInstancesSetFromPointerType[*models.PartiallyGrowthCurve2DTrajectoryP1P2](partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback.probe.stageOfInterest, _partiallygrowthcurve2dtrajectoryp1p2)
				
				// if PartiallyGrowthCurve2DTrajectoryP1P2 is selected
				if targetPartiallyGrowthCurve2DTrajectoryP1P2IDs[id] {
					// ensure partiallygrowthcurve2dtrajectoryp1pointshape_ is in _partiallygrowthcurve2dtrajectoryp1p2.P1PointShapes
					found := false
					for _, _b := range _partiallygrowthcurve2dtrajectoryp1p2.P1PointShapes {
						if _b == partiallygrowthcurve2dtrajectoryp1pointshape_ {
							found = true
							break
						}
					}
					if !found {
						_partiallygrowthcurve2dtrajectoryp1p2.P1PointShapes = append(_partiallygrowthcurve2dtrajectoryp1p2.P1PointShapes, partiallygrowthcurve2dtrajectoryp1pointshape_)
						partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback.probe.UpdateSliceOfPointersCallback(_partiallygrowthcurve2dtrajectoryp1p2, "P1PointShapes", &_partiallygrowthcurve2dtrajectoryp1p2.P1PointShapes)
					}
				} else {
					// ensure partiallygrowthcurve2dtrajectoryp1pointshape_ is NOT in _partiallygrowthcurve2dtrajectoryp1p2.P1PointShapes
					idx := slices.Index(_partiallygrowthcurve2dtrajectoryp1p2.P1PointShapes, partiallygrowthcurve2dtrajectoryp1pointshape_)
					if idx != -1 {
						_partiallygrowthcurve2dtrajectoryp1p2.P1PointShapes = slices.Delete(_partiallygrowthcurve2dtrajectoryp1p2.P1PointShapes, idx, idx+1)
						partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback.probe.UpdateSliceOfPointersCallback(_partiallygrowthcurve2dtrajectoryp1p2, "P1PointShapes", &_partiallygrowthcurve2dtrajectoryp1p2.P1PointShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		partiallygrowthcurve2dtrajectoryp1pointshape_.Unstage(partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback.probe.stageOfInterest)
	}

	partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.PartiallyGrowthCurve2DTrajectoryP1PointShape](
		partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback.CreationMode || partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PartiallyGrowthCurve2DTrajectoryP1PointShapeFormCallback(
			nil,
			partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback.probe,
			newFormGroup,
		)
		partiallygrowthcurve2dtrajectoryp1pointshape := new(models.PartiallyGrowthCurve2DTrajectoryP1PointShape)
		FillUpForm(partiallygrowthcurve2dtrajectoryp1pointshape, newFormGroup, partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback.probe)
		partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback.probe.formStage.Commit()
	}

	partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback.probe.ux_tree()
}
func __gong__New__PartiallyGrowthCurve2DTrajectoryP2CurveShapeFormCallback(
	partiallygrowthcurve2dtrajectoryp2curveshape *models.PartiallyGrowthCurve2DTrajectoryP2CurveShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback *PartiallyGrowthCurve2DTrajectoryP2CurveShapeFormCallback) {
	partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback = new(PartiallyGrowthCurve2DTrajectoryP2CurveShapeFormCallback)
	partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback.probe = probe
	partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback.partiallygrowthcurve2dtrajectoryp2curveshape = partiallygrowthcurve2dtrajectoryp2curveshape
	partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback.formGroup = formGroup

	partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback.CreationMode = (partiallygrowthcurve2dtrajectoryp2curveshape == nil)

	return
}

type PartiallyGrowthCurve2DTrajectoryP2CurveShapeFormCallback struct {
	partiallygrowthcurve2dtrajectoryp2curveshape *models.PartiallyGrowthCurve2DTrajectoryP2CurveShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback *PartiallyGrowthCurve2DTrajectoryP2CurveShapeFormCallback) OnSave() {
	partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback.probe.stageOfInterest.Lock()
	defer partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PartiallyGrowthCurve2DTrajectoryP2CurveShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback.probe.formStage.Checkout()

	if partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback.partiallygrowthcurve2dtrajectoryp2curveshape == nil {
		partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback.partiallygrowthcurve2dtrajectoryp2curveshape = new(models.PartiallyGrowthCurve2DTrajectoryP2CurveShape).Stage(partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback.probe.stageOfInterest)
	}
	partiallygrowthcurve2dtrajectoryp2curveshape_ := partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback.partiallygrowthcurve2dtrajectoryp2curveshape
	_ = partiallygrowthcurve2dtrajectoryp2curveshape_

	for _, formDiv := range partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dtrajectoryp2curveshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dtrajectoryp2curveshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dtrajectoryp2curveshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dtrajectoryp2curveshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dtrajectoryp2curveshape_.EndY), formDiv)
		case "PartiallyGrowthCurve2DTrajectoryP1P2:P2CurveShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the PartiallyGrowthCurve2DTrajectoryP1P2 instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target PartiallyGrowthCurve2DTrajectoryP1P2 instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.PartiallyGrowthCurve2DTrajectoryP1P2](partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback.probe.stageOfInterest)
			targetPartiallyGrowthCurve2DTrajectoryP1P2IDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetPartiallyGrowthCurve2DTrajectoryP1P2IDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all PartiallyGrowthCurve2DTrajectoryP1P2 instances and update their P2CurveShapes slice
			for _partiallygrowthcurve2dtrajectoryp1p2 := range *models.GetGongstructInstancesSetFromPointerType[*models.PartiallyGrowthCurve2DTrajectoryP1P2](partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback.probe.stageOfInterest, _partiallygrowthcurve2dtrajectoryp1p2)
				
				// if PartiallyGrowthCurve2DTrajectoryP1P2 is selected
				if targetPartiallyGrowthCurve2DTrajectoryP1P2IDs[id] {
					// ensure partiallygrowthcurve2dtrajectoryp2curveshape_ is in _partiallygrowthcurve2dtrajectoryp1p2.P2CurveShapes
					found := false
					for _, _b := range _partiallygrowthcurve2dtrajectoryp1p2.P2CurveShapes {
						if _b == partiallygrowthcurve2dtrajectoryp2curveshape_ {
							found = true
							break
						}
					}
					if !found {
						_partiallygrowthcurve2dtrajectoryp1p2.P2CurveShapes = append(_partiallygrowthcurve2dtrajectoryp1p2.P2CurveShapes, partiallygrowthcurve2dtrajectoryp2curveshape_)
						partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback.probe.UpdateSliceOfPointersCallback(_partiallygrowthcurve2dtrajectoryp1p2, "P2CurveShapes", &_partiallygrowthcurve2dtrajectoryp1p2.P2CurveShapes)
					}
				} else {
					// ensure partiallygrowthcurve2dtrajectoryp2curveshape_ is NOT in _partiallygrowthcurve2dtrajectoryp1p2.P2CurveShapes
					idx := slices.Index(_partiallygrowthcurve2dtrajectoryp1p2.P2CurveShapes, partiallygrowthcurve2dtrajectoryp2curveshape_)
					if idx != -1 {
						_partiallygrowthcurve2dtrajectoryp1p2.P2CurveShapes = slices.Delete(_partiallygrowthcurve2dtrajectoryp1p2.P2CurveShapes, idx, idx+1)
						partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback.probe.UpdateSliceOfPointersCallback(_partiallygrowthcurve2dtrajectoryp1p2, "P2CurveShapes", &_partiallygrowthcurve2dtrajectoryp1p2.P2CurveShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		partiallygrowthcurve2dtrajectoryp2curveshape_.Unstage(partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback.probe.stageOfInterest)
	}

	partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.PartiallyGrowthCurve2DTrajectoryP2CurveShape](
		partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback.CreationMode || partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PartiallyGrowthCurve2DTrajectoryP2CurveShapeFormCallback(
			nil,
			partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback.probe,
			newFormGroup,
		)
		partiallygrowthcurve2dtrajectoryp2curveshape := new(models.PartiallyGrowthCurve2DTrajectoryP2CurveShape)
		FillUpForm(partiallygrowthcurve2dtrajectoryp2curveshape, newFormGroup, partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback.probe)
		partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback.probe.formStage.Commit()
	}

	partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback.probe.ux_tree()
}
func __gong__New__PartiallyGrowthCurve2DTrajectoryP2PointShapeFormCallback(
	partiallygrowthcurve2dtrajectoryp2pointshape *models.PartiallyGrowthCurve2DTrajectoryP2PointShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback *PartiallyGrowthCurve2DTrajectoryP2PointShapeFormCallback) {
	partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback = new(PartiallyGrowthCurve2DTrajectoryP2PointShapeFormCallback)
	partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback.probe = probe
	partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback.partiallygrowthcurve2dtrajectoryp2pointshape = partiallygrowthcurve2dtrajectoryp2pointshape
	partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback.formGroup = formGroup

	partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback.CreationMode = (partiallygrowthcurve2dtrajectoryp2pointshape == nil)

	return
}

type PartiallyGrowthCurve2DTrajectoryP2PointShapeFormCallback struct {
	partiallygrowthcurve2dtrajectoryp2pointshape *models.PartiallyGrowthCurve2DTrajectoryP2PointShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback *PartiallyGrowthCurve2DTrajectoryP2PointShapeFormCallback) OnSave() {
	partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback.probe.stageOfInterest.Lock()
	defer partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PartiallyGrowthCurve2DTrajectoryP2PointShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback.probe.formStage.Checkout()

	if partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback.partiallygrowthcurve2dtrajectoryp2pointshape == nil {
		partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback.partiallygrowthcurve2dtrajectoryp2pointshape = new(models.PartiallyGrowthCurve2DTrajectoryP2PointShape).Stage(partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback.probe.stageOfInterest)
	}
	partiallygrowthcurve2dtrajectoryp2pointshape_ := partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback.partiallygrowthcurve2dtrajectoryp2pointshape
	_ = partiallygrowthcurve2dtrajectoryp2pointshape_

	for _, formDiv := range partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dtrajectoryp2pointshape_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dtrajectoryp2pointshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dtrajectoryp2pointshape_.Y), formDiv)
		case "PartiallyGrowthCurve2DTrajectoryP1P2:P2PointShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the PartiallyGrowthCurve2DTrajectoryP1P2 instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target PartiallyGrowthCurve2DTrajectoryP1P2 instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.PartiallyGrowthCurve2DTrajectoryP1P2](partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback.probe.stageOfInterest)
			targetPartiallyGrowthCurve2DTrajectoryP1P2IDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetPartiallyGrowthCurve2DTrajectoryP1P2IDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all PartiallyGrowthCurve2DTrajectoryP1P2 instances and update their P2PointShapes slice
			for _partiallygrowthcurve2dtrajectoryp1p2 := range *models.GetGongstructInstancesSetFromPointerType[*models.PartiallyGrowthCurve2DTrajectoryP1P2](partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback.probe.stageOfInterest, _partiallygrowthcurve2dtrajectoryp1p2)
				
				// if PartiallyGrowthCurve2DTrajectoryP1P2 is selected
				if targetPartiallyGrowthCurve2DTrajectoryP1P2IDs[id] {
					// ensure partiallygrowthcurve2dtrajectoryp2pointshape_ is in _partiallygrowthcurve2dtrajectoryp1p2.P2PointShapes
					found := false
					for _, _b := range _partiallygrowthcurve2dtrajectoryp1p2.P2PointShapes {
						if _b == partiallygrowthcurve2dtrajectoryp2pointshape_ {
							found = true
							break
						}
					}
					if !found {
						_partiallygrowthcurve2dtrajectoryp1p2.P2PointShapes = append(_partiallygrowthcurve2dtrajectoryp1p2.P2PointShapes, partiallygrowthcurve2dtrajectoryp2pointshape_)
						partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback.probe.UpdateSliceOfPointersCallback(_partiallygrowthcurve2dtrajectoryp1p2, "P2PointShapes", &_partiallygrowthcurve2dtrajectoryp1p2.P2PointShapes)
					}
				} else {
					// ensure partiallygrowthcurve2dtrajectoryp2pointshape_ is NOT in _partiallygrowthcurve2dtrajectoryp1p2.P2PointShapes
					idx := slices.Index(_partiallygrowthcurve2dtrajectoryp1p2.P2PointShapes, partiallygrowthcurve2dtrajectoryp2pointshape_)
					if idx != -1 {
						_partiallygrowthcurve2dtrajectoryp1p2.P2PointShapes = slices.Delete(_partiallygrowthcurve2dtrajectoryp1p2.P2PointShapes, idx, idx+1)
						partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback.probe.UpdateSliceOfPointersCallback(_partiallygrowthcurve2dtrajectoryp1p2, "P2PointShapes", &_partiallygrowthcurve2dtrajectoryp1p2.P2PointShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		partiallygrowthcurve2dtrajectoryp2pointshape_.Unstage(partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback.probe.stageOfInterest)
	}

	partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.PartiallyGrowthCurve2DTrajectoryP2PointShape](
		partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback.CreationMode || partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PartiallyGrowthCurve2DTrajectoryP2PointShapeFormCallback(
			nil,
			partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback.probe,
			newFormGroup,
		)
		partiallygrowthcurve2dtrajectoryp2pointshape := new(models.PartiallyGrowthCurve2DTrajectoryP2PointShape)
		FillUpForm(partiallygrowthcurve2dtrajectoryp2pointshape, newFormGroup, partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback.probe)
		partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback.probe.formStage.Commit()
	}

	partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback.probe.ux_tree()
}
func __gong__New__PartiallyGrowthCurve2DTrajectoryShapeFormCallback(
	partiallygrowthcurve2dtrajectoryshape *models.PartiallyGrowthCurve2DTrajectoryShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (partiallygrowthcurve2dtrajectoryshapeFormCallback *PartiallyGrowthCurve2DTrajectoryShapeFormCallback) {
	partiallygrowthcurve2dtrajectoryshapeFormCallback = new(PartiallyGrowthCurve2DTrajectoryShapeFormCallback)
	partiallygrowthcurve2dtrajectoryshapeFormCallback.probe = probe
	partiallygrowthcurve2dtrajectoryshapeFormCallback.partiallygrowthcurve2dtrajectoryshape = partiallygrowthcurve2dtrajectoryshape
	partiallygrowthcurve2dtrajectoryshapeFormCallback.formGroup = formGroup

	partiallygrowthcurve2dtrajectoryshapeFormCallback.CreationMode = (partiallygrowthcurve2dtrajectoryshape == nil)

	return
}

type PartiallyGrowthCurve2DTrajectoryShapeFormCallback struct {
	partiallygrowthcurve2dtrajectoryshape *models.PartiallyGrowthCurve2DTrajectoryShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (partiallygrowthcurve2dtrajectoryshapeFormCallback *PartiallyGrowthCurve2DTrajectoryShapeFormCallback) OnSave() {
	partiallygrowthcurve2dtrajectoryshapeFormCallback.probe.stageOfInterest.Lock()
	defer partiallygrowthcurve2dtrajectoryshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PartiallyGrowthCurve2DTrajectoryShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	partiallygrowthcurve2dtrajectoryshapeFormCallback.probe.formStage.Checkout()

	if partiallygrowthcurve2dtrajectoryshapeFormCallback.partiallygrowthcurve2dtrajectoryshape == nil {
		partiallygrowthcurve2dtrajectoryshapeFormCallback.partiallygrowthcurve2dtrajectoryshape = new(models.PartiallyGrowthCurve2DTrajectoryShape).Stage(partiallygrowthcurve2dtrajectoryshapeFormCallback.probe.stageOfInterest)
	}
	partiallygrowthcurve2dtrajectoryshape_ := partiallygrowthcurve2dtrajectoryshapeFormCallback.partiallygrowthcurve2dtrajectoryshape
	_ = partiallygrowthcurve2dtrajectoryshape_

	for _, formDiv := range partiallygrowthcurve2dtrajectoryshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dtrajectoryshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dtrajectoryshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dtrajectoryshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dtrajectoryshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(partiallygrowthcurve2dtrajectoryshape_.EndY), formDiv)
		case "PartiallyGrowthCurve2DTrajectory:PartiallyGrowthCurve2DTrajectoryShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the PartiallyGrowthCurve2DTrajectory instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target PartiallyGrowthCurve2DTrajectory instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.PartiallyGrowthCurve2DTrajectory](partiallygrowthcurve2dtrajectoryshapeFormCallback.probe.stageOfInterest)
			targetPartiallyGrowthCurve2DTrajectoryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetPartiallyGrowthCurve2DTrajectoryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all PartiallyGrowthCurve2DTrajectory instances and update their PartiallyGrowthCurve2DTrajectoryShapes slice
			for _partiallygrowthcurve2dtrajectory := range *models.GetGongstructInstancesSetFromPointerType[*models.PartiallyGrowthCurve2DTrajectory](partiallygrowthcurve2dtrajectoryshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(partiallygrowthcurve2dtrajectoryshapeFormCallback.probe.stageOfInterest, _partiallygrowthcurve2dtrajectory)
				
				// if PartiallyGrowthCurve2DTrajectory is selected
				if targetPartiallyGrowthCurve2DTrajectoryIDs[id] {
					// ensure partiallygrowthcurve2dtrajectoryshape_ is in _partiallygrowthcurve2dtrajectory.PartiallyGrowthCurve2DTrajectoryShapes
					found := false
					for _, _b := range _partiallygrowthcurve2dtrajectory.PartiallyGrowthCurve2DTrajectoryShapes {
						if _b == partiallygrowthcurve2dtrajectoryshape_ {
							found = true
							break
						}
					}
					if !found {
						_partiallygrowthcurve2dtrajectory.PartiallyGrowthCurve2DTrajectoryShapes = append(_partiallygrowthcurve2dtrajectory.PartiallyGrowthCurve2DTrajectoryShapes, partiallygrowthcurve2dtrajectoryshape_)
						partiallygrowthcurve2dtrajectoryshapeFormCallback.probe.UpdateSliceOfPointersCallback(_partiallygrowthcurve2dtrajectory, "PartiallyGrowthCurve2DTrajectoryShapes", &_partiallygrowthcurve2dtrajectory.PartiallyGrowthCurve2DTrajectoryShapes)
					}
				} else {
					// ensure partiallygrowthcurve2dtrajectoryshape_ is NOT in _partiallygrowthcurve2dtrajectory.PartiallyGrowthCurve2DTrajectoryShapes
					idx := slices.Index(_partiallygrowthcurve2dtrajectory.PartiallyGrowthCurve2DTrajectoryShapes, partiallygrowthcurve2dtrajectoryshape_)
					if idx != -1 {
						_partiallygrowthcurve2dtrajectory.PartiallyGrowthCurve2DTrajectoryShapes = slices.Delete(_partiallygrowthcurve2dtrajectory.PartiallyGrowthCurve2DTrajectoryShapes, idx, idx+1)
						partiallygrowthcurve2dtrajectoryshapeFormCallback.probe.UpdateSliceOfPointersCallback(_partiallygrowthcurve2dtrajectory, "PartiallyGrowthCurve2DTrajectoryShapes", &_partiallygrowthcurve2dtrajectory.PartiallyGrowthCurve2DTrajectoryShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if partiallygrowthcurve2dtrajectoryshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		partiallygrowthcurve2dtrajectoryshape_.Unstage(partiallygrowthcurve2dtrajectoryshapeFormCallback.probe.stageOfInterest)
	}

	partiallygrowthcurve2dtrajectoryshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.PartiallyGrowthCurve2DTrajectoryShape](
		partiallygrowthcurve2dtrajectoryshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if partiallygrowthcurve2dtrajectoryshapeFormCallback.CreationMode || partiallygrowthcurve2dtrajectoryshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		partiallygrowthcurve2dtrajectoryshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(partiallygrowthcurve2dtrajectoryshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PartiallyGrowthCurve2DTrajectoryShapeFormCallback(
			nil,
			partiallygrowthcurve2dtrajectoryshapeFormCallback.probe,
			newFormGroup,
		)
		partiallygrowthcurve2dtrajectoryshape := new(models.PartiallyGrowthCurve2DTrajectoryShape)
		FillUpForm(partiallygrowthcurve2dtrajectoryshape, newFormGroup, partiallygrowthcurve2dtrajectoryshapeFormCallback.probe)
		partiallygrowthcurve2dtrajectoryshapeFormCallback.probe.formStage.Commit()
	}

	partiallygrowthcurve2dtrajectoryshapeFormCallback.probe.ux_tree()
}
func __gong__New__PartiallyRotatedTorusShapeFormCallback(
	partiallyrotatedtorusshape *models.PartiallyRotatedTorusShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (partiallyrotatedtorusshapeFormCallback *PartiallyRotatedTorusShapeFormCallback) {
	partiallyrotatedtorusshapeFormCallback = new(PartiallyRotatedTorusShapeFormCallback)
	partiallyrotatedtorusshapeFormCallback.probe = probe
	partiallyrotatedtorusshapeFormCallback.partiallyrotatedtorusshape = partiallyrotatedtorusshape
	partiallyrotatedtorusshapeFormCallback.formGroup = formGroup

	partiallyrotatedtorusshapeFormCallback.CreationMode = (partiallyrotatedtorusshape == nil)

	return
}

type PartiallyRotatedTorusShapeFormCallback struct {
	partiallyrotatedtorusshape *models.PartiallyRotatedTorusShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (partiallyrotatedtorusshapeFormCallback *PartiallyRotatedTorusShapeFormCallback) OnSave() {
	partiallyrotatedtorusshapeFormCallback.probe.stageOfInterest.Lock()
	defer partiallyrotatedtorusshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PartiallyRotatedTorusShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	partiallyrotatedtorusshapeFormCallback.probe.formStage.Checkout()

	if partiallyrotatedtorusshapeFormCallback.partiallyrotatedtorusshape == nil {
		partiallyrotatedtorusshapeFormCallback.partiallyrotatedtorusshape = new(models.PartiallyRotatedTorusShape).Stage(partiallyrotatedtorusshapeFormCallback.probe.stageOfInterest)
	}
	partiallyrotatedtorusshape_ := partiallyrotatedtorusshapeFormCallback.partiallyrotatedtorusshape
	_ = partiallyrotatedtorusshape_

	for _, formDiv := range partiallyrotatedtorusshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(partiallyrotatedtorusshape_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if partiallyrotatedtorusshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		partiallyrotatedtorusshape_.Unstage(partiallyrotatedtorusshapeFormCallback.probe.stageOfInterest)
	}

	partiallyrotatedtorusshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.PartiallyRotatedTorusShape](
		partiallyrotatedtorusshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if partiallyrotatedtorusshapeFormCallback.CreationMode || partiallyrotatedtorusshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		partiallyrotatedtorusshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(partiallyrotatedtorusshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PartiallyRotatedTorusShapeFormCallback(
			nil,
			partiallyrotatedtorusshapeFormCallback.probe,
			newFormGroup,
		)
		partiallyrotatedtorusshape := new(models.PartiallyRotatedTorusShape)
		FillUpForm(partiallyrotatedtorusshape, newFormGroup, partiallyrotatedtorusshapeFormCallback.probe)
		partiallyrotatedtorusshapeFormCallback.probe.formStage.Commit()
	}

	partiallyrotatedtorusshapeFormCallback.probe.ux_tree()
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
		case "RelativeVerticalThickness":
			FormDivBasicFieldToField(&(plant_.RelativeVerticalThickness), formDiv)
		case "RelativeRadialThickness":
			FormDivBasicFieldToField(&(plant_.RelativeRadialThickness), formDiv)
		case "RhombusSideLength":
			FormDivBasicFieldToField(&(plant_.RhombusSideLength), formDiv)
		case "RelativeCuttedStackFloorHeight":
			FormDivBasicFieldToField(&(plant_.RelativeCuttedStackFloorHeight), formDiv)
		case "RelativeRotatedTorusSeparation":
			FormDivBasicFieldToField(&(plant_.RelativeRotatedTorusSeparation), formDiv)
		case "RotationRatio":
			FormDivBasicFieldToField(&(plant_.RotationRatio), formDiv)
		case "ThreeDModulo":
			FormDivBasicFieldToField(&(plant_.ThreeDModulo), formDiv)
		case "RelativeTrajectoryOffsetX":
			FormDivBasicFieldToField(&(plant_.RelativeTrajectoryOffsetX), formDiv)
		case "RelativeTrajectoryOffsetY":
			FormDivBasicFieldToField(&(plant_.RelativeTrajectoryOffsetY), formDiv)
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
		case "RhombusStuff":
			FormDivSelectFieldToField(&(plant_.RhombusStuff), plantFormCallback.probe.stageOfInterest, formDiv)
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
		case "ShiftedBottomTopStartArcShapeGrid":
			FormDivSelectFieldToField(&(plant_.ShiftedBottomTopStartArcShapeGrid), plantFormCallback.probe.stageOfInterest, formDiv)
		case "MidArcVectorShapeGrid":
			FormDivSelectFieldToField(&(plant_.MidArcVectorShapeGrid), plantFormCallback.probe.stageOfInterest, formDiv)
		case "TopMidArcVectorShapeGrid":
			FormDivSelectFieldToField(&(plant_.TopMidArcVectorShapeGrid), plantFormCallback.probe.stageOfInterest, formDiv)
		case "StartHalfwayArcShapeGrid":
			FormDivSelectFieldToField(&(plant_.StartHalfwayArcShapeGrid), plantFormCallback.probe.stageOfInterest, formDiv)
		case "TopStartHalfwayArcShapeGrid":
			FormDivSelectFieldToField(&(plant_.TopStartHalfwayArcShapeGrid), plantFormCallback.probe.stageOfInterest, formDiv)
		case "EndHalfwayArcShapeGrid":
			FormDivSelectFieldToField(&(plant_.EndHalfwayArcShapeGrid), plantFormCallback.probe.stageOfInterest, formDiv)
		case "TopEndHalfwayArcShapeGrid":
			FormDivSelectFieldToField(&(plant_.TopEndHalfwayArcShapeGrid), plantFormCallback.probe.stageOfInterest, formDiv)
		case "StackOfRotatedGrowthCurve2D":
			FormDivSelectFieldToField(&(plant_.StackOfRotatedGrowthCurve2D), plantFormCallback.probe.stageOfInterest, formDiv)
		case "TopStackOfRotatedGrowthCurve2D":
			FormDivSelectFieldToField(&(plant_.TopStackOfRotatedGrowthCurve2D), plantFormCallback.probe.stageOfInterest, formDiv)
		case "GrowthCurve2D":
			FormDivSelectFieldToField(&(plant_.GrowthCurve2D), plantFormCallback.probe.stageOfInterest, formDiv)
		case "TopGrowthCurve2D":
			FormDivSelectFieldToField(&(plant_.TopGrowthCurve2D), plantFormCallback.probe.stageOfInterest, formDiv)
		case "StackOfGrowthCurve2D":
			FormDivSelectFieldToField(&(plant_.StackOfGrowthCurve2D), plantFormCallback.probe.stageOfInterest, formDiv)
		case "TopStackOfGrowthCurve2D":
			FormDivSelectFieldToField(&(plant_.TopStackOfGrowthCurve2D), plantFormCallback.probe.stageOfInterest, formDiv)
		case "StackOfGrowthCurve2DRibbon":
			FormDivSelectFieldToField(&(plant_.StackOfGrowthCurve2DRibbon), plantFormCallback.probe.stageOfInterest, formDiv)
		case "StackOfRotatedGrowthCurve2DRibbon":
			FormDivSelectFieldToField(&(plant_.StackOfRotatedGrowthCurve2DRibbon), plantFormCallback.probe.stageOfInterest, formDiv)
		case "GrowthCurve2DRibbon":
			FormDivSelectFieldToField(&(plant_.GrowthCurve2DRibbon), plantFormCallback.probe.stageOfInterest, formDiv)
		case "ShiftedRightGrowthCurve2DRibbon":
			FormDivSelectFieldToField(&(plant_.ShiftedRightGrowthCurve2DRibbon), plantFormCallback.probe.stageOfInterest, formDiv)
		case "PartiallyGrowthCurve2DRibbon":
			FormDivSelectFieldToField(&(plant_.PartiallyGrowthCurve2DRibbon), plantFormCallback.probe.stageOfInterest, formDiv)
		case "PartiallyGrowthCurve2DTrajectory":
			FormDivSelectFieldToField(&(plant_.PartiallyGrowthCurve2DTrajectory), plantFormCallback.probe.stageOfInterest, formDiv)
		case "PartiallyGrowthCurve2DTrajectoryP1P2":
			FormDivSelectFieldToField(&(plant_.PartiallyGrowthCurve2DTrajectoryP1P2), plantFormCallback.probe.stageOfInterest, formDiv)
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
		case "IsArcNodesExpanded":
			FormDivBasicFieldToField(&(plantdiagram_.IsArcNodesExpanded), formDiv)
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
		case "IsHiddenShiftedBottomTopStartArcShapeGrid":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenShiftedBottomTopStartArcShapeGrid), formDiv)
		case "IsHiddenMidArcVectorShapeGrid":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenMidArcVectorShapeGrid), formDiv)
		case "IsHiddenTopMidArcVectorShapeGrid":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenTopMidArcVectorShapeGrid), formDiv)
		case "IsHiddenStartHalfwayArcShapeGrid":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenStartHalfwayArcShapeGrid), formDiv)
		case "IsHiddenTopStartHalfwayArcShapeGrid":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenTopStartHalfwayArcShapeGrid), formDiv)
		case "IsHiddenEndHalfwayArcShapeGrid":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenEndHalfwayArcShapeGrid), formDiv)
		case "IsHiddenTopEndHalfwayArcShapeGrid":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenTopEndHalfwayArcShapeGrid), formDiv)
		case "IsHiddenEndArcShapeGrid":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenEndArcShapeGrid), formDiv)
		case "IsHiddenTopEndArcShapeGrid":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenTopEndArcShapeGrid), formDiv)
		case "IsHiddenBottomStartArcShapeGrid":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenBottomStartArcShapeGrid), formDiv)
		case "IsHiddenBottomEndArcShapeGrid":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenBottomEndArcShapeGrid), formDiv)
		case "IsHiddenStackOfGrowthCurve":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenStackOfGrowthCurve), formDiv)
		case "IsHiddenTopStackOfGrowthCurve":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenTopStackOfGrowthCurve), formDiv)
		case "IsHiddenBottomStackOfGrowthCurve":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenBottomStackOfGrowthCurve), formDiv)
		case "IsHiddenShiftedLeftStackOfGrowthCurve":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenShiftedLeftStackOfGrowthCurve), formDiv)
		case "IsHiddenShiftedLeftStackOfNormalVector":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenShiftedLeftStackOfNormalVector), formDiv)
		case "IsHiddenGrowthCurve2D":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenGrowthCurve2D), formDiv)
		case "IsHiddenTopGrowthCurve2D":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenTopGrowthCurve2D), formDiv)
		case "IsHiddenStackOfGrowthCurve2D":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenStackOfGrowthCurve2D), formDiv)
		case "IsHiddenTopStackOfGrowthCurve2D":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenTopStackOfGrowthCurve2D), formDiv)
		case "IsHiddenGrowthCurve2DRibbon":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenGrowthCurve2DRibbon), formDiv)
		case "IsHiddenShiftedRightGrowthCurve2DRibbon":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenShiftedRightGrowthCurve2DRibbon), formDiv)
		case "IsHiddenStackOfGrowthCurve2DRibbon":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenStackOfGrowthCurve2DRibbon), formDiv)
		case "IsHiddenStackOfRotatedGrowthCurve2DRibbon":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenStackOfRotatedGrowthCurve2DRibbon), formDiv)
		case "IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon), formDiv)
		case "IsHiddenPartiallyGrowthCurve2DRibbon":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenPartiallyGrowthCurve2DRibbon), formDiv)
		case "IsHiddenPartiallyGrowthCurve2DTrajectory":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenPartiallyGrowthCurve2DTrajectory), formDiv)
		case "IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2), formDiv)
		case "IsHiddenTorusStackShape":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenTorusStackShape), formDiv)
		case "IsHiddenVerticalTorusStackShape":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenVerticalTorusStackShape), formDiv)
		case "IsHiddenPartiallyRotatedTorusShape":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenPartiallyRotatedTorusShape), formDiv)
		case "IsHiddenStackOfPartiallyRotatedTorusShape":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenStackOfPartiallyRotatedTorusShape), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(plantdiagram_.IsChecked), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(plantdiagram_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(plantdiagram_.IsExpanded), formDiv)
		case "Rendered3DShape":
			FormDivSelectFieldToField(&(plantdiagram_.Rendered3DShape), plantdiagramFormCallback.probe.stageOfInterest, formDiv)
		case "GrowthCurve2DRibbon":
			FormDivSelectFieldToField(&(plantdiagram_.GrowthCurve2DRibbon), plantdiagramFormCallback.probe.stageOfInterest, formDiv)
		case "ShiftedRightGrowthCurve2DRibbon":
			FormDivSelectFieldToField(&(plantdiagram_.ShiftedRightGrowthCurve2DRibbon), plantdiagramFormCallback.probe.stageOfInterest, formDiv)
		case "TorusStackShape":
			FormDivSelectFieldToField(&(plantdiagram_.TorusStackShape), plantdiagramFormCallback.probe.stageOfInterest, formDiv)
		case "VerticalTorusStackShape":
			FormDivSelectFieldToField(&(plantdiagram_.VerticalTorusStackShape), plantdiagramFormCallback.probe.stageOfInterest, formDiv)
		case "PartiallyRotatedTorusShape":
			FormDivSelectFieldToField(&(plantdiagram_.PartiallyRotatedTorusShape), plantdiagramFormCallback.probe.stageOfInterest, formDiv)
		case "StackOfPartiallyRotatedTorusShape":
			FormDivSelectFieldToField(&(plantdiagram_.StackOfPartiallyRotatedTorusShape), plantdiagramFormCallback.probe.stageOfInterest, formDiv)
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
func __gong__New__RhombusStuffFormCallback(
	rhombusstuff *models.RhombusStuff,
	probe *Probe,
	formGroup *form.FormGroup,
) (rhombusstuffFormCallback *RhombusStuffFormCallback) {
	rhombusstuffFormCallback = new(RhombusStuffFormCallback)
	rhombusstuffFormCallback.probe = probe
	rhombusstuffFormCallback.rhombusstuff = rhombusstuff
	rhombusstuffFormCallback.formGroup = formGroup

	rhombusstuffFormCallback.CreationMode = (rhombusstuff == nil)

	return
}

type RhombusStuffFormCallback struct {
	rhombusstuff *models.RhombusStuff

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (rhombusstuffFormCallback *RhombusStuffFormCallback) OnSave() {
	rhombusstuffFormCallback.probe.stageOfInterest.Lock()
	defer rhombusstuffFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("RhombusStuffFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rhombusstuffFormCallback.probe.formStage.Checkout()

	if rhombusstuffFormCallback.rhombusstuff == nil {
		rhombusstuffFormCallback.rhombusstuff = new(models.RhombusStuff).Stage(rhombusstuffFormCallback.probe.stageOfInterest)
	}
	rhombusstuff_ := rhombusstuffFormCallback.rhombusstuff
	_ = rhombusstuff_

	for _, formDiv := range rhombusstuffFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rhombusstuff_.Name), formDiv)
		case "ReferenceRhombus":
			FormDivSelectFieldToField(&(rhombusstuff_.ReferenceRhombus), rhombusstuffFormCallback.probe.stageOfInterest, formDiv)
		case "PlantCircumferenceShape":
			FormDivSelectFieldToField(&(rhombusstuff_.PlantCircumferenceShape), rhombusstuffFormCallback.probe.stageOfInterest, formDiv)
		case "GridPathShape":
			FormDivSelectFieldToField(&(rhombusstuff_.GridPathShape), rhombusstuffFormCallback.probe.stageOfInterest, formDiv)
		case "InitialRhombusGridShape":
			FormDivSelectFieldToField(&(rhombusstuff_.InitialRhombusGridShape), rhombusstuffFormCallback.probe.stageOfInterest, formDiv)
		case "ExplanationTextShape":
			FormDivSelectFieldToField(&(rhombusstuff_.ExplanationTextShape), rhombusstuffFormCallback.probe.stageOfInterest, formDiv)
		case "RotatedReferenceRhombus":
			FormDivSelectFieldToField(&(rhombusstuff_.RotatedReferenceRhombus), rhombusstuffFormCallback.probe.stageOfInterest, formDiv)
		case "RotatedPlantCircumferenceShape":
			FormDivSelectFieldToField(&(rhombusstuff_.RotatedPlantCircumferenceShape), rhombusstuffFormCallback.probe.stageOfInterest, formDiv)
		case "RotatedGridPathShape":
			FormDivSelectFieldToField(&(rhombusstuff_.RotatedGridPathShape), rhombusstuffFormCallback.probe.stageOfInterest, formDiv)
		case "RotatedRhombusGridShape2":
			FormDivSelectFieldToField(&(rhombusstuff_.RotatedRhombusGridShape2), rhombusstuffFormCallback.probe.stageOfInterest, formDiv)
		case "GrowthCurveRhombusGridShape":
			FormDivSelectFieldToField(&(rhombusstuff_.GrowthCurveRhombusGridShape), rhombusstuffFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	// manage the suppress operation
	if rhombusstuffFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rhombusstuff_.Unstage(rhombusstuffFormCallback.probe.stageOfInterest)
	}

	rhombusstuffFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.RhombusStuff](
		rhombusstuffFormCallback.probe,
	)

	// display a new form by reset the form stage
	if rhombusstuffFormCallback.CreationMode || rhombusstuffFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rhombusstuffFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(rhombusstuffFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RhombusStuffFormCallback(
			nil,
			rhombusstuffFormCallback.probe,
			newFormGroup,
		)
		rhombusstuff := new(models.RhombusStuff)
		FillUpForm(rhombusstuff, newFormGroup, rhombusstuffFormCallback.probe)
		rhombusstuffFormCallback.probe.formStage.Commit()
	}

	rhombusstuffFormCallback.probe.ux_tree()
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
func __gong__New__ShiftedBottomTopStartArcShapeFormCallback(
	shiftedbottomtopstartarcshape *models.ShiftedBottomTopStartArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (shiftedbottomtopstartarcshapeFormCallback *ShiftedBottomTopStartArcShapeFormCallback) {
	shiftedbottomtopstartarcshapeFormCallback = new(ShiftedBottomTopStartArcShapeFormCallback)
	shiftedbottomtopstartarcshapeFormCallback.probe = probe
	shiftedbottomtopstartarcshapeFormCallback.shiftedbottomtopstartarcshape = shiftedbottomtopstartarcshape
	shiftedbottomtopstartarcshapeFormCallback.formGroup = formGroup

	shiftedbottomtopstartarcshapeFormCallback.CreationMode = (shiftedbottomtopstartarcshape == nil)

	return
}

type ShiftedBottomTopStartArcShapeFormCallback struct {
	shiftedbottomtopstartarcshape *models.ShiftedBottomTopStartArcShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (shiftedbottomtopstartarcshapeFormCallback *ShiftedBottomTopStartArcShapeFormCallback) OnSave() {
	shiftedbottomtopstartarcshapeFormCallback.probe.stageOfInterest.Lock()
	defer shiftedbottomtopstartarcshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ShiftedBottomTopStartArcShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	shiftedbottomtopstartarcshapeFormCallback.probe.formStage.Checkout()

	if shiftedbottomtopstartarcshapeFormCallback.shiftedbottomtopstartarcshape == nil {
		shiftedbottomtopstartarcshapeFormCallback.shiftedbottomtopstartarcshape = new(models.ShiftedBottomTopStartArcShape).Stage(shiftedbottomtopstartarcshapeFormCallback.probe.stageOfInterest)
	}
	shiftedbottomtopstartarcshape_ := shiftedbottomtopstartarcshapeFormCallback.shiftedbottomtopstartarcshape
	_ = shiftedbottomtopstartarcshape_

	for _, formDiv := range shiftedbottomtopstartarcshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(shiftedbottomtopstartarcshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(shiftedbottomtopstartarcshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(shiftedbottomtopstartarcshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(shiftedbottomtopstartarcshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(shiftedbottomtopstartarcshape_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(shiftedbottomtopstartarcshape_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(shiftedbottomtopstartarcshape_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(shiftedbottomtopstartarcshape_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(shiftedbottomtopstartarcshape_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(shiftedbottomtopstartarcshape_.RadiusY), formDiv)
		case "ShiftedBottomTopStartArcShapeGrid:ShiftedBottomTopStartArcShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the ShiftedBottomTopStartArcShapeGrid instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target ShiftedBottomTopStartArcShapeGrid instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.ShiftedBottomTopStartArcShapeGrid](shiftedbottomtopstartarcshapeFormCallback.probe.stageOfInterest)
			targetShiftedBottomTopStartArcShapeGridIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetShiftedBottomTopStartArcShapeGridIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all ShiftedBottomTopStartArcShapeGrid instances and update their ShiftedBottomTopStartArcShapes slice
			for _shiftedbottomtopstartarcshapegrid := range *models.GetGongstructInstancesSetFromPointerType[*models.ShiftedBottomTopStartArcShapeGrid](shiftedbottomtopstartarcshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(shiftedbottomtopstartarcshapeFormCallback.probe.stageOfInterest, _shiftedbottomtopstartarcshapegrid)
				
				// if ShiftedBottomTopStartArcShapeGrid is selected
				if targetShiftedBottomTopStartArcShapeGridIDs[id] {
					// ensure shiftedbottomtopstartarcshape_ is in _shiftedbottomtopstartarcshapegrid.ShiftedBottomTopStartArcShapes
					found := false
					for _, _b := range _shiftedbottomtopstartarcshapegrid.ShiftedBottomTopStartArcShapes {
						if _b == shiftedbottomtopstartarcshape_ {
							found = true
							break
						}
					}
					if !found {
						_shiftedbottomtopstartarcshapegrid.ShiftedBottomTopStartArcShapes = append(_shiftedbottomtopstartarcshapegrid.ShiftedBottomTopStartArcShapes, shiftedbottomtopstartarcshape_)
						shiftedbottomtopstartarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_shiftedbottomtopstartarcshapegrid, "ShiftedBottomTopStartArcShapes", &_shiftedbottomtopstartarcshapegrid.ShiftedBottomTopStartArcShapes)
					}
				} else {
					// ensure shiftedbottomtopstartarcshape_ is NOT in _shiftedbottomtopstartarcshapegrid.ShiftedBottomTopStartArcShapes
					idx := slices.Index(_shiftedbottomtopstartarcshapegrid.ShiftedBottomTopStartArcShapes, shiftedbottomtopstartarcshape_)
					if idx != -1 {
						_shiftedbottomtopstartarcshapegrid.ShiftedBottomTopStartArcShapes = slices.Delete(_shiftedbottomtopstartarcshapegrid.ShiftedBottomTopStartArcShapes, idx, idx+1)
						shiftedbottomtopstartarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_shiftedbottomtopstartarcshapegrid, "ShiftedBottomTopStartArcShapes", &_shiftedbottomtopstartarcshapegrid.ShiftedBottomTopStartArcShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if shiftedbottomtopstartarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		shiftedbottomtopstartarcshape_.Unstage(shiftedbottomtopstartarcshapeFormCallback.probe.stageOfInterest)
	}

	shiftedbottomtopstartarcshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ShiftedBottomTopStartArcShape](
		shiftedbottomtopstartarcshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if shiftedbottomtopstartarcshapeFormCallback.CreationMode || shiftedbottomtopstartarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		shiftedbottomtopstartarcshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(shiftedbottomtopstartarcshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ShiftedBottomTopStartArcShapeFormCallback(
			nil,
			shiftedbottomtopstartarcshapeFormCallback.probe,
			newFormGroup,
		)
		shiftedbottomtopstartarcshape := new(models.ShiftedBottomTopStartArcShape)
		FillUpForm(shiftedbottomtopstartarcshape, newFormGroup, shiftedbottomtopstartarcshapeFormCallback.probe)
		shiftedbottomtopstartarcshapeFormCallback.probe.formStage.Commit()
	}

	shiftedbottomtopstartarcshapeFormCallback.probe.ux_tree()
}
func __gong__New__ShiftedBottomTopStartArcShapeGridFormCallback(
	shiftedbottomtopstartarcshapegrid *models.ShiftedBottomTopStartArcShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (shiftedbottomtopstartarcshapegridFormCallback *ShiftedBottomTopStartArcShapeGridFormCallback) {
	shiftedbottomtopstartarcshapegridFormCallback = new(ShiftedBottomTopStartArcShapeGridFormCallback)
	shiftedbottomtopstartarcshapegridFormCallback.probe = probe
	shiftedbottomtopstartarcshapegridFormCallback.shiftedbottomtopstartarcshapegrid = shiftedbottomtopstartarcshapegrid
	shiftedbottomtopstartarcshapegridFormCallback.formGroup = formGroup

	shiftedbottomtopstartarcshapegridFormCallback.CreationMode = (shiftedbottomtopstartarcshapegrid == nil)

	return
}

type ShiftedBottomTopStartArcShapeGridFormCallback struct {
	shiftedbottomtopstartarcshapegrid *models.ShiftedBottomTopStartArcShapeGrid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (shiftedbottomtopstartarcshapegridFormCallback *ShiftedBottomTopStartArcShapeGridFormCallback) OnSave() {
	shiftedbottomtopstartarcshapegridFormCallback.probe.stageOfInterest.Lock()
	defer shiftedbottomtopstartarcshapegridFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ShiftedBottomTopStartArcShapeGridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	shiftedbottomtopstartarcshapegridFormCallback.probe.formStage.Checkout()

	if shiftedbottomtopstartarcshapegridFormCallback.shiftedbottomtopstartarcshapegrid == nil {
		shiftedbottomtopstartarcshapegridFormCallback.shiftedbottomtopstartarcshapegrid = new(models.ShiftedBottomTopStartArcShapeGrid).Stage(shiftedbottomtopstartarcshapegridFormCallback.probe.stageOfInterest)
	}
	shiftedbottomtopstartarcshapegrid_ := shiftedbottomtopstartarcshapegridFormCallback.shiftedbottomtopstartarcshapegrid
	_ = shiftedbottomtopstartarcshapegrid_

	for _, formDiv := range shiftedbottomtopstartarcshapegridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(shiftedbottomtopstartarcshapegrid_.Name), formDiv)
		case "ShiftedBottomTopStartArcShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ShiftedBottomTopStartArcShape](shiftedbottomtopstartarcshapegridFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ShiftedBottomTopStartArcShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ShiftedBottomTopStartArcShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					shiftedbottomtopstartarcshapegridFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ShiftedBottomTopStartArcShape](shiftedbottomtopstartarcshapegridFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			shiftedbottomtopstartarcshapegrid_.ShiftedBottomTopStartArcShapes = instanceSlice
			shiftedbottomtopstartarcshapegridFormCallback.probe.UpdateSliceOfPointersCallback(shiftedbottomtopstartarcshapegrid_, "ShiftedBottomTopStartArcShapes", &shiftedbottomtopstartarcshapegrid_.ShiftedBottomTopStartArcShapes)

		}
	}

	// manage the suppress operation
	if shiftedbottomtopstartarcshapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		shiftedbottomtopstartarcshapegrid_.Unstage(shiftedbottomtopstartarcshapegridFormCallback.probe.stageOfInterest)
	}

	shiftedbottomtopstartarcshapegridFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ShiftedBottomTopStartArcShapeGrid](
		shiftedbottomtopstartarcshapegridFormCallback.probe,
	)

	// display a new form by reset the form stage
	if shiftedbottomtopstartarcshapegridFormCallback.CreationMode || shiftedbottomtopstartarcshapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		shiftedbottomtopstartarcshapegridFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(shiftedbottomtopstartarcshapegridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ShiftedBottomTopStartArcShapeGridFormCallback(
			nil,
			shiftedbottomtopstartarcshapegridFormCallback.probe,
			newFormGroup,
		)
		shiftedbottomtopstartarcshapegrid := new(models.ShiftedBottomTopStartArcShapeGrid)
		FillUpForm(shiftedbottomtopstartarcshapegrid, newFormGroup, shiftedbottomtopstartarcshapegridFormCallback.probe)
		shiftedbottomtopstartarcshapegridFormCallback.probe.formStage.Commit()
	}

	shiftedbottomtopstartarcshapegridFormCallback.probe.ux_tree()
}
func __gong__New__ShiftedLeftStackGrowthCurveEndArcShapeFormCallback(
	shiftedleftstackgrowthcurveendarcshape *models.ShiftedLeftStackGrowthCurveEndArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (shiftedleftstackgrowthcurveendarcshapeFormCallback *ShiftedLeftStackGrowthCurveEndArcShapeFormCallback) {
	shiftedleftstackgrowthcurveendarcshapeFormCallback = new(ShiftedLeftStackGrowthCurveEndArcShapeFormCallback)
	shiftedleftstackgrowthcurveendarcshapeFormCallback.probe = probe
	shiftedleftstackgrowthcurveendarcshapeFormCallback.shiftedleftstackgrowthcurveendarcshape = shiftedleftstackgrowthcurveendarcshape
	shiftedleftstackgrowthcurveendarcshapeFormCallback.formGroup = formGroup

	shiftedleftstackgrowthcurveendarcshapeFormCallback.CreationMode = (shiftedleftstackgrowthcurveendarcshape == nil)

	return
}

type ShiftedLeftStackGrowthCurveEndArcShapeFormCallback struct {
	shiftedleftstackgrowthcurveendarcshape *models.ShiftedLeftStackGrowthCurveEndArcShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (shiftedleftstackgrowthcurveendarcshapeFormCallback *ShiftedLeftStackGrowthCurveEndArcShapeFormCallback) OnSave() {
	shiftedleftstackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest.Lock()
	defer shiftedleftstackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ShiftedLeftStackGrowthCurveEndArcShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	shiftedleftstackgrowthcurveendarcshapeFormCallback.probe.formStage.Checkout()

	if shiftedleftstackgrowthcurveendarcshapeFormCallback.shiftedleftstackgrowthcurveendarcshape == nil {
		shiftedleftstackgrowthcurveendarcshapeFormCallback.shiftedleftstackgrowthcurveendarcshape = new(models.ShiftedLeftStackGrowthCurveEndArcShape).Stage(shiftedleftstackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest)
	}
	shiftedleftstackgrowthcurveendarcshape_ := shiftedleftstackgrowthcurveendarcshapeFormCallback.shiftedleftstackgrowthcurveendarcshape
	_ = shiftedleftstackgrowthcurveendarcshape_

	for _, formDiv := range shiftedleftstackgrowthcurveendarcshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(shiftedleftstackgrowthcurveendarcshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(shiftedleftstackgrowthcurveendarcshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(shiftedleftstackgrowthcurveendarcshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(shiftedleftstackgrowthcurveendarcshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(shiftedleftstackgrowthcurveendarcshape_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(shiftedleftstackgrowthcurveendarcshape_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(shiftedleftstackgrowthcurveendarcshape_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(shiftedleftstackgrowthcurveendarcshape_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(shiftedleftstackgrowthcurveendarcshape_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(shiftedleftstackgrowthcurveendarcshape_.RadiusY), formDiv)
		case "ShiftedLeftStackOfGrowthCurve:ShiftedLeftStackGrowthCurveEndArcShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the ShiftedLeftStackOfGrowthCurve instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target ShiftedLeftStackOfGrowthCurve instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.ShiftedLeftStackOfGrowthCurve](shiftedleftstackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest)
			targetShiftedLeftStackOfGrowthCurveIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetShiftedLeftStackOfGrowthCurveIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all ShiftedLeftStackOfGrowthCurve instances and update their ShiftedLeftStackGrowthCurveEndArcShapes slice
			for _shiftedleftstackofgrowthcurve := range *models.GetGongstructInstancesSetFromPointerType[*models.ShiftedLeftStackOfGrowthCurve](shiftedleftstackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(shiftedleftstackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest, _shiftedleftstackofgrowthcurve)
				
				// if ShiftedLeftStackOfGrowthCurve is selected
				if targetShiftedLeftStackOfGrowthCurveIDs[id] {
					// ensure shiftedleftstackgrowthcurveendarcshape_ is in _shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveEndArcShapes
					found := false
					for _, _b := range _shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveEndArcShapes {
						if _b == shiftedleftstackgrowthcurveendarcshape_ {
							found = true
							break
						}
					}
					if !found {
						_shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveEndArcShapes = append(_shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveEndArcShapes, shiftedleftstackgrowthcurveendarcshape_)
						shiftedleftstackgrowthcurveendarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_shiftedleftstackofgrowthcurve, "ShiftedLeftStackGrowthCurveEndArcShapes", &_shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveEndArcShapes)
					}
				} else {
					// ensure shiftedleftstackgrowthcurveendarcshape_ is NOT in _shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveEndArcShapes
					idx := slices.Index(_shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveEndArcShapes, shiftedleftstackgrowthcurveendarcshape_)
					if idx != -1 {
						_shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveEndArcShapes = slices.Delete(_shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveEndArcShapes, idx, idx+1)
						shiftedleftstackgrowthcurveendarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_shiftedleftstackofgrowthcurve, "ShiftedLeftStackGrowthCurveEndArcShapes", &_shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveEndArcShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if shiftedleftstackgrowthcurveendarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		shiftedleftstackgrowthcurveendarcshape_.Unstage(shiftedleftstackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest)
	}

	shiftedleftstackgrowthcurveendarcshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ShiftedLeftStackGrowthCurveEndArcShape](
		shiftedleftstackgrowthcurveendarcshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if shiftedleftstackgrowthcurveendarcshapeFormCallback.CreationMode || shiftedleftstackgrowthcurveendarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		shiftedleftstackgrowthcurveendarcshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(shiftedleftstackgrowthcurveendarcshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ShiftedLeftStackGrowthCurveEndArcShapeFormCallback(
			nil,
			shiftedleftstackgrowthcurveendarcshapeFormCallback.probe,
			newFormGroup,
		)
		shiftedleftstackgrowthcurveendarcshape := new(models.ShiftedLeftStackGrowthCurveEndArcShape)
		FillUpForm(shiftedleftstackgrowthcurveendarcshape, newFormGroup, shiftedleftstackgrowthcurveendarcshapeFormCallback.probe)
		shiftedleftstackgrowthcurveendarcshapeFormCallback.probe.formStage.Commit()
	}

	shiftedleftstackgrowthcurveendarcshapeFormCallback.probe.ux_tree()
}
func __gong__New__ShiftedLeftStackGrowthCurveStartArcShapeFormCallback(
	shiftedleftstackgrowthcurvestartarcshape *models.ShiftedLeftStackGrowthCurveStartArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (shiftedleftstackgrowthcurvestartarcshapeFormCallback *ShiftedLeftStackGrowthCurveStartArcShapeFormCallback) {
	shiftedleftstackgrowthcurvestartarcshapeFormCallback = new(ShiftedLeftStackGrowthCurveStartArcShapeFormCallback)
	shiftedleftstackgrowthcurvestartarcshapeFormCallback.probe = probe
	shiftedleftstackgrowthcurvestartarcshapeFormCallback.shiftedleftstackgrowthcurvestartarcshape = shiftedleftstackgrowthcurvestartarcshape
	shiftedleftstackgrowthcurvestartarcshapeFormCallback.formGroup = formGroup

	shiftedleftstackgrowthcurvestartarcshapeFormCallback.CreationMode = (shiftedleftstackgrowthcurvestartarcshape == nil)

	return
}

type ShiftedLeftStackGrowthCurveStartArcShapeFormCallback struct {
	shiftedleftstackgrowthcurvestartarcshape *models.ShiftedLeftStackGrowthCurveStartArcShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (shiftedleftstackgrowthcurvestartarcshapeFormCallback *ShiftedLeftStackGrowthCurveStartArcShapeFormCallback) OnSave() {
	shiftedleftstackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest.Lock()
	defer shiftedleftstackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ShiftedLeftStackGrowthCurveStartArcShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	shiftedleftstackgrowthcurvestartarcshapeFormCallback.probe.formStage.Checkout()

	if shiftedleftstackgrowthcurvestartarcshapeFormCallback.shiftedleftstackgrowthcurvestartarcshape == nil {
		shiftedleftstackgrowthcurvestartarcshapeFormCallback.shiftedleftstackgrowthcurvestartarcshape = new(models.ShiftedLeftStackGrowthCurveStartArcShape).Stage(shiftedleftstackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest)
	}
	shiftedleftstackgrowthcurvestartarcshape_ := shiftedleftstackgrowthcurvestartarcshapeFormCallback.shiftedleftstackgrowthcurvestartarcshape
	_ = shiftedleftstackgrowthcurvestartarcshape_

	for _, formDiv := range shiftedleftstackgrowthcurvestartarcshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(shiftedleftstackgrowthcurvestartarcshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(shiftedleftstackgrowthcurvestartarcshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(shiftedleftstackgrowthcurvestartarcshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(shiftedleftstackgrowthcurvestartarcshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(shiftedleftstackgrowthcurvestartarcshape_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(shiftedleftstackgrowthcurvestartarcshape_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(shiftedleftstackgrowthcurvestartarcshape_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(shiftedleftstackgrowthcurvestartarcshape_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(shiftedleftstackgrowthcurvestartarcshape_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(shiftedleftstackgrowthcurvestartarcshape_.RadiusY), formDiv)
		case "ShiftedLeftStackOfGrowthCurve:ShiftedLeftStackGrowthCurveStartArcShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the ShiftedLeftStackOfGrowthCurve instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target ShiftedLeftStackOfGrowthCurve instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.ShiftedLeftStackOfGrowthCurve](shiftedleftstackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest)
			targetShiftedLeftStackOfGrowthCurveIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetShiftedLeftStackOfGrowthCurveIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all ShiftedLeftStackOfGrowthCurve instances and update their ShiftedLeftStackGrowthCurveStartArcShapes slice
			for _shiftedleftstackofgrowthcurve := range *models.GetGongstructInstancesSetFromPointerType[*models.ShiftedLeftStackOfGrowthCurve](shiftedleftstackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(shiftedleftstackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest, _shiftedleftstackofgrowthcurve)
				
				// if ShiftedLeftStackOfGrowthCurve is selected
				if targetShiftedLeftStackOfGrowthCurveIDs[id] {
					// ensure shiftedleftstackgrowthcurvestartarcshape_ is in _shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveStartArcShapes
					found := false
					for _, _b := range _shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveStartArcShapes {
						if _b == shiftedleftstackgrowthcurvestartarcshape_ {
							found = true
							break
						}
					}
					if !found {
						_shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveStartArcShapes = append(_shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveStartArcShapes, shiftedleftstackgrowthcurvestartarcshape_)
						shiftedleftstackgrowthcurvestartarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_shiftedleftstackofgrowthcurve, "ShiftedLeftStackGrowthCurveStartArcShapes", &_shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveStartArcShapes)
					}
				} else {
					// ensure shiftedleftstackgrowthcurvestartarcshape_ is NOT in _shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveStartArcShapes
					idx := slices.Index(_shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveStartArcShapes, shiftedleftstackgrowthcurvestartarcshape_)
					if idx != -1 {
						_shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveStartArcShapes = slices.Delete(_shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveStartArcShapes, idx, idx+1)
						shiftedleftstackgrowthcurvestartarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_shiftedleftstackofgrowthcurve, "ShiftedLeftStackGrowthCurveStartArcShapes", &_shiftedleftstackofgrowthcurve.ShiftedLeftStackGrowthCurveStartArcShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if shiftedleftstackgrowthcurvestartarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		shiftedleftstackgrowthcurvestartarcshape_.Unstage(shiftedleftstackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest)
	}

	shiftedleftstackgrowthcurvestartarcshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ShiftedLeftStackGrowthCurveStartArcShape](
		shiftedleftstackgrowthcurvestartarcshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if shiftedleftstackgrowthcurvestartarcshapeFormCallback.CreationMode || shiftedleftstackgrowthcurvestartarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		shiftedleftstackgrowthcurvestartarcshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(shiftedleftstackgrowthcurvestartarcshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ShiftedLeftStackGrowthCurveStartArcShapeFormCallback(
			nil,
			shiftedleftstackgrowthcurvestartarcshapeFormCallback.probe,
			newFormGroup,
		)
		shiftedleftstackgrowthcurvestartarcshape := new(models.ShiftedLeftStackGrowthCurveStartArcShape)
		FillUpForm(shiftedleftstackgrowthcurvestartarcshape, newFormGroup, shiftedleftstackgrowthcurvestartarcshapeFormCallback.probe)
		shiftedleftstackgrowthcurvestartarcshapeFormCallback.probe.formStage.Commit()
	}

	shiftedleftstackgrowthcurvestartarcshapeFormCallback.probe.ux_tree()
}
func __gong__New__ShiftedLeftStackNormalVectorFormCallback(
	shiftedleftstacknormalvector *models.ShiftedLeftStackNormalVector,
	probe *Probe,
	formGroup *form.FormGroup,
) (shiftedleftstacknormalvectorFormCallback *ShiftedLeftStackNormalVectorFormCallback) {
	shiftedleftstacknormalvectorFormCallback = new(ShiftedLeftStackNormalVectorFormCallback)
	shiftedleftstacknormalvectorFormCallback.probe = probe
	shiftedleftstacknormalvectorFormCallback.shiftedleftstacknormalvector = shiftedleftstacknormalvector
	shiftedleftstacknormalvectorFormCallback.formGroup = formGroup

	shiftedleftstacknormalvectorFormCallback.CreationMode = (shiftedleftstacknormalvector == nil)

	return
}

type ShiftedLeftStackNormalVectorFormCallback struct {
	shiftedleftstacknormalvector *models.ShiftedLeftStackNormalVector

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (shiftedleftstacknormalvectorFormCallback *ShiftedLeftStackNormalVectorFormCallback) OnSave() {
	shiftedleftstacknormalvectorFormCallback.probe.stageOfInterest.Lock()
	defer shiftedleftstacknormalvectorFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ShiftedLeftStackNormalVectorFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	shiftedleftstacknormalvectorFormCallback.probe.formStage.Checkout()

	if shiftedleftstacknormalvectorFormCallback.shiftedleftstacknormalvector == nil {
		shiftedleftstacknormalvectorFormCallback.shiftedleftstacknormalvector = new(models.ShiftedLeftStackNormalVector).Stage(shiftedleftstacknormalvectorFormCallback.probe.stageOfInterest)
	}
	shiftedleftstacknormalvector_ := shiftedleftstacknormalvectorFormCallback.shiftedleftstacknormalvector
	_ = shiftedleftstacknormalvector_

	for _, formDiv := range shiftedleftstacknormalvectorFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(shiftedleftstacknormalvector_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(shiftedleftstacknormalvector_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(shiftedleftstacknormalvector_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(shiftedleftstacknormalvector_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(shiftedleftstacknormalvector_.EndY), formDiv)
		case "ShiftedLeftStackOfNormalVector:ShiftedLeftStackNormalVectors":
			// 1. Decode the AssociationStorage which contains the rowIDs of the ShiftedLeftStackOfNormalVector instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target ShiftedLeftStackOfNormalVector instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.ShiftedLeftStackOfNormalVector](shiftedleftstacknormalvectorFormCallback.probe.stageOfInterest)
			targetShiftedLeftStackOfNormalVectorIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetShiftedLeftStackOfNormalVectorIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all ShiftedLeftStackOfNormalVector instances and update their ShiftedLeftStackNormalVectors slice
			for _shiftedleftstackofnormalvector := range *models.GetGongstructInstancesSetFromPointerType[*models.ShiftedLeftStackOfNormalVector](shiftedleftstacknormalvectorFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(shiftedleftstacknormalvectorFormCallback.probe.stageOfInterest, _shiftedleftstackofnormalvector)
				
				// if ShiftedLeftStackOfNormalVector is selected
				if targetShiftedLeftStackOfNormalVectorIDs[id] {
					// ensure shiftedleftstacknormalvector_ is in _shiftedleftstackofnormalvector.ShiftedLeftStackNormalVectors
					found := false
					for _, _b := range _shiftedleftstackofnormalvector.ShiftedLeftStackNormalVectors {
						if _b == shiftedleftstacknormalvector_ {
							found = true
							break
						}
					}
					if !found {
						_shiftedleftstackofnormalvector.ShiftedLeftStackNormalVectors = append(_shiftedleftstackofnormalvector.ShiftedLeftStackNormalVectors, shiftedleftstacknormalvector_)
						shiftedleftstacknormalvectorFormCallback.probe.UpdateSliceOfPointersCallback(_shiftedleftstackofnormalvector, "ShiftedLeftStackNormalVectors", &_shiftedleftstackofnormalvector.ShiftedLeftStackNormalVectors)
					}
				} else {
					// ensure shiftedleftstacknormalvector_ is NOT in _shiftedleftstackofnormalvector.ShiftedLeftStackNormalVectors
					idx := slices.Index(_shiftedleftstackofnormalvector.ShiftedLeftStackNormalVectors, shiftedleftstacknormalvector_)
					if idx != -1 {
						_shiftedleftstackofnormalvector.ShiftedLeftStackNormalVectors = slices.Delete(_shiftedleftstackofnormalvector.ShiftedLeftStackNormalVectors, idx, idx+1)
						shiftedleftstacknormalvectorFormCallback.probe.UpdateSliceOfPointersCallback(_shiftedleftstackofnormalvector, "ShiftedLeftStackNormalVectors", &_shiftedleftstackofnormalvector.ShiftedLeftStackNormalVectors)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if shiftedleftstacknormalvectorFormCallback.formGroup.HasSuppressButtonBeenPressed {
		shiftedleftstacknormalvector_.Unstage(shiftedleftstacknormalvectorFormCallback.probe.stageOfInterest)
	}

	shiftedleftstacknormalvectorFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ShiftedLeftStackNormalVector](
		shiftedleftstacknormalvectorFormCallback.probe,
	)

	// display a new form by reset the form stage
	if shiftedleftstacknormalvectorFormCallback.CreationMode || shiftedleftstacknormalvectorFormCallback.formGroup.HasSuppressButtonBeenPressed {
		shiftedleftstacknormalvectorFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(shiftedleftstacknormalvectorFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ShiftedLeftStackNormalVectorFormCallback(
			nil,
			shiftedleftstacknormalvectorFormCallback.probe,
			newFormGroup,
		)
		shiftedleftstacknormalvector := new(models.ShiftedLeftStackNormalVector)
		FillUpForm(shiftedleftstacknormalvector, newFormGroup, shiftedleftstacknormalvectorFormCallback.probe)
		shiftedleftstacknormalvectorFormCallback.probe.formStage.Commit()
	}

	shiftedleftstacknormalvectorFormCallback.probe.ux_tree()
}
func __gong__New__ShiftedLeftStackOfGrowthCurveFormCallback(
	shiftedleftstackofgrowthcurve *models.ShiftedLeftStackOfGrowthCurve,
	probe *Probe,
	formGroup *form.FormGroup,
) (shiftedleftstackofgrowthcurveFormCallback *ShiftedLeftStackOfGrowthCurveFormCallback) {
	shiftedleftstackofgrowthcurveFormCallback = new(ShiftedLeftStackOfGrowthCurveFormCallback)
	shiftedleftstackofgrowthcurveFormCallback.probe = probe
	shiftedleftstackofgrowthcurveFormCallback.shiftedleftstackofgrowthcurve = shiftedleftstackofgrowthcurve
	shiftedleftstackofgrowthcurveFormCallback.formGroup = formGroup

	shiftedleftstackofgrowthcurveFormCallback.CreationMode = (shiftedleftstackofgrowthcurve == nil)

	return
}

type ShiftedLeftStackOfGrowthCurveFormCallback struct {
	shiftedleftstackofgrowthcurve *models.ShiftedLeftStackOfGrowthCurve

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (shiftedleftstackofgrowthcurveFormCallback *ShiftedLeftStackOfGrowthCurveFormCallback) OnSave() {
	shiftedleftstackofgrowthcurveFormCallback.probe.stageOfInterest.Lock()
	defer shiftedleftstackofgrowthcurveFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ShiftedLeftStackOfGrowthCurveFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	shiftedleftstackofgrowthcurveFormCallback.probe.formStage.Checkout()

	if shiftedleftstackofgrowthcurveFormCallback.shiftedleftstackofgrowthcurve == nil {
		shiftedleftstackofgrowthcurveFormCallback.shiftedleftstackofgrowthcurve = new(models.ShiftedLeftStackOfGrowthCurve).Stage(shiftedleftstackofgrowthcurveFormCallback.probe.stageOfInterest)
	}
	shiftedleftstackofgrowthcurve_ := shiftedleftstackofgrowthcurveFormCallback.shiftedleftstackofgrowthcurve
	_ = shiftedleftstackofgrowthcurve_

	for _, formDiv := range shiftedleftstackofgrowthcurveFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(shiftedleftstackofgrowthcurve_.Name), formDiv)
		case "ShiftedLeftStackGrowthCurveStartArcShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ShiftedLeftStackGrowthCurveStartArcShape](shiftedleftstackofgrowthcurveFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ShiftedLeftStackGrowthCurveStartArcShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ShiftedLeftStackGrowthCurveStartArcShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					shiftedleftstackofgrowthcurveFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ShiftedLeftStackGrowthCurveStartArcShape](shiftedleftstackofgrowthcurveFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			shiftedleftstackofgrowthcurve_.ShiftedLeftStackGrowthCurveStartArcShapes = instanceSlice
			shiftedleftstackofgrowthcurveFormCallback.probe.UpdateSliceOfPointersCallback(shiftedleftstackofgrowthcurve_, "ShiftedLeftStackGrowthCurveStartArcShapes", &shiftedleftstackofgrowthcurve_.ShiftedLeftStackGrowthCurveStartArcShapes)

		case "ShiftedLeftStackGrowthCurveEndArcShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ShiftedLeftStackGrowthCurveEndArcShape](shiftedleftstackofgrowthcurveFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ShiftedLeftStackGrowthCurveEndArcShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ShiftedLeftStackGrowthCurveEndArcShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					shiftedleftstackofgrowthcurveFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ShiftedLeftStackGrowthCurveEndArcShape](shiftedleftstackofgrowthcurveFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			shiftedleftstackofgrowthcurve_.ShiftedLeftStackGrowthCurveEndArcShapes = instanceSlice
			shiftedleftstackofgrowthcurveFormCallback.probe.UpdateSliceOfPointersCallback(shiftedleftstackofgrowthcurve_, "ShiftedLeftStackGrowthCurveEndArcShapes", &shiftedleftstackofgrowthcurve_.ShiftedLeftStackGrowthCurveEndArcShapes)

		}
	}

	// manage the suppress operation
	if shiftedleftstackofgrowthcurveFormCallback.formGroup.HasSuppressButtonBeenPressed {
		shiftedleftstackofgrowthcurve_.Unstage(shiftedleftstackofgrowthcurveFormCallback.probe.stageOfInterest)
	}

	shiftedleftstackofgrowthcurveFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ShiftedLeftStackOfGrowthCurve](
		shiftedleftstackofgrowthcurveFormCallback.probe,
	)

	// display a new form by reset the form stage
	if shiftedleftstackofgrowthcurveFormCallback.CreationMode || shiftedleftstackofgrowthcurveFormCallback.formGroup.HasSuppressButtonBeenPressed {
		shiftedleftstackofgrowthcurveFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(shiftedleftstackofgrowthcurveFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ShiftedLeftStackOfGrowthCurveFormCallback(
			nil,
			shiftedleftstackofgrowthcurveFormCallback.probe,
			newFormGroup,
		)
		shiftedleftstackofgrowthcurve := new(models.ShiftedLeftStackOfGrowthCurve)
		FillUpForm(shiftedleftstackofgrowthcurve, newFormGroup, shiftedleftstackofgrowthcurveFormCallback.probe)
		shiftedleftstackofgrowthcurveFormCallback.probe.formStage.Commit()
	}

	shiftedleftstackofgrowthcurveFormCallback.probe.ux_tree()
}
func __gong__New__ShiftedLeftStackOfNormalVectorFormCallback(
	shiftedleftstackofnormalvector *models.ShiftedLeftStackOfNormalVector,
	probe *Probe,
	formGroup *form.FormGroup,
) (shiftedleftstackofnormalvectorFormCallback *ShiftedLeftStackOfNormalVectorFormCallback) {
	shiftedleftstackofnormalvectorFormCallback = new(ShiftedLeftStackOfNormalVectorFormCallback)
	shiftedleftstackofnormalvectorFormCallback.probe = probe
	shiftedleftstackofnormalvectorFormCallback.shiftedleftstackofnormalvector = shiftedleftstackofnormalvector
	shiftedleftstackofnormalvectorFormCallback.formGroup = formGroup

	shiftedleftstackofnormalvectorFormCallback.CreationMode = (shiftedleftstackofnormalvector == nil)

	return
}

type ShiftedLeftStackOfNormalVectorFormCallback struct {
	shiftedleftstackofnormalvector *models.ShiftedLeftStackOfNormalVector

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (shiftedleftstackofnormalvectorFormCallback *ShiftedLeftStackOfNormalVectorFormCallback) OnSave() {
	shiftedleftstackofnormalvectorFormCallback.probe.stageOfInterest.Lock()
	defer shiftedleftstackofnormalvectorFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ShiftedLeftStackOfNormalVectorFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	shiftedleftstackofnormalvectorFormCallback.probe.formStage.Checkout()

	if shiftedleftstackofnormalvectorFormCallback.shiftedleftstackofnormalvector == nil {
		shiftedleftstackofnormalvectorFormCallback.shiftedleftstackofnormalvector = new(models.ShiftedLeftStackOfNormalVector).Stage(shiftedleftstackofnormalvectorFormCallback.probe.stageOfInterest)
	}
	shiftedleftstackofnormalvector_ := shiftedleftstackofnormalvectorFormCallback.shiftedleftstackofnormalvector
	_ = shiftedleftstackofnormalvector_

	for _, formDiv := range shiftedleftstackofnormalvectorFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(shiftedleftstackofnormalvector_.Name), formDiv)
		case "ShiftedLeftStackNormalVectors":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ShiftedLeftStackNormalVector](shiftedleftstackofnormalvectorFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ShiftedLeftStackNormalVector, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ShiftedLeftStackNormalVector)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					shiftedleftstackofnormalvectorFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ShiftedLeftStackNormalVector](shiftedleftstackofnormalvectorFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			shiftedleftstackofnormalvector_.ShiftedLeftStackNormalVectors = instanceSlice
			shiftedleftstackofnormalvectorFormCallback.probe.UpdateSliceOfPointersCallback(shiftedleftstackofnormalvector_, "ShiftedLeftStackNormalVectors", &shiftedleftstackofnormalvector_.ShiftedLeftStackNormalVectors)

		}
	}

	// manage the suppress operation
	if shiftedleftstackofnormalvectorFormCallback.formGroup.HasSuppressButtonBeenPressed {
		shiftedleftstackofnormalvector_.Unstage(shiftedleftstackofnormalvectorFormCallback.probe.stageOfInterest)
	}

	shiftedleftstackofnormalvectorFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ShiftedLeftStackOfNormalVector](
		shiftedleftstackofnormalvectorFormCallback.probe,
	)

	// display a new form by reset the form stage
	if shiftedleftstackofnormalvectorFormCallback.CreationMode || shiftedleftstackofnormalvectorFormCallback.formGroup.HasSuppressButtonBeenPressed {
		shiftedleftstackofnormalvectorFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(shiftedleftstackofnormalvectorFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ShiftedLeftStackOfNormalVectorFormCallback(
			nil,
			shiftedleftstackofnormalvectorFormCallback.probe,
			newFormGroup,
		)
		shiftedleftstackofnormalvector := new(models.ShiftedLeftStackOfNormalVector)
		FillUpForm(shiftedleftstackofnormalvector, newFormGroup, shiftedleftstackofnormalvectorFormCallback.probe)
		shiftedleftstackofnormalvectorFormCallback.probe.formStage.Commit()
	}

	shiftedleftstackofnormalvectorFormCallback.probe.ux_tree()
}
func __gong__New__ShiftedRightGrowthCurve2DRibbonFormCallback(
	shiftedrightgrowthcurve2dribbon *models.ShiftedRightGrowthCurve2DRibbon,
	probe *Probe,
	formGroup *form.FormGroup,
) (shiftedrightgrowthcurve2dribbonFormCallback *ShiftedRightGrowthCurve2DRibbonFormCallback) {
	shiftedrightgrowthcurve2dribbonFormCallback = new(ShiftedRightGrowthCurve2DRibbonFormCallback)
	shiftedrightgrowthcurve2dribbonFormCallback.probe = probe
	shiftedrightgrowthcurve2dribbonFormCallback.shiftedrightgrowthcurve2dribbon = shiftedrightgrowthcurve2dribbon
	shiftedrightgrowthcurve2dribbonFormCallback.formGroup = formGroup

	shiftedrightgrowthcurve2dribbonFormCallback.CreationMode = (shiftedrightgrowthcurve2dribbon == nil)

	return
}

type ShiftedRightGrowthCurve2DRibbonFormCallback struct {
	shiftedrightgrowthcurve2dribbon *models.ShiftedRightGrowthCurve2DRibbon

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (shiftedrightgrowthcurve2dribbonFormCallback *ShiftedRightGrowthCurve2DRibbonFormCallback) OnSave() {
	shiftedrightgrowthcurve2dribbonFormCallback.probe.stageOfInterest.Lock()
	defer shiftedrightgrowthcurve2dribbonFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ShiftedRightGrowthCurve2DRibbonFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	shiftedrightgrowthcurve2dribbonFormCallback.probe.formStage.Checkout()

	if shiftedrightgrowthcurve2dribbonFormCallback.shiftedrightgrowthcurve2dribbon == nil {
		shiftedrightgrowthcurve2dribbonFormCallback.shiftedrightgrowthcurve2dribbon = new(models.ShiftedRightGrowthCurve2DRibbon).Stage(shiftedrightgrowthcurve2dribbonFormCallback.probe.stageOfInterest)
	}
	shiftedrightgrowthcurve2dribbon_ := shiftedrightgrowthcurve2dribbonFormCallback.shiftedrightgrowthcurve2dribbon
	_ = shiftedrightgrowthcurve2dribbon_

	for _, formDiv := range shiftedrightgrowthcurve2dribbonFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbon_.Name), formDiv)
		case "ShiftedRightGrowthCurve2DRibbonStartShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ShiftedRightGrowthCurve2DRibbonStartShape](shiftedrightgrowthcurve2dribbonFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ShiftedRightGrowthCurve2DRibbonStartShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ShiftedRightGrowthCurve2DRibbonStartShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					shiftedrightgrowthcurve2dribbonFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ShiftedRightGrowthCurve2DRibbonStartShape](shiftedrightgrowthcurve2dribbonFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			shiftedrightgrowthcurve2dribbon_.ShiftedRightGrowthCurve2DRibbonStartShapes = instanceSlice
			shiftedrightgrowthcurve2dribbonFormCallback.probe.UpdateSliceOfPointersCallback(shiftedrightgrowthcurve2dribbon_, "ShiftedRightGrowthCurve2DRibbonStartShapes", &shiftedrightgrowthcurve2dribbon_.ShiftedRightGrowthCurve2DRibbonStartShapes)

		case "ShiftedRightGrowthCurve2DRibbonEndShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ShiftedRightGrowthCurve2DRibbonEndShape](shiftedrightgrowthcurve2dribbonFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ShiftedRightGrowthCurve2DRibbonEndShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ShiftedRightGrowthCurve2DRibbonEndShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					shiftedrightgrowthcurve2dribbonFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ShiftedRightGrowthCurve2DRibbonEndShape](shiftedrightgrowthcurve2dribbonFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			shiftedrightgrowthcurve2dribbon_.ShiftedRightGrowthCurve2DRibbonEndShapes = instanceSlice
			shiftedrightgrowthcurve2dribbonFormCallback.probe.UpdateSliceOfPointersCallback(shiftedrightgrowthcurve2dribbon_, "ShiftedRightGrowthCurve2DRibbonEndShapes", &shiftedrightgrowthcurve2dribbon_.ShiftedRightGrowthCurve2DRibbonEndShapes)

		}
	}

	// manage the suppress operation
	if shiftedrightgrowthcurve2dribbonFormCallback.formGroup.HasSuppressButtonBeenPressed {
		shiftedrightgrowthcurve2dribbon_.Unstage(shiftedrightgrowthcurve2dribbonFormCallback.probe.stageOfInterest)
	}

	shiftedrightgrowthcurve2dribbonFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ShiftedRightGrowthCurve2DRibbon](
		shiftedrightgrowthcurve2dribbonFormCallback.probe,
	)

	// display a new form by reset the form stage
	if shiftedrightgrowthcurve2dribbonFormCallback.CreationMode || shiftedrightgrowthcurve2dribbonFormCallback.formGroup.HasSuppressButtonBeenPressed {
		shiftedrightgrowthcurve2dribbonFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(shiftedrightgrowthcurve2dribbonFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ShiftedRightGrowthCurve2DRibbonFormCallback(
			nil,
			shiftedrightgrowthcurve2dribbonFormCallback.probe,
			newFormGroup,
		)
		shiftedrightgrowthcurve2dribbon := new(models.ShiftedRightGrowthCurve2DRibbon)
		FillUpForm(shiftedrightgrowthcurve2dribbon, newFormGroup, shiftedrightgrowthcurve2dribbonFormCallback.probe)
		shiftedrightgrowthcurve2dribbonFormCallback.probe.formStage.Commit()
	}

	shiftedrightgrowthcurve2dribbonFormCallback.probe.ux_tree()
}
func __gong__New__ShiftedRightGrowthCurve2DRibbonEndShapeFormCallback(
	shiftedrightgrowthcurve2dribbonendshape *models.ShiftedRightGrowthCurve2DRibbonEndShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (shiftedrightgrowthcurve2dribbonendshapeFormCallback *ShiftedRightGrowthCurve2DRibbonEndShapeFormCallback) {
	shiftedrightgrowthcurve2dribbonendshapeFormCallback = new(ShiftedRightGrowthCurve2DRibbonEndShapeFormCallback)
	shiftedrightgrowthcurve2dribbonendshapeFormCallback.probe = probe
	shiftedrightgrowthcurve2dribbonendshapeFormCallback.shiftedrightgrowthcurve2dribbonendshape = shiftedrightgrowthcurve2dribbonendshape
	shiftedrightgrowthcurve2dribbonendshapeFormCallback.formGroup = formGroup

	shiftedrightgrowthcurve2dribbonendshapeFormCallback.CreationMode = (shiftedrightgrowthcurve2dribbonendshape == nil)

	return
}

type ShiftedRightGrowthCurve2DRibbonEndShapeFormCallback struct {
	shiftedrightgrowthcurve2dribbonendshape *models.ShiftedRightGrowthCurve2DRibbonEndShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (shiftedrightgrowthcurve2dribbonendshapeFormCallback *ShiftedRightGrowthCurve2DRibbonEndShapeFormCallback) OnSave() {
	shiftedrightgrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest.Lock()
	defer shiftedrightgrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ShiftedRightGrowthCurve2DRibbonEndShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	shiftedrightgrowthcurve2dribbonendshapeFormCallback.probe.formStage.Checkout()

	if shiftedrightgrowthcurve2dribbonendshapeFormCallback.shiftedrightgrowthcurve2dribbonendshape == nil {
		shiftedrightgrowthcurve2dribbonendshapeFormCallback.shiftedrightgrowthcurve2dribbonendshape = new(models.ShiftedRightGrowthCurve2DRibbonEndShape).Stage(shiftedrightgrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest)
	}
	shiftedrightgrowthcurve2dribbonendshape_ := shiftedrightgrowthcurve2dribbonendshapeFormCallback.shiftedrightgrowthcurve2dribbonendshape
	_ = shiftedrightgrowthcurve2dribbonendshape_

	for _, formDiv := range shiftedrightgrowthcurve2dribbonendshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonendshape_.Name), formDiv)
		case "BottomStartX":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonendshape_.BottomStartX), formDiv)
		case "BottomStartY":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonendshape_.BottomStartY), formDiv)
		case "BottomEndX":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonendshape_.BottomEndX), formDiv)
		case "BottomEndY":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonendshape_.BottomEndY), formDiv)
		case "BottomRadiusX":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonendshape_.BottomRadiusX), formDiv)
		case "BottomRadiusY":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonendshape_.BottomRadiusY), formDiv)
		case "BottomXAxisRotation":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonendshape_.BottomXAxisRotation), formDiv)
		case "BottomLargeArcFlag":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonendshape_.BottomLargeArcFlag), formDiv)
		case "BottomSweepFlag":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonendshape_.BottomSweepFlag), formDiv)
		case "TopStartX":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonendshape_.TopStartX), formDiv)
		case "TopStartY":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonendshape_.TopStartY), formDiv)
		case "TopEndX":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonendshape_.TopEndX), formDiv)
		case "TopEndY":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonendshape_.TopEndY), formDiv)
		case "TopRadiusX":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonendshape_.TopRadiusX), formDiv)
		case "TopRadiusY":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonendshape_.TopRadiusY), formDiv)
		case "TopXAxisRotation":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonendshape_.TopXAxisRotation), formDiv)
		case "TopLargeArcFlag":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonendshape_.TopLargeArcFlag), formDiv)
		case "TopSweepFlag":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonendshape_.TopSweepFlag), formDiv)
		case "ShiftedRightGrowthCurve2DRibbon:ShiftedRightGrowthCurve2DRibbonEndShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the ShiftedRightGrowthCurve2DRibbon instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target ShiftedRightGrowthCurve2DRibbon instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.ShiftedRightGrowthCurve2DRibbon](shiftedrightgrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest)
			targetShiftedRightGrowthCurve2DRibbonIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetShiftedRightGrowthCurve2DRibbonIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all ShiftedRightGrowthCurve2DRibbon instances and update their ShiftedRightGrowthCurve2DRibbonEndShapes slice
			for _shiftedrightgrowthcurve2dribbon := range *models.GetGongstructInstancesSetFromPointerType[*models.ShiftedRightGrowthCurve2DRibbon](shiftedrightgrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(shiftedrightgrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest, _shiftedrightgrowthcurve2dribbon)
				
				// if ShiftedRightGrowthCurve2DRibbon is selected
				if targetShiftedRightGrowthCurve2DRibbonIDs[id] {
					// ensure shiftedrightgrowthcurve2dribbonendshape_ is in _shiftedrightgrowthcurve2dribbon.ShiftedRightGrowthCurve2DRibbonEndShapes
					found := false
					for _, _b := range _shiftedrightgrowthcurve2dribbon.ShiftedRightGrowthCurve2DRibbonEndShapes {
						if _b == shiftedrightgrowthcurve2dribbonendshape_ {
							found = true
							break
						}
					}
					if !found {
						_shiftedrightgrowthcurve2dribbon.ShiftedRightGrowthCurve2DRibbonEndShapes = append(_shiftedrightgrowthcurve2dribbon.ShiftedRightGrowthCurve2DRibbonEndShapes, shiftedrightgrowthcurve2dribbonendshape_)
						shiftedrightgrowthcurve2dribbonendshapeFormCallback.probe.UpdateSliceOfPointersCallback(_shiftedrightgrowthcurve2dribbon, "ShiftedRightGrowthCurve2DRibbonEndShapes", &_shiftedrightgrowthcurve2dribbon.ShiftedRightGrowthCurve2DRibbonEndShapes)
					}
				} else {
					// ensure shiftedrightgrowthcurve2dribbonendshape_ is NOT in _shiftedrightgrowthcurve2dribbon.ShiftedRightGrowthCurve2DRibbonEndShapes
					idx := slices.Index(_shiftedrightgrowthcurve2dribbon.ShiftedRightGrowthCurve2DRibbonEndShapes, shiftedrightgrowthcurve2dribbonendshape_)
					if idx != -1 {
						_shiftedrightgrowthcurve2dribbon.ShiftedRightGrowthCurve2DRibbonEndShapes = slices.Delete(_shiftedrightgrowthcurve2dribbon.ShiftedRightGrowthCurve2DRibbonEndShapes, idx, idx+1)
						shiftedrightgrowthcurve2dribbonendshapeFormCallback.probe.UpdateSliceOfPointersCallback(_shiftedrightgrowthcurve2dribbon, "ShiftedRightGrowthCurve2DRibbonEndShapes", &_shiftedrightgrowthcurve2dribbon.ShiftedRightGrowthCurve2DRibbonEndShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if shiftedrightgrowthcurve2dribbonendshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		shiftedrightgrowthcurve2dribbonendshape_.Unstage(shiftedrightgrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest)
	}

	shiftedrightgrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ShiftedRightGrowthCurve2DRibbonEndShape](
		shiftedrightgrowthcurve2dribbonendshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if shiftedrightgrowthcurve2dribbonendshapeFormCallback.CreationMode || shiftedrightgrowthcurve2dribbonendshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		shiftedrightgrowthcurve2dribbonendshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(shiftedrightgrowthcurve2dribbonendshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ShiftedRightGrowthCurve2DRibbonEndShapeFormCallback(
			nil,
			shiftedrightgrowthcurve2dribbonendshapeFormCallback.probe,
			newFormGroup,
		)
		shiftedrightgrowthcurve2dribbonendshape := new(models.ShiftedRightGrowthCurve2DRibbonEndShape)
		FillUpForm(shiftedrightgrowthcurve2dribbonendshape, newFormGroup, shiftedrightgrowthcurve2dribbonendshapeFormCallback.probe)
		shiftedrightgrowthcurve2dribbonendshapeFormCallback.probe.formStage.Commit()
	}

	shiftedrightgrowthcurve2dribbonendshapeFormCallback.probe.ux_tree()
}
func __gong__New__ShiftedRightGrowthCurve2DRibbonStartShapeFormCallback(
	shiftedrightgrowthcurve2dribbonstartshape *models.ShiftedRightGrowthCurve2DRibbonStartShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (shiftedrightgrowthcurve2dribbonstartshapeFormCallback *ShiftedRightGrowthCurve2DRibbonStartShapeFormCallback) {
	shiftedrightgrowthcurve2dribbonstartshapeFormCallback = new(ShiftedRightGrowthCurve2DRibbonStartShapeFormCallback)
	shiftedrightgrowthcurve2dribbonstartshapeFormCallback.probe = probe
	shiftedrightgrowthcurve2dribbonstartshapeFormCallback.shiftedrightgrowthcurve2dribbonstartshape = shiftedrightgrowthcurve2dribbonstartshape
	shiftedrightgrowthcurve2dribbonstartshapeFormCallback.formGroup = formGroup

	shiftedrightgrowthcurve2dribbonstartshapeFormCallback.CreationMode = (shiftedrightgrowthcurve2dribbonstartshape == nil)

	return
}

type ShiftedRightGrowthCurve2DRibbonStartShapeFormCallback struct {
	shiftedrightgrowthcurve2dribbonstartshape *models.ShiftedRightGrowthCurve2DRibbonStartShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (shiftedrightgrowthcurve2dribbonstartshapeFormCallback *ShiftedRightGrowthCurve2DRibbonStartShapeFormCallback) OnSave() {
	shiftedrightgrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest.Lock()
	defer shiftedrightgrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ShiftedRightGrowthCurve2DRibbonStartShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	shiftedrightgrowthcurve2dribbonstartshapeFormCallback.probe.formStage.Checkout()

	if shiftedrightgrowthcurve2dribbonstartshapeFormCallback.shiftedrightgrowthcurve2dribbonstartshape == nil {
		shiftedrightgrowthcurve2dribbonstartshapeFormCallback.shiftedrightgrowthcurve2dribbonstartshape = new(models.ShiftedRightGrowthCurve2DRibbonStartShape).Stage(shiftedrightgrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest)
	}
	shiftedrightgrowthcurve2dribbonstartshape_ := shiftedrightgrowthcurve2dribbonstartshapeFormCallback.shiftedrightgrowthcurve2dribbonstartshape
	_ = shiftedrightgrowthcurve2dribbonstartshape_

	for _, formDiv := range shiftedrightgrowthcurve2dribbonstartshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonstartshape_.Name), formDiv)
		case "BottomStartX":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonstartshape_.BottomStartX), formDiv)
		case "BottomStartY":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonstartshape_.BottomStartY), formDiv)
		case "BottomEndX":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonstartshape_.BottomEndX), formDiv)
		case "BottomEndY":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonstartshape_.BottomEndY), formDiv)
		case "BottomRadiusX":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonstartshape_.BottomRadiusX), formDiv)
		case "BottomRadiusY":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonstartshape_.BottomRadiusY), formDiv)
		case "BottomXAxisRotation":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonstartshape_.BottomXAxisRotation), formDiv)
		case "BottomLargeArcFlag":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonstartshape_.BottomLargeArcFlag), formDiv)
		case "BottomSweepFlag":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonstartshape_.BottomSweepFlag), formDiv)
		case "TopStartX":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonstartshape_.TopStartX), formDiv)
		case "TopStartY":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonstartshape_.TopStartY), formDiv)
		case "TopEndX":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonstartshape_.TopEndX), formDiv)
		case "TopEndY":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonstartshape_.TopEndY), formDiv)
		case "TopRadiusX":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonstartshape_.TopRadiusX), formDiv)
		case "TopRadiusY":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonstartshape_.TopRadiusY), formDiv)
		case "TopXAxisRotation":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonstartshape_.TopXAxisRotation), formDiv)
		case "TopLargeArcFlag":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonstartshape_.TopLargeArcFlag), formDiv)
		case "TopSweepFlag":
			FormDivBasicFieldToField(&(shiftedrightgrowthcurve2dribbonstartshape_.TopSweepFlag), formDiv)
		case "ShiftedRightGrowthCurve2DRibbon:ShiftedRightGrowthCurve2DRibbonStartShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the ShiftedRightGrowthCurve2DRibbon instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target ShiftedRightGrowthCurve2DRibbon instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.ShiftedRightGrowthCurve2DRibbon](shiftedrightgrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest)
			targetShiftedRightGrowthCurve2DRibbonIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetShiftedRightGrowthCurve2DRibbonIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all ShiftedRightGrowthCurve2DRibbon instances and update their ShiftedRightGrowthCurve2DRibbonStartShapes slice
			for _shiftedrightgrowthcurve2dribbon := range *models.GetGongstructInstancesSetFromPointerType[*models.ShiftedRightGrowthCurve2DRibbon](shiftedrightgrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(shiftedrightgrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest, _shiftedrightgrowthcurve2dribbon)
				
				// if ShiftedRightGrowthCurve2DRibbon is selected
				if targetShiftedRightGrowthCurve2DRibbonIDs[id] {
					// ensure shiftedrightgrowthcurve2dribbonstartshape_ is in _shiftedrightgrowthcurve2dribbon.ShiftedRightGrowthCurve2DRibbonStartShapes
					found := false
					for _, _b := range _shiftedrightgrowthcurve2dribbon.ShiftedRightGrowthCurve2DRibbonStartShapes {
						if _b == shiftedrightgrowthcurve2dribbonstartshape_ {
							found = true
							break
						}
					}
					if !found {
						_shiftedrightgrowthcurve2dribbon.ShiftedRightGrowthCurve2DRibbonStartShapes = append(_shiftedrightgrowthcurve2dribbon.ShiftedRightGrowthCurve2DRibbonStartShapes, shiftedrightgrowthcurve2dribbonstartshape_)
						shiftedrightgrowthcurve2dribbonstartshapeFormCallback.probe.UpdateSliceOfPointersCallback(_shiftedrightgrowthcurve2dribbon, "ShiftedRightGrowthCurve2DRibbonStartShapes", &_shiftedrightgrowthcurve2dribbon.ShiftedRightGrowthCurve2DRibbonStartShapes)
					}
				} else {
					// ensure shiftedrightgrowthcurve2dribbonstartshape_ is NOT in _shiftedrightgrowthcurve2dribbon.ShiftedRightGrowthCurve2DRibbonStartShapes
					idx := slices.Index(_shiftedrightgrowthcurve2dribbon.ShiftedRightGrowthCurve2DRibbonStartShapes, shiftedrightgrowthcurve2dribbonstartshape_)
					if idx != -1 {
						_shiftedrightgrowthcurve2dribbon.ShiftedRightGrowthCurve2DRibbonStartShapes = slices.Delete(_shiftedrightgrowthcurve2dribbon.ShiftedRightGrowthCurve2DRibbonStartShapes, idx, idx+1)
						shiftedrightgrowthcurve2dribbonstartshapeFormCallback.probe.UpdateSliceOfPointersCallback(_shiftedrightgrowthcurve2dribbon, "ShiftedRightGrowthCurve2DRibbonStartShapes", &_shiftedrightgrowthcurve2dribbon.ShiftedRightGrowthCurve2DRibbonStartShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if shiftedrightgrowthcurve2dribbonstartshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		shiftedrightgrowthcurve2dribbonstartshape_.Unstage(shiftedrightgrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest)
	}

	shiftedrightgrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ShiftedRightGrowthCurve2DRibbonStartShape](
		shiftedrightgrowthcurve2dribbonstartshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if shiftedrightgrowthcurve2dribbonstartshapeFormCallback.CreationMode || shiftedrightgrowthcurve2dribbonstartshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		shiftedrightgrowthcurve2dribbonstartshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(shiftedrightgrowthcurve2dribbonstartshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ShiftedRightGrowthCurve2DRibbonStartShapeFormCallback(
			nil,
			shiftedrightgrowthcurve2dribbonstartshapeFormCallback.probe,
			newFormGroup,
		)
		shiftedrightgrowthcurve2dribbonstartshape := new(models.ShiftedRightGrowthCurve2DRibbonStartShape)
		FillUpForm(shiftedrightgrowthcurve2dribbonstartshape, newFormGroup, shiftedrightgrowthcurve2dribbonstartshapeFormCallback.probe)
		shiftedrightgrowthcurve2dribbonstartshapeFormCallback.probe.formStage.Commit()
	}

	shiftedrightgrowthcurve2dribbonstartshapeFormCallback.probe.ux_tree()
}
func __gong__New__StackGrowthCurve2DEndHalfwayArcShapeFormCallback(
	stackgrowthcurve2dendhalfwayarcshape *models.StackGrowthCurve2DEndHalfwayArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackgrowthcurve2dendhalfwayarcshapeFormCallback *StackGrowthCurve2DEndHalfwayArcShapeFormCallback) {
	stackgrowthcurve2dendhalfwayarcshapeFormCallback = new(StackGrowthCurve2DEndHalfwayArcShapeFormCallback)
	stackgrowthcurve2dendhalfwayarcshapeFormCallback.probe = probe
	stackgrowthcurve2dendhalfwayarcshapeFormCallback.stackgrowthcurve2dendhalfwayarcshape = stackgrowthcurve2dendhalfwayarcshape
	stackgrowthcurve2dendhalfwayarcshapeFormCallback.formGroup = formGroup

	stackgrowthcurve2dendhalfwayarcshapeFormCallback.CreationMode = (stackgrowthcurve2dendhalfwayarcshape == nil)

	return
}

type StackGrowthCurve2DEndHalfwayArcShapeFormCallback struct {
	stackgrowthcurve2dendhalfwayarcshape *models.StackGrowthCurve2DEndHalfwayArcShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (stackgrowthcurve2dendhalfwayarcshapeFormCallback *StackGrowthCurve2DEndHalfwayArcShapeFormCallback) OnSave() {
	stackgrowthcurve2dendhalfwayarcshapeFormCallback.probe.stageOfInterest.Lock()
	defer stackgrowthcurve2dendhalfwayarcshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StackGrowthCurve2DEndHalfwayArcShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	stackgrowthcurve2dendhalfwayarcshapeFormCallback.probe.formStage.Checkout()

	if stackgrowthcurve2dendhalfwayarcshapeFormCallback.stackgrowthcurve2dendhalfwayarcshape == nil {
		stackgrowthcurve2dendhalfwayarcshapeFormCallback.stackgrowthcurve2dendhalfwayarcshape = new(models.StackGrowthCurve2DEndHalfwayArcShape).Stage(stackgrowthcurve2dendhalfwayarcshapeFormCallback.probe.stageOfInterest)
	}
	stackgrowthcurve2dendhalfwayarcshape_ := stackgrowthcurve2dendhalfwayarcshapeFormCallback.stackgrowthcurve2dendhalfwayarcshape
	_ = stackgrowthcurve2dendhalfwayarcshape_

	for _, formDiv := range stackgrowthcurve2dendhalfwayarcshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(stackgrowthcurve2dendhalfwayarcshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(stackgrowthcurve2dendhalfwayarcshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(stackgrowthcurve2dendhalfwayarcshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(stackgrowthcurve2dendhalfwayarcshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(stackgrowthcurve2dendhalfwayarcshape_.EndY), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(stackgrowthcurve2dendhalfwayarcshape_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(stackgrowthcurve2dendhalfwayarcshape_.RadiusY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(stackgrowthcurve2dendhalfwayarcshape_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(stackgrowthcurve2dendhalfwayarcshape_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(stackgrowthcurve2dendhalfwayarcshape_.SweepFlag), formDiv)
		case "StackOfGrowthCurve2D:StackGrowthCurve2DEndHalfwayArcShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the StackOfGrowthCurve2D instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target StackOfGrowthCurve2D instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.StackOfGrowthCurve2D](stackgrowthcurve2dendhalfwayarcshapeFormCallback.probe.stageOfInterest)
			targetStackOfGrowthCurve2DIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetStackOfGrowthCurve2DIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all StackOfGrowthCurve2D instances and update their StackGrowthCurve2DEndHalfwayArcShapes slice
			for _stackofgrowthcurve2d := range *models.GetGongstructInstancesSetFromPointerType[*models.StackOfGrowthCurve2D](stackgrowthcurve2dendhalfwayarcshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(stackgrowthcurve2dendhalfwayarcshapeFormCallback.probe.stageOfInterest, _stackofgrowthcurve2d)
				
				// if StackOfGrowthCurve2D is selected
				if targetStackOfGrowthCurve2DIDs[id] {
					// ensure stackgrowthcurve2dendhalfwayarcshape_ is in _stackofgrowthcurve2d.StackGrowthCurve2DEndHalfwayArcShapes
					found := false
					for _, _b := range _stackofgrowthcurve2d.StackGrowthCurve2DEndHalfwayArcShapes {
						if _b == stackgrowthcurve2dendhalfwayarcshape_ {
							found = true
							break
						}
					}
					if !found {
						_stackofgrowthcurve2d.StackGrowthCurve2DEndHalfwayArcShapes = append(_stackofgrowthcurve2d.StackGrowthCurve2DEndHalfwayArcShapes, stackgrowthcurve2dendhalfwayarcshape_)
						stackgrowthcurve2dendhalfwayarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_stackofgrowthcurve2d, "StackGrowthCurve2DEndHalfwayArcShapes", &_stackofgrowthcurve2d.StackGrowthCurve2DEndHalfwayArcShapes)
					}
				} else {
					// ensure stackgrowthcurve2dendhalfwayarcshape_ is NOT in _stackofgrowthcurve2d.StackGrowthCurve2DEndHalfwayArcShapes
					idx := slices.Index(_stackofgrowthcurve2d.StackGrowthCurve2DEndHalfwayArcShapes, stackgrowthcurve2dendhalfwayarcshape_)
					if idx != -1 {
						_stackofgrowthcurve2d.StackGrowthCurve2DEndHalfwayArcShapes = slices.Delete(_stackofgrowthcurve2d.StackGrowthCurve2DEndHalfwayArcShapes, idx, idx+1)
						stackgrowthcurve2dendhalfwayarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_stackofgrowthcurve2d, "StackGrowthCurve2DEndHalfwayArcShapes", &_stackofgrowthcurve2d.StackGrowthCurve2DEndHalfwayArcShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if stackgrowthcurve2dendhalfwayarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackgrowthcurve2dendhalfwayarcshape_.Unstage(stackgrowthcurve2dendhalfwayarcshapeFormCallback.probe.stageOfInterest)
	}

	stackgrowthcurve2dendhalfwayarcshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StackGrowthCurve2DEndHalfwayArcShape](
		stackgrowthcurve2dendhalfwayarcshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if stackgrowthcurve2dendhalfwayarcshapeFormCallback.CreationMode || stackgrowthcurve2dendhalfwayarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackgrowthcurve2dendhalfwayarcshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(stackgrowthcurve2dendhalfwayarcshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StackGrowthCurve2DEndHalfwayArcShapeFormCallback(
			nil,
			stackgrowthcurve2dendhalfwayarcshapeFormCallback.probe,
			newFormGroup,
		)
		stackgrowthcurve2dendhalfwayarcshape := new(models.StackGrowthCurve2DEndHalfwayArcShape)
		FillUpForm(stackgrowthcurve2dendhalfwayarcshape, newFormGroup, stackgrowthcurve2dendhalfwayarcshapeFormCallback.probe)
		stackgrowthcurve2dendhalfwayarcshapeFormCallback.probe.formStage.Commit()
	}

	stackgrowthcurve2dendhalfwayarcshapeFormCallback.probe.ux_tree()
}
func __gong__New__StackGrowthCurve2DRibbonEndShapeFormCallback(
	stackgrowthcurve2dribbonendshape *models.StackGrowthCurve2DRibbonEndShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackgrowthcurve2dribbonendshapeFormCallback *StackGrowthCurve2DRibbonEndShapeFormCallback) {
	stackgrowthcurve2dribbonendshapeFormCallback = new(StackGrowthCurve2DRibbonEndShapeFormCallback)
	stackgrowthcurve2dribbonendshapeFormCallback.probe = probe
	stackgrowthcurve2dribbonendshapeFormCallback.stackgrowthcurve2dribbonendshape = stackgrowthcurve2dribbonendshape
	stackgrowthcurve2dribbonendshapeFormCallback.formGroup = formGroup

	stackgrowthcurve2dribbonendshapeFormCallback.CreationMode = (stackgrowthcurve2dribbonendshape == nil)

	return
}

type StackGrowthCurve2DRibbonEndShapeFormCallback struct {
	stackgrowthcurve2dribbonendshape *models.StackGrowthCurve2DRibbonEndShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (stackgrowthcurve2dribbonendshapeFormCallback *StackGrowthCurve2DRibbonEndShapeFormCallback) OnSave() {
	stackgrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest.Lock()
	defer stackgrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StackGrowthCurve2DRibbonEndShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	stackgrowthcurve2dribbonendshapeFormCallback.probe.formStage.Checkout()

	if stackgrowthcurve2dribbonendshapeFormCallback.stackgrowthcurve2dribbonendshape == nil {
		stackgrowthcurve2dribbonendshapeFormCallback.stackgrowthcurve2dribbonendshape = new(models.StackGrowthCurve2DRibbonEndShape).Stage(stackgrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest)
	}
	stackgrowthcurve2dribbonendshape_ := stackgrowthcurve2dribbonendshapeFormCallback.stackgrowthcurve2dribbonendshape
	_ = stackgrowthcurve2dribbonendshape_

	for _, formDiv := range stackgrowthcurve2dribbonendshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonendshape_.Name), formDiv)
		case "BottomStartX":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonendshape_.BottomStartX), formDiv)
		case "BottomStartY":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonendshape_.BottomStartY), formDiv)
		case "BottomEndX":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonendshape_.BottomEndX), formDiv)
		case "BottomEndY":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonendshape_.BottomEndY), formDiv)
		case "BottomRadiusX":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonendshape_.BottomRadiusX), formDiv)
		case "BottomRadiusY":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonendshape_.BottomRadiusY), formDiv)
		case "BottomXAxisRotation":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonendshape_.BottomXAxisRotation), formDiv)
		case "BottomLargeArcFlag":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonendshape_.BottomLargeArcFlag), formDiv)
		case "BottomSweepFlag":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonendshape_.BottomSweepFlag), formDiv)
		case "TopStartX":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonendshape_.TopStartX), formDiv)
		case "TopStartY":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonendshape_.TopStartY), formDiv)
		case "TopEndX":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonendshape_.TopEndX), formDiv)
		case "TopEndY":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonendshape_.TopEndY), formDiv)
		case "TopRadiusX":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonendshape_.TopRadiusX), formDiv)
		case "TopRadiusY":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonendshape_.TopRadiusY), formDiv)
		case "TopXAxisRotation":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonendshape_.TopXAxisRotation), formDiv)
		case "TopLargeArcFlag":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonendshape_.TopLargeArcFlag), formDiv)
		case "TopSweepFlag":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonendshape_.TopSweepFlag), formDiv)
		case "StackOfGrowthCurve2DRibbon:StackGrowthCurve2DRibbonEndShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the StackOfGrowthCurve2DRibbon instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target StackOfGrowthCurve2DRibbon instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.StackOfGrowthCurve2DRibbon](stackgrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest)
			targetStackOfGrowthCurve2DRibbonIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetStackOfGrowthCurve2DRibbonIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all StackOfGrowthCurve2DRibbon instances and update their StackGrowthCurve2DRibbonEndShapes slice
			for _stackofgrowthcurve2dribbon := range *models.GetGongstructInstancesSetFromPointerType[*models.StackOfGrowthCurve2DRibbon](stackgrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(stackgrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest, _stackofgrowthcurve2dribbon)
				
				// if StackOfGrowthCurve2DRibbon is selected
				if targetStackOfGrowthCurve2DRibbonIDs[id] {
					// ensure stackgrowthcurve2dribbonendshape_ is in _stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonEndShapes
					found := false
					for _, _b := range _stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonEndShapes {
						if _b == stackgrowthcurve2dribbonendshape_ {
							found = true
							break
						}
					}
					if !found {
						_stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonEndShapes = append(_stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonEndShapes, stackgrowthcurve2dribbonendshape_)
						stackgrowthcurve2dribbonendshapeFormCallback.probe.UpdateSliceOfPointersCallback(_stackofgrowthcurve2dribbon, "StackGrowthCurve2DRibbonEndShapes", &_stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonEndShapes)
					}
				} else {
					// ensure stackgrowthcurve2dribbonendshape_ is NOT in _stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonEndShapes
					idx := slices.Index(_stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonEndShapes, stackgrowthcurve2dribbonendshape_)
					if idx != -1 {
						_stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonEndShapes = slices.Delete(_stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonEndShapes, idx, idx+1)
						stackgrowthcurve2dribbonendshapeFormCallback.probe.UpdateSliceOfPointersCallback(_stackofgrowthcurve2dribbon, "StackGrowthCurve2DRibbonEndShapes", &_stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonEndShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if stackgrowthcurve2dribbonendshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackgrowthcurve2dribbonendshape_.Unstage(stackgrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest)
	}

	stackgrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StackGrowthCurve2DRibbonEndShape](
		stackgrowthcurve2dribbonendshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if stackgrowthcurve2dribbonendshapeFormCallback.CreationMode || stackgrowthcurve2dribbonendshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackgrowthcurve2dribbonendshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(stackgrowthcurve2dribbonendshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StackGrowthCurve2DRibbonEndShapeFormCallback(
			nil,
			stackgrowthcurve2dribbonendshapeFormCallback.probe,
			newFormGroup,
		)
		stackgrowthcurve2dribbonendshape := new(models.StackGrowthCurve2DRibbonEndShape)
		FillUpForm(stackgrowthcurve2dribbonendshape, newFormGroup, stackgrowthcurve2dribbonendshapeFormCallback.probe)
		stackgrowthcurve2dribbonendshapeFormCallback.probe.formStage.Commit()
	}

	stackgrowthcurve2dribbonendshapeFormCallback.probe.ux_tree()
}
func __gong__New__StackGrowthCurve2DRibbonStartShapeFormCallback(
	stackgrowthcurve2dribbonstartshape *models.StackGrowthCurve2DRibbonStartShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackgrowthcurve2dribbonstartshapeFormCallback *StackGrowthCurve2DRibbonStartShapeFormCallback) {
	stackgrowthcurve2dribbonstartshapeFormCallback = new(StackGrowthCurve2DRibbonStartShapeFormCallback)
	stackgrowthcurve2dribbonstartshapeFormCallback.probe = probe
	stackgrowthcurve2dribbonstartshapeFormCallback.stackgrowthcurve2dribbonstartshape = stackgrowthcurve2dribbonstartshape
	stackgrowthcurve2dribbonstartshapeFormCallback.formGroup = formGroup

	stackgrowthcurve2dribbonstartshapeFormCallback.CreationMode = (stackgrowthcurve2dribbonstartshape == nil)

	return
}

type StackGrowthCurve2DRibbonStartShapeFormCallback struct {
	stackgrowthcurve2dribbonstartshape *models.StackGrowthCurve2DRibbonStartShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (stackgrowthcurve2dribbonstartshapeFormCallback *StackGrowthCurve2DRibbonStartShapeFormCallback) OnSave() {
	stackgrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest.Lock()
	defer stackgrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StackGrowthCurve2DRibbonStartShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	stackgrowthcurve2dribbonstartshapeFormCallback.probe.formStage.Checkout()

	if stackgrowthcurve2dribbonstartshapeFormCallback.stackgrowthcurve2dribbonstartshape == nil {
		stackgrowthcurve2dribbonstartshapeFormCallback.stackgrowthcurve2dribbonstartshape = new(models.StackGrowthCurve2DRibbonStartShape).Stage(stackgrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest)
	}
	stackgrowthcurve2dribbonstartshape_ := stackgrowthcurve2dribbonstartshapeFormCallback.stackgrowthcurve2dribbonstartshape
	_ = stackgrowthcurve2dribbonstartshape_

	for _, formDiv := range stackgrowthcurve2dribbonstartshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonstartshape_.Name), formDiv)
		case "BottomStartX":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonstartshape_.BottomStartX), formDiv)
		case "BottomStartY":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonstartshape_.BottomStartY), formDiv)
		case "BottomEndX":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonstartshape_.BottomEndX), formDiv)
		case "BottomEndY":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonstartshape_.BottomEndY), formDiv)
		case "BottomRadiusX":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonstartshape_.BottomRadiusX), formDiv)
		case "BottomRadiusY":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonstartshape_.BottomRadiusY), formDiv)
		case "BottomXAxisRotation":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonstartshape_.BottomXAxisRotation), formDiv)
		case "BottomLargeArcFlag":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonstartshape_.BottomLargeArcFlag), formDiv)
		case "BottomSweepFlag":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonstartshape_.BottomSweepFlag), formDiv)
		case "TopStartX":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonstartshape_.TopStartX), formDiv)
		case "TopStartY":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonstartshape_.TopStartY), formDiv)
		case "TopEndX":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonstartshape_.TopEndX), formDiv)
		case "TopEndY":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonstartshape_.TopEndY), formDiv)
		case "TopRadiusX":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonstartshape_.TopRadiusX), formDiv)
		case "TopRadiusY":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonstartshape_.TopRadiusY), formDiv)
		case "TopXAxisRotation":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonstartshape_.TopXAxisRotation), formDiv)
		case "TopLargeArcFlag":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonstartshape_.TopLargeArcFlag), formDiv)
		case "TopSweepFlag":
			FormDivBasicFieldToField(&(stackgrowthcurve2dribbonstartshape_.TopSweepFlag), formDiv)
		case "StackOfGrowthCurve2DRibbon:StackGrowthCurve2DRibbonStartShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the StackOfGrowthCurve2DRibbon instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target StackOfGrowthCurve2DRibbon instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.StackOfGrowthCurve2DRibbon](stackgrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest)
			targetStackOfGrowthCurve2DRibbonIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetStackOfGrowthCurve2DRibbonIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all StackOfGrowthCurve2DRibbon instances and update their StackGrowthCurve2DRibbonStartShapes slice
			for _stackofgrowthcurve2dribbon := range *models.GetGongstructInstancesSetFromPointerType[*models.StackOfGrowthCurve2DRibbon](stackgrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(stackgrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest, _stackofgrowthcurve2dribbon)
				
				// if StackOfGrowthCurve2DRibbon is selected
				if targetStackOfGrowthCurve2DRibbonIDs[id] {
					// ensure stackgrowthcurve2dribbonstartshape_ is in _stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonStartShapes
					found := false
					for _, _b := range _stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonStartShapes {
						if _b == stackgrowthcurve2dribbonstartshape_ {
							found = true
							break
						}
					}
					if !found {
						_stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonStartShapes = append(_stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonStartShapes, stackgrowthcurve2dribbonstartshape_)
						stackgrowthcurve2dribbonstartshapeFormCallback.probe.UpdateSliceOfPointersCallback(_stackofgrowthcurve2dribbon, "StackGrowthCurve2DRibbonStartShapes", &_stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonStartShapes)
					}
				} else {
					// ensure stackgrowthcurve2dribbonstartshape_ is NOT in _stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonStartShapes
					idx := slices.Index(_stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonStartShapes, stackgrowthcurve2dribbonstartshape_)
					if idx != -1 {
						_stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonStartShapes = slices.Delete(_stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonStartShapes, idx, idx+1)
						stackgrowthcurve2dribbonstartshapeFormCallback.probe.UpdateSliceOfPointersCallback(_stackofgrowthcurve2dribbon, "StackGrowthCurve2DRibbonStartShapes", &_stackofgrowthcurve2dribbon.StackGrowthCurve2DRibbonStartShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if stackgrowthcurve2dribbonstartshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackgrowthcurve2dribbonstartshape_.Unstage(stackgrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest)
	}

	stackgrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StackGrowthCurve2DRibbonStartShape](
		stackgrowthcurve2dribbonstartshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if stackgrowthcurve2dribbonstartshapeFormCallback.CreationMode || stackgrowthcurve2dribbonstartshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackgrowthcurve2dribbonstartshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(stackgrowthcurve2dribbonstartshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StackGrowthCurve2DRibbonStartShapeFormCallback(
			nil,
			stackgrowthcurve2dribbonstartshapeFormCallback.probe,
			newFormGroup,
		)
		stackgrowthcurve2dribbonstartshape := new(models.StackGrowthCurve2DRibbonStartShape)
		FillUpForm(stackgrowthcurve2dribbonstartshape, newFormGroup, stackgrowthcurve2dribbonstartshapeFormCallback.probe)
		stackgrowthcurve2dribbonstartshapeFormCallback.probe.formStage.Commit()
	}

	stackgrowthcurve2dribbonstartshapeFormCallback.probe.ux_tree()
}
func __gong__New__StackGrowthCurve2DStartHalfwayArcShapeFormCallback(
	stackgrowthcurve2dstarthalfwayarcshape *models.StackGrowthCurve2DStartHalfwayArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackgrowthcurve2dstarthalfwayarcshapeFormCallback *StackGrowthCurve2DStartHalfwayArcShapeFormCallback) {
	stackgrowthcurve2dstarthalfwayarcshapeFormCallback = new(StackGrowthCurve2DStartHalfwayArcShapeFormCallback)
	stackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe = probe
	stackgrowthcurve2dstarthalfwayarcshapeFormCallback.stackgrowthcurve2dstarthalfwayarcshape = stackgrowthcurve2dstarthalfwayarcshape
	stackgrowthcurve2dstarthalfwayarcshapeFormCallback.formGroup = formGroup

	stackgrowthcurve2dstarthalfwayarcshapeFormCallback.CreationMode = (stackgrowthcurve2dstarthalfwayarcshape == nil)

	return
}

type StackGrowthCurve2DStartHalfwayArcShapeFormCallback struct {
	stackgrowthcurve2dstarthalfwayarcshape *models.StackGrowthCurve2DStartHalfwayArcShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (stackgrowthcurve2dstarthalfwayarcshapeFormCallback *StackGrowthCurve2DStartHalfwayArcShapeFormCallback) OnSave() {
	stackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe.stageOfInterest.Lock()
	defer stackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StackGrowthCurve2DStartHalfwayArcShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	stackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe.formStage.Checkout()

	if stackgrowthcurve2dstarthalfwayarcshapeFormCallback.stackgrowthcurve2dstarthalfwayarcshape == nil {
		stackgrowthcurve2dstarthalfwayarcshapeFormCallback.stackgrowthcurve2dstarthalfwayarcshape = new(models.StackGrowthCurve2DStartHalfwayArcShape).Stage(stackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe.stageOfInterest)
	}
	stackgrowthcurve2dstarthalfwayarcshape_ := stackgrowthcurve2dstarthalfwayarcshapeFormCallback.stackgrowthcurve2dstarthalfwayarcshape
	_ = stackgrowthcurve2dstarthalfwayarcshape_

	for _, formDiv := range stackgrowthcurve2dstarthalfwayarcshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(stackgrowthcurve2dstarthalfwayarcshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(stackgrowthcurve2dstarthalfwayarcshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(stackgrowthcurve2dstarthalfwayarcshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(stackgrowthcurve2dstarthalfwayarcshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(stackgrowthcurve2dstarthalfwayarcshape_.EndY), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(stackgrowthcurve2dstarthalfwayarcshape_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(stackgrowthcurve2dstarthalfwayarcshape_.RadiusY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(stackgrowthcurve2dstarthalfwayarcshape_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(stackgrowthcurve2dstarthalfwayarcshape_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(stackgrowthcurve2dstarthalfwayarcshape_.SweepFlag), formDiv)
		case "StackOfGrowthCurve2D:StackGrowthCurve2DStartHalfwayArcShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the StackOfGrowthCurve2D instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target StackOfGrowthCurve2D instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.StackOfGrowthCurve2D](stackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe.stageOfInterest)
			targetStackOfGrowthCurve2DIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetStackOfGrowthCurve2DIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all StackOfGrowthCurve2D instances and update their StackGrowthCurve2DStartHalfwayArcShapes slice
			for _stackofgrowthcurve2d := range *models.GetGongstructInstancesSetFromPointerType[*models.StackOfGrowthCurve2D](stackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(stackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe.stageOfInterest, _stackofgrowthcurve2d)
				
				// if StackOfGrowthCurve2D is selected
				if targetStackOfGrowthCurve2DIDs[id] {
					// ensure stackgrowthcurve2dstarthalfwayarcshape_ is in _stackofgrowthcurve2d.StackGrowthCurve2DStartHalfwayArcShapes
					found := false
					for _, _b := range _stackofgrowthcurve2d.StackGrowthCurve2DStartHalfwayArcShapes {
						if _b == stackgrowthcurve2dstarthalfwayarcshape_ {
							found = true
							break
						}
					}
					if !found {
						_stackofgrowthcurve2d.StackGrowthCurve2DStartHalfwayArcShapes = append(_stackofgrowthcurve2d.StackGrowthCurve2DStartHalfwayArcShapes, stackgrowthcurve2dstarthalfwayarcshape_)
						stackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_stackofgrowthcurve2d, "StackGrowthCurve2DStartHalfwayArcShapes", &_stackofgrowthcurve2d.StackGrowthCurve2DStartHalfwayArcShapes)
					}
				} else {
					// ensure stackgrowthcurve2dstarthalfwayarcshape_ is NOT in _stackofgrowthcurve2d.StackGrowthCurve2DStartHalfwayArcShapes
					idx := slices.Index(_stackofgrowthcurve2d.StackGrowthCurve2DStartHalfwayArcShapes, stackgrowthcurve2dstarthalfwayarcshape_)
					if idx != -1 {
						_stackofgrowthcurve2d.StackGrowthCurve2DStartHalfwayArcShapes = slices.Delete(_stackofgrowthcurve2d.StackGrowthCurve2DStartHalfwayArcShapes, idx, idx+1)
						stackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_stackofgrowthcurve2d, "StackGrowthCurve2DStartHalfwayArcShapes", &_stackofgrowthcurve2d.StackGrowthCurve2DStartHalfwayArcShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if stackgrowthcurve2dstarthalfwayarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackgrowthcurve2dstarthalfwayarcshape_.Unstage(stackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe.stageOfInterest)
	}

	stackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StackGrowthCurve2DStartHalfwayArcShape](
		stackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if stackgrowthcurve2dstarthalfwayarcshapeFormCallback.CreationMode || stackgrowthcurve2dstarthalfwayarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(stackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StackGrowthCurve2DStartHalfwayArcShapeFormCallback(
			nil,
			stackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe,
			newFormGroup,
		)
		stackgrowthcurve2dstarthalfwayarcshape := new(models.StackGrowthCurve2DStartHalfwayArcShape)
		FillUpForm(stackgrowthcurve2dstarthalfwayarcshape, newFormGroup, stackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe)
		stackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe.formStage.Commit()
	}

	stackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe.ux_tree()
}
func __gong__New__StackOfGrowthCurve2DFormCallback(
	stackofgrowthcurve2d *models.StackOfGrowthCurve2D,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackofgrowthcurve2dFormCallback *StackOfGrowthCurve2DFormCallback) {
	stackofgrowthcurve2dFormCallback = new(StackOfGrowthCurve2DFormCallback)
	stackofgrowthcurve2dFormCallback.probe = probe
	stackofgrowthcurve2dFormCallback.stackofgrowthcurve2d = stackofgrowthcurve2d
	stackofgrowthcurve2dFormCallback.formGroup = formGroup

	stackofgrowthcurve2dFormCallback.CreationMode = (stackofgrowthcurve2d == nil)

	return
}

type StackOfGrowthCurve2DFormCallback struct {
	stackofgrowthcurve2d *models.StackOfGrowthCurve2D

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (stackofgrowthcurve2dFormCallback *StackOfGrowthCurve2DFormCallback) OnSave() {
	stackofgrowthcurve2dFormCallback.probe.stageOfInterest.Lock()
	defer stackofgrowthcurve2dFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StackOfGrowthCurve2DFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	stackofgrowthcurve2dFormCallback.probe.formStage.Checkout()

	if stackofgrowthcurve2dFormCallback.stackofgrowthcurve2d == nil {
		stackofgrowthcurve2dFormCallback.stackofgrowthcurve2d = new(models.StackOfGrowthCurve2D).Stage(stackofgrowthcurve2dFormCallback.probe.stageOfInterest)
	}
	stackofgrowthcurve2d_ := stackofgrowthcurve2dFormCallback.stackofgrowthcurve2d
	_ = stackofgrowthcurve2d_

	for _, formDiv := range stackofgrowthcurve2dFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(stackofgrowthcurve2d_.Name), formDiv)
		case "StackGrowthCurve2DStartHalfwayArcShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.StackGrowthCurve2DStartHalfwayArcShape](stackofgrowthcurve2dFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.StackGrowthCurve2DStartHalfwayArcShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.StackGrowthCurve2DStartHalfwayArcShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					stackofgrowthcurve2dFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.StackGrowthCurve2DStartHalfwayArcShape](stackofgrowthcurve2dFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			stackofgrowthcurve2d_.StackGrowthCurve2DStartHalfwayArcShapes = instanceSlice
			stackofgrowthcurve2dFormCallback.probe.UpdateSliceOfPointersCallback(stackofgrowthcurve2d_, "StackGrowthCurve2DStartHalfwayArcShapes", &stackofgrowthcurve2d_.StackGrowthCurve2DStartHalfwayArcShapes)

		case "StackGrowthCurve2DEndHalfwayArcShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.StackGrowthCurve2DEndHalfwayArcShape](stackofgrowthcurve2dFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.StackGrowthCurve2DEndHalfwayArcShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.StackGrowthCurve2DEndHalfwayArcShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					stackofgrowthcurve2dFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.StackGrowthCurve2DEndHalfwayArcShape](stackofgrowthcurve2dFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			stackofgrowthcurve2d_.StackGrowthCurve2DEndHalfwayArcShapes = instanceSlice
			stackofgrowthcurve2dFormCallback.probe.UpdateSliceOfPointersCallback(stackofgrowthcurve2d_, "StackGrowthCurve2DEndHalfwayArcShapes", &stackofgrowthcurve2d_.StackGrowthCurve2DEndHalfwayArcShapes)

		}
	}

	// manage the suppress operation
	if stackofgrowthcurve2dFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackofgrowthcurve2d_.Unstage(stackofgrowthcurve2dFormCallback.probe.stageOfInterest)
	}

	stackofgrowthcurve2dFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StackOfGrowthCurve2D](
		stackofgrowthcurve2dFormCallback.probe,
	)

	// display a new form by reset the form stage
	if stackofgrowthcurve2dFormCallback.CreationMode || stackofgrowthcurve2dFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackofgrowthcurve2dFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(stackofgrowthcurve2dFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StackOfGrowthCurve2DFormCallback(
			nil,
			stackofgrowthcurve2dFormCallback.probe,
			newFormGroup,
		)
		stackofgrowthcurve2d := new(models.StackOfGrowthCurve2D)
		FillUpForm(stackofgrowthcurve2d, newFormGroup, stackofgrowthcurve2dFormCallback.probe)
		stackofgrowthcurve2dFormCallback.probe.formStage.Commit()
	}

	stackofgrowthcurve2dFormCallback.probe.ux_tree()
}
func __gong__New__StackOfGrowthCurve2DRibbonFormCallback(
	stackofgrowthcurve2dribbon *models.StackOfGrowthCurve2DRibbon,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackofgrowthcurve2dribbonFormCallback *StackOfGrowthCurve2DRibbonFormCallback) {
	stackofgrowthcurve2dribbonFormCallback = new(StackOfGrowthCurve2DRibbonFormCallback)
	stackofgrowthcurve2dribbonFormCallback.probe = probe
	stackofgrowthcurve2dribbonFormCallback.stackofgrowthcurve2dribbon = stackofgrowthcurve2dribbon
	stackofgrowthcurve2dribbonFormCallback.formGroup = formGroup

	stackofgrowthcurve2dribbonFormCallback.CreationMode = (stackofgrowthcurve2dribbon == nil)

	return
}

type StackOfGrowthCurve2DRibbonFormCallback struct {
	stackofgrowthcurve2dribbon *models.StackOfGrowthCurve2DRibbon

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (stackofgrowthcurve2dribbonFormCallback *StackOfGrowthCurve2DRibbonFormCallback) OnSave() {
	stackofgrowthcurve2dribbonFormCallback.probe.stageOfInterest.Lock()
	defer stackofgrowthcurve2dribbonFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StackOfGrowthCurve2DRibbonFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	stackofgrowthcurve2dribbonFormCallback.probe.formStage.Checkout()

	if stackofgrowthcurve2dribbonFormCallback.stackofgrowthcurve2dribbon == nil {
		stackofgrowthcurve2dribbonFormCallback.stackofgrowthcurve2dribbon = new(models.StackOfGrowthCurve2DRibbon).Stage(stackofgrowthcurve2dribbonFormCallback.probe.stageOfInterest)
	}
	stackofgrowthcurve2dribbon_ := stackofgrowthcurve2dribbonFormCallback.stackofgrowthcurve2dribbon
	_ = stackofgrowthcurve2dribbon_

	for _, formDiv := range stackofgrowthcurve2dribbonFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(stackofgrowthcurve2dribbon_.Name), formDiv)
		case "StackGrowthCurve2DRibbonStartShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.StackGrowthCurve2DRibbonStartShape](stackofgrowthcurve2dribbonFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.StackGrowthCurve2DRibbonStartShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.StackGrowthCurve2DRibbonStartShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					stackofgrowthcurve2dribbonFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.StackGrowthCurve2DRibbonStartShape](stackofgrowthcurve2dribbonFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			stackofgrowthcurve2dribbon_.StackGrowthCurve2DRibbonStartShapes = instanceSlice
			stackofgrowthcurve2dribbonFormCallback.probe.UpdateSliceOfPointersCallback(stackofgrowthcurve2dribbon_, "StackGrowthCurve2DRibbonStartShapes", &stackofgrowthcurve2dribbon_.StackGrowthCurve2DRibbonStartShapes)

		case "StackGrowthCurve2DRibbonEndShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.StackGrowthCurve2DRibbonEndShape](stackofgrowthcurve2dribbonFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.StackGrowthCurve2DRibbonEndShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.StackGrowthCurve2DRibbonEndShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					stackofgrowthcurve2dribbonFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.StackGrowthCurve2DRibbonEndShape](stackofgrowthcurve2dribbonFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			stackofgrowthcurve2dribbon_.StackGrowthCurve2DRibbonEndShapes = instanceSlice
			stackofgrowthcurve2dribbonFormCallback.probe.UpdateSliceOfPointersCallback(stackofgrowthcurve2dribbon_, "StackGrowthCurve2DRibbonEndShapes", &stackofgrowthcurve2dribbon_.StackGrowthCurve2DRibbonEndShapes)

		}
	}

	// manage the suppress operation
	if stackofgrowthcurve2dribbonFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackofgrowthcurve2dribbon_.Unstage(stackofgrowthcurve2dribbonFormCallback.probe.stageOfInterest)
	}

	stackofgrowthcurve2dribbonFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StackOfGrowthCurve2DRibbon](
		stackofgrowthcurve2dribbonFormCallback.probe,
	)

	// display a new form by reset the form stage
	if stackofgrowthcurve2dribbonFormCallback.CreationMode || stackofgrowthcurve2dribbonFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackofgrowthcurve2dribbonFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(stackofgrowthcurve2dribbonFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StackOfGrowthCurve2DRibbonFormCallback(
			nil,
			stackofgrowthcurve2dribbonFormCallback.probe,
			newFormGroup,
		)
		stackofgrowthcurve2dribbon := new(models.StackOfGrowthCurve2DRibbon)
		FillUpForm(stackofgrowthcurve2dribbon, newFormGroup, stackofgrowthcurve2dribbonFormCallback.probe)
		stackofgrowthcurve2dribbonFormCallback.probe.formStage.Commit()
	}

	stackofgrowthcurve2dribbonFormCallback.probe.ux_tree()
}
func __gong__New__StackOfPartiallyRotatedTorusShapeFormCallback(
	stackofpartiallyrotatedtorusshape *models.StackOfPartiallyRotatedTorusShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackofpartiallyrotatedtorusshapeFormCallback *StackOfPartiallyRotatedTorusShapeFormCallback) {
	stackofpartiallyrotatedtorusshapeFormCallback = new(StackOfPartiallyRotatedTorusShapeFormCallback)
	stackofpartiallyrotatedtorusshapeFormCallback.probe = probe
	stackofpartiallyrotatedtorusshapeFormCallback.stackofpartiallyrotatedtorusshape = stackofpartiallyrotatedtorusshape
	stackofpartiallyrotatedtorusshapeFormCallback.formGroup = formGroup

	stackofpartiallyrotatedtorusshapeFormCallback.CreationMode = (stackofpartiallyrotatedtorusshape == nil)

	return
}

type StackOfPartiallyRotatedTorusShapeFormCallback struct {
	stackofpartiallyrotatedtorusshape *models.StackOfPartiallyRotatedTorusShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (stackofpartiallyrotatedtorusshapeFormCallback *StackOfPartiallyRotatedTorusShapeFormCallback) OnSave() {
	stackofpartiallyrotatedtorusshapeFormCallback.probe.stageOfInterest.Lock()
	defer stackofpartiallyrotatedtorusshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StackOfPartiallyRotatedTorusShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	stackofpartiallyrotatedtorusshapeFormCallback.probe.formStage.Checkout()

	if stackofpartiallyrotatedtorusshapeFormCallback.stackofpartiallyrotatedtorusshape == nil {
		stackofpartiallyrotatedtorusshapeFormCallback.stackofpartiallyrotatedtorusshape = new(models.StackOfPartiallyRotatedTorusShape).Stage(stackofpartiallyrotatedtorusshapeFormCallback.probe.stageOfInterest)
	}
	stackofpartiallyrotatedtorusshape_ := stackofpartiallyrotatedtorusshapeFormCallback.stackofpartiallyrotatedtorusshape
	_ = stackofpartiallyrotatedtorusshape_

	for _, formDiv := range stackofpartiallyrotatedtorusshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(stackofpartiallyrotatedtorusshape_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if stackofpartiallyrotatedtorusshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackofpartiallyrotatedtorusshape_.Unstage(stackofpartiallyrotatedtorusshapeFormCallback.probe.stageOfInterest)
	}

	stackofpartiallyrotatedtorusshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StackOfPartiallyRotatedTorusShape](
		stackofpartiallyrotatedtorusshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if stackofpartiallyrotatedtorusshapeFormCallback.CreationMode || stackofpartiallyrotatedtorusshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackofpartiallyrotatedtorusshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(stackofpartiallyrotatedtorusshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StackOfPartiallyRotatedTorusShapeFormCallback(
			nil,
			stackofpartiallyrotatedtorusshapeFormCallback.probe,
			newFormGroup,
		)
		stackofpartiallyrotatedtorusshape := new(models.StackOfPartiallyRotatedTorusShape)
		FillUpForm(stackofpartiallyrotatedtorusshape, newFormGroup, stackofpartiallyrotatedtorusshapeFormCallback.probe)
		stackofpartiallyrotatedtorusshapeFormCallback.probe.formStage.Commit()
	}

	stackofpartiallyrotatedtorusshapeFormCallback.probe.ux_tree()
}
func __gong__New__StackOfRotatedGrowthCurve2DFormCallback(
	stackofrotatedgrowthcurve2d *models.StackOfRotatedGrowthCurve2D,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackofrotatedgrowthcurve2dFormCallback *StackOfRotatedGrowthCurve2DFormCallback) {
	stackofrotatedgrowthcurve2dFormCallback = new(StackOfRotatedGrowthCurve2DFormCallback)
	stackofrotatedgrowthcurve2dFormCallback.probe = probe
	stackofrotatedgrowthcurve2dFormCallback.stackofrotatedgrowthcurve2d = stackofrotatedgrowthcurve2d
	stackofrotatedgrowthcurve2dFormCallback.formGroup = formGroup

	stackofrotatedgrowthcurve2dFormCallback.CreationMode = (stackofrotatedgrowthcurve2d == nil)

	return
}

type StackOfRotatedGrowthCurve2DFormCallback struct {
	stackofrotatedgrowthcurve2d *models.StackOfRotatedGrowthCurve2D

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (stackofrotatedgrowthcurve2dFormCallback *StackOfRotatedGrowthCurve2DFormCallback) OnSave() {
	stackofrotatedgrowthcurve2dFormCallback.probe.stageOfInterest.Lock()
	defer stackofrotatedgrowthcurve2dFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StackOfRotatedGrowthCurve2DFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	stackofrotatedgrowthcurve2dFormCallback.probe.formStage.Checkout()

	if stackofrotatedgrowthcurve2dFormCallback.stackofrotatedgrowthcurve2d == nil {
		stackofrotatedgrowthcurve2dFormCallback.stackofrotatedgrowthcurve2d = new(models.StackOfRotatedGrowthCurve2D).Stage(stackofrotatedgrowthcurve2dFormCallback.probe.stageOfInterest)
	}
	stackofrotatedgrowthcurve2d_ := stackofrotatedgrowthcurve2dFormCallback.stackofrotatedgrowthcurve2d
	_ = stackofrotatedgrowthcurve2d_

	for _, formDiv := range stackofrotatedgrowthcurve2dFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(stackofrotatedgrowthcurve2d_.Name), formDiv)
		case "StackRotatedGrowthCurve2DStartArcShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.StackRotatedGrowthCurve2DStartArcShape](stackofrotatedgrowthcurve2dFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.StackRotatedGrowthCurve2DStartArcShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.StackRotatedGrowthCurve2DStartArcShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					stackofrotatedgrowthcurve2dFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.StackRotatedGrowthCurve2DStartArcShape](stackofrotatedgrowthcurve2dFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			stackofrotatedgrowthcurve2d_.StackRotatedGrowthCurve2DStartArcShapes = instanceSlice
			stackofrotatedgrowthcurve2dFormCallback.probe.UpdateSliceOfPointersCallback(stackofrotatedgrowthcurve2d_, "StackRotatedGrowthCurve2DStartArcShapes", &stackofrotatedgrowthcurve2d_.StackRotatedGrowthCurve2DStartArcShapes)

		case "StackRotatedGrowthCurve2DEndArcShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.StackRotatedGrowthCurve2DEndArcShape](stackofrotatedgrowthcurve2dFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.StackRotatedGrowthCurve2DEndArcShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.StackRotatedGrowthCurve2DEndArcShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					stackofrotatedgrowthcurve2dFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.StackRotatedGrowthCurve2DEndArcShape](stackofrotatedgrowthcurve2dFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			stackofrotatedgrowthcurve2d_.StackRotatedGrowthCurve2DEndArcShapes = instanceSlice
			stackofrotatedgrowthcurve2dFormCallback.probe.UpdateSliceOfPointersCallback(stackofrotatedgrowthcurve2d_, "StackRotatedGrowthCurve2DEndArcShapes", &stackofrotatedgrowthcurve2d_.StackRotatedGrowthCurve2DEndArcShapes)

		}
	}

	// manage the suppress operation
	if stackofrotatedgrowthcurve2dFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackofrotatedgrowthcurve2d_.Unstage(stackofrotatedgrowthcurve2dFormCallback.probe.stageOfInterest)
	}

	stackofrotatedgrowthcurve2dFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StackOfRotatedGrowthCurve2D](
		stackofrotatedgrowthcurve2dFormCallback.probe,
	)

	// display a new form by reset the form stage
	if stackofrotatedgrowthcurve2dFormCallback.CreationMode || stackofrotatedgrowthcurve2dFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackofrotatedgrowthcurve2dFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(stackofrotatedgrowthcurve2dFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StackOfRotatedGrowthCurve2DFormCallback(
			nil,
			stackofrotatedgrowthcurve2dFormCallback.probe,
			newFormGroup,
		)
		stackofrotatedgrowthcurve2d := new(models.StackOfRotatedGrowthCurve2D)
		FillUpForm(stackofrotatedgrowthcurve2d, newFormGroup, stackofrotatedgrowthcurve2dFormCallback.probe)
		stackofrotatedgrowthcurve2dFormCallback.probe.formStage.Commit()
	}

	stackofrotatedgrowthcurve2dFormCallback.probe.ux_tree()
}
func __gong__New__StackOfRotatedGrowthCurve2DRibbonFormCallback(
	stackofrotatedgrowthcurve2dribbon *models.StackOfRotatedGrowthCurve2DRibbon,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackofrotatedgrowthcurve2dribbonFormCallback *StackOfRotatedGrowthCurve2DRibbonFormCallback) {
	stackofrotatedgrowthcurve2dribbonFormCallback = new(StackOfRotatedGrowthCurve2DRibbonFormCallback)
	stackofrotatedgrowthcurve2dribbonFormCallback.probe = probe
	stackofrotatedgrowthcurve2dribbonFormCallback.stackofrotatedgrowthcurve2dribbon = stackofrotatedgrowthcurve2dribbon
	stackofrotatedgrowthcurve2dribbonFormCallback.formGroup = formGroup

	stackofrotatedgrowthcurve2dribbonFormCallback.CreationMode = (stackofrotatedgrowthcurve2dribbon == nil)

	return
}

type StackOfRotatedGrowthCurve2DRibbonFormCallback struct {
	stackofrotatedgrowthcurve2dribbon *models.StackOfRotatedGrowthCurve2DRibbon

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (stackofrotatedgrowthcurve2dribbonFormCallback *StackOfRotatedGrowthCurve2DRibbonFormCallback) OnSave() {
	stackofrotatedgrowthcurve2dribbonFormCallback.probe.stageOfInterest.Lock()
	defer stackofrotatedgrowthcurve2dribbonFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StackOfRotatedGrowthCurve2DRibbonFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	stackofrotatedgrowthcurve2dribbonFormCallback.probe.formStage.Checkout()

	if stackofrotatedgrowthcurve2dribbonFormCallback.stackofrotatedgrowthcurve2dribbon == nil {
		stackofrotatedgrowthcurve2dribbonFormCallback.stackofrotatedgrowthcurve2dribbon = new(models.StackOfRotatedGrowthCurve2DRibbon).Stage(stackofrotatedgrowthcurve2dribbonFormCallback.probe.stageOfInterest)
	}
	stackofrotatedgrowthcurve2dribbon_ := stackofrotatedgrowthcurve2dribbonFormCallback.stackofrotatedgrowthcurve2dribbon
	_ = stackofrotatedgrowthcurve2dribbon_

	for _, formDiv := range stackofrotatedgrowthcurve2dribbonFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(stackofrotatedgrowthcurve2dribbon_.Name), formDiv)
		case "StackRotatedGrowthCurve2DRibbonStartShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.StackRotatedGrowthCurve2DRibbonStartShape](stackofrotatedgrowthcurve2dribbonFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.StackRotatedGrowthCurve2DRibbonStartShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.StackRotatedGrowthCurve2DRibbonStartShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					stackofrotatedgrowthcurve2dribbonFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.StackRotatedGrowthCurve2DRibbonStartShape](stackofrotatedgrowthcurve2dribbonFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			stackofrotatedgrowthcurve2dribbon_.StackRotatedGrowthCurve2DRibbonStartShapes = instanceSlice
			stackofrotatedgrowthcurve2dribbonFormCallback.probe.UpdateSliceOfPointersCallback(stackofrotatedgrowthcurve2dribbon_, "StackRotatedGrowthCurve2DRibbonStartShapes", &stackofrotatedgrowthcurve2dribbon_.StackRotatedGrowthCurve2DRibbonStartShapes)

		case "StackRotatedGrowthCurve2DRibbonEndShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.StackRotatedGrowthCurve2DRibbonEndShape](stackofrotatedgrowthcurve2dribbonFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.StackRotatedGrowthCurve2DRibbonEndShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.StackRotatedGrowthCurve2DRibbonEndShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					stackofrotatedgrowthcurve2dribbonFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.StackRotatedGrowthCurve2DRibbonEndShape](stackofrotatedgrowthcurve2dribbonFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			stackofrotatedgrowthcurve2dribbon_.StackRotatedGrowthCurve2DRibbonEndShapes = instanceSlice
			stackofrotatedgrowthcurve2dribbonFormCallback.probe.UpdateSliceOfPointersCallback(stackofrotatedgrowthcurve2dribbon_, "StackRotatedGrowthCurve2DRibbonEndShapes", &stackofrotatedgrowthcurve2dribbon_.StackRotatedGrowthCurve2DRibbonEndShapes)

		}
	}

	// manage the suppress operation
	if stackofrotatedgrowthcurve2dribbonFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackofrotatedgrowthcurve2dribbon_.Unstage(stackofrotatedgrowthcurve2dribbonFormCallback.probe.stageOfInterest)
	}

	stackofrotatedgrowthcurve2dribbonFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StackOfRotatedGrowthCurve2DRibbon](
		stackofrotatedgrowthcurve2dribbonFormCallback.probe,
	)

	// display a new form by reset the form stage
	if stackofrotatedgrowthcurve2dribbonFormCallback.CreationMode || stackofrotatedgrowthcurve2dribbonFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackofrotatedgrowthcurve2dribbonFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(stackofrotatedgrowthcurve2dribbonFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StackOfRotatedGrowthCurve2DRibbonFormCallback(
			nil,
			stackofrotatedgrowthcurve2dribbonFormCallback.probe,
			newFormGroup,
		)
		stackofrotatedgrowthcurve2dribbon := new(models.StackOfRotatedGrowthCurve2DRibbon)
		FillUpForm(stackofrotatedgrowthcurve2dribbon, newFormGroup, stackofrotatedgrowthcurve2dribbonFormCallback.probe)
		stackofrotatedgrowthcurve2dribbonFormCallback.probe.formStage.Commit()
	}

	stackofrotatedgrowthcurve2dribbonFormCallback.probe.ux_tree()
}
func __gong__New__StackRotatedGrowthCurve2DEndArcShapeFormCallback(
	stackrotatedgrowthcurve2dendarcshape *models.StackRotatedGrowthCurve2DEndArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackrotatedgrowthcurve2dendarcshapeFormCallback *StackRotatedGrowthCurve2DEndArcShapeFormCallback) {
	stackrotatedgrowthcurve2dendarcshapeFormCallback = new(StackRotatedGrowthCurve2DEndArcShapeFormCallback)
	stackrotatedgrowthcurve2dendarcshapeFormCallback.probe = probe
	stackrotatedgrowthcurve2dendarcshapeFormCallback.stackrotatedgrowthcurve2dendarcshape = stackrotatedgrowthcurve2dendarcshape
	stackrotatedgrowthcurve2dendarcshapeFormCallback.formGroup = formGroup

	stackrotatedgrowthcurve2dendarcshapeFormCallback.CreationMode = (stackrotatedgrowthcurve2dendarcshape == nil)

	return
}

type StackRotatedGrowthCurve2DEndArcShapeFormCallback struct {
	stackrotatedgrowthcurve2dendarcshape *models.StackRotatedGrowthCurve2DEndArcShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (stackrotatedgrowthcurve2dendarcshapeFormCallback *StackRotatedGrowthCurve2DEndArcShapeFormCallback) OnSave() {
	stackrotatedgrowthcurve2dendarcshapeFormCallback.probe.stageOfInterest.Lock()
	defer stackrotatedgrowthcurve2dendarcshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StackRotatedGrowthCurve2DEndArcShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	stackrotatedgrowthcurve2dendarcshapeFormCallback.probe.formStage.Checkout()

	if stackrotatedgrowthcurve2dendarcshapeFormCallback.stackrotatedgrowthcurve2dendarcshape == nil {
		stackrotatedgrowthcurve2dendarcshapeFormCallback.stackrotatedgrowthcurve2dendarcshape = new(models.StackRotatedGrowthCurve2DEndArcShape).Stage(stackrotatedgrowthcurve2dendarcshapeFormCallback.probe.stageOfInterest)
	}
	stackrotatedgrowthcurve2dendarcshape_ := stackrotatedgrowthcurve2dendarcshapeFormCallback.stackrotatedgrowthcurve2dendarcshape
	_ = stackrotatedgrowthcurve2dendarcshape_

	for _, formDiv := range stackrotatedgrowthcurve2dendarcshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dendarcshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dendarcshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dendarcshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dendarcshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dendarcshape_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dendarcshape_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dendarcshape_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dendarcshape_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dendarcshape_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dendarcshape_.RadiusY), formDiv)
		case "StackOfRotatedGrowthCurve2D:StackRotatedGrowthCurve2DEndArcShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the StackOfRotatedGrowthCurve2D instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target StackOfRotatedGrowthCurve2D instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.StackOfRotatedGrowthCurve2D](stackrotatedgrowthcurve2dendarcshapeFormCallback.probe.stageOfInterest)
			targetStackOfRotatedGrowthCurve2DIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetStackOfRotatedGrowthCurve2DIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all StackOfRotatedGrowthCurve2D instances and update their StackRotatedGrowthCurve2DEndArcShapes slice
			for _stackofrotatedgrowthcurve2d := range *models.GetGongstructInstancesSetFromPointerType[*models.StackOfRotatedGrowthCurve2D](stackrotatedgrowthcurve2dendarcshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(stackrotatedgrowthcurve2dendarcshapeFormCallback.probe.stageOfInterest, _stackofrotatedgrowthcurve2d)
				
				// if StackOfRotatedGrowthCurve2D is selected
				if targetStackOfRotatedGrowthCurve2DIDs[id] {
					// ensure stackrotatedgrowthcurve2dendarcshape_ is in _stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DEndArcShapes
					found := false
					for _, _b := range _stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DEndArcShapes {
						if _b == stackrotatedgrowthcurve2dendarcshape_ {
							found = true
							break
						}
					}
					if !found {
						_stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DEndArcShapes = append(_stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DEndArcShapes, stackrotatedgrowthcurve2dendarcshape_)
						stackrotatedgrowthcurve2dendarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_stackofrotatedgrowthcurve2d, "StackRotatedGrowthCurve2DEndArcShapes", &_stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DEndArcShapes)
					}
				} else {
					// ensure stackrotatedgrowthcurve2dendarcshape_ is NOT in _stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DEndArcShapes
					idx := slices.Index(_stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DEndArcShapes, stackrotatedgrowthcurve2dendarcshape_)
					if idx != -1 {
						_stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DEndArcShapes = slices.Delete(_stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DEndArcShapes, idx, idx+1)
						stackrotatedgrowthcurve2dendarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_stackofrotatedgrowthcurve2d, "StackRotatedGrowthCurve2DEndArcShapes", &_stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DEndArcShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if stackrotatedgrowthcurve2dendarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackrotatedgrowthcurve2dendarcshape_.Unstage(stackrotatedgrowthcurve2dendarcshapeFormCallback.probe.stageOfInterest)
	}

	stackrotatedgrowthcurve2dendarcshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StackRotatedGrowthCurve2DEndArcShape](
		stackrotatedgrowthcurve2dendarcshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if stackrotatedgrowthcurve2dendarcshapeFormCallback.CreationMode || stackrotatedgrowthcurve2dendarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackrotatedgrowthcurve2dendarcshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(stackrotatedgrowthcurve2dendarcshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StackRotatedGrowthCurve2DEndArcShapeFormCallback(
			nil,
			stackrotatedgrowthcurve2dendarcshapeFormCallback.probe,
			newFormGroup,
		)
		stackrotatedgrowthcurve2dendarcshape := new(models.StackRotatedGrowthCurve2DEndArcShape)
		FillUpForm(stackrotatedgrowthcurve2dendarcshape, newFormGroup, stackrotatedgrowthcurve2dendarcshapeFormCallback.probe)
		stackrotatedgrowthcurve2dendarcshapeFormCallback.probe.formStage.Commit()
	}

	stackrotatedgrowthcurve2dendarcshapeFormCallback.probe.ux_tree()
}
func __gong__New__StackRotatedGrowthCurve2DRibbonEndShapeFormCallback(
	stackrotatedgrowthcurve2dribbonendshape *models.StackRotatedGrowthCurve2DRibbonEndShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackrotatedgrowthcurve2dribbonendshapeFormCallback *StackRotatedGrowthCurve2DRibbonEndShapeFormCallback) {
	stackrotatedgrowthcurve2dribbonendshapeFormCallback = new(StackRotatedGrowthCurve2DRibbonEndShapeFormCallback)
	stackrotatedgrowthcurve2dribbonendshapeFormCallback.probe = probe
	stackrotatedgrowthcurve2dribbonendshapeFormCallback.stackrotatedgrowthcurve2dribbonendshape = stackrotatedgrowthcurve2dribbonendshape
	stackrotatedgrowthcurve2dribbonendshapeFormCallback.formGroup = formGroup

	stackrotatedgrowthcurve2dribbonendshapeFormCallback.CreationMode = (stackrotatedgrowthcurve2dribbonendshape == nil)

	return
}

type StackRotatedGrowthCurve2DRibbonEndShapeFormCallback struct {
	stackrotatedgrowthcurve2dribbonendshape *models.StackRotatedGrowthCurve2DRibbonEndShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (stackrotatedgrowthcurve2dribbonendshapeFormCallback *StackRotatedGrowthCurve2DRibbonEndShapeFormCallback) OnSave() {
	stackrotatedgrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest.Lock()
	defer stackrotatedgrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StackRotatedGrowthCurve2DRibbonEndShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	stackrotatedgrowthcurve2dribbonendshapeFormCallback.probe.formStage.Checkout()

	if stackrotatedgrowthcurve2dribbonendshapeFormCallback.stackrotatedgrowthcurve2dribbonendshape == nil {
		stackrotatedgrowthcurve2dribbonendshapeFormCallback.stackrotatedgrowthcurve2dribbonendshape = new(models.StackRotatedGrowthCurve2DRibbonEndShape).Stage(stackrotatedgrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest)
	}
	stackrotatedgrowthcurve2dribbonendshape_ := stackrotatedgrowthcurve2dribbonendshapeFormCallback.stackrotatedgrowthcurve2dribbonendshape
	_ = stackrotatedgrowthcurve2dribbonendshape_

	for _, formDiv := range stackrotatedgrowthcurve2dribbonendshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonendshape_.Name), formDiv)
		case "BottomStartX":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonendshape_.BottomStartX), formDiv)
		case "BottomStartY":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonendshape_.BottomStartY), formDiv)
		case "BottomEndX":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonendshape_.BottomEndX), formDiv)
		case "BottomEndY":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonendshape_.BottomEndY), formDiv)
		case "BottomRadiusX":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonendshape_.BottomRadiusX), formDiv)
		case "BottomRadiusY":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonendshape_.BottomRadiusY), formDiv)
		case "BottomXAxisRotation":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonendshape_.BottomXAxisRotation), formDiv)
		case "BottomLargeArcFlag":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonendshape_.BottomLargeArcFlag), formDiv)
		case "BottomSweepFlag":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonendshape_.BottomSweepFlag), formDiv)
		case "TopStartX":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonendshape_.TopStartX), formDiv)
		case "TopStartY":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonendshape_.TopStartY), formDiv)
		case "TopEndX":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonendshape_.TopEndX), formDiv)
		case "TopEndY":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonendshape_.TopEndY), formDiv)
		case "TopRadiusX":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonendshape_.TopRadiusX), formDiv)
		case "TopRadiusY":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonendshape_.TopRadiusY), formDiv)
		case "TopXAxisRotation":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonendshape_.TopXAxisRotation), formDiv)
		case "TopLargeArcFlag":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonendshape_.TopLargeArcFlag), formDiv)
		case "TopSweepFlag":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonendshape_.TopSweepFlag), formDiv)
		case "StackOfRotatedGrowthCurve2DRibbon:StackRotatedGrowthCurve2DRibbonEndShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the StackOfRotatedGrowthCurve2DRibbon instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target StackOfRotatedGrowthCurve2DRibbon instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.StackOfRotatedGrowthCurve2DRibbon](stackrotatedgrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest)
			targetStackOfRotatedGrowthCurve2DRibbonIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetStackOfRotatedGrowthCurve2DRibbonIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all StackOfRotatedGrowthCurve2DRibbon instances and update their StackRotatedGrowthCurve2DRibbonEndShapes slice
			for _stackofrotatedgrowthcurve2dribbon := range *models.GetGongstructInstancesSetFromPointerType[*models.StackOfRotatedGrowthCurve2DRibbon](stackrotatedgrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(stackrotatedgrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest, _stackofrotatedgrowthcurve2dribbon)
				
				// if StackOfRotatedGrowthCurve2DRibbon is selected
				if targetStackOfRotatedGrowthCurve2DRibbonIDs[id] {
					// ensure stackrotatedgrowthcurve2dribbonendshape_ is in _stackofrotatedgrowthcurve2dribbon.StackRotatedGrowthCurve2DRibbonEndShapes
					found := false
					for _, _b := range _stackofrotatedgrowthcurve2dribbon.StackRotatedGrowthCurve2DRibbonEndShapes {
						if _b == stackrotatedgrowthcurve2dribbonendshape_ {
							found = true
							break
						}
					}
					if !found {
						_stackofrotatedgrowthcurve2dribbon.StackRotatedGrowthCurve2DRibbonEndShapes = append(_stackofrotatedgrowthcurve2dribbon.StackRotatedGrowthCurve2DRibbonEndShapes, stackrotatedgrowthcurve2dribbonendshape_)
						stackrotatedgrowthcurve2dribbonendshapeFormCallback.probe.UpdateSliceOfPointersCallback(_stackofrotatedgrowthcurve2dribbon, "StackRotatedGrowthCurve2DRibbonEndShapes", &_stackofrotatedgrowthcurve2dribbon.StackRotatedGrowthCurve2DRibbonEndShapes)
					}
				} else {
					// ensure stackrotatedgrowthcurve2dribbonendshape_ is NOT in _stackofrotatedgrowthcurve2dribbon.StackRotatedGrowthCurve2DRibbonEndShapes
					idx := slices.Index(_stackofrotatedgrowthcurve2dribbon.StackRotatedGrowthCurve2DRibbonEndShapes, stackrotatedgrowthcurve2dribbonendshape_)
					if idx != -1 {
						_stackofrotatedgrowthcurve2dribbon.StackRotatedGrowthCurve2DRibbonEndShapes = slices.Delete(_stackofrotatedgrowthcurve2dribbon.StackRotatedGrowthCurve2DRibbonEndShapes, idx, idx+1)
						stackrotatedgrowthcurve2dribbonendshapeFormCallback.probe.UpdateSliceOfPointersCallback(_stackofrotatedgrowthcurve2dribbon, "StackRotatedGrowthCurve2DRibbonEndShapes", &_stackofrotatedgrowthcurve2dribbon.StackRotatedGrowthCurve2DRibbonEndShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if stackrotatedgrowthcurve2dribbonendshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackrotatedgrowthcurve2dribbonendshape_.Unstage(stackrotatedgrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest)
	}

	stackrotatedgrowthcurve2dribbonendshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StackRotatedGrowthCurve2DRibbonEndShape](
		stackrotatedgrowthcurve2dribbonendshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if stackrotatedgrowthcurve2dribbonendshapeFormCallback.CreationMode || stackrotatedgrowthcurve2dribbonendshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackrotatedgrowthcurve2dribbonendshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(stackrotatedgrowthcurve2dribbonendshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StackRotatedGrowthCurve2DRibbonEndShapeFormCallback(
			nil,
			stackrotatedgrowthcurve2dribbonendshapeFormCallback.probe,
			newFormGroup,
		)
		stackrotatedgrowthcurve2dribbonendshape := new(models.StackRotatedGrowthCurve2DRibbonEndShape)
		FillUpForm(stackrotatedgrowthcurve2dribbonendshape, newFormGroup, stackrotatedgrowthcurve2dribbonendshapeFormCallback.probe)
		stackrotatedgrowthcurve2dribbonendshapeFormCallback.probe.formStage.Commit()
	}

	stackrotatedgrowthcurve2dribbonendshapeFormCallback.probe.ux_tree()
}
func __gong__New__StackRotatedGrowthCurve2DRibbonStartShapeFormCallback(
	stackrotatedgrowthcurve2dribbonstartshape *models.StackRotatedGrowthCurve2DRibbonStartShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackrotatedgrowthcurve2dribbonstartshapeFormCallback *StackRotatedGrowthCurve2DRibbonStartShapeFormCallback) {
	stackrotatedgrowthcurve2dribbonstartshapeFormCallback = new(StackRotatedGrowthCurve2DRibbonStartShapeFormCallback)
	stackrotatedgrowthcurve2dribbonstartshapeFormCallback.probe = probe
	stackrotatedgrowthcurve2dribbonstartshapeFormCallback.stackrotatedgrowthcurve2dribbonstartshape = stackrotatedgrowthcurve2dribbonstartshape
	stackrotatedgrowthcurve2dribbonstartshapeFormCallback.formGroup = formGroup

	stackrotatedgrowthcurve2dribbonstartshapeFormCallback.CreationMode = (stackrotatedgrowthcurve2dribbonstartshape == nil)

	return
}

type StackRotatedGrowthCurve2DRibbonStartShapeFormCallback struct {
	stackrotatedgrowthcurve2dribbonstartshape *models.StackRotatedGrowthCurve2DRibbonStartShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (stackrotatedgrowthcurve2dribbonstartshapeFormCallback *StackRotatedGrowthCurve2DRibbonStartShapeFormCallback) OnSave() {
	stackrotatedgrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest.Lock()
	defer stackrotatedgrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StackRotatedGrowthCurve2DRibbonStartShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	stackrotatedgrowthcurve2dribbonstartshapeFormCallback.probe.formStage.Checkout()

	if stackrotatedgrowthcurve2dribbonstartshapeFormCallback.stackrotatedgrowthcurve2dribbonstartshape == nil {
		stackrotatedgrowthcurve2dribbonstartshapeFormCallback.stackrotatedgrowthcurve2dribbonstartshape = new(models.StackRotatedGrowthCurve2DRibbonStartShape).Stage(stackrotatedgrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest)
	}
	stackrotatedgrowthcurve2dribbonstartshape_ := stackrotatedgrowthcurve2dribbonstartshapeFormCallback.stackrotatedgrowthcurve2dribbonstartshape
	_ = stackrotatedgrowthcurve2dribbonstartshape_

	for _, formDiv := range stackrotatedgrowthcurve2dribbonstartshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonstartshape_.Name), formDiv)
		case "BottomStartX":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonstartshape_.BottomStartX), formDiv)
		case "BottomStartY":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonstartshape_.BottomStartY), formDiv)
		case "BottomEndX":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonstartshape_.BottomEndX), formDiv)
		case "BottomEndY":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonstartshape_.BottomEndY), formDiv)
		case "BottomRadiusX":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonstartshape_.BottomRadiusX), formDiv)
		case "BottomRadiusY":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonstartshape_.BottomRadiusY), formDiv)
		case "BottomXAxisRotation":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonstartshape_.BottomXAxisRotation), formDiv)
		case "BottomLargeArcFlag":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonstartshape_.BottomLargeArcFlag), formDiv)
		case "BottomSweepFlag":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonstartshape_.BottomSweepFlag), formDiv)
		case "TopStartX":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonstartshape_.TopStartX), formDiv)
		case "TopStartY":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonstartshape_.TopStartY), formDiv)
		case "TopEndX":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonstartshape_.TopEndX), formDiv)
		case "TopEndY":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonstartshape_.TopEndY), formDiv)
		case "TopRadiusX":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonstartshape_.TopRadiusX), formDiv)
		case "TopRadiusY":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonstartshape_.TopRadiusY), formDiv)
		case "TopXAxisRotation":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonstartshape_.TopXAxisRotation), formDiv)
		case "TopLargeArcFlag":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonstartshape_.TopLargeArcFlag), formDiv)
		case "TopSweepFlag":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dribbonstartshape_.TopSweepFlag), formDiv)
		case "StackOfRotatedGrowthCurve2DRibbon:StackRotatedGrowthCurve2DRibbonStartShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the StackOfRotatedGrowthCurve2DRibbon instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target StackOfRotatedGrowthCurve2DRibbon instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.StackOfRotatedGrowthCurve2DRibbon](stackrotatedgrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest)
			targetStackOfRotatedGrowthCurve2DRibbonIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetStackOfRotatedGrowthCurve2DRibbonIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all StackOfRotatedGrowthCurve2DRibbon instances and update their StackRotatedGrowthCurve2DRibbonStartShapes slice
			for _stackofrotatedgrowthcurve2dribbon := range *models.GetGongstructInstancesSetFromPointerType[*models.StackOfRotatedGrowthCurve2DRibbon](stackrotatedgrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(stackrotatedgrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest, _stackofrotatedgrowthcurve2dribbon)
				
				// if StackOfRotatedGrowthCurve2DRibbon is selected
				if targetStackOfRotatedGrowthCurve2DRibbonIDs[id] {
					// ensure stackrotatedgrowthcurve2dribbonstartshape_ is in _stackofrotatedgrowthcurve2dribbon.StackRotatedGrowthCurve2DRibbonStartShapes
					found := false
					for _, _b := range _stackofrotatedgrowthcurve2dribbon.StackRotatedGrowthCurve2DRibbonStartShapes {
						if _b == stackrotatedgrowthcurve2dribbonstartshape_ {
							found = true
							break
						}
					}
					if !found {
						_stackofrotatedgrowthcurve2dribbon.StackRotatedGrowthCurve2DRibbonStartShapes = append(_stackofrotatedgrowthcurve2dribbon.StackRotatedGrowthCurve2DRibbonStartShapes, stackrotatedgrowthcurve2dribbonstartshape_)
						stackrotatedgrowthcurve2dribbonstartshapeFormCallback.probe.UpdateSliceOfPointersCallback(_stackofrotatedgrowthcurve2dribbon, "StackRotatedGrowthCurve2DRibbonStartShapes", &_stackofrotatedgrowthcurve2dribbon.StackRotatedGrowthCurve2DRibbonStartShapes)
					}
				} else {
					// ensure stackrotatedgrowthcurve2dribbonstartshape_ is NOT in _stackofrotatedgrowthcurve2dribbon.StackRotatedGrowthCurve2DRibbonStartShapes
					idx := slices.Index(_stackofrotatedgrowthcurve2dribbon.StackRotatedGrowthCurve2DRibbonStartShapes, stackrotatedgrowthcurve2dribbonstartshape_)
					if idx != -1 {
						_stackofrotatedgrowthcurve2dribbon.StackRotatedGrowthCurve2DRibbonStartShapes = slices.Delete(_stackofrotatedgrowthcurve2dribbon.StackRotatedGrowthCurve2DRibbonStartShapes, idx, idx+1)
						stackrotatedgrowthcurve2dribbonstartshapeFormCallback.probe.UpdateSliceOfPointersCallback(_stackofrotatedgrowthcurve2dribbon, "StackRotatedGrowthCurve2DRibbonStartShapes", &_stackofrotatedgrowthcurve2dribbon.StackRotatedGrowthCurve2DRibbonStartShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if stackrotatedgrowthcurve2dribbonstartshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackrotatedgrowthcurve2dribbonstartshape_.Unstage(stackrotatedgrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest)
	}

	stackrotatedgrowthcurve2dribbonstartshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StackRotatedGrowthCurve2DRibbonStartShape](
		stackrotatedgrowthcurve2dribbonstartshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if stackrotatedgrowthcurve2dribbonstartshapeFormCallback.CreationMode || stackrotatedgrowthcurve2dribbonstartshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackrotatedgrowthcurve2dribbonstartshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(stackrotatedgrowthcurve2dribbonstartshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StackRotatedGrowthCurve2DRibbonStartShapeFormCallback(
			nil,
			stackrotatedgrowthcurve2dribbonstartshapeFormCallback.probe,
			newFormGroup,
		)
		stackrotatedgrowthcurve2dribbonstartshape := new(models.StackRotatedGrowthCurve2DRibbonStartShape)
		FillUpForm(stackrotatedgrowthcurve2dribbonstartshape, newFormGroup, stackrotatedgrowthcurve2dribbonstartshapeFormCallback.probe)
		stackrotatedgrowthcurve2dribbonstartshapeFormCallback.probe.formStage.Commit()
	}

	stackrotatedgrowthcurve2dribbonstartshapeFormCallback.probe.ux_tree()
}
func __gong__New__StackRotatedGrowthCurve2DStartArcShapeFormCallback(
	stackrotatedgrowthcurve2dstartarcshape *models.StackRotatedGrowthCurve2DStartArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackrotatedgrowthcurve2dstartarcshapeFormCallback *StackRotatedGrowthCurve2DStartArcShapeFormCallback) {
	stackrotatedgrowthcurve2dstartarcshapeFormCallback = new(StackRotatedGrowthCurve2DStartArcShapeFormCallback)
	stackrotatedgrowthcurve2dstartarcshapeFormCallback.probe = probe
	stackrotatedgrowthcurve2dstartarcshapeFormCallback.stackrotatedgrowthcurve2dstartarcshape = stackrotatedgrowthcurve2dstartarcshape
	stackrotatedgrowthcurve2dstartarcshapeFormCallback.formGroup = formGroup

	stackrotatedgrowthcurve2dstartarcshapeFormCallback.CreationMode = (stackrotatedgrowthcurve2dstartarcshape == nil)

	return
}

type StackRotatedGrowthCurve2DStartArcShapeFormCallback struct {
	stackrotatedgrowthcurve2dstartarcshape *models.StackRotatedGrowthCurve2DStartArcShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (stackrotatedgrowthcurve2dstartarcshapeFormCallback *StackRotatedGrowthCurve2DStartArcShapeFormCallback) OnSave() {
	stackrotatedgrowthcurve2dstartarcshapeFormCallback.probe.stageOfInterest.Lock()
	defer stackrotatedgrowthcurve2dstartarcshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StackRotatedGrowthCurve2DStartArcShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	stackrotatedgrowthcurve2dstartarcshapeFormCallback.probe.formStage.Checkout()

	if stackrotatedgrowthcurve2dstartarcshapeFormCallback.stackrotatedgrowthcurve2dstartarcshape == nil {
		stackrotatedgrowthcurve2dstartarcshapeFormCallback.stackrotatedgrowthcurve2dstartarcshape = new(models.StackRotatedGrowthCurve2DStartArcShape).Stage(stackrotatedgrowthcurve2dstartarcshapeFormCallback.probe.stageOfInterest)
	}
	stackrotatedgrowthcurve2dstartarcshape_ := stackrotatedgrowthcurve2dstartarcshapeFormCallback.stackrotatedgrowthcurve2dstartarcshape
	_ = stackrotatedgrowthcurve2dstartarcshape_

	for _, formDiv := range stackrotatedgrowthcurve2dstartarcshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dstartarcshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dstartarcshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dstartarcshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dstartarcshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dstartarcshape_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dstartarcshape_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dstartarcshape_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dstartarcshape_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dstartarcshape_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(stackrotatedgrowthcurve2dstartarcshape_.RadiusY), formDiv)
		case "StackOfRotatedGrowthCurve2D:StackRotatedGrowthCurve2DStartArcShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the StackOfRotatedGrowthCurve2D instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target StackOfRotatedGrowthCurve2D instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.StackOfRotatedGrowthCurve2D](stackrotatedgrowthcurve2dstartarcshapeFormCallback.probe.stageOfInterest)
			targetStackOfRotatedGrowthCurve2DIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetStackOfRotatedGrowthCurve2DIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all StackOfRotatedGrowthCurve2D instances and update their StackRotatedGrowthCurve2DStartArcShapes slice
			for _stackofrotatedgrowthcurve2d := range *models.GetGongstructInstancesSetFromPointerType[*models.StackOfRotatedGrowthCurve2D](stackrotatedgrowthcurve2dstartarcshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(stackrotatedgrowthcurve2dstartarcshapeFormCallback.probe.stageOfInterest, _stackofrotatedgrowthcurve2d)
				
				// if StackOfRotatedGrowthCurve2D is selected
				if targetStackOfRotatedGrowthCurve2DIDs[id] {
					// ensure stackrotatedgrowthcurve2dstartarcshape_ is in _stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DStartArcShapes
					found := false
					for _, _b := range _stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DStartArcShapes {
						if _b == stackrotatedgrowthcurve2dstartarcshape_ {
							found = true
							break
						}
					}
					if !found {
						_stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DStartArcShapes = append(_stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DStartArcShapes, stackrotatedgrowthcurve2dstartarcshape_)
						stackrotatedgrowthcurve2dstartarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_stackofrotatedgrowthcurve2d, "StackRotatedGrowthCurve2DStartArcShapes", &_stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DStartArcShapes)
					}
				} else {
					// ensure stackrotatedgrowthcurve2dstartarcshape_ is NOT in _stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DStartArcShapes
					idx := slices.Index(_stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DStartArcShapes, stackrotatedgrowthcurve2dstartarcshape_)
					if idx != -1 {
						_stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DStartArcShapes = slices.Delete(_stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DStartArcShapes, idx, idx+1)
						stackrotatedgrowthcurve2dstartarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_stackofrotatedgrowthcurve2d, "StackRotatedGrowthCurve2DStartArcShapes", &_stackofrotatedgrowthcurve2d.StackRotatedGrowthCurve2DStartArcShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if stackrotatedgrowthcurve2dstartarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackrotatedgrowthcurve2dstartarcshape_.Unstage(stackrotatedgrowthcurve2dstartarcshapeFormCallback.probe.stageOfInterest)
	}

	stackrotatedgrowthcurve2dstartarcshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StackRotatedGrowthCurve2DStartArcShape](
		stackrotatedgrowthcurve2dstartarcshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if stackrotatedgrowthcurve2dstartarcshapeFormCallback.CreationMode || stackrotatedgrowthcurve2dstartarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stackrotatedgrowthcurve2dstartarcshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(stackrotatedgrowthcurve2dstartarcshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StackRotatedGrowthCurve2DStartArcShapeFormCallback(
			nil,
			stackrotatedgrowthcurve2dstartarcshapeFormCallback.probe,
			newFormGroup,
		)
		stackrotatedgrowthcurve2dstartarcshape := new(models.StackRotatedGrowthCurve2DStartArcShape)
		FillUpForm(stackrotatedgrowthcurve2dstartarcshape, newFormGroup, stackrotatedgrowthcurve2dstartarcshapeFormCallback.probe)
		stackrotatedgrowthcurve2dstartarcshapeFormCallback.probe.formStage.Commit()
	}

	stackrotatedgrowthcurve2dstartarcshapeFormCallback.probe.ux_tree()
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
func __gong__New__StartHalfwayArcShapeFormCallback(
	starthalfwayarcshape *models.StartHalfwayArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (starthalfwayarcshapeFormCallback *StartHalfwayArcShapeFormCallback) {
	starthalfwayarcshapeFormCallback = new(StartHalfwayArcShapeFormCallback)
	starthalfwayarcshapeFormCallback.probe = probe
	starthalfwayarcshapeFormCallback.starthalfwayarcshape = starthalfwayarcshape
	starthalfwayarcshapeFormCallback.formGroup = formGroup

	starthalfwayarcshapeFormCallback.CreationMode = (starthalfwayarcshape == nil)

	return
}

type StartHalfwayArcShapeFormCallback struct {
	starthalfwayarcshape *models.StartHalfwayArcShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (starthalfwayarcshapeFormCallback *StartHalfwayArcShapeFormCallback) OnSave() {
	starthalfwayarcshapeFormCallback.probe.stageOfInterest.Lock()
	defer starthalfwayarcshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StartHalfwayArcShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	starthalfwayarcshapeFormCallback.probe.formStage.Checkout()

	if starthalfwayarcshapeFormCallback.starthalfwayarcshape == nil {
		starthalfwayarcshapeFormCallback.starthalfwayarcshape = new(models.StartHalfwayArcShape).Stage(starthalfwayarcshapeFormCallback.probe.stageOfInterest)
	}
	starthalfwayarcshape_ := starthalfwayarcshapeFormCallback.starthalfwayarcshape
	_ = starthalfwayarcshape_

	for _, formDiv := range starthalfwayarcshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(starthalfwayarcshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(starthalfwayarcshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(starthalfwayarcshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(starthalfwayarcshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(starthalfwayarcshape_.EndY), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(starthalfwayarcshape_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(starthalfwayarcshape_.RadiusY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(starthalfwayarcshape_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(starthalfwayarcshape_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(starthalfwayarcshape_.SweepFlag), formDiv)
		case "StartHalfwayArcShapeGrid:StartHalfwayArcShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the StartHalfwayArcShapeGrid instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target StartHalfwayArcShapeGrid instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.StartHalfwayArcShapeGrid](starthalfwayarcshapeFormCallback.probe.stageOfInterest)
			targetStartHalfwayArcShapeGridIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetStartHalfwayArcShapeGridIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all StartHalfwayArcShapeGrid instances and update their StartHalfwayArcShapes slice
			for _starthalfwayarcshapegrid := range *models.GetGongstructInstancesSetFromPointerType[*models.StartHalfwayArcShapeGrid](starthalfwayarcshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(starthalfwayarcshapeFormCallback.probe.stageOfInterest, _starthalfwayarcshapegrid)
				
				// if StartHalfwayArcShapeGrid is selected
				if targetStartHalfwayArcShapeGridIDs[id] {
					// ensure starthalfwayarcshape_ is in _starthalfwayarcshapegrid.StartHalfwayArcShapes
					found := false
					for _, _b := range _starthalfwayarcshapegrid.StartHalfwayArcShapes {
						if _b == starthalfwayarcshape_ {
							found = true
							break
						}
					}
					if !found {
						_starthalfwayarcshapegrid.StartHalfwayArcShapes = append(_starthalfwayarcshapegrid.StartHalfwayArcShapes, starthalfwayarcshape_)
						starthalfwayarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_starthalfwayarcshapegrid, "StartHalfwayArcShapes", &_starthalfwayarcshapegrid.StartHalfwayArcShapes)
					}
				} else {
					// ensure starthalfwayarcshape_ is NOT in _starthalfwayarcshapegrid.StartHalfwayArcShapes
					idx := slices.Index(_starthalfwayarcshapegrid.StartHalfwayArcShapes, starthalfwayarcshape_)
					if idx != -1 {
						_starthalfwayarcshapegrid.StartHalfwayArcShapes = slices.Delete(_starthalfwayarcshapegrid.StartHalfwayArcShapes, idx, idx+1)
						starthalfwayarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_starthalfwayarcshapegrid, "StartHalfwayArcShapes", &_starthalfwayarcshapegrid.StartHalfwayArcShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if starthalfwayarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		starthalfwayarcshape_.Unstage(starthalfwayarcshapeFormCallback.probe.stageOfInterest)
	}

	starthalfwayarcshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StartHalfwayArcShape](
		starthalfwayarcshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if starthalfwayarcshapeFormCallback.CreationMode || starthalfwayarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		starthalfwayarcshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(starthalfwayarcshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StartHalfwayArcShapeFormCallback(
			nil,
			starthalfwayarcshapeFormCallback.probe,
			newFormGroup,
		)
		starthalfwayarcshape := new(models.StartHalfwayArcShape)
		FillUpForm(starthalfwayarcshape, newFormGroup, starthalfwayarcshapeFormCallback.probe)
		starthalfwayarcshapeFormCallback.probe.formStage.Commit()
	}

	starthalfwayarcshapeFormCallback.probe.ux_tree()
}
func __gong__New__StartHalfwayArcShapeGridFormCallback(
	starthalfwayarcshapegrid *models.StartHalfwayArcShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (starthalfwayarcshapegridFormCallback *StartHalfwayArcShapeGridFormCallback) {
	starthalfwayarcshapegridFormCallback = new(StartHalfwayArcShapeGridFormCallback)
	starthalfwayarcshapegridFormCallback.probe = probe
	starthalfwayarcshapegridFormCallback.starthalfwayarcshapegrid = starthalfwayarcshapegrid
	starthalfwayarcshapegridFormCallback.formGroup = formGroup

	starthalfwayarcshapegridFormCallback.CreationMode = (starthalfwayarcshapegrid == nil)

	return
}

type StartHalfwayArcShapeGridFormCallback struct {
	starthalfwayarcshapegrid *models.StartHalfwayArcShapeGrid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (starthalfwayarcshapegridFormCallback *StartHalfwayArcShapeGridFormCallback) OnSave() {
	starthalfwayarcshapegridFormCallback.probe.stageOfInterest.Lock()
	defer starthalfwayarcshapegridFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StartHalfwayArcShapeGridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	starthalfwayarcshapegridFormCallback.probe.formStage.Checkout()

	if starthalfwayarcshapegridFormCallback.starthalfwayarcshapegrid == nil {
		starthalfwayarcshapegridFormCallback.starthalfwayarcshapegrid = new(models.StartHalfwayArcShapeGrid).Stage(starthalfwayarcshapegridFormCallback.probe.stageOfInterest)
	}
	starthalfwayarcshapegrid_ := starthalfwayarcshapegridFormCallback.starthalfwayarcshapegrid
	_ = starthalfwayarcshapegrid_

	for _, formDiv := range starthalfwayarcshapegridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(starthalfwayarcshapegrid_.Name), formDiv)
		case "StartHalfwayArcShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.StartHalfwayArcShape](starthalfwayarcshapegridFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.StartHalfwayArcShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.StartHalfwayArcShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					starthalfwayarcshapegridFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.StartHalfwayArcShape](starthalfwayarcshapegridFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			starthalfwayarcshapegrid_.StartHalfwayArcShapes = instanceSlice
			starthalfwayarcshapegridFormCallback.probe.UpdateSliceOfPointersCallback(starthalfwayarcshapegrid_, "StartHalfwayArcShapes", &starthalfwayarcshapegrid_.StartHalfwayArcShapes)

		}
	}

	// manage the suppress operation
	if starthalfwayarcshapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		starthalfwayarcshapegrid_.Unstage(starthalfwayarcshapegridFormCallback.probe.stageOfInterest)
	}

	starthalfwayarcshapegridFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StartHalfwayArcShapeGrid](
		starthalfwayarcshapegridFormCallback.probe,
	)

	// display a new form by reset the form stage
	if starthalfwayarcshapegridFormCallback.CreationMode || starthalfwayarcshapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		starthalfwayarcshapegridFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(starthalfwayarcshapegridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StartHalfwayArcShapeGridFormCallback(
			nil,
			starthalfwayarcshapegridFormCallback.probe,
			newFormGroup,
		)
		starthalfwayarcshapegrid := new(models.StartHalfwayArcShapeGrid)
		FillUpForm(starthalfwayarcshapegrid, newFormGroup, starthalfwayarcshapegridFormCallback.probe)
		starthalfwayarcshapegridFormCallback.probe.formStage.Commit()
	}

	starthalfwayarcshapegridFormCallback.probe.ux_tree()
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
func __gong__New__TopEndHalfwayArcShapeFormCallback(
	topendhalfwayarcshape *models.TopEndHalfwayArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (topendhalfwayarcshapeFormCallback *TopEndHalfwayArcShapeFormCallback) {
	topendhalfwayarcshapeFormCallback = new(TopEndHalfwayArcShapeFormCallback)
	topendhalfwayarcshapeFormCallback.probe = probe
	topendhalfwayarcshapeFormCallback.topendhalfwayarcshape = topendhalfwayarcshape
	topendhalfwayarcshapeFormCallback.formGroup = formGroup

	topendhalfwayarcshapeFormCallback.CreationMode = (topendhalfwayarcshape == nil)

	return
}

type TopEndHalfwayArcShapeFormCallback struct {
	topendhalfwayarcshape *models.TopEndHalfwayArcShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (topendhalfwayarcshapeFormCallback *TopEndHalfwayArcShapeFormCallback) OnSave() {
	topendhalfwayarcshapeFormCallback.probe.stageOfInterest.Lock()
	defer topendhalfwayarcshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TopEndHalfwayArcShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	topendhalfwayarcshapeFormCallback.probe.formStage.Checkout()

	if topendhalfwayarcshapeFormCallback.topendhalfwayarcshape == nil {
		topendhalfwayarcshapeFormCallback.topendhalfwayarcshape = new(models.TopEndHalfwayArcShape).Stage(topendhalfwayarcshapeFormCallback.probe.stageOfInterest)
	}
	topendhalfwayarcshape_ := topendhalfwayarcshapeFormCallback.topendhalfwayarcshape
	_ = topendhalfwayarcshape_

	for _, formDiv := range topendhalfwayarcshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(topendhalfwayarcshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(topendhalfwayarcshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(topendhalfwayarcshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(topendhalfwayarcshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(topendhalfwayarcshape_.EndY), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(topendhalfwayarcshape_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(topendhalfwayarcshape_.RadiusY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(topendhalfwayarcshape_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(topendhalfwayarcshape_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(topendhalfwayarcshape_.SweepFlag), formDiv)
		case "TopEndHalfwayArcShapeGrid:TopEndHalfwayArcShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the TopEndHalfwayArcShapeGrid instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target TopEndHalfwayArcShapeGrid instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.TopEndHalfwayArcShapeGrid](topendhalfwayarcshapeFormCallback.probe.stageOfInterest)
			targetTopEndHalfwayArcShapeGridIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetTopEndHalfwayArcShapeGridIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all TopEndHalfwayArcShapeGrid instances and update their TopEndHalfwayArcShapes slice
			for _topendhalfwayarcshapegrid := range *models.GetGongstructInstancesSetFromPointerType[*models.TopEndHalfwayArcShapeGrid](topendhalfwayarcshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(topendhalfwayarcshapeFormCallback.probe.stageOfInterest, _topendhalfwayarcshapegrid)
				
				// if TopEndHalfwayArcShapeGrid is selected
				if targetTopEndHalfwayArcShapeGridIDs[id] {
					// ensure topendhalfwayarcshape_ is in _topendhalfwayarcshapegrid.TopEndHalfwayArcShapes
					found := false
					for _, _b := range _topendhalfwayarcshapegrid.TopEndHalfwayArcShapes {
						if _b == topendhalfwayarcshape_ {
							found = true
							break
						}
					}
					if !found {
						_topendhalfwayarcshapegrid.TopEndHalfwayArcShapes = append(_topendhalfwayarcshapegrid.TopEndHalfwayArcShapes, topendhalfwayarcshape_)
						topendhalfwayarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_topendhalfwayarcshapegrid, "TopEndHalfwayArcShapes", &_topendhalfwayarcshapegrid.TopEndHalfwayArcShapes)
					}
				} else {
					// ensure topendhalfwayarcshape_ is NOT in _topendhalfwayarcshapegrid.TopEndHalfwayArcShapes
					idx := slices.Index(_topendhalfwayarcshapegrid.TopEndHalfwayArcShapes, topendhalfwayarcshape_)
					if idx != -1 {
						_topendhalfwayarcshapegrid.TopEndHalfwayArcShapes = slices.Delete(_topendhalfwayarcshapegrid.TopEndHalfwayArcShapes, idx, idx+1)
						topendhalfwayarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_topendhalfwayarcshapegrid, "TopEndHalfwayArcShapes", &_topendhalfwayarcshapegrid.TopEndHalfwayArcShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if topendhalfwayarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topendhalfwayarcshape_.Unstage(topendhalfwayarcshapeFormCallback.probe.stageOfInterest)
	}

	topendhalfwayarcshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TopEndHalfwayArcShape](
		topendhalfwayarcshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if topendhalfwayarcshapeFormCallback.CreationMode || topendhalfwayarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topendhalfwayarcshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(topendhalfwayarcshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TopEndHalfwayArcShapeFormCallback(
			nil,
			topendhalfwayarcshapeFormCallback.probe,
			newFormGroup,
		)
		topendhalfwayarcshape := new(models.TopEndHalfwayArcShape)
		FillUpForm(topendhalfwayarcshape, newFormGroup, topendhalfwayarcshapeFormCallback.probe)
		topendhalfwayarcshapeFormCallback.probe.formStage.Commit()
	}

	topendhalfwayarcshapeFormCallback.probe.ux_tree()
}
func __gong__New__TopEndHalfwayArcShapeGridFormCallback(
	topendhalfwayarcshapegrid *models.TopEndHalfwayArcShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (topendhalfwayarcshapegridFormCallback *TopEndHalfwayArcShapeGridFormCallback) {
	topendhalfwayarcshapegridFormCallback = new(TopEndHalfwayArcShapeGridFormCallback)
	topendhalfwayarcshapegridFormCallback.probe = probe
	topendhalfwayarcshapegridFormCallback.topendhalfwayarcshapegrid = topendhalfwayarcshapegrid
	topendhalfwayarcshapegridFormCallback.formGroup = formGroup

	topendhalfwayarcshapegridFormCallback.CreationMode = (topendhalfwayarcshapegrid == nil)

	return
}

type TopEndHalfwayArcShapeGridFormCallback struct {
	topendhalfwayarcshapegrid *models.TopEndHalfwayArcShapeGrid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (topendhalfwayarcshapegridFormCallback *TopEndHalfwayArcShapeGridFormCallback) OnSave() {
	topendhalfwayarcshapegridFormCallback.probe.stageOfInterest.Lock()
	defer topendhalfwayarcshapegridFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TopEndHalfwayArcShapeGridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	topendhalfwayarcshapegridFormCallback.probe.formStage.Checkout()

	if topendhalfwayarcshapegridFormCallback.topendhalfwayarcshapegrid == nil {
		topendhalfwayarcshapegridFormCallback.topendhalfwayarcshapegrid = new(models.TopEndHalfwayArcShapeGrid).Stage(topendhalfwayarcshapegridFormCallback.probe.stageOfInterest)
	}
	topendhalfwayarcshapegrid_ := topendhalfwayarcshapegridFormCallback.topendhalfwayarcshapegrid
	_ = topendhalfwayarcshapegrid_

	for _, formDiv := range topendhalfwayarcshapegridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(topendhalfwayarcshapegrid_.Name), formDiv)
		case "TopEndHalfwayArcShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.TopEndHalfwayArcShape](topendhalfwayarcshapegridFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.TopEndHalfwayArcShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.TopEndHalfwayArcShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					topendhalfwayarcshapegridFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.TopEndHalfwayArcShape](topendhalfwayarcshapegridFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			topendhalfwayarcshapegrid_.TopEndHalfwayArcShapes = instanceSlice
			topendhalfwayarcshapegridFormCallback.probe.UpdateSliceOfPointersCallback(topendhalfwayarcshapegrid_, "TopEndHalfwayArcShapes", &topendhalfwayarcshapegrid_.TopEndHalfwayArcShapes)

		}
	}

	// manage the suppress operation
	if topendhalfwayarcshapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topendhalfwayarcshapegrid_.Unstage(topendhalfwayarcshapegridFormCallback.probe.stageOfInterest)
	}

	topendhalfwayarcshapegridFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TopEndHalfwayArcShapeGrid](
		topendhalfwayarcshapegridFormCallback.probe,
	)

	// display a new form by reset the form stage
	if topendhalfwayarcshapegridFormCallback.CreationMode || topendhalfwayarcshapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topendhalfwayarcshapegridFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(topendhalfwayarcshapegridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TopEndHalfwayArcShapeGridFormCallback(
			nil,
			topendhalfwayarcshapegridFormCallback.probe,
			newFormGroup,
		)
		topendhalfwayarcshapegrid := new(models.TopEndHalfwayArcShapeGrid)
		FillUpForm(topendhalfwayarcshapegrid, newFormGroup, topendhalfwayarcshapegridFormCallback.probe)
		topendhalfwayarcshapegridFormCallback.probe.formStage.Commit()
	}

	topendhalfwayarcshapegridFormCallback.probe.ux_tree()
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
		case "TopStartHalfwayArcShapeGrid":
			FormDivSelectFieldToField(&(topgrowthcurve2d_.TopStartHalfwayArcShapeGrid), topgrowthcurve2dFormCallback.probe.stageOfInterest, formDiv)
		case "TopEndHalfwayArcShapeGrid":
			FormDivSelectFieldToField(&(topgrowthcurve2d_.TopEndHalfwayArcShapeGrid), topgrowthcurve2dFormCallback.probe.stageOfInterest, formDiv)
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
func __gong__New__TopMidArcVectorShapeFormCallback(
	topmidarcvectorshape *models.TopMidArcVectorShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (topmidarcvectorshapeFormCallback *TopMidArcVectorShapeFormCallback) {
	topmidarcvectorshapeFormCallback = new(TopMidArcVectorShapeFormCallback)
	topmidarcvectorshapeFormCallback.probe = probe
	topmidarcvectorshapeFormCallback.topmidarcvectorshape = topmidarcvectorshape
	topmidarcvectorshapeFormCallback.formGroup = formGroup

	topmidarcvectorshapeFormCallback.CreationMode = (topmidarcvectorshape == nil)

	return
}

type TopMidArcVectorShapeFormCallback struct {
	topmidarcvectorshape *models.TopMidArcVectorShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (topmidarcvectorshapeFormCallback *TopMidArcVectorShapeFormCallback) OnSave() {
	topmidarcvectorshapeFormCallback.probe.stageOfInterest.Lock()
	defer topmidarcvectorshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TopMidArcVectorShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	topmidarcvectorshapeFormCallback.probe.formStage.Checkout()

	if topmidarcvectorshapeFormCallback.topmidarcvectorshape == nil {
		topmidarcvectorshapeFormCallback.topmidarcvectorshape = new(models.TopMidArcVectorShape).Stage(topmidarcvectorshapeFormCallback.probe.stageOfInterest)
	}
	topmidarcvectorshape_ := topmidarcvectorshapeFormCallback.topmidarcvectorshape
	_ = topmidarcvectorshape_

	for _, formDiv := range topmidarcvectorshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(topmidarcvectorshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(topmidarcvectorshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(topmidarcvectorshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(topmidarcvectorshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(topmidarcvectorshape_.EndY), formDiv)
		case "TopMidArcVectorShapeGrid:TopMidArcVectorShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the TopMidArcVectorShapeGrid instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target TopMidArcVectorShapeGrid instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.TopMidArcVectorShapeGrid](topmidarcvectorshapeFormCallback.probe.stageOfInterest)
			targetTopMidArcVectorShapeGridIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetTopMidArcVectorShapeGridIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all TopMidArcVectorShapeGrid instances and update their TopMidArcVectorShapes slice
			for _topmidarcvectorshapegrid := range *models.GetGongstructInstancesSetFromPointerType[*models.TopMidArcVectorShapeGrid](topmidarcvectorshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(topmidarcvectorshapeFormCallback.probe.stageOfInterest, _topmidarcvectorshapegrid)
				
				// if TopMidArcVectorShapeGrid is selected
				if targetTopMidArcVectorShapeGridIDs[id] {
					// ensure topmidarcvectorshape_ is in _topmidarcvectorshapegrid.TopMidArcVectorShapes
					found := false
					for _, _b := range _topmidarcvectorshapegrid.TopMidArcVectorShapes {
						if _b == topmidarcvectorshape_ {
							found = true
							break
						}
					}
					if !found {
						_topmidarcvectorshapegrid.TopMidArcVectorShapes = append(_topmidarcvectorshapegrid.TopMidArcVectorShapes, topmidarcvectorshape_)
						topmidarcvectorshapeFormCallback.probe.UpdateSliceOfPointersCallback(_topmidarcvectorshapegrid, "TopMidArcVectorShapes", &_topmidarcvectorshapegrid.TopMidArcVectorShapes)
					}
				} else {
					// ensure topmidarcvectorshape_ is NOT in _topmidarcvectorshapegrid.TopMidArcVectorShapes
					idx := slices.Index(_topmidarcvectorshapegrid.TopMidArcVectorShapes, topmidarcvectorshape_)
					if idx != -1 {
						_topmidarcvectorshapegrid.TopMidArcVectorShapes = slices.Delete(_topmidarcvectorshapegrid.TopMidArcVectorShapes, idx, idx+1)
						topmidarcvectorshapeFormCallback.probe.UpdateSliceOfPointersCallback(_topmidarcvectorshapegrid, "TopMidArcVectorShapes", &_topmidarcvectorshapegrid.TopMidArcVectorShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if topmidarcvectorshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topmidarcvectorshape_.Unstage(topmidarcvectorshapeFormCallback.probe.stageOfInterest)
	}

	topmidarcvectorshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TopMidArcVectorShape](
		topmidarcvectorshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if topmidarcvectorshapeFormCallback.CreationMode || topmidarcvectorshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topmidarcvectorshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(topmidarcvectorshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TopMidArcVectorShapeFormCallback(
			nil,
			topmidarcvectorshapeFormCallback.probe,
			newFormGroup,
		)
		topmidarcvectorshape := new(models.TopMidArcVectorShape)
		FillUpForm(topmidarcvectorshape, newFormGroup, topmidarcvectorshapeFormCallback.probe)
		topmidarcvectorshapeFormCallback.probe.formStage.Commit()
	}

	topmidarcvectorshapeFormCallback.probe.ux_tree()
}
func __gong__New__TopMidArcVectorShapeGridFormCallback(
	topmidarcvectorshapegrid *models.TopMidArcVectorShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (topmidarcvectorshapegridFormCallback *TopMidArcVectorShapeGridFormCallback) {
	topmidarcvectorshapegridFormCallback = new(TopMidArcVectorShapeGridFormCallback)
	topmidarcvectorshapegridFormCallback.probe = probe
	topmidarcvectorshapegridFormCallback.topmidarcvectorshapegrid = topmidarcvectorshapegrid
	topmidarcvectorshapegridFormCallback.formGroup = formGroup

	topmidarcvectorshapegridFormCallback.CreationMode = (topmidarcvectorshapegrid == nil)

	return
}

type TopMidArcVectorShapeGridFormCallback struct {
	topmidarcvectorshapegrid *models.TopMidArcVectorShapeGrid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (topmidarcvectorshapegridFormCallback *TopMidArcVectorShapeGridFormCallback) OnSave() {
	topmidarcvectorshapegridFormCallback.probe.stageOfInterest.Lock()
	defer topmidarcvectorshapegridFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TopMidArcVectorShapeGridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	topmidarcvectorshapegridFormCallback.probe.formStage.Checkout()

	if topmidarcvectorshapegridFormCallback.topmidarcvectorshapegrid == nil {
		topmidarcvectorshapegridFormCallback.topmidarcvectorshapegrid = new(models.TopMidArcVectorShapeGrid).Stage(topmidarcvectorshapegridFormCallback.probe.stageOfInterest)
	}
	topmidarcvectorshapegrid_ := topmidarcvectorshapegridFormCallback.topmidarcvectorshapegrid
	_ = topmidarcvectorshapegrid_

	for _, formDiv := range topmidarcvectorshapegridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(topmidarcvectorshapegrid_.Name), formDiv)
		case "TopMidArcVectorShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.TopMidArcVectorShape](topmidarcvectorshapegridFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.TopMidArcVectorShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.TopMidArcVectorShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					topmidarcvectorshapegridFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.TopMidArcVectorShape](topmidarcvectorshapegridFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			topmidarcvectorshapegrid_.TopMidArcVectorShapes = instanceSlice
			topmidarcvectorshapegridFormCallback.probe.UpdateSliceOfPointersCallback(topmidarcvectorshapegrid_, "TopMidArcVectorShapes", &topmidarcvectorshapegrid_.TopMidArcVectorShapes)

		}
	}

	// manage the suppress operation
	if topmidarcvectorshapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topmidarcvectorshapegrid_.Unstage(topmidarcvectorshapegridFormCallback.probe.stageOfInterest)
	}

	topmidarcvectorshapegridFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TopMidArcVectorShapeGrid](
		topmidarcvectorshapegridFormCallback.probe,
	)

	// display a new form by reset the form stage
	if topmidarcvectorshapegridFormCallback.CreationMode || topmidarcvectorshapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topmidarcvectorshapegridFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(topmidarcvectorshapegridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TopMidArcVectorShapeGridFormCallback(
			nil,
			topmidarcvectorshapegridFormCallback.probe,
			newFormGroup,
		)
		topmidarcvectorshapegrid := new(models.TopMidArcVectorShapeGrid)
		FillUpForm(topmidarcvectorshapegrid, newFormGroup, topmidarcvectorshapegridFormCallback.probe)
		topmidarcvectorshapegridFormCallback.probe.formStage.Commit()
	}

	topmidarcvectorshapegridFormCallback.probe.ux_tree()
}
func __gong__New__TopStackGrowthCurve2DEndHalfwayArcShapeFormCallback(
	topstackgrowthcurve2dendhalfwayarcshape *models.TopStackGrowthCurve2DEndHalfwayArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (topstackgrowthcurve2dendhalfwayarcshapeFormCallback *TopStackGrowthCurve2DEndHalfwayArcShapeFormCallback) {
	topstackgrowthcurve2dendhalfwayarcshapeFormCallback = new(TopStackGrowthCurve2DEndHalfwayArcShapeFormCallback)
	topstackgrowthcurve2dendhalfwayarcshapeFormCallback.probe = probe
	topstackgrowthcurve2dendhalfwayarcshapeFormCallback.topstackgrowthcurve2dendhalfwayarcshape = topstackgrowthcurve2dendhalfwayarcshape
	topstackgrowthcurve2dendhalfwayarcshapeFormCallback.formGroup = formGroup

	topstackgrowthcurve2dendhalfwayarcshapeFormCallback.CreationMode = (topstackgrowthcurve2dendhalfwayarcshape == nil)

	return
}

type TopStackGrowthCurve2DEndHalfwayArcShapeFormCallback struct {
	topstackgrowthcurve2dendhalfwayarcshape *models.TopStackGrowthCurve2DEndHalfwayArcShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (topstackgrowthcurve2dendhalfwayarcshapeFormCallback *TopStackGrowthCurve2DEndHalfwayArcShapeFormCallback) OnSave() {
	topstackgrowthcurve2dendhalfwayarcshapeFormCallback.probe.stageOfInterest.Lock()
	defer topstackgrowthcurve2dendhalfwayarcshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TopStackGrowthCurve2DEndHalfwayArcShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	topstackgrowthcurve2dendhalfwayarcshapeFormCallback.probe.formStage.Checkout()

	if topstackgrowthcurve2dendhalfwayarcshapeFormCallback.topstackgrowthcurve2dendhalfwayarcshape == nil {
		topstackgrowthcurve2dendhalfwayarcshapeFormCallback.topstackgrowthcurve2dendhalfwayarcshape = new(models.TopStackGrowthCurve2DEndHalfwayArcShape).Stage(topstackgrowthcurve2dendhalfwayarcshapeFormCallback.probe.stageOfInterest)
	}
	topstackgrowthcurve2dendhalfwayarcshape_ := topstackgrowthcurve2dendhalfwayarcshapeFormCallback.topstackgrowthcurve2dendhalfwayarcshape
	_ = topstackgrowthcurve2dendhalfwayarcshape_

	for _, formDiv := range topstackgrowthcurve2dendhalfwayarcshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(topstackgrowthcurve2dendhalfwayarcshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(topstackgrowthcurve2dendhalfwayarcshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(topstackgrowthcurve2dendhalfwayarcshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(topstackgrowthcurve2dendhalfwayarcshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(topstackgrowthcurve2dendhalfwayarcshape_.EndY), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(topstackgrowthcurve2dendhalfwayarcshape_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(topstackgrowthcurve2dendhalfwayarcshape_.RadiusY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(topstackgrowthcurve2dendhalfwayarcshape_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(topstackgrowthcurve2dendhalfwayarcshape_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(topstackgrowthcurve2dendhalfwayarcshape_.SweepFlag), formDiv)
		case "TopStackOfGrowthCurve2D:TopStackGrowthCurve2DEndHalfwayArcShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the TopStackOfGrowthCurve2D instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target TopStackOfGrowthCurve2D instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.TopStackOfGrowthCurve2D](topstackgrowthcurve2dendhalfwayarcshapeFormCallback.probe.stageOfInterest)
			targetTopStackOfGrowthCurve2DIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetTopStackOfGrowthCurve2DIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all TopStackOfGrowthCurve2D instances and update their TopStackGrowthCurve2DEndHalfwayArcShapes slice
			for _topstackofgrowthcurve2d := range *models.GetGongstructInstancesSetFromPointerType[*models.TopStackOfGrowthCurve2D](topstackgrowthcurve2dendhalfwayarcshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(topstackgrowthcurve2dendhalfwayarcshapeFormCallback.probe.stageOfInterest, _topstackofgrowthcurve2d)
				
				// if TopStackOfGrowthCurve2D is selected
				if targetTopStackOfGrowthCurve2DIDs[id] {
					// ensure topstackgrowthcurve2dendhalfwayarcshape_ is in _topstackofgrowthcurve2d.TopStackGrowthCurve2DEndHalfwayArcShapes
					found := false
					for _, _b := range _topstackofgrowthcurve2d.TopStackGrowthCurve2DEndHalfwayArcShapes {
						if _b == topstackgrowthcurve2dendhalfwayarcshape_ {
							found = true
							break
						}
					}
					if !found {
						_topstackofgrowthcurve2d.TopStackGrowthCurve2DEndHalfwayArcShapes = append(_topstackofgrowthcurve2d.TopStackGrowthCurve2DEndHalfwayArcShapes, topstackgrowthcurve2dendhalfwayarcshape_)
						topstackgrowthcurve2dendhalfwayarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_topstackofgrowthcurve2d, "TopStackGrowthCurve2DEndHalfwayArcShapes", &_topstackofgrowthcurve2d.TopStackGrowthCurve2DEndHalfwayArcShapes)
					}
				} else {
					// ensure topstackgrowthcurve2dendhalfwayarcshape_ is NOT in _topstackofgrowthcurve2d.TopStackGrowthCurve2DEndHalfwayArcShapes
					idx := slices.Index(_topstackofgrowthcurve2d.TopStackGrowthCurve2DEndHalfwayArcShapes, topstackgrowthcurve2dendhalfwayarcshape_)
					if idx != -1 {
						_topstackofgrowthcurve2d.TopStackGrowthCurve2DEndHalfwayArcShapes = slices.Delete(_topstackofgrowthcurve2d.TopStackGrowthCurve2DEndHalfwayArcShapes, idx, idx+1)
						topstackgrowthcurve2dendhalfwayarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_topstackofgrowthcurve2d, "TopStackGrowthCurve2DEndHalfwayArcShapes", &_topstackofgrowthcurve2d.TopStackGrowthCurve2DEndHalfwayArcShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if topstackgrowthcurve2dendhalfwayarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstackgrowthcurve2dendhalfwayarcshape_.Unstage(topstackgrowthcurve2dendhalfwayarcshapeFormCallback.probe.stageOfInterest)
	}

	topstackgrowthcurve2dendhalfwayarcshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TopStackGrowthCurve2DEndHalfwayArcShape](
		topstackgrowthcurve2dendhalfwayarcshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if topstackgrowthcurve2dendhalfwayarcshapeFormCallback.CreationMode || topstackgrowthcurve2dendhalfwayarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstackgrowthcurve2dendhalfwayarcshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(topstackgrowthcurve2dendhalfwayarcshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TopStackGrowthCurve2DEndHalfwayArcShapeFormCallback(
			nil,
			topstackgrowthcurve2dendhalfwayarcshapeFormCallback.probe,
			newFormGroup,
		)
		topstackgrowthcurve2dendhalfwayarcshape := new(models.TopStackGrowthCurve2DEndHalfwayArcShape)
		FillUpForm(topstackgrowthcurve2dendhalfwayarcshape, newFormGroup, topstackgrowthcurve2dendhalfwayarcshapeFormCallback.probe)
		topstackgrowthcurve2dendhalfwayarcshapeFormCallback.probe.formStage.Commit()
	}

	topstackgrowthcurve2dendhalfwayarcshapeFormCallback.probe.ux_tree()
}
func __gong__New__TopStackGrowthCurve2DStartHalfwayArcShapeFormCallback(
	topstackgrowthcurve2dstarthalfwayarcshape *models.TopStackGrowthCurve2DStartHalfwayArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (topstackgrowthcurve2dstarthalfwayarcshapeFormCallback *TopStackGrowthCurve2DStartHalfwayArcShapeFormCallback) {
	topstackgrowthcurve2dstarthalfwayarcshapeFormCallback = new(TopStackGrowthCurve2DStartHalfwayArcShapeFormCallback)
	topstackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe = probe
	topstackgrowthcurve2dstarthalfwayarcshapeFormCallback.topstackgrowthcurve2dstarthalfwayarcshape = topstackgrowthcurve2dstarthalfwayarcshape
	topstackgrowthcurve2dstarthalfwayarcshapeFormCallback.formGroup = formGroup

	topstackgrowthcurve2dstarthalfwayarcshapeFormCallback.CreationMode = (topstackgrowthcurve2dstarthalfwayarcshape == nil)

	return
}

type TopStackGrowthCurve2DStartHalfwayArcShapeFormCallback struct {
	topstackgrowthcurve2dstarthalfwayarcshape *models.TopStackGrowthCurve2DStartHalfwayArcShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (topstackgrowthcurve2dstarthalfwayarcshapeFormCallback *TopStackGrowthCurve2DStartHalfwayArcShapeFormCallback) OnSave() {
	topstackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe.stageOfInterest.Lock()
	defer topstackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TopStackGrowthCurve2DStartHalfwayArcShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	topstackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe.formStage.Checkout()

	if topstackgrowthcurve2dstarthalfwayarcshapeFormCallback.topstackgrowthcurve2dstarthalfwayarcshape == nil {
		topstackgrowthcurve2dstarthalfwayarcshapeFormCallback.topstackgrowthcurve2dstarthalfwayarcshape = new(models.TopStackGrowthCurve2DStartHalfwayArcShape).Stage(topstackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe.stageOfInterest)
	}
	topstackgrowthcurve2dstarthalfwayarcshape_ := topstackgrowthcurve2dstarthalfwayarcshapeFormCallback.topstackgrowthcurve2dstarthalfwayarcshape
	_ = topstackgrowthcurve2dstarthalfwayarcshape_

	for _, formDiv := range topstackgrowthcurve2dstarthalfwayarcshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(topstackgrowthcurve2dstarthalfwayarcshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(topstackgrowthcurve2dstarthalfwayarcshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(topstackgrowthcurve2dstarthalfwayarcshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(topstackgrowthcurve2dstarthalfwayarcshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(topstackgrowthcurve2dstarthalfwayarcshape_.EndY), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(topstackgrowthcurve2dstarthalfwayarcshape_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(topstackgrowthcurve2dstarthalfwayarcshape_.RadiusY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(topstackgrowthcurve2dstarthalfwayarcshape_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(topstackgrowthcurve2dstarthalfwayarcshape_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(topstackgrowthcurve2dstarthalfwayarcshape_.SweepFlag), formDiv)
		case "TopStackOfGrowthCurve2D:TopStackGrowthCurve2DStartHalfwayArcShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the TopStackOfGrowthCurve2D instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target TopStackOfGrowthCurve2D instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.TopStackOfGrowthCurve2D](topstackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe.stageOfInterest)
			targetTopStackOfGrowthCurve2DIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetTopStackOfGrowthCurve2DIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all TopStackOfGrowthCurve2D instances and update their TopStackGrowthCurve2DStartHalfwayArcShapes slice
			for _topstackofgrowthcurve2d := range *models.GetGongstructInstancesSetFromPointerType[*models.TopStackOfGrowthCurve2D](topstackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(topstackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe.stageOfInterest, _topstackofgrowthcurve2d)
				
				// if TopStackOfGrowthCurve2D is selected
				if targetTopStackOfGrowthCurve2DIDs[id] {
					// ensure topstackgrowthcurve2dstarthalfwayarcshape_ is in _topstackofgrowthcurve2d.TopStackGrowthCurve2DStartHalfwayArcShapes
					found := false
					for _, _b := range _topstackofgrowthcurve2d.TopStackGrowthCurve2DStartHalfwayArcShapes {
						if _b == topstackgrowthcurve2dstarthalfwayarcshape_ {
							found = true
							break
						}
					}
					if !found {
						_topstackofgrowthcurve2d.TopStackGrowthCurve2DStartHalfwayArcShapes = append(_topstackofgrowthcurve2d.TopStackGrowthCurve2DStartHalfwayArcShapes, topstackgrowthcurve2dstarthalfwayarcshape_)
						topstackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_topstackofgrowthcurve2d, "TopStackGrowthCurve2DStartHalfwayArcShapes", &_topstackofgrowthcurve2d.TopStackGrowthCurve2DStartHalfwayArcShapes)
					}
				} else {
					// ensure topstackgrowthcurve2dstarthalfwayarcshape_ is NOT in _topstackofgrowthcurve2d.TopStackGrowthCurve2DStartHalfwayArcShapes
					idx := slices.Index(_topstackofgrowthcurve2d.TopStackGrowthCurve2DStartHalfwayArcShapes, topstackgrowthcurve2dstarthalfwayarcshape_)
					if idx != -1 {
						_topstackofgrowthcurve2d.TopStackGrowthCurve2DStartHalfwayArcShapes = slices.Delete(_topstackofgrowthcurve2d.TopStackGrowthCurve2DStartHalfwayArcShapes, idx, idx+1)
						topstackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_topstackofgrowthcurve2d, "TopStackGrowthCurve2DStartHalfwayArcShapes", &_topstackofgrowthcurve2d.TopStackGrowthCurve2DStartHalfwayArcShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if topstackgrowthcurve2dstarthalfwayarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstackgrowthcurve2dstarthalfwayarcshape_.Unstage(topstackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe.stageOfInterest)
	}

	topstackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TopStackGrowthCurve2DStartHalfwayArcShape](
		topstackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if topstackgrowthcurve2dstarthalfwayarcshapeFormCallback.CreationMode || topstackgrowthcurve2dstarthalfwayarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(topstackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TopStackGrowthCurve2DStartHalfwayArcShapeFormCallback(
			nil,
			topstackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe,
			newFormGroup,
		)
		topstackgrowthcurve2dstarthalfwayarcshape := new(models.TopStackGrowthCurve2DStartHalfwayArcShape)
		FillUpForm(topstackgrowthcurve2dstarthalfwayarcshape, newFormGroup, topstackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe)
		topstackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe.formStage.Commit()
	}

	topstackgrowthcurve2dstarthalfwayarcshapeFormCallback.probe.ux_tree()
}
func __gong__New__TopStackOfGrowthCurve2DFormCallback(
	topstackofgrowthcurve2d *models.TopStackOfGrowthCurve2D,
	probe *Probe,
	formGroup *form.FormGroup,
) (topstackofgrowthcurve2dFormCallback *TopStackOfGrowthCurve2DFormCallback) {
	topstackofgrowthcurve2dFormCallback = new(TopStackOfGrowthCurve2DFormCallback)
	topstackofgrowthcurve2dFormCallback.probe = probe
	topstackofgrowthcurve2dFormCallback.topstackofgrowthcurve2d = topstackofgrowthcurve2d
	topstackofgrowthcurve2dFormCallback.formGroup = formGroup

	topstackofgrowthcurve2dFormCallback.CreationMode = (topstackofgrowthcurve2d == nil)

	return
}

type TopStackOfGrowthCurve2DFormCallback struct {
	topstackofgrowthcurve2d *models.TopStackOfGrowthCurve2D

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (topstackofgrowthcurve2dFormCallback *TopStackOfGrowthCurve2DFormCallback) OnSave() {
	topstackofgrowthcurve2dFormCallback.probe.stageOfInterest.Lock()
	defer topstackofgrowthcurve2dFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TopStackOfGrowthCurve2DFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	topstackofgrowthcurve2dFormCallback.probe.formStage.Checkout()

	if topstackofgrowthcurve2dFormCallback.topstackofgrowthcurve2d == nil {
		topstackofgrowthcurve2dFormCallback.topstackofgrowthcurve2d = new(models.TopStackOfGrowthCurve2D).Stage(topstackofgrowthcurve2dFormCallback.probe.stageOfInterest)
	}
	topstackofgrowthcurve2d_ := topstackofgrowthcurve2dFormCallback.topstackofgrowthcurve2d
	_ = topstackofgrowthcurve2d_

	for _, formDiv := range topstackofgrowthcurve2dFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(topstackofgrowthcurve2d_.Name), formDiv)
		case "TopStackGrowthCurve2DStartHalfwayArcShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.TopStackGrowthCurve2DStartHalfwayArcShape](topstackofgrowthcurve2dFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.TopStackGrowthCurve2DStartHalfwayArcShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.TopStackGrowthCurve2DStartHalfwayArcShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					topstackofgrowthcurve2dFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.TopStackGrowthCurve2DStartHalfwayArcShape](topstackofgrowthcurve2dFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			topstackofgrowthcurve2d_.TopStackGrowthCurve2DStartHalfwayArcShapes = instanceSlice
			topstackofgrowthcurve2dFormCallback.probe.UpdateSliceOfPointersCallback(topstackofgrowthcurve2d_, "TopStackGrowthCurve2DStartHalfwayArcShapes", &topstackofgrowthcurve2d_.TopStackGrowthCurve2DStartHalfwayArcShapes)

		case "TopStackGrowthCurve2DEndHalfwayArcShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.TopStackGrowthCurve2DEndHalfwayArcShape](topstackofgrowthcurve2dFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.TopStackGrowthCurve2DEndHalfwayArcShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.TopStackGrowthCurve2DEndHalfwayArcShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					topstackofgrowthcurve2dFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.TopStackGrowthCurve2DEndHalfwayArcShape](topstackofgrowthcurve2dFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			topstackofgrowthcurve2d_.TopStackGrowthCurve2DEndHalfwayArcShapes = instanceSlice
			topstackofgrowthcurve2dFormCallback.probe.UpdateSliceOfPointersCallback(topstackofgrowthcurve2d_, "TopStackGrowthCurve2DEndHalfwayArcShapes", &topstackofgrowthcurve2d_.TopStackGrowthCurve2DEndHalfwayArcShapes)

		}
	}

	// manage the suppress operation
	if topstackofgrowthcurve2dFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstackofgrowthcurve2d_.Unstage(topstackofgrowthcurve2dFormCallback.probe.stageOfInterest)
	}

	topstackofgrowthcurve2dFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TopStackOfGrowthCurve2D](
		topstackofgrowthcurve2dFormCallback.probe,
	)

	// display a new form by reset the form stage
	if topstackofgrowthcurve2dFormCallback.CreationMode || topstackofgrowthcurve2dFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstackofgrowthcurve2dFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(topstackofgrowthcurve2dFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TopStackOfGrowthCurve2DFormCallback(
			nil,
			topstackofgrowthcurve2dFormCallback.probe,
			newFormGroup,
		)
		topstackofgrowthcurve2d := new(models.TopStackOfGrowthCurve2D)
		FillUpForm(topstackofgrowthcurve2d, newFormGroup, topstackofgrowthcurve2dFormCallback.probe)
		topstackofgrowthcurve2dFormCallback.probe.formStage.Commit()
	}

	topstackofgrowthcurve2dFormCallback.probe.ux_tree()
}
func __gong__New__TopStackOfRotatedGrowthCurve2DFormCallback(
	topstackofrotatedgrowthcurve2d *models.TopStackOfRotatedGrowthCurve2D,
	probe *Probe,
	formGroup *form.FormGroup,
) (topstackofrotatedgrowthcurve2dFormCallback *TopStackOfRotatedGrowthCurve2DFormCallback) {
	topstackofrotatedgrowthcurve2dFormCallback = new(TopStackOfRotatedGrowthCurve2DFormCallback)
	topstackofrotatedgrowthcurve2dFormCallback.probe = probe
	topstackofrotatedgrowthcurve2dFormCallback.topstackofrotatedgrowthcurve2d = topstackofrotatedgrowthcurve2d
	topstackofrotatedgrowthcurve2dFormCallback.formGroup = formGroup

	topstackofrotatedgrowthcurve2dFormCallback.CreationMode = (topstackofrotatedgrowthcurve2d == nil)

	return
}

type TopStackOfRotatedGrowthCurve2DFormCallback struct {
	topstackofrotatedgrowthcurve2d *models.TopStackOfRotatedGrowthCurve2D

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (topstackofrotatedgrowthcurve2dFormCallback *TopStackOfRotatedGrowthCurve2DFormCallback) OnSave() {
	topstackofrotatedgrowthcurve2dFormCallback.probe.stageOfInterest.Lock()
	defer topstackofrotatedgrowthcurve2dFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TopStackOfRotatedGrowthCurve2DFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	topstackofrotatedgrowthcurve2dFormCallback.probe.formStage.Checkout()

	if topstackofrotatedgrowthcurve2dFormCallback.topstackofrotatedgrowthcurve2d == nil {
		topstackofrotatedgrowthcurve2dFormCallback.topstackofrotatedgrowthcurve2d = new(models.TopStackOfRotatedGrowthCurve2D).Stage(topstackofrotatedgrowthcurve2dFormCallback.probe.stageOfInterest)
	}
	topstackofrotatedgrowthcurve2d_ := topstackofrotatedgrowthcurve2dFormCallback.topstackofrotatedgrowthcurve2d
	_ = topstackofrotatedgrowthcurve2d_

	for _, formDiv := range topstackofrotatedgrowthcurve2dFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(topstackofrotatedgrowthcurve2d_.Name), formDiv)
		case "TopStackOfRotatedGrowthCurve2DStartArcShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.TopStackOfRotatedGrowthCurve2DStartArcShape](topstackofrotatedgrowthcurve2dFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.TopStackOfRotatedGrowthCurve2DStartArcShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.TopStackOfRotatedGrowthCurve2DStartArcShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					topstackofrotatedgrowthcurve2dFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.TopStackOfRotatedGrowthCurve2DStartArcShape](topstackofrotatedgrowthcurve2dFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			topstackofrotatedgrowthcurve2d_.TopStackOfRotatedGrowthCurve2DStartArcShapes = instanceSlice
			topstackofrotatedgrowthcurve2dFormCallback.probe.UpdateSliceOfPointersCallback(topstackofrotatedgrowthcurve2d_, "TopStackOfRotatedGrowthCurve2DStartArcShapes", &topstackofrotatedgrowthcurve2d_.TopStackOfRotatedGrowthCurve2DStartArcShapes)

		case "TopStackOfRotatedGrowthCurve2DEndArcShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.TopStackOfRotatedGrowthCurve2DEndArcShape](topstackofrotatedgrowthcurve2dFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.TopStackOfRotatedGrowthCurve2DEndArcShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.TopStackOfRotatedGrowthCurve2DEndArcShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					topstackofrotatedgrowthcurve2dFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.TopStackOfRotatedGrowthCurve2DEndArcShape](topstackofrotatedgrowthcurve2dFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			topstackofrotatedgrowthcurve2d_.TopStackOfRotatedGrowthCurve2DEndArcShapes = instanceSlice
			topstackofrotatedgrowthcurve2dFormCallback.probe.UpdateSliceOfPointersCallback(topstackofrotatedgrowthcurve2d_, "TopStackOfRotatedGrowthCurve2DEndArcShapes", &topstackofrotatedgrowthcurve2d_.TopStackOfRotatedGrowthCurve2DEndArcShapes)

		}
	}

	// manage the suppress operation
	if topstackofrotatedgrowthcurve2dFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstackofrotatedgrowthcurve2d_.Unstage(topstackofrotatedgrowthcurve2dFormCallback.probe.stageOfInterest)
	}

	topstackofrotatedgrowthcurve2dFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TopStackOfRotatedGrowthCurve2D](
		topstackofrotatedgrowthcurve2dFormCallback.probe,
	)

	// display a new form by reset the form stage
	if topstackofrotatedgrowthcurve2dFormCallback.CreationMode || topstackofrotatedgrowthcurve2dFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstackofrotatedgrowthcurve2dFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(topstackofrotatedgrowthcurve2dFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TopStackOfRotatedGrowthCurve2DFormCallback(
			nil,
			topstackofrotatedgrowthcurve2dFormCallback.probe,
			newFormGroup,
		)
		topstackofrotatedgrowthcurve2d := new(models.TopStackOfRotatedGrowthCurve2D)
		FillUpForm(topstackofrotatedgrowthcurve2d, newFormGroup, topstackofrotatedgrowthcurve2dFormCallback.probe)
		topstackofrotatedgrowthcurve2dFormCallback.probe.formStage.Commit()
	}

	topstackofrotatedgrowthcurve2dFormCallback.probe.ux_tree()
}
func __gong__New__TopStackOfRotatedGrowthCurve2DEndArcShapeFormCallback(
	topstackofrotatedgrowthcurve2dendarcshape *models.TopStackOfRotatedGrowthCurve2DEndArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (topstackofrotatedgrowthcurve2dendarcshapeFormCallback *TopStackOfRotatedGrowthCurve2DEndArcShapeFormCallback) {
	topstackofrotatedgrowthcurve2dendarcshapeFormCallback = new(TopStackOfRotatedGrowthCurve2DEndArcShapeFormCallback)
	topstackofrotatedgrowthcurve2dendarcshapeFormCallback.probe = probe
	topstackofrotatedgrowthcurve2dendarcshapeFormCallback.topstackofrotatedgrowthcurve2dendarcshape = topstackofrotatedgrowthcurve2dendarcshape
	topstackofrotatedgrowthcurve2dendarcshapeFormCallback.formGroup = formGroup

	topstackofrotatedgrowthcurve2dendarcshapeFormCallback.CreationMode = (topstackofrotatedgrowthcurve2dendarcshape == nil)

	return
}

type TopStackOfRotatedGrowthCurve2DEndArcShapeFormCallback struct {
	topstackofrotatedgrowthcurve2dendarcshape *models.TopStackOfRotatedGrowthCurve2DEndArcShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (topstackofrotatedgrowthcurve2dendarcshapeFormCallback *TopStackOfRotatedGrowthCurve2DEndArcShapeFormCallback) OnSave() {
	topstackofrotatedgrowthcurve2dendarcshapeFormCallback.probe.stageOfInterest.Lock()
	defer topstackofrotatedgrowthcurve2dendarcshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TopStackOfRotatedGrowthCurve2DEndArcShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	topstackofrotatedgrowthcurve2dendarcshapeFormCallback.probe.formStage.Checkout()

	if topstackofrotatedgrowthcurve2dendarcshapeFormCallback.topstackofrotatedgrowthcurve2dendarcshape == nil {
		topstackofrotatedgrowthcurve2dendarcshapeFormCallback.topstackofrotatedgrowthcurve2dendarcshape = new(models.TopStackOfRotatedGrowthCurve2DEndArcShape).Stage(topstackofrotatedgrowthcurve2dendarcshapeFormCallback.probe.stageOfInterest)
	}
	topstackofrotatedgrowthcurve2dendarcshape_ := topstackofrotatedgrowthcurve2dendarcshapeFormCallback.topstackofrotatedgrowthcurve2dendarcshape
	_ = topstackofrotatedgrowthcurve2dendarcshape_

	for _, formDiv := range topstackofrotatedgrowthcurve2dendarcshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(topstackofrotatedgrowthcurve2dendarcshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(topstackofrotatedgrowthcurve2dendarcshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(topstackofrotatedgrowthcurve2dendarcshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(topstackofrotatedgrowthcurve2dendarcshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(topstackofrotatedgrowthcurve2dendarcshape_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(topstackofrotatedgrowthcurve2dendarcshape_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(topstackofrotatedgrowthcurve2dendarcshape_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(topstackofrotatedgrowthcurve2dendarcshape_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(topstackofrotatedgrowthcurve2dendarcshape_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(topstackofrotatedgrowthcurve2dendarcshape_.RadiusY), formDiv)
		case "TopStackOfRotatedGrowthCurve2D:TopStackOfRotatedGrowthCurve2DEndArcShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the TopStackOfRotatedGrowthCurve2D instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target TopStackOfRotatedGrowthCurve2D instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.TopStackOfRotatedGrowthCurve2D](topstackofrotatedgrowthcurve2dendarcshapeFormCallback.probe.stageOfInterest)
			targetTopStackOfRotatedGrowthCurve2DIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetTopStackOfRotatedGrowthCurve2DIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all TopStackOfRotatedGrowthCurve2D instances and update their TopStackOfRotatedGrowthCurve2DEndArcShapes slice
			for _topstackofrotatedgrowthcurve2d := range *models.GetGongstructInstancesSetFromPointerType[*models.TopStackOfRotatedGrowthCurve2D](topstackofrotatedgrowthcurve2dendarcshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(topstackofrotatedgrowthcurve2dendarcshapeFormCallback.probe.stageOfInterest, _topstackofrotatedgrowthcurve2d)
				
				// if TopStackOfRotatedGrowthCurve2D is selected
				if targetTopStackOfRotatedGrowthCurve2DIDs[id] {
					// ensure topstackofrotatedgrowthcurve2dendarcshape_ is in _topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DEndArcShapes
					found := false
					for _, _b := range _topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DEndArcShapes {
						if _b == topstackofrotatedgrowthcurve2dendarcshape_ {
							found = true
							break
						}
					}
					if !found {
						_topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DEndArcShapes = append(_topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DEndArcShapes, topstackofrotatedgrowthcurve2dendarcshape_)
						topstackofrotatedgrowthcurve2dendarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_topstackofrotatedgrowthcurve2d, "TopStackOfRotatedGrowthCurve2DEndArcShapes", &_topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DEndArcShapes)
					}
				} else {
					// ensure topstackofrotatedgrowthcurve2dendarcshape_ is NOT in _topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DEndArcShapes
					idx := slices.Index(_topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DEndArcShapes, topstackofrotatedgrowthcurve2dendarcshape_)
					if idx != -1 {
						_topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DEndArcShapes = slices.Delete(_topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DEndArcShapes, idx, idx+1)
						topstackofrotatedgrowthcurve2dendarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_topstackofrotatedgrowthcurve2d, "TopStackOfRotatedGrowthCurve2DEndArcShapes", &_topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DEndArcShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if topstackofrotatedgrowthcurve2dendarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstackofrotatedgrowthcurve2dendarcshape_.Unstage(topstackofrotatedgrowthcurve2dendarcshapeFormCallback.probe.stageOfInterest)
	}

	topstackofrotatedgrowthcurve2dendarcshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TopStackOfRotatedGrowthCurve2DEndArcShape](
		topstackofrotatedgrowthcurve2dendarcshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if topstackofrotatedgrowthcurve2dendarcshapeFormCallback.CreationMode || topstackofrotatedgrowthcurve2dendarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstackofrotatedgrowthcurve2dendarcshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(topstackofrotatedgrowthcurve2dendarcshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TopStackOfRotatedGrowthCurve2DEndArcShapeFormCallback(
			nil,
			topstackofrotatedgrowthcurve2dendarcshapeFormCallback.probe,
			newFormGroup,
		)
		topstackofrotatedgrowthcurve2dendarcshape := new(models.TopStackOfRotatedGrowthCurve2DEndArcShape)
		FillUpForm(topstackofrotatedgrowthcurve2dendarcshape, newFormGroup, topstackofrotatedgrowthcurve2dendarcshapeFormCallback.probe)
		topstackofrotatedgrowthcurve2dendarcshapeFormCallback.probe.formStage.Commit()
	}

	topstackofrotatedgrowthcurve2dendarcshapeFormCallback.probe.ux_tree()
}
func __gong__New__TopStackOfRotatedGrowthCurve2DStartArcShapeFormCallback(
	topstackofrotatedgrowthcurve2dstartarcshape *models.TopStackOfRotatedGrowthCurve2DStartArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (topstackofrotatedgrowthcurve2dstartarcshapeFormCallback *TopStackOfRotatedGrowthCurve2DStartArcShapeFormCallback) {
	topstackofrotatedgrowthcurve2dstartarcshapeFormCallback = new(TopStackOfRotatedGrowthCurve2DStartArcShapeFormCallback)
	topstackofrotatedgrowthcurve2dstartarcshapeFormCallback.probe = probe
	topstackofrotatedgrowthcurve2dstartarcshapeFormCallback.topstackofrotatedgrowthcurve2dstartarcshape = topstackofrotatedgrowthcurve2dstartarcshape
	topstackofrotatedgrowthcurve2dstartarcshapeFormCallback.formGroup = formGroup

	topstackofrotatedgrowthcurve2dstartarcshapeFormCallback.CreationMode = (topstackofrotatedgrowthcurve2dstartarcshape == nil)

	return
}

type TopStackOfRotatedGrowthCurve2DStartArcShapeFormCallback struct {
	topstackofrotatedgrowthcurve2dstartarcshape *models.TopStackOfRotatedGrowthCurve2DStartArcShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (topstackofrotatedgrowthcurve2dstartarcshapeFormCallback *TopStackOfRotatedGrowthCurve2DStartArcShapeFormCallback) OnSave() {
	topstackofrotatedgrowthcurve2dstartarcshapeFormCallback.probe.stageOfInterest.Lock()
	defer topstackofrotatedgrowthcurve2dstartarcshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TopStackOfRotatedGrowthCurve2DStartArcShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	topstackofrotatedgrowthcurve2dstartarcshapeFormCallback.probe.formStage.Checkout()

	if topstackofrotatedgrowthcurve2dstartarcshapeFormCallback.topstackofrotatedgrowthcurve2dstartarcshape == nil {
		topstackofrotatedgrowthcurve2dstartarcshapeFormCallback.topstackofrotatedgrowthcurve2dstartarcshape = new(models.TopStackOfRotatedGrowthCurve2DStartArcShape).Stage(topstackofrotatedgrowthcurve2dstartarcshapeFormCallback.probe.stageOfInterest)
	}
	topstackofrotatedgrowthcurve2dstartarcshape_ := topstackofrotatedgrowthcurve2dstartarcshapeFormCallback.topstackofrotatedgrowthcurve2dstartarcshape
	_ = topstackofrotatedgrowthcurve2dstartarcshape_

	for _, formDiv := range topstackofrotatedgrowthcurve2dstartarcshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(topstackofrotatedgrowthcurve2dstartarcshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(topstackofrotatedgrowthcurve2dstartarcshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(topstackofrotatedgrowthcurve2dstartarcshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(topstackofrotatedgrowthcurve2dstartarcshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(topstackofrotatedgrowthcurve2dstartarcshape_.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(topstackofrotatedgrowthcurve2dstartarcshape_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(topstackofrotatedgrowthcurve2dstartarcshape_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(topstackofrotatedgrowthcurve2dstartarcshape_.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(topstackofrotatedgrowthcurve2dstartarcshape_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(topstackofrotatedgrowthcurve2dstartarcshape_.RadiusY), formDiv)
		case "TopStackOfRotatedGrowthCurve2D:TopStackOfRotatedGrowthCurve2DStartArcShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the TopStackOfRotatedGrowthCurve2D instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target TopStackOfRotatedGrowthCurve2D instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.TopStackOfRotatedGrowthCurve2D](topstackofrotatedgrowthcurve2dstartarcshapeFormCallback.probe.stageOfInterest)
			targetTopStackOfRotatedGrowthCurve2DIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetTopStackOfRotatedGrowthCurve2DIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all TopStackOfRotatedGrowthCurve2D instances and update their TopStackOfRotatedGrowthCurve2DStartArcShapes slice
			for _topstackofrotatedgrowthcurve2d := range *models.GetGongstructInstancesSetFromPointerType[*models.TopStackOfRotatedGrowthCurve2D](topstackofrotatedgrowthcurve2dstartarcshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(topstackofrotatedgrowthcurve2dstartarcshapeFormCallback.probe.stageOfInterest, _topstackofrotatedgrowthcurve2d)
				
				// if TopStackOfRotatedGrowthCurve2D is selected
				if targetTopStackOfRotatedGrowthCurve2DIDs[id] {
					// ensure topstackofrotatedgrowthcurve2dstartarcshape_ is in _topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DStartArcShapes
					found := false
					for _, _b := range _topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DStartArcShapes {
						if _b == topstackofrotatedgrowthcurve2dstartarcshape_ {
							found = true
							break
						}
					}
					if !found {
						_topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DStartArcShapes = append(_topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DStartArcShapes, topstackofrotatedgrowthcurve2dstartarcshape_)
						topstackofrotatedgrowthcurve2dstartarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_topstackofrotatedgrowthcurve2d, "TopStackOfRotatedGrowthCurve2DStartArcShapes", &_topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DStartArcShapes)
					}
				} else {
					// ensure topstackofrotatedgrowthcurve2dstartarcshape_ is NOT in _topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DStartArcShapes
					idx := slices.Index(_topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DStartArcShapes, topstackofrotatedgrowthcurve2dstartarcshape_)
					if idx != -1 {
						_topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DStartArcShapes = slices.Delete(_topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DStartArcShapes, idx, idx+1)
						topstackofrotatedgrowthcurve2dstartarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_topstackofrotatedgrowthcurve2d, "TopStackOfRotatedGrowthCurve2DStartArcShapes", &_topstackofrotatedgrowthcurve2d.TopStackOfRotatedGrowthCurve2DStartArcShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if topstackofrotatedgrowthcurve2dstartarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstackofrotatedgrowthcurve2dstartarcshape_.Unstage(topstackofrotatedgrowthcurve2dstartarcshapeFormCallback.probe.stageOfInterest)
	}

	topstackofrotatedgrowthcurve2dstartarcshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TopStackOfRotatedGrowthCurve2DStartArcShape](
		topstackofrotatedgrowthcurve2dstartarcshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if topstackofrotatedgrowthcurve2dstartarcshapeFormCallback.CreationMode || topstackofrotatedgrowthcurve2dstartarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstackofrotatedgrowthcurve2dstartarcshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(topstackofrotatedgrowthcurve2dstartarcshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TopStackOfRotatedGrowthCurve2DStartArcShapeFormCallback(
			nil,
			topstackofrotatedgrowthcurve2dstartarcshapeFormCallback.probe,
			newFormGroup,
		)
		topstackofrotatedgrowthcurve2dstartarcshape := new(models.TopStackOfRotatedGrowthCurve2DStartArcShape)
		FillUpForm(topstackofrotatedgrowthcurve2dstartarcshape, newFormGroup, topstackofrotatedgrowthcurve2dstartarcshapeFormCallback.probe)
		topstackofrotatedgrowthcurve2dstartarcshapeFormCallback.probe.formStage.Commit()
	}

	topstackofrotatedgrowthcurve2dstartarcshapeFormCallback.probe.ux_tree()
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
func __gong__New__TopStartHalfwayArcShapeFormCallback(
	topstarthalfwayarcshape *models.TopStartHalfwayArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (topstarthalfwayarcshapeFormCallback *TopStartHalfwayArcShapeFormCallback) {
	topstarthalfwayarcshapeFormCallback = new(TopStartHalfwayArcShapeFormCallback)
	topstarthalfwayarcshapeFormCallback.probe = probe
	topstarthalfwayarcshapeFormCallback.topstarthalfwayarcshape = topstarthalfwayarcshape
	topstarthalfwayarcshapeFormCallback.formGroup = formGroup

	topstarthalfwayarcshapeFormCallback.CreationMode = (topstarthalfwayarcshape == nil)

	return
}

type TopStartHalfwayArcShapeFormCallback struct {
	topstarthalfwayarcshape *models.TopStartHalfwayArcShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (topstarthalfwayarcshapeFormCallback *TopStartHalfwayArcShapeFormCallback) OnSave() {
	topstarthalfwayarcshapeFormCallback.probe.stageOfInterest.Lock()
	defer topstarthalfwayarcshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TopStartHalfwayArcShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	topstarthalfwayarcshapeFormCallback.probe.formStage.Checkout()

	if topstarthalfwayarcshapeFormCallback.topstarthalfwayarcshape == nil {
		topstarthalfwayarcshapeFormCallback.topstarthalfwayarcshape = new(models.TopStartHalfwayArcShape).Stage(topstarthalfwayarcshapeFormCallback.probe.stageOfInterest)
	}
	topstarthalfwayarcshape_ := topstarthalfwayarcshapeFormCallback.topstarthalfwayarcshape
	_ = topstarthalfwayarcshape_

	for _, formDiv := range topstarthalfwayarcshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(topstarthalfwayarcshape_.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(topstarthalfwayarcshape_.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(topstarthalfwayarcshape_.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(topstarthalfwayarcshape_.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(topstarthalfwayarcshape_.EndY), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(topstarthalfwayarcshape_.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(topstarthalfwayarcshape_.RadiusY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(topstarthalfwayarcshape_.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(topstarthalfwayarcshape_.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(topstarthalfwayarcshape_.SweepFlag), formDiv)
		case "TopStartHalfwayArcShapeGrid:TopStartHalfwayArcShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the TopStartHalfwayArcShapeGrid instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target TopStartHalfwayArcShapeGrid instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.TopStartHalfwayArcShapeGrid](topstarthalfwayarcshapeFormCallback.probe.stageOfInterest)
			targetTopStartHalfwayArcShapeGridIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetTopStartHalfwayArcShapeGridIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all TopStartHalfwayArcShapeGrid instances and update their TopStartHalfwayArcShapes slice
			for _topstarthalfwayarcshapegrid := range *models.GetGongstructInstancesSetFromPointerType[*models.TopStartHalfwayArcShapeGrid](topstarthalfwayarcshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(topstarthalfwayarcshapeFormCallback.probe.stageOfInterest, _topstarthalfwayarcshapegrid)
				
				// if TopStartHalfwayArcShapeGrid is selected
				if targetTopStartHalfwayArcShapeGridIDs[id] {
					// ensure topstarthalfwayarcshape_ is in _topstarthalfwayarcshapegrid.TopStartHalfwayArcShapes
					found := false
					for _, _b := range _topstarthalfwayarcshapegrid.TopStartHalfwayArcShapes {
						if _b == topstarthalfwayarcshape_ {
							found = true
							break
						}
					}
					if !found {
						_topstarthalfwayarcshapegrid.TopStartHalfwayArcShapes = append(_topstarthalfwayarcshapegrid.TopStartHalfwayArcShapes, topstarthalfwayarcshape_)
						topstarthalfwayarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_topstarthalfwayarcshapegrid, "TopStartHalfwayArcShapes", &_topstarthalfwayarcshapegrid.TopStartHalfwayArcShapes)
					}
				} else {
					// ensure topstarthalfwayarcshape_ is NOT in _topstarthalfwayarcshapegrid.TopStartHalfwayArcShapes
					idx := slices.Index(_topstarthalfwayarcshapegrid.TopStartHalfwayArcShapes, topstarthalfwayarcshape_)
					if idx != -1 {
						_topstarthalfwayarcshapegrid.TopStartHalfwayArcShapes = slices.Delete(_topstarthalfwayarcshapegrid.TopStartHalfwayArcShapes, idx, idx+1)
						topstarthalfwayarcshapeFormCallback.probe.UpdateSliceOfPointersCallback(_topstarthalfwayarcshapegrid, "TopStartHalfwayArcShapes", &_topstarthalfwayarcshapegrid.TopStartHalfwayArcShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if topstarthalfwayarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstarthalfwayarcshape_.Unstage(topstarthalfwayarcshapeFormCallback.probe.stageOfInterest)
	}

	topstarthalfwayarcshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TopStartHalfwayArcShape](
		topstarthalfwayarcshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if topstarthalfwayarcshapeFormCallback.CreationMode || topstarthalfwayarcshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstarthalfwayarcshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(topstarthalfwayarcshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TopStartHalfwayArcShapeFormCallback(
			nil,
			topstarthalfwayarcshapeFormCallback.probe,
			newFormGroup,
		)
		topstarthalfwayarcshape := new(models.TopStartHalfwayArcShape)
		FillUpForm(topstarthalfwayarcshape, newFormGroup, topstarthalfwayarcshapeFormCallback.probe)
		topstarthalfwayarcshapeFormCallback.probe.formStage.Commit()
	}

	topstarthalfwayarcshapeFormCallback.probe.ux_tree()
}
func __gong__New__TopStartHalfwayArcShapeGridFormCallback(
	topstarthalfwayarcshapegrid *models.TopStartHalfwayArcShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (topstarthalfwayarcshapegridFormCallback *TopStartHalfwayArcShapeGridFormCallback) {
	topstarthalfwayarcshapegridFormCallback = new(TopStartHalfwayArcShapeGridFormCallback)
	topstarthalfwayarcshapegridFormCallback.probe = probe
	topstarthalfwayarcshapegridFormCallback.topstarthalfwayarcshapegrid = topstarthalfwayarcshapegrid
	topstarthalfwayarcshapegridFormCallback.formGroup = formGroup

	topstarthalfwayarcshapegridFormCallback.CreationMode = (topstarthalfwayarcshapegrid == nil)

	return
}

type TopStartHalfwayArcShapeGridFormCallback struct {
	topstarthalfwayarcshapegrid *models.TopStartHalfwayArcShapeGrid

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (topstarthalfwayarcshapegridFormCallback *TopStartHalfwayArcShapeGridFormCallback) OnSave() {
	topstarthalfwayarcshapegridFormCallback.probe.stageOfInterest.Lock()
	defer topstarthalfwayarcshapegridFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TopStartHalfwayArcShapeGridFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	topstarthalfwayarcshapegridFormCallback.probe.formStage.Checkout()

	if topstarthalfwayarcshapegridFormCallback.topstarthalfwayarcshapegrid == nil {
		topstarthalfwayarcshapegridFormCallback.topstarthalfwayarcshapegrid = new(models.TopStartHalfwayArcShapeGrid).Stage(topstarthalfwayarcshapegridFormCallback.probe.stageOfInterest)
	}
	topstarthalfwayarcshapegrid_ := topstarthalfwayarcshapegridFormCallback.topstarthalfwayarcshapegrid
	_ = topstarthalfwayarcshapegrid_

	for _, formDiv := range topstarthalfwayarcshapegridFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(topstarthalfwayarcshapegrid_.Name), formDiv)
		case "TopStartHalfwayArcShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.TopStartHalfwayArcShape](topstarthalfwayarcshapegridFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.TopStartHalfwayArcShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.TopStartHalfwayArcShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					topstarthalfwayarcshapegridFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.TopStartHalfwayArcShape](topstarthalfwayarcshapegridFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			topstarthalfwayarcshapegrid_.TopStartHalfwayArcShapes = instanceSlice
			topstarthalfwayarcshapegridFormCallback.probe.UpdateSliceOfPointersCallback(topstarthalfwayarcshapegrid_, "TopStartHalfwayArcShapes", &topstarthalfwayarcshapegrid_.TopStartHalfwayArcShapes)

		}
	}

	// manage the suppress operation
	if topstarthalfwayarcshapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstarthalfwayarcshapegrid_.Unstage(topstarthalfwayarcshapegridFormCallback.probe.stageOfInterest)
	}

	topstarthalfwayarcshapegridFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TopStartHalfwayArcShapeGrid](
		topstarthalfwayarcshapegridFormCallback.probe,
	)

	// display a new form by reset the form stage
	if topstarthalfwayarcshapegridFormCallback.CreationMode || topstarthalfwayarcshapegridFormCallback.formGroup.HasSuppressButtonBeenPressed {
		topstarthalfwayarcshapegridFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(topstarthalfwayarcshapegridFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TopStartHalfwayArcShapeGridFormCallback(
			nil,
			topstarthalfwayarcshapegridFormCallback.probe,
			newFormGroup,
		)
		topstarthalfwayarcshapegrid := new(models.TopStartHalfwayArcShapeGrid)
		FillUpForm(topstarthalfwayarcshapegrid, newFormGroup, topstarthalfwayarcshapegridFormCallback.probe)
		topstarthalfwayarcshapegridFormCallback.probe.formStage.Commit()
	}

	topstarthalfwayarcshapegridFormCallback.probe.ux_tree()
}
func __gong__New__TorusStackShapeFormCallback(
	torusstackshape *models.TorusStackShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (torusstackshapeFormCallback *TorusStackShapeFormCallback) {
	torusstackshapeFormCallback = new(TorusStackShapeFormCallback)
	torusstackshapeFormCallback.probe = probe
	torusstackshapeFormCallback.torusstackshape = torusstackshape
	torusstackshapeFormCallback.formGroup = formGroup

	torusstackshapeFormCallback.CreationMode = (torusstackshape == nil)

	return
}

type TorusStackShapeFormCallback struct {
	torusstackshape *models.TorusStackShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (torusstackshapeFormCallback *TorusStackShapeFormCallback) OnSave() {
	torusstackshapeFormCallback.probe.stageOfInterest.Lock()
	defer torusstackshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("TorusStackShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	torusstackshapeFormCallback.probe.formStage.Checkout()

	if torusstackshapeFormCallback.torusstackshape == nil {
		torusstackshapeFormCallback.torusstackshape = new(models.TorusStackShape).Stage(torusstackshapeFormCallback.probe.stageOfInterest)
	}
	torusstackshape_ := torusstackshapeFormCallback.torusstackshape
	_ = torusstackshape_

	for _, formDiv := range torusstackshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(torusstackshape_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if torusstackshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		torusstackshape_.Unstage(torusstackshapeFormCallback.probe.stageOfInterest)
	}

	torusstackshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.TorusStackShape](
		torusstackshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if torusstackshapeFormCallback.CreationMode || torusstackshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		torusstackshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(torusstackshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TorusStackShapeFormCallback(
			nil,
			torusstackshapeFormCallback.probe,
			newFormGroup,
		)
		torusstackshape := new(models.TorusStackShape)
		FillUpForm(torusstackshape, newFormGroup, torusstackshapeFormCallback.probe)
		torusstackshapeFormCallback.probe.formStage.Commit()
	}

	torusstackshapeFormCallback.probe.ux_tree()
}
func __gong__New__VerticalTorusStackShapeFormCallback(
	verticaltorusstackshape *models.VerticalTorusStackShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (verticaltorusstackshapeFormCallback *VerticalTorusStackShapeFormCallback) {
	verticaltorusstackshapeFormCallback = new(VerticalTorusStackShapeFormCallback)
	verticaltorusstackshapeFormCallback.probe = probe
	verticaltorusstackshapeFormCallback.verticaltorusstackshape = verticaltorusstackshape
	verticaltorusstackshapeFormCallback.formGroup = formGroup

	verticaltorusstackshapeFormCallback.CreationMode = (verticaltorusstackshape == nil)

	return
}

type VerticalTorusStackShapeFormCallback struct {
	verticaltorusstackshape *models.VerticalTorusStackShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (verticaltorusstackshapeFormCallback *VerticalTorusStackShapeFormCallback) OnSave() {
	verticaltorusstackshapeFormCallback.probe.stageOfInterest.Lock()
	defer verticaltorusstackshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("VerticalTorusStackShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	verticaltorusstackshapeFormCallback.probe.formStage.Checkout()

	if verticaltorusstackshapeFormCallback.verticaltorusstackshape == nil {
		verticaltorusstackshapeFormCallback.verticaltorusstackshape = new(models.VerticalTorusStackShape).Stage(verticaltorusstackshapeFormCallback.probe.stageOfInterest)
	}
	verticaltorusstackshape_ := verticaltorusstackshapeFormCallback.verticaltorusstackshape
	_ = verticaltorusstackshape_

	for _, formDiv := range verticaltorusstackshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(verticaltorusstackshape_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if verticaltorusstackshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		verticaltorusstackshape_.Unstage(verticaltorusstackshapeFormCallback.probe.stageOfInterest)
	}

	verticaltorusstackshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.VerticalTorusStackShape](
		verticaltorusstackshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if verticaltorusstackshapeFormCallback.CreationMode || verticaltorusstackshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		verticaltorusstackshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(verticaltorusstackshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__VerticalTorusStackShapeFormCallback(
			nil,
			verticaltorusstackshapeFormCallback.probe,
			newFormGroup,
		)
		verticaltorusstackshape := new(models.VerticalTorusStackShape)
		FillUpForm(verticaltorusstackshape, newFormGroup, verticaltorusstackshapeFormCallback.probe)
		verticaltorusstackshapeFormCallback.probe.formStage.Commit()
	}

	verticaltorusstackshapeFormCallback.probe.ux_tree()
}
