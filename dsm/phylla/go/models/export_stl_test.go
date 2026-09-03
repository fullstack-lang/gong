package models

import (
	"strings"
	"testing"
)

func TestGenerateSTLParts(t *testing.T) {
	stage := NewStage("test")

	lib := (&Library{
		Name:              "MainLibrary",
		NbPixPerCharacter: 8.0,
	}).Stage(stage)

	plant := (&PlantAbstract{
		Name:               "TestVase",
		PlantType:          Vase,
		RhombusSideLength:  350.0,
		RhombusInsideAngle: 127.0,
		N:                  1,
		M:                  2,
		StackHeight:        3,
	}).Stage(stage)

	lib.Plants = append(lib.Plants, plant)

	vase := (&VaseAbstract{
		Name:                          "TestVase-VaseAbstract",
		RelativeRadialThickness:       0.13,
		RelativeVerticalThickness:     0.25,
		RadialRepetitions:             3,
		RelativeHorizontalRingsHeight: 0.08,
	}).Stage(stage)

	plant.VaseAbstract = vase

	stager := &Stager{stage: stage}
	stager.enforceSemantic()

	topSTL := GenerateTopRingSTL(plant)
	if !strings.HasPrefix(topSTL, "solid TestVase_top_ring\n") || !strings.HasSuffix(topSTL, "endsolid TestVase_top_ring\n") {
		t.Errorf("Unexpected top STL output")
	}

	oneSTL := GenerateOneRingSTL(plant)
	if !strings.HasPrefix(oneSTL, "solid TestVase_one_ring\n") || !strings.HasSuffix(oneSTL, "endsolid TestVase_one_ring\n") {
		t.Errorf("Unexpected one STL output")
	}

	bottomSTL := GenerateBottomRingSTL(plant)
	if !strings.HasPrefix(bottomSTL, "solid TestVase_bottom_ring\n") || !strings.HasSuffix(bottomSTL, "endsolid TestVase_bottom_ring\n") {
		t.Errorf("Unexpected bottom STL output")
	}

	allSTL := GenerateSTL(plant)
	if !strings.HasPrefix(allSTL, "solid TestVase\n") || !strings.HasSuffix(allSTL, "endsolid TestVase\n") {
		t.Errorf("Unexpected all STL output")
	}

	facetsTop := strings.Count(topSTL, "facet normal")
	facetsOne := strings.Count(oneSTL, "facet normal")
	facetsBottom := strings.Count(bottomSTL, "facet normal")

	t.Logf("Facets count: Top=%d, One=%d, Bottom=%d", facetsTop, facetsOne, facetsBottom)

	if facetsOne <= 0 {
		t.Errorf("Expected facets for one ring, got %d", facetsOne)
	}
	if facetsTop <= facetsOne {
		t.Errorf("Expected Top ring to have more facets than one ring (%d <= %d)", facetsTop, facetsOne)
	}
	if facetsBottom <= facetsOne {
		t.Errorf("Expected Bottom ring to have more facets than one ring (%d <= %d)", facetsBottom, facetsOne)
	}
}
