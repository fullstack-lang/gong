package models

import tree "github.com/fullstack-lang/gong/lib/tree/go/models"

type DiagramRenameButtonProxy struct {
	stager     *Stager
	diagram    *Diagram
	buttonType ButtonType
}

func (p *DiagramRenameButtonProxy) ButtonUpdated(stage *tree.Stage,
	button *tree.Button, updatedButton *tree.Button) {
	// perfectly boring code
	switch p.buttonType {
	case RENAME:
		p.diagram.IsInRenameMode = true
	case RENAME_CANCEL:
		p.diagram.IsInRenameMode = false
	}

	p.stager.stage.Commit()
}

type ButtonType int

const (
	DUPLICATE ButtonType = iota
	EDIT_CANCEL
	EDIT
	REMOVE
	RENAME_CANCEL
	RENAME
	SAVE
)
