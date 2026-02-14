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

	__Diagram__00000042_ := (&models.Diagram{Name: `NewDiagram`}).Stage(stage)
	__Diagram__00000057_ := (&models.Diagram{Name: `NewDiagram`}).Stage(stage)
	__Diagram__00000061_ := (&models.Diagram{Name: `Work diagram x`}).Stage(stage)

	__Note__00000000_ := (&models.Note{Name: `This is an example to explain a particular product

This has lot of lines of explaining

another line`}).Stage(stage)
	__Note__00000001_ := (&models.Note{Name: `Second Note`}).Stage(stage)
	__Note__00000002_ := (&models.Note{Name: `The algo uses the stage.reference to computes all new/deleted/modified instances of the stage`}).Stage(stage)
	__Note__00000003_ := (&models.Note{Name: `The probe Notfication Table Allows the verification of commit diff`}).Stage(stage)
	__Note__00000004_ := (&models.Note{Name: `https://github.com/fullstack-lang/gong/issues/998`}).Stage(stage)

	__NoteProductShape__00000001_ := (&models.NoteProductShape{Name: `The probe Notfication Table Allows the verification of commit diff to Probe display the marshalled commit`}).Stage(stage)
	__NoteProductShape__00000002_ := (&models.NoteProductShape{Name: `The probe Notfication Table Allows the verification of commit diff to Probe Notification Table`}).Stage(stage)

	__NoteShape__00000000_ := (&models.NoteShape{Name: `This is an example to explain a particular product

This has lot of lines of explaining

another line-NewDiagram`}).Stage(stage)
	__NoteShape__00000002_ := (&models.NoteShape{Name: `NewNote-NewDiagram`}).Stage(stage)
	__NoteShape__00000003_ := (&models.NoteShape{Name: `NewNote-NewDiagram`}).Stage(stage)
	__NoteShape__00000004_ := (&models.NoteShape{Name: `NewNote-NewDiagram`}).Stage(stage)

	__NoteTaskShape__00000001_ := (&models.NoteTaskShape{Name: `The algo uses the stage.reference to computes all new/deleted/modified instances of the stage to Compute all Instance Diff`}).Stage(stage)
	__NoteTaskShape__00000003_ := (&models.NoteTaskShape{Name: `https://github.com/fullstack-lang/gong/issues/998 to Develop Probe Notifications`}).Stage(stage)

	__Product__00000011_ := (&models.Product{Name: `Product 1`}).Stage(stage)
	__Product__00000013_ := (&models.Product{Name: `Product 1.1`}).Stage(stage)
	__Product__00000014_ := (&models.Product{Name: `Product 1.2`}).Stage(stage)
	__Product__00000015_ := (&models.Product{Name: `Product 2`}).Stage(stage)
	__Product__00000016_ := (&models.Product{Name: `Product 1.2.1`}).Stage(stage)
	__Product__00000017_ := (&models.Product{Name: `Product 2.1`}).Stage(stage)
	__Product__00000018_ := (&models.Product{Name: `Product 2.2`}).Stage(stage)
	__Product__00000021_ := (&models.Product{Name: `String of all New+Stage/ Updated/ Unstage`}).Stage(stage)
	__Product__00000022_ := (&models.Product{Name: `New Commit code`}).Stage(stage)
	__Product__00000025_ := (&models.Product{Name: `Probe Notification Table`}).Stage(stage)
	__Product__00000026_ := (&models.Product{Name: `<instance>. GongDiff(another instance)`}).Stage(stage)
	__Product__00000027_ := (&models.Product{Name: `append log when marshalling`}).Stage(stage)
	__Product__00000028_ := (&models.Product{Name: `New Umarshall code`}).Stage(stage)
	__Product__00000029_ := (&models.Product{Name: `Probe Undo`}).Stage(stage)
	__Product__00000030_ := (&models.Product{Name: `Probe Redo`}).Stage(stage)
	__Product__00000031_ := (&models.Product{Name: `New Probe`}).Stage(stage)
	__Product__00000033_ := (&models.Product{Name: `Reverse Commit log Code`}).Stage(stage)
	__Product__00000034_ := (&models.Product{Name: `Rebase Log Commits`}).Stage(stage)
	__Product__00000035_ := (&models.Product{Name: `Probe display the marshalled commit`}).Stage(stage)
	__Product__00000036_ := (&models.Product{Name: `Display the marshalled reverse commit`}).Stage(stage)
	__Product__00000037_ := (&models.Product{Name: `Instance method for marshalling Unstage`}).Stage(stage)
	__Product__00000038_ := (&models.Product{Name: `NewProduct`}).Stage(stage)
	__Product__00000039_ := (&models.Product{Name: `NewProduct`}).Stage(stage)
	__Product__00000040_ := (&models.Product{Name: `NewProduct`}).Stage(stage)
	__Product__00000041_ := (&models.Product{Name: `NewProduct`}).Stage(stage)
	__Product__00000042_ := (&models.Product{Name: `NewProduct`}).Stage(stage)

	__ProductCompositionShape__00000204_ := (&models.ProductCompositionShape{Name: `Product 1 to Product 1.1`}).Stage(stage)
	__ProductCompositionShape__00000205_ := (&models.ProductCompositionShape{Name: `Product 1.2 to Product 1.2.1`}).Stage(stage)
	__ProductCompositionShape__00000206_ := (&models.ProductCompositionShape{Name: `Product 1 to Product 1.2`}).Stage(stage)
	__ProductCompositionShape__00000207_ := (&models.ProductCompositionShape{Name: `Product 2 to Product 2.1`}).Stage(stage)
	__ProductCompositionShape__00000208_ := (&models.ProductCompositionShape{Name: `Product 2 to Product 2.2`}).Stage(stage)
	__ProductCompositionShape__00000298_ := (&models.ProductCompositionShape{Name: `New Commit code to String of New+Stage/Updated/Unstage`}).Stage(stage)
	__ProductCompositionShape__00000302_ := (&models.ProductCompositionShape{Name: `New Commit code to NewProduct`}).Stage(stage)
	__ProductCompositionShape__00000307_ := (&models.ProductCompositionShape{Name: `New Commit code to <instance>. GongDiff(another instance)`}).Stage(stage)
	__ProductCompositionShape__00000308_ := (&models.ProductCompositionShape{Name: `New Commit code to append log when marshalling`}).Stage(stage)
	__ProductCompositionShape__00000309_ := (&models.ProductCompositionShape{Name: `New Commit code to New Umarshall code`}).Stage(stage)
	__ProductCompositionShape__00000310_ := (&models.ProductCompositionShape{Name: `New Commit code to Instance method for marshalling Unstage`}).Stage(stage)
	__ProductCompositionShape__00000311_ := (&models.ProductCompositionShape{Name: `New Probe to Probe Notification Table`}).Stage(stage)
	__ProductCompositionShape__00000312_ := (&models.ProductCompositionShape{Name: `New Probe to Probe display the marshalled commit`}).Stage(stage)
	__ProductCompositionShape__00000313_ := (&models.ProductCompositionShape{Name: `New Probe to Display the marshalled reverse commit`}).Stage(stage)
	__ProductCompositionShape__00000314_ := (&models.ProductCompositionShape{Name: `New Probe to Probe Undo`}).Stage(stage)
	__ProductCompositionShape__00000315_ := (&models.ProductCompositionShape{Name: `New Probe to Probe Redo`}).Stage(stage)
	__ProductCompositionShape__00000316_ := (&models.ProductCompositionShape{Name: `New Probe to Rebase Log Commits`}).Stage(stage)

	__ProductShape__00000330_ := (&models.ProductShape{Name: `Product 1-NewDiagram`}).Stage(stage)
	__ProductShape__00000331_ := (&models.ProductShape{Name: `Product 1.1-NewDiagram`}).Stage(stage)
	__ProductShape__00000332_ := (&models.ProductShape{Name: `Product 1.2-NewDiagram`}).Stage(stage)
	__ProductShape__00000333_ := (&models.ProductShape{Name: `Product 1.2.1-NewDiagram`}).Stage(stage)
	__ProductShape__00000334_ := (&models.ProductShape{Name: `Product 2-NewDiagram`}).Stage(stage)
	__ProductShape__00000335_ := (&models.ProductShape{Name: `Product 2.1-NewDiagram`}).Stage(stage)
	__ProductShape__00000336_ := (&models.ProductShape{Name: `Product 2.2-NewDiagram`}).Stage(stage)
	__ProductShape__00000466_ := (&models.ProductShape{Name: `String of New+Stage/Updated/Unstage-NewDiagram`}).Stage(stage)
	__ProductShape__00000467_ := (&models.ProductShape{Name: `New Commit code-NewDiagram`}).Stage(stage)
	__ProductShape__00000468_ := (&models.ProductShape{Name: `NewProduct-NewDiagram`}).Stage(stage)
	__ProductShape__00000469_ := (&models.ProductShape{Name: `NewProduct-NewDiagram`}).Stage(stage)
	__ProductShape__00000470_ := (&models.ProductShape{Name: `NewProduct-NewDiagram`}).Stage(stage)
	__ProductShape__00000471_ := (&models.ProductShape{Name: `New Umarshall code-NewDiagram`}).Stage(stage)
	__ProductShape__00000472_ := (&models.ProductShape{Name: `NewProduct-NewDiagram`}).Stage(stage)
	__ProductShape__00000473_ := (&models.ProductShape{Name: `Probe Redo-NewDiagram`}).Stage(stage)
	__ProductShape__00000475_ := (&models.ProductShape{Name: `NewProduct-NewDiagram`}).Stage(stage)
	__ProductShape__00000476_ := (&models.ProductShape{Name: `Rebase Log Commits-NewDiagram`}).Stage(stage)
	__ProductShape__00000477_ := (&models.ProductShape{Name: `Probe display the marshalled commit-Work diagram x`}).Stage(stage)
	__ProductShape__00000478_ := (&models.ProductShape{Name: `Display the marshalled reverse commit-Work diagram x`}).Stage(stage)
	__ProductShape__00000479_ := (&models.ProductShape{Name: `Instance method for marshalling Unstage-Work diagram x`}).Stage(stage)
	__ProductShape__00000480_ := (&models.ProductShape{Name: `NewProduct-NewDiagram`}).Stage(stage)
	__ProductShape__00000481_ := (&models.ProductShape{Name: `New Probe-Work diagram x`}).Stage(stage)

	__Project__00000000_ := (&models.Project{Name: `Project Editor`}).Stage(stage)
	__Project__00000001_ := (&models.Project{Name: `DSME Docx`}).Stage(stage)
	__Project__00000002_ := (&models.Project{Name: `gong : persist stage.go as an history of commits  #934`}).Stage(stage)

	__Resource__00000000_ := (&models.Resource{Name: `T`}).Stage(stage)
	__Resource__00000001_ := (&models.Resource{Name: `NewResource`}).Stage(stage)

	__ResourceShape__00000000_ := (&models.ResourceShape{Name: `NewResource-Work diagram x`}).Stage(stage)
	__ResourceShape__00000001_ := (&models.ResourceShape{Name: `T-Work diagram x`}).Stage(stage)
	__ResourceShape__00000002_ := (&models.ResourceShape{Name: `T-Work diagram x`}).Stage(stage)
	__ResourceShape__00000003_ := (&models.ResourceShape{Name: `T-Work diagram x`}).Stage(stage)
	__ResourceShape__00000004_ := (&models.ResourceShape{Name: `NewResource-Work diagram x`}).Stage(stage)

	__Root__00000000_ := (&models.Root{Name: `Root`}).Stage(stage)

	__Task__00000003_ := (&models.Task{Name: `Task 1`}).Stage(stage)
	__Task__00000007_ := (&models.Task{Name: `Compute all Instance Diff`}).Stage(stage)
	__Task__00000008_ := (&models.Task{Name: `code GongDiff per instance`}).Stage(stage)
	__Task__00000009_ := (&models.Task{Name: `Develop Probe Notifications`}).Stage(stage)
	__Task__00000010_ := (&models.Task{Name: `Display the the commit diff in the probe as marshalled go statements #1038 `}).Stage(stage)
	__Task__00000011_ := (&models.Task{Name: `Put the commit forward text in one notification #1044`}).Stage(stage)
	__Task__00000012_ := (&models.Task{Name: `Compute diff in slices with minimal edit distance #1045`}).Stage(stage)
	__Task__00000013_ := (&models.Task{Name: `refactor AST generation`}).Stage(stage)

	__TaskInputShape__00000001_ := (&models.TaskInputShape{Name: `code GongDiff to Probe Notification Table`}).Stage(stage)
	__TaskInputShape__00000002_ := (&models.TaskInputShape{Name: `Compute all Instance Diff to Probe Notification Table`}).Stage(stage)

	__TaskOutputShape__00000063_ := (&models.TaskOutputShape{Name: `Compute Instance Diff to String of New+Stage/Updated/Unstage`}).Stage(stage)
	__TaskOutputShape__00000064_ := (&models.TaskOutputShape{Name: `Develop Probe Notifications to Probe Notification Table`}).Stage(stage)
	__TaskOutputShape__00000065_ := (&models.TaskOutputShape{Name: `code GongDiff per instance to <instance>. GongDiff(another instance)`}).Stage(stage)
	__TaskOutputShape__00000067_ := (&models.TaskOutputShape{Name: `#1038 Display the the commit diff in the probe as marshalled go statements to Probe display the marshalled commit`}).Stage(stage)
	__TaskOutputShape__00000068_ := (&models.TaskOutputShape{Name: `Put the commit forward text in one notification #1044 to Probe display the marshalled commit`}).Stage(stage)
	__TaskOutputShape__00000069_ := (&models.TaskOutputShape{Name: `Compute diff in slices with minimal edit distance #1045 to Probe display the marshalled commit`}).Stage(stage)

	__TaskShape__00000019_ := (&models.TaskShape{Name: `zzz_gong : compute instance.GongDiff(another_instance)-NewDiagram`}).Stage(stage)
	__TaskShape__00000020_ := (&models.TaskShape{Name: `NewTask-NewDiagram`}).Stage(stage)
	__TaskShape__00000021_ := (&models.TaskShape{Name: `NewTask-NewDiagram`}).Stage(stage)
	__TaskShape__00000022_ := (&models.TaskShape{Name: `Unstage Method Marshalling-Work diagram x`}).Stage(stage)
	__TaskShape__00000023_ := (&models.TaskShape{Name: `NewTask-Work diagram x`}).Stage(stage)
	__TaskShape__00000024_ := (&models.TaskShape{Name: `Compute diff in slices with minimal edit distance #1045-Work diagram x`}).Stage(stage)
	__TaskShape__00000025_ := (&models.TaskShape{Name: `refactor AST generation-Work diagram x`}).Stage(stage)

	// insertion point for initialization of values

	__Diagram__00000042_.Name = `NewDiagram`
	__Diagram__00000042_.IsChecked = false
	__Diagram__00000042_.IsEditable_ = true
	__Diagram__00000042_.IsInRenameMode = false
	__Diagram__00000042_.ShowPrefix = true
	__Diagram__00000042_.DefaultBoxWidth = 150.000000
	__Diagram__00000042_.DefaultBoxHeigth = 70.000000
	__Diagram__00000042_.IsExpanded = false
	__Diagram__00000042_.ComputedPrefix = ``
	__Diagram__00000042_.IsPBSNodeExpanded = true
	__Diagram__00000042_.IsWBSNodeExpanded = true
	__Diagram__00000042_.IsNotesNodeExpanded = false
	__Diagram__00000042_.IsResourcesNodeExpanded = false

	__Diagram__00000057_.Name = `NewDiagram`
	__Diagram__00000057_.IsChecked = false
	__Diagram__00000057_.IsEditable_ = true
	__Diagram__00000057_.IsInRenameMode = false
	__Diagram__00000057_.ShowPrefix = false
	__Diagram__00000057_.DefaultBoxWidth = 150.000000
	__Diagram__00000057_.DefaultBoxHeigth = 70.000000
	__Diagram__00000057_.IsExpanded = false
	__Diagram__00000057_.ComputedPrefix = ``
	__Diagram__00000057_.IsPBSNodeExpanded = true
	__Diagram__00000057_.IsWBSNodeExpanded = false
	__Diagram__00000057_.IsNotesNodeExpanded = false
	__Diagram__00000057_.IsResourcesNodeExpanded = false

	__Diagram__00000061_.Name = `Work diagram x`
	__Diagram__00000061_.IsChecked = true
	__Diagram__00000061_.IsEditable_ = true
	__Diagram__00000061_.IsInRenameMode = false
	__Diagram__00000061_.ShowPrefix = true
	__Diagram__00000061_.DefaultBoxWidth = 250.000000
	__Diagram__00000061_.DefaultBoxHeigth = 100.000000
	__Diagram__00000061_.IsExpanded = true
	__Diagram__00000061_.ComputedPrefix = ``
	__Diagram__00000061_.IsPBSNodeExpanded = true
	__Diagram__00000061_.IsWBSNodeExpanded = true
	__Diagram__00000061_.IsNotesNodeExpanded = true
	__Diagram__00000061_.IsResourcesNodeExpanded = true

	__Note__00000000_.Name = `This is an example to explain a particular product

This has lot of lines of explaining

another line`
	__Note__00000000_.IsExpanded = false
	__Note__00000000_.ComputedPrefix = ``

	__Note__00000001_.Name = `Second Note`
	__Note__00000001_.IsExpanded = false
	__Note__00000001_.ComputedPrefix = ``

	__Note__00000002_.Name = `The algo uses the stage.reference to computes all new/deleted/modified instances of the stage`
	__Note__00000002_.IsExpanded = false
	__Note__00000002_.ComputedPrefix = ``

	__Note__00000003_.Name = `The probe Notfication Table Allows the verification of commit diff`
	__Note__00000003_.IsExpanded = false
	__Note__00000003_.ComputedPrefix = ``

	__Note__00000004_.Name = `https://github.com/fullstack-lang/gong/issues/998`
	__Note__00000004_.IsExpanded = false
	__Note__00000004_.ComputedPrefix = ``

	__NoteProductShape__00000001_.Name = `The probe Notfication Table Allows the verification of commit diff to Probe display the marshalled commit`
	__NoteProductShape__00000001_.StartRatio = 0.500000
	__NoteProductShape__00000001_.EndRatio = 0.500000
	__NoteProductShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000001_.CornerOffsetRatio = 1.680000

	__NoteProductShape__00000002_.Name = `The probe Notfication Table Allows the verification of commit diff to Probe Notification Table`
	__NoteProductShape__00000002_.StartRatio = 0.500000
	__NoteProductShape__00000002_.EndRatio = 0.500000
	__NoteProductShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteProductShape__00000002_.CornerOffsetRatio = 1.680000

	__NoteShape__00000000_.Name = `This is an example to explain a particular product

This has lot of lines of explaining

another line-NewDiagram`
	__NoteShape__00000000_.IsExpanded = false
	__NoteShape__00000000_.X = 421.108494
	__NoteShape__00000000_.Y = 147.407017
	__NoteShape__00000000_.Width = 394.000000
	__NoteShape__00000000_.Height = 87.000000

	__NoteShape__00000002_.Name = `NewNote-NewDiagram`
	__NoteShape__00000002_.IsExpanded = false
	__NoteShape__00000002_.X = 569.404861
	__NoteShape__00000002_.Y = 46.960749
	__NoteShape__00000002_.Width = 250.000000
	__NoteShape__00000002_.Height = 100.000000

	__NoteShape__00000003_.Name = `NewNote-NewDiagram`
	__NoteShape__00000003_.IsExpanded = false
	__NoteShape__00000003_.X = 60.407081
	__NoteShape__00000003_.Y = 1156.370815
	__NoteShape__00000003_.Width = 250.000000
	__NoteShape__00000003_.Height = 85.000000

	__NoteShape__00000004_.Name = `NewNote-NewDiagram`
	__NoteShape__00000004_.IsExpanded = false
	__NoteShape__00000004_.X = 49.079720
	__NoteShape__00000004_.Y = 106.901476
	__NoteShape__00000004_.Width = 411.000000
	__NoteShape__00000004_.Height = 53.000000

	__NoteTaskShape__00000001_.Name = `The algo uses the stage.reference to computes all new/deleted/modified instances of the stage to Compute all Instance Diff`
	__NoteTaskShape__00000001_.StartRatio = 0.500000
	__NoteTaskShape__00000001_.EndRatio = 0.500000
	__NoteTaskShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteTaskShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteTaskShape__00000001_.CornerOffsetRatio = 1.680000

	__NoteTaskShape__00000003_.Name = `https://github.com/fullstack-lang/gong/issues/998 to Develop Probe Notifications`
	__NoteTaskShape__00000003_.StartRatio = 0.500000
	__NoteTaskShape__00000003_.EndRatio = 0.500000
	__NoteTaskShape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__NoteTaskShape__00000003_.EndOrientation = models.ORIENTATION_VERTICAL
	__NoteTaskShape__00000003_.CornerOffsetRatio = 1.680000

	__Product__00000011_.Name = `Product 1`
	__Product__00000011_.Description = ``
	__Product__00000011_.IsExpanded = true
	__Product__00000011_.ComputedPrefix = `1`
	__Product__00000011_.IsProducersNodeExpanded = false
	__Product__00000011_.IsConsumersNodeExpanded = false

	__Product__00000013_.Name = `Product 1.1`
	__Product__00000013_.Description = ``
	__Product__00000013_.IsExpanded = false
	__Product__00000013_.ComputedPrefix = `1.1`
	__Product__00000013_.IsProducersNodeExpanded = false
	__Product__00000013_.IsConsumersNodeExpanded = false

	__Product__00000014_.Name = `Product 1.2`
	__Product__00000014_.Description = ``
	__Product__00000014_.IsExpanded = true
	__Product__00000014_.ComputedPrefix = `1.2`
	__Product__00000014_.IsProducersNodeExpanded = false
	__Product__00000014_.IsConsumersNodeExpanded = false

	__Product__00000015_.Name = `Product 2`
	__Product__00000015_.Description = ``
	__Product__00000015_.IsExpanded = true
	__Product__00000015_.ComputedPrefix = `2`
	__Product__00000015_.IsProducersNodeExpanded = false
	__Product__00000015_.IsConsumersNodeExpanded = false

	__Product__00000016_.Name = `Product 1.2.1`
	__Product__00000016_.Description = ``
	__Product__00000016_.IsExpanded = false
	__Product__00000016_.ComputedPrefix = `1.2.1`
	__Product__00000016_.IsProducersNodeExpanded = false
	__Product__00000016_.IsConsumersNodeExpanded = false

	__Product__00000017_.Name = `Product 2.1`
	__Product__00000017_.Description = ``
	__Product__00000017_.IsExpanded = false
	__Product__00000017_.ComputedPrefix = `2.1`
	__Product__00000017_.IsProducersNodeExpanded = false
	__Product__00000017_.IsConsumersNodeExpanded = false

	__Product__00000018_.Name = `Product 2.2`
	__Product__00000018_.Description = ``
	__Product__00000018_.IsExpanded = false
	__Product__00000018_.ComputedPrefix = `2.2`
	__Product__00000018_.IsProducersNodeExpanded = false
	__Product__00000018_.IsConsumersNodeExpanded = false

	__Product__00000021_.Name = `String of all New+Stage/ Updated/ Unstage`
	__Product__00000021_.Description = ``
	__Product__00000021_.IsExpanded = false
	__Product__00000021_.ComputedPrefix = `1.3`
	__Product__00000021_.IsProducersNodeExpanded = false
	__Product__00000021_.IsConsumersNodeExpanded = false

	__Product__00000022_.Name = `New Commit code`
	__Product__00000022_.Description = ``
	__Product__00000022_.IsExpanded = false
	__Product__00000022_.ComputedPrefix = `1`
	__Product__00000022_.IsProducersNodeExpanded = false
	__Product__00000022_.IsConsumersNodeExpanded = false

	__Product__00000025_.Name = `Probe Notification Table`
	__Product__00000025_.Description = ``
	__Product__00000025_.IsExpanded = false
	__Product__00000025_.ComputedPrefix = `2.1`
	__Product__00000025_.IsProducersNodeExpanded = false
	__Product__00000025_.IsConsumersNodeExpanded = false

	__Product__00000026_.Name = `<instance>. GongDiff(another instance)`
	__Product__00000026_.Description = ``
	__Product__00000026_.IsExpanded = false
	__Product__00000026_.ComputedPrefix = `1.1`
	__Product__00000026_.IsProducersNodeExpanded = false
	__Product__00000026_.IsConsumersNodeExpanded = false

	__Product__00000027_.Name = `append log when marshalling`
	__Product__00000027_.Description = ``
	__Product__00000027_.IsExpanded = false
	__Product__00000027_.ComputedPrefix = `1.5`
	__Product__00000027_.IsProducersNodeExpanded = false
	__Product__00000027_.IsConsumersNodeExpanded = false

	__Product__00000028_.Name = `New Umarshall code`
	__Product__00000028_.Description = ``
	__Product__00000028_.IsExpanded = false
	__Product__00000028_.ComputedPrefix = `1.6`
	__Product__00000028_.IsProducersNodeExpanded = false
	__Product__00000028_.IsConsumersNodeExpanded = false

	__Product__00000029_.Name = `Probe Undo`
	__Product__00000029_.Description = ``
	__Product__00000029_.IsExpanded = false
	__Product__00000029_.ComputedPrefix = `2.4`
	__Product__00000029_.IsProducersNodeExpanded = false
	__Product__00000029_.IsConsumersNodeExpanded = false

	__Product__00000030_.Name = `Probe Redo`
	__Product__00000030_.Description = ``
	__Product__00000030_.IsExpanded = false
	__Product__00000030_.ComputedPrefix = `2.5`
	__Product__00000030_.IsProducersNodeExpanded = false
	__Product__00000030_.IsConsumersNodeExpanded = false

	__Product__00000031_.Name = `New Probe`
	__Product__00000031_.Description = ``
	__Product__00000031_.IsExpanded = true
	__Product__00000031_.ComputedPrefix = `2`
	__Product__00000031_.IsProducersNodeExpanded = false
	__Product__00000031_.IsConsumersNodeExpanded = false

	__Product__00000033_.Name = `Reverse Commit log Code`
	__Product__00000033_.Description = ``
	__Product__00000033_.IsExpanded = false
	__Product__00000033_.ComputedPrefix = `1.4`
	__Product__00000033_.IsProducersNodeExpanded = false
	__Product__00000033_.IsConsumersNodeExpanded = false

	__Product__00000034_.Name = `Rebase Log Commits`
	__Product__00000034_.Description = ``
	__Product__00000034_.IsExpanded = false
	__Product__00000034_.ComputedPrefix = `2.6`
	__Product__00000034_.IsProducersNodeExpanded = false
	__Product__00000034_.IsConsumersNodeExpanded = false

	__Product__00000035_.Name = `Probe display the marshalled commit`
	__Product__00000035_.Description = ``
	__Product__00000035_.IsExpanded = false
	__Product__00000035_.ComputedPrefix = `2.2`
	__Product__00000035_.IsProducersNodeExpanded = false
	__Product__00000035_.IsConsumersNodeExpanded = false

	__Product__00000036_.Name = `Display the marshalled reverse commit`
	__Product__00000036_.Description = ``
	__Product__00000036_.IsExpanded = false
	__Product__00000036_.ComputedPrefix = `2.3`
	__Product__00000036_.IsProducersNodeExpanded = false
	__Product__00000036_.IsConsumersNodeExpanded = false

	__Product__00000037_.Name = `Instance method for marshalling Unstage`
	__Product__00000037_.Description = ``
	__Product__00000037_.IsExpanded = false
	__Product__00000037_.ComputedPrefix = `1.2`
	__Product__00000037_.IsProducersNodeExpanded = false
	__Product__00000037_.IsConsumersNodeExpanded = false

	__Product__00000038_.Name = `NewProduct`
	__Product__00000038_.Description = ``
	__Product__00000038_.IsExpanded = true
	__Product__00000038_.ComputedPrefix = `1`
	__Product__00000038_.IsProducersNodeExpanded = false
	__Product__00000038_.IsConsumersNodeExpanded = false

	__Product__00000039_.Name = `NewProduct`
	__Product__00000039_.Description = ``
	__Product__00000039_.IsExpanded = false
	__Product__00000039_.ComputedPrefix = `1.1`
	__Product__00000039_.IsProducersNodeExpanded = false
	__Product__00000039_.IsConsumersNodeExpanded = false

	__Product__00000040_.Name = `NewProduct`
	__Product__00000040_.Description = ``
	__Product__00000040_.IsExpanded = false
	__Product__00000040_.ComputedPrefix = `1.2`
	__Product__00000040_.IsProducersNodeExpanded = false
	__Product__00000040_.IsConsumersNodeExpanded = false

	__Product__00000041_.Name = `NewProduct`
	__Product__00000041_.Description = ``
	__Product__00000041_.IsExpanded = false
	__Product__00000041_.ComputedPrefix = `1.3`
	__Product__00000041_.IsProducersNodeExpanded = false
	__Product__00000041_.IsConsumersNodeExpanded = false

	__Product__00000042_.Name = `NewProduct`
	__Product__00000042_.Description = ``
	__Product__00000042_.IsExpanded = false
	__Product__00000042_.ComputedPrefix = `2`
	__Product__00000042_.IsProducersNodeExpanded = false
	__Product__00000042_.IsConsumersNodeExpanded = false

	__ProductCompositionShape__00000204_.Name = `Product 1 to Product 1.1`
	__ProductCompositionShape__00000204_.StartRatio = 0.500000
	__ProductCompositionShape__00000204_.EndRatio = 0.500000
	__ProductCompositionShape__00000204_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000204_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000204_.CornerOffsetRatio = 1.500000

	__ProductCompositionShape__00000205_.Name = `Product 1.2 to Product 1.2.1`
	__ProductCompositionShape__00000205_.StartRatio = 0.500000
	__ProductCompositionShape__00000205_.EndRatio = 0.500000
	__ProductCompositionShape__00000205_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000205_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000205_.CornerOffsetRatio = 1.500000

	__ProductCompositionShape__00000206_.Name = `Product 1 to Product 1.2`
	__ProductCompositionShape__00000206_.StartRatio = 0.500000
	__ProductCompositionShape__00000206_.EndRatio = 0.500000
	__ProductCompositionShape__00000206_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000206_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000206_.CornerOffsetRatio = 1.500000

	__ProductCompositionShape__00000207_.Name = `Product 2 to Product 2.1`
	__ProductCompositionShape__00000207_.StartRatio = 0.500000
	__ProductCompositionShape__00000207_.EndRatio = 0.500000
	__ProductCompositionShape__00000207_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000207_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000207_.CornerOffsetRatio = 1.500000

	__ProductCompositionShape__00000208_.Name = `Product 2 to Product 2.2`
	__ProductCompositionShape__00000208_.StartRatio = 0.500000
	__ProductCompositionShape__00000208_.EndRatio = 0.500000
	__ProductCompositionShape__00000208_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000208_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000208_.CornerOffsetRatio = 1.500000

	__ProductCompositionShape__00000298_.Name = `New Commit code to String of New+Stage/Updated/Unstage`
	__ProductCompositionShape__00000298_.StartRatio = 0.500000
	__ProductCompositionShape__00000298_.EndRatio = 0.500000
	__ProductCompositionShape__00000298_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000298_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000298_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000302_.Name = `New Commit code to NewProduct`
	__ProductCompositionShape__00000302_.StartRatio = 0.500000
	__ProductCompositionShape__00000302_.EndRatio = 0.500000
	__ProductCompositionShape__00000302_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000302_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000302_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000307_.Name = `New Commit code to <instance>. GongDiff(another instance)`
	__ProductCompositionShape__00000307_.StartRatio = 0.500000
	__ProductCompositionShape__00000307_.EndRatio = 0.500000
	__ProductCompositionShape__00000307_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000307_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000307_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000308_.Name = `New Commit code to append log when marshalling`
	__ProductCompositionShape__00000308_.StartRatio = 0.500000
	__ProductCompositionShape__00000308_.EndRatio = 0.500000
	__ProductCompositionShape__00000308_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000308_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000308_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000309_.Name = `New Commit code to New Umarshall code`
	__ProductCompositionShape__00000309_.StartRatio = 0.500000
	__ProductCompositionShape__00000309_.EndRatio = 0.500000
	__ProductCompositionShape__00000309_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000309_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000309_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000310_.Name = `New Commit code to Instance method for marshalling Unstage`
	__ProductCompositionShape__00000310_.StartRatio = 0.500000
	__ProductCompositionShape__00000310_.EndRatio = 0.500000
	__ProductCompositionShape__00000310_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000310_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000310_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000311_.Name = `New Probe to Probe Notification Table`
	__ProductCompositionShape__00000311_.StartRatio = 0.500000
	__ProductCompositionShape__00000311_.EndRatio = 0.500000
	__ProductCompositionShape__00000311_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000311_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000311_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000312_.Name = `New Probe to Probe display the marshalled commit`
	__ProductCompositionShape__00000312_.StartRatio = 0.500000
	__ProductCompositionShape__00000312_.EndRatio = 0.500000
	__ProductCompositionShape__00000312_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000312_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000312_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000313_.Name = `New Probe to Display the marshalled reverse commit`
	__ProductCompositionShape__00000313_.StartRatio = 0.500000
	__ProductCompositionShape__00000313_.EndRatio = 0.500000
	__ProductCompositionShape__00000313_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000313_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000313_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000314_.Name = `New Probe to Probe Undo`
	__ProductCompositionShape__00000314_.StartRatio = 0.500000
	__ProductCompositionShape__00000314_.EndRatio = 0.500000
	__ProductCompositionShape__00000314_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000314_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000314_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000315_.Name = `New Probe to Probe Redo`
	__ProductCompositionShape__00000315_.StartRatio = 0.500000
	__ProductCompositionShape__00000315_.EndRatio = 0.500000
	__ProductCompositionShape__00000315_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000315_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000315_.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000316_.Name = `New Probe to Rebase Log Commits`
	__ProductCompositionShape__00000316_.StartRatio = 0.500000
	__ProductCompositionShape__00000316_.EndRatio = 0.500000
	__ProductCompositionShape__00000316_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000316_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000316_.CornerOffsetRatio = 1.680000

	__ProductShape__00000330_.Name = `Product 1-NewDiagram`
	__ProductShape__00000330_.IsExpanded = false
	__ProductShape__00000330_.X = 122.500000
	__ProductShape__00000330_.Y = 10.000000
	__ProductShape__00000330_.Width = 150.000000
	__ProductShape__00000330_.Height = 70.000000

	__ProductShape__00000331_.Name = `Product 1.1-NewDiagram`
	__ProductShape__00000331_.IsExpanded = false
	__ProductShape__00000331_.X = 10.000000
	__ProductShape__00000331_.Y = 150.000000
	__ProductShape__00000331_.Width = 150.000000
	__ProductShape__00000331_.Height = 70.000000

	__ProductShape__00000332_.Name = `Product 1.2-NewDiagram`
	__ProductShape__00000332_.IsExpanded = false
	__ProductShape__00000332_.X = 235.000000
	__ProductShape__00000332_.Y = 150.000000
	__ProductShape__00000332_.Width = 150.000000
	__ProductShape__00000332_.Height = 70.000000

	__ProductShape__00000333_.Name = `Product 1.2.1-NewDiagram`
	__ProductShape__00000333_.IsExpanded = false
	__ProductShape__00000333_.X = 235.000000
	__ProductShape__00000333_.Y = 290.000000
	__ProductShape__00000333_.Width = 150.000000
	__ProductShape__00000333_.Height = 70.000000

	__ProductShape__00000334_.Name = `Product 2-NewDiagram`
	__ProductShape__00000334_.IsExpanded = false
	__ProductShape__00000334_.X = 572.500000
	__ProductShape__00000334_.Y = 10.000000
	__ProductShape__00000334_.Width = 150.000000
	__ProductShape__00000334_.Height = 70.000000

	__ProductShape__00000335_.Name = `Product 2.1-NewDiagram`
	__ProductShape__00000335_.IsExpanded = false
	__ProductShape__00000335_.X = 460.000000
	__ProductShape__00000335_.Y = 150.000000
	__ProductShape__00000335_.Width = 150.000000
	__ProductShape__00000335_.Height = 70.000000

	__ProductShape__00000336_.Name = `Product 2.2-NewDiagram`
	__ProductShape__00000336_.IsExpanded = false
	__ProductShape__00000336_.X = 685.000000
	__ProductShape__00000336_.Y = 150.000000
	__ProductShape__00000336_.Width = 150.000000
	__ProductShape__00000336_.Height = 70.000000

	__ProductShape__00000466_.Name = `String of New+Stage/Updated/Unstage-NewDiagram`
	__ProductShape__00000466_.IsExpanded = false
	__ProductShape__00000466_.X = 850.183991
	__ProductShape__00000466_.Y = 599.413298
	__ProductShape__00000466_.Width = 176.000000
	__ProductShape__00000466_.Height = 78.000000

	__ProductShape__00000467_.Name = `New Commit code-NewDiagram`
	__ProductShape__00000467_.IsExpanded = false
	__ProductShape__00000467_.X = 1189.350565
	__ProductShape__00000467_.Y = 386.227692
	__ProductShape__00000467_.Width = 154.000000
	__ProductShape__00000467_.Height = 65.000000

	__ProductShape__00000468_.Name = `NewProduct-NewDiagram`
	__ProductShape__00000468_.IsExpanded = false
	__ProductShape__00000468_.X = 55.066958
	__ProductShape__00000468_.Y = 955.727459
	__ProductShape__00000468_.Width = 202.000000
	__ProductShape__00000468_.Height = 52.000000

	__ProductShape__00000469_.Name = `NewProduct-NewDiagram`
	__ProductShape__00000469_.IsExpanded = false
	__ProductShape__00000469_.X = 443.358054
	__ProductShape__00000469_.Y = 590.125900
	__ProductShape__00000469_.Width = 155.000000
	__ProductShape__00000469_.Height = 79.000000

	__ProductShape__00000470_.Name = `NewProduct-NewDiagram`
	__ProductShape__00000470_.IsExpanded = false
	__ProductShape__00000470_.X = 1243.329023
	__ProductShape__00000470_.Y = 602.974873
	__ProductShape__00000470_.Width = 154.000000
	__ProductShape__00000470_.Height = 75.000000

	__ProductShape__00000471_.Name = `New Umarshall code-NewDiagram`
	__ProductShape__00000471_.IsExpanded = false
	__ProductShape__00000471_.X = 1450.716678
	__ProductShape__00000471_.Y = 603.610369
	__ProductShape__00000471_.Width = 134.000000
	__ProductShape__00000471_.Height = 72.000000

	__ProductShape__00000472_.Name = `NewProduct-NewDiagram`
	__ProductShape__00000472_.IsExpanded = false
	__ProductShape__00000472_.X = 950.696003
	__ProductShape__00000472_.Y = 962.377332
	__ProductShape__00000472_.Width = 122.000000
	__ProductShape__00000472_.Height = 55.000000

	__ProductShape__00000473_.Name = `Probe Redo-NewDiagram`
	__ProductShape__00000473_.IsExpanded = false
	__ProductShape__00000473_.X = 1167.554776
	__ProductShape__00000473_.Y = 959.390410
	__ProductShape__00000473_.Width = 116.000000
	__ProductShape__00000473_.Height = 63.000000

	__ProductShape__00000475_.Name = `NewProduct-NewDiagram`
	__ProductShape__00000475_.IsExpanded = false
	__ProductShape__00000475_.X = 1069.724901
	__ProductShape__00000475_.Y = 607.341831
	__ProductShape__00000475_.Width = 141.000000
	__ProductShape__00000475_.Height = 66.000000

	__ProductShape__00000476_.Name = `Rebase Log Commits-NewDiagram`
	__ProductShape__00000476_.IsExpanded = false
	__ProductShape__00000476_.X = 1371.656434
	__ProductShape__00000476_.Y = 962.230597
	__ProductShape__00000476_.Width = 126.000000
	__ProductShape__00000476_.Height = 65.000000

	__ProductShape__00000477_.Name = `Probe display the marshalled commit-Work diagram x`
	__ProductShape__00000477_.IsExpanded = false
	__ProductShape__00000477_.X = 305.092048
	__ProductShape__00000477_.Y = 958.996299
	__ProductShape__00000477_.Width = 250.000000
	__ProductShape__00000477_.Height = 70.000000

	__ProductShape__00000478_.Name = `Display the marshalled reverse commit-Work diagram x`
	__ProductShape__00000478_.IsExpanded = false
	__ProductShape__00000478_.X = 665.717385
	__ProductShape__00000478_.Y = 963.308469
	__ProductShape__00000478_.Width = 250.000000
	__ProductShape__00000478_.Height = 64.000000

	__ProductShape__00000479_.Name = `Instance method for marshalling Unstage-Work diagram x`
	__ProductShape__00000479_.IsExpanded = false
	__ProductShape__00000479_.X = 643.481878
	__ProductShape__00000479_.Y = 600.302589
	__ProductShape__00000479_.Width = 187.000000
	__ProductShape__00000479_.Height = 75.000000

	__ProductShape__00000480_.Name = `NewProduct-NewDiagram`
	__ProductShape__00000480_.IsExpanded = false
	__ProductShape__00000480_.X = 168.035150
	__ProductShape__00000480_.Y = 129.647586
	__ProductShape__00000480_.Width = 150.000000
	__ProductShape__00000480_.Height = 70.000000

	__ProductShape__00000481_.Name = `New Probe-Work diagram x`
	__ProductShape__00000481_.IsExpanded = false
	__ProductShape__00000481_.X = 236.035692
	__ProductShape__00000481_.Y = 815.987810
	__ProductShape__00000481_.Width = 250.000000
	__ProductShape__00000481_.Height = 57.000000

	__Project__00000000_.Name = `Project Editor`
	__Project__00000000_.IsExpanded = false
	__Project__00000000_.ComputedPrefix = ``

	__Project__00000001_.Name = `DSME Docx`
	__Project__00000001_.IsExpanded = false
	__Project__00000001_.ComputedPrefix = ``

	__Project__00000002_.Name = `gong : persist stage.go as an history of commits  #934`
	__Project__00000002_.IsExpanded = true
	__Project__00000002_.ComputedPrefix = ``

	__Resource__00000000_.Name = `T`
	__Resource__00000000_.Description = ``
	__Resource__00000000_.IsExpanded = false
	__Resource__00000000_.ComputedPrefix = ``

	__Resource__00000001_.Name = `NewResource`
	__Resource__00000001_.Description = ``
	__Resource__00000001_.IsExpanded = false
	__Resource__00000001_.ComputedPrefix = ``

	__ResourceShape__00000000_.Name = `NewResource-Work diagram x`
	__ResourceShape__00000000_.IsExpanded = false
	__ResourceShape__00000000_.X = 197.905888
	__ResourceShape__00000000_.Y = 164.497153
	__ResourceShape__00000000_.Width = 250.000000
	__ResourceShape__00000000_.Height = 100.000000

	__ResourceShape__00000001_.Name = `T-Work diagram x`
	__ResourceShape__00000001_.IsExpanded = false
	__ResourceShape__00000001_.X = 186.392291
	__ResourceShape__00000001_.Y = 108.018591
	__ResourceShape__00000001_.Width = 250.000000
	__ResourceShape__00000001_.Height = 100.000000

	__ResourceShape__00000002_.Name = `T-Work diagram x`
	__ResourceShape__00000002_.IsExpanded = false
	__ResourceShape__00000002_.X = 126.560696
	__ResourceShape__00000002_.Y = 182.001251
	__ResourceShape__00000002_.Width = 250.000000
	__ResourceShape__00000002_.Height = 100.000000

	__ResourceShape__00000003_.Name = `T-Work diagram x`
	__ResourceShape__00000003_.IsExpanded = false
	__ResourceShape__00000003_.X = 118.393086
	__ResourceShape__00000003_.Y = 141.518392
	__ResourceShape__00000003_.Width = 250.000000
	__ResourceShape__00000003_.Height = 100.000000

	__ResourceShape__00000004_.Name = `NewResource-Work diagram x`
	__ResourceShape__00000004_.IsExpanded = false
	__ResourceShape__00000004_.X = 157.220401
	__ResourceShape__00000004_.Y = 183.322406
	__ResourceShape__00000004_.Width = 250.000000
	__ResourceShape__00000004_.Height = 100.000000

	__Root__00000000_.Name = `Root`
	__Root__00000000_.NbPixPerCharacter = 8.000000

	__Task__00000003_.Name = `Task 1`
	__Task__00000003_.Description = ``
	__Task__00000003_.IsExpanded = true
	__Task__00000003_.ComputedPrefix = `1`
	__Task__00000003_.IsInputsNodeExpanded = true
	__Task__00000003_.IsOutputsNodeExpanded = true
	__Task__00000003_.IsWithCompletion = false
	__Task__00000003_.Completion = ""

	__Task__00000007_.Name = `Compute all Instance Diff`
	__Task__00000007_.Description = ``
	__Task__00000007_.IsExpanded = false
	__Task__00000007_.ComputedPrefix = `2`
	__Task__00000007_.IsInputsNodeExpanded = false
	__Task__00000007_.IsOutputsNodeExpanded = false
	__Task__00000007_.IsWithCompletion = true
	__Task__00000007_.Completion = models.PERCENT_100

	__Task__00000008_.Name = `code GongDiff per instance`
	__Task__00000008_.Description = ``
	__Task__00000008_.IsExpanded = false
	__Task__00000008_.ComputedPrefix = `1`
	__Task__00000008_.IsInputsNodeExpanded = false
	__Task__00000008_.IsOutputsNodeExpanded = false
	__Task__00000008_.IsWithCompletion = true
	__Task__00000008_.Completion = models.PERCENT_100

	__Task__00000009_.Name = `Develop Probe Notifications`
	__Task__00000009_.Description = ``
	__Task__00000009_.IsExpanded = false
	__Task__00000009_.ComputedPrefix = `6`
	__Task__00000009_.IsInputsNodeExpanded = false
	__Task__00000009_.IsOutputsNodeExpanded = false
	__Task__00000009_.IsWithCompletion = true
	__Task__00000009_.Completion = models.PERCENT_100

	__Task__00000010_.Name = `Display the the commit diff in the probe as marshalled go statements #1038 `
	__Task__00000010_.Description = ``
	__Task__00000010_.IsExpanded = false
	__Task__00000010_.ComputedPrefix = `3`
	__Task__00000010_.IsInputsNodeExpanded = false
	__Task__00000010_.IsOutputsNodeExpanded = false
	__Task__00000010_.IsWithCompletion = true
	__Task__00000010_.Completion = models.PERCENT_100

	__Task__00000011_.Name = `Put the commit forward text in one notification #1044`
	__Task__00000011_.Description = ``
	__Task__00000011_.IsExpanded = false
	__Task__00000011_.ComputedPrefix = `4`
	__Task__00000011_.IsInputsNodeExpanded = false
	__Task__00000011_.IsOutputsNodeExpanded = false
	__Task__00000011_.IsWithCompletion = true
	__Task__00000011_.Completion = models.PERCENT_100

	__Task__00000012_.Name = `Compute diff in slices with minimal edit distance #1045`
	__Task__00000012_.Description = ``
	__Task__00000012_.IsExpanded = false
	__Task__00000012_.ComputedPrefix = `5`
	__Task__00000012_.IsInputsNodeExpanded = false
	__Task__00000012_.IsOutputsNodeExpanded = false
	__Task__00000012_.IsWithCompletion = true
	__Task__00000012_.Completion = models.PERCENT_100

	__Task__00000013_.Name = `refactor AST generation`
	__Task__00000013_.Description = ``
	__Task__00000013_.IsExpanded = false
	__Task__00000013_.ComputedPrefix = `7`
	__Task__00000013_.IsInputsNodeExpanded = false
	__Task__00000013_.IsOutputsNodeExpanded = false
	__Task__00000013_.IsWithCompletion = true
	__Task__00000013_.Completion = models.PERCENT_100

	__TaskInputShape__00000001_.Name = `code GongDiff to Probe Notification Table`
	__TaskInputShape__00000001_.StartRatio = 0.664301
	__TaskInputShape__00000001_.EndRatio = 0.500000
	__TaskInputShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskInputShape__00000001_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskInputShape__00000001_.CornerOffsetRatio = -0.725415

	__TaskInputShape__00000002_.Name = `Compute all Instance Diff to Probe Notification Table`
	__TaskInputShape__00000002_.StartRatio = 0.392024
	__TaskInputShape__00000002_.EndRatio = 0.440078
	__TaskInputShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskInputShape__00000002_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskInputShape__00000002_.CornerOffsetRatio = -0.552338

	__TaskOutputShape__00000063_.Name = `Compute Instance Diff to String of New+Stage/Updated/Unstage`
	__TaskOutputShape__00000063_.StartRatio = 0.500000
	__TaskOutputShape__00000063_.EndRatio = 0.000000
	__TaskOutputShape__00000063_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000063_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000063_.CornerOffsetRatio = 1.268543

	__TaskOutputShape__00000064_.Name = `Develop Probe Notifications to Probe Notification Table`
	__TaskOutputShape__00000064_.StartRatio = 0.266561
	__TaskOutputShape__00000064_.EndRatio = 0.194004
	__TaskOutputShape__00000064_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskOutputShape__00000064_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskOutputShape__00000064_.CornerOffsetRatio = 2.418918

	__TaskOutputShape__00000065_.Name = `code GongDiff per instance to <instance>. GongDiff(another instance)`
	__TaskOutputShape__00000065_.StartRatio = 0.414548
	__TaskOutputShape__00000065_.EndRatio = 0.296115
	__TaskOutputShape__00000065_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000065_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskOutputShape__00000065_.CornerOffsetRatio = 1.000288

	__TaskOutputShape__00000067_.Name = `#1038 Display the the commit diff in the probe as marshalled go statements to Probe display the marshalled commit`
	__TaskOutputShape__00000067_.StartRatio = 0.386529
	__TaskOutputShape__00000067_.EndRatio = 0.691171
	__TaskOutputShape__00000067_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskOutputShape__00000067_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskOutputShape__00000067_.CornerOffsetRatio = -0.274121

	__TaskOutputShape__00000068_.Name = `Put the commit forward text in one notification #1044 to Probe display the marshalled commit`
	__TaskOutputShape__00000068_.StartRatio = 0.320898
	__TaskOutputShape__00000068_.EndRatio = 0.856655
	__TaskOutputShape__00000068_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskOutputShape__00000068_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskOutputShape__00000068_.CornerOffsetRatio = -1.026162

	__TaskOutputShape__00000069_.Name = `Compute diff in slices with minimal edit distance #1045 to Probe display the marshalled commit`
	__TaskOutputShape__00000069_.StartRatio = 0.133143
	__TaskOutputShape__00000069_.EndRatio = 0.988655
	__TaskOutputShape__00000069_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskOutputShape__00000069_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskOutputShape__00000069_.CornerOffsetRatio = -1.387283

	__TaskShape__00000019_.Name = `zzz_gong : compute instance.GongDiff(another_instance)-NewDiagram`
	__TaskShape__00000019_.IsExpanded = false
	__TaskShape__00000019_.X = 790.000000
	__TaskShape__00000019_.Y = 280.000000
	__TaskShape__00000019_.Width = 178.000000
	__TaskShape__00000019_.Height = 75.000000

	__TaskShape__00000020_.Name = `NewTask-NewDiagram`
	__TaskShape__00000020_.IsExpanded = false
	__TaskShape__00000020_.X = 215.214318
	__TaskShape__00000020_.Y = 390.572924
	__TaskShape__00000020_.Width = 195.000000
	__TaskShape__00000020_.Height = 72.000000

	__TaskShape__00000021_.Name = `NewTask-NewDiagram`
	__TaskShape__00000021_.IsExpanded = false
	__TaskShape__00000021_.X = 12.662171
	__TaskShape__00000021_.Y = 203.681593
	__TaskShape__00000021_.Width = 217.000000
	__TaskShape__00000021_.Height = 82.000000

	__TaskShape__00000022_.Name = `Unstage Method Marshalling-Work diagram x`
	__TaskShape__00000022_.IsExpanded = false
	__TaskShape__00000022_.X = 389.623689
	__TaskShape__00000022_.Y = 1187.641570
	__TaskShape__00000022_.Width = 293.000000
	__TaskShape__00000022_.Height = 130.000000

	__TaskShape__00000023_.Name = `NewTask-Work diagram x`
	__TaskShape__00000023_.IsExpanded = false
	__TaskShape__00000023_.X = 733.031262
	__TaskShape__00000023_.Y = 1229.997063
	__TaskShape__00000023_.Width = 250.000000
	__TaskShape__00000023_.Height = 100.000000

	__TaskShape__00000024_.Name = `Compute diff in slices with minimal edit distance #1045-Work diagram x`
	__TaskShape__00000024_.IsExpanded = false
	__TaskShape__00000024_.X = 1025.970163
	__TaskShape__00000024_.Y = 1230.109182
	__TaskShape__00000024_.Width = 250.000000
	__TaskShape__00000024_.Height = 100.000000

	__TaskShape__00000025_.Name = `refactor AST generation-Work diagram x`
	__TaskShape__00000025_.IsExpanded = false
	__TaskShape__00000025_.X = 594.989280
	__TaskShape__00000025_.Y = 753.694736
	__TaskShape__00000025_.Width = 250.000000
	__TaskShape__00000025_.Height = 100.000000

	// insertion point for setup of pointers
	__Diagram__00000042_.Product_Shapes = append(__Diagram__00000042_.Product_Shapes, __ProductShape__00000330_)
	__Diagram__00000042_.Product_Shapes = append(__Diagram__00000042_.Product_Shapes, __ProductShape__00000331_)
	__Diagram__00000042_.Product_Shapes = append(__Diagram__00000042_.Product_Shapes, __ProductShape__00000332_)
	__Diagram__00000042_.Product_Shapes = append(__Diagram__00000042_.Product_Shapes, __ProductShape__00000333_)
	__Diagram__00000042_.Product_Shapes = append(__Diagram__00000042_.Product_Shapes, __ProductShape__00000334_)
	__Diagram__00000042_.Product_Shapes = append(__Diagram__00000042_.Product_Shapes, __ProductShape__00000335_)
	__Diagram__00000042_.Product_Shapes = append(__Diagram__00000042_.Product_Shapes, __ProductShape__00000336_)
	__Diagram__00000042_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000042_.ProductsWhoseNodeIsExpanded, __Product__00000015_)
	__Diagram__00000042_.ProductComposition_Shapes = append(__Diagram__00000042_.ProductComposition_Shapes, __ProductCompositionShape__00000204_)
	__Diagram__00000042_.ProductComposition_Shapes = append(__Diagram__00000042_.ProductComposition_Shapes, __ProductCompositionShape__00000205_)
	__Diagram__00000042_.ProductComposition_Shapes = append(__Diagram__00000042_.ProductComposition_Shapes, __ProductCompositionShape__00000206_)
	__Diagram__00000042_.ProductComposition_Shapes = append(__Diagram__00000042_.ProductComposition_Shapes, __ProductCompositionShape__00000207_)
	__Diagram__00000042_.ProductComposition_Shapes = append(__Diagram__00000042_.ProductComposition_Shapes, __ProductCompositionShape__00000208_)
	__Diagram__00000057_.Product_Shapes = append(__Diagram__00000057_.Product_Shapes, __ProductShape__00000480_)
	__Diagram__00000057_.Note_Shapes = append(__Diagram__00000057_.Note_Shapes, __NoteShape__00000000_)
	__Diagram__00000061_.Product_Shapes = append(__Diagram__00000061_.Product_Shapes, __ProductShape__00000466_)
	__Diagram__00000061_.Product_Shapes = append(__Diagram__00000061_.Product_Shapes, __ProductShape__00000467_)
	__Diagram__00000061_.Product_Shapes = append(__Diagram__00000061_.Product_Shapes, __ProductShape__00000468_)
	__Diagram__00000061_.Product_Shapes = append(__Diagram__00000061_.Product_Shapes, __ProductShape__00000469_)
	__Diagram__00000061_.Product_Shapes = append(__Diagram__00000061_.Product_Shapes, __ProductShape__00000470_)
	__Diagram__00000061_.Product_Shapes = append(__Diagram__00000061_.Product_Shapes, __ProductShape__00000471_)
	__Diagram__00000061_.Product_Shapes = append(__Diagram__00000061_.Product_Shapes, __ProductShape__00000472_)
	__Diagram__00000061_.Product_Shapes = append(__Diagram__00000061_.Product_Shapes, __ProductShape__00000473_)
	__Diagram__00000061_.Product_Shapes = append(__Diagram__00000061_.Product_Shapes, __ProductShape__00000475_)
	__Diagram__00000061_.Product_Shapes = append(__Diagram__00000061_.Product_Shapes, __ProductShape__00000476_)
	__Diagram__00000061_.Product_Shapes = append(__Diagram__00000061_.Product_Shapes, __ProductShape__00000477_)
	__Diagram__00000061_.Product_Shapes = append(__Diagram__00000061_.Product_Shapes, __ProductShape__00000478_)
	__Diagram__00000061_.Product_Shapes = append(__Diagram__00000061_.Product_Shapes, __ProductShape__00000479_)
	__Diagram__00000061_.Product_Shapes = append(__Diagram__00000061_.Product_Shapes, __ProductShape__00000481_)
	__Diagram__00000061_.ProductComposition_Shapes = append(__Diagram__00000061_.ProductComposition_Shapes, __ProductCompositionShape__00000298_)
	__Diagram__00000061_.ProductComposition_Shapes = append(__Diagram__00000061_.ProductComposition_Shapes, __ProductCompositionShape__00000302_)
	__Diagram__00000061_.ProductComposition_Shapes = append(__Diagram__00000061_.ProductComposition_Shapes, __ProductCompositionShape__00000307_)
	__Diagram__00000061_.ProductComposition_Shapes = append(__Diagram__00000061_.ProductComposition_Shapes, __ProductCompositionShape__00000308_)
	__Diagram__00000061_.ProductComposition_Shapes = append(__Diagram__00000061_.ProductComposition_Shapes, __ProductCompositionShape__00000309_)
	__Diagram__00000061_.ProductComposition_Shapes = append(__Diagram__00000061_.ProductComposition_Shapes, __ProductCompositionShape__00000310_)
	__Diagram__00000061_.ProductComposition_Shapes = append(__Diagram__00000061_.ProductComposition_Shapes, __ProductCompositionShape__00000311_)
	__Diagram__00000061_.ProductComposition_Shapes = append(__Diagram__00000061_.ProductComposition_Shapes, __ProductCompositionShape__00000312_)
	__Diagram__00000061_.ProductComposition_Shapes = append(__Diagram__00000061_.ProductComposition_Shapes, __ProductCompositionShape__00000313_)
	__Diagram__00000061_.ProductComposition_Shapes = append(__Diagram__00000061_.ProductComposition_Shapes, __ProductCompositionShape__00000314_)
	__Diagram__00000061_.ProductComposition_Shapes = append(__Diagram__00000061_.ProductComposition_Shapes, __ProductCompositionShape__00000315_)
	__Diagram__00000061_.ProductComposition_Shapes = append(__Diagram__00000061_.ProductComposition_Shapes, __ProductCompositionShape__00000316_)
	__Diagram__00000061_.Task_Shapes = append(__Diagram__00000061_.Task_Shapes, __TaskShape__00000019_)
	__Diagram__00000061_.Task_Shapes = append(__Diagram__00000061_.Task_Shapes, __TaskShape__00000020_)
	__Diagram__00000061_.Task_Shapes = append(__Diagram__00000061_.Task_Shapes, __TaskShape__00000021_)
	__Diagram__00000061_.Task_Shapes = append(__Diagram__00000061_.Task_Shapes, __TaskShape__00000022_)
	__Diagram__00000061_.Task_Shapes = append(__Diagram__00000061_.Task_Shapes, __TaskShape__00000023_)
	__Diagram__00000061_.Task_Shapes = append(__Diagram__00000061_.Task_Shapes, __TaskShape__00000024_)
	__Diagram__00000061_.Task_Shapes = append(__Diagram__00000061_.Task_Shapes, __TaskShape__00000025_)
	__Diagram__00000061_.TasksWhoseOutputNodeIsExpanded = append(__Diagram__00000061_.TasksWhoseOutputNodeIsExpanded, __Task__00000007_)
	__Diagram__00000061_.TasksWhoseOutputNodeIsExpanded = append(__Diagram__00000061_.TasksWhoseOutputNodeIsExpanded, __Task__00000010_)
	__Diagram__00000061_.TaskInputShapes = append(__Diagram__00000061_.TaskInputShapes, __TaskInputShape__00000001_)
	__Diagram__00000061_.TaskInputShapes = append(__Diagram__00000061_.TaskInputShapes, __TaskInputShape__00000002_)
	__Diagram__00000061_.TaskOutputShapes = append(__Diagram__00000061_.TaskOutputShapes, __TaskOutputShape__00000063_)
	__Diagram__00000061_.TaskOutputShapes = append(__Diagram__00000061_.TaskOutputShapes, __TaskOutputShape__00000064_)
	__Diagram__00000061_.TaskOutputShapes = append(__Diagram__00000061_.TaskOutputShapes, __TaskOutputShape__00000065_)
	__Diagram__00000061_.TaskOutputShapes = append(__Diagram__00000061_.TaskOutputShapes, __TaskOutputShape__00000067_)
	__Diagram__00000061_.TaskOutputShapes = append(__Diagram__00000061_.TaskOutputShapes, __TaskOutputShape__00000068_)
	__Diagram__00000061_.TaskOutputShapes = append(__Diagram__00000061_.TaskOutputShapes, __TaskOutputShape__00000069_)
	__Diagram__00000061_.Note_Shapes = append(__Diagram__00000061_.Note_Shapes, __NoteShape__00000002_)
	__Diagram__00000061_.Note_Shapes = append(__Diagram__00000061_.Note_Shapes, __NoteShape__00000003_)
	__Diagram__00000061_.Note_Shapes = append(__Diagram__00000061_.Note_Shapes, __NoteShape__00000004_)
	__Diagram__00000061_.NoteProductShapes = append(__Diagram__00000061_.NoteProductShapes, __NoteProductShape__00000001_)
	__Diagram__00000061_.NoteProductShapes = append(__Diagram__00000061_.NoteProductShapes, __NoteProductShape__00000002_)
	__Diagram__00000061_.NoteTaskShapes = append(__Diagram__00000061_.NoteTaskShapes, __NoteTaskShape__00000001_)
	__Diagram__00000061_.NoteTaskShapes = append(__Diagram__00000061_.NoteTaskShapes, __NoteTaskShape__00000003_)
	__Diagram__00000061_.Resource_Shapes = append(__Diagram__00000061_.Resource_Shapes, __ResourceShape__00000000_)
	__Diagram__00000061_.Resource_Shapes = append(__Diagram__00000061_.Resource_Shapes, __ResourceShape__00000001_)
	__Diagram__00000061_.Resource_Shapes = append(__Diagram__00000061_.Resource_Shapes, __ResourceShape__00000002_)
	__Diagram__00000061_.Resource_Shapes = append(__Diagram__00000061_.Resource_Shapes, __ResourceShape__00000003_)
	__Diagram__00000061_.Resource_Shapes = append(__Diagram__00000061_.Resource_Shapes, __ResourceShape__00000004_)
	__Note__00000002_.Tasks = append(__Note__00000002_.Tasks, __Task__00000007_)
	__Note__00000003_.Products = append(__Note__00000003_.Products, __Product__00000035_)
	__Note__00000003_.Products = append(__Note__00000003_.Products, __Product__00000025_)
	__Note__00000003_.Products = append(__Note__00000003_.Products, __Product__00000036_)
	__Note__00000004_.Tasks = append(__Note__00000004_.Tasks, __Task__00000009_)
	__NoteProductShape__00000001_.Note = __Note__00000003_
	__NoteProductShape__00000001_.Product = __Product__00000035_
	__NoteProductShape__00000002_.Note = __Note__00000003_
	__NoteProductShape__00000002_.Product = __Product__00000025_
	__NoteShape__00000000_.Note = __Note__00000000_
	__NoteShape__00000002_.Note = __Note__00000002_
	__NoteShape__00000003_.Note = __Note__00000003_
	__NoteShape__00000004_.Note = __Note__00000004_
	__NoteTaskShape__00000001_.Note = __Note__00000002_
	__NoteTaskShape__00000001_.Task = __Task__00000007_
	__NoteTaskShape__00000003_.Note = __Note__00000004_
	__NoteTaskShape__00000003_.Task = __Task__00000009_
	__Product__00000011_.SubProducts = append(__Product__00000011_.SubProducts, __Product__00000013_)
	__Product__00000011_.SubProducts = append(__Product__00000011_.SubProducts, __Product__00000014_)
	__Product__00000014_.SubProducts = append(__Product__00000014_.SubProducts, __Product__00000016_)
	__Product__00000015_.SubProducts = append(__Product__00000015_.SubProducts, __Product__00000017_)
	__Product__00000015_.SubProducts = append(__Product__00000015_.SubProducts, __Product__00000018_)
	__Product__00000022_.SubProducts = append(__Product__00000022_.SubProducts, __Product__00000026_)
	__Product__00000022_.SubProducts = append(__Product__00000022_.SubProducts, __Product__00000037_)
	__Product__00000022_.SubProducts = append(__Product__00000022_.SubProducts, __Product__00000021_)
	__Product__00000022_.SubProducts = append(__Product__00000022_.SubProducts, __Product__00000033_)
	__Product__00000022_.SubProducts = append(__Product__00000022_.SubProducts, __Product__00000027_)
	__Product__00000022_.SubProducts = append(__Product__00000022_.SubProducts, __Product__00000028_)
	__Product__00000031_.SubProducts = append(__Product__00000031_.SubProducts, __Product__00000025_)
	__Product__00000031_.SubProducts = append(__Product__00000031_.SubProducts, __Product__00000035_)
	__Product__00000031_.SubProducts = append(__Product__00000031_.SubProducts, __Product__00000036_)
	__Product__00000031_.SubProducts = append(__Product__00000031_.SubProducts, __Product__00000029_)
	__Product__00000031_.SubProducts = append(__Product__00000031_.SubProducts, __Product__00000030_)
	__Product__00000031_.SubProducts = append(__Product__00000031_.SubProducts, __Product__00000034_)
	__Product__00000038_.SubProducts = append(__Product__00000038_.SubProducts, __Product__00000039_)
	__Product__00000038_.SubProducts = append(__Product__00000038_.SubProducts, __Product__00000040_)
	__Product__00000038_.SubProducts = append(__Product__00000038_.SubProducts, __Product__00000041_)
	__ProductCompositionShape__00000204_.Product = __Product__00000013_
	__ProductCompositionShape__00000205_.Product = __Product__00000016_
	__ProductCompositionShape__00000206_.Product = __Product__00000014_
	__ProductCompositionShape__00000207_.Product = __Product__00000017_
	__ProductCompositionShape__00000208_.Product = __Product__00000018_
	__ProductCompositionShape__00000298_.Product = __Product__00000021_
	__ProductCompositionShape__00000302_.Product = __Product__00000033_
	__ProductCompositionShape__00000307_.Product = __Product__00000026_
	__ProductCompositionShape__00000308_.Product = __Product__00000027_
	__ProductCompositionShape__00000309_.Product = __Product__00000028_
	__ProductCompositionShape__00000310_.Product = __Product__00000037_
	__ProductCompositionShape__00000311_.Product = __Product__00000025_
	__ProductCompositionShape__00000312_.Product = __Product__00000035_
	__ProductCompositionShape__00000313_.Product = __Product__00000036_
	__ProductCompositionShape__00000314_.Product = __Product__00000029_
	__ProductCompositionShape__00000315_.Product = __Product__00000030_
	__ProductCompositionShape__00000316_.Product = __Product__00000034_
	__ProductShape__00000330_.Product = __Product__00000011_
	__ProductShape__00000331_.Product = __Product__00000013_
	__ProductShape__00000332_.Product = __Product__00000014_
	__ProductShape__00000333_.Product = __Product__00000016_
	__ProductShape__00000334_.Product = __Product__00000015_
	__ProductShape__00000335_.Product = __Product__00000017_
	__ProductShape__00000336_.Product = __Product__00000018_
	__ProductShape__00000466_.Product = __Product__00000021_
	__ProductShape__00000467_.Product = __Product__00000022_
	__ProductShape__00000468_.Product = __Product__00000025_
	__ProductShape__00000469_.Product = __Product__00000026_
	__ProductShape__00000470_.Product = __Product__00000027_
	__ProductShape__00000471_.Product = __Product__00000028_
	__ProductShape__00000472_.Product = __Product__00000029_
	__ProductShape__00000473_.Product = __Product__00000030_
	__ProductShape__00000475_.Product = __Product__00000033_
	__ProductShape__00000476_.Product = __Product__00000034_
	__ProductShape__00000477_.Product = __Product__00000035_
	__ProductShape__00000478_.Product = __Product__00000036_
	__ProductShape__00000479_.Product = __Product__00000037_
	__ProductShape__00000480_.Product = __Product__00000038_
	__ProductShape__00000481_.Product = __Product__00000031_
	__Project__00000000_.RootProducts = append(__Project__00000000_.RootProducts, __Product__00000038_)
	__Project__00000000_.RootProducts = append(__Project__00000000_.RootProducts, __Product__00000042_)
	__Project__00000000_.Notes = append(__Project__00000000_.Notes, __Note__00000000_)
	__Project__00000000_.Notes = append(__Project__00000000_.Notes, __Note__00000001_)
	__Project__00000000_.Diagrams = append(__Project__00000000_.Diagrams, __Diagram__00000057_)
	__Project__00000001_.RootProducts = append(__Project__00000001_.RootProducts, __Product__00000011_)
	__Project__00000001_.RootProducts = append(__Project__00000001_.RootProducts, __Product__00000015_)
	__Project__00000001_.RootTasks = append(__Project__00000001_.RootTasks, __Task__00000003_)
	__Project__00000001_.Diagrams = append(__Project__00000001_.Diagrams, __Diagram__00000042_)
	__Project__00000002_.RootProducts = append(__Project__00000002_.RootProducts, __Product__00000022_)
	__Project__00000002_.RootProducts = append(__Project__00000002_.RootProducts, __Product__00000031_)
	__Project__00000002_.RootTasks = append(__Project__00000002_.RootTasks, __Task__00000008_)
	__Project__00000002_.RootTasks = append(__Project__00000002_.RootTasks, __Task__00000007_)
	__Project__00000002_.RootTasks = append(__Project__00000002_.RootTasks, __Task__00000010_)
	__Project__00000002_.RootTasks = append(__Project__00000002_.RootTasks, __Task__00000011_)
	__Project__00000002_.RootTasks = append(__Project__00000002_.RootTasks, __Task__00000012_)
	__Project__00000002_.RootTasks = append(__Project__00000002_.RootTasks, __Task__00000009_)
	__Project__00000002_.RootTasks = append(__Project__00000002_.RootTasks, __Task__00000013_)
	__Project__00000002_.RootResources = append(__Project__00000002_.RootResources, __Resource__00000000_)
	__Project__00000002_.RootResources = append(__Project__00000002_.RootResources, __Resource__00000001_)
	__Project__00000002_.Notes = append(__Project__00000002_.Notes, __Note__00000002_)
	__Project__00000002_.Notes = append(__Project__00000002_.Notes, __Note__00000003_)
	__Project__00000002_.Notes = append(__Project__00000002_.Notes, __Note__00000004_)
	__Project__00000002_.Diagrams = append(__Project__00000002_.Diagrams, __Diagram__00000061_)
	__ResourceShape__00000000_.Resource = __Resource__00000000_
	__ResourceShape__00000001_.Resource = __Resource__00000000_
	__ResourceShape__00000002_.Resource = __Resource__00000000_
	__ResourceShape__00000003_.Resource = __Resource__00000000_
	__ResourceShape__00000004_.Resource = __Resource__00000001_
	__Root__00000000_.Projects = append(__Root__00000000_.Projects, __Project__00000000_)
	__Root__00000000_.Projects = append(__Root__00000000_.Projects, __Project__00000001_)
	__Root__00000000_.Projects = append(__Root__00000000_.Projects, __Project__00000002_)
	__Task__00000007_.Inputs = append(__Task__00000007_.Inputs, __Product__00000025_)
	__Task__00000007_.Inputs = append(__Task__00000007_.Inputs, __Product__00000026_)
	__Task__00000007_.Outputs = append(__Task__00000007_.Outputs, __Product__00000021_)
	__Task__00000007_.Outputs = append(__Task__00000007_.Outputs, __Product__00000033_)
	__Task__00000008_.Inputs = append(__Task__00000008_.Inputs, __Product__00000025_)
	__Task__00000008_.Outputs = append(__Task__00000008_.Outputs, __Product__00000026_)
	__Task__00000009_.Outputs = append(__Task__00000009_.Outputs, __Product__00000025_)
	__Task__00000010_.Outputs = append(__Task__00000010_.Outputs, __Product__00000035_)
	__Task__00000011_.Outputs = append(__Task__00000011_.Outputs, __Product__00000035_)
	__Task__00000012_.Outputs = append(__Task__00000012_.Outputs, __Product__00000035_)
	__TaskInputShape__00000001_.Product = __Product__00000025_
	__TaskInputShape__00000001_.Task = __Task__00000008_
	__TaskInputShape__00000002_.Product = __Product__00000025_
	__TaskInputShape__00000002_.Task = __Task__00000007_
	__TaskOutputShape__00000063_.Task = __Task__00000007_
	__TaskOutputShape__00000063_.Product = __Product__00000021_
	__TaskOutputShape__00000064_.Task = __Task__00000009_
	__TaskOutputShape__00000064_.Product = __Product__00000025_
	__TaskOutputShape__00000065_.Task = __Task__00000008_
	__TaskOutputShape__00000065_.Product = __Product__00000026_
	__TaskOutputShape__00000067_.Task = __Task__00000010_
	__TaskOutputShape__00000067_.Product = __Product__00000035_
	__TaskOutputShape__00000068_.Task = __Task__00000011_
	__TaskOutputShape__00000068_.Product = __Product__00000035_
	__TaskOutputShape__00000069_.Task = __Task__00000012_
	__TaskOutputShape__00000069_.Product = __Product__00000035_
	__TaskShape__00000019_.Task = __Task__00000007_
	__TaskShape__00000020_.Task = __Task__00000008_
	__TaskShape__00000021_.Task = __Task__00000009_
	__TaskShape__00000022_.Task = __Task__00000010_
	__TaskShape__00000023_.Task = __Task__00000011_
	__TaskShape__00000024_.Task = __Task__00000012_
	__TaskShape__00000025_.Task = __Task__00000013_
}
