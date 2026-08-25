package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceDefaultValues() (needCommit bool) {
	architectures := GetGongstrucsSorted[*Architecture](stager.stage)
	if len(architectures) == 0 {
		stager.architecture = (&Architecture{Name: "Architecture"}).Stage(stager.stage)
		needCommit = true
		if stager.probeForm != nil {
			stager.probeForm.AddNotification(time.Now(), "Created default Architecture")
		}
	} else {
		stager.architecture = architectures[0]
	}

	if stager.architecture.NbPixPerCharacter == 0 {
		stager.architecture.NbPixPerCharacter = 8
		needCommit = true
		if stager.probeForm != nil {
			stager.probeForm.AddNotification(time.Now(),
				fmt.Sprintf("Architecture: setting NbPixPerCharacter to %f", stager.architecture.NbPixPerCharacter))
		}
	}

	root := stager.getRootLibrary()
	if root != nil && root.NbPixPerCharacter == 0 {
		root.NbPixPerCharacter = 8
		needCommit = true
	}

	return
}
