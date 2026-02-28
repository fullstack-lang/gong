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

	__Product__00000000_ := (&models.Product{Name: `P1`}).Stage(stage)

	__ProductShape__00000000_ := (&models.ProductShape{Name: `-D5`}).Stage(stage)

	__Project__00000001_ := (&models.Project{Name: `P1`}).Stage(stage)

	__Resource__00000003_ := (&models.Resource{Name: `R3`}).Stage(stage)
	__Resource__00000004_ := (&models.Resource{Name: `R4`}).Stage(stage)
	__Resource__00000005_ := (&models.Resource{Name: `R5`}).Stage(stage)
	__Resource__00000006_ := (&models.Resource{Name: `R4.1`}).Stage(stage)
	__Resource__00000007_ := (&models.Resource{Name: `R4.2`}).Stage(stage)
	__Resource__00000008_ := (&models.Resource{Name: `R4.1.1`}).Stage(stage)

	__ResourceCompositionShape__00000000_ := (&models.ResourceCompositionShape{Name: `R4 to `}).Stage(stage)
	__ResourceCompositionShape__00000001_ := (&models.ResourceCompositionShape{Name: `R4 to `}).Stage(stage)
	__ResourceCompositionShape__00000002_ := (&models.ResourceCompositionShape{Name: ` to `}).Stage(stage)

	__ResourceShape__00000003_ := (&models.ResourceShape{Name: `R3-D5`}).Stage(stage)
	__ResourceShape__00000005_ := (&models.ResourceShape{Name: `R5-D5`}).Stage(stage)
	__ResourceShape__00000006_ := (&models.ResourceShape{Name: `R4-D5`}).Stage(stage)
	__ResourceShape__00000007_ := (&models.ResourceShape{Name: `-D5`}).Stage(stage)
	__ResourceShape__00000008_ := (&models.ResourceShape{Name: `-D5`}).Stage(stage)
	__ResourceShape__00000009_ := (&models.ResourceShape{Name: `-D5`}).Stage(stage)

	__ResourceTaskShape__00000000_ := (&models.ResourceTaskShape{Name: ` to T1`}).Stage(stage)
	__ResourceTaskShape__00000001_ := (&models.ResourceTaskShape{Name: `R4 to T1`}).Stage(stage)
	__ResourceTaskShape__00000002_ := (&models.ResourceTaskShape{Name: `R4.1 to T1`}).Stage(stage)

	__Root__00000004_ := (&models.Root{Name: `Root 4`}).Stage(stage)

	__Task__00000000_ := (&models.Task{Name: `T1`}).Stage(stage)

	__TaskOutputShape__00000000_ := (&models.TaskOutputShape{Name: `T1 to P1`}).Stage(stage)

	__TaskShape__00000000_ := (&models.TaskShape{Name: `T1-D5`}).Stage(stage)

	// insertion point for initialization of values

	__Diagram__00000005_.Name = `D5`
	__Diagram__00000005_.IsChecked = true
	__Diagram__00000005_.IsEditable_ = true
	__Diagram__00000005_.ShowPrefix = true
	__Diagram__00000005_.DefaultBoxWidth = 250.000000
	__Diagram__00000005_.DefaultBoxHeigth = 100.000000
	__Diagram__00000005_.Width = 7708.862056
	__Diagram__00000005_.Height = 8076.845003
	__Diagram__00000005_.IsExpanded = true
	__Diagram__00000005_.ComputedPrefix = ``
	__Diagram__00000005_.IsInRenameMode = false
	__Diagram__00000005_.IsPBSNodeExpanded = true
	__Diagram__00000005_.IsWBSNodeExpanded = true
	__Diagram__00000005_.IsNotesNodeExpanded = false
	__Diagram__00000005_.IsResourcesNodeExpanded = true

	__Product__00000000_.Name = `P1`
	__Product__00000000_.Description = ``
	__Product__00000000_.IsExpanded = false
	__Product__00000000_.ComputedPrefix = `1`
	__Product__00000000_.IsInRenameMode = false
	__Product__00000000_.IsProducersNodeExpanded = false
	__Product__00000000_.IsConsumersNodeExpanded = false

	__ProductShape__00000000_.Name = `-D5`
	__ProductShape__00000000_.IsExpanded = false
	__ProductShape__00000000_.X = 1075.448177
	__ProductShape__00000000_.Y = 111.024928
	__ProductShape__00000000_.Width = 250.000000
	__ProductShape__00000000_.Height = 100.000000

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

	__Resource__00000006_.Name = `R4.1`
	__Resource__00000006_.Description = ``
	__Resource__00000006_.IsExpanded = false
	__Resource__00000006_.ComputedPrefix = `2.1`
	__Resource__00000006_.IsInRenameMode = false

	__Resource__00000007_.Name = `R4.2`
	__Resource__00000007_.Description = ``
	__Resource__00000007_.IsExpanded = false
	__Resource__00000007_.ComputedPrefix = `2.2`
	__Resource__00000007_.IsInRenameMode = false

	__Resource__00000008_.Name = `R4.1.1`
	__Resource__00000008_.Description = ``
	__Resource__00000008_.IsExpanded = false
	__Resource__00000008_.ComputedPrefix = `2.1.1`
	__Resource__00000008_.IsInRenameMode = false

	__ResourceCompositionShape__00000000_.Name = `R4 to `
	__ResourceCompositionShape__00000000_.StartRatio = 0.500000
	__ResourceCompositionShape__00000000_.EndRatio = 0.760630
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
	__ResourceCompositionShape__00000002_.CornerOffsetRatio = 1.842251

	__ResourceShape__00000003_.Name = `R3-D5`
	__ResourceShape__00000003_.IsExpanded = false
	__ResourceShape__00000003_.X = 125.659628
	__ResourceShape__00000003_.Y = 229.647694
	__ResourceShape__00000003_.Width = 250.000000
	__ResourceShape__00000003_.Height = 100.000000

	__ResourceShape__00000005_.Name = `R5-D5`
	__ResourceShape__00000005_.IsExpanded = false
	__ResourceShape__00000005_.X = 94.228419
	__ResourceShape__00000005_.Y = 638.987657
	__ResourceShape__00000005_.Width = 250.000000
	__ResourceShape__00000005_.Height = 100.000000

	__ResourceShape__00000006_.Name = `R4-D5`
	__ResourceShape__00000006_.IsExpanded = false
	__ResourceShape__00000006_.X = 521.940130
	__ResourceShape__00000006_.Y = 91.780748
	__ResourceShape__00000006_.Width = 250.000000
	__ResourceShape__00000006_.Height = 100.000000

	__ResourceShape__00000007_.Name = `-D5`
	__ResourceShape__00000007_.IsExpanded = false
	__ResourceShape__00000007_.X = 1170.940130
	__ResourceShape__00000007_.Y = 532.780748
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
	__ResourceShape__00000009_.X = 599.940130
	__ResourceShape__00000009_.Y = 754.780748
	__ResourceShape__00000009_.Width = 250.000000
	__ResourceShape__00000009_.Height = 100.000000

	__ResourceTaskShape__00000000_.Name = ` to T1`
	__ResourceTaskShape__00000000_.StartRatio = 0.500000
	__ResourceTaskShape__00000000_.EndRatio = 0.500000
	__ResourceTaskShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceTaskShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceTaskShape__00000000_.CornerOffsetRatio = 1.680000

	__ResourceTaskShape__00000001_.Name = `R4 to T1`
	__ResourceTaskShape__00000001_.StartRatio = 0.352251
	__ResourceTaskShape__00000001_.EndRatio = 0.500000
	__ResourceTaskShape__00000001_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__ResourceTaskShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceTaskShape__00000001_.CornerOffsetRatio = 1.036630

	__ResourceTaskShape__00000002_.Name = `R4.1 to T1`
	__ResourceTaskShape__00000002_.StartRatio = 0.500000
	__ResourceTaskShape__00000002_.EndRatio = 0.390133
	__ResourceTaskShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceTaskShape__00000002_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ResourceTaskShape__00000002_.CornerOffsetRatio = 0.559165

	__Root__00000004_.Name = `Root 4`
	__Root__00000004_.NbPixPerCharacter = 8.000000

	__Task__00000000_.Name = `T1`
	__Task__00000000_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000000_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000000_.Description = ``
	__Task__00000000_.IsExpanded = false
	__Task__00000000_.ComputedPrefix = `1`
	__Task__00000000_.IsInRenameMode = false
	__Task__00000000_.IsInputsNodeExpanded = false
	__Task__00000000_.IsOutputsNodeExpanded = false
	__Task__00000000_.IsWithCompletion = false
	__Task__00000000_.Completion = ""

	__TaskOutputShape__00000000_.Name = `T1 to P1`
	__TaskOutputShape__00000000_.StartRatio = 0.721465
	__TaskOutputShape__00000000_.EndRatio = 0.500000
	__TaskOutputShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskOutputShape__00000000_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000000_.CornerOffsetRatio = -0.039867

	__TaskShape__00000000_.Name = `T1-D5`
	__TaskShape__00000000_.IsExpanded = false
	__TaskShape__00000000_.X = 854.731320
	__TaskShape__00000000_.Y = 350.992601
	__TaskShape__00000000_.Width = 250.000000
	__TaskShape__00000000_.Height = 100.000000

	// insertion point for setup of pointers
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000000_)
	__Diagram__00000005_.Task_Shapes = append(__Diagram__00000005_.Task_Shapes, __TaskShape__00000000_)
	__Diagram__00000005_.TaskOutputShapes = append(__Diagram__00000005_.TaskOutputShapes, __TaskOutputShape__00000000_)
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
	__Diagram__00000005_.ResourceTaskShapes = append(__Diagram__00000005_.ResourceTaskShapes, __ResourceTaskShape__00000000_)
	__Diagram__00000005_.ResourceTaskShapes = append(__Diagram__00000005_.ResourceTaskShapes, __ResourceTaskShape__00000001_)
	__Diagram__00000005_.ResourceTaskShapes = append(__Diagram__00000005_.ResourceTaskShapes, __ResourceTaskShape__00000002_)
	__ProductShape__00000000_.Product = __Product__00000000_
	__Project__00000001_.RootProducts = append(__Project__00000001_.RootProducts, __Product__00000000_)
	__Project__00000001_.RootTasks = append(__Project__00000001_.RootTasks, __Task__00000000_)
	__Project__00000001_.RootResources = append(__Project__00000001_.RootResources, __Resource__00000003_)
	__Project__00000001_.RootResources = append(__Project__00000001_.RootResources, __Resource__00000004_)
	__Project__00000001_.RootResources = append(__Project__00000001_.RootResources, __Resource__00000005_)
	__Project__00000001_.Diagrams = append(__Project__00000001_.Diagrams, __Diagram__00000005_)
	__Resource__00000004_.Tasks = append(__Resource__00000004_.Tasks, __Task__00000000_)
	__Resource__00000004_.SubResources = append(__Resource__00000004_.SubResources, __Resource__00000006_)
	__Resource__00000004_.SubResources = append(__Resource__00000004_.SubResources, __Resource__00000007_)
	__Resource__00000006_.Tasks = append(__Resource__00000006_.Tasks, __Task__00000000_)
	__Resource__00000006_.SubResources = append(__Resource__00000006_.SubResources, __Resource__00000008_)
	__Resource__00000007_.Tasks = append(__Resource__00000007_.Tasks, __Task__00000000_)
	__ResourceCompositionShape__00000000_.Resource = __Resource__00000006_
	__ResourceCompositionShape__00000001_.Resource = __Resource__00000007_
	__ResourceCompositionShape__00000002_.Resource = __Resource__00000008_
	__ResourceShape__00000003_.Resource = __Resource__00000003_
	__ResourceShape__00000005_.Resource = __Resource__00000005_
	__ResourceShape__00000006_.Resource = __Resource__00000004_
	__ResourceShape__00000007_.Resource = __Resource__00000006_
	__ResourceShape__00000008_.Resource = __Resource__00000007_
	__ResourceShape__00000009_.Resource = __Resource__00000008_
	__ResourceTaskShape__00000000_.Resource = __Resource__00000007_
	__ResourceTaskShape__00000000_.Task = __Task__00000000_
	__ResourceTaskShape__00000001_.Resource = __Resource__00000004_
	__ResourceTaskShape__00000001_.Task = __Task__00000000_
	__ResourceTaskShape__00000002_.Resource = __Resource__00000006_
	__ResourceTaskShape__00000002_.Task = __Task__00000000_
	__Root__00000004_.Projects = append(__Root__00000004_.Projects, __Project__00000001_)
	__Task__00000000_.Outputs = append(__Task__00000000_.Outputs, __Product__00000000_)
	__TaskOutputShape__00000000_.Task = __Task__00000000_
	__TaskOutputShape__00000000_.Product = __Product__00000000_
	__TaskShape__00000000_.Task = __Task__00000000_
}
