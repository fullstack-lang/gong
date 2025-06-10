package models

import (
	"log"

	gong "github.com/fullstack-lang/gong/go/models"
)

type ModelElement interface {
}

type Shape interface {
}

func (stager *Stager) compute_map_modelElement_shape(
	classdiagram *Classdiagram,
	gongStage *gong.Stage) (
	map_ModelElement_Shape map[ModelElement]Shape) {
	map_ModelElement_Shape = make(map[ModelElement]Shape)

	gongStructSet := *gong.GetGongstructInstancesMap[gong.GongStruct](stager.gongStage)

	for _, gongStructShape := range classdiagram.GongStructShapes {

		gongStructName := IdentifierMetaToGongStructName(gongStructShape.IdentifierMeta)
		gongStruct, ok := gongStructSet[gongStructName]

		if !ok {
			log.Println("Diagram", classdiagram.GetName(), "has a shape named", gongStructName, "but no gongstruct exists")
			continue
		}
		map_ModelElement_Shape[gongStruct] = gongStructShape

		for _, fieldShape := range gongStructShape.AttributeShapes {
			fieldShapeName := IdentifierMetaToFieldName(fieldShape.IdentifierMeta)

			var fieldFound bool
			for _, field := range gongStruct.Fields {

				if field.GetName() == fieldShapeName {
					map_ModelElement_Shape[field] = fieldShape
					fieldFound = true
				}
			}
			if !fieldFound {
				log.Panicln("Diagram", classdiagram.GetName(), "has a shape named", fieldShapeName, "but no gongstruct exists")
			}
		}

		for _, linkShape := range gongStructShape.LinkShapes {
			linkShapeName := IdentifierMetaToFieldName(linkShape.IdentifierMeta)

			for _, link := range gongStruct.SliceOfPointerToGongStructFields {
				if link.GetName() == linkShapeName {
					map_ModelElement_Shape[link] = linkShape
				}
			}

			for _, link := range gongStruct.PointerToGongStructFields {
				if link.GetName() == linkShapeName {
					map_ModelElement_Shape[link] = linkShape
				}
			}
		}
	}

	gongEnumSet := *gong.GetGongstructInstancesMap[gong.GongEnum](gongStage)
	for _, gongEnumShape := range classdiagram.GongEnumShapes {

		gongEnumName := GongEnumIdentifierMetaToGongEnumName(gongEnumShape.IdentifierMeta)
		gongEnum, ok := gongEnumSet[gongEnumName]
		if !ok {
			log.Fatalln("unkown element", gongEnumName)
		}

		map_ModelElement_Shape[gongEnum] = gongEnumShape

		for _, gongEnumValueShape := range gongEnumShape.GongEnumValueShapes {
			valueShapeName := GongEnumValueShapeIdentifierMetaToValueName(gongEnumValueShape.IdentifierMeta)

			for _, value := range gongEnum.GongEnumValues {
				if value.GetName() == valueShapeName {
					map_ModelElement_Shape[value] = gongEnumValueShape
				}
			}
		}
	}

	gongNoteSet := *gong.GetGongstructInstancesMap[gong.GongNote](gongStage)
	for _, gongNoteShape := range classdiagram.GongNoteShapes {

		gongNoteName := IdentifierToGongStructName(gongNoteShape.Identifier)
		gongNote, ok := gongNoteSet[gongNoteName]
		if !ok {
			log.Fatalln("unkown element", gongNoteShape.Identifier)
		}

		map_ModelElement_Shape[gongNote] = gongNoteShape

		for _, noteLinkShape := range gongNoteShape.GongNoteLinkShapes {

			switch noteLinkShape.Type {
			case NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE:
				nodeLinkShapeTarget := IdentifierToGongStructName(noteLinkShape.Identifier)
				for _, link := range gongNote.Links {
					if link.GetName() == nodeLinkShapeTarget {
						map_ModelElement_Shape[link] = noteLinkShape
					}
				}
			case NOTE_SHAPE_LINK_TO_GONG_FIELD:
				receiver, fieldName := IdentifierToReceiverAndFieldName(noteLinkShape.Identifier)
				for _, link := range gongNote.Links {
					if link.GetName() == fieldName && link.Recv == receiver {
						map_ModelElement_Shape[link] = noteLinkShape
					}
				}

			}
		}
	}
	return
}
