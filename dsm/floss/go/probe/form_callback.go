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
		case "System:Complexitys":
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

			// 3. Iterate over all System instances and update their Complexitys slice
			for _system := range *models.GetGongstructInstancesSetFromPointerType[*models.System](complexityFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(complexityFormCallback.probe.stageOfInterest, _system)
				
				// if System is selected
				if targetSystemIDs[id] {
					// ensure complexity_ is in _system.Complexitys
					found := false
					for _, _b := range _system.Complexitys {
						if _b == complexity_ {
							found = true
							break
						}
					}
					if !found {
						_system.Complexitys = append(_system.Complexitys, complexity_)
						complexityFormCallback.probe.UpdateSliceOfPointersCallback(_system, "Complexitys", &_system.Complexitys)
					}
				} else {
					// ensure complexity_ is NOT in _system.Complexitys
					idx := slices.Index(_system.Complexitys, complexity_)
					if idx != -1 {
						_system.Complexitys = slices.Delete(_system.Complexitys, idx, idx+1)
						complexityFormCallback.probe.UpdateSliceOfPointersCallback(_system, "Complexitys", &_system.Complexitys)
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
func __gong__New__DiagramFlossFormCallback(
	diagramfloss *models.DiagramFloss,
	probe *Probe,
	formGroup *form.FormGroup,
) (diagramflossFormCallback *DiagramFlossFormCallback) {
	diagramflossFormCallback = new(DiagramFlossFormCallback)
	diagramflossFormCallback.probe = probe
	diagramflossFormCallback.diagramfloss = diagramfloss
	diagramflossFormCallback.formGroup = formGroup

	diagramflossFormCallback.CreationMode = (diagramfloss == nil)

	return
}

type DiagramFlossFormCallback struct {
	diagramfloss *models.DiagramFloss

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (diagramflossFormCallback *DiagramFlossFormCallback) OnSave() {
	diagramflossFormCallback.probe.stageOfInterest.Lock()
	defer diagramflossFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DiagramFlossFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	diagramflossFormCallback.probe.formStage.Checkout()

	if diagramflossFormCallback.diagramfloss == nil {
		diagramflossFormCallback.diagramfloss = new(models.DiagramFloss).Stage(diagramflossFormCallback.probe.stageOfInterest)
	}
	diagramfloss_ := diagramflossFormCallback.diagramfloss
	_ = diagramfloss_

	for _, formDiv := range diagramflossFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(diagramfloss_.Name), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(diagramfloss_.Description), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(diagramfloss_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(diagramfloss_.IsExpanded), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(diagramfloss_.IsChecked), formDiv)
		case "IsEditable_":
			FormDivBasicFieldToField(&(diagramfloss_.IsEditable_), formDiv)
		case "IsShowPrefix":
			FormDivBasicFieldToField(&(diagramfloss_.IsShowPrefix), formDiv)
		case "DefaultBoxWidth":
			FormDivBasicFieldToField(&(diagramfloss_.DefaultBoxWidth), formDiv)
		case "DefaultBoxHeigth":
			FormDivBasicFieldToField(&(diagramfloss_.DefaultBoxHeigth), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(diagramfloss_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(diagramfloss_.Height), formDiv)
		case "System_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.SystemShape](diagramflossFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.SystemShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.SystemShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramflossFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.SystemShape](diagramflossFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramfloss_.System_Shapes = instanceSlice
			diagramflossFormCallback.probe.UpdateSliceOfPointersCallback(diagramfloss_, "System_Shapes", &diagramfloss_.System_Shapes)

		case "IsSystemsNodeExpanded":
			FormDivBasicFieldToField(&(diagramfloss_.IsSystemsNodeExpanded), formDiv)
		case "SystemsWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.System](diagramflossFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.System, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.System)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramflossFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.System](diagramflossFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramfloss_.SystemsWhoseNodeIsExpanded = instanceSlice
			diagramflossFormCallback.probe.UpdateSliceOfPointersCallback(diagramfloss_, "SystemsWhoseNodeIsExpanded", &diagramfloss_.SystemsWhoseNodeIsExpanded)

		case "System:DiagramFlosses":
			// 1. Decode the AssociationStorage which contains the rowIDs of the System instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target System instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.System](diagramflossFormCallback.probe.stageOfInterest)
			targetSystemIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSystemIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all System instances and update their DiagramFlosses slice
			for _system := range *models.GetGongstructInstancesSetFromPointerType[*models.System](diagramflossFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(diagramflossFormCallback.probe.stageOfInterest, _system)
				
				// if System is selected
				if targetSystemIDs[id] {
					// ensure diagramfloss_ is in _system.DiagramFlosses
					found := false
					for _, _b := range _system.DiagramFlosses {
						if _b == diagramfloss_ {
							found = true
							break
						}
					}
					if !found {
						_system.DiagramFlosses = append(_system.DiagramFlosses, diagramfloss_)
						diagramflossFormCallback.probe.UpdateSliceOfPointersCallback(_system, "DiagramFlosses", &_system.DiagramFlosses)
					}
				} else {
					// ensure diagramfloss_ is NOT in _system.DiagramFlosses
					idx := slices.Index(_system.DiagramFlosses, diagramfloss_)
					if idx != -1 {
						_system.DiagramFlosses = slices.Delete(_system.DiagramFlosses, idx, idx+1)
						diagramflossFormCallback.probe.UpdateSliceOfPointersCallback(_system, "DiagramFlosses", &_system.DiagramFlosses)
					}
				}
			}
		case "System:DiagramFlossWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the System instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target System instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.System](diagramflossFormCallback.probe.stageOfInterest)
			targetSystemIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSystemIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all System instances and update their DiagramFlossWhoseNodeIsExpanded slice
			for _system := range *models.GetGongstructInstancesSetFromPointerType[*models.System](diagramflossFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(diagramflossFormCallback.probe.stageOfInterest, _system)
				
				// if System is selected
				if targetSystemIDs[id] {
					// ensure diagramfloss_ is in _system.DiagramFlossWhoseNodeIsExpanded
					found := false
					for _, _b := range _system.DiagramFlossWhoseNodeIsExpanded {
						if _b == diagramfloss_ {
							found = true
							break
						}
					}
					if !found {
						_system.DiagramFlossWhoseNodeIsExpanded = append(_system.DiagramFlossWhoseNodeIsExpanded, diagramfloss_)
						diagramflossFormCallback.probe.UpdateSliceOfPointersCallback(_system, "DiagramFlossWhoseNodeIsExpanded", &_system.DiagramFlossWhoseNodeIsExpanded)
					}
				} else {
					// ensure diagramfloss_ is NOT in _system.DiagramFlossWhoseNodeIsExpanded
					idx := slices.Index(_system.DiagramFlossWhoseNodeIsExpanded, diagramfloss_)
					if idx != -1 {
						_system.DiagramFlossWhoseNodeIsExpanded = slices.Delete(_system.DiagramFlossWhoseNodeIsExpanded, idx, idx+1)
						diagramflossFormCallback.probe.UpdateSliceOfPointersCallback(_system, "DiagramFlossWhoseNodeIsExpanded", &_system.DiagramFlossWhoseNodeIsExpanded)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if diagramflossFormCallback.formGroup.HasSuppressButtonBeenPressed {
		diagramfloss_.Unstage(diagramflossFormCallback.probe.stageOfInterest)
	}

	diagramflossFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.DiagramFloss](
		diagramflossFormCallback.probe,
	)

	// display a new form by reset the form stage
	if diagramflossFormCallback.CreationMode || diagramflossFormCallback.formGroup.HasSuppressButtonBeenPressed {
		diagramflossFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(diagramflossFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DiagramFlossFormCallback(
			nil,
			diagramflossFormCallback.probe,
			newFormGroup,
		)
		diagramfloss := new(models.DiagramFloss)
		FillUpForm(diagramfloss, newFormGroup, diagramflossFormCallback.probe)
		diagramflossFormCallback.probe.formStage.Commit()
	}

	diagramflossFormCallback.probe.ux_tree()
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
		case "Complexitys":
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
			system_.Complexitys = instanceSlice
			systemFormCallback.probe.UpdateSliceOfPointersCallback(system_, "Complexitys", &system_.Complexitys)

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
		case "DiagramFlosses":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DiagramFloss](systemFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DiagramFloss, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DiagramFloss)

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
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramFloss](systemFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			system_.DiagramFlosses = instanceSlice
			systemFormCallback.probe.UpdateSliceOfPointersCallback(system_, "DiagramFlosses", &system_.DiagramFlosses)

		case "DiagramFlossWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DiagramFloss](systemFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DiagramFloss, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DiagramFloss)

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
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramFloss](systemFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			system_.DiagramFlossWhoseNodeIsExpanded = instanceSlice
			systemFormCallback.probe.UpdateSliceOfPointersCallback(system_, "DiagramFlossWhoseNodeIsExpanded", &system_.DiagramFlossWhoseNodeIsExpanded)

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

		case "DiagramFloss:SystemsWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramFloss instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramFloss instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramFloss](systemFormCallback.probe.stageOfInterest)
			targetDiagramFlossIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramFlossIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramFloss instances and update their SystemsWhoseNodeIsExpanded slice
			for _diagramfloss := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramFloss](systemFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(systemFormCallback.probe.stageOfInterest, _diagramfloss)
				
				// if DiagramFloss is selected
				if targetDiagramFlossIDs[id] {
					// ensure system_ is in _diagramfloss.SystemsWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramfloss.SystemsWhoseNodeIsExpanded {
						if _b == system_ {
							found = true
							break
						}
					}
					if !found {
						_diagramfloss.SystemsWhoseNodeIsExpanded = append(_diagramfloss.SystemsWhoseNodeIsExpanded, system_)
						systemFormCallback.probe.UpdateSliceOfPointersCallback(_diagramfloss, "SystemsWhoseNodeIsExpanded", &_diagramfloss.SystemsWhoseNodeIsExpanded)
					}
				} else {
					// ensure system_ is NOT in _diagramfloss.SystemsWhoseNodeIsExpanded
					idx := slices.Index(_diagramfloss.SystemsWhoseNodeIsExpanded, system_)
					if idx != -1 {
						_diagramfloss.SystemsWhoseNodeIsExpanded = slices.Delete(_diagramfloss.SystemsWhoseNodeIsExpanded, idx, idx+1)
						systemFormCallback.probe.UpdateSliceOfPointersCallback(_diagramfloss, "SystemsWhoseNodeIsExpanded", &_diagramfloss.SystemsWhoseNodeIsExpanded)
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
		case "DiagramFloss:System_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramFloss instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramFloss instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramFloss](systemshapeFormCallback.probe.stageOfInterest)
			targetDiagramFlossIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramFlossIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramFloss instances and update their System_Shapes slice
			for _diagramfloss := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramFloss](systemshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(systemshapeFormCallback.probe.stageOfInterest, _diagramfloss)
				
				// if DiagramFloss is selected
				if targetDiagramFlossIDs[id] {
					// ensure systemshape_ is in _diagramfloss.System_Shapes
					found := false
					for _, _b := range _diagramfloss.System_Shapes {
						if _b == systemshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramfloss.System_Shapes = append(_diagramfloss.System_Shapes, systemshape_)
						systemshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramfloss, "System_Shapes", &_diagramfloss.System_Shapes)
					}
				} else {
					// ensure systemshape_ is NOT in _diagramfloss.System_Shapes
					idx := slices.Index(_diagramfloss.System_Shapes, systemshape_)
					if idx != -1 {
						_diagramfloss.System_Shapes = slices.Delete(_diagramfloss.System_Shapes, idx, idx+1)
						systemshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramfloss, "System_Shapes", &_diagramfloss.System_Shapes)
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
