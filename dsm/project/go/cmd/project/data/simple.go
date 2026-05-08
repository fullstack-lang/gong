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

	__Diagram__00000000_ := (&models.Diagram{Name: `D1`}).Stage(stage)
	__Diagram__00000007_ := (&models.Diagram{Name: `Default Diagram copy`}).Stage(stage)
	__Library__00000000_ := (&models.Library{Name: ``}).Stage(stage)
	__NoteShape__00000000_ := (&models.NoteShape{Name: `-Default Diagram`}).Stage(stage)
	__NoteShape__00000003_ := (&models.NoteShape{Name: `-Default Diagram`}).Stage(stage)
	__Note__00000000_ := (&models.Note{Name: ``}).Stage(stage)
	__ProductCompositionShape__00000000_ := (&models.ProductCompositionShape{Name: `G3 to `}).Stage(stage)
	__ProductShape__00000000_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000001_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000023_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__Product__00000000_ := (&models.Product{Name: `G3`}).Stage(stage)
	__Product__00000001_ := (&models.Product{Name: `G3.1`}).Stage(stage)
	__Diagram__00000000_.Name = `D1`
	__Diagram__00000000_.ComputedPrefix = `1`
	__Diagram__00000000_.IsChecked = false
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.IsShowPrefix = false
	__Diagram__00000000_.DefaultBoxWidth = 250.000000
	__Diagram__00000000_.DefaultBoxHeigth = 70.000000
	__Diagram__00000000_.Width = 15300.000000
	__Diagram__00000000_.Height = 15300.000000
	__Diagram__00000000_.IsPBSNodeExpanded = true
	__Diagram__00000000_.IsWBSNodeExpanded = false
	__Diagram__00000000_.IsNotesNodeExpanded = true
	__Diagram__00000000_.IsResourcesNodeExpanded = false
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000000_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000001_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, __Product__00000000_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000000_)
	__Diagram__00000000_.Note_Shapes = append(__Diagram__00000000_.Note_Shapes, __NoteShape__00000000_)
	__Diagram__00000007_.Name = `Default Diagram copy`
	__Diagram__00000007_.ComputedPrefix = `2`
	__Diagram__00000007_.IsChecked = true
	__Diagram__00000007_.IsEditable_ = true
	__Diagram__00000007_.IsShowPrefix = false
	__Diagram__00000007_.DefaultBoxWidth = 250.000000
	__Diagram__00000007_.DefaultBoxHeigth = 70.000000
	__Diagram__00000007_.Width = 3559.078303
	__Diagram__00000007_.Height = 3417.001756
	__Diagram__00000007_.IsPBSNodeExpanded = true
	__Diagram__00000007_.IsWBSNodeExpanded = false
	__Diagram__00000007_.IsNotesNodeExpanded = false
	__Diagram__00000007_.IsResourcesNodeExpanded = false
	__Diagram__00000007_.Product_Shapes = append(__Diagram__00000007_.Product_Shapes, __ProductShape__00000023_)
	__Diagram__00000007_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000007_.ProductsWhoseNodeIsExpanded, __Product__00000000_)
	__Diagram__00000007_.Note_Shapes = append(__Diagram__00000007_.Note_Shapes, __NoteShape__00000003_)
	__Library__00000000_.Name = ``
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.NbPixPerCharacter = 8.000000
	__Library__00000000_.LogoSVGFile = ``
	__Library__00000000_.RootProducts = append(__Library__00000000_.RootProducts, __Product__00000000_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000000_)
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000000_)
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000007_)
	__NoteShape__00000000_.Name = `-Default Diagram`
	__NoteShape__00000000_.X = 409.078303
	__NoteShape__00000000_.Y = 447.001756
	__NoteShape__00000000_.Width = 250.000000
	__NoteShape__00000000_.Height = 70.000000
	__NoteShape__00000000_.IsHidden = false
	__NoteShape__00000000_.Note = __Note__00000000_
	__NoteShape__00000003_.Name = `-Default Diagram`
	__NoteShape__00000003_.X = 409.078303
	__NoteShape__00000003_.Y = 447.001756
	__NoteShape__00000003_.Width = 250.000000
	__NoteShape__00000003_.Height = 70.000000
	__NoteShape__00000003_.IsHidden = false
	__NoteShape__00000003_.Note = __Note__00000000_
	__Note__00000000_.Name = ``
	__Note__00000000_.ComputedPrefix = `1`
	__ProductCompositionShape__00000000_.Name = `G3 to `
	__ProductCompositionShape__00000000_.StartRatio = 0.500000
	__ProductCompositionShape__00000000_.EndRatio = 0.500000
	__ProductCompositionShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000000_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000000_.IsHidden = false
	__ProductCompositionShape__00000000_.Product = __Product__00000001_
	__ProductShape__00000000_.Name = `-Default Diagram`
	__ProductShape__00000000_.X = 159.729721
	__ProductShape__00000000_.Y = 112.660089
	__ProductShape__00000000_.Width = 250.000000
	__ProductShape__00000000_.Height = 70.000000
	__ProductShape__00000000_.IsHidden = false
	__ProductShape__00000000_.Product = __Product__00000000_
	__ProductShape__00000001_.Name = `-Default Diagram`
	__ProductShape__00000001_.X = 159.729721
	__ProductShape__00000001_.Y = 252.660089
	__ProductShape__00000001_.Width = 250.000000
	__ProductShape__00000001_.Height = 70.000000
	__ProductShape__00000001_.IsHidden = false
	__ProductShape__00000001_.Product = __Product__00000001_
	__ProductShape__00000023_.Name = `-Default Diagram`
	__ProductShape__00000023_.X = 159.729721
	__ProductShape__00000023_.Y = 112.660089
	__ProductShape__00000023_.Width = 250.000000
	__ProductShape__00000023_.Height = 70.000000
	__ProductShape__00000023_.IsHidden = false
	__ProductShape__00000023_.Product = __Product__00000000_
	__Product__00000000_.Name = `G3`
	__Product__00000000_.ComputedPrefix = `1`
	__Product__00000000_.Description = ``
	__Product__00000000_.IsProducersNodeExpanded = false
	__Product__00000000_.IsConsumersNodeExpanded = false
	__Product__00000000_.SubProducts = append(__Product__00000000_.SubProducts, __Product__00000001_)
	__Product__00000001_.Name = `G3.1`
	__Product__00000001_.ComputedPrefix = `1.1`
	__Product__00000001_.Description = ``
	__Product__00000001_.IsProducersNodeExpanded = false
	__Product__00000001_.IsConsumersNodeExpanded = false
	stage.Commit()
}
