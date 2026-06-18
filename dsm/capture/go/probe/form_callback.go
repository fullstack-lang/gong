// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/capture/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__AnalysisNeedFormCallback(
	analysisneed *models.AnalysisNeed,
	probe *Probe,
	formGroup *form.FormGroup,
) (analysisneedFormCallback *AnalysisNeedFormCallback) {
	analysisneedFormCallback = new(AnalysisNeedFormCallback)
	analysisneedFormCallback.probe = probe
	analysisneedFormCallback.analysisneed = analysisneed
	analysisneedFormCallback.formGroup = formGroup

	analysisneedFormCallback.CreationMode = (analysisneed == nil)

	return
}

type AnalysisNeedFormCallback struct {
	analysisneed *models.AnalysisNeed

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (analysisneedFormCallback *AnalysisNeedFormCallback) OnSave() {
	analysisneedFormCallback.probe.stageOfInterest.Lock()
	defer analysisneedFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("AnalysisNeedFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	analysisneedFormCallback.probe.formStage.Checkout()

	if analysisneedFormCallback.analysisneed == nil {
		analysisneedFormCallback.analysisneed = new(models.AnalysisNeed).Stage(analysisneedFormCallback.probe.stageOfInterest)
	}
	analysisneed_ := analysisneedFormCallback.analysisneed
	_ = analysisneed_

	for _, formDiv := range analysisneedFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(analysisneed_.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(analysisneed_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(analysisneed_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(analysisneed_.LayoutDirection), formDiv)
		case "Library:AnalysisNeeds":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](analysisneedFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their AnalysisNeeds slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](analysisneedFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(analysisneedFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure analysisneed_ is in _library.AnalysisNeeds
					found := false
					for _, _b := range _library.AnalysisNeeds {
						if _b == analysisneed_ {
							found = true
							break
						}
					}
					if !found {
						_library.AnalysisNeeds = append(_library.AnalysisNeeds, analysisneed_)
						analysisneedFormCallback.probe.UpdateSliceOfPointersCallback(_library, "AnalysisNeeds", &_library.AnalysisNeeds)
					}
				} else {
					// ensure analysisneed_ is NOT in _library.AnalysisNeeds
					idx := slices.Index(_library.AnalysisNeeds, analysisneed_)
					if idx != -1 {
						_library.AnalysisNeeds = slices.Delete(_library.AnalysisNeeds, idx, idx+1)
						analysisneedFormCallback.probe.UpdateSliceOfPointersCallback(_library, "AnalysisNeeds", &_library.AnalysisNeeds)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if analysisneedFormCallback.formGroup.HasSuppressButtonBeenPressed {
		analysisneed_.Unstage(analysisneedFormCallback.probe.stageOfInterest)
	}

	analysisneedFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.AnalysisNeed](
		analysisneedFormCallback.probe,
	)

	// display a new form by reset the form stage
	if analysisneedFormCallback.CreationMode || analysisneedFormCallback.formGroup.HasSuppressButtonBeenPressed {
		analysisneedFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(analysisneedFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__AnalysisNeedFormCallback(
			nil,
			analysisneedFormCallback.probe,
			newFormGroup,
		)
		analysisneed := new(models.AnalysisNeed)
		FillUpForm(analysisneed, newFormGroup, analysisneedFormCallback.probe)
		analysisneedFormCallback.probe.formStage.Commit()
	}

	analysisneedFormCallback.probe.ux_tree()
}
func __gong__New__ConceptFormCallback(
	concept *models.Concept,
	probe *Probe,
	formGroup *form.FormGroup,
) (conceptFormCallback *ConceptFormCallback) {
	conceptFormCallback = new(ConceptFormCallback)
	conceptFormCallback.probe = probe
	conceptFormCallback.concept = concept
	conceptFormCallback.formGroup = formGroup

	conceptFormCallback.CreationMode = (concept == nil)

	return
}

type ConceptFormCallback struct {
	concept *models.Concept

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (conceptFormCallback *ConceptFormCallback) OnSave() {
	conceptFormCallback.probe.stageOfInterest.Lock()
	defer conceptFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ConceptFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	conceptFormCallback.probe.formStage.Checkout()

	if conceptFormCallback.concept == nil {
		conceptFormCallback.concept = new(models.Concept).Stage(conceptFormCallback.probe.stageOfInterest)
	}
	concept_ := conceptFormCallback.concept
	_ = concept_

	for _, formDiv := range conceptFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(concept_.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(concept_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(concept_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(concept_.LayoutDirection), formDiv)
		case "Tools":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Tool](conceptFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Tool, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Tool)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					conceptFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Tool](conceptFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			concept_.Tools = instanceSlice
			conceptFormCallback.probe.UpdateSliceOfPointersCallback(concept_, "Tools", &concept_.Tools)

		case "Deliverable:Concepts":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Deliverable instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Deliverable instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Deliverable](conceptFormCallback.probe.stageOfInterest)
			targetDeliverableIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDeliverableIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Deliverable instances and update their Concepts slice
			for _deliverable := range *models.GetGongstructInstancesSetFromPointerType[*models.Deliverable](conceptFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(conceptFormCallback.probe.stageOfInterest, _deliverable)
				
				// if Deliverable is selected
				if targetDeliverableIDs[id] {
					// ensure concept_ is in _deliverable.Concepts
					found := false
					for _, _b := range _deliverable.Concepts {
						if _b == concept_ {
							found = true
							break
						}
					}
					if !found {
						_deliverable.Concepts = append(_deliverable.Concepts, concept_)
						conceptFormCallback.probe.UpdateSliceOfPointersCallback(_deliverable, "Concepts", &_deliverable.Concepts)
					}
				} else {
					// ensure concept_ is NOT in _deliverable.Concepts
					idx := slices.Index(_deliverable.Concepts, concept_)
					if idx != -1 {
						_deliverable.Concepts = slices.Delete(_deliverable.Concepts, idx, idx+1)
						conceptFormCallback.probe.UpdateSliceOfPointersCallback(_deliverable, "Concepts", &_deliverable.Concepts)
					}
				}
			}
		case "Library:RootConcepts":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](conceptFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their RootConcepts slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](conceptFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(conceptFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure concept_ is in _library.RootConcepts
					found := false
					for _, _b := range _library.RootConcepts {
						if _b == concept_ {
							found = true
							break
						}
					}
					if !found {
						_library.RootConcepts = append(_library.RootConcepts, concept_)
						conceptFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootConcepts", &_library.RootConcepts)
					}
				} else {
					// ensure concept_ is NOT in _library.RootConcepts
					idx := slices.Index(_library.RootConcepts, concept_)
					if idx != -1 {
						_library.RootConcepts = slices.Delete(_library.RootConcepts, idx, idx+1)
						conceptFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootConcepts", &_library.RootConcepts)
					}
				}
			}
		case "Requirement:Concepts":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Requirement instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Requirement instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Requirement](conceptFormCallback.probe.stageOfInterest)
			targetRequirementIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetRequirementIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Requirement instances and update their Concepts slice
			for _requirement := range *models.GetGongstructInstancesSetFromPointerType[*models.Requirement](conceptFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(conceptFormCallback.probe.stageOfInterest, _requirement)
				
				// if Requirement is selected
				if targetRequirementIDs[id] {
					// ensure concept_ is in _requirement.Concepts
					found := false
					for _, _b := range _requirement.Concepts {
						if _b == concept_ {
							found = true
							break
						}
					}
					if !found {
						_requirement.Concepts = append(_requirement.Concepts, concept_)
						conceptFormCallback.probe.UpdateSliceOfPointersCallback(_requirement, "Concepts", &_requirement.Concepts)
					}
				} else {
					// ensure concept_ is NOT in _requirement.Concepts
					idx := slices.Index(_requirement.Concepts, concept_)
					if idx != -1 {
						_requirement.Concepts = slices.Delete(_requirement.Concepts, idx, idx+1)
						conceptFormCallback.probe.UpdateSliceOfPointersCallback(_requirement, "Concepts", &_requirement.Concepts)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if conceptFormCallback.formGroup.HasSuppressButtonBeenPressed {
		concept_.Unstage(conceptFormCallback.probe.stageOfInterest)
	}

	conceptFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Concept](
		conceptFormCallback.probe,
	)

	// display a new form by reset the form stage
	if conceptFormCallback.CreationMode || conceptFormCallback.formGroup.HasSuppressButtonBeenPressed {
		conceptFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(conceptFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ConceptFormCallback(
			nil,
			conceptFormCallback.probe,
			newFormGroup,
		)
		concept := new(models.Concept)
		FillUpForm(concept, newFormGroup, conceptFormCallback.probe)
		conceptFormCallback.probe.formStage.Commit()
	}

	conceptFormCallback.probe.ux_tree()
}
func __gong__New__ConcernFormCallback(
	concern *models.Concern,
	probe *Probe,
	formGroup *form.FormGroup,
) (concernFormCallback *ConcernFormCallback) {
	concernFormCallback = new(ConcernFormCallback)
	concernFormCallback.probe = probe
	concernFormCallback.concern = concern
	concernFormCallback.formGroup = formGroup

	concernFormCallback.CreationMode = (concern == nil)

	return
}

type ConcernFormCallback struct {
	concern *models.Concern

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (concernFormCallback *ConcernFormCallback) OnSave() {
	concernFormCallback.probe.stageOfInterest.Lock()
	defer concernFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ConcernFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	concernFormCallback.probe.formStage.Checkout()

	if concernFormCallback.concern == nil {
		concernFormCallback.concern = new(models.Concern).Stage(concernFormCallback.probe.stageOfInterest)
	}
	concern_ := concernFormCallback.concern
	_ = concern_

	for _, formDiv := range concernFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(concern_.Name), formDiv)
		case "IDAirbus":
			FormDivBasicFieldToField(&(concern_.IDAirbus), formDiv)
		case "Priority":
			FormDivEnumStringFieldToField(&(concern_.Priority), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(concern_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(concern_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(concern_.LayoutDirection), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(concern_.Description), formDiv)
		case "SubConcerns":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Concern](concernFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Concern, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Concern)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					concernFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Concern](concernFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			concern_.SubConcerns = instanceSlice
			concernFormCallback.probe.UpdateSliceOfPointersCallback(concern_, "SubConcerns", &concern_.SubConcerns)

		case "Inputs":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Deliverable](concernFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Deliverable, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Deliverable)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					concernFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Deliverable](concernFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			concern_.Inputs = instanceSlice
			concernFormCallback.probe.UpdateSliceOfPointersCallback(concern_, "Inputs", &concern_.Inputs)

		case "IsInputsNodeExpanded":
			FormDivBasicFieldToField(&(concern_.IsInputsNodeExpanded), formDiv)
		case "Outputs":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Deliverable](concernFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Deliverable, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Deliverable)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					concernFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Deliverable](concernFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			concern_.Outputs = instanceSlice
			concernFormCallback.probe.UpdateSliceOfPointersCallback(concern_, "Outputs", &concern_.Outputs)

		case "IsOutputsNodeExpanded":
			FormDivBasicFieldToField(&(concern_.IsOutputsNodeExpanded), formDiv)
		case "IsWithCompletion":
			FormDivBasicFieldToField(&(concern_.IsWithCompletion), formDiv)
		case "Completion":
			FormDivEnumStringFieldToField(&(concern_.Completion), formDiv)
		case "Requirements":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Requirement](concernFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Requirement, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Requirement)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					concernFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Requirement](concernFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			concern_.Requirements = instanceSlice
			concernFormCallback.probe.UpdateSliceOfPointersCallback(concern_, "Requirements", &concern_.Requirements)

		case "Concern:SubConcerns":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Concern instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Concern instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Concern](concernFormCallback.probe.stageOfInterest)
			targetConcernIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetConcernIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Concern instances and update their SubConcerns slice
			for _concern := range *models.GetGongstructInstancesSetFromPointerType[*models.Concern](concernFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(concernFormCallback.probe.stageOfInterest, _concern)
				
				// if Concern is selected
				if targetConcernIDs[id] {
					// ensure concern_ is in _concern.SubConcerns
					found := false
					for _, _b := range _concern.SubConcerns {
						if _b == concern_ {
							found = true
							break
						}
					}
					if !found {
						_concern.SubConcerns = append(_concern.SubConcerns, concern_)
						concernFormCallback.probe.UpdateSliceOfPointersCallback(_concern, "SubConcerns", &_concern.SubConcerns)
					}
				} else {
					// ensure concern_ is NOT in _concern.SubConcerns
					idx := slices.Index(_concern.SubConcerns, concern_)
					if idx != -1 {
						_concern.SubConcerns = slices.Delete(_concern.SubConcerns, idx, idx+1)
						concernFormCallback.probe.UpdateSliceOfPointersCallback(_concern, "SubConcerns", &_concern.SubConcerns)
					}
				}
			}
		case "Diagram:ConcernsWhoseRequirementsNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](concernFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their ConcernsWhoseRequirementsNodeIsExpanded slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](concernFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(concernFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure concern_ is in _diagram.ConcernsWhoseRequirementsNodeIsExpanded
					found := false
					for _, _b := range _diagram.ConcernsWhoseRequirementsNodeIsExpanded {
						if _b == concern_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.ConcernsWhoseRequirementsNodeIsExpanded = append(_diagram.ConcernsWhoseRequirementsNodeIsExpanded, concern_)
						concernFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ConcernsWhoseRequirementsNodeIsExpanded", &_diagram.ConcernsWhoseRequirementsNodeIsExpanded)
					}
				} else {
					// ensure concern_ is NOT in _diagram.ConcernsWhoseRequirementsNodeIsExpanded
					idx := slices.Index(_diagram.ConcernsWhoseRequirementsNodeIsExpanded, concern_)
					if idx != -1 {
						_diagram.ConcernsWhoseRequirementsNodeIsExpanded = slices.Delete(_diagram.ConcernsWhoseRequirementsNodeIsExpanded, idx, idx+1)
						concernFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ConcernsWhoseRequirementsNodeIsExpanded", &_diagram.ConcernsWhoseRequirementsNodeIsExpanded)
					}
				}
			}
		case "Diagram:ConcernsWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](concernFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their ConcernsWhoseNodeIsExpanded slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](concernFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(concernFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure concern_ is in _diagram.ConcernsWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagram.ConcernsWhoseNodeIsExpanded {
						if _b == concern_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.ConcernsWhoseNodeIsExpanded = append(_diagram.ConcernsWhoseNodeIsExpanded, concern_)
						concernFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ConcernsWhoseNodeIsExpanded", &_diagram.ConcernsWhoseNodeIsExpanded)
					}
				} else {
					// ensure concern_ is NOT in _diagram.ConcernsWhoseNodeIsExpanded
					idx := slices.Index(_diagram.ConcernsWhoseNodeIsExpanded, concern_)
					if idx != -1 {
						_diagram.ConcernsWhoseNodeIsExpanded = slices.Delete(_diagram.ConcernsWhoseNodeIsExpanded, idx, idx+1)
						concernFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ConcernsWhoseNodeIsExpanded", &_diagram.ConcernsWhoseNodeIsExpanded)
					}
				}
			}
		case "Diagram:ConcernsWhoseInputNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](concernFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their ConcernsWhoseInputNodeIsExpanded slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](concernFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(concernFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure concern_ is in _diagram.ConcernsWhoseInputNodeIsExpanded
					found := false
					for _, _b := range _diagram.ConcernsWhoseInputNodeIsExpanded {
						if _b == concern_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.ConcernsWhoseInputNodeIsExpanded = append(_diagram.ConcernsWhoseInputNodeIsExpanded, concern_)
						concernFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ConcernsWhoseInputNodeIsExpanded", &_diagram.ConcernsWhoseInputNodeIsExpanded)
					}
				} else {
					// ensure concern_ is NOT in _diagram.ConcernsWhoseInputNodeIsExpanded
					idx := slices.Index(_diagram.ConcernsWhoseInputNodeIsExpanded, concern_)
					if idx != -1 {
						_diagram.ConcernsWhoseInputNodeIsExpanded = slices.Delete(_diagram.ConcernsWhoseInputNodeIsExpanded, idx, idx+1)
						concernFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ConcernsWhoseInputNodeIsExpanded", &_diagram.ConcernsWhoseInputNodeIsExpanded)
					}
				}
			}
		case "Diagram:ConcernsWhoseStakeholderNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](concernFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their ConcernsWhoseStakeholderNodeIsExpanded slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](concernFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(concernFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure concern_ is in _diagram.ConcernsWhoseStakeholderNodeIsExpanded
					found := false
					for _, _b := range _diagram.ConcernsWhoseStakeholderNodeIsExpanded {
						if _b == concern_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.ConcernsWhoseStakeholderNodeIsExpanded = append(_diagram.ConcernsWhoseStakeholderNodeIsExpanded, concern_)
						concernFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ConcernsWhoseStakeholderNodeIsExpanded", &_diagram.ConcernsWhoseStakeholderNodeIsExpanded)
					}
				} else {
					// ensure concern_ is NOT in _diagram.ConcernsWhoseStakeholderNodeIsExpanded
					idx := slices.Index(_diagram.ConcernsWhoseStakeholderNodeIsExpanded, concern_)
					if idx != -1 {
						_diagram.ConcernsWhoseStakeholderNodeIsExpanded = slices.Delete(_diagram.ConcernsWhoseStakeholderNodeIsExpanded, idx, idx+1)
						concernFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ConcernsWhoseStakeholderNodeIsExpanded", &_diagram.ConcernsWhoseStakeholderNodeIsExpanded)
					}
				}
			}
		case "Diagram:ConcernssWhoseOutputNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](concernFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their ConcernssWhoseOutputNodeIsExpanded slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](concernFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(concernFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure concern_ is in _diagram.ConcernssWhoseOutputNodeIsExpanded
					found := false
					for _, _b := range _diagram.ConcernssWhoseOutputNodeIsExpanded {
						if _b == concern_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.ConcernssWhoseOutputNodeIsExpanded = append(_diagram.ConcernssWhoseOutputNodeIsExpanded, concern_)
						concernFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ConcernssWhoseOutputNodeIsExpanded", &_diagram.ConcernssWhoseOutputNodeIsExpanded)
					}
				} else {
					// ensure concern_ is NOT in _diagram.ConcernssWhoseOutputNodeIsExpanded
					idx := slices.Index(_diagram.ConcernssWhoseOutputNodeIsExpanded, concern_)
					if idx != -1 {
						_diagram.ConcernssWhoseOutputNodeIsExpanded = slices.Delete(_diagram.ConcernssWhoseOutputNodeIsExpanded, idx, idx+1)
						concernFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ConcernssWhoseOutputNodeIsExpanded", &_diagram.ConcernssWhoseOutputNodeIsExpanded)
					}
				}
			}
		case "Library:RootConcerns":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](concernFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their RootConcerns slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](concernFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(concernFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure concern_ is in _library.RootConcerns
					found := false
					for _, _b := range _library.RootConcerns {
						if _b == concern_ {
							found = true
							break
						}
					}
					if !found {
						_library.RootConcerns = append(_library.RootConcerns, concern_)
						concernFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootConcerns", &_library.RootConcerns)
					}
				} else {
					// ensure concern_ is NOT in _library.RootConcerns
					idx := slices.Index(_library.RootConcerns, concern_)
					if idx != -1 {
						_library.RootConcerns = slices.Delete(_library.RootConcerns, idx, idx+1)
						concernFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootConcerns", &_library.RootConcerns)
					}
				}
			}
		case "Note:Tasks":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Note instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Note instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Note](concernFormCallback.probe.stageOfInterest)
			targetNoteIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetNoteIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Note instances and update their Tasks slice
			for _note := range *models.GetGongstructInstancesSetFromPointerType[*models.Note](concernFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(concernFormCallback.probe.stageOfInterest, _note)
				
				// if Note is selected
				if targetNoteIDs[id] {
					// ensure concern_ is in _note.Tasks
					found := false
					for _, _b := range _note.Tasks {
						if _b == concern_ {
							found = true
							break
						}
					}
					if !found {
						_note.Tasks = append(_note.Tasks, concern_)
						concernFormCallback.probe.UpdateSliceOfPointersCallback(_note, "Tasks", &_note.Tasks)
					}
				} else {
					// ensure concern_ is NOT in _note.Tasks
					idx := slices.Index(_note.Tasks, concern_)
					if idx != -1 {
						_note.Tasks = slices.Delete(_note.Tasks, idx, idx+1)
						concernFormCallback.probe.UpdateSliceOfPointersCallback(_note, "Tasks", &_note.Tasks)
					}
				}
			}
		case "Stakeholder:Concerns":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Stakeholder instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Stakeholder instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Stakeholder](concernFormCallback.probe.stageOfInterest)
			targetStakeholderIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetStakeholderIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Stakeholder instances and update their Concerns slice
			for _stakeholder := range *models.GetGongstructInstancesSetFromPointerType[*models.Stakeholder](concernFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(concernFormCallback.probe.stageOfInterest, _stakeholder)
				
				// if Stakeholder is selected
				if targetStakeholderIDs[id] {
					// ensure concern_ is in _stakeholder.Concerns
					found := false
					for _, _b := range _stakeholder.Concerns {
						if _b == concern_ {
							found = true
							break
						}
					}
					if !found {
						_stakeholder.Concerns = append(_stakeholder.Concerns, concern_)
						concernFormCallback.probe.UpdateSliceOfPointersCallback(_stakeholder, "Concerns", &_stakeholder.Concerns)
					}
				} else {
					// ensure concern_ is NOT in _stakeholder.Concerns
					idx := slices.Index(_stakeholder.Concerns, concern_)
					if idx != -1 {
						_stakeholder.Concerns = slices.Delete(_stakeholder.Concerns, idx, idx+1)
						concernFormCallback.probe.UpdateSliceOfPointersCallback(_stakeholder, "Concerns", &_stakeholder.Concerns)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if concernFormCallback.formGroup.HasSuppressButtonBeenPressed {
		concern_.Unstage(concernFormCallback.probe.stageOfInterest)
	}

	concernFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Concern](
		concernFormCallback.probe,
	)

	// display a new form by reset the form stage
	if concernFormCallback.CreationMode || concernFormCallback.formGroup.HasSuppressButtonBeenPressed {
		concernFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(concernFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ConcernFormCallback(
			nil,
			concernFormCallback.probe,
			newFormGroup,
		)
		concern := new(models.Concern)
		FillUpForm(concern, newFormGroup, concernFormCallback.probe)
		concernFormCallback.probe.formStage.Commit()
	}

	concernFormCallback.probe.ux_tree()
}
func __gong__New__ConcernCompositionShapeFormCallback(
	concerncompositionshape *models.ConcernCompositionShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (concerncompositionshapeFormCallback *ConcernCompositionShapeFormCallback) {
	concerncompositionshapeFormCallback = new(ConcernCompositionShapeFormCallback)
	concerncompositionshapeFormCallback.probe = probe
	concerncompositionshapeFormCallback.concerncompositionshape = concerncompositionshape
	concerncompositionshapeFormCallback.formGroup = formGroup

	concerncompositionshapeFormCallback.CreationMode = (concerncompositionshape == nil)

	return
}

type ConcernCompositionShapeFormCallback struct {
	concerncompositionshape *models.ConcernCompositionShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (concerncompositionshapeFormCallback *ConcernCompositionShapeFormCallback) OnSave() {
	concerncompositionshapeFormCallback.probe.stageOfInterest.Lock()
	defer concerncompositionshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ConcernCompositionShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	concerncompositionshapeFormCallback.probe.formStage.Checkout()

	if concerncompositionshapeFormCallback.concerncompositionshape == nil {
		concerncompositionshapeFormCallback.concerncompositionshape = new(models.ConcernCompositionShape).Stage(concerncompositionshapeFormCallback.probe.stageOfInterest)
	}
	concerncompositionshape_ := concerncompositionshapeFormCallback.concerncompositionshape
	_ = concerncompositionshape_

	for _, formDiv := range concerncompositionshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(concerncompositionshape_.Name), formDiv)
		case "Concern":
			FormDivSelectFieldToField(&(concerncompositionshape_.Concern), concerncompositionshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(concerncompositionshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(concerncompositionshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(concerncompositionshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(concerncompositionshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(concerncompositionshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(concerncompositionshape_.IsHidden), formDiv)
		case "Diagram:ConcernComposition_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](concerncompositionshapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their ConcernComposition_Shapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](concerncompositionshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(concerncompositionshapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure concerncompositionshape_ is in _diagram.ConcernComposition_Shapes
					found := false
					for _, _b := range _diagram.ConcernComposition_Shapes {
						if _b == concerncompositionshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.ConcernComposition_Shapes = append(_diagram.ConcernComposition_Shapes, concerncompositionshape_)
						concerncompositionshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ConcernComposition_Shapes", &_diagram.ConcernComposition_Shapes)
					}
				} else {
					// ensure concerncompositionshape_ is NOT in _diagram.ConcernComposition_Shapes
					idx := slices.Index(_diagram.ConcernComposition_Shapes, concerncompositionshape_)
					if idx != -1 {
						_diagram.ConcernComposition_Shapes = slices.Delete(_diagram.ConcernComposition_Shapes, idx, idx+1)
						concerncompositionshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ConcernComposition_Shapes", &_diagram.ConcernComposition_Shapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if concerncompositionshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		concerncompositionshape_.Unstage(concerncompositionshapeFormCallback.probe.stageOfInterest)
	}

	concerncompositionshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ConcernCompositionShape](
		concerncompositionshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if concerncompositionshapeFormCallback.CreationMode || concerncompositionshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		concerncompositionshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(concerncompositionshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ConcernCompositionShapeFormCallback(
			nil,
			concerncompositionshapeFormCallback.probe,
			newFormGroup,
		)
		concerncompositionshape := new(models.ConcernCompositionShape)
		FillUpForm(concerncompositionshape, newFormGroup, concerncompositionshapeFormCallback.probe)
		concerncompositionshapeFormCallback.probe.formStage.Commit()
	}

	concerncompositionshapeFormCallback.probe.ux_tree()
}
func __gong__New__ConcernInputShapeFormCallback(
	concerninputshape *models.ConcernInputShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (concerninputshapeFormCallback *ConcernInputShapeFormCallback) {
	concerninputshapeFormCallback = new(ConcernInputShapeFormCallback)
	concerninputshapeFormCallback.probe = probe
	concerninputshapeFormCallback.concerninputshape = concerninputshape
	concerninputshapeFormCallback.formGroup = formGroup

	concerninputshapeFormCallback.CreationMode = (concerninputshape == nil)

	return
}

type ConcernInputShapeFormCallback struct {
	concerninputshape *models.ConcernInputShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (concerninputshapeFormCallback *ConcernInputShapeFormCallback) OnSave() {
	concerninputshapeFormCallback.probe.stageOfInterest.Lock()
	defer concerninputshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ConcernInputShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	concerninputshapeFormCallback.probe.formStage.Checkout()

	if concerninputshapeFormCallback.concerninputshape == nil {
		concerninputshapeFormCallback.concerninputshape = new(models.ConcernInputShape).Stage(concerninputshapeFormCallback.probe.stageOfInterest)
	}
	concerninputshape_ := concerninputshapeFormCallback.concerninputshape
	_ = concerninputshape_

	for _, formDiv := range concerninputshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(concerninputshape_.Name), formDiv)
		case "Deliverable":
			FormDivSelectFieldToField(&(concerninputshape_.Deliverable), concerninputshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Concern":
			FormDivSelectFieldToField(&(concerninputshape_.Concern), concerninputshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(concerninputshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(concerninputshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(concerninputshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(concerninputshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(concerninputshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(concerninputshape_.IsHidden), formDiv)
		case "Diagram:ConcernInputShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](concerninputshapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their ConcernInputShapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](concerninputshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(concerninputshapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure concerninputshape_ is in _diagram.ConcernInputShapes
					found := false
					for _, _b := range _diagram.ConcernInputShapes {
						if _b == concerninputshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.ConcernInputShapes = append(_diagram.ConcernInputShapes, concerninputshape_)
						concerninputshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ConcernInputShapes", &_diagram.ConcernInputShapes)
					}
				} else {
					// ensure concerninputshape_ is NOT in _diagram.ConcernInputShapes
					idx := slices.Index(_diagram.ConcernInputShapes, concerninputshape_)
					if idx != -1 {
						_diagram.ConcernInputShapes = slices.Delete(_diagram.ConcernInputShapes, idx, idx+1)
						concerninputshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ConcernInputShapes", &_diagram.ConcernInputShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if concerninputshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		concerninputshape_.Unstage(concerninputshapeFormCallback.probe.stageOfInterest)
	}

	concerninputshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ConcernInputShape](
		concerninputshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if concerninputshapeFormCallback.CreationMode || concerninputshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		concerninputshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(concerninputshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ConcernInputShapeFormCallback(
			nil,
			concerninputshapeFormCallback.probe,
			newFormGroup,
		)
		concerninputshape := new(models.ConcernInputShape)
		FillUpForm(concerninputshape, newFormGroup, concerninputshapeFormCallback.probe)
		concerninputshapeFormCallback.probe.formStage.Commit()
	}

	concerninputshapeFormCallback.probe.ux_tree()
}
func __gong__New__ConcernOutputShapeFormCallback(
	concernoutputshape *models.ConcernOutputShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (concernoutputshapeFormCallback *ConcernOutputShapeFormCallback) {
	concernoutputshapeFormCallback = new(ConcernOutputShapeFormCallback)
	concernoutputshapeFormCallback.probe = probe
	concernoutputshapeFormCallback.concernoutputshape = concernoutputshape
	concernoutputshapeFormCallback.formGroup = formGroup

	concernoutputshapeFormCallback.CreationMode = (concernoutputshape == nil)

	return
}

type ConcernOutputShapeFormCallback struct {
	concernoutputshape *models.ConcernOutputShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (concernoutputshapeFormCallback *ConcernOutputShapeFormCallback) OnSave() {
	concernoutputshapeFormCallback.probe.stageOfInterest.Lock()
	defer concernoutputshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ConcernOutputShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	concernoutputshapeFormCallback.probe.formStage.Checkout()

	if concernoutputshapeFormCallback.concernoutputshape == nil {
		concernoutputshapeFormCallback.concernoutputshape = new(models.ConcernOutputShape).Stage(concernoutputshapeFormCallback.probe.stageOfInterest)
	}
	concernoutputshape_ := concernoutputshapeFormCallback.concernoutputshape
	_ = concernoutputshape_

	for _, formDiv := range concernoutputshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(concernoutputshape_.Name), formDiv)
		case "Task":
			FormDivSelectFieldToField(&(concernoutputshape_.Task), concernoutputshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Product":
			FormDivSelectFieldToField(&(concernoutputshape_.Product), concernoutputshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(concernoutputshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(concernoutputshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(concernoutputshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(concernoutputshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(concernoutputshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(concernoutputshape_.IsHidden), formDiv)
		case "Diagram:ConcernOutputShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](concernoutputshapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their ConcernOutputShapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](concernoutputshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(concernoutputshapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure concernoutputshape_ is in _diagram.ConcernOutputShapes
					found := false
					for _, _b := range _diagram.ConcernOutputShapes {
						if _b == concernoutputshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.ConcernOutputShapes = append(_diagram.ConcernOutputShapes, concernoutputshape_)
						concernoutputshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ConcernOutputShapes", &_diagram.ConcernOutputShapes)
					}
				} else {
					// ensure concernoutputshape_ is NOT in _diagram.ConcernOutputShapes
					idx := slices.Index(_diagram.ConcernOutputShapes, concernoutputshape_)
					if idx != -1 {
						_diagram.ConcernOutputShapes = slices.Delete(_diagram.ConcernOutputShapes, idx, idx+1)
						concernoutputshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ConcernOutputShapes", &_diagram.ConcernOutputShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if concernoutputshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		concernoutputshape_.Unstage(concernoutputshapeFormCallback.probe.stageOfInterest)
	}

	concernoutputshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ConcernOutputShape](
		concernoutputshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if concernoutputshapeFormCallback.CreationMode || concernoutputshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		concernoutputshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(concernoutputshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ConcernOutputShapeFormCallback(
			nil,
			concernoutputshapeFormCallback.probe,
			newFormGroup,
		)
		concernoutputshape := new(models.ConcernOutputShape)
		FillUpForm(concernoutputshape, newFormGroup, concernoutputshapeFormCallback.probe)
		concernoutputshapeFormCallback.probe.formStage.Commit()
	}

	concernoutputshapeFormCallback.probe.ux_tree()
}
func __gong__New__ConcernShapeFormCallback(
	concernshape *models.ConcernShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (concernshapeFormCallback *ConcernShapeFormCallback) {
	concernshapeFormCallback = new(ConcernShapeFormCallback)
	concernshapeFormCallback.probe = probe
	concernshapeFormCallback.concernshape = concernshape
	concernshapeFormCallback.formGroup = formGroup

	concernshapeFormCallback.CreationMode = (concernshape == nil)

	return
}

type ConcernShapeFormCallback struct {
	concernshape *models.ConcernShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (concernshapeFormCallback *ConcernShapeFormCallback) OnSave() {
	concernshapeFormCallback.probe.stageOfInterest.Lock()
	defer concernshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ConcernShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	concernshapeFormCallback.probe.formStage.Checkout()

	if concernshapeFormCallback.concernshape == nil {
		concernshapeFormCallback.concernshape = new(models.ConcernShape).Stage(concernshapeFormCallback.probe.stageOfInterest)
	}
	concernshape_ := concernshapeFormCallback.concernshape
	_ = concernshape_

	for _, formDiv := range concernshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(concernshape_.Name), formDiv)
		case "Concern":
			FormDivSelectFieldToField(&(concernshape_.Concern), concernshapeFormCallback.probe.stageOfInterest, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(concernshape_.IsExpanded), formDiv)
		case "X":
			FormDivBasicFieldToField(&(concernshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(concernshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(concernshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(concernshape_.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(concernshape_.IsHidden), formDiv)
		case "Diagram:Concern_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](concernshapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their Concern_Shapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](concernshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(concernshapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure concernshape_ is in _diagram.Concern_Shapes
					found := false
					for _, _b := range _diagram.Concern_Shapes {
						if _b == concernshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.Concern_Shapes = append(_diagram.Concern_Shapes, concernshape_)
						concernshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "Concern_Shapes", &_diagram.Concern_Shapes)
					}
				} else {
					// ensure concernshape_ is NOT in _diagram.Concern_Shapes
					idx := slices.Index(_diagram.Concern_Shapes, concernshape_)
					if idx != -1 {
						_diagram.Concern_Shapes = slices.Delete(_diagram.Concern_Shapes, idx, idx+1)
						concernshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "Concern_Shapes", &_diagram.Concern_Shapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if concernshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		concernshape_.Unstage(concernshapeFormCallback.probe.stageOfInterest)
	}

	concernshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ConcernShape](
		concernshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if concernshapeFormCallback.CreationMode || concernshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		concernshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(concernshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ConcernShapeFormCallback(
			nil,
			concernshapeFormCallback.probe,
			newFormGroup,
		)
		concernshape := new(models.ConcernShape)
		FillUpForm(concernshape, newFormGroup, concernshapeFormCallback.probe)
		concernshapeFormCallback.probe.formStage.Commit()
	}

	concernshapeFormCallback.probe.ux_tree()
}
func __gong__New__DeliverableFormCallback(
	deliverable *models.Deliverable,
	probe *Probe,
	formGroup *form.FormGroup,
) (deliverableFormCallback *DeliverableFormCallback) {
	deliverableFormCallback = new(DeliverableFormCallback)
	deliverableFormCallback.probe = probe
	deliverableFormCallback.deliverable = deliverable
	deliverableFormCallback.formGroup = formGroup

	deliverableFormCallback.CreationMode = (deliverable == nil)

	return
}

type DeliverableFormCallback struct {
	deliverable *models.Deliverable

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (deliverableFormCallback *DeliverableFormCallback) OnSave() {
	deliverableFormCallback.probe.stageOfInterest.Lock()
	defer deliverableFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DeliverableFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	deliverableFormCallback.probe.formStage.Checkout()

	if deliverableFormCallback.deliverable == nil {
		deliverableFormCallback.deliverable = new(models.Deliverable).Stage(deliverableFormCallback.probe.stageOfInterest)
	}
	deliverable_ := deliverableFormCallback.deliverable
	_ = deliverable_

	for _, formDiv := range deliverableFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(deliverable_.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(deliverable_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(deliverable_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(deliverable_.LayoutDirection), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(deliverable_.Description), formDiv)
		case "SubProducts":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Deliverable](deliverableFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Deliverable, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Deliverable)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					deliverableFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Deliverable](deliverableFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			deliverable_.SubProducts = instanceSlice
			deliverableFormCallback.probe.UpdateSliceOfPointersCallback(deliverable_, "SubProducts", &deliverable_.SubProducts)

		case "IsProducersNodeExpanded":
			FormDivBasicFieldToField(&(deliverable_.IsProducersNodeExpanded), formDiv)
		case "IsConsumersNodeExpanded":
			FormDivBasicFieldToField(&(deliverable_.IsConsumersNodeExpanded), formDiv)
		case "Concepts":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Concept](deliverableFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Concept, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Concept)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					deliverableFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Concept](deliverableFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			deliverable_.Concepts = instanceSlice
			deliverableFormCallback.probe.UpdateSliceOfPointersCallback(deliverable_, "Concepts", &deliverable_.Concepts)

		case "Concern:Inputs":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Concern instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Concern instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Concern](deliverableFormCallback.probe.stageOfInterest)
			targetConcernIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetConcernIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Concern instances and update their Inputs slice
			for _concern := range *models.GetGongstructInstancesSetFromPointerType[*models.Concern](deliverableFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(deliverableFormCallback.probe.stageOfInterest, _concern)
				
				// if Concern is selected
				if targetConcernIDs[id] {
					// ensure deliverable_ is in _concern.Inputs
					found := false
					for _, _b := range _concern.Inputs {
						if _b == deliverable_ {
							found = true
							break
						}
					}
					if !found {
						_concern.Inputs = append(_concern.Inputs, deliverable_)
						deliverableFormCallback.probe.UpdateSliceOfPointersCallback(_concern, "Inputs", &_concern.Inputs)
					}
				} else {
					// ensure deliverable_ is NOT in _concern.Inputs
					idx := slices.Index(_concern.Inputs, deliverable_)
					if idx != -1 {
						_concern.Inputs = slices.Delete(_concern.Inputs, idx, idx+1)
						deliverableFormCallback.probe.UpdateSliceOfPointersCallback(_concern, "Inputs", &_concern.Inputs)
					}
				}
			}
		case "Concern:Outputs":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Concern instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Concern instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Concern](deliverableFormCallback.probe.stageOfInterest)
			targetConcernIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetConcernIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Concern instances and update their Outputs slice
			for _concern := range *models.GetGongstructInstancesSetFromPointerType[*models.Concern](deliverableFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(deliverableFormCallback.probe.stageOfInterest, _concern)
				
				// if Concern is selected
				if targetConcernIDs[id] {
					// ensure deliverable_ is in _concern.Outputs
					found := false
					for _, _b := range _concern.Outputs {
						if _b == deliverable_ {
							found = true
							break
						}
					}
					if !found {
						_concern.Outputs = append(_concern.Outputs, deliverable_)
						deliverableFormCallback.probe.UpdateSliceOfPointersCallback(_concern, "Outputs", &_concern.Outputs)
					}
				} else {
					// ensure deliverable_ is NOT in _concern.Outputs
					idx := slices.Index(_concern.Outputs, deliverable_)
					if idx != -1 {
						_concern.Outputs = slices.Delete(_concern.Outputs, idx, idx+1)
						deliverableFormCallback.probe.UpdateSliceOfPointersCallback(_concern, "Outputs", &_concern.Outputs)
					}
				}
			}
		case "Deliverable:SubProducts":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Deliverable instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Deliverable instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Deliverable](deliverableFormCallback.probe.stageOfInterest)
			targetDeliverableIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDeliverableIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Deliverable instances and update their SubProducts slice
			for _deliverable := range *models.GetGongstructInstancesSetFromPointerType[*models.Deliverable](deliverableFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(deliverableFormCallback.probe.stageOfInterest, _deliverable)
				
				// if Deliverable is selected
				if targetDeliverableIDs[id] {
					// ensure deliverable_ is in _deliverable.SubProducts
					found := false
					for _, _b := range _deliverable.SubProducts {
						if _b == deliverable_ {
							found = true
							break
						}
					}
					if !found {
						_deliverable.SubProducts = append(_deliverable.SubProducts, deliverable_)
						deliverableFormCallback.probe.UpdateSliceOfPointersCallback(_deliverable, "SubProducts", &_deliverable.SubProducts)
					}
				} else {
					// ensure deliverable_ is NOT in _deliverable.SubProducts
					idx := slices.Index(_deliverable.SubProducts, deliverable_)
					if idx != -1 {
						_deliverable.SubProducts = slices.Delete(_deliverable.SubProducts, idx, idx+1)
						deliverableFormCallback.probe.UpdateSliceOfPointersCallback(_deliverable, "SubProducts", &_deliverable.SubProducts)
					}
				}
			}
		case "Diagram:ProductsWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](deliverableFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their ProductsWhoseNodeIsExpanded slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](deliverableFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(deliverableFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure deliverable_ is in _diagram.ProductsWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagram.ProductsWhoseNodeIsExpanded {
						if _b == deliverable_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.ProductsWhoseNodeIsExpanded = append(_diagram.ProductsWhoseNodeIsExpanded, deliverable_)
						deliverableFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ProductsWhoseNodeIsExpanded", &_diagram.ProductsWhoseNodeIsExpanded)
					}
				} else {
					// ensure deliverable_ is NOT in _diagram.ProductsWhoseNodeIsExpanded
					idx := slices.Index(_diagram.ProductsWhoseNodeIsExpanded, deliverable_)
					if idx != -1 {
						_diagram.ProductsWhoseNodeIsExpanded = slices.Delete(_diagram.ProductsWhoseNodeIsExpanded, idx, idx+1)
						deliverableFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ProductsWhoseNodeIsExpanded", &_diagram.ProductsWhoseNodeIsExpanded)
					}
				}
			}
		case "Library:RootDeliverables":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](deliverableFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their RootDeliverables slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](deliverableFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(deliverableFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure deliverable_ is in _library.RootDeliverables
					found := false
					for _, _b := range _library.RootDeliverables {
						if _b == deliverable_ {
							found = true
							break
						}
					}
					if !found {
						_library.RootDeliverables = append(_library.RootDeliverables, deliverable_)
						deliverableFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootDeliverables", &_library.RootDeliverables)
					}
				} else {
					// ensure deliverable_ is NOT in _library.RootDeliverables
					idx := slices.Index(_library.RootDeliverables, deliverable_)
					if idx != -1 {
						_library.RootDeliverables = slices.Delete(_library.RootDeliverables, idx, idx+1)
						deliverableFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootDeliverables", &_library.RootDeliverables)
					}
				}
			}
		case "Note:Products":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Note instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Note instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Note](deliverableFormCallback.probe.stageOfInterest)
			targetNoteIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetNoteIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Note instances and update their Products slice
			for _note := range *models.GetGongstructInstancesSetFromPointerType[*models.Note](deliverableFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(deliverableFormCallback.probe.stageOfInterest, _note)
				
				// if Note is selected
				if targetNoteIDs[id] {
					// ensure deliverable_ is in _note.Products
					found := false
					for _, _b := range _note.Products {
						if _b == deliverable_ {
							found = true
							break
						}
					}
					if !found {
						_note.Products = append(_note.Products, deliverable_)
						deliverableFormCallback.probe.UpdateSliceOfPointersCallback(_note, "Products", &_note.Products)
					}
				} else {
					// ensure deliverable_ is NOT in _note.Products
					idx := slices.Index(_note.Products, deliverable_)
					if idx != -1 {
						_note.Products = slices.Delete(_note.Products, idx, idx+1)
						deliverableFormCallback.probe.UpdateSliceOfPointersCallback(_note, "Products", &_note.Products)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if deliverableFormCallback.formGroup.HasSuppressButtonBeenPressed {
		deliverable_.Unstage(deliverableFormCallback.probe.stageOfInterest)
	}

	deliverableFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Deliverable](
		deliverableFormCallback.probe,
	)

	// display a new form by reset the form stage
	if deliverableFormCallback.CreationMode || deliverableFormCallback.formGroup.HasSuppressButtonBeenPressed {
		deliverableFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(deliverableFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DeliverableFormCallback(
			nil,
			deliverableFormCallback.probe,
			newFormGroup,
		)
		deliverable := new(models.Deliverable)
		FillUpForm(deliverable, newFormGroup, deliverableFormCallback.probe)
		deliverableFormCallback.probe.formStage.Commit()
	}

	deliverableFormCallback.probe.ux_tree()
}
func __gong__New__DiagramFormCallback(
	diagram *models.Diagram,
	probe *Probe,
	formGroup *form.FormGroup,
) (diagramFormCallback *DiagramFormCallback) {
	diagramFormCallback = new(DiagramFormCallback)
	diagramFormCallback.probe = probe
	diagramFormCallback.diagram = diagram
	diagramFormCallback.formGroup = formGroup

	diagramFormCallback.CreationMode = (diagram == nil)

	return
}

type DiagramFormCallback struct {
	diagram *models.Diagram

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (diagramFormCallback *DiagramFormCallback) OnSave() {
	diagramFormCallback.probe.stageOfInterest.Lock()
	defer diagramFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DiagramFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	diagramFormCallback.probe.formStage.Checkout()

	if diagramFormCallback.diagram == nil {
		diagramFormCallback.diagram = new(models.Diagram).Stage(diagramFormCallback.probe.stageOfInterest)
	}
	diagram_ := diagramFormCallback.diagram
	_ = diagram_

	for _, formDiv := range diagramFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(diagram_.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(diagram_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(diagram_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(diagram_.LayoutDirection), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(diagram_.IsChecked), formDiv)
		case "IsEditable_":
			FormDivBasicFieldToField(&(diagram_.IsEditable_), formDiv)
		case "ShowPrefix":
			FormDivBasicFieldToField(&(diagram_.ShowPrefix), formDiv)
		case "DefaultBoxWidth":
			FormDivBasicFieldToField(&(diagram_.DefaultBoxWidth), formDiv)
		case "DefaultBoxHeigth":
			FormDivBasicFieldToField(&(diagram_.DefaultBoxHeigth), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(diagram_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(diagram_.Height), formDiv)
		case "ConcernsWhoseRequirementsNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Concern](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Concern, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Concern)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Concern](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ConcernsWhoseRequirementsNodeIsExpanded = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "ConcernsWhoseRequirementsNodeIsExpanded", &diagram_.ConcernsWhoseRequirementsNodeIsExpanded)

		case "IsRequirementsNodeExpanded":
			FormDivBasicFieldToField(&(diagram_.IsRequirementsNodeExpanded), formDiv)
		case "IsConceptsNodeExpanded":
			FormDivBasicFieldToField(&(diagram_.IsConceptsNodeExpanded), formDiv)
		case "Product_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ProductShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ProductShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ProductShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ProductShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.Product_Shapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "Product_Shapes", &diagram_.Product_Shapes)

		case "ProductsWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Deliverable](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Deliverable, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Deliverable)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Deliverable](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ProductsWhoseNodeIsExpanded = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "ProductsWhoseNodeIsExpanded", &diagram_.ProductsWhoseNodeIsExpanded)

		case "IsPBSNodeExpanded":
			FormDivBasicFieldToField(&(diagram_.IsPBSNodeExpanded), formDiv)
		case "ProductComposition_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ProductCompositionShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ProductCompositionShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ProductCompositionShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ProductCompositionShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ProductComposition_Shapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "ProductComposition_Shapes", &diagram_.ProductComposition_Shapes)

		case "IsConcernsNodeExpanded":
			FormDivBasicFieldToField(&(diagram_.IsConcernsNodeExpanded), formDiv)
		case "Concern_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ConcernShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ConcernShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ConcernShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ConcernShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.Concern_Shapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "Concern_Shapes", &diagram_.Concern_Shapes)

		case "ConcernsWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Concern](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Concern, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Concern)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Concern](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ConcernsWhoseNodeIsExpanded = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "ConcernsWhoseNodeIsExpanded", &diagram_.ConcernsWhoseNodeIsExpanded)

		case "ConcernsWhoseInputNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Concern](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Concern, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Concern)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Concern](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ConcernsWhoseInputNodeIsExpanded = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "ConcernsWhoseInputNodeIsExpanded", &diagram_.ConcernsWhoseInputNodeIsExpanded)

		case "ConcernsWhoseStakeholderNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Concern](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Concern, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Concern)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Concern](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ConcernsWhoseStakeholderNodeIsExpanded = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "ConcernsWhoseStakeholderNodeIsExpanded", &diagram_.ConcernsWhoseStakeholderNodeIsExpanded)

		case "ConcernssWhoseOutputNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Concern](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Concern, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Concern)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Concern](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ConcernssWhoseOutputNodeIsExpanded = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "ConcernssWhoseOutputNodeIsExpanded", &diagram_.ConcernssWhoseOutputNodeIsExpanded)

		case "ConcernComposition_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ConcernCompositionShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ConcernCompositionShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ConcernCompositionShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ConcernCompositionShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ConcernComposition_Shapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "ConcernComposition_Shapes", &diagram_.ConcernComposition_Shapes)

		case "ConcernInputShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ConcernInputShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ConcernInputShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ConcernInputShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ConcernInputShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ConcernInputShapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "ConcernInputShapes", &diagram_.ConcernInputShapes)

		case "ConcernOutputShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ConcernOutputShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ConcernOutputShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ConcernOutputShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ConcernOutputShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ConcernOutputShapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "ConcernOutputShapes", &diagram_.ConcernOutputShapes)

		case "Note_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.NoteShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.NoteShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.NoteShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.NoteShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.Note_Shapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "Note_Shapes", &diagram_.Note_Shapes)

		case "NotesWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Note](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Note, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Note)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Note](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.NotesWhoseNodeIsExpanded = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "NotesWhoseNodeIsExpanded", &diagram_.NotesWhoseNodeIsExpanded)

		case "IsNotesNodeExpanded":
			FormDivBasicFieldToField(&(diagram_.IsNotesNodeExpanded), formDiv)
		case "NoteProductShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.NoteProductShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.NoteProductShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.NoteProductShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.NoteProductShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.NoteProductShapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "NoteProductShapes", &diagram_.NoteProductShapes)

		case "NoteTaskShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.NoteTaskShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.NoteTaskShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.NoteTaskShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.NoteTaskShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.NoteTaskShapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "NoteTaskShapes", &diagram_.NoteTaskShapes)

		case "NoteResourceShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.NoteStakeholderShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.NoteStakeholderShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.NoteStakeholderShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.NoteStakeholderShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.NoteResourceShapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "NoteResourceShapes", &diagram_.NoteResourceShapes)

		case "Stakeholder_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.StakeholderShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.StakeholderShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.StakeholderShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.StakeholderShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.Stakeholder_Shapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "Stakeholder_Shapes", &diagram_.Stakeholder_Shapes)

		case "ResourcesWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Stakeholder](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Stakeholder, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Stakeholder)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Stakeholder](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ResourcesWhoseNodeIsExpanded = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "ResourcesWhoseNodeIsExpanded", &diagram_.ResourcesWhoseNodeIsExpanded)

		case "IsStakeholdersNodeExpanded":
			FormDivBasicFieldToField(&(diagram_.IsStakeholdersNodeExpanded), formDiv)
		case "ResourceComposition_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.StakeholderCompositionShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.StakeholderCompositionShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.StakeholderCompositionShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.StakeholderCompositionShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ResourceComposition_Shapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "ResourceComposition_Shapes", &diagram_.ResourceComposition_Shapes)

		case "StakeholderConcernShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.StakeholderConcernShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.StakeholderConcernShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.StakeholderConcernShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					diagramFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.StakeholderConcernShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.StakeholderConcernShapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "StakeholderConcernShapes", &diagram_.StakeholderConcernShapes)

		case "Library:Diagrams":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](diagramFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their Diagrams slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](diagramFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(diagramFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure diagram_ is in _library.Diagrams
					found := false
					for _, _b := range _library.Diagrams {
						if _b == diagram_ {
							found = true
							break
						}
					}
					if !found {
						_library.Diagrams = append(_library.Diagrams, diagram_)
						diagramFormCallback.probe.UpdateSliceOfPointersCallback(_library, "Diagrams", &_library.Diagrams)
					}
				} else {
					// ensure diagram_ is NOT in _library.Diagrams
					idx := slices.Index(_library.Diagrams, diagram_)
					if idx != -1 {
						_library.Diagrams = slices.Delete(_library.Diagrams, idx, idx+1)
						diagramFormCallback.probe.UpdateSliceOfPointersCallback(_library, "Diagrams", &_library.Diagrams)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if diagramFormCallback.formGroup.HasSuppressButtonBeenPressed {
		diagram_.Unstage(diagramFormCallback.probe.stageOfInterest)
	}

	diagramFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Diagram](
		diagramFormCallback.probe,
	)

	// display a new form by reset the form stage
	if diagramFormCallback.CreationMode || diagramFormCallback.formGroup.HasSuppressButtonBeenPressed {
		diagramFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(diagramFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DiagramFormCallback(
			nil,
			diagramFormCallback.probe,
			newFormGroup,
		)
		diagram := new(models.Diagram)
		FillUpForm(diagram, newFormGroup, diagramFormCallback.probe)
		diagramFormCallback.probe.formStage.Commit()
	}

	diagramFormCallback.probe.ux_tree()
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
		case "IsRootLibrary":
			FormDivBasicFieldToField(&(library_.IsRootLibrary), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(library_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(library_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(library_.LayoutDirection), formDiv)
		case "RootDeliverables":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Deliverable](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Deliverable, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Deliverable)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Deliverable](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.RootDeliverables = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "RootDeliverables", &library_.RootDeliverables)

		case "RootConcerns":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Concern](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Concern, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Concern)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Concern](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.RootConcerns = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "RootConcerns", &library_.RootConcerns)

		case "RootStakeholders":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Stakeholder](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Stakeholder, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Stakeholder)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Stakeholder](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.RootStakeholders = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "RootStakeholders", &library_.RootStakeholders)

		case "RootRequirements":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Requirement](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Requirement, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Requirement)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Requirement](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.RootRequirements = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "RootRequirements", &library_.RootRequirements)

		case "RootConcepts":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Concept](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Concept, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Concept)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Concept](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.RootConcepts = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "RootConcepts", &library_.RootConcepts)

		case "AnalysisNeeds":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.AnalysisNeed](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.AnalysisNeed, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.AnalysisNeed)

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
			map_RowID_ID := GetMap_RowID_ID[*models.AnalysisNeed](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.AnalysisNeeds = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "AnalysisNeeds", &library_.AnalysisNeeds)

		case "Notes":
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
			library_.Notes = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "Notes", &library_.Notes)

		case "Diagrams":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](libraryFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Diagram, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Diagram)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](libraryFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			library_.Diagrams = instanceSlice
			libraryFormCallback.probe.UpdateSliceOfPointersCallback(library_, "Diagrams", &library_.Diagrams)

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
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(note_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(note_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(note_.LayoutDirection), formDiv)
		case "Products":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Deliverable](noteFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Deliverable, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Deliverable)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Deliverable](noteFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			note_.Products = instanceSlice
			noteFormCallback.probe.UpdateSliceOfPointersCallback(note_, "Products", &note_.Products)

		case "Tasks":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Concern](noteFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Concern, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Concern)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Concern](noteFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			note_.Tasks = instanceSlice
			noteFormCallback.probe.UpdateSliceOfPointersCallback(note_, "Tasks", &note_.Tasks)

		case "Resources":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Stakeholder](noteFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Stakeholder, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Stakeholder)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Stakeholder](noteFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			note_.Resources = instanceSlice
			noteFormCallback.probe.UpdateSliceOfPointersCallback(note_, "Resources", &note_.Resources)

		case "Diagram:NotesWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](noteFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their NotesWhoseNodeIsExpanded slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](noteFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(noteFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure note_ is in _diagram.NotesWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagram.NotesWhoseNodeIsExpanded {
						if _b == note_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.NotesWhoseNodeIsExpanded = append(_diagram.NotesWhoseNodeIsExpanded, note_)
						noteFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "NotesWhoseNodeIsExpanded", &_diagram.NotesWhoseNodeIsExpanded)
					}
				} else {
					// ensure note_ is NOT in _diagram.NotesWhoseNodeIsExpanded
					idx := slices.Index(_diagram.NotesWhoseNodeIsExpanded, note_)
					if idx != -1 {
						_diagram.NotesWhoseNodeIsExpanded = slices.Delete(_diagram.NotesWhoseNodeIsExpanded, idx, idx+1)
						noteFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "NotesWhoseNodeIsExpanded", &_diagram.NotesWhoseNodeIsExpanded)
					}
				}
			}
		case "Library:Notes":
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

			// 3. Iterate over all Library instances and update their Notes slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](noteFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(noteFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure note_ is in _library.Notes
					found := false
					for _, _b := range _library.Notes {
						if _b == note_ {
							found = true
							break
						}
					}
					if !found {
						_library.Notes = append(_library.Notes, note_)
						noteFormCallback.probe.UpdateSliceOfPointersCallback(_library, "Notes", &_library.Notes)
					}
				} else {
					// ensure note_ is NOT in _library.Notes
					idx := slices.Index(_library.Notes, note_)
					if idx != -1 {
						_library.Notes = slices.Delete(_library.Notes, idx, idx+1)
						noteFormCallback.probe.UpdateSliceOfPointersCallback(_library, "Notes", &_library.Notes)
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
func __gong__New__NoteProductShapeFormCallback(
	noteproductshape *models.NoteProductShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (noteproductshapeFormCallback *NoteProductShapeFormCallback) {
	noteproductshapeFormCallback = new(NoteProductShapeFormCallback)
	noteproductshapeFormCallback.probe = probe
	noteproductshapeFormCallback.noteproductshape = noteproductshape
	noteproductshapeFormCallback.formGroup = formGroup

	noteproductshapeFormCallback.CreationMode = (noteproductshape == nil)

	return
}

type NoteProductShapeFormCallback struct {
	noteproductshape *models.NoteProductShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (noteproductshapeFormCallback *NoteProductShapeFormCallback) OnSave() {
	noteproductshapeFormCallback.probe.stageOfInterest.Lock()
	defer noteproductshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("NoteProductShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	noteproductshapeFormCallback.probe.formStage.Checkout()

	if noteproductshapeFormCallback.noteproductshape == nil {
		noteproductshapeFormCallback.noteproductshape = new(models.NoteProductShape).Stage(noteproductshapeFormCallback.probe.stageOfInterest)
	}
	noteproductshape_ := noteproductshapeFormCallback.noteproductshape
	_ = noteproductshape_

	for _, formDiv := range noteproductshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(noteproductshape_.Name), formDiv)
		case "Note":
			FormDivSelectFieldToField(&(noteproductshape_.Note), noteproductshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Product":
			FormDivSelectFieldToField(&(noteproductshape_.Product), noteproductshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(noteproductshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(noteproductshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(noteproductshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(noteproductshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(noteproductshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(noteproductshape_.IsHidden), formDiv)
		case "Diagram:NoteProductShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](noteproductshapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their NoteProductShapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](noteproductshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(noteproductshapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure noteproductshape_ is in _diagram.NoteProductShapes
					found := false
					for _, _b := range _diagram.NoteProductShapes {
						if _b == noteproductshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.NoteProductShapes = append(_diagram.NoteProductShapes, noteproductshape_)
						noteproductshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "NoteProductShapes", &_diagram.NoteProductShapes)
					}
				} else {
					// ensure noteproductshape_ is NOT in _diagram.NoteProductShapes
					idx := slices.Index(_diagram.NoteProductShapes, noteproductshape_)
					if idx != -1 {
						_diagram.NoteProductShapes = slices.Delete(_diagram.NoteProductShapes, idx, idx+1)
						noteproductshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "NoteProductShapes", &_diagram.NoteProductShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if noteproductshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteproductshape_.Unstage(noteproductshapeFormCallback.probe.stageOfInterest)
	}

	noteproductshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.NoteProductShape](
		noteproductshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if noteproductshapeFormCallback.CreationMode || noteproductshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		noteproductshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(noteproductshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__NoteProductShapeFormCallback(
			nil,
			noteproductshapeFormCallback.probe,
			newFormGroup,
		)
		noteproductshape := new(models.NoteProductShape)
		FillUpForm(noteproductshape, newFormGroup, noteproductshapeFormCallback.probe)
		noteproductshapeFormCallback.probe.formStage.Commit()
	}

	noteproductshapeFormCallback.probe.ux_tree()
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
		case "IsExpanded":
			FormDivBasicFieldToField(&(noteshape_.IsExpanded), formDiv)
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
		case "Diagram:Note_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](noteshapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their Note_Shapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](noteshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(noteshapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure noteshape_ is in _diagram.Note_Shapes
					found := false
					for _, _b := range _diagram.Note_Shapes {
						if _b == noteshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.Note_Shapes = append(_diagram.Note_Shapes, noteshape_)
						noteshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "Note_Shapes", &_diagram.Note_Shapes)
					}
				} else {
					// ensure noteshape_ is NOT in _diagram.Note_Shapes
					idx := slices.Index(_diagram.Note_Shapes, noteshape_)
					if idx != -1 {
						_diagram.Note_Shapes = slices.Delete(_diagram.Note_Shapes, idx, idx+1)
						noteshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "Note_Shapes", &_diagram.Note_Shapes)
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
func __gong__New__NoteStakeholderShapeFormCallback(
	notestakeholdershape *models.NoteStakeholderShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (notestakeholdershapeFormCallback *NoteStakeholderShapeFormCallback) {
	notestakeholdershapeFormCallback = new(NoteStakeholderShapeFormCallback)
	notestakeholdershapeFormCallback.probe = probe
	notestakeholdershapeFormCallback.notestakeholdershape = notestakeholdershape
	notestakeholdershapeFormCallback.formGroup = formGroup

	notestakeholdershapeFormCallback.CreationMode = (notestakeholdershape == nil)

	return
}

type NoteStakeholderShapeFormCallback struct {
	notestakeholdershape *models.NoteStakeholderShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (notestakeholdershapeFormCallback *NoteStakeholderShapeFormCallback) OnSave() {
	notestakeholdershapeFormCallback.probe.stageOfInterest.Lock()
	defer notestakeholdershapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("NoteStakeholderShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	notestakeholdershapeFormCallback.probe.formStage.Checkout()

	if notestakeholdershapeFormCallback.notestakeholdershape == nil {
		notestakeholdershapeFormCallback.notestakeholdershape = new(models.NoteStakeholderShape).Stage(notestakeholdershapeFormCallback.probe.stageOfInterest)
	}
	notestakeholdershape_ := notestakeholdershapeFormCallback.notestakeholdershape
	_ = notestakeholdershape_

	for _, formDiv := range notestakeholdershapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(notestakeholdershape_.Name), formDiv)
		case "Note":
			FormDivSelectFieldToField(&(notestakeholdershape_.Note), notestakeholdershapeFormCallback.probe.stageOfInterest, formDiv)
		case "Stakeholder":
			FormDivSelectFieldToField(&(notestakeholdershape_.Stakeholder), notestakeholdershapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(notestakeholdershape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(notestakeholdershape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(notestakeholdershape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(notestakeholdershape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(notestakeholdershape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(notestakeholdershape_.IsHidden), formDiv)
		case "Diagram:NoteResourceShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](notestakeholdershapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their NoteResourceShapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](notestakeholdershapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(notestakeholdershapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure notestakeholdershape_ is in _diagram.NoteResourceShapes
					found := false
					for _, _b := range _diagram.NoteResourceShapes {
						if _b == notestakeholdershape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.NoteResourceShapes = append(_diagram.NoteResourceShapes, notestakeholdershape_)
						notestakeholdershapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "NoteResourceShapes", &_diagram.NoteResourceShapes)
					}
				} else {
					// ensure notestakeholdershape_ is NOT in _diagram.NoteResourceShapes
					idx := slices.Index(_diagram.NoteResourceShapes, notestakeholdershape_)
					if idx != -1 {
						_diagram.NoteResourceShapes = slices.Delete(_diagram.NoteResourceShapes, idx, idx+1)
						notestakeholdershapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "NoteResourceShapes", &_diagram.NoteResourceShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if notestakeholdershapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		notestakeholdershape_.Unstage(notestakeholdershapeFormCallback.probe.stageOfInterest)
	}

	notestakeholdershapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.NoteStakeholderShape](
		notestakeholdershapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if notestakeholdershapeFormCallback.CreationMode || notestakeholdershapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		notestakeholdershapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(notestakeholdershapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__NoteStakeholderShapeFormCallback(
			nil,
			notestakeholdershapeFormCallback.probe,
			newFormGroup,
		)
		notestakeholdershape := new(models.NoteStakeholderShape)
		FillUpForm(notestakeholdershape, newFormGroup, notestakeholdershapeFormCallback.probe)
		notestakeholdershapeFormCallback.probe.formStage.Commit()
	}

	notestakeholdershapeFormCallback.probe.ux_tree()
}
func __gong__New__NoteTaskShapeFormCallback(
	notetaskshape *models.NoteTaskShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (notetaskshapeFormCallback *NoteTaskShapeFormCallback) {
	notetaskshapeFormCallback = new(NoteTaskShapeFormCallback)
	notetaskshapeFormCallback.probe = probe
	notetaskshapeFormCallback.notetaskshape = notetaskshape
	notetaskshapeFormCallback.formGroup = formGroup

	notetaskshapeFormCallback.CreationMode = (notetaskshape == nil)

	return
}

type NoteTaskShapeFormCallback struct {
	notetaskshape *models.NoteTaskShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (notetaskshapeFormCallback *NoteTaskShapeFormCallback) OnSave() {
	notetaskshapeFormCallback.probe.stageOfInterest.Lock()
	defer notetaskshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("NoteTaskShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	notetaskshapeFormCallback.probe.formStage.Checkout()

	if notetaskshapeFormCallback.notetaskshape == nil {
		notetaskshapeFormCallback.notetaskshape = new(models.NoteTaskShape).Stage(notetaskshapeFormCallback.probe.stageOfInterest)
	}
	notetaskshape_ := notetaskshapeFormCallback.notetaskshape
	_ = notetaskshape_

	for _, formDiv := range notetaskshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(notetaskshape_.Name), formDiv)
		case "Note":
			FormDivSelectFieldToField(&(notetaskshape_.Note), notetaskshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Task":
			FormDivSelectFieldToField(&(notetaskshape_.Task), notetaskshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(notetaskshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(notetaskshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(notetaskshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(notetaskshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(notetaskshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(notetaskshape_.IsHidden), formDiv)
		case "Diagram:NoteTaskShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](notetaskshapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their NoteTaskShapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](notetaskshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(notetaskshapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure notetaskshape_ is in _diagram.NoteTaskShapes
					found := false
					for _, _b := range _diagram.NoteTaskShapes {
						if _b == notetaskshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.NoteTaskShapes = append(_diagram.NoteTaskShapes, notetaskshape_)
						notetaskshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "NoteTaskShapes", &_diagram.NoteTaskShapes)
					}
				} else {
					// ensure notetaskshape_ is NOT in _diagram.NoteTaskShapes
					idx := slices.Index(_diagram.NoteTaskShapes, notetaskshape_)
					if idx != -1 {
						_diagram.NoteTaskShapes = slices.Delete(_diagram.NoteTaskShapes, idx, idx+1)
						notetaskshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "NoteTaskShapes", &_diagram.NoteTaskShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if notetaskshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		notetaskshape_.Unstage(notetaskshapeFormCallback.probe.stageOfInterest)
	}

	notetaskshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.NoteTaskShape](
		notetaskshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if notetaskshapeFormCallback.CreationMode || notetaskshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		notetaskshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(notetaskshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__NoteTaskShapeFormCallback(
			nil,
			notetaskshapeFormCallback.probe,
			newFormGroup,
		)
		notetaskshape := new(models.NoteTaskShape)
		FillUpForm(notetaskshape, newFormGroup, notetaskshapeFormCallback.probe)
		notetaskshapeFormCallback.probe.formStage.Commit()
	}

	notetaskshapeFormCallback.probe.ux_tree()
}
func __gong__New__ProductCompositionShapeFormCallback(
	productcompositionshape *models.ProductCompositionShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (productcompositionshapeFormCallback *ProductCompositionShapeFormCallback) {
	productcompositionshapeFormCallback = new(ProductCompositionShapeFormCallback)
	productcompositionshapeFormCallback.probe = probe
	productcompositionshapeFormCallback.productcompositionshape = productcompositionshape
	productcompositionshapeFormCallback.formGroup = formGroup

	productcompositionshapeFormCallback.CreationMode = (productcompositionshape == nil)

	return
}

type ProductCompositionShapeFormCallback struct {
	productcompositionshape *models.ProductCompositionShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (productcompositionshapeFormCallback *ProductCompositionShapeFormCallback) OnSave() {
	productcompositionshapeFormCallback.probe.stageOfInterest.Lock()
	defer productcompositionshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ProductCompositionShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	productcompositionshapeFormCallback.probe.formStage.Checkout()

	if productcompositionshapeFormCallback.productcompositionshape == nil {
		productcompositionshapeFormCallback.productcompositionshape = new(models.ProductCompositionShape).Stage(productcompositionshapeFormCallback.probe.stageOfInterest)
	}
	productcompositionshape_ := productcompositionshapeFormCallback.productcompositionshape
	_ = productcompositionshape_

	for _, formDiv := range productcompositionshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(productcompositionshape_.Name), formDiv)
		case "Product":
			FormDivSelectFieldToField(&(productcompositionshape_.Product), productcompositionshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(productcompositionshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(productcompositionshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(productcompositionshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(productcompositionshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(productcompositionshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(productcompositionshape_.IsHidden), formDiv)
		case "Diagram:ProductComposition_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](productcompositionshapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their ProductComposition_Shapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](productcompositionshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(productcompositionshapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure productcompositionshape_ is in _diagram.ProductComposition_Shapes
					found := false
					for _, _b := range _diagram.ProductComposition_Shapes {
						if _b == productcompositionshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.ProductComposition_Shapes = append(_diagram.ProductComposition_Shapes, productcompositionshape_)
						productcompositionshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ProductComposition_Shapes", &_diagram.ProductComposition_Shapes)
					}
				} else {
					// ensure productcompositionshape_ is NOT in _diagram.ProductComposition_Shapes
					idx := slices.Index(_diagram.ProductComposition_Shapes, productcompositionshape_)
					if idx != -1 {
						_diagram.ProductComposition_Shapes = slices.Delete(_diagram.ProductComposition_Shapes, idx, idx+1)
						productcompositionshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ProductComposition_Shapes", &_diagram.ProductComposition_Shapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if productcompositionshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		productcompositionshape_.Unstage(productcompositionshapeFormCallback.probe.stageOfInterest)
	}

	productcompositionshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ProductCompositionShape](
		productcompositionshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if productcompositionshapeFormCallback.CreationMode || productcompositionshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		productcompositionshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(productcompositionshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ProductCompositionShapeFormCallback(
			nil,
			productcompositionshapeFormCallback.probe,
			newFormGroup,
		)
		productcompositionshape := new(models.ProductCompositionShape)
		FillUpForm(productcompositionshape, newFormGroup, productcompositionshapeFormCallback.probe)
		productcompositionshapeFormCallback.probe.formStage.Commit()
	}

	productcompositionshapeFormCallback.probe.ux_tree()
}
func __gong__New__ProductShapeFormCallback(
	productshape *models.ProductShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (productshapeFormCallback *ProductShapeFormCallback) {
	productshapeFormCallback = new(ProductShapeFormCallback)
	productshapeFormCallback.probe = probe
	productshapeFormCallback.productshape = productshape
	productshapeFormCallback.formGroup = formGroup

	productshapeFormCallback.CreationMode = (productshape == nil)

	return
}

type ProductShapeFormCallback struct {
	productshape *models.ProductShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (productshapeFormCallback *ProductShapeFormCallback) OnSave() {
	productshapeFormCallback.probe.stageOfInterest.Lock()
	defer productshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ProductShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	productshapeFormCallback.probe.formStage.Checkout()

	if productshapeFormCallback.productshape == nil {
		productshapeFormCallback.productshape = new(models.ProductShape).Stage(productshapeFormCallback.probe.stageOfInterest)
	}
	productshape_ := productshapeFormCallback.productshape
	_ = productshape_

	for _, formDiv := range productshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(productshape_.Name), formDiv)
		case "Product":
			FormDivSelectFieldToField(&(productshape_.Product), productshapeFormCallback.probe.stageOfInterest, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(productshape_.IsExpanded), formDiv)
		case "X":
			FormDivBasicFieldToField(&(productshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(productshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(productshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(productshape_.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(productshape_.IsHidden), formDiv)
		case "Diagram:Product_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](productshapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their Product_Shapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](productshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(productshapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure productshape_ is in _diagram.Product_Shapes
					found := false
					for _, _b := range _diagram.Product_Shapes {
						if _b == productshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.Product_Shapes = append(_diagram.Product_Shapes, productshape_)
						productshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "Product_Shapes", &_diagram.Product_Shapes)
					}
				} else {
					// ensure productshape_ is NOT in _diagram.Product_Shapes
					idx := slices.Index(_diagram.Product_Shapes, productshape_)
					if idx != -1 {
						_diagram.Product_Shapes = slices.Delete(_diagram.Product_Shapes, idx, idx+1)
						productshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "Product_Shapes", &_diagram.Product_Shapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if productshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		productshape_.Unstage(productshapeFormCallback.probe.stageOfInterest)
	}

	productshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ProductShape](
		productshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if productshapeFormCallback.CreationMode || productshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		productshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(productshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ProductShapeFormCallback(
			nil,
			productshapeFormCallback.probe,
			newFormGroup,
		)
		productshape := new(models.ProductShape)
		FillUpForm(productshape, newFormGroup, productshapeFormCallback.probe)
		productshapeFormCallback.probe.formStage.Commit()
	}

	productshapeFormCallback.probe.ux_tree()
}
func __gong__New__RequirementFormCallback(
	requirement *models.Requirement,
	probe *Probe,
	formGroup *form.FormGroup,
) (requirementFormCallback *RequirementFormCallback) {
	requirementFormCallback = new(RequirementFormCallback)
	requirementFormCallback.probe = probe
	requirementFormCallback.requirement = requirement
	requirementFormCallback.formGroup = formGroup

	requirementFormCallback.CreationMode = (requirement == nil)

	return
}

type RequirementFormCallback struct {
	requirement *models.Requirement

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (requirementFormCallback *RequirementFormCallback) OnSave() {
	requirementFormCallback.probe.stageOfInterest.Lock()
	defer requirementFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("RequirementFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	requirementFormCallback.probe.formStage.Checkout()

	if requirementFormCallback.requirement == nil {
		requirementFormCallback.requirement = new(models.Requirement).Stage(requirementFormCallback.probe.stageOfInterest)
	}
	requirement_ := requirementFormCallback.requirement
	_ = requirement_

	for _, formDiv := range requirementFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(requirement_.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(requirement_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(requirement_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(requirement_.LayoutDirection), formDiv)
		case "SupportLevels":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.SupportLevel](requirementFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.SupportLevel, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.SupportLevel)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					requirementFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.SupportLevel](requirementFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			requirement_.SupportLevels = instanceSlice
			requirementFormCallback.probe.UpdateSliceOfPointersCallback(requirement_, "SupportLevels", &requirement_.SupportLevels)

		case "Concepts":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Concept](requirementFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Concept, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Concept)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					requirementFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Concept](requirementFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			requirement_.Concepts = instanceSlice
			requirementFormCallback.probe.UpdateSliceOfPointersCallback(requirement_, "Concepts", &requirement_.Concepts)

		case "Concern:Requirements":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Concern instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Concern instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Concern](requirementFormCallback.probe.stageOfInterest)
			targetConcernIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetConcernIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Concern instances and update their Requirements slice
			for _concern := range *models.GetGongstructInstancesSetFromPointerType[*models.Concern](requirementFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(requirementFormCallback.probe.stageOfInterest, _concern)
				
				// if Concern is selected
				if targetConcernIDs[id] {
					// ensure requirement_ is in _concern.Requirements
					found := false
					for _, _b := range _concern.Requirements {
						if _b == requirement_ {
							found = true
							break
						}
					}
					if !found {
						_concern.Requirements = append(_concern.Requirements, requirement_)
						requirementFormCallback.probe.UpdateSliceOfPointersCallback(_concern, "Requirements", &_concern.Requirements)
					}
				} else {
					// ensure requirement_ is NOT in _concern.Requirements
					idx := slices.Index(_concern.Requirements, requirement_)
					if idx != -1 {
						_concern.Requirements = slices.Delete(_concern.Requirements, idx, idx+1)
						requirementFormCallback.probe.UpdateSliceOfPointersCallback(_concern, "Requirements", &_concern.Requirements)
					}
				}
			}
		case "Library:RootRequirements":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](requirementFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their RootRequirements slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](requirementFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(requirementFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure requirement_ is in _library.RootRequirements
					found := false
					for _, _b := range _library.RootRequirements {
						if _b == requirement_ {
							found = true
							break
						}
					}
					if !found {
						_library.RootRequirements = append(_library.RootRequirements, requirement_)
						requirementFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootRequirements", &_library.RootRequirements)
					}
				} else {
					// ensure requirement_ is NOT in _library.RootRequirements
					idx := slices.Index(_library.RootRequirements, requirement_)
					if idx != -1 {
						_library.RootRequirements = slices.Delete(_library.RootRequirements, idx, idx+1)
						requirementFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootRequirements", &_library.RootRequirements)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if requirementFormCallback.formGroup.HasSuppressButtonBeenPressed {
		requirement_.Unstage(requirementFormCallback.probe.stageOfInterest)
	}

	requirementFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Requirement](
		requirementFormCallback.probe,
	)

	// display a new form by reset the form stage
	if requirementFormCallback.CreationMode || requirementFormCallback.formGroup.HasSuppressButtonBeenPressed {
		requirementFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(requirementFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RequirementFormCallback(
			nil,
			requirementFormCallback.probe,
			newFormGroup,
		)
		requirement := new(models.Requirement)
		FillUpForm(requirement, newFormGroup, requirementFormCallback.probe)
		requirementFormCallback.probe.formStage.Commit()
	}

	requirementFormCallback.probe.ux_tree()
}
func __gong__New__StakeholderFormCallback(
	stakeholder *models.Stakeholder,
	probe *Probe,
	formGroup *form.FormGroup,
) (stakeholderFormCallback *StakeholderFormCallback) {
	stakeholderFormCallback = new(StakeholderFormCallback)
	stakeholderFormCallback.probe = probe
	stakeholderFormCallback.stakeholder = stakeholder
	stakeholderFormCallback.formGroup = formGroup

	stakeholderFormCallback.CreationMode = (stakeholder == nil)

	return
}

type StakeholderFormCallback struct {
	stakeholder *models.Stakeholder

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (stakeholderFormCallback *StakeholderFormCallback) OnSave() {
	stakeholderFormCallback.probe.stageOfInterest.Lock()
	defer stakeholderFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StakeholderFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	stakeholderFormCallback.probe.formStage.Checkout()

	if stakeholderFormCallback.stakeholder == nil {
		stakeholderFormCallback.stakeholder = new(models.Stakeholder).Stage(stakeholderFormCallback.probe.stageOfInterest)
	}
	stakeholder_ := stakeholderFormCallback.stakeholder
	_ = stakeholder_

	for _, formDiv := range stakeholderFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(stakeholder_.Name), formDiv)
		case "IDAirbus":
			FormDivBasicFieldToField(&(stakeholder_.IDAirbus), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(stakeholder_.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(stakeholder_.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(stakeholder_.LayoutDirection), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(stakeholder_.Description), formDiv)
		case "Concerns":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Concern](stakeholderFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Concern, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Concern)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					stakeholderFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Concern](stakeholderFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			stakeholder_.Concerns = instanceSlice
			stakeholderFormCallback.probe.UpdateSliceOfPointersCallback(stakeholder_, "Concerns", &stakeholder_.Concerns)

		case "SubStakeholders":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Stakeholder](stakeholderFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Stakeholder, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Stakeholder)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					stakeholderFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Stakeholder](stakeholderFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			stakeholder_.SubStakeholders = instanceSlice
			stakeholderFormCallback.probe.UpdateSliceOfPointersCallback(stakeholder_, "SubStakeholders", &stakeholder_.SubStakeholders)

		case "Diagram:ResourcesWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](stakeholderFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their ResourcesWhoseNodeIsExpanded slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](stakeholderFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(stakeholderFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure stakeholder_ is in _diagram.ResourcesWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagram.ResourcesWhoseNodeIsExpanded {
						if _b == stakeholder_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.ResourcesWhoseNodeIsExpanded = append(_diagram.ResourcesWhoseNodeIsExpanded, stakeholder_)
						stakeholderFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ResourcesWhoseNodeIsExpanded", &_diagram.ResourcesWhoseNodeIsExpanded)
					}
				} else {
					// ensure stakeholder_ is NOT in _diagram.ResourcesWhoseNodeIsExpanded
					idx := slices.Index(_diagram.ResourcesWhoseNodeIsExpanded, stakeholder_)
					if idx != -1 {
						_diagram.ResourcesWhoseNodeIsExpanded = slices.Delete(_diagram.ResourcesWhoseNodeIsExpanded, idx, idx+1)
						stakeholderFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ResourcesWhoseNodeIsExpanded", &_diagram.ResourcesWhoseNodeIsExpanded)
					}
				}
			}
		case "Library:RootStakeholders":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Library instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Library instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Library](stakeholderFormCallback.probe.stageOfInterest)
			targetLibraryIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetLibraryIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Library instances and update their RootStakeholders slice
			for _library := range *models.GetGongstructInstancesSetFromPointerType[*models.Library](stakeholderFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(stakeholderFormCallback.probe.stageOfInterest, _library)
				
				// if Library is selected
				if targetLibraryIDs[id] {
					// ensure stakeholder_ is in _library.RootStakeholders
					found := false
					for _, _b := range _library.RootStakeholders {
						if _b == stakeholder_ {
							found = true
							break
						}
					}
					if !found {
						_library.RootStakeholders = append(_library.RootStakeholders, stakeholder_)
						stakeholderFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootStakeholders", &_library.RootStakeholders)
					}
				} else {
					// ensure stakeholder_ is NOT in _library.RootStakeholders
					idx := slices.Index(_library.RootStakeholders, stakeholder_)
					if idx != -1 {
						_library.RootStakeholders = slices.Delete(_library.RootStakeholders, idx, idx+1)
						stakeholderFormCallback.probe.UpdateSliceOfPointersCallback(_library, "RootStakeholders", &_library.RootStakeholders)
					}
				}
			}
		case "Note:Resources":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Note instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Note instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Note](stakeholderFormCallback.probe.stageOfInterest)
			targetNoteIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetNoteIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Note instances and update their Resources slice
			for _note := range *models.GetGongstructInstancesSetFromPointerType[*models.Note](stakeholderFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(stakeholderFormCallback.probe.stageOfInterest, _note)
				
				// if Note is selected
				if targetNoteIDs[id] {
					// ensure stakeholder_ is in _note.Resources
					found := false
					for _, _b := range _note.Resources {
						if _b == stakeholder_ {
							found = true
							break
						}
					}
					if !found {
						_note.Resources = append(_note.Resources, stakeholder_)
						stakeholderFormCallback.probe.UpdateSliceOfPointersCallback(_note, "Resources", &_note.Resources)
					}
				} else {
					// ensure stakeholder_ is NOT in _note.Resources
					idx := slices.Index(_note.Resources, stakeholder_)
					if idx != -1 {
						_note.Resources = slices.Delete(_note.Resources, idx, idx+1)
						stakeholderFormCallback.probe.UpdateSliceOfPointersCallback(_note, "Resources", &_note.Resources)
					}
				}
			}
		case "Stakeholder:SubStakeholders":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Stakeholder instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Stakeholder instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Stakeholder](stakeholderFormCallback.probe.stageOfInterest)
			targetStakeholderIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetStakeholderIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Stakeholder instances and update their SubStakeholders slice
			for _stakeholder := range *models.GetGongstructInstancesSetFromPointerType[*models.Stakeholder](stakeholderFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(stakeholderFormCallback.probe.stageOfInterest, _stakeholder)
				
				// if Stakeholder is selected
				if targetStakeholderIDs[id] {
					// ensure stakeholder_ is in _stakeholder.SubStakeholders
					found := false
					for _, _b := range _stakeholder.SubStakeholders {
						if _b == stakeholder_ {
							found = true
							break
						}
					}
					if !found {
						_stakeholder.SubStakeholders = append(_stakeholder.SubStakeholders, stakeholder_)
						stakeholderFormCallback.probe.UpdateSliceOfPointersCallback(_stakeholder, "SubStakeholders", &_stakeholder.SubStakeholders)
					}
				} else {
					// ensure stakeholder_ is NOT in _stakeholder.SubStakeholders
					idx := slices.Index(_stakeholder.SubStakeholders, stakeholder_)
					if idx != -1 {
						_stakeholder.SubStakeholders = slices.Delete(_stakeholder.SubStakeholders, idx, idx+1)
						stakeholderFormCallback.probe.UpdateSliceOfPointersCallback(_stakeholder, "SubStakeholders", &_stakeholder.SubStakeholders)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if stakeholderFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stakeholder_.Unstage(stakeholderFormCallback.probe.stageOfInterest)
	}

	stakeholderFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Stakeholder](
		stakeholderFormCallback.probe,
	)

	// display a new form by reset the form stage
	if stakeholderFormCallback.CreationMode || stakeholderFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stakeholderFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(stakeholderFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StakeholderFormCallback(
			nil,
			stakeholderFormCallback.probe,
			newFormGroup,
		)
		stakeholder := new(models.Stakeholder)
		FillUpForm(stakeholder, newFormGroup, stakeholderFormCallback.probe)
		stakeholderFormCallback.probe.formStage.Commit()
	}

	stakeholderFormCallback.probe.ux_tree()
}
func __gong__New__StakeholderCompositionShapeFormCallback(
	stakeholdercompositionshape *models.StakeholderCompositionShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (stakeholdercompositionshapeFormCallback *StakeholderCompositionShapeFormCallback) {
	stakeholdercompositionshapeFormCallback = new(StakeholderCompositionShapeFormCallback)
	stakeholdercompositionshapeFormCallback.probe = probe
	stakeholdercompositionshapeFormCallback.stakeholdercompositionshape = stakeholdercompositionshape
	stakeholdercompositionshapeFormCallback.formGroup = formGroup

	stakeholdercompositionshapeFormCallback.CreationMode = (stakeholdercompositionshape == nil)

	return
}

type StakeholderCompositionShapeFormCallback struct {
	stakeholdercompositionshape *models.StakeholderCompositionShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (stakeholdercompositionshapeFormCallback *StakeholderCompositionShapeFormCallback) OnSave() {
	stakeholdercompositionshapeFormCallback.probe.stageOfInterest.Lock()
	defer stakeholdercompositionshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StakeholderCompositionShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	stakeholdercompositionshapeFormCallback.probe.formStage.Checkout()

	if stakeholdercompositionshapeFormCallback.stakeholdercompositionshape == nil {
		stakeholdercompositionshapeFormCallback.stakeholdercompositionshape = new(models.StakeholderCompositionShape).Stage(stakeholdercompositionshapeFormCallback.probe.stageOfInterest)
	}
	stakeholdercompositionshape_ := stakeholdercompositionshapeFormCallback.stakeholdercompositionshape
	_ = stakeholdercompositionshape_

	for _, formDiv := range stakeholdercompositionshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(stakeholdercompositionshape_.Name), formDiv)
		case "Stakeholder":
			FormDivSelectFieldToField(&(stakeholdercompositionshape_.Stakeholder), stakeholdercompositionshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(stakeholdercompositionshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(stakeholdercompositionshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(stakeholdercompositionshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(stakeholdercompositionshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(stakeholdercompositionshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(stakeholdercompositionshape_.IsHidden), formDiv)
		case "Diagram:ResourceComposition_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](stakeholdercompositionshapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their ResourceComposition_Shapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](stakeholdercompositionshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(stakeholdercompositionshapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure stakeholdercompositionshape_ is in _diagram.ResourceComposition_Shapes
					found := false
					for _, _b := range _diagram.ResourceComposition_Shapes {
						if _b == stakeholdercompositionshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.ResourceComposition_Shapes = append(_diagram.ResourceComposition_Shapes, stakeholdercompositionshape_)
						stakeholdercompositionshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ResourceComposition_Shapes", &_diagram.ResourceComposition_Shapes)
					}
				} else {
					// ensure stakeholdercompositionshape_ is NOT in _diagram.ResourceComposition_Shapes
					idx := slices.Index(_diagram.ResourceComposition_Shapes, stakeholdercompositionshape_)
					if idx != -1 {
						_diagram.ResourceComposition_Shapes = slices.Delete(_diagram.ResourceComposition_Shapes, idx, idx+1)
						stakeholdercompositionshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ResourceComposition_Shapes", &_diagram.ResourceComposition_Shapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if stakeholdercompositionshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stakeholdercompositionshape_.Unstage(stakeholdercompositionshapeFormCallback.probe.stageOfInterest)
	}

	stakeholdercompositionshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StakeholderCompositionShape](
		stakeholdercompositionshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if stakeholdercompositionshapeFormCallback.CreationMode || stakeholdercompositionshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stakeholdercompositionshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(stakeholdercompositionshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StakeholderCompositionShapeFormCallback(
			nil,
			stakeholdercompositionshapeFormCallback.probe,
			newFormGroup,
		)
		stakeholdercompositionshape := new(models.StakeholderCompositionShape)
		FillUpForm(stakeholdercompositionshape, newFormGroup, stakeholdercompositionshapeFormCallback.probe)
		stakeholdercompositionshapeFormCallback.probe.formStage.Commit()
	}

	stakeholdercompositionshapeFormCallback.probe.ux_tree()
}
func __gong__New__StakeholderConcernShapeFormCallback(
	stakeholderconcernshape *models.StakeholderConcernShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (stakeholderconcernshapeFormCallback *StakeholderConcernShapeFormCallback) {
	stakeholderconcernshapeFormCallback = new(StakeholderConcernShapeFormCallback)
	stakeholderconcernshapeFormCallback.probe = probe
	stakeholderconcernshapeFormCallback.stakeholderconcernshape = stakeholderconcernshape
	stakeholderconcernshapeFormCallback.formGroup = formGroup

	stakeholderconcernshapeFormCallback.CreationMode = (stakeholderconcernshape == nil)

	return
}

type StakeholderConcernShapeFormCallback struct {
	stakeholderconcernshape *models.StakeholderConcernShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (stakeholderconcernshapeFormCallback *StakeholderConcernShapeFormCallback) OnSave() {
	stakeholderconcernshapeFormCallback.probe.stageOfInterest.Lock()
	defer stakeholderconcernshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StakeholderConcernShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	stakeholderconcernshapeFormCallback.probe.formStage.Checkout()

	if stakeholderconcernshapeFormCallback.stakeholderconcernshape == nil {
		stakeholderconcernshapeFormCallback.stakeholderconcernshape = new(models.StakeholderConcernShape).Stage(stakeholderconcernshapeFormCallback.probe.stageOfInterest)
	}
	stakeholderconcernshape_ := stakeholderconcernshapeFormCallback.stakeholderconcernshape
	_ = stakeholderconcernshape_

	for _, formDiv := range stakeholderconcernshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(stakeholderconcernshape_.Name), formDiv)
		case "Stakeholder":
			FormDivSelectFieldToField(&(stakeholderconcernshape_.Stakeholder), stakeholderconcernshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Concern":
			FormDivSelectFieldToField(&(stakeholderconcernshape_.Concern), stakeholderconcernshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(stakeholderconcernshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(stakeholderconcernshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(stakeholderconcernshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(stakeholderconcernshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(stakeholderconcernshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(stakeholderconcernshape_.IsHidden), formDiv)
		case "Diagram:StakeholderConcernShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](stakeholderconcernshapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their StakeholderConcernShapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](stakeholderconcernshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(stakeholderconcernshapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure stakeholderconcernshape_ is in _diagram.StakeholderConcernShapes
					found := false
					for _, _b := range _diagram.StakeholderConcernShapes {
						if _b == stakeholderconcernshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.StakeholderConcernShapes = append(_diagram.StakeholderConcernShapes, stakeholderconcernshape_)
						stakeholderconcernshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "StakeholderConcernShapes", &_diagram.StakeholderConcernShapes)
					}
				} else {
					// ensure stakeholderconcernshape_ is NOT in _diagram.StakeholderConcernShapes
					idx := slices.Index(_diagram.StakeholderConcernShapes, stakeholderconcernshape_)
					if idx != -1 {
						_diagram.StakeholderConcernShapes = slices.Delete(_diagram.StakeholderConcernShapes, idx, idx+1)
						stakeholderconcernshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "StakeholderConcernShapes", &_diagram.StakeholderConcernShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if stakeholderconcernshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stakeholderconcernshape_.Unstage(stakeholderconcernshapeFormCallback.probe.stageOfInterest)
	}

	stakeholderconcernshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StakeholderConcernShape](
		stakeholderconcernshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if stakeholderconcernshapeFormCallback.CreationMode || stakeholderconcernshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stakeholderconcernshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(stakeholderconcernshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StakeholderConcernShapeFormCallback(
			nil,
			stakeholderconcernshapeFormCallback.probe,
			newFormGroup,
		)
		stakeholderconcernshape := new(models.StakeholderConcernShape)
		FillUpForm(stakeholderconcernshape, newFormGroup, stakeholderconcernshapeFormCallback.probe)
		stakeholderconcernshapeFormCallback.probe.formStage.Commit()
	}

	stakeholderconcernshapeFormCallback.probe.ux_tree()
}
func __gong__New__StakeholderShapeFormCallback(
	stakeholdershape *models.StakeholderShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (stakeholdershapeFormCallback *StakeholderShapeFormCallback) {
	stakeholdershapeFormCallback = new(StakeholderShapeFormCallback)
	stakeholdershapeFormCallback.probe = probe
	stakeholdershapeFormCallback.stakeholdershape = stakeholdershape
	stakeholdershapeFormCallback.formGroup = formGroup

	stakeholdershapeFormCallback.CreationMode = (stakeholdershape == nil)

	return
}

type StakeholderShapeFormCallback struct {
	stakeholdershape *models.StakeholderShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (stakeholdershapeFormCallback *StakeholderShapeFormCallback) OnSave() {
	stakeholdershapeFormCallback.probe.stageOfInterest.Lock()
	defer stakeholdershapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("StakeholderShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	stakeholdershapeFormCallback.probe.formStage.Checkout()

	if stakeholdershapeFormCallback.stakeholdershape == nil {
		stakeholdershapeFormCallback.stakeholdershape = new(models.StakeholderShape).Stage(stakeholdershapeFormCallback.probe.stageOfInterest)
	}
	stakeholdershape_ := stakeholdershapeFormCallback.stakeholdershape
	_ = stakeholdershape_

	for _, formDiv := range stakeholdershapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(stakeholdershape_.Name), formDiv)
		case "Stakeholder":
			FormDivSelectFieldToField(&(stakeholdershape_.Stakeholder), stakeholdershapeFormCallback.probe.stageOfInterest, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(stakeholdershape_.IsExpanded), formDiv)
		case "X":
			FormDivBasicFieldToField(&(stakeholdershape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(stakeholdershape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(stakeholdershape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(stakeholdershape_.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(stakeholdershape_.IsHidden), formDiv)
		case "Diagram:Stakeholder_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](stakeholdershapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their Stakeholder_Shapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](stakeholdershapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(stakeholdershapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure stakeholdershape_ is in _diagram.Stakeholder_Shapes
					found := false
					for _, _b := range _diagram.Stakeholder_Shapes {
						if _b == stakeholdershape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.Stakeholder_Shapes = append(_diagram.Stakeholder_Shapes, stakeholdershape_)
						stakeholdershapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "Stakeholder_Shapes", &_diagram.Stakeholder_Shapes)
					}
				} else {
					// ensure stakeholdershape_ is NOT in _diagram.Stakeholder_Shapes
					idx := slices.Index(_diagram.Stakeholder_Shapes, stakeholdershape_)
					if idx != -1 {
						_diagram.Stakeholder_Shapes = slices.Delete(_diagram.Stakeholder_Shapes, idx, idx+1)
						stakeholdershapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "Stakeholder_Shapes", &_diagram.Stakeholder_Shapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if stakeholdershapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stakeholdershape_.Unstage(stakeholdershapeFormCallback.probe.stageOfInterest)
	}

	stakeholdershapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.StakeholderShape](
		stakeholdershapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if stakeholdershapeFormCallback.CreationMode || stakeholdershapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		stakeholdershapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(stakeholdershapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__StakeholderShapeFormCallback(
			nil,
			stakeholdershapeFormCallback.probe,
			newFormGroup,
		)
		stakeholdershape := new(models.StakeholderShape)
		FillUpForm(stakeholdershape, newFormGroup, stakeholdershapeFormCallback.probe)
		stakeholdershapeFormCallback.probe.formStage.Commit()
	}

	stakeholdershapeFormCallback.probe.ux_tree()
}
func __gong__New__SupportLevelFormCallback(
	supportlevel *models.SupportLevel,
	probe *Probe,
	formGroup *form.FormGroup,
) (supportlevelFormCallback *SupportLevelFormCallback) {
	supportlevelFormCallback = new(SupportLevelFormCallback)
	supportlevelFormCallback.probe = probe
	supportlevelFormCallback.supportlevel = supportlevel
	supportlevelFormCallback.formGroup = formGroup

	supportlevelFormCallback.CreationMode = (supportlevel == nil)

	return
}

type SupportLevelFormCallback struct {
	supportlevel *models.SupportLevel

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (supportlevelFormCallback *SupportLevelFormCallback) OnSave() {
	supportlevelFormCallback.probe.stageOfInterest.Lock()
	defer supportlevelFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("SupportLevelFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	supportlevelFormCallback.probe.formStage.Checkout()

	if supportlevelFormCallback.supportlevel == nil {
		supportlevelFormCallback.supportlevel = new(models.SupportLevel).Stage(supportlevelFormCallback.probe.stageOfInterest)
	}
	supportlevel_ := supportlevelFormCallback.supportlevel
	_ = supportlevel_

	for _, formDiv := range supportlevelFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(supportlevel_.Name), formDiv)
		case "Tool":
			FormDivSelectFieldToField(&(supportlevel_.Tool), supportlevelFormCallback.probe.stageOfInterest, formDiv)
		case "Requirement:SupportLevels":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Requirement instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Requirement instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Requirement](supportlevelFormCallback.probe.stageOfInterest)
			targetRequirementIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetRequirementIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Requirement instances and update their SupportLevels slice
			for _requirement := range *models.GetGongstructInstancesSetFromPointerType[*models.Requirement](supportlevelFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(supportlevelFormCallback.probe.stageOfInterest, _requirement)
				
				// if Requirement is selected
				if targetRequirementIDs[id] {
					// ensure supportlevel_ is in _requirement.SupportLevels
					found := false
					for _, _b := range _requirement.SupportLevels {
						if _b == supportlevel_ {
							found = true
							break
						}
					}
					if !found {
						_requirement.SupportLevels = append(_requirement.SupportLevels, supportlevel_)
						supportlevelFormCallback.probe.UpdateSliceOfPointersCallback(_requirement, "SupportLevels", &_requirement.SupportLevels)
					}
				} else {
					// ensure supportlevel_ is NOT in _requirement.SupportLevels
					idx := slices.Index(_requirement.SupportLevels, supportlevel_)
					if idx != -1 {
						_requirement.SupportLevels = slices.Delete(_requirement.SupportLevels, idx, idx+1)
						supportlevelFormCallback.probe.UpdateSliceOfPointersCallback(_requirement, "SupportLevels", &_requirement.SupportLevels)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if supportlevelFormCallback.formGroup.HasSuppressButtonBeenPressed {
		supportlevel_.Unstage(supportlevelFormCallback.probe.stageOfInterest)
	}

	supportlevelFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.SupportLevel](
		supportlevelFormCallback.probe,
	)

	// display a new form by reset the form stage
	if supportlevelFormCallback.CreationMode || supportlevelFormCallback.formGroup.HasSuppressButtonBeenPressed {
		supportlevelFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(supportlevelFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SupportLevelFormCallback(
			nil,
			supportlevelFormCallback.probe,
			newFormGroup,
		)
		supportlevel := new(models.SupportLevel)
		FillUpForm(supportlevel, newFormGroup, supportlevelFormCallback.probe)
		supportlevelFormCallback.probe.formStage.Commit()
	}

	supportlevelFormCallback.probe.ux_tree()
}
func __gong__New__ToolFormCallback(
	tool *models.Tool,
	probe *Probe,
	formGroup *form.FormGroup,
) (toolFormCallback *ToolFormCallback) {
	toolFormCallback = new(ToolFormCallback)
	toolFormCallback.probe = probe
	toolFormCallback.tool = tool
	toolFormCallback.formGroup = formGroup

	toolFormCallback.CreationMode = (tool == nil)

	return
}

type ToolFormCallback struct {
	tool *models.Tool

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (toolFormCallback *ToolFormCallback) OnSave() {
	toolFormCallback.probe.stageOfInterest.Lock()
	defer toolFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ToolFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	toolFormCallback.probe.formStage.Checkout()

	if toolFormCallback.tool == nil {
		toolFormCallback.tool = new(models.Tool).Stage(toolFormCallback.probe.stageOfInterest)
	}
	tool_ := toolFormCallback.tool
	_ = tool_

	for _, formDiv := range toolFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(tool_.Name), formDiv)
		case "Concept:Tools":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Concept instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Concept instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Concept](toolFormCallback.probe.stageOfInterest)
			targetConceptIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetConceptIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Concept instances and update their Tools slice
			for _concept := range *models.GetGongstructInstancesSetFromPointerType[*models.Concept](toolFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(toolFormCallback.probe.stageOfInterest, _concept)
				
				// if Concept is selected
				if targetConceptIDs[id] {
					// ensure tool_ is in _concept.Tools
					found := false
					for _, _b := range _concept.Tools {
						if _b == tool_ {
							found = true
							break
						}
					}
					if !found {
						_concept.Tools = append(_concept.Tools, tool_)
						toolFormCallback.probe.UpdateSliceOfPointersCallback(_concept, "Tools", &_concept.Tools)
					}
				} else {
					// ensure tool_ is NOT in _concept.Tools
					idx := slices.Index(_concept.Tools, tool_)
					if idx != -1 {
						_concept.Tools = slices.Delete(_concept.Tools, idx, idx+1)
						toolFormCallback.probe.UpdateSliceOfPointersCallback(_concept, "Tools", &_concept.Tools)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if toolFormCallback.formGroup.HasSuppressButtonBeenPressed {
		tool_.Unstage(toolFormCallback.probe.stageOfInterest)
	}

	toolFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Tool](
		toolFormCallback.probe,
	)

	// display a new form by reset the form stage
	if toolFormCallback.CreationMode || toolFormCallback.formGroup.HasSuppressButtonBeenPressed {
		toolFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(toolFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ToolFormCallback(
			nil,
			toolFormCallback.probe,
			newFormGroup,
		)
		tool := new(models.Tool)
		FillUpForm(tool, newFormGroup, toolFormCallback.probe)
		toolFormCallback.probe.formStage.Commit()
	}

	toolFormCallback.probe.ux_tree()
}
