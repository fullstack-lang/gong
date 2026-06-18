// generated code - do not edit
package models

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

var (
	__GongSliceTemplate_time__dummyDeclaration time.Duration
	_                                          = __GongSliceTemplate_time__dummyDeclaration
)

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct AnalysisNeed
	// insertion point per field

	// Compute reverse map for named struct Concept
	// insertion point per field
	stage.Concept_Tools_reverseMap = make(map[*Tool]*Concept)
	for concept := range stage.Concepts {
		_ = concept
		for _, _tool := range concept.Tools {
			stage.Concept_Tools_reverseMap[_tool] = concept
		}
	}

	// Compute reverse map for named struct ConceptShape
	// insertion point per field

	// Compute reverse map for named struct Concern
	// insertion point per field
	stage.Concern_SubConcerns_reverseMap = make(map[*Concern]*Concern)
	for concern := range stage.Concerns {
		_ = concern
		for _, _concern := range concern.SubConcerns {
			stage.Concern_SubConcerns_reverseMap[_concern] = concern
		}
	}
	stage.Concern_Inputs_reverseMap = make(map[*Deliverable]*Concern)
	for concern := range stage.Concerns {
		_ = concern
		for _, _deliverable := range concern.Inputs {
			stage.Concern_Inputs_reverseMap[_deliverable] = concern
		}
	}
	stage.Concern_Outputs_reverseMap = make(map[*Deliverable]*Concern)
	for concern := range stage.Concerns {
		_ = concern
		for _, _deliverable := range concern.Outputs {
			stage.Concern_Outputs_reverseMap[_deliverable] = concern
		}
	}
	stage.Concern_Requirements_reverseMap = make(map[*Requirement]*Concern)
	for concern := range stage.Concerns {
		_ = concern
		for _, _requirement := range concern.Requirements {
			stage.Concern_Requirements_reverseMap[_requirement] = concern
		}
	}

	// Compute reverse map for named struct ConcernCompositionShape
	// insertion point per field

	// Compute reverse map for named struct ConcernInputShape
	// insertion point per field

	// Compute reverse map for named struct ConcernOutputShape
	// insertion point per field

	// Compute reverse map for named struct ConcernShape
	// insertion point per field

	// Compute reverse map for named struct Deliverable
	// insertion point per field
	stage.Deliverable_SubDeliverables_reverseMap = make(map[*Deliverable]*Deliverable)
	for deliverable := range stage.Deliverables {
		_ = deliverable
		for _, _deliverable := range deliverable.SubDeliverables {
			stage.Deliverable_SubDeliverables_reverseMap[_deliverable] = deliverable
		}
	}
	stage.Deliverable_Concepts_reverseMap = make(map[*Concept]*Deliverable)
	for deliverable := range stage.Deliverables {
		_ = deliverable
		for _, _concept := range deliverable.Concepts {
			stage.Deliverable_Concepts_reverseMap[_concept] = deliverable
		}
	}

	// Compute reverse map for named struct DeliverableCompositionShape
	// insertion point per field

	// Compute reverse map for named struct DeliverableConceptShape
	// insertion point per field

	// Compute reverse map for named struct DeliverableShape
	// insertion point per field

	// Compute reverse map for named struct Diagram
	// insertion point per field
	stage.Diagram_ConcernsWhoseRequirementsNodeIsExpanded_reverseMap = make(map[*Concern]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _concern := range diagram.ConcernsWhoseRequirementsNodeIsExpanded {
			stage.Diagram_ConcernsWhoseRequirementsNodeIsExpanded_reverseMap[_concern] = diagram
		}
	}
	stage.Diagram_Deliverable_Shapes_reverseMap = make(map[*DeliverableShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _deliverableshape := range diagram.Deliverable_Shapes {
			stage.Diagram_Deliverable_Shapes_reverseMap[_deliverableshape] = diagram
		}
	}
	stage.Diagram_DeliverablesWhoseNodeIsExpanded_reverseMap = make(map[*Deliverable]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _deliverable := range diagram.DeliverablesWhoseNodeIsExpanded {
			stage.Diagram_DeliverablesWhoseNodeIsExpanded_reverseMap[_deliverable] = diagram
		}
	}
	stage.Diagram_DeliverablesWhoseConceptsNodeIsExpanded_reverseMap = make(map[*Deliverable]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _deliverable := range diagram.DeliverablesWhoseConceptsNodeIsExpanded {
			stage.Diagram_DeliverablesWhoseConceptsNodeIsExpanded_reverseMap[_deliverable] = diagram
		}
	}
	stage.Diagram_DeliverableComposition_Shapes_reverseMap = make(map[*DeliverableCompositionShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _deliverablecompositionshape := range diagram.DeliverableComposition_Shapes {
			stage.Diagram_DeliverableComposition_Shapes_reverseMap[_deliverablecompositionshape] = diagram
		}
	}
	stage.Diagram_Concern_Shapes_reverseMap = make(map[*ConcernShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _concernshape := range diagram.Concern_Shapes {
			stage.Diagram_Concern_Shapes_reverseMap[_concernshape] = diagram
		}
	}
	stage.Diagram_ConcernsWhoseNodeIsExpanded_reverseMap = make(map[*Concern]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _concern := range diagram.ConcernsWhoseNodeIsExpanded {
			stage.Diagram_ConcernsWhoseNodeIsExpanded_reverseMap[_concern] = diagram
		}
	}
	stage.Diagram_ConcernsWhoseInputNodeIsExpanded_reverseMap = make(map[*Concern]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _concern := range diagram.ConcernsWhoseInputNodeIsExpanded {
			stage.Diagram_ConcernsWhoseInputNodeIsExpanded_reverseMap[_concern] = diagram
		}
	}
	stage.Diagram_ConcernsWhoseStakeholderNodeIsExpanded_reverseMap = make(map[*Concern]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _concern := range diagram.ConcernsWhoseStakeholderNodeIsExpanded {
			stage.Diagram_ConcernsWhoseStakeholderNodeIsExpanded_reverseMap[_concern] = diagram
		}
	}
	stage.Diagram_ConcernssWhoseOutputNodeIsExpanded_reverseMap = make(map[*Concern]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _concern := range diagram.ConcernssWhoseOutputNodeIsExpanded {
			stage.Diagram_ConcernssWhoseOutputNodeIsExpanded_reverseMap[_concern] = diagram
		}
	}
	stage.Diagram_ConcernComposition_Shapes_reverseMap = make(map[*ConcernCompositionShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _concerncompositionshape := range diagram.ConcernComposition_Shapes {
			stage.Diagram_ConcernComposition_Shapes_reverseMap[_concerncompositionshape] = diagram
		}
	}
	stage.Diagram_ConcernInputShapes_reverseMap = make(map[*ConcernInputShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _concerninputshape := range diagram.ConcernInputShapes {
			stage.Diagram_ConcernInputShapes_reverseMap[_concerninputshape] = diagram
		}
	}
	stage.Diagram_ConcernOutputShapes_reverseMap = make(map[*ConcernOutputShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _concernoutputshape := range diagram.ConcernOutputShapes {
			stage.Diagram_ConcernOutputShapes_reverseMap[_concernoutputshape] = diagram
		}
	}
	stage.Diagram_Note_Shapes_reverseMap = make(map[*NoteShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _noteshape := range diagram.Note_Shapes {
			stage.Diagram_Note_Shapes_reverseMap[_noteshape] = diagram
		}
	}
	stage.Diagram_NotesWhoseNodeIsExpanded_reverseMap = make(map[*Note]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _note := range diagram.NotesWhoseNodeIsExpanded {
			stage.Diagram_NotesWhoseNodeIsExpanded_reverseMap[_note] = diagram
		}
	}
	stage.Diagram_NoteDeliverableShapes_reverseMap = make(map[*NoteDeliverableShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _notedeliverableshape := range diagram.NoteDeliverableShapes {
			stage.Diagram_NoteDeliverableShapes_reverseMap[_notedeliverableshape] = diagram
		}
	}
	stage.Diagram_NoteTaskShapes_reverseMap = make(map[*NoteTaskShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _notetaskshape := range diagram.NoteTaskShapes {
			stage.Diagram_NoteTaskShapes_reverseMap[_notetaskshape] = diagram
		}
	}
	stage.Diagram_NoteResourceShapes_reverseMap = make(map[*NoteStakeholderShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _notestakeholdershape := range diagram.NoteResourceShapes {
			stage.Diagram_NoteResourceShapes_reverseMap[_notestakeholdershape] = diagram
		}
	}
	stage.Diagram_Stakeholder_Shapes_reverseMap = make(map[*StakeholderShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _stakeholdershape := range diagram.Stakeholder_Shapes {
			stage.Diagram_Stakeholder_Shapes_reverseMap[_stakeholdershape] = diagram
		}
	}
	stage.Diagram_ResourcesWhoseNodeIsExpanded_reverseMap = make(map[*Stakeholder]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _stakeholder := range diagram.ResourcesWhoseNodeIsExpanded {
			stage.Diagram_ResourcesWhoseNodeIsExpanded_reverseMap[_stakeholder] = diagram
		}
	}
	stage.Diagram_ResourceComposition_Shapes_reverseMap = make(map[*StakeholderCompositionShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _stakeholdercompositionshape := range diagram.ResourceComposition_Shapes {
			stage.Diagram_ResourceComposition_Shapes_reverseMap[_stakeholdercompositionshape] = diagram
		}
	}
	stage.Diagram_StakeholderConcernShapes_reverseMap = make(map[*StakeholderConcernShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _stakeholderconcernshape := range diagram.StakeholderConcernShapes {
			stage.Diagram_StakeholderConcernShapes_reverseMap[_stakeholderconcernshape] = diagram
		}
	}
	stage.Diagram_Requirement_Shapes_reverseMap = make(map[*RequirementShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _requirementshape := range diagram.Requirement_Shapes {
			stage.Diagram_Requirement_Shapes_reverseMap[_requirementshape] = diagram
		}
	}
	stage.Diagram_RequirementsWhoseNodeIsExpanded_reverseMap = make(map[*Requirement]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _requirement := range diagram.RequirementsWhoseNodeIsExpanded {
			stage.Diagram_RequirementsWhoseNodeIsExpanded_reverseMap[_requirement] = diagram
		}
	}
	stage.Diagram_Concept_Shapes_reverseMap = make(map[*ConceptShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _conceptshape := range diagram.Concept_Shapes {
			stage.Diagram_Concept_Shapes_reverseMap[_conceptshape] = diagram
		}
	}
	stage.Diagram_ConceptsWhoseNodeIsExpanded_reverseMap = make(map[*Concept]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _concept := range diagram.ConceptsWhoseNodeIsExpanded {
			stage.Diagram_ConceptsWhoseNodeIsExpanded_reverseMap[_concept] = diagram
		}
	}
	stage.Diagram_ConceptsWhoseDeliverablesNodeIsExpanded_reverseMap = make(map[*Concept]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _concept := range diagram.ConceptsWhoseDeliverablesNodeIsExpanded {
			stage.Diagram_ConceptsWhoseDeliverablesNodeIsExpanded_reverseMap[_concept] = diagram
		}
	}
	stage.Diagram_DeliverableConceptShapes_reverseMap = make(map[*DeliverableConceptShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _deliverableconceptshape := range diagram.DeliverableConceptShapes {
			stage.Diagram_DeliverableConceptShapes_reverseMap[_deliverableconceptshape] = diagram
		}
	}

	// Compute reverse map for named struct Library
	// insertion point per field
	stage.Library_RootDeliverables_reverseMap = make(map[*Deliverable]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _deliverable := range library.RootDeliverables {
			stage.Library_RootDeliverables_reverseMap[_deliverable] = library
		}
	}
	stage.Library_RootConcerns_reverseMap = make(map[*Concern]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _concern := range library.RootConcerns {
			stage.Library_RootConcerns_reverseMap[_concern] = library
		}
	}
	stage.Library_RootStakeholders_reverseMap = make(map[*Stakeholder]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _stakeholder := range library.RootStakeholders {
			stage.Library_RootStakeholders_reverseMap[_stakeholder] = library
		}
	}
	stage.Library_RootRequirements_reverseMap = make(map[*Requirement]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _requirement := range library.RootRequirements {
			stage.Library_RootRequirements_reverseMap[_requirement] = library
		}
	}
	stage.Library_RootConcepts_reverseMap = make(map[*Concept]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _concept := range library.RootConcepts {
			stage.Library_RootConcepts_reverseMap[_concept] = library
		}
	}
	stage.Library_AnalysisNeeds_reverseMap = make(map[*AnalysisNeed]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _analysisneed := range library.AnalysisNeeds {
			stage.Library_AnalysisNeeds_reverseMap[_analysisneed] = library
		}
	}
	stage.Library_Notes_reverseMap = make(map[*Note]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _note := range library.Notes {
			stage.Library_Notes_reverseMap[_note] = library
		}
	}
	stage.Library_Diagrams_reverseMap = make(map[*Diagram]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _diagram := range library.Diagrams {
			stage.Library_Diagrams_reverseMap[_diagram] = library
		}
	}
	stage.Library_SubLibraries_reverseMap = make(map[*Library]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _library := range library.SubLibraries {
			stage.Library_SubLibraries_reverseMap[_library] = library
		}
	}

	// Compute reverse map for named struct Note
	// insertion point per field
	stage.Note_Deliverables_reverseMap = make(map[*Deliverable]*Note)
	for note := range stage.Notes {
		_ = note
		for _, _deliverable := range note.Deliverables {
			stage.Note_Deliverables_reverseMap[_deliverable] = note
		}
	}
	stage.Note_Tasks_reverseMap = make(map[*Concern]*Note)
	for note := range stage.Notes {
		_ = note
		for _, _concern := range note.Tasks {
			stage.Note_Tasks_reverseMap[_concern] = note
		}
	}
	stage.Note_Resources_reverseMap = make(map[*Stakeholder]*Note)
	for note := range stage.Notes {
		_ = note
		for _, _stakeholder := range note.Resources {
			stage.Note_Resources_reverseMap[_stakeholder] = note
		}
	}

	// Compute reverse map for named struct NoteDeliverableShape
	// insertion point per field

	// Compute reverse map for named struct NoteShape
	// insertion point per field

	// Compute reverse map for named struct NoteStakeholderShape
	// insertion point per field

	// Compute reverse map for named struct NoteTaskShape
	// insertion point per field

	// Compute reverse map for named struct Requirement
	// insertion point per field
	stage.Requirement_SupportLevels_reverseMap = make(map[*SupportLevel]*Requirement)
	for requirement := range stage.Requirements {
		_ = requirement
		for _, _supportlevel := range requirement.SupportLevels {
			stage.Requirement_SupportLevels_reverseMap[_supportlevel] = requirement
		}
	}
	stage.Requirement_Concepts_reverseMap = make(map[*Concept]*Requirement)
	for requirement := range stage.Requirements {
		_ = requirement
		for _, _concept := range requirement.Concepts {
			stage.Requirement_Concepts_reverseMap[_concept] = requirement
		}
	}

	// Compute reverse map for named struct RequirementShape
	// insertion point per field

	// Compute reverse map for named struct Stakeholder
	// insertion point per field
	stage.Stakeholder_Concerns_reverseMap = make(map[*Concern]*Stakeholder)
	for stakeholder := range stage.Stakeholders {
		_ = stakeholder
		for _, _concern := range stakeholder.Concerns {
			stage.Stakeholder_Concerns_reverseMap[_concern] = stakeholder
		}
	}
	stage.Stakeholder_SubStakeholders_reverseMap = make(map[*Stakeholder]*Stakeholder)
	for stakeholder := range stage.Stakeholders {
		_ = stakeholder
		for _, _stakeholder := range stakeholder.SubStakeholders {
			stage.Stakeholder_SubStakeholders_reverseMap[_stakeholder] = stakeholder
		}
	}

	// Compute reverse map for named struct StakeholderCompositionShape
	// insertion point per field

	// Compute reverse map for named struct StakeholderConcernShape
	// insertion point per field

	// Compute reverse map for named struct StakeholderShape
	// insertion point per field

	// Compute reverse map for named struct SupportLevel
	// insertion point per field

	// Compute reverse map for named struct Tool
	// insertion point per field

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.AnalysisNeeds {
		res = append(res, instance)
	}

	for instance := range stage.Concepts {
		res = append(res, instance)
	}

	for instance := range stage.ConceptShapes {
		res = append(res, instance)
	}

	for instance := range stage.Concerns {
		res = append(res, instance)
	}

	for instance := range stage.ConcernCompositionShapes {
		res = append(res, instance)
	}

	for instance := range stage.ConcernInputShapes {
		res = append(res, instance)
	}

	for instance := range stage.ConcernOutputShapes {
		res = append(res, instance)
	}

	for instance := range stage.ConcernShapes {
		res = append(res, instance)
	}

	for instance := range stage.Deliverables {
		res = append(res, instance)
	}

	for instance := range stage.DeliverableCompositionShapes {
		res = append(res, instance)
	}

	for instance := range stage.DeliverableConceptShapes {
		res = append(res, instance)
	}

	for instance := range stage.DeliverableShapes {
		res = append(res, instance)
	}

	for instance := range stage.Diagrams {
		res = append(res, instance)
	}

	for instance := range stage.Librarys {
		res = append(res, instance)
	}

	for instance := range stage.Notes {
		res = append(res, instance)
	}

	for instance := range stage.NoteDeliverableShapes {
		res = append(res, instance)
	}

	for instance := range stage.NoteShapes {
		res = append(res, instance)
	}

	for instance := range stage.NoteStakeholderShapes {
		res = append(res, instance)
	}

	for instance := range stage.NoteTaskShapes {
		res = append(res, instance)
	}

	for instance := range stage.Requirements {
		res = append(res, instance)
	}

	for instance := range stage.RequirementShapes {
		res = append(res, instance)
	}

	for instance := range stage.Stakeholders {
		res = append(res, instance)
	}

	for instance := range stage.StakeholderCompositionShapes {
		res = append(res, instance)
	}

	for instance := range stage.StakeholderConcernShapes {
		res = append(res, instance)
	}

	for instance := range stage.StakeholderShapes {
		res = append(res, instance)
	}

	for instance := range stage.SupportLevels {
		res = append(res, instance)
	}

	for instance := range stage.Tools {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (analysisneed *AnalysisNeed) GongCopy() GongstructIF {
	newInstance := new(AnalysisNeed)
	analysisneed.CopyBasicFields(newInstance)
	return newInstance
}

func (concept *Concept) GongCopy() GongstructIF {
	newInstance := new(Concept)
	concept.CopyBasicFields(newInstance)
	return newInstance
}

func (conceptshape *ConceptShape) GongCopy() GongstructIF {
	newInstance := new(ConceptShape)
	conceptshape.CopyBasicFields(newInstance)
	return newInstance
}

func (concern *Concern) GongCopy() GongstructIF {
	newInstance := new(Concern)
	concern.CopyBasicFields(newInstance)
	return newInstance
}

func (concerncompositionshape *ConcernCompositionShape) GongCopy() GongstructIF {
	newInstance := new(ConcernCompositionShape)
	concerncompositionshape.CopyBasicFields(newInstance)
	return newInstance
}

func (concerninputshape *ConcernInputShape) GongCopy() GongstructIF {
	newInstance := new(ConcernInputShape)
	concerninputshape.CopyBasicFields(newInstance)
	return newInstance
}

func (concernoutputshape *ConcernOutputShape) GongCopy() GongstructIF {
	newInstance := new(ConcernOutputShape)
	concernoutputshape.CopyBasicFields(newInstance)
	return newInstance
}

func (concernshape *ConcernShape) GongCopy() GongstructIF {
	newInstance := new(ConcernShape)
	concernshape.CopyBasicFields(newInstance)
	return newInstance
}

func (deliverable *Deliverable) GongCopy() GongstructIF {
	newInstance := new(Deliverable)
	deliverable.CopyBasicFields(newInstance)
	return newInstance
}

func (deliverablecompositionshape *DeliverableCompositionShape) GongCopy() GongstructIF {
	newInstance := new(DeliverableCompositionShape)
	deliverablecompositionshape.CopyBasicFields(newInstance)
	return newInstance
}

func (deliverableconceptshape *DeliverableConceptShape) GongCopy() GongstructIF {
	newInstance := new(DeliverableConceptShape)
	deliverableconceptshape.CopyBasicFields(newInstance)
	return newInstance
}

func (deliverableshape *DeliverableShape) GongCopy() GongstructIF {
	newInstance := new(DeliverableShape)
	deliverableshape.CopyBasicFields(newInstance)
	return newInstance
}

func (diagram *Diagram) GongCopy() GongstructIF {
	newInstance := new(Diagram)
	diagram.CopyBasicFields(newInstance)
	return newInstance
}

func (library *Library) GongCopy() GongstructIF {
	newInstance := new(Library)
	library.CopyBasicFields(newInstance)
	return newInstance
}

func (note *Note) GongCopy() GongstructIF {
	newInstance := new(Note)
	note.CopyBasicFields(newInstance)
	return newInstance
}

func (notedeliverableshape *NoteDeliverableShape) GongCopy() GongstructIF {
	newInstance := new(NoteDeliverableShape)
	notedeliverableshape.CopyBasicFields(newInstance)
	return newInstance
}

func (noteshape *NoteShape) GongCopy() GongstructIF {
	newInstance := new(NoteShape)
	noteshape.CopyBasicFields(newInstance)
	return newInstance
}

func (notestakeholdershape *NoteStakeholderShape) GongCopy() GongstructIF {
	newInstance := new(NoteStakeholderShape)
	notestakeholdershape.CopyBasicFields(newInstance)
	return newInstance
}

func (notetaskshape *NoteTaskShape) GongCopy() GongstructIF {
	newInstance := new(NoteTaskShape)
	notetaskshape.CopyBasicFields(newInstance)
	return newInstance
}

func (requirement *Requirement) GongCopy() GongstructIF {
	newInstance := new(Requirement)
	requirement.CopyBasicFields(newInstance)
	return newInstance
}

func (requirementshape *RequirementShape) GongCopy() GongstructIF {
	newInstance := new(RequirementShape)
	requirementshape.CopyBasicFields(newInstance)
	return newInstance
}

func (stakeholder *Stakeholder) GongCopy() GongstructIF {
	newInstance := new(Stakeholder)
	stakeholder.CopyBasicFields(newInstance)
	return newInstance
}

func (stakeholdercompositionshape *StakeholderCompositionShape) GongCopy() GongstructIF {
	newInstance := new(StakeholderCompositionShape)
	stakeholdercompositionshape.CopyBasicFields(newInstance)
	return newInstance
}

func (stakeholderconcernshape *StakeholderConcernShape) GongCopy() GongstructIF {
	newInstance := new(StakeholderConcernShape)
	stakeholderconcernshape.CopyBasicFields(newInstance)
	return newInstance
}

func (stakeholdershape *StakeholderShape) GongCopy() GongstructIF {
	newInstance := new(StakeholderShape)
	stakeholdershape.CopyBasicFields(newInstance)
	return newInstance
}

func (supportlevel *SupportLevel) GongCopy() GongstructIF {
	newInstance := new(SupportLevel)
	supportlevel.CopyBasicFields(newInstance)
	return newInstance
}

func (tool *Tool) GongCopy() GongstructIF {
	newInstance := new(Tool)
	tool.CopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (analysisneed *AnalysisNeed) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(analysisneed).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(analysisneed), uint64(GetOrderPointerGongstruct(stage, analysisneed)))
	return
}

func (concept *Concept) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(concept).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(concept), uint64(GetOrderPointerGongstruct(stage, concept)))
	return
}

func (conceptshape *ConceptShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(conceptshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(conceptshape), uint64(GetOrderPointerGongstruct(stage, conceptshape)))
	return
}

func (concern *Concern) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(concern).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(concern), uint64(GetOrderPointerGongstruct(stage, concern)))
	return
}

func (concerncompositionshape *ConcernCompositionShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(concerncompositionshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(concerncompositionshape), uint64(GetOrderPointerGongstruct(stage, concerncompositionshape)))
	return
}

func (concerninputshape *ConcernInputShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(concerninputshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(concerninputshape), uint64(GetOrderPointerGongstruct(stage, concerninputshape)))
	return
}

func (concernoutputshape *ConcernOutputShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(concernoutputshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(concernoutputshape), uint64(GetOrderPointerGongstruct(stage, concernoutputshape)))
	return
}

func (concernshape *ConcernShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(concernshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(concernshape), uint64(GetOrderPointerGongstruct(stage, concernshape)))
	return
}

func (deliverable *Deliverable) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(deliverable).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(deliverable), uint64(GetOrderPointerGongstruct(stage, deliverable)))
	return
}

func (deliverablecompositionshape *DeliverableCompositionShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(deliverablecompositionshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(deliverablecompositionshape), uint64(GetOrderPointerGongstruct(stage, deliverablecompositionshape)))
	return
}

func (deliverableconceptshape *DeliverableConceptShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(deliverableconceptshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(deliverableconceptshape), uint64(GetOrderPointerGongstruct(stage, deliverableconceptshape)))
	return
}

func (deliverableshape *DeliverableShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(deliverableshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(deliverableshape), uint64(GetOrderPointerGongstruct(stage, deliverableshape)))
	return
}

func (diagram *Diagram) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(diagram).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(diagram), uint64(GetOrderPointerGongstruct(stage, diagram)))
	return
}

func (library *Library) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(library).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(library), uint64(GetOrderPointerGongstruct(stage, library)))
	return
}

func (note *Note) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(note).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(note), uint64(GetOrderPointerGongstruct(stage, note)))
	return
}

func (notedeliverableshape *NoteDeliverableShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(notedeliverableshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(notedeliverableshape), uint64(GetOrderPointerGongstruct(stage, notedeliverableshape)))
	return
}

func (noteshape *NoteShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(noteshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(noteshape), uint64(GetOrderPointerGongstruct(stage, noteshape)))
	return
}

func (notestakeholdershape *NoteStakeholderShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(notestakeholdershape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(notestakeholdershape), uint64(GetOrderPointerGongstruct(stage, notestakeholdershape)))
	return
}

func (notetaskshape *NoteTaskShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(notetaskshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(notetaskshape), uint64(GetOrderPointerGongstruct(stage, notetaskshape)))
	return
}

func (requirement *Requirement) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(requirement).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(requirement), uint64(GetOrderPointerGongstruct(stage, requirement)))
	return
}

func (requirementshape *RequirementShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(requirementshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(requirementshape), uint64(GetOrderPointerGongstruct(stage, requirementshape)))
	return
}

func (stakeholder *Stakeholder) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stakeholder).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(stakeholder), uint64(GetOrderPointerGongstruct(stage, stakeholder)))
	return
}

func (stakeholdercompositionshape *StakeholderCompositionShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stakeholdercompositionshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(stakeholdercompositionshape), uint64(GetOrderPointerGongstruct(stage, stakeholdercompositionshape)))
	return
}

func (stakeholderconcernshape *StakeholderConcernShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stakeholderconcernshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(stakeholderconcernshape), uint64(GetOrderPointerGongstruct(stage, stakeholderconcernshape)))
	return
}

func (stakeholdershape *StakeholderShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stakeholdershape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(stakeholdershape), uint64(GetOrderPointerGongstruct(stage, stakeholdershape)))
	return
}

func (supportlevel *SupportLevel) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(supportlevel).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(supportlevel), uint64(GetOrderPointerGongstruct(stage, supportlevel)))
	return
}

func (tool *Tool) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(tool).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(tool), uint64(GetOrderPointerGongstruct(stage, tool)))
	return
}

func (stage *Stage) ComputeForwardAndBackwardCommits() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	var analysisneeds_newInstances []*AnalysisNeed
	var analysisneeds_deletedInstances []*AnalysisNeed

	// parse all staged instances and check if they have a reference
	for analysisneed := range stage.AnalysisNeeds {
		if ref, ok := stage.AnalysisNeeds_reference[analysisneed]; !ok {
			analysisneeds_newInstances = append(analysisneeds_newInstances, analysisneed)
			newInstancesSlice = append(newInstancesSlice, analysisneed.GongMarshallIdentifier(stage))
			if stage.AnalysisNeeds_referenceOrder == nil {
				stage.AnalysisNeeds_referenceOrder = make(map[*AnalysisNeed]uint)
			}
			stage.AnalysisNeeds_referenceOrder[analysisneed] = stage.AnalysisNeed_stagedOrder[analysisneed]
			newInstancesReverseSlice = append(newInstancesReverseSlice, analysisneed.GongMarshallUnstaging(stage))
			// delete(stage.AnalysisNeeds_referenceOrder, analysisneed)
			fieldInitializers, pointersInitializations := analysisneed.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.AnalysisNeed_stagedOrder[ref] = stage.AnalysisNeed_stagedOrder[analysisneed]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := analysisneed.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, analysisneed)
			// delete(stage.AnalysisNeed_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if analysisneed.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", analysisneed.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.AnalysisNeeds_reference {
		instance := stage.AnalysisNeeds_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.AnalysisNeeds[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			analysisneeds_deletedInstances = append(analysisneeds_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(analysisneeds_newInstances)
	lenDeletedInstances += len(analysisneeds_deletedInstances)
	var concepts_newInstances []*Concept
	var concepts_deletedInstances []*Concept

	// parse all staged instances and check if they have a reference
	for concept := range stage.Concepts {
		if ref, ok := stage.Concepts_reference[concept]; !ok {
			concepts_newInstances = append(concepts_newInstances, concept)
			newInstancesSlice = append(newInstancesSlice, concept.GongMarshallIdentifier(stage))
			if stage.Concepts_referenceOrder == nil {
				stage.Concepts_referenceOrder = make(map[*Concept]uint)
			}
			stage.Concepts_referenceOrder[concept] = stage.Concept_stagedOrder[concept]
			newInstancesReverseSlice = append(newInstancesReverseSlice, concept.GongMarshallUnstaging(stage))
			// delete(stage.Concepts_referenceOrder, concept)
			fieldInitializers, pointersInitializations := concept.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Concept_stagedOrder[ref] = stage.Concept_stagedOrder[concept]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := concept.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, concept)
			// delete(stage.Concept_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if concept.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", concept.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Concepts_reference {
		instance := stage.Concepts_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Concepts[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			concepts_deletedInstances = append(concepts_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(concepts_newInstances)
	lenDeletedInstances += len(concepts_deletedInstances)
	var conceptshapes_newInstances []*ConceptShape
	var conceptshapes_deletedInstances []*ConceptShape

	// parse all staged instances and check if they have a reference
	for conceptshape := range stage.ConceptShapes {
		if ref, ok := stage.ConceptShapes_reference[conceptshape]; !ok {
			conceptshapes_newInstances = append(conceptshapes_newInstances, conceptshape)
			newInstancesSlice = append(newInstancesSlice, conceptshape.GongMarshallIdentifier(stage))
			if stage.ConceptShapes_referenceOrder == nil {
				stage.ConceptShapes_referenceOrder = make(map[*ConceptShape]uint)
			}
			stage.ConceptShapes_referenceOrder[conceptshape] = stage.ConceptShape_stagedOrder[conceptshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, conceptshape.GongMarshallUnstaging(stage))
			// delete(stage.ConceptShapes_referenceOrder, conceptshape)
			fieldInitializers, pointersInitializations := conceptshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ConceptShape_stagedOrder[ref] = stage.ConceptShape_stagedOrder[conceptshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := conceptshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, conceptshape)
			// delete(stage.ConceptShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if conceptshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", conceptshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.ConceptShapes_reference {
		instance := stage.ConceptShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ConceptShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			conceptshapes_deletedInstances = append(conceptshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(conceptshapes_newInstances)
	lenDeletedInstances += len(conceptshapes_deletedInstances)
	var concerns_newInstances []*Concern
	var concerns_deletedInstances []*Concern

	// parse all staged instances and check if they have a reference
	for concern := range stage.Concerns {
		if ref, ok := stage.Concerns_reference[concern]; !ok {
			concerns_newInstances = append(concerns_newInstances, concern)
			newInstancesSlice = append(newInstancesSlice, concern.GongMarshallIdentifier(stage))
			if stage.Concerns_referenceOrder == nil {
				stage.Concerns_referenceOrder = make(map[*Concern]uint)
			}
			stage.Concerns_referenceOrder[concern] = stage.Concern_stagedOrder[concern]
			newInstancesReverseSlice = append(newInstancesReverseSlice, concern.GongMarshallUnstaging(stage))
			// delete(stage.Concerns_referenceOrder, concern)
			fieldInitializers, pointersInitializations := concern.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Concern_stagedOrder[ref] = stage.Concern_stagedOrder[concern]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := concern.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, concern)
			// delete(stage.Concern_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if concern.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", concern.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Concerns_reference {
		instance := stage.Concerns_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Concerns[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			concerns_deletedInstances = append(concerns_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(concerns_newInstances)
	lenDeletedInstances += len(concerns_deletedInstances)
	var concerncompositionshapes_newInstances []*ConcernCompositionShape
	var concerncompositionshapes_deletedInstances []*ConcernCompositionShape

	// parse all staged instances and check if they have a reference
	for concerncompositionshape := range stage.ConcernCompositionShapes {
		if ref, ok := stage.ConcernCompositionShapes_reference[concerncompositionshape]; !ok {
			concerncompositionshapes_newInstances = append(concerncompositionshapes_newInstances, concerncompositionshape)
			newInstancesSlice = append(newInstancesSlice, concerncompositionshape.GongMarshallIdentifier(stage))
			if stage.ConcernCompositionShapes_referenceOrder == nil {
				stage.ConcernCompositionShapes_referenceOrder = make(map[*ConcernCompositionShape]uint)
			}
			stage.ConcernCompositionShapes_referenceOrder[concerncompositionshape] = stage.ConcernCompositionShape_stagedOrder[concerncompositionshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, concerncompositionshape.GongMarshallUnstaging(stage))
			// delete(stage.ConcernCompositionShapes_referenceOrder, concerncompositionshape)
			fieldInitializers, pointersInitializations := concerncompositionshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ConcernCompositionShape_stagedOrder[ref] = stage.ConcernCompositionShape_stagedOrder[concerncompositionshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := concerncompositionshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, concerncompositionshape)
			// delete(stage.ConcernCompositionShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if concerncompositionshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", concerncompositionshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.ConcernCompositionShapes_reference {
		instance := stage.ConcernCompositionShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ConcernCompositionShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			concerncompositionshapes_deletedInstances = append(concerncompositionshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(concerncompositionshapes_newInstances)
	lenDeletedInstances += len(concerncompositionshapes_deletedInstances)
	var concerninputshapes_newInstances []*ConcernInputShape
	var concerninputshapes_deletedInstances []*ConcernInputShape

	// parse all staged instances and check if they have a reference
	for concerninputshape := range stage.ConcernInputShapes {
		if ref, ok := stage.ConcernInputShapes_reference[concerninputshape]; !ok {
			concerninputshapes_newInstances = append(concerninputshapes_newInstances, concerninputshape)
			newInstancesSlice = append(newInstancesSlice, concerninputshape.GongMarshallIdentifier(stage))
			if stage.ConcernInputShapes_referenceOrder == nil {
				stage.ConcernInputShapes_referenceOrder = make(map[*ConcernInputShape]uint)
			}
			stage.ConcernInputShapes_referenceOrder[concerninputshape] = stage.ConcernInputShape_stagedOrder[concerninputshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, concerninputshape.GongMarshallUnstaging(stage))
			// delete(stage.ConcernInputShapes_referenceOrder, concerninputshape)
			fieldInitializers, pointersInitializations := concerninputshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ConcernInputShape_stagedOrder[ref] = stage.ConcernInputShape_stagedOrder[concerninputshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := concerninputshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, concerninputshape)
			// delete(stage.ConcernInputShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if concerninputshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", concerninputshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.ConcernInputShapes_reference {
		instance := stage.ConcernInputShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ConcernInputShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			concerninputshapes_deletedInstances = append(concerninputshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(concerninputshapes_newInstances)
	lenDeletedInstances += len(concerninputshapes_deletedInstances)
	var concernoutputshapes_newInstances []*ConcernOutputShape
	var concernoutputshapes_deletedInstances []*ConcernOutputShape

	// parse all staged instances and check if they have a reference
	for concernoutputshape := range stage.ConcernOutputShapes {
		if ref, ok := stage.ConcernOutputShapes_reference[concernoutputshape]; !ok {
			concernoutputshapes_newInstances = append(concernoutputshapes_newInstances, concernoutputshape)
			newInstancesSlice = append(newInstancesSlice, concernoutputshape.GongMarshallIdentifier(stage))
			if stage.ConcernOutputShapes_referenceOrder == nil {
				stage.ConcernOutputShapes_referenceOrder = make(map[*ConcernOutputShape]uint)
			}
			stage.ConcernOutputShapes_referenceOrder[concernoutputshape] = stage.ConcernOutputShape_stagedOrder[concernoutputshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, concernoutputshape.GongMarshallUnstaging(stage))
			// delete(stage.ConcernOutputShapes_referenceOrder, concernoutputshape)
			fieldInitializers, pointersInitializations := concernoutputshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ConcernOutputShape_stagedOrder[ref] = stage.ConcernOutputShape_stagedOrder[concernoutputshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := concernoutputshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, concernoutputshape)
			// delete(stage.ConcernOutputShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if concernoutputshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", concernoutputshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.ConcernOutputShapes_reference {
		instance := stage.ConcernOutputShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ConcernOutputShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			concernoutputshapes_deletedInstances = append(concernoutputshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(concernoutputshapes_newInstances)
	lenDeletedInstances += len(concernoutputshapes_deletedInstances)
	var concernshapes_newInstances []*ConcernShape
	var concernshapes_deletedInstances []*ConcernShape

	// parse all staged instances and check if they have a reference
	for concernshape := range stage.ConcernShapes {
		if ref, ok := stage.ConcernShapes_reference[concernshape]; !ok {
			concernshapes_newInstances = append(concernshapes_newInstances, concernshape)
			newInstancesSlice = append(newInstancesSlice, concernshape.GongMarshallIdentifier(stage))
			if stage.ConcernShapes_referenceOrder == nil {
				stage.ConcernShapes_referenceOrder = make(map[*ConcernShape]uint)
			}
			stage.ConcernShapes_referenceOrder[concernshape] = stage.ConcernShape_stagedOrder[concernshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, concernshape.GongMarshallUnstaging(stage))
			// delete(stage.ConcernShapes_referenceOrder, concernshape)
			fieldInitializers, pointersInitializations := concernshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ConcernShape_stagedOrder[ref] = stage.ConcernShape_stagedOrder[concernshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := concernshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, concernshape)
			// delete(stage.ConcernShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if concernshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", concernshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.ConcernShapes_reference {
		instance := stage.ConcernShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ConcernShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			concernshapes_deletedInstances = append(concernshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(concernshapes_newInstances)
	lenDeletedInstances += len(concernshapes_deletedInstances)
	var deliverables_newInstances []*Deliverable
	var deliverables_deletedInstances []*Deliverable

	// parse all staged instances and check if they have a reference
	for deliverable := range stage.Deliverables {
		if ref, ok := stage.Deliverables_reference[deliverable]; !ok {
			deliverables_newInstances = append(deliverables_newInstances, deliverable)
			newInstancesSlice = append(newInstancesSlice, deliverable.GongMarshallIdentifier(stage))
			if stage.Deliverables_referenceOrder == nil {
				stage.Deliverables_referenceOrder = make(map[*Deliverable]uint)
			}
			stage.Deliverables_referenceOrder[deliverable] = stage.Deliverable_stagedOrder[deliverable]
			newInstancesReverseSlice = append(newInstancesReverseSlice, deliverable.GongMarshallUnstaging(stage))
			// delete(stage.Deliverables_referenceOrder, deliverable)
			fieldInitializers, pointersInitializations := deliverable.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Deliverable_stagedOrder[ref] = stage.Deliverable_stagedOrder[deliverable]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := deliverable.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, deliverable)
			// delete(stage.Deliverable_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if deliverable.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", deliverable.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Deliverables_reference {
		instance := stage.Deliverables_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Deliverables[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			deliverables_deletedInstances = append(deliverables_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(deliverables_newInstances)
	lenDeletedInstances += len(deliverables_deletedInstances)
	var deliverablecompositionshapes_newInstances []*DeliverableCompositionShape
	var deliverablecompositionshapes_deletedInstances []*DeliverableCompositionShape

	// parse all staged instances and check if they have a reference
	for deliverablecompositionshape := range stage.DeliverableCompositionShapes {
		if ref, ok := stage.DeliverableCompositionShapes_reference[deliverablecompositionshape]; !ok {
			deliverablecompositionshapes_newInstances = append(deliverablecompositionshapes_newInstances, deliverablecompositionshape)
			newInstancesSlice = append(newInstancesSlice, deliverablecompositionshape.GongMarshallIdentifier(stage))
			if stage.DeliverableCompositionShapes_referenceOrder == nil {
				stage.DeliverableCompositionShapes_referenceOrder = make(map[*DeliverableCompositionShape]uint)
			}
			stage.DeliverableCompositionShapes_referenceOrder[deliverablecompositionshape] = stage.DeliverableCompositionShape_stagedOrder[deliverablecompositionshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, deliverablecompositionshape.GongMarshallUnstaging(stage))
			// delete(stage.DeliverableCompositionShapes_referenceOrder, deliverablecompositionshape)
			fieldInitializers, pointersInitializations := deliverablecompositionshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DeliverableCompositionShape_stagedOrder[ref] = stage.DeliverableCompositionShape_stagedOrder[deliverablecompositionshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := deliverablecompositionshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, deliverablecompositionshape)
			// delete(stage.DeliverableCompositionShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if deliverablecompositionshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", deliverablecompositionshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.DeliverableCompositionShapes_reference {
		instance := stage.DeliverableCompositionShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.DeliverableCompositionShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			deliverablecompositionshapes_deletedInstances = append(deliverablecompositionshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(deliverablecompositionshapes_newInstances)
	lenDeletedInstances += len(deliverablecompositionshapes_deletedInstances)
	var deliverableconceptshapes_newInstances []*DeliverableConceptShape
	var deliverableconceptshapes_deletedInstances []*DeliverableConceptShape

	// parse all staged instances and check if they have a reference
	for deliverableconceptshape := range stage.DeliverableConceptShapes {
		if ref, ok := stage.DeliverableConceptShapes_reference[deliverableconceptshape]; !ok {
			deliverableconceptshapes_newInstances = append(deliverableconceptshapes_newInstances, deliverableconceptshape)
			newInstancesSlice = append(newInstancesSlice, deliverableconceptshape.GongMarshallIdentifier(stage))
			if stage.DeliverableConceptShapes_referenceOrder == nil {
				stage.DeliverableConceptShapes_referenceOrder = make(map[*DeliverableConceptShape]uint)
			}
			stage.DeliverableConceptShapes_referenceOrder[deliverableconceptshape] = stage.DeliverableConceptShape_stagedOrder[deliverableconceptshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, deliverableconceptshape.GongMarshallUnstaging(stage))
			// delete(stage.DeliverableConceptShapes_referenceOrder, deliverableconceptshape)
			fieldInitializers, pointersInitializations := deliverableconceptshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DeliverableConceptShape_stagedOrder[ref] = stage.DeliverableConceptShape_stagedOrder[deliverableconceptshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := deliverableconceptshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, deliverableconceptshape)
			// delete(stage.DeliverableConceptShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if deliverableconceptshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", deliverableconceptshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.DeliverableConceptShapes_reference {
		instance := stage.DeliverableConceptShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.DeliverableConceptShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			deliverableconceptshapes_deletedInstances = append(deliverableconceptshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(deliverableconceptshapes_newInstances)
	lenDeletedInstances += len(deliverableconceptshapes_deletedInstances)
	var deliverableshapes_newInstances []*DeliverableShape
	var deliverableshapes_deletedInstances []*DeliverableShape

	// parse all staged instances and check if they have a reference
	for deliverableshape := range stage.DeliverableShapes {
		if ref, ok := stage.DeliverableShapes_reference[deliverableshape]; !ok {
			deliverableshapes_newInstances = append(deliverableshapes_newInstances, deliverableshape)
			newInstancesSlice = append(newInstancesSlice, deliverableshape.GongMarshallIdentifier(stage))
			if stage.DeliverableShapes_referenceOrder == nil {
				stage.DeliverableShapes_referenceOrder = make(map[*DeliverableShape]uint)
			}
			stage.DeliverableShapes_referenceOrder[deliverableshape] = stage.DeliverableShape_stagedOrder[deliverableshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, deliverableshape.GongMarshallUnstaging(stage))
			// delete(stage.DeliverableShapes_referenceOrder, deliverableshape)
			fieldInitializers, pointersInitializations := deliverableshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DeliverableShape_stagedOrder[ref] = stage.DeliverableShape_stagedOrder[deliverableshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := deliverableshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, deliverableshape)
			// delete(stage.DeliverableShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if deliverableshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", deliverableshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.DeliverableShapes_reference {
		instance := stage.DeliverableShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.DeliverableShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			deliverableshapes_deletedInstances = append(deliverableshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(deliverableshapes_newInstances)
	lenDeletedInstances += len(deliverableshapes_deletedInstances)
	var diagrams_newInstances []*Diagram
	var diagrams_deletedInstances []*Diagram

	// parse all staged instances and check if they have a reference
	for diagram := range stage.Diagrams {
		if ref, ok := stage.Diagrams_reference[diagram]; !ok {
			diagrams_newInstances = append(diagrams_newInstances, diagram)
			newInstancesSlice = append(newInstancesSlice, diagram.GongMarshallIdentifier(stage))
			if stage.Diagrams_referenceOrder == nil {
				stage.Diagrams_referenceOrder = make(map[*Diagram]uint)
			}
			stage.Diagrams_referenceOrder[diagram] = stage.Diagram_stagedOrder[diagram]
			newInstancesReverseSlice = append(newInstancesReverseSlice, diagram.GongMarshallUnstaging(stage))
			// delete(stage.Diagrams_referenceOrder, diagram)
			fieldInitializers, pointersInitializations := diagram.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Diagram_stagedOrder[ref] = stage.Diagram_stagedOrder[diagram]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := diagram.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, diagram)
			// delete(stage.Diagram_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if diagram.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", diagram.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Diagrams_reference {
		instance := stage.Diagrams_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Diagrams[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			diagrams_deletedInstances = append(diagrams_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(diagrams_newInstances)
	lenDeletedInstances += len(diagrams_deletedInstances)
	var librarys_newInstances []*Library
	var librarys_deletedInstances []*Library

	// parse all staged instances and check if they have a reference
	for library := range stage.Librarys {
		if ref, ok := stage.Librarys_reference[library]; !ok {
			librarys_newInstances = append(librarys_newInstances, library)
			newInstancesSlice = append(newInstancesSlice, library.GongMarshallIdentifier(stage))
			if stage.Librarys_referenceOrder == nil {
				stage.Librarys_referenceOrder = make(map[*Library]uint)
			}
			stage.Librarys_referenceOrder[library] = stage.Library_stagedOrder[library]
			newInstancesReverseSlice = append(newInstancesReverseSlice, library.GongMarshallUnstaging(stage))
			// delete(stage.Librarys_referenceOrder, library)
			fieldInitializers, pointersInitializations := library.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Library_stagedOrder[ref] = stage.Library_stagedOrder[library]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := library.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, library)
			// delete(stage.Library_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if library.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", library.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Librarys_reference {
		instance := stage.Librarys_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Librarys[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			librarys_deletedInstances = append(librarys_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(librarys_newInstances)
	lenDeletedInstances += len(librarys_deletedInstances)
	var notes_newInstances []*Note
	var notes_deletedInstances []*Note

	// parse all staged instances and check if they have a reference
	for note := range stage.Notes {
		if ref, ok := stage.Notes_reference[note]; !ok {
			notes_newInstances = append(notes_newInstances, note)
			newInstancesSlice = append(newInstancesSlice, note.GongMarshallIdentifier(stage))
			if stage.Notes_referenceOrder == nil {
				stage.Notes_referenceOrder = make(map[*Note]uint)
			}
			stage.Notes_referenceOrder[note] = stage.Note_stagedOrder[note]
			newInstancesReverseSlice = append(newInstancesReverseSlice, note.GongMarshallUnstaging(stage))
			// delete(stage.Notes_referenceOrder, note)
			fieldInitializers, pointersInitializations := note.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Note_stagedOrder[ref] = stage.Note_stagedOrder[note]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := note.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, note)
			// delete(stage.Note_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if note.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", note.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Notes_reference {
		instance := stage.Notes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Notes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			notes_deletedInstances = append(notes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(notes_newInstances)
	lenDeletedInstances += len(notes_deletedInstances)
	var notedeliverableshapes_newInstances []*NoteDeliverableShape
	var notedeliverableshapes_deletedInstances []*NoteDeliverableShape

	// parse all staged instances and check if they have a reference
	for notedeliverableshape := range stage.NoteDeliverableShapes {
		if ref, ok := stage.NoteDeliverableShapes_reference[notedeliverableshape]; !ok {
			notedeliverableshapes_newInstances = append(notedeliverableshapes_newInstances, notedeliverableshape)
			newInstancesSlice = append(newInstancesSlice, notedeliverableshape.GongMarshallIdentifier(stage))
			if stage.NoteDeliverableShapes_referenceOrder == nil {
				stage.NoteDeliverableShapes_referenceOrder = make(map[*NoteDeliverableShape]uint)
			}
			stage.NoteDeliverableShapes_referenceOrder[notedeliverableshape] = stage.NoteDeliverableShape_stagedOrder[notedeliverableshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, notedeliverableshape.GongMarshallUnstaging(stage))
			// delete(stage.NoteDeliverableShapes_referenceOrder, notedeliverableshape)
			fieldInitializers, pointersInitializations := notedeliverableshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.NoteDeliverableShape_stagedOrder[ref] = stage.NoteDeliverableShape_stagedOrder[notedeliverableshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := notedeliverableshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, notedeliverableshape)
			// delete(stage.NoteDeliverableShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if notedeliverableshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", notedeliverableshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.NoteDeliverableShapes_reference {
		instance := stage.NoteDeliverableShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.NoteDeliverableShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			notedeliverableshapes_deletedInstances = append(notedeliverableshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(notedeliverableshapes_newInstances)
	lenDeletedInstances += len(notedeliverableshapes_deletedInstances)
	var noteshapes_newInstances []*NoteShape
	var noteshapes_deletedInstances []*NoteShape

	// parse all staged instances and check if they have a reference
	for noteshape := range stage.NoteShapes {
		if ref, ok := stage.NoteShapes_reference[noteshape]; !ok {
			noteshapes_newInstances = append(noteshapes_newInstances, noteshape)
			newInstancesSlice = append(newInstancesSlice, noteshape.GongMarshallIdentifier(stage))
			if stage.NoteShapes_referenceOrder == nil {
				stage.NoteShapes_referenceOrder = make(map[*NoteShape]uint)
			}
			stage.NoteShapes_referenceOrder[noteshape] = stage.NoteShape_stagedOrder[noteshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, noteshape.GongMarshallUnstaging(stage))
			// delete(stage.NoteShapes_referenceOrder, noteshape)
			fieldInitializers, pointersInitializations := noteshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.NoteShape_stagedOrder[ref] = stage.NoteShape_stagedOrder[noteshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := noteshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, noteshape)
			// delete(stage.NoteShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if noteshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", noteshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.NoteShapes_reference {
		instance := stage.NoteShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.NoteShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			noteshapes_deletedInstances = append(noteshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(noteshapes_newInstances)
	lenDeletedInstances += len(noteshapes_deletedInstances)
	var notestakeholdershapes_newInstances []*NoteStakeholderShape
	var notestakeholdershapes_deletedInstances []*NoteStakeholderShape

	// parse all staged instances and check if they have a reference
	for notestakeholdershape := range stage.NoteStakeholderShapes {
		if ref, ok := stage.NoteStakeholderShapes_reference[notestakeholdershape]; !ok {
			notestakeholdershapes_newInstances = append(notestakeholdershapes_newInstances, notestakeholdershape)
			newInstancesSlice = append(newInstancesSlice, notestakeholdershape.GongMarshallIdentifier(stage))
			if stage.NoteStakeholderShapes_referenceOrder == nil {
				stage.NoteStakeholderShapes_referenceOrder = make(map[*NoteStakeholderShape]uint)
			}
			stage.NoteStakeholderShapes_referenceOrder[notestakeholdershape] = stage.NoteStakeholderShape_stagedOrder[notestakeholdershape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, notestakeholdershape.GongMarshallUnstaging(stage))
			// delete(stage.NoteStakeholderShapes_referenceOrder, notestakeholdershape)
			fieldInitializers, pointersInitializations := notestakeholdershape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.NoteStakeholderShape_stagedOrder[ref] = stage.NoteStakeholderShape_stagedOrder[notestakeholdershape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := notestakeholdershape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, notestakeholdershape)
			// delete(stage.NoteStakeholderShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if notestakeholdershape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", notestakeholdershape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.NoteStakeholderShapes_reference {
		instance := stage.NoteStakeholderShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.NoteStakeholderShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			notestakeholdershapes_deletedInstances = append(notestakeholdershapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(notestakeholdershapes_newInstances)
	lenDeletedInstances += len(notestakeholdershapes_deletedInstances)
	var notetaskshapes_newInstances []*NoteTaskShape
	var notetaskshapes_deletedInstances []*NoteTaskShape

	// parse all staged instances and check if they have a reference
	for notetaskshape := range stage.NoteTaskShapes {
		if ref, ok := stage.NoteTaskShapes_reference[notetaskshape]; !ok {
			notetaskshapes_newInstances = append(notetaskshapes_newInstances, notetaskshape)
			newInstancesSlice = append(newInstancesSlice, notetaskshape.GongMarshallIdentifier(stage))
			if stage.NoteTaskShapes_referenceOrder == nil {
				stage.NoteTaskShapes_referenceOrder = make(map[*NoteTaskShape]uint)
			}
			stage.NoteTaskShapes_referenceOrder[notetaskshape] = stage.NoteTaskShape_stagedOrder[notetaskshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, notetaskshape.GongMarshallUnstaging(stage))
			// delete(stage.NoteTaskShapes_referenceOrder, notetaskshape)
			fieldInitializers, pointersInitializations := notetaskshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.NoteTaskShape_stagedOrder[ref] = stage.NoteTaskShape_stagedOrder[notetaskshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := notetaskshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, notetaskshape)
			// delete(stage.NoteTaskShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if notetaskshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", notetaskshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.NoteTaskShapes_reference {
		instance := stage.NoteTaskShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.NoteTaskShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			notetaskshapes_deletedInstances = append(notetaskshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(notetaskshapes_newInstances)
	lenDeletedInstances += len(notetaskshapes_deletedInstances)
	var requirements_newInstances []*Requirement
	var requirements_deletedInstances []*Requirement

	// parse all staged instances and check if they have a reference
	for requirement := range stage.Requirements {
		if ref, ok := stage.Requirements_reference[requirement]; !ok {
			requirements_newInstances = append(requirements_newInstances, requirement)
			newInstancesSlice = append(newInstancesSlice, requirement.GongMarshallIdentifier(stage))
			if stage.Requirements_referenceOrder == nil {
				stage.Requirements_referenceOrder = make(map[*Requirement]uint)
			}
			stage.Requirements_referenceOrder[requirement] = stage.Requirement_stagedOrder[requirement]
			newInstancesReverseSlice = append(newInstancesReverseSlice, requirement.GongMarshallUnstaging(stage))
			// delete(stage.Requirements_referenceOrder, requirement)
			fieldInitializers, pointersInitializations := requirement.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Requirement_stagedOrder[ref] = stage.Requirement_stagedOrder[requirement]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := requirement.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, requirement)
			// delete(stage.Requirement_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if requirement.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", requirement.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Requirements_reference {
		instance := stage.Requirements_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Requirements[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			requirements_deletedInstances = append(requirements_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(requirements_newInstances)
	lenDeletedInstances += len(requirements_deletedInstances)
	var requirementshapes_newInstances []*RequirementShape
	var requirementshapes_deletedInstances []*RequirementShape

	// parse all staged instances and check if they have a reference
	for requirementshape := range stage.RequirementShapes {
		if ref, ok := stage.RequirementShapes_reference[requirementshape]; !ok {
			requirementshapes_newInstances = append(requirementshapes_newInstances, requirementshape)
			newInstancesSlice = append(newInstancesSlice, requirementshape.GongMarshallIdentifier(stage))
			if stage.RequirementShapes_referenceOrder == nil {
				stage.RequirementShapes_referenceOrder = make(map[*RequirementShape]uint)
			}
			stage.RequirementShapes_referenceOrder[requirementshape] = stage.RequirementShape_stagedOrder[requirementshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, requirementshape.GongMarshallUnstaging(stage))
			// delete(stage.RequirementShapes_referenceOrder, requirementshape)
			fieldInitializers, pointersInitializations := requirementshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.RequirementShape_stagedOrder[ref] = stage.RequirementShape_stagedOrder[requirementshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := requirementshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, requirementshape)
			// delete(stage.RequirementShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if requirementshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", requirementshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.RequirementShapes_reference {
		instance := stage.RequirementShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.RequirementShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			requirementshapes_deletedInstances = append(requirementshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(requirementshapes_newInstances)
	lenDeletedInstances += len(requirementshapes_deletedInstances)
	var stakeholders_newInstances []*Stakeholder
	var stakeholders_deletedInstances []*Stakeholder

	// parse all staged instances and check if they have a reference
	for stakeholder := range stage.Stakeholders {
		if ref, ok := stage.Stakeholders_reference[stakeholder]; !ok {
			stakeholders_newInstances = append(stakeholders_newInstances, stakeholder)
			newInstancesSlice = append(newInstancesSlice, stakeholder.GongMarshallIdentifier(stage))
			if stage.Stakeholders_referenceOrder == nil {
				stage.Stakeholders_referenceOrder = make(map[*Stakeholder]uint)
			}
			stage.Stakeholders_referenceOrder[stakeholder] = stage.Stakeholder_stagedOrder[stakeholder]
			newInstancesReverseSlice = append(newInstancesReverseSlice, stakeholder.GongMarshallUnstaging(stage))
			// delete(stage.Stakeholders_referenceOrder, stakeholder)
			fieldInitializers, pointersInitializations := stakeholder.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Stakeholder_stagedOrder[ref] = stage.Stakeholder_stagedOrder[stakeholder]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := stakeholder.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, stakeholder)
			// delete(stage.Stakeholder_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if stakeholder.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", stakeholder.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Stakeholders_reference {
		instance := stage.Stakeholders_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Stakeholders[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			stakeholders_deletedInstances = append(stakeholders_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(stakeholders_newInstances)
	lenDeletedInstances += len(stakeholders_deletedInstances)
	var stakeholdercompositionshapes_newInstances []*StakeholderCompositionShape
	var stakeholdercompositionshapes_deletedInstances []*StakeholderCompositionShape

	// parse all staged instances and check if they have a reference
	for stakeholdercompositionshape := range stage.StakeholderCompositionShapes {
		if ref, ok := stage.StakeholderCompositionShapes_reference[stakeholdercompositionshape]; !ok {
			stakeholdercompositionshapes_newInstances = append(stakeholdercompositionshapes_newInstances, stakeholdercompositionshape)
			newInstancesSlice = append(newInstancesSlice, stakeholdercompositionshape.GongMarshallIdentifier(stage))
			if stage.StakeholderCompositionShapes_referenceOrder == nil {
				stage.StakeholderCompositionShapes_referenceOrder = make(map[*StakeholderCompositionShape]uint)
			}
			stage.StakeholderCompositionShapes_referenceOrder[stakeholdercompositionshape] = stage.StakeholderCompositionShape_stagedOrder[stakeholdercompositionshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, stakeholdercompositionshape.GongMarshallUnstaging(stage))
			// delete(stage.StakeholderCompositionShapes_referenceOrder, stakeholdercompositionshape)
			fieldInitializers, pointersInitializations := stakeholdercompositionshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StakeholderCompositionShape_stagedOrder[ref] = stage.StakeholderCompositionShape_stagedOrder[stakeholdercompositionshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := stakeholdercompositionshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, stakeholdercompositionshape)
			// delete(stage.StakeholderCompositionShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if stakeholdercompositionshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", stakeholdercompositionshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.StakeholderCompositionShapes_reference {
		instance := stage.StakeholderCompositionShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.StakeholderCompositionShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			stakeholdercompositionshapes_deletedInstances = append(stakeholdercompositionshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(stakeholdercompositionshapes_newInstances)
	lenDeletedInstances += len(stakeholdercompositionshapes_deletedInstances)
	var stakeholderconcernshapes_newInstances []*StakeholderConcernShape
	var stakeholderconcernshapes_deletedInstances []*StakeholderConcernShape

	// parse all staged instances and check if they have a reference
	for stakeholderconcernshape := range stage.StakeholderConcernShapes {
		if ref, ok := stage.StakeholderConcernShapes_reference[stakeholderconcernshape]; !ok {
			stakeholderconcernshapes_newInstances = append(stakeholderconcernshapes_newInstances, stakeholderconcernshape)
			newInstancesSlice = append(newInstancesSlice, stakeholderconcernshape.GongMarshallIdentifier(stage))
			if stage.StakeholderConcernShapes_referenceOrder == nil {
				stage.StakeholderConcernShapes_referenceOrder = make(map[*StakeholderConcernShape]uint)
			}
			stage.StakeholderConcernShapes_referenceOrder[stakeholderconcernshape] = stage.StakeholderConcernShape_stagedOrder[stakeholderconcernshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, stakeholderconcernshape.GongMarshallUnstaging(stage))
			// delete(stage.StakeholderConcernShapes_referenceOrder, stakeholderconcernshape)
			fieldInitializers, pointersInitializations := stakeholderconcernshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StakeholderConcernShape_stagedOrder[ref] = stage.StakeholderConcernShape_stagedOrder[stakeholderconcernshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := stakeholderconcernshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, stakeholderconcernshape)
			// delete(stage.StakeholderConcernShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if stakeholderconcernshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", stakeholderconcernshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.StakeholderConcernShapes_reference {
		instance := stage.StakeholderConcernShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.StakeholderConcernShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			stakeholderconcernshapes_deletedInstances = append(stakeholderconcernshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(stakeholderconcernshapes_newInstances)
	lenDeletedInstances += len(stakeholderconcernshapes_deletedInstances)
	var stakeholdershapes_newInstances []*StakeholderShape
	var stakeholdershapes_deletedInstances []*StakeholderShape

	// parse all staged instances and check if they have a reference
	for stakeholdershape := range stage.StakeholderShapes {
		if ref, ok := stage.StakeholderShapes_reference[stakeholdershape]; !ok {
			stakeholdershapes_newInstances = append(stakeholdershapes_newInstances, stakeholdershape)
			newInstancesSlice = append(newInstancesSlice, stakeholdershape.GongMarshallIdentifier(stage))
			if stage.StakeholderShapes_referenceOrder == nil {
				stage.StakeholderShapes_referenceOrder = make(map[*StakeholderShape]uint)
			}
			stage.StakeholderShapes_referenceOrder[stakeholdershape] = stage.StakeholderShape_stagedOrder[stakeholdershape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, stakeholdershape.GongMarshallUnstaging(stage))
			// delete(stage.StakeholderShapes_referenceOrder, stakeholdershape)
			fieldInitializers, pointersInitializations := stakeholdershape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StakeholderShape_stagedOrder[ref] = stage.StakeholderShape_stagedOrder[stakeholdershape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := stakeholdershape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, stakeholdershape)
			// delete(stage.StakeholderShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if stakeholdershape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", stakeholdershape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.StakeholderShapes_reference {
		instance := stage.StakeholderShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.StakeholderShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			stakeholdershapes_deletedInstances = append(stakeholdershapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(stakeholdershapes_newInstances)
	lenDeletedInstances += len(stakeholdershapes_deletedInstances)
	var supportlevels_newInstances []*SupportLevel
	var supportlevels_deletedInstances []*SupportLevel

	// parse all staged instances and check if they have a reference
	for supportlevel := range stage.SupportLevels {
		if ref, ok := stage.SupportLevels_reference[supportlevel]; !ok {
			supportlevels_newInstances = append(supportlevels_newInstances, supportlevel)
			newInstancesSlice = append(newInstancesSlice, supportlevel.GongMarshallIdentifier(stage))
			if stage.SupportLevels_referenceOrder == nil {
				stage.SupportLevels_referenceOrder = make(map[*SupportLevel]uint)
			}
			stage.SupportLevels_referenceOrder[supportlevel] = stage.SupportLevel_stagedOrder[supportlevel]
			newInstancesReverseSlice = append(newInstancesReverseSlice, supportlevel.GongMarshallUnstaging(stage))
			// delete(stage.SupportLevels_referenceOrder, supportlevel)
			fieldInitializers, pointersInitializations := supportlevel.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SupportLevel_stagedOrder[ref] = stage.SupportLevel_stagedOrder[supportlevel]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := supportlevel.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, supportlevel)
			// delete(stage.SupportLevel_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if supportlevel.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", supportlevel.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.SupportLevels_reference {
		instance := stage.SupportLevels_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.SupportLevels[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			supportlevels_deletedInstances = append(supportlevels_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(supportlevels_newInstances)
	lenDeletedInstances += len(supportlevels_deletedInstances)
	var tools_newInstances []*Tool
	var tools_deletedInstances []*Tool

	// parse all staged instances and check if they have a reference
	for tool := range stage.Tools {
		if ref, ok := stage.Tools_reference[tool]; !ok {
			tools_newInstances = append(tools_newInstances, tool)
			newInstancesSlice = append(newInstancesSlice, tool.GongMarshallIdentifier(stage))
			if stage.Tools_referenceOrder == nil {
				stage.Tools_referenceOrder = make(map[*Tool]uint)
			}
			stage.Tools_referenceOrder[tool] = stage.Tool_stagedOrder[tool]
			newInstancesReverseSlice = append(newInstancesReverseSlice, tool.GongMarshallUnstaging(stage))
			// delete(stage.Tools_referenceOrder, tool)
			fieldInitializers, pointersInitializations := tool.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Tool_stagedOrder[ref] = stage.Tool_stagedOrder[tool]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := tool.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, tool)
			// delete(stage.Tool_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if tool.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", tool.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Tools_reference {
		instance := stage.Tools_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Tools[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			tools_deletedInstances = append(tools_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(tools_newInstances)
	lenDeletedInstances += len(tools_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.forwardCommits = append(stage.forwardCommits, forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.backwardCommits = append(stage.backwardCommits, backwardCommit)
		stage.modified = true
	} else {
		stage.modified = false
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	stage.AnalysisNeeds_reference = make(map[*AnalysisNeed]*AnalysisNeed)
	stage.AnalysisNeeds_referenceOrder = make(map[*AnalysisNeed]uint) // diff Unstage needs the reference order
	stage.AnalysisNeeds_instance = make(map[*AnalysisNeed]*AnalysisNeed)
	for instance := range stage.AnalysisNeeds {
		_copy := instance.GongCopy().(*AnalysisNeed)
		stage.AnalysisNeeds_reference[instance] = _copy
		stage.AnalysisNeeds_instance[_copy] = instance
		stage.AnalysisNeeds_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Concepts_reference = make(map[*Concept]*Concept)
	stage.Concepts_referenceOrder = make(map[*Concept]uint) // diff Unstage needs the reference order
	stage.Concepts_instance = make(map[*Concept]*Concept)
	for instance := range stage.Concepts {
		_copy := instance.GongCopy().(*Concept)
		stage.Concepts_reference[instance] = _copy
		stage.Concepts_instance[_copy] = instance
		stage.Concepts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ConceptShapes_reference = make(map[*ConceptShape]*ConceptShape)
	stage.ConceptShapes_referenceOrder = make(map[*ConceptShape]uint) // diff Unstage needs the reference order
	stage.ConceptShapes_instance = make(map[*ConceptShape]*ConceptShape)
	for instance := range stage.ConceptShapes {
		_copy := instance.GongCopy().(*ConceptShape)
		stage.ConceptShapes_reference[instance] = _copy
		stage.ConceptShapes_instance[_copy] = instance
		stage.ConceptShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Concerns_reference = make(map[*Concern]*Concern)
	stage.Concerns_referenceOrder = make(map[*Concern]uint) // diff Unstage needs the reference order
	stage.Concerns_instance = make(map[*Concern]*Concern)
	for instance := range stage.Concerns {
		_copy := instance.GongCopy().(*Concern)
		stage.Concerns_reference[instance] = _copy
		stage.Concerns_instance[_copy] = instance
		stage.Concerns_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ConcernCompositionShapes_reference = make(map[*ConcernCompositionShape]*ConcernCompositionShape)
	stage.ConcernCompositionShapes_referenceOrder = make(map[*ConcernCompositionShape]uint) // diff Unstage needs the reference order
	stage.ConcernCompositionShapes_instance = make(map[*ConcernCompositionShape]*ConcernCompositionShape)
	for instance := range stage.ConcernCompositionShapes {
		_copy := instance.GongCopy().(*ConcernCompositionShape)
		stage.ConcernCompositionShapes_reference[instance] = _copy
		stage.ConcernCompositionShapes_instance[_copy] = instance
		stage.ConcernCompositionShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ConcernInputShapes_reference = make(map[*ConcernInputShape]*ConcernInputShape)
	stage.ConcernInputShapes_referenceOrder = make(map[*ConcernInputShape]uint) // diff Unstage needs the reference order
	stage.ConcernInputShapes_instance = make(map[*ConcernInputShape]*ConcernInputShape)
	for instance := range stage.ConcernInputShapes {
		_copy := instance.GongCopy().(*ConcernInputShape)
		stage.ConcernInputShapes_reference[instance] = _copy
		stage.ConcernInputShapes_instance[_copy] = instance
		stage.ConcernInputShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ConcernOutputShapes_reference = make(map[*ConcernOutputShape]*ConcernOutputShape)
	stage.ConcernOutputShapes_referenceOrder = make(map[*ConcernOutputShape]uint) // diff Unstage needs the reference order
	stage.ConcernOutputShapes_instance = make(map[*ConcernOutputShape]*ConcernOutputShape)
	for instance := range stage.ConcernOutputShapes {
		_copy := instance.GongCopy().(*ConcernOutputShape)
		stage.ConcernOutputShapes_reference[instance] = _copy
		stage.ConcernOutputShapes_instance[_copy] = instance
		stage.ConcernOutputShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ConcernShapes_reference = make(map[*ConcernShape]*ConcernShape)
	stage.ConcernShapes_referenceOrder = make(map[*ConcernShape]uint) // diff Unstage needs the reference order
	stage.ConcernShapes_instance = make(map[*ConcernShape]*ConcernShape)
	for instance := range stage.ConcernShapes {
		_copy := instance.GongCopy().(*ConcernShape)
		stage.ConcernShapes_reference[instance] = _copy
		stage.ConcernShapes_instance[_copy] = instance
		stage.ConcernShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Deliverables_reference = make(map[*Deliverable]*Deliverable)
	stage.Deliverables_referenceOrder = make(map[*Deliverable]uint) // diff Unstage needs the reference order
	stage.Deliverables_instance = make(map[*Deliverable]*Deliverable)
	for instance := range stage.Deliverables {
		_copy := instance.GongCopy().(*Deliverable)
		stage.Deliverables_reference[instance] = _copy
		stage.Deliverables_instance[_copy] = instance
		stage.Deliverables_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.DeliverableCompositionShapes_reference = make(map[*DeliverableCompositionShape]*DeliverableCompositionShape)
	stage.DeliverableCompositionShapes_referenceOrder = make(map[*DeliverableCompositionShape]uint) // diff Unstage needs the reference order
	stage.DeliverableCompositionShapes_instance = make(map[*DeliverableCompositionShape]*DeliverableCompositionShape)
	for instance := range stage.DeliverableCompositionShapes {
		_copy := instance.GongCopy().(*DeliverableCompositionShape)
		stage.DeliverableCompositionShapes_reference[instance] = _copy
		stage.DeliverableCompositionShapes_instance[_copy] = instance
		stage.DeliverableCompositionShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.DeliverableConceptShapes_reference = make(map[*DeliverableConceptShape]*DeliverableConceptShape)
	stage.DeliverableConceptShapes_referenceOrder = make(map[*DeliverableConceptShape]uint) // diff Unstage needs the reference order
	stage.DeliverableConceptShapes_instance = make(map[*DeliverableConceptShape]*DeliverableConceptShape)
	for instance := range stage.DeliverableConceptShapes {
		_copy := instance.GongCopy().(*DeliverableConceptShape)
		stage.DeliverableConceptShapes_reference[instance] = _copy
		stage.DeliverableConceptShapes_instance[_copy] = instance
		stage.DeliverableConceptShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.DeliverableShapes_reference = make(map[*DeliverableShape]*DeliverableShape)
	stage.DeliverableShapes_referenceOrder = make(map[*DeliverableShape]uint) // diff Unstage needs the reference order
	stage.DeliverableShapes_instance = make(map[*DeliverableShape]*DeliverableShape)
	for instance := range stage.DeliverableShapes {
		_copy := instance.GongCopy().(*DeliverableShape)
		stage.DeliverableShapes_reference[instance] = _copy
		stage.DeliverableShapes_instance[_copy] = instance
		stage.DeliverableShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Diagrams_reference = make(map[*Diagram]*Diagram)
	stage.Diagrams_referenceOrder = make(map[*Diagram]uint) // diff Unstage needs the reference order
	stage.Diagrams_instance = make(map[*Diagram]*Diagram)
	for instance := range stage.Diagrams {
		_copy := instance.GongCopy().(*Diagram)
		stage.Diagrams_reference[instance] = _copy
		stage.Diagrams_instance[_copy] = instance
		stage.Diagrams_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Librarys_reference = make(map[*Library]*Library)
	stage.Librarys_referenceOrder = make(map[*Library]uint) // diff Unstage needs the reference order
	stage.Librarys_instance = make(map[*Library]*Library)
	for instance := range stage.Librarys {
		_copy := instance.GongCopy().(*Library)
		stage.Librarys_reference[instance] = _copy
		stage.Librarys_instance[_copy] = instance
		stage.Librarys_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Notes_reference = make(map[*Note]*Note)
	stage.Notes_referenceOrder = make(map[*Note]uint) // diff Unstage needs the reference order
	stage.Notes_instance = make(map[*Note]*Note)
	for instance := range stage.Notes {
		_copy := instance.GongCopy().(*Note)
		stage.Notes_reference[instance] = _copy
		stage.Notes_instance[_copy] = instance
		stage.Notes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.NoteDeliverableShapes_reference = make(map[*NoteDeliverableShape]*NoteDeliverableShape)
	stage.NoteDeliverableShapes_referenceOrder = make(map[*NoteDeliverableShape]uint) // diff Unstage needs the reference order
	stage.NoteDeliverableShapes_instance = make(map[*NoteDeliverableShape]*NoteDeliverableShape)
	for instance := range stage.NoteDeliverableShapes {
		_copy := instance.GongCopy().(*NoteDeliverableShape)
		stage.NoteDeliverableShapes_reference[instance] = _copy
		stage.NoteDeliverableShapes_instance[_copy] = instance
		stage.NoteDeliverableShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.NoteShapes_reference = make(map[*NoteShape]*NoteShape)
	stage.NoteShapes_referenceOrder = make(map[*NoteShape]uint) // diff Unstage needs the reference order
	stage.NoteShapes_instance = make(map[*NoteShape]*NoteShape)
	for instance := range stage.NoteShapes {
		_copy := instance.GongCopy().(*NoteShape)
		stage.NoteShapes_reference[instance] = _copy
		stage.NoteShapes_instance[_copy] = instance
		stage.NoteShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.NoteStakeholderShapes_reference = make(map[*NoteStakeholderShape]*NoteStakeholderShape)
	stage.NoteStakeholderShapes_referenceOrder = make(map[*NoteStakeholderShape]uint) // diff Unstage needs the reference order
	stage.NoteStakeholderShapes_instance = make(map[*NoteStakeholderShape]*NoteStakeholderShape)
	for instance := range stage.NoteStakeholderShapes {
		_copy := instance.GongCopy().(*NoteStakeholderShape)
		stage.NoteStakeholderShapes_reference[instance] = _copy
		stage.NoteStakeholderShapes_instance[_copy] = instance
		stage.NoteStakeholderShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.NoteTaskShapes_reference = make(map[*NoteTaskShape]*NoteTaskShape)
	stage.NoteTaskShapes_referenceOrder = make(map[*NoteTaskShape]uint) // diff Unstage needs the reference order
	stage.NoteTaskShapes_instance = make(map[*NoteTaskShape]*NoteTaskShape)
	for instance := range stage.NoteTaskShapes {
		_copy := instance.GongCopy().(*NoteTaskShape)
		stage.NoteTaskShapes_reference[instance] = _copy
		stage.NoteTaskShapes_instance[_copy] = instance
		stage.NoteTaskShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Requirements_reference = make(map[*Requirement]*Requirement)
	stage.Requirements_referenceOrder = make(map[*Requirement]uint) // diff Unstage needs the reference order
	stage.Requirements_instance = make(map[*Requirement]*Requirement)
	for instance := range stage.Requirements {
		_copy := instance.GongCopy().(*Requirement)
		stage.Requirements_reference[instance] = _copy
		stage.Requirements_instance[_copy] = instance
		stage.Requirements_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.RequirementShapes_reference = make(map[*RequirementShape]*RequirementShape)
	stage.RequirementShapes_referenceOrder = make(map[*RequirementShape]uint) // diff Unstage needs the reference order
	stage.RequirementShapes_instance = make(map[*RequirementShape]*RequirementShape)
	for instance := range stage.RequirementShapes {
		_copy := instance.GongCopy().(*RequirementShape)
		stage.RequirementShapes_reference[instance] = _copy
		stage.RequirementShapes_instance[_copy] = instance
		stage.RequirementShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Stakeholders_reference = make(map[*Stakeholder]*Stakeholder)
	stage.Stakeholders_referenceOrder = make(map[*Stakeholder]uint) // diff Unstage needs the reference order
	stage.Stakeholders_instance = make(map[*Stakeholder]*Stakeholder)
	for instance := range stage.Stakeholders {
		_copy := instance.GongCopy().(*Stakeholder)
		stage.Stakeholders_reference[instance] = _copy
		stage.Stakeholders_instance[_copy] = instance
		stage.Stakeholders_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StakeholderCompositionShapes_reference = make(map[*StakeholderCompositionShape]*StakeholderCompositionShape)
	stage.StakeholderCompositionShapes_referenceOrder = make(map[*StakeholderCompositionShape]uint) // diff Unstage needs the reference order
	stage.StakeholderCompositionShapes_instance = make(map[*StakeholderCompositionShape]*StakeholderCompositionShape)
	for instance := range stage.StakeholderCompositionShapes {
		_copy := instance.GongCopy().(*StakeholderCompositionShape)
		stage.StakeholderCompositionShapes_reference[instance] = _copy
		stage.StakeholderCompositionShapes_instance[_copy] = instance
		stage.StakeholderCompositionShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StakeholderConcernShapes_reference = make(map[*StakeholderConcernShape]*StakeholderConcernShape)
	stage.StakeholderConcernShapes_referenceOrder = make(map[*StakeholderConcernShape]uint) // diff Unstage needs the reference order
	stage.StakeholderConcernShapes_instance = make(map[*StakeholderConcernShape]*StakeholderConcernShape)
	for instance := range stage.StakeholderConcernShapes {
		_copy := instance.GongCopy().(*StakeholderConcernShape)
		stage.StakeholderConcernShapes_reference[instance] = _copy
		stage.StakeholderConcernShapes_instance[_copy] = instance
		stage.StakeholderConcernShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StakeholderShapes_reference = make(map[*StakeholderShape]*StakeholderShape)
	stage.StakeholderShapes_referenceOrder = make(map[*StakeholderShape]uint) // diff Unstage needs the reference order
	stage.StakeholderShapes_instance = make(map[*StakeholderShape]*StakeholderShape)
	for instance := range stage.StakeholderShapes {
		_copy := instance.GongCopy().(*StakeholderShape)
		stage.StakeholderShapes_reference[instance] = _copy
		stage.StakeholderShapes_instance[_copy] = instance
		stage.StakeholderShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SupportLevels_reference = make(map[*SupportLevel]*SupportLevel)
	stage.SupportLevels_referenceOrder = make(map[*SupportLevel]uint) // diff Unstage needs the reference order
	stage.SupportLevels_instance = make(map[*SupportLevel]*SupportLevel)
	for instance := range stage.SupportLevels {
		_copy := instance.GongCopy().(*SupportLevel)
		stage.SupportLevels_reference[instance] = _copy
		stage.SupportLevels_instance[_copy] = instance
		stage.SupportLevels_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Tools_reference = make(map[*Tool]*Tool)
	stage.Tools_referenceOrder = make(map[*Tool]uint) // diff Unstage needs the reference order
	stage.Tools_instance = make(map[*Tool]*Tool)
	for instance := range stage.Tools {
		_copy := instance.GongCopy().(*Tool)
		stage.Tools_reference[instance] = _copy
		stage.Tools_instance[_copy] = instance
		stage.Tools_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.AnalysisNeeds {
		reference := stage.AnalysisNeeds_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Concepts {
		reference := stage.Concepts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ConceptShapes {
		reference := stage.ConceptShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Concerns {
		reference := stage.Concerns_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ConcernCompositionShapes {
		reference := stage.ConcernCompositionShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ConcernInputShapes {
		reference := stage.ConcernInputShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ConcernOutputShapes {
		reference := stage.ConcernOutputShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ConcernShapes {
		reference := stage.ConcernShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Deliverables {
		reference := stage.Deliverables_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.DeliverableCompositionShapes {
		reference := stage.DeliverableCompositionShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.DeliverableConceptShapes {
		reference := stage.DeliverableConceptShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.DeliverableShapes {
		reference := stage.DeliverableShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Diagrams {
		reference := stage.Diagrams_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Librarys {
		reference := stage.Librarys_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Notes {
		reference := stage.Notes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.NoteDeliverableShapes {
		reference := stage.NoteDeliverableShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.NoteShapes {
		reference := stage.NoteShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.NoteStakeholderShapes {
		reference := stage.NoteStakeholderShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.NoteTaskShapes {
		reference := stage.NoteTaskShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Requirements {
		reference := stage.Requirements_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.RequirementShapes {
		reference := stage.RequirementShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Stakeholders {
		reference := stage.Stakeholders_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StakeholderCompositionShapes {
		reference := stage.StakeholderCompositionShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StakeholderConcernShapes {
		reference := stage.StakeholderConcernShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StakeholderShapes {
		reference := stage.StakeholderShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SupportLevels {
		reference := stage.SupportLevels_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Tools {
		reference := stage.Tools_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (analysisneed *AnalysisNeed) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.AnalysisNeed_stagedOrder[analysisneed]; ok {
		return order
	}
	if order, ok := stage.AnalysisNeeds_referenceOrder[analysisneed]; ok {
		return order
	} else {
		log.Printf("instance %p of type AnalysisNeed was not staged and does not have a reference order", analysisneed)
		return 0
	}
}

func (concept *Concept) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Concept_stagedOrder[concept]; ok {
		return order
	}
	if order, ok := stage.Concepts_referenceOrder[concept]; ok {
		return order
	} else {
		log.Printf("instance %p of type Concept was not staged and does not have a reference order", concept)
		return 0
	}
}

func (conceptshape *ConceptShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ConceptShape_stagedOrder[conceptshape]; ok {
		return order
	}
	if order, ok := stage.ConceptShapes_referenceOrder[conceptshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ConceptShape was not staged and does not have a reference order", conceptshape)
		return 0
	}
}

func (concern *Concern) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Concern_stagedOrder[concern]; ok {
		return order
	}
	if order, ok := stage.Concerns_referenceOrder[concern]; ok {
		return order
	} else {
		log.Printf("instance %p of type Concern was not staged and does not have a reference order", concern)
		return 0
	}
}

func (concerncompositionshape *ConcernCompositionShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ConcernCompositionShape_stagedOrder[concerncompositionshape]; ok {
		return order
	}
	if order, ok := stage.ConcernCompositionShapes_referenceOrder[concerncompositionshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ConcernCompositionShape was not staged and does not have a reference order", concerncompositionshape)
		return 0
	}
}

func (concerninputshape *ConcernInputShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ConcernInputShape_stagedOrder[concerninputshape]; ok {
		return order
	}
	if order, ok := stage.ConcernInputShapes_referenceOrder[concerninputshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ConcernInputShape was not staged and does not have a reference order", concerninputshape)
		return 0
	}
}

func (concernoutputshape *ConcernOutputShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ConcernOutputShape_stagedOrder[concernoutputshape]; ok {
		return order
	}
	if order, ok := stage.ConcernOutputShapes_referenceOrder[concernoutputshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ConcernOutputShape was not staged and does not have a reference order", concernoutputshape)
		return 0
	}
}

func (concernshape *ConcernShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ConcernShape_stagedOrder[concernshape]; ok {
		return order
	}
	if order, ok := stage.ConcernShapes_referenceOrder[concernshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ConcernShape was not staged and does not have a reference order", concernshape)
		return 0
	}
}

func (deliverable *Deliverable) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Deliverable_stagedOrder[deliverable]; ok {
		return order
	}
	if order, ok := stage.Deliverables_referenceOrder[deliverable]; ok {
		return order
	} else {
		log.Printf("instance %p of type Deliverable was not staged and does not have a reference order", deliverable)
		return 0
	}
}

func (deliverablecompositionshape *DeliverableCompositionShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DeliverableCompositionShape_stagedOrder[deliverablecompositionshape]; ok {
		return order
	}
	if order, ok := stage.DeliverableCompositionShapes_referenceOrder[deliverablecompositionshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type DeliverableCompositionShape was not staged and does not have a reference order", deliverablecompositionshape)
		return 0
	}
}

func (deliverableconceptshape *DeliverableConceptShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DeliverableConceptShape_stagedOrder[deliverableconceptshape]; ok {
		return order
	}
	if order, ok := stage.DeliverableConceptShapes_referenceOrder[deliverableconceptshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type DeliverableConceptShape was not staged and does not have a reference order", deliverableconceptshape)
		return 0
	}
}

func (deliverableshape *DeliverableShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DeliverableShape_stagedOrder[deliverableshape]; ok {
		return order
	}
	if order, ok := stage.DeliverableShapes_referenceOrder[deliverableshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type DeliverableShape was not staged and does not have a reference order", deliverableshape)
		return 0
	}
}

func (diagram *Diagram) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Diagram_stagedOrder[diagram]; ok {
		return order
	}
	if order, ok := stage.Diagrams_referenceOrder[diagram]; ok {
		return order
	} else {
		log.Printf("instance %p of type Diagram was not staged and does not have a reference order", diagram)
		return 0
	}
}

func (library *Library) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Library_stagedOrder[library]; ok {
		return order
	}
	if order, ok := stage.Librarys_referenceOrder[library]; ok {
		return order
	} else {
		log.Printf("instance %p of type Library was not staged and does not have a reference order", library)
		return 0
	}
}

func (note *Note) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Note_stagedOrder[note]; ok {
		return order
	}
	if order, ok := stage.Notes_referenceOrder[note]; ok {
		return order
	} else {
		log.Printf("instance %p of type Note was not staged and does not have a reference order", note)
		return 0
	}
}

func (notedeliverableshape *NoteDeliverableShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.NoteDeliverableShape_stagedOrder[notedeliverableshape]; ok {
		return order
	}
	if order, ok := stage.NoteDeliverableShapes_referenceOrder[notedeliverableshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type NoteDeliverableShape was not staged and does not have a reference order", notedeliverableshape)
		return 0
	}
}

func (noteshape *NoteShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.NoteShape_stagedOrder[noteshape]; ok {
		return order
	}
	if order, ok := stage.NoteShapes_referenceOrder[noteshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type NoteShape was not staged and does not have a reference order", noteshape)
		return 0
	}
}

func (notestakeholdershape *NoteStakeholderShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.NoteStakeholderShape_stagedOrder[notestakeholdershape]; ok {
		return order
	}
	if order, ok := stage.NoteStakeholderShapes_referenceOrder[notestakeholdershape]; ok {
		return order
	} else {
		log.Printf("instance %p of type NoteStakeholderShape was not staged and does not have a reference order", notestakeholdershape)
		return 0
	}
}

func (notetaskshape *NoteTaskShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.NoteTaskShape_stagedOrder[notetaskshape]; ok {
		return order
	}
	if order, ok := stage.NoteTaskShapes_referenceOrder[notetaskshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type NoteTaskShape was not staged and does not have a reference order", notetaskshape)
		return 0
	}
}

func (requirement *Requirement) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Requirement_stagedOrder[requirement]; ok {
		return order
	}
	if order, ok := stage.Requirements_referenceOrder[requirement]; ok {
		return order
	} else {
		log.Printf("instance %p of type Requirement was not staged and does not have a reference order", requirement)
		return 0
	}
}

func (requirementshape *RequirementShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.RequirementShape_stagedOrder[requirementshape]; ok {
		return order
	}
	if order, ok := stage.RequirementShapes_referenceOrder[requirementshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type RequirementShape was not staged and does not have a reference order", requirementshape)
		return 0
	}
}

func (stakeholder *Stakeholder) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Stakeholder_stagedOrder[stakeholder]; ok {
		return order
	}
	if order, ok := stage.Stakeholders_referenceOrder[stakeholder]; ok {
		return order
	} else {
		log.Printf("instance %p of type Stakeholder was not staged and does not have a reference order", stakeholder)
		return 0
	}
}

func (stakeholdercompositionshape *StakeholderCompositionShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StakeholderCompositionShape_stagedOrder[stakeholdercompositionshape]; ok {
		return order
	}
	if order, ok := stage.StakeholderCompositionShapes_referenceOrder[stakeholdercompositionshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type StakeholderCompositionShape was not staged and does not have a reference order", stakeholdercompositionshape)
		return 0
	}
}

func (stakeholderconcernshape *StakeholderConcernShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StakeholderConcernShape_stagedOrder[stakeholderconcernshape]; ok {
		return order
	}
	if order, ok := stage.StakeholderConcernShapes_referenceOrder[stakeholderconcernshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type StakeholderConcernShape was not staged and does not have a reference order", stakeholderconcernshape)
		return 0
	}
}

func (stakeholdershape *StakeholderShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StakeholderShape_stagedOrder[stakeholdershape]; ok {
		return order
	}
	if order, ok := stage.StakeholderShapes_referenceOrder[stakeholdershape]; ok {
		return order
	} else {
		log.Printf("instance %p of type StakeholderShape was not staged and does not have a reference order", stakeholdershape)
		return 0
	}
}

func (supportlevel *SupportLevel) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SupportLevel_stagedOrder[supportlevel]; ok {
		return order
	}
	if order, ok := stage.SupportLevels_referenceOrder[supportlevel]; ok {
		return order
	} else {
		log.Printf("instance %p of type SupportLevel was not staged and does not have a reference order", supportlevel)
		return 0
	}
}

func (tool *Tool) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Tool_stagedOrder[tool]; ok {
		return order
	}
	if order, ok := stage.Tools_referenceOrder[tool]; ok {
		return order
	} else {
		log.Printf("instance %p of type Tool was not staged and does not have a reference order", tool)
		return 0
	}
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (analysisneed *AnalysisNeed) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", analysisneed.GongGetGongstructName(), analysisneed.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (analysisneed *AnalysisNeed) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", analysisneed.GongGetGongstructName(), analysisneed.GongGetOrder(stage))
}

func (concept *Concept) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", concept.GongGetGongstructName(), concept.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (concept *Concept) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", concept.GongGetGongstructName(), concept.GongGetOrder(stage))
}

func (conceptshape *ConceptShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", conceptshape.GongGetGongstructName(), conceptshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (conceptshape *ConceptShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", conceptshape.GongGetGongstructName(), conceptshape.GongGetOrder(stage))
}

func (concern *Concern) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", concern.GongGetGongstructName(), concern.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (concern *Concern) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", concern.GongGetGongstructName(), concern.GongGetOrder(stage))
}

func (concerncompositionshape *ConcernCompositionShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", concerncompositionshape.GongGetGongstructName(), concerncompositionshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (concerncompositionshape *ConcernCompositionShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", concerncompositionshape.GongGetGongstructName(), concerncompositionshape.GongGetOrder(stage))
}

func (concerninputshape *ConcernInputShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", concerninputshape.GongGetGongstructName(), concerninputshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (concerninputshape *ConcernInputShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", concerninputshape.GongGetGongstructName(), concerninputshape.GongGetOrder(stage))
}

func (concernoutputshape *ConcernOutputShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", concernoutputshape.GongGetGongstructName(), concernoutputshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (concernoutputshape *ConcernOutputShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", concernoutputshape.GongGetGongstructName(), concernoutputshape.GongGetOrder(stage))
}

func (concernshape *ConcernShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", concernshape.GongGetGongstructName(), concernshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (concernshape *ConcernShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", concernshape.GongGetGongstructName(), concernshape.GongGetOrder(stage))
}

func (deliverable *Deliverable) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", deliverable.GongGetGongstructName(), deliverable.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (deliverable *Deliverable) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", deliverable.GongGetGongstructName(), deliverable.GongGetOrder(stage))
}

func (deliverablecompositionshape *DeliverableCompositionShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", deliverablecompositionshape.GongGetGongstructName(), deliverablecompositionshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (deliverablecompositionshape *DeliverableCompositionShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", deliverablecompositionshape.GongGetGongstructName(), deliverablecompositionshape.GongGetOrder(stage))
}

func (deliverableconceptshape *DeliverableConceptShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", deliverableconceptshape.GongGetGongstructName(), deliverableconceptshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (deliverableconceptshape *DeliverableConceptShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", deliverableconceptshape.GongGetGongstructName(), deliverableconceptshape.GongGetOrder(stage))
}

func (deliverableshape *DeliverableShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", deliverableshape.GongGetGongstructName(), deliverableshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (deliverableshape *DeliverableShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", deliverableshape.GongGetGongstructName(), deliverableshape.GongGetOrder(stage))
}

func (diagram *Diagram) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagram.GongGetGongstructName(), diagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (diagram *Diagram) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagram.GongGetGongstructName(), diagram.GongGetOrder(stage))
}

func (library *Library) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", library.GongGetGongstructName(), library.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (library *Library) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", library.GongGetGongstructName(), library.GongGetOrder(stage))
}

func (note *Note) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", note.GongGetGongstructName(), note.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (note *Note) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", note.GongGetGongstructName(), note.GongGetOrder(stage))
}

func (notedeliverableshape *NoteDeliverableShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", notedeliverableshape.GongGetGongstructName(), notedeliverableshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notedeliverableshape *NoteDeliverableShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", notedeliverableshape.GongGetGongstructName(), notedeliverableshape.GongGetOrder(stage))
}

func (noteshape *NoteShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", noteshape.GongGetGongstructName(), noteshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (noteshape *NoteShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", noteshape.GongGetGongstructName(), noteshape.GongGetOrder(stage))
}

func (notestakeholdershape *NoteStakeholderShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", notestakeholdershape.GongGetGongstructName(), notestakeholdershape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notestakeholdershape *NoteStakeholderShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", notestakeholdershape.GongGetGongstructName(), notestakeholdershape.GongGetOrder(stage))
}

func (notetaskshape *NoteTaskShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", notetaskshape.GongGetGongstructName(), notetaskshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notetaskshape *NoteTaskShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", notetaskshape.GongGetGongstructName(), notetaskshape.GongGetOrder(stage))
}

func (requirement *Requirement) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", requirement.GongGetGongstructName(), requirement.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (requirement *Requirement) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", requirement.GongGetGongstructName(), requirement.GongGetOrder(stage))
}

func (requirementshape *RequirementShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", requirementshape.GongGetGongstructName(), requirementshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (requirementshape *RequirementShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", requirementshape.GongGetGongstructName(), requirementshape.GongGetOrder(stage))
}

func (stakeholder *Stakeholder) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stakeholder.GongGetGongstructName(), stakeholder.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stakeholder *Stakeholder) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stakeholder.GongGetGongstructName(), stakeholder.GongGetOrder(stage))
}

func (stakeholdercompositionshape *StakeholderCompositionShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stakeholdercompositionshape.GongGetGongstructName(), stakeholdercompositionshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stakeholdercompositionshape *StakeholderCompositionShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stakeholdercompositionshape.GongGetGongstructName(), stakeholdercompositionshape.GongGetOrder(stage))
}

func (stakeholderconcernshape *StakeholderConcernShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stakeholderconcernshape.GongGetGongstructName(), stakeholderconcernshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stakeholderconcernshape *StakeholderConcernShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stakeholderconcernshape.GongGetGongstructName(), stakeholderconcernshape.GongGetOrder(stage))
}

func (stakeholdershape *StakeholderShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stakeholdershape.GongGetGongstructName(), stakeholdershape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stakeholdershape *StakeholderShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stakeholdershape.GongGetGongstructName(), stakeholdershape.GongGetOrder(stage))
}

func (supportlevel *SupportLevel) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", supportlevel.GongGetGongstructName(), supportlevel.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (supportlevel *SupportLevel) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", supportlevel.GongGetGongstructName(), supportlevel.GongGetOrder(stage))
}

func (tool *Tool) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tool.GongGetGongstructName(), tool.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tool *Tool) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tool.GongGetGongstructName(), tool.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (analysisneed *AnalysisNeed) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", analysisneed.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AnalysisNeed")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(analysisneed.Name))
	return
}

func (concept *Concept) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", concept.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Concept")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(concept.Name))
	return
}

func (conceptshape *ConceptShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", conceptshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ConceptShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(conceptshape.Name))
	return
}

func (concern *Concern) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", concern.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Concern")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(concern.Name))
	return
}

func (concerncompositionshape *ConcernCompositionShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", concerncompositionshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ConcernCompositionShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(concerncompositionshape.Name))
	return
}

func (concerninputshape *ConcernInputShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", concerninputshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ConcernInputShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(concerninputshape.Name))
	return
}

func (concernoutputshape *ConcernOutputShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", concernoutputshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ConcernOutputShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(concernoutputshape.Name))
	return
}

func (concernshape *ConcernShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", concernshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ConcernShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(concernshape.Name))
	return
}

func (deliverable *Deliverable) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", deliverable.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Deliverable")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(deliverable.Name))
	return
}

func (deliverablecompositionshape *DeliverableCompositionShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", deliverablecompositionshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DeliverableCompositionShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(deliverablecompositionshape.Name))
	return
}

func (deliverableconceptshape *DeliverableConceptShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", deliverableconceptshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DeliverableConceptShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(deliverableconceptshape.Name))
	return
}

func (deliverableshape *DeliverableShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", deliverableshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DeliverableShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(deliverableshape.Name))
	return
}

func (diagram *Diagram) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagram.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Diagram")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.Name))
	return
}

func (library *Library) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", library.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Library")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(library.Name))
	return
}

func (note *Note) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", note.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Note")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(note.Name))
	return
}

func (notedeliverableshape *NoteDeliverableShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", notedeliverableshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteDeliverableShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(notedeliverableshape.Name))
	return
}

func (noteshape *NoteShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(noteshape.Name))
	return
}

func (notestakeholdershape *NoteStakeholderShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", notestakeholdershape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteStakeholderShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(notestakeholdershape.Name))
	return
}

func (notetaskshape *NoteTaskShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteTaskShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(notetaskshape.Name))
	return
}

func (requirement *Requirement) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", requirement.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Requirement")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(requirement.Name))
	return
}

func (requirementshape *RequirementShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", requirementshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RequirementShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(requirementshape.Name))
	return
}

func (stakeholder *Stakeholder) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stakeholder.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Stakeholder")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(stakeholder.Name))
	return
}

func (stakeholdercompositionshape *StakeholderCompositionShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stakeholdercompositionshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StakeholderCompositionShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(stakeholdercompositionshape.Name))
	return
}

func (stakeholderconcernshape *StakeholderConcernShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stakeholderconcernshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StakeholderConcernShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(stakeholderconcernshape.Name))
	return
}

func (stakeholdershape *StakeholderShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stakeholdershape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StakeholderShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(stakeholdershape.Name))
	return
}

func (supportlevel *SupportLevel) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", supportlevel.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SupportLevel")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(supportlevel.Name))
	return
}

func (tool *Tool) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tool.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Tool")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(tool.Name))
	return
}

// insertion point for unstaging
func (analysisneed *AnalysisNeed) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", analysisneed.GongGetReferenceIdentifier(stage))
	return
}

func (concept *Concept) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", concept.GongGetReferenceIdentifier(stage))
	return
}

func (conceptshape *ConceptShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", conceptshape.GongGetReferenceIdentifier(stage))
	return
}

func (concern *Concern) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", concern.GongGetReferenceIdentifier(stage))
	return
}

func (concerncompositionshape *ConcernCompositionShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", concerncompositionshape.GongGetReferenceIdentifier(stage))
	return
}

func (concerninputshape *ConcernInputShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", concerninputshape.GongGetReferenceIdentifier(stage))
	return
}

func (concernoutputshape *ConcernOutputShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", concernoutputshape.GongGetReferenceIdentifier(stage))
	return
}

func (concernshape *ConcernShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", concernshape.GongGetReferenceIdentifier(stage))
	return
}

func (deliverable *Deliverable) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", deliverable.GongGetReferenceIdentifier(stage))
	return
}

func (deliverablecompositionshape *DeliverableCompositionShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", deliverablecompositionshape.GongGetReferenceIdentifier(stage))
	return
}

func (deliverableconceptshape *DeliverableConceptShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", deliverableconceptshape.GongGetReferenceIdentifier(stage))
	return
}

func (deliverableshape *DeliverableShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", deliverableshape.GongGetReferenceIdentifier(stage))
	return
}

func (diagram *Diagram) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagram.GongGetReferenceIdentifier(stage))
	return
}

func (library *Library) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", library.GongGetReferenceIdentifier(stage))
	return
}

func (note *Note) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", note.GongGetReferenceIdentifier(stage))
	return
}

func (notedeliverableshape *NoteDeliverableShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", notedeliverableshape.GongGetReferenceIdentifier(stage))
	return
}

func (noteshape *NoteShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteshape.GongGetReferenceIdentifier(stage))
	return
}

func (notestakeholdershape *NoteStakeholderShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", notestakeholdershape.GongGetReferenceIdentifier(stage))
	return
}

func (notetaskshape *NoteTaskShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", notetaskshape.GongGetReferenceIdentifier(stage))
	return
}

func (requirement *Requirement) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", requirement.GongGetReferenceIdentifier(stage))
	return
}

func (requirementshape *RequirementShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", requirementshape.GongGetReferenceIdentifier(stage))
	return
}

func (stakeholder *Stakeholder) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stakeholder.GongGetReferenceIdentifier(stage))
	return
}

func (stakeholdercompositionshape *StakeholderCompositionShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stakeholdercompositionshape.GongGetReferenceIdentifier(stage))
	return
}

func (stakeholderconcernshape *StakeholderConcernShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stakeholderconcernshape.GongGetReferenceIdentifier(stage))
	return
}

func (stakeholdershape *StakeholderShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stakeholdershape.GongGetReferenceIdentifier(stage))
	return
}

func (supportlevel *SupportLevel) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", supportlevel.GongGetReferenceIdentifier(stage))
	return
}

func (tool *Tool) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tool.GongGetReferenceIdentifier(stage))
	return
}

func IntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += IntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
	// 1. Create a deterministic hash from the inputs using SHA-256
	h := sha256.New()

	// Write the string to the hash
	h.Write([]byte(seedStr))

	// Write the integer to the hash (using BigEndian to ensure consistency across architectures)
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, seedInt)
	h.Write(intBytes)

	// 2. Extract the first 16 bytes from our resulting hash
	hashBytes := h.Sum(nil)
	uuid := make([]byte, 16)
	copy(uuid, hashBytes[:16])

	// 3. Set the Version to 4 (0100 in binary)
	// We take the 7th byte, clear the top 4 bits with & 0x0f, and set the top bits to 0100 with | 0x40
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// 4. Set the Variant to RFC4122 (10 in binary)
	// We take the 9th byte, clear the top 2 bits with & 0x3f, and set the top bits to 10 with | 0x80
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// 5. Format and return the byte array as a standard UUID string
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

// end of template
