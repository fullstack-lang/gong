package models_test

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	embeddedgo "github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/probe"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/diagramlayout"
	cesium "github.com/fullstack-lang/gong/lib/cesium/go/models"
	maplibre "github.com/fullstack-lang/gong/lib/maplibre/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

const (
	domainName         = "airspace"
	domainImportPrefix = "github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models"
)

func TestRoundTripAndPopulateStage(t *testing.T) {
	xmlPath := filepath.Join("..", "..", "sample_"+domainName+".xml")
	xmlBytes, err := os.ReadFile(xmlPath)
	if err != nil {
		t.Fatalf("failed reading %s: %v", xmlPath, err)
	}

	stage := models.NewStage(domainName)
	if err := models.UnmarshallXMLToStage(stage, xmlBytes); err != nil {
		t.Fatalf("UnmarshallXMLToStage failed: %v", err)
	}

	stage.ComputeReferenceAndOrders()
	goCode, err := stage.MarshallToString("", "main")
	if err != nil {
		t.Fatalf("MarshallToString failed: %v", err)
	}

	stageGoPath := filepath.Join("..", "cmd", domainName, "data", "stage.go")
	if err := os.MkdirAll(filepath.Dir(stageGoPath), 0755); err != nil {
		t.Fatalf("failed creating data dir: %v", err)
	}
	if err := os.WriteFile(stageGoPath, []byte(goCode), 0644); err != nil {
		t.Fatalf("failed writing %s: %v", stageGoPath, err)
	}

	stage2 := models.NewStage(domainName + "_ast")
	if err := stage2.ParseAstString(goCode, true); err != nil {
		t.Fatalf("ParseAstString failed: %v", err)
	}

	xmlOut, err := models.MarshallStageToXML(stage2)
	if err != nil {
		t.Fatalf("MarshallStageToXML failed: %v", err)
	}

	stage3 := models.NewStage(domainName + "_xml")
	if err := models.UnmarshallXMLToStage(stage3, xmlOut); err != nil {
		t.Fatalf("round-trip UnmarshallXMLToStage failed: %v", err)
	}
}

func TestUnmarshalAirspaceBaseline(t *testing.T) {
	baselinePath := filepath.Join("..", "cmd", domainName, "data", "Airspace.BASELINE")
	xmlBytes, err := os.ReadFile(baselinePath)
	if err != nil {
		if os.IsNotExist(err) {
			t.Skipf("skipping baseline test, file not found: %s", baselinePath)
		}
		t.Fatalf("failed reading %s: %v", baselinePath, err)
	}
	if len(xmlBytes) >= 1024*1024 {
		t.Fatalf("expected Airspace.BASELINE subset to be < 1 MB, got %d bytes", len(xmlBytes))
	}

	mux := http.NewServeMux()
	stage := models.NewStage(domainName + "_baseline")
	prb := probe.NewProbe(mux, embeddedgo.GoModelsDir, embeddedgo.GoDiagramsDir, true, stage)
	stage.SetProbeIF(prb)
	stager := models.NewStager(mux, stage, prb)

	stage.Reset()
	if err := models.UnmarshallXMLToStage(stage, xmlBytes); err != nil {
		t.Fatalf("UnmarshallXMLToStage on Airspace.BASELINE failed: %v", err)
	}
	stage.Commit()

	msgs := stage.GetInstancesSorted[*models.AIXMBasicMessage]()
	if len(msgs) != 1 {
		t.Fatalf("expected 1 AIXMBasicMessage, got %d", len(msgs))
	}

	checks := []struct {
		name  string
		count int
	}{
		{"AIXMBasicMessage", len(*stage.GetInstancesSet[*models.AIXMBasicMessage]())},
		{"MessageMetadataPropertyType", len(*stage.AixmabstractgmlobjecttypesStage.GetInstancesSet[*aixmabstractgmlobjecttypes.MessageMetadataPropertyType]())},
		{"BasicMessageMemberAIXMPropertyType", len(*stage.GetInstancesSet[*models.BasicMessageMemberAIXMPropertyType]())},
		{"AirspaceType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceType]())},
		{"AirspaceTimeSlicePropertyType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceTimeSlicePropertyType]())},
		{"AirspaceTimeSliceType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceTimeSliceType]())},
		{"TimePrimitivePropertyType", len(*stage.GmlStage.GetInstancesSet[*gml.TimePrimitivePropertyType]())},
		{"AirspaceGeometryComponentType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceGeometryComponentType]())},
		{"AirspaceVolumeType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceVolumeType]())},
		{"AirspaceVolumeDependencyType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceVolumeDependencyType]())},
		{"SurfaceType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SurfaceType]())},
		{"AirspaceActivationType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceActivationType]())},
		{"AirspaceLayerClassType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceLayerClassType]())},
		{"AirspaceLayerType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirspaceLayerType]())},
		{"TimesheetType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TimesheetType]())},
		{"AuthorityForAirspaceType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AuthorityForAirspaceType]())},
		{"AuthorityForAirspaceTimeSliceType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AuthorityForAirspaceTimeSliceType]())},
		{"GeoBorderType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.GeoBorderType]())},
		{"GeoBorderTimeSliceType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.GeoBorderTimeSliceType]())},
		{"CurveType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CurveType]())},
		{"ElevatedCurveType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedCurveType]())},
		{"NoteType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NoteType]())},
		{"DirectPositionListType", len(*stage.GmlStage.GetInstancesSet[*gml.DirectPositionListType]())},
		{"DirectPositionType", len(*stage.GmlStage.GetInstancesSet[*gml.DirectPositionType]())},
		{"SurfacePatchArrayPropertyType", len(*stage.GmlStage.GetInstancesSet[*gml.SurfacePatchArrayPropertyType]())},
		{"AbstractRingPropertyType", len(*stage.GmlStage.GetInstancesSet[*gml.AbstractRingPropertyType]())},
		{"LinearRingType", len(*stage.GmlStage.GetInstancesSet[*gml.LinearRingType]())},
		{"CurveSegmentArrayPropertyType", len(*stage.GmlStage.GetInstancesSet[*gml.CurveSegmentArrayPropertyType]())},
		{"PolygonPatchType", len(*stage.GmlStage.GetInstancesSet[*gml.PolygonPatchType]())},
		{"GeodesicStringType", len(*stage.GmlStage.GetInstancesSet[*gml.GeodesicStringType]())},
		{"ArcByCenterPointType", len(*stage.GmlStage.GetInstancesSet[*gml.ArcByCenterPointType]())},
		{"CircleByCenterPointType", len(*stage.GmlStage.GetInstancesSet[*gml.CircleByCenterPointType]())},
		{"Maplibre.MapView", len(*stager.GetMaplibreStage().GetInstancesSet[*maplibre.MapView]())},
		{"Maplibre.Layer", len(*stager.GetMaplibreStage().GetInstancesSet[*maplibre.Layer]())},
		{"Maplibre.Polygon", len(*stager.GetMaplibreStage().GetInstancesSet[*maplibre.Polygon]())},
		{"Maplibre.Polyline", len(*stager.GetMaplibreStage().GetInstancesSet[*maplibre.Polyline]())},
		{"Maplibre.Marker", len(*stager.GetMaplibreStage().GetInstancesSet[*maplibre.Marker]())},
		{"Cesium.MapView", len(*stager.GetCesiumStage().GetInstancesSet[*cesium.MapView]())},
		{"Cesium.Layer", len(*stager.GetCesiumStage().GetInstancesSet[*cesium.Layer]())},
		{"Cesium.Polygon", len(*stager.GetCesiumStage().GetInstancesSet[*cesium.Polygon]())},
		{"Cesium.Polyline", len(*stager.GetCesiumStage().GetInstancesSet[*cesium.Polyline]())},
		{"Cesium.Marker", len(*stager.GetCesiumStage().GetInstancesSet[*cesium.Marker]())},
	}

	for _, c := range checks {
		t.Logf("  %-35s: %d instances", c.name, c.count)
		if c.count == 0 {
			t.Errorf("expected > 0 instances for %s, got 0", c.name)
		}
	}

	errs := stager.GetSemanticErrors()
	t.Logf("ValidateXSDConstraints reported %d errors", len(errs))
	for i := 0; i < len(errs) && i < 20; i++ {
		t.Logf("  err[%d]: %s", i, errs[i])
	}

	// Verify that Eurocontrol ADR zones (e.g. LKTRA19Z, UKT732A, LFBB FIR) display their Authority node in the tree
	baselineTreeNodes := stager.GetTreeStage().GetInstancesMapByName[*tree.Node]()
	var lktra19zNode, ukt732aNode, lfbbNode *tree.Node
	for name, n := range baselineTreeNodes {
		if strings.HasPrefix(name, "[TRA] LKTRA19Z ") {
			lktra19zNode = n
		}
		if strings.HasPrefix(name, "[TSA] UKT732A ") {
			ukt732aNode = n
		}
		if strings.HasPrefix(name, "[PART] ") && strings.Contains(name, "LFBB FIR") {
			lfbbNode = n
		}
	}
	if lktra19zNode == nil {
		t.Fatalf("expected '[TRA] LKTRA19Z ...' FBZ zone node in Airspace.BASELINE tree")
	}
	if len(lktra19zNode.Children) == 0 || !strings.HasPrefix(lktra19zNode.Children[0].Name, "Authority: ") {
		t.Fatalf("expected first child of LKTRA19Z to start with 'Authority: ', got %v", lktra19zNode.Children)
	}
	t.Logf("LKTRA19Z Authority node: %s (%d sub-items)", lktra19zNode.Children[0].Name, len(lktra19zNode.Children[0].Children))
	hasOwnerRSA := false
	for _, sub := range lktra19zNode.Children[0].Children {
		if strings.Contains(sub.Name, "Owner RSA") && strings.Contains(sub.Name, "LKTRA19") {
			hasOwnerRSA = true
			break
		}
	}
	if !hasOwnerRSA {
		t.Errorf("expected LKTRA19Z Authority node to include resolved Owner RSA link to LKTRA19")
	}

	if ukt732aNode == nil {
		t.Fatalf("expected '[TSA] UKT732A ...' zone node in Airspace.BASELINE tree")
	}
	if len(ukt732aNode.Children) == 0 || !strings.HasPrefix(ukt732aNode.Children[0].Name, "Authority: ") {
		t.Fatalf("expected first child of UKT732A to start with 'Authority: ', got %v", ukt732aNode.Children)
	}
	t.Logf("UKT732A Authority node: %s", ukt732aNode.Children[0].Name)

	if lfbbNode == nil {
		t.Fatalf("expected '[PART] LFBB FIR' zone node in Airspace.BASELINE tree")
	}
	if len(lfbbNode.Children) == 0 || !strings.HasPrefix(lfbbNode.Children[0].Name, "Authority: ") {
		t.Fatalf("expected first child of LFBB FIR to start with 'Authority: ', got %v", lfbbNode.Children)
	}
	t.Logf("LFBB FIR Authority node: %s", lfbbNode.Children[0].Name)
}

func TestUnmarshalDonlon(t *testing.T) {
	donlonPath := filepath.Join("..", "cmd", domainName, "data", "donlon_"+domainName+".xml")
	xmlBytes, err := os.ReadFile(donlonPath)
	if err != nil {
		t.Fatalf("failed reading %s: %v", donlonPath, err)
	}

	mux := http.NewServeMux()
	stage := models.NewStage(domainName + "_donlon")
	prb := probe.NewProbe(mux, embeddedgo.GoModelsDir, embeddedgo.GoDiagramsDir, true, stage)
	stage.SetProbeIF(prb)
	stager := models.NewStager(mux, stage, prb)

	stage.Reset()
	if err := models.UnmarshallXMLToStage(stage, xmlBytes); err != nil {
		t.Fatalf("UnmarshallXMLToStage on %s failed: %v", donlonPath, err)
	}
	stage.Commit()

	msgs := stage.GetInstancesSorted[*models.AIXMBasicMessage]()
	if len(msgs) != 1 {
		t.Fatalf("expected 1 AIXMBasicMessage, got %d", len(msgs))
	}
	if len(msgs[0].HasMember) == 0 {
		t.Fatalf("expected > 0 HasMember in Donlon %s, got 0", domainName)
	}
	if len(*stage.GmlStage.GetInstancesSet[*gml.RingType]()) == 0 {
		t.Fatalf("expected > 0 RingType in Donlon %s, got 0", domainName)
	}
	nPolys := len(*stager.GetMaplibreStage().GetInstancesSet[*maplibre.Polygon]())
	nLines := len(*stager.GetMaplibreStage().GetInstancesSet[*maplibre.Polyline]())
	nMarkers := len(*stager.GetMaplibreStage().GetInstancesSet[*maplibre.Marker]())
	if nPolys == 0 || nMarkers == 0 {
		t.Fatalf("expected > 0 Maplibre polygons and markers for Donlon airspace, got %d polygons, %d markers", nPolys, nMarkers)
	}
	errs := models.ValidateXSDConstraints(stage)
	t.Logf("Donlon %s: unmarshalled %d members -> Maplibre: %d Polygons, %d Polylines, %d SVG Markers (%d XSD constraint notices)",
		domainName, len(msgs[0].HasMember), nPolys, nLines, nMarkers, len(errs))
	for i := 0; i < len(errs) && i < 10; i++ {
		t.Logf("  err[%d]: %s", i, errs[i])
	}

	// Verify intermediate area-first-letter level between Zone Categories and Zones in the UX tree
	treeNodes := stager.GetTreeStage().GetInstancesMapByName[*tree.Node]()
	firCatNode, ok := treeNodes["[FIR] Flight Information Region (2)"]
	if !ok {
		t.Fatalf("expected '[FIR] Flight Information Region (2)' category node in tree")
	}
	if len(firCatNode.Children) != 2 {
		t.Fatalf("expected 2 first-letter group nodes ([E], [K]) under FIR, got %d", len(firCatNode.Children))
	}
	expectedFIRLetterGroups := []string{
		"[E] EAAD (1)",
		"[K] KAAD (1)",
	}
	for idx, wantName := range expectedFIRLetterGroups {
		grpNode := firCatNode.Children[idx]
		if grpNode.Name != wantName {
			t.Errorf("FIR letter group[%d]: expected Name %q, got %q", idx, wantName, grpNode.Name)
		}
		if len(grpNode.Buttons) == 0 {
			t.Errorf("FIR letter group[%d]: expected visibility toggle button on intermediate group node", idx)
		}
	}

	dOtherCatNode, ok := treeNodes["[D_OTHER] Danger / Activity Area (13)"]
	if !ok {
		t.Fatalf("expected '[D_OTHER] Danger / Activity Area (13)' category node in tree")
	}
	if len(dOtherCatNode.Children) != 1 {
		t.Fatalf("expected 1 first-letter group node ([E] EAV1 -> EAV13 (13)) under D_OTHER, got %d", len(dOtherCatNode.Children))
	}
	eGroup := dOtherCatNode.Children[0]
	if eGroup.Name != "[E] EAV1 -> EAV13 (13)" {
		t.Fatalf("expected '[E] EAV1 -> EAV13 (13)', got %q", eGroup.Name)
	}
	if len(eGroup.Children) != 13 {
		t.Fatalf("expected 13 zone nodes inside '[E] EAV1 -> EAV13 (13)', got %d", len(eGroup.Children))
	}
	for i := 1; i <= 13; i++ {
		wantZonePrefix := fmt.Sprintf("[D_OTHER] EAV%d ", i)
		if !strings.HasPrefix(eGroup.Children[i-1].Name, wantZonePrefix) {
			t.Errorf("zone[%d] in [E]: expected Name starting with %q, got %q", i-1, wantZonePrefix, eGroup.Children[i-1].Name)
		}
		if len(eGroup.Children[i-1].Children) == 0 || !strings.HasPrefix(eGroup.Children[i-1].Children[0].Name, "Authority: ") {
			t.Errorf("zone[%d] (%s): expected first child to start with 'Authority: '", i-1, eGroup.Children[i-1].Name)
		}
	}

	// Toggle visibility on the [E] EAAD (1) intermediate group button and verify [K] KAAD (1) remains visible
	firstFIRGrp := firCatNode.Children[0]
	firstFIRGrp.Buttons[0].OnClick()
	treeNodes = stager.GetTreeStage().GetInstancesMapByName[*tree.Node]()
	firstFIRGrp = treeNodes["[E] EAAD (1)"]
	if firstFIRGrp.FontStyle != tree.ITALIC {
		t.Errorf("expected '[E] EAAD (1)' to have FontStyle ITALIC after hiding group, got %q", firstFIRGrp.FontStyle)
	}
	secondFIRGrp := treeNodes["[K] KAAD (1)"]
	if secondFIRGrp.FontStyle == tree.ITALIC {
		t.Errorf("expected '[K] KAAD (1)' to remain NORMAL after hiding [E], got %q", secondFIRGrp.FontStyle)
	}
}

func TestPopulateDomainUMLDiagrams(t *testing.T) {
	diagramlayout.RunDomainUMLDiagramsTest(t, diagramlayout.DomainConfig{
		DomainName:         domainName,
		DomainImportPrefix: domainImportPrefix,
		GoModelsDir:        embeddedgo.GoModelsDir,
		GoDiagramsDir:      embeddedgo.GoDiagramsDir,
		Specs:              buildDomainDiagramSpecs(),
		StrictMaxWidth:     1850,
		StrictMaxHeight:    1300,
	})
}

func buildDomainDiagramSpecs() []diagramlayout.DiagramSpec {
	return []diagramlayout.DiagramSpec{
		{
			Name:        "01. Airspace Message & Temporality",
			NoteName:    "Note01AirspaceMessage",
			Description: "Root AIXMBasicMessage container for Airspace features and the AIXM Temporality Model (Feature -> TimeSliceProperty -> TimeSlice).",
			ExtraCols: [5][]diagramlayout.StructRef{
				0: {
					{Pkg: "models", Name: "AIXMBasicMessage"},
				},
				1: {
					{Pkg: "aixmabstractgmlobjecttypes", Name: "MessageMetadataPropertyType"},
					{Pkg: "models", Name: "BasicMessageMemberAIXMPropertyType"},
				},
				2: {
					{Pkg: "aixmfeatures", Name: "AirspaceType"},
					{Pkg: "aixmfeatures", Name: "GeoBorderType"},
				},
				3: {
					{Pkg: "aixmfeatures", Name: "AirspaceTimeSlicePropertyType"},
					{Pkg: "aixmfeatures", Name: "GeoBorderTimeSlicePropertyType"},
				},
				4: {
					{Pkg: "aixmfeatures", Name: "AirspaceTimeSliceType"},
					{Pkg: "aixmfeatures", Name: "GeoBorderTimeSliceType"},
					{Pkg: "gml", Name: "TimePrimitivePropertyType"},
				},
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmabstractgmlobjecttypes", Name: "Enum_Interpretation_SimpleType"},
			},
		},
		{
			Name:        "02. Airspace & Geometry Volumes",
			NoteName:    "Note02AirspaceVolumes",
			Description: "Airspace feature, AirspaceGeometryComponent, AirspaceVolume, AirspaceVolumeDependency, vertical distance limits, and Surface projection.",
			Features: []string{
				"Airspace",
			},
			Objects: []string{
				"AirspaceGeometryComponent",
				"AirspaceVolume",
				"AirspaceVolumeDependency",
				"Surface",
			},
			ExtraCols: [5][]diagramlayout.StructRef{
				3: {
					{Pkg: "aixmdatatypes", Name: "ValDistanceVerticalType"},
					{Pkg: "aixmdatatypes", Name: "ValDistanceType"},
				},
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeAirspaceBaseType_Member_1"},
			},
		},
		{
			Name:        "03. Airspace Activation & Layer Classification",
			NoteName:    "Note03AirspaceActivation",
			Description: "AirspaceActivation, AirspaceLayerClass, AirspaceLayer, and Timesheet schedule properties.",
			Features: []string{
				"Airspace",
			},
			Objects: []string{
				"AirspaceActivation",
				"AirspaceLayerClass",
				"AirspaceLayer",
				"Timesheet",
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeAirspaceClassificationBaseType_Member_1"},
			},
		},
		{
			Name:        "04. AuthorityForAirspace & GeoBorder",
			NoteName:    "Note04GeoBorder",
			Description: "AuthorityForAirspace, GeoBorder, and border Curve/ElevatedCurve geometry.",
			Features: []string{
				"AuthorityForAirspace",
				"GeoBorder",
			},
			Objects: []string{
				"Curve",
				"ElevatedCurve",
				"Note",
			},
			ExtraCols: [5][]diagramlayout.StructRef{
				1: {
					{Pkg: "aixmfeatures", Name: "OrganisationAuthorityPropertyType"},
					{Pkg: "aixmfeatures", Name: "AirspacePropertyType"},
				},
			},
		},
		{
			Name:        "05. Airspace GML 3.2.1 Surface & Curve Segments",
			NoteName:    "Note05GMLGeometry",
			Description: "GML 3.2.1 PolygonPatch, LinearRing, GeodesicString, ArcByCenterPoint, and CircleByCenterPoint used by Airspace volumes and borders.",
			Objects: []string{
				"Curve",
			},
			ExtraCols: [5][]diagramlayout.StructRef{
				0: {
					{Pkg: "aixmfeatures", Name: "SurfaceType"},
				},
				1: {
					{Pkg: "gml", Name: "SurfacePatchArrayPropertyType"},
					{Pkg: "gml", Name: "PolygonPatchType"},
					{Pkg: "gml", Name: "AbstractRingPropertyType"},
					{Pkg: "gml", Name: "RingType"},
					{Pkg: "gml", Name: "LinearRingType"},
				},
				2: {
					{Pkg: "gml", Name: "CurveSegmentArrayPropertyType"},
					{Pkg: "gml", Name: "GeodesicStringType"},
					{Pkg: "gml", Name: "ArcByCenterPointType"},
					{Pkg: "gml", Name: "CircleByCenterPointType"},
				},
				3: {
					{Pkg: "gml", Name: "DirectPositionListType"},
					{Pkg: "gml", Name: "DirectPositionType"},
					{Pkg: "gml", Name: "PointPropertyType"},
				},
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "gml", Name: "Enum_CurveInterpolationType"},
			},
		},
	}
}
