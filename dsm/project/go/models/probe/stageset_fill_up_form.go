// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/project/go/models"
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
	case *models.Diagram:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultBoxWidth", inst.DefaultBoxWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DefaultBoxHeigth", inst.DefaultBoxHeigth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DateFormat", inst.DateFormat, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsTimeDiagram", inst.IsTimeDiagram, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedStart", inst.ComputedStart, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedEnd", inst.ComputedEnd, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedDuration", inst.ComputedDuration, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DrawVerticalTimeLines", inst.DrawVerticalTimeLines, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DrawSecondaryVerticalTimeLines", inst.DrawSecondaryVerticalTimeLines, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HideWeekendsPeriod", inst.HideWeekendsPeriod, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("UseManualStartAndEndDates", inst.UseManualStartAndEndDates, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ManualStart", inst.ManualStart, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ManualEnd", inst.ManualEnd, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TimeStep", inst.TimeStep, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("TimeStepScale", inst.TimeStepScale, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("SecondaryTimeStep", inst.SecondaryTimeStep, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("SecondaryTimeStepScale", inst.SecondaryTimeStepScale, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("LaneHeight", inst.LaneHeight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RatioBarToLaneHeight", inst.RatioBarToLaneHeight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("YTopMargin", inst.YTopMargin, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("XLeftText", inst.XLeftText, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TextHeight", inst.TextHeight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("XLeftLanes", inst.XLeftLanes, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("XRightMargin", inst.XRightMargin, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ArrowLengthToTheRightOfStartBar", inst.ArrowLengthToTheRightOfStartBar, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ArrowTipLenght", inst.ArrowTipLenght, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TimeLine_Color", inst.TimeLine_Color, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TimeLine_FillOpacity", inst.TimeLine_FillOpacity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TimeLine_Stroke", inst.TimeLine_Stroke, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TimeLine_StrokeWidth", inst.TimeLine_StrokeWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Group_Stroke", inst.Group_Stroke, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Group_StrokeWidth", inst.Group_StrokeWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Group_StrokeDashArray", inst.Group_StrokeDashArray, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DateYOffset", inst.DateYOffset, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("AlignOnBeginningOfTimeScale", inst.AlignOnBeginningOfTimeScale, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsChecked", inst.IsChecked, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsEditable_", inst.IsEditable_, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsShowPrefix", inst.IsShowPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsInAutoLayoutMode", inst.IsInAutoLayoutMode, probe.formStage, formGroup)

		{
			// Slice of pointers: Product_Shapes
			div := (&form.FormDiv{Name: "Product_Shapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Product_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Product_Shapes",
				Label: "Product_Shapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ProductsWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "ProductsWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ProductsWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ProductsWhoseNodeIsExpanded",
				Label: "ProductsWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsPBSNodeExpanded", inst.IsPBSNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: ProductComposition_Shapes
			div := (&form.FormDiv{Name: "ProductComposition_Shapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ProductComposition_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ProductComposition_Shapes",
				Label: "ProductComposition_Shapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: ProductReference_Shapes
			div := (&form.FormDiv{Name: "ProductReference_Shapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ProductReference_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ProductReference_Shapes",
				Label: "ProductReference_Shapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsWBSNodeExpanded", inst.IsWBSNodeExpanded, probe.formStage, formGroup)

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
			// Slice of pointers: TasksWhoseInputNodeIsExpanded
			div := (&form.FormDiv{Name: "TasksWhoseInputNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TasksWhoseInputNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TasksWhoseInputNodeIsExpanded",
				Label: "TasksWhoseInputNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: TasksWhoseOutputNodeIsExpanded
			div := (&form.FormDiv{Name: "TasksWhoseOutputNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TasksWhoseOutputNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TasksWhoseOutputNodeIsExpanded",
				Label: "TasksWhoseOutputNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: TasksWhosePredecessorNodeIsExpanded
			div := (&form.FormDiv{Name: "TasksWhosePredecessorNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TasksWhosePredecessorNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TasksWhosePredecessorNodeIsExpanded",
				Label: "TasksWhosePredecessorNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsTaskGroupsNodeExpanded", inst.IsTaskGroupsNodeExpanded, probe.formStage, formGroup)

		{
			// Slice of pointers: TaskGroupShapes
			div := (&form.FormDiv{Name: "TaskGroupShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TaskGroupShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TaskGroupShapes",
				Label: "TaskGroupShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: TaskGroupsWhoseNodeIsExpanded
			div := (&form.FormDiv{Name: "TaskGroupsWhoseNodeIsExpanded"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TaskGroupsWhoseNodeIsExpanded {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TaskGroupsWhoseNodeIsExpanded",
				Label: "TaskGroupsWhoseNodeIsExpanded",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: TaskComposition_Shapes
			div := (&form.FormDiv{Name: "TaskComposition_Shapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TaskComposition_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TaskComposition_Shapes",
				Label: "TaskComposition_Shapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: TaskInputShapes
			div := (&form.FormDiv{Name: "TaskInputShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TaskInputShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TaskInputShapes",
				Label: "TaskInputShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: TaskOutputShapes
			div := (&form.FormDiv{Name: "TaskOutputShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TaskOutputShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TaskOutputShapes",
				Label: "TaskOutputShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: TaskPredecessorShapes
			div := (&form.FormDiv{Name: "TaskPredecessorShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TaskPredecessorShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TaskPredecessorShapes",
				Label: "TaskPredecessorShapes",
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
			// Slice of pointers: NoteProductShapes
			div := (&form.FormDiv{Name: "NoteProductShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.NoteProductShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "NoteProductShapes",
				Label: "NoteProductShapes",
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
			// Slice of pointers: Resource_Shapes
			div := (&form.FormDiv{Name: "Resource_Shapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Resource_Shapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Resource_Shapes",
				Label: "Resource_Shapes",
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
		StageSetBasicFieldtoForm("IsResourcesNodeExpanded", inst.IsResourcesNodeExpanded, probe.formStage, formGroup)

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
			// Slice of pointers: ResourceTaskShapes
			div := (&form.FormDiv{Name: "ResourceTaskShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ResourceTaskShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ResourceTaskShapes",
				Label: "ResourceTaskShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
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
	case *models.Library:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

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
		StageSetBasicFieldtoForm("LogoSVGFile", inst.LogoSVGFile, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsRootLibrary", inst.IsRootLibrary, probe.formStage, formGroup)

		{
			// Slice of pointers: RootProducts
			div := (&form.FormDiv{Name: "RootProducts"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RootProducts {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RootProducts",
				Label: "RootProducts",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: RootTasks
			div := (&form.FormDiv{Name: "RootTasks"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RootTasks {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RootTasks",
				Label: "RootTasks",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: RootTaskGroups
			div := (&form.FormDiv{Name: "RootTaskGroups"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RootTaskGroups {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RootTaskGroups",
				Label: "RootTaskGroups",
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
		StageSetEnumIntFieldToForm("LayoutDirection", inst.LayoutDirection, formGroup, probe.formStage)

		{
			// Slice of pointers: Products
			div := (&form.FormDiv{Name: "Products"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Products {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Products",
				Label: "Products",
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
			for src := range probe.stageSet.Stage.NoteProductShapes {
				if src.Note == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.NoteProductShape", "Note", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.NoteResourceShapes {
				if src.Note == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.NoteResourceShape", "Note", refNames, formGroup, probe.formStage)
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
	case *models.NoteProductShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Note", inst.Note, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Note](), probe.formStage)
		StageSetAssociationFieldToForm("Product", inst.Product, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Product](), probe.formStage)
		StageSetBasicFieldtoForm("StartRatio", inst.StartRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndRatio", inst.EndRatio, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("StartOrientation", inst.StartOrientation, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EndOrientation", inst.EndOrientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("CornerOffsetRatio", inst.CornerOffsetRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.NoteProductShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "NoteProductShapes", refNames, formGroup, probe.formStage)
		}
	case *models.NoteResourceShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Note", inst.Note, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Note](), probe.formStage)
		StageSetAssociationFieldToForm("Resource", inst.Resource, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Resource](), probe.formStage)
		StageSetBasicFieldtoForm("StartRatio", inst.StartRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndRatio", inst.EndRatio, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("StartOrientation", inst.StartOrientation, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EndOrientation", inst.EndOrientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("CornerOffsetRatio", inst.CornerOffsetRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

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
	case *models.NoteShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Note", inst.Note, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Note](), probe.formStage)
		StageSetBasicFieldtoForm("IsLayoutDirectionDifferent", inst.IsLayoutDirectionDifferent, probe.formStage, formGroup)
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
	case *models.Product:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)

		{
			// Slice of pointers: SubProducts
			div := (&form.FormDiv{Name: "SubProducts"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SubProducts {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SubProducts",
				Label: "SubProducts",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsProducersNodeExpanded", inst.IsProducersNodeExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsConsumersNodeExpanded", inst.IsConsumersNodeExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsImport", inst.IsImport, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ReferencedProduct", inst.ReferencedProduct, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Product](), probe.formStage)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetEnumIntFieldToForm("LayoutDirection", inst.LayoutDirection, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ProductsWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "ProductsWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.RootProducts {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "RootProducts", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Notes {
				for _, target := range src.Products {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Note", "Products", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.NoteProductShapes {
				if src.Product == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.NoteProductShape", "Product", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Products {
				for _, target := range src.SubProducts {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Product", "SubProducts", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Products {
				if src.ReferencedProduct == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Product", "ReferencedProduct", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ProductCompositionShapes {
				if src.Product == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ProductCompositionShape", "Product", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ProductReferenceShapes {
				if src.Product == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ProductReferenceShape", "Product", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ProductReferenceShapes {
				if src.ReferencedProduct == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ProductReferenceShape", "ReferencedProduct", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ProductShapes {
				if src.Product == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ProductShape", "Product", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Tasks {
				for _, target := range src.Inputs {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Task", "Inputs", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Tasks {
				for _, target := range src.Outputs {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Task", "Outputs", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.TaskInputShapes {
				if src.Product == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.TaskInputShape", "Product", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.TaskOutputShapes {
				if src.Product == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.TaskOutputShape", "Product", refNames, formGroup, probe.formStage)
		}
	case *models.ProductCompositionShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Product", inst.Product, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Product](), probe.formStage)
		StageSetBasicFieldtoForm("StartRatio", inst.StartRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndRatio", inst.EndRatio, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("StartOrientation", inst.StartOrientation, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EndOrientation", inst.EndOrientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("CornerOffsetRatio", inst.CornerOffsetRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ProductComposition_Shapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "ProductComposition_Shapes", refNames, formGroup, probe.formStage)
		}
	case *models.ProductReferenceShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Product", inst.Product, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Product](), probe.formStage)
		StageSetAssociationFieldToForm("ReferencedProduct", inst.ReferencedProduct, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Product](), probe.formStage)
		StageSetBasicFieldtoForm("StartRatio", inst.StartRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndRatio", inst.EndRatio, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("StartOrientation", inst.StartOrientation, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EndOrientation", inst.EndOrientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("CornerOffsetRatio", inst.CornerOffsetRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ProductReference_Shapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "ProductReference_Shapes", refNames, formGroup, probe.formStage)
		}
	case *models.ProductShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Product", inst.Product, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Product](), probe.formStage)
		StageSetBasicFieldtoForm("IsShowType", inst.IsShowType, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsLayoutDirectionDifferent", inst.IsLayoutDirectionDifferent, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.Product_Shapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "Product_Shapes", refNames, formGroup, probe.formStage)
		}
	case *models.Resource:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)

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
			// Slice of pointers: SubResources
			div := (&form.FormDiv{Name: "SubResources"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SubResources {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SubResources",
				Label: "SubResources",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetEnumIntFieldToForm("LayoutDirection", inst.LayoutDirection, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("IsImport", inst.IsImport, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ReferencedResource", inst.ReferencedResource, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Resource](), probe.formStage)

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
			for src := range probe.stageSet.Stage.NoteResourceShapes {
				if src.Resource == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.NoteResourceShape", "Resource", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Resources {
				for _, target := range src.SubResources {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Resource", "SubResources", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Resources {
				if src.ReferencedResource == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Resource", "ReferencedResource", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ResourceCompositionShapes {
				if src.Resource == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ResourceCompositionShape", "Resource", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ResourceShapes {
				if src.Resource == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ResourceShape", "Resource", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ResourceTaskShapes {
				if src.Resource == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ResourceTaskShape", "Resource", refNames, formGroup, probe.formStage)
		}
	case *models.ResourceCompositionShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Resource", inst.Resource, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Resource](), probe.formStage)
		StageSetBasicFieldtoForm("StartRatio", inst.StartRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndRatio", inst.EndRatio, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("StartOrientation", inst.StartOrientation, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EndOrientation", inst.EndOrientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("CornerOffsetRatio", inst.CornerOffsetRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

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
	case *models.ResourceShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Resource", inst.Resource, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Resource](), probe.formStage)
		StageSetBasicFieldtoForm("IsLayoutDirectionDifferent", inst.IsLayoutDirectionDifferent, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.Resource_Shapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "Resource_Shapes", refNames, formGroup, probe.formStage)
		}
	case *models.ResourceTaskShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Resource", inst.Resource, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Resource](), probe.formStage)
		StageSetAssociationFieldToForm("Task", inst.Task, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Task](), probe.formStage)
		StageSetBasicFieldtoForm("StartRatio", inst.StartRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndRatio", inst.EndRatio, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("StartOrientation", inst.StartOrientation, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EndOrientation", inst.EndOrientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("CornerOffsetRatio", inst.CornerOffsetRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.ResourceTaskShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "ResourceTaskShapes", refNames, formGroup, probe.formStage)
		}
	case *models.Task:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Start", inst.Start, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("End", inst.End, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsAllDay", inst.IsAllDay, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsMilestone", inst.IsMilestone, probe.formStage, formGroup)

		{
			// Slice of pointers: TaskGroups
			div := (&form.FormDiv{Name: "TaskGroups"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TaskGroups {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TaskGroups",
				Label: "TaskGroups",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Predecessors
			div := (&form.FormDiv{Name: "Predecessors"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Predecessors {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Predecessors",
				Label: "Predecessors",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetEnumStringFieldToForm("DependencyType", inst.DependencyType, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("DependencyDurationYears", inst.DependencyDurationYears, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DependencyDurationMonths", inst.DependencyDurationMonths, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DependencyDurationWeeks", inst.DependencyDurationWeeks, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DependencyDurationDays", inst.DependencyDurationDays, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DependencyDurationHours", inst.DependencyDurationHours, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DurationYears", inst.DurationYears, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DurationMonths", inst.DurationMonths, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DurationWeeks", inst.DurationWeeks, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DurationDays", inst.DurationDays, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DurationHours", inst.DurationHours, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsEndDateComputedFromDuration", inst.IsEndDateComputedFromDuration, probe.formStage, formGroup)

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

		{
			// Slice of pointers: SubTasks
			div := (&form.FormDiv{Name: "SubTasks"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SubTasks {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SubTasks",
				Label: "SubTasks",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsWithCompletion", inst.IsWithCompletion, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Completion", inst.Completion, formGroup, probe.formStage)

		{
			// Slice of pointers: TaskGroupsToDisplay
			div := (&form.FormDiv{Name: "TaskGroupsToDisplay"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TaskGroupsToDisplay {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TaskGroupsToDisplay",
				Label: "TaskGroupsToDisplay",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetEnumStringFieldToForm("TextPosition", inst.TextPosition, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("XOffset", inst.XOffset, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("YOffset", inst.YOffset, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsImport", inst.IsImport, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ReferencedTask", inst.ReferencedTask, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Task](), probe.formStage)
		StageSetBasicFieldtoForm("IsInputsNodeExpanded", inst.IsInputsNodeExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsOutputsNodeExpanded", inst.IsOutputsNodeExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetEnumIntFieldToForm("LayoutDirection", inst.LayoutDirection, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.TasksWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "TasksWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.TasksWhoseInputNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "TasksWhoseInputNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.TasksWhoseOutputNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "TasksWhoseOutputNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.TasksWhosePredecessorNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "TasksWhosePredecessorNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.RootTasks {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "RootTasks", refNames, formGroup, probe.formStage)
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
			for src := range probe.stageSet.Stage.Resources {
				for _, target := range src.Tasks {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Resource", "Tasks", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ResourceTaskShapes {
				if src.Task == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ResourceTaskShape", "Task", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Tasks {
				for _, target := range src.Predecessors {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Task", "Predecessors", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Tasks {
				for _, target := range src.SubTasks {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Task", "SubTasks", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Tasks {
				if src.ReferencedTask == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Task", "ReferencedTask", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.TaskCompositionShapes {
				if src.Task == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.TaskCompositionShape", "Task", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.TaskGroups {
				for _, target := range src.Tasks {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.TaskGroup", "Tasks", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.TaskInputShapes {
				if src.Task == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.TaskInputShape", "Task", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.TaskOutputShapes {
				if src.Task == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.TaskOutputShape", "Task", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.TaskPredecessorShapes {
				if src.Predecessor == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.TaskPredecessorShape", "Predecessor", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.TaskPredecessorShapes {
				if src.Task == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.TaskPredecessorShape", "Task", refNames, formGroup, probe.formStage)
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
	case *models.TaskCompositionShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Task", inst.Task, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Task](), probe.formStage)
		StageSetBasicFieldtoForm("StartRatio", inst.StartRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndRatio", inst.EndRatio, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("StartOrientation", inst.StartOrientation, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EndOrientation", inst.EndOrientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("CornerOffsetRatio", inst.CornerOffsetRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.TaskComposition_Shapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "TaskComposition_Shapes", refNames, formGroup, probe.formStage)
		}
	case *models.TaskGroup:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ComputedPrefix", inst.ComputedPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

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
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.TaskGroupsWhoseNodeIsExpanded {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "TaskGroupsWhoseNodeIsExpanded", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.RootTaskGroups {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Library", "RootTaskGroups", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Tasks {
				for _, target := range src.TaskGroups {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Task", "TaskGroups", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Tasks {
				for _, target := range src.TaskGroupsToDisplay {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Task", "TaskGroupsToDisplay", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.TaskGroupShapes {
				if src.TaskGroup == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.TaskGroupShape", "TaskGroup", refNames, formGroup, probe.formStage)
		}
	case *models.TaskGroupShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("TaskGroup", inst.TaskGroup, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.TaskGroup](), probe.formStage)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.TaskGroupShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "TaskGroupShapes", refNames, formGroup, probe.formStage)
		}
	case *models.TaskInputShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Product", inst.Product, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Product](), probe.formStage)
		StageSetAssociationFieldToForm("Task", inst.Task, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Task](), probe.formStage)
		StageSetBasicFieldtoForm("StartRatio", inst.StartRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndRatio", inst.EndRatio, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("StartOrientation", inst.StartOrientation, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EndOrientation", inst.EndOrientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("CornerOffsetRatio", inst.CornerOffsetRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.TaskInputShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "TaskInputShapes", refNames, formGroup, probe.formStage)
		}
	case *models.TaskOutputShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Task", inst.Task, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Task](), probe.formStage)
		StageSetAssociationFieldToForm("Product", inst.Product, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Product](), probe.formStage)
		StageSetBasicFieldtoForm("StartRatio", inst.StartRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndRatio", inst.EndRatio, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("StartOrientation", inst.StartOrientation, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EndOrientation", inst.EndOrientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("CornerOffsetRatio", inst.CornerOffsetRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.TaskOutputShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "TaskOutputShapes", refNames, formGroup, probe.formStage)
		}
	case *models.TaskPredecessorShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Predecessor", inst.Predecessor, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Task](), probe.formStage)
		StageSetAssociationFieldToForm("Task", inst.Task, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Task](), probe.formStage)
		StageSetBasicFieldtoForm("StartRatio", inst.StartRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EndRatio", inst.EndRatio, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("StartOrientation", inst.StartOrientation, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("EndOrientation", inst.EndOrientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("CornerOffsetRatio", inst.CornerOffsetRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.TaskPredecessorShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "TaskPredecessorShapes", refNames, formGroup, probe.formStage)
		}
	case *models.TaskShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Task", inst.Task, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Task](), probe.formStage)
		StageSetBasicFieldtoForm("IsShowDate", inst.IsShowDate, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("VerticalOffset", inst.VerticalOffset, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DisplayVerticalBar", inst.DisplayVerticalBar, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsLayoutDirectionDifferent", inst.IsLayoutDirectionDifferent, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Diagrams {
				for _, target := range src.Task_Shapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Diagram", "Task_Shapes", refNames, formGroup, probe.formStage)
		}
	default:
		_ = inst
	}
}
