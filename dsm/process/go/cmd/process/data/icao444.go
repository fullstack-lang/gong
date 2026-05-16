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

	__Data__00000000_ := (&models.Data{Name: `FPL`}).Stage(stage)
	__Data__00000001_ := (&models.Data{Name: `CHG`}).Stage(stage)
	__Data__00000002_ := (&models.Data{Name: `CNL`}).Stage(stage)
	__Data__00000003_ := (&models.Data{Name: `DLA`}).Stage(stage)
	__Data__00000004_ := (&models.Data{Name: `DEP`}).Stage(stage)
	__Data__00000005_ := (&models.Data{Name: `ARR`}).Stage(stage)

	__DiagramProcess__00000000_ := (&models.DiagramProcess{Name: `DiagramProcess`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: `ICAO 4444 Chapter 11`}).Stage(stage)

	__Process__00000000_ := (&models.Process{Name: `Flight Plan Management`}).Stage(stage)

	__ProcessShape__00000000_ := (&models.ProcessShape{Name: `ProcessShape`}).Stage(stage)

	__Resource__00000000_ := (&models.Resource{Name: `AO`}).Stage(stage)
	__Resource__00000001_ := (&models.Resource{Name: `ARO`}).Stage(stage)
	__Resource__00000002_ := (&models.Resource{Name: `ATC`}).Stage(stage)

	// insertion point for initialization of values

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
	__DiagramProcess__00000000_.Width = 900.000000
	__DiagramProcess__00000000_.Height = 1350.000000
	__DiagramProcess__00000000_.IsProcesssNodeExpanded = false
	__DiagramProcess__00000000_.IsParticipantsNodeExpanded = false
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

	__Process__00000000_.Name = `Flight Plan Management`
	__Process__00000000_.Description = ``
	__Process__00000000_.ComputedPrefix = ``
	__Process__00000000_.SVG_Path = ``
	__Process__00000000_.InverseAppliedScaling = 0.000000
	__Process__00000000_.IsSubProcessNodeExpanded = false
	__Process__00000000_.IsDataFlowsNodeExpanded = false

	__ProcessShape__00000000_.Name = `ProcessShape`
	__ProcessShape__00000000_.IsExpanded = false
	__ProcessShape__00000000_.X = 100.000000
	__ProcessShape__00000000_.Y = 50.000000
	__ProcessShape__00000000_.Width = 500.000000
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
	__DiagramProcess__00000000_.Process_Shapes = append(__DiagramProcess__00000000_.Process_Shapes, __ProcessShape__00000000_)
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
	__Process__00000000_.DiagramProcesss = append(__Process__00000000_.DiagramProcesss, __DiagramProcess__00000000_)
	__Process__00000000_.DiagramProcessWhoseNodeIsExpanded = append(__Process__00000000_.DiagramProcessWhoseNodeIsExpanded, __DiagramProcess__00000000_)
	__ProcessShape__00000000_.Process = __Process__00000000_
}
