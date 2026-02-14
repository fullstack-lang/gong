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

	__Note__00000000_ := (&models.Note{Name: `NewNote`}).Stage(stage)

	__NoteShape__00000000_ := (&models.NoteShape{Name: `NewNote-Default Diagram`}).Stage(stage)

	__Product__00000000_ := (&models.Product{Name: `NewProduct`}).Stage(stage)

	__ProductShape__00000000_ := (&models.ProductShape{Name: `NewProduct-Default Diagram`}).Stage(stage)

	__Project__00000000_ := (&models.Project{Name: `NewProject`}).Stage(stage)

	__Resource__00000000_ := (&models.Resource{Name: `NewResource`}).Stage(stage)

	__ResourceShape__00000000_ := (&models.ResourceShape{Name: `NewResource-Default Diagram`}).Stage(stage)

	__Root__00000000_ := (&models.Root{Name: `Root`}).Stage(stage)

	__Task__00000000_ := (&models.Task{Name: `NewTask`}).Stage(stage)

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
	__Diagram__00000000_.IsNotesNodeExpanded = true
	__Diagram__00000000_.IsResourcesNodeExpanded = true

	__Note__00000000_.Name = `NewNote`
	__Note__00000000_.IsExpanded = false
	__Note__00000000_.ComputedPrefix = `1`

	__NoteShape__00000000_.Name = `NewNote-Default Diagram`
	__NoteShape__00000000_.IsExpanded = false
	__NoteShape__00000000_.X = 120.497483
	__NoteShape__00000000_.Y = 132.852682
	__NoteShape__00000000_.Width = 250.000000
	__NoteShape__00000000_.Height = 100.000000

	__Product__00000000_.Name = `NewProduct`
	__Product__00000000_.Description = ``
	__Product__00000000_.IsExpanded = false
	__Product__00000000_.ComputedPrefix = `1`
	__Product__00000000_.IsProducersNodeExpanded = false
	__Product__00000000_.IsConsumersNodeExpanded = false

	__ProductShape__00000000_.Name = `NewProduct-Default Diagram`
	__ProductShape__00000000_.IsExpanded = false
	__ProductShape__00000000_.X = 183.915471
	__ProductShape__00000000_.Y = 100.307227
	__ProductShape__00000000_.Width = 250.000000
	__ProductShape__00000000_.Height = 100.000000

	__Project__00000000_.Name = `NewProject`
	__Project__00000000_.IsExpanded = true
	__Project__00000000_.ComputedPrefix = ``

	__Resource__00000000_.Name = `NewResource`
	__Resource__00000000_.Description = ``
	__Resource__00000000_.IsExpanded = false
	__Resource__00000000_.ComputedPrefix = `1`

	__ResourceShape__00000000_.Name = `NewResource-Default Diagram`
	__ResourceShape__00000000_.IsExpanded = false
	__ResourceShape__00000000_.X = 106.501544
	__ResourceShape__00000000_.Y = 104.249957
	__ResourceShape__00000000_.Width = 250.000000
	__ResourceShape__00000000_.Height = 100.000000

	__Root__00000000_.Name = `Root`
	__Root__00000000_.NbPixPerCharacter = 8.000000

	__Task__00000000_.Name = `NewTask`
	__Task__00000000_.Description = ``
	__Task__00000000_.IsExpanded = false
	__Task__00000000_.ComputedPrefix = `1`
	__Task__00000000_.IsInputsNodeExpanded = false
	__Task__00000000_.IsOutputsNodeExpanded = false
	__Task__00000000_.IsWithCompletion = false
	__Task__00000000_.Completion = ""

	__TaskShape__00000000_.Name = `NewTask-Default Diagram`
	__TaskShape__00000000_.IsExpanded = false
	__TaskShape__00000000_.X = 154.604908
	__TaskShape__00000000_.Y = 143.374505
	__TaskShape__00000000_.Width = 250.000000
	__TaskShape__00000000_.Height = 100.000000

	// insertion point for setup of pointers
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000000_)
	__Diagram__00000000_.Task_Shapes = append(__Diagram__00000000_.Task_Shapes, __TaskShape__00000000_)
	__Diagram__00000000_.Note_Shapes = append(__Diagram__00000000_.Note_Shapes, __NoteShape__00000000_)
	__Diagram__00000000_.Resource_Shapes = append(__Diagram__00000000_.Resource_Shapes, __ResourceShape__00000000_)
	__NoteShape__00000000_.Note = __Note__00000000_
	__ProductShape__00000000_.Product = __Product__00000000_
	__Project__00000000_.RootProducts = append(__Project__00000000_.RootProducts, __Product__00000000_)
	__Project__00000000_.RootTasks = append(__Project__00000000_.RootTasks, __Task__00000000_)
	__Project__00000000_.RootResources = append(__Project__00000000_.RootResources, __Resource__00000000_)
	__Project__00000000_.Notes = append(__Project__00000000_.Notes, __Note__00000000_)
	__Project__00000000_.Diagrams = append(__Project__00000000_.Diagrams, __Diagram__00000000_)
	__ResourceShape__00000000_.Resource = __Resource__00000000_
	__Root__00000000_.Projects = append(__Root__00000000_.Projects, __Project__00000000_)
	__TaskShape__00000000_.Task = __Task__00000000_
}
