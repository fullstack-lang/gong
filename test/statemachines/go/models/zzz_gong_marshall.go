// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"os"
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
}
`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{}).Stage(stage)`

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

	id := ""
	_ = id
	decl := ""
	_ = decl
	setValueField := ""
	_ = setValueField

	// insertion initialization of objects to stage
	map_Architecture_Identifiers := make(map[*Architecture]string)
	_ = map_Architecture_Identifiers

	architectureOrdered := []*Architecture{}
	for architecture := range stage.Architectures {
		architectureOrdered = append(architectureOrdered, architecture)
	}
	sort.Slice(architectureOrdered[:], func(i, j int) bool {
		architecturei := architectureOrdered[i]
		architecturej := architectureOrdered[j]
		architecturei_order, oki := stage.ArchitectureMap_Staged_Order[architecturei]
		architecturej_order, okj := stage.ArchitectureMap_Staged_Order[architecturej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return architecturei_order < architecturej_order
	})
	if len(architectureOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, architecture := range architectureOrdered {

		id = generatesIdentifier("Architecture", idx, architecture.Name)
		map_Architecture_Identifiers[architecture] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Architecture")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", architecture.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(architecture.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NbPixPerCharacter")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", architecture.NbPixPerCharacter))
		initializerStatements += setValueField

	}

	map_Diagram_Identifiers := make(map[*Diagram]string)
	_ = map_Diagram_Identifiers

	diagramOrdered := []*Diagram{}
	for diagram := range stage.Diagrams {
		diagramOrdered = append(diagramOrdered, diagram)
	}
	sort.Slice(diagramOrdered[:], func(i, j int) bool {
		diagrami := diagramOrdered[i]
		diagramj := diagramOrdered[j]
		diagrami_order, oki := stage.DiagramMap_Staged_Order[diagrami]
		diagramj_order, okj := stage.DiagramMap_Staged_Order[diagramj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return diagrami_order < diagramj_order
	})
	if len(diagramOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, diagram := range diagramOrdered {

		id = generatesIdentifier("Diagram", idx, diagram.Name)
		map_Diagram_Identifiers[diagram] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Diagram")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", diagram.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsChecked")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsChecked))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsEditable_")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsEditable_))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsInRenameMode")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsInRenameMode))
		initializerStatements += setValueField

	}

	map_DoAction_Identifiers := make(map[*DoAction]string)
	_ = map_DoAction_Identifiers

	doactionOrdered := []*DoAction{}
	for doaction := range stage.DoActions {
		doactionOrdered = append(doactionOrdered, doaction)
	}
	sort.Slice(doactionOrdered[:], func(i, j int) bool {
		doactioni := doactionOrdered[i]
		doactionj := doactionOrdered[j]
		doactioni_order, oki := stage.DoActionMap_Staged_Order[doactioni]
		doactionj_order, okj := stage.DoActionMap_Staged_Order[doactionj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return doactioni_order < doactionj_order
	})
	if len(doactionOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, doaction := range doactionOrdered {

		id = generatesIdentifier("DoAction", idx, doaction.Name)
		map_DoAction_Identifiers[doaction] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DoAction")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", doaction.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(doaction.Name))
		initializerStatements += setValueField

		if doaction.Criticality != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Criticality")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+doaction.Criticality.ToCodeString())
			initializerStatements += setValueField
		}

	}

	map_Kill_Identifiers := make(map[*Kill]string)
	_ = map_Kill_Identifiers

	killOrdered := []*Kill{}
	for kill := range stage.Kills {
		killOrdered = append(killOrdered, kill)
	}
	sort.Slice(killOrdered[:], func(i, j int) bool {
		killi := killOrdered[i]
		killj := killOrdered[j]
		killi_order, oki := stage.KillMap_Staged_Order[killi]
		killj_order, okj := stage.KillMap_Staged_Order[killj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return killi_order < killj_order
	})
	if len(killOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, kill := range killOrdered {

		id = generatesIdentifier("Kill", idx, kill.Name)
		map_Kill_Identifiers[kill] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Kill")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", kill.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(kill.Name))
		initializerStatements += setValueField

	}

	map_Message_Identifiers := make(map[*Message]string)
	_ = map_Message_Identifiers

	messageOrdered := []*Message{}
	for message := range stage.Messages {
		messageOrdered = append(messageOrdered, message)
	}
	sort.Slice(messageOrdered[:], func(i, j int) bool {
		messagei := messageOrdered[i]
		messagej := messageOrdered[j]
		messagei_order, oki := stage.MessageMap_Staged_Order[messagei]
		messagej_order, okj := stage.MessageMap_Staged_Order[messagej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return messagei_order < messagej_order
	})
	if len(messageOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, message := range messageOrdered {

		id = generatesIdentifier("Message", idx, message.Name)
		map_Message_Identifiers[message] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Message")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", message.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(message.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", message.IsSelected))
		initializerStatements += setValueField

	}

	map_MessageType_Identifiers := make(map[*MessageType]string)
	_ = map_MessageType_Identifiers

	messagetypeOrdered := []*MessageType{}
	for messagetype := range stage.MessageTypes {
		messagetypeOrdered = append(messagetypeOrdered, messagetype)
	}
	sort.Slice(messagetypeOrdered[:], func(i, j int) bool {
		messagetypei := messagetypeOrdered[i]
		messagetypej := messagetypeOrdered[j]
		messagetypei_order, oki := stage.MessageTypeMap_Staged_Order[messagetypei]
		messagetypej_order, okj := stage.MessageTypeMap_Staged_Order[messagetypej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return messagetypei_order < messagetypej_order
	})
	if len(messagetypeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, messagetype := range messagetypeOrdered {

		id = generatesIdentifier("MessageType", idx, messagetype.Name)
		map_MessageType_Identifiers[messagetype] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MessageType")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", messagetype.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(messagetype.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Description")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(messagetype.Description))
		initializerStatements += setValueField

	}

	map_Object_Identifiers := make(map[*Object]string)
	_ = map_Object_Identifiers

	objectOrdered := []*Object{}
	for object := range stage.Objects {
		objectOrdered = append(objectOrdered, object)
	}
	sort.Slice(objectOrdered[:], func(i, j int) bool {
		objecti := objectOrdered[i]
		objectj := objectOrdered[j]
		objecti_order, oki := stage.ObjectMap_Staged_Order[objecti]
		objectj_order, okj := stage.ObjectMap_Staged_Order[objectj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return objecti_order < objectj_order
	})
	if len(objectOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, object := range objectOrdered {

		id = generatesIdentifier("Object", idx, object.Name)
		map_Object_Identifiers[object] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Object")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", object.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(object.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", object.IsSelected))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Rank")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", object.Rank))
		initializerStatements += setValueField

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DOF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", object.DOF.String())
		initializerStatements += setValueField

	}

	map_Role_Identifiers := make(map[*Role]string)
	_ = map_Role_Identifiers

	roleOrdered := []*Role{}
	for role := range stage.Roles {
		roleOrdered = append(roleOrdered, role)
	}
	sort.Slice(roleOrdered[:], func(i, j int) bool {
		rolei := roleOrdered[i]
		rolej := roleOrdered[j]
		rolei_order, oki := stage.RoleMap_Staged_Order[rolei]
		rolej_order, okj := stage.RoleMap_Staged_Order[rolej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return rolei_order < rolej_order
	})
	if len(roleOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, role := range roleOrdered {

		id = generatesIdentifier("Role", idx, role.Name)
		map_Role_Identifiers[role] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Role")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", role.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(role.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Acronym")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(role.Acronym))
		initializerStatements += setValueField

	}

	map_State_Identifiers := make(map[*State]string)
	_ = map_State_Identifiers

	stateOrdered := []*State{}
	for state := range stage.States {
		stateOrdered = append(stateOrdered, state)
	}
	sort.Slice(stateOrdered[:], func(i, j int) bool {
		statei := stateOrdered[i]
		statej := stateOrdered[j]
		statei_order, oki := stage.StateMap_Staged_Order[statei]
		statej_order, okj := stage.StateMap_Staged_Order[statej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return statei_order < statej_order
	})
	if len(stateOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, state := range stateOrdered {

		id = generatesIdentifier("State", idx, state.Name)
		map_State_Identifiers[state] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "State")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", state.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(state.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDecisionNode")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", state.IsDecisionNode))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsFictif")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", state.IsFictif))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsEndState")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", state.IsEndState))
		initializerStatements += setValueField

	}

	map_StateMachine_Identifiers := make(map[*StateMachine]string)
	_ = map_StateMachine_Identifiers

	statemachineOrdered := []*StateMachine{}
	for statemachine := range stage.StateMachines {
		statemachineOrdered = append(statemachineOrdered, statemachine)
	}
	sort.Slice(statemachineOrdered[:], func(i, j int) bool {
		statemachinei := statemachineOrdered[i]
		statemachinej := statemachineOrdered[j]
		statemachinei_order, oki := stage.StateMachineMap_Staged_Order[statemachinei]
		statemachinej_order, okj := stage.StateMachineMap_Staged_Order[statemachinej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return statemachinei_order < statemachinej_order
	})
	if len(statemachineOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, statemachine := range statemachineOrdered {

		id = generatesIdentifier("StateMachine", idx, statemachine.Name)
		map_StateMachine_Identifiers[statemachine] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StateMachine")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", statemachine.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(statemachine.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", statemachine.IsNodeExpanded))
		initializerStatements += setValueField

	}

	map_StateShape_Identifiers := make(map[*StateShape]string)
	_ = map_StateShape_Identifiers

	stateshapeOrdered := []*StateShape{}
	for stateshape := range stage.StateShapes {
		stateshapeOrdered = append(stateshapeOrdered, stateshape)
	}
	sort.Slice(stateshapeOrdered[:], func(i, j int) bool {
		stateshapei := stateshapeOrdered[i]
		stateshapej := stateshapeOrdered[j]
		stateshapei_order, oki := stage.StateShapeMap_Staged_Order[stateshapei]
		stateshapej_order, okj := stage.StateShapeMap_Staged_Order[stateshapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return stateshapei_order < stateshapej_order
	})
	if len(stateshapeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, stateshape := range stateshapeOrdered {

		id = generatesIdentifier("StateShape", idx, stateshape.Name)
		map_StateShape_Identifiers[stateshape] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StateShape")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", stateshape.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(stateshape.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", stateshape.IsExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", stateshape.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", stateshape.Y))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", stateshape.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", stateshape.Height))
		initializerStatements += setValueField

	}

	map_Transition_Identifiers := make(map[*Transition]string)
	_ = map_Transition_Identifiers

	transitionOrdered := []*Transition{}
	for transition := range stage.Transitions {
		transitionOrdered = append(transitionOrdered, transition)
	}
	sort.Slice(transitionOrdered[:], func(i, j int) bool {
		transitioni := transitionOrdered[i]
		transitionj := transitionOrdered[j]
		transitioni_order, oki := stage.TransitionMap_Staged_Order[transitioni]
		transitionj_order, okj := stage.TransitionMap_Staged_Order[transitionj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return transitioni_order < transitionj_order
	})
	if len(transitionOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, transition := range transitionOrdered {

		id = generatesIdentifier("Transition", idx, transition.Name)
		map_Transition_Identifiers[transition] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Transition")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", transition.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(transition.Name))
		initializerStatements += setValueField

	}

	map_Transition_Shape_Identifiers := make(map[*Transition_Shape]string)
	_ = map_Transition_Shape_Identifiers

	transition_shapeOrdered := []*Transition_Shape{}
	for transition_shape := range stage.Transition_Shapes {
		transition_shapeOrdered = append(transition_shapeOrdered, transition_shape)
	}
	sort.Slice(transition_shapeOrdered[:], func(i, j int) bool {
		transition_shapei := transition_shapeOrdered[i]
		transition_shapej := transition_shapeOrdered[j]
		transition_shapei_order, oki := stage.Transition_ShapeMap_Staged_Order[transition_shapei]
		transition_shapej_order, okj := stage.Transition_ShapeMap_Staged_Order[transition_shapej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return transition_shapei_order < transition_shapej_order
	})
	if len(transition_shapeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, transition_shape := range transition_shapeOrdered {

		id = generatesIdentifier("Transition_Shape", idx, transition_shape.Name)
		map_Transition_Shape_Identifiers[transition_shape] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Transition_Shape")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", transition_shape.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(transition_shape.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", transition_shape.StartRatio))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", transition_shape.EndRatio))
		initializerStatements += setValueField

		if transition_shape.StartOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+transition_shape.StartOrientation.ToCodeString())
			initializerStatements += setValueField
		}

		if transition_shape.EndOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+transition_shape.EndOrientation.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", transition_shape.CornerOffsetRatio))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	if len(architectureOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Architecture instances pointers"
	}
	for idx, architecture := range architectureOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Architecture", idx, architecture.Name)
		map_Architecture_Identifiers[architecture] = id

		// Initialisation of values
		for _, _statemachine := range architecture.StateMachines {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "StateMachines")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_StateMachine_Identifiers[_statemachine])
			pointersInitializesStatements += setPointerField
		}

		for _, _role := range architecture.Roles {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Roles")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Role_Identifiers[_role])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(diagramOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Diagram instances pointers"
	}
	for idx, diagram := range diagramOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Diagram", idx, diagram.Name)
		map_Diagram_Identifiers[diagram] = id

		// Initialisation of values
		for _, _stateshape := range diagram.State_Shapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "State_Shapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_StateShape_Identifiers[_stateshape])
			pointersInitializesStatements += setPointerField
		}

		for _, _transition_shape := range diagram.Transition_Shapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Transition_Shapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Transition_Shape_Identifiers[_transition_shape])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(doactionOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of DoAction instances pointers"
	}
	for idx, doaction := range doactionOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DoAction", idx, doaction.Name)
		map_DoAction_Identifiers[doaction] = id

		// Initialisation of values
	}

	if len(killOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Kill instances pointers"
	}
	for idx, kill := range killOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Kill", idx, kill.Name)
		map_Kill_Identifiers[kill] = id

		// Initialisation of values
	}

	if len(messageOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Message instances pointers"
	}
	for idx, message := range messageOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Message", idx, message.Name)
		map_Message_Identifiers[message] = id

		// Initialisation of values
		if message.MessageType != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "MessageType")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_MessageType_Identifiers[message.MessageType])
			pointersInitializesStatements += setPointerField
		}

		if message.OriginTransition != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "OriginTransition")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Transition_Identifiers[message.OriginTransition])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(messagetypeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of MessageType instances pointers"
	}
	for idx, messagetype := range messagetypeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("MessageType", idx, messagetype.Name)
		map_MessageType_Identifiers[messagetype] = id

		// Initialisation of values
	}

	if len(objectOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Object instances pointers"
	}
	for idx, object := range objectOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Object", idx, object.Name)
		map_Object_Identifiers[object] = id

		// Initialisation of values
		if object.State != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "State")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_State_Identifiers[object.State])
			pointersInitializesStatements += setPointerField
		}

		for _, _message := range object.Messages {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Messages")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Message_Identifiers[_message])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(roleOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Role instances pointers"
	}
	for idx, role := range roleOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Role", idx, role.Name)
		map_Role_Identifiers[role] = id

		// Initialisation of values
		for _, _role := range role.RolesWithSamePermissions {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RolesWithSamePermissions")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Role_Identifiers[_role])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(stateOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of State instances pointers"
	}
	for idx, state := range stateOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("State", idx, state.Name)
		map_State_Identifiers[state] = id

		// Initialisation of values
		if state.Parent != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Parent")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_State_Identifiers[state.Parent])
			pointersInitializesStatements += setPointerField
		}

		for _, _state := range state.SubStates {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SubStates")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_State_Identifiers[_state])
			pointersInitializesStatements += setPointerField
		}

		for _, _diagram := range state.Diagrams {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Diagrams")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Diagram_Identifiers[_diagram])
			pointersInitializesStatements += setPointerField
		}

		for _, _doaction := range state.DoActions {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DoActions")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_DoAction_Identifiers[_doaction])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(statemachineOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of StateMachine instances pointers"
	}
	for idx, statemachine := range statemachineOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("StateMachine", idx, statemachine.Name)
		map_StateMachine_Identifiers[statemachine] = id

		// Initialisation of values
		for _, _state := range statemachine.States {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "States")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_State_Identifiers[_state])
			pointersInitializesStatements += setPointerField
		}

		for _, _diagram := range statemachine.Diagrams {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Diagrams")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Diagram_Identifiers[_diagram])
			pointersInitializesStatements += setPointerField
		}

		if statemachine.InitialState != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "InitialState")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_State_Identifiers[statemachine.InitialState])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(stateshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of StateShape instances pointers"
	}
	for idx, stateshape := range stateshapeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("StateShape", idx, stateshape.Name)
		map_StateShape_Identifiers[stateshape] = id

		// Initialisation of values
		if stateshape.State != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "State")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_State_Identifiers[stateshape.State])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(transitionOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Transition instances pointers"
	}
	for idx, transition := range transitionOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Transition", idx, transition.Name)
		map_Transition_Identifiers[transition] = id

		// Initialisation of values
		if transition.Start != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Start")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_State_Identifiers[transition.Start])
			pointersInitializesStatements += setPointerField
		}

		if transition.End != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "End")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_State_Identifiers[transition.End])
			pointersInitializesStatements += setPointerField
		}

		for _, _role := range transition.RolesWithPermissions {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RolesWithPermissions")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Role_Identifiers[_role])
			pointersInitializesStatements += setPointerField
		}

		for _, _messagetype := range transition.GeneratedMessages {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GeneratedMessages")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_MessageType_Identifiers[_messagetype])
			pointersInitializesStatements += setPointerField
		}

		for _, _diagram := range transition.Diagrams {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Diagrams")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Diagram_Identifiers[_diagram])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(transition_shapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Transition_Shape instances pointers"
	}
	for idx, transition_shape := range transition_shapeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Transition_Shape", idx, transition_shape.Name)
		map_Transition_Shape_Identifiers[transition_shape] = id

		// Initialisation of values
		if transition_shape.Transition != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Transition")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Transition_Identifiers[transition_shape.Transition])
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
