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

	// Declaration of instances to stage

	__Diagram__00000011_NewDiagram := (&models.Diagram{}).Stage(stage)

	__Product__00000000_UX := (&models.Product{}).Stage(stage)
	__Product__00000001_Backend := (&models.Product{}).Stage(stage)
	__Product__00000002_WBS_tree := (&models.Product{}).Stage(stage)
	__Product__00000004_PBS_tree := (&models.Product{}).Stage(stage)
	__Product__00000005_views := (&models.Product{}).Stage(stage)
	__Product__00000006_Semantic_Enforcer := (&models.Product{}).Stage(stage)
	__Product__00000009_Docx_Backend := (&models.Product{}).Stage(stage)
	__Product__00000010_Specifications := (&models.Product{}).Stage(stage)

	__ProductCompositionShape__00000059_Backend_to_Semantic_Enforcer := (&models.ProductCompositionShape{}).Stage(stage)
	__ProductCompositionShape__00000060_Backend_to_Docx_Backend := (&models.ProductCompositionShape{}).Stage(stage)
	__ProductCompositionShape__00000061_UX_to_WBS_tree := (&models.ProductCompositionShape{}).Stage(stage)
	__ProductCompositionShape__00000062_UX_to_PBS_tree := (&models.ProductCompositionShape{}).Stage(stage)
	__ProductCompositionShape__00000063_UX_to_views := (&models.ProductCompositionShape{}).Stage(stage)

	__ProductShape__00000100_Backend := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000101_Docx_Backend := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000102_PBS_tree := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000103_Semantic_Enforcer := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000104_Specifications := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000105_UX := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000106_WBS_tree := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000107_views := (&models.ProductShape{}).Stage(stage)

	__Project__00000000_Project_Editor := (&models.Project{}).Stage(stage)
	__Project__00000001_DSME_Docx := (&models.Project{}).Stage(stage)

	__Root__00000000_Root := (&models.Root{}).Stage(stage)

	__Task__00000000_Develop_Backend := (&models.Task{}).Stage(stage)
	__Task__00000001_Dev_WBS_Tree := (&models.Task{}).Stage(stage)
	__Task__00000002_Dev_PBS_Tree := (&models.Task{}).Stage(stage)
	__Task__00000003_Dev_docx_Backend := (&models.Task{}).Stage(stage)
	__Task__00000004_Dev_views := (&models.Task{}).Stage(stage)
	__Task__00000005_Dev_UXx := (&models.Task{}).Stage(stage)
	__Task__00000006_Write_Specs := (&models.Task{}).Stage(stage)

	__TaskCompositionShape__00000020_Develop_Backend_to_Dev_views := (&models.TaskCompositionShape{}).Stage(stage)
	__TaskCompositionShape__00000021_Develop_Backend_to_Dev_UXx := (&models.TaskCompositionShape{}).Stage(stage)

	__TaskInputShape__00000006_Develop_Backend_to_Specifications := (&models.TaskInputShape{}).Stage(stage)

	__TaskOutputShape__00000006_Dev_UXx_WBS_tree := (&models.TaskOutputShape{}).Stage(stage)
	__TaskOutputShape__00000007_Dev_UXx_PBS_tree := (&models.TaskOutputShape{}).Stage(stage)
	__TaskOutputShape__00000008_Dev_UXx_views := (&models.TaskOutputShape{}).Stage(stage)
	__TaskOutputShape__00000009_Dev_docx_Backend_Docx_Backend := (&models.TaskOutputShape{}).Stage(stage)
	__TaskOutputShape__00000010_Dev_docx_Backend_Semantic_Enforcer := (&models.TaskOutputShape{}).Stage(stage)
	__TaskOutputShape__00000011_Write_Specs_Specifications := (&models.TaskOutputShape{}).Stage(stage)
	__TaskOutputShape__00000018_Dev_UXx_WBS_tree := (&models.TaskOutputShape{}).Stage(stage)
	__TaskOutputShape__00000019_Dev_UXx_PBS_tree := (&models.TaskOutputShape{}).Stage(stage)
	__TaskOutputShape__00000020_Dev_UXx_views := (&models.TaskOutputShape{}).Stage(stage)
	__TaskOutputShape__00000021_Dev_docx_Backend_Docx_Backend := (&models.TaskOutputShape{}).Stage(stage)
	__TaskOutputShape__00000022_Dev_docx_Backend_Semantic_Enforcer := (&models.TaskOutputShape{}).Stage(stage)
	__TaskOutputShape__00000041_Dev_UXx_to_WBS_tree := (&models.TaskOutputShape{}).Stage(stage)
	__TaskOutputShape__00000042_Dev_UXx_to_PBS_tree := (&models.TaskOutputShape{}).Stage(stage)
	__TaskOutputShape__00000043_Dev_UXx_to_views := (&models.TaskOutputShape{}).Stage(stage)
	__TaskOutputShape__00000044_Dev_docx_Backend_to_Docx_Backend := (&models.TaskOutputShape{}).Stage(stage)
	__TaskOutputShape__00000045_Dev_docx_Backend_to_Semantic_Enforcer := (&models.TaskOutputShape{}).Stage(stage)
	__TaskOutputShape__00000046_Write_Specs_to_Specifications := (&models.TaskOutputShape{}).Stage(stage)
	__TaskOutputShape__00000053_Dev_UXx_to_WBS_tree := (&models.TaskOutputShape{}).Stage(stage)
	__TaskOutputShape__00000054_Dev_UXx_to_PBS_tree := (&models.TaskOutputShape{}).Stage(stage)
	__TaskOutputShape__00000055_Dev_UXx_to_views := (&models.TaskOutputShape{}).Stage(stage)
	__TaskOutputShape__00000056_Dev_docx_Backend_to_Docx_Backend := (&models.TaskOutputShape{}).Stage(stage)
	__TaskOutputShape__00000057_Dev_docx_Backend_to_Semantic_Enforcer := (&models.TaskOutputShape{}).Stage(stage)
	__TaskOutputShape__00000058_Write_Specs_to_Specifications := (&models.TaskOutputShape{}).Stage(stage)
	__TaskOutputShape__00000059_Dev_UXx_to_WBS_tree := (&models.TaskOutputShape{}).Stage(stage)
	__TaskOutputShape__00000060_Dev_UXx_to_PBS_tree := (&models.TaskOutputShape{}).Stage(stage)
	__TaskOutputShape__00000061_Dev_UXx_to_views := (&models.TaskOutputShape{}).Stage(stage)
	__TaskOutputShape__00000062_Dev_docx_Backend_to_Docx_Backend := (&models.TaskOutputShape{}).Stage(stage)
	__TaskOutputShape__00000063_Dev_docx_Backend_to_Semantic_Enforcer := (&models.TaskOutputShape{}).Stage(stage)
	__TaskOutputShape__00000064_Write_Specs_to_Specifications := (&models.TaskOutputShape{}).Stage(stage)

	__TaskShape__00000080_Dev_PBS_Tree := (&models.TaskShape{}).Stage(stage)
	__TaskShape__00000081_Dev_UXx := (&models.TaskShape{}).Stage(stage)
	__TaskShape__00000082_Dev_WBS_Tree := (&models.TaskShape{}).Stage(stage)
	__TaskShape__00000083_Dev_docx_Backend := (&models.TaskShape{}).Stage(stage)
	__TaskShape__00000084_Dev_views := (&models.TaskShape{}).Stage(stage)
	__TaskShape__00000085_Develop_Backend := (&models.TaskShape{}).Stage(stage)
	__TaskShape__00000086_Write_Specs := (&models.TaskShape{}).Stage(stage)

	// Setup of values

	__Diagram__00000011_NewDiagram.Name = `NewDiagram`
	__Diagram__00000011_NewDiagram.IsChecked = true
	__Diagram__00000011_NewDiagram.IsEditable_ = true
	__Diagram__00000011_NewDiagram.IsInRenameMode = false
	__Diagram__00000011_NewDiagram.IsExpanded = false
	__Diagram__00000011_NewDiagram.ComputedPrefix = ``
	__Diagram__00000011_NewDiagram.IsPBSNodeExpanded = false
	__Diagram__00000011_NewDiagram.IsWBSNodeExpanded = false

	__Product__00000000_UX.Name = `UX`
	__Product__00000000_UX.IsExpanded = false
	__Product__00000000_UX.ComputedPrefix = `2`
	__Product__00000000_UX.IsProducersNodeExpanded = false
	__Product__00000000_UX.IsConsumersNodeExpanded = false

	__Product__00000001_Backend.Name = `Backend`
	__Product__00000001_Backend.IsExpanded = false
	__Product__00000001_Backend.ComputedPrefix = `3`
	__Product__00000001_Backend.IsProducersNodeExpanded = false
	__Product__00000001_Backend.IsConsumersNodeExpanded = false

	__Product__00000002_WBS_tree.Name = `WBS tree`
	__Product__00000002_WBS_tree.IsExpanded = false
	__Product__00000002_WBS_tree.ComputedPrefix = `2.1`
	__Product__00000002_WBS_tree.IsProducersNodeExpanded = false
	__Product__00000002_WBS_tree.IsConsumersNodeExpanded = false

	__Product__00000004_PBS_tree.Name = `PBS tree`
	__Product__00000004_PBS_tree.IsExpanded = false
	__Product__00000004_PBS_tree.ComputedPrefix = `2.2`
	__Product__00000004_PBS_tree.IsProducersNodeExpanded = true
	__Product__00000004_PBS_tree.IsConsumersNodeExpanded = false

	__Product__00000005_views.Name = `views`
	__Product__00000005_views.IsExpanded = false
	__Product__00000005_views.ComputedPrefix = `2.3`
	__Product__00000005_views.IsProducersNodeExpanded = false
	__Product__00000005_views.IsConsumersNodeExpanded = false

	__Product__00000006_Semantic_Enforcer.Name = `Semantic Enforcer`
	__Product__00000006_Semantic_Enforcer.IsExpanded = false
	__Product__00000006_Semantic_Enforcer.ComputedPrefix = `3.1`
	__Product__00000006_Semantic_Enforcer.IsProducersNodeExpanded = false
	__Product__00000006_Semantic_Enforcer.IsConsumersNodeExpanded = false

	__Product__00000009_Docx_Backend.Name = `Docx Backend`
	__Product__00000009_Docx_Backend.IsExpanded = false
	__Product__00000009_Docx_Backend.ComputedPrefix = `3.2`
	__Product__00000009_Docx_Backend.IsProducersNodeExpanded = false
	__Product__00000009_Docx_Backend.IsConsumersNodeExpanded = false

	__Product__00000010_Specifications.Name = `Specifications`
	__Product__00000010_Specifications.IsExpanded = false
	__Product__00000010_Specifications.ComputedPrefix = `1`
	__Product__00000010_Specifications.IsProducersNodeExpanded = false
	__Product__00000010_Specifications.IsConsumersNodeExpanded = false

	__ProductCompositionShape__00000059_Backend_to_Semantic_Enforcer.Name = `Backend to Semantic Enforcer`
	__ProductCompositionShape__00000059_Backend_to_Semantic_Enforcer.StartRatio = 0.500000
	__ProductCompositionShape__00000059_Backend_to_Semantic_Enforcer.EndRatio = 0.500000
	__ProductCompositionShape__00000059_Backend_to_Semantic_Enforcer.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000059_Backend_to_Semantic_Enforcer.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000059_Backend_to_Semantic_Enforcer.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000060_Backend_to_Docx_Backend.Name = `Backend to Docx Backend`
	__ProductCompositionShape__00000060_Backend_to_Docx_Backend.StartRatio = 0.500000
	__ProductCompositionShape__00000060_Backend_to_Docx_Backend.EndRatio = 0.500000
	__ProductCompositionShape__00000060_Backend_to_Docx_Backend.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000060_Backend_to_Docx_Backend.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000060_Backend_to_Docx_Backend.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000061_UX_to_WBS_tree.Name = `UX to WBS tree`
	__ProductCompositionShape__00000061_UX_to_WBS_tree.StartRatio = 0.500000
	__ProductCompositionShape__00000061_UX_to_WBS_tree.EndRatio = 0.500000
	__ProductCompositionShape__00000061_UX_to_WBS_tree.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000061_UX_to_WBS_tree.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000061_UX_to_WBS_tree.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000062_UX_to_PBS_tree.Name = `UX to PBS tree`
	__ProductCompositionShape__00000062_UX_to_PBS_tree.StartRatio = 0.500000
	__ProductCompositionShape__00000062_UX_to_PBS_tree.EndRatio = 0.500000
	__ProductCompositionShape__00000062_UX_to_PBS_tree.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000062_UX_to_PBS_tree.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000062_UX_to_PBS_tree.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000063_UX_to_views.Name = `UX to views`
	__ProductCompositionShape__00000063_UX_to_views.StartRatio = 0.500000
	__ProductCompositionShape__00000063_UX_to_views.EndRatio = 0.500000
	__ProductCompositionShape__00000063_UX_to_views.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000063_UX_to_views.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000063_UX_to_views.CornerOffsetRatio = 1.680000

	__ProductShape__00000100_Backend.Name = `Backend`
	__ProductShape__00000100_Backend.IsExpanded = false
	__ProductShape__00000100_Backend.X = 50.000000
	__ProductShape__00000100_Backend.Y = 50.000000
	__ProductShape__00000100_Backend.Width = 200.000000
	__ProductShape__00000100_Backend.Height = 60.000000

	__ProductShape__00000101_Docx_Backend.Name = `Docx Backend`
	__ProductShape__00000101_Docx_Backend.IsExpanded = false
	__ProductShape__00000101_Docx_Backend.X = 484.000000
	__ProductShape__00000101_Docx_Backend.Y = 493.999985
	__ProductShape__00000101_Docx_Backend.Width = 200.000000
	__ProductShape__00000101_Docx_Backend.Height = 60.000000

	__ProductShape__00000102_PBS_tree.Name = `PBS tree`
	__ProductShape__00000102_PBS_tree.IsExpanded = false
	__ProductShape__00000102_PBS_tree.X = 1250.000000
	__ProductShape__00000102_PBS_tree.Y = 50.000000
	__ProductShape__00000102_PBS_tree.Width = 200.000000
	__ProductShape__00000102_PBS_tree.Height = 60.000000

	__ProductShape__00000103_Semantic_Enforcer.Name = `Semantic Enforcer`
	__ProductShape__00000103_Semantic_Enforcer.IsExpanded = false
	__ProductShape__00000103_Semantic_Enforcer.X = 350.000000
	__ProductShape__00000103_Semantic_Enforcer.Y = 150.000000
	__ProductShape__00000103_Semantic_Enforcer.Width = 200.000000
	__ProductShape__00000103_Semantic_Enforcer.Height = 60.000000

	__ProductShape__00000104_Specifications.Name = `Specifications`
	__ProductShape__00000104_Specifications.IsExpanded = false
	__ProductShape__00000104_Specifications.X = 350.000000
	__ProductShape__00000104_Specifications.Y = 250.000000
	__ProductShape__00000104_Specifications.Width = 200.000000
	__ProductShape__00000104_Specifications.Height = 60.000000

	__ProductShape__00000105_UX.Name = `UX`
	__ProductShape__00000105_UX.IsExpanded = false
	__ProductShape__00000105_UX.X = 50.000000
	__ProductShape__00000105_UX.Y = 150.000000
	__ProductShape__00000105_UX.Width = 200.000000
	__ProductShape__00000105_UX.Height = 60.000000

	__ProductShape__00000106_WBS_tree.Name = `WBS tree`
	__ProductShape__00000106_WBS_tree.IsExpanded = false
	__ProductShape__00000106_WBS_tree.X = 1250.000000
	__ProductShape__00000106_WBS_tree.Y = 150.000000
	__ProductShape__00000106_WBS_tree.Width = 200.000000
	__ProductShape__00000106_WBS_tree.Height = 60.000000

	__ProductShape__00000107_views.Name = `views`
	__ProductShape__00000107_views.IsExpanded = false
	__ProductShape__00000107_views.X = 1250.000000
	__ProductShape__00000107_views.Y = 250.000000
	__ProductShape__00000107_views.Width = 200.000000
	__ProductShape__00000107_views.Height = 60.000000

	__Project__00000000_Project_Editor.Name = `Project Editor`
	__Project__00000000_Project_Editor.IsPBSNodeExpanded = true
	__Project__00000000_Project_Editor.IsWBSNodeExpanded = true
	__Project__00000000_Project_Editor.IsDiagramsNodeExpanded = true
	__Project__00000000_Project_Editor.IsExpanded = true
	__Project__00000000_Project_Editor.ComputedPrefix = ``

	__Project__00000001_DSME_Docx.Name = `DSME Docx`
	__Project__00000001_DSME_Docx.IsPBSNodeExpanded = false
	__Project__00000001_DSME_Docx.IsWBSNodeExpanded = true
	__Project__00000001_DSME_Docx.IsDiagramsNodeExpanded = false
	__Project__00000001_DSME_Docx.IsExpanded = true
	__Project__00000001_DSME_Docx.ComputedPrefix = ``

	__Root__00000000_Root.Name = `Root`
	__Root__00000000_Root.NbPixPerCharacter = 8.000000

	__Task__00000000_Develop_Backend.Name = `Develop Backend`
	__Task__00000000_Develop_Backend.IsExpanded = false
	__Task__00000000_Develop_Backend.ComputedPrefix = `2`
	__Task__00000000_Develop_Backend.IsInputsNodeExpanded = false
	__Task__00000000_Develop_Backend.IsOutputsNodeExpanded = false

	__Task__00000001_Dev_WBS_Tree.Name = `Dev WBS Tree`
	__Task__00000001_Dev_WBS_Tree.IsExpanded = false
	__Task__00000001_Dev_WBS_Tree.ComputedPrefix = `3`
	__Task__00000001_Dev_WBS_Tree.IsInputsNodeExpanded = false
	__Task__00000001_Dev_WBS_Tree.IsOutputsNodeExpanded = false

	__Task__00000002_Dev_PBS_Tree.Name = `Dev PBS Tree`
	__Task__00000002_Dev_PBS_Tree.IsExpanded = false
	__Task__00000002_Dev_PBS_Tree.ComputedPrefix = `4`
	__Task__00000002_Dev_PBS_Tree.IsInputsNodeExpanded = false
	__Task__00000002_Dev_PBS_Tree.IsOutputsNodeExpanded = false

	__Task__00000003_Dev_docx_Backend.Name = `Dev docx Backend`
	__Task__00000003_Dev_docx_Backend.IsExpanded = true
	__Task__00000003_Dev_docx_Backend.ComputedPrefix = `1`
	__Task__00000003_Dev_docx_Backend.IsInputsNodeExpanded = true
	__Task__00000003_Dev_docx_Backend.IsOutputsNodeExpanded = true

	__Task__00000004_Dev_views.Name = `Dev views`
	__Task__00000004_Dev_views.IsExpanded = false
	__Task__00000004_Dev_views.ComputedPrefix = `2.1`
	__Task__00000004_Dev_views.IsInputsNodeExpanded = false
	__Task__00000004_Dev_views.IsOutputsNodeExpanded = false

	__Task__00000005_Dev_UXx.Name = `Dev UXx`
	__Task__00000005_Dev_UXx.IsExpanded = true
	__Task__00000005_Dev_UXx.ComputedPrefix = `2.2`
	__Task__00000005_Dev_UXx.IsInputsNodeExpanded = false
	__Task__00000005_Dev_UXx.IsOutputsNodeExpanded = false

	__Task__00000006_Write_Specs.Name = `Write Specs`
	__Task__00000006_Write_Specs.IsExpanded = false
	__Task__00000006_Write_Specs.ComputedPrefix = `1`
	__Task__00000006_Write_Specs.IsInputsNodeExpanded = false
	__Task__00000006_Write_Specs.IsOutputsNodeExpanded = false

	__TaskCompositionShape__00000020_Develop_Backend_to_Dev_views.Name = `Develop Backend to Dev views`
	__TaskCompositionShape__00000020_Develop_Backend_to_Dev_views.StartRatio = 0.500000
	__TaskCompositionShape__00000020_Develop_Backend_to_Dev_views.EndRatio = 0.500000
	__TaskCompositionShape__00000020_Develop_Backend_to_Dev_views.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000020_Develop_Backend_to_Dev_views.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000020_Develop_Backend_to_Dev_views.CornerOffsetRatio = 1.680000

	__TaskCompositionShape__00000021_Develop_Backend_to_Dev_UXx.Name = `Develop Backend to Dev UXx`
	__TaskCompositionShape__00000021_Develop_Backend_to_Dev_UXx.StartRatio = 0.500000
	__TaskCompositionShape__00000021_Develop_Backend_to_Dev_UXx.EndRatio = 0.500000
	__TaskCompositionShape__00000021_Develop_Backend_to_Dev_UXx.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000021_Develop_Backend_to_Dev_UXx.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000021_Develop_Backend_to_Dev_UXx.CornerOffsetRatio = 1.680000

	__TaskInputShape__00000006_Develop_Backend_to_Specifications.Name = `Develop Backend to Specifications`
	__TaskInputShape__00000006_Develop_Backend_to_Specifications.StartRatio = 0.500000
	__TaskInputShape__00000006_Develop_Backend_to_Specifications.EndRatio = 0.500000
	__TaskInputShape__00000006_Develop_Backend_to_Specifications.StartOrientation = models.ORIENTATION_HORIZONTAL
	__TaskInputShape__00000006_Develop_Backend_to_Specifications.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskInputShape__00000006_Develop_Backend_to_Specifications.CornerOffsetRatio = 1.680000

	__TaskOutputShape__00000006_Dev_UXx_WBS_tree.Name = `Dev UXx->WBS tree`
	__TaskOutputShape__00000006_Dev_UXx_WBS_tree.StartRatio = 0.000000
	__TaskOutputShape__00000006_Dev_UXx_WBS_tree.EndRatio = 0.000000
	__TaskOutputShape__00000006_Dev_UXx_WBS_tree.CornerOffsetRatio = 0.000000

	__TaskOutputShape__00000007_Dev_UXx_PBS_tree.Name = `Dev UXx->PBS tree`
	__TaskOutputShape__00000007_Dev_UXx_PBS_tree.StartRatio = 0.000000
	__TaskOutputShape__00000007_Dev_UXx_PBS_tree.EndRatio = 0.000000
	__TaskOutputShape__00000007_Dev_UXx_PBS_tree.CornerOffsetRatio = 0.000000

	__TaskOutputShape__00000008_Dev_UXx_views.Name = `Dev UXx->views`
	__TaskOutputShape__00000008_Dev_UXx_views.StartRatio = 0.000000
	__TaskOutputShape__00000008_Dev_UXx_views.EndRatio = 0.000000
	__TaskOutputShape__00000008_Dev_UXx_views.CornerOffsetRatio = 0.000000

	__TaskOutputShape__00000009_Dev_docx_Backend_Docx_Backend.Name = `Dev docx Backend->Docx Backend`
	__TaskOutputShape__00000009_Dev_docx_Backend_Docx_Backend.StartRatio = 0.000000
	__TaskOutputShape__00000009_Dev_docx_Backend_Docx_Backend.EndRatio = 0.000000
	__TaskOutputShape__00000009_Dev_docx_Backend_Docx_Backend.CornerOffsetRatio = 0.000000

	__TaskOutputShape__00000010_Dev_docx_Backend_Semantic_Enforcer.Name = `Dev docx Backend->Semantic Enforcer`
	__TaskOutputShape__00000010_Dev_docx_Backend_Semantic_Enforcer.StartRatio = 0.000000
	__TaskOutputShape__00000010_Dev_docx_Backend_Semantic_Enforcer.EndRatio = 0.000000
	__TaskOutputShape__00000010_Dev_docx_Backend_Semantic_Enforcer.CornerOffsetRatio = 0.000000

	__TaskOutputShape__00000011_Write_Specs_Specifications.Name = `Write Specs->Specifications`
	__TaskOutputShape__00000011_Write_Specs_Specifications.StartRatio = 0.000000
	__TaskOutputShape__00000011_Write_Specs_Specifications.EndRatio = 0.000000
	__TaskOutputShape__00000011_Write_Specs_Specifications.CornerOffsetRatio = 0.000000

	__TaskOutputShape__00000018_Dev_UXx_WBS_tree.Name = `Dev UXx->WBS tree`
	__TaskOutputShape__00000018_Dev_UXx_WBS_tree.StartRatio = 0.000000
	__TaskOutputShape__00000018_Dev_UXx_WBS_tree.EndRatio = 0.000000
	__TaskOutputShape__00000018_Dev_UXx_WBS_tree.CornerOffsetRatio = 0.000000

	__TaskOutputShape__00000019_Dev_UXx_PBS_tree.Name = `Dev UXx->PBS tree`
	__TaskOutputShape__00000019_Dev_UXx_PBS_tree.StartRatio = 0.000000
	__TaskOutputShape__00000019_Dev_UXx_PBS_tree.EndRatio = 0.000000
	__TaskOutputShape__00000019_Dev_UXx_PBS_tree.CornerOffsetRatio = 0.000000

	__TaskOutputShape__00000020_Dev_UXx_views.Name = `Dev UXx->views`
	__TaskOutputShape__00000020_Dev_UXx_views.StartRatio = 0.000000
	__TaskOutputShape__00000020_Dev_UXx_views.EndRatio = 0.000000
	__TaskOutputShape__00000020_Dev_UXx_views.CornerOffsetRatio = 0.000000

	__TaskOutputShape__00000021_Dev_docx_Backend_Docx_Backend.Name = `Dev docx Backend->Docx Backend`
	__TaskOutputShape__00000021_Dev_docx_Backend_Docx_Backend.StartRatio = 0.000000
	__TaskOutputShape__00000021_Dev_docx_Backend_Docx_Backend.EndRatio = 0.000000
	__TaskOutputShape__00000021_Dev_docx_Backend_Docx_Backend.CornerOffsetRatio = 0.000000

	__TaskOutputShape__00000022_Dev_docx_Backend_Semantic_Enforcer.Name = `Dev docx Backend->Semantic Enforcer`
	__TaskOutputShape__00000022_Dev_docx_Backend_Semantic_Enforcer.StartRatio = 0.000000
	__TaskOutputShape__00000022_Dev_docx_Backend_Semantic_Enforcer.EndRatio = 0.000000
	__TaskOutputShape__00000022_Dev_docx_Backend_Semantic_Enforcer.CornerOffsetRatio = 0.000000

	__TaskOutputShape__00000041_Dev_UXx_to_WBS_tree.Name = `Dev UXx to WBS tree`
	__TaskOutputShape__00000041_Dev_UXx_to_WBS_tree.StartRatio = 0.500000
	__TaskOutputShape__00000041_Dev_UXx_to_WBS_tree.EndRatio = 0.500000
	__TaskOutputShape__00000041_Dev_UXx_to_WBS_tree.StartOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000041_Dev_UXx_to_WBS_tree.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000041_Dev_UXx_to_WBS_tree.CornerOffsetRatio = 1.680000

	__TaskOutputShape__00000042_Dev_UXx_to_PBS_tree.Name = `Dev UXx to PBS tree`
	__TaskOutputShape__00000042_Dev_UXx_to_PBS_tree.StartRatio = 0.500000
	__TaskOutputShape__00000042_Dev_UXx_to_PBS_tree.EndRatio = 0.500000
	__TaskOutputShape__00000042_Dev_UXx_to_PBS_tree.StartOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000042_Dev_UXx_to_PBS_tree.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000042_Dev_UXx_to_PBS_tree.CornerOffsetRatio = 1.680000

	__TaskOutputShape__00000043_Dev_UXx_to_views.Name = `Dev UXx to views`
	__TaskOutputShape__00000043_Dev_UXx_to_views.StartRatio = 0.500000
	__TaskOutputShape__00000043_Dev_UXx_to_views.EndRatio = 0.500000
	__TaskOutputShape__00000043_Dev_UXx_to_views.StartOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000043_Dev_UXx_to_views.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000043_Dev_UXx_to_views.CornerOffsetRatio = 1.680000

	__TaskOutputShape__00000044_Dev_docx_Backend_to_Docx_Backend.Name = `Dev docx Backend to Docx Backend`
	__TaskOutputShape__00000044_Dev_docx_Backend_to_Docx_Backend.StartRatio = 0.500000
	__TaskOutputShape__00000044_Dev_docx_Backend_to_Docx_Backend.EndRatio = 0.500000
	__TaskOutputShape__00000044_Dev_docx_Backend_to_Docx_Backend.StartOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000044_Dev_docx_Backend_to_Docx_Backend.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000044_Dev_docx_Backend_to_Docx_Backend.CornerOffsetRatio = 1.680000

	__TaskOutputShape__00000045_Dev_docx_Backend_to_Semantic_Enforcer.Name = `Dev docx Backend to Semantic Enforcer`
	__TaskOutputShape__00000045_Dev_docx_Backend_to_Semantic_Enforcer.StartRatio = 0.522222
	__TaskOutputShape__00000045_Dev_docx_Backend_to_Semantic_Enforcer.EndRatio = 0.500000
	__TaskOutputShape__00000045_Dev_docx_Backend_to_Semantic_Enforcer.StartOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000045_Dev_docx_Backend_to_Semantic_Enforcer.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000045_Dev_docx_Backend_to_Semantic_Enforcer.CornerOffsetRatio = 1.229010

	__TaskOutputShape__00000046_Write_Specs_to_Specifications.Name = `Write Specs to Specifications`
	__TaskOutputShape__00000046_Write_Specs_to_Specifications.StartRatio = 0.500000
	__TaskOutputShape__00000046_Write_Specs_to_Specifications.EndRatio = 0.500000
	__TaskOutputShape__00000046_Write_Specs_to_Specifications.StartOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000046_Write_Specs_to_Specifications.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000046_Write_Specs_to_Specifications.CornerOffsetRatio = 1.680000

	__TaskOutputShape__00000053_Dev_UXx_to_WBS_tree.Name = `Dev UXx to WBS tree`
	__TaskOutputShape__00000053_Dev_UXx_to_WBS_tree.StartRatio = 0.500000
	__TaskOutputShape__00000053_Dev_UXx_to_WBS_tree.EndRatio = 0.500000
	__TaskOutputShape__00000053_Dev_UXx_to_WBS_tree.StartOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000053_Dev_UXx_to_WBS_tree.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000053_Dev_UXx_to_WBS_tree.CornerOffsetRatio = 1.680000

	__TaskOutputShape__00000054_Dev_UXx_to_PBS_tree.Name = `Dev UXx to PBS tree`
	__TaskOutputShape__00000054_Dev_UXx_to_PBS_tree.StartRatio = 0.500000
	__TaskOutputShape__00000054_Dev_UXx_to_PBS_tree.EndRatio = 0.500000
	__TaskOutputShape__00000054_Dev_UXx_to_PBS_tree.StartOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000054_Dev_UXx_to_PBS_tree.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000054_Dev_UXx_to_PBS_tree.CornerOffsetRatio = 1.680000

	__TaskOutputShape__00000055_Dev_UXx_to_views.Name = `Dev UXx to views`
	__TaskOutputShape__00000055_Dev_UXx_to_views.StartRatio = 0.500000
	__TaskOutputShape__00000055_Dev_UXx_to_views.EndRatio = 0.500000
	__TaskOutputShape__00000055_Dev_UXx_to_views.StartOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000055_Dev_UXx_to_views.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000055_Dev_UXx_to_views.CornerOffsetRatio = 1.680000

	__TaskOutputShape__00000056_Dev_docx_Backend_to_Docx_Backend.Name = `Dev docx Backend to Docx Backend`
	__TaskOutputShape__00000056_Dev_docx_Backend_to_Docx_Backend.StartRatio = 0.500000
	__TaskOutputShape__00000056_Dev_docx_Backend_to_Docx_Backend.EndRatio = 0.500000
	__TaskOutputShape__00000056_Dev_docx_Backend_to_Docx_Backend.StartOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000056_Dev_docx_Backend_to_Docx_Backend.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000056_Dev_docx_Backend_to_Docx_Backend.CornerOffsetRatio = 1.680000

	__TaskOutputShape__00000057_Dev_docx_Backend_to_Semantic_Enforcer.Name = `Dev docx Backend to Semantic Enforcer`
	__TaskOutputShape__00000057_Dev_docx_Backend_to_Semantic_Enforcer.StartRatio = 0.500000
	__TaskOutputShape__00000057_Dev_docx_Backend_to_Semantic_Enforcer.EndRatio = 0.500000
	__TaskOutputShape__00000057_Dev_docx_Backend_to_Semantic_Enforcer.StartOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000057_Dev_docx_Backend_to_Semantic_Enforcer.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000057_Dev_docx_Backend_to_Semantic_Enforcer.CornerOffsetRatio = 1.680000

	__TaskOutputShape__00000058_Write_Specs_to_Specifications.Name = `Write Specs to Specifications`
	__TaskOutputShape__00000058_Write_Specs_to_Specifications.StartRatio = 0.500000
	__TaskOutputShape__00000058_Write_Specs_to_Specifications.EndRatio = 0.500000
	__TaskOutputShape__00000058_Write_Specs_to_Specifications.StartOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000058_Write_Specs_to_Specifications.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000058_Write_Specs_to_Specifications.CornerOffsetRatio = 1.209010

	__TaskOutputShape__00000059_Dev_UXx_to_WBS_tree.Name = `Dev UXx to WBS tree`
	__TaskOutputShape__00000059_Dev_UXx_to_WBS_tree.StartRatio = 0.500000
	__TaskOutputShape__00000059_Dev_UXx_to_WBS_tree.EndRatio = 0.500000
	__TaskOutputShape__00000059_Dev_UXx_to_WBS_tree.StartOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000059_Dev_UXx_to_WBS_tree.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000059_Dev_UXx_to_WBS_tree.CornerOffsetRatio = 1.680000

	__TaskOutputShape__00000060_Dev_UXx_to_PBS_tree.Name = `Dev UXx to PBS tree`
	__TaskOutputShape__00000060_Dev_UXx_to_PBS_tree.StartRatio = 0.500000
	__TaskOutputShape__00000060_Dev_UXx_to_PBS_tree.EndRatio = 0.500000
	__TaskOutputShape__00000060_Dev_UXx_to_PBS_tree.StartOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000060_Dev_UXx_to_PBS_tree.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000060_Dev_UXx_to_PBS_tree.CornerOffsetRatio = 1.680000

	__TaskOutputShape__00000061_Dev_UXx_to_views.Name = `Dev UXx to views`
	__TaskOutputShape__00000061_Dev_UXx_to_views.StartRatio = 0.500000
	__TaskOutputShape__00000061_Dev_UXx_to_views.EndRatio = 0.500000
	__TaskOutputShape__00000061_Dev_UXx_to_views.StartOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000061_Dev_UXx_to_views.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000061_Dev_UXx_to_views.CornerOffsetRatio = 1.680000

	__TaskOutputShape__00000062_Dev_docx_Backend_to_Docx_Backend.Name = `Dev docx Backend to Docx Backend`
	__TaskOutputShape__00000062_Dev_docx_Backend_to_Docx_Backend.StartRatio = 0.500000
	__TaskOutputShape__00000062_Dev_docx_Backend_to_Docx_Backend.EndRatio = 0.500000
	__TaskOutputShape__00000062_Dev_docx_Backend_to_Docx_Backend.StartOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000062_Dev_docx_Backend_to_Docx_Backend.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000062_Dev_docx_Backend_to_Docx_Backend.CornerOffsetRatio = 1.680000

	__TaskOutputShape__00000063_Dev_docx_Backend_to_Semantic_Enforcer.Name = `Dev docx Backend to Semantic Enforcer`
	__TaskOutputShape__00000063_Dev_docx_Backend_to_Semantic_Enforcer.StartRatio = 0.500000
	__TaskOutputShape__00000063_Dev_docx_Backend_to_Semantic_Enforcer.EndRatio = 0.500000
	__TaskOutputShape__00000063_Dev_docx_Backend_to_Semantic_Enforcer.StartOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000063_Dev_docx_Backend_to_Semantic_Enforcer.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000063_Dev_docx_Backend_to_Semantic_Enforcer.CornerOffsetRatio = 1.680000

	__TaskOutputShape__00000064_Write_Specs_to_Specifications.Name = `Write Specs to Specifications`
	__TaskOutputShape__00000064_Write_Specs_to_Specifications.StartRatio = 0.500000
	__TaskOutputShape__00000064_Write_Specs_to_Specifications.EndRatio = 0.500000
	__TaskOutputShape__00000064_Write_Specs_to_Specifications.StartOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000064_Write_Specs_to_Specifications.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskOutputShape__00000064_Write_Specs_to_Specifications.CornerOffsetRatio = 1.680000

	__TaskShape__00000080_Dev_PBS_Tree.Name = `Dev PBS Tree`
	__TaskShape__00000080_Dev_PBS_Tree.IsExpanded = false
	__TaskShape__00000080_Dev_PBS_Tree.X = 50.000000
	__TaskShape__00000080_Dev_PBS_Tree.Y = 250.000000
	__TaskShape__00000080_Dev_PBS_Tree.Width = 200.000000
	__TaskShape__00000080_Dev_PBS_Tree.Height = 60.000000

	__TaskShape__00000081_Dev_UXx.Name = `Dev UXx`
	__TaskShape__00000081_Dev_UXx.IsExpanded = false
	__TaskShape__00000081_Dev_UXx.X = 950.000000
	__TaskShape__00000081_Dev_UXx.Y = 50.000000
	__TaskShape__00000081_Dev_UXx.Width = 200.000000
	__TaskShape__00000081_Dev_UXx.Height = 60.000000

	__TaskShape__00000082_Dev_WBS_Tree.Name = `Dev WBS Tree`
	__TaskShape__00000082_Dev_WBS_Tree.IsExpanded = false
	__TaskShape__00000082_Dev_WBS_Tree.X = 50.000000
	__TaskShape__00000082_Dev_WBS_Tree.Y = 350.000000
	__TaskShape__00000082_Dev_WBS_Tree.Width = 200.000000
	__TaskShape__00000082_Dev_WBS_Tree.Height = 60.000000

	__TaskShape__00000083_Dev_docx_Backend.Name = `Dev docx Backend`
	__TaskShape__00000083_Dev_docx_Backend.IsExpanded = false
	__TaskShape__00000083_Dev_docx_Backend.X = 50.000000
	__TaskShape__00000083_Dev_docx_Backend.Y = 450.000000
	__TaskShape__00000083_Dev_docx_Backend.Width = 200.000000
	__TaskShape__00000083_Dev_docx_Backend.Height = 60.000000

	__TaskShape__00000084_Dev_views.Name = `Dev views`
	__TaskShape__00000084_Dev_views.IsExpanded = false
	__TaskShape__00000084_Dev_views.X = 950.000000
	__TaskShape__00000084_Dev_views.Y = 150.000000
	__TaskShape__00000084_Dev_views.Width = 200.000000
	__TaskShape__00000084_Dev_views.Height = 60.000000

	__TaskShape__00000085_Develop_Backend.Name = `Develop Backend`
	__TaskShape__00000085_Develop_Backend.IsExpanded = false
	__TaskShape__00000085_Develop_Backend.X = 650.000000
	__TaskShape__00000085_Develop_Backend.Y = 50.000000
	__TaskShape__00000085_Develop_Backend.Width = 200.000000
	__TaskShape__00000085_Develop_Backend.Height = 60.000000

	__TaskShape__00000086_Write_Specs.Name = `Write Specs`
	__TaskShape__00000086_Write_Specs.IsExpanded = false
	__TaskShape__00000086_Write_Specs.X = 50.000000
	__TaskShape__00000086_Write_Specs.Y = 550.000000
	__TaskShape__00000086_Write_Specs.Width = 200.000000
	__TaskShape__00000086_Write_Specs.Height = 60.000000

	// Setup of pointers
	// setup of Diagram instances pointers
	__Diagram__00000011_NewDiagram.Product_Shapes = append(__Diagram__00000011_NewDiagram.Product_Shapes, __ProductShape__00000100_Backend)
	__Diagram__00000011_NewDiagram.Product_Shapes = append(__Diagram__00000011_NewDiagram.Product_Shapes, __ProductShape__00000101_Docx_Backend)
	__Diagram__00000011_NewDiagram.Product_Shapes = append(__Diagram__00000011_NewDiagram.Product_Shapes, __ProductShape__00000102_PBS_tree)
	__Diagram__00000011_NewDiagram.Product_Shapes = append(__Diagram__00000011_NewDiagram.Product_Shapes, __ProductShape__00000103_Semantic_Enforcer)
	__Diagram__00000011_NewDiagram.Product_Shapes = append(__Diagram__00000011_NewDiagram.Product_Shapes, __ProductShape__00000104_Specifications)
	__Diagram__00000011_NewDiagram.Product_Shapes = append(__Diagram__00000011_NewDiagram.Product_Shapes, __ProductShape__00000105_UX)
	__Diagram__00000011_NewDiagram.Product_Shapes = append(__Diagram__00000011_NewDiagram.Product_Shapes, __ProductShape__00000106_WBS_tree)
	__Diagram__00000011_NewDiagram.Product_Shapes = append(__Diagram__00000011_NewDiagram.Product_Shapes, __ProductShape__00000107_views)
	__Diagram__00000011_NewDiagram.ProductComposition_Shapes = append(__Diagram__00000011_NewDiagram.ProductComposition_Shapes, __ProductCompositionShape__00000059_Backend_to_Semantic_Enforcer)
	__Diagram__00000011_NewDiagram.ProductComposition_Shapes = append(__Diagram__00000011_NewDiagram.ProductComposition_Shapes, __ProductCompositionShape__00000060_Backend_to_Docx_Backend)
	__Diagram__00000011_NewDiagram.ProductComposition_Shapes = append(__Diagram__00000011_NewDiagram.ProductComposition_Shapes, __ProductCompositionShape__00000061_UX_to_WBS_tree)
	__Diagram__00000011_NewDiagram.ProductComposition_Shapes = append(__Diagram__00000011_NewDiagram.ProductComposition_Shapes, __ProductCompositionShape__00000062_UX_to_PBS_tree)
	__Diagram__00000011_NewDiagram.ProductComposition_Shapes = append(__Diagram__00000011_NewDiagram.ProductComposition_Shapes, __ProductCompositionShape__00000063_UX_to_views)
	__Diagram__00000011_NewDiagram.Task_Shapes = append(__Diagram__00000011_NewDiagram.Task_Shapes, __TaskShape__00000080_Dev_PBS_Tree)
	__Diagram__00000011_NewDiagram.Task_Shapes = append(__Diagram__00000011_NewDiagram.Task_Shapes, __TaskShape__00000081_Dev_UXx)
	__Diagram__00000011_NewDiagram.Task_Shapes = append(__Diagram__00000011_NewDiagram.Task_Shapes, __TaskShape__00000082_Dev_WBS_Tree)
	__Diagram__00000011_NewDiagram.Task_Shapes = append(__Diagram__00000011_NewDiagram.Task_Shapes, __TaskShape__00000083_Dev_docx_Backend)
	__Diagram__00000011_NewDiagram.Task_Shapes = append(__Diagram__00000011_NewDiagram.Task_Shapes, __TaskShape__00000084_Dev_views)
	__Diagram__00000011_NewDiagram.Task_Shapes = append(__Diagram__00000011_NewDiagram.Task_Shapes, __TaskShape__00000085_Develop_Backend)
	__Diagram__00000011_NewDiagram.Task_Shapes = append(__Diagram__00000011_NewDiagram.Task_Shapes, __TaskShape__00000086_Write_Specs)
	__Diagram__00000011_NewDiagram.TaskComposition_Shapes = append(__Diagram__00000011_NewDiagram.TaskComposition_Shapes, __TaskCompositionShape__00000020_Develop_Backend_to_Dev_views)
	__Diagram__00000011_NewDiagram.TaskComposition_Shapes = append(__Diagram__00000011_NewDiagram.TaskComposition_Shapes, __TaskCompositionShape__00000021_Develop_Backend_to_Dev_UXx)
	__Diagram__00000011_NewDiagram.TaskInputShapes = append(__Diagram__00000011_NewDiagram.TaskInputShapes, __TaskInputShape__00000006_Develop_Backend_to_Specifications)
	__Diagram__00000011_NewDiagram.TaskOutputShapes = append(__Diagram__00000011_NewDiagram.TaskOutputShapes, __TaskOutputShape__00000059_Dev_UXx_to_WBS_tree)
	__Diagram__00000011_NewDiagram.TaskOutputShapes = append(__Diagram__00000011_NewDiagram.TaskOutputShapes, __TaskOutputShape__00000060_Dev_UXx_to_PBS_tree)
	__Diagram__00000011_NewDiagram.TaskOutputShapes = append(__Diagram__00000011_NewDiagram.TaskOutputShapes, __TaskOutputShape__00000061_Dev_UXx_to_views)
	__Diagram__00000011_NewDiagram.TaskOutputShapes = append(__Diagram__00000011_NewDiagram.TaskOutputShapes, __TaskOutputShape__00000062_Dev_docx_Backend_to_Docx_Backend)
	__Diagram__00000011_NewDiagram.TaskOutputShapes = append(__Diagram__00000011_NewDiagram.TaskOutputShapes, __TaskOutputShape__00000063_Dev_docx_Backend_to_Semantic_Enforcer)
	__Diagram__00000011_NewDiagram.TaskOutputShapes = append(__Diagram__00000011_NewDiagram.TaskOutputShapes, __TaskOutputShape__00000064_Write_Specs_to_Specifications)
	// setup of Product instances pointers
	__Product__00000000_UX.SubProducts = append(__Product__00000000_UX.SubProducts, __Product__00000002_WBS_tree)
	__Product__00000000_UX.SubProducts = append(__Product__00000000_UX.SubProducts, __Product__00000004_PBS_tree)
	__Product__00000000_UX.SubProducts = append(__Product__00000000_UX.SubProducts, __Product__00000005_views)
	__Product__00000001_Backend.SubProducts = append(__Product__00000001_Backend.SubProducts, __Product__00000006_Semantic_Enforcer)
	__Product__00000001_Backend.SubProducts = append(__Product__00000001_Backend.SubProducts, __Product__00000009_Docx_Backend)
	// setup of ProductCompositionShape instances pointers
	__ProductCompositionShape__00000059_Backend_to_Semantic_Enforcer.Product = __Product__00000006_Semantic_Enforcer
	__ProductCompositionShape__00000060_Backend_to_Docx_Backend.Product = __Product__00000009_Docx_Backend
	__ProductCompositionShape__00000061_UX_to_WBS_tree.Product = __Product__00000002_WBS_tree
	__ProductCompositionShape__00000062_UX_to_PBS_tree.Product = __Product__00000004_PBS_tree
	__ProductCompositionShape__00000063_UX_to_views.Product = __Product__00000005_views
	// setup of ProductShape instances pointers
	__ProductShape__00000100_Backend.Product = __Product__00000001_Backend
	__ProductShape__00000101_Docx_Backend.Product = __Product__00000009_Docx_Backend
	__ProductShape__00000102_PBS_tree.Product = __Product__00000004_PBS_tree
	__ProductShape__00000103_Semantic_Enforcer.Product = __Product__00000006_Semantic_Enforcer
	__ProductShape__00000104_Specifications.Product = __Product__00000010_Specifications
	__ProductShape__00000105_UX.Product = __Product__00000000_UX
	__ProductShape__00000106_WBS_tree.Product = __Product__00000002_WBS_tree
	__ProductShape__00000107_views.Product = __Product__00000005_views
	// setup of Project instances pointers
	__Project__00000000_Project_Editor.RootProducts = append(__Project__00000000_Project_Editor.RootProducts, __Product__00000010_Specifications)
	__Project__00000000_Project_Editor.RootProducts = append(__Project__00000000_Project_Editor.RootProducts, __Product__00000000_UX)
	__Project__00000000_Project_Editor.RootProducts = append(__Project__00000000_Project_Editor.RootProducts, __Product__00000001_Backend)
	__Project__00000000_Project_Editor.RootTasks = append(__Project__00000000_Project_Editor.RootTasks, __Task__00000006_Write_Specs)
	__Project__00000000_Project_Editor.RootTasks = append(__Project__00000000_Project_Editor.RootTasks, __Task__00000000_Develop_Backend)
	__Project__00000000_Project_Editor.RootTasks = append(__Project__00000000_Project_Editor.RootTasks, __Task__00000001_Dev_WBS_Tree)
	__Project__00000000_Project_Editor.RootTasks = append(__Project__00000000_Project_Editor.RootTasks, __Task__00000002_Dev_PBS_Tree)
	__Project__00000000_Project_Editor.Diagrams = append(__Project__00000000_Project_Editor.Diagrams, __Diagram__00000011_NewDiagram)
	__Project__00000001_DSME_Docx.RootTasks = append(__Project__00000001_DSME_Docx.RootTasks, __Task__00000003_Dev_docx_Backend)
	// setup of Root instances pointers
	__Root__00000000_Root.Projects = append(__Root__00000000_Root.Projects, __Project__00000000_Project_Editor)
	__Root__00000000_Root.Projects = append(__Root__00000000_Root.Projects, __Project__00000001_DSME_Docx)
	// setup of Task instances pointers
	__Task__00000000_Develop_Backend.SubTasks = append(__Task__00000000_Develop_Backend.SubTasks, __Task__00000004_Dev_views)
	__Task__00000000_Develop_Backend.SubTasks = append(__Task__00000000_Develop_Backend.SubTasks, __Task__00000005_Dev_UXx)
	__Task__00000000_Develop_Backend.Inputs = append(__Task__00000000_Develop_Backend.Inputs, __Product__00000010_Specifications)
	__Task__00000005_Dev_UXx.Outputs = append(__Task__00000005_Dev_UXx.Outputs, __Product__00000002_WBS_tree)
	__Task__00000005_Dev_UXx.Outputs = append(__Task__00000005_Dev_UXx.Outputs, __Product__00000004_PBS_tree)
	__Task__00000005_Dev_UXx.Outputs = append(__Task__00000005_Dev_UXx.Outputs, __Product__00000005_views)
	__Task__00000006_Write_Specs.Outputs = append(__Task__00000006_Write_Specs.Outputs, __Product__00000010_Specifications)
	// setup of TaskCompositionShape instances pointers
	__TaskCompositionShape__00000020_Develop_Backend_to_Dev_views.Task = __Task__00000004_Dev_views
	__TaskCompositionShape__00000021_Develop_Backend_to_Dev_UXx.Task = __Task__00000005_Dev_UXx
	// setup of TaskInputShape instances pointers
	__TaskInputShape__00000006_Develop_Backend_to_Specifications.Task = __Task__00000000_Develop_Backend
	__TaskInputShape__00000006_Develop_Backend_to_Specifications.Product = __Product__00000010_Specifications
	// setup of TaskOutputShape instances pointers
	__TaskOutputShape__00000006_Dev_UXx_WBS_tree.Task = __Task__00000005_Dev_UXx
	__TaskOutputShape__00000006_Dev_UXx_WBS_tree.Product = __Product__00000002_WBS_tree
	__TaskOutputShape__00000007_Dev_UXx_PBS_tree.Task = __Task__00000005_Dev_UXx
	__TaskOutputShape__00000007_Dev_UXx_PBS_tree.Product = __Product__00000004_PBS_tree
	__TaskOutputShape__00000008_Dev_UXx_views.Task = __Task__00000005_Dev_UXx
	__TaskOutputShape__00000008_Dev_UXx_views.Product = __Product__00000005_views
	__TaskOutputShape__00000009_Dev_docx_Backend_Docx_Backend.Task = __Task__00000003_Dev_docx_Backend
	__TaskOutputShape__00000009_Dev_docx_Backend_Docx_Backend.Product = __Product__00000009_Docx_Backend
	__TaskOutputShape__00000010_Dev_docx_Backend_Semantic_Enforcer.Task = __Task__00000003_Dev_docx_Backend
	__TaskOutputShape__00000010_Dev_docx_Backend_Semantic_Enforcer.Product = __Product__00000006_Semantic_Enforcer
	__TaskOutputShape__00000011_Write_Specs_Specifications.Task = __Task__00000006_Write_Specs
	__TaskOutputShape__00000011_Write_Specs_Specifications.Product = __Product__00000010_Specifications
	__TaskOutputShape__00000018_Dev_UXx_WBS_tree.Task = __Task__00000005_Dev_UXx
	__TaskOutputShape__00000018_Dev_UXx_WBS_tree.Product = __Product__00000002_WBS_tree
	__TaskOutputShape__00000019_Dev_UXx_PBS_tree.Task = __Task__00000005_Dev_UXx
	__TaskOutputShape__00000019_Dev_UXx_PBS_tree.Product = __Product__00000004_PBS_tree
	__TaskOutputShape__00000020_Dev_UXx_views.Task = __Task__00000005_Dev_UXx
	__TaskOutputShape__00000020_Dev_UXx_views.Product = __Product__00000005_views
	__TaskOutputShape__00000021_Dev_docx_Backend_Docx_Backend.Task = __Task__00000003_Dev_docx_Backend
	__TaskOutputShape__00000021_Dev_docx_Backend_Docx_Backend.Product = __Product__00000009_Docx_Backend
	__TaskOutputShape__00000022_Dev_docx_Backend_Semantic_Enforcer.Task = __Task__00000003_Dev_docx_Backend
	__TaskOutputShape__00000022_Dev_docx_Backend_Semantic_Enforcer.Product = __Product__00000006_Semantic_Enforcer
	__TaskOutputShape__00000041_Dev_UXx_to_WBS_tree.Task = __Task__00000005_Dev_UXx
	__TaskOutputShape__00000041_Dev_UXx_to_WBS_tree.Product = __Product__00000002_WBS_tree
	__TaskOutputShape__00000042_Dev_UXx_to_PBS_tree.Task = __Task__00000005_Dev_UXx
	__TaskOutputShape__00000042_Dev_UXx_to_PBS_tree.Product = __Product__00000004_PBS_tree
	__TaskOutputShape__00000043_Dev_UXx_to_views.Task = __Task__00000005_Dev_UXx
	__TaskOutputShape__00000043_Dev_UXx_to_views.Product = __Product__00000005_views
	__TaskOutputShape__00000044_Dev_docx_Backend_to_Docx_Backend.Task = __Task__00000003_Dev_docx_Backend
	__TaskOutputShape__00000044_Dev_docx_Backend_to_Docx_Backend.Product = __Product__00000009_Docx_Backend
	__TaskOutputShape__00000045_Dev_docx_Backend_to_Semantic_Enforcer.Task = __Task__00000003_Dev_docx_Backend
	__TaskOutputShape__00000045_Dev_docx_Backend_to_Semantic_Enforcer.Product = __Product__00000006_Semantic_Enforcer
	__TaskOutputShape__00000046_Write_Specs_to_Specifications.Task = __Task__00000006_Write_Specs
	__TaskOutputShape__00000046_Write_Specs_to_Specifications.Product = __Product__00000010_Specifications
	__TaskOutputShape__00000053_Dev_UXx_to_WBS_tree.Task = __Task__00000005_Dev_UXx
	__TaskOutputShape__00000053_Dev_UXx_to_WBS_tree.Product = __Product__00000002_WBS_tree
	__TaskOutputShape__00000054_Dev_UXx_to_PBS_tree.Task = __Task__00000005_Dev_UXx
	__TaskOutputShape__00000054_Dev_UXx_to_PBS_tree.Product = __Product__00000004_PBS_tree
	__TaskOutputShape__00000055_Dev_UXx_to_views.Task = __Task__00000005_Dev_UXx
	__TaskOutputShape__00000055_Dev_UXx_to_views.Product = __Product__00000005_views
	__TaskOutputShape__00000056_Dev_docx_Backend_to_Docx_Backend.Task = __Task__00000003_Dev_docx_Backend
	__TaskOutputShape__00000056_Dev_docx_Backend_to_Docx_Backend.Product = __Product__00000009_Docx_Backend
	__TaskOutputShape__00000057_Dev_docx_Backend_to_Semantic_Enforcer.Task = __Task__00000003_Dev_docx_Backend
	__TaskOutputShape__00000057_Dev_docx_Backend_to_Semantic_Enforcer.Product = __Product__00000006_Semantic_Enforcer
	__TaskOutputShape__00000058_Write_Specs_to_Specifications.Task = __Task__00000006_Write_Specs
	__TaskOutputShape__00000058_Write_Specs_to_Specifications.Product = __Product__00000010_Specifications
	__TaskOutputShape__00000059_Dev_UXx_to_WBS_tree.Task = __Task__00000005_Dev_UXx
	__TaskOutputShape__00000059_Dev_UXx_to_WBS_tree.Product = __Product__00000002_WBS_tree
	__TaskOutputShape__00000060_Dev_UXx_to_PBS_tree.Task = __Task__00000005_Dev_UXx
	__TaskOutputShape__00000060_Dev_UXx_to_PBS_tree.Product = __Product__00000004_PBS_tree
	__TaskOutputShape__00000061_Dev_UXx_to_views.Task = __Task__00000005_Dev_UXx
	__TaskOutputShape__00000061_Dev_UXx_to_views.Product = __Product__00000005_views
	__TaskOutputShape__00000062_Dev_docx_Backend_to_Docx_Backend.Task = __Task__00000003_Dev_docx_Backend
	__TaskOutputShape__00000062_Dev_docx_Backend_to_Docx_Backend.Product = __Product__00000009_Docx_Backend
	__TaskOutputShape__00000063_Dev_docx_Backend_to_Semantic_Enforcer.Task = __Task__00000003_Dev_docx_Backend
	__TaskOutputShape__00000063_Dev_docx_Backend_to_Semantic_Enforcer.Product = __Product__00000006_Semantic_Enforcer
	__TaskOutputShape__00000064_Write_Specs_to_Specifications.Task = __Task__00000006_Write_Specs
	__TaskOutputShape__00000064_Write_Specs_to_Specifications.Product = __Product__00000010_Specifications
	// setup of TaskShape instances pointers
	__TaskShape__00000080_Dev_PBS_Tree.Task = __Task__00000002_Dev_PBS_Tree
	__TaskShape__00000081_Dev_UXx.Task = __Task__00000005_Dev_UXx
	__TaskShape__00000082_Dev_WBS_Tree.Task = __Task__00000001_Dev_WBS_Tree
	__TaskShape__00000083_Dev_docx_Backend.Task = __Task__00000003_Dev_docx_Backend
	__TaskShape__00000084_Dev_views.Task = __Task__00000004_Dev_views
	__TaskShape__00000085_Develop_Backend.Task = __Task__00000000_Develop_Backend
	__TaskShape__00000086_Write_Specs.Task = __Task__00000006_Write_Specs
}

