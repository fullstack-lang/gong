package models

import (
	"sort"
)

type layoutNode[AT interface {
	AbstractType
	comparable
}, CT LayoutConcreteType] struct {
	shape    CT
	children []*layoutNode[AT, CT]
	parent   *layoutNode[AT, CT]
}

func layoutGenericShapes[AT interface {
	AbstractType
	comparable
}, CT LayoutConcreteType, ACT LayoutAssociationType](
	shapes []CT,
	rootElements []AT,
	getChildren func(AT) []AT,
	getParent func(AT) AT,
	compositionShapes []ACT,
	getLinkElement func(ACT) AT,
	getName func(AT) string,
	nextX *float64,
	startY float64,
	margin float64,
) {
	nodesByElement := make(map[AT]*layoutNode[AT, CT])
	var rootNodes []*layoutNode[AT, CT]

	for _, shape := range shapes {
		// Type assert to get the AT
		abstractElement := shape.GetAbstractElement().(AT)
		nodesByElement[abstractElement] = &layoutNode[AT, CT]{shape: shape}
	}

	for _, shape := range shapes {
		abstractElement := shape.GetAbstractElement().(AT)
		node := nodesByElement[abstractElement]
		isRoot := true

		parent := getParent(abstractElement)
		// Go interfaces comparison against nil is tricky. We can check if it's valid.
		// Let's assume parent != nil implies it has a parent if it exists in the map
		if any(parent) != nil {
			if _, exists := nodesByElement[parent]; exists {
				isRoot = false
			}
		}

		if isRoot {
			rootNodes = append(rootNodes, node)
		}
	}

	// Populate children respecting order
	for _, node := range nodesByElement {
		abstractElement := node.shape.GetAbstractElement().(AT)
		for _, subElement := range getChildren(abstractElement) {
			if childNode, exists := nodesByElement[subElement]; exists {
				node.children = append(node.children, childNode)
				childNode.parent = node
			}
		}
	}

	rootElementIndex := make(map[AT]int)
	for i, sub := range rootElements {
		rootElementIndex[sub] = i
	}

	sort.Slice(rootNodes, func(i, j int) bool {
		idxI, okI := rootElementIndex[rootNodes[i].shape.GetAbstractElement().(AT)]
		idxJ, okJ := rootElementIndex[rootNodes[j].shape.GetAbstractElement().(AT)]
		if okI && okJ {
			return idxI < idxJ
		}
		if okI {
			return true
		}
		if okJ {
			return false
		}
		return getName(rootNodes[i].shape.GetAbstractElement().(AT)) < getName(rootNodes[j].shape.GetAbstractElement().(AT))
	})

	for _, root := range rootNodes {
		maxX, _ := layoutGenericDFS(root, *nextX, startY, margin)
		*nextX = maxX
	}

	// Find the parent of each element
	parentByElement := make(map[AT]*layoutNode[AT, CT])
	for _, node := range nodesByElement {
		abstractElement := node.shape.GetAbstractElement().(AT)
		for _, subElement := range getChildren(abstractElement) {
			parentByElement[subElement] = node
		}
	}

	for _, link := range compositionShapes {
		link.SetCornerOffsetRatio(1.5)
		element := getLinkElement(link)
		if any(element) != nil {
			if parentNode, ok := parentByElement[element]; ok {
				layoutDirection := parentNode.shape.GetAbstractElement().GetLayoutDirection()
				if parentNode.shape.GetOverideLayoutDirection() {
					layoutDirection = parentNode.shape.GetConcreteLayoutDirection()
				}
				if layoutDirection == Horizontal {
					link.SetStartOrientation(ORIENTATION_VERTICAL)
					link.SetEndOrientation(ORIENTATION_HORIZONTAL)
				} else {
					link.SetStartOrientation(ORIENTATION_VERTICAL)
					link.SetEndOrientation(ORIENTATION_VERTICAL)
				}
			}
		}
	}
}

func layoutGenericDFS[AT interface {
	AbstractType
	comparable
}, CT LayoutConcreteType](node *layoutNode[AT, CT], currentX float64, currentY float64, margin float64) (float64, float64) {
	node.shape.SetX(currentX)
	node.shape.SetY(currentY)

	w := node.shape.GetWidth()
	h := node.shape.GetHeight()

	if len(node.children) == 0 {
		return currentX + w + margin, currentY + h + margin
	}

	var maxX float64 = currentX + w + margin
	var maxY float64 = currentY + h + margin

	layoutDirection := node.shape.GetAbstractElement().GetLayoutDirection()
	if node.shape.GetOverideLayoutDirection() {
		layoutDirection = node.shape.GetConcreteLayoutDirection()
	}

	if layoutDirection == Vertical {
		// Children are arranged horizontally.
		childX := currentX

		isParentHorizontal := false
		if node.parent != nil {
			parentLayout := node.parent.shape.GetAbstractElement().GetLayoutDirection()
			if node.parent.shape.GetOverideLayoutDirection() {
				parentLayout = node.parent.shape.GetConcreteLayoutDirection()
			}
			if parentLayout == Horizontal {
				isParentHorizontal = true
			}
		}

		if isParentHorizontal {
			childX = currentX + w/2.0 + margin
		}

		for _, child := range node.children {
			childMaxX, childMaxY := layoutGenericDFS(child, childX, currentY+h*2.0, margin)
			childX = childMaxX

			if childMaxX > maxX {
				maxX = childMaxX
			}
			if childMaxY > maxY {
				maxY = childMaxY
			}
		}
		if !isParentHorizontal {
			node.shape.SetX(node.children[0].shape.GetX())
		}
		maxX = childX
	} else {
		// Children are arranged vertically (indented tree)
		verticalMargin := 15.0
		childY := currentY + h + verticalMargin
		for _, child := range node.children {
			childMaxX, childMaxY := layoutGenericDFS(child, currentX+w/2.0+margin, childY, margin)
			childY = (childMaxY - margin) + verticalMargin

			if childMaxX > maxX {
				maxX = childMaxX
			}
			if childMaxY > maxY {
				maxY = childMaxY
			}
		}
		// Do not align parent vertically with its children. Parent stays at (currentX, currentY).
		maxY = (childY - verticalMargin) + margin
	}

	return maxX, maxY
}

func layoutDiagram(diagram *Diagram, stager *Stager) {
	rootLibrary := stager.getRootLibrary()

	nextX := 50.0
	startY := 50.0
	margin := 50.0

	var rootProducts []*Product
	var rootTasks []*Task
	var rootResources []*Resource

	if rootLibrary != nil {
		rootProducts = rootLibrary.RootProducts
		rootTasks = rootLibrary.RootTasks
		rootResources = rootLibrary.RootResources
	}

	// Layout Products
	var productShapes []*ProductShape
	for _, shape := range diagram.Product_Shapes {
		productShapes = append(productShapes, shape)
	}
	var productLinks []*ProductCompositionShape
	for _, link := range diagram.ProductComposition_Shapes {
		productLinks = append(productLinks, link)
	}

	layoutGenericShapes(
		productShapes,
		rootProducts,
		func(p *Product) []*Product { return p.SubProducts },
		func(p *Product) *Product { return p.parentProduct },
		productLinks,
		func(l *ProductCompositionShape) *Product { return l.Product },
		func(p *Product) string { return p.Name },
		&nextX, startY, margin,
	)

	// Layout Tasks
	var taskShapes []*TaskShape
	for _, shape := range diagram.Task_Shapes {
		taskShapes = append(taskShapes, shape)
	}
	var taskLinks []*TaskCompositionShape
	for _, link := range diagram.TaskComposition_Shapes {
		taskLinks = append(taskLinks, link)
	}

	layoutGenericShapes(
		taskShapes,
		rootTasks,
		func(t *Task) []*Task { return t.SubTasks },
		func(t *Task) *Task { return t.parentTask },
		taskLinks,
		func(l *TaskCompositionShape) *Task { return l.Task },
		func(t *Task) string { return t.Name },
		&nextX, startY, margin,
	)

	// Layout Resources
	var resourceShapes []*ResourceShape
	for _, shape := range diagram.Resource_Shapes {
		resourceShapes = append(resourceShapes, shape)
	}
	var resourceLinks []*ResourceCompositionShape
	for _, link := range diagram.ResourceComposition_Shapes {
		resourceLinks = append(resourceLinks, link)
	}

	layoutGenericShapes(
		resourceShapes,
		rootResources,
		func(r *Resource) []*Resource { return r.SubResources },
		func(r *Resource) *Resource { return r.parentResource },
		resourceLinks,
		func(l *ResourceCompositionShape) *Resource { return l.Resource },
		func(r *Resource) string { return r.Name },
		&nextX, startY, margin,
	)
}
