// new file
package models

import "slices"

// enforceCompositionShapes iterates through all diagrams and removes any
// CompositionShape whose associated Product or parent Product is not present
// in the diagram's Product_Shapes. This ensures that links are only displayed
// when both connected nodes are visible.
func (stager *Stager) enforceCompositionShapes() (needCommit bool) {
	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {

		// map of products present in the diagram
		set_Product := make(map[*Product]struct{})
		for _, productShape := range diagram.Product_Shapes {
			if productShape.Product != nil {
				set_Product[productShape.Product] = struct{}{}
			}
		}

		// check that the composition shape is valid
		for i := len(diagram.ProductComposition_Shapes) - 1; i >= 0; i-- {
			compositionShape := diagram.ProductComposition_Shapes[i]
			product := compositionShape.Product

			// check if the product is present in the diagram
			// if not, remove the composition shape
			if _, ok := set_Product[product]; !ok {
				diagram.ProductComposition_Shapes = slices.Delete(diagram.ProductComposition_Shapes, i, i+1)
				compositionShape.Unstage(stager.stage)
				needCommit = true
				continue
			}

			// check if the parent product is present in the diagram
			// if not, remove the composition shape
			parentProduct := product.parentProduct
			if _, ok := set_Product[parentProduct]; !ok {
				diagram.ProductComposition_Shapes = slices.Delete(diagram.ProductComposition_Shapes, i, i+1)
				compositionShape.Unstage(stager.stage)
				needCommit = true
				continue
			}
		}
	}
	return
}
