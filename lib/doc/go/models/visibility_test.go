package models

import (
	"testing"

	gong "github.com/fullstack-lang/gong/go/models"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
)

func TestShapeVisibilityInterface(t *testing.T) {
	// Verify that all 7 shapes implement ShapeWithVisibility
	shapes := []ShapeWithVisibility{
		&GongStructShape{},
		&LinkShape{},
		&AttributeShape{},
		&GongEnumShape{},
		&GongEnumValueShape{},
		&GongNoteShape{},
		&GongNoteLinkShape{},
	}

	for i, s := range shapes {
		if s.GetIsHidden() {
			t.Errorf("shape %d initially hidden", i)
		}
		s.SetIsHidden(true)
		if !s.GetIsHidden() {
			t.Errorf("shape %d was not hidden after SetIsHidden(true)", i)
		}
		s.SetIsHidden(false)
		if s.GetIsHidden() {
			t.Errorf("shape %d was not unhidden after SetIsHidden(false)", i)
		}
	}
}

func TestTreeVisibilityButton(t *testing.T) {
	docStage := NewStage("test_doc_vis_tree")
	treeStage := tree.NewStage("test_tree_vis")

	stager := &Stager{
		stage:     docStage,
		treeStage: treeStage,
	}

	shape := &GongStructShape{}
	node := &tree.Node{Name: "TestNode"}

	stager.addVisibilityButton(node, shape)
	if len(node.Buttons) != 1 {
		t.Fatalf("expected 1 button, got %d", len(node.Buttons))
	}
	btn := node.Buttons[0]
	if btn.Name != "Hide" || btn.Icon != string(buttons.BUTTON_visibility_off) {
		t.Errorf("unexpected button config: Name=%s, Icon=%s", btn.Name, btn.Icon)
	}

	shape.SetIsHidden(true)
	node2 := &tree.Node{Name: "TestNode2"}
	stager.addVisibilityButton(node2, shape)
	if len(node2.Buttons) != 1 {
		t.Fatalf("expected 1 button, got %d", len(node2.Buttons))
	}
	btn2 := node2.Buttons[0]
	if btn2.Name != "Show" || btn2.Icon != string(buttons.BUTTON_visibility) {
		t.Errorf("unexpected button config for hidden shape: Name=%s, Icon=%s", btn2.Name, btn2.Icon)
	}
}

func TestSvgHiding(t *testing.T) {
	docStage := NewStage("test_doc_svg")
	svgStage := svg.NewStage("test_svg")
	gongStage := gong.NewStage("test_gong_svg")

	structB := (&gong.GongStruct{Name: "B"}).Stage(gongStage)
	structA := (&gong.GongStruct{Name: "A"}).Stage(gongStage)
	fieldBs := (&gong.PointerToGongStructField{
		Name:       "Bs",
		GongStruct: structB,
	}).Stage(gongStage)
	structA.Fields = append(structA.Fields, fieldBs)
	structA.PointerToGongStructFields = append(structA.PointerToGongStructFields, fieldBs)

	basicField := (&gong.GongBasicField{
		Name:         "Name",
		DeclaredType: "string",
	}).Stage(gongStage)
	structA.Fields = append(structA.Fields, basicField)
	structA.GongBasicFields = append(structA.GongBasicFields, basicField)

	classdiagram := (&Classdiagram{Name: "Diagram1"}).Stage(docStage)
	diagramPkg := (&DiagramPackage{
		Name:                "Pkg",
		SelectedClassdiagram: classdiagram,
		Classdiagrams:       []*Classdiagram{classdiagram},
	}).Stage(docStage)
	_ = diagramPkg

	gongStructShapeA := (&GongStructShape{
		Name:           "Diagram1-A",
		IdentifierMeta: "ref_models.A{}",
	}).Stage(docStage)
	gongStructShapeB := (&GongStructShape{
		Name:           "Diagram1-B",
		IdentifierMeta: "ref_models.B{}",
	}).Stage(docStage)

	attrShape := (&AttributeShape{
		Name:           "Name",
		IdentifierMeta: "ref_models.A{}.Name",
		Fieldtypename:  "string",
	}).Stage(docStage)
	gongStructShapeA.AttributeShapes = append(gongStructShapeA.AttributeShapes, attrShape)

	linkShapeBs := (&LinkShape{
		Name:                    "Bs",
		IdentifierMeta:          "ref_models.A{}.Bs",
		FieldTypeIdentifierMeta: "ref_models.B{}",
		TargetMultiplicity:      ZERO_ONE,
		SourceMultiplicity:      MANY,
	}).Stage(docStage)
	gongStructShapeA.LinkShapes = append(gongStructShapeA.LinkShapes, linkShapeBs)

	classdiagram.GongStructShapes = append(classdiagram.GongStructShapes, gongStructShapeA, gongStructShapeB)

	treeNavStage := tree.NewStage("test_nav")
	stager := &Stager{
		stage:               docStage,
		svgStage:            svgStage,
		gongStage:           gongStage,
		treeNavigationStage: treeNavStage,
	}

	// 1. Initial render: everything visible
	stager.Svg()
	svgInst := svgStage.GetInstancesSorted[*svg.SVG]()[0]
	if len(svgInst.Layers) < 3 { // LayerA, LayerB, LinkLayer
		t.Fatalf("expected at least 3 layers, got %d", len(svgInst.Layers))
	}

	// 2. Hide link: link layer should be skipped
	linkShapeBs.IsHidden = true
	stager.Svg()
	svgInst = svgStage.GetInstancesSorted[*svg.SVG]()[0]
	foundLinkLayer := false
	for _, layer := range svgInst.Layers {
		if len(layer.Links) > 0 {
			foundLinkLayer = true
		}
	}
	if foundLinkLayer {
		t.Errorf("link was hidden, but found a link in svg layers")
	}

	// 3. Hide struct A: LayerA should be skipped
	linkShapeBs.IsHidden = false
	gongStructShapeA.IsHidden = true
	stager.Svg()
	svgInst = svgStage.GetInstancesSorted[*svg.SVG]()[0]
	foundLayerA := false
	for _, layer := range svgInst.Layers {
		if layer.Name == "LayerA" {
			foundLayerA = true
		}
	}
	if foundLayerA {
		t.Errorf("struct A was hidden, but found LayerA in svg layers")
	}

	// 4. Attribute hidden: should not appear in rect anchored texts
	gongStructShapeA.IsHidden = false
	attrShape.IsHidden = true
	stager.Svg()
	rectA := stager.map_GongstructShape_Rect[gongStructShapeA]
	for _, text := range rectA.RectAnchoredTexts {
		if text.Name == "Name : string" {
			t.Errorf("attribute Name was hidden, but text was rendered")
		}
	}

	// 5. Enum and EnumValue hiding
	gongEnum := (&gong.GongEnum{
		Name: "MyEnum",
		GongEnumValues: []*gong.GongEnumValue{
			{Name: "VAL1"},
			{Name: "VAL2"},
		},
	}).Stage(gongStage)

	enumShape := (&GongEnumShape{
		Name:           "Diagram1-MyEnum",
		IdentifierMeta: "new(models.MyEnum)",
	}).Stage(docStage)
	valShape1 := (&GongEnumValueShape{
		Name:           "VAL1",
		IdentifierMeta: "models.VAL1",
	}).Stage(docStage)
	valShape2 := (&GongEnumValueShape{
		Name:           "VAL2",
		IdentifierMeta: "models.VAL2",
	}).Stage(docStage)
	enumShape.GongEnumValueShapes = append(enumShape.GongEnumValueShapes, valShape1, valShape2)
	classdiagram.GongEnumShapes = append(classdiagram.GongEnumShapes, enumShape)

	stager.Svg()
	rectEnum := stager.map_GongenumShape_Rect[enumShape]
	if rectEnum == nil {
		t.Fatalf("expected rect for enumShape")
	}
	if len(rectEnum.RectAnchoredTexts) != 3 { // title + 2 values
		t.Errorf("expected 3 texts for enum, got %d", len(rectEnum.RectAnchoredTexts))
	}

	valShape1.IsHidden = true
	stager.Svg()
	rectEnum = stager.map_GongenumShape_Rect[enumShape]
	if len(rectEnum.RectAnchoredTexts) != 2 { // title + 1 value
		t.Errorf("expected 2 texts for enum after hiding VAL1, got %d", len(rectEnum.RectAnchoredTexts))
	}

	enumShape.IsHidden = true
	stager.Svg()
	if stager.map_GongenumShape_Rect[enumShape] != nil {
		t.Errorf("expected enum to be omitted when hidden")
	}

	// 6. Note and NoteLink hiding
	_ = gongEnum
	gongNote := (&gong.GongNote{
		Name: "MyNote",
		Body: "MyNote Title\nNote body text",
		Links: []*gong.GongLink{
			{Name: "A"},
		},
	}).Stage(gongStage)

	noteShape := (&GongNoteShape{
		Name:       "Diagram1-MyNote",
		Identifier: "models.MyNote",
		Body:       gongNote.Body,
	}).Stage(docStage)
	noteLinkShape := (&GongNoteLinkShape{
		Name:       "Diagram1-MyNote-A",
		Identifier: "models.A",
		Type:       NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE,
	}).Stage(docStage)
	noteShape.GongNoteLinkShapes = append(noteShape.GongNoteLinkShapes, noteLinkShape)
	classdiagram.GongNoteShapes = append(classdiagram.GongNoteShapes, noteShape)

	stager.Svg()
	if stager.map_NoteShape_Rect[noteShape] == nil {
		t.Fatalf("expected rect for noteShape")
	}
	svgInst = svgStage.GetInstancesSorted[*svg.SVG]()[0]
	hasNoteLink := false
	for _, layer := range svgInst.Layers {
		for _, l := range layer.Links {
			if l.Name == "models.MyNote - to - A" {
				hasNoteLink = true
			}
		}
	}
	if !hasNoteLink {
		t.Errorf("expected note link in svg layers")
	}

	noteLinkShape.IsHidden = true
	stager.Svg()
	svgInst = svgStage.GetInstancesSorted[*svg.SVG]()[0]
	hasNoteLink = false
	for _, layer := range svgInst.Layers {
		for _, l := range layer.Links {
			if l.Name == "models.MyNote - to - A" {
				hasNoteLink = true
			}
		}
	}
	if hasNoteLink {
		t.Errorf("expected note link to be hidden")
	}

	noteShape.IsHidden = true
	stager.Svg()
	if stager.map_NoteShape_Rect[noteShape] != nil {
		t.Errorf("expected note to be omitted when hidden")
	}
}
