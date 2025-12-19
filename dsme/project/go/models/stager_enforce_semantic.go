package models

import "slices"

func (stager *Stager) enforceSemantic() (needCommit bool) {

	stage := stager.stage

	// VERY important because the probe only unstages objects
	// this is the Clean that delete them from slices and pointers that reference
	// them. If the checkout is not performed, the stage might be dirty
	// with slices of pointer or pointer to unstaged instance
	stage.Clean()

	// Ensures that there is one and only one root
	// prune the other
	// check that there is at least one root
	// and that one can safely access stager.root
	roots := GetGongstrucsSorted[*Root](stager.stage)
	if len(roots) == 0 {
		stager.root = (&Root{Name: "Root"}).Stage(stager.stage)
		needCommit = true
	} else {
		stager.root = roots[0]

		if len(roots) > 1 {
			for _, root := range roots[1:] {
				root.Unstage(stage)
			}
		}
	}

	root := stager.root

	// Enforce that all projects are appended to the [root]
	// if one project is not appended, append it
	for _, project := range GetGongstrucsSorted[*Project](stage) {

		if slices.Index(root.Projects, project) == -1 {
			root.Projects = append(root.Projects, project)
		}
	}

	if needCommit {
		stager.stage.Clean()
		stager.stage.CommitWithSuspendedCallbacks()
	}

	if stager.enforceDAG() {
		needCommit = true
	}
	if stager.collectOrphans() {
		needCommit = true
	}

	return
}
