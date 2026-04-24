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
	__Library__00000000_ := (&models.Library{Name: ``}).Stage(stage)
	__NoteShape__00000000_ := (&models.NoteShape{Name: `-Default Diagram`}).Stage(stage)
	__Note__00000000_ := (&models.Note{Name: ``}).Stage(stage)
	__ProductCompositionShape__00000000_ := (&models.ProductCompositionShape{Name: `G3 to `}).Stage(stage)
	__ProductShape__00000000_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000001_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__Product__00000000_ := (&models.Product{Name: `G3`}).Stage(stage)
	__Product__00000001_ := (&models.Product{Name: `G3.1`}).Stage(stage)
	__Diagram__00000000_.Name = `Default Diagram`
	__Diagram__00000000_.ComputedPrefix = `1`
	__Diagram__00000000_.IsInRenameMode = false
	__Diagram__00000000_.IsExpanded = true
	__Diagram__00000000_.IsChecked = true
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.IsShowPrefix = false
	__Diagram__00000000_.DefaultBoxWidth = 250.000000
	__Diagram__00000000_.DefaultBoxHeigth = 70.000000
	__Diagram__00000000_.Width = 11300.000000
	__Diagram__00000000_.Height = 11300.000000
	__Diagram__00000000_.IsPBSNodeExpanded = true
	__Diagram__00000000_.IsWBSNodeExpanded = false
	__Diagram__00000000_.IsNotesNodeExpanded = true
	__Diagram__00000000_.IsResourcesNodeExpanded = false
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000000_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000001_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, __Product__00000000_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000000_)
	__Diagram__00000000_.Note_Shapes = append(__Diagram__00000000_.Note_Shapes, __NoteShape__00000000_)
	__Library__00000000_.Name = ``
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsInRenameMode = false
	__Library__00000000_.IsExpanded = true
	__Library__00000000_.NbPixPerCharacter = 8.000000
	__Library__00000000_.LogoSVGFile = ``
	__Library__00000000_.RootProducts = append(__Library__00000000_.RootProducts, __Product__00000000_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000000_)
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000000_)
	__NoteShape__00000000_.Name = `-Default Diagram`
	__NoteShape__00000000_.IsExpanded = false
	__NoteShape__00000000_.X = 409.078303
	__NoteShape__00000000_.Y = 447.001756
	__NoteShape__00000000_.Width = 250.000000
	__NoteShape__00000000_.Height = 70.000000
	__NoteShape__00000000_.IsHidden = false
	__NoteShape__00000000_.Note = __Note__00000000_
	__Note__00000000_.Name = ``
	__Note__00000000_.ComputedPrefix = `1`
	__Note__00000000_.IsInRenameMode = false
	__Note__00000000_.IsExpanded = false
	__ProductCompositionShape__00000000_.Name = `G3 to `
	__ProductCompositionShape__00000000_.StartRatio = 0.500000
	__ProductCompositionShape__00000000_.EndRatio = 0.500000
	__ProductCompositionShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000000_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000000_.IsHidden = false
	__ProductCompositionShape__00000000_.Product = __Product__00000001_
	__ProductShape__00000000_.Name = `-Default Diagram`
	__ProductShape__00000000_.IsExpanded = false
	__ProductShape__00000000_.X = 159.729721
	__ProductShape__00000000_.Y = 112.660089
	__ProductShape__00000000_.Width = 250.000000
	__ProductShape__00000000_.Height = 70.000000
	__ProductShape__00000000_.IsHidden = false
	__ProductShape__00000000_.Product = __Product__00000000_
	__ProductShape__00000001_.Name = `-Default Diagram`
	__ProductShape__00000001_.IsExpanded = false
	__ProductShape__00000001_.X = 159.729721
	__ProductShape__00000001_.Y = 252.660089
	__ProductShape__00000001_.Width = 250.000000
	__ProductShape__00000001_.Height = 70.000000
	__ProductShape__00000001_.IsHidden = false
	__ProductShape__00000001_.Product = __Product__00000001_
	__Product__00000000_.Name = `G3`
	__Product__00000000_.ComputedPrefix = `1`
	__Product__00000000_.IsInRenameMode = false
	__Product__00000000_.IsExpanded = false
	__Product__00000000_.Description = ``
	__Product__00000000_.IsProducersNodeExpanded = false
	__Product__00000000_.IsConsumersNodeExpanded = false
	__Product__00000000_.SubProducts = append(__Product__00000000_.SubProducts, __Product__00000001_)
	__Product__00000001_.Name = `G3.1`
	__Product__00000001_.ComputedPrefix = `1.1`
	__Product__00000001_.IsInRenameMode = false
	__Product__00000001_.IsExpanded = false
	__Product__00000001_.Description = ``
	__Product__00000001_.IsProducersNodeExpanded = false
	__Product__00000001_.IsConsumersNodeExpanded = false
	stage.Commit()

	__Diagram__00000005_ := (&models.Diagram{Name: `Default Diagram copy`}).Stage(stage)
	__NoteShape__00000001_ := (&models.NoteShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductCompositionShape__00000011_ := (&models.ProductCompositionShape{Name: `G3 to `}).Stage(stage)
	__ProductShape__00000019_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000020_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	//
	__Library__00000000_.Diagrams = slices.Insert(__Library__00000000_.Diagrams, 1, __Diagram__00000005_)
	// Default Diagram
	__Diagram__00000000_.Width = 11400.000000
	__Diagram__00000000_.Height = 11400.000000
	__Diagram__00000005_.Name = `Default Diagram copy`
	__Diagram__00000005_.ComputedPrefix = ``
	__Diagram__00000005_.IsInRenameMode = false
	__Diagram__00000005_.IsExpanded = false
	__Diagram__00000005_.IsChecked = false
	__Diagram__00000005_.IsEditable_ = true
	__Diagram__00000005_.IsShowPrefix = false
	__Diagram__00000005_.DefaultBoxWidth = 0.000000
	__Diagram__00000005_.DefaultBoxHeigth = 0.000000
	__Diagram__00000005_.Width = 0.000000
	__Diagram__00000005_.Height = 0.000000
	__Diagram__00000005_.IsPBSNodeExpanded = false
	__Diagram__00000005_.IsWBSNodeExpanded = false
	__Diagram__00000005_.IsNotesNodeExpanded = false
	__Diagram__00000005_.IsResourcesNodeExpanded = false
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000019_)
	__Diagram__00000005_.Product_Shapes = append(__Diagram__00000005_.Product_Shapes, __ProductShape__00000020_)
	__Diagram__00000005_.ProductComposition_Shapes = append(__Diagram__00000005_.ProductComposition_Shapes, __ProductCompositionShape__00000011_)
	__Diagram__00000005_.Note_Shapes = append(__Diagram__00000005_.Note_Shapes, __NoteShape__00000001_)
	__NoteShape__00000001_.Name = `-Default Diagram`
	__NoteShape__00000001_.IsExpanded = false
	__NoteShape__00000001_.X = 409.078303
	__NoteShape__00000001_.Y = 447.001756
	__NoteShape__00000001_.Width = 250.000000
	__NoteShape__00000001_.Height = 70.000000
	__NoteShape__00000001_.IsHidden = false
	__NoteShape__00000001_.Note = nil
	__ProductCompositionShape__00000011_.Name = `G3 to `
	__ProductCompositionShape__00000011_.StartRatio = 0.500000
	__ProductCompositionShape__00000011_.EndRatio = 0.500000
	__ProductCompositionShape__00000011_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000011_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000011_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000011_.IsHidden = false
	__ProductCompositionShape__00000011_.Product = nil
	__ProductShape__00000019_.Name = `-Default Diagram`
	__ProductShape__00000019_.IsExpanded = false
	__ProductShape__00000019_.X = 159.729721
	__ProductShape__00000019_.Y = 112.660089
	__ProductShape__00000019_.Width = 250.000000
	__ProductShape__00000019_.Height = 70.000000
	__ProductShape__00000019_.IsHidden = false
	__ProductShape__00000019_.Product = nil
	__ProductShape__00000020_.Name = `-Default Diagram`
	__ProductShape__00000020_.IsExpanded = false
	__ProductShape__00000020_.X = 159.729721
	__ProductShape__00000020_.Y = 252.660089
	__ProductShape__00000020_.Width = 250.000000
	__ProductShape__00000020_.Height = 70.000000
	__ProductShape__00000020_.IsHidden = false
	__ProductShape__00000020_.Product = nil
	stage.Commit()

	// Default Diagram
	__Diagram__00000000_.Width = 11700.000000
	__Diagram__00000000_.Height = 11700.000000
	// Default Diagram copy
	__Diagram__00000005_.ComputedPrefix = `2`
	__Diagram__00000005_.DefaultBoxWidth = 250.000000
	__Diagram__00000005_.DefaultBoxHeigth = 70.000000
	__Diagram__00000005_.Width = 959.078303
	__Diagram__00000005_.Height = 817.001756
	__Diagram__00000005_.Product_Shapes = slices.Delete(__Diagram__00000005_.Product_Shapes, 1, 2)
	__Diagram__00000005_.Product_Shapes = slices.Delete(__Diagram__00000005_.Product_Shapes, 0, 1)
	__Diagram__00000005_.ProductComposition_Shapes = slices.Delete(__Diagram__00000005_.ProductComposition_Shapes, 0, 1)
	__Diagram__00000005_.Note_Shapes = slices.Delete(__Diagram__00000005_.Note_Shapes, 0, 1)
	__NoteShape__00000001_.Unstage(stage)
	__ProductCompositionShape__00000011_.Unstage(stage)
	__ProductShape__00000019_.Unstage(stage)
	__ProductShape__00000020_.Unstage(stage)
	stage.Commit()

	// Default Diagram copy
	__Diagram__00000005_.IsExpanded = true
	stage.Commit()

	// Default Diagram copy
	__Diagram__00000005_.IsPBSNodeExpanded = true
	stage.Commit()

	// Default Diagram copy
	__Diagram__00000005_.ProductsWhoseNodeIsExpanded = slices.Insert(__Diagram__00000005_.ProductsWhoseNodeIsExpanded, 0, __Product__00000000_)
	stage.Commit()

	//
	__Library__00000000_.Diagrams = slices.Delete(__Library__00000000_.Diagrams, 1, 2)
	// Default Diagram
	__Diagram__00000000_.Width = 11800.000000
	__Diagram__00000000_.Height = 11800.000000
	__Diagram__00000005_.Unstage(stage)
	stage.Commit()

	__Diagram__00000006_ := (&models.Diagram{Name: `Default Diagram copy`}).Stage(stage)
	__NoteShape__00000002_ := (&models.NoteShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductCompositionShape__00000012_ := (&models.ProductCompositionShape{Name: `G3 to `}).Stage(stage)
	__ProductShape__00000021_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000022_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	//
	__Library__00000000_.Diagrams = slices.Insert(__Library__00000000_.Diagrams, 1, __Diagram__00000006_)
	__Diagram__00000006_.Name = `Default Diagram copy`
	__Diagram__00000006_.ComputedPrefix = ``
	__Diagram__00000006_.IsInRenameMode = false
	__Diagram__00000006_.IsExpanded = false
	__Diagram__00000006_.IsChecked = false
	__Diagram__00000006_.IsEditable_ = true
	__Diagram__00000006_.IsShowPrefix = false
	__Diagram__00000006_.DefaultBoxWidth = 0.000000
	__Diagram__00000006_.DefaultBoxHeigth = 0.000000
	__Diagram__00000006_.Width = 0.000000
	__Diagram__00000006_.Height = 0.000000
	__Diagram__00000006_.IsPBSNodeExpanded = false
	__Diagram__00000006_.IsWBSNodeExpanded = false
	__Diagram__00000006_.IsNotesNodeExpanded = false
	__Diagram__00000006_.IsResourcesNodeExpanded = false
	__Diagram__00000006_.Product_Shapes = append(__Diagram__00000006_.Product_Shapes, __ProductShape__00000021_)
	__Diagram__00000006_.Product_Shapes = append(__Diagram__00000006_.Product_Shapes, __ProductShape__00000022_)
	__Diagram__00000006_.ProductComposition_Shapes = append(__Diagram__00000006_.ProductComposition_Shapes, __ProductCompositionShape__00000012_)
	__Diagram__00000006_.Note_Shapes = append(__Diagram__00000006_.Note_Shapes, __NoteShape__00000002_)
	__NoteShape__00000002_.Name = `-Default Diagram`
	__NoteShape__00000002_.IsExpanded = false
	__NoteShape__00000002_.X = 409.078303
	__NoteShape__00000002_.Y = 447.001756
	__NoteShape__00000002_.Width = 250.000000
	__NoteShape__00000002_.Height = 70.000000
	__NoteShape__00000002_.IsHidden = false
	__NoteShape__00000002_.Note = __Note__00000000_
	__ProductCompositionShape__00000012_.Name = `G3 to `
	__ProductCompositionShape__00000012_.StartRatio = 0.500000
	__ProductCompositionShape__00000012_.EndRatio = 0.500000
	__ProductCompositionShape__00000012_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000012_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000012_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000012_.IsHidden = false
	__ProductCompositionShape__00000012_.Product = __Product__00000001_
	__ProductShape__00000021_.Name = `-Default Diagram`
	__ProductShape__00000021_.IsExpanded = false
	__ProductShape__00000021_.X = 159.729721
	__ProductShape__00000021_.Y = 112.660089
	__ProductShape__00000021_.Width = 250.000000
	__ProductShape__00000021_.Height = 70.000000
	__ProductShape__00000021_.IsHidden = false
	__ProductShape__00000021_.Product = __Product__00000000_
	__ProductShape__00000022_.Name = `-Default Diagram`
	__ProductShape__00000022_.IsExpanded = false
	__ProductShape__00000022_.X = 159.729721
	__ProductShape__00000022_.Y = 252.660089
	__ProductShape__00000022_.Width = 250.000000
	__ProductShape__00000022_.Height = 70.000000
	__ProductShape__00000022_.IsHidden = false
	__ProductShape__00000022_.Product = __Product__00000001_
	stage.Commit()

	// Default Diagram
	__Diagram__00000000_.Width = 12100.000000
	__Diagram__00000000_.Height = 12100.000000
	// Default Diagram copy
	__Diagram__00000006_.ComputedPrefix = `2`
	__Diagram__00000006_.DefaultBoxWidth = 250.000000
	__Diagram__00000006_.DefaultBoxHeigth = 70.000000
	__Diagram__00000006_.Width = 859.078303
	__Diagram__00000006_.Height = 717.001756
	stage.Commit()

	// Default Diagram
	__Diagram__00000000_.IsChecked = false
	// Default Diagram copy
	__Diagram__00000006_.IsChecked = true
	stage.Commit()

	// Default Diagram
	__Diagram__00000000_.IsChecked = true
	// Default Diagram copy
	__Diagram__00000006_.IsChecked = false
	stage.Commit()

	//
	__Library__00000000_.Diagrams = slices.Delete(__Library__00000000_.Diagrams, 1, 2)
	// Default Diagram
	__Diagram__00000000_.Width = 12200.000000
	__Diagram__00000000_.Height = 12200.000000
	__Diagram__00000006_.Unstage(stage)
	stage.Commit()

	// Default Diagram
	__Diagram__00000000_.Width = 12400.000000
	__Diagram__00000000_.Height = 12400.000000
	__NoteShape__00000002_.Unstage(stage)
	__ProductCompositionShape__00000012_.Unstage(stage)
	__ProductShape__00000021_.Unstage(stage)
	__ProductShape__00000022_.Unstage(stage)
	stage.Commit()

	__Diagram__00000007_ := (&models.Diagram{Name: `Default Diagram copy`}).Stage(stage)
	__NoteShape__00000003_ := (&models.NoteShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductCompositionShape__00000013_ := (&models.ProductCompositionShape{Name: `G3 to `}).Stage(stage)
	__ProductShape__00000023_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000024_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	//
	__Library__00000000_.Diagrams = slices.Insert(__Library__00000000_.Diagrams, 1, __Diagram__00000007_)
	__Diagram__00000007_.Name = `Default Diagram copy`
	__Diagram__00000007_.ComputedPrefix = ``
	__Diagram__00000007_.IsInRenameMode = false
	__Diagram__00000007_.IsExpanded = false
	__Diagram__00000007_.IsChecked = false
	__Diagram__00000007_.IsEditable_ = true
	__Diagram__00000007_.IsShowPrefix = false
	__Diagram__00000007_.DefaultBoxWidth = 0.000000
	__Diagram__00000007_.DefaultBoxHeigth = 0.000000
	__Diagram__00000007_.Width = 0.000000
	__Diagram__00000007_.Height = 0.000000
	__Diagram__00000007_.IsPBSNodeExpanded = false
	__Diagram__00000007_.IsWBSNodeExpanded = false
	__Diagram__00000007_.IsNotesNodeExpanded = false
	__Diagram__00000007_.IsResourcesNodeExpanded = false
	__Diagram__00000007_.Product_Shapes = append(__Diagram__00000007_.Product_Shapes, __ProductShape__00000023_)
	__Diagram__00000007_.Product_Shapes = append(__Diagram__00000007_.Product_Shapes, __ProductShape__00000024_)
	__Diagram__00000007_.ProductComposition_Shapes = append(__Diagram__00000007_.ProductComposition_Shapes, __ProductCompositionShape__00000013_)
	__Diagram__00000007_.Note_Shapes = append(__Diagram__00000007_.Note_Shapes, __NoteShape__00000003_)
	__NoteShape__00000003_.Name = `-Default Diagram`
	__NoteShape__00000003_.IsExpanded = false
	__NoteShape__00000003_.X = 409.078303
	__NoteShape__00000003_.Y = 447.001756
	__NoteShape__00000003_.Width = 250.000000
	__NoteShape__00000003_.Height = 70.000000
	__NoteShape__00000003_.IsHidden = false
	__NoteShape__00000003_.Note = __Note__00000000_
	__ProductCompositionShape__00000013_.Name = `G3 to `
	__ProductCompositionShape__00000013_.StartRatio = 0.500000
	__ProductCompositionShape__00000013_.EndRatio = 0.500000
	__ProductCompositionShape__00000013_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000013_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000013_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000013_.IsHidden = false
	__ProductCompositionShape__00000013_.Product = __Product__00000001_
	__ProductShape__00000023_.Name = `-Default Diagram`
	__ProductShape__00000023_.IsExpanded = false
	__ProductShape__00000023_.X = 159.729721
	__ProductShape__00000023_.Y = 112.660089
	__ProductShape__00000023_.Width = 250.000000
	__ProductShape__00000023_.Height = 70.000000
	__ProductShape__00000023_.IsHidden = false
	__ProductShape__00000023_.Product = __Product__00000000_
	__ProductShape__00000024_.Name = `-Default Diagram`
	__ProductShape__00000024_.IsExpanded = false
	__ProductShape__00000024_.X = 159.729721
	__ProductShape__00000024_.Y = 252.660089
	__ProductShape__00000024_.Width = 250.000000
	__ProductShape__00000024_.Height = 70.000000
	__ProductShape__00000024_.IsHidden = false
	__ProductShape__00000024_.Product = __Product__00000001_
	stage.Commit()

	// Default Diagram
	__Diagram__00000000_.Width = 12600.000000
	__Diagram__00000000_.Height = 12600.000000
	// Default Diagram copy
	__Diagram__00000007_.ComputedPrefix = `2`
	__Diagram__00000007_.DefaultBoxWidth = 250.000000
	__Diagram__00000007_.DefaultBoxHeigth = 70.000000
	__Diagram__00000007_.Width = 859.078303
	__Diagram__00000007_.Height = 717.001756
	stage.Commit()

	// Default Diagram
	__Diagram__00000000_.IsChecked = false
	// Default Diagram copy
	__Diagram__00000007_.IsChecked = true
	stage.Commit()

	// Default Diagram
	__Diagram__00000000_.IsChecked = true
	// Default Diagram copy
	__Diagram__00000007_.IsChecked = false
	stage.Commit()

	// Default Diagram
	__Diagram__00000000_.Width = 12700.000000
	__Diagram__00000000_.Height = 12700.000000
	// Default Diagram copy
	__Diagram__00000007_.Width = 959.078303
	__Diagram__00000007_.Height = 817.001756
	stage.Commit()

	// Default Diagram
	__Diagram__00000000_.IsChecked = false
	// Default Diagram copy
	__Diagram__00000007_.IsChecked = true
	stage.Commit()

	// Default Diagram
	__Diagram__00000000_.IsChecked = true
	// Default Diagram copy
	__Diagram__00000007_.IsChecked = false
	stage.Commit()

	// Default Diagram
	__Diagram__00000000_.IsChecked = false
	// Default Diagram copy
	__Diagram__00000007_.IsChecked = true
	stage.Commit()

	// Default Diagram copy
	__Diagram__00000007_.IsExpanded = true
	stage.Commit()

	// Default Diagram copy
	__Diagram__00000007_.IsPBSNodeExpanded = true
	stage.Commit()

	// Default Diagram copy
	__Diagram__00000007_.ProductsWhoseNodeIsExpanded = slices.Insert( __Diagram__00000007_.ProductsWhoseNodeIsExpanded, 0, __Product__00000000_)
	stage.Commit()

	// Default Diagram copy
	__Diagram__00000007_.Product_Shapes = slices.Delete( __Diagram__00000007_.Product_Shapes, 1, 2)
	__ProductShape__00000024_.Unstage(stage)
	stage.Commit()

	// Default Diagram
	__Diagram__00000000_.Width = 13700.000000
	__Diagram__00000000_.Height = 13700.000000
	// Default Diagram copy
	__Diagram__00000007_.Width = 1959.078303
	__Diagram__00000007_.Height = 1817.001756
	__Diagram__00000007_.ProductComposition_Shapes = slices.Delete( __Diagram__00000007_.ProductComposition_Shapes, 0, 1)
	__ProductCompositionShape__00000013_.Unstage(stage)
	stage.Commit()

	// Default Diagram
	__Diagram__00000000_.IsChecked = true
	// Default Diagram copy
	__Diagram__00000007_.IsChecked = false
	stage.Commit()

	// Default Diagram
	__Diagram__00000000_.IsChecked = false
	// Default Diagram copy
	__Diagram__00000007_.IsChecked = true
	stage.Commit()
}
