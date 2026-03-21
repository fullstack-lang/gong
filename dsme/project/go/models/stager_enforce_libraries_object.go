package models

// enforceLibrariesObject repopulates the "objects" collection field for each library
// by gathering all abstract elements that declare it as their OwningLibrary.
// It also assigns a default library to any abstract elements that are orphaned.
func (stager *Stager) enforceLibrariesObject() {
	stage := stager.stage

	// fetch the first library as the default one
	var defaultLibrary *Library
	for library_ := range *GetGongstructInstancesSetFromPointerType[*Library](stage) {
		defaultLibrary = library_
		break
	}
	if defaultLibrary == nil {
		return
	}

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
			if abstractObject.GetOwnlingLibrary() == nil {
				abstractObject.SetOwningLibrary(defaultLibrary)
			}
			abstractObject.GetOwnlingLibrary().objects = append(abstractObject.GetOwnlingLibrary().objects, abstractObject)
		}
	}
}
