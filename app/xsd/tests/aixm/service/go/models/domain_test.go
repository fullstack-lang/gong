package models_test

import (
	"os"
	"path/filepath"
	"testing"

	embeddedgo "github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/diagramlayout"
)

const (
	domainName         = "service"
	domainImportPrefix = "github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models"
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
			Name:        "01. Service & Organisation Message & Temporality",
			NoteName:    "Note01ServiceMessage",
			Description: "Root AIXMBasicMessage container for Service/Organisation features and the AIXM Temporality Model.",
			ExtraCols: [5][]diagramlayout.StructRef{
				0: {
					{Pkg: "models", Name: "AIXMBasicMessage"},
				},
				1: {
					{Pkg: "aixmabstractgmlobjecttypes", Name: "MessageMetadataPropertyType"},
					{Pkg: "models", Name: "BasicMessageMemberAIXMPropertyType"},
				},
				2: {
					{Pkg: "aixmfeatures", Name: "OrganisationAuthorityType"},
					{Pkg: "aixmfeatures", Name: "UnitType"},
				},
				3: {
					{Pkg: "aixmfeatures", Name: "OrganisationAuthorityTimeSlicePropertyType"},
					{Pkg: "aixmfeatures", Name: "UnitTimeSlicePropertyType"},
				},
				4: {
					{Pkg: "aixmfeatures", Name: "OrganisationAuthorityTimeSliceType"},
					{Pkg: "aixmfeatures", Name: "UnitTimeSliceType"},
					{Pkg: "gml", Name: "TimePrimitivePropertyType"},
				},
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmabstractgmlobjecttypes", Name: "Enum_Interpretation_SimpleType"},
			},
		},
		{
			Name:        "02. OrganisationAuthority & ContactInformation",
			NoteName:    "Note02OrganisationAuthority",
			Description: "OrganisationAuthority feature, OrganisationAuthorityAssociation, ContactInformation, PostalAddress, TelephoneContact, and OnlineContact.",
			Features: []string{
				"OrganisationAuthority",
			},
			Objects: []string{
				"OrganisationAuthorityAssociation",
				"ContactInformation",
				"PostalAddress",
				"TelephoneContact",
				"OnlineContact",
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeOrganisationBaseType_Member_1"},
			},
		},
		{
			Name:        "03. Unit, Availability & Dependency",
			NoteName:    "Note03Unit",
			Description: "Unit (ATC/FIC/ARO/MET) feature, UnitAvailability, UnitDependency, and ContactInformation.",
			Features: []string{
				"Unit",
			},
			Objects: []string{
				"UnitAvailability",
				"UnitDependency",
				"ContactInformation",
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeUnitBaseType_Member_1"},
			},
		},
		{
			Name:        "04. Traffic Separation Services (ATC, ATM, Ground)",
			NoteName:    "Note04TrafficSeparation",
			Description: "AirTrafficControlService, AirTrafficManagementService, GroundTrafficControlService, and ServiceOperationalStatus.",
			Features: []string{
				"AirTrafficControlService",
				"AirTrafficManagementService",
				"GroundTrafficControlService",
			},
			Objects: []string{
				"ServiceOperationalStatus",
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeServiceATCBaseType_Member_1"},
				{Pkg: "aixmdatatypes", Name: "Enum_CodeStatusServiceBaseType_Member_1"},
			},
		},
		{
			Name:        "05. Information & SearchRescue Services",
			NoteName:    "Note05InformationSAR",
			Description: "InformationService (ATIS/AFIS/VOLMET), SearchRescueService, and ServiceOperationalStatus.",
			Features: []string{
				"InformationService",
				"SearchRescueService",
			},
			Objects: []string{
				"ServiceOperationalStatus",
			},
			Enums: []diagramlayout.StructRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeServiceInformationBaseType_Member_1"},
				{Pkg: "aixmdatatypes", Name: "Enum_CodeServiceSARBaseType_Member_1"},
			},
		},
		{
			Name:        "06. Airport Ground Services & Supplies",
			NoteName:    "Note06AirportGroundServices",
			Description: "AircraftGroundService, AirportClearanceService, AirportSuppliesService, FireFightingService, PassengerService, Fuel, Oil, Oxygen, and Nitrogen.",
			Features: []string{
				"AircraftGroundService",
				"AirportClearanceService",
				"AirportSuppliesService",
				"FireFightingService",
				"PassengerService",
			},
			Objects: []string{
				"Fuel",
				"Oil",
				"Oxygen",
				"Nitrogen",
			},
		},
		{
			Name:        "07. RadioCommunicationChannel & PilotControlledLighting",
			NoteName:    "Note07RadioComm",
			Description: "RadioCommunicationChannel, PilotControlledLighting, RadioCommunicationOperationalStatus, CallsignDetail, and CircleSector.",
			Features: []string{
				"RadioCommunicationChannel",
				"PilotControlledLighting",
			},
			Objects: []string{
				"RadioCommunicationOperationalStatus",
				"CallsignDetail",
				"CircleSector",
			},
		},
		{
			Name:        "08. Shared Notes, Timesheets & StandardLevel Tables",
			NoteName:    "Note08StandardLevels",
			Description: "SpecialDate, StandardLevelTable, StandardLevelColumn, StandardLevelSector, StandardLevel, Note, LinguisticNote, and Timesheet.",
			Features: []string{
				"SpecialDate",
				"StandardLevelTable",
				"StandardLevelColumn",
				"StandardLevelSector",
				"StandardLevel",
			},
			Objects: []string{
				"Note",
				"LinguisticNote",
				"Timesheet",
			},
		},
	}
}
