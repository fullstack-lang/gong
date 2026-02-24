package main

import (
	"slices"
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

	// Forward commits:

	__Diagram__00000000_ := (&models.Diagram{Name: `Default Diagram`}).Stage(stage)
	__Project__00000000_ := (&models.Project{Name: ``}).Stage(stage)
	__Root__00000000_ := (&models.Root{Name: `Root`}).Stage(stage)
	__Diagram__00000000_.Name = `Default Diagram`
	__Diagram__00000000_.IsChecked = true
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.ShowPrefix = false
	__Diagram__00000000_.DefaultBoxWidth = 250.000000
	__Diagram__00000000_.DefaultBoxHeigth = 70.000000
	__Diagram__00000000_.Width = 300.000000
	__Diagram__00000000_.Height = 300.000000
	__Diagram__00000000_.IsExpanded = true
	__Diagram__00000000_.ComputedPrefix = ``
	__Diagram__00000000_.IsInRenameMode = false
	__Diagram__00000000_.IsPBSNodeExpanded = false
	__Diagram__00000000_.IsWBSNodeExpanded = false
	__Diagram__00000000_.IsNotesNodeExpanded = false
	__Diagram__00000000_.IsResourcesNodeExpanded = false
	__Project__00000000_.Name = ``
	__Project__00000000_.IsExpanded = true
	__Project__00000000_.ComputedPrefix = ``
	__Project__00000000_.IsInRenameMode = false
	__Project__00000000_.Diagrams = append(__Project__00000000_.Diagrams, __Diagram__00000000_)
	__Root__00000000_.Name = `Root`
	__Root__00000000_.NbPixPerCharacter = 8.000000
	__Root__00000000_.Projects = append(__Root__00000000_.Projects, __Project__00000000_)
	stage.Commit()
	//
	__Project__00000000_.Diagrams = slices.Insert(__Project__00000000_.Diagrams, 0, __Diagram__00000000_)
	// Default Diagram
	__Diagram__00000000_.Width = 400.000000
	__Diagram__00000000_.Height = 400.000000
	__Diagram__00000000_.IsInRenameMode = true
	// Root
	__Root__00000000_.Projects = slices.Insert(__Root__00000000_.Projects, 0, __Project__00000000_)
	stage.Commit()
	//
	__Project__00000000_.Diagrams = slices.Insert(__Project__00000000_.Diagrams, 0, __Diagram__00000000_)
	// D1
	__Diagram__00000000_.Name = `D1`
	__Diagram__00000000_.Width = 500.000000
	__Diagram__00000000_.Height = 500.000000
	__Diagram__00000000_.IsInRenameMode = false
	// Root
	__Root__00000000_.Projects = slices.Insert(__Root__00000000_.Projects, 0, __Project__00000000_)
	stage.Commit()
	__ProductShape__00000000_ := (&models.ProductShape{Name: `-D1`}).Stage(stage)
	__Product__00000000_ := (&models.Product{Name: ``}).Stage(stage)
	//
	__Project__00000000_.RootProducts = slices.Insert(__Project__00000000_.RootProducts, 0, __Product__00000000_)
	__Project__00000000_.Diagrams = slices.Insert(__Project__00000000_.Diagrams, 0, __Diagram__00000000_)
	// D1
	__Diagram__00000000_.Width = 800.000000
	__Diagram__00000000_.Height = 800.000000
	__Diagram__00000000_.Product_Shapes = slices.Insert(__Diagram__00000000_.Product_Shapes, 0, __ProductShape__00000000_)
	__Diagram__00000000_.IsPBSNodeExpanded = true
	// Root
	__Root__00000000_.Projects = slices.Insert(__Root__00000000_.Projects, 0, __Project__00000000_)
	__ProductShape__00000000_.Name = `-D1`
	__ProductShape__00000000_.IsExpanded = false
	__ProductShape__00000000_.X = 198.086491
	__ProductShape__00000000_.Y = 188.841519
	__ProductShape__00000000_.Width = 250.000000
	__ProductShape__00000000_.Height = 70.000000
	__ProductShape__00000000_.Product = __Product__00000000_
	__Product__00000000_.Name = ``
	__Product__00000000_.Description = ``
	__Product__00000000_.IsExpanded = false
	__Product__00000000_.ComputedPrefix = `1`
	__Product__00000000_.IsInRenameMode = false
	__Product__00000000_.IsProducersNodeExpanded = false
	__Product__00000000_.IsConsumersNodeExpanded = false
	stage.Commit()
	//
	__Product__00000000_.IsInRenameMode = true
	//
	__Project__00000000_.RootProducts = slices.Insert(__Project__00000000_.RootProducts, 0, __Product__00000000_)
	__Project__00000000_.Diagrams = slices.Insert(__Project__00000000_.Diagrams, 0, __Diagram__00000000_)
	// -D1
	__ProductShape__00000000_.Product = __Product__00000000_
	// D1
	__Diagram__00000000_.Width = 900.000000
	__Diagram__00000000_.Height = 900.000000
	__Diagram__00000000_.Product_Shapes = slices.Insert(__Diagram__00000000_.Product_Shapes, 0, __ProductShape__00000000_)
	// Root
	__Root__00000000_.Projects = slices.Insert(__Root__00000000_.Projects, 0, __Project__00000000_)
	stage.Commit()
	//
	__Project__00000000_.RootProducts = slices.Insert(__Project__00000000_.RootProducts, 0, __Product__00000000_)
	__Project__00000000_.Diagrams = slices.Insert(__Project__00000000_.Diagrams, 0, __Diagram__00000000_)
	// -D1
	__ProductShape__00000000_.Product = __Product__00000000_
	// D1
	__Diagram__00000000_.Width = 1000.000000
	__Diagram__00000000_.Height = 1000.000000
	__Diagram__00000000_.Product_Shapes = slices.Insert(__Diagram__00000000_.Product_Shapes, 0, __ProductShape__00000000_)
	// P1
	__Product__00000000_.Name = `P1`
	__Product__00000000_.IsInRenameMode = false
	// Root
	__Root__00000000_.Projects = slices.Insert(__Root__00000000_.Projects, 0, __Project__00000000_)
	stage.Commit()
	__ProductCompositionShape__00000000_ := (&models.ProductCompositionShape{Name: `P1 to `}).Stage(stage)
	__ProductShape__00000001_ := (&models.ProductShape{Name: `-D1`}).Stage(stage)
	__Product__00000001_ := (&models.Product{Name: ``}).Stage(stage)
	//
	__Project__00000000_.RootProducts = slices.Insert(__Project__00000000_.RootProducts, 0, __Product__00000000_)
	__Project__00000000_.Diagrams = slices.Insert(__Project__00000000_.Diagrams, 0, __Diagram__00000000_)
	// -D1
	__ProductShape__00000000_.Product = __Product__00000000_
	// D1
	__Diagram__00000000_.Width = 1300.000000
	__Diagram__00000000_.Height = 1300.000000
	__Diagram__00000000_.Product_Shapes = slices.Insert(__Diagram__00000000_.Product_Shapes, 0, __ProductShape__00000000_)
	__Diagram__00000000_.Product_Shapes = slices.Insert(__Diagram__00000000_.Product_Shapes, 1, __ProductShape__00000001_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = slices.Insert(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, 0, __Product__00000000_)
	__Diagram__00000000_.ProductComposition_Shapes = slices.Insert(__Diagram__00000000_.ProductComposition_Shapes, 0, __ProductCompositionShape__00000000_)
	// P1
	__Product__00000000_.SubProducts = slices.Insert(__Product__00000000_.SubProducts, 0, __Product__00000001_)
	// Root
	__Root__00000000_.Projects = slices.Insert(__Root__00000000_.Projects, 0, __Project__00000000_)
	__ProductCompositionShape__00000000_.Name = `P1 to `
	__ProductCompositionShape__00000000_.StartRatio = 0.500000
	__ProductCompositionShape__00000000_.EndRatio = 0.500000
	__ProductCompositionShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000000_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000000_.Product = __Product__00000001_
	__ProductShape__00000001_.Name = `-D1`
	__ProductShape__00000001_.IsExpanded = false
	__ProductShape__00000001_.X = 198.086491
	__ProductShape__00000001_.Y = 328.841519
	__ProductShape__00000001_.Width = 250.000000
	__ProductShape__00000001_.Height = 70.000000
	__ProductShape__00000001_.Product = __Product__00000001_
	__Product__00000001_.Name = ``
	__Product__00000001_.Description = ``
	__Product__00000001_.IsExpanded = false
	__Product__00000001_.ComputedPrefix = `1.1`
	__Product__00000001_.IsInRenameMode = false
	__Product__00000001_.IsProducersNodeExpanded = false
	__Product__00000001_.IsConsumersNodeExpanded = false
	stage.Commit()
	//
	__Product__00000001_.IsInRenameMode = true
	//
	__Project__00000000_.RootProducts = slices.Insert(__Project__00000000_.RootProducts, 0, __Product__00000000_)
	__Project__00000000_.Diagrams = slices.Insert(__Project__00000000_.Diagrams, 0, __Diagram__00000000_)
	// -D1
	__ProductShape__00000000_.Product = __Product__00000000_
	// -D1
	__ProductShape__00000001_.Product = __Product__00000001_
	// D1
	__Diagram__00000000_.Width = 1400.000000
	__Diagram__00000000_.Height = 1400.000000
	__Diagram__00000000_.Product_Shapes = slices.Insert(__Diagram__00000000_.Product_Shapes, 0, __ProductShape__00000000_)
	__Diagram__00000000_.Product_Shapes = slices.Insert(__Diagram__00000000_.Product_Shapes, 1, __ProductShape__00000001_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = slices.Insert(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, 0, __Product__00000000_)
	__Diagram__00000000_.ProductComposition_Shapes = slices.Insert(__Diagram__00000000_.ProductComposition_Shapes, 0, __ProductCompositionShape__00000000_)
	// P1
	__Product__00000000_.SubProducts = slices.Insert(__Product__00000000_.SubProducts, 0, __Product__00000001_)
	// P1 to
	__ProductCompositionShape__00000000_.Product = __Product__00000001_
	// Root
	__Root__00000000_.Projects = slices.Insert(__Root__00000000_.Projects, 0, __Project__00000000_)
	stage.Commit()
	//
	__Project__00000000_.RootProducts = slices.Insert(__Project__00000000_.RootProducts, 0, __Product__00000000_)
	__Project__00000000_.Diagrams = slices.Insert(__Project__00000000_.Diagrams, 0, __Diagram__00000000_)
	// -D1
	__ProductShape__00000000_.Product = __Product__00000000_
	// -D1
	__ProductShape__00000001_.Product = __Product__00000001_
	// D1
	__Diagram__00000000_.Width = 1500.000000
	__Diagram__00000000_.Height = 1500.000000
	__Diagram__00000000_.Product_Shapes = slices.Insert(__Diagram__00000000_.Product_Shapes, 0, __ProductShape__00000000_)
	__Diagram__00000000_.Product_Shapes = slices.Insert(__Diagram__00000000_.Product_Shapes, 1, __ProductShape__00000001_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = slices.Insert(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, 0, __Product__00000000_)
	__Diagram__00000000_.ProductComposition_Shapes = slices.Insert(__Diagram__00000000_.ProductComposition_Shapes, 0, __ProductCompositionShape__00000000_)
	// P1
	__Product__00000000_.SubProducts = slices.Insert(__Product__00000000_.SubProducts, 0, __Product__00000001_)
	// P1 to
	__ProductCompositionShape__00000000_.Product = __Product__00000001_
	// P1.1
	__Product__00000001_.Name = `P1.1`
	__Product__00000001_.IsInRenameMode = false
	// Root
	__Root__00000000_.Projects = slices.Insert(__Root__00000000_.Projects, 0, __Project__00000000_)
	stage.Commit()
	//
	__Project__00000000_.RootProducts = slices.Insert(__Project__00000000_.RootProducts, 0, __Product__00000000_)
	__Project__00000000_.Diagrams = slices.Insert(__Project__00000000_.Diagrams, 0, __Diagram__00000000_)
	// -D1
	__ProductShape__00000000_.Product = __Product__00000000_
	// -D1
	__ProductShape__00000001_.Product = __Product__00000001_
	// D1
	__Diagram__00000000_.Product_Shapes = slices.Insert(__Diagram__00000000_.Product_Shapes, 0, __ProductShape__00000000_)
	__Diagram__00000000_.Product_Shapes = slices.Insert(__Diagram__00000000_.Product_Shapes, 1, __ProductShape__00000001_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = slices.Insert(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, 0, __Product__00000000_)
	__Diagram__00000000_.ProductComposition_Shapes = slices.Insert(__Diagram__00000000_.ProductComposition_Shapes, 0, __ProductCompositionShape__00000000_)
	// P1
	__Product__00000000_.SubProducts = slices.Insert(__Product__00000000_.SubProducts, 0, __Product__00000001_)
	// P1 to
	__ProductCompositionShape__00000000_.Product = __Product__00000001_
	// Root
	__Root__00000000_.Projects = slices.Insert(__Root__00000000_.Projects, 0, __Project__00000000_)
	stage.Commit()
	//
	__Project__00000000_.RootProducts = slices.Insert(__Project__00000000_.RootProducts, 0, __Product__00000000_)
	__Project__00000000_.Diagrams = slices.Insert(__Project__00000000_.Diagrams, 0, __Diagram__00000000_)
	// -D1
	__ProductShape__00000000_.Product = __Product__00000000_
	// D1
	__Diagram__00000000_.Width = 1800.000000
	__Diagram__00000000_.Height = 1800.000000
	__Diagram__00000000_.Product_Shapes = slices.Insert(__Diagram__00000000_.Product_Shapes, 0, __ProductShape__00000000_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = slices.Insert(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, 0, __Product__00000000_)
	// Root
	__Root__00000000_.Projects = slices.Insert(__Root__00000000_.Projects, 0, __Project__00000000_)
	__ProductCompositionShape__00000000_.Unstage(stage)
	__ProductShape__00000001_.Unstage(stage)
	__Product__00000001_.Unstage(stage)
	stage.Commit()
}
