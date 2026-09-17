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
	stage.ConcernCompositionShape_ControlPointShapes_reverseMap = make(map[*ControlPointShape]*ConcernCompositionShape)
	for concerncompositionshape := range stage.ConcernCompositionShapes {
		_ = concerncompositionshape
		for _, _controlpointshape := range concerncompositionshape.ControlPointShapes {
			stage.ConcernCompositionShape_ControlPointShapes_reverseMap[_controlpointshape] = concerncompositionshape
		}
	}

	// Compute reverse map for named struct ConcernInputShape
	// insertion point per field
	stage.ConcernInputShape_ControlPointShapes_reverseMap = make(map[*ControlPointShape]*ConcernInputShape)
	for concerninputshape := range stage.ConcernInputShapes {
		_ = concerninputshape
		for _, _controlpointshape := range concerninputshape.ControlPointShapes {
			stage.ConcernInputShape_ControlPointShapes_reverseMap[_controlpointshape] = concerninputshape
		}
	}

	// Compute reverse map for named struct ConcernOutputShape
	// insertion point per field
	stage.ConcernOutputShape_ControlPointShapes_reverseMap = make(map[*ControlPointShape]*ConcernOutputShape)
	for concernoutputshape := range stage.ConcernOutputShapes {
		_ = concernoutputshape
		for _, _controlpointshape := range concernoutputshape.ControlPointShapes {
			stage.ConcernOutputShape_ControlPointShapes_reverseMap[_controlpointshape] = concernoutputshape
		}
	}

	// Compute reverse map for named struct ConcernShape
	// insertion point per field

	// Compute reverse map for named struct ControlPointShape
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
	stage.DeliverableCompositionShape_ControlPointShapes_reverseMap = make(map[*ControlPointShape]*DeliverableCompositionShape)
	for deliverablecompositionshape := range stage.DeliverableCompositionShapes {
		_ = deliverablecompositionshape
		for _, _controlpointshape := range deliverablecompositionshape.ControlPointShapes {
			stage.DeliverableCompositionShape_ControlPointShapes_reverseMap[_controlpointshape] = deliverablecompositionshape
		}
	}

	// Compute reverse map for named struct DeliverableConceptShape
	// insertion point per field
	stage.DeliverableConceptShape_ControlPointShapes_reverseMap = make(map[*ControlPointShape]*DeliverableConceptShape)
	for deliverableconceptshape := range stage.DeliverableConceptShapes {
		_ = deliverableconceptshape
		for _, _controlpointshape := range deliverableconceptshape.ControlPointShapes {
			stage.DeliverableConceptShape_ControlPointShapes_reverseMap[_controlpointshape] = deliverableconceptshape
		}
	}

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
	stage.Diagram_Diagram_Shapes_reverseMap = make(map[*DiagramShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _diagramshape := range diagram.Diagram_Shapes {
			stage.Diagram_Diagram_Shapes_reverseMap[_diagramshape] = diagram
		}
	}
	stage.Diagram_DiagramsWhoseNodeIsExpanded_reverseMap = make(map[*Diagram]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _diagram := range diagram.DiagramsWhoseNodeIsExpanded {
			stage.Diagram_DiagramsWhoseNodeIsExpanded_reverseMap[_diagram] = diagram
		}
	}

	// Compute reverse map for named struct DiagramShape
	// insertion point per field

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
	stage.NoteDeliverableShape_ControlPointShapes_reverseMap = make(map[*ControlPointShape]*NoteDeliverableShape)
	for notedeliverableshape := range stage.NoteDeliverableShapes {
		_ = notedeliverableshape
		for _, _controlpointshape := range notedeliverableshape.ControlPointShapes {
			stage.NoteDeliverableShape_ControlPointShapes_reverseMap[_controlpointshape] = notedeliverableshape
		}
	}

	// Compute reverse map for named struct NoteShape
	// insertion point per field

	// Compute reverse map for named struct NoteStakeholderShape
	// insertion point per field
	stage.NoteStakeholderShape_ControlPointShapes_reverseMap = make(map[*ControlPointShape]*NoteStakeholderShape)
	for notestakeholdershape := range stage.NoteStakeholderShapes {
		_ = notestakeholdershape
		for _, _controlpointshape := range notestakeholdershape.ControlPointShapes {
			stage.NoteStakeholderShape_ControlPointShapes_reverseMap[_controlpointshape] = notestakeholdershape
		}
	}

	// Compute reverse map for named struct NoteTaskShape
	// insertion point per field
	stage.NoteTaskShape_ControlPointShapes_reverseMap = make(map[*ControlPointShape]*NoteTaskShape)
	for notetaskshape := range stage.NoteTaskShapes {
		_ = notetaskshape
		for _, _controlpointshape := range notetaskshape.ControlPointShapes {
			stage.NoteTaskShape_ControlPointShapes_reverseMap[_controlpointshape] = notetaskshape
		}
	}

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
	stage.StakeholderCompositionShape_ControlPointShapes_reverseMap = make(map[*ControlPointShape]*StakeholderCompositionShape)
	for stakeholdercompositionshape := range stage.StakeholderCompositionShapes {
		_ = stakeholdercompositionshape
		for _, _controlpointshape := range stakeholdercompositionshape.ControlPointShapes {
			stage.StakeholderCompositionShape_ControlPointShapes_reverseMap[_controlpointshape] = stakeholdercompositionshape
		}
	}

	// Compute reverse map for named struct StakeholderConcernShape
	// insertion point per field
	stage.StakeholderConcernShape_ControlPointShapes_reverseMap = make(map[*ControlPointShape]*StakeholderConcernShape)
	for stakeholderconcernshape := range stage.StakeholderConcernShapes {
		_ = stakeholderconcernshape
		for _, _controlpointshape := range stakeholderconcernshape.ControlPointShapes {
			stage.StakeholderConcernShape_ControlPointShapes_reverseMap[_controlpointshape] = stakeholderconcernshape
		}
	}

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

	for instance := range stage.ControlPointShapes {
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

	for instance := range stage.DiagramShapes {
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
	analysisneed.GongCopyBasicFields(newInstance)
	return newInstance
}

func (concept *Concept) GongCopy() GongstructIF {
	newInstance := new(Concept)
	concept.GongCopyBasicFields(newInstance)
	return newInstance
}

func (conceptshape *ConceptShape) GongCopy() GongstructIF {
	newInstance := new(ConceptShape)
	conceptshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (concern *Concern) GongCopy() GongstructIF {
	newInstance := new(Concern)
	concern.GongCopyBasicFields(newInstance)
	return newInstance
}

func (concerncompositionshape *ConcernCompositionShape) GongCopy() GongstructIF {
	newInstance := new(ConcernCompositionShape)
	concerncompositionshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (concerninputshape *ConcernInputShape) GongCopy() GongstructIF {
	newInstance := new(ConcernInputShape)
	concerninputshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (concernoutputshape *ConcernOutputShape) GongCopy() GongstructIF {
	newInstance := new(ConcernOutputShape)
	concernoutputshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (concernshape *ConcernShape) GongCopy() GongstructIF {
	newInstance := new(ConcernShape)
	concernshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (controlpointshape *ControlPointShape) GongCopy() GongstructIF {
	newInstance := new(ControlPointShape)
	controlpointshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (deliverable *Deliverable) GongCopy() GongstructIF {
	newInstance := new(Deliverable)
	deliverable.GongCopyBasicFields(newInstance)
	return newInstance
}

func (deliverablecompositionshape *DeliverableCompositionShape) GongCopy() GongstructIF {
	newInstance := new(DeliverableCompositionShape)
	deliverablecompositionshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (deliverableconceptshape *DeliverableConceptShape) GongCopy() GongstructIF {
	newInstance := new(DeliverableConceptShape)
	deliverableconceptshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (deliverableshape *DeliverableShape) GongCopy() GongstructIF {
	newInstance := new(DeliverableShape)
	deliverableshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (diagram *Diagram) GongCopy() GongstructIF {
	newInstance := new(Diagram)
	diagram.GongCopyBasicFields(newInstance)
	return newInstance
}

func (diagramshape *DiagramShape) GongCopy() GongstructIF {
	newInstance := new(DiagramShape)
	diagramshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (library *Library) GongCopy() GongstructIF {
	newInstance := new(Library)
	library.GongCopyBasicFields(newInstance)
	return newInstance
}

func (note *Note) GongCopy() GongstructIF {
	newInstance := new(Note)
	note.GongCopyBasicFields(newInstance)
	return newInstance
}

func (notedeliverableshape *NoteDeliverableShape) GongCopy() GongstructIF {
	newInstance := new(NoteDeliverableShape)
	notedeliverableshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (noteshape *NoteShape) GongCopy() GongstructIF {
	newInstance := new(NoteShape)
	noteshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (notestakeholdershape *NoteStakeholderShape) GongCopy() GongstructIF {
	newInstance := new(NoteStakeholderShape)
	notestakeholdershape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (notetaskshape *NoteTaskShape) GongCopy() GongstructIF {
	newInstance := new(NoteTaskShape)
	notetaskshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (requirement *Requirement) GongCopy() GongstructIF {
	newInstance := new(Requirement)
	requirement.GongCopyBasicFields(newInstance)
	return newInstance
}

func (requirementshape *RequirementShape) GongCopy() GongstructIF {
	newInstance := new(RequirementShape)
	requirementshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stakeholder *Stakeholder) GongCopy() GongstructIF {
	newInstance := new(Stakeholder)
	stakeholder.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stakeholdercompositionshape *StakeholderCompositionShape) GongCopy() GongstructIF {
	newInstance := new(StakeholderCompositionShape)
	stakeholdercompositionshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stakeholderconcernshape *StakeholderConcernShape) GongCopy() GongstructIF {
	newInstance := new(StakeholderConcernShape)
	stakeholderconcernshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stakeholdershape *StakeholderShape) GongCopy() GongstructIF {
	newInstance := new(StakeholderShape)
	stakeholdershape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (supportlevel *SupportLevel) GongCopy() GongstructIF {
	newInstance := new(SupportLevel)
	supportlevel.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tool *Tool) GongCopy() GongstructIF {
	newInstance := new(Tool)
	tool.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (analysisneed *AnalysisNeed) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(analysisneed).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(analysisneed), uint64(stage.GetOrder(analysisneed)))
	return
}

func (concept *Concept) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(concept).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(concept), uint64(stage.GetOrder(concept)))
	return
}

func (conceptshape *ConceptShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(conceptshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(conceptshape), uint64(stage.GetOrder(conceptshape)))
	return
}

func (concern *Concern) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(concern).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(concern), uint64(stage.GetOrder(concern)))
	return
}

func (concerncompositionshape *ConcernCompositionShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(concerncompositionshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(concerncompositionshape), uint64(stage.GetOrder(concerncompositionshape)))
	return
}

func (concerninputshape *ConcernInputShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(concerninputshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(concerninputshape), uint64(stage.GetOrder(concerninputshape)))
	return
}

func (concernoutputshape *ConcernOutputShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(concernoutputshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(concernoutputshape), uint64(stage.GetOrder(concernoutputshape)))
	return
}

func (concernshape *ConcernShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(concernshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(concernshape), uint64(stage.GetOrder(concernshape)))
	return
}

func (controlpointshape *ControlPointShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(controlpointshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(controlpointshape), uint64(stage.GetOrder(controlpointshape)))
	return
}

func (deliverable *Deliverable) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(deliverable).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(deliverable), uint64(stage.GetOrder(deliverable)))
	return
}

func (deliverablecompositionshape *DeliverableCompositionShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(deliverablecompositionshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(deliverablecompositionshape), uint64(stage.GetOrder(deliverablecompositionshape)))
	return
}

func (deliverableconceptshape *DeliverableConceptShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(deliverableconceptshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(deliverableconceptshape), uint64(stage.GetOrder(deliverableconceptshape)))
	return
}

func (deliverableshape *DeliverableShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(deliverableshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(deliverableshape), uint64(stage.GetOrder(deliverableshape)))
	return
}

func (diagram *Diagram) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(diagram).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(diagram), uint64(stage.GetOrder(diagram)))
	return
}

func (diagramshape *DiagramShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(diagramshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(diagramshape), uint64(stage.GetOrder(diagramshape)))
	return
}

func (library *Library) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(library).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(library), uint64(stage.GetOrder(library)))
	return
}

func (note *Note) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(note).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(note), uint64(stage.GetOrder(note)))
	return
}

func (notedeliverableshape *NoteDeliverableShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(notedeliverableshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(notedeliverableshape), uint64(stage.GetOrder(notedeliverableshape)))
	return
}

func (noteshape *NoteShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(noteshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(noteshape), uint64(stage.GetOrder(noteshape)))
	return
}

func (notestakeholdershape *NoteStakeholderShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(notestakeholdershape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(notestakeholdershape), uint64(stage.GetOrder(notestakeholdershape)))
	return
}

func (notetaskshape *NoteTaskShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(notetaskshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(notetaskshape), uint64(stage.GetOrder(notetaskshape)))
	return
}

func (requirement *Requirement) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(requirement).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(requirement), uint64(stage.GetOrder(requirement)))
	return
}

func (requirementshape *RequirementShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(requirementshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(requirementshape), uint64(stage.GetOrder(requirementshape)))
	return
}

func (stakeholder *Stakeholder) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stakeholder).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(stakeholder), uint64(stage.GetOrder(stakeholder)))
	return
}

func (stakeholdercompositionshape *StakeholderCompositionShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stakeholdercompositionshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(stakeholdercompositionshape), uint64(stage.GetOrder(stakeholdercompositionshape)))
	return
}

func (stakeholderconcernshape *StakeholderConcernShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stakeholderconcernshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(stakeholderconcernshape), uint64(stage.GetOrder(stakeholderconcernshape)))
	return
}

func (stakeholdershape *StakeholderShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stakeholdershape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(stakeholdershape), uint64(stage.GetOrder(stakeholdershape)))
	return
}

func (supportlevel *SupportLevel) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(supportlevel).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(supportlevel), uint64(stage.GetOrder(supportlevel)))
	return
}

func (tool *Tool) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(tool).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(tool), uint64(stage.GetOrder(tool)))
	return
}


type GongstructDiffable[T any] interface {
	PointerToGongstruct
	GongMarshallIdentifier(stage *Stage) string
	GongMarshallUnstaging(stage *Stage) string
	GongMarshallAllFields(stage *Stage) (string, string)
	GongReconstructPointersFromInstances(stage *Stage)
	GongDiff(stage *Stage, other T) []string
}

func computeCommitsForType[T GongstructDiffable[T]](
	stage *Stage,
	stagedInstances map[T]struct{},
	stagedOrder map[T]uint,
	referenceInstances map[T]T,
	referenceOrder *map[T]uint,
	instancesMap map[T]T,
	newInstancesSlice *[]string,
	fieldsEditSlice *[]string,
	deletedInstancesSlice *[]string,
	newInstancesReverseSlice *[]string,
	fieldsEditReverseSlice *[]string,
	deletedInstancesReverseSlice *[]string,
	lenNewInstances *int,
	lenDeletedInstances *int,
	lenModifiedInstances *int,
) {
	var newInstances []T
	var deletedInstances []T

	// parse all staged instances and check if they have a reference
	for instance := range stagedInstances {
		if ref, ok := referenceInstances[instance]; !ok {
			newInstances = append(newInstances, instance)
			*newInstancesSlice = append(*newInstancesSlice, instance.GongMarshallIdentifier(stage))
			if *referenceOrder == nil {
				*referenceOrder = make(map[T]uint)
			}
			(*referenceOrder)[instance] = stagedOrder[instance]
			*newInstancesReverseSlice = append(*newInstancesReverseSlice, instance.GongMarshallUnstaging(stage))
			fieldInitializers, pointersInitializations := instance.GongMarshallAllFields(stage)
			*fieldsEditSlice = append(*fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stagedOrder[ref] = stagedOrder[instance]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := instance.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, instance)
			if len(diffs) > 0 {
				var fieldsEdit string
				if instance.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", instance.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				*fieldsEditSlice = append(*fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, reverseDiff)
				}
				*lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range referenceInstances {
		instance := instancesMap[ref] // get the instance corresponding to the reference
		if _, ok := stagedInstances[instance]; !ok { // if the instance is not staged anymore, it means it has been unstaged
			deletedInstances = append(deletedInstances, ref)
			*deletedInstancesSlice = append(*deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			*deletedInstancesReverseSlice = append(*deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	*lenNewInstances += len(newInstances)
	*lenDeletedInstances += len(deletedInstances)
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
	computeCommitsForType(
		stage,
		stage.AnalysisNeeds,
		stage.AnalysisNeed_stagedOrder,
		stage.AnalysisNeeds_reference,
		&stage.AnalysisNeeds_referenceOrder,
		stage.AnalysisNeeds_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Concepts,
		stage.Concept_stagedOrder,
		stage.Concepts_reference,
		&stage.Concepts_referenceOrder,
		stage.Concepts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ConceptShapes,
		stage.ConceptShape_stagedOrder,
		stage.ConceptShapes_reference,
		&stage.ConceptShapes_referenceOrder,
		stage.ConceptShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Concerns,
		stage.Concern_stagedOrder,
		stage.Concerns_reference,
		&stage.Concerns_referenceOrder,
		stage.Concerns_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ConcernCompositionShapes,
		stage.ConcernCompositionShape_stagedOrder,
		stage.ConcernCompositionShapes_reference,
		&stage.ConcernCompositionShapes_referenceOrder,
		stage.ConcernCompositionShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ConcernInputShapes,
		stage.ConcernInputShape_stagedOrder,
		stage.ConcernInputShapes_reference,
		&stage.ConcernInputShapes_referenceOrder,
		stage.ConcernInputShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ConcernOutputShapes,
		stage.ConcernOutputShape_stagedOrder,
		stage.ConcernOutputShapes_reference,
		&stage.ConcernOutputShapes_referenceOrder,
		stage.ConcernOutputShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ConcernShapes,
		stage.ConcernShape_stagedOrder,
		stage.ConcernShapes_reference,
		&stage.ConcernShapes_referenceOrder,
		stage.ConcernShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ControlPointShapes,
		stage.ControlPointShape_stagedOrder,
		stage.ControlPointShapes_reference,
		&stage.ControlPointShapes_referenceOrder,
		stage.ControlPointShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Deliverables,
		stage.Deliverable_stagedOrder,
		stage.Deliverables_reference,
		&stage.Deliverables_referenceOrder,
		stage.Deliverables_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DeliverableCompositionShapes,
		stage.DeliverableCompositionShape_stagedOrder,
		stage.DeliverableCompositionShapes_reference,
		&stage.DeliverableCompositionShapes_referenceOrder,
		stage.DeliverableCompositionShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DeliverableConceptShapes,
		stage.DeliverableConceptShape_stagedOrder,
		stage.DeliverableConceptShapes_reference,
		&stage.DeliverableConceptShapes_referenceOrder,
		stage.DeliverableConceptShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DeliverableShapes,
		stage.DeliverableShape_stagedOrder,
		stage.DeliverableShapes_reference,
		&stage.DeliverableShapes_referenceOrder,
		stage.DeliverableShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Diagrams,
		stage.Diagram_stagedOrder,
		stage.Diagrams_reference,
		&stage.Diagrams_referenceOrder,
		stage.Diagrams_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DiagramShapes,
		stage.DiagramShape_stagedOrder,
		stage.DiagramShapes_reference,
		&stage.DiagramShapes_referenceOrder,
		stage.DiagramShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Librarys,
		stage.Library_stagedOrder,
		stage.Librarys_reference,
		&stage.Librarys_referenceOrder,
		stage.Librarys_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Notes,
		stage.Note_stagedOrder,
		stage.Notes_reference,
		&stage.Notes_referenceOrder,
		stage.Notes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.NoteDeliverableShapes,
		stage.NoteDeliverableShape_stagedOrder,
		stage.NoteDeliverableShapes_reference,
		&stage.NoteDeliverableShapes_referenceOrder,
		stage.NoteDeliverableShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.NoteShapes,
		stage.NoteShape_stagedOrder,
		stage.NoteShapes_reference,
		&stage.NoteShapes_referenceOrder,
		stage.NoteShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.NoteStakeholderShapes,
		stage.NoteStakeholderShape_stagedOrder,
		stage.NoteStakeholderShapes_reference,
		&stage.NoteStakeholderShapes_referenceOrder,
		stage.NoteStakeholderShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.NoteTaskShapes,
		stage.NoteTaskShape_stagedOrder,
		stage.NoteTaskShapes_reference,
		&stage.NoteTaskShapes_referenceOrder,
		stage.NoteTaskShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Requirements,
		stage.Requirement_stagedOrder,
		stage.Requirements_reference,
		&stage.Requirements_referenceOrder,
		stage.Requirements_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.RequirementShapes,
		stage.RequirementShape_stagedOrder,
		stage.RequirementShapes_reference,
		&stage.RequirementShapes_referenceOrder,
		stage.RequirementShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Stakeholders,
		stage.Stakeholder_stagedOrder,
		stage.Stakeholders_reference,
		&stage.Stakeholders_referenceOrder,
		stage.Stakeholders_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StakeholderCompositionShapes,
		stage.StakeholderCompositionShape_stagedOrder,
		stage.StakeholderCompositionShapes_reference,
		&stage.StakeholderCompositionShapes_referenceOrder,
		stage.StakeholderCompositionShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StakeholderConcernShapes,
		stage.StakeholderConcernShape_stagedOrder,
		stage.StakeholderConcernShapes_reference,
		&stage.StakeholderConcernShapes_referenceOrder,
		stage.StakeholderConcernShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StakeholderShapes,
		stage.StakeholderShape_stagedOrder,
		stage.StakeholderShapes_reference,
		&stage.StakeholderShapes_referenceOrder,
		stage.StakeholderShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SupportLevels,
		stage.SupportLevel_stagedOrder,
		stage.SupportLevels_reference,
		&stage.SupportLevels_referenceOrder,
		stage.SupportLevels_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Tools,
		stage.Tool_stagedOrder,
		stage.Tools_reference,
		&stage.Tools_referenceOrder,
		stage.Tools_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)

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

	stage.ControlPointShapes_reference = make(map[*ControlPointShape]*ControlPointShape)
	stage.ControlPointShapes_referenceOrder = make(map[*ControlPointShape]uint) // diff Unstage needs the reference order
	stage.ControlPointShapes_instance = make(map[*ControlPointShape]*ControlPointShape)
	for instance := range stage.ControlPointShapes {
		_copy := instance.GongCopy().(*ControlPointShape)
		stage.ControlPointShapes_reference[instance] = _copy
		stage.ControlPointShapes_instance[_copy] = instance
		stage.ControlPointShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.DiagramShapes_reference = make(map[*DiagramShape]*DiagramShape)
	stage.DiagramShapes_referenceOrder = make(map[*DiagramShape]uint) // diff Unstage needs the reference order
	stage.DiagramShapes_instance = make(map[*DiagramShape]*DiagramShape)
	for instance := range stage.DiagramShapes {
		_copy := instance.GongCopy().(*DiagramShape)
		stage.DiagramShapes_reference[instance] = _copy
		stage.DiagramShapes_instance[_copy] = instance
		stage.DiagramShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	for instance := range stage.ControlPointShapes {
		reference := stage.ControlPointShapes_reference[instance]
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

	for instance := range stage.DiagramShapes {
		reference := stage.DiagramShapes_reference[instance]
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

func (controlpointshape *ControlPointShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ControlPointShape_stagedOrder[controlpointshape]; ok {
		return order
	}
	if order, ok := stage.ControlPointShapes_referenceOrder[controlpointshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ControlPointShape was not staged and does not have a reference order", controlpointshape)
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

func (diagramshape *DiagramShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DiagramShape_stagedOrder[diagramshape]; ok {
		return order
	}
	if order, ok := stage.DiagramShapes_referenceOrder[diagramshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type DiagramShape was not staged and does not have a reference order", diagramshape)
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

func (controlpointshape *ControlPointShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", controlpointshape.GongGetGongstructName(), controlpointshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (controlpointshape *ControlPointShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", controlpointshape.GongGetGongstructName(), controlpointshape.GongGetOrder(stage))
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

func (diagramshape *DiagramShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagramshape.GongGetGongstructName(), diagramshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (diagramshape *DiagramShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagramshape.GongGetGongstructName(), diagramshape.GongGetOrder(stage))
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
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(analysisneed.Name))
	return
}

func (concept *Concept) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", concept.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Concept")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(concept.Name))
	return
}

func (conceptshape *ConceptShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", conceptshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ConceptShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(conceptshape.Name))
	return
}

func (concern *Concern) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", concern.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Concern")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(concern.Name))
	return
}

func (concerncompositionshape *ConcernCompositionShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", concerncompositionshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ConcernCompositionShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(concerncompositionshape.Name))
	return
}

func (concerninputshape *ConcernInputShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", concerninputshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ConcernInputShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(concerninputshape.Name))
	return
}

func (concernoutputshape *ConcernOutputShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", concernoutputshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ConcernOutputShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(concernoutputshape.Name))
	return
}

func (concernshape *ConcernShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", concernshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ConcernShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(concernshape.Name))
	return
}

func (controlpointshape *ControlPointShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", controlpointshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ControlPointShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(controlpointshape.Name))
	return
}

func (deliverable *Deliverable) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", deliverable.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Deliverable")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(deliverable.Name))
	return
}

func (deliverablecompositionshape *DeliverableCompositionShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", deliverablecompositionshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DeliverableCompositionShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(deliverablecompositionshape.Name))
	return
}

func (deliverableconceptshape *DeliverableConceptShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", deliverableconceptshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DeliverableConceptShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(deliverableconceptshape.Name))
	return
}

func (deliverableshape *DeliverableShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", deliverableshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DeliverableShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(deliverableshape.Name))
	return
}

func (diagram *Diagram) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagram.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Diagram")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(diagram.Name))
	return
}

func (diagramshape *DiagramShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagramshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DiagramShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(diagramshape.Name))
	return
}

func (library *Library) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", library.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Library")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(library.Name))
	return
}

func (note *Note) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", note.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Note")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(note.Name))
	return
}

func (notedeliverableshape *NoteDeliverableShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", notedeliverableshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteDeliverableShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(notedeliverableshape.Name))
	return
}

func (noteshape *NoteShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(noteshape.Name))
	return
}

func (notestakeholdershape *NoteStakeholderShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", notestakeholdershape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteStakeholderShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(notestakeholdershape.Name))
	return
}

func (notetaskshape *NoteTaskShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteTaskShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(notetaskshape.Name))
	return
}

func (requirement *Requirement) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", requirement.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Requirement")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(requirement.Name))
	return
}

func (requirementshape *RequirementShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", requirementshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RequirementShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(requirementshape.Name))
	return
}

func (stakeholder *Stakeholder) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stakeholder.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Stakeholder")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(stakeholder.Name))
	return
}

func (stakeholdercompositionshape *StakeholderCompositionShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stakeholdercompositionshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StakeholderCompositionShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(stakeholdercompositionshape.Name))
	return
}

func (stakeholderconcernshape *StakeholderConcernShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stakeholderconcernshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StakeholderConcernShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(stakeholderconcernshape.Name))
	return
}

func (stakeholdershape *StakeholderShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stakeholdershape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StakeholderShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(stakeholdershape.Name))
	return
}

func (supportlevel *SupportLevel) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", supportlevel.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SupportLevel")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(supportlevel.Name))
	return
}

func (tool *Tool) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tool.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Tool")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(tool.Name))
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

func (controlpointshape *ControlPointShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", controlpointshape.GongGetReferenceIdentifier(stage))
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

func (diagramshape *DiagramShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagramshape.GongGetReferenceIdentifier(stage))
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

func GongIntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += GongIntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GongGenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GongGenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
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
