package models

import "time"

func (stager *Stager) enforceAProcessDiagramHasItsOwningProcess() (needCommit bool) {

	rm := GetSliceOfPointersReverseMap[Process, DiagramProcess](GetAssociationName[Process]().DiagramProcesss[0].Name, stager.stage)

	for diagramProcess := range *GetGongstructInstancesSetFromPointerType[*DiagramProcess](stager.stage) {

		owningProcesses := rm[diagramProcess]
		if len(owningProcesses) > 1 {
			stager.probeForm.AddNotification(time.Now(), "DiagramProcess "+diagramProcess.Name+" has more than one owning process")
			needCommit = true
			diagramProcess.UnstageVoid(stager.stage)
			continue
		}

		if len(owningProcesses) == 0 {
			stager.probeForm.AddNotification(time.Now(), "DiagramProcess "+diagramProcess.Name+" has no owning process")
			diagramProcess.UnstageVoid(stager.stage)
			needCommit = true
			continue
		}

		diagramProcess.owningProcess = owningProcesses[0]

		// check that that there is a processShape that have the process as its process
		isProcessShapeOfOwningProcessFound := false
		for _, processShape := range diagramProcess.Process_Shapes {
			if processShape.Process == diagramProcess.owningProcess {
				isProcessShapeOfOwningProcessFound = true
				break
			}
		}

		if !isProcessShapeOfOwningProcessFound {
			stager.probeForm.AddNotification(time.Now(), "DiagramProcess "+diagramProcess.Name+" has no process shape for its owning process "+diagramProcess.owningProcess.Name)
			processShape := (&ProcessShape{
				Name:    "ProcessShape",
				Process: diagramProcess.owningProcess,
				RectShape: RectShape{
					X:      100,
					Y:      50,
					Width:  500,
					Height: 1000,
				},
			}).Stage(stager.stage)
			diagramProcess.Process_Shapes = append(diagramProcess.Process_Shapes, processShape)
			needCommit = true
			continue
		}

	}
	return
}
