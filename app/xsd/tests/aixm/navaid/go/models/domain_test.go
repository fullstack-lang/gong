package models_test

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	embeddedgo "github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models/probe"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/diagramlayout"
	cesium "github.com/fullstack-lang/gong/lib/cesium/go/models"
	maplibre "github.com/fullstack-lang/gong/lib/maplibre/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

const (
	domainName         = "navaid"
	domainImportPrefix = "github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models"
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
	if errs := models.ValidateXSDConstraints(stage); len(errs) > 0 {
		t.Fatalf("expected 0 XSD validation errors on sample_%s.xml, got %d: %v", domainName, len(errs), errs)
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
	if errs := models.ValidateXSDConstraints(stage3); len(errs) > 0 {
		t.Fatalf("expected 0 XSD validation errors after round-trip, got %d: %v", len(errs), errs)
	}

	// Verify 2D MapLibre, 3D Cesium, and Tree projections on sample_navaid.xml
	mux := http.NewServeMux()
	prb := probe.NewProbe(mux, embeddedgo.GoModelsDir, embeddedgo.GoDiagramsDir, true, stage)
	stage.SetProbeIF(prb)
	stager := models.NewStager(mux, stage, prb)
	if len(stager.GetSemanticErrors()) > 0 {
		t.Fatalf("expected 0 stager semantic errors on sample_%s.xml, got %v", domainName, stager.GetSemanticErrors())
	}

	sampleChecks := []struct {
		name  string
		count int
	}{
		{"Sample Maplibre.MapView", len(*stager.GetMaplibreStage().GetInstancesSet[*maplibre.MapView]())},
		{"Sample Maplibre.Layer", len(*stager.GetMaplibreStage().GetInstancesSet[*maplibre.Layer]())},
		{"Sample Maplibre.Polygon", len(*stager.GetMaplibreStage().GetInstancesSet[*maplibre.Polygon]())},
		{"Sample Maplibre.Polyline", len(*stager.GetMaplibreStage().GetInstancesSet[*maplibre.Polyline]())},
		{"Sample Maplibre.Marker", len(*stager.GetMaplibreStage().GetInstancesSet[*maplibre.Marker]())},
		{"Sample Cesium.MapView", len(*stager.GetCesiumStage().GetInstancesSet[*cesium.MapView]())},
		{"Sample Cesium.Layer", len(*stager.GetCesiumStage().GetInstancesSet[*cesium.Layer]())},
		{"Sample Cesium.Polygon", len(*stager.GetCesiumStage().GetInstancesSet[*cesium.Polygon]())},
		{"Sample Cesium.Polyline", len(*stager.GetCesiumStage().GetInstancesSet[*cesium.Polyline]())},
		{"Sample Cesium.Marker", len(*stager.GetCesiumStage().GetInstancesSet[*cesium.Marker]())},
	}
	for _, c := range sampleChecks {
		t.Logf("  %-28s: %d instances", c.name, c.count)
		if c.count == 0 {
			t.Errorf("expected > 0 instances for %s, got 0", c.name)
		}
	}
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

	checks := []struct {
		name  string
		count int
	}{
		{"AIXMBasicMessage", len(*stage.GetInstancesSet[*models.AIXMBasicMessage]())},
		{"NavaidType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NavaidType]())},
		{"VORType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.VORType]())},
		{"DMEType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DMEType]())},
		{"TACANType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.TACANType]())},
		{"NDBType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.NDBType]())},
		{"LocalizerType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.LocalizerType]())},
		{"GlidepathType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.GlidepathType]())},
		{"MarkerBeaconType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.MarkerBeaconType]())},
		{"DesignatedPointType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DesignatedPointType]())},
		{"SignificantPointInAirspaceType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.SignificantPointInAirspaceType]())},
		{"AngleIndicationType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AngleIndicationType]())},
		{"DistanceIndicationType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.DistanceIndicationType]())},
		{"CheckpointVORType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CheckpointVORType]())},
		{"CheckpointINSType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.CheckpointINSType]())},
		{"ElevatedPointType", len(*stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.ElevatedPointType]())},
		{"DirectPositionType", len(*stage.GmlStage.GetInstancesSet[*gml.DirectPositionType]())},
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
	t.Logf("Donlon %s: unmarshalled %d members (%d XSD constraint notices)", domainName, len(msgs[0].HasMember), len(errs))
	for i := 0; i < len(errs) && i < 10; i++ {
		t.Logf("  err[%d]: %s", i, errs[i])
	}

	// Verify domain folders, selection sync, and visibility eye toggle callbacks
	treeNodes := stager.GetTreeStage().GetInstancesMapByName[*tree.Node]()
	expectedFolderPrefixes := []string{
		"Composite Navaid Systems (",
		"Radio Navaid Equipment (",
		"Approach & Landing Guidance (",
		"Waypoints & Airspace Points (",
		"Bearings & Distance Indications (",
		"Checkpoints & Surveillance Radars (",
	}
	for _, pfx := range expectedFolderPrefixes {
		found := false
		for name := range treeNodes {
			if strings.HasPrefix(name, pfx) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected domain folder starting with %q in Donlon navaid tree", pfx)
		}
	}

	// Find BOR VOR/DME node and test OnClick selection + visibility toggle
	var borNode *tree.Node
	for name, n := range treeNodes {
		if strings.HasPrefix(name, "[VOR_DME] BOR ") {
			borNode = n
			break
		}
	}
	if borNode == nil {
		t.Fatalf("expected '[VOR_DME] BOR ...' node in Donlon navaid tree")
	}
	initialMarkers := len(*stager.GetMaplibreStage().GetInstancesSet[*maplibre.Marker]())
	if borNode.OnClick != nil {
		borNode.OnClick(borNode)
	}
	afterSelectNodes := stager.GetTreeStage().GetInstancesMapByName[*tree.Node]()
	selectedFound := false
	var selectedBorNode *tree.Node
	for name, n := range afterSelectNodes {
		if strings.HasPrefix(name, "> [VOR_DME] BOR ") {
			selectedFound = true
			selectedBorNode = n
			break
		}
	}
	if !selectedFound || selectedBorNode == nil {
		t.Fatalf("expected BOR node to be prefixed with '> ' after tree click")
	}
	if len(selectedBorNode.Buttons) == 0 || selectedBorNode.Buttons[0].OnClick == nil {
		t.Fatalf("expected visibility toggle button on BOR node")
	}
	selectedBorNode.Buttons[0].OnClick()
	hiddenMarkers := len(*stager.GetMaplibreStage().GetInstancesSet[*maplibre.Marker]())
	if hiddenMarkers >= initialMarkers {
		t.Errorf("expected marker count to decrease when hiding BOR (before=%d, after=%d)", initialMarkers, hiddenMarkers)
	}
}

func TestPopulateDomainUMLDiagrams(t *testing.T) {
	diagramlayout.RunDomainUMLDiagramsTest(t, diagramlayout.DomainConfig{
		DomainName:         domainName,
		DomainImportPrefix: domainImportPrefix,
		GoModelsDir:        embeddedgo.GoModelsDir,
		GoDiagramsDir:      embeddedgo.GoDiagramsDir,
		Specs:              buildDomainDiagramSpecs(),
		StrictMaxWidth:     1900,
		StrictMaxHeight:    1300,
	})
}

func buildDomainDiagramSpecs() []diagramlayout.DiagramSpec {
	return []diagramlayout.DiagramSpec{
		{
			Name:        "01. Navaid & Point Message & Temporality",
			NoteName:    "Note01NavaidMessage",
			Description: "Root AIXMBasicMessage container for Navaids, DesignatedPoints, and Radar features.",
			ExtraCols: [5][]diagramlayout.StructRef{
				0: {
					{Pkg: "models", Name: "AIXMBasicMessage"},
				},
				1: {
					{Pkg: "aixmabstractgmlobjecttypes", Name: "MessageMetadataPropertyType"},
					{Pkg: "models", Name: "BasicMessageMemberAIXMPropertyType"},
				},
				2: {
					{Pkg: "aixmfeatures", Name: "NavaidType"},
					{Pkg: "aixmfeatures", Name: "DesignatedPointType"},
				},
				3: {
					{Pkg: "aixmfeatures", Name: "NavaidTimeSlicePropertyType"},
					{Pkg: "aixmfeatures", Name: "DesignatedPointTimeSlicePropertyType"},
				},
				4: {
					{Pkg: "aixmfeatures", Name: "NavaidTimeSliceType"},
					{Pkg: "aixmfeatures", Name: "DesignatedPointTimeSliceType"},
					{Pkg: "gml", Name: "TimePrimitivePropertyType"},
				},
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmabstractgmlobjecttypes", Name: "Enum_Interpretation_SimpleType"},
			},
		},
		{
			Name:        "02. Navaid, Component & Operational Status",
			NoteName:    "Note02NavaidComponent",
			Description: "Navaid system, NavaidComponent, NavaidOperationalStatus, Monitoring, and Authority.",
			Features: []string{
				"Navaid",
			},
			Objects: []string{
				"NavaidComponent",
				"NavaidOperationalStatus",
				"NavaidEquipmentMonitoring",
				"AuthorityForNavaidEquipment",
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeNavaidServiceBaseType_Member_1"},
				{Pkg: "aixmdatatypes", Name: "Enum_CodeNavaidPurposeBaseType_Member_1"},
				{Pkg: "aixmdatatypes", Name: "Enum_CodeStatusNavaidBaseType_Member_1"},
			},
		},
		{
			Name:        "03. En-Route Radio Navaids (VOR, DME, TACAN, NDB)",
			NoteName:    "Note03RadioNavaids",
			Description: "VOR, DME, TACAN, and NDB radio navigation equipment features.",
			Features: []string{
				"VOR",
				"DME",
				"TACAN",
				"NDB",
			},
			Objects: []string{
				"NavaidOperationalStatus",
				"ElevatedPoint",
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeVORBaseType_Member_1"},
				{Pkg: "aixmdatatypes", Name: "Enum_CodeDMEBaseType_Member_1"},
			},
		},
		{
			Name:        "04. ILS & MLS Precision Approach Navaids",
			NoteName:    "Note04ILSMLS",
			Description: "Localizer, Glidepath, MarkerBeacon, SDF, Azimuth, Elevation, and DirectionFinder.",
			Features: []string{
				"Localizer",
				"Glidepath",
				"MarkerBeacon",
				"SDF",
				"Azimuth",
				"Elevation",
				"DirectionFinder",
			},
			Objects: []string{
				"NavaidOperationalStatus",
			},
		},
		{
			Name:        "05. Navigation Checkpoints (VOR & INS)",
			NoteName:    "Note05Checkpoints",
			Description: "CheckpointVOR and CheckpointINS features.",
			Features: []string{
				"CheckpointVOR",
				"CheckpointINS",
			},
			Objects: []string{
				"ElevatedPoint",
				"Note",
			},
		},
		{
			Name:        "06. Special Navigation System & Station",
			NoteName:    "Note06SpecialNav",
			Description: "SpecialNavigationSystem, SpecialNavigationStation, SpecialNavigationStationStatus, and Authority objects.",
			Features: []string{
				"SpecialNavigationSystem",
				"SpecialNavigationStation",
			},
			Objects: []string{
				"SpecialNavigationStationStatus",
				"AuthorityForSpecialNavigationSystem",
				"AuthorityForSpecialNavigationStation",
			},
		},
		{
			Name:        "07. DesignatedPoint & SignificantPointInAirspace",
			NoteName:    "Note07DesignatedPoint",
			Description: "DesignatedPoint, SignificantPointInAirspace, EnRouteSegmentPoint, and TerminalSegmentPoint.",
			Features: []string{
				"DesignatedPoint",
				"SignificantPointInAirspace",
			},
			Objects: []string{
				"EnRouteSegmentPoint",
				"TerminalSegmentPoint",
				"PointReference",
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeDesignatedPointBaseType_Member_1"},
			},
		},
		{
			Name:        "08. AngleIndication & DistanceIndication",
			NoteName:    "Note08AngleDistance",
			Description: "AngleIndication, DistanceIndication, AngleUse, and PointReference.",
			Features: []string{
				"AngleIndication",
				"DistanceIndication",
			},
			Objects: []string{
				"AngleUse",
				"PointReference",
			},
		},
		{
			Name:        "09. Surveillance Radar Systems (PSR, SSR, PAR)",
			NoteName:    "Note09RadarSystems",
			Description: "RadarSystem, PrimarySurveillanceRadar, SecondarySurveillanceRadar, PrecisionApproachRadar, RadarComponent, and Reflector.",
			Features: []string{
				"RadarSystem",
				"PrimarySurveillanceRadar",
				"SecondarySurveillanceRadar",
				"PrecisionApproachRadar",
			},
			Objects: []string{
				"RadarComponent",
				"Reflector",
				"SurveillanceGroundStation",
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeRadarServiceBaseType_Member_1"},
				{Pkg: "aixmdatatypes", Name: "Enum_CodePrimaryRadarBaseType_Member_1"},
			},
		},
	}
}
