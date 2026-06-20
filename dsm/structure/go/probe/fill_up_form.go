// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/structure/go/models"
)

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.DiagramStructure:
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
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DefaultBoxWidth", instanceWithInferedType.DefaultBoxWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("DefaultBoxHeigth", instanceWithInferedType.DefaultBoxHeigth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("System_Shapes", instanceWithInferedType, &instanceWithInferedType.System_Shapes, formGroup, probe)
		AssociationSliceToForm("Part_Shapes", instanceWithInferedType, &instanceWithInferedType.Part_Shapes, formGroup, probe)
		BasicFieldtoForm("IsPartsNodeExpanded", instanceWithInferedType.IsPartsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("PartsWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.PartsWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("Link_Shapes", instanceWithInferedType, &instanceWithInferedType.Link_Shapes, formGroup, probe)
		BasicFieldtoForm("IsLinksNodeExpanded", instanceWithInferedType.IsLinksNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("LinksWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.LinksWhoseNodeIsExpanded, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.System, *models.DiagramStructure](
				"System",
				"DiagramStructures",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.System) []*models.DiagramStructure {
					return owner.DiagramStructures
				})
		}
		{
			AssociationReverseSliceToForm[*models.System, *models.DiagramStructure](
				"System",
				"DiagramStructuresWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.System) []*models.DiagramStructure {
					return owner.DiagramStructuresWhoseNodeIsExpanded
				})
		}

	case *models.Library:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("SubLibraries", instanceWithInferedType, &instanceWithInferedType.SubLibraries, formGroup, probe)
		BasicFieldtoForm("IsSubLibrariesNodeExpanded", instanceWithInferedType.IsSubLibrariesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("SubLibrariesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.SubLibrariesWhoseNodeIsExpanded, formGroup, probe)
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
		AssociationSliceToForm("RootSystems", instanceWithInferedType, &instanceWithInferedType.RootSystems, formGroup, probe)
		BasicFieldtoForm("IsSystemsNodeExpanded", instanceWithInferedType.IsSystemsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("SystemsWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.SystemsWhoseNodeIsExpanded, formGroup, probe)
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

	case *models.Link:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		AssociationFieldToForm("Source", instanceWithInferedType.Source, formGroup, probe)
		AssociationFieldToForm("Target", instanceWithInferedType.Target, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramStructure, *models.Link](
				"DiagramStructure",
				"LinksWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramStructure) []*models.Link {
					return owner.LinksWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.System, *models.Link](
				"System",
				"Links",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.System) []*models.Link {
					return owner.Links
				})
		}
		{
			AssociationReverseSliceToForm[*models.System, *models.Link](
				"System",
				"LinksWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.System) []*models.Link {
					return owner.LinksWhoseNodeIsExpanded
				})
		}

	case *models.LinkAssociationShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Link", instanceWithInferedType.Link, formGroup, probe)
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
			AssociationReverseSliceToForm[*models.DiagramStructure, *models.LinkAssociationShape](
				"DiagramStructure",
				"Link_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramStructure) []*models.LinkAssociationShape {
					return owner.Link_Shapes
				})
		}

	case *models.Part:
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
			AssociationReverseSliceToForm[*models.DiagramStructure, *models.Part](
				"DiagramStructure",
				"PartsWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramStructure) []*models.Part {
					return owner.PartsWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.System, *models.Part](
				"System",
				"Parts",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.System) []*models.Part {
					return owner.Parts
				})
		}
		{
			AssociationReverseSliceToForm[*models.System, *models.Part](
				"System",
				"PartsWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.System) []*models.Part {
					return owner.PartsWhoseNodeIsExpanded
				})
		}

	case *models.PartShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Part", instanceWithInferedType.Part, formGroup, probe)
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
		BasicFieldtoForm("WidthWeight", instanceWithInferedType.WidthWeight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("OverideLayoutDirection", instanceWithInferedType.OverideLayoutDirection, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramStructure, *models.PartShape](
				"DiagramStructure",
				"Part_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramStructure) []*models.PartShape {
					return owner.Part_Shapes
				})
		}

	case *models.System:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		AssociationSliceToForm("Parts", instanceWithInferedType, &instanceWithInferedType.Parts, formGroup, probe)
		BasicFieldtoForm("IsPartsNodeExpanded", instanceWithInferedType.IsPartsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("PartsWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.PartsWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("SubSystems", instanceWithInferedType, &instanceWithInferedType.SubSystems, formGroup, probe)
		BasicFieldtoForm("IsSubSystemsNodeExpanded", instanceWithInferedType.IsSubSystemsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("SubSystemsWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.SubSystemsWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("Links", instanceWithInferedType, &instanceWithInferedType.Links, formGroup, probe)
		BasicFieldtoForm("IsLinksNodeExpanded", instanceWithInferedType.IsLinksNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("LinksWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.LinksWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("DiagramStructures", instanceWithInferedType, &instanceWithInferedType.DiagramStructures, formGroup, probe)
		BasicFieldtoForm("IsDiagramStructuresNodeExpanded", instanceWithInferedType.IsDiagramStructuresNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("DiagramStructuresWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.DiagramStructuresWhoseNodeIsExpanded, formGroup, probe)
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
		{
			AssociationReverseSliceToForm[*models.System, *models.System](
				"System",
				"SubSystemsWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.System) []*models.System {
					return owner.SubSystemsWhoseNodeIsExpanded
				})
		}

	case *models.SystemShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("System", instanceWithInferedType.System, formGroup, probe)
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
		BasicFieldtoForm("OverideLayoutDirection", instanceWithInferedType.OverideLayoutDirection, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeIntToForm("LayoutDirection", instanceWithInferedType.LayoutDirection, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramStructure, *models.SystemShape](
				"DiagramStructure",
				"System_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramStructure) []*models.SystemShape {
					return owner.System_Shapes
				})
		}

	default:
		_ = instanceWithInferedType
	}
}
