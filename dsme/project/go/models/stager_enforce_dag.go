package models

import "slices"

// Enforce DAG
// All [models.Product] are within at least one DAG
// This function check that there is no cycles within the DAG.
// If a cycle is found the culprit association is deleted
func (stager *Stager) enforceDAG() (needCommit bool) {

	// inspired by https://www.geeksforgeeks.org/detect-cycle-in-a-graph/
	whiteSet := make(map[*Product]struct{})
	graySet := make(map[*Product]struct{})
	blackSet := make(map[*Product]struct{})

	for product := range stager.stage.Products {
		whiteSet[product] = struct{}{}
	}

	var dfs func(product *Product)

	dfs = func(product *Product) {
		// move product from white to gray
		delete(whiteSet, product)
		graySet[product] = struct{}{}

		// we iterate on a copy of sub products because the original slice can be modified
		subProductsCopy := make([]*Product, len(product.SubProducts))
		copy(subProductsCopy, product.SubProducts)

		for _, subProduct := range subProductsCopy {
			// if the subProduct is in the gray set, we have a cycle
			if _, ok := graySet[subProduct]; ok {
				// we need to remove the edge from product to subProduct
				// we find the index of subProduct in the original slice
				for j, p := range product.SubProducts {
					if p == subProduct {
						product.SubProducts = slices.Delete(product.SubProducts, j, j+1)
						needCommit = true
						break // we found it, we can break
					}
				}
				continue
			}

			// if the subProduct is in the black set, it means it has been visited
			// and there is no cycle from it, so we can skip it
			if _, ok := blackSet[subProduct]; ok {
				continue
			}

			// if the subProduct is in the white set, we visit it
			if _, ok := whiteSet[subProduct]; ok {
				dfs(subProduct)
			}
		}

		// move product from gray to black
		delete(graySet, product)
		blackSet[product] = struct{}{}
	}

	products := make([]*Product, 0, len(stager.stage.Products))
	for product := range stager.stage.Products {
		products = append(products, product)
	}

	for _, product := range products {
		if _, ok := whiteSet[product]; ok {
			dfs(product)
		}
	}

	return
}
