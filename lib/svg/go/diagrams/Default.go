package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/lib/doc/go/models"
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

	__Field__000000_Name := (&models.Field{}).Stage(stage)
	__Field__000001_IsEditable := (&models.Field{}).Stage(stage)
	__Field__000002_Display := (&models.Field{}).Stage(stage)

	__GongStructShape__000000_Default_SVG := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_Layer := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000002_Default_Link := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000003_Default_Rect := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000004_Default_RectLinkLink := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000005_Default_RectAnchoredRect := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000006_Default_RectAnchoredPath := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000007_Default_RectAnchoredText := (&models.GongStructShape{}).Stage(stage)

	__Link__000000_Layers := (&models.Link{}).Stage(stage)
	__Link__000001_Rects := (&models.Link{}).Stage(stage)
	__Link__000002_Links := (&models.Link{}).Stage(stage)
	__Link__000003_RectLinkLinks := (&models.Link{}).Stage(stage)
	__Link__000004_RectAnchoredTexts := (&models.Link{}).Stage(stage)
	__Link__000005_RectAnchoredRects := (&models.Link{}).Stage(stage)
	__Link__000006_RectAnchoredPaths := (&models.Link{}).Stage(stage)

	__Position__000000_Pos_Default_SVG := (&models.Position{}).Stage(stage)
	__Position__000001_Pos_Default_Layer := (&models.Position{}).Stage(stage)
	__Position__000002_Pos_Default_Link := (&models.Position{}).Stage(stage)
	__Position__000003_Pos_Default_Rect := (&models.Position{}).Stage(stage)
	__Position__000004_Pos_Default_RectLinkLink := (&models.Position{}).Stage(stage)
	__Position__000005_Pos_Default_RectAnchoredRect := (&models.Position{}).Stage(stage)
	__Position__000006_Pos_Default_RectAnchoredPath := (&models.Position{}).Stage(stage)
	__Position__000007_Pos_Default_RectAnchoredText := (&models.Position{}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_SVG_and_Default_Layer := (&models.Vertice{}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Layer_and_Default_Rect := (&models.Vertice{}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Layer_and_Default_Link := (&models.Vertice{}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Layer_and_Default_RectLinkLink := (&models.Vertice{}).Stage(stage)
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_Rect_and_Default_RectAnchoredText := (&models.Vertice{}).Stage(stage)
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_Rect_and_Default_RectAnchoredRect := (&models.Vertice{}).Stage(stage)
	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_Rect_and_Default_RectAnchoredPath := (&models.Vertice{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	__Field__000000_Name.Name = `Name`

	//gong:ident [ref_models.SVG.Name] comment added to overcome the problem with the comment map association
	__Field__000000_Name.Identifier = `ref_models.SVG.Name`
	__Field__000000_Name.FieldTypeAsString = ``
	__Field__000000_Name.Structname = `SVG`
	__Field__000000_Name.Fieldtypename = `string`

	__Field__000001_IsEditable.Name = `IsEditable`

	//gong:ident [ref_models.SVG.IsEditable] comment added to overcome the problem with the comment map association
	__Field__000001_IsEditable.Identifier = `ref_models.SVG.IsEditable`
	__Field__000001_IsEditable.FieldTypeAsString = ``
	__Field__000001_IsEditable.Structname = `SVG`
	__Field__000001_IsEditable.Fieldtypename = `bool`

	__Field__000002_Display.Name = `Display`

	//gong:ident [ref_models.Layer.Display] comment added to overcome the problem with the comment map association
	__Field__000002_Display.Identifier = `ref_models.Layer.Display`
	__Field__000002_Display.FieldTypeAsString = ``
	__Field__000002_Display.Structname = `Layer`
	__Field__000002_Display.Fieldtypename = `bool`

	__GongStructShape__000000_Default_SVG.Name = `Default-SVG`

	//gong:ident [ref_models.SVG] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_SVG.Identifier = `ref_models.SVG`
	__GongStructShape__000000_Default_SVG.ShowNbInstances = false
	__GongStructShape__000000_Default_SVG.NbInstances = 0
	__GongStructShape__000000_Default_SVG.Width = 240.000000
	__GongStructShape__000000_Default_SVG.Height = 93.000000
	__GongStructShape__000000_Default_SVG.IsSelected = false

	__GongStructShape__000001_Default_Layer.Name = `Default-Layer`

	//gong:ident [ref_models.Layer] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_Layer.Identifier = `ref_models.Layer`
	__GongStructShape__000001_Default_Layer.ShowNbInstances = false
	__GongStructShape__000001_Default_Layer.NbInstances = 0
	__GongStructShape__000001_Default_Layer.Width = 240.000000
	__GongStructShape__000001_Default_Layer.Height = 78.000000
	__GongStructShape__000001_Default_Layer.IsSelected = false

	__GongStructShape__000002_Default_Link.Name = `Default-Link`

	//gong:ident [ref_models.Link] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_Default_Link.Identifier = `ref_models.Link`
	__GongStructShape__000002_Default_Link.ShowNbInstances = false
	__GongStructShape__000002_Default_Link.NbInstances = 0
	__GongStructShape__000002_Default_Link.Width = 240.000000
	__GongStructShape__000002_Default_Link.Height = 63.000000
	__GongStructShape__000002_Default_Link.IsSelected = false

	__GongStructShape__000003_Default_Rect.Name = `Default-Rect`

	//gong:ident [ref_models.Rect] comment added to overcome the problem with the comment map association
	__GongStructShape__000003_Default_Rect.Identifier = `ref_models.Rect`
	__GongStructShape__000003_Default_Rect.ShowNbInstances = false
	__GongStructShape__000003_Default_Rect.NbInstances = 0
	__GongStructShape__000003_Default_Rect.Width = 240.000000
	__GongStructShape__000003_Default_Rect.Height = 123.000000
	__GongStructShape__000003_Default_Rect.IsSelected = false

	__GongStructShape__000004_Default_RectLinkLink.Name = `Default-RectLinkLink`

	//gong:ident [ref_models.RectLinkLink] comment added to overcome the problem with the comment map association
	__GongStructShape__000004_Default_RectLinkLink.Identifier = `ref_models.RectLinkLink`
	__GongStructShape__000004_Default_RectLinkLink.ShowNbInstances = false
	__GongStructShape__000004_Default_RectLinkLink.NbInstances = 0
	__GongStructShape__000004_Default_RectLinkLink.Width = 240.000000
	__GongStructShape__000004_Default_RectLinkLink.Height = 63.000000
	__GongStructShape__000004_Default_RectLinkLink.IsSelected = false

	__GongStructShape__000005_Default_RectAnchoredRect.Name = `Default-RectAnchoredRect`

	//gong:ident [ref_models.RectAnchoredRect] comment added to overcome the problem with the comment map association
	__GongStructShape__000005_Default_RectAnchoredRect.Identifier = `ref_models.RectAnchoredRect`
	__GongStructShape__000005_Default_RectAnchoredRect.ShowNbInstances = false
	__GongStructShape__000005_Default_RectAnchoredRect.NbInstances = 0
	__GongStructShape__000005_Default_RectAnchoredRect.Width = 240.000000
	__GongStructShape__000005_Default_RectAnchoredRect.Height = 63.000000
	__GongStructShape__000005_Default_RectAnchoredRect.IsSelected = false

	__GongStructShape__000006_Default_RectAnchoredPath.Name = `Default-RectAnchoredPath`

	//gong:ident [ref_models.RectAnchoredPath] comment added to overcome the problem with the comment map association
	__GongStructShape__000006_Default_RectAnchoredPath.Identifier = `ref_models.RectAnchoredPath`
	__GongStructShape__000006_Default_RectAnchoredPath.ShowNbInstances = false
	__GongStructShape__000006_Default_RectAnchoredPath.NbInstances = 0
	__GongStructShape__000006_Default_RectAnchoredPath.Width = 240.000000
	__GongStructShape__000006_Default_RectAnchoredPath.Height = 63.000000
	__GongStructShape__000006_Default_RectAnchoredPath.IsSelected = false

	__GongStructShape__000007_Default_RectAnchoredText.Name = `Default-RectAnchoredText`

	//gong:ident [ref_models.RectAnchoredText] comment added to overcome the problem with the comment map association
	__GongStructShape__000007_Default_RectAnchoredText.Identifier = `ref_models.RectAnchoredText`
	__GongStructShape__000007_Default_RectAnchoredText.ShowNbInstances = false
	__GongStructShape__000007_Default_RectAnchoredText.NbInstances = 0
	__GongStructShape__000007_Default_RectAnchoredText.Width = 240.000000
	__GongStructShape__000007_Default_RectAnchoredText.Height = 63.000000
	__GongStructShape__000007_Default_RectAnchoredText.IsSelected = false

	__Link__000000_Layers.Name = `Layers`

	//gong:ident [ref_models.SVG.Layers] comment added to overcome the problem with the comment map association
	__Link__000000_Layers.Identifier = `ref_models.SVG.Layers`

	//gong:ident [ref_models.Layer] comment added to overcome the problem with the comment map association
	__Link__000000_Layers.Fieldtypename = `ref_models.Layer`
	__Link__000000_Layers.FieldOffsetX = 0.000000
	__Link__000000_Layers.FieldOffsetY = 0.000000
	__Link__000000_Layers.TargetMultiplicity = models.MANY
	__Link__000000_Layers.TargetMultiplicityOffsetX = 0.000000
	__Link__000000_Layers.TargetMultiplicityOffsetY = 0.000000
	__Link__000000_Layers.SourceMultiplicity = models.MANY
	__Link__000000_Layers.SourceMultiplicityOffsetX = 0.000000
	__Link__000000_Layers.SourceMultiplicityOffsetY = 0.000000
	__Link__000000_Layers.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Layers.StartRatio = 0.500000
	__Link__000000_Layers.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Layers.EndRatio = 0.500000
	__Link__000000_Layers.CornerOffsetRatio = 1.380000

	__Link__000001_Rects.Name = `Rects`

	//gong:ident [ref_models.Layer.Rects] comment added to overcome the problem with the comment map association
	__Link__000001_Rects.Identifier = `ref_models.Layer.Rects`

	//gong:ident [ref_models.Rect] comment added to overcome the problem with the comment map association
	__Link__000001_Rects.Fieldtypename = `ref_models.Rect`
	__Link__000001_Rects.FieldOffsetX = 0.000000
	__Link__000001_Rects.FieldOffsetY = 0.000000
	__Link__000001_Rects.TargetMultiplicity = models.MANY
	__Link__000001_Rects.TargetMultiplicityOffsetX = 0.000000
	__Link__000001_Rects.TargetMultiplicityOffsetY = 0.000000
	__Link__000001_Rects.SourceMultiplicity = models.MANY
	__Link__000001_Rects.SourceMultiplicityOffsetX = 0.000000
	__Link__000001_Rects.SourceMultiplicityOffsetY = 0.000000
	__Link__000001_Rects.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Rects.StartRatio = 0.500000
	__Link__000001_Rects.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Rects.EndRatio = 0.500000
	__Link__000001_Rects.CornerOffsetRatio = 1.380000

	__Link__000002_Links.Name = `Links`

	//gong:ident [ref_models.Layer.Links] comment added to overcome the problem with the comment map association
	__Link__000002_Links.Identifier = `ref_models.Layer.Links`

	//gong:ident [ref_models.Link] comment added to overcome the problem with the comment map association
	__Link__000002_Links.Fieldtypename = `ref_models.Link`
	__Link__000002_Links.FieldOffsetX = 0.000000
	__Link__000002_Links.FieldOffsetY = 0.000000
	__Link__000002_Links.TargetMultiplicity = models.MANY
	__Link__000002_Links.TargetMultiplicityOffsetX = 0.000000
	__Link__000002_Links.TargetMultiplicityOffsetY = 0.000000
	__Link__000002_Links.SourceMultiplicity = models.MANY
	__Link__000002_Links.SourceMultiplicityOffsetX = 0.000000
	__Link__000002_Links.SourceMultiplicityOffsetY = 0.000000
	__Link__000002_Links.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_Links.StartRatio = 0.500000
	__Link__000002_Links.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_Links.EndRatio = 0.500000
	__Link__000002_Links.CornerOffsetRatio = 1.380000

	__Link__000003_RectLinkLinks.Name = `RectLinkLinks`

	//gong:ident [ref_models.Layer.RectLinkLinks] comment added to overcome the problem with the comment map association
	__Link__000003_RectLinkLinks.Identifier = `ref_models.Layer.RectLinkLinks`

	//gong:ident [ref_models.RectLinkLink] comment added to overcome the problem with the comment map association
	__Link__000003_RectLinkLinks.Fieldtypename = `ref_models.RectLinkLink`
	__Link__000003_RectLinkLinks.FieldOffsetX = 0.000000
	__Link__000003_RectLinkLinks.FieldOffsetY = 0.000000
	__Link__000003_RectLinkLinks.TargetMultiplicity = models.MANY
	__Link__000003_RectLinkLinks.TargetMultiplicityOffsetX = 0.000000
	__Link__000003_RectLinkLinks.TargetMultiplicityOffsetY = 0.000000
	__Link__000003_RectLinkLinks.SourceMultiplicity = models.MANY
	__Link__000003_RectLinkLinks.SourceMultiplicityOffsetX = 0.000000
	__Link__000003_RectLinkLinks.SourceMultiplicityOffsetY = 0.000000
	__Link__000003_RectLinkLinks.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_RectLinkLinks.StartRatio = 0.500000
	__Link__000003_RectLinkLinks.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_RectLinkLinks.EndRatio = 0.500000
	__Link__000003_RectLinkLinks.CornerOffsetRatio = 1.380000

	__Link__000004_RectAnchoredTexts.Name = `RectAnchoredTexts`

	//gong:ident [ref_models.Rect.RectAnchoredTexts] comment added to overcome the problem with the comment map association
	__Link__000004_RectAnchoredTexts.Identifier = `ref_models.Rect.RectAnchoredTexts`

	//gong:ident [ref_models.RectAnchoredText] comment added to overcome the problem with the comment map association
	__Link__000004_RectAnchoredTexts.Fieldtypename = `ref_models.RectAnchoredText`
	__Link__000004_RectAnchoredTexts.FieldOffsetX = 0.000000
	__Link__000004_RectAnchoredTexts.FieldOffsetY = 0.000000
	__Link__000004_RectAnchoredTexts.TargetMultiplicity = models.MANY
	__Link__000004_RectAnchoredTexts.TargetMultiplicityOffsetX = 0.000000
	__Link__000004_RectAnchoredTexts.TargetMultiplicityOffsetY = 0.000000
	__Link__000004_RectAnchoredTexts.SourceMultiplicity = models.MANY
	__Link__000004_RectAnchoredTexts.SourceMultiplicityOffsetX = 0.000000
	__Link__000004_RectAnchoredTexts.SourceMultiplicityOffsetY = 0.000000
	__Link__000004_RectAnchoredTexts.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_RectAnchoredTexts.StartRatio = 0.787109
	__Link__000004_RectAnchoredTexts.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_RectAnchoredTexts.EndRatio = 0.500000
	__Link__000004_RectAnchoredTexts.CornerOffsetRatio = 1.380000

	__Link__000005_RectAnchoredRects.Name = `RectAnchoredRects`

	//gong:ident [ref_models.Rect.RectAnchoredRects] comment added to overcome the problem with the comment map association
	__Link__000005_RectAnchoredRects.Identifier = `ref_models.Rect.RectAnchoredRects`

	//gong:ident [ref_models.RectAnchoredRect] comment added to overcome the problem with the comment map association
	__Link__000005_RectAnchoredRects.Fieldtypename = `ref_models.RectAnchoredRect`
	__Link__000005_RectAnchoredRects.FieldOffsetX = 0.000000
	__Link__000005_RectAnchoredRects.FieldOffsetY = 0.000000
	__Link__000005_RectAnchoredRects.TargetMultiplicity = models.MANY
	__Link__000005_RectAnchoredRects.TargetMultiplicityOffsetX = 0.000000
	__Link__000005_RectAnchoredRects.TargetMultiplicityOffsetY = 0.000000
	__Link__000005_RectAnchoredRects.SourceMultiplicity = models.MANY
	__Link__000005_RectAnchoredRects.SourceMultiplicityOffsetX = 0.000000
	__Link__000005_RectAnchoredRects.SourceMultiplicityOffsetY = 0.000000
	__Link__000005_RectAnchoredRects.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000005_RectAnchoredRects.StartRatio = 0.273072
	__Link__000005_RectAnchoredRects.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000005_RectAnchoredRects.EndRatio = 0.500000
	__Link__000005_RectAnchoredRects.CornerOffsetRatio = 1.380000

	__Link__000006_RectAnchoredPaths.Name = `RectAnchoredPaths`

	//gong:ident [ref_models.Rect.RectAnchoredPaths] comment added to overcome the problem with the comment map association
	__Link__000006_RectAnchoredPaths.Identifier = `ref_models.Rect.RectAnchoredPaths`

	//gong:ident [ref_models.RectAnchoredPath] comment added to overcome the problem with the comment map association
	__Link__000006_RectAnchoredPaths.Fieldtypename = `ref_models.RectAnchoredPath`
	__Link__000006_RectAnchoredPaths.FieldOffsetX = 0.000000
	__Link__000006_RectAnchoredPaths.FieldOffsetY = 0.000000
	__Link__000006_RectAnchoredPaths.TargetMultiplicity = models.MANY
	__Link__000006_RectAnchoredPaths.TargetMultiplicityOffsetX = 0.000000
	__Link__000006_RectAnchoredPaths.TargetMultiplicityOffsetY = 0.000000
	__Link__000006_RectAnchoredPaths.SourceMultiplicity = models.MANY
	__Link__000006_RectAnchoredPaths.SourceMultiplicityOffsetX = 0.000000
	__Link__000006_RectAnchoredPaths.SourceMultiplicityOffsetY = 0.000000
	__Link__000006_RectAnchoredPaths.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000006_RectAnchoredPaths.StartRatio = 0.533141
	__Link__000006_RectAnchoredPaths.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000006_RectAnchoredPaths.EndRatio = 0.500000
	__Link__000006_RectAnchoredPaths.CornerOffsetRatio = 1.380000

	__Position__000000_Pos_Default_SVG.X = 54.000000
	__Position__000000_Pos_Default_SVG.Y = 23.000000
	__Position__000000_Pos_Default_SVG.Name = `Pos-Default-SVG`

	__Position__000001_Pos_Default_Layer.X = 469.000000
	__Position__000001_Pos_Default_Layer.Y = 22.000000
	__Position__000001_Pos_Default_Layer.Name = `Pos-Default-Layer`

	__Position__000002_Pos_Default_Link.X = 971.000000
	__Position__000002_Pos_Default_Link.Y = 33.000000
	__Position__000002_Pos_Default_Link.Name = `Pos-Default-Link`

	__Position__000003_Pos_Default_Rect.X = 975.000000
	__Position__000003_Pos_Default_Rect.Y = 137.000000
	__Position__000003_Pos_Default_Rect.Name = `Pos-Default-Rect`

	__Position__000004_Pos_Default_RectLinkLink.X = 979.000000
	__Position__000004_Pos_Default_RectLinkLink.Y = 293.000000
	__Position__000004_Pos_Default_RectLinkLink.Name = `Pos-Default-RectLinkLink`

	__Position__000005_Pos_Default_RectAnchoredRect.X = 1548.000000
	__Position__000005_Pos_Default_RectAnchoredRect.Y = 59.000000
	__Position__000005_Pos_Default_RectAnchoredRect.Name = `Pos-Default-RectAnchoredRect`

	__Position__000006_Pos_Default_RectAnchoredPath.X = 1551.000000
	__Position__000006_Pos_Default_RectAnchoredPath.Y = 173.000000
	__Position__000006_Pos_Default_RectAnchoredPath.Name = `Pos-Default-RectAnchoredPath`

	__Position__000007_Pos_Default_RectAnchoredText.X = 1557.000000
	__Position__000007_Pos_Default_RectAnchoredText.Y = 283.000000
	__Position__000007_Pos_Default_RectAnchoredText.Name = `Pos-Default-RectAnchoredText`

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_SVG_and_Default_Layer.X = 621.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_SVG_and_Default_Layer.Y = 61.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_SVG_and_Default_Layer.Name = `Verticle in class diagram Default in middle between Default-SVG and Default-Layer`

	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Layer_and_Default_Rect.X = 1050.000000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Layer_and_Default_Rect.Y = 130.000000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Layer_and_Default_Rect.Name = `Verticle in class diagram Default in middle between Default-Layer and Default-Rect`

	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Layer_and_Default_Link.X = 1051.000000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Layer_and_Default_Link.Y = 65.000000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Layer_and_Default_Link.Name = `Verticle in class diagram Default in middle between Default-Layer and Default-Link`

	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Layer_and_Default_RectLinkLink.X = 1044.000000
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Layer_and_Default_RectLinkLink.Y = 199.500000
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Layer_and_Default_RectLinkLink.Name = `Verticle in class diagram Default in middle between Default-Layer and Default-RectLinkLink`

	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_Rect_and_Default_RectAnchoredText.X = 864.000000
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_Rect_and_Default_RectAnchoredText.Y = 165.500000
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_Rect_and_Default_RectAnchoredText.Name = `Verticle in class diagram Default in middle between Default-Rect and Default-RectAnchoredText`

	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_Rect_and_Default_RectAnchoredRect.X = 893.500000
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_Rect_and_Default_RectAnchoredRect.Y = 149.000000
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_Rect_and_Default_RectAnchoredRect.Name = `Verticle in class diagram Default in middle between Default-Rect and Default-RectAnchoredRect`

	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_Rect_and_Default_RectAnchoredPath.X = 885.000000
	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_Rect_and_Default_RectAnchoredPath.Y = 155.500000
	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_Rect_and_Default_RectAnchoredPath.Name = `Verticle in class diagram Default in middle between Default-Rect and Default-RectAnchoredPath`

	// Setup of pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_SVG)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_Layer)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_Link)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000003_Default_Rect)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000004_Default_RectLinkLink)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000005_Default_RectAnchoredRect)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000006_Default_RectAnchoredPath)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000007_Default_RectAnchoredText)
	// setup of Field instances pointers
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_SVG.Position = __Position__000000_Pos_Default_SVG
	__GongStructShape__000000_Default_SVG.Fields = append(__GongStructShape__000000_Default_SVG.Fields, __Field__000000_Name)
	__GongStructShape__000000_Default_SVG.Fields = append(__GongStructShape__000000_Default_SVG.Fields, __Field__000001_IsEditable)
	__GongStructShape__000000_Default_SVG.Links = append(__GongStructShape__000000_Default_SVG.Links, __Link__000000_Layers)
	__GongStructShape__000001_Default_Layer.Position = __Position__000001_Pos_Default_Layer
	__GongStructShape__000001_Default_Layer.Fields = append(__GongStructShape__000001_Default_Layer.Fields, __Field__000002_Display)
	__GongStructShape__000001_Default_Layer.Links = append(__GongStructShape__000001_Default_Layer.Links, __Link__000001_Rects)
	__GongStructShape__000001_Default_Layer.Links = append(__GongStructShape__000001_Default_Layer.Links, __Link__000002_Links)
	__GongStructShape__000001_Default_Layer.Links = append(__GongStructShape__000001_Default_Layer.Links, __Link__000003_RectLinkLinks)
	__GongStructShape__000002_Default_Link.Position = __Position__000002_Pos_Default_Link
	__GongStructShape__000003_Default_Rect.Position = __Position__000003_Pos_Default_Rect
	__GongStructShape__000003_Default_Rect.Links = append(__GongStructShape__000003_Default_Rect.Links, __Link__000004_RectAnchoredTexts)
	__GongStructShape__000003_Default_Rect.Links = append(__GongStructShape__000003_Default_Rect.Links, __Link__000005_RectAnchoredRects)
	__GongStructShape__000003_Default_Rect.Links = append(__GongStructShape__000003_Default_Rect.Links, __Link__000006_RectAnchoredPaths)
	__GongStructShape__000004_Default_RectLinkLink.Position = __Position__000004_Pos_Default_RectLinkLink
	__GongStructShape__000005_Default_RectAnchoredRect.Position = __Position__000005_Pos_Default_RectAnchoredRect
	__GongStructShape__000006_Default_RectAnchoredPath.Position = __Position__000006_Pos_Default_RectAnchoredPath
	__GongStructShape__000007_Default_RectAnchoredText.Position = __Position__000007_Pos_Default_RectAnchoredText
	// setup of Link instances pointers
	__Link__000000_Layers.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_SVG_and_Default_Layer
	__Link__000001_Rects.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Layer_and_Default_Rect
	__Link__000002_Links.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Layer_and_Default_Link
	__Link__000003_RectLinkLinks.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Layer_and_Default_RectLinkLink
	__Link__000004_RectAnchoredTexts.Middlevertice = __Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_Rect_and_Default_RectAnchoredText
	__Link__000005_RectAnchoredRects.Middlevertice = __Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_Rect_and_Default_RectAnchoredRect
	__Link__000006_RectAnchoredPaths.Middlevertice = __Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_Rect_and_Default_RectAnchoredPath
	// setup of Position instances pointers
	// setup of Vertice instances pointers
}
