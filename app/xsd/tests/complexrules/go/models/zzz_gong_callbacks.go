// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront(instance GongstructIF) {
	if instance != nil {
		instance.GongAfterCreateFromFront(stage)
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront(old, new GongstructIF) {
	if old != nil {
		old.GongOnAfterUpdateFromFront(stage, new)
	}
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront(staged, front GongstructIF) {
	if staged != nil {
		staged.GongAfterDeleteFromFront(stage, front)
	}
}

// insertion point
func (catalog *Catalog) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCatalogCreateCallback != nil {
		stage.OnAfterCatalogCreateCallback.OnAfterCreate(stage, catalog)
	}
}

func (catalog *Catalog) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCatalogUpdateCallback != nil {
		var frontCatalog *Catalog
		if front != nil {
			frontCatalog, _ = front.(*Catalog)
		}
		stage.OnAfterCatalogUpdateCallback.OnAfterUpdate(stage, catalog, frontCatalog)
	}
}

func (catalog *Catalog) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCatalogDeleteCallback != nil {
		var frontCatalog *Catalog
		if front != nil {
			frontCatalog, _ = front.(*Catalog)
		}
		stage.OnAfterCatalogDeleteCallback.OnAfterDelete(stage, catalog, frontCatalog)
	}
}

func (deriveditemtype *DerivedItemType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDerivedItemTypeCreateCallback != nil {
		stage.OnAfterDerivedItemTypeCreateCallback.OnAfterCreate(stage, deriveditemtype)
	}
}

func (deriveditemtype *DerivedItemType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDerivedItemTypeUpdateCallback != nil {
		var frontDerivedItemType *DerivedItemType
		if front != nil {
			frontDerivedItemType, _ = front.(*DerivedItemType)
		}
		stage.OnAfterDerivedItemTypeUpdateCallback.OnAfterUpdate(stage, deriveditemtype, frontDerivedItemType)
	}
}

func (deriveditemtype *DerivedItemType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDerivedItemTypeDeleteCallback != nil {
		var frontDerivedItemType *DerivedItemType
		if front != nil {
			frontDerivedItemType, _ = front.(*DerivedItemType)
		}
		stage.OnAfterDerivedItemTypeDeleteCallback.OnAfterDelete(stage, deriveditemtype, frontDerivedItemType)
	}
}

