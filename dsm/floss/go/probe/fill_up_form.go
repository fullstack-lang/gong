// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/floss/go/models"
)

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.CompareAnalysis:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("FromSystem", instanceWithInferedType.FromSystem, formGroup, probe)
		AssociationFieldToForm("ToSystem", instanceWithInferedType.ToSystem, formGroup, probe)
		BasicFieldtoForm("Mu", instanceWithInferedType.Mu, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Epsilon", instanceWithInferedType.Epsilon, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("DiagramFlossEquations", instanceWithInferedType, &instanceWithInferedType.DiagramFlossEquations, formGroup, probe)
		AssociationSliceToForm("DiagramFlossEquationsWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.DiagramFlossEquationsWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Library, *models.CompareAnalysis](
				"Library",
				"RootCompareAnalysis",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.CompareAnalysis {
					return owner.RootCompareAnalysis
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.CompareAnalysis](
				"Library",
				"CompareAnalysisWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.CompareAnalysis {
					return owner.CompareAnalysisWhoseNodeIsExpanded
				})
		}

	case *models.Complexity:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Strength", instanceWithInferedType.Strength, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0, false)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramFlossEquation, *models.Complexity](
				"DiagramFlossEquation",
				"ComplexitysWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramFlossEquation) []*models.Complexity {
					return owner.ComplexitysWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.Complexity](
				"Library",
				"RootComplexitys",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Complexity {
					return owner.RootComplexitys
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.Complexity](
				"Library",
				"ComplexitysWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Complexity {
					return owner.ComplexitysWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Note, *models.Complexity](
				"Note",
				"Complexities",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Note) []*models.Complexity {
					return owner.Complexities
				})
		}
		{
			AssociationReverseSliceToForm[*models.System, *models.Complexity](
				"System",
				"Complexities",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.System) []*models.Complexity {
					return owner.Complexities
				})
		}
		{
			AssociationReverseSliceToForm[*models.System, *models.Complexity](
				"System",
				"ComplexitysWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.System) []*models.Complexity {
					return owner.ComplexitysWhoseNodeIsExpanded
				})
		}

	case *models.DiagramFlossEquation:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0, false)
		BasicFieldtoForm("Scale", instanceWithInferedType.Scale, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsChecked", instanceWithInferedType.IsChecked, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsEditable_", instanceWithInferedType.IsEditable_, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsInDelta3ColumnsMode", instanceWithInferedType.IsInDelta3ColumnsMode, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("AreQuantitativeElementsVisible", instanceWithInferedType.AreQuantitativeElementsVisible, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("AreSubsystemsVisible", instanceWithInferedType.AreSubsystemsVisible, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DefaultBoxWidth", instanceWithInferedType.DefaultBoxWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DefaultBoxHeigth", instanceWithInferedType.DefaultBoxHeigth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Note_Shapes", instanceWithInferedType, &instanceWithInferedType.Note_Shapes, formGroup, probe)
		AssociationSliceToForm("NoteComplexityShapes", instanceWithInferedType, &instanceWithInferedType.NoteComplexityShapes, formGroup, probe)
		AssociationSliceToForm("NotePerformanceShapes", instanceWithInferedType, &instanceWithInferedType.NotePerformanceShapes, formGroup, probe)
		AssociationSliceToForm("NoteEffortShapes", instanceWithInferedType, &instanceWithInferedType.NoteEffortShapes, formGroup, probe)
		BasicFieldtoForm("IsNotesNodeExpanded", instanceWithInferedType.IsNotesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("NotesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.NotesWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsComplexitysNodeExpanded", instanceWithInferedType.IsComplexitysNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("ComplexitysWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ComplexitysWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsPerformancesNodeExpanded", instanceWithInferedType.IsPerformancesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("PerformancesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.PerformancesWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsEffortsNodeExpanded", instanceWithInferedType.IsEffortsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("EffortsWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.EffortsWhoseNodeIsExpanded, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.CompareAnalysis, *models.DiagramFlossEquation](
				"CompareAnalysis",
				"DiagramFlossEquations",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.CompareAnalysis) []*models.DiagramFlossEquation {
					return owner.DiagramFlossEquations
				})
		}
		{
			AssociationReverseSliceToForm[*models.CompareAnalysis, *models.DiagramFlossEquation](
				"CompareAnalysis",
				"DiagramFlossEquationsWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.CompareAnalysis) []*models.DiagramFlossEquation {
					return owner.DiagramFlossEquationsWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.System, *models.DiagramFlossEquation](
				"System",
				"DiagramFlossEquations",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.System) []*models.DiagramFlossEquation {
					return owner.DiagramFlossEquations
				})
		}
		{
			AssociationReverseSliceToForm[*models.System, *models.DiagramFlossEquation](
				"System",
				"DiagramFlossEquationsWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.System) []*models.DiagramFlossEquation {
					return owner.DiagramFlossEquationsWhoseNodeIsExpanded
				})
		}

	case *models.Effort:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Strength", instanceWithInferedType.Strength, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0, false)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramFlossEquation, *models.Effort](
				"DiagramFlossEquation",
				"EffortsWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramFlossEquation) []*models.Effort {
					return owner.EffortsWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.Effort](
				"Library",
				"RootEfforts",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Effort {
					return owner.RootEfforts
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.Effort](
				"Library",
				"EffortsWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Effort {
					return owner.EffortsWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Note, *models.Effort](
				"Note",
				"Efforts",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Note) []*models.Effort {
					return owner.Efforts
				})
		}
		{
			AssociationReverseSliceToForm[*models.System, *models.Effort](
				"System",
				"Efforts",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.System) []*models.Effort {
					return owner.Efforts
				})
		}
		{
			AssociationReverseSliceToForm[*models.System, *models.Effort](
				"System",
				"EffortsWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.System) []*models.Effort {
					return owner.EffortsWhoseNodeIsExpanded
				})
		}

	case *models.Library:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0, false)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("SubLibraries", instanceWithInferedType, &instanceWithInferedType.SubLibraries, formGroup, probe)
		AssociationSliceToForm("RootSystems", instanceWithInferedType, &instanceWithInferedType.RootSystems, formGroup, probe)
		AssociationSliceToForm("RootComplexitys", instanceWithInferedType, &instanceWithInferedType.RootComplexitys, formGroup, probe)
		AssociationSliceToForm("RootPerformances", instanceWithInferedType, &instanceWithInferedType.RootPerformances, formGroup, probe)
		AssociationSliceToForm("RootEfforts", instanceWithInferedType, &instanceWithInferedType.RootEfforts, formGroup, probe)
		AssociationSliceToForm("RootCompareAnalysis", instanceWithInferedType, &instanceWithInferedType.RootCompareAnalysis, formGroup, probe)
		AssociationSliceToForm("RootNotes", instanceWithInferedType, &instanceWithInferedType.RootNotes, formGroup, probe)
		BasicFieldtoForm("IsRootLibrary", instanceWithInferedType.IsRootLibrary, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsSubLibrariesNodeExpanded", instanceWithInferedType.IsSubLibrariesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("SubLibrariesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.SubLibrariesWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("NbPixPerCharacter", instanceWithInferedType.NbPixPerCharacter, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LogoSVGFile", instanceWithInferedType.LogoSVGFile, instanceWithInferedType, probe.formStage, formGroup,
			false, true, 600, true, 300, false)
		BasicFieldtoForm("IsSystemsNodeExpanded", instanceWithInferedType.IsSystemsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("SystemsWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.SystemsWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsComplexitysNodeExpanded", instanceWithInferedType.IsComplexitysNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("ComplexitysWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ComplexitysWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsPerformancesNodeExpanded", instanceWithInferedType.IsPerformancesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("PerformancesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.PerformancesWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsEffortsNodeExpanded", instanceWithInferedType.IsEffortsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("EffortsWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.EffortsWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsCompareAnalysisNodeExpanded", instanceWithInferedType.IsCompareAnalysisNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("CompareAnalysisWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.CompareAnalysisWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsNotesNodeExpanded", instanceWithInferedType.IsNotesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("NotesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.NotesWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsExpandedTmp", instanceWithInferedType.IsExpandedTmp, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
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
		{
			AssociationReverseSliceToForm[*models.Library, *models.Library](
				"Library",
				"SubLibrariesWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Library {
					return owner.SubLibrariesWhoseNodeIsExpanded
				})
		}

	case *models.Note:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0, false)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0, false)
		AssociationSliceToForm("Complexities", instanceWithInferedType, &instanceWithInferedType.Complexities, formGroup, probe)
		AssociationSliceToForm("Performances", instanceWithInferedType, &instanceWithInferedType.Performances, formGroup, probe)
		AssociationSliceToForm("Efforts", instanceWithInferedType, &instanceWithInferedType.Efforts, formGroup, probe)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsComplexitysNodeExpanded", instanceWithInferedType.IsComplexitysNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsPerformancesNodeExpanded", instanceWithInferedType.IsPerformancesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsEffortsNodeExpanded", instanceWithInferedType.IsEffortsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramFlossEquation, *models.Note](
				"DiagramFlossEquation",
				"NotesWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramFlossEquation) []*models.Note {
					return owner.NotesWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.Note](
				"Library",
				"RootNotes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Note {
					return owner.RootNotes
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.Note](
				"Library",
				"NotesWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Note {
					return owner.NotesWhoseNodeIsExpanded
				})
		}

	case *models.NoteComplexityShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Note", instanceWithInferedType.Note, formGroup, probe)
		AssociationFieldToForm("Complexity", instanceWithInferedType.Complexity, formGroup, probe)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramFlossEquation, *models.NoteComplexityShape](
				"DiagramFlossEquation",
				"NoteComplexityShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramFlossEquation) []*models.NoteComplexityShape {
					return owner.NoteComplexityShapes
				})
		}

	case *models.NoteEffortShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Note", instanceWithInferedType.Note, formGroup, probe)
		AssociationFieldToForm("Effort", instanceWithInferedType.Effort, formGroup, probe)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramFlossEquation, *models.NoteEffortShape](
				"DiagramFlossEquation",
				"NoteEffortShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramFlossEquation) []*models.NoteEffortShape {
					return owner.NoteEffortShapes
				})
		}

	case *models.NotePerformanceShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Note", instanceWithInferedType.Note, formGroup, probe)
		AssociationFieldToForm("Performance", instanceWithInferedType.Performance, formGroup, probe)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramFlossEquation, *models.NotePerformanceShape](
				"DiagramFlossEquation",
				"NotePerformanceShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramFlossEquation) []*models.NotePerformanceShape {
					return owner.NotePerformanceShapes
				})
		}

	case *models.NoteShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Note", instanceWithInferedType.Note, formGroup, probe)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHidden", instanceWithInferedType.IsHidden, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramFlossEquation, *models.NoteShape](
				"DiagramFlossEquation",
				"Note_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramFlossEquation) []*models.NoteShape {
					return owner.Note_Shapes
				})
		}

	case *models.Performance:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Strength", instanceWithInferedType.Strength, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0, false)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramFlossEquation, *models.Performance](
				"DiagramFlossEquation",
				"PerformancesWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramFlossEquation) []*models.Performance {
					return owner.PerformancesWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.Performance](
				"Library",
				"RootPerformances",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Performance {
					return owner.RootPerformances
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.Performance](
				"Library",
				"PerformancesWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Performance {
					return owner.PerformancesWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Note, *models.Performance](
				"Note",
				"Performances",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Note) []*models.Performance {
					return owner.Performances
				})
		}
		{
			AssociationReverseSliceToForm[*models.System, *models.Performance](
				"System",
				"Performances",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.System) []*models.Performance {
					return owner.Performances
				})
		}
		{
			AssociationReverseSliceToForm[*models.System, *models.Performance](
				"System",
				"PerformancesWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.System) []*models.Performance {
					return owner.PerformancesWhoseNodeIsExpanded
				})
		}

	case *models.System:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0, false)
		AssociationSliceToForm("Complexities", instanceWithInferedType, &instanceWithInferedType.Complexities, formGroup, probe)
		AssociationSliceToForm("Performances", instanceWithInferedType, &instanceWithInferedType.Performances, formGroup, probe)
		AssociationSliceToForm("Efforts", instanceWithInferedType, &instanceWithInferedType.Efforts, formGroup, probe)
		AssociationSliceToForm("SubSystems", instanceWithInferedType, &instanceWithInferedType.SubSystems, formGroup, probe)
		BasicFieldtoForm("AreCPEsCompoundedFromSubSystems", instanceWithInferedType.AreCPEsCompoundedFromSubSystems, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("SVG_Path", instanceWithInferedType.SVG_Path, instanceWithInferedType, probe.formStage, formGroup,
			false, true, 600, true, 300, false)
		BasicFieldtoForm("InverseAppliedScaling", instanceWithInferedType.InverseAppliedScaling, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("DiagramFlossEquations", instanceWithInferedType, &instanceWithInferedType.DiagramFlossEquations, formGroup, probe)
		AssociationSliceToForm("DiagramFlossEquationsWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.DiagramFlossEquationsWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsSubSystemNodeExpanded", instanceWithInferedType.IsSubSystemNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsComplexitysNodeExpanded", instanceWithInferedType.IsComplexitysNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("ComplexitysWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ComplexitysWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsPerformancesNodeExpanded", instanceWithInferedType.IsPerformancesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("PerformancesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.PerformancesWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsEffortsNodeExpanded", instanceWithInferedType.IsEffortsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("EffortsWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.EffortsWhoseNodeIsExpanded, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Library, *models.System](
				"Library",
				"RootSystems",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.System {
					return owner.RootSystems
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.System](
				"Library",
				"SystemsWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.System {
					return owner.SystemsWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.System, *models.System](
				"System",
				"SubSystems",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.System) []*models.System {
					return owner.SubSystems
				})
		}

	default:
		_ = instanceWithInferedType
	}
}
