package tests

import (
	"testing"

	"github.com/fullstack-lang/gong/stacks/gongdoc/go/models"
)

func TestUmarshallDocument(t *testing.T) {

	var pkgelt models.Pkgelt

	// unmarshall diagrams of the document
	pkgelt.Unmarshall("geometry/diagrams")

}
