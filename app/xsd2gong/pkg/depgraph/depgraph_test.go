package depgraph_test

import (
	"path/filepath"
	"testing"

	"github.com/fullstack-lang/gong/app/xsd2gong/pkg/depgraph"
	"github.com/fullstack-lang/gong/app/xsd2gong/pkg/namer"
	"github.com/fullstack-lang/gong/app/xsd2gong/pkg/xsdloader"
)

func TestMusicXMLSchemaDAGAndPartitioning(t *testing.T) {
	xsdPath := filepath.Join("..", "..", "testbed", "musicxml", "schemas", "musicxml.xsd")
	schemaSet, err := xsdloader.Load(xsdPath)
	if err != nil {
		t.Fatalf("xsdloader.Load(%s) failed: %v", xsdPath, err)
	}
	if len(schemaSet.Schemas) < 3 {
		t.Fatalf("expected at least 3 loaded schemas (musicxml.xsd, xlink.xsd, xml.xsd), got %d", len(schemaSet.Schemas))
	}

	g, err := depgraph.Build(schemaSet)
	if err != nil {
		t.Fatalf("depgraph.Build failed: %v", err)
	}

	expectedOrder := []string{"xlink", "common", "barline", "attributes", "direction", "note", "score"}
	if len(g.TopologicalOrd) != len(expectedOrder) {
		t.Fatalf("expected %d packages in topological order, got %d", len(expectedOrder), len(g.TopologicalOrd))
	}
	for i, exp := range expectedOrder {
		if g.TopologicalOrd[i].Name != exp {
			t.Errorf("topological package[%d] = %q, want %q", i, g.TopologicalOrd[i].Name, exp)
		}
	}

	// Verify strict DAG property: no package may import a package of equal or higher topological rank
	rankOf := make(map[string]int)
	for i, pkg := range g.TopologicalOrd {
		rankOf[pkg.Name] = i
	}
	for _, pkg := range g.TopologicalOrd {
		for _, dep := range pkg.Imports {
			if rankOf[dep] >= rankOf[pkg.Name] {
				t.Errorf("DAG violation: package %q (rank %d) imports %q (rank %d)",
					pkg.Name, rankOf[pkg.Name], dep, rankOf[dep])
			}
		}
	}
}

func TestNamerDeterministicCollisionResolution(t *testing.T) {
	scope := namer.NewScope()
	scope.Reserve("Name")

	if got := scope.Unique("name"); got != "NameAttr" {
		t.Errorf("expected reserved collision 'name' -> 'NameAttr', got %q", got)
	}
	if got := scope.Unique("score-partwise"); got != "ScorePartwise" {
		t.Errorf("expected 'score-partwise' -> 'ScorePartwise', got %q", got)
	}
	if got := scope.Unique("score_partwise"); got != "ScorePartwise_1" {
		t.Errorf("expected second 'score_partwise' -> 'ScorePartwise_1', got %q", got)
	}
}
