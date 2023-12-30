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
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"{{PkgPathRoot}}/models"
	"{{PkgPathRoot}}/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

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
func __gong__New__{{Structname}}FormCallback(
	{{structname}} *models.{{Structname}},
	probe *Probe,
	formGroup *table.FormGroup,
) ({{structname}}FormCallback *{{Structname}}FormCallback) {
	{{structname}}FormCallback = new({{Structname}}FormCallback)
	{{structname}}FormCallback.probe = probe
	{{structname}}FormCallback.{{structname}} = {{structname}}
	{{structname}}FormCallback.formGroup = formGroup

	{{structname}}FormCallback.CreationMode = ({{structname}} == nil)

	return
}

type {{Structname}}FormCallback struct {
	{{structname}} *models.{{Structname}}

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func ({{structname}}FormCallback *{{Structname}}FormCallback) OnSave() {

	log.Println("{{Structname}}FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	{{structname}}FormCallback.probe.formStage.Checkout()

	if {{structname}}FormCallback.{{structname}} == nil {
		{{structname}}FormCallback.{{structname}} = new(models.{{Structname}}).Stage({{structname}}FormCallback.probe.stageOfInterest)
	}
	{{structname}}_ := {{structname}}FormCallback.{{structname}}
	_ = {{structname}}_

	for _, formDiv := range {{structname}}FormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field{{fieldToFormCode}}
		}
	}

	// manage the suppress operation
	if {{structname}}FormCallback.formGroup.HasSuppressButtonBeenPressed {
		{{structname}}_.Unstage({{structname}}FormCallback.probe.stageOfInterest)
	}

	{{structname}}FormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.{{Structname}}](
		{{structname}}FormCallback.probe,
	)
	{{structname}}FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if {{structname}}FormCallback.CreationMode || {{structname}}FormCallback.formGroup.HasSuppressButtonBeenPressed {
		{{structname}}FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage({{structname}}FormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__{{Structname}}FormCallback(
			nil,
			{{structname}}FormCallback.probe,
			newFormGroup,
		)
		{{structname}} := new(models.{{Structname}})
		FillUpForm({{structname}}, newFormGroup, {{structname}}FormCallback.probe)
		{{structname}}FormCallback.probe.formStage.Commit()
	}

	fillUpTree({{structname}}FormCallback.probe)
}`,
}

type FormCallbackSubTemplateId int

const (
	FormCallbackSubTmplBasicField FormCallbackSubTemplateId = iota
	FormCallbackSubTmplBasicFieldEnumString
	FormCallbackSubTmplBasicFieldEnumInt
	FormCallbackSubTmplPointerToStruct
	FormCallbackSubTmplSliceOfPointersReversePointer
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
			FormDivSelectFieldToField(&({{structname}}_.{{FieldName}}), {{structname}}FormCallback.probe.stageOfInterest, formDiv)`,
	FormCallbackSubTmplSliceOfPointersReversePointer: `
		case "{{AssocStructName}}:{{FieldName}}":
			// we need to retrieve the field owner before the change
			var past{{AssocStructName}}Owner *models.{{AssocStructName}}
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "{{AssocStructName}}"
			rf.Fieldname = "{{FieldName}}"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				{{structname}}FormCallback.probe.stageOfInterest,
				{{structname}}FormCallback.probe.backRepoOfInterest,
				{{structname}}_,
				&rf)

			if reverseFieldOwner != nil {
				past{{AssocStructName}}Owner = reverseFieldOwner.(*models.{{AssocStructName}})
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if past{{AssocStructName}}Owner != nil {
					idx := slices.Index(past{{AssocStructName}}Owner.{{FieldName}}, {{structname}}_)
					past{{AssocStructName}}Owner.{{FieldName}} = slices.Delete(past{{AssocStructName}}Owner.{{FieldName}}, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _{{assocStructName}} := range *models.GetGongstructInstancesSet[models.{{AssocStructName}}]({{structname}}FormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _{{assocStructName}}.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						new{{AssocStructName}}Owner := _{{assocStructName}} // we have a match
						if past{{AssocStructName}}Owner != nil {
							if new{{AssocStructName}}Owner != past{{AssocStructName}}Owner {
								idx := slices.Index(past{{AssocStructName}}Owner.{{FieldName}}, {{structname}}_)
								past{{AssocStructName}}Owner.{{FieldName}} = slices.Delete(past{{AssocStructName}}Owner.{{FieldName}}, idx, idx+1)
								new{{AssocStructName}}Owner.{{FieldName}} = append(new{{AssocStructName}}Owner.{{FieldName}}, {{structname}}_)
							}
						} else {
							new{{AssocStructName}}Owner.{{FieldName}} = append(new{{AssocStructName}}Owner.{{FieldName}}, {{structname}}_)
						}
					}
				}
			}`,
}

func CodeGeneratorModelFormCallback(
	modelPkg *models.ModelPkg,
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
	for _, _struct := range modelPkg.GongStructs {
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
				case *models.GongTimeField:
					fieldToFormCode += models.Replace1(
						FormCallbackFileFieldFieldSubTemplateCode[FormCallbackSubTmplBasicField],
						"{{FieldName}}", field.Name)
				case *models.PointerToGongStructField:
					fieldToFormCode += models.Replace1(
						FormCallbackFileFieldFieldSubTemplateCode[FormCallbackSubTmplPointerToStruct],
						"{{FieldName}}", field.Name)
				default:
				}

			}

			for _, __struct := range gongStructs {

				for _, field := range __struct.Fields {
					switch field := field.(type) {
					case *models.SliceOfPointerToGongStructField:
						if field.GongStruct == gongStruct {
							fieldToFormCode += models.Replace3(
								FormCallbackFileFieldFieldSubTemplateCode[FormCallbackSubTmplSliceOfPointersReversePointer],
								"{{AssocStructName}}", __struct.Name,
								"{{assocStructName}}", strings.ToLower(__struct.Name),
								"{{FieldName}}", field.Name)
						}
					}
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
