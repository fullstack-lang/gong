package models

func (stager *Stager) enforceSemantic() (needCommit bool) {
	stage := stager.stage

	pass := 0
	for {
		if stager.enforceSemanticOnePass(false, stage) {
			needCommit = true
			pass++
		} else {
			break
		}
	}

	return
}

func (stager *Stager) enforceSemanticOnePass(needCommit bool, stage *Stage) bool {
	// VERY important because the probe only unstages objects
	// this is the Clean that delete them from slices and pointers that reference
	// them. If the checkout is not performed, the stage might be dirty
	// with slices of pointer or pointer to unstaged instance
	needCommit = stage.Clean() || needCommit

	return needCommit
}
