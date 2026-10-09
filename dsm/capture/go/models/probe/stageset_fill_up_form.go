// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/capture/go/models"
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func StageSetFillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *StageSetProbe,
) {
	switch inst := instance.(type) {
	case *models.AnalysisNeed:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.AnalysisNeeds {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "AnalysisNeeds", refNames, formGroup, probe.formStage)
		}
	case *models.Concept:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: Tools
			div := (&form.FormDiv{Name: "Tools"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Tools {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Tools",
				Label: "Tools",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ConceptShapes {
				if src.Concept == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ConceptShape", "Concept", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Deliverables {
				for _, target := range src.Concepts {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Deliverable", "Concepts", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DeliverableConceptShapes {
				if src.Concept == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DeliverableConceptShape", "Concept", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ConceptsWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "ConceptsWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ConceptsWhoseDeliverablesNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "ConceptsWhoseDeliverablesNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.RootConcepts {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "RootConcepts", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Requirements {
				for _, target := range src.Concepts {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Requirement", "Concepts", refNames, formGroup, probe.formStage)
		}
	case *models.ConceptShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Concept", inst.Concept, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Concept](), probe.formStage)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.Concept_Shapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "Concept_Shapes", refNames, formGroup, probe.formStage)
		}
	case *models.Concern:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IDAirbus", inst.IDAirbus, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Priority", inst.Priority, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)

		{
			// Slice of pointers: SubConcerns
			div := (&form.FormDiv{Name: "SubConcerns"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SubConcerns {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SubConcerns",
				Label: "SubConcerns",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Inputs
			div := (&form.FormDiv{Name: "Inputs"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Inputs {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Inputs",
				Label: "Inputs",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsInputsNodeExpanded", inst.IsInputsNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: Outputs
			div := (&form.FormDiv{Name: "Outputs"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Outputs {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Outputs",
				Label: "Outputs",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsOutputsNodeExpanded", inst.IsOutputsNodeExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsWithCompletion", inst.IsWithCompletion, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Completion", inst.Completion, formGroup, probe.formStage)

		{
			// Slice of pointers: Requirements
			div := (&form.FormDiv{Name: "Requirements"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Requirements {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Requirements",
				Label: "Requirements",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Concerns {
				for _, target := range src.SubConcerns {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Concern", "SubConcerns", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ConcernCompositionShapes {
				if src.Concern == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ConcernCompositionShape", "Concern", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ConcernInputShapes {
				if src.Concern == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ConcernInputShape", "Concern", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ConcernOutputShapes {
				if src.Concern == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ConcernOutputShape", "Concern", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ConcernShapes {
				if src.Concern == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ConcernShape", "Concern", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ConcernsWhoseRequirementsNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "ConcernsWhoseRequirementsNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ConcernsWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "ConcernsWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ConcernsWhoseInputNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "ConcernsWhoseInputNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ConcernsWhoseStakeholderNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "ConcernsWhoseStakeholderNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ConcernssWhoseOutputNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "ConcernssWhoseOutputNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.RootConcerns {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "RootConcerns", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notes {
				for _, target := range src.Tasks {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Tasks", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.NoteTaskShapes {
				if src.Task == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.NoteTaskShape", "Task", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Stakeholders {
				for _, target := range src.Concerns {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Stakeholder", "Concerns", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.StakeholderConcernShapes {
				if src.Concern == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.StakeholderConcernShape", "Concern", refNames, formGroup, probe.formStage)
		}
	case *models.ConcernCompositionShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Concern", inst.Concern, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Concern](), probe.formStage)
		StageSetBasicFieldtoForm("StartRatio", inst.StartRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndRatio", inst.EndRatio, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("StartOrientation", inst.StartOrientation, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EndOrientation", inst.EndOrientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("CornerOffsetRatio", inst.CornerOffsetRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			// Slice of pointers: ControlPointShapes
			div := (&form.FormDiv{Name: "ControlPointShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ControlPointShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ControlPointShapes",
				Label: "ControlPointShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ConcernComposition_Shapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "ConcernComposition_Shapes", refNames, formGroup, probe.formStage)
		}
	case *models.ConcernInputShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Deliverable", inst.Deliverable, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Deliverable](), probe.formStage)
		StageSetAssociationFieldToForm("Concern", inst.Concern, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Concern](), probe.formStage)
		StageSetBasicFieldtoForm("StartRatio", inst.StartRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndRatio", inst.EndRatio, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("StartOrientation", inst.StartOrientation, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EndOrientation", inst.EndOrientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("CornerOffsetRatio", inst.CornerOffsetRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			// Slice of pointers: ControlPointShapes
			div := (&form.FormDiv{Name: "ControlPointShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ControlPointShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ControlPointShapes",
				Label: "ControlPointShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ConcernInputShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "ConcernInputShapes", refNames, formGroup, probe.formStage)
		}
	case *models.ConcernOutputShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Concern", inst.Concern, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Concern](), probe.formStage)
		StageSetAssociationFieldToForm("Deliverable", inst.Deliverable, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Deliverable](), probe.formStage)
		StageSetBasicFieldtoForm("StartRatio", inst.StartRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndRatio", inst.EndRatio, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("StartOrientation", inst.StartOrientation, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EndOrientation", inst.EndOrientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("CornerOffsetRatio", inst.CornerOffsetRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			// Slice of pointers: ControlPointShapes
			div := (&form.FormDiv{Name: "ControlPointShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ControlPointShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ControlPointShapes",
				Label: "ControlPointShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ConcernOutputShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "ConcernOutputShapes", refNames, formGroup, probe.formStage)
		}
	case *models.ConcernShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Concern", inst.Concern, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Concern](), probe.formStage)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.Concern_Shapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "Concern_Shapes", refNames, formGroup, probe.formStage)
		}
	case *models.ControlPointShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X_Relative", inst.X_Relative, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y_Relative", inst.Y_Relative, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsStartShapeTheClosestShape", inst.IsStartShapeTheClosestShape, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ConcernCompositionShapes {
				for _, target := range src.ControlPointShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ConcernCompositionShape", "ControlPointShapes", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ConcernInputShapes {
				for _, target := range src.ControlPointShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ConcernInputShape", "ControlPointShapes", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ConcernOutputShapes {
				for _, target := range src.ControlPointShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ConcernOutputShape", "ControlPointShapes", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DeliverableCompositionShapes {
				for _, target := range src.ControlPointShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DeliverableCompositionShape", "ControlPointShapes", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DeliverableConceptShapes {
				for _, target := range src.ControlPointShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DeliverableConceptShape", "ControlPointShapes", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.NoteDeliverableShapes {
				for _, target := range src.ControlPointShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.NoteDeliverableShape", "ControlPointShapes", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.NoteStakeholderShapes {
				for _, target := range src.ControlPointShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.NoteStakeholderShape", "ControlPointShapes", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.NoteTaskShapes {
				for _, target := range src.ControlPointShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.NoteTaskShape", "ControlPointShapes", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.StakeholderCompositionShapes {
				for _, target := range src.ControlPointShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.StakeholderCompositionShape", "ControlPointShapes", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.StakeholderConcernShapes {
				for _, target := range src.ControlPointShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.StakeholderConcernShape", "ControlPointShapes", refNames, formGroup, probe.formStage)
		}
	case *models.Deliverable:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)

		{
			// Slice of pointers: SubDeliverables
			div := (&form.FormDiv{Name: "SubDeliverables"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SubDeliverables {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SubDeliverables",
				Label: "SubDeliverables",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsProducersNodeExpanded", inst.IsProducersNodeExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsConsumersNodeExpanded", inst.IsConsumersNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: Concepts
			div := (&form.FormDiv{Name: "Concepts"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Concepts {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Concepts",
				Label: "Concepts",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Concerns {
				for _, target := range src.Inputs {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Concern", "Inputs", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Concerns {
				for _, target := range src.Outputs {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Concern", "Outputs", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ConcernInputShapes {
				if src.Deliverable == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ConcernInputShape", "Deliverable", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ConcernOutputShapes {
				if src.Deliverable == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ConcernOutputShape", "Deliverable", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Deliverables {
				for _, target := range src.SubDeliverables {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Deliverable", "SubDeliverables", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DeliverableCompositionShapes {
				if src.Deliverable == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DeliverableCompositionShape", "Deliverable", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DeliverableConceptShapes {
				if src.Deliverable == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DeliverableConceptShape", "Deliverable", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DeliverableShapes {
				if src.Deliverable == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DeliverableShape", "Deliverable", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.DeliverablesWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "DeliverablesWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.DeliverablesWhoseConceptsNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "DeliverablesWhoseConceptsNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.RootDeliverables {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "RootDeliverables", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notes {
				for _, target := range src.Deliverables {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Deliverables", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.NoteDeliverableShapes {
				if src.Deliverable == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.NoteDeliverableShape", "Deliverable", refNames, formGroup, probe.formStage)
		}
	case *models.DeliverableCompositionShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Deliverable", inst.Deliverable, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Deliverable](), probe.formStage)
		StageSetBasicFieldtoForm("StartRatio", inst.StartRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndRatio", inst.EndRatio, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("StartOrientation", inst.StartOrientation, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EndOrientation", inst.EndOrientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("CornerOffsetRatio", inst.CornerOffsetRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			// Slice of pointers: ControlPointShapes
			div := (&form.FormDiv{Name: "ControlPointShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ControlPointShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ControlPointShapes",
				Label: "ControlPointShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.DeliverableComposition_Shapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "DeliverableComposition_Shapes", refNames, formGroup, probe.formStage)
		}
	case *models.DeliverableConceptShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Deliverable", inst.Deliverable, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Deliverable](), probe.formStage)
		StageSetAssociationFieldToForm("Concept", inst.Concept, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Concept](), probe.formStage)
		StageSetBasicFieldtoForm("StartRatio", inst.StartRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndRatio", inst.EndRatio, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("StartOrientation", inst.StartOrientation, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EndOrientation", inst.EndOrientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("CornerOffsetRatio", inst.CornerOffsetRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			// Slice of pointers: ControlPointShapes
			div := (&form.FormDiv{Name: "ControlPointShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ControlPointShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ControlPointShapes",
				Label: "ControlPointShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.DeliverableConceptShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "DeliverableConceptShapes", refNames, formGroup, probe.formStage)
		}
	case *models.DeliverableShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Deliverable", inst.Deliverable, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Deliverable](), probe.formStage)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.Deliverable_Shapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "Deliverable_Shapes", refNames, formGroup, probe.formStage)
		}
	case *models.Diagram:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsChecked", inst.IsChecked, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsEditable_", inst.IsEditable_, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ShowPrefix", inst.ShowPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultBoxWidth", inst.DefaultBoxWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultBoxHeigth", inst.DefaultBoxHeigth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)

		{
			// Slice of pointers: ConcernsWhoseRequirementsNodeIsExpanded
			div := (&form.FormDiv{Name: "ConcernsWhoseRequirementsNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ConcernsWhoseRequirementsNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ConcernsWhoseRequirementsNodeIsExpanded",
				Label: "ConcernsWhoseRequirementsNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsRequirementsNodeExpanded", inst.IsRequirementsNodeExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsConceptsNodeExpanded", inst.IsConceptsNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: Deliverable_Shapes
			div := (&form.FormDiv{Name: "Deliverable_Shapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Deliverable_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Deliverable_Shapes",
				Label: "Deliverable_Shapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: DeliverablesWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "DeliverablesWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.DeliverablesWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "DeliverablesWhoseNodeIsExpanded",
				Label: "DeliverablesWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: DeliverablesWhoseConceptsNodeIsExpanded
			div := (&form.FormDiv{Name: "DeliverablesWhoseConceptsNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.DeliverablesWhoseConceptsNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "DeliverablesWhoseConceptsNodeIsExpanded",
				Label: "DeliverablesWhoseConceptsNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsPBSNodeExpanded", inst.IsPBSNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: DeliverableComposition_Shapes
			div := (&form.FormDiv{Name: "DeliverableComposition_Shapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.DeliverableComposition_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "DeliverableComposition_Shapes",
				Label: "DeliverableComposition_Shapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsConcernsNodeExpanded", inst.IsConcernsNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: Concern_Shapes
			div := (&form.FormDiv{Name: "Concern_Shapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Concern_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Concern_Shapes",
				Label: "Concern_Shapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ConcernsWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "ConcernsWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ConcernsWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ConcernsWhoseNodeIsExpanded",
				Label: "ConcernsWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ConcernsWhoseInputNodeIsExpanded
			div := (&form.FormDiv{Name: "ConcernsWhoseInputNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ConcernsWhoseInputNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ConcernsWhoseInputNodeIsExpanded",
				Label: "ConcernsWhoseInputNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ConcernsWhoseStakeholderNodeIsExpanded
			div := (&form.FormDiv{Name: "ConcernsWhoseStakeholderNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ConcernsWhoseStakeholderNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ConcernsWhoseStakeholderNodeIsExpanded",
				Label: "ConcernsWhoseStakeholderNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ConcernssWhoseOutputNodeIsExpanded
			div := (&form.FormDiv{Name: "ConcernssWhoseOutputNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ConcernssWhoseOutputNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ConcernssWhoseOutputNodeIsExpanded",
				Label: "ConcernssWhoseOutputNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ConcernComposition_Shapes
			div := (&form.FormDiv{Name: "ConcernComposition_Shapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ConcernComposition_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ConcernComposition_Shapes",
				Label: "ConcernComposition_Shapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ConcernInputShapes
			div := (&form.FormDiv{Name: "ConcernInputShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ConcernInputShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ConcernInputShapes",
				Label: "ConcernInputShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ConcernOutputShapes
			div := (&form.FormDiv{Name: "ConcernOutputShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ConcernOutputShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ConcernOutputShapes",
				Label: "ConcernOutputShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Note_Shapes
			div := (&form.FormDiv{Name: "Note_Shapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Note_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Note_Shapes",
				Label: "Note_Shapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: NotesWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "NotesWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NotesWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NotesWhoseNodeIsExpanded",
				Label: "NotesWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsNotesNodeExpanded", inst.IsNotesNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: NoteDeliverableShapes
			div := (&form.FormDiv{Name: "NoteDeliverableShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NoteDeliverableShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NoteDeliverableShapes",
				Label: "NoteDeliverableShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: NoteTaskShapes
			div := (&form.FormDiv{Name: "NoteTaskShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NoteTaskShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NoteTaskShapes",
				Label: "NoteTaskShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: NoteResourceShapes
			div := (&form.FormDiv{Name: "NoteResourceShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NoteResourceShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NoteResourceShapes",
				Label: "NoteResourceShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Stakeholder_Shapes
			div := (&form.FormDiv{Name: "Stakeholder_Shapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Stakeholder_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Stakeholder_Shapes",
				Label: "Stakeholder_Shapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ResourcesWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "ResourcesWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ResourcesWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ResourcesWhoseNodeIsExpanded",
				Label: "ResourcesWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsStakeholdersNodeExpanded", inst.IsStakeholdersNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: ResourceComposition_Shapes
			div := (&form.FormDiv{Name: "ResourceComposition_Shapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ResourceComposition_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ResourceComposition_Shapes",
				Label: "ResourceComposition_Shapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: StakeholderConcernShapes
			div := (&form.FormDiv{Name: "StakeholderConcernShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.StakeholderConcernShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "StakeholderConcernShapes",
				Label: "StakeholderConcernShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Requirement_Shapes
			div := (&form.FormDiv{Name: "Requirement_Shapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Requirement_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Requirement_Shapes",
				Label: "Requirement_Shapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: RequirementsWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "RequirementsWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RequirementsWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RequirementsWhoseNodeIsExpanded",
				Label: "RequirementsWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Concept_Shapes
			div := (&form.FormDiv{Name: "Concept_Shapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Concept_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Concept_Shapes",
				Label: "Concept_Shapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ConceptsWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "ConceptsWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ConceptsWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ConceptsWhoseNodeIsExpanded",
				Label: "ConceptsWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ConceptsWhoseDeliverablesNodeIsExpanded
			div := (&form.FormDiv{Name: "ConceptsWhoseDeliverablesNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ConceptsWhoseDeliverablesNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ConceptsWhoseDeliverablesNodeIsExpanded",
				Label: "ConceptsWhoseDeliverablesNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: DeliverableConceptShapes
			div := (&form.FormDiv{Name: "DeliverableConceptShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.DeliverableConceptShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "DeliverableConceptShapes",
				Label: "DeliverableConceptShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Diagram_Shapes
			div := (&form.FormDiv{Name: "Diagram_Shapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Diagram_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Diagram_Shapes",
				Label: "Diagram_Shapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsDiagramsNodeExpanded", inst.IsDiagramsNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: DiagramsWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "DiagramsWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.DiagramsWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "DiagramsWhoseNodeIsExpanded",
				Label: "DiagramsWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.DiagramsWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "DiagramsWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DiagramShapes {
				if src.Diagram == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DiagramShape", "Diagram", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.Diagrams {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "Diagrams", refNames, formGroup, probe.formStage)
		}
	case *models.DiagramShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Diagram", inst.Diagram, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Diagram](), probe.formStage)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.Diagram_Shapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "Diagram_Shapes", refNames, formGroup, probe.formStage)
		}
	case *models.Library:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsRootLibrary", inst.IsRootLibrary, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: RootDeliverables
			div := (&form.FormDiv{Name: "RootDeliverables"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RootDeliverables {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RootDeliverables",
				Label: "RootDeliverables",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: RootConcerns
			div := (&form.FormDiv{Name: "RootConcerns"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RootConcerns {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RootConcerns",
				Label: "RootConcerns",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: RootStakeholders
			div := (&form.FormDiv{Name: "RootStakeholders"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RootStakeholders {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RootStakeholders",
				Label: "RootStakeholders",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: RootRequirements
			div := (&form.FormDiv{Name: "RootRequirements"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RootRequirements {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RootRequirements",
				Label: "RootRequirements",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: RootConcepts
			div := (&form.FormDiv{Name: "RootConcepts"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RootConcepts {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RootConcepts",
				Label: "RootConcepts",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AnalysisNeeds
			div := (&form.FormDiv{Name: "AnalysisNeeds"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AnalysisNeeds {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AnalysisNeeds",
				Label: "AnalysisNeeds",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Notes
			div := (&form.FormDiv{Name: "Notes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Notes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Notes",
				Label: "Notes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Diagrams
			div := (&form.FormDiv{Name: "Diagrams"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Diagrams {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Diagrams",
				Label: "Diagrams",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: SubLibraries
			div := (&form.FormDiv{Name: "SubLibraries"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SubLibraries {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SubLibraries",
				Label: "SubLibraries",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("NbPixPerCharacter", inst.NbPixPerCharacter, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.SubLibraries {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "SubLibraries", refNames, formGroup, probe.formStage)
		}
	case *models.Note:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: Deliverables
			div := (&form.FormDiv{Name: "Deliverables"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Deliverables {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Deliverables",
				Label: "Deliverables",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Tasks
			div := (&form.FormDiv{Name: "Tasks"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Tasks {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Tasks",
				Label: "Tasks",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Resources
			div := (&form.FormDiv{Name: "Resources"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Resources {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Resources",
				Label: "Resources",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.NotesWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "NotesWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.Notes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "Notes", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.NoteDeliverableShapes {
				if src.Note == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.NoteDeliverableShape", "Note", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.NoteShapes {
				if src.Note == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.NoteShape", "Note", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.NoteStakeholderShapes {
				if src.Note == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.NoteStakeholderShape", "Note", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.NoteTaskShapes {
				if src.Note == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.NoteTaskShape", "Note", refNames, formGroup, probe.formStage)
		}
	case *models.NoteDeliverableShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Note", inst.Note, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Note](), probe.formStage)
		StageSetAssociationFieldToForm("Deliverable", inst.Deliverable, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Deliverable](), probe.formStage)
		StageSetBasicFieldtoForm("StartRatio", inst.StartRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndRatio", inst.EndRatio, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("StartOrientation", inst.StartOrientation, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EndOrientation", inst.EndOrientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("CornerOffsetRatio", inst.CornerOffsetRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			// Slice of pointers: ControlPointShapes
			div := (&form.FormDiv{Name: "ControlPointShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ControlPointShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ControlPointShapes",
				Label: "ControlPointShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.NoteDeliverableShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "NoteDeliverableShapes", refNames, formGroup, probe.formStage)
		}
	case *models.NoteShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Note", inst.Note, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Note](), probe.formStage)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.Note_Shapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "Note_Shapes", refNames, formGroup, probe.formStage)
		}
	case *models.NoteStakeholderShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Note", inst.Note, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Note](), probe.formStage)
		StageSetAssociationFieldToForm("Stakeholder", inst.Stakeholder, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Stakeholder](), probe.formStage)
		StageSetBasicFieldtoForm("StartRatio", inst.StartRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndRatio", inst.EndRatio, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("StartOrientation", inst.StartOrientation, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EndOrientation", inst.EndOrientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("CornerOffsetRatio", inst.CornerOffsetRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			// Slice of pointers: ControlPointShapes
			div := (&form.FormDiv{Name: "ControlPointShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ControlPointShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ControlPointShapes",
				Label: "ControlPointShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.NoteResourceShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "NoteResourceShapes", refNames, formGroup, probe.formStage)
		}
	case *models.NoteTaskShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Note", inst.Note, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Note](), probe.formStage)
		StageSetAssociationFieldToForm("Task", inst.Task, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Concern](), probe.formStage)
		StageSetBasicFieldtoForm("StartRatio", inst.StartRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndRatio", inst.EndRatio, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("StartOrientation", inst.StartOrientation, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EndOrientation", inst.EndOrientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("CornerOffsetRatio", inst.CornerOffsetRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			// Slice of pointers: ControlPointShapes
			div := (&form.FormDiv{Name: "ControlPointShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ControlPointShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ControlPointShapes",
				Label: "ControlPointShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.NoteTaskShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "NoteTaskShapes", refNames, formGroup, probe.formStage)
		}
	case *models.Requirement:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: SupportLevels
			div := (&form.FormDiv{Name: "SupportLevels"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SupportLevels {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SupportLevels",
				Label: "SupportLevels",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Concepts
			div := (&form.FormDiv{Name: "Concepts"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Concepts {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Concepts",
				Label: "Concepts",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Concerns {
				for _, target := range src.Requirements {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Concern", "Requirements", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.RequirementsWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "RequirementsWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.RootRequirements {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "RootRequirements", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.RequirementShapes {
				if src.Requirement == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.RequirementShape", "Requirement", refNames, formGroup, probe.formStage)
		}
	case *models.RequirementShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Requirement", inst.Requirement, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Requirement](), probe.formStage)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.Requirement_Shapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "Requirement_Shapes", refNames, formGroup, probe.formStage)
		}
	case *models.Stakeholder:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IDAirbus", inst.IDAirbus, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)

		{
			// Slice of pointers: Concerns
			div := (&form.FormDiv{Name: "Concerns"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Concerns {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Concerns",
				Label: "Concerns",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: SubStakeholders
			div := (&form.FormDiv{Name: "SubStakeholders"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SubStakeholders {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SubStakeholders",
				Label: "SubStakeholders",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ResourcesWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "ResourcesWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.RootStakeholders {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "RootStakeholders", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notes {
				for _, target := range src.Resources {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Resources", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.NoteStakeholderShapes {
				if src.Stakeholder == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.NoteStakeholderShape", "Stakeholder", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Stakeholders {
				for _, target := range src.SubStakeholders {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Stakeholder", "SubStakeholders", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.StakeholderCompositionShapes {
				if src.Stakeholder == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.StakeholderCompositionShape", "Stakeholder", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.StakeholderConcernShapes {
				if src.Stakeholder == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.StakeholderConcernShape", "Stakeholder", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.StakeholderShapes {
				if src.Stakeholder == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.StakeholderShape", "Stakeholder", refNames, formGroup, probe.formStage)
		}
	case *models.StakeholderCompositionShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Stakeholder", inst.Stakeholder, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Stakeholder](), probe.formStage)
		StageSetBasicFieldtoForm("StartRatio", inst.StartRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndRatio", inst.EndRatio, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("StartOrientation", inst.StartOrientation, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EndOrientation", inst.EndOrientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("CornerOffsetRatio", inst.CornerOffsetRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			// Slice of pointers: ControlPointShapes
			div := (&form.FormDiv{Name: "ControlPointShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ControlPointShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ControlPointShapes",
				Label: "ControlPointShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ResourceComposition_Shapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "ResourceComposition_Shapes", refNames, formGroup, probe.formStage)
		}
	case *models.StakeholderConcernShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Stakeholder", inst.Stakeholder, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Stakeholder](), probe.formStage)
		StageSetAssociationFieldToForm("Concern", inst.Concern, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Concern](), probe.formStage)
		StageSetBasicFieldtoForm("StartRatio", inst.StartRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndRatio", inst.EndRatio, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("StartOrientation", inst.StartOrientation, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EndOrientation", inst.EndOrientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("CornerOffsetRatio", inst.CornerOffsetRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			// Slice of pointers: ControlPointShapes
			div := (&form.FormDiv{Name: "ControlPointShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ControlPointShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ControlPointShapes",
				Label: "ControlPointShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.StakeholderConcernShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "StakeholderConcernShapes", refNames, formGroup, probe.formStage)
		}
	case *models.StakeholderShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Stakeholder", inst.Stakeholder, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Stakeholder](), probe.formStage)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.Stakeholder_Shapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "Stakeholder_Shapes", refNames, formGroup, probe.formStage)
		}
	case *models.SupportLevel:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Tool", inst.Tool, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Tool](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Requirements {
				for _, target := range src.SupportLevels {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Requirement", "SupportLevels", refNames, formGroup, probe.formStage)
		}
	case *models.Tool:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Concepts {
				for _, target := range src.Tools {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Concept", "Tools", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.SupportLevels {
				if src.Tool == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SupportLevel", "Tool", refNames, formGroup, probe.formStage)
		}
	default:
		_ = inst
	}
}
