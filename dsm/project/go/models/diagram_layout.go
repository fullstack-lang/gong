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

	rootLibrary := stager.GetRootLibrary()

	if len(diagram.Product_Shapes) > 0 {
		layoutProductShapes(diagram, rootLibrary, &startX, startY, margin)
	}
	if len(diagram.Task_Shapes) > 0 {
		layoutTaskShapes(diagram, rootLibrary, &startX, startY, margin)
	}
	if len(diagram.Resource_Shapes) > 0 {
		layoutResourceShapes(diagram, rootLibrary, &startX, startY, margin)
	}
}

// Product layout
type productNode struct {
	shape    *ProductShape
	children []*productNode
}

func layoutProductShapes(diagram *Diagram, rootLibrary *Library, nextX *float64, startY float64, margin float64) {
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
			if _, exists := nodesByProduct[product.parentProduct]; exists {
				isRoot = false
			}
		}

		if isRoot {
			rootNodes = append(rootNodes, node)
		}
	}

	// Populate children respecting SubProducts order
	for _, node := range nodesByProduct {
		for _, subProduct := range node.shape.Product.SubProducts {
			if childNode, exists := nodesByProduct[subProduct]; exists {
				node.children = append(node.children, childNode)
			}
		}
	}

	rootProductIndex := make(map[*Product]int)
	if rootLibrary != nil {
		for i, sub := range rootLibrary.RootProducts {
			rootProductIndex[sub] = i
		}
	}

	sort.Slice(rootNodes, func(i, j int) bool {
		idxI, okI := rootProductIndex[rootNodes[i].shape.Product]
		idxJ, okJ := rootProductIndex[rootNodes[j].shape.Product]
		if okI && okJ {
			return idxI < idxJ
		}
		if okI {
			return true
		}
		if okJ {
			return false
		}
		return rootNodes[i].shape.Product.Name < rootNodes[j].shape.Product.Name
	})

	for _, root := range rootNodes {
		maxX, _ := layoutProductDFS(root, *nextX, startY, margin)
		*nextX = maxX
	}
	
	// Find the parent of each product based on SubProducts
	parentByProduct := make(map[*Product]*productNode)
	for _, node := range nodesByProduct {
		for _, subProduct := range node.shape.Product.SubProducts {
			parentByProduct[subProduct] = node
		}
	}

	for _, link := range diagram.ProductComposition_Shapes {
		link.SetCornerOffsetRatio(1.5)
		if link.Product != nil {
			if parentNode, ok := parentByProduct[link.Product]; ok {
				if parentNode.shape.LayoutDirection == Horizontal {
					link.StartOrientation = ORIENTATION_VERTICAL
					link.EndOrientation = ORIENTATION_HORIZONTAL
				} else {
					link.StartOrientation = ORIENTATION_VERTICAL
					link.EndOrientation = ORIENTATION_VERTICAL
				}
			}
		}
	}
}

func layoutProductDFS(node *productNode, currentX float64, currentY float64, margin float64) (float64, float64) {
	node.shape.SetX(currentX)
	node.shape.SetY(currentY)

	w := node.shape.GetWidth()
	h := node.shape.GetHeight()

	if len(node.children) == 0 {
		return currentX + w + margin, currentY + h + margin
	}

	var maxX float64 = currentX + w + margin
	var maxY float64 = currentY + h + margin

	if node.shape.LayoutDirection == Vertical {
		// Children are arranged horizontally.
		childX := currentX
		for _, child := range node.children {
			childMaxX, childMaxY := layoutProductDFS(child, childX, currentY+h*2.0, margin)
			childX = childMaxX
			
			if childMaxX > maxX {
				maxX = childMaxX
			}
			if childMaxY > maxY {
				maxY = childMaxY
			}
		}
		node.shape.SetX(node.children[0].shape.GetX())
		maxX = childX
	} else {
		// Children are arranged vertically (indented tree)
		childY := currentY + h + margin
		for _, child := range node.children {
			childMaxX, childMaxY := layoutProductDFS(child, currentX+w/2.0+margin, childY, margin)
			childY = childMaxY

			if childMaxX > maxX {
				maxX = childMaxX
			}
			if childMaxY > maxY {
				maxY = childMaxY
			}
		}
		// Do not align parent vertically with its children. Parent stays at (currentX, currentY).
		maxY = childY
	}

	return maxX, maxY
}

// Task layout
type taskNode struct {
	shape    *TaskShape
	children []*taskNode
}

func layoutTaskShapes(diagram *Diagram, rootLibrary *Library, nextX *float64, startY float64, margin float64) {
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
			if _, exists := nodesByTask[task.parentTask]; exists {
				isRoot = false
			}
		}

		if isRoot {
			rootNodes = append(rootNodes, node)
		}
	}

	// Populate children respecting SubTasks order
	for _, node := range nodesByTask {
		for _, subTask := range node.shape.Task.SubTasks {
			if childNode, exists := nodesByTask[subTask]; exists {
				node.children = append(node.children, childNode)
			}
		}
	}

	rootTaskIndex := make(map[*Task]int)
	if rootLibrary != nil {
		for i, sub := range rootLibrary.RootTasks {
			rootTaskIndex[sub] = i
		}
	}

	sort.Slice(rootNodes, func(i, j int) bool {
		idxI, okI := rootTaskIndex[rootNodes[i].shape.Task]
		idxJ, okJ := rootTaskIndex[rootNodes[j].shape.Task]
		if okI && okJ {
			return idxI < idxJ
		}
		if okI {
			return true
		}
		if okJ {
			return false
		}
		return rootNodes[i].shape.Task.Name < rootNodes[j].shape.Task.Name
	})

	for _, root := range rootNodes {
		layoutTaskDFS(root, nextX, startY, margin)
	}
	
	for _, link := range diagram.TaskComposition_Shapes {
		link.SetCornerOffsetRatio(1.5)
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

func layoutResourceShapes(diagram *Diagram, rootLibrary *Library, nextX *float64, startY float64, margin float64) {
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
			if _, exists := nodesByResource[resource.parentResource]; exists {
				isRoot = false
			}
		}

		if isRoot {
			rootNodes = append(rootNodes, node)
		}
	}

	// Populate children respecting SubResources order
	for _, node := range nodesByResource {
		for _, subResource := range node.shape.Resource.SubResources {
			if childNode, exists := nodesByResource[subResource]; exists {
				node.children = append(node.children, childNode)
			}
		}
	}

	rootResourceIndex := make(map[*Resource]int)
	if rootLibrary != nil {
		for i, sub := range rootLibrary.RootResources {
			rootResourceIndex[sub] = i
		}
	}

	sort.Slice(rootNodes, func(i, j int) bool {
		idxI, okI := rootResourceIndex[rootNodes[i].shape.Resource]
		idxJ, okJ := rootResourceIndex[rootNodes[j].shape.Resource]
		if okI && okJ {
			return idxI < idxJ
		}
		if okI {
			return true
		}
		if okJ {
			return false
		}
		return rootNodes[i].shape.Resource.Name < rootNodes[j].shape.Resource.Name
	})

	for _, root := range rootNodes {
		layoutResourceDFS(root, nextX, startY, margin)
	}
	
	for _, link := range diagram.ResourceComposition_Shapes {
		link.SetCornerOffsetRatio(1.5)
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
