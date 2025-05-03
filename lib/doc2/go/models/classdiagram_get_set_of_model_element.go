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

		gongStructName := IdentifierToGongObjectName(gongStructShape.Identifier)
		gongStruct, ok := gongStructSet[gongStructName]

		if !ok {
			log.Panicln("Diagram", classdiagram.GetName(), "has a shape named", gongStructName, "but no gongstruct exists")
			continue
		}
		map_ModelElement_Shape[gongStruct] = gongStructShape

		for _, fieldShape := range gongStructShape.AttributeShapes {
			fieldShapeName := IdentifierToFieldName(fieldShape.Identifier)

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
			linkShapeName := IdentifierToFieldName(linkShape.Identifier)

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

		gongEnumName := IdentifierToGongObjectName(gongEnumShape.Identifier)
		gongEnum, ok := gongEnumSet[gongEnumName]
		if !ok {
			log.Fatalln("unkown element", gongEnumShape.Identifier)
		}

		map_ModelElement_Shape[gongEnum] = gongEnumShape

		for _, valueShape := range gongEnumShape.GongEnumValueShapes {
			valueShapeName := IdentifierToFieldName(valueShape.Identifier)

			for _, value := range gongEnum.GongEnumValues {
				if value.GetName() == valueShapeName {
					map_ModelElement_Shape[value] = valueShape
				}
			}
		}
	}

	gongNoteSet := *gong.GetGongstructInstancesMap[gong.GongNote](gongStage)
	for _, gongNoteShape := range classdiagram.NoteShapes {

		gongNoteName := IdentifierToGongObjectName(gongNoteShape.Identifier)
		gongNote, ok := gongNoteSet[gongNoteName]
		if !ok {
			log.Fatalln("unkown element", gongNoteShape.Identifier)
		}

		map_ModelElement_Shape[gongNote] = gongNoteShape

		for _, noteLinkShape := range gongNoteShape.NoteShapeLinks {

			switch noteLinkShape.Type {
			case NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE:
				nodeLinkShapeTarget := IdentifierToGongObjectName(noteLinkShape.Identifier)
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
