package models

import (
	"testing"

	gong "github.com/fullstack-lang/gong/go/models"
)

func TestEnforceSemanticMultiplicity(t *testing.T) {
	docStage := NewStage("test_doc")
	gongStage := gong.NewStage("test_gong")

	// Gong struct A with field Bs of type *B (PointerToGongStructField)
	structB := (&gong.GongStruct{Name: "B"}).Stage(gongStage)
	structA := (&gong.GongStruct{Name: "A"}).Stage(gongStage)
	fieldBs := (&gong.PointerToGongStructField{
		Name:       "Bs",
		GongStruct: structB,
	}).Stage(gongStage)
	structA.Fields = append(structA.Fields, fieldBs)
	structA.PointerToGongStructFields = append(structA.PointerToGongStructFields, fieldBs)

	// In doc stage, GongStructShape A and LinkShape Bs with TargetMultiplicity = MANY (outdated)
	classdiagram := (&Classdiagram{Name: "Default"}).Stage(docStage)
	gongStructShapeA := (&GongStructShape{
		Name:           "Default-A",
		IdentifierMeta: "ref_models.A{}",
	}).Stage(docStage)
	gongStructShapeB := (&GongStructShape{
		Name:           "Default-B",
		IdentifierMeta: "ref_models.B{}",
	}).Stage(docStage)

	linkShapeBs := (&LinkShape{
		Name:                    "Bs",
		IdentifierMeta:          "ref_models.A{}.Bs",
		FieldTypeIdentifierMeta: "ref_models.B{}",
		TargetMultiplicity:      MANY, // outdated multiplicity
		SourceMultiplicity:      MANY,
	}).Stage(docStage)

	gongStructShapeA.LinkShapes = append(gongStructShapeA.LinkShapes, linkShapeBs)
	classdiagram.GongStructShapes = append(classdiagram.GongStructShapes, gongStructShapeA, gongStructShapeB)

	stager := &Stager{
		stage:     docStage,
		gongStage: gongStage,
	}

	stager.enforceSemantic()

	if linkShapeBs.TargetMultiplicity != ZERO_ONE {
		t.Fatalf("expected TargetMultiplicity to be ZERO_ONE (0..1), got %v", linkShapeBs.TargetMultiplicity)
	}

	// Now change field in gongStage to SliceOfPointerToGongStructField
	sliceFieldBs := (&gong.SliceOfPointerToGongStructField{
		Name:       "Bs",
		GongStruct: structB,
	}).Stage(gongStage)
	structA.Fields[0] = sliceFieldBs
	structA.PointerToGongStructFields = nil
	structA.SliceOfPointerToGongStructFields = append(structA.SliceOfPointerToGongStructFields, sliceFieldBs)

	stager.enforceSemantic()

	if linkShapeBs.TargetMultiplicity != MANY {
		t.Fatalf("expected TargetMultiplicity to be MANY (*), got %v", linkShapeBs.TargetMultiplicity)
	}

	// Test AttributeShape field type change
	basicField := (&gong.GongBasicField{
		Name:         "Name",
		DeclaredType: "string",
	}).Stage(gongStage)
	structA.Fields = append(structA.Fields, basicField)
	structA.GongBasicFields = append(structA.GongBasicFields, basicField)

	attrShape := (&AttributeShape{
		Name:           "Name",
		IdentifierMeta: "ref_models.A{}.Name",
		Fieldtypename:  "int", // outdated type
	}).Stage(docStage)
	gongStructShapeA.AttributeShapes = append(gongStructShapeA.AttributeShapes, attrShape)

	stager.enforceSemantic()

	if attrShape.Fieldtypename != "string" {
		t.Fatalf("expected Fieldtypename to be 'string', got %v", attrShape.Fieldtypename)
	}

	// Change basicField type to "int"
	basicField.DeclaredType = "int"
	stager.enforceSemantic()
	if attrShape.Fieldtypename != "int" {
		t.Fatalf("expected Fieldtypename to be 'int', got %v", attrShape.Fieldtypename)
	}
}

func TestEnforceSemanticMultiPackage(t *testing.T) {
	docStage := NewStage("test_doc_mp")
	docStage.MetaPackageImports = []*MetaPackageImport{
		{Alias: "ref_models", Path: `"github.com/fullstack-lang/gong/test/test2/go/models"`},
		{Alias: "ref_x", Path: `"github.com/fullstack-lang/gong/test/test2/go/models/x"`},
		{Alias: "ref_y", Path: `"github.com/fullstack-lang/gong/test/test2/go/models/y"`},
	}

	gongStage := gong.NewStage("test_gong_mp")

	modelPkgModels := &gong.ModelPkg{PkgGoName: "models", PkgPath: "github.com/fullstack-lang/gong/test/test2/go/models"}
	modelPkgX := &gong.ModelPkg{PkgGoName: "x", PkgPath: "github.com/fullstack-lang/gong/test/test2/go/models/x"}
	modelPkgY := &gong.ModelPkg{PkgGoName: "y", PkgPath: "github.com/fullstack-lang/gong/test/test2/go/models/y"}

	structY := (&gong.GongStruct{Name: "Y", ModelPkg: modelPkgY}).Stage(gongStage)
	structX := (&gong.GongStruct{Name: "X", ModelPkg: modelPkgX}).Stage(gongStage)
	structA := (&gong.GongStruct{Name: "A", ModelPkg: modelPkgModels}).Stage(gongStage)

	fieldY := (&gong.PointerToGongStructField{
		Name:       "Y",
		GongStruct: structY,
	}).Stage(gongStage)
	structX.Fields = append(structX.Fields, fieldY)
	structX.PointerToGongStructFields = append(structX.PointerToGongStructFields, fieldY)

	fieldX := (&gong.PointerToGongStructField{
		Name:       "X",
		GongStruct: structX,
	}).Stage(gongStage)
	structA.Fields = append(structA.Fields, fieldX)
	structA.PointerToGongStructFields = append(structA.PointerToGongStructFields, fieldX)

	classdiagram := (&Classdiagram{Name: "StageSet_Diagram"}).Stage(docStage)
	gongStructShapeA := (&GongStructShape{
		Name:           "StageSet_Diagram-A",
		IdentifierMeta: "ref_models.A{}",
	}).Stage(docStage)
	gongStructShapeX := (&GongStructShape{
		Name:           "StageSet_Diagram-X",
		IdentifierMeta: "ref_x.X{}",
	}).Stage(docStage)
	gongStructShapeY := (&GongStructShape{
		Name:           "StageSet_Diagram-Y",
		IdentifierMeta: "ref_y.Y{}",
	}).Stage(docStage)

	linkShapeX := (&LinkShape{
		Name:                    "X",
		IdentifierMeta:          "ref_models.A{}.X",
		FieldTypeIdentifierMeta: "ref_x.X{}",
		TargetMultiplicity:      ZERO_ONE,
		SourceMultiplicity:      MANY,
	}).Stage(docStage)

	linkShapeY := (&LinkShape{
		Name:                    "Y",
		IdentifierMeta:          "ref_x.X{}.Y",
		FieldTypeIdentifierMeta: "ref_y.Y{}",
		TargetMultiplicity:      ZERO_ONE,
		SourceMultiplicity:      MANY,
	}).Stage(docStage)

	gongStructShapeA.LinkShapes = append(gongStructShapeA.LinkShapes, linkShapeX)
	gongStructShapeX.LinkShapes = append(gongStructShapeX.LinkShapes, linkShapeY)
	classdiagram.GongStructShapes = append(classdiagram.GongStructShapes, gongStructShapeA, gongStructShapeX, gongStructShapeY)

	stager := &Stager{
		stage:     docStage,
		gongStage: gongStage,
	}

	stager.enforceSemantic()

	if linkShapeX.FieldTypeIdentifierMeta != "ref_x.X{}" {
		t.Errorf("expected FieldTypeIdentifierMeta ref_x.X{}, got %v", linkShapeX.FieldTypeIdentifierMeta)
	}
	if linkShapeY.FieldTypeIdentifierMeta != "ref_y.Y{}" {
		t.Errorf("expected FieldTypeIdentifierMeta ref_y.Y{}, got %v", linkShapeY.FieldTypeIdentifierMeta)
	}

	// Also test recovery if FieldTypeIdentifierMeta was corrupted/outdated to ref_models
	linkShapeX.FieldTypeIdentifierMeta = "ref_models.X{}"
	linkShapeY.FieldTypeIdentifierMeta = "ref_models.Y{}"
	stager.enforceSemantic()

	if linkShapeX.FieldTypeIdentifierMeta != "ref_x.X{}" {
		t.Errorf("expected recovered FieldTypeIdentifierMeta ref_x.X{}, got %v", linkShapeX.FieldTypeIdentifierMeta)
	}
	if linkShapeY.FieldTypeIdentifierMeta != "ref_y.Y{}" {
		t.Errorf("expected recovered FieldTypeIdentifierMeta ref_y.Y{}, got %v", linkShapeY.FieldTypeIdentifierMeta)
	}
}
