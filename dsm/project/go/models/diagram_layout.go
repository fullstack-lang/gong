package models

import (
	"sort"
)

// layoutDiagram performs a hierarchical vertical layout of all visible nodes in the diagram.
// The parent node aligns horizontally with its FIRST child.
func layoutDiagram(diagram *Diagram, stager *Stager) {
	var margin float64 = 50

	var startX float64 = 50
	startY := 50.0

	if len(diagram.Product_Shapes) > 0 {
		layoutProductShapes(diagram, &startX, startY, margin)
	}
	if len(diagram.Task_Shapes) > 0 {
		layoutTaskShapes(diagram, &startX, startY, margin)
	}
	if len(diagram.Resource_Shapes) > 0 {
		layoutResourceShapes(diagram, &startX, startY, margin)
	}
}

// Product layout
type productNode struct {
	shape    *ProductShape
	children []*productNode
}

func layoutProductShapes(diagram *Diagram, nextX *float64, startY float64, margin float64) {
	nodesByProduct := make(map[*Product]*productNode)
	var rootNodes []*productNode

	for _, shape := range diagram.Product_Shapes {
		nodesByProduct[shape.Product] = &productNode{shape: shape}
	}

	for _, shape := range diagram.Product_Shapes {
		product := shape.Product
		node := nodesByProduct[product]
		isRoot := true

		if product.parentProduct != nil {
			if parentNode, exists := nodesByProduct[product.parentProduct]; exists {
				parentNode.children = append(parentNode.children, node)
				isRoot = false
			}
		}

		if isRoot {
			rootNodes = append(rootNodes, node)
		}
	}

	sort.Slice(rootNodes, func(i, j int) bool {
		return rootNodes[i].shape.Product.Name < rootNodes[j].shape.Product.Name
	})
	for _, root := range rootNodes {
		sortProductChildren(root)
	}

	for _, root := range rootNodes {
		layoutProductDFS(root, nextX, startY, margin)
	}
	
	// Update composition link corner ratio to 1.5 to enforce orthogonal routing perfectly matching parent-child alignment
	for _, link := range diagram.ProductComposition_Shapes {
		link.SetCornerOffsetRatio(1.5)
	}
}

func sortProductChildren(node *productNode) {
	sort.Slice(node.children, func(i, j int) bool {
		return node.children[i].shape.Product.Name < node.children[j].shape.Product.Name
	})
	for _, child := range node.children {
		sortProductChildren(child)
	}
}

func layoutProductDFS(node *productNode, nextX *float64, currentY float64, margin float64) {
	node.shape.SetY(currentY)
	w := node.shape.GetWidth()
	h := node.shape.GetHeight()

	if len(node.children) == 0 {
		node.shape.SetX(*nextX)
		*nextX += w + margin
		return
	}

	for _, child := range node.children {
		layoutProductDFS(child, nextX, currentY+h*2.0, margin)
	}
	
	node.shape.SetX(node.children[0].shape.GetX())
}

// Task layout
type taskNode struct {
	shape    *TaskShape
	children []*taskNode
}

func layoutTaskShapes(diagram *Diagram, nextX *float64, startY float64, margin float64) {
	nodesByTask := make(map[*Task]*taskNode)
	var rootNodes []*taskNode

	for _, shape := range diagram.Task_Shapes {
		nodesByTask[shape.Task] = &taskNode{shape: shape}
	}

	for _, shape := range diagram.Task_Shapes {
		task := shape.Task
		node := nodesByTask[task]
		isRoot := true

		if task.parentTask != nil {
			if parentNode, exists := nodesByTask[task.parentTask]; exists {
				parentNode.children = append(parentNode.children, node)
				isRoot = false
			}
		}

		if isRoot {
			rootNodes = append(rootNodes, node)
		}
	}

	sort.Slice(rootNodes, func(i, j int) bool {
		return rootNodes[i].shape.Task.Name < rootNodes[j].shape.Task.Name
	})
	for _, root := range rootNodes {
		sortTaskChildren(root)
	}

	for _, root := range rootNodes {
		layoutTaskDFS(root, nextX, startY, margin)
	}
	
	for _, link := range diagram.TaskComposition_Shapes {
		link.SetCornerOffsetRatio(1.5)
	}
}

func sortTaskChildren(node *taskNode) {
	sort.Slice(node.children, func(i, j int) bool {
		return node.children[i].shape.Task.Name < node.children[j].shape.Task.Name
	})
	for _, child := range node.children {
		sortTaskChildren(child)
	}
}

func layoutTaskDFS(node *taskNode, nextX *float64, currentY float64, margin float64) {
	node.shape.SetY(currentY)
	w := node.shape.GetWidth()
	h := node.shape.GetHeight()

	if len(node.children) == 0 {
		node.shape.SetX(*nextX)
		*nextX += w + margin
		return
	}

	for _, child := range node.children {
		layoutTaskDFS(child, nextX, currentY+h*2.0, margin)
	}

	node.shape.SetX(node.children[0].shape.GetX())
}

// Resource layout
type resourceNode struct {
	shape    *ResourceShape
	children []*resourceNode
}

func layoutResourceShapes(diagram *Diagram, nextX *float64, startY float64, margin float64) {
	nodesByResource := make(map[*Resource]*resourceNode)
	var rootNodes []*resourceNode

	for _, shape := range diagram.Resource_Shapes {
		nodesByResource[shape.Resource] = &resourceNode{shape: shape}
	}

	for _, shape := range diagram.Resource_Shapes {
		resource := shape.Resource
		node := nodesByResource[resource]
		isRoot := true

		if resource.parentResource != nil {
			if parentNode, exists := nodesByResource[resource.parentResource]; exists {
				parentNode.children = append(parentNode.children, node)
				isRoot = false
			}
		}

		if isRoot {
			rootNodes = append(rootNodes, node)
		}
	}

	sort.Slice(rootNodes, func(i, j int) bool {
		return rootNodes[i].shape.Resource.Name < rootNodes[j].shape.Resource.Name
	})
	for _, root := range rootNodes {
		sortResourceChildren(root)
	}

	for _, root := range rootNodes {
		layoutResourceDFS(root, nextX, startY, margin)
	}
	
	for _, link := range diagram.ResourceComposition_Shapes {
		link.SetCornerOffsetRatio(1.5)
	}
}

func sortResourceChildren(node *resourceNode) {
	sort.Slice(node.children, func(i, j int) bool {
		return node.children[i].shape.Resource.Name < node.children[j].shape.Resource.Name
	})
	for _, child := range node.children {
		sortResourceChildren(child)
	}
}

func layoutResourceDFS(node *resourceNode, nextX *float64, currentY float64, margin float64) {
	node.shape.SetY(currentY)
	w := node.shape.GetWidth()
	h := node.shape.GetHeight()

	if len(node.children) == 0 {
		node.shape.SetX(*nextX)
		*nextX += w + margin
		return
	}

	for _, child := range node.children {
		layoutResourceDFS(child, nextX, currentY+h*2.0, margin)
	}

	node.shape.SetX(node.children[0].shape.GetX())
}
