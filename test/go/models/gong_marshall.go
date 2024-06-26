// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
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

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
var _ = func(stage *models.StageStruct) {

	// Declaration of instances to stage{{Identifiers}}

	// Setup of values{{ValueInitializers}}

	// Setup of pointers{{PointersInitializers}}
}`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: ` + "`" + `{{GeneratedFieldNameValue}}` + "`" + `}).Stage(stage)`

const StringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + "`" + `{{GeneratedFieldNameValue}}` + "`"

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
func (stage *StageStruct) Marshall(file *os.File, modelsPackageName, packageName string) {

	name := file.Name()

	if !strings.HasSuffix(name, ".go") {
		log.Fatalln(name + " is not a go filename")
	}

	log.Println("filename of marshall output is " + name)
	newBase := filepath.Base(file.Name())

	res := marshallRes
	res = strings.ReplaceAll(res, "{{databaseName}}", strings.ReplaceAll(newBase, ".go", ""))
	res = strings.ReplaceAll(res, "{{PackageName}}", packageName)
	res = strings.ReplaceAll(res, "{{ModelsPackageName}}", modelsPackageName)

	// map of identifiers
	// var StageMapDstructIds map[*Dstruct]string
	identifiersDecl := ""
	initializerStatements := ""
	pointersInitializesStatements := ""

	id := ""
	_ = id
	decl := ""
	_ = decl
	setValueField := ""
	_ = setValueField

	// insertion initialization of objects to stage
	map_Astruct_Identifiers := make(map[*Astruct]string)
	_ = map_Astruct_Identifiers

	astructOrdered := []*Astruct{}
	for astruct := range stage.Astructs {
		astructOrdered = append(astructOrdered, astruct)
	}
	sort.Slice(astructOrdered[:], func(i, j int) bool {
		return astructOrdered[i].Name < astructOrdered[j].Name
	})
	if len(astructOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, astruct := range astructOrdered {

		id = generatesIdentifier("Astruct", idx, astruct.Name)
		map_Astruct_Identifiers[astruct] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Astruct")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", astruct.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(astruct.Name))
		initializerStatements += setValueField

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Date")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", astruct.Date.String())
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Booleanfield")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", astruct.Booleanfield))
		initializerStatements += setValueField

		if astruct.Aenum != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Aenum")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+astruct.Aenum.ToCodeString())
			initializerStatements += setValueField
		}

		if astruct.Aenum_2 != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Aenum_2")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+astruct.Aenum_2.ToCodeString())
			initializerStatements += setValueField
		}

		if astruct.Benum != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Benum")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+astruct.Benum.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CEnum")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+astruct.CEnum.ToCodeString())
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(astruct.CName))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CFloatfield")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", astruct.CFloatfield))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Floatfield")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", astruct.Floatfield))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Intfield")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", astruct.Intfield))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Anotherbooleanfield")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", astruct.Anotherbooleanfield))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Duration1")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", astruct.Duration1))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "\n\t{{Identifier}}",
			fmt.Sprintf("\n\n\t//gong:ident [%s] comment added to overcome the problem with the comment map association\n\t{{Identifier}}",
				string(astruct.StructRef)))
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StructRef")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(astruct.StructRef))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "\n\t{{Identifier}}",
			fmt.Sprintf("\n\n\t//gong:ident [%s] comment added to overcome the problem with the comment map association\n\t{{Identifier}}",
				string(astruct.FieldRef)))
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "FieldRef")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(astruct.FieldRef))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "\n\t{{Identifier}}",
			fmt.Sprintf("\n\n\t//gong:ident [%s] comment added to overcome the problem with the comment map association\n\t{{Identifier}}",
				string(astruct.EnumIntRef)))
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EnumIntRef")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(astruct.EnumIntRef))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "\n\t{{Identifier}}",
			fmt.Sprintf("\n\n\t//gong:ident [%s] comment added to overcome the problem with the comment map association\n\t{{Identifier}}",
				string(astruct.EnumStringRef)))
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EnumStringRef")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(astruct.EnumStringRef))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "\n\t{{Identifier}}",
			fmt.Sprintf("\n\n\t//gong:ident [%s] comment added to overcome the problem with the comment map association\n\t{{Identifier}}",
				string(astruct.EnumValue)))
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EnumValue")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(astruct.EnumValue))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "\n\t{{Identifier}}",
			fmt.Sprintf("\n\n\t//gong:ident [%s] comment added to overcome the problem with the comment map association\n\t{{Identifier}}",
				string(astruct.ConstIdentifierValue)))
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ConstIdentifierValue")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(astruct.ConstIdentifierValue))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TextFieldBespokeSize")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(astruct.TextFieldBespokeSize))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TextArea")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(astruct.TextArea))
		initializerStatements += setValueField

	}

	map_AstructBstruct2Use_Identifiers := make(map[*AstructBstruct2Use]string)
	_ = map_AstructBstruct2Use_Identifiers

	astructbstruct2useOrdered := []*AstructBstruct2Use{}
	for astructbstruct2use := range stage.AstructBstruct2Uses {
		astructbstruct2useOrdered = append(astructbstruct2useOrdered, astructbstruct2use)
	}
	sort.Slice(astructbstruct2useOrdered[:], func(i, j int) bool {
		return astructbstruct2useOrdered[i].Name < astructbstruct2useOrdered[j].Name
	})
	if len(astructbstruct2useOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, astructbstruct2use := range astructbstruct2useOrdered {

		id = generatesIdentifier("AstructBstruct2Use", idx, astructbstruct2use.Name)
		map_AstructBstruct2Use_Identifiers[astructbstruct2use] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AstructBstruct2Use")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", astructbstruct2use.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(astructbstruct2use.Name))
		initializerStatements += setValueField

	}

	map_AstructBstructUse_Identifiers := make(map[*AstructBstructUse]string)
	_ = map_AstructBstructUse_Identifiers

	astructbstructuseOrdered := []*AstructBstructUse{}
	for astructbstructuse := range stage.AstructBstructUses {
		astructbstructuseOrdered = append(astructbstructuseOrdered, astructbstructuse)
	}
	sort.Slice(astructbstructuseOrdered[:], func(i, j int) bool {
		return astructbstructuseOrdered[i].Name < astructbstructuseOrdered[j].Name
	})
	if len(astructbstructuseOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, astructbstructuse := range astructbstructuseOrdered {

		id = generatesIdentifier("AstructBstructUse", idx, astructbstructuse.Name)
		map_AstructBstructUse_Identifiers[astructbstructuse] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AstructBstructUse")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", astructbstructuse.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(astructbstructuse.Name))
		initializerStatements += setValueField

	}

	map_Bstruct_Identifiers := make(map[*Bstruct]string)
	_ = map_Bstruct_Identifiers

	bstructOrdered := []*Bstruct{}
	for bstruct := range stage.Bstructs {
		bstructOrdered = append(bstructOrdered, bstruct)
	}
	sort.Slice(bstructOrdered[:], func(i, j int) bool {
		return bstructOrdered[i].Name < bstructOrdered[j].Name
	})
	if len(bstructOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, bstruct := range bstructOrdered {

		id = generatesIdentifier("Bstruct", idx, bstruct.Name)
		map_Bstruct_Identifiers[bstruct] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Bstruct")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", bstruct.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bstruct.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Floatfield")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bstruct.Floatfield))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Floatfield2")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bstruct.Floatfield2))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Intfield")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", bstruct.Intfield))
		initializerStatements += setValueField

	}

	map_Dstruct_Identifiers := make(map[*Dstruct]string)
	_ = map_Dstruct_Identifiers

	dstructOrdered := []*Dstruct{}
	for dstruct := range stage.Dstructs {
		dstructOrdered = append(dstructOrdered, dstruct)
	}
	sort.Slice(dstructOrdered[:], func(i, j int) bool {
		return dstructOrdered[i].Name < dstructOrdered[j].Name
	})
	if len(dstructOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, dstruct := range dstructOrdered {

		id = generatesIdentifier("Dstruct", idx, dstruct.Name)
		map_Dstruct_Identifiers[dstruct] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Dstruct")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", dstruct.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(dstruct.Name))
		initializerStatements += setValueField

	}

	map_Fstruct_Identifiers := make(map[*Fstruct]string)
	_ = map_Fstruct_Identifiers

	fstructOrdered := []*Fstruct{}
	for fstruct := range stage.Fstructs {
		fstructOrdered = append(fstructOrdered, fstruct)
	}
	sort.Slice(fstructOrdered[:], func(i, j int) bool {
		return fstructOrdered[i].Name < fstructOrdered[j].Name
	})
	if len(fstructOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, fstruct := range fstructOrdered {

		id = generatesIdentifier("Fstruct", idx, fstruct.Name)
		map_Fstruct_Identifiers[fstruct] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Fstruct")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", fstruct.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(fstruct.Name))
		initializerStatements += setValueField

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Date")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fstruct.Date.String())
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	for idx, astruct := range astructOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Astruct", idx, astruct.Name)
		map_Astruct_Identifiers[astruct] = id

		// Initialisation of values
		if astruct.Associationtob != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Associationtob")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Bstruct_Identifiers[astruct.Associationtob])
			pointersInitializesStatements += setPointerField
		}

		for _, _bstruct := range astruct.Anarrayofb {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Anarrayofb")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Bstruct_Identifiers[_bstruct])
			pointersInitializesStatements += setPointerField
		}

		if astruct.Anotherassociationtob_2 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Anotherassociationtob_2")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Bstruct_Identifiers[astruct.Anotherassociationtob_2])
			pointersInitializesStatements += setPointerField
		}

		if astruct.Bstruct != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Bstruct")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Bstruct_Identifiers[astruct.Bstruct])
			pointersInitializesStatements += setPointerField
		}

		if astruct.Bstruct2 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Bstruct2")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Bstruct_Identifiers[astruct.Bstruct2])
			pointersInitializesStatements += setPointerField
		}

		if astruct.Dstruct != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Dstruct")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Dstruct_Identifiers[astruct.Dstruct])
			pointersInitializesStatements += setPointerField
		}

		if astruct.Dstruct2 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Dstruct2")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Dstruct_Identifiers[astruct.Dstruct2])
			pointersInitializesStatements += setPointerField
		}

		if astruct.Dstruct3 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Dstruct3")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Dstruct_Identifiers[astruct.Dstruct3])
			pointersInitializesStatements += setPointerField
		}

		if astruct.Dstruct4 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Dstruct4")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Dstruct_Identifiers[astruct.Dstruct4])
			pointersInitializesStatements += setPointerField
		}

		for _, _astruct := range astruct.Anarrayofa {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Anarrayofa")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Astruct_Identifiers[_astruct])
			pointersInitializesStatements += setPointerField
		}

		for _, _bstruct := range astruct.Anotherarrayofb {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Anotherarrayofb")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Bstruct_Identifiers[_bstruct])
			pointersInitializesStatements += setPointerField
		}

		for _, _astructbstructuse := range astruct.AnarrayofbUse {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "AnarrayofbUse")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_AstructBstructUse_Identifiers[_astructbstructuse])
			pointersInitializesStatements += setPointerField
		}

		for _, _astructbstruct2use := range astruct.Anarrayofb2Use {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Anarrayofb2Use")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_AstructBstruct2Use_Identifiers[_astructbstruct2use])
			pointersInitializesStatements += setPointerField
		}

		if astruct.AnAstruct != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "AnAstruct")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Astruct_Identifiers[astruct.AnAstruct])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, astructbstruct2use := range astructbstruct2useOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("AstructBstruct2Use", idx, astructbstruct2use.Name)
		map_AstructBstruct2Use_Identifiers[astructbstruct2use] = id

		// Initialisation of values
		if astructbstruct2use.Bstrcut2 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Bstrcut2")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Bstruct_Identifiers[astructbstruct2use.Bstrcut2])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, astructbstructuse := range astructbstructuseOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("AstructBstructUse", idx, astructbstructuse.Name)
		map_AstructBstructUse_Identifiers[astructbstructuse] = id

		// Initialisation of values
		if astructbstructuse.Bstruct2 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Bstruct2")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Bstruct_Identifiers[astructbstructuse.Bstruct2])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, bstruct := range bstructOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Bstruct", idx, bstruct.Name)
		map_Bstruct_Identifiers[bstruct] = id

		// Initialisation of values
	}

	for idx, dstruct := range dstructOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Dstruct", idx, dstruct.Name)
		map_Dstruct_Identifiers[dstruct] = id

		// Initialisation of values
		for _, _bstruct := range dstruct.Anarrayofb {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Anarrayofb")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Bstruct_Identifiers[_bstruct])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, fstruct := range fstructOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Fstruct", idx, fstruct.Name)
		map_Fstruct_Identifiers[fstruct] = id

		// Initialisation of values
	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl)
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements)
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements)

	if stage.MetaPackageImportAlias != "" {
		res = strings.ReplaceAll(res, "{{ImportPackageDeclaration}}",
			fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))

		res = strings.ReplaceAll(res, "{{ImportPackageDummyDeclaration}}",
			fmt.Sprintf("\nvar _ %s.StageStruct",
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

		res = strings.ReplaceAll(res, "{{EntriesDocLinkStringDocLinkIdentifier}}", entries)
	}

	fmt.Fprintln(file, res)
}

// unique identifier per struct
func generatesIdentifier(gongStructName string, idx int, instanceName string) (identifier string) {

	identifier = instanceName
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(instanceName, "_")

	identifier = fmt.Sprintf("__%s__%06d_%s", gongStructName, idx, processedString)

	return
}
