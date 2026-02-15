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

	__Note__00000000_ := (&models.Note{Name: `N1`}).Stage(stage)

	__NoteProductShape__00000000_ := (&models.NoteProductShape{Name: `N1 to P1`}).Stage(stage)

	__NoteShape__00000000_ := (&models.NoteShape{Name: `NewNote-Default Diagram`}).Stage(stage)

	__NoteTaskShape__00000000_ := (&models.NoteTaskShape{Name: `N1 to T1`}).Stage(stage)

	__Product__00000000_ := (&models.Product{Name: `P1`}).Stage(stage)
	__Product__00000001_ := (&models.Product{Name: `P1.1`}).Stage(stage)
	__Product__00000002_ := (&models.Product{Name: `P1.2`}).Stage(stage)

	__ProductCompositionShape__00000000_ := (&models.ProductCompositionShape{Name: `P1 to P1.1`}).Stage(stage)
	__ProductCompositionShape__00000001_ := (&models.ProductCompositionShape{Name: `P1 to NewProduct`}).Stage(stage)

	__ProductShape__00000000_ := (&models.ProductShape{Name: `NewProduct-Default Diagram`}).Stage(stage)
	__ProductShape__00000001_ := (&models.ProductShape{Name: `NewProduct-Default Diagram`}).Stage(stage)
	__ProductShape__00000002_ := (&models.ProductShape{Name: `NewProduct-Default Diagram`}).Stage(stage)

	__Project__00000000_ := (&models.Project{Name: `NewProject`}).Stage(stage)

	__Resource__00000000_ := (&models.Resource{Name: `R1`}).Stage(stage)
	__Resource__00000001_ := (&models.Resource{Name: `R1.1`}).Stage(stage)
	__Resource__00000002_ := (&models.Resource{Name: `R1.2`}).Stage(stage)

	__ResourceCompositionShape__00000001_ := (&models.ResourceCompositionShape{Name: `R1 to R1.2`}).Stage(stage)
	__ResourceCompositionShape__00000002_ := (&models.ResourceCompositionShape{Name: `R1 to R1.1`}).Stage(stage)

	__ResourceShape__00000000_ := (&models.ResourceShape{Name: `NewResource-Default Diagram`}).Stage(stage)
	__ResourceShape__00000002_ := (&models.ResourceShape{Name: `NewResource-Default Diagram`}).Stage(stage)
	__ResourceShape__00000003_ := (&models.ResourceShape{Name: `R1.1-Default Diagram`}).Stage(stage)

	__ResourceTaskShape__00000000_ := (&models.ResourceTaskShape{Name: `R1.1 to T1.1`}).Stage(stage)

	__Root__00000000_ := (&models.Root{Name: `Root`}).Stage(stage)

	__Task__00000000_ := (&models.Task{Name: `T1`}).Stage(stage)
	__Task__00000001_ := (&models.Task{Name: `T1.1`}).Stage(stage)

	__TaskCompositionShape__00000000_ := (&models.TaskCompositionShape{Name: `T1 to NewTask`}).Stage(stage)

	__TaskInputShape__00000000_ := (&models.TaskInputShape{Name: `T1.1 to P1.2`}).Stage(stage)

	__TaskOutputShape__00000000_ := (&models.TaskOutputShape{Name: `T1.1 to P1.1`}).Stage(stage)

	__TaskShape__00000000_ := (&models.TaskShape{Name: `NewTask-Default Diagram`}).Stage(stage)
	__TaskShape__00000001_ := (&models.TaskShape{Name: `NewTask-Default Diagram`}).Stage(stage)

	// insertion point for initialization of values

	__Diagram__00000000_.Name = `Default Diagram`
	__Diagram__00000000_.IsChecked = true
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.IsInRenameMode = false
	__Diagram__00000000_.ShowPrefix = true
	__Diagram__00000000_.DefaultBoxWidth = 250.000000
	__Diagram__00000000_.DefaultBoxHeigth = 100.000000
	__Diagram__00000000_.IsExpanded = true
	__Diagram__00000000_.ComputedPrefix = ``
	__Diagram__00000000_.IsPBSNodeExpanded = true
	__Diagram__00000000_.IsWBSNodeExpanded = true
	__Diagram__00000000_.IsNotesNodeExpanded = true
	__Diagram__00000000_.IsResourcesNodeExpanded = true

	__Note__00000000_.Name = `N1`
	__Note__00000000_.IsExpanded = false
	__Note__00000000_.ComputedPrefix = `1`

	__NoteProductShape__00000000_.Name = `N1 to P1`
	__NoteProductShape__00000000_.StartRatio = 0.500000
	__NoteProductShape__00000000_.EndRatio = 0.500000
	__NoteProductShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000000_.CornerOffsetRatio = 1.680000

	__NoteShape__00000000_.Name = `NewNote-Default Diagram`
	__NoteShape__00000000_.IsExpanded = false
	__NoteShape__00000000_.X = 357.643327
	__NoteShape__00000000_.Y = 111.847509
	__NoteShape__00000000_.Width = 250.000000
	__NoteShape__00000000_.Height = 50.000000

	__NoteTaskShape__00000000_.Name = `N1 to T1`
	__NoteTaskShape__00000000_.StartRatio = 0.500000
	__NoteTaskShape__00000000_.EndRatio = 0.500000
	__NoteTaskShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteTaskShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteTaskShape__00000000_.CornerOffsetRatio = 1.680000

	__Product__00000000_.Name = `P1`
	__Product__00000000_.Description = ``
	__Product__00000000_.IsExpanded = false
	__Product__00000000_.ComputedPrefix = `1`
	__Product__00000000_.IsProducersNodeExpanded = false
	__Product__00000000_.IsConsumersNodeExpanded = false

	__Product__00000001_.Name = `P1.1`
	__Product__00000001_.Description = ``
	__Product__00000001_.IsExpanded = false
	__Product__00000001_.ComputedPrefix = `1.1`
	__Product__00000001_.IsProducersNodeExpanded = false
	__Product__00000001_.IsConsumersNodeExpanded = false

	__Product__00000002_.Name = `P1.2`
	__Product__00000002_.Description = ``
	__Product__00000002_.IsExpanded = false
	__Product__00000002_.ComputedPrefix = `1.2`
	__Product__00000002_.IsProducersNodeExpanded = false
	__Product__00000002_.IsConsumersNodeExpanded = false

	__ProductCompositionShape__00000000_.Name = `P1 to P1.1`
	__ProductCompositionShape__00000000_.StartRatio = 0.500000
	__ProductCompositionShape__00000000_.EndRatio = 0.500000
	__ProductCompositionShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000000_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000001_.Name = `P1 to NewProduct`
	__ProductCompositionShape__00000001_.StartRatio = 0.500000
	__ProductCompositionShape__00000001_.EndRatio = 0.500000
	__ProductCompositionShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000001_.CornerOffsetRatio = 1.680000

	__ProductShape__00000000_.Name = `NewProduct-Default Diagram`
	__ProductShape__00000000_.IsExpanded = false
	__ProductShape__00000000_.X = 661.915440
	__ProductShape__00000000_.Y = 46.307227
	__ProductShape__00000000_.Width = 250.000000
	__ProductShape__00000000_.Height = 54.000000

	__ProductShape__00000001_.Name = `NewProduct-Default Diagram`
	__ProductShape__00000001_.IsExpanded = false
	__ProductShape__00000001_.X = 900.915440
	__ProductShape__00000001_.Y = 220.307227
	__ProductShape__00000001_.Width = 250.000000
	__ProductShape__00000001_.Height = 41.000000

	__ProductShape__00000002_.Name = `NewProduct-Default Diagram`
	__ProductShape__00000002_.IsExpanded = false
	__ProductShape__00000002_.X = 528.915440
	__ProductShape__00000002_.Y = 194.307227
	__ProductShape__00000002_.Width = 250.000000
	__ProductShape__00000002_.Height = 59.000000

	__Project__00000000_.Name = `NewProject`
	__Project__00000000_.IsExpanded = true
	__Project__00000000_.ComputedPrefix = ``

	__Resource__00000000_.Name = `R1`
	__Resource__00000000_.Description = ``
	__Resource__00000000_.IsExpanded = false
	__Resource__00000000_.ComputedPrefix = `1`

	__Resource__00000001_.Name = `R1.1`
	__Resource__00000001_.Description = ``
	__Resource__00000001_.IsExpanded = false
	__Resource__00000001_.ComputedPrefix = `1.1`

	__Resource__00000002_.Name = `R1.2`
	__Resource__00000002_.Description = ``
	__Resource__00000002_.IsExpanded = false
	__Resource__00000002_.ComputedPrefix = `1.2`

	__ResourceCompositionShape__00000001_.Name = `R1 to R1.2`
	__ResourceCompositionShape__00000001_.StartRatio = 0.235952
	__ResourceCompositionShape__00000001_.EndRatio = 0.440574
	__ResourceCompositionShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000001_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ResourceCompositionShape__00000001_.CornerOffsetRatio = 0.097994

	__ResourceCompositionShape__00000002_.Name = `R1 to R1.1`
	__ResourceCompositionShape__00000002_.StartRatio = 0.500000
	__ResourceCompositionShape__00000002_.EndRatio = 0.500000
	__ResourceCompositionShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000002_.CornerOffsetRatio = 1.680000

	__ResourceShape__00000000_.Name = `NewResource-Default Diagram`
	__ResourceShape__00000000_.IsExpanded = false
	__ResourceShape__00000000_.X = 25.501544
	__ResourceShape__00000000_.Y = 474.249957
	__ResourceShape__00000000_.Width = 250.000000
	__ResourceShape__00000000_.Height = 71.000000

	__ResourceShape__00000002_.Name = `NewResource-Default Diagram`
	__ResourceShape__00000002_.IsExpanded = false
	__ResourceShape__00000002_.X = 158.862056
	__ResourceShape__00000002_.Y = 703.845003
	__ResourceShape__00000002_.Width = 250.000000
	__ResourceShape__00000002_.Height = 73.000000

	__ResourceShape__00000003_.Name = `R1.1-Default Diagram`
	__ResourceShape__00000003_.IsExpanded = false
	__ResourceShape__00000003_.X = 269.043648
	__ResourceShape__00000003_.Y = 618.652577
	__ResourceShape__00000003_.Width = 250.000000
	__ResourceShape__00000003_.Height = 72.000000

	__ResourceTaskShape__00000000_.Name = `R1.1 to T1.1`
	__ResourceTaskShape__00000000_.StartRatio = 0.727144
	__ResourceTaskShape__00000000_.EndRatio = 1.000000
	__ResourceTaskShape__00000000_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__ResourceTaskShape__00000000_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ResourceTaskShape__00000000_.CornerOffsetRatio = 1.616679

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

	__Task__00000001_.Name = `T1.1`
	__Task__00000001_.Description = ``
	__Task__00000001_.IsExpanded = false
	__Task__00000001_.ComputedPrefix = `1.1`
	__Task__00000001_.IsInputsNodeExpanded = false
	__Task__00000001_.IsOutputsNodeExpanded = false
	__Task__00000001_.IsWithCompletion = true
	__Task__00000001_.Completion = models.PERCENT_050

	__TaskCompositionShape__00000000_.Name = `T1 to NewTask`
	__TaskCompositionShape__00000000_.StartRatio = 0.500000
	__TaskCompositionShape__00000000_.EndRatio = 0.500000
	__TaskCompositionShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000000_.CornerOffsetRatio = 1.680000

	__TaskInputShape__00000000_.Name = `T1.1 to P1.2`
	__TaskInputShape__00000000_.StartRatio = 0.500000
	__TaskInputShape__00000000_.EndRatio = 0.806434
	__TaskInputShape__00000000_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__TaskInputShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskInputShape__00000000_.CornerOffsetRatio = 1.944063

	__TaskOutputShape__00000000_.Name = `T1.1 to P1.1`
	__TaskOutputShape__00000000_.StartRatio = 0.500000
	__TaskOutputShape__00000000_.EndRatio = 0.405192
	__TaskOutputShape__00000000_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskOutputShape__00000000_.CornerOffsetRatio = -0.577653

	__TaskShape__00000000_.Name = `NewTask-Default Diagram`
	__TaskShape__00000000_.IsExpanded = false
	__TaskShape__00000000_.X = 66.604908
	__TaskShape__00000000_.Y = 27.374505
	__TaskShape__00000000_.Width = 250.000000
	__TaskShape__00000000_.Height = 58.000000

	__TaskShape__00000001_.Name = `NewTask-Default Diagram`
	__TaskShape__00000001_.IsExpanded = false
	__TaskShape__00000001_.X = 130.604908
	__TaskShape__00000001_.Y = 309.374505
	__TaskShape__00000001_.Width = 250.000000
	__TaskShape__00000001_.Height = 82.000000

	// insertion point for setup of pointers
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000000_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000001_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000002_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, __Product__00000000_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000000_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000001_)
	__Diagram__00000000_.Task_Shapes = append(__Diagram__00000000_.Task_Shapes, __TaskShape__00000000_)
	__Diagram__00000000_.Task_Shapes = append(__Diagram__00000000_.Task_Shapes, __TaskShape__00000001_)
	__Diagram__00000000_.TasksWhoseNodeIsExpanded = append(__Diagram__00000000_.TasksWhoseNodeIsExpanded, __Task__00000000_)
	__Diagram__00000000_.TasksWhoseNodeIsExpanded = append(__Diagram__00000000_.TasksWhoseNodeIsExpanded, __Task__00000001_)
	__Diagram__00000000_.TasksWhoseOutputNodeIsExpanded = append(__Diagram__00000000_.TasksWhoseOutputNodeIsExpanded, __Task__00000001_)
	__Diagram__00000000_.TaskComposition_Shapes = append(__Diagram__00000000_.TaskComposition_Shapes, __TaskCompositionShape__00000000_)
	__Diagram__00000000_.TaskInputShapes = append(__Diagram__00000000_.TaskInputShapes, __TaskInputShape__00000000_)
	__Diagram__00000000_.TaskOutputShapes = append(__Diagram__00000000_.TaskOutputShapes, __TaskOutputShape__00000000_)
	__Diagram__00000000_.Note_Shapes = append(__Diagram__00000000_.Note_Shapes, __NoteShape__00000000_)
	__Diagram__00000000_.NotesWhoseNodeIsExpanded = append(__Diagram__00000000_.NotesWhoseNodeIsExpanded, __Note__00000000_)
	__Diagram__00000000_.NoteProductShapes = append(__Diagram__00000000_.NoteProductShapes, __NoteProductShape__00000000_)
	__Diagram__00000000_.NoteTaskShapes = append(__Diagram__00000000_.NoteTaskShapes, __NoteTaskShape__00000000_)
	__Diagram__00000000_.Resource_Shapes = append(__Diagram__00000000_.Resource_Shapes, __ResourceShape__00000000_)
	__Diagram__00000000_.Resource_Shapes = append(__Diagram__00000000_.Resource_Shapes, __ResourceShape__00000002_)
	__Diagram__00000000_.Resource_Shapes = append(__Diagram__00000000_.Resource_Shapes, __ResourceShape__00000003_)
	__Diagram__00000000_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000000_.ResourcesWhoseNodeIsExpanded, __Resource__00000000_)
	__Diagram__00000000_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000000_.ResourcesWhoseNodeIsExpanded, __Resource__00000001_)
	__Diagram__00000000_.ResourceComposition_Shapes = append(__Diagram__00000000_.ResourceComposition_Shapes, __ResourceCompositionShape__00000001_)
	__Diagram__00000000_.ResourceComposition_Shapes = append(__Diagram__00000000_.ResourceComposition_Shapes, __ResourceCompositionShape__00000002_)
	__Diagram__00000000_.ResourceTaskShapes = append(__Diagram__00000000_.ResourceTaskShapes, __ResourceTaskShape__00000000_)
	__Note__00000000_.Products = append(__Note__00000000_.Products, __Product__00000000_)
	__Note__00000000_.Tasks = append(__Note__00000000_.Tasks, __Task__00000000_)
	__NoteProductShape__00000000_.Note = __Note__00000000_
	__NoteProductShape__00000000_.Product = __Product__00000000_
	__NoteShape__00000000_.Note = __Note__00000000_
	__NoteTaskShape__00000000_.Note = __Note__00000000_
	__NoteTaskShape__00000000_.Task = __Task__00000000_
	__Product__00000000_.SubProducts = append(__Product__00000000_.SubProducts, __Product__00000001_)
	__Product__00000000_.SubProducts = append(__Product__00000000_.SubProducts, __Product__00000002_)
	__ProductCompositionShape__00000000_.Product = __Product__00000001_
	__ProductCompositionShape__00000001_.Product = __Product__00000002_
	__ProductShape__00000000_.Product = __Product__00000000_
	__ProductShape__00000001_.Product = __Product__00000001_
	__ProductShape__00000002_.Product = __Product__00000002_
	__Project__00000000_.RootProducts = append(__Project__00000000_.RootProducts, __Product__00000000_)
	__Project__00000000_.RootTasks = append(__Project__00000000_.RootTasks, __Task__00000000_)
	__Project__00000000_.RootResources = append(__Project__00000000_.RootResources, __Resource__00000000_)
	__Project__00000000_.Notes = append(__Project__00000000_.Notes, __Note__00000000_)
	__Project__00000000_.Diagrams = append(__Project__00000000_.Diagrams, __Diagram__00000000_)
	__Resource__00000000_.SubResources = append(__Resource__00000000_.SubResources, __Resource__00000001_)
	__Resource__00000000_.SubResources = append(__Resource__00000000_.SubResources, __Resource__00000002_)
	__Resource__00000001_.Tasks = append(__Resource__00000001_.Tasks, __Task__00000001_)
	__ResourceCompositionShape__00000001_.Resource = __Resource__00000002_
	__ResourceCompositionShape__00000002_.Resource = __Resource__00000001_
	__ResourceShape__00000000_.Resource = __Resource__00000000_
	__ResourceShape__00000002_.Resource = __Resource__00000002_
	__ResourceShape__00000003_.Resource = __Resource__00000001_
	__ResourceTaskShape__00000000_.Resource = __Resource__00000001_
	__ResourceTaskShape__00000000_.Task = __Task__00000001_
	__Root__00000000_.Projects = append(__Root__00000000_.Projects, __Project__00000000_)
	__Task__00000000_.SubTasks = append(__Task__00000000_.SubTasks, __Task__00000001_)
	__Task__00000001_.Inputs = append(__Task__00000001_.Inputs, __Product__00000002_)
	__Task__00000001_.Outputs = append(__Task__00000001_.Outputs, __Product__00000001_)
	__TaskCompositionShape__00000000_.Task = __Task__00000001_
	__TaskInputShape__00000000_.Product = __Product__00000002_
	__TaskInputShape__00000000_.Task = __Task__00000001_
	__TaskOutputShape__00000000_.Task = __Task__00000001_
	__TaskOutputShape__00000000_.Product = __Product__00000001_
	__TaskShape__00000000_.Task = __Task__00000000_
	__TaskShape__00000001_.Task = __Task__00000001_
}
