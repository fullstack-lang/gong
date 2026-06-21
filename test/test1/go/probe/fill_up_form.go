// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/test/test1/go/models"
)

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Astruct:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Associationtob", instanceWithInferedType.Associationtob, formGroup, probe)
		AssociationSliceToForm("Anarrayofb", instanceWithInferedType, &instanceWithInferedType.Anarrayofb, formGroup, probe)
		AssociationFieldToForm("Anotherassociationtob_2", instanceWithInferedType.Anotherassociationtob_2, formGroup, probe)
		BasicFieldtoForm("Date", instanceWithInferedType.Date, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Date2", instanceWithInferedType.Date2, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Booleanfield", instanceWithInferedType.Booleanfield, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Aenum", instanceWithInferedType.Aenum, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Aenum_2", instanceWithInferedType.Aenum_2, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Benum", instanceWithInferedType.Benum, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeIntToForm("CEnum", instanceWithInferedType.CEnum, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("CName", instanceWithInferedType.CName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("CFloatfield", instanceWithInferedType.CFloatfield, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Bstruct", instanceWithInferedType.Bstruct, formGroup, probe)
		AssociationFieldToForm("Bstruct2", instanceWithInferedType.Bstruct2, formGroup, probe)
		AssociationFieldToForm("Dstruct", instanceWithInferedType.Dstruct, formGroup, probe)
		AssociationFieldToForm("Dstruct2", instanceWithInferedType.Dstruct2, formGroup, probe)
		AssociationFieldToForm("Dstruct3", instanceWithInferedType.Dstruct3, formGroup, probe)
		AssociationFieldToForm("Dstruct4", instanceWithInferedType.Dstruct4, formGroup, probe)
		AssociationSliceToForm("Dstruct4s", instanceWithInferedType, &instanceWithInferedType.Dstruct4s, formGroup, probe)
		BasicFieldtoForm("Floatfield", instanceWithInferedType.Floatfield, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Intfield", instanceWithInferedType.Intfield, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Anotherbooleanfield", instanceWithInferedType.Anotherbooleanfield, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Duration1", instanceWithInferedType.Duration1, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Anarrayofa", instanceWithInferedType, &instanceWithInferedType.Anarrayofa, formGroup, probe)
		AssociationSliceToForm("Anotherarrayofb", instanceWithInferedType, &instanceWithInferedType.Anotherarrayofb, formGroup, probe)
		AssociationSliceToForm("AnarrayofbUse", instanceWithInferedType, &instanceWithInferedType.AnarrayofbUse, formGroup, probe)
		AssociationSliceToForm("Anarrayofb2Use", instanceWithInferedType, &instanceWithInferedType.Anarrayofb2Use, formGroup, probe)
		AssociationFieldToForm("AnAstruct", instanceWithInferedType.AnAstruct, formGroup, probe)
		BasicFieldtoForm("TextFieldBespokeSize", instanceWithInferedType.TextFieldBespokeSize, instanceWithInferedType, probe.formStage, formGroup,
			false, true, 600, true, 300)
		BasicFieldtoForm("TextArea", instanceWithInferedType.TextArea, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Astruct, *models.Astruct](
				"Astruct",
				"Anarrayofa",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Astruct) []*models.Astruct {
					return owner.Anarrayofa
				})
		}

	case *models.AstructBstruct2Use:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Bstrcut2", instanceWithInferedType.Bstrcut2, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Astruct, *models.AstructBstruct2Use](
				"Astruct",
				"Anarrayofb2Use",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Astruct) []*models.AstructBstruct2Use {
					return owner.Anarrayofb2Use
				})
		}

	case *models.AstructBstructUse:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Bstruct2", instanceWithInferedType.Bstruct2, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Astruct, *models.AstructBstructUse](
				"Astruct",
				"AnarrayofbUse",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Astruct) []*models.AstructBstructUse {
					return owner.AnarrayofbUse
				})
		}

	case *models.Bstruct:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Floatfield", instanceWithInferedType.Floatfield, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Floatfield2", instanceWithInferedType.Floatfield2, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Intfield", instanceWithInferedType.Intfield, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Astruct, *models.Bstruct](
				"Astruct",
				"Anarrayofb",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Astruct) []*models.Bstruct {
					return owner.Anarrayofb
				})
		}
		{
			AssociationReverseSliceToForm[*models.Astruct, *models.Bstruct](
				"Astruct",
				"Anotherarrayofb",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Astruct) []*models.Bstruct {
					return owner.Anotherarrayofb
				})
		}
		{
			AssociationReverseSliceToForm[*models.Dstruct, *models.Bstruct](
				"Dstruct",
				"Anarrayofb",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Dstruct) []*models.Bstruct {
					return owner.Anarrayofb
				})
		}

	case *models.Dstruct:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Anarrayofb", instanceWithInferedType, &instanceWithInferedType.Anarrayofb, formGroup, probe)
		AssociationFieldToForm("Gstruct", instanceWithInferedType.Gstruct, formGroup, probe)
		AssociationSliceToForm("Gstructs", instanceWithInferedType, &instanceWithInferedType.Gstructs, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Astruct, *models.Dstruct](
				"Astruct",
				"Dstruct4s",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Astruct) []*models.Dstruct {
					return owner.Dstruct4s
				})
		}

	case *models.F0123456789012345678901234567890:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Date", instanceWithInferedType.Date, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Gstruct:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Floatfield", instanceWithInferedType.Floatfield, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Floatfield2", instanceWithInferedType.Floatfield2, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Intfield", instanceWithInferedType.Intfield, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Dstruct, *models.Gstruct](
				"Dstruct",
				"Gstructs",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Dstruct) []*models.Gstruct {
					return owner.Gstructs
				})
		}

	default:
		_ = instanceWithInferedType
	}
}
