package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceShapeOrphans() (needCommit bool) {
	stage := stager.stage

	// remove orphean movement shapes
	{
		rm := stager.stage.GetSliceOfPointersReverseMap[Diagram, MovementShape](GongGetAssociationName[Diagram]().MovementShapes[0].Name)
		for _, shape := range stager.stage.GetInstancesSorted[*MovementShape]() {
			if shape.GetArtElement() == nil {
				shape.Unstage(stager.stage)
				needCommit = true
				if stage.probeIF != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("MovementShape %s is orphan, unstaging it", shape.GetName()))
				}
			} else if _, ok := rm[shape]; !ok {
				shape.Unstage(stager.stage)
				needCommit = true
				if stage.probeIF != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("MovementShape %s is orphan, unstaging it", shape.GetName()))
				}
			}
		}
	}
	{
		rm := stager.stage.GetSliceOfPointersReverseMap[Diagram, ArtistShape](GongGetAssociationName[Diagram]().ArtistShapes[0].Name)
		for _, shape := range stager.stage.GetInstancesSorted[*ArtistShape]() {
			if shape.GetArtElement() == nil {
				shape.Unstage(stager.stage)
				needCommit = true
				if stage.probeIF != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("ArtistShape %s is orphan, unstaging it", shape.GetName()))
				}
			} else if _, ok := rm[shape]; !ok {
				shape.Unstage(stager.stage)
				needCommit = true
				if stage.probeIF != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("ArtistShape %s is orphan, unstaging it", shape.GetName()))
				}
			}
		}
	}
	{
		rm := stager.stage.GetSliceOfPointersReverseMap[Diagram, ArtefactTypeShape](GongGetAssociationName[Diagram]().ArtefactTypeShapes[0].Name)
		for _, shape := range stager.stage.GetInstancesSorted[*ArtefactTypeShape]() {
			if shape.GetArtElement() == nil {
				shape.Unstage(stager.stage)
				needCommit = true
				if stage.probeIF != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("ArtefactTypeShape %s is orphan, unstaging it", shape.GetName()))
				}
			} else if _, ok := rm[shape]; !ok {
				shape.Unstage(stager.stage)
				needCommit = true
				if stage.probeIF != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("ArtefactTypeShape %s is orphan, unstaging it", shape.GetName()))
				}
			}
		}
	}
	{
		rm := stager.stage.GetSliceOfPointersReverseMap[Diagram, InfluenceShape](GongGetAssociationName[Diagram]().InfluenceShapes[0].Name)
		for _, shape := range stager.stage.GetInstancesSorted[*InfluenceShape]() {
			if shape.GetArtElement() == nil {
				shape.Unstage(stager.stage)
				needCommit = true
				if stage.probeIF != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("InfluenceShape %s is orphan, unstaging it", shape.GetName()))
				}
			} else if _, ok := rm[shape]; !ok {
				shape.Unstage(stager.stage)
				needCommit = true
				if stage.probeIF != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("InfluenceShape %s is orphan, unstaging it", shape.GetName()))
				}
			}
		}
	}
	{
		rm := stager.stage.GetSliceOfPointersReverseMap[InfluenceShape, ControlPointShape](GongGetAssociationName[InfluenceShape]().ControlPointShapes[0].Name)
		for _, shape := range stager.stage.GetInstancesSorted[*ControlPointShape]() {
			if _, ok := rm[shape]; !ok {
				shape.Unstage(stager.stage)
				needCommit = true
				if stage.probeIF != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("ControlPointShape %s is orphan, unstaging it", shape.GetName()))
				}
			}
		}
	}

	return
}
