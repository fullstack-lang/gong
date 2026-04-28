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

	__Diagram__00000000_ := (&models.Diagram{Name: `Default Diagram`}).Stage(stage)
	__Diagram__00000001_ := (&models.Diagram{Name: ``}).Stage(stage)
	__Library__00000000_ := (&models.Library{Name: ``}).Stage(stage)
	__ProductShape__00000000_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__Product__00000000_ := (&models.Product{Name: `a`}).Stage(stage)
	__Diagram__00000000_.Name = `Default Diagram`
	__Diagram__00000000_.ComputedPrefix = `1`
	__Diagram__00000000_.IsChecked = true
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.IsShowPrefix = false
	__Diagram__00000000_.DefaultBoxWidth = 250.000000
	__Diagram__00000000_.DefaultBoxHeigth = 70.000000
	__Diagram__00000000_.Width = 2300.000000
	__Diagram__00000000_.Height = 2300.000000
	__Diagram__00000000_.IsPBSNodeExpanded = true
	__Diagram__00000000_.IsWBSNodeExpanded = false
	__Diagram__00000000_.IsNotesNodeExpanded = false
	__Diagram__00000000_.IsResourcesNodeExpanded = false
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000000_)
	__Diagram__00000001_.Name = ``
	__Diagram__00000001_.ComputedPrefix = `2`
	__Diagram__00000001_.IsChecked = false
	__Diagram__00000001_.IsEditable_ = true
	__Diagram__00000001_.IsShowPrefix = false
	__Diagram__00000001_.DefaultBoxWidth = 250.000000
	__Diagram__00000001_.DefaultBoxHeigth = 70.000000
	__Diagram__00000001_.Width = 1100.000000
	__Diagram__00000001_.Height = 1100.000000
	__Diagram__00000001_.IsPBSNodeExpanded = false
	__Diagram__00000001_.IsWBSNodeExpanded = false
	__Diagram__00000001_.IsNotesNodeExpanded = false
	__Diagram__00000001_.IsResourcesNodeExpanded = false
	__Library__00000000_.Name = ``
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.NbPixPerCharacter = 8.000000
	__Library__00000000_.LogoSVGFile = ``
	__Library__00000000_.RootProducts = append(__Library__00000000_.RootProducts, __Product__00000000_)
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000000_)
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000001_)
	__ProductShape__00000000_.Name = `-Default Diagram`
	__ProductShape__00000000_.X = 153.106061
	__ProductShape__00000000_.Y = 153.126803
	__ProductShape__00000000_.Width = 250.000000
	__ProductShape__00000000_.Height = 70.000000
	__ProductShape__00000000_.IsHidden = false
	__ProductShape__00000000_.Product = __Product__00000000_
	__Product__00000000_.Name = `a`
	__Product__00000000_.ComputedPrefix = `1`
	__Product__00000000_.Description = ``
	__Product__00000000_.IsProducersNodeExpanded = false
	__Product__00000000_.IsConsumersNodeExpanded = false
	stage.Commit()

	// 
	__Diagram__00000001_.Width = 100.000000
	__Diagram__00000001_.Height = 100.000000
	// Default Diagram
	__Diagram__00000000_.Width = 503.106061
	__Diagram__00000000_.Height = 323.126803
	stage.Commit()

	// -Default Diagram
	__ProductShape__00000000_.X = 392.106061
	__ProductShape__00000000_.Y = 221.126803
	stage.Commit()

	// -Default Diagram
	__ProductShape__00000000_.X = 177.106061
	__ProductShape__00000000_.Y = 180.126803
	stage.Commit()

	// -Default Diagram
	__ProductShape__00000000_.X = 206.106061
	__ProductShape__00000000_.Y = 219.126788
	stage.Commit()

	// -Default Diagram
	__ProductShape__00000000_.X = 383.106061
	__ProductShape__00000000_.Y = 212.126803
	stage.Commit()

	// Default Diagram
	__Diagram__00000000_.Width = 733.106061
	__Diagram__00000000_.Height = 382.126803
	stage.Commit()

	// Default Diagram
	__Diagram__00000000_.Product_Shapes = slices.Delete( __Diagram__00000000_.Product_Shapes, 0, 1)
	__ProductShape__00000000_.Unstage(stage)
	stage.Commit()

	// Default Diagram
	__Diagram__00000000_.Width = 100.000000
	__Diagram__00000000_.Height = 100.000000
	stage.Commit()

	__ProductShape__00000001_ := (&models.ProductShape{Name: `a-Default Diagram`}).Stage(stage)
	// Default Diagram
	__Diagram__00000000_.Product_Shapes = slices.Insert( __Diagram__00000000_.Product_Shapes, 0, __ProductShape__00000001_)
	__ProductShape__00000001_.Name = `a-Default Diagram`
	__ProductShape__00000001_.X = 145.056982
	__ProductShape__00000001_.Y = 125.618199
	__ProductShape__00000001_.Width = 250.000000
	__ProductShape__00000001_.Height = 70.000000
	__ProductShape__00000001_.IsHidden = false
	__ProductShape__00000001_.Product = __Product__00000000_
	stage.Commit()

	// Default Diagram
	__Diagram__00000000_.Width = 495.056982
	__Diagram__00000000_.Height = 295.618199
	stage.Commit()

	// a-Default Diagram
	__ProductShape__00000001_.IsHidden = true
	stage.Commit()

	// a-Default Diagram
	__ProductShape__00000001_.IsHidden = false
	stage.Commit()

	__ProductCompositionShape__00000000_ := (&models.ProductCompositionShape{Name: `a to `}).Stage(stage)
	__ProductShape__00000002_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__Product__00000001_ := (&models.Product{Name: ``}).Stage(stage)
	// Default Diagram
	__Diagram__00000000_.Product_Shapes = slices.Insert( __Diagram__00000000_.Product_Shapes, 1, __ProductShape__00000002_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = slices.Insert( __Diagram__00000000_.ProductsWhoseNodeIsExpanded, 0, __Product__00000000_)
	__Diagram__00000000_.ProductComposition_Shapes = slices.Insert( __Diagram__00000000_.ProductComposition_Shapes, 0, __ProductCompositionShape__00000000_)
	// a
	__Product__00000000_.SubProducts = slices.Insert( __Product__00000000_.SubProducts, 0, __Product__00000001_)
	__ProductCompositionShape__00000000_.Name = `a to `
	__ProductCompositionShape__00000000_.StartRatio = 0.500000
	__ProductCompositionShape__00000000_.EndRatio = 0.500000
	__ProductCompositionShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000000_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000000_.IsHidden = false
	__ProductCompositionShape__00000000_.Product = __Product__00000001_
	__ProductShape__00000002_.Name = `-Default Diagram`
	__ProductShape__00000002_.X = 145.056982
	__ProductShape__00000002_.Y = 265.618199
	__ProductShape__00000002_.Width = 250.000000
	__ProductShape__00000002_.Height = 70.000000
	__ProductShape__00000002_.IsHidden = false
	__ProductShape__00000002_.Product = __Product__00000001_
	__Product__00000001_.Name = ``
	__Product__00000001_.ComputedPrefix = ``
	__Product__00000001_.Description = ``
	__Product__00000001_.IsProducersNodeExpanded = false
	__Product__00000001_.IsConsumersNodeExpanded = false
	stage.Commit()

	// 
	__Product__00000001_.ComputedPrefix = `1.1`
	// Default Diagram
	__Diagram__00000000_.Height = 435.618199
	stage.Commit()

	// A.1
	__Product__00000001_.Name = `A.1`
	stage.Commit()

	// Default Diagram
	__Diagram__00000000_.ProductComposition_Shapes = slices.Delete( __Diagram__00000000_.ProductComposition_Shapes, 0, 1)
	__ProductCompositionShape__00000000_.Unstage(stage)
	stage.Commit()

	__ProductCompositionShape__00000001_ := (&models.ProductCompositionShape{Name: `a to A.1`}).Stage(stage)
	// Default Diagram
	__Diagram__00000000_.ProductComposition_Shapes = slices.Insert( __Diagram__00000000_.ProductComposition_Shapes, 0, __ProductCompositionShape__00000001_)
	__ProductCompositionShape__00000001_.Name = `a to A.1`
	__ProductCompositionShape__00000001_.StartRatio = 0.500000
	__ProductCompositionShape__00000001_.EndRatio = 0.500000
	__ProductCompositionShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000001_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000001_.IsHidden = false
	__ProductCompositionShape__00000001_.Product = __Product__00000001_
	stage.Commit()
}