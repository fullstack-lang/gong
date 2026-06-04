package models

import (
	"fmt"
	"slices"
	"time"
)

func (stager *Stager) enforceSemantic() (needCommit bool) {
	stage := stager.stage

	pass := 0
	for {
		if stager.enforceSemanticOnePass(false, stage) {
			needCommit = true
			stager.probeForm.AddNotification(time.Now(), fmt.Sprint("Stage was modified to enforce semantic, pass ", pass))
			pass++
		} else {
			break
		}
	}

	if needCommit {
		stage.CommitWithSuspendedCallbacks()
		stager.probeForm.CommitNotificationTable()
	}

	return
}

func (stager *Stager) enforceSemanticOnePass(needCommit bool, stage *Stage) bool {
	needCommit = stage.Clean() || needCommit
	if needCommit {
		stager.probeForm.AddNotification(time.Now(), "Stage clean generated a modification")
	}

	// VERY important because the probe only unstages objects
	// the stage might be dirty
	// with slices of pointer or pointer to unstaged instance
	if stager.stage.Clean() {
		needCommit = true
		if stage.probeIF != nil {
			stager.probeForm.AddNotification(time.Now(), "Stage clean generated a modification")
		}
	}

	// check that there is a Desk
	needCommit = stager.enforce_semantic_singlotons() || needCommit
	needCommit = stager.enforce_semantic_diagrams() || needCommit

	for _, movement := range GetGongstrucsSorted[*Movement](stager.stage) {
		//  movement cannot be minor AND major
		if movement.IsMajor && movement.IsMinor {
			movement.IsMinor = false
			needCommit = true
			if stage.probeIF != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Movement %s cannot be both major and minor, setting IsMinor to false", movement.GetName()))
			}
		}
	}

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
			}
			if _, ok := rm[shape]; !ok {
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
			}
			if _, ok := rm[shape]; !ok {
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
			}
			if _, ok := rm[shape]; !ok {
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
			}
			if _, ok := rm[shape]; !ok {
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

	for _, influence := range GetGongstrucsSorted[*Influence](stager.stage) {
		if influence.SourceMovement != nil {
			influence.source = influence.SourceMovement
		}
		if influence.SourceArtefactType != nil {
			influence.source = influence.SourceArtefactType
		}
		if influence.SourceArtist != nil {
			influence.source = influence.SourceArtist
		}

		if influence.TargetMovement != nil {
			influence.target = influence.TargetMovement
		}
		if influence.TargetArtefactType != nil {
			influence.target = influence.TargetArtefactType
		}
		if influence.TargetArtist != nil {
			influence.target = influence.TargetArtist
		}

		if influence.Name != influence.source.GetName()+" to "+influence.target.GetName() {
			influence.Name = influence.source.GetName() + " to " + influence.target.GetName()
			needCommit = true
			if stage.probeIF != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Influence %s has a name not consistent with its source and target, setting it to %s", influence.GetName(), influence.Name))
			}
		}
	}

	for _, influenceShape := range GetGongstrucsSorted[*InfluenceShape](stager.stage) {
		if influenceShape.Influence == nil {
			influenceShape.Unstage(stager.stage)
			needCommit = true
			if stage.probeIF != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("InfluenceShape %s has no influence, unstaging it", influenceShape.GetName()))
			}
			continue
		}
		if influenceShape.Name != influenceShape.Influence.Name {
			influenceShape.Name = influenceShape.Influence.Name
			needCommit = true
			if stage.probeIF != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("InfluenceShape %s has a name not consistent with its influence, setting it to %s", influenceShape.GetName(), influenceShape.Name))
			}
			continue
		}
	}

	return needCommit
}

func (stager *Stager) enforceThereIsARootLibrary() (needCommit bool) {
	stage := stager.stage
	libraries := GetStructInstancesByOrderAuto[*Library](stage)
	if len(libraries) == 0 {
		rootLibrary := (&Library{Name: "", IsRootLibrary: true}).Stage(stage)
		if stager.probeForm != nil {
			stager.probeForm.AddNotification(time.Now(),
				"Created root library")
		}
		rootLibrary.Stage(stage)
		needCommit = true
		return
	}

	var rootLibrary *Library
	for _, library := range libraries {
		if library.IsRootLibrary {
			if rootLibrary == nil {
				rootLibrary = library
			} else {
				library.IsRootLibrary = false
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(),
						"Detected an extra root library, unsetting it is a root library: "+library.GetName())
				}
				needCommit = true
			}
		}
	}

	if rootLibrary == nil {
		// we set the first library that has no parent as the root library
		hasParent := make(map[*Library]bool)
		for _, library := range libraries {
			for _, sub := range library.SubLibraries {
				hasParent[sub] = true
			}
		}

		var rootCandidates []*Library
		for _, library := range libraries {
			if !hasParent[library] {
				rootCandidates = append(rootCandidates, library)
			}
		}

		if len(rootCandidates) > 0 {
			slices.SortFunc(rootCandidates, CompareGongstructByName[*Library])
			rootLibrary = rootCandidates[0]
			rootLibrary.IsRootLibrary = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					"Set existing library as root: "+rootLibrary.GetName())
			}
			needCommit = true
		} else if len(libraries) > 0 {
			// Fallback in case of circular dependencies
			rootLibrary = libraries[0]
			rootLibrary.IsRootLibrary = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					"Set existing library as root (fallback): "+rootLibrary.GetName())
			}
			needCommit = true
		}
	}

	return
}

func (stager *Stager) GetRootLibrary() (rootLibrary *Library) {
	if stager.getRootLibrary() == nil {
		stager.enforceThereIsARootLibrary()
	}
	return stager.getRootLibrary()
}
