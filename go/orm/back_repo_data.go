// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	GongBasicFieldAPIs []*GongBasicFieldAPI

	GongEnumAPIs []*GongEnumAPI

	GongEnumValueAPIs []*GongEnumValueAPI

	GongLinkAPIs []*GongLinkAPI

	GongNoteAPIs []*GongNoteAPI

	GongStructAPIs []*GongStructAPI

	GongTimeFieldAPIs []*GongTimeFieldAPI

	MetaAPIs []*MetaAPI

	MetaReferenceAPIs []*MetaReferenceAPI

	ModelPkgAPIs []*ModelPkgAPI

	PointerToGongStructFieldAPIs []*PointerToGongStructFieldAPI

	SliceOfPointerToGongStructFieldAPIs []*SliceOfPointerToGongStructFieldAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {
	// insertion point for slices copies
	for _, gongbasicfieldDB := range backRepo.BackRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldDB {

		var gongbasicfieldAPI GongBasicFieldAPI
		gongbasicfieldAPI.ID = gongbasicfieldDB.ID
		gongbasicfieldAPI.GongBasicFieldPointersEncoding = gongbasicfieldDB.GongBasicFieldPointersEncoding
		gongbasicfieldDB.CopyBasicFieldsToGongBasicField_WOP(&gongbasicfieldAPI.GongBasicField_WOP)

		backRepoData.GongBasicFieldAPIs = append(backRepoData.GongBasicFieldAPIs, &gongbasicfieldAPI)
	}

	for _, gongenumDB := range backRepo.BackRepoGongEnum.Map_GongEnumDBID_GongEnumDB {

		var gongenumAPI GongEnumAPI
		gongenumAPI.ID = gongenumDB.ID
		gongenumAPI.GongEnumPointersEncoding = gongenumDB.GongEnumPointersEncoding
		gongenumDB.CopyBasicFieldsToGongEnum_WOP(&gongenumAPI.GongEnum_WOP)

		backRepoData.GongEnumAPIs = append(backRepoData.GongEnumAPIs, &gongenumAPI)
	}

	for _, gongenumvalueDB := range backRepo.BackRepoGongEnumValue.Map_GongEnumValueDBID_GongEnumValueDB {

		var gongenumvalueAPI GongEnumValueAPI
		gongenumvalueAPI.ID = gongenumvalueDB.ID
		gongenumvalueAPI.GongEnumValuePointersEncoding = gongenumvalueDB.GongEnumValuePointersEncoding
		gongenumvalueDB.CopyBasicFieldsToGongEnumValue_WOP(&gongenumvalueAPI.GongEnumValue_WOP)

		backRepoData.GongEnumValueAPIs = append(backRepoData.GongEnumValueAPIs, &gongenumvalueAPI)
	}

	for _, gonglinkDB := range backRepo.BackRepoGongLink.Map_GongLinkDBID_GongLinkDB {

		var gonglinkAPI GongLinkAPI
		gonglinkAPI.ID = gonglinkDB.ID
		gonglinkAPI.GongLinkPointersEncoding = gonglinkDB.GongLinkPointersEncoding
		gonglinkDB.CopyBasicFieldsToGongLink_WOP(&gonglinkAPI.GongLink_WOP)

		backRepoData.GongLinkAPIs = append(backRepoData.GongLinkAPIs, &gonglinkAPI)
	}

	for _, gongnoteDB := range backRepo.BackRepoGongNote.Map_GongNoteDBID_GongNoteDB {

		var gongnoteAPI GongNoteAPI
		gongnoteAPI.ID = gongnoteDB.ID
		gongnoteAPI.GongNotePointersEncoding = gongnoteDB.GongNotePointersEncoding
		gongnoteDB.CopyBasicFieldsToGongNote_WOP(&gongnoteAPI.GongNote_WOP)

		backRepoData.GongNoteAPIs = append(backRepoData.GongNoteAPIs, &gongnoteAPI)
	}

	for _, gongstructDB := range backRepo.BackRepoGongStruct.Map_GongStructDBID_GongStructDB {

		var gongstructAPI GongStructAPI
		gongstructAPI.ID = gongstructDB.ID
		gongstructAPI.GongStructPointersEncoding = gongstructDB.GongStructPointersEncoding
		gongstructDB.CopyBasicFieldsToGongStruct_WOP(&gongstructAPI.GongStruct_WOP)

		backRepoData.GongStructAPIs = append(backRepoData.GongStructAPIs, &gongstructAPI)
	}

	for _, gongtimefieldDB := range backRepo.BackRepoGongTimeField.Map_GongTimeFieldDBID_GongTimeFieldDB {

		var gongtimefieldAPI GongTimeFieldAPI
		gongtimefieldAPI.ID = gongtimefieldDB.ID
		gongtimefieldAPI.GongTimeFieldPointersEncoding = gongtimefieldDB.GongTimeFieldPointersEncoding
		gongtimefieldDB.CopyBasicFieldsToGongTimeField_WOP(&gongtimefieldAPI.GongTimeField_WOP)

		backRepoData.GongTimeFieldAPIs = append(backRepoData.GongTimeFieldAPIs, &gongtimefieldAPI)
	}

	for _, metaDB := range backRepo.BackRepoMeta.Map_MetaDBID_MetaDB {

		var metaAPI MetaAPI
		metaAPI.ID = metaDB.ID
		metaAPI.MetaPointersEncoding = metaDB.MetaPointersEncoding
		metaDB.CopyBasicFieldsToMeta_WOP(&metaAPI.Meta_WOP)

		backRepoData.MetaAPIs = append(backRepoData.MetaAPIs, &metaAPI)
	}

	for _, metareferenceDB := range backRepo.BackRepoMetaReference.Map_MetaReferenceDBID_MetaReferenceDB {

		var metareferenceAPI MetaReferenceAPI
		metareferenceAPI.ID = metareferenceDB.ID
		metareferenceAPI.MetaReferencePointersEncoding = metareferenceDB.MetaReferencePointersEncoding
		metareferenceDB.CopyBasicFieldsToMetaReference_WOP(&metareferenceAPI.MetaReference_WOP)

		backRepoData.MetaReferenceAPIs = append(backRepoData.MetaReferenceAPIs, &metareferenceAPI)
	}

	for _, modelpkgDB := range backRepo.BackRepoModelPkg.Map_ModelPkgDBID_ModelPkgDB {

		var modelpkgAPI ModelPkgAPI
		modelpkgAPI.ID = modelpkgDB.ID
		modelpkgAPI.ModelPkgPointersEncoding = modelpkgDB.ModelPkgPointersEncoding
		modelpkgDB.CopyBasicFieldsToModelPkg_WOP(&modelpkgAPI.ModelPkg_WOP)

		backRepoData.ModelPkgAPIs = append(backRepoData.ModelPkgAPIs, &modelpkgAPI)
	}

	for _, pointertogongstructfieldDB := range backRepo.BackRepoPointerToGongStructField.Map_PointerToGongStructFieldDBID_PointerToGongStructFieldDB {

		var pointertogongstructfieldAPI PointerToGongStructFieldAPI
		pointertogongstructfieldAPI.ID = pointertogongstructfieldDB.ID
		pointertogongstructfieldAPI.PointerToGongStructFieldPointersEncoding = pointertogongstructfieldDB.PointerToGongStructFieldPointersEncoding
		pointertogongstructfieldDB.CopyBasicFieldsToPointerToGongStructField_WOP(&pointertogongstructfieldAPI.PointerToGongStructField_WOP)

		backRepoData.PointerToGongStructFieldAPIs = append(backRepoData.PointerToGongStructFieldAPIs, &pointertogongstructfieldAPI)
	}

	for _, sliceofpointertogongstructfieldDB := range backRepo.BackRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldDB {

		var sliceofpointertogongstructfieldAPI SliceOfPointerToGongStructFieldAPI
		sliceofpointertogongstructfieldAPI.ID = sliceofpointertogongstructfieldDB.ID
		sliceofpointertogongstructfieldAPI.SliceOfPointerToGongStructFieldPointersEncoding = sliceofpointertogongstructfieldDB.SliceOfPointerToGongStructFieldPointersEncoding
		sliceofpointertogongstructfieldDB.CopyBasicFieldsToSliceOfPointerToGongStructField_WOP(&sliceofpointertogongstructfieldAPI.SliceOfPointerToGongStructField_WOP)

		backRepoData.SliceOfPointerToGongStructFieldAPIs = append(backRepoData.SliceOfPointerToGongStructFieldAPIs, &sliceofpointertogongstructfieldAPI)
	}

}
