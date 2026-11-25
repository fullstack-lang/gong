package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceDataStoreSemanticRules() (needCommit bool) {
	stage := stager.stage

	// 1. Ensure each DataStore appears at most once across Library.RootDataStores and remove nil entries.
	seenRootDataStores := make(map[*DataStore]bool)
	for _, lib := range stage.GetInstancesSorted[*Library]() {
		var filtered []*DataStore
		modified := false
		for _, ds := range lib.RootDataStores {
			if ds == nil {
				modified = true
				continue
			}
			if seenRootDataStores[ds] {
				modified = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(),
						fmt.Sprintf("DataStore \"%s\" is duplicated in RootDataStores, removing extra entry from library \"%s\"", ds.GetName(), lib.GetName()))
				}
				continue
			}
			seenRootDataStores[ds] = true
			filtered = append(filtered, ds)
		}
		if modified {
			lib.RootDataStores = filtered
			needCommit = true
		}
	}

	// 2. Clean nil and duplicate Data entries inside DataStore.Datas.
	for _, ds := range stage.GetInstancesSorted[*DataStore]() {
		seenData := make(map[*Data]bool, len(ds.Datas))
		var cleaned []*Data
		modified := false
		for _, d := range ds.Datas {
			if d == nil {
				modified = true
				continue
			}
			if seenData[d] {
				modified = true
				continue
			}
			seenData[d] = true
			cleaned = append(cleaned, d)
		}
		if modified {
			ds.Datas = cleaned
			needCommit = true
		}
	}

	return
}

func (stager *Stager) enforceDataStoreShapesRules() (needCommit bool) {
	for _, diagram := range stager.stage.GetInstancesSorted[*DiagramProcess]() {
		for _, shape := range diagram.DataStore_Shapes {
			if shape == nil || shape.DataStore == nil {
				continue
			}
			if shape.Width <= 0 {
				shape.Width = DataStoreDefaultWidth
				needCommit = true
			} else if shape.Width < DataStoreMinWidth {
				shape.Width = DataStoreMinWidth
				needCommit = true
			}
			if shape.Height <= 0 {
				shape.Height = DataStoreDefaultHeight
				needCommit = true
			} else if shape.Height < DataStoreMinHeight {
				shape.Height = DataStoreMinHeight
				needCommit = true
			}
			expectedName := shape.DataStore.GetName() + "-" + diagram.GetName()
			if shape.Name != expectedName {
				shape.Name = expectedName
				needCommit = true
			}
		}
	}
	return
}
