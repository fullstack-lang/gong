// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	ClassdiagramAPIs []*ClassdiagramAPI

	DiagramPackageAPIs []*DiagramPackageAPI

	FieldAPIs []*FieldAPI

	GongEnumShapeAPIs []*GongEnumShapeAPI

	GongEnumValueEntryAPIs []*GongEnumValueEntryAPI

	GongStructShapeAPIs []*GongStructShapeAPI

	LinkAPIs []*LinkAPI

	NoteShapeAPIs []*NoteShapeAPI

	NoteShapeLinkAPIs []*NoteShapeLinkAPI

	PositionAPIs []*PositionAPI

	UmlStateAPIs []*UmlStateAPI

	UmlscAPIs []*UmlscAPI

	VerticeAPIs []*VerticeAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {
	// insertion point for slices copies
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

	for _, fieldDB := range backRepo.BackRepoField.Map_FieldDBID_FieldDB {

		var fieldAPI FieldAPI
		fieldAPI.ID = fieldDB.ID
		fieldAPI.FieldPointersEncoding = fieldDB.FieldPointersEncoding
		fieldDB.CopyBasicFieldsToField_WOP(&fieldAPI.Field_WOP)

		backRepoData.FieldAPIs = append(backRepoData.FieldAPIs, &fieldAPI)
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

	for _, linkDB := range backRepo.BackRepoLink.Map_LinkDBID_LinkDB {

		var linkAPI LinkAPI
		linkAPI.ID = linkDB.ID
		linkAPI.LinkPointersEncoding = linkDB.LinkPointersEncoding
		linkDB.CopyBasicFieldsToLink_WOP(&linkAPI.Link_WOP)

		backRepoData.LinkAPIs = append(backRepoData.LinkAPIs, &linkAPI)
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

	for _, positionDB := range backRepo.BackRepoPosition.Map_PositionDBID_PositionDB {

		var positionAPI PositionAPI
		positionAPI.ID = positionDB.ID
		positionAPI.PositionPointersEncoding = positionDB.PositionPointersEncoding
		positionDB.CopyBasicFieldsToPosition_WOP(&positionAPI.Position_WOP)

		backRepoData.PositionAPIs = append(backRepoData.PositionAPIs, &positionAPI)
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

	for _, verticeDB := range backRepo.BackRepoVertice.Map_VerticeDBID_VerticeDB {

		var verticeAPI VerticeAPI
		verticeAPI.ID = verticeDB.ID
		verticeAPI.VerticePointersEncoding = verticeDB.VerticePointersEncoding
		verticeDB.CopyBasicFieldsToVertice_WOP(&verticeAPI.Vertice_WOP)

		backRepoData.VerticeAPIs = append(backRepoData.VerticeAPIs, &verticeAPI)
	}

}
