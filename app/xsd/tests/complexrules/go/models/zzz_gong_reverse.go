// generated code - do not edit
package models

// insertion point
func (inst *Catalog) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DerivedItemType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Catalog":
		switch reverseField.Fieldname {
		case "Item":
			if _catalog, ok := stage.Catalog_Item_reverseMap[inst]; ok {
				res = _catalog.Name
			}
		}
	}
	return
}
