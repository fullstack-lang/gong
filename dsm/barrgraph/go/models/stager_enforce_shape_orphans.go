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
		for _, shape := range GetGongstrucsSorted[*MovementShape](stager.stage) {
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
		for _, shape := range GetGongstrucsSorted[*ArtistShape](stager.stage) {
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
		for _, shape := range GetGongstrucsSorted[*ArtefactTypeShape](stager.stage) {
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
		for _, shape := range GetGongstrucsSorted[*InfluenceShape](stager.stage) {
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
		for _, shape := range GetGongstrucsSorted[*ControlPointShape](stager.stage) {
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
