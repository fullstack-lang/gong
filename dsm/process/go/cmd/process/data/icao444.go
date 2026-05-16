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

	__AllocatedResourceShape__00000000_ := (&models.AllocatedResourceShape{Name: `DiagramProcess-Aircraft Operator-AO`}).Stage(stage)
	__AllocatedResourceShape__00000001_ := (&models.AllocatedResourceShape{Name: `DiagramProcess-ATS Reporting Office-ARO`}).Stage(stage)
	__AllocatedResourceShape__00000002_ := (&models.AllocatedResourceShape{Name: `DiagramProcess-ATS Unit-ATC`}).Stage(stage)

	__Data__00000000_ := (&models.Data{Name: `FPL`}).Stage(stage)
	__Data__00000001_ := (&models.Data{Name: `CHG`}).Stage(stage)
	__Data__00000002_ := (&models.Data{Name: `CNL`}).Stage(stage)
	__Data__00000003_ := (&models.Data{Name: `DLA`}).Stage(stage)
	__Data__00000004_ := (&models.Data{Name: `DEP`}).Stage(stage)
	__Data__00000005_ := (&models.Data{Name: `ARR`}).Stage(stage)

	__DiagramProcess__00000000_ := (&models.DiagramProcess{Name: `DiagramProcess`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: `ICAO 4444 Chapter 11`}).Stage(stage)

	__Participant__00000000_ := (&models.Participant{Name: `Aircraft Operator`}).Stage(stage)
	__Participant__00000001_ := (&models.Participant{Name: `ATS Reporting Office`}).Stage(stage)
	__Participant__00000002_ := (&models.Participant{Name: `ATS Unit`}).Stage(stage)

	__ParticipantShape__00000000_ := (&models.ParticipantShape{Name: `Aircraft Operator-DiagramProcess`}).Stage(stage)
	__ParticipantShape__00000001_ := (&models.ParticipantShape{Name: `ATS Reporting Office-DiagramProcess`}).Stage(stage)
	__ParticipantShape__00000002_ := (&models.ParticipantShape{Name: `ATS Unit-DiagramProcess`}).Stage(stage)

	__Process__00000000_ := (&models.Process{Name: `ARO Message Exchanges`}).Stage(stage)

	__ProcessShape__00000000_ := (&models.ProcessShape{Name: `ProcessShape`}).Stage(stage)

	__Resource__00000000_ := (&models.Resource{Name: `AO`}).Stage(stage)
	__Resource__00000001_ := (&models.Resource{Name: `ARO`}).Stage(stage)
	__Resource__00000002_ := (&models.Resource{Name: `ATC`}).Stage(stage)

	// insertion point for initialization of values

	__AllocatedResourceShape__00000000_.Name = `DiagramProcess-Aircraft Operator-AO`

	__AllocatedResourceShape__00000001_.Name = `DiagramProcess-ATS Reporting Office-ARO`

	__AllocatedResourceShape__00000002_.Name = `DiagramProcess-ATS Unit-ATC`

	__Data__00000000_.Name = `FPL`
	__Data__00000000_.Acronym = ``
	__Data__00000000_.Description = `Filed Flight Plan Message`
	__Data__00000000_.ComputedPrefix = ``
	__Data__00000000_.SVG_Path = ``
	__Data__00000000_.InverseAppliedScaling = 0.000000

	__Data__00000001_.Name = `CHG`
	__Data__00000001_.Acronym = ``
	__Data__00000001_.Description = `Modification Message`
	__Data__00000001_.ComputedPrefix = ``
	__Data__00000001_.SVG_Path = ``
	__Data__00000001_.InverseAppliedScaling = 0.000000

	__Data__00000002_.Name = `CNL`
	__Data__00000002_.Acronym = ``
	__Data__00000002_.Description = `Flight Plan Cancellation Message`
	__Data__00000002_.ComputedPrefix = ``
	__Data__00000002_.SVG_Path = ``
	__Data__00000002_.InverseAppliedScaling = 0.000000

	__Data__00000003_.Name = `DLA`
	__Data__00000003_.Acronym = ``
	__Data__00000003_.Description = `Delay Message`
	__Data__00000003_.ComputedPrefix = ``
	__Data__00000003_.SVG_Path = ``
	__Data__00000003_.InverseAppliedScaling = 0.000000

	__Data__00000004_.Name = `DEP`
	__Data__00000004_.Acronym = ``
	__Data__00000004_.Description = `Departure Message`
	__Data__00000004_.ComputedPrefix = ``
	__Data__00000004_.SVG_Path = ``
	__Data__00000004_.InverseAppliedScaling = 0.000000

	__Data__00000005_.Name = `ARR`
	__Data__00000005_.Acronym = ``
	__Data__00000005_.Description = `Arrival Message`
	__Data__00000005_.ComputedPrefix = ``
	__Data__00000005_.SVG_Path = ``
	__Data__00000005_.InverseAppliedScaling = 0.000000

	__DiagramProcess__00000000_.Name = `DiagramProcess`
	__DiagramProcess__00000000_.Description = ``
	__DiagramProcess__00000000_.ComputedPrefix = ``
	__DiagramProcess__00000000_.IsChecked = true
	__DiagramProcess__00000000_.IsEditable_ = true
	__DiagramProcess__00000000_.IsShowPrefix = false
	__DiagramProcess__00000000_.DefaultBoxWidth = 250.000000
	__DiagramProcess__00000000_.DefaultBoxHeigth = 70.000000
	__DiagramProcess__00000000_.Width = 1167.000000
	__DiagramProcess__00000000_.Height = 1327.000000
	__DiagramProcess__00000000_.IsProcesssNodeExpanded = false
	__DiagramProcess__00000000_.IsParticipantsNodeExpanded = true
	__DiagramProcess__00000000_.IsExternalParticipantsNodeExpanded = false
	__DiagramProcess__00000000_.IsNotesNodeExpanded = false

	__Library__00000000_.Name = `ICAO 4444 Chapter 11`
	__Library__00000000_.Description = `Modeling of message exchanges handled by the ATS Reporting Office (ARO)`
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsSubLibrariesNodeExpanded = false
	__Library__00000000_.NbPixPerCharacter = 8.000000
	__Library__00000000_.LogoSVGFile = ``
	__Library__00000000_.IsProcessesNodeExpanded = true
	__Library__00000000_.IsDataFlowsNodeExpanded = false
	__Library__00000000_.IsDatasNodeExpanded = true
	__Library__00000000_.IsResourcesNodeExpanded = true
	__Library__00000000_.IsNotesNodeExpanded = false
	__Library__00000000_.IsExpandedTmp = true

	__Participant__00000000_.Name = `Aircraft Operator`
	__Participant__00000000_.IsProcessResource = false
	__Participant__00000000_.Description = `Originator of initial flight plan data and operational modifications`
	__Participant__00000000_.IsResourcesNodeExpanded = true
	__Participant__00000000_.IsProcessesNodeExpanded = false
	__Participant__00000000_.ComputedPrefix = ``
	__Participant__00000000_.IsTasksNodeExpanded = false
	__Participant__00000000_.IsControlFlowsNodeExpanded = false
	__Participant__00000000_.IsDataFlowsNodeExpanded = false

	__Participant__00000001_.Name = `ATS Reporting Office`
	__Participant__00000001_.IsProcessResource = false
	__Participant__00000001_.Description = `Receives, validates, accepts, and distributes flight plan-related messages`
	__Participant__00000001_.IsResourcesNodeExpanded = true
	__Participant__00000001_.IsProcessesNodeExpanded = false
	__Participant__00000001_.ComputedPrefix = ``
	__Participant__00000001_.IsTasksNodeExpanded = false
	__Participant__00000001_.IsControlFlowsNodeExpanded = false
	__Participant__00000001_.IsDataFlowsNodeExpanded = false

	__Participant__00000002_.Name = `ATS Unit`
	__Participant__00000002_.IsProcessResource = false
	__Participant__00000002_.Description = `Recipient of distributed movement and control data used for traffic separation`
	__Participant__00000002_.IsResourcesNodeExpanded = true
	__Participant__00000002_.IsProcessesNodeExpanded = false
	__Participant__00000002_.ComputedPrefix = ``
	__Participant__00000002_.IsTasksNodeExpanded = false
	__Participant__00000002_.IsControlFlowsNodeExpanded = false
	__Participant__00000002_.IsDataFlowsNodeExpanded = false

	__ParticipantShape__00000000_.Name = `Aircraft Operator-DiagramProcess`
	__ParticipantShape__00000000_.IsExpanded = false
	__ParticipantShape__00000000_.X = 153.916317
	__ParticipantShape__00000000_.Y = 178.327215
	__ParticipantShape__00000000_.Width = 250.000000
	__ParticipantShape__00000000_.Height = 70.000000
	__ParticipantShape__00000000_.IsHidden = false
	__ParticipantShape__00000000_.WidthWeight = 1.000000

	__ParticipantShape__00000001_.Name = `ATS Reporting Office-DiagramProcess`
	__ParticipantShape__00000001_.IsExpanded = false
	__ParticipantShape__00000001_.X = 108.634825
	__ParticipantShape__00000001_.Y = 100.269741
	__ParticipantShape__00000001_.Width = 250.000000
	__ParticipantShape__00000001_.Height = 70.000000
	__ParticipantShape__00000001_.IsHidden = false
	__ParticipantShape__00000001_.WidthWeight = 1.000000

	__ParticipantShape__00000002_.Name = `ATS Unit-DiagramProcess`
	__ParticipantShape__00000002_.IsExpanded = false
	__ParticipantShape__00000002_.X = 143.845785
	__ParticipantShape__00000002_.Y = 168.047525
	__ParticipantShape__00000002_.Width = 250.000000
	__ParticipantShape__00000002_.Height = 70.000000
	__ParticipantShape__00000002_.IsHidden = false
	__ParticipantShape__00000002_.WidthWeight = 1.000000

	__Process__00000000_.Name = `ARO Message Exchanges`
	__Process__00000000_.Description = `Process governing the submission, validation, modification, and distribution of ATS messages via the ARO.`
	__Process__00000000_.ComputedPrefix = ``
	__Process__00000000_.SVG_Path = ``
	__Process__00000000_.InverseAppliedScaling = 0.000000
	__Process__00000000_.IsSubProcessNodeExpanded = false
	__Process__00000000_.IsDataFlowsNodeExpanded = true

	__ProcessShape__00000000_.Name = `ProcessShape`
	__ProcessShape__00000000_.IsExpanded = false
	__ProcessShape__00000000_.X = 138.000000
	__ProcessShape__00000000_.Y = 27.000000
	__ProcessShape__00000000_.Width = 729.000000
	__ProcessShape__00000000_.Height = 1000.000000
	__ProcessShape__00000000_.IsHidden = false

	__Resource__00000000_.Name = `AO`
	__Resource__00000000_.Acronym = ``
	__Resource__00000000_.Description = `Aircraft Operator / Flight Crew`
	__Resource__00000000_.ComputedPrefix = ``
	__Resource__00000000_.SVG_Path = ``
	__Resource__00000000_.InverseAppliedScaling = 0.000000

	__Resource__00000001_.Name = `ARO`
	__Resource__00000001_.Acronym = ``
	__Resource__00000001_.Description = `ATS Reporting Office`
	__Resource__00000001_.ComputedPrefix = ``
	__Resource__00000001_.SVG_Path = ``
	__Resource__00000001_.InverseAppliedScaling = 0.000000

	__Resource__00000002_.Name = `ATC`
	__Resource__00000002_.Acronym = ``
	__Resource__00000002_.Description = `Air Traffic Control Unit / ACC / TWR`
	__Resource__00000002_.ComputedPrefix = ``
	__Resource__00000002_.SVG_Path = ``
	__Resource__00000002_.InverseAppliedScaling = 0.000000

	// insertion point for setup of pointers
	__AllocatedResourceShape__00000000_.Participant = __Participant__00000000_
	__AllocatedResourceShape__00000000_.Resource = __Resource__00000000_
	__AllocatedResourceShape__00000001_.Participant = __Participant__00000001_
	__AllocatedResourceShape__00000001_.Resource = __Resource__00000001_
	__AllocatedResourceShape__00000002_.Participant = __Participant__00000002_
	__AllocatedResourceShape__00000002_.Resource = __Resource__00000002_
	__DiagramProcess__00000000_.Process_Shapes = append(__DiagramProcess__00000000_.Process_Shapes, __ProcessShape__00000000_)
	__DiagramProcess__00000000_.Participant_Shapes = append(__DiagramProcess__00000000_.Participant_Shapes, __ParticipantShape__00000000_)
	__DiagramProcess__00000000_.Participant_Shapes = append(__DiagramProcess__00000000_.Participant_Shapes, __ParticipantShape__00000001_)
	__DiagramProcess__00000000_.Participant_Shapes = append(__DiagramProcess__00000000_.Participant_Shapes, __ParticipantShape__00000002_)
	__DiagramProcess__00000000_.ParticipantWhoseNodeIsExpanded = append(__DiagramProcess__00000000_.ParticipantWhoseNodeIsExpanded, __Participant__00000000_)
	__DiagramProcess__00000000_.ParticipantWhoseNodeIsExpanded = append(__DiagramProcess__00000000_.ParticipantWhoseNodeIsExpanded, __Participant__00000001_)
	__DiagramProcess__00000000_.ParticipantWhoseNodeIsExpanded = append(__DiagramProcess__00000000_.ParticipantWhoseNodeIsExpanded, __Participant__00000002_)
	__DiagramProcess__00000000_.AllocatedResourceShapes = append(__DiagramProcess__00000000_.AllocatedResourceShapes, __AllocatedResourceShape__00000000_)
	__DiagramProcess__00000000_.AllocatedResourceShapes = append(__DiagramProcess__00000000_.AllocatedResourceShapes, __AllocatedResourceShape__00000001_)
	__DiagramProcess__00000000_.AllocatedResourceShapes = append(__DiagramProcess__00000000_.AllocatedResourceShapes, __AllocatedResourceShape__00000002_)
	__Library__00000000_.RootProcesses = append(__Library__00000000_.RootProcesses, __Process__00000000_)
	__Library__00000000_.ProcesssWhoseNodeIsExpanded = append(__Library__00000000_.ProcesssWhoseNodeIsExpanded, __Process__00000000_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000000_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000001_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000002_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000003_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000004_)
	__Library__00000000_.RootDatas = append(__Library__00000000_.RootDatas, __Data__00000005_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000000_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000001_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000002_)
	__Participant__00000000_.Resources = append(__Participant__00000000_.Resources, __Resource__00000000_)
	__Participant__00000001_.Resources = append(__Participant__00000001_.Resources, __Resource__00000001_)
	__Participant__00000002_.Resources = append(__Participant__00000002_.Resources, __Resource__00000002_)
	__ParticipantShape__00000000_.Participant = __Participant__00000000_
	__ParticipantShape__00000001_.Participant = __Participant__00000001_
	__ParticipantShape__00000002_.Participant = __Participant__00000002_
	__Process__00000000_.DiagramProcesss = append(__Process__00000000_.DiagramProcesss, __DiagramProcess__00000000_)
	__Process__00000000_.DiagramProcessWhoseNodeIsExpanded = append(__Process__00000000_.DiagramProcessWhoseNodeIsExpanded, __DiagramProcess__00000000_)
	__Process__00000000_.Participants = append(__Process__00000000_.Participants, __Participant__00000000_)
	__Process__00000000_.Participants = append(__Process__00000000_.Participants, __Participant__00000001_)
	__Process__00000000_.Participants = append(__Process__00000000_.Participants, __Participant__00000002_)
	__ProcessShape__00000000_.Process = __Process__00000000_
}
