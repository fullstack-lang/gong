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

	__Product__00000000_ := (&models.Product{Name: `P1`}).Stage(stage)
	__Product__00000001_ := (&models.Product{Name: `P2`}).Stage(stage)

	__ProductShape__00000000_ := (&models.ProductShape{Name: `NewProduct-Default Diagram`}).Stage(stage)
	__ProductShape__00000001_ := (&models.ProductShape{Name: `NewProduct-Default Diagram`}).Stage(stage)

	__Project__00000000_ := (&models.Project{Name: `NewProject`}).Stage(stage)

	__Root__00000000_ := (&models.Root{Name: `Root`}).Stage(stage)

	__Task__00000000_ := (&models.Task{Name: `T1`}).Stage(stage)

	__TaskOutputShape__00000000_ := (&models.TaskOutputShape{Name: `T1 to P1`}).Stage(stage)

	__TaskShape__00000000_ := (&models.TaskShape{Name: `NewTask-Default Diagram`}).Stage(stage)

	// insertion point for initialization of values

	__Diagram__00000000_.Name = `Default Diagram`
	__Diagram__00000000_.IsChecked = true
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.IsInRenameMode = false
	__Diagram__00000000_.ShowPrefix = false
	__Diagram__00000000_.DefaultBoxWidth = 250.000000
	__Diagram__00000000_.DefaultBoxHeigth = 100.000000
	__Diagram__00000000_.IsExpanded = true
	__Diagram__00000000_.ComputedPrefix = ``
	__Diagram__00000000_.IsPBSNodeExpanded = true
	__Diagram__00000000_.IsWBSNodeExpanded = true
	__Diagram__00000000_.IsNotesNodeExpanded = false

	__Product__00000000_.Name = `P1`
	__Product__00000000_.Description = ``
	__Product__00000000_.IsExpanded = false
	__Product__00000000_.ComputedPrefix = `1`
	__Product__00000000_.IsProducersNodeExpanded = false
	__Product__00000000_.IsConsumersNodeExpanded = false

	__Product__00000001_.Name = `P2`
	__Product__00000001_.Description = ``
	__Product__00000001_.IsExpanded = false
	__Product__00000001_.ComputedPrefix = `2`
	__Product__00000001_.IsProducersNodeExpanded = false
	__Product__00000001_.IsConsumersNodeExpanded = false

	__ProductShape__00000000_.Name = `NewProduct-Default Diagram`
	__ProductShape__00000000_.IsExpanded = false
	__ProductShape__00000000_.X = 636.468244
	__ProductShape__00000000_.Y = 83.048137
	__ProductShape__00000000_.Width = 250.000000
	__ProductShape__00000000_.Height = 100.000000

	__ProductShape__00000001_.Name = `NewProduct-Default Diagram`
	__ProductShape__00000001_.IsExpanded = false
	__ProductShape__00000001_.X = 47.503241
	__ProductShape__00000001_.Y = 296.651901
	__ProductShape__00000001_.Width = 250.000000
	__ProductShape__00000001_.Height = 100.000000

	__Project__00000000_.Name = `NewProject`
	__Project__00000000_.IsExpanded = true
	__Project__00000000_.ComputedPrefix = ``

	__Root__00000000_.Name = `Root`
	__Root__00000000_.NbPixPerCharacter = 8.000000

	__Task__00000000_.Name = `T1`
	__Task__00000000_.Description = ``
	__Task__00000000_.IsExpanded = false
	__Task__00000000_.ComputedPrefix = `1`
	__Task__00000000_.IsInputsNodeExpanded = false
	__Task__00000000_.IsOutputsNodeExpanded = false
	__Task__00000000_.IsWithCompletion = false
	__Task__00000000_.Completion = ""

	__TaskOutputShape__00000000_.Name = `T1 to P1`
	__TaskOutputShape__00000000_.StartRatio = 0.500000
	__TaskOutputShape__00000000_.EndRatio = 0.500000
	__TaskOutputShape__00000000_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000000_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000000_.CornerOffsetRatio = 1.273106

	__TaskShape__00000000_.Name = `NewTask-Default Diagram`
	__TaskShape__00000000_.IsExpanded = false
	__TaskShape__00000000_.X = 262.994393
	__TaskShape__00000000_.Y = 111.277418
	__TaskShape__00000000_.Width = 250.000000
	__TaskShape__00000000_.Height = 100.000000

	// insertion point for setup of pointers
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000000_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000001_)
	__Diagram__00000000_.Task_Shapes = append(__Diagram__00000000_.Task_Shapes, __TaskShape__00000000_)
	__Diagram__00000000_.TasksWhoseNodeIsExpanded = append(__Diagram__00000000_.TasksWhoseNodeIsExpanded, __Task__00000000_)
	__Diagram__00000000_.TasksWhoseOutputNodeIsExpanded = append(__Diagram__00000000_.TasksWhoseOutputNodeIsExpanded, __Task__00000000_)
	__Diagram__00000000_.TaskOutputShapes = append(__Diagram__00000000_.TaskOutputShapes, __TaskOutputShape__00000000_)
	__ProductShape__00000000_.Product = __Product__00000000_
	__ProductShape__00000001_.Product = __Product__00000001_
	__Project__00000000_.RootProducts = append(__Project__00000000_.RootProducts, __Product__00000000_)
	__Project__00000000_.RootProducts = append(__Project__00000000_.RootProducts, __Product__00000001_)
	__Project__00000000_.RootTasks = append(__Project__00000000_.RootTasks, __Task__00000000_)
	__Project__00000000_.Diagrams = append(__Project__00000000_.Diagrams, __Diagram__00000000_)
	__Root__00000000_.Projects = append(__Root__00000000_.Projects, __Project__00000000_)
	__Task__00000000_.Outputs = append(__Task__00000000_.Outputs, __Product__00000000_)
	__Task__00000000_.Outputs = append(__Task__00000000_.Outputs, __Product__00000001_)
	__TaskOutputShape__00000000_.Task = __Task__00000000_
	__TaskOutputShape__00000000_.Product = __Product__00000000_
	__TaskShape__00000000_.Task = __Task__00000000_
}
