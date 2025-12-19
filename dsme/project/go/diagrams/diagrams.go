package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/dsme/project/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration
var _ ref_models.Stage

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// Declaration of instances to stage

	__Classdiagram__00000000_Default := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__00000000_Diagram_Package_created_the_2025_12_18T17_41_37Z := (&models.DiagramPackage{}).Stage(stage)

	__GongStructShape__00000000_Default_Product := (&models.GongStructShape{}).Stage(stage)

	__LinkShape__00000000_SubProducts := (&models.LinkShape{}).Stage(stage)

	// Setup of values

	__Classdiagram__00000000_Default.Name = `Default`
	__Classdiagram__00000000_Default.Description = ``
	__Classdiagram__00000000_Default.IsIncludedInStaticWebSite = false
	__Classdiagram__00000000_Default.ShowNbInstances = true
	__Classdiagram__00000000_Default.ShowMultiplicity = true
	__Classdiagram__00000000_Default.ShowLinkNames = true
	__Classdiagram__00000000_Default.IsInRenameMode = false
	__Classdiagram__00000000_Default.IsExpanded = true
	__Classdiagram__00000000_Default.NodeGongStructsIsExpanded = true
	__Classdiagram__00000000_Default.NodeGongStructNodeExpansion = `[true]`
	__Classdiagram__00000000_Default.NodeGongEnumsIsExpanded = false
	__Classdiagram__00000000_Default.NodeGongEnumNodeExpansion = ``
	__Classdiagram__00000000_Default.NodeGongNotesIsExpanded = false
	__Classdiagram__00000000_Default.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__00000000_Diagram_Package_created_the_2025_12_18T17_41_37Z.Name = `Diagram Package created the 2025-12-18T17:41:37Z`
	__DiagramPackage__00000000_Diagram_Package_created_the_2025_12_18T17_41_37Z.Path = ``
	__DiagramPackage__00000000_Diagram_Package_created_the_2025_12_18T17_41_37Z.GongModelPath = ``
	__DiagramPackage__00000000_Diagram_Package_created_the_2025_12_18T17_41_37Z.AbsolutePathToDiagramPackage = ``

	__GongStructShape__00000000_Default_Product.Name = `Default-Product`
	__GongStructShape__00000000_Default_Product.X = 11.000000
	__GongStructShape__00000000_Default_Product.Y = 20.000000
	__GongStructShape__00000000_Default_Product.IdentifierMeta = ref_models.Product{}
	__GongStructShape__00000000_Default_Product.Width = 240.000000
	__GongStructShape__00000000_Default_Product.Height = 146.000000
	__GongStructShape__00000000_Default_Product.IsSelected = false

	__LinkShape__00000000_SubProducts.Name = `SubProducts`
	__LinkShape__00000000_SubProducts.IdentifierMeta = ref_models.Product{}.SubProducts
	__LinkShape__00000000_SubProducts.FieldTypeIdentifierMeta = ref_models.Product{}
	__LinkShape__00000000_SubProducts.FieldOffsetX = 0.000000
	__LinkShape__00000000_SubProducts.FieldOffsetY = 0.000000
	__LinkShape__00000000_SubProducts.TargetMultiplicity = models.MANY
	__LinkShape__00000000_SubProducts.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000000_SubProducts.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000000_SubProducts.SourceMultiplicity = models.MANY
	__LinkShape__00000000_SubProducts.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000000_SubProducts.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000000_SubProducts.X = 371.000000
	__LinkShape__00000000_SubProducts.Y = 51.500000
	__LinkShape__00000000_SubProducts.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000000_SubProducts.StartRatio = 0.216650
	__LinkShape__00000000_SubProducts.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000000_SubProducts.EndRatio = 0.730348
	__LinkShape__00000000_SubProducts.CornerOffsetRatio = 1.656893

	// Setup of pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__00000000_Default.GongStructShapes = append(__Classdiagram__00000000_Default.GongStructShapes, __GongStructShape__00000000_Default_Product)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__00000000_Diagram_Package_created_the_2025_12_18T17_41_37Z.Classdiagrams = append(__DiagramPackage__00000000_Diagram_Package_created_the_2025_12_18T17_41_37Z.Classdiagrams, __Classdiagram__00000000_Default)
	__DiagramPackage__00000000_Diagram_Package_created_the_2025_12_18T17_41_37Z.SelectedClassdiagram = __Classdiagram__00000000_Default
	// setup of GongStructShape instances pointers
	__GongStructShape__00000000_Default_Product.LinkShapes = append(__GongStructShape__00000000_Default_Product.LinkShapes, __LinkShape__00000000_SubProducts)
	// setup of LinkShape instances pointers
}

