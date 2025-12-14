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

	__Architecture__000000_Gong_UX_loop_Architecture := (&models.Architecture{}).Stage(stage)

	__Diagram__000000_UX_Loop_Diagram := (&models.Diagram{}).Stage(stage)

	__DoAction__000000_Lock_Stage := (&models.DoAction{}).Stage(stage)
	__DoAction__000001_Unlock_Stage := (&models.DoAction{}).Stage(stage)
	__DoAction__000002_RLock_Stage := (&models.DoAction{}).Stage(stage)
	__DoAction__000003_RUnlock_Stage := (&models.DoAction{}).Stage(stage)

	__State__000000_UX_1_Waiting_for_User_Input := (&models.State{}).Stage(stage)
	__State__000001_UX_2_Waiting_for_User_Input := (&models.State{}).Stage(stage)
	__State__000002_UX_n_Waiting_for_User_Input := (&models.State{}).Stage(stage)
	__State__000003_UX_1_Update_Stage := (&models.State{}).Stage(stage)
	__State__000004_UX_2_Update_Stage := (&models.State{}).Stage(stage)
	__State__000005_UX_n_Update_Stage := (&models.State{}).Stage(stage)
	__State__000006_Enforce_Model_Semantic := (&models.State{}).Stage(stage)
	__State__000007_UX_1_Update_UX := (&models.State{}).Stage(stage)
	__State__000008_UX_2_Update_UX := (&models.State{}).Stage(stage)
	__State__000009_UX_n_Update_UX := (&models.State{}).Stage(stage)
	__State__000010_Waiting_for_user_Input_ := (&models.State{}).Stage(stage)
	__State__000011_Load_Stage := (&models.State{}).Stage(stage)
	__State__000012_Persist_Stage := (&models.State{}).Stage(stage)
	__State__000013_Stage_Modified_ := (&models.State{}).Stage(stage)
	__State__000014_Update_UX_ := (&models.State{}).Stage(stage)

	__StateMachine__000000_UX_Loop := (&models.StateMachine{}).Stage(stage)

	__StateShape__000000_UX_1_Waiting_for_User_Input := (&models.StateShape{}).Stage(stage)
	__StateShape__000001_UX_2_Waiting_for_User_Input := (&models.StateShape{}).Stage(stage)
	__StateShape__000002_UX_n_Waiting_for_User_Input := (&models.StateShape{}).Stage(stage)
	__StateShape__000003_UX_1_Update_Stage := (&models.StateShape{}).Stage(stage)
	__StateShape__000004_UX_2_Update_Stage := (&models.StateShape{}).Stage(stage)
	__StateShape__000005_UX_n_Update_Stage := (&models.StateShape{}).Stage(stage)
	__StateShape__000006_Enforce_Model_Semantic := (&models.StateShape{}).Stage(stage)
	__StateShape__000007_UX_1_Update_UX := (&models.StateShape{}).Stage(stage)
	__StateShape__000008_UX_2_Update_UX := (&models.StateShape{}).Stage(stage)
	__StateShape__000009_UX_n_Update_UX := (&models.StateShape{}).Stage(stage)
	__StateShape__000010_Waiting_for_user_Input_ := (&models.StateShape{}).Stage(stage)
	__StateShape__000011_Load_Stage := (&models.StateShape{}).Stage(stage)
	__StateShape__000012_Persist_Stage := (&models.StateShape{}).Stage(stage)
	__StateShape__000013_Stage_Modified_ := (&models.StateShape{}).Stage(stage)
	__StateShape__000014_Update_UX_ := (&models.StateShape{}).Stage(stage)

	__Transition__000000__User_Input := (&models.Transition{}).Stage(stage)
	__Transition__000001__User_Input := (&models.Transition{}).Stage(stage)
	__Transition__000002_User_Input := (&models.Transition{}).Stage(stage)
	__Transition__000003_Stage_updated := (&models.Transition{}).Stage(stage)
	__Transition__000004_Stage_updated := (&models.Transition{}).Stage(stage)
	__Transition__000005_Stage_updated := (&models.Transition{}).Stage(stage)
	__Transition__000006__Model_Updated := (&models.Transition{}).Stage(stage)
	__Transition__000007_UX_1_updated := (&models.Transition{}).Stage(stage)
	__Transition__000008_UX_2_updated := (&models.Transition{}).Stage(stage)
	__Transition__000009_UX_n_updated := (&models.Transition{}).Stage(stage)
	__Transition__000010__Stage_loaded := (&models.Transition{}).Stage(stage)
	__Transition__000011__Stage_persisted := (&models.Transition{}).Stage(stage)
	__Transition__000012_No := (&models.Transition{}).Stage(stage)
	__Transition__000013_Yes := (&models.Transition{}).Stage(stage)
	__Transition__000014_Yes := (&models.Transition{}).Stage(stage)
	__Transition__000015_No := (&models.Transition{}).Stage(stage)

	__Transition_Shape__000000_UX_1_Waiting_for_User_Input_to_UX_1_Update_Stage := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000001_UX_2_Waiting_for_User_Input_to_UX_2_Update_Stage := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000002_UX_n_Waiting_for_User_Input_to_UX_n_Update_Stage := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000003_UX_1_Update_Stage_to_Enforce_Model_Semantic := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000004_UX_2_Update_Stage_to_Enforce_Model_Semantic := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000005_UX_n_Update_Stage_to_Enforce_Model_Semantic := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000006_Enforce_Model_Semantic_to_UX_1_Update_UX := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000007_Enforce_Model_Semantic_to_UX_2_Update_UX := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000008_Enforce_Model_Semantic_to_UX_n_Update_UX := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000009_UX_n_Update_UX_to_Waiting_for_user_Input_ := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000010_Load_Stage_to_Enforce_Model_Semantic := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000011_Persist_Stage_to_UX_1_Update_UX := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000012_Stage_Modified_to_UX_1_Update_UX := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000013_Stage_Modified_to_Persist_Stage := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000014_Update_UX_to_UX_1_Update_UX := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000015_Update_UX_to_Waiting_for_user_Input_ := (&models.Transition_Shape{}).Stage(stage)

	// Setup of values

	__Architecture__000000_Gong_UX_loop_Architecture.Name = `Gong UX loop Architecture`
	__Architecture__000000_Gong_UX_loop_Architecture.NbPixPerCharacter = 8.000000

	__Diagram__000000_UX_Loop_Diagram.Name = `UX Loop Diagram`
	__Diagram__000000_UX_Loop_Diagram.IsChecked = true
	__Diagram__000000_UX_Loop_Diagram.IsExpanded = true
	__Diagram__000000_UX_Loop_Diagram.IsEditable_ = true
	__Diagram__000000_UX_Loop_Diagram.IsInRenameMode = false

	__DoAction__000000_Lock_Stage.Name = `Lock Stage`
	__DoAction__000000_Lock_Stage.Criticality = models.DoActionCritical

	__DoAction__000001_Unlock_Stage.Name = `Unlock Stage`
	__DoAction__000001_Unlock_Stage.Criticality = models.DoActionCritical

	__DoAction__000002_RLock_Stage.Name = `RLock Stage`
	__DoAction__000002_RLock_Stage.Criticality = models.DoActionCritical

	__DoAction__000003_RUnlock_Stage.Name = `RUnlock Stage`
	__DoAction__000003_RUnlock_Stage.Criticality = models.DoActionCritical

	__State__000000_UX_1_Waiting_for_User_Input.Name = `UX 1 - Waiting for User Input`
	__State__000000_UX_1_Waiting_for_User_Input.IsDecisionNode = false
	__State__000000_UX_1_Waiting_for_User_Input.IsFictif = false
	__State__000000_UX_1_Waiting_for_User_Input.IsEndState = false

	__State__000001_UX_2_Waiting_for_User_Input.Name = `UX 2 - Waiting for User Input`
	__State__000001_UX_2_Waiting_for_User_Input.IsDecisionNode = false
	__State__000001_UX_2_Waiting_for_User_Input.IsFictif = false
	__State__000001_UX_2_Waiting_for_User_Input.IsEndState = false

	__State__000002_UX_n_Waiting_for_User_Input.Name = `UX n - Waiting for User Input`
	__State__000002_UX_n_Waiting_for_User_Input.IsDecisionNode = false
	__State__000002_UX_n_Waiting_for_User_Input.IsFictif = false
	__State__000002_UX_n_Waiting_for_User_Input.IsEndState = false

	__State__000003_UX_1_Update_Stage.Name = `UX 1 - Update Stage`
	__State__000003_UX_1_Update_Stage.IsDecisionNode = false
	__State__000003_UX_1_Update_Stage.IsFictif = false
	__State__000003_UX_1_Update_Stage.IsEndState = false

	__State__000004_UX_2_Update_Stage.Name = `UX 2 - Update Stage`
	__State__000004_UX_2_Update_Stage.IsDecisionNode = false
	__State__000004_UX_2_Update_Stage.IsFictif = false
	__State__000004_UX_2_Update_Stage.IsEndState = false

	__State__000005_UX_n_Update_Stage.Name = `UX n - Update Stage`
	__State__000005_UX_n_Update_Stage.IsDecisionNode = false
	__State__000005_UX_n_Update_Stage.IsFictif = false
	__State__000005_UX_n_Update_Stage.IsEndState = false

	__State__000006_Enforce_Model_Semantic.Name = `Enforce Model Semantic`
	__State__000006_Enforce_Model_Semantic.IsDecisionNode = false
	__State__000006_Enforce_Model_Semantic.IsFictif = false
	__State__000006_Enforce_Model_Semantic.IsEndState = false

	__State__000007_UX_1_Update_UX.Name = `UX 1 - Update UX`
	__State__000007_UX_1_Update_UX.IsDecisionNode = false
	__State__000007_UX_1_Update_UX.IsFictif = false
	__State__000007_UX_1_Update_UX.IsEndState = false

	__State__000008_UX_2_Update_UX.Name = `UX 2 - Update UX`
	__State__000008_UX_2_Update_UX.IsDecisionNode = false
	__State__000008_UX_2_Update_UX.IsFictif = false
	__State__000008_UX_2_Update_UX.IsEndState = false

	__State__000009_UX_n_Update_UX.Name = `UX n - Update UX`
	__State__000009_UX_n_Update_UX.IsDecisionNode = false
	__State__000009_UX_n_Update_UX.IsFictif = false
	__State__000009_UX_n_Update_UX.IsEndState = false

	__State__000010_Waiting_for_user_Input_.Name = `Waiting for user Input `
	__State__000010_Waiting_for_user_Input_.IsDecisionNode = false
	__State__000010_Waiting_for_user_Input_.IsFictif = false
	__State__000010_Waiting_for_user_Input_.IsEndState = false

	__State__000011_Load_Stage.Name = `Load Stage`
	__State__000011_Load_Stage.IsDecisionNode = false
	__State__000011_Load_Stage.IsFictif = false
	__State__000011_Load_Stage.IsEndState = false

	__State__000012_Persist_Stage.Name = `Persist Stage`
	__State__000012_Persist_Stage.IsDecisionNode = false
	__State__000012_Persist_Stage.IsFictif = false
	__State__000012_Persist_Stage.IsEndState = false

	__State__000013_Stage_Modified_.Name = `Stage Modified ?`
	__State__000013_Stage_Modified_.IsDecisionNode = true
	__State__000013_Stage_Modified_.IsFictif = true
	__State__000013_Stage_Modified_.IsEndState = false

	__State__000014_Update_UX_.Name = `Update UX ?`
	__State__000014_Update_UX_.IsDecisionNode = true
	__State__000014_Update_UX_.IsFictif = false
	__State__000014_Update_UX_.IsEndState = false

	__StateMachine__000000_UX_Loop.Name = `UX Loop`
	__StateMachine__000000_UX_Loop.IsNodeExpanded = true

	__StateShape__000000_UX_1_Waiting_for_User_Input.Name = `UX 1 - Waiting for User Input`
	__StateShape__000000_UX_1_Waiting_for_User_Input.IsExpanded = false
	__StateShape__000000_UX_1_Waiting_for_User_Input.X = 359.000000
	__StateShape__000000_UX_1_Waiting_for_User_Input.Y = 154.000000
	__StateShape__000000_UX_1_Waiting_for_User_Input.Width = 221.000000
	__StateShape__000000_UX_1_Waiting_for_User_Input.Height = 80.000000

	__StateShape__000001_UX_2_Waiting_for_User_Input.Name = `UX 2 - Waiting for User Input`
	__StateShape__000001_UX_2_Waiting_for_User_Input.IsExpanded = false
	__StateShape__000001_UX_2_Waiting_for_User_Input.X = 601.000000
	__StateShape__000001_UX_2_Waiting_for_User_Input.Y = 152.000000
	__StateShape__000001_UX_2_Waiting_for_User_Input.Width = 200.000000
	__StateShape__000001_UX_2_Waiting_for_User_Input.Height = 80.000000

	__StateShape__000002_UX_n_Waiting_for_User_Input.Name = `UX n - Waiting for User Input`
	__StateShape__000002_UX_n_Waiting_for_User_Input.IsExpanded = false
	__StateShape__000002_UX_n_Waiting_for_User_Input.X = 821.000000
	__StateShape__000002_UX_n_Waiting_for_User_Input.Y = 151.000000
	__StateShape__000002_UX_n_Waiting_for_User_Input.Width = 200.000000
	__StateShape__000002_UX_n_Waiting_for_User_Input.Height = 80.000000

	__StateShape__000003_UX_1_Update_Stage.Name = `UX 1 - Update Stage`
	__StateShape__000003_UX_1_Update_Stage.IsExpanded = false
	__StateShape__000003_UX_1_Update_Stage.X = 379.000000
	__StateShape__000003_UX_1_Update_Stage.Y = 381.999985
	__StateShape__000003_UX_1_Update_Stage.Width = 200.000000
	__StateShape__000003_UX_1_Update_Stage.Height = 80.000000

	__StateShape__000004_UX_2_Update_Stage.Name = `UX 2 - Update Stage`
	__StateShape__000004_UX_2_Update_Stage.IsExpanded = false
	__StateShape__000004_UX_2_Update_Stage.X = 606.000000
	__StateShape__000004_UX_2_Update_Stage.Y = 380.999985
	__StateShape__000004_UX_2_Update_Stage.Width = 200.000000
	__StateShape__000004_UX_2_Update_Stage.Height = 80.000000

	__StateShape__000005_UX_n_Update_Stage.Name = `UX n - Update Stage`
	__StateShape__000005_UX_n_Update_Stage.IsExpanded = false
	__StateShape__000005_UX_n_Update_Stage.X = 827.000000
	__StateShape__000005_UX_n_Update_Stage.Y = 383.999985
	__StateShape__000005_UX_n_Update_Stage.Width = 200.000000
	__StateShape__000005_UX_n_Update_Stage.Height = 78.000000

	__StateShape__000006_Enforce_Model_Semantic.Name = `Enforce Model Semantic`
	__StateShape__000006_Enforce_Model_Semantic.IsExpanded = false
	__StateShape__000006_Enforce_Model_Semantic.X = 383.000000
	__StateShape__000006_Enforce_Model_Semantic.Y = 588.000015
	__StateShape__000006_Enforce_Model_Semantic.Width = 646.000000
	__StateShape__000006_Enforce_Model_Semantic.Height = 80.000000

	__StateShape__000007_UX_1_Update_UX.Name = `UX 1 - Update UX`
	__StateShape__000007_UX_1_Update_UX.IsExpanded = false
	__StateShape__000007_UX_1_Update_UX.X = 377.000000
	__StateShape__000007_UX_1_Update_UX.Y = 922.999954
	__StateShape__000007_UX_1_Update_UX.Width = 200.000000
	__StateShape__000007_UX_1_Update_UX.Height = 80.000000

	__StateShape__000008_UX_2_Update_UX.Name = `UX 2 - Update UX`
	__StateShape__000008_UX_2_Update_UX.IsExpanded = false
	__StateShape__000008_UX_2_Update_UX.X = 600.000000
	__StateShape__000008_UX_2_Update_UX.Y = 920.999954
	__StateShape__000008_UX_2_Update_UX.Width = 200.000000
	__StateShape__000008_UX_2_Update_UX.Height = 80.000000

	__StateShape__000009_UX_n_Update_UX.Name = `UX n - Update UX`
	__StateShape__000009_UX_n_Update_UX.IsExpanded = false
	__StateShape__000009_UX_n_Update_UX.X = 820.000000
	__StateShape__000009_UX_n_Update_UX.Y = 920.999954
	__StateShape__000009_UX_n_Update_UX.Width = 200.000000
	__StateShape__000009_UX_n_Update_UX.Height = 80.000000

	__StateShape__000010_Waiting_for_user_Input_.Name = `Waiting for user Input `
	__StateShape__000010_Waiting_for_user_Input_.IsExpanded = false
	__StateShape__000010_Waiting_for_user_Input_.X = 318.000000
	__StateShape__000010_Waiting_for_user_Input_.Y = 93.000000
	__StateShape__000010_Waiting_for_user_Input_.Width = 745.000000
	__StateShape__000010_Waiting_for_user_Input_.Height = 204.999985

	__StateShape__000011_Load_Stage.Name = `Load Stage`
	__StateShape__000011_Load_Stage.IsExpanded = false
	__StateShape__000011_Load_Stage.X = 63.000000
	__StateShape__000011_Load_Stage.Y = 617.000015
	__StateShape__000011_Load_Stage.Width = 107.000000
	__StateShape__000011_Load_Stage.Height = 20.000000

	__StateShape__000012_Persist_Stage.Name = `Persist Stage`
	__StateShape__000012_Persist_Stage.IsExpanded = false
	__StateShape__000012_Persist_Stage.X = 53.000000
	__StateShape__000012_Persist_Stage.Y = 803.999954
	__StateShape__000012_Persist_Stage.Width = 200.000000
	__StateShape__000012_Persist_Stage.Height = 80.000000

	__StateShape__000013_Stage_Modified_.Name = `Stage Modified ?`
	__StateShape__000013_Stage_Modified_.IsExpanded = false
	__StateShape__000013_Stage_Modified_.X = 586.000000
	__StateShape__000013_Stage_Modified_.Y = 736.999954
	__StateShape__000013_Stage_Modified_.Width = 153.000000
	__StateShape__000013_Stage_Modified_.Height = 50.000000

	__StateShape__000014_Update_UX_.Name = `Update UX ?`
	__StateShape__000014_Update_UX_.IsExpanded = false
	__StateShape__000014_Update_UX_.X = 106.000000
	__StateShape__000014_Update_UX_.Y = 1071.000015
	__StateShape__000014_Update_UX_.Width = 120.000000
	__StateShape__000014_Update_UX_.Height = 50.000000

	__Transition__000000__User_Input.Name = `/User Input`

	__Transition__000001__User_Input.Name = `/User Input`

	__Transition__000002_User_Input.Name = `User Input`

	__Transition__000003_Stage_updated.Name = `Stage updated`

	__Transition__000004_Stage_updated.Name = `Stage updated`

	__Transition__000005_Stage_updated.Name = `Stage updated`

	__Transition__000006__Model_Updated.Name = `/Model Updated`

	__Transition__000007_UX_1_updated.Name = `UX 1 updated`

	__Transition__000008_UX_2_updated.Name = `UX 2 updated`

	__Transition__000009_UX_n_updated.Name = `UX n updated`

	__Transition__000010__Stage_loaded.Name = `/Stage loaded`

	__Transition__000011__Stage_persisted.Name = `/Stage persisted`

	__Transition__000012_No.Name = `No`

	__Transition__000013_Yes.Name = `Yes`

	__Transition__000014_Yes.Name = `Yes`

	__Transition__000015_No.Name = `No`

	__Transition_Shape__000000_UX_1_Waiting_for_User_Input_to_UX_1_Update_Stage.Name = `UX 1 - Waiting for User Input to UX 1 - Update Stage`
	__Transition_Shape__000000_UX_1_Waiting_for_User_Input_to_UX_1_Update_Stage.StartRatio = 0.398958
	__Transition_Shape__000000_UX_1_Waiting_for_User_Input_to_UX_1_Update_Stage.EndRatio = 0.258932
	__Transition_Shape__000000_UX_1_Waiting_for_User_Input_to_UX_1_Update_Stage.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000000_UX_1_Waiting_for_User_Input_to_UX_1_Update_Stage.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000000_UX_1_Waiting_for_User_Input_to_UX_1_Update_Stage.CornerOffsetRatio = 1.500087

	__Transition_Shape__000001_UX_2_Waiting_for_User_Input_to_UX_2_Update_Stage.Name = `UX 2 - Waiting for User Input to UX 2 - Update Stage`
	__Transition_Shape__000001_UX_2_Waiting_for_User_Input_to_UX_2_Update_Stage.StartRatio = 0.583932
	__Transition_Shape__000001_UX_2_Waiting_for_User_Input_to_UX_2_Update_Stage.EndRatio = 0.473932
	__Transition_Shape__000001_UX_2_Waiting_for_User_Input_to_UX_2_Update_Stage.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000001_UX_2_Waiting_for_User_Input_to_UX_2_Update_Stage.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000001_UX_2_Waiting_for_User_Input_to_UX_2_Update_Stage.CornerOffsetRatio = 1.350087

	__Transition_Shape__000002_UX_n_Waiting_for_User_Input_to_UX_n_Update_Stage.Name = `UX n - Waiting for User Input to UX n - Update Stage`
	__Transition_Shape__000002_UX_n_Waiting_for_User_Input_to_UX_n_Update_Stage.StartRatio = 0.468932
	__Transition_Shape__000002_UX_n_Waiting_for_User_Input_to_UX_n_Update_Stage.EndRatio = 0.458932
	__Transition_Shape__000002_UX_n_Waiting_for_User_Input_to_UX_n_Update_Stage.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000002_UX_n_Waiting_for_User_Input_to_UX_n_Update_Stage.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000002_UX_n_Waiting_for_User_Input_to_UX_n_Update_Stage.CornerOffsetRatio = 1.275087

	__Transition_Shape__000003_UX_1_Update_Stage_to_Enforce_Model_Semantic.Name = `UX 1 - Update Stage to Enforce Model Semantic`
	__Transition_Shape__000003_UX_1_Update_Stage_to_Enforce_Model_Semantic.StartRatio = 0.558932
	__Transition_Shape__000003_UX_1_Update_Stage_to_Enforce_Model_Semantic.EndRatio = 0.166852
	__Transition_Shape__000003_UX_1_Update_Stage_to_Enforce_Model_Semantic.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000003_UX_1_Update_Stage_to_Enforce_Model_Semantic.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000003_UX_1_Update_Stage_to_Enforce_Model_Semantic.CornerOffsetRatio = 1.912587

	__Transition_Shape__000004_UX_2_Update_Stage_to_Enforce_Model_Semantic.Name = `UX 2 - Update Stage to Enforce Model Semantic`
	__Transition_Shape__000004_UX_2_Update_Stage_to_Enforce_Model_Semantic.StartRatio = 0.513932
	__Transition_Shape__000004_UX_2_Update_Stage_to_Enforce_Model_Semantic.EndRatio = 0.548932
	__Transition_Shape__000004_UX_2_Update_Stage_to_Enforce_Model_Semantic.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000004_UX_2_Update_Stage_to_Enforce_Model_Semantic.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000004_UX_2_Update_Stage_to_Enforce_Model_Semantic.CornerOffsetRatio = 1.887587

	__Transition_Shape__000005_UX_n_Update_Stage_to_Enforce_Model_Semantic.Name = `UX n - Update Stage to Enforce Model Semantic`
	__Transition_Shape__000005_UX_n_Update_Stage_to_Enforce_Model_Semantic.StartRatio = 0.478932
	__Transition_Shape__000005_UX_n_Update_Stage_to_Enforce_Model_Semantic.EndRatio = 0.792239
	__Transition_Shape__000005_UX_n_Update_Stage_to_Enforce_Model_Semantic.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000005_UX_n_Update_Stage_to_Enforce_Model_Semantic.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000005_UX_n_Update_Stage_to_Enforce_Model_Semantic.CornerOffsetRatio = 1.887587

	__Transition_Shape__000006_Enforce_Model_Semantic_to_UX_1_Update_UX.Name = `Enforce Model Semantic to UX 1 - Update UX`
	__Transition_Shape__000006_Enforce_Model_Semantic_to_UX_1_Update_UX.StartRatio = 0.512053
	__Transition_Shape__000006_Enforce_Model_Semantic_to_UX_1_Update_UX.EndRatio = 0.831915
	__Transition_Shape__000006_Enforce_Model_Semantic_to_UX_1_Update_UX.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000006_Enforce_Model_Semantic_to_UX_1_Update_UX.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000006_Enforce_Model_Semantic_to_UX_1_Update_UX.CornerOffsetRatio = 1.612587

	__Transition_Shape__000007_Enforce_Model_Semantic_to_UX_2_Update_UX.Name = `Enforce Model Semantic to UX 2 - Update UX`
	__Transition_Shape__000007_Enforce_Model_Semantic_to_UX_2_Update_UX.StartRatio = 0.821415
	__Transition_Shape__000007_Enforce_Model_Semantic_to_UX_2_Update_UX.EndRatio = 0.166415
	__Transition_Shape__000007_Enforce_Model_Semantic_to_UX_2_Update_UX.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000007_Enforce_Model_Semantic_to_UX_2_Update_UX.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000007_Enforce_Model_Semantic_to_UX_2_Update_UX.CornerOffsetRatio = 1.625087

	__Transition_Shape__000008_Enforce_Model_Semantic_to_UX_n_Update_UX.Name = `Enforce Model Semantic to UX n - Update UX`
	__Transition_Shape__000008_Enforce_Model_Semantic_to_UX_n_Update_UX.StartRatio = 0.830939
	__Transition_Shape__000008_Enforce_Model_Semantic_to_UX_n_Update_UX.EndRatio = 0.801415
	__Transition_Shape__000008_Enforce_Model_Semantic_to_UX_n_Update_UX.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000008_Enforce_Model_Semantic_to_UX_n_Update_UX.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000008_Enforce_Model_Semantic_to_UX_n_Update_UX.CornerOffsetRatio = 1.662587

	__Transition_Shape__000009_UX_n_Update_UX_to_Waiting_for_user_Input_.Name = `UX n - Update UX to Waiting for user Input `
	__Transition_Shape__000009_UX_n_Update_UX_to_Waiting_for_user_Input_.StartRatio = 0.500000
	__Transition_Shape__000009_UX_n_Update_UX_to_Waiting_for_user_Input_.EndRatio = 0.443936
	__Transition_Shape__000009_UX_n_Update_UX_to_Waiting_for_user_Input_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__000009_UX_n_Update_UX_to_Waiting_for_user_Input_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__000009_UX_n_Update_UX_to_Waiting_for_user_Input_.CornerOffsetRatio = 1.971415

	__Transition_Shape__000010_Load_Stage_to_Enforce_Model_Semantic.Name = `Load Stage to Enforce Model Semantic`
	__Transition_Shape__000010_Load_Stage_to_Enforce_Model_Semantic.StartRatio = 0.550347
	__Transition_Shape__000010_Load_Stage_to_Enforce_Model_Semantic.EndRatio = 0.500000
	__Transition_Shape__000010_Load_Stage_to_Enforce_Model_Semantic.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__000010_Load_Stage_to_Enforce_Model_Semantic.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__000010_Load_Stage_to_Enforce_Model_Semantic.CornerOffsetRatio = 1.544701

	__Transition_Shape__000011_Persist_Stage_to_UX_1_Update_UX.Name = `Persist Stage to UX 1 - Update UX`
	__Transition_Shape__000011_Persist_Stage_to_UX_1_Update_UX.StartRatio = 0.741415
	__Transition_Shape__000011_Persist_Stage_to_UX_1_Update_UX.EndRatio = 0.760691
	__Transition_Shape__000011_Persist_Stage_to_UX_1_Update_UX.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000011_Persist_Stage_to_UX_1_Update_UX.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000011_Persist_Stage_to_UX_1_Update_UX.CornerOffsetRatio = 1.737587

	__Transition_Shape__000012_Stage_Modified_to_UX_1_Update_UX.Name = `Stage Modified ? to UX 1 - Update UX`
	__Transition_Shape__000012_Stage_Modified_to_UX_1_Update_UX.StartRatio = 0.500000
	__Transition_Shape__000012_Stage_Modified_to_UX_1_Update_UX.EndRatio = 0.643936
	__Transition_Shape__000012_Stage_Modified_to_UX_1_Update_UX.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__000012_Stage_Modified_to_UX_1_Update_UX.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__000012_Stage_Modified_to_UX_1_Update_UX.CornerOffsetRatio = 3.890738

	__Transition_Shape__000013_Stage_Modified_to_Persist_Stage.Name = `Stage Modified ? to Persist Stage`
	__Transition_Shape__000013_Stage_Modified_to_Persist_Stage.StartRatio = 0.831915
	__Transition_Shape__000013_Stage_Modified_to_Persist_Stage.EndRatio = 0.350087
	__Transition_Shape__000013_Stage_Modified_to_Persist_Stage.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000013_Stage_Modified_to_Persist_Stage.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__000013_Stage_Modified_to_Persist_Stage.CornerOffsetRatio = 1.100139

	__Transition_Shape__000014_Update_UX_to_UX_1_Update_UX.Name = `Update UX ? to UX 1 - Update UX`
	__Transition_Shape__000014_Update_UX_to_UX_1_Update_UX.StartRatio = 0.120139
	__Transition_Shape__000014_Update_UX_to_UX_1_Update_UX.EndRatio = 0.500000
	__Transition_Shape__000014_Update_UX_to_UX_1_Update_UX.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__000014_Update_UX_to_UX_1_Update_UX.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__000014_Update_UX_to_UX_1_Update_UX.CornerOffsetRatio = 1.560691

	__Transition_Shape__000015_Update_UX_to_Waiting_for_user_Input_.Name = `Update UX ? to Waiting for user Input `
	__Transition_Shape__000015_Update_UX_to_Waiting_for_user_Input_.StartRatio = 0.980139
	__Transition_Shape__000015_Update_UX_to_Waiting_for_user_Input_.EndRatio = 0.234180
	__Transition_Shape__000015_Update_UX_to_Waiting_for_user_Input_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__000015_Update_UX_to_Waiting_for_user_Input_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__000015_Update_UX_to_Waiting_for_user_Input_.CornerOffsetRatio = 9.527358

	// Setup of pointers
	// setup of Architecture instances pointers
	__Architecture__000000_Gong_UX_loop_Architecture.StateMachines = append(__Architecture__000000_Gong_UX_loop_Architecture.StateMachines, __StateMachine__000000_UX_Loop)
	// setup of Diagram instances pointers
	__Diagram__000000_UX_Loop_Diagram.State_Shapes = append(__Diagram__000000_UX_Loop_Diagram.State_Shapes, __StateShape__000010_Waiting_for_user_Input_)
	__Diagram__000000_UX_Loop_Diagram.State_Shapes = append(__Diagram__000000_UX_Loop_Diagram.State_Shapes, __StateShape__000000_UX_1_Waiting_for_User_Input)
	__Diagram__000000_UX_Loop_Diagram.State_Shapes = append(__Diagram__000000_UX_Loop_Diagram.State_Shapes, __StateShape__000001_UX_2_Waiting_for_User_Input)
	__Diagram__000000_UX_Loop_Diagram.State_Shapes = append(__Diagram__000000_UX_Loop_Diagram.State_Shapes, __StateShape__000002_UX_n_Waiting_for_User_Input)
	__Diagram__000000_UX_Loop_Diagram.State_Shapes = append(__Diagram__000000_UX_Loop_Diagram.State_Shapes, __StateShape__000003_UX_1_Update_Stage)
	__Diagram__000000_UX_Loop_Diagram.State_Shapes = append(__Diagram__000000_UX_Loop_Diagram.State_Shapes, __StateShape__000004_UX_2_Update_Stage)
	__Diagram__000000_UX_Loop_Diagram.State_Shapes = append(__Diagram__000000_UX_Loop_Diagram.State_Shapes, __StateShape__000005_UX_n_Update_Stage)
	__Diagram__000000_UX_Loop_Diagram.State_Shapes = append(__Diagram__000000_UX_Loop_Diagram.State_Shapes, __StateShape__000006_Enforce_Model_Semantic)
	__Diagram__000000_UX_Loop_Diagram.State_Shapes = append(__Diagram__000000_UX_Loop_Diagram.State_Shapes, __StateShape__000007_UX_1_Update_UX)
	__Diagram__000000_UX_Loop_Diagram.State_Shapes = append(__Diagram__000000_UX_Loop_Diagram.State_Shapes, __StateShape__000008_UX_2_Update_UX)
	__Diagram__000000_UX_Loop_Diagram.State_Shapes = append(__Diagram__000000_UX_Loop_Diagram.State_Shapes, __StateShape__000009_UX_n_Update_UX)
	__Diagram__000000_UX_Loop_Diagram.State_Shapes = append(__Diagram__000000_UX_Loop_Diagram.State_Shapes, __StateShape__000011_Load_Stage)
	__Diagram__000000_UX_Loop_Diagram.State_Shapes = append(__Diagram__000000_UX_Loop_Diagram.State_Shapes, __StateShape__000012_Persist_Stage)
	__Diagram__000000_UX_Loop_Diagram.State_Shapes = append(__Diagram__000000_UX_Loop_Diagram.State_Shapes, __StateShape__000013_Stage_Modified_)
	__Diagram__000000_UX_Loop_Diagram.State_Shapes = append(__Diagram__000000_UX_Loop_Diagram.State_Shapes, __StateShape__000014_Update_UX_)
	__Diagram__000000_UX_Loop_Diagram.Transition_Shapes = append(__Diagram__000000_UX_Loop_Diagram.Transition_Shapes, __Transition_Shape__000000_UX_1_Waiting_for_User_Input_to_UX_1_Update_Stage)
	__Diagram__000000_UX_Loop_Diagram.Transition_Shapes = append(__Diagram__000000_UX_Loop_Diagram.Transition_Shapes, __Transition_Shape__000001_UX_2_Waiting_for_User_Input_to_UX_2_Update_Stage)
	__Diagram__000000_UX_Loop_Diagram.Transition_Shapes = append(__Diagram__000000_UX_Loop_Diagram.Transition_Shapes, __Transition_Shape__000002_UX_n_Waiting_for_User_Input_to_UX_n_Update_Stage)
	__Diagram__000000_UX_Loop_Diagram.Transition_Shapes = append(__Diagram__000000_UX_Loop_Diagram.Transition_Shapes, __Transition_Shape__000003_UX_1_Update_Stage_to_Enforce_Model_Semantic)
	__Diagram__000000_UX_Loop_Diagram.Transition_Shapes = append(__Diagram__000000_UX_Loop_Diagram.Transition_Shapes, __Transition_Shape__000004_UX_2_Update_Stage_to_Enforce_Model_Semantic)
	__Diagram__000000_UX_Loop_Diagram.Transition_Shapes = append(__Diagram__000000_UX_Loop_Diagram.Transition_Shapes, __Transition_Shape__000005_UX_n_Update_Stage_to_Enforce_Model_Semantic)
	__Diagram__000000_UX_Loop_Diagram.Transition_Shapes = append(__Diagram__000000_UX_Loop_Diagram.Transition_Shapes, __Transition_Shape__000006_Enforce_Model_Semantic_to_UX_1_Update_UX)
	__Diagram__000000_UX_Loop_Diagram.Transition_Shapes = append(__Diagram__000000_UX_Loop_Diagram.Transition_Shapes, __Transition_Shape__000007_Enforce_Model_Semantic_to_UX_2_Update_UX)
	__Diagram__000000_UX_Loop_Diagram.Transition_Shapes = append(__Diagram__000000_UX_Loop_Diagram.Transition_Shapes, __Transition_Shape__000008_Enforce_Model_Semantic_to_UX_n_Update_UX)
	__Diagram__000000_UX_Loop_Diagram.Transition_Shapes = append(__Diagram__000000_UX_Loop_Diagram.Transition_Shapes, __Transition_Shape__000009_UX_n_Update_UX_to_Waiting_for_user_Input_)
	__Diagram__000000_UX_Loop_Diagram.Transition_Shapes = append(__Diagram__000000_UX_Loop_Diagram.Transition_Shapes, __Transition_Shape__000010_Load_Stage_to_Enforce_Model_Semantic)
	__Diagram__000000_UX_Loop_Diagram.Transition_Shapes = append(__Diagram__000000_UX_Loop_Diagram.Transition_Shapes, __Transition_Shape__000011_Persist_Stage_to_UX_1_Update_UX)
	__Diagram__000000_UX_Loop_Diagram.Transition_Shapes = append(__Diagram__000000_UX_Loop_Diagram.Transition_Shapes, __Transition_Shape__000012_Stage_Modified_to_UX_1_Update_UX)
	__Diagram__000000_UX_Loop_Diagram.Transition_Shapes = append(__Diagram__000000_UX_Loop_Diagram.Transition_Shapes, __Transition_Shape__000013_Stage_Modified_to_Persist_Stage)
	__Diagram__000000_UX_Loop_Diagram.Transition_Shapes = append(__Diagram__000000_UX_Loop_Diagram.Transition_Shapes, __Transition_Shape__000014_Update_UX_to_UX_1_Update_UX)
	__Diagram__000000_UX_Loop_Diagram.Transition_Shapes = append(__Diagram__000000_UX_Loop_Diagram.Transition_Shapes, __Transition_Shape__000015_Update_UX_to_Waiting_for_user_Input_)
	// setup of DoAction instances pointers
	// setup of State instances pointers
	__State__000000_UX_1_Waiting_for_User_Input.Diagrams = append(__State__000000_UX_1_Waiting_for_User_Input.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000001_UX_2_Waiting_for_User_Input.Diagrams = append(__State__000001_UX_2_Waiting_for_User_Input.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000002_UX_n_Waiting_for_User_Input.Diagrams = append(__State__000002_UX_n_Waiting_for_User_Input.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000003_UX_1_Update_Stage.Diagrams = append(__State__000003_UX_1_Update_Stage.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000003_UX_1_Update_Stage.DoActions = append(__State__000003_UX_1_Update_Stage.DoActions, __DoAction__000000_Lock_Stage)
	__State__000003_UX_1_Update_Stage.DoActions = append(__State__000003_UX_1_Update_Stage.DoActions, __DoAction__000001_Unlock_Stage)
	__State__000004_UX_2_Update_Stage.Diagrams = append(__State__000004_UX_2_Update_Stage.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000004_UX_2_Update_Stage.DoActions = append(__State__000004_UX_2_Update_Stage.DoActions, __DoAction__000000_Lock_Stage)
	__State__000004_UX_2_Update_Stage.DoActions = append(__State__000004_UX_2_Update_Stage.DoActions, __DoAction__000001_Unlock_Stage)
	__State__000005_UX_n_Update_Stage.Diagrams = append(__State__000005_UX_n_Update_Stage.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000005_UX_n_Update_Stage.DoActions = append(__State__000005_UX_n_Update_Stage.DoActions, __DoAction__000000_Lock_Stage)
	__State__000005_UX_n_Update_Stage.DoActions = append(__State__000005_UX_n_Update_Stage.DoActions, __DoAction__000001_Unlock_Stage)
	__State__000006_Enforce_Model_Semantic.Diagrams = append(__State__000006_Enforce_Model_Semantic.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000006_Enforce_Model_Semantic.DoActions = append(__State__000006_Enforce_Model_Semantic.DoActions, __DoAction__000000_Lock_Stage)
	__State__000006_Enforce_Model_Semantic.DoActions = append(__State__000006_Enforce_Model_Semantic.DoActions, __DoAction__000001_Unlock_Stage)
	__State__000007_UX_1_Update_UX.Diagrams = append(__State__000007_UX_1_Update_UX.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000007_UX_1_Update_UX.DoActions = append(__State__000007_UX_1_Update_UX.DoActions, __DoAction__000002_RLock_Stage)
	__State__000007_UX_1_Update_UX.DoActions = append(__State__000007_UX_1_Update_UX.DoActions, __DoAction__000003_RUnlock_Stage)
	__State__000008_UX_2_Update_UX.Diagrams = append(__State__000008_UX_2_Update_UX.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000008_UX_2_Update_UX.DoActions = append(__State__000008_UX_2_Update_UX.DoActions, __DoAction__000002_RLock_Stage)
	__State__000008_UX_2_Update_UX.DoActions = append(__State__000008_UX_2_Update_UX.DoActions, __DoAction__000003_RUnlock_Stage)
	__State__000009_UX_n_Update_UX.Diagrams = append(__State__000009_UX_n_Update_UX.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000009_UX_n_Update_UX.DoActions = append(__State__000009_UX_n_Update_UX.DoActions, __DoAction__000002_RLock_Stage)
	__State__000009_UX_n_Update_UX.DoActions = append(__State__000009_UX_n_Update_UX.DoActions, __DoAction__000003_RUnlock_Stage)
	__State__000010_Waiting_for_user_Input_.Diagrams = append(__State__000010_Waiting_for_user_Input_.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000011_Load_Stage.Diagrams = append(__State__000011_Load_Stage.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000012_Persist_Stage.Diagrams = append(__State__000012_Persist_Stage.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000012_Persist_Stage.DoActions = append(__State__000012_Persist_Stage.DoActions, __DoAction__000002_RLock_Stage)
	__State__000012_Persist_Stage.DoActions = append(__State__000012_Persist_Stage.DoActions, __DoAction__000003_RUnlock_Stage)
	__State__000013_Stage_Modified_.Diagrams = append(__State__000013_Stage_Modified_.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000014_Update_UX_.Diagrams = append(__State__000014_Update_UX_.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	// setup of StateMachine instances pointers
	__StateMachine__000000_UX_Loop.States = append(__StateMachine__000000_UX_Loop.States, __State__000000_UX_1_Waiting_for_User_Input)
	__StateMachine__000000_UX_Loop.States = append(__StateMachine__000000_UX_Loop.States, __State__000001_UX_2_Waiting_for_User_Input)
	__StateMachine__000000_UX_Loop.States = append(__StateMachine__000000_UX_Loop.States, __State__000002_UX_n_Waiting_for_User_Input)
	__StateMachine__000000_UX_Loop.States = append(__StateMachine__000000_UX_Loop.States, __State__000003_UX_1_Update_Stage)
	__StateMachine__000000_UX_Loop.States = append(__StateMachine__000000_UX_Loop.States, __State__000004_UX_2_Update_Stage)
	__StateMachine__000000_UX_Loop.States = append(__StateMachine__000000_UX_Loop.States, __State__000005_UX_n_Update_Stage)
	__StateMachine__000000_UX_Loop.States = append(__StateMachine__000000_UX_Loop.States, __State__000006_Enforce_Model_Semantic)
	__StateMachine__000000_UX_Loop.States = append(__StateMachine__000000_UX_Loop.States, __State__000007_UX_1_Update_UX)
	__StateMachine__000000_UX_Loop.States = append(__StateMachine__000000_UX_Loop.States, __State__000008_UX_2_Update_UX)
	__StateMachine__000000_UX_Loop.States = append(__StateMachine__000000_UX_Loop.States, __State__000009_UX_n_Update_UX)
	__StateMachine__000000_UX_Loop.States = append(__StateMachine__000000_UX_Loop.States, __State__000010_Waiting_for_user_Input_)
	__StateMachine__000000_UX_Loop.States = append(__StateMachine__000000_UX_Loop.States, __State__000011_Load_Stage)
	__StateMachine__000000_UX_Loop.States = append(__StateMachine__000000_UX_Loop.States, __State__000012_Persist_Stage)
	__StateMachine__000000_UX_Loop.States = append(__StateMachine__000000_UX_Loop.States, __State__000013_Stage_Modified_)
	__StateMachine__000000_UX_Loop.States = append(__StateMachine__000000_UX_Loop.States, __State__000014_Update_UX_)
	__StateMachine__000000_UX_Loop.Diagrams = append(__StateMachine__000000_UX_Loop.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__StateMachine__000000_UX_Loop.InitialState = __State__000011_Load_Stage
	// setup of StateShape instances pointers
	__StateShape__000000_UX_1_Waiting_for_User_Input.State = __State__000000_UX_1_Waiting_for_User_Input
	__StateShape__000001_UX_2_Waiting_for_User_Input.State = __State__000001_UX_2_Waiting_for_User_Input
	__StateShape__000002_UX_n_Waiting_for_User_Input.State = __State__000002_UX_n_Waiting_for_User_Input
	__StateShape__000003_UX_1_Update_Stage.State = __State__000003_UX_1_Update_Stage
	__StateShape__000004_UX_2_Update_Stage.State = __State__000004_UX_2_Update_Stage
	__StateShape__000005_UX_n_Update_Stage.State = __State__000005_UX_n_Update_Stage
	__StateShape__000006_Enforce_Model_Semantic.State = __State__000006_Enforce_Model_Semantic
	__StateShape__000007_UX_1_Update_UX.State = __State__000007_UX_1_Update_UX
	__StateShape__000008_UX_2_Update_UX.State = __State__000008_UX_2_Update_UX
	__StateShape__000009_UX_n_Update_UX.State = __State__000009_UX_n_Update_UX
	__StateShape__000010_Waiting_for_user_Input_.State = __State__000010_Waiting_for_user_Input_
	__StateShape__000011_Load_Stage.State = __State__000011_Load_Stage
	__StateShape__000012_Persist_Stage.State = __State__000012_Persist_Stage
	__StateShape__000013_Stage_Modified_.State = __State__000013_Stage_Modified_
	__StateShape__000014_Update_UX_.State = __State__000014_Update_UX_
	// setup of Transition instances pointers
	__Transition__000000__User_Input.Start = __State__000000_UX_1_Waiting_for_User_Input
	__Transition__000000__User_Input.End = __State__000003_UX_1_Update_Stage
	__Transition__000000__User_Input.Diagrams = append(__Transition__000000__User_Input.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__Transition__000001__User_Input.Start = __State__000001_UX_2_Waiting_for_User_Input
	__Transition__000001__User_Input.End = __State__000004_UX_2_Update_Stage
	__Transition__000001__User_Input.Diagrams = append(__Transition__000001__User_Input.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__Transition__000002_User_Input.Start = __State__000002_UX_n_Waiting_for_User_Input
	__Transition__000002_User_Input.End = __State__000005_UX_n_Update_Stage
	__Transition__000002_User_Input.Diagrams = append(__Transition__000002_User_Input.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__Transition__000003_Stage_updated.Start = __State__000003_UX_1_Update_Stage
	__Transition__000003_Stage_updated.End = __State__000006_Enforce_Model_Semantic
	__Transition__000003_Stage_updated.Diagrams = append(__Transition__000003_Stage_updated.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__Transition__000004_Stage_updated.Start = __State__000004_UX_2_Update_Stage
	__Transition__000004_Stage_updated.End = __State__000006_Enforce_Model_Semantic
	__Transition__000004_Stage_updated.Diagrams = append(__Transition__000004_Stage_updated.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__Transition__000005_Stage_updated.Start = __State__000005_UX_n_Update_Stage
	__Transition__000005_Stage_updated.End = __State__000006_Enforce_Model_Semantic
	__Transition__000005_Stage_updated.Diagrams = append(__Transition__000005_Stage_updated.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__Transition__000006__Model_Updated.Start = __State__000006_Enforce_Model_Semantic
	__Transition__000006__Model_Updated.End = __State__000013_Stage_Modified_
	__Transition__000006__Model_Updated.Diagrams = append(__Transition__000006__Model_Updated.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__Transition__000007_UX_1_updated.Start = __State__000007_UX_1_Update_UX
	__Transition__000007_UX_1_updated.End = __State__000008_UX_2_Update_UX
	__Transition__000007_UX_1_updated.Diagrams = append(__Transition__000007_UX_1_updated.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__Transition__000008_UX_2_updated.Start = __State__000008_UX_2_Update_UX
	__Transition__000008_UX_2_updated.End = __State__000009_UX_n_Update_UX
	__Transition__000008_UX_2_updated.Diagrams = append(__Transition__000008_UX_2_updated.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__Transition__000009_UX_n_updated.Start = __State__000009_UX_n_Update_UX
	__Transition__000009_UX_n_updated.End = __State__000010_Waiting_for_user_Input_
	__Transition__000009_UX_n_updated.Diagrams = append(__Transition__000009_UX_n_updated.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__Transition__000010__Stage_loaded.Start = __State__000011_Load_Stage
	__Transition__000010__Stage_loaded.End = __State__000006_Enforce_Model_Semantic
	__Transition__000010__Stage_loaded.Diagrams = append(__Transition__000010__Stage_loaded.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__Transition__000011__Stage_persisted.Start = __State__000012_Persist_Stage
	__Transition__000011__Stage_persisted.End = __State__000014_Update_UX_
	__Transition__000011__Stage_persisted.Diagrams = append(__Transition__000011__Stage_persisted.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__Transition__000012_No.Start = __State__000013_Stage_Modified_
	__Transition__000012_No.End = __State__000010_Waiting_for_user_Input_
	__Transition__000012_No.Diagrams = append(__Transition__000012_No.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__Transition__000013_Yes.Start = __State__000013_Stage_Modified_
	__Transition__000013_Yes.End = __State__000012_Persist_Stage
	__Transition__000013_Yes.Diagrams = append(__Transition__000013_Yes.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__Transition__000014_Yes.Start = __State__000014_Update_UX_
	__Transition__000014_Yes.End = __State__000007_UX_1_Update_UX
	__Transition__000014_Yes.Diagrams = append(__Transition__000014_Yes.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__Transition__000015_No.Start = __State__000014_Update_UX_
	__Transition__000015_No.End = __State__000010_Waiting_for_user_Input_
	__Transition__000015_No.Diagrams = append(__Transition__000015_No.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	// setup of Transition_Shape instances pointers
	__Transition_Shape__000000_UX_1_Waiting_for_User_Input_to_UX_1_Update_Stage.Transition = __Transition__000000__User_Input
	__Transition_Shape__000001_UX_2_Waiting_for_User_Input_to_UX_2_Update_Stage.Transition = __Transition__000001__User_Input
	__Transition_Shape__000002_UX_n_Waiting_for_User_Input_to_UX_n_Update_Stage.Transition = __Transition__000002_User_Input
	__Transition_Shape__000003_UX_1_Update_Stage_to_Enforce_Model_Semantic.Transition = __Transition__000003_Stage_updated
	__Transition_Shape__000004_UX_2_Update_Stage_to_Enforce_Model_Semantic.Transition = __Transition__000004_Stage_updated
	__Transition_Shape__000005_UX_n_Update_Stage_to_Enforce_Model_Semantic.Transition = __Transition__000005_Stage_updated
	__Transition_Shape__000006_Enforce_Model_Semantic_to_UX_1_Update_UX.Transition = __Transition__000006__Model_Updated
	__Transition_Shape__000007_Enforce_Model_Semantic_to_UX_2_Update_UX.Transition = __Transition__000007_UX_1_updated
	__Transition_Shape__000008_Enforce_Model_Semantic_to_UX_n_Update_UX.Transition = __Transition__000008_UX_2_updated
	__Transition_Shape__000009_UX_n_Update_UX_to_Waiting_for_user_Input_.Transition = __Transition__000009_UX_n_updated
	__Transition_Shape__000010_Load_Stage_to_Enforce_Model_Semantic.Transition = __Transition__000010__Stage_loaded
	__Transition_Shape__000011_Persist_Stage_to_UX_1_Update_UX.Transition = __Transition__000011__Stage_persisted
	__Transition_Shape__000012_Stage_Modified_to_UX_1_Update_UX.Transition = __Transition__000012_No
	__Transition_Shape__000013_Stage_Modified_to_Persist_Stage.Transition = __Transition__000013_Yes
	__Transition_Shape__000014_Update_UX_to_UX_1_Update_UX.Transition = __Transition__000014_Yes
	__Transition_Shape__000015_Update_UX_to_Waiting_for_user_Input_.Transition = __Transition__000015_No
}

