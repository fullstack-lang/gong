// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	AttributeShapeAPIs []*AttributeShapeAPI

	ClassdiagramAPIs []*ClassdiagramAPI

	DiagramPackageAPIs []*DiagramPackageAPI

	GongEnumShapeAPIs []*GongEnumShapeAPI

	GongEnumValueShapeAPIs []*GongEnumValueShapeAPI

	GongNoteLinkShapeAPIs []*GongNoteLinkShapeAPI

	GongNoteShapeAPIs []*GongNoteShapeAPI

	GongStructShapeAPIs []*GongStructShapeAPI

	LinkShapeAPIs []*LinkShapeAPI

	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index int
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, attributeshapeDB := range backRepo.BackRepoAttributeShape.Map_AttributeShapeDBID_AttributeShapeDB {

		var attributeshapeAPI AttributeShapeAPI
		attributeshapeAPI.ID = attributeshapeDB.ID
		attributeshapeAPI.AttributeShapePointersEncoding = attributeshapeDB.AttributeShapePointersEncoding
		attributeshapeDB.CopyBasicFieldsToAttributeShape_WOP(&attributeshapeAPI.AttributeShape_WOP)

		backRepoData.AttributeShapeAPIs = append(backRepoData.AttributeShapeAPIs, &attributeshapeAPI)
	}

	for _, classdiagramDB := range backRepo.BackRepoClassdiagram.Map_ClassdiagramDBID_ClassdiagramDB {

		var classdiagramAPI ClassdiagramAPI
		classdiagramAPI.ID = classdiagramDB.ID
		classdiagramAPI.ClassdiagramPointersEncoding = classdiagramDB.ClassdiagramPointersEncoding
		classdiagramDB.CopyBasicFieldsToClassdiagram_WOP(&classdiagramAPI.Classdiagram_WOP)

		backRepoData.ClassdiagramAPIs = append(backRepoData.ClassdiagramAPIs, &classdiagramAPI)
	}

	for _, diagrampackageDB := range backRepo.BackRepoDiagramPackage.Map_DiagramPackageDBID_DiagramPackageDB {

		var diagrampackageAPI DiagramPackageAPI
		diagrampackageAPI.ID = diagrampackageDB.ID
		diagrampackageAPI.DiagramPackagePointersEncoding = diagrampackageDB.DiagramPackagePointersEncoding
		diagrampackageDB.CopyBasicFieldsToDiagramPackage_WOP(&diagrampackageAPI.DiagramPackage_WOP)

		backRepoData.DiagramPackageAPIs = append(backRepoData.DiagramPackageAPIs, &diagrampackageAPI)
	}

	for _, gongenumshapeDB := range backRepo.BackRepoGongEnumShape.Map_GongEnumShapeDBID_GongEnumShapeDB {

		var gongenumshapeAPI GongEnumShapeAPI
		gongenumshapeAPI.ID = gongenumshapeDB.ID
		gongenumshapeAPI.GongEnumShapePointersEncoding = gongenumshapeDB.GongEnumShapePointersEncoding
		gongenumshapeDB.CopyBasicFieldsToGongEnumShape_WOP(&gongenumshapeAPI.GongEnumShape_WOP)

		backRepoData.GongEnumShapeAPIs = append(backRepoData.GongEnumShapeAPIs, &gongenumshapeAPI)
	}

	for _, gongenumvalueshapeDB := range backRepo.BackRepoGongEnumValueShape.Map_GongEnumValueShapeDBID_GongEnumValueShapeDB {

		var gongenumvalueshapeAPI GongEnumValueShapeAPI
		gongenumvalueshapeAPI.ID = gongenumvalueshapeDB.ID
		gongenumvalueshapeAPI.GongEnumValueShapePointersEncoding = gongenumvalueshapeDB.GongEnumValueShapePointersEncoding
		gongenumvalueshapeDB.CopyBasicFieldsToGongEnumValueShape_WOP(&gongenumvalueshapeAPI.GongEnumValueShape_WOP)

		backRepoData.GongEnumValueShapeAPIs = append(backRepoData.GongEnumValueShapeAPIs, &gongenumvalueshapeAPI)
	}

	for _, gongnotelinkshapeDB := range backRepo.BackRepoGongNoteLinkShape.Map_GongNoteLinkShapeDBID_GongNoteLinkShapeDB {

		var gongnotelinkshapeAPI GongNoteLinkShapeAPI
		gongnotelinkshapeAPI.ID = gongnotelinkshapeDB.ID
		gongnotelinkshapeAPI.GongNoteLinkShapePointersEncoding = gongnotelinkshapeDB.GongNoteLinkShapePointersEncoding
		gongnotelinkshapeDB.CopyBasicFieldsToGongNoteLinkShape_WOP(&gongnotelinkshapeAPI.GongNoteLinkShape_WOP)

		backRepoData.GongNoteLinkShapeAPIs = append(backRepoData.GongNoteLinkShapeAPIs, &gongnotelinkshapeAPI)
	}

	for _, gongnoteshapeDB := range backRepo.BackRepoGongNoteShape.Map_GongNoteShapeDBID_GongNoteShapeDB {

		var gongnoteshapeAPI GongNoteShapeAPI
		gongnoteshapeAPI.ID = gongnoteshapeDB.ID
		gongnoteshapeAPI.GongNoteShapePointersEncoding = gongnoteshapeDB.GongNoteShapePointersEncoding
		gongnoteshapeDB.CopyBasicFieldsToGongNoteShape_WOP(&gongnoteshapeAPI.GongNoteShape_WOP)

		backRepoData.GongNoteShapeAPIs = append(backRepoData.GongNoteShapeAPIs, &gongnoteshapeAPI)
	}

	for _, gongstructshapeDB := range backRepo.BackRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapeDB {

		var gongstructshapeAPI GongStructShapeAPI
		gongstructshapeAPI.ID = gongstructshapeDB.ID
		gongstructshapeAPI.GongStructShapePointersEncoding = gongstructshapeDB.GongStructShapePointersEncoding
		gongstructshapeDB.CopyBasicFieldsToGongStructShape_WOP(&gongstructshapeAPI.GongStructShape_WOP)

		backRepoData.GongStructShapeAPIs = append(backRepoData.GongStructShapeAPIs, &gongstructshapeAPI)
	}

	for _, linkshapeDB := range backRepo.BackRepoLinkShape.Map_LinkShapeDBID_LinkShapeDB {

		var linkshapeAPI LinkShapeAPI
		linkshapeAPI.ID = linkshapeDB.ID
		linkshapeAPI.LinkShapePointersEncoding = linkshapeDB.LinkShapePointersEncoding
		linkshapeDB.CopyBasicFieldsToLinkShape_WOP(&linkshapeAPI.LinkShape_WOP)

		backRepoData.LinkShapeAPIs = append(backRepoData.LinkShapeAPIs, &linkshapeAPI)
	}

}
