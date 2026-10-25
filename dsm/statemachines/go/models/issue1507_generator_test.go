package models

import (
	"os"
	"testing"
	"time"
)

type issue1507Probe struct {
	ProbeIF
	notifications []string
}

func (m *issue1507Probe) AddNotification(date time.Time, message string) {
	m.notifications = append(m.notifications, message)
}

func (m *issue1507Probe) CommitNotificationTable() {}

func TestGenerateIssue1507StateMachine(t *testing.T) {
	stage := NewStage("issue1507")
	probe := &issue1507Probe{}
	stager := &Stager{
		stage:     stage,
		probeForm: probe,
	}

	rootLib := (&Library{
		Name:                        "Issue 1507: The 3 Stories Explaining the Bug",
		IsRootLibrary:               true,
		NbPixPerCharacter:           8.0,
		IsStateMachinesNodeExpanded: true,
		IsSubLibrariesNodeExpanded:  false,
		IsExpandedTmp:               true,
	}).Stage(stage)

	// Helper to build a clean top-to-bottom story state machine where:
	// - States are clean boxes on the left (X = 120..480) with no cryptic /entry or /do code dumps
	// - Every state has a yellow Note on the right (X = 560..1100) explaining the story in plain English
	type stepSpec struct {
		stateName      string
		transitionFrom string // label on incoming transition from previous state
		noteText       string
		noteHeight     float64
	}

	buildStoryStateMachine := func(smName, diagName string, isChecked bool, steps []stepSpec) (*StateMachine, *State) {
		sm := (&StateMachine{
			Name:       smName,
			IsExpanded: true,
		}).Stage(stage)

		diag := (&Diagram{
			Name:                 diagName,
			IsChecked:            isChecked,
			IsExpanded:           true,
			IsEditable_:          true,
			IsStatesNodeExpanded: true,
			ShowRoles:            false,
			ShowMessages:         false,
		}).Stage(stage)
		sm.Diagrams = append(sm.Diagrams, diag)

		startState := (&State{
			Name: "Start",
		}).Stage(stage)
		sm.InitialState = startState
		sm.States = append(sm.States, startState)

		startShape := (&StateShape{
			Name:  startState.Name,
			State: startState,
			RectShape: RectShape{
				X:      210,
				Y:      20,
				Width:  100,
				Height: 20,
			},
		}).Stage(stage)
		diag.State_Shapes = append(diag.State_Shapes, startShape)

		prevState := startState
		var firstRealState *State

		currentY := 85.0
		for i, spec := range steps {
			st := (&State{
				Name: spec.stateName,
			}).Stage(stage)
			sm.States = append(sm.States, st)
			if i == 0 {
				firstRealState = st
			}

			stateH := 80.0
			ss := (&StateShape{
				Name:  st.Name,
				State: st,
				RectShape: RectShape{
					X:      120,
					Y:      currentY,
					Width:  360,
					Height: stateH,
				},
			}).Stage(stage)
			diag.State_Shapes = append(diag.State_Shapes, ss)

			tr := (&Transition{
				Name:  spec.transitionFrom,
				Start: prevState,
				End:   st,
			}).Stage(stage)
			ts := (&Transition_Shape{
				Name:       prevState.Name + " to " + st.Name,
				Transition: tr,
				LinkShape: LinkShape{
					StartOrientation:  ORIENTATION_VERTICAL,
					EndOrientation:    ORIENTATION_VERTICAL,
					StartRatio:        0.5,
					EndRatio:          0.5,
					CornerOffsetRatio: 1.2,
				},
			}).Stage(stage)
			diag.Transition_Shapes = append(diag.Transition_Shapes, ts)

			if spec.noteText != "" {
				note := (&Note{
					Name:  spec.noteText,
					State: st,
				}).Stage(stage)
				st.Notes = append(st.Notes, note)

				noteH := spec.noteHeight
				if noteH == 0 {
					noteH = 130
				}
				noteY := currentY - (noteH-stateH)/2.0
				if noteY < 15 {
					noteY = 15
				}

				ns := (&NoteShape{
					Name: note.Name + "-" + diag.Name,
					Note: note,
					RectShape: RectShape{
						X:      560,
						Y:      noteY,
						Width:  540,
						Height: noteH,
					},
				}).Stage(stage)
				diag.Note_Shapes = append(diag.Note_Shapes, ns)

				nss := (&NoteStateShape{
					Name:  note.Name + "-" + st.Name,
					Note:  note,
					State: st,
					LinkShape: LinkShape{
						StartOrientation:  ORIENTATION_HORIZONTAL,
						EndOrientation:    ORIENTATION_HORIZONTAL,
						StartRatio:        0.5,
						EndRatio:          0.5,
						CornerOffsetRatio: 1.2,
					},
				}).Stage(stage)
				diag.NoteState_Shapes = append(diag.NoteState_Shapes, nss)

				currentY = noteY + noteH + 45.0
			} else {
				currentY += stateH + 80.0
			}

			prevState = st
		}

		rootLib.RootStateMachines = append(rootLib.RootStateMachines, sm)
		rootLib.StateMachinesWhoseNodeIsExpanded = append(rootLib.StateMachinesWhoseNodeIsExpanded, sm)
		return sm, firstRealState
	}

	// =========================================================================
	// DIAGRAM 1: Story 1 - Why Drag-and-Drop Created Two Checked Diagrams
	// =========================================================================
	_, firstState1 := buildStoryStateMachine(
		"Story 1: Drag-and-Drop Creates 2 Checked Diagrams",
		"1. Story 1: Why Drag-and-Drop Created 2 Checked Diagrams",
		true,
		[]stepSpec{
			{
				stateName:      "1. App Starts with Default Sample Data",
				transitionFrom: "App launch",
				noteText: "In the 'project' DSM, each Diagram has a boolean field: IsChecked.\n" +
					"In the left tree, each diagram has a checkbox bound to IsChecked, and stager.svg() draws whichever diagram has IsChecked == true.\n" +
					"At startup, only 1 diagram is checked.",
				noteHeight: 130,
			},
			{
				stateName:      "2. User Drags & Drops Confidential Project File",
				transitionFrom: "Drag & drop .go file",
				noteText: "loadProxy.OnFileUpload resets the stage and parses the uploaded .go file (ParseAstFromBytes).\n" +
					"In the uploaded file:\n" +
					"- The main Library has the user's 'Time diagram' with IsChecked = true.\n" +
					"- A SubLibrary in the project has 0 diagrams of its own (len(subLibrary.Diagrams) == 0).",
				noteHeight: 145,
			},
			{
				stateName:      "3. Semantic Check Creates a Second Checked Diagram",
				transitionFrom: "ParseAstFromBytes calls stage.Commit()",
				noteText: "stage.Commit() runs enforceLibraryHasAtLeastOneDiagram().\n" +
					"Because the SubLibrary has 0 diagrams, Gong automatically creates a new 'Default Diagram' for that SubLibrary with IsChecked = true, WITHOUT unchecking existing diagrams!\n" +
					"Now TWO diagrams in the stage have IsChecked = true at the same time:\n" +
					"1) The user's 'Time diagram' (IsChecked = true)\n" +
					"2) The SubLibrary's new 'Default Diagram' (IsChecked = true)",
				noteHeight: 175,
			},
			{
				stateName:      "4. SVG Renderer Picks From an Unordered Go Map",
				transitionFrom: "stager.svg() redraws canvas",
				noteText: "In stager_ux_svg.go, stager.svg() selects the diagram to draw using a map loop without 'break':\n" +
					"  for diagram_ := range *stager.stage.GetInstancesSet[*Diagram]() {\n" +
					"      if diagram_.IsChecked { diagram = diagram_ }\n" +
					"  }\n" +
					"Because TWO diagrams have IsChecked = true, whichever diagram happens to be visited last in Go's random map iteration wins and gets drawn!",
				noteHeight: 175,
			},
		},
	)

	// =========================================================================
	// DIAGRAM 2: Story 2 - Why Clicking the Diagram Name in the Tree Still Failed
	// =========================================================================
	buildStoryStateMachine(
		"Story 2: Clicking Diagram Name in Tree",
		"2. Story 2: Why Clicking the Diagram Name Still Failed",
		false,
		[]stepSpec{
			{
				stateName:      "1. Stage Has Two Diagrams with IsChecked = true",
				transitionFrom: "After drag & drop (Story 1)",
				noteText:       "Both the user's 'Time diagram' and the SubLibrary's 'Default Diagram' currently have IsChecked = true in the stage.",
				noteHeight:     110,
			},
			{
				stateName:      "2. User Clicks 'Time diagram' Name in the Tree",
				transitionFrom: "User clicks diagram text label",
				noteText: "When the user clicks the text label 'Time diagram' in the left tree (not the checkbox), diagramNode.OnClick runs onNodeClicked(stager, diagram).\n" +
					"This calls FillUpFormFromGongstruct(diagram, 'Diagram').\n" +
					"So the right-hand Probe form IS reloaded with 'Time diagram's field values (YTopMargin, LaneHeight, etc.).",
				noteHeight: 150,
			},
			{
				stateName:      "3. OnClick Did Not Uncheck the Other Diagram",
				transitionFrom: "What OnClick forgot to do",
				noteText: "Unlike the checkbox callback, diagramNode.OnClick ONLY reloaded the Probe form.\n" +
					"It did NOT uncheck the SubLibrary's 'Default Diagram'!\n" +
					"So even after selecting 'Time diagram' in the tree, BOTH diagrams still have IsChecked = true in the stage.",
				noteHeight: 135,
			},
			{
				stateName:      "4. User Changes YTopMargin or LaneHeight & Clicks Save",
				transitionFrom: "User edits parameter & clicks Save",
				noteText: "Saving the Probe form updates YTopMargin or LaneHeight on 'Time diagram' and calls stage.Commit().\n" +
					"However, when stager.svg() runs to redraw the SVG canvas, both diagrams still have IsChecked = true.\n" +
					"When the map loop visits 'Default Diagram' after 'Time diagram', stager.svg() draws 'Default Diagram' instead of the updated 'Time diagram'!",
				noteHeight: 165,
			},
		},
	)

	// =========================================================================
	// DIAGRAM 3: Story 3 - Why Clicking the Checkbox Causes 2 Successive Form Stage Commits
	// =========================================================================
	buildStoryStateMachine(
		"Story 3: 2 Successive Commits to Form Stage",
		"3. Story 3: 2 Successive Commits to Form Stage (Checkbox Click)",
		false,
		[]stepSpec{
			{
				stateName:      "1. User Clicks the Diagram Checkbox in the Tree",
				transitionFrom: "User toggles checkbox",
				noteText: "In stager_ux_tree_diagram.go, clicking the checkbox next to a diagram node runs the callback:\n" +
					"  diagramNode.OnIsCheckedChanged\n" +
					"It unchecks all other diagrams and sets diagram.IsChecked = true.",
				noteHeight: 125,
			},
			{
				stateName:      "2. First Commit of Form Stage (Direct Call)",
				transitionFrom: "1st call inside OnIsCheckedChanged",
				noteText: "Still inside OnIsCheckedChanged, the code directly calls:\n" +
					"  stager.probeForm.FillUpFormFromGongstruct(diagram, 'Diagram')\n" +
					"This resets formStage, creates new FormGroup and FormFieldFloat64 instances (for YTopMargin, LaneHeight, etc.), and calls:\n" +
					"  formStage.Commit()   <-- COMMIT #1 of Form Stage!\n" +
					"The form websocket notifies the Angular frontend of Commit #1.",
				noteHeight: 165,
			},
			{
				stateName:      "3. Main Stage Commit Runs After-Commit Hook",
				transitionFrom: "Very next line in OnIsCheckedChanged",
				noteText: "On the very next line of OnIsCheckedChanged, the code calls:\n" +
					"  stager.stage.Commit()\n" +
					"Committing the main stage triggers its afterCommitHooks, which calls:\n" +
					"  stager.ux() -> stager.probeForm.Refresh() -> probe.ux_form()",
				noteHeight: 135,
			},
			{
				stateName:      "4. Second Commit of Form Stage in the Same Callback",
				transitionFrom: "probe.ux_form() rebuilds the form again",
				noteText: "Inside probe.ux_form(), it calls FillUpFormFromGongstruct(diagram) a SECOND time!\n" +
					"This calls formStage.Reset() (destroying all FormFieldFloat64 instances created in State 2), creates a brand-new set of FormFieldFloat64 instances, and calls:\n" +
					"  formStage.Commit()   <-- COMMIT #2 of Form Stage in the same callback!",
				noteHeight: 160,
			},
			{
				stateName:      "5. Angular Form and Backend Form Stage Are Out of Sync",
				transitionFrom: "User edits YTopMargin / LaneHeight",
				noteText: "This violates the Gong rule 'Avoid Successive Stack Updates':\n" +
					"Because formStage was committed twice in one callback, the Angular frontend can display the FormFieldFloat64 instances from Commit #1 while the backend formStage has already replaced them with Commit #2.\n" +
					"When the user edits YTopMargin or LaneHeight and clicks Save, the edited field does not match the backend's live FormFieldFloat64 instance!",
				noteHeight: 175,
			},
		},
	)

	// Highlight State 1 in Story 1
	_ = (&Object{
		Name:       "Reader",
		State:      firstState1,
		IsSelected: true,
		DOF:        time.Date(2026, 10, 25, 9, 0, 0, 0, time.UTC),
	}).Stage(stage)

	stage.ComputeReverseMaps()
	stager.enforceSemantic()

	if len(probe.notifications) > 0 {
		t.Fatalf("unexpected semantic notifications: %v", probe.notifications)
	}

	if os.Getenv("GENERATE_ISSUE_1507") == "true" {
		stage.MarshallFile("../cmd/statemachines/data/issue1507.go", "github.com/fullstack-lang/gong/dsm/statemachines/go/models", "main")
	}
}
