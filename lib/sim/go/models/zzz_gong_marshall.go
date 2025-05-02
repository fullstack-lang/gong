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
	map_Command_Identifiers := make(map[*Command]string)
	_ = map_Command_Identifiers

	commandOrdered := []*Command{}
	for command := range stage.Commands {
		commandOrdered = append(commandOrdered, command)
	}
	sort.Slice(commandOrdered[:], func(i, j int) bool {
		commandi := commandOrdered[i]
		commandj := commandOrdered[j]
		commandi_order, oki := stage.CommandMap_Staged_Order[commandi]
		commandj_order, okj := stage.CommandMap_Staged_Order[commandj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return commandi_order < commandj_order
	})
	if len(commandOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, command := range commandOrdered {

		id = generatesIdentifier("Command", idx, command.Name)
		map_Command_Identifiers[command] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Command")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", command.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(command.Name))
		initializerStatements += setValueField

		if command.Command != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Command")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+command.Command.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CommandDate")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(command.CommandDate))
		initializerStatements += setValueField

	}

	map_DummyAgent_Identifiers := make(map[*DummyAgent]string)
	_ = map_DummyAgent_Identifiers

	dummyagentOrdered := []*DummyAgent{}
	for dummyagent := range stage.DummyAgents {
		dummyagentOrdered = append(dummyagentOrdered, dummyagent)
	}
	sort.Slice(dummyagentOrdered[:], func(i, j int) bool {
		dummyagenti := dummyagentOrdered[i]
		dummyagentj := dummyagentOrdered[j]
		dummyagenti_order, oki := stage.DummyAgentMap_Staged_Order[dummyagenti]
		dummyagentj_order, okj := stage.DummyAgentMap_Staged_Order[dummyagentj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return dummyagenti_order < dummyagentj_order
	})
	if len(dummyagentOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, dummyagent := range dummyagentOrdered {

		id = generatesIdentifier("DummyAgent", idx, dummyagent.Name)
		map_DummyAgent_Identifiers[dummyagent] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DummyAgent")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", dummyagent.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TechName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(dummyagent.TechName))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(dummyagent.Name))
		initializerStatements += setValueField

	}

	map_Engine_Identifiers := make(map[*Engine]string)
	_ = map_Engine_Identifiers

	engineOrdered := []*Engine{}
	for engine := range stage.Engines {
		engineOrdered = append(engineOrdered, engine)
	}
	sort.Slice(engineOrdered[:], func(i, j int) bool {
		enginei := engineOrdered[i]
		enginej := engineOrdered[j]
		enginei_order, oki := stage.EngineMap_Staged_Order[enginei]
		enginej_order, okj := stage.EngineMap_Staged_Order[enginej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return enginei_order < enginej_order
	})
	if len(engineOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, engine := range engineOrdered {

		id = generatesIdentifier("Engine", idx, engine.Name)
		map_Engine_Identifiers[engine] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Engine")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", engine.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(engine.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndTime")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(engine.EndTime))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CurrentTime")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(engine.CurrentTime))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DisplayFormat")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(engine.DisplayFormat))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SecondsSinceStart")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", engine.SecondsSinceStart))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Fired")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", engine.Fired))
		initializerStatements += setValueField

		if engine.ControlMode != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ControlMode")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+engine.ControlMode.ToCodeString())
			initializerStatements += setValueField
		}

		if engine.State != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "State")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+engine.State.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Speed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", engine.Speed))
		initializerStatements += setValueField

	}

	map_Event_Identifiers := make(map[*Event]string)
	_ = map_Event_Identifiers

	eventOrdered := []*Event{}
	for event := range stage.Events {
		eventOrdered = append(eventOrdered, event)
	}
	sort.Slice(eventOrdered[:], func(i, j int) bool {
		eventi := eventOrdered[i]
		eventj := eventOrdered[j]
		eventi_order, oki := stage.EventMap_Staged_Order[eventi]
		eventj_order, okj := stage.EventMap_Staged_Order[eventj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return eventi_order < eventj_order
	})
	if len(eventOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, event := range eventOrdered {

		id = generatesIdentifier("Event", idx, event.Name)
		map_Event_Identifiers[event] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Event")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", event.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(event.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Duration")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", event.Duration))
		initializerStatements += setValueField

	}

	map_Status_Identifiers := make(map[*Status]string)
	_ = map_Status_Identifiers

	statusOrdered := []*Status{}
	for status := range stage.Statuss {
		statusOrdered = append(statusOrdered, status)
	}
	sort.Slice(statusOrdered[:], func(i, j int) bool {
		statusi := statusOrdered[i]
		statusj := statusOrdered[j]
		statusi_order, oki := stage.StatusMap_Staged_Order[statusi]
		statusj_order, okj := stage.StatusMap_Staged_Order[statusj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return statusi_order < statusj_order
	})
	if len(statusOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, status := range statusOrdered {

		id = generatesIdentifier("Status", idx, status.Name)
		map_Status_Identifiers[status] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Status")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", status.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(status.Name))
		initializerStatements += setValueField

		if status.CurrentCommand != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CurrentCommand")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+status.CurrentCommand.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CompletionDate")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(status.CompletionDate))
		initializerStatements += setValueField

		if status.CurrentSpeedCommand != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CurrentSpeedCommand")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+status.CurrentSpeedCommand.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SpeedCommandCompletionDate")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(status.SpeedCommandCompletionDate))
		initializerStatements += setValueField

	}

	map_UpdateState_Identifiers := make(map[*UpdateState]string)
	_ = map_UpdateState_Identifiers

	updatestateOrdered := []*UpdateState{}
	for updatestate := range stage.UpdateStates {
		updatestateOrdered = append(updatestateOrdered, updatestate)
	}
	sort.Slice(updatestateOrdered[:], func(i, j int) bool {
		updatestatei := updatestateOrdered[i]
		updatestatej := updatestateOrdered[j]
		updatestatei_order, oki := stage.UpdateStateMap_Staged_Order[updatestatei]
		updatestatej_order, okj := stage.UpdateStateMap_Staged_Order[updatestatej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return updatestatei_order < updatestatej_order
	})
	if len(updatestateOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, updatestate := range updatestateOrdered {

		id = generatesIdentifier("UpdateState", idx, updatestate.Name)
		map_UpdateState_Identifiers[updatestate] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "UpdateState")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", updatestate.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(updatestate.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Duration")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", updatestate.Duration))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Period")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", updatestate.Period))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	if len(commandOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Command instances pointers"
	}
	for idx, command := range commandOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Command", idx, command.Name)
		map_Command_Identifiers[command] = id

		// Initialisation of values
		if command.Engine != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Engine")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Engine_Identifiers[command.Engine])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(dummyagentOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of DummyAgent instances pointers"
	}
	for idx, dummyagent := range dummyagentOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DummyAgent", idx, dummyagent.Name)
		map_DummyAgent_Identifiers[dummyagent] = id

		// Initialisation of values
	}

	if len(engineOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Engine instances pointers"
	}
	for idx, engine := range engineOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Engine", idx, engine.Name)
		map_Engine_Identifiers[engine] = id

		// Initialisation of values
	}

	if len(eventOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Event instances pointers"
	}
	for idx, event := range eventOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Event", idx, event.Name)
		map_Event_Identifiers[event] = id

		// Initialisation of values
	}

	if len(statusOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Status instances pointers"
	}
	for idx, status := range statusOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Status", idx, status.Name)
		map_Status_Identifiers[status] = id

		// Initialisation of values
	}

	if len(updatestateOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of UpdateState instances pointers"
	}
	for idx, updatestate := range updatestateOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("UpdateState", idx, updatestate.Name)
		map_UpdateState_Identifiers[updatestate] = id

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
