package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gong/test/statemachines/go/models"
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

	__AttributeShape__000000_IsDecisionNode := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000001_IsFictif := (&models.AttributeShape{}).Stage(stage)
	__AttributeShape__000002_IsEndState := (&models.AttributeShape{}).Stage(stage)

	__Classdiagram__000000_abstract_syntax := (&models.Classdiagram{}).Stage(stage)
	__Classdiagram__000001_concrete_syntax := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_11_29T13_01_53Z := (&models.DiagramPackage{}).Stage(stage)

	__GongStructShape__000000_Default_Architecture := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000001_Default_MessageType := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000002_Default_Role := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000003_Default_State := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000004_Default_StateMachine := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000005_Default_Transition := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000006_concrete_syntax_Diagram := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000007_concrete_syntax_StateMachine := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000008_concrete_syntax_StateShape := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000009_concrete_syntax_Transition_Shape := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000010_concrete_syntax_Transition := (&models.GongStructShape{}).Stage(stage)
	__GongStructShape__000011_concrete_syntax_State := (&models.GongStructShape{}).Stage(stage)

	__LinkShape__000000_Start := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000001_End := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000002_RolesWithPermissions := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000003_GeneratedMessages := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000004_States := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000005_InitialState := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000006_SubStates := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000007_RolesWithSamePermissions := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000008_StateMachines := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000009_Roles := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000010_Diagrams := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000011_State_Shapes := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000012_Transition_Shapes := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000013_State := (&models.LinkShape{}).Stage(stage)
	__LinkShape__000014_Transition := (&models.LinkShape{}).Stage(stage)

	// Setup of values

	__AttributeShape__000000_IsDecisionNode.Name = `IsDecisionNode`
	__AttributeShape__000000_IsDecisionNode.IdentifierMeta = ref_models.State{}.IsDecisionNode
	__AttributeShape__000000_IsDecisionNode.FieldTypeAsString = ``
	__AttributeShape__000000_IsDecisionNode.Structname = `State`
	__AttributeShape__000000_IsDecisionNode.Fieldtypename = `bool`

	__AttributeShape__000001_IsFictif.Name = `IsFictif`
	__AttributeShape__000001_IsFictif.IdentifierMeta = ref_models.State{}.IsFictif
	__AttributeShape__000001_IsFictif.FieldTypeAsString = ``
	__AttributeShape__000001_IsFictif.Structname = `State`
	__AttributeShape__000001_IsFictif.Fieldtypename = `bool`

	__AttributeShape__000002_IsEndState.Name = `IsEndState`
	__AttributeShape__000002_IsEndState.IdentifierMeta = ref_models.State{}.IsEndState
	__AttributeShape__000002_IsEndState.FieldTypeAsString = ``
	__AttributeShape__000002_IsEndState.Structname = `State`
	__AttributeShape__000002_IsEndState.Fieldtypename = `bool`

	__Classdiagram__000000_abstract_syntax.Name = `abstract syntax`
	__Classdiagram__000000_abstract_syntax.Description = ``
	__Classdiagram__000000_abstract_syntax.IsIncludedInStaticWebSite = false
	__Classdiagram__000000_abstract_syntax.ShowNbInstances = true
	__Classdiagram__000000_abstract_syntax.ShowMultiplicity = true
	__Classdiagram__000000_abstract_syntax.ShowLinkNames = true
	__Classdiagram__000000_abstract_syntax.IsInRenameMode = false
	__Classdiagram__000000_abstract_syntax.IsExpanded = false
	__Classdiagram__000000_abstract_syntax.NodeGongStructsIsExpanded = false
	__Classdiagram__000000_abstract_syntax.NodeGongStructNodeExpansion = `[true,false,false,true,true,false,true,true,true,false,false,false]`
	__Classdiagram__000000_abstract_syntax.NodeGongEnumsIsExpanded = false
	__Classdiagram__000000_abstract_syntax.NodeGongEnumNodeExpansion = ``
	__Classdiagram__000000_abstract_syntax.NodeGongNotesIsExpanded = false
	__Classdiagram__000000_abstract_syntax.NodeGongNoteNodeExpansion = ``

	__Classdiagram__000001_concrete_syntax.Name = `concrete syntax`
	__Classdiagram__000001_concrete_syntax.Description = ``
	__Classdiagram__000001_concrete_syntax.IsIncludedInStaticWebSite = false
	__Classdiagram__000001_concrete_syntax.ShowNbInstances = true
	__Classdiagram__000001_concrete_syntax.ShowMultiplicity = true
	__Classdiagram__000001_concrete_syntax.ShowLinkNames = true
	__Classdiagram__000001_concrete_syntax.IsInRenameMode = false
	__Classdiagram__000001_concrete_syntax.IsExpanded = true
	__Classdiagram__000001_concrete_syntax.NodeGongStructsIsExpanded = true
	__Classdiagram__000001_concrete_syntax.NodeGongStructNodeExpansion = `[false,true,false,false,false,false,false,false,true,false,false,false]`
	__Classdiagram__000001_concrete_syntax.NodeGongEnumsIsExpanded = false
	__Classdiagram__000001_concrete_syntax.NodeGongEnumNodeExpansion = ``
	__Classdiagram__000001_concrete_syntax.NodeGongNotesIsExpanded = false
	__Classdiagram__000001_concrete_syntax.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__000000_Diagram_Package_created_the_2025_11_29T13_01_53Z.Name = `Diagram Package created the 2025-11-29T13:01:53Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_11_29T13_01_53Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_11_29T13_01_53Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_11_29T13_01_53Z.AbsolutePathToDiagramPackage = ``

	__GongStructShape__000000_Default_Architecture.Name = `Default-Architecture`
	__GongStructShape__000000_Default_Architecture.X = 91.000000
	__GongStructShape__000000_Default_Architecture.Y = 45.000000
	__GongStructShape__000000_Default_Architecture.IdentifierMeta = ref_models.Architecture{}
	__GongStructShape__000000_Default_Architecture.Width = 1476.000000
	__GongStructShape__000000_Default_Architecture.Height = 70.000000
	__GongStructShape__000000_Default_Architecture.IsSelected = false

	__GongStructShape__000001_Default_MessageType.Name = `Default-MessageType`
	__GongStructShape__000001_Default_MessageType.X = 1269.000000
	__GongStructShape__000001_Default_MessageType.Y = 600.000015
	__GongStructShape__000001_Default_MessageType.IdentifierMeta = ref_models.MessageType{}
	__GongStructShape__000001_Default_MessageType.Width = 240.000000
	__GongStructShape__000001_Default_MessageType.Height = 63.000000
	__GongStructShape__000001_Default_MessageType.IsSelected = false

	__GongStructShape__000002_Default_Role.Name = `Default-Role`
	__GongStructShape__000002_Default_Role.X = 1262.000000
	__GongStructShape__000002_Default_Role.Y = 247.999985
	__GongStructShape__000002_Default_Role.IdentifierMeta = ref_models.Role{}
	__GongStructShape__000002_Default_Role.Width = 240.000000
	__GongStructShape__000002_Default_Role.Height = 178.000000
	__GongStructShape__000002_Default_Role.IsSelected = false

	__GongStructShape__000003_Default_State.Name = `Default-State`
	__GongStructShape__000003_Default_State.X = 379.000000
	__GongStructShape__000003_Default_State.Y = 375.999985
	__GongStructShape__000003_Default_State.IdentifierMeta = ref_models.State{}
	__GongStructShape__000003_Default_State.Width = 240.000000
	__GongStructShape__000003_Default_State.Height = 123.000000
	__GongStructShape__000003_Default_State.IsSelected = false

	__GongStructShape__000004_Default_StateMachine.Name = `Default-StateMachine`
	__GongStructShape__000004_Default_StateMachine.X = 85.000000
	__GongStructShape__000004_Default_StateMachine.Y = 199.000000
	__GongStructShape__000004_Default_StateMachine.IdentifierMeta = ref_models.StateMachine{}
	__GongStructShape__000004_Default_StateMachine.Width = 240.000000
	__GongStructShape__000004_Default_StateMachine.Height = 141.999985
	__GongStructShape__000004_Default_StateMachine.IsSelected = false

	__GongStructShape__000005_Default_Transition.Name = `Default-Transition`
	__GongStructShape__000005_Default_Transition.X = 810.000000
	__GongStructShape__000005_Default_Transition.Y = 250.999985
	__GongStructShape__000005_Default_Transition.IdentifierMeta = ref_models.Transition{}
	__GongStructShape__000005_Default_Transition.Width = 240.000000
	__GongStructShape__000005_Default_Transition.Height = 315.000031
	__GongStructShape__000005_Default_Transition.IsSelected = false

	__GongStructShape__000006_concrete_syntax_Diagram.Name = `concrete syntax-Diagram`
	__GongStructShape__000006_concrete_syntax_Diagram.X = 98.000000
	__GongStructShape__000006_concrete_syntax_Diagram.Y = 274.999985
	__GongStructShape__000006_concrete_syntax_Diagram.IdentifierMeta = ref_models.Diagram{}
	__GongStructShape__000006_concrete_syntax_Diagram.Width = 240.000000
	__GongStructShape__000006_concrete_syntax_Diagram.Height = 63.000000
	__GongStructShape__000006_concrete_syntax_Diagram.IsSelected = false

	__GongStructShape__000007_concrete_syntax_StateMachine.Name = `concrete syntax-StateMachine`
	__GongStructShape__000007_concrete_syntax_StateMachine.X = 93.000000
	__GongStructShape__000007_concrete_syntax_StateMachine.Y = 106.000000
	__GongStructShape__000007_concrete_syntax_StateMachine.IdentifierMeta = ref_models.StateMachine{}
	__GongStructShape__000007_concrete_syntax_StateMachine.Width = 240.000000
	__GongStructShape__000007_concrete_syntax_StateMachine.Height = 63.000000
	__GongStructShape__000007_concrete_syntax_StateMachine.IsSelected = false

	__GongStructShape__000008_concrete_syntax_StateShape.Name = `concrete syntax-StateShape`
	__GongStructShape__000008_concrete_syntax_StateShape.X = 656.000000
	__GongStructShape__000008_concrete_syntax_StateShape.Y = 110.000000
	__GongStructShape__000008_concrete_syntax_StateShape.IdentifierMeta = ref_models.StateShape{}
	__GongStructShape__000008_concrete_syntax_StateShape.Width = 240.000000
	__GongStructShape__000008_concrete_syntax_StateShape.Height = 63.000000
	__GongStructShape__000008_concrete_syntax_StateShape.IsSelected = false

	__GongStructShape__000009_concrete_syntax_Transition_Shape.Name = `concrete syntax-Transition_Shape`
	__GongStructShape__000009_concrete_syntax_Transition_Shape.X = 672.000000
	__GongStructShape__000009_concrete_syntax_Transition_Shape.Y = 274.999985
	__GongStructShape__000009_concrete_syntax_Transition_Shape.IdentifierMeta = ref_models.Transition_Shape{}
	__GongStructShape__000009_concrete_syntax_Transition_Shape.Width = 240.000000
	__GongStructShape__000009_concrete_syntax_Transition_Shape.Height = 63.000000
	__GongStructShape__000009_concrete_syntax_Transition_Shape.IsSelected = false

	__GongStructShape__000010_concrete_syntax_Transition.Name = `concrete syntax-Transition`
	__GongStructShape__000010_concrete_syntax_Transition.X = 1072.000000
	__GongStructShape__000010_concrete_syntax_Transition.Y = 272.999985
	__GongStructShape__000010_concrete_syntax_Transition.IdentifierMeta = ref_models.Transition{}
	__GongStructShape__000010_concrete_syntax_Transition.Width = 240.000000
	__GongStructShape__000010_concrete_syntax_Transition.Height = 63.000000
	__GongStructShape__000010_concrete_syntax_Transition.IsSelected = false

	__GongStructShape__000011_concrete_syntax_State.Name = `concrete syntax-State`
	__GongStructShape__000011_concrete_syntax_State.X = 1063.000000
	__GongStructShape__000011_concrete_syntax_State.Y = 111.000000
	__GongStructShape__000011_concrete_syntax_State.IdentifierMeta = ref_models.State{}
	__GongStructShape__000011_concrete_syntax_State.Width = 240.000000
	__GongStructShape__000011_concrete_syntax_State.Height = 63.000000
	__GongStructShape__000011_concrete_syntax_State.IsSelected = false

	__LinkShape__000000_Start.Name = `Start`
	__LinkShape__000000_Start.IdentifierMeta = ref_models.Transition{}.Start
	__LinkShape__000000_Start.FieldTypeIdentifierMeta = ref_models.State{}
	__LinkShape__000000_Start.FieldOffsetX = 0.000000
	__LinkShape__000000_Start.FieldOffsetY = 0.000000
	__LinkShape__000000_Start.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000000_Start.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000000_Start.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000000_Start.SourceMultiplicity = models.MANY
	__LinkShape__000000_Start.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000000_Start.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000000_Start.X = 406.000000
	__LinkShape__000000_Start.Y = 76.500000
	__LinkShape__000000_Start.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_Start.StartRatio = 0.479387
	__LinkShape__000000_Start.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000000_Start.EndRatio = 0.284609
	__LinkShape__000000_Start.CornerOffsetRatio = -0.331120

	__LinkShape__000001_End.Name = `End`
	__LinkShape__000001_End.IdentifierMeta = ref_models.Transition{}.End
	__LinkShape__000001_End.FieldTypeIdentifierMeta = ref_models.State{}
	__LinkShape__000001_End.FieldOffsetX = 0.000000
	__LinkShape__000001_End.FieldOffsetY = 0.000000
	__LinkShape__000001_End.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000001_End.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000001_End.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000001_End.SourceMultiplicity = models.MANY
	__LinkShape__000001_End.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000001_End.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000001_End.X = 406.000000
	__LinkShape__000001_End.Y = 76.500000
	__LinkShape__000001_End.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_End.StartRatio = 0.707958
	__LinkShape__000001_End.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000001_End.EndRatio = 0.731764
	__LinkShape__000001_End.CornerOffsetRatio = -0.414453

	__LinkShape__000002_RolesWithPermissions.Name = `RolesWithPermissions`
	__LinkShape__000002_RolesWithPermissions.IdentifierMeta = ref_models.Transition{}.RolesWithPermissions
	__LinkShape__000002_RolesWithPermissions.FieldTypeIdentifierMeta = ref_models.Role{}
	__LinkShape__000002_RolesWithPermissions.FieldOffsetX = 0.000000
	__LinkShape__000002_RolesWithPermissions.FieldOffsetY = 0.000000
	__LinkShape__000002_RolesWithPermissions.TargetMultiplicity = models.MANY
	__LinkShape__000002_RolesWithPermissions.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000002_RolesWithPermissions.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000002_RolesWithPermissions.SourceMultiplicity = models.MANY
	__LinkShape__000002_RolesWithPermissions.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000002_RolesWithPermissions.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000002_RolesWithPermissions.X = 396.000000
	__LinkShape__000002_RolesWithPermissions.Y = 87.500000
	__LinkShape__000002_RolesWithPermissions.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000002_RolesWithPermissions.StartRatio = 0.149228
	__LinkShape__000002_RolesWithPermissions.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000002_RolesWithPermissions.EndRatio = 0.258466
	__LinkShape__000002_RolesWithPermissions.CornerOffsetRatio = 1.380000

	__LinkShape__000003_GeneratedMessages.Name = `GeneratedMessages`
	__LinkShape__000003_GeneratedMessages.IdentifierMeta = ref_models.Transition{}.GeneratedMessages
	__LinkShape__000003_GeneratedMessages.FieldTypeIdentifierMeta = ref_models.MessageType{}
	__LinkShape__000003_GeneratedMessages.FieldOffsetX = 0.000000
	__LinkShape__000003_GeneratedMessages.FieldOffsetY = 0.000000
	__LinkShape__000003_GeneratedMessages.TargetMultiplicity = models.MANY
	__LinkShape__000003_GeneratedMessages.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000003_GeneratedMessages.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000003_GeneratedMessages.SourceMultiplicity = models.MANY
	__LinkShape__000003_GeneratedMessages.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000003_GeneratedMessages.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000003_GeneratedMessages.X = 384.000000
	__LinkShape__000003_GeneratedMessages.Y = 51.000000
	__LinkShape__000003_GeneratedMessages.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__000003_GeneratedMessages.StartRatio = 0.618880
	__LinkShape__000003_GeneratedMessages.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000003_GeneratedMessages.EndRatio = 0.500000
	__LinkShape__000003_GeneratedMessages.CornerOffsetRatio = 1.047729

	__LinkShape__000004_States.Name = `States`
	__LinkShape__000004_States.IdentifierMeta = ref_models.StateMachine{}.States
	__LinkShape__000004_States.FieldTypeIdentifierMeta = ref_models.State{}
	__LinkShape__000004_States.FieldOffsetX = 0.000000
	__LinkShape__000004_States.FieldOffsetY = 0.000000
	__LinkShape__000004_States.TargetMultiplicity = models.MANY
	__LinkShape__000004_States.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000004_States.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000004_States.SourceMultiplicity = models.MANY
	__LinkShape__000004_States.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000004_States.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000004_States.X = 838.500000
	__LinkShape__000004_States.Y = 199.499992
	__LinkShape__000004_States.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000004_States.StartRatio = 0.232443
	__LinkShape__000004_States.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__000004_States.EndRatio = 0.748047
	__LinkShape__000004_States.CornerOffsetRatio = 2.254078

	__LinkShape__000005_InitialState.Name = `InitialState`
	__LinkShape__000005_InitialState.IdentifierMeta = ref_models.StateMachine{}.InitialState
	__LinkShape__000005_InitialState.FieldTypeIdentifierMeta = ref_models.State{}
	__LinkShape__000005_InitialState.FieldOffsetX = 0.000000
	__LinkShape__000005_InitialState.FieldOffsetY = 0.000000
	__LinkShape__000005_InitialState.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000005_InitialState.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000005_InitialState.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000005_InitialState.SourceMultiplicity = models.MANY
	__LinkShape__000005_InitialState.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000005_InitialState.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000005_InitialState.X = 838.500000
	__LinkShape__000005_InitialState.Y = 199.499992
	__LinkShape__000005_InitialState.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000005_InitialState.StartRatio = 0.795824
	__LinkShape__000005_InitialState.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__000005_InitialState.EndRatio = 0.323047
	__LinkShape__000005_InitialState.CornerOffsetRatio = 2.301698

	__LinkShape__000006_SubStates.Name = `SubStates`
	__LinkShape__000006_SubStates.IdentifierMeta = ref_models.State{}.SubStates
	__LinkShape__000006_SubStates.FieldTypeIdentifierMeta = ref_models.State{}
	__LinkShape__000006_SubStates.FieldOffsetX = 0.000000
	__LinkShape__000006_SubStates.FieldOffsetY = 0.000000
	__LinkShape__000006_SubStates.TargetMultiplicity = models.MANY
	__LinkShape__000006_SubStates.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000006_SubStates.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000006_SubStates.SourceMultiplicity = models.MANY
	__LinkShape__000006_SubStates.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000006_SubStates.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000006_SubStates.X = 1244.000000
	__LinkShape__000006_SubStates.Y = 444.499985
	__LinkShape__000006_SubStates.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__000006_SubStates.StartRatio = 0.227214
	__LinkShape__000006_SubStates.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__000006_SubStates.EndRatio = 0.793880
	__LinkShape__000006_SubStates.CornerOffsetRatio = 1.674853

	__LinkShape__000007_RolesWithSamePermissions.Name = `RolesWithSamePermissions`
	__LinkShape__000007_RolesWithSamePermissions.IdentifierMeta = ref_models.Role{}.RolesWithSamePermissions
	__LinkShape__000007_RolesWithSamePermissions.FieldTypeIdentifierMeta = ref_models.Role{}
	__LinkShape__000007_RolesWithSamePermissions.FieldOffsetX = 0.000000
	__LinkShape__000007_RolesWithSamePermissions.FieldOffsetY = 0.000000
	__LinkShape__000007_RolesWithSamePermissions.TargetMultiplicity = models.MANY
	__LinkShape__000007_RolesWithSamePermissions.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000007_RolesWithSamePermissions.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000007_RolesWithSamePermissions.SourceMultiplicity = models.MANY
	__LinkShape__000007_RolesWithSamePermissions.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000007_RolesWithSamePermissions.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000007_RolesWithSamePermissions.X = 443.000000
	__LinkShape__000007_RolesWithSamePermissions.Y = 404.499985
	__LinkShape__000007_RolesWithSamePermissions.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__000007_RolesWithSamePermissions.StartRatio = 0.760547
	__LinkShape__000007_RolesWithSamePermissions.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__000007_RolesWithSamePermissions.EndRatio = 0.468880
	__LinkShape__000007_RolesWithSamePermissions.CornerOffsetRatio = 1.427005

	__LinkShape__000008_StateMachines.Name = `StateMachines`
	__LinkShape__000008_StateMachines.IdentifierMeta = ref_models.Architecture{}.StateMachines
	__LinkShape__000008_StateMachines.FieldTypeIdentifierMeta = ref_models.StateMachine{}
	__LinkShape__000008_StateMachines.FieldOffsetX = 0.000000
	__LinkShape__000008_StateMachines.FieldOffsetY = 0.000000
	__LinkShape__000008_StateMachines.TargetMultiplicity = models.MANY
	__LinkShape__000008_StateMachines.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000008_StateMachines.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000008_StateMachines.SourceMultiplicity = models.MANY
	__LinkShape__000008_StateMachines.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000008_StateMachines.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000008_StateMachines.X = 829.500000
	__LinkShape__000008_StateMachines.Y = 197.499992
	__LinkShape__000008_StateMachines.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__000008_StateMachines.StartRatio = 0.146024
	__LinkShape__000008_StateMachines.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__000008_StateMachines.EndRatio = 0.985547
	__LinkShape__000008_StateMachines.CornerOffsetRatio = 1.442956

	__LinkShape__000009_Roles.Name = `Roles`
	__LinkShape__000009_Roles.IdentifierMeta = ref_models.Architecture{}.Roles
	__LinkShape__000009_Roles.FieldTypeIdentifierMeta = ref_models.Role{}
	__LinkShape__000009_Roles.FieldOffsetX = 0.000000
	__LinkShape__000009_Roles.FieldOffsetY = 0.000000
	__LinkShape__000009_Roles.TargetMultiplicity = models.MANY
	__LinkShape__000009_Roles.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000009_Roles.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000009_Roles.SourceMultiplicity = models.MANY
	__LinkShape__000009_Roles.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000009_Roles.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000009_Roles.X = 471.500000
	__LinkShape__000009_Roles.Y = 254.499992
	__LinkShape__000009_Roles.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__000009_Roles.StartRatio = 0.927189
	__LinkShape__000009_Roles.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__000009_Roles.EndRatio = 0.843880
	__LinkShape__000009_Roles.CornerOffsetRatio = 1.700099

	__LinkShape__000010_Diagrams.Name = `Diagrams`
	__LinkShape__000010_Diagrams.IdentifierMeta = ref_models.StateMachine{}.Diagrams
	__LinkShape__000010_Diagrams.FieldTypeIdentifierMeta = ref_models.Diagram{}
	__LinkShape__000010_Diagrams.FieldOffsetX = 0.000000
	__LinkShape__000010_Diagrams.FieldOffsetY = 0.000000
	__LinkShape__000010_Diagrams.TargetMultiplicity = models.MANY
	__LinkShape__000010_Diagrams.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000010_Diagrams.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000010_Diagrams.SourceMultiplicity = models.MANY
	__LinkShape__000010_Diagrams.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000010_Diagrams.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000010_Diagrams.X = 443.000000
	__LinkShape__000010_Diagrams.Y = 65.500000
	__LinkShape__000010_Diagrams.StartOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__000010_Diagrams.StartRatio = 0.437630
	__LinkShape__000010_Diagrams.EndOrientation = models.ORIENTATION_VERTICAL
	__LinkShape__000010_Diagrams.EndRatio = 0.387630
	__LinkShape__000010_Diagrams.CornerOffsetRatio = 1.571539

	__LinkShape__000011_State_Shapes.Name = `State_Shapes`
	__LinkShape__000011_State_Shapes.IdentifierMeta = ref_models.Diagram{}.State_Shapes
	__LinkShape__000011_State_Shapes.FieldTypeIdentifierMeta = ref_models.StateShape{}
	__LinkShape__000011_State_Shapes.FieldOffsetX = 0.000000
	__LinkShape__000011_State_Shapes.FieldOffsetY = 0.000000
	__LinkShape__000011_State_Shapes.TargetMultiplicity = models.MANY
	__LinkShape__000011_State_Shapes.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000011_State_Shapes.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000011_State_Shapes.SourceMultiplicity = models.MANY
	__LinkShape__000011_State_Shapes.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000011_State_Shapes.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000011_State_Shapes.X = 655.000000
	__LinkShape__000011_State_Shapes.Y = 223.499992
	__LinkShape__000011_State_Shapes.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000011_State_Shapes.StartRatio = 0.500000
	__LinkShape__000011_State_Shapes.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000011_State_Shapes.EndRatio = 0.500000
	__LinkShape__000011_State_Shapes.CornerOffsetRatio = 1.380000

	__LinkShape__000012_Transition_Shapes.Name = `Transition_Shapes`
	__LinkShape__000012_Transition_Shapes.IdentifierMeta = ref_models.Diagram{}.Transition_Shapes
	__LinkShape__000012_Transition_Shapes.FieldTypeIdentifierMeta = ref_models.Transition_Shape{}
	__LinkShape__000012_Transition_Shapes.FieldOffsetX = 0.000000
	__LinkShape__000012_Transition_Shapes.FieldOffsetY = 0.000000
	__LinkShape__000012_Transition_Shapes.TargetMultiplicity = models.MANY
	__LinkShape__000012_Transition_Shapes.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000012_Transition_Shapes.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000012_Transition_Shapes.SourceMultiplicity = models.MANY
	__LinkShape__000012_Transition_Shapes.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000012_Transition_Shapes.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000012_Transition_Shapes.X = 658.000000
	__LinkShape__000012_Transition_Shapes.Y = 303.999985
	__LinkShape__000012_Transition_Shapes.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000012_Transition_Shapes.StartRatio = 0.500000
	__LinkShape__000012_Transition_Shapes.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000012_Transition_Shapes.EndRatio = 0.500000
	__LinkShape__000012_Transition_Shapes.CornerOffsetRatio = 1.380000

	__LinkShape__000013_State.Name = `State`
	__LinkShape__000013_State.IdentifierMeta = ref_models.StateShape{}.State
	__LinkShape__000013_State.FieldTypeIdentifierMeta = ref_models.State{}
	__LinkShape__000013_State.FieldOffsetX = 0.000000
	__LinkShape__000013_State.FieldOffsetY = 0.000000
	__LinkShape__000013_State.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000013_State.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000013_State.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000013_State.SourceMultiplicity = models.MANY
	__LinkShape__000013_State.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000013_State.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000013_State.X = 1224.000000
	__LinkShape__000013_State.Y = 221.999992
	__LinkShape__000013_State.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000013_State.StartRatio = 0.500000
	__LinkShape__000013_State.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000013_State.EndRatio = 0.500000
	__LinkShape__000013_State.CornerOffsetRatio = 1.380000

	__LinkShape__000014_Transition.Name = `Transition`
	__LinkShape__000014_Transition.IdentifierMeta = ref_models.Transition_Shape{}.Transition
	__LinkShape__000014_Transition.FieldTypeIdentifierMeta = ref_models.Transition{}
	__LinkShape__000014_Transition.FieldOffsetX = 0.000000
	__LinkShape__000014_Transition.FieldOffsetY = 0.000000
	__LinkShape__000014_Transition.TargetMultiplicity = models.ZERO_ONE
	__LinkShape__000014_Transition.TargetMultiplicityOffsetX = 0.000000
	__LinkShape__000014_Transition.TargetMultiplicityOffsetY = 0.000000
	__LinkShape__000014_Transition.SourceMultiplicity = models.MANY
	__LinkShape__000014_Transition.SourceMultiplicityOffsetX = 0.000000
	__LinkShape__000014_Transition.SourceMultiplicityOffsetY = 0.000000
	__LinkShape__000014_Transition.X = 1228.500000
	__LinkShape__000014_Transition.Y = 223.999992
	__LinkShape__000014_Transition.StartOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000014_Transition.StartRatio = 0.500000
	__LinkShape__000014_Transition.EndOrientation = models.ORIENTATION_HORIZONTAL
	__LinkShape__000014_Transition.EndRatio = 0.500000
	__LinkShape__000014_Transition.CornerOffsetRatio = 1.380000

	// Setup of pointers
	// setup of AttributeShape instances pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_abstract_syntax.GongStructShapes = append(__Classdiagram__000000_abstract_syntax.GongStructShapes, __GongStructShape__000000_Default_Architecture)
	__Classdiagram__000000_abstract_syntax.GongStructShapes = append(__Classdiagram__000000_abstract_syntax.GongStructShapes, __GongStructShape__000001_Default_MessageType)
	__Classdiagram__000000_abstract_syntax.GongStructShapes = append(__Classdiagram__000000_abstract_syntax.GongStructShapes, __GongStructShape__000002_Default_Role)
	__Classdiagram__000000_abstract_syntax.GongStructShapes = append(__Classdiagram__000000_abstract_syntax.GongStructShapes, __GongStructShape__000003_Default_State)
	__Classdiagram__000000_abstract_syntax.GongStructShapes = append(__Classdiagram__000000_abstract_syntax.GongStructShapes, __GongStructShape__000004_Default_StateMachine)
	__Classdiagram__000000_abstract_syntax.GongStructShapes = append(__Classdiagram__000000_abstract_syntax.GongStructShapes, __GongStructShape__000005_Default_Transition)
	__Classdiagram__000001_concrete_syntax.GongStructShapes = append(__Classdiagram__000001_concrete_syntax.GongStructShapes, __GongStructShape__000006_concrete_syntax_Diagram)
	__Classdiagram__000001_concrete_syntax.GongStructShapes = append(__Classdiagram__000001_concrete_syntax.GongStructShapes, __GongStructShape__000007_concrete_syntax_StateMachine)
	__Classdiagram__000001_concrete_syntax.GongStructShapes = append(__Classdiagram__000001_concrete_syntax.GongStructShapes, __GongStructShape__000008_concrete_syntax_StateShape)
	__Classdiagram__000001_concrete_syntax.GongStructShapes = append(__Classdiagram__000001_concrete_syntax.GongStructShapes, __GongStructShape__000009_concrete_syntax_Transition_Shape)
	__Classdiagram__000001_concrete_syntax.GongStructShapes = append(__Classdiagram__000001_concrete_syntax.GongStructShapes, __GongStructShape__000010_concrete_syntax_Transition)
	__Classdiagram__000001_concrete_syntax.GongStructShapes = append(__Classdiagram__000001_concrete_syntax.GongStructShapes, __GongStructShape__000011_concrete_syntax_State)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_11_29T13_01_53Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_11_29T13_01_53Z.Classdiagrams, __Classdiagram__000000_abstract_syntax)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_11_29T13_01_53Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_11_29T13_01_53Z.Classdiagrams, __Classdiagram__000001_concrete_syntax)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_11_29T13_01_53Z.SelectedClassdiagram = __Classdiagram__000001_concrete_syntax
	// setup of GongStructShape instances pointers
	__GongStructShape__000000_Default_Architecture.LinkShapes = append(__GongStructShape__000000_Default_Architecture.LinkShapes, __LinkShape__000008_StateMachines)
	__GongStructShape__000000_Default_Architecture.LinkShapes = append(__GongStructShape__000000_Default_Architecture.LinkShapes, __LinkShape__000009_Roles)
	__GongStructShape__000002_Default_Role.LinkShapes = append(__GongStructShape__000002_Default_Role.LinkShapes, __LinkShape__000007_RolesWithSamePermissions)
	__GongStructShape__000003_Default_State.AttributeShapes = append(__GongStructShape__000003_Default_State.AttributeShapes, __AttributeShape__000000_IsDecisionNode)
	__GongStructShape__000003_Default_State.AttributeShapes = append(__GongStructShape__000003_Default_State.AttributeShapes, __AttributeShape__000001_IsFictif)
	__GongStructShape__000003_Default_State.AttributeShapes = append(__GongStructShape__000003_Default_State.AttributeShapes, __AttributeShape__000002_IsEndState)
	__GongStructShape__000003_Default_State.LinkShapes = append(__GongStructShape__000003_Default_State.LinkShapes, __LinkShape__000006_SubStates)
	__GongStructShape__000004_Default_StateMachine.LinkShapes = append(__GongStructShape__000004_Default_StateMachine.LinkShapes, __LinkShape__000004_States)
	__GongStructShape__000004_Default_StateMachine.LinkShapes = append(__GongStructShape__000004_Default_StateMachine.LinkShapes, __LinkShape__000005_InitialState)
	__GongStructShape__000005_Default_Transition.LinkShapes = append(__GongStructShape__000005_Default_Transition.LinkShapes, __LinkShape__000000_Start)
	__GongStructShape__000005_Default_Transition.LinkShapes = append(__GongStructShape__000005_Default_Transition.LinkShapes, __LinkShape__000001_End)
	__GongStructShape__000005_Default_Transition.LinkShapes = append(__GongStructShape__000005_Default_Transition.LinkShapes, __LinkShape__000002_RolesWithPermissions)
	__GongStructShape__000005_Default_Transition.LinkShapes = append(__GongStructShape__000005_Default_Transition.LinkShapes, __LinkShape__000003_GeneratedMessages)
	__GongStructShape__000006_concrete_syntax_Diagram.LinkShapes = append(__GongStructShape__000006_concrete_syntax_Diagram.LinkShapes, __LinkShape__000011_State_Shapes)
	__GongStructShape__000006_concrete_syntax_Diagram.LinkShapes = append(__GongStructShape__000006_concrete_syntax_Diagram.LinkShapes, __LinkShape__000012_Transition_Shapes)
	__GongStructShape__000007_concrete_syntax_StateMachine.LinkShapes = append(__GongStructShape__000007_concrete_syntax_StateMachine.LinkShapes, __LinkShape__000010_Diagrams)
	__GongStructShape__000008_concrete_syntax_StateShape.LinkShapes = append(__GongStructShape__000008_concrete_syntax_StateShape.LinkShapes, __LinkShape__000013_State)
	__GongStructShape__000009_concrete_syntax_Transition_Shape.LinkShapes = append(__GongStructShape__000009_concrete_syntax_Transition_Shape.LinkShapes, __LinkShape__000014_Transition)
	// setup of LinkShape instances pointers
}

