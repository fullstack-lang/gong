package golang

import (
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"

	test3_models "github.com/fullstack-lang/gong/test/test3/go/models"
)

func TestMergeStageBytes_Basic(t *testing.T) {
	file1 := `package main

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/test/test3/go/models"
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]
)

func _(stage *models.Stage) {
	// insertion point for declaration of instances to stage
	__A__00000000_ := (&models.A{Name: "A0"}).Stage(stage)
	__A__00000001_ := (&models.A{Name: "A1"}).Stage(stage)

	// insertion point for initialization of values
	__A__00000000_.Name = "A0"
	__A__00000000_.IntValue = 10
	__A__00000001_.Name = "A1"
	__A__00000001_.IntValue = 20

	// insertion point for setup of pointers
	__A__00000000_.B = nil
	__A__00000001_.B = nil
}
`

	file2 := `package main

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/test/test3/go/models"
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]
)

func _(stage *models.Stage) {
	// insertion point for declaration of instances to stage
	__A__00000000_ := (&models.A{Name: "A2"}).Stage(stage)
	__B__00000000_ := (&models.B{Name: "B0"}).Stage(stage)

	// insertion point for initialization of values
	__A__00000000_.Name = "A2"
	__A__00000000_.IntValue = 30
	__B__00000000_.Name = "B0"

	// insertion point for setup of pointers
	__A__00000000_.B = __B__00000000_
}
`

	merged, err := MergeStageBytesToString([][]byte{[]byte(file1), []byte(file2)}, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Verify syntax correctness
	fset := token.NewFileSet()
	_, err = parser.ParseFile(fset, "merged.go", merged, parser.ParseComments)
	if err != nil {
		t.Fatalf("merged file has syntax errors:\n%s\nError: %v", merged, err)
	}

	// Check declarations
	if !strings.Contains(merged, "__A__00000000_ := (&models.A{Name: \"A0\"}).Stage(stage)") {
		t.Errorf("missing declaration of __A__00000000_ in:\n%s", merged)
	}
	if !strings.Contains(merged, "__A__00000001_ := (&models.A{Name: \"A1\"}).Stage(stage)") {
		t.Errorf("missing declaration of __A__00000001_ in:\n%s", merged)
	}
	// From file 2, A0 should become A2
	if !strings.Contains(merged, "__A__00000002_ := (&models.A{Name: \"A2\"}).Stage(stage)") {
		t.Errorf("missing remapped declaration of __A__00000002_ in:\n%s", merged)
	}
	// From file 2, B0 should remain B0 (first B seen)
	if !strings.Contains(merged, "__B__00000000_ := (&models.B{Name: \"B0\"}).Stage(stage)") {
		t.Errorf("missing declaration of __B__00000000_ in:\n%s", merged)
	}

	// Check pointer setup
	if !strings.Contains(merged, "__A__00000002_.B = __B__00000000_") {
		t.Errorf("missing remapped pointer assignment in:\n%s", merged)
	}
}

func TestMergeStageBytes_PointersAndSlices(t *testing.T) {
	file1 := `package main

import "github.com/fullstack-lang/gong/test/test3/go/models"

func _(stage *models.Stage) {
	__A__00000000_ := (&models.A{Name: "A0"}).Stage(stage)
	__B__00000000_ := (&models.B{Name: "B0"}).Stage(stage)
	__A__00000000_.B = __B__00000000_
	__A__00000000_.Bs = append(__A__00000000_.Bs, __B__00000000_)
}
`

	file2 := `package main

import "github.com/fullstack-lang/gong/test/test3/go/models"

func _(stage *models.Stage) {
	__A__00000000_ := (&models.A{Name: "A1"}).Stage(stage)
	__B__00000000_ := (&models.B{Name: "B1"}).Stage(stage)
	__B__00000001_ := (&models.B{Name: "B2"}).Stage(stage)
	__A__00000000_.B = __B__00000001_
	__A__00000000_.Bs = append(__A__00000000_.Bs, __B__00000000_)
	__A__00000000_.Bs = slices.Insert(__A__00000000_.Bs, 0, __B__00000001_)
}
`

	merged, err := MergeStageBytesToString([][]byte{[]byte(file1), []byte(file2)}, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	fset := token.NewFileSet()
	_, err = parser.ParseFile(fset, "merged.go", merged, parser.ParseComments)
	if err != nil {
		t.Fatalf("merged file syntax error:\n%s\nError: %v", merged, err)
	}

	// File 1 has A0, B0 -> offset for A is 1, offset for B is 1
	// File 2's A0 -> A1, B0 -> B1, B1 -> B2
	expectedPointers := []string{
		"__A__00000000_.B = __B__00000000_",
		"__A__00000000_.Bs = append(__A__00000000_.Bs, __B__00000000_)",
		"__A__00000001_.B = __B__00000002_",
		"__A__00000001_.Bs = append(__A__00000001_.Bs, __B__00000001_)",
		"__A__00000001_.Bs = slices.Insert(__A__00000001_.Bs, 0, __B__00000002_)",
	}

	for _, expected := range expectedPointers {
		if !strings.Contains(merged, expected) {
			t.Errorf("missing expected pointer line '%s' in:\n%s", expected, merged)
		}
	}
}

func TestMergeStageBytes_ThreeFiles(t *testing.T) {
	file1 := `package main
import "github.com/fullstack-lang/gong/test/test3/go/models"
func _(stage *models.Stage) {
	__A__00000000_ := (&models.A{Name: "1"}).Stage(stage)
}
`
	file2 := `package main
import "github.com/fullstack-lang/gong/test/test3/go/models"
func _(stage *models.Stage) {
	__A__00000000_ := (&models.A{Name: "2"}).Stage(stage)
}
`
	file3 := `package main
import "github.com/fullstack-lang/gong/test/test3/go/models"
func _(stage *models.Stage) {
	__A__00000000_ := (&models.A{Name: "3"}).Stage(stage)
}
`

	merged, err := MergeStageBytesToString([][]byte{[]byte(file1), []byte(file2), []byte(file3)}, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	fset := token.NewFileSet()
	_, err = parser.ParseFile(fset, "merged.go", merged, parser.ParseComments)
	if err != nil {
		t.Fatalf("syntax error: %v\n%s", err, merged)
	}

	if !strings.Contains(merged, "__A__00000000_ :=") ||
		!strings.Contains(merged, "__A__00000001_ :=") ||
		!strings.Contains(merged, "__A__00000002_ :=") {
		t.Errorf("expected A0, A1, A2 declarations in:\n%s", merged)
	}
}

func TestMergeStageFiles_ActualWorkspaceFiles(t *testing.T) {
	file1 := filepath.Join("..", "..", "test", "test3", "go", "cmd", "test3", "data", "stage.go")
	file2 := filepath.Join("..", "..", "test", "test3", "go", "cmd", "test3", "data", "issue1128.go")

	tmpDir, err := os.MkdirTemp("", "gong_merge_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	outFile := filepath.Join(tmpDir, "merged.go")
	err = MergeStageFiles([]string{file1, file2}, outFile, "")
	if err != nil {
		t.Fatalf("error merging workspace files: %v", err)
	}

	content, err := os.ReadFile(outFile)
	if err != nil {
		t.Fatalf("error reading output: %v", err)
	}

	fset := token.NewFileSet()
	_, err = parser.ParseFile(fset, outFile, content, parser.ParseComments)
	if err != nil {
		t.Fatalf("merged file syntax error: %v\n%s", err, string(content))
	}

	// Verify unmarshalling with test3 models
	testStage := test3_models.NewStage("test")
	err = test3_models.ParseAstFile(testStage, outFile, false)
	if err != nil {
		t.Fatalf("failed to unmarshall merged stage with test3 models: %v", err)
	}

	if len(testStage.As) != 3 {
		t.Errorf("expected 3 instances of A in unmarshalled stage, got %d", len(testStage.As))
	}
	if len(testStage.Bs) != 1 {
		t.Errorf("expected 1 instance of B in unmarshalled stage, got %d", len(testStage.Bs))
	}
}
