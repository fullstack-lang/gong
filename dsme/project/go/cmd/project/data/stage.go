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

	__Diagram__00000000_NewDiagram := (&models.Diagram{}).Stage(stage)

	__Product__00000000_UX := (&models.Product{}).Stage(stage)
	__Product__00000001_Backend := (&models.Product{}).Stage(stage)
	__Product__00000002_WBS_tree := (&models.Product{}).Stage(stage)
	__Product__00000004_PBS_tree := (&models.Product{}).Stage(stage)
	__Product__00000005_views := (&models.Product{}).Stage(stage)
	__Product__00000006_Semantic_Enforcer := (&models.Product{}).Stage(stage)
	__Product__00000009_Docx_Backend := (&models.Product{}).Stage(stage)

	__ProductCompositionShape__00000003_UX_to_WBS_tree := (&models.ProductCompositionShape{}).Stage(stage)
	__ProductCompositionShape__00000005_UX_to_PBS_tree := (&models.ProductCompositionShape{}).Stage(stage)
	__ProductCompositionShape__00000006_Backend_to_Semantic_Enforcer := (&models.ProductCompositionShape{}).Stage(stage)

	__ProductShape__00000002_WBS_tree_NewDiagram := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000004_UX_NewDiagram := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000005_PBS_tree_NewDiagram := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000006_Backend_NewDiagram := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000007_Semantic_Enforcer_NewDiagram := (&models.ProductShape{}).Stage(stage)

	__Project__00000000_Project_Editor := (&models.Project{}).Stage(stage)
	__Project__00000001_DSME_Docx := (&models.Project{}).Stage(stage)

	__Root__00000000_Root := (&models.Root{}).Stage(stage)

	__Task__00000000_Develop_Backend := (&models.Task{}).Stage(stage)
	__Task__00000001_Dev_WBS_Tree := (&models.Task{}).Stage(stage)
	__Task__00000002_Dev_PBS_Tree := (&models.Task{}).Stage(stage)
	__Task__00000003_Dev_docx_Backend := (&models.Task{}).Stage(stage)
	__Task__00000004_Dev_views := (&models.Task{}).Stage(stage)
	__Task__00000005_Dev_UXx := (&models.Task{}).Stage(stage)

	__TaskCompositionShape__00000000_Develop_Backend_to_Dev_views := (&models.TaskCompositionShape{}).Stage(stage)
	__TaskCompositionShape__00000001_Develop_Backend_to_Dev_UXx := (&models.TaskCompositionShape{}).Stage(stage)

	__TaskShape__00000000_Develop_Backend_NewDiagram := (&models.TaskShape{}).Stage(stage)
	__TaskShape__00000001_Dev_views_NewDiagram := (&models.TaskShape{}).Stage(stage)
	__TaskShape__00000002_Dev_UXx_NewDiagram := (&models.TaskShape{}).Stage(stage)
	__TaskShape__00000003_Dev_WBS_Tree_NewDiagram := (&models.TaskShape{}).Stage(stage)

	// Setup of values

	__Diagram__00000000_NewDiagram.Name = `NewDiagram`
	__Diagram__00000000_NewDiagram.IsChecked = true
	__Diagram__00000000_NewDiagram.IsEditable_ = true
	__Diagram__00000000_NewDiagram.IsInRenameMode = false
	__Diagram__00000000_NewDiagram.IsExpanded = true
	__Diagram__00000000_NewDiagram.ComputedPrefix = ``
	__Diagram__00000000_NewDiagram.IsPBSNodeExpanded = true
	__Diagram__00000000_NewDiagram.IsWBSNodeExpanded = true

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
	__Product__00000004_PBS_tree.IsExpanded = true
	__Product__00000004_PBS_tree.ComputedPrefix = `1.2`
	__Product__00000004_PBS_tree.IsProducersNodeExpanded = true
	__Product__00000004_PBS_tree.IsConsumersNodeExpanded = false

	__Product__00000005_views.Name = `views`
	__Product__00000005_views.IsExpanded = false
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

	__ProductCompositionShape__00000003_UX_to_WBS_tree.Name = `UX to WBS tree`
	__ProductCompositionShape__00000003_UX_to_WBS_tree.StartRatio = 0.500000
	__ProductCompositionShape__00000003_UX_to_WBS_tree.EndRatio = 0.500000
	__ProductCompositionShape__00000003_UX_to_WBS_tree.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000003_UX_to_WBS_tree.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000003_UX_to_WBS_tree.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000005_UX_to_PBS_tree.Name = `UX to PBS tree`
	__ProductCompositionShape__00000005_UX_to_PBS_tree.StartRatio = 0.500000
	__ProductCompositionShape__00000005_UX_to_PBS_tree.EndRatio = 0.500000
	__ProductCompositionShape__00000005_UX_to_PBS_tree.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000005_UX_to_PBS_tree.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000005_UX_to_PBS_tree.CornerOffsetRatio = 1.680000

	__ProductCompositionShape__00000006_Backend_to_Semantic_Enforcer.Name = `Backend to Semantic Enforcer`
	__ProductCompositionShape__00000006_Backend_to_Semantic_Enforcer.StartRatio = 0.500000
	__ProductCompositionShape__00000006_Backend_to_Semantic_Enforcer.EndRatio = 0.500000
	__ProductCompositionShape__00000006_Backend_to_Semantic_Enforcer.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000006_Backend_to_Semantic_Enforcer.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000006_Backend_to_Semantic_Enforcer.CornerOffsetRatio = 1.680000

	__ProductShape__00000002_WBS_tree_NewDiagram.Name = `WBS tree-NewDiagram`
	__ProductShape__00000002_WBS_tree_NewDiagram.IsExpanded = false
	__ProductShape__00000002_WBS_tree_NewDiagram.X = 151.710110
	__ProductShape__00000002_WBS_tree_NewDiagram.Y = 523.165460
	__ProductShape__00000002_WBS_tree_NewDiagram.Width = 200.000000
	__ProductShape__00000002_WBS_tree_NewDiagram.Height = 80.000000

	__ProductShape__00000004_UX_NewDiagram.Name = `UX-NewDiagram`
	__ProductShape__00000004_UX_NewDiagram.IsExpanded = false
	__ProductShape__00000004_UX_NewDiagram.X = 457.575064
	__ProductShape__00000004_UX_NewDiagram.Y = 286.316429
	__ProductShape__00000004_UX_NewDiagram.Width = 200.000000
	__ProductShape__00000004_UX_NewDiagram.Height = 80.000000

	__ProductShape__00000005_PBS_tree_NewDiagram.Name = `PBS tree-NewDiagram`
	__ProductShape__00000005_PBS_tree_NewDiagram.IsExpanded = false
	__ProductShape__00000005_PBS_tree_NewDiagram.X = 831.532973
	__ProductShape__00000005_PBS_tree_NewDiagram.Y = 493.239329
	__ProductShape__00000005_PBS_tree_NewDiagram.Width = 200.000000
	__ProductShape__00000005_PBS_tree_NewDiagram.Height = 80.000000

	__ProductShape__00000006_Backend_NewDiagram.Name = `Backend-NewDiagram`
	__ProductShape__00000006_Backend_NewDiagram.IsExpanded = false
	__ProductShape__00000006_Backend_NewDiagram.X = 917.979641
	__ProductShape__00000006_Backend_NewDiagram.Y = 234.154442
	__ProductShape__00000006_Backend_NewDiagram.Width = 200.000000
	__ProductShape__00000006_Backend_NewDiagram.Height = 80.000000

	__ProductShape__00000007_Semantic_Enforcer_NewDiagram.Name = `Semantic Enforcer-NewDiagram`
	__ProductShape__00000007_Semantic_Enforcer_NewDiagram.IsExpanded = false
	__ProductShape__00000007_Semantic_Enforcer_NewDiagram.X = 1187.293838
	__ProductShape__00000007_Semantic_Enforcer_NewDiagram.Y = 485.853802
	__ProductShape__00000007_Semantic_Enforcer_NewDiagram.Width = 200.000000
	__ProductShape__00000007_Semantic_Enforcer_NewDiagram.Height = 80.000000

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

	__TaskCompositionShape__00000000_Develop_Backend_to_Dev_views.Name = `Develop Backend to Dev views`
	__TaskCompositionShape__00000000_Develop_Backend_to_Dev_views.StartRatio = 0.500000
	__TaskCompositionShape__00000000_Develop_Backend_to_Dev_views.EndRatio = 1.000000
	__TaskCompositionShape__00000000_Develop_Backend_to_Dev_views.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000000_Develop_Backend_to_Dev_views.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000000_Develop_Backend_to_Dev_views.CornerOffsetRatio = 1.680000

	__TaskCompositionShape__00000001_Develop_Backend_to_Dev_UXx.Name = `Develop Backend to Dev UXx`
	__TaskCompositionShape__00000001_Develop_Backend_to_Dev_UXx.StartRatio = 0.500000
	__TaskCompositionShape__00000001_Develop_Backend_to_Dev_UXx.EndRatio = 0.500000
	__TaskCompositionShape__00000001_Develop_Backend_to_Dev_UXx.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000001_Develop_Backend_to_Dev_UXx.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000001_Develop_Backend_to_Dev_UXx.CornerOffsetRatio = 1.680000

	__TaskShape__00000000_Develop_Backend_NewDiagram.Name = `Develop Backend-NewDiagram`
	__TaskShape__00000000_Develop_Backend_NewDiagram.IsExpanded = false
	__TaskShape__00000000_Develop_Backend_NewDiagram.X = 565.825736
	__TaskShape__00000000_Develop_Backend_NewDiagram.Y = 611.338520
	__TaskShape__00000000_Develop_Backend_NewDiagram.Width = 200.000000
	__TaskShape__00000000_Develop_Backend_NewDiagram.Height = 80.000000

	__TaskShape__00000001_Dev_views_NewDiagram.Name = `Dev views-NewDiagram`
	__TaskShape__00000001_Dev_views_NewDiagram.IsExpanded = false
	__TaskShape__00000001_Dev_views_NewDiagram.X = 328.923936
	__TaskShape__00000001_Dev_views_NewDiagram.Y = 840.432698
	__TaskShape__00000001_Dev_views_NewDiagram.Width = 200.000000
	__TaskShape__00000001_Dev_views_NewDiagram.Height = 80.000000

	__TaskShape__00000002_Dev_UXx_NewDiagram.Name = `Dev UXx-NewDiagram`
	__TaskShape__00000002_Dev_UXx_NewDiagram.IsExpanded = false
	__TaskShape__00000002_Dev_UXx_NewDiagram.X = 845.853316
	__TaskShape__00000002_Dev_UXx_NewDiagram.Y = 852.695171
	__TaskShape__00000002_Dev_UXx_NewDiagram.Width = 200.000000
	__TaskShape__00000002_Dev_UXx_NewDiagram.Height = 80.000000

	__TaskShape__00000003_Dev_WBS_Tree_NewDiagram.Name = `Dev WBS Tree-NewDiagram`
	__TaskShape__00000003_Dev_WBS_Tree_NewDiagram.IsExpanded = false
	__TaskShape__00000003_Dev_WBS_Tree_NewDiagram.X = 116.250055
	__TaskShape__00000003_Dev_WBS_Tree_NewDiagram.Y = 354.790003
	__TaskShape__00000003_Dev_WBS_Tree_NewDiagram.Width = 200.000000
	__TaskShape__00000003_Dev_WBS_Tree_NewDiagram.Height = 80.000000

	// Setup of pointers
	// setup of Diagram instances pointers
	__Diagram__00000000_NewDiagram.Product_Shapes = append(__Diagram__00000000_NewDiagram.Product_Shapes, __ProductShape__00000002_WBS_tree_NewDiagram)
	__Diagram__00000000_NewDiagram.Product_Shapes = append(__Diagram__00000000_NewDiagram.Product_Shapes, __ProductShape__00000004_UX_NewDiagram)
	__Diagram__00000000_NewDiagram.Product_Shapes = append(__Diagram__00000000_NewDiagram.Product_Shapes, __ProductShape__00000005_PBS_tree_NewDiagram)
	__Diagram__00000000_NewDiagram.Product_Shapes = append(__Diagram__00000000_NewDiagram.Product_Shapes, __ProductShape__00000006_Backend_NewDiagram)
	__Diagram__00000000_NewDiagram.Product_Shapes = append(__Diagram__00000000_NewDiagram.Product_Shapes, __ProductShape__00000007_Semantic_Enforcer_NewDiagram)
	__Diagram__00000000_NewDiagram.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_NewDiagram.ProductsWhoseNodeIsExpanded, __Product__00000001_Backend)
	__Diagram__00000000_NewDiagram.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_NewDiagram.ProductsWhoseNodeIsExpanded, __Product__00000000_UX)
	__Diagram__00000000_NewDiagram.ProductComposition_Shapes = append(__Diagram__00000000_NewDiagram.ProductComposition_Shapes, __ProductCompositionShape__00000003_UX_to_WBS_tree)
	__Diagram__00000000_NewDiagram.ProductComposition_Shapes = append(__Diagram__00000000_NewDiagram.ProductComposition_Shapes, __ProductCompositionShape__00000005_UX_to_PBS_tree)
	__Diagram__00000000_NewDiagram.ProductComposition_Shapes = append(__Diagram__00000000_NewDiagram.ProductComposition_Shapes, __ProductCompositionShape__00000006_Backend_to_Semantic_Enforcer)
	__Diagram__00000000_NewDiagram.Task_Shapes = append(__Diagram__00000000_NewDiagram.Task_Shapes, __TaskShape__00000000_Develop_Backend_NewDiagram)
	__Diagram__00000000_NewDiagram.Task_Shapes = append(__Diagram__00000000_NewDiagram.Task_Shapes, __TaskShape__00000001_Dev_views_NewDiagram)
	__Diagram__00000000_NewDiagram.Task_Shapes = append(__Diagram__00000000_NewDiagram.Task_Shapes, __TaskShape__00000002_Dev_UXx_NewDiagram)
	__Diagram__00000000_NewDiagram.Task_Shapes = append(__Diagram__00000000_NewDiagram.Task_Shapes, __TaskShape__00000003_Dev_WBS_Tree_NewDiagram)
	__Diagram__00000000_NewDiagram.TasksWhoseNodeIsExpanded = append(__Diagram__00000000_NewDiagram.TasksWhoseNodeIsExpanded, __Task__00000000_Develop_Backend)
	__Diagram__00000000_NewDiagram.TaskComposition_Shapes = append(__Diagram__00000000_NewDiagram.TaskComposition_Shapes, __TaskCompositionShape__00000000_Develop_Backend_to_Dev_views)
	__Diagram__00000000_NewDiagram.TaskComposition_Shapes = append(__Diagram__00000000_NewDiagram.TaskComposition_Shapes, __TaskCompositionShape__00000001_Develop_Backend_to_Dev_UXx)
	// setup of Product instances pointers
	__Product__00000000_UX.SubProducts = append(__Product__00000000_UX.SubProducts, __Product__00000002_WBS_tree)
	__Product__00000000_UX.SubProducts = append(__Product__00000000_UX.SubProducts, __Product__00000004_PBS_tree)
	__Product__00000000_UX.SubProducts = append(__Product__00000000_UX.SubProducts, __Product__00000005_views)
	__Product__00000001_Backend.SubProducts = append(__Product__00000001_Backend.SubProducts, __Product__00000006_Semantic_Enforcer)
	__Product__00000001_Backend.SubProducts = append(__Product__00000001_Backend.SubProducts, __Product__00000009_Docx_Backend)
	// setup of ProductCompositionShape instances pointers
	__ProductCompositionShape__00000003_UX_to_WBS_tree.Product = __Product__00000002_WBS_tree
	__ProductCompositionShape__00000005_UX_to_PBS_tree.Product = __Product__00000004_PBS_tree
	__ProductCompositionShape__00000006_Backend_to_Semantic_Enforcer.Product = __Product__00000006_Semantic_Enforcer
	// setup of ProductShape instances pointers
	__ProductShape__00000002_WBS_tree_NewDiagram.Product = __Product__00000002_WBS_tree
	__ProductShape__00000004_UX_NewDiagram.Product = __Product__00000000_UX
	__ProductShape__00000005_PBS_tree_NewDiagram.Product = __Product__00000004_PBS_tree
	__ProductShape__00000006_Backend_NewDiagram.Product = __Product__00000001_Backend
	__ProductShape__00000007_Semantic_Enforcer_NewDiagram.Product = __Product__00000006_Semantic_Enforcer
	// setup of Project instances pointers
	__Project__00000000_Project_Editor.RootProducts = append(__Project__00000000_Project_Editor.RootProducts, __Product__00000000_UX)
	__Project__00000000_Project_Editor.RootProducts = append(__Project__00000000_Project_Editor.RootProducts, __Product__00000001_Backend)
	__Project__00000000_Project_Editor.RootTasks = append(__Project__00000000_Project_Editor.RootTasks, __Task__00000000_Develop_Backend)
	__Project__00000000_Project_Editor.RootTasks = append(__Project__00000000_Project_Editor.RootTasks, __Task__00000001_Dev_WBS_Tree)
	__Project__00000000_Project_Editor.RootTasks = append(__Project__00000000_Project_Editor.RootTasks, __Task__00000002_Dev_PBS_Tree)
	__Project__00000000_Project_Editor.Diagrams = append(__Project__00000000_Project_Editor.Diagrams, __Diagram__00000000_NewDiagram)
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
	// setup of TaskCompositionShape instances pointers
	__TaskCompositionShape__00000000_Develop_Backend_to_Dev_views.Task = __Task__00000004_Dev_views
	__TaskCompositionShape__00000001_Develop_Backend_to_Dev_UXx.Task = __Task__00000005_Dev_UXx
	// setup of TaskShape instances pointers
	__TaskShape__00000000_Develop_Backend_NewDiagram.Task = __Task__00000000_Develop_Backend
	__TaskShape__00000001_Dev_views_NewDiagram.Task = __Task__00000004_Dev_views
	__TaskShape__00000002_Dev_UXx_NewDiagram.Task = __Task__00000005_Dev_UXx
	__TaskShape__00000003_Dev_WBS_Tree_NewDiagram.Task = __Task__00000001_Dev_WBS_Tree
}

