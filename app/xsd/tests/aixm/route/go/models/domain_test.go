package models_test

import (
	"os"
	"path/filepath"
	"testing"

	embeddedgo "github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/diagramlayout"
)

const (
	domainName         = "route"
	domainImportPrefix = "github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models"
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
			Name:        "01. Route Message & Temporality",
			NoteName:    "Note01RouteMessage",
			Description: "Root AIXMBasicMessage container for Route features and the AIXM Temporality Model (Feature -> TimeSliceProperty -> TimeSlice).",
			ExtraCols: [5][]diagramlayout.StructRef{
				0: {
					{Pkg: "models", Name: "AIXMBasicMessage"},
				},
				1: {
					{Pkg: "aixmabstractgmlobjecttypes", Name: "MessageMetadataPropertyType"},
					{Pkg: "models", Name: "BasicMessageMemberAIXMPropertyType"},
				},
				2: {
					{Pkg: "aixmfeatures", Name: "RouteType"},
					{Pkg: "aixmfeatures", Name: "RouteSegmentType"},
				},
				3: {
					{Pkg: "aixmfeatures", Name: "RouteTimeSlicePropertyType"},
					{Pkg: "aixmfeatures", Name: "RouteSegmentTimeSlicePropertyType"},
				},
				4: {
					{Pkg: "aixmfeatures", Name: "RouteTimeSliceType"},
					{Pkg: "aixmfeatures", Name: "RouteSegmentTimeSliceType"},
					{Pkg: "gml", Name: "TimePrimitivePropertyType"},
				},
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmabstractgmlobjecttypes", Name: "Enum_Interpretation_SimpleType"},
			},
		},
		{
			Name:        "02. Route & RoutePortion",
			NoteName:    "Note02RoutePortion",
			Description: "En-route ATS/RNAV Route feature, RoutePortion, and Note annotations.",
			Features: []string{
				"Route",
			},
			Objects: []string{
				"RoutePortion",
				"Note",
				"LinguisticNote",
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeRouteBaseType_Member_1"},
				{Pkg: "aixmdatatypes", Name: "Enum_CodeRouteNavigationBaseType_Member_1"},
			},
		},
		{
			Name:        "03. RouteSegment & RouteAvailability",
			NoteName:    "Note03RouteSegment",
			Description: "RouteSegment feature, start/end EnRouteSegmentPoint, RouteAvailability, AirspaceLayer, and Curve extent.",
			Features: []string{
				"RouteSegment",
			},
			Objects: []string{
				"RouteAvailability",
				"EnRouteSegmentPoint",
				"AirspaceLayer",
				"Curve",
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeRouteSegmentPathBaseType_Member_1"},
			},
		},
		{
			Name:        "04. RouteDME & ChangeOverPoint",
			NoteName:    "Note04ChangeOverPoint",
			Description: "RouteDME and ChangeOverPoint features along a RouteSegment.",
			Features: []string{
				"RouteDME",
				"ChangeOverPoint",
			},
			Objects: []string{
				"EnRouteSegmentPoint",
				"PointReference",
			},
		},
		{
			Name:        "05. FlightRestriction & Routing Elements",
			NoteName:    "Note05FlightRestriction",
			Description: "FlightRestriction feature, FlightRestrictionRoute, FlightRestrictionLevel, and FlightRoutingElement.",
			Features: []string{
				"FlightRestriction",
			},
			Objects: []string{
				"FlightRestrictionRoute",
				"FlightRestrictionLevel",
				"FlightRoutingElement",
			},
		},
		{
			Name:        "06. Flight Conditions, DirectFlight & RulesProcedures",
			NoteName:    "Note06FlightConditions",
			Description: "AirspaceBorderCrossing, RulesProcedures, FlightConditionCombination, FlightConditionElement, and DirectFlight.",
			Features: []string{
				"AirspaceBorderCrossing",
				"RulesProcedures",
			},
			Objects: []string{
				"FlightConditionCombination",
				"FlightConditionElement",
				"FlightConditionCircumstance",
				"DirectFlightSegment",
				"DirectFlightClass",
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeRuleProcedureBaseType_Member_1"},
			},
		},
		{
			Name:        "07. Aerial Refuelling",
			NoteName:    "Note07AerialRefuelling",
			Description: "AerialRefuelling feature, AerialRefuellingTrack, AerialRefuellingAnchor, AerialRefuellingPoint, and AuthorityForAerialRefuelling.",
			Features: []string{
				"AerialRefuelling",
			},
			Objects: []string{
				"AerialRefuellingTrack",
				"AerialRefuellingAnchor",
				"AerialRefuellingPoint",
				"AuthorityForAerialRefuelling",
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeAerialRefuellingBaseType_Member_1"},
				{Pkg: "aixmdatatypes", Name: "Enum_CodeAerialRefuellingPointBaseType_Member_1"},
			},
		},
	}
}
