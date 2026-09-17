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
	case *models.AllocatedResourceShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Part", instanceWithInferedType.Part, formGroup, probe)
		AssociationFieldToForm("Resource", instanceWithInferedType.Resource, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		AssociationReverseSliceToForm(
			"DiagramStructure",
			"AllocatedResourceShapes",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.DiagramStructure) []*models.AllocatedResourceShape {
				return owner.AllocatedResourceShapes
			})

	case *models.AllocatedSystemShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Part", instanceWithInferedType.Part, formGroup, probe)
		AssociationFieldToForm("System", instanceWithInferedType.System, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		AssociationReverseSliceToForm(
			"DiagramStructure",
			"AllocatedSystemShapes",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.DiagramStructure) []*models.AllocatedSystemShape {
				return owner.AllocatedSystemShapes
			})

	case *models.ControlFlow:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0, false)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Start", instanceWithInferedType.Start, formGroup, probe)
		AssociationFieldToForm("End", instanceWithInferedType.End, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		AssociationReverseSliceToForm(
			"DiagramStructure",
			"ControlFlowsWhoseNodeIsExpanded",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.DiagramStructure) []*models.ControlFlow {
				return owner.ControlFlowsWhoseNodeIsExpanded
			})
		AssociationReverseSliceToForm(
			"Part",
			"ControlFlows",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.Part) []*models.ControlFlow {
				return owner.ControlFlows
			})

	case *models.ControlFlowShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("ControlFlow", instanceWithInferedType.ControlFlow, formGroup, probe)
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
		AssociationReverseSliceToForm(
			"DiagramStructure",
			"ControlFlow_Shapes",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.DiagramStructure) []*models.ControlFlowShape {
				return owner.ControlFlow_Shapes
			})

	case *models.Data:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Acronym", instanceWithInferedType.Acronym, instanceWithInferedType, probe.formStage, formGroup,
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
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		AssociationReverseSliceToForm(
			"DataFlow",
			"Datas",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.DataFlow) []*models.Data {
				return owner.Datas
			})
		AssociationReverseSliceToForm(
			"DiagramStructure",
			"DatasWhoseNodeIsExpanded",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.DiagramStructure) []*models.Data {
				return owner.DatasWhoseNodeIsExpanded
			})
		AssociationReverseSliceToForm(
			"Library",
			"RootDatas",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.Library) []*models.Data {
				return owner.RootDatas
			})
		AssociationReverseSliceToForm(
			"Library",
			"DatasWhoseNodeIsExpanded",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.Library) []*models.Data {
				return owner.DatasWhoseNodeIsExpanded
			})

	case *models.DataFlow:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("StartPort", instanceWithInferedType.StartPort, formGroup, probe)
		AssociationFieldToForm("EndPort", instanceWithInferedType.EndPort, formGroup, probe)
		AssociationFieldToForm("StartExternalPart", instanceWithInferedType.StartExternalPart, formGroup, probe)
		AssociationFieldToForm("EndExternalPart", instanceWithInferedType.EndExternalPart, formGroup, probe)
		AssociationSliceToForm("Datas", instanceWithInferedType, &instanceWithInferedType.Datas, formGroup, probe)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0, false)
		EnumTypeStringToForm("Direction", instanceWithInferedType.Direction, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("IsDatasNodeExpanded", instanceWithInferedType.IsDatasNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		AssociationReverseSliceToForm(
			"DiagramStructure",
			"DataFlowsWhoseNodeIsExpanded",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.DiagramStructure) []*models.DataFlow {
				return owner.DataFlowsWhoseNodeIsExpanded
			})
		AssociationReverseSliceToForm(
			"DiagramStructure",
			"DataFlowsWhoseDataNodeIsExpanded",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.DiagramStructure) []*models.DataFlow {
				return owner.DataFlowsWhoseDataNodeIsExpanded
			})
		AssociationReverseSliceToForm(
			"Library",
			"RootDataFlows",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.Library) []*models.DataFlow {
				return owner.RootDataFlows
			})
		AssociationReverseSliceToForm(
			"Library",
			"DataFlowsWhoseNodeIsExpanded",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.Library) []*models.DataFlow {
				return owner.DataFlowsWhoseNodeIsExpanded
			})
		AssociationReverseSliceToForm(
			"System",
			"DataFlows",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.System) []*models.DataFlow {
				return owner.DataFlows
			})

	case *models.DataFlowShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("DataFlow", instanceWithInferedType.DataFlow, formGroup, probe)
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
		AssociationReverseSliceToForm(
			"DiagramStructure",
			"DataFlow_Shapes",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.DiagramStructure) []*models.DataFlowShape {
				return owner.DataFlow_Shapes
			})

	case *models.DataShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Data", instanceWithInferedType.Data, formGroup, probe)
		AssociationFieldToForm("DataFlow", instanceWithInferedType.DataFlow, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		AssociationReverseSliceToForm(
			"DiagramStructure",
			"Data_Shapes",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.DiagramStructure) []*models.DataShape {
				return owner.Data_Shapes
			})

	case *models.DiagramLayerState:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("DiagramStructure", instanceWithInferedType.DiagramStructure, formGroup, probe)
		AssociationFieldToForm("LayerDefinition", instanceWithInferedType.LayerDefinition, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.DiagramStructure:
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
		BasicFieldtoForm("IsWithDiscretePorts", instanceWithInferedType.IsWithDiscretePorts, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("System_Shapes", instanceWithInferedType, &instanceWithInferedType.System_Shapes, formGroup, probe)
		BasicFieldtoForm("IsSystemsNodeExpanded", instanceWithInferedType.IsSystemsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("SystemsWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.SystemsWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("Part_Shapes", instanceWithInferedType, &instanceWithInferedType.Part_Shapes, formGroup, probe)
		BasicFieldtoForm("IsPartsNodeExpanded", instanceWithInferedType.IsPartsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("PartWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.PartWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("ExternalPart_Shapes", instanceWithInferedType, &instanceWithInferedType.ExternalPart_Shapes, formGroup, probe)
		BasicFieldtoForm("IsExternalPartsNodeExpanded", instanceWithInferedType.IsExternalPartsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("ExternalPartWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ExternalPartWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("ExternalPartsWhoseOutDataFlowsNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ExternalPartsWhoseOutDataFlowsNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("ExternalPartsWhoseInDataFlowsNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ExternalPartsWhoseInDataFlowsNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("PortsWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.PortsWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("Port_Shapes", instanceWithInferedType, &instanceWithInferedType.Port_Shapes, formGroup, probe)
		AssociationSliceToForm("ControlFlowsWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ControlFlowsWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("ControlFlow_Shapes", instanceWithInferedType, &instanceWithInferedType.ControlFlow_Shapes, formGroup, probe)
		AssociationSliceToForm("DataFlowsWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.DataFlowsWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("DataFlow_Shapes", instanceWithInferedType, &instanceWithInferedType.DataFlow_Shapes, formGroup, probe)
		AssociationSliceToForm("DatasWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.DatasWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("Data_Shapes", instanceWithInferedType, &instanceWithInferedType.Data_Shapes, formGroup, probe)
		AssociationSliceToForm("DataFlowsWhoseDataNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.DataFlowsWhoseDataNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("AllocatedResourcesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.AllocatedResourcesWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("AllocatedResourceShapes", instanceWithInferedType, &instanceWithInferedType.AllocatedResourceShapes, formGroup, probe)
		AssociationSliceToForm("AllocatedSystemesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.AllocatedSystemesWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("AllocatedSystemShapes", instanceWithInferedType, &instanceWithInferedType.AllocatedSystemShapes, formGroup, probe)
		AssociationSliceToForm("Note_Shapes", instanceWithInferedType, &instanceWithInferedType.Note_Shapes, formGroup, probe)
		AssociationSliceToForm("NotesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.NotesWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsNotesNodeExpanded", instanceWithInferedType.IsNotesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("NotePortShapes", instanceWithInferedType, &instanceWithInferedType.NotePortShapes, formGroup, probe)
		AssociationSliceToForm("NotePartShapes", instanceWithInferedType, &instanceWithInferedType.NotePartShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		AssociationReverseSliceToForm(
			"System",
			"DiagramStructures",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.System) []*models.DiagramStructure {
				return owner.DiagramStructures
			})
		AssociationReverseSliceToForm(
			"System",
			"DiagramStructureWhoseNodeIsExpanded",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.System) []*models.DiagramStructure {
				return owner.DiagramStructureWhoseNodeIsExpanded
			})

	case *models.ExternalPartShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Part", instanceWithInferedType.Part, formGroup, probe)
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
		BasicFieldtoForm("TailHeigth", instanceWithInferedType.TailHeigth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		AssociationReverseSliceToForm(
			"DiagramStructure",
			"ExternalPart_Shapes",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.DiagramStructure) []*models.ExternalPartShape {
				return owner.ExternalPart_Shapes
			})

	case *models.LayerDefinition:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Query", instanceWithInferedType, &instanceWithInferedType.Query, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

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
		AssociationSliceToForm("RootDataFlows", instanceWithInferedType, &instanceWithInferedType.RootDataFlows, formGroup, probe)
		BasicFieldtoForm("IsDataFlowsNodeExpanded", instanceWithInferedType.IsDataFlowsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("DataFlowsWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.DataFlowsWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("RootDatas", instanceWithInferedType, &instanceWithInferedType.RootDatas, formGroup, probe)
		BasicFieldtoForm("IsDatasNodeExpanded", instanceWithInferedType.IsDatasNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("DatasWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.DatasWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("RootResources", instanceWithInferedType, &instanceWithInferedType.RootResources, formGroup, probe)
		BasicFieldtoForm("IsResourcesNodeExpanded", instanceWithInferedType.IsResourcesNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("ResourcesWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ResourcesWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("PartsWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.PartsWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("RootNotes", instanceWithInferedType, &instanceWithInferedType.RootNotes, formGroup, probe)
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
		AssociationReverseSliceToForm(
			"Library",
			"SubLibraries",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.Library) []*models.Library {
				return owner.SubLibraries
			})
		AssociationReverseSliceToForm(
			"Library",
			"SubLibrariesWhoseNodeIsExpanded",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.Library) []*models.Library {
				return owner.SubLibrariesWhoseNodeIsExpanded
			})

	case *models.Note:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0, false)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0, false)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsPartsNodeExpanded", instanceWithInferedType.IsPartsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Parts", instanceWithInferedType, &instanceWithInferedType.Parts, formGroup, probe)
		BasicFieldtoForm("IsPortsNodeExpanded", instanceWithInferedType.IsPortsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Ports", instanceWithInferedType, &instanceWithInferedType.Ports, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		AssociationReverseSliceToForm(
			"DiagramStructure",
			"NotesWhoseNodeIsExpanded",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.DiagramStructure) []*models.Note {
				return owner.NotesWhoseNodeIsExpanded
			})
		AssociationReverseSliceToForm(
			"Library",
			"RootNotes",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.Library) []*models.Note {
				return owner.RootNotes
			})
		AssociationReverseSliceToForm(
			"Library",
			"NotesWhoseNodeIsExpanded",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.Library) []*models.Note {
				return owner.NotesWhoseNodeIsExpanded
			})

	case *models.NotePartShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Note", instanceWithInferedType.Note, formGroup, probe)
		AssociationFieldToForm("Part", instanceWithInferedType.Part, formGroup, probe)
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
		AssociationReverseSliceToForm(
			"DiagramStructure",
			"NotePartShapes",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.DiagramStructure) []*models.NotePartShape {
				return owner.NotePartShapes
			})

	case *models.NotePortShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Note", instanceWithInferedType.Note, formGroup, probe)
		AssociationFieldToForm("Port", instanceWithInferedType.Port, formGroup, probe)
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
		AssociationReverseSliceToForm(
			"DiagramStructure",
			"NotePortShapes",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.DiagramStructure) []*models.NotePortShape {
				return owner.NotePortShapes
			})

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
		AssociationReverseSliceToForm(
			"DiagramStructure",
			"Note_Shapes",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.DiagramStructure) []*models.NoteShape {
				return owner.Note_Shapes
			})

	case *models.Part:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0, false)
		AssociationSliceToForm("Ports", instanceWithInferedType, &instanceWithInferedType.Ports, formGroup, probe)
		AssociationFieldToForm("TypeOfPart", instanceWithInferedType.TypeOfPart, formGroup, probe)
		BasicFieldtoForm("IsPartNameNotSystemName", instanceWithInferedType.IsPartNameNotSystemName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsControlFlowsNodeExpanded", instanceWithInferedType.IsControlFlowsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("ControlFlows", instanceWithInferedType, &instanceWithInferedType.ControlFlows, formGroup, probe)
		AssociationSliceToForm("PortWhoseOutControlFlowsNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.PortWhoseOutControlFlowsNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("PortWhoseInControlFlowsNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.PortWhoseInControlFlowsNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsDataFlowsNodeExpanded", instanceWithInferedType.IsDataFlowsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("PortWhoseOutDataFlowsNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.PortWhoseOutDataFlowsNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("PortWhoseInDataFlowsNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.PortWhoseInDataFlowsNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("PartAnchoredPath", instanceWithInferedType, &instanceWithInferedType.PartAnchoredPath, formGroup, probe)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsPortsNodeExpanded", instanceWithInferedType.IsPortsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		AssociationReverseSliceToForm(
			"DiagramStructure",
			"PartWhoseNodeIsExpanded",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.DiagramStructure) []*models.Part {
				return owner.PartWhoseNodeIsExpanded
			})
		AssociationReverseSliceToForm(
			"DiagramStructure",
			"ExternalPartWhoseNodeIsExpanded",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.DiagramStructure) []*models.Part {
				return owner.ExternalPartWhoseNodeIsExpanded
			})
		AssociationReverseSliceToForm(
			"DiagramStructure",
			"ExternalPartsWhoseOutDataFlowsNodeIsExpanded",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.DiagramStructure) []*models.Part {
				return owner.ExternalPartsWhoseOutDataFlowsNodeIsExpanded
			})
		AssociationReverseSliceToForm(
			"DiagramStructure",
			"ExternalPartsWhoseInDataFlowsNodeIsExpanded",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.DiagramStructure) []*models.Part {
				return owner.ExternalPartsWhoseInDataFlowsNodeIsExpanded
			})
		AssociationReverseSliceToForm(
			"Library",
			"PartsWhoseNodeIsExpanded",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.Library) []*models.Part {
				return owner.PartsWhoseNodeIsExpanded
			})
		AssociationReverseSliceToForm(
			"Note",
			"Parts",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.Note) []*models.Part {
				return owner.Parts
			})
		AssociationReverseSliceToForm(
			"SemanticTag",
			"Parts",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.SemanticTag) []*models.Part {
				return owner.Parts
			})
		AssociationReverseSliceToForm(
			"System",
			"Parts",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.System) []*models.Part {
				return owner.Parts
			})
		AssociationReverseSliceToForm(
			"System",
			"PartWhoseNodeIsExpanded",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.System) []*models.Part {
				return owner.PartWhoseNodeIsExpanded
			})
		AssociationReverseSliceToForm(
			"System",
			"ExternalParts",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.System) []*models.Part {
				return owner.ExternalParts
			})
		AssociationReverseSliceToForm(
			"System",
			"ExternalPartWhoseNodeIsExpanded",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.System) []*models.Part {
				return owner.ExternalPartWhoseNodeIsExpanded
			})

	case *models.PartAnchoredPath:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Definition", instanceWithInferedType.Definition, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("X_Offset", instanceWithInferedType.X_Offset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Y_Offset", instanceWithInferedType.Y_Offset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("RectAnchorType", instanceWithInferedType.RectAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("ScalePropotionnally", instanceWithInferedType.ScalePropotionnally, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("AppliedScaling", instanceWithInferedType.AppliedScaling, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Stroke", instanceWithInferedType.Stroke, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("StrokeDashArrayWhenSelected", instanceWithInferedType.StrokeDashArrayWhenSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Transform", instanceWithInferedType.Transform, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		AssociationReverseSliceToForm(
			"Part",
			"PartAnchoredPath",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.Part) []*models.PartAnchoredPath {
				return owner.PartAnchoredPath
			})

	case *models.PartShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Part", instanceWithInferedType.Part, formGroup, probe)
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
		AssociationReverseSliceToForm(
			"DiagramStructure",
			"Part_Shapes",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.DiagramStructure) []*models.PartShape {
				return owner.Part_Shapes
			})

	case *models.Port:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
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
		AssociationReverseSliceToForm(
			"DiagramStructure",
			"PortsWhoseNodeIsExpanded",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.DiagramStructure) []*models.Port {
				return owner.PortsWhoseNodeIsExpanded
			})
		AssociationReverseSliceToForm(
			"Note",
			"Ports",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.Note) []*models.Port {
				return owner.Ports
			})
		AssociationReverseSliceToForm(
			"Part",
			"Ports",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.Part) []*models.Port {
				return owner.Ports
			})
		AssociationReverseSliceToForm(
			"Part",
			"PortWhoseOutControlFlowsNodeIsExpanded",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.Part) []*models.Port {
				return owner.PortWhoseOutControlFlowsNodeIsExpanded
			})
		AssociationReverseSliceToForm(
			"Part",
			"PortWhoseInControlFlowsNodeIsExpanded",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.Part) []*models.Port {
				return owner.PortWhoseInControlFlowsNodeIsExpanded
			})
		AssociationReverseSliceToForm(
			"Part",
			"PortWhoseOutDataFlowsNodeIsExpanded",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.Part) []*models.Port {
				return owner.PortWhoseOutDataFlowsNodeIsExpanded
			})
		AssociationReverseSliceToForm(
			"Part",
			"PortWhoseInDataFlowsNodeIsExpanded",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.Part) []*models.Port {
				return owner.PortWhoseInDataFlowsNodeIsExpanded
			})

	case *models.PortShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Port", instanceWithInferedType.Port, formGroup, probe)
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
		AssociationReverseSliceToForm(
			"DiagramStructure",
			"Port_Shapes",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.DiagramStructure) []*models.PortShape {
				return owner.Port_Shapes
			})

	case *models.Resource:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Acronym", instanceWithInferedType.Acronym, instanceWithInferedType, probe.formStage, formGroup,
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
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		AssociationReverseSliceToForm(
			"DiagramStructure",
			"AllocatedResourcesWhoseNodeIsExpanded",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.DiagramStructure) []*models.Resource {
				return owner.AllocatedResourcesWhoseNodeIsExpanded
			})
		AssociationReverseSliceToForm(
			"Library",
			"RootResources",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.Library) []*models.Resource {
				return owner.RootResources
			})
		AssociationReverseSliceToForm(
			"Library",
			"ResourcesWhoseNodeIsExpanded",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.Library) []*models.Resource {
				return owner.ResourcesWhoseNodeIsExpanded
			})

	case *models.SemanticTag:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Parts", instanceWithInferedType, &instanceWithInferedType.Parts, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		AssociationReverseSliceToForm(
			"LayerDefinition",
			"Query",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.LayerDefinition) []*models.SemanticTag {
				return owner.Query
			})

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
		AssociationSliceToForm("DiagramStructures", instanceWithInferedType, &instanceWithInferedType.DiagramStructures, formGroup, probe)
		AssociationSliceToForm("DiagramStructureWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.DiagramStructureWhoseNodeIsExpanded, formGroup, probe)
		BasicFieldtoForm("IsSubSystemNodeExpanded", instanceWithInferedType.IsSubSystemNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("SubSystemes", instanceWithInferedType, &instanceWithInferedType.SubSystemes, formGroup, probe)
		AssociationSliceToForm("Parts", instanceWithInferedType, &instanceWithInferedType.Parts, formGroup, probe)
		AssociationSliceToForm("PartWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.PartWhoseNodeIsExpanded, formGroup, probe)
		AssociationSliceToForm("DataFlows", instanceWithInferedType, &instanceWithInferedType.DataFlows, formGroup, probe)
		BasicFieldtoForm("IsDataFlowsNodeExpanded", instanceWithInferedType.IsDataFlowsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("ExternalParts", instanceWithInferedType, &instanceWithInferedType.ExternalParts, formGroup, probe)
		AssociationSliceToForm("ExternalPartWhoseNodeIsExpanded", instanceWithInferedType, &instanceWithInferedType.ExternalPartWhoseNodeIsExpanded, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		AssociationReverseSliceToForm(
			"DiagramStructure",
			"SystemsWhoseNodeIsExpanded",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.DiagramStructure) []*models.System {
				return owner.SystemsWhoseNodeIsExpanded
			})
		AssociationReverseSliceToForm(
			"DiagramStructure",
			"AllocatedSystemesWhoseNodeIsExpanded",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.DiagramStructure) []*models.System {
				return owner.AllocatedSystemesWhoseNodeIsExpanded
			})
		AssociationReverseSliceToForm(
			"Library",
			"RootSystemes",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.Library) []*models.System {
				return owner.RootSystemes
			})
		AssociationReverseSliceToForm(
			"Library",
			"SystemsWhoseNodeIsExpanded",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.Library) []*models.System {
				return owner.SystemsWhoseNodeIsExpanded
			})
		AssociationReverseSliceToForm(
			"System",
			"SubSystemes",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.System) []*models.System {
				return owner.SubSystemes
			})

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
		AssociationReverseSliceToForm(
			"DiagramStructure",
			"System_Shapes",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.DiagramStructure) []*models.SystemShape {
				return owner.System_Shapes
			})

	default:
		_ = instanceWithInferedType
	}
}
