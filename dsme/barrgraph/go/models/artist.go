package models

import (
	"math/rand/v2"
	"slices"
	"time"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type Artist struct {
	Name string

	IsDead      bool
	DateOfDeath time.Time
	Place       *Place
}

func (*Artist) IsArtElement() {
}

type ArtistShape struct {
	Name   string
	Artist *Artist

	X, Y float64

	Width, Height float64
}

func (shape *ArtistShape) GetArtElement() *Artist {
	return shape.Artist
}

type ArtistNodeProxy struct {
	stager  *Stager
	diagram *Diagram
	artist  *Artist
	node    *tree.Node
}

// OnAfterUpdate implements models.NodeImplInterface.
func (d *ArtistNodeProxy) OnAfterUpdate(stage *tree.Stage, stagedNode *tree.Node, frontNode *tree.Node) {
	if frontNode.IsChecked && !stagedNode.IsChecked {
		artistShape := &ArtistShape{
			Artist: d.artist,
			Width:  80,
			Height: 30,
			X:      float64(int(rand.Float32()*100) + 10),
			Y:      float64(int(rand.Float32()*100) + 10),
		}
		artistShape.Stage(d.stager.stage)
		d.diagram.ArtistShapes = append(d.diagram.ArtistShapes, artistShape)
	}
	if !frontNode.IsChecked && stagedNode.IsChecked {
		for idx, shape := range d.diagram.ArtistShapes {
			if shape.Artist == d.artist {
				shape.Unstage(d.stager.stage)
				d.diagram.ArtistShapes = slices.Delete(d.diagram.ArtistShapes, idx, idx+1)
				continue
			}
		}
	}

	d.stager.stage.Commit()
}
