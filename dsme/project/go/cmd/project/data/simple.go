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

	__Product__00000000_ := (&models.Product{Name: `P1.12`}).Stage(stage)

	__ProductShape__00000000_ := (&models.ProductShape{Name: `-D1`}).Stage(stage)

	__Project__00000000_ := (&models.Project{Name: `NewProject`}).Stage(stage)

	__Resource__00000002_ := (&models.Resource{Name: `R1`}).Stage(stage)

	__ResourceShape__00000002_ := (&models.ResourceShape{Name: `D1-R1.2`}).Stage(stage)

	__ResourceTaskShape__00000000_ := (&models.ResourceTaskShape{Name: `R1 to W1`}).Stage(stage)

	__Root__00000000_ := (&models.Root{Name: `Root`}).Stage(stage)

	__Task__00000000_ := (&models.Task{Name: `W1.1`}).Stage(stage)

	__TaskOutputShape__00000000_ := (&models.TaskOutputShape{Name: `W1 to P1`}).Stage(stage)

	__TaskShape__00000000_ := (&models.TaskShape{Name: `-D1`}).Stage(stage)

	// insertion point for initialization of values

	__Diagram__00000000_.Name = `D1`
	__Diagram__00000000_.IsChecked = true
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.ShowPrefix = true
	__Diagram__00000000_.DefaultBoxWidth = 250.000000
	__Diagram__00000000_.DefaultBoxHeigth = 100.000000
	__Diagram__00000000_.Width = 6108.862056
	__Diagram__00000000_.Height = 6476.845003
	__Diagram__00000000_.IsExpanded = true
	__Diagram__00000000_.ComputedPrefix = ``
	__Diagram__00000000_.IsInRenameMode = false
	__Diagram__00000000_.IsPBSNodeExpanded = true
	__Diagram__00000000_.IsWBSNodeExpanded = true
	__Diagram__00000000_.IsNotesNodeExpanded = false
	__Diagram__00000000_.IsResourcesNodeExpanded = true

	__Product__00000000_.Name = `P1.12`
	__Product__00000000_.Description = ``
	__Product__00000000_.IsExpanded = false
	__Product__00000000_.ComputedPrefix = `1`
	__Product__00000000_.IsInRenameMode = false
	__Product__00000000_.IsProducersNodeExpanded = false
	__Product__00000000_.IsConsumersNodeExpanded = false

	__ProductShape__00000000_.Name = `-D1`
	__ProductShape__00000000_.IsExpanded = false
	__ProductShape__00000000_.X = 871.835096
	__ProductShape__00000000_.Y = 295.442323
	__ProductShape__00000000_.Width = 250.000000
	__ProductShape__00000000_.Height = 100.000000
	__ProductShape__00000000_.IsHidden = false

	__Project__00000000_.Name = `NewProject`
	__Project__00000000_.IsExpanded = true
	__Project__00000000_.ComputedPrefix = ``
	__Project__00000000_.IsInRenameMode = false

	__Resource__00000002_.Name = `R1`
	__Resource__00000002_.Description = ``
	__Resource__00000002_.IsExpanded = false
	__Resource__00000002_.ComputedPrefix = `1`
	__Resource__00000002_.IsInRenameMode = false

	__ResourceShape__00000002_.Name = `D1-R1.2`
	__ResourceShape__00000002_.IsExpanded = false
	__ResourceShape__00000002_.X = 46.862056
	__ResourceShape__00000002_.Y = 303.845003
	__ResourceShape__00000002_.Width = 250.000000
	__ResourceShape__00000002_.Height = 73.000000
	__ResourceShape__00000002_.IsHidden = false

	__ResourceTaskShape__00000000_.Name = `R1 to W1`
	__ResourceTaskShape__00000000_.StartRatio = 0.563847
	__ResourceTaskShape__00000000_.EndRatio = 0.476371
	__ResourceTaskShape__00000000_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__ResourceTaskShape__00000000_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ResourceTaskShape__00000000_.CornerOffsetRatio = 1.236942

	__Root__00000000_.Name = `Root`
	__Root__00000000_.NbPixPerCharacter = 8.000000

	__Task__00000000_.Name = `W1.1`
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

	__TaskOutputShape__00000000_.Name = `W1 to P1`
	__TaskOutputShape__00000000_.StartRatio = 0.500000
	__TaskOutputShape__00000000_.EndRatio = 0.500000
	__TaskOutputShape__00000000_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000000_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000000_.CornerOffsetRatio = 1.177566

	__TaskShape__00000000_.Name = `-D1`
	__TaskShape__00000000_.IsExpanded = false
	__TaskShape__00000000_.X = 442.706056
	__TaskShape__00000000_.Y = 300.368751
	__TaskShape__00000000_.Width = 250.000000
	__TaskShape__00000000_.Height = 100.000000
	__TaskShape__00000000_.IsHidden = false

	// insertion point for setup of pointers
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000000_)
	__Diagram__00000000_.Task_Shapes = append(__Diagram__00000000_.Task_Shapes, __TaskShape__00000000_)
	__Diagram__00000000_.TaskOutputShapes = append(__Diagram__00000000_.TaskOutputShapes, __TaskOutputShape__00000000_)
	__Diagram__00000000_.Resource_Shapes = append(__Diagram__00000000_.Resource_Shapes, __ResourceShape__00000002_)
	__Diagram__00000000_.ResourceTaskShapes = append(__Diagram__00000000_.ResourceTaskShapes, __ResourceTaskShape__00000000_)
	__ProductShape__00000000_.Product = __Product__00000000_
	__Project__00000000_.RootProducts = append(__Project__00000000_.RootProducts, __Product__00000000_)
	__Project__00000000_.RootTasks = append(__Project__00000000_.RootTasks, __Task__00000000_)
	__Project__00000000_.RootResources = append(__Project__00000000_.RootResources, __Resource__00000002_)
	__Project__00000000_.Diagrams = append(__Project__00000000_.Diagrams, __Diagram__00000000_)
	__Resource__00000002_.Tasks = append(__Resource__00000002_.Tasks, __Task__00000000_)
	__ResourceShape__00000002_.Resource = __Resource__00000002_
	__ResourceTaskShape__00000000_.Resource = __Resource__00000002_
	__ResourceTaskShape__00000000_.Task = __Task__00000000_
	__Root__00000000_.Projects = append(__Root__00000000_.Projects, __Project__00000000_)
	__Task__00000000_.Outputs = append(__Task__00000000_.Outputs, __Product__00000000_)
	__TaskOutputShape__00000000_.Task = __Task__00000000_
	__TaskOutputShape__00000000_.Product = __Product__00000000_
	__TaskShape__00000000_.Task = __Task__00000000_
}
