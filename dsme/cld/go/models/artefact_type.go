package models

import (
	"math/rand/v2"
	"slices"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type ArtefactType struct {
	Name string
}

func (*ArtefactType) IsArtElement() {

}

func (shape *ArtefactTypeShape) GetArtElement() *ArtefactType {
	return shape.ArtefactType
}

type ArtefactTypeShape struct {
	Name string

	ArtefactType *ArtefactType

	X, Y float64

	Width, Height float64
}

type ArtefactTypeShapeProxy struct {
	shape  *ArtefactTypeShape
	stager *Stager
}

func (p *ArtefactTypeShapeProxy) RectUpdated(updatedRect *svg.Rect) {

	// update the shape
	diff := p.shape.X != updatedRect.X ||
		p.shape.Y != updatedRect.Y ||
		p.shape.Width != updatedRect.Width ||
		p.shape.Height != updatedRect.Height

	// update the shape
	p.shape.X = updatedRect.X
	p.shape.Y = updatedRect.Y
	p.shape.Width = updatedRect.Width
	p.shape.Height = updatedRect.Height

	if !diff {
		p.stager.stage.CommitWithSuspendedCallbacks()
	} else {
		p.stager.stage.Commit()
	}
}

type ArtefactTypeNodeProxy struct {
	stager       *Stager
	diagram      *Diagram
	artefactType *ArtefactType
	node         *tree.Node
}

// OnAfterUpdate implements models.NodeImplInterface.
func (d *ArtefactTypeNodeProxy) OnAfterUpdate(stage *tree.Stage, stagedNode *tree.Node, frontNode *tree.Node) {

	if frontNode.IsChecked && !stagedNode.IsChecked {
		artefactTypeShape := &ArtefactTypeShape{
			ArtefactType: d.artefactType,
			Width:        150,
			Height:       25,
			X:            float64(int(rand.Float32()*100) + 10),
			Y:            float64(int(rand.Float32()*100) + 10),
		}
		artefactTypeShape.Stage(d.stager.stage)
		d.diagram.ArtefactTypeShapes = append(d.diagram.ArtefactTypeShapes, artefactTypeShape)
	}
	if !frontNode.IsChecked && stagedNode.IsChecked {

		for idx, shape := range d.diagram.ArtefactTypeShapes {
			if shape.ArtefactType == d.artefactType {
				shape.Unstage(d.stager.stage)
				d.diagram.ArtefactTypeShapes = slices.Delete(d.diagram.ArtefactTypeShapes, idx, idx+1)
				continue
			}
		}
	}

	d.stager.stage.Commit()
}
