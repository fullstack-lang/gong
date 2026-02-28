package models

import (
	"math/rand/v2"
	"slices"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type Category3 struct {
	Name string
}

func (*Category3) IsCategory() {

}

func (shape *Category3Shape) GetCategory() *Category3 {
	return shape.Category3
}

type Category3Shape struct {
	Name string

	Category3 *Category3

	X, Y float64

	Width, Height float64
}

type Category3ShapeProxy struct {
	shape  *Category3Shape
	stager *Stager
}

func (p *Category3ShapeProxy) RectUpdated(updatedRect *svg.Rect) {

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

type Category3NodeProxy struct {
	stager   *Stager
	diagram  *Diagram
	category *Category3
	node     *tree.Node
}

// OnAfterUpdate implements models.NodeImplInterface.
func (d *Category3NodeProxy) OnAfterUpdate(stage *tree.Stage, stagedNode *tree.Node, frontNode *tree.Node) {

	if frontNode.IsChecked && !stagedNode.IsChecked {
		s := &Category3Shape{
			Category3: d.category,
			Width:     150,
			Height:    25,
			X:         float64(int(rand.Float32()*100) + 10),
			Y:         float64(int(rand.Float32()*100) + 10),
		}
		s.Stage(d.stager.stage)
		d.diagram.Category3Shapes = append(d.diagram.Category3Shapes, s)
	}
	if !frontNode.IsChecked && stagedNode.IsChecked {

		for idx, shape := range d.diagram.Category3Shapes {
			if shape.Category3 == d.category {
				shape.Unstage(d.stager.stage)
				d.diagram.Category3Shapes = slices.Delete(d.diagram.Category3Shapes, idx, idx+1)
				continue
			}
		}
	}

	d.stager.stage.Commit()
}
