package models

import "log"

func (stager *Stager) assignProbeCallback() bool {
	if stager.probeForm == nil {
		return false
	}

	// We only assign it once. But since it's an assignment, doing it every time is harmless,
	// though we might want to ensure it's not nil before assigning.
	stager.probeForm.SetUpdateSliceOfPointersCallback(func(instance any, fieldName string, slicePtr any) {
		if dataFlow, ok := instance.(*DataFlow); ok {
			if slicePtr == &dataFlow.Datas {
				stager.onDataFlowDatasUpdate(dataFlow)
			}
		}
	})

	return false
}

func (stager *Stager) onDataFlowDatasUpdate(dataFlow *DataFlow) {
	// when a Data is added/removed to/from a DataFlow
	// update all diagrams accordingly
	for _, diagram := range GetGongstrucsSorted[*DiagramProcess](stager.stage) {
		// check if this diagram has a DataFlowShape for this dataFlow
		var hasDataFlowShape bool
		for _, dataFlowShape := range diagram.DataFlow_Shapes {
			if dataFlowShape.DataFlow == dataFlow {
				hasDataFlowShape = true
				break
			}
		}

		if hasDataFlowShape {
			existingDataShapes := make(map[*Data]*DataShape)
			var validDataShapes []*DataShape

			for _, dataShape := range diagram.Data_Shapes {
				if dataShape.DataFlow == dataFlow {
					if dataShape.Data != nil {
						existingDataShapes[dataShape.Data] = dataShape

						stillExists := false
						for _, d := range dataFlow.Datas {
							if d == dataShape.Data {
								stillExists = true
								break
							}
						}
						if stillExists {
							validDataShapes = append(validDataShapes, dataShape)
						} else {
							dataShape.Unstage(stager.stage)
						}
					} else {
						dataShape.Unstage(stager.stage)
					}
				} else {
					validDataShapes = append(validDataShapes, dataShape)
				}
			}

			diagram.Data_Shapes = validDataShapes

			// add missing DataShapes
			for _, data := range dataFlow.Datas {
				if _, ok := existingDataShapes[data]; !ok {
					// create a new DataShape
					newDataShape := new(DataShape).Stage(stager.stage)
					newDataShape.Name = data.Name
					newDataShape.Data = data
					newDataShape.DataFlow = dataFlow
					diagram.Data_Shapes = append(diagram.Data_Shapes, newDataShape)
				}
			}
		}
	}

	// The callback is invoked synchronously inside OnSave of the form,
	// and OnSave calls commit *after* the callback. But in case it doesn't,
	// we do not need to call commit explicitly if the form is saving.
	// However, if we modified diagram.Data_Shapes, we might want to log it.
	log.Printf("Updated Data_Shapes for DataFlow %s", dataFlow.Name)
}
