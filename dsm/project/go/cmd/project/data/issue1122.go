package main

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/project/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var (
	_ time.Time
	_ = slices.Index[[]int, int]
)

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__Diagram__00000000_ := (&models.Diagram{Name: `Default Diagram`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: `Root`}).Stage(stage)

	__Product__00000000_ := (&models.Product{Name: `P1`}).Stage(stage)
	__Product__00000001_ := (&models.Product{Name: `P1.1`}).Stage(stage)
	__Product__00000002_ := (&models.Product{Name: `P1.2`}).Stage(stage)

	__ProductCompositionShape__00000000_ := (&models.ProductCompositionShape{Name: `P1 to P1.1`}).Stage(stage)

	__ProductShape__00000000_ := (&models.ProductShape{Name: `P1-Default Diagram`}).Stage(stage)
	__ProductShape__00000001_ := (&models.ProductShape{Name: `P1.1-Default Diagram`}).Stage(stage)

	__Task__00000000_ := (&models.Task{Name: `L1.W1`}).Stage(stage)
	__Task__00000001_ := (&models.Task{Name: `L1.W2`}).Stage(stage)
	__Task__00000002_ := (&models.Task{Name: `L2.W1`}).Stage(stage)

	__TaskGroup__00000000_ := (&models.TaskGroup{Name: `L1`}).Stage(stage)
	__TaskGroup__00000002_ := (&models.TaskGroup{Name: `L2`}).Stage(stage)
	__TaskGroup__00000003_ := (&models.TaskGroup{Name: `L3`}).Stage(stage)

	__TaskGroupShape__00000000_ := (&models.TaskGroupShape{Name: `L1-Default Diagram`}).Stage(stage)
	__TaskGroupShape__00000001_ := (&models.TaskGroupShape{Name: `L3-Default Diagram`}).Stage(stage)
	__TaskGroupShape__00000002_ := (&models.TaskGroupShape{Name: `L2-Default Diagram`}).Stage(stage)

	__TaskShape__00000002_ := (&models.TaskShape{Name: `-Default Diagram`}).Stage(stage)
	__TaskShape__00000008_ := (&models.TaskShape{Name: `L1.W1-Default Diagram`}).Stage(stage)
	__TaskShape__00000009_ := (&models.TaskShape{Name: `L1.W2-Default Diagram`}).Stage(stage)

	// insertion point for initialization of values

	__Diagram__00000000_.Name = `Default Diagram`
	__Diagram__00000000_.ComputedPrefix = `1`
	__Diagram__00000000_.IsChecked = true
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.IsShowPrefix = false
	__Diagram__00000000_.DefaultBoxWidth = 250.000000
	__Diagram__00000000_.DefaultBoxHeigth = 70.000000
	__Diagram__00000000_.Width = 1100.000000
	__Diagram__00000000_.Height = 816.403016
	__Diagram__00000000_.IsPBSNodeExpanded = true
	__Diagram__00000000_.IsWBSNodeExpanded = true
	__Diagram__00000000_.IsTaskGroupsNodeExpanded = true
	__Diagram__00000000_.DateFormat = `2006`
	__Diagram__00000000_.IsNotesNodeExpanded = false
	__Diagram__00000000_.IsResourcesNodeExpanded = false
	__Diagram__00000000_.IsTimeDiagram = true
	__Diagram__00000000_.ComputedStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-01-01 00:00:00 +0000 UTC")
	__Diagram__00000000_.ComputedEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-06-01 00:00:00 +0000 UTC")
	__Diagram__00000000_.ComputedDuration = 13046400000000000
	__Diagram__00000000_.UseManualStartAndEndDates = false
	__Diagram__00000000_.ManualStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000000_.ManualEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000000_.TimeStep = 1
	__Diagram__00000000_.TimeStepScale = models.MONTHS
	__Diagram__00000000_.LaneHeight = 100.000000
	__Diagram__00000000_.RatioBarToLaneHeight = 0.800000
	__Diagram__00000000_.YTopMargin = 40.000000
	__Diagram__00000000_.XLeftText = 10.000000
	__Diagram__00000000_.TextHeight = 15.000000
	__Diagram__00000000_.XLeftLanes = 200.000000
	__Diagram__00000000_.XRightMargin = 1000.000000
	__Diagram__00000000_.ArrowLengthToTheRightOfStartBar = 15.000000
	__Diagram__00000000_.ArrowTipLenght = 5.000000
	__Diagram__00000000_.TimeLine_Color = `grey`
	__Diagram__00000000_.TimeLine_FillOpacity = 0.100000
	__Diagram__00000000_.TimeLine_Stroke = `grey`
	__Diagram__00000000_.TimeLine_StrokeWidth = 1.000000
	__Diagram__00000000_.Group_Stroke = `black`
	__Diagram__00000000_.Group_StrokeWidth = 1.000000
	__Diagram__00000000_.Group_StrokeDashArray = `2 2`
	__Diagram__00000000_.DateYOffset = 15.000000
	__Diagram__00000000_.AlignOnStartEndOnYearStart = false

	__Library__00000000_.Name = `Root`
	__Library__00000000_.NbPixPerCharacter = 8.000000
	__Library__00000000_.LogoSVGFile = ``
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsRootLibrary = true

	__Product__00000000_.Name = `P1`
	__Product__00000000_.ComputedPrefix = `1`
	__Product__00000000_.IsImport = false
	__Product__00000000_.Description = ``
	__Product__00000000_.IsProducersNodeExpanded = false
	__Product__00000000_.IsConsumersNodeExpanded = false

	__Product__00000001_.Name = `P1.1`
	__Product__00000001_.ComputedPrefix = `1.1`
	__Product__00000001_.IsImport = false
	__Product__00000001_.Description = ``
	__Product__00000001_.IsProducersNodeExpanded = false
	__Product__00000001_.IsConsumersNodeExpanded = false

	__Product__00000002_.Name = `P1.2`
	__Product__00000002_.ComputedPrefix = `1.2`
	__Product__00000002_.IsImport = false
	__Product__00000002_.Description = ``
	__Product__00000002_.IsProducersNodeExpanded = false
	__Product__00000002_.IsConsumersNodeExpanded = false

	__ProductCompositionShape__00000000_.Name = `P1 to P1.1`
	__ProductCompositionShape__00000000_.StartRatio = 0.500000
	__ProductCompositionShape__00000000_.EndRatio = 0.500000
	__ProductCompositionShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000000_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000000_.IsHidden = false

	__ProductShape__00000000_.Name = `P1-Default Diagram`
	__ProductShape__00000000_.X = 241.214319
	__ProductShape__00000000_.Y = 506.403016
	__ProductShape__00000000_.Width = 250.000000
	__ProductShape__00000000_.Height = 70.000000
	__ProductShape__00000000_.IsHidden = false

	__ProductShape__00000001_.Name = `P1.1-Default Diagram`
	__ProductShape__00000001_.X = 241.214319
	__ProductShape__00000001_.Y = 646.403016
	__ProductShape__00000001_.Width = 250.000000
	__ProductShape__00000001_.Height = 70.000000
	__ProductShape__00000001_.IsHidden = false

	__Task__00000000_.Name = `L1.W1`
	__Task__00000000_.ComputedPrefix = `1`
	__Task__00000000_.IsImport = false
	__Task__00000000_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-01-01 00:00:00 +0000 UTC")
	__Task__00000000_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-03-01 00:00:00 +0000 UTC")
	__Task__00000000_.Description = ``
	__Task__00000000_.IsInputsNodeExpanded = false
	__Task__00000000_.IsOutputsNodeExpanded = false
	__Task__00000000_.IsWithCompletion = false
	__Task__00000000_.Completion = ""

	__Task__00000001_.Name = `L1.W2`
	__Task__00000001_.ComputedPrefix = `2`
	__Task__00000001_.IsImport = false
	__Task__00000001_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-03-15 00:00:00 +0000 UTC")
	__Task__00000001_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-04-15 00:00:00 +0000 UTC")
	__Task__00000001_.Description = ``
	__Task__00000001_.IsInputsNodeExpanded = false
	__Task__00000001_.IsOutputsNodeExpanded = false
	__Task__00000001_.IsWithCompletion = false
	__Task__00000001_.Completion = ""

	__Task__00000002_.Name = `L2.W1`
	__Task__00000002_.ComputedPrefix = `3`
	__Task__00000002_.IsImport = false
	__Task__00000002_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-04-01 00:00:00 +0000 UTC")
	__Task__00000002_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-05-15 00:00:00 +0000 UTC")
	__Task__00000002_.Description = ``
	__Task__00000002_.IsInputsNodeExpanded = false
	__Task__00000002_.IsOutputsNodeExpanded = false
	__Task__00000002_.IsWithCompletion = false
	__Task__00000002_.Completion = ""

	__TaskGroup__00000000_.Name = `L1`
	__TaskGroup__00000000_.ComputedPrefix = ``

	__TaskGroup__00000002_.Name = `L2`
	__TaskGroup__00000002_.ComputedPrefix = ``

	__TaskGroup__00000003_.Name = `L3`
	__TaskGroup__00000003_.ComputedPrefix = ``

	__TaskGroupShape__00000000_.Name = `L1-Default Diagram`
	__TaskGroupShape__00000000_.X = 127.204461
	__TaskGroupShape__00000000_.Y = 102.225927
	__TaskGroupShape__00000000_.Width = 250.000000
	__TaskGroupShape__00000000_.Height = 70.000000
	__TaskGroupShape__00000000_.IsHidden = false

	__TaskGroupShape__00000001_.Name = `L3-Default Diagram`
	__TaskGroupShape__00000001_.X = 177.848696
	__TaskGroupShape__00000001_.Y = 179.428679
	__TaskGroupShape__00000001_.Width = 250.000000
	__TaskGroupShape__00000001_.Height = 70.000000
	__TaskGroupShape__00000001_.IsHidden = false

	__TaskGroupShape__00000002_.Name = `L2-Default Diagram`
	__TaskGroupShape__00000002_.X = 185.231479
	__TaskGroupShape__00000002_.Y = 175.285965
	__TaskGroupShape__00000002_.Width = 250.000000
	__TaskGroupShape__00000002_.Height = 70.000000
	__TaskGroupShape__00000002_.IsHidden = false

	__TaskShape__00000002_.Name = `-Default Diagram`
	__TaskShape__00000002_.IsShowDate = false
	__TaskShape__00000002_.X = 27.917788
	__TaskShape__00000002_.Y = 269.902722
	__TaskShape__00000002_.Width = 250.000000
	__TaskShape__00000002_.Height = 70.000000
	__TaskShape__00000002_.IsHidden = false

	__TaskShape__00000008_.Name = `L1.W1-Default Diagram`
	__TaskShape__00000008_.IsShowDate = false
	__TaskShape__00000008_.X = 159.552233
	__TaskShape__00000008_.Y = 135.522115
	__TaskShape__00000008_.Width = 250.000000
	__TaskShape__00000008_.Height = 70.000000
	__TaskShape__00000008_.IsHidden = false

	__TaskShape__00000009_.Name = `L1.W2-Default Diagram`
	__TaskShape__00000009_.IsShowDate = false
	__TaskShape__00000009_.X = 119.165805
	__TaskShape__00000009_.Y = 139.255674
	__TaskShape__00000009_.Width = 250.000000
	__TaskShape__00000009_.Height = 70.000000
	__TaskShape__00000009_.IsHidden = false

	// insertion point for setup of pointers
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000000_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000001_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, __Product__00000000_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000000_)
	__Diagram__00000000_.Task_Shapes = append(__Diagram__00000000_.Task_Shapes, __TaskShape__00000002_)
	__Diagram__00000000_.Task_Shapes = append(__Diagram__00000000_.Task_Shapes, __TaskShape__00000008_)
	__Diagram__00000000_.Task_Shapes = append(__Diagram__00000000_.Task_Shapes, __TaskShape__00000009_)
	__Diagram__00000000_.TaskGroupShapes = append(__Diagram__00000000_.TaskGroupShapes, __TaskGroupShape__00000000_)
	__Diagram__00000000_.TaskGroupShapes = append(__Diagram__00000000_.TaskGroupShapes, __TaskGroupShape__00000002_)
	__Diagram__00000000_.TaskGroupShapes = append(__Diagram__00000000_.TaskGroupShapes, __TaskGroupShape__00000001_)
	__Diagram__00000000_.TaskGroupsWhoseNodeIsExpanded = append(__Diagram__00000000_.TaskGroupsWhoseNodeIsExpanded, __TaskGroup__00000000_)
	__Library__00000000_.RootProducts = append(__Library__00000000_.RootProducts, __Product__00000000_)
	__Library__00000000_.RootTasks = append(__Library__00000000_.RootTasks, __Task__00000000_)
	__Library__00000000_.RootTasks = append(__Library__00000000_.RootTasks, __Task__00000001_)
	__Library__00000000_.RootTasks = append(__Library__00000000_.RootTasks, __Task__00000002_)
	__Library__00000000_.RootTaskGroups = append(__Library__00000000_.RootTaskGroups, __TaskGroup__00000000_)
	__Library__00000000_.RootTaskGroups = append(__Library__00000000_.RootTaskGroups, __TaskGroup__00000002_)
	__Library__00000000_.RootTaskGroups = append(__Library__00000000_.RootTaskGroups, __TaskGroup__00000003_)
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000000_)
	__Product__00000000_.ReferencedProduct = nil
	__Product__00000000_.SubProducts = append(__Product__00000000_.SubProducts, __Product__00000001_)
	__Product__00000000_.SubProducts = append(__Product__00000000_.SubProducts, __Product__00000002_)
	__Product__00000001_.ReferencedProduct = nil
	__Product__00000002_.ReferencedProduct = nil
	__ProductCompositionShape__00000000_.Product = __Product__00000001_
	__ProductShape__00000000_.Product = __Product__00000000_
	__ProductShape__00000001_.Product = __Product__00000001_
	__Task__00000000_.ReferencedTask = nil
	__Task__00000001_.ReferencedTask = nil
	__Task__00000002_.ReferencedTask = nil
	__TaskGroup__00000000_.Tasks = append(__TaskGroup__00000000_.Tasks, __Task__00000001_)
	__TaskGroup__00000000_.Tasks = append(__TaskGroup__00000000_.Tasks, __Task__00000000_)
	__TaskGroup__00000002_.Tasks = append(__TaskGroup__00000002_.Tasks, __Task__00000002_)
	__TaskGroupShape__00000000_.TaskGroup = __TaskGroup__00000000_
	__TaskGroupShape__00000001_.TaskGroup = __TaskGroup__00000003_
	__TaskGroupShape__00000002_.TaskGroup = __TaskGroup__00000002_
	__TaskShape__00000002_.Task = __Task__00000002_
	__TaskShape__00000008_.Task = __Task__00000000_
	__TaskShape__00000009_.Task = __Task__00000001_
}
