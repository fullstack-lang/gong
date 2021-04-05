package tests

import (
	"testing"

	"github.com/fullstack-lang/gong/stacks/gongdoc/go/models"
	"github.com/fullstack-lang/gong/stacks/gongdoc/go/tests/geometry/diagrams"
)

func TestUnmarshall(t *testing.T) {

	// scope diagram package
	var pkgelt models.Pkgelt
	pkgelt.Classdiagrams = make([]*models.Classdiagram, 0)
	pkgelt.Unmarshall("geometry/diagrams")

	{
		got := len(pkgelt.Classdiagrams)
		if got != 2 {
			t.Errorf("Number of diagrams, got = %d; want 2", got)
		}
	}
	{
		got := pkgelt.Classdiagrams[0].Classshapes[0].Position.X
		want := diagrams.Diagram1.Classshapes[0].Position.X
		if got != want {
			t.Errorf("got = %f; want %f", got, want)
		}
	}
	{
		got := pkgelt.Umlscs[0].States[0].X
		want := diagrams.UmlscDiagram1.States[0].X
		if got != want {
			t.Errorf("got= %f; want %f", got, want)
		}
	}
	{
		got := pkgelt.Umlscs[0].States[0].Name
		want := diagrams.UmlscDiagram1.States[0].Name
		if got != want {
			t.Errorf("got= %s; want %s", got, want)
		}
	}
	{
		got := pkgelt.Umlscs[0].Activestate
		want := diagrams.UmlscDiagram1.Activestate
		if got != want {
			t.Errorf("got= %s; want %s", got, want)
		}
	}

}

func TestMarshall(t *testing.T) {

	var pkgelt models.Pkgelt

	pkgelt.Unmarshall("geometry/diagrams")

	pkgelt.Marshall("geometry/diagrams")

}
