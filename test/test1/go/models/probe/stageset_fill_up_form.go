// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/test/test1/go/models"
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
	case *models.Astruct:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Associationtob", inst.Associationtob, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Bstruct](), probe.formStage)

		{
			// Slice of pointers: Anarrayofb
			div := (&form.FormDiv{Name: "Anarrayofb"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Anarrayofb {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Anarrayofb",
				Label: "Anarrayofb",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Anotherassociationtob_2", inst.Anotherassociationtob_2, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Bstruct](), probe.formStage)
		StageSetBasicFieldtoForm("Date", inst.Date, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Date2", inst.Date2, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Booleanfield", inst.Booleanfield, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Aenum", inst.Aenum, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Aenum_2", inst.Aenum_2, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("Benum", inst.Benum, formGroup, probe.formStage)
		StageSetEnumIntFieldToForm("CEnum", inst.CEnum, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("CName", inst.CName, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("CFloatfield", inst.CFloatfield, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Bstruct", inst.Bstruct, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Bstruct](), probe.formStage)
		StageSetAssociationFieldToForm("Bstruct2", inst.Bstruct2, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Bstruct](), probe.formStage)
		StageSetAssociationFieldToForm("Dstruct", inst.Dstruct, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Dstruct](), probe.formStage)
		StageSetAssociationFieldToForm("Dstruct2", inst.Dstruct2, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Dstruct](), probe.formStage)
		StageSetAssociationFieldToForm("Dstruct3", inst.Dstruct3, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Dstruct](), probe.formStage)
		StageSetAssociationFieldToForm("Dstruct4", inst.Dstruct4, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Dstruct](), probe.formStage)

		{
			// Slice of pointers: Dstruct4s
			div := (&form.FormDiv{Name: "Dstruct4s"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Dstruct4s {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Dstruct4s",
				Label: "Dstruct4s",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("Floatfield", inst.Floatfield, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Intfield", inst.Intfield, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Anotherbooleanfield", inst.Anotherbooleanfield, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Duration1", inst.Duration1, probe.formStage, formGroup)

		{
			// Slice of pointers: Anarrayofa
			div := (&form.FormDiv{Name: "Anarrayofa"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Anarrayofa {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Anarrayofa",
				Label: "Anarrayofa",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Anotherarrayofb
			div := (&form.FormDiv{Name: "Anotherarrayofb"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Anotherarrayofb {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Anotherarrayofb",
				Label: "Anotherarrayofb",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: AnarrayofbUse
			div := (&form.FormDiv{Name: "AnarrayofbUse"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AnarrayofbUse {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AnarrayofbUse",
				Label: "AnarrayofbUse",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Anarrayofb2Use
			div := (&form.FormDiv{Name: "Anarrayofb2Use"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Anarrayofb2Use {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Anarrayofb2Use",
				Label: "Anarrayofb2Use",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("AnAstruct", inst.AnAstruct, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Astruct](), probe.formStage)
		StageSetBasicFieldtoForm("TextFieldBespokeSize", inst.TextFieldBespokeSize, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TextArea", inst.TextArea, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Astructs {
				for _, target := range src.Anarrayofa {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Astruct", "Anarrayofa", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Astructs {
				if src.AnAstruct == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Astruct", "AnAstruct", refNames, formGroup, probe.formStage)
		}
	case *models.AstructBstruct2Use:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Bstrcut2", inst.Bstrcut2, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Bstruct](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Astructs {
				for _, target := range src.Anarrayofb2Use {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Astruct", "Anarrayofb2Use", refNames, formGroup, probe.formStage)
		}
	case *models.AstructBstructUse:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Bstruct2", inst.Bstruct2, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Bstruct](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Astructs {
				for _, target := range src.AnarrayofbUse {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Astruct", "AnarrayofbUse", refNames, formGroup, probe.formStage)
		}
	case *models.Bstruct:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Floatfield", inst.Floatfield, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Floatfield2", inst.Floatfield2, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Intfield", inst.Intfield, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ToBeIgnored", inst.ToBeIgnored, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Astructs {
				if src.Associationtob == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Astruct", "Associationtob", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Astructs {
				for _, target := range src.Anarrayofb {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Astruct", "Anarrayofb", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Astructs {
				if src.Anotherassociationtob_2 == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Astruct", "Anotherassociationtob_2", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Astructs {
				if src.Bstruct == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Astruct", "Bstruct", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Astructs {
				if src.Bstruct2 == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Astruct", "Bstruct2", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Astructs {
				for _, target := range src.Anotherarrayofb {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Astruct", "Anotherarrayofb", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.AstructBstruct2Uses {
				if src.Bstrcut2 == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AstructBstruct2Use", "Bstrcut2", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.AstructBstructUses {
				if src.Bstruct2 == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AstructBstructUse", "Bstruct2", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Dstructs {
				for _, target := range src.Anarrayofb {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Dstruct", "Anarrayofb", refNames, formGroup, probe.formStage)
		}
	case *models.Dstruct:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: Anarrayofb
			div := (&form.FormDiv{Name: "Anarrayofb"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Anarrayofb {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Anarrayofb",
				Label: "Anarrayofb",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Gstruct", inst.Gstruct, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Gstruct](), probe.formStage)

		{
			// Slice of pointers: Gstructs
			div := (&form.FormDiv{Name: "Gstructs"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Gstructs {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Gstructs",
				Label: "Gstructs",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Astructs {
				if src.Dstruct == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Astruct", "Dstruct", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Astructs {
				if src.Dstruct2 == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Astruct", "Dstruct2", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Astructs {
				if src.Dstruct3 == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Astruct", "Dstruct3", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Astructs {
				if src.Dstruct4 == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Astruct", "Dstruct4", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Astructs {
				for _, target := range src.Dstruct4s {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Astruct", "Dstruct4s", refNames, formGroup, probe.formStage)
		}
	case *models.F0123456789012345678901234567890:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Date", inst.Date, probe.formStage, formGroup)
	case *models.Gstruct:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Floatfield", inst.Floatfield, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Floatfield2", inst.Floatfield2, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Intfield", inst.Intfield, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ToBeIgnored", inst.ToBeIgnored, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Dstructs {
				if src.Gstruct == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Dstruct", "Gstruct", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Dstructs {
				for _, target := range src.Gstructs {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Dstruct", "Gstructs", refNames, formGroup, probe.formStage)
		}
	default:
		_ = inst
	}
}
