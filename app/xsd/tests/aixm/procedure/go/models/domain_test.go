package models_test

import (
	"os"
	"path/filepath"
	"testing"

	embeddedgo "github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/diagramlayout"
)

const (
	domainName         = "procedure"
	domainImportPrefix = "github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models"
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
			Name:        "01. Procedure Message & Temporality",
			NoteName:    "Note01ProcedureMessage",
			Description: "Root AIXMBasicMessage container for SID, STAR, IAP, SegmentLeg, and Holding features.",
			ExtraCols: [5][]diagramlayout.StructRef{
				0: {
					{Pkg: "models", Name: "AIXMBasicMessage"},
				},
				1: {
					{Pkg: "aixmabstractgmlobjecttypes", Name: "MessageMetadataPropertyType"},
					{Pkg: "models", Name: "BasicMessageMemberAIXMPropertyType"},
				},
				2: {
					{Pkg: "aixmfeatures", Name: "StandardInstrumentDepartureType"},
					{Pkg: "aixmfeatures", Name: "InstrumentApproachProcedureType"},
				},
				3: {
					{Pkg: "aixmfeatures", Name: "StandardInstrumentDepartureTimeSlicePropertyType"},
					{Pkg: "aixmfeatures", Name: "InstrumentApproachProcedureTimeSlicePropertyType"},
				},
				4: {
					{Pkg: "aixmfeatures", Name: "StandardInstrumentDepartureTimeSliceType"},
					{Pkg: "aixmfeatures", Name: "InstrumentApproachProcedureTimeSliceType"},
					{Pkg: "gml", Name: "TimePrimitivePropertyType"},
				},
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmabstractgmlobjecttypes", Name: "Enum_Interpretation_SimpleType"},
			},
		},
		{
			Name:        "02. SID & STAR Procedures",
			NoteName:    "Note02SIDSTAR",
			Description: "StandardInstrumentDeparture (SID), StandardInstrumentArrival (STAR), ProcedureTransition, and ProcedureAvailability.",
			Features: []string{
				"StandardInstrumentDeparture",
				"StandardInstrumentArrival",
			},
			Objects: []string{
				"ProcedureTransition",
				"ProcedureTransitionLeg",
				"ProcedureAvailability",
				"DepartureArrivalCondition",
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeProcedurePhaseBaseType_Member_1"},
				{Pkg: "aixmdatatypes", Name: "Enum_CodeProcedureCodingStandardBaseType_Member_1"},
			},
		},
		{
			Name:        "03. Instrument Approach Procedure (IAP)",
			NoteName:    "Note03IAP",
			Description: "InstrumentApproachProcedure, ApproachCondition, Minima, FinalProfile, FASDataBlock, and MissedApproachGroup.",
			Features: []string{
				"InstrumentApproachProcedure",
			},
			Objects: []string{
				"ApproachCondition",
				"Minima",
				"FinalProfile",
				"FASDataBlock",
				"MissedApproachGroup",
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeApproachBaseType_Member_1"},
				{Pkg: "aixmdatatypes", Name: "Enum_CodeApproachGuidanceBaseType_Member_1"},
			},
		},
		{
			Name:        "04. NavigationArea, TAA & CirclingArea",
			NoteName:    "Note04NavigationArea",
			Description: "NavigationArea, TerminalArrivalArea (TAA), CirclingArea, NavigationAreaRestriction, and SectorDesign.",
			Features: []string{
				"NavigationArea",
				"TerminalArrivalArea",
				"CirclingArea",
				"NavigationAreaRestriction",
			},
			Objects: []string{
				"NavigationAreaSector",
				"TerminalArrivalAreaSector",
				"SectorDesign",
			},
		},
		{
			Name:        "05. Departure & Arrival Segment Legs",
			NoteName:    "Note05DepartureArrivalLegs",
			Description: "DepartureLeg, ArrivalLeg, ArrivalFeederLeg, TerminalSegmentPoint, and LandingTakeoffAreaCollection.",
			Features: []string{
				"DepartureLeg",
				"ArrivalLeg",
				"ArrivalFeederLeg",
			},
			Objects: []string{
				"TerminalSegmentPoint",
				"LandingTakeoffAreaCollection",
			},
		},
		{
			Name:        "06. Approach & Missed Approach Segment Legs",
			NoteName:    "Note06ApproachLegs",
			Description: "InitialLeg, IntermediateLeg, FinalLeg, MissedApproachLeg, ProcedureDME, and HoldingUse.",
			Features: []string{
				"InitialLeg",
				"IntermediateLeg",
				"FinalLeg",
				"MissedApproachLeg",
				"ProcedureDME",
			},
			Objects: []string{
				"HoldingUse",
				"TerminalSegmentPoint",
			},
		},
		{
			Name:        "07. SafeAltitudeArea",
			NoteName:    "Note07SafeAltitude",
			Description: "SafeAltitudeArea (MSA/ESA), SafeAltitudeAreaSector, and CircleSector.",
			Features: []string{
				"SafeAltitudeArea",
			},
			Objects: []string{
				"SafeAltitudeAreaSector",
				"CircleSector",
			},
		},
		{
			Name:        "08. HoldingPattern & UnplannedHolding",
			NoteName:    "Note08HoldingPattern",
			Description: "HoldingPattern, UnplannedHolding, HoldingPatternDistance, and HoldingPatternDuration.",
			Features: []string{
				"HoldingPattern",
				"UnplannedHolding",
			},
			Objects: []string{
				"HoldingPatternDistance",
				"HoldingPatternDuration",
				"TerminalSegmentPoint",
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeHoldingCategoryBaseType_Member_1"},
				{Pkg: "aixmdatatypes", Name: "Enum_CodeHoldingUseBaseType_Member_1"},
			},
		},
		{
			Name:        "09. HoldingAssessment & Obstacle Assessment",
			NoteName:    "Note09HoldingAssessment",
			Description: "HoldingAssessment, AltitudeAdjustment, ObstacleAssessmentArea, ObstaclePlacement, and Obstruction.",
			Features: []string{
				"HoldingAssessment",
			},
			Objects: []string{
				"AltitudeAdjustment",
				"ObstacleAssessmentArea",
				"ObstaclePlacement",
				"Obstruction",
			},
		},
	}
}
