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
	__Diagram__00000001_ := (&models.Diagram{Name: `Default Diagram`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: `Root Lib`}).Stage(stage)
	__Library__00000001_ := (&models.Library{Name: `Sub Library`}).Stage(stage)

	__Note__00000000_ := (&models.Note{Name: `Example of note`}).Stage(stage)

	__NoteShape__00000000_ := (&models.NoteShape{Name: `-Default Diagram`}).Stage(stage)

	__Product__00000000_ := (&models.Product{Name: `P1`}).Stage(stage)
	__Product__00000001_ := (&models.Product{Name: `P1* (sub lib)`}).Stage(stage)
	__Product__00000002_ := (&models.Product{Name: `Imported P1 (from sub lib)`}).Stage(stage)

	__ProductShape__00000000_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000001_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000002_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)

	// insertion point for initialization of values

	__Diagram__00000000_.Name = `Default Diagram`
	__Diagram__00000000_.ComputedPrefix = `1`
	__Diagram__00000000_.IsChecked = true
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.IsShowPrefix = false
	__Diagram__00000000_.DefaultBoxWidth = 250.000000
	__Diagram__00000000_.DefaultBoxHeigth = 70.000000
	__Diagram__00000000_.Width = 865.429257
	__Diagram__00000000_.Height = 474.638598
	__Diagram__00000000_.IsPBSNodeExpanded = true
	__Diagram__00000000_.IsWBSNodeExpanded = true
	__Diagram__00000000_.IsNotesNodeExpanded = true
	__Diagram__00000000_.IsResourcesNodeExpanded = false

	__Diagram__00000001_.Name = `Default Diagram`
	__Diagram__00000001_.ComputedPrefix = `1`
	__Diagram__00000001_.IsChecked = false
	__Diagram__00000001_.IsEditable_ = true
	__Diagram__00000001_.IsShowPrefix = false
	__Diagram__00000001_.DefaultBoxWidth = 250.000000
	__Diagram__00000001_.DefaultBoxHeigth = 70.000000
	__Diagram__00000001_.Width = 539.502921
	__Diagram__00000001_.Height = 315.645142
	__Diagram__00000001_.IsPBSNodeExpanded = true
	__Diagram__00000001_.IsWBSNodeExpanded = true
	__Diagram__00000001_.IsNotesNodeExpanded = false
	__Diagram__00000001_.IsResourcesNodeExpanded = false

	__Library__00000000_.Name = `Root Lib`
	__Library__00000000_.NbPixPerCharacter = 8.000000
	__Library__00000000_.LogoSVGFile = ``
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsRootLibrary = true

	__Library__00000001_.Name = `Sub Library`
	__Library__00000001_.NbPixPerCharacter = 0.000000
	__Library__00000001_.LogoSVGFile = ``
	__Library__00000001_.ComputedPrefix = ``
	__Library__00000001_.IsRootLibrary = false

	__Note__00000000_.Name = `Example of note`
	__Note__00000000_.ComputedPrefix = `1`

	__NoteShape__00000000_.Name = `-Default Diagram`
	__NoteShape__00000000_.X = 192.541401
	__NoteShape__00000000_.Y = 304.638598
	__NoteShape__00000000_.Width = 250.000000
	__NoteShape__00000000_.Height = 70.000000
	__NoteShape__00000000_.IsHidden = false

	__Product__00000000_.Name = `P1`
	__Product__00000000_.ComputedPrefix = `1`
	__Product__00000000_.IsImport = false
	__Product__00000000_.Description = ``
	__Product__00000000_.IsProducersNodeExpanded = false
	__Product__00000000_.IsConsumersNodeExpanded = false

	__Product__00000001_.Name = `P1* (sub lib)`
	__Product__00000001_.ComputedPrefix = `1`
	__Product__00000001_.IsImport = false
	__Product__00000001_.Description = ``
	__Product__00000001_.IsProducersNodeExpanded = false
	__Product__00000001_.IsConsumersNodeExpanded = false

	__Product__00000002_.Name = `Imported P1 (from sub lib)`
	__Product__00000002_.ComputedPrefix = `2`
	__Product__00000002_.IsImport = true
	__Product__00000002_.Description = ``
	__Product__00000002_.IsProducersNodeExpanded = false
	__Product__00000002_.IsConsumersNodeExpanded = false

	__ProductShape__00000000_.Name = `-Default Diagram`
	__ProductShape__00000000_.X = 137.007620
	__ProductShape__00000000_.Y = 157.296725
	__ProductShape__00000000_.Width = 250.000000
	__ProductShape__00000000_.Height = 70.000000
	__ProductShape__00000000_.IsHidden = false

	__ProductShape__00000001_.Name = `-Default Diagram`
	__ProductShape__00000001_.X = 189.502921
	__ProductShape__00000001_.Y = 145.645142
	__ProductShape__00000001_.Width = 250.000000
	__ProductShape__00000001_.Height = 70.000000
	__ProductShape__00000001_.IsHidden = false

	__ProductShape__00000002_.Name = `-Default Diagram`
	__ProductShape__00000002_.X = 515.429257
	__ProductShape__00000002_.Y = 210.687668
	__ProductShape__00000002_.Width = 250.000000
	__ProductShape__00000002_.Height = 70.000000
	__ProductShape__00000002_.IsHidden = false

	// insertion point for setup of pointers
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000000_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000002_)
	__Diagram__00000000_.Note_Shapes = append(__Diagram__00000000_.Note_Shapes, __NoteShape__00000000_)
	__Diagram__00000001_.Product_Shapes = append(__Diagram__00000001_.Product_Shapes, __ProductShape__00000001_)
	__Library__00000000_.SubLibraries = append(__Library__00000000_.SubLibraries, __Library__00000001_)
	__Library__00000000_.RootProducts = append(__Library__00000000_.RootProducts, __Product__00000000_)
	__Library__00000000_.RootProducts = append(__Library__00000000_.RootProducts, __Product__00000002_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000000_)
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000000_)
	__Library__00000001_.RootProducts = append(__Library__00000001_.RootProducts, __Product__00000001_)
	__Library__00000001_.Diagrams = append(__Library__00000001_.Diagrams, __Diagram__00000001_)
	__NoteShape__00000000_.Note = __Note__00000000_
	__Product__00000000_.ReferencedProduct = nil
	__Product__00000001_.ReferencedProduct = nil
	__Product__00000002_.ReferencedProduct = __Product__00000001_
	__ProductShape__00000000_.Product = __Product__00000000_
	__ProductShape__00000001_.Product = __Product__00000001_
	__ProductShape__00000002_.Product = __Product__00000002_
}
