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
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(diagramstructure_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(diagramstructure_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(diagramstructure_.LayoutDirection), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(diagramstructure_.IsChecked), formDiv)
		case "IsEditable_":
			FormDivBasicFieldToField(&(diagramstructure_.IsEditable_), formDiv)
		case "IsShowPrefix":
			FormDivBasicFieldToField(&(diagramstructure_.IsShowPrefix), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(diagramstructure_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(diagramstructure_.Height), formDiv)
		case "DefaultBoxWidth":
			FormDivBasicFieldToField(&(diagramstructure_.DefaultBoxWidth), formDiv)
		case "DefaultBoxHeigth":
			FormDivBasicFieldToField(&(diagramstructure_.DefaultBoxHeigth), formDiv)
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
		case "PartsWhoseNodeIsExpanded":
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
			diagramstructure_.PartsWhoseNodeIsExpanded = instanceSlice
			diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(diagramstructure_, "PartsWhoseNodeIsExpanded", &diagramstructure_.PartsWhoseNodeIsExpanded)

		case "Link_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.LinkAssociationShape](diagramstructureFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.LinkAssociationShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.LinkAssociationShape)

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
			map_RowID_ID := GetMap_RowID_ID[*models.LinkAssociationShape](diagramstructureFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramstructure_.Link_Shapes = instanceSlice
			diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(diagramstructure_, "Link_Shapes", &diagramstructure_.Link_Shapes)

		case "LinksWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Link](diagramstructureFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Link, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Link)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Link](diagramstructureFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagramstructure_.LinksWhoseNodeIsExpanded = instanceSlice
			diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(diagramstructure_, "LinksWhoseNodeIsExpanded", &diagramstructure_.LinksWhoseNodeIsExpanded)

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
		case "System:DiagramStructuresWhoseNodeIsExpanded":
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

			// 3. Iterate over all System instances and update their DiagramStructuresWhoseNodeIsExpanded slice
			for _system := range *models.GetGongstructInstancesSetFromPointerType[*models.System](diagramstructureFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(diagramstructureFormCallback.probe.stageOfInterest, _system)
				
				// if System is selected
				if targetSystemIDs[id] {
					// ensure diagramstructure_ is in _system.DiagramStructuresWhoseNodeIsExpanded
					found := false
					for _, _b := range _system.DiagramStructuresWhoseNodeIsExpanded {
						if _b == diagramstructure_ {
							found = true
							break
						}
					}
					if !found {
						_system.DiagramStructuresWhoseNodeIsExpanded = append(_system.DiagramStructuresWhoseNodeIsExpanded, diagramstructure_)
						diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(_system, "DiagramStructuresWhoseNodeIsExpanded", &_system.DiagramStructuresWhoseNodeIsExpanded)
					}
				} else {
					// ensure diagramstructure_ is NOT in _system.DiagramStructuresWhoseNodeIsExpanded
					idx := slices.Index(_system.DiagramStructuresWhoseNodeIsExpanded, diagramstructure_)
					if idx != -1 {
						_system.DiagramStructuresWhoseNodeIsExpanded = slices.Delete(_system.DiagramStructuresWhoseNodeIsExpanded, idx, idx+1)
						diagramstructureFormCallback.probe.UpdateSliceOfPointersCallback(_system, "DiagramStructuresWhoseNodeIsExpanded", &_system.DiagramStructuresWhoseNodeIsExpanded)
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
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(library_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(library_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(library_.LayoutDirection), formDiv)
		case "IsRootLibrary":
			FormDivBasicFieldToField(&(library_.IsRootLibrary), formDiv)
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
func __gong__New__LinkFormCallback(
	link *models.Link,
	probe *Probe,
	formGroup *form.FormGroup,
) (linkFormCallback *LinkFormCallback) {
	linkFormCallback = new(LinkFormCallback)
	linkFormCallback.probe = probe
	linkFormCallback.link = link
	linkFormCallback.formGroup = formGroup

	linkFormCallback.CreationMode = (link == nil)

	return
}

type LinkFormCallback struct {
	link *models.Link

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (linkFormCallback *LinkFormCallback) OnSave() {
	linkFormCallback.probe.stageOfInterest.Lock()
	defer linkFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("LinkFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	linkFormCallback.probe.formStage.Checkout()

	if linkFormCallback.link == nil {
		linkFormCallback.link = new(models.Link).Stage(linkFormCallback.probe.stageOfInterest)
	}
	link_ := linkFormCallback.link
	_ = link_

	for _, formDiv := range linkFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(link_.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(link_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(link_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(link_.LayoutDirection), formDiv)
		case "Source":
			FormDivSelectFieldToField(&(link_.Source), linkFormCallback.probe.stageOfInterest, formDiv)
		case "Target":
			FormDivSelectFieldToField(&(link_.Target), linkFormCallback.probe.stageOfInterest, formDiv)
		case "DiagramStructure:LinksWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramStructure instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramStructure instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramStructure](linkFormCallback.probe.stageOfInterest)
			targetDiagramStructureIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramStructureIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramStructure instances and update their LinksWhoseNodeIsExpanded slice
			for _diagramstructure := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramStructure](linkFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(linkFormCallback.probe.stageOfInterest, _diagramstructure)
				
				// if DiagramStructure is selected
				if targetDiagramStructureIDs[id] {
					// ensure link_ is in _diagramstructure.LinksWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramstructure.LinksWhoseNodeIsExpanded {
						if _b == link_ {
							found = true
							break
						}
					}
					if !found {
						_diagramstructure.LinksWhoseNodeIsExpanded = append(_diagramstructure.LinksWhoseNodeIsExpanded, link_)
						linkFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "LinksWhoseNodeIsExpanded", &_diagramstructure.LinksWhoseNodeIsExpanded)
					}
				} else {
					// ensure link_ is NOT in _diagramstructure.LinksWhoseNodeIsExpanded
					idx := slices.Index(_diagramstructure.LinksWhoseNodeIsExpanded, link_)
					if idx != -1 {
						_diagramstructure.LinksWhoseNodeIsExpanded = slices.Delete(_diagramstructure.LinksWhoseNodeIsExpanded, idx, idx+1)
						linkFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "LinksWhoseNodeIsExpanded", &_diagramstructure.LinksWhoseNodeIsExpanded)
					}
				}
			}
		case "System:Links":
			// 1. Decode the AssociationStorage which contains the rowIDs of the System instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target System instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.System](linkFormCallback.probe.stageOfInterest)
			targetSystemIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSystemIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all System instances and update their Links slice
			for _system := range *models.GetGongstructInstancesSetFromPointerType[*models.System](linkFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(linkFormCallback.probe.stageOfInterest, _system)
				
				// if System is selected
				if targetSystemIDs[id] {
					// ensure link_ is in _system.Links
					found := false
					for _, _b := range _system.Links {
						if _b == link_ {
							found = true
							break
						}
					}
					if !found {
						_system.Links = append(_system.Links, link_)
						linkFormCallback.probe.UpdateSliceOfPointersCallback(_system, "Links", &_system.Links)
					}
				} else {
					// ensure link_ is NOT in _system.Links
					idx := slices.Index(_system.Links, link_)
					if idx != -1 {
						_system.Links = slices.Delete(_system.Links, idx, idx+1)
						linkFormCallback.probe.UpdateSliceOfPointersCallback(_system, "Links", &_system.Links)
					}
				}
			}
		case "System:LinksWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the System instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target System instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.System](linkFormCallback.probe.stageOfInterest)
			targetSystemIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetSystemIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all System instances and update their LinksWhoseNodeIsExpanded slice
			for _system := range *models.GetGongstructInstancesSetFromPointerType[*models.System](linkFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(linkFormCallback.probe.stageOfInterest, _system)
				
				// if System is selected
				if targetSystemIDs[id] {
					// ensure link_ is in _system.LinksWhoseNodeIsExpanded
					found := false
					for _, _b := range _system.LinksWhoseNodeIsExpanded {
						if _b == link_ {
							found = true
							break
						}
					}
					if !found {
						_system.LinksWhoseNodeIsExpanded = append(_system.LinksWhoseNodeIsExpanded, link_)
						linkFormCallback.probe.UpdateSliceOfPointersCallback(_system, "LinksWhoseNodeIsExpanded", &_system.LinksWhoseNodeIsExpanded)
					}
				} else {
					// ensure link_ is NOT in _system.LinksWhoseNodeIsExpanded
					idx := slices.Index(_system.LinksWhoseNodeIsExpanded, link_)
					if idx != -1 {
						_system.LinksWhoseNodeIsExpanded = slices.Delete(_system.LinksWhoseNodeIsExpanded, idx, idx+1)
						linkFormCallback.probe.UpdateSliceOfPointersCallback(_system, "LinksWhoseNodeIsExpanded", &_system.LinksWhoseNodeIsExpanded)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if linkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		link_.Unstage(linkFormCallback.probe.stageOfInterest)
	}

	linkFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Link](
		linkFormCallback.probe,
	)

	// display a new form by reset the form stage
	if linkFormCallback.CreationMode || linkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		linkFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(linkFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LinkFormCallback(
			nil,
			linkFormCallback.probe,
			newFormGroup,
		)
		link := new(models.Link)
		FillUpForm(link, newFormGroup, linkFormCallback.probe)
		linkFormCallback.probe.formStage.Commit()
	}

	linkFormCallback.probe.ux_tree()
}
func __gong__New__LinkAssociationShapeFormCallback(
	linkassociationshape *models.LinkAssociationShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (linkassociationshapeFormCallback *LinkAssociationShapeFormCallback) {
	linkassociationshapeFormCallback = new(LinkAssociationShapeFormCallback)
	linkassociationshapeFormCallback.probe = probe
	linkassociationshapeFormCallback.linkassociationshape = linkassociationshape
	linkassociationshapeFormCallback.formGroup = formGroup

	linkassociationshapeFormCallback.CreationMode = (linkassociationshape == nil)

	return
}

type LinkAssociationShapeFormCallback struct {
	linkassociationshape *models.LinkAssociationShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (linkassociationshapeFormCallback *LinkAssociationShapeFormCallback) OnSave() {
	linkassociationshapeFormCallback.probe.stageOfInterest.Lock()
	defer linkassociationshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("LinkAssociationShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	linkassociationshapeFormCallback.probe.formStage.Checkout()

	if linkassociationshapeFormCallback.linkassociationshape == nil {
		linkassociationshapeFormCallback.linkassociationshape = new(models.LinkAssociationShape).Stage(linkassociationshapeFormCallback.probe.stageOfInterest)
	}
	linkassociationshape_ := linkassociationshapeFormCallback.linkassociationshape
	_ = linkassociationshape_

	for _, formDiv := range linkassociationshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(linkassociationshape_.Name), formDiv)
		case "Link":
			FormDivSelectFieldToField(&(linkassociationshape_.Link), linkassociationshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(linkassociationshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(linkassociationshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(linkassociationshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(linkassociationshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(linkassociationshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(linkassociationshape_.IsHidden), formDiv)
		case "DiagramStructure:Link_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DiagramStructure instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DiagramStructure instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DiagramStructure](linkassociationshapeFormCallback.probe.stageOfInterest)
			targetDiagramStructureIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramStructureIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DiagramStructure instances and update their Link_Shapes slice
			for _diagramstructure := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramStructure](linkassociationshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(linkassociationshapeFormCallback.probe.stageOfInterest, _diagramstructure)
				
				// if DiagramStructure is selected
				if targetDiagramStructureIDs[id] {
					// ensure linkassociationshape_ is in _diagramstructure.Link_Shapes
					found := false
					for _, _b := range _diagramstructure.Link_Shapes {
						if _b == linkassociationshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagramstructure.Link_Shapes = append(_diagramstructure.Link_Shapes, linkassociationshape_)
						linkassociationshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "Link_Shapes", &_diagramstructure.Link_Shapes)
					}
				} else {
					// ensure linkassociationshape_ is NOT in _diagramstructure.Link_Shapes
					idx := slices.Index(_diagramstructure.Link_Shapes, linkassociationshape_)
					if idx != -1 {
						_diagramstructure.Link_Shapes = slices.Delete(_diagramstructure.Link_Shapes, idx, idx+1)
						linkassociationshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "Link_Shapes", &_diagramstructure.Link_Shapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if linkassociationshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		linkassociationshape_.Unstage(linkassociationshapeFormCallback.probe.stageOfInterest)
	}

	linkassociationshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.LinkAssociationShape](
		linkassociationshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if linkassociationshapeFormCallback.CreationMode || linkassociationshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		linkassociationshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(linkassociationshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LinkAssociationShapeFormCallback(
			nil,
			linkassociationshapeFormCallback.probe,
			newFormGroup,
		)
		linkassociationshape := new(models.LinkAssociationShape)
		FillUpForm(linkassociationshape, newFormGroup, linkassociationshapeFormCallback.probe)
		linkassociationshapeFormCallback.probe.formStage.Commit()
	}

	linkassociationshapeFormCallback.probe.ux_tree()
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
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(part_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(part_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(part_.LayoutDirection), formDiv)
		case "DiagramStructure:PartsWhoseNodeIsExpanded":
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

			// 3. Iterate over all DiagramStructure instances and update their PartsWhoseNodeIsExpanded slice
			for _diagramstructure := range *models.GetGongstructInstancesSetFromPointerType[*models.DiagramStructure](partFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(partFormCallback.probe.stageOfInterest, _diagramstructure)
				
				// if DiagramStructure is selected
				if targetDiagramStructureIDs[id] {
					// ensure part_ is in _diagramstructure.PartsWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagramstructure.PartsWhoseNodeIsExpanded {
						if _b == part_ {
							found = true
							break
						}
					}
					if !found {
						_diagramstructure.PartsWhoseNodeIsExpanded = append(_diagramstructure.PartsWhoseNodeIsExpanded, part_)
						partFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "PartsWhoseNodeIsExpanded", &_diagramstructure.PartsWhoseNodeIsExpanded)
					}
				} else {
					// ensure part_ is NOT in _diagramstructure.PartsWhoseNodeIsExpanded
					idx := slices.Index(_diagramstructure.PartsWhoseNodeIsExpanded, part_)
					if idx != -1 {
						_diagramstructure.PartsWhoseNodeIsExpanded = slices.Delete(_diagramstructure.PartsWhoseNodeIsExpanded, idx, idx+1)
						partFormCallback.probe.UpdateSliceOfPointersCallback(_diagramstructure, "PartsWhoseNodeIsExpanded", &_diagramstructure.PartsWhoseNodeIsExpanded)
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
		case "System:PartsWhoseNodeIsExpanded":
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

			// 3. Iterate over all System instances and update their PartsWhoseNodeIsExpanded slice
			for _system := range *models.GetGongstructInstancesSetFromPointerType[*models.System](partFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(partFormCallback.probe.stageOfInterest, _system)
				
				// if System is selected
				if targetSystemIDs[id] {
					// ensure part_ is in _system.PartsWhoseNodeIsExpanded
					found := false
					for _, _b := range _system.PartsWhoseNodeIsExpanded {
						if _b == part_ {
							found = true
							break
						}
					}
					if !found {
						_system.PartsWhoseNodeIsExpanded = append(_system.PartsWhoseNodeIsExpanded, part_)
						partFormCallback.probe.UpdateSliceOfPointersCallback(_system, "PartsWhoseNodeIsExpanded", &_system.PartsWhoseNodeIsExpanded)
					}
				} else {
					// ensure part_ is NOT in _system.PartsWhoseNodeIsExpanded
					idx := slices.Index(_system.PartsWhoseNodeIsExpanded, part_)
					if idx != -1 {
						_system.PartsWhoseNodeIsExpanded = slices.Delete(_system.PartsWhoseNodeIsExpanded, idx, idx+1)
						partFormCallback.probe.UpdateSliceOfPointersCallback(_system, "PartsWhoseNodeIsExpanded", &_system.PartsWhoseNodeIsExpanded)
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
		case "WidthWeight":
			FormDivBasicFieldToField(&(partshape_.WidthWeight), formDiv)
		case "OverideLayoutDirection":
			FormDivBasicFieldToField(&(partshape_.OverideLayoutDirection), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(partshape_.LayoutDirection), formDiv)
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
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(system_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(system_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(system_.LayoutDirection), formDiv)
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

		case "IsPartsNodeExpanded":
			FormDivBasicFieldToField(&(system_.IsPartsNodeExpanded), formDiv)
		case "PartsWhoseNodeIsExpanded":
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
			system_.PartsWhoseNodeIsExpanded = instanceSlice
			systemFormCallback.probe.UpdateSliceOfPointersCallback(system_, "PartsWhoseNodeIsExpanded", &system_.PartsWhoseNodeIsExpanded)

		case "SubSystems":
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
			system_.SubSystems = instanceSlice
			systemFormCallback.probe.UpdateSliceOfPointersCallback(system_, "SubSystems", &system_.SubSystems)

		case "IsSubSystemsNodeExpanded":
			FormDivBasicFieldToField(&(system_.IsSubSystemsNodeExpanded), formDiv)
		case "SubSystemsWhoseNodeIsExpanded":
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
			system_.SubSystemsWhoseNodeIsExpanded = instanceSlice
			systemFormCallback.probe.UpdateSliceOfPointersCallback(system_, "SubSystemsWhoseNodeIsExpanded", &system_.SubSystemsWhoseNodeIsExpanded)

		case "Links":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Link](systemFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Link, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Link)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Link](systemFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			system_.Links = instanceSlice
			systemFormCallback.probe.UpdateSliceOfPointersCallback(system_, "Links", &system_.Links)

		case "IsLinksNodeExpanded":
			FormDivBasicFieldToField(&(system_.IsLinksNodeExpanded), formDiv)
		case "LinksWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Link](systemFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Link, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Link)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Link](systemFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			system_.LinksWhoseNodeIsExpanded = instanceSlice
			systemFormCallback.probe.UpdateSliceOfPointersCallback(system_, "LinksWhoseNodeIsExpanded", &system_.LinksWhoseNodeIsExpanded)

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

		case "IsDiagramStructuresNodeExpanded":
			FormDivBasicFieldToField(&(system_.IsDiagramStructuresNodeExpanded), formDiv)
		case "DiagramStructuresWhoseNodeIsExpanded":
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
			system_.DiagramStructuresWhoseNodeIsExpanded = instanceSlice
			systemFormCallback.probe.UpdateSliceOfPointersCallback(system_, "DiagramStructuresWhoseNodeIsExpanded", &system_.DiagramStructuresWhoseNodeIsExpanded)

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
		case "System:SubSystems":
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

			// 3. Iterate over all System instances and update their SubSystems slice
			for _system := range *models.GetGongstructInstancesSetFromPointerType[*models.System](systemFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(systemFormCallback.probe.stageOfInterest, _system)
				
				// if System is selected
				if targetSystemIDs[id] {
					// ensure system_ is in _system.SubSystems
					found := false
					for _, _b := range _system.SubSystems {
						if _b == system_ {
							found = true
							break
						}
					}
					if !found {
						_system.SubSystems = append(_system.SubSystems, system_)
						systemFormCallback.probe.UpdateSliceOfPointersCallback(_system, "SubSystems", &_system.SubSystems)
					}
				} else {
					// ensure system_ is NOT in _system.SubSystems
					idx := slices.Index(_system.SubSystems, system_)
					if idx != -1 {
						_system.SubSystems = slices.Delete(_system.SubSystems, idx, idx+1)
						systemFormCallback.probe.UpdateSliceOfPointersCallback(_system, "SubSystems", &_system.SubSystems)
					}
				}
			}
		case "System:SubSystemsWhoseNodeIsExpanded":
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

			// 3. Iterate over all System instances and update their SubSystemsWhoseNodeIsExpanded slice
			for _system := range *models.GetGongstructInstancesSetFromPointerType[*models.System](systemFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(systemFormCallback.probe.stageOfInterest, _system)
				
				// if System is selected
				if targetSystemIDs[id] {
					// ensure system_ is in _system.SubSystemsWhoseNodeIsExpanded
					found := false
					for _, _b := range _system.SubSystemsWhoseNodeIsExpanded {
						if _b == system_ {
							found = true
							break
						}
					}
					if !found {
						_system.SubSystemsWhoseNodeIsExpanded = append(_system.SubSystemsWhoseNodeIsExpanded, system_)
						systemFormCallback.probe.UpdateSliceOfPointersCallback(_system, "SubSystemsWhoseNodeIsExpanded", &_system.SubSystemsWhoseNodeIsExpanded)
					}
				} else {
					// ensure system_ is NOT in _system.SubSystemsWhoseNodeIsExpanded
					idx := slices.Index(_system.SubSystemsWhoseNodeIsExpanded, system_)
					if idx != -1 {
						_system.SubSystemsWhoseNodeIsExpanded = slices.Delete(_system.SubSystemsWhoseNodeIsExpanded, idx, idx+1)
						systemFormCallback.probe.UpdateSliceOfPointersCallback(_system, "SubSystemsWhoseNodeIsExpanded", &_system.SubSystemsWhoseNodeIsExpanded)
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
