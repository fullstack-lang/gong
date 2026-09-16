package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceShapeOrphans() (needCommit bool) {
	stage := stager.stage

	// remove orphean movement shapes
	{
		rm := GetSliceOfPointersReverseMap[Diagram, MovementShape](GetAssociationName[Diagram]().MovementShapes[0].Name, stager.stage)
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
		rm := GetSliceOfPointersReverseMap[Diagram, ArtistShape](GetAssociationName[Diagram]().ArtistShapes[0].Name, stager.stage)
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
		rm := GetSliceOfPointersReverseMap[Diagram, ArtefactTypeShape](GetAssociationName[Diagram]().ArtefactTypeShapes[0].Name, stager.stage)
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
		rm := GetSliceOfPointersReverseMap[Diagram, InfluenceShape](GetAssociationName[Diagram]().InfluenceShapes[0].Name, stager.stage)
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
		rm := GetSliceOfPointersReverseMap[InfluenceShape, ControlPointShape](GetAssociationName[InfluenceShape]().ControlPointShapes[0].Name, stager.stage)
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
