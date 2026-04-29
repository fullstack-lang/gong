package main

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/process/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var (
	_ time.Time
	_ = slices.Index[[]int, int]
)

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__DiagramProcess__00000005_ := (&models.DiagramProcess{Name: `DiagramProcess`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: `Root`}).Stage(stage)

	__Participant__00000003_ := (&models.Participant{Name: `PP1`}).Stage(stage)
	__Participant__00000004_ := (&models.Participant{Name: `PP2`}).Stage(stage)
	__Participant__00000005_ := (&models.Participant{Name: `PP3`}).Stage(stage)

	__ParticipantShape__00000006_ := (&models.ParticipantShape{Name: `PP2-DiagramProcess`}).Stage(stage)
	__ParticipantShape__00000007_ := (&models.ParticipantShape{Name: `PP3-DiagramProcess`}).Stage(stage)
	__ParticipantShape__00000009_ := (&models.ParticipantShape{Name: `PP1-DiagramProcess`}).Stage(stage)

	__Process__00000005_ := (&models.Process{Name: `P1`}).Stage(stage)

	__ProcessShape__00000005_ := (&models.ProcessShape{Name: `ProcessShape`}).Stage(stage)

	__Task__00000001_ := (&models.Task{Name: `T1`}).Stage(stage)
	__Task__00000002_ := (&models.Task{Name: `T2`}).Stage(stage)
	__Task__00000003_ := (&models.Task{Name: `PP2 T1`}).Stage(stage)
	__Task__00000004_ := (&models.Task{Name: `Start`}).Stage(stage)
	__Task__00000005_ := (&models.Task{Name: `End`}).Stage(stage)

	__TaskShape__00000000_ := (&models.TaskShape{Name: `T1-DiagramProcess`}).Stage(stage)
	__TaskShape__00000002_ := (&models.TaskShape{Name: `T2-DiagramProcess`}).Stage(stage)
	__TaskShape__00000004_ := (&models.TaskShape{Name: `Start-DiagramProcess`}).Stage(stage)
	__TaskShape__00000005_ := (&models.TaskShape{Name: `PP2 T1-DiagramProcess`}).Stage(stage)
	__TaskShape__00000006_ := (&models.TaskShape{Name: `End-DiagramProcess`}).Stage(stage)

	// insertion point for initialization of values

	__DiagramProcess__00000005_.Name = `DiagramProcess`
	__DiagramProcess__00000005_.ComputedPrefix = ``
	__DiagramProcess__00000005_.IsChecked = true
	__DiagramProcess__00000005_.IsEditable_ = true
	__DiagramProcess__00000005_.IsShowPrefix = false
	__DiagramProcess__00000005_.DefaultBoxWidth = 250.000000
	__DiagramProcess__00000005_.DefaultBoxHeigth = 70.000000
	__DiagramProcess__00000005_.Width = 1521.000000
	__DiagramProcess__00000005_.Height = 1318.000000
	__DiagramProcess__00000005_.IsProcesssNodeExpanded = false
	__DiagramProcess__00000005_.IsParticipantsNodeExpanded = true

	__Library__00000000_.Name = `Root`
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.NbPixPerCharacter = 8.000000
	__Library__00000000_.LogoSVGFile = ``
	__Library__00000000_.IsExpandedTmp = true

	__Participant__00000003_.Name = `PP1`
	__Participant__00000003_.ComputedPrefix = ``

	__Participant__00000004_.Name = `PP2`
	__Participant__00000004_.ComputedPrefix = ``

	__Participant__00000005_.Name = `PP3`
	__Participant__00000005_.ComputedPrefix = ``

	__ParticipantShape__00000006_.Name = `PP2-DiagramProcess`
	__ParticipantShape__00000006_.IsExpanded = false
	__ParticipantShape__00000006_.X = 461.000000
	__ParticipantShape__00000006_.Y = 68.000000
	__ParticipantShape__00000006_.Width = 375.000000
	__ParticipantShape__00000006_.Height = 940.000000
	__ParticipantShape__00000006_.IsHidden = false

	__ParticipantShape__00000007_.Name = `PP3-DiagramProcess`
	__ParticipantShape__00000007_.IsExpanded = false
	__ParticipantShape__00000007_.X = 673.250000
	__ParticipantShape__00000007_.Y = 68.000000
	__ParticipantShape__00000007_.Width = 195.750000
	__ParticipantShape__00000007_.Height = 940.000000
	__ParticipantShape__00000007_.IsHidden = false

	__ParticipantShape__00000009_.Name = `PP1-DiagramProcess`
	__ParticipantShape__00000009_.IsExpanded = false
	__ParticipantShape__00000009_.X = 86.000000
	__ParticipantShape__00000009_.Y = 68.000000
	__ParticipantShape__00000009_.Width = 375.000000
	__ParticipantShape__00000009_.Height = 940.000000
	__ParticipantShape__00000009_.IsHidden = false

	__Process__00000005_.Name = `P1`
	__Process__00000005_.ComputedPrefix = ``
	__Process__00000005_.IsSubProcessNodeExpanded = false

	__ProcessShape__00000005_.Name = `ProcessShape`
	__ProcessShape__00000005_.IsExpanded = false
	__ProcessShape__00000005_.X = 76.000000
	__ProcessShape__00000005_.Y = 18.000000
	__ProcessShape__00000005_.Width = 1145.000000
	__ProcessShape__00000005_.Height = 1000.000000
	__ProcessShape__00000005_.IsHidden = false

	__Task__00000001_.Name = `T1`
	__Task__00000001_.ComputedPrefix = ``
	__Task__00000001_.IsStartTask = false
	__Task__00000001_.IsEndTask = false

	__Task__00000002_.Name = `T2`
	__Task__00000002_.ComputedPrefix = ``
	__Task__00000002_.IsStartTask = false
	__Task__00000002_.IsEndTask = false

	__Task__00000003_.Name = `PP2 T1`
	__Task__00000003_.ComputedPrefix = ``
	__Task__00000003_.IsStartTask = true
	__Task__00000003_.IsEndTask = false

	__Task__00000004_.Name = `Start`
	__Task__00000004_.ComputedPrefix = ``
	__Task__00000004_.IsStartTask = true
	__Task__00000004_.IsEndTask = false

	__Task__00000005_.Name = `End`
	__Task__00000005_.ComputedPrefix = ``
	__Task__00000005_.IsStartTask = false
	__Task__00000005_.IsEndTask = true

	__TaskShape__00000000_.Name = `T1-DiagramProcess`
	__TaskShape__00000000_.IsExpanded = false
	__TaskShape__00000000_.X = 142.000000
	__TaskShape__00000000_.Y = 258.141353
	__TaskShape__00000000_.Width = 163.000000
	__TaskShape__00000000_.Height = 70.000000
	__TaskShape__00000000_.IsHidden = false

	__TaskShape__00000002_.Name = `T2-DiagramProcess`
	__TaskShape__00000002_.IsExpanded = false
	__TaskShape__00000002_.X = 126.000000
	__TaskShape__00000002_.Y = 402.786461
	__TaskShape__00000002_.Width = 152.000000
	__TaskShape__00000002_.Height = 70.000000
	__TaskShape__00000002_.IsHidden = false

	__TaskShape__00000004_.Name = `Start-DiagramProcess`
	__TaskShape__00000004_.IsExpanded = false
	__TaskShape__00000004_.X = 191.000000
	__TaskShape__00000004_.Y = 143.683844
	__TaskShape__00000004_.Width = 58.000000
	__TaskShape__00000004_.Height = 20.000000
	__TaskShape__00000004_.IsHidden = false

	__TaskShape__00000005_.Name = `PP2 T1-DiagramProcess`
	__TaskShape__00000005_.IsExpanded = false
	__TaskShape__00000005_.X = 607.000000
	__TaskShape__00000005_.Y = 146.449476
	__TaskShape__00000005_.Width = 75.000000
	__TaskShape__00000005_.Height = 20.000000
	__TaskShape__00000005_.IsHidden = false

	__TaskShape__00000006_.Name = `End-DiagramProcess`
	__TaskShape__00000006_.IsExpanded = false
	__TaskShape__00000006_.X = 195.000000
	__TaskShape__00000006_.Y = 578.101302
	__TaskShape__00000006_.Width = 68.000000
	__TaskShape__00000006_.Height = 36.000000
	__TaskShape__00000006_.IsHidden = false

	// insertion point for setup of pointers
	__DiagramProcess__00000005_.Process_Shapes = append(__DiagramProcess__00000005_.Process_Shapes, __ProcessShape__00000005_)
	__DiagramProcess__00000005_.Participant_Shapes = append(__DiagramProcess__00000005_.Participant_Shapes, __ParticipantShape__00000009_)
	__DiagramProcess__00000005_.Participant_Shapes = append(__DiagramProcess__00000005_.Participant_Shapes, __ParticipantShape__00000006_)
	__DiagramProcess__00000005_.Participant_Shapes = append(__DiagramProcess__00000005_.Participant_Shapes, __ParticipantShape__00000007_)
	__DiagramProcess__00000005_.ParticipantWhoseNodeIsExpanded = append(__DiagramProcess__00000005_.ParticipantWhoseNodeIsExpanded, __Participant__00000003_)
	__DiagramProcess__00000005_.ParticipantWhoseNodeIsExpanded = append(__DiagramProcess__00000005_.ParticipantWhoseNodeIsExpanded, __Participant__00000004_)
	__DiagramProcess__00000005_.TaskShapes = append(__DiagramProcess__00000005_.TaskShapes, __TaskShape__00000000_)
	__DiagramProcess__00000005_.TaskShapes = append(__DiagramProcess__00000005_.TaskShapes, __TaskShape__00000002_)
	__DiagramProcess__00000005_.TaskShapes = append(__DiagramProcess__00000005_.TaskShapes, __TaskShape__00000004_)
	__DiagramProcess__00000005_.TaskShapes = append(__DiagramProcess__00000005_.TaskShapes, __TaskShape__00000005_)
	__DiagramProcess__00000005_.TaskShapes = append(__DiagramProcess__00000005_.TaskShapes, __TaskShape__00000006_)
	__Library__00000000_.RootProcesses = append(__Library__00000000_.RootProcesses, __Process__00000005_)
	__Library__00000000_.ProcesssWhoseNodeIsExpanded = append(__Library__00000000_.ProcesssWhoseNodeIsExpanded, __Process__00000005_)
	__Participant__00000003_.Tasks = append(__Participant__00000003_.Tasks, __Task__00000001_)
	__Participant__00000003_.Tasks = append(__Participant__00000003_.Tasks, __Task__00000002_)
	__Participant__00000003_.Tasks = append(__Participant__00000003_.Tasks, __Task__00000004_)
	__Participant__00000003_.Tasks = append(__Participant__00000003_.Tasks, __Task__00000005_)
	__Participant__00000004_.Tasks = append(__Participant__00000004_.Tasks, __Task__00000003_)
	__ParticipantShape__00000006_.Participant = __Participant__00000004_
	__ParticipantShape__00000007_.Participant = __Participant__00000005_
	__ParticipantShape__00000009_.Participant = __Participant__00000003_
	__Process__00000005_.DiagramProcesss = append(__Process__00000005_.DiagramProcesss, __DiagramProcess__00000005_)
	__Process__00000005_.DiagramProcessWhoseNodeIsExpanded = append(__Process__00000005_.DiagramProcessWhoseNodeIsExpanded, __DiagramProcess__00000005_)
	__Process__00000005_.Participants = append(__Process__00000005_.Participants, __Participant__00000003_)
	__Process__00000005_.Participants = append(__Process__00000005_.Participants, __Participant__00000004_)
	__Process__00000005_.Participants = append(__Process__00000005_.Participants, __Participant__00000005_)
	__Process__00000005_.ParticipantWhoseNodeIsExpanded = append(__Process__00000005_.ParticipantWhoseNodeIsExpanded, __Participant__00000003_)
	__ProcessShape__00000005_.Process = __Process__00000005_
	__TaskShape__00000000_.Task = __Task__00000001_
	__TaskShape__00000002_.Task = __Task__00000002_
	__TaskShape__00000004_.Task = __Task__00000004_
	__TaskShape__00000005_.Task = __Task__00000003_
	__TaskShape__00000006_.Task = __Task__00000005_
}
