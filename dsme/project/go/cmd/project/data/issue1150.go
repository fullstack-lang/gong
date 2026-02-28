package main

import (
	"time"

	"github.com/fullstack-lang/gong/dsme/project/go/models"
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

	// insertion point for declaration of instances to stage

	__Diagram__00000005_ := (&models.Diagram{Name: `D5`}).Stage(stage)

	__Project__00000001_ := (&models.Project{Name: `P1`}).Stage(stage)

	__Resource__00000003_ := (&models.Resource{Name: `R3`}).Stage(stage)

	__ResourceShape__00000003_ := (&models.ResourceShape{Name: `R3-D5`}).Stage(stage)

	__Root__00000004_ := (&models.Root{Name: `Root 4`}).Stage(stage)

	// insertion point for initialization of values

	__Diagram__00000005_.Name = `D5`
	__Diagram__00000005_.IsChecked = true
	__Diagram__00000005_.IsEditable_ = true
	__Diagram__00000005_.ShowPrefix = true
	__Diagram__00000005_.DefaultBoxWidth = 250.000000
	__Diagram__00000005_.DefaultBoxHeigth = 100.000000
	__Diagram__00000005_.Width = 2108.862056
	__Diagram__00000005_.Height = 2476.845003
	__Diagram__00000005_.IsExpanded = true
	__Diagram__00000005_.ComputedPrefix = ``
	__Diagram__00000005_.IsInRenameMode = false
	__Diagram__00000005_.IsPBSNodeExpanded = false
	__Diagram__00000005_.IsWBSNodeExpanded = false
	__Diagram__00000005_.IsNotesNodeExpanded = false
	__Diagram__00000005_.IsResourcesNodeExpanded = true

	__Project__00000001_.Name = `P1`
	__Project__00000001_.IsExpanded = true
	__Project__00000001_.ComputedPrefix = ``
	__Project__00000001_.IsInRenameMode = false

	__Resource__00000003_.Name = `R3`
	__Resource__00000003_.Description = ``
	__Resource__00000003_.IsExpanded = false
	__Resource__00000003_.ComputedPrefix = `1`
	__Resource__00000003_.IsInRenameMode = false

	__ResourceShape__00000003_.Name = `R3-D5`
	__ResourceShape__00000003_.IsExpanded = false
	__ResourceShape__00000003_.X = 152.659628
	__ResourceShape__00000003_.Y = 184.647694
	__ResourceShape__00000003_.Width = 250.000000
	__ResourceShape__00000003_.Height = 100.000000

	__Root__00000004_.Name = `Root 4`
	__Root__00000004_.NbPixPerCharacter = 8.000000

	__Root__00000004_.Projects = append(__Root__00000004_.Projects, __Project__00000001_)
	__Project__00000001_.RootResources = append(__Project__00000001_.RootResources, __Resource__00000003_)
	__Project__00000001_.Diagrams = append(__Project__00000001_.Diagrams, __Diagram__00000005_)
	__ResourceShape__00000003_.Resource = __Resource__00000003_
	__Diagram__00000005_.Resource_Shapes = append(__Diagram__00000005_.Resource_Shapes, __ResourceShape__00000003_)

	stage.Commit()

	// __Diagram__00000005_.Resource_Shapes = slices.Delete(__Diagram__00000005_.Resource_Shapes, 0, 1)
	// __ResourceShape__00000003_.Unstage(stage)
	// stage.Commit()
}
