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

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"{{PkgPathRoot}}/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

type FormCallbackIF interface {
	GetCreationMode() bool
	GetInstance() any
	GetGongstructName() string
	OnSave()
}

type FormCallback[T models.PointerToGongstruct] struct {
	Instance     T
	CreationMode bool
	probe        *Probe
	formGroup    *form.FormGroup
	saveFields   func(instance T, probe *Probe, formGroup *form.FormGroup)
}

func NewFormCallback[T models.PointerToGongstruct](
	instance T,
	probe *Probe,
	formGroup *form.FormGroup,
	saveFields func(instance T, probe *Probe, formGroup *form.FormGroup),
) *FormCallback[T] {
	return &FormCallback[T]{
		Instance:     instance,
		CreationMode: any(instance) == nil,
		probe:        probe,
		formGroup:    formGroup,
		saveFields:   saveFields,
	}
}

func (cb *FormCallback[T]) GetCreationMode() bool     { return cb.CreationMode }
func (cb *FormCallback[T]) GetInstance() any           { return cb.Instance }
func (cb *FormCallback[T]) GetGongstructName() string { return models.GetPointerToGongstructName[T]() }

func (cb *FormCallback[T]) OnSave() {
	cb.probe.stageOfInterest.Lock()
	defer cb.probe.stageOfInterest.Unlock()

	cb.probe.formStage.Checkout()

	if any(cb.Instance) == nil {
		cb.Instance = cb.probe.stageOfInterest.GongNewInstance[T]()
	}

	cb.saveFields(cb.Instance, cb.probe, cb.formGroup)

	if cb.formGroup.HasSuppressButtonBeenPressed {
		cb.Instance.UnstageVoid(cb.probe.stageOfInterest)
	}

	cb.probe.stageOfInterest.Commit()
	updateProbeTable[T](cb.probe)

	if cb.CreationMode || cb.formGroup.HasSuppressButtonBeenPressed {
		cb.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(cb.probe.formStage)
		newFormGroup.OnSave = NewFormCallback[T](
			*new(T),
			cb.probe,
			newFormGroup,
			cb.saveFields,
		)
		newInstance := models.GongNewInstance[T]()
		FillUpForm(newInstance, newFormGroup, cb.probe)
		cb.probe.formStage.Commit()
	}

	cb.probe.ux_tree()
}

// insertion point{{` + string(rune(FormCallbackPerGongstructCode)) + `}}
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
	_instance *models.{{Structname}},
	probe *Probe,
	formGroup *form.FormGroup,
) ({{structname}}FormCallback *FormCallback[*models.{{Structname}}]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		save{{Structname}}Fields,
	)
}

type {{Structname}}FormCallback = FormCallback[*models.{{Structname}}]

func save{{Structname}}Fields(
	_instance *models.{{Structname}},
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field{{fieldToFormCode}}
		}
	}
}
`,
}

type FormCallbackSubTemplateId int

const (
	FormCallbackSubTmplBasicField FormCallbackSubTemplateId = iota
	FormCallbackSubTmplTimeField
	FormCallbackSubTmplBasicFieldEnumString
	FormCallbackSubTmplBasicFieldEnumInt
	FormCallbackSubTmplPointerToStruct
	FormCallbackSubTmplSliceOfPointers
	FormCallbackSubTmplSliceOfPointersReversePointer
)

var FormCallbackFileFieldFieldSubTemplateCode map[FormCallbackSubTemplateId]string = // declaration of the sub templates
map[FormCallbackSubTemplateId]string{

	FormCallbackSubTmplBasicField: `
		case "{{FieldName}}":
			FormDivBasicFieldToField(&(_instance.{{FieldName}}), formDiv)`,
	FormCallbackSubTmplTimeField: `
		case "{{FieldName}}":
			FormDivTimeFieldToField(&(_instance.{{FieldName}}), formDiv, {{isTimeFormOnly}})`,
	FormCallbackSubTmplBasicFieldEnumString: `
		case "{{FieldName}}":
			FormDivEnumStringFieldToField(&(_instance.{{FieldName}}), formDiv)`,
	FormCallbackSubTmplBasicFieldEnumInt: `
		case "{{FieldName}}":
			FormDivEnumIntFieldToField(&(_instance.{{FieldName}}), formDiv)`,
	FormCallbackSubTmplPointerToStruct: `
		case "{{FieldName}}":
			FormDivSelectFieldToField(&(_instance.{{FieldName}}), probe.stageOfInterest, formDiv)`,
	FormCallbackSubTmplSliceOfPointers: `
		case "{{FieldName}}":
			FormDivSliceOfPointersToField(_instance, "{{FieldName}}", &(_instance.{{FieldName}}), formDiv, probe)`,
	FormCallbackSubTmplSliceOfPointersReversePointer: `
		case "{{AssocStructName}}:{{FieldName}}":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "{{FieldName}}", func(owner *models.{{AssocStructName}}) *[]*models.{{Structname}} { return &owner.{{FieldName}} })`,
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
					fieldToFormCode += models.Replace2(
						FormCallbackFileFieldFieldSubTemplateCode[FormCallbackSubTmplTimeField],
						"{{FieldName}}", field.Name,
						"{{isTimeFormOnly}}", fmt.Sprintf("%t", field.TimeFormOnly))
				case *models.PointerToGongStructField:
					fieldToFormCode += models.Replace1(
						FormCallbackFileFieldFieldSubTemplateCode[FormCallbackSubTmplPointerToStruct],
						"{{FieldName}}", field.Name)
				case *models.SliceOfPointerToGongStructField:
					fieldToFormCode += models.Replace2(
						FormCallbackFileFieldFieldSubTemplateCode[FormCallbackSubTmplSliceOfPointers],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name)
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
	for insertionPerStructId := range FormCallbackGongstructInsertionNb {
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

	file, err := os.Create(filepath.Join(pkgPath, "probe/form_callback.go"))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeGO)

}
