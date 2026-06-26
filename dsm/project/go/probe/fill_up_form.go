// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/project/go/models"
)

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
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
		BasicFieldtoForm("IsShowPrefix", instanceWithInferedType.IsShowPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DefaultBoxWidth", instanceWithInferedType.DefaultBoxWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DefaultBoxHeigth", instanceWithInferedType.DefaultBoxHeigth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Product_Shapes", instanceWithInferedType, &instanceWithInferedType.Product_Shapes, formGroup, probe)
		AssociationSliceToForm("ProductsWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ProductsWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsPBSNodeExpanded", instanceWithInferedType.IsPBSNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ProductComposition_Shapes", instanceWithInferedType, &instanceWithInferedType.ProductComposition_Shapes, formGroup, probe)
		BasicFieldtoForm("IsWBSNodeExpanded", instanceWithInferedType.IsWBSNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Task_Shapes", instanceWithInferedType, &instanceWithInferedType.Task_Shapes, formGroup, probe)
		AssociationSliceToForm("TasksWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.TasksWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("TasksWhoseInputNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.TasksWhoseInputNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("TasksWhoseOutputNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.TasksWhoseOutputNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsTaskGroupsNodeExpanded", instanceWithInferedType.IsTaskGroupsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("TaskGroupShapes", instanceWithInferedType, &instanceWithInferedType.TaskGroupShapes, formGroup, probe)
		AssociationSliceToForm("TaskGroupsWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.TaskGroupsWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("DateFormat", instanceWithInferedType.DateFormat, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("TaskComposition_Shapes", instanceWithInferedType, &instanceWithInferedType.TaskComposition_Shapes, formGroup, probe)
		AssociationSliceToForm("TaskInputShapes", instanceWithInferedType, &instanceWithInferedType.TaskInputShapes, formGroup, probe)
		AssociationSliceToForm("TaskOutputShapes", instanceWithInferedType, &instanceWithInferedType.TaskOutputShapes, formGroup, probe)
		AssociationSliceToForm("Note_Shapes", instanceWithInferedType, &instanceWithInferedType.Note_Shapes, formGroup, probe)
		AssociationSliceToForm("NotesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.NotesWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsNotesNodeExpanded", instanceWithInferedType.IsNotesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("NoteProductShapes", instanceWithInferedType, &instanceWithInferedType.NoteProductShapes, formGroup, probe)
		AssociationSliceToForm("NoteTaskShapes", instanceWithInferedType, &instanceWithInferedType.NoteTaskShapes, formGroup, probe)
		AssociationSliceToForm("NoteResourceShapes", instanceWithInferedType, &instanceWithInferedType.NoteResourceShapes, formGroup, probe)
		AssociationSliceToForm("Resource_Shapes", instanceWithInferedType, &instanceWithInferedType.Resource_Shapes, formGroup, probe)
		AssociationSliceToForm("ResourcesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ResourcesWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsResourcesNodeExpanded", instanceWithInferedType.IsResourcesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("ResourceComposition_Shapes", instanceWithInferedType, &instanceWithInferedType.ResourceComposition_Shapes, formGroup, probe)
		AssociationSliceToForm("ResourceTaskShapes", instanceWithInferedType, &instanceWithInferedType.ResourceTaskShapes, formGroup, probe)
		BasicFieldtoForm("IsTimeDiagram", instanceWithInferedType.IsTimeDiagram, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedStart", instanceWithInferedType.ComputedStart, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedEnd", instanceWithInferedType.ComputedEnd, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedDuration", instanceWithInferedType.ComputedDuration, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("UseManualStartAndEndDates", instanceWithInferedType.UseManualStartAndEndDates, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ManualStart", instanceWithInferedType.ManualStart, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ManualEnd", instanceWithInferedType.ManualEnd, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("TimeStep", instanceWithInferedType.TimeStep, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("TimeStepScale", instanceWithInferedType.TimeStepScale, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("LaneHeight", instanceWithInferedType.LaneHeight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("RatioBarToLaneHeight", instanceWithInferedType.RatioBarToLaneHeight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("YTopMargin", instanceWithInferedType.YTopMargin, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("XLeftText", instanceWithInferedType.XLeftText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("TextHeight", instanceWithInferedType.TextHeight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("XLeftLanes", instanceWithInferedType.XLeftLanes, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("XRightMargin", instanceWithInferedType.XRightMargin, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ArrowLengthToTheRightOfStartBar", instanceWithInferedType.ArrowLengthToTheRightOfStartBar, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ArrowTipLenght", instanceWithInferedType.ArrowTipLenght, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("TimeLine_Color", instanceWithInferedType.TimeLine_Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("TimeLine_FillOpacity", instanceWithInferedType.TimeLine_FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("TimeLine_Stroke", instanceWithInferedType.TimeLine_Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("TimeLine_StrokeWidth", instanceWithInferedType.TimeLine_StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DrawVerticalTimeLines", instanceWithInferedType.DrawVerticalTimeLines, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Group_Stroke", instanceWithInferedType.Group_Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Group_StrokeWidth", instanceWithInferedType.Group_StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Group_StrokeDashArray", instanceWithInferedType.Group_StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DateYOffset", instanceWithInferedType.DateYOffset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("AlignOnStartEndOnYearStart", instanceWithInferedType.AlignOnStartEndOnYearStart, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
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
		AssociationSliceToForm("SubLibraries", instanceWithInferedType, &instanceWithInferedType.SubLibraries, formGroup, probe)
		BasicFieldtoForm("NbPixPerCharacter", instanceWithInferedType.NbPixPerCharacter, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LogoSVGFile", instanceWithInferedType.LogoSVGFile, instanceWithInferedType, probe.formStage, formGroup,
			false, true, 600, true, 300)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("IsRootLibrary", instanceWithInferedType.IsRootLibrary, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("RootProducts", instanceWithInferedType, &instanceWithInferedType.RootProducts, formGroup, probe)
		AssociationSliceToForm("RootTasks", instanceWithInferedType, &instanceWithInferedType.RootTasks, formGroup, probe)
		AssociationSliceToForm("RootTaskGroups", instanceWithInferedType, &instanceWithInferedType.RootTaskGroups, formGroup, probe)
		AssociationSliceToForm("RootResources", instanceWithInferedType, &instanceWithInferedType.RootResources, formGroup, probe)
		AssociationSliceToForm("Notes", instanceWithInferedType, &instanceWithInferedType.Notes, formGroup, probe)
		AssociationSliceToForm("Diagrams", instanceWithInferedType, &instanceWithInferedType.Diagrams, formGroup, probe)
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
		AssociationSliceToForm("Products", instanceWithInferedType, &instanceWithInferedType.Products, formGroup, probe)
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

	case *models.NoteProductShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Note", instanceWithInferedType.Note, formGroup, probe)
		AssociationFieldToForm("Product", instanceWithInferedType.Product, formGroup, probe)
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
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.NoteProductShape](
				"Diagram",
				"NoteProductShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.NoteProductShape {
					return owner.NoteProductShapes
				})
		}

	case *models.NoteResourceShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Note", instanceWithInferedType.Note, formGroup, probe)
		AssociationFieldToForm("Resource", instanceWithInferedType.Resource, formGroup, probe)
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
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.NoteResourceShape](
				"Diagram",
				"NoteResourceShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.NoteResourceShape {
					return owner.NoteResourceShapes
				})
		}

	case *models.NoteShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Note", instanceWithInferedType.Note, formGroup, probe)
		BasicFieldtoForm("OverideLayoutDirection", instanceWithInferedType.OverideLayoutDirection, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
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

	case *models.Product:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("IsImport", instanceWithInferedType.IsImport, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ReferencedProduct", instanceWithInferedType.ReferencedProduct, formGroup, probe)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		AssociationSliceToForm("SubProducts", instanceWithInferedType, &instanceWithInferedType.SubProducts, formGroup, probe)
		BasicFieldtoForm("IsProducersNodeExpanded", instanceWithInferedType.IsProducersNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsConsumersNodeExpanded", instanceWithInferedType.IsConsumersNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.Product](
				"Diagram",
				"ProductsWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.Product {
					return owner.ProductsWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.Product](
				"Library",
				"RootProducts",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Product {
					return owner.RootProducts
				})
		}
		{
			AssociationReverseSliceToForm[*models.Note, *models.Product](
				"Note",
				"Products",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Note) []*models.Product {
					return owner.Products
				})
		}
		{
			AssociationReverseSliceToForm[*models.Product, *models.Product](
				"Product",
				"SubProducts",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Product) []*models.Product {
					return owner.SubProducts
				})
		}
		{
			AssociationReverseSliceToForm[*models.Task, *models.Product](
				"Task",
				"Inputs",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Task) []*models.Product {
					return owner.Inputs
				})
		}
		{
			AssociationReverseSliceToForm[*models.Task, *models.Product](
				"Task",
				"Outputs",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Task) []*models.Product {
					return owner.Outputs
				})
		}

	case *models.ProductCompositionShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Product", instanceWithInferedType.Product, formGroup, probe)
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
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.ProductCompositionShape](
				"Diagram",
				"ProductComposition_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.ProductCompositionShape {
					return owner.ProductComposition_Shapes
				})
		}

	case *models.ProductShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Product", instanceWithInferedType.Product, formGroup, probe)
		BasicFieldtoForm("OverideLayoutDirection", instanceWithInferedType.OverideLayoutDirection, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
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
			AssociationReverseSliceToForm[*models.Diagram, *models.ProductShape](
				"Diagram",
				"Product_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.ProductShape {
					return owner.Product_Shapes
				})
		}

	case *models.Resource:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("IsImport", instanceWithInferedType.IsImport, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ReferencedResource", instanceWithInferedType.ReferencedResource, formGroup, probe)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		AssociationSliceToForm("Tasks", instanceWithInferedType, &instanceWithInferedType.Tasks, formGroup, probe)
		AssociationSliceToForm("SubResources", instanceWithInferedType, &instanceWithInferedType.SubResources, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.Resource](
				"Diagram",
				"ResourcesWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.Resource {
					return owner.ResourcesWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.Resource](
				"Library",
				"RootResources",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Resource {
					return owner.RootResources
				})
		}
		{
			AssociationReverseSliceToForm[*models.Note, *models.Resource](
				"Note",
				"Resources",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Note) []*models.Resource {
					return owner.Resources
				})
		}
		{
			AssociationReverseSliceToForm[*models.Resource, *models.Resource](
				"Resource",
				"SubResources",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Resource) []*models.Resource {
					return owner.SubResources
				})
		}

	case *models.ResourceCompositionShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Resource", instanceWithInferedType.Resource, formGroup, probe)
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
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.ResourceCompositionShape](
				"Diagram",
				"ResourceComposition_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.ResourceCompositionShape {
					return owner.ResourceComposition_Shapes
				})
		}

	case *models.ResourceShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Resource", instanceWithInferedType.Resource, formGroup, probe)
		BasicFieldtoForm("OverideLayoutDirection", instanceWithInferedType.OverideLayoutDirection, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
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
			AssociationReverseSliceToForm[*models.Diagram, *models.ResourceShape](
				"Diagram",
				"Resource_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.ResourceShape {
					return owner.Resource_Shapes
				})
		}

	case *models.ResourceTaskShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Resource", instanceWithInferedType.Resource, formGroup, probe)
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
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.ResourceTaskShape](
				"Diagram",
				"ResourceTaskShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.ResourceTaskShape {
					return owner.ResourceTaskShapes
				})
		}

	case *models.Task:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		AssociationSliceToForm("SubTasks", instanceWithInferedType, &instanceWithInferedType.SubTasks, formGroup, probe)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("IsImport", instanceWithInferedType.IsImport, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ReferencedTask", instanceWithInferedType.ReferencedTask, formGroup, probe)
		BasicFieldtoForm("Start", instanceWithInferedType.Start, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("End", instanceWithInferedType.End, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formGroup.FormDivs = append(formGroup.FormDivs, (&form.FormDiv{
			Name:       "Duration",
			IsAStartAccordionGroup: true,
			AccordionGroupName: "Duration",
		}).Stage(probe.formStage))
		BasicFieldtoForm("DurationYears", instanceWithInferedType.DurationYears, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DurationMonths", instanceWithInferedType.DurationMonths, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DurationWeeks", instanceWithInferedType.DurationWeeks, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DurationDays", instanceWithInferedType.DurationDays, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DurationHours", instanceWithInferedType.DurationHours, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsEndDateComputedFromDuration", instanceWithInferedType.IsEndDateComputedFromDuration, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formGroup.FormDivs = append(formGroup.FormDivs, (&form.FormDiv{
			Name:       "",
			IsAEndAccordionGroup:   true,
		}).Stage(probe.formStage))
		formGroup.FormDivs = append(formGroup.FormDivs, (&form.FormDiv{
			Name:       "Predecessors",
			IsAStartAccordionGroup: true,
			AccordionGroupName: "Predecessors",
		}).Stage(probe.formStage))
		AssociationSliceToForm("Predecessors", instanceWithInferedType, &instanceWithInferedType.Predecessors, formGroup, probe)
		BasicFieldtoForm("IsStartDateComputedFromPredecessors", instanceWithInferedType.IsStartDateComputedFromPredecessors, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formGroup.FormDivs = append(formGroup.FormDivs, (&form.FormDiv{
			Name:       "",
			IsAEndAccordionGroup:   true,
		}).Stage(probe.formStage))
		BasicFieldtoForm("IsMilestone", instanceWithInferedType.IsMilestone, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Inputs", instanceWithInferedType, &instanceWithInferedType.Inputs, formGroup, probe)
		BasicFieldtoForm("IsInputsNodeExpanded", instanceWithInferedType.IsInputsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Outputs", instanceWithInferedType, &instanceWithInferedType.Outputs, formGroup, probe)
		BasicFieldtoForm("IsOutputsNodeExpanded", instanceWithInferedType.IsOutputsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formGroup.FormDivs = append(formGroup.FormDivs, (&form.FormDiv{
			Name:       "Completion Display",
			IsAStartAccordionGroup: true,
			AccordionGroupName: "Completion Display",
		}).Stage(probe.formStage))
		BasicFieldtoForm("IsWithCompletion", instanceWithInferedType.IsWithCompletion, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Completion", instanceWithInferedType.Completion, instanceWithInferedType, probe.formStage, formGroup)
		formGroup.FormDivs = append(formGroup.FormDivs, (&form.FormDiv{
			Name:       "",
			IsAEndAccordionGroup:   true,
		}).Stage(probe.formStage))
		BasicFieldtoForm("DisplayVerticalBar", instanceWithInferedType.DisplayVerticalBar, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("TaskGroupsToDisplay", instanceWithInferedType, &instanceWithInferedType.TaskGroupsToDisplay, formGroup, probe)
		formGroup.FormDivs = append(formGroup.FormDivs, (&form.FormDiv{
			Name:       "Custom positions",
			IsAStartAccordionGroup: true,
			AccordionGroupName: "Custom positions",
		}).Stage(probe.formStage))
		EnumTypeStringToForm("TextPosition", instanceWithInferedType.TextPosition, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("XOffset", instanceWithInferedType.XOffset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("YOffset", instanceWithInferedType.YOffset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formGroup.FormDivs = append(formGroup.FormDivs, (&form.FormDiv{
			Name:       "",
			IsAEndAccordionGroup:   true,
		}).Stage(probe.formStage))
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.Task](
				"Diagram",
				"TasksWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.Task {
					return owner.TasksWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.Task](
				"Diagram",
				"TasksWhoseInputNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.Task {
					return owner.TasksWhoseInputNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.Task](
				"Diagram",
				"TasksWhoseOutputNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.Task {
					return owner.TasksWhoseOutputNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.Task](
				"Library",
				"RootTasks",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Task {
					return owner.RootTasks
				})
		}
		{
			AssociationReverseSliceToForm[*models.Note, *models.Task](
				"Note",
				"Tasks",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Note) []*models.Task {
					return owner.Tasks
				})
		}
		{
			AssociationReverseSliceToForm[*models.Resource, *models.Task](
				"Resource",
				"Tasks",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Resource) []*models.Task {
					return owner.Tasks
				})
		}
		{
			AssociationReverseSliceToForm[*models.Task, *models.Task](
				"Task",
				"SubTasks",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Task) []*models.Task {
					return owner.SubTasks
				})
		}
		{
			AssociationReverseSliceToForm[*models.Task, *models.Task](
				"Task",
				"Predecessors",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Task) []*models.Task {
					return owner.Predecessors
				})
		}
		{
			AssociationReverseSliceToForm[*models.TaskGroup, *models.Task](
				"TaskGroup",
				"Tasks",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.TaskGroup) []*models.Task {
					return owner.Tasks
				})
		}

	case *models.TaskCompositionShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
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
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.TaskCompositionShape](
				"Diagram",
				"TaskComposition_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.TaskCompositionShape {
					return owner.TaskComposition_Shapes
				})
		}

	case *models.TaskGroup:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		AssociationSliceToForm("Tasks", instanceWithInferedType, &instanceWithInferedType.Tasks, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.TaskGroup](
				"Diagram",
				"TaskGroupsWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.TaskGroup {
					return owner.TaskGroupsWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.TaskGroup](
				"Library",
				"RootTaskGroups",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.TaskGroup {
					return owner.RootTaskGroups
				})
		}
		{
			AssociationReverseSliceToForm[*models.Task, *models.TaskGroup](
				"Task",
				"TaskGroupsToDisplay",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Task) []*models.TaskGroup {
					return owner.TaskGroupsToDisplay
				})
		}

	case *models.TaskGroupShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("TaskGroup", instanceWithInferedType.TaskGroup, formGroup, probe)
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
			AssociationReverseSliceToForm[*models.Diagram, *models.TaskGroupShape](
				"Diagram",
				"TaskGroupShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.TaskGroupShape {
					return owner.TaskGroupShapes
				})
		}

	case *models.TaskInputShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Product", instanceWithInferedType.Product, formGroup, probe)
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
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.TaskInputShape](
				"Diagram",
				"TaskInputShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.TaskInputShape {
					return owner.TaskInputShapes
				})
		}

	case *models.TaskOutputShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Task", instanceWithInferedType.Task, formGroup, probe)
		AssociationFieldToForm("Product", instanceWithInferedType.Product, formGroup, probe)
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
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Diagram, *models.TaskOutputShape](
				"Diagram",
				"TaskOutputShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.TaskOutputShape {
					return owner.TaskOutputShapes
				})
		}

	case *models.TaskShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Task", instanceWithInferedType.Task, formGroup, probe)
		BasicFieldtoForm("IsShowDate", instanceWithInferedType.IsShowDate, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("OverideLayoutDirection", instanceWithInferedType.OverideLayoutDirection, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
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
			AssociationReverseSliceToForm[*models.Diagram, *models.TaskShape](
				"Diagram",
				"Task_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Diagram) []*models.TaskShape {
					return owner.Task_Shapes
				})
		}

	default:
		_ = instanceWithInferedType
	}
}
