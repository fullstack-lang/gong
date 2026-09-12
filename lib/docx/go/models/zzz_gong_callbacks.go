// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront[Type Gongstruct](instance *Type) {

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

// AfterCreateFromFront is a backward-compatible package-level forwarder.
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {
	stage.AfterCreateFromFront(instance)
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront[Type Gongstruct](old, new *Type) {

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

// OnAfterUpdateFromFront is a backward-compatible package-level forwarder.
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {
	stage.OnAfterUpdateFromFront(old, new)
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront[Type Gongstruct](staged, front *Type) {

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

// AfterDeleteFromFront is a backward-compatible package-level forwarder.
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {
	stage.AfterDeleteFromFront(staged, front)
}
