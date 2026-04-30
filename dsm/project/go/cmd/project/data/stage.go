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

	__Product__00000000_ := (&models.Product{Name: `A`}).Stage(stage)
	__Product__00000001_ := (&models.Product{Name: `A.1`}).Stage(stage)
	__Product__00000002_ := (&models.Product{Name: `A.1.1`}).Stage(stage)
	__Product__00000003_ := (&models.Product{Name: `A.2`}).Stage(stage)
	__Product__00000004_ := (&models.Product{Name: `A.3`}).Stage(stage)
	__Product__00000005_ := (&models.Product{Name: ``}).Stage(stage)

	__ProductCompositionShape__00000001_ := (&models.ProductCompositionShape{Name: `a to A.1`}).Stage(stage)
	__ProductCompositionShape__00000002_ := (&models.ProductCompositionShape{Name: `A.1 to `}).Stage(stage)
	__ProductCompositionShape__00000003_ := (&models.ProductCompositionShape{Name: `A to `}).Stage(stage)
	__ProductCompositionShape__00000004_ := (&models.ProductCompositionShape{Name: `A to `}).Stage(stage)
	__ProductCompositionShape__00000005_ := (&models.ProductCompositionShape{Name: `A to `}).Stage(stage)

	__ProductShape__00000001_ := (&models.ProductShape{Name: `a-Default Diagram`}).Stage(stage)
	__ProductShape__00000002_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000003_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000004_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000005_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000006_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)

	// insertion point for initialization of values

	__Diagram__00000000_.Name = `Default Diagram`
	__Diagram__00000000_.ComputedPrefix = `1`
	__Diagram__00000000_.IsChecked = true
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.IsShowPrefix = false
	__Diagram__00000000_.DefaultBoxWidth = 250.000000
	__Diagram__00000000_.DefaultBoxHeigth = 70.000000
	__Diagram__00000000_.Width = 1395.056982
	__Diagram__00000000_.Height = 575.618199
	__Diagram__00000000_.IsPBSNodeExpanded = true
	__Diagram__00000000_.IsWBSNodeExpanded = false
	__Diagram__00000000_.IsNotesNodeExpanded = false
	__Diagram__00000000_.IsResourcesNodeExpanded = false

	__Library__00000000_.Name = ``
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.NbPixPerCharacter = 8.000000
	__Library__00000000_.LogoSVGFile = ``

	__Product__00000000_.Name = `A`
	__Product__00000000_.ComputedPrefix = `1`
	__Product__00000000_.Description = ``
	__Product__00000000_.IsProducersNodeExpanded = false
	__Product__00000000_.IsConsumersNodeExpanded = false

	__Product__00000001_.Name = `A.1`
	__Product__00000001_.ComputedPrefix = `1.1`
	__Product__00000001_.Description = ``
	__Product__00000001_.IsProducersNodeExpanded = false
	__Product__00000001_.IsConsumersNodeExpanded = false

	__Product__00000002_.Name = `A.1.1`
	__Product__00000002_.ComputedPrefix = `1.1.1`
	__Product__00000002_.Description = ``
	__Product__00000002_.IsProducersNodeExpanded = false
	__Product__00000002_.IsConsumersNodeExpanded = false

	__Product__00000003_.Name = `A.2`
	__Product__00000003_.ComputedPrefix = `1.2`
	__Product__00000003_.Description = ``
	__Product__00000003_.IsProducersNodeExpanded = false
	__Product__00000003_.IsConsumersNodeExpanded = false

	__Product__00000004_.Name = `A.3`
	__Product__00000004_.ComputedPrefix = `1.3`
	__Product__00000004_.Description = ``
	__Product__00000004_.IsProducersNodeExpanded = false
	__Product__00000004_.IsConsumersNodeExpanded = false

	__Product__00000005_.Name = ``
	__Product__00000005_.ComputedPrefix = `1.4`
	__Product__00000005_.Description = ``
	__Product__00000005_.IsProducersNodeExpanded = false
	__Product__00000005_.IsConsumersNodeExpanded = false

	__ProductCompositionShape__00000001_.Name = `a to A.1`
	__ProductCompositionShape__00000001_.StartRatio = 0.500000
	__ProductCompositionShape__00000001_.EndRatio = 0.500000
	__ProductCompositionShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000001_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000001_.IsHidden = false

	__ProductCompositionShape__00000002_.Name = `A.1 to `
	__ProductCompositionShape__00000002_.StartRatio = 0.500000
	__ProductCompositionShape__00000002_.EndRatio = 0.500000
	__ProductCompositionShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000002_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000002_.IsHidden = false

	__ProductCompositionShape__00000003_.Name = `A to `
	__ProductCompositionShape__00000003_.StartRatio = 0.500000
	__ProductCompositionShape__00000003_.EndRatio = 0.500000
	__ProductCompositionShape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000003_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000003_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000003_.IsHidden = false

	__ProductCompositionShape__00000004_.Name = `A to `
	__ProductCompositionShape__00000004_.StartRatio = 0.500000
	__ProductCompositionShape__00000004_.EndRatio = 0.500000
	__ProductCompositionShape__00000004_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000004_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000004_.IsHidden = false

	__ProductCompositionShape__00000005_.Name = `A to `
	__ProductCompositionShape__00000005_.StartRatio = 0.500000
	__ProductCompositionShape__00000005_.EndRatio = 0.500000
	__ProductCompositionShape__00000005_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000005_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000005_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000005_.IsHidden = false

	__ProductShape__00000001_.Name = `a-Default Diagram`
	__ProductShape__00000001_.X = 145.056982
	__ProductShape__00000001_.Y = 125.618199
	__ProductShape__00000001_.Width = 250.000000
	__ProductShape__00000001_.Height = 70.000000
	__ProductShape__00000001_.IsHidden = false

	__ProductShape__00000002_.Name = `-Default Diagram`
	__ProductShape__00000002_.X = 145.056982
	__ProductShape__00000002_.Y = 265.618199
	__ProductShape__00000002_.Width = 250.000000
	__ProductShape__00000002_.Height = 70.000000
	__ProductShape__00000002_.IsHidden = false

	__ProductShape__00000003_.Name = `-Default Diagram`
	__ProductShape__00000003_.X = 145.056982
	__ProductShape__00000003_.Y = 405.618199
	__ProductShape__00000003_.Width = 250.000000
	__ProductShape__00000003_.Height = 70.000000
	__ProductShape__00000003_.IsHidden = false

	__ProductShape__00000004_.Name = `-Default Diagram`
	__ProductShape__00000004_.X = 445.056982
	__ProductShape__00000004_.Y = 265.618199
	__ProductShape__00000004_.Width = 250.000000
	__ProductShape__00000004_.Height = 70.000000
	__ProductShape__00000004_.IsHidden = false

	__ProductShape__00000005_.Name = `-Default Diagram`
	__ProductShape__00000005_.X = 745.056982
	__ProductShape__00000005_.Y = 265.618199
	__ProductShape__00000005_.Width = 250.000000
	__ProductShape__00000005_.Height = 70.000000
	__ProductShape__00000005_.IsHidden = false

	__ProductShape__00000006_.Name = `-Default Diagram`
	__ProductShape__00000006_.X = 1045.056982
	__ProductShape__00000006_.Y = 265.618199
	__ProductShape__00000006_.Width = 250.000000
	__ProductShape__00000006_.Height = 70.000000
	__ProductShape__00000006_.IsHidden = false

	// insertion point for setup of pointers
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000001_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000002_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000003_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000004_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000005_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000006_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, __Product__00000001_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, __Product__00000000_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000001_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000002_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000003_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000004_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000005_)
	__Library__00000000_.RootProducts = append(__Library__00000000_.RootProducts, __Product__00000000_)
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000000_)
	__Product__00000000_.SubProducts = append(__Product__00000000_.SubProducts, __Product__00000001_)
	__Product__00000000_.SubProducts = append(__Product__00000000_.SubProducts, __Product__00000003_)
	__Product__00000000_.SubProducts = append(__Product__00000000_.SubProducts, __Product__00000004_)
	__Product__00000000_.SubProducts = append(__Product__00000000_.SubProducts, __Product__00000005_)
	__Product__00000001_.SubProducts = append(__Product__00000001_.SubProducts, __Product__00000002_)
	__ProductCompositionShape__00000001_.Product = __Product__00000001_
	__ProductCompositionShape__00000002_.Product = __Product__00000002_
	__ProductCompositionShape__00000003_.Product = __Product__00000003_
	__ProductCompositionShape__00000004_.Product = __Product__00000004_
	__ProductCompositionShape__00000005_.Product = __Product__00000005_
	__ProductShape__00000001_.Product = __Product__00000000_
	__ProductShape__00000002_.Product = __Product__00000001_
	__ProductShape__00000003_.Product = __Product__00000002_
	__ProductShape__00000004_.Product = __Product__00000003_
	__ProductShape__00000005_.Product = __Product__00000004_
	__ProductShape__00000006_.Product = __Product__00000005_

	stage.Commit()

	__Library__00000001_ := (&models.Library{Name: ``}).Stage(stage)
	// 
	__Library__00000000_.SubLibraries = slices.Insert( __Library__00000000_.SubLibraries, 0, __Library__00000001_)
	__Library__00000001_.Name = ``
	__Library__00000001_.ComputedPrefix = ``
	__Library__00000001_.NbPixPerCharacter = 0.000000
	__Library__00000001_.LogoSVGFile = ``
	stage.Commit()

	__Diagram__00000001_ := (&models.Diagram{Name: `Default Diagram`}).Stage(stage)
	// 
	__Library__00000001_.Diagrams = slices.Insert( __Library__00000001_.Diagrams, 0, __Diagram__00000001_)
	// Default Diagram
	__Diagram__00000000_.IsChecked = false
	__Diagram__00000001_.Name = `Default Diagram`
	__Diagram__00000001_.ComputedPrefix = `1`
	__Diagram__00000001_.IsChecked = true
	__Diagram__00000001_.IsEditable_ = true
	__Diagram__00000001_.IsShowPrefix = false
	__Diagram__00000001_.DefaultBoxWidth = 250.000000
	__Diagram__00000001_.DefaultBoxHeigth = 70.000000
	__Diagram__00000001_.Width = 100.000000
	__Diagram__00000001_.Height = 100.000000
	__Diagram__00000001_.IsPBSNodeExpanded = false
	__Diagram__00000001_.IsWBSNodeExpanded = false
	__Diagram__00000001_.IsNotesNodeExpanded = false
	__Diagram__00000001_.IsResourcesNodeExpanded = false
	stage.Commit()

	// L1
	__Library__00000001_.Name = `L1`
	stage.Commit()

	__ProductShape__00000007_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__Product__00000006_ := (&models.Product{Name: ``}).Stage(stage)
	// Default Diagram
	__Diagram__00000001_.Product_Shapes = slices.Insert( __Diagram__00000001_.Product_Shapes, 0, __ProductShape__00000007_)
	__Diagram__00000001_.IsPBSNodeExpanded = true
	// L1
	__Library__00000001_.RootProducts = slices.Insert( __Library__00000001_.RootProducts, 0, __Product__00000006_)
	__ProductShape__00000007_.Name = `-Default Diagram`
	__ProductShape__00000007_.X = 183.836804
	__ProductShape__00000007_.Y = 151.095882
	__ProductShape__00000007_.Width = 250.000000
	__ProductShape__00000007_.Height = 70.000000
	__ProductShape__00000007_.IsHidden = false
	__ProductShape__00000007_.Product = __Product__00000006_
	__Product__00000006_.Name = ``
	__Product__00000006_.ComputedPrefix = ``
	__Product__00000006_.Description = ``
	__Product__00000006_.IsProducersNodeExpanded = false
	__Product__00000006_.IsConsumersNodeExpanded = false
	stage.Commit()

	// 
	__Product__00000006_.ComputedPrefix = `1`
	// Default Diagram
	__Diagram__00000001_.Width = 533.836804
	__Diagram__00000001_.Height = 321.095882
	stage.Commit()

	// LA1
	__Product__00000006_.Name = `LA1`
	stage.Commit()

	// Default Diagram
	__Diagram__00000001_.Height = 321.095882
	stage.Commit()

	__Library__00000002_ := (&models.Library{Name: ``}).Stage(stage)
	// L1
	__Library__00000001_.SubLibraries = slices.Insert( __Library__00000001_.SubLibraries, 0, __Library__00000002_)
	__Library__00000002_.Name = ``
	__Library__00000002_.ComputedPrefix = ``
	__Library__00000002_.NbPixPerCharacter = 0.000000
	__Library__00000002_.LogoSVGFile = ``
	stage.Commit()

	__Diagram__00000002_ := (&models.Diagram{Name: `Default Diagram`}).Stage(stage)
	// 
	__Library__00000002_.Diagrams = slices.Insert( __Library__00000002_.Diagrams, 0, __Diagram__00000002_)
	// Default Diagram
	__Diagram__00000001_.IsChecked = false
	__Diagram__00000002_.Name = `Default Diagram`
	__Diagram__00000002_.ComputedPrefix = `1`
	__Diagram__00000002_.IsChecked = true
	__Diagram__00000002_.IsEditable_ = true
	__Diagram__00000002_.IsShowPrefix = false
	__Diagram__00000002_.DefaultBoxWidth = 250.000000
	__Diagram__00000002_.DefaultBoxHeigth = 70.000000
	__Diagram__00000002_.Width = 100.000000
	__Diagram__00000002_.Height = 100.000000
	__Diagram__00000002_.IsPBSNodeExpanded = false
	__Diagram__00000002_.IsWBSNodeExpanded = false
	__Diagram__00000002_.IsNotesNodeExpanded = false
	__Diagram__00000002_.IsResourcesNodeExpanded = false
	stage.Commit()

	// L1.1
	__Library__00000002_.Name = `L1.1`
	stage.Commit()

	// Default Diagram
	__Diagram__00000001_.IsChecked = true
	// Default Diagram
	__Diagram__00000002_.IsChecked = false
	stage.Commit()

	// Default Diagram
	__Diagram__00000000_.IsChecked = true
	// Default Diagram
	__Diagram__00000001_.IsChecked = false
	stage.Commit()
}
