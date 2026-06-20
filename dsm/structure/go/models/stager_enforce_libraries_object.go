package models

// enforceOwningLibraryAndObjects repopulates the "objects" collection field for each library
// by gathering all abstract elements that declare it as their OwningLibrary.
// It also assigns a default library to any abstract elements that are orphaned.
func (stager *Stager) enforceOwningLibraryAndObjects() {
	stage := stager.stage
	// Clear the objects slice of all libraries
	for _, library := range GetGongstrucsSorted[*Library](stage) {
		library.objects = nil
	}

	// redeem objects without owning libraries
	for _, instance := range stage.GetInstances() {
		if abstractObject, ok := instance.(interface {
			AbstractType
			LibraryOwnedType
		}); ok {
			if abstractObject.GetOwningLibrary() == nil && abstractObject != any(stager.getRootLibrary()) {
				abstractObject.SetOwningLibrary(stager.getRootLibrary())
				abstractObject.GetOwningLibrary().objects = append(abstractObject.GetOwningLibrary().objects, abstractObject)
			}
		}
	}
}
