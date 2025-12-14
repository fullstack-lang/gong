package models

import (
	"math/rand/v2"
	"slices"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type Category2 struct {
	Name string
}

func (*Category2) IsCategory() {

}

type Category2Shape struct {
	Name      string
	Category2 *Category2

	X, Y float64

	Width, Height float64
}

func (shape *Category2Shape) GetCategory() *Category2 {
	return shape.Category2
}

type Category2Proxy struct {
	shape  *Category2Shape
	stager *Stager
}

func (p *Category2Proxy) RectUpdated(updatedRect *svg.Rect) {

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

type Category2NodeProxy struct {
	stager   *Stager
	diagram  *Diagram
	category *Category2
	node     *tree.Node
}

// OnAfterUpdate implements models.NodeImplInterface.
func (d *Category2NodeProxy) OnAfterUpdate(stage *tree.Stage, stagedNode *tree.Node, frontNode *tree.Node) {

	if frontNode.IsChecked && !stagedNode.IsChecked {
		s := &Category2Shape{
			Category2: d.category,
			Width:     80,
			Height:    30,
			X:         float64(int(rand.Float32()*100) + 10),
			Y:         float64(int(rand.Float32()*100) + 10),
		}
		s.Stage(d.stager.stage)
		d.diagram.Category2Shapes = append(d.diagram.Category2Shapes, s)
	}
	if !frontNode.IsChecked && stagedNode.IsChecked {

		for idx, shape := range d.diagram.Category2Shapes {
			if shape.Category2 == d.category {
				shape.Unstage(d.stager.stage)
				d.diagram.Category2Shapes = slices.Delete(d.diagram.Category2Shapes, idx, idx+1)
				continue
			}
		}
	}

	d.stager.stage.Commit()
}
