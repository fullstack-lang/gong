package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/lib/split/go/models"
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

	__AttributeShape__00000000_ := (&models.AttributeShape{Name: `StackName`}).Stage(stage)
	__AttributeShape__00000001_ := (&models.AttributeShape{Name: `IsSelectedView`}).Stage(stage)
	__AttributeShape__00000002_ := (&models.AttributeShape{Name: `ShowViewName`}).Stage(stage)
	__AttributeShape__00000003_ := (&models.AttributeShape{Name: `SVG`}).Stage(stage)
	__AttributeShape__00000004_ := (&models.AttributeShape{Name: `Height`}).Stage(stage)
	__AttributeShape__00000005_ := (&models.AttributeShape{Name: `Width`}).Stage(stage)
	__AttributeShape__00000006_ := (&models.AttributeShape{Name: `Width`}).Stage(stage)
	__AttributeShape__00000007_ := (&models.AttributeShape{Name: `Height`}).Stage(stage)
	__AttributeShape__00000008_ := (&models.AttributeShape{Name: `SVG`}).Stage(stage)
	__AttributeShape__00000009_ := (&models.AttributeShape{Name: `ShowNameInHeader`}).Stage(stage)
	__AttributeShape__00000010_ := (&models.AttributeShape{Name: `Size`}).Stage(stage)
	__AttributeShape__00000011_ := (&models.AttributeShape{Name: `IsAny`}).Stage(stage)
	__AttributeShape__00000012_ := (&models.AttributeShape{Name: `HasDiv`}).Stage(stage)
	__AttributeShape__00000013_ := (&models.AttributeShape{Name: `DivStyle`}).Stage(stage)
	__AttributeShape__00000014_ := (&models.AttributeShape{Name: `NoName yet`}).Stage(stage)

	__Classdiagram__00000000_ := (&models.Classdiagram{Name: `Default`}).Stage(stage)

	__DiagramPackage__00000000_ := (&models.DiagramPackage{Name: `Diagram Package created the 2025-06-12T00:16:22Z`}).Stage(stage)

	__GongEnumShape__00000000_ := (&models.GongEnumShape{Name: `Default-Direction`}).Stage(stage)

	__GongEnumValueShape__00000000_ := (&models.GongEnumValueShape{Name: `Horizontal`}).Stage(stage)
	__GongEnumValueShape__00000001_ := (&models.GongEnumValueShape{Name: `Vertical`}).Stage(stage)

	__GongStructShape__00000000_ := (&models.GongStructShape{Name: `Default-Split`}).Stage(stage)
	__GongStructShape__00000001_ := (&models.GongStructShape{Name: `Default-Xlsx`}).Stage(stage)
	__GongStructShape__00000002_ := (&models.GongStructShape{Name: `Default-View`}).Stage(stage)
	__GongStructShape__00000003_ := (&models.GongStructShape{Name: `Default-Tree`}).Stage(stage)
	__GongStructShape__00000004_ := (&models.GongStructShape{Name: `Default-Tone`}).Stage(stage)
	__GongStructShape__00000005_ := (&models.GongStructShape{Name: `Default-Title`}).Stage(stage)
	__GongStructShape__00000006_ := (&models.GongStructShape{Name: `Default-Table`}).Stage(stage)
	__GongStructShape__00000007_ := (&models.GongStructShape{Name: `Default-Svg`}).Stage(stage)
	__GongStructShape__00000008_ := (&models.GongStructShape{Name: `Default-Slider`}).Stage(stage)
	__GongStructShape__00000009_ := (&models.GongStructShape{Name: `Default-Markdown`}).Stage(stage)
	__GongStructShape__00000010_ := (&models.GongStructShape{Name: `Default-LogoOnTheRight`}).Stage(stage)
	__GongStructShape__00000011_ := (&models.GongStructShape{Name: `Default-LogoOnTheLeft`}).Stage(stage)
	__GongStructShape__00000012_ := (&models.GongStructShape{Name: `Default-Load`}).Stage(stage)
	__GongStructShape__00000013_ := (&models.GongStructShape{Name: `Default-Form`}).Stage(stage)
	__GongStructShape__00000014_ := (&models.GongStructShape{Name: `Default-FavIcon`}).Stage(stage)
	__GongStructShape__00000015_ := (&models.GongStructShape{Name: `Default-Cursor`}).Stage(stage)
	__GongStructShape__00000016_ := (&models.GongStructShape{Name: `Default-Button`}).Stage(stage)
	__GongStructShape__00000017_ := (&models.GongStructShape{Name: `Default-AsSplitArea`}).Stage(stage)
	__GongStructShape__00000019_ := (&models.GongStructShape{Name: `Default-AsSplit`}).Stage(stage)

	__LinkShape__00000001_ := (&models.LinkShape{Name: `AsSplit`}).Stage(stage)
	__LinkShape__00000002_ := (&models.LinkShape{Name: `Button`}).Stage(stage)
	__LinkShape__00000003_ := (&models.LinkShape{Name: `Cursor`}).Stage(stage)
	__LinkShape__00000004_ := (&models.LinkShape{Name: `Form`}).Stage(stage)
	__LinkShape__00000005_ := (&models.LinkShape{Name: `Load`}).Stage(stage)
	__LinkShape__00000006_ := (&models.LinkShape{Name: `Markdown`}).Stage(stage)
	__LinkShape__00000007_ := (&models.LinkShape{Name: `Slider`}).Stage(stage)
	__LinkShape__00000008_ := (&models.LinkShape{Name: `Split`}).Stage(stage)
	__LinkShape__00000009_ := (&models.LinkShape{Name: `Svg`}).Stage(stage)
	__LinkShape__00000010_ := (&models.LinkShape{Name: `Table`}).Stage(stage)
	__LinkShape__00000011_ := (&models.LinkShape{Name: `Tone`}).Stage(stage)
	__LinkShape__00000012_ := (&models.LinkShape{Name: `Tree`}).Stage(stage)
	__LinkShape__00000013_ := (&models.LinkShape{Name: `Xlsx`}).Stage(stage)
	__LinkShape__00000014_ := (&models.LinkShape{Name: `NoName yet`}).Stage(stage)
	__LinkShape__00000015_ := (&models.LinkShape{Name: `AsSplitAreas`}).Stage(stage)
	__LinkShape__00000016_ := (&models.LinkShape{Name: `RootAsSplitAreas`}).Stage(stage)

	// insertion point for initialization of values

	__AttributeShape__00000000_.Name = `StackName`
	__AttributeShape__00000000_.IdentifierMeta = ref_models.Split{}.StackName
	__AttributeShape__00000000_.FieldTypeAsString = ``
	__AttributeShape__00000000_.Structname = `Split`
	__AttributeShape__00000000_.Fieldtypename = `string`

	__AttributeShape__00000001_.Name = `IsSelectedView`
	__AttributeShape__00000001_.IdentifierMeta = ref_models.View{}.IsSelectedView
	__AttributeShape__00000001_.FieldTypeAsString = ``
	__AttributeShape__00000001_.Structname = `View`
	__AttributeShape__00000001_.Fieldtypename = `bool`

	__AttributeShape__00000002_.Name = `ShowViewName`
	__AttributeShape__00000002_.IdentifierMeta = ref_models.View{}.ShowViewName
	__AttributeShape__00000002_.FieldTypeAsString = ``
	__AttributeShape__00000002_.Structname = `View`
	__AttributeShape__00000002_.Fieldtypename = `bool`

	__AttributeShape__00000003_.Name = `SVG`
	__AttributeShape__00000003_.IdentifierMeta = ref_models.LogoOnTheRight{}.SVG
	__AttributeShape__00000003_.FieldTypeAsString = ``
	__AttributeShape__00000003_.Structname = `LogoOnTheRight`
	__AttributeShape__00000003_.Fieldtypename = `string`

	__AttributeShape__00000004_.Name = `Height`
	__AttributeShape__00000004_.IdentifierMeta = ref_models.LogoOnTheRight{}.Height
	__AttributeShape__00000004_.FieldTypeAsString = ``
	__AttributeShape__00000004_.Structname = `LogoOnTheRight`
	__AttributeShape__00000004_.Fieldtypename = `int`

	__AttributeShape__00000005_.Name = `Width`
	__AttributeShape__00000005_.IdentifierMeta = ref_models.LogoOnTheRight{}.Width
	__AttributeShape__00000005_.FieldTypeAsString = ``
	__AttributeShape__00000005_.Structname = `LogoOnTheRight`
	__AttributeShape__00000005_.Fieldtypename = `int`

	__AttributeShape__00000006_.Name = `Width`
	__AttributeShape__00000006_.IdentifierMeta = ref_models.LogoOnTheLeft{}.Width
	__AttributeShape__00000006_.FieldTypeAsString = ``
	__AttributeShape__00000006_.Structname = `LogoOnTheLeft`
	__AttributeShape__00000006_.Fieldtypename = `int`

	__AttributeShape__00000007_.Name = `Height`
	__AttributeShape__00000007_.IdentifierMeta = ref_models.LogoOnTheLeft{}.Height
	__AttributeShape__00000007_.FieldTypeAsString = ``
	__AttributeShape__00000007_.Structname = `LogoOnTheLeft`
	__AttributeShape__00000007_.Fieldtypename = `int`

	__AttributeShape__00000008_.Name = `SVG`
	__AttributeShape__00000008_.IdentifierMeta = ref_models.LogoOnTheLeft{}.SVG
	__AttributeShape__00000008_.FieldTypeAsString = ``
	__AttributeShape__00000008_.Structname = `LogoOnTheLeft`
	__AttributeShape__00000008_.Fieldtypename = `string`

	__AttributeShape__00000009_.Name = `ShowNameInHeader`
	__AttributeShape__00000009_.IdentifierMeta = ref_models.AsSplitArea{}.ShowNameInHeader
	__AttributeShape__00000009_.FieldTypeAsString = ``
	__AttributeShape__00000009_.Structname = `AsSplitArea`
	__AttributeShape__00000009_.Fieldtypename = `bool`

	__AttributeShape__00000010_.Name = `Size`
	__AttributeShape__00000010_.IdentifierMeta = ref_models.AsSplitArea{}.Size
	__AttributeShape__00000010_.FieldTypeAsString = ``
	__AttributeShape__00000010_.Structname = `AsSplitArea`
	__AttributeShape__00000010_.Fieldtypename = `float64`

	__AttributeShape__00000011_.Name = `IsAny`
	__AttributeShape__00000011_.IdentifierMeta = ref_models.AsSplitArea{}.IsAny
	__AttributeShape__00000011_.FieldTypeAsString = ``
	__AttributeShape__00000011_.Structname = `AsSplitArea`
	__AttributeShape__00000011_.Fieldtypename = `bool`

	__AttributeShape__00000012_.Name = `HasDiv`
	__AttributeShape__00000012_.IdentifierMeta = ref_models.AsSplitArea{}.HasDiv
	__AttributeShape__00000012_.FieldTypeAsString = ``
	__AttributeShape__00000012_.Structname = `AsSplitArea`
	__AttributeShape__00000012_.Fieldtypename = `bool`

	__AttributeShape__00000013_.Name = `DivStyle`
	__AttributeShape__00000013_.IdentifierMeta = ref_models.AsSplitArea{}.DivStyle
	__AttributeShape__00000013_.FieldTypeAsString = ``
	__AttributeShape__00000013_.Structname = `AsSplitArea`
	__AttributeShape__00000013_.Fieldtypename = `string`

	__AttributeShape__00000014_.Name = `NoName yet`
	__AttributeShape__00000014_.FieldTypeAsString = ``
	__AttributeShape__00000014_.Structname = ``
	__AttributeShape__00000014_.Fieldtypename = ``

	__Classdiagram__00000000_.Name = `Default`
	__Classdiagram__00000000_.Description = ``
	__Classdiagram__00000000_.IsIncludedInStaticWebSite = false
	__Classdiagram__00000000_.ShowNbInstances = true
	__Classdiagram__00000000_.ShowMultiplicity = true
	__Classdiagram__00000000_.ShowLinkNames = true
	__Classdiagram__00000000_.IsInRenameMode = false
	__Classdiagram__00000000_.IsExpanded = true
	__Classdiagram__00000000_.NodeGongStructsIsExpanded = false
	__Classdiagram__00000000_.NodeGongStructNodeExpansion = `[true,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,true]`
	__Classdiagram__00000000_.NodeGongEnumsIsExpanded = true
	__Classdiagram__00000000_.NodeGongEnumNodeExpansion = `[false]`
	__Classdiagram__00000000_.NodeGongNotesIsExpanded = false
	__Classdiagram__00000000_.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__00000000_.Name = `Diagram Package created the 2025-06-12T00:16:22Z`
	__DiagramPackage__00000000_.Path = ``
	__DiagramPackage__00000000_.GongModelPath = ``
	__DiagramPackage__00000000_.AbsolutePathToDiagramPackage = ``

	__GongEnumShape__00000000_.Name = `Default-Direction`
	__GongEnumShape__00000000_.X = 935.000000
	__GongEnumShape__00000000_.Y = 88.000000
	__GongEnumShape__00000000_.IdentifierMeta = new(ref_models.Direction)
	__GongEnumShape__00000000_.Width = 240.000000
	__GongEnumShape__00000000_.Height = 103.000000
	__GongEnumShape__00000000_.IsExpanded = false

	__GongEnumValueShape__00000000_.Name = `Horizontal`
	__GongEnumValueShape__00000000_.IdentifierMeta = ref_models.Horizontal

	__GongEnumValueShape__00000001_.Name = `Vertical`
	__GongEnumValueShape__00000001_.IdentifierMeta = ref_models.Vertical

	__GongStructShape__00000000_.Name = `Default-Split`
	__GongStructShape__00000000_.X = 579.000000
	__GongStructShape__00000000_.Y = 994.000000
	__GongStructShape__00000000_.IdentifierMeta = ref_models.Split{}
	__GongStructShape__00000000_.Width = 240.000000
	__GongStructShape__00000000_.Height = 83.000000
	__GongStructShape__00000000_.IsSelected = false

	__GongStructShape__00000001_.Name = `Default-Xlsx`
	__GongStructShape__00000001_.X = 579.000000
	__GongStructShape__00000001_.Y = 908.000000
	__GongStructShape__00000001_.IdentifierMeta = ref_models.Xlsx{}
	__GongStructShape__00000001_.Width = 240.000000
	__GongStructShape__00000001_.Height = 63.000000
	__GongStructShape__00000001_.IsSelected = false

	__GongStructShape__00000002_.Name = `Default-View`
	__GongStructShape__00000002_.X = 49.000000
	__GongStructShape__00000002_.Y = 155.000000
	__GongStructShape__00000002_.IdentifierMeta = ref_models.View{}
	__GongStructShape__00000002_.Width = 240.000000
	__GongStructShape__00000002_.Height = 123.000000
	__GongStructShape__00000002_.IsSelected = false

	__GongStructShape__00000003_.Name = `Default-Tree`
	__GongStructShape__00000003_.X = 581.000000
	__GongStructShape__00000003_.Y = 840.000000
	__GongStructShape__00000003_.IdentifierMeta = ref_models.Tree{}
	__GongStructShape__00000003_.Width = 240.000000
	__GongStructShape__00000003_.Height = 63.000000
	__GongStructShape__00000003_.IsSelected = false

	__GongStructShape__00000004_.Name = `Default-Tone`
	__GongStructShape__00000004_.X = 581.000000
	__GongStructShape__00000004_.Y = 749.000000
	__GongStructShape__00000004_.IdentifierMeta = ref_models.Tone{}
	__GongStructShape__00000004_.Width = 240.000000
	__GongStructShape__00000004_.Height = 63.000000
	__GongStructShape__00000004_.IsSelected = false

	__GongStructShape__00000005_.Name = `Default-Title`
	__GongStructShape__00000005_.X = 1275.000000
	__GongStructShape__00000005_.Y = 497.000000
	__GongStructShape__00000005_.IdentifierMeta = ref_models.Title{}
	__GongStructShape__00000005_.Width = 240.000000
	__GongStructShape__00000005_.Height = 63.000000
	__GongStructShape__00000005_.IsSelected = false

	__GongStructShape__00000006_.Name = `Default-Table`
	__GongStructShape__00000006_.X = 571.000000
	__GongStructShape__00000006_.Y = 489.000000
	__GongStructShape__00000006_.IdentifierMeta = ref_models.Table{}
	__GongStructShape__00000006_.Width = 240.000000
	__GongStructShape__00000006_.Height = 63.000000
	__GongStructShape__00000006_.IsSelected = false

	__GongStructShape__00000007_.Name = `Default-Svg`
	__GongStructShape__00000007_.X = 582.000000
	__GongStructShape__00000007_.Y = 667.000000
	__GongStructShape__00000007_.IdentifierMeta = ref_models.Svg{}
	__GongStructShape__00000007_.Width = 240.000000
	__GongStructShape__00000007_.Height = 63.000000
	__GongStructShape__00000007_.IsSelected = false

	__GongStructShape__00000008_.Name = `Default-Slider`
	__GongStructShape__00000008_.X = 557.000000
	__GongStructShape__00000008_.Y = 400.000000
	__GongStructShape__00000008_.IdentifierMeta = ref_models.Slider{}
	__GongStructShape__00000008_.Width = 240.000000
	__GongStructShape__00000008_.Height = 63.000000
	__GongStructShape__00000008_.IsSelected = false

	__GongStructShape__00000009_.Name = `Default-Markdown`
	__GongStructShape__00000009_.X = 576.000000
	__GongStructShape__00000009_.Y = 575.000000
	__GongStructShape__00000009_.IdentifierMeta = ref_models.Markdown{}
	__GongStructShape__00000009_.Width = 240.000000
	__GongStructShape__00000009_.Height = 63.000000
	__GongStructShape__00000009_.IsSelected = false

	__GongStructShape__00000010_.Name = `Default-LogoOnTheRight`
	__GongStructShape__00000010_.X = 1292.000000
	__GongStructShape__00000010_.Y = 229.000000
	__GongStructShape__00000010_.IdentifierMeta = ref_models.LogoOnTheRight{}
	__GongStructShape__00000010_.Width = 240.000000
	__GongStructShape__00000010_.Height = 143.000000
	__GongStructShape__00000010_.IsSelected = false

	__GongStructShape__00000011_.Name = `Default-LogoOnTheLeft`
	__GongStructShape__00000011_.X = 1306.000000
	__GongStructShape__00000011_.Y = 51.000000
	__GongStructShape__00000011_.IdentifierMeta = ref_models.LogoOnTheLeft{}
	__GongStructShape__00000011_.Width = 240.000000
	__GongStructShape__00000011_.Height = 143.000000
	__GongStructShape__00000011_.IsSelected = false

	__GongStructShape__00000012_.Name = `Default-Load`
	__GongStructShape__00000012_.X = 555.000000
	__GongStructShape__00000012_.Y = 308.000000
	__GongStructShape__00000012_.IdentifierMeta = ref_models.Load{}
	__GongStructShape__00000012_.Width = 240.000000
	__GongStructShape__00000012_.Height = 63.000000
	__GongStructShape__00000012_.IsSelected = false

	__GongStructShape__00000013_.Name = `Default-Form`
	__GongStructShape__00000013_.X = 562.000000
	__GongStructShape__00000013_.Y = 229.000000
	__GongStructShape__00000013_.IdentifierMeta = ref_models.Form{}
	__GongStructShape__00000013_.Width = 240.000000
	__GongStructShape__00000013_.Height = 63.000000
	__GongStructShape__00000013_.IsSelected = false

	__GongStructShape__00000014_.Name = `Default-FavIcon`
	__GongStructShape__00000014_.X = 1283.000000
	__GongStructShape__00000014_.Y = 402.000000
	__GongStructShape__00000014_.IdentifierMeta = ref_models.FavIcon{}
	__GongStructShape__00000014_.Width = 240.000000
	__GongStructShape__00000014_.Height = 63.000000
	__GongStructShape__00000014_.IsSelected = false

	__GongStructShape__00000015_.Name = `Default-Cursor`
	__GongStructShape__00000015_.X = 553.000000
	__GongStructShape__00000015_.Y = 123.000000
	__GongStructShape__00000015_.IdentifierMeta = ref_models.Cursor{}
	__GongStructShape__00000015_.Width = 240.000000
	__GongStructShape__00000015_.Height = 63.000000
	__GongStructShape__00000015_.IsSelected = false

	__GongStructShape__00000016_.Name = `Default-Button`
	__GongStructShape__00000016_.X = 551.000000
	__GongStructShape__00000016_.Y = 36.000000
	__GongStructShape__00000016_.IdentifierMeta = ref_models.Button{}
	__GongStructShape__00000016_.Width = 240.000000
	__GongStructShape__00000016_.Height = 63.000000
	__GongStructShape__00000016_.IsSelected = false

	__GongStructShape__00000017_.Name = `Default-AsSplitArea`
	__GongStructShape__00000017_.X = 35.000000
	__GongStructShape__00000017_.Y = 512.000000
	__GongStructShape__00000017_.IdentifierMeta = ref_models.AsSplitArea{}
	__GongStructShape__00000017_.Width = 240.000000
	__GongStructShape__00000017_.Height = 183.000000
	__GongStructShape__00000017_.IsSelected = false

	__GongStructShape__00000019_.Name = `Default-AsSplit`
	__GongStructShape__00000019_.X = 38.000000
	__GongStructShape__00000019_.Y = 830.000000
	__GongStructShape__00000019_.IdentifierMeta = ref_models.AsSplit{}
	__GongStructShape__00000019_.Width = 240.000000
	__GongStructShape__00000019_.Height = 83.000000
	__GongStructShape__00000019_.IsSelected = false

	__LinkShape__00000001_.Name = `AsSplit`
	__LinkShape__00000001_.IdentifierMeta = ref_models.AsSplitArea{}.AsSplit
	__LinkShape__00000001_.FieldTypeIdentifierMeta = ref_models.AsSplit{}
	__LinkShape__00000001_.FieldOffsetX = 0.000000
	__LinkShape__00000001_.FieldOffsetY = 0.000000
	__LinkShape__00000001_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000001_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000001_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000001_.SourceMultiplicity = models.MANY
	__LinkShape__00000001_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000001_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000001_.X = 631.500000
	__LinkShape__00000001_.Y = 436.500000
	__LinkShape__00000001_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000001_.StartRatio = 0.500000
	__LinkShape__00000001_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000001_.EndRatio = 0.500000
	__LinkShape__00000001_.CornerOffsetRatio = 1.380000

	__LinkShape__00000002_.Name = `Button`
	__LinkShape__00000002_.IdentifierMeta = ref_models.AsSplitArea{}.Button
	__LinkShape__00000002_.FieldTypeIdentifierMeta = ref_models.Button{}
	__LinkShape__00000002_.FieldOffsetX = 0.000000
	__LinkShape__00000002_.FieldOffsetY = 0.000000
	__LinkShape__00000002_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000002_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000002_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000002_.SourceMultiplicity = models.MANY
	__LinkShape__00000002_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000002_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000002_.X = 571.000000
	__LinkShape__00000002_.Y = 354.000000
	__LinkShape__00000002_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000002_.StartRatio = 0.500000
	__LinkShape__00000002_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000002_.EndRatio = 0.500000
	__LinkShape__00000002_.CornerOffsetRatio = 1.380000

	__LinkShape__00000003_.Name = `Cursor`
	__LinkShape__00000003_.IdentifierMeta = ref_models.AsSplitArea{}.Cursor
	__LinkShape__00000003_.FieldTypeIdentifierMeta = ref_models.Cursor{}
	__LinkShape__00000003_.FieldOffsetX = 0.000000
	__LinkShape__00000003_.FieldOffsetY = 0.000000
	__LinkShape__00000003_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000003_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000003_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000003_.SourceMultiplicity = models.MANY
	__LinkShape__00000003_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000003_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000003_.X = 672.500000
	__LinkShape__00000003_.Y = 381.000000
	__LinkShape__00000003_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000003_.StartRatio = 0.500000
	__LinkShape__00000003_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000003_.EndRatio = 0.500000
	__LinkShape__00000003_.CornerOffsetRatio = 1.380000

	__LinkShape__00000004_.Name = `Form`
	__LinkShape__00000004_.IdentifierMeta = ref_models.AsSplitArea{}.Form
	__LinkShape__00000004_.FieldTypeIdentifierMeta = ref_models.Form{}
	__LinkShape__00000004_.FieldOffsetX = 0.000000
	__LinkShape__00000004_.FieldOffsetY = 0.000000
	__LinkShape__00000004_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000004_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000004_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000004_.SourceMultiplicity = models.MANY
	__LinkShape__00000004_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000004_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000004_.X = 540.000000
	__LinkShape__00000004_.Y = 439.000000
	__LinkShape__00000004_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000004_.StartRatio = 0.500000
	__LinkShape__00000004_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000004_.EndRatio = 0.500000
	__LinkShape__00000004_.CornerOffsetRatio = 1.380000

	__LinkShape__00000005_.Name = `Load`
	__LinkShape__00000005_.IdentifierMeta = ref_models.AsSplitArea{}.Load
	__LinkShape__00000005_.FieldTypeIdentifierMeta = ref_models.Load{}
	__LinkShape__00000005_.FieldOffsetX = 0.000000
	__LinkShape__00000005_.FieldOffsetY = 0.000000
	__LinkShape__00000005_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000005_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000005_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000005_.SourceMultiplicity = models.MANY
	__LinkShape__00000005_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000005_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000005_.X = 434.000000
	__LinkShape__00000005_.Y = 444.500000
	__LinkShape__00000005_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000005_.StartRatio = 0.500000
	__LinkShape__00000005_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000005_.EndRatio = 0.500000
	__LinkShape__00000005_.CornerOffsetRatio = 1.380000

	__LinkShape__00000006_.Name = `Markdown`
	__LinkShape__00000006_.IdentifierMeta = ref_models.AsSplitArea{}.Markdown
	__LinkShape__00000006_.FieldTypeIdentifierMeta = ref_models.Markdown{}
	__LinkShape__00000006_.FieldOffsetX = 0.000000
	__LinkShape__00000006_.FieldOffsetY = 0.000000
	__LinkShape__00000006_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000006_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000006_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000006_.SourceMultiplicity = models.MANY
	__LinkShape__00000006_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000006_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000006_.X = 631.500000
	__LinkShape__00000006_.Y = 520.500000
	__LinkShape__00000006_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000006_.StartRatio = 0.500000
	__LinkShape__00000006_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000006_.EndRatio = 0.500000
	__LinkShape__00000006_.CornerOffsetRatio = 1.380000

	__LinkShape__00000007_.Name = `Slider`
	__LinkShape__00000007_.IdentifierMeta = ref_models.AsSplitArea{}.Slider
	__LinkShape__00000007_.FieldTypeIdentifierMeta = ref_models.Slider{}
	__LinkShape__00000007_.FieldOffsetX = 0.000000
	__LinkShape__00000007_.FieldOffsetY = 0.000000
	__LinkShape__00000007_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000007_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000007_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000007_.SourceMultiplicity = models.MANY
	__LinkShape__00000007_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000007_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000007_.X = 666.500000
	__LinkShape__00000007_.Y = 572.000000
	__LinkShape__00000007_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000007_.StartRatio = 0.500000
	__LinkShape__00000007_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000007_.EndRatio = 0.500000
	__LinkShape__00000007_.CornerOffsetRatio = 1.380000

	__LinkShape__00000008_.Name = `Split`
	__LinkShape__00000008_.IdentifierMeta = ref_models.AsSplitArea{}.Split
	__LinkShape__00000008_.FieldTypeIdentifierMeta = ref_models.Split{}
	__LinkShape__00000008_.FieldOffsetX = 0.000000
	__LinkShape__00000008_.FieldOffsetY = 0.000000
	__LinkShape__00000008_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000008_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000008_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000008_.SourceMultiplicity = models.MANY
	__LinkShape__00000008_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000008_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000008_.X = 390.500000
	__LinkShape__00000008_.Y = 351.500000
	__LinkShape__00000008_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000008_.StartRatio = 0.500000
	__LinkShape__00000008_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000008_.EndRatio = 0.500000
	__LinkShape__00000008_.CornerOffsetRatio = 1.380000

	__LinkShape__00000009_.Name = `Svg`
	__LinkShape__00000009_.IdentifierMeta = ref_models.AsSplitArea{}.Svg
	__LinkShape__00000009_.FieldTypeIdentifierMeta = ref_models.Svg{}
	__LinkShape__00000009_.FieldOffsetX = 0.000000
	__LinkShape__00000009_.FieldOffsetY = 0.000000
	__LinkShape__00000009_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000009_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000009_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000009_.SourceMultiplicity = models.MANY
	__LinkShape__00000009_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000009_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000009_.X = 805.500000
	__LinkShape__00000009_.Y = 519.000000
	__LinkShape__00000009_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000009_.StartRatio = 0.500000
	__LinkShape__00000009_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000009_.EndRatio = 0.500000
	__LinkShape__00000009_.CornerOffsetRatio = 1.380000

	__LinkShape__00000010_.Name = `Table`
	__LinkShape__00000010_.IdentifierMeta = ref_models.AsSplitArea{}.Table
	__LinkShape__00000010_.FieldTypeIdentifierMeta = ref_models.Table{}
	__LinkShape__00000010_.FieldOffsetX = 0.000000
	__LinkShape__00000010_.FieldOffsetY = 0.000000
	__LinkShape__00000010_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000010_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000010_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000010_.SourceMultiplicity = models.MANY
	__LinkShape__00000010_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000010_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000010_.X = 695.000000
	__LinkShape__00000010_.Y = 520.500000
	__LinkShape__00000010_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000010_.StartRatio = 0.500000
	__LinkShape__00000010_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000010_.EndRatio = 0.500000
	__LinkShape__00000010_.CornerOffsetRatio = 1.380000

	__LinkShape__00000011_.Name = `Tone`
	__LinkShape__00000011_.IdentifierMeta = ref_models.AsSplitArea{}.Tone
	__LinkShape__00000011_.FieldTypeIdentifierMeta = ref_models.Tone{}
	__LinkShape__00000011_.FieldOffsetX = 0.000000
	__LinkShape__00000011_.FieldOffsetY = 0.000000
	__LinkShape__00000011_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000011_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000011_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000011_.SourceMultiplicity = models.MANY
	__LinkShape__00000011_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000011_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000011_.X = 853.500000
	__LinkShape__00000011_.Y = 456.500000
	__LinkShape__00000011_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000011_.StartRatio = 0.500000
	__LinkShape__00000011_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000011_.EndRatio = 0.500000
	__LinkShape__00000011_.CornerOffsetRatio = 1.380000

	__LinkShape__00000012_.Name = `Tree`
	__LinkShape__00000012_.IdentifierMeta = ref_models.AsSplitArea{}.Tree
	__LinkShape__00000012_.FieldTypeIdentifierMeta = ref_models.Tree{}
	__LinkShape__00000012_.FieldOffsetX = 0.000000
	__LinkShape__00000012_.FieldOffsetY = 0.000000
	__LinkShape__00000012_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000012_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000012_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000012_.SourceMultiplicity = models.MANY
	__LinkShape__00000012_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000012_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000012_.X = 894.000000
	__LinkShape__00000012_.Y = 562.000000
	__LinkShape__00000012_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000012_.StartRatio = 0.500000
	__LinkShape__00000012_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000012_.EndRatio = 0.500000
	__LinkShape__00000012_.CornerOffsetRatio = 1.380000

	__LinkShape__00000013_.Name = `Xlsx`
	__LinkShape__00000013_.IdentifierMeta = ref_models.AsSplitArea{}.Xlsx
	__LinkShape__00000013_.FieldTypeIdentifierMeta = ref_models.Xlsx{}
	__LinkShape__00000013_.FieldOffsetX = 0.000000
	__LinkShape__00000013_.FieldOffsetY = 0.000000
	__LinkShape__00000013_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000013_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000013_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000013_.SourceMultiplicity = models.MANY
	__LinkShape__00000013_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000013_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000013_.X = 848.500000
	__LinkShape__00000013_.Y = 656.000000
	__LinkShape__00000013_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000013_.StartRatio = 0.498709
	__LinkShape__00000013_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000013_.EndRatio = 0.500000
	__LinkShape__00000013_.CornerOffsetRatio = 1.382324

	__LinkShape__00000014_.Name = `NoName yet`
	__LinkShape__00000014_.FieldOffsetX = 0.000000
	__LinkShape__00000014_.FieldOffsetY = 0.000000
	__LinkShape__00000014_.TargetMultiplicity = ""
	__LinkShape__00000014_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000014_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000014_.SourceMultiplicity = ""
	__LinkShape__00000014_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000014_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000014_.X = 0.000000
	__LinkShape__00000014_.Y = 0.000000
	__LinkShape__00000014_.StartOrientation = ""
	__LinkShape__00000014_.StartRatio = 0.000000
	__LinkShape__00000014_.EndOrientation = ""
	__LinkShape__00000014_.EndRatio = 0.000000
	__LinkShape__00000014_.CornerOffsetRatio = 0.000000

	__LinkShape__00000015_.Name = `AsSplitAreas`
	__LinkShape__00000015_.IdentifierMeta = ref_models.AsSplit{}.AsSplitAreas
	__LinkShape__00000015_.FieldTypeIdentifierMeta = ref_models.AsSplitArea{}
	__LinkShape__00000015_.FieldOffsetX = 0.000000
	__LinkShape__00000015_.FieldOffsetY = 0.000000
	__LinkShape__00000015_.TargetMultiplicity = models.MANY
	__LinkShape__00000015_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000015_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000015_.SourceMultiplicity = models.MANY
	__LinkShape__00000015_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000015_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000015_.X = 396.500000
	__LinkShape__00000015_.Y = 712.500000
	__LinkShape__00000015_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000015_.StartRatio = 0.786491
	__LinkShape__00000015_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000015_.EndRatio = 0.798991
	__LinkShape__00000015_.CornerOffsetRatio = -0.623329

	__LinkShape__00000016_.Name = `RootAsSplitAreas`
	__LinkShape__00000016_.IdentifierMeta = ref_models.View{}.RootAsSplitAreas
	__LinkShape__00000016_.FieldTypeIdentifierMeta = ref_models.AsSplitArea{}
	__LinkShape__00000016_.FieldOffsetX = 0.000000
	__LinkShape__00000016_.FieldOffsetY = 0.000000
	__LinkShape__00000016_.TargetMultiplicity = models.MANY
	__LinkShape__00000016_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000016_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000016_.SourceMultiplicity = models.MANY
	__LinkShape__00000016_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000016_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000016_.X = 402.000000
	__LinkShape__00000016_.Y = 395.000000
	__LinkShape__00000016_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000016_.StartRatio = 0.757324
	__LinkShape__00000016_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000016_.EndRatio = 0.844824
	__LinkShape__00000016_.CornerOffsetRatio = 2.026534

	// insertion point for setup of pointers
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000000_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000001_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000002_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000003_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000004_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000005_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000006_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000007_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000008_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000009_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000010_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000011_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000012_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000013_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000014_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000015_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000016_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000017_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000019_)
	__Classdiagram__00000000_.GongEnumShapes = append(__Classdiagram__00000000_.GongEnumShapes, __GongEnumShape__00000000_)
	__DiagramPackage__00000000_.Classdiagrams = append(__DiagramPackage__00000000_.Classdiagrams, __Classdiagram__00000000_)
	__DiagramPackage__00000000_.SelectedClassdiagram = __Classdiagram__00000000_
	__GongEnumShape__00000000_.GongEnumValueShapes = append(__GongEnumShape__00000000_.GongEnumValueShapes, __GongEnumValueShape__00000001_)
	__GongEnumShape__00000000_.GongEnumValueShapes = append(__GongEnumShape__00000000_.GongEnumValueShapes, __GongEnumValueShape__00000000_)
	__GongStructShape__00000000_.AttributeShapes = append(__GongStructShape__00000000_.AttributeShapes, __AttributeShape__00000000_)
	__GongStructShape__00000002_.AttributeShapes = append(__GongStructShape__00000002_.AttributeShapes, __AttributeShape__00000002_)
	__GongStructShape__00000002_.AttributeShapes = append(__GongStructShape__00000002_.AttributeShapes, __AttributeShape__00000001_)
	__GongStructShape__00000002_.LinkShapes = append(__GongStructShape__00000002_.LinkShapes, __LinkShape__00000016_)
	__GongStructShape__00000010_.AttributeShapes = append(__GongStructShape__00000010_.AttributeShapes, __AttributeShape__00000005_)
	__GongStructShape__00000010_.AttributeShapes = append(__GongStructShape__00000010_.AttributeShapes, __AttributeShape__00000004_)
	__GongStructShape__00000010_.AttributeShapes = append(__GongStructShape__00000010_.AttributeShapes, __AttributeShape__00000003_)
	__GongStructShape__00000011_.AttributeShapes = append(__GongStructShape__00000011_.AttributeShapes, __AttributeShape__00000006_)
	__GongStructShape__00000011_.AttributeShapes = append(__GongStructShape__00000011_.AttributeShapes, __AttributeShape__00000007_)
	__GongStructShape__00000011_.AttributeShapes = append(__GongStructShape__00000011_.AttributeShapes, __AttributeShape__00000008_)
	__GongStructShape__00000017_.AttributeShapes = append(__GongStructShape__00000017_.AttributeShapes, __AttributeShape__00000009_)
	__GongStructShape__00000017_.AttributeShapes = append(__GongStructShape__00000017_.AttributeShapes, __AttributeShape__00000010_)
	__GongStructShape__00000017_.AttributeShapes = append(__GongStructShape__00000017_.AttributeShapes, __AttributeShape__00000011_)
	__GongStructShape__00000017_.AttributeShapes = append(__GongStructShape__00000017_.AttributeShapes, __AttributeShape__00000012_)
	__GongStructShape__00000017_.AttributeShapes = append(__GongStructShape__00000017_.AttributeShapes, __AttributeShape__00000013_)
	__GongStructShape__00000017_.LinkShapes = append(__GongStructShape__00000017_.LinkShapes, __LinkShape__00000001_)
	__GongStructShape__00000017_.LinkShapes = append(__GongStructShape__00000017_.LinkShapes, __LinkShape__00000002_)
	__GongStructShape__00000017_.LinkShapes = append(__GongStructShape__00000017_.LinkShapes, __LinkShape__00000003_)
	__GongStructShape__00000017_.LinkShapes = append(__GongStructShape__00000017_.LinkShapes, __LinkShape__00000004_)
	__GongStructShape__00000017_.LinkShapes = append(__GongStructShape__00000017_.LinkShapes, __LinkShape__00000005_)
	__GongStructShape__00000017_.LinkShapes = append(__GongStructShape__00000017_.LinkShapes, __LinkShape__00000006_)
	__GongStructShape__00000017_.LinkShapes = append(__GongStructShape__00000017_.LinkShapes, __LinkShape__00000007_)
	__GongStructShape__00000017_.LinkShapes = append(__GongStructShape__00000017_.LinkShapes, __LinkShape__00000008_)
	__GongStructShape__00000017_.LinkShapes = append(__GongStructShape__00000017_.LinkShapes, __LinkShape__00000009_)
	__GongStructShape__00000017_.LinkShapes = append(__GongStructShape__00000017_.LinkShapes, __LinkShape__00000010_)
	__GongStructShape__00000017_.LinkShapes = append(__GongStructShape__00000017_.LinkShapes, __LinkShape__00000011_)
	__GongStructShape__00000017_.LinkShapes = append(__GongStructShape__00000017_.LinkShapes, __LinkShape__00000012_)
	__GongStructShape__00000017_.LinkShapes = append(__GongStructShape__00000017_.LinkShapes, __LinkShape__00000013_)
	__GongStructShape__00000019_.LinkShapes = append(__GongStructShape__00000019_.LinkShapes, __LinkShape__00000015_)
}
