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

	__Diagram__00000000_ := (&models.Diagram{Name: `Default Diagram`}).Stage(stage)
	__Diagram__00000001_ := (&models.Diagram{Name: `WBS`}).Stage(stage)
	__Diagram__00000002_ := (&models.Diagram{Name: `PBS`}).Stage(stage)

	__Note__00000000_ := (&models.Note{Name: `CFT ended in march 2025`}).Stage(stage)

	__NoteShape__00000000_ := (&models.NoteShape{Name: `-Default Diagram`}).Stage(stage)

	__NoteTaskShape__00000000_ := (&models.NoteTaskShape{Name: `CFT ended in march 2025 to Starliner Crewed Flight Test (CFT)`}).Stage(stage)

	__Product__00000000_ := (&models.Product{Name: `International Space Station (ISS)`}).Stage(stage)
	__Product__00000001_ := (&models.Product{Name: `Dragon`}).Stage(stage)
	__Product__00000002_ := (&models.Product{Name: `Starliner`}).Stage(stage)

	__ProductShape__00000000_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000001_ := (&models.ProductShape{Name: `International Space Station (ISS)-PBS`}).Stage(stage)
	__ProductShape__00000002_ := (&models.ProductShape{Name: `Dragon-PBS`}).Stage(stage)
	__ProductShape__00000003_ := (&models.ProductShape{Name: `Starliner-PBS`}).Stage(stage)

	__Project__00000000_ := (&models.Project{Name: `Startliner Mishape Report`}).Stage(stage)

	__Resource__00000000_ := (&models.Resource{Name: `PITProgram Investigation Team (PIT)`}).Stage(stage)
	__Resource__00000001_ := (&models.Resource{Name: `Barry "Butch" Wilmore`}).Stage(stage)
	__Resource__00000002_ := (&models.Resource{Name: `Sunita "Suni" Williams`}).Stage(stage)

	__ResourceShape__00000000_ := (&models.ResourceShape{Name: `-Default Diagram`}).Stage(stage)
	__ResourceShape__00000001_ := (&models.ResourceShape{Name: `-Default Diagram`}).Stage(stage)
	__ResourceShape__00000002_ := (&models.ResourceShape{Name: `-Default Diagram`}).Stage(stage)

	__ResourceTaskShape__00000000_ := (&models.ResourceTaskShape{Name: `PITProgram Investigation Team (PIT) to Mishap investigation`}).Stage(stage)

	__Root__00000000_ := (&models.Root{Name: `Root`}).Stage(stage)

	__Task__00000000_ := (&models.Task{Name: `Starliner Crewed Flight Test (CFT)`}).Stage(stage)
	__Task__00000001_ := (&models.Task{Name: `Mishap investigation`}).Stage(stage)
	__Task__00000002_ := (&models.Task{Name: `Commercial Crew Program (CCP),`}).Stage(stage)
	__Task__00000003_ := (&models.Task{Name: ` Commercial ReSupply (CRS) `}).Stage(stage)

	__TaskCompositionShape__00000000_ := (&models.TaskCompositionShape{Name: `Commercial Crew Program (CCP), to Starliner Crewed Flight Test (CFT)`}).Stage(stage)
	__TaskCompositionShape__00000001_ := (&models.TaskCompositionShape{Name: `Commercial Crew Program (CCP), to Starliner Crewed Flight Test (CFT)`}).Stage(stage)

	__TaskShape__00000000_ := (&models.TaskShape{Name: `NewTask-Default Diagram`}).Stage(stage)
	__TaskShape__00000001_ := (&models.TaskShape{Name: `NewTask-Default Diagram`}).Stage(stage)
	__TaskShape__00000002_ := (&models.TaskShape{Name: `-Default Diagram`}).Stage(stage)
	__TaskShape__00000003_ := (&models.TaskShape{Name: `Mishap investigation-WBS`}).Stage(stage)
	__TaskShape__00000004_ := (&models.TaskShape{Name: `Commercial Crew Program (CCP),-WBS`}).Stage(stage)
	__TaskShape__00000005_ := (&models.TaskShape{Name: `Starliner Crewed Flight Test (CFT)-WBS`}).Stage(stage)
	__TaskShape__00000006_ := (&models.TaskShape{Name: `-WBS`}).Stage(stage)

	// insertion point for initialization of values

	__Diagram__00000000_.Name = `Default Diagram`
	__Diagram__00000000_.IsChecked = false
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.ShowPrefix = false
	__Diagram__00000000_.DefaultBoxWidth = 250.000000
	__Diagram__00000000_.DefaultBoxHeigth = 70.000000
	__Diagram__00000000_.IsExpanded = false
	__Diagram__00000000_.ComputedPrefix = ``
	__Diagram__00000000_.IsInRenameMode = false
	__Diagram__00000000_.IsPBSNodeExpanded = true
	__Diagram__00000000_.IsWBSNodeExpanded = true
	__Diagram__00000000_.IsNotesNodeExpanded = true
	__Diagram__00000000_.IsResourcesNodeExpanded = true

	__Diagram__00000001_.Name = `WBS`
	__Diagram__00000001_.IsChecked = false
	__Diagram__00000001_.IsEditable_ = true
	__Diagram__00000001_.ShowPrefix = false
	__Diagram__00000001_.DefaultBoxWidth = 250.000000
	__Diagram__00000001_.DefaultBoxHeigth = 70.000000
	__Diagram__00000001_.IsExpanded = false
	__Diagram__00000001_.ComputedPrefix = ``
	__Diagram__00000001_.IsInRenameMode = false
	__Diagram__00000001_.IsPBSNodeExpanded = true
	__Diagram__00000001_.IsWBSNodeExpanded = true
	__Diagram__00000001_.IsNotesNodeExpanded = false
	__Diagram__00000001_.IsResourcesNodeExpanded = false

	__Diagram__00000002_.Name = `PBS`
	__Diagram__00000002_.IsChecked = true
	__Diagram__00000002_.IsEditable_ = true
	__Diagram__00000002_.ShowPrefix = false
	__Diagram__00000002_.DefaultBoxWidth = 250.000000
	__Diagram__00000002_.DefaultBoxHeigth = 70.000000
	__Diagram__00000002_.IsExpanded = true
	__Diagram__00000002_.ComputedPrefix = ``
	__Diagram__00000002_.IsInRenameMode = false
	__Diagram__00000002_.IsPBSNodeExpanded = true
	__Diagram__00000002_.IsWBSNodeExpanded = false
	__Diagram__00000002_.IsNotesNodeExpanded = false
	__Diagram__00000002_.IsResourcesNodeExpanded = false

	__Note__00000000_.Name = `CFT ended in march 2025`
	__Note__00000000_.IsExpanded = false
	__Note__00000000_.ComputedPrefix = `1`
	__Note__00000000_.IsInRenameMode = false

	__NoteShape__00000000_.Name = `-Default Diagram`
	__NoteShape__00000000_.IsExpanded = false
	__NoteShape__00000000_.X = 92.441418
	__NoteShape__00000000_.Y = 1047.751617
	__NoteShape__00000000_.Width = 250.000000
	__NoteShape__00000000_.Height = 70.000000

	__NoteTaskShape__00000000_.Name = `CFT ended in march 2025 to Starliner Crewed Flight Test (CFT)`
	__NoteTaskShape__00000000_.StartRatio = 0.500000
	__NoteTaskShape__00000000_.EndRatio = 0.500000
	__NoteTaskShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteTaskShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteTaskShape__00000000_.CornerOffsetRatio = 1.680000

	__Product__00000000_.Name = `International Space Station (ISS)`
	__Product__00000000_.Description = ``
	__Product__00000000_.IsExpanded = false
	__Product__00000000_.ComputedPrefix = `1`
	__Product__00000000_.IsInRenameMode = false
	__Product__00000000_.IsProducersNodeExpanded = false
	__Product__00000000_.IsConsumersNodeExpanded = false

	__Product__00000001_.Name = `Dragon`
	__Product__00000001_.Description = ``
	__Product__00000001_.IsExpanded = false
	__Product__00000001_.ComputedPrefix = `2`
	__Product__00000001_.IsInRenameMode = false
	__Product__00000001_.IsProducersNodeExpanded = false
	__Product__00000001_.IsConsumersNodeExpanded = false

	__Product__00000002_.Name = `Starliner`
	__Product__00000002_.Description = ``
	__Product__00000002_.IsExpanded = false
	__Product__00000002_.ComputedPrefix = `3`
	__Product__00000002_.IsInRenameMode = false
	__Product__00000002_.IsProducersNodeExpanded = false
	__Product__00000002_.IsConsumersNodeExpanded = false

	__ProductShape__00000000_.Name = `-Default Diagram`
	__ProductShape__00000000_.IsExpanded = false
	__ProductShape__00000000_.X = 38.591007
	__ProductShape__00000000_.Y = 409.214483
	__ProductShape__00000000_.Width = 250.000000
	__ProductShape__00000000_.Height = 70.000000

	__ProductShape__00000001_.Name = `International Space Station (ISS)-PBS`
	__ProductShape__00000001_.IsExpanded = false
	__ProductShape__00000001_.X = 32.344404
	__ProductShape__00000001_.Y = 59.322372
	__ProductShape__00000001_.Width = 250.000000
	__ProductShape__00000001_.Height = 70.000000

	__ProductShape__00000002_.Name = `Dragon-PBS`
	__ProductShape__00000002_.IsExpanded = false
	__ProductShape__00000002_.X = 40.020302
	__ProductShape__00000002_.Y = 276.039798
	__ProductShape__00000002_.Width = 250.000000
	__ProductShape__00000002_.Height = 70.000000

	__ProductShape__00000003_.Name = `Starliner-PBS`
	__ProductShape__00000003_.IsExpanded = false
	__ProductShape__00000003_.X = 29.852191
	__ProductShape__00000003_.Y = 380.054421
	__ProductShape__00000003_.Width = 250.000000
	__ProductShape__00000003_.Height = 70.000000

	__Project__00000000_.Name = `Startliner Mishape Report`
	__Project__00000000_.IsExpanded = true
	__Project__00000000_.ComputedPrefix = ``
	__Project__00000000_.IsInRenameMode = false

	__Resource__00000000_.Name = `PITProgram Investigation Team (PIT)`
	__Resource__00000000_.Description = ``
	__Resource__00000000_.IsExpanded = false
	__Resource__00000000_.ComputedPrefix = `1`
	__Resource__00000000_.IsInRenameMode = false

	__Resource__00000001_.Name = `Barry "Butch" Wilmore`
	__Resource__00000001_.Description = ``
	__Resource__00000001_.IsExpanded = false
	__Resource__00000001_.ComputedPrefix = `2`
	__Resource__00000001_.IsInRenameMode = false

	__Resource__00000002_.Name = `Sunita "Suni" Williams`
	__Resource__00000002_.Description = ``
	__Resource__00000002_.IsExpanded = false
	__Resource__00000002_.ComputedPrefix = `3`
	__Resource__00000002_.IsInRenameMode = false

	__ResourceShape__00000000_.Name = `-Default Diagram`
	__ResourceShape__00000000_.IsExpanded = false
	__ResourceShape__00000000_.X = 52.114853
	__ResourceShape__00000000_.Y = 34.127119
	__ResourceShape__00000000_.Width = 250.000000
	__ResourceShape__00000000_.Height = 70.000000

	__ResourceShape__00000001_.Name = `-Default Diagram`
	__ResourceShape__00000001_.IsExpanded = false
	__ResourceShape__00000001_.X = 34.180389
	__ResourceShape__00000001_.Y = 553.010316
	__ResourceShape__00000001_.Width = 250.000000
	__ResourceShape__00000001_.Height = 70.000000

	__ResourceShape__00000002_.Name = `-Default Diagram`
	__ResourceShape__00000002_.IsExpanded = false
	__ResourceShape__00000002_.X = 33.201744
	__ResourceShape__00000002_.Y = 650.719328
	__ResourceShape__00000002_.Width = 250.000000
	__ResourceShape__00000002_.Height = 70.000000

	__ResourceTaskShape__00000000_.Name = `PITProgram Investigation Team (PIT) to Mishap investigation`
	__ResourceTaskShape__00000000_.StartRatio = 0.500000
	__ResourceTaskShape__00000000_.EndRatio = 0.500000
	__ResourceTaskShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceTaskShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceTaskShape__00000000_.CornerOffsetRatio = 1.680000

	__Root__00000000_.Name = `Root`
	__Root__00000000_.NbPixPerCharacter = 8.000000

	__Task__00000000_.Name = `Starliner Crewed Flight Test (CFT)`
	__Task__00000000_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2024-06-05 00:00:00 +0000 UTC")
	__Task__00000000_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2024-06-10 00:00:00 +0000 UTC")
	__Task__00000000_.Description = `The mission of interest to the report.`
	__Task__00000000_.IsExpanded = false
	__Task__00000000_.ComputedPrefix = `2.1`
	__Task__00000000_.IsInRenameMode = false
	__Task__00000000_.IsInputsNodeExpanded = false
	__Task__00000000_.IsOutputsNodeExpanded = false
	__Task__00000000_.IsWithCompletion = false
	__Task__00000000_.Completion = ""

	__Task__00000001_.Name = `Mishap investigation`
	__Task__00000001_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2025-02-01 00:00:00 +0000 UTC")
	__Task__00000001_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2025-11-01 00:00:00 +0000 UTC")
	__Task__00000001_.Description = ``
	__Task__00000001_.IsExpanded = false
	__Task__00000001_.ComputedPrefix = `1`
	__Task__00000001_.IsInRenameMode = false
	__Task__00000001_.IsInputsNodeExpanded = false
	__Task__00000001_.IsOutputsNodeExpanded = false
	__Task__00000001_.IsWithCompletion = false
	__Task__00000001_.Completion = ""

	__Task__00000002_.Name = `Commercial Crew Program (CCP),`
	__Task__00000002_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000002_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000002_.Description = ``
	__Task__00000002_.IsExpanded = false
	__Task__00000002_.ComputedPrefix = `2`
	__Task__00000002_.IsInRenameMode = false
	__Task__00000002_.IsInputsNodeExpanded = false
	__Task__00000002_.IsOutputsNodeExpanded = false
	__Task__00000002_.IsWithCompletion = false
	__Task__00000002_.Completion = ""

	__Task__00000003_.Name = ` Commercial ReSupply (CRS) `
	__Task__00000003_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000003_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000003_.Description = ``
	__Task__00000003_.IsExpanded = false
	__Task__00000003_.ComputedPrefix = `3`
	__Task__00000003_.IsInRenameMode = false
	__Task__00000003_.IsInputsNodeExpanded = false
	__Task__00000003_.IsOutputsNodeExpanded = false
	__Task__00000003_.IsWithCompletion = false
	__Task__00000003_.Completion = ""

	__TaskCompositionShape__00000000_.Name = `Commercial Crew Program (CCP), to Starliner Crewed Flight Test (CFT)`
	__TaskCompositionShape__00000000_.StartRatio = 0.500000
	__TaskCompositionShape__00000000_.EndRatio = 0.500000
	__TaskCompositionShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000000_.CornerOffsetRatio = 1.680000

	__TaskCompositionShape__00000001_.Name = `Commercial Crew Program (CCP), to Starliner Crewed Flight Test (CFT)`
	__TaskCompositionShape__00000001_.StartRatio = 0.500000
	__TaskCompositionShape__00000001_.EndRatio = 0.500000
	__TaskCompositionShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000001_.CornerOffsetRatio = 1.680000

	__TaskShape__00000000_.Name = `NewTask-Default Diagram`
	__TaskShape__00000000_.IsExpanded = false
	__TaskShape__00000000_.X = 61.409529
	__TaskShape__00000000_.Y = 952.589880
	__TaskShape__00000000_.Width = 250.000000
	__TaskShape__00000000_.Height = 70.000000

	__TaskShape__00000001_.Name = `NewTask-Default Diagram`
	__TaskShape__00000001_.IsExpanded = false
	__TaskShape__00000001_.X = 105.187127
	__TaskShape__00000001_.Y = 246.958864
	__TaskShape__00000001_.Width = 250.000000
	__TaskShape__00000001_.Height = 70.000000

	__TaskShape__00000002_.Name = `-Default Diagram`
	__TaskShape__00000002_.IsExpanded = false
	__TaskShape__00000002_.X = 62.073018
	__TaskShape__00000002_.Y = 795.651222
	__TaskShape__00000002_.Width = 250.000000
	__TaskShape__00000002_.Height = 70.000000

	__TaskShape__00000003_.Name = `Mishap investigation-WBS`
	__TaskShape__00000003_.IsExpanded = false
	__TaskShape__00000003_.X = 28.416829
	__TaskShape__00000003_.Y = 546.674065
	__TaskShape__00000003_.Width = 250.000000
	__TaskShape__00000003_.Height = 70.000000

	__TaskShape__00000004_.Name = `Commercial Crew Program (CCP),-WBS`
	__TaskShape__00000004_.IsExpanded = false
	__TaskShape__00000004_.X = 16.071983
	__TaskShape__00000004_.Y = 166.517160
	__TaskShape__00000004_.Width = 250.000000
	__TaskShape__00000004_.Height = 70.000000

	__TaskShape__00000005_.Name = `Starliner Crewed Flight Test (CFT)-WBS`
	__TaskShape__00000005_.IsExpanded = false
	__TaskShape__00000005_.X = 16.617766
	__TaskShape__00000005_.Y = 336.295606
	__TaskShape__00000005_.Width = 250.000000
	__TaskShape__00000005_.Height = 70.000000

	__TaskShape__00000006_.Name = `-WBS`
	__TaskShape__00000006_.IsExpanded = false
	__TaskShape__00000006_.X = 14.409715
	__TaskShape__00000006_.Y = 32.470060
	__TaskShape__00000006_.Width = 250.000000
	__TaskShape__00000006_.Height = 70.000000

	// insertion point for setup of pointers
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000000_)
	__Diagram__00000000_.Task_Shapes = append(__Diagram__00000000_.Task_Shapes, __TaskShape__00000000_)
	__Diagram__00000000_.Task_Shapes = append(__Diagram__00000000_.Task_Shapes, __TaskShape__00000001_)
	__Diagram__00000000_.Task_Shapes = append(__Diagram__00000000_.Task_Shapes, __TaskShape__00000002_)
	__Diagram__00000000_.TasksWhoseNodeIsExpanded = append(__Diagram__00000000_.TasksWhoseNodeIsExpanded, __Task__00000002_)
	__Diagram__00000000_.TaskComposition_Shapes = append(__Diagram__00000000_.TaskComposition_Shapes, __TaskCompositionShape__00000000_)
	__Diagram__00000000_.Note_Shapes = append(__Diagram__00000000_.Note_Shapes, __NoteShape__00000000_)
	__Diagram__00000000_.NoteTaskShapes = append(__Diagram__00000000_.NoteTaskShapes, __NoteTaskShape__00000000_)
	__Diagram__00000000_.Resource_Shapes = append(__Diagram__00000000_.Resource_Shapes, __ResourceShape__00000000_)
	__Diagram__00000000_.Resource_Shapes = append(__Diagram__00000000_.Resource_Shapes, __ResourceShape__00000001_)
	__Diagram__00000000_.Resource_Shapes = append(__Diagram__00000000_.Resource_Shapes, __ResourceShape__00000002_)
	__Diagram__00000000_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000000_.ResourcesWhoseNodeIsExpanded, __Resource__00000000_)
	__Diagram__00000000_.ResourceTaskShapes = append(__Diagram__00000000_.ResourceTaskShapes, __ResourceTaskShape__00000000_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000003_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000004_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000005_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000006_)
	__Diagram__00000001_.TasksWhoseNodeIsExpanded = append(__Diagram__00000001_.TasksWhoseNodeIsExpanded, __Task__00000002_)
	__Diagram__00000001_.TaskComposition_Shapes = append(__Diagram__00000001_.TaskComposition_Shapes, __TaskCompositionShape__00000001_)
	__Diagram__00000002_.Product_Shapes = append(__Diagram__00000002_.Product_Shapes, __ProductShape__00000001_)
	__Diagram__00000002_.Product_Shapes = append(__Diagram__00000002_.Product_Shapes, __ProductShape__00000002_)
	__Diagram__00000002_.Product_Shapes = append(__Diagram__00000002_.Product_Shapes, __ProductShape__00000003_)
	__Note__00000000_.Tasks = append(__Note__00000000_.Tasks, __Task__00000000_)
	__NoteShape__00000000_.Note = __Note__00000000_
	__NoteTaskShape__00000000_.Note = __Note__00000000_
	__NoteTaskShape__00000000_.Task = __Task__00000000_
	__ProductShape__00000000_.Product = __Product__00000000_
	__ProductShape__00000001_.Product = __Product__00000000_
	__ProductShape__00000002_.Product = __Product__00000001_
	__ProductShape__00000003_.Product = __Product__00000002_
	__Project__00000000_.RootProducts = append(__Project__00000000_.RootProducts, __Product__00000000_)
	__Project__00000000_.RootProducts = append(__Project__00000000_.RootProducts, __Product__00000001_)
	__Project__00000000_.RootProducts = append(__Project__00000000_.RootProducts, __Product__00000002_)
	__Project__00000000_.RootTasks = append(__Project__00000000_.RootTasks, __Task__00000001_)
	__Project__00000000_.RootTasks = append(__Project__00000000_.RootTasks, __Task__00000002_)
	__Project__00000000_.RootTasks = append(__Project__00000000_.RootTasks, __Task__00000003_)
	__Project__00000000_.RootResources = append(__Project__00000000_.RootResources, __Resource__00000000_)
	__Project__00000000_.RootResources = append(__Project__00000000_.RootResources, __Resource__00000001_)
	__Project__00000000_.RootResources = append(__Project__00000000_.RootResources, __Resource__00000002_)
	__Project__00000000_.Notes = append(__Project__00000000_.Notes, __Note__00000000_)
	__Project__00000000_.Diagrams = append(__Project__00000000_.Diagrams, __Diagram__00000000_)
	__Project__00000000_.Diagrams = append(__Project__00000000_.Diagrams, __Diagram__00000001_)
	__Project__00000000_.Diagrams = append(__Project__00000000_.Diagrams, __Diagram__00000002_)
	__Resource__00000000_.Tasks = append(__Resource__00000000_.Tasks, __Task__00000001_)
	__ResourceShape__00000000_.Resource = __Resource__00000000_
	__ResourceShape__00000001_.Resource = __Resource__00000001_
	__ResourceShape__00000002_.Resource = __Resource__00000002_
	__ResourceTaskShape__00000000_.Resource = __Resource__00000000_
	__ResourceTaskShape__00000000_.Task = __Task__00000001_
	__Root__00000000_.Projects = append(__Root__00000000_.Projects, __Project__00000000_)
	__Task__00000002_.SubTasks = append(__Task__00000002_.SubTasks, __Task__00000000_)
	__TaskCompositionShape__00000000_.Task = __Task__00000000_
	__TaskCompositionShape__00000001_.Task = __Task__00000000_
	__TaskShape__00000000_.Task = __Task__00000000_
	__TaskShape__00000001_.Task = __Task__00000001_
	__TaskShape__00000002_.Task = __Task__00000002_
	__TaskShape__00000003_.Task = __Task__00000001_
	__TaskShape__00000004_.Task = __Task__00000002_
	__TaskShape__00000005_.Task = __Task__00000000_
	__TaskShape__00000006_.Task = __Task__00000003_
}
