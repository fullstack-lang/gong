package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Classdiagram:
		if stage.OnAfterClassdiagramCreateCallback != nil {
			stage.OnAfterClassdiagramCreateCallback.OnAfterCreate(stage, target)
		}
	case *DiagramPackage:
		if stage.OnAfterDiagramPackageCreateCallback != nil {
			stage.OnAfterDiagramPackageCreateCallback.OnAfterCreate(stage, target)
		}
	case *Field:
		if stage.OnAfterFieldCreateCallback != nil {
			stage.OnAfterFieldCreateCallback.OnAfterCreate(stage, target)
		}
	case *GongEnumShape:
		if stage.OnAfterGongEnumShapeCreateCallback != nil {
			stage.OnAfterGongEnumShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *GongStructShape:
		if stage.OnAfterGongStructShapeCreateCallback != nil {
			stage.OnAfterGongStructShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Link:
		if stage.OnAfterLinkCreateCallback != nil {
			stage.OnAfterLinkCreateCallback.OnAfterCreate(stage, target)
		}
	case *Node:
		if stage.OnAfterNodeCreateCallback != nil {
			stage.OnAfterNodeCreateCallback.OnAfterCreate(stage, target)
		}
	case *NoteShape:
		if stage.OnAfterNoteShapeCreateCallback != nil {
			stage.OnAfterNoteShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *NoteShapeLink:
		if stage.OnAfterNoteShapeLinkCreateCallback != nil {
			stage.OnAfterNoteShapeLinkCreateCallback.OnAfterCreate(stage, target)
		}
	case *Position:
		if stage.OnAfterPositionCreateCallback != nil {
			stage.OnAfterPositionCreateCallback.OnAfterCreate(stage, target)
		}
	case *Tree:
		if stage.OnAfterTreeCreateCallback != nil {
			stage.OnAfterTreeCreateCallback.OnAfterCreate(stage, target)
		}
	case *UmlState:
		if stage.OnAfterUmlStateCreateCallback != nil {
			stage.OnAfterUmlStateCreateCallback.OnAfterCreate(stage, target)
		}
	case *Umlsc:
		if stage.OnAfterUmlscCreateCallback != nil {
			stage.OnAfterUmlscCreateCallback.OnAfterCreate(stage, target)
		}
	case *Vertice:
		if stage.OnAfterVerticeCreateCallback != nil {
			stage.OnAfterVerticeCreateCallback.OnAfterCreate(stage, target)
		}
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *Classdiagram:
		newTarget := any(new).(*Classdiagram)
		if stage.OnAfterClassdiagramUpdateCallback != nil {
			stage.OnAfterClassdiagramUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DiagramPackage:
		newTarget := any(new).(*DiagramPackage)
		if stage.OnAfterDiagramPackageUpdateCallback != nil {
			stage.OnAfterDiagramPackageUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Field:
		newTarget := any(new).(*Field)
		if stage.OnAfterFieldUpdateCallback != nil {
			stage.OnAfterFieldUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GongEnumShape:
		newTarget := any(new).(*GongEnumShape)
		if stage.OnAfterGongEnumShapeUpdateCallback != nil {
			stage.OnAfterGongEnumShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GongStructShape:
		newTarget := any(new).(*GongStructShape)
		if stage.OnAfterGongStructShapeUpdateCallback != nil {
			stage.OnAfterGongStructShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Link:
		newTarget := any(new).(*Link)
		if stage.OnAfterLinkUpdateCallback != nil {
			stage.OnAfterLinkUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Node:
		newTarget := any(new).(*Node)
		if stage.OnAfterNodeUpdateCallback != nil {
			stage.OnAfterNodeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *NoteShape:
		newTarget := any(new).(*NoteShape)
		if stage.OnAfterNoteShapeUpdateCallback != nil {
			stage.OnAfterNoteShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *NoteShapeLink:
		newTarget := any(new).(*NoteShapeLink)
		if stage.OnAfterNoteShapeLinkUpdateCallback != nil {
			stage.OnAfterNoteShapeLinkUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Position:
		newTarget := any(new).(*Position)
		if stage.OnAfterPositionUpdateCallback != nil {
			stage.OnAfterPositionUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Tree:
		newTarget := any(new).(*Tree)
		if stage.OnAfterTreeUpdateCallback != nil {
			stage.OnAfterTreeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *UmlState:
		newTarget := any(new).(*UmlState)
		if stage.OnAfterUmlStateUpdateCallback != nil {
			stage.OnAfterUmlStateUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Umlsc:
		newTarget := any(new).(*Umlsc)
		if stage.OnAfterUmlscUpdateCallback != nil {
			stage.OnAfterUmlscUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Vertice:
		newTarget := any(new).(*Vertice)
		if stage.OnAfterVerticeUpdateCallback != nil {
			stage.OnAfterVerticeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Classdiagram:
		if stage.OnAfterClassdiagramDeleteCallback != nil {
			staged := any(staged).(*Classdiagram)
			stage.OnAfterClassdiagramDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DiagramPackage:
		if stage.OnAfterDiagramPackageDeleteCallback != nil {
			staged := any(staged).(*DiagramPackage)
			stage.OnAfterDiagramPackageDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Field:
		if stage.OnAfterFieldDeleteCallback != nil {
			staged := any(staged).(*Field)
			stage.OnAfterFieldDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GongEnumShape:
		if stage.OnAfterGongEnumShapeDeleteCallback != nil {
			staged := any(staged).(*GongEnumShape)
			stage.OnAfterGongEnumShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GongStructShape:
		if stage.OnAfterGongStructShapeDeleteCallback != nil {
			staged := any(staged).(*GongStructShape)
			stage.OnAfterGongStructShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Link:
		if stage.OnAfterLinkDeleteCallback != nil {
			staged := any(staged).(*Link)
			stage.OnAfterLinkDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Node:
		if stage.OnAfterNodeDeleteCallback != nil {
			staged := any(staged).(*Node)
			stage.OnAfterNodeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *NoteShape:
		if stage.OnAfterNoteShapeDeleteCallback != nil {
			staged := any(staged).(*NoteShape)
			stage.OnAfterNoteShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *NoteShapeLink:
		if stage.OnAfterNoteShapeLinkDeleteCallback != nil {
			staged := any(staged).(*NoteShapeLink)
			stage.OnAfterNoteShapeLinkDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Position:
		if stage.OnAfterPositionDeleteCallback != nil {
			staged := any(staged).(*Position)
			stage.OnAfterPositionDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Tree:
		if stage.OnAfterTreeDeleteCallback != nil {
			staged := any(staged).(*Tree)
			stage.OnAfterTreeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *UmlState:
		if stage.OnAfterUmlStateDeleteCallback != nil {
			staged := any(staged).(*UmlState)
			stage.OnAfterUmlStateDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Umlsc:
		if stage.OnAfterUmlscDeleteCallback != nil {
			staged := any(staged).(*Umlsc)
			stage.OnAfterUmlscDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Vertice:
		if stage.OnAfterVerticeDeleteCallback != nil {
			staged := any(staged).(*Vertice)
			stage.OnAfterVerticeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Classdiagram:
		if stage.OnAfterClassdiagramReadCallback != nil {
			stage.OnAfterClassdiagramReadCallback.OnAfterRead(stage, target)
		}
	case *DiagramPackage:
		if stage.OnAfterDiagramPackageReadCallback != nil {
			stage.OnAfterDiagramPackageReadCallback.OnAfterRead(stage, target)
		}
	case *Field:
		if stage.OnAfterFieldReadCallback != nil {
			stage.OnAfterFieldReadCallback.OnAfterRead(stage, target)
		}
	case *GongEnumShape:
		if stage.OnAfterGongEnumShapeReadCallback != nil {
			stage.OnAfterGongEnumShapeReadCallback.OnAfterRead(stage, target)
		}
	case *GongStructShape:
		if stage.OnAfterGongStructShapeReadCallback != nil {
			stage.OnAfterGongStructShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Link:
		if stage.OnAfterLinkReadCallback != nil {
			stage.OnAfterLinkReadCallback.OnAfterRead(stage, target)
		}
	case *Node:
		if stage.OnAfterNodeReadCallback != nil {
			stage.OnAfterNodeReadCallback.OnAfterRead(stage, target)
		}
	case *NoteShape:
		if stage.OnAfterNoteShapeReadCallback != nil {
			stage.OnAfterNoteShapeReadCallback.OnAfterRead(stage, target)
		}
	case *NoteShapeLink:
		if stage.OnAfterNoteShapeLinkReadCallback != nil {
			stage.OnAfterNoteShapeLinkReadCallback.OnAfterRead(stage, target)
		}
	case *Position:
		if stage.OnAfterPositionReadCallback != nil {
			stage.OnAfterPositionReadCallback.OnAfterRead(stage, target)
		}
	case *Tree:
		if stage.OnAfterTreeReadCallback != nil {
			stage.OnAfterTreeReadCallback.OnAfterRead(stage, target)
		}
	case *UmlState:
		if stage.OnAfterUmlStateReadCallback != nil {
			stage.OnAfterUmlStateReadCallback.OnAfterRead(stage, target)
		}
	case *Umlsc:
		if stage.OnAfterUmlscReadCallback != nil {
			stage.OnAfterUmlscReadCallback.OnAfterRead(stage, target)
		}
	case *Vertice:
		if stage.OnAfterVerticeReadCallback != nil {
			stage.OnAfterVerticeReadCallback.OnAfterRead(stage, target)
		}
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Classdiagram:
		stage.OnAfterClassdiagramUpdateCallback = any(callback).(OnAfterUpdateInterface[Classdiagram])
	
	case *DiagramPackage:
		stage.OnAfterDiagramPackageUpdateCallback = any(callback).(OnAfterUpdateInterface[DiagramPackage])
	
	case *Field:
		stage.OnAfterFieldUpdateCallback = any(callback).(OnAfterUpdateInterface[Field])
	
	case *GongEnumShape:
		stage.OnAfterGongEnumShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[GongEnumShape])
	
	case *GongStructShape:
		stage.OnAfterGongStructShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[GongStructShape])
	
	case *Link:
		stage.OnAfterLinkUpdateCallback = any(callback).(OnAfterUpdateInterface[Link])
	
	case *Node:
		stage.OnAfterNodeUpdateCallback = any(callback).(OnAfterUpdateInterface[Node])
	
	case *NoteShape:
		stage.OnAfterNoteShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[NoteShape])
	
	case *NoteShapeLink:
		stage.OnAfterNoteShapeLinkUpdateCallback = any(callback).(OnAfterUpdateInterface[NoteShapeLink])
	
	case *Position:
		stage.OnAfterPositionUpdateCallback = any(callback).(OnAfterUpdateInterface[Position])
	
	case *Tree:
		stage.OnAfterTreeUpdateCallback = any(callback).(OnAfterUpdateInterface[Tree])
	
	case *UmlState:
		stage.OnAfterUmlStateUpdateCallback = any(callback).(OnAfterUpdateInterface[UmlState])
	
	case *Umlsc:
		stage.OnAfterUmlscUpdateCallback = any(callback).(OnAfterUpdateInterface[Umlsc])
	
	case *Vertice:
		stage.OnAfterVerticeUpdateCallback = any(callback).(OnAfterUpdateInterface[Vertice])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Classdiagram:
		stage.OnAfterClassdiagramCreateCallback = any(callback).(OnAfterCreateInterface[Classdiagram])
	
	case *DiagramPackage:
		stage.OnAfterDiagramPackageCreateCallback = any(callback).(OnAfterCreateInterface[DiagramPackage])
	
	case *Field:
		stage.OnAfterFieldCreateCallback = any(callback).(OnAfterCreateInterface[Field])
	
	case *GongEnumShape:
		stage.OnAfterGongEnumShapeCreateCallback = any(callback).(OnAfterCreateInterface[GongEnumShape])
	
	case *GongStructShape:
		stage.OnAfterGongStructShapeCreateCallback = any(callback).(OnAfterCreateInterface[GongStructShape])
	
	case *Link:
		stage.OnAfterLinkCreateCallback = any(callback).(OnAfterCreateInterface[Link])
	
	case *Node:
		stage.OnAfterNodeCreateCallback = any(callback).(OnAfterCreateInterface[Node])
	
	case *NoteShape:
		stage.OnAfterNoteShapeCreateCallback = any(callback).(OnAfterCreateInterface[NoteShape])
	
	case *NoteShapeLink:
		stage.OnAfterNoteShapeLinkCreateCallback = any(callback).(OnAfterCreateInterface[NoteShapeLink])
	
	case *Position:
		stage.OnAfterPositionCreateCallback = any(callback).(OnAfterCreateInterface[Position])
	
	case *Tree:
		stage.OnAfterTreeCreateCallback = any(callback).(OnAfterCreateInterface[Tree])
	
	case *UmlState:
		stage.OnAfterUmlStateCreateCallback = any(callback).(OnAfterCreateInterface[UmlState])
	
	case *Umlsc:
		stage.OnAfterUmlscCreateCallback = any(callback).(OnAfterCreateInterface[Umlsc])
	
	case *Vertice:
		stage.OnAfterVerticeCreateCallback = any(callback).(OnAfterCreateInterface[Vertice])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Classdiagram:
		stage.OnAfterClassdiagramDeleteCallback = any(callback).(OnAfterDeleteInterface[Classdiagram])
	
	case *DiagramPackage:
		stage.OnAfterDiagramPackageDeleteCallback = any(callback).(OnAfterDeleteInterface[DiagramPackage])
	
	case *Field:
		stage.OnAfterFieldDeleteCallback = any(callback).(OnAfterDeleteInterface[Field])
	
	case *GongEnumShape:
		stage.OnAfterGongEnumShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[GongEnumShape])
	
	case *GongStructShape:
		stage.OnAfterGongStructShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[GongStructShape])
	
	case *Link:
		stage.OnAfterLinkDeleteCallback = any(callback).(OnAfterDeleteInterface[Link])
	
	case *Node:
		stage.OnAfterNodeDeleteCallback = any(callback).(OnAfterDeleteInterface[Node])
	
	case *NoteShape:
		stage.OnAfterNoteShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[NoteShape])
	
	case *NoteShapeLink:
		stage.OnAfterNoteShapeLinkDeleteCallback = any(callback).(OnAfterDeleteInterface[NoteShapeLink])
	
	case *Position:
		stage.OnAfterPositionDeleteCallback = any(callback).(OnAfterDeleteInterface[Position])
	
	case *Tree:
		stage.OnAfterTreeDeleteCallback = any(callback).(OnAfterDeleteInterface[Tree])
	
	case *UmlState:
		stage.OnAfterUmlStateDeleteCallback = any(callback).(OnAfterDeleteInterface[UmlState])
	
	case *Umlsc:
		stage.OnAfterUmlscDeleteCallback = any(callback).(OnAfterDeleteInterface[Umlsc])
	
	case *Vertice:
		stage.OnAfterVerticeDeleteCallback = any(callback).(OnAfterDeleteInterface[Vertice])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Classdiagram:
		stage.OnAfterClassdiagramReadCallback = any(callback).(OnAfterReadInterface[Classdiagram])
	
	case *DiagramPackage:
		stage.OnAfterDiagramPackageReadCallback = any(callback).(OnAfterReadInterface[DiagramPackage])
	
	case *Field:
		stage.OnAfterFieldReadCallback = any(callback).(OnAfterReadInterface[Field])
	
	case *GongEnumShape:
		stage.OnAfterGongEnumShapeReadCallback = any(callback).(OnAfterReadInterface[GongEnumShape])
	
	case *GongStructShape:
		stage.OnAfterGongStructShapeReadCallback = any(callback).(OnAfterReadInterface[GongStructShape])
	
	case *Link:
		stage.OnAfterLinkReadCallback = any(callback).(OnAfterReadInterface[Link])
	
	case *Node:
		stage.OnAfterNodeReadCallback = any(callback).(OnAfterReadInterface[Node])
	
	case *NoteShape:
		stage.OnAfterNoteShapeReadCallback = any(callback).(OnAfterReadInterface[NoteShape])
	
	case *NoteShapeLink:
		stage.OnAfterNoteShapeLinkReadCallback = any(callback).(OnAfterReadInterface[NoteShapeLink])
	
	case *Position:
		stage.OnAfterPositionReadCallback = any(callback).(OnAfterReadInterface[Position])
	
	case *Tree:
		stage.OnAfterTreeReadCallback = any(callback).(OnAfterReadInterface[Tree])
	
	case *UmlState:
		stage.OnAfterUmlStateReadCallback = any(callback).(OnAfterReadInterface[UmlState])
	
	case *Umlsc:
		stage.OnAfterUmlscReadCallback = any(callback).(OnAfterReadInterface[Umlsc])
	
	case *Vertice:
		stage.OnAfterVerticeReadCallback = any(callback).(OnAfterReadInterface[Vertice])
	
	}
}
