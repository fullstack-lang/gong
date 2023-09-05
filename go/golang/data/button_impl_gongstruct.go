package data

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

const ButtonImplGongstructFileTemplate = `// generated code - do not edit
package data

import (
	"log"

	"github.com/gin-gonic/gin"

	gong_models "github.com/fullstack-lang/gong/go/models"
	form "github.com/fullstack-lang/gongtable/go/models"
	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"{{PkgPathRoot}}/models"
	"{{PkgPathRoot}}/orm"
)

type ButtonImplGongstruct struct {
	gongStruct         *gong_models.GongStruct
	Icon               gongtree_buttons.ButtonType
	tableStage         *form.StageStruct
	formStage          *form.StageStruct
	stageOfInterest    *models.StageStruct
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func NewButtonImplGongstruct(
	gongStruct *gong_models.GongStruct,
	icon gongtree_buttons.ButtonType,
	tableStage *form.StageStruct,
	formStage *form.StageStruct,
	stageOfInterest *models.StageStruct,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (buttonImplGongstruct *ButtonImplGongstruct) {

	buttonImplGongstruct = new(ButtonImplGongstruct)
	buttonImplGongstruct.Icon = icon
	buttonImplGongstruct.gongStruct = gongStruct
	buttonImplGongstruct.tableStage = tableStage
	buttonImplGongstruct.formStage = formStage
	buttonImplGongstruct.stageOfInterest = stageOfInterest
	buttonImplGongstruct.r = r
	buttonImplGongstruct.backRepoOfInterest = backRepoOfInterest

	return
}

func (buttonImpl *ButtonImplGongstruct) ButtonUpdated(
	gongtreeStage *gongtree_models.StageStruct,
	stageButton, front *gongtree_models.Button) {

	log.Println("ButtonImplGongstruct: ButtonUpdated")

	formStage := buttonImpl.formStage
	formStage.Reset()
	formStage.Commit()

	switch buttonImpl.gongStruct.Name {
	// insertion point{{` + string(rune(ButtonImplPerGongstructCallToForm)) + `}}
	}
	formStage.Commit()
}

func FillUpForm[T models.Gongstruct](
	instance *T,
	stageOfInterest *models.StageStruct,
	formStage *form.StageStruct,
	formGroup *form.FormGroup,
	r *gin.Engine,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point{{` + string(rune(ButtonImplPerGongstructCallToFormGenerator)) + `}}
	default:
		_ = instanceWithInferedType
	}
}

`

type ButtonImplGongstructInsertionId int

const (
	ButtonImplPerGongstructCallToForm ButtonImplGongstructInsertionId = iota
	ButtonImplPerGongstructCallToFormGenerator
	ButtonImplGongstructInsertionNb
)

var ButtonImplGongstructSubTemplateCode map[ButtonImplGongstructInsertionId]string = // new line
map[ButtonImplGongstructInsertionId]string{
	ButtonImplPerGongstructCallToForm: `
	case "{{Structname}}":
		formGroup := (&form.FormGroup{
			Name: form.FormGroupDefaultName.ToString(),
			OnSave: New{{Structname}}FormCallback(
				buttonImpl.stageOfInterest,
				buttonImpl.tableStage,
				formStage,
				nil,
				buttonImpl.r,
				buttonImpl.backRepoOfInterest,
			),
		}).Stage(formStage)
		{{structname}} := new(models.{{Structname}})
		FillUpForm({{structname}}, buttonImpl.stageOfInterest, formStage, formGroup, buttonImpl.r)`,
	ButtonImplPerGongstructCallToFormGenerator: `
	case *models.{{Structname}}:
		`,
}

type ButtonImplFilePerStructSubTemplateId int

const (
	ButtonImplFileFieldSubTmplSetBasicFieldBool ButtonImplFilePerStructSubTemplateId = iota
	ButtonImplFileFieldSubTmplSetBasicFieldInt
	ButtonImplFileFieldSubTmplSetBasicFieldEnumString
	ButtonImplFileFieldSubTmplSetBasicFieldEnumInt
	ButtonImplFileFieldSubTmplSetBasicFieldFloat64
	ButtonImplFileFieldSubTmplSetBasicFieldString
	ButtonImplFileFieldSubTmplSetBasicFieldStringDocLink
	ButtonImplFileFieldSubTmplSetTimeField
	ButtonImplFileFieldSubTmplSetPointerField
	ButtonImplFileFieldSubTmplSetSliceOfPointersField
)

var ButtonImplFileFieldFieldSubTemplateCode map[ButtonImplFilePerStructSubTemplateId]string = // declaration of the sub templates
map[ButtonImplFilePerStructSubTemplateId]string{

	ButtonImplFileFieldSubTmplSetBasicFieldBool: `
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "{{FieldName}}")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", {{structname}}.{{FieldName}}))
		initializerStatements += setValueField
`,
	ButtonImplFileFieldSubTmplSetTimeField: `
		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "{{FieldName}}")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", {{structname}}.{{FieldName}}.String())
		initializerStatements += setValueField
`,
	ButtonImplFileFieldSubTmplSetBasicFieldInt: `
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "{{FieldName}}")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", {{structname}}.{{FieldName}}))
		initializerStatements += setValueField
`,
	ButtonImplFileFieldSubTmplSetBasicFieldEnumString: `
		if {{structname}}.{{FieldName}} != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "{{FieldName}}")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+{{structname}}.{{FieldName}}.ToCodeString())
			initializerStatements += setValueField
		}
`,
	ButtonImplFileFieldSubTmplSetBasicFieldEnumInt: `
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "{{FieldName}}")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+{{structname}}.{{FieldName}}.ToCodeString())
		initializerStatements += setValueField
`,
	ButtonImplFileFieldSubTmplSetBasicFieldFloat64: `
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "{{FieldName}}")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", {{structname}}.{{FieldName}}))
		initializerStatements += setValueField
`,
	ButtonImplFileFieldSubTmplSetBasicFieldString: `
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "{{FieldName}}")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string({{structname}}.{{FieldName}}))
		initializerStatements += setValueField
`,
	ButtonImplFileFieldSubTmplSetBasicFieldStringDocLink: `
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "\n\t{{Identifier}}",
			fmt.Sprintf("\n\n\t// comment added to overcome the problem with the comment map association\n\n\t//gong:ident [%s]\n\t{{Identifier}}",
				string({{structname}}.{{FieldName}})))
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "{{FieldName}}")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string({{structname}}.{{FieldName}}))
		initializerStatements += setValueField
`,
	ButtonImplFileFieldSubTmplSetPointerField: `
		if {{structname}}.{{FieldName}} != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "{{FieldName}}")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_{{AssocStructName}}_Identifiers[{{structname}}.{{FieldName}}])
			pointersInitializesStatements += setPointerField
		}
`,
	ButtonImplFileFieldSubTmplSetSliceOfPointersField: `
		for _, _{{assocstructname}} := range {{structname}}.{{FieldName}} {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "{{FieldName}}")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_{{AssocStructName}}_Identifiers[_{{assocstructname}}])
			pointersInitializesStatements += setPointerField
		}
`,
}

func CodeGeneratorModelButtonImpl(
	mdlPkg *models.ModelPkg,
	pkgName string,
	pkgPath string,
	pkgGoPath string) {

	// generate the typescript file
	codeGO := ButtonImplGongstructFileTemplate

	subStructCodes := make(map[ButtonImplGongstructInsertionId]string)
	for subStructTemplate := range ButtonImplGongstructSubTemplateCode {
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

		for subStructTemplate := range ButtonImplGongstructSubTemplateCode {

			valInitCode := ""
			pointerInitCode := ""

			for _, field := range gongStruct.Fields {

				switch field := field.(type) {
				case *models.GongBasicField:

					switch field.GetBasicKind() {
					case types.String:
						if field.GongEnum == nil {

							if !field.IsDocLink {
								valInitCode += models.Replace1(
									ButtonImplFileFieldFieldSubTemplateCode[ButtonImplFileFieldSubTmplSetBasicFieldString],
									"{{FieldName}}", field.Name)
							} else {
								valInitCode += models.Replace1(
									ButtonImplFileFieldFieldSubTemplateCode[ButtonImplFileFieldSubTmplSetBasicFieldStringDocLink],
									"{{FieldName}}", field.Name)
							}

						} else {
							valInitCode += models.Replace1(
								ButtonImplFileFieldFieldSubTemplateCode[ButtonImplFileFieldSubTmplSetBasicFieldEnumString],
								"{{FieldName}}", field.Name)
						}
					case types.Bool:
						valInitCode += models.Replace1(
							ButtonImplFileFieldFieldSubTemplateCode[ButtonImplFileFieldSubTmplSetBasicFieldBool],
							"{{FieldName}}", field.Name)
					case types.Float64:
						valInitCode += models.Replace1(
							ButtonImplFileFieldFieldSubTemplateCode[ButtonImplFileFieldSubTmplSetBasicFieldFloat64],
							"{{FieldName}}", field.Name)
					case types.Int, types.Int64:
						if field.GongEnum == nil {
							valInitCode += models.Replace1(
								ButtonImplFileFieldFieldSubTemplateCode[ButtonImplFileFieldSubTmplSetBasicFieldInt],
								"{{FieldName}}", field.Name)
						} else {
							valInitCode += models.Replace1(
								ButtonImplFileFieldFieldSubTemplateCode[ButtonImplFileFieldSubTmplSetBasicFieldEnumInt],
								"{{FieldName}}", field.Name)
						}
					default:
					}
				case *models.GongTimeField:
					valInitCode += models.Replace1(
						ButtonImplFileFieldFieldSubTemplateCode[ButtonImplFileFieldSubTmplSetTimeField],
						"{{FieldName}}", field.Name)
				case *models.PointerToGongStructField:
					pointerInitCode += models.Replace2(
						ButtonImplFileFieldFieldSubTemplateCode[ButtonImplFileFieldSubTmplSetPointerField],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name)
				case *models.SliceOfPointerToGongStructField:
					pointerInitCode += models.Replace3(
						ButtonImplFileFieldFieldSubTemplateCode[ButtonImplFileFieldSubTmplSetSliceOfPointersField],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name,
						"{{assocstructname}}", strings.ToLower(field.GongStruct.Name))
				default:
				}

			}

			valInitCode = models.Replace2(valInitCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			pointerInitCode = models.Replace2(pointerInitCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			generatedCodeFromSubTemplate := models.Replace4(ButtonImplGongstructSubTemplateCode[subStructTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{ValuesInitialization}}", valInitCode,
				"{{PointersInitialization}}", pointerInitCode)

			subStructCodes[subStructTemplate] += generatedCodeFromSubTemplate
		}

	}

	// substitutes {{<<insertionPerStructId points>>}} stuff with generated code
	for insertionPerStructId := ButtonImplGongstructInsertionId(0); insertionPerStructId < ButtonImplGongstructInsertionNb; insertionPerStructId++ {
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

	file, err := os.Create(filepath.Join(pkgPath, "../data/button_impl_gongstruct.go"))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeGO)

}
