package models

import (
	"strings"
	"testing"
)

func TestToSysMLIdent(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Red", "Red"},
		{"Traffic Light UK", "'Traffic Light UK'"},
		{"state", "'state'"},
		{"package", "'package'"},
		{"123State", "'123State'"},
		{"State-A", "'State-A'"},
		{"State_B", "State_B"},
		{"State's", "'State\\'s'"},
		{"", "''"},
	}

	for _, tt := range tests {
		actual := toSysMLIdent(tt.input)
		if actual != tt.expected {
			t.Errorf("toSysMLIdent(%q) = %q; want %q", tt.input, actual, tt.expected)
		}
	}
}

func TestGetStatePath(t *testing.T) {
	parent := &State{Name: "On"}
	child := &State{Name: "Red", Parent: parent}
	grandchild := &State{Name: "SubRed", Parent: child}

	parentMap := map[*State]*State{
		child:      parent,
		grandchild: child,
	}

	if p := getStatePath(parent, parentMap); p != "On" {
		t.Errorf("expected 'On', got %q", p)
	}

	if p := getStatePath(child, parentMap); p != "On::Red" {
		t.Errorf("expected 'On::Red', got %q", p)
	}

	if p := getStatePath(grandchild, parentMap); p != "On::Red::SubRed" {
		t.Errorf("expected 'On::Red::SubRed', got %q", p)
	}
}

func TestGenerateSysML_TrafficLight(t *testing.T) {
	stage := NewStage("test")
	stager := &Stager{
		stage: stage,
	}

	// Create model elements
	lib := (&Library{Name: "Traffic Lights Library"}).Stage(stage)

	msgType := (&MessageType{Name: "Repair Report", Description: "Report sent when repairing"}).Stage(stage)

	roleTech := (&Role{Name: "Technician"}).Stage(stage)

	actWait := (&Activities{Name: "Wait for timer"}).Stage(stage)
	actFlash := (&Activities{Name: "Flash orange light"}).Stage(stage)
	entryAct := (&Action{Name: "PowerOn"}).Stage(stage)
	exitAct := (&Action{Name: "PowerOff"}).Stage(stage)

	initialState := (&State{Name: ""}).Stage(stage)
	stateRed := (&State{Name: "Red", Activities: []*Activities{actWait}}).Stage(stage)
	stateYellow := (&State{Name: "Yellow", Activities: []*Activities{actFlash, actWait}}).Stage(stage)
	stateGreen := (&State{Name: "Green", Activities: []*Activities{actWait}}).Stage(stage)

	stateOn := (&State{
		Name:      "On",
		Entry:     entryAct,
		Exit:      exitAct,
		SubStates: []*State{stateRed, stateYellow, stateGreen},
	}).Stage(stage)

	stateOff := (&State{Name: "Off", IsEndState: true}).Stage(stage)

	sm := (&StateMachine{
		Name:         "Traffic Light UK",
		InitialState: initialState,
		States:       []*State{initialState, stateOn, stateRed, stateYellow, stateGreen, stateOff},
	}).Stage(stage)

	lib.RootStateMachines = append(lib.RootStateMachines, sm)

	// Transitions
	// 1. Initial transition: initialState -> Red
	tInit := (&Transition{Start: initialState, End: stateRed}).Stage(stage)
	_ = tInit

	// 2. Red -> Yellow
	(&Transition{
		Name:                 "T_Red_Yellow",
		Start:                stateRed,
		End:                  stateYellow,
		RolesWithPermissions: []*Role{roleTech},
	}).Stage(stage)

	// 3. Yellow -> Green
	(&Transition{
		Name:  "T_Yellow_Green",
		Start: stateYellow,
		End:   stateGreen,
	}).Stage(stage)

	// 4. On -> Off
	(&Transition{
		Start: stateOn,
		End:   stateOff,
		Guard: &Guard{Name: "PowerFailure"},
	}).Stage(stage)

	// 5. Off -> Red (with generated message)
	(&Transition{
		Start:             stateOff,
		End:               stateRed,
		GeneratedMessages: []*MessageType{msgType},
	}).Stage(stage)

	sysml := stager.generateSysML(lib)

	t.Logf("Generated SysML:\n%s", sysml)

	// Verify package declaration
	if !strings.Contains(sysml, "package 'Traffic Lights Library' {") {
		t.Errorf("Expected package declaration, got:\n%s", sysml)
	}

	// Verify item def for MessageType with doc
	if !strings.Contains(sysml, "item def 'Repair Report' {") || !strings.Contains(sysml, "doc /* Report sent when repairing */") {
		t.Errorf("Expected item def with doc for Repair Report, got:\n%s", sysml)
	}

	// Verify state def
	if !strings.Contains(sysml, "state def 'Traffic Light UK' {") {
		t.Errorf("Expected state def 'Traffic Light UK', got:\n%s", sysml)
	}

	// Verify entry transition to On::Red
	if !strings.Contains(sysml, "entry; then On::Red;") {
		t.Errorf("Expected entry; then On::Red;, got:\n%s", sysml)
	}

	// Verify composite state On with entry, exit, substates and actions
	if !strings.Contains(sysml, "state On {") {
		t.Errorf("Expected state On {, got:\n%s", sysml)
	}
	if !strings.Contains(sysml, "entry action PowerOn;") {
		t.Errorf("Expected entry action PowerOn;, got:\n%s", sysml)
	}
	if !strings.Contains(sysml, "exit action PowerOff;") {
		t.Errorf("Expected exit action PowerOff;, got:\n%s", sysml)
	}
	if !strings.Contains(sysml, "state Red {") || !strings.Contains(sysml, "do action 'Wait for timer';") {
		t.Errorf("Expected state Red with do action, got:\n%s", sysml)
	}
	if !strings.Contains(sysml, "state Yellow {") || !strings.Contains(sysml, "do action 'Flash orange light';") {
		t.Errorf("Expected state Yellow with do action, got:\n%s", sysml)
	}

	// Verify end state Off
	if !strings.Contains(sysml, "state Off; // end state") {
		t.Errorf("Expected state Off; // end state, got:\n%s", sysml)
	}

	// Verify transitions
	if !strings.Contains(sysml, "transition T_Red_Yellow first On::Red then On::Yellow; // Roles: Technician") {
		t.Errorf("Expected transition T_Red_Yellow with role comment, got:\n%s", sysml)
	}
	if !strings.Contains(sysml, "transition first On if PowerFailure then Off;") {
		t.Errorf("Expected guarded transition from On to Off, got:\n%s", sysml)
	}
	if !strings.Contains(sysml, "transition first Off do action 'Repair Report' then On::Red;") {
		t.Errorf("Expected transition with action, got:\n%s", sysml)
	}

	// Ensure anonymous pseudo-state is not declared
	if strings.Contains(sysml, "state '';") || strings.Contains(sysml, "state '' {") {
		t.Errorf("Anonymous pseudo-state should not be declared as state '';")
	}
}
