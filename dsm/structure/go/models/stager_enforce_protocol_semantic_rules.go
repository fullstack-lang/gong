package models

import (
	"fmt"
	"slices"
	"time"
)

func (stager *Stager) enforceProtocolSemanticRules() (needCommit bool) {
	stage := stager.stage

	sortedDFs := stage.GetInstancesSorted[*DataFlow]()
	sortedPorts := stage.GetInstancesSorted[*Port]()

	// Build owning library maps for DataFlows and Parts/Ports
	dfToLibrary := make(map[*DataFlow]*Library)
	for _, lib := range stage.GetInstancesSorted[*Library]() {
		for _, df := range lib.RootDataFlows {
			if df != nil {
				dfToLibrary[df] = lib
			}
		}
	}
	_, partToLibrary := stager.buildSystemAndPartLibraryMaps()
	portToLibrary := make(map[*Port]*Library)
	for _, part := range stage.GetInstancesSorted[*Part]() {
		lib := partToLibrary[part]
		if lib == nil {
			lib = part.GetOwningLibrary()
		}
		for _, port := range part.Ports {
			if port != nil {
				portToLibrary[port] = lib
			}
		}
	}

	isDFRelaxed := func(df *DataFlow) bool {
		lib := dfToLibrary[df]
		if lib == nil {
			lib = df.GetOwningLibrary()
		}
		return lib != nil && lib.IsPortAndDataFlowProtocolConsistencyRelaxed
	}
	isPortRelaxed := func(p *Port) bool {
		lib := portToLibrary[p]
		if lib == nil {
			lib = p.GetOwningLibrary()
		}
		return lib != nil && lib.IsPortAndDataFlowProtocolConsistencyRelaxed
	}

	dfPorts := func(df *DataFlow) []*Port {
		if isDFRelaxed(df) {
			return nil
		}
		var res []*Port
		switch df.Type {
		case DataFlow_Port2Port:
			if df.StartPort != nil && !isPortRelaxed(df.StartPort) {
				res = append(res, df.StartPort)
			}
			if df.EndPort != nil && df.EndPort != df.StartPort && !isPortRelaxed(df.EndPort) {
				res = append(res, df.EndPort)
			}
		case DataFlow_ExternalPart2Port:
			if df.EndPort != nil && !isPortRelaxed(df.EndPort) {
				res = append(res, df.EndPort)
			}
		case DataFlow_Port2ExternalPart:
			if df.StartPort != nil && !isPortRelaxed(df.StartPort) {
				res = append(res, df.StartPort)
			}
		}
		return res
	}

	portToDFs := make(map[*Port][]*DataFlow)
	for _, df := range sortedDFs {
		for _, p := range dfPorts(df) {
			portToDFs[p] = append(portToDFs[p], df)
		}
	}

	visitedDF := make(map[*DataFlow]bool)
	visitedPort := make(map[*Port]bool)

	for _, rootDF := range sortedDFs {
		if visitedDF[rootDF] {
			continue
		}
		portsOfRoot := dfPorts(rootDF)
		if len(portsOfRoot) == 0 {
			visitedDF[rootDF] = true
			continue
		}

		var compDFs []*DataFlow
		var compPorts []*Port
		queueDF := []*DataFlow{rootDF}
		visitedDF[rootDF] = true

		for len(queueDF) > 0 {
			currDF := queueDF[0]
			queueDF = queueDF[1:]
			compDFs = append(compDFs, currDF)

			for _, p := range dfPorts(currDF) {
				if !visitedPort[p] {
					visitedPort[p] = true
					compPorts = append(compPorts, p)
					for _, nextDF := range portToDFs[p] {
						if !visitedDF[nextDF] {
							visitedDF[nextDF] = true
							queueDF = append(queueDF, nextDF)
						}
					}
				}
			}
		}

		slices.SortFunc(compDFs, GongCompareGongstructByName[*DataFlow])
		slices.SortFunc(compPorts, GongCompareGongstructByName[*Port])

		var targetProtocol *Protocol
		var hasTarget bool

		for _, df := range compDFs {
			if stager.map_DataFlow_ProtocolTracked[df] && df.Protocol != stager.map_DataFlow_PrevProtocol[df] {
				targetProtocol = df.Protocol
				hasTarget = true
				break
			}
		}
		if !hasTarget {
			for _, p := range compPorts {
				if stager.map_Port_ProtocolTracked[p] && p.Protocol != stager.map_Port_PrevProtocol[p] {
					targetProtocol = p.Protocol
					hasTarget = true
					break
				}
			}
		}
		if !hasTarget {
			for _, df := range compDFs {
				if df.Protocol != nil {
					targetProtocol = df.Protocol
					hasTarget = true
					break
				}
			}
		}
		if !hasTarget {
			for _, p := range compPorts {
				if p.Protocol != nil {
					targetProtocol = p.Protocol
					hasTarget = true
					break
				}
			}
		}

		if !hasTarget {
			continue
		}

		protoLabel := "<nil>"
		if targetProtocol != nil {
			protoLabel = targetProtocol.GetName()
		}

		for _, df := range compDFs {
			if df.Protocol != targetProtocol {
				df.Protocol = targetProtocol
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(),
						fmt.Sprintf("Data flow \"%s\" protocol synchronized to \"%s\"", df.GetName(), protoLabel))
				}
			}
		}
		for _, p := range compPorts {
			if p.Protocol != targetProtocol {
				p.Protocol = targetProtocol
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(),
						fmt.Sprintf("Port \"%s\" protocol synchronized to \"%s\"", p.GetName(), protoLabel))
				}
			}
		}
	}

	stager.map_DataFlow_PrevProtocol = make(map[*DataFlow]*Protocol, len(sortedDFs))
	stager.map_DataFlow_ProtocolTracked = make(map[*DataFlow]bool, len(sortedDFs))
	for _, df := range sortedDFs {
		stager.map_DataFlow_PrevProtocol[df] = df.Protocol
		stager.map_DataFlow_ProtocolTracked[df] = true
	}

	stager.map_Port_PrevProtocol = make(map[*Port]*Protocol, len(sortedPorts))
	stager.map_Port_ProtocolTracked = make(map[*Port]bool, len(sortedPorts))
	for _, p := range sortedPorts {
		stager.map_Port_PrevProtocol[p] = p.Protocol
		stager.map_Port_ProtocolTracked[p] = true
	}

	return
}
