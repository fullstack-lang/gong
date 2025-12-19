package models

func (stager *Stager) collectOrphans() (needCommit bool) {

	root := stager.root

	// 1. Find all reachable products
	reachableProducts := make(map[*Product]struct{})
	var collectReachableProducts func(product *Product)
	collectReachableProducts = func(product *Product) {
		if product == nil {
			return
		}
		if _, ok := reachableProducts[product]; ok {
			return // already visited
		}
		reachableProducts[product] = struct{}{}

		for _, subProduct := range product.SubProducts {
			collectReachableProducts(subProduct)
		}
	}

	for _, project := range GetGongstrucsSorted[*Project](stager.stage) {
		for _, rootProduct := range project.RootProducts {
			collectReachableProducts(rootProduct)
		}
	}

	// 2. Find all products and check for orphans
	orphanProducts := make([]*Product, 0)
	allProducts := GetGongstrucsSorted[*Product](stager.stage) // to have a deterministic order
	for _, product := range allProducts {
		if _, ok := reachableProducts[product]; !ok {
			orphanProducts = append(orphanProducts, product)
		}
	}

	// 3. check if the slice has changed (order does not matter)
	if len(orphanProducts) != len(root.OrphanedProducts) {
		needCommit = true
	} else {
		currentOrphansMap := make(map[*Product]struct{})
		for _, p := range root.OrphanedProducts {
			currentOrphansMap[p] = struct{}{}
		}
		for _, p := range orphanProducts {
			if _, ok := currentOrphansMap[p]; !ok {
				needCommit = true
				break
			}
		}
	}

	if needCommit {
		root.OrphanedProducts = root.OrphanedProducts[:0]
		root.OrphanedProducts = append(root.OrphanedProducts, orphanProducts...)
	}

	return
}
