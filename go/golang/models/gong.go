package models

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

// insertion points are places where the code is
// generated per gong struct
type ModelGongStructInsertionId int

const (
	ModelGongStructInsertionCommitCheckout ModelGongStructInsertionId = iota
	ModelGongStructInsertionStageFunctions
	ModelGongStructInsertionCreateCallback
	ModelGongStructInsertionDeleteCallback
	ModelGongStructInsertionArrayDefintion
	ModelGongStructInsertionArrayInitialisation
	ModelGongStructInsertionArrayReset
	ModelGongStructInsertionArrayNil
	ModelGongStructInsertionArrayUnstage
	ModelGongStructInsertionUnmarshallDeclarations
	ModelGongStructInsertionUnmarshallPointersInitializations
	ModelGongStructInsertionComputeNbInstances

	ModelGongStructInsertionGenericGetFields
	ModelGongStructInsertionGenericGetReverseFields
	ModelGongStructInsertionGenericGetFieldsFromPointer
	ModelGongStructInsertionGenericGetFieldValues
	ModelGongStructInsertionGenericGetFieldValuesFromPointer

	ModelGongStructInsertionGenericReversePointerAssociationsMaps
	ModelGongStructInsertionGenericReverseSliceOfPointersAssociationsMaps

	ModelGongStructInsertionGenericGongSetTypes
	ModelGongStructInsertionGenericGongstructName
	ModelGongStructInsertionGenericPointerToGongstructName
	ModelGongStructInsertionGenericGongMapTypes
	ModelGongStructInsertionGenericGetSetFunctions
	ModelGongStructInsertionGenericGetMapFunctions
	ModelGongStructInsertionGenericInstancesSetFunctions
	ModelGongStructInsertionGenericInstancesSetFromPointerTypeFunctions

	ModelGongStructInsertionGenericInstancesMapFunctions
	ModelGongStructInsertionGenericGetAssociationNameFunctions

	ModelGongStructInsertionGenericGongstructTypes
	ModelGongStructInsertionGenericPointerToGongstructTypes

	ModelGongStructInsertionsNb
)

var ModelGongStructSubTemplateCode map[ModelGongStructInsertionId]string = // new line
map[ModelGongStructInsertionId]string{
	ModelGongStructInsertionCommitCheckout: `
	Commit{{Structname}}({{structname}} *{{Structname}})
	Checkout{{Structname}}({{structname}} *{{Structname}})`,

	ModelGongStructInsertionGenericGetFields: `
	case {{Structname}}:{{ListOfFieldsName}}`,

	ModelGongStructInsertionGenericGetReverseFields: `
	case {{Structname}}:{{ListOfReverseFields}}`,

	ModelGongStructInsertionGenericGetFieldsFromPointer: `
	case *{{Structname}}:{{ListOfFieldsName}}`,

	ModelGongStructInsertionGenericGetFieldValues: `
	case {{Structname}}:
		switch fieldName {
		// string value of fields{{StringValueOfFields}}
		}`,

	ModelGongStructInsertionGenericGetFieldValuesFromPointer: `
	case *{{Structname}}:
		switch fieldName {
		// string value of fields{{StringValueOfFields}}
		}`,

	ModelGongStructInsertionStageFunctions: `
// Stage puts {{structname}} to the model stage
func ({{structname}} *{{Structname}}) Stage(stage *StageStruct) *{{Structname}} {
	stage.{{Structname}}s[{{structname}}] = __member
	stage.{{Structname}}s_mapString[{{structname}}.Name] = {{structname}}

	return {{structname}}
}

// Unstage removes {{structname}} off the model stage
func ({{structname}} *{{Structname}}) Unstage(stage *StageStruct) *{{Structname}} {
	delete(stage.{{Structname}}s, {{structname}})
	delete(stage.{{Structname}}s_mapString, {{structname}}.Name)
	return {{structname}}
}

// UnstageVoid removes {{structname}} off the model stage
func ({{structname}} *{{Structname}}) UnstageVoid(stage *StageStruct) {
	delete(stage.{{Structname}}s, {{structname}})
	delete(stage.{{Structname}}s_mapString, {{structname}}.Name)
}

// commit {{structname}} to the back repo (if it is already staged)
func ({{structname}} *{{Structname}}) Commit(stage *StageStruct) *{{Structname}} {
	if _, ok := stage.{{Structname}}s[{{structname}}]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.Commit{{Structname}}({{structname}})
		}
	}
	return {{structname}}
}

func ({{structname}} *{{Structname}}) CommitVoid(stage *StageStruct) {
	{{structname}}.Commit(stage)
}

// Checkout {{structname}} to the back repo (if it is already staged)
func ({{structname}} *{{Structname}}) Checkout(stage *StageStruct) *{{Structname}} {
	if _, ok := stage.{{Structname}}s[{{structname}}]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.Checkout{{Structname}}({{structname}})
		}
	}
	return {{structname}}
}

// for satisfaction of GongStruct interface
func ({{structname}} *{{Structname}}) GetName() (res string) {
	return {{structname}}.Name
}
`,

	ModelGongStructInsertionCreateCallback: `
	CreateORM{{Structname}}({{Structname}} *{{Structname}})`,

	ModelGongStructInsertionDeleteCallback: `
	DeleteORM{{Structname}}({{Structname}} *{{Structname}})`,

	ModelGongStructInsertionArrayDefintion: `
	{{Structname}}s           map[*{{Structname}}]any
	{{Structname}}s_mapString map[string]*{{Structname}}

	OnAfter{{Structname}}CreateCallback OnAfterCreateInterface[{{Structname}}]
	OnAfter{{Structname}}UpdateCallback OnAfterUpdateInterface[{{Structname}}]
	OnAfter{{Structname}}DeleteCallback OnAfterDeleteInterface[{{Structname}}]
	OnAfter{{Structname}}ReadCallback   OnAfterReadInterface[{{Structname}}]
`,

	ModelGongStructInsertionArrayInitialisation: `
		{{Structname}}s:           make(map[*{{Structname}}]any),
		{{Structname}}s_mapString: make(map[string]*{{Structname}}),
`,

	ModelGongStructInsertionArrayReset: `
	stage.{{Structname}}s = make(map[*{{Structname}}]any)
	stage.{{Structname}}s_mapString = make(map[string]*{{Structname}})
`,

	ModelGongStructInsertionArrayNil: `
	stage.{{Structname}}s = nil
	stage.{{Structname}}s_mapString = nil
`,
	ModelGongStructInsertionArrayUnstage: `
	for {{structname}} := range stage.{{Structname}}s {
		{{structname}}.Unstage(stage)
	}
`,
	ModelGongStructInsertionUnmarshallDeclarations: `
	map_{{Structname}}_Identifiers := make(map[*{{Structname}}]string)
	_ = map_{{Structname}}_Identifiers

	{{structname}}Ordered := []*{{Structname}}{}
	for {{structname}} := range stage.{{Structname}}s {
		{{structname}}Ordered = append({{structname}}Ordered, {{structname}})
	}
	sort.Slice({{structname}}Ordered[:], func(i, j int) bool {
		return {{structname}}Ordered[i].Name < {{structname}}Ordered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of {{Structname}}"
	for idx, {{structname}} := range {{structname}}Ordered {

		id = generatesIdentifier("{{Structname}}", idx, {{structname}}.Name)
		map_{{Structname}}_Identifiers[{{structname}}] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "{{Structname}}")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", {{structname}}.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// {{Structname}} values setup"
		// Initialisation of values{{ValuesInitialization}}
	}
`,

	ModelGongStructInsertionUnmarshallPointersInitializations: `
	for idx, {{structname}} := range {{structname}}Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("{{Structname}}", idx, {{structname}}.Name)
		map_{{Structname}}_Identifiers[{{structname}}] = id

		// Initialisation of values{{PointersInitialization}}
	}
`,

	ModelGongStructInsertionComputeNbInstances: `
	stage.Map_GongStructName_InstancesNb["{{Structname}}"] = len(stage.{{Structname}}s)`,

	ModelGongStructInsertionGenericReversePointerAssociationsMaps: `
	// reverse maps of direct associations of {{Structname}}
	case {{Structname}}:
		switch fieldname {
		// insertion point for per direct association field{{fieldReversePointerAssociationMapCode}}
		}`,

	ModelGongStructInsertionGenericReverseSliceOfPointersAssociationsMaps: `
	// reverse maps of direct associations of {{Structname}}
	case {{Structname}}:
		switch fieldname {
		// insertion point for per direct association field{{fieldReverseSliceOfPointersAssociationMapCode}}
		}`,

	ModelGongStructInsertionGenericGongstructTypes: ` | {{Structname}}`,

	ModelGongStructInsertionGenericPointerToGongstructTypes: ` | *{{Structname}}`,

	ModelGongStructInsertionGenericGongSetTypes: `
		map[*{{Structname}}]any |`,

	ModelGongStructInsertionGenericGongstructName: `
	case {{Structname}}:
		res = "{{Structname}}"`,

	ModelGongStructInsertionGenericPointerToGongstructName: `
	case *{{Structname}}:
		res = "{{Structname}}"`,

	ModelGongStructInsertionGenericGongMapTypes: `
		map[string]*{{Structname}} |`,

	ModelGongStructInsertionGenericGetSetFunctions: `
	case map[*{{Structname}}]any:
		return any(&stage.{{Structname}}s).(*Type)`,

	ModelGongStructInsertionGenericGetMapFunctions: `
	case map[string]*{{Structname}}:
		return any(&stage.{{Structname}}s_mapString).(*Type)`,

	ModelGongStructInsertionGenericInstancesSetFunctions: `
	case {{Structname}}:
		return any(&stage.{{Structname}}s).(*map[*Type]any)`,

	ModelGongStructInsertionGenericInstancesSetFromPointerTypeFunctions: `
	case *{{Structname}}:
		return any(&stage.{{Structname}}s).(*map[Type]any)`,

	ModelGongStructInsertionGenericInstancesMapFunctions: `
	case {{Structname}}:
		return any(&stage.{{Structname}}s_mapString).(*map[string]*Type)`,

	ModelGongStructInsertionGenericGetAssociationNameFunctions: `
	case {{Structname}}:
		return any(&{{Structname}}{
			// Initialisation of associations{{associationFieldInitialization}}
		}).(*Type)`,
}

// Sub sub Templates identifiers per gong field
//
// For each gongstruct, a code snippet will be generated from each sub template
type GongFilePerStructSubTemplateId int

const (
	GongFileFieldSubTmplStringFieldName GongFilePerStructSubTemplateId = iota
	GongFileFieldSubTmplReverseField

	GongFileFieldSubTmplStringValueBasicFieldBool
	GongFileFieldSubTmplStringValueBasicFieldInt
	GongFileFieldSubTmplStringValueBasicFieldEnumString
	GongFileFieldSubTmplStringValueBasicFieldEnumInt
	GongFileFieldSubTmplStringValueBasicFieldFloat64
	GongFileFieldSubTmplStringValueBasicFieldString
	GongFileFieldSubTmplStringValueTimeField
	GongFileFieldSubTmplStringValuePointerField
	GongFileFieldSubTmplStringValueSliceOfPointersField

	GongFileFieldSubTmplAssociationNamePointerField
	GongFileFieldSubTmplAssociationNameSliceOfPointersField
	GongFileFieldSubTmplAssociationNameEnclosingCompositePointerField
	GongFileFieldSubTmplAssociationNameCompositePointerField

	GongFileFieldSubTmplPointerFieldPointerAssociationMapFunction
	GongFileFieldSubTmplPointerFieldSliceOfPointersAssociationMapFunction
)

// for each sub template code, there is the sub template code
var GongFileFieldFieldSubTemplateCode map[GongFilePerStructSubTemplateId]string = // declaration of the sub templates
map[GongFilePerStructSubTemplateId]string{

	GongFileFieldSubTmplStringFieldName: `"{{FieldName}}"`,

	GongFileFieldSubTmplReverseField: `
		rf.GongstructName = "{{AssocStructName}}"
		rf.Fieldname = "{{FieldName}}"
		res = append(res, rf)`,

	GongFileFieldSubTmplStringValueBasicFieldBool: `
		case "{{FieldName}}":
			res = fmt.Sprintf("%t", inferedInstance.{{FieldName}})`,
	GongFileFieldSubTmplStringValueBasicFieldInt: `
		case "{{FieldName}}":
			res = fmt.Sprintf("%d", inferedInstance.{{FieldName}})`,
	GongFileFieldSubTmplStringValueBasicFieldEnumString: `
		case "{{FieldName}}":
			enum := inferedInstance.{{FieldName}}
			res = enum.ToCodeString()`,
	GongFileFieldSubTmplStringValueBasicFieldEnumInt: `
		case "{{FieldName}}":
			enum := inferedInstance.{{FieldName}}
			res = enum.ToCodeString()`,
	GongFileFieldSubTmplStringValueBasicFieldFloat64: `
		case "{{FieldName}}":
			res = fmt.Sprintf("%f", inferedInstance.{{FieldName}})`,
	GongFileFieldSubTmplStringValueBasicFieldString: `
		case "{{FieldName}}":
			res = inferedInstance.{{FieldName}}`,
	GongFileFieldSubTmplStringValueTimeField: `
		case "{{FieldName}}":
			res = inferedInstance.{{FieldName}}.String()`,
	GongFileFieldSubTmplStringValuePointerField: `
		case "{{FieldName}}":
			if inferedInstance.{{FieldName}} != nil {
				res = inferedInstance.{{FieldName}}.Name
			}`,
	GongFileFieldSubTmplStringValueSliceOfPointersField: `
		case "{{FieldName}}":
			for idx, __instance__ := range inferedInstance.{{FieldName}} {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}`,

	GongFileFieldSubTmplAssociationNamePointerField: `
			// field is initialized with an instance of {{AssocStructName}} with the name of the field
			{{FieldName}}: &{{AssocStructName}}{Name: "{{FieldName}}"},`,

	GongFileFieldSubTmplAssociationNameSliceOfPointersField: `
			// field is initialized with an instance of {{AssocStructName}} with the name of the field
			{{FieldName}}: []*{{AssocStructName}}{{Name: "{{FieldName}}"}},`,

	GongFileFieldSubTmplAssociationNameEnclosingCompositePointerField: `
			// field is initialized with {{AssocCompositeStructName}} as it is a composite
			{{AssocCompositeStructName}}: {{AssocCompositeStructName}}{
				// per field init{{PerCompositeFieldInit}}
			},`,
	GongFileFieldSubTmplAssociationNameCompositePointerField: `
				//
				{{FieldName}}: &{{AssocStructName}}{Name: "{{FieldName}}"},`,

	GongFileFieldSubTmplPointerFieldPointerAssociationMapFunction: `
		case "{{FieldName}}":
			res := make(map[*{{AssocStructName}}][]*{{Structname}})
			for {{structname}} := range stage.{{Structname}}s {
				if {{structname}}.{{FieldName}} != nil {
					{{assocstructname}}_ := {{structname}}.{{FieldName}}
					var {{structname}}s []*{{Structname}}
					_, ok := res[{{assocstructname}}_]
					if ok {
						{{structname}}s = res[{{assocstructname}}_]
					} else {
						{{structname}}s = make([]*{{Structname}}, 0)
					}
					{{structname}}s = append({{structname}}s, {{structname}})
					res[{{assocstructname}}_] = {{structname}}s
				}
			}
			return any(res).(map[*End][]*Start)`,

	GongFileFieldSubTmplPointerFieldSliceOfPointersAssociationMapFunction: `
		case "{{FieldName}}":
			res := make(map[*{{AssocStructName}}]*{{Structname}})
			for {{structname}} := range stage.{{Structname}}s {
				for _, {{assocstructname}}_ := range {{structname}}.{{FieldName}} {
					res[{{assocstructname}}_] = {{structname}}
				}
			}
			return any(res).(map[*End]*Start)`,
}

func CodeGeneratorModelGong(
	mdlPkg *models.ModelPkg,
	pkgName string,
	pkgPath string) {

	// generate the typescript file
	codeGO := ModelGongFileTemplate

	subStructCodes := make(map[ModelGongStructInsertionId]string)
	for subStructTemplate := range ModelGongStructSubTemplateCode {
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

		for subStructTemplate := range ModelGongStructSubTemplateCode {

			// replace {{ValuesInitialization}}
			fieldNames := `
		res = []string{`
			reverseFields := `
		var rf ReverseField
		_ = rf`
			fieldStringValues := ``
			fieldReversePointerAssociationMapCode := ``
			fieldReverseSliceOfPointersAssociationMapCode := ``
			associationFieldInitialization := ``

			associationFieldInitializationPerCompositeStruct := make(map[string]string, 0)

			for idx, field := range gongStruct.Fields {

				switch field := field.(type) {
				case *models.GongBasicField:

					switch field.GetBasicKind() {
					case types.String:
						if field.GongEnum == nil {
							fieldStringValues += models.Replace1(
								GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringValueBasicFieldString],
								"{{FieldName}}", field.Name)
						} else {
							fieldStringValues += models.Replace1(
								GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringValueBasicFieldEnumString],
								"{{FieldName}}", field.Name)
						}
					case types.Bool:
						fieldStringValues += models.Replace1(
							GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringValueBasicFieldBool],
							"{{FieldName}}", field.Name)
					case types.Float64:
						fieldStringValues += models.Replace1(
							GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringValueBasicFieldFloat64],
							"{{FieldName}}", field.Name)
					case types.Int, types.Int64:
						if field.GongEnum == nil {
							fieldStringValues += models.Replace1(
								GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringValueBasicFieldInt],
								"{{FieldName}}", field.Name)
						} else {
							fieldStringValues += models.Replace1(
								GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringValueBasicFieldEnumInt],
								"{{FieldName}}", field.Name)
						}
					default:
					}
				case *models.GongTimeField:
					fieldStringValues += models.Replace1(
						GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringValueTimeField],
						"{{FieldName}}", field.Name)
				case *models.PointerToGongStructField:
					fieldStringValues += models.Replace1(
						GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringValuePointerField],
						"{{FieldName}}", field.Name)
					fieldReversePointerAssociationMapCode += models.Replace3(
						GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplPointerFieldPointerAssociationMapFunction],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name,
						"{{assocstructname}}", strings.ToLower(field.GongStruct.Name))
					if field.CompositeStructName == "" {
						associationFieldInitialization += models.Replace3(
							GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplAssociationNamePointerField],
							"{{FieldName}}", field.Name,
							"{{AssocStructName}}", field.GongStruct.Name,
							"{{assocstructname}}", strings.ToLower(field.GongStruct.Name))
					} else {
						associationFieldInitializationPerCompositeStruct[field.CompositeStructName] += models.Replace4(
							GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplAssociationNameCompositePointerField],
							"{{FieldName}}", field.Name,
							"{{AssocStructName}}", field.GongStruct.Name,
							"{{AssocCompositeStructName}}", field.CompositeStructName,
							"{{assocstructname}}", strings.ToLower(field.GongStruct.Name))
					}
				case *models.SliceOfPointerToGongStructField:
					fieldStringValues += models.Replace1(
						GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringValueSliceOfPointersField],
						"{{FieldName}}", field.Name)
					fieldReverseSliceOfPointersAssociationMapCode += models.Replace3(
						GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplPointerFieldSliceOfPointersAssociationMapFunction],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name,
						"{{assocstructname}}", strings.ToLower(field.GongStruct.Name))
					associationFieldInitialization += models.Replace3(
						GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplAssociationNameSliceOfPointersField],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name,
						"{{assocstructname}}", strings.ToLower(field.GongStruct.Name))
				default:
				}

				if idx > 0 {
					fieldNames += ", "
				}
				fieldNames += models.Replace1(
					GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringFieldName],
					"{{FieldName}}", field.GetName())
			}

			//
			// Parse all fields from other structs that points to this struct
			//
			for _, __struct := range gongStructs {
				for _, field := range __struct.Fields {
					switch field := field.(type) {
					case *models.SliceOfPointerToGongStructField:

						if field.GongStruct == gongStruct {

							reverseFields += models.Replace2(
								GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplReverseField],
								"{{AssocStructName}}", __struct.Name,
								"{{FieldName}}", field.GetName())
						}
					}
				}
			}

			fieldStringValues = models.Replace2(fieldStringValues,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			fieldReversePointerAssociationMapCode = models.Replace2(fieldReversePointerAssociationMapCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			fieldReverseSliceOfPointersAssociationMapCode = models.Replace2(fieldReverseSliceOfPointersAssociationMapCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			fieldNames += `}`

			// The generation has to be be reproductible, therefore the map
			// associationFieldInitializationPerCompositeStruct has to be ordered
			keys := make([]string, 0)
			for k := range associationFieldInitializationPerCompositeStruct {
				keys = append(keys, k)
			}

			// sort the keys in alphabetical order
			sort.Strings(keys)

			for _, compositeStructName := range keys {
				associationFieldInitialization += models.Replace2(
					GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplAssociationNameEnclosingCompositePointerField],
					"{{AssocCompositeStructName}}", compositeStructName,
					"{{PerCompositeFieldInit}}", associationFieldInitializationPerCompositeStruct[compositeStructName])
			}

			generatedCodeFromSubTemplate := models.Replace8(ModelGongStructSubTemplateCode[subStructTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{ListOfFieldsName}}", fieldNames,
				"{{ListOfReverseFields}}", reverseFields,
				"{{StringValueOfFields}}", fieldStringValues,
				"{{fieldReversePointerAssociationMapCode}}", fieldReversePointerAssociationMapCode,
				"{{fieldReverseSliceOfPointersAssociationMapCode}}", fieldReverseSliceOfPointersAssociationMapCode,
				"{{associationFieldInitialization}}", associationFieldInitialization,
			)

			subStructCodes[subStructTemplate] += generatedCodeFromSubTemplate
		}
	}

	// substitutes {{<<insertionPerStructId points>>}} stuff with generated code
	for insertionPerStructId := ModelGongStructInsertionId(0); insertionPerStructId < ModelGongStructInsertionsNb; insertionPerStructId++ {
		toReplace := "{{" + string(rune(insertionPerStructId)) + "}}"
		codeGO = strings.ReplaceAll(codeGO, toReplace, subStructCodes[insertionPerStructId])
	}

	caserEnglish := cases.Title(language.English)
	codeGO = models.Replace4(codeGO,
		"{{PkgName}}", pkgName,
		"{{TitlePkgName}}", caserEnglish.String(pkgName),
		"{{pkgname}}", strings.ToLower(pkgName),
		"	 | ", "	", // for the replacement of the of the first bar in the Gongstruct Type def
	)

	file, err := os.Create(filepath.Join(pkgPath, "gong.go"))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeGO)

}
