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
		case "IsHidden":
			FormDivBasicFieldToField(&(axesshape_.IsHidden), formDiv)
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
		case "IsHidden":
			FormDivBasicFieldToField(&(circlegridshape_.IsHidden), formDiv)
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
		case "IsHidden":
			FormDivBasicFieldToField(&(explanationtextshape_.IsHidden), formDiv)
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
		case "IsHidden":
			FormDivBasicFieldToField(&(gridpathshape_.IsHidden), formDiv)
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
		case "IsHidden":
			FormDivBasicFieldToField(&(nextcircleshape_.IsHidden), formDiv)
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
		case "IsHidden":
			FormDivBasicFieldToField(&(plantcircumferenceshape_.IsHidden), formDiv)
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
		case "AxesShape":
			FormDivSelectFieldToField(&(plantdiagram_.AxesShape), plantdiagramFormCallback.probe.stageOfInterest, formDiv)
		case "ReferenceRhombus":
			FormDivSelectFieldToField(&(plantdiagram_.ReferenceRhombus), plantdiagramFormCallback.probe.stageOfInterest, formDiv)
		case "PlantCircumferenceShape":
			FormDivSelectFieldToField(&(plantdiagram_.PlantCircumferenceShape), plantdiagramFormCallback.probe.stageOfInterest, formDiv)
		case "GridPathShape":
			FormDivSelectFieldToField(&(plantdiagram_.GridPathShape), plantdiagramFormCallback.probe.stageOfInterest, formDiv)
		case "RhombusGridShape":
			FormDivSelectFieldToField(&(plantdiagram_.RhombusGridShape), plantdiagramFormCallback.probe.stageOfInterest, formDiv)
		case "ExplanationTextShape":
			FormDivSelectFieldToField(&(plantdiagram_.ExplanationTextShape), plantdiagramFormCallback.probe.stageOfInterest, formDiv)
		case "RotatedReferenceRhombus":
			FormDivSelectFieldToField(&(plantdiagram_.RotatedReferenceRhombus), plantdiagramFormCallback.probe.stageOfInterest, formDiv)
		case "RotatedPlantCircumferenceShape":
			FormDivSelectFieldToField(&(plantdiagram_.RotatedPlantCircumferenceShape), plantdiagramFormCallback.probe.stageOfInterest, formDiv)
		case "RotatedGridPathShape":
			FormDivSelectFieldToField(&(plantdiagram_.RotatedGridPathShape), plantdiagramFormCallback.probe.stageOfInterest, formDiv)
		case "RotatedRhombusGridShape":
			FormDivSelectFieldToField(&(plantdiagram_.RotatedRhombusGridShape), plantdiagramFormCallback.probe.stageOfInterest, formDiv)
		case "GrowthPathRhombusGridShape":
			FormDivSelectFieldToField(&(plantdiagram_.GrowthPathRhombusGridShape), plantdiagramFormCallback.probe.stageOfInterest, formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(plantdiagram_.IsChecked), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(plantdiagram_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(plantdiagram_.IsExpanded), formDiv)
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
func __gong__New__RhombusGridShapeFormCallback(
	rhombusgridshape *models.RhombusGridShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (rhombusgridshapeFormCallback *RhombusGridShapeFormCallback) {
	rhombusgridshapeFormCallback = new(RhombusGridShapeFormCallback)
	rhombusgridshapeFormCallback.probe = probe
	rhombusgridshapeFormCallback.rhombusgridshape = rhombusgridshape
	rhombusgridshapeFormCallback.formGroup = formGroup

	rhombusgridshapeFormCallback.CreationMode = (rhombusgridshape == nil)

	return
}

type RhombusGridShapeFormCallback struct {
	rhombusgridshape *models.RhombusGridShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (rhombusgridshapeFormCallback *RhombusGridShapeFormCallback) OnSave() {
	rhombusgridshapeFormCallback.probe.stageOfInterest.Lock()
	defer rhombusgridshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("RhombusGridShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rhombusgridshapeFormCallback.probe.formStage.Checkout()

	if rhombusgridshapeFormCallback.rhombusgridshape == nil {
		rhombusgridshapeFormCallback.rhombusgridshape = new(models.RhombusGridShape).Stage(rhombusgridshapeFormCallback.probe.stageOfInterest)
	}
	rhombusgridshape_ := rhombusgridshapeFormCallback.rhombusgridshape
	_ = rhombusgridshape_

	for _, formDiv := range rhombusgridshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rhombusgridshape_.Name), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(rhombusgridshape_.IsHidden), formDiv)
		case "RhombusShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.RhombusShape](rhombusgridshapeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.RhombusShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.RhombusShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					rhombusgridshapeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.RhombusShape](rhombusgridshapeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			rhombusgridshape_.RhombusShapes = instanceSlice
			rhombusgridshapeFormCallback.probe.UpdateSliceOfPointersCallback(rhombusgridshape_, "RhombusShapes", &rhombusgridshape_.RhombusShapes)

		}
	}

	// manage the suppress operation
	if rhombusgridshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rhombusgridshape_.Unstage(rhombusgridshapeFormCallback.probe.stageOfInterest)
	}

	rhombusgridshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.RhombusGridShape](
		rhombusgridshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if rhombusgridshapeFormCallback.CreationMode || rhombusgridshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		rhombusgridshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(rhombusgridshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RhombusGridShapeFormCallback(
			nil,
			rhombusgridshapeFormCallback.probe,
			newFormGroup,
		)
		rhombusgridshape := new(models.RhombusGridShape)
		FillUpForm(rhombusgridshape, newFormGroup, rhombusgridshapeFormCallback.probe)
		rhombusgridshapeFormCallback.probe.formStage.Commit()
	}

	rhombusgridshapeFormCallback.probe.ux_tree()
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
		case "IsHidden":
			FormDivBasicFieldToField(&(rhombusshape_.IsHidden), formDiv)
		case "RhombusGridShape:RhombusShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the RhombusGridShape instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target RhombusGridShape instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.RhombusGridShape](rhombusshapeFormCallback.probe.stageOfInterest)
			targetRhombusGridShapeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetRhombusGridShapeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all RhombusGridShape instances and update their RhombusShapes slice
			for _rhombusgridshape := range *models.GetGongstructInstancesSetFromPointerType[*models.RhombusGridShape](rhombusshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(rhombusshapeFormCallback.probe.stageOfInterest, _rhombusgridshape)
				
				// if RhombusGridShape is selected
				if targetRhombusGridShapeIDs[id] {
					// ensure rhombusshape_ is in _rhombusgridshape.RhombusShapes
					found := false
					for _, _b := range _rhombusgridshape.RhombusShapes {
						if _b == rhombusshape_ {
							found = true
							break
						}
					}
					if !found {
						_rhombusgridshape.RhombusShapes = append(_rhombusgridshape.RhombusShapes, rhombusshape_)
						rhombusshapeFormCallback.probe.UpdateSliceOfPointersCallback(_rhombusgridshape, "RhombusShapes", &_rhombusgridshape.RhombusShapes)
					}
				} else {
					// ensure rhombusshape_ is NOT in _rhombusgridshape.RhombusShapes
					idx := slices.Index(_rhombusgridshape.RhombusShapes, rhombusshape_)
					if idx != -1 {
						_rhombusgridshape.RhombusShapes = slices.Delete(_rhombusgridshape.RhombusShapes, idx, idx+1)
						rhombusshapeFormCallback.probe.UpdateSliceOfPointersCallback(_rhombusgridshape, "RhombusShapes", &_rhombusgridshape.RhombusShapes)
					}
				}
			}
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
