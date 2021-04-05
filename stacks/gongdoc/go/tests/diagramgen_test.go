package tests

import (
	"testing"

	"github.com/fullstack-lang/gong/stacks/gongdoc/go/models"
)

func TestGenerateSVG(t *testing.T) {

	var pkgelt models.Pkgelt

	pkgelt.Unmarshall("geometry/diagrams")

	for _, classDiagram := range pkgelt.Classdiagrams {
		classDiagram.OutputSVG("geometry/diagrams")
	}
}
