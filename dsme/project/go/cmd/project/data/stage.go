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

	__Diagram__00000000_diagram_1 := (&models.Diagram{}).Stage(stage)
	__Diagram__00000001_diagram_2 := (&models.Diagram{}).Stage(stage)

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
	__ProductShape__00000003_views_diagram_1 := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000004_Backend_diagram_1 := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000005_Semantic_Enforcer_diagram_1 := (&models.ProductShape{}).Stage(stage)
	__ProductShape__00000006_Docx_Backend_diagram_1 := (&models.ProductShape{}).Stage(stage)

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

	__Diagram__00000000_diagram_1.Name = `diagram 1`
	__Diagram__00000000_diagram_1.IsChecked = true
	__Diagram__00000000_diagram_1.IsEditable_ = true
	__Diagram__00000000_diagram_1.IsInRenameMode = false
	__Diagram__00000000_diagram_1.IsExpanded = true
	__Diagram__00000000_diagram_1.ComputedPrefix = ``

	__Diagram__00000001_diagram_2.Name = `diagram 2`
	__Diagram__00000001_diagram_2.IsChecked = false
	__Diagram__00000001_diagram_2.IsEditable_ = false
	__Diagram__00000001_diagram_2.IsInRenameMode = false
	__Diagram__00000001_diagram_2.IsExpanded = false
	__Diagram__00000001_diagram_2.ComputedPrefix = ``

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

	__ProductShape__00000003_views_diagram_1.Name = `views-diagram 1`
	__ProductShape__00000003_views_diagram_1.IsExpanded = false
	__ProductShape__00000003_views_diagram_1.X = 68.000000
	__ProductShape__00000003_views_diagram_1.Y = 337.999985
	__ProductShape__00000003_views_diagram_1.Width = 200.000000
	__ProductShape__00000003_views_diagram_1.Height = 80.000000

	__ProductShape__00000004_Backend_diagram_1.Name = `Backend-diagram 1`
	__ProductShape__00000004_Backend_diagram_1.IsExpanded = false
	__ProductShape__00000004_Backend_diagram_1.X = 97.000000
	__ProductShape__00000004_Backend_diagram_1.Y = 524.999985
	__ProductShape__00000004_Backend_diagram_1.Width = 805.000000
	__ProductShape__00000004_Backend_diagram_1.Height = 80.000000

	__ProductShape__00000005_Semantic_Enforcer_diagram_1.Name = `Semantic Enforcer-diagram 1`
	__ProductShape__00000005_Semantic_Enforcer_diagram_1.IsExpanded = false
	__ProductShape__00000005_Semantic_Enforcer_diagram_1.X = 95.000000
	__ProductShape__00000005_Semantic_Enforcer_diagram_1.Y = 708.999985
	__ProductShape__00000005_Semantic_Enforcer_diagram_1.Width = 200.000000
	__ProductShape__00000005_Semantic_Enforcer_diagram_1.Height = 80.000000

	__ProductShape__00000006_Docx_Backend_diagram_1.Name = `Docx Backend-diagram 1`
	__ProductShape__00000006_Docx_Backend_diagram_1.IsExpanded = false
	__ProductShape__00000006_Docx_Backend_diagram_1.X = 379.000000
	__ProductShape__00000006_Docx_Backend_diagram_1.Y = 705.999985
	__ProductShape__00000006_Docx_Backend_diagram_1.Width = 200.000000
	__ProductShape__00000006_Docx_Backend_diagram_1.Height = 80.000000

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
	// setup of Diagram instances pointers
	__Diagram__00000000_diagram_1.Product_Shapes = append(__Diagram__00000000_diagram_1.Product_Shapes, __ProductShape__00000000_UX_diagram_1)
	__Diagram__00000000_diagram_1.Product_Shapes = append(__Diagram__00000000_diagram_1.Product_Shapes, __ProductShape__00000001_WBS_tree_diagram_1)
	__Diagram__00000000_diagram_1.Product_Shapes = append(__Diagram__00000000_diagram_1.Product_Shapes, __ProductShape__00000002_PBS_tree_diagram_1)
	__Diagram__00000000_diagram_1.Product_Shapes = append(__Diagram__00000000_diagram_1.Product_Shapes, __ProductShape__00000003_views_diagram_1)
	__Diagram__00000000_diagram_1.Product_Shapes = append(__Diagram__00000000_diagram_1.Product_Shapes, __ProductShape__00000004_Backend_diagram_1)
	__Diagram__00000000_diagram_1.Product_Shapes = append(__Diagram__00000000_diagram_1.Product_Shapes, __ProductShape__00000005_Semantic_Enforcer_diagram_1)
	__Diagram__00000000_diagram_1.Product_Shapes = append(__Diagram__00000000_diagram_1.Product_Shapes, __ProductShape__00000006_Docx_Backend_diagram_1)
	__Diagram__00000000_diagram_1.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_diagram_1.ProductsWhoseNodeIsExpanded, __Product__00000000_UX)
	__Diagram__00000000_diagram_1.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_diagram_1.ProductsWhoseNodeIsExpanded, __Product__00000001_Backend)
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
	__ProductShape__00000003_views_diagram_1.Product = __Product__00000005_views
	__ProductShape__00000004_Backend_diagram_1.Product = __Product__00000001_Backend
	__ProductShape__00000005_Semantic_Enforcer_diagram_1.Product = __Product__00000006_Semantic_Enforcer
	__ProductShape__00000006_Docx_Backend_diagram_1.Product = __Product__00000009_Docx_Backend
	// setup of Project instances pointers
	__Project__00000000_Project_Editor.RootProducts = append(__Project__00000000_Project_Editor.RootProducts, __Product__00000000_UX)
	__Project__00000000_Project_Editor.RootProducts = append(__Project__00000000_Project_Editor.RootProducts, __Product__00000001_Backend)
	__Project__00000000_Project_Editor.RootTasks = append(__Project__00000000_Project_Editor.RootTasks, __Task__00000000_Develop_Backend)
	__Project__00000000_Project_Editor.RootTasks = append(__Project__00000000_Project_Editor.RootTasks, __Task__00000001_Dev_WBS_Tree)
	__Project__00000000_Project_Editor.RootTasks = append(__Project__00000000_Project_Editor.RootTasks, __Task__00000002_Dev_PBS_Tree)
	__Project__00000000_Project_Editor.Diagrams = append(__Project__00000000_Project_Editor.Diagrams, __Diagram__00000000_diagram_1)
	__Project__00000000_Project_Editor.Diagrams = append(__Project__00000000_Project_Editor.Diagrams, __Diagram__00000001_diagram_2)
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

