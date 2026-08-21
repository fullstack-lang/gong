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
	case *models.DiagramFloss:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0, false)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsChecked", instanceWithInferedType.IsChecked, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsEditable_", instanceWithInferedType.IsEditable_, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsShowPrefix", instanceWithInferedType.IsShowPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DefaultBoxWidth", instanceWithInferedType.DefaultBoxWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DefaultBoxHeigth", instanceWithInferedType.DefaultBoxHeigth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("System_Shapes", instanceWithInferedType, &instanceWithInferedType.System_Shapes, formGroup, probe)
		BasicFieldtoForm("IsSystemsNodeExpanded", instanceWithInferedType.IsSystemsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("SystemsWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.SystemsWhoseNodeIsExpanded, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.System, *models.DiagramFloss](
				"System",
				"DiagramFlosses",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.System) []*models.DiagramFloss {
					return owner.DiagramFlosses
				})
		}
		{
			AssociationReverseSliceToForm[*models.System, *models.DiagramFloss](
				"System",
				"DiagramFlossWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.System) []*models.DiagramFloss {
					return owner.DiagramFlossWhoseNodeIsExpanded
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
		BasicFieldtoForm("IsRootLibrary", instanceWithInferedType.IsRootLibrary, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("SubLibraries", instanceWithInferedType, &instanceWithInferedType.SubLibraries, formGroup, probe)
		BasicFieldtoForm("IsSubLibrariesNodeExpanded", instanceWithInferedType.IsSubLibrariesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("SubLibrariesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.SubLibrariesWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("NbPixPerCharacter", instanceWithInferedType.NbPixPerCharacter, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LogoSVGFile", instanceWithInferedType.LogoSVGFile, instanceWithInferedType, probe.formStage, formGroup,
			false, true, 600, true, 300, false)
		AssociationSliceToForm("RootSystemes", instanceWithInferedType, &instanceWithInferedType.RootSystemes, formGroup, probe)
		BasicFieldtoForm("IsSystemesNodeExpanded", instanceWithInferedType.IsSystemesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("SystemsWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.SystemsWhoseNodeIsExpanded, formGroup, probe)
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

	case *models.System:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0, false)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("SVG_Path", instanceWithInferedType.SVG_Path, instanceWithInferedType, probe.formStage, formGroup,
			false, true, 600, true, 300, false)
		BasicFieldtoForm("InverseAppliedScaling", instanceWithInferedType.InverseAppliedScaling, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("DiagramFlosses", instanceWithInferedType, &instanceWithInferedType.DiagramFlosses, formGroup, probe)
		AssociationSliceToForm("DiagramFlossWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.DiagramFlossWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsSubSystemNodeExpanded", instanceWithInferedType.IsSubSystemNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("SubSystemes", instanceWithInferedType, &instanceWithInferedType.SubSystemes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramFloss, *models.System](
				"DiagramFloss",
				"SystemsWhoseNodeIsExpanded",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramFloss) []*models.System {
					return owner.SystemsWhoseNodeIsExpanded
				})
		}
		{
			AssociationReverseSliceToForm[*models.Library, *models.System](
				"Library",
				"RootSystemes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.System {
					return owner.RootSystemes
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
				"SubSystemes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.System) []*models.System {
					return owner.SubSystemes
				})
		}

	case *models.SystemShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("System", instanceWithInferedType.System, formGroup, probe)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
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
			AssociationReverseSliceToForm[*models.DiagramFloss, *models.SystemShape](
				"DiagramFloss",
				"System_Shapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramFloss) []*models.SystemShape {
					return owner.System_Shapes
				})
		}

	default:
		_ = instanceWithInferedType
	}
}
