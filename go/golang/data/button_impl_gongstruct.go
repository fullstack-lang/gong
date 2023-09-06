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
)

var ButtonImplFileFieldFieldSubTemplateCode map[ButtonImplSubTemplateId]string = // declaration of the sub templates
map[ButtonImplSubTemplateId]string{

	ButtonImplSubTmplBasicField: `
		BasicFieldtoForm("{{FieldName}}", instanceWithInferedType.{{FieldName}}, instanceWithInferedType, formStage, formGroup)`,
	ButtonImplSubTmplBasicFieldEnumString: `
		EnumTypeStringToForm("{{FieldName}}", instanceWithInferedType.{{FieldName}}, instanceWithInferedType, formStage, formGroup)`,
	ButtonImplSubTmplBasicFieldEnumInt: `
		EnumTypeIntToForm("{{FieldName}}", instanceWithInferedType.{{FieldName}}, instanceWithInferedType, formStage, formGroup)`,
	ButtonImplSubTmplPointerField: `
		AssociationFieldToForm("{{FieldName}}", instanceWithInferedType.{{FieldName}}, stageOfInterest, formStage, formGroup)`,
	ButtonImplSubTmplSliceOfPointersField: `
		AssociationSliceToForm("{{FieldName}}", instanceWithInferedType, &instanceWithInferedType.{{FieldName}}, stageOfInterest, formStage, formGroup, r)`,
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

			fieldToFormCode := ""

			for _, field := range gongStruct.Fields {

				switch field := field.(type) {
				case *models.GongBasicField:

					switch field.GetBasicKind() {
					case types.String, types.Bool, types.Float64, types.Int, types.Int64:
						if field.GongEnum == nil {
							fieldToFormCode += models.Replace1(
								ButtonImplFileFieldFieldSubTemplateCode[ButtonImplSubTmplBasicField],
								"{{FieldName}}", field.Name)
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

			fieldToFormCode = models.Replace2(fieldToFormCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			generatedCodeFromSubTemplate := models.Replace3(ButtonImplGongstructSubTemplateCode[subStructTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{fieldToFormCode}}", fieldToFormCode)

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
