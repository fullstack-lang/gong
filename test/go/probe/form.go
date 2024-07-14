package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/test/go/models"
	"github.com/fullstack-lang/gong/test/go/orm"
)

// AstructForm provides the interface functions for
// generating the form
type AstructNode struct {
	instance *models.Astruct
	probe    *Probe
}

func (astructForm *AstructNode) CreateForm() {
	formGroup := (&gongtable.FormGroup{
		Name:  gongtable.FormGroupDefaultName.ToString(),
		Label: "Astruct Form",
	}).Stage(astructForm.probe.formStage)
	formGroup.OnSave = __gong__New__AstructFormCallback(
		astructForm.instance,
		astructForm.probe,
		formGroup,
	)
	formGroup.HasSuppressButton = true

	BasicFieldtoForm("Name", astructForm.instance.Name, astructForm.instance, astructForm.probe.formStage, formGroup,
		false, false, 0, false, 0)
	AssociationFieldToForm("Associationtob", astructForm.instance.Associationtob, formGroup, astructForm.probe)
	AssociationSliceToForm("Anarrayofb", astructForm.instance, &astructForm.instance.Anarrayofb, formGroup, astructForm.probe)
	AssociationFieldToForm("Anotherassociationtob_2", astructForm.instance.Anotherassociationtob_2, formGroup, astructForm.probe)
	BasicFieldtoForm("Date", astructForm.instance.Date, astructForm.instance, astructForm.probe.formStage, formGroup,
		false, false, 0, false, 0)
	BasicFieldtoForm("Booleanfield", astructForm.instance.Booleanfield, astructForm.instance, astructForm.probe.formStage, formGroup,
		false, false, 0, false, 0)
	EnumTypeStringToForm("Aenum", astructForm.instance.Aenum, astructForm.instance, astructForm.probe.formStage, formGroup)
	EnumTypeStringToForm("Aenum_2", astructForm.instance.Aenum_2, astructForm.instance, astructForm.probe.formStage, formGroup)
	EnumTypeStringToForm("Benum", astructForm.instance.Benum, astructForm.instance, astructForm.probe.formStage, formGroup)
	EnumTypeIntToForm("CEnum", astructForm.instance.CEnum, astructForm.instance, astructForm.probe.formStage, formGroup)
	BasicFieldtoForm("CName", astructForm.instance.CName, astructForm.instance, astructForm.probe.formStage, formGroup,
		false, false, 0, false, 0)
	BasicFieldtoForm("CFloatfield", astructForm.instance.CFloatfield, astructForm.instance, astructForm.probe.formStage, formGroup,
		false, false, 0, false, 0)
	AssociationFieldToForm("Bstruct", astructForm.instance.Bstruct, formGroup, astructForm.probe)
	AssociationFieldToForm("Bstruct2", astructForm.instance.Bstruct2, formGroup, astructForm.probe)
	AssociationFieldToForm("Dstruct", astructForm.instance.Dstruct, formGroup, astructForm.probe)
	AssociationFieldToForm("Dstruct2", astructForm.instance.Dstruct2, formGroup, astructForm.probe)
	AssociationFieldToForm("Dstruct3", astructForm.instance.Dstruct3, formGroup, astructForm.probe)
	AssociationFieldToForm("Dstruct4", astructForm.instance.Dstruct4, formGroup, astructForm.probe)
	BasicFieldtoForm("Floatfield", astructForm.instance.Floatfield, astructForm.instance, astructForm.probe.formStage, formGroup,
		false, false, 0, false, 0)
	BasicFieldtoForm("Intfield", astructForm.instance.Intfield, astructForm.instance, astructForm.probe.formStage, formGroup,
		false, false, 0, false, 0)
	BasicFieldtoForm("Anotherbooleanfield", astructForm.instance.Anotherbooleanfield, astructForm.instance, astructForm.probe.formStage, formGroup,
		false, false, 0, false, 0)
	BasicFieldtoForm("Duration1", astructForm.instance.Duration1, astructForm.instance, astructForm.probe.formStage, formGroup,
		false, false, 0, false, 0)
	AssociationSliceToForm("Anarrayofa", astructForm.instance, &astructForm.instance.Anarrayofa, formGroup, astructForm.probe)
	AssociationSliceToForm("Anotherarrayofb", astructForm.instance, &astructForm.instance.Anotherarrayofb, formGroup, astructForm.probe)
	AssociationSliceToForm("AnarrayofbUse", astructForm.instance, &astructForm.instance.AnarrayofbUse, formGroup, astructForm.probe)
	AssociationSliceToForm("Anarrayofb2Use", astructForm.instance, &astructForm.instance.Anarrayofb2Use, formGroup, astructForm.probe)
	AssociationFieldToForm("AnAstruct", astructForm.instance.AnAstruct, formGroup, astructForm.probe)
	BasicFieldtoForm("StructRef", astructForm.instance.StructRef, astructForm.instance, astructForm.probe.formStage, formGroup,
		false, false, 0, false, 0)
	BasicFieldtoForm("FieldRef", astructForm.instance.FieldRef, astructForm.instance, astructForm.probe.formStage, formGroup,
		false, false, 0, false, 0)
	BasicFieldtoForm("EnumIntRef", astructForm.instance.EnumIntRef, astructForm.instance, astructForm.probe.formStage, formGroup,
		false, false, 0, false, 0)
	BasicFieldtoForm("EnumStringRef", astructForm.instance.EnumStringRef, astructForm.instance, astructForm.probe.formStage, formGroup,
		false, false, 0, false, 0)
	BasicFieldtoForm("EnumValue", astructForm.instance.EnumValue, astructForm.instance, astructForm.probe.formStage, formGroup,
		false, false, 0, false, 0)
	BasicFieldtoForm("ConstIdentifierValue", astructForm.instance.ConstIdentifierValue, astructForm.instance, astructForm.probe.formStage, formGroup,
		false, false, 0, false, 0)
	BasicFieldtoForm("TextFieldBespokeSize", astructForm.instance.TextFieldBespokeSize, astructForm.instance, astructForm.probe.formStage, formGroup,
		false, true, 600, true, 300)
	BasicFieldtoForm("TextArea", astructForm.instance.TextArea, astructForm.instance, astructForm.probe.formStage, formGroup,
		true, false, 0, false, 0)
	{
		var rf models.ReverseField
		_ = rf
		rf.GongstructName = "Astruct"
		rf.Fieldname = "Anarrayofa"
		reverseFieldOwner := orm.GetReverseFieldOwner(astructForm.probe.stageOfInterest, astructForm.probe.backRepoOfInterest, astructForm.instance, &rf)
		if reverseFieldOwner != nil {
			AssociationReverseFieldToForm(
				reverseFieldOwner.(*models.Astruct),
				"Anarrayofa",
				astructForm.instance,
				formGroup,
				astructForm.probe)
		} else {
			AssociationReverseFieldToForm[*models.Astruct, *models.Astruct](
				nil,
				"Anarrayofa",
				astructForm.instance,
				formGroup,
				astructForm.probe)
		}
	}
}
