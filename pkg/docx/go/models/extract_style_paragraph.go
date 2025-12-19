package models

import "log"

func ExtractStyleText(styleName string, gongdocxStage *Stage) (res []string) {

	res = make([]string, 0)

	map_ParagraphStyle_StyleName := make(map[ParagraphStyle]string)
	map_StyleName_Occurence := make(map[string]int)
	map_StyleName_SliceOfParagraphStyle := make(map[string][]*ParagraphStyle)

	for paragraphStyle := range *GetGongstructInstancesSet[ParagraphStyle](gongdocxStage) {

		if paragraphStyle.ValAttr != "" {
			map_ParagraphStyle_StyleName[*paragraphStyle] = paragraphStyle.ValAttr

			if val, ok := map_StyleName_Occurence[paragraphStyle.ValAttr]; !ok {
				map_StyleName_Occurence[paragraphStyle.ValAttr] = 1
				map_StyleName_SliceOfParagraphStyle[paragraphStyle.ValAttr] = make([]*ParagraphStyle, 1)
				map_StyleName_SliceOfParagraphStyle[paragraphStyle.ValAttr] =
					append(map_StyleName_SliceOfParagraphStyle[paragraphStyle.ValAttr], paragraphStyle)
			} else {
				map_StyleName_Occurence[paragraphStyle.ValAttr] = val + 1
				map_StyleName_SliceOfParagraphStyle[paragraphStyle.ValAttr] =
					append(map_StyleName_SliceOfParagraphStyle[paragraphStyle.ValAttr], paragraphStyle)
			}
		}
	}

	reverseMapParagraph_ParagraphProperties :=
		GetPointerReverseMap[
			Paragraph, ParagraphProperties](
			GetAssociationName[Paragraph]().ParagraphProperties.Name, gongdocxStage)
	_ = reverseMapParagraph_ParagraphProperties

	reverseMapParagraphProperties_ParagraphStyles :=
		GetPointerReverseMap[
			ParagraphProperties, ParagraphStyle](
			GetAssociationName[ParagraphProperties]().ParagraphStyle.Name, gongdocxStage)
	_ = reverseMapParagraphProperties_ParagraphStyles

	log.Println("Used styles are ")
	for styleName, styleOccurences := range map_StyleName_Occurence {
		log.Println(styleName, styleOccurences)
	}

	if paragraphStyleSlice, ok := map_StyleName_SliceOfParagraphStyle[styleName]; ok {
		for _, paragraphStyle := range paragraphStyleSlice {
			for _, paragraphProperties := range reverseMapParagraphProperties_ParagraphStyles[paragraphStyle] {
				for _, paragraph := range reverseMapParagraph_ParagraphProperties[paragraphProperties] {
					for _, rune_ := range paragraph.Runes {
						res = append(res, rune_.Text.Content)
					}
				}
			}
		}
	}

	return
}
