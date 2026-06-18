package diagrams

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/dsm/scenario/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var (
	_ time.Time
	_ = slices.Index[[]int, int]
)

// _ point for meta package dummy declaration
var _ ref_models.Stage

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__AttributeShape__00000000_ := (&models.AttributeShape{Name: `Description`}).Stage(stage)
	__AttributeShape__00000001_ := (&models.AttributeShape{Name: `IsWithProbaility`}).Stage(stage)
	__AttributeShape__00000002_ := (&models.AttributeShape{Name: `Probability`}).Stage(stage)
	__AttributeShape__00000004_ := (&models.AttributeShape{Name: `Start`}).Stage(stage)
	__AttributeShape__00000005_ := (&models.AttributeShape{Name: `End`}).Stage(stage)
	__AttributeShape__00000006_ := (&models.AttributeShape{Name: `NumberOfYearsBetweenTicks`}).Stage(stage)
	__AttributeShape__00000007_ := (&models.AttributeShape{Name: `HorizontalAxis_Right_X`}).Stage(stage)
	__AttributeShape__00000008_ := (&models.AttributeShape{Name: `VerticalAxis_StrokeWidth`}).Stage(stage)
	__AttributeShape__00000009_ := (&models.AttributeShape{Name: `VerticalAxis_Bottom_Y`}).Stage(stage)
	__AttributeShape__00000010_ := (&models.AttributeShape{Name: `VerticalAxis_Top_Y`}).Stage(stage)
	__AttributeShape__00000011_ := (&models.AttributeShape{Name: `AxisOrign_Y`}).Stage(stage)
	__AttributeShape__00000012_ := (&models.AttributeShape{Name: `AxisOrign_X`}).Stage(stage)

	__Classdiagram__00000000_ := (&models.Classdiagram{Name: `Abstract core`}).Stage(stage)
	__Classdiagram__00000001_ := (&models.Classdiagram{Name: `Concrete core`}).Stage(stage)

	__DiagramPackage__00000000_ := (&models.DiagramPackage{Name: `Diagram Package created the 2026-06-02T15:48:46Z`}).Stage(stage)

	__GongStructShape__00000001_ := (&models.GongStructShape{Name: `Abstract core-ActorState`}).Stage(stage)
	__GongStructShape__00000002_ := (&models.GongStructShape{Name: `Abstract core-ActorStateTransition`}).Stage(stage)
	__GongStructShape__00000003_ := (&models.GongStructShape{Name: `Abstract core-Analysis`}).Stage(stage)
	__GongStructShape__00000004_ := (&models.GongStructShape{Name: `Abstract core-Scenario`}).Stage(stage)
	__GongStructShape__00000006_ := (&models.GongStructShape{Name: `Abstract core-Parameter`}).Stage(stage)
	__GongStructShape__00000009_ := (&models.GongStructShape{Name: `Abstract core-EvolutionDirection`}).Stage(stage)
	__GongStructShape__00000010_ := (&models.GongStructShape{Name: `Concrete core-Diagram`}).Stage(stage)
	__GongStructShape__00000011_ := (&models.GongStructShape{Name: `Concrete core-ActorStateShape`}).Stage(stage)
	__GongStructShape__00000012_ := (&models.GongStructShape{Name: `Concrete core-ActorStateTransitionShape`}).Stage(stage)
	__GongStructShape__00000013_ := (&models.GongStructShape{Name: `Concrete core-EvolutionDirectionShape`}).Stage(stage)
	__GongStructShape__00000014_ := (&models.GongStructShape{Name: `Concrete core-ParameterShape`}).Stage(stage)
	__GongStructShape__00000015_ := (&models.GongStructShape{Name: `Concrete core-ParametersAggregateShape`}).Stage(stage)
	__GongStructShape__00000017_ := (&models.GongStructShape{Name: `Abstract core-ParametersAggregate`}).Stage(stage)

	__LinkShape__00000000_ := (&models.LinkShape{Name: `StartState`}).Stage(stage)
	__LinkShape__00000001_ := (&models.LinkShape{Name: `EndState`}).Stage(stage)
	__LinkShape__00000002_ := (&models.LinkShape{Name: `Scenarios`}).Stage(stage)
	__LinkShape__00000003_ := (&models.LinkShape{Name: `ActorStates`}).Stage(stage)
	__LinkShape__00000004_ := (&models.LinkShape{Name: `ActorStateTransitions`}).Stage(stage)
	__LinkShape__00000005_ := (&models.LinkShape{Name: `Parameters`}).Stage(stage)
	__LinkShape__00000006_ := (&models.LinkShape{Name: `EvolutionDirections`}).Stage(stage)
	__LinkShape__00000007_ := (&models.LinkShape{Name: `ParametersAggretates`}).Stage(stage)
	__LinkShape__00000008_ := (&models.LinkShape{Name: `Parameters`}).Stage(stage)
	__LinkShape__00000009_ := (&models.LinkShape{Name: `Justifications`}).Stage(stage)
	__LinkShape__00000010_ := (&models.LinkShape{Name: `ActorStateShapes`}).Stage(stage)
	__LinkShape__00000011_ := (&models.LinkShape{Name: `ParameterShapes`}).Stage(stage)
	__LinkShape__00000012_ := (&models.LinkShape{Name: `ActorStateTransitionShapes`}).Stage(stage)
	__LinkShape__00000013_ := (&models.LinkShape{Name: `ScenarioParameterShapes`}).Stage(stage)
	__LinkShape__00000014_ := (&models.LinkShape{Name: `EvolutionDirectionShapes`}).Stage(stage)

	// insertion point for initialization of values

	__AttributeShape__00000000_.Name = `Description`
	__AttributeShape__00000000_.IdentifierMeta = ref_models.ActorState{}.Description
	__AttributeShape__00000000_.FieldTypeAsString = ``
	__AttributeShape__00000000_.Structname = `ActorState`
	__AttributeShape__00000000_.Fieldtypename = `string`

	__AttributeShape__00000001_.Name = `IsWithProbaility`
	__AttributeShape__00000001_.IdentifierMeta = ref_models.ActorState{}.IsWithProbaility
	__AttributeShape__00000001_.FieldTypeAsString = ``
	__AttributeShape__00000001_.Structname = `ActorState`
	__AttributeShape__00000001_.Fieldtypename = `bool`

	__AttributeShape__00000002_.Name = `Probability`
	__AttributeShape__00000002_.IdentifierMeta = ref_models.ActorState{}.Probability
	__AttributeShape__00000002_.FieldTypeAsString = ``
	__AttributeShape__00000002_.Structname = `ActorState`
	__AttributeShape__00000002_.Fieldtypename = `ProbabilityEnum`

	__AttributeShape__00000004_.Name = `Start`
	__AttributeShape__00000004_.IdentifierMeta = ref_models.Diagram{}.Start
	__AttributeShape__00000004_.FieldTypeAsString = ``
	__AttributeShape__00000004_.Structname = `Diagram`
	__AttributeShape__00000004_.Fieldtypename = `Time`

	__AttributeShape__00000005_.Name = `End`
	__AttributeShape__00000005_.IdentifierMeta = ref_models.Diagram{}.End
	__AttributeShape__00000005_.FieldTypeAsString = ``
	__AttributeShape__00000005_.Structname = `Diagram`
	__AttributeShape__00000005_.Fieldtypename = `Time`

	__AttributeShape__00000006_.Name = `NumberOfYearsBetweenTicks`
	__AttributeShape__00000006_.IdentifierMeta = ref_models.Diagram{}.NumberOfYearsBetweenTicks
	__AttributeShape__00000006_.FieldTypeAsString = ``
	__AttributeShape__00000006_.Structname = `Diagram`
	__AttributeShape__00000006_.Fieldtypename = `int`

	__AttributeShape__00000007_.Name = `HorizontalAxis_Right_X`
	__AttributeShape__00000007_.IdentifierMeta = ref_models.Diagram{}.HorizontalAxis_Right_X
	__AttributeShape__00000007_.FieldTypeAsString = ``
	__AttributeShape__00000007_.Structname = `Diagram`
	__AttributeShape__00000007_.Fieldtypename = `float64`

	__AttributeShape__00000008_.Name = `VerticalAxis_StrokeWidth`
	__AttributeShape__00000008_.IdentifierMeta = ref_models.Diagram{}.VerticalAxis_StrokeWidth
	__AttributeShape__00000008_.FieldTypeAsString = ``
	__AttributeShape__00000008_.Structname = `Diagram`
	__AttributeShape__00000008_.Fieldtypename = `float64`

	__AttributeShape__00000009_.Name = `VerticalAxis_Bottom_Y`
	__AttributeShape__00000009_.IdentifierMeta = ref_models.Diagram{}.VerticalAxis_Bottom_Y
	__AttributeShape__00000009_.FieldTypeAsString = ``
	__AttributeShape__00000009_.Structname = `Diagram`
	__AttributeShape__00000009_.Fieldtypename = `float64`

	__AttributeShape__00000010_.Name = `VerticalAxis_Top_Y`
	__AttributeShape__00000010_.IdentifierMeta = ref_models.Diagram{}.VerticalAxis_Top_Y
	__AttributeShape__00000010_.FieldTypeAsString = ``
	__AttributeShape__00000010_.Structname = `Diagram`
	__AttributeShape__00000010_.Fieldtypename = `float64`

	__AttributeShape__00000011_.Name = `AxisOrign_Y`
	__AttributeShape__00000011_.IdentifierMeta = ref_models.Diagram{}.AxisOrign_Y
	__AttributeShape__00000011_.FieldTypeAsString = ``
	__AttributeShape__00000011_.Structname = `Diagram`
	__AttributeShape__00000011_.Fieldtypename = `float64`

	__AttributeShape__00000012_.Name = `AxisOrign_X`
	__AttributeShape__00000012_.IdentifierMeta = ref_models.Diagram{}.AxisOrign_X
	__AttributeShape__00000012_.FieldTypeAsString = ``
	__AttributeShape__00000012_.Structname = `Diagram`
	__AttributeShape__00000012_.Fieldtypename = `float64`

	__Classdiagram__00000000_.Name = `Abstract core`
	__Classdiagram__00000000_.Description = ``
	__Classdiagram__00000000_.IsIncludedInStaticWebSite = false
	__Classdiagram__00000000_.ShowNbInstances = true
	__Classdiagram__00000000_.ShowMultiplicity = true
	__Classdiagram__00000000_.ShowLinkNames = true
	__Classdiagram__00000000_.IsInRenameMode = false
	__Classdiagram__00000000_.IsExpanded = false
	__Classdiagram__00000000_.NodeGongStructsIsExpanded = true
	__Classdiagram__00000000_.NodeGongStructNodeExpansion = `[true,false,true,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false]`
	__Classdiagram__00000000_.NodeGongEnumsIsExpanded = false
	__Classdiagram__00000000_.NodeGongEnumNodeExpansion = ``
	__Classdiagram__00000000_.NodeGongNotesIsExpanded = false
	__Classdiagram__00000000_.NodeGongNoteNodeExpansion = ``

	__Classdiagram__00000001_.Name = `Concrete core`
	__Classdiagram__00000001_.Description = ``
	__Classdiagram__00000001_.IsIncludedInStaticWebSite = false
	__Classdiagram__00000001_.ShowNbInstances = true
	__Classdiagram__00000001_.ShowMultiplicity = true
	__Classdiagram__00000001_.ShowLinkNames = true
	__Classdiagram__00000001_.IsInRenameMode = false
	__Classdiagram__00000001_.IsExpanded = true
	__Classdiagram__00000001_.NodeGongStructsIsExpanded = true
	__Classdiagram__00000001_.NodeGongStructNodeExpansion = `[false,false,false,false,false,false]`
	__Classdiagram__00000001_.NodeGongEnumsIsExpanded = false
	__Classdiagram__00000001_.NodeGongEnumNodeExpansion = ``
	__Classdiagram__00000001_.NodeGongNotesIsExpanded = false
	__Classdiagram__00000001_.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__00000000_.Name = `Diagram Package created the 2026-06-02T15:48:46Z`
	__DiagramPackage__00000000_.Path = ``
	__DiagramPackage__00000000_.GongModelPath = ``
	__DiagramPackage__00000000_.AbsolutePathToDiagramPackage = ``

	__GongStructShape__00000001_.Name = `Abstract core-ActorState`
	__GongStructShape__00000001_.X = 905.999939
	__GongStructShape__00000001_.Y = 308.000000
	__GongStructShape__00000001_.IdentifierMeta = ref_models.ActorState{}
	__GongStructShape__00000001_.Width = 240.000000
	__GongStructShape__00000001_.Height = 252.000000
	__GongStructShape__00000001_.IsSelected = false

	__GongStructShape__00000002_.Name = `Abstract core-ActorStateTransition`
	__GongStructShape__00000002_.X = 1407.999939
	__GongStructShape__00000002_.Y = 311.000000
	__GongStructShape__00000002_.IdentifierMeta = ref_models.ActorStateTransition{}
	__GongStructShape__00000002_.Width = 240.000000
	__GongStructShape__00000002_.Height = 240.999969
	__GongStructShape__00000002_.IsSelected = false

	__GongStructShape__00000003_.Name = `Abstract core-Analysis`
	__GongStructShape__00000003_.X = 23.000000
	__GongStructShape__00000003_.Y = 50.000000
	__GongStructShape__00000003_.IdentifierMeta = ref_models.Analysis{}
	__GongStructShape__00000003_.Width = 240.000000
	__GongStructShape__00000003_.Height = 63.000000
	__GongStructShape__00000003_.IsSelected = false

	__GongStructShape__00000004_.Name = `Abstract core-Scenario`
	__GongStructShape__00000004_.X = 652.000000
	__GongStructShape__00000004_.Y = 52.000000
	__GongStructShape__00000004_.IdentifierMeta = ref_models.Scenario{}
	__GongStructShape__00000004_.Width = 240.000000
	__GongStructShape__00000004_.Height = 63.000000
	__GongStructShape__00000004_.IsSelected = false

	__GongStructShape__00000006_.Name = `Abstract core-Parameter`
	__GongStructShape__00000006_.X = 332.000000
	__GongStructShape__00000006_.Y = 614.999969
	__GongStructShape__00000006_.IdentifierMeta = ref_models.Parameter{}
	__GongStructShape__00000006_.Width = 240.000000
	__GongStructShape__00000006_.Height = 63.000000
	__GongStructShape__00000006_.IsSelected = false

	__GongStructShape__00000009_.Name = `Abstract core-EvolutionDirection`
	__GongStructShape__00000009_.X = 576.000000
	__GongStructShape__00000009_.Y = 327.000000
	__GongStructShape__00000009_.IdentifierMeta = ref_models.EvolutionDirection{}
	__GongStructShape__00000009_.Width = 240.000000
	__GongStructShape__00000009_.Height = 63.000000
	__GongStructShape__00000009_.IsSelected = false

	__GongStructShape__00000010_.Name = `Concrete core-Diagram`
	__GongStructShape__00000010_.X = 81.000000
	__GongStructShape__00000010_.Y = 38.000000
	__GongStructShape__00000010_.IdentifierMeta = ref_models.Diagram{}
	__GongStructShape__00000010_.Width = 362.000000
	__GongStructShape__00000010_.Height = 686.999969
	__GongStructShape__00000010_.IsSelected = false

	__GongStructShape__00000011_.Name = `Concrete core-ActorStateShape`
	__GongStructShape__00000011_.X = 785.000000
	__GongStructShape__00000011_.Y = 52.000000
	__GongStructShape__00000011_.IdentifierMeta = ref_models.ActorStateShape{}
	__GongStructShape__00000011_.Width = 346.999939
	__GongStructShape__00000011_.Height = 63.000000
	__GongStructShape__00000011_.IsSelected = false

	__GongStructShape__00000012_.Name = `Concrete core-ActorStateTransitionShape`
	__GongStructShape__00000012_.X = 787.000000
	__GongStructShape__00000012_.Y = 189.000000
	__GongStructShape__00000012_.IdentifierMeta = ref_models.ActorStateTransitionShape{}
	__GongStructShape__00000012_.Width = 346.000000
	__GongStructShape__00000012_.Height = 63.000000
	__GongStructShape__00000012_.IsSelected = false

	__GongStructShape__00000013_.Name = `Concrete core-EvolutionDirectionShape`
	__GongStructShape__00000013_.X = 791.000000
	__GongStructShape__00000013_.Y = 326.000000
	__GongStructShape__00000013_.IdentifierMeta = ref_models.EvolutionDirectionShape{}
	__GongStructShape__00000013_.Width = 346.000000
	__GongStructShape__00000013_.Height = 63.000000
	__GongStructShape__00000013_.IsSelected = false

	__GongStructShape__00000014_.Name = `Concrete core-ParameterShape`
	__GongStructShape__00000014_.X = 791.000000
	__GongStructShape__00000014_.Y = 486.999969
	__GongStructShape__00000014_.IdentifierMeta = ref_models.ParameterShape{}
	__GongStructShape__00000014_.Width = 345.000000
	__GongStructShape__00000014_.Height = 63.000000
	__GongStructShape__00000014_.IsSelected = false

	__GongStructShape__00000015_.Name = `Concrete core-ParametersAggregateShape`
	__GongStructShape__00000015_.X = 794.000000
	__GongStructShape__00000015_.Y = 650.999969
	__GongStructShape__00000015_.IdentifierMeta = ref_models.ParametersAggregateShape{}
	__GongStructShape__00000015_.Width = 347.000000
	__GongStructShape__00000015_.Height = 63.000000
	__GongStructShape__00000015_.IsSelected = false

	__GongStructShape__00000017_.Name = `Abstract core-ParametersAggregate`
	__GongStructShape__00000017_.X = 19.000000
	__GongStructShape__00000017_.Y = 329.000000
	__GongStructShape__00000017_.IdentifierMeta = ref_models.ParametersAggregate{}
	__GongStructShape__00000017_.Width = 240.000000
	__GongStructShape__00000017_.Height = 63.000000
	__GongStructShape__00000017_.IsSelected = false

	__LinkShape__00000000_.Name = `StartState`
	__LinkShape__00000000_.IdentifierMeta = ref_models.ActorStateTransition{}.StartState
	__LinkShape__00000000_.FieldTypeIdentifierMeta = ref_models.ActorState{}
	__LinkShape__00000000_.FieldOffsetX = 0.000000
	__LinkShape__00000000_.FieldOffsetY = 0.000000
	__LinkShape__00000000_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000000_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000000_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000000_.SourceMultiplicity = models.MANY
	__LinkShape__00000000_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000000_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000000_.X = 739.000000
	__LinkShape__00000000_.Y = 152.500000
	__LinkShape__00000000_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000000_.StartRatio = 0.179167
	__LinkShape__00000000_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000000_.EndRatio = 0.228836
	__LinkShape__00000000_.CornerOffsetRatio = -0.455903

	__LinkShape__00000001_.Name = `EndState`
	__LinkShape__00000001_.IdentifierMeta = ref_models.ActorStateTransition{}.EndState
	__LinkShape__00000001_.FieldTypeIdentifierMeta = ref_models.ActorState{}
	__LinkShape__00000001_.FieldOffsetX = 0.000000
	__LinkShape__00000001_.FieldOffsetY = 0.000000
	__LinkShape__00000001_.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__00000001_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000001_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000001_.SourceMultiplicity = models.MANY
	__LinkShape__00000001_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000001_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000001_.X = 739.000000
	__LinkShape__00000001_.Y = 152.500000
	__LinkShape__00000001_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000001_.StartRatio = 0.803596
	__LinkShape__00000001_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000001_.EndRatio = 0.748677
	__LinkShape__00000001_.CornerOffsetRatio = -0.551736

	__LinkShape__00000002_.Name = `Scenarios`
	__LinkShape__00000002_.IdentifierMeta = ref_models.Analysis{}.Scenarios
	__LinkShape__00000002_.FieldTypeIdentifierMeta = ref_models.Scenario{}
	__LinkShape__00000002_.FieldOffsetX = 0.000000
	__LinkShape__00000002_.FieldOffsetY = 0.000000
	__LinkShape__00000002_.TargetMultiplicity = models.MANY
	__LinkShape__00000002_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000002_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000002_.SourceMultiplicity = models.MANY
	__LinkShape__00000002_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000002_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000002_.X = 721.500000
	__LinkShape__00000002_.Y = 83.500000
	__LinkShape__00000002_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000002_.StartRatio = 0.500000
	__LinkShape__00000002_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000002_.EndRatio = 0.500000
	__LinkShape__00000002_.CornerOffsetRatio = 1.380000

	__LinkShape__00000003_.Name = `ActorStates`
	__LinkShape__00000003_.IdentifierMeta = ref_models.Scenario{}.ActorStates
	__LinkShape__00000003_.FieldTypeIdentifierMeta = ref_models.ActorState{}
	__LinkShape__00000003_.FieldOffsetX = 0.000000
	__LinkShape__00000003_.FieldOffsetY = 0.000000
	__LinkShape__00000003_.TargetMultiplicity = models.MANY
	__LinkShape__00000003_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000003_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000003_.SourceMultiplicity = models.MANY
	__LinkShape__00000003_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000003_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000003_.X = 1176.999969
	__LinkShape__00000003_.Y = 208.500000
	__LinkShape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000003_.StartRatio = 0.673264
	__LinkShape__00000003_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000003_.EndRatio = 0.623264
	__LinkShape__00000003_.CornerOffsetRatio = 2.201058

	__LinkShape__00000004_.Name = `ActorStateTransitions`
	__LinkShape__00000004_.IdentifierMeta = ref_models.Scenario{}.ActorStateTransitions
	__LinkShape__00000004_.FieldTypeIdentifierMeta = ref_models.ActorStateTransition{}
	__LinkShape__00000004_.FieldOffsetX = 0.000000
	__LinkShape__00000004_.FieldOffsetY = 0.000000
	__LinkShape__00000004_.TargetMultiplicity = models.MANY
	__LinkShape__00000004_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000004_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000004_.SourceMultiplicity = models.MANY
	__LinkShape__00000004_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000004_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000004_.X = 896.000000
	__LinkShape__00000004_.Y = 209.000000
	__LinkShape__00000004_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000004_.StartRatio = 0.877430
	__LinkShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000004_.EndRatio = 0.823264
	__LinkShape__00000004_.CornerOffsetRatio = 1.423280

	__LinkShape__00000005_.Name = `Parameters`
	__LinkShape__00000005_.IdentifierMeta = ref_models.Scenario{}.Parameters
	__LinkShape__00000005_.FieldTypeIdentifierMeta = ref_models.Parameter{}
	__LinkShape__00000005_.FieldOffsetX = 0.000000
	__LinkShape__00000005_.FieldOffsetY = 0.000000
	__LinkShape__00000005_.TargetMultiplicity = models.MANY
	__LinkShape__00000005_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000005_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000005_.SourceMultiplicity = models.MANY
	__LinkShape__00000005_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000005_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000005_.X = 706.000000
	__LinkShape__00000005_.Y = 205.500000
	__LinkShape__00000005_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000005_.StartRatio = 0.331597
	__LinkShape__00000005_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000005_.EndRatio = 0.460764
	__LinkShape__00000005_.CornerOffsetRatio = 2.201058

	__LinkShape__00000006_.Name = `EvolutionDirections`
	__LinkShape__00000006_.IdentifierMeta = ref_models.Scenario{}.EvolutionDirections
	__LinkShape__00000006_.FieldTypeIdentifierMeta = ref_models.EvolutionDirection{}
	__LinkShape__00000006_.FieldOffsetX = 0.000000
	__LinkShape__00000006_.FieldOffsetY = 0.000000
	__LinkShape__00000006_.TargetMultiplicity = models.MANY
	__LinkShape__00000006_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000006_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000006_.SourceMultiplicity = models.MANY
	__LinkShape__00000006_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000006_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000006_.X = 937.000000
	__LinkShape__00000006_.Y = 211.000000
	__LinkShape__00000006_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000006_.StartRatio = 0.464930
	__LinkShape__00000006_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000006_.EndRatio = 0.773264
	__LinkShape__00000006_.CornerOffsetRatio = 1.804233

	__LinkShape__00000007_.Name = `ParametersAggretates`
	__LinkShape__00000007_.IdentifierMeta = ref_models.Scenario{}.ParametersAggretates
	__LinkShape__00000007_.FieldTypeIdentifierMeta = ref_models.ParametersAggregate{}
	__LinkShape__00000007_.FieldOffsetX = 0.000000
	__LinkShape__00000007_.FieldOffsetY = 0.000000
	__LinkShape__00000007_.TargetMultiplicity = models.MANY
	__LinkShape__00000007_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000007_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000007_.SourceMultiplicity = models.MANY
	__LinkShape__00000007_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000007_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000007_.X = 746.000000
	__LinkShape__00000007_.Y = 312.999985
	__LinkShape__00000007_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000007_.StartRatio = 0.135764
	__LinkShape__00000007_.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000007_.EndRatio = 0.777430
	__LinkShape__00000007_.CornerOffsetRatio = 1.804233

	__LinkShape__00000008_.Name = `Parameters`
	__LinkShape__00000008_.IdentifierMeta = ref_models.ParametersAggregate{}.Parameters
	__LinkShape__00000008_.FieldTypeIdentifierMeta = ref_models.Parameter{}
	__LinkShape__00000008_.FieldOffsetX = 0.000000
	__LinkShape__00000008_.FieldOffsetY = 0.000000
	__LinkShape__00000008_.TargetMultiplicity = models.MANY
	__LinkShape__00000008_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000008_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000008_.SourceMultiplicity = models.MANY
	__LinkShape__00000008_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000008_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000008_.X = 523.000000
	__LinkShape__00000008_.Y = 359.500000
	__LinkShape__00000008_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000008_.StartRatio = 0.756597
	__LinkShape__00000008_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000008_.EndRatio = 0.500000
	__LinkShape__00000008_.CornerOffsetRatio = 1.042328

	__LinkShape__00000009_.Name = `Justifications`
	__LinkShape__00000009_.IdentifierMeta = ref_models.ActorStateTransition{}.Justifications
	__LinkShape__00000009_.FieldTypeIdentifierMeta = ref_models.Parameter{}
	__LinkShape__00000009_.FieldOffsetX = 0.000000
	__LinkShape__00000009_.FieldOffsetY = 0.000000
	__LinkShape__00000009_.TargetMultiplicity = models.MANY
	__LinkShape__00000009_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000009_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000009_.SourceMultiplicity = models.MANY
	__LinkShape__00000009_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000009_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000009_.X = 1225.499969
	__LinkShape__00000009_.Y = 505.999985
	__LinkShape__00000009_.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__00000009_.StartRatio = 0.769097
	__LinkShape__00000009_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000009_.EndRatio = 0.500000
	__LinkShape__00000009_.CornerOffsetRatio = 1.002766

	__LinkShape__00000010_.Name = `ActorStateShapes`
	__LinkShape__00000010_.IdentifierMeta = ref_models.Diagram{}.ActorStateShapes
	__LinkShape__00000010_.FieldTypeIdentifierMeta = ref_models.ActorStateShape{}
	__LinkShape__00000010_.FieldOffsetX = 0.000000
	__LinkShape__00000010_.FieldOffsetY = 0.000000
	__LinkShape__00000010_.TargetMultiplicity = models.MANY
	__LinkShape__00000010_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000010_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000010_.SourceMultiplicity = models.MANY
	__LinkShape__00000010_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000010_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000010_.X = 975.500000
	__LinkShape__00000010_.Y = 388.499985
	__LinkShape__00000010_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000010_.StartRatio = 0.063561
	__LinkShape__00000010_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000010_.EndRatio = 0.500000
	__LinkShape__00000010_.CornerOffsetRatio = 1.380000

	__LinkShape__00000011_.Name = `ParameterShapes`
	__LinkShape__00000011_.IdentifierMeta = ref_models.Diagram{}.ParameterShapes
	__LinkShape__00000011_.FieldTypeIdentifierMeta = ref_models.ParameterShape{}
	__LinkShape__00000011_.FieldOffsetX = 0.000000
	__LinkShape__00000011_.FieldOffsetY = 0.000000
	__LinkShape__00000011_.TargetMultiplicity = models.MANY
	__LinkShape__00000011_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000011_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000011_.SourceMultiplicity = models.MANY
	__LinkShape__00000011_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000011_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000011_.X = 979.000000
	__LinkShape__00000011_.Y = 605.999969
	__LinkShape__00000011_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000011_.StartRatio = 0.696749
	__LinkShape__00000011_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000011_.EndRatio = 0.500000
	__LinkShape__00000011_.CornerOffsetRatio = 1.380000

	__LinkShape__00000012_.Name = `ActorStateTransitionShapes`
	__LinkShape__00000012_.IdentifierMeta = ref_models.Diagram{}.ActorStateTransitionShapes
	__LinkShape__00000012_.FieldTypeIdentifierMeta = ref_models.ActorStateTransitionShape{}
	__LinkShape__00000012_.FieldOffsetX = 0.000000
	__LinkShape__00000012_.FieldOffsetY = 0.000000
	__LinkShape__00000012_.TargetMultiplicity = models.MANY
	__LinkShape__00000012_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000012_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000012_.SourceMultiplicity = models.MANY
	__LinkShape__00000012_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000012_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000012_.X = 976.500000
	__LinkShape__00000012_.Y = 456.999985
	__LinkShape__00000012_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000012_.StartRatio = 0.261524
	__LinkShape__00000012_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000012_.EndRatio = 0.500000
	__LinkShape__00000012_.CornerOffsetRatio = 1.380000

	__LinkShape__00000013_.Name = `ScenarioParameterShapes`
	__LinkShape__00000013_.IdentifierMeta = ref_models.Diagram{}.ScenarioParameterShapes
	__LinkShape__00000013_.FieldTypeIdentifierMeta = ref_models.ParametersAggregateShape{}
	__LinkShape__00000013_.FieldOffsetX = 0.000000
	__LinkShape__00000013_.FieldOffsetY = 0.000000
	__LinkShape__00000013_.TargetMultiplicity = models.MANY
	__LinkShape__00000013_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000013_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000013_.SourceMultiplicity = models.MANY
	__LinkShape__00000013_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000013_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000013_.X = 980.500000
	__LinkShape__00000013_.Y = 687.999969
	__LinkShape__00000013_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000013_.StartRatio = 0.935468
	__LinkShape__00000013_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000013_.EndRatio = 0.500000
	__LinkShape__00000013_.CornerOffsetRatio = 1.380000

	__LinkShape__00000014_.Name = `EvolutionDirectionShapes`
	__LinkShape__00000014_.IdentifierMeta = ref_models.Diagram{}.EvolutionDirectionShapes
	__LinkShape__00000014_.FieldTypeIdentifierMeta = ref_models.EvolutionDirectionShape{}
	__LinkShape__00000014_.FieldOffsetX = 0.000000
	__LinkShape__00000014_.FieldOffsetY = 0.000000
	__LinkShape__00000014_.TargetMultiplicity = models.MANY
	__LinkShape__00000014_.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__00000014_.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__00000014_.SourceMultiplicity = models.MANY
	__LinkShape__00000014_.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__00000014_.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__00000014_.X = 979.000000
	__LinkShape__00000014_.Y = 525.499985
	__LinkShape__00000014_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000014_.StartRatio = 0.481320
	__LinkShape__00000014_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__00000014_.EndRatio = 0.500000
	__LinkShape__00000014_.CornerOffsetRatio = 1.380000

	// insertion point for setup of pointers
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000001_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000002_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000003_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000004_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000006_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000009_)
	__Classdiagram__00000000_.GongStructShapes = append(__Classdiagram__00000000_.GongStructShapes, __GongStructShape__00000017_)
	__Classdiagram__00000001_.GongStructShapes = append(__Classdiagram__00000001_.GongStructShapes, __GongStructShape__00000010_)
	__Classdiagram__00000001_.GongStructShapes = append(__Classdiagram__00000001_.GongStructShapes, __GongStructShape__00000011_)
	__Classdiagram__00000001_.GongStructShapes = append(__Classdiagram__00000001_.GongStructShapes, __GongStructShape__00000012_)
	__Classdiagram__00000001_.GongStructShapes = append(__Classdiagram__00000001_.GongStructShapes, __GongStructShape__00000013_)
	__Classdiagram__00000001_.GongStructShapes = append(__Classdiagram__00000001_.GongStructShapes, __GongStructShape__00000014_)
	__Classdiagram__00000001_.GongStructShapes = append(__Classdiagram__00000001_.GongStructShapes, __GongStructShape__00000015_)
	__DiagramPackage__00000000_.Classdiagrams = append(__DiagramPackage__00000000_.Classdiagrams, __Classdiagram__00000000_)
	__DiagramPackage__00000000_.Classdiagrams = append(__DiagramPackage__00000000_.Classdiagrams, __Classdiagram__00000001_)
	__DiagramPackage__00000000_.SelectedClassdiagram = __Classdiagram__00000001_
	__GongStructShape__00000001_.AttributeShapes = append(__GongStructShape__00000001_.AttributeShapes, __AttributeShape__00000000_)
	__GongStructShape__00000001_.AttributeShapes = append(__GongStructShape__00000001_.AttributeShapes, __AttributeShape__00000001_)
	__GongStructShape__00000001_.AttributeShapes = append(__GongStructShape__00000001_.AttributeShapes, __AttributeShape__00000002_)
	__GongStructShape__00000002_.LinkShapes = append(__GongStructShape__00000002_.LinkShapes, __LinkShape__00000000_)
	__GongStructShape__00000002_.LinkShapes = append(__GongStructShape__00000002_.LinkShapes, __LinkShape__00000001_)
	__GongStructShape__00000002_.LinkShapes = append(__GongStructShape__00000002_.LinkShapes, __LinkShape__00000009_)
	__GongStructShape__00000003_.LinkShapes = append(__GongStructShape__00000003_.LinkShapes, __LinkShape__00000002_)
	__GongStructShape__00000004_.LinkShapes = append(__GongStructShape__00000004_.LinkShapes, __LinkShape__00000003_)
	__GongStructShape__00000004_.LinkShapes = append(__GongStructShape__00000004_.LinkShapes, __LinkShape__00000004_)
	__GongStructShape__00000004_.LinkShapes = append(__GongStructShape__00000004_.LinkShapes, __LinkShape__00000005_)
	__GongStructShape__00000004_.LinkShapes = append(__GongStructShape__00000004_.LinkShapes, __LinkShape__00000006_)
	__GongStructShape__00000004_.LinkShapes = append(__GongStructShape__00000004_.LinkShapes, __LinkShape__00000007_)
	__GongStructShape__00000010_.AttributeShapes = append(__GongStructShape__00000010_.AttributeShapes, __AttributeShape__00000012_)
	__GongStructShape__00000010_.AttributeShapes = append(__GongStructShape__00000010_.AttributeShapes, __AttributeShape__00000011_)
	__GongStructShape__00000010_.AttributeShapes = append(__GongStructShape__00000010_.AttributeShapes, __AttributeShape__00000010_)
	__GongStructShape__00000010_.AttributeShapes = append(__GongStructShape__00000010_.AttributeShapes, __AttributeShape__00000009_)
	__GongStructShape__00000010_.AttributeShapes = append(__GongStructShape__00000010_.AttributeShapes, __AttributeShape__00000008_)
	__GongStructShape__00000010_.AttributeShapes = append(__GongStructShape__00000010_.AttributeShapes, __AttributeShape__00000007_)
	__GongStructShape__00000010_.AttributeShapes = append(__GongStructShape__00000010_.AttributeShapes, __AttributeShape__00000004_)
	__GongStructShape__00000010_.AttributeShapes = append(__GongStructShape__00000010_.AttributeShapes, __AttributeShape__00000005_)
	__GongStructShape__00000010_.AttributeShapes = append(__GongStructShape__00000010_.AttributeShapes, __AttributeShape__00000006_)
	__GongStructShape__00000010_.LinkShapes = append(__GongStructShape__00000010_.LinkShapes, __LinkShape__00000010_)
	__GongStructShape__00000010_.LinkShapes = append(__GongStructShape__00000010_.LinkShapes, __LinkShape__00000011_)
	__GongStructShape__00000010_.LinkShapes = append(__GongStructShape__00000010_.LinkShapes, __LinkShape__00000012_)
	__GongStructShape__00000010_.LinkShapes = append(__GongStructShape__00000010_.LinkShapes, __LinkShape__00000013_)
	__GongStructShape__00000010_.LinkShapes = append(__GongStructShape__00000010_.LinkShapes, __LinkShape__00000014_)
	__GongStructShape__00000017_.LinkShapes = append(__GongStructShape__00000017_.LinkShapes, __LinkShape__00000008_)
}
