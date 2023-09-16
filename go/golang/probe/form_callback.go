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

	"{{PkgPathRoot}}/models"
)

const __dummmy__time = time.Nanosecond

// insertion point{{` + string(rune(FillUpTreeStructCase)) + `}}
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
	{{structname}} *models.{{Structname}},
	playground *Playground,
) ({{structname}}FormCallback *{{Structname}}FormCallback) {
	{{structname}}FormCallback = new({{Structname}}FormCallback)
	{{structname}}FormCallback.playground = playground
	{{structname}}FormCallback.{{structname}} = {{structname}}

	{{structname}}FormCallback.CreationMode = ({{structname}} == nil)

	return
}

type {{Structname}}FormCallback struct {
	{{structname}} *models.{{Structname}}

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func ({{structname}}FormCallback *{{Structname}}FormCallback) OnSave() {

	log.Println("{{Structname}}FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	{{structname}}FormCallback.playground.formStage.Checkout()

	if {{structname}}FormCallback.{{structname}} == nil {
		{{structname}}FormCallback.{{structname}} = new(models.{{Structname}}).Stage({{structname}}FormCallback.playground.stageOfInterest)
	}
	{{structname}}_ := {{structname}}FormCallback.{{structname}}
	_ = {{structname}}_

	// get the formGroup
	formGroup := {{structname}}FormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field{{fieldToFormCode}}
		}
	}

	{{structname}}FormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.{{Structname}}](
		{{structname}}FormCallback.playground,
	)
	{{structname}}FormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if {{structname}}FormCallback.CreationMode {
		{{structname}}FormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: New{{Structname}}FormCallback(
				nil,
				{{structname}}FormCallback.playground,
			),
		}).Stage({{structname}}FormCallback.playground.formStage)
		{{structname}} := new(models.{{Structname}})
		FillUpForm({{structname}}, newFormGroup, {{structname}}FormCallback.playground)
		{{structname}}FormCallback.playground.formStage.Commit()
	}

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
			FormDivSelectFieldToField(&({{structname}}_.{{FieldName}}), {{structname}}FormCallback.playground.stageOfInterest, formDiv)`,
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
