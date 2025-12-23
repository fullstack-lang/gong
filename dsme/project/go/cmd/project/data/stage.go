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

	__CompositionShape__00000011_Backend_to_Semantic_Enforcer := (&models.CompositionShape{}).Stage(stage)
	__CompositionShape__00000012_Backend_to_Docx_Backend := (&models.CompositionShape{}).Stage(stage)
	__CompositionShape__00000013_UX_to_WBS_tree := (&models.CompositionShape{}).Stage(stage)
	__CompositionShape__00000014_UX_to_PBS_tree := (&models.CompositionShape{}).Stage(stage)
	__CompositionShape__00000017_UX_to_views := (&models.CompositionShape{}).Stage(stage)
	__CompositionShape__00000019_UX_to_PBS_tree := (&models.CompositionShape{}).Stage(stage)
	__CompositionShape__00000020_UX_to_views := (&models.CompositionShape{}).Stage(stage)
	__CompositionShape__00000021_UX_to_WBS_tree := (&models.CompositionShape{}).Stage(stage)

	__Product__00000000_UX := (&models.Product{}).Stage(stage)
	__Product__00000001_Backend := (&models.Product{}).Stage(stage)
	__Product__00000002_WBS_tree := (&models.Product{}).Stage(stage)
	__Product__00000004_PBS_tree := (&models.Product{}).Stage(stage)
	__Product__00000005_views := (&models.Product{}).Stage(stage)
	__Product__00000006_Semantic_Enforcer := (&models.Product{}).Stage(stage)
	__Product__00000009_Docx_Backend := (&models.Product{}).Stage(stage)

	__ProductShape__00000000_UX_diagram_1 := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000001_WBS_tree_diagram_1 := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000002_PBS_tree_diagram_1 := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000005_Semantic_Enforcer_diagram_1 := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000007_Docx_Backend_diagram_1 := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000008_Backend_diagram_1 := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000010_views_diagram_1 := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000011_NewProduct_diagram_1 := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000012_UX_diagram_2 := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000013_WBS_tree_diagram_2 := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000022_UX_NewDiagram := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000024_PBS_tree_NewDiagram := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000025_views_NewDiagram := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000026_WBS_tree_NewDiagram := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000027_Backend_NewDiagram := (&models.ProductShape{}).Stage(stage)

	__Project__00000000_Project_Editor := (&models.Project{}).Stage(stage)
	__Project__00000001_DSME_Docx := (&models.Project{}).Stage(stage)

	__Root__00000000_Root := (&models.Root{}).Stage(stage)

	__Task__00000000_Develop_Backend := (&models.Task{}).Stage(stage)
	__Task__00000001_Dev_WBS_Tree := (&models.Task{}).Stage(stage)
	__Task__00000002_Dev_PBS_Tree := (&models.Task{}).Stage(stage)
	__Task__00000003_Dev_docx_Backend := (&models.Task{}).Stage(stage)
	__Task__00000004_Dev_views := (&models.Task{}).Stage(stage)
	__Task__00000005_Dev_UXx := (&models.Task{}).Stage(stage)

	// Setup of values

	__CompositionShape__00000011_Backend_to_Semantic_Enforcer.Name = `Backend to Semantic Enforcer`
	__CompositionShape__00000011_Backend_to_Semantic_Enforcer.StartRatio = 0.500000
	__CompositionShape__00000011_Backend_to_Semantic_Enforcer.EndRatio = 0.500000
	__CompositionShape__00000011_Backend_to_Semantic_Enforcer.StartOrientation = models.ORIENTATION_VERTICAL
	__CompositionShape__00000011_Backend_to_Semantic_Enforcer.EndOrientation = models.ORIENTATION_VERTICAL
	__CompositionShape__00000011_Backend_to_Semantic_Enforcer.CornerOffsetRatio = 2.000000

	__CompositionShape__00000012_Backend_to_Docx_Backend.Name = `Backend to Docx Backend`
	__CompositionShape__00000012_Backend_to_Docx_Backend.StartRatio = 0.500000
	__CompositionShape__00000012_Backend_to_Docx_Backend.EndRatio = 0.500000
	__CompositionShape__00000012_Backend_to_Docx_Backend.StartOrientation = models.ORIENTATION_VERTICAL
	__CompositionShape__00000012_Backend_to_Docx_Backend.EndOrientation = models.ORIENTATION_VERTICAL
	__CompositionShape__00000012_Backend_to_Docx_Backend.CornerOffsetRatio = 2.000000

	__CompositionShape__00000013_UX_to_WBS_tree.Name = `UX to WBS tree`
	__CompositionShape__00000013_UX_to_WBS_tree.StartRatio = 0.500000
	__CompositionShape__00000013_UX_to_WBS_tree.EndRatio = 0.500000
	__CompositionShape__00000013_UX_to_WBS_tree.StartOrientation = models.ORIENTATION_VERTICAL
	__CompositionShape__00000013_UX_to_WBS_tree.EndOrientation = models.ORIENTATION_VERTICAL
	__CompositionShape__00000013_UX_to_WBS_tree.CornerOffsetRatio = 1.680000

	__CompositionShape__00000014_UX_to_PBS_tree.Name = `UX to PBS tree`
	__CompositionShape__00000014_UX_to_PBS_tree.StartRatio = 0.500000
	__CompositionShape__00000014_UX_to_PBS_tree.EndRatio = 0.500000
	__CompositionShape__00000014_UX_to_PBS_tree.StartOrientation = models.ORIENTATION_VERTICAL
	__CompositionShape__00000014_UX_to_PBS_tree.EndOrientation = models.ORIENTATION_VERTICAL
	__CompositionShape__00000014_UX_to_PBS_tree.CornerOffsetRatio = 1.680000

	__CompositionShape__00000017_UX_to_views.Name = `UX to views`
	__CompositionShape__00000017_UX_to_views.StartRatio = 0.500000
	__CompositionShape__00000017_UX_to_views.EndRatio = 0.500000
	__CompositionShape__00000017_UX_to_views.StartOrientation = models.ORIENTATION_VERTICAL
	__CompositionShape__00000017_UX_to_views.EndOrientation = models.ORIENTATION_VERTICAL
	__CompositionShape__00000017_UX_to_views.CornerOffsetRatio = 1.680000

	__CompositionShape__00000019_UX_to_PBS_tree.Name = `UX to PBS tree`
	__CompositionShape__00000019_UX_to_PBS_tree.StartRatio = 0.500000
	__CompositionShape__00000019_UX_to_PBS_tree.EndRatio = 0.500000
	__CompositionShape__00000019_UX_to_PBS_tree.StartOrientation = models.ORIENTATION_VERTICAL
	__CompositionShape__00000019_UX_to_PBS_tree.EndOrientation = models.ORIENTATION_VERTICAL
	__CompositionShape__00000019_UX_to_PBS_tree.CornerOffsetRatio = 1.680000

	__CompositionShape__00000020_UX_to_views.Name = `UX to views`
	__CompositionShape__00000020_UX_to_views.StartRatio = 0.500000
	__CompositionShape__00000020_UX_to_views.EndRatio = 0.500000
	__CompositionShape__00000020_UX_to_views.StartOrientation = models.ORIENTATION_VERTICAL
	__CompositionShape__00000020_UX_to_views.EndOrientation = models.ORIENTATION_VERTICAL
	__CompositionShape__00000020_UX_to_views.CornerOffsetRatio = 1.680000

	__CompositionShape__00000021_UX_to_WBS_tree.Name = `UX to WBS tree`
	__CompositionShape__00000021_UX_to_WBS_tree.StartRatio = 0.500000
	__CompositionShape__00000021_UX_to_WBS_tree.EndRatio = 0.500000
	__CompositionShape__00000021_UX_to_WBS_tree.StartOrientation = models.ORIENTATION_VERTICAL
	__CompositionShape__00000021_UX_to_WBS_tree.EndOrientation = models.ORIENTATION_VERTICAL
	__CompositionShape__00000021_UX_to_WBS_tree.CornerOffsetRatio = 1.680000

	__Product__00000000_UX.Name = `UX`
	__Product__00000000_UX.IsExpanded = true
	__Product__00000000_UX.ComputedPrefix = `1`
	__Product__00000000_UX.IsProducersNodeExpanded = false
	__Product__00000000_UX.IsConsumersNodeExpanded = false

	__Product__00000001_Backend.Name = `Backend`
	__Product__00000001_Backend.IsExpanded = false
	__Product__00000001_Backend.ComputedPrefix = `2`
	__Product__00000001_Backend.IsProducersNodeExpanded = false
	__Product__00000001_Backend.IsConsumersNodeExpanded = false

	__Product__00000002_WBS_tree.Name = `WBS tree`
	__Product__00000002_WBS_tree.IsExpanded = false
	__Product__00000002_WBS_tree.ComputedPrefix = `1.1`
	__Product__00000002_WBS_tree.IsProducersNodeExpanded = false
	__Product__00000002_WBS_tree.IsConsumersNodeExpanded = false

	__Product__00000004_PBS_tree.Name = `PBS tree`
	__Product__00000004_PBS_tree.IsExpanded = false
	__Product__00000004_PBS_tree.ComputedPrefix = `1.2`
	__Product__00000004_PBS_tree.IsProducersNodeExpanded = false
	__Product__00000004_PBS_tree.IsConsumersNodeExpanded = false

	__Product__00000005_views.Name = `views`
	__Product__00000005_views.IsExpanded = true
	__Product__00000005_views.ComputedPrefix = `1.3`
	__Product__00000005_views.IsProducersNodeExpanded = false
	__Product__00000005_views.IsConsumersNodeExpanded = false

	__Product__00000006_Semantic_Enforcer.Name = `Semantic Enforcer`
	__Product__00000006_Semantic_Enforcer.IsExpanded = false
	__Product__00000006_Semantic_Enforcer.ComputedPrefix = `2.1`
	__Product__00000006_Semantic_Enforcer.IsProducersNodeExpanded = false
	__Product__00000006_Semantic_Enforcer.IsConsumersNodeExpanded = false

	__Product__00000009_Docx_Backend.Name = `Docx Backend`
	__Product__00000009_Docx_Backend.IsExpanded = false
	__Product__00000009_Docx_Backend.ComputedPrefix = `2.2`
	__Product__00000009_Docx_Backend.IsProducersNodeExpanded = false
	__Product__00000009_Docx_Backend.IsConsumersNodeExpanded = false

	__ProductShape__00000000_UX_diagram_1.Name = `UX-diagram 1`
	__ProductShape__00000000_UX_diagram_1.IsExpanded = false
	__ProductShape__00000000_UX_diagram_1.X = 72.000000
	__ProductShape__00000000_UX_diagram_1.Y = 134.000000
	__ProductShape__00000000_UX_diagram_1.Width = 810.000000
	__ProductShape__00000000_UX_diagram_1.Height = 66.000015

	__ProductShape__00000001_WBS_tree_diagram_1.Name = `WBS tree-diagram 1`
	__ProductShape__00000001_WBS_tree_diagram_1.IsExpanded = false
	__ProductShape__00000001_WBS_tree_diagram_1.X = 666.000000
	__ProductShape__00000001_WBS_tree_diagram_1.Y = 344.999985
	__ProductShape__00000001_WBS_tree_diagram_1.Width = 200.000000
	__ProductShape__00000001_WBS_tree_diagram_1.Height = 80.000000

	__ProductShape__00000002_PBS_tree_diagram_1.Name = `PBS tree-diagram 1`
	__ProductShape__00000002_PBS_tree_diagram_1.IsExpanded = false
	__ProductShape__00000002_PBS_tree_diagram_1.X = 384.000000
	__ProductShape__00000002_PBS_tree_diagram_1.Y = 353.999985
	__ProductShape__00000002_PBS_tree_diagram_1.Width = 200.000000
	__ProductShape__00000002_PBS_tree_diagram_1.Height = 80.000000

	__ProductShape__00000005_Semantic_Enforcer_diagram_1.Name = `Semantic Enforcer-diagram 1`
	__ProductShape__00000005_Semantic_Enforcer_diagram_1.IsExpanded = false
	__ProductShape__00000005_Semantic_Enforcer_diagram_1.X = 95.000000
	__ProductShape__00000005_Semantic_Enforcer_diagram_1.Y = 708.999985
	__ProductShape__00000005_Semantic_Enforcer_diagram_1.Width = 200.000000
	__ProductShape__00000005_Semantic_Enforcer_diagram_1.Height = 80.000000

	__ProductShape__00000007_Docx_Backend_diagram_1.Name = `Docx Backend-diagram 1`
	__ProductShape__00000007_Docx_Backend_diagram_1.IsExpanded = false
	__ProductShape__00000007_Docx_Backend_diagram_1.X = 521.000000
	__ProductShape__00000007_Docx_Backend_diagram_1.Y = 690.999985
	__ProductShape__00000007_Docx_Backend_diagram_1.Width = 200.000000
	__ProductShape__00000007_Docx_Backend_diagram_1.Height = 80.000000

	__ProductShape__00000008_Backend_diagram_1.Name = `Backend-diagram 1`
	__ProductShape__00000008_Backend_diagram_1.IsExpanded = false
	__ProductShape__00000008_Backend_diagram_1.X = 336.000000
	__ProductShape__00000008_Backend_diagram_1.Y = 514.999985
	__ProductShape__00000008_Backend_diagram_1.Width = 200.000000
	__ProductShape__00000008_Backend_diagram_1.Height = 80.000000

	__ProductShape__00000010_views_diagram_1.Name = `views-diagram 1`
	__ProductShape__00000010_views_diagram_1.IsExpanded = false
	__ProductShape__00000010_views_diagram_1.X = 69.000000
	__ProductShape__00000010_views_diagram_1.Y = 366.999985
	__ProductShape__00000010_views_diagram_1.Width = 200.000000
	__ProductShape__00000010_views_diagram_1.Height = 80.000000

	__ProductShape__00000011_NewProduct_diagram_1.Name = `NewProduct-diagram 1`
	__ProductShape__00000011_NewProduct_diagram_1.IsExpanded = false
	__ProductShape__00000011_NewProduct_diagram_1.X = 938.999939
	__ProductShape__00000011_NewProduct_diagram_1.Y = 351.999985
	__ProductShape__00000011_NewProduct_diagram_1.Width = 200.000000
	__ProductShape__00000011_NewProduct_diagram_1.Height = 80.000000

	__ProductShape__00000012_UX_diagram_2.Name = `UX-diagram 2`
	__ProductShape__00000012_UX_diagram_2.IsExpanded = false
	__ProductShape__00000012_UX_diagram_2.X = 130.000000
	__ProductShape__00000012_UX_diagram_2.Y = 207.999985
	__ProductShape__00000012_UX_diagram_2.Width = 200.000000
	__ProductShape__00000012_UX_diagram_2.Height = 80.000000

	__ProductShape__00000013_WBS_tree_diagram_2.Name = `WBS tree-diagram 2`
	__ProductShape__00000013_WBS_tree_diagram_2.IsExpanded = false
	__ProductShape__00000013_WBS_tree_diagram_2.X = 282.000000
	__ProductShape__00000013_WBS_tree_diagram_2.Y = 340.999985
	__ProductShape__00000013_WBS_tree_diagram_2.Width = 200.000000
	__ProductShape__00000013_WBS_tree_diagram_2.Height = 80.000000

	__ProductShape__00000022_UX_NewDiagram.Name = `UX-NewDiagram`
	__ProductShape__00000022_UX_NewDiagram.IsExpanded = false
	__ProductShape__00000022_UX_NewDiagram.X = 312.924106
	__ProductShape__00000022_UX_NewDiagram.Y = 158.411418
	__ProductShape__00000022_UX_NewDiagram.Width = 200.000000
	__ProductShape__00000022_UX_NewDiagram.Height = 80.000000

	__ProductShape__00000024_PBS_tree_NewDiagram.Name = `PBS tree-NewDiagram`
	__ProductShape__00000024_PBS_tree_NewDiagram.IsExpanded = false
	__ProductShape__00000024_PBS_tree_NewDiagram.X = 76.164411
	__ProductShape__00000024_PBS_tree_NewDiagram.Y = 364.524270
	__ProductShape__00000024_PBS_tree_NewDiagram.Width = 200.000000
	__ProductShape__00000024_PBS_tree_NewDiagram.Height = 80.000000

	__ProductShape__00000025_views_NewDiagram.Name = `views-NewDiagram`
	__ProductShape__00000025_views_NewDiagram.IsExpanded = false
	__ProductShape__00000025_views_NewDiagram.X = 611.240168
	__ProductShape__00000025_views_NewDiagram.Y = 387.667699
	__ProductShape__00000025_views_NewDiagram.Width = 200.000000
	__ProductShape__00000025_views_NewDiagram.Height = 80.000000

	__ProductShape__00000026_WBS_tree_NewDiagram.Name = `WBS tree-NewDiagram`
	__ProductShape__00000026_WBS_tree_NewDiagram.IsExpanded = false
	__ProductShape__00000026_WBS_tree_NewDiagram.X = 344.053372
	__ProductShape__00000026_WBS_tree_NewDiagram.Y = 531.116054
	__ProductShape__00000026_WBS_tree_NewDiagram.Width = 200.000000
	__ProductShape__00000026_WBS_tree_NewDiagram.Height = 80.000000

	__ProductShape__00000027_Backend_NewDiagram.Name = `Backend-NewDiagram`
	__ProductShape__00000027_Backend_NewDiagram.IsExpanded = false
	__ProductShape__00000027_Backend_NewDiagram.X = 948.520260
	__ProductShape__00000027_Backend_NewDiagram.Y = 183.170288
	__ProductShape__00000027_Backend_NewDiagram.Width = 200.000000
	__ProductShape__00000027_Backend_NewDiagram.Height = 80.000000

	__Project__00000000_Project_Editor.Name = `Project Editor`
	__Project__00000000_Project_Editor.IsPBSNodeExpanded = true
	__Project__00000000_Project_Editor.IsWBSNodeExpanded = false
	__Project__00000000_Project_Editor.IsDiagramsNodeExpanded = true
	__Project__00000000_Project_Editor.IsExpanded = true
	__Project__00000000_Project_Editor.ComputedPrefix = ``

	__Project__00000001_DSME_Docx.Name = `DSME Docx`
	__Project__00000001_DSME_Docx.IsPBSNodeExpanded = false
	__Project__00000001_DSME_Docx.IsWBSNodeExpanded = false
	__Project__00000001_DSME_Docx.IsDiagramsNodeExpanded = false
	__Project__00000001_DSME_Docx.IsExpanded = true
	__Project__00000001_DSME_Docx.ComputedPrefix = ``

	__Root__00000000_Root.Name = `Root`
	__Root__00000000_Root.NbPixPerCharacter = 8.000000

	__Task__00000000_Develop_Backend.Name = `Develop Backend`
	__Task__00000000_Develop_Backend.IsExpanded = true
	__Task__00000000_Develop_Backend.ComputedPrefix = `1`
	__Task__00000000_Develop_Backend.IsInputsNodeExpanded = false
	__Task__00000000_Develop_Backend.IsOutputsNodeExpanded = false

	__Task__00000001_Dev_WBS_Tree.Name = `Dev WBS Tree`
	__Task__00000001_Dev_WBS_Tree.IsExpanded = false
	__Task__00000001_Dev_WBS_Tree.ComputedPrefix = `2`
	__Task__00000001_Dev_WBS_Tree.IsInputsNodeExpanded = false
	__Task__00000001_Dev_WBS_Tree.IsOutputsNodeExpanded = false

	__Task__00000002_Dev_PBS_Tree.Name = `Dev PBS Tree`
	__Task__00000002_Dev_PBS_Tree.IsExpanded = false
	__Task__00000002_Dev_PBS_Tree.ComputedPrefix = `3`
	__Task__00000002_Dev_PBS_Tree.IsInputsNodeExpanded = false
	__Task__00000002_Dev_PBS_Tree.IsOutputsNodeExpanded = false

	__Task__00000003_Dev_docx_Backend.Name = `Dev docx Backend`
	__Task__00000003_Dev_docx_Backend.IsExpanded = true
	__Task__00000003_Dev_docx_Backend.ComputedPrefix = `1`
	__Task__00000003_Dev_docx_Backend.IsInputsNodeExpanded = true
	__Task__00000003_Dev_docx_Backend.IsOutputsNodeExpanded = false

	__Task__00000004_Dev_views.Name = `Dev views`
	__Task__00000004_Dev_views.IsExpanded = false
	__Task__00000004_Dev_views.ComputedPrefix = `1.1`
	__Task__00000004_Dev_views.IsInputsNodeExpanded = false
	__Task__00000004_Dev_views.IsOutputsNodeExpanded = false

	__Task__00000005_Dev_UXx.Name = `Dev UXx`
	__Task__00000005_Dev_UXx.IsExpanded = false
	__Task__00000005_Dev_UXx.ComputedPrefix = `1.2`
	__Task__00000005_Dev_UXx.IsInputsNodeExpanded = false
	__Task__00000005_Dev_UXx.IsOutputsNodeExpanded = false

	// Setup of pointers
	// setup of CompositionShape instances pointers
	__CompositionShape__00000011_Backend_to_Semantic_Enforcer.Product = __Product__00000006_Semantic_Enforcer
	__CompositionShape__00000012_Backend_to_Docx_Backend.Product = __Product__00000009_Docx_Backend
	__CompositionShape__00000013_UX_to_WBS_tree.Product = __Product__00000002_WBS_tree
	__CompositionShape__00000014_UX_to_PBS_tree.Product = __Product__00000004_PBS_tree
	__CompositionShape__00000017_UX_to_views.Product = __Product__00000005_views
	__CompositionShape__00000019_UX_to_PBS_tree.Product = __Product__00000004_PBS_tree
	__CompositionShape__00000020_UX_to_views.Product = __Product__00000005_views
	__CompositionShape__00000021_UX_to_WBS_tree.Product = __Product__00000002_WBS_tree
	// setup of Product instances pointers
	__Product__00000000_UX.SubProducts = append(__Product__00000000_UX.SubProducts, __Product__00000002_WBS_tree)
	__Product__00000000_UX.SubProducts = append(__Product__00000000_UX.SubProducts, __Product__00000004_PBS_tree)
	__Product__00000000_UX.SubProducts = append(__Product__00000000_UX.SubProducts, __Product__00000005_views)
	__Product__00000001_Backend.SubProducts = append(__Product__00000001_Backend.SubProducts, __Product__00000006_Semantic_Enforcer)
	__Product__00000001_Backend.SubProducts = append(__Product__00000001_Backend.SubProducts, __Product__00000009_Docx_Backend)
	// setup of ProductShape instances pointers
	__ProductShape__00000000_UX_diagram_1.Product = __Product__00000000_UX
	__ProductShape__00000001_WBS_tree_diagram_1.Product = __Product__00000002_WBS_tree
	__ProductShape__00000002_PBS_tree_diagram_1.Product = __Product__00000004_PBS_tree
	__ProductShape__00000005_Semantic_Enforcer_diagram_1.Product = __Product__00000006_Semantic_Enforcer
	__ProductShape__00000007_Docx_Backend_diagram_1.Product = __Product__00000009_Docx_Backend
	__ProductShape__00000008_Backend_diagram_1.Product = __Product__00000001_Backend
	__ProductShape__00000010_views_diagram_1.Product = __Product__00000005_views
	__ProductShape__00000012_UX_diagram_2.Product = __Product__00000000_UX
	__ProductShape__00000013_WBS_tree_diagram_2.Product = __Product__00000002_WBS_tree
	__ProductShape__00000022_UX_NewDiagram.Product = __Product__00000000_UX
	__ProductShape__00000024_PBS_tree_NewDiagram.Product = __Product__00000004_PBS_tree
	__ProductShape__00000025_views_NewDiagram.Product = __Product__00000005_views
	__ProductShape__00000026_WBS_tree_NewDiagram.Product = __Product__00000002_WBS_tree
	__ProductShape__00000027_Backend_NewDiagram.Product = __Product__00000001_Backend
	// setup of Project instances pointers
	__Project__00000000_Project_Editor.RootProducts = append(__Project__00000000_Project_Editor.RootProducts, __Product__00000000_UX)
	__Project__00000000_Project_Editor.RootProducts = append(__Project__00000000_Project_Editor.RootProducts, __Product__00000001_Backend)
	__Project__00000000_Project_Editor.RootTasks = append(__Project__00000000_Project_Editor.RootTasks, __Task__00000000_Develop_Backend)
	__Project__00000000_Project_Editor.RootTasks = append(__Project__00000000_Project_Editor.RootTasks, __Task__00000001_Dev_WBS_Tree)
	__Project__00000000_Project_Editor.RootTasks = append(__Project__00000000_Project_Editor.RootTasks, __Task__00000002_Dev_PBS_Tree)
	__Project__00000000_Project_Editor.Diagrams = append(__Project__00000000_Project_Editor.Diagrams, )
	__Project__00000001_DSME_Docx.RootTasks = append(__Project__00000001_DSME_Docx.RootTasks, __Task__00000003_Dev_docx_Backend)
	// setup of Root instances pointers
	__Root__00000000_Root.Projects = append(__Root__00000000_Root.Projects, __Project__00000000_Project_Editor)
	__Root__00000000_Root.Projects = append(__Root__00000000_Root.Projects, __Project__00000001_DSME_Docx)
	// setup of Task instances pointers
	__Task__00000000_Develop_Backend.SubTasks = append(__Task__00000000_Develop_Backend.SubTasks, __Task__00000004_Dev_views)
	__Task__00000000_Develop_Backend.SubTasks = append(__Task__00000000_Develop_Backend.SubTasks, __Task__00000005_Dev_UXx)
	__Task__00000003_Dev_docx_Backend.Outputs = append(__Task__00000003_Dev_docx_Backend.Outputs, __Product__00000009_Docx_Backend)
	__Task__00000003_Dev_docx_Backend.Outputs = append(__Task__00000003_Dev_docx_Backend.Outputs, __Product__00000006_Semantic_Enforcer)
	__Task__00000005_Dev_UXx.Outputs = append(__Task__00000005_Dev_UXx.Outputs, __Product__00000002_WBS_tree)
	__Task__00000005_Dev_UXx.Outputs = append(__Task__00000005_Dev_UXx.Outputs, __Product__00000004_PBS_tree)
	__Task__00000005_Dev_UXx.Outputs = append(__Task__00000005_Dev_UXx.Outputs, __Product__00000005_views)
}

