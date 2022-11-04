package models

import "log"

type ClassshapeCallbacksSingloton struct {
	ClassshapeCallback ClassshapeCallback
}

func (classshapeCallbacksSingloton *ClassshapeCallbacksSingloton) OnAfterUpdate(
	stage *StageStruct,
	stagedClassshape, frontClassshape *Classshape) {

	if stagedClassshape.IsSelected != frontClassshape.IsSelected {

		// reset the IsSelected to false
		stagedClassshape.Checkout()
		stagedClassshape.IsSelected = false
		stagedClassshape.Commit()

		log.Println("UML Shape selected ", stagedClassshape.Reference.Name)
		gongStruct, ok := (*GetGongstructInstancesMap[Reference]())[stagedClassshape.Reference.Name]
		if ok {
			if classshapeCallbacksSingloton.ClassshapeCallback != nil {
				classshapeCallbacksSingloton.ClassshapeCallback.HasSelected(gongStruct.Name)
			}
		}
	}
}

type ClassshapeCallback interface {
	HasSelected(gongstructName string)
}
