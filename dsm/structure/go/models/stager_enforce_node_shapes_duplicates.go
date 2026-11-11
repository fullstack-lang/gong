package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceNodeShapeDuplicates() (needCommit bool) {
	for _, diagram := range stager.stage.GetInstancesSorted[*DiagramStructure]() {
		needCommit = removeDuplicateNodeShape(stager, &diagram.System_Shapes) || needCommit
		needCommit = removeDuplicateShapeByKey(stager, &diagram.Part_Shapes, func(ps *PartShape) (partShapeKey, bool) {
			if ps == nil || ps.Part == nil {
				return partShapeKey{}, false
			}
			return partShapeKey{part: ps.Part, parentPartShape: ps.ParentPartShape}, true
		}) || needCommit
		needCommit = removeDuplicateShapeByKey(stager, &diagram.Port_Shapes, func(ps *PortShape) (portShapeKey, bool) {
			if ps == nil || ps.Port == nil {
				return portShapeKey{}, false
			}
			return portShapeKey{port: ps.Port, parentPartShape: effectivePortParentPartShape(diagram, ps)}, true
		}) || needCommit
		needCommit = removeDuplicateShapeByKey(stager, &diagram.ControlFlow_Shapes, func(cfs *ControlFlowShape) (controlFlowShapeKey, bool) {
			if cfs == nil || cfs.ControlFlow == nil {
				return controlFlowShapeKey{}, false
			}
			return controlFlowShapeKey{controlFlow: cfs.ControlFlow, parentPartShape: effectiveControlFlowParentPartShape(diagram, cfs)}, true
		}) || needCommit
		needCommit = removeDuplicateShapeByKey(stager, &diagram.DataFlow_Shapes, func(dfs *DataFlowShape) (dataFlowShapeKey, bool) {
			if dfs == nil || dfs.DataFlow == nil {
				return dataFlowShapeKey{}, false
			}
			return dataFlowShapeKey{dataFlow: dfs.DataFlow, parentPartShape: dfs.ParentPartShape}, true
		}) || needCommit
	}
	return
}

func removeDuplicateNodeShape[T ConcreteType](stager *Stager, shapes *[]T) (needCommit bool) {
	seen := make(map[AbstractType]bool)
	var newShapes []T

	for _, shape := range *shapes {
		abstractElement := shape.GetAbstractElement()
		if abstractElement != nil {
			if seen[abstractElement] {
				shape.UnstageVoid(stager.stage)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged duplicate shape %s", shape.GetName()))
				}
				needCommit = true
				continue
			}
			seen[abstractElement] = true
		}
		newShapes = append(newShapes, shape)
	}
	if needCommit {
		*shapes = newShapes
	}
	return
}

func removeDuplicateShapeByKey[T ConcreteType, K comparable](
	stager *Stager,
	shapes *[]T,
	keyFn func(T) (K, bool),
) (needCommit bool) {
	seen := make(map[K]bool)
	var newShapes []T

	for _, shape := range *shapes {
		if key, ok := keyFn(shape); ok {
			if seen[key] {
				shape.UnstageVoid(stager.stage)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged duplicate shape %s", shape.GetName()))
				}
				needCommit = true
				continue
			}
			seen[key] = true
		}
		newShapes = append(newShapes, shape)
	}
	if needCommit {
		*shapes = newShapes
	}
	return
}
