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
	var identifiersDecl strings.Builder
	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder

	decl := ""
	_ = decl
	setValueField := ""
	_ = setValueField

	// insertion initialization of objects to stage
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
		identifiersDecl.WriteString("\n")
	}
	for _, command := range commandOrdered {

		identifiersDecl.WriteString(command.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(command.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(command.GongMarshallField(stage, "Command"))
		initializerStatements.WriteString(command.GongMarshallField(stage, "CommandDate"))
		pointersInitializesStatements.WriteString(command.GongMarshallField(stage, "Engine"))
	}

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
		identifiersDecl.WriteString("\n")
	}
	for _, dummyagent := range dummyagentOrdered {

		identifiersDecl.WriteString(dummyagent.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(dummyagent.GongMarshallField(stage, "TechName"))
		initializerStatements.WriteString(dummyagent.GongMarshallField(stage, "Name"))
	}

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
		identifiersDecl.WriteString("\n")
	}
	for _, engine := range engineOrdered {

		identifiersDecl.WriteString(engine.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(engine.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(engine.GongMarshallField(stage, "EndTime"))
		initializerStatements.WriteString(engine.GongMarshallField(stage, "CurrentTime"))
		initializerStatements.WriteString(engine.GongMarshallField(stage, "DisplayFormat"))
		initializerStatements.WriteString(engine.GongMarshallField(stage, "SecondsSinceStart"))
		initializerStatements.WriteString(engine.GongMarshallField(stage, "Fired"))
		initializerStatements.WriteString(engine.GongMarshallField(stage, "ControlMode"))
		initializerStatements.WriteString(engine.GongMarshallField(stage, "State"))
		initializerStatements.WriteString(engine.GongMarshallField(stage, "Speed"))
	}

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
		identifiersDecl.WriteString("\n")
	}
	for _, event := range eventOrdered {

		identifiersDecl.WriteString(event.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(event.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(event.GongMarshallField(stage, "Duration"))
	}

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
		identifiersDecl.WriteString("\n")
	}
	for _, status := range statusOrdered {

		identifiersDecl.WriteString(status.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(status.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(status.GongMarshallField(stage, "CurrentCommand"))
		initializerStatements.WriteString(status.GongMarshallField(stage, "CompletionDate"))
		initializerStatements.WriteString(status.GongMarshallField(stage, "CurrentSpeedCommand"))
		initializerStatements.WriteString(status.GongMarshallField(stage, "SpeedCommandCompletionDate"))
	}

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
		identifiersDecl.WriteString("\n")
	}
	for _, updatestate := range updatestateOrdered {

		identifiersDecl.WriteString(updatestate.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(updatestate.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(updatestate.GongMarshallField(stage, "Duration"))
		initializerStatements.WriteString(updatestate.GongMarshallField(stage, "Period"))
	}

	// insertion initialization of objects to stage
	for _, command := range commandOrdered {
		_ = command
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, dummyagent := range dummyagentOrdered {
		_ = dummyagent
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, engine := range engineOrdered {
		_ = engine
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, event := range eventOrdered {
		_ = event
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, status := range statusOrdered {
		_ = status
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, updatestate := range updatestateOrdered {
		_ = updatestate
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl.String())
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements.String())
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements.String())

	if stage.MetaPackageImportAlias != "" {
		res = strings.ReplaceAll(res, "{{ImportPackageDeclaration}}",
			fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))

		res = strings.ReplaceAll(res, "{{ImportPackageDummyDeclaration}}",
			fmt.Sprintf("\nvar _ %s.Stage",
				stage.MetaPackageImportAlias))

		var entries strings.Builder

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
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": %s(0),", value.Ident, value.Ident))
			case GONG__ENUM_CAST_STRING:
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": %s(\"\"),", value.Ident, value.Ident))
			case GONG__FIELD_VALUE:
				// substitute the second point with "{})."
				joker := "__substitute_for_first_point__"
				valueIdentifier := strings.Replace(value.Ident, ".", joker, 1)
				valueIdentifier = strings.Replace(valueIdentifier, ".", "{}).", 1)
				valueIdentifier = strings.Replace(valueIdentifier, joker, ".", 1)
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": (%s,", value.Ident, valueIdentifier))
			case GONG__IDENTIFIER_CONST:
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": %s,", value.Ident, value.Ident))
			case GONG__STRUCT_INSTANCE:
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": &(%s{}),", value.Ident, value.Ident))
			}
		}

		// res = strings.ReplaceAll(res, "{{EntriesDocLinkStringDocLinkIdentifier}}", entries.String())
	}
	return
}

// insertion point for marshall field methods
func (command *Command) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", command.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(command.Name))
	case "Command":
		if command.Command.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", command.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Command")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+command.Command.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", command.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Command")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "CommandDate":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", command.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CommandDate")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(command.CommandDate))

	case "Engine":
		if command.Engine != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", command.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Engine")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", command.Engine.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", command.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Engine")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Command", fieldName)
	}
	return
}

func (dummyagent *DummyAgent) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "TechName":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", dummyagent.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TechName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(dummyagent.TechName))
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", dummyagent.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(dummyagent.Name))

	default:
		log.Panicf("Unknown field %s for Gongstruct DummyAgent", fieldName)
	}
	return
}

func (engine *Engine) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", engine.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(engine.Name))
	case "EndTime":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", engine.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndTime")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(engine.EndTime))
	case "CurrentTime":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", engine.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CurrentTime")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(engine.CurrentTime))
	case "DisplayFormat":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", engine.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DisplayFormat")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(engine.DisplayFormat))
	case "SecondsSinceStart":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", engine.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SecondsSinceStart")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", engine.SecondsSinceStart))
	case "Fired":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", engine.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Fired")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", engine.Fired))
	case "ControlMode":
		if engine.ControlMode.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", engine.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ControlMode")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+engine.ControlMode.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", engine.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ControlMode")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "State":
		if engine.State.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", engine.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "State")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+engine.State.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", engine.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "State")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "Speed":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", engine.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Speed")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", engine.Speed))

	default:
		log.Panicf("Unknown field %s for Gongstruct Engine", fieldName)
	}
	return
}

func (event *Event) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", event.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(event.Name))
	case "Duration":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", event.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Duration")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", event.Duration))

	default:
		log.Panicf("Unknown field %s for Gongstruct Event", fieldName)
	}
	return
}

func (status *Status) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", status.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(status.Name))
	case "CurrentCommand":
		if status.CurrentCommand.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", status.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CurrentCommand")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+status.CurrentCommand.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", status.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CurrentCommand")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "CompletionDate":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", status.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CompletionDate")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(status.CompletionDate))
	case "CurrentSpeedCommand":
		if status.CurrentSpeedCommand.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", status.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CurrentSpeedCommand")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+status.CurrentSpeedCommand.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", status.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CurrentSpeedCommand")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "SpeedCommandCompletionDate":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", status.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SpeedCommandCompletionDate")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(status.SpeedCommandCompletionDate))

	default:
		log.Panicf("Unknown field %s for Gongstruct Status", fieldName)
	}
	return
}

func (updatestate *UpdateState) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", updatestate.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(updatestate.Name))
	case "Duration":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", updatestate.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Duration")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", updatestate.Duration))
	case "Period":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", updatestate.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Period")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", updatestate.Period))

	default:
		log.Panicf("Unknown field %s for Gongstruct UpdateState", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (command *Command) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(command.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(command.GongMarshallField(stage, "Command"))
		initializerStatements.WriteString(command.GongMarshallField(stage, "CommandDate"))
		pointersInitializesStatements.WriteString(command.GongMarshallField(stage, "Engine"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (dummyagent *DummyAgent) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(dummyagent.GongMarshallField(stage, "TechName"))
		initializerStatements.WriteString(dummyagent.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (engine *Engine) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(engine.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(engine.GongMarshallField(stage, "EndTime"))
		initializerStatements.WriteString(engine.GongMarshallField(stage, "CurrentTime"))
		initializerStatements.WriteString(engine.GongMarshallField(stage, "DisplayFormat"))
		initializerStatements.WriteString(engine.GongMarshallField(stage, "SecondsSinceStart"))
		initializerStatements.WriteString(engine.GongMarshallField(stage, "Fired"))
		initializerStatements.WriteString(engine.GongMarshallField(stage, "ControlMode"))
		initializerStatements.WriteString(engine.GongMarshallField(stage, "State"))
		initializerStatements.WriteString(engine.GongMarshallField(stage, "Speed"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (event *Event) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(event.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(event.GongMarshallField(stage, "Duration"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (status *Status) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(status.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(status.GongMarshallField(stage, "CurrentCommand"))
		initializerStatements.WriteString(status.GongMarshallField(stage, "CompletionDate"))
		initializerStatements.WriteString(status.GongMarshallField(stage, "CurrentSpeedCommand"))
		initializerStatements.WriteString(status.GongMarshallField(stage, "SpeedCommandCompletionDate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (updatestate *UpdateState) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(updatestate.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(updatestate.GongMarshallField(stage, "Duration"))
		initializerStatements.WriteString(updatestate.GongMarshallField(stage, "Period"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
