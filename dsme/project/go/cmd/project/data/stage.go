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

	__Diagram__00000042_NewDiagram := (&models.Diagram{}).Stage(stage)
	__Diagram__00000049_NewDiagram := (&models.Diagram{}).Stage(stage)

	__Product__00000000_UX := (&models.Product{}).Stage(stage)
	__Product__00000001_Backend := (&models.Product{}).Stage(stage)
	__Product__00000002_WBS_tree := (&models.Product{}).Stage(stage)
	__Product__00000004_PBS_tree := (&models.Product{}).Stage(stage)
	__Product__00000005_views := (&models.Product{}).Stage(stage)
	__Product__00000006_Semantic_Enforcer := (&models.Product{}).Stage(stage)
	__Product__00000009_Docx_Backend := (&models.Product{}).Stage(stage)
	__Product__00000010_Specifications := (&models.Product{}).Stage(stage)
	__Product__00000011_Product_1 := (&models.Product{}).Stage(stage)
	__Product__00000013_Product_1_1 := (&models.Product{}).Stage(stage)
	__Product__00000014_Product_1_2 := (&models.Product{}).Stage(stage)
	__Product__00000015_Product_2 := (&models.Product{}).Stage(stage)
	__Product__00000016_Product_1_2_1 := (&models.Product{}).Stage(stage)
	__Product__00000017_Product_2_1 := (&models.Product{}).Stage(stage)
	__Product__00000018_Product_2_2 := (&models.Product{}).Stage(stage)
	__Product__00000019_Probe_Views := (&models.Product{}).Stage(stage)
	__Product__00000020_Application_Views := (&models.Product{}).Stage(stage)

	__ProductCompositionShape__00000204_Product_1_to_Product_1_1 := (&models.ProductCompositionShape{}).Stage(stage)
	__ProductCompositionShape__00000205_Product_1_2_to_Product_1_2_1 := (&models.ProductCompositionShape{}).Stage(stage)
	__ProductCompositionShape__00000206_Product_1_to_Product_1_2 := (&models.ProductCompositionShape{}).Stage(stage)
	__ProductCompositionShape__00000207_Product_2_to_Product_2_1 := (&models.ProductCompositionShape{}).Stage(stage)
	__ProductCompositionShape__00000208_Product_2_to_Product_2_2 := (&models.ProductCompositionShape{}).Stage(stage)
	__ProductCompositionShape__00000263_UX_to_WBS_tree := (&models.ProductCompositionShape{}).Stage(stage)
	__ProductCompositionShape__00000264_UX_to_PBS_tree := (&models.ProductCompositionShape{}).Stage(stage)
	__ProductCompositionShape__00000265_views_to_Probe_Views := (&models.ProductCompositionShape{}).Stage(stage)
	__ProductCompositionShape__00000266_views_to_Application_Views := (&models.ProductCompositionShape{}).Stage(stage)
	__ProductCompositionShape__00000267_UX_to_views := (&models.ProductCompositionShape{}).Stage(stage)
	__ProductCompositionShape__00000268_Backend_to_Semantic_Enforcer := (&models.ProductCompositionShape{}).Stage(stage)
	__ProductCompositionShape__00000269_Backend_to_Docx_Backend := (&models.ProductCompositionShape{}).Stage(stage)

	__ProductShape__00000330_Product_1_NewDiagram := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000331_Product_1_1_NewDiagram := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000332_Product_1_2_NewDiagram := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000333_Product_1_2_1_NewDiagram := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000334_Product_2_NewDiagram := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000335_Product_2_1_NewDiagram := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000336_Product_2_2_NewDiagram := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000415_Specifications_NewDiagram := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000416_UX_NewDiagram := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000417_WBS_tree_NewDiagram := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000418_PBS_tree_NewDiagram := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000419_views_NewDiagram := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000420_Probe_Views_NewDiagram := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000421_Application_Views_NewDiagram := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000422_Backend_NewDiagram := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000423_Semantic_Enforcer_NewDiagram := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000424_Docx_Backend_NewDiagram := (&models.ProductShape{}).Stage(stage)

	__Project__00000000_Project_Editor := (&models.Project{}).Stage(stage)
	__Project__00000001_DSME_Docx := (&models.Project{}).Stage(stage)

	__Root__00000000_Root := (&models.Root{}).Stage(stage)

	__Task__00000000_Develop_Backend := (&models.Task{}).Stage(stage)
	__Task__00000001_Dev_WBS_Tree := (&models.Task{}).Stage(stage)
	__Task__00000002_Dev_PBS_Tree := (&models.Task{}).Stage(stage)
	__Task__00000003_Task_1 := (&models.Task{}).Stage(stage)
	__Task__00000004_Dev_views := (&models.Task{}).Stage(stage)
	__Task__00000005_Dev_UXx := (&models.Task{}).Stage(stage)
	__Task__00000006_Write_Specs := (&models.Task{}).Stage(stage)

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

	// Setup of values

	__Diagram__00000042_NewDiagram.Name = `NewDiagram`
	__Diagram__00000042_NewDiagram.IsChecked = false
	__Diagram__00000042_NewDiagram.IsEditable_ = true
	__Diagram__00000042_NewDiagram.IsInRenameMode = false
	__Diagram__00000042_NewDiagram.ShowPrefix = true
	__Diagram__00000042_NewDiagram.DefaultBoxWidth = 150.000000
	__Diagram__00000042_NewDiagram.DefaultBoxHeigth = 70.000000
	__Diagram__00000042_NewDiagram.IsExpanded = true
	__Diagram__00000042_NewDiagram.ComputedPrefix = ``
	__Diagram__00000042_NewDiagram.IsPBSNodeExpanded = true
	__Diagram__00000042_NewDiagram.IsWBSNodeExpanded = true

	__Diagram__00000049_NewDiagram.Name = `NewDiagram`
	__Diagram__00000049_NewDiagram.IsChecked = true
	__Diagram__00000049_NewDiagram.IsEditable_ = true
	__Diagram__00000049_NewDiagram.IsInRenameMode = false
	__Diagram__00000049_NewDiagram.ShowPrefix = false
	__Diagram__00000049_NewDiagram.DefaultBoxWidth = 150.000000
	__Diagram__00000049_NewDiagram.DefaultBoxHeigth = 70.000000
	__Diagram__00000049_NewDiagram.IsExpanded = true
	__Diagram__00000049_NewDiagram.ComputedPrefix = ``
	__Diagram__00000049_NewDiagram.IsPBSNodeExpanded = true
	__Diagram__00000049_NewDiagram.IsWBSNodeExpanded = false

	__Product__00000000_UX.Name = `UX`
	__Product__00000000_UX.IsExpanded = true
	__Product__00000000_UX.ComputedPrefix = `2`
	__Product__00000000_UX.IsProducersNodeExpanded = false
	__Product__00000000_UX.IsConsumersNodeExpanded = false

	__Product__00000001_Backend.Name = `Backend`
	__Product__00000001_Backend.IsExpanded = true
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
	__Product__00000005_views.IsExpanded = true
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

	__Product__00000011_Product_1.Name = `Product 1`
	__Product__00000011_Product_1.IsExpanded = true
	__Product__00000011_Product_1.ComputedPrefix = `1`
	__Product__00000011_Product_1.IsProducersNodeExpanded = false
	__Product__00000011_Product_1.IsConsumersNodeExpanded = false

	__Product__00000013_Product_1_1.Name = `Product 1.1`
	__Product__00000013_Product_1_1.IsExpanded = false
	__Product__00000013_Product_1_1.ComputedPrefix = `1.1`
	__Product__00000013_Product_1_1.IsProducersNodeExpanded = false
	__Product__00000013_Product_1_1.IsConsumersNodeExpanded = false

	__Product__00000014_Product_1_2.Name = `Product 1.2`
	__Product__00000014_Product_1_2.IsExpanded = true
	__Product__00000014_Product_1_2.ComputedPrefix = `1.2`
	__Product__00000014_Product_1_2.IsProducersNodeExpanded = false
	__Product__00000014_Product_1_2.IsConsumersNodeExpanded = false

	__Product__00000015_Product_2.Name = `Product 2`
	__Product__00000015_Product_2.IsExpanded = true
	__Product__00000015_Product_2.ComputedPrefix = `2`
	__Product__00000015_Product_2.IsProducersNodeExpanded = false
	__Product__00000015_Product_2.IsConsumersNodeExpanded = false

	__Product__00000016_Product_1_2_1.Name = `Product 1.2.1`
	__Product__00000016_Product_1_2_1.IsExpanded = false
	__Product__00000016_Product_1_2_1.ComputedPrefix = `1.2.1`
	__Product__00000016_Product_1_2_1.IsProducersNodeExpanded = false
	__Product__00000016_Product_1_2_1.IsConsumersNodeExpanded = false

	__Product__00000017_Product_2_1.Name = `Product 2.1`
	__Product__00000017_Product_2_1.IsExpanded = false
	__Product__00000017_Product_2_1.ComputedPrefix = `2.1`
	__Product__00000017_Product_2_1.IsProducersNodeExpanded = false
	__Product__00000017_Product_2_1.IsConsumersNodeExpanded = false

	__Product__00000018_Product_2_2.Name = `Product 2.2`
	__Product__00000018_Product_2_2.IsExpanded = false
	__Product__00000018_Product_2_2.ComputedPrefix = `2.2`
	__Product__00000018_Product_2_2.IsProducersNodeExpanded = false
	__Product__00000018_Product_2_2.IsConsumersNodeExpanded = false

	__Product__00000019_Probe_Views.Name = `Probe Views`
	__Product__00000019_Probe_Views.IsExpanded = false
	__Product__00000019_Probe_Views.ComputedPrefix = `2.3.1`
	__Product__00000019_Probe_Views.IsProducersNodeExpanded = false
	__Product__00000019_Probe_Views.IsConsumersNodeExpanded = false

	__Product__00000020_Application_Views.Name = `Application Views`
	__Product__00000020_Application_Views.IsExpanded = false
	__Product__00000020_Application_Views.ComputedPrefix = `2.3.2`
	__Product__00000020_Application_Views.IsProducersNodeExpanded = false
	__Product__00000020_Application_Views.IsConsumersNodeExpanded = false

	__ProductCompositionShape__00000204_Product_1_to_Product_1_1.Name = `Product 1 to Product 1.1`
	__ProductCompositionShape__00000204_Product_1_to_Product_1_1.StartRatio = 0.500000
	__ProductCompositionShape__00000204_Product_1_to_Product_1_1.EndRatio = 0.500000
	__ProductCompositionShape__00000204_Product_1_to_Product_1_1.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000204_Product_1_to_Product_1_1.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000204_Product_1_to_Product_1_1.CornerOffsetRatio = 1.500000

	__ProductCompositionShape__00000205_Product_1_2_to_Product_1_2_1.Name = `Product 1.2 to Product 1.2.1`
	__ProductCompositionShape__00000205_Product_1_2_to_Product_1_2_1.StartRatio = 0.500000
	__ProductCompositionShape__00000205_Product_1_2_to_Product_1_2_1.EndRatio = 0.500000
	__ProductCompositionShape__00000205_Product_1_2_to_Product_1_2_1.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000205_Product_1_2_to_Product_1_2_1.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000205_Product_1_2_to_Product_1_2_1.CornerOffsetRatio = 1.500000

	__ProductCompositionShape__00000206_Product_1_to_Product_1_2.Name = `Product 1 to Product 1.2`
	__ProductCompositionShape__00000206_Product_1_to_Product_1_2.StartRatio = 0.500000
	__ProductCompositionShape__00000206_Product_1_to_Product_1_2.EndRatio = 0.500000
	__ProductCompositionShape__00000206_Product_1_to_Product_1_2.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000206_Product_1_to_Product_1_2.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000206_Product_1_to_Product_1_2.CornerOffsetRatio = 1.500000

	__ProductCompositionShape__00000207_Product_2_to_Product_2_1.Name = `Product 2 to Product 2.1`
	__ProductCompositionShape__00000207_Product_2_to_Product_2_1.StartRatio = 0.500000
	__ProductCompositionShape__00000207_Product_2_to_Product_2_1.EndRatio = 0.500000
	__ProductCompositionShape__00000207_Product_2_to_Product_2_1.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000207_Product_2_to_Product_2_1.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000207_Product_2_to_Product_2_1.CornerOffsetRatio = 1.500000

	__ProductCompositionShape__00000208_Product_2_to_Product_2_2.Name = `Product 2 to Product 2.2`
	__ProductCompositionShape__00000208_Product_2_to_Product_2_2.StartRatio = 0.500000
	__ProductCompositionShape__00000208_Product_2_to_Product_2_2.EndRatio = 0.500000
	__ProductCompositionShape__00000208_Product_2_to_Product_2_2.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000208_Product_2_to_Product_2_2.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000208_Product_2_to_Product_2_2.CornerOffsetRatio = 1.500000

	__ProductCompositionShape__00000263_UX_to_WBS_tree.Name = `UX to WBS tree`
	__ProductCompositionShape__00000263_UX_to_WBS_tree.StartRatio = 0.500000
	__ProductCompositionShape__00000263_UX_to_WBS_tree.EndRatio = 0.500000
	__ProductCompositionShape__00000263_UX_to_WBS_tree.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000263_UX_to_WBS_tree.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000263_UX_to_WBS_tree.CornerOffsetRatio = 1.500000

	__ProductCompositionShape__00000264_UX_to_PBS_tree.Name = `UX to PBS tree`
	__ProductCompositionShape__00000264_UX_to_PBS_tree.StartRatio = 0.500000
	__ProductCompositionShape__00000264_UX_to_PBS_tree.EndRatio = 0.500000
	__ProductCompositionShape__00000264_UX_to_PBS_tree.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000264_UX_to_PBS_tree.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000264_UX_to_PBS_tree.CornerOffsetRatio = 1.500000

	__ProductCompositionShape__00000265_views_to_Probe_Views.Name = `views to Probe Views`
	__ProductCompositionShape__00000265_views_to_Probe_Views.StartRatio = 0.500000
	__ProductCompositionShape__00000265_views_to_Probe_Views.EndRatio = 0.500000
	__ProductCompositionShape__00000265_views_to_Probe_Views.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000265_views_to_Probe_Views.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000265_views_to_Probe_Views.CornerOffsetRatio = 1.500000

	__ProductCompositionShape__00000266_views_to_Application_Views.Name = `views to Application Views`
	__ProductCompositionShape__00000266_views_to_Application_Views.StartRatio = 0.500000
	__ProductCompositionShape__00000266_views_to_Application_Views.EndRatio = 0.500000
	__ProductCompositionShape__00000266_views_to_Application_Views.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000266_views_to_Application_Views.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000266_views_to_Application_Views.CornerOffsetRatio = 1.500000

	__ProductCompositionShape__00000267_UX_to_views.Name = `UX to views`
	__ProductCompositionShape__00000267_UX_to_views.StartRatio = 0.500000
	__ProductCompositionShape__00000267_UX_to_views.EndRatio = 0.500000
	__ProductCompositionShape__00000267_UX_to_views.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000267_UX_to_views.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000267_UX_to_views.CornerOffsetRatio = 1.500000

	__ProductCompositionShape__00000268_Backend_to_Semantic_Enforcer.Name = `Backend to Semantic Enforcer`
	__ProductCompositionShape__00000268_Backend_to_Semantic_Enforcer.StartRatio = 0.500000
	__ProductCompositionShape__00000268_Backend_to_Semantic_Enforcer.EndRatio = 0.500000
	__ProductCompositionShape__00000268_Backend_to_Semantic_Enforcer.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000268_Backend_to_Semantic_Enforcer.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000268_Backend_to_Semantic_Enforcer.CornerOffsetRatio = 1.500000

	__ProductCompositionShape__00000269_Backend_to_Docx_Backend.Name = `Backend to Docx Backend`
	__ProductCompositionShape__00000269_Backend_to_Docx_Backend.StartRatio = 0.500000
	__ProductCompositionShape__00000269_Backend_to_Docx_Backend.EndRatio = 0.500000
	__ProductCompositionShape__00000269_Backend_to_Docx_Backend.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000269_Backend_to_Docx_Backend.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000269_Backend_to_Docx_Backend.CornerOffsetRatio = 1.500000

	__ProductShape__00000330_Product_1_NewDiagram.Name = `Product 1-NewDiagram`
	__ProductShape__00000330_Product_1_NewDiagram.IsExpanded = false
	__ProductShape__00000330_Product_1_NewDiagram.X = 122.500000
	__ProductShape__00000330_Product_1_NewDiagram.Y = 10.000000
	__ProductShape__00000330_Product_1_NewDiagram.Width = 150.000000
	__ProductShape__00000330_Product_1_NewDiagram.Height = 70.000000

	__ProductShape__00000331_Product_1_1_NewDiagram.Name = `Product 1.1-NewDiagram`
	__ProductShape__00000331_Product_1_1_NewDiagram.IsExpanded = false
	__ProductShape__00000331_Product_1_1_NewDiagram.X = 10.000000
	__ProductShape__00000331_Product_1_1_NewDiagram.Y = 150.000000
	__ProductShape__00000331_Product_1_1_NewDiagram.Width = 150.000000
	__ProductShape__00000331_Product_1_1_NewDiagram.Height = 70.000000

	__ProductShape__00000332_Product_1_2_NewDiagram.Name = `Product 1.2-NewDiagram`
	__ProductShape__00000332_Product_1_2_NewDiagram.IsExpanded = false
	__ProductShape__00000332_Product_1_2_NewDiagram.X = 235.000000
	__ProductShape__00000332_Product_1_2_NewDiagram.Y = 150.000000
	__ProductShape__00000332_Product_1_2_NewDiagram.Width = 150.000000
	__ProductShape__00000332_Product_1_2_NewDiagram.Height = 70.000000

	__ProductShape__00000333_Product_1_2_1_NewDiagram.Name = `Product 1.2.1-NewDiagram`
	__ProductShape__00000333_Product_1_2_1_NewDiagram.IsExpanded = false
	__ProductShape__00000333_Product_1_2_1_NewDiagram.X = 235.000000
	__ProductShape__00000333_Product_1_2_1_NewDiagram.Y = 290.000000
	__ProductShape__00000333_Product_1_2_1_NewDiagram.Width = 150.000000
	__ProductShape__00000333_Product_1_2_1_NewDiagram.Height = 70.000000

	__ProductShape__00000334_Product_2_NewDiagram.Name = `Product 2-NewDiagram`
	__ProductShape__00000334_Product_2_NewDiagram.IsExpanded = false
	__ProductShape__00000334_Product_2_NewDiagram.X = 572.500000
	__ProductShape__00000334_Product_2_NewDiagram.Y = 10.000000
	__ProductShape__00000334_Product_2_NewDiagram.Width = 150.000000
	__ProductShape__00000334_Product_2_NewDiagram.Height = 70.000000

	__ProductShape__00000335_Product_2_1_NewDiagram.Name = `Product 2.1-NewDiagram`
	__ProductShape__00000335_Product_2_1_NewDiagram.IsExpanded = false
	__ProductShape__00000335_Product_2_1_NewDiagram.X = 460.000000
	__ProductShape__00000335_Product_2_1_NewDiagram.Y = 150.000000
	__ProductShape__00000335_Product_2_1_NewDiagram.Width = 150.000000
	__ProductShape__00000335_Product_2_1_NewDiagram.Height = 70.000000

	__ProductShape__00000336_Product_2_2_NewDiagram.Name = `Product 2.2-NewDiagram`
	__ProductShape__00000336_Product_2_2_NewDiagram.IsExpanded = false
	__ProductShape__00000336_Product_2_2_NewDiagram.X = 685.000000
	__ProductShape__00000336_Product_2_2_NewDiagram.Y = 150.000000
	__ProductShape__00000336_Product_2_2_NewDiagram.Width = 150.000000
	__ProductShape__00000336_Product_2_2_NewDiagram.Height = 70.000000

	__ProductShape__00000415_Specifications_NewDiagram.Name = `Specifications-NewDiagram`
	__ProductShape__00000415_Specifications_NewDiagram.IsExpanded = false
	__ProductShape__00000415_Specifications_NewDiagram.X = 10.000000
	__ProductShape__00000415_Specifications_NewDiagram.Y = 10.000000
	__ProductShape__00000415_Specifications_NewDiagram.Width = 150.000000
	__ProductShape__00000415_Specifications_NewDiagram.Height = 70.000000

	__ProductShape__00000416_UX_NewDiagram.Name = `UX-NewDiagram`
	__ProductShape__00000416_UX_NewDiagram.IsExpanded = false
	__ProductShape__00000416_UX_NewDiagram.X = 572.500000
	__ProductShape__00000416_UX_NewDiagram.Y = 10.000000
	__ProductShape__00000416_UX_NewDiagram.Width = 150.000000
	__ProductShape__00000416_UX_NewDiagram.Height = 70.000000

	__ProductShape__00000417_WBS_tree_NewDiagram.Name = `WBS tree-NewDiagram`
	__ProductShape__00000417_WBS_tree_NewDiagram.IsExpanded = false
	__ProductShape__00000417_WBS_tree_NewDiagram.X = 235.000000
	__ProductShape__00000417_WBS_tree_NewDiagram.Y = 150.000000
	__ProductShape__00000417_WBS_tree_NewDiagram.Width = 150.000000
	__ProductShape__00000417_WBS_tree_NewDiagram.Height = 70.000000

	__ProductShape__00000418_PBS_tree_NewDiagram.Name = `PBS tree-NewDiagram`
	__ProductShape__00000418_PBS_tree_NewDiagram.IsExpanded = false
	__ProductShape__00000418_PBS_tree_NewDiagram.X = 460.000000
	__ProductShape__00000418_PBS_tree_NewDiagram.Y = 150.000000
	__ProductShape__00000418_PBS_tree_NewDiagram.Width = 150.000000
	__ProductShape__00000418_PBS_tree_NewDiagram.Height = 70.000000

	__ProductShape__00000419_views_NewDiagram.Name = `views-NewDiagram`
	__ProductShape__00000419_views_NewDiagram.IsExpanded = false
	__ProductShape__00000419_views_NewDiagram.X = 797.500000
	__ProductShape__00000419_views_NewDiagram.Y = 150.000000
	__ProductShape__00000419_views_NewDiagram.Width = 150.000000
	__ProductShape__00000419_views_NewDiagram.Height = 70.000000

	__ProductShape__00000420_Probe_Views_NewDiagram.Name = `Probe Views-NewDiagram`
	__ProductShape__00000420_Probe_Views_NewDiagram.IsExpanded = false
	__ProductShape__00000420_Probe_Views_NewDiagram.X = 685.000000
	__ProductShape__00000420_Probe_Views_NewDiagram.Y = 290.000000
	__ProductShape__00000420_Probe_Views_NewDiagram.Width = 150.000000
	__ProductShape__00000420_Probe_Views_NewDiagram.Height = 70.000000

	__ProductShape__00000421_Application_Views_NewDiagram.Name = `Application Views-NewDiagram`
	__ProductShape__00000421_Application_Views_NewDiagram.IsExpanded = false
	__ProductShape__00000421_Application_Views_NewDiagram.X = 910.000000
	__ProductShape__00000421_Application_Views_NewDiagram.Y = 290.000000
	__ProductShape__00000421_Application_Views_NewDiagram.Width = 150.000000
	__ProductShape__00000421_Application_Views_NewDiagram.Height = 70.000000

	__ProductShape__00000422_Backend_NewDiagram.Name = `Backend-NewDiagram`
	__ProductShape__00000422_Backend_NewDiagram.IsExpanded = false
	__ProductShape__00000422_Backend_NewDiagram.X = 1247.500000
	__ProductShape__00000422_Backend_NewDiagram.Y = 10.000000
	__ProductShape__00000422_Backend_NewDiagram.Width = 150.000000
	__ProductShape__00000422_Backend_NewDiagram.Height = 70.000000

	__ProductShape__00000423_Semantic_Enforcer_NewDiagram.Name = `Semantic Enforcer-NewDiagram`
	__ProductShape__00000423_Semantic_Enforcer_NewDiagram.IsExpanded = false
	__ProductShape__00000423_Semantic_Enforcer_NewDiagram.X = 1135.000000
	__ProductShape__00000423_Semantic_Enforcer_NewDiagram.Y = 150.000000
	__ProductShape__00000423_Semantic_Enforcer_NewDiagram.Width = 150.000000
	__ProductShape__00000423_Semantic_Enforcer_NewDiagram.Height = 70.000000

	__ProductShape__00000424_Docx_Backend_NewDiagram.Name = `Docx Backend-NewDiagram`
	__ProductShape__00000424_Docx_Backend_NewDiagram.IsExpanded = false
	__ProductShape__00000424_Docx_Backend_NewDiagram.X = 1360.000000
	__ProductShape__00000424_Docx_Backend_NewDiagram.Y = 150.000000
	__ProductShape__00000424_Docx_Backend_NewDiagram.Width = 150.000000
	__ProductShape__00000424_Docx_Backend_NewDiagram.Height = 70.000000

	__Project__00000000_Project_Editor.Name = `Project Editor`
	__Project__00000000_Project_Editor.IsPBSNodeExpanded = true
	__Project__00000000_Project_Editor.IsWBSNodeExpanded = false
	__Project__00000000_Project_Editor.IsDiagramsNodeExpanded = true
	__Project__00000000_Project_Editor.IsExpanded = false
	__Project__00000000_Project_Editor.ComputedPrefix = ``

	__Project__00000001_DSME_Docx.Name = `DSME Docx`
	__Project__00000001_DSME_Docx.IsPBSNodeExpanded = false
	__Project__00000001_DSME_Docx.IsWBSNodeExpanded = false
	__Project__00000001_DSME_Docx.IsDiagramsNodeExpanded = true
	__Project__00000001_DSME_Docx.IsExpanded = false
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

	__Task__00000003_Task_1.Name = `Task 1`
	__Task__00000003_Task_1.IsExpanded = true
	__Task__00000003_Task_1.ComputedPrefix = `1`
	__Task__00000003_Task_1.IsInputsNodeExpanded = true
	__Task__00000003_Task_1.IsOutputsNodeExpanded = true

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

	// Setup of pointers
	// setup of Diagram instances pointers
	__Diagram__00000042_NewDiagram.Product_Shapes = append(__Diagram__00000042_NewDiagram.Product_Shapes, __ProductShape__00000330_Product_1_NewDiagram)
	__Diagram__00000042_NewDiagram.Product_Shapes = append(__Diagram__00000042_NewDiagram.Product_Shapes, __ProductShape__00000331_Product_1_1_NewDiagram)
	__Diagram__00000042_NewDiagram.Product_Shapes = append(__Diagram__00000042_NewDiagram.Product_Shapes, __ProductShape__00000332_Product_1_2_NewDiagram)
	__Diagram__00000042_NewDiagram.Product_Shapes = append(__Diagram__00000042_NewDiagram.Product_Shapes, __ProductShape__00000333_Product_1_2_1_NewDiagram)
	__Diagram__00000042_NewDiagram.Product_Shapes = append(__Diagram__00000042_NewDiagram.Product_Shapes, __ProductShape__00000334_Product_2_NewDiagram)
	__Diagram__00000042_NewDiagram.Product_Shapes = append(__Diagram__00000042_NewDiagram.Product_Shapes, __ProductShape__00000335_Product_2_1_NewDiagram)
	__Diagram__00000042_NewDiagram.Product_Shapes = append(__Diagram__00000042_NewDiagram.Product_Shapes, __ProductShape__00000336_Product_2_2_NewDiagram)
	__Diagram__00000042_NewDiagram.ProductsWhoseNodeIsExpanded = append(__Diagram__00000042_NewDiagram.ProductsWhoseNodeIsExpanded, __Product__00000015_Product_2)
	__Diagram__00000042_NewDiagram.ProductComposition_Shapes = append(__Diagram__00000042_NewDiagram.ProductComposition_Shapes, __ProductCompositionShape__00000204_Product_1_to_Product_1_1)
	__Diagram__00000042_NewDiagram.ProductComposition_Shapes = append(__Diagram__00000042_NewDiagram.ProductComposition_Shapes, __ProductCompositionShape__00000205_Product_1_2_to_Product_1_2_1)
	__Diagram__00000042_NewDiagram.ProductComposition_Shapes = append(__Diagram__00000042_NewDiagram.ProductComposition_Shapes, __ProductCompositionShape__00000206_Product_1_to_Product_1_2)
	__Diagram__00000042_NewDiagram.ProductComposition_Shapes = append(__Diagram__00000042_NewDiagram.ProductComposition_Shapes, __ProductCompositionShape__00000207_Product_2_to_Product_2_1)
	__Diagram__00000042_NewDiagram.ProductComposition_Shapes = append(__Diagram__00000042_NewDiagram.ProductComposition_Shapes, __ProductCompositionShape__00000208_Product_2_to_Product_2_2)
	__Diagram__00000049_NewDiagram.Product_Shapes = append(__Diagram__00000049_NewDiagram.Product_Shapes, __ProductShape__00000415_Specifications_NewDiagram)
	__Diagram__00000049_NewDiagram.Product_Shapes = append(__Diagram__00000049_NewDiagram.Product_Shapes, __ProductShape__00000416_UX_NewDiagram)
	__Diagram__00000049_NewDiagram.Product_Shapes = append(__Diagram__00000049_NewDiagram.Product_Shapes, __ProductShape__00000417_WBS_tree_NewDiagram)
	__Diagram__00000049_NewDiagram.Product_Shapes = append(__Diagram__00000049_NewDiagram.Product_Shapes, __ProductShape__00000418_PBS_tree_NewDiagram)
	__Diagram__00000049_NewDiagram.Product_Shapes = append(__Diagram__00000049_NewDiagram.Product_Shapes, __ProductShape__00000419_views_NewDiagram)
	__Diagram__00000049_NewDiagram.Product_Shapes = append(__Diagram__00000049_NewDiagram.Product_Shapes, __ProductShape__00000420_Probe_Views_NewDiagram)
	__Diagram__00000049_NewDiagram.Product_Shapes = append(__Diagram__00000049_NewDiagram.Product_Shapes, __ProductShape__00000421_Application_Views_NewDiagram)
	__Diagram__00000049_NewDiagram.Product_Shapes = append(__Diagram__00000049_NewDiagram.Product_Shapes, __ProductShape__00000422_Backend_NewDiagram)
	__Diagram__00000049_NewDiagram.Product_Shapes = append(__Diagram__00000049_NewDiagram.Product_Shapes, __ProductShape__00000423_Semantic_Enforcer_NewDiagram)
	__Diagram__00000049_NewDiagram.Product_Shapes = append(__Diagram__00000049_NewDiagram.Product_Shapes, __ProductShape__00000424_Docx_Backend_NewDiagram)
	__Diagram__00000049_NewDiagram.ProductsWhoseNodeIsExpanded = append(__Diagram__00000049_NewDiagram.ProductsWhoseNodeIsExpanded, __Product__00000000_UX)
	__Diagram__00000049_NewDiagram.ProductsWhoseNodeIsExpanded = append(__Diagram__00000049_NewDiagram.ProductsWhoseNodeIsExpanded, __Product__00000005_views)
	__Diagram__00000049_NewDiagram.ProductComposition_Shapes = append(__Diagram__00000049_NewDiagram.ProductComposition_Shapes, __ProductCompositionShape__00000263_UX_to_WBS_tree)
	__Diagram__00000049_NewDiagram.ProductComposition_Shapes = append(__Diagram__00000049_NewDiagram.ProductComposition_Shapes, __ProductCompositionShape__00000264_UX_to_PBS_tree)
	__Diagram__00000049_NewDiagram.ProductComposition_Shapes = append(__Diagram__00000049_NewDiagram.ProductComposition_Shapes, __ProductCompositionShape__00000265_views_to_Probe_Views)
	__Diagram__00000049_NewDiagram.ProductComposition_Shapes = append(__Diagram__00000049_NewDiagram.ProductComposition_Shapes, __ProductCompositionShape__00000266_views_to_Application_Views)
	__Diagram__00000049_NewDiagram.ProductComposition_Shapes = append(__Diagram__00000049_NewDiagram.ProductComposition_Shapes, __ProductCompositionShape__00000267_UX_to_views)
	__Diagram__00000049_NewDiagram.ProductComposition_Shapes = append(__Diagram__00000049_NewDiagram.ProductComposition_Shapes, __ProductCompositionShape__00000268_Backend_to_Semantic_Enforcer)
	__Diagram__00000049_NewDiagram.ProductComposition_Shapes = append(__Diagram__00000049_NewDiagram.ProductComposition_Shapes, __ProductCompositionShape__00000269_Backend_to_Docx_Backend)
	// setup of Product instances pointers
	__Product__00000000_UX.SubProducts = append(__Product__00000000_UX.SubProducts, __Product__00000002_WBS_tree)
	__Product__00000000_UX.SubProducts = append(__Product__00000000_UX.SubProducts, __Product__00000004_PBS_tree)
	__Product__00000000_UX.SubProducts = append(__Product__00000000_UX.SubProducts, __Product__00000005_views)
	__Product__00000001_Backend.SubProducts = append(__Product__00000001_Backend.SubProducts, __Product__00000006_Semantic_Enforcer)
	__Product__00000001_Backend.SubProducts = append(__Product__00000001_Backend.SubProducts, __Product__00000009_Docx_Backend)
	__Product__00000005_views.SubProducts = append(__Product__00000005_views.SubProducts, __Product__00000019_Probe_Views)
	__Product__00000005_views.SubProducts = append(__Product__00000005_views.SubProducts, __Product__00000020_Application_Views)
	__Product__00000011_Product_1.SubProducts = append(__Product__00000011_Product_1.SubProducts, __Product__00000013_Product_1_1)
	__Product__00000011_Product_1.SubProducts = append(__Product__00000011_Product_1.SubProducts, __Product__00000014_Product_1_2)
	__Product__00000014_Product_1_2.SubProducts = append(__Product__00000014_Product_1_2.SubProducts, __Product__00000016_Product_1_2_1)
	__Product__00000015_Product_2.SubProducts = append(__Product__00000015_Product_2.SubProducts, __Product__00000017_Product_2_1)
	__Product__00000015_Product_2.SubProducts = append(__Product__00000015_Product_2.SubProducts, __Product__00000018_Product_2_2)
	// setup of ProductCompositionShape instances pointers
	__ProductCompositionShape__00000204_Product_1_to_Product_1_1.Product = __Product__00000013_Product_1_1
	__ProductCompositionShape__00000205_Product_1_2_to_Product_1_2_1.Product = __Product__00000016_Product_1_2_1
	__ProductCompositionShape__00000206_Product_1_to_Product_1_2.Product = __Product__00000014_Product_1_2
	__ProductCompositionShape__00000207_Product_2_to_Product_2_1.Product = __Product__00000017_Product_2_1
	__ProductCompositionShape__00000208_Product_2_to_Product_2_2.Product = __Product__00000018_Product_2_2
	__ProductCompositionShape__00000263_UX_to_WBS_tree.Product = __Product__00000002_WBS_tree
	__ProductCompositionShape__00000264_UX_to_PBS_tree.Product = __Product__00000004_PBS_tree
	__ProductCompositionShape__00000265_views_to_Probe_Views.Product = __Product__00000019_Probe_Views
	__ProductCompositionShape__00000266_views_to_Application_Views.Product = __Product__00000020_Application_Views
	__ProductCompositionShape__00000267_UX_to_views.Product = __Product__00000005_views
	__ProductCompositionShape__00000268_Backend_to_Semantic_Enforcer.Product = __Product__00000006_Semantic_Enforcer
	__ProductCompositionShape__00000269_Backend_to_Docx_Backend.Product = __Product__00000009_Docx_Backend
	// setup of ProductShape instances pointers
	__ProductShape__00000330_Product_1_NewDiagram.Product = __Product__00000011_Product_1
	__ProductShape__00000331_Product_1_1_NewDiagram.Product = __Product__00000013_Product_1_1
	__ProductShape__00000332_Product_1_2_NewDiagram.Product = __Product__00000014_Product_1_2
	__ProductShape__00000333_Product_1_2_1_NewDiagram.Product = __Product__00000016_Product_1_2_1
	__ProductShape__00000334_Product_2_NewDiagram.Product = __Product__00000015_Product_2
	__ProductShape__00000335_Product_2_1_NewDiagram.Product = __Product__00000017_Product_2_1
	__ProductShape__00000336_Product_2_2_NewDiagram.Product = __Product__00000018_Product_2_2
	__ProductShape__00000415_Specifications_NewDiagram.Product = __Product__00000010_Specifications
	__ProductShape__00000416_UX_NewDiagram.Product = __Product__00000000_UX
	__ProductShape__00000417_WBS_tree_NewDiagram.Product = __Product__00000002_WBS_tree
	__ProductShape__00000418_PBS_tree_NewDiagram.Product = __Product__00000004_PBS_tree
	__ProductShape__00000419_views_NewDiagram.Product = __Product__00000005_views
	__ProductShape__00000420_Probe_Views_NewDiagram.Product = __Product__00000019_Probe_Views
	__ProductShape__00000421_Application_Views_NewDiagram.Product = __Product__00000020_Application_Views
	__ProductShape__00000422_Backend_NewDiagram.Product = __Product__00000001_Backend
	__ProductShape__00000423_Semantic_Enforcer_NewDiagram.Product = __Product__00000006_Semantic_Enforcer
	__ProductShape__00000424_Docx_Backend_NewDiagram.Product = __Product__00000009_Docx_Backend
	// setup of Project instances pointers
	__Project__00000000_Project_Editor.RootProducts = append(__Project__00000000_Project_Editor.RootProducts, __Product__00000010_Specifications)
	__Project__00000000_Project_Editor.RootProducts = append(__Project__00000000_Project_Editor.RootProducts, __Product__00000000_UX)
	__Project__00000000_Project_Editor.RootProducts = append(__Project__00000000_Project_Editor.RootProducts, __Product__00000001_Backend)
	__Project__00000000_Project_Editor.RootTasks = append(__Project__00000000_Project_Editor.RootTasks, __Task__00000006_Write_Specs)
	__Project__00000000_Project_Editor.RootTasks = append(__Project__00000000_Project_Editor.RootTasks, __Task__00000000_Develop_Backend)
	__Project__00000000_Project_Editor.RootTasks = append(__Project__00000000_Project_Editor.RootTasks, __Task__00000001_Dev_WBS_Tree)
	__Project__00000000_Project_Editor.RootTasks = append(__Project__00000000_Project_Editor.RootTasks, __Task__00000002_Dev_PBS_Tree)
	__Project__00000000_Project_Editor.Diagrams = append(__Project__00000000_Project_Editor.Diagrams, __Diagram__00000049_NewDiagram)
	__Project__00000001_DSME_Docx.RootProducts = append(__Project__00000001_DSME_Docx.RootProducts, __Product__00000011_Product_1)
	__Project__00000001_DSME_Docx.RootProducts = append(__Project__00000001_DSME_Docx.RootProducts, __Product__00000015_Product_2)
	__Project__00000001_DSME_Docx.RootTasks = append(__Project__00000001_DSME_Docx.RootTasks, __Task__00000003_Task_1)
	__Project__00000001_DSME_Docx.Diagrams = append(__Project__00000001_DSME_Docx.Diagrams, __Diagram__00000042_NewDiagram)
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
	// setup of TaskOutputShape instances pointers
	__TaskOutputShape__00000006_Dev_UXx_WBS_tree.Task = __Task__00000005_Dev_UXx
	__TaskOutputShape__00000006_Dev_UXx_WBS_tree.Product = __Product__00000002_WBS_tree
	__TaskOutputShape__00000007_Dev_UXx_PBS_tree.Task = __Task__00000005_Dev_UXx
	__TaskOutputShape__00000007_Dev_UXx_PBS_tree.Product = __Product__00000004_PBS_tree
	__TaskOutputShape__00000008_Dev_UXx_views.Task = __Task__00000005_Dev_UXx
	__TaskOutputShape__00000008_Dev_UXx_views.Product = __Product__00000005_views
	__TaskOutputShape__00000009_Dev_docx_Backend_Docx_Backend.Task = __Task__00000003_Task_1
	__TaskOutputShape__00000009_Dev_docx_Backend_Docx_Backend.Product = __Product__00000009_Docx_Backend
	__TaskOutputShape__00000010_Dev_docx_Backend_Semantic_Enforcer.Task = __Task__00000003_Task_1
	__TaskOutputShape__00000010_Dev_docx_Backend_Semantic_Enforcer.Product = __Product__00000006_Semantic_Enforcer
	__TaskOutputShape__00000011_Write_Specs_Specifications.Task = __Task__00000006_Write_Specs
	__TaskOutputShape__00000011_Write_Specs_Specifications.Product = __Product__00000010_Specifications
	__TaskOutputShape__00000018_Dev_UXx_WBS_tree.Task = __Task__00000005_Dev_UXx
	__TaskOutputShape__00000018_Dev_UXx_WBS_tree.Product = __Product__00000002_WBS_tree
	__TaskOutputShape__00000019_Dev_UXx_PBS_tree.Task = __Task__00000005_Dev_UXx
	__TaskOutputShape__00000019_Dev_UXx_PBS_tree.Product = __Product__00000004_PBS_tree
	__TaskOutputShape__00000020_Dev_UXx_views.Task = __Task__00000005_Dev_UXx
	__TaskOutputShape__00000020_Dev_UXx_views.Product = __Product__00000005_views
	__TaskOutputShape__00000021_Dev_docx_Backend_Docx_Backend.Task = __Task__00000003_Task_1
	__TaskOutputShape__00000021_Dev_docx_Backend_Docx_Backend.Product = __Product__00000009_Docx_Backend
	__TaskOutputShape__00000022_Dev_docx_Backend_Semantic_Enforcer.Task = __Task__00000003_Task_1
	__TaskOutputShape__00000022_Dev_docx_Backend_Semantic_Enforcer.Product = __Product__00000006_Semantic_Enforcer
	__TaskOutputShape__00000041_Dev_UXx_to_WBS_tree.Task = __Task__00000005_Dev_UXx
	__TaskOutputShape__00000041_Dev_UXx_to_WBS_tree.Product = __Product__00000002_WBS_tree
	__TaskOutputShape__00000042_Dev_UXx_to_PBS_tree.Task = __Task__00000005_Dev_UXx
	__TaskOutputShape__00000042_Dev_UXx_to_PBS_tree.Product = __Product__00000004_PBS_tree
	__TaskOutputShape__00000043_Dev_UXx_to_views.Task = __Task__00000005_Dev_UXx
	__TaskOutputShape__00000043_Dev_UXx_to_views.Product = __Product__00000005_views
	__TaskOutputShape__00000044_Dev_docx_Backend_to_Docx_Backend.Task = __Task__00000003_Task_1
	__TaskOutputShape__00000044_Dev_docx_Backend_to_Docx_Backend.Product = __Product__00000009_Docx_Backend
	__TaskOutputShape__00000045_Dev_docx_Backend_to_Semantic_Enforcer.Task = __Task__00000003_Task_1
	__TaskOutputShape__00000045_Dev_docx_Backend_to_Semantic_Enforcer.Product = __Product__00000006_Semantic_Enforcer
	__TaskOutputShape__00000046_Write_Specs_to_Specifications.Task = __Task__00000006_Write_Specs
	__TaskOutputShape__00000046_Write_Specs_to_Specifications.Product = __Product__00000010_Specifications
	__TaskOutputShape__00000053_Dev_UXx_to_WBS_tree.Task = __Task__00000005_Dev_UXx
	__TaskOutputShape__00000053_Dev_UXx_to_WBS_tree.Product = __Product__00000002_WBS_tree
	__TaskOutputShape__00000054_Dev_UXx_to_PBS_tree.Task = __Task__00000005_Dev_UXx
	__TaskOutputShape__00000054_Dev_UXx_to_PBS_tree.Product = __Product__00000004_PBS_tree
	__TaskOutputShape__00000055_Dev_UXx_to_views.Task = __Task__00000005_Dev_UXx
	__TaskOutputShape__00000055_Dev_UXx_to_views.Product = __Product__00000005_views
	__TaskOutputShape__00000056_Dev_docx_Backend_to_Docx_Backend.Task = __Task__00000003_Task_1
	__TaskOutputShape__00000056_Dev_docx_Backend_to_Docx_Backend.Product = __Product__00000009_Docx_Backend
	__TaskOutputShape__00000057_Dev_docx_Backend_to_Semantic_Enforcer.Task = __Task__00000003_Task_1
	__TaskOutputShape__00000057_Dev_docx_Backend_to_Semantic_Enforcer.Product = __Product__00000006_Semantic_Enforcer
	__TaskOutputShape__00000058_Write_Specs_to_Specifications.Task = __Task__00000006_Write_Specs
	__TaskOutputShape__00000058_Write_Specs_to_Specifications.Product = __Product__00000010_Specifications
}

