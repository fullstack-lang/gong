package models

func (stager *Stager) collectOrphans() (needCommit bool) {
	needCommit = unstageOrphans(
		stager,
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
	)

	needCommit = needCommit || unstageOrphans(
		stager,
		func() []*Task {
			roots := make([]*Task, 0)
			for _, task := range GetGongstrucsSorted[*Project](stager.stage) {
				roots = append(roots, task.RootTasks...)
			}
			return roots
		},
		func(product *Task) []*Task {
			return product.SubTasks
		},
	)

	return
}

func unstageOrphans[T PointerToGongstruct](
	stager *Stager,
	getRoots func() []T,
	getChildren func(T) []T,
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

	// 2. Find all nodes and delete them
	unstageUnreachableOrphans(stager, reachable)

	return
}
