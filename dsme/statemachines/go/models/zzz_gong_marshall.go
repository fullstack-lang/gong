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
		identifiersDecl.WriteString("\n")
	}
	for _, action := range actionOrdered {

		identifiersDecl.WriteString(action.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(action.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(action.GongMarshallField(stage, "Criticality"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, activities := range activitiesOrdered {

		identifiersDecl.WriteString(activities.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(activities.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(activities.GongMarshallField(stage, "Criticality"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, architecture := range architectureOrdered {

		identifiersDecl.WriteString(architecture.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(architecture.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(architecture.GongMarshallField(stage, "StateMachines"))
		pointersInitializesStatements.WriteString(architecture.GongMarshallField(stage, "Roles"))
		initializerStatements.WriteString(architecture.GongMarshallField(stage, "NbPixPerCharacter"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, diagram := range diagramOrdered {

		identifiersDecl.WriteString(diagram.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsChecked"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsEditable_"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsInRenameMode"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "State_Shapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "Transition_Shapes"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, guard := range guardOrdered {

		identifiersDecl.WriteString(guard.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(guard.GongMarshallField(stage, "Name"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, kill := range killOrdered {

		identifiersDecl.WriteString(kill.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(kill.GongMarshallField(stage, "Name"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, message := range messageOrdered {

		identifiersDecl.WriteString(message.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(message.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(message.GongMarshallField(stage, "IsSelected"))
		pointersInitializesStatements.WriteString(message.GongMarshallField(stage, "MessageType"))
		pointersInitializesStatements.WriteString(message.GongMarshallField(stage, "OriginTransition"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, messagetype := range messagetypeOrdered {

		identifiersDecl.WriteString(messagetype.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(messagetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(messagetype.GongMarshallField(stage, "Description"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, object := range objectOrdered {

		identifiersDecl.WriteString(object.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(object.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(object.GongMarshallField(stage, "State"))
		initializerStatements.WriteString(object.GongMarshallField(stage, "IsSelected"))
		initializerStatements.WriteString(object.GongMarshallField(stage, "Rank"))
		initializerStatements.WriteString(object.GongMarshallField(stage, "DOF"))
		pointersInitializesStatements.WriteString(object.GongMarshallField(stage, "Messages"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, role := range roleOrdered {

		identifiersDecl.WriteString(role.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(role.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(role.GongMarshallField(stage, "Acronym"))
		pointersInitializesStatements.WriteString(role.GongMarshallField(stage, "RolesWithSamePermissions"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, state := range stateOrdered {

		identifiersDecl.WriteString(state.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(state.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(state.GongMarshallField(stage, "Parent"))
		initializerStatements.WriteString(state.GongMarshallField(stage, "IsDecisionNode"))
		initializerStatements.WriteString(state.GongMarshallField(stage, "IsFictif"))
		initializerStatements.WriteString(state.GongMarshallField(stage, "IsEndState"))
		pointersInitializesStatements.WriteString(state.GongMarshallField(stage, "SubStates"))
		pointersInitializesStatements.WriteString(state.GongMarshallField(stage, "Diagrams"))
		pointersInitializesStatements.WriteString(state.GongMarshallField(stage, "Entry"))
		pointersInitializesStatements.WriteString(state.GongMarshallField(stage, "Activities"))
		pointersInitializesStatements.WriteString(state.GongMarshallField(stage, "Exit"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, statemachine := range statemachineOrdered {

		identifiersDecl.WriteString(statemachine.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(statemachine.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(statemachine.GongMarshallField(stage, "IsNodeExpanded"))
		pointersInitializesStatements.WriteString(statemachine.GongMarshallField(stage, "States"))
		pointersInitializesStatements.WriteString(statemachine.GongMarshallField(stage, "Diagrams"))
		pointersInitializesStatements.WriteString(statemachine.GongMarshallField(stage, "InitialState"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, stateshape := range stateshapeOrdered {

		identifiersDecl.WriteString(stateshape.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(stateshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(stateshape.GongMarshallField(stage, "State"))
		initializerStatements.WriteString(stateshape.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(stateshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(stateshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(stateshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(stateshape.GongMarshallField(stage, "Height"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, transition := range transitionOrdered {

		identifiersDecl.WriteString(transition.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(transition.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(transition.GongMarshallField(stage, "Start"))
		pointersInitializesStatements.WriteString(transition.GongMarshallField(stage, "End"))
		pointersInitializesStatements.WriteString(transition.GongMarshallField(stage, "RolesWithPermissions"))
		pointersInitializesStatements.WriteString(transition.GongMarshallField(stage, "GeneratedMessages"))
		pointersInitializesStatements.WriteString(transition.GongMarshallField(stage, "Guard"))
		pointersInitializesStatements.WriteString(transition.GongMarshallField(stage, "Diagrams"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, transition_shape := range transition_shapeOrdered {

		identifiersDecl.WriteString(transition_shape.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(transition_shape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(transition_shape.GongMarshallField(stage, "Transition"))
		initializerStatements.WriteString(transition_shape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(transition_shape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(transition_shape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(transition_shape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(transition_shape.GongMarshallField(stage, "CornerOffsetRatio"))
	}

	// insertion initialization of objects to stage
	for _, action := range actionOrdered {
		_ = action
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, activities := range activitiesOrdered {
		_ = activities
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, architecture := range architectureOrdered {
		_ = architecture
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, diagram := range diagramOrdered {
		_ = diagram
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, guard := range guardOrdered {
		_ = guard
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, kill := range killOrdered {
		_ = kill
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, message := range messageOrdered {
		_ = message
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, messagetype := range messagetypeOrdered {
		_ = messagetype
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, object := range objectOrdered {
		_ = object
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, role := range roleOrdered {
		_ = role
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, state := range stateOrdered {
		_ = state
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, statemachine := range statemachineOrdered {
		_ = statemachine
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, stateshape := range stateshapeOrdered {
		_ = stateshape
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, transition := range transitionOrdered {
		_ = transition
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, transition_shape := range transition_shapeOrdered {
		_ = transition_shape
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
func (action *Action) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", action.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(action.Name))
	case "Criticality":
		if action.Criticality.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", action.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Criticality")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+action.Criticality.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", action.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Criticality")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct Action", fieldName)
	}
	return
}

func (activities *Activities) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", activities.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(activities.Name))
	case "Criticality":
		if activities.Criticality.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", activities.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Criticality")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+activities.Criticality.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", activities.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Criticality")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct Activities", fieldName)
	}
	return
}

func (architecture *Architecture) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", architecture.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(architecture.Name))
	case "NbPixPerCharacter":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", architecture.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NbPixPerCharacter")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", architecture.NbPixPerCharacter))

	case "StateMachines":
		var sb strings.Builder
		for _, _statemachine := range architecture.StateMachines {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", architecture.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "StateMachines")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _statemachine.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "Roles":
		var sb strings.Builder
		for _, _role := range architecture.Roles {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", architecture.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Roles")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _role.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Architecture", fieldName)
	}
	return
}

func (diagram *Diagram) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(diagram.Name))
	case "IsChecked":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsChecked")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsChecked))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsExpanded))
	case "IsEditable_":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsEditable_")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsEditable_))
	case "IsInRenameMode":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", diagram.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsInRenameMode")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", diagram.IsInRenameMode))

	case "State_Shapes":
		var sb strings.Builder
		for _, _stateshape := range diagram.State_Shapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "State_Shapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _stateshape.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "Transition_Shapes":
		var sb strings.Builder
		for _, _transition_shape := range diagram.Transition_Shapes {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", diagram.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Transition_Shapes")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _transition_shape.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Diagram", fieldName)
	}
	return
}

func (guard *Guard) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", guard.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(guard.Name))

	default:
		log.Panicf("Unknown field %s for Gongstruct Guard", fieldName)
	}
	return
}

func (kill *Kill) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", kill.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(kill.Name))

	default:
		log.Panicf("Unknown field %s for Gongstruct Kill", fieldName)
	}
	return
}

func (message *Message) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", message.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(message.Name))
	case "IsSelected":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", message.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", message.IsSelected))

	case "MessageType":
		if message.MessageType != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", message.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MessageType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", message.MessageType.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", message.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MessageType")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "OriginTransition":
		if message.OriginTransition != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", message.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "OriginTransition")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", message.OriginTransition.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", message.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "OriginTransition")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Message", fieldName)
	}
	return
}

func (messagetype *MessageType) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", messagetype.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(messagetype.Name))
	case "Description":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", messagetype.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Description")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(messagetype.Description))

	default:
		log.Panicf("Unknown field %s for Gongstruct MessageType", fieldName)
	}
	return
}

func (object *Object) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", object.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(object.Name))
	case "IsSelected":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", object.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsSelected")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", object.IsSelected))
	case "Rank":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", object.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Rank")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", object.Rank))
	case "DOF":
		res = TimeInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", object.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DOF")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", object.DOF.String())

	case "State":
		if object.State != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", object.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "State")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", object.State.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", object.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "State")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Messages":
		var sb strings.Builder
		for _, _message := range object.Messages {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", object.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Messages")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _message.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Object", fieldName)
	}
	return
}

func (role *Role) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", role.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(role.Name))
	case "Acronym":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", role.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Acronym")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(role.Acronym))

	case "RolesWithSamePermissions":
		var sb strings.Builder
		for _, _role := range role.RolesWithSamePermissions {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", role.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "RolesWithSamePermissions")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _role.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Role", fieldName)
	}
	return
}

func (state *State) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", state.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(state.Name))
	case "IsDecisionNode":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", state.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDecisionNode")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", state.IsDecisionNode))
	case "IsFictif":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", state.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsFictif")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", state.IsFictif))
	case "IsEndState":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", state.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsEndState")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", state.IsEndState))

	case "Parent":
		if state.Parent != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", state.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Parent")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", state.Parent.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", state.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Parent")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SubStates":
		var sb strings.Builder
		for _, _state := range state.SubStates {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", state.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SubStates")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _state.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "Diagrams":
		var sb strings.Builder
		for _, _diagram := range state.Diagrams {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", state.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Diagrams")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _diagram.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "Entry":
		if state.Entry != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", state.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Entry")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", state.Entry.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", state.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Entry")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Activities":
		var sb strings.Builder
		for _, _activities := range state.Activities {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", state.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Activities")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _activities.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "Exit":
		if state.Exit != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", state.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Exit")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", state.Exit.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", state.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Exit")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct State", fieldName)
	}
	return
}

func (statemachine *StateMachine) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", statemachine.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(statemachine.Name))
	case "IsNodeExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", statemachine.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsNodeExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", statemachine.IsNodeExpanded))

	case "States":
		var sb strings.Builder
		for _, _state := range statemachine.States {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", statemachine.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "States")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _state.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "Diagrams":
		var sb strings.Builder
		for _, _diagram := range statemachine.Diagrams {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", statemachine.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Diagrams")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _diagram.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "InitialState":
		if statemachine.InitialState != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", statemachine.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InitialState")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", statemachine.InitialState.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", statemachine.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "InitialState")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct StateMachine", fieldName)
	}
	return
}

func (stateshape *StateShape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", stateshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(stateshape.Name))
	case "IsExpanded":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", stateshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsExpanded")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", stateshape.IsExpanded))
	case "X":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", stateshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", stateshape.X))
	case "Y":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", stateshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", stateshape.Y))
	case "Width":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", stateshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Width")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", stateshape.Width))
	case "Height":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", stateshape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Height")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", stateshape.Height))

	case "State":
		if stateshape.State != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", stateshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "State")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", stateshape.State.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", stateshape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "State")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct StateShape", fieldName)
	}
	return
}

func (transition *Transition) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", transition.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(transition.Name))

	case "Start":
		if transition.Start != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", transition.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Start")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", transition.Start.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", transition.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Start")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "End":
		if transition.End != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", transition.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "End")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", transition.End.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", transition.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "End")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "RolesWithPermissions":
		var sb strings.Builder
		for _, _role := range transition.RolesWithPermissions {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", transition.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "RolesWithPermissions")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _role.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "GeneratedMessages":
		var sb strings.Builder
		for _, _messagetype := range transition.GeneratedMessages {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", transition.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "GeneratedMessages")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _messagetype.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "Guard":
		if transition.Guard != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", transition.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Guard")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", transition.Guard.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", transition.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Guard")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "Diagrams":
		var sb strings.Builder
		for _, _diagram := range transition.Diagrams {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", transition.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Diagrams")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _diagram.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Transition", fieldName)
	}
	return
}

func (transition_shape *Transition_Shape) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", transition_shape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(transition_shape.Name))
	case "StartRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", transition_shape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", transition_shape.StartRatio))
	case "EndRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", transition_shape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", transition_shape.EndRatio))
	case "StartOrientation":
		if transition_shape.StartOrientation.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", transition_shape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+transition_shape.StartOrientation.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", transition_shape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StartOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "EndOrientation":
		if transition_shape.EndOrientation.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", transition_shape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+transition_shape.EndOrientation.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", transition_shape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EndOrientation")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "CornerOffsetRatio":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", transition_shape.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CornerOffsetRatio")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", transition_shape.CornerOffsetRatio))

	case "Transition":
		if transition_shape.Transition != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", transition_shape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transition")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", transition_shape.Transition.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", transition_shape.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Transition")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Transition_Shape", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (action *Action) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(action.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(action.GongMarshallField(stage, "Criticality"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (activities *Activities) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(activities.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(activities.GongMarshallField(stage, "Criticality"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (architecture *Architecture) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(architecture.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(architecture.GongMarshallField(stage, "StateMachines"))
		pointersInitializesStatements.WriteString(architecture.GongMarshallField(stage, "Roles"))
		initializerStatements.WriteString(architecture.GongMarshallField(stage, "NbPixPerCharacter"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (diagram *Diagram) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsChecked"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsEditable_"))
		initializerStatements.WriteString(diagram.GongMarshallField(stage, "IsInRenameMode"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "State_Shapes"))
		pointersInitializesStatements.WriteString(diagram.GongMarshallField(stage, "Transition_Shapes"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (guard *Guard) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(guard.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (kill *Kill) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(kill.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (message *Message) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(message.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(message.GongMarshallField(stage, "IsSelected"))
		pointersInitializesStatements.WriteString(message.GongMarshallField(stage, "MessageType"))
		pointersInitializesStatements.WriteString(message.GongMarshallField(stage, "OriginTransition"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (messagetype *MessageType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(messagetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(messagetype.GongMarshallField(stage, "Description"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (object *Object) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(object.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(object.GongMarshallField(stage, "State"))
		initializerStatements.WriteString(object.GongMarshallField(stage, "IsSelected"))
		initializerStatements.WriteString(object.GongMarshallField(stage, "Rank"))
		initializerStatements.WriteString(object.GongMarshallField(stage, "DOF"))
		pointersInitializesStatements.WriteString(object.GongMarshallField(stage, "Messages"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (role *Role) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(role.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(role.GongMarshallField(stage, "Acronym"))
		pointersInitializesStatements.WriteString(role.GongMarshallField(stage, "RolesWithSamePermissions"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (state *State) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(state.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(state.GongMarshallField(stage, "Parent"))
		initializerStatements.WriteString(state.GongMarshallField(stage, "IsDecisionNode"))
		initializerStatements.WriteString(state.GongMarshallField(stage, "IsFictif"))
		initializerStatements.WriteString(state.GongMarshallField(stage, "IsEndState"))
		pointersInitializesStatements.WriteString(state.GongMarshallField(stage, "SubStates"))
		pointersInitializesStatements.WriteString(state.GongMarshallField(stage, "Diagrams"))
		pointersInitializesStatements.WriteString(state.GongMarshallField(stage, "Entry"))
		pointersInitializesStatements.WriteString(state.GongMarshallField(stage, "Activities"))
		pointersInitializesStatements.WriteString(state.GongMarshallField(stage, "Exit"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (statemachine *StateMachine) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(statemachine.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(statemachine.GongMarshallField(stage, "IsNodeExpanded"))
		pointersInitializesStatements.WriteString(statemachine.GongMarshallField(stage, "States"))
		pointersInitializesStatements.WriteString(statemachine.GongMarshallField(stage, "Diagrams"))
		pointersInitializesStatements.WriteString(statemachine.GongMarshallField(stage, "InitialState"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (stateshape *StateShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(stateshape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(stateshape.GongMarshallField(stage, "State"))
		initializerStatements.WriteString(stateshape.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(stateshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(stateshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(stateshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(stateshape.GongMarshallField(stage, "Height"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (transition *Transition) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(transition.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(transition.GongMarshallField(stage, "Start"))
		pointersInitializesStatements.WriteString(transition.GongMarshallField(stage, "End"))
		pointersInitializesStatements.WriteString(transition.GongMarshallField(stage, "RolesWithPermissions"))
		pointersInitializesStatements.WriteString(transition.GongMarshallField(stage, "GeneratedMessages"))
		pointersInitializesStatements.WriteString(transition.GongMarshallField(stage, "Guard"))
		pointersInitializesStatements.WriteString(transition.GongMarshallField(stage, "Diagrams"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (transition_shape *Transition_Shape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(transition_shape.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(transition_shape.GongMarshallField(stage, "Transition"))
		initializerStatements.WriteString(transition_shape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(transition_shape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(transition_shape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(transition_shape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(transition_shape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
