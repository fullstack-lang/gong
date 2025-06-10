package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc2/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/lib/doc2/go/models"
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

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)
	__Classdiagram__000001_spoil_diagram := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_53_27Z := (&models.DiagramPackage{}).Stage(stage)

	__GongNoteShape__000000_Default_NoteOnGongdoc := (&models.GongNoteShape{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.Description = `This diagram describes the model of the doc2 package. A DiagramPackage is composed of ClassDiagram and each ClassDiagram has shapes.`
	__Classdiagram__000000_Default.IsIncludedInStaticWebSite = true
	__Classdiagram__000000_Default.IsInRenameMode = false
	__Classdiagram__000000_Default.IsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructNodeExpansion = `[false,false,false,true,false,false,false,true]`
	__Classdiagram__000000_Default.NodeGongEnumsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongEnumNodeExpansion = ``
	__Classdiagram__000000_Default.NodeGongNotesIsExpanded = true
	__Classdiagram__000000_Default.NodeGongNoteNodeExpansion = `[true]`

	__Classdiagram__000001_spoil_diagram.Name = `spoil diagram`
	__Classdiagram__000001_spoil_diagram.Description = `Spoil diagram`
	__Classdiagram__000001_spoil_diagram.IsIncludedInStaticWebSite = false
	__Classdiagram__000001_spoil_diagram.IsInRenameMode = false
	__Classdiagram__000001_spoil_diagram.IsExpanded = true
	__Classdiagram__000001_spoil_diagram.NodeGongStructsIsExpanded = true
	__Classdiagram__000001_spoil_diagram.NodeGongStructNodeExpansion = ``
	__Classdiagram__000001_spoil_diagram.NodeGongEnumsIsExpanded = true
	__Classdiagram__000001_spoil_diagram.NodeGongEnumNodeExpansion = ``
	__Classdiagram__000001_spoil_diagram.NodeGongNotesIsExpanded = false
	__Classdiagram__000001_spoil_diagram.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_53_27Z.Name = `Diagram Package created the 2025-05-04T22:53:27Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_53_27Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_53_27Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_53_27Z.AbsolutePathToDiagramPackage = ``

	__GongNoteShape__000000_Default_NoteOnGongdoc.Name = `Default-NoteOnGongdoc`
	__GongNoteShape__000000_Default_NoteOnGongdoc.Identifier = `ref_models.NoteOnGongdoc`
	__GongNoteShape__000000_Default_NoteOnGongdoc.Body = `Note Example

This note can refers to [models.GongNoteShape]
or to [models.Classdiagram.GongNoteShapes]
or to [models.OrientationType]
`
	__GongNoteShape__000000_Default_NoteOnGongdoc.BodyHTML = `<p>Note Example
<p>This note can refers to <a href="/models#GongNoteShape">models.GongNoteShape</a>
or to <a href="/models#Classdiagram.GongNoteShapes">models.Classdiagram.GongNoteShapes</a>
or to <a href="/models#OrientationType">models.OrientationType</a>
`
	__GongNoteShape__000000_Default_NoteOnGongdoc.X = 176.000000
	__GongNoteShape__000000_Default_NoteOnGongdoc.Y = 231.000000
	__GongNoteShape__000000_Default_NoteOnGongdoc.Width = 497.999969
	__GongNoteShape__000000_Default_NoteOnGongdoc.Height = 119.000000
	__GongNoteShape__000000_Default_NoteOnGongdoc.Matched = false
	__GongNoteShape__000000_Default_NoteOnGongdoc.IsExpanded = false

	// Setup of pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongNoteShapes = append(__Classdiagram__000000_Default.GongNoteShapes, __GongNoteShape__000000_Default_NoteOnGongdoc)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_53_27Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_53_27Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_53_27Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_53_27Z.Classdiagrams, __Classdiagram__000001_spoil_diagram)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_04T22_53_27Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongNoteShape instances pointers
}
