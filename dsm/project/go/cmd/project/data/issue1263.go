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

	// insertion point for declaration of instances to stage

	__Diagram__00000000_ := (&models.Diagram{Name: `Default Diagram`}).Stage(stage)
	__Diagram__00000001_ := (&models.Diagram{Name: `Default Diagram`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: `Pain points`}).Stage(stage)
	__Library__00000001_ := (&models.Library{Name: `Goals`}).Stage(stage)

	__Note__00000000_ := (&models.Note{Name: `Invariant : generated code never import a gong package `}).Stage(stage)
	__Note__00000001_ := (&models.Note{Name: `A typical DSM is "concept map", a tool to creates Noun, Verbs and directed relations between them. The visual diagram is very simple. `}).Stage(stage)
	__Note__00000002_ := (&models.Note{Name: `A DSM is 
 - abstract models
- concrete concepts
- library organisation with the root object
- semantic rules
- a Tree
- a SVG

`}).Stage(stage)
	__Note__00000003_ := (&models.Note{Name: `Default functions like Rename, Delete, ...`}).Stage(stage)
	__Note__00000004_ := (&models.Note{Name: `sqfsdqfqd`}).Stage(stage)
	__Note__00000005_ := (&models.Note{Name: `For instance, a project DSM needs a Requirement DSM and a Statemachine DSM `}).Stage(stage)

	__NoteProductShape__00000000_ := (&models.NoteProductShape{Name: `Default Diagram-Invariant : generated code never import a gong package -confusion between interface & generic code`}).Stage(stage)

	__NoteShape__00000000_ := (&models.NoteShape{Name: `Default Diagram-Invariant : generated code never import a gong package `}).Stage(stage)
	__NoteShape__00000001_ := (&models.NoteShape{Name: `Default Diagram-A typical DSM is "concept map", a tool to creates Noun, Verbs and directed relations between them. The visual diagram is very simple. `}).Stage(stage)
	__NoteShape__00000002_ := (&models.NoteShape{Name: `Default Diagram-A DSM is 
 - abstract models
- concrete concepts
- library organisation with the root object
- semantic rules
- a Tree
- a SVG

`}).Stage(stage)
	__NoteShape__00000003_ := (&models.NoteShape{Name: `Default Diagram-Default functions like Rename, Delete, ...`}).Stage(stage)
	__NoteShape__00000004_ := (&models.NoteShape{Name: `Default Diagram-sqfsdqfqd`}).Stage(stage)
	__NoteShape__00000007_ := (&models.NoteShape{Name: `Default Diagram-For instance, a project DSM needs a Requirement DSM and a Statemachine DSM `}).Stage(stage)

	__Product__00000000_ := (&models.Product{Name: `Paint points`}).Stage(stage)
	__Product__00000001_ := (&models.Product{Name: `lot of yyy files and no clear mental model `}).Stage(stage)
	__Product__00000002_ := (&models.Product{Name: `lot of stager is a hybrid concept`}).Stage(stage)
	__Product__00000003_ := (&models.Product{Name: `the stager concept is not clearly delineated`}).Stage(stage)
	__Product__00000004_ := (&models.Product{Name: `confusion between interface & generic code`}).Stage(stage)
	__Product__00000005_ := (&models.Product{Name: `There is a lifecycle "enforce sematntic", "marhsall", generate UXs`}).Stage(stage)
	__Product__00000006_ := (&models.Product{Name: `TBC`}).Stage(stage)
	__Product__00000007_ := (&models.Product{Name: `yyy files are helpers for UX updates of trees and svg`}).Stage(stage)
	__Product__00000008_ := (&models.Product{Name: `Core Abstractions & Interfaces`}).Stage(stage)
	__Product__00000009_ := (&models.Product{Name: `Tree UI & Context Menus`}).Stage(stage)
	__Product__00000010_ := (&models.Product{Name: `UI Callbacks & State Syncing`}).Stage(stage)
	__Product__00000011_ := (&models.Product{Name: `SVG Diagram Rendering & Interaction`}).Stage(stage)
	__Product__00000012_ := (&models.Product{Name: `Graph Data Integrity & Traversal`}).Stage(stage)
	__Product__00000013_ := (&models.Product{Name: `Rely on heavily typed Stage`}).Stage(stage)
	__Product__00000014_ := (&models.Product{Name: `lots of boilerplate code`}).Stage(stage)
	__Product__00000015_ := (&models.Product{Name: `Goals`}).Stage(stage)
	__Product__00000016_ := (&models.Product{Name: `creates a DSM in les than 1/2 hour`}).Stage(stage)
	__Product__00000017_ := (&models.Product{Name: `Provides cleans abstraction for the development of the tree`}).Stage(stage)
	__Product__00000018_ := (&models.Product{Name: `Abstraction for semantic rules`}).Stage(stage)
	__Product__00000023_ := (&models.Product{Name: `Allows combination of DSM `}).Stage(stage)
	__Product__00000024_ := (&models.Product{Name: ``}).Stage(stage)
	__Product__00000025_ := (&models.Product{Name: ``}).Stage(stage)
	__Product__00000026_ := (&models.Product{Name: `sss`}).Stage(stage)

	__ProductCompositionShape__00000000_ := (&models.ProductCompositionShape{Name: `Default Diagram-Paint points-lot of yyy files and no clear mental model `}).Stage(stage)
	__ProductCompositionShape__00000001_ := (&models.ProductCompositionShape{Name: `Default Diagram-Paint points-lot of stager is a hybrid concept`}).Stage(stage)
	__ProductCompositionShape__00000002_ := (&models.ProductCompositionShape{Name: `Default Diagram-Paint points-the stager concept is not clearly delineated`}).Stage(stage)
	__ProductCompositionShape__00000003_ := (&models.ProductCompositionShape{Name: `Default Diagram-Paint points-confusion between interface & generic code`}).Stage(stage)
	__ProductCompositionShape__00000004_ := (&models.ProductCompositionShape{Name: `Default Diagram-the stager concept is not clearly delineated-There is a lifecycle "enforce sematntic", "marhsall", generate UXs`}).Stage(stage)
	__ProductCompositionShape__00000005_ := (&models.ProductCompositionShape{Name: `Default Diagram-the stager concept is not clearly delineated-TBC`}).Stage(stage)
	__ProductCompositionShape__00000006_ := (&models.ProductCompositionShape{Name: `Default Diagram-lot of yyy files and no clear mental model -yyy files are helpers for UX updates of trees and svg`}).Stage(stage)
	__ProductCompositionShape__00000007_ := (&models.ProductCompositionShape{Name: `Default Diagram-yyy files are helpers for UX updates of trees and svg-Core Abstractions & Interfaces`}).Stage(stage)
	__ProductCompositionShape__00000008_ := (&models.ProductCompositionShape{Name: `Default Diagram-yyy files are helpers for UX updates of trees and svg-Tree UI & Context Menus`}).Stage(stage)
	__ProductCompositionShape__00000009_ := (&models.ProductCompositionShape{Name: `Default Diagram-yyy files are helpers for UX updates of trees and svg-UI Callbacks & State Syncing`}).Stage(stage)
	__ProductCompositionShape__00000010_ := (&models.ProductCompositionShape{Name: `Default Diagram-yyy files are helpers for UX updates of trees and svg-SVG Diagram Rendering & Interaction`}).Stage(stage)
	__ProductCompositionShape__00000011_ := (&models.ProductCompositionShape{Name: `Default Diagram-yyy files are helpers for UX updates of trees and svg-Graph Data Integrity & Traversal`}).Stage(stage)
	__ProductCompositionShape__00000012_ := (&models.ProductCompositionShape{Name: `Default Diagram-lot of stager is a hybrid concept-Rely on heavily typed Stage`}).Stage(stage)
	__ProductCompositionShape__00000013_ := (&models.ProductCompositionShape{Name: `Default Diagram-lot of stager is a hybrid concept-lots of boilerplate code`}).Stage(stage)
	__ProductCompositionShape__00000014_ := (&models.ProductCompositionShape{Name: `Default Diagram-Goals-creates a DSM in les than 1/2 hour`}).Stage(stage)
	__ProductCompositionShape__00000015_ := (&models.ProductCompositionShape{Name: `Default Diagram-creates a DSM in les than 1/2 hour-Provides cleans abstraction for the development of the tree`}).Stage(stage)
	__ProductCompositionShape__00000017_ := (&models.ProductCompositionShape{Name: `Default Diagram-creates a DSM in les than 1/2 hour-Abstraction for semantic rules`}).Stage(stage)
	__ProductCompositionShape__00000020_ := (&models.ProductCompositionShape{Name: `Default Diagram-Goals-Allows combination of DSM `}).Stage(stage)

	__ProductShape__00000000_ := (&models.ProductShape{Name: `Default Diagram-Paint points`}).Stage(stage)
	__ProductShape__00000001_ := (&models.ProductShape{Name: `Default Diagram-lot of yyy files and no clear mental model `}).Stage(stage)
	__ProductShape__00000002_ := (&models.ProductShape{Name: `Default Diagram-lot of stager is a hybrid concept`}).Stage(stage)
	__ProductShape__00000003_ := (&models.ProductShape{Name: `Default Diagram-the stager concept is not clearly delineated`}).Stage(stage)
	__ProductShape__00000004_ := (&models.ProductShape{Name: `Default Diagram-confusion between interface & generic code`}).Stage(stage)
	__ProductShape__00000005_ := (&models.ProductShape{Name: `Default Diagram-There is a lifecycle "enforce sematntic", "marhsall", generate UXs`}).Stage(stage)
	__ProductShape__00000006_ := (&models.ProductShape{Name: `Default Diagram-TBC`}).Stage(stage)
	__ProductShape__00000007_ := (&models.ProductShape{Name: `Default Diagram-yyy files are helpers for UX updates of trees and svg`}).Stage(stage)
	__ProductShape__00000008_ := (&models.ProductShape{Name: `Default Diagram-Core Abstractions & Interfaces`}).Stage(stage)
	__ProductShape__00000009_ := (&models.ProductShape{Name: `Default Diagram-Tree UI & Context Menus`}).Stage(stage)
	__ProductShape__00000010_ := (&models.ProductShape{Name: `Default Diagram-UI Callbacks & State Syncing`}).Stage(stage)
	__ProductShape__00000011_ := (&models.ProductShape{Name: `Default Diagram-SVG Diagram Rendering & Interaction`}).Stage(stage)
	__ProductShape__00000012_ := (&models.ProductShape{Name: `Default Diagram-Graph Data Integrity & Traversal`}).Stage(stage)
	__ProductShape__00000013_ := (&models.ProductShape{Name: `Default Diagram-Rely on heavily typed Stage`}).Stage(stage)
	__ProductShape__00000014_ := (&models.ProductShape{Name: `Default Diagram-lots of boilerplate code`}).Stage(stage)
	__ProductShape__00000015_ := (&models.ProductShape{Name: `Default Diagram-Goals`}).Stage(stage)
	__ProductShape__00000016_ := (&models.ProductShape{Name: `Default Diagram-creates a DSM in les than 1/2 hour`}).Stage(stage)
	__ProductShape__00000017_ := (&models.ProductShape{Name: `Default Diagram-Provides cleans abstraction for the development of the tree`}).Stage(stage)
	__ProductShape__00000018_ := (&models.ProductShape{Name: `Default Diagram-Abstraction for semantic rules`}).Stage(stage)
	__ProductShape__00000025_ := (&models.ProductShape{Name: `Default Diagram-Allows combination of DSM `}).Stage(stage)
	__ProductShape__00000028_ := (&models.ProductShape{Name: `Default Diagram-sss`}).Stage(stage)

	// insertion point for initialization of values

	__Diagram__00000000_.Name = `Default Diagram`
	__Diagram__00000000_.DefaultBoxWidth = 250.000000
	__Diagram__00000000_.DefaultBoxHeigth = 70.000000
	__Diagram__00000000_.DateFormat = ``
	__Diagram__00000000_.Width = 1269.706973
	__Diagram__00000000_.Height = 1405.000000
	__Diagram__00000000_.IsTimeDiagram = false
	__Diagram__00000000_.ComputedStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000000_.ComputedEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000000_.ComputedDuration = 0
	__Diagram__00000000_.UseManualStartAndEndDates = false
	__Diagram__00000000_.ManualStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000000_.ManualEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000000_.TimeStep = 0
	__Diagram__00000000_.TimeStepScale = ""
	__Diagram__00000000_.LaneHeight = 0.000000
	__Diagram__00000000_.RatioBarToLaneHeight = 0.000000
	__Diagram__00000000_.YTopMargin = 0.000000
	__Diagram__00000000_.XLeftText = 0.000000
	__Diagram__00000000_.TextHeight = 0.000000
	__Diagram__00000000_.XLeftLanes = 0.000000
	__Diagram__00000000_.XRightMargin = 0.000000
	__Diagram__00000000_.ArrowLengthToTheRightOfStartBar = 0.000000
	__Diagram__00000000_.ArrowTipLenght = 0.000000
	__Diagram__00000000_.TimeLine_Color = ``
	__Diagram__00000000_.TimeLine_FillOpacity = 0.000000
	__Diagram__00000000_.TimeLine_Stroke = ``
	__Diagram__00000000_.TimeLine_StrokeWidth = 0.000000
	__Diagram__00000000_.DrawVerticalTimeLines = false
	__Diagram__00000000_.Group_Stroke = ``
	__Diagram__00000000_.Group_StrokeWidth = 0.000000
	__Diagram__00000000_.Group_StrokeDashArray = ``
	__Diagram__00000000_.DateYOffset = 0.000000
	__Diagram__00000000_.AlignOnStartEndOnYearStart = false
	__Diagram__00000000_.ComputedPrefix = `1`
	__Diagram__00000000_.IsExpanded = true
	__Diagram__00000000_.LayoutDirection = models.Vertical
	__Diagram__00000000_.IsChecked = false
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.IsShowPrefix = false
	__Diagram__00000000_.IsPBSNodeExpanded = true
	__Diagram__00000000_.IsWBSNodeExpanded = true
	__Diagram__00000000_.IsTaskGroupsNodeExpanded = false
	__Diagram__00000000_.IsNotesNodeExpanded = true
	__Diagram__00000000_.IsResourcesNodeExpanded = false

	__Diagram__00000001_.Name = `Default Diagram`
	__Diagram__00000001_.DefaultBoxWidth = 250.000000
	__Diagram__00000001_.DefaultBoxHeigth = 70.000000
	__Diagram__00000001_.DateFormat = ``
	__Diagram__00000001_.Width = 1335.155913
	__Diagram__00000001_.Height = 1514.000000
	__Diagram__00000001_.IsTimeDiagram = false
	__Diagram__00000001_.ComputedStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000001_.ComputedEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000001_.ComputedDuration = 0
	__Diagram__00000001_.UseManualStartAndEndDates = false
	__Diagram__00000001_.ManualStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000001_.ManualEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000001_.TimeStep = 0
	__Diagram__00000001_.TimeStepScale = ""
	__Diagram__00000001_.LaneHeight = 0.000000
	__Diagram__00000001_.RatioBarToLaneHeight = 0.000000
	__Diagram__00000001_.YTopMargin = 0.000000
	__Diagram__00000001_.XLeftText = 0.000000
	__Diagram__00000001_.TextHeight = 0.000000
	__Diagram__00000001_.XLeftLanes = 0.000000
	__Diagram__00000001_.XRightMargin = 0.000000
	__Diagram__00000001_.ArrowLengthToTheRightOfStartBar = 0.000000
	__Diagram__00000001_.ArrowTipLenght = 0.000000
	__Diagram__00000001_.TimeLine_Color = ``
	__Diagram__00000001_.TimeLine_FillOpacity = 0.000000
	__Diagram__00000001_.TimeLine_Stroke = ``
	__Diagram__00000001_.TimeLine_StrokeWidth = 0.000000
	__Diagram__00000001_.DrawVerticalTimeLines = false
	__Diagram__00000001_.Group_Stroke = ``
	__Diagram__00000001_.Group_StrokeWidth = 0.000000
	__Diagram__00000001_.Group_StrokeDashArray = ``
	__Diagram__00000001_.DateYOffset = 0.000000
	__Diagram__00000001_.AlignOnStartEndOnYearStart = false
	__Diagram__00000001_.ComputedPrefix = `1`
	__Diagram__00000001_.IsExpanded = true
	__Diagram__00000001_.LayoutDirection = models.Vertical
	__Diagram__00000001_.IsChecked = true
	__Diagram__00000001_.IsEditable_ = true
	__Diagram__00000001_.IsShowPrefix = false
	__Diagram__00000001_.IsPBSNodeExpanded = true
	__Diagram__00000001_.IsWBSNodeExpanded = true
	__Diagram__00000001_.IsTaskGroupsNodeExpanded = false
	__Diagram__00000001_.IsNotesNodeExpanded = true
	__Diagram__00000001_.IsResourcesNodeExpanded = false

	__Library__00000000_.Name = `Pain points`
	__Library__00000000_.NbPixPerCharacter = 8.000000
	__Library__00000000_.LogoSVGFile = ``
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsExpanded = true
	__Library__00000000_.LayoutDirection = models.Vertical
	__Library__00000000_.IsRootLibrary = true

	__Library__00000001_.Name = `Goals`
	__Library__00000001_.NbPixPerCharacter = 0.000000
	__Library__00000001_.LogoSVGFile = ``
	__Library__00000001_.ComputedPrefix = ``
	__Library__00000001_.IsExpanded = true
	__Library__00000001_.LayoutDirection = models.Vertical
	__Library__00000001_.IsRootLibrary = false

	__Note__00000000_.Name = `Invariant : generated code never import a gong package `
	__Note__00000000_.ComputedPrefix = `1`
	__Note__00000000_.IsExpanded = false
	__Note__00000000_.LayoutDirection = models.Vertical

	__Note__00000001_.Name = `A typical DSM is "concept map", a tool to creates Noun, Verbs and directed relations between them. The visual diagram is very simple. `
	__Note__00000001_.ComputedPrefix = `1`
	__Note__00000001_.IsExpanded = false
	__Note__00000001_.LayoutDirection = models.Vertical

	__Note__00000002_.Name = `A DSM is 
 - abstract models
- concrete concepts
- library organisation with the root object
- semantic rules
- a Tree
- a SVG

`
	__Note__00000002_.ComputedPrefix = `2`
	__Note__00000002_.IsExpanded = false
	__Note__00000002_.LayoutDirection = models.Vertical

	__Note__00000003_.Name = `Default functions like Rename, Delete, ...`
	__Note__00000003_.ComputedPrefix = `3`
	__Note__00000003_.IsExpanded = false
	__Note__00000003_.LayoutDirection = models.Vertical

	__Note__00000004_.Name = `sqfsdqfqd`
	__Note__00000004_.ComputedPrefix = `4`
	__Note__00000004_.IsExpanded = false
	__Note__00000004_.LayoutDirection = models.Vertical

	__Note__00000005_.Name = `For instance, a project DSM needs a Requirement DSM and a Statemachine DSM `
	__Note__00000005_.ComputedPrefix = `5`
	__Note__00000005_.IsExpanded = false
	__Note__00000005_.LayoutDirection = models.Vertical

	__NoteProductShape__00000000_.Name = `Default Diagram-Invariant : generated code never import a gong package -confusion between interface & generic code`
	__NoteProductShape__00000000_.StartRatio = 0.500000
	__NoteProductShape__00000000_.EndRatio = 0.500000
	__NoteProductShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000000_.CornerOffsetRatio = 1.680000
	__NoteProductShape__00000000_.IsHidden = false

	__NoteShape__00000000_.Name = `Default Diagram-Invariant : generated code never import a gong package `
	__NoteShape__00000000_.OverideLayoutDirection = false
	__NoteShape__00000000_.LayoutDirection = models.Vertical
	__NoteShape__00000000_.X = 919.706973
	__NoteShape__00000000_.Y = 1125.285144
	__NoteShape__00000000_.Width = 250.000000
	__NoteShape__00000000_.Height = 70.000000
	__NoteShape__00000000_.IsHidden = false

	__NoteShape__00000001_.Name = `Default Diagram-A typical DSM is "concept map", a tool to creates Noun, Verbs and directed relations between them. The visual diagram is very simple. `
	__NoteShape__00000001_.OverideLayoutDirection = false
	__NoteShape__00000001_.LayoutDirection = models.Vertical
	__NoteShape__00000001_.X = 522.625303
	__NoteShape__00000001_.Y = 472.810898
	__NoteShape__00000001_.Width = 250.000000
	__NoteShape__00000001_.Height = 122.000000
	__NoteShape__00000001_.IsHidden = false

	__NoteShape__00000002_.Name = `Default Diagram-A DSM is 
 - abstract models
- concrete concepts
- library organisation with the root object
- semantic rules
- a Tree
- a SVG

`
	__NoteShape__00000002_.OverideLayoutDirection = false
	__NoteShape__00000002_.LayoutDirection = models.Vertical
	__NoteShape__00000002_.X = 985.155913
	__NoteShape__00000002_.Y = 37.256246
	__NoteShape__00000002_.Width = 250.000000
	__NoteShape__00000002_.Height = 163.000000
	__NoteShape__00000002_.IsHidden = false

	__NoteShape__00000003_.Name = `Default Diagram-Default functions like Rename, Delete, ...`
	__NoteShape__00000003_.OverideLayoutDirection = false
	__NoteShape__00000003_.LayoutDirection = models.Vertical
	__NoteShape__00000003_.X = 950.184253
	__NoteShape__00000003_.Y = 479.000000
	__NoteShape__00000003_.Width = 250.000000
	__NoteShape__00000003_.Height = 70.000000
	__NoteShape__00000003_.IsHidden = false

	__NoteShape__00000004_.Name = `Default Diagram-sqfsdqfqd`
	__NoteShape__00000004_.OverideLayoutDirection = false
	__NoteShape__00000004_.LayoutDirection = models.Vertical
	__NoteShape__00000004_.X = 58.558884
	__NoteShape__00000004_.Y = 1344.000000
	__NoteShape__00000004_.Width = 250.000000
	__NoteShape__00000004_.Height = 70.000000
	__NoteShape__00000004_.IsHidden = false

	__NoteShape__00000007_.Name = `Default Diagram-For instance, a project DSM needs a Requirement DSM and a Statemachine DSM `
	__NoteShape__00000007_.OverideLayoutDirection = false
	__NoteShape__00000007_.LayoutDirection = models.Vertical
	__NoteShape__00000007_.X = 600.450625
	__NoteShape__00000007_.Y = 621.387352
	__NoteShape__00000007_.Width = 250.000000
	__NoteShape__00000007_.Height = 70.000000
	__NoteShape__00000007_.IsHidden = false

	__Product__00000000_.Name = `Paint points`
	__Product__00000000_.Description = ``
	__Product__00000000_.IsProducersNodeExpanded = false
	__Product__00000000_.IsConsumersNodeExpanded = false
	__Product__00000000_.IsImport = false
	__Product__00000000_.ComputedPrefix = `1`
	__Product__00000000_.IsExpanded = false
	__Product__00000000_.LayoutDirection = models.Horizontal

	__Product__00000001_.Name = `lot of yyy files and no clear mental model `
	__Product__00000001_.Description = ``
	__Product__00000001_.IsProducersNodeExpanded = false
	__Product__00000001_.IsConsumersNodeExpanded = false
	__Product__00000001_.IsImport = false
	__Product__00000001_.ComputedPrefix = `1.1`
	__Product__00000001_.IsExpanded = false
	__Product__00000001_.LayoutDirection = models.Vertical

	__Product__00000002_.Name = `lot of stager is a hybrid concept`
	__Product__00000002_.Description = ``
	__Product__00000002_.IsProducersNodeExpanded = false
	__Product__00000002_.IsConsumersNodeExpanded = false
	__Product__00000002_.IsImport = false
	__Product__00000002_.ComputedPrefix = `1.2`
	__Product__00000002_.IsExpanded = false
	__Product__00000002_.LayoutDirection = models.Vertical

	__Product__00000003_.Name = `the stager concept is not clearly delineated`
	__Product__00000003_.Description = ``
	__Product__00000003_.IsProducersNodeExpanded = false
	__Product__00000003_.IsConsumersNodeExpanded = false
	__Product__00000003_.IsImport = false
	__Product__00000003_.ComputedPrefix = `1.3`
	__Product__00000003_.IsExpanded = false
	__Product__00000003_.LayoutDirection = models.Vertical

	__Product__00000004_.Name = `confusion between interface & generic code`
	__Product__00000004_.Description = ``
	__Product__00000004_.IsProducersNodeExpanded = false
	__Product__00000004_.IsConsumersNodeExpanded = false
	__Product__00000004_.IsImport = false
	__Product__00000004_.ComputedPrefix = `1.4`
	__Product__00000004_.IsExpanded = false
	__Product__00000004_.LayoutDirection = models.Vertical

	__Product__00000005_.Name = `There is a lifecycle "enforce sematntic", "marhsall", generate UXs`
	__Product__00000005_.Description = ``
	__Product__00000005_.IsProducersNodeExpanded = false
	__Product__00000005_.IsConsumersNodeExpanded = false
	__Product__00000005_.IsImport = false
	__Product__00000005_.ComputedPrefix = `1.3.1`
	__Product__00000005_.IsExpanded = false
	__Product__00000005_.LayoutDirection = models.Vertical

	__Product__00000006_.Name = `TBC`
	__Product__00000006_.Description = ``
	__Product__00000006_.IsProducersNodeExpanded = false
	__Product__00000006_.IsConsumersNodeExpanded = false
	__Product__00000006_.IsImport = false
	__Product__00000006_.ComputedPrefix = `1.3.2`
	__Product__00000006_.IsExpanded = false
	__Product__00000006_.LayoutDirection = models.Vertical

	__Product__00000007_.Name = `yyy files are helpers for UX updates of trees and svg`
	__Product__00000007_.Description = ``
	__Product__00000007_.IsProducersNodeExpanded = false
	__Product__00000007_.IsConsumersNodeExpanded = false
	__Product__00000007_.IsImport = false
	__Product__00000007_.ComputedPrefix = `1.1.1`
	__Product__00000007_.IsExpanded = false
	__Product__00000007_.LayoutDirection = models.Horizontal

	__Product__00000008_.Name = `Core Abstractions & Interfaces`
	__Product__00000008_.Description = ``
	__Product__00000008_.IsProducersNodeExpanded = false
	__Product__00000008_.IsConsumersNodeExpanded = false
	__Product__00000008_.IsImport = false
	__Product__00000008_.ComputedPrefix = `1.1.1.1`
	__Product__00000008_.IsExpanded = false
	__Product__00000008_.LayoutDirection = models.Vertical

	__Product__00000009_.Name = `Tree UI & Context Menus`
	__Product__00000009_.Description = ``
	__Product__00000009_.IsProducersNodeExpanded = false
	__Product__00000009_.IsConsumersNodeExpanded = false
	__Product__00000009_.IsImport = false
	__Product__00000009_.ComputedPrefix = `1.1.1.2`
	__Product__00000009_.IsExpanded = false
	__Product__00000009_.LayoutDirection = models.Vertical

	__Product__00000010_.Name = `UI Callbacks & State Syncing`
	__Product__00000010_.Description = ``
	__Product__00000010_.IsProducersNodeExpanded = false
	__Product__00000010_.IsConsumersNodeExpanded = false
	__Product__00000010_.IsImport = false
	__Product__00000010_.ComputedPrefix = `1.1.1.3`
	__Product__00000010_.IsExpanded = false
	__Product__00000010_.LayoutDirection = models.Vertical

	__Product__00000011_.Name = `SVG Diagram Rendering & Interaction`
	__Product__00000011_.Description = ``
	__Product__00000011_.IsProducersNodeExpanded = false
	__Product__00000011_.IsConsumersNodeExpanded = false
	__Product__00000011_.IsImport = false
	__Product__00000011_.ComputedPrefix = `1.1.1.4`
	__Product__00000011_.IsExpanded = false
	__Product__00000011_.LayoutDirection = models.Vertical

	__Product__00000012_.Name = `Graph Data Integrity & Traversal`
	__Product__00000012_.Description = ``
	__Product__00000012_.IsProducersNodeExpanded = false
	__Product__00000012_.IsConsumersNodeExpanded = false
	__Product__00000012_.IsImport = false
	__Product__00000012_.ComputedPrefix = `1.1.1.5`
	__Product__00000012_.IsExpanded = false
	__Product__00000012_.LayoutDirection = models.Vertical

	__Product__00000013_.Name = `Rely on heavily typed Stage`
	__Product__00000013_.Description = ``
	__Product__00000013_.IsProducersNodeExpanded = false
	__Product__00000013_.IsConsumersNodeExpanded = false
	__Product__00000013_.IsImport = false
	__Product__00000013_.ComputedPrefix = `1.2.1`
	__Product__00000013_.IsExpanded = false
	__Product__00000013_.LayoutDirection = models.Vertical

	__Product__00000014_.Name = `lots of boilerplate code`
	__Product__00000014_.Description = ``
	__Product__00000014_.IsProducersNodeExpanded = false
	__Product__00000014_.IsConsumersNodeExpanded = false
	__Product__00000014_.IsImport = false
	__Product__00000014_.ComputedPrefix = `1.2.2`
	__Product__00000014_.IsExpanded = false
	__Product__00000014_.LayoutDirection = models.Vertical

	__Product__00000015_.Name = `Goals`
	__Product__00000015_.Description = ``
	__Product__00000015_.IsProducersNodeExpanded = false
	__Product__00000015_.IsConsumersNodeExpanded = false
	__Product__00000015_.IsImport = false
	__Product__00000015_.ComputedPrefix = `1`
	__Product__00000015_.IsExpanded = false
	__Product__00000015_.LayoutDirection = models.Vertical

	__Product__00000016_.Name = `creates a DSM in les than 1/2 hour`
	__Product__00000016_.Description = ``
	__Product__00000016_.IsProducersNodeExpanded = false
	__Product__00000016_.IsConsumersNodeExpanded = false
	__Product__00000016_.IsImport = false
	__Product__00000016_.ComputedPrefix = `1.1`
	__Product__00000016_.IsExpanded = false
	__Product__00000016_.LayoutDirection = models.Horizontal

	__Product__00000017_.Name = `Provides cleans abstraction for the development of the tree`
	__Product__00000017_.Description = ``
	__Product__00000017_.IsProducersNodeExpanded = false
	__Product__00000017_.IsConsumersNodeExpanded = false
	__Product__00000017_.IsImport = false
	__Product__00000017_.ComputedPrefix = `1.1.1`
	__Product__00000017_.IsExpanded = false
	__Product__00000017_.LayoutDirection = models.Vertical

	__Product__00000018_.Name = `Abstraction for semantic rules`
	__Product__00000018_.Description = ``
	__Product__00000018_.IsProducersNodeExpanded = false
	__Product__00000018_.IsConsumersNodeExpanded = false
	__Product__00000018_.IsImport = false
	__Product__00000018_.ComputedPrefix = `1.1.2`
	__Product__00000018_.IsExpanded = false
	__Product__00000018_.LayoutDirection = models.Vertical

	__Product__00000023_.Name = `Allows combination of DSM `
	__Product__00000023_.Description = ``
	__Product__00000023_.IsProducersNodeExpanded = false
	__Product__00000023_.IsConsumersNodeExpanded = false
	__Product__00000023_.IsImport = false
	__Product__00000023_.ComputedPrefix = `1.2`
	__Product__00000023_.IsExpanded = false
	__Product__00000023_.LayoutDirection = models.Vertical

	__Product__00000024_.Name = ``
	__Product__00000024_.Description = ``
	__Product__00000024_.IsProducersNodeExpanded = false
	__Product__00000024_.IsConsumersNodeExpanded = false
	__Product__00000024_.IsImport = false
	__Product__00000024_.ComputedPrefix = `2`
	__Product__00000024_.IsExpanded = false
	__Product__00000024_.LayoutDirection = models.Vertical

	__Product__00000025_.Name = ``
	__Product__00000025_.Description = ``
	__Product__00000025_.IsProducersNodeExpanded = false
	__Product__00000025_.IsConsumersNodeExpanded = false
	__Product__00000025_.IsImport = false
	__Product__00000025_.ComputedPrefix = `3`
	__Product__00000025_.IsExpanded = false
	__Product__00000025_.LayoutDirection = models.Vertical

	__Product__00000026_.Name = `sss`
	__Product__00000026_.Description = ``
	__Product__00000026_.IsProducersNodeExpanded = false
	__Product__00000026_.IsConsumersNodeExpanded = false
	__Product__00000026_.IsImport = false
	__Product__00000026_.ComputedPrefix = `4`
	__Product__00000026_.IsExpanded = false
	__Product__00000026_.LayoutDirection = models.Vertical

	__ProductCompositionShape__00000000_.Name = `Default Diagram-Paint points-lot of yyy files and no clear mental model `
	__ProductCompositionShape__00000000_.StartRatio = 0.500000
	__ProductCompositionShape__00000000_.EndRatio = 0.500000
	__ProductCompositionShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000000_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000000_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000000_.IsHidden = false

	__ProductCompositionShape__00000001_.Name = `Default Diagram-Paint points-lot of stager is a hybrid concept`
	__ProductCompositionShape__00000001_.StartRatio = 0.500000
	__ProductCompositionShape__00000001_.EndRatio = 0.500000
	__ProductCompositionShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000001_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000001_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000001_.IsHidden = false

	__ProductCompositionShape__00000002_.Name = `Default Diagram-Paint points-the stager concept is not clearly delineated`
	__ProductCompositionShape__00000002_.StartRatio = 0.500000
	__ProductCompositionShape__00000002_.EndRatio = 0.500000
	__ProductCompositionShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000002_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000002_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000002_.IsHidden = false

	__ProductCompositionShape__00000003_.Name = `Default Diagram-Paint points-confusion between interface & generic code`
	__ProductCompositionShape__00000003_.StartRatio = 0.500000
	__ProductCompositionShape__00000003_.EndRatio = 0.500000
	__ProductCompositionShape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000003_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000003_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000003_.IsHidden = false

	__ProductCompositionShape__00000004_.Name = `Default Diagram-the stager concept is not clearly delineated-There is a lifecycle "enforce sematntic", "marhsall", generate UXs`
	__ProductCompositionShape__00000004_.StartRatio = 0.500000
	__ProductCompositionShape__00000004_.EndRatio = 0.500000
	__ProductCompositionShape__00000004_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000004_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000004_.IsHidden = false

	__ProductCompositionShape__00000005_.Name = `Default Diagram-the stager concept is not clearly delineated-TBC`
	__ProductCompositionShape__00000005_.StartRatio = 0.500000
	__ProductCompositionShape__00000005_.EndRatio = 0.500000
	__ProductCompositionShape__00000005_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000005_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000005_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000005_.IsHidden = false

	__ProductCompositionShape__00000006_.Name = `Default Diagram-lot of yyy files and no clear mental model -yyy files are helpers for UX updates of trees and svg`
	__ProductCompositionShape__00000006_.StartRatio = 0.500000
	__ProductCompositionShape__00000006_.EndRatio = 0.500000
	__ProductCompositionShape__00000006_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000006_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000006_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000006_.IsHidden = false

	__ProductCompositionShape__00000007_.Name = `Default Diagram-yyy files are helpers for UX updates of trees and svg-Core Abstractions & Interfaces`
	__ProductCompositionShape__00000007_.StartRatio = 0.500000
	__ProductCompositionShape__00000007_.EndRatio = 0.500000
	__ProductCompositionShape__00000007_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000007_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000007_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000007_.IsHidden = false

	__ProductCompositionShape__00000008_.Name = `Default Diagram-yyy files are helpers for UX updates of trees and svg-Tree UI & Context Menus`
	__ProductCompositionShape__00000008_.StartRatio = 0.500000
	__ProductCompositionShape__00000008_.EndRatio = 0.500000
	__ProductCompositionShape__00000008_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000008_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000008_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000008_.IsHidden = false

	__ProductCompositionShape__00000009_.Name = `Default Diagram-yyy files are helpers for UX updates of trees and svg-UI Callbacks & State Syncing`
	__ProductCompositionShape__00000009_.StartRatio = 0.500000
	__ProductCompositionShape__00000009_.EndRatio = 0.500000
	__ProductCompositionShape__00000009_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000009_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000009_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000009_.IsHidden = false

	__ProductCompositionShape__00000010_.Name = `Default Diagram-yyy files are helpers for UX updates of trees and svg-SVG Diagram Rendering & Interaction`
	__ProductCompositionShape__00000010_.StartRatio = 0.500000
	__ProductCompositionShape__00000010_.EndRatio = 0.500000
	__ProductCompositionShape__00000010_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000010_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000010_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000010_.IsHidden = false

	__ProductCompositionShape__00000011_.Name = `Default Diagram-yyy files are helpers for UX updates of trees and svg-Graph Data Integrity & Traversal`
	__ProductCompositionShape__00000011_.StartRatio = 0.500000
	__ProductCompositionShape__00000011_.EndRatio = 0.500000
	__ProductCompositionShape__00000011_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000011_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000011_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000011_.IsHidden = false

	__ProductCompositionShape__00000012_.Name = `Default Diagram-lot of stager is a hybrid concept-Rely on heavily typed Stage`
	__ProductCompositionShape__00000012_.StartRatio = 0.500000
	__ProductCompositionShape__00000012_.EndRatio = 0.500000
	__ProductCompositionShape__00000012_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000012_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000012_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000012_.IsHidden = false

	__ProductCompositionShape__00000013_.Name = `Default Diagram-lot of stager is a hybrid concept-lots of boilerplate code`
	__ProductCompositionShape__00000013_.StartRatio = 0.500000
	__ProductCompositionShape__00000013_.EndRatio = 0.500000
	__ProductCompositionShape__00000013_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000013_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000013_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000013_.IsHidden = false

	__ProductCompositionShape__00000014_.Name = `Default Diagram-Goals-creates a DSM in les than 1/2 hour`
	__ProductCompositionShape__00000014_.StartRatio = 0.500000
	__ProductCompositionShape__00000014_.EndRatio = 0.500000
	__ProductCompositionShape__00000014_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000014_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000014_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000014_.IsHidden = false

	__ProductCompositionShape__00000015_.Name = `Default Diagram-creates a DSM in les than 1/2 hour-Provides cleans abstraction for the development of the tree`
	__ProductCompositionShape__00000015_.StartRatio = 0.500000
	__ProductCompositionShape__00000015_.EndRatio = 0.500000
	__ProductCompositionShape__00000015_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000015_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000015_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000015_.IsHidden = false

	__ProductCompositionShape__00000017_.Name = `Default Diagram-creates a DSM in les than 1/2 hour-Abstraction for semantic rules`
	__ProductCompositionShape__00000017_.StartRatio = 0.500000
	__ProductCompositionShape__00000017_.EndRatio = 0.500000
	__ProductCompositionShape__00000017_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000017_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000017_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000017_.IsHidden = false

	__ProductCompositionShape__00000020_.Name = `Default Diagram-Goals-Allows combination of DSM `
	__ProductCompositionShape__00000020_.StartRatio = 0.500000
	__ProductCompositionShape__00000020_.EndRatio = 0.500000
	__ProductCompositionShape__00000020_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000020_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000020_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000020_.IsHidden = false

	__ProductShape__00000000_.Name = `Default Diagram-Paint points`
	__ProductShape__00000000_.OverideLayoutDirection = false
	__ProductShape__00000000_.LayoutDirection = models.Vertical
	__ProductShape__00000000_.X = 50.000000
	__ProductShape__00000000_.Y = 50.000000
	__ProductShape__00000000_.Width = 250.000000
	__ProductShape__00000000_.Height = 70.000000
	__ProductShape__00000000_.IsHidden = false

	__ProductShape__00000001_.Name = `Default Diagram-lot of yyy files and no clear mental model `
	__ProductShape__00000001_.OverideLayoutDirection = false
	__ProductShape__00000001_.LayoutDirection = models.Vertical
	__ProductShape__00000001_.X = 225.000000
	__ProductShape__00000001_.Y = 135.000000
	__ProductShape__00000001_.Width = 250.000000
	__ProductShape__00000001_.Height = 70.000000
	__ProductShape__00000001_.IsHidden = false

	__ProductShape__00000002_.Name = `Default Diagram-lot of stager is a hybrid concept`
	__ProductShape__00000002_.OverideLayoutDirection = false
	__ProductShape__00000002_.LayoutDirection = models.Vertical
	__ProductShape__00000002_.X = 225.000000
	__ProductShape__00000002_.Y = 785.000000
	__ProductShape__00000002_.Width = 250.000000
	__ProductShape__00000002_.Height = 70.000000
	__ProductShape__00000002_.IsHidden = false

	__ProductShape__00000003_.Name = `Default Diagram-the stager concept is not clearly delineated`
	__ProductShape__00000003_.OverideLayoutDirection = false
	__ProductShape__00000003_.LayoutDirection = models.Vertical
	__ProductShape__00000003_.X = 225.000000
	__ProductShape__00000003_.Y = 1010.000000
	__ProductShape__00000003_.Width = 250.000000
	__ProductShape__00000003_.Height = 70.000000
	__ProductShape__00000003_.IsHidden = false

	__ProductShape__00000004_.Name = `Default Diagram-confusion between interface & generic code`
	__ProductShape__00000004_.OverideLayoutDirection = false
	__ProductShape__00000004_.LayoutDirection = models.Vertical
	__ProductShape__00000004_.X = 225.000000
	__ProductShape__00000004_.Y = 1235.000000
	__ProductShape__00000004_.Width = 250.000000
	__ProductShape__00000004_.Height = 70.000000
	__ProductShape__00000004_.IsHidden = false

	__ProductShape__00000005_.Name = `Default Diagram-There is a lifecycle "enforce sematntic", "marhsall", generate UXs`
	__ProductShape__00000005_.OverideLayoutDirection = false
	__ProductShape__00000005_.LayoutDirection = models.Vertical
	__ProductShape__00000005_.X = 400.000000
	__ProductShape__00000005_.Y = 1150.000000
	__ProductShape__00000005_.Width = 250.000000
	__ProductShape__00000005_.Height = 70.000000
	__ProductShape__00000005_.IsHidden = false

	__ProductShape__00000006_.Name = `Default Diagram-TBC`
	__ProductShape__00000006_.OverideLayoutDirection = false
	__ProductShape__00000006_.LayoutDirection = models.Vertical
	__ProductShape__00000006_.X = 700.000000
	__ProductShape__00000006_.Y = 1150.000000
	__ProductShape__00000006_.Width = 250.000000
	__ProductShape__00000006_.Height = 70.000000
	__ProductShape__00000006_.IsHidden = false

	__ProductShape__00000007_.Name = `Default Diagram-yyy files are helpers for UX updates of trees and svg`
	__ProductShape__00000007_.OverideLayoutDirection = false
	__ProductShape__00000007_.LayoutDirection = models.Vertical
	__ProductShape__00000007_.X = 400.000000
	__ProductShape__00000007_.Y = 275.000000
	__ProductShape__00000007_.Width = 250.000000
	__ProductShape__00000007_.Height = 70.000000
	__ProductShape__00000007_.IsHidden = false

	__ProductShape__00000008_.Name = `Default Diagram-Core Abstractions & Interfaces`
	__ProductShape__00000008_.OverideLayoutDirection = false
	__ProductShape__00000008_.LayoutDirection = models.Vertical
	__ProductShape__00000008_.X = 575.000000
	__ProductShape__00000008_.Y = 360.000000
	__ProductShape__00000008_.Width = 250.000000
	__ProductShape__00000008_.Height = 70.000000
	__ProductShape__00000008_.IsHidden = false

	__ProductShape__00000009_.Name = `Default Diagram-Tree UI & Context Menus`
	__ProductShape__00000009_.OverideLayoutDirection = false
	__ProductShape__00000009_.LayoutDirection = models.Vertical
	__ProductShape__00000009_.X = 575.000000
	__ProductShape__00000009_.Y = 445.000000
	__ProductShape__00000009_.Width = 250.000000
	__ProductShape__00000009_.Height = 70.000000
	__ProductShape__00000009_.IsHidden = false

	__ProductShape__00000010_.Name = `Default Diagram-UI Callbacks & State Syncing`
	__ProductShape__00000010_.OverideLayoutDirection = false
	__ProductShape__00000010_.LayoutDirection = models.Vertical
	__ProductShape__00000010_.X = 575.000000
	__ProductShape__00000010_.Y = 530.000000
	__ProductShape__00000010_.Width = 250.000000
	__ProductShape__00000010_.Height = 70.000000
	__ProductShape__00000010_.IsHidden = false

	__ProductShape__00000011_.Name = `Default Diagram-SVG Diagram Rendering & Interaction`
	__ProductShape__00000011_.OverideLayoutDirection = false
	__ProductShape__00000011_.LayoutDirection = models.Vertical
	__ProductShape__00000011_.X = 575.000000
	__ProductShape__00000011_.Y = 615.000000
	__ProductShape__00000011_.Width = 250.000000
	__ProductShape__00000011_.Height = 70.000000
	__ProductShape__00000011_.IsHidden = false

	__ProductShape__00000012_.Name = `Default Diagram-Graph Data Integrity & Traversal`
	__ProductShape__00000012_.OverideLayoutDirection = false
	__ProductShape__00000012_.LayoutDirection = models.Vertical
	__ProductShape__00000012_.X = 575.000000
	__ProductShape__00000012_.Y = 700.000000
	__ProductShape__00000012_.Width = 250.000000
	__ProductShape__00000012_.Height = 70.000000
	__ProductShape__00000012_.IsHidden = false

	__ProductShape__00000013_.Name = `Default Diagram-Rely on heavily typed Stage`
	__ProductShape__00000013_.OverideLayoutDirection = false
	__ProductShape__00000013_.LayoutDirection = models.Vertical
	__ProductShape__00000013_.X = 400.000000
	__ProductShape__00000013_.Y = 925.000000
	__ProductShape__00000013_.Width = 250.000000
	__ProductShape__00000013_.Height = 70.000000
	__ProductShape__00000013_.IsHidden = false

	__ProductShape__00000014_.Name = `Default Diagram-lots of boilerplate code`
	__ProductShape__00000014_.OverideLayoutDirection = false
	__ProductShape__00000014_.LayoutDirection = models.Vertical
	__ProductShape__00000014_.X = 700.000000
	__ProductShape__00000014_.Y = 925.000000
	__ProductShape__00000014_.Width = 250.000000
	__ProductShape__00000014_.Height = 70.000000
	__ProductShape__00000014_.IsHidden = false

	__ProductShape__00000015_.Name = `Default Diagram-Goals`
	__ProductShape__00000015_.OverideLayoutDirection = false
	__ProductShape__00000015_.LayoutDirection = models.Vertical
	__ProductShape__00000015_.X = 523.000000
	__ProductShape__00000015_.Y = 51.000000
	__ProductShape__00000015_.Width = 250.000000
	__ProductShape__00000015_.Height = 70.000000
	__ProductShape__00000015_.IsHidden = false

	__ProductShape__00000016_.Name = `Default Diagram-creates a DSM in les than 1/2 hour`
	__ProductShape__00000016_.OverideLayoutDirection = false
	__ProductShape__00000016_.LayoutDirection = models.Vertical
	__ProductShape__00000016_.X = 523.000000
	__ProductShape__00000016_.Y = 191.000000
	__ProductShape__00000016_.Width = 250.000000
	__ProductShape__00000016_.Height = 70.000000
	__ProductShape__00000016_.IsHidden = false

	__ProductShape__00000017_.Name = `Default Diagram-Provides cleans abstraction for the development of the tree`
	__ProductShape__00000017_.OverideLayoutDirection = false
	__ProductShape__00000017_.LayoutDirection = models.Vertical
	__ProductShape__00000017_.X = 698.000000
	__ProductShape__00000017_.Y = 276.000000
	__ProductShape__00000017_.Width = 250.000000
	__ProductShape__00000017_.Height = 70.000000
	__ProductShape__00000017_.IsHidden = false

	__ProductShape__00000018_.Name = `Default Diagram-Abstraction for semantic rules`
	__ProductShape__00000018_.OverideLayoutDirection = false
	__ProductShape__00000018_.LayoutDirection = models.Vertical
	__ProductShape__00000018_.X = 698.000000
	__ProductShape__00000018_.Y = 361.000000
	__ProductShape__00000018_.Width = 250.000000
	__ProductShape__00000018_.Height = 70.000000
	__ProductShape__00000018_.IsHidden = false

	__ProductShape__00000025_.Name = `Default Diagram-Allows combination of DSM `
	__ProductShape__00000025_.OverideLayoutDirection = false
	__ProductShape__00000025_.LayoutDirection = models.Vertical
	__ProductShape__00000025_.X = 823.000000
	__ProductShape__00000025_.Y = 191.000000
	__ProductShape__00000025_.Width = 250.000000
	__ProductShape__00000025_.Height = 70.000000
	__ProductShape__00000025_.IsHidden = false

	__ProductShape__00000028_.Name = `Default Diagram-sss`
	__ProductShape__00000028_.OverideLayoutDirection = false
	__ProductShape__00000028_.LayoutDirection = models.Vertical
	__ProductShape__00000028_.X = 169.084049
	__ProductShape__00000028_.Y = 764.000000
	__ProductShape__00000028_.Width = 250.000000
	__ProductShape__00000028_.Height = 70.000000
	__ProductShape__00000028_.IsHidden = false

	// insertion point for setup of pointers
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000000_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000001_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000002_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000003_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000004_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000005_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000006_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000007_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000008_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000009_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000010_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000011_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000012_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000013_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000014_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, __Product__00000003_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, __Product__00000001_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, __Product__00000007_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, __Product__00000002_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, __Product__00000005_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000000_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000001_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000002_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000003_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000004_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000005_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000006_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000007_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000008_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000009_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000010_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000011_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000012_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000013_)
	__Diagram__00000000_.Note_Shapes = append(__Diagram__00000000_.Note_Shapes, __NoteShape__00000000_)
	__Diagram__00000000_.NoteProductShapes = append(__Diagram__00000000_.NoteProductShapes, __NoteProductShape__00000000_)
	__Diagram__00000001_.Product_Shapes = append(__Diagram__00000001_.Product_Shapes, __ProductShape__00000015_)
	__Diagram__00000001_.Product_Shapes = append(__Diagram__00000001_.Product_Shapes, __ProductShape__00000016_)
	__Diagram__00000001_.Product_Shapes = append(__Diagram__00000001_.Product_Shapes, __ProductShape__00000017_)
	__Diagram__00000001_.Product_Shapes = append(__Diagram__00000001_.Product_Shapes, __ProductShape__00000018_)
	__Diagram__00000001_.Product_Shapes = append(__Diagram__00000001_.Product_Shapes, __ProductShape__00000025_)
	__Diagram__00000001_.Product_Shapes = append(__Diagram__00000001_.Product_Shapes, __ProductShape__00000028_)
	__Diagram__00000001_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000001_.ProductsWhoseNodeIsExpanded, __Product__00000016_)
	__Diagram__00000001_.ProductComposition_Shapes = append(__Diagram__00000001_.ProductComposition_Shapes, __ProductCompositionShape__00000014_)
	__Diagram__00000001_.ProductComposition_Shapes = append(__Diagram__00000001_.ProductComposition_Shapes, __ProductCompositionShape__00000015_)
	__Diagram__00000001_.ProductComposition_Shapes = append(__Diagram__00000001_.ProductComposition_Shapes, __ProductCompositionShape__00000017_)
	__Diagram__00000001_.ProductComposition_Shapes = append(__Diagram__00000001_.ProductComposition_Shapes, __ProductCompositionShape__00000020_)
	__Diagram__00000001_.Note_Shapes = append(__Diagram__00000001_.Note_Shapes, __NoteShape__00000001_)
	__Diagram__00000001_.Note_Shapes = append(__Diagram__00000001_.Note_Shapes, __NoteShape__00000002_)
	__Diagram__00000001_.Note_Shapes = append(__Diagram__00000001_.Note_Shapes, __NoteShape__00000003_)
	__Diagram__00000001_.Note_Shapes = append(__Diagram__00000001_.Note_Shapes, __NoteShape__00000004_)
	__Diagram__00000001_.Note_Shapes = append(__Diagram__00000001_.Note_Shapes, __NoteShape__00000007_)
	__Library__00000000_.SubLibraries = append(__Library__00000000_.SubLibraries, __Library__00000001_)
	__Library__00000000_.RootProducts = append(__Library__00000000_.RootProducts, __Product__00000000_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000000_)
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000000_)
	__Library__00000001_.RootProducts = append(__Library__00000001_.RootProducts, __Product__00000015_)
	__Library__00000001_.RootProducts = append(__Library__00000001_.RootProducts, __Product__00000024_)
	__Library__00000001_.RootProducts = append(__Library__00000001_.RootProducts, __Product__00000025_)
	__Library__00000001_.RootProducts = append(__Library__00000001_.RootProducts, __Product__00000026_)
	__Library__00000001_.Notes = append(__Library__00000001_.Notes, __Note__00000001_)
	__Library__00000001_.Notes = append(__Library__00000001_.Notes, __Note__00000002_)
	__Library__00000001_.Notes = append(__Library__00000001_.Notes, __Note__00000003_)
	__Library__00000001_.Notes = append(__Library__00000001_.Notes, __Note__00000004_)
	__Library__00000001_.Notes = append(__Library__00000001_.Notes, __Note__00000005_)
	__Library__00000001_.Diagrams = append(__Library__00000001_.Diagrams, __Diagram__00000001_)
	__Note__00000000_.Products = append(__Note__00000000_.Products, __Product__00000004_)
	__Note__00000005_.Products = append(__Note__00000005_.Products, __Product__00000023_)
	__NoteProductShape__00000000_.Note = __Note__00000000_
	__NoteProductShape__00000000_.Product = __Product__00000004_
	__NoteShape__00000000_.Note = __Note__00000000_
	__NoteShape__00000001_.Note = __Note__00000001_
	__NoteShape__00000002_.Note = __Note__00000002_
	__NoteShape__00000003_.Note = __Note__00000003_
	__NoteShape__00000004_.Note = __Note__00000004_
	__NoteShape__00000007_.Note = __Note__00000005_
	__Product__00000000_.SubProducts = append(__Product__00000000_.SubProducts, __Product__00000001_)
	__Product__00000000_.SubProducts = append(__Product__00000000_.SubProducts, __Product__00000002_)
	__Product__00000000_.SubProducts = append(__Product__00000000_.SubProducts, __Product__00000003_)
	__Product__00000000_.SubProducts = append(__Product__00000000_.SubProducts, __Product__00000004_)
	__Product__00000000_.ReferencedProduct = nil
	__Product__00000001_.SubProducts = append(__Product__00000001_.SubProducts, __Product__00000007_)
	__Product__00000001_.ReferencedProduct = nil
	__Product__00000002_.SubProducts = append(__Product__00000002_.SubProducts, __Product__00000013_)
	__Product__00000002_.SubProducts = append(__Product__00000002_.SubProducts, __Product__00000014_)
	__Product__00000002_.ReferencedProduct = nil
	__Product__00000003_.SubProducts = append(__Product__00000003_.SubProducts, __Product__00000005_)
	__Product__00000003_.SubProducts = append(__Product__00000003_.SubProducts, __Product__00000006_)
	__Product__00000003_.ReferencedProduct = nil
	__Product__00000004_.ReferencedProduct = nil
	__Product__00000005_.ReferencedProduct = nil
	__Product__00000006_.ReferencedProduct = nil
	__Product__00000007_.SubProducts = append(__Product__00000007_.SubProducts, __Product__00000008_)
	__Product__00000007_.SubProducts = append(__Product__00000007_.SubProducts, __Product__00000009_)
	__Product__00000007_.SubProducts = append(__Product__00000007_.SubProducts, __Product__00000010_)
	__Product__00000007_.SubProducts = append(__Product__00000007_.SubProducts, __Product__00000011_)
	__Product__00000007_.SubProducts = append(__Product__00000007_.SubProducts, __Product__00000012_)
	__Product__00000007_.ReferencedProduct = nil
	__Product__00000008_.ReferencedProduct = nil
	__Product__00000009_.ReferencedProduct = nil
	__Product__00000010_.ReferencedProduct = nil
	__Product__00000011_.ReferencedProduct = nil
	__Product__00000012_.ReferencedProduct = nil
	__Product__00000013_.ReferencedProduct = nil
	__Product__00000014_.ReferencedProduct = nil
	__Product__00000015_.SubProducts = append(__Product__00000015_.SubProducts, __Product__00000016_)
	__Product__00000015_.SubProducts = append(__Product__00000015_.SubProducts, __Product__00000023_)
	__Product__00000015_.ReferencedProduct = nil
	__Product__00000016_.SubProducts = append(__Product__00000016_.SubProducts, __Product__00000017_)
	__Product__00000016_.SubProducts = append(__Product__00000016_.SubProducts, __Product__00000018_)
	__Product__00000016_.ReferencedProduct = nil
	__Product__00000017_.ReferencedProduct = nil
	__Product__00000018_.ReferencedProduct = nil
	__Product__00000023_.ReferencedProduct = nil
	__Product__00000024_.ReferencedProduct = nil
	__Product__00000025_.ReferencedProduct = nil
	__Product__00000026_.ReferencedProduct = nil
	__ProductCompositionShape__00000000_.Product = __Product__00000001_
	__ProductCompositionShape__00000001_.Product = __Product__00000002_
	__ProductCompositionShape__00000002_.Product = __Product__00000003_
	__ProductCompositionShape__00000003_.Product = __Product__00000004_
	__ProductCompositionShape__00000004_.Product = __Product__00000005_
	__ProductCompositionShape__00000005_.Product = __Product__00000006_
	__ProductCompositionShape__00000006_.Product = __Product__00000007_
	__ProductCompositionShape__00000007_.Product = __Product__00000008_
	__ProductCompositionShape__00000008_.Product = __Product__00000009_
	__ProductCompositionShape__00000009_.Product = __Product__00000010_
	__ProductCompositionShape__00000010_.Product = __Product__00000011_
	__ProductCompositionShape__00000011_.Product = __Product__00000012_
	__ProductCompositionShape__00000012_.Product = __Product__00000013_
	__ProductCompositionShape__00000013_.Product = __Product__00000014_
	__ProductCompositionShape__00000014_.Product = __Product__00000016_
	__ProductCompositionShape__00000015_.Product = __Product__00000017_
	__ProductCompositionShape__00000017_.Product = __Product__00000018_
	__ProductCompositionShape__00000020_.Product = __Product__00000023_
	__ProductShape__00000000_.Product = __Product__00000000_
	__ProductShape__00000001_.Product = __Product__00000001_
	__ProductShape__00000002_.Product = __Product__00000002_
	__ProductShape__00000003_.Product = __Product__00000003_
	__ProductShape__00000004_.Product = __Product__00000004_
	__ProductShape__00000005_.Product = __Product__00000005_
	__ProductShape__00000006_.Product = __Product__00000006_
	__ProductShape__00000007_.Product = __Product__00000007_
	__ProductShape__00000008_.Product = __Product__00000008_
	__ProductShape__00000009_.Product = __Product__00000009_
	__ProductShape__00000010_.Product = __Product__00000010_
	__ProductShape__00000011_.Product = __Product__00000011_
	__ProductShape__00000012_.Product = __Product__00000012_
	__ProductShape__00000013_.Product = __Product__00000013_
	__ProductShape__00000014_.Product = __Product__00000014_
	__ProductShape__00000015_.Product = __Product__00000015_
	__ProductShape__00000016_.Product = __Product__00000016_
	__ProductShape__00000017_.Product = __Product__00000017_
	__ProductShape__00000018_.Product = __Product__00000018_
	__ProductShape__00000025_.Product = __Product__00000023_
	__ProductShape__00000028_.Product = __Product__00000026_
}
