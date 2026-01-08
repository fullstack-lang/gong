package models

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/fullstack-lang/gong/go/models"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const GongSliceTemplate = `// generated code - do not edit
package models

import (
	"fmt"
	"strings"
	"time"
)

var __GongSliceTemplate_time__dummyDeclaration time.Duration
var _ = __GongSliceTemplate_time__dummyDeclaration

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct{{` + string(rune(GongSliceReverseMapCompute)) + `}}
}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct{{` + string(rune(GongSliceGetInstances)) + `}}
	return
}

// insertion point per named struct{{` + string(rune(GongSliceGongCopy)) + `}}
func (stage *Stage) ComputeDifference() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesStmt string
	_ = newInstancesStmt
	var fieldsEditStmt string
	_ = fieldsEditStmt
	var deletedInstancesStmt string
	_ = deletedInstancesStmt

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct{{` + string(rune(GongSliceGongComputeDifference)) + `}}

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {
		notif := newInstancesStmt+fieldsEditStmt+deletedInstancesStmt
		notif += fmt.Sprintf("\n\t// %s", time.Now().Format(time.RFC3339Nano))
		notif += fmt.Sprintf("\n\tstage.Commit()")
		if stage.GetProbeIF() != nil {
			stage.GetProbeIF().AddNotification(
				time.Now(),
				notif,
			)
			stage.GetProbeIF().CommitNotificationTable()
		}
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct{{` + string(rune(GongSliceGongComputeReference)) + `}}
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct{{` + string(rune(GongSliceGongGetOrder)) + `}}
// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct{{` + string(rune(GongSliceGongGetIdentifier)) + `}}
// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct{{` + string(rune(GongSliceMarshallDeclaration)) + `}}

// insertion point for unstaging{{` + string(rune(GongSliceMarshallUnstaging)) + `}}
`

type GongSliceGongstructInsertionId int

const (
	GongSliceCase GongSliceGongstructInsertionId = iota
	GongSliceReverseMapCompute
	GongSliceGetInstances
	GongSliceGongCopy
	GongSliceGongComputeDifference
	GongSliceGongComputeReference
	GongSliceGongGetOrder
	GongSliceGongGetIdentifier
	GongSliceMarshallDeclaration
	GongSliceMarshallUnstaging
	GongSliceGongstructInsertionNb
)

var GongSliceGongstructSubTemplateCode map[GongSliceGongstructInsertionId]string = // new line
map[GongSliceGongstructInsertionId]string{
	GongSliceCase: `
	case *{{Structname}}:
		// insertion point per field{{perFieldCode}}
`,
	GongSliceReverseMapCompute: `
	// Compute reverse map for named struct {{Structname}}
	// insertion point per field{{sliceOfPointerFieldReverseMapComputationCode}}
`,
	GongSliceGetInstances: `
	for instance := range stage.{{Structname}}s {
		res = append(res, instance)
	}
`,
	GongSliceGongCopy: `
func ({{structname}} *{{Structname}}) GongCopy() GongstructIF {
	newInstance := *{{structname}}
	return &newInstance
}
`,
	GongSliceGongComputeDifference: `
	var {{structname}}s_newInstances []*{{Structname}}
	var {{structname}}s_deletedInstances []*{{Structname}}

	// parse all staged instances and check if they have a reference
	for {{structname}} := range stage.{{Structname}}s {
		if ref, ok := stage.{{Structname}}s_reference[{{structname}}]; !ok {
			{{structname}}s_newInstances = append({{structname}}s_newInstances, {{structname}})
			newInstancesStmt += {{structname}}.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := {{structname}}.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := {{structname}}.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", {{structname}}.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for {{structname}} := range stage.{{Structname}}s_reference {
		if _, ok := stage.{{Structname}}s[{{structname}}]; !ok {
			{{structname}}s_deletedInstances = append({{structname}}s_deletedInstances, {{structname}})
			deletedInstancesStmt += {{structname}}.GongMarshallUnstaging(stage)
		}
	}

	lenNewInstances += len({{structname}}s_newInstances)
	lenDeletedInstances += len({{structname}}s_deletedInstances)`,

	GongSliceGongComputeReference: `
	stage.{{Structname}}s_reference = make(map[*{{Structname}}]*{{Structname}})
	for instance := range stage.{{Structname}}s {
		stage.{{Structname}}s_reference[instance] = instance.GongCopy().(*{{Structname}})
	}
`,

	GongSliceGongGetOrder: `
func ({{structname}} *{{Structname}}) GongGetOrder(stage *Stage) uint {
	return stage.{{Structname}}Map_Staged_Order[{{structname}}]
}
`,
	GongSliceGongGetIdentifier: `
func ({{structname}} *{{Structname}}) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", {{structname}}.GongGetGongstructName(), {{structname}}.GongGetOrder(stage))
}
`,
	GongSliceMarshallDeclaration: `
func ({{structname}} *{{Structname}}) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", {{structname}}.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "{{Structname}}")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", {{structname}}.Name)
	return
}`,
	GongSliceMarshallUnstaging: `
func ({{structname}} *{{Structname}}) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", {{structname}}.GongGetIdentifier(stage))
	return
}`,
}

type GongSliceSubTemplateId int

const (
	GongSliceSubTmplSliceOfPointersToStruct GongSliceSubTemplateId = iota
	GongSliceSubTmplSliceOfPointersReverseMapComputation
)

var GongSliceFileFieldFieldSubTemplateCode map[GongSliceSubTemplateId]string = // declaration of the sub templates
map[GongSliceSubTemplateId]string{

	GongSliceSubTmplSliceOfPointersToStruct: `
		if fieldName == "{{FieldName}}" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*{{Structname}}) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*{{Structname}})
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.{{FieldName}}).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.{{FieldName}} = _inferedTypeInstance.{{FieldName}}[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.{{FieldName}} =
								append(_inferedTypeInstance.{{FieldName}}, any(fieldInstance).(*{{AssociationStructName}}))
						}
					}
				}
			}
		}`,
	GongSliceSubTmplSliceOfPointersReverseMapComputation: `
	stage.{{Structname}}_{{FieldNameForReverseMapField}}_reverseMap = make(map[*{{AssociationStructName}}]*{{Structname}})
	for {{structname}} := range stage.{{Structname}}s {
		_ = {{structname}}
		for _, _{{associationStructName}} := range {{structname}}.{{FieldName}} {
			stage.{{Structname}}_{{FieldNameForReverseMapField}}_reverseMap[_{{associationStructName}}] = {{structname}}
		}
	}`,
}

func CodeGeneratorModelGongSlice(
	modelPkg *models.ModelPkg,
	pkgName string,
	pkgPath string,
	pkgGoPath string) {

	// this code is not robust to empty models
	// map[Gongstruct]any cannot compile
	if len(modelPkg.GongStructs) == 0 {
		return
	}

	// generate the typescript file
	codeGO := GongSliceTemplate

	subStructCodes := make(map[GongSliceGongstructInsertionId]string)
	for subStructTemplate := range GongSliceGongstructSubTemplateCode {
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

		for subStructTemplate := range GongSliceGongstructSubTemplateCode {

			perFieldCode := ""
			sliceOfPointerFieldReverseMapComputationCode := ""

			for _, field := range gongStruct.Fields {

				fieldName := field.GetName()
				fieldNameForReverseMapField := fieldName

				// in case of a field within an anonymous struct, one needs
				// to strip the prefix
				fieldNameSplitted := strings.Split(fieldName, ".")
				isWithinAnonymousStruct := len(fieldNameSplitted) > 1
				if isWithinAnonymousStruct {
					fieldNameForReverseMapField = fieldNameSplitted[0] + "_" + fieldNameSplitted[1]
				}

				switch field := field.(type) {
				case *models.SliceOfPointerToGongStructField:
					perFieldCode += models.Replace3(
						GongSliceFileFieldFieldSubTemplateCode[GongSliceSubTmplSliceOfPointersToStruct],
						"{{FieldName}}", field.Name,
						"{{AssociationStructName}}", field.GongStruct.Name,
						"{{associationStructName}}", strings.ToLower(field.GongStruct.Name))
					sliceOfPointerFieldReverseMapComputationCode += models.Replace4(
						GongSliceFileFieldFieldSubTemplateCode[GongSliceSubTmplSliceOfPointersReverseMapComputation],
						"{{FieldNameForReverseMapField}}", fieldNameForReverseMapField,
						"{{FieldName}}", fieldName,
						"{{AssociationStructName}}", field.GongStruct.Name,
						"{{associationStructName}}", strings.ToLower(field.GongStruct.Name))

				default:
				}

			}

			perFieldCode = models.Replace2(perFieldCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			sliceOfPointerFieldReverseMapComputationCode = models.Replace2(sliceOfPointerFieldReverseMapComputationCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			generatedCodeFromSubTemplate := models.Replace4(GongSliceGongstructSubTemplateCode[subStructTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{perFieldCode}}", perFieldCode,
				"{{sliceOfPointerFieldReverseMapComputationCode}}", sliceOfPointerFieldReverseMapComputationCode)

			// case where the i
			if GongSliceMarshallDeclaration == subStructTemplate {
				// find the "Name" field
				nameFieldIsEmbedded := false
				for _, field := range gongStruct.Fields {
					if field.GetName() == "Name" && field.GetCompositeStructName() != "" {
						nameFieldIsEmbedded = true
					}
				}
				if nameFieldIsEmbedded {
					// go does not hanldledirect reference to embedded fields in struct literals (https://github.com/golang/go/issues/9859)
					// therefore, we simplify the code generation
					generatedCodeFromSubTemplate = strings.ReplaceAll(generatedCodeFromSubTemplate, "GongIdentifiersDecls", "IdentifiersDeclsWithoutNameInit")
				}
			}

			subStructCodes[subStructTemplate] += generatedCodeFromSubTemplate
		}

	}

	// substitutes {{<<insertionPerStructId points>>}} stuff with generated code
	for insertionPerStructId := GongSliceGongstructInsertionId(0); insertionPerStructId < GongSliceGongstructInsertionNb; insertionPerStructId++ {
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

	file, err := os.Create(filepath.Join(pkgPath, string(models.GeneratedGongSlicesGoFilePath)))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeGO)

}
