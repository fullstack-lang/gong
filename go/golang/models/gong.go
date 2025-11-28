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

	ModelGongStructInsertionGenericGetReverseFields
	ModelGongStructInsertionGenericGetFieldsFromPointer
	ModelGongStructInsertionGenericGetFieldsHeadersMethod
	ModelGongStructInsertionGenericGetFieldValues
	ModelGongStructInsertionGenericGetFieldValuesFromPointer

	ModelGongStructInsertionGenericReversePointerAssociationsMaps
	ModelGongStructInsertionGenericReverseSliceOfPointersAssociationsMaps

	ModelGongStructInsertionGenericPointerToGongstructName
	ModelGongStructInsertionGenericGetSetFunctions
	ModelGongStructInsertionGenericGetMapFunctions
	ModelGongStructInsertionGenericInstancesSetFunctions
	ModelGongStructInsertionGenericInstancesSetFromPointerTypeFunctions
	ModelGongStructInsertionGenericInstancesSetFromPointerTypeFunctionsReturnType

	ModelGongStructInsertionGenericInstancesMapFunctions
	ModelGongStructInsertionGenericGetAssociationNameFunctions

	ModelGongStructInsertionGenericSetFieldValuesFromPointer
	ModelGongStructInsertionGenericGetGongstructName

	ModelGongOrderFields
	ModelGongOrderMapsInit
	ModelGongOrderSwitchGet

	ModelGongNamedStructsSliceInit
	ModelGongNamedStructsInstancesNames

	ModelGongNamedStructSortedOrderInstances

	ModelGongStructInsertionsNb
)

var ModelGongStructSubTemplateCode map[ModelGongStructInsertionId]string = // new line
map[ModelGongStructInsertionId]string{
	ModelGongStructInsertionCommitCheckout: `
	Commit{{Structname}}({{structname}} *{{Structname}})
	Checkout{{Structname}}({{structname}} *{{Structname}})`,

	ModelGongStructInsertionGenericGetReverseFields: `
	case *{{Structname}}:{{ListOfReverseFields}}`,

	ModelGongStructInsertionGenericGetFieldsHeadersMethod: `
func ({{structname}} *{{Structname}}) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers{{ListOfFieldHeaders}}
	return
}
`,

	ModelGongStructInsertionGenericGetFieldsFromPointer: `
	case *{{Structname}}:{{ListOfFieldHeaders}}`,

	ModelGongStructInsertionGenericGetFieldValues: `
	case {{Structname}}:
		switch fieldName {
		// string value of fields{{StringValueOfFields}}
		}`,

	ModelGongStructInsertionGenericGetFieldValuesFromPointer: `
func ({{structname}} *{{Structname}}) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields{{StringValueOfFields}}
	}
	return
}`,

	ModelGongStructInsertionGenericSetFieldValuesFromPointer: `
func ({{structname}} *{{Structname}}) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code{{SetFieldValues}}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}
`,

	ModelGongStructInsertionGenericGetGongstructName: `
func ({{structname}} *{{Structname}}) GongGetGongstructName() string {
	return "{{Structname}}"
}
`,

	ModelGongStructInsertionStageFunctions: `
// Stage puts {{structname}} to the model stage
func ({{structname}} *{{Structname}}) Stage(stage *Stage) *{{Structname}} {

	if _, ok := stage.{{Structname}}s[{{structname}}]; !ok {
		stage.{{Structname}}s[{{structname}}] = __member
		stage.{{Structname}}Map_Staged_Order[{{structname}}] = stage.{{Structname}}Order
		stage.{{Structname}}Order++
		stage.new[{{structname}}] = struct{}{}
		delete(stage.deleted, {{structname}})
	} else {
		if _, ok := stage.new[{{structname}}]; !ok {
			stage.modified[{{structname}}] = struct{}{}
		}
	}
	stage.{{Structname}}s_mapString[{{structname}}.Name] = {{structname}}

	return {{structname}}
}

// Unstage removes {{structname}} off the model stage
func ({{structname}} *{{Structname}}) Unstage(stage *Stage) *{{Structname}} {
	delete(stage.{{Structname}}s, {{structname}})
	delete(stage.{{Structname}}s_mapString, {{structname}}.Name)

	if _, ok := stage.reference[{{structname}}]; ok {
		stage.deleted[{{structname}}] = struct{}{}
	} else {
		delete(stage.new, {{structname}})
	}
	return {{structname}}
}

// UnstageVoid removes {{structname}} off the model stage
func ({{structname}} *{{Structname}}) UnstageVoid(stage *Stage) {
	delete(stage.{{Structname}}s, {{structname}})
	delete(stage.{{Structname}}s_mapString, {{structname}}.Name)
}

// commit {{structname}} to the back repo (if it is already staged)
func ({{structname}} *{{Structname}}) Commit(stage *Stage) *{{Structname}} {
	if _, ok := stage.{{Structname}}s[{{structname}}]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.Commit{{Structname}}({{structname}})
		}
	}
	return {{structname}}
}

func ({{structname}} *{{Structname}}) CommitVoid(stage *Stage) {
	{{structname}}.Commit(stage)
}

func ({{structname}} *{{Structname}}) StageVoid(stage *Stage) {
	{{structname}}.Stage(stage)
}

// Checkout {{structname}} to the back repo (if it is already staged)
func ({{structname}} *{{Structname}}) Checkout(stage *Stage) *{{Structname}} {
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

	// insertion point for slice of pointers maps{{SliceOfPointersReverseMaps}}
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
	stage.{{Structname}}Map_Staged_Order = make(map[*{{Structname}}]uint)
	stage.{{Structname}}Order = 0
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
	for idx, {{structname}} := range {{structname}}Ordered {

		id = generatesIdentifier("{{Structname}}", idx, {{structname}}.Name)
		map_{{Structname}}_Identifiers[{{structname}}] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "{{Structname}}")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", {{structname}}.Name)
		identifiersDecl += decl

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

	ModelGongStructInsertionGenericPointerToGongstructName: `
	case *{{Structname}}:
		res = "{{Structname}}"`,

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

	ModelGongOrderFields: `
	{{Structname}}Order            uint
	{{Structname}}Map_Staged_Order map[*{{Structname}}]uint
`,

	ModelGongOrderMapsInit: `
		{{Structname}}Map_Staged_Order: make(map[*{{Structname}}]uint),
`,
	ModelGongOrderSwitchGet: `
	case *{{Structname}}:
		return stage.{{Structname}}Map_Staged_Order[instance]`,

	ModelGongNamedStructsSliceInit: `
			{name: "{{Structname}}"},`,
	ModelGongNamedStructsInstancesNames: `
	case "{{Structname}}":
		res = GetNamedStructInstances(stage.{{Structname}}s, stage.{{Structname}}Map_Staged_Order)`,

	ModelGongNamedStructSortedOrderInstances: `
	case *{{Structname}}:
		tmp := GetStructInstancesByOrder(stage.{{Structname}}s, stage.{{Structname}}Map_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *{{Structname}} implements.
			res = append(res, any(v).(T))
		}
		return res`,
}

// Sub sub Templates identifiers per gong field
//
// For each gongstruct, a code snippet will be generated from each sub template
type GongFilePerStructSubTemplateId int

const (
	GongFileFieldSubTmplStringFieldName GongFilePerStructSubTemplateId = iota
	GongFileFieldSubTmplReverseField

	GongFileFieldSubTmplStringHeaderBasicKindField
	GongFileFieldSubTmplStringHeaderPointerField
	GongFileFieldSubTmplStringHeaderSliceOfPointersField

	GongFileFieldSubTmplStringValueBasicFieldBool
	GongFileFieldSubTmplStringValueBasicFieldInt
	GongFileFieldSubTmplStringValueBasicFieldIntDuration
	GongFileFieldSubTmplStringValueBasicFieldEnumString
	GongFileFieldSubTmplStringValueBasicFieldEnumInt
	GongFileFieldSubTmplStringValueBasicFieldFloat64
	GongFileFieldSubTmplStringValueBasicFieldString
	GongFileFieldSubTmplStringValueTimeField
	GongFileFieldSubTmplStringValueTimeFieldBespokeFormat
	GongFileFieldSubTmplStringValuePointerField
	GongFileFieldSubTmplStringValueSliceOfPointersField

	GongFileFieldSubTmplSetBasicString
	GongFileFieldSubTmplSetBasicBool
	GongFileFieldSubTmplSetBasicInt
	GongFileFieldSubTmplSetBasicFloat
	GongFileFieldSubTmplSetEnumString
	GongFileFieldSubTmplSetEnumInt
	GongFileFieldSubTmplSetPointer
	GongFileFieldSubTmplSetSliceOfPointers

	GongFileFieldSubTmplAssociationNamePointerField
	GongFileFieldSubTmplAssociationNameSliceOfPointersField
	GongFileFieldSubTmplAssociationPromotedNamePointerField
	GongFileFieldSubTmplAssociationPromotedNameSliceOfPointersField
	GongFileFieldSubTmplAssociationNameEnclosingCompositePointerField
	GongFileFieldSubTmplAssociationNameCompositePointerField

	GongFileFieldSubTmplPointerFieldPointerAssociationMapFunction
	GongFileFieldSubTmplPointerFieldSliceOfPointersAssociationMapFunction

	GongFileSliceOfPointersReverseMap
)

// for each sub template code, there is the sub template code
var GongFileFieldFieldSubTemplateCode map[GongFilePerStructSubTemplateId]string = // declaration of the sub templates
map[GongFilePerStructSubTemplateId]string{

	GongFileFieldSubTmplStringFieldName: `"{{FieldName}}"`,

	GongFileFieldSubTmplReverseField: `
		rf.GongstructName = "{{AssocStructName}}"
		rf.Fieldname = "{{FieldName}}"
		res = append(res, rf)`,

	GongFileFieldSubTmplStringHeaderBasicKindField: `
		{
			Name:               "{{FieldName}}",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},`,
	GongFileFieldSubTmplStringHeaderPointerField: `
		{
			Name:               "{{FieldName}}",
			GongFieldValueType: GongFieldValueTypePointer,
			TargetGongstructName: "{{AssocStructName}}",
		},`,
	GongFileFieldSubTmplStringHeaderSliceOfPointersField: `
		{
			Name:                 "{{FieldName}}",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "{{AssocStructName}}",
		},`,
	GongFileFieldSubTmplStringValueBasicFieldBool: `
	case "{{FieldName}}":
		res.valueString = fmt.Sprintf("%t", {{structname}}.{{FieldName}})
		res.valueBool = {{structname}}.{{FieldName}}
		res.GongFieldValueType = GongFieldValueTypeBool`,
	GongFileFieldSubTmplStringValueBasicFieldInt: `
	case "{{FieldName}}":
		res.valueString = fmt.Sprintf("%d", {{structname}}.{{FieldName}})
		res.valueInt = {{structname}}.{{FieldName}}
		res.GongFieldValueType = GongFieldValueTypeInt`,
	GongFileFieldSubTmplStringValueBasicFieldIntDuration: `
	case "{{FieldName}}":
		if math.Abs({{structname}}.{{FieldName}}.Hours()) >= 24 {
			days := __Gong__Abs(int(int({{structname}}.{{FieldName}}.Hours()) / 24))
			months := int(days / 31)
			days = days - months*31

			remainingHours := int({{structname}}.{{FieldName}}.Hours()) % 24
			remainingMinutes := int({{structname}}.{{FieldName}}.Minutes()) % 60
			remainingSeconds := int({{structname}}.{{FieldName}}.Seconds()) % 60

			if {{structname}}.{{FieldName}}.Hours() < 0 {
				res.valueString = "- "
			}

			if months > 0 {
				if months > 1 {
					res.valueString = res.valueString + fmt.Sprintf("%d months", months)
				} else {
					res.valueString = res.valueString + fmt.Sprintf("%d month", months)
				}
			}
			if days > 0 {
				if months != 0 {
					res.valueString = res.valueString + ", "
				}
				if days > 1 {
					res.valueString = res.valueString + fmt.Sprintf("%d days", days)
				} else {
					res.valueString = res.valueString + fmt.Sprintf("%d day", days)
				}

			}
			if remainingHours != 0 || remainingMinutes != 0 || remainingSeconds != 0 {
				if days != 0 || (days == 0 && months != 0) {
					res.valueString = res.valueString + ", "
				}
				res.valueString = res.valueString + fmt.Sprintf("%d hours, %d minutes, %d seconds\n", remainingHours, remainingMinutes, remainingSeconds)
			}
		} else {
			res.valueString = fmt.Sprintf("%s\n", {{structname}}.{{FieldName}}.String())
		}`,
	GongFileFieldSubTmplStringValueBasicFieldEnumString: `
	case "{{FieldName}}":
		enum := {{structname}}.{{FieldName}}
		res.valueString = enum.ToCodeString()`,
	GongFileFieldSubTmplStringValueBasicFieldEnumInt: `
	case "{{FieldName}}":
		enum := {{structname}}.{{FieldName}}
		res.valueString = enum.ToCodeString()`,
	GongFileFieldSubTmplStringValueBasicFieldFloat64: `
	case "{{FieldName}}":
		res.valueString = fmt.Sprintf("%f", {{structname}}.{{FieldName}})
		res.valueFloat = {{structname}}.{{FieldName}}
		res.GongFieldValueType = GongFieldValueTypeFloat`,
	GongFileFieldSubTmplStringValueBasicFieldString: `
	case "{{FieldName}}":
		res.valueString = {{structname}}.{{FieldName}}`,
	GongFileFieldSubTmplStringValueTimeField: `
	case "{{FieldName}}":
		res.valueString = {{structname}}.{{FieldName}}.String()`,
	GongFileFieldSubTmplStringValueTimeFieldBespokeFormat: `
	case "{{FieldName}}":
		res.valueString = {{structname}}.{{FieldName}}.Format("{{TimeFormat}}")`,
	GongFileFieldSubTmplStringValuePointerField: `
	case "{{FieldName}}":
		res.GongFieldValueType = GongFieldValueTypePointer
		if {{structname}}.{{FieldName}} != nil {
			res.valueString = {{structname}}.{{FieldName}}.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, {{structname}}.{{FieldName}}))
		}`,
	GongFileFieldSubTmplStringValueSliceOfPointersField: `
	case "{{FieldName}}":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range {{structname}}.{{FieldName}} {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}`,

	GongFileFieldSubTmplSetBasicString: `
	case "{{FieldName}}":
		{{structname}}.{{FieldName}} = value.GetValueString()`,
	GongFileFieldSubTmplSetBasicBool: `
	case "{{FieldName}}":
		{{structname}}.{{FieldName}} = value.GetValueBool()`,
	GongFileFieldSubTmplSetBasicInt: `
	case "{{FieldName}}":
		{{structname}}.{{FieldName}} = int(value.GetValueInt())`,
	GongFileFieldSubTmplSetBasicFloat: `
	case "{{FieldName}}":
		{{structname}}.{{FieldName}} = value.GetValueFloat()`,
	GongFileFieldSubTmplSetEnumString: `
	case "{{FieldName}}":
		{{structname}}.{{FieldName}}.FromCodeString(value.GetValueString())`,
	GongFileFieldSubTmplSetEnumInt: `
	case "{{FieldName}}":
		{{structname}}.{{FieldName}}.FromCodeString(value.GetValueString())`,
	GongFileFieldSubTmplSetPointer: `
	case "{{FieldName}}":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			{{structname}}.{{FieldName}} = nil
			for __instance__ := range stage.{{AssocStructName}}s {
				if stage.{{AssocStructName}}Map_Staged_Order[__instance__] == uint(id) {
					{{structname}}.{{FieldName}} = __instance__
					break
				}
			}
		}`,
	GongFileFieldSubTmplSetSliceOfPointers: `
	case "{{FieldName}}":
		{{structname}}.{{FieldName}} = make([]*{{AssocStructName}}, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.{{AssocStructName}}s {
					if stage.{{AssocStructName}}Map_Staged_Order[__instance__] == uint(id) {
						{{structname}}.{{FieldName}} = append({{structname}}.{{FieldName}}, __instance__)
						break
					}
				}
			}
		}`,

	GongFileFieldSubTmplAssociationNamePointerField: `
			// field is initialized with an instance of {{AssocStructName}} with the name of the field
			{{FieldName}}: &{{AssocStructName}}{Name: "{{FieldName}}"},`,

	GongFileFieldSubTmplAssociationNameSliceOfPointersField: `
			// field is initialized with an instance of {{AssocStructName}} with the name of the field
			{{FieldName}}: []*{{AssocStructName}}{{Name: "{{FieldName}}"}},`,

	GongFileFieldSubTmplAssociationPromotedNamePointerField: `
			// field is initialized with an instance of {{AssocStructName}} with the name of the field
			{{FieldName}}: &{{AssocStructName}}{{{CompositeAssocStructName}}: {{CompositeAssocStructName}}{Name: "{{FieldName}}"}},`,

	GongFileFieldSubTmplAssociationPromotedNameSliceOfPointersField: `
			// field is initialized with an instance of {{AssocStructName}} with the name of the field
			{{FieldName}}: []*{{AssocStructName}}{{{{CompositeAssocStructName}}: {{CompositeAssocStructName}}{Name: "{{FieldName}}"}}},`,

	// GongFileFieldSubTmplAssociationNameEnclosingCompositePointerField: `
	// 		// field is initialized with {{AssocCompositeStructName}} as it is a composite
	// 		{{AssocCompositeStructName}}: {{AssocCompositeStructName}}{
	// 			// per field init{{PerCompositeFieldInit}}
	// 		},`,
	GongFileFieldSubTmplAssociationNameEnclosingCompositePointerField: `
			// field is initialized with {{AssocCompositeStructName}} problem with composites
`,
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

	GongFileSliceOfPointersReverseMap: `
	{{Structname}}_{{FieldName}}_reverseMap map[*{{AssocStructName}}]*{{Structname}}
`,

	GongFileFieldSubTmplPointerFieldSliceOfPointersAssociationMapFunction: `
		case "{{FieldName}}":
			res := make(map[*{{AssocStructName}}][]*{{Structname}})
			for {{structname}} := range stage.{{Structname}}s {
				for _, {{assocstructname}}_ := range {{structname}}.{{FieldName}} {
					res[{{assocstructname}}_] = append(res[{{assocstructname}}_], {{structname}})
				}
			}
			return any(res).(map[*End][]*Start)`,
}

func CodeGeneratorModelGong(
	modelPkg *models.ModelPkg,
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

		for subStructTemplate := range ModelGongStructSubTemplateCode {

			// replace {{ValuesInitialization}}
			fieldNames := `
		res = []string{`
			fieldHeaders := `
	res = []GongFieldHeader{`
			reverseFields := `
		var rf ReverseField
		_ = rf`
			fieldStringValues := ``
			fieldSetValues := ``
			fieldReversePointerAssociationMapCode := ``
			fieldReverseSliceOfPointersAssociationMapCode := ``
			associationFieldInitialization := ``
			sliceOfPointersReverseMapStorageCode := ``

			associationFieldInitializationPerCompositeStruct := make(map[string]string, 0)

			for idx, field := range gongStruct.Fields {

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
				case *models.GongBasicField:

					switch field.GetBasicKind() {
					case types.String:
						if field.GongEnum == nil {
							fieldStringValues += models.Replace1(
								GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringValueBasicFieldString],
								"{{FieldName}}", field.Name)
							fieldSetValues += models.Replace1(
								GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetBasicString],
								"{{FieldName}}", field.Name)
						} else {
							fieldStringValues += models.Replace1(
								GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringValueBasicFieldEnumString],
								"{{FieldName}}", field.Name)
							fieldSetValues += models.Replace1(
								GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetEnumString],
								"{{FieldName}}", field.Name)
						}
					case types.Bool:
						fieldStringValues += models.Replace1(
							GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringValueBasicFieldBool],
							"{{FieldName}}", field.Name)
						fieldSetValues += models.Replace1(
							GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetBasicBool],
							"{{FieldName}}", field.Name)
					case types.Float64:
						fieldStringValues += models.Replace1(
							GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringValueBasicFieldFloat64],
							"{{FieldName}}", field.Name)
						fieldSetValues += models.Replace1(
							GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetBasicFloat],
							"{{FieldName}}", field.Name)
					case types.Int, types.Int64:
						if field.DeclaredType == "time.Duration" {
							fieldStringValues += models.Replace1(
								GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringValueBasicFieldIntDuration],
								"{{FieldName}}", field.Name)
							break
						}
						if field.GongEnum == nil {
							fieldStringValues += models.Replace1(
								GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringValueBasicFieldInt],
								"{{FieldName}}", field.Name)
							fieldSetValues += models.Replace1(
								GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetBasicInt],
								"{{FieldName}}", field.Name)
							break
						}
						fieldStringValues += models.Replace1(
							GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringValueBasicFieldEnumInt],
							"{{FieldName}}", field.Name)
						fieldSetValues += models.Replace1(
							GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetEnumInt],
							"{{FieldName}}", field.Name)
					default:
					}
					fieldHeaders += models.Replace1(
						GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringHeaderBasicKindField],
						"{{FieldName}}", field.GetName())
				case *models.GongTimeField:
					if field.BespokeTimeFormat == "" {
						fieldStringValues += models.Replace1(
							GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringValueTimeField],
							"{{FieldName}}", field.Name)
					} else {
						fieldStringValues += models.Replace2(
							GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringValueTimeFieldBespokeFormat],
							"{{FieldName}}", field.Name,
							"{{TimeFormat}}", field.BespokeTimeFormat,
						)
					}
					fieldHeaders += models.Replace1(
						GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringHeaderBasicKindField],
						"{{FieldName}}", field.GetName())
				case *models.PointerToGongStructField:
					fieldStringValues += models.Replace1(
						GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringValuePointerField],
						"{{FieldName}}", field.Name)
					fieldSetValues += models.Replace2(
						GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetPointer],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name)
					fieldReversePointerAssociationMapCode += models.Replace3(
						GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplPointerFieldPointerAssociationMapFunction],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name,
						"{{assocstructname}}", strings.ToLower(field.GongStruct.Name))

					// for the case where the "Name" is a promoted field, one cannot generate
					// the code (cf. waiting for https://github.com/golang/go/issues/9859)
					var assocStructNameHasNameAsPromotedField bool
					var assocCompositeStrucName string
					for _, __field := range field.GongStruct.Fields {
						if __field.GetName() == "Name" && __field.GetCompositeStructName() != "" {
							assocStructNameHasNameAsPromotedField = true
							assocCompositeStrucName = __field.GetCompositeStructName()
						}
					}
					if field.CompositeStructName == "" && !isWithinAnonymousStruct && !assocStructNameHasNameAsPromotedField {
						associationFieldInitialization += models.Replace3(
							GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplAssociationNamePointerField],
							"{{FieldName}}", field.Name,
							"{{AssocStructName}}", field.GongStruct.Name,
							"{{assocstructname}}", strings.ToLower(field.GongStruct.Name))
					} else {
						if !isWithinAnonymousStruct && !assocStructNameHasNameAsPromotedField {
							associationFieldInitializationPerCompositeStruct[field.CompositeStructName] += models.Replace4(
								GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplAssociationNameCompositePointerField],
								"{{FieldName}}", field.Name,
								"{{AssocStructName}}", field.GongStruct.Name,
								"{{AssocCompositeStructName}}", field.CompositeStructName,
								"{{assocstructname}}", strings.ToLower(field.GongStruct.Name))
						}
					}
					if assocStructNameHasNameAsPromotedField {
						associationFieldInitialization += models.Replace4(
							GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplAssociationPromotedNamePointerField],
							"{{FieldName}}", field.Name,
							"{{AssocStructName}}", field.GongStruct.Name,
							"{{assocstructname}}", strings.ToLower(field.GongStruct.Name),
							"{{CompositeAssocStructName}}", assocCompositeStrucName,
						)
					}
					fieldHeaders += models.Replace2(
						GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringHeaderPointerField],
						"{{AssocStructName}}", field.GongStruct.Name,
						"{{FieldName}}", field.GetName())
				case *models.SliceOfPointerToGongStructField:
					fieldStringValues += models.Replace1(
						GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringValueSliceOfPointersField],
						"{{FieldName}}", field.Name)
					fieldSetValues += models.Replace2(
						GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplSetSliceOfPointers],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name)
					fieldReverseSliceOfPointersAssociationMapCode += models.Replace3(
						GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplPointerFieldSliceOfPointersAssociationMapFunction],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name,
						"{{assocstructname}}", strings.ToLower(field.GongStruct.Name))

					sliceOfPointersReverseMapStorageCode += models.Replace3(
						GongFileFieldFieldSubTemplateCode[GongFileSliceOfPointersReverseMap],
						"{{FieldName}}", fieldNameForReverseMapField,
						"{{AssocStructName}}", field.GongStruct.Name,
						"{{assocstructname}}", strings.ToLower(field.GongStruct.Name))

					// for the case where the "Name" is a promoted field, one cannot generate
					// the code (cf. waiting for https://github.com/golang/go/issues/9859)
					var assocStructNameHasNameAsPromotedField bool
					var assocCompositeStrucName string
					for _, __field := range field.GongStruct.Fields {
						if __field.GetName() == "Name" && __field.GetCompositeStructName() != "" {
							assocStructNameHasNameAsPromotedField = true
							assocCompositeStrucName = __field.GetCompositeStructName()
						}
					}
					if field.CompositeStructName == "" && !isWithinAnonymousStruct && !assocStructNameHasNameAsPromotedField {
						associationFieldInitialization += models.Replace3(
							GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplAssociationNameSliceOfPointersField],
							"{{FieldName}}", field.Name,
							"{{AssocStructName}}", field.GongStruct.Name,
							"{{assocstructname}}", strings.ToLower(field.GongStruct.Name))
					}
					if assocStructNameHasNameAsPromotedField {
						associationFieldInitialization += models.Replace4(
							GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplAssociationPromotedNameSliceOfPointersField],
							"{{FieldName}}", field.Name,
							"{{AssocStructName}}", field.GongStruct.Name,
							"{{assocstructname}}", strings.ToLower(field.GongStruct.Name),
							"{{CompositeAssocStructName}}", assocCompositeStrucName,
						)
					}
					fieldHeaders += models.Replace2(
						GongFileFieldFieldSubTemplateCode[GongFileFieldSubTmplStringHeaderSliceOfPointersField],
						"{{AssocStructName}}", field.GongStruct.Name,
						"{{FieldName}}", field.GetName())
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

			fieldSetValues = models.Replace2(fieldSetValues,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			fieldReversePointerAssociationMapCode = models.Replace2(fieldReversePointerAssociationMapCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			fieldReverseSliceOfPointersAssociationMapCode = models.Replace2(fieldReverseSliceOfPointersAssociationMapCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			sliceOfPointersReverseMapStorageCode = models.Replace2(sliceOfPointersReverseMapStorageCode,
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name)

			fieldNames += `}`
			fieldHeaders += `
	}`

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

			generatedCodeFromSubTemplate := models.Replace11(ModelGongStructSubTemplateCode[subStructTemplate],
				"{{structname}}", strings.ToLower(gongStruct.Name),
				"{{Structname}}", gongStruct.Name,
				"{{ListOfFieldsName}}", fieldNames,
				"{{ListOfFieldHeaders}}", fieldHeaders,
				"{{ListOfReverseFields}}", reverseFields,
				"{{StringValueOfFields}}", fieldStringValues,
				"{{SetFieldValues}}", fieldSetValues,
				"{{fieldReversePointerAssociationMapCode}}", fieldReversePointerAssociationMapCode,
				"{{SliceOfPointersReverseMaps}}", sliceOfPointersReverseMapStorageCode,
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
	returnType := "*map[Type]any" // to solve the issue of empty model
	_ = returnType
	if len(modelPkg.GongStructs) == 0 {
		returnType = "any"
	}
	codeGO = models.Replace6(codeGO,
		"{{PkgName}}", pkgName,
		"{{TitlePkgName}}", caserEnglish.String(pkgName),
		"{{pkgname}}", strings.ToLower(pkgName),
		"	 | ", "	", // for the replacement of the of the first bar in the Gongstruct Type def,
		"{{mapReturnType}}", returnType,
		"{{PkgPathRoot}}", strings.ReplaceAll(modelPkg.PkgPath, "/models", ""),
	)

	file, err := os.Create(filepath.Join(pkgPath, string(models.GeneratedGongGoFilePath)))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeGO)

}
