package models

import (
	"sort"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
func SetupMapDocLinkRenaming() {

	// set up Map_DocLink_Renaming
	//  TO BE REMOVED
	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	Stage.Map_DocLink_Renaming = make(map[string]GONG__Identifier)

	gongstructOrdered := []*gong_models.GongStruct{}
	for gongstruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct]() {
		gongstructOrdered = append(gongstructOrdered, gongstruct)
	}
	sort.Slice(gongstructOrdered[:], func(i, j int) bool {
		return gongstructOrdered[i].Name < gongstructOrdered[j].Name
	})
	for _, gongStruct := range gongstructOrdered {

		ident := GongStructNameToIdentifier(gongStruct.Name)
		var identifier GONG__Identifier
		identifier.Ident = ident
		identifier.Type = GONG__STRUCT_INSTANCE

		Stage.Map_DocLink_Renaming[ident] = identifier

		for _, field := range gongStruct.Fields {
			ident := GongstructAndFieldnameToFieldIdentifier(
				gongStruct.Name, field.GetName())

			var identifier GONG__Identifier
			identifier.Ident = ident
			identifier.Type = GONG__FIELD_VALUE
			Stage.Map_DocLink_Renaming[ident] = identifier
		}
	}
	for gongEnum := range *gong_models.GetGongstructInstancesSet[gong_models.GongEnum]() {
		ident := GongStructNameToIdentifier(gongEnum.Name)

		var identifier GONG__Identifier
		identifier.Ident = ident
		switch gongEnum.Type {
		case gong_models.Int:
			identifier.Type = GONG__ENUM_CAST_INT
		case gong_models.String:
			identifier.Type = GONG__ENUM_CAST_STRING
		}

		Stage.Map_DocLink_Renaming[ident] = identifier

		for _, value := range gongEnum.GongEnumValues {
			ident := GongStructNameToIdentifier(value.Name)

			var identifier GONG__Identifier
			identifier.Ident = ident
			identifier.Type = GONG__IDENTIFIER_CONST
			Stage.Map_DocLink_Renaming[ident] = identifier
		}

		// to do after fix of https://github.com/fullstack-lang/gongdoc/issues/100
		// Stage.Map_DocLink_Renaming[ident] = ident
	}

	for gongNote := range *gong_models.GetGongstructInstancesSet[gong_models.GongNote]() {
		ident := GongStructNameToIdentifier(gongNote.Name)

		var identifier GONG__Identifier
		identifier.Ident = ident
		identifier.Type = GONG__IDENTIFIER_CONST
		Stage.Map_DocLink_Renaming[ident] = identifier
	}

	// end of TO BE REMOVED

}
