package models_test

import (
	"net/http"
	"os"
	"path/filepath"
	"testing"

	embeddedgo "github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/probe"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/diagramlayout"
)

const (
	domainName         = "airport"
	domainImportPrefix = "github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models"
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

func TestUnmarshalAirportHeliportBaseline(t *testing.T) {
	baselinePath := filepath.Join("..", "cmd", domainName, "data", "AirportHeliport.BASELINE")
	xmlBytes, err := os.ReadFile(baselinePath)
	if err != nil {
		if os.IsNotExist(err) {
			t.Skipf("skipping baseline test, file not found: %s", baselinePath)
		}
		t.Fatalf("failed reading %s: %v", baselinePath, err)
	}

	mux := http.NewServeMux()
	stage := models.NewStage(domainName + "_baseline")
	prb := probe.NewProbe(mux, embeddedgo.GoModelsDir, embeddedgo.GoDiagramsDir, true, stage)
	stage.SetProbeIF(prb)
	stager := models.NewStager(mux, stage, prb)

	stage.Reset()
	if err := models.UnmarshallXMLToStage(stage, xmlBytes); err != nil {
		t.Fatalf("UnmarshallXMLToStage on AirportHeliport.BASELINE failed: %v", err)
	}
	stage.Commit()

	msgs := stage.GetInstancesSorted[*models.AIXMBasicMessage]()
	if len(msgs) != 1 {
		t.Fatalf("expected 1 AIXMBasicMessage, got %d", len(msgs))
	}
	airports := *stage.AixmfeaturesStage.GetInstancesSet[*aixmfeatures.AirportHeliportType]()
	if len(airports) != 10134 {
		t.Fatalf("expected 10134 AirportHeliportType instances, got %d", len(airports))
	}
	t.Logf("Unmarshalled ADRMessage into AIXMBasicMessage (Id=%s) with %d members and %d AirportHeliportType instances",
		msgs[0].Id, len(msgs[0].HasMember), len(airports))

	errs := stager.GetSemanticErrors()
	t.Logf("ValidateXSDConstraints reported %d errors", len(errs))
	for i := 0; i < len(errs) && i < 20; i++ {
		t.Logf("  err[%d]: %s", i, errs[i])
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
			Name:        "01. Airport Message & Temporality",
			NoteName:    "Note01AirportMessage",
			Description: "Root AIXMBasicMessage container for AirportHeliport, Runway, Taxiway, Apron, Lighting, and Marking features.",
			ExtraCols: [5][]diagramlayout.StructRef{
				0: {
					{Pkg: "models", Name: "AIXMBasicMessage"},
				},
				1: {
					{Pkg: "aixmabstractgmlobjecttypes", Name: "MessageMetadataPropertyType"},
					{Pkg: "models", Name: "BasicMessageMemberAIXMPropertyType"},
				},
				2: {
					{Pkg: "aixmfeatures", Name: "AirportHeliportType"},
					{Pkg: "aixmfeatures", Name: "RunwayType"},
				},
				3: {
					{Pkg: "aixmfeatures", Name: "AirportHeliportTimeSlicePropertyType"},
					{Pkg: "aixmfeatures", Name: "RunwayTimeSlicePropertyType"},
				},
				4: {
					{Pkg: "aixmfeatures", Name: "AirportHeliportTimeSliceType"},
					{Pkg: "aixmfeatures", Name: "RunwayTimeSliceType"},
					{Pkg: "gml", Name: "TimePrimitivePropertyType"},
				},
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmabstractgmlobjecttypes", Name: "Enum_Interpretation_SimpleType"},
			},
		},
		{
			Name:        "02. AirportHeliport, Availability, Usage & City",
			NoteName:    "Note02AirportOverview",
			Description: "AirportHeliport, AirportHeliportCollocation, Availability, Usage, ResponsibilityOrganisation, and Served City.",
			Features: []string{
				"AirportHeliport",
				"AirportHeliportCollocation",
			},
			Objects: []string{
				"AirportHeliportAvailability",
				"AirportHeliportUsage",
				"AirportHeliportResponsibilityOrganisation",
				"City",
				"ConditionCombination",
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeAirportHeliportBaseType_Member_1"},
				{Pkg: "aixmdatatypes", Name: "Enum_CodeOperationAirportHeliportBaseType_Member_1"},
			},
		},
		{
			Name:        "03. HotSpot, AltimeterSource, NonMovement & WorkArea",
			NoteName:    "Note03HotSpotWorkArea",
			Description: "AirportHotSpot, AltimeterSource, NonMovementArea, SurveyControlPoint, WorkArea, and WorkareaActivity.",
			Features: []string{
				"AirportHotSpot",
				"AltimeterSource",
				"NonMovementArea",
				"SurveyControlPoint",
				"WorkArea",
			},
			Objects: []string{
				"AltimeterSourceStatus",
				"WorkareaActivity",
			},
		},
		{
			Name:        "04. Apron & ApronElement",
			NoteName:    "Note04Apron",
			Description: "Apron, ApronElement, ApronAreaAvailability, ApronAreaUsage, and SurfaceCharacteristics.",
			Features: []string{
				"Apron",
				"ApronElement",
			},
			Objects: []string{
				"ApronAreaAvailability",
				"ApronAreaUsage",
				"SurfaceCharacteristics",
			},
		},
		{
			Name:        "05. AircraftStand, LoadingBridge, DeicingArea & Road",
			NoteName:    "Note05AircraftStand",
			Description: "AircraftStand, PassengerLoadingBridge, DeicingArea, Road, and SurfaceCharacteristics.",
			Features: []string{
				"AircraftStand",
				"PassengerLoadingBridge",
				"DeicingArea",
				"Road",
			},
			Objects: []string{
				"ApronAreaAvailability",
				"SurfaceCharacteristics",
			},
		},
		{
			Name:        "06. TouchDownLiftOff (TLOF) & SafeArea",
			NoteName:    "Note06TLOF",
			Description: "Heliport TouchDownLiftOff (TLOF), TouchDownLiftOffSafeArea, ManoeuvringAreaAvailability, and SurfaceCharacteristics.",
			Features: []string{
				"TouchDownLiftOff",
				"TouchDownLiftOffSafeArea",
			},
			Objects: []string{
				"ManoeuvringAreaAvailability",
				"ManoeuvringAreaUsage",
				"SurfaceCharacteristics",
			},
		},
		{
			Name:        "07. Runway, RunwayElement & Surface",
			NoteName:    "Note07Runway",
			Description: "Runway, RunwayElement, ManoeuvringAreaAvailability, ManoeuvringAreaUsage, and SurfaceCharacteristics.",
			Features: []string{
				"Runway",
				"RunwayElement",
			},
			Objects: []string{
				"ManoeuvringAreaAvailability",
				"ManoeuvringAreaUsage",
				"SurfaceCharacteristics",
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeRunwayBaseType_Member_1"},
				{Pkg: "aixmdatatypes", Name: "Enum_CodeRunwayElementBaseType_Member_1"},
			},
		},
		{
			Name:        "08. RunwayDirection, CentrelinePoint & DeclaredDistance",
			NoteName:    "Note08RunwayDirection",
			Description: "RunwayDirection, RunwayCentrelinePoint, RunwayDeclaredDistance, RunwayDeclaredDistanceValue, and NavaidEquipmentDistance.",
			Features: []string{
				"RunwayDirection",
				"RunwayCentrelinePoint",
			},
			Objects: []string{
				"RunwayDeclaredDistance",
				"RunwayDeclaredDistanceValue",
				"NavaidEquipmentDistance",
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeRunwayPointRoleBaseType_Member_1"},
			},
		},
		{
			Name:        "09. BlastPad, ProtectArea, RVR, VGSI & ArrestingGear",
			NoteName:    "Note09BlastPadVGSI",
			Description: "RunwayBlastPad, RunwayProtectArea, RunwayVisualRange (RVR), VisualGlideSlopeIndicator (VGSI), and ArrestingGear.",
			Features: []string{
				"RunwayBlastPad",
				"RunwayProtectArea",
				"RunwayVisualRange",
				"VisualGlideSlopeIndicator",
				"ArrestingGear",
			},
			Objects: []string{
				"SurfaceCharacteristics",
			},
		},
		{
			Name:        "10. Taxiway, TaxiwayElement, HoldingPosition & GuidanceLine",
			NoteName:    "Note10Taxiway",
			Description: "Taxiway, TaxiwayElement, TaxiHoldingPosition, GuidanceLine, and SurfaceCharacteristics.",
			Features: []string{
				"Taxiway",
				"TaxiwayElement",
				"TaxiHoldingPosition",
				"GuidanceLine",
			},
			Objects: []string{
				"ManoeuvringAreaAvailability",
				"SurfaceCharacteristics",
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeTaxiwayBaseType_Member_1"},
				{Pkg: "aixmdatatypes", Name: "Enum_CodeTaxiwayElementBaseType_Member_1"},
			},
		},
		{
			Name:        "11. SeaplaneLandingArea, RampSite, DockSite & Buoy",
			NoteName:    "Note11Seaplane",
			Description: "SeaplaneLandingArea, SeaplaneRampSite, FloatingDockSite, and MarkingBuoy.",
			Features: []string{
				"SeaplaneLandingArea",
				"SeaplaneRampSite",
				"FloatingDockSite",
				"MarkingBuoy",
			},
			Objects: []string{
				"ManoeuvringAreaAvailability",
			},
		},
		{
			Name:        "12. Approach & Runway Ground Lighting",
			NoteName:    "Note12RunwayLighting",
			Description: "ApproachLightingSystem, RunwayDirectionLightSystem, RunwayProtectAreaLightSystem, and LightElement.",
			Features: []string{
				"ApproachLightingSystem",
				"RunwayDirectionLightSystem",
				"RunwayProtectAreaLightSystem",
			},
			Objects: []string{
				"GroundLightingAvailability",
				"LightActivation",
				"LightElement",
				"LightElementStatus",
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeApproachLightingBaseType_Member_1"},
				{Pkg: "aixmdatatypes", Name: "Enum_CodeLightIntensityBaseType_Member_1"},
			},
		},
		{
			Name:        "13. Taxiway, Apron, GuidanceLine & TLOF Lighting",
			NoteName:    "Note13TaxiwayLighting",
			Description: "TaxiwayLightSystem, TaxiHoldingPositionLightSystem, GuidanceLineLightSystem, ApronLightSystem, and TouchDownLiftOffLightSystem.",
			Features: []string{
				"TaxiwayLightSystem",
				"TaxiHoldingPositionLightSystem",
				"GuidanceLineLightSystem",
				"ApronLightSystem",
				"TouchDownLiftOffLightSystem",
			},
			Objects: []string{
				"GroundLightingAvailability",
				"LightElement",
			},
		},
		{
			Name:        "14. Runway, Taxiway & HoldingPosition Marking",
			NoteName:    "Note14RunwayMarking",
			Description: "RunwayMarking, TaxiwayMarking, TaxiHoldingPositionMarking, and MarkingElement.",
			Features: []string{
				"RunwayMarking",
				"TaxiwayMarking",
				"TaxiHoldingPositionMarking",
			},
			Objects: []string{
				"MarkingElement",
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeMarkingConditionBaseType_Member_1"},
				{Pkg: "aixmdatatypes", Name: "Enum_CodeMarkingStyleBaseType_Member_1"},
			},
		},
		{
			Name:        "15. Stand, Apron, Deicing & TLOF Marking",
			NoteName:    "Note15StandContamination",
			Description: "StandMarking, ApronMarking, DeicingAreaMarking, GuidanceLineMarking, TouchDownLiftOffMarking, AirportProtectionAreaMarking, and MarkingElement.",
			Features: []string{
				"StandMarking",
				"ApronMarking",
				"DeicingAreaMarking",
				"GuidanceLineMarking",
				"TouchDownLiftOffMarking",
				"AirportProtectionAreaMarking",
			},
			Objects: []string{
				"MarkingElement",
			},
		},
		{
			Name:        "16. Airport & Runway Surface Contamination",
			NoteName:    "Note15StandContamination",
			Description: "SNOWTAM surface contamination: AirportHeliportContamination, RunwayContamination, RunwaySectionContamination, SurfaceContaminationLayer, and Ridge.",
			Features: []string{
				"AirportHeliport",
				"Runway",
			},
			Objects: []string{
				"AirportHeliportContamination",
				"RunwayContamination",
				"RunwaySectionContamination",
				"SurfaceContaminationLayer",
				"Ridge",
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeContaminationBaseType_Member_1"},
			},
		},
	}
}
