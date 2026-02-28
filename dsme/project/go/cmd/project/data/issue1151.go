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
	__Resource__00000004_ := (&models.Resource{Name: `R4`}).Stage(stage)
	__Resource__00000005_ := (&models.Resource{Name: `R5`}).Stage(stage)
	__Resource__00000006_ := (&models.Resource{Name: ``}).Stage(stage)
	__Resource__00000007_ := (&models.Resource{Name: ``}).Stage(stage)
	__Resource__00000008_ := (&models.Resource{Name: ``}).Stage(stage)

	__ResourceCompositionShape__00000000_ := (&models.ResourceCompositionShape{Name: `R4 to `}).Stage(stage)
	__ResourceCompositionShape__00000001_ := (&models.ResourceCompositionShape{Name: `R4 to `}).Stage(stage)
	__ResourceCompositionShape__00000002_ := (&models.ResourceCompositionShape{Name: ` to `}).Stage(stage)

	__ResourceShape__00000003_ := (&models.ResourceShape{Name: `R3-D5`}).Stage(stage)
	__ResourceShape__00000005_ := (&models.ResourceShape{Name: `R5-D5`}).Stage(stage)
	__ResourceShape__00000006_ := (&models.ResourceShape{Name: `R4-D5`}).Stage(stage)
	__ResourceShape__00000007_ := (&models.ResourceShape{Name: `-D5`}).Stage(stage)
	__ResourceShape__00000008_ := (&models.ResourceShape{Name: `-D5`}).Stage(stage)
	__ResourceShape__00000009_ := (&models.ResourceShape{Name: `-D5`}).Stage(stage)

	__Root__00000004_ := (&models.Root{Name: `Root 4`}).Stage(stage)

	__Task__00000000_ := (&models.Task{Name: `T1`}).Stage(stage)

	__TaskShape__00000000_ := (&models.TaskShape{Name: `T1-D5`}).Stage(stage)

	// insertion point for initialization of values

	__Diagram__00000005_.Name = `D5`
	__Diagram__00000005_.IsChecked = true
	__Diagram__00000005_.IsEditable_ = true
	__Diagram__00000005_.ShowPrefix = true
	__Diagram__00000005_.DefaultBoxWidth = 250.000000
	__Diagram__00000005_.DefaultBoxHeigth = 100.000000
	__Diagram__00000005_.Width = 4208.862056
	__Diagram__00000005_.Height = 4576.845003
	__Diagram__00000005_.IsExpanded = true
	__Diagram__00000005_.ComputedPrefix = ``
	__Diagram__00000005_.IsInRenameMode = false
	__Diagram__00000005_.IsPBSNodeExpanded = false
	__Diagram__00000005_.IsWBSNodeExpanded = true
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

	__Resource__00000004_.Name = `R4`
	__Resource__00000004_.Description = ``
	__Resource__00000004_.IsExpanded = false
	__Resource__00000004_.ComputedPrefix = `2`
	__Resource__00000004_.IsInRenameMode = false

	__Resource__00000005_.Name = `R5`
	__Resource__00000005_.Description = ``
	__Resource__00000005_.IsExpanded = false
	__Resource__00000005_.ComputedPrefix = `3`
	__Resource__00000005_.IsInRenameMode = false

	__Resource__00000006_.Name = ``
	__Resource__00000006_.Description = ``
	__Resource__00000006_.IsExpanded = false
	__Resource__00000006_.ComputedPrefix = `2.1`
	__Resource__00000006_.IsInRenameMode = false

	__Resource__00000007_.Name = ``
	__Resource__00000007_.Description = ``
	__Resource__00000007_.IsExpanded = false
	__Resource__00000007_.ComputedPrefix = `2.2`
	__Resource__00000007_.IsInRenameMode = false

	__Resource__00000008_.Name = ``
	__Resource__00000008_.Description = ``
	__Resource__00000008_.IsExpanded = false
	__Resource__00000008_.ComputedPrefix = `2.1.1`
	__Resource__00000008_.IsInRenameMode = false

	__ResourceCompositionShape__00000000_.Name = `R4 to `
	__ResourceCompositionShape__00000000_.StartRatio = 0.500000
	__ResourceCompositionShape__00000000_.EndRatio = 0.500000
	__ResourceCompositionShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000000_.CornerOffsetRatio = 1.680000

	__ResourceCompositionShape__00000001_.Name = `R4 to `
	__ResourceCompositionShape__00000001_.StartRatio = 0.500000
	__ResourceCompositionShape__00000001_.EndRatio = 0.500000
	__ResourceCompositionShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000001_.CornerOffsetRatio = 1.680000

	__ResourceCompositionShape__00000002_.Name = ` to `
	__ResourceCompositionShape__00000002_.StartRatio = 0.500000
	__ResourceCompositionShape__00000002_.EndRatio = 0.500000
	__ResourceCompositionShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000002_.CornerOffsetRatio = 1.680000

	__ResourceShape__00000003_.Name = `R3-D5`
	__ResourceShape__00000003_.IsExpanded = false
	__ResourceShape__00000003_.X = 125.659628
	__ResourceShape__00000003_.Y = 314.647694
	__ResourceShape__00000003_.Width = 250.000000
	__ResourceShape__00000003_.Height = 100.000000

	__ResourceShape__00000005_.Name = `R5-D5`
	__ResourceShape__00000005_.IsExpanded = false
	__ResourceShape__00000005_.X = 130.228419
	__ResourceShape__00000005_.Y = 467.987657
	__ResourceShape__00000005_.Width = 250.000000
	__ResourceShape__00000005_.Height = 100.000000

	__ResourceShape__00000006_.Name = `R4-D5`
	__ResourceShape__00000006_.IsExpanded = false
	__ResourceShape__00000006_.X = 162.940130
	__ResourceShape__00000006_.Y = 190.780748
	__ResourceShape__00000006_.Width = 250.000000
	__ResourceShape__00000006_.Height = 100.000000

	__ResourceShape__00000007_.Name = `-D5`
	__ResourceShape__00000007_.IsExpanded = false
	__ResourceShape__00000007_.X = 162.940130
	__ResourceShape__00000007_.Y = 390.780748
	__ResourceShape__00000007_.Width = 250.000000
	__ResourceShape__00000007_.Height = 100.000000

	__ResourceShape__00000008_.Name = `-D5`
	__ResourceShape__00000008_.IsExpanded = false
	__ResourceShape__00000008_.X = 462.940130
	__ResourceShape__00000008_.Y = 390.780748
	__ResourceShape__00000008_.Width = 250.000000
	__ResourceShape__00000008_.Height = 100.000000

	__ResourceShape__00000009_.Name = `-D5`
	__ResourceShape__00000009_.IsExpanded = false
	__ResourceShape__00000009_.X = 162.940130
	__ResourceShape__00000009_.Y = 590.780748
	__ResourceShape__00000009_.Width = 250.000000
	__ResourceShape__00000009_.Height = 100.000000

	__Root__00000004_.Name = `Root 4`
	__Root__00000004_.NbPixPerCharacter = 8.000000

	__Task__00000000_.Name = `T1`
	__Task__00000000_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000000_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000000_.Description = ``
	__Task__00000000_.IsExpanded = false
	__Task__00000000_.ComputedPrefix = `1`
	__Task__00000000_.IsInRenameMode = true
	__Task__00000000_.IsInputsNodeExpanded = false
	__Task__00000000_.IsOutputsNodeExpanded = false
	__Task__00000000_.IsWithCompletion = false
	__Task__00000000_.Completion = ""

	__TaskShape__00000000_.Name = `T1-D5`
	__TaskShape__00000000_.IsExpanded = false
	__TaskShape__00000000_.X = 114.731320
	__TaskShape__00000000_.Y = 114.992601
	__TaskShape__00000000_.Width = 250.000000
	__TaskShape__00000000_.Height = 100.000000

	// insertion point for setup of pointers
	__Diagram__00000005_.Task_Shapes = append(__Diagram__00000005_.Task_Shapes, __TaskShape__00000000_)
	__Diagram__00000005_.Resource_Shapes = append(__Diagram__00000005_.Resource_Shapes, __ResourceShape__00000003_)
	__Diagram__00000005_.Resource_Shapes = append(__Diagram__00000005_.Resource_Shapes, __ResourceShape__00000005_)
	__Diagram__00000005_.Resource_Shapes = append(__Diagram__00000005_.Resource_Shapes, __ResourceShape__00000006_)
	__Diagram__00000005_.Resource_Shapes = append(__Diagram__00000005_.Resource_Shapes, __ResourceShape__00000007_)
	__Diagram__00000005_.Resource_Shapes = append(__Diagram__00000005_.Resource_Shapes, __ResourceShape__00000008_)
	__Diagram__00000005_.Resource_Shapes = append(__Diagram__00000005_.Resource_Shapes, __ResourceShape__00000009_)
	__Diagram__00000005_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000005_.ResourcesWhoseNodeIsExpanded, __Resource__00000004_)
	__Diagram__00000005_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000005_.ResourcesWhoseNodeIsExpanded, __Resource__00000006_)
	__Diagram__00000005_.ResourceComposition_Shapes = append(__Diagram__00000005_.ResourceComposition_Shapes, __ResourceCompositionShape__00000000_)
	__Diagram__00000005_.ResourceComposition_Shapes = append(__Diagram__00000005_.ResourceComposition_Shapes, __ResourceCompositionShape__00000001_)
	__Diagram__00000005_.ResourceComposition_Shapes = append(__Diagram__00000005_.ResourceComposition_Shapes, __ResourceCompositionShape__00000002_)
	__Project__00000001_.RootTasks = append(__Project__00000001_.RootTasks, __Task__00000000_)
	__Project__00000001_.RootResources = append(__Project__00000001_.RootResources, __Resource__00000003_)
	__Project__00000001_.RootResources = append(__Project__00000001_.RootResources, __Resource__00000004_)
	__Project__00000001_.RootResources = append(__Project__00000001_.RootResources, __Resource__00000005_)
	__Project__00000001_.Diagrams = append(__Project__00000001_.Diagrams, __Diagram__00000005_)
	__Resource__00000004_.SubResources = append(__Resource__00000004_.SubResources, __Resource__00000006_)
	__Resource__00000004_.SubResources = append(__Resource__00000004_.SubResources, __Resource__00000007_)
	__Resource__00000006_.SubResources = append(__Resource__00000006_.SubResources, __Resource__00000008_)
	__ResourceCompositionShape__00000000_.Resource = __Resource__00000006_
	__ResourceCompositionShape__00000001_.Resource = __Resource__00000007_
	__ResourceCompositionShape__00000002_.Resource = __Resource__00000008_
	__ResourceShape__00000003_.Resource = __Resource__00000003_
	__ResourceShape__00000005_.Resource = __Resource__00000005_
	__ResourceShape__00000006_.Resource = __Resource__00000004_
	__ResourceShape__00000007_.Resource = __Resource__00000006_
	__ResourceShape__00000008_.Resource = __Resource__00000007_
	__ResourceShape__00000009_.Resource = __Resource__00000008_
	__Root__00000004_.Projects = append(__Root__00000004_.Projects, __Project__00000001_)
	__TaskShape__00000000_.Task = __Task__00000000_
}
