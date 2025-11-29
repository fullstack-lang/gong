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

