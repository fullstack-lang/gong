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

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T05_37_56Z := (&models.DiagramPackage{}).Stage(stage)

	__GongEnumShape__000000_Default_AEnumType := (&models.GongEnumShape{}).Stage(stage)

	__GongEnumValueShape__000000_NoName_yet := (&models.GongEnumValueShape{}).Stage(stage)
	__GongEnumValueShape__000001_ENUM_VAL1 := (&models.GongEnumValueShape{}).Stage(stage)
	__GongEnumValueShape__000002_ENUM_VAL2 := (&models.GongEnumValueShape{}).Stage(stage)

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

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.Description = ``
	__Classdiagram__000000_Default.IsIncludedInStaticWebSite = false
	__Classdiagram__000000_Default.IsInRenameMode = false
	__Classdiagram__000000_Default.IsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructNodeExpansion = `[true]`
	__Classdiagram__000000_Default.NodeGongEnumsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongEnumNodeExpansion = `[true]`
	__Classdiagram__000000_Default.NodeGongNotesIsExpanded = false
	__Classdiagram__000000_Default.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T05_37_56Z.Name = `Diagram Package created the 2025-06-04T05:37:56Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T05_37_56Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T05_37_56Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T05_37_56Z.AbsolutePathToDiagramPackage = ``

	__GongEnumShape__000000_Default_AEnumType.Name = `Default-AEnumType`
	__GongEnumShape__000000_Default_AEnumType.X = 72.000000
	__GongEnumShape__000000_Default_AEnumType.Y = 226.000000

	//gong:ident [ref_models.AEnumType] comment added to overcome the problem with the comment map association
	__GongEnumShape__000000_Default_AEnumType.Identifier = `ref_models.AEnumType`
	__GongEnumShape__000000_Default_AEnumType.IdentifierMeta = new(ref_models.AEnumType)
	__GongEnumShape__000000_Default_AEnumType.Width = 240.000000
	__GongEnumShape__000000_Default_AEnumType.Height = 103.000000
	__GongEnumShape__000000_Default_AEnumType.IsExpanded = false

	__GongEnumValueShape__000000_NoName_yet.Name = `NoName yet`
	__GongEnumValueShape__000000_NoName_yet.IdentifierMeta = ref_models.ENUM_VAL1

	__GongEnumValueShape__000001_ENUM_VAL1.Name = `ENUM_VAL1`
	__GongEnumValueShape__000001_ENUM_VAL1.IdentifierMeta = ref_models.ENUM_VAL1

	__GongEnumValueShape__000002_ENUM_VAL2.Name = `ENUM_VAL2`
	__GongEnumValueShape__000002_ENUM_VAL2.IdentifierMeta = ref_models.ENUM_VAL2

	__GongStructShape__000000_Default_Astruct.Name = `Default-Astruct`
	__GongStructShape__000000_Default_Astruct.X = 153.000000
	__GongStructShape__000000_Default_Astruct.Y = 41.000000
	__GongStructShape__000000_Default_Astruct.IdentifierMeta = ref_models.Astruct{}
	__GongStructShape__000000_Default_Astruct.ShowNbInstances = false
	__GongStructShape__000000_Default_Astruct.NbInstances = 0
	__GongStructShape__000000_Default_Astruct.Width = 240.000000
	__GongStructShape__000000_Default_Astruct.Height = 83.000000
	__GongStructShape__000000_Default_Astruct.IsSelected = false

	__GongStructShape__000001_Default_Bstruct.Name = `Default-Bstruct`
	__GongStructShape__000001_Default_Bstruct.X = 964.000000
	__GongStructShape__000001_Default_Bstruct.Y = 110.000000
	__GongStructShape__000001_Default_Bstruct.IdentifierMeta = ref_models.Bstruct{}
	__GongStructShape__000001_Default_Bstruct.ShowNbInstances = false
	__GongStructShape__000001_Default_Bstruct.NbInstances = 0
	__GongStructShape__000001_Default_Bstruct.Width = 240.000000
	__GongStructShape__000001_Default_Bstruct.Height = 227.000000
	__GongStructShape__000001_Default_Bstruct.IsSelected = false

	__LinkShape__000000_Associationtob.Name = `Associationtob`

	//gong:ident [ref_models.Astruct.Associationtob] comment added to overcome the problem with the comment map association
	__LinkShape__000000_Associationtob.Identifier = `ref_models.Astruct.Associationtob`
	__LinkShape__000000_Associationtob.IdentifierMeta = ref_models.Astruct{}.Associationtob

	//gong:ident [ref_models.Bstruct] comment added to overcome the problem with the comment map association

	__LinkShape__000000_Associationtob.FieldTypeIdentifierMeta = ref_models.Bstruct{}
	__LinkShape__000000_Associationtob.FieldOffsetX = 0.000000
	__LinkShape__000000_Associationtob.FieldOffsetY = 0.000000
	__LinkShape__000000_Associationtob.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000000_Associationtob.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000000_Associationtob.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000000_Associationtob.SourceMultiplicity = models.MANY
	__LinkShape__000000_Associationtob.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000000_Associationtob.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000000_Associationtob.X = 414.500000
	__LinkShape__000000_Associationtob.Y = 119.000000
	__LinkShape__000000_Associationtob.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_Associationtob.StartRatio = 0.500000
	__LinkShape__000000_Associationtob.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_Associationtob.EndRatio = 0.233159
	__LinkShape__000000_Associationtob.CornerOffsetRatio = 1.380000

	__LinkShape__000001_Anarrayofb.Name = `Anarrayofb`

	//gong:ident [ref_models.Astruct.Anarrayofb] comment added to overcome the problem with the comment map association
	__LinkShape__000001_Anarrayofb.Identifier = `ref_models.Astruct.Anarrayofb`
	__LinkShape__000001_Anarrayofb.IdentifierMeta = ref_models.Astruct{}.Anarrayofb

	//gong:ident [ref_models.Bstruct] comment added to overcome the problem with the comment map association

	__LinkShape__000001_Anarrayofb.FieldTypeIdentifierMeta = ref_models.Bstruct{}
	__LinkShape__000001_Anarrayofb.FieldOffsetX = 0.000000
	__LinkShape__000001_Anarrayofb.FieldOffsetY = 0.000000
	__LinkShape__000001_Anarrayofb.TargetMultiplicity = models.MANY
	__LinkShape__000001_Anarrayofb.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000001_Anarrayofb.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000001_Anarrayofb.SourceMultiplicity = models.MANY
	__LinkShape__000001_Anarrayofb.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000001_Anarrayofb.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000001_Anarrayofb.X = 918.500000
	__LinkShape__000001_Anarrayofb.Y = 117.000000
	__LinkShape__000001_Anarrayofb.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_Anarrayofb.StartRatio = 0.500000
	__LinkShape__000001_Anarrayofb.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_Anarrayofb.EndRatio = 0.863115
	__LinkShape__000001_Anarrayofb.CornerOffsetRatio = 1.285547

	// Setup of pointers
	// setup of AttributeShape instances pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Astruct)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_Bstruct)
	__Classdiagram__000000_Default.GongEnumShapes = append(__Classdiagram__000000_Default.GongEnumShapes, __GongEnumShape__000000_Default_AEnumType)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T05_37_56Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T05_37_56Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T05_37_56Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongEnumShape instances pointers
	__GongEnumShape__000000_Default_AEnumType.GongEnumValueShapes = append(__GongEnumShape__000000_Default_AEnumType.GongEnumValueShapes, __GongEnumValueShape__000001_ENUM_VAL1)
	__GongEnumShape__000000_Default_AEnumType.GongEnumValueShapes = append(__GongEnumShape__000000_Default_AEnumType.GongEnumValueShapes, __GongEnumValueShape__000002_ENUM_VAL2)
	// setup of GongEnumValueShape instances pointers
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_Astruct.AttributeShapes = append(__GongStructShape__000000_Default_Astruct.AttributeShapes, __AttributeShape__000000_Name)
	__GongStructShape__000000_Default_Astruct.LinkShapes = append(__GongStructShape__000000_Default_Astruct.LinkShapes, __LinkShape__000000_Associationtob)
	__GongStructShape__000000_Default_Astruct.LinkShapes = append(__GongStructShape__000000_Default_Astruct.LinkShapes, __LinkShape__000001_Anarrayofb)
	// setup of LinkShape instances pointers
}
