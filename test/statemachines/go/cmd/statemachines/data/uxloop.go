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

	__Transition__000000__User_Input := (&models.Transition{}).Stage(stage)
	__Transition__000001__User_Input := (&models.Transition{}).Stage(stage)
	__Transition__000002_User_Input := (&models.Transition{}).Stage(stage)
	__Transition__000003_Stage_updated := (&models.Transition{}).Stage(stage)
	__Transition__000004_Stage_updated := (&models.Transition{}).Stage(stage)
	__Transition__000005_Stage_updated := (&models.Transition{}).Stage(stage)
	__Transition__000006_Model_Updated := (&models.Transition{}).Stage(stage)
	__Transition__000007_UX_updated := (&models.Transition{}).Stage(stage)
	__Transition__000008_UX_updated := (&models.Transition{}).Stage(stage)
	__Transition__000009_UX_updated := (&models.Transition{}).Stage(stage)

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

	// Setup of values

	__Architecture__000000_Gong_UX_loop_Architecture.Name = `Gong UX loop Architecture`
	__Architecture__000000_Gong_UX_loop_Architecture.NbPixPerCharacter = 8.000000

	__Diagram__000000_UX_Loop_Diagram.Name = `UX Loop Diagram`
	__Diagram__000000_UX_Loop_Diagram.IsChecked = true
	__Diagram__000000_UX_Loop_Diagram.IsExpanded = true
	__Diagram__000000_UX_Loop_Diagram.IsEditable_ = true
	__Diagram__000000_UX_Loop_Diagram.IsInRenameMode = false

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

	__StateMachine__000000_UX_Loop.Name = `UX Loop`
	__StateMachine__000000_UX_Loop.IsNodeExpanded = true

	__StateShape__000000_UX_1_Waiting_for_User_Input.Name = `UX 1 - Waiting for User Input`
	__StateShape__000000_UX_1_Waiting_for_User_Input.IsExpanded = false
	__StateShape__000000_UX_1_Waiting_for_User_Input.X = 308.000000
	__StateShape__000000_UX_1_Waiting_for_User_Input.Y = 151.000000
	__StateShape__000000_UX_1_Waiting_for_User_Input.Width = 205.000000
	__StateShape__000000_UX_1_Waiting_for_User_Input.Height = 80.000000

	__StateShape__000001_UX_2_Waiting_for_User_Input.Name = `UX 2 - Waiting for User Input`
	__StateShape__000001_UX_2_Waiting_for_User_Input.IsExpanded = false
	__StateShape__000001_UX_2_Waiting_for_User_Input.X = 534.000000
	__StateShape__000001_UX_2_Waiting_for_User_Input.Y = 149.000000
	__StateShape__000001_UX_2_Waiting_for_User_Input.Width = 200.000000
	__StateShape__000001_UX_2_Waiting_for_User_Input.Height = 80.000000

	__StateShape__000002_UX_n_Waiting_for_User_Input.Name = `UX n - Waiting for User Input`
	__StateShape__000002_UX_n_Waiting_for_User_Input.IsExpanded = false
	__StateShape__000002_UX_n_Waiting_for_User_Input.X = 754.000000
	__StateShape__000002_UX_n_Waiting_for_User_Input.Y = 148.000000
	__StateShape__000002_UX_n_Waiting_for_User_Input.Width = 200.000000
	__StateShape__000002_UX_n_Waiting_for_User_Input.Height = 80.000000

	__StateShape__000003_UX_1_Update_Stage.Name = `UX 1 - Update Stage`
	__StateShape__000003_UX_1_Update_Stage.IsExpanded = false
	__StateShape__000003_UX_1_Update_Stage.X = 312.000000
	__StateShape__000003_UX_1_Update_Stage.Y = 378.999985
	__StateShape__000003_UX_1_Update_Stage.Width = 200.000000
	__StateShape__000003_UX_1_Update_Stage.Height = 80.000000

	__StateShape__000004_UX_2_Update_Stage.Name = `UX 2 - Update Stage`
	__StateShape__000004_UX_2_Update_Stage.IsExpanded = false
	__StateShape__000004_UX_2_Update_Stage.X = 539.000000
	__StateShape__000004_UX_2_Update_Stage.Y = 377.999985
	__StateShape__000004_UX_2_Update_Stage.Width = 200.000000
	__StateShape__000004_UX_2_Update_Stage.Height = 80.000000

	__StateShape__000005_UX_n_Update_Stage.Name = `UX n - Update Stage`
	__StateShape__000005_UX_n_Update_Stage.IsExpanded = false
	__StateShape__000005_UX_n_Update_Stage.X = 760.000000
	__StateShape__000005_UX_n_Update_Stage.Y = 380.999985
	__StateShape__000005_UX_n_Update_Stage.Width = 200.000000
	__StateShape__000005_UX_n_Update_Stage.Height = 78.000000

	__StateShape__000006_Enforce_Model_Semantic.Name = `Enforce Model Semantic`
	__StateShape__000006_Enforce_Model_Semantic.IsExpanded = false
	__StateShape__000006_Enforce_Model_Semantic.X = 316.000000
	__StateShape__000006_Enforce_Model_Semantic.Y = 585.000015
	__StateShape__000006_Enforce_Model_Semantic.Width = 646.000000
	__StateShape__000006_Enforce_Model_Semantic.Height = 80.000000

	__StateShape__000007_UX_1_Update_UX.Name = `UX 1 - Update UX`
	__StateShape__000007_UX_1_Update_UX.IsExpanded = false
	__StateShape__000007_UX_1_Update_UX.X = 315.000000
	__StateShape__000007_UX_1_Update_UX.Y = 754.999954
	__StateShape__000007_UX_1_Update_UX.Width = 200.000000
	__StateShape__000007_UX_1_Update_UX.Height = 80.000000

	__StateShape__000008_UX_2_Update_UX.Name = `UX 2 - Update UX`
	__StateShape__000008_UX_2_Update_UX.IsExpanded = false
	__StateShape__000008_UX_2_Update_UX.X = 546.000000
	__StateShape__000008_UX_2_Update_UX.Y = 751.999954
	__StateShape__000008_UX_2_Update_UX.Width = 200.000000
	__StateShape__000008_UX_2_Update_UX.Height = 80.000000

	__StateShape__000009_UX_n_Update_UX.Name = `UX n - Update UX`
	__StateShape__000009_UX_n_Update_UX.IsExpanded = false
	__StateShape__000009_UX_n_Update_UX.X = 774.000000
	__StateShape__000009_UX_n_Update_UX.Y = 750.999954
	__StateShape__000009_UX_n_Update_UX.Width = 200.000000
	__StateShape__000009_UX_n_Update_UX.Height = 80.000000

	__StateShape__000010_Waiting_for_user_Input_.Name = `Waiting for user Input `
	__StateShape__000010_Waiting_for_user_Input_.IsExpanded = false
	__StateShape__000010_Waiting_for_user_Input_.X = 251.000000
	__StateShape__000010_Waiting_for_user_Input_.Y = 90.000000
	__StateShape__000010_Waiting_for_user_Input_.Width = 745.000000
	__StateShape__000010_Waiting_for_user_Input_.Height = 204.999985

	__Transition__000000__User_Input.Name = `/User Input`

	__Transition__000001__User_Input.Name = `/User Input`

	__Transition__000002_User_Input.Name = `User Input`

	__Transition__000003_Stage_updated.Name = `Stage updated`

	__Transition__000004_Stage_updated.Name = `Stage updated`

	__Transition__000005_Stage_updated.Name = `Stage updated`

	__Transition__000006_Model_Updated.Name = `Model Updated`

	__Transition__000007_UX_updated.Name = `UX updated`

	__Transition__000008_UX_updated.Name = `UX updated`

	__Transition__000009_UX_updated.Name = `UX updated`

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
	__Transition_Shape__000006_Enforce_Model_Semantic_to_UX_1_Update_UX.EndRatio = 0.528932
	__Transition_Shape__000006_Enforce_Model_Semantic_to_UX_1_Update_UX.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000006_Enforce_Model_Semantic_to_UX_1_Update_UX.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000006_Enforce_Model_Semantic_to_UX_1_Update_UX.CornerOffsetRatio = 1.625087

	__Transition_Shape__000007_Enforce_Model_Semantic_to_UX_2_Update_UX.Name = `Enforce Model Semantic to UX 2 - Update UX`
	__Transition_Shape__000007_Enforce_Model_Semantic_to_UX_2_Update_UX.StartRatio = 0.490381
	__Transition_Shape__000007_Enforce_Model_Semantic_to_UX_2_Update_UX.EndRatio = 0.103932
	__Transition_Shape__000007_Enforce_Model_Semantic_to_UX_2_Update_UX.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000007_Enforce_Model_Semantic_to_UX_2_Update_UX.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000007_Enforce_Model_Semantic_to_UX_2_Update_UX.CornerOffsetRatio = 1.625087

	__Transition_Shape__000008_Enforce_Model_Semantic_to_UX_n_Update_UX.Name = `Enforce Model Semantic to UX n - Update UX`
	__Transition_Shape__000008_Enforce_Model_Semantic_to_UX_n_Update_UX.StartRatio = 0.830939
	__Transition_Shape__000008_Enforce_Model_Semantic_to_UX_n_Update_UX.EndRatio = 0.398932
	__Transition_Shape__000008_Enforce_Model_Semantic_to_UX_n_Update_UX.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000008_Enforce_Model_Semantic_to_UX_n_Update_UX.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000008_Enforce_Model_Semantic_to_UX_n_Update_UX.CornerOffsetRatio = 1.662587

	__Transition_Shape__000009_UX_n_Update_UX_to_Waiting_for_user_Input_.Name = `UX n - Update UX to Waiting for user Input `
	__Transition_Shape__000009_UX_n_Update_UX_to_Waiting_for_user_Input_.StartRatio = 0.500000
	__Transition_Shape__000009_UX_n_Update_UX_to_Waiting_for_user_Input_.EndRatio = 0.482961
	__Transition_Shape__000009_UX_n_Update_UX_to_Waiting_for_user_Input_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__000009_UX_n_Update_UX_to_Waiting_for_user_Input_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__000009_UX_n_Update_UX_to_Waiting_for_user_Input_.CornerOffsetRatio = 1.553932

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
	// setup of State instances pointers
	__State__000000_UX_1_Waiting_for_User_Input.Diagrams = append(__State__000000_UX_1_Waiting_for_User_Input.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000001_UX_2_Waiting_for_User_Input.Diagrams = append(__State__000001_UX_2_Waiting_for_User_Input.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000002_UX_n_Waiting_for_User_Input.Diagrams = append(__State__000002_UX_n_Waiting_for_User_Input.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000003_UX_1_Update_Stage.Diagrams = append(__State__000003_UX_1_Update_Stage.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000004_UX_2_Update_Stage.Diagrams = append(__State__000004_UX_2_Update_Stage.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000005_UX_n_Update_Stage.Diagrams = append(__State__000005_UX_n_Update_Stage.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000006_Enforce_Model_Semantic.Diagrams = append(__State__000006_Enforce_Model_Semantic.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000007_UX_1_Update_UX.Diagrams = append(__State__000007_UX_1_Update_UX.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000008_UX_2_Update_UX.Diagrams = append(__State__000008_UX_2_Update_UX.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000009_UX_n_Update_UX.Diagrams = append(__State__000009_UX_n_Update_UX.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__State__000010_Waiting_for_user_Input_.Diagrams = append(__State__000010_Waiting_for_user_Input_.Diagrams, __Diagram__000000_UX_Loop_Diagram)
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
	__StateMachine__000000_UX_Loop.Diagrams = append(__StateMachine__000000_UX_Loop.Diagrams, __Diagram__000000_UX_Loop_Diagram)
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
	__Transition__000006_Model_Updated.Start = __State__000006_Enforce_Model_Semantic
	__Transition__000006_Model_Updated.End = __State__000007_UX_1_Update_UX
	__Transition__000006_Model_Updated.Diagrams = append(__Transition__000006_Model_Updated.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__Transition__000007_UX_updated.Start = __State__000007_UX_1_Update_UX
	__Transition__000007_UX_updated.End = __State__000008_UX_2_Update_UX
	__Transition__000007_UX_updated.Diagrams = append(__Transition__000007_UX_updated.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__Transition__000008_UX_updated.Start = __State__000008_UX_2_Update_UX
	__Transition__000008_UX_updated.End = __State__000009_UX_n_Update_UX
	__Transition__000008_UX_updated.Diagrams = append(__Transition__000008_UX_updated.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	__Transition__000009_UX_updated.Start = __State__000009_UX_n_Update_UX
	__Transition__000009_UX_updated.End = __State__000010_Waiting_for_user_Input_
	__Transition__000009_UX_updated.Diagrams = append(__Transition__000009_UX_updated.Diagrams, __Diagram__000000_UX_Loop_Diagram)
	// setup of Transition_Shape instances pointers
	__Transition_Shape__000000_UX_1_Waiting_for_User_Input_to_UX_1_Update_Stage.Transition = __Transition__000000__User_Input
	__Transition_Shape__000001_UX_2_Waiting_for_User_Input_to_UX_2_Update_Stage.Transition = __Transition__000001__User_Input
	__Transition_Shape__000002_UX_n_Waiting_for_User_Input_to_UX_n_Update_Stage.Transition = __Transition__000002_User_Input
	__Transition_Shape__000003_UX_1_Update_Stage_to_Enforce_Model_Semantic.Transition = __Transition__000003_Stage_updated
	__Transition_Shape__000004_UX_2_Update_Stage_to_Enforce_Model_Semantic.Transition = __Transition__000004_Stage_updated
	__Transition_Shape__000005_UX_n_Update_Stage_to_Enforce_Model_Semantic.Transition = __Transition__000005_Stage_updated
	__Transition_Shape__000006_Enforce_Model_Semantic_to_UX_1_Update_UX.Transition = __Transition__000006_Model_Updated
	__Transition_Shape__000007_Enforce_Model_Semantic_to_UX_2_Update_UX.Transition = __Transition__000007_UX_updated
	__Transition_Shape__000008_Enforce_Model_Semantic_to_UX_n_Update_UX.Transition = __Transition__000008_UX_updated
	__Transition_Shape__000009_UX_n_Update_UX_to_Waiting_for_user_Input_.Transition = __Transition__000009_UX_updated
}

