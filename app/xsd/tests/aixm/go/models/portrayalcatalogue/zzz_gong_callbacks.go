// generated code - do not edit
package portrayalcatalogue

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
func (md_portrayalcataloguereference_propertytype *MD_PortrayalCatalogueReference_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_PortrayalCatalogueReference_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_PortrayalCatalogueReference_PropertyTypeCreateCallback.OnAfterCreate(stage, md_portrayalcataloguereference_propertytype)
	}
}

func (md_portrayalcataloguereference_propertytype *MD_PortrayalCatalogueReference_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_PortrayalCatalogueReference_PropertyTypeUpdateCallback != nil {
		var frontMD_PortrayalCatalogueReference_PropertyType *MD_PortrayalCatalogueReference_PropertyType
		if front != nil {
			frontMD_PortrayalCatalogueReference_PropertyType, _ = front.(*MD_PortrayalCatalogueReference_PropertyType)
		}
		stage.OnAfterMD_PortrayalCatalogueReference_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_portrayalcataloguereference_propertytype, frontMD_PortrayalCatalogueReference_PropertyType)
	}
}

func (md_portrayalcataloguereference_propertytype *MD_PortrayalCatalogueReference_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_PortrayalCatalogueReference_PropertyTypeDeleteCallback != nil {
		var frontMD_PortrayalCatalogueReference_PropertyType *MD_PortrayalCatalogueReference_PropertyType
		if front != nil {
			frontMD_PortrayalCatalogueReference_PropertyType, _ = front.(*MD_PortrayalCatalogueReference_PropertyType)
		}
		stage.OnAfterMD_PortrayalCatalogueReference_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_portrayalcataloguereference_propertytype, frontMD_PortrayalCatalogueReference_PropertyType)
	}
}

func (md_portrayalcataloguereference_type *MD_PortrayalCatalogueReference_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_PortrayalCatalogueReference_TypeCreateCallback != nil {
		stage.OnAfterMD_PortrayalCatalogueReference_TypeCreateCallback.OnAfterCreate(stage, md_portrayalcataloguereference_type)
	}
}

func (md_portrayalcataloguereference_type *MD_PortrayalCatalogueReference_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_PortrayalCatalogueReference_TypeUpdateCallback != nil {
		var frontMD_PortrayalCatalogueReference_Type *MD_PortrayalCatalogueReference_Type
		if front != nil {
			frontMD_PortrayalCatalogueReference_Type, _ = front.(*MD_PortrayalCatalogueReference_Type)
		}
		stage.OnAfterMD_PortrayalCatalogueReference_TypeUpdateCallback.OnAfterUpdate(stage, md_portrayalcataloguereference_type, frontMD_PortrayalCatalogueReference_Type)
	}
}

func (md_portrayalcataloguereference_type *MD_PortrayalCatalogueReference_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_PortrayalCatalogueReference_TypeDeleteCallback != nil {
		var frontMD_PortrayalCatalogueReference_Type *MD_PortrayalCatalogueReference_Type
		if front != nil {
			frontMD_PortrayalCatalogueReference_Type, _ = front.(*MD_PortrayalCatalogueReference_Type)
		}
		stage.OnAfterMD_PortrayalCatalogueReference_TypeDeleteCallback.OnAfterDelete(stage, md_portrayalcataloguereference_type, frontMD_PortrayalCatalogueReference_Type)
	}
}

