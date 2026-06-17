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

	__Library__00000000_ := (&models.Library{Name: ``}).Stage(stage)

	__Note__00000000_ := (&models.Note{Name: `Invariant : generated code never import a gong package `}).Stage(stage)

	__NoteProductShape__00000000_ := (&models.NoteProductShape{Name: `Invariant : generated code never import a gong package  to confusion between interface & generic code`}).Stage(stage)

	__NoteShape__00000000_ := (&models.NoteShape{Name: `-Default Diagram`}).Stage(stage)

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

	__ProductCompositionShape__00000000_ := (&models.ProductCompositionShape{Name: `Paint points to `}).Stage(stage)
	__ProductCompositionShape__00000001_ := (&models.ProductCompositionShape{Name: `Paint points to `}).Stage(stage)
	__ProductCompositionShape__00000002_ := (&models.ProductCompositionShape{Name: `Paint points to `}).Stage(stage)
	__ProductCompositionShape__00000003_ := (&models.ProductCompositionShape{Name: `Paint points to `}).Stage(stage)
	__ProductCompositionShape__00000004_ := (&models.ProductCompositionShape{Name: `the stager concept is not clearly delineated to `}).Stage(stage)
	__ProductCompositionShape__00000005_ := (&models.ProductCompositionShape{Name: `the stager concept is not clearly delineated to `}).Stage(stage)
	__ProductCompositionShape__00000006_ := (&models.ProductCompositionShape{Name: `lot of yyy files and no clear mental model  to `}).Stage(stage)
	__ProductCompositionShape__00000007_ := (&models.ProductCompositionShape{Name: `yyy files are helpers for UX updates of trees and svg to `}).Stage(stage)
	__ProductCompositionShape__00000008_ := (&models.ProductCompositionShape{Name: `yyy files are helpers for UX updates of trees and svg to `}).Stage(stage)
	__ProductCompositionShape__00000009_ := (&models.ProductCompositionShape{Name: `yyy files are helpers for UX updates of trees and svg to `}).Stage(stage)
	__ProductCompositionShape__00000010_ := (&models.ProductCompositionShape{Name: `yyy files are helpers for UX updates of trees and svg to `}).Stage(stage)
	__ProductCompositionShape__00000011_ := (&models.ProductCompositionShape{Name: `yyy files are helpers for UX updates of trees and svg to `}).Stage(stage)
	__ProductCompositionShape__00000012_ := (&models.ProductCompositionShape{Name: `lot of boilerplate code for stager to `}).Stage(stage)
	__ProductCompositionShape__00000013_ := (&models.ProductCompositionShape{Name: `lot of boilerplate code for stager. stager is a hybrid concept to `}).Stage(stage)

	__ProductShape__00000000_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000001_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000002_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000003_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000004_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000005_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000006_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000007_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000008_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000009_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000010_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000011_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000012_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000013_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)
	__ProductShape__00000014_ := (&models.ProductShape{Name: `-Default Diagram`}).Stage(stage)

	// insertion point for initialization of values

	__Diagram__00000000_.Name = `Default Diagram`
	__Diagram__00000000_.ComputedPrefix = `1`
	__Diagram__00000000_.IsExpanded = true
	__Diagram__00000000_.LayoutDirection = models.Vertical
	__Diagram__00000000_.IsChecked = true
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.IsShowPrefix = false
	__Diagram__00000000_.DefaultBoxWidth = 250.000000
	__Diagram__00000000_.DefaultBoxHeigth = 70.000000
	__Diagram__00000000_.Width = 1269.706973
	__Diagram__00000000_.Height = 1405.000000
	__Diagram__00000000_.IsPBSNodeExpanded = true
	__Diagram__00000000_.IsWBSNodeExpanded = true
	__Diagram__00000000_.IsTaskGroupsNodeExpanded = false
	__Diagram__00000000_.DateFormat = ``
	__Diagram__00000000_.IsNotesNodeExpanded = true
	__Diagram__00000000_.IsResourcesNodeExpanded = false
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

	__Library__00000000_.Name = ``
	__Library__00000000_.NbPixPerCharacter = 8.000000
	__Library__00000000_.LogoSVGFile = ``
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsExpanded = true
	__Library__00000000_.LayoutDirection = models.Vertical
	__Library__00000000_.IsRootLibrary = true

	__Note__00000000_.Name = `Invariant : generated code never import a gong package `
	__Note__00000000_.ComputedPrefix = `1`
	__Note__00000000_.IsExpanded = false
	__Note__00000000_.LayoutDirection = models.Vertical

	__NoteProductShape__00000000_.Name = `Invariant : generated code never import a gong package  to confusion between interface & generic code`
	__NoteProductShape__00000000_.StartRatio = 0.500000
	__NoteProductShape__00000000_.EndRatio = 0.500000
	__NoteProductShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000000_.CornerOffsetRatio = 1.680000
	__NoteProductShape__00000000_.IsHidden = false

	__NoteShape__00000000_.Name = `-Default Diagram`
	__NoteShape__00000000_.OverideLayoutDirection = false
	__NoteShape__00000000_.LayoutDirection = models.Vertical
	__NoteShape__00000000_.X = 919.706973
	__NoteShape__00000000_.Y = 1125.285144
	__NoteShape__00000000_.Width = 250.000000
	__NoteShape__00000000_.Height = 70.000000
	__NoteShape__00000000_.IsHidden = false

	__Product__00000000_.Name = `Paint points`
	__Product__00000000_.ComputedPrefix = `1`
	__Product__00000000_.IsExpanded = false
	__Product__00000000_.LayoutDirection = models.Horizontal
	__Product__00000000_.IsImport = false
	__Product__00000000_.Description = ``
	__Product__00000000_.IsProducersNodeExpanded = false
	__Product__00000000_.IsConsumersNodeExpanded = false

	__Product__00000001_.Name = `lot of yyy files and no clear mental model `
	__Product__00000001_.ComputedPrefix = `1.1`
	__Product__00000001_.IsExpanded = false
	__Product__00000001_.LayoutDirection = models.Vertical
	__Product__00000001_.IsImport = false
	__Product__00000001_.Description = ``
	__Product__00000001_.IsProducersNodeExpanded = false
	__Product__00000001_.IsConsumersNodeExpanded = false

	__Product__00000002_.Name = `lot of stager is a hybrid concept`
	__Product__00000002_.ComputedPrefix = `1.2`
	__Product__00000002_.IsExpanded = false
	__Product__00000002_.LayoutDirection = models.Vertical
	__Product__00000002_.IsImport = false
	__Product__00000002_.Description = ``
	__Product__00000002_.IsProducersNodeExpanded = false
	__Product__00000002_.IsConsumersNodeExpanded = false

	__Product__00000003_.Name = `the stager concept is not clearly delineated`
	__Product__00000003_.ComputedPrefix = `1.3`
	__Product__00000003_.IsExpanded = false
	__Product__00000003_.LayoutDirection = models.Vertical
	__Product__00000003_.IsImport = false
	__Product__00000003_.Description = ``
	__Product__00000003_.IsProducersNodeExpanded = false
	__Product__00000003_.IsConsumersNodeExpanded = false

	__Product__00000004_.Name = `confusion between interface & generic code`
	__Product__00000004_.ComputedPrefix = `1.4`
	__Product__00000004_.IsExpanded = false
	__Product__00000004_.LayoutDirection = models.Vertical
	__Product__00000004_.IsImport = false
	__Product__00000004_.Description = ``
	__Product__00000004_.IsProducersNodeExpanded = false
	__Product__00000004_.IsConsumersNodeExpanded = false

	__Product__00000005_.Name = `There is a lifecycle "enforce sematntic", "marhsall", generate UXs`
	__Product__00000005_.ComputedPrefix = `1.3.1`
	__Product__00000005_.IsExpanded = false
	__Product__00000005_.LayoutDirection = models.Vertical
	__Product__00000005_.IsImport = false
	__Product__00000005_.Description = ``
	__Product__00000005_.IsProducersNodeExpanded = false
	__Product__00000005_.IsConsumersNodeExpanded = false

	__Product__00000006_.Name = `TBC`
	__Product__00000006_.ComputedPrefix = `1.3.2`
	__Product__00000006_.IsExpanded = false
	__Product__00000006_.LayoutDirection = models.Vertical
	__Product__00000006_.IsImport = false
	__Product__00000006_.Description = ``
	__Product__00000006_.IsProducersNodeExpanded = false
	__Product__00000006_.IsConsumersNodeExpanded = false

	__Product__00000007_.Name = `yyy files are helpers for UX updates of trees and svg`
	__Product__00000007_.ComputedPrefix = `1.1.1`
	__Product__00000007_.IsExpanded = false
	__Product__00000007_.LayoutDirection = models.Horizontal
	__Product__00000007_.IsImport = false
	__Product__00000007_.Description = ``
	__Product__00000007_.IsProducersNodeExpanded = false
	__Product__00000007_.IsConsumersNodeExpanded = false

	__Product__00000008_.Name = `Core Abstractions & Interfaces`
	__Product__00000008_.ComputedPrefix = `1.1.1.1`
	__Product__00000008_.IsExpanded = false
	__Product__00000008_.LayoutDirection = models.Vertical
	__Product__00000008_.IsImport = false
	__Product__00000008_.Description = ``
	__Product__00000008_.IsProducersNodeExpanded = false
	__Product__00000008_.IsConsumersNodeExpanded = false

	__Product__00000009_.Name = `Tree UI & Context Menus`
	__Product__00000009_.ComputedPrefix = `1.1.1.2`
	__Product__00000009_.IsExpanded = false
	__Product__00000009_.LayoutDirection = models.Vertical
	__Product__00000009_.IsImport = false
	__Product__00000009_.Description = ``
	__Product__00000009_.IsProducersNodeExpanded = false
	__Product__00000009_.IsConsumersNodeExpanded = false

	__Product__00000010_.Name = `UI Callbacks & State Syncing`
	__Product__00000010_.ComputedPrefix = `1.1.1.3`
	__Product__00000010_.IsExpanded = false
	__Product__00000010_.LayoutDirection = models.Vertical
	__Product__00000010_.IsImport = false
	__Product__00000010_.Description = ``
	__Product__00000010_.IsProducersNodeExpanded = false
	__Product__00000010_.IsConsumersNodeExpanded = false

	__Product__00000011_.Name = `SVG Diagram Rendering & Interaction`
	__Product__00000011_.ComputedPrefix = `1.1.1.4`
	__Product__00000011_.IsExpanded = false
	__Product__00000011_.LayoutDirection = models.Vertical
	__Product__00000011_.IsImport = false
	__Product__00000011_.Description = ``
	__Product__00000011_.IsProducersNodeExpanded = false
	__Product__00000011_.IsConsumersNodeExpanded = false

	__Product__00000012_.Name = `Graph Data Integrity & Traversal`
	__Product__00000012_.ComputedPrefix = `1.1.1.5`
	__Product__00000012_.IsExpanded = false
	__Product__00000012_.LayoutDirection = models.Vertical
	__Product__00000012_.IsImport = false
	__Product__00000012_.Description = ``
	__Product__00000012_.IsProducersNodeExpanded = false
	__Product__00000012_.IsConsumersNodeExpanded = false

	__Product__00000013_.Name = `Rely on heavily typed Stage`
	__Product__00000013_.ComputedPrefix = `1.2.1`
	__Product__00000013_.IsExpanded = false
	__Product__00000013_.LayoutDirection = models.Vertical
	__Product__00000013_.IsImport = false
	__Product__00000013_.Description = ``
	__Product__00000013_.IsProducersNodeExpanded = false
	__Product__00000013_.IsConsumersNodeExpanded = false

	__Product__00000014_.Name = `lots of boilerplate code`
	__Product__00000014_.ComputedPrefix = `1.2.2`
	__Product__00000014_.IsExpanded = false
	__Product__00000014_.LayoutDirection = models.Vertical
	__Product__00000014_.IsImport = false
	__Product__00000014_.Description = ``
	__Product__00000014_.IsProducersNodeExpanded = false
	__Product__00000014_.IsConsumersNodeExpanded = false

	__ProductCompositionShape__00000000_.Name = `Paint points to `
	__ProductCompositionShape__00000000_.StartRatio = 0.500000
	__ProductCompositionShape__00000000_.EndRatio = 0.500000
	__ProductCompositionShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000000_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000000_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000000_.IsHidden = false

	__ProductCompositionShape__00000001_.Name = `Paint points to `
	__ProductCompositionShape__00000001_.StartRatio = 0.500000
	__ProductCompositionShape__00000001_.EndRatio = 0.500000
	__ProductCompositionShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000001_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000001_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000001_.IsHidden = false

	__ProductCompositionShape__00000002_.Name = `Paint points to `
	__ProductCompositionShape__00000002_.StartRatio = 0.500000
	__ProductCompositionShape__00000002_.EndRatio = 0.500000
	__ProductCompositionShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000002_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000002_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000002_.IsHidden = false

	__ProductCompositionShape__00000003_.Name = `Paint points to `
	__ProductCompositionShape__00000003_.StartRatio = 0.500000
	__ProductCompositionShape__00000003_.EndRatio = 0.500000
	__ProductCompositionShape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000003_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000003_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000003_.IsHidden = false

	__ProductCompositionShape__00000004_.Name = `the stager concept is not clearly delineated to `
	__ProductCompositionShape__00000004_.StartRatio = 0.500000
	__ProductCompositionShape__00000004_.EndRatio = 0.500000
	__ProductCompositionShape__00000004_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000004_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000004_.IsHidden = false

	__ProductCompositionShape__00000005_.Name = `the stager concept is not clearly delineated to `
	__ProductCompositionShape__00000005_.StartRatio = 0.500000
	__ProductCompositionShape__00000005_.EndRatio = 0.500000
	__ProductCompositionShape__00000005_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000005_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000005_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000005_.IsHidden = false

	__ProductCompositionShape__00000006_.Name = `lot of yyy files and no clear mental model  to `
	__ProductCompositionShape__00000006_.StartRatio = 0.500000
	__ProductCompositionShape__00000006_.EndRatio = 0.500000
	__ProductCompositionShape__00000006_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000006_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000006_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000006_.IsHidden = false

	__ProductCompositionShape__00000007_.Name = `yyy files are helpers for UX updates of trees and svg to `
	__ProductCompositionShape__00000007_.StartRatio = 0.500000
	__ProductCompositionShape__00000007_.EndRatio = 0.500000
	__ProductCompositionShape__00000007_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000007_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000007_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000007_.IsHidden = false

	__ProductCompositionShape__00000008_.Name = `yyy files are helpers for UX updates of trees and svg to `
	__ProductCompositionShape__00000008_.StartRatio = 0.500000
	__ProductCompositionShape__00000008_.EndRatio = 0.500000
	__ProductCompositionShape__00000008_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000008_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000008_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000008_.IsHidden = false

	__ProductCompositionShape__00000009_.Name = `yyy files are helpers for UX updates of trees and svg to `
	__ProductCompositionShape__00000009_.StartRatio = 0.500000
	__ProductCompositionShape__00000009_.EndRatio = 0.500000
	__ProductCompositionShape__00000009_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000009_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000009_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000009_.IsHidden = false

	__ProductCompositionShape__00000010_.Name = `yyy files are helpers for UX updates of trees and svg to `
	__ProductCompositionShape__00000010_.StartRatio = 0.500000
	__ProductCompositionShape__00000010_.EndRatio = 0.500000
	__ProductCompositionShape__00000010_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000010_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000010_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000010_.IsHidden = false

	__ProductCompositionShape__00000011_.Name = `yyy files are helpers for UX updates of trees and svg to `
	__ProductCompositionShape__00000011_.StartRatio = 0.500000
	__ProductCompositionShape__00000011_.EndRatio = 0.500000
	__ProductCompositionShape__00000011_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000011_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000011_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000011_.IsHidden = false

	__ProductCompositionShape__00000012_.Name = `lot of boilerplate code for stager to `
	__ProductCompositionShape__00000012_.StartRatio = 0.500000
	__ProductCompositionShape__00000012_.EndRatio = 0.500000
	__ProductCompositionShape__00000012_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000012_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000012_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000012_.IsHidden = false

	__ProductCompositionShape__00000013_.Name = `lot of boilerplate code for stager. stager is a hybrid concept to `
	__ProductCompositionShape__00000013_.StartRatio = 0.500000
	__ProductCompositionShape__00000013_.EndRatio = 0.500000
	__ProductCompositionShape__00000013_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000013_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000013_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000013_.IsHidden = false

	__ProductShape__00000000_.Name = `-Default Diagram`
	__ProductShape__00000000_.OverideLayoutDirection = false
	__ProductShape__00000000_.LayoutDirection = models.Vertical
	__ProductShape__00000000_.X = 50.000000
	__ProductShape__00000000_.Y = 50.000000
	__ProductShape__00000000_.Width = 250.000000
	__ProductShape__00000000_.Height = 70.000000
	__ProductShape__00000000_.IsHidden = false

	__ProductShape__00000001_.Name = `-Default Diagram`
	__ProductShape__00000001_.OverideLayoutDirection = false
	__ProductShape__00000001_.LayoutDirection = models.Vertical
	__ProductShape__00000001_.X = 225.000000
	__ProductShape__00000001_.Y = 135.000000
	__ProductShape__00000001_.Width = 250.000000
	__ProductShape__00000001_.Height = 70.000000
	__ProductShape__00000001_.IsHidden = false

	__ProductShape__00000002_.Name = `-Default Diagram`
	__ProductShape__00000002_.OverideLayoutDirection = false
	__ProductShape__00000002_.LayoutDirection = models.Vertical
	__ProductShape__00000002_.X = 225.000000
	__ProductShape__00000002_.Y = 785.000000
	__ProductShape__00000002_.Width = 250.000000
	__ProductShape__00000002_.Height = 70.000000
	__ProductShape__00000002_.IsHidden = false

	__ProductShape__00000003_.Name = `-Default Diagram`
	__ProductShape__00000003_.OverideLayoutDirection = false
	__ProductShape__00000003_.LayoutDirection = models.Vertical
	__ProductShape__00000003_.X = 225.000000
	__ProductShape__00000003_.Y = 1010.000000
	__ProductShape__00000003_.Width = 250.000000
	__ProductShape__00000003_.Height = 70.000000
	__ProductShape__00000003_.IsHidden = false

	__ProductShape__00000004_.Name = `-Default Diagram`
	__ProductShape__00000004_.OverideLayoutDirection = false
	__ProductShape__00000004_.LayoutDirection = models.Vertical
	__ProductShape__00000004_.X = 225.000000
	__ProductShape__00000004_.Y = 1235.000000
	__ProductShape__00000004_.Width = 250.000000
	__ProductShape__00000004_.Height = 70.000000
	__ProductShape__00000004_.IsHidden = false

	__ProductShape__00000005_.Name = `-Default Diagram`
	__ProductShape__00000005_.OverideLayoutDirection = false
	__ProductShape__00000005_.LayoutDirection = models.Vertical
	__ProductShape__00000005_.X = 400.000000
	__ProductShape__00000005_.Y = 1150.000000
	__ProductShape__00000005_.Width = 250.000000
	__ProductShape__00000005_.Height = 70.000000
	__ProductShape__00000005_.IsHidden = false

	__ProductShape__00000006_.Name = `-Default Diagram`
	__ProductShape__00000006_.OverideLayoutDirection = false
	__ProductShape__00000006_.LayoutDirection = models.Vertical
	__ProductShape__00000006_.X = 700.000000
	__ProductShape__00000006_.Y = 1150.000000
	__ProductShape__00000006_.Width = 250.000000
	__ProductShape__00000006_.Height = 70.000000
	__ProductShape__00000006_.IsHidden = false

	__ProductShape__00000007_.Name = `-Default Diagram`
	__ProductShape__00000007_.OverideLayoutDirection = false
	__ProductShape__00000007_.LayoutDirection = models.Vertical
	__ProductShape__00000007_.X = 400.000000
	__ProductShape__00000007_.Y = 275.000000
	__ProductShape__00000007_.Width = 250.000000
	__ProductShape__00000007_.Height = 70.000000
	__ProductShape__00000007_.IsHidden = false

	__ProductShape__00000008_.Name = `-Default Diagram`
	__ProductShape__00000008_.OverideLayoutDirection = false
	__ProductShape__00000008_.LayoutDirection = models.Vertical
	__ProductShape__00000008_.X = 575.000000
	__ProductShape__00000008_.Y = 360.000000
	__ProductShape__00000008_.Width = 250.000000
	__ProductShape__00000008_.Height = 70.000000
	__ProductShape__00000008_.IsHidden = false

	__ProductShape__00000009_.Name = `-Default Diagram`
	__ProductShape__00000009_.OverideLayoutDirection = false
	__ProductShape__00000009_.LayoutDirection = models.Vertical
	__ProductShape__00000009_.X = 575.000000
	__ProductShape__00000009_.Y = 445.000000
	__ProductShape__00000009_.Width = 250.000000
	__ProductShape__00000009_.Height = 70.000000
	__ProductShape__00000009_.IsHidden = false

	__ProductShape__00000010_.Name = `-Default Diagram`
	__ProductShape__00000010_.OverideLayoutDirection = false
	__ProductShape__00000010_.LayoutDirection = models.Vertical
	__ProductShape__00000010_.X = 575.000000
	__ProductShape__00000010_.Y = 530.000000
	__ProductShape__00000010_.Width = 250.000000
	__ProductShape__00000010_.Height = 70.000000
	__ProductShape__00000010_.IsHidden = false

	__ProductShape__00000011_.Name = `-Default Diagram`
	__ProductShape__00000011_.OverideLayoutDirection = false
	__ProductShape__00000011_.LayoutDirection = models.Vertical
	__ProductShape__00000011_.X = 575.000000
	__ProductShape__00000011_.Y = 615.000000
	__ProductShape__00000011_.Width = 250.000000
	__ProductShape__00000011_.Height = 70.000000
	__ProductShape__00000011_.IsHidden = false

	__ProductShape__00000012_.Name = `-Default Diagram`
	__ProductShape__00000012_.OverideLayoutDirection = false
	__ProductShape__00000012_.LayoutDirection = models.Vertical
	__ProductShape__00000012_.X = 575.000000
	__ProductShape__00000012_.Y = 700.000000
	__ProductShape__00000012_.Width = 250.000000
	__ProductShape__00000012_.Height = 70.000000
	__ProductShape__00000012_.IsHidden = false

	__ProductShape__00000013_.Name = `-Default Diagram`
	__ProductShape__00000013_.OverideLayoutDirection = false
	__ProductShape__00000013_.LayoutDirection = models.Vertical
	__ProductShape__00000013_.X = 400.000000
	__ProductShape__00000013_.Y = 925.000000
	__ProductShape__00000013_.Width = 250.000000
	__ProductShape__00000013_.Height = 70.000000
	__ProductShape__00000013_.IsHidden = false

	__ProductShape__00000014_.Name = `-Default Diagram`
	__ProductShape__00000014_.OverideLayoutDirection = false
	__ProductShape__00000014_.LayoutDirection = models.Vertical
	__ProductShape__00000014_.X = 700.000000
	__ProductShape__00000014_.Y = 925.000000
	__ProductShape__00000014_.Width = 250.000000
	__ProductShape__00000014_.Height = 70.000000
	__ProductShape__00000014_.IsHidden = false

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
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, __Product__00000000_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, __Product__00000003_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, __Product__00000001_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, __Product__00000007_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, __Product__00000002_)
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
	__Library__00000000_.RootProducts = append(__Library__00000000_.RootProducts, __Product__00000000_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000000_)
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000000_)
	__Note__00000000_.Products = append(__Note__00000000_.Products, __Product__00000004_)
	__NoteProductShape__00000000_.Note = __Note__00000000_
	__NoteProductShape__00000000_.Product = __Product__00000004_
	__NoteShape__00000000_.Note = __Note__00000000_
	__Product__00000000_.ReferencedProduct = nil
	__Product__00000000_.SubProducts = append(__Product__00000000_.SubProducts, __Product__00000001_)
	__Product__00000000_.SubProducts = append(__Product__00000000_.SubProducts, __Product__00000002_)
	__Product__00000000_.SubProducts = append(__Product__00000000_.SubProducts, __Product__00000003_)
	__Product__00000000_.SubProducts = append(__Product__00000000_.SubProducts, __Product__00000004_)
	__Product__00000001_.ReferencedProduct = nil
	__Product__00000001_.SubProducts = append(__Product__00000001_.SubProducts, __Product__00000007_)
	__Product__00000002_.ReferencedProduct = nil
	__Product__00000002_.SubProducts = append(__Product__00000002_.SubProducts, __Product__00000013_)
	__Product__00000002_.SubProducts = append(__Product__00000002_.SubProducts, __Product__00000014_)
	__Product__00000003_.ReferencedProduct = nil
	__Product__00000003_.SubProducts = append(__Product__00000003_.SubProducts, __Product__00000005_)
	__Product__00000003_.SubProducts = append(__Product__00000003_.SubProducts, __Product__00000006_)
	__Product__00000004_.ReferencedProduct = nil
	__Product__00000005_.ReferencedProduct = nil
	__Product__00000006_.ReferencedProduct = nil
	__Product__00000007_.ReferencedProduct = nil
	__Product__00000007_.SubProducts = append(__Product__00000007_.SubProducts, __Product__00000008_)
	__Product__00000007_.SubProducts = append(__Product__00000007_.SubProducts, __Product__00000009_)
	__Product__00000007_.SubProducts = append(__Product__00000007_.SubProducts, __Product__00000010_)
	__Product__00000007_.SubProducts = append(__Product__00000007_.SubProducts, __Product__00000011_)
	__Product__00000007_.SubProducts = append(__Product__00000007_.SubProducts, __Product__00000012_)
	__Product__00000008_.ReferencedProduct = nil
	__Product__00000009_.ReferencedProduct = nil
	__Product__00000010_.ReferencedProduct = nil
	__Product__00000011_.ReferencedProduct = nil
	__Product__00000012_.ReferencedProduct = nil
	__Product__00000013_.ReferencedProduct = nil
	__Product__00000014_.ReferencedProduct = nil
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
}
