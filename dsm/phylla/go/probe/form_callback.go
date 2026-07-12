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
		case "RhombusSideLength":
			FormDivBasicFieldToField(&(plant_.RhombusSideLength), formDiv)
		case "ShiftToNearestCircle":
			FormDivBasicFieldToField(&(plant_.ShiftToNearestCircle), formDiv)
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
		case "GrowthCurveBezierShapeGrid":
			FormDivSelectFieldToField(&(plant_.GrowthCurveBezierShapeGrid), plantFormCallback.probe.stageOfInterest, formDiv)
		case "StackOfGrowthCurve":
			FormDivSelectFieldToField(&(plant_.StackOfGrowthCurve), plantFormCallback.probe.stageOfInterest, formDiv)
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
		case "IsHiddenGrowthCurveBezierShapeGrid":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenGrowthCurveBezierShapeGrid), formDiv)
		case "IsHiddenStackOfGrowthCurve":
			FormDivBasicFieldToField(&(plantdiagram_.IsHiddenStackOfGrowthCurve), formDiv)
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
