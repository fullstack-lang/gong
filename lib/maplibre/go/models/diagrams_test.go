package models_test

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/diagramlayout"
	gong_models "github.com/fullstack-lang/gong/go/models"
	doc_models "github.com/fullstack-lang/gong/lib/doc/go/models"
	"github.com/fullstack-lang/gong/lib/doc/go/prepare"
	embeddedgo "github.com/fullstack-lang/gong/lib/maplibre/go"
	split "github.com/fullstack-lang/gong/lib/split/go/models"
)

const (
	maplibreGoPrefix     = "github.com/fullstack-lang/gong/lib/maplibre/go"
	maplibreModelsImport = "github.com/fullstack-lang/gong/lib/maplibre/go/models"
	maplibreDiagramTitle = "MapLibre 2D/3D Map Scene Graph"
	maplibreNoteName     = "NoteOnMaplibre"
)

func TestMaplibreA3UMLDiagram(t *testing.T) {
	metaImports := []*doc_models.MetaPackageImport{
		{Alias: "ref_models", Path: fmt.Sprintf("%q", maplibreModelsImport)},
	}

	if os.Getenv("GONG_REGEN_DIAGRAMS") == "1" {
		regenerateMaplibreA3Diagram(t, metaImports)
		return
	}

	const diagramsPath = "diagrams/diagrams_set.go"
	beforeBytes, err := os.ReadFile(diagramsPath)
	if err != nil {
		t.Fatalf("failed reading %s: %v", diagramsPath, err)
	}

	rawStage := doc_models.NewStage(maplibreModelsImport + ":raw")
	rawStage.MetaPackageImports = metaImports
	if err := rawStage.ParseAstFile(diagramsPath, false); err != nil {
		t.Fatalf("failed parsing %s: %v", diagramsPath, err)
	}

	mux := http.NewServeMux()
	diagramEditor := &split.AsSplitArea{Name: "Bottom", Size: 50}
	docStager := prepare.Prepare(
		mux,
		true,
		maplibreGoPrefix+":verify",
		embeddedgo.GoModelsDir,
		embeddedgo.GoDiagramsDir,
		diagramEditor,
		map[string]int{},
	)
	verifyStage := docStager.GetStage()

	afterBytes, err := os.ReadFile(diagramsPath)
	if err != nil {
		t.Fatalf("failed re-reading %s: %v", diagramsPath, err)
	}
	if string(beforeBytes) != string(afterBytes) {
		t.Fatalf("%s was unexpectedly modified during read-only test", diagramsPath)
	}

	assertMaplibreDiagramStageCounts(t, verifyStage)

	diagrams := rawStage.GetInstancesSorted[*doc_models.Classdiagram]()
	if len(diagrams) != 1 {
		t.Fatalf("expected 1 Classdiagram, got %d", len(diagrams))
	}
	assertMaplibreA3GeometryClean(t, diagrams[0])
}

func regenerateMaplibreA3Diagram(t *testing.T, metaImports []*doc_models.MetaPackageImport) {
	t.Helper()

	mux := http.NewServeMux()
	diagramEditor := &split.AsSplitArea{Name: "Bottom", Size: 50}
	docStager := prepare.Prepare(
		mux,
		true,
		maplibreGoPrefix+":regen",
		embeddedgo.GoModelsDir,
		embeddedgo.GoDiagramsDir,
		diagramEditor,
		map[string]int{},
	)
	docStage := docStager.GetStage()
	gongStage := docStager.GetGongStage()

	structsByName := gongStage.GetInstancesMapByName[*gong_models.GongStruct]()
	enumsByName := gongStage.GetInstancesMapByName[*gong_models.GongEnum]()
	notesByName := gongStage.GetInstancesMapByName[*gong_models.GongNote]()

	docStage.Reset()
	docStage.MetaPackageImportAlias = "ref_models"
	docStage.MetaPackageImportPath = fmt.Sprintf("%q", maplibreModelsImport)
	docStage.MetaPackageImports = metaImports

	pkg := (&doc_models.DiagramPackage{
		Name: "MapLibre 2D/3D Map UML Diagrams",
	}).Stage(docStage)

	cd := (&doc_models.Classdiagram{
		Name:                      maplibreDiagramTitle,
		Description:               "A3 Landscape overview of the MapLibre 2D/3D vector map scene graph: MapView, Layer, extruded Polygon, Polyline, SVG Marker, Coordinate, and alignment/anchor enums.",
		ShowNbInstances:           true,
		ShowMultiplicity:          true,
		ShowLinkNames:             true,
		IsExpanded:                true,
		NodeGongStructsIsExpanded: false,
		NodeGongEnumsIsExpanded:   false,
		NodeGongNotesIsExpanded:   true,
		NodeGongNoteNodeExpansion: `[true]`,
	}).Stage(docStage)
	pkg.Classdiagrams = []*doc_models.Classdiagram{cd}
	pkg.SelectedClassdiagram = cd

	gongNote, ok := notesByName[maplibreNoteName]
	if !ok {
		t.Fatalf("missing GONGNOTE %s in notes.go", maplibreNoteName)
	}
	noteW, noteH := diagramlayout.NoteBoxSize(gongNote.Body, 536)
	noteShape := (&doc_models.GongNoteShape{
		Name:       cd.Name + "-" + gongNote.Name,
		Identifier: doc_models.GongNoteNameToIdentifier(gongNote.Name),
		Body:       gongNote.Body,
		BodyHTML:   gongNote.BodyHTML,
		X:          56,
		Y:          56,
		Width:      noteW,
		Height:     noteH,
	}).Stage(docStage)
	noteLink := (&doc_models.GongNoteLinkShape{
		Name:       cd.Name + "-" + noteShape.Name + "-MapView",
		Identifier: doc_models.GongStructNameToIdentifier("MapView"),
		Type:       doc_models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE,
		IsHidden:   false,
	}).Stage(docStage)
	noteShape.GongNoteLinkShapes = []*doc_models.GongNoteLinkShape{noteLink}
	cd.GongNoteShapes = []*doc_models.GongNoteShape{noteShape}

	buildStruct := func(name string, x, y, minW, minH float64) *doc_models.GongStructShape {
		gs, exists := structsByName[name]
		if !exists {
			t.Fatalf("missing GongStruct %s", name)
		}
		var attrShapes []*doc_models.AttributeShape
		var attrLines []string
		for _, field := range gs.Fields {
			if bf, ok := field.(*gong_models.GongBasicField); ok {
				parts := strings.Split(bf.DeclaredType, ".")
				ftName := parts[len(parts)-1]
				as := (&doc_models.AttributeShape{
					Name:           bf.Name,
					IdentifierMeta: fmt.Sprintf("ref_models.%s{}.%s", gs.Name, bf.Name),
					Structname:     gs.Name,
					Fieldtypename:  ftName,
					IsHidden:       false,
				}).Stage(docStage)
				attrShapes = append(attrShapes, as)
				attrLines = append(attrLines, bf.Name+" : "+ftName)
			}
		}
		w := math.Max(minW, diagramlayout.StructBoxMinWidth(gs.Name, attrLines, true))
		h := math.Max(minH, 54.0+float64(len(attrShapes))*20.0)
		gss := (&doc_models.GongStructShape{
			Name:            cd.Name + "-" + gs.Name,
			IdentifierMeta:  fmt.Sprintf("ref_models.%s{}", gs.Name),
			AttributeShapes: attrShapes,
			X:               x,
			Y:               y,
			Width:           w,
			Height:          h,
		}).Stage(docStage)
		cd.GongStructShapes = append(cd.GongStructShapes, gss)
		return gss
	}

	buildEnum := func(name string, x, y, minW, minH float64) *doc_models.GongEnumShape {
		ge, exists := enumsByName[name]
		if !exists {
			t.Fatalf("missing GongEnum %s", name)
		}
		var valShapes []*doc_models.GongEnumValueShape
		var valNames []string
		for _, ev := range ge.GongEnumValues {
			evs := (&doc_models.GongEnumValueShape{
				Name:           ev.Name,
				IdentifierMeta: fmt.Sprintf("ref_models.%s", ev.Name),
				IsHidden:       false,
			}).Stage(docStage)
			valShapes = append(valShapes, evs)
			valNames = append(valNames, ev.Name)
		}
		ew, eh := diagramlayout.EnumBoxSize(ge.Name, valNames)
		ges := (&doc_models.GongEnumShape{
			Name:                cd.Name + "-" + ge.Name,
			IdentifierMeta:      fmt.Sprintf("new(ref_models.%s)", ge.Name),
			GongEnumValueShapes: valShapes,
			X:                   x,
			Y:                   y,
			Width:               math.Max(minW, ew),
			Height:              math.Max(minH, eh),
		}).Stage(docStage)
		cd.GongEnumShapes = append(cd.GongEnumShapes, ges)
		return ges
	}

	// Horizontal spine at Y = 550 aligns MapView -> Layer -> Polyline cleanly left-to-right.
	const spineY = 550.0
	mapViewShape := buildStruct("MapView", 56, spineY-100, 248, 200)
	layerShape := buildStruct("Layer", 424, spineY-84, 216, 168)

	polygonShape := buildStruct("Polygon", 796, 56, 292, 278)
	polylineShape := buildStruct("Polyline", 796, spineY-100, 292, 200)
	markerShape := buildStruct("Marker", 796, 726, 412, 340)

	// Column 4 right-aligned at X = 1264..1564 (Width = 300)
	const col4X = 1264.0
	const col4W = 300.0
	coordCenterY := 0.5 * ((polygonShape.Y + 0.5*polygonShape.Height) + (polylineShape.Y + 0.5*polylineShape.Height))
	coordinateShape := buildStruct("Coordinate", col4X, math.Round(coordCenterY-72), col4W, 144)

	buildEnum("MarkerPitchAlignmentEnum", col4X, 694, col4W, 116)
	anchorEnum := buildEnum("MarkerAnchorEnum", col4X, 830, col4W, 236)
	markerShape.Y = (anchorEnum.Y + anchorEnum.Height) - markerShape.Height

	addHorizLink := func(src, dst *doc_models.GongStructShape, fieldName string, startRatio, endRatio, corridorX float64) {
		_, srcName := doc_models.IdentifierMetaToPackageAndGongStructName(src.IdentifierMeta)
		_, dstName := doc_models.IdentifierMetaToPackageAndGongStructName(dst.IdentifierMeta)
		cornerRatio := (corridorX - src.X) / src.Width
		ls := (&doc_models.LinkShape{
			Name:                    fieldName,
			IdentifierMeta:          fmt.Sprintf("ref_models.%s{}.%s", srcName, fieldName),
			FieldTypeIdentifierMeta: fmt.Sprintf("ref_models.%s{}", dstName),
			SourceMultiplicity:      doc_models.MANY,
			TargetMultiplicity:      doc_models.MANY,
			StartOrientation:        doc_models.ORIENTATION_HORIZONTAL,
			StartRatio:              math.Round(startRatio*10000) / 10000,
			EndOrientation:          doc_models.ORIENTATION_HORIZONTAL,
			EndRatio:                math.Round(endRatio*10000) / 10000,
			CornerOffsetRatio:       math.Round(cornerRatio*10000) / 10000,
			X:                       math.Round((src.X+dst.X)*0.5*10) / 10,
			Y:                       math.Round((src.Y+dst.Y)*0.5*10) / 10,
			IsHidden:                false,
		}).Stage(docStage)
		src.LinkShapes = append(src.LinkShapes, ls)
	}

	addHorizLink(mapViewShape, layerShape, "Layers", 0.50, 0.50, 360)
	addHorizLink(layerShape, polygonShape, "Polygons", 0.24, 0.50, 672)
	addHorizLink(layerShape, polylineShape, "Polylines", 0.50, 0.50, 710)
	addHorizLink(layerShape, markerShape, "Markers", 0.76, 0.25, 672)
	addHorizLink(polygonShape, coordinateShape, "Coordinates", 0.65, 0.28, 1122)
	addHorizLink(polylineShape, coordinateShape, "Coordinates", 0.35, 0.72, 1122)

	assertMaplibreA3GeometryClean(t, cd)

	docStage.ComputeReferenceAndOrders()
	const outPath = "diagrams/diagrams_set.go"
	outFile, err := os.Create(outPath)
	if err != nil {
		t.Fatalf("failed creating %s: %v", outPath, err)
	}
	docStage.Marshall(outFile, "github.com/fullstack-lang/gong/lib/doc/go/models", "diagrams")
	if err := outFile.Close(); err != nil {
		t.Fatalf("failed closing %s: %v", outPath, err)
	}

	verifyMux := http.NewServeMux()
	verifyEditor := &split.AsSplitArea{Name: "VerifyDisk", Size: 50}
	verifyStager := prepare.Prepare(
		verifyMux,
		false,
		maplibreGoPrefix+":verify_disk",
		embeddedgo.GoModelsDir,
		embeddedgo.GoDiagramsDir,
		verifyEditor,
		map[string]int{},
	)
	assertMaplibreDiagramStageCounts(t, verifyStager.GetStage())
}

func assertMaplibreDiagramStageCounts(t *testing.T, stage *doc_models.Stage) {
	t.Helper()
	if got := len(*stage.GetInstancesSet[*doc_models.Classdiagram]()); got != 1 {
		t.Errorf("expected 1 Classdiagram, got %d", got)
	}
	if got := len(*stage.GetInstancesSet[*doc_models.GongStructShape]()); got != 6 {
		t.Errorf("expected 6 GongStructShapes, got %d", got)
	}
	if got := len(*stage.GetInstancesSet[*doc_models.AttributeShape]()); got != 45 {
		t.Errorf("expected 45 AttributeShapes (all basic/enum fields), got %d", got)
	}
	if got := len(*stage.GetInstancesSet[*doc_models.LinkShape]()); got != 6 {
		t.Errorf("expected 6 LinkShapes, got %d", got)
	}
	if got := len(*stage.GetInstancesSet[*doc_models.GongEnumShape]()); got != 2 {
		t.Errorf("expected 2 GongEnumShapes, got %d", got)
	}
	if got := len(*stage.GetInstancesSet[*doc_models.GongEnumValueShape]()); got != 12 {
		t.Errorf("expected 12 GongEnumValueShapes, got %d", got)
	}
	if got := len(*stage.GetInstancesSet[*doc_models.GongNoteShape]()); got != 1 {
		t.Errorf("expected 1 GongNoteShape, got %d", got)
	}
	if got := len(*stage.GetInstancesSet[*doc_models.GongNoteLinkShape]()); got != 1 {
		t.Errorf("expected 1 GongNoteLinkShape, got %d", got)
	}
}

func assertMaplibreA3GeometryClean(t *testing.T, cd *doc_models.Classdiagram) {
	t.Helper()
	dm := diagramlayout.AnalyzeClassdiagram(cd)
	t.Logf("[%s] %.0f x %.0f px (aspect ratio %.2f)", cd.Name, dm.Width, dm.Height, dm.AspectRatio)

	if len(dm.BoxOverlaps) > 0 {
		t.Errorf("box overlaps: %v", dm.BoxOverlaps)
	}
	if len(dm.TextClipIssues) > 0 {
		t.Errorf("text clip issues: %v", dm.TextClipIssues)
	}
	if len(dm.SegmentPiercings) > 0 {
		t.Errorf("segment piercings: %v", dm.SegmentPiercings)
	}
	if len(dm.LabelBoxHits) > 0 {
		t.Errorf("label-box collisions: %v", dm.LabelBoxHits)
	}
	if len(dm.LabelLabelHits) > 0 {
		t.Errorf("label-label collisions: %v", dm.LabelLabelHits)
	}
	if len(dm.NoteLinkHits) > 0 {
		t.Errorf("note-link collisions: %v", dm.NoteLinkHits)
	}
	if dm.Width < 1350 || dm.Width > 1650 {
		t.Errorf("expected A3 landscape width in [1350, 1650] px, got %.0f", dm.Width)
	}
	if dm.Height < 950 || dm.Height > 1140 {
		t.Errorf("expected A3 landscape height in [950, 1140] px, got %.0f", dm.Height)
	}
	if dm.AspectRatio < 1.32 || dm.AspectRatio > 1.56 {
		t.Errorf("expected A3 landscape aspect ratio near 1.414 (in [1.32, 1.56]), got %.2f", dm.AspectRatio)
	}
}
