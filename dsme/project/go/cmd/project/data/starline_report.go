package main

import (
	"time"

	"github.com/fullstack-lang/gong/dsme/project/go/models"
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

	// insertion point for declaration of instances to stage

	__Diagram__00000000_ := (&models.Diagram{Name: `Default Diagram`}).Stage(stage)
	__Diagram__00000001_ := (&models.Diagram{Name: `WBS`}).Stage(stage)
	__Diagram__00000002_ := (&models.Diagram{Name: `PBS`}).Stage(stage)
	__Diagram__00000003_ := (&models.Diagram{Name: `RBS`}).Stage(stage)
	__Diagram__00000004_ := (&models.Diagram{Name: `PIT focus`}).Stage(stage)

	__Note__00000000_ := (&models.Note{Name: `CFT ended in march 2025`}).Stage(stage)
	__Note__00000001_ := (&models.Note{Name: `A thorough review of the STAR report is advised (p14)
`}).Stage(stage)

	__NoteProductShape__00000000_ := (&models.NoteProductShape{Name: `A thorough review of the STAR report is advised to  Starliner Tests and Anomalies Review (STAR) Investigation Report`}).Stage(stage)

	__NoteShape__00000000_ := (&models.NoteShape{Name: `-Default Diagram`}).Stage(stage)
	__NoteShape__00000001_ := (&models.NoteShape{Name: `-PIT focus`}).Stage(stage)

	__NoteTaskShape__00000000_ := (&models.NoteTaskShape{Name: `CFT ended in march 2025 to Starliner Crewed Flight Test (CFT)`}).Stage(stage)

	__Product__00000001_ := (&models.Product{Name: `Dragon`}).Stage(stage)
	__Product__00000002_ := (&models.Product{Name: `Starliner`}).Stage(stage)
	__Product__00000003_ := (&models.Product{Name: `Reports`}).Stage(stage)
	__Product__00000004_ := (&models.Product{Name: ` Starliner Tests and Anomalies Review (STAR) Investigation Report`}).Stage(stage)
	__Product__00000005_ := (&models.Product{Name: `Program Investigation Team (PIT) Report`}).Stage(stage)
	__Product__00000006_ := (&models.Product{Name: `Commercial Crew Transportation Capability (CCtCap).`}).Stage(stage)
	__Product__00000007_ := (&models.Product{Name: `NASA Assets/Capabities`}).Stage(stage)
	__Product__00000008_ := (&models.Product{Name: `ISS`}).Stage(stage)

	__ProductCompositionShape__00000000_ := (&models.ProductCompositionShape{Name: `Reports to `}).Stage(stage)
	__ProductCompositionShape__00000001_ := (&models.ProductCompositionShape{Name: `Reports to `}).Stage(stage)
	__ProductCompositionShape__00000002_ := (&models.ProductCompositionShape{Name: `Commercial Crew Transportation Capability (CCtCap). to Starliner`}).Stage(stage)
	__ProductCompositionShape__00000003_ := (&models.ProductCompositionShape{Name: `Commercial Crew Transportation Capability (CCtCap). to Dragon`}).Stage(stage)
	__ProductCompositionShape__00000004_ := (&models.ProductCompositionShape{Name: `NASA Assets/Capabities to Commercial Crew Transportation Capability (CCtCap).`}).Stage(stage)
	__ProductCompositionShape__00000005_ := (&models.ProductCompositionShape{Name: `Commercial Crew Transportation Capability (CCtCap). to `}).Stage(stage)

	__ProductShape__00000002_ := (&models.ProductShape{Name: `Dragon-PBS`}).Stage(stage)
	__ProductShape__00000003_ := (&models.ProductShape{Name: `Starliner-PBS`}).Stage(stage)
	__ProductShape__00000004_ := (&models.ProductShape{Name: `-PBS`}).Stage(stage)
	__ProductShape__00000005_ := (&models.ProductShape{Name: `-PBS`}).Stage(stage)
	__ProductShape__00000006_ := (&models.ProductShape{Name: `-PBS`}).Stage(stage)
	__ProductShape__00000007_ := (&models.ProductShape{Name: `Program Investigation Team (PIT) Report-PIT focus`}).Stage(stage)
	__ProductShape__00000008_ := (&models.ProductShape{Name: ` Starliner Tests and Anomalies Review (STAR) Investigation Report-PIT focus`}).Stage(stage)
	__ProductShape__00000009_ := (&models.ProductShape{Name: `-PBS`}).Stage(stage)
	__ProductShape__00000010_ := (&models.ProductShape{Name: `-PBS`}).Stage(stage)
	__ProductShape__00000011_ := (&models.ProductShape{Name: `-PBS`}).Stage(stage)

	__Project__00000000_ := (&models.Project{Name: `Startliner Mishape Report`}).Stage(stage)

	__Resource__00000000_ := (&models.Resource{Name: `Program Investigation Team (PIT)`}).Stage(stage)
	__Resource__00000001_ := (&models.Resource{Name: `Barry "Butch" Wilmore`}).Stage(stage)
	__Resource__00000002_ := (&models.Resource{Name: `Sunita "Suni" Williams`}).Stage(stage)
	__Resource__00000003_ := (&models.Resource{Name: `NASA`}).Stage(stage)
	__Resource__00000004_ := (&models.Resource{Name: `Crew Commercial Program (CPP)`}).Stage(stage)
	__Resource__00000005_ := (&models.Resource{Name: `Crews`}).Stage(stage)
	__Resource__00000006_ := (&models.Resource{Name: `Boeing`}).Stage(stage)
	__Resource__00000007_ := (&models.Resource{Name: ``}).Stage(stage)
	__Resource__00000008_ := (&models.Resource{Name: ``}).Stage(stage)
	__Resource__00000009_ := (&models.Resource{Name: ` Starliner Tests and Anomalies Review (STAR) Investigation Team`}).Stage(stage)

	__ResourceCompositionShape__00000000_ := (&models.ResourceCompositionShape{Name: `NASA to `}).Stage(stage)
	__ResourceCompositionShape__00000001_ := (&models.ResourceCompositionShape{Name: `Crew Commercial Program (CPP) to `}).Stage(stage)
	__ResourceCompositionShape__00000002_ := (&models.ResourceCompositionShape{Name: `Crews to Barry "Butch" Wilmore`}).Stage(stage)
	__ResourceCompositionShape__00000003_ := (&models.ResourceCompositionShape{Name: `Crews to Sunita "Suni" Williams`}).Stage(stage)
	__ResourceCompositionShape__00000004_ := (&models.ResourceCompositionShape{Name: `NASA to PITProgram Investigation Team (PIT)`}).Stage(stage)
	__ResourceCompositionShape__00000005_ := (&models.ResourceCompositionShape{Name: `Boeing to `}).Stage(stage)
	__ResourceCompositionShape__00000006_ := (&models.ResourceCompositionShape{Name: `NASA to `}).Stage(stage)
	__ResourceCompositionShape__00000007_ := (&models.ResourceCompositionShape{Name: `Crew Commercial Program (CPP) to `}).Stage(stage)

	__ResourceShape__00000000_ := (&models.ResourceShape{Name: `-Default Diagram`}).Stage(stage)
	__ResourceShape__00000001_ := (&models.ResourceShape{Name: `-Default Diagram`}).Stage(stage)
	__ResourceShape__00000002_ := (&models.ResourceShape{Name: `-Default Diagram`}).Stage(stage)
	__ResourceShape__00000003_ := (&models.ResourceShape{Name: `-RBS`}).Stage(stage)
	__ResourceShape__00000004_ := (&models.ResourceShape{Name: `-RBS`}).Stage(stage)
	__ResourceShape__00000005_ := (&models.ResourceShape{Name: `-RBS`}).Stage(stage)
	__ResourceShape__00000006_ := (&models.ResourceShape{Name: `Barry "Butch" Wilmore-RBS`}).Stage(stage)
	__ResourceShape__00000007_ := (&models.ResourceShape{Name: `Sunita "Suni" Williams-RBS`}).Stage(stage)
	__ResourceShape__00000008_ := (&models.ResourceShape{Name: `PITProgram Investigation Team (PIT)-RBS`}).Stage(stage)
	__ResourceShape__00000009_ := (&models.ResourceShape{Name: `-RBS`}).Stage(stage)
	__ResourceShape__00000010_ := (&models.ResourceShape{Name: `-RBS`}).Stage(stage)
	__ResourceShape__00000011_ := (&models.ResourceShape{Name: `-RBS`}).Stage(stage)
	__ResourceShape__00000012_ := (&models.ResourceShape{Name: `-RBS`}).Stage(stage)
	__ResourceShape__00000013_ := (&models.ResourceShape{Name: `Program Investigation Team (PIT)-PBS`}).Stage(stage)
	__ResourceShape__00000014_ := (&models.ResourceShape{Name: `Program Investigation Team (PIT)-PIT focus`}).Stage(stage)

	__ResourceTaskShape__00000000_ := (&models.ResourceTaskShape{Name: `PITProgram Investigation Team (PIT) to Mishap investigation`}).Stage(stage)

	__Root__00000000_ := (&models.Root{Name: `Root`}).Stage(stage)

	__Task__00000000_ := (&models.Task{Name: `Starliner Crewed Flight Test (CFT)`}).Stage(stage)
	__Task__00000001_ := (&models.Task{Name: `Mishap investigations`}).Stage(stage)
	__Task__00000002_ := (&models.Task{Name: `Commercial Crew Program (CCP),`}).Stage(stage)
	__Task__00000003_ := (&models.Task{Name: ` Commercial ReSupply (CRS) `}).Stage(stage)
	__Task__00000004_ := (&models.Task{Name: `Program Investigation Team (PIT) Report`}).Stage(stage)

	__TaskCompositionShape__00000000_ := (&models.TaskCompositionShape{Name: `Commercial Crew Program (CCP), to Starliner Crewed Flight Test (CFT)`}).Stage(stage)
	__TaskCompositionShape__00000001_ := (&models.TaskCompositionShape{Name: `Commercial Crew Program (CCP), to Starliner Crewed Flight Test (CFT)`}).Stage(stage)
	__TaskCompositionShape__00000002_ := (&models.TaskCompositionShape{Name: `Mishap investigations to `}).Stage(stage)

	__TaskInputShape__00000000_ := (&models.TaskInputShape{Name: `Program Investigation Team (PIT) Report to  Starliner Tests and Anomalies Review (STAR) Investigation Report`}).Stage(stage)

	__TaskOutputShape__00000000_ := (&models.TaskOutputShape{Name: `Program Investigation Team (PIT) Report to Program Investigation Team (PIT) Report`}).Stage(stage)

	__TaskShape__00000000_ := (&models.TaskShape{Name: `NewTask-Default Diagram`}).Stage(stage)
	__TaskShape__00000001_ := (&models.TaskShape{Name: `NewTask-Default Diagram`}).Stage(stage)
	__TaskShape__00000002_ := (&models.TaskShape{Name: `-Default Diagram`}).Stage(stage)
	__TaskShape__00000003_ := (&models.TaskShape{Name: `Mishap investigation-WBS`}).Stage(stage)
	__TaskShape__00000004_ := (&models.TaskShape{Name: `Commercial Crew Program (CCP),-WBS`}).Stage(stage)
	__TaskShape__00000005_ := (&models.TaskShape{Name: `Starliner Crewed Flight Test (CFT)-WBS`}).Stage(stage)
	__TaskShape__00000006_ := (&models.TaskShape{Name: `-WBS`}).Stage(stage)
	__TaskShape__00000007_ := (&models.TaskShape{Name: `Mishap investigation-PBS`}).Stage(stage)
	__TaskShape__00000008_ := (&models.TaskShape{Name: `-WBS`}).Stage(stage)
	__TaskShape__00000009_ := (&models.TaskShape{Name: `Program Investigation Team (PIT) Report-PIT focus`}).Stage(stage)

	// insertion point for initialization of values

	__Diagram__00000000_.Name = `Default Diagram`
	__Diagram__00000000_.IsChecked = false
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.ShowPrefix = false
	__Diagram__00000000_.DefaultBoxWidth = 250.000000
	__Diagram__00000000_.DefaultBoxHeigth = 70.000000
	__Diagram__00000000_.IsExpanded = false
	__Diagram__00000000_.ComputedPrefix = ``
	__Diagram__00000000_.IsInRenameMode = false
	__Diagram__00000000_.IsPBSNodeExpanded = true
	__Diagram__00000000_.IsWBSNodeExpanded = true
	__Diagram__00000000_.IsNotesNodeExpanded = true
	__Diagram__00000000_.IsResourcesNodeExpanded = true

	__Diagram__00000001_.Name = `WBS`
	__Diagram__00000001_.IsChecked = false
	__Diagram__00000001_.IsEditable_ = true
	__Diagram__00000001_.ShowPrefix = false
	__Diagram__00000001_.DefaultBoxWidth = 250.000000
	__Diagram__00000001_.DefaultBoxHeigth = 70.000000
	__Diagram__00000001_.IsExpanded = true
	__Diagram__00000001_.ComputedPrefix = ``
	__Diagram__00000001_.IsInRenameMode = false
	__Diagram__00000001_.IsPBSNodeExpanded = false
	__Diagram__00000001_.IsWBSNodeExpanded = false
	__Diagram__00000001_.IsNotesNodeExpanded = false
	__Diagram__00000001_.IsResourcesNodeExpanded = false

	__Diagram__00000002_.Name = `PBS`
	__Diagram__00000002_.IsChecked = true
	__Diagram__00000002_.IsEditable_ = true
	__Diagram__00000002_.ShowPrefix = true
	__Diagram__00000002_.DefaultBoxWidth = 250.000000
	__Diagram__00000002_.DefaultBoxHeigth = 70.000000
	__Diagram__00000002_.IsExpanded = true
	__Diagram__00000002_.ComputedPrefix = ``
	__Diagram__00000002_.IsInRenameMode = false
	__Diagram__00000002_.IsPBSNodeExpanded = true
	__Diagram__00000002_.IsWBSNodeExpanded = true
	__Diagram__00000002_.IsNotesNodeExpanded = false
	__Diagram__00000002_.IsResourcesNodeExpanded = true

	__Diagram__00000003_.Name = `RBS`
	__Diagram__00000003_.IsChecked = false
	__Diagram__00000003_.IsEditable_ = true
	__Diagram__00000003_.ShowPrefix = false
	__Diagram__00000003_.DefaultBoxWidth = 250.000000
	__Diagram__00000003_.DefaultBoxHeigth = 70.000000
	__Diagram__00000003_.IsExpanded = false
	__Diagram__00000003_.ComputedPrefix = ``
	__Diagram__00000003_.IsInRenameMode = false
	__Diagram__00000003_.IsPBSNodeExpanded = true
	__Diagram__00000003_.IsWBSNodeExpanded = false
	__Diagram__00000003_.IsNotesNodeExpanded = false
	__Diagram__00000003_.IsResourcesNodeExpanded = true

	__Diagram__00000004_.Name = `PIT focus`
	__Diagram__00000004_.IsChecked = false
	__Diagram__00000004_.IsEditable_ = true
	__Diagram__00000004_.ShowPrefix = false
	__Diagram__00000004_.DefaultBoxWidth = 250.000000
	__Diagram__00000004_.DefaultBoxHeigth = 70.000000
	__Diagram__00000004_.IsExpanded = true
	__Diagram__00000004_.ComputedPrefix = ``
	__Diagram__00000004_.IsInRenameMode = false
	__Diagram__00000004_.IsPBSNodeExpanded = true
	__Diagram__00000004_.IsWBSNodeExpanded = false
	__Diagram__00000004_.IsNotesNodeExpanded = true
	__Diagram__00000004_.IsResourcesNodeExpanded = false

	__Note__00000000_.Name = `CFT ended in march 2025`
	__Note__00000000_.IsExpanded = false
	__Note__00000000_.ComputedPrefix = `1`
	__Note__00000000_.IsInRenameMode = false

	__Note__00000001_.Name = `A thorough review of the STAR report is advised (p14)
`
	__Note__00000001_.IsExpanded = false
	__Note__00000001_.ComputedPrefix = `2`
	__Note__00000001_.IsInRenameMode = false

	__NoteProductShape__00000000_.Name = `A thorough review of the STAR report is advised to  Starliner Tests and Anomalies Review (STAR) Investigation Report`
	__NoteProductShape__00000000_.StartRatio = 0.500000
	__NoteProductShape__00000000_.EndRatio = 0.500000
	__NoteProductShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000000_.CornerOffsetRatio = 1.680000

	__NoteShape__00000000_.Name = `-Default Diagram`
	__NoteShape__00000000_.IsExpanded = false
	__NoteShape__00000000_.X = 92.441418
	__NoteShape__00000000_.Y = 1047.751617
	__NoteShape__00000000_.Width = 250.000000
	__NoteShape__00000000_.Height = 70.000000

	__NoteShape__00000001_.Name = `-PIT focus`
	__NoteShape__00000001_.IsExpanded = false
	__NoteShape__00000001_.X = 892.606649
	__NoteShape__00000001_.Y = 74.204707
	__NoteShape__00000001_.Width = 250.000000
	__NoteShape__00000001_.Height = 70.000000

	__NoteTaskShape__00000000_.Name = `CFT ended in march 2025 to Starliner Crewed Flight Test (CFT)`
	__NoteTaskShape__00000000_.StartRatio = 0.500000
	__NoteTaskShape__00000000_.EndRatio = 0.500000
	__NoteTaskShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteTaskShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteTaskShape__00000000_.CornerOffsetRatio = 1.680000

	__Product__00000001_.Name = `Dragon`
	__Product__00000001_.Description = ``
	__Product__00000001_.IsExpanded = false
	__Product__00000001_.ComputedPrefix = `1.1.2`
	__Product__00000001_.IsInRenameMode = false
	__Product__00000001_.IsProducersNodeExpanded = false
	__Product__00000001_.IsConsumersNodeExpanded = false

	__Product__00000002_.Name = `Starliner`
	__Product__00000002_.Description = ``
	__Product__00000002_.IsExpanded = false
	__Product__00000002_.ComputedPrefix = `1.1.1`
	__Product__00000002_.IsInRenameMode = false
	__Product__00000002_.IsProducersNodeExpanded = false
	__Product__00000002_.IsConsumersNodeExpanded = false

	__Product__00000003_.Name = `Reports`
	__Product__00000003_.Description = ``
	__Product__00000003_.IsExpanded = false
	__Product__00000003_.ComputedPrefix = `2`
	__Product__00000003_.IsInRenameMode = false
	__Product__00000003_.IsProducersNodeExpanded = false
	__Product__00000003_.IsConsumersNodeExpanded = false

	__Product__00000004_.Name = ` Starliner Tests and Anomalies Review (STAR) Investigation Report`
	__Product__00000004_.Description = ``
	__Product__00000004_.IsExpanded = false
	__Product__00000004_.ComputedPrefix = `2.1`
	__Product__00000004_.IsInRenameMode = false
	__Product__00000004_.IsProducersNodeExpanded = false
	__Product__00000004_.IsConsumersNodeExpanded = false

	__Product__00000005_.Name = `Program Investigation Team (PIT) Report`
	__Product__00000005_.Description = ``
	__Product__00000005_.IsExpanded = false
	__Product__00000005_.ComputedPrefix = `2.2`
	__Product__00000005_.IsInRenameMode = false
	__Product__00000005_.IsProducersNodeExpanded = false
	__Product__00000005_.IsConsumersNodeExpanded = false

	__Product__00000006_.Name = `Commercial Crew Transportation Capability (CCtCap).`
	__Product__00000006_.Description = ``
	__Product__00000006_.IsExpanded = false
	__Product__00000006_.ComputedPrefix = `1.1`
	__Product__00000006_.IsInRenameMode = false
	__Product__00000006_.IsProducersNodeExpanded = false
	__Product__00000006_.IsConsumersNodeExpanded = false

	__Product__00000007_.Name = `NASA Assets/Capabities`
	__Product__00000007_.Description = ``
	__Product__00000007_.IsExpanded = false
	__Product__00000007_.ComputedPrefix = `1`
	__Product__00000007_.IsInRenameMode = false
	__Product__00000007_.IsProducersNodeExpanded = false
	__Product__00000007_.IsConsumersNodeExpanded = false

	__Product__00000008_.Name = `ISS`
	__Product__00000008_.Description = ``
	__Product__00000008_.IsExpanded = false
	__Product__00000008_.ComputedPrefix = `1.1.3`
	__Product__00000008_.IsInRenameMode = false
	__Product__00000008_.IsProducersNodeExpanded = false
	__Product__00000008_.IsConsumersNodeExpanded = false

	__ProductCompositionShape__00000000_.Name = `Reports to `
	__ProductCompositionShape__00000000_.StartRatio = 0.500000
	__ProductCompositionShape__00000000_.EndRatio = 0.500000
	__ProductCompositionShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000000_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000001_.Name = `Reports to `
	__ProductCompositionShape__00000001_.StartRatio = 0.500000
	__ProductCompositionShape__00000001_.EndRatio = 0.500000
	__ProductCompositionShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000001_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000002_.Name = `Commercial Crew Transportation Capability (CCtCap). to Starliner`
	__ProductCompositionShape__00000002_.StartRatio = 0.500000
	__ProductCompositionShape__00000002_.EndRatio = 0.500000
	__ProductCompositionShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000002_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000003_.Name = `Commercial Crew Transportation Capability (CCtCap). to Dragon`
	__ProductCompositionShape__00000003_.StartRatio = 0.500000
	__ProductCompositionShape__00000003_.EndRatio = 0.500000
	__ProductCompositionShape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000003_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000003_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000004_.Name = `NASA Assets/Capabities to Commercial Crew Transportation Capability (CCtCap).`
	__ProductCompositionShape__00000004_.StartRatio = 0.500000
	__ProductCompositionShape__00000004_.EndRatio = 0.500000
	__ProductCompositionShape__00000004_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000004_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000005_.Name = `Commercial Crew Transportation Capability (CCtCap). to `
	__ProductCompositionShape__00000005_.StartRatio = 0.500000
	__ProductCompositionShape__00000005_.EndRatio = 0.500000
	__ProductCompositionShape__00000005_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000005_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000005_.CornerOffsetRatio = 1.680000

	__ProductShape__00000002_.Name = `Dragon-PBS`
	__ProductShape__00000002_.IsExpanded = false
	__ProductShape__00000002_.X = 363.020302
	__ProductShape__00000002_.Y = 904.039767
	__ProductShape__00000002_.Width = 250.000000
	__ProductShape__00000002_.Height = 70.000000

	__ProductShape__00000003_.Name = `Starliner-PBS`
	__ProductShape__00000003_.IsExpanded = false
	__ProductShape__00000003_.X = 49.852191
	__ProductShape__00000003_.Y = 897.054390
	__ProductShape__00000003_.Width = 250.000000
	__ProductShape__00000003_.Height = 70.000000

	__ProductShape__00000004_.Name = `-PBS`
	__ProductShape__00000004_.IsExpanded = false
	__ProductShape__00000004_.X = 495.980870
	__ProductShape__00000004_.Y = 389.166629
	__ProductShape__00000004_.Width = 250.000000
	__ProductShape__00000004_.Height = 70.000000

	__ProductShape__00000005_.Name = `-PBS`
	__ProductShape__00000005_.IsExpanded = false
	__ProductShape__00000005_.X = 495.980870
	__ProductShape__00000005_.Y = 529.166629
	__ProductShape__00000005_.Width = 250.000000
	__ProductShape__00000005_.Height = 70.000000

	__ProductShape__00000006_.Name = `-PBS`
	__ProductShape__00000006_.IsExpanded = false
	__ProductShape__00000006_.X = 795.980870
	__ProductShape__00000006_.Y = 529.166629
	__ProductShape__00000006_.Width = 250.000000
	__ProductShape__00000006_.Height = 70.000000

	__ProductShape__00000007_.Name = `Program Investigation Team (PIT) Report-PIT focus`
	__ProductShape__00000007_.IsExpanded = false
	__ProductShape__00000007_.X = 912.017987
	__ProductShape__00000007_.Y = 275.867255
	__ProductShape__00000007_.Width = 250.000000
	__ProductShape__00000007_.Height = 70.000000

	__ProductShape__00000008_.Name = ` Starliner Tests and Anomalies Review (STAR) Investigation Report-PIT focus`
	__ProductShape__00000008_.IsExpanded = false
	__ProductShape__00000008_.X = 499.019526
	__ProductShape__00000008_.Y = 77.170905
	__ProductShape__00000008_.Width = 250.000000
	__ProductShape__00000008_.Height = 70.000000

	__ProductShape__00000009_.Name = `-PBS`
	__ProductShape__00000009_.IsExpanded = false
	__ProductShape__00000009_.X = 144.071765
	__ProductShape__00000009_.Y = 732.961381
	__ProductShape__00000009_.Width = 250.000000
	__ProductShape__00000009_.Height = 70.000000

	__ProductShape__00000010_.Name = `-PBS`
	__ProductShape__00000010_.IsExpanded = false
	__ProductShape__00000010_.X = 114.619741
	__ProductShape__00000010_.Y = 565.878708
	__ProductShape__00000010_.Width = 250.000000
	__ProductShape__00000010_.Height = 70.000000

	__ProductShape__00000011_.Name = `-PBS`
	__ProductShape__00000011_.IsExpanded = false
	__ProductShape__00000011_.X = 670.071765
	__ProductShape__00000011_.Y = 906.961381
	__ProductShape__00000011_.Width = 250.000000
	__ProductShape__00000011_.Height = 70.000000

	__Project__00000000_.Name = `Startliner Mishape Report`
	__Project__00000000_.IsExpanded = true
	__Project__00000000_.ComputedPrefix = ``
	__Project__00000000_.IsInRenameMode = false

	__Resource__00000000_.Name = `Program Investigation Team (PIT)`
	__Resource__00000000_.Description = ``
	__Resource__00000000_.IsExpanded = false
	__Resource__00000000_.ComputedPrefix = `1.2`
	__Resource__00000000_.IsInRenameMode = false

	__Resource__00000001_.Name = `Barry "Butch" Wilmore`
	__Resource__00000001_.Description = ``
	__Resource__00000001_.IsExpanded = false
	__Resource__00000001_.ComputedPrefix = `1.1.1.1`
	__Resource__00000001_.IsInRenameMode = false

	__Resource__00000002_.Name = `Sunita "Suni" Williams`
	__Resource__00000002_.Description = ``
	__Resource__00000002_.IsExpanded = false
	__Resource__00000002_.ComputedPrefix = `1.1.1.2`
	__Resource__00000002_.IsInRenameMode = false

	__Resource__00000003_.Name = `NASA`
	__Resource__00000003_.Description = ``
	__Resource__00000003_.IsExpanded = false
	__Resource__00000003_.ComputedPrefix = `1`
	__Resource__00000003_.IsInRenameMode = false

	__Resource__00000004_.Name = `Crew Commercial Program (CPP)`
	__Resource__00000004_.Description = ``
	__Resource__00000004_.IsExpanded = false
	__Resource__00000004_.ComputedPrefix = `1.1`
	__Resource__00000004_.IsInRenameMode = false

	__Resource__00000005_.Name = `Crews`
	__Resource__00000005_.Description = ``
	__Resource__00000005_.IsExpanded = false
	__Resource__00000005_.ComputedPrefix = `1.1.1`
	__Resource__00000005_.IsInRenameMode = false

	__Resource__00000006_.Name = `Boeing`
	__Resource__00000006_.Description = ``
	__Resource__00000006_.IsExpanded = false
	__Resource__00000006_.ComputedPrefix = `2`
	__Resource__00000006_.IsInRenameMode = false

	__Resource__00000007_.Name = ``
	__Resource__00000007_.Description = ``
	__Resource__00000007_.IsExpanded = false
	__Resource__00000007_.ComputedPrefix = `2.1`
	__Resource__00000007_.IsInRenameMode = false

	__Resource__00000008_.Name = ``
	__Resource__00000008_.Description = ``
	__Resource__00000008_.IsExpanded = false
	__Resource__00000008_.ComputedPrefix = `1.3`
	__Resource__00000008_.IsInRenameMode = false

	__Resource__00000009_.Name = ` Starliner Tests and Anomalies Review (STAR) Investigation Team`
	__Resource__00000009_.Description = ``
	__Resource__00000009_.IsExpanded = false
	__Resource__00000009_.ComputedPrefix = `1.1.2`
	__Resource__00000009_.IsInRenameMode = false

	__ResourceCompositionShape__00000000_.Name = `NASA to `
	__ResourceCompositionShape__00000000_.StartRatio = 0.500000
	__ResourceCompositionShape__00000000_.EndRatio = 0.500000
	__ResourceCompositionShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000000_.CornerOffsetRatio = 1.680000

	__ResourceCompositionShape__00000001_.Name = `Crew Commercial Program (CPP) to `
	__ResourceCompositionShape__00000001_.StartRatio = 0.500000
	__ResourceCompositionShape__00000001_.EndRatio = 0.500000
	__ResourceCompositionShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000001_.CornerOffsetRatio = 1.680000

	__ResourceCompositionShape__00000002_.Name = `Crews to Barry "Butch" Wilmore`
	__ResourceCompositionShape__00000002_.StartRatio = 0.500000
	__ResourceCompositionShape__00000002_.EndRatio = 0.500000
	__ResourceCompositionShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000002_.CornerOffsetRatio = 1.680000

	__ResourceCompositionShape__00000003_.Name = `Crews to Sunita "Suni" Williams`
	__ResourceCompositionShape__00000003_.StartRatio = 0.500000
	__ResourceCompositionShape__00000003_.EndRatio = 0.500000
	__ResourceCompositionShape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000003_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000003_.CornerOffsetRatio = 1.680000

	__ResourceCompositionShape__00000004_.Name = `NASA to PITProgram Investigation Team (PIT)`
	__ResourceCompositionShape__00000004_.StartRatio = 0.500000
	__ResourceCompositionShape__00000004_.EndRatio = 0.500000
	__ResourceCompositionShape__00000004_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000004_.CornerOffsetRatio = 1.680000

	__ResourceCompositionShape__00000005_.Name = `Boeing to `
	__ResourceCompositionShape__00000005_.StartRatio = 0.500000
	__ResourceCompositionShape__00000005_.EndRatio = 0.500000
	__ResourceCompositionShape__00000005_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000005_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000005_.CornerOffsetRatio = 1.680000

	__ResourceCompositionShape__00000006_.Name = `NASA to `
	__ResourceCompositionShape__00000006_.StartRatio = 0.500000
	__ResourceCompositionShape__00000006_.EndRatio = 0.500000
	__ResourceCompositionShape__00000006_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000006_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000006_.CornerOffsetRatio = 1.680000

	__ResourceCompositionShape__00000007_.Name = `Crew Commercial Program (CPP) to `
	__ResourceCompositionShape__00000007_.StartRatio = 0.500000
	__ResourceCompositionShape__00000007_.EndRatio = 0.500000
	__ResourceCompositionShape__00000007_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000007_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000007_.CornerOffsetRatio = 1.680000

	__ResourceShape__00000000_.Name = `-Default Diagram`
	__ResourceShape__00000000_.IsExpanded = false
	__ResourceShape__00000000_.X = 52.114853
	__ResourceShape__00000000_.Y = 34.127119
	__ResourceShape__00000000_.Width = 250.000000
	__ResourceShape__00000000_.Height = 70.000000

	__ResourceShape__00000001_.Name = `-Default Diagram`
	__ResourceShape__00000001_.IsExpanded = false
	__ResourceShape__00000001_.X = 34.180389
	__ResourceShape__00000001_.Y = 553.010316
	__ResourceShape__00000001_.Width = 250.000000
	__ResourceShape__00000001_.Height = 70.000000

	__ResourceShape__00000002_.Name = `-Default Diagram`
	__ResourceShape__00000002_.IsExpanded = false
	__ResourceShape__00000002_.X = 33.201744
	__ResourceShape__00000002_.Y = 650.719328
	__ResourceShape__00000002_.Width = 250.000000
	__ResourceShape__00000002_.Height = 70.000000

	__ResourceShape__00000003_.Name = `-RBS`
	__ResourceShape__00000003_.IsExpanded = false
	__ResourceShape__00000003_.X = 134.737820
	__ResourceShape__00000003_.Y = 123.124371
	__ResourceShape__00000003_.Width = 250.000000
	__ResourceShape__00000003_.Height = 70.000000

	__ResourceShape__00000004_.Name = `-RBS`
	__ResourceShape__00000004_.IsExpanded = false
	__ResourceShape__00000004_.X = 137.737820
	__ResourceShape__00000004_.Y = 300.124371
	__ResourceShape__00000004_.Width = 250.000000
	__ResourceShape__00000004_.Height = 70.000000

	__ResourceShape__00000005_.Name = `-RBS`
	__ResourceShape__00000005_.IsExpanded = false
	__ResourceShape__00000005_.X = 137.737820
	__ResourceShape__00000005_.Y = 440.124371
	__ResourceShape__00000005_.Width = 250.000000
	__ResourceShape__00000005_.Height = 70.000000

	__ResourceShape__00000006_.Name = `Barry "Butch" Wilmore-RBS`
	__ResourceShape__00000006_.IsExpanded = false
	__ResourceShape__00000006_.X = 132.423066
	__ResourceShape__00000006_.Y = 597.719206
	__ResourceShape__00000006_.Width = 250.000000
	__ResourceShape__00000006_.Height = 70.000000

	__ResourceShape__00000007_.Name = `Sunita "Suni" Williams-RBS`
	__ResourceShape__00000007_.IsExpanded = false
	__ResourceShape__00000007_.X = 454.596590
	__ResourceShape__00000007_.Y = 600.173621
	__ResourceShape__00000007_.Width = 250.000000
	__ResourceShape__00000007_.Height = 70.000000

	__ResourceShape__00000008_.Name = `PITProgram Investigation Team (PIT)-RBS`
	__ResourceShape__00000008_.IsExpanded = false
	__ResourceShape__00000008_.X = 835.651635
	__ResourceShape__00000008_.Y = 297.886860
	__ResourceShape__00000008_.Width = 250.000000
	__ResourceShape__00000008_.Height = 70.000000

	__ResourceShape__00000009_.Name = `-RBS`
	__ResourceShape__00000009_.IsExpanded = false
	__ResourceShape__00000009_.X = 126.434105
	__ResourceShape__00000009_.Y = 846.201664
	__ResourceShape__00000009_.Width = 250.000000
	__ResourceShape__00000009_.Height = 70.000000

	__ResourceShape__00000010_.Name = `-RBS`
	__ResourceShape__00000010_.IsExpanded = false
	__ResourceShape__00000010_.X = 110.434105
	__ResourceShape__00000010_.Y = 1029.201740
	__ResourceShape__00000010_.Width = 250.000000
	__ResourceShape__00000010_.Height = 70.000000

	__ResourceShape__00000011_.Name = `-RBS`
	__ResourceShape__00000011_.IsExpanded = false
	__ResourceShape__00000011_.X = 136.986696
	__ResourceShape__00000011_.Y = 123.307852
	__ResourceShape__00000011_.Width = 250.000000
	__ResourceShape__00000011_.Height = 70.000000

	__ResourceShape__00000012_.Name = `-RBS`
	__ResourceShape__00000012_.IsExpanded = false
	__ResourceShape__00000012_.X = 437.737820
	__ResourceShape__00000012_.Y = 440.124371
	__ResourceShape__00000012_.Width = 250.000000
	__ResourceShape__00000012_.Height = 70.000000

	__ResourceShape__00000013_.Name = `Program Investigation Team (PIT)-PBS`
	__ResourceShape__00000013_.IsExpanded = false
	__ResourceShape__00000013_.X = 782.684113
	__ResourceShape__00000013_.Y = 776.417682
	__ResourceShape__00000013_.Width = 250.000000
	__ResourceShape__00000013_.Height = 70.000000

	__ResourceShape__00000014_.Name = `Program Investigation Team (PIT)-PIT focus`
	__ResourceShape__00000014_.IsExpanded = false
	__ResourceShape__00000014_.X = 104.267926
	__ResourceShape__00000014_.Y = 288.658205
	__ResourceShape__00000014_.Width = 250.000000
	__ResourceShape__00000014_.Height = 70.000000

	__ResourceTaskShape__00000000_.Name = `PITProgram Investigation Team (PIT) to Mishap investigation`
	__ResourceTaskShape__00000000_.StartRatio = 0.500000
	__ResourceTaskShape__00000000_.EndRatio = 0.500000
	__ResourceTaskShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceTaskShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceTaskShape__00000000_.CornerOffsetRatio = 1.680000

	__Root__00000000_.Name = `Root`
	__Root__00000000_.NbPixPerCharacter = 8.000000

	__Task__00000000_.Name = `Starliner Crewed Flight Test (CFT)`
	__Task__00000000_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2024-06-05 00:00:00 +0000 UTC")
	__Task__00000000_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2024-06-10 00:00:00 +0000 UTC")
	__Task__00000000_.Description = `The mission of interest to the report.`
	__Task__00000000_.IsExpanded = false
	__Task__00000000_.ComputedPrefix = `2.1`
	__Task__00000000_.IsInRenameMode = false
	__Task__00000000_.IsInputsNodeExpanded = false
	__Task__00000000_.IsOutputsNodeExpanded = false
	__Task__00000000_.IsWithCompletion = false
	__Task__00000000_.Completion = ""

	__Task__00000001_.Name = `Mishap investigations`
	__Task__00000001_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2025-02-01 00:00:00 +0000 UTC")
	__Task__00000001_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2025-11-01 00:00:00 +0000 UTC")
	__Task__00000001_.Description = ``
	__Task__00000001_.IsExpanded = false
	__Task__00000001_.ComputedPrefix = `1`
	__Task__00000001_.IsInRenameMode = false
	__Task__00000001_.IsInputsNodeExpanded = false
	__Task__00000001_.IsOutputsNodeExpanded = false
	__Task__00000001_.IsWithCompletion = false
	__Task__00000001_.Completion = ""

	__Task__00000002_.Name = `Commercial Crew Program (CCP),`
	__Task__00000002_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000002_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000002_.Description = ``
	__Task__00000002_.IsExpanded = false
	__Task__00000002_.ComputedPrefix = `2`
	__Task__00000002_.IsInRenameMode = false
	__Task__00000002_.IsInputsNodeExpanded = false
	__Task__00000002_.IsOutputsNodeExpanded = false
	__Task__00000002_.IsWithCompletion = false
	__Task__00000002_.Completion = ""

	__Task__00000003_.Name = ` Commercial ReSupply (CRS) `
	__Task__00000003_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000003_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000003_.Description = ``
	__Task__00000003_.IsExpanded = false
	__Task__00000003_.ComputedPrefix = `3`
	__Task__00000003_.IsInRenameMode = false
	__Task__00000003_.IsInputsNodeExpanded = false
	__Task__00000003_.IsOutputsNodeExpanded = false
	__Task__00000003_.IsWithCompletion = false
	__Task__00000003_.Completion = ""

	__Task__00000004_.Name = `Program Investigation Team (PIT) Report`
	__Task__00000004_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000004_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000004_.Description = ``
	__Task__00000004_.IsExpanded = false
	__Task__00000004_.ComputedPrefix = `1.1`
	__Task__00000004_.IsInRenameMode = false
	__Task__00000004_.IsInputsNodeExpanded = false
	__Task__00000004_.IsOutputsNodeExpanded = false
	__Task__00000004_.IsWithCompletion = false
	__Task__00000004_.Completion = ""

	__TaskCompositionShape__00000000_.Name = `Commercial Crew Program (CCP), to Starliner Crewed Flight Test (CFT)`
	__TaskCompositionShape__00000000_.StartRatio = 0.500000
	__TaskCompositionShape__00000000_.EndRatio = 0.500000
	__TaskCompositionShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000000_.CornerOffsetRatio = 1.680000

	__TaskCompositionShape__00000001_.Name = `Commercial Crew Program (CCP), to Starliner Crewed Flight Test (CFT)`
	__TaskCompositionShape__00000001_.StartRatio = 0.500000
	__TaskCompositionShape__00000001_.EndRatio = 0.500000
	__TaskCompositionShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000001_.CornerOffsetRatio = 1.680000

	__TaskCompositionShape__00000002_.Name = `Mishap investigations to `
	__TaskCompositionShape__00000002_.StartRatio = 0.500000
	__TaskCompositionShape__00000002_.EndRatio = 0.500000
	__TaskCompositionShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000002_.CornerOffsetRatio = 1.680000

	__TaskInputShape__00000000_.Name = `Program Investigation Team (PIT) Report to  Starliner Tests and Anomalies Review (STAR) Investigation Report`
	__TaskInputShape__00000000_.StartRatio = 0.278248
	__TaskInputShape__00000000_.EndRatio = 0.277627
	__TaskInputShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskInputShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskInputShape__00000000_.CornerOffsetRatio = 1.883372

	__TaskOutputShape__00000000_.Name = `Program Investigation Team (PIT) Report to Program Investigation Team (PIT) Report`
	__TaskOutputShape__00000000_.StartRatio = 0.439622
	__TaskOutputShape__00000000_.EndRatio = 0.501995
	__TaskOutputShape__00000000_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000000_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000000_.CornerOffsetRatio = 1.313627

	__TaskShape__00000000_.Name = `NewTask-Default Diagram`
	__TaskShape__00000000_.IsExpanded = false
	__TaskShape__00000000_.X = 61.409529
	__TaskShape__00000000_.Y = 952.589880
	__TaskShape__00000000_.Width = 250.000000
	__TaskShape__00000000_.Height = 70.000000

	__TaskShape__00000001_.Name = `NewTask-Default Diagram`
	__TaskShape__00000001_.IsExpanded = false
	__TaskShape__00000001_.X = 105.187127
	__TaskShape__00000001_.Y = 246.958864
	__TaskShape__00000001_.Width = 250.000000
	__TaskShape__00000001_.Height = 70.000000

	__TaskShape__00000002_.Name = `-Default Diagram`
	__TaskShape__00000002_.IsExpanded = false
	__TaskShape__00000002_.X = 62.073018
	__TaskShape__00000002_.Y = 795.651222
	__TaskShape__00000002_.Width = 250.000000
	__TaskShape__00000002_.Height = 70.000000

	__TaskShape__00000003_.Name = `Mishap investigation-WBS`
	__TaskShape__00000003_.IsExpanded = false
	__TaskShape__00000003_.X = 28.416829
	__TaskShape__00000003_.Y = 546.674065
	__TaskShape__00000003_.Width = 250.000000
	__TaskShape__00000003_.Height = 70.000000

	__TaskShape__00000004_.Name = `Commercial Crew Program (CCP),-WBS`
	__TaskShape__00000004_.IsExpanded = false
	__TaskShape__00000004_.X = 16.071983
	__TaskShape__00000004_.Y = 166.517160
	__TaskShape__00000004_.Width = 250.000000
	__TaskShape__00000004_.Height = 70.000000

	__TaskShape__00000005_.Name = `Starliner Crewed Flight Test (CFT)-WBS`
	__TaskShape__00000005_.IsExpanded = false
	__TaskShape__00000005_.X = 16.617766
	__TaskShape__00000005_.Y = 336.295606
	__TaskShape__00000005_.Width = 250.000000
	__TaskShape__00000005_.Height = 70.000000

	__TaskShape__00000006_.Name = `-WBS`
	__TaskShape__00000006_.IsExpanded = false
	__TaskShape__00000006_.X = 14.409715
	__TaskShape__00000006_.Y = 32.470060
	__TaskShape__00000006_.Width = 250.000000
	__TaskShape__00000006_.Height = 70.000000

	__TaskShape__00000007_.Name = `Mishap investigation-PBS`
	__TaskShape__00000007_.IsExpanded = false
	__TaskShape__00000007_.X = 782.178384
	__TaskShape__00000007_.Y = 665.328500
	__TaskShape__00000007_.Width = 250.000000
	__TaskShape__00000007_.Height = 70.000000

	__TaskShape__00000008_.Name = `-WBS`
	__TaskShape__00000008_.IsExpanded = false
	__TaskShape__00000008_.X = 28.416829
	__TaskShape__00000008_.Y = 686.674065
	__TaskShape__00000008_.Width = 250.000000
	__TaskShape__00000008_.Height = 70.000000

	__TaskShape__00000009_.Name = `Program Investigation Team (PIT) Report-PIT focus`
	__TaskShape__00000009_.IsExpanded = false
	__TaskShape__00000009_.X = 499.174738
	__TaskShape__00000009_.Y = 283.233412
	__TaskShape__00000009_.Width = 250.000000
	__TaskShape__00000009_.Height = 70.000000

	// insertion point for setup of pointers
	__Diagram__00000000_.Task_Shapes = append(__Diagram__00000000_.Task_Shapes, __TaskShape__00000000_)
	__Diagram__00000000_.Task_Shapes = append(__Diagram__00000000_.Task_Shapes, __TaskShape__00000001_)
	__Diagram__00000000_.Task_Shapes = append(__Diagram__00000000_.Task_Shapes, __TaskShape__00000002_)
	__Diagram__00000000_.TasksWhoseNodeIsExpanded = append(__Diagram__00000000_.TasksWhoseNodeIsExpanded, __Task__00000002_)
	__Diagram__00000000_.TaskComposition_Shapes = append(__Diagram__00000000_.TaskComposition_Shapes, __TaskCompositionShape__00000000_)
	__Diagram__00000000_.Note_Shapes = append(__Diagram__00000000_.Note_Shapes, __NoteShape__00000000_)
	__Diagram__00000000_.NoteTaskShapes = append(__Diagram__00000000_.NoteTaskShapes, __NoteTaskShape__00000000_)
	__Diagram__00000000_.Resource_Shapes = append(__Diagram__00000000_.Resource_Shapes, __ResourceShape__00000000_)
	__Diagram__00000000_.Resource_Shapes = append(__Diagram__00000000_.Resource_Shapes, __ResourceShape__00000001_)
	__Diagram__00000000_.Resource_Shapes = append(__Diagram__00000000_.Resource_Shapes, __ResourceShape__00000002_)
	__Diagram__00000000_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000000_.ResourcesWhoseNodeIsExpanded, __Resource__00000000_)
	__Diagram__00000000_.ResourceTaskShapes = append(__Diagram__00000000_.ResourceTaskShapes, __ResourceTaskShape__00000000_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000003_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000004_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000005_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000006_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000008_)
	__Diagram__00000001_.TasksWhoseNodeIsExpanded = append(__Diagram__00000001_.TasksWhoseNodeIsExpanded, __Task__00000002_)
	__Diagram__00000001_.TasksWhoseNodeIsExpanded = append(__Diagram__00000001_.TasksWhoseNodeIsExpanded, __Task__00000001_)
	__Diagram__00000001_.TaskComposition_Shapes = append(__Diagram__00000001_.TaskComposition_Shapes, __TaskCompositionShape__00000001_)
	__Diagram__00000001_.TaskComposition_Shapes = append(__Diagram__00000001_.TaskComposition_Shapes, __TaskCompositionShape__00000002_)
	__Diagram__00000001_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000001_.ResourcesWhoseNodeIsExpanded, __Resource__00000003_)
	__Diagram__00000002_.Product_Shapes = append(__Diagram__00000002_.Product_Shapes, __ProductShape__00000002_)
	__Diagram__00000002_.Product_Shapes = append(__Diagram__00000002_.Product_Shapes, __ProductShape__00000003_)
	__Diagram__00000002_.Product_Shapes = append(__Diagram__00000002_.Product_Shapes, __ProductShape__00000004_)
	__Diagram__00000002_.Product_Shapes = append(__Diagram__00000002_.Product_Shapes, __ProductShape__00000005_)
	__Diagram__00000002_.Product_Shapes = append(__Diagram__00000002_.Product_Shapes, __ProductShape__00000006_)
	__Diagram__00000002_.Product_Shapes = append(__Diagram__00000002_.Product_Shapes, __ProductShape__00000009_)
	__Diagram__00000002_.Product_Shapes = append(__Diagram__00000002_.Product_Shapes, __ProductShape__00000010_)
	__Diagram__00000002_.Product_Shapes = append(__Diagram__00000002_.Product_Shapes, __ProductShape__00000011_)
	__Diagram__00000002_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000002_.ProductsWhoseNodeIsExpanded, __Product__00000003_)
	__Diagram__00000002_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000002_.ProductsWhoseNodeIsExpanded, __Product__00000006_)
	__Diagram__00000002_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000002_.ProductsWhoseNodeIsExpanded, __Product__00000007_)
	__Diagram__00000002_.ProductComposition_Shapes = append(__Diagram__00000002_.ProductComposition_Shapes, __ProductCompositionShape__00000000_)
	__Diagram__00000002_.ProductComposition_Shapes = append(__Diagram__00000002_.ProductComposition_Shapes, __ProductCompositionShape__00000001_)
	__Diagram__00000002_.ProductComposition_Shapes = append(__Diagram__00000002_.ProductComposition_Shapes, __ProductCompositionShape__00000002_)
	__Diagram__00000002_.ProductComposition_Shapes = append(__Diagram__00000002_.ProductComposition_Shapes, __ProductCompositionShape__00000003_)
	__Diagram__00000002_.ProductComposition_Shapes = append(__Diagram__00000002_.ProductComposition_Shapes, __ProductCompositionShape__00000004_)
	__Diagram__00000002_.ProductComposition_Shapes = append(__Diagram__00000002_.ProductComposition_Shapes, __ProductCompositionShape__00000005_)
	__Diagram__00000002_.Task_Shapes = append(__Diagram__00000002_.Task_Shapes, __TaskShape__00000007_)
	__Diagram__00000002_.TasksWhoseNodeIsExpanded = append(__Diagram__00000002_.TasksWhoseNodeIsExpanded, __Task__00000002_)
	__Diagram__00000002_.Resource_Shapes = append(__Diagram__00000002_.Resource_Shapes, __ResourceShape__00000013_)
	__Diagram__00000002_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000002_.ResourcesWhoseNodeIsExpanded, __Resource__00000003_)
	__Diagram__00000002_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000002_.ResourcesWhoseNodeIsExpanded, __Resource__00000000_)
	__Diagram__00000003_.Resource_Shapes = append(__Diagram__00000003_.Resource_Shapes, __ResourceShape__00000003_)
	__Diagram__00000003_.Resource_Shapes = append(__Diagram__00000003_.Resource_Shapes, __ResourceShape__00000004_)
	__Diagram__00000003_.Resource_Shapes = append(__Diagram__00000003_.Resource_Shapes, __ResourceShape__00000005_)
	__Diagram__00000003_.Resource_Shapes = append(__Diagram__00000003_.Resource_Shapes, __ResourceShape__00000006_)
	__Diagram__00000003_.Resource_Shapes = append(__Diagram__00000003_.Resource_Shapes, __ResourceShape__00000007_)
	__Diagram__00000003_.Resource_Shapes = append(__Diagram__00000003_.Resource_Shapes, __ResourceShape__00000008_)
	__Diagram__00000003_.Resource_Shapes = append(__Diagram__00000003_.Resource_Shapes, __ResourceShape__00000009_)
	__Diagram__00000003_.Resource_Shapes = append(__Diagram__00000003_.Resource_Shapes, __ResourceShape__00000010_)
	__Diagram__00000003_.Resource_Shapes = append(__Diagram__00000003_.Resource_Shapes, __ResourceShape__00000011_)
	__Diagram__00000003_.Resource_Shapes = append(__Diagram__00000003_.Resource_Shapes, __ResourceShape__00000012_)
	__Diagram__00000003_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000003_.ResourcesWhoseNodeIsExpanded, __Resource__00000000_)
	__Diagram__00000003_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000003_.ResourcesWhoseNodeIsExpanded, __Resource__00000003_)
	__Diagram__00000003_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000003_.ResourcesWhoseNodeIsExpanded, __Resource__00000004_)
	__Diagram__00000003_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000003_.ResourcesWhoseNodeIsExpanded, __Resource__00000006_)
	__Diagram__00000003_.ResourceComposition_Shapes = append(__Diagram__00000003_.ResourceComposition_Shapes, __ResourceCompositionShape__00000000_)
	__Diagram__00000003_.ResourceComposition_Shapes = append(__Diagram__00000003_.ResourceComposition_Shapes, __ResourceCompositionShape__00000001_)
	__Diagram__00000003_.ResourceComposition_Shapes = append(__Diagram__00000003_.ResourceComposition_Shapes, __ResourceCompositionShape__00000002_)
	__Diagram__00000003_.ResourceComposition_Shapes = append(__Diagram__00000003_.ResourceComposition_Shapes, __ResourceCompositionShape__00000003_)
	__Diagram__00000003_.ResourceComposition_Shapes = append(__Diagram__00000003_.ResourceComposition_Shapes, __ResourceCompositionShape__00000004_)
	__Diagram__00000003_.ResourceComposition_Shapes = append(__Diagram__00000003_.ResourceComposition_Shapes, __ResourceCompositionShape__00000005_)
	__Diagram__00000003_.ResourceComposition_Shapes = append(__Diagram__00000003_.ResourceComposition_Shapes, __ResourceCompositionShape__00000006_)
	__Diagram__00000003_.ResourceComposition_Shapes = append(__Diagram__00000003_.ResourceComposition_Shapes, __ResourceCompositionShape__00000007_)
	__Diagram__00000004_.Product_Shapes = append(__Diagram__00000004_.Product_Shapes, __ProductShape__00000007_)
	__Diagram__00000004_.Product_Shapes = append(__Diagram__00000004_.Product_Shapes, __ProductShape__00000008_)
	__Diagram__00000004_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000004_.ProductsWhoseNodeIsExpanded, __Product__00000003_)
	__Diagram__00000004_.Task_Shapes = append(__Diagram__00000004_.Task_Shapes, __TaskShape__00000009_)
	__Diagram__00000004_.TasksWhoseNodeIsExpanded = append(__Diagram__00000004_.TasksWhoseNodeIsExpanded, __Task__00000001_)
	__Diagram__00000004_.TasksWhoseNodeIsExpanded = append(__Diagram__00000004_.TasksWhoseNodeIsExpanded, __Task__00000004_)
	__Diagram__00000004_.TasksWhoseInputNodeIsExpanded = append(__Diagram__00000004_.TasksWhoseInputNodeIsExpanded, __Task__00000004_)
	__Diagram__00000004_.TasksWhoseOutputNodeIsExpanded = append(__Diagram__00000004_.TasksWhoseOutputNodeIsExpanded, __Task__00000004_)
	__Diagram__00000004_.TaskInputShapes = append(__Diagram__00000004_.TaskInputShapes, __TaskInputShape__00000000_)
	__Diagram__00000004_.TaskOutputShapes = append(__Diagram__00000004_.TaskOutputShapes, __TaskOutputShape__00000000_)
	__Diagram__00000004_.Note_Shapes = append(__Diagram__00000004_.Note_Shapes, __NoteShape__00000001_)
	__Diagram__00000004_.NoteProductShapes = append(__Diagram__00000004_.NoteProductShapes, __NoteProductShape__00000000_)
	__Diagram__00000004_.Resource_Shapes = append(__Diagram__00000004_.Resource_Shapes, __ResourceShape__00000014_)
	__Diagram__00000004_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000004_.ResourcesWhoseNodeIsExpanded, __Resource__00000003_)
	__Diagram__00000004_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000004_.ResourcesWhoseNodeIsExpanded, __Resource__00000000_)
	__Note__00000000_.Tasks = append(__Note__00000000_.Tasks, __Task__00000000_)
	__Note__00000001_.Products = append(__Note__00000001_.Products, __Product__00000004_)
	__NoteProductShape__00000000_.Note = __Note__00000001_
	__NoteProductShape__00000000_.Product = __Product__00000004_
	__NoteShape__00000000_.Note = __Note__00000000_
	__NoteShape__00000001_.Note = __Note__00000001_
	__NoteTaskShape__00000000_.Note = __Note__00000000_
	__NoteTaskShape__00000000_.Task = __Task__00000000_
	__Product__00000003_.SubProducts = append(__Product__00000003_.SubProducts, __Product__00000004_)
	__Product__00000003_.SubProducts = append(__Product__00000003_.SubProducts, __Product__00000005_)
	__Product__00000006_.SubProducts = append(__Product__00000006_.SubProducts, __Product__00000002_)
	__Product__00000006_.SubProducts = append(__Product__00000006_.SubProducts, __Product__00000001_)
	__Product__00000006_.SubProducts = append(__Product__00000006_.SubProducts, __Product__00000008_)
	__Product__00000007_.SubProducts = append(__Product__00000007_.SubProducts, __Product__00000006_)
	__ProductCompositionShape__00000000_.Product = __Product__00000004_
	__ProductCompositionShape__00000001_.Product = __Product__00000005_
	__ProductCompositionShape__00000002_.Product = __Product__00000002_
	__ProductCompositionShape__00000003_.Product = __Product__00000001_
	__ProductCompositionShape__00000004_.Product = __Product__00000006_
	__ProductCompositionShape__00000005_.Product = __Product__00000008_
	__ProductShape__00000002_.Product = __Product__00000001_
	__ProductShape__00000003_.Product = __Product__00000002_
	__ProductShape__00000004_.Product = __Product__00000003_
	__ProductShape__00000005_.Product = __Product__00000004_
	__ProductShape__00000006_.Product = __Product__00000005_
	__ProductShape__00000007_.Product = __Product__00000005_
	__ProductShape__00000008_.Product = __Product__00000004_
	__ProductShape__00000009_.Product = __Product__00000006_
	__ProductShape__00000010_.Product = __Product__00000007_
	__ProductShape__00000011_.Product = __Product__00000008_
	__Project__00000000_.RootProducts = append(__Project__00000000_.RootProducts, __Product__00000007_)
	__Project__00000000_.RootProducts = append(__Project__00000000_.RootProducts, __Product__00000003_)
	__Project__00000000_.RootTasks = append(__Project__00000000_.RootTasks, __Task__00000001_)
	__Project__00000000_.RootTasks = append(__Project__00000000_.RootTasks, __Task__00000002_)
	__Project__00000000_.RootTasks = append(__Project__00000000_.RootTasks, __Task__00000003_)
	__Project__00000000_.RootResources = append(__Project__00000000_.RootResources, __Resource__00000003_)
	__Project__00000000_.RootResources = append(__Project__00000000_.RootResources, __Resource__00000006_)
	__Project__00000000_.Notes = append(__Project__00000000_.Notes, __Note__00000000_)
	__Project__00000000_.Notes = append(__Project__00000000_.Notes, __Note__00000001_)
	__Project__00000000_.Diagrams = append(__Project__00000000_.Diagrams, __Diagram__00000000_)
	__Project__00000000_.Diagrams = append(__Project__00000000_.Diagrams, __Diagram__00000001_)
	__Project__00000000_.Diagrams = append(__Project__00000000_.Diagrams, __Diagram__00000002_)
	__Project__00000000_.Diagrams = append(__Project__00000000_.Diagrams, __Diagram__00000003_)
	__Project__00000000_.Diagrams = append(__Project__00000000_.Diagrams, __Diagram__00000004_)
	__Resource__00000000_.Tasks = append(__Resource__00000000_.Tasks, __Task__00000001_)
	__Resource__00000003_.SubResources = append(__Resource__00000003_.SubResources, __Resource__00000004_)
	__Resource__00000003_.SubResources = append(__Resource__00000003_.SubResources, __Resource__00000000_)
	__Resource__00000003_.SubResources = append(__Resource__00000003_.SubResources, __Resource__00000008_)
	__Resource__00000004_.SubResources = append(__Resource__00000004_.SubResources, __Resource__00000005_)
	__Resource__00000004_.SubResources = append(__Resource__00000004_.SubResources, __Resource__00000009_)
	__Resource__00000005_.SubResources = append(__Resource__00000005_.SubResources, __Resource__00000001_)
	__Resource__00000005_.SubResources = append(__Resource__00000005_.SubResources, __Resource__00000002_)
	__Resource__00000006_.SubResources = append(__Resource__00000006_.SubResources, __Resource__00000007_)
	__ResourceCompositionShape__00000000_.Resource = __Resource__00000004_
	__ResourceCompositionShape__00000001_.Resource = __Resource__00000005_
	__ResourceCompositionShape__00000002_.Resource = __Resource__00000001_
	__ResourceCompositionShape__00000003_.Resource = __Resource__00000002_
	__ResourceCompositionShape__00000004_.Resource = __Resource__00000000_
	__ResourceCompositionShape__00000005_.Resource = __Resource__00000007_
	__ResourceCompositionShape__00000006_.Resource = __Resource__00000008_
	__ResourceCompositionShape__00000007_.Resource = __Resource__00000009_
	__ResourceShape__00000000_.Resource = __Resource__00000000_
	__ResourceShape__00000001_.Resource = __Resource__00000001_
	__ResourceShape__00000002_.Resource = __Resource__00000002_
	__ResourceShape__00000003_.Resource = __Resource__00000003_
	__ResourceShape__00000004_.Resource = __Resource__00000004_
	__ResourceShape__00000005_.Resource = __Resource__00000005_
	__ResourceShape__00000006_.Resource = __Resource__00000001_
	__ResourceShape__00000007_.Resource = __Resource__00000002_
	__ResourceShape__00000008_.Resource = __Resource__00000000_
	__ResourceShape__00000009_.Resource = __Resource__00000006_
	__ResourceShape__00000010_.Resource = __Resource__00000007_
	__ResourceShape__00000011_.Resource = __Resource__00000008_
	__ResourceShape__00000012_.Resource = __Resource__00000009_
	__ResourceShape__00000013_.Resource = __Resource__00000000_
	__ResourceShape__00000014_.Resource = __Resource__00000000_
	__ResourceTaskShape__00000000_.Resource = __Resource__00000000_
	__ResourceTaskShape__00000000_.Task = __Task__00000001_
	__Root__00000000_.Projects = append(__Root__00000000_.Projects, __Project__00000000_)
	__Task__00000001_.SubTasks = append(__Task__00000001_.SubTasks, __Task__00000004_)
	__Task__00000002_.SubTasks = append(__Task__00000002_.SubTasks, __Task__00000000_)
	__Task__00000004_.Inputs = append(__Task__00000004_.Inputs, __Product__00000004_)
	__Task__00000004_.Outputs = append(__Task__00000004_.Outputs, __Product__00000005_)
	__TaskCompositionShape__00000000_.Task = __Task__00000000_
	__TaskCompositionShape__00000001_.Task = __Task__00000000_
	__TaskCompositionShape__00000002_.Task = __Task__00000004_
	__TaskInputShape__00000000_.Product = __Product__00000004_
	__TaskInputShape__00000000_.Task = __Task__00000004_
	__TaskOutputShape__00000000_.Task = __Task__00000004_
	__TaskOutputShape__00000000_.Product = __Product__00000005_
	__TaskShape__00000000_.Task = __Task__00000000_
	__TaskShape__00000001_.Task = __Task__00000001_
	__TaskShape__00000002_.Task = __Task__00000002_
	__TaskShape__00000003_.Task = __Task__00000001_
	__TaskShape__00000004_.Task = __Task__00000002_
	__TaskShape__00000005_.Task = __Task__00000000_
	__TaskShape__00000006_.Task = __Task__00000003_
	__TaskShape__00000007_.Task = __Task__00000001_
	__TaskShape__00000008_.Task = __Task__00000004_
	__TaskShape__00000009_.Task = __Task__00000004_
}
