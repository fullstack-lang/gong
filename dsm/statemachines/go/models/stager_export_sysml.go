package models

import (
	"encoding/base64"
	"fmt"
	"log"
	"regexp"
	"slices"
	"strings"
	"time"

	load "github.com/fullstack-lang/gong/lib/load/go/models"
)

var sysmlKeywords = map[string]bool{
	"package": true, "state": true, "def": true, "transition": true,
	"first": true, "then": true, "entry": true, "exit": true, "do": true,
	"action": true, "accept": true, "if": true, "item": true, "import": true,
	"part": true, "port": true, "attribute": true, "doc": true, "send": true,
	"done": true, "terminate": true, "return": true, "in": true, "out": true,
	"inout": true, "as": true, "alias": true, "private": true, "public": true,
	"protected": true, "abstract": true, "variation": true, "connect": true,
	"flow": true, "bind": true,
}

func toSysMLIdent(name string) string {
	if name == "" {
		return "''"
	}
	needQuotes := false
	if sysmlKeywords[name] {
		needQuotes = true
	} else {
		for i, r := range name {
			if i == 0 && (r >= '0' && r <= '9') {
				needQuotes = true
				break
			}
			if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_') {
				needQuotes = true
				break
			}
		}
	}
	if !needQuotes {
		return name
	}
	escaped := strings.ReplaceAll(name, "'", "\\'")
	return "'" + escaped + "'"
}

func getStatePath(state *State, parentMap map[*State]*State) string {
	if state == nil {
		return "''"
	}
	var parts []string
	curr := state
	for curr != nil {
		if curr.Name != "" {
			parts = append([]string{toSysMLIdent(curr.Name)}, parts...)
		}
		curr = parentMap[curr]
	}
	if len(parts) == 0 {
		return "''"
	}
	return strings.Join(parts, "::")
}

func writeState(sb *strings.Builder, state *State, indent string, parentMap map[*State]*State, visited map[*State]bool) {
	if state == nil || state.Name == "" || visited[state] {
		return
	}
	visited[state] = true

	ident := toSysMLIdent(state.Name)

	hasSubStates := false
	for _, sub := range state.SubStates {
		if sub != nil && sub.Name != "" {
			hasSubStates = true
			break
		}
	}

	hasBody := hasSubStates ||
		len(state.Notes) > 0 ||
		(state.Entry != nil && state.Entry.Name != "") ||
		len(state.Activities) > 0 ||
		(state.Exit != nil && state.Exit.Name != "")

	if !hasBody {
		if state.IsEndState {
			sb.WriteString(fmt.Sprintf("%sstate %s; // end state\n", indent, ident))
		} else {
			sb.WriteString(fmt.Sprintf("%sstate %s;\n", indent, ident))
		}
		return
	}

	sb.WriteString(fmt.Sprintf("%sstate %s {\n", indent, ident))
	innerIndent := indent + "    "

	if state.IsEndState {
		sb.WriteString(fmt.Sprintf("%s// end state\n", innerIndent))
	}

	for _, note := range state.Notes {
		if note != nil && note.Name != "" {
			sb.WriteString(fmt.Sprintf("%sdoc /* Note: %s */\n", innerIndent, strings.ReplaceAll(note.Name, "*/", "* /")))
		}
	}

	if state.Entry != nil && state.Entry.Name != "" {
		sb.WriteString(fmt.Sprintf("%sentry action %s;\n", innerIndent, toSysMLIdent(state.Entry.Name)))
	}

	for _, act := range state.Activities {
		if act != nil && act.Name != "" {
			sb.WriteString(fmt.Sprintf("%sdo action %s;\n", innerIndent, toSysMLIdent(act.Name)))
		}
	}

	if state.Exit != nil && state.Exit.Name != "" {
		sb.WriteString(fmt.Sprintf("%sexit action %s;\n", innerIndent, toSysMLIdent(state.Exit.Name)))
	}

	for _, sub := range state.SubStates {
		if sub != nil {
			writeState(sb, sub, innerIndent, parentMap, visited)
		}
	}

	sb.WriteString(fmt.Sprintf("%s}\n", indent))
}

func (stager *Stager) writeStateMachine(sb *strings.Builder, sm *StateMachine, indent string, allTransitions []*Transition) {
	if sm == nil {
		return
	}
	smName := sm.Name
	if smName == "" {
		smName = "StateMachine"
	}
	sb.WriteString(fmt.Sprintf("%sstate def %s {\n", indent, toSysMLIdent(smName)))
	innerIndent := indent + "    "

	// Build map of parent state for each substate
	parentMap := make(map[*State]*State)
	for _, state := range sm.States {
		if state == nil {
			continue
		}
		for _, sub := range state.SubStates {
			if sub != nil {
				parentMap[sub] = state
			}
		}
		if state.Parent != nil {
			parentMap[state] = state.Parent
		}
	}

	// State membership set for this SM
	smStates := make(map[*State]bool)
	for _, state := range sm.States {
		if state != nil {
			smStates[state] = true
		}
	}
	if sm.InitialState != nil {
		smStates[sm.InitialState] = true
	}

	// Identify transitions for this SM and the initial transition
	var initialTransition *Transition
	var smTransitions []*Transition

	for _, t := range allTransitions {
		if t == nil || t.Start == nil || t.End == nil {
			continue
		}
		// A transition belongs to this state machine if both Start and End belong to this SM
		if smStates[t.Start] && smStates[t.End] {
			if sm.InitialState != nil && t.Start == sm.InitialState && initialTransition == nil {
				initialTransition = t
			} else {
				smTransitions = append(smTransitions, t)
			}
		}
	}

	// Entry / initial transition
	if initialTransition != nil && initialTransition.End != nil {
		targetPath := getStatePath(initialTransition.End, parentMap)
		sb.WriteString(fmt.Sprintf("%sentry; then %s;\n\n", innerIndent, targetPath))
	} else if sm.InitialState != nil && sm.InitialState.Name != "" {
		sb.WriteString(fmt.Sprintf("%sentry; then %s;\n\n", innerIndent, getStatePath(sm.InitialState, parentMap)))
	}

	// Write states
	visited := make(map[*State]bool)
	// If InitialState is an anonymous pseudo-state or simple start marker, mark it as visited so it's not rendered as state '';
	if sm.InitialState != nil && (sm.InitialState.Name == "" || (initialTransition != nil && len(sm.InitialState.SubStates) == 0 && sm.InitialState.Entry == nil && sm.InitialState.Exit == nil && len(sm.InitialState.Activities) == 0)) {
		visited[sm.InitialState] = true
	}

	for _, state := range sm.States {
		if state == nil || state.Name == "" {
			continue
		}
		if parentMap[state] == nil && !visited[state] {
			writeState(sb, state, innerIndent, parentMap, visited)
		}
	}

	// Write transitions
	if len(smTransitions) > 0 {
		sb.WriteString("\n")
		for _, t := range smTransitions {
			nameStr := ""
			if t.Name != "" {
				nameStr = toSysMLIdent(t.Name) + " "
			}

			sourcePath := getStatePath(t.Start, parentMap)
			targetPath := getStatePath(t.End, parentMap)

			guardStr := ""
			if t.Guard != nil && t.Guard.Name != "" {
				guardStr = "if " + toSysMLIdent(t.Guard.Name) + " "
			}

			effectStr := ""
			for _, msg := range t.GeneratedMessages {
				if msg != nil && msg.Name != "" {
					effectStr += "do action " + toSysMLIdent(msg.Name) + " "
				}
			}

			rolesComment := ""
			if len(t.RolesWithPermissions) > 0 {
				var roleNames []string
				for _, r := range t.RolesWithPermissions {
					if r != nil && r.Name != "" {
						roleNames = append(roleNames, r.Name)
					}
				}
				if len(roleNames) > 0 {
					rolesComment = " // Roles: " + strings.Join(roleNames, ", ")
				}
			}

			sb.WriteString(fmt.Sprintf("%stransition %sfirst %s %s%sthen %s;%s\n",
				innerIndent, nameStr, sourcePath, guardStr, effectStr, targetPath, rolesComment))
		}
	}

	sb.WriteString(fmt.Sprintf("%s}\n", indent))
}

func (stager *Stager) writeLibraryPackage(sb *strings.Builder, lib *Library, indent string, allTransitions []*Transition) {
	if lib == nil {
		return
	}
	libName := lib.Name
	if libName == "" {
		libName = "Library"
	}
	sb.WriteString(fmt.Sprintf("%spackage %s {\n", indent, toSysMLIdent(libName)))
	innerIndent := indent + "    "

	// Gather state machines belonging to this library
	sms := lib.RootStateMachines
	if len(sms) == 0 && lib.IsRootLibrary {
		sms = GetGongstrucsSorted[*StateMachine](stager.stage)
	}

	// Collect message types used across these state machines
	usedMessageTypes := make(map[*MessageType]bool)
	smStateSet := make(map[*State]bool)
	for _, sm := range sms {
		if sm == nil {
			continue
		}
		for _, s := range sm.States {
			if s != nil {
				smStateSet[s] = true
			}
		}
		if sm.InitialState != nil {
			smStateSet[sm.InitialState] = true
		}
	}

	for _, t := range allTransitions {
		if t == nil || t.Start == nil || t.End == nil {
			continue
		}
		if smStateSet[t.Start] && smStateSet[t.End] {
			for _, msg := range t.GeneratedMessages {
				if msg != nil {
					usedMessageTypes[msg] = true
				}
			}
		}
	}

	// If root library, also include message types defined in stage if none gathered
	if lib.IsRootLibrary && len(usedMessageTypes) == 0 {
		for _, msg := range GetGongstrucsSorted[*MessageType](stager.stage) {
			usedMessageTypes[msg] = true
		}
	}

	var sortedMessageTypes []*MessageType
	for msg := range usedMessageTypes {
		sortedMessageTypes = append(sortedMessageTypes, msg)
	}
	slices.SortFunc(sortedMessageTypes, func(a, b *MessageType) int {
		return strings.Compare(a.Name, b.Name)
	})

	if len(sortedMessageTypes) > 0 {
		for _, msg := range sortedMessageTypes {
			if msg.Name == "" {
				continue
			}
			if msg.Description != "" {
				sb.WriteString(fmt.Sprintf("%sitem def %s {\n%s    doc /* %s */\n%s}\n",
					innerIndent, toSysMLIdent(msg.Name), innerIndent, strings.ReplaceAll(msg.Description, "*/", "* /"), innerIndent))
			} else {
				sb.WriteString(fmt.Sprintf("%sitem def %s;\n", innerIndent, toSysMLIdent(msg.Name)))
			}
		}
		sb.WriteString("\n")
	}



	// Write state machines
	for _, sm := range sms {
		if sm != nil {
			stager.writeStateMachine(sb, sm, innerIndent, allTransitions)
			sb.WriteString("\n")
		}
	}

	// Write sub-libraries recursively
	for _, subLib := range lib.SubLibraries {
		if subLib != nil {
			stager.writeLibraryPackage(sb, subLib, innerIndent, allTransitions)
			sb.WriteString("\n")
		}
	}

	sb.WriteString(fmt.Sprintf("%s}\n", indent))
}

func (stager *Stager) generateSysML(library *Library) string {
	var sb strings.Builder

	allTransitions := GetGongstrucsSorted[*Transition](stager.stage)
	stager.writeLibraryPackage(&sb, library, "", allTransitions)

	return sb.String()
}

func (stager *Stager) exportSysML(library *Library) {
	log.Println("Exporting SysML V2 format for library:", library.Name)

	sysmlContent := stager.generateSysML(library)

	stager.loadStage.Reset()

	cleanLibraryName := library.Name
	if cleanLibraryName == "" {
		cleanLibraryName = "statemachines"
	}
	reg := regexp.MustCompile(`[^a-zA-Z0-9_\-]+`)
	cleanLibraryName = reg.ReplaceAllString(cleanLibraryName, "_")

	fileName := time.Now().Format("20060102 1504 ") + cleanLibraryName + ".sysml"

	fileToDownload := &load.FileToDownload{
		Name:                 fileName,
		Base64EncodedContent: base64.StdEncoding.EncodeToString([]byte(sysmlContent)),
	}

	fileToUpload := &load.FileToUpload{
		Name: "Name of file",
		FileToUploadProxy: &loadProxy{
			stager: stager,
		},
	}

	load.StageBranch(stager.loadStage, fileToDownload)
	load.StageBranch(stager.loadStage, fileToUpload)

	message := &load.Message{
		Name: "Drop your <library>.go file here or ",
	}
	message.Stage(stager.loadStage)

	stager.loadStage.Commit()
}
