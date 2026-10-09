// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/process/go/models"
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
	case *models.AllocatedProcessShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Participant", inst.Participant, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Participant](), probe.formStage)
		StageSetAssociationFieldToForm("Process", inst.Process, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Process](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DiagramProcesss {
				for _, target := range src.AllocatedProcessShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DiagramProcess", "AllocatedProcessShapes", refNames, formGroup, probe.formStage)
		}
	case *models.AllocatedResourceShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Participant", inst.Participant, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Participant](), probe.formStage)
		StageSetAssociationFieldToForm("Resource", inst.Resource, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Resource](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DiagramProcesss {
				for _, target := range src.AllocatedResourceShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DiagramProcess", "AllocatedResourceShapes", refNames, formGroup, probe.formStage)
		}
	case *models.ControlFlow:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Start", inst.Start, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Task](), probe.formStage)
		StageSetAssociationFieldToForm("End", inst.End, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Task](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ControlFlowShapes {
				if src.ControlFlow == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ControlFlowShape", "ControlFlow", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DiagramProcesss {
				for _, target := range src.ControlFlowsWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DiagramProcess", "ControlFlowsWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Participants {
				for _, target := range src.ControlFlows {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Participant", "ControlFlows", refNames, formGroup, probe.formStage)
		}
	case *models.ControlFlowShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ControlFlow", inst.ControlFlow, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ControlFlow](), probe.formStage)
		StageSetBasicFieldtoForm("StartRatio", inst.StartRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndRatio", inst.EndRatio, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("StartOrientation", inst.StartOrientation, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EndOrientation", inst.EndOrientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("CornerOffsetRatio", inst.CornerOffsetRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DiagramProcesss {
				for _, target := range src.ControlFlow_Shapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DiagramProcess", "ControlFlow_Shapes", refNames, formGroup, probe.formStage)
		}
	case *models.Data:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Acronym", inst.Acronym, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SVG_Path", inst.SVG_Path, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("InverseAppliedScaling", inst.InverseAppliedScaling, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DataFlows {
				for _, target := range src.Datas {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DataFlow", "Datas", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DataShapes {
				if src.Data == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DataShape", "Data", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DiagramProcesss {
				for _, target := range src.DatasWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DiagramProcess", "DatasWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.RootDatas {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "RootDatas", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.DatasWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "DatasWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}
	case *models.DataFlow:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: Datas
			div := (&form.FormDiv{Name: "Datas"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Datas {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Datas",
				Label: "Datas",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetAssociationFieldToForm("StartTask", inst.StartTask, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Task](), probe.formStage)
		StageSetAssociationFieldToForm("EndTask", inst.EndTask, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Task](), probe.formStage)
		StageSetAssociationFieldToForm("StartExternalParticipant", inst.StartExternalParticipant, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Participant](), probe.formStage)
		StageSetAssociationFieldToForm("EndExternalParticipant", inst.EndExternalParticipant, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Participant](), probe.formStage)
		StageSetBasicFieldtoForm("IsDatasNodeExpanded", inst.IsDatasNodeExpanded, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DataFlowShapes {
				if src.DataFlow == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DataFlowShape", "DataFlow", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DataShapes {
				if src.DataFlow == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DataShape", "DataFlow", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DiagramProcesss {
				for _, target := range src.DataFlowsWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DiagramProcess", "DataFlowsWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DiagramProcesss {
				for _, target := range src.DataFlowsWhoseDataNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DiagramProcess", "DataFlowsWhoseDataNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.RootDataFlows {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "RootDataFlows", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.DataFlowsWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "DataFlowsWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Processs {
				for _, target := range src.DataFlows {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Process", "DataFlows", refNames, formGroup, probe.formStage)
		}
	case *models.DataFlowShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("DataFlow", inst.DataFlow, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.DataFlow](), probe.formStage)
		StageSetBasicFieldtoForm("StartRatio", inst.StartRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndRatio", inst.EndRatio, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("StartOrientation", inst.StartOrientation, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EndOrientation", inst.EndOrientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("CornerOffsetRatio", inst.CornerOffsetRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DiagramProcesss {
				for _, target := range src.DataFlow_Shapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DiagramProcess", "DataFlow_Shapes", refNames, formGroup, probe.formStage)
		}
	case *models.DataShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Data", inst.Data, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Data](), probe.formStage)
		StageSetAssociationFieldToForm("DataFlow", inst.DataFlow, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.DataFlow](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DiagramProcesss {
				for _, target := range src.Data_Shapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DiagramProcess", "Data_Shapes", refNames, formGroup, probe.formStage)
		}
	case *models.DiagramProcess:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsChecked", inst.IsChecked, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsEditable_", inst.IsEditable_, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsShowPrefix", inst.IsShowPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultBoxWidth", inst.DefaultBoxWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultBoxHeigth", inst.DefaultBoxHeigth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)

		{
			// Slice of pointers: Process_Shapes
			div := (&form.FormDiv{Name: "Process_Shapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Process_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Process_Shapes",
				Label: "Process_Shapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsProcesssNodeExpanded", inst.IsProcesssNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: ProcesssWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "ProcesssWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ProcesssWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ProcesssWhoseNodeIsExpanded",
				Label: "ProcesssWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Participant_Shapes
			div := (&form.FormDiv{Name: "Participant_Shapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Participant_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Participant_Shapes",
				Label: "Participant_Shapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsParticipantsNodeExpanded", inst.IsParticipantsNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: ParticipantWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "ParticipantWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ParticipantWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ParticipantWhoseNodeIsExpanded",
				Label: "ParticipantWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ExternalParticipant_Shapes
			div := (&form.FormDiv{Name: "ExternalParticipant_Shapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ExternalParticipant_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ExternalParticipant_Shapes",
				Label: "ExternalParticipant_Shapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsExternalParticipantsNodeExpanded", inst.IsExternalParticipantsNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: ExternalParticipantWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "ExternalParticipantWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ExternalParticipantWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ExternalParticipantWhoseNodeIsExpanded",
				Label: "ExternalParticipantWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded
			div := (&form.FormDiv{Name: "ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded",
				Label: "ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ExternalParticipantsWhoseInDataFlowsNodeIsExpanded
			div := (&form.FormDiv{Name: "ExternalParticipantsWhoseInDataFlowsNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ExternalParticipantsWhoseInDataFlowsNodeIsExpanded",
				Label: "ExternalParticipantsWhoseInDataFlowsNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: TasksWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "TasksWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TasksWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TasksWhoseNodeIsExpanded",
				Label: "TasksWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Task_Shapes
			div := (&form.FormDiv{Name: "Task_Shapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Task_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Task_Shapes",
				Label: "Task_Shapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ControlFlowsWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "ControlFlowsWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ControlFlowsWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ControlFlowsWhoseNodeIsExpanded",
				Label: "ControlFlowsWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ControlFlow_Shapes
			div := (&form.FormDiv{Name: "ControlFlow_Shapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ControlFlow_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ControlFlow_Shapes",
				Label: "ControlFlow_Shapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: DataFlowsWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "DataFlowsWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.DataFlowsWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "DataFlowsWhoseNodeIsExpanded",
				Label: "DataFlowsWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: DataFlow_Shapes
			div := (&form.FormDiv{Name: "DataFlow_Shapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.DataFlow_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "DataFlow_Shapes",
				Label: "DataFlow_Shapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: DatasWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "DatasWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.DatasWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "DatasWhoseNodeIsExpanded",
				Label: "DatasWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Data_Shapes
			div := (&form.FormDiv{Name: "Data_Shapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Data_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Data_Shapes",
				Label: "Data_Shapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: DataFlowsWhoseDataNodeIsExpanded
			div := (&form.FormDiv{Name: "DataFlowsWhoseDataNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.DataFlowsWhoseDataNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "DataFlowsWhoseDataNodeIsExpanded",
				Label: "DataFlowsWhoseDataNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AllocatedResourcesWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "AllocatedResourcesWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AllocatedResourcesWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AllocatedResourcesWhoseNodeIsExpanded",
				Label: "AllocatedResourcesWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AllocatedResourceShapes
			div := (&form.FormDiv{Name: "AllocatedResourceShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AllocatedResourceShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AllocatedResourceShapes",
				Label: "AllocatedResourceShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AllocatedProcessesWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "AllocatedProcessesWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AllocatedProcessesWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AllocatedProcessesWhoseNodeIsExpanded",
				Label: "AllocatedProcessesWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AllocatedProcessShapes
			div := (&form.FormDiv{Name: "AllocatedProcessShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AllocatedProcessShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AllocatedProcessShapes",
				Label: "AllocatedProcessShapes",
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
			var refNames []string
			for src := range probe.stageSet.Stage.Processs {
				for _, target := range src.DiagramProcesss {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Process", "DiagramProcesss", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Processs {
				for _, target := range src.DiagramProcessWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Process", "DiagramProcessWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}
	case *models.ExternalParticipantShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Participant", inst.Participant, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Participant](), probe.formStage)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TailHeigth", inst.TailHeigth, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DiagramProcesss {
				for _, target := range src.ExternalParticipant_Shapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DiagramProcess", "ExternalParticipant_Shapes", refNames, formGroup, probe.formStage)
		}
	case *models.Library:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsRootLibrary", inst.IsRootLibrary, probe.formStage, formGroup)

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
		StageSetBasicFieldtoForm("IsSubLibrariesNodeExpanded", inst.IsSubLibrariesNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: SubLibrariesWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "SubLibrariesWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SubLibrariesWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SubLibrariesWhoseNodeIsExpanded",
				Label: "SubLibrariesWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("NbPixPerCharacter", inst.NbPixPerCharacter, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LogoSVGFile", inst.LogoSVGFile, probe.formStage, formGroup)

		{
			// Slice of pointers: RootProcesses
			div := (&form.FormDiv{Name: "RootProcesses"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RootProcesses {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RootProcesses",
				Label: "RootProcesses",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsProcessesNodeExpanded", inst.IsProcessesNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: ProcesssWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "ProcesssWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ProcesssWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ProcesssWhoseNodeIsExpanded",
				Label: "ProcesssWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: RootDataFlows
			div := (&form.FormDiv{Name: "RootDataFlows"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RootDataFlows {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RootDataFlows",
				Label: "RootDataFlows",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsDataFlowsNodeExpanded", inst.IsDataFlowsNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: DataFlowsWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "DataFlowsWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.DataFlowsWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "DataFlowsWhoseNodeIsExpanded",
				Label: "DataFlowsWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: RootDatas
			div := (&form.FormDiv{Name: "RootDatas"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RootDatas {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RootDatas",
				Label: "RootDatas",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsDatasNodeExpanded", inst.IsDatasNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: DatasWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "DatasWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.DatasWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "DatasWhoseNodeIsExpanded",
				Label: "DatasWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: RootResources
			div := (&form.FormDiv{Name: "RootResources"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RootResources {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RootResources",
				Label: "RootResources",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsResourcesNodeExpanded", inst.IsResourcesNodeExpanded, probe.formStage, formGroup)

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

		{
			// Slice of pointers: ParticipantsWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "ParticipantsWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ParticipantsWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ParticipantsWhoseNodeIsExpanded",
				Label: "ParticipantsWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: RootNotes
			div := (&form.FormDiv{Name: "RootNotes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RootNotes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RootNotes",
				Label: "RootNotes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsNotesNodeExpanded", inst.IsNotesNodeExpanded, probe.formStage, formGroup)

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
		StageSetBasicFieldtoForm("IsExpandedTmp", inst.IsExpandedTmp, probe.formStage, formGroup)

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

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.SubLibrariesWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "SubLibrariesWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}
	case *models.Note:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsTasksNodeExpanded", inst.IsTasksNodeExpanded, probe.formStage, formGroup)

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
			var refNames []string
			for src := range probe.stageSet.Stage.DiagramProcesss {
				for _, target := range src.NotesWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DiagramProcess", "NotesWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.RootNotes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "RootNotes", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.NotesWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "NotesWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageSet.Stage.NoteTaskShapes {
				if src.Note == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.NoteTaskShape", "Note", refNames, formGroup, probe.formStage)
		}
	case *models.NoteShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Note", inst.Note, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Note](), probe.formStage)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DiagramProcesss {
				for _, target := range src.Note_Shapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DiagramProcess", "Note_Shapes", refNames, formGroup, probe.formStage)
		}
	case *models.NoteTaskShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Note", inst.Note, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Note](), probe.formStage)
		StageSetAssociationFieldToForm("Task", inst.Task, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Task](), probe.formStage)
		StageSetBasicFieldtoForm("StartRatio", inst.StartRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndRatio", inst.EndRatio, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("StartOrientation", inst.StartOrientation, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EndOrientation", inst.EndOrientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("CornerOffsetRatio", inst.CornerOffsetRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DiagramProcesss {
				for _, target := range src.NoteTaskShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DiagramProcess", "NoteTaskShapes", refNames, formGroup, probe.formStage)
		}
	case *models.Participant:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsProcessResource", inst.IsProcessResource, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)

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
		StageSetBasicFieldtoForm("IsResourcesNodeExpanded", inst.IsResourcesNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: Processes
			div := (&form.FormDiv{Name: "Processes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Processes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Processes",
				Label: "Processes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsProcessesNodeExpanded", inst.IsProcessesNodeExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsTasksNodeExpanded", inst.IsTasksNodeExpanded, probe.formStage, formGroup)

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
		StageSetBasicFieldtoForm("IsControlFlowsNodeExpanded", inst.IsControlFlowsNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: ControlFlows
			div := (&form.FormDiv{Name: "ControlFlows"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ControlFlows {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ControlFlows",
				Label: "ControlFlows",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: TaskWhoseOutControlFlowsNodeIsExpanded
			div := (&form.FormDiv{Name: "TaskWhoseOutControlFlowsNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TaskWhoseOutControlFlowsNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TaskWhoseOutControlFlowsNodeIsExpanded",
				Label: "TaskWhoseOutControlFlowsNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: TaskWhoseInControlFlowsNodeIsExpanded
			div := (&form.FormDiv{Name: "TaskWhoseInControlFlowsNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TaskWhoseInControlFlowsNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TaskWhoseInControlFlowsNodeIsExpanded",
				Label: "TaskWhoseInControlFlowsNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsDataFlowsNodeExpanded", inst.IsDataFlowsNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: TaskWhoseOutDataFlowsNodeIsExpanded
			div := (&form.FormDiv{Name: "TaskWhoseOutDataFlowsNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TaskWhoseOutDataFlowsNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TaskWhoseOutDataFlowsNodeIsExpanded",
				Label: "TaskWhoseOutDataFlowsNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: TaskWhoseInDataFlowsNodeIsExpanded
			div := (&form.FormDiv{Name: "TaskWhoseInDataFlowsNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TaskWhoseInDataFlowsNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TaskWhoseInDataFlowsNodeIsExpanded",
				Label: "TaskWhoseInDataFlowsNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.AllocatedProcessShapes {
				if src.Participant == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AllocatedProcessShape", "Participant", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.AllocatedResourceShapes {
				if src.Participant == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AllocatedResourceShape", "Participant", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DataFlows {
				if src.StartExternalParticipant == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DataFlow", "StartExternalParticipant", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DataFlows {
				if src.EndExternalParticipant == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DataFlow", "EndExternalParticipant", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DiagramProcesss {
				for _, target := range src.ParticipantWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DiagramProcess", "ParticipantWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DiagramProcesss {
				for _, target := range src.ExternalParticipantWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DiagramProcess", "ExternalParticipantWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DiagramProcesss {
				for _, target := range src.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DiagramProcess", "ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DiagramProcesss {
				for _, target := range src.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DiagramProcess", "ExternalParticipantsWhoseInDataFlowsNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ExternalParticipantShapes {
				if src.Participant == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ExternalParticipantShape", "Participant", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.ParticipantsWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "ParticipantsWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ParticipantShapes {
				if src.Participant == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ParticipantShape", "Participant", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Processs {
				for _, target := range src.Participants {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Process", "Participants", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Processs {
				for _, target := range src.ParticipantWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Process", "ParticipantWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Processs {
				for _, target := range src.ExternalParticipants {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Process", "ExternalParticipants", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Processs {
				for _, target := range src.ExternalParticipantWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Process", "ExternalParticipantWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}
	case *models.ParticipantShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Participant", inst.Participant, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Participant](), probe.formStage)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("WidthWeight", inst.WidthWeight, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DiagramProcesss {
				for _, target := range src.Participant_Shapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DiagramProcess", "Participant_Shapes", refNames, formGroup, probe.formStage)
		}
	case *models.Process:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SVG_Path", inst.SVG_Path, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("InverseAppliedScaling", inst.InverseAppliedScaling, probe.formStage, formGroup)

		{
			// Slice of pointers: DiagramProcesss
			div := (&form.FormDiv{Name: "DiagramProcesss"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.DiagramProcesss {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "DiagramProcesss",
				Label: "DiagramProcesss",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: DiagramProcessWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "DiagramProcessWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.DiagramProcessWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "DiagramProcessWhoseNodeIsExpanded",
				Label: "DiagramProcessWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsSubProcessNodeExpanded", inst.IsSubProcessNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: SubProcesses
			div := (&form.FormDiv{Name: "SubProcesses"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SubProcesses {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SubProcesses",
				Label: "SubProcesses",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Participants
			div := (&form.FormDiv{Name: "Participants"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Participants {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Participants",
				Label: "Participants",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ParticipantWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "ParticipantWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ParticipantWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ParticipantWhoseNodeIsExpanded",
				Label: "ParticipantWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: DataFlows
			div := (&form.FormDiv{Name: "DataFlows"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.DataFlows {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "DataFlows",
				Label: "DataFlows",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsDataFlowsNodeExpanded", inst.IsDataFlowsNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: ExternalParticipants
			div := (&form.FormDiv{Name: "ExternalParticipants"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ExternalParticipants {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ExternalParticipants",
				Label: "ExternalParticipants",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ExternalParticipantWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "ExternalParticipantWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ExternalParticipantWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ExternalParticipantWhoseNodeIsExpanded",
				Label: "ExternalParticipantWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.AllocatedProcessShapes {
				if src.Process == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AllocatedProcessShape", "Process", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DiagramProcesss {
				for _, target := range src.ProcesssWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DiagramProcess", "ProcesssWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DiagramProcesss {
				for _, target := range src.AllocatedProcessesWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DiagramProcess", "AllocatedProcessesWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.RootProcesses {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "RootProcesses", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.ProcesssWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "ProcesssWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Participants {
				for _, target := range src.Processes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Participant", "Processes", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Processs {
				for _, target := range src.SubProcesses {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Process", "SubProcesses", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ProcessShapes {
				if src.Process == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ProcessShape", "Process", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Tasks {
				if src.Type == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Task", "Type", refNames, formGroup, probe.formStage)
		}
	case *models.ProcessShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Process", inst.Process, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Process](), probe.formStage)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DiagramProcesss {
				for _, target := range src.Process_Shapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DiagramProcess", "Process_Shapes", refNames, formGroup, probe.formStage)
		}
	case *models.Resource:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Acronym", inst.Acronym, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SVG_Path", inst.SVG_Path, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("InverseAppliedScaling", inst.InverseAppliedScaling, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.AllocatedResourceShapes {
				if src.Resource == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AllocatedResourceShape", "Resource", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DiagramProcesss {
				for _, target := range src.AllocatedResourcesWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DiagramProcess", "AllocatedResourcesWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.RootResources {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "RootResources", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.ResourcesWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "ResourcesWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Participants {
				for _, target := range src.Resources {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Participant", "Resources", refNames, formGroup, probe.formStage)
		}
	case *models.Task:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsStartTask", inst.IsStartTask, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsEndTask", inst.IsEndTask, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Type", inst.Type, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Process](), probe.formStage)
		StageSetBasicFieldtoForm("IsTaskNameNotProcessName", inst.IsTaskNameNotProcessName, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ControlFlows {
				if src.Start == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ControlFlow", "Start", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ControlFlows {
				if src.End == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ControlFlow", "End", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DataFlows {
				if src.StartTask == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DataFlow", "StartTask", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DataFlows {
				if src.EndTask == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DataFlow", "EndTask", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DiagramProcesss {
				for _, target := range src.TasksWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DiagramProcess", "TasksWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageSet.Stage.Participants {
				for _, target := range src.Tasks {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Participant", "Tasks", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Participants {
				for _, target := range src.TaskWhoseOutControlFlowsNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Participant", "TaskWhoseOutControlFlowsNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Participants {
				for _, target := range src.TaskWhoseInControlFlowsNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Participant", "TaskWhoseInControlFlowsNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Participants {
				for _, target := range src.TaskWhoseOutDataFlowsNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Participant", "TaskWhoseOutDataFlowsNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Participants {
				for _, target := range src.TaskWhoseInDataFlowsNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Participant", "TaskWhoseInDataFlowsNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.TaskShapes {
				if src.Task == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.TaskShape", "Task", refNames, formGroup, probe.formStage)
		}
	case *models.TaskShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Task", inst.Task, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Task](), probe.formStage)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DiagramProcesss {
				for _, target := range src.Task_Shapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DiagramProcess", "Task_Shapes", refNames, formGroup, probe.formStage)
		}
	default:
		_ = inst
	}
}
