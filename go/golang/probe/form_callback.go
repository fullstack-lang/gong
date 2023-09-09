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

const FormCallbackGongstructFileTemplate = `// generated code - do not edit
package probe

import (
	"log"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"
	"github.com/gin-gonic/gin"

	"{{PkgPathRoot}}/models"
	"{{PkgPathRoot}}/orm"
)

const __dummmy__time = time.Nanosecond

// insertion point{{` + string(rune(ButtonImplPerGongstructCallToForm)) + `}}
`

type FormCallbackGongstructInsertionId int

const (
	FormCallbackPerGongstructCode FormCallbackGongstructInsertionId = iota
	FormCallbackGongstructInsertionNb
)

var FormCallbackGongstructSubTemplateCode map[FormCallbackGongstructInsertionId]string = // new line
map[FormCallbackGongstructInsertionId]string{
	FormCallbackPerGongstructCode: `
func New{{Structname}}FormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	{{structname}} *models.{{Structname}},
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) ({{structname}}FormCallback *{{Structname}}FormCallback) {
	{{structname}}FormCallback = new({{Structname}}FormCallback)
	{{structname}}FormCallback.stageOfInterest = stageOfInterest
	{{structname}}FormCallback.tableStage = tableStage
	{{structname}}FormCallback.formStage = formStage
	{{structname}}FormCallback.{{structname}} = {{structname}}
	{{structname}}FormCallback.r = r
	{{structname}}FormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type {{Structname}}FormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	{{structname}}            *models.{{Structname}}
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func ({{structname}}FormCallback *{{Structname}}FormCallback) OnSave() {

	log.Println("{{Structname}}FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	{{structname}}FormCallback.formStage.Checkout()

	if {{structname}}FormCallback.{{structname}} == nil {
		{{structname}}FormCallback.{{structname}} = new(models.{{Structname}}).Stage({{structname}}FormCallback.stageOfInterest)
	}
	{{structname}}_ := {{structname}}FormCallback.{{structname}}
	_ = {{structname}}_

	// get the formGroup
	formGroup := {{structname}}FormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field{{fieldToFormCode}}
		}
	}

	{{structname}}FormCallback.stageOfInterest.Commit()
	fillUpTable[models.{{Structname}}](
		{{structname}}FormCallback.stageOfInterest,
		{{structname}}FormCallback.tableStage,
		{{structname}}FormCallback.formStage,
		{{structname}}FormCallback.r,
		{{structname}}FormCallback.backRepoOfInterest,
	)
	{{structname}}FormCallback.tableStage.Commit()
}`,
}

type FormCallbackSubTemplateId int

const (
	FormCallbackSubTmplBasicField FormCallbackSubTemplateId = iota
	FormCallbackSubTmplBasicFieldEnumString
	FormCallbackSubTmplBasicFieldEnumInt
	FormCallbackSubTmplPointerToStruct
)

var FormCallbackFileFieldFieldSubTemplateCode map[FormCallbackSubTemplateId]string = // declaration of the sub templates
map[FormCallbackSubTemplateId]string{

	FormCallbackSubTmplBasicField: `
		case "{{FieldName}}":
			FormDivBasicFieldToField(&({{structname}}_.{{FieldName}}), formDiv)`,
	FormCallbackSubTmplBasicFieldEnumString: `
		case "{{FieldName}}":
			FormDivEnumStringFieldToField(&({{structname}}_.{{FieldName}}), formDiv)`,
	FormCallbackSubTmplBasicFieldEnumInt: `
		case "{{FieldName}}":
			FormDivEnumIntFieldToField(&({{structname}}_.{{FieldName}}), formDiv)`,
	FormCallbackSubTmplPointerToStruct: `
		case "{{FieldName}}":
			FormDivSelectFieldToField(&({{structname}}_.{{FieldName}}), {{structname}}FormCallback.stageOfInterest, formDiv)`,
}

func CodeGeneratorModelFormCallback(
	mdlPkg *models.ModelPkg,
	pkgName string,
	pkgPath string,
	pkgGoPath string) {

	// generate the typescript file
	codeGO := FormCallbackGongstructFileTemplate

	subStructCodes := make(map[FormCallbackGongstructInsertionId]string)
	for subStructTemplate := range FormCallbackGongstructSubTemplateCode {
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

		for subStructTemplate := range FormCallbackGongstructSubTemplateCode {

			fieldToFormCode := ""

			for _, field := range gongStruct.Fields {

				switch field := field.(type) {
				case *models.GongBasicField:

					switch field.GetBasicKind() {
					case types.String, types.Bool, types.Float64, types.Int, types.Int64:
						if field.GongEnum == nil {
							fieldToFormCode += models.Replace1(
								FormCallbackFileFieldFieldSubTemplateCode[FormCallbackSubTmplBasicField],
								"{{FieldName}}", field.Name)
						} else {
							if field.GongEnum.Type == models.Int {
								fieldToFormCode += models.Replace1(
									FormCallbackFileFieldFieldSubTemplateCode[FormCallbackSubTmplBasicFieldEnumInt],
									"{{FieldName}}", field.Name)
							} else {
								fieldToFormCode += models.Replace1(
									FormCallbackFileFieldFieldSubTemplateCode[FormCallbackSubTmplBasicFieldEnumString],
									"{{FieldName}}", field.Name)
							}
						}
					default:
					}
				case *models.PointerToGongStructField:
					fieldToFormCode += models.Replace1(
						FormCallbackFileFieldFieldSubTemplateCode[FormCallbackSubTmplPointerToStruct],
						"{{FieldName}}", field.Name)
				default:
				}

			}

			fieldToFormCode = models.Replace2(fieldToFormCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			generatedCodeFromSubTemplate := models.Replace3(FormCallbackGongstructSubTemplateCode[subStructTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{fieldToFormCode}}", fieldToFormCode)

			subStructCodes[subStructTemplate] += generatedCodeFromSubTemplate
		}

	}

	// substitutes {{<<insertionPerStructId points>>}} stuff with generated code
	for insertionPerStructId := FormCallbackGongstructInsertionId(0); insertionPerStructId < FormCallbackGongstructInsertionNb; insertionPerStructId++ {
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

	file, err := os.Create(filepath.Join(pkgPath, "../probe/form_callback.go"))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeGO)

}
