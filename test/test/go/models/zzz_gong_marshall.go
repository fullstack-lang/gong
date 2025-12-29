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
	astructOrdered := []*Astruct{}
	for astruct := range stage.Astructs {
		astructOrdered = append(astructOrdered, astruct)
	}
	sort.Slice(astructOrdered[:], func(i, j int) bool {
		astructi := astructOrdered[i]
		astructj := astructOrdered[j]
		astructi_order, oki := stage.AstructMap_Staged_Order[astructi]
		astructj_order, okj := stage.AstructMap_Staged_Order[astructj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return astructi_order < astructj_order
	})
	if len(astructOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, astruct := range astructOrdered {
	
		identifiersDecl += astruct.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(astruct.Name))
		initializerStatements += setValueField

		if str, ok := astruct.Field.(string); ok {
			setValueField = MetaFieldStructInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Field")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", str)
			initializerStatements += setValueField
		}

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Date")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", astruct.Date.String())
		initializerStatements += setValueField

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Date2")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", astruct.Date2.String())
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Booleanfield")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", astruct.Booleanfield))
		initializerStatements += setValueField

		if astruct.Aenum != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Aenum")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+astruct.Aenum.ToCodeString())
			initializerStatements += setValueField
		}

		if astruct.Aenum_2 != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Aenum_2")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+astruct.Aenum_2.ToCodeString())
			initializerStatements += setValueField
		}

		if astruct.Benum != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Benum")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+astruct.Benum.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CEnum")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+astruct.CEnum.ToCodeString())
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(astruct.CName))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CFloatfield")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", astruct.CFloatfield))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Floatfield")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", astruct.Floatfield))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Intfield")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", astruct.Intfield))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Anotherbooleanfield")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", astruct.Anotherbooleanfield))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Duration1")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", astruct.Duration1))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TextFieldBespokeSize")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(astruct.TextFieldBespokeSize))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TextArea")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(astruct.TextArea))
		initializerStatements += setValueField

	}

	astructbstruct2useOrdered := []*AstructBstruct2Use{}
	for astructbstruct2use := range stage.AstructBstruct2Uses {
		astructbstruct2useOrdered = append(astructbstruct2useOrdered, astructbstruct2use)
	}
	sort.Slice(astructbstruct2useOrdered[:], func(i, j int) bool {
		astructbstruct2usei := astructbstruct2useOrdered[i]
		astructbstruct2usej := astructbstruct2useOrdered[j]
		astructbstruct2usei_order, oki := stage.AstructBstruct2UseMap_Staged_Order[astructbstruct2usei]
		astructbstruct2usej_order, okj := stage.AstructBstruct2UseMap_Staged_Order[astructbstruct2usej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return astructbstruct2usei_order < astructbstruct2usej_order
	})
	if len(astructbstruct2useOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, astructbstruct2use := range astructbstruct2useOrdered {
	
		identifiersDecl += astructbstruct2use.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astructbstruct2use.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(astructbstruct2use.Name))
		initializerStatements += setValueField

	}

	astructbstructuseOrdered := []*AstructBstructUse{}
	for astructbstructuse := range stage.AstructBstructUses {
		astructbstructuseOrdered = append(astructbstructuseOrdered, astructbstructuse)
	}
	sort.Slice(astructbstructuseOrdered[:], func(i, j int) bool {
		astructbstructusei := astructbstructuseOrdered[i]
		astructbstructusej := astructbstructuseOrdered[j]
		astructbstructusei_order, oki := stage.AstructBstructUseMap_Staged_Order[astructbstructusei]
		astructbstructusej_order, okj := stage.AstructBstructUseMap_Staged_Order[astructbstructusej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return astructbstructusei_order < astructbstructusej_order
	})
	if len(astructbstructuseOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, astructbstructuse := range astructbstructuseOrdered {
	
		identifiersDecl += astructbstructuse.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astructbstructuse.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(astructbstructuse.Name))
		initializerStatements += setValueField

	}

	bstructOrdered := []*Bstruct{}
	for bstruct := range stage.Bstructs {
		bstructOrdered = append(bstructOrdered, bstruct)
	}
	sort.Slice(bstructOrdered[:], func(i, j int) bool {
		bstructi := bstructOrdered[i]
		bstructj := bstructOrdered[j]
		bstructi_order, oki := stage.BstructMap_Staged_Order[bstructi]
		bstructj_order, okj := stage.BstructMap_Staged_Order[bstructj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return bstructi_order < bstructj_order
	})
	if len(bstructOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, bstruct := range bstructOrdered {
	
		identifiersDecl += bstruct.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", bstruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bstruct.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", bstruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Floatfield")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bstruct.Floatfield))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", bstruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Floatfield2")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bstruct.Floatfield2))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", bstruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Intfield")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", bstruct.Intfield))
		initializerStatements += setValueField

	}

	dstructOrdered := []*Dstruct{}
	for dstruct := range stage.Dstructs {
		dstructOrdered = append(dstructOrdered, dstruct)
	}
	sort.Slice(dstructOrdered[:], func(i, j int) bool {
		dstructi := dstructOrdered[i]
		dstructj := dstructOrdered[j]
		dstructi_order, oki := stage.DstructMap_Staged_Order[dstructi]
		dstructj_order, okj := stage.DstructMap_Staged_Order[dstructj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return dstructi_order < dstructj_order
	})
	if len(dstructOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, dstruct := range dstructOrdered {
	
		identifiersDecl += dstruct.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", dstruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(dstruct.Name))
		initializerStatements += setValueField

	}

	f0123456789012345678901234567890Ordered := []*F0123456789012345678901234567890{}
	for f0123456789012345678901234567890 := range stage.F0123456789012345678901234567890s {
		f0123456789012345678901234567890Ordered = append(f0123456789012345678901234567890Ordered, f0123456789012345678901234567890)
	}
	sort.Slice(f0123456789012345678901234567890Ordered[:], func(i, j int) bool {
		f0123456789012345678901234567890i := f0123456789012345678901234567890Ordered[i]
		f0123456789012345678901234567890j := f0123456789012345678901234567890Ordered[j]
		f0123456789012345678901234567890i_order, oki := stage.F0123456789012345678901234567890Map_Staged_Order[f0123456789012345678901234567890i]
		f0123456789012345678901234567890j_order, okj := stage.F0123456789012345678901234567890Map_Staged_Order[f0123456789012345678901234567890j]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return f0123456789012345678901234567890i_order < f0123456789012345678901234567890j_order
	})
	if len(f0123456789012345678901234567890Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for _, f0123456789012345678901234567890 := range f0123456789012345678901234567890Ordered {
	
		identifiersDecl += f0123456789012345678901234567890.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", f0123456789012345678901234567890.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(f0123456789012345678901234567890.Name))
		initializerStatements += setValueField

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", f0123456789012345678901234567890.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Date")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", f0123456789012345678901234567890.Date.String())
		initializerStatements += setValueField

	}

	gstructOrdered := []*Gstruct{}
	for gstruct := range stage.Gstructs {
		gstructOrdered = append(gstructOrdered, gstruct)
	}
	sort.Slice(gstructOrdered[:], func(i, j int) bool {
		gstructi := gstructOrdered[i]
		gstructj := gstructOrdered[j]
		gstructi_order, oki := stage.GstructMap_Staged_Order[gstructi]
		gstructj_order, okj := stage.GstructMap_Staged_Order[gstructj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return gstructi_order < gstructj_order
	})
	if len(gstructOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, gstruct := range gstructOrdered {
	
		identifiersDecl += gstruct.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gstruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gstruct.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gstruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Floatfield")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gstruct.Floatfield))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gstruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Floatfield2")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gstruct.Floatfield2))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gstruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Intfield")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", gstruct.Intfield))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	if len(astructOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Astruct instances pointers"
	}
	for _, astruct := range astructOrdered {
		_ = astruct
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if astruct.Associationtob != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Associationtob")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", astruct.Associationtob.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _bstruct := range astruct.Anarrayofb {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Anarrayofb")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _bstruct.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if astruct.Anotherassociationtob_2 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Anotherassociationtob_2")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", astruct.Anotherassociationtob_2.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if astruct.Bstruct != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Bstruct")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", astruct.Bstruct.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if astruct.Bstruct2 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Bstruct2")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", astruct.Bstruct2.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if astruct.Dstruct != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Dstruct")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", astruct.Dstruct.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if astruct.Dstruct2 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Dstruct2")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", astruct.Dstruct2.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if astruct.Dstruct3 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Dstruct3")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", astruct.Dstruct3.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if astruct.Dstruct4 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Dstruct4")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", astruct.Dstruct4.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _dstruct := range astruct.Dstruct4s {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Dstruct4s")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _dstruct.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _astruct := range astruct.Anarrayofa {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Anarrayofa")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _astruct.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _bstruct := range astruct.Anotherarrayofb {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Anotherarrayofb")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _bstruct.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _astructbstructuse := range astruct.AnarrayofbUse {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "AnarrayofbUse")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _astructbstructuse.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _astructbstruct2use := range astruct.Anarrayofb2Use {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Anarrayofb2Use")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _astructbstruct2use.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if astruct.AnAstruct != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "AnAstruct")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", astruct.AnAstruct.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(astructbstruct2useOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of AstructBstruct2Use instances pointers"
	}
	for _, astructbstruct2use := range astructbstruct2useOrdered {
		_ = astructbstruct2use
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if astructbstruct2use.Bstrcut2 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astructbstruct2use.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Bstrcut2")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", astructbstruct2use.Bstrcut2.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(astructbstructuseOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of AstructBstructUse instances pointers"
	}
	for _, astructbstructuse := range astructbstructuseOrdered {
		_ = astructbstructuse
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if astructbstructuse.Bstruct2 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astructbstructuse.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Bstruct2")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", astructbstructuse.Bstruct2.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(bstructOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Bstruct instances pointers"
	}
	for _, bstruct := range bstructOrdered {
		_ = bstruct
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(dstructOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Dstruct instances pointers"
	}
	for _, dstruct := range dstructOrdered {
		_ = dstruct
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		for _, _bstruct := range dstruct.Anarrayofb {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", dstruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Anarrayofb")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _bstruct.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if dstruct.Gstruct != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", dstruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Gstruct")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", dstruct.Gstruct.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _gstruct := range dstruct.Gstructs {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", dstruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Gstructs")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _gstruct.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(f0123456789012345678901234567890Ordered) > 0 {
		pointersInitializesStatements += "\n\t// setup of F0123456789012345678901234567890 instances pointers"
	}
	for _, f0123456789012345678901234567890 := range f0123456789012345678901234567890Ordered {
		_ = f0123456789012345678901234567890
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(gstructOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Gstruct instances pointers"
	}
	for _, gstruct := range gstructOrdered {
		_ = gstruct
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
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
func (astruct *Astruct) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(astruct.Name))
		initializerStatements += setValueField
	case "Field":
		if str, ok := astruct.Field.(string); ok {
			setValueField = MetaFieldStructInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Field")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", str)
			initializerStatements += setValueField
		}
	case "Date":
		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Date")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", astruct.Date.String())
		initializerStatements += setValueField
	case "Date2":
		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Date2")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", astruct.Date2.String())
		initializerStatements += setValueField
	case "Booleanfield":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Booleanfield")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", astruct.Booleanfield))
		initializerStatements += setValueField
	case "Aenum":
		if astruct.Aenum != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Aenum")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+astruct.Aenum.ToCodeString())
			initializerStatements += setValueField
		}
	case "Aenum_2":
		if astruct.Aenum_2 != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Aenum_2")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+astruct.Aenum_2.ToCodeString())
			initializerStatements += setValueField
		}
	case "Benum":
		if astruct.Benum != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Benum")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+astruct.Benum.ToCodeString())
			initializerStatements += setValueField
		}
	case "CEnum":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CEnum")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+astruct.CEnum.ToCodeString())
		initializerStatements += setValueField
	case "CName":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(astruct.CName))
		initializerStatements += setValueField
	case "CFloatfield":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CFloatfield")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", astruct.CFloatfield))
		initializerStatements += setValueField
	case "Floatfield":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Floatfield")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", astruct.Floatfield))
		initializerStatements += setValueField
	case "Intfield":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Intfield")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", astruct.Intfield))
		initializerStatements += setValueField
	case "Anotherbooleanfield":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Anotherbooleanfield")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", astruct.Anotherbooleanfield))
		initializerStatements += setValueField
	case "Duration1":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Duration1")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", astruct.Duration1))
		initializerStatements += setValueField
	case "TextFieldBespokeSize":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TextFieldBespokeSize")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(astruct.TextFieldBespokeSize))
		initializerStatements += setValueField
	case "TextArea":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TextArea")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(astruct.TextArea))
		initializerStatements += setValueField

	case "Associationtob":
		if astruct.Associationtob != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Associationtob")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", astruct.Associationtob.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Anarrayofb":
		for _, _bstruct := range astruct.Anarrayofb {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Anarrayofb")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _bstruct.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Anotherassociationtob_2":
		if astruct.Anotherassociationtob_2 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Anotherassociationtob_2")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", astruct.Anotherassociationtob_2.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Bstruct":
		if astruct.Bstruct != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Bstruct")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", astruct.Bstruct.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Bstruct2":
		if astruct.Bstruct2 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Bstruct2")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", astruct.Bstruct2.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Dstruct":
		if astruct.Dstruct != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Dstruct")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", astruct.Dstruct.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Dstruct2":
		if astruct.Dstruct2 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Dstruct2")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", astruct.Dstruct2.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Dstruct3":
		if astruct.Dstruct3 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Dstruct3")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", astruct.Dstruct3.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Dstruct4":
		if astruct.Dstruct4 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Dstruct4")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", astruct.Dstruct4.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Dstruct4s":
		for _, _dstruct := range astruct.Dstruct4s {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Dstruct4s")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _dstruct.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Anarrayofa":
		for _, _astruct := range astruct.Anarrayofa {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Anarrayofa")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _astruct.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Anotherarrayofb":
		for _, _bstruct := range astruct.Anotherarrayofb {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Anotherarrayofb")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _bstruct.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "AnarrayofbUse":
		for _, _astructbstructuse := range astruct.AnarrayofbUse {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "AnarrayofbUse")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _astructbstructuse.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Anarrayofb2Use":
		for _, _astructbstruct2use := range astruct.Anarrayofb2Use {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Anarrayofb2Use")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _astructbstruct2use.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "AnAstruct":
		if astruct.AnAstruct != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "AnAstruct")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", astruct.AnAstruct.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct Astruct", fieldName)
	}
	return
}

func (astructbstruct2use *AstructBstruct2Use) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astructbstruct2use.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(astructbstruct2use.Name))
		initializerStatements += setValueField

	case "Bstrcut2":
		if astructbstruct2use.Bstrcut2 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astructbstruct2use.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Bstrcut2")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", astructbstruct2use.Bstrcut2.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct AstructBstruct2Use", fieldName)
	}
	return
}

func (astructbstructuse *AstructBstructUse) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", astructbstructuse.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(astructbstructuse.Name))
		initializerStatements += setValueField

	case "Bstruct2":
		if astructbstructuse.Bstruct2 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", astructbstructuse.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Bstruct2")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", astructbstructuse.Bstruct2.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct AstructBstructUse", fieldName)
	}
	return
}

func (bstruct *Bstruct) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", bstruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(bstruct.Name))
		initializerStatements += setValueField
	case "Floatfield":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", bstruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Floatfield")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bstruct.Floatfield))
		initializerStatements += setValueField
	case "Floatfield2":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", bstruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Floatfield2")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bstruct.Floatfield2))
		initializerStatements += setValueField
	case "Intfield":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", bstruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Intfield")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", bstruct.Intfield))
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct Bstruct", fieldName)
	}
	return
}

func (dstruct *Dstruct) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", dstruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(dstruct.Name))
		initializerStatements += setValueField

	case "Anarrayofb":
		for _, _bstruct := range dstruct.Anarrayofb {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", dstruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Anarrayofb")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _bstruct.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Gstruct":
		if dstruct.Gstruct != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", dstruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Gstruct")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", dstruct.Gstruct.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Gstructs":
		for _, _gstruct := range dstruct.Gstructs {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", dstruct.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Gstructs")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _gstruct.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct Dstruct", fieldName)
	}
	return
}

func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", f0123456789012345678901234567890.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(f0123456789012345678901234567890.Name))
		initializerStatements += setValueField
	case "Date":
		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", f0123456789012345678901234567890.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Date")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", f0123456789012345678901234567890.Date.String())
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct F0123456789012345678901234567890", fieldName)
	}
	return
}

func (gstruct *Gstruct) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gstruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(gstruct.Name))
		initializerStatements += setValueField
	case "Floatfield":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gstruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Floatfield")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gstruct.Floatfield))
		initializerStatements += setValueField
	case "Floatfield2":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gstruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Floatfield2")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gstruct.Floatfield2))
		initializerStatements += setValueField
	case "Intfield":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", gstruct.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Intfield")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", gstruct.Intfield))
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct Gstruct", fieldName)
	}
	return
}
