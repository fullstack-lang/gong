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

	// insertion point for initialization of values

	__Diagram__00000000_.Name = `Default Diagram`
	__Diagram__00000000_.ComputedPrefix = `1`
	__Diagram__00000000_.IsInRenameMode = false
	__Diagram__00000000_.IsExpanded = true
	__Diagram__00000000_.IsChecked = true
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.IsShowPrefix = false
	__Diagram__00000000_.DefaultBoxWidth = 250.000000
	__Diagram__00000000_.DefaultBoxHeigth = 70.000000
	__Diagram__00000000_.Width = 500.000000
	__Diagram__00000000_.Height = 500.000000
	__Diagram__00000000_.IsPBSNodeExpanded = false
	__Diagram__00000000_.IsWBSNodeExpanded = false
	__Diagram__00000000_.IsNotesNodeExpanded = false
	__Diagram__00000000_.IsResourcesNodeExpanded = false

	__Library__00000000_.Name = ``
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsInRenameMode = false
	__Library__00000000_.IsExpanded = true
	__Library__00000000_.NbPixPerCharacter = 8.000000

	// insertion point for setup of pointers
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000000_)

	stage.Commit()

	// Default Diagram
	__Diagram__00000000_.IsShowPrefix = true
	__Diagram__00000000_.Width = 600.000000
	__Diagram__00000000_.Height = 600.000000
	stage.Commit()

	// Default Diagram
	__Diagram__00000000_.IsShowPrefix = false
	stage.Commit()

	__ProductShape__00000000_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__Product__00000000_ := (&models.Product{Name: ``}).Stage(stage)
	// 
	__Library__00000000_.RootProducts = slices.Insert( __Library__00000000_.RootProducts, 0, __Product__00000000_)
	// Default Diagram
	__Diagram__00000000_.Product_Shapes = slices.Insert( __Diagram__00000000_.Product_Shapes, 0, __ProductShape__00000000_)
	__Diagram__00000000_.IsPBSNodeExpanded = true
	__ProductShape__00000000_.Name = `-Default Diagram`
	__ProductShape__00000000_.IsExpanded = false
	__ProductShape__00000000_.X = 153.106061
	__ProductShape__00000000_.Y = 153.126803
	__ProductShape__00000000_.Width = 250.000000
	__ProductShape__00000000_.Height = 70.000000
	__ProductShape__00000000_.IsHidden = false
	__ProductShape__00000000_.Product = __Product__00000000_
	__Product__00000000_.Name = ``
	__Product__00000000_.ComputedPrefix = ``
	__Product__00000000_.IsInRenameMode = true
	__Product__00000000_.IsExpanded = false
	__Product__00000000_.Description = ``
	__Product__00000000_.IsProducersNodeExpanded = false
	__Product__00000000_.IsConsumersNodeExpanded = false
	stage.Commit()

	// 
	__Product__00000000_.ComputedPrefix = `1`
	// Default Diagram
	__Diagram__00000000_.Width = 1000.000000
	__Diagram__00000000_.Height = 1000.000000
	stage.Commit()

	// a
	__Product__00000000_.Name = `a`
	__Product__00000000_.IsInRenameMode = false
	stage.Commit()

	// Default Diagram
	__Diagram__00000000_.IsShowPrefix = true
	stage.Commit()

	// Default Diagram
	__Diagram__00000000_.IsShowPrefix = false
	stage.Commit()
}
