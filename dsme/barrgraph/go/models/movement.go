package models

import (
	"math/rand/v2"
	"slices"
	"time"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type Movement struct {
	Name   string
	Date   time.Time
	Places []*Place

	IsAbstract bool // some mouvements are abstracts

	// the "Modern" in "Modern Architecture" is rendered above.
	// We surmise this is a Alfred Barr"s intentional caterogy choice
	IsModern bool

	IsMajor bool // cubism
	IsMinor bool // orphism

	// For DE STILJ and NEOPLATICISM, the
	// rendering is different the "and" is in lower cases
	// and NEOPLATICISM is on the next line
	// So, for this movement, the
	// Name is "DE STILJ" and the AdditionnalName
	// AdditionnalName is NEOPLATICISM
	AdditionnalName string

	// NON GEOMETRICAL and GEOMTRICAL ABSRTACT ART
	// have no dates shown
	HideDate bool
}

func (*Movement) IsArtElement() {

}

func (shape *MovementShape) GetArtElement() *Movement {
	return shape.Movement
}

type MovementShape struct {
	Name     string
	Movement *Movement

	X, Y float64

	Width, Height float64
}

type MovementShapeProxy struct {
	shape  *MovementShape
	stager *Stager
}

func (p *MovementShapeProxy) RectUpdated(updatedRect *svg.Rect) {

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

type MovementNodeProxy struct {
	stager   *Stager
	diagram  *Diagram
	movement *Movement
	node     *tree.Node
}

// OnAfterUpdate implements models.NodeImplInterface.
func (d *MovementNodeProxy) OnAfterUpdate(stage *tree.Stage, stagedNode *tree.Node, frontNode *tree.Node) {

	if frontNode.IsChecked && !stagedNode.IsChecked {
		movementShape := &MovementShape{
			Movement: d.movement,
			Width:    240,
			Height:   80,
			X:        float64(int(rand.Float32()*100) + 10),
			Y:        float64(int(rand.Float32()*100) + 10),
		}
		movementShape.Stage(d.stager.stage)
		d.diagram.MovementShapes = append(d.diagram.MovementShapes, movementShape)
	}
	if !frontNode.IsChecked && stagedNode.IsChecked {

		for idx, shape := range d.diagram.MovementShapes {
			if shape.Movement == d.movement {
				d.diagram.MovementShapes = slices.Delete(d.diagram.MovementShapes, idx, idx+1)
				continue
			}
		}
	}

	d.stager.stage.Commit()
}
