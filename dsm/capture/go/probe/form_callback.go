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
		case "Diagram:ConceptsWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](conceptFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their ConceptsWhoseNodeIsExpanded slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](conceptFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(conceptFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure concept_ is in _diagram.ConceptsWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagram.ConceptsWhoseNodeIsExpanded {
						if _b == concept_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.ConceptsWhoseNodeIsExpanded = append(_diagram.ConceptsWhoseNodeIsExpanded, concept_)
						conceptFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ConceptsWhoseNodeIsExpanded", &_diagram.ConceptsWhoseNodeIsExpanded)
					}
				} else {
					// ensure concept_ is NOT in _diagram.ConceptsWhoseNodeIsExpanded
					idx := slices.Index(_diagram.ConceptsWhoseNodeIsExpanded, concept_)
					if idx != -1 {
						_diagram.ConceptsWhoseNodeIsExpanded = slices.Delete(_diagram.ConceptsWhoseNodeIsExpanded, idx, idx+1)
						conceptFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ConceptsWhoseNodeIsExpanded", &_diagram.ConceptsWhoseNodeIsExpanded)
					}
				}
			}
		case "Diagram:ConceptsWhoseDeliverablesNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](conceptFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their ConceptsWhoseDeliverablesNodeIsExpanded slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](conceptFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(conceptFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure concept_ is in _diagram.ConceptsWhoseDeliverablesNodeIsExpanded
					found := false
					for _, _b := range _diagram.ConceptsWhoseDeliverablesNodeIsExpanded {
						if _b == concept_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.ConceptsWhoseDeliverablesNodeIsExpanded = append(_diagram.ConceptsWhoseDeliverablesNodeIsExpanded, concept_)
						conceptFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ConceptsWhoseDeliverablesNodeIsExpanded", &_diagram.ConceptsWhoseDeliverablesNodeIsExpanded)
					}
				} else {
					// ensure concept_ is NOT in _diagram.ConceptsWhoseDeliverablesNodeIsExpanded
					idx := slices.Index(_diagram.ConceptsWhoseDeliverablesNodeIsExpanded, concept_)
					if idx != -1 {
						_diagram.ConceptsWhoseDeliverablesNodeIsExpanded = slices.Delete(_diagram.ConceptsWhoseDeliverablesNodeIsExpanded, idx, idx+1)
						conceptFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "ConceptsWhoseDeliverablesNodeIsExpanded", &_diagram.ConceptsWhoseDeliverablesNodeIsExpanded)
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
func __gong__New__ConceptShapeFormCallback(
	conceptshape *models.ConceptShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (conceptshapeFormCallback *ConceptShapeFormCallback) {
	conceptshapeFormCallback = new(ConceptShapeFormCallback)
	conceptshapeFormCallback.probe = probe
	conceptshapeFormCallback.conceptshape = conceptshape
	conceptshapeFormCallback.formGroup = formGroup

	conceptshapeFormCallback.CreationMode = (conceptshape == nil)

	return
}

type ConceptShapeFormCallback struct {
	conceptshape *models.ConceptShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (conceptshapeFormCallback *ConceptShapeFormCallback) OnSave() {
	conceptshapeFormCallback.probe.stageOfInterest.Lock()
	defer conceptshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ConceptShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	conceptshapeFormCallback.probe.formStage.Checkout()

	if conceptshapeFormCallback.conceptshape == nil {
		conceptshapeFormCallback.conceptshape = new(models.ConceptShape).Stage(conceptshapeFormCallback.probe.stageOfInterest)
	}
	conceptshape_ := conceptshapeFormCallback.conceptshape
	_ = conceptshape_

	for _, formDiv := range conceptshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(conceptshape_.Name), formDiv)
		case "Concept":
			FormDivSelectFieldToField(&(conceptshape_.Concept), conceptshapeFormCallback.probe.stageOfInterest, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(conceptshape_.IsExpanded), formDiv)
		case "X":
			FormDivBasicFieldToField(&(conceptshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(conceptshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(conceptshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(conceptshape_.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(conceptshape_.IsHidden), formDiv)
		case "Diagram:Concept_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](conceptshapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their Concept_Shapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](conceptshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(conceptshapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure conceptshape_ is in _diagram.Concept_Shapes
					found := false
					for _, _b := range _diagram.Concept_Shapes {
						if _b == conceptshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.Concept_Shapes = append(_diagram.Concept_Shapes, conceptshape_)
						conceptshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "Concept_Shapes", &_diagram.Concept_Shapes)
					}
				} else {
					// ensure conceptshape_ is NOT in _diagram.Concept_Shapes
					idx := slices.Index(_diagram.Concept_Shapes, conceptshape_)
					if idx != -1 {
						_diagram.Concept_Shapes = slices.Delete(_diagram.Concept_Shapes, idx, idx+1)
						conceptshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "Concept_Shapes", &_diagram.Concept_Shapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if conceptshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		conceptshape_.Unstage(conceptshapeFormCallback.probe.stageOfInterest)
	}

	conceptshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ConceptShape](
		conceptshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if conceptshapeFormCallback.CreationMode || conceptshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		conceptshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(conceptshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ConceptShapeFormCallback(
			nil,
			conceptshapeFormCallback.probe,
			newFormGroup,
		)
		conceptshape := new(models.ConceptShape)
		FillUpForm(conceptshape, newFormGroup, conceptshapeFormCallback.probe)
		conceptshapeFormCallback.probe.formStage.Commit()
	}

	conceptshapeFormCallback.probe.ux_tree()
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
		case "ControlPointShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ControlPointShape](concerncompositionshapeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ControlPointShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ControlPointShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					concerncompositionshapeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ControlPointShape](concerncompositionshapeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			concerncompositionshape_.ControlPointShapes = instanceSlice
			concerncompositionshapeFormCallback.probe.UpdateSliceOfPointersCallback(concerncompositionshape_, "ControlPointShapes", &concerncompositionshape_.ControlPointShapes)

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
		case "ControlPointShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ControlPointShape](concerninputshapeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ControlPointShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ControlPointShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					concerninputshapeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ControlPointShape](concerninputshapeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			concerninputshape_.ControlPointShapes = instanceSlice
			concerninputshapeFormCallback.probe.UpdateSliceOfPointersCallback(concerninputshape_, "ControlPointShapes", &concerninputshape_.ControlPointShapes)

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
		case "Concern":
			FormDivSelectFieldToField(&(concernoutputshape_.Concern), concernoutputshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Deliverable":
			FormDivSelectFieldToField(&(concernoutputshape_.Deliverable), concernoutputshapeFormCallback.probe.stageOfInterest, formDiv)
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
		case "ControlPointShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ControlPointShape](concernoutputshapeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ControlPointShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ControlPointShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					concernoutputshapeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ControlPointShape](concernoutputshapeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			concernoutputshape_.ControlPointShapes = instanceSlice
			concernoutputshapeFormCallback.probe.UpdateSliceOfPointersCallback(concernoutputshape_, "ControlPointShapes", &concernoutputshape_.ControlPointShapes)

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
func __gong__New__ControlPointShapeFormCallback(
	controlpointshape *models.ControlPointShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (controlpointshapeFormCallback *ControlPointShapeFormCallback) {
	controlpointshapeFormCallback = new(ControlPointShapeFormCallback)
	controlpointshapeFormCallback.probe = probe
	controlpointshapeFormCallback.controlpointshape = controlpointshape
	controlpointshapeFormCallback.formGroup = formGroup

	controlpointshapeFormCallback.CreationMode = (controlpointshape == nil)

	return
}

type ControlPointShapeFormCallback struct {
	controlpointshape *models.ControlPointShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (controlpointshapeFormCallback *ControlPointShapeFormCallback) OnSave() {
	controlpointshapeFormCallback.probe.stageOfInterest.Lock()
	defer controlpointshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("ControlPointShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	controlpointshapeFormCallback.probe.formStage.Checkout()

	if controlpointshapeFormCallback.controlpointshape == nil {
		controlpointshapeFormCallback.controlpointshape = new(models.ControlPointShape).Stage(controlpointshapeFormCallback.probe.stageOfInterest)
	}
	controlpointshape_ := controlpointshapeFormCallback.controlpointshape
	_ = controlpointshape_

	for _, formDiv := range controlpointshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(controlpointshape_.Name), formDiv)
		case "X_Relative":
			FormDivBasicFieldToField(&(controlpointshape_.X_Relative), formDiv)
		case "Y_Relative":
			FormDivBasicFieldToField(&(controlpointshape_.Y_Relative), formDiv)
		case "IsStartShapeTheClosestShape":
			FormDivBasicFieldToField(&(controlpointshape_.IsStartShapeTheClosestShape), formDiv)
		case "ConcernCompositionShape:ControlPointShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the ConcernCompositionShape instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target ConcernCompositionShape instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.ConcernCompositionShape](controlpointshapeFormCallback.probe.stageOfInterest)
			targetConcernCompositionShapeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetConcernCompositionShapeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all ConcernCompositionShape instances and update their ControlPointShapes slice
			for _concerncompositionshape := range *models.GetGongstructInstancesSetFromPointerType[*models.ConcernCompositionShape](controlpointshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(controlpointshapeFormCallback.probe.stageOfInterest, _concerncompositionshape)
				
				// if ConcernCompositionShape is selected
				if targetConcernCompositionShapeIDs[id] {
					// ensure controlpointshape_ is in _concerncompositionshape.ControlPointShapes
					found := false
					for _, _b := range _concerncompositionshape.ControlPointShapes {
						if _b == controlpointshape_ {
							found = true
							break
						}
					}
					if !found {
						_concerncompositionshape.ControlPointShapes = append(_concerncompositionshape.ControlPointShapes, controlpointshape_)
						controlpointshapeFormCallback.probe.UpdateSliceOfPointersCallback(_concerncompositionshape, "ControlPointShapes", &_concerncompositionshape.ControlPointShapes)
					}
				} else {
					// ensure controlpointshape_ is NOT in _concerncompositionshape.ControlPointShapes
					idx := slices.Index(_concerncompositionshape.ControlPointShapes, controlpointshape_)
					if idx != -1 {
						_concerncompositionshape.ControlPointShapes = slices.Delete(_concerncompositionshape.ControlPointShapes, idx, idx+1)
						controlpointshapeFormCallback.probe.UpdateSliceOfPointersCallback(_concerncompositionshape, "ControlPointShapes", &_concerncompositionshape.ControlPointShapes)
					}
				}
			}
		case "ConcernInputShape:ControlPointShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the ConcernInputShape instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target ConcernInputShape instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.ConcernInputShape](controlpointshapeFormCallback.probe.stageOfInterest)
			targetConcernInputShapeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetConcernInputShapeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all ConcernInputShape instances and update their ControlPointShapes slice
			for _concerninputshape := range *models.GetGongstructInstancesSetFromPointerType[*models.ConcernInputShape](controlpointshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(controlpointshapeFormCallback.probe.stageOfInterest, _concerninputshape)
				
				// if ConcernInputShape is selected
				if targetConcernInputShapeIDs[id] {
					// ensure controlpointshape_ is in _concerninputshape.ControlPointShapes
					found := false
					for _, _b := range _concerninputshape.ControlPointShapes {
						if _b == controlpointshape_ {
							found = true
							break
						}
					}
					if !found {
						_concerninputshape.ControlPointShapes = append(_concerninputshape.ControlPointShapes, controlpointshape_)
						controlpointshapeFormCallback.probe.UpdateSliceOfPointersCallback(_concerninputshape, "ControlPointShapes", &_concerninputshape.ControlPointShapes)
					}
				} else {
					// ensure controlpointshape_ is NOT in _concerninputshape.ControlPointShapes
					idx := slices.Index(_concerninputshape.ControlPointShapes, controlpointshape_)
					if idx != -1 {
						_concerninputshape.ControlPointShapes = slices.Delete(_concerninputshape.ControlPointShapes, idx, idx+1)
						controlpointshapeFormCallback.probe.UpdateSliceOfPointersCallback(_concerninputshape, "ControlPointShapes", &_concerninputshape.ControlPointShapes)
					}
				}
			}
		case "ConcernOutputShape:ControlPointShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the ConcernOutputShape instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target ConcernOutputShape instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.ConcernOutputShape](controlpointshapeFormCallback.probe.stageOfInterest)
			targetConcernOutputShapeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetConcernOutputShapeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all ConcernOutputShape instances and update their ControlPointShapes slice
			for _concernoutputshape := range *models.GetGongstructInstancesSetFromPointerType[*models.ConcernOutputShape](controlpointshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(controlpointshapeFormCallback.probe.stageOfInterest, _concernoutputshape)
				
				// if ConcernOutputShape is selected
				if targetConcernOutputShapeIDs[id] {
					// ensure controlpointshape_ is in _concernoutputshape.ControlPointShapes
					found := false
					for _, _b := range _concernoutputshape.ControlPointShapes {
						if _b == controlpointshape_ {
							found = true
							break
						}
					}
					if !found {
						_concernoutputshape.ControlPointShapes = append(_concernoutputshape.ControlPointShapes, controlpointshape_)
						controlpointshapeFormCallback.probe.UpdateSliceOfPointersCallback(_concernoutputshape, "ControlPointShapes", &_concernoutputshape.ControlPointShapes)
					}
				} else {
					// ensure controlpointshape_ is NOT in _concernoutputshape.ControlPointShapes
					idx := slices.Index(_concernoutputshape.ControlPointShapes, controlpointshape_)
					if idx != -1 {
						_concernoutputshape.ControlPointShapes = slices.Delete(_concernoutputshape.ControlPointShapes, idx, idx+1)
						controlpointshapeFormCallback.probe.UpdateSliceOfPointersCallback(_concernoutputshape, "ControlPointShapes", &_concernoutputshape.ControlPointShapes)
					}
				}
			}
		case "DeliverableCompositionShape:ControlPointShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DeliverableCompositionShape instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DeliverableCompositionShape instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DeliverableCompositionShape](controlpointshapeFormCallback.probe.stageOfInterest)
			targetDeliverableCompositionShapeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDeliverableCompositionShapeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DeliverableCompositionShape instances and update their ControlPointShapes slice
			for _deliverablecompositionshape := range *models.GetGongstructInstancesSetFromPointerType[*models.DeliverableCompositionShape](controlpointshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(controlpointshapeFormCallback.probe.stageOfInterest, _deliverablecompositionshape)
				
				// if DeliverableCompositionShape is selected
				if targetDeliverableCompositionShapeIDs[id] {
					// ensure controlpointshape_ is in _deliverablecompositionshape.ControlPointShapes
					found := false
					for _, _b := range _deliverablecompositionshape.ControlPointShapes {
						if _b == controlpointshape_ {
							found = true
							break
						}
					}
					if !found {
						_deliverablecompositionshape.ControlPointShapes = append(_deliverablecompositionshape.ControlPointShapes, controlpointshape_)
						controlpointshapeFormCallback.probe.UpdateSliceOfPointersCallback(_deliverablecompositionshape, "ControlPointShapes", &_deliverablecompositionshape.ControlPointShapes)
					}
				} else {
					// ensure controlpointshape_ is NOT in _deliverablecompositionshape.ControlPointShapes
					idx := slices.Index(_deliverablecompositionshape.ControlPointShapes, controlpointshape_)
					if idx != -1 {
						_deliverablecompositionshape.ControlPointShapes = slices.Delete(_deliverablecompositionshape.ControlPointShapes, idx, idx+1)
						controlpointshapeFormCallback.probe.UpdateSliceOfPointersCallback(_deliverablecompositionshape, "ControlPointShapes", &_deliverablecompositionshape.ControlPointShapes)
					}
				}
			}
		case "DeliverableConceptShape:ControlPointShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the DeliverableConceptShape instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target DeliverableConceptShape instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.DeliverableConceptShape](controlpointshapeFormCallback.probe.stageOfInterest)
			targetDeliverableConceptShapeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDeliverableConceptShapeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all DeliverableConceptShape instances and update their ControlPointShapes slice
			for _deliverableconceptshape := range *models.GetGongstructInstancesSetFromPointerType[*models.DeliverableConceptShape](controlpointshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(controlpointshapeFormCallback.probe.stageOfInterest, _deliverableconceptshape)
				
				// if DeliverableConceptShape is selected
				if targetDeliverableConceptShapeIDs[id] {
					// ensure controlpointshape_ is in _deliverableconceptshape.ControlPointShapes
					found := false
					for _, _b := range _deliverableconceptshape.ControlPointShapes {
						if _b == controlpointshape_ {
							found = true
							break
						}
					}
					if !found {
						_deliverableconceptshape.ControlPointShapes = append(_deliverableconceptshape.ControlPointShapes, controlpointshape_)
						controlpointshapeFormCallback.probe.UpdateSliceOfPointersCallback(_deliverableconceptshape, "ControlPointShapes", &_deliverableconceptshape.ControlPointShapes)
					}
				} else {
					// ensure controlpointshape_ is NOT in _deliverableconceptshape.ControlPointShapes
					idx := slices.Index(_deliverableconceptshape.ControlPointShapes, controlpointshape_)
					if idx != -1 {
						_deliverableconceptshape.ControlPointShapes = slices.Delete(_deliverableconceptshape.ControlPointShapes, idx, idx+1)
						controlpointshapeFormCallback.probe.UpdateSliceOfPointersCallback(_deliverableconceptshape, "ControlPointShapes", &_deliverableconceptshape.ControlPointShapes)
					}
				}
			}
		case "NoteDeliverableShape:ControlPointShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the NoteDeliverableShape instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target NoteDeliverableShape instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.NoteDeliverableShape](controlpointshapeFormCallback.probe.stageOfInterest)
			targetNoteDeliverableShapeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetNoteDeliverableShapeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all NoteDeliverableShape instances and update their ControlPointShapes slice
			for _notedeliverableshape := range *models.GetGongstructInstancesSetFromPointerType[*models.NoteDeliverableShape](controlpointshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(controlpointshapeFormCallback.probe.stageOfInterest, _notedeliverableshape)
				
				// if NoteDeliverableShape is selected
				if targetNoteDeliverableShapeIDs[id] {
					// ensure controlpointshape_ is in _notedeliverableshape.ControlPointShapes
					found := false
					for _, _b := range _notedeliverableshape.ControlPointShapes {
						if _b == controlpointshape_ {
							found = true
							break
						}
					}
					if !found {
						_notedeliverableshape.ControlPointShapes = append(_notedeliverableshape.ControlPointShapes, controlpointshape_)
						controlpointshapeFormCallback.probe.UpdateSliceOfPointersCallback(_notedeliverableshape, "ControlPointShapes", &_notedeliverableshape.ControlPointShapes)
					}
				} else {
					// ensure controlpointshape_ is NOT in _notedeliverableshape.ControlPointShapes
					idx := slices.Index(_notedeliverableshape.ControlPointShapes, controlpointshape_)
					if idx != -1 {
						_notedeliverableshape.ControlPointShapes = slices.Delete(_notedeliverableshape.ControlPointShapes, idx, idx+1)
						controlpointshapeFormCallback.probe.UpdateSliceOfPointersCallback(_notedeliverableshape, "ControlPointShapes", &_notedeliverableshape.ControlPointShapes)
					}
				}
			}
		case "NoteStakeholderShape:ControlPointShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the NoteStakeholderShape instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target NoteStakeholderShape instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.NoteStakeholderShape](controlpointshapeFormCallback.probe.stageOfInterest)
			targetNoteStakeholderShapeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetNoteStakeholderShapeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all NoteStakeholderShape instances and update their ControlPointShapes slice
			for _notestakeholdershape := range *models.GetGongstructInstancesSetFromPointerType[*models.NoteStakeholderShape](controlpointshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(controlpointshapeFormCallback.probe.stageOfInterest, _notestakeholdershape)
				
				// if NoteStakeholderShape is selected
				if targetNoteStakeholderShapeIDs[id] {
					// ensure controlpointshape_ is in _notestakeholdershape.ControlPointShapes
					found := false
					for _, _b := range _notestakeholdershape.ControlPointShapes {
						if _b == controlpointshape_ {
							found = true
							break
						}
					}
					if !found {
						_notestakeholdershape.ControlPointShapes = append(_notestakeholdershape.ControlPointShapes, controlpointshape_)
						controlpointshapeFormCallback.probe.UpdateSliceOfPointersCallback(_notestakeholdershape, "ControlPointShapes", &_notestakeholdershape.ControlPointShapes)
					}
				} else {
					// ensure controlpointshape_ is NOT in _notestakeholdershape.ControlPointShapes
					idx := slices.Index(_notestakeholdershape.ControlPointShapes, controlpointshape_)
					if idx != -1 {
						_notestakeholdershape.ControlPointShapes = slices.Delete(_notestakeholdershape.ControlPointShapes, idx, idx+1)
						controlpointshapeFormCallback.probe.UpdateSliceOfPointersCallback(_notestakeholdershape, "ControlPointShapes", &_notestakeholdershape.ControlPointShapes)
					}
				}
			}
		case "NoteTaskShape:ControlPointShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the NoteTaskShape instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target NoteTaskShape instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.NoteTaskShape](controlpointshapeFormCallback.probe.stageOfInterest)
			targetNoteTaskShapeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetNoteTaskShapeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all NoteTaskShape instances and update their ControlPointShapes slice
			for _notetaskshape := range *models.GetGongstructInstancesSetFromPointerType[*models.NoteTaskShape](controlpointshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(controlpointshapeFormCallback.probe.stageOfInterest, _notetaskshape)
				
				// if NoteTaskShape is selected
				if targetNoteTaskShapeIDs[id] {
					// ensure controlpointshape_ is in _notetaskshape.ControlPointShapes
					found := false
					for _, _b := range _notetaskshape.ControlPointShapes {
						if _b == controlpointshape_ {
							found = true
							break
						}
					}
					if !found {
						_notetaskshape.ControlPointShapes = append(_notetaskshape.ControlPointShapes, controlpointshape_)
						controlpointshapeFormCallback.probe.UpdateSliceOfPointersCallback(_notetaskshape, "ControlPointShapes", &_notetaskshape.ControlPointShapes)
					}
				} else {
					// ensure controlpointshape_ is NOT in _notetaskshape.ControlPointShapes
					idx := slices.Index(_notetaskshape.ControlPointShapes, controlpointshape_)
					if idx != -1 {
						_notetaskshape.ControlPointShapes = slices.Delete(_notetaskshape.ControlPointShapes, idx, idx+1)
						controlpointshapeFormCallback.probe.UpdateSliceOfPointersCallback(_notetaskshape, "ControlPointShapes", &_notetaskshape.ControlPointShapes)
					}
				}
			}
		case "StakeholderCompositionShape:ControlPointShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the StakeholderCompositionShape instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target StakeholderCompositionShape instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.StakeholderCompositionShape](controlpointshapeFormCallback.probe.stageOfInterest)
			targetStakeholderCompositionShapeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetStakeholderCompositionShapeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all StakeholderCompositionShape instances and update their ControlPointShapes slice
			for _stakeholdercompositionshape := range *models.GetGongstructInstancesSetFromPointerType[*models.StakeholderCompositionShape](controlpointshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(controlpointshapeFormCallback.probe.stageOfInterest, _stakeholdercompositionshape)
				
				// if StakeholderCompositionShape is selected
				if targetStakeholderCompositionShapeIDs[id] {
					// ensure controlpointshape_ is in _stakeholdercompositionshape.ControlPointShapes
					found := false
					for _, _b := range _stakeholdercompositionshape.ControlPointShapes {
						if _b == controlpointshape_ {
							found = true
							break
						}
					}
					if !found {
						_stakeholdercompositionshape.ControlPointShapes = append(_stakeholdercompositionshape.ControlPointShapes, controlpointshape_)
						controlpointshapeFormCallback.probe.UpdateSliceOfPointersCallback(_stakeholdercompositionshape, "ControlPointShapes", &_stakeholdercompositionshape.ControlPointShapes)
					}
				} else {
					// ensure controlpointshape_ is NOT in _stakeholdercompositionshape.ControlPointShapes
					idx := slices.Index(_stakeholdercompositionshape.ControlPointShapes, controlpointshape_)
					if idx != -1 {
						_stakeholdercompositionshape.ControlPointShapes = slices.Delete(_stakeholdercompositionshape.ControlPointShapes, idx, idx+1)
						controlpointshapeFormCallback.probe.UpdateSliceOfPointersCallback(_stakeholdercompositionshape, "ControlPointShapes", &_stakeholdercompositionshape.ControlPointShapes)
					}
				}
			}
		case "StakeholderConcernShape:ControlPointShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the StakeholderConcernShape instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target StakeholderConcernShape instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.StakeholderConcernShape](controlpointshapeFormCallback.probe.stageOfInterest)
			targetStakeholderConcernShapeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetStakeholderConcernShapeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all StakeholderConcernShape instances and update their ControlPointShapes slice
			for _stakeholderconcernshape := range *models.GetGongstructInstancesSetFromPointerType[*models.StakeholderConcernShape](controlpointshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(controlpointshapeFormCallback.probe.stageOfInterest, _stakeholderconcernshape)
				
				// if StakeholderConcernShape is selected
				if targetStakeholderConcernShapeIDs[id] {
					// ensure controlpointshape_ is in _stakeholderconcernshape.ControlPointShapes
					found := false
					for _, _b := range _stakeholderconcernshape.ControlPointShapes {
						if _b == controlpointshape_ {
							found = true
							break
						}
					}
					if !found {
						_stakeholderconcernshape.ControlPointShapes = append(_stakeholderconcernshape.ControlPointShapes, controlpointshape_)
						controlpointshapeFormCallback.probe.UpdateSliceOfPointersCallback(_stakeholderconcernshape, "ControlPointShapes", &_stakeholderconcernshape.ControlPointShapes)
					}
				} else {
					// ensure controlpointshape_ is NOT in _stakeholderconcernshape.ControlPointShapes
					idx := slices.Index(_stakeholderconcernshape.ControlPointShapes, controlpointshape_)
					if idx != -1 {
						_stakeholderconcernshape.ControlPointShapes = slices.Delete(_stakeholderconcernshape.ControlPointShapes, idx, idx+1)
						controlpointshapeFormCallback.probe.UpdateSliceOfPointersCallback(_stakeholderconcernshape, "ControlPointShapes", &_stakeholderconcernshape.ControlPointShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if controlpointshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		controlpointshape_.Unstage(controlpointshapeFormCallback.probe.stageOfInterest)
	}

	controlpointshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.ControlPointShape](
		controlpointshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if controlpointshapeFormCallback.CreationMode || controlpointshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		controlpointshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(controlpointshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ControlPointShapeFormCallback(
			nil,
			controlpointshapeFormCallback.probe,
			newFormGroup,
		)
		controlpointshape := new(models.ControlPointShape)
		FillUpForm(controlpointshape, newFormGroup, controlpointshapeFormCallback.probe)
		controlpointshapeFormCallback.probe.formStage.Commit()
	}

	controlpointshapeFormCallback.probe.ux_tree()
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
		case "SubDeliverables":
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
			deliverable_.SubDeliverables = instanceSlice
			deliverableFormCallback.probe.UpdateSliceOfPointersCallback(deliverable_, "SubDeliverables", &deliverable_.SubDeliverables)

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
		case "Deliverable:SubDeliverables":
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

			// 3. Iterate over all Deliverable instances and update their SubDeliverables slice
			for _deliverable := range *models.GetGongstructInstancesSetFromPointerType[*models.Deliverable](deliverableFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(deliverableFormCallback.probe.stageOfInterest, _deliverable)
				
				// if Deliverable is selected
				if targetDeliverableIDs[id] {
					// ensure deliverable_ is in _deliverable.SubDeliverables
					found := false
					for _, _b := range _deliverable.SubDeliverables {
						if _b == deliverable_ {
							found = true
							break
						}
					}
					if !found {
						_deliverable.SubDeliverables = append(_deliverable.SubDeliverables, deliverable_)
						deliverableFormCallback.probe.UpdateSliceOfPointersCallback(_deliverable, "SubDeliverables", &_deliverable.SubDeliverables)
					}
				} else {
					// ensure deliverable_ is NOT in _deliverable.SubDeliverables
					idx := slices.Index(_deliverable.SubDeliverables, deliverable_)
					if idx != -1 {
						_deliverable.SubDeliverables = slices.Delete(_deliverable.SubDeliverables, idx, idx+1)
						deliverableFormCallback.probe.UpdateSliceOfPointersCallback(_deliverable, "SubDeliverables", &_deliverable.SubDeliverables)
					}
				}
			}
		case "Diagram:DeliverablesWhoseNodeIsExpanded":
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

			// 3. Iterate over all Diagram instances and update their DeliverablesWhoseNodeIsExpanded slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](deliverableFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(deliverableFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure deliverable_ is in _diagram.DeliverablesWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagram.DeliverablesWhoseNodeIsExpanded {
						if _b == deliverable_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.DeliverablesWhoseNodeIsExpanded = append(_diagram.DeliverablesWhoseNodeIsExpanded, deliverable_)
						deliverableFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "DeliverablesWhoseNodeIsExpanded", &_diagram.DeliverablesWhoseNodeIsExpanded)
					}
				} else {
					// ensure deliverable_ is NOT in _diagram.DeliverablesWhoseNodeIsExpanded
					idx := slices.Index(_diagram.DeliverablesWhoseNodeIsExpanded, deliverable_)
					if idx != -1 {
						_diagram.DeliverablesWhoseNodeIsExpanded = slices.Delete(_diagram.DeliverablesWhoseNodeIsExpanded, idx, idx+1)
						deliverableFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "DeliverablesWhoseNodeIsExpanded", &_diagram.DeliverablesWhoseNodeIsExpanded)
					}
				}
			}
		case "Diagram:DeliverablesWhoseConceptsNodeIsExpanded":
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

			// 3. Iterate over all Diagram instances and update their DeliverablesWhoseConceptsNodeIsExpanded slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](deliverableFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(deliverableFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure deliverable_ is in _diagram.DeliverablesWhoseConceptsNodeIsExpanded
					found := false
					for _, _b := range _diagram.DeliverablesWhoseConceptsNodeIsExpanded {
						if _b == deliverable_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.DeliverablesWhoseConceptsNodeIsExpanded = append(_diagram.DeliverablesWhoseConceptsNodeIsExpanded, deliverable_)
						deliverableFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "DeliverablesWhoseConceptsNodeIsExpanded", &_diagram.DeliverablesWhoseConceptsNodeIsExpanded)
					}
				} else {
					// ensure deliverable_ is NOT in _diagram.DeliverablesWhoseConceptsNodeIsExpanded
					idx := slices.Index(_diagram.DeliverablesWhoseConceptsNodeIsExpanded, deliverable_)
					if idx != -1 {
						_diagram.DeliverablesWhoseConceptsNodeIsExpanded = slices.Delete(_diagram.DeliverablesWhoseConceptsNodeIsExpanded, idx, idx+1)
						deliverableFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "DeliverablesWhoseConceptsNodeIsExpanded", &_diagram.DeliverablesWhoseConceptsNodeIsExpanded)
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
		case "Note:Deliverables":
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

			// 3. Iterate over all Note instances and update their Deliverables slice
			for _note := range *models.GetGongstructInstancesSetFromPointerType[*models.Note](deliverableFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(deliverableFormCallback.probe.stageOfInterest, _note)
				
				// if Note is selected
				if targetNoteIDs[id] {
					// ensure deliverable_ is in _note.Deliverables
					found := false
					for _, _b := range _note.Deliverables {
						if _b == deliverable_ {
							found = true
							break
						}
					}
					if !found {
						_note.Deliverables = append(_note.Deliverables, deliverable_)
						deliverableFormCallback.probe.UpdateSliceOfPointersCallback(_note, "Deliverables", &_note.Deliverables)
					}
				} else {
					// ensure deliverable_ is NOT in _note.Deliverables
					idx := slices.Index(_note.Deliverables, deliverable_)
					if idx != -1 {
						_note.Deliverables = slices.Delete(_note.Deliverables, idx, idx+1)
						deliverableFormCallback.probe.UpdateSliceOfPointersCallback(_note, "Deliverables", &_note.Deliverables)
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
func __gong__New__DeliverableCompositionShapeFormCallback(
	deliverablecompositionshape *models.DeliverableCompositionShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (deliverablecompositionshapeFormCallback *DeliverableCompositionShapeFormCallback) {
	deliverablecompositionshapeFormCallback = new(DeliverableCompositionShapeFormCallback)
	deliverablecompositionshapeFormCallback.probe = probe
	deliverablecompositionshapeFormCallback.deliverablecompositionshape = deliverablecompositionshape
	deliverablecompositionshapeFormCallback.formGroup = formGroup

	deliverablecompositionshapeFormCallback.CreationMode = (deliverablecompositionshape == nil)

	return
}

type DeliverableCompositionShapeFormCallback struct {
	deliverablecompositionshape *models.DeliverableCompositionShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (deliverablecompositionshapeFormCallback *DeliverableCompositionShapeFormCallback) OnSave() {
	deliverablecompositionshapeFormCallback.probe.stageOfInterest.Lock()
	defer deliverablecompositionshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DeliverableCompositionShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	deliverablecompositionshapeFormCallback.probe.formStage.Checkout()

	if deliverablecompositionshapeFormCallback.deliverablecompositionshape == nil {
		deliverablecompositionshapeFormCallback.deliverablecompositionshape = new(models.DeliverableCompositionShape).Stage(deliverablecompositionshapeFormCallback.probe.stageOfInterest)
	}
	deliverablecompositionshape_ := deliverablecompositionshapeFormCallback.deliverablecompositionshape
	_ = deliverablecompositionshape_

	for _, formDiv := range deliverablecompositionshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(deliverablecompositionshape_.Name), formDiv)
		case "Deliverable":
			FormDivSelectFieldToField(&(deliverablecompositionshape_.Deliverable), deliverablecompositionshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(deliverablecompositionshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(deliverablecompositionshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(deliverablecompositionshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(deliverablecompositionshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(deliverablecompositionshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(deliverablecompositionshape_.IsHidden), formDiv)
		case "ControlPointShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ControlPointShape](deliverablecompositionshapeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ControlPointShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ControlPointShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					deliverablecompositionshapeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ControlPointShape](deliverablecompositionshapeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			deliverablecompositionshape_.ControlPointShapes = instanceSlice
			deliverablecompositionshapeFormCallback.probe.UpdateSliceOfPointersCallback(deliverablecompositionshape_, "ControlPointShapes", &deliverablecompositionshape_.ControlPointShapes)

		case "Diagram:DeliverableComposition_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](deliverablecompositionshapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their DeliverableComposition_Shapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](deliverablecompositionshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(deliverablecompositionshapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure deliverablecompositionshape_ is in _diagram.DeliverableComposition_Shapes
					found := false
					for _, _b := range _diagram.DeliverableComposition_Shapes {
						if _b == deliverablecompositionshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.DeliverableComposition_Shapes = append(_diagram.DeliverableComposition_Shapes, deliverablecompositionshape_)
						deliverablecompositionshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "DeliverableComposition_Shapes", &_diagram.DeliverableComposition_Shapes)
					}
				} else {
					// ensure deliverablecompositionshape_ is NOT in _diagram.DeliverableComposition_Shapes
					idx := slices.Index(_diagram.DeliverableComposition_Shapes, deliverablecompositionshape_)
					if idx != -1 {
						_diagram.DeliverableComposition_Shapes = slices.Delete(_diagram.DeliverableComposition_Shapes, idx, idx+1)
						deliverablecompositionshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "DeliverableComposition_Shapes", &_diagram.DeliverableComposition_Shapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if deliverablecompositionshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		deliverablecompositionshape_.Unstage(deliverablecompositionshapeFormCallback.probe.stageOfInterest)
	}

	deliverablecompositionshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.DeliverableCompositionShape](
		deliverablecompositionshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if deliverablecompositionshapeFormCallback.CreationMode || deliverablecompositionshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		deliverablecompositionshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(deliverablecompositionshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DeliverableCompositionShapeFormCallback(
			nil,
			deliverablecompositionshapeFormCallback.probe,
			newFormGroup,
		)
		deliverablecompositionshape := new(models.DeliverableCompositionShape)
		FillUpForm(deliverablecompositionshape, newFormGroup, deliverablecompositionshapeFormCallback.probe)
		deliverablecompositionshapeFormCallback.probe.formStage.Commit()
	}

	deliverablecompositionshapeFormCallback.probe.ux_tree()
}
func __gong__New__DeliverableConceptShapeFormCallback(
	deliverableconceptshape *models.DeliverableConceptShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (deliverableconceptshapeFormCallback *DeliverableConceptShapeFormCallback) {
	deliverableconceptshapeFormCallback = new(DeliverableConceptShapeFormCallback)
	deliverableconceptshapeFormCallback.probe = probe
	deliverableconceptshapeFormCallback.deliverableconceptshape = deliverableconceptshape
	deliverableconceptshapeFormCallback.formGroup = formGroup

	deliverableconceptshapeFormCallback.CreationMode = (deliverableconceptshape == nil)

	return
}

type DeliverableConceptShapeFormCallback struct {
	deliverableconceptshape *models.DeliverableConceptShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (deliverableconceptshapeFormCallback *DeliverableConceptShapeFormCallback) OnSave() {
	deliverableconceptshapeFormCallback.probe.stageOfInterest.Lock()
	defer deliverableconceptshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DeliverableConceptShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	deliverableconceptshapeFormCallback.probe.formStage.Checkout()

	if deliverableconceptshapeFormCallback.deliverableconceptshape == nil {
		deliverableconceptshapeFormCallback.deliverableconceptshape = new(models.DeliverableConceptShape).Stage(deliverableconceptshapeFormCallback.probe.stageOfInterest)
	}
	deliverableconceptshape_ := deliverableconceptshapeFormCallback.deliverableconceptshape
	_ = deliverableconceptshape_

	for _, formDiv := range deliverableconceptshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(deliverableconceptshape_.Name), formDiv)
		case "Deliverable":
			FormDivSelectFieldToField(&(deliverableconceptshape_.Deliverable), deliverableconceptshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Concept":
			FormDivSelectFieldToField(&(deliverableconceptshape_.Concept), deliverableconceptshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(deliverableconceptshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(deliverableconceptshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(deliverableconceptshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(deliverableconceptshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(deliverableconceptshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(deliverableconceptshape_.IsHidden), formDiv)
		case "ControlPointShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ControlPointShape](deliverableconceptshapeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ControlPointShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ControlPointShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					deliverableconceptshapeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ControlPointShape](deliverableconceptshapeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			deliverableconceptshape_.ControlPointShapes = instanceSlice
			deliverableconceptshapeFormCallback.probe.UpdateSliceOfPointersCallback(deliverableconceptshape_, "ControlPointShapes", &deliverableconceptshape_.ControlPointShapes)

		case "Diagram:DeliverableConceptShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](deliverableconceptshapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their DeliverableConceptShapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](deliverableconceptshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(deliverableconceptshapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure deliverableconceptshape_ is in _diagram.DeliverableConceptShapes
					found := false
					for _, _b := range _diagram.DeliverableConceptShapes {
						if _b == deliverableconceptshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.DeliverableConceptShapes = append(_diagram.DeliverableConceptShapes, deliverableconceptshape_)
						deliverableconceptshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "DeliverableConceptShapes", &_diagram.DeliverableConceptShapes)
					}
				} else {
					// ensure deliverableconceptshape_ is NOT in _diagram.DeliverableConceptShapes
					idx := slices.Index(_diagram.DeliverableConceptShapes, deliverableconceptshape_)
					if idx != -1 {
						_diagram.DeliverableConceptShapes = slices.Delete(_diagram.DeliverableConceptShapes, idx, idx+1)
						deliverableconceptshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "DeliverableConceptShapes", &_diagram.DeliverableConceptShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if deliverableconceptshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		deliverableconceptshape_.Unstage(deliverableconceptshapeFormCallback.probe.stageOfInterest)
	}

	deliverableconceptshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.DeliverableConceptShape](
		deliverableconceptshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if deliverableconceptshapeFormCallback.CreationMode || deliverableconceptshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		deliverableconceptshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(deliverableconceptshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DeliverableConceptShapeFormCallback(
			nil,
			deliverableconceptshapeFormCallback.probe,
			newFormGroup,
		)
		deliverableconceptshape := new(models.DeliverableConceptShape)
		FillUpForm(deliverableconceptshape, newFormGroup, deliverableconceptshapeFormCallback.probe)
		deliverableconceptshapeFormCallback.probe.formStage.Commit()
	}

	deliverableconceptshapeFormCallback.probe.ux_tree()
}
func __gong__New__DeliverableShapeFormCallback(
	deliverableshape *models.DeliverableShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (deliverableshapeFormCallback *DeliverableShapeFormCallback) {
	deliverableshapeFormCallback = new(DeliverableShapeFormCallback)
	deliverableshapeFormCallback.probe = probe
	deliverableshapeFormCallback.deliverableshape = deliverableshape
	deliverableshapeFormCallback.formGroup = formGroup

	deliverableshapeFormCallback.CreationMode = (deliverableshape == nil)

	return
}

type DeliverableShapeFormCallback struct {
	deliverableshape *models.DeliverableShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (deliverableshapeFormCallback *DeliverableShapeFormCallback) OnSave() {
	deliverableshapeFormCallback.probe.stageOfInterest.Lock()
	defer deliverableshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("DeliverableShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	deliverableshapeFormCallback.probe.formStage.Checkout()

	if deliverableshapeFormCallback.deliverableshape == nil {
		deliverableshapeFormCallback.deliverableshape = new(models.DeliverableShape).Stage(deliverableshapeFormCallback.probe.stageOfInterest)
	}
	deliverableshape_ := deliverableshapeFormCallback.deliverableshape
	_ = deliverableshape_

	for _, formDiv := range deliverableshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(deliverableshape_.Name), formDiv)
		case "Deliverable":
			FormDivSelectFieldToField(&(deliverableshape_.Deliverable), deliverableshapeFormCallback.probe.stageOfInterest, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(deliverableshape_.IsExpanded), formDiv)
		case "X":
			FormDivBasicFieldToField(&(deliverableshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(deliverableshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(deliverableshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(deliverableshape_.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(deliverableshape_.IsHidden), formDiv)
		case "Diagram:Deliverable_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](deliverableshapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their Deliverable_Shapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](deliverableshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(deliverableshapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure deliverableshape_ is in _diagram.Deliverable_Shapes
					found := false
					for _, _b := range _diagram.Deliverable_Shapes {
						if _b == deliverableshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.Deliverable_Shapes = append(_diagram.Deliverable_Shapes, deliverableshape_)
						deliverableshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "Deliverable_Shapes", &_diagram.Deliverable_Shapes)
					}
				} else {
					// ensure deliverableshape_ is NOT in _diagram.Deliverable_Shapes
					idx := slices.Index(_diagram.Deliverable_Shapes, deliverableshape_)
					if idx != -1 {
						_diagram.Deliverable_Shapes = slices.Delete(_diagram.Deliverable_Shapes, idx, idx+1)
						deliverableshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "Deliverable_Shapes", &_diagram.Deliverable_Shapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if deliverableshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		deliverableshape_.Unstage(deliverableshapeFormCallback.probe.stageOfInterest)
	}

	deliverableshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.DeliverableShape](
		deliverableshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if deliverableshapeFormCallback.CreationMode || deliverableshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		deliverableshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(deliverableshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__DeliverableShapeFormCallback(
			nil,
			deliverableshapeFormCallback.probe,
			newFormGroup,
		)
		deliverableshape := new(models.DeliverableShape)
		FillUpForm(deliverableshape, newFormGroup, deliverableshapeFormCallback.probe)
		deliverableshapeFormCallback.probe.formStage.Commit()
	}

	deliverableshapeFormCallback.probe.ux_tree()
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
		case "Deliverable_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DeliverableShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DeliverableShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DeliverableShape)

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
			map_RowID_ID := GetMap_RowID_ID[*models.DeliverableShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.Deliverable_Shapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "Deliverable_Shapes", &diagram_.Deliverable_Shapes)

		case "DeliverablesWhoseNodeIsExpanded":
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
			diagram_.DeliverablesWhoseNodeIsExpanded = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "DeliverablesWhoseNodeIsExpanded", &diagram_.DeliverablesWhoseNodeIsExpanded)

		case "DeliverablesWhoseConceptsNodeIsExpanded":
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
			diagram_.DeliverablesWhoseConceptsNodeIsExpanded = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "DeliverablesWhoseConceptsNodeIsExpanded", &diagram_.DeliverablesWhoseConceptsNodeIsExpanded)

		case "IsPBSNodeExpanded":
			FormDivBasicFieldToField(&(diagram_.IsPBSNodeExpanded), formDiv)
		case "DeliverableComposition_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DeliverableCompositionShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DeliverableCompositionShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DeliverableCompositionShape)

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
			map_RowID_ID := GetMap_RowID_ID[*models.DeliverableCompositionShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.DeliverableComposition_Shapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "DeliverableComposition_Shapes", &diagram_.DeliverableComposition_Shapes)

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
		case "NoteDeliverableShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.NoteDeliverableShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.NoteDeliverableShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.NoteDeliverableShape)

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
			map_RowID_ID := GetMap_RowID_ID[*models.NoteDeliverableShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.NoteDeliverableShapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "NoteDeliverableShapes", &diagram_.NoteDeliverableShapes)

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

		case "Requirement_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.RequirementShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.RequirementShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.RequirementShape)

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
			map_RowID_ID := GetMap_RowID_ID[*models.RequirementShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.Requirement_Shapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "Requirement_Shapes", &diagram_.Requirement_Shapes)

		case "RequirementsWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Requirement](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Requirement, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Requirement)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Requirement](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.RequirementsWhoseNodeIsExpanded = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "RequirementsWhoseNodeIsExpanded", &diagram_.RequirementsWhoseNodeIsExpanded)

		case "Concept_Shapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ConceptShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ConceptShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ConceptShape)

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
			map_RowID_ID := GetMap_RowID_ID[*models.ConceptShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.Concept_Shapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "Concept_Shapes", &diagram_.Concept_Shapes)

		case "ConceptsWhoseNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Concept](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Concept, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Concept)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Concept](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ConceptsWhoseNodeIsExpanded = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "ConceptsWhoseNodeIsExpanded", &diagram_.ConceptsWhoseNodeIsExpanded)

		case "ConceptsWhoseDeliverablesNodeIsExpanded":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Concept](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Concept, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Concept)

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
			map_RowID_ID := GetMap_RowID_ID[*models.Concept](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.ConceptsWhoseDeliverablesNodeIsExpanded = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "ConceptsWhoseDeliverablesNodeIsExpanded", &diagram_.ConceptsWhoseDeliverablesNodeIsExpanded)

		case "DeliverableConceptShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.DeliverableConceptShape](diagramFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.DeliverableConceptShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.DeliverableConceptShape)

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
			map_RowID_ID := GetMap_RowID_ID[*models.DeliverableConceptShape](diagramFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			diagram_.DeliverableConceptShapes = instanceSlice
			diagramFormCallback.probe.UpdateSliceOfPointersCallback(diagram_, "DeliverableConceptShapes", &diagram_.DeliverableConceptShapes)

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
		case "Deliverables":
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
			note_.Deliverables = instanceSlice
			noteFormCallback.probe.UpdateSliceOfPointersCallback(note_, "Deliverables", &note_.Deliverables)

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
func __gong__New__NoteDeliverableShapeFormCallback(
	notedeliverableshape *models.NoteDeliverableShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (notedeliverableshapeFormCallback *NoteDeliverableShapeFormCallback) {
	notedeliverableshapeFormCallback = new(NoteDeliverableShapeFormCallback)
	notedeliverableshapeFormCallback.probe = probe
	notedeliverableshapeFormCallback.notedeliverableshape = notedeliverableshape
	notedeliverableshapeFormCallback.formGroup = formGroup

	notedeliverableshapeFormCallback.CreationMode = (notedeliverableshape == nil)

	return
}

type NoteDeliverableShapeFormCallback struct {
	notedeliverableshape *models.NoteDeliverableShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (notedeliverableshapeFormCallback *NoteDeliverableShapeFormCallback) OnSave() {
	notedeliverableshapeFormCallback.probe.stageOfInterest.Lock()
	defer notedeliverableshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("NoteDeliverableShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	notedeliverableshapeFormCallback.probe.formStage.Checkout()

	if notedeliverableshapeFormCallback.notedeliverableshape == nil {
		notedeliverableshapeFormCallback.notedeliverableshape = new(models.NoteDeliverableShape).Stage(notedeliverableshapeFormCallback.probe.stageOfInterest)
	}
	notedeliverableshape_ := notedeliverableshapeFormCallback.notedeliverableshape
	_ = notedeliverableshape_

	for _, formDiv := range notedeliverableshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(notedeliverableshape_.Name), formDiv)
		case "Note":
			FormDivSelectFieldToField(&(notedeliverableshape_.Note), notedeliverableshapeFormCallback.probe.stageOfInterest, formDiv)
		case "Deliverable":
			FormDivSelectFieldToField(&(notedeliverableshape_.Deliverable), notedeliverableshapeFormCallback.probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(notedeliverableshape_.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(notedeliverableshape_.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(notedeliverableshape_.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(notedeliverableshape_.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(notedeliverableshape_.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(notedeliverableshape_.IsHidden), formDiv)
		case "ControlPointShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ControlPointShape](notedeliverableshapeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ControlPointShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ControlPointShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					notedeliverableshapeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ControlPointShape](notedeliverableshapeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			notedeliverableshape_.ControlPointShapes = instanceSlice
			notedeliverableshapeFormCallback.probe.UpdateSliceOfPointersCallback(notedeliverableshape_, "ControlPointShapes", &notedeliverableshape_.ControlPointShapes)

		case "Diagram:NoteDeliverableShapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](notedeliverableshapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their NoteDeliverableShapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](notedeliverableshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(notedeliverableshapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure notedeliverableshape_ is in _diagram.NoteDeliverableShapes
					found := false
					for _, _b := range _diagram.NoteDeliverableShapes {
						if _b == notedeliverableshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.NoteDeliverableShapes = append(_diagram.NoteDeliverableShapes, notedeliverableshape_)
						notedeliverableshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "NoteDeliverableShapes", &_diagram.NoteDeliverableShapes)
					}
				} else {
					// ensure notedeliverableshape_ is NOT in _diagram.NoteDeliverableShapes
					idx := slices.Index(_diagram.NoteDeliverableShapes, notedeliverableshape_)
					if idx != -1 {
						_diagram.NoteDeliverableShapes = slices.Delete(_diagram.NoteDeliverableShapes, idx, idx+1)
						notedeliverableshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "NoteDeliverableShapes", &_diagram.NoteDeliverableShapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if notedeliverableshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		notedeliverableshape_.Unstage(notedeliverableshapeFormCallback.probe.stageOfInterest)
	}

	notedeliverableshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.NoteDeliverableShape](
		notedeliverableshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if notedeliverableshapeFormCallback.CreationMode || notedeliverableshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		notedeliverableshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(notedeliverableshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__NoteDeliverableShapeFormCallback(
			nil,
			notedeliverableshapeFormCallback.probe,
			newFormGroup,
		)
		notedeliverableshape := new(models.NoteDeliverableShape)
		FillUpForm(notedeliverableshape, newFormGroup, notedeliverableshapeFormCallback.probe)
		notedeliverableshapeFormCallback.probe.formStage.Commit()
	}

	notedeliverableshapeFormCallback.probe.ux_tree()
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
		case "ControlPointShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ControlPointShape](notestakeholdershapeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ControlPointShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ControlPointShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					notestakeholdershapeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ControlPointShape](notestakeholdershapeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			notestakeholdershape_.ControlPointShapes = instanceSlice
			notestakeholdershapeFormCallback.probe.UpdateSliceOfPointersCallback(notestakeholdershape_, "ControlPointShapes", &notestakeholdershape_.ControlPointShapes)

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
		case "ControlPointShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ControlPointShape](notetaskshapeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ControlPointShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ControlPointShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					notetaskshapeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ControlPointShape](notetaskshapeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			notetaskshape_.ControlPointShapes = instanceSlice
			notetaskshapeFormCallback.probe.UpdateSliceOfPointersCallback(notetaskshape_, "ControlPointShapes", &notetaskshape_.ControlPointShapes)

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
		case "Diagram:RequirementsWhoseNodeIsExpanded":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](requirementFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their RequirementsWhoseNodeIsExpanded slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](requirementFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(requirementFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure requirement_ is in _diagram.RequirementsWhoseNodeIsExpanded
					found := false
					for _, _b := range _diagram.RequirementsWhoseNodeIsExpanded {
						if _b == requirement_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.RequirementsWhoseNodeIsExpanded = append(_diagram.RequirementsWhoseNodeIsExpanded, requirement_)
						requirementFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "RequirementsWhoseNodeIsExpanded", &_diagram.RequirementsWhoseNodeIsExpanded)
					}
				} else {
					// ensure requirement_ is NOT in _diagram.RequirementsWhoseNodeIsExpanded
					idx := slices.Index(_diagram.RequirementsWhoseNodeIsExpanded, requirement_)
					if idx != -1 {
						_diagram.RequirementsWhoseNodeIsExpanded = slices.Delete(_diagram.RequirementsWhoseNodeIsExpanded, idx, idx+1)
						requirementFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "RequirementsWhoseNodeIsExpanded", &_diagram.RequirementsWhoseNodeIsExpanded)
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
func __gong__New__RequirementShapeFormCallback(
	requirementshape *models.RequirementShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (requirementshapeFormCallback *RequirementShapeFormCallback) {
	requirementshapeFormCallback = new(RequirementShapeFormCallback)
	requirementshapeFormCallback.probe = probe
	requirementshapeFormCallback.requirementshape = requirementshape
	requirementshapeFormCallback.formGroup = formGroup

	requirementshapeFormCallback.CreationMode = (requirementshape == nil)

	return
}

type RequirementShapeFormCallback struct {
	requirementshape *models.RequirementShape

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (requirementshapeFormCallback *RequirementShapeFormCallback) OnSave() {
	requirementshapeFormCallback.probe.stageOfInterest.Lock()
	defer requirementshapeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("RequirementShapeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	requirementshapeFormCallback.probe.formStage.Checkout()

	if requirementshapeFormCallback.requirementshape == nil {
		requirementshapeFormCallback.requirementshape = new(models.RequirementShape).Stage(requirementshapeFormCallback.probe.stageOfInterest)
	}
	requirementshape_ := requirementshapeFormCallback.requirementshape
	_ = requirementshape_

	for _, formDiv := range requirementshapeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(requirementshape_.Name), formDiv)
		case "Requirement":
			FormDivSelectFieldToField(&(requirementshape_.Requirement), requirementshapeFormCallback.probe.stageOfInterest, formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(requirementshape_.IsExpanded), formDiv)
		case "X":
			FormDivBasicFieldToField(&(requirementshape_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(requirementshape_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(requirementshape_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(requirementshape_.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(requirementshape_.IsHidden), formDiv)
		case "Diagram:Requirement_Shapes":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Diagram instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Diagram instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Diagram](requirementshapeFormCallback.probe.stageOfInterest)
			targetDiagramIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetDiagramIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Diagram instances and update their Requirement_Shapes slice
			for _diagram := range *models.GetGongstructInstancesSetFromPointerType[*models.Diagram](requirementshapeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(requirementshapeFormCallback.probe.stageOfInterest, _diagram)
				
				// if Diagram is selected
				if targetDiagramIDs[id] {
					// ensure requirementshape_ is in _diagram.Requirement_Shapes
					found := false
					for _, _b := range _diagram.Requirement_Shapes {
						if _b == requirementshape_ {
							found = true
							break
						}
					}
					if !found {
						_diagram.Requirement_Shapes = append(_diagram.Requirement_Shapes, requirementshape_)
						requirementshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "Requirement_Shapes", &_diagram.Requirement_Shapes)
					}
				} else {
					// ensure requirementshape_ is NOT in _diagram.Requirement_Shapes
					idx := slices.Index(_diagram.Requirement_Shapes, requirementshape_)
					if idx != -1 {
						_diagram.Requirement_Shapes = slices.Delete(_diagram.Requirement_Shapes, idx, idx+1)
						requirementshapeFormCallback.probe.UpdateSliceOfPointersCallback(_diagram, "Requirement_Shapes", &_diagram.Requirement_Shapes)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if requirementshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		requirementshape_.Unstage(requirementshapeFormCallback.probe.stageOfInterest)
	}

	requirementshapeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.RequirementShape](
		requirementshapeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if requirementshapeFormCallback.CreationMode || requirementshapeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		requirementshapeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(requirementshapeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__RequirementShapeFormCallback(
			nil,
			requirementshapeFormCallback.probe,
			newFormGroup,
		)
		requirementshape := new(models.RequirementShape)
		FillUpForm(requirementshape, newFormGroup, requirementshapeFormCallback.probe)
		requirementshapeFormCallback.probe.formStage.Commit()
	}

	requirementshapeFormCallback.probe.ux_tree()
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
		case "ControlPointShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ControlPointShape](stakeholdercompositionshapeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ControlPointShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ControlPointShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					stakeholdercompositionshapeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ControlPointShape](stakeholdercompositionshapeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			stakeholdercompositionshape_.ControlPointShapes = instanceSlice
			stakeholdercompositionshapeFormCallback.probe.UpdateSliceOfPointersCallback(stakeholdercompositionshape_, "ControlPointShapes", &stakeholdercompositionshape_.ControlPointShapes)

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
		case "ControlPointShapes":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.ControlPointShape](stakeholderconcernshapeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.ControlPointShape, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.ControlPointShape)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					stakeholderconcernshapeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.ControlPointShape](stakeholderconcernshapeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			stakeholderconcernshape_.ControlPointShapes = instanceSlice
			stakeholderconcernshapeFormCallback.probe.UpdateSliceOfPointersCallback(stakeholderconcernshape_, "ControlPointShapes", &stakeholderconcernshape_.ControlPointShapes)

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
