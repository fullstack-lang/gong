package models

func (stager *Stager) collectOrphans() (needCommit bool) {

	root := stager.root

	return CollectOrphans(
		GetGongstrucsSorted[*Product](stager.stage),
		func() []*Product {
			roots := make([]*Product, 0)
			for _, project := range GetGongstrucsSorted[*Project](stager.stage) {
				roots = append(roots, project.RootProducts...)
			}
			return roots
		},
		func(product *Product) []*Product {
			return product.SubProducts
		},
		&root.OrphanedProducts,
	)
}

func CollectOrphans[T comparable](
	allNodes []T,
	getRoots func() []T,
	getChildren func(T) []T,
	orphanedSlice *[]T,
) (needCommit bool) {

	// 1. Find all reachable nodes
	reachable := make(map[T]struct{})
	var collectReachable func(node T)
	collectReachable = func(node T) {
		var zero T
		if node == zero {
			return
		}
		if _, ok := reachable[node]; ok {
			return // already visited
		}
		reachable[node] = struct{}{}

		for _, child := range getChildren(node) {
			collectReachable(child)
		}
	}

	for _, root := range getRoots() {
		collectReachable(root)
	}

	// 2. Find all nodes and check for orphans
	orphanNodes := make([]T, 0)
	for _, node := range allNodes {
		if _, ok := reachable[node]; !ok {
			orphanNodes = append(orphanNodes, node)
		}
	}

	// 3. check if the slice has changed (order does not matter)
	if len(orphanNodes) != len(*orphanedSlice) {
		needCommit = true
	} else {
		currentOrphansMap := make(map[T]struct{})
		for _, p := range *orphanedSlice {
			currentOrphansMap[p] = struct{}{}
		}
		for _, p := range orphanNodes {
			if _, ok := currentOrphansMap[p]; !ok {
				needCommit = true
				break
			}
		}
	}

	if needCommit {
		*orphanedSlice = (*orphanedSlice)[:0]
		*orphanedSlice = append(*orphanedSlice, orphanNodes...)
	}

	return
}
