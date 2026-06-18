// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/capture/go/models"
)

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.AnalysisNeed:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Library, *models.AnalysisNeed](
				"Library",
				"AnalysisNeeds",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.AnalysisNeed {
					return owner.AnalysisNeeds
				})
		}

	case *models.Concept:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		AssociationSliceToForm("Tools", instanceWithInferedType, &instanceWithInferedType.Tools, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Deliverable, *models.Concept](
				"Deliverable",
				"Concepts",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Deliverable) []*models.Concept {
					return owner.Concepts
				})
		}
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.Concept](
				"Diagram",
				"ConceptsWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.Concept {
					return owner.ConceptsWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.Concept](
				"Diagram",
				"ConceptsWhoseDeliverablesNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.Concept {
					return owner.ConceptsWhoseDeliverablesNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.Concept](
				"Library",
				"RootConcepts",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Concept {
					return owner.RootConcepts
				})
		}
		{
			AssociationReverseSliceToForm[*models.Requirement, *models.Concept](
				"Requirement",
				"Concepts",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Requirement) []*models.Concept {
					return owner.Concepts
				})
		}

	case *models.ConceptShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Concept", instanceWithInferedType.Concept, formGroup, probe)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.ConceptShape](
				"Diagram",
				"Concept_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.ConceptShape {
					return owner.Concept_Shapes
				})
		}

	case *models.Concern:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDAirbus", instanceWithInferedType.IDAirbus, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Priority", instanceWithInferedType.Priority, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		AssociationSliceToForm("SubConcerns", instanceWithInferedType, &instanceWithInferedType.SubConcerns, formGroup, probe)
		AssociationSliceToForm("Inputs", instanceWithInferedType, &instanceWithInferedType.Inputs, formGroup, probe)
		BasicFieldtoForm("IsInputsNodeExpanded", instanceWithInferedType.IsInputsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Outputs", instanceWithInferedType, &instanceWithInferedType.Outputs, formGroup, probe)
		BasicFieldtoForm("IsOutputsNodeExpanded", instanceWithInferedType.IsOutputsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsWithCompletion", instanceWithInferedType.IsWithCompletion, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Completion", instanceWithInferedType.Completion, instanceWithInferedType, probe.formStage, formGroup)
		AssociationSliceToForm("Requirements", instanceWithInferedType, &instanceWithInferedType.Requirements, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Concern, *models.Concern](
				"Concern",
				"SubConcerns",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Concern) []*models.Concern {
					return owner.SubConcerns
				})
		}
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.Concern](
				"Diagram",
				"ConcernsWhoseRequirementsNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.Concern {
					return owner.ConcernsWhoseRequirementsNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.Concern](
				"Diagram",
				"ConcernsWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.Concern {
					return owner.ConcernsWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.Concern](
				"Diagram",
				"ConcernsWhoseInputNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.Concern {
					return owner.ConcernsWhoseInputNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.Concern](
				"Diagram",
				"ConcernsWhoseStakeholderNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.Concern {
					return owner.ConcernsWhoseStakeholderNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.Concern](
				"Diagram",
				"ConcernssWhoseOutputNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.Concern {
					return owner.ConcernssWhoseOutputNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.Concern](
				"Library",
				"RootConcerns",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Concern {
					return owner.RootConcerns
				})
		}
		{
			AssociationReverseSliceToForm[*models.Note, *models.Concern](
				"Note",
				"Tasks",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Note) []*models.Concern {
					return owner.Tasks
				})
		}
		{
			AssociationReverseSliceToForm[*models.Stakeholder, *models.Concern](
				"Stakeholder",
				"Concerns",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Stakeholder) []*models.Concern {
					return owner.Concerns
				})
		}

	case *models.ConcernCompositionShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Concern", instanceWithInferedType.Concern, formGroup, probe)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ControlPointShapes", instanceWithInferedType, &instanceWithInferedType.ControlPointShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.ConcernCompositionShape](
				"Diagram",
				"ConcernComposition_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.ConcernCompositionShape {
					return owner.ConcernComposition_Shapes
				})
		}

	case *models.ConcernInputShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Deliverable", instanceWithInferedType.Deliverable, formGroup, probe)
		AssociationFieldToForm("Concern", instanceWithInferedType.Concern, formGroup, probe)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ControlPointShapes", instanceWithInferedType, &instanceWithInferedType.ControlPointShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.ConcernInputShape](
				"Diagram",
				"ConcernInputShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.ConcernInputShape {
					return owner.ConcernInputShapes
				})
		}

	case *models.ConcernOutputShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Task", instanceWithInferedType.Task, formGroup, probe)
		AssociationFieldToForm("Deliverable", instanceWithInferedType.Deliverable, formGroup, probe)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ControlPointShapes", instanceWithInferedType, &instanceWithInferedType.ControlPointShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.ConcernOutputShape](
				"Diagram",
				"ConcernOutputShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.ConcernOutputShape {
					return owner.ConcernOutputShapes
				})
		}

	case *models.ConcernShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Concern", instanceWithInferedType.Concern, formGroup, probe)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.ConcernShape](
				"Diagram",
				"Concern_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.ConcernShape {
					return owner.Concern_Shapes
				})
		}

	case *models.ControlPointShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X_Relative", instanceWithInferedType.X_Relative, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y_Relative", instanceWithInferedType.Y_Relative, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsStartShapeTheClosestShape", instanceWithInferedType.IsStartShapeTheClosestShape, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.ConcernCompositionShape, *models.ControlPointShape](
				"ConcernCompositionShape",
				"ControlPointShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.ConcernCompositionShape) []*models.ControlPointShape {
					return owner.ControlPointShapes
				})
		}
		{
			AssociationReverseSliceToForm[*models.ConcernInputShape, *models.ControlPointShape](
				"ConcernInputShape",
				"ControlPointShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.ConcernInputShape) []*models.ControlPointShape {
					return owner.ControlPointShapes
				})
		}
		{
			AssociationReverseSliceToForm[*models.ConcernOutputShape, *models.ControlPointShape](
				"ConcernOutputShape",
				"ControlPointShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.ConcernOutputShape) []*models.ControlPointShape {
					return owner.ControlPointShapes
				})
		}
		{
			AssociationReverseSliceToForm[*models.DeliverableCompositionShape, *models.ControlPointShape](
				"DeliverableCompositionShape",
				"ControlPointShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DeliverableCompositionShape) []*models.ControlPointShape {
					return owner.ControlPointShapes
				})
		}
		{
			AssociationReverseSliceToForm[*models.DeliverableConceptShape, *models.ControlPointShape](
				"DeliverableConceptShape",
				"ControlPointShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DeliverableConceptShape) []*models.ControlPointShape {
					return owner.ControlPointShapes
				})
		}
		{
			AssociationReverseSliceToForm[*models.NoteDeliverableShape, *models.ControlPointShape](
				"NoteDeliverableShape",
				"ControlPointShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.NoteDeliverableShape) []*models.ControlPointShape {
					return owner.ControlPointShapes
				})
		}
		{
			AssociationReverseSliceToForm[*models.NoteStakeholderShape, *models.ControlPointShape](
				"NoteStakeholderShape",
				"ControlPointShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.NoteStakeholderShape) []*models.ControlPointShape {
					return owner.ControlPointShapes
				})
		}
		{
			AssociationReverseSliceToForm[*models.NoteTaskShape, *models.ControlPointShape](
				"NoteTaskShape",
				"ControlPointShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.NoteTaskShape) []*models.ControlPointShape {
					return owner.ControlPointShapes
				})
		}
		{
			AssociationReverseSliceToForm[*models.StakeholderCompositionShape, *models.ControlPointShape](
				"StakeholderCompositionShape",
				"ControlPointShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.StakeholderCompositionShape) []*models.ControlPointShape {
					return owner.ControlPointShapes
				})
		}
		{
			AssociationReverseSliceToForm[*models.StakeholderConcernShape, *models.ControlPointShape](
				"StakeholderConcernShape",
				"ControlPointShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.StakeholderConcernShape) []*models.ControlPointShape {
					return owner.ControlPointShapes
				})
		}

	case *models.Deliverable:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		AssociationSliceToForm("SubDeliverables", instanceWithInferedType, &instanceWithInferedType.SubDeliverables, formGroup, probe)
		BasicFieldtoForm("IsProducersNodeExpanded", instanceWithInferedType.IsProducersNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsConsumersNodeExpanded", instanceWithInferedType.IsConsumersNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Concepts", instanceWithInferedType, &instanceWithInferedType.Concepts, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Concern, *models.Deliverable](
				"Concern",
				"Inputs",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Concern) []*models.Deliverable {
					return owner.Inputs
				})
		}
		{
			AssociationReverseSliceToForm[*models.Concern, *models.Deliverable](
				"Concern",
				"Outputs",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Concern) []*models.Deliverable {
					return owner.Outputs
				})
		}
		{
			AssociationReverseSliceToForm[*models.Deliverable, *models.Deliverable](
				"Deliverable",
				"SubDeliverables",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Deliverable) []*models.Deliverable {
					return owner.SubDeliverables
				})
		}
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.Deliverable](
				"Diagram",
				"DeliverablesWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.Deliverable {
					return owner.DeliverablesWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.Deliverable](
				"Diagram",
				"DeliverablesWhoseConceptsNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.Deliverable {
					return owner.DeliverablesWhoseConceptsNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.Deliverable](
				"Library",
				"RootDeliverables",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Deliverable {
					return owner.RootDeliverables
				})
		}
		{
			AssociationReverseSliceToForm[*models.Note, *models.Deliverable](
				"Note",
				"Deliverables",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Note) []*models.Deliverable {
					return owner.Deliverables
				})
		}

	case *models.DeliverableCompositionShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Deliverable", instanceWithInferedType.Deliverable, formGroup, probe)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ControlPointShapes", instanceWithInferedType, &instanceWithInferedType.ControlPointShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.DeliverableCompositionShape](
				"Diagram",
				"DeliverableComposition_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.DeliverableCompositionShape {
					return owner.DeliverableComposition_Shapes
				})
		}

	case *models.DeliverableConceptShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Deliverable", instanceWithInferedType.Deliverable, formGroup, probe)
		AssociationFieldToForm("Concept", instanceWithInferedType.Concept, formGroup, probe)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ControlPointShapes", instanceWithInferedType, &instanceWithInferedType.ControlPointShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.DeliverableConceptShape](
				"Diagram",
				"DeliverableConceptShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.DeliverableConceptShape {
					return owner.DeliverableConceptShapes
				})
		}

	case *models.DeliverableShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Deliverable", instanceWithInferedType.Deliverable, formGroup, probe)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.DeliverableShape](
				"Diagram",
				"Deliverable_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.DeliverableShape {
					return owner.Deliverable_Shapes
				})
		}

	case *models.Diagram:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("IsChecked", instanceWithInferedType.IsChecked, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsEditable_", instanceWithInferedType.IsEditable_, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ShowPrefix", instanceWithInferedType.ShowPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DefaultBoxWidth", instanceWithInferedType.DefaultBoxWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DefaultBoxHeigth", instanceWithInferedType.DefaultBoxHeigth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ConcernsWhoseRequirementsNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ConcernsWhoseRequirementsNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsRequirementsNodeExpanded", instanceWithInferedType.IsRequirementsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsConceptsNodeExpanded", instanceWithInferedType.IsConceptsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Deliverable_Shapes", instanceWithInferedType, &instanceWithInferedType.Deliverable_Shapes, formGroup, probe)
		AssociationSliceToForm("DeliverablesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.DeliverablesWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("DeliverablesWhoseConceptsNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.DeliverablesWhoseConceptsNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsPBSNodeExpanded", instanceWithInferedType.IsPBSNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("DeliverableComposition_Shapes", instanceWithInferedType, &instanceWithInferedType.DeliverableComposition_Shapes, formGroup, probe)
		BasicFieldtoForm("IsConcernsNodeExpanded", instanceWithInferedType.IsConcernsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Concern_Shapes", instanceWithInferedType, &instanceWithInferedType.Concern_Shapes, formGroup, probe)
		AssociationSliceToForm("ConcernsWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ConcernsWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("ConcernsWhoseInputNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ConcernsWhoseInputNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("ConcernsWhoseStakeholderNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ConcernsWhoseStakeholderNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("ConcernssWhoseOutputNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ConcernssWhoseOutputNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("ConcernComposition_Shapes", instanceWithInferedType, &instanceWithInferedType.ConcernComposition_Shapes, formGroup, probe)
		AssociationSliceToForm("ConcernInputShapes", instanceWithInferedType, &instanceWithInferedType.ConcernInputShapes, formGroup, probe)
		AssociationSliceToForm("ConcernOutputShapes", instanceWithInferedType, &instanceWithInferedType.ConcernOutputShapes, formGroup, probe)
		AssociationSliceToForm("Note_Shapes", instanceWithInferedType, &instanceWithInferedType.Note_Shapes, formGroup, probe)
		AssociationSliceToForm("NotesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.NotesWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsNotesNodeExpanded", instanceWithInferedType.IsNotesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("NoteDeliverableShapes", instanceWithInferedType, &instanceWithInferedType.NoteDeliverableShapes, formGroup, probe)
		AssociationSliceToForm("NoteTaskShapes", instanceWithInferedType, &instanceWithInferedType.NoteTaskShapes, formGroup, probe)
		AssociationSliceToForm("NoteResourceShapes", instanceWithInferedType, &instanceWithInferedType.NoteResourceShapes, formGroup, probe)
		AssociationSliceToForm("Stakeholder_Shapes", instanceWithInferedType, &instanceWithInferedType.Stakeholder_Shapes, formGroup, probe)
		AssociationSliceToForm("ResourcesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ResourcesWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsStakeholdersNodeExpanded", instanceWithInferedType.IsStakeholdersNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ResourceComposition_Shapes", instanceWithInferedType, &instanceWithInferedType.ResourceComposition_Shapes, formGroup, probe)
		AssociationSliceToForm("StakeholderConcernShapes", instanceWithInferedType, &instanceWithInferedType.StakeholderConcernShapes, formGroup, probe)
		AssociationSliceToForm("Requirement_Shapes", instanceWithInferedType, &instanceWithInferedType.Requirement_Shapes, formGroup, probe)
		AssociationSliceToForm("RequirementsWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.RequirementsWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("Concept_Shapes", instanceWithInferedType, &instanceWithInferedType.Concept_Shapes, formGroup, probe)
		AssociationSliceToForm("ConceptsWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ConceptsWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("ConceptsWhoseDeliverablesNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ConceptsWhoseDeliverablesNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("DeliverableConceptShapes", instanceWithInferedType, &instanceWithInferedType.DeliverableConceptShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Library, *models.Diagram](
				"Library",
				"Diagrams",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Diagram {
					return owner.Diagrams
				})
		}

	case *models.Library:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsRootLibrary", instanceWithInferedType.IsRootLibrary, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		AssociationSliceToForm("RootDeliverables", instanceWithInferedType, &instanceWithInferedType.RootDeliverables, formGroup, probe)
		AssociationSliceToForm("RootConcerns", instanceWithInferedType, &instanceWithInferedType.RootConcerns, formGroup, probe)
		AssociationSliceToForm("RootStakeholders", instanceWithInferedType, &instanceWithInferedType.RootStakeholders, formGroup, probe)
		AssociationSliceToForm("RootRequirements", instanceWithInferedType, &instanceWithInferedType.RootRequirements, formGroup, probe)
		AssociationSliceToForm("RootConcepts", instanceWithInferedType, &instanceWithInferedType.RootConcepts, formGroup, probe)
		AssociationSliceToForm("AnalysisNeeds", instanceWithInferedType, &instanceWithInferedType.AnalysisNeeds, formGroup, probe)
		AssociationSliceToForm("Notes", instanceWithInferedType, &instanceWithInferedType.Notes, formGroup, probe)
		AssociationSliceToForm("Diagrams", instanceWithInferedType, &instanceWithInferedType.Diagrams, formGroup, probe)
		AssociationSliceToForm("SubLibraries", instanceWithInferedType, &instanceWithInferedType.SubLibraries, formGroup, probe)
		BasicFieldtoForm("NbPixPerCharacter", instanceWithInferedType.NbPixPerCharacter, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Library, *models.Library](
				"Library",
				"SubLibraries",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Library {
					return owner.SubLibraries
				})
		}

	case *models.Note:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		AssociationSliceToForm("Deliverables", instanceWithInferedType, &instanceWithInferedType.Deliverables, formGroup, probe)
		AssociationSliceToForm("Tasks", instanceWithInferedType, &instanceWithInferedType.Tasks, formGroup, probe)
		AssociationSliceToForm("Resources", instanceWithInferedType, &instanceWithInferedType.Resources, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.Note](
				"Diagram",
				"NotesWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.Note {
					return owner.NotesWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.Note](
				"Library",
				"Notes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Note {
					return owner.Notes
				})
		}

	case *models.NoteDeliverableShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Note", instanceWithInferedType.Note, formGroup, probe)
		AssociationFieldToForm("Deliverable", instanceWithInferedType.Deliverable, formGroup, probe)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ControlPointShapes", instanceWithInferedType, &instanceWithInferedType.ControlPointShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.NoteDeliverableShape](
				"Diagram",
				"NoteDeliverableShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.NoteDeliverableShape {
					return owner.NoteDeliverableShapes
				})
		}

	case *models.NoteShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Note", instanceWithInferedType.Note, formGroup, probe)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.NoteShape](
				"Diagram",
				"Note_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.NoteShape {
					return owner.Note_Shapes
				})
		}

	case *models.NoteStakeholderShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Note", instanceWithInferedType.Note, formGroup, probe)
		AssociationFieldToForm("Stakeholder", instanceWithInferedType.Stakeholder, formGroup, probe)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ControlPointShapes", instanceWithInferedType, &instanceWithInferedType.ControlPointShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.NoteStakeholderShape](
				"Diagram",
				"NoteResourceShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.NoteStakeholderShape {
					return owner.NoteResourceShapes
				})
		}

	case *models.NoteTaskShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Note", instanceWithInferedType.Note, formGroup, probe)
		AssociationFieldToForm("Task", instanceWithInferedType.Task, formGroup, probe)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ControlPointShapes", instanceWithInferedType, &instanceWithInferedType.ControlPointShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.NoteTaskShape](
				"Diagram",
				"NoteTaskShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.NoteTaskShape {
					return owner.NoteTaskShapes
				})
		}

	case *models.Requirement:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		AssociationSliceToForm("SupportLevels", instanceWithInferedType, &instanceWithInferedType.SupportLevels, formGroup, probe)
		AssociationSliceToForm("Concepts", instanceWithInferedType, &instanceWithInferedType.Concepts, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Concern, *models.Requirement](
				"Concern",
				"Requirements",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Concern) []*models.Requirement {
					return owner.Requirements
				})
		}
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.Requirement](
				"Diagram",
				"RequirementsWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.Requirement {
					return owner.RequirementsWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.Requirement](
				"Library",
				"RootRequirements",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Requirement {
					return owner.RootRequirements
				})
		}

	case *models.RequirementShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Requirement", instanceWithInferedType.Requirement, formGroup, probe)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.RequirementShape](
				"Diagram",
				"Requirement_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.RequirementShape {
					return owner.Requirement_Shapes
				})
		}

	case *models.Stakeholder:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IDAirbus", instanceWithInferedType.IDAirbus, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		AssociationSliceToForm("Concerns", instanceWithInferedType, &instanceWithInferedType.Concerns, formGroup, probe)
		AssociationSliceToForm("SubStakeholders", instanceWithInferedType, &instanceWithInferedType.SubStakeholders, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.Stakeholder](
				"Diagram",
				"ResourcesWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.Stakeholder {
					return owner.ResourcesWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.Stakeholder](
				"Library",
				"RootStakeholders",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Stakeholder {
					return owner.RootStakeholders
				})
		}
		{
			AssociationReverseSliceToForm[*models.Note, *models.Stakeholder](
				"Note",
				"Resources",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Note) []*models.Stakeholder {
					return owner.Resources
				})
		}
		{
			AssociationReverseSliceToForm[*models.Stakeholder, *models.Stakeholder](
				"Stakeholder",
				"SubStakeholders",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Stakeholder) []*models.Stakeholder {
					return owner.SubStakeholders
				})
		}

	case *models.StakeholderCompositionShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Stakeholder", instanceWithInferedType.Stakeholder, formGroup, probe)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ControlPointShapes", instanceWithInferedType, &instanceWithInferedType.ControlPointShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.StakeholderCompositionShape](
				"Diagram",
				"ResourceComposition_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.StakeholderCompositionShape {
					return owner.ResourceComposition_Shapes
				})
		}

	case *models.StakeholderConcernShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Stakeholder", instanceWithInferedType.Stakeholder, formGroup, probe)
		AssociationFieldToForm("Concern", instanceWithInferedType.Concern, formGroup, probe)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ControlPointShapes", instanceWithInferedType, &instanceWithInferedType.ControlPointShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.StakeholderConcernShape](
				"Diagram",
				"StakeholderConcernShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.StakeholderConcernShape {
					return owner.StakeholderConcernShapes
				})
		}

	case *models.StakeholderShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Stakeholder", instanceWithInferedType.Stakeholder, formGroup, probe)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.StakeholderShape](
				"Diagram",
				"Stakeholder_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.StakeholderShape {
					return owner.Stakeholder_Shapes
				})
		}

	case *models.SupportLevel:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Tool", instanceWithInferedType.Tool, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Requirement, *models.SupportLevel](
				"Requirement",
				"SupportLevels",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Requirement) []*models.SupportLevel {
					return owner.SupportLevels
				})
		}

	case *models.Tool:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Concept, *models.Tool](
				"Concept",
				"Tools",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Concept) []*models.Tool {
					return owner.Tools
				})
		}

	default:
		_ = instanceWithInferedType
	}
}
