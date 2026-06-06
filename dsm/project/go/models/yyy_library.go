// generated code (do not edit)
package models

import (
	"log"
	"slices"
	"time"
)

type LibraryAbstractFields struct {
	owningLibrary *Library
}

type LibraryOwnedType interface {
	GetOwningLibrary() *Library
	SetOwningLibrary(library *Library)
}

func (r *LibraryAbstractFields) GetOwningLibrary() *Library {
	return r.owningLibrary
}

func (r *LibraryAbstractFields) SetOwningLibrary(library *Library) {
	r.owningLibrary = library
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

func (stager *Stager) getRootLibrary() *Library {
	for library := range *GetGongstructInstancesSet[Library](stager.stage) {
		if library.IsRootLibrary {
			return library
		}
	}

	// should not happen
	log.Panic("No root library found")
	return nil
}
