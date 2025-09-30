package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc2/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/test/test3/go/models"
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

	const __write__local_time = "2025-09-30 02:19:15.788226 CEST"
	const __write__utc_time__ = "2025-09-30 00:19:15.788226 UTC"

	const __commitId__ = "0000000019"

	// Declaration of instances to stage

	__AttributeShape__000000_Name := (&models.AttributeShape{}).Stage(stage)

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_13T23_34_45Z := (&models.DiagramPackage{}).Stage(stage)

	__GongStructShape__000000_Default_A := (&models.GongStructShape{}).Stage(stage)

	__LinkShape__000000_As := (&models.LinkShape{}).Stage(stage)

	// Setup of values

	__AttributeShape__000000_Name.Name = `Name`
	__AttributeShape__000000_Name.IdentifierMeta = ref_models.A{}.Name
	__AttributeShape__000000_Name.FieldTypeAsString = ``
	__AttributeShape__000000_Name.Structname = `A`
	__AttributeShape__000000_Name.Fieldtypename = `string`

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.Description = ``
	__Classdiagram__000000_Default.IsIncludedInStaticWebSite = true
	__Classdiagram__000000_Default.IsInRenameMode = false
	__Classdiagram__000000_Default.IsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructNodeExpansion = `[true]`
	__Classdiagram__000000_Default.NodeGongEnumsIsExpanded = false
	__Classdiagram__000000_Default.NodeGongEnumNodeExpansion = ``
	__Classdiagram__000000_Default.NodeGongNotesIsExpanded = false
	__Classdiagram__000000_Default.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_13T23_34_45Z.Name = `Diagram Package created the 2025-06-13T23:34:45Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_13T23_34_45Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_13T23_34_45Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_13T23_34_45Z.AbsolutePathToDiagramPackage = ``

	__GongStructShape__000000_Default_A.Name = `Default-A`
	__GongStructShape__000000_Default_A.X = 76.000000
	__GongStructShape__000000_Default_A.Y = 73.000000
	__GongStructShape__000000_Default_A.IdentifierMeta = ref_models.A{}
	__GongStructShape__000000_Default_A.Width = 240.000000
	__GongStructShape__000000_Default_A.Height = 124.000000
	__GongStructShape__000000_Default_A.IsSelected = false

	__LinkShape__000000_As.Name = `As`
	__LinkShape__000000_As.IdentifierMeta = ref_models.A{}.As
	__LinkShape__000000_As.FieldTypeIdentifierMeta = ref_models.A{}
	__LinkShape__000000_As.FieldOffsetX = 0.000000
	__LinkShape__000000_As.FieldOffsetY = 0.000000
	__LinkShape__000000_As.TargetMultiplicity = models.MANY
	__LinkShape__000000_As.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000000_As.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000000_As.SourceMultiplicity = models.MANY
	__LinkShape__000000_As.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000000_As.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000000_As.X = 437.000000
	__LinkShape__000000_As.Y = 71.500000
	__LinkShape__000000_As.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_As.StartRatio = 0.160486
	__LinkShape__000000_As.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_As.EndRatio = 0.883377
	__LinkShape__000000_As.CornerOffsetRatio = 1.522949

	// Setup of pointers
	// setup of AttributeShape instances pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_A)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_13T23_34_45Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_06_13T23_34_45Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_13T23_34_45Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_A.AttributeShapes = append(__GongStructShape__000000_Default_A.AttributeShapes, __AttributeShape__000000_Name)
	__GongStructShape__000000_Default_A.LinkShapes = append(__GongStructShape__000000_Default_A.LinkShapes, __LinkShape__000000_As)
	// setup of LinkShape instances pointers
}

