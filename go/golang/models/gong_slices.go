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
	"sort"
	"strings"
	"time"
)

var (
	__GongSliceTemplate_time__dummyDeclaration time.Duration
	_                                          = __GongSliceTemplate_time__dummyDeclaration
)

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct{{` + string(rune(GongSliceReverseMapCompute)) + `}}
	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct{{` + string(rune(GongSliceGetInstances)) + `}}
	return
}

// insertion point per named struct{{` + string(rune(GongSliceGongCopy)) + `}}
func (stage *Stage) ComputeForwardAndBackwardCommits() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct{{` + string(rune(GongSliceGongComputeDifference)) + `}}

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.forwardCommits = append(stage.forwardCommits, forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.backwardCommits = append(stage.backwardCommits, backwardCommit)
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct{{` + string(rune(GongSliceGongComputeReference)) + `}}
	stage.recomputeOrders()
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
// end of template
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
			newInstancesSlice = append(newInstancesSlice, {{structname}}.GongMarshallIdentifier(stage))
			if stage.{{Structname}}s_referenceOrder == nil {
				stage.{{Structname}}s_referenceOrder = make(map[*{{Structname}}]uint)
			}
			stage.{{Structname}}s_referenceOrder[{{structname}}] = stage.{{Structname}}Map_Staged_Order[{{structname}}]
			newInstancesReverseSlice = append(newInstancesReverseSlice, {{structname}}.GongMarshallUnstaging(stage))
			delete(stage.{{Structname}}s_referenceOrder, {{structname}})
			fieldInitializers, pointersInitializations := {{structname}}.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.{{Structname}}Map_Staged_Order[ref] = stage.{{Structname}}Map_Staged_Order[{{structname}}]
			diffs := {{structname}}.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, {{structname}})
			delete(stage.{{Structname}}Map_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", {{structname}}.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.{{Structname}}s_reference {
		if _, ok := stage.{{Structname}}s[ref]; !ok {
			{{structname}}s_deletedInstances = append({{structname}}s_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len({{structname}}s_newInstances)
	lenDeletedInstances += len({{structname}}s_deletedInstances)`,

	GongSliceGongComputeReference: `
	stage.{{Structname}}s_reference = make(map[*{{Structname}}]*{{Structname}})
	stage.{{Structname}}s_referenceOrder = make(map[*{{Structname}}]uint) // diff Unstage needs the reference order
	for instance := range stage.{{Structname}}s {
		stage.{{Structname}}s_reference[instance] = instance.GongCopy().(*{{Structname}})
		stage.{{Structname}}s_referenceOrder[instance] = instance.GongGetOrder(stage)
	}
`,

	GongSliceGongGetOrder: `
func ({{structname}} *{{Structname}}) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.{{Structname}}Map_Staged_Order[{{structname}}]; ok {
		return order
	}
	return stage.{{Structname}}s_referenceOrder[{{structname}}]
}

func ({{structname}} *{{Structname}}) GongGetReferenceOrder(stage *Stage) uint {
	return stage.{{Structname}}s_referenceOrder[{{structname}}]
}
`,
	GongSliceGongGetIdentifier: `
func ({{structname}} *{{Structname}}) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", {{structname}}.GongGetGongstructName(), {{structname}}.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func ({{structname}} *{{Structname}}) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", {{structname}}.GongGetGongstructName(), {{structname}}.GongGetReferenceOrder(stage))
}
`,
	GongSliceMarshallDeclaration: `
func ({{structname}} *{{Structname}}) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", {{structname}}.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "{{Structname}}")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", {{structname}}.Name)
	return
}
`,
	GongSliceMarshallUnstaging: `
func ({{structname}} *{{Structname}}) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", {{structname}}.GongGetReferenceIdentifier(stage))
	return
}
`,
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
	pkgGoPath string,
) {
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
