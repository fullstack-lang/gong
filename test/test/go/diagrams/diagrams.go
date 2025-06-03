package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc2/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/test/test/go/models"
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

	__AttributeShape__000000_Name := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000001_Field := (&models.AttributeShape{}).Stage(stage)

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_03T03_09_51Z := (&models.DiagramPackage{}).Stage(stage)

	__GongStructShape__000000_Default_Astruct := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_Bstruct := (&models.GongStructShape{}).Stage(stage)

	__LinkShape__000000_Associationtob := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000001_Anarrayofb := (&models.LinkShape{}).Stage(stage)

	// Setup of values

	__AttributeShape__000000_Name.Name = `Name`

	//gong:ident [ref_models.Astruct.Name] comment added to overcome the problem with the comment map association
	__AttributeShape__000000_Name.Identifier = `ref_models.Astruct.Name`
	__AttributeShape__000000_Name.IdentifierMeta = ref_models.Astruct{}.Name
	__AttributeShape__000000_Name.FieldTypeAsString = ``
	__AttributeShape__000000_Name.Structname = `Astruct`
	__AttributeShape__000000_Name.Fieldtypename = `string`

	__AttributeShape__000001_Field.Name = `Field`

	//gong:ident [ref_models.Astruct.Field] comment added to overcome the problem with the comment map association
	__AttributeShape__000001_Field.Identifier = `ref_models.Astruct.Field`
	__AttributeShape__000001_Field.IdentifierMeta = ref_models.Astruct{}.Field
	__AttributeShape__000001_Field.FieldTypeAsString = ``
	__AttributeShape__000001_Field.Structname = `Astruct`
	__AttributeShape__000001_Field.Fieldtypename = `any`

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.Description = ``
	__Classdiagram__000000_Default.IsIncludedInStaticWebSite = false
	__Classdiagram__000000_Default.IsInRenameMode = false
	__Classdiagram__000000_Default.IsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructNodeExpansion = `[true]`
	__Classdiagram__000000_Default.NodeGongEnumsIsExpanded = false
	__Classdiagram__000000_Default.NodeGongEnumNodeExpansion = ``
	__Classdiagram__000000_Default.NodeGongNotesIsExpanded = false
	__Classdiagram__000000_Default.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_03T03_09_51Z.Name = `Diagram Package created the 2025-06-03T03:09:51Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_03T03_09_51Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_03T03_09_51Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_03T03_09_51Z.AbsolutePathToDiagramPackage = ``

	__GongStructShape__000000_Default_Astruct.Name = `Default-Astruct`
	__GongStructShape__000000_Default_Astruct.X = 83.000000
	__GongStructShape__000000_Default_Astruct.Y = 21.000000

	//gong:ident [.] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_Astruct.Identifier = `.`
	__GongStructShape__000000_Default_Astruct.IdentifierMeta = ref_models.Astruct{}
	__GongStructShape__000000_Default_Astruct.ShowNbInstances = false
	__GongStructShape__000000_Default_Astruct.NbInstances = 0
	__GongStructShape__000000_Default_Astruct.Width = 240.000000
	__GongStructShape__000000_Default_Astruct.Height = 103.000000
	__GongStructShape__000000_Default_Astruct.IsSelected = false

	__GongStructShape__000001_Default_Bstruct.Name = `Default-Bstruct`
	__GongStructShape__000001_Default_Bstruct.X = 602.000000
	__GongStructShape__000001_Default_Bstruct.Y = 79.000000

	//gong:ident [.] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_Bstruct.Identifier = `.`
	__GongStructShape__000001_Default_Bstruct.IdentifierMeta = ref_models.Bstruct{}
	__GongStructShape__000001_Default_Bstruct.ShowNbInstances = false
	__GongStructShape__000001_Default_Bstruct.NbInstances = 0
	__GongStructShape__000001_Default_Bstruct.Width = 240.000000
	__GongStructShape__000001_Default_Bstruct.Height = 151.000000
	__GongStructShape__000001_Default_Bstruct.IsSelected = false

	__LinkShape__000000_Associationtob.Name = `Associationtob`

	//gong:ident [ref_models.Astruct.Associationtob] comment added to overcome the problem with the comment map association
	__LinkShape__000000_Associationtob.Identifier = `ref_models.Astruct.Associationtob`
	__LinkShape__000000_Associationtob.IdentifierMeta = ref_models.Astruct{}.Associationtob

	//gong:ident [ref_models.Bstruct] comment added to overcome the problem with the comment map association
	__LinkShape__000000_Associationtob.Fieldtypename = `ref_models.Bstruct`
	__LinkShape__000000_Associationtob.FieldTypeIdentifierMeta = ref_models.Bstruct{}
	__LinkShape__000000_Associationtob.FieldOffsetX = 0.000000
	__LinkShape__000000_Associationtob.FieldOffsetY = 0.000000
	__LinkShape__000000_Associationtob.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000000_Associationtob.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000000_Associationtob.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000000_Associationtob.SourceMultiplicity = models.MANY
	__LinkShape__000000_Associationtob.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000000_Associationtob.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000000_Associationtob.X = 412.000000
	__LinkShape__000000_Associationtob.Y = 75.000000
	__LinkShape__000000_Associationtob.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_Associationtob.StartRatio = 0.500000
	__LinkShape__000000_Associationtob.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_Associationtob.EndRatio = 0.120964
	__LinkShape__000000_Associationtob.CornerOffsetRatio = 1.380000

	__LinkShape__000001_Anarrayofb.Name = `Anarrayofb`

	//gong:ident [ref_models.Astruct.Anarrayofb] comment added to overcome the problem with the comment map association
	__LinkShape__000001_Anarrayofb.Identifier = `ref_models.Astruct.Anarrayofb`
	__LinkShape__000001_Anarrayofb.IdentifierMeta = ref_models.Astruct{}.Anarrayofb

	//gong:ident [ref_models.Bstruct] comment added to overcome the problem with the comment map association
	__LinkShape__000001_Anarrayofb.Fieldtypename = `ref_models.Bstruct`
	__LinkShape__000001_Anarrayofb.FieldTypeIdentifierMeta = ref_models.Bstruct{}
	__LinkShape__000001_Anarrayofb.FieldOffsetX = 0.000000
	__LinkShape__000001_Anarrayofb.FieldOffsetY = 0.000000
	__LinkShape__000001_Anarrayofb.TargetMultiplicity = models.MANY
	__LinkShape__000001_Anarrayofb.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000001_Anarrayofb.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000001_Anarrayofb.SourceMultiplicity = models.MANY
	__LinkShape__000001_Anarrayofb.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000001_Anarrayofb.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000001_Anarrayofb.X = 702.500000
	__LinkShape__000001_Anarrayofb.Y = 81.500000
	__LinkShape__000001_Anarrayofb.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_Anarrayofb.StartRatio = 0.500000
	__LinkShape__000001_Anarrayofb.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_Anarrayofb.EndRatio = 0.955401
	__LinkShape__000001_Anarrayofb.CornerOffsetRatio = 1.380000

	// Setup of pointers
	// setup of AttributeShape instances pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Astruct)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_Bstruct)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_03T03_09_51Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_06_03T03_09_51Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_03T03_09_51Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_Astruct.AttributeShapes = append(__GongStructShape__000000_Default_Astruct.AttributeShapes, __AttributeShape__000000_Name)
	__GongStructShape__000000_Default_Astruct.AttributeShapes = append(__GongStructShape__000000_Default_Astruct.AttributeShapes, __AttributeShape__000001_Field)
	__GongStructShape__000000_Default_Astruct.LinkShapes = append(__GongStructShape__000000_Default_Astruct.LinkShapes, __LinkShape__000000_Associationtob)
	__GongStructShape__000000_Default_Astruct.LinkShapes = append(__GongStructShape__000000_Default_Astruct.LinkShapes, __LinkShape__000001_Anarrayofb)
	// setup of LinkShape instances pointers
}
