package models

import (
	"fmt"
	"os"
	"testing"
	"time"
)

type issue1507ProcessProbe struct {
	ProbeIF
	notifications []string
}

func (m *issue1507ProcessProbe) AddNotification(date time.Time, message string) {
	m.notifications = append(m.notifications, message)
}

func (m *issue1507ProcessProbe) CommitNotificationTable() {}

func (m *issue1507ProcessProbe) SetCommitMode(b bool) {}

func (m *issue1507ProcessProbe) SetUpdateSliceOfPointersCallback(cb func(instance any, fieldName string, slicePtr any)) {
}

func TestGenerateIssue1507ProcessModel(t *testing.T) {
	stage := NewStage("issue1507_process")
	probe := &issue1507ProcessProbe{}
	stager := &Stager{
		stage:     stage,
		probeForm: probe,
	}

	rootLibrary := (&Library{
		Name:                    "Issue 1507: The 3 Stories Explaining the Bug",
		IsRootLibrary:           true,
		NbPixPerCharacter:       8.0,
		IsProcessesNodeExpanded: true,
		IsDataFlowsNodeExpanded: true,
		IsDatasNodeExpanded:     true,
		IsNotesNodeExpanded:     true,
	}).Stage(stage)

	proc := (&Process{
		Name:                    "Issue 1507 - 3 Stories",
		Description:             "Three separate diagrams telling the story of Issue 1507 step by step.",
		IsDataFlowsNodeExpanded: true,
	}).Stage(stage)
	rootLibrary.RootProcesses = append(rootLibrary.RootProcesses, proc)
	rootLibrary.ProcesssWhoseNodeIsExpanded = append(rootLibrary.ProcesssWhoseNodeIsExpanded, proc)

	// Swimlanes (Participants)
	pTree := (&Participant{
		Name:                "1. Tree Callbacks",
		IsTasksNodeExpanded: true,
	}).Stage(stage)
	pStage := (&Participant{
		Name:                "2. Main Stage & Semantic Check",
		IsTasksNodeExpanded: true,
	}).Stage(stage)
	pProbe := (&Participant{
		Name:                "3. Probe Form Stage",
		IsTasksNodeExpanded: true,
	}).Stage(stage)
	pSvg := (&Participant{
		Name:                "4. SVG Canvas Renderer",
		IsTasksNodeExpanded: true,
	}).Stage(stage)
	proc.Participants = append(proc.Participants, pTree, pStage, pProbe, pSvg)

	// =========================================================================
	// STORY 3 DIAGRAM (Checked by default):
	// How clicking the diagram checkbox causes 2 successive commits to formStage
	// =========================================================================
	t1Checkbox := (&Task{
		Name: "1. User clicks the checkbox next to 'Time diagram' in the Tree (diagramNode.OnIsCheckedChanged)",
	}).Stage(stage)
	pTree.Tasks = append(pTree.Tasks, t1Checkbox)

	t2FirstFormCommit := (&Task{
		Name: "2. COMMIT #1 of Form Stage: OnIsCheckedChanged calls FillUpFormFromGongstruct(diagram), which calls formStage.Reset() + formStage.Commit()",
	}).Stage(stage)
	pProbe.Tasks = append(pProbe.Tasks, t2FirstFormCommit)

	t3MainStageCommit := (&Task{
		Name: "3. Next line in OnIsCheckedChanged calls stager.stage.Commit(), which triggers afterCommitHook -> probeForm.Refresh()",
	}).Stage(stage)
	pStage.Tasks = append(pStage.Tasks, t3MainStageCommit)

	t4SecondFormCommit := (&Task{
		Name: "4. COMMIT #2 of Form Stage: probe.ux_form() calls FillUpFormFromGongstruct(diagram) AGAIN -> formStage.Reset() + formStage.Commit()",
	}).Stage(stage)
	pProbe.Tasks = append(pProbe.Tasks, t4SecondFormCommit)

	cfProbe := (&ControlFlow{
		Name:  fmt.Sprintf(`"%s" to "%s"`, t2FirstFormCommit.Name, t4SecondFormCommit.Name),
		Start: t2FirstFormCommit,
		End:   t4SecondFormCommit,
	}).Stage(stage)
	pProbe.ControlFlows = append(pProbe.ControlFlows, cfProbe)

	dDirectCall := (&Data{Name: "1st call: FillUpFormFromGongstruct()"}).Stage(stage)
	dNextLine := (&Data{Name: "Next line: stager.stage.Commit()"}).Stage(stage)
	dAfterHook := (&Data{Name: "afterCommitHook: probeForm.Refresh()"}).Stage(stage)
	rootLibrary.RootDatas = append(rootLibrary.RootDatas, dDirectCall, dNextLine, dAfterHook)

	df1 := (&DataFlow{
		Name:      fmt.Sprintf(`"%s" to "%s"`, t1Checkbox.Name, t2FirstFormCommit.Name),
		Type:      DataFlow_Task2Task,
		StartTask: t1Checkbox,
		EndTask:   t2FirstFormCommit,
		Datas:     []*Data{dDirectCall},
	}).Stage(stage)

	df2 := (&DataFlow{
		Name:      fmt.Sprintf(`"%s" to "%s"`, t1Checkbox.Name, t3MainStageCommit.Name),
		Type:      DataFlow_Task2Task,
		StartTask: t1Checkbox,
		EndTask:   t3MainStageCommit,
		Datas:     []*Data{dNextLine},
	}).Stage(stage)

	df3 := (&DataFlow{
		Name:      fmt.Sprintf(`"%s" to "%s"`, t3MainStageCommit.Name, t4SecondFormCommit.Name),
		Type:      DataFlow_Task2Task,
		StartTask: t3MainStageCommit,
		EndTask:   t4SecondFormCommit,
		Datas:     []*Data{dAfterHook},
	}).Stage(stage)

	rootLibrary.RootDataFlows = append(rootLibrary.RootDataFlows, df1, df2, df3)
	proc.DataFlows = append(proc.DataFlows, df1, df2, df3)

	diag := (&DiagramProcess{
		Name:                       "Story 3: Why Clicking the Checkbox Causes 2 Successive Form Stage Commits",
		IsChecked:                  true,
		IsEditable_:                true,
		DefaultBoxWidth:            300,
		DefaultBoxHeigth:           90,
		Width:                      1650,
		Height:                     950,
		IsParticipantsNodeExpanded: true,
		IsNotesNodeExpanded:        true,
	}).Stage(stage)
	proc.DiagramProcesss = append(proc.DiagramProcesss, diag)
	proc.DiagramProcessWhoseNodeIsExpanded = append(proc.DiagramProcessWhoseNodeIsExpanded, diag)

	pShape := (&ProcessShape{
		Name:    "ProcessShape",
		Process: proc,
	}).Stage(stage)
	pShape.SetX(40)
	pShape.SetY(40)
	pShape.SetWidth(1520)
	pShape.SetHeight(840)
	diag.Process_Shapes = append(diag.Process_Shapes, pShape)

	for _, p := range []*Participant{pTree, pStage, pProbe} {
		ps := (&ParticipantShape{
			Name:        fmt.Sprintf("%s-%s", p.Name, diag.Name),
			Participant: p,
			WidthWeight: 1.0,
		}).Stage(stage)
		diag.Participant_Shapes = append(diag.Participant_Shapes, ps)
	}

	makeTShape := func(task *Task, x, y, w, h float64) *TaskShape {
		ts := (&TaskShape{
			Name: fmt.Sprintf("%s-%s", task.Name, diag.Name),
			Task: task,
		}).Stage(stage)
		ts.SetX(x)
		ts.SetY(y)
		ts.SetWidth(w)
		ts.SetHeight(h)
		diag.Task_Shapes = append(diag.Task_Shapes, ts)
		return ts
	}

	// 3 wide lanes (~500px each), generously spaced vertically
	makeTShape(t1Checkbox, 130, 150, 320, 95)
	makeTShape(t2FirstFormCommit, 1130, 290, 340, 105)
	makeTShape(t3MainStageCommit, 630, 460, 320, 95)
	makeTShape(t4SecondFormCommit, 1130, 620, 340, 105)

	cfs := (&ControlFlowShape{
		Name:        cfProbe.Name,
		ControlFlow: cfProbe,
		LinkShape: LinkShape{
			StartOrientation:  ORIENTATION_VERTICAL,
			EndOrientation:    ORIENTATION_VERTICAL,
			StartRatio:        0.5,
			EndRatio:          0.5,
			CornerOffsetRatio: 1.5,
		},
	}).Stage(stage)
	diag.ControlFlow_Shapes = append(diag.ControlFlow_Shapes, cfs)

	addDFShape := func(df *DataFlow, startRatio, endRatio, cornerOffset float64) {
		dfs := (&DataFlowShape{
			Name:     df.Name,
			DataFlow: df,
			LinkShape: LinkShape{
				StartOrientation:  ORIENTATION_HORIZONTAL,
				EndOrientation:    ORIENTATION_HORIZONTAL,
				StartRatio:        startRatio,
				EndRatio:          endRatio,
				CornerOffsetRatio: cornerOffset,
			},
		}).Stage(stage)
		diag.DataFlow_Shapes = append(diag.DataFlow_Shapes, dfs)

		for _, d := range df.Datas {
			ds := (&DataShape{
				Name:     fmt.Sprintf("%s-%s-%s", df.Name, d.Name, diag.Name),
				Data:     d,
				DataFlow: df,
			}).Stage(stage)
			diag.Data_Shapes = append(diag.Data_Shapes, ds)
		}
	}

	addDFShape(df1, 0.35, 0.5, 0.25)
	addDFShape(df2, 0.75, 0.5, 0.25)
	addDFShape(df3, 0.5, 0.5, 0.25)

	note := (&Note{
		Name: "Why this breaks editing YTopMargin or LaneHeight:\n" +
			"Because formStage is reset and committed TWICE inside the same OnIsCheckedChanged callback:\n" +
			"- Commit #1 creates FormFieldFloat64 instances (for YTopMargin, LaneHeight) and sends them to Angular over the websocket.\n" +
			"- Immediately after, Commit #2 calls formStage.Reset(), destroying Commit #1's FormFieldFloat64 instances and creating new ones.\n" +
			"The Angular frontend can stay bound to Commit #1's FormFieldFloat64 IDs while the backend holds Commit #2's instances!",
		Tasks: []*Task{t4SecondFormCommit},
	}).Stage(stage)
	rootLibrary.RootNotes = append(rootLibrary.RootNotes, note)

	ns := (&NoteShape{
		Name: note.Name + "-" + diag.Name,
		Note: note,
		RectShape: RectShape{
			X:      110,
			Y:      600,
			Width:  840,
			Height: 165,
		},
	}).Stage(stage)
	diag.Note_Shapes = append(diag.Note_Shapes, ns)

	nts := (&NoteTaskShape{
		Name: note.Name + " to " + t4SecondFormCommit.Name,
		Note: note,
		Task: t4SecondFormCommit,
		LinkShape: LinkShape{
			StartOrientation:  ORIENTATION_HORIZONTAL,
			EndOrientation:    ORIENTATION_HORIZONTAL,
			StartRatio:        0.5,
			EndRatio:          0.5,
			CornerOffsetRatio: 1.2,
		},
	}).Stage(stage)
	diag.NoteTaskShapes = append(diag.NoteTaskShapes, nts)

	stage.ComputeReverseMaps()
	stager.enforceSemantic()

	if os.Getenv("GENERATE_ISSUE_1507") == "true" {
		stage.MarshallFile("../cmd/process/data/issue1507.go", "github.com/fullstack-lang/gong/dsm/process/go/models", "main")
	}
}
