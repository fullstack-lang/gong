package models

import (
	"math/rand/v2"
	"slices"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type Category1 struct {
	Name string
}

func (*Category1) IsCategory() {

}

func (shape *Category1Shape) GetCategory() *Category1 {
	return shape.Category1
}

type Category1Shape struct {
	Name      string
	Category1 *Category1

	X, Y float64

	Width, Height float64
}

type Category1ShapeProxy struct {
	shape  *Category1Shape
	stager *Stager
}

func (p *Category1ShapeProxy) RectUpdated(updatedRect *svg.Rect) {

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

type Category1NodeProxy struct {
	stager   *Stager
	diagram  *Diagram
	category *Category1
	node     *tree.Node
}

// OnAfterUpdate implements models.NodeImplInterface.
func (p *Category1NodeProxy) OnAfterUpdate(stage *tree.Stage, stagedNode *tree.Node, frontNode *tree.Node) {

	if frontNode.IsChecked && !stagedNode.IsChecked {
		s := &Category1Shape{
			Category1: p.category,
			Width:     240,
			Height:    80,
			X:         float64(int(rand.Float32()*100) + 10),
			Y:         float64(int(rand.Float32()*100) + 10),
		}
		s.Stage(p.stager.stage)
		p.diagram.Category1Shapes = append(p.diagram.Category1Shapes, s)
	}
	if !frontNode.IsChecked && stagedNode.IsChecked {

		for idx, shape := range p.diagram.Category1Shapes {
			if shape.Category1 == p.category {
				p.diagram.Category1Shapes = slices.Delete(p.diagram.Category1Shapes, idx, idx+1)
				continue
			}
		}
	}

	p.stager.stage.Commit()
}
