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
