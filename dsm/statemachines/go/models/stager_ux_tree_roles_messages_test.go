package models

import (
	"slices"
	"strings"
	"testing"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func TestRolesAndMessagesInTreeAndSvg(t *testing.T) {
	stage := NewStage("test")
	stager := &Stager{
		stage: stage,
	}

	// 1. Create root library with roles and messages
	lib := (&Library{
		Name:          "Root Lib",
		IsRootLibrary: true,
	}).Stage(stage)

	r1 := (&Role{Name: "Role1", Acronym: "R1"}).Stage(stage)
	m1 := (&MessageType{Name: "Msg1"}).Stage(stage)

	lib.Roles = append(lib.Roles, r1)
	lib.MessageTypes = append(lib.MessageTypes, m1)

	// 2. Test treeLibrary creates Roles and Messages nodes
	var rootNodes []*tree.Node
	stager.treeLibrary(lib, &rootNodes)

	if len(rootNodes) == 0 {
		t.Fatalf("expected at least one root node, got 0")
	}
	libNode := rootNodes[0]

	var rolesNode *tree.Node
	var messagesNode *tree.Node
	for _, child := range libNode.Children {
		if child.Name == "Roles" {
			rolesNode = child
		}
		if child.Name == "Messages" {
			messagesNode = child
		}
	}

	if rolesNode == nil {
		t.Fatalf("expected 'Roles' node under library")
	}
	if len(rolesNode.Children) != 1 || rolesNode.Children[0].Name != "Role1" {
		t.Fatalf("expected Role1 under Roles node")
	}

	if messagesNode == nil {
		t.Fatalf("expected 'Messages' node under library")
	}
	if len(messagesNode.Children) != 1 || messagesNode.Children[0].Name != "Msg1" {
		t.Fatalf("expected Msg1 under Messages node")
	}

	// 3. Test SVG link generation respects ShowRoles and ShowMessages
	s1 := (&State{Name: "State1"}).Stage(stage)
	s2 := (&State{Name: "State2"}).Stage(stage)
	trans := (&Transition{
		Name:                 "T1",
		Start:                s1,
		End:                  s2,
		RolesWithPermissions: []*Role{r1},
		GeneratedMessages:    []*MessageType{m1},
	}).Stage(stage)

	diagram := (&Diagram{
		Name:         "Diagram1",
		ShowRoles:    true,
		ShowMessages: true,
	}).Stage(stage)

	startRect := &svg.Rect{Name: "R1"}
	endRect := &svg.Rect{Name: "R2"}
	linkShape := &LinkShape{}
	layer := &svg.Layer{}

	// When both are true: roles and messages should be present
	stager.svgGenerateLink(diagram, startRect, endRect, linkShape, trans, layer, false)
	if len(layer.Links) != 1 {
		t.Fatalf("expected 1 link, got %d", len(layer.Links))
	}
	link := layer.Links[0]
	if !strings.Contains(link.TextAtArrowStart[0].Content, "/R1") {
		t.Errorf("expected role /R1 in link label, got: %s", link.TextAtArrowStart[0].Content)
	}
	if len(link.TextAtArrowStart) < 2 || !strings.Contains(link.TextAtArrowStart[1].Content, "Msg1") {
		t.Errorf("expected Msg1 in anchored text, got: %v", link.TextAtArrowStart)
	}

	// When ShowRoles is false: roles should NOT be in link label
	diagram.ShowRoles = false
	layer.Links = nil
	stager.svgGenerateLink(diagram, startRect, endRect, linkShape, trans, layer, false)
	link = layer.Links[0]
	if strings.Contains(link.TextAtArrowStart[0].Content, "/R1") {
		t.Errorf("expected role /R1 to be hidden when ShowRoles=false, got: %s", link.TextAtArrowStart[0].Content)
	}
	if len(link.TextAtArrowStart) < 2 || !strings.Contains(link.TextAtArrowStart[1].Content, "Msg1") {
		t.Errorf("expected Msg1 still present when ShowMessages=true")
	}

	// When ShowMessages is false: messages should NOT be in anchored text
	diagram.ShowMessages = false
	layer.Links = nil
	stager.svgGenerateLink(diagram, startRect, endRect, linkShape, trans, layer, false)
	link = layer.Links[0]
	if len(link.TextAtArrowStart) != 1 {
		t.Errorf("expected only 1 text item (no message) when ShowMessages=false, got %d", len(link.TextAtArrowStart))
	}

	// 4. Test treeStateMachines generates toggle buttons on diagramNode
	sm := (&StateMachine{
		Name:     "SM1",
		Diagrams: []*Diagram{diagram},
		States:   []*State{s1, s2},
	}).Stage(stage)
	lib.RootStateMachines = []*StateMachine{sm}

	var parentNode tree.Node
	var expandedSM []*StateMachine
	stager.treeStateMachines(sm, &parentNode, &expandedSM)

	if len(parentNode.Children) == 0 {
		t.Fatalf("expected state machine node")
	}
	smNode := parentNode.Children[0]
	var diagNode *tree.Node
	for _, child := range smNode.Children {
		if child.Name == "Diagram1" {
			diagNode = child
			break
		}
	}
	if diagNode == nil {
		t.Fatalf("expected Diagram1 node under state machine")
	}

	hasRolesToggle := false
	hasMessagesToggle := false
	for _, btn := range diagNode.Buttons {
		if btn.Name == "Show/Hide Roles" {
			hasRolesToggle = true
		}
		if btn.Name == "Show/Hide Messages" {
			hasMessagesToggle = true
		}
	}
	if !hasRolesToggle {
		t.Errorf("expected Show/Hide Roles button on diagram node")
	}
	if !hasMessagesToggle {
		t.Errorf("expected Show/Hide Messages button on diagram node")
	}

	// 5. Test transitionNode has Roles and Messages sub-nodes
	// Add state shape so state is checked and transitions are listed
	diagNode.Children = nil
	diagram.State_Shapes = append(diagram.State_Shapes, &StateShape{
		Name:  "State1-Diagram1",
		State: s1,
	})
	parentNode.Children = nil
	stager.treeStateMachines(sm, &parentNode, &expandedSM)

	smNode = parentNode.Children[0]
	for _, child := range smNode.Children {
		if child.Name == "Diagram1" {
			diagNode = child
			break
		}
	}

	var statesFolder *tree.Node
	for _, child := range diagNode.Children {
		if child.Name == "States" {
			statesFolder = child
			break
		}
	}
	if statesFolder == nil || len(statesFolder.Children) == 0 {
		t.Fatalf("expected States folder with children")
	}
	s1Node := statesFolder.Children[0]
	var transNode *tree.Node
	for _, child := range s1Node.Children {
		if strings.HasPrefix(child.Name, "T1") {
			transNode = child
			break
		}
	}
	if transNode == nil {
		t.Fatalf("expected T1 transition node under state 1")
	}

	var transRolesNode *tree.Node
	var transMessagesNode *tree.Node
	for _, child := range transNode.Children {
		if child.Name == "Roles" {
			transRolesNode = child
		}
		if child.Name == "Messages" {
			transMessagesNode = child
		}
	}
	if transRolesNode == nil {
		t.Fatalf("expected Roles folder under transition node")
	}
	if len(transRolesNode.Children) != 1 || !transRolesNode.Children[0].IsChecked {
		t.Errorf("expected Role1 to be checked under transition Roles")
	}
	if transMessagesNode == nil {
		t.Fatalf("expected Messages folder under transition node")
	}
	if len(transMessagesNode.Children) != 1 || !transMessagesNode.Children[0].IsChecked {
		t.Errorf("expected Msg1 to be checked under transition Messages")
	}

	// Test unchecking role removes it from transition
	transRolesNode.Children[0].OnIsCheckedChanged(false)
	if slices.Contains(trans.RolesWithPermissions, r1) {
		t.Errorf("expected role to be removed from transition after unchecking")
	}

	// Test unchecking message removes it from transition
	transMessagesNode.Children[0].OnIsCheckedChanged(false)
	if slices.Contains(trans.GeneratedMessages, m1) {
		t.Errorf("expected message to be removed from transition after unchecking")
	}
}
