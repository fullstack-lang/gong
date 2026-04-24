package models

func (stager *Stager) enforceAtLeastOneDiagramPerProcess() (needCommit bool) {
	stage := stager.stage

	// enforce that there is at least one diagram per process
	for process := range *GetGongstructInstancesSetFromPointerType[*Process](stage) {
		if len(process.DiagramProcesss) == 0 {
			diagramProcess := (&DiagramProcess{
				Name: "DiagramProcess",
			}).Stage(stage)
			process.DiagramProcesss = append(process.DiagramProcesss, diagramProcess)

			processShape := (&ProcessShape{
				Name:    "ProcessShape",
				Process: process,
				RectShape: RectShape{
					X:      100,
					Y:      50,
					Width:  500,
					Height: 1000,
				},
			}).Stage(stage)
			diagramProcess.Process_Shapes = append(diagramProcess.Process_Shapes, processShape)

			needCommit = true
		}
	}

	return
}
