package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceShapeOrphans() (needCommit bool) {
	// 1. collect all shapes that are attached to a diagram
	reachablePartShapes := make(map[*PartShape]struct{})
	reachableLinkAssociationShapes := make(map[*LinkAssociationShape]struct{})

	for _, diagram := range GetGongstrucsSorted[*DiagramStructure](stager.stage) {
		collectShapes(diagram.Part_Shapes, reachablePartShapes)
		collectShapes(diagram.Link_Shapes, reachableLinkAssociationShapes)
	}

	// 2. unstage shapes that are not attached to a diagram
	needCommit = unstageUnreachableOrphans(stager, reachablePartShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableLinkAssociationShapes) || needCommit

	return
}

func collectShapes[T comparable](shapes []T, reachable map[T]struct{}) {
	for _, shape := range shapes {
		reachable[shape] = struct{}{}
	}
}

func unstageUnreachableOrphans[T PointerToGongstruct](stager *Stager, reachable map[T]struct{}) (needCommit bool) {
	for _, object := range GetGongstrucsSorted[T](stager.stage) {
		if _, ok := reachable[object]; !ok {
			object.UnstageVoid(stager.stage)
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaging orphan shape object \"%s\" of type \"%s\"",
					object.GetName(), object.GongGetGongstructName()))
			}
		}
	}
	return
}
