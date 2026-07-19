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
		case "RadialThickness":
			FormDivBasicFieldToField(&(plant_.RadialThickness), formDiv)
		case "RhombusSideLength":
			FormDivBasicFieldToField(&(plant_.RhombusSideLength), formDiv)
		case "CuttedStackFloorHeight":
			FormDivBasicFieldToField(&(plant_.CuttedStackFloorHeight), formDiv)
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
		case "StackOfGrowthCurve":
			FormDivSelectFieldToField(&(plant_.StackOfGrowthCurve), plantFormCallback.probe.stageOfInterest, formDiv)
		case "TopStackOfGrowthCurve":
			FormDivSelectFieldToField(&(plant_.TopStackOfGrowthCurve), plantFormCallback.probe.stageOfInterest, formDiv)
		case "ShiftedLeftStackOfGrowthCurve":
			FormDivSelectFieldToField(&(plant_.ShiftedLeftStackOfGrowthCurve), plantFormCallback.probe.stageOfInterest, formDiv)
		case "ShiftedLeftStackOfNormalVector":
			FormDivSelectFieldToField(&(plant_.ShiftedLeftStackOfNormalVector), plantFormCallback.probe.stageOfInterest, formDiv)
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
		case "IsHiddenTorusStackShape":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenTorusStackShape), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(plantdiagram_.IsChecked), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(plantdiagram_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(plantdiagram_.IsExpanded), formDiv)
		case "Rendered3DShape":
			FormDivSelectFieldToField(&(plantdiagram_.Rendered3DShape), plantdiagramFormCallback.probe.stageOfInterest, formDiv)
		case "TorusStackShape":
			FormDivSelectFieldToField(&(plantdiagram_.TorusStackShape), plantdiagramFormCallback.probe.stageOfInterest, formDiv)
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
