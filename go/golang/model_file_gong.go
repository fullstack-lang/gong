package golang

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
	ModelGongStructInsertionUnmarshallDeclarations
	ModelGongStructInsertionUnmarshallPointersInitializations
	ModelGongStructInsertionComputeNbInstances
	ModelGongStructInsertionReverseAssociationsMaps
	ModelGongStructInsertionGenericGetFields
	ModelGongStructInsertionGenericGetFieldValues
	ModelGongStructInsertionGenericReversePointerAssociationsMaps
	ModelGongStructInsertionGenericReverseSliceOfPointersAssociationsMaps
	ModelGongStructInsertionGenericGongstructTypes
	ModelGongStructInsertionGenericPointerToGongstructTypes
	ModelGongStructInsertionGenericGongSetTypes
	ModelGongStructInsertionGenericGongstructName
	ModelGongStructInsertionGenericGongMapTypes
	ModelGongStructInsertionGenericGetSetFunctions
	ModelGongStructInsertionGenericGetMapFunctions
	ModelGongStructInsertionGenericInstancesSetFunctions
	ModelGongStructInsertionGenericInstancesMapFunctions
	ModelGongStructInsertionGenericGetAssociationNameFunctions
	ModelGongStructInsertionsNb
)

// insertion code for all enums
type ModelGongEnumInsertionId int

const (
	// iota + 40 is to separate the insertion code of gongstruct from insertion code of gongenum
	ModelGongEnumUtilityFunctions ModelGongEnumInsertionId = iota + 40
	ModelGongEnumInsertionsNb
)

var ModelGongEnumSubTemplateCode map[ModelGongEnumInsertionId]string = // new line
map[ModelGongEnumInsertionId]string{
	ModelGongEnumUtilityFunctions: `
// Utility function for {{EnumName}}
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func ({{enumName}} {{EnumName}}) To{{Type}}() (res {{type}}) {

	// migration of former implementation of enum
	switch {{enumName}} {
	// insertion code per enum code{{ToStringPerCodeCode}}
	}
	return
}

func ({{enumName}} *{{EnumName}}) From{{Type}}(input {{type}}) (err error) {

	switch input {
	// insertion code per enum code{{FromStringPerCodeCode}}
	default:
		return errUnkownEnum
	}
	return
}

func ({{enumName}} *{{EnumName}}) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code{{FromCodeStringPerCodeCode}}
	default:
		return errUnkownEnum
	}
	return
}

func ({{enumName}} *{{EnumName}}) ToCodeString() (res string) {

	switch *{{enumName}} {
	// insertion code per enum code{{ToCodeStringPerCodeCode}}
	}
	return
}
`,
}

var ModelGongStructSubTemplateCode map[ModelGongStructInsertionId]string = // new line
map[ModelGongStructInsertionId]string{
	ModelGongStructInsertionCommitCheckout: `
	Commit{{Structname}}({{structname}} *{{Structname}})
	Checkout{{Structname}}({{structname}} *{{Structname}})`,

	ModelGongStructInsertionGenericGetFields: `
	case {{Structname}}:{{ListOfFieldsName}}`,

	ModelGongStructInsertionGenericGetFieldValues: `
	case {{Structname}}:
		switch fieldName {
		// string value of fields{{StringValueOfFields}}
		}`,

	ModelGongStructInsertionStageFunctions: `
// Stage puts {{structname}} to the model stage
func ({{structname}} *{{Structname}}) Stage() *{{Structname}} {
	Stage.{{Structname}}s[{{structname}}] = __member
	Stage.{{Structname}}s_mapString[{{structname}}.Name] = {{structname}}

	return {{structname}}
}

// Unstage removes {{structname}} off the model stage
func ({{structname}} *{{Structname}}) Unstage() *{{Structname}} {
	delete(Stage.{{Structname}}s, {{structname}})
	delete(Stage.{{Structname}}s_mapString, {{structname}}.Name)
	return {{structname}}
}

// commit {{structname}} to the back repo (if it is already staged)
func ({{structname}} *{{Structname}}) Commit() *{{Structname}} {
	if _, ok := Stage.{{Structname}}s[{{structname}}]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.Commit{{Structname}}({{structname}})
		}
	}
	return {{structname}}
}

// Checkout {{structname}} to the back repo (if it is already staged)
func ({{structname}} *{{Structname}}) Checkout() *{{Structname}} {
	if _, ok := Stage.{{Structname}}s[{{structname}}]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.Checkout{{Structname}}({{structname}})
		}
	}
	return {{structname}}
}

//
// Legacy, to be deleted
//

// StageCopy appends a copy of {{structname}} to the model stage
func ({{structname}} *{{Structname}}) StageCopy() *{{Structname}} {
	_{{structname}} := new({{Structname}})
	*_{{structname}} = *{{structname}}
	_{{structname}}.Stage()
	return _{{structname}}
}

// StageAndCommit appends {{structname}} to the model stage and commit to the orm repo
func ({{structname}} *{{Structname}}) StageAndCommit() *{{Structname}} {
	{{structname}}.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORM{{Structname}}({{structname}})
	}
	return {{structname}}
}

// DeleteStageAndCommit appends {{structname}} to the model stage and commit to the orm repo
func ({{structname}} *{{Structname}}) DeleteStageAndCommit() *{{Structname}} {
	{{structname}}.Unstage()
	DeleteORM{{Structname}}({{structname}})
	return {{structname}}
}

// StageCopyAndCommit appends a copy of {{structname}} to the model stage and commit to the orm repo
func ({{structname}} *{{Structname}}) StageCopyAndCommit() *{{Structname}} {
	_{{structname}} := new({{Structname}})
	*_{{structname}} = *{{structname}}
	_{{structname}}.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORM{{Structname}}({{structname}})
	}
	return _{{structname}}
}

// CreateORM{{Structname}} enables dynamic staging of a {{Structname}} instance
func CreateORM{{Structname}}({{structname}} *{{Structname}}) {
	{{structname}}.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORM{{Structname}}({{structname}})
	}
}

// DeleteORM{{Structname}} enables dynamic staging of a {{Structname}} instance
func DeleteORM{{Structname}}({{structname}} *{{Structname}}) {
	{{structname}}.Unstage()
	if Stage.AllModelsStructDeleteCallback != nil {
		Stage.AllModelsStructDeleteCallback.DeleteORM{{Structname}}({{structname}})
	}
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

	ModelGongStructInsertionReverseAssociationsMaps: `

// generate function for reverse association maps of {{Structname}}{{ReverseAssociationMapFunctions}}`,

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

	ModelGongStructInsertionGenericGongMapTypes: `
		map[string]*{{Structname}} |`,

	ModelGongStructInsertionGenericGetSetFunctions: `
	case map[*{{Structname}}]any:
		return any(&Stage.{{Structname}}s).(*Type)`,

	ModelGongStructInsertionGenericGetMapFunctions: `
	case map[string]*{{Structname}}:
		return any(&Stage.{{Structname}}s_mapString).(*Type)`,

	ModelGongStructInsertionGenericInstancesSetFunctions: `
	case {{Structname}}:
		return any(&Stage.{{Structname}}s).(*map[*Type]any)`,

	ModelGongStructInsertionGenericInstancesMapFunctions: `
	case {{Structname}}:
		return any(&Stage.{{Structname}}s_mapString).(*map[string]*Type)`,

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
	GongFileFieldSubTmplSetBasicFieldBool GongFilePerStructSubTemplateId = iota
	GongFileFieldSubTmplSetBasicFieldInt
	GongFileFieldSubTmplSetBasicFieldEnumString
	GongFileFieldSubTmplSetBasicFieldEnumInt
	GongFileFieldSubTmplSetBasicFieldFloat64
	GongFileFieldSubTmplSetBasicFieldString
	GongFileFieldSubTmplSetTimeField
	GongFileFieldSubTmplSetPointerField
	GongFileFieldSubTmplSetSliceOfPointersField

	GongFileFieldSubTmplStringFieldName

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
	GongFileFieldSubTmplAssociationNameCompositePointerField
	GongFileFieldSubTmplAssociationNameCompositeSliceOfPointersField

	GongFileFieldSubTmplPointerFieldAssociationMapFunction
	GongFileFieldSubTmplSliceOfPointersFieldAssociationMapFunction

	GongFileFieldSubTmplPointerFieldPointerAssociationMapFunction
	GongFileFieldSubTmplPointerFieldSliceOfPointersAssociationMapFunction
)

// for each sub template code, there is the sub template code
var GongFileFieldFieldSubTemplateCode map[GongFilePerStructSubTemplateId]string = // declaration of the sub templates
map[GongFilePerStructSubTemplateId]string{

	GongFileFieldSubTmplSetBasicFieldBool: `
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "{{FieldName}}")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", {{structname}}.{{FieldName}}))
		initializerStatements += setValueField
`,
	GongFileFieldSubTmplSetTimeField: `
		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "{{FieldName}}")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", {{structname}}.{{FieldName}}.String())
		initializerStatements += setValueField
`,
	GongFileFieldSubTmplSetBasicFieldInt: `
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "{{FieldName}}")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", {{structname}}.{{FieldName}}))
		initializerStatements += setValueField
`,
	GongFileFieldSubTmplSetBasicFieldEnumString: `
		if {{structname}}.{{FieldName}} != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "{{FieldName}}")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+{{structname}}.{{FieldName}}.ToCodeString())
			initializerStatements += setValueField
		}
`,
	GongFileFieldSubTmplSetBasicFieldEnumInt: `
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "{{FieldName}}")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+{{structname}}.{{FieldName}}.ToCodeString())
		initializerStatements += setValueField
`,
	GongFileFieldSubTmplSetBasicFieldFloat64: `
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "{{FieldName}}")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", {{structname}}.{{FieldName}}))
		initializerStatements += setValueField
`,
	GongFileFieldSubTmplSetBasicFieldString: `
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "{{FieldName}}")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string({{structname}}.{{FieldName}}))
		initializerStatements += setValueField
`,
	GongFileFieldSubTmplSetPointerField: `
		if {{structname}}.{{FieldName}} != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "{{FieldName}}")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_{{AssocStructName}}_Identifiers[{{structname}}.{{FieldName}}])
			pointersInitializesStatements += setPointerField
		}
`,
	GongFileFieldSubTmplSetSliceOfPointersField: `
		for _, _{{assocstructname}} := range {{structname}}.{{FieldName}} {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "{{FieldName}}")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_{{AssocStructName}}_Identifiers[_{{assocstructname}}])
			pointersInitializesStatements += setPointerField
		}
`,
	GongFileFieldSubTmplStringFieldName: `"{{FieldName}}"`,

	GongFileFieldSubTmplStringValueBasicFieldBool: `
		case "{{FieldName}}":
			res = fmt.Sprintf("%t", any(instance).({{Structname}}).{{FieldName}})`,
	GongFileFieldSubTmplStringValueBasicFieldInt: `
		case "{{FieldName}}":
			res = fmt.Sprintf("%d", any(instance).({{Structname}}).{{FieldName}})`,
	GongFileFieldSubTmplStringValueBasicFieldEnumString: `
		case "{{FieldName}}":
			enum := any(instance).({{Structname}}).{{FieldName}}
			res = enum.ToCodeString()`,
	GongFileFieldSubTmplStringValueBasicFieldEnumInt: `
		case "{{FieldName}}":
			enum := any(instance).({{Structname}}).{{FieldName}}
			res = enum.ToCodeString()`,
	GongFileFieldSubTmplStringValueBasicFieldFloat64: `
		case "{{FieldName}}":
			res = fmt.Sprintf("%f", any(instance).({{Structname}}).{{FieldName}})`,
	GongFileFieldSubTmplStringValueBasicFieldString: `
		case "{{FieldName}}":
			res = any(instance).({{Structname}}).{{FieldName}}`,
	GongFileFieldSubTmplStringValueTimeField: `
		case "{{FieldName}}":
			res = any(instance).({{Structname}}).{{FieldName}}.String()`,
	GongFileFieldSubTmplStringValuePointerField: `
		case "{{FieldName}}":
			if any(instance).({{Structname}}).{{FieldName}} != nil {
				res = any(instance).({{Structname}}).{{FieldName}}.Name
			}`,
	GongFileFieldSubTmplStringValueSliceOfPointersField: `
		case "{{FieldName}}":
			for idx, __instance__ := range any(instance).({{Structname}}).{{FieldName}} {
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

	GongFileFieldSubTmplAssociationNameCompositePointerField: `
			// field is initialized with an instance of {{AssocStructName}} ({{AssocCompositeStructName}} as it is a composite) with the name of the field
			{{AssocCompositeStructName}}: {{AssocCompositeStructName}}{
				{{FieldName}}: &{{AssocStructName}}{Name: "{{FieldName}}"},
			},`,

	GongFileFieldSubTmplAssociationNameCompositeSliceOfPointersField: `
			// field is initialized with an instance of {{AssocStructName}} ({{AssocCompositeStructName}} as it is a composite) with the name of the field
			{{AssocCompositeStructName}}: {{AssocCompositeStructName}}{
				{{FieldName}}: []*{{AssocStructName}}{{Name: "{{FieldName}}"}},
			},`,

	GongFileFieldSubTmplPointerFieldAssociationMapFunction: `
func (stageStruct *StageStruct) CreateReverseMap_{{Structname}}_{{FieldName}}() (res map[*{{AssocStructName}}][]*{{Structname}}) {
	res = make(map[*{{AssocStructName}}][]*{{Structname}})

	for {{structname}} := range stageStruct.{{Structname}}s {
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

	return
}`,
	GongFileFieldSubTmplPointerFieldPointerAssociationMapFunction: `
		case "{{FieldName}}":
			res := make(map[*{{AssocStructName}}][]*{{Structname}})
			for {{structname}} := range Stage.{{Structname}}s {
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
			for {{structname}} := range Stage.{{Structname}}s {
				for _, {{assocstructname}}_ := range {{structname}}.{{FieldName}} {
					res[{{assocstructname}}_] = {{structname}}
				}
			}
			return any(res).(map[*End]*Start)`,

	GongFileFieldSubTmplSliceOfPointersFieldAssociationMapFunction: `
func (stageStruct *StageStruct) CreateReverseMap_{{Structname}}_{{FieldName}}() (res map[*{{AssocStructName}}]*{{Structname}}) {
	res = make(map[*{{AssocStructName}}]*{{Structname}})

	for {{structname}} := range stageStruct.{{Structname}}s {
		for _, {{assocstructname}}_ := range {{structname}}.{{FieldName}} {
			res[{{assocstructname}}_] = {{structname}}
		}
	}

	return
}
`,
}

// gongenum value template
type GongModelEnumValueSubTemplateId int

const (
	GongModelEnumValueFromString GongModelEnumValueSubTemplateId = iota
	GongModelEnumValueFromCodeString
	GongModelEnumValueToString
	GongModelEnumValueToCodeString
)

var GongModelEnumValueSubTemplateCode map[GongModelEnumValueSubTemplateId]string = // declaration of the sub templates
map[GongModelEnumValueSubTemplateId]string{

	GongModelEnumValueFromString: `
	case {{GongEnumValue}}:
		*{{enumName}} = {{GongEnumCode}}`,
	GongModelEnumValueFromCodeString: `
	case "{{GongEnumCode}}":
		*{{enumName}} = {{GongEnumCode}}`,
	GongModelEnumValueToString: `
	case {{GongEnumCode}}:
		res = {{GongEnumValue}}`,
	GongModelEnumValueToCodeString: `
	case {{GongEnumCode}}:
		res = "{{GongEnumCode}}"`,
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
			valInitCode := ""
			pointerInitCode := ""
			fieldNames := `
		res = []string{`
			fieldStringValues := ``
			fieldReverseAssociationMapCreationCode := ``
			fieldReversePointerAssociationMapCode := ``
			fieldReverseSliceOfPointersAssociationMapCode := ``
			associationFieldInitialization := ``

			for idx, field := range gongStruct.Fields {

				switch field := field.(type) {
				case *models.GongBasicField:

					switch field.GetBasicKind() {
					case types.String:
						if field.GongEnum == nil {
							valInitCode += models.Replace1(
								GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetBasicFieldString],
								"{{FieldName}}", field.Name)

							fieldStringValues += models.Replace1(
								GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringValueBasicFieldString],
								"{{FieldName}}", field.Name)
						} else {
							valInitCode += models.Replace1(
								GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetBasicFieldEnumString],
								"{{FieldName}}", field.Name)
							fieldStringValues += models.Replace1(
								GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringValueBasicFieldEnumString],
								"{{FieldName}}", field.Name)
						}
					case types.Bool:
						valInitCode += models.Replace1(
							GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetBasicFieldBool],
							"{{FieldName}}", field.Name)
						fieldStringValues += models.Replace1(
							GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringValueBasicFieldBool],
							"{{FieldName}}", field.Name)
					case types.Float64:
						valInitCode += models.Replace1(
							GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetBasicFieldFloat64],
							"{{FieldName}}", field.Name)
						fieldStringValues += models.Replace1(
							GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringValueBasicFieldFloat64],
							"{{FieldName}}", field.Name)
					case types.Int, types.Int64:
						if field.GongEnum == nil {
							valInitCode += models.Replace1(
								GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetBasicFieldInt],
								"{{FieldName}}", field.Name)
							fieldStringValues += models.Replace1(
								GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringValueBasicFieldInt],
								"{{FieldName}}", field.Name)
						} else {
							valInitCode += models.Replace1(
								GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetBasicFieldEnumInt],
								"{{FieldName}}", field.Name)
							fieldStringValues += models.Replace1(
								GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringValueBasicFieldEnumInt],
								"{{FieldName}}", field.Name)
						}
					default:
					}
				case *models.GongTimeField:
					valInitCode += models.Replace1(
						GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetTimeField],
						"{{FieldName}}", field.Name)
					fieldStringValues += models.Replace1(
						GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringValueTimeField],
						"{{FieldName}}", field.Name)
				case *models.PointerToGongStructField:
					pointerInitCode += models.Replace2(
						GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetPointerField],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name)
					fieldStringValues += models.Replace1(
						GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringValuePointerField],
						"{{FieldName}}", field.Name)
					fieldReverseAssociationMapCreationCode += models.Replace3(
						GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplPointerFieldAssociationMapFunction],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name,
						"{{assocstructname}}", strings.ToLower(field.GongStruct.Name))
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
						associationFieldInitialization += models.Replace4(
							GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplAssociationNameCompositePointerField],
							"{{FieldName}}", field.Name,
							"{{AssocStructName}}", field.GongStruct.Name,
							"{{AssocCompositeStructName}}", field.CompositeStructName,
							"{{assocstructname}}", strings.ToLower(field.GongStruct.Name))
					}
				case *models.SliceOfPointerToGongStructField:
					pointerInitCode += models.Replace3(
						GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetSliceOfPointersField],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name,
						"{{assocstructname}}", strings.ToLower(field.GongStruct.Name))
					fieldStringValues += models.Replace1(
						GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringValueSliceOfPointersField],
						"{{FieldName}}", field.Name)
					fieldReverseAssociationMapCreationCode += models.Replace3(
						GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSliceOfPointersFieldAssociationMapFunction],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name,
						"{{assocstructname}}", strings.ToLower(field.GongStruct.Name))
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

			valInitCode = models.Replace2(valInitCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			pointerInitCode = models.Replace2(pointerInitCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			fieldStringValues = models.Replace2(fieldStringValues,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			fieldReverseAssociationMapCreationCode = models.Replace2(fieldReverseAssociationMapCreationCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			fieldReversePointerAssociationMapCode = models.Replace2(fieldReversePointerAssociationMapCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			fieldReverseSliceOfPointersAssociationMapCode = models.Replace2(fieldReverseSliceOfPointersAssociationMapCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			fieldNames += `}`
			generatedCodeFromSubTemplate := models.Replace10(ModelGongStructSubTemplateCode[subStructTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{ValuesInitialization}}", valInitCode,
				"{{PointersInitialization}}", pointerInitCode,
				"{{ListOfFieldsName}}", fieldNames,
				"{{StringValueOfFields}}", fieldStringValues,
				"{{ReverseAssociationMapFunctions}}", fieldReverseAssociationMapCreationCode,
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

	subEnumCodes := make(map[ModelGongEnumInsertionId]string)
	for subEnumTemplate := range ModelGongEnumSubTemplateCode {
		subEnumCodes[subEnumTemplate] = ""
	}

	// sort gong enums per name (for reproductibility)
	gongEnums := []*models.GongEnum{}
	for _, _enum := range mdlPkg.GongEnums {
		gongEnums = append(gongEnums, _enum)
	}
	sort.Slice(gongEnums[:], func(i, j int) bool {
		return gongEnums[i].Name < gongEnums[j].Name
	})

	for _, gongEnum := range gongEnums {

		for subEnumTemplate := range ModelGongEnumSubTemplateCode {

			codeFromStringPerGongValue := ""
			codeFromCodeStringPerGongValue := ""
			codeToStringPerGongValue := ""
			codeToCodeStringPerGongValue := ""

			for _, enumValue := range gongEnum.GongEnumValues {
				codeFromStringPerGongValue += models.Replace2(GongModelEnumValueSubTemplateCode[GongModelEnumValueFromString],
					"{{GongEnumValue}}", enumValue.Value,
					"{{GongEnumCode}}", enumValue.Name)
				codeFromCodeStringPerGongValue += models.Replace1(GongModelEnumValueSubTemplateCode[GongModelEnumValueFromCodeString],
					"{{GongEnumCode}}", enumValue.Name)
				codeToStringPerGongValue += models.Replace2(GongModelEnumValueSubTemplateCode[GongModelEnumValueToString],
					"{{GongEnumValue}}", enumValue.Value,
					"{{GongEnumCode}}", enumValue.Name)

				codeToCodeStringPerGongValue += models.Replace2(GongModelEnumValueSubTemplateCode[GongModelEnumValueToCodeString],
					"{{GongEnumValue}}", enumValue.Value,
					"{{GongEnumCode}}", enumValue.Name)
			}

			generatedCodeFromSubTemplate := models.Replace4(ModelGongEnumSubTemplateCode[subEnumTemplate],
				"{{ToStringPerCodeCode}}", codeToStringPerGongValue,
				"{{FromStringPerCodeCode}}", codeFromStringPerGongValue,
				"{{FromCodeStringPerCodeCode}}", codeFromCodeStringPerGongValue,
				"{{ToCodeStringPerCodeCode}}", codeToCodeStringPerGongValue)

			var typeOfEnumAsString string
			if gongEnum.Type == models.String {
				typeOfEnumAsString = "String"
			} else {
				typeOfEnumAsString = "Int"
			}

			generatedCodeFromSubTemplate = models.Replace4(generatedCodeFromSubTemplate,
				"{{enumName}}", strings.ToLower(gongEnum.Name),
				"{{EnumName}}", gongEnum.Name,
				"{{Type}}", typeOfEnumAsString,
				"{{type}}", strings.ToLower(typeOfEnumAsString))

			subEnumCodes[subEnumTemplate] += generatedCodeFromSubTemplate
		}
	}

	// substitutes {{<<insertionPerEnumId points>>}} stuff with generated code
	for insertionPerEnumId := ModelGongEnumUtilityFunctions; insertionPerEnumId < ModelGongEnumInsertionsNb; insertionPerEnumId++ {
		toReplace := "{{" + string(rune(insertionPerEnumId)) + "}}"
		codeGO = strings.ReplaceAll(codeGO, toReplace, subEnumCodes[insertionPerEnumId])
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
