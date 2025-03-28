// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
	"github.com/fullstack-lang/gong/lib/doc/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Classdiagram:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("GongStructShapes", instanceWithInferedType, &instanceWithInferedType.GongStructShapes, formGroup, probe)
		AssociationSliceToForm("GongEnumShapes", instanceWithInferedType, &instanceWithInferedType.GongEnumShapes, formGroup, probe)
		AssociationSliceToForm("NoteShapes", instanceWithInferedType, &instanceWithInferedType.NoteShapes, formGroup, probe)
		BasicFieldtoForm("IsInDrawMode", instanceWithInferedType.IsInDrawMode, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DiagramPackage"
			rf.Fieldname = "Classdiagrams"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.DiagramPackage),
					"Classdiagrams",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.DiagramPackage, *models.Classdiagram](
					nil,
					"Classdiagrams",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.DiagramPackage:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Path", instanceWithInferedType.Path, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("GongModelPath", instanceWithInferedType.GongModelPath, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Classdiagrams", instanceWithInferedType, &instanceWithInferedType.Classdiagrams, formGroup, probe)
		AssociationFieldToForm("SelectedClassdiagram", instanceWithInferedType.SelectedClassdiagram, formGroup, probe)
		AssociationSliceToForm("Umlscs", instanceWithInferedType, &instanceWithInferedType.Umlscs, formGroup, probe)
		BasicFieldtoForm("IsEditable", instanceWithInferedType.IsEditable, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsReloaded", instanceWithInferedType.IsReloaded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("AbsolutePathToDiagramPackage", instanceWithInferedType.AbsolutePathToDiagramPackage, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Field:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FieldTypeAsString", instanceWithInferedType.FieldTypeAsString, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Structname", instanceWithInferedType.Structname, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Fieldtypename", instanceWithInferedType.Fieldtypename, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongStructShape"
			rf.Fieldname = "Fields"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.GongStructShape),
					"Fields",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.GongStructShape, *models.Field](
					nil,
					"Fields",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.GongEnumShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Position", instanceWithInferedType.Position, formGroup, probe)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("GongEnumValueEntrys", instanceWithInferedType, &instanceWithInferedType.GongEnumValueEntrys, formGroup, probe)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Classdiagram"
			rf.Fieldname = "GongEnumShapes"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Classdiagram),
					"GongEnumShapes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Classdiagram, *models.GongEnumShape](
					nil,
					"GongEnumShapes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.GongEnumValueEntry:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongEnumShape"
			rf.Fieldname = "GongEnumValueEntrys"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.GongEnumShape),
					"GongEnumValueEntrys",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.GongEnumShape, *models.GongEnumValueEntry](
					nil,
					"GongEnumValueEntrys",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.GongStructShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Position", instanceWithInferedType.Position, formGroup, probe)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ShowNbInstances", instanceWithInferedType.ShowNbInstances, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NbInstances", instanceWithInferedType.NbInstances, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Fields", instanceWithInferedType, &instanceWithInferedType.Fields, formGroup, probe)
		AssociationSliceToForm("Links", instanceWithInferedType, &instanceWithInferedType.Links, formGroup, probe)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsSelected", instanceWithInferedType.IsSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Classdiagram"
			rf.Fieldname = "GongStructShapes"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Classdiagram),
					"GongStructShapes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Classdiagram, *models.GongStructShape](
					nil,
					"GongStructShapes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Link:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Fieldtypename", instanceWithInferedType.Fieldtypename, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FieldOffsetX", instanceWithInferedType.FieldOffsetX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FieldOffsetY", instanceWithInferedType.FieldOffsetY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("TargetMultiplicity", instanceWithInferedType.TargetMultiplicity, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("TargetMultiplicityOffsetX", instanceWithInferedType.TargetMultiplicityOffsetX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("TargetMultiplicityOffsetY", instanceWithInferedType.TargetMultiplicityOffsetY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("SourceMultiplicity", instanceWithInferedType.SourceMultiplicity, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("SourceMultiplicityOffsetX", instanceWithInferedType.SourceMultiplicityOffsetX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SourceMultiplicityOffsetY", instanceWithInferedType.SourceMultiplicityOffsetY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Middlevertice", instanceWithInferedType.Middlevertice, formGroup, probe)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "GongStructShape"
			rf.Fieldname = "Links"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.GongStructShape),
					"Links",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.GongStructShape, *models.Link](
					nil,
					"Links",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.NoteShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Body", instanceWithInferedType.Body, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("BodyHTML", instanceWithInferedType.BodyHTML, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Matched", instanceWithInferedType.Matched, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("NoteShapeLinks", instanceWithInferedType, &instanceWithInferedType.NoteShapeLinks, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Classdiagram"
			rf.Fieldname = "NoteShapes"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Classdiagram),
					"NoteShapes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Classdiagram, *models.NoteShape](
					nil,
					"NoteShapes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.NoteShapeLink:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "NoteShape"
			rf.Fieldname = "NoteShapeLinks"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.NoteShape),
					"NoteShapeLinks",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.NoteShape, *models.NoteShapeLink](
					nil,
					"NoteShapeLinks",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Position:
		// insertion point
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.UmlState:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Umlsc"
			rf.Fieldname = "States"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Umlsc),
					"States",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Umlsc, *models.UmlState](
					nil,
					"States",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Umlsc:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("States", instanceWithInferedType, &instanceWithInferedType.States, formGroup, probe)
		BasicFieldtoForm("Activestate", instanceWithInferedType.Activestate, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsInDrawMode", instanceWithInferedType.IsInDrawMode, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "DiagramPackage"
			rf.Fieldname = "Umlscs"
			reverseFieldOwner := models.GetReverseFieldOwner(probe.stageOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.DiagramPackage),
					"Umlscs",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.DiagramPackage, *models.Umlsc](
					nil,
					"Umlscs",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Vertice:
		// insertion point
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	default:
		_ = instanceWithInferedType
	}
}
