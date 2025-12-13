package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
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

	const __write__local_time = "2025-08-31 10:26:36.005588 CEST"
	const __write__utc_time__ = "2025-08-31 08:26:36.005588 UTC"

	const __commitId__ = "0000000124"

	// Declaration of instances to stage

	__AttributeShape__000000_Name := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000001_Booleanfield_renamed_renamed := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000002_Date := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000003_Date2 := (&models.AttributeShape{}).Stage(stage)

	__Classdiagram__000000_1_First := (&models.Classdiagram{}).Stage(stage)
	__Classdiagram__000001_2_Second_ := (&models.Classdiagram{}).Stage(stage)
	__Classdiagram__000002_0_Zero := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T05_37_56Z := (&models.DiagramPackage{}).Stage(stage)

	__GongEnumShape__000000_Default_AEnumType := (&models.GongEnumShape{}).Stage(stage)
	__GongEnumShape__000001_Default_BEnumType := (&models.GongEnumShape{}).Stage(stage)

	__GongEnumValueShape__000000_NoName_yet := (&models.GongEnumValueShape{}).Stage(stage)
	__GongEnumValueShape__000001_ENUM_VAL2 := (&models.GongEnumValueShape{}).Stage(stage)
	__GongEnumValueShape__000002_ENUM_VAL1_ := (&models.GongEnumValueShape{}).Stage(stage)

	__GongNoteShape__000000_Default_NoteOnOrganisation := (&models.GongNoteShape{}).Stage(stage)

	__GongStructShape__000000_Default_Astruct := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_Bstruct := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000002_0_Zero_Astruct := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000003_0_Zero_AstructBstruct2Use := (&models.GongStructShape{}).Stage(stage)

	__LinkShape__000000_Anarrayofb := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000001_Associationtob := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000002_Anarrayofb2Use := (&models.LinkShape{}).Stage(stage)

	// Setup of values

	__AttributeShape__000000_Name.Name = `Name`
	__AttributeShape__000000_Name.IdentifierMeta = ref_models.Astruct{}.Name
	__AttributeShape__000000_Name.FieldTypeAsString = ``
	__AttributeShape__000000_Name.Structname = `Astruct`
	__AttributeShape__000000_Name.Fieldtypename = `string`

	__AttributeShape__000001_Booleanfield_renamed_renamed.Name = `Booleanfield_renamed_renamed`
	__AttributeShape__000001_Booleanfield_renamed_renamed.IdentifierMeta = ref_models.Astruct{}.Booleanfield
	__AttributeShape__000001_Booleanfield_renamed_renamed.FieldTypeAsString = ``
	__AttributeShape__000001_Booleanfield_renamed_renamed.Structname = `Astruct`
	__AttributeShape__000001_Booleanfield_renamed_renamed.Fieldtypename = `bool`

	__AttributeShape__000002_Date.Name = `Date`
	__AttributeShape__000002_Date.IdentifierMeta = ref_models.Astruct{}.Date
	__AttributeShape__000002_Date.FieldTypeAsString = ``
	__AttributeShape__000002_Date.Structname = `Astruct`
	__AttributeShape__000002_Date.Fieldtypename = `Time`

	__AttributeShape__000003_Date2.Name = `Date2`
	__AttributeShape__000003_Date2.IdentifierMeta = ref_models.Astruct{}.Date2
	__AttributeShape__000003_Date2.FieldTypeAsString = ``
	__AttributeShape__000003_Date2.Structname = `Astruct`
	__AttributeShape__000003_Date2.Fieldtypename = `Time`

	__Classdiagram__000000_1_First.Name = `1. First`
	__Classdiagram__000000_1_First.Description = ``
	__Classdiagram__000000_1_First.IsIncludedInStaticWebSite = false
	__Classdiagram__000000_1_First.IsInRenameMode = false
	__Classdiagram__000000_1_First.IsExpanded = true
	__Classdiagram__000000_1_First.NodeGongStructsIsExpanded = false
	__Classdiagram__000000_1_First.NodeGongStructNodeExpansion = `[false]`
	__Classdiagram__000000_1_First.NodeGongEnumsIsExpanded = false
	__Classdiagram__000000_1_First.NodeGongEnumNodeExpansion = `[false]`
	__Classdiagram__000000_1_First.NodeGongNotesIsExpanded = true
	__Classdiagram__000000_1_First.NodeGongNoteNodeExpansion = `[true]`

	__Classdiagram__000001_2_Second_.Name = `2. Second.`
	__Classdiagram__000001_2_Second_.Description = ``
	__Classdiagram__000001_2_Second_.IsIncludedInStaticWebSite = false
	__Classdiagram__000001_2_Second_.IsInRenameMode = false
	__Classdiagram__000001_2_Second_.IsExpanded = false
	__Classdiagram__000001_2_Second_.NodeGongStructsIsExpanded = false
	__Classdiagram__000001_2_Second_.NodeGongStructNodeExpansion = ``
	__Classdiagram__000001_2_Second_.NodeGongEnumsIsExpanded = false
	__Classdiagram__000001_2_Second_.NodeGongEnumNodeExpansion = ``
	__Classdiagram__000001_2_Second_.NodeGongNotesIsExpanded = false
	__Classdiagram__000001_2_Second_.NodeGongNoteNodeExpansion = ``

	__Classdiagram__000002_0_Zero.Name = `0. Zero`
	__Classdiagram__000002_0_Zero.Description = ``
	__Classdiagram__000002_0_Zero.IsIncludedInStaticWebSite = false
	__Classdiagram__000002_0_Zero.IsInRenameMode = false
	__Classdiagram__000002_0_Zero.IsExpanded = true
	__Classdiagram__000002_0_Zero.NodeGongStructsIsExpanded = true
	__Classdiagram__000002_0_Zero.NodeGongStructNodeExpansion = `[true]`
	__Classdiagram__000002_0_Zero.NodeGongEnumsIsExpanded = false
	__Classdiagram__000002_0_Zero.NodeGongEnumNodeExpansion = ``
	__Classdiagram__000002_0_Zero.NodeGongNotesIsExpanded = false
	__Classdiagram__000002_0_Zero.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T05_37_56Z.Name = `Diagram Package created the 2025-06-04T05:37:56Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T05_37_56Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T05_37_56Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T05_37_56Z.AbsolutePathToDiagramPackage = ``

	__GongEnumShape__000000_Default_AEnumType.Name = `Default-AEnumType`
	__GongEnumShape__000000_Default_AEnumType.X = 534.000000
	__GongEnumShape__000000_Default_AEnumType.Y = 27.000000
	__GongEnumShape__000000_Default_AEnumType.IdentifierMeta = new(ref_models.AEnumType)
	__GongEnumShape__000000_Default_AEnumType.Width = 240.000000
	__GongEnumShape__000000_Default_AEnumType.Height = 103.000000
	__GongEnumShape__000000_Default_AEnumType.IsExpanded = false

	__GongEnumShape__000001_Default_BEnumType.Name = `Default-BEnumType`
	__GongEnumShape__000001_Default_BEnumType.X = 74.000000
	__GongEnumShape__000001_Default_BEnumType.Y = 191.000000
	__GongEnumShape__000001_Default_BEnumType.IdentifierMeta = new(ref_models.BEnumType)
	__GongEnumShape__000001_Default_BEnumType.Width = 240.000000
	__GongEnumShape__000001_Default_BEnumType.Height = 63.000000
	__GongEnumShape__000001_Default_BEnumType.IsExpanded = false

	__GongEnumValueShape__000000_NoName_yet.Name = `NoName yet`
	__GongEnumValueShape__000000_NoName_yet.IdentifierMeta = ref_models.ENUM_VAL1

	__GongEnumValueShape__000001_ENUM_VAL2.Name = `ENUM_VAL2`
	__GongEnumValueShape__000001_ENUM_VAL2.IdentifierMeta = ref_models.ENUM_VAL2

	__GongEnumValueShape__000002_ENUM_VAL1_.Name = `ENUM_VAL1_____`
	__GongEnumValueShape__000002_ENUM_VAL1_.IdentifierMeta = ref_models.ENUM_VAL1

	__GongNoteShape__000000_Default_NoteOnOrganisation.Name = `Default-NoteOnOrganisation`
	__GongNoteShape__000000_Default_NoteOnOrganisation.Identifier = `ref_models.NoteOnOrganisation`
	__GongNoteShape__000000_Default_NoteOnOrganisation.Body = `Note on the organisation

and [models.Bstruct],
and [models.Astruct],
and [models.Astruct],
and [models.Astruct],
and [models.Astruct.Associationtob],
having the following const exported identifier allows for
referencing the note from the UML note and allows for
renaming

# This is heading 1

## This is heading 1.1

 -
`
	__GongNoteShape__000000_Default_NoteOnOrganisation.BodyHTML = `<p>Note on the organisation
<p>and <a href="/models#Bstruct">models.Bstruct</a>,
and <a href="/models#Astruct">models.Astruct</a>,
and <a href="/models#Astruct">models.Astruct</a>,
and <a href="/models#Astruct">models.Astruct</a>,
and <a href="/models#Astruct.Associationtob">models.Astruct.Associationtob</a>,
having the following const exported identifier allows for
referencing the note from the UML note and allows for
renaming
<h3 id="hdr-This_is_heading_1">This is heading 1</h3>
<p>## This is heading 1.1
<pre>-
</pre>
`
	__GongNoteShape__000000_Default_NoteOnOrganisation.X = 44.000000
	__GongNoteShape__000000_Default_NoteOnOrganisation.Y = 344.000000
	__GongNoteShape__000000_Default_NoteOnOrganisation.Width = 625.000000
	__GongNoteShape__000000_Default_NoteOnOrganisation.Height = 226.000000
	__GongNoteShape__000000_Default_NoteOnOrganisation.Matched = false
	__GongNoteShape__000000_Default_NoteOnOrganisation.IsExpanded = false

	__GongStructShape__000000_Default_Astruct.Name = `Default-Astruct`
	__GongStructShape__000000_Default_Astruct.X = 127.000000
	__GongStructShape__000000_Default_Astruct.Y = 46.000000
	__GongStructShape__000000_Default_Astruct.IdentifierMeta = ref_models.Astruct{}
	__GongStructShape__000000_Default_Astruct.Width = 240.000000
	__GongStructShape__000000_Default_Astruct.Height = 103.000000
	__GongStructShape__000000_Default_Astruct.IsSelected = false

	__GongStructShape__000001_Default_Bstruct.Name = `Default-Bstruct`
	__GongStructShape__000001_Default_Bstruct.X = 813.000000
	__GongStructShape__000001_Default_Bstruct.Y = 97.000000
	__GongStructShape__000001_Default_Bstruct.IdentifierMeta = ref_models.Bstruct{}
	__GongStructShape__000001_Default_Bstruct.Width = 240.000000
	__GongStructShape__000001_Default_Bstruct.Height = 227.000000
	__GongStructShape__000001_Default_Bstruct.IsSelected = false

	__GongStructShape__000002_0_Zero_Astruct.Name = `0. Zero-Astruct`
	__GongStructShape__000002_0_Zero_Astruct.X = 56.000000
	__GongStructShape__000002_0_Zero_Astruct.Y = 105.000000
	__GongStructShape__000002_0_Zero_Astruct.IdentifierMeta = ref_models.Astruct{}
	__GongStructShape__000002_0_Zero_Astruct.Width = 240.000000
	__GongStructShape__000002_0_Zero_Astruct.Height = 103.000000
	__GongStructShape__000002_0_Zero_Astruct.IsSelected = false

	__GongStructShape__000003_0_Zero_AstructBstruct2Use.Name = `0. Zero-AstructBstruct2Use`
	__GongStructShape__000003_0_Zero_AstructBstruct2Use.X = 607.000000
	__GongStructShape__000003_0_Zero_AstructBstruct2Use.Y = 131.000000
	__GongStructShape__000003_0_Zero_AstructBstruct2Use.IdentifierMeta = ref_models.AstructBstruct2Use{}
	__GongStructShape__000003_0_Zero_AstructBstruct2Use.Width = 240.000000
	__GongStructShape__000003_0_Zero_AstructBstruct2Use.Height = 63.000000
	__GongStructShape__000003_0_Zero_AstructBstruct2Use.IsSelected = false

	__LinkShape__000000_Anarrayofb.Name = `Anarrayofb`
	__LinkShape__000000_Anarrayofb.IdentifierMeta = ref_models.Astruct{}.Anarrayofb
	__LinkShape__000000_Anarrayofb.FieldTypeIdentifierMeta = ref_models.Bstruct{}
	__LinkShape__000000_Anarrayofb.FieldOffsetX = 0.000000
	__LinkShape__000000_Anarrayofb.FieldOffsetY = 0.000000
	__LinkShape__000000_Anarrayofb.TargetMultiplicity = models.MANY
	__LinkShape__000000_Anarrayofb.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000000_Anarrayofb.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000000_Anarrayofb.SourceMultiplicity = models.MANY
	__LinkShape__000000_Anarrayofb.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000000_Anarrayofb.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000000_Anarrayofb.X = 918.500000
	__LinkShape__000000_Anarrayofb.Y = 117.000000
	__LinkShape__000000_Anarrayofb.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_Anarrayofb.StartRatio = 0.500000
	__LinkShape__000000_Anarrayofb.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_Anarrayofb.EndRatio = 0.863115
	__LinkShape__000000_Anarrayofb.CornerOffsetRatio = 1.285547

	__LinkShape__000001_Associationtob.Name = `Associationtob`
	__LinkShape__000001_Associationtob.IdentifierMeta = ref_models.Astruct{}.Associationtob
	__LinkShape__000001_Associationtob.FieldTypeIdentifierMeta = ref_models.Bstruct{}
	__LinkShape__000001_Associationtob.FieldOffsetX = 0.000000
	__LinkShape__000001_Associationtob.FieldOffsetY = 0.000000
	__LinkShape__000001_Associationtob.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000001_Associationtob.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000001_Associationtob.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000001_Associationtob.SourceMultiplicity = models.MANY
	__LinkShape__000001_Associationtob.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000001_Associationtob.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000001_Associationtob.X = 918.500000
	__LinkShape__000001_Associationtob.Y = 117.000000
	__LinkShape__000001_Associationtob.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_Associationtob.StartRatio = 0.500000
	__LinkShape__000001_Associationtob.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_Associationtob.EndRatio = 0.500000
	__LinkShape__000001_Associationtob.CornerOffsetRatio = 1.380000

	__LinkShape__000002_Anarrayofb2Use.Name = `Anarrayofb2Use`
	__LinkShape__000002_Anarrayofb2Use.IdentifierMeta = ref_models.Astruct{}.Anarrayofb2Use
	__LinkShape__000002_Anarrayofb2Use.FieldTypeIdentifierMeta = ref_models.AstructBstruct2Use{}
	__LinkShape__000002_Anarrayofb2Use.FieldOffsetX = 0.000000
	__LinkShape__000002_Anarrayofb2Use.FieldOffsetY = 0.000000
	__LinkShape__000002_Anarrayofb2Use.TargetMultiplicity = models.MANY
	__LinkShape__000002_Anarrayofb2Use.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000002_Anarrayofb2Use.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000002_Anarrayofb2Use.SourceMultiplicity = models.MANY
	__LinkShape__000002_Anarrayofb2Use.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000002_Anarrayofb2Use.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000002_Anarrayofb2Use.X = 540.000000
	__LinkShape__000002_Anarrayofb2Use.Y = 233.500000
	__LinkShape__000002_Anarrayofb2Use.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000002_Anarrayofb2Use.StartRatio = 0.500000
	__LinkShape__000002_Anarrayofb2Use.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000002_Anarrayofb2Use.EndRatio = 0.500000
	__LinkShape__000002_Anarrayofb2Use.CornerOffsetRatio = 1.380000

	// Setup of pointers
	// setup of AttributeShape instances pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_1_First.GongStructShapes = append(__Classdiagram__000000_1_First.GongStructShapes, __GongStructShape__000000_Default_Astruct)
	__Classdiagram__000000_1_First.GongStructShapes = append(__Classdiagram__000000_1_First.GongStructShapes, __GongStructShape__000001_Default_Bstruct)
	__Classdiagram__000000_1_First.GongEnumShapes = append(__Classdiagram__000000_1_First.GongEnumShapes, __GongEnumShape__000000_Default_AEnumType)
	__Classdiagram__000000_1_First.GongEnumShapes = append(__Classdiagram__000000_1_First.GongEnumShapes, __GongEnumShape__000001_Default_BEnumType)
	__Classdiagram__000000_1_First.GongNoteShapes = append(__Classdiagram__000000_1_First.GongNoteShapes, __GongNoteShape__000000_Default_NoteOnOrganisation)
	__Classdiagram__000002_0_Zero.GongStructShapes = append(__Classdiagram__000002_0_Zero.GongStructShapes, __GongStructShape__000002_0_Zero_Astruct)
	__Classdiagram__000002_0_Zero.GongStructShapes = append(__Classdiagram__000002_0_Zero.GongStructShapes, __GongStructShape__000003_0_Zero_AstructBstruct2Use)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T05_37_56Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T05_37_56Z.Classdiagrams, __Classdiagram__000000_1_First)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T05_37_56Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T05_37_56Z.Classdiagrams, __Classdiagram__000001_2_Second_)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T05_37_56Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T05_37_56Z.Classdiagrams, __Classdiagram__000002_0_Zero)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_06_04T05_37_56Z.SelectedClassdiagram = __Classdiagram__000002_0_Zero
	// setup of GongEnumShape instances pointers
	__GongEnumShape__000000_Default_AEnumType.GongEnumValueShapes = append(__GongEnumShape__000000_Default_AEnumType.GongEnumValueShapes, __GongEnumValueShape__000002_ENUM_VAL1_)
	__GongEnumShape__000000_Default_AEnumType.GongEnumValueShapes = append(__GongEnumShape__000000_Default_AEnumType.GongEnumValueShapes, __GongEnumValueShape__000001_ENUM_VAL2)
	// setup of GongEnumValueShape instances pointers
	// setup of GongNoteShape instances pointers
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_Astruct.AttributeShapes = append(__GongStructShape__000000_Default_Astruct.AttributeShapes, __AttributeShape__000000_Name)
	__GongStructShape__000000_Default_Astruct.AttributeShapes = append(__GongStructShape__000000_Default_Astruct.AttributeShapes, __AttributeShape__000001_Booleanfield_renamed_renamed)
	__GongStructShape__000000_Default_Astruct.LinkShapes = append(__GongStructShape__000000_Default_Astruct.LinkShapes, __LinkShape__000000_Anarrayofb)
	__GongStructShape__000000_Default_Astruct.LinkShapes = append(__GongStructShape__000000_Default_Astruct.LinkShapes, __LinkShape__000001_Associationtob)
	__GongStructShape__000002_0_Zero_Astruct.AttributeShapes = append(__GongStructShape__000002_0_Zero_Astruct.AttributeShapes, __AttributeShape__000002_Date)
	__GongStructShape__000002_0_Zero_Astruct.AttributeShapes = append(__GongStructShape__000002_0_Zero_Astruct.AttributeShapes, __AttributeShape__000003_Date2)
	__GongStructShape__000002_0_Zero_Astruct.LinkShapes = append(__GongStructShape__000002_0_Zero_Astruct.LinkShapes, __LinkShape__000002_Anarrayofb2Use)
	// setup of LinkShape instances pointers
}
