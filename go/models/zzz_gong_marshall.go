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

	// Declaration of instances to stage{{Identifiers}}

	// Setup of values{{ValueInitializers}}

	// Setup of pointers{{PointersInitializers}}
}
`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: ` + "`" + `{{GeneratedFieldNameValue}}` + "`" + `}).Stage(stage)`

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
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongbasicfield.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BasicKindName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongbasicfield.BasicKindName))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DeclaredType")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongbasicfield.DeclaredType))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CompositeStructName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongbasicfield.CompositeStructName))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Index")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", gongbasicfield.Index))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsTextArea")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongbasicfield.IsTextArea))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsBespokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongbasicfield.IsBespokeWidth))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BespokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", gongbasicfield.BespokeWidth))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsBespokeHeight")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongbasicfield.IsBespokeHeight))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BespokeHeight")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", gongbasicfield.BespokeHeight))
		initializerStatements += setValueField

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
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongenum.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongenum.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongenum.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Type")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+gongenum.Type.ToCodeString())
		initializerStatements += setValueField

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
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongenumvalue.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongenumvalue.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongenumvalue.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongenumvalue.Value))
		initializerStatements += setValueField

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
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gonglink.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gonglink.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gonglink.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Recv")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gonglink.Recv))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gonglink.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ImportPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gonglink.ImportPath))
		initializerStatements += setValueField

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
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnote.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongnote.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnote.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Body")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongnote.Body))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnote.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BodyHTML")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongnote.BodyHTML))
		initializerStatements += setValueField

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
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongstruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongstruct.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongstruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasOnAfterUpdateSignature")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongstruct.HasOnAfterUpdateSignature))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongstruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsIgnoredForFront")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongstruct.IsIgnoredForFront))
		initializerStatements += setValueField

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
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongtimefield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongtimefield.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongtimefield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Index")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", gongtimefield.Index))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongtimefield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CompositeStructName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongtimefield.CompositeStructName))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongtimefield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BespokeTimeFormat")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongtimefield.BespokeTimeFormat))
		initializerStatements += setValueField

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
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", metareference.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(metareference.Name))
		initializerStatements += setValueField

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
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "PkgPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.PkgPath))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "PathToGoSubDirectory")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.PathToGoSubDirectory))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OrmPkgGenPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.OrmPkgGenPath))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DbOrmPkgGenPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.DbOrmPkgGenPath))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DbLiteOrmPkgGenPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.DbLiteOrmPkgGenPath))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DbPkgGenPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.DbPkgGenPath))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ControllersPkgGenPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.ControllersPkgGenPath))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FullstackPkgGenPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.FullstackPkgGenPath))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackPkgGenPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.StackPkgGenPath))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Level1StackPkgGenPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.Level1StackPkgGenPath))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StaticPkgGenPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.StaticPkgGenPath))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ProbePkgGenPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.ProbePkgGenPath))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NgWorkspacePath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.NgWorkspacePath))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NgWorkspaceName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.NgWorkspaceName))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NgDataLibrarySourceCodeDirectory")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.NgDataLibrarySourceCodeDirectory))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NgSpecificLibrarySourceCodeDirectory")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.NgSpecificLibrarySourceCodeDirectory))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MaterialLibDatamodelTargetPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.MaterialLibDatamodelTargetPath))
		initializerStatements += setValueField

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
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", pointertogongstructfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(pointertogongstructfield.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", pointertogongstructfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Index")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", pointertogongstructfield.Index))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", pointertogongstructfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CompositeStructName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(pointertogongstructfield.CompositeStructName))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", pointertogongstructfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsType")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", pointertogongstructfield.IsType))
		initializerStatements += setValueField

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
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", sliceofpointertogongstructfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(sliceofpointertogongstructfield.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", sliceofpointertogongstructfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Index")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", sliceofpointertogongstructfield.Index))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", sliceofpointertogongstructfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CompositeStructName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(sliceofpointertogongstructfield.CompositeStructName))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	if len(gongbasicfieldOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of GongBasicField instances pointers"
	}
	for _, gongbasicfield := range gongbasicfieldOrdered {
		_ = gongbasicfield
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if gongbasicfield.GongEnum != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GongEnum")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", gongbasicfield.GongEnum.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(gongenumOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of GongEnum instances pointers"
	}
	for _, gongenum := range gongenumOrdered {
		_ = gongenum
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		for _, _gongenumvalue := range gongenum.GongEnumValues {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", gongenum.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GongEnumValues")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _gongenumvalue.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(gongenumvalueOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of GongEnumValue instances pointers"
	}
	for _, gongenumvalue := range gongenumvalueOrdered {
		_ = gongenumvalue
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(gonglinkOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of GongLink instances pointers"
	}
	for _, gonglink := range gonglinkOrdered {
		_ = gonglink
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(gongnoteOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of GongNote instances pointers"
	}
	for _, gongnote := range gongnoteOrdered {
		_ = gongnote
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		for _, _gonglink := range gongnote.Links {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", gongnote.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Links")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _gonglink.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(gongstructOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of GongStruct instances pointers"
	}
	for _, gongstruct := range gongstructOrdered {
		_ = gongstruct
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		for _, _gongbasicfield := range gongstruct.GongBasicFields {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", gongstruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GongBasicFields")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _gongbasicfield.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _gongtimefield := range gongstruct.GongTimeFields {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", gongstruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GongTimeFields")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _gongtimefield.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _pointertogongstructfield := range gongstruct.PointerToGongStructFields {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", gongstruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "PointerToGongStructFields")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _pointertogongstructfield.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _sliceofpointertogongstructfield := range gongstruct.SliceOfPointerToGongStructFields {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", gongstruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SliceOfPointerToGongStructFields")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _sliceofpointertogongstructfield.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(gongtimefieldOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of GongTimeField instances pointers"
	}
	for _, gongtimefield := range gongtimefieldOrdered {
		_ = gongtimefield
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(metareferenceOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of MetaReference instances pointers"
	}
	for _, metareference := range metareferenceOrdered {
		_ = metareference
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(modelpkgOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ModelPkg instances pointers"
	}
	for _, modelpkg := range modelpkgOrdered {
		_ = modelpkg
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(pointertogongstructfieldOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of PointerToGongStructField instances pointers"
	}
	for _, pointertogongstructfield := range pointertogongstructfieldOrdered {
		_ = pointertogongstructfield
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if pointertogongstructfield.GongStruct != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", pointertogongstructfield.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GongStruct")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", pointertogongstructfield.GongStruct.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(sliceofpointertogongstructfieldOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of SliceOfPointerToGongStructField instances pointers"
	}
	for _, sliceofpointertogongstructfield := range sliceofpointertogongstructfieldOrdered {
		_ = sliceofpointertogongstructfield
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if sliceofpointertogongstructfield.GongStruct != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", sliceofpointertogongstructfield.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GongStruct")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", sliceofpointertogongstructfield.GongStruct.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

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

// insertion initialization of objects to stage
func (gongbasicfield *GongBasicField) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongbasicfield.Name))
		initializerStatements += setValueField
	case "BasicKindName":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BasicKindName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongbasicfield.BasicKindName))
		initializerStatements += setValueField
	case "DeclaredType":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DeclaredType")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongbasicfield.DeclaredType))
		initializerStatements += setValueField
	case "CompositeStructName":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CompositeStructName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongbasicfield.CompositeStructName))
		initializerStatements += setValueField
	case "Index":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Index")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", gongbasicfield.Index))
		initializerStatements += setValueField
	case "IsTextArea":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsTextArea")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongbasicfield.IsTextArea))
		initializerStatements += setValueField
	case "IsBespokeWidth":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsBespokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongbasicfield.IsBespokeWidth))
		initializerStatements += setValueField
	case "BespokeWidth":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BespokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", gongbasicfield.BespokeWidth))
		initializerStatements += setValueField
	case "IsBespokeHeight":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsBespokeHeight")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongbasicfield.IsBespokeHeight))
		initializerStatements += setValueField
	case "BespokeHeight":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BespokeHeight")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", gongbasicfield.BespokeHeight))
		initializerStatements += setValueField

	case "GongEnum":
		if gongbasicfield.GongEnum != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GongEnum")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", gongbasicfield.GongEnum.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct GongBasicField", fieldName)
	}
	return
}

func (gongenum *GongEnum) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongenum.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongenum.Name))
		initializerStatements += setValueField
	case "Type":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongenum.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Type")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+gongenum.Type.ToCodeString())
		initializerStatements += setValueField

	case "GongEnumValues":
		for _, _gongenumvalue := range gongenum.GongEnumValues {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", gongenum.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GongEnumValues")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _gongenumvalue.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct GongEnum", fieldName)
	}
	return
}

func (gongenumvalue *GongEnumValue) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongenumvalue.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongenumvalue.Name))
		initializerStatements += setValueField
	case "Value":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongenumvalue.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongenumvalue.Value))
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct GongEnumValue", fieldName)
	}
	return
}

func (gonglink *GongLink) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gonglink.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gonglink.Name))
		initializerStatements += setValueField
	case "Recv":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gonglink.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Recv")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gonglink.Recv))
		initializerStatements += setValueField
	case "ImportPath":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gonglink.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ImportPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gonglink.ImportPath))
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct GongLink", fieldName)
	}
	return
}

func (gongnote *GongNote) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnote.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongnote.Name))
		initializerStatements += setValueField
	case "Body":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnote.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Body")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongnote.Body))
		initializerStatements += setValueField
	case "BodyHTML":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongnote.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BodyHTML")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongnote.BodyHTML))
		initializerStatements += setValueField

	case "Links":
		for _, _gonglink := range gongnote.Links {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", gongnote.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Links")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _gonglink.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct GongNote", fieldName)
	}
	return
}

func (gongstruct *GongStruct) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongstruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongstruct.Name))
		initializerStatements += setValueField
	case "HasOnAfterUpdateSignature":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongstruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasOnAfterUpdateSignature")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongstruct.HasOnAfterUpdateSignature))
		initializerStatements += setValueField
	case "IsIgnoredForFront":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongstruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsIgnoredForFront")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", gongstruct.IsIgnoredForFront))
		initializerStatements += setValueField

	case "GongBasicFields":
		for _, _gongbasicfield := range gongstruct.GongBasicFields {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", gongstruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GongBasicFields")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _gongbasicfield.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "GongTimeFields":
		for _, _gongtimefield := range gongstruct.GongTimeFields {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", gongstruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GongTimeFields")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _gongtimefield.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "PointerToGongStructFields":
		for _, _pointertogongstructfield := range gongstruct.PointerToGongStructFields {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", gongstruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "PointerToGongStructFields")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _pointertogongstructfield.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "SliceOfPointerToGongStructFields":
		for _, _sliceofpointertogongstructfield := range gongstruct.SliceOfPointerToGongStructFields {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", gongstruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SliceOfPointerToGongStructFields")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _sliceofpointertogongstructfield.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct GongStruct", fieldName)
	}
	return
}

func (gongtimefield *GongTimeField) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongtimefield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongtimefield.Name))
		initializerStatements += setValueField
	case "Index":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongtimefield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Index")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", gongtimefield.Index))
		initializerStatements += setValueField
	case "CompositeStructName":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongtimefield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CompositeStructName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongtimefield.CompositeStructName))
		initializerStatements += setValueField
	case "BespokeTimeFormat":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gongtimefield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BespokeTimeFormat")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gongtimefield.BespokeTimeFormat))
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct GongTimeField", fieldName)
	}
	return
}

func (metareference *MetaReference) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", metareference.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(metareference.Name))
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct MetaReference", fieldName)
	}
	return
}

func (modelpkg *ModelPkg) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.Name))
		initializerStatements += setValueField
	case "PkgPath":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "PkgPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.PkgPath))
		initializerStatements += setValueField
	case "PathToGoSubDirectory":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "PathToGoSubDirectory")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.PathToGoSubDirectory))
		initializerStatements += setValueField
	case "OrmPkgGenPath":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OrmPkgGenPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.OrmPkgGenPath))
		initializerStatements += setValueField
	case "DbOrmPkgGenPath":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DbOrmPkgGenPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.DbOrmPkgGenPath))
		initializerStatements += setValueField
	case "DbLiteOrmPkgGenPath":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DbLiteOrmPkgGenPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.DbLiteOrmPkgGenPath))
		initializerStatements += setValueField
	case "DbPkgGenPath":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DbPkgGenPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.DbPkgGenPath))
		initializerStatements += setValueField
	case "ControllersPkgGenPath":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ControllersPkgGenPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.ControllersPkgGenPath))
		initializerStatements += setValueField
	case "FullstackPkgGenPath":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FullstackPkgGenPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.FullstackPkgGenPath))
		initializerStatements += setValueField
	case "StackPkgGenPath":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StackPkgGenPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.StackPkgGenPath))
		initializerStatements += setValueField
	case "Level1StackPkgGenPath":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Level1StackPkgGenPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.Level1StackPkgGenPath))
		initializerStatements += setValueField
	case "StaticPkgGenPath":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StaticPkgGenPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.StaticPkgGenPath))
		initializerStatements += setValueField
	case "ProbePkgGenPath":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ProbePkgGenPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.ProbePkgGenPath))
		initializerStatements += setValueField
	case "NgWorkspacePath":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NgWorkspacePath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.NgWorkspacePath))
		initializerStatements += setValueField
	case "NgWorkspaceName":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NgWorkspaceName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.NgWorkspaceName))
		initializerStatements += setValueField
	case "NgDataLibrarySourceCodeDirectory":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NgDataLibrarySourceCodeDirectory")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.NgDataLibrarySourceCodeDirectory))
		initializerStatements += setValueField
	case "NgSpecificLibrarySourceCodeDirectory":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NgSpecificLibrarySourceCodeDirectory")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.NgSpecificLibrarySourceCodeDirectory))
		initializerStatements += setValueField
	case "MaterialLibDatamodelTargetPath":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MaterialLibDatamodelTargetPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(modelpkg.MaterialLibDatamodelTargetPath))
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct ModelPkg", fieldName)
	}
	return
}

func (pointertogongstructfield *PointerToGongStructField) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", pointertogongstructfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(pointertogongstructfield.Name))
		initializerStatements += setValueField
	case "Index":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", pointertogongstructfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Index")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", pointertogongstructfield.Index))
		initializerStatements += setValueField
	case "CompositeStructName":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", pointertogongstructfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CompositeStructName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(pointertogongstructfield.CompositeStructName))
		initializerStatements += setValueField
	case "IsType":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", pointertogongstructfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsType")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", pointertogongstructfield.IsType))
		initializerStatements += setValueField

	case "GongStruct":
		if pointertogongstructfield.GongStruct != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", pointertogongstructfield.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GongStruct")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", pointertogongstructfield.GongStruct.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct PointerToGongStructField", fieldName)
	}
	return
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", sliceofpointertogongstructfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(sliceofpointertogongstructfield.Name))
		initializerStatements += setValueField
	case "Index":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", sliceofpointertogongstructfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Index")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", sliceofpointertogongstructfield.Index))
		initializerStatements += setValueField
	case "CompositeStructName":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", sliceofpointertogongstructfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CompositeStructName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(sliceofpointertogongstructfield.CompositeStructName))
		initializerStatements += setValueField

	case "GongStruct":
		if sliceofpointertogongstructfield.GongStruct != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", sliceofpointertogongstructfield.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GongStruct")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", sliceofpointertogongstructfield.GongStruct.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct SliceOfPointerToGongStructField", fieldName)
	}
	return
}
