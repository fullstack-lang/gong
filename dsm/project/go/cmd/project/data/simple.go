package main

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/project/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time
var _ = slices.Index[[]int, int]

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

	__Note__00000000_ := (&models.Note{Name: ``}).Stage(stage)

	__NoteShape__00000000_ := (&models.NoteShape{Name: `-Default Diagram`}).Stage(stage)

	__Product__00000000_ := (&models.Product{Name: `G3`}).Stage(stage)
	__Product__00000001_ := (&models.Product{Name: `G3.1`}).Stage(stage)
	__Product__00000002_ := (&models.Product{Name: ``}).Stage(stage)

	__ProductCompositionShape__00000000_ := (&models.ProductCompositionShape{Name: `G3 to `}).Stage(stage)
	__ProductCompositionShape__00000001_ := (&models.ProductCompositionShape{Name: `G3 to `}).Stage(stage)

	__ProductShape__00000000_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000001_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000002_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)

	// insertion point for initialization of values

	__Diagram__00000000_.Name = `Default Diagram`
	__Diagram__00000000_.ComputedPrefix = `1`
	__Diagram__00000000_.IsInRenameMode = false
	__Diagram__00000000_.IsExpanded = true
	__Diagram__00000000_.IsChecked = true
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.ShowPrefix = false
	__Diagram__00000000_.DefaultBoxWidth = 250.000000
	__Diagram__00000000_.DefaultBoxHeigth = 70.000000
	__Diagram__00000000_.Width = 6700.000000
	__Diagram__00000000_.Height = 6700.000000
	__Diagram__00000000_.IsPBSNodeExpanded = true
	__Diagram__00000000_.IsWBSNodeExpanded = false
	__Diagram__00000000_.IsNotesNodeExpanded = true
	__Diagram__00000000_.IsResourcesNodeExpanded = false

	__Library__00000000_.Name = ``
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsInRenameMode = false
	__Library__00000000_.IsExpanded = true
	__Library__00000000_.NbPixPerCharacter = 8.000000
	__Library__00000000_.LogoSVGFile = ``

	__Note__00000000_.Name = ``
	__Note__00000000_.ComputedPrefix = `1`
	__Note__00000000_.IsInRenameMode = false
	__Note__00000000_.IsExpanded = false

	__NoteShape__00000000_.Name = `-Default Diagram`
	__NoteShape__00000000_.IsExpanded = false
	__NoteShape__00000000_.X = 555.078303
	__NoteShape__00000000_.Y = 266.001756
	__NoteShape__00000000_.Width = 250.000000
	__NoteShape__00000000_.Height = 70.000000
	__NoteShape__00000000_.IsHidden = false

	__Product__00000000_.Name = `G3`
	__Product__00000000_.ComputedPrefix = `1`
	__Product__00000000_.IsInRenameMode = false
	__Product__00000000_.IsExpanded = false
	__Product__00000000_.Description = ``
	__Product__00000000_.IsProducersNodeExpanded = false
	__Product__00000000_.IsConsumersNodeExpanded = false

	__Product__00000001_.Name = `G3.1`
	__Product__00000001_.ComputedPrefix = `1.1`
	__Product__00000001_.IsInRenameMode = false
	__Product__00000001_.IsExpanded = false
	__Product__00000001_.Description = ``
	__Product__00000001_.IsProducersNodeExpanded = false
	__Product__00000001_.IsConsumersNodeExpanded = false

	__Product__00000002_.Name = ``
	__Product__00000002_.ComputedPrefix = `1.2`
	__Product__00000002_.IsInRenameMode = false
	__Product__00000002_.IsExpanded = false
	__Product__00000002_.Description = ``
	__Product__00000002_.IsProducersNodeExpanded = false
	__Product__00000002_.IsConsumersNodeExpanded = false

	__ProductCompositionShape__00000000_.Name = `G3 to `
	__ProductCompositionShape__00000000_.StartRatio = 0.500000
	__ProductCompositionShape__00000000_.EndRatio = 0.500000
	__ProductCompositionShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000000_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000000_.IsHidden = false

	__ProductCompositionShape__00000001_.Name = `G3 to `
	__ProductCompositionShape__00000001_.StartRatio = 0.500000
	__ProductCompositionShape__00000001_.EndRatio = 0.500000
	__ProductCompositionShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000001_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000001_.IsHidden = false

	__ProductShape__00000000_.Name = `-Default Diagram`
	__ProductShape__00000000_.IsExpanded = false
	__ProductShape__00000000_.X = 159.729721
	__ProductShape__00000000_.Y = 112.660089
	__ProductShape__00000000_.Width = 250.000000
	__ProductShape__00000000_.Height = 70.000000
	__ProductShape__00000000_.IsHidden = false

	__ProductShape__00000001_.Name = `-Default Diagram`
	__ProductShape__00000001_.IsExpanded = false
	__ProductShape__00000001_.X = 159.729721
	__ProductShape__00000001_.Y = 252.660089
	__ProductShape__00000001_.Width = 250.000000
	__ProductShape__00000001_.Height = 70.000000
	__ProductShape__00000001_.IsHidden = false

	__ProductShape__00000002_.Name = `-Default Diagram`
	__ProductShape__00000002_.IsExpanded = false
	__ProductShape__00000002_.X = 459.729721
	__ProductShape__00000002_.Y = 252.660089
	__ProductShape__00000002_.Width = 250.000000
	__ProductShape__00000002_.Height = 70.000000
	__ProductShape__00000002_.IsHidden = false

	// insertion point for setup of pointers
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000000_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000001_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000002_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, __Product__00000000_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000000_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000001_)
	__Diagram__00000000_.Note_Shapes = append(__Diagram__00000000_.Note_Shapes, __NoteShape__00000000_)
	__Library__00000000_.RootProducts = append(__Library__00000000_.RootProducts, __Product__00000000_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000000_)
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000000_)
	__NoteShape__00000000_.Note = __Note__00000000_
	__Product__00000000_.SubProducts = append(__Product__00000000_.SubProducts, __Product__00000001_)
	__Product__00000000_.SubProducts = append(__Product__00000000_.SubProducts, __Product__00000002_)
	__ProductCompositionShape__00000000_.Product = __Product__00000001_
	__ProductCompositionShape__00000001_.Product = __Product__00000002_
	__ProductShape__00000000_.Product = __Product__00000000_
	__ProductShape__00000001_.Product = __Product__00000001_
	__ProductShape__00000002_.Product = __Product__00000002_

	stage.Commit()

	// -Default Diagram
	__ProductShape__00000002_.Product = nil
	// G3
	__Product__00000000_.SubProducts = slices.Delete( __Product__00000000_.SubProducts, 1, 2)
	// G3 to 
	__ProductCompositionShape__00000001_.Product = nil
	__Product__00000002_.Unstage(stage)
	stage.Commit()

	// Default Diagram
	__Diagram__00000000_.Width = 7100.000000
	__Diagram__00000000_.Height = 7100.000000
	__Diagram__00000000_.Product_Shapes = slices.Delete( __Diagram__00000000_.Product_Shapes, 2, 3)
	__Diagram__00000000_.ProductComposition_Shapes = slices.Delete( __Diagram__00000000_.ProductComposition_Shapes, 1, 2)
	__ProductCompositionShape__00000001_.Unstage(stage)
	__ProductShape__00000002_.Unstage(stage)
	stage.Commit()

	// Default Diagram
	__Diagram__00000000_.IsExpanded = false
	stage.Commit()

	// Default Diagram
	__Diagram__00000000_.IsExpanded = true
	stage.Commit()

	// Default Diagram
	__Diagram__00000000_.Width = 7200.000000
	__Diagram__00000000_.Height = 7200.000000
	stage.Commit()

	__Diagram__00000001_ := (&models.Diagram{Name: ``}).Stage(stage)
	// 
	__Library__00000000_.Diagrams = slices.Insert( __Library__00000000_.Diagrams, 1, __Diagram__00000001_)
	// Default Diagram
	__Diagram__00000000_.IsChecked = false
	__Diagram__00000000_.Width = 7300.000000
	__Diagram__00000000_.Height = 7300.000000
	__Diagram__00000001_.Name = ``
	__Diagram__00000001_.ComputedPrefix = ``
	__Diagram__00000001_.IsInRenameMode = false
	__Diagram__00000001_.IsExpanded = true
	__Diagram__00000001_.IsChecked = true
	__Diagram__00000001_.IsEditable_ = true
	__Diagram__00000001_.ShowPrefix = false
	__Diagram__00000001_.DefaultBoxWidth = 0.000000
	__Diagram__00000001_.DefaultBoxHeigth = 0.000000
	__Diagram__00000001_.Width = 0.000000
	__Diagram__00000001_.Height = 0.000000
	__Diagram__00000001_.IsPBSNodeExpanded = false
	__Diagram__00000001_.IsWBSNodeExpanded = false
	__Diagram__00000001_.IsNotesNodeExpanded = false
	__Diagram__00000001_.IsResourcesNodeExpanded = false
	stage.Commit()

	// 
	__Diagram__00000001_.ComputedPrefix = `2`
	__Diagram__00000001_.DefaultBoxWidth = 250.000000
	__Diagram__00000001_.DefaultBoxHeigth = 70.000000
	__Diagram__00000001_.Width = 200.000000
	__Diagram__00000001_.Height = 200.000000
	// Default Diagram
	__Diagram__00000000_.Width = 7500.000000
	__Diagram__00000000_.Height = 7500.000000
	stage.Commit()

	// 
	__Diagram__00000001_.IsPBSNodeExpanded = true
	stage.Commit()

	__ProductShape__00000003_ := (&models.ProductShape{Name: `G3-`}).Stage(stage)
	// 
	__Diagram__00000001_.Product_Shapes = slices.Insert( __Diagram__00000001_.Product_Shapes, 0, __ProductShape__00000003_)
	__ProductShape__00000003_.Name = `G3-`
	__ProductShape__00000003_.IsExpanded = false
	__ProductShape__00000003_.X = 184.385391
	__ProductShape__00000003_.Y = 170.877611
	__ProductShape__00000003_.Width = 250.000000
	__ProductShape__00000003_.Height = 70.000000
	__ProductShape__00000003_.IsHidden = false
	__ProductShape__00000003_.Product = __Product__00000000_
	stage.Commit()

	// 
	__Diagram__00000001_.Width = 634.385391
	__Diagram__00000001_.Height = 500.000000
	// Default Diagram
	__Diagram__00000000_.Width = 7800.000000
	__Diagram__00000000_.Height = 7800.000000
	stage.Commit()

	// 
	__Diagram__00000001_.ProductsWhoseNodeIsExpanded = slices.Insert( __Diagram__00000001_.ProductsWhoseNodeIsExpanded, 0, __Product__00000000_)
	stage.Commit()

	__ProductShape__00000004_ := (&models.ProductShape{Name: `G3.1-`}).Stage(stage)
	// 
	__Diagram__00000001_.Width = 734.385391
	__Diagram__00000001_.Height = 600.000000
	__Diagram__00000001_.Product_Shapes = slices.Insert( __Diagram__00000001_.Product_Shapes, 1, __ProductShape__00000004_)
	// Default Diagram
	__Diagram__00000000_.Width = 7900.000000
	__Diagram__00000000_.Height = 7900.000000
	__ProductShape__00000004_.Name = `G3.1-`
	__ProductShape__00000004_.IsExpanded = false
	__ProductShape__00000004_.X = 163.488662
	__ProductShape__00000004_.Y = 177.098443
	__ProductShape__00000004_.Width = 250.000000
	__ProductShape__00000004_.Height = 70.000000
	__ProductShape__00000004_.IsHidden = false
	__ProductShape__00000004_.Product = __Product__00000001_
	stage.Commit()

	// 
	__Diagram__00000001_.Product_Shapes = slices.Delete( __Diagram__00000001_.Product_Shapes, 1, 2)
	__ProductShape__00000004_.Unstage(stage)
	stage.Commit()

	__ProductShape__00000005_ := (&models.ProductShape{Name: `G3.1-`}).Stage(stage)
	// 
	__Diagram__00000001_.Width = 834.385391
	__Diagram__00000001_.Height = 700.000000
	__Diagram__00000001_.Product_Shapes = slices.Insert( __Diagram__00000001_.Product_Shapes, 1, __ProductShape__00000005_)
	// Default Diagram
	__Diagram__00000000_.Width = 8000.000000
	__Diagram__00000000_.Height = 8000.000000
	__ProductShape__00000005_.Name = `G3.1-`
	__ProductShape__00000005_.IsExpanded = false
	__ProductShape__00000005_.X = 194.229556
	__ProductShape__00000005_.Y = 106.519097
	__ProductShape__00000005_.Width = 250.000000
	__ProductShape__00000005_.Height = 70.000000
	__ProductShape__00000005_.IsHidden = false
	__ProductShape__00000005_.Product = __Product__00000001_
	stage.Commit()

	// 
	__Diagram__00000001_.Product_Shapes = slices.Delete( __Diagram__00000001_.Product_Shapes, 1, 2)
	__ProductShape__00000005_.Unstage(stage)
	stage.Commit()

	__ProductShape__00000006_ := (&models.ProductShape{Name: `G3.1-`}).Stage(stage)
	// 
	__Diagram__00000001_.Product_Shapes = slices.Insert( __Diagram__00000001_.Product_Shapes, 1, __ProductShape__00000006_)
	__ProductShape__00000006_.Name = `G3.1-`
	__ProductShape__00000006_.IsExpanded = false
	__ProductShape__00000006_.X = 128.155354
	__ProductShape__00000006_.Y = 105.728001
	__ProductShape__00000006_.Width = 250.000000
	__ProductShape__00000006_.Height = 70.000000
	__ProductShape__00000006_.IsHidden = false
	__ProductShape__00000006_.Product = __Product__00000001_
	stage.Commit()

	// 
	__Diagram__00000001_.Product_Shapes = slices.Delete( __Diagram__00000001_.Product_Shapes, 1, 2)
	__ProductShape__00000006_.Unstage(stage)
	stage.Commit()

	__ProductCompositionShape__00000002_ := (&models.ProductCompositionShape{Name: `G3 to G3.1`}).Stage(stage)
	__ProductShape__00000007_ := (&models.ProductShape{Name: `G3.1-`}).Stage(stage)
	// 
	__Diagram__00000001_.Width = 934.385391
	__Diagram__00000001_.Height = 800.000000
	__Diagram__00000001_.Product_Shapes = slices.Insert( __Diagram__00000001_.Product_Shapes, 1, __ProductShape__00000007_)
	__Diagram__00000001_.ProductComposition_Shapes = slices.Insert( __Diagram__00000001_.ProductComposition_Shapes, 0, __ProductCompositionShape__00000002_)
	// Default Diagram
	__Diagram__00000000_.Width = 8100.000000
	__Diagram__00000000_.Height = 8100.000000
	__ProductCompositionShape__00000002_.Name = `G3 to G3.1`
	__ProductCompositionShape__00000002_.StartRatio = 0.500000
	__ProductCompositionShape__00000002_.EndRatio = 0.500000
	__ProductCompositionShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000002_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000002_.IsHidden = false
	__ProductCompositionShape__00000002_.Product = __Product__00000001_
	__ProductShape__00000007_.Name = `G3.1-`
	__ProductShape__00000007_.IsExpanded = false
	__ProductShape__00000007_.X = 184.385391
	__ProductShape__00000007_.Y = 310.877611
	__ProductShape__00000007_.Width = 250.000000
	__ProductShape__00000007_.Height = 70.000000
	__ProductShape__00000007_.IsHidden = false
	__ProductShape__00000007_.Product = __Product__00000001_
	stage.Commit()

	// 
	__Diagram__00000001_.IsChecked = false
	// Default Diagram
	__Diagram__00000000_.IsChecked = true
	stage.Commit()

	// -Default Diagram
	__NoteShape__00000000_.X = 409.078303
	__NoteShape__00000000_.Y = 447.001756
	stage.Commit()

	// 
	__Diagram__00000001_.IsChecked = true
	// Default Diagram
	__Diagram__00000000_.IsChecked = false
	stage.Commit()

	// 
	__Diagram__00000001_.IsChecked = false
	__Diagram__00000001_.Width = 1034.385391
	__Diagram__00000001_.Height = 900.000000
	// Default Diagram
	__Diagram__00000000_.IsChecked = true
	__Diagram__00000000_.Width = 8200.000000
	__Diagram__00000000_.Height = 8200.000000
	stage.Commit()

	__ProductCompositionShape__00000003_ := (&models.ProductCompositionShape{Name: `G3 to `}).Stage(stage)
	__ProductShape__00000008_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__Product__00000003_ := (&models.Product{Name: ``}).Stage(stage)
	// 
	__Diagram__00000001_.Width = 1134.385391
	__Diagram__00000001_.Height = 1000.000000
	// Default Diagram
	__Diagram__00000000_.Width = 8300.000000
	__Diagram__00000000_.Height = 8300.000000
	__Diagram__00000000_.Product_Shapes = slices.Insert( __Diagram__00000000_.Product_Shapes, 2, __ProductShape__00000008_)
	__Diagram__00000000_.ProductComposition_Shapes = slices.Insert( __Diagram__00000000_.ProductComposition_Shapes, 1, __ProductCompositionShape__00000003_)
	// G3
	__Product__00000000_.SubProducts = slices.Insert( __Product__00000000_.SubProducts, 1, __Product__00000003_)
	__ProductCompositionShape__00000003_.Name = `G3 to `
	__ProductCompositionShape__00000003_.StartRatio = 0.500000
	__ProductCompositionShape__00000003_.EndRatio = 0.500000
	__ProductCompositionShape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000003_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000003_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000003_.IsHidden = false
	__ProductCompositionShape__00000003_.Product = __Product__00000003_
	__ProductShape__00000008_.Name = `-Default Diagram`
	__ProductShape__00000008_.IsExpanded = false
	__ProductShape__00000008_.X = 459.729721
	__ProductShape__00000008_.Y = 252.660089
	__ProductShape__00000008_.Width = 250.000000
	__ProductShape__00000008_.Height = 70.000000
	__ProductShape__00000008_.IsHidden = false
	__ProductShape__00000008_.Product = __Product__00000003_
	__Product__00000003_.Name = ``
	__Product__00000003_.ComputedPrefix = ``
	__Product__00000003_.IsInRenameMode = false
	__Product__00000003_.IsExpanded = false
	__Product__00000003_.Description = ``
	__Product__00000003_.IsProducersNodeExpanded = false
	__Product__00000003_.IsConsumersNodeExpanded = false
	stage.Commit()

	// 
	__Diagram__00000001_.Width = 1334.385391
	__Diagram__00000001_.Height = 1200.000000
	// 
	__Product__00000003_.ComputedPrefix = `1.2`
	// Default Diagram
	__Diagram__00000000_.Width = 8500.000000
	__Diagram__00000000_.Height = 8500.000000
	stage.Commit()

	// 
	__Diagram__00000001_.IsChecked = true
	// Default Diagram
	__Diagram__00000000_.IsChecked = false
	stage.Commit()

	__ProductCompositionShape__00000004_ := (&models.ProductCompositionShape{Name: `G3 to `}).Stage(stage)
	__ProductShape__00000009_ := (&models.ProductShape{Name: `-`}).Stage(stage)
	// 
	__Diagram__00000001_.Product_Shapes = slices.Insert( __Diagram__00000001_.Product_Shapes, 2, __ProductShape__00000009_)
	__Diagram__00000001_.ProductComposition_Shapes = slices.Insert( __Diagram__00000001_.ProductComposition_Shapes, 1, __ProductCompositionShape__00000004_)
	__ProductCompositionShape__00000004_.Name = `G3 to `
	__ProductCompositionShape__00000004_.StartRatio = 0.500000
	__ProductCompositionShape__00000004_.EndRatio = 0.500000
	__ProductCompositionShape__00000004_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000004_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000004_.IsHidden = false
	__ProductCompositionShape__00000004_.Product = __Product__00000003_
	__ProductShape__00000009_.Name = `-`
	__ProductShape__00000009_.IsExpanded = false
	__ProductShape__00000009_.X = 484.385391
	__ProductShape__00000009_.Y = 310.877611
	__ProductShape__00000009_.Width = 250.000000
	__ProductShape__00000009_.Height = 70.000000
	__ProductShape__00000009_.IsHidden = false
	__ProductShape__00000009_.Product = __Product__00000003_
	stage.Commit()

	__ProductCompositionShape__00000005_ := (&models.ProductCompositionShape{Name: `G3 to G3.1`}).Stage(stage)
	__ProductCompositionShape__00000006_ := (&models.ProductCompositionShape{Name: `G3 to `}).Stage(stage)
	__ProductShape__00000010_ := (&models.ProductShape{Name: `G3-`}).Stage(stage)
	__ProductShape__00000011_ := (&models.ProductShape{Name: `G3.1-`}).Stage(stage)
	__ProductShape__00000012_ := (&models.ProductShape{Name: `-`}).Stage(stage)
	// 
	__Diagram__00000002_.Product_Shapes = slices.Delete( __Diagram__00000002_.Product_Shapes, 2, 3)
	__Diagram__00000002_.Product_Shapes = slices.Delete( __Diagram__00000002_.Product_Shapes, 1, 2)
	__Diagram__00000002_.Product_Shapes = slices.Delete( __Diagram__00000002_.Product_Shapes, 0, 1)
	__Diagram__00000002_.Product_Shapes = slices.Insert( __Diagram__00000002_.Product_Shapes, 0, __ProductShape__00000010_)
	__Diagram__00000002_.Product_Shapes = slices.Insert( __Diagram__00000002_.Product_Shapes, 1, __ProductShape__00000011_)
	__Diagram__00000002_.Product_Shapes = slices.Insert( __Diagram__00000002_.Product_Shapes, 2, __ProductShape__00000012_)
	__Diagram__00000002_.ProductComposition_Shapes = slices.Insert( __Diagram__00000002_.ProductComposition_Shapes, 2, __ProductCompositionShape__00000005_)
	__Diagram__00000002_.ProductComposition_Shapes = slices.Insert( __Diagram__00000002_.ProductComposition_Shapes, 3, __ProductCompositionShape__00000006_)
	__ProductCompositionShape__00000005_.Name = `G3 to G3.1`
	__ProductCompositionShape__00000005_.StartRatio = 0.500000
	__ProductCompositionShape__00000005_.EndRatio = 0.500000
	__ProductCompositionShape__00000005_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000005_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000005_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000005_.IsHidden = false
	__ProductCompositionShape__00000005_.Product = __Product__00000001_
	__ProductCompositionShape__00000006_.Name = `G3 to `
	__ProductCompositionShape__00000006_.StartRatio = 0.500000
	__ProductCompositionShape__00000006_.EndRatio = 0.500000
	__ProductCompositionShape__00000006_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000006_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000006_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000006_.IsHidden = false
	__ProductCompositionShape__00000006_.Product = __Product__00000003_
	__ProductShape__00000010_.Name = `G3-`
	__ProductShape__00000010_.IsExpanded = false
	__ProductShape__00000010_.X = 237.500000
	__ProductShape__00000010_.Y = 50.000000
	__ProductShape__00000010_.Width = 250.000000
	__ProductShape__00000010_.Height = 70.000000
	__ProductShape__00000010_.IsHidden = false
	__ProductShape__00000010_.Product = __Product__00000000_
	__ProductShape__00000011_.Name = `G3.1-`
	__ProductShape__00000011_.IsExpanded = false
	__ProductShape__00000011_.X = 50.000000
	__ProductShape__00000011_.Y = 190.000000
	__ProductShape__00000011_.Width = 250.000000
	__ProductShape__00000011_.Height = 70.000000
	__ProductShape__00000011_.IsHidden = false
	__ProductShape__00000011_.Product = __Product__00000001_
	__ProductShape__00000012_.Name = `-`
	__ProductShape__00000012_.IsExpanded = false
	__ProductShape__00000012_.X = 425.000000
	__ProductShape__00000012_.Y = 190.000000
	__ProductShape__00000012_.Width = 250.000000
	__ProductShape__00000012_.Height = 70.000000
	__ProductShape__00000012_.IsHidden = false
	__ProductShape__00000012_.Product = __Product__00000003_
	stage.Commit()

	// 
	__Diagram__00000002_.Width = 1834.385391
	__Diagram__00000002_.Height = 1700.000000
	__Diagram__00000002_.ProductComposition_Shapes = slices.Delete( __Diagram__00000002_.ProductComposition_Shapes, 3, 4)
	__Diagram__00000002_.ProductComposition_Shapes = slices.Delete( __Diagram__00000002_.ProductComposition_Shapes, 2, 3)
	// Default Diagram
	__Diagram__00000000_.Width = 9000.000000
	__Diagram__00000000_.Height = 9000.000000
	__ProductCompositionShape__00000005_.Unstage(stage)
	__ProductCompositionShape__00000006_.Unstage(stage)
	__ProductShape__00000003_.Unstage(stage)
	__ProductShape__00000007_.Unstage(stage)
	__ProductShape__00000009_.Unstage(stage)
	stage.Commit()

	// 
	__Diagram__00000002_.Product_Shapes = slices.Delete( __Diagram__00000002_.Product_Shapes, 2, 3)
	__ProductShape__00000012_.Unstage(stage)
	stage.Commit()

	// 
	__Diagram__00000002_.Width = 2034.385391
	__Diagram__00000002_.Height = 1900.000000
	__Diagram__00000002_.ProductComposition_Shapes = slices.Delete( __Diagram__00000002_.ProductComposition_Shapes, 1, 2)
	// Default Diagram
	__Diagram__00000000_.Width = 9200.000000
	__Diagram__00000000_.Height = 9200.000000
	__ProductCompositionShape__00000004_.Unstage(stage)
	stage.Commit()

	// 
	__Diagram__00000002_.Product_Shapes = slices.Delete( __Diagram__00000002_.Product_Shapes, 1, 2)
	__ProductShape__00000011_.Unstage(stage)
	stage.Commit()

	// 
	__Diagram__00000002_.Width = 2234.385391
	__Diagram__00000002_.Height = 2100.000000
	__Diagram__00000002_.ProductComposition_Shapes = slices.Delete( __Diagram__00000002_.ProductComposition_Shapes, 0, 1)
	// Default Diagram
	__Diagram__00000000_.Width = 9400.000000
	__Diagram__00000000_.Height = 9400.000000
	__ProductCompositionShape__00000002_.Unstage(stage)
	stage.Commit()

	// 
	__Diagram__00000002_.Product_Shapes = slices.Delete( __Diagram__00000002_.Product_Shapes, 0, 1)
	__ProductShape__00000010_.Unstage(stage)
	stage.Commit()

	__ProductCompositionShape__00000007_ := (&models.ProductCompositionShape{Name: `G3 to G3.1`}).Stage(stage)
	__ProductCompositionShape__00000008_ := (&models.ProductCompositionShape{Name: `G3 to `}).Stage(stage)
	__ProductShape__00000013_ := (&models.ProductShape{Name: `G3-`}).Stage(stage)
	__ProductShape__00000014_ := (&models.ProductShape{Name: `G3.1-`}).Stage(stage)
	__ProductShape__00000015_ := (&models.ProductShape{Name: `-`}).Stage(stage)
	// 
	__Diagram__00000003_.Product_Shapes = slices.Insert( __Diagram__00000003_.Product_Shapes, 0, __ProductShape__00000013_)
	__Diagram__00000003_.Product_Shapes = slices.Insert( __Diagram__00000003_.Product_Shapes, 1, __ProductShape__00000014_)
	__Diagram__00000003_.Product_Shapes = slices.Insert( __Diagram__00000003_.Product_Shapes, 2, __ProductShape__00000015_)
	__Diagram__00000003_.ProductComposition_Shapes = slices.Insert( __Diagram__00000003_.ProductComposition_Shapes, 0, __ProductCompositionShape__00000007_)
	__Diagram__00000003_.ProductComposition_Shapes = slices.Insert( __Diagram__00000003_.ProductComposition_Shapes, 1, __ProductCompositionShape__00000008_)
	__ProductCompositionShape__00000007_.Name = `G3 to G3.1`
	__ProductCompositionShape__00000007_.StartRatio = 0.500000
	__ProductCompositionShape__00000007_.EndRatio = 0.500000
	__ProductCompositionShape__00000007_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000007_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000007_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000007_.IsHidden = false
	__ProductCompositionShape__00000007_.Product = __Product__00000001_
	__ProductCompositionShape__00000008_.Name = `G3 to `
	__ProductCompositionShape__00000008_.StartRatio = 0.500000
	__ProductCompositionShape__00000008_.EndRatio = 0.500000
	__ProductCompositionShape__00000008_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000008_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000008_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000008_.IsHidden = false
	__ProductCompositionShape__00000008_.Product = __Product__00000003_
	__ProductShape__00000013_.Name = `G3-`
	__ProductShape__00000013_.IsExpanded = false
	__ProductShape__00000013_.X = 237.500000
	__ProductShape__00000013_.Y = 50.000000
	__ProductShape__00000013_.Width = 250.000000
	__ProductShape__00000013_.Height = 70.000000
	__ProductShape__00000013_.IsHidden = false
	__ProductShape__00000013_.Product = __Product__00000000_
	__ProductShape__00000014_.Name = `G3.1-`
	__ProductShape__00000014_.IsExpanded = false
	__ProductShape__00000014_.X = 50.000000
	__ProductShape__00000014_.Y = 190.000000
	__ProductShape__00000014_.Width = 250.000000
	__ProductShape__00000014_.Height = 70.000000
	__ProductShape__00000014_.IsHidden = false
	__ProductShape__00000014_.Product = __Product__00000001_
	__ProductShape__00000015_.Name = `-`
	__ProductShape__00000015_.IsExpanded = false
	__ProductShape__00000015_.X = 425.000000
	__ProductShape__00000015_.Y = 190.000000
	__ProductShape__00000015_.Width = 250.000000
	__ProductShape__00000015_.Height = 70.000000
	__ProductShape__00000015_.IsHidden = false
	__ProductShape__00000015_.Product = __Product__00000003_
	stage.Commit()
}
