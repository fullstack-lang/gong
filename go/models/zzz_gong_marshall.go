// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

const marshallRes = `package {{PackageName}}

import (
	"time"

	"{{ModelsPackageName}}"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage{{Identifiers}}

	// insertion point for initialization of values{{ValueInitializers}}

	// insertion point for setup of pointers{{PointersInitializers}}
}`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: ` + "`" + `{{GeneratedFieldNameValue}}` + "`" + `}).Stage(stage)`

// previous version does not hanldle embedded structs (https://github.com/golang/go/issues/9859)
// simpler version but the name of the instance cannot be human read before the fields initialization
const IdentifiersDeclsWithoutNameInit = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{}).Stage(stage)` /* */

const StringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + "`" + `{{GeneratedFieldNameValue}}` + "`"

const MetaFieldStructInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + `{{GeneratedFieldNameValue}}`

const StringEnumInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const NumberInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const PointerFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const SliceOfPointersFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = append({{Identifier}}.{{GeneratedFieldName}}, {{GeneratedFieldNameValue}})`

const TimeInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}}, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "{{GeneratedFieldNameValue}}")`

// Marshall marshall the stage content into the file as an instanciation into a stage
func (stage *Stage) Marshall(file *os.File, modelsPackageName, packageName string) {

	name := file.Name()

	if !strings.HasSuffix(name, ".go") {
		log.Println(name + " is not a go filename")
	}

	log.Printf("Marshalling %s", name)

	res, err := stage.MarshallToString(modelsPackageName, packageName)
	if err != nil {
		log.Fatalln("Error marshalling to string:", err)
	}

	fmt.Fprintln(file, res)
}

// MarshallToString marshall the stage content into a string
func (stage *Stage) MarshallToString(modelsPackageName, packageName string) (res string, err error) {

	res = marshallRes
	res = strings.ReplaceAll(res, "{{PackageName}}", packageName)
	res = strings.ReplaceAll(res, "{{ModelsPackageName}}", modelsPackageName)

	// map of identifiers
	// var StageMapDstructIds map[*Dstruct]string
	identifiersDecl := ""
	initializerStatements := ""
	pointersInitializesStatements := ""

	decl := ""
	_ = decl
	setValueField := ""
	_ = setValueField

	// insertion initialization of objects to stage
	gongbasicfieldOrdered := []*GongBasicField{}
	for gongbasicfield := range stage.GongBasicFields {
		gongbasicfieldOrdered = append(gongbasicfieldOrdered, gongbasicfield)
	}
	sort.Slice(gongbasicfieldOrdered[:], func(i, j int) bool {
		gongbasicfieldi := gongbasicfieldOrdered[i]
		gongbasicfieldj := gongbasicfieldOrdered[j]
		gongbasicfieldi_order, oki := stage.GongBasicFieldMap_Staged_Order[gongbasicfieldi]
		gongbasicfieldj_order, okj := stage.GongBasicFieldMap_Staged_Order[gongbasicfieldj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return gongbasicfieldi_order < gongbasicfieldj_order
	})
	if len(gongbasicfieldOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, gongbasicfield := range gongbasicfieldOrdered {

		identifiersDecl += gongbasicfield.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += gongbasicfield.GongMarshallField(stage, "Name")
		initializerStatements += gongbasicfield.GongMarshallField(stage, "BasicKindName")
		pointersInitializesStatements += gongbasicfield.GongMarshallField(stage, "GongEnum")
		initializerStatements += gongbasicfield.GongMarshallField(stage, "DeclaredType")
		initializerStatements += gongbasicfield.GongMarshallField(stage, "CompositeStructName")
		initializerStatements += gongbasicfield.GongMarshallField(stage, "Index")
		initializerStatements += gongbasicfield.GongMarshallField(stage, "IsTextArea")
		initializerStatements += gongbasicfield.GongMarshallField(stage, "IsBespokeWidth")
		initializerStatements += gongbasicfield.GongMarshallField(stage, "BespokeWidth")
		initializerStatements += gongbasicfield.GongMarshallField(stage, "IsBespokeHeight")
		initializerStatements += gongbasicfield.GongMarshallField(stage, "BespokeHeight")
	}

	gongenumOrdered := []*GongEnum{}
	for gongenum := range stage.GongEnums {
		gongenumOrdered = append(gongenumOrdered, gongenum)
	}
	sort.Slice(gongenumOrdered[:], func(i, j int) bool {
		gongenumi := gongenumOrdered[i]
		gongenumj := gongenumOrdered[j]
		gongenumi_order, oki := stage.GongEnumMap_Staged_Order[gongenumi]
		gongenumj_order, okj := stage.GongEnumMap_Staged_Order[gongenumj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return gongenumi_order < gongenumj_order
	})
	if len(gongenumOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, gongenum := range gongenumOrdered {

		identifiersDecl += gongenum.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += gongenum.GongMarshallField(stage, "Name")
		initializerStatements += gongenum.GongMarshallField(stage, "Type")
		pointersInitializesStatements += gongenum.GongMarshallField(stage, "GongEnumValues")
	}

	gongenumvalueOrdered := []*GongEnumValue{}
	for gongenumvalue := range stage.GongEnumValues {
		gongenumvalueOrdered = append(gongenumvalueOrdered, gongenumvalue)
	}
	sort.Slice(gongenumvalueOrdered[:], func(i, j int) bool {
		gongenumvaluei := gongenumvalueOrdered[i]
		gongenumvaluej := gongenumvalueOrdered[j]
		gongenumvaluei_order, oki := stage.GongEnumValueMap_Staged_Order[gongenumvaluei]
		gongenumvaluej_order, okj := stage.GongEnumValueMap_Staged_Order[gongenumvaluej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return gongenumvaluei_order < gongenumvaluej_order
	})
	if len(gongenumvalueOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, gongenumvalue := range gongenumvalueOrdered {

		identifiersDecl += gongenumvalue.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += gongenumvalue.GongMarshallField(stage, "Name")
		initializerStatements += gongenumvalue.GongMarshallField(stage, "Value")
	}

	gonglinkOrdered := []*GongLink{}
	for gonglink := range stage.GongLinks {
		gonglinkOrdered = append(gonglinkOrdered, gonglink)
	}
	sort.Slice(gonglinkOrdered[:], func(i, j int) bool {
		gonglinki := gonglinkOrdered[i]
		gonglinkj := gonglinkOrdered[j]
		gonglinki_order, oki := stage.GongLinkMap_Staged_Order[gonglinki]
		gonglinkj_order, okj := stage.GongLinkMap_Staged_Order[gonglinkj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return gonglinki_order < gonglinkj_order
	})
	if len(gonglinkOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, gonglink := range gonglinkOrdered {

		identifiersDecl += gonglink.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += gonglink.GongMarshallField(stage, "Name")
		initializerStatements += gonglink.GongMarshallField(stage, "Recv")
		initializerStatements += gonglink.GongMarshallField(stage, "ImportPath")
	}

	gongnoteOrdered := []*GongNote{}
	for gongnote := range stage.GongNotes {
		gongnoteOrdered = append(gongnoteOrdered, gongnote)
	}
	sort.Slice(gongnoteOrdered[:], func(i, j int) bool {
		gongnotei := gongnoteOrdered[i]
		gongnotej := gongnoteOrdered[j]
		gongnotei_order, oki := stage.GongNoteMap_Staged_Order[gongnotei]
		gongnotej_order, okj := stage.GongNoteMap_Staged_Order[gongnotej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return gongnotei_order < gongnotej_order
	})
	if len(gongnoteOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, gongnote := range gongnoteOrdered {

		identifiersDecl += gongnote.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += gongnote.GongMarshallField(stage, "Name")
		initializerStatements += gongnote.GongMarshallField(stage, "Body")
		initializerStatements += gongnote.GongMarshallField(stage, "BodyHTML")
		pointersInitializesStatements += gongnote.GongMarshallField(stage, "Links")
	}

	gongstructOrdered := []*GongStruct{}
	for gongstruct := range stage.GongStructs {
		gongstructOrdered = append(gongstructOrdered, gongstruct)
	}
	sort.Slice(gongstructOrdered[:], func(i, j int) bool {
		gongstructi := gongstructOrdered[i]
		gongstructj := gongstructOrdered[j]
		gongstructi_order, oki := stage.GongStructMap_Staged_Order[gongstructi]
		gongstructj_order, okj := stage.GongStructMap_Staged_Order[gongstructj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return gongstructi_order < gongstructj_order
	})
	if len(gongstructOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, gongstruct := range gongstructOrdered {

		identifiersDecl += gongstruct.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += gongstruct.GongMarshallField(stage, "Name")
		pointersInitializesStatements += gongstruct.GongMarshallField(stage, "GongBasicFields")
		pointersInitializesStatements += gongstruct.GongMarshallField(stage, "GongTimeFields")
		pointersInitializesStatements += gongstruct.GongMarshallField(stage, "PointerToGongStructFields")
		pointersInitializesStatements += gongstruct.GongMarshallField(stage, "SliceOfPointerToGongStructFields")
		initializerStatements += gongstruct.GongMarshallField(stage, "HasOnAfterUpdateSignature")
		initializerStatements += gongstruct.GongMarshallField(stage, "IsIgnoredForFront")
	}

	gongtimefieldOrdered := []*GongTimeField{}
	for gongtimefield := range stage.GongTimeFields {
		gongtimefieldOrdered = append(gongtimefieldOrdered, gongtimefield)
	}
	sort.Slice(gongtimefieldOrdered[:], func(i, j int) bool {
		gongtimefieldi := gongtimefieldOrdered[i]
		gongtimefieldj := gongtimefieldOrdered[j]
		gongtimefieldi_order, oki := stage.GongTimeFieldMap_Staged_Order[gongtimefieldi]
		gongtimefieldj_order, okj := stage.GongTimeFieldMap_Staged_Order[gongtimefieldj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return gongtimefieldi_order < gongtimefieldj_order
	})
	if len(gongtimefieldOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, gongtimefield := range gongtimefieldOrdered {

		identifiersDecl += gongtimefield.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += gongtimefield.GongMarshallField(stage, "Name")
		initializerStatements += gongtimefield.GongMarshallField(stage, "Index")
		initializerStatements += gongtimefield.GongMarshallField(stage, "CompositeStructName")
		initializerStatements += gongtimefield.GongMarshallField(stage, "BespokeTimeFormat")
	}

	metareferenceOrdered := []*MetaReference{}
	for metareference := range stage.MetaReferences {
		metareferenceOrdered = append(metareferenceOrdered, metareference)
	}
	sort.Slice(metareferenceOrdered[:], func(i, j int) bool {
		metareferencei := metareferenceOrdered[i]
		metareferencej := metareferenceOrdered[j]
		metareferencei_order, oki := stage.MetaReferenceMap_Staged_Order[metareferencei]
		metareferencej_order, okj := stage.MetaReferenceMap_Staged_Order[metareferencej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return metareferencei_order < metareferencej_order
	})
	if len(metareferenceOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, metareference := range metareferenceOrdered {

		identifiersDecl += metareference.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += metareference.GongMarshallField(stage, "Name")
	}

	modelpkgOrdered := []*ModelPkg{}
	for modelpkg := range stage.ModelPkgs {
		modelpkgOrdered = append(modelpkgOrdered, modelpkg)
	}
	sort.Slice(modelpkgOrdered[:], func(i, j int) bool {
		modelpkgi := modelpkgOrdered[i]
		modelpkgj := modelpkgOrdered[j]
		modelpkgi_order, oki := stage.ModelPkgMap_Staged_Order[modelpkgi]
		modelpkgj_order, okj := stage.ModelPkgMap_Staged_Order[modelpkgj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return modelpkgi_order < modelpkgj_order
	})
	if len(modelpkgOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, modelpkg := range modelpkgOrdered {

		identifiersDecl += modelpkg.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += modelpkg.GongMarshallField(stage, "Name")
		initializerStatements += modelpkg.GongMarshallField(stage, "PkgPath")
		initializerStatements += modelpkg.GongMarshallField(stage, "PathToGoSubDirectory")
		initializerStatements += modelpkg.GongMarshallField(stage, "OrmPkgGenPath")
		initializerStatements += modelpkg.GongMarshallField(stage, "DbOrmPkgGenPath")
		initializerStatements += modelpkg.GongMarshallField(stage, "DbLiteOrmPkgGenPath")
		initializerStatements += modelpkg.GongMarshallField(stage, "DbPkgGenPath")
		initializerStatements += modelpkg.GongMarshallField(stage, "ControllersPkgGenPath")
		initializerStatements += modelpkg.GongMarshallField(stage, "FullstackPkgGenPath")
		initializerStatements += modelpkg.GongMarshallField(stage, "StackPkgGenPath")
		initializerStatements += modelpkg.GongMarshallField(stage, "Level1StackPkgGenPath")
		initializerStatements += modelpkg.GongMarshallField(stage, "StaticPkgGenPath")
		initializerStatements += modelpkg.GongMarshallField(stage, "ProbePkgGenPath")
		initializerStatements += modelpkg.GongMarshallField(stage, "NgWorkspacePath")
		initializerStatements += modelpkg.GongMarshallField(stage, "NgWorkspaceName")
		initializerStatements += modelpkg.GongMarshallField(stage, "NgDataLibrarySourceCodeDirectory")
		initializerStatements += modelpkg.GongMarshallField(stage, "NgSpecificLibrarySourceCodeDirectory")
		initializerStatements += modelpkg.GongMarshallField(stage, "MaterialLibDatamodelTargetPath")
	}

	pointertogongstructfieldOrdered := []*PointerToGongStructField{}
	for pointertogongstructfield := range stage.PointerToGongStructFields {
		pointertogongstructfieldOrdered = append(pointertogongstructfieldOrdered, pointertogongstructfield)
	}
	sort.Slice(pointertogongstructfieldOrdered[:], func(i, j int) bool {
		pointertogongstructfieldi := pointertogongstructfieldOrdered[i]
		pointertogongstructfieldj := pointertogongstructfieldOrdered[j]
		pointertogongstructfieldi_order, oki := stage.PointerToGongStructFieldMap_Staged_Order[pointertogongstructfieldi]
		pointertogongstructfieldj_order, okj := stage.PointerToGongStructFieldMap_Staged_Order[pointertogongstructfieldj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return pointertogongstructfieldi_order < pointertogongstructfieldj_order
	})
	if len(pointertogongstructfieldOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, pointertogongstructfield := range pointertogongstructfieldOrdered {

		identifiersDecl += pointertogongstructfield.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += pointertogongstructfield.GongMarshallField(stage, "Name")
		pointersInitializesStatements += pointertogongstructfield.GongMarshallField(stage, "GongStruct")
		initializerStatements += pointertogongstructfield.GongMarshallField(stage, "Index")
		initializerStatements += pointertogongstructfield.GongMarshallField(stage, "CompositeStructName")
		initializerStatements += pointertogongstructfield.GongMarshallField(stage, "IsType")
	}

	sliceofpointertogongstructfieldOrdered := []*SliceOfPointerToGongStructField{}
	for sliceofpointertogongstructfield := range stage.SliceOfPointerToGongStructFields {
		sliceofpointertogongstructfieldOrdered = append(sliceofpointertogongstructfieldOrdered, sliceofpointertogongstructfield)
	}
	sort.Slice(sliceofpointertogongstructfieldOrdered[:], func(i, j int) bool {
		sliceofpointertogongstructfieldi := sliceofpointertogongstructfieldOrdered[i]
		sliceofpointertogongstructfieldj := sliceofpointertogongstructfieldOrdered[j]
		sliceofpointertogongstructfieldi_order, oki := stage.SliceOfPointerToGongStructFieldMap_Staged_Order[sliceofpointertogongstructfieldi]
		sliceofpointertogongstructfieldj_order, okj := stage.SliceOfPointerToGongStructFieldMap_Staged_Order[sliceofpointertogongstructfieldj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return sliceofpointertogongstructfieldi_order < sliceofpointertogongstructfieldj_order
	})
	if len(sliceofpointertogongstructfieldOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, sliceofpointertogongstructfield := range sliceofpointertogongstructfieldOrdered {

		identifiersDecl += sliceofpointertogongstructfield.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += sliceofpointertogongstructfield.GongMarshallField(stage, "Name")
		pointersInitializesStatements += sliceofpointertogongstructfield.GongMarshallField(stage, "GongStruct")
		initializerStatements += sliceofpointertogongstructfield.GongMarshallField(stage, "Index")
		initializerStatements += sliceofpointertogongstructfield.GongMarshallField(stage, "CompositeStructName")
	}

	// insertion initialization of objects to stage
	for _, gongbasicfield := range gongbasicfieldOrdered {
		_ = gongbasicfield
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, gongenum := range gongenumOrdered {
		_ = gongenum
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, gongenumvalue := range gongenumvalueOrdered {
		_ = gongenumvalue
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, gonglink := range gonglinkOrdered {
		_ = gonglink
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, gongnote := range gongnoteOrdered {
		_ = gongnote
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, gongstruct := range gongstructOrdered {
		_ = gongstruct
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, gongtimefield := range gongtimefieldOrdered {
		_ = gongtimefield
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, metareference := range metareferenceOrdered {
		_ = metareference
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, modelpkg := range modelpkgOrdered {
		_ = modelpkg
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, pointertogongstructfield := range pointertogongstructfieldOrdered {
		_ = pointertogongstructfield
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, sliceofpointertogongstructfield := range sliceofpointertogongstructfieldOrdered {
		_ = sliceofpointertogongstructfield
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl)
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements)
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements)

	if stage.MetaPackageImportAlias != "" {
		res = strings.ReplaceAll(res, "{{ImportPackageDeclaration}}",
			fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))

		res = strings.ReplaceAll(res, "{{ImportPackageDummyDeclaration}}",
			fmt.Sprintf("\nvar _ %s.Stage",
				stage.MetaPackageImportAlias))

		var entries string

		// regenerate the map of doc link renaming
		// the key and value are set to the value because
		// if it has been renamed, this is the new value that matters
		valuesOrdered := make([]GONG__Identifier, 0)
		for _, value := range stage.Map_DocLink_Renaming {
			valuesOrdered = append(valuesOrdered, value)
		}
		sort.Slice(valuesOrdered[:], func(i, j int) bool {
			return valuesOrdered[i].Ident < valuesOrdered[j].Ident
		})
		for _, value := range valuesOrdered {

			// get the number of points in the value to find if it is a field
			// or a struct

			switch value.Type {
			case GONG__ENUM_CAST_INT:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s(0),", value.Ident, value.Ident)
			case GONG__ENUM_CAST_STRING:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s(\"\"),", value.Ident, value.Ident)
			case GONG__FIELD_VALUE:
				// substitute the second point with "{})."
				joker := "__substitute_for_first_point__"
				valueIdentifier := strings.Replace(value.Ident, ".", joker, 1)
				valueIdentifier = strings.Replace(valueIdentifier, ".", "{}).", 1)
				valueIdentifier = strings.Replace(valueIdentifier, joker, ".", 1)
				entries += fmt.Sprintf("\n\n\t\"%s\": (%s,", value.Ident, valueIdentifier)
			case GONG__IDENTIFIER_CONST:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s,", value.Ident, value.Ident)
			case GONG__STRUCT_INSTANCE:
				entries += fmt.Sprintf("\n\n\t\"%s\": &(%s{}),", value.Ident, value.Ident)
			}
		}

		// res = strings.ReplaceAll(res, "{{EntriesDocLinkStringDocLinkIdentifier}}", entries)
	}
	return
}

// insertion point for marshall field methods
func (gongbasicfield *GongBasicField) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gongbasicfield.Name))
	case "BasicKindName":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BasicKindName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gongbasicfield.BasicKindName))
	case "DeclaredType":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DeclaredType")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gongbasicfield.DeclaredType))
	case "CompositeStructName":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CompositeStructName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gongbasicfield.CompositeStructName))
	case "Index":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Index")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", gongbasicfield.Index))
	case "IsTextArea":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsTextArea")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongbasicfield.IsTextArea))
	case "IsBespokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsBespokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongbasicfield.IsBespokeWidth))
	case "BespokeWidth":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BespokeWidth")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", gongbasicfield.BespokeWidth))
	case "IsBespokeHeight":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsBespokeHeight")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongbasicfield.IsBespokeHeight))
	case "BespokeHeight":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BespokeHeight")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", gongbasicfield.BespokeHeight))

	case "GongEnum":
		if gongbasicfield.GongEnum != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GongEnum")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", gongbasicfield.GongEnum.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct GongBasicField", fieldName)
	}
	return
}

func (gongenum *GongEnum) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongenum.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gongenum.Name))
	case "Type":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongenum.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Type")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+gongenum.Type.ToCodeString())

	case "GongEnumValues":
		for _, _gongenumvalue := range gongenum.GongEnumValues {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", gongenum.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "GongEnumValues")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _gongenumvalue.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct GongEnum", fieldName)
	}
	return
}

func (gongenumvalue *GongEnumValue) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongenumvalue.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gongenumvalue.Name))
	case "Value":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongenumvalue.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Value")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gongenumvalue.Value))

	default:
		log.Panicf("Unknown field %s for Gongstruct GongEnumValue", fieldName)
	}
	return
}

func (gonglink *GongLink) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gonglink.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gonglink.Name))
	case "Recv":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gonglink.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Recv")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gonglink.Recv))
	case "ImportPath":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gonglink.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ImportPath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gonglink.ImportPath))

	default:
		log.Panicf("Unknown field %s for Gongstruct GongLink", fieldName)
	}
	return
}

func (gongnote *GongNote) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongnote.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gongnote.Name))
	case "Body":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongnote.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Body")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gongnote.Body))
	case "BodyHTML":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongnote.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BodyHTML")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gongnote.BodyHTML))

	case "Links":
		for _, _gonglink := range gongnote.Links {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", gongnote.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Links")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _gonglink.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct GongNote", fieldName)
	}
	return
}

func (gongstruct *GongStruct) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongstruct.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gongstruct.Name))
	case "HasOnAfterUpdateSignature":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongstruct.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasOnAfterUpdateSignature")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongstruct.HasOnAfterUpdateSignature))
	case "IsIgnoredForFront":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongstruct.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsIgnoredForFront")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongstruct.IsIgnoredForFront))

	case "GongBasicFields":
		for _, _gongbasicfield := range gongstruct.GongBasicFields {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", gongstruct.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "GongBasicFields")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _gongbasicfield.GongGetIdentifier(stage))
			res += tmp
		}
	case "GongTimeFields":
		for _, _gongtimefield := range gongstruct.GongTimeFields {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", gongstruct.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "GongTimeFields")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _gongtimefield.GongGetIdentifier(stage))
			res += tmp
		}
	case "PointerToGongStructFields":
		for _, _pointertogongstructfield := range gongstruct.PointerToGongStructFields {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", gongstruct.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "PointerToGongStructFields")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _pointertogongstructfield.GongGetIdentifier(stage))
			res += tmp
		}
	case "SliceOfPointerToGongStructFields":
		for _, _sliceofpointertogongstructfield := range gongstruct.SliceOfPointerToGongStructFields {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", gongstruct.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SliceOfPointerToGongStructFields")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _sliceofpointertogongstructfield.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct GongStruct", fieldName)
	}
	return
}

func (gongtimefield *GongTimeField) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongtimefield.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gongtimefield.Name))
	case "Index":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongtimefield.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Index")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", gongtimefield.Index))
	case "CompositeStructName":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongtimefield.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CompositeStructName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gongtimefield.CompositeStructName))
	case "BespokeTimeFormat":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gongtimefield.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BespokeTimeFormat")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gongtimefield.BespokeTimeFormat))

	default:
		log.Panicf("Unknown field %s for Gongstruct GongTimeField", fieldName)
	}
	return
}

func (metareference *MetaReference) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", metareference.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(metareference.Name))

	default:
		log.Panicf("Unknown field %s for Gongstruct MetaReference", fieldName)
	}
	return
}

func (modelpkg *ModelPkg) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(modelpkg.Name))
	case "PkgPath":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "PkgPath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(modelpkg.PkgPath))
	case "PathToGoSubDirectory":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "PathToGoSubDirectory")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(modelpkg.PathToGoSubDirectory))
	case "OrmPkgGenPath":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "OrmPkgGenPath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(modelpkg.OrmPkgGenPath))
	case "DbOrmPkgGenPath":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DbOrmPkgGenPath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(modelpkg.DbOrmPkgGenPath))
	case "DbLiteOrmPkgGenPath":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DbLiteOrmPkgGenPath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(modelpkg.DbLiteOrmPkgGenPath))
	case "DbPkgGenPath":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DbPkgGenPath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(modelpkg.DbPkgGenPath))
	case "ControllersPkgGenPath":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ControllersPkgGenPath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(modelpkg.ControllersPkgGenPath))
	case "FullstackPkgGenPath":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "FullstackPkgGenPath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(modelpkg.FullstackPkgGenPath))
	case "StackPkgGenPath":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StackPkgGenPath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(modelpkg.StackPkgGenPath))
	case "Level1StackPkgGenPath":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Level1StackPkgGenPath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(modelpkg.Level1StackPkgGenPath))
	case "StaticPkgGenPath":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StaticPkgGenPath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(modelpkg.StaticPkgGenPath))
	case "ProbePkgGenPath":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ProbePkgGenPath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(modelpkg.ProbePkgGenPath))
	case "NgWorkspacePath":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NgWorkspacePath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(modelpkg.NgWorkspacePath))
	case "NgWorkspaceName":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NgWorkspaceName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(modelpkg.NgWorkspaceName))
	case "NgDataLibrarySourceCodeDirectory":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NgDataLibrarySourceCodeDirectory")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(modelpkg.NgDataLibrarySourceCodeDirectory))
	case "NgSpecificLibrarySourceCodeDirectory":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NgSpecificLibrarySourceCodeDirectory")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(modelpkg.NgSpecificLibrarySourceCodeDirectory))
	case "MaterialLibDatamodelTargetPath":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MaterialLibDatamodelTargetPath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(modelpkg.MaterialLibDatamodelTargetPath))

	default:
		log.Panicf("Unknown field %s for Gongstruct ModelPkg", fieldName)
	}
	return
}

func (pointertogongstructfield *PointerToGongStructField) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", pointertogongstructfield.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(pointertogongstructfield.Name))
	case "Index":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", pointertogongstructfield.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Index")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", pointertogongstructfield.Index))
	case "CompositeStructName":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", pointertogongstructfield.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CompositeStructName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(pointertogongstructfield.CompositeStructName))
	case "IsType":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", pointertogongstructfield.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsType")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", pointertogongstructfield.IsType))

	case "GongStruct":
		if pointertogongstructfield.GongStruct != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", pointertogongstructfield.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GongStruct")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", pointertogongstructfield.GongStruct.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct PointerToGongStructField", fieldName)
	}
	return
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", sliceofpointertogongstructfield.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(sliceofpointertogongstructfield.Name))
	case "Index":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", sliceofpointertogongstructfield.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Index")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", sliceofpointertogongstructfield.Index))
	case "CompositeStructName":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", sliceofpointertogongstructfield.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CompositeStructName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(sliceofpointertogongstructfield.CompositeStructName))

	case "GongStruct":
		if sliceofpointertogongstructfield.GongStruct != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", sliceofpointertogongstructfield.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "GongStruct")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", sliceofpointertogongstructfield.GongStruct.GongGetIdentifier(stage))
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SliceOfPointerToGongStructField", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (gongbasicfield *GongBasicField) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += gongbasicfield.GongMarshallField(stage, "Name")
		initializerStatements += gongbasicfield.GongMarshallField(stage, "BasicKindName")
		pointersInitializesStatements += gongbasicfield.GongMarshallField(stage, "GongEnum")
		initializerStatements += gongbasicfield.GongMarshallField(stage, "DeclaredType")
		initializerStatements += gongbasicfield.GongMarshallField(stage, "CompositeStructName")
		initializerStatements += gongbasicfield.GongMarshallField(stage, "Index")
		initializerStatements += gongbasicfield.GongMarshallField(stage, "IsTextArea")
		initializerStatements += gongbasicfield.GongMarshallField(stage, "IsBespokeWidth")
		initializerStatements += gongbasicfield.GongMarshallField(stage, "BespokeWidth")
		initializerStatements += gongbasicfield.GongMarshallField(stage, "IsBespokeHeight")
		initializerStatements += gongbasicfield.GongMarshallField(stage, "BespokeHeight")
	}
	return
}
func (gongenum *GongEnum) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += gongenum.GongMarshallField(stage, "Name")
		initializerStatements += gongenum.GongMarshallField(stage, "Type")
		pointersInitializesStatements += gongenum.GongMarshallField(stage, "GongEnumValues")
	}
	return
}
func (gongenumvalue *GongEnumValue) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += gongenumvalue.GongMarshallField(stage, "Name")
		initializerStatements += gongenumvalue.GongMarshallField(stage, "Value")
	}
	return
}
func (gonglink *GongLink) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += gonglink.GongMarshallField(stage, "Name")
		initializerStatements += gonglink.GongMarshallField(stage, "Recv")
		initializerStatements += gonglink.GongMarshallField(stage, "ImportPath")
	}
	return
}
func (gongnote *GongNote) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += gongnote.GongMarshallField(stage, "Name")
		initializerStatements += gongnote.GongMarshallField(stage, "Body")
		initializerStatements += gongnote.GongMarshallField(stage, "BodyHTML")
		pointersInitializesStatements += gongnote.GongMarshallField(stage, "Links")
	}
	return
}
func (gongstruct *GongStruct) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += gongstruct.GongMarshallField(stage, "Name")
		pointersInitializesStatements += gongstruct.GongMarshallField(stage, "GongBasicFields")
		pointersInitializesStatements += gongstruct.GongMarshallField(stage, "GongTimeFields")
		pointersInitializesStatements += gongstruct.GongMarshallField(stage, "PointerToGongStructFields")
		pointersInitializesStatements += gongstruct.GongMarshallField(stage, "SliceOfPointerToGongStructFields")
		initializerStatements += gongstruct.GongMarshallField(stage, "HasOnAfterUpdateSignature")
		initializerStatements += gongstruct.GongMarshallField(stage, "IsIgnoredForFront")
	}
	return
}
func (gongtimefield *GongTimeField) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += gongtimefield.GongMarshallField(stage, "Name")
		initializerStatements += gongtimefield.GongMarshallField(stage, "Index")
		initializerStatements += gongtimefield.GongMarshallField(stage, "CompositeStructName")
		initializerStatements += gongtimefield.GongMarshallField(stage, "BespokeTimeFormat")
	}
	return
}
func (metareference *MetaReference) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += metareference.GongMarshallField(stage, "Name")
	}
	return
}
func (modelpkg *ModelPkg) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += modelpkg.GongMarshallField(stage, "Name")
		initializerStatements += modelpkg.GongMarshallField(stage, "PkgPath")
		initializerStatements += modelpkg.GongMarshallField(stage, "PathToGoSubDirectory")
		initializerStatements += modelpkg.GongMarshallField(stage, "OrmPkgGenPath")
		initializerStatements += modelpkg.GongMarshallField(stage, "DbOrmPkgGenPath")
		initializerStatements += modelpkg.GongMarshallField(stage, "DbLiteOrmPkgGenPath")
		initializerStatements += modelpkg.GongMarshallField(stage, "DbPkgGenPath")
		initializerStatements += modelpkg.GongMarshallField(stage, "ControllersPkgGenPath")
		initializerStatements += modelpkg.GongMarshallField(stage, "FullstackPkgGenPath")
		initializerStatements += modelpkg.GongMarshallField(stage, "StackPkgGenPath")
		initializerStatements += modelpkg.GongMarshallField(stage, "Level1StackPkgGenPath")
		initializerStatements += modelpkg.GongMarshallField(stage, "StaticPkgGenPath")
		initializerStatements += modelpkg.GongMarshallField(stage, "ProbePkgGenPath")
		initializerStatements += modelpkg.GongMarshallField(stage, "NgWorkspacePath")
		initializerStatements += modelpkg.GongMarshallField(stage, "NgWorkspaceName")
		initializerStatements += modelpkg.GongMarshallField(stage, "NgDataLibrarySourceCodeDirectory")
		initializerStatements += modelpkg.GongMarshallField(stage, "NgSpecificLibrarySourceCodeDirectory")
		initializerStatements += modelpkg.GongMarshallField(stage, "MaterialLibDatamodelTargetPath")
	}
	return
}
func (pointertogongstructfield *PointerToGongStructField) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += pointertogongstructfield.GongMarshallField(stage, "Name")
		pointersInitializesStatements += pointertogongstructfield.GongMarshallField(stage, "GongStruct")
		initializerStatements += pointertogongstructfield.GongMarshallField(stage, "Index")
		initializerStatements += pointertogongstructfield.GongMarshallField(stage, "CompositeStructName")
		initializerStatements += pointertogongstructfield.GongMarshallField(stage, "IsType")
	}
	return
}
func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	initializerStatements += "\n"
	pointersInitializesStatements += "\n"
	{ // Insertion point for basic fields value assignment
		initializerStatements += sliceofpointertogongstructfield.GongMarshallField(stage, "Name")
		pointersInitializesStatements += sliceofpointertogongstructfield.GongMarshallField(stage, "GongStruct")
		initializerStatements += sliceofpointertogongstructfield.GongMarshallField(stage, "Index")
		initializerStatements += sliceofpointertogongstructfield.GongMarshallField(stage, "CompositeStructName")
	}
	return
}