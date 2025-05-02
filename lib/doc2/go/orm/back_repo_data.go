// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	AttributeShapeAPIs []*AttributeShapeAPI

	ClassdiagramAPIs []*ClassdiagramAPI

	DiagramPackageAPIs []*DiagramPackageAPI

	GongEnumShapeAPIs []*GongEnumShapeAPI

	GongEnumValueEntryAPIs []*GongEnumValueEntryAPI

	GongStructShapeAPIs []*GongStructShapeAPI

	LinkShapeAPIs []*LinkShapeAPI

	NoteShapeAPIs []*NoteShapeAPI

	NoteShapeLinkAPIs []*NoteShapeLinkAPI

	UmlStateAPIs []*UmlStateAPI

	UmlscAPIs []*UmlscAPI

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

	for _, gongenumvalueentryDB := range backRepo.BackRepoGongEnumValueEntry.Map_GongEnumValueEntryDBID_GongEnumValueEntryDB {

		var gongenumvalueentryAPI GongEnumValueEntryAPI
		gongenumvalueentryAPI.ID = gongenumvalueentryDB.ID
		gongenumvalueentryAPI.GongEnumValueEntryPointersEncoding = gongenumvalueentryDB.GongEnumValueEntryPointersEncoding
		gongenumvalueentryDB.CopyBasicFieldsToGongEnumValueEntry_WOP(&gongenumvalueentryAPI.GongEnumValueEntry_WOP)

		backRepoData.GongEnumValueEntryAPIs = append(backRepoData.GongEnumValueEntryAPIs, &gongenumvalueentryAPI)
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

	for _, noteshapeDB := range backRepo.BackRepoNoteShape.Map_NoteShapeDBID_NoteShapeDB {

		var noteshapeAPI NoteShapeAPI
		noteshapeAPI.ID = noteshapeDB.ID
		noteshapeAPI.NoteShapePointersEncoding = noteshapeDB.NoteShapePointersEncoding
		noteshapeDB.CopyBasicFieldsToNoteShape_WOP(&noteshapeAPI.NoteShape_WOP)

		backRepoData.NoteShapeAPIs = append(backRepoData.NoteShapeAPIs, &noteshapeAPI)
	}

	for _, noteshapelinkDB := range backRepo.BackRepoNoteShapeLink.Map_NoteShapeLinkDBID_NoteShapeLinkDB {

		var noteshapelinkAPI NoteShapeLinkAPI
		noteshapelinkAPI.ID = noteshapelinkDB.ID
		noteshapelinkAPI.NoteShapeLinkPointersEncoding = noteshapelinkDB.NoteShapeLinkPointersEncoding
		noteshapelinkDB.CopyBasicFieldsToNoteShapeLink_WOP(&noteshapelinkAPI.NoteShapeLink_WOP)

		backRepoData.NoteShapeLinkAPIs = append(backRepoData.NoteShapeLinkAPIs, &noteshapelinkAPI)
	}

	for _, umlstateDB := range backRepo.BackRepoUmlState.Map_UmlStateDBID_UmlStateDB {

		var umlstateAPI UmlStateAPI
		umlstateAPI.ID = umlstateDB.ID
		umlstateAPI.UmlStatePointersEncoding = umlstateDB.UmlStatePointersEncoding
		umlstateDB.CopyBasicFieldsToUmlState_WOP(&umlstateAPI.UmlState_WOP)

		backRepoData.UmlStateAPIs = append(backRepoData.UmlStateAPIs, &umlstateAPI)
	}

	for _, umlscDB := range backRepo.BackRepoUmlsc.Map_UmlscDBID_UmlscDB {

		var umlscAPI UmlscAPI
		umlscAPI.ID = umlscDB.ID
		umlscAPI.UmlscPointersEncoding = umlscDB.UmlscPointersEncoding
		umlscDB.CopyBasicFieldsToUmlsc_WOP(&umlscAPI.Umlsc_WOP)

		backRepoData.UmlscAPIs = append(backRepoData.UmlscAPIs, &umlscAPI)
	}

}
