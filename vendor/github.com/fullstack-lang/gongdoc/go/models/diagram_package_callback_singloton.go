package models

import "log"

type DiagramPackageCallbacksSingloton struct {
	DiagramPackageCallback DiagramPackageCallback
}

func (classshapeCallbacksSingloton *DiagramPackageCallbacksSingloton) OnAfterUpdate(
	stage *StageStruct,
	stagedDiagramPackage, frontDiagramPackage *DiagramPackage) {

	if stagedDiagramPackage.IsReloaded != frontDiagramPackage.IsReloaded {

		// reset the IsReloaded to false
		stagedDiagramPackage.Checkout()
		stagedDiagramPackage.IsReloaded = false
		stagedDiagramPackage.Commit()

		log.Println("Reload requested")
		if stagedDiagramPackage.IsEditable {
			stagedDiagramPackage.Reload()
		}
	}
}

type DiagramPackageCallback interface {
	HasSelected(gongstructName string)
}
