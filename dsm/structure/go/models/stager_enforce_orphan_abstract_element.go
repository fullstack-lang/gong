package models

func (stager *Stager) enforceOrphansAbstractElement() (needCommit bool) {
	needCommit = reattachToLibraryRoots(
		stager,
		func() []*System {
			roots := make([]*System, 0)
			for _, library := range stager.stage.GetInstancesSorted[*Library]() {
				roots = append(roots, library.RootSystemes...)
			}
			return roots
		},
		func(system *System) {
			system.GetOwningLibrary().RootSystemes = append(system.GetOwningLibrary().RootSystemes, system)
		},
		func(system *System) []*System {
			return system.SubSystemes
		},
	)

	if reattachToLibraryRoots(
		stager,
		func() []*DataFlow {
			roots := make([]*DataFlow, 0)
			for _, library := range stager.stage.GetInstancesSorted[*Library]() {
				roots = append(roots, library.RootDataFlows...)
			}
			return roots
		},
		func(dataflow *DataFlow) {
			dataflow.GetOwningLibrary().RootDataFlows = append(dataflow.GetOwningLibrary().RootDataFlows, dataflow)
		},
		func(dataflow *DataFlow) []*DataFlow {
			return []*DataFlow{}
		},
	) {
		needCommit = true
	}

	if reattachToLibraryRoots(
		stager,
		func() []*Protocol {
			roots := make([]*Protocol, 0)
			for _, library := range stager.stage.GetInstancesSorted[*Library]() {
				roots = append(roots, library.RootProtocols...)
			}
			return roots
		},
		func(protocol *Protocol) {
			protocol.GetOwningLibrary().RootProtocols = append(protocol.GetOwningLibrary().RootProtocols, protocol)
		},
		func(protocol *Protocol) []*Protocol {
			return []*Protocol{}
		},
	) {
		needCommit = true
	}

	if reattachToLibraryRoots(
		stager,
		func() []*ProtocolLayer {
			roots := make([]*ProtocolLayer, 0)
			for _, library := range stager.stage.GetInstancesSorted[*Library]() {
				roots = append(roots, library.RootProtocolLayers...)
			}
			return roots
		},
		func(protocolLayer *ProtocolLayer) {
			protocolLayer.GetOwningLibrary().RootProtocolLayers = append(protocolLayer.GetOwningLibrary().RootProtocolLayers, protocolLayer)
		},
		func(protocolLayer *ProtocolLayer) []*ProtocolLayer {
			return []*ProtocolLayer{}
		},
	) {
		needCommit = true
	}

	if reattachToLibraryRoots(
		stager,
		func() []*Library {
			return stager.getRootLibrary().SubLibraries
		},
		func(library *Library) {
			// attach to root, only if it is not the root library
			// (which is the only one without an owning library)
			if library != stager.getRootLibrary() {
				stager.getRootLibrary().SubLibraries = append(stager.getRootLibrary().SubLibraries, library)
				library.SetOwningLibrary(stager.getRootLibrary())
			}
		},
		func(library *Library) []*Library {
			return library.SubLibraries
		},
	) {
		needCommit = true
	}

	return
}
