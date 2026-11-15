package models_test

import (
	"os"
	"path/filepath"
	"testing"

	embeddedgo "github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/diagramlayout"
)

const (
	domainName         = "obstacle"
	domainImportPrefix = "github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models"
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

func TestUnmarshalDonlon(t *testing.T) {
	donlonPath := filepath.Join("..", "cmd", domainName, "data", "donlon_"+domainName+".xml")
	xmlBytes, err := os.ReadFile(donlonPath)
	if err != nil {
		t.Fatalf("failed reading %s: %v", donlonPath, err)
	}

	stage := models.NewStage(domainName + "_donlon")
	if err := models.UnmarshallXMLToStage(stage, xmlBytes); err != nil {
		t.Fatalf("UnmarshallXMLToStage on %s failed: %v", donlonPath, err)
	}

	msgs := stage.GetInstancesSorted[*models.AIXMBasicMessage]()
	if len(msgs) != 1 {
		t.Fatalf("expected 1 AIXMBasicMessage, got %d", len(msgs))
	}
	if len(msgs[0].HasMember) == 0 {
		t.Fatalf("expected > 0 HasMember in Donlon %s, got 0", domainName)
	}
	errs := models.ValidateXSDConstraints(stage)
	t.Logf("Donlon %s: unmarshalled %d members (%d XSD constraint notices)", domainName, len(msgs[0].HasMember), len(errs))
	for i := 0; i < len(errs) && i < 10; i++ {
		t.Logf("  err[%d]: %s", i, errs[i])
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
			Name:        "01. Obstacle Message & Temporality",
			NoteName:    "Note01ObstacleMessage",
			Description: "Root AIXMBasicMessage container for Obstacle features and the AIXM Temporality Model (Feature -> TimeSliceProperty -> TimeSlice).",
			ExtraCols: [5][]diagramlayout.StructRef{
				0: {
					{Pkg: "models", Name: "AIXMBasicMessage"},
				},
				1: {
					{Pkg: "aixmabstractgmlobjecttypes", Name: "MessageMetadataPropertyType"},
					{Pkg: "models", Name: "BasicMessageMemberAIXMPropertyType"},
				},
				2: {
					{Pkg: "aixmfeatures", Name: "VerticalStructureType"},
					{Pkg: "aixmfeatures", Name: "ObstacleAreaType"},
				},
				3: {
					{Pkg: "aixmfeatures", Name: "VerticalStructureTimeSlicePropertyType"},
					{Pkg: "aixmfeatures", Name: "ObstacleAreaTimeSlicePropertyType"},
				},
				4: {
					{Pkg: "aixmfeatures", Name: "VerticalStructureTimeSliceType"},
					{Pkg: "aixmfeatures", Name: "ObstacleAreaTimeSliceType"},
					{Pkg: "gml", Name: "TimePrimitivePropertyType"},
				},
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmabstractgmlobjecttypes", Name: "Enum_Interpretation_SimpleType"},
			},
		},
		{
			Name:        "02. VerticalStructure & Parts",
			NoteName:    "Note02VerticalStructure",
			Description: "VerticalStructure feature, VerticalStructurePart, VerticalStructureLightingStatus, and LightElement.",
			Features: []string{
				"VerticalStructure",
			},
			Objects: []string{
				"VerticalStructurePart",
				"VerticalStructureLightingStatus",
				"LightElement",
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeVerticalStructureBaseType_Member_1"},
			},
		},
		{
			Name:        "03. ObstacleArea",
			NoteName:    "Note03ObstacleArea",
			Description: "ObstacleArea (Area 1/2/3/4), linked VerticalStructures, and Surface/ElevatedSurface extent.",
			Features: []string{
				"ObstacleArea",
			},
			Objects: []string{
				"Surface",
				"ElevatedSurface",
				"Note",
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeObstacleAreaBaseType_Member_1"},
			},
		},
		{
			Name:        "04. AeronauticalGroundLight",
			NoteName:    "Note04GroundLight",
			Description: "AeronauticalGroundLight beacon/lighthouse, LightElement, LightElementStatus, and ElevatedPoint location.",
			Features: []string{
				"AeronauticalGroundLight",
			},
			Objects: []string{
				"LightElement",
				"LightElementStatus",
				"ElevatedPoint",
			},
		},
		{
			Name:        "05. Elevated 3D Geometry (Point, Curve, Surface)",
			NoteName:    "Note05ElevatedGeometry",
			Description: "ElevatedPoint, ElevatedCurve, and ElevatedSurface used for 3D obstacle coordinates and vertical elevations.",
			ExtraCols: [5][]diagramlayout.StructRef{
				0: {
					{Pkg: "aixmfeatures", Name: "ElevatedPointType"},
					{Pkg: "aixmfeatures", Name: "ElevatedCurveType"},
					{Pkg: "aixmfeatures", Name: "ElevatedSurfaceType"},
				},
				1: {
					{Pkg: "gml", Name: "CurveSegmentArrayPropertyType"},
					{Pkg: "gml", Name: "SurfacePatchArrayPropertyType"},
				},
				2: {
					{Pkg: "gml", Name: "GeodesicStringType"},
					{Pkg: "gml", Name: "PolygonPatchType"},
					{Pkg: "gml", Name: "AbstractRingPropertyType"},
					{Pkg: "gml", Name: "LinearRingType"},
				},
				3: {
					{Pkg: "gml", Name: "DirectPositionType"},
					{Pkg: "gml", Name: "DirectPositionListType"},
				},
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeVerticalDatumBaseType_Member_1"},
				{Pkg: "aixmdatatypes", Name: "Enum_UomDistanceVerticalType_Member_1"},
			},
		},
	}
}
