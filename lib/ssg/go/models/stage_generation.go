package models

func (stage *Stage) Generation() {

	contents := GetGongstructInstancesSet[Content](stage)

	if len(*contents) != 1 {
		return
	}

	var content *Content
	for k, _ := range *contents {
		content = k
	}

	_ = content
	contentPath := content.ContentPath

}
