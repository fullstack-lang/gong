// generated code - do not edit
package models

// insertion point
func (inst *Chapter) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Content":
		switch reverseField.Fieldname {
		case "Chapters":
			if _content, ok := stage.Content_Chapters_reverseMap[inst]; ok {
				res = _content.Name
			}
		}
	}
	return
}

func (inst *Content) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *DownloadableFile) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *JpgImage) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Page) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Chapter":
		switch reverseField.Fieldname {
		case "Pages":
			if _chapter, ok := stage.Chapter_Pages_reverseMap[inst]; ok {
				res = _chapter.Name
			}
		}
	}
	return
}

func (inst *PngImage) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Section) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Page":
		switch reverseField.Fieldname {
		case "Sections":
			if _page, ok := stage.Page_Sections_reverseMap[inst]; ok {
				res = _page.Name
			}
		}
	}
	return
}

func (inst *SvgImage) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

// insertion point
func (inst *Chapter) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Content":
		switch reverseField.Fieldname {
		case "Chapters":
			res = stage.Content_Chapters_reverseMap[inst]
		}
	}
	return res
}

func (inst *Content) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *DownloadableFile) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *JpgImage) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Page) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Chapter":
		switch reverseField.Fieldname {
		case "Pages":
			res = stage.Chapter_Pages_reverseMap[inst]
		}
	}
	return res
}

func (inst *PngImage) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Section) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Page":
		switch reverseField.Fieldname {
		case "Sections":
			res = stage.Page_Sections_reverseMap[inst]
		}
	}
	return res
}

func (inst *SvgImage) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}
