package models

import gong_models "github.com/fullstack-lang/gong/go/models"

// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
func SetupMapDocLinkRenamingNew(gongdocStage *StageStruct, diagramPackage *DiagramPackage) {

	// set up Map_DocLink_Renaming
	//  TO BE REMOVED
	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	map_DocLink_Renaming := make(map[string]GONG__Identifier)

	for field := range *GetGongstructInstancesSet[Field](gongdocStage) {

		var identifier GONG__Identifier
		identifier.Ident = field.Identifier
		identifier.Type = GONG__FIELD_VALUE

		map_DocLink_Renaming[field.Identifier] = identifier
	}

	for gongEnumShape := range *GetGongstructInstancesSet[GongEnumShape](gongdocStage) {

		var identifier GONG__Identifier
		identifier.Ident = gongEnumShape.Identifier

		shapeEnumName := IdentifierToGongObjectName(gongEnumShape.Identifier)
		for _, gongEnum := range diagramPackage.ModelPkg.GongEnums {
			if gongEnum.Name == shapeEnumName {
				switch gongEnum.Type {
				case gong_models.Int:
					identifier.Type = GONG__ENUM_CAST_INT
				case gong_models.String:
					identifier.Type = GONG__ENUM_CAST_STRING
				}
			}
		}
		map_DocLink_Renaming[gongEnumShape.Identifier] = identifier
	}

	for gongStructShape := range *GetGongstructInstancesSet[GongStructShape](gongdocStage) {

		var identifier GONG__Identifier
		identifier.Ident = gongStructShape.Identifier
		identifier.Type = GONG__STRUCT_INSTANCE
		map_DocLink_Renaming[gongStructShape.Identifier] = identifier
	}

	for link := range *GetGongstructInstancesSet[Link](gongdocStage) {

		var identifier GONG__Identifier
		identifier.Ident = link.Identifier
		identifier.Type = GONG__FIELD_VALUE
		map_DocLink_Renaming[link.Identifier] = identifier
	}

	for noteShape := range *GetGongstructInstancesSet[NoteShape](gongdocStage) {

		var identifier GONG__Identifier
		identifier.Ident = noteShape.Identifier
		identifier.Type = GONG__IDENTIFIER_CONST
		map_DocLink_Renaming[noteShape.Identifier] = identifier
	}

	Stage.Map_DocLink_Renaming = map_DocLink_Renaming
}
