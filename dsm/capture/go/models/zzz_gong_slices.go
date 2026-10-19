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
	// Compute reverse map for named struct Concept
	// insertion point per field
	stage.Concept_Tools_reverseMap = make(map[*Tool]*Concept)
	for concept := range stage.Concepts {
		_ = concept
		for _, _tool := range concept.Tools {
			stage.Concept_Tools_reverseMap[_tool] = concept
		}
	}

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

	// Compute reverse map for named struct Deliverable
	// insertion point per field
	stage.Deliverable_SubDeliverables_reverseMap = make(map[*Deliverable]*Deliverable)
	for deliverable := range stage.Deliverables {
		_ = deliverable
		for _, _deliverable := range deliverable.SubDeliverables {
			stage.Deliverable_SubDeliverables_reverseMap[_deliverable] = deliverable
		}
	}
	stage.Deliverable_Views_reverseMap = make(map[*View]*Deliverable)
	for deliverable := range stage.Deliverables {
		_ = deliverable
		for _, _view := range deliverable.Views {
			stage.Deliverable_Views_reverseMap[_view] = deliverable
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

	// Compute reverse map for named struct DeliverableViewShape
	// insertion point per field
	stage.DeliverableViewShape_ControlPointShapes_reverseMap = make(map[*ControlPointShape]*DeliverableViewShape)
	for deliverableviewshape := range stage.DeliverableViewShapes {
		_ = deliverableviewshape
		for _, _controlpointshape := range deliverableviewshape.ControlPointShapes {
			stage.DeliverableViewShape_ControlPointShapes_reverseMap[_controlpointshape] = deliverableviewshape
		}
	}

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
	stage.Diagram_DeliverablesWhoseViewsNodeIsExpanded_reverseMap = make(map[*Deliverable]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _deliverable := range diagram.DeliverablesWhoseViewsNodeIsExpanded {
			stage.Diagram_DeliverablesWhoseViewsNodeIsExpanded_reverseMap[_deliverable] = diagram
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
	stage.Diagram_RequirementsWhoseConcernsNodeIsExpanded_reverseMap = make(map[*Requirement]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _requirement := range diagram.RequirementsWhoseConcernsNodeIsExpanded {
			stage.Diagram_RequirementsWhoseConcernsNodeIsExpanded_reverseMap[_requirement] = diagram
		}
	}
	stage.Diagram_RequirementsWhoseViewsNodeIsExpanded_reverseMap = make(map[*Requirement]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _requirement := range diagram.RequirementsWhoseViewsNodeIsExpanded {
			stage.Diagram_RequirementsWhoseViewsNodeIsExpanded_reverseMap[_requirement] = diagram
		}
	}
	stage.Diagram_RequirementsWhoseConceptsNodeIsExpanded_reverseMap = make(map[*Requirement]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _requirement := range diagram.RequirementsWhoseConceptsNodeIsExpanded {
			stage.Diagram_RequirementsWhoseConceptsNodeIsExpanded_reverseMap[_requirement] = diagram
		}
	}
	stage.Diagram_RequirementConcernShapes_reverseMap = make(map[*RequirementConcernShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _requirementconcernshape := range diagram.RequirementConcernShapes {
			stage.Diagram_RequirementConcernShapes_reverseMap[_requirementconcernshape] = diagram
		}
	}
	stage.Diagram_RequirementViewShapes_reverseMap = make(map[*RequirementViewShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _requirementviewshape := range diagram.RequirementViewShapes {
			stage.Diagram_RequirementViewShapes_reverseMap[_requirementviewshape] = diagram
		}
	}
	stage.Diagram_RequirementConceptShapes_reverseMap = make(map[*RequirementConceptShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _requirementconceptshape := range diagram.RequirementConceptShapes {
			stage.Diagram_RequirementConceptShapes_reverseMap[_requirementconceptshape] = diagram
		}
	}
	stage.Diagram_View_Shapes_reverseMap = make(map[*ViewShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _viewshape := range diagram.View_Shapes {
			stage.Diagram_View_Shapes_reverseMap[_viewshape] = diagram
		}
	}
	stage.Diagram_ViewsWhoseNodeIsExpanded_reverseMap = make(map[*View]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _view := range diagram.ViewsWhoseNodeIsExpanded {
			stage.Diagram_ViewsWhoseNodeIsExpanded_reverseMap[_view] = diagram
		}
	}
	stage.Diagram_ViewsWhoseDeliverablesNodeIsExpanded_reverseMap = make(map[*View]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _view := range diagram.ViewsWhoseDeliverablesNodeIsExpanded {
			stage.Diagram_ViewsWhoseDeliverablesNodeIsExpanded_reverseMap[_view] = diagram
		}
	}
	stage.Diagram_ViewsWhoseConceptsNodeIsExpanded_reverseMap = make(map[*View]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _view := range diagram.ViewsWhoseConceptsNodeIsExpanded {
			stage.Diagram_ViewsWhoseConceptsNodeIsExpanded_reverseMap[_view] = diagram
		}
	}
	stage.Diagram_ViewsWhoseRequirementsNodeIsExpanded_reverseMap = make(map[*View]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _view := range diagram.ViewsWhoseRequirementsNodeIsExpanded {
			stage.Diagram_ViewsWhoseRequirementsNodeIsExpanded_reverseMap[_view] = diagram
		}
	}
	stage.Diagram_DeliverableViewShapes_reverseMap = make(map[*DeliverableViewShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _deliverableviewshape := range diagram.DeliverableViewShapes {
			stage.Diagram_DeliverableViewShapes_reverseMap[_deliverableviewshape] = diagram
		}
	}
	stage.Diagram_ViewConceptShapes_reverseMap = make(map[*ViewConceptShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _viewconceptshape := range diagram.ViewConceptShapes {
			stage.Diagram_ViewConceptShapes_reverseMap[_viewconceptshape] = diagram
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
	stage.Diagram_ConceptsWhoseViewsNodeIsExpanded_reverseMap = make(map[*Concept]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _concept := range diagram.ConceptsWhoseViewsNodeIsExpanded {
			stage.Diagram_ConceptsWhoseViewsNodeIsExpanded_reverseMap[_concept] = diagram
		}
	}
	stage.Diagram_ConceptsWhoseDeliverablesNodeIsExpanded_reverseMap = make(map[*Concept]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _concept := range diagram.ConceptsWhoseDeliverablesNodeIsExpanded {
			stage.Diagram_ConceptsWhoseDeliverablesNodeIsExpanded_reverseMap[_concept] = diagram
		}
	}
	stage.Diagram_ConceptsWhoseRequirementsNodeIsExpanded_reverseMap = make(map[*Concept]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _concept := range diagram.ConceptsWhoseRequirementsNodeIsExpanded {
			stage.Diagram_ConceptsWhoseRequirementsNodeIsExpanded_reverseMap[_concept] = diagram
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
	stage.Library_RootViews_reverseMap = make(map[*View]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _view := range library.RootViews {
			stage.Library_RootViews_reverseMap[_view] = library
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
	stage.Requirement_Concerns_reverseMap = make(map[*Concern]*Requirement)
	for requirement := range stage.Requirements {
		_ = requirement
		for _, _concern := range requirement.Concerns {
			stage.Requirement_Concerns_reverseMap[_concern] = requirement
		}
	}
	stage.Requirement_Views_reverseMap = make(map[*View]*Requirement)
	for requirement := range stage.Requirements {
		_ = requirement
		for _, _view := range requirement.Views {
			stage.Requirement_Views_reverseMap[_view] = requirement
		}
	}
	stage.Requirement_Concepts_reverseMap = make(map[*Concept]*Requirement)
	for requirement := range stage.Requirements {
		_ = requirement
		for _, _concept := range requirement.Concepts {
			stage.Requirement_Concepts_reverseMap[_concept] = requirement
		}
	}

	// Compute reverse map for named struct RequirementConceptShape
	// insertion point per field
	stage.RequirementConceptShape_ControlPointShapes_reverseMap = make(map[*ControlPointShape]*RequirementConceptShape)
	for requirementconceptshape := range stage.RequirementConceptShapes {
		_ = requirementconceptshape
		for _, _controlpointshape := range requirementconceptshape.ControlPointShapes {
			stage.RequirementConceptShape_ControlPointShapes_reverseMap[_controlpointshape] = requirementconceptshape
		}
	}

	// Compute reverse map for named struct RequirementConcernShape
	// insertion point per field
	stage.RequirementConcernShape_ControlPointShapes_reverseMap = make(map[*ControlPointShape]*RequirementConcernShape)
	for requirementconcernshape := range stage.RequirementConcernShapes {
		_ = requirementconcernshape
		for _, _controlpointshape := range requirementconcernshape.ControlPointShapes {
			stage.RequirementConcernShape_ControlPointShapes_reverseMap[_controlpointshape] = requirementconcernshape
		}
	}

	// Compute reverse map for named struct RequirementViewShape
	// insertion point per field
	stage.RequirementViewShape_ControlPointShapes_reverseMap = make(map[*ControlPointShape]*RequirementViewShape)
	for requirementviewshape := range stage.RequirementViewShapes {
		_ = requirementviewshape
		for _, _controlpointshape := range requirementviewshape.ControlPointShapes {
			stage.RequirementViewShape_ControlPointShapes_reverseMap[_controlpointshape] = requirementviewshape
		}
	}

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

	// Compute reverse map for named struct View
	// insertion point per field
	stage.View_Concepts_reverseMap = make(map[*Concept]*View)
	for view := range stage.Views {
		_ = view
		for _, _concept := range view.Concepts {
			stage.View_Concepts_reverseMap[_concept] = view
		}
	}

	// Compute reverse map for named struct ViewConceptShape
	// insertion point per field
	stage.ViewConceptShape_ControlPointShapes_reverseMap = make(map[*ControlPointShape]*ViewConceptShape)
	for viewconceptshape := range stage.ViewConceptShapes {
		_ = viewconceptshape
		for _, _controlpointshape := range viewconceptshape.ControlPointShapes {
			stage.ViewConceptShape_ControlPointShapes_reverseMap[_controlpointshape] = viewconceptshape
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.AnalysisNeeds)

	res = __gong__appendInstances(res, stage.Concepts)

	res = __gong__appendInstances(res, stage.ConceptShapes)

	res = __gong__appendInstances(res, stage.Concerns)

	res = __gong__appendInstances(res, stage.ConcernCompositionShapes)

	res = __gong__appendInstances(res, stage.ConcernInputShapes)

	res = __gong__appendInstances(res, stage.ConcernOutputShapes)

	res = __gong__appendInstances(res, stage.ConcernShapes)

	res = __gong__appendInstances(res, stage.ControlPointShapes)

	res = __gong__appendInstances(res, stage.Deliverables)

	res = __gong__appendInstances(res, stage.DeliverableCompositionShapes)

	res = __gong__appendInstances(res, stage.DeliverableConceptShapes)

	res = __gong__appendInstances(res, stage.DeliverableShapes)

	res = __gong__appendInstances(res, stage.DeliverableViewShapes)

	res = __gong__appendInstances(res, stage.Diagrams)

	res = __gong__appendInstances(res, stage.DiagramShapes)

	res = __gong__appendInstances(res, stage.Librarys)

	res = __gong__appendInstances(res, stage.Notes)

	res = __gong__appendInstances(res, stage.NoteDeliverableShapes)

	res = __gong__appendInstances(res, stage.NoteShapes)

	res = __gong__appendInstances(res, stage.NoteStakeholderShapes)

	res = __gong__appendInstances(res, stage.NoteTaskShapes)

	res = __gong__appendInstances(res, stage.Requirements)

	res = __gong__appendInstances(res, stage.RequirementConceptShapes)

	res = __gong__appendInstances(res, stage.RequirementConcernShapes)

	res = __gong__appendInstances(res, stage.RequirementShapes)

	res = __gong__appendInstances(res, stage.RequirementViewShapes)

	res = __gong__appendInstances(res, stage.Stakeholders)

	res = __gong__appendInstances(res, stage.StakeholderCompositionShapes)

	res = __gong__appendInstances(res, stage.StakeholderConcernShapes)

	res = __gong__appendInstances(res, stage.StakeholderShapes)

	res = __gong__appendInstances(res, stage.SupportLevels)

	res = __gong__appendInstances(res, stage.Tools)

	res = __gong__appendInstances(res, stage.Views)

	res = __gong__appendInstances(res, stage.ViewConceptShapes)

	res = __gong__appendInstances(res, stage.ViewShapes)

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

func (deliverableviewshape *DeliverableViewShape) GongCopy() GongstructIF {
	newInstance := new(DeliverableViewShape)
	deliverableviewshape.GongCopyBasicFields(newInstance)
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

func (requirementconceptshape *RequirementConceptShape) GongCopy() GongstructIF {
	newInstance := new(RequirementConceptShape)
	requirementconceptshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (requirementconcernshape *RequirementConcernShape) GongCopy() GongstructIF {
	newInstance := new(RequirementConcernShape)
	requirementconcernshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (requirementshape *RequirementShape) GongCopy() GongstructIF {
	newInstance := new(RequirementShape)
	requirementshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (requirementviewshape *RequirementViewShape) GongCopy() GongstructIF {
	newInstance := new(RequirementViewShape)
	requirementviewshape.GongCopyBasicFields(newInstance)
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

func (view *View) GongCopy() GongstructIF {
	newInstance := new(View)
	view.GongCopyBasicFields(newInstance)
	return newInstance
}

func (viewconceptshape *ViewConceptShape) GongCopy() GongstructIF {
	newInstance := new(ViewConceptShape)
	viewconceptshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (viewshape *ViewShape) GongCopy() GongstructIF {
	newInstance := new(ViewShape)
	viewshape.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (analysisneed *AnalysisNeed) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, analysisneed)
}

func (concept *Concept) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, concept)
}

func (conceptshape *ConceptShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, conceptshape)
}

func (concern *Concern) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, concern)
}

func (concerncompositionshape *ConcernCompositionShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, concerncompositionshape)
}

func (concerninputshape *ConcernInputShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, concerninputshape)
}

func (concernoutputshape *ConcernOutputShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, concernoutputshape)
}

func (concernshape *ConcernShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, concernshape)
}

func (controlpointshape *ControlPointShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, controlpointshape)
}

func (deliverable *Deliverable) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, deliverable)
}

func (deliverablecompositionshape *DeliverableCompositionShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, deliverablecompositionshape)
}

func (deliverableconceptshape *DeliverableConceptShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, deliverableconceptshape)
}

func (deliverableshape *DeliverableShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, deliverableshape)
}

func (deliverableviewshape *DeliverableViewShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, deliverableviewshape)
}

func (diagram *Diagram) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, diagram)
}

func (diagramshape *DiagramShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, diagramshape)
}

func (library *Library) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, library)
}

func (note *Note) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, note)
}

func (notedeliverableshape *NoteDeliverableShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, notedeliverableshape)
}

func (noteshape *NoteShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, noteshape)
}

func (notestakeholdershape *NoteStakeholderShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, notestakeholdershape)
}

func (notetaskshape *NoteTaskShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, notetaskshape)
}

func (requirement *Requirement) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, requirement)
}

func (requirementconceptshape *RequirementConceptShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, requirementconceptshape)
}

func (requirementconcernshape *RequirementConcernShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, requirementconcernshape)
}

func (requirementshape *RequirementShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, requirementshape)
}

func (requirementviewshape *RequirementViewShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, requirementviewshape)
}

func (stakeholder *Stakeholder) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stakeholder)
}

func (stakeholdercompositionshape *StakeholderCompositionShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stakeholdercompositionshape)
}

func (stakeholderconcernshape *StakeholderConcernShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stakeholderconcernshape)
}

func (stakeholdershape *StakeholderShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stakeholdershape)
}

func (supportlevel *SupportLevel) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, supportlevel)
}

func (tool *Tool) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tool)
}

func (view *View) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, view)
}

func (viewconceptshape *ViewConceptShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, viewconceptshape)
}

func (viewshape *ViewShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, viewshape)
}


type GongstructDiffable[T any] interface {
	GongstructPtr
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
		stage.DeliverableViewShapes,
		stage.DeliverableViewShape_stagedOrder,
		stage.DeliverableViewShapes_reference,
		&stage.DeliverableViewShapes_referenceOrder,
		stage.DeliverableViewShapes_instance,
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
		stage.RequirementConceptShapes,
		stage.RequirementConceptShape_stagedOrder,
		stage.RequirementConceptShapes_reference,
		&stage.RequirementConceptShapes_referenceOrder,
		stage.RequirementConceptShapes_instance,
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
		stage.RequirementConcernShapes,
		stage.RequirementConcernShape_stagedOrder,
		stage.RequirementConcernShapes_reference,
		&stage.RequirementConcernShapes_referenceOrder,
		stage.RequirementConcernShapes_instance,
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
		stage.RequirementViewShapes,
		stage.RequirementViewShape_stagedOrder,
		stage.RequirementViewShapes_reference,
		&stage.RequirementViewShapes_referenceOrder,
		stage.RequirementViewShapes_instance,
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
	computeCommitsForType(
		stage,
		stage.Views,
		stage.View_stagedOrder,
		stage.Views_reference,
		&stage.Views_referenceOrder,
		stage.Views_instance,
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
		stage.ViewConceptShapes,
		stage.ViewConceptShape_stagedOrder,
		stage.ViewConceptShapes_reference,
		&stage.ViewConceptShapes_referenceOrder,
		stage.ViewConceptShapes_instance,
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
		stage.ViewShapes,
		stage.ViewShape_stagedOrder,
		stage.ViewShapes_reference,
		&stage.ViewShapes_referenceOrder,
		stage.ViewShapes_instance,
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
		stage.AppendForwardCommit(forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.AppendBackwardCommit(backwardCommit)
		stage.SetModified(true)
	} else {
		stage.SetModified(false)
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	__gong__computeReferencePass1(stage, stage.AnalysisNeeds, &stage.AnalysisNeeds_reference, &stage.AnalysisNeeds_referenceOrder, &stage.AnalysisNeeds_instance)

	__gong__computeReferencePass1(stage, stage.Concepts, &stage.Concepts_reference, &stage.Concepts_referenceOrder, &stage.Concepts_instance)

	__gong__computeReferencePass1(stage, stage.ConceptShapes, &stage.ConceptShapes_reference, &stage.ConceptShapes_referenceOrder, &stage.ConceptShapes_instance)

	__gong__computeReferencePass1(stage, stage.Concerns, &stage.Concerns_reference, &stage.Concerns_referenceOrder, &stage.Concerns_instance)

	__gong__computeReferencePass1(stage, stage.ConcernCompositionShapes, &stage.ConcernCompositionShapes_reference, &stage.ConcernCompositionShapes_referenceOrder, &stage.ConcernCompositionShapes_instance)

	__gong__computeReferencePass1(stage, stage.ConcernInputShapes, &stage.ConcernInputShapes_reference, &stage.ConcernInputShapes_referenceOrder, &stage.ConcernInputShapes_instance)

	__gong__computeReferencePass1(stage, stage.ConcernOutputShapes, &stage.ConcernOutputShapes_reference, &stage.ConcernOutputShapes_referenceOrder, &stage.ConcernOutputShapes_instance)

	__gong__computeReferencePass1(stage, stage.ConcernShapes, &stage.ConcernShapes_reference, &stage.ConcernShapes_referenceOrder, &stage.ConcernShapes_instance)

	__gong__computeReferencePass1(stage, stage.ControlPointShapes, &stage.ControlPointShapes_reference, &stage.ControlPointShapes_referenceOrder, &stage.ControlPointShapes_instance)

	__gong__computeReferencePass1(stage, stage.Deliverables, &stage.Deliverables_reference, &stage.Deliverables_referenceOrder, &stage.Deliverables_instance)

	__gong__computeReferencePass1(stage, stage.DeliverableCompositionShapes, &stage.DeliverableCompositionShapes_reference, &stage.DeliverableCompositionShapes_referenceOrder, &stage.DeliverableCompositionShapes_instance)

	__gong__computeReferencePass1(stage, stage.DeliverableConceptShapes, &stage.DeliverableConceptShapes_reference, &stage.DeliverableConceptShapes_referenceOrder, &stage.DeliverableConceptShapes_instance)

	__gong__computeReferencePass1(stage, stage.DeliverableShapes, &stage.DeliverableShapes_reference, &stage.DeliverableShapes_referenceOrder, &stage.DeliverableShapes_instance)

	__gong__computeReferencePass1(stage, stage.DeliverableViewShapes, &stage.DeliverableViewShapes_reference, &stage.DeliverableViewShapes_referenceOrder, &stage.DeliverableViewShapes_instance)

	__gong__computeReferencePass1(stage, stage.Diagrams, &stage.Diagrams_reference, &stage.Diagrams_referenceOrder, &stage.Diagrams_instance)

	__gong__computeReferencePass1(stage, stage.DiagramShapes, &stage.DiagramShapes_reference, &stage.DiagramShapes_referenceOrder, &stage.DiagramShapes_instance)

	__gong__computeReferencePass1(stage, stage.Librarys, &stage.Librarys_reference, &stage.Librarys_referenceOrder, &stage.Librarys_instance)

	__gong__computeReferencePass1(stage, stage.Notes, &stage.Notes_reference, &stage.Notes_referenceOrder, &stage.Notes_instance)

	__gong__computeReferencePass1(stage, stage.NoteDeliverableShapes, &stage.NoteDeliverableShapes_reference, &stage.NoteDeliverableShapes_referenceOrder, &stage.NoteDeliverableShapes_instance)

	__gong__computeReferencePass1(stage, stage.NoteShapes, &stage.NoteShapes_reference, &stage.NoteShapes_referenceOrder, &stage.NoteShapes_instance)

	__gong__computeReferencePass1(stage, stage.NoteStakeholderShapes, &stage.NoteStakeholderShapes_reference, &stage.NoteStakeholderShapes_referenceOrder, &stage.NoteStakeholderShapes_instance)

	__gong__computeReferencePass1(stage, stage.NoteTaskShapes, &stage.NoteTaskShapes_reference, &stage.NoteTaskShapes_referenceOrder, &stage.NoteTaskShapes_instance)

	__gong__computeReferencePass1(stage, stage.Requirements, &stage.Requirements_reference, &stage.Requirements_referenceOrder, &stage.Requirements_instance)

	__gong__computeReferencePass1(stage, stage.RequirementConceptShapes, &stage.RequirementConceptShapes_reference, &stage.RequirementConceptShapes_referenceOrder, &stage.RequirementConceptShapes_instance)

	__gong__computeReferencePass1(stage, stage.RequirementConcernShapes, &stage.RequirementConcernShapes_reference, &stage.RequirementConcernShapes_referenceOrder, &stage.RequirementConcernShapes_instance)

	__gong__computeReferencePass1(stage, stage.RequirementShapes, &stage.RequirementShapes_reference, &stage.RequirementShapes_referenceOrder, &stage.RequirementShapes_instance)

	__gong__computeReferencePass1(stage, stage.RequirementViewShapes, &stage.RequirementViewShapes_reference, &stage.RequirementViewShapes_referenceOrder, &stage.RequirementViewShapes_instance)

	__gong__computeReferencePass1(stage, stage.Stakeholders, &stage.Stakeholders_reference, &stage.Stakeholders_referenceOrder, &stage.Stakeholders_instance)

	__gong__computeReferencePass1(stage, stage.StakeholderCompositionShapes, &stage.StakeholderCompositionShapes_reference, &stage.StakeholderCompositionShapes_referenceOrder, &stage.StakeholderCompositionShapes_instance)

	__gong__computeReferencePass1(stage, stage.StakeholderConcernShapes, &stage.StakeholderConcernShapes_reference, &stage.StakeholderConcernShapes_referenceOrder, &stage.StakeholderConcernShapes_instance)

	__gong__computeReferencePass1(stage, stage.StakeholderShapes, &stage.StakeholderShapes_reference, &stage.StakeholderShapes_referenceOrder, &stage.StakeholderShapes_instance)

	__gong__computeReferencePass1(stage, stage.SupportLevels, &stage.SupportLevels_reference, &stage.SupportLevels_referenceOrder, &stage.SupportLevels_instance)

	__gong__computeReferencePass1(stage, stage.Tools, &stage.Tools_reference, &stage.Tools_referenceOrder, &stage.Tools_instance)

	__gong__computeReferencePass1(stage, stage.Views, &stage.Views_reference, &stage.Views_referenceOrder, &stage.Views_instance)

	__gong__computeReferencePass1(stage, stage.ViewConceptShapes, &stage.ViewConceptShapes_reference, &stage.ViewConceptShapes_referenceOrder, &stage.ViewConceptShapes_instance)

	__gong__computeReferencePass1(stage, stage.ViewShapes, &stage.ViewShapes_reference, &stage.ViewShapes_referenceOrder, &stage.ViewShapes_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.AnalysisNeeds, stage.AnalysisNeeds_reference, stage)

	__gong__computeReferencePass2(stage.Concepts, stage.Concepts_reference, stage)

	__gong__computeReferencePass2(stage.ConceptShapes, stage.ConceptShapes_reference, stage)

	__gong__computeReferencePass2(stage.Concerns, stage.Concerns_reference, stage)

	__gong__computeReferencePass2(stage.ConcernCompositionShapes, stage.ConcernCompositionShapes_reference, stage)

	__gong__computeReferencePass2(stage.ConcernInputShapes, stage.ConcernInputShapes_reference, stage)

	__gong__computeReferencePass2(stage.ConcernOutputShapes, stage.ConcernOutputShapes_reference, stage)

	__gong__computeReferencePass2(stage.ConcernShapes, stage.ConcernShapes_reference, stage)

	__gong__computeReferencePass2(stage.ControlPointShapes, stage.ControlPointShapes_reference, stage)

	__gong__computeReferencePass2(stage.Deliverables, stage.Deliverables_reference, stage)

	__gong__computeReferencePass2(stage.DeliverableCompositionShapes, stage.DeliverableCompositionShapes_reference, stage)

	__gong__computeReferencePass2(stage.DeliverableConceptShapes, stage.DeliverableConceptShapes_reference, stage)

	__gong__computeReferencePass2(stage.DeliverableShapes, stage.DeliverableShapes_reference, stage)

	__gong__computeReferencePass2(stage.DeliverableViewShapes, stage.DeliverableViewShapes_reference, stage)

	__gong__computeReferencePass2(stage.Diagrams, stage.Diagrams_reference, stage)

	__gong__computeReferencePass2(stage.DiagramShapes, stage.DiagramShapes_reference, stage)

	__gong__computeReferencePass2(stage.Librarys, stage.Librarys_reference, stage)

	__gong__computeReferencePass2(stage.Notes, stage.Notes_reference, stage)

	__gong__computeReferencePass2(stage.NoteDeliverableShapes, stage.NoteDeliverableShapes_reference, stage)

	__gong__computeReferencePass2(stage.NoteShapes, stage.NoteShapes_reference, stage)

	__gong__computeReferencePass2(stage.NoteStakeholderShapes, stage.NoteStakeholderShapes_reference, stage)

	__gong__computeReferencePass2(stage.NoteTaskShapes, stage.NoteTaskShapes_reference, stage)

	__gong__computeReferencePass2(stage.Requirements, stage.Requirements_reference, stage)

	__gong__computeReferencePass2(stage.RequirementConceptShapes, stage.RequirementConceptShapes_reference, stage)

	__gong__computeReferencePass2(stage.RequirementConcernShapes, stage.RequirementConcernShapes_reference, stage)

	__gong__computeReferencePass2(stage.RequirementShapes, stage.RequirementShapes_reference, stage)

	__gong__computeReferencePass2(stage.RequirementViewShapes, stage.RequirementViewShapes_reference, stage)

	__gong__computeReferencePass2(stage.Stakeholders, stage.Stakeholders_reference, stage)

	__gong__computeReferencePass2(stage.StakeholderCompositionShapes, stage.StakeholderCompositionShapes_reference, stage)

	__gong__computeReferencePass2(stage.StakeholderConcernShapes, stage.StakeholderConcernShapes_reference, stage)

	__gong__computeReferencePass2(stage.StakeholderShapes, stage.StakeholderShapes_reference, stage)

	__gong__computeReferencePass2(stage.SupportLevels, stage.SupportLevels_reference, stage)

	__gong__computeReferencePass2(stage.Tools, stage.Tools_reference, stage)

	__gong__computeReferencePass2(stage.Views, stage.Views_reference, stage)

	__gong__computeReferencePass2(stage.ViewConceptShapes, stage.ViewConceptShapes_reference, stage)

	__gong__computeReferencePass2(stage.ViewShapes, stage.ViewShapes_reference, stage)

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
	return __gong__getOrder(stage.AnalysisNeed_stagedOrder, stage.AnalysisNeeds_referenceOrder, analysisneed, "AnalysisNeed")
}

func (concept *Concept) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Concept_stagedOrder, stage.Concepts_referenceOrder, concept, "Concept")
}

func (conceptshape *ConceptShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ConceptShape_stagedOrder, stage.ConceptShapes_referenceOrder, conceptshape, "ConceptShape")
}

func (concern *Concern) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Concern_stagedOrder, stage.Concerns_referenceOrder, concern, "Concern")
}

func (concerncompositionshape *ConcernCompositionShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ConcernCompositionShape_stagedOrder, stage.ConcernCompositionShapes_referenceOrder, concerncompositionshape, "ConcernCompositionShape")
}

func (concerninputshape *ConcernInputShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ConcernInputShape_stagedOrder, stage.ConcernInputShapes_referenceOrder, concerninputshape, "ConcernInputShape")
}

func (concernoutputshape *ConcernOutputShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ConcernOutputShape_stagedOrder, stage.ConcernOutputShapes_referenceOrder, concernoutputshape, "ConcernOutputShape")
}

func (concernshape *ConcernShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ConcernShape_stagedOrder, stage.ConcernShapes_referenceOrder, concernshape, "ConcernShape")
}

func (controlpointshape *ControlPointShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ControlPointShape_stagedOrder, stage.ControlPointShapes_referenceOrder, controlpointshape, "ControlPointShape")
}

func (deliverable *Deliverable) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Deliverable_stagedOrder, stage.Deliverables_referenceOrder, deliverable, "Deliverable")
}

func (deliverablecompositionshape *DeliverableCompositionShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DeliverableCompositionShape_stagedOrder, stage.DeliverableCompositionShapes_referenceOrder, deliverablecompositionshape, "DeliverableCompositionShape")
}

func (deliverableconceptshape *DeliverableConceptShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DeliverableConceptShape_stagedOrder, stage.DeliverableConceptShapes_referenceOrder, deliverableconceptshape, "DeliverableConceptShape")
}

func (deliverableshape *DeliverableShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DeliverableShape_stagedOrder, stage.DeliverableShapes_referenceOrder, deliverableshape, "DeliverableShape")
}

func (deliverableviewshape *DeliverableViewShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DeliverableViewShape_stagedOrder, stage.DeliverableViewShapes_referenceOrder, deliverableviewshape, "DeliverableViewShape")
}

func (diagram *Diagram) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Diagram_stagedOrder, stage.Diagrams_referenceOrder, diagram, "Diagram")
}

func (diagramshape *DiagramShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DiagramShape_stagedOrder, stage.DiagramShapes_referenceOrder, diagramshape, "DiagramShape")
}

func (library *Library) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Library_stagedOrder, stage.Librarys_referenceOrder, library, "Library")
}

func (note *Note) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Note_stagedOrder, stage.Notes_referenceOrder, note, "Note")
}

func (notedeliverableshape *NoteDeliverableShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoteDeliverableShape_stagedOrder, stage.NoteDeliverableShapes_referenceOrder, notedeliverableshape, "NoteDeliverableShape")
}

func (noteshape *NoteShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoteShape_stagedOrder, stage.NoteShapes_referenceOrder, noteshape, "NoteShape")
}

func (notestakeholdershape *NoteStakeholderShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoteStakeholderShape_stagedOrder, stage.NoteStakeholderShapes_referenceOrder, notestakeholdershape, "NoteStakeholderShape")
}

func (notetaskshape *NoteTaskShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoteTaskShape_stagedOrder, stage.NoteTaskShapes_referenceOrder, notetaskshape, "NoteTaskShape")
}

func (requirement *Requirement) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Requirement_stagedOrder, stage.Requirements_referenceOrder, requirement, "Requirement")
}

func (requirementconceptshape *RequirementConceptShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RequirementConceptShape_stagedOrder, stage.RequirementConceptShapes_referenceOrder, requirementconceptshape, "RequirementConceptShape")
}

func (requirementconcernshape *RequirementConcernShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RequirementConcernShape_stagedOrder, stage.RequirementConcernShapes_referenceOrder, requirementconcernshape, "RequirementConcernShape")
}

func (requirementshape *RequirementShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RequirementShape_stagedOrder, stage.RequirementShapes_referenceOrder, requirementshape, "RequirementShape")
}

func (requirementviewshape *RequirementViewShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RequirementViewShape_stagedOrder, stage.RequirementViewShapes_referenceOrder, requirementviewshape, "RequirementViewShape")
}

func (stakeholder *Stakeholder) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Stakeholder_stagedOrder, stage.Stakeholders_referenceOrder, stakeholder, "Stakeholder")
}

func (stakeholdercompositionshape *StakeholderCompositionShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StakeholderCompositionShape_stagedOrder, stage.StakeholderCompositionShapes_referenceOrder, stakeholdercompositionshape, "StakeholderCompositionShape")
}

func (stakeholderconcernshape *StakeholderConcernShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StakeholderConcernShape_stagedOrder, stage.StakeholderConcernShapes_referenceOrder, stakeholderconcernshape, "StakeholderConcernShape")
}

func (stakeholdershape *StakeholderShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StakeholderShape_stagedOrder, stage.StakeholderShapes_referenceOrder, stakeholdershape, "StakeholderShape")
}

func (supportlevel *SupportLevel) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SupportLevel_stagedOrder, stage.SupportLevels_referenceOrder, supportlevel, "SupportLevel")
}

func (tool *Tool) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Tool_stagedOrder, stage.Tools_referenceOrder, tool, "Tool")
}

func (view *View) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.View_stagedOrder, stage.Views_referenceOrder, view, "View")
}

func (viewconceptshape *ViewConceptShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ViewConceptShape_stagedOrder, stage.ViewConceptShapes_referenceOrder, viewconceptshape, "ViewConceptShape")
}

func (viewshape *ViewShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ViewShape_stagedOrder, stage.ViewShapes_referenceOrder, viewshape, "ViewShape")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (analysisneed *AnalysisNeed) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(analysisneed, analysisneed.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (analysisneed *AnalysisNeed) GongGetReferenceIdentifier(stage *Stage) string {
	return analysisneed.GongGetIdentifier(stage)
}

func (concept *Concept) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(concept, concept.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (concept *Concept) GongGetReferenceIdentifier(stage *Stage) string {
	return concept.GongGetIdentifier(stage)
}

func (conceptshape *ConceptShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(conceptshape, conceptshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (conceptshape *ConceptShape) GongGetReferenceIdentifier(stage *Stage) string {
	return conceptshape.GongGetIdentifier(stage)
}

func (concern *Concern) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(concern, concern.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (concern *Concern) GongGetReferenceIdentifier(stage *Stage) string {
	return concern.GongGetIdentifier(stage)
}

func (concerncompositionshape *ConcernCompositionShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(concerncompositionshape, concerncompositionshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (concerncompositionshape *ConcernCompositionShape) GongGetReferenceIdentifier(stage *Stage) string {
	return concerncompositionshape.GongGetIdentifier(stage)
}

func (concerninputshape *ConcernInputShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(concerninputshape, concerninputshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (concerninputshape *ConcernInputShape) GongGetReferenceIdentifier(stage *Stage) string {
	return concerninputshape.GongGetIdentifier(stage)
}

func (concernoutputshape *ConcernOutputShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(concernoutputshape, concernoutputshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (concernoutputshape *ConcernOutputShape) GongGetReferenceIdentifier(stage *Stage) string {
	return concernoutputshape.GongGetIdentifier(stage)
}

func (concernshape *ConcernShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(concernshape, concernshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (concernshape *ConcernShape) GongGetReferenceIdentifier(stage *Stage) string {
	return concernshape.GongGetIdentifier(stage)
}

func (controlpointshape *ControlPointShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(controlpointshape, controlpointshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (controlpointshape *ControlPointShape) GongGetReferenceIdentifier(stage *Stage) string {
	return controlpointshape.GongGetIdentifier(stage)
}

func (deliverable *Deliverable) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(deliverable, deliverable.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (deliverable *Deliverable) GongGetReferenceIdentifier(stage *Stage) string {
	return deliverable.GongGetIdentifier(stage)
}

func (deliverablecompositionshape *DeliverableCompositionShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(deliverablecompositionshape, deliverablecompositionshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (deliverablecompositionshape *DeliverableCompositionShape) GongGetReferenceIdentifier(stage *Stage) string {
	return deliverablecompositionshape.GongGetIdentifier(stage)
}

func (deliverableconceptshape *DeliverableConceptShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(deliverableconceptshape, deliverableconceptshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (deliverableconceptshape *DeliverableConceptShape) GongGetReferenceIdentifier(stage *Stage) string {
	return deliverableconceptshape.GongGetIdentifier(stage)
}

func (deliverableshape *DeliverableShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(deliverableshape, deliverableshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (deliverableshape *DeliverableShape) GongGetReferenceIdentifier(stage *Stage) string {
	return deliverableshape.GongGetIdentifier(stage)
}

func (deliverableviewshape *DeliverableViewShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(deliverableviewshape, deliverableviewshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (deliverableviewshape *DeliverableViewShape) GongGetReferenceIdentifier(stage *Stage) string {
	return deliverableviewshape.GongGetIdentifier(stage)
}

func (diagram *Diagram) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(diagram, diagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (diagram *Diagram) GongGetReferenceIdentifier(stage *Stage) string {
	return diagram.GongGetIdentifier(stage)
}

func (diagramshape *DiagramShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(diagramshape, diagramshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (diagramshape *DiagramShape) GongGetReferenceIdentifier(stage *Stage) string {
	return diagramshape.GongGetIdentifier(stage)
}

func (library *Library) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(library, library.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (library *Library) GongGetReferenceIdentifier(stage *Stage) string {
	return library.GongGetIdentifier(stage)
}

func (note *Note) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(note, note.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (note *Note) GongGetReferenceIdentifier(stage *Stage) string {
	return note.GongGetIdentifier(stage)
}

func (notedeliverableshape *NoteDeliverableShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(notedeliverableshape, notedeliverableshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notedeliverableshape *NoteDeliverableShape) GongGetReferenceIdentifier(stage *Stage) string {
	return notedeliverableshape.GongGetIdentifier(stage)
}

func (noteshape *NoteShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(noteshape, noteshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (noteshape *NoteShape) GongGetReferenceIdentifier(stage *Stage) string {
	return noteshape.GongGetIdentifier(stage)
}

func (notestakeholdershape *NoteStakeholderShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(notestakeholdershape, notestakeholdershape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notestakeholdershape *NoteStakeholderShape) GongGetReferenceIdentifier(stage *Stage) string {
	return notestakeholdershape.GongGetIdentifier(stage)
}

func (notetaskshape *NoteTaskShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(notetaskshape, notetaskshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notetaskshape *NoteTaskShape) GongGetReferenceIdentifier(stage *Stage) string {
	return notetaskshape.GongGetIdentifier(stage)
}

func (requirement *Requirement) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(requirement, requirement.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (requirement *Requirement) GongGetReferenceIdentifier(stage *Stage) string {
	return requirement.GongGetIdentifier(stage)
}

func (requirementconceptshape *RequirementConceptShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(requirementconceptshape, requirementconceptshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (requirementconceptshape *RequirementConceptShape) GongGetReferenceIdentifier(stage *Stage) string {
	return requirementconceptshape.GongGetIdentifier(stage)
}

func (requirementconcernshape *RequirementConcernShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(requirementconcernshape, requirementconcernshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (requirementconcernshape *RequirementConcernShape) GongGetReferenceIdentifier(stage *Stage) string {
	return requirementconcernshape.GongGetIdentifier(stage)
}

func (requirementshape *RequirementShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(requirementshape, requirementshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (requirementshape *RequirementShape) GongGetReferenceIdentifier(stage *Stage) string {
	return requirementshape.GongGetIdentifier(stage)
}

func (requirementviewshape *RequirementViewShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(requirementviewshape, requirementviewshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (requirementviewshape *RequirementViewShape) GongGetReferenceIdentifier(stage *Stage) string {
	return requirementviewshape.GongGetIdentifier(stage)
}

func (stakeholder *Stakeholder) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stakeholder, stakeholder.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stakeholder *Stakeholder) GongGetReferenceIdentifier(stage *Stage) string {
	return stakeholder.GongGetIdentifier(stage)
}

func (stakeholdercompositionshape *StakeholderCompositionShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stakeholdercompositionshape, stakeholdercompositionshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stakeholdercompositionshape *StakeholderCompositionShape) GongGetReferenceIdentifier(stage *Stage) string {
	return stakeholdercompositionshape.GongGetIdentifier(stage)
}

func (stakeholderconcernshape *StakeholderConcernShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stakeholderconcernshape, stakeholderconcernshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stakeholderconcernshape *StakeholderConcernShape) GongGetReferenceIdentifier(stage *Stage) string {
	return stakeholderconcernshape.GongGetIdentifier(stage)
}

func (stakeholdershape *StakeholderShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stakeholdershape, stakeholdershape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stakeholdershape *StakeholderShape) GongGetReferenceIdentifier(stage *Stage) string {
	return stakeholdershape.GongGetIdentifier(stage)
}

func (supportlevel *SupportLevel) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(supportlevel, supportlevel.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (supportlevel *SupportLevel) GongGetReferenceIdentifier(stage *Stage) string {
	return supportlevel.GongGetIdentifier(stage)
}

func (tool *Tool) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tool, tool.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tool *Tool) GongGetReferenceIdentifier(stage *Stage) string {
	return tool.GongGetIdentifier(stage)
}

func (view *View) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(view, view.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (view *View) GongGetReferenceIdentifier(stage *Stage) string {
	return view.GongGetIdentifier(stage)
}

func (viewconceptshape *ViewConceptShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(viewconceptshape, viewconceptshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (viewconceptshape *ViewConceptShape) GongGetReferenceIdentifier(stage *Stage) string {
	return viewconceptshape.GongGetIdentifier(stage)
}

func (viewshape *ViewShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(viewshape, viewshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (viewshape *ViewShape) GongGetReferenceIdentifier(stage *Stage) string {
	return viewshape.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (analysisneed *AnalysisNeed) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(analysisneed.GongGetIdentifier(stage), "AnalysisNeed", analysisneed.Name)
}

func (concept *Concept) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(concept.GongGetIdentifier(stage), "Concept", concept.Name)
}

func (conceptshape *ConceptShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(conceptshape.GongGetIdentifier(stage), "ConceptShape", conceptshape.Name)
}

func (concern *Concern) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(concern.GongGetIdentifier(stage), "Concern", concern.Name)
}

func (concerncompositionshape *ConcernCompositionShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(concerncompositionshape.GongGetIdentifier(stage), "ConcernCompositionShape", concerncompositionshape.Name)
}

func (concerninputshape *ConcernInputShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(concerninputshape.GongGetIdentifier(stage), "ConcernInputShape", concerninputshape.Name)
}

func (concernoutputshape *ConcernOutputShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(concernoutputshape.GongGetIdentifier(stage), "ConcernOutputShape", concernoutputshape.Name)
}

func (concernshape *ConcernShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(concernshape.GongGetIdentifier(stage), "ConcernShape", concernshape.Name)
}

func (controlpointshape *ControlPointShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(controlpointshape.GongGetIdentifier(stage), "ControlPointShape", controlpointshape.Name)
}

func (deliverable *Deliverable) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(deliverable.GongGetIdentifier(stage), "Deliverable", deliverable.Name)
}

func (deliverablecompositionshape *DeliverableCompositionShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(deliverablecompositionshape.GongGetIdentifier(stage), "DeliverableCompositionShape", deliverablecompositionshape.Name)
}

func (deliverableconceptshape *DeliverableConceptShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(deliverableconceptshape.GongGetIdentifier(stage), "DeliverableConceptShape", deliverableconceptshape.Name)
}

func (deliverableshape *DeliverableShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(deliverableshape.GongGetIdentifier(stage), "DeliverableShape", deliverableshape.Name)
}

func (deliverableviewshape *DeliverableViewShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(deliverableviewshape.GongGetIdentifier(stage), "DeliverableViewShape", deliverableviewshape.Name)
}

func (diagram *Diagram) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(diagram.GongGetIdentifier(stage), "Diagram", diagram.Name)
}

func (diagramshape *DiagramShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(diagramshape.GongGetIdentifier(stage), "DiagramShape", diagramshape.Name)
}

func (library *Library) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(library.GongGetIdentifier(stage), "Library", library.Name)
}

func (note *Note) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(note.GongGetIdentifier(stage), "Note", note.Name)
}

func (notedeliverableshape *NoteDeliverableShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(notedeliverableshape.GongGetIdentifier(stage), "NoteDeliverableShape", notedeliverableshape.Name)
}

func (noteshape *NoteShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(noteshape.GongGetIdentifier(stage), "NoteShape", noteshape.Name)
}

func (notestakeholdershape *NoteStakeholderShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(notestakeholdershape.GongGetIdentifier(stage), "NoteStakeholderShape", notestakeholdershape.Name)
}

func (notetaskshape *NoteTaskShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(notetaskshape.GongGetIdentifier(stage), "NoteTaskShape", notetaskshape.Name)
}

func (requirement *Requirement) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(requirement.GongGetIdentifier(stage), "Requirement", requirement.Name)
}

func (requirementconceptshape *RequirementConceptShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(requirementconceptshape.GongGetIdentifier(stage), "RequirementConceptShape", requirementconceptshape.Name)
}

func (requirementconcernshape *RequirementConcernShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(requirementconcernshape.GongGetIdentifier(stage), "RequirementConcernShape", requirementconcernshape.Name)
}

func (requirementshape *RequirementShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(requirementshape.GongGetIdentifier(stage), "RequirementShape", requirementshape.Name)
}

func (requirementviewshape *RequirementViewShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(requirementviewshape.GongGetIdentifier(stage), "RequirementViewShape", requirementviewshape.Name)
}

func (stakeholder *Stakeholder) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stakeholder.GongGetIdentifier(stage), "Stakeholder", stakeholder.Name)
}

func (stakeholdercompositionshape *StakeholderCompositionShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stakeholdercompositionshape.GongGetIdentifier(stage), "StakeholderCompositionShape", stakeholdercompositionshape.Name)
}

func (stakeholderconcernshape *StakeholderConcernShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stakeholderconcernshape.GongGetIdentifier(stage), "StakeholderConcernShape", stakeholderconcernshape.Name)
}

func (stakeholdershape *StakeholderShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stakeholdershape.GongGetIdentifier(stage), "StakeholderShape", stakeholdershape.Name)
}

func (supportlevel *SupportLevel) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(supportlevel.GongGetIdentifier(stage), "SupportLevel", supportlevel.Name)
}

func (tool *Tool) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tool.GongGetIdentifier(stage), "Tool", tool.Name)
}

func (view *View) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(view.GongGetIdentifier(stage), "View", view.Name)
}

func (viewconceptshape *ViewConceptShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(viewconceptshape.GongGetIdentifier(stage), "ViewConceptShape", viewconceptshape.Name)
}

func (viewshape *ViewShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(viewshape.GongGetIdentifier(stage), "ViewShape", viewshape.Name)
}

// insertion point for unstaging
func (analysisneed *AnalysisNeed) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(analysisneed.GongGetReferenceIdentifier(stage))
}

func (concept *Concept) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(concept.GongGetReferenceIdentifier(stage))
}

func (conceptshape *ConceptShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(conceptshape.GongGetReferenceIdentifier(stage))
}

func (concern *Concern) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(concern.GongGetReferenceIdentifier(stage))
}

func (concerncompositionshape *ConcernCompositionShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(concerncompositionshape.GongGetReferenceIdentifier(stage))
}

func (concerninputshape *ConcernInputShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(concerninputshape.GongGetReferenceIdentifier(stage))
}

func (concernoutputshape *ConcernOutputShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(concernoutputshape.GongGetReferenceIdentifier(stage))
}

func (concernshape *ConcernShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(concernshape.GongGetReferenceIdentifier(stage))
}

func (controlpointshape *ControlPointShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(controlpointshape.GongGetReferenceIdentifier(stage))
}

func (deliverable *Deliverable) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(deliverable.GongGetReferenceIdentifier(stage))
}

func (deliverablecompositionshape *DeliverableCompositionShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(deliverablecompositionshape.GongGetReferenceIdentifier(stage))
}

func (deliverableconceptshape *DeliverableConceptShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(deliverableconceptshape.GongGetReferenceIdentifier(stage))
}

func (deliverableshape *DeliverableShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(deliverableshape.GongGetReferenceIdentifier(stage))
}

func (deliverableviewshape *DeliverableViewShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(deliverableviewshape.GongGetReferenceIdentifier(stage))
}

func (diagram *Diagram) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(diagram.GongGetReferenceIdentifier(stage))
}

func (diagramshape *DiagramShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(diagramshape.GongGetReferenceIdentifier(stage))
}

func (library *Library) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(library.GongGetReferenceIdentifier(stage))
}

func (note *Note) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(note.GongGetReferenceIdentifier(stage))
}

func (notedeliverableshape *NoteDeliverableShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(notedeliverableshape.GongGetReferenceIdentifier(stage))
}

func (noteshape *NoteShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(noteshape.GongGetReferenceIdentifier(stage))
}

func (notestakeholdershape *NoteStakeholderShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(notestakeholdershape.GongGetReferenceIdentifier(stage))
}

func (notetaskshape *NoteTaskShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(notetaskshape.GongGetReferenceIdentifier(stage))
}

func (requirement *Requirement) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(requirement.GongGetReferenceIdentifier(stage))
}

func (requirementconceptshape *RequirementConceptShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(requirementconceptshape.GongGetReferenceIdentifier(stage))
}

func (requirementconcernshape *RequirementConcernShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(requirementconcernshape.GongGetReferenceIdentifier(stage))
}

func (requirementshape *RequirementShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(requirementshape.GongGetReferenceIdentifier(stage))
}

func (requirementviewshape *RequirementViewShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(requirementviewshape.GongGetReferenceIdentifier(stage))
}

func (stakeholder *Stakeholder) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stakeholder.GongGetReferenceIdentifier(stage))
}

func (stakeholdercompositionshape *StakeholderCompositionShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stakeholdercompositionshape.GongGetReferenceIdentifier(stage))
}

func (stakeholderconcernshape *StakeholderConcernShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stakeholderconcernshape.GongGetReferenceIdentifier(stage))
}

func (stakeholdershape *StakeholderShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stakeholdershape.GongGetReferenceIdentifier(stage))
}

func (supportlevel *SupportLevel) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(supportlevel.GongGetReferenceIdentifier(stage))
}

func (tool *Tool) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tool.GongGetReferenceIdentifier(stage))
}

func (view *View) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(view.GongGetReferenceIdentifier(stage))
}

func (viewconceptshape *ViewConceptShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(viewconceptshape.GongGetReferenceIdentifier(stage))
}

func (viewshape *ViewShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(viewshape.GongGetReferenceIdentifier(stage))
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

func __gong__appendInstances[T interface {
	comparable
	GongstructIF
}](res []GongstructIF, m map[T]struct{}) []GongstructIF {
	for instance := range m {
		res = append(res, instance)
	}
	return res
}

func __gong__getUUID(stage *Stage, instance GongstructIF) string {
	if __gong__, ok := any(instance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}
	return GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(instance), uint64(stage.GetOrder(instance)))
}

func __gong__computeReferencePass1[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	staged map[T]struct{},
	ref *map[T]T,
	refOrder *map[T]uint,
	inst *map[T]T,
) {
	*ref = make(map[T]T, len(staged))
	*refOrder = make(map[T]uint, len(staged))
	*inst = make(map[T]T, len(staged))
	for instance := range staged {
		_copy := instance.GongCopy().(T)
		(*ref)[instance] = _copy
		(*inst)[_copy] = instance
		(*refOrder)[_copy] = instance.GongGetOrder(stage)
	}
}

func __gong__computeReferencePass2[T interface {
	comparable
	GongstructIF
	GongReconstructPointersFromReferences(*Stage, T)
}](staged map[T]struct{}, reference map[T]T, stage *Stage) {
	for instance := range staged {
		reference[instance].GongReconstructPointersFromReferences(stage, instance)
	}
}

func __gong__getOrder[T comparable](stagedOrder, refOrder map[T]uint, instance T, typeName string) uint {
	if order, ok := stagedOrder[instance]; ok {
		return order
	}
	if order, ok := refOrder[instance]; ok {
		return order
	}
	log.Printf("instance %p of type %s was not staged and does not have a reference order", any(instance), typeName)
	return 0
}

func __gong__formatIdentifier(s GongstructIF, order uint) string {
	return fmt.Sprintf("__models__%s__%08d_", s.GongGetGongstructName(), order)
}

func __gong__marshallIdentifier(identifier, structName, name string) string {
	decl := strings.ReplaceAll(GongIdentifiersDecls, "{{Identifier}}", identifier)
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", structName)
	return strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(name))
}

func __gong__marshallUnstaging(identifier string) string {
	return strings.ReplaceAll(GongUnstageStmt, "{{Identifier}}", identifier)
}

// end of template
