// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/go/models"
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
	case *models.GongBasicField:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BasicKindName", inst.BasicKindName, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("GongEnum", inst.GongEnum, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.GongEnum](), probe.formStage)
		StageSetBasicFieldtoForm("DeclaredType", inst.DeclaredType, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("CompositeStructName", inst.CompositeStructName, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsAccordionStart", inst.IsAccordionStart, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("AccordionName", inst.AccordionName, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsAccordionEnd", inst.IsAccordionEnd, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Index", inst.Index, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsTextArea", inst.IsTextArea, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsBespokeWidth", inst.IsBespokeWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BespokeWidth", inst.BespokeWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsBespokeHeight", inst.IsBespokeHeight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BespokeHeight", inst.BespokeHeight, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.GongStructs {
				for _, target := range src.GongBasicFields {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.GongStruct", "GongBasicFields", refNames, formGroup, probe.formStage)
		}
	case *models.GongEnum:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumIntFieldToForm("Type", inst.Type, formGroup, probe.formStage)

		{
			// Slice of pointers: GongEnumValues
			div := (&form.FormDiv{Name: "GongEnumValues"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GongEnumValues {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GongEnumValues",
				Label: "GongEnumValues",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.GongBasicFields {
				if src.GongEnum == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.GongBasicField", "GongEnum", refNames, formGroup, probe.formStage)
		}
	case *models.GongEnumValue:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.GongEnums {
				for _, target := range src.GongEnumValues {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.GongEnum", "GongEnumValues", refNames, formGroup, probe.formStage)
		}
	case *models.GongLink:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Recv", inst.Recv, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ImportPath", inst.ImportPath, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.GongNotes {
				for _, target := range src.Links {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.GongNote", "Links", refNames, formGroup, probe.formStage)
		}
	case *models.GongNote:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Body", inst.Body, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BodyHTML", inst.BodyHTML, probe.formStage, formGroup)

		{
			// Slice of pointers: Links
			div := (&form.FormDiv{Name: "Links"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Links {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Links",
				Label: "Links",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
	case *models.GongStruct:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: GongBasicFields
			div := (&form.FormDiv{Name: "GongBasicFields"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GongBasicFields {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GongBasicFields",
				Label: "GongBasicFields",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: GongTimeFields
			div := (&form.FormDiv{Name: "GongTimeFields"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GongTimeFields {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GongTimeFields",
				Label: "GongTimeFields",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: PointerToGongStructFields
			div := (&form.FormDiv{Name: "PointerToGongStructFields"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.PointerToGongStructFields {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "PointerToGongStructFields",
				Label: "PointerToGongStructFields",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: SliceOfPointerToGongStructFields
			div := (&form.FormDiv{Name: "SliceOfPointerToGongStructFields"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SliceOfPointerToGongStructFields {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SliceOfPointerToGongStructFields",
				Label: "SliceOfPointerToGongStructFields",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("HasOnAfterUpdateSignature", inst.HasOnAfterUpdateSignature, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsIgnoredForFront", inst.IsIgnoredForFront, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsOmittedForMarshalling", inst.IsOmittedForMarshalling, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ModelPkg", inst.ModelPkg, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ModelPkg](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.PointerToGongStructFields {
				if src.GongStruct == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.PointerToGongStructField", "GongStruct", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.SliceOfPointerToGongStructFields {
				if src.GongStruct == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SliceOfPointerToGongStructField", "GongStruct", refNames, formGroup, probe.formStage)
		}
	case *models.GongTimeField:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Index", inst.Index, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("CompositeStructName", inst.CompositeStructName, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsAccordionStart", inst.IsAccordionStart, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("AccordionName", inst.AccordionName, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsAccordionEnd", inst.IsAccordionEnd, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BespokeTimeFormat", inst.BespokeTimeFormat, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TimeFormOnly", inst.TimeFormOnly, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.GongStructs {
				for _, target := range src.GongTimeFields {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.GongStruct", "GongTimeFields", refNames, formGroup, probe.formStage)
		}
	case *models.MetaReference:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
	case *models.ModelPkg:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("PkgGoName", inst.PkgGoName, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("PkgPath", inst.PkgPath, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("StageSet", inst.StageSet, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.StageSetModel](), probe.formStage)
		StageSetBasicFieldtoForm("PathToGoSubDirectory", inst.PathToGoSubDirectory, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("OrmPkgGenPath", inst.OrmPkgGenPath, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DbOrmPkgGenPath", inst.DbOrmPkgGenPath, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DbLiteOrmPkgGenPath", inst.DbLiteOrmPkgGenPath, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DbPkgGenPath", inst.DbPkgGenPath, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ControllersPkgGenPath", inst.ControllersPkgGenPath, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FullstackPkgGenPath", inst.FullstackPkgGenPath, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StackPkgGenPath", inst.StackPkgGenPath, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Level1StackPkgGenPath", inst.Level1StackPkgGenPath, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StaticPkgGenPath", inst.StaticPkgGenPath, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ProbePkgGenPath", inst.ProbePkgGenPath, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NgWorkspacePath", inst.NgWorkspacePath, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NgWorkspaceName", inst.NgWorkspaceName, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NgDataLibrarySourceCodeDirectory", inst.NgDataLibrarySourceCodeDirectory, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NgSpecificLibrarySourceCodeDirectory", inst.NgSpecificLibrarySourceCodeDirectory, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MaterialLibDatamodelTargetPath", inst.MaterialLibDatamodelTargetPath, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.GongStructs {
				if src.ModelPkg == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.GongStruct", "ModelPkg", refNames, formGroup, probe.formStage)
		}
	case *models.PointerToGongStructField:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("GongStruct", inst.GongStruct, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.GongStruct](), probe.formStage)
		StageSetBasicFieldtoForm("Index", inst.Index, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("CompositeStructName", inst.CompositeStructName, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsAccordionStart", inst.IsAccordionStart, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("AccordionName", inst.AccordionName, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsAccordionEnd", inst.IsAccordionEnd, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsType", inst.IsType, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.GongStructs {
				for _, target := range src.PointerToGongStructFields {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.GongStruct", "PointerToGongStructFields", refNames, formGroup, probe.formStage)
		}
	case *models.SliceOfPointerToGongStructField:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("GongStruct", inst.GongStruct, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.GongStruct](), probe.formStage)
		StageSetBasicFieldtoForm("Index", inst.Index, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("CompositeStructName", inst.CompositeStructName, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsAccordionStart", inst.IsAccordionStart, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("AccordionName", inst.AccordionName, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsAccordionEnd", inst.IsAccordionEnd, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.GongStructs {
				for _, target := range src.SliceOfPointerToGongStructFields {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.GongStruct", "SliceOfPointerToGongStructFields", refNames, formGroup, probe.formStage)
		}
	case *models.StageSetField:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("PackageName", inst.PackageName, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("PackagePath", inst.PackagePath, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsLocal", inst.IsLocal, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ImportAlias", inst.ImportAlias, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.StageSetModels {
				for _, target := range src.Fields {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.StageSetModel", "Fields", refNames, formGroup, probe.formStage)
		}
	case *models.StageSetModel:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: Fields
			div := (&form.FormDiv{Name: "Fields"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Fields {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Fields",
				Label: "Fields",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsManual", inst.IsManual, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ModelPkgs {
				if src.StageSet == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ModelPkg", "StageSet", refNames, formGroup, probe.formStage)
		}
	default:
		_ = inst
	}
}
