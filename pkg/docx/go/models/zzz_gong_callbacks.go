// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Body:
		if stage.OnAfterBodyCreateCallback != nil {
			stage.OnAfterBodyCreateCallback.OnAfterCreate(stage, target)
		}
	case *Document:
		if stage.OnAfterDocumentCreateCallback != nil {
			stage.OnAfterDocumentCreateCallback.OnAfterCreate(stage, target)
		}
	case *Docx:
		if stage.OnAfterDocxCreateCallback != nil {
			stage.OnAfterDocxCreateCallback.OnAfterCreate(stage, target)
		}
	case *File:
		if stage.OnAfterFileCreateCallback != nil {
			stage.OnAfterFileCreateCallback.OnAfterCreate(stage, target)
		}
	case *Node:
		if stage.OnAfterNodeCreateCallback != nil {
			stage.OnAfterNodeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Paragraph:
		if stage.OnAfterParagraphCreateCallback != nil {
			stage.OnAfterParagraphCreateCallback.OnAfterCreate(stage, target)
		}
	case *ParagraphProperties:
		if stage.OnAfterParagraphPropertiesCreateCallback != nil {
			stage.OnAfterParagraphPropertiesCreateCallback.OnAfterCreate(stage, target)
		}
	case *ParagraphStyle:
		if stage.OnAfterParagraphStyleCreateCallback != nil {
			stage.OnAfterParagraphStyleCreateCallback.OnAfterCreate(stage, target)
		}
	case *Rune:
		if stage.OnAfterRuneCreateCallback != nil {
			stage.OnAfterRuneCreateCallback.OnAfterCreate(stage, target)
		}
	case *RuneProperties:
		if stage.OnAfterRunePropertiesCreateCallback != nil {
			stage.OnAfterRunePropertiesCreateCallback.OnAfterCreate(stage, target)
		}
	case *Table:
		if stage.OnAfterTableCreateCallback != nil {
			stage.OnAfterTableCreateCallback.OnAfterCreate(stage, target)
		}
	case *TableColumn:
		if stage.OnAfterTableColumnCreateCallback != nil {
			stage.OnAfterTableColumnCreateCallback.OnAfterCreate(stage, target)
		}
	case *TableProperties:
		if stage.OnAfterTablePropertiesCreateCallback != nil {
			stage.OnAfterTablePropertiesCreateCallback.OnAfterCreate(stage, target)
		}
	case *TableRow:
		if stage.OnAfterTableRowCreateCallback != nil {
			stage.OnAfterTableRowCreateCallback.OnAfterCreate(stage, target)
		}
	case *TableStyle:
		if stage.OnAfterTableStyleCreateCallback != nil {
			stage.OnAfterTableStyleCreateCallback.OnAfterCreate(stage, target)
		}
	case *Text:
		if stage.OnAfterTextCreateCallback != nil {
			stage.OnAfterTextCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is called after a update from front
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *Body:
		newTarget := any(new).(*Body)
		if stage.OnAfterBodyUpdateCallback != nil {
			stage.OnAfterBodyUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Document:
		newTarget := any(new).(*Document)
		if stage.OnAfterDocumentUpdateCallback != nil {
			stage.OnAfterDocumentUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Docx:
		newTarget := any(new).(*Docx)
		if stage.OnAfterDocxUpdateCallback != nil {
			stage.OnAfterDocxUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *File:
		newTarget := any(new).(*File)
		if stage.OnAfterFileUpdateCallback != nil {
			stage.OnAfterFileUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Node:
		newTarget := any(new).(*Node)
		if stage.OnAfterNodeUpdateCallback != nil {
			stage.OnAfterNodeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Paragraph:
		newTarget := any(new).(*Paragraph)
		if stage.OnAfterParagraphUpdateCallback != nil {
			stage.OnAfterParagraphUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ParagraphProperties:
		newTarget := any(new).(*ParagraphProperties)
		if stage.OnAfterParagraphPropertiesUpdateCallback != nil {
			stage.OnAfterParagraphPropertiesUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ParagraphStyle:
		newTarget := any(new).(*ParagraphStyle)
		if stage.OnAfterParagraphStyleUpdateCallback != nil {
			stage.OnAfterParagraphStyleUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Rune:
		newTarget := any(new).(*Rune)
		if stage.OnAfterRuneUpdateCallback != nil {
			stage.OnAfterRuneUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *RuneProperties:
		newTarget := any(new).(*RuneProperties)
		if stage.OnAfterRunePropertiesUpdateCallback != nil {
			stage.OnAfterRunePropertiesUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Table:
		newTarget := any(new).(*Table)
		if stage.OnAfterTableUpdateCallback != nil {
			stage.OnAfterTableUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TableColumn:
		newTarget := any(new).(*TableColumn)
		if stage.OnAfterTableColumnUpdateCallback != nil {
			stage.OnAfterTableColumnUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TableProperties:
		newTarget := any(new).(*TableProperties)
		if stage.OnAfterTablePropertiesUpdateCallback != nil {
			stage.OnAfterTablePropertiesUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TableRow:
		newTarget := any(new).(*TableRow)
		if stage.OnAfterTableRowUpdateCallback != nil {
			stage.OnAfterTableRowUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TableStyle:
		newTarget := any(new).(*TableStyle)
		if stage.OnAfterTableStyleUpdateCallback != nil {
			stage.OnAfterTableStyleUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Text:
		newTarget := any(new).(*Text)
		if stage.OnAfterTextUpdateCallback != nil {
			stage.OnAfterTextUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Body:
		if stage.OnAfterBodyDeleteCallback != nil {
			staged := any(staged).(*Body)
			stage.OnAfterBodyDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Document:
		if stage.OnAfterDocumentDeleteCallback != nil {
			staged := any(staged).(*Document)
			stage.OnAfterDocumentDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Docx:
		if stage.OnAfterDocxDeleteCallback != nil {
			staged := any(staged).(*Docx)
			stage.OnAfterDocxDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *File:
		if stage.OnAfterFileDeleteCallback != nil {
			staged := any(staged).(*File)
			stage.OnAfterFileDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Node:
		if stage.OnAfterNodeDeleteCallback != nil {
			staged := any(staged).(*Node)
			stage.OnAfterNodeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Paragraph:
		if stage.OnAfterParagraphDeleteCallback != nil {
			staged := any(staged).(*Paragraph)
			stage.OnAfterParagraphDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ParagraphProperties:
		if stage.OnAfterParagraphPropertiesDeleteCallback != nil {
			staged := any(staged).(*ParagraphProperties)
			stage.OnAfterParagraphPropertiesDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ParagraphStyle:
		if stage.OnAfterParagraphStyleDeleteCallback != nil {
			staged := any(staged).(*ParagraphStyle)
			stage.OnAfterParagraphStyleDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Rune:
		if stage.OnAfterRuneDeleteCallback != nil {
			staged := any(staged).(*Rune)
			stage.OnAfterRuneDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *RuneProperties:
		if stage.OnAfterRunePropertiesDeleteCallback != nil {
			staged := any(staged).(*RuneProperties)
			stage.OnAfterRunePropertiesDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Table:
		if stage.OnAfterTableDeleteCallback != nil {
			staged := any(staged).(*Table)
			stage.OnAfterTableDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TableColumn:
		if stage.OnAfterTableColumnDeleteCallback != nil {
			staged := any(staged).(*TableColumn)
			stage.OnAfterTableColumnDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TableProperties:
		if stage.OnAfterTablePropertiesDeleteCallback != nil {
			staged := any(staged).(*TableProperties)
			stage.OnAfterTablePropertiesDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TableRow:
		if stage.OnAfterTableRowDeleteCallback != nil {
			staged := any(staged).(*TableRow)
			stage.OnAfterTableRowDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TableStyle:
		if stage.OnAfterTableStyleDeleteCallback != nil {
			staged := any(staged).(*TableStyle)
			stage.OnAfterTableStyleDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Text:
		if stage.OnAfterTextDeleteCallback != nil {
			staged := any(staged).(*Text)
			stage.OnAfterTextDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Body:
		if stage.OnAfterBodyReadCallback != nil {
			stage.OnAfterBodyReadCallback.OnAfterRead(stage, target)
		}
	case *Document:
		if stage.OnAfterDocumentReadCallback != nil {
			stage.OnAfterDocumentReadCallback.OnAfterRead(stage, target)
		}
	case *Docx:
		if stage.OnAfterDocxReadCallback != nil {
			stage.OnAfterDocxReadCallback.OnAfterRead(stage, target)
		}
	case *File:
		if stage.OnAfterFileReadCallback != nil {
			stage.OnAfterFileReadCallback.OnAfterRead(stage, target)
		}
	case *Node:
		if stage.OnAfterNodeReadCallback != nil {
			stage.OnAfterNodeReadCallback.OnAfterRead(stage, target)
		}
	case *Paragraph:
		if stage.OnAfterParagraphReadCallback != nil {
			stage.OnAfterParagraphReadCallback.OnAfterRead(stage, target)
		}
	case *ParagraphProperties:
		if stage.OnAfterParagraphPropertiesReadCallback != nil {
			stage.OnAfterParagraphPropertiesReadCallback.OnAfterRead(stage, target)
		}
	case *ParagraphStyle:
		if stage.OnAfterParagraphStyleReadCallback != nil {
			stage.OnAfterParagraphStyleReadCallback.OnAfterRead(stage, target)
		}
	case *Rune:
		if stage.OnAfterRuneReadCallback != nil {
			stage.OnAfterRuneReadCallback.OnAfterRead(stage, target)
		}
	case *RuneProperties:
		if stage.OnAfterRunePropertiesReadCallback != nil {
			stage.OnAfterRunePropertiesReadCallback.OnAfterRead(stage, target)
		}
	case *Table:
		if stage.OnAfterTableReadCallback != nil {
			stage.OnAfterTableReadCallback.OnAfterRead(stage, target)
		}
	case *TableColumn:
		if stage.OnAfterTableColumnReadCallback != nil {
			stage.OnAfterTableColumnReadCallback.OnAfterRead(stage, target)
		}
	case *TableProperties:
		if stage.OnAfterTablePropertiesReadCallback != nil {
			stage.OnAfterTablePropertiesReadCallback.OnAfterRead(stage, target)
		}
	case *TableRow:
		if stage.OnAfterTableRowReadCallback != nil {
			stage.OnAfterTableRowReadCallback.OnAfterRead(stage, target)
		}
	case *TableStyle:
		if stage.OnAfterTableStyleReadCallback != nil {
			stage.OnAfterTableStyleReadCallback.OnAfterRead(stage, target)
		}
	case *Text:
		if stage.OnAfterTextReadCallback != nil {
			stage.OnAfterTextReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *Stage, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Body:
		stage.OnAfterBodyUpdateCallback = any(callback).(OnAfterUpdateInterface[Body])
	
	case *Document:
		stage.OnAfterDocumentUpdateCallback = any(callback).(OnAfterUpdateInterface[Document])
	
	case *Docx:
		stage.OnAfterDocxUpdateCallback = any(callback).(OnAfterUpdateInterface[Docx])
	
	case *File:
		stage.OnAfterFileUpdateCallback = any(callback).(OnAfterUpdateInterface[File])
	
	case *Node:
		stage.OnAfterNodeUpdateCallback = any(callback).(OnAfterUpdateInterface[Node])
	
	case *Paragraph:
		stage.OnAfterParagraphUpdateCallback = any(callback).(OnAfterUpdateInterface[Paragraph])
	
	case *ParagraphProperties:
		stage.OnAfterParagraphPropertiesUpdateCallback = any(callback).(OnAfterUpdateInterface[ParagraphProperties])
	
	case *ParagraphStyle:
		stage.OnAfterParagraphStyleUpdateCallback = any(callback).(OnAfterUpdateInterface[ParagraphStyle])
	
	case *Rune:
		stage.OnAfterRuneUpdateCallback = any(callback).(OnAfterUpdateInterface[Rune])
	
	case *RuneProperties:
		stage.OnAfterRunePropertiesUpdateCallback = any(callback).(OnAfterUpdateInterface[RuneProperties])
	
	case *Table:
		stage.OnAfterTableUpdateCallback = any(callback).(OnAfterUpdateInterface[Table])
	
	case *TableColumn:
		stage.OnAfterTableColumnUpdateCallback = any(callback).(OnAfterUpdateInterface[TableColumn])
	
	case *TableProperties:
		stage.OnAfterTablePropertiesUpdateCallback = any(callback).(OnAfterUpdateInterface[TableProperties])
	
	case *TableRow:
		stage.OnAfterTableRowUpdateCallback = any(callback).(OnAfterUpdateInterface[TableRow])
	
	case *TableStyle:
		stage.OnAfterTableStyleUpdateCallback = any(callback).(OnAfterUpdateInterface[TableStyle])
	
	case *Text:
		stage.OnAfterTextUpdateCallback = any(callback).(OnAfterUpdateInterface[Text])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Body:
		stage.OnAfterBodyCreateCallback = any(callback).(OnAfterCreateInterface[Body])
	
	case *Document:
		stage.OnAfterDocumentCreateCallback = any(callback).(OnAfterCreateInterface[Document])
	
	case *Docx:
		stage.OnAfterDocxCreateCallback = any(callback).(OnAfterCreateInterface[Docx])
	
	case *File:
		stage.OnAfterFileCreateCallback = any(callback).(OnAfterCreateInterface[File])
	
	case *Node:
		stage.OnAfterNodeCreateCallback = any(callback).(OnAfterCreateInterface[Node])
	
	case *Paragraph:
		stage.OnAfterParagraphCreateCallback = any(callback).(OnAfterCreateInterface[Paragraph])
	
	case *ParagraphProperties:
		stage.OnAfterParagraphPropertiesCreateCallback = any(callback).(OnAfterCreateInterface[ParagraphProperties])
	
	case *ParagraphStyle:
		stage.OnAfterParagraphStyleCreateCallback = any(callback).(OnAfterCreateInterface[ParagraphStyle])
	
	case *Rune:
		stage.OnAfterRuneCreateCallback = any(callback).(OnAfterCreateInterface[Rune])
	
	case *RuneProperties:
		stage.OnAfterRunePropertiesCreateCallback = any(callback).(OnAfterCreateInterface[RuneProperties])
	
	case *Table:
		stage.OnAfterTableCreateCallback = any(callback).(OnAfterCreateInterface[Table])
	
	case *TableColumn:
		stage.OnAfterTableColumnCreateCallback = any(callback).(OnAfterCreateInterface[TableColumn])
	
	case *TableProperties:
		stage.OnAfterTablePropertiesCreateCallback = any(callback).(OnAfterCreateInterface[TableProperties])
	
	case *TableRow:
		stage.OnAfterTableRowCreateCallback = any(callback).(OnAfterCreateInterface[TableRow])
	
	case *TableStyle:
		stage.OnAfterTableStyleCreateCallback = any(callback).(OnAfterCreateInterface[TableStyle])
	
	case *Text:
		stage.OnAfterTextCreateCallback = any(callback).(OnAfterCreateInterface[Text])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Body:
		stage.OnAfterBodyDeleteCallback = any(callback).(OnAfterDeleteInterface[Body])
	
	case *Document:
		stage.OnAfterDocumentDeleteCallback = any(callback).(OnAfterDeleteInterface[Document])
	
	case *Docx:
		stage.OnAfterDocxDeleteCallback = any(callback).(OnAfterDeleteInterface[Docx])
	
	case *File:
		stage.OnAfterFileDeleteCallback = any(callback).(OnAfterDeleteInterface[File])
	
	case *Node:
		stage.OnAfterNodeDeleteCallback = any(callback).(OnAfterDeleteInterface[Node])
	
	case *Paragraph:
		stage.OnAfterParagraphDeleteCallback = any(callback).(OnAfterDeleteInterface[Paragraph])
	
	case *ParagraphProperties:
		stage.OnAfterParagraphPropertiesDeleteCallback = any(callback).(OnAfterDeleteInterface[ParagraphProperties])
	
	case *ParagraphStyle:
		stage.OnAfterParagraphStyleDeleteCallback = any(callback).(OnAfterDeleteInterface[ParagraphStyle])
	
	case *Rune:
		stage.OnAfterRuneDeleteCallback = any(callback).(OnAfterDeleteInterface[Rune])
	
	case *RuneProperties:
		stage.OnAfterRunePropertiesDeleteCallback = any(callback).(OnAfterDeleteInterface[RuneProperties])
	
	case *Table:
		stage.OnAfterTableDeleteCallback = any(callback).(OnAfterDeleteInterface[Table])
	
	case *TableColumn:
		stage.OnAfterTableColumnDeleteCallback = any(callback).(OnAfterDeleteInterface[TableColumn])
	
	case *TableProperties:
		stage.OnAfterTablePropertiesDeleteCallback = any(callback).(OnAfterDeleteInterface[TableProperties])
	
	case *TableRow:
		stage.OnAfterTableRowDeleteCallback = any(callback).(OnAfterDeleteInterface[TableRow])
	
	case *TableStyle:
		stage.OnAfterTableStyleDeleteCallback = any(callback).(OnAfterDeleteInterface[TableStyle])
	
	case *Text:
		stage.OnAfterTextDeleteCallback = any(callback).(OnAfterDeleteInterface[Text])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Body:
		stage.OnAfterBodyReadCallback = any(callback).(OnAfterReadInterface[Body])
	
	case *Document:
		stage.OnAfterDocumentReadCallback = any(callback).(OnAfterReadInterface[Document])
	
	case *Docx:
		stage.OnAfterDocxReadCallback = any(callback).(OnAfterReadInterface[Docx])
	
	case *File:
		stage.OnAfterFileReadCallback = any(callback).(OnAfterReadInterface[File])
	
	case *Node:
		stage.OnAfterNodeReadCallback = any(callback).(OnAfterReadInterface[Node])
	
	case *Paragraph:
		stage.OnAfterParagraphReadCallback = any(callback).(OnAfterReadInterface[Paragraph])
	
	case *ParagraphProperties:
		stage.OnAfterParagraphPropertiesReadCallback = any(callback).(OnAfterReadInterface[ParagraphProperties])
	
	case *ParagraphStyle:
		stage.OnAfterParagraphStyleReadCallback = any(callback).(OnAfterReadInterface[ParagraphStyle])
	
	case *Rune:
		stage.OnAfterRuneReadCallback = any(callback).(OnAfterReadInterface[Rune])
	
	case *RuneProperties:
		stage.OnAfterRunePropertiesReadCallback = any(callback).(OnAfterReadInterface[RuneProperties])
	
	case *Table:
		stage.OnAfterTableReadCallback = any(callback).(OnAfterReadInterface[Table])
	
	case *TableColumn:
		stage.OnAfterTableColumnReadCallback = any(callback).(OnAfterReadInterface[TableColumn])
	
	case *TableProperties:
		stage.OnAfterTablePropertiesReadCallback = any(callback).(OnAfterReadInterface[TableProperties])
	
	case *TableRow:
		stage.OnAfterTableRowReadCallback = any(callback).(OnAfterReadInterface[TableRow])
	
	case *TableStyle:
		stage.OnAfterTableStyleReadCallback = any(callback).(OnAfterReadInterface[TableStyle])
	
	case *Text:
		stage.OnAfterTextReadCallback = any(callback).(OnAfterReadInterface[Text])
	
	}
}
