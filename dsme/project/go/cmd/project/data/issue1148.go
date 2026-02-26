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

	__Diagram__00000000_ := (&models.Diagram{Name: `D1`}).Stage(stage)

	__Project__00000000_ := (&models.Project{Name: `NewProject`}).Stage(stage)

	__Resource__00000002_ := (&models.Resource{Name: `R1.2`}).Stage(stage)
	__Resource__00000003_ := (&models.Resource{Name: `R1`}).Stage(stage)

	__ResourceCompositionShape__00000000_ := (&models.ResourceCompositionShape{Name: `R1 to R1.2`}).Stage(stage)

	__ResourceShape__00000002_ := (&models.ResourceShape{Name: `NewResource-Default Diagram`}).Stage(stage)
	__ResourceShape__00000003_ := (&models.ResourceShape{Name: `-D1`}).Stage(stage)

	__Root__00000000_ := (&models.Root{Name: `Root`}).Stage(stage)

	// insertion point for initialization of values

	__Diagram__00000000_.Name = `D1`
	__Diagram__00000000_.IsChecked = true
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.ShowPrefix = true
	__Diagram__00000000_.DefaultBoxWidth = 250.000000
	__Diagram__00000000_.DefaultBoxHeigth = 100.000000
	__Diagram__00000000_.Width = 1708.862056
	__Diagram__00000000_.Height = 2076.845003
	__Diagram__00000000_.IsExpanded = true
	__Diagram__00000000_.ComputedPrefix = ``
	__Diagram__00000000_.IsInRenameMode = false
	__Diagram__00000000_.IsPBSNodeExpanded = false
	__Diagram__00000000_.IsWBSNodeExpanded = false
	__Diagram__00000000_.IsNotesNodeExpanded = false
	__Diagram__00000000_.IsResourcesNodeExpanded = true

	__Project__00000000_.Name = `NewProject`
	__Project__00000000_.IsExpanded = true
	__Project__00000000_.ComputedPrefix = ``
	__Project__00000000_.IsInRenameMode = false

	__Resource__00000002_.Name = `R1.2`
	__Resource__00000002_.Description = ``
	__Resource__00000002_.IsExpanded = false
	__Resource__00000002_.ComputedPrefix = `1.1`
	__Resource__00000002_.IsInRenameMode = false

	__Resource__00000003_.Name = `R1`
	__Resource__00000003_.Description = ``
	__Resource__00000003_.IsExpanded = false
	__Resource__00000003_.ComputedPrefix = `1`
	__Resource__00000003_.IsInRenameMode = false

	__ResourceCompositionShape__00000000_.Name = `R1 to R1.2`
	__ResourceCompositionShape__00000000_.StartRatio = 0.500000
	__ResourceCompositionShape__00000000_.EndRatio = 0.500000
	__ResourceCompositionShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000000_.CornerOffsetRatio = 1.680000

	__ResourceShape__00000002_.Name = `NewResource-Default Diagram`
	__ResourceShape__00000002_.IsExpanded = false
	__ResourceShape__00000002_.X = 283.862056
	__ResourceShape__00000002_.Y = 413.845003
	__ResourceShape__00000002_.Width = 250.000000
	__ResourceShape__00000002_.Height = 73.000000

	__ResourceShape__00000003_.Name = `-D1`
	__ResourceShape__00000003_.IsExpanded = false
	__ResourceShape__00000003_.X = 152.659628
	__ResourceShape__00000003_.Y = 184.647694
	__ResourceShape__00000003_.Width = 250.000000
	__ResourceShape__00000003_.Height = 100.000000

	__Root__00000000_.Name = `Root`
	__Root__00000000_.NbPixPerCharacter = 8.000000

	// insertion point for setup of pointers
	__Diagram__00000000_.Resource_Shapes = append(__Diagram__00000000_.Resource_Shapes, __ResourceShape__00000002_)
	__Diagram__00000000_.Resource_Shapes = append(__Diagram__00000000_.Resource_Shapes, __ResourceShape__00000003_)
	__Diagram__00000000_.ResourceComposition_Shapes = append(__Diagram__00000000_.ResourceComposition_Shapes, __ResourceCompositionShape__00000000_)
	__Project__00000000_.RootResources = append(__Project__00000000_.RootResources, __Resource__00000003_)
	__Project__00000000_.Diagrams = append(__Project__00000000_.Diagrams, __Diagram__00000000_)
	__Resource__00000003_.SubResources = append(__Resource__00000003_.SubResources, __Resource__00000002_)
	__ResourceCompositionShape__00000000_.Resource = __Resource__00000002_
	__ResourceShape__00000002_.Resource = __Resource__00000002_
	__ResourceShape__00000003_.Resource = __Resource__00000003_
	__Root__00000000_.Projects = append(__Root__00000000_.Projects, __Project__00000000_)
}
