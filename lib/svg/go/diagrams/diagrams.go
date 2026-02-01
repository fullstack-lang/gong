package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/lib/svg/go/models"
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

	// insertion point for declaration of instances to stage

	__AttributeShape__00000000_ := (&models.AttributeShape{Name: `IsEditable`}).Stage(stage)
	__AttributeShape__00000001_ := (&models.AttributeShape{Name: `IsSVGFrontEndFileGenerated`}).Stage(stage)
	__AttributeShape__00000002_ := (&models.AttributeShape{Name: `IsSVGBackEndFileGenerated`}).Stage(stage)
	__AttributeShape__00000003_ := (&models.AttributeShape{Name: `DefaultDirectoryForGeneratedImages`}).Stage(stage)
	__AttributeShape__00000004_ := (&models.AttributeShape{Name: `IsControlBannerHidden`}).Stage(stage)
	__AttributeShape__00000005_ := (&models.AttributeShape{Name: `OverrideWidth`}).Stage(stage)
	__AttributeShape__00000006_ := (&models.AttributeShape{Name: `OverriddenWidth`}).Stage(stage)
	__AttributeShape__00000007_ := (&models.AttributeShape{Name: `OverrideHeight`}).Stage(stage)
	__AttributeShape__00000008_ := (&models.AttributeShape{Name: `OverriddenHeight`}).Stage(stage)
	__AttributeShape__00000009_ := (&models.AttributeShape{Name: `Name`}).Stage(stage)
	__AttributeShape__00000010_ := (&models.AttributeShape{Name: `Name`}).Stage(stage)
	__AttributeShape__00000011_ := (&models.AttributeShape{Name: `X`}).Stage(stage)
	__AttributeShape__00000012_ := (&models.AttributeShape{Name: `Y`}).Stage(stage)
	__AttributeShape__00000013_ := (&models.AttributeShape{Name: `Width`}).Stage(stage)
	__AttributeShape__00000014_ := (&models.AttributeShape{Name: `Height`}).Stage(stage)
	__AttributeShape__00000015_ := (&models.AttributeShape{Name: `RX`}).Stage(stage)
	__AttributeShape__00000016_ := (&models.AttributeShape{Name: `StartOrientation`}).Stage(stage)
	__AttributeShape__00000017_ := (&models.AttributeShape{Name: `StartRatio`}).Stage(stage)
	__AttributeShape__00000018_ := (&models.AttributeShape{Name: `EndOrientation`}).Stage(stage)
	__AttributeShape__00000019_ := (&models.AttributeShape{Name: `EndRatio`}).Stage(stage)
	__AttributeShape__00000020_ := (&models.AttributeShape{Name: `HasStartArrow`}).Stage(stage)
	__AttributeShape__00000021_ := (&models.AttributeShape{Name: `StartArrowSize`}).Stage(stage)
	__AttributeShape__00000022_ := (&models.AttributeShape{Name: `EndArrowOffset`}).Stage(stage)
	__AttributeShape__00000023_ := (&models.AttributeShape{Name: `EndArrowSize`}).Stage(stage)
	__AttributeShape__00000024_ := (&models.AttributeShape{Name: `StartArrowOffset`}).Stage(stage)
	__AttributeShape__00000025_ := (&models.AttributeShape{Name: `Content`}).Stage(stage)
	__AttributeShape__00000026_ := (&models.AttributeShape{Name: `AutomaticLayout`}).Stage(stage)
	__AttributeShape__00000027_ := (&models.AttributeShape{Name: `LinkAnchorType`}).Stage(stage)
	__AttributeShape__00000028_ := (&models.AttributeShape{Name: `X_Offset`}).Stage(stage)
	__AttributeShape__00000029_ := (&models.AttributeShape{Name: `Y_Offset`}).Stage(stage)

	__Classdiagram__00000000_ := (&models.Classdiagram{Name: `Default`}).Stage(stage)

	__DiagramPackage__00000000_ := (&models.DiagramPackage{Name: `Diagram Package created the 2026-01-31T17:53:12Z`}).Stage(stage)

	__GongEnumShape__00000000_ := (&models.GongEnumShape{Name: `Default-AnchorType`}).Stage(stage)

	__GongEnumValueShape__00000000_ := (&models.GongEnumValueShape{Name: `ANCHOR_TOP`}).Stage(stage)
	__GongEnumValueShape__00000004_ := (&models.GongEnumValueShape{Name: `ANCHOR_BOTTOM`}).Stage(stage)
	__GongEnumValueShape__00000005_ := (&models.GongEnumValueShape{Name: `ANCHOR_LEFT`}).Stage(stage)
	__GongEnumValueShape__00000006_ := (&models.GongEnumValueShape{Name: `ANCHOR_RIGHT`}).Stage(stage)
	__GongEnumValueShape__00000007_ := (&models.GongEnumValueShape{Name: `ANCHOR_CENTER`}).Stage(stage)

	__GongStructShape__00000000_ := (&models.GongStructShape{Name: `Default-SVG`}).Stage(stage)
	__GongStructShape__00000001_ := (&models.GongStructShape{Name: `Default-Layer`}).Stage(stage)
	__GongStructShape__00000006_ := (&models.GongStructShape{Name: `Default-Rect`}).Stage(stage)
	__GongStructShape__00000007_ := (&models.GongStructShape{Name: `Default-Link`}).Stage(stage)
	__GongStructShape__00000008_ := (&models.GongStructShape{Name: `Default-RectAnchoredPath`}).Stage(stage)
	__GongStructShape__00000009_ := (&models.GongStructShape{Name: `Default-RectAnchoredRect`}).Stage(stage)
	__GongStructShape__00000010_ := (&models.GongStructShape{Name: `Default-RectAnchoredText`}).Stage(stage)
	__GongStructShape__00000011_ := (&models.GongStructShape{Name: `Default-LinkAnchoredText`}).Stage(stage)

	__LinkShape__00000000_ := (&models.LinkShape{Name: `Layers`}).Stage(stage)
	__LinkShape__00000001_ := (&models.LinkShape{Name: `Links`}).Stage(stage)
	__LinkShape__00000002_ := (&models.LinkShape{Name: `Rects`}).Stage(stage)
	__LinkShape__00000003_ := (&models.LinkShape{Name: `RectAnchoredTexts`}).Stage(stage)
	__LinkShape__00000004_ := (&models.LinkShape{Name: `RectAnchoredRects`}).Stage(stage)
	__LinkShape__00000005_ := (&models.LinkShape{Name: `RectAnchoredPaths`}).Stage(stage)
	__LinkShape__00000006_ := (&models.LinkShape{Name: `Start`}).Stage(stage)
	__LinkShape__00000007_ := (&models.LinkShape{Name: `End`}).Stage(stage)
	__LinkShape__00000008_ := (&models.LinkShape{Name: `TextAtArrowStart`}).Stage(stage)
	__LinkShape__00000009_ := (&models.LinkShape{Name: `TextAtArrowEnd`}).Stage(stage)

	// insertion point for initialization of values

	__AttributeShape__00000000_.Name = `IsEditable`
	__AttributeShape__00000000_.IdentifierMeta = ref_models.SVG{}.IsEditable
	__AttributeShape__00000000_.FieldTypeAsString = ``
	__AttributeShape__00000000_.Structname = `SVG`
	__AttributeShape__00000000_.Fieldtypename = `bool`

	__AttributeShape__00000001_.Name = `IsSVGFrontEndFileGenerated`
	__AttributeShape__00000001_.IdentifierMeta = ref_models.SVG{}.IsSVGFrontEndFileGenerated
	__AttributeShape__00000001_.FieldTypeAsString = ``
	__AttributeShape__00000001_.Structname = `SVG`
	__AttributeShape__00000001_.Fieldtypename = `bool`

	__AttributeShape__00000002_.Name = `IsSVGBackEndFileGenerated`
	__AttributeShape__00000002_.IdentifierMeta = ref_models.SVG{}.IsSVGBackEndFileGenerated
	__AttributeShape__00000002_.FieldTypeAsString = ``
	__AttributeShape__00000002_.Structname = `SVG`
	__AttributeShape__00000002_.Fieldtypename = `bool`

	__AttributeShape__00000003_.Name = `DefaultDirectoryForGeneratedImages`
	__AttributeShape__00000003_.IdentifierMeta = ref_models.SVG{}.DefaultDirectoryForGeneratedImages
	__AttributeShape__00000003_.FieldTypeAsString = ``
	__AttributeShape__00000003_.Structname = `SVG`
	__AttributeShape__00000003_.Fieldtypename = `string`

	__AttributeShape__00000004_.Name = `IsControlBannerHidden`
	__AttributeShape__00000004_.IdentifierMeta = ref_models.SVG{}.IsControlBannerHidden
	__AttributeShape__00000004_.FieldTypeAsString = ``
	__AttributeShape__00000004_.Structname = `SVG`
	__AttributeShape__00000004_.Fieldtypename = `bool`

	__AttributeShape__00000005_.Name = `OverrideWidth`
	__AttributeShape__00000005_.IdentifierMeta = ref_models.SVG{}.OverrideWidth
	__AttributeShape__00000005_.FieldTypeAsString = ``
	__AttributeShape__00000005_.Structname = `SVG`
	__AttributeShape__00000005_.Fieldtypename = `bool`

	__AttributeShape__00000006_.Name = `OverriddenWidth`
	__AttributeShape__00000006_.IdentifierMeta = ref_models.SVG{}.OverriddenWidth
	__AttributeShape__00000006_.FieldTypeAsString = ``
	__AttributeShape__00000006_.Structname = `SVG`
	__AttributeShape__00000006_.Fieldtypename = `float64`

	__AttributeShape__00000007_.Name = `OverrideHeight`
	__AttributeShape__00000007_.IdentifierMeta = ref_models.SVG{}.OverrideHeight
	__AttributeShape__00000007_.FieldTypeAsString = ``
	__AttributeShape__00000007_.Structname = `SVG`
	__AttributeShape__00000007_.Fieldtypename = `bool`

	__AttributeShape__00000008_.Name = `OverriddenHeight`
	__AttributeShape__00000008_.IdentifierMeta = ref_models.SVG{}.OverriddenHeight
	__AttributeShape__00000008_.FieldTypeAsString = ``
	__AttributeShape__00000008_.Structname = `SVG`
	__AttributeShape__00000008_.Fieldtypename = `float64`

	__AttributeShape__00000009_.Name = `Name`
	__AttributeShape__00000009_.IdentifierMeta = ref_models.Layer{}.Name
	__AttributeShape__00000009_.FieldTypeAsString = ``
	__AttributeShape__00000009_.Structname = `Layer`
	__AttributeShape__00000009_.Fieldtypename = `string`

	__AttributeShape__00000010_.Name = `Name`
	__AttributeShape__00000010_.IdentifierMeta = ref_models.Rect{}.Name
	__AttributeShape__00000010_.FieldTypeAsString = ``
	__AttributeShape__00000010_.Structname = `Rect`
	__AttributeShape__00000010_.Fieldtypename = `string`

	__AttributeShape__00000011_.Name = `X`
	__AttributeShape__00000011_.IdentifierMeta = ref_models.Rect{}.X
	__AttributeShape__00000011_.FieldTypeAsString = ``
	__AttributeShape__00000011_.Structname = `Rect`
	__AttributeShape__00000011_.Fieldtypename = `float64`

	__AttributeShape__00000012_.Name = `Y`
	__AttributeShape__00000012_.IdentifierMeta = ref_models.Rect{}.Y
	__AttributeShape__00000012_.FieldTypeAsString = ``
	__AttributeShape__00000012_.Structname = `Rect`
	__AttributeShape__00000012_.Fieldtypename = `float64`

	__AttributeShape__00000013_.Name = `Width`
	__AttributeShape__00000013_.IdentifierMeta = ref_models.Rect{}.Width
	__AttributeShape__00000013_.FieldTypeAsString = ``
	__AttributeShape__00000013_.Structname = `Rect`
	__AttributeShape__00000013_.Fieldtypename = `float64`

	__AttributeShape__00000014_.Name = `Height`
	__AttributeShape__00000014_.IdentifierMeta = ref_models.Rect{}.Height
	__AttributeShape__00000014_.FieldTypeAsString = ``
	__AttributeShape__00000014_.Structname = `Rect`
	__AttributeShape__00000014_.Fieldtypename = `float64`

	__AttributeShape__00000015_.Name = `RX`
	__AttributeShape__00000015_.IdentifierMeta = ref_models.Rect{}.RX
	__AttributeShape__00000015_.FieldTypeAsString = ``
	__AttributeShape__00000015_.Structname = `Rect`
	__AttributeShape__00000015_.Fieldtypename = `float64`

	__AttributeShape__00000016_.Name = `StartOrientation`
	__AttributeShape__00000016_.IdentifierMeta = ref_models.Link{}.StartOrientation
	__AttributeShape__00000016_.FieldTypeAsString = ``
	__AttributeShape__00000016_.Structname = `Link`
	__AttributeShape__00000016_.Fieldtypename = `OrientationType`

	__AttributeShape__00000017_.Name = `StartRatio`
	__AttributeShape__00000017_.IdentifierMeta = ref_models.Link{}.StartRatio
	__AttributeShape__00000017_.FieldTypeAsString = ``
	__AttributeShape__00000017_.Structname = `Link`
	__AttributeShape__00000017_.Fieldtypename = `float64`

	__AttributeShape__00000018_.Name = `EndOrientation`
	__AttributeShape__00000018_.IdentifierMeta = ref_models.Link{}.EndOrientation
	__AttributeShape__00000018_.FieldTypeAsString = ``
	__AttributeShape__00000018_.Structname = `Link`
	__AttributeShape__00000018_.Fieldtypename = `OrientationType`

	__AttributeShape__00000019_.Name = `EndRatio`
	__AttributeShape__00000019_.IdentifierMeta = ref_models.Link{}.EndRatio
	__AttributeShape__00000019_.FieldTypeAsString = ``
	__AttributeShape__00000019_.Structname = `Link`
	__AttributeShape__00000019_.Fieldtypename = `float64`

	__AttributeShape__00000020_.Name = `HasStartArrow`
	__AttributeShape__00000020_.IdentifierMeta = ref_models.Link{}.HasStartArrow
	__AttributeShape__00000020_.FieldTypeAsString = ``
	__AttributeShape__00000020_.Structname = `Link`
	__AttributeShape__00000020_.Fieldtypename = `bool`

	__AttributeShape__00000021_.Name = `StartArrowSize`
	__AttributeShape__00000021_.IdentifierMeta = ref_models.Link{}.StartArrowSize
	__AttributeShape__00000021_.FieldTypeAsString = ``
	__AttributeShape__00000021_.Structname = `Link`
	__AttributeShape__00000021_.Fieldtypename = `float64`

	__AttributeShape__00000022_.Name = `EndArrowOffset`
	__AttributeShape__00000022_.IdentifierMeta = ref_models.Link{}.EndArrowOffset
	__AttributeShape__00000022_.FieldTypeAsString = ``
	__AttributeShape__00000022_.Structname = `Link`
	__AttributeShape__00000022_.Fieldtypename = `float64`

	__AttributeShape__00000023_.Name = `EndArrowSize`
	__AttributeShape__00000023_.IdentifierMeta = ref_models.Link{}.EndArrowSize
	__AttributeShape__00000023_.FieldTypeAsString = ``
	__AttributeShape__00000023_.Structname = `Link`
	__AttributeShape__00000023_.Fieldtypename = `float64`

	__AttributeShape__00000024_.Name = `StartArrowOffset`
	__AttributeShape__00000024_.IdentifierMeta = ref_models.Link{}.StartArrowOffset
	__AttributeShape__00000024_.FieldTypeAsString = ``
	__AttributeShape__00000024_.Structname = `Link`
	__AttributeShape__00000024_.Fieldtypename = `float64`

	__AttributeShape__00000025_.Name = `Content`
	__AttributeShape__00000025_.IdentifierMeta = ref_models.LinkAnchoredText{}.Content
	__AttributeShape__00000025_.FieldTypeAsString = ``
	__AttributeShape__00000025_.Structname = `LinkAnchoredText`
	__AttributeShape__00000025_.Fieldtypename = `string`

	__AttributeShape__00000026_.Name = `AutomaticLayout`
	__AttributeShape__00000026_.IdentifierMeta = ref_models.LinkAnchoredText{}.AutomaticLayout
	__AttributeShape__00000026_.FieldTypeAsString = ``
	__AttributeShape__00000026_.Structname = `LinkAnchoredText`
	__AttributeShape__00000026_.Fieldtypename = `bool`

	__AttributeShape__00000027_.Name = `LinkAnchorType`
	__AttributeShape__00000027_.IdentifierMeta = ref_models.LinkAnchoredText{}.LinkAnchorType
	__AttributeShape__00000027_.FieldTypeAsString = ``
	__AttributeShape__00000027_.Structname = `LinkAnchoredText`
	__AttributeShape__00000027_.Fieldtypename = `LinkAnchorType`

	__AttributeShape__00000028_.Name = `X_Offset`
	__AttributeShape__00000028_.IdentifierMeta = ref_models.LinkAnchoredText{}.X_Offset
	__AttributeShape__00000028_.FieldTypeAsString = ``
	__AttributeShape__00000028_.Structname = `LinkAnchoredText`
	__AttributeShape__00000028_.Fieldtypename = `float64`

	__AttributeShape__00000029_.Name = `Y_Offset`
	__AttributeShape__00000029_.IdentifierMeta = ref_models.LinkAnchoredText{}.Y_Offset
	__AttributeShape__00000029_.FieldTypeAsString = ``
	__AttributeShape__00000029_.Structname = `LinkAnchoredText`
	__AttributeShape__00000029_.Fieldtypename = `float64`

	__Classdiagram__00000000_.Name = `Default`
	__Classdiagram__00000000_.Description = ``
	__Classdiagram__00000000_.IsIncludedInStaticWebSite = false
	__Classdiagram__00000000_.ShowNbInstances = true
	__Classdiagram__00000000_.ShowMultiplicity = true
	__Classdiagram__00000000_.ShowLinkNames = true
	__Classdiagram__00000000_.IsInRenameMode = false
	__Classdiagram__00000000_.IsExpanded = false
	__Classdiagram__00000000_.NodeGongStructsIsExpanded = true
	__Classdiagram__00000000_.NodeGongStructNodeExpansion = `[false,false,false,false,false,false,false,true,true,false,false,false,false,false,false,false,false,false,false]`
	__Classdiagram__00000000_.NodeGongEnumsIsExpanded = false
	__Classdiagram__00000000_.NodeGongEnumNodeExpansion = `[true]`
	__Classdiagram__00000000_.NodeGongNotesIsExpanded = false
	__Classdiagram__00000000_.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__00000000_.Name = `Diagram Package created the 2026-01-31T17:53:12Z`
	__DiagramPackage__00000000_.Path = ``
	__DiagramPackage__00000000_.GongModelPath = ``
	__DiagramPackage__00000000_.AbsolutePathToDiagramPackage = ``

	__GongEnumShape__00000000_.Name = `Default-AnchorType`
	__GongEnumShape__00000000_.X = 344.000000
	__GongEnumShape__00000000_.Y = 301.000000
	__GongEnumShape__00000000_.IdentifierMeta = new(ref_models.AnchorType)
	__GongEnumShape__00000000_.Width = 240.000000
	__GongEnumShape__00000000_.Height = 163.000000
	__GongEnumShape__00000000_.IsExpanded = false

	__GongEnumValueShape__00000000_.Name = `ANCHOR_TOP`
	__GongEnumValueShape__00000000_.IdentifierMeta = ref_models.ANCHOR_TOP

	__GongEnumValueShape__00000004_.Name = `ANCHOR_BOTTOM`
	__GongEnumValueShape__00000004_.IdentifierMeta = ref_models.ANCHOR_BOTTOM

	__GongEnumValueShape__00000005_.Name = `ANCHOR_LEFT`
	__GongEnumValueShape__00000005_.IdentifierMeta = ref_models.ANCHOR_LEFT

	__GongEnumValueShape__00000006_.Name = `ANCHOR_RIGHT`
	__GongEnumValueShape__00000006_.IdentifierMeta = ref_models.ANCHOR_RIGHT

	__GongEnumValueShape__00000007_.Name = `ANCHOR_CENTER`
	__GongEnumValueShape__00000007_.IdentifierMeta = ref_models.ANCHOR_CENTER

	__GongStructShape__00000000_.Name = `Default-SVG`
	__GongStructShape__00000000_.X = 58.000000
	__GongStructShape__00000000_.Y = 20.000000
	__GongStructShape__00000000_.IdentifierMeta = ref_models.SVG{}
	__GongStructShape__00000000_.Width = 382.000000
	__GongStructShape__00000000_.Height = 243.000000
	__GongStructShape__00000000_.IsSelected = false

	__GongStructShape__00000001_.Name = `Default-Layer`
	__GongStructShape__00000001_.X = 579.000000
	__GongStructShape__00000001_.Y = 71.000000
	__GongStructShape__00000001_.IdentifierMeta = ref_models.Layer{}
	__GongStructShape__00000001_.Width = 240.000000
	__GongStructShape__00000001_.Height = 83.000000
	__GongStructShape__00000001_.IsSelected = false

	__GongStructShape__00000006_.Name = `Default-Rect`
	__GongStructShape__00000006_.X = 971.000000
	__GongStructShape__00000006_.Y = 318.000000
	__GongStructShape__00000006_.IdentifierMeta = ref_models.Rect{}
	__GongStructShape__00000006_.Width = 295.000000
	__GongStructShape__00000006_.Height = 248.000000
	__GongStructShape__00000006_.IsSelected = false

	__GongStructShape__00000007_.Name = `Default-Link`
	__GongStructShape__00000007_.X = 977.000000
	__GongStructShape__00000007_.Y = 46.000000
	__GongStructShape__00000007_.IdentifierMeta = ref_models.Link{}
	__GongStructShape__00000007_.Width = 287.000000
	__GongStructShape__00000007_.Height = 243.000000
	__GongStructShape__00000007_.IsSelected = false

	__GongStructShape__00000008_.Name = `Default-RectAnchoredPath`
	__GongStructShape__00000008_.X = 1483.000000
	__GongStructShape__00000008_.Y = 507.000000
	__GongStructShape__00000008_.IdentifierMeta = ref_models.RectAnchoredPath{}
	__GongStructShape__00000008_.Width = 240.000000
	__GongStructShape__00000008_.Height = 63.000000
	__GongStructShape__00000008_.IsSelected = false

	__GongStructShape__00000009_.Name = `Default-RectAnchoredRect`
	__GongStructShape__00000009_.X = 1479.000000
	__GongStructShape__00000009_.Y = 411.000000
	__GongStructShape__00000009_.IdentifierMeta = ref_models.RectAnchoredRect{}
	__GongStructShape__00000009_.Width = 240.000000
	__GongStructShape__00000009_.Height = 63.000000
	__GongStructShape__00000009_.IsSelected = false

	__GongStructShape__00000010_.Name = `Default-RectAnchoredText`
	__GongStructShape__00000010_.X = 1487.000000
	__GongStructShape__00000010_.Y = 324.000000
	__GongStructShape__00000010_.IdentifierMeta = ref_models.RectAnchoredText{}
	__GongStructShape__00000010_.Width = 240.000000
	__GongStructShape__00000010_.Height = 63.000000
	__GongStructShape__00000010_.IsSelected = false

	__GongStructShape__00000011_.Name = `Default-LinkAnchoredText`
	__GongStructShape__00000011_.X = 1470.000000
	__GongStructShape__00000011_.Y = 33.000000
	__GongStructShape__00000011_.IdentifierMeta = ref_models.LinkAnchoredText{}
	__GongStructShape__00000011_.Width = 268.000000
	__GongStructShape__00000011_.Height = 163.000000
	__GongStructShape__00000011_.IsSelected = false

	__LinkShape__00000000_.Name = `Layers`
	__LinkShape__00000000_.IdentifierMeta = ref_models.SVG{}.Layers
	__LinkShape__00000000_.FieldTypeIdentifierMeta = ref_models.Layer{}
	__LinkShape__00000000_.FieldOffsetX = 0.000000
	__LinkShape__00000000_.FieldOffsetY = 0.000000
	__LinkShape__00000000_.TargetMultiplicity = models.MANY
	__LinkShape__00000000_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000000_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000000_.SourceMultiplicity = models.MANY
	__LinkShape__00000000_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000000_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000000_.X = 421.000000
	__LinkShape__00000000_.Y = 52.500000
	__LinkShape__00000000_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000000_.StartRatio = 0.370249
	__LinkShape__00000000_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000000_.EndRatio = 0.602054
	__LinkShape__00000000_.CornerOffsetRatio = 1.082543

	__LinkShape__00000001_.Name = `Links`
	__LinkShape__00000001_.IdentifierMeta = ref_models.Layer{}.Links
	__LinkShape__00000001_.FieldTypeIdentifierMeta = ref_models.Link{}
	__LinkShape__00000001_.FieldOffsetX = 0.000000
	__LinkShape__00000001_.FieldOffsetY = 0.000000
	__LinkShape__00000001_.TargetMultiplicity = models.MANY
	__LinkShape__00000001_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000001_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000001_.SourceMultiplicity = models.MANY
	__LinkShape__00000001_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000001_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000001_.X = 1139.000000
	__LinkShape__00000001_.Y = 112.000000
	__LinkShape__00000001_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000001_.StartRatio = 0.500000
	__LinkShape__00000001_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000001_.EndRatio = 0.139796
	__LinkShape__00000001_.CornerOffsetRatio = 1.339714

	__LinkShape__00000002_.Name = `Rects`
	__LinkShape__00000002_.IdentifierMeta = ref_models.Layer{}.Rects
	__LinkShape__00000002_.FieldTypeIdentifierMeta = ref_models.Rect{}
	__LinkShape__00000002_.FieldOffsetX = 0.000000
	__LinkShape__00000002_.FieldOffsetY = 0.000000
	__LinkShape__00000002_.TargetMultiplicity = models.MANY
	__LinkShape__00000002_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000002_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000002_.SourceMultiplicity = models.MANY
	__LinkShape__00000002_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000002_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000002_.X = 1137.500000
	__LinkShape__00000002_.Y = 176.000000
	__LinkShape__00000002_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000002_.StartRatio = 0.500000
	__LinkShape__00000002_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000002_.EndRatio = 0.181332
	__LinkShape__00000002_.CornerOffsetRatio = 1.327214

	__LinkShape__00000003_.Name = `RectAnchoredTexts`
	__LinkShape__00000003_.IdentifierMeta = ref_models.Rect{}.RectAnchoredTexts
	__LinkShape__00000003_.FieldTypeIdentifierMeta = ref_models.RectAnchoredText{}
	__LinkShape__00000003_.FieldOffsetX = 0.000000
	__LinkShape__00000003_.FieldOffsetY = 0.000000
	__LinkShape__00000003_.TargetMultiplicity = models.MANY
	__LinkShape__00000003_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000003_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000003_.SourceMultiplicity = models.MANY
	__LinkShape__00000003_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000003_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000003_.X = 1538.500000
	__LinkShape__00000003_.Y = 288.500000
	__LinkShape__00000003_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000003_.StartRatio = 0.500000
	__LinkShape__00000003_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000003_.EndRatio = 0.500000
	__LinkShape__00000003_.CornerOffsetRatio = 1.134004

	__LinkShape__00000004_.Name = `RectAnchoredRects`
	__LinkShape__00000004_.IdentifierMeta = ref_models.Rect{}.RectAnchoredRects
	__LinkShape__00000004_.FieldTypeIdentifierMeta = ref_models.RectAnchoredRect{}
	__LinkShape__00000004_.FieldOffsetX = 0.000000
	__LinkShape__00000004_.FieldOffsetY = 0.000000
	__LinkShape__00000004_.TargetMultiplicity = models.MANY
	__LinkShape__00000004_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000004_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000004_.SourceMultiplicity = models.MANY
	__LinkShape__00000004_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000004_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000004_.X = 1540.000000
	__LinkShape__00000004_.Y = 340.500000
	__LinkShape__00000004_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000004_.StartRatio = 0.500000
	__LinkShape__00000004_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000004_.EndRatio = 0.500000
	__LinkShape__00000004_.CornerOffsetRatio = 1.380000

	__LinkShape__00000005_.Name = `RectAnchoredPaths`
	__LinkShape__00000005_.IdentifierMeta = ref_models.Rect{}.RectAnchoredPaths
	__LinkShape__00000005_.FieldTypeIdentifierMeta = ref_models.RectAnchoredPath{}
	__LinkShape__00000005_.FieldOffsetX = 0.000000
	__LinkShape__00000005_.FieldOffsetY = 0.000000
	__LinkShape__00000005_.TargetMultiplicity = models.MANY
	__LinkShape__00000005_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000005_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000005_.SourceMultiplicity = models.MANY
	__LinkShape__00000005_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000005_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000005_.X = 1541.000000
	__LinkShape__00000005_.Y = 391.000000
	__LinkShape__00000005_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000005_.StartRatio = 0.508035
	__LinkShape__00000005_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000005_.EndRatio = 0.500000
	__LinkShape__00000005_.CornerOffsetRatio = 1.150953

	__LinkShape__00000006_.Name = `Start`
	__LinkShape__00000006_.IdentifierMeta = ref_models.Link{}.Start
	__LinkShape__00000006_.FieldTypeIdentifierMeta = ref_models.Rect{}
	__LinkShape__00000006_.FieldOffsetX = 0.000000
	__LinkShape__00000006_.FieldOffsetY = 0.000000
	__LinkShape__00000006_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000006_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000006_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000006_.SourceMultiplicity = models.MANY
	__LinkShape__00000006_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000006_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000006_.X = 1405.500000
	__LinkShape__00000006_.Y = 297.000000
	__LinkShape__00000006_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000006_.StartRatio = 0.654199
	__LinkShape__00000006_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000006_.EndRatio = 0.935365
	__LinkShape__00000006_.CornerOffsetRatio = -0.524281

	__LinkShape__00000007_.Name = `End`
	__LinkShape__00000007_.IdentifierMeta = ref_models.Link{}.End
	__LinkShape__00000007_.FieldTypeIdentifierMeta = ref_models.Rect{}
	__LinkShape__00000007_.FieldOffsetX = 0.000000
	__LinkShape__00000007_.FieldOffsetY = 0.000000
	__LinkShape__00000007_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000007_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000007_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000007_.SourceMultiplicity = models.MANY
	__LinkShape__00000007_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000007_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000007_.X = 1405.500000
	__LinkShape__00000007_.Y = 297.000000
	__LinkShape__00000007_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000007_.StartRatio = 0.777656
	__LinkShape__00000007_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000007_.EndRatio = 0.657139
	__LinkShape__00000007_.CornerOffsetRatio = -0.454595

	__LinkShape__00000008_.Name = `TextAtArrowStart`
	__LinkShape__00000008_.IdentifierMeta = ref_models.Link{}.TextAtArrowStart
	__LinkShape__00000008_.FieldTypeIdentifierMeta = ref_models.LinkAnchoredText{}
	__LinkShape__00000008_.FieldOffsetX = 0.000000
	__LinkShape__00000008_.FieldOffsetY = 0.000000
	__LinkShape__00000008_.TargetMultiplicity = models.MANY
	__LinkShape__00000008_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000008_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000008_.SourceMultiplicity = models.MANY
	__LinkShape__00000008_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000008_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000008_.X = 1653.000000
	__LinkShape__00000008_.Y = 160.500000
	__LinkShape__00000008_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000008_.StartRatio = 0.308520
	__LinkShape__00000008_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000008_.EndRatio = 0.269757
	__LinkShape__00000008_.CornerOffsetRatio = 1.217879

	__LinkShape__00000009_.Name = `TextAtArrowEnd`
	__LinkShape__00000009_.IdentifierMeta = ref_models.Link{}.TextAtArrowEnd
	__LinkShape__00000009_.FieldTypeIdentifierMeta = ref_models.LinkAnchoredText{}
	__LinkShape__00000009_.FieldOffsetX = 0.000000
	__LinkShape__00000009_.FieldOffsetY = 0.000000
	__LinkShape__00000009_.TargetMultiplicity = models.MANY
	__LinkShape__00000009_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000009_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000009_.SourceMultiplicity = models.MANY
	__LinkShape__00000009_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000009_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000009_.X = 1653.000000
	__LinkShape__00000009_.Y = 160.500000
	__LinkShape__00000009_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000009_.StartRatio = 0.308520
	__LinkShape__00000009_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000009_.EndRatio = 0.852580
	__LinkShape__00000009_.CornerOffsetRatio = 1.228332

	// insertion point for setup of pointers
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000000_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000001_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000006_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000007_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000008_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000009_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000010_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000011_)
	__Classdiagram__00000000_.GongEnumShapes = append(__Classdiagram__00000000_.GongEnumShapes, __GongEnumShape__00000000_)
	__DiagramPackage__00000000_.Classdiagrams = append(__DiagramPackage__00000000_.Classdiagrams, __Classdiagram__00000000_)
	__DiagramPackage__00000000_.SelectedClassdiagram = __Classdiagram__00000000_
	__GongEnumShape__00000000_.GongEnumValueShapes = append(__GongEnumShape__00000000_.GongEnumValueShapes, __GongEnumValueShape__00000000_)
	__GongEnumShape__00000000_.GongEnumValueShapes = append(__GongEnumShape__00000000_.GongEnumValueShapes, __GongEnumValueShape__00000004_)
	__GongEnumShape__00000000_.GongEnumValueShapes = append(__GongEnumShape__00000000_.GongEnumValueShapes, __GongEnumValueShape__00000005_)
	__GongEnumShape__00000000_.GongEnumValueShapes = append(__GongEnumShape__00000000_.GongEnumValueShapes, __GongEnumValueShape__00000006_)
	__GongEnumShape__00000000_.GongEnumValueShapes = append(__GongEnumShape__00000000_.GongEnumValueShapes, __GongEnumValueShape__00000007_)
	__GongStructShape__00000000_.AttributeShapes = append(__GongStructShape__00000000_.AttributeShapes, __AttributeShape__00000000_)
	__GongStructShape__00000000_.AttributeShapes = append(__GongStructShape__00000000_.AttributeShapes, __AttributeShape__00000001_)
	__GongStructShape__00000000_.AttributeShapes = append(__GongStructShape__00000000_.AttributeShapes, __AttributeShape__00000002_)
	__GongStructShape__00000000_.AttributeShapes = append(__GongStructShape__00000000_.AttributeShapes, __AttributeShape__00000003_)
	__GongStructShape__00000000_.AttributeShapes = append(__GongStructShape__00000000_.AttributeShapes, __AttributeShape__00000004_)
	__GongStructShape__00000000_.AttributeShapes = append(__GongStructShape__00000000_.AttributeShapes, __AttributeShape__00000005_)
	__GongStructShape__00000000_.AttributeShapes = append(__GongStructShape__00000000_.AttributeShapes, __AttributeShape__00000006_)
	__GongStructShape__00000000_.AttributeShapes = append(__GongStructShape__00000000_.AttributeShapes, __AttributeShape__00000007_)
	__GongStructShape__00000000_.AttributeShapes = append(__GongStructShape__00000000_.AttributeShapes, __AttributeShape__00000008_)
	__GongStructShape__00000000_.LinkShapes = append(__GongStructShape__00000000_.LinkShapes, __LinkShape__00000000_)
	__GongStructShape__00000001_.AttributeShapes = append(__GongStructShape__00000001_.AttributeShapes, __AttributeShape__00000009_)
	__GongStructShape__00000001_.LinkShapes = append(__GongStructShape__00000001_.LinkShapes, __LinkShape__00000001_)
	__GongStructShape__00000001_.LinkShapes = append(__GongStructShape__00000001_.LinkShapes, __LinkShape__00000002_)
	__GongStructShape__00000006_.AttributeShapes = append(__GongStructShape__00000006_.AttributeShapes, __AttributeShape__00000010_)
	__GongStructShape__00000006_.AttributeShapes = append(__GongStructShape__00000006_.AttributeShapes, __AttributeShape__00000011_)
	__GongStructShape__00000006_.AttributeShapes = append(__GongStructShape__00000006_.AttributeShapes, __AttributeShape__00000012_)
	__GongStructShape__00000006_.AttributeShapes = append(__GongStructShape__00000006_.AttributeShapes, __AttributeShape__00000013_)
	__GongStructShape__00000006_.AttributeShapes = append(__GongStructShape__00000006_.AttributeShapes, __AttributeShape__00000014_)
	__GongStructShape__00000006_.AttributeShapes = append(__GongStructShape__00000006_.AttributeShapes, __AttributeShape__00000015_)
	__GongStructShape__00000006_.LinkShapes = append(__GongStructShape__00000006_.LinkShapes, __LinkShape__00000003_)
	__GongStructShape__00000006_.LinkShapes = append(__GongStructShape__00000006_.LinkShapes, __LinkShape__00000004_)
	__GongStructShape__00000006_.LinkShapes = append(__GongStructShape__00000006_.LinkShapes, __LinkShape__00000005_)
	__GongStructShape__00000007_.AttributeShapes = append(__GongStructShape__00000007_.AttributeShapes, __AttributeShape__00000016_)
	__GongStructShape__00000007_.AttributeShapes = append(__GongStructShape__00000007_.AttributeShapes, __AttributeShape__00000017_)
	__GongStructShape__00000007_.AttributeShapes = append(__GongStructShape__00000007_.AttributeShapes, __AttributeShape__00000018_)
	__GongStructShape__00000007_.AttributeShapes = append(__GongStructShape__00000007_.AttributeShapes, __AttributeShape__00000019_)
	__GongStructShape__00000007_.AttributeShapes = append(__GongStructShape__00000007_.AttributeShapes, __AttributeShape__00000023_)
	__GongStructShape__00000007_.AttributeShapes = append(__GongStructShape__00000007_.AttributeShapes, __AttributeShape__00000022_)
	__GongStructShape__00000007_.AttributeShapes = append(__GongStructShape__00000007_.AttributeShapes, __AttributeShape__00000020_)
	__GongStructShape__00000007_.AttributeShapes = append(__GongStructShape__00000007_.AttributeShapes, __AttributeShape__00000021_)
	__GongStructShape__00000007_.AttributeShapes = append(__GongStructShape__00000007_.AttributeShapes, __AttributeShape__00000024_)
	__GongStructShape__00000007_.LinkShapes = append(__GongStructShape__00000007_.LinkShapes, __LinkShape__00000006_)
	__GongStructShape__00000007_.LinkShapes = append(__GongStructShape__00000007_.LinkShapes, __LinkShape__00000007_)
	__GongStructShape__00000007_.LinkShapes = append(__GongStructShape__00000007_.LinkShapes, __LinkShape__00000008_)
	__GongStructShape__00000007_.LinkShapes = append(__GongStructShape__00000007_.LinkShapes, __LinkShape__00000009_)
	__GongStructShape__00000011_.AttributeShapes = append(__GongStructShape__00000011_.AttributeShapes, __AttributeShape__00000025_)
	__GongStructShape__00000011_.AttributeShapes = append(__GongStructShape__00000011_.AttributeShapes, __AttributeShape__00000026_)
	__GongStructShape__00000011_.AttributeShapes = append(__GongStructShape__00000011_.AttributeShapes, __AttributeShape__00000027_)
	__GongStructShape__00000011_.AttributeShapes = append(__GongStructShape__00000011_.AttributeShapes, __AttributeShape__00000028_)
	__GongStructShape__00000011_.AttributeShapes = append(__GongStructShape__00000011_.AttributeShapes, __AttributeShape__00000029_)
}
