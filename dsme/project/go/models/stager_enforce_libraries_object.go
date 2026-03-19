package models

// enforceLibrariesObject recomputes the objects fields of all libraries
// from the OwningLibrary field of the abstract types.
func (stager *Stager) enforceLibrariesObject() {
	stage := stager.stage

	// Clear the objects slice of all libraries
	for _, library := range GetGongstrucsSorted[*Library](stage) {
		library.objects = nil
	}

	for _, product := range GetGongstrucsSorted[*Product](stage) {
		if product.OwningLibrary != nil {
			product.OwningLibrary.objects = append(product.OwningLibrary.objects, product)
		}
	}

	for _, task := range GetGongstrucsSorted[*Task](stage) {
		if task.OwningLibrary != nil {
			task.OwningLibrary.objects = append(task.OwningLibrary.objects, task)
		}
	}

	for _, resource := range GetGongstrucsSorted[*Resource](stage) {
		if resource.OwningLibrary != nil {
			resource.OwningLibrary.objects = append(resource.OwningLibrary.objects, resource)
		}
	}

	for _, note := range GetGongstrucsSorted[*Note](stage) {
		if note.OwningLibrary != nil {
			note.OwningLibrary.objects = append(note.OwningLibrary.objects, note)
		}
	}

	for _, diagram := range GetGongstrucsSorted[*Diagram](stage) {
		if diagram.OwningLibrary != nil {
			diagram.OwningLibrary.objects = append(diagram.OwningLibrary.objects, diagram)
		}
	}

	for _, lib := range GetGongstrucsSorted[*Library](stage) {
		if lib.OwningLibrary != nil {
			lib.OwningLibrary.objects = append(lib.OwningLibrary.objects, lib)
		}
	}
}
