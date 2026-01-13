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

const GongIdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: ` + "`" + `{{GeneratedFieldNameValue}}` + "`" + `}).Stage(stage)`

const GongUnstageStmt = `
	{{Identifier}}.Unstage(stage)`

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
		// Insertion point for basic fields value assignment
		initializerStatements += astruct.GongMarshallField(stage, "Name")
		initializerStatements += astruct.GongMarshallField(stage, "Field")
		pointersInitializesStatements += astruct.GongMarshallField(stage, "Associationtob")
		pointersInitializesStatements += astruct.GongMarshallField(stage, "Anarrayofb")
		pointersInitializesStatements += astruct.GongMarshallField(stage, "Anotherassociationtob_2")
		initializerStatements += astruct.GongMarshallField(stage, "Date")
		initializerStatements += astruct.GongMarshallField(stage, "Date2")
		initializerStatements += astruct.GongMarshallField(stage, "Booleanfield")
		initializerStatements += astruct.GongMarshallField(stage, "Aenum")
		initializerStatements += astruct.GongMarshallField(stage, "Aenum_2")
		initializerStatements += astruct.GongMarshallField(stage, "Benum")
		initializerStatements += astruct.GongMarshallField(stage, "CEnum")
		initializerStatements += astruct.GongMarshallField(stage, "CName")
		initializerStatements += astruct.GongMarshallField(stage, "CFloatfield")
		pointersInitializesStatements += astruct.GongMarshallField(stage, "Bstruct")
		pointersInitializesStatements += astruct.GongMarshallField(stage, "Bstruct2")
		pointersInitializesStatements += astruct.GongMarshallField(stage, "Dstruct")
		pointersInitializesStatements += astruct.GongMarshallField(stage, "Dstruct2")
		pointersInitializesStatements += astruct.GongMarshallField(stage, "Dstruct3")
		pointersInitializesStatements += astruct.GongMarshallField(stage, "Dstruct4")
		pointersInitializesStatements += astruct.GongMarshallField(stage, "Dstruct4s")
		initializerStatements += astruct.GongMarshallField(stage, "Floatfield")
		initializerStatements += astruct.GongMarshallField(stage, "Intfield")
		initializerStatements += astruct.GongMarshallField(stage, "Anotherbooleanfield")
		initializerStatements += astruct.GongMarshallField(stage, "Duration1")
		pointersInitializesStatements += astruct.GongMarshallField(stage, "Anarrayofa")
		pointersInitializesStatements += astruct.GongMarshallField(stage, "Anotherarrayofb")
		pointersInitializesStatements += astruct.GongMarshallField(stage, "AnarrayofbUse")
		pointersInitializesStatements += astruct.GongMarshallField(stage, "Anarrayofb2Use")
		pointersInitializesStatements += astruct.GongMarshallField(stage, "AnAstruct")
		initializerStatements += astruct.GongMarshallField(stage, "TextFieldBespokeSize")
		initializerStatements += astruct.GongMarshallField(stage, "TextArea")
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
		// Insertion point for basic fields value assignment
		initializerStatements += astructbstruct2use.GongMarshallField(stage, "Name")
		pointersInitializesStatements += astructbstruct2use.GongMarshallField(stage, "Bstrcut2")
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
		// Insertion point for basic fields value assignment
		initializerStatements += astructbstructuse.GongMarshallField(stage, "Name")
		pointersInitializesStatements += astructbstructuse.GongMarshallField(stage, "Bstruct2")
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
		// Insertion point for basic fields value assignment
		initializerStatements += bstruct.GongMarshallField(stage, "Name")
		initializerStatements += bstruct.GongMarshallField(stage, "Floatfield")
		initializerStatements += bstruct.GongMarshallField(stage, "Floatfield2")
		initializerStatements += bstruct.GongMarshallField(stage, "Intfield")
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
		// Insertion point for basic fields value assignment
		initializerStatements += dstruct.GongMarshallField(stage, "Name")
		pointersInitializesStatements += dstruct.GongMarshallField(stage, "Anarrayofb")
		pointersInitializesStatements += dstruct.GongMarshallField(stage, "Gstruct")
		pointersInitializesStatements += dstruct.GongMarshallField(stage, "Gstructs")
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
		// Insertion point for basic fields value assignment
		initializerStatements += f0123456789012345678901234567890.GongMarshallField(stage, "Name")
		initializerStatements += f0123456789012345678901234567890.GongMarshallField(stage, "Date")
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
		// Insertion point for basic fields value assignment
		initializerStatements += gstruct.GongMarshallField(stage, "Name")
		initializerStatements += gstruct.GongMarshallField(stage, "Floatfield")
		initializerStatements += gstruct.GongMarshallField(stage, "Floatfield2")
		initializerStatements += gstruct.GongMarshallField(stage, "Intfield")
	}

	// insertion initialization of objects to stage
	for _, astruct := range astructOrdered {
		_ = astruct
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, astructbstruct2use := range astructbstruct2useOrdered {
		_ = astructbstruct2use
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, astructbstructuse := range astructbstructuseOrdered {
		_ = astructbstructuse
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, bstruct := range bstructOrdered {
		_ = bstruct
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, dstruct := range dstructOrdered {
		_ = dstruct
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, f0123456789012345678901234567890 := range f0123456789012345678901234567890Ordered {
		_ = f0123456789012345678901234567890
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, gstruct := range gstructOrdered {
		_ = gstruct
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
func (astruct *Astruct) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(astruct.Name))
	case "Field":
		if str, ok := astruct.Field.(string); ok {
			res = MetaFieldStructInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Field")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", str)
		}
	case "Date":
		res = TimeInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Date")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", astruct.Date.String())
	case "Date2":
		res = TimeInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Date2")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", astruct.Date2.String())
	case "Booleanfield":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Booleanfield")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", astruct.Booleanfield))
	case "Aenum":
		if astruct.Aenum.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Aenum")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+astruct.Aenum.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Aenum")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "Aenum_2":
		if astruct.Aenum_2.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Aenum_2")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+astruct.Aenum_2.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Aenum_2")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "Benum":
		if astruct.Benum.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Benum")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+astruct.Benum.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Benum")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "CEnum":
		if astruct.CEnum.ToCodeString() != "" {
			res = NumberInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CEnum")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+astruct.CEnum.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = NumberInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CEnum")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "0")
		}
	case "CName":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(astruct.CName))
	case "CFloatfield":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CFloatfield")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", astruct.CFloatfield))
	case "Floatfield":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Floatfield")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", astruct.Floatfield))
	case "Intfield":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Intfield")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", astruct.Intfield))
	case "Anotherbooleanfield":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Anotherbooleanfield")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", astruct.Anotherbooleanfield))
	case "Duration1":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Duration1")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", astruct.Duration1))
	case "TextFieldBespokeSize":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TextFieldBespokeSize")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(astruct.TextFieldBespokeSize))
	case "TextArea":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TextArea")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(astruct.TextArea))

	case "Associationtob":
		if astruct.Associationtob != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Associationtob")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", astruct.Associationtob.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Associationtob")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Anarrayofb":
		for _, _bstruct := range astruct.Anarrayofb {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Anarrayofb")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _bstruct.GongGetIdentifier(stage))
			res += tmp
		}
	case "Anotherassociationtob_2":
		if astruct.Anotherassociationtob_2 != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Anotherassociationtob_2")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", astruct.Anotherassociationtob_2.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Anotherassociationtob_2")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Bstruct":
		if astruct.Bstruct != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Bstruct")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", astruct.Bstruct.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Bstruct")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Bstruct2":
		if astruct.Bstruct2 != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Bstruct2")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", astruct.Bstruct2.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Bstruct2")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Dstruct":
		if astruct.Dstruct != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Dstruct")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", astruct.Dstruct.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Dstruct")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Dstruct2":
		if astruct.Dstruct2 != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Dstruct2")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", astruct.Dstruct2.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Dstruct2")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Dstruct3":
		if astruct.Dstruct3 != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Dstruct3")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", astruct.Dstruct3.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Dstruct3")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Dstruct4":
		if astruct.Dstruct4 != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Dstruct4")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", astruct.Dstruct4.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Dstruct4")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Dstruct4s":
		for _, _dstruct := range astruct.Dstruct4s {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Dstruct4s")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _dstruct.GongGetIdentifier(stage))
			res += tmp
		}
	case "Anarrayofa":
		for _, _astruct := range astruct.Anarrayofa {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Anarrayofa")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _astruct.GongGetIdentifier(stage))
			res += tmp
		}
	case "Anotherarrayofb":
		for _, _bstruct := range astruct.Anotherarrayofb {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Anotherarrayofb")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _bstruct.GongGetIdentifier(stage))
			res += tmp
		}
	case "AnarrayofbUse":
		for _, _astructbstructuse := range astruct.AnarrayofbUse {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "AnarrayofbUse")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _astructbstructuse.GongGetIdentifier(stage))
			res += tmp
		}
	case "Anarrayofb2Use":
		for _, _astructbstruct2use := range astruct.Anarrayofb2Use {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Anarrayofb2Use")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _astructbstruct2use.GongGetIdentifier(stage))
			res += tmp
		}
	case "AnAstruct":
		if astruct.AnAstruct != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "AnAstruct")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", astruct.AnAstruct.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astruct.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "AnAstruct")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Astruct", fieldName)
	}
	return
}

func (astructbstruct2use *AstructBstruct2Use) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", astructbstruct2use.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(astructbstruct2use.Name))

	case "Bstrcut2":
		if astructbstruct2use.Bstrcut2 != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astructbstruct2use.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Bstrcut2")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", astructbstruct2use.Bstrcut2.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astructbstruct2use.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Bstrcut2")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AstructBstruct2Use", fieldName)
	}
	return
}

func (astructbstructuse *AstructBstructUse) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", astructbstructuse.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(astructbstructuse.Name))

	case "Bstruct2":
		if astructbstructuse.Bstruct2 != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astructbstructuse.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Bstruct2")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", astructbstructuse.Bstruct2.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", astructbstructuse.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Bstruct2")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AstructBstructUse", fieldName)
	}
	return
}

func (bstruct *Bstruct) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bstruct.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(bstruct.Name))
	case "Floatfield":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bstruct.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Floatfield")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bstruct.Floatfield))
	case "Floatfield2":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bstruct.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Floatfield2")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", bstruct.Floatfield2))
	case "Intfield":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", bstruct.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Intfield")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", bstruct.Intfield))

	default:
		log.Panicf("Unknown field %s for Gongstruct Bstruct", fieldName)
	}
	return
}

func (dstruct *Dstruct) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", dstruct.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(dstruct.Name))

	case "Anarrayofb":
		for _, _bstruct := range dstruct.Anarrayofb {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", dstruct.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Anarrayofb")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _bstruct.GongGetIdentifier(stage))
			res += tmp
		}
	case "Gstruct":
		if dstruct.Gstruct != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", dstruct.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Gstruct")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", dstruct.Gstruct.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", dstruct.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Gstruct")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Gstructs":
		for _, _gstruct := range dstruct.Gstructs {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", dstruct.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Gstructs")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _gstruct.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Dstruct", fieldName)
	}
	return
}

func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", f0123456789012345678901234567890.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(f0123456789012345678901234567890.Name))
	case "Date":
		res = TimeInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", f0123456789012345678901234567890.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Date")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", f0123456789012345678901234567890.Date.String())

	default:
		log.Panicf("Unknown field %s for Gongstruct F0123456789012345678901234567890", fieldName)
	}
	return
}

func (gstruct *Gstruct) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gstruct.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(gstruct.Name))
	case "Floatfield":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gstruct.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Floatfield")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gstruct.Floatfield))
	case "Floatfield2":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gstruct.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Floatfield2")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", gstruct.Floatfield2))
	case "Intfield":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", gstruct.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Intfield")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", gstruct.Intfield))

	default:
		log.Panicf("Unknown field %s for Gongstruct Gstruct", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (astruct *Astruct) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += astruct.GongMarshallField(stage, "Name")
		initializerStatements += astruct.GongMarshallField(stage, "Field")
		pointersInitializesStatements += astruct.GongMarshallField(stage, "Associationtob")
		pointersInitializesStatements += astruct.GongMarshallField(stage, "Anarrayofb")
		pointersInitializesStatements += astruct.GongMarshallField(stage, "Anotherassociationtob_2")
		initializerStatements += astruct.GongMarshallField(stage, "Date")
		initializerStatements += astruct.GongMarshallField(stage, "Date2")
		initializerStatements += astruct.GongMarshallField(stage, "Booleanfield")
		initializerStatements += astruct.GongMarshallField(stage, "Aenum")
		initializerStatements += astruct.GongMarshallField(stage, "Aenum_2")
		initializerStatements += astruct.GongMarshallField(stage, "Benum")
		initializerStatements += astruct.GongMarshallField(stage, "CEnum")
		initializerStatements += astruct.GongMarshallField(stage, "CName")
		initializerStatements += astruct.GongMarshallField(stage, "CFloatfield")
		pointersInitializesStatements += astruct.GongMarshallField(stage, "Bstruct")
		pointersInitializesStatements += astruct.GongMarshallField(stage, "Bstruct2")
		pointersInitializesStatements += astruct.GongMarshallField(stage, "Dstruct")
		pointersInitializesStatements += astruct.GongMarshallField(stage, "Dstruct2")
		pointersInitializesStatements += astruct.GongMarshallField(stage, "Dstruct3")
		pointersInitializesStatements += astruct.GongMarshallField(stage, "Dstruct4")
		pointersInitializesStatements += astruct.GongMarshallField(stage, "Dstruct4s")
		initializerStatements += astruct.GongMarshallField(stage, "Floatfield")
		initializerStatements += astruct.GongMarshallField(stage, "Intfield")
		initializerStatements += astruct.GongMarshallField(stage, "Anotherbooleanfield")
		initializerStatements += astruct.GongMarshallField(stage, "Duration1")
		pointersInitializesStatements += astruct.GongMarshallField(stage, "Anarrayofa")
		pointersInitializesStatements += astruct.GongMarshallField(stage, "Anotherarrayofb")
		pointersInitializesStatements += astruct.GongMarshallField(stage, "AnarrayofbUse")
		pointersInitializesStatements += astruct.GongMarshallField(stage, "Anarrayofb2Use")
		pointersInitializesStatements += astruct.GongMarshallField(stage, "AnAstruct")
		initializerStatements += astruct.GongMarshallField(stage, "TextFieldBespokeSize")
		initializerStatements += astruct.GongMarshallField(stage, "TextArea")
	}
	return
}
func (astructbstruct2use *AstructBstruct2Use) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += astructbstruct2use.GongMarshallField(stage, "Name")
		pointersInitializesStatements += astructbstruct2use.GongMarshallField(stage, "Bstrcut2")
	}
	return
}
func (astructbstructuse *AstructBstructUse) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += astructbstructuse.GongMarshallField(stage, "Name")
		pointersInitializesStatements += astructbstructuse.GongMarshallField(stage, "Bstruct2")
	}
	return
}
func (bstruct *Bstruct) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += bstruct.GongMarshallField(stage, "Name")
		initializerStatements += bstruct.GongMarshallField(stage, "Floatfield")
		initializerStatements += bstruct.GongMarshallField(stage, "Floatfield2")
		initializerStatements += bstruct.GongMarshallField(stage, "Intfield")
	}
	return
}
func (dstruct *Dstruct) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += dstruct.GongMarshallField(stage, "Name")
		pointersInitializesStatements += dstruct.GongMarshallField(stage, "Anarrayofb")
		pointersInitializesStatements += dstruct.GongMarshallField(stage, "Gstruct")
		pointersInitializesStatements += dstruct.GongMarshallField(stage, "Gstructs")
	}
	return
}
func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += f0123456789012345678901234567890.GongMarshallField(stage, "Name")
		initializerStatements += f0123456789012345678901234567890.GongMarshallField(stage, "Date")
	}
	return
}
func (gstruct *Gstruct) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += gstruct.GongMarshallField(stage, "Name")
		initializerStatements += gstruct.GongMarshallField(stage, "Floatfield")
		initializerStatements += gstruct.GongMarshallField(stage, "Floatfield2")
		initializerStatements += gstruct.GongMarshallField(stage, "Intfield")
	}
	return
}
