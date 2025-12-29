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
	actionOrdered := []*Action{}
	for action := range stage.Actions {
		actionOrdered = append(actionOrdered, action)
	}
	sort.Slice(actionOrdered[:], func(i, j int) bool {
		actioni := actionOrdered[i]
		actionj := actionOrdered[j]
		actioni_order, oki := stage.ActionMap_Staged_Order[actioni]
		actionj_order, okj := stage.ActionMap_Staged_Order[actionj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return actioni_order < actionj_order
	})
	if len(actionOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, action := range actionOrdered {
	
		identifiersDecl += action.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", action.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(action.Name))
		initializerStatements += setValueField

		if action.Criticality != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", action.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Criticality")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+action.Criticality.ToCodeString())
			initializerStatements += setValueField
		}

	}

	activitiesOrdered := []*Activities{}
	for activities := range stage.Activitiess {
		activitiesOrdered = append(activitiesOrdered, activities)
	}
	sort.Slice(activitiesOrdered[:], func(i, j int) bool {
		activitiesi := activitiesOrdered[i]
		activitiesj := activitiesOrdered[j]
		activitiesi_order, oki := stage.ActivitiesMap_Staged_Order[activitiesi]
		activitiesj_order, okj := stage.ActivitiesMap_Staged_Order[activitiesj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return activitiesi_order < activitiesj_order
	})
	if len(activitiesOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, activities := range activitiesOrdered {
	
		identifiersDecl += activities.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", activities.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(activities.Name))
		initializerStatements += setValueField

		if activities.Criticality != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", activities.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Criticality")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+activities.Criticality.ToCodeString())
			initializerStatements += setValueField
		}

	}

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
	for _, architecture := range architectureOrdered {
	
		identifiersDecl += architecture.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", architecture.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(architecture.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", architecture.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NbPixPerCharacter")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", architecture.NbPixPerCharacter))
		initializerStatements += setValueField

	}

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
	for _, diagram := range diagramOrdered {
	
		identifiersDecl += diagram.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsChecked")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsChecked))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsEditable_")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsEditable_))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsInRenameMode")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsInRenameMode))
		initializerStatements += setValueField

	}

	guardOrdered := []*Guard{}
	for guard := range stage.Guards {
		guardOrdered = append(guardOrdered, guard)
	}
	sort.Slice(guardOrdered[:], func(i, j int) bool {
		guardi := guardOrdered[i]
		guardj := guardOrdered[j]
		guardi_order, oki := stage.GuardMap_Staged_Order[guardi]
		guardj_order, okj := stage.GuardMap_Staged_Order[guardj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return guardi_order < guardj_order
	})
	if len(guardOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, guard := range guardOrdered {
	
		identifiersDecl += guard.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", guard.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(guard.Name))
		initializerStatements += setValueField

	}

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
	for _, kill := range killOrdered {
	
		identifiersDecl += kill.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", kill.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(kill.Name))
		initializerStatements += setValueField

	}

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
	for _, message := range messageOrdered {
	
		identifiersDecl += message.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", message.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(message.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", message.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", message.IsSelected))
		initializerStatements += setValueField

	}

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
	for _, messagetype := range messagetypeOrdered {
	
		identifiersDecl += messagetype.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", messagetype.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(messagetype.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", messagetype.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Description")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(messagetype.Description))
		initializerStatements += setValueField

	}

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
	for _, object := range objectOrdered {
	
		identifiersDecl += object.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", object.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(object.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", object.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", object.IsSelected))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", object.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Rank")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", object.Rank))
		initializerStatements += setValueField

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", object.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DOF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", object.DOF.String())
		initializerStatements += setValueField

	}

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
	for _, role := range roleOrdered {
	
		identifiersDecl += role.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", role.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(role.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", role.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Acronym")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(role.Acronym))
		initializerStatements += setValueField

	}

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
	for _, state := range stateOrdered {
	
		identifiersDecl += state.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", state.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(state.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", state.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDecisionNode")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", state.IsDecisionNode))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", state.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsFictif")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", state.IsFictif))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", state.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsEndState")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", state.IsEndState))
		initializerStatements += setValueField

	}

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
	for _, statemachine := range statemachineOrdered {
	
		identifiersDecl += statemachine.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", statemachine.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(statemachine.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", statemachine.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", statemachine.IsNodeExpanded))
		initializerStatements += setValueField

	}

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
	for _, stateshape := range stateshapeOrdered {
	
		identifiersDecl += stateshape.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", stateshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(stateshape.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", stateshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", stateshape.IsExpanded))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", stateshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", stateshape.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", stateshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", stateshape.Y))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", stateshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", stateshape.Width))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", stateshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", stateshape.Height))
		initializerStatements += setValueField

	}

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
	for _, transition := range transitionOrdered {
	
		identifiersDecl += transition.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", transition.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(transition.Name))
		initializerStatements += setValueField

	}

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
	for _, transition_shape := range transition_shapeOrdered {
	
		identifiersDecl += transition_shape.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", transition_shape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(transition_shape.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", transition_shape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", transition_shape.StartRatio))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", transition_shape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", transition_shape.EndRatio))
		initializerStatements += setValueField

		if transition_shape.StartOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", transition_shape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+transition_shape.StartOrientation.ToCodeString())
			initializerStatements += setValueField
		}

		if transition_shape.EndOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", transition_shape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+transition_shape.EndOrientation.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", transition_shape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", transition_shape.CornerOffsetRatio))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	if len(actionOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Action instances pointers"
	}
	for _, action := range actionOrdered {
		_ = action
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(activitiesOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Activities instances pointers"
	}
	for _, activities := range activitiesOrdered {
		_ = activities
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(architectureOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Architecture instances pointers"
	}
	for _, architecture := range architectureOrdered {
		_ = architecture
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		for _, _statemachine := range architecture.StateMachines {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", architecture.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "StateMachines")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _statemachine.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _role := range architecture.Roles {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", architecture.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Roles")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _role.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(diagramOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Diagram instances pointers"
	}
	for _, diagram := range diagramOrdered {
		_ = diagram
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		for _, _stateshape := range diagram.State_Shapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "State_Shapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _stateshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _transition_shape := range diagram.Transition_Shapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Transition_Shapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _transition_shape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(guardOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Guard instances pointers"
	}
	for _, guard := range guardOrdered {
		_ = guard
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(killOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Kill instances pointers"
	}
	for _, kill := range killOrdered {
		_ = kill
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(messageOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Message instances pointers"
	}
	for _, message := range messageOrdered {
		_ = message
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if message.MessageType != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", message.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "MessageType")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", message.MessageType.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if message.OriginTransition != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", message.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "OriginTransition")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", message.OriginTransition.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(messagetypeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of MessageType instances pointers"
	}
	for _, messagetype := range messagetypeOrdered {
		_ = messagetype
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(objectOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Object instances pointers"
	}
	for _, object := range objectOrdered {
		_ = object
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if object.State != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", object.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "State")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", object.State.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _message := range object.Messages {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", object.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Messages")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _message.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(roleOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Role instances pointers"
	}
	for _, role := range roleOrdered {
		_ = role
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		for _, _role := range role.RolesWithSamePermissions {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", role.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RolesWithSamePermissions")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _role.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(stateOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of State instances pointers"
	}
	for _, state := range stateOrdered {
		_ = state
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if state.Parent != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", state.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Parent")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", state.Parent.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _state := range state.SubStates {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", state.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SubStates")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _state.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _diagram := range state.Diagrams {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", state.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Diagrams")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _diagram.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if state.Entry != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", state.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Entry")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", state.Entry.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _activities := range state.Activities {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", state.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Activities")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _activities.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if state.Exit != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", state.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Exit")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", state.Exit.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(statemachineOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of StateMachine instances pointers"
	}
	for _, statemachine := range statemachineOrdered {
		_ = statemachine
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		for _, _state := range statemachine.States {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", statemachine.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "States")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _state.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _diagram := range statemachine.Diagrams {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", statemachine.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Diagrams")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _diagram.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if statemachine.InitialState != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", statemachine.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "InitialState")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", statemachine.InitialState.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(stateshapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of StateShape instances pointers"
	}
	for _, stateshape := range stateshapeOrdered {
		_ = stateshape
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if stateshape.State != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", stateshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "State")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", stateshape.State.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(transitionOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Transition instances pointers"
	}
	for _, transition := range transitionOrdered {
		_ = transition
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if transition.Start != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", transition.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Start")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", transition.Start.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if transition.End != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", transition.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "End")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", transition.End.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _role := range transition.RolesWithPermissions {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", transition.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RolesWithPermissions")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _role.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _messagetype := range transition.GeneratedMessages {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", transition.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GeneratedMessages")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _messagetype.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if transition.Guard != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", transition.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Guard")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", transition.Guard.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _diagram := range transition.Diagrams {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", transition.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Diagrams")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _diagram.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(transition_shapeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Transition_Shape instances pointers"
	}
	for _, transition_shape := range transition_shapeOrdered {
		_ = transition_shape
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if transition_shape.Transition != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", transition_shape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Transition")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", transition_shape.Transition.GongGetIdentifier(stage))
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
func (action *Action) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", action.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(action.Name))
		initializerStatements += setValueField
	case "Criticality":
		if action.Criticality != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", action.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Criticality")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+action.Criticality.ToCodeString())
			initializerStatements += setValueField
		}


	default:
		log.Panicf("Unknown field %s for Gongstruct Action", fieldName)
	}
	return
}

func (activities *Activities) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", activities.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(activities.Name))
		initializerStatements += setValueField
	case "Criticality":
		if activities.Criticality != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", activities.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Criticality")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+activities.Criticality.ToCodeString())
			initializerStatements += setValueField
		}


	default:
		log.Panicf("Unknown field %s for Gongstruct Activities", fieldName)
	}
	return
}

func (architecture *Architecture) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", architecture.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(architecture.Name))
		initializerStatements += setValueField
	case "NbPixPerCharacter":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", architecture.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NbPixPerCharacter")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", architecture.NbPixPerCharacter))
		initializerStatements += setValueField

	case "StateMachines":
		for _, _statemachine := range architecture.StateMachines {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", architecture.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "StateMachines")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _statemachine.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Roles":
		for _, _role := range architecture.Roles {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", architecture.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Roles")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _role.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct Architecture", fieldName)
	}
	return
}

func (diagram *Diagram) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(diagram.Name))
		initializerStatements += setValueField
	case "IsChecked":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsChecked")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsChecked))
		initializerStatements += setValueField
	case "IsExpanded":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsExpanded))
		initializerStatements += setValueField
	case "IsEditable_":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsEditable_")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsEditable_))
		initializerStatements += setValueField
	case "IsInRenameMode":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsInRenameMode")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsInRenameMode))
		initializerStatements += setValueField

	case "State_Shapes":
		for _, _stateshape := range diagram.State_Shapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "State_Shapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _stateshape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Transition_Shapes":
		for _, _transition_shape := range diagram.Transition_Shapes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Transition_Shapes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _transition_shape.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct Diagram", fieldName)
	}
	return
}

func (guard *Guard) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", guard.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(guard.Name))
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct Guard", fieldName)
	}
	return
}

func (kill *Kill) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", kill.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(kill.Name))
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct Kill", fieldName)
	}
	return
}

func (message *Message) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", message.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(message.Name))
		initializerStatements += setValueField
	case "IsSelected":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", message.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", message.IsSelected))
		initializerStatements += setValueField

	case "MessageType":
		if message.MessageType != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", message.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "MessageType")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", message.MessageType.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "OriginTransition":
		if message.OriginTransition != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", message.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "OriginTransition")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", message.OriginTransition.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct Message", fieldName)
	}
	return
}

func (messagetype *MessageType) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", messagetype.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(messagetype.Name))
		initializerStatements += setValueField
	case "Description":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", messagetype.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Description")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(messagetype.Description))
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct MessageType", fieldName)
	}
	return
}

func (object *Object) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", object.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(object.Name))
		initializerStatements += setValueField
	case "IsSelected":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", object.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsSelected")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", object.IsSelected))
		initializerStatements += setValueField
	case "Rank":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", object.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Rank")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", object.Rank))
		initializerStatements += setValueField
	case "DOF":
		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", object.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "DOF")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", object.DOF.String())
		initializerStatements += setValueField

	case "State":
		if object.State != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", object.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "State")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", object.State.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Messages":
		for _, _message := range object.Messages {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", object.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Messages")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _message.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct Object", fieldName)
	}
	return
}

func (role *Role) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", role.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(role.Name))
		initializerStatements += setValueField
	case "Acronym":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", role.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Acronym")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(role.Acronym))
		initializerStatements += setValueField

	case "RolesWithSamePermissions":
		for _, _role := range role.RolesWithSamePermissions {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", role.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RolesWithSamePermissions")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _role.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct Role", fieldName)
	}
	return
}

func (state *State) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", state.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(state.Name))
		initializerStatements += setValueField
	case "IsDecisionNode":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", state.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDecisionNode")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", state.IsDecisionNode))
		initializerStatements += setValueField
	case "IsFictif":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", state.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsFictif")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", state.IsFictif))
		initializerStatements += setValueField
	case "IsEndState":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", state.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsEndState")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", state.IsEndState))
		initializerStatements += setValueField

	case "Parent":
		if state.Parent != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", state.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Parent")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", state.Parent.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "SubStates":
		for _, _state := range state.SubStates {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", state.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SubStates")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _state.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Diagrams":
		for _, _diagram := range state.Diagrams {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", state.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Diagrams")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _diagram.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Entry":
		if state.Entry != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", state.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Entry")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", state.Entry.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Activities":
		for _, _activities := range state.Activities {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", state.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Activities")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _activities.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Exit":
		if state.Exit != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", state.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Exit")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", state.Exit.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct State", fieldName)
	}
	return
}

func (statemachine *StateMachine) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", statemachine.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(statemachine.Name))
		initializerStatements += setValueField
	case "IsNodeExpanded":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", statemachine.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsNodeExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", statemachine.IsNodeExpanded))
		initializerStatements += setValueField

	case "States":
		for _, _state := range statemachine.States {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", statemachine.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "States")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _state.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Diagrams":
		for _, _diagram := range statemachine.Diagrams {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", statemachine.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Diagrams")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _diagram.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "InitialState":
		if statemachine.InitialState != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", statemachine.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "InitialState")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", statemachine.InitialState.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct StateMachine", fieldName)
	}
	return
}

func (stateshape *StateShape) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", stateshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(stateshape.Name))
		initializerStatements += setValueField
	case "IsExpanded":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", stateshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsExpanded")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", stateshape.IsExpanded))
		initializerStatements += setValueField
	case "X":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", stateshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", stateshape.X))
		initializerStatements += setValueField
	case "Y":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", stateshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", stateshape.Y))
		initializerStatements += setValueField
	case "Width":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", stateshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Width")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", stateshape.Width))
		initializerStatements += setValueField
	case "Height":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", stateshape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Height")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", stateshape.Height))
		initializerStatements += setValueField

	case "State":
		if stateshape.State != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", stateshape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "State")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", stateshape.State.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct StateShape", fieldName)
	}
	return
}

func (transition *Transition) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", transition.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(transition.Name))
		initializerStatements += setValueField

	case "Start":
		if transition.Start != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", transition.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Start")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", transition.Start.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "End":
		if transition.End != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", transition.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "End")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", transition.End.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "RolesWithPermissions":
		for _, _role := range transition.RolesWithPermissions {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", transition.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "RolesWithPermissions")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _role.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "GeneratedMessages":
		for _, _messagetype := range transition.GeneratedMessages {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", transition.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "GeneratedMessages")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _messagetype.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Guard":
		if transition.Guard != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", transition.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Guard")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", transition.Guard.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Diagrams":
		for _, _diagram := range transition.Diagrams {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", transition.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Diagrams")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _diagram.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct Transition", fieldName)
	}
	return
}

func (transition_shape *Transition_Shape) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", transition_shape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(transition_shape.Name))
		initializerStatements += setValueField
	case "StartRatio":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", transition_shape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", transition_shape.StartRatio))
		initializerStatements += setValueField
	case "EndRatio":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", transition_shape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", transition_shape.EndRatio))
		initializerStatements += setValueField
	case "StartOrientation":
		if transition_shape.StartOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", transition_shape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StartOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+transition_shape.StartOrientation.ToCodeString())
			initializerStatements += setValueField
		}
	case "EndOrientation":
		if transition_shape.EndOrientation != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", transition_shape.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "EndOrientation")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+transition_shape.EndOrientation.ToCodeString())
			initializerStatements += setValueField
		}
	case "CornerOffsetRatio":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", transition_shape.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", transition_shape.CornerOffsetRatio))
		initializerStatements += setValueField

	case "Transition":
		if transition_shape.Transition != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", transition_shape.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Transition")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", transition_shape.Transition.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct Transition_Shape", fieldName)
	}
	return
}
