package probe

import (
	"fmt"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/fullstack-lang/gong/go/models"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const FillUpFormTemplate = `// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"{{PkgPathRoot}}/models"
	"{{PkgPathRoot}}/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point{{` + string(rune(ButtonImplPerGongstructCallToFormGenerator)) + `}}
	default:
		_ = instanceWithInferedType
	}
}
`

type FillUpFormInsertionId int

const (
	ButtonImplPerGongstructCallToFormGenerator FillUpFormInsertionId = iota
	FillUpFormInsertionNb
)

var FillUpFormSubTemplateCode map[FillUpFormInsertionId]string = // new line
map[FillUpFormInsertionId]string{
	ButtonImplPerGongstructCallToFormGenerator: `
	case *models.{{Structname}}:
		// insertion point{{fieldToFormCode}}
`,
}

type ButtonImplSubTemplateId int

const (
	ButtonImplSubTmplBasicField ButtonImplSubTemplateId = iota
	// ButtonImplFileFieldSubTmplSetBasicFieldInt
	ButtonImplSubTmplBasicFieldEnumString
	ButtonImplSubTmplBasicFieldEnumInt
	// ButtonImplFileFieldSubTmplSetBasicFieldFloat64
	// ButtonImplFileFieldSubTmplSetBasicFieldString
	// ButtonImplFileFieldSubTmplSetBasicFieldStringDocLink
	// ButtonImplFileFieldSubTmplSetTimeField
	ButtonImplSubTmplPointerField
	ButtonImplSubTmplSliceOfPointersField
	ButtonImplSubTmplSliceOfPointersReversePointer
)

var ButtonImplFileFieldFieldSubTemplateCode map[ButtonImplSubTemplateId]string = // declaration of the sub templates
map[ButtonImplSubTemplateId]string{

	ButtonImplSubTmplBasicField: `
		BasicFieldtoForm("{{FieldName}}", instanceWithInferedType.{{FieldName}}, instanceWithInferedType, probe.formStage, formGroup, {{isTextArea}})`,
	ButtonImplSubTmplBasicFieldEnumString: `
		EnumTypeStringToForm("{{FieldName}}", instanceWithInferedType.{{FieldName}}, instanceWithInferedType, probe.formStage, formGroup)`,
	ButtonImplSubTmplBasicFieldEnumInt: `
		EnumTypeIntToForm("{{FieldName}}", instanceWithInferedType.{{FieldName}}, instanceWithInferedType, probe.formStage, formGroup)`,
	ButtonImplSubTmplPointerField: `
		AssociationFieldToForm("{{FieldName}}", instanceWithInferedType.{{FieldName}}, formGroup, probe)`,
	ButtonImplSubTmplSliceOfPointersField: `
		AssociationSliceToForm("{{FieldName}}", instanceWithInferedType, &instanceWithInferedType.{{FieldName}}, formGroup, probe)`,
	ButtonImplSubTmplSliceOfPointersReversePointer: `
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "{{AssocStructName}}"
			rf.Fieldname = "{{FieldName}}"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.{{AssocStructName}}),
					"{{FieldName}}",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.{{AssocStructName}}, *models.{{Structname}}](
					nil,
					"{{FieldName}}",
					instanceWithInferedType,
					formGroup,
					probe)
			}	
		}`,
}

func CodeGeneratorFillUpForm(
	mdlPkg *models.ModelPkg,
	pkgName string,
	pkgPath string,
	pkgGoPath string) {

	// generate the typescript file
	codeGO := FillUpFormTemplate

	subStructCodes := make(map[FillUpFormInsertionId]string)
	for subStructTemplate := range FillUpFormSubTemplateCode {
		subStructCodes[subStructTemplate] = ""
	}

	// sort gong structs per name (for reproductibility)
	gongStructs := []*models.GongStruct{}
	for _, _struct := range mdlPkg.GongStructs {
		gongStructs = append(gongStructs, _struct)
	}
	sort.Slice(gongStructs[:], func(i, j int) bool {
		return gongStructs[i].Name < gongStructs[j].Name
	})

	for _, gongStruct := range gongStructs {

		if !gongStruct.HasNameField() {
			continue
		}

		for subStructTemplate := range FillUpFormSubTemplateCode {

			fieldToFormCode := ""

			for _, field := range gongStruct.Fields {

				switch field := field.(type) {
				case *models.GongBasicField:

					switch field.GetBasicKind() {
					case types.String, types.Bool, types.Float64, types.Int, types.Int64:
						if field.GongEnum == nil {
							var isTextArea = "false"
							if field.IsTextArea {
								isTextArea = "true"
							}
							fieldToFormCode += models.Replace2(
								ButtonImplFileFieldFieldSubTemplateCode[ButtonImplSubTmplBasicField],
								"{{FieldName}}", field.Name,
								"{{isTextArea}}", isTextArea)
						} else {
							if field.GongEnum.Type == models.Int {
								fieldToFormCode += models.Replace1(
									ButtonImplFileFieldFieldSubTemplateCode[ButtonImplSubTmplBasicFieldEnumInt],
									"{{FieldName}}", field.Name)
							} else {
								fieldToFormCode += models.Replace1(
									ButtonImplFileFieldFieldSubTemplateCode[ButtonImplSubTmplBasicFieldEnumString],
									"{{FieldName}}", field.Name)
							}
						}
					default:
					}
				case *models.GongTimeField:
					fieldToFormCode += models.Replace2(
						ButtonImplFileFieldFieldSubTemplateCode[ButtonImplSubTmplBasicField],
						"{{FieldName}}", field.Name,
						"{{isTextArea}}", "false")

				case *models.PointerToGongStructField:
					fieldToFormCode += models.Replace1(
						ButtonImplFileFieldFieldSubTemplateCode[ButtonImplSubTmplPointerField],
						"{{FieldName}}", field.Name)
				case *models.SliceOfPointerToGongStructField:
					fieldToFormCode += models.Replace1(
						ButtonImplFileFieldFieldSubTemplateCode[ButtonImplSubTmplSliceOfPointersField],
						"{{FieldName}}", field.Name)
				default:
				}

			}

			//
			// Parse all fields from other structs that points to this struct
			//
			for _, __struct := range gongStructs {
				for _, field := range __struct.Fields {
					switch field := field.(type) {
					case *models.SliceOfPointerToGongStructField:

						if field.GongStruct == gongStruct {
							fieldToFormCode += models.Replace2(
								ButtonImplFileFieldFieldSubTemplateCode[ButtonImplSubTmplSliceOfPointersReversePointer],
								"{{AssocStructName}}", __struct.Name,
								"{{FieldName}}", field.GetName())
						}
					}
				}
			}

			fieldToFormCode = models.Replace2(fieldToFormCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			generatedCodeFromSubTemplate := models.Replace3(FillUpFormSubTemplateCode[subStructTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{fieldToFormCode}}", fieldToFormCode)

			subStructCodes[subStructTemplate] += generatedCodeFromSubTemplate
		}

	}

	// substitutes {{<<insertionPerStructId points>>}} stuff with generated code
	for insertionPerStructId := FillUpFormInsertionId(0); insertionPerStructId < FillUpFormInsertionNb; insertionPerStructId++ {
		toReplace := "{{" + string(rune(insertionPerStructId)) + "}}"
		codeGO = strings.ReplaceAll(codeGO, toReplace, subStructCodes[insertionPerStructId])
	}

	caserEnglish := cases.Title(language.English)
	codeGO = models.Replace5(codeGO,
		"{{PkgName}}", pkgName,
		"{{TitlePkgName}}", caserEnglish.String(pkgName),
		"{{pkgname}}", strings.ToLower(pkgName),
		"	 | ", "	", // for the replacement of the of the first bar in the Gongstruct Type def

		"{{PkgPathRoot}}", strings.ReplaceAll(pkgGoPath, "/models", ""),
	)

	file, err := os.Create(filepath.Join(pkgPath, "../probe/fill_up_form.go"))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeGO)

}
