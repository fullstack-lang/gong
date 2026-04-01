package models

func (stager *Stager) initMap_Objects_Relations() {

	relations := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPEC_RELATIONS
	if relations == nil {
		return
	}

	for _, specRelation := range relations.SPEC_RELATION {

		{
			specObject, ok := stager.Map_id_SPEC_OBJECT[specRelation.SOURCE.SPEC_OBJECT_REF]
			if ok {
				stager.Map_SPEC_OBJECT_relations_sources[specObject] =
					append(stager.Map_SPEC_OBJECT_relations_sources[specObject],
						specRelation,
					)

			}
		}
		{
			specObject, ok := stager.Map_id_SPEC_OBJECT[specRelation.TARGET.SPEC_OBJECT_REF]
			if ok {
				stager.Map_SPEC_OBJECT_relations_targets[specObject] =
					append(stager.Map_SPEC_OBJECT_relations_targets[specObject],
						specRelation,
					)
			}
		}
	}
}
