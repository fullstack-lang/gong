package models

func (stager *Stager) enforceSemantic() {

	stage := stager.stage

	// VERY important because the probe only unstages objects
	// this is the Clean that delete them from slices and pointers that reference
	// them. If the checkout is not performed, the stage might be dirty
	// with slices of pointer or pointer to unstaged instance
	stage.Clean()

	needCommit := false

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

	if needCommit {
		stager.stage.Clean()
		stager.stage.CommitWithSuspendedCallbacks()
	}
}
