package main

import (
	"time"

	"github.com/fullstack-lang/gong/test/statemachines/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// Declaration of instances to stage

	__Action__000000_Lock := (&models.Action{}).Stage(stage)
	__Action__000001_Unlock := (&models.Action{}).Stage(stage)
	__Action__000002_RLock := (&models.Action{}).Stage(stage)
	__Action__000003_Runlock := (&models.Action{}).Stage(stage)

	__Architecture__000000_Gong_UX_loop_Architecture := (&models.Architecture{}).Stage(stage)

	__Diagram__000000_UX_Loop_Diagram := (&models.Diagram{}).Stage(stage)

	__Guard__000000_yes := (&models.Guard{}).Stage(stage)
	__Guard__000001_no := (&models.Guard{}).Stage(stage)
	__Guard__000002_a_suppress_action := (&models.Guard{}).Stage(stage)
	__Guard__000003_not_a_suppress_action := (&models.Guard{}).Stage(stage)

	__State__000000_Probe_Form := (&models.State{}).Stage(stage)
	__State__000001_Probe_Tree := (&models.State{}).Stage(stage)
	__State__000002_Probe_Table := (&models.State{}).Stage(stage)
	__State__000003_Form_Update_Stage_Of_Interest := (&models.State{}).Stage(stage)
	__State__000004_Enforce_Model_Semantic := (&models.State{}).Stage(stage)
	__State__000005_Update_Probe_Form := (&models.State{}).Stage(stage)
	__State__000006_Update_Probe_Table := (&models.State{}).Stage(stage)
	__State__000007_Update_Probe_Tree := (&models.State{}).Stage(stage)
	__State__000008_Load_Stage := (&models.State{}).Stage(stage)
	__State__000009_Persist_Stage := (&models.State{}).Stage(stage)
	__State__000010_Stage_Modified_ := (&models.State{}).Stage(stage)
	__State__000011_Update_UX_ := (&models.State{}).Stage(stage)
	__State__000012_Clean_Stage := (&models.State{}).Stage(stage)

	__StateMachine__000000_UX_Loop := (&models.StateMachine{}).Stage(stage)

	__StateShape__000000_Probe_Form := (&models.StateShape{}).Stage(stage)
	__StateShape__000001_Probe_Tree := (&models.StateShape{}).Stage(stage)
	__StateShape__000002_Probe_Table := (&models.StateShape{}).Stage(stage)
	__StateShape__000003_Form_Update_Stage_Of_Interest := (&models.StateShape{}).Stage(stage)
	__StateShape__000004_Enforce_Model_Semantic := (&models.StateShape{}).Stage(stage)
	__StateShape__000005_Update_Probe_Form := (&models.StateShape{}).Stage(stage)
	__StateShape__000006_Update_Probe_Table := (&models.StateShape{}).Stage(stage)
	__StateShape__000007_Update_Probe_Tree := (&models.StateShape{}).Stage(stage)
	__StateShape__000008_Load_Stage := (&models.StateShape{}).Stage(stage)
	__StateShape__000009_Persist_Stage := (&models.StateShape{}).Stage(stage)
	__StateShape__000010_Stage_Modified_ := (&models.StateShape{}).Stage(stage)
	__StateShape__000011_Update_UX_ := (&models.StateShape{}).Stage(stage)
	__StateShape__000012_Clean_Stage := (&models.StateShape{}).Stage(stage)

	__Transition__000000_Submit := (&models.Transition{}).Stage(stage)
	__Transition__000001_ := (&models.Transition{}).Stage(stage)
	__Transition__000002__Model_Updated := (&models.Transition{}).Stage(stage)
	__Transition__000003_UX_1_updated := (&models.Transition{}).Stage(stage)
	__Transition__000004_UX_2_updated := (&models.Transition{}).Stage(stage)
	__Transition__000005__Stage_loaded := (&models.Transition{}).Stage(stage)
	__Transition__000006__Stage_persisted := (&models.Transition{}).Stage(stage)
	__Transition__000007_Yes := (&models.Transition{}).Stage(stage)
	__Transition__000008_Yes := (&models.Transition{}).Stage(stage)
	__Transition__000009_Form_Update_Stage_Of_Interest_to_Clean_Stage := (&models.Transition{}).Stage(stage)
	__Transition__000010_ := (&models.Transition{}).Stage(stage)

	__Transition_Shape__000000_UX_1_Waiting_for_User_Input_to_UX_1_Update_Stage := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000001_UX_1_Update_Stage_to_Enforce_Model_Semantic := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000002_Enforce_Model_Semantic_to_UX_1_Update_UX := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000003_Enforce_Model_Semantic_to_UX_2_Update_UX := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000004_Enforce_Model_Semantic_to_UX_n_Update_UX := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000005_Load_Stage_to_Enforce_Model_Semantic := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000006_Persist_Stage_to_UX_1_Update_UX := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000007_Stage_Modified_to_Persist_Stage := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000008_Update_UX_to_UX_1_Update_UX := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000009_Form_Update_Stage_Of_Interest_to_Clean_Stage := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000010_Clean_Stage_to_Enforce_Model_Semantic := (&models.Transition_Shape{}).Stage(stage)

	// Setup of values

	__Action__000000_Lock.Name = `Lock`
	__Action__000000_Lock.Criticality = models.CriticalityCritical

	__Action__000001_Unlock.Name = `Unlock`
	__Action__000001_Unlock.Criticality = models.CriticalityCritical

	__Action__000002_RLock.Name = `RLock`
	__Action__000002_RLock.Criticality = models.CriticalityDefault

	__Action__000003_Runlock.Name = `Runlock`
	__Action__000003_Runlock.Criticality = models.CriticalityDefault

	__Architecture__000000_Gong_UX_loop_Architecture.Name = `Gong UX loop Architecture`
	__Architecture__000000_Gong_UX_loop_Architecture.NbPixPerCharacter = 8.000000

	__Diagram__000000_UX_Loop_Diagram.Name = `UX Loop Diagram`
	__Diagram__000000_UX_Loop_Diagram.IsChecked = true
	__Diagram__000000_UX_Loop_Diagram.IsExpanded = true
	__Diagram__000000_UX_Loop_Diagram.IsEditable_ = true
	__Diagram__000000_UX_Loop_Diagram.IsInRenameMode = false

	__Guard__000000_yes.Name = `yes`

	__Guard__000001_no.Name = `no`

	__Guard__000002_a_suppress_action.Name = `a suppress action`

	__Guard__000003_not_a_suppress_action.Name = `not a suppress action`

	__State__000000_Probe_Form.Name = `Probe Form`
	__State__000000_Probe_Form.IsDecisionNode = false
	__State__000000_Probe_Form.IsFictif = false
	__State__000000_Probe_Form.IsEndState = false

	__State__000001_Probe_Tree.Name = `Probe Tree`
	__State__000001_Probe_Tree.IsDecisionNode = false
	__State__000001_Probe_Tree.IsFictif = false
	__State__000001_Probe_Tree.IsEndState = false

	__State__000002_Probe_Table.Name = `Probe Table`
	__State__000002_Probe_Table.IsDecisionNode = false
	__State__000002_Probe_Table.IsFictif = false
	__State__000002_Probe_Table.IsEndState = false

	__State__000003_Form_Update_Stage_Of_Interest.Name = `Form - Update Stage Of Interest`
	__State__000003_Form_Update_Stage_Of_Interest.IsDecisionNode = false
	__State__000003_Form_Update_Stage_Of_Interest.IsFictif = false
	__State__000003_Form_Update_Stage_Of_Interest.IsEndState = false

	__State__000004_Enforce_Model_Semantic.Name = `Enforce Model Semantic`
	__State__000004_Enforce_Model_Semantic.IsDecisionNode = false
	__State__000004_Enforce_Model_Semantic.IsFictif = false
	__State__000004_Enforce_Model_Semantic.IsEndState = false

	__State__000005_Update_Probe_Form.Name = `Update Probe Form`
	__State__000005_Update_Probe_Form.IsDecisionNode = false
	__State__000005_Update_Probe_Form.IsFictif = false
	__State__000005_Update_Probe_Form.IsEndState = false

	__State__000006_Update_Probe_Table.Name = `Update Probe Table`
	__State__000006_Update_Probe_Table.IsDecisionNode = false
	__State__000006_Update_Probe_Table.IsFictif = false
	__State__000006_Update_Probe_Table.IsEndState = false

	__State__000007_Update_Probe_Tree.Name = `Update Probe Tree`
	__State__000007_Update_Probe_Tree.IsDecisionNode = false
	__State__000007_Update_Probe_Tree.IsFictif = false
	__State__000007_Update_Probe_Tree.IsEndState = false

	__State__000008_Load_Stage.Name = `Load Stage`
	__State__000008_Load_Stage.IsDecisionNode = false
	__State__000008_Load_Stage.IsFictif = false
	__State__000008_Load_Stage.IsEndState = false

	__State__000009_Persist_Stage.Name = `Persist Stage`
	__State__000009_Persist_Stage.IsDecisionNode = false
	__State__000009_Persist_Stage.IsFictif = false
	__State__000009_Persist_Stage.IsEndState = false

	__State__000010_Stage_Modified_.Name = `Stage Modified ?`
	__State__000010_Stage_Modified_.IsDecisionNode = true
	__State__000010_Stage_Modified_.IsFictif = true
	__State__000010_Stage_Modified_.IsEndState = false

	__State__000011_Update_UX_.Name = `Update UX ?`
	__State__000011_Update_UX_.IsDecisionNode = true
	__State__000011_Update_UX_.IsFictif = false
	__State__000011_Update_UX_.IsEndState = false

	__State__000012_Clean_Stage.Name = `Clean Stage`
	__State__000012_Clean_Stage.IsDecisionNode = false
	__State__000012_Clean_Stage.IsFictif = false
	__State__000012_Clean_Stage.IsEndState = false

	__StateMachine__000000_UX_Loop.Name = `UX Loop`
	__StateMachine__000000_UX_Loop.IsNodeExpanded = true

	__StateShape__000000_Probe_Form.Name = `Probe Form`
	__StateShape__000000_Probe_Form.IsExpanded = false
	__StateShape__000000_Probe_Form.X = 360.000000
	__StateShape__000000_Probe_Form.Y = 154.000000
	__StateShape__000000_Probe_Form.Width = 221.000000
	__StateShape__000000_Probe_Form.Height = 80.000000

	__StateShape__000001_Probe_Tree.Name = `Probe Tree`
	__StateShape__000001_Probe_Tree.IsExpanded = false
	__StateShape__000001_Probe_Tree.X = 601.000000
	__StateShape__000001_Probe_Tree.Y = 152.000000
	__StateShape__000001_Probe_Tree.Width = 200.000000
	__StateShape__000001_Probe_Tree.Height = 80.000000

	__StateShape__000002_Probe_Table.Name = `Probe Table`
	__StateShape__000002_Probe_Table.IsExpanded = false
	__StateShape__000002_Probe_Table.X = 821.000000
	__StateShape__000002_Probe_Table.Y = 151.000000
	__StateShape__000002_Probe_Table.Width = 200.000000
	__StateShape__000002_Probe_Table.Height = 80.000000

	__StateShape__000003_Form_Update_Stage_Of_Interest.Name = `Form - Update Stage Of Interest`
	__StateShape__000003_Form_Update_Stage_Of_Interest.IsExpanded = false
	__StateShape__000003_Form_Update_Stage_Of_Interest.X = 379.000000
	__StateShape__000003_Form_Update_Stage_Of_Interest.Y = 404.999985
	__StateShape__000003_Form_Update_Stage_Of_Interest.Width = 200.000000
	__StateShape__000003_Form_Update_Stage_Of_Interest.Height = 108.000000

	__StateShape__000004_Enforce_Model_Semantic.Name = `Enforce Model Semantic`
	__StateShape__000004_Enforce_Model_Semantic.IsExpanded = false
	__StateShape__000004_Enforce_Model_Semantic.X = 380.000000
	__StateShape__000004_Enforce_Model_Semantic.Y = 677.000015
	__StateShape__000004_Enforce_Model_Semantic.Width = 646.000000
	__StateShape__000004_Enforce_Model_Semantic.Height = 80.000000

	__StateShape__000005_Update_Probe_Form.Name = `Update Probe Form`
	__StateShape__000005_Update_Probe_Form.IsExpanded = false
	__StateShape__000005_Update_Probe_Form.X = 329.000000
	__StateShape__000005_Update_Probe_Form.Y = 1005.999954
	__StateShape__000005_Update_Probe_Form.Width = 200.000000
	__StateShape__000005_Update_Probe_Form.Height = 80.000000

	__StateShape__000006_Update_Probe_Table.Name = `Update Probe Table`
	__StateShape__000006_Update_Probe_Table.IsExpanded = false
	__StateShape__000006_Update_Probe_Table.X = 552.000000
	__StateShape__000006_Update_Probe_Table.Y = 1003.999954
	__StateShape__000006_Update_Probe_Table.Width = 200.000000
	__StateShape__000006_Update_Probe_Table.Height = 80.000000

	__StateShape__000007_Update_Probe_Tree.Name = `Update Probe Tree`
	__StateShape__000007_Update_Probe_Tree.IsExpanded = false
	__StateShape__000007_Update_Probe_Tree.X = 772.000000
	__StateShape__000007_Update_Probe_Tree.Y = 1003.999954
	__StateShape__000007_Update_Probe_Tree.Width = 200.000000
	__StateShape__000007_Update_Probe_Tree.Height = 80.000000

	__StateShape__000008_Load_Stage.Name = `Load Stage`
	__StateShape__000008_Load_Stage.IsExpanded = false
	__StateShape__000008_Load_Stage.X = 14.000000
	__StateShape__000008_Load_Stage.Y = 549.999984
	__StateShape__000008_Load_Stage.Width = 107.000000
	__StateShape__000008_Load_Stage.Height = 20.000000

	__StateShape__000009_Persist_Stage.Name = `Persist Stage`
	__StateShape__000009_Persist_Stage.IsExpanded = false
	__StateShape__000009_Persist_Stage.X = 52.000000
	__StateShape__000009_Persist_Stage.Y = 888.999954
	__StateShape__000009_Persist_Stage.Width = 200.000000
	__StateShape__000009_Persist_Stage.Height = 80.000000

	__StateShape__000010_Stage_Modified_.Name = `Stage Modified ?`
	__StateShape__000010_Stage_Modified_.IsExpanded = false
	__StateShape__000010_Stage_Modified_.X = 585.000000
	__StateShape__000010_Stage_Modified_.Y = 821.999954
	__StateShape__000010_Stage_Modified_.Width = 153.000000
	__StateShape__000010_Stage_Modified_.Height = 50.000000

	__StateShape__000011_Update_UX_.Name = `Update UX ?`
	__StateShape__000011_Update_UX_.IsExpanded = false
	__StateShape__000011_Update_UX_.X = 105.000000
	__StateShape__000011_Update_UX_.Y = 1156.000015
	__StateShape__000011_Update_UX_.Width = 120.000000
	__StateShape__000011_Update_UX_.Height = 50.000000

	__StateShape__000012_Clean_Stage.Name = `Clean Stage`
	__StateShape__000012_Clean_Stage.IsExpanded = false
	__StateShape__000012_Clean_Stage.X = 786.000000
	__StateShape__000012_Clean_Stage.Y = 411.000000
	__StateShape__000012_Clean_Stage.Width = 200.000000
	__StateShape__000012_Clean_Stage.Height = 80.000000

	__Transition__000000_Submit.Name = `Submit`

	__Transition__000001_.Name = ``

	__Transition__000002__Model_Updated.Name = `/Model Updated`

	__Transition__000003_UX_1_updated.Name = `UX 1 updated`

	__Transition__000004_UX_2_updated.Name = `UX 2 updated`

	__Transition__000005__Stage_loaded.Name = `/Stage loaded`

	__Transition__000006__Stage_persisted.Name = `/Stage persisted`

	__Transition__000007_Yes.Name = `Yes`

	__Transition__000008_Yes.Name = `Yes`

	__Transition__000009_Form_Update_Stage_Of_Interest_to_Clean_Stage.Name = `Form - Update Stage Of Interest to Clean Stage`

	__Transition__000010_.Name = ``

	__Transition_Shape__000000_UX_1_Waiting_for_User_Input_to_UX_1_Update_Stage.Name = `UX 1 - Waiting for User Input to UX 1 - Update Stage`
	__Transition_Shape__000000_UX_1_Waiting_for_User_Input_to_UX_1_Update_Stage.StartRatio = 0.815760
	__Transition_Shape__000000_UX_1_Waiting_for_User_Input_to_UX_1_Update_Stage.EndRatio = 0.866415
	__Transition_Shape__000000_UX_1_Waiting_for_User_Input_to_UX_1_Update_Stage.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000000_UX_1_Waiting_for_User_Input_to_UX_1_Update_Stage.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000000_UX_1_Waiting_for_User_Input_to_UX_1_Update_Stage.CornerOffsetRatio = 2.175087

	__Transition_Shape__000001_UX_1_Update_Stage_to_Enforce_Model_Semantic.Name = `UX 1 - Update Stage to Enforce Model Semantic`
	__Transition_Shape__000001_UX_1_Update_Stage_to_Enforce_Model_Semantic.StartRatio = 0.557656
	__Transition_Shape__000001_UX_1_Update_Stage_to_Enforce_Model_Semantic.EndRatio = 0.166852
	__Transition_Shape__000001_UX_1_Update_Stage_to_Enforce_Model_Semantic.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000001_UX_1_Update_Stage_to_Enforce_Model_Semantic.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000001_UX_1_Update_Stage_to_Enforce_Model_Semantic.CornerOffsetRatio = 1.912587

	__Transition_Shape__000002_Enforce_Model_Semantic_to_UX_1_Update_UX.Name = `Enforce Model Semantic to UX 1 - Update UX`
	__Transition_Shape__000002_Enforce_Model_Semantic_to_UX_1_Update_UX.StartRatio = 0.512053
	__Transition_Shape__000002_Enforce_Model_Semantic_to_UX_1_Update_UX.EndRatio = 0.831915
	__Transition_Shape__000002_Enforce_Model_Semantic_to_UX_1_Update_UX.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000002_Enforce_Model_Semantic_to_UX_1_Update_UX.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000002_Enforce_Model_Semantic_to_UX_1_Update_UX.CornerOffsetRatio = 1.612587

	__Transition_Shape__000003_Enforce_Model_Semantic_to_UX_2_Update_UX.Name = `Enforce Model Semantic to UX 2 - Update UX`
	__Transition_Shape__000003_Enforce_Model_Semantic_to_UX_2_Update_UX.StartRatio = 0.821415
	__Transition_Shape__000003_Enforce_Model_Semantic_to_UX_2_Update_UX.EndRatio = 0.166415
	__Transition_Shape__000003_Enforce_Model_Semantic_to_UX_2_Update_UX.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000003_Enforce_Model_Semantic_to_UX_2_Update_UX.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000003_Enforce_Model_Semantic_to_UX_2_Update_UX.CornerOffsetRatio = 1.625087

	__Transition_Shape__000004_Enforce_Model_Semantic_to_UX_n_Update_UX.Name = `Enforce Model Semantic to UX n - Update UX`
	__Transition_Shape__000004_Enforce_Model_Semantic_to_UX_n_Update_UX.StartRatio = 0.830939
	__Transition_Shape__000004_Enforce_Model_Semantic_to_UX_n_Update_UX.EndRatio = 0.801415
	__Transition_Shape__000004_Enforce_Model_Semantic_to_UX_n_Update_UX.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000004_Enforce_Model_Semantic_to_UX_n_Update_UX.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000004_Enforce_Model_Semantic_to_UX_n_Update_UX.CornerOffsetRatio = 1.662587

	__Transition_Shape__000005_Load_Stage_to_Enforce_Model_Semantic.Name = `Load Stage to Enforce Model Semantic`
	__Transition_Shape__000005_Load_Stage_to_Enforce_Model_Semantic.StartRatio = 0.550349
	__Transition_Shape__000005_Load_Stage_to_Enforce_Model_Semantic.EndRatio = 0.298697
	__Transition_Shape__000005_Load_Stage_to_Enforce_Model_Semantic.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__000005_Load_Stage_to_Enforce_Model_Semantic.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__000005_Load_Stage_to_Enforce_Model_Semantic.CornerOffsetRatio = 1.320401

	__Transition_Shape__000006_Persist_Stage_to_UX_1_Update_UX.Name = `Persist Stage to UX 1 - Update UX`
	__Transition_Shape__000006_Persist_Stage_to_UX_1_Update_UX.StartRatio = 0.741415
	__Transition_Shape__000006_Persist_Stage_to_UX_1_Update_UX.EndRatio = 0.760691
	__Transition_Shape__000006_Persist_Stage_to_UX_1_Update_UX.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000006_Persist_Stage_to_UX_1_Update_UX.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000006_Persist_Stage_to_UX_1_Update_UX.CornerOffsetRatio = 1.737587

	__Transition_Shape__000007_Stage_Modified_to_Persist_Stage.Name = `Stage Modified ? to Persist Stage`
	__Transition_Shape__000007_Stage_Modified_to_Persist_Stage.StartRatio = 0.831915
	__Transition_Shape__000007_Stage_Modified_to_Persist_Stage.EndRatio = 0.350087
	__Transition_Shape__000007_Stage_Modified_to_Persist_Stage.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000007_Stage_Modified_to_Persist_Stage.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__000007_Stage_Modified_to_Persist_Stage.CornerOffsetRatio = 1.100139

	__Transition_Shape__000008_Update_UX_to_UX_1_Update_UX.Name = `Update UX ? to UX 1 - Update UX`
	__Transition_Shape__000008_Update_UX_to_UX_1_Update_UX.StartRatio = 0.475693
	__Transition_Shape__000008_Update_UX_to_UX_1_Update_UX.EndRatio = 0.500000
	__Transition_Shape__000008_Update_UX_to_UX_1_Update_UX.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__000008_Update_UX_to_UX_1_Update_UX.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__000008_Update_UX_to_UX_1_Update_UX.CornerOffsetRatio = 1.560691

	__Transition_Shape__000009_Form_Update_Stage_Of_Interest_to_Clean_Stage.Name = `Form - Update Stage Of Interest to Clean Stage`
	__Transition_Shape__000009_Form_Update_Stage_Of_Interest_to_Clean_Stage.StartRatio = 0.500000
	__Transition_Shape__000009_Form_Update_Stage_Of_Interest_to_Clean_Stage.EndRatio = 0.500000
	__Transition_Shape__000009_Form_Update_Stage_Of_Interest_to_Clean_Stage.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__000009_Form_Update_Stage_Of_Interest_to_Clean_Stage.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__000009_Form_Update_Stage_Of_Interest_to_Clean_Stage.CornerOffsetRatio = 1.200000

	__Transition_Shape__000010_Clean_Stage_to_Enforce_Model_Semantic.Name = `Clean Stage to Enforce Model Semantic`
	__Transition_Shape__000010_Clean_Stage_to_Enforce_Model_Semantic.StartRatio = 0.786415
	__Transition_Shape__000010_Clean_Stage_to_Enforce_Model_Semantic.EndRatio = 0.887435
	__Transition_Shape__000010_Clean_Stage_to_Enforce_Model_Semantic.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000010_Clean_Stage_to_Enforce_Model_Semantic.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000010_Clean_Stage_to_Enforce_Model_Semantic.CornerOffsetRatio = 1.950087

	// Setup of pointers
	// setup of Action instances pointers
	// setup of Architecture instances pointers
	__Architecture__000000_Gong_UX_loop_Architecture.StateMachines = append(__Architecture__000000_Gong_UX_loop_Architecture.StateMachines, __StateMachine__000000_UX_Loop)
	// setup of Diagram instances pointers
	__Diagram__000000_UX_Loop_Diagram.State_Shapes = append(__Diagram__000000_UX_Loop_Diagram.State_Shapes, __StateShape__000011_Update_UX_)
	__Diagram__000000_UX_Loop_Diagram.State_Shapes = append(__Diagram__000000_UX_Loop_Diagram.State_Shapes, __StateShape__000000_Probe_Form)
	__Diagram__000000_UX_Loop_Diagram.State_Shapes = append(__Diagram__000000_UX_Loop_Diagram.State_Shapes, __StateShape__000001_Probe_Tree)
	__Diagram__000000_UX_Loop_Diagram.State_Shapes = append(__Diagram__000000_UX_Loop_Diagram.State_Shapes, __StateShape__000002_Probe_Table)
	__Diagram__000000_UX_Loop_Diagram.State_Shapes = append(__Diagram__000000_UX_Loop_Diagram.State_Shapes, __StateShape__000003_Form_Update_Stage_Of_Interest)
	__Diagram__000000_UX_Loop_Diagram.State_Shapes = append(__Diagram__000000_UX_Loop_Diagram.State_Shapes, __StateShape__000004_Enforce_Model_Semantic)
	__Diagram__000000_UX_Loop_Diagram.State_Shapes = append(__Diagram__000000_UX_Loop_Diagram.State_Shapes, __StateShape__000005_Update_Probe_Form)
	__Diagram__000000_UX_Loop_Diagram.State_Shapes = append(__Diagram__000000_UX_Loop_Diagram.State_Shapes, __StateShape__000006_Update_Probe_Table)
	__Diagram__000000_UX_Loop_Diagram.State_Shapes = append(__Diagram__000000_UX_Loop_Diagram.State_Shapes, __StateShape__000007_Update_Probe_Tree)
	__Diagram__000000_UX_Loop_Diagram.State_Shapes = append(__Diagram__000000_UX_Loop_Diagram.State_Shapes, __StateShape__000008_Load_Stage)
	__Diagram__000000_UX_Loop_Diagram.State_Shapes = append(__Diagram__000000_UX_Loop_Diagram.State_Shapes, __StateShape__000009_Persist_Stage)
	__Diagram__000000_UX_Loop_Diagram.State_Shapes = append(__Diagram__000000_UX_Loop_Diagram.State_Shapes, __StateShape__000010_Stage_Modified_)
	__Diagram__000000_UX_Loop_Diagram.State_Shapes = append(__Diagram__000000_UX_Loop_Diagram.State_Shapes, __StateShape__000012_Clean_Stage)
	__Diagram__000000_UX_Loop_Diagram.Transition_Shapes = append(__Diagram__000000_UX_Loop_Diagram.Transition_Shapes, __Transition_Shape__000000_UX_1_Waiting_for_User_Input_to_UX_1_Update_Stage)
	__Diagram__000000_UX_Loop_Diagram.Transition_Shapes = append(__Diagram__000000_UX_Loop_Diagram.Transition_Shapes, __Transition_Shape__000001_UX_1_Update_Stage_to_Enforce_Model_Semantic)
	__Diagram__000000_UX_Loop_Diagram.Transition_Shapes = append(__Diagram__000000_UX_Loop_Diagram.Transition_Shapes, __Transition_Shape__000002_Enforce_Model_Semantic_to_UX_1_Update_UX)
	__Diagram__000000_UX_Loop_Diagram.Transition_Shapes = append(__Diagram__000000_UX_Loop_Diagram.Transition_Shapes, __Transition_Shape__000003_Enforce_Model_Semantic_to_UX_2_Update_UX)
	__Diagram__000000_UX_Loop_Diagram.Transition_Shapes = append(__Diagram__000000_UX_Loop_Diagram.Transition_Shapes, __Transition_Shape__000004_Enforce_Model_Semantic_to_UX_n_Update_UX)
	__Diagram__000000_UX_Loop_Diagram.Transition_Shapes = append(__Diagram__000000_UX_Loop_Diagram.Transition_Shapes, __Transition_Shape__000005_Load_Stage_to_Enforce_Model_Semantic)
	__Diagram__000000_UX_Loop_Diagram.Transition_Shapes = append(__Diagram__000000_UX_Loop_Diagram.Transition_Shapes, __Transition_Shape__000006_Persist_Stage_to_UX_1_Update_UX)
	__Diagram__000000_UX_Loop_Diagram.Transition_Shapes = append(__Diagram__000000_UX_Loop_Diagram.Transition_Shapes, __Transition_Shape__000007_Stage_Modified_to_Persist_Stage)
	__Diagram__000000_UX_Loop_Diagram.Transition_Shapes = append(__Diagram__000000_UX_Loop_Diagram.Transition_Shapes, __Transition_Shape__000008_Update_UX_to_UX_1_Update_UX)
	__Diagram__000000_UX_Loop_Diagram.Transition_Shapes = append(__Diagram__000000_UX_Loop_Diagram.Transition_Shapes, __Transition_Shape__000009_Form_Update_Stage_Of_Interest_to_Clean_Stage)
	__Diagram__000000_UX_Loop_Diagram.Transition_Shapes = append(__Diagram__000000_UX_Loop_Diagram.Transition_Shapes, __Transition_Shape__000010_Clean_Stage_to_Enforce_Model_Semantic)
	// setup of Guard instances pointers
	// setup of State instances pointers
	__State__000000_Probe_Form.Diagrams = append(__State__000000_Probe_Form.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000001_Probe_Tree.Diagrams = append(__State__000001_Probe_Tree.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000002_Probe_Table.Diagrams = append(__State__000002_Probe_Table.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000003_Form_Update_Stage_Of_Interest.Diagrams = append(__State__000003_Form_Update_Stage_Of_Interest.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000003_Form_Update_Stage_Of_Interest.Entry = __Action__000000_Lock
	__State__000003_Form_Update_Stage_Of_Interest.Exit = __Action__000001_Unlock
	__State__000004_Enforce_Model_Semantic.Diagrams = append(__State__000004_Enforce_Model_Semantic.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000004_Enforce_Model_Semantic.Entry = __Action__000000_Lock
	__State__000004_Enforce_Model_Semantic.Exit = __Action__000001_Unlock
	__State__000005_Update_Probe_Form.Diagrams = append(__State__000005_Update_Probe_Form.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000006_Update_Probe_Table.Diagrams = append(__State__000006_Update_Probe_Table.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000007_Update_Probe_Tree.Diagrams = append(__State__000007_Update_Probe_Tree.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000008_Load_Stage.Diagrams = append(__State__000008_Load_Stage.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000009_Persist_Stage.Diagrams = append(__State__000009_Persist_Stage.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000010_Stage_Modified_.Diagrams = append(__State__000010_Stage_Modified_.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000011_Update_UX_.Diagrams = append(__State__000011_Update_UX_.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000012_Clean_Stage.Diagrams = append(__State__000012_Clean_Stage.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	// setup of StateMachine instances pointers
	__StateMachine__000000_UX_Loop.States = append(__StateMachine__000000_UX_Loop.States, __State__000000_Probe_Form)
	__StateMachine__000000_UX_Loop.States = append(__StateMachine__000000_UX_Loop.States, __State__000001_Probe_Tree)
	__StateMachine__000000_UX_Loop.States = append(__StateMachine__000000_UX_Loop.States, __State__000002_Probe_Table)
	__StateMachine__000000_UX_Loop.States = append(__StateMachine__000000_UX_Loop.States, __State__000003_Form_Update_Stage_Of_Interest)
	__StateMachine__000000_UX_Loop.States = append(__StateMachine__000000_UX_Loop.States, __State__000004_Enforce_Model_Semantic)
	__StateMachine__000000_UX_Loop.States = append(__StateMachine__000000_UX_Loop.States, __State__000005_Update_Probe_Form)
	__StateMachine__000000_UX_Loop.States = append(__StateMachine__000000_UX_Loop.States, __State__000006_Update_Probe_Table)
	__StateMachine__000000_UX_Loop.States = append(__StateMachine__000000_UX_Loop.States, __State__000007_Update_Probe_Tree)
	__StateMachine__000000_UX_Loop.States = append(__StateMachine__000000_UX_Loop.States, __State__000008_Load_Stage)
	__StateMachine__000000_UX_Loop.States = append(__StateMachine__000000_UX_Loop.States, __State__000009_Persist_Stage)
	__StateMachine__000000_UX_Loop.States = append(__StateMachine__000000_UX_Loop.States, __State__000010_Stage_Modified_)
	__StateMachine__000000_UX_Loop.States = append(__StateMachine__000000_UX_Loop.States, __State__000011_Update_UX_)
	__StateMachine__000000_UX_Loop.States = append(__StateMachine__000000_UX_Loop.States, __State__000012_Clean_Stage)
	__StateMachine__000000_UX_Loop.Diagrams = append(__StateMachine__000000_UX_Loop.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__StateMachine__000000_UX_Loop.InitialState = __State__000008_Load_Stage
	// setup of StateShape instances pointers
	__StateShape__000000_Probe_Form.State = __State__000000_Probe_Form
	__StateShape__000001_Probe_Tree.State = __State__000001_Probe_Tree
	__StateShape__000002_Probe_Table.State = __State__000002_Probe_Table
	__StateShape__000003_Form_Update_Stage_Of_Interest.State = __State__000003_Form_Update_Stage_Of_Interest
	__StateShape__000004_Enforce_Model_Semantic.State = __State__000004_Enforce_Model_Semantic
	__StateShape__000005_Update_Probe_Form.State = __State__000005_Update_Probe_Form
	__StateShape__000006_Update_Probe_Table.State = __State__000006_Update_Probe_Table
	__StateShape__000007_Update_Probe_Tree.State = __State__000007_Update_Probe_Tree
	__StateShape__000008_Load_Stage.State = __State__000008_Load_Stage
	__StateShape__000009_Persist_Stage.State = __State__000009_Persist_Stage
	__StateShape__000010_Stage_Modified_.State = __State__000010_Stage_Modified_
	__StateShape__000011_Update_UX_.State = __State__000011_Update_UX_
	__StateShape__000012_Clean_Stage.State = __State__000012_Clean_Stage
	// setup of Transition instances pointers
	__Transition__000000_Submit.Start = __State__000000_Probe_Form
	__Transition__000000_Submit.End = __State__000003_Form_Update_Stage_Of_Interest
	__Transition__000000_Submit.Diagrams = append(__Transition__000000_Submit.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__Transition__000001_.Start = __State__000003_Form_Update_Stage_Of_Interest
	__Transition__000001_.End = __State__000004_Enforce_Model_Semantic
	__Transition__000001_.Guard = __Guard__000003_not_a_suppress_action
	__Transition__000001_.Diagrams = append(__Transition__000001_.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__Transition__000002__Model_Updated.Start = __State__000004_Enforce_Model_Semantic
	__Transition__000002__Model_Updated.End = __State__000010_Stage_Modified_
	__Transition__000002__Model_Updated.Diagrams = append(__Transition__000002__Model_Updated.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__Transition__000003_UX_1_updated.Start = __State__000005_Update_Probe_Form
	__Transition__000003_UX_1_updated.End = __State__000006_Update_Probe_Table
	__Transition__000003_UX_1_updated.Diagrams = append(__Transition__000003_UX_1_updated.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__Transition__000004_UX_2_updated.Start = __State__000006_Update_Probe_Table
	__Transition__000004_UX_2_updated.End = __State__000007_Update_Probe_Tree
	__Transition__000004_UX_2_updated.Diagrams = append(__Transition__000004_UX_2_updated.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__Transition__000005__Stage_loaded.Start = __State__000008_Load_Stage
	__Transition__000005__Stage_loaded.End = __State__000004_Enforce_Model_Semantic
	__Transition__000005__Stage_loaded.Diagrams = append(__Transition__000005__Stage_loaded.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__Transition__000006__Stage_persisted.Start = __State__000009_Persist_Stage
	__Transition__000006__Stage_persisted.End = __State__000011_Update_UX_
	__Transition__000006__Stage_persisted.Diagrams = append(__Transition__000006__Stage_persisted.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__Transition__000007_Yes.Start = __State__000010_Stage_Modified_
	__Transition__000007_Yes.End = __State__000009_Persist_Stage
	__Transition__000007_Yes.Diagrams = append(__Transition__000007_Yes.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__Transition__000008_Yes.Start = __State__000011_Update_UX_
	__Transition__000008_Yes.End = __State__000005_Update_Probe_Form
	__Transition__000008_Yes.Diagrams = append(__Transition__000008_Yes.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__Transition__000009_Form_Update_Stage_Of_Interest_to_Clean_Stage.Start = __State__000003_Form_Update_Stage_Of_Interest
	__Transition__000009_Form_Update_Stage_Of_Interest_to_Clean_Stage.End = __State__000012_Clean_Stage
	__Transition__000009_Form_Update_Stage_Of_Interest_to_Clean_Stage.Guard = __Guard__000002_a_suppress_action
	__Transition__000009_Form_Update_Stage_Of_Interest_to_Clean_Stage.Diagrams = append(__Transition__000009_Form_Update_Stage_Of_Interest_to_Clean_Stage.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__Transition__000010_.Start = __State__000012_Clean_Stage
	__Transition__000010_.End = __State__000004_Enforce_Model_Semantic
	__Transition__000010_.Diagrams = append(__Transition__000010_.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	// setup of Transition_Shape instances pointers
	__Transition_Shape__000000_UX_1_Waiting_for_User_Input_to_UX_1_Update_Stage.Transition = __Transition__000000_Submit
	__Transition_Shape__000001_UX_1_Update_Stage_to_Enforce_Model_Semantic.Transition = __Transition__000001_
	__Transition_Shape__000002_Enforce_Model_Semantic_to_UX_1_Update_UX.Transition = __Transition__000002__Model_Updated
	__Transition_Shape__000003_Enforce_Model_Semantic_to_UX_2_Update_UX.Transition = __Transition__000003_UX_1_updated
	__Transition_Shape__000004_Enforce_Model_Semantic_to_UX_n_Update_UX.Transition = __Transition__000004_UX_2_updated
	__Transition_Shape__000005_Load_Stage_to_Enforce_Model_Semantic.Transition = __Transition__000005__Stage_loaded
	__Transition_Shape__000006_Persist_Stage_to_UX_1_Update_UX.Transition = __Transition__000006__Stage_persisted
	__Transition_Shape__000007_Stage_Modified_to_Persist_Stage.Transition = __Transition__000007_Yes
	__Transition_Shape__000008_Update_UX_to_UX_1_Update_UX.Transition = __Transition__000008_Yes
	__Transition_Shape__000009_Form_Update_Stage_Of_Interest_to_Clean_Stage.Transition = __Transition__000009_Form_Update_Stage_Of_Interest_to_Clean_Stage
	__Transition_Shape__000010_Clean_Stage_to_Enforce_Model_Semantic.Transition = __Transition__000010_
}

