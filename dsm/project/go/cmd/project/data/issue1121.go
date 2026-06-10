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

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__Diagram__00000000_ := (&models.Diagram{Name: `Default Diagram`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: ``}).Stage(stage)

	__Product__00000000_ := (&models.Product{Name: `P1`}).Stage(stage)
	__Product__00000001_ := (&models.Product{Name: `P1.1`}).Stage(stage)
	__Product__00000002_ := (&models.Product{Name: `P1.2`}).Stage(stage)

	__ProductCompositionShape__00000002_ := (&models.ProductCompositionShape{Name: `P1 to P1.1`}).Stage(stage)
	__ProductCompositionShape__00000003_ := (&models.ProductCompositionShape{Name: `P1 to P1.2`}).Stage(stage)

	__ProductShape__00000003_ := (&models.ProductShape{Name: `P1-Default Diagram`}).Stage(stage)
	__ProductShape__00000004_ := (&models.ProductShape{Name: `P1.1-Default Diagram`}).Stage(stage)
	__ProductShape__00000005_ := (&models.ProductShape{Name: `P1.2-Default Diagram`}).Stage(stage)

	__Task__00000000_ := (&models.Task{Name: `L1.W1`}).Stage(stage)
	__Task__00000001_ := (&models.Task{Name: `L1.W2`}).Stage(stage)
	__Task__00000002_ := (&models.Task{Name: `L2.W1`}).Stage(stage)

	__TaskGroup__00000000_ := (&models.TaskGroup{Name: `L1`}).Stage(stage)
	__TaskGroup__00000001_ := (&models.TaskGroup{Name: `L2`}).Stage(stage)
	__TaskGroup__00000002_ := (&models.TaskGroup{Name: `L2`}).Stage(stage)
	__TaskGroup__00000003_ := (&models.TaskGroup{Name: `L3`}).Stage(stage)

	__TaskShape__00000001_ := (&models.TaskShape{Name: `-Default Diagram`}).Stage(stage)
	__TaskShape__00000002_ := (&models.TaskShape{Name: `-Default Diagram`}).Stage(stage)
	__TaskShape__00000003_ := (&models.TaskShape{Name: `L1.W1-Default Diagram`}).Stage(stage)

	// insertion point for initialization of values

	__Diagram__00000000_.Name = `Default Diagram`
	__Diagram__00000000_.ComputedPrefix = `1`
	__Diagram__00000000_.IsChecked = true
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.IsShowPrefix = false
	__Diagram__00000000_.DefaultBoxWidth = 250.000000
	__Diagram__00000000_.DefaultBoxHeigth = 70.000000
	__Diagram__00000000_.Width = 1029.002714
	__Diagram__00000000_.Height = 721.589747
	__Diagram__00000000_.IsPBSNodeExpanded = true
	__Diagram__00000000_.IsWBSNodeExpanded = true
	__Diagram__00000000_.IsTaskGroupsNodeExpanded = true
	__Diagram__00000000_.DateFormat = `2006`
	__Diagram__00000000_.IsNotesNodeExpanded = false
	__Diagram__00000000_.IsResourcesNodeExpanded = false
	__Diagram__00000000_.IsTimeDiagram = true
	__Diagram__00000000_.ComputedStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2000-01-01 00:00:00 +0000 UTC")
	__Diagram__00000000_.ComputedEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2030-01-01 00:00:00 +0000 UTC")
	__Diagram__00000000_.ComputedDuration = 946771200000000000
	__Diagram__00000000_.UseManualStartAndEndDates = false
	__Diagram__00000000_.ManualStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000000_.ManualEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000000_.NumberOfYearsBetweenTicks = 5
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

	__Library__00000000_.Name = ``
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

	__ProductCompositionShape__00000002_.Name = `P1 to P1.1`
	__ProductCompositionShape__00000002_.StartRatio = 0.500000
	__ProductCompositionShape__00000002_.EndRatio = 0.500000
	__ProductCompositionShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000002_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000002_.IsHidden = false

	__ProductCompositionShape__00000003_.Name = `P1 to P1.2`
	__ProductCompositionShape__00000003_.StartRatio = 0.500000
	__ProductCompositionShape__00000003_.EndRatio = 0.500000
	__ProductCompositionShape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000003_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000003_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000003_.IsHidden = false

	__ProductShape__00000003_.Name = `P1-Default Diagram`
	__ProductShape__00000003_.X = 379.002714
	__ProductShape__00000003_.Y = 411.589747
	__ProductShape__00000003_.Width = 250.000000
	__ProductShape__00000003_.Height = 70.000000
	__ProductShape__00000003_.IsHidden = false

	__ProductShape__00000004_.Name = `P1.1-Default Diagram`
	__ProductShape__00000004_.X = 379.002714
	__ProductShape__00000004_.Y = 551.589747
	__ProductShape__00000004_.Width = 250.000000
	__ProductShape__00000004_.Height = 70.000000
	__ProductShape__00000004_.IsHidden = false

	__ProductShape__00000005_.Name = `P1.2-Default Diagram`
	__ProductShape__00000005_.X = 679.002714
	__ProductShape__00000005_.Y = 551.589747
	__ProductShape__00000005_.Width = 250.000000
	__ProductShape__00000005_.Height = 70.000000
	__ProductShape__00000005_.IsHidden = false

	__Task__00000000_.Name = `L1.W1`
	__Task__00000000_.ComputedPrefix = `1`
	__Task__00000000_.IsImport = false
	__Task__00000000_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2000-01-01 00:00:00 +0000 UTC")
	__Task__00000000_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2030-01-01 00:00:00 +0000 UTC")
	__Task__00000000_.Description = ``
	__Task__00000000_.IsInputsNodeExpanded = false
	__Task__00000000_.IsOutputsNodeExpanded = false
	__Task__00000000_.IsWithCompletion = false
	__Task__00000000_.Completion = ""

	__Task__00000001_.Name = `L1.W2`
	__Task__00000001_.ComputedPrefix = `2`
	__Task__00000001_.IsImport = false
	__Task__00000001_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2020-01-01 00:00:00 +0000 UTC")
	__Task__00000001_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2025-01-01 00:00:00 +0000 UTC")
	__Task__00000001_.Description = ``
	__Task__00000001_.IsInputsNodeExpanded = false
	__Task__00000001_.IsOutputsNodeExpanded = false
	__Task__00000001_.IsWithCompletion = false
	__Task__00000001_.Completion = ""

	__Task__00000002_.Name = `L2.W1`
	__Task__00000002_.ComputedPrefix = `3`
	__Task__00000002_.IsImport = false
	__Task__00000002_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2021-01-01 00:00:00 +0000 UTC")
	__Task__00000002_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-01-01 00:00:00 +0000 UTC")
	__Task__00000002_.Description = ``
	__Task__00000002_.IsInputsNodeExpanded = false
	__Task__00000002_.IsOutputsNodeExpanded = false
	__Task__00000002_.IsWithCompletion = false
	__Task__00000002_.Completion = ""

	__TaskGroup__00000000_.Name = `L1`

	__TaskGroup__00000001_.Name = `L2`

	__TaskGroup__00000002_.Name = `L2`

	__TaskGroup__00000003_.Name = `L3`

	__TaskShape__00000001_.Name = `-Default Diagram`
	__TaskShape__00000001_.IsShowDate = true
	__TaskShape__00000001_.X = 140.810427
	__TaskShape__00000001_.Y = 60.899550
	__TaskShape__00000001_.Width = 250.000000
	__TaskShape__00000001_.Height = 70.000000
	__TaskShape__00000001_.IsHidden = false

	__TaskShape__00000002_.Name = `-Default Diagram`
	__TaskShape__00000002_.IsShowDate = false
	__TaskShape__00000002_.X = 27.917788
	__TaskShape__00000002_.Y = 269.902722
	__TaskShape__00000002_.Width = 250.000000
	__TaskShape__00000002_.Height = 70.000000
	__TaskShape__00000002_.IsHidden = false

	__TaskShape__00000003_.Name = `L1.W1-Default Diagram`
	__TaskShape__00000003_.IsShowDate = true
	__TaskShape__00000003_.X = 106.789018
	__TaskShape__00000003_.Y = 184.181955
	__TaskShape__00000003_.Width = 250.000000
	__TaskShape__00000003_.Height = 70.000000
	__TaskShape__00000003_.IsHidden = false

	// insertion point for setup of pointers
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000003_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000004_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000005_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, __Product__00000000_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000002_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000003_)
	__Diagram__00000000_.Task_Shapes = append(__Diagram__00000000_.Task_Shapes, __TaskShape__00000001_)
	__Diagram__00000000_.Task_Shapes = append(__Diagram__00000000_.Task_Shapes, __TaskShape__00000002_)
	__Diagram__00000000_.Task_Shapes = append(__Diagram__00000000_.Task_Shapes, __TaskShape__00000003_)
	__Diagram__00000000_.TaskGroups = append(__Diagram__00000000_.TaskGroups, __TaskGroup__00000000_)
	__Diagram__00000000_.TaskGroups = append(__Diagram__00000000_.TaskGroups, __TaskGroup__00000002_)
	__Diagram__00000000_.TaskGroups = append(__Diagram__00000000_.TaskGroups, __TaskGroup__00000003_)
	__Diagram__00000000_.TaskGroupsWhoseNodeIsExpanded = append(__Diagram__00000000_.TaskGroupsWhoseNodeIsExpanded, __TaskGroup__00000002_)
	__Library__00000000_.RootProducts = append(__Library__00000000_.RootProducts, __Product__00000000_)
	__Library__00000000_.RootTasks = append(__Library__00000000_.RootTasks, __Task__00000000_)
	__Library__00000000_.RootTasks = append(__Library__00000000_.RootTasks, __Task__00000001_)
	__Library__00000000_.RootTasks = append(__Library__00000000_.RootTasks, __Task__00000002_)
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000000_)
	__Product__00000000_.ReferencedProduct = nil
	__Product__00000000_.SubProducts = append(__Product__00000000_.SubProducts, __Product__00000001_)
	__Product__00000000_.SubProducts = append(__Product__00000000_.SubProducts, __Product__00000002_)
	__Product__00000001_.ReferencedProduct = nil
	__Product__00000002_.ReferencedProduct = nil
	__ProductCompositionShape__00000002_.Product = __Product__00000001_
	__ProductCompositionShape__00000003_.Product = __Product__00000002_
	__ProductShape__00000003_.Product = __Product__00000000_
	__ProductShape__00000004_.Product = __Product__00000001_
	__ProductShape__00000005_.Product = __Product__00000002_
	__Task__00000000_.ReferencedTask = nil
	__Task__00000001_.ReferencedTask = nil
	__Task__00000002_.ReferencedTask = nil
	__TaskGroup__00000000_.Tasks = append(__TaskGroup__00000000_.Tasks, __Task__00000001_)
	__TaskGroup__00000000_.Tasks = append(__TaskGroup__00000000_.Tasks, __Task__00000000_)
	__TaskGroup__00000002_.Tasks = append(__TaskGroup__00000002_.Tasks, __Task__00000002_)
	__TaskShape__00000001_.Task = __Task__00000001_
	__TaskShape__00000002_.Task = __Task__00000002_
	__TaskShape__00000003_.Task = __Task__00000000_
}
