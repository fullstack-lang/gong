package models

import (
	"strings"
	"testing"
	"time"
)

type mockProbe struct {
	ProbeIF
	notifications []string
}

func (m *mockProbe) AddNotification(date time.Time, message string) {
	m.notifications = append(m.notifications, message)
}

func TestEnforceStateMachineSemanticRules_OneClick(t *testing.T) {
	t.Run("single connected component (one click)", func(t *testing.T) {
		stage := NewStage("test")
		probe := &mockProbe{}
		stager := &Stager{
			stage:     stage,
			probeForm: probe,
		}

		sm := (&StateMachine{Name: "TrafficLight"}).Stage(stage)
		s1 := (&State{Name: "Red"}).Stage(stage)
		s2 := (&State{Name: "Yellow"}).Stage(stage)
		s3 := (&State{Name: "Green"}).Stage(stage)

		sm.InitialState = s1
		sm.States = append(sm.States, s1, s2, s3)

		t1 := (&Transition{Name: "T1", Start: s1, End: s2}).Stage(stage)
		t2 := (&Transition{Name: "T2", Start: s2, End: s3}).Stage(stage)
		_ = t1
		_ = t2

		stager.enforceStagerMaps()
		stager.enforceStateMachineSemanticRules()

		for _, n := range probe.notifications {
			if strings.Contains(n, "state machine graph has to be one click") {
				t.Fatalf("unexpected notification: %s", n)
			}
		}
	})

	t.Run("two disconnected components (two cliques)", func(t *testing.T) {
		stage := NewStage("test")
		probe := &mockProbe{}
		stager := &Stager{
			stage:     stage,
			probeForm: probe,
		}

		sm := (&StateMachine{Name: "TwoCliquesSM"}).Stage(stage)
		s1 := (&State{Name: "S1"}).Stage(stage)
		s2 := (&State{Name: "S2"}).Stage(stage)
		s3 := (&State{Name: "S3"}).Stage(stage)
		s4 := (&State{Name: "S4"}).Stage(stage)

		sm.InitialState = s1
		sm.States = append(sm.States, s1, s2, s3, s4)

		// Clique 1: S1 -> S2
		_ = (&Transition{Name: "T1", Start: s1, End: s2}).Stage(stage)
		// Clique 2: S3 -> S4
		_ = (&Transition{Name: "T2", Start: s3, End: s4}).Stage(stage)

		stager.enforceStagerMaps()
		stager.enforceStateMachineSemanticRules()

		found := false
		foundCount := false
		for _, n := range probe.notifications {
			if strings.Contains(n, "state machine graph has to be one click") {
				found = true
				if strings.Contains(n, "computed 2 clicks") {
					foundCount = true
				}
				break
			}
		}
		if !found {
			t.Fatalf("expected 'state machine graph has to be one click' notification, got: %v", probe.notifications)
		}
		if !foundCount {
			t.Fatalf("expected notification to contain 'computed 2 clicks', got: %v", probe.notifications)
		}
	})

	t.Run("isolated state (not one click)", func(t *testing.T) {
		stage := NewStage("test")
		probe := &mockProbe{}
		stager := &Stager{
			stage:     stage,
			probeForm: probe,
		}

		sm := (&StateMachine{Name: "IsolatedStateSM"}).Stage(stage)
		s1 := (&State{Name: "S1"}).Stage(stage)
		s2 := (&State{Name: "S2"}).Stage(stage)
		s3 := (&State{Name: "S3_Isolated"}).Stage(stage)

		sm.InitialState = s1
		sm.States = append(sm.States, s1, s2, s3)

		_ = (&Transition{Name: "T1", Start: s1, End: s2}).Stage(stage)

		stager.enforceStagerMaps()
		stager.enforceStateMachineSemanticRules()

		found := false
		for _, n := range probe.notifications {
			if strings.Contains(n, "state machine graph has to be one click") {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("expected 'state machine graph has to be one click' notification, got: %v", probe.notifications)
		}
	})

	t.Run("composite state hierarchy (one click)", func(t *testing.T) {
		stage := NewStage("test")
		probe := &mockProbe{}
		stager := &Stager{
			stage:     stage,
			probeForm: probe,
		}

		sm := (&StateMachine{Name: "CompositeSM"}).Stage(stage)
		sParent := (&State{Name: "Parent"}).Stage(stage)
		s1 := (&State{Name: "Sub1"}).Stage(stage)
		s2 := (&State{Name: "Sub2"}).Stage(stage)

		sParent.SubStates = append(sParent.SubStates, s1, s2)
		sm.InitialState = s1
		sm.States = append(sm.States, sParent, s1, s2)

		_ = (&Transition{Name: "T1", Start: s1, End: s2}).Stage(stage)

		stager.enforceStagerMaps()
		stager.enforceStateMachineSemanticRules()

		for _, n := range probe.notifications {
			if strings.Contains(n, "state machine graph has to be one click") {
				t.Fatalf("unexpected notification for composite state: %s", n)
			}
		}
	})

	t.Run("multiple state machines independent check", func(t *testing.T) {
		stage := NewStage("test")
		probe := &mockProbe{}
		stager := &Stager{
			stage:     stage,
			probeForm: probe,
		}

		// SM1 is connected (1 click)
		sm1 := (&StateMachine{Name: "SM1"}).Stage(stage)
		s1 := (&State{Name: "S1"}).Stage(stage)
		s2 := (&State{Name: "S2"}).Stage(stage)
		sm1.InitialState = s1
		sm1.States = append(sm1.States, s1, s2)
		_ = (&Transition{Name: "T1", Start: s1, End: s2}).Stage(stage)

		// SM2 has 2 cliques (2 disconnected components)
		sm2 := (&StateMachine{Name: "SM2"}).Stage(stage)
		s3 := (&State{Name: "S3"}).Stage(stage)
		s4 := (&State{Name: "S4"}).Stage(stage)
		s5 := (&State{Name: "S5"}).Stage(stage)
		s6 := (&State{Name: "S6"}).Stage(stage)
		sm2.InitialState = s3
		sm2.States = append(sm2.States, s3, s4, s5, s6)
		_ = (&Transition{Name: "T2", Start: s3, End: s4}).Stage(stage)
		_ = (&Transition{Name: "T3", Start: s5, End: s6}).Stage(stage)

		stager.enforceStagerMaps()
		stager.enforceStateMachineSemanticRules()

		sm1Failed := false
		sm2Failed := false
		for _, n := range probe.notifications {
			if strings.Contains(n, "state machine graph has to be one click") {
				if strings.Contains(n, "SM1") {
					sm1Failed = true
				}
				if strings.Contains(n, "SM2") {
					sm2Failed = true
				}
			}
		}

		if sm1Failed {
			t.Fatalf("SM1 should not have failed one click rule")
		}
		if !sm2Failed {
			t.Fatalf("SM2 should have failed one click rule")
		}
	})
}

func TestEnforceAtLeastOneDiagramPerStateMachine(t *testing.T) {
	stage := NewStage("test")
	probe := &mockProbe{}
	stager := &Stager{
		stage:     stage,
		probeForm: probe,
	}

	sm := (&StateMachine{Name: "TrafficLight"}).Stage(stage)
	if len(sm.Diagrams) != 0 {
		t.Fatalf("expected 0 diagrams initially")
	}

	needCommit := stager.enforceAtLeastOneDiagramPerStateMachine()
	if !needCommit {
		t.Errorf("expected needCommit to be true")
	}
	if len(sm.Diagrams) != 1 {
		t.Fatalf("expected 1 diagram after enforcement, got %d", len(sm.Diagrams))
	}
	if sm.Diagrams[0].Name != "New Diagram" {
		t.Errorf("expected diagram name 'New Diagram', got %q", sm.Diagrams[0].Name)
	}
	if !sm.Diagrams[0].IsChecked {
		t.Errorf("expected first diagram to be checked")
	}

	foundNotification := false
	for _, n := range probe.notifications {
		if strings.Contains(n, "each state machine has to have at least one diagram") {
			foundNotification = true
			break
		}
	}
	if !foundNotification {
		t.Errorf("expected notification about missing diagram, got: %v", probe.notifications)
	}

	// second pass should not add another diagram
	needCommit2 := stager.enforceAtLeastOneDiagramPerStateMachine()
	if needCommit2 {
		t.Errorf("expected needCommit to be false on second pass")
	}
	if len(sm.Diagrams) != 1 {
		t.Fatalf("expected still 1 diagram, got %d", len(sm.Diagrams))
	}
}

