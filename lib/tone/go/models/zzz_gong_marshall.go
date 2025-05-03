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
	"time"
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
}`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{}).Stage(stage)`

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
func (stage *Stage) Marshall(file *os.File, modelsPackageName, packageName string) {

	name := file.Name()

	if !strings.HasSuffix(name, ".go") {
		log.Fatalln(name + " is not a go filename")
	}

	log.Printf("%s Marshalling %s", time.Now().Format("2006-01-02 15:04:05.000000"), name)
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
	map_Freqency_Identifiers := make(map[*Freqency]string)
	_ = map_Freqency_Identifiers

	freqencyOrdered := []*Freqency{}
	for freqency := range stage.Freqencys {
		freqencyOrdered = append(freqencyOrdered, freqency)
	}
	sort.Slice(freqencyOrdered[:], func(i, j int) bool {
		freqencyi := freqencyOrdered[i]
		freqencyj := freqencyOrdered[j]
		freqencyi_order, oki := stage.FreqencyMap_Staged_Order[freqencyi]
		freqencyj_order, okj := stage.FreqencyMap_Staged_Order[freqencyj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return freqencyi_order < freqencyj_order
	})
	if len(freqencyOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, freqency := range freqencyOrdered {

		id = generatesIdentifier("Freqency", idx, freqency.Name)
		map_Freqency_Identifiers[freqency] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Freqency")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", freqency.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(freqency.Name))
		initializerStatements += setValueField

	}

	map_Note_Identifiers := make(map[*Note]string)
	_ = map_Note_Identifiers

	noteOrdered := []*Note{}
	for note := range stage.Notes {
		noteOrdered = append(noteOrdered, note)
	}
	sort.Slice(noteOrdered[:], func(i, j int) bool {
		notei := noteOrdered[i]
		notej := noteOrdered[j]
		notei_order, oki := stage.NoteMap_Staged_Order[notei]
		notej_order, okj := stage.NoteMap_Staged_Order[notej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return notei_order < notej_order
	})
	if len(noteOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, note := range noteOrdered {

		id = generatesIdentifier("Note", idx, note.Name)
		map_Note_Identifiers[note] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Note")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", note.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(note.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Start")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", note.Start))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Duration")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", note.Duration))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Velocity")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", note.Velocity))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Info")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(note.Info))
		initializerStatements += setValueField

	}

	map_Player_Identifiers := make(map[*Player]string)
	_ = map_Player_Identifiers

	playerOrdered := []*Player{}
	for player := range stage.Players {
		playerOrdered = append(playerOrdered, player)
	}
	sort.Slice(playerOrdered[:], func(i, j int) bool {
		playeri := playerOrdered[i]
		playerj := playerOrdered[j]
		playeri_order, oki := stage.PlayerMap_Staged_Order[playeri]
		playerj_order, okj := stage.PlayerMap_Staged_Order[playerj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return playeri_order < playerj_order
	})
	if len(playerOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, player := range playerOrdered {

		id = generatesIdentifier("Player", idx, player.Name)
		map_Player_Identifiers[player] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Player")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", player.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(player.Name))
		initializerStatements += setValueField

		if player.Status != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Status")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+player.Status.ToCodeString())
			initializerStatements += setValueField
		}

	}

	// insertion initialization of objects to stage
	if len(freqencyOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Freqency instances pointers"
	}
	for idx, freqency := range freqencyOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Freqency", idx, freqency.Name)
		map_Freqency_Identifiers[freqency] = id

		// Initialisation of values
	}

	if len(noteOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Note instances pointers"
	}
	for idx, note := range noteOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Note", idx, note.Name)
		map_Note_Identifiers[note] = id

		// Initialisation of values
		for _, _freqency := range note.Frequencies {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Frequencies")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Freqency_Identifiers[_freqency])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(playerOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Player instances pointers"
	}
	for idx, player := range playerOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Player", idx, player.Name)
		map_Player_Identifiers[player] = id

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
