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

	const __write__local_time = "2025-11-25 08:38:29.486224 CET"
	const __write__utc_time__ = "2025-11-25 07:38:29.486224 UTC"

	const __commitId__ = "0000001144"

	// Declaration of instances to stage

	__Architecture__000000_Traffic_Lights_of_the_world := (&models.Architecture{}).Stage(stage)

	__Diagram__000000_Traffic_Light_FR_Diagram := (&models.Diagram{}).Stage(stage)
	__Diagram__000001_Traffic_Light_UK_Diagram := (&models.Diagram{}).Stage(stage)

	__Message__000000_08_38_29_Off_UK_to_Red_UK_Repair_Report := (&models.Message{}).Stage(stage)

	__MessageType__000000_Repair_Report := (&models.MessageType{}).Stage(stage)

	__Object__000000_02_MI_DOF_2025_11_26_DEP_083652 := (&models.Object{}).Stage(stage)

	__Role__000000_Technician := (&models.Role{}).Stage(stage)
	__Role__000001_Timer := (&models.Role{}).Stage(stage)

	__State__000000_Initial_State_FR := (&models.State{}).Stage(stage)
	__State__000001_Red_FR := (&models.State{}).Stage(stage)
	__State__000002_Green_FR := (&models.State{}).Stage(stage)
	__State__000003_Yellow_FR := (&models.State{}).Stage(stage)
	__State__000004_On_FR := (&models.State{}).Stage(stage)
	__State__000005_Off_FR := (&models.State{}).Stage(stage)
	__State__000006_Initial_State_UK := (&models.State{}).Stage(stage)
	__State__000007_Red_UK := (&models.State{}).Stage(stage)
	__State__000008_Yellow_UK := (&models.State{}).Stage(stage)
	__State__000009_Green_UK := (&models.State{}).Stage(stage)
	__State__000010_On_UK := (&models.State{}).Stage(stage)
	__State__000011_Off_UK := (&models.State{}).Stage(stage)

	__StateMachine__000000_Traffic_Lights_FR := (&models.StateMachine{}).Stage(stage)
	__StateMachine__000001_Traffic_Light_UK := (&models.StateMachine{}).Stage(stage)

	__StateShape__000000_Red_FR := (&models.StateShape{}).Stage(stage)
	__StateShape__000001_Initial_State_FR := (&models.StateShape{}).Stage(stage)
	__StateShape__000002_Green_FR := (&models.StateShape{}).Stage(stage)
	__StateShape__000003_Yellow_FR := (&models.StateShape{}).Stage(stage)
	__StateShape__000004_On_FR := (&models.StateShape{}).Stage(stage)
	__StateShape__000005_Off_FR := (&models.StateShape{}).Stage(stage)
	__StateShape__000006_Red_UK := (&models.StateShape{}).Stage(stage)
	__StateShape__000007_Initial_State_UK := (&models.StateShape{}).Stage(stage)
	__StateShape__000008_Yellow_UK := (&models.StateShape{}).Stage(stage)
	__StateShape__000009_Green_UK := (&models.StateShape{}).Stage(stage)
	__StateShape__000010_On_UK := (&models.StateShape{}).Stage(stage)
	__StateShape__000011_Off_UK := (&models.StateShape{}).Stage(stage)

	__Transition__000000_Initial_State_FR_to_Red_FR := (&models.Transition{}).Stage(stage)
	__Transition__000001_Red_FR_to_Green_FR := (&models.Transition{}).Stage(stage)
	__Transition__000002_Green_FR_to_Yellow_FR := (&models.Transition{}).Stage(stage)
	__Transition__000003_Yellow_FR_to_Red_FR := (&models.Transition{}).Stage(stage)
	__Transition__000004_On_to_Off := (&models.Transition{}).Stage(stage)
	__Transition__000005_Off_to_On := (&models.Transition{}).Stage(stage)
	__Transition__000006_Initial_State_UK_to_New_State := (&models.Transition{}).Stage(stage)
	__Transition__000007_Red_UK_to_Yellow_UK := (&models.Transition{}).Stage(stage)
	__Transition__000008_Yellow_UK_to_Green_UK := (&models.Transition{}).Stage(stage)
	__Transition__000009_On_UK_to_Off_UK := (&models.Transition{}).Stage(stage)
	__Transition__000010_Off_UK_to_Red_UK := (&models.Transition{}).Stage(stage)
	__Transition__000011_Green_UK_to_Red_UK := (&models.Transition{}).Stage(stage)

	__Transition_Shape__000000_Initial_State_FR_to_Red_FR := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000001_Red_FR_to_Green_FR := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000002_Green_FR_to_Yellow_FR := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000003_Yellow_FR_to_Red_FR := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000004_New_State_to_New_State := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000005_Off_to_On := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000006_Initial_State_UK_to_New_State := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000007_Red_UK_to_Yellow_UK := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000008_Yellow_UK_to_Green_UK := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000009_On_UK_to_Off_UK := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000010_Off_UK_to_On_UK := (&models.Transition_Shape{}).Stage(stage)
	__Transition_Shape__000011_Green_UK_to_Red_UK := (&models.Transition_Shape{}).Stage(stage)

	// Setup of values

	__Architecture__000000_Traffic_Lights_of_the_world.Name = `Traffic Lights of the world`
	__Architecture__000000_Traffic_Lights_of_the_world.NbPixPerCharacter = 8.000000

	__Diagram__000000_Traffic_Light_FR_Diagram.Name = `Traffic Light FR Diagram`
	__Diagram__000000_Traffic_Light_FR_Diagram.IsChecked = false
	__Diagram__000000_Traffic_Light_FR_Diagram.IsExpanded = true
	__Diagram__000000_Traffic_Light_FR_Diagram.IsEditable_ = true
	__Diagram__000000_Traffic_Light_FR_Diagram.IsInRenameMode = false

	__Diagram__000001_Traffic_Light_UK_Diagram.Name = `Traffic Light UK Diagram`
	__Diagram__000001_Traffic_Light_UK_Diagram.IsChecked = true
	__Diagram__000001_Traffic_Light_UK_Diagram.IsExpanded = true
	__Diagram__000001_Traffic_Light_UK_Diagram.IsEditable_ = true
	__Diagram__000001_Traffic_Light_UK_Diagram.IsInRenameMode = false

	__Message__000000_08_38_29_Off_UK_to_Red_UK_Repair_Report.Name = `08:38:29 Off UK to Red UK -> Repair Report`
	__Message__000000_08_38_29_Off_UK_to_Red_UK_Repair_Report.IsSelected = true

	__MessageType__000000_Repair_Report.Name = `Repair Report`
	__MessageType__000000_Repair_Report.Description = ``

	__Object__000000_02_MI_DOF_2025_11_26_DEP_083652.Name = `02/MI DOF/ 2025-11-26 DEP/ 083652`
	__Object__000000_02_MI_DOF_2025_11_26_DEP_083652.IsSelected = true
	__Object__000000_02_MI_DOF_2025_11_26_DEP_083652.Rank = 0
	__Object__000000_02_MI_DOF_2025_11_26_DEP_083652.DOF, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2025-11-26 08:36:52.163215 +0100 CET m=+86453.154167293")

	__Role__000000_Technician.Name = `Technician`
	__Role__000000_Technician.Acronym = `Tech`

	__Role__000001_Timer.Name = `Timer`
	__Role__000001_Timer.Acronym = `Timer`

	__State__000000_Initial_State_FR.Name = `Initial State FR`
	__State__000000_Initial_State_FR.IsDecisionNode = false
	__State__000000_Initial_State_FR.IsFictif = false
	__State__000000_Initial_State_FR.IsEndState = false

	__State__000001_Red_FR.Name = `Red FR`
	__State__000001_Red_FR.IsDecisionNode = false
	__State__000001_Red_FR.IsFictif = false
	__State__000001_Red_FR.IsEndState = false

	__State__000002_Green_FR.Name = `Green FR`
	__State__000002_Green_FR.IsDecisionNode = false
	__State__000002_Green_FR.IsFictif = false
	__State__000002_Green_FR.IsEndState = false

	__State__000003_Yellow_FR.Name = `Yellow FR`
	__State__000003_Yellow_FR.IsDecisionNode = false
	__State__000003_Yellow_FR.IsFictif = false
	__State__000003_Yellow_FR.IsEndState = false

	__State__000004_On_FR.Name = `On FR`
	__State__000004_On_FR.IsDecisionNode = false
	__State__000004_On_FR.IsFictif = false
	__State__000004_On_FR.IsEndState = false

	__State__000005_Off_FR.Name = `Off FR`
	__State__000005_Off_FR.IsDecisionNode = false
	__State__000005_Off_FR.IsFictif = false
	__State__000005_Off_FR.IsEndState = true

	__State__000006_Initial_State_UK.Name = `Initial State UK`
	__State__000006_Initial_State_UK.IsDecisionNode = false
	__State__000006_Initial_State_UK.IsFictif = false
	__State__000006_Initial_State_UK.IsEndState = false

	__State__000007_Red_UK.Name = `Red UK`
	__State__000007_Red_UK.IsDecisionNode = false
	__State__000007_Red_UK.IsFictif = false
	__State__000007_Red_UK.IsEndState = false

	__State__000008_Yellow_UK.Name = `Yellow UK`
	__State__000008_Yellow_UK.IsDecisionNode = false
	__State__000008_Yellow_UK.IsFictif = false
	__State__000008_Yellow_UK.IsEndState = false

	__State__000009_Green_UK.Name = `Green UK`
	__State__000009_Green_UK.IsDecisionNode = false
	__State__000009_Green_UK.IsFictif = false
	__State__000009_Green_UK.IsEndState = false

	__State__000010_On_UK.Name = `On UK`
	__State__000010_On_UK.IsDecisionNode = false
	__State__000010_On_UK.IsFictif = false
	__State__000010_On_UK.IsEndState = false

	__State__000011_Off_UK.Name = `Off UK`
	__State__000011_Off_UK.IsDecisionNode = false
	__State__000011_Off_UK.IsFictif = false
	__State__000011_Off_UK.IsEndState = true

	__StateMachine__000000_Traffic_Lights_FR.Name = `Traffic Lights FR`
	__StateMachine__000000_Traffic_Lights_FR.IsNodeExpanded = true

	__StateMachine__000001_Traffic_Light_UK.Name = `Traffic Light UK`
	__StateMachine__000001_Traffic_Light_UK.IsNodeExpanded = true

	__StateShape__000000_Red_FR.Name = `Red FR`
	__StateShape__000000_Red_FR.IsExpanded = false
	__StateShape__000000_Red_FR.X = 568.000000
	__StateShape__000000_Red_FR.Y = 209.000000
	__StateShape__000000_Red_FR.Width = 200.000000
	__StateShape__000000_Red_FR.Height = 80.000000

	__StateShape__000001_Initial_State_FR.Name = `Initial State FR`
	__StateShape__000001_Initial_State_FR.IsExpanded = false
	__StateShape__000001_Initial_State_FR.X = 297.000000
	__StateShape__000001_Initial_State_FR.Y = 55.000000
	__StateShape__000001_Initial_State_FR.Width = 148.000000
	__StateShape__000001_Initial_State_FR.Height = 20.000000

	__StateShape__000002_Green_FR.Name = `Green FR`
	__StateShape__000002_Green_FR.IsExpanded = false
	__StateShape__000002_Green_FR.X = 418.000000
	__StateShape__000002_Green_FR.Y = 358.000000
	__StateShape__000002_Green_FR.Width = 200.000000
	__StateShape__000002_Green_FR.Height = 80.000000

	__StateShape__000003_Yellow_FR.Name = `Yellow FR`
	__StateShape__000003_Yellow_FR.IsExpanded = false
	__StateShape__000003_Yellow_FR.X = 566.000000
	__StateShape__000003_Yellow_FR.Y = 524.000000
	__StateShape__000003_Yellow_FR.Width = 200.000000
	__StateShape__000003_Yellow_FR.Height = 80.000000

	__StateShape__000004_On_FR.Name = `On FR`
	__StateShape__000004_On_FR.IsExpanded = false
	__StateShape__000004_On_FR.X = 351.000000
	__StateShape__000004_On_FR.Y = 156.000000
	__StateShape__000004_On_FR.Width = 678.000000
	__StateShape__000004_On_FR.Height = 478.000000

	__StateShape__000005_Off_FR.Name = `Off FR`
	__StateShape__000005_Off_FR.IsExpanded = false
	__StateShape__000005_Off_FR.X = 1120.000000
	__StateShape__000005_Off_FR.Y = 379.875000
	__StateShape__000005_Off_FR.Width = 91.000000
	__StateShape__000005_Off_FR.Height = 36.000000

	__StateShape__000006_Red_UK.Name = `Red UK`
	__StateShape__000006_Red_UK.IsExpanded = false
	__StateShape__000006_Red_UK.X = 346.000000
	__StateShape__000006_Red_UK.Y = 253.000000
	__StateShape__000006_Red_UK.Width = 200.000000
	__StateShape__000006_Red_UK.Height = 80.000000

	__StateShape__000007_Initial_State_UK.Name = `Initial State UK`
	__StateShape__000007_Initial_State_UK.IsExpanded = false
	__StateShape__000007_Initial_State_UK.X = 324.000000
	__StateShape__000007_Initial_State_UK.Y = 91.000000
	__StateShape__000007_Initial_State_UK.Width = 121.000000
	__StateShape__000007_Initial_State_UK.Height = 20.000000

	__StateShape__000008_Yellow_UK.Name = `Yellow UK`
	__StateShape__000008_Yellow_UK.IsExpanded = false
	__StateShape__000008_Yellow_UK.X = 342.000000
	__StateShape__000008_Yellow_UK.Y = 400.000000
	__StateShape__000008_Yellow_UK.Width = 200.000000
	__StateShape__000008_Yellow_UK.Height = 80.000000

	__StateShape__000009_Green_UK.Name = `Green UK`
	__StateShape__000009_Green_UK.IsExpanded = false
	__StateShape__000009_Green_UK.X = 344.000000
	__StateShape__000009_Green_UK.Y = 576.000000
	__StateShape__000009_Green_UK.Width = 200.000000
	__StateShape__000009_Green_UK.Height = 80.000000

	__StateShape__000010_On_UK.Name = `On UK`
	__StateShape__000010_On_UK.IsExpanded = false
	__StateShape__000010_On_UK.X = 197.000000
	__StateShape__000010_On_UK.Y = 209.000000
	__StateShape__000010_On_UK.Width = 742.000000
	__StateShape__000010_On_UK.Height = 479.000000

	__StateShape__000011_Off_UK.Name = `Off UK`
	__StateShape__000011_Off_UK.IsExpanded = false
	__StateShape__000011_Off_UK.X = 1133.000000
	__StateShape__000011_Off_UK.Y = 373.000000
	__StateShape__000011_Off_UK.Width = 94.000000
	__StateShape__000011_Off_UK.Height = 36.000000

	__Transition__000000_Initial_State_FR_to_Red_FR.Name = `Initial State FR to Red FR`

	__Transition__000001_Red_FR_to_Green_FR.Name = `Red FR to Green FR`

	__Transition__000002_Green_FR_to_Yellow_FR.Name = `Green FR to Yellow FR`

	__Transition__000003_Yellow_FR_to_Red_FR.Name = `Yellow FR to Red FR`

	__Transition__000004_On_to_Off.Name = `On to Off`

	__Transition__000005_Off_to_On.Name = `Off to On`

	__Transition__000006_Initial_State_UK_to_New_State.Name = `Initial State UK to New State`

	__Transition__000007_Red_UK_to_Yellow_UK.Name = `Red UK to Yellow UK`

	__Transition__000008_Yellow_UK_to_Green_UK.Name = `Yellow UK to Green UK`

	__Transition__000009_On_UK_to_Off_UK.Name = `On UK to Off UK`

	__Transition__000010_Off_UK_to_Red_UK.Name = `Off UK to Red UK`

	__Transition__000011_Green_UK_to_Red_UK.Name = `Green UK to Red UK`

	__Transition_Shape__000000_Initial_State_FR_to_Red_FR.Name = `Initial State FR to Red FR`
	__Transition_Shape__000000_Initial_State_FR_to_Red_FR.StartRatio = 1.000000
	__Transition_Shape__000000_Initial_State_FR_to_Red_FR.EndRatio = 0.100073
	__Transition_Shape__000000_Initial_State_FR_to_Red_FR.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000000_Initial_State_FR_to_Red_FR.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__000000_Initial_State_FR_to_Red_FR.CornerOffsetRatio = 1.100293

	__Transition_Shape__000001_Red_FR_to_Green_FR.Name = `Red FR to Green FR`
	__Transition_Shape__000001_Red_FR_to_Green_FR.StartRatio = 0.598789
	__Transition_Shape__000001_Red_FR_to_Green_FR.EndRatio = 0.608789
	__Transition_Shape__000001_Red_FR_to_Green_FR.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000001_Red_FR_to_Green_FR.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000001_Red_FR_to_Green_FR.CornerOffsetRatio = 1.425073

	__Transition_Shape__000002_Green_FR_to_Yellow_FR.Name = `Green FR to Yellow FR`
	__Transition_Shape__000002_Green_FR_to_Yellow_FR.StartRatio = 0.617656
	__Transition_Shape__000002_Green_FR_to_Yellow_FR.EndRatio = 0.623789
	__Transition_Shape__000002_Green_FR_to_Yellow_FR.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000002_Green_FR_to_Yellow_FR.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000002_Green_FR_to_Yellow_FR.CornerOffsetRatio = 1.687573

	__Transition_Shape__000003_Yellow_FR_to_Red_FR.Name = `Yellow FR to Red FR`
	__Transition_Shape__000003_Yellow_FR_to_Red_FR.StartRatio = 0.500000
	__Transition_Shape__000003_Yellow_FR_to_Red_FR.EndRatio = 0.700073
	__Transition_Shape__000003_Yellow_FR_to_Red_FR.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__000003_Yellow_FR_to_Red_FR.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__000003_Yellow_FR_to_Red_FR.CornerOffsetRatio = 2.103789

	__Transition_Shape__000004_New_State_to_New_State.Name = `New State to New State`
	__Transition_Shape__000004_New_State_to_New_State.StartRatio = 0.884949
	__Transition_Shape__000004_New_State_to_New_State.EndRatio = 0.503635
	__Transition_Shape__000004_New_State_to_New_State.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__000004_New_State_to_New_State.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__000004_New_State_to_New_State.CornerOffsetRatio = 1.436221

	__Transition_Shape__000005_Off_to_On.Name = `Off to On`
	__Transition_Shape__000005_Off_to_On.StartRatio = 0.799536
	__Transition_Shape__000005_Off_to_On.EndRatio = 0.225073
	__Transition_Shape__000005_Off_to_On.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000005_Off_to_On.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__000005_Off_to_On.CornerOffsetRatio = -8.888889

	__Transition_Shape__000006_Initial_State_UK_to_New_State.Name = `Initial State UK to New State`
	__Transition_Shape__000006_Initial_State_UK_to_New_State.StartRatio = 1.000000
	__Transition_Shape__000006_Initial_State_UK_to_New_State.EndRatio = 0.523789
	__Transition_Shape__000006_Initial_State_UK_to_New_State.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000006_Initial_State_UK_to_New_State.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000006_Initial_State_UK_to_New_State.CornerOffsetRatio = 4.350293

	__Transition_Shape__000007_Red_UK_to_Yellow_UK.Name = `Red UK to Yellow UK`
	__Transition_Shape__000007_Red_UK_to_Yellow_UK.StartRatio = 0.458789
	__Transition_Shape__000007_Red_UK_to_Yellow_UK.EndRatio = 0.478789
	__Transition_Shape__000007_Red_UK_to_Yellow_UK.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000007_Red_UK_to_Yellow_UK.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000007_Red_UK_to_Yellow_UK.CornerOffsetRatio = 1.625073

	__Transition_Shape__000008_Yellow_UK_to_Green_UK.Name = `Yellow UK to Green UK`
	__Transition_Shape__000008_Yellow_UK_to_Green_UK.StartRatio = 0.478789
	__Transition_Shape__000008_Yellow_UK_to_Green_UK.EndRatio = 0.363789
	__Transition_Shape__000008_Yellow_UK_to_Green_UK.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000008_Yellow_UK_to_Green_UK.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000008_Yellow_UK_to_Green_UK.CornerOffsetRatio = 1.537573

	__Transition_Shape__000009_On_UK_to_Off_UK.Name = `On UK to Off UK`
	__Transition_Shape__000009_On_UK_to_Off_UK.StartRatio = 0.722350
	__Transition_Shape__000009_On_UK_to_Off_UK.EndRatio = 0.816572
	__Transition_Shape__000009_On_UK_to_Off_UK.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__000009_On_UK_to_Off_UK.EndOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000009_On_UK_to_Off_UK.CornerOffsetRatio = 0.488530

	__Transition_Shape__000010_Off_UK_to_On_UK.Name = `Off UK to On UK`
	__Transition_Shape__000010_Off_UK_to_On_UK.StartRatio = 0.805934
	__Transition_Shape__000010_Off_UK_to_On_UK.EndRatio = 0.187573
	__Transition_Shape__000010_Off_UK_to_On_UK.StartOrientation = models.ORIENTATION_VERTICAL
	__Transition_Shape__000010_Off_UK_to_On_UK.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__000010_Off_UK_to_On_UK.CornerOffsetRatio = 1.000163

	__Transition_Shape__000011_Green_UK_to_Red_UK.Name = `Green UK to Red UK`
	__Transition_Shape__000011_Green_UK_to_Red_UK.StartRatio = 0.500000
	__Transition_Shape__000011_Green_UK_to_Red_UK.EndRatio = 0.675073
	__Transition_Shape__000011_Green_UK_to_Red_UK.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__000011_Green_UK_to_Red_UK.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Transition_Shape__000011_Green_UK_to_Red_UK.CornerOffsetRatio = 2.403789

	// Setup of pointers
	// setup of Architecture instances pointers
	__Architecture__000000_Traffic_Lights_of_the_world.StateMachines = append(__Architecture__000000_Traffic_Lights_of_the_world.StateMachines, __StateMachine__000000_Traffic_Lights_FR)
	__Architecture__000000_Traffic_Lights_of_the_world.StateMachines = append(__Architecture__000000_Traffic_Lights_of_the_world.StateMachines, __StateMachine__000001_Traffic_Light_UK)
	__Architecture__000000_Traffic_Lights_of_the_world.Roles = append(__Architecture__000000_Traffic_Lights_of_the_world.Roles, __Role__000000_Technician)
	__Architecture__000000_Traffic_Lights_of_the_world.Roles = append(__Architecture__000000_Traffic_Lights_of_the_world.Roles, __Role__000001_Timer)
	// setup of Diagram instances pointers
	__Diagram__000000_Traffic_Light_FR_Diagram.State_Shapes = append(__Diagram__000000_Traffic_Light_FR_Diagram.State_Shapes, __StateShape__000004_On_FR)
	__Diagram__000000_Traffic_Light_FR_Diagram.State_Shapes = append(__Diagram__000000_Traffic_Light_FR_Diagram.State_Shapes, __StateShape__000000_Red_FR)
	__Diagram__000000_Traffic_Light_FR_Diagram.State_Shapes = append(__Diagram__000000_Traffic_Light_FR_Diagram.State_Shapes, __StateShape__000001_Initial_State_FR)
	__Diagram__000000_Traffic_Light_FR_Diagram.State_Shapes = append(__Diagram__000000_Traffic_Light_FR_Diagram.State_Shapes, __StateShape__000002_Green_FR)
	__Diagram__000000_Traffic_Light_FR_Diagram.State_Shapes = append(__Diagram__000000_Traffic_Light_FR_Diagram.State_Shapes, __StateShape__000003_Yellow_FR)
	__Diagram__000000_Traffic_Light_FR_Diagram.State_Shapes = append(__Diagram__000000_Traffic_Light_FR_Diagram.State_Shapes, __StateShape__000005_Off_FR)
	__Diagram__000000_Traffic_Light_FR_Diagram.Transition_Shapes = append(__Diagram__000000_Traffic_Light_FR_Diagram.Transition_Shapes, __Transition_Shape__000000_Initial_State_FR_to_Red_FR)
	__Diagram__000000_Traffic_Light_FR_Diagram.Transition_Shapes = append(__Diagram__000000_Traffic_Light_FR_Diagram.Transition_Shapes, __Transition_Shape__000001_Red_FR_to_Green_FR)
	__Diagram__000000_Traffic_Light_FR_Diagram.Transition_Shapes = append(__Diagram__000000_Traffic_Light_FR_Diagram.Transition_Shapes, __Transition_Shape__000002_Green_FR_to_Yellow_FR)
	__Diagram__000000_Traffic_Light_FR_Diagram.Transition_Shapes = append(__Diagram__000000_Traffic_Light_FR_Diagram.Transition_Shapes, __Transition_Shape__000003_Yellow_FR_to_Red_FR)
	__Diagram__000000_Traffic_Light_FR_Diagram.Transition_Shapes = append(__Diagram__000000_Traffic_Light_FR_Diagram.Transition_Shapes, __Transition_Shape__000004_New_State_to_New_State)
	__Diagram__000000_Traffic_Light_FR_Diagram.Transition_Shapes = append(__Diagram__000000_Traffic_Light_FR_Diagram.Transition_Shapes, __Transition_Shape__000005_Off_to_On)
	__Diagram__000001_Traffic_Light_UK_Diagram.State_Shapes = append(__Diagram__000001_Traffic_Light_UK_Diagram.State_Shapes, __StateShape__000010_On_UK)
	__Diagram__000001_Traffic_Light_UK_Diagram.State_Shapes = append(__Diagram__000001_Traffic_Light_UK_Diagram.State_Shapes, __StateShape__000006_Red_UK)
	__Diagram__000001_Traffic_Light_UK_Diagram.State_Shapes = append(__Diagram__000001_Traffic_Light_UK_Diagram.State_Shapes, __StateShape__000007_Initial_State_UK)
	__Diagram__000001_Traffic_Light_UK_Diagram.State_Shapes = append(__Diagram__000001_Traffic_Light_UK_Diagram.State_Shapes, __StateShape__000008_Yellow_UK)
	__Diagram__000001_Traffic_Light_UK_Diagram.State_Shapes = append(__Diagram__000001_Traffic_Light_UK_Diagram.State_Shapes, __StateShape__000009_Green_UK)
	__Diagram__000001_Traffic_Light_UK_Diagram.State_Shapes = append(__Diagram__000001_Traffic_Light_UK_Diagram.State_Shapes, __StateShape__000011_Off_UK)
	__Diagram__000001_Traffic_Light_UK_Diagram.Transition_Shapes = append(__Diagram__000001_Traffic_Light_UK_Diagram.Transition_Shapes, __Transition_Shape__000006_Initial_State_UK_to_New_State)
	__Diagram__000001_Traffic_Light_UK_Diagram.Transition_Shapes = append(__Diagram__000001_Traffic_Light_UK_Diagram.Transition_Shapes, __Transition_Shape__000007_Red_UK_to_Yellow_UK)
	__Diagram__000001_Traffic_Light_UK_Diagram.Transition_Shapes = append(__Diagram__000001_Traffic_Light_UK_Diagram.Transition_Shapes, __Transition_Shape__000008_Yellow_UK_to_Green_UK)
	__Diagram__000001_Traffic_Light_UK_Diagram.Transition_Shapes = append(__Diagram__000001_Traffic_Light_UK_Diagram.Transition_Shapes, __Transition_Shape__000009_On_UK_to_Off_UK)
	__Diagram__000001_Traffic_Light_UK_Diagram.Transition_Shapes = append(__Diagram__000001_Traffic_Light_UK_Diagram.Transition_Shapes, __Transition_Shape__000010_Off_UK_to_On_UK)
	__Diagram__000001_Traffic_Light_UK_Diagram.Transition_Shapes = append(__Diagram__000001_Traffic_Light_UK_Diagram.Transition_Shapes, __Transition_Shape__000011_Green_UK_to_Red_UK)
	// setup of Message instances pointers
	__Message__000000_08_38_29_Off_UK_to_Red_UK_Repair_Report.MessageType = __MessageType__000000_Repair_Report
	__Message__000000_08_38_29_Off_UK_to_Red_UK_Repair_Report.OriginTransition = __Transition__000010_Off_UK_to_Red_UK
	// setup of MessageType instances pointers
	// setup of Object instances pointers
	__Object__000000_02_MI_DOF_2025_11_26_DEP_083652.State = __State__000007_Red_UK
	__Object__000000_02_MI_DOF_2025_11_26_DEP_083652.Messages = append(__Object__000000_02_MI_DOF_2025_11_26_DEP_083652.Messages, __Message__000000_08_38_29_Off_UK_to_Red_UK_Repair_Report)
	// setup of Role instances pointers
	// setup of State instances pointers
	__State__000000_Initial_State_FR.Diagrams = append(__State__000000_Initial_State_FR.Diagrams, __Diagram__000000_Traffic_Light_FR_Diagram)
	__State__000001_Red_FR.Diagrams = append(__State__000001_Red_FR.Diagrams, __Diagram__000000_Traffic_Light_FR_Diagram)
	__State__000002_Green_FR.Diagrams = append(__State__000002_Green_FR.Diagrams, __Diagram__000000_Traffic_Light_FR_Diagram)
	__State__000003_Yellow_FR.Diagrams = append(__State__000003_Yellow_FR.Diagrams, __Diagram__000000_Traffic_Light_FR_Diagram)
	__State__000004_On_FR.SubStates = append(__State__000004_On_FR.SubStates, __State__000001_Red_FR)
	__State__000004_On_FR.SubStates = append(__State__000004_On_FR.SubStates, __State__000002_Green_FR)
	__State__000004_On_FR.SubStates = append(__State__000004_On_FR.SubStates, __State__000003_Yellow_FR)
	__State__000004_On_FR.Diagrams = append(__State__000004_On_FR.Diagrams, __Diagram__000000_Traffic_Light_FR_Diagram)
	__State__000005_Off_FR.Diagrams = append(__State__000005_Off_FR.Diagrams, __Diagram__000000_Traffic_Light_FR_Diagram)
	__State__000006_Initial_State_UK.Diagrams = append(__State__000006_Initial_State_UK.Diagrams, __Diagram__000001_Traffic_Light_UK_Diagram)
	__State__000007_Red_UK.Diagrams = append(__State__000007_Red_UK.Diagrams, __Diagram__000001_Traffic_Light_UK_Diagram)
	__State__000008_Yellow_UK.Diagrams = append(__State__000008_Yellow_UK.Diagrams, __Diagram__000001_Traffic_Light_UK_Diagram)
	__State__000009_Green_UK.Diagrams = append(__State__000009_Green_UK.Diagrams, __Diagram__000001_Traffic_Light_UK_Diagram)
	__State__000010_On_UK.SubStates = append(__State__000010_On_UK.SubStates, __State__000008_Yellow_UK)
	__State__000010_On_UK.SubStates = append(__State__000010_On_UK.SubStates, __State__000009_Green_UK)
	__State__000010_On_UK.SubStates = append(__State__000010_On_UK.SubStates, __State__000007_Red_UK)
	__State__000010_On_UK.Diagrams = append(__State__000010_On_UK.Diagrams, __Diagram__000001_Traffic_Light_UK_Diagram)
	__State__000011_Off_UK.Diagrams = append(__State__000011_Off_UK.Diagrams, __Diagram__000001_Traffic_Light_UK_Diagram)
	// setup of StateMachine instances pointers
	__StateMachine__000000_Traffic_Lights_FR.States = append(__StateMachine__000000_Traffic_Lights_FR.States, __State__000000_Initial_State_FR)
	__StateMachine__000000_Traffic_Lights_FR.States = append(__StateMachine__000000_Traffic_Lights_FR.States, __State__000001_Red_FR)
	__StateMachine__000000_Traffic_Lights_FR.States = append(__StateMachine__000000_Traffic_Lights_FR.States, __State__000002_Green_FR)
	__StateMachine__000000_Traffic_Lights_FR.States = append(__StateMachine__000000_Traffic_Lights_FR.States, __State__000003_Yellow_FR)
	__StateMachine__000000_Traffic_Lights_FR.States = append(__StateMachine__000000_Traffic_Lights_FR.States, __State__000004_On_FR)
	__StateMachine__000000_Traffic_Lights_FR.States = append(__StateMachine__000000_Traffic_Lights_FR.States, __State__000005_Off_FR)
	__StateMachine__000000_Traffic_Lights_FR.Diagrams = append(__StateMachine__000000_Traffic_Lights_FR.Diagrams, __Diagram__000000_Traffic_Light_FR_Diagram)
	__StateMachine__000000_Traffic_Lights_FR.InitialState = __State__000000_Initial_State_FR
	__StateMachine__000001_Traffic_Light_UK.States = append(__StateMachine__000001_Traffic_Light_UK.States, __State__000006_Initial_State_UK)
	__StateMachine__000001_Traffic_Light_UK.States = append(__StateMachine__000001_Traffic_Light_UK.States, __State__000007_Red_UK)
	__StateMachine__000001_Traffic_Light_UK.States = append(__StateMachine__000001_Traffic_Light_UK.States, __State__000008_Yellow_UK)
	__StateMachine__000001_Traffic_Light_UK.States = append(__StateMachine__000001_Traffic_Light_UK.States, __State__000009_Green_UK)
	__StateMachine__000001_Traffic_Light_UK.States = append(__StateMachine__000001_Traffic_Light_UK.States, __State__000010_On_UK)
	__StateMachine__000001_Traffic_Light_UK.States = append(__StateMachine__000001_Traffic_Light_UK.States, __State__000011_Off_UK)
	__StateMachine__000001_Traffic_Light_UK.Diagrams = append(__StateMachine__000001_Traffic_Light_UK.Diagrams, __Diagram__000001_Traffic_Light_UK_Diagram)
	__StateMachine__000001_Traffic_Light_UK.InitialState = __State__000006_Initial_State_UK
	// setup of StateShape instances pointers
	__StateShape__000000_Red_FR.State = __State__000001_Red_FR
	__StateShape__000001_Initial_State_FR.State = __State__000000_Initial_State_FR
	__StateShape__000002_Green_FR.State = __State__000002_Green_FR
	__StateShape__000003_Yellow_FR.State = __State__000003_Yellow_FR
	__StateShape__000004_On_FR.State = __State__000004_On_FR
	__StateShape__000005_Off_FR.State = __State__000005_Off_FR
	__StateShape__000006_Red_UK.State = __State__000007_Red_UK
	__StateShape__000007_Initial_State_UK.State = __State__000006_Initial_State_UK
	__StateShape__000008_Yellow_UK.State = __State__000008_Yellow_UK
	__StateShape__000009_Green_UK.State = __State__000009_Green_UK
	__StateShape__000010_On_UK.State = __State__000010_On_UK
	__StateShape__000011_Off_UK.State = __State__000011_Off_UK
	// setup of Transition instances pointers
	__Transition__000000_Initial_State_FR_to_Red_FR.Start = __State__000000_Initial_State_FR
	__Transition__000000_Initial_State_FR_to_Red_FR.End = __State__000001_Red_FR
	__Transition__000000_Initial_State_FR_to_Red_FR.RolesWithPermissions = append(__Transition__000000_Initial_State_FR_to_Red_FR.RolesWithPermissions, __Role__000000_Technician)
	__Transition__000000_Initial_State_FR_to_Red_FR.Diagrams = append(__Transition__000000_Initial_State_FR_to_Red_FR.Diagrams, __Diagram__000000_Traffic_Light_FR_Diagram)
	__Transition__000001_Red_FR_to_Green_FR.Start = __State__000001_Red_FR
	__Transition__000001_Red_FR_to_Green_FR.End = __State__000002_Green_FR
	__Transition__000001_Red_FR_to_Green_FR.RolesWithPermissions = append(__Transition__000001_Red_FR_to_Green_FR.RolesWithPermissions, __Role__000001_Timer)
	__Transition__000001_Red_FR_to_Green_FR.Diagrams = append(__Transition__000001_Red_FR_to_Green_FR.Diagrams, __Diagram__000000_Traffic_Light_FR_Diagram)
	__Transition__000002_Green_FR_to_Yellow_FR.Start = __State__000002_Green_FR
	__Transition__000002_Green_FR_to_Yellow_FR.End = __State__000003_Yellow_FR
	__Transition__000002_Green_FR_to_Yellow_FR.RolesWithPermissions = append(__Transition__000002_Green_FR_to_Yellow_FR.RolesWithPermissions, __Role__000001_Timer)
	__Transition__000002_Green_FR_to_Yellow_FR.Diagrams = append(__Transition__000002_Green_FR_to_Yellow_FR.Diagrams, __Diagram__000000_Traffic_Light_FR_Diagram)
	__Transition__000003_Yellow_FR_to_Red_FR.Start = __State__000003_Yellow_FR
	__Transition__000003_Yellow_FR_to_Red_FR.End = __State__000001_Red_FR
	__Transition__000003_Yellow_FR_to_Red_FR.RolesWithPermissions = append(__Transition__000003_Yellow_FR_to_Red_FR.RolesWithPermissions, __Role__000001_Timer)
	__Transition__000003_Yellow_FR_to_Red_FR.Diagrams = append(__Transition__000003_Yellow_FR_to_Red_FR.Diagrams, __Diagram__000000_Traffic_Light_FR_Diagram)
	__Transition__000004_On_to_Off.Start = __State__000004_On_FR
	__Transition__000004_On_to_Off.End = __State__000005_Off_FR
	__Transition__000004_On_to_Off.RolesWithPermissions = append(__Transition__000004_On_to_Off.RolesWithPermissions, __Role__000000_Technician)
	__Transition__000004_On_to_Off.Diagrams = append(__Transition__000004_On_to_Off.Diagrams, __Diagram__000000_Traffic_Light_FR_Diagram)
	__Transition__000005_Off_to_On.Start = __State__000005_Off_FR
	__Transition__000005_Off_to_On.End = __State__000001_Red_FR
	__Transition__000005_Off_to_On.RolesWithPermissions = append(__Transition__000005_Off_to_On.RolesWithPermissions, __Role__000000_Technician)
	__Transition__000005_Off_to_On.Diagrams = append(__Transition__000005_Off_to_On.Diagrams, __Diagram__000000_Traffic_Light_FR_Diagram)
	__Transition__000006_Initial_State_UK_to_New_State.Start = __State__000006_Initial_State_UK
	__Transition__000006_Initial_State_UK_to_New_State.End = __State__000007_Red_UK
	__Transition__000006_Initial_State_UK_to_New_State.RolesWithPermissions = append(__Transition__000006_Initial_State_UK_to_New_State.RolesWithPermissions, __Role__000000_Technician)
	__Transition__000006_Initial_State_UK_to_New_State.Diagrams = append(__Transition__000006_Initial_State_UK_to_New_State.Diagrams, __Diagram__000001_Traffic_Light_UK_Diagram)
	__Transition__000007_Red_UK_to_Yellow_UK.Start = __State__000007_Red_UK
	__Transition__000007_Red_UK_to_Yellow_UK.End = __State__000008_Yellow_UK
	__Transition__000007_Red_UK_to_Yellow_UK.RolesWithPermissions = append(__Transition__000007_Red_UK_to_Yellow_UK.RolesWithPermissions, __Role__000001_Timer)
	__Transition__000007_Red_UK_to_Yellow_UK.Diagrams = append(__Transition__000007_Red_UK_to_Yellow_UK.Diagrams, __Diagram__000001_Traffic_Light_UK_Diagram)
	__Transition__000008_Yellow_UK_to_Green_UK.Start = __State__000008_Yellow_UK
	__Transition__000008_Yellow_UK_to_Green_UK.End = __State__000009_Green_UK
	__Transition__000008_Yellow_UK_to_Green_UK.RolesWithPermissions = append(__Transition__000008_Yellow_UK_to_Green_UK.RolesWithPermissions, __Role__000001_Timer)
	__Transition__000008_Yellow_UK_to_Green_UK.Diagrams = append(__Transition__000008_Yellow_UK_to_Green_UK.Diagrams, __Diagram__000001_Traffic_Light_UK_Diagram)
	__Transition__000009_On_UK_to_Off_UK.Start = __State__000010_On_UK
	__Transition__000009_On_UK_to_Off_UK.End = __State__000011_Off_UK
	__Transition__000009_On_UK_to_Off_UK.RolesWithPermissions = append(__Transition__000009_On_UK_to_Off_UK.RolesWithPermissions, __Role__000000_Technician)
	__Transition__000009_On_UK_to_Off_UK.Diagrams = append(__Transition__000009_On_UK_to_Off_UK.Diagrams, __Diagram__000001_Traffic_Light_UK_Diagram)
	__Transition__000010_Off_UK_to_Red_UK.Start = __State__000011_Off_UK
	__Transition__000010_Off_UK_to_Red_UK.End = __State__000007_Red_UK
	__Transition__000010_Off_UK_to_Red_UK.RolesWithPermissions = append(__Transition__000010_Off_UK_to_Red_UK.RolesWithPermissions, __Role__000000_Technician)
	__Transition__000010_Off_UK_to_Red_UK.GeneratedMessages = append(__Transition__000010_Off_UK_to_Red_UK.GeneratedMessages, __MessageType__000000_Repair_Report)
	__Transition__000010_Off_UK_to_Red_UK.Diagrams = append(__Transition__000010_Off_UK_to_Red_UK.Diagrams, __Diagram__000001_Traffic_Light_UK_Diagram)
	__Transition__000011_Green_UK_to_Red_UK.Start = __State__000009_Green_UK
	__Transition__000011_Green_UK_to_Red_UK.End = __State__000007_Red_UK
	__Transition__000011_Green_UK_to_Red_UK.RolesWithPermissions = append(__Transition__000011_Green_UK_to_Red_UK.RolesWithPermissions, __Role__000001_Timer)
	__Transition__000011_Green_UK_to_Red_UK.Diagrams = append(__Transition__000011_Green_UK_to_Red_UK.Diagrams, __Diagram__000001_Traffic_Light_UK_Diagram)
	// setup of Transition_Shape instances pointers
	__Transition_Shape__000000_Initial_State_FR_to_Red_FR.Transition = __Transition__000000_Initial_State_FR_to_Red_FR
	__Transition_Shape__000001_Red_FR_to_Green_FR.Transition = __Transition__000001_Red_FR_to_Green_FR
	__Transition_Shape__000002_Green_FR_to_Yellow_FR.Transition = __Transition__000002_Green_FR_to_Yellow_FR
	__Transition_Shape__000003_Yellow_FR_to_Red_FR.Transition = __Transition__000003_Yellow_FR_to_Red_FR
	__Transition_Shape__000004_New_State_to_New_State.Transition = __Transition__000004_On_to_Off
	__Transition_Shape__000005_Off_to_On.Transition = __Transition__000005_Off_to_On
	__Transition_Shape__000006_Initial_State_UK_to_New_State.Transition = __Transition__000006_Initial_State_UK_to_New_State
	__Transition_Shape__000007_Red_UK_to_Yellow_UK.Transition = __Transition__000007_Red_UK_to_Yellow_UK
	__Transition_Shape__000008_Yellow_UK_to_Green_UK.Transition = __Transition__000008_Yellow_UK_to_Green_UK
	__Transition_Shape__000009_On_UK_to_Off_UK.Transition = __Transition__000009_On_UK_to_Off_UK
	__Transition_Shape__000010_Off_UK_to_On_UK.Transition = __Transition__000010_Off_UK_to_Red_UK
	__Transition_Shape__000011_Green_UK_to_Red_UK.Transition = __Transition__000011_Green_UK_to_Red_UK
}

