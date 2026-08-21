// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/floss/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__CompareAnalysisFormCallback(
	compareanalysis *models.CompareAnalysis,
	probe *Probe,
	formGroup *form.FormGroup,
) (compareanalysisFormCallback *CompareAnalysisFormCallback) {
	compareanalysisFormCallback = new(CompareAnalysisFormCallback)
	compareanalysisFormCallback.probe = probe
	compareanalysisFormCallback.compareanalysis = compareanalysis
	compareanalysisFormCallback.formGroup = formGroup

	compareanalysisFormCallback.CreationMode = (compareanalysis == nil)

	return
}

type CompareAnalysisFormCallback struct {
	compareanalysis *models.CompareAnalysis

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (compareanalysisFormCallback *CompareAnalysisFormCallback) OnSave() {
	compareanalysisFormCallback.probe.stageOfInterest.Lock()
	defer compareanalysisFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("CompareAnalysisFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	compareanalysisFormCallback.probe.formStage.Checkout()

	if compareanalysisFormCallback.compareanalysis == nil {
		compareanalysisFormCallback.compareanalysis = new(models.CompareAnalysis).Stage(compareanalysisFormCallback.probe.stageOfInterest)
	}
	compareanalysis_ := compareanalysisFormCallback.compareanalysis
	_ = compareanalysis_

	for _, formDiv := range compareanalysisFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(compareanalysis_.Name), formDiv)
		case "FromSystem":
			FormDivSelectFieldToField(&(compareanalysis_.FromSystem), compareanalysisFormCallback.probe.stageOfInterest, formDiv)
		case "ToSystem":
			FormDivSelectFieldToField(&(compareanalysis_.ToSystem), compareanalysisFormCallback.probe.stageOfInterest, formDiv)
		case "Alpha":
			FormDivBasicFieldToField(&(compareanalysis_.Alpha), formDiv)
		case "Beta":
			FormDivBasicFieldToField(&(compareanalysis_.Beta), formDiv)
		case "DiagramFlossEquations":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DiagramFlossEquation](compareanalysisFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DiagramFlossEquation, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DiagramFlossEquation)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					compareanalysisFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramFlossEquation](compareanalysisFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			compareanalysis_.DiagramFlossEquations = instanceSlice
			compareanalysisFormCallback.probe.UpdateSliceOfPointersCallback(compareanalysis_, "DiagramFlossEquations", &compareanalysis_.DiagramFlossEquations)

		case "DiagramFlossEquationsWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DiagramFlossEquation](compareanalysisFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DiagramFlossEquation, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DiagramFlossEquation)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					compareanalysisFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramFlossEquation](compareanalysisFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			compareanalysis_.DiagramFlossEquationsWhoseNodeIsExpanded = instanceSlice
			compareanalysisFormCallback.probe.UpdateSliceOfPointersCallback(compareanalysis_, "DiagramFlossEquationsWhoseNodeIsExpanded", &compareanalysis_.DiagramFlossEquationsWhoseNodeIsExpanded)

		case "ComputedPrefix":
			FormDivBasicFieldToField(&(compareanalysis_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(compareanalysis_.IsExpanded), formDiv)
		case "Library:RootCompareAnalysis":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](compareanalysisFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their RootCompareAnalysis slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](compareanalysisFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(compareanalysisFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure compareanalysis_ is in _library.RootCompareAnalysis
					found := false
					for _, _b := range _library.RootCompareAnalysis {
						if _b == compareanalysis_ {
							found = true
							break
						}
					}
					if !found {
						_library.RootCompareAnalysis = append(_library.RootCompareAnalysis, compareanalysis_)
						compareanalysisFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootCompareAnalysis", &_library.RootCompareAnalysis)
					}
				} else {
					// ensure compareanalysis_ is NOT in _library.RootCompareAnalysis
					idx := slices.Index(_library.RootCompareAnalysis, compareanalysis_)
					if idx != -1 {
						_library.RootCompareAnalysis = slices.Delete(_library.RootCompareAnalysis, idx, idx+1)
						compareanalysisFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootCompareAnalysis", &_library.RootCompareAnalysis)
					}
				}
			}
		case "Library:CompareAnalysisWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](compareanalysisFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their CompareAnalysisWhoseNodeIsExpanded slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](compareanalysisFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(compareanalysisFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure compareanalysis_ is in _library.CompareAnalysisWhoseNodeIsExpanded
					found := false
					for _, _b := range _library.CompareAnalysisWhoseNodeIsExpanded {
						if _b == compareanalysis_ {
							found = true
							break
						}
					}
					if !found {
						_library.CompareAnalysisWhoseNodeIsExpanded = append(_library.CompareAnalysisWhoseNodeIsExpanded, compareanalysis_)
						compareanalysisFormCallback.probe.UpdateSliceOfPointersCallback(_library, "CompareAnalysisWhoseNodeIsExpanded", &_library.CompareAnalysisWhoseNodeIsExpanded)
					}
				} else {
					// ensure compareanalysis_ is NOT in _library.CompareAnalysisWhoseNodeIsExpanded
					idx := slices.Index(_library.CompareAnalysisWhoseNodeIsExpanded, compareanalysis_)
					if idx != -1 {
						_library.CompareAnalysisWhoseNodeIsExpanded = slices.Delete(_library.CompareAnalysisWhoseNodeIsExpanded, idx, idx+1)
						compareanalysisFormCallback.probe.UpdateSliceOfPointersCallback(_library, "CompareAnalysisWhoseNodeIsExpanded", &_library.CompareAnalysisWhoseNodeIsExpanded)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if compareanalysisFormCallback.formGroup.HasSuppressButtonBeenPressed {
		compareanalysis_.Unstage(compareanalysisFormCallback.probe.stageOfInterest)
	}

	compareanalysisFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.CompareAnalysis](
		compareanalysisFormCallback.probe,
	)

	// display a new form by reset the form stage
	if compareanalysisFormCallback.CreationMode || compareanalysisFormCallback.formGroup.HasSuppressButtonBeenPressed {
		compareanalysisFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(compareanalysisFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CompareAnalysisFormCallback(
			nil,
			compareanalysisFormCallback.probe,
			newFormGroup,
		)
		compareanalysis := new(models.CompareAnalysis)
		FillUpForm(compareanalysis, newFormGroup, compareanalysisFormCallback.probe)
		compareanalysisFormCallback.probe.formStage.Commit()
	}

	compareanalysisFormCallback.probe.ux_tree()
}
func __gong__New__ComplexityFormCallback(
	complexity *models.Complexity,
	probe *Probe,
	formGroup *form.FormGroup,
) (complexityFormCallback *ComplexityFormCallback) {
	complexityFormCallback = new(ComplexityFormCallback)
	complexityFormCallback.probe = probe
	complexityFormCallback.complexity = complexity
	complexityFormCallback.formGroup = formGroup

	complexityFormCallback.CreationMode = (complexity == nil)

	return
}

type ComplexityFormCallback struct {
	complexity *models.Complexity

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (complexityFormCallback *ComplexityFormCallback) OnSave() {
	complexityFormCallback.probe.stageOfInterest.Lock()
	defer complexityFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ComplexityFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	complexityFormCallback.probe.formStage.Checkout()

	if complexityFormCallback.complexity == nil {
		complexityFormCallback.complexity = new(models.Complexity).Stage(complexityFormCallback.probe.stageOfInterest)
	}
	complexity_ := complexityFormCallback.complexity
	_ = complexity_

	for _, formDiv := range complexityFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(complexity_.Name), formDiv)
		case "Strength":
			FormDivBasicFieldToField(&(complexity_.Strength), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(complexity_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(complexity_.IsExpanded), formDiv)
		case "DiagramFlossEquation:ComplexitysWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramFlossEquation instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramFlossEquation instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramFlossEquation](complexityFormCallback.probe.stageOfInterest)
			targetDiagramFlossEquationIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramFlossEquationIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramFlossEquation instances and update their ComplexitysWhoseNodeIsExpanded slice
			for _diagramflossequation := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramFlossEquation](complexityFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(complexityFormCallback.probe.stageOfInterest, _diagramflossequation)
				
				// if DiagramFlossEquation is selected
				if targetDiagramFlossEquationIDs[id] {
					// ensure complexity_ is in _diagramflossequation.ComplexitysWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramflossequation.ComplexitysWhoseNodeIsExpanded {
						if _b == complexity_ {
							found = true
							break
						}
					}
					if !found {
						_diagramflossequation.ComplexitysWhoseNodeIsExpanded = append(_diagramflossequation.ComplexitysWhoseNodeIsExpanded, complexity_)
						complexityFormCallback.probe.UpdateSliceOfPointersCallback(_diagramflossequation, "ComplexitysWhoseNodeIsExpanded", &_diagramflossequation.ComplexitysWhoseNodeIsExpanded)
					}
				} else {
					// ensure complexity_ is NOT in _diagramflossequation.ComplexitysWhoseNodeIsExpanded
					idx := slices.Index(_diagramflossequation.ComplexitysWhoseNodeIsExpanded, complexity_)
					if idx != -1 {
						_diagramflossequation.ComplexitysWhoseNodeIsExpanded = slices.Delete(_diagramflossequation.ComplexitysWhoseNodeIsExpanded, idx, idx+1)
						complexityFormCallback.probe.UpdateSliceOfPointersCallback(_diagramflossequation, "ComplexitysWhoseNodeIsExpanded", &_diagramflossequation.ComplexitysWhoseNodeIsExpanded)
					}
				}
			}
		case "Library:RootComplexitys":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](complexityFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their RootComplexitys slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](complexityFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(complexityFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure complexity_ is in _library.RootComplexitys
					found := false
					for _, _b := range _library.RootComplexitys {
						if _b == complexity_ {
							found = true
							break
						}
					}
					if !found {
						_library.RootComplexitys = append(_library.RootComplexitys, complexity_)
						complexityFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootComplexitys", &_library.RootComplexitys)
					}
				} else {
					// ensure complexity_ is NOT in _library.RootComplexitys
					idx := slices.Index(_library.RootComplexitys, complexity_)
					if idx != -1 {
						_library.RootComplexitys = slices.Delete(_library.RootComplexitys, idx, idx+1)
						complexityFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootComplexitys", &_library.RootComplexitys)
					}
				}
			}
		case "Library:ComplexitysWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](complexityFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their ComplexitysWhoseNodeIsExpanded slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](complexityFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(complexityFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure complexity_ is in _library.ComplexitysWhoseNodeIsExpanded
					found := false
					for _, _b := range _library.ComplexitysWhoseNodeIsExpanded {
						if _b == complexity_ {
							found = true
							break
						}
					}
					if !found {
						_library.ComplexitysWhoseNodeIsExpanded = append(_library.ComplexitysWhoseNodeIsExpanded, complexity_)
						complexityFormCallback.probe.UpdateSliceOfPointersCallback(_library, "ComplexitysWhoseNodeIsExpanded", &_library.ComplexitysWhoseNodeIsExpanded)
					}
				} else {
					// ensure complexity_ is NOT in _library.ComplexitysWhoseNodeIsExpanded
					idx := slices.Index(_library.ComplexitysWhoseNodeIsExpanded, complexity_)
					if idx != -1 {
						_library.ComplexitysWhoseNodeIsExpanded = slices.Delete(_library.ComplexitysWhoseNodeIsExpanded, idx, idx+1)
						complexityFormCallback.probe.UpdateSliceOfPointersCallback(_library, "ComplexitysWhoseNodeIsExpanded", &_library.ComplexitysWhoseNodeIsExpanded)
					}
				}
			}
		case "Note:Complexities":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Note instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Note instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Note](complexityFormCallback.probe.stageOfInterest)
			targetNoteIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetNoteIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Note instances and update their Complexities slice
			for _note := range *models.GetGongstructInstancesSetFromPointerType[*models.Note](complexityFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(complexityFormCallback.probe.stageOfInterest, _note)
				
				// if Note is selected
				if targetNoteIDs[id] {
					// ensure complexity_ is in _note.Complexities
					found := false
					for _, _b := range _note.Complexities {
						if _b == complexity_ {
							found = true
							break
						}
					}
					if !found {
						_note.Complexities = append(_note.Complexities, complexity_)
						complexityFormCallback.probe.UpdateSliceOfPointersCallback(_note, "Complexities", &_note.Complexities)
					}
				} else {
					// ensure complexity_ is NOT in _note.Complexities
					idx := slices.Index(_note.Complexities, complexity_)
					if idx != -1 {
						_note.Complexities = slices.Delete(_note.Complexities, idx, idx+1)
						complexityFormCallback.probe.UpdateSliceOfPointersCallback(_note, "Complexities", &_note.Complexities)
					}
				}
			}
		case "System:Complexities":
			// 1. Decode the AssociationStorage which contains the rowIDs of the System instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target System instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.System](complexityFormCallback.probe.stageOfInterest)
			targetSystemIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSystemIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all System instances and update their Complexities slice
			for _system := range *models.GetGongstructInstancesSetFromPointerType[*models.System](complexityFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(complexityFormCallback.probe.stageOfInterest, _system)
				
				// if System is selected
				if targetSystemIDs[id] {
					// ensure complexity_ is in _system.Complexities
					found := false
					for _, _b := range _system.Complexities {
						if _b == complexity_ {
							found = true
							break
						}
					}
					if !found {
						_system.Complexities = append(_system.Complexities, complexity_)
						complexityFormCallback.probe.UpdateSliceOfPointersCallback(_system, "Complexities", &_system.Complexities)
					}
				} else {
					// ensure complexity_ is NOT in _system.Complexities
					idx := slices.Index(_system.Complexities, complexity_)
					if idx != -1 {
						_system.Complexities = slices.Delete(_system.Complexities, idx, idx+1)
						complexityFormCallback.probe.UpdateSliceOfPointersCallback(_system, "Complexities", &_system.Complexities)
					}
				}
			}
		case "System:ComplexitysWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the System instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target System instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.System](complexityFormCallback.probe.stageOfInterest)
			targetSystemIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSystemIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all System instances and update their ComplexitysWhoseNodeIsExpanded slice
			for _system := range *models.GetGongstructInstancesSetFromPointerType[*models.System](complexityFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(complexityFormCallback.probe.stageOfInterest, _system)
				
				// if System is selected
				if targetSystemIDs[id] {
					// ensure complexity_ is in _system.ComplexitysWhoseNodeIsExpanded
					found := false
					for _, _b := range _system.ComplexitysWhoseNodeIsExpanded {
						if _b == complexity_ {
							found = true
							break
						}
					}
					if !found {
						_system.ComplexitysWhoseNodeIsExpanded = append(_system.ComplexitysWhoseNodeIsExpanded, complexity_)
						complexityFormCallback.probe.UpdateSliceOfPointersCallback(_system, "ComplexitysWhoseNodeIsExpanded", &_system.ComplexitysWhoseNodeIsExpanded)
					}
				} else {
					// ensure complexity_ is NOT in _system.ComplexitysWhoseNodeIsExpanded
					idx := slices.Index(_system.ComplexitysWhoseNodeIsExpanded, complexity_)
					if idx != -1 {
						_system.ComplexitysWhoseNodeIsExpanded = slices.Delete(_system.ComplexitysWhoseNodeIsExpanded, idx, idx+1)
						complexityFormCallback.probe.UpdateSliceOfPointersCallback(_system, "ComplexitysWhoseNodeIsExpanded", &_system.ComplexitysWhoseNodeIsExpanded)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if complexityFormCallback.formGroup.HasSuppressButtonBeenPressed {
		complexity_.Unstage(complexityFormCallback.probe.stageOfInterest)
	}

	complexityFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Complexity](
		complexityFormCallback.probe,
	)

	// display a new form by reset the form stage
	if complexityFormCallback.CreationMode || complexityFormCallback.formGroup.HasSuppressButtonBeenPressed {
		complexityFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(complexityFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ComplexityFormCallback(
			nil,
			complexityFormCallback.probe,
			newFormGroup,
		)
		complexity := new(models.Complexity)
		FillUpForm(complexity, newFormGroup, complexityFormCallback.probe)
		complexityFormCallback.probe.formStage.Commit()
	}

	complexityFormCallback.probe.ux_tree()
}
func __gong__New__DiagramFlossEquationFormCallback(
	diagramflossequation *models.DiagramFlossEquation,
	probe *Probe,
	formGroup *form.FormGroup,
) (diagramflossequationFormCallback *DiagramFlossEquationFormCallback) {
	diagramflossequationFormCallback = new(DiagramFlossEquationFormCallback)
	diagramflossequationFormCallback.probe = probe
	diagramflossequationFormCallback.diagramflossequation = diagramflossequation
	diagramflossequationFormCallback.formGroup = formGroup

	diagramflossequationFormCallback.CreationMode = (diagramflossequation == nil)

	return
}

type DiagramFlossEquationFormCallback struct {
	diagramflossequation *models.DiagramFlossEquation

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (diagramflossequationFormCallback *DiagramFlossEquationFormCallback) OnSave() {
	diagramflossequationFormCallback.probe.stageOfInterest.Lock()
	defer diagramflossequationFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DiagramFlossEquationFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	diagramflossequationFormCallback.probe.formStage.Checkout()

	if diagramflossequationFormCallback.diagramflossequation == nil {
		diagramflossequationFormCallback.diagramflossequation = new(models.DiagramFlossEquation).Stage(diagramflossequationFormCallback.probe.stageOfInterest)
	}
	diagramflossequation_ := diagramflossequationFormCallback.diagramflossequation
	_ = diagramflossequation_

	for _, formDiv := range diagramflossequationFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(diagramflossequation_.Name), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(diagramflossequation_.Description), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(diagramflossequation_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(diagramflossequation_.IsExpanded), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(diagramflossequation_.IsChecked), formDiv)
		case "IsEditable_":
			FormDivBasicFieldToField(&(diagramflossequation_.IsEditable_), formDiv)
		case "AreQuantitativeElementsVisible":
			FormDivBasicFieldToField(&(diagramflossequation_.AreQuantitativeElementsVisible), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(diagramflossequation_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(diagramflossequation_.Height), formDiv)
		case "Scale":
			FormDivBasicFieldToField(&(diagramflossequation_.Scale), formDiv)
		case "DefaultBoxWidth":
			FormDivBasicFieldToField(&(diagramflossequation_.DefaultBoxWidth), formDiv)
		case "DefaultBoxHeigth":
			FormDivBasicFieldToField(&(diagramflossequation_.DefaultBoxHeigth), formDiv)
		case "Note_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.NoteShape](diagramflossequationFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.NoteShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.NoteShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramflossequationFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.NoteShape](diagramflossequationFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramflossequation_.Note_Shapes = instanceSlice
			diagramflossequationFormCallback.probe.UpdateSliceOfPointersCallback(diagramflossequation_, "Note_Shapes", &diagramflossequation_.Note_Shapes)

		case "NoteComplexityShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.NoteComplexityShape](diagramflossequationFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.NoteComplexityShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.NoteComplexityShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramflossequationFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.NoteComplexityShape](diagramflossequationFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramflossequation_.NoteComplexityShapes = instanceSlice
			diagramflossequationFormCallback.probe.UpdateSliceOfPointersCallback(diagramflossequation_, "NoteComplexityShapes", &diagramflossequation_.NoteComplexityShapes)

		case "NotePerformanceShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.NotePerformanceShape](diagramflossequationFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.NotePerformanceShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.NotePerformanceShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramflossequationFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.NotePerformanceShape](diagramflossequationFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramflossequation_.NotePerformanceShapes = instanceSlice
			diagramflossequationFormCallback.probe.UpdateSliceOfPointersCallback(diagramflossequation_, "NotePerformanceShapes", &diagramflossequation_.NotePerformanceShapes)

		case "NoteEffortShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.NoteEffortShape](diagramflossequationFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.NoteEffortShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.NoteEffortShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramflossequationFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.NoteEffortShape](diagramflossequationFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramflossequation_.NoteEffortShapes = instanceSlice
			diagramflossequationFormCallback.probe.UpdateSliceOfPointersCallback(diagramflossequation_, "NoteEffortShapes", &diagramflossequation_.NoteEffortShapes)

		case "IsNotesNodeExpanded":
			FormDivBasicFieldToField(&(diagramflossequation_.IsNotesNodeExpanded), formDiv)
		case "NotesWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Note](diagramflossequationFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Note, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Note)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramflossequationFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Note](diagramflossequationFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramflossequation_.NotesWhoseNodeIsExpanded = instanceSlice
			diagramflossequationFormCallback.probe.UpdateSliceOfPointersCallback(diagramflossequation_, "NotesWhoseNodeIsExpanded", &diagramflossequation_.NotesWhoseNodeIsExpanded)

		case "IsComplexitysNodeExpanded":
			FormDivBasicFieldToField(&(diagramflossequation_.IsComplexitysNodeExpanded), formDiv)
		case "ComplexitysWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Complexity](diagramflossequationFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Complexity, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Complexity)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramflossequationFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Complexity](diagramflossequationFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramflossequation_.ComplexitysWhoseNodeIsExpanded = instanceSlice
			diagramflossequationFormCallback.probe.UpdateSliceOfPointersCallback(diagramflossequation_, "ComplexitysWhoseNodeIsExpanded", &diagramflossequation_.ComplexitysWhoseNodeIsExpanded)

		case "IsPerformancesNodeExpanded":
			FormDivBasicFieldToField(&(diagramflossequation_.IsPerformancesNodeExpanded), formDiv)
		case "PerformancesWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Performance](diagramflossequationFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Performance, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Performance)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramflossequationFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Performance](diagramflossequationFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramflossequation_.PerformancesWhoseNodeIsExpanded = instanceSlice
			diagramflossequationFormCallback.probe.UpdateSliceOfPointersCallback(diagramflossequation_, "PerformancesWhoseNodeIsExpanded", &diagramflossequation_.PerformancesWhoseNodeIsExpanded)

		case "IsEffortsNodeExpanded":
			FormDivBasicFieldToField(&(diagramflossequation_.IsEffortsNodeExpanded), formDiv)
		case "EffortsWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Effort](diagramflossequationFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Effort, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Effort)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramflossequationFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Effort](diagramflossequationFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramflossequation_.EffortsWhoseNodeIsExpanded = instanceSlice
			diagramflossequationFormCallback.probe.UpdateSliceOfPointersCallback(diagramflossequation_, "EffortsWhoseNodeIsExpanded", &diagramflossequation_.EffortsWhoseNodeIsExpanded)

		case "CompareAnalysis:DiagramFlossEquations":
			// 1. Decode the AssociationStorage which contains the rowIDs of the CompareAnalysis instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target CompareAnalysis instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.CompareAnalysis](diagramflossequationFormCallback.probe.stageOfInterest)
			targetCompareAnalysisIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetCompareAnalysisIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all CompareAnalysis instances and update their DiagramFlossEquations slice
			for _compareanalysis := range *models.GetGongstructInstancesSetFromPointerType[*models.CompareAnalysis](diagramflossequationFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(diagramflossequationFormCallback.probe.stageOfInterest, _compareanalysis)
				
				// if CompareAnalysis is selected
				if targetCompareAnalysisIDs[id] {
					// ensure diagramflossequation_ is in _compareanalysis.DiagramFlossEquations
					found := false
					for _, _b := range _compareanalysis.DiagramFlossEquations {
						if _b == diagramflossequation_ {
							found = true
							break
						}
					}
					if !found {
						_compareanalysis.DiagramFlossEquations = append(_compareanalysis.DiagramFlossEquations, diagramflossequation_)
						diagramflossequationFormCallback.probe.UpdateSliceOfPointersCallback(_compareanalysis, "DiagramFlossEquations", &_compareanalysis.DiagramFlossEquations)
					}
				} else {
					// ensure diagramflossequation_ is NOT in _compareanalysis.DiagramFlossEquations
					idx := slices.Index(_compareanalysis.DiagramFlossEquations, diagramflossequation_)
					if idx != -1 {
						_compareanalysis.DiagramFlossEquations = slices.Delete(_compareanalysis.DiagramFlossEquations, idx, idx+1)
						diagramflossequationFormCallback.probe.UpdateSliceOfPointersCallback(_compareanalysis, "DiagramFlossEquations", &_compareanalysis.DiagramFlossEquations)
					}
				}
			}
		case "CompareAnalysis:DiagramFlossEquationsWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the CompareAnalysis instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target CompareAnalysis instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.CompareAnalysis](diagramflossequationFormCallback.probe.stageOfInterest)
			targetCompareAnalysisIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetCompareAnalysisIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all CompareAnalysis instances and update their DiagramFlossEquationsWhoseNodeIsExpanded slice
			for _compareanalysis := range *models.GetGongstructInstancesSetFromPointerType[*models.CompareAnalysis](diagramflossequationFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(diagramflossequationFormCallback.probe.stageOfInterest, _compareanalysis)
				
				// if CompareAnalysis is selected
				if targetCompareAnalysisIDs[id] {
					// ensure diagramflossequation_ is in _compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded
					found := false
					for _, _b := range _compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded {
						if _b == diagramflossequation_ {
							found = true
							break
						}
					}
					if !found {
						_compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded = append(_compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded, diagramflossequation_)
						diagramflossequationFormCallback.probe.UpdateSliceOfPointersCallback(_compareanalysis, "DiagramFlossEquationsWhoseNodeIsExpanded", &_compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded)
					}
				} else {
					// ensure diagramflossequation_ is NOT in _compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded
					idx := slices.Index(_compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded, diagramflossequation_)
					if idx != -1 {
						_compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded = slices.Delete(_compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded, idx, idx+1)
						diagramflossequationFormCallback.probe.UpdateSliceOfPointersCallback(_compareanalysis, "DiagramFlossEquationsWhoseNodeIsExpanded", &_compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded)
					}
				}
			}
		case "System:DiagramFlossEquations":
			// 1. Decode the AssociationStorage which contains the rowIDs of the System instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target System instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.System](diagramflossequationFormCallback.probe.stageOfInterest)
			targetSystemIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSystemIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all System instances and update their DiagramFlossEquations slice
			for _system := range *models.GetGongstructInstancesSetFromPointerType[*models.System](diagramflossequationFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(diagramflossequationFormCallback.probe.stageOfInterest, _system)
				
				// if System is selected
				if targetSystemIDs[id] {
					// ensure diagramflossequation_ is in _system.DiagramFlossEquations
					found := false
					for _, _b := range _system.DiagramFlossEquations {
						if _b == diagramflossequation_ {
							found = true
							break
						}
					}
					if !found {
						_system.DiagramFlossEquations = append(_system.DiagramFlossEquations, diagramflossequation_)
						diagramflossequationFormCallback.probe.UpdateSliceOfPointersCallback(_system, "DiagramFlossEquations", &_system.DiagramFlossEquations)
					}
				} else {
					// ensure diagramflossequation_ is NOT in _system.DiagramFlossEquations
					idx := slices.Index(_system.DiagramFlossEquations, diagramflossequation_)
					if idx != -1 {
						_system.DiagramFlossEquations = slices.Delete(_system.DiagramFlossEquations, idx, idx+1)
						diagramflossequationFormCallback.probe.UpdateSliceOfPointersCallback(_system, "DiagramFlossEquations", &_system.DiagramFlossEquations)
					}
				}
			}
		case "System:DiagramFlossEquationsWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the System instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target System instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.System](diagramflossequationFormCallback.probe.stageOfInterest)
			targetSystemIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSystemIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all System instances and update their DiagramFlossEquationsWhoseNodeIsExpanded slice
			for _system := range *models.GetGongstructInstancesSetFromPointerType[*models.System](diagramflossequationFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(diagramflossequationFormCallback.probe.stageOfInterest, _system)
				
				// if System is selected
				if targetSystemIDs[id] {
					// ensure diagramflossequation_ is in _system.DiagramFlossEquationsWhoseNodeIsExpanded
					found := false
					for _, _b := range _system.DiagramFlossEquationsWhoseNodeIsExpanded {
						if _b == diagramflossequation_ {
							found = true
							break
						}
					}
					if !found {
						_system.DiagramFlossEquationsWhoseNodeIsExpanded = append(_system.DiagramFlossEquationsWhoseNodeIsExpanded, diagramflossequation_)
						diagramflossequationFormCallback.probe.UpdateSliceOfPointersCallback(_system, "DiagramFlossEquationsWhoseNodeIsExpanded", &_system.DiagramFlossEquationsWhoseNodeIsExpanded)
					}
				} else {
					// ensure diagramflossequation_ is NOT in _system.DiagramFlossEquationsWhoseNodeIsExpanded
					idx := slices.Index(_system.DiagramFlossEquationsWhoseNodeIsExpanded, diagramflossequation_)
					if idx != -1 {
						_system.DiagramFlossEquationsWhoseNodeIsExpanded = slices.Delete(_system.DiagramFlossEquationsWhoseNodeIsExpanded, idx, idx+1)
						diagramflossequationFormCallback.probe.UpdateSliceOfPointersCallback(_system, "DiagramFlossEquationsWhoseNodeIsExpanded", &_system.DiagramFlossEquationsWhoseNodeIsExpanded)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if diagramflossequationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		diagramflossequation_.Unstage(diagramflossequationFormCallback.probe.stageOfInterest)
	}

	diagramflossequationFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.DiagramFlossEquation](
		diagramflossequationFormCallback.probe,
	)

	// display a new form by reset the form stage
	if diagramflossequationFormCallback.CreationMode || diagramflossequationFormCallback.formGroup.HasSuppressButtonBeenPressed {
		diagramflossequationFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(diagramflossequationFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DiagramFlossEquationFormCallback(
			nil,
			diagramflossequationFormCallback.probe,
			newFormGroup,
		)
		diagramflossequation := new(models.DiagramFlossEquation)
		FillUpForm(diagramflossequation, newFormGroup, diagramflossequationFormCallback.probe)
		diagramflossequationFormCallback.probe.formStage.Commit()
	}

	diagramflossequationFormCallback.probe.ux_tree()
}
func __gong__New__EffortFormCallback(
	effort *models.Effort,
	probe *Probe,
	formGroup *form.FormGroup,
) (effortFormCallback *EffortFormCallback) {
	effortFormCallback = new(EffortFormCallback)
	effortFormCallback.probe = probe
	effortFormCallback.effort = effort
	effortFormCallback.formGroup = formGroup

	effortFormCallback.CreationMode = (effort == nil)

	return
}

type EffortFormCallback struct {
	effort *models.Effort

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (effortFormCallback *EffortFormCallback) OnSave() {
	effortFormCallback.probe.stageOfInterest.Lock()
	defer effortFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("EffortFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	effortFormCallback.probe.formStage.Checkout()

	if effortFormCallback.effort == nil {
		effortFormCallback.effort = new(models.Effort).Stage(effortFormCallback.probe.stageOfInterest)
	}
	effort_ := effortFormCallback.effort
	_ = effort_

	for _, formDiv := range effortFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(effort_.Name), formDiv)
		case "Strength":
			FormDivBasicFieldToField(&(effort_.Strength), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(effort_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(effort_.IsExpanded), formDiv)
		case "DiagramFlossEquation:EffortsWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramFlossEquation instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramFlossEquation instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramFlossEquation](effortFormCallback.probe.stageOfInterest)
			targetDiagramFlossEquationIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramFlossEquationIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramFlossEquation instances and update their EffortsWhoseNodeIsExpanded slice
			for _diagramflossequation := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramFlossEquation](effortFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(effortFormCallback.probe.stageOfInterest, _diagramflossequation)
				
				// if DiagramFlossEquation is selected
				if targetDiagramFlossEquationIDs[id] {
					// ensure effort_ is in _diagramflossequation.EffortsWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramflossequation.EffortsWhoseNodeIsExpanded {
						if _b == effort_ {
							found = true
							break
						}
					}
					if !found {
						_diagramflossequation.EffortsWhoseNodeIsExpanded = append(_diagramflossequation.EffortsWhoseNodeIsExpanded, effort_)
						effortFormCallback.probe.UpdateSliceOfPointersCallback(_diagramflossequation, "EffortsWhoseNodeIsExpanded", &_diagramflossequation.EffortsWhoseNodeIsExpanded)
					}
				} else {
					// ensure effort_ is NOT in _diagramflossequation.EffortsWhoseNodeIsExpanded
					idx := slices.Index(_diagramflossequation.EffortsWhoseNodeIsExpanded, effort_)
					if idx != -1 {
						_diagramflossequation.EffortsWhoseNodeIsExpanded = slices.Delete(_diagramflossequation.EffortsWhoseNodeIsExpanded, idx, idx+1)
						effortFormCallback.probe.UpdateSliceOfPointersCallback(_diagramflossequation, "EffortsWhoseNodeIsExpanded", &_diagramflossequation.EffortsWhoseNodeIsExpanded)
					}
				}
			}
		case "Library:RootEfforts":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](effortFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their RootEfforts slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](effortFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(effortFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure effort_ is in _library.RootEfforts
					found := false
					for _, _b := range _library.RootEfforts {
						if _b == effort_ {
							found = true
							break
						}
					}
					if !found {
						_library.RootEfforts = append(_library.RootEfforts, effort_)
						effortFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootEfforts", &_library.RootEfforts)
					}
				} else {
					// ensure effort_ is NOT in _library.RootEfforts
					idx := slices.Index(_library.RootEfforts, effort_)
					if idx != -1 {
						_library.RootEfforts = slices.Delete(_library.RootEfforts, idx, idx+1)
						effortFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootEfforts", &_library.RootEfforts)
					}
				}
			}
		case "Library:EffortsWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](effortFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their EffortsWhoseNodeIsExpanded slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](effortFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(effortFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure effort_ is in _library.EffortsWhoseNodeIsExpanded
					found := false
					for _, _b := range _library.EffortsWhoseNodeIsExpanded {
						if _b == effort_ {
							found = true
							break
						}
					}
					if !found {
						_library.EffortsWhoseNodeIsExpanded = append(_library.EffortsWhoseNodeIsExpanded, effort_)
						effortFormCallback.probe.UpdateSliceOfPointersCallback(_library, "EffortsWhoseNodeIsExpanded", &_library.EffortsWhoseNodeIsExpanded)
					}
				} else {
					// ensure effort_ is NOT in _library.EffortsWhoseNodeIsExpanded
					idx := slices.Index(_library.EffortsWhoseNodeIsExpanded, effort_)
					if idx != -1 {
						_library.EffortsWhoseNodeIsExpanded = slices.Delete(_library.EffortsWhoseNodeIsExpanded, idx, idx+1)
						effortFormCallback.probe.UpdateSliceOfPointersCallback(_library, "EffortsWhoseNodeIsExpanded", &_library.EffortsWhoseNodeIsExpanded)
					}
				}
			}
		case "Note:Efforts":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Note instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Note instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Note](effortFormCallback.probe.stageOfInterest)
			targetNoteIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetNoteIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Note instances and update their Efforts slice
			for _note := range *models.GetGongstructInstancesSetFromPointerType[*models.Note](effortFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(effortFormCallback.probe.stageOfInterest, _note)
				
				// if Note is selected
				if targetNoteIDs[id] {
					// ensure effort_ is in _note.Efforts
					found := false
					for _, _b := range _note.Efforts {
						if _b == effort_ {
							found = true
							break
						}
					}
					if !found {
						_note.Efforts = append(_note.Efforts, effort_)
						effortFormCallback.probe.UpdateSliceOfPointersCallback(_note, "Efforts", &_note.Efforts)
					}
				} else {
					// ensure effort_ is NOT in _note.Efforts
					idx := slices.Index(_note.Efforts, effort_)
					if idx != -1 {
						_note.Efforts = slices.Delete(_note.Efforts, idx, idx+1)
						effortFormCallback.probe.UpdateSliceOfPointersCallback(_note, "Efforts", &_note.Efforts)
					}
				}
			}
		case "System:Efforts":
			// 1. Decode the AssociationStorage which contains the rowIDs of the System instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target System instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.System](effortFormCallback.probe.stageOfInterest)
			targetSystemIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSystemIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all System instances and update their Efforts slice
			for _system := range *models.GetGongstructInstancesSetFromPointerType[*models.System](effortFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(effortFormCallback.probe.stageOfInterest, _system)
				
				// if System is selected
				if targetSystemIDs[id] {
					// ensure effort_ is in _system.Efforts
					found := false
					for _, _b := range _system.Efforts {
						if _b == effort_ {
							found = true
							break
						}
					}
					if !found {
						_system.Efforts = append(_system.Efforts, effort_)
						effortFormCallback.probe.UpdateSliceOfPointersCallback(_system, "Efforts", &_system.Efforts)
					}
				} else {
					// ensure effort_ is NOT in _system.Efforts
					idx := slices.Index(_system.Efforts, effort_)
					if idx != -1 {
						_system.Efforts = slices.Delete(_system.Efforts, idx, idx+1)
						effortFormCallback.probe.UpdateSliceOfPointersCallback(_system, "Efforts", &_system.Efforts)
					}
				}
			}
		case "System:EffortsWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the System instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target System instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.System](effortFormCallback.probe.stageOfInterest)
			targetSystemIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSystemIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all System instances and update their EffortsWhoseNodeIsExpanded slice
			for _system := range *models.GetGongstructInstancesSetFromPointerType[*models.System](effortFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(effortFormCallback.probe.stageOfInterest, _system)
				
				// if System is selected
				if targetSystemIDs[id] {
					// ensure effort_ is in _system.EffortsWhoseNodeIsExpanded
					found := false
					for _, _b := range _system.EffortsWhoseNodeIsExpanded {
						if _b == effort_ {
							found = true
							break
						}
					}
					if !found {
						_system.EffortsWhoseNodeIsExpanded = append(_system.EffortsWhoseNodeIsExpanded, effort_)
						effortFormCallback.probe.UpdateSliceOfPointersCallback(_system, "EffortsWhoseNodeIsExpanded", &_system.EffortsWhoseNodeIsExpanded)
					}
				} else {
					// ensure effort_ is NOT in _system.EffortsWhoseNodeIsExpanded
					idx := slices.Index(_system.EffortsWhoseNodeIsExpanded, effort_)
					if idx != -1 {
						_system.EffortsWhoseNodeIsExpanded = slices.Delete(_system.EffortsWhoseNodeIsExpanded, idx, idx+1)
						effortFormCallback.probe.UpdateSliceOfPointersCallback(_system, "EffortsWhoseNodeIsExpanded", &_system.EffortsWhoseNodeIsExpanded)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if effortFormCallback.formGroup.HasSuppressButtonBeenPressed {
		effort_.Unstage(effortFormCallback.probe.stageOfInterest)
	}

	effortFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Effort](
		effortFormCallback.probe,
	)

	// display a new form by reset the form stage
	if effortFormCallback.CreationMode || effortFormCallback.formGroup.HasSuppressButtonBeenPressed {
		effortFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(effortFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__EffortFormCallback(
			nil,
			effortFormCallback.probe,
			newFormGroup,
		)
		effort := new(models.Effort)
		FillUpForm(effort, newFormGroup, effortFormCallback.probe)
		effortFormCallback.probe.formStage.Commit()
	}

	effortFormCallback.probe.ux_tree()
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

		case "RootSystems":
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
			library_.RootSystems = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "RootSystems", &library_.RootSystems)

		case "RootComplexitys":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Complexity](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Complexity, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Complexity)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Complexity](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.RootComplexitys = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "RootComplexitys", &library_.RootComplexitys)

		case "RootPerformances":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Performance](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Performance, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Performance)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Performance](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.RootPerformances = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "RootPerformances", &library_.RootPerformances)

		case "RootEfforts":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Effort](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Effort, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Effort)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Effort](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.RootEfforts = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "RootEfforts", &library_.RootEfforts)

		case "RootCompareAnalysis":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.CompareAnalysis](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.CompareAnalysis, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.CompareAnalysis)

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
			map_RowID_ID := GetMap_RowID_ID[*models.CompareAnalysis](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.RootCompareAnalysis = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "RootCompareAnalysis", &library_.RootCompareAnalysis)

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

		case "IsRootLibrary":
			FormDivBasicFieldToField(&(library_.IsRootLibrary), formDiv)
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
		case "IsSystemsNodeExpanded":
			FormDivBasicFieldToField(&(library_.IsSystemsNodeExpanded), formDiv)
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

		case "IsComplexitysNodeExpanded":
			FormDivBasicFieldToField(&(library_.IsComplexitysNodeExpanded), formDiv)
		case "ComplexitysWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Complexity](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Complexity, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Complexity)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Complexity](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.ComplexitysWhoseNodeIsExpanded = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "ComplexitysWhoseNodeIsExpanded", &library_.ComplexitysWhoseNodeIsExpanded)

		case "IsPerformancesNodeExpanded":
			FormDivBasicFieldToField(&(library_.IsPerformancesNodeExpanded), formDiv)
		case "PerformancesWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Performance](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Performance, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Performance)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Performance](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.PerformancesWhoseNodeIsExpanded = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "PerformancesWhoseNodeIsExpanded", &library_.PerformancesWhoseNodeIsExpanded)

		case "IsEffortsNodeExpanded":
			FormDivBasicFieldToField(&(library_.IsEffortsNodeExpanded), formDiv)
		case "EffortsWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Effort](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Effort, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Effort)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Effort](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.EffortsWhoseNodeIsExpanded = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "EffortsWhoseNodeIsExpanded", &library_.EffortsWhoseNodeIsExpanded)

		case "IsCompareAnalysisNodeExpanded":
			FormDivBasicFieldToField(&(library_.IsCompareAnalysisNodeExpanded), formDiv)
		case "CompareAnalysisWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.CompareAnalysis](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.CompareAnalysis, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.CompareAnalysis)

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
			map_RowID_ID := GetMap_RowID_ID[*models.CompareAnalysis](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.CompareAnalysisWhoseNodeIsExpanded = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "CompareAnalysisWhoseNodeIsExpanded", &library_.CompareAnalysisWhoseNodeIsExpanded)

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
		case "Complexities":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Complexity](noteFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Complexity, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Complexity)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Complexity](noteFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			note_.Complexities = instanceSlice
			noteFormCallback.probe.UpdateSliceOfPointersCallback(note_, "Complexities", &note_.Complexities)

		case "Performances":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Performance](noteFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Performance, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Performance)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Performance](noteFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			note_.Performances = instanceSlice
			noteFormCallback.probe.UpdateSliceOfPointersCallback(note_, "Performances", &note_.Performances)

		case "Efforts":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Effort](noteFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Effort, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Effort)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Effort](noteFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			note_.Efforts = instanceSlice
			noteFormCallback.probe.UpdateSliceOfPointersCallback(note_, "Efforts", &note_.Efforts)

		case "ComputedPrefix":
			FormDivBasicFieldToField(&(note_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(note_.IsExpanded), formDiv)
		case "IsComplexitysNodeExpanded":
			FormDivBasicFieldToField(&(note_.IsComplexitysNodeExpanded), formDiv)
		case "IsPerformancesNodeExpanded":
			FormDivBasicFieldToField(&(note_.IsPerformancesNodeExpanded), formDiv)
		case "IsEffortsNodeExpanded":
			FormDivBasicFieldToField(&(note_.IsEffortsNodeExpanded), formDiv)
		case "DiagramFlossEquation:NotesWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramFlossEquation instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramFlossEquation instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramFlossEquation](noteFormCallback.probe.stageOfInterest)
			targetDiagramFlossEquationIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramFlossEquationIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramFlossEquation instances and update their NotesWhoseNodeIsExpanded slice
			for _diagramflossequation := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramFlossEquation](noteFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(noteFormCallback.probe.stageOfInterest, _diagramflossequation)
				
				// if DiagramFlossEquation is selected
				if targetDiagramFlossEquationIDs[id] {
					// ensure note_ is in _diagramflossequation.NotesWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramflossequation.NotesWhoseNodeIsExpanded {
						if _b == note_ {
							found = true
							break
						}
					}
					if !found {
						_diagramflossequation.NotesWhoseNodeIsExpanded = append(_diagramflossequation.NotesWhoseNodeIsExpanded, note_)
						noteFormCallback.probe.UpdateSliceOfPointersCallback(_diagramflossequation, "NotesWhoseNodeIsExpanded", &_diagramflossequation.NotesWhoseNodeIsExpanded)
					}
				} else {
					// ensure note_ is NOT in _diagramflossequation.NotesWhoseNodeIsExpanded
					idx := slices.Index(_diagramflossequation.NotesWhoseNodeIsExpanded, note_)
					if idx != -1 {
						_diagramflossequation.NotesWhoseNodeIsExpanded = slices.Delete(_diagramflossequation.NotesWhoseNodeIsExpanded, idx, idx+1)
						noteFormCallback.probe.UpdateSliceOfPointersCallback(_diagramflossequation, "NotesWhoseNodeIsExpanded", &_diagramflossequation.NotesWhoseNodeIsExpanded)
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
func __gong__New__NoteComplexityShapeFormCallback(
	notecomplexityshape *models.NoteComplexityShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (notecomplexityshapeFormCallback *NoteComplexityShapeFormCallback) {
	notecomplexityshapeFormCallback = new(NoteComplexityShapeFormCallback)
	notecomplexityshapeFormCallback.probe = probe
	notecomplexityshapeFormCallback.notecomplexityshape = notecomplexityshape
	notecomplexityshapeFormCallback.formGroup = formGroup

	notecomplexityshapeFormCallback.CreationMode = (notecomplexityshape == nil)

	return
}

type NoteComplexityShapeFormCallback struct {
	notecomplexityshape *models.NoteComplexityShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (notecomplexityshapeFormCallback *NoteComplexityShapeFormCallback) OnSave() {
	notecomplexityshapeFormCallback.probe.stageOfInterest.Lock()
	defer notecomplexityshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("NoteComplexityShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	notecomplexityshapeFormCallback.probe.formStage.Checkout()

	if notecomplexityshapeFormCallback.notecomplexityshape == nil {
		notecomplexityshapeFormCallback.notecomplexityshape = new(models.NoteComplexityShape).Stage(notecomplexityshapeFormCallback.probe.stageOfInterest)
	}
	notecomplexityshape_ := notecomplexityshapeFormCallback.notecomplexityshape
	_ = notecomplexityshape_

	for _, formDiv := range notecomplexityshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(notecomplexityshape_.Name), formDiv)
		case "Note":
			FormDivSelectFieldToField(&(notecomplexityshape_.Note), notecomplexityshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Complexity":
			FormDivSelectFieldToField(&(notecomplexityshape_.Complexity), notecomplexityshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(notecomplexityshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(notecomplexityshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(notecomplexityshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(notecomplexityshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(notecomplexityshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(notecomplexityshape_.IsHidden), formDiv)
		case "DiagramFlossEquation:NoteComplexityShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramFlossEquation instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramFlossEquation instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramFlossEquation](notecomplexityshapeFormCallback.probe.stageOfInterest)
			targetDiagramFlossEquationIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramFlossEquationIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramFlossEquation instances and update their NoteComplexityShapes slice
			for _diagramflossequation := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramFlossEquation](notecomplexityshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(notecomplexityshapeFormCallback.probe.stageOfInterest, _diagramflossequation)
				
				// if DiagramFlossEquation is selected
				if targetDiagramFlossEquationIDs[id] {
					// ensure notecomplexityshape_ is in _diagramflossequation.NoteComplexityShapes
					found := false
					for _, _b := range _diagramflossequation.NoteComplexityShapes {
						if _b == notecomplexityshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramflossequation.NoteComplexityShapes = append(_diagramflossequation.NoteComplexityShapes, notecomplexityshape_)
						notecomplexityshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramflossequation, "NoteComplexityShapes", &_diagramflossequation.NoteComplexityShapes)
					}
				} else {
					// ensure notecomplexityshape_ is NOT in _diagramflossequation.NoteComplexityShapes
					idx := slices.Index(_diagramflossequation.NoteComplexityShapes, notecomplexityshape_)
					if idx != -1 {
						_diagramflossequation.NoteComplexityShapes = slices.Delete(_diagramflossequation.NoteComplexityShapes, idx, idx+1)
						notecomplexityshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramflossequation, "NoteComplexityShapes", &_diagramflossequation.NoteComplexityShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if notecomplexityshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		notecomplexityshape_.Unstage(notecomplexityshapeFormCallback.probe.stageOfInterest)
	}

	notecomplexityshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.NoteComplexityShape](
		notecomplexityshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if notecomplexityshapeFormCallback.CreationMode || notecomplexityshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		notecomplexityshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(notecomplexityshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__NoteComplexityShapeFormCallback(
			nil,
			notecomplexityshapeFormCallback.probe,
			newFormGroup,
		)
		notecomplexityshape := new(models.NoteComplexityShape)
		FillUpForm(notecomplexityshape, newFormGroup, notecomplexityshapeFormCallback.probe)
		notecomplexityshapeFormCallback.probe.formStage.Commit()
	}

	notecomplexityshapeFormCallback.probe.ux_tree()
}
func __gong__New__NoteEffortShapeFormCallback(
	noteeffortshape *models.NoteEffortShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (noteeffortshapeFormCallback *NoteEffortShapeFormCallback) {
	noteeffortshapeFormCallback = new(NoteEffortShapeFormCallback)
	noteeffortshapeFormCallback.probe = probe
	noteeffortshapeFormCallback.noteeffortshape = noteeffortshape
	noteeffortshapeFormCallback.formGroup = formGroup

	noteeffortshapeFormCallback.CreationMode = (noteeffortshape == nil)

	return
}

type NoteEffortShapeFormCallback struct {
	noteeffortshape *models.NoteEffortShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (noteeffortshapeFormCallback *NoteEffortShapeFormCallback) OnSave() {
	noteeffortshapeFormCallback.probe.stageOfInterest.Lock()
	defer noteeffortshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("NoteEffortShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	noteeffortshapeFormCallback.probe.formStage.Checkout()

	if noteeffortshapeFormCallback.noteeffortshape == nil {
		noteeffortshapeFormCallback.noteeffortshape = new(models.NoteEffortShape).Stage(noteeffortshapeFormCallback.probe.stageOfInterest)
	}
	noteeffortshape_ := noteeffortshapeFormCallback.noteeffortshape
	_ = noteeffortshape_

	for _, formDiv := range noteeffortshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(noteeffortshape_.Name), formDiv)
		case "Note":
			FormDivSelectFieldToField(&(noteeffortshape_.Note), noteeffortshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Effort":
			FormDivSelectFieldToField(&(noteeffortshape_.Effort), noteeffortshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(noteeffortshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(noteeffortshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(noteeffortshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(noteeffortshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(noteeffortshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(noteeffortshape_.IsHidden), formDiv)
		case "DiagramFlossEquation:NoteEffortShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramFlossEquation instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramFlossEquation instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramFlossEquation](noteeffortshapeFormCallback.probe.stageOfInterest)
			targetDiagramFlossEquationIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramFlossEquationIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramFlossEquation instances and update their NoteEffortShapes slice
			for _diagramflossequation := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramFlossEquation](noteeffortshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(noteeffortshapeFormCallback.probe.stageOfInterest, _diagramflossequation)
				
				// if DiagramFlossEquation is selected
				if targetDiagramFlossEquationIDs[id] {
					// ensure noteeffortshape_ is in _diagramflossequation.NoteEffortShapes
					found := false
					for _, _b := range _diagramflossequation.NoteEffortShapes {
						if _b == noteeffortshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramflossequation.NoteEffortShapes = append(_diagramflossequation.NoteEffortShapes, noteeffortshape_)
						noteeffortshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramflossequation, "NoteEffortShapes", &_diagramflossequation.NoteEffortShapes)
					}
				} else {
					// ensure noteeffortshape_ is NOT in _diagramflossequation.NoteEffortShapes
					idx := slices.Index(_diagramflossequation.NoteEffortShapes, noteeffortshape_)
					if idx != -1 {
						_diagramflossequation.NoteEffortShapes = slices.Delete(_diagramflossequation.NoteEffortShapes, idx, idx+1)
						noteeffortshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramflossequation, "NoteEffortShapes", &_diagramflossequation.NoteEffortShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if noteeffortshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteeffortshape_.Unstage(noteeffortshapeFormCallback.probe.stageOfInterest)
	}

	noteeffortshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.NoteEffortShape](
		noteeffortshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if noteeffortshapeFormCallback.CreationMode || noteeffortshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteeffortshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(noteeffortshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__NoteEffortShapeFormCallback(
			nil,
			noteeffortshapeFormCallback.probe,
			newFormGroup,
		)
		noteeffortshape := new(models.NoteEffortShape)
		FillUpForm(noteeffortshape, newFormGroup, noteeffortshapeFormCallback.probe)
		noteeffortshapeFormCallback.probe.formStage.Commit()
	}

	noteeffortshapeFormCallback.probe.ux_tree()
}
func __gong__New__NotePerformanceShapeFormCallback(
	noteperformanceshape *models.NotePerformanceShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (noteperformanceshapeFormCallback *NotePerformanceShapeFormCallback) {
	noteperformanceshapeFormCallback = new(NotePerformanceShapeFormCallback)
	noteperformanceshapeFormCallback.probe = probe
	noteperformanceshapeFormCallback.noteperformanceshape = noteperformanceshape
	noteperformanceshapeFormCallback.formGroup = formGroup

	noteperformanceshapeFormCallback.CreationMode = (noteperformanceshape == nil)

	return
}

type NotePerformanceShapeFormCallback struct {
	noteperformanceshape *models.NotePerformanceShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (noteperformanceshapeFormCallback *NotePerformanceShapeFormCallback) OnSave() {
	noteperformanceshapeFormCallback.probe.stageOfInterest.Lock()
	defer noteperformanceshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("NotePerformanceShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	noteperformanceshapeFormCallback.probe.formStage.Checkout()

	if noteperformanceshapeFormCallback.noteperformanceshape == nil {
		noteperformanceshapeFormCallback.noteperformanceshape = new(models.NotePerformanceShape).Stage(noteperformanceshapeFormCallback.probe.stageOfInterest)
	}
	noteperformanceshape_ := noteperformanceshapeFormCallback.noteperformanceshape
	_ = noteperformanceshape_

	for _, formDiv := range noteperformanceshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(noteperformanceshape_.Name), formDiv)
		case "Note":
			FormDivSelectFieldToField(&(noteperformanceshape_.Note), noteperformanceshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Performance":
			FormDivSelectFieldToField(&(noteperformanceshape_.Performance), noteperformanceshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(noteperformanceshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(noteperformanceshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(noteperformanceshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(noteperformanceshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(noteperformanceshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(noteperformanceshape_.IsHidden), formDiv)
		case "DiagramFlossEquation:NotePerformanceShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramFlossEquation instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramFlossEquation instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramFlossEquation](noteperformanceshapeFormCallback.probe.stageOfInterest)
			targetDiagramFlossEquationIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramFlossEquationIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramFlossEquation instances and update their NotePerformanceShapes slice
			for _diagramflossequation := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramFlossEquation](noteperformanceshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(noteperformanceshapeFormCallback.probe.stageOfInterest, _diagramflossequation)
				
				// if DiagramFlossEquation is selected
				if targetDiagramFlossEquationIDs[id] {
					// ensure noteperformanceshape_ is in _diagramflossequation.NotePerformanceShapes
					found := false
					for _, _b := range _diagramflossequation.NotePerformanceShapes {
						if _b == noteperformanceshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramflossequation.NotePerformanceShapes = append(_diagramflossequation.NotePerformanceShapes, noteperformanceshape_)
						noteperformanceshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramflossequation, "NotePerformanceShapes", &_diagramflossequation.NotePerformanceShapes)
					}
				} else {
					// ensure noteperformanceshape_ is NOT in _diagramflossequation.NotePerformanceShapes
					idx := slices.Index(_diagramflossequation.NotePerformanceShapes, noteperformanceshape_)
					if idx != -1 {
						_diagramflossequation.NotePerformanceShapes = slices.Delete(_diagramflossequation.NotePerformanceShapes, idx, idx+1)
						noteperformanceshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramflossequation, "NotePerformanceShapes", &_diagramflossequation.NotePerformanceShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if noteperformanceshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteperformanceshape_.Unstage(noteperformanceshapeFormCallback.probe.stageOfInterest)
	}

	noteperformanceshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.NotePerformanceShape](
		noteperformanceshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if noteperformanceshapeFormCallback.CreationMode || noteperformanceshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteperformanceshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(noteperformanceshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__NotePerformanceShapeFormCallback(
			nil,
			noteperformanceshapeFormCallback.probe,
			newFormGroup,
		)
		noteperformanceshape := new(models.NotePerformanceShape)
		FillUpForm(noteperformanceshape, newFormGroup, noteperformanceshapeFormCallback.probe)
		noteperformanceshapeFormCallback.probe.formStage.Commit()
	}

	noteperformanceshapeFormCallback.probe.ux_tree()
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
		case "DiagramFlossEquation:Note_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramFlossEquation instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramFlossEquation instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramFlossEquation](noteshapeFormCallback.probe.stageOfInterest)
			targetDiagramFlossEquationIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramFlossEquationIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramFlossEquation instances and update their Note_Shapes slice
			for _diagramflossequation := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramFlossEquation](noteshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(noteshapeFormCallback.probe.stageOfInterest, _diagramflossequation)
				
				// if DiagramFlossEquation is selected
				if targetDiagramFlossEquationIDs[id] {
					// ensure noteshape_ is in _diagramflossequation.Note_Shapes
					found := false
					for _, _b := range _diagramflossequation.Note_Shapes {
						if _b == noteshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramflossequation.Note_Shapes = append(_diagramflossequation.Note_Shapes, noteshape_)
						noteshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramflossequation, "Note_Shapes", &_diagramflossequation.Note_Shapes)
					}
				} else {
					// ensure noteshape_ is NOT in _diagramflossequation.Note_Shapes
					idx := slices.Index(_diagramflossequation.Note_Shapes, noteshape_)
					if idx != -1 {
						_diagramflossequation.Note_Shapes = slices.Delete(_diagramflossequation.Note_Shapes, idx, idx+1)
						noteshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramflossequation, "Note_Shapes", &_diagramflossequation.Note_Shapes)
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
func __gong__New__PerformanceFormCallback(
	performance *models.Performance,
	probe *Probe,
	formGroup *form.FormGroup,
) (performanceFormCallback *PerformanceFormCallback) {
	performanceFormCallback = new(PerformanceFormCallback)
	performanceFormCallback.probe = probe
	performanceFormCallback.performance = performance
	performanceFormCallback.formGroup = formGroup

	performanceFormCallback.CreationMode = (performance == nil)

	return
}

type PerformanceFormCallback struct {
	performance *models.Performance

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (performanceFormCallback *PerformanceFormCallback) OnSave() {
	performanceFormCallback.probe.stageOfInterest.Lock()
	defer performanceFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("PerformanceFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	performanceFormCallback.probe.formStage.Checkout()

	if performanceFormCallback.performance == nil {
		performanceFormCallback.performance = new(models.Performance).Stage(performanceFormCallback.probe.stageOfInterest)
	}
	performance_ := performanceFormCallback.performance
	_ = performance_

	for _, formDiv := range performanceFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(performance_.Name), formDiv)
		case "Strength":
			FormDivBasicFieldToField(&(performance_.Strength), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(performance_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(performance_.IsExpanded), formDiv)
		case "DiagramFlossEquation:PerformancesWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramFlossEquation instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramFlossEquation instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramFlossEquation](performanceFormCallback.probe.stageOfInterest)
			targetDiagramFlossEquationIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramFlossEquationIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramFlossEquation instances and update their PerformancesWhoseNodeIsExpanded slice
			for _diagramflossequation := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramFlossEquation](performanceFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(performanceFormCallback.probe.stageOfInterest, _diagramflossequation)
				
				// if DiagramFlossEquation is selected
				if targetDiagramFlossEquationIDs[id] {
					// ensure performance_ is in _diagramflossequation.PerformancesWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramflossequation.PerformancesWhoseNodeIsExpanded {
						if _b == performance_ {
							found = true
							break
						}
					}
					if !found {
						_diagramflossequation.PerformancesWhoseNodeIsExpanded = append(_diagramflossequation.PerformancesWhoseNodeIsExpanded, performance_)
						performanceFormCallback.probe.UpdateSliceOfPointersCallback(_diagramflossequation, "PerformancesWhoseNodeIsExpanded", &_diagramflossequation.PerformancesWhoseNodeIsExpanded)
					}
				} else {
					// ensure performance_ is NOT in _diagramflossequation.PerformancesWhoseNodeIsExpanded
					idx := slices.Index(_diagramflossequation.PerformancesWhoseNodeIsExpanded, performance_)
					if idx != -1 {
						_diagramflossequation.PerformancesWhoseNodeIsExpanded = slices.Delete(_diagramflossequation.PerformancesWhoseNodeIsExpanded, idx, idx+1)
						performanceFormCallback.probe.UpdateSliceOfPointersCallback(_diagramflossequation, "PerformancesWhoseNodeIsExpanded", &_diagramflossequation.PerformancesWhoseNodeIsExpanded)
					}
				}
			}
		case "Library:RootPerformances":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](performanceFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their RootPerformances slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](performanceFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(performanceFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure performance_ is in _library.RootPerformances
					found := false
					for _, _b := range _library.RootPerformances {
						if _b == performance_ {
							found = true
							break
						}
					}
					if !found {
						_library.RootPerformances = append(_library.RootPerformances, performance_)
						performanceFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootPerformances", &_library.RootPerformances)
					}
				} else {
					// ensure performance_ is NOT in _library.RootPerformances
					idx := slices.Index(_library.RootPerformances, performance_)
					if idx != -1 {
						_library.RootPerformances = slices.Delete(_library.RootPerformances, idx, idx+1)
						performanceFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootPerformances", &_library.RootPerformances)
					}
				}
			}
		case "Library:PerformancesWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](performanceFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their PerformancesWhoseNodeIsExpanded slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](performanceFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(performanceFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure performance_ is in _library.PerformancesWhoseNodeIsExpanded
					found := false
					for _, _b := range _library.PerformancesWhoseNodeIsExpanded {
						if _b == performance_ {
							found = true
							break
						}
					}
					if !found {
						_library.PerformancesWhoseNodeIsExpanded = append(_library.PerformancesWhoseNodeIsExpanded, performance_)
						performanceFormCallback.probe.UpdateSliceOfPointersCallback(_library, "PerformancesWhoseNodeIsExpanded", &_library.PerformancesWhoseNodeIsExpanded)
					}
				} else {
					// ensure performance_ is NOT in _library.PerformancesWhoseNodeIsExpanded
					idx := slices.Index(_library.PerformancesWhoseNodeIsExpanded, performance_)
					if idx != -1 {
						_library.PerformancesWhoseNodeIsExpanded = slices.Delete(_library.PerformancesWhoseNodeIsExpanded, idx, idx+1)
						performanceFormCallback.probe.UpdateSliceOfPointersCallback(_library, "PerformancesWhoseNodeIsExpanded", &_library.PerformancesWhoseNodeIsExpanded)
					}
				}
			}
		case "Note:Performances":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Note instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Note instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Note](performanceFormCallback.probe.stageOfInterest)
			targetNoteIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetNoteIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Note instances and update their Performances slice
			for _note := range *models.GetGongstructInstancesSetFromPointerType[*models.Note](performanceFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(performanceFormCallback.probe.stageOfInterest, _note)
				
				// if Note is selected
				if targetNoteIDs[id] {
					// ensure performance_ is in _note.Performances
					found := false
					for _, _b := range _note.Performances {
						if _b == performance_ {
							found = true
							break
						}
					}
					if !found {
						_note.Performances = append(_note.Performances, performance_)
						performanceFormCallback.probe.UpdateSliceOfPointersCallback(_note, "Performances", &_note.Performances)
					}
				} else {
					// ensure performance_ is NOT in _note.Performances
					idx := slices.Index(_note.Performances, performance_)
					if idx != -1 {
						_note.Performances = slices.Delete(_note.Performances, idx, idx+1)
						performanceFormCallback.probe.UpdateSliceOfPointersCallback(_note, "Performances", &_note.Performances)
					}
				}
			}
		case "System:Performances":
			// 1. Decode the AssociationStorage which contains the rowIDs of the System instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target System instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.System](performanceFormCallback.probe.stageOfInterest)
			targetSystemIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSystemIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all System instances and update their Performances slice
			for _system := range *models.GetGongstructInstancesSetFromPointerType[*models.System](performanceFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(performanceFormCallback.probe.stageOfInterest, _system)
				
				// if System is selected
				if targetSystemIDs[id] {
					// ensure performance_ is in _system.Performances
					found := false
					for _, _b := range _system.Performances {
						if _b == performance_ {
							found = true
							break
						}
					}
					if !found {
						_system.Performances = append(_system.Performances, performance_)
						performanceFormCallback.probe.UpdateSliceOfPointersCallback(_system, "Performances", &_system.Performances)
					}
				} else {
					// ensure performance_ is NOT in _system.Performances
					idx := slices.Index(_system.Performances, performance_)
					if idx != -1 {
						_system.Performances = slices.Delete(_system.Performances, idx, idx+1)
						performanceFormCallback.probe.UpdateSliceOfPointersCallback(_system, "Performances", &_system.Performances)
					}
				}
			}
		case "System:PerformancesWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the System instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target System instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.System](performanceFormCallback.probe.stageOfInterest)
			targetSystemIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSystemIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all System instances and update their PerformancesWhoseNodeIsExpanded slice
			for _system := range *models.GetGongstructInstancesSetFromPointerType[*models.System](performanceFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(performanceFormCallback.probe.stageOfInterest, _system)
				
				// if System is selected
				if targetSystemIDs[id] {
					// ensure performance_ is in _system.PerformancesWhoseNodeIsExpanded
					found := false
					for _, _b := range _system.PerformancesWhoseNodeIsExpanded {
						if _b == performance_ {
							found = true
							break
						}
					}
					if !found {
						_system.PerformancesWhoseNodeIsExpanded = append(_system.PerformancesWhoseNodeIsExpanded, performance_)
						performanceFormCallback.probe.UpdateSliceOfPointersCallback(_system, "PerformancesWhoseNodeIsExpanded", &_system.PerformancesWhoseNodeIsExpanded)
					}
				} else {
					// ensure performance_ is NOT in _system.PerformancesWhoseNodeIsExpanded
					idx := slices.Index(_system.PerformancesWhoseNodeIsExpanded, performance_)
					if idx != -1 {
						_system.PerformancesWhoseNodeIsExpanded = slices.Delete(_system.PerformancesWhoseNodeIsExpanded, idx, idx+1)
						performanceFormCallback.probe.UpdateSliceOfPointersCallback(_system, "PerformancesWhoseNodeIsExpanded", &_system.PerformancesWhoseNodeIsExpanded)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if performanceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		performance_.Unstage(performanceFormCallback.probe.stageOfInterest)
	}

	performanceFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Performance](
		performanceFormCallback.probe,
	)

	// display a new form by reset the form stage
	if performanceFormCallback.CreationMode || performanceFormCallback.formGroup.HasSuppressButtonBeenPressed {
		performanceFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(performanceFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PerformanceFormCallback(
			nil,
			performanceFormCallback.probe,
			newFormGroup,
		)
		performance := new(models.Performance)
		FillUpForm(performance, newFormGroup, performanceFormCallback.probe)
		performanceFormCallback.probe.formStage.Commit()
	}

	performanceFormCallback.probe.ux_tree()
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
		case "Complexities":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Complexity](systemFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Complexity, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Complexity)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Complexity](systemFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			system_.Complexities = instanceSlice
			systemFormCallback.probe.UpdateSliceOfPointersCallback(system_, "Complexities", &system_.Complexities)

		case "Performances":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Performance](systemFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Performance, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Performance)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Performance](systemFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			system_.Performances = instanceSlice
			systemFormCallback.probe.UpdateSliceOfPointersCallback(system_, "Performances", &system_.Performances)

		case "Efforts":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Effort](systemFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Effort, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Effort)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Effort](systemFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			system_.Efforts = instanceSlice
			systemFormCallback.probe.UpdateSliceOfPointersCallback(system_, "Efforts", &system_.Efforts)

		case "ComputedPrefix":
			FormDivBasicFieldToField(&(system_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(system_.IsExpanded), formDiv)
		case "SVG_Path":
			FormDivBasicFieldToField(&(system_.SVG_Path), formDiv)
		case "InverseAppliedScaling":
			FormDivBasicFieldToField(&(system_.InverseAppliedScaling), formDiv)
		case "DiagramFlossEquations":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DiagramFlossEquation](systemFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DiagramFlossEquation, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DiagramFlossEquation)

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
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramFlossEquation](systemFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			system_.DiagramFlossEquations = instanceSlice
			systemFormCallback.probe.UpdateSliceOfPointersCallback(system_, "DiagramFlossEquations", &system_.DiagramFlossEquations)

		case "DiagramFlossEquationsWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DiagramFlossEquation](systemFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DiagramFlossEquation, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DiagramFlossEquation)

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
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramFlossEquation](systemFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			system_.DiagramFlossEquationsWhoseNodeIsExpanded = instanceSlice
			systemFormCallback.probe.UpdateSliceOfPointersCallback(system_, "DiagramFlossEquationsWhoseNodeIsExpanded", &system_.DiagramFlossEquationsWhoseNodeIsExpanded)

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

		case "IsComplexitysNodeExpanded":
			FormDivBasicFieldToField(&(system_.IsComplexitysNodeExpanded), formDiv)
		case "ComplexitysWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Complexity](systemFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Complexity, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Complexity)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Complexity](systemFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			system_.ComplexitysWhoseNodeIsExpanded = instanceSlice
			systemFormCallback.probe.UpdateSliceOfPointersCallback(system_, "ComplexitysWhoseNodeIsExpanded", &system_.ComplexitysWhoseNodeIsExpanded)

		case "IsPerformancesNodeExpanded":
			FormDivBasicFieldToField(&(system_.IsPerformancesNodeExpanded), formDiv)
		case "PerformancesWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Performance](systemFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Performance, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Performance)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Performance](systemFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			system_.PerformancesWhoseNodeIsExpanded = instanceSlice
			systemFormCallback.probe.UpdateSliceOfPointersCallback(system_, "PerformancesWhoseNodeIsExpanded", &system_.PerformancesWhoseNodeIsExpanded)

		case "IsEffortsNodeExpanded":
			FormDivBasicFieldToField(&(system_.IsEffortsNodeExpanded), formDiv)
		case "EffortsWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Effort](systemFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Effort, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Effort)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Effort](systemFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			system_.EffortsWhoseNodeIsExpanded = instanceSlice
			systemFormCallback.probe.UpdateSliceOfPointersCallback(system_, "EffortsWhoseNodeIsExpanded", &system_.EffortsWhoseNodeIsExpanded)

		case "Library:RootSystems":
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

			// 3. Iterate over all Library instances and update their RootSystems slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](systemFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(systemFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure system_ is in _library.RootSystems
					found := false
					for _, _b := range _library.RootSystems {
						if _b == system_ {
							found = true
							break
						}
					}
					if !found {
						_library.RootSystems = append(_library.RootSystems, system_)
						systemFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootSystems", &_library.RootSystems)
					}
				} else {
					// ensure system_ is NOT in _library.RootSystems
					idx := slices.Index(_library.RootSystems, system_)
					if idx != -1 {
						_library.RootSystems = slices.Delete(_library.RootSystems, idx, idx+1)
						systemFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootSystems", &_library.RootSystems)
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
