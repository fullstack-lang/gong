package models

// GONGNOTE(Note01ServiceMessage): 01. Service & Organisation Message & Temporality Story
//
// An [models.AIXMBasicMessage] groups aviation authorities, ATC units, and services via [models.BasicMessageMemberAIXMPropertyType].
// Each feature ([models.OrganisationAuthorityType], [models.AirTrafficControlServiceType]) follows the AIXM Temporality Model
// using [models.OrganisationAuthorityTimeSliceType] slices bounded by ValidTime ([models.TimePrimitivePropertyType]).
const Note01ServiceMessage = ""

// GONGNOTE(Note02OrganisationAuthority): 02. OrganisationAuthority & ContactInformation Story
//
// [models.OrganisationAuthorityTimeSliceType] represents states, civil aviation authorities, ANSPs, or airlines.
// Inter-organisation relationships use [models.OrganisationAuthorityAssociationType], while reachability is modeled via
// [models.ContactInformationType] ([models.PostalAddressType], [models.TelephoneContactType], and [models.OnlineContactType]).
const Note02OrganisationAuthority = ""

// GONGNOTE(Note03Unit): 03. Unit, Availability & Dependency Story
//
// A [models.UnitTimeSliceType] models an operational ATS unit (ACC, APP, TWR, FIC, ARO, or MET office).
// Its operational hours are given by [models.UnitAvailabilityType], inter-unit delegation by [models.UnitDependencyType], and contact details by [models.ContactInformationType].
const Note03Unit = ""

// GONGNOTE(Note04TrafficSeparation): 04. Traffic Separation Services (ATC, ATM, Ground) Story
//
// Air traffic control services ([models.AirTrafficControlServiceTimeSliceType], [models.AirTrafficManagementServiceTimeSliceType], and
// [models.GroundTrafficControlServiceTimeSliceType]) link an operational unit to the airspace or airport served and track status via [models.ServiceOperationalStatusType].
const Note04TrafficSeparation = ""

// GONGNOTE(Note05InformationSAR): 05. Information & SearchRescue Services Story
//
// [models.InformationServiceTimeSliceType] (FIS, ATIS, AFIS, VOLMET) and [models.SearchRescueServiceTimeSliceType] (RCC/RSC)
// provide flight information and alerting/SAR coverage with real-time availability tracked in [models.ServiceOperationalStatusType].
const Note05InformationSAR = ""

// GONGNOTE(Note06AirportGroundServices): 06. Airport Ground Services & Supplies Story
//
// Aerodrome ground handling is modeled by [models.AircraftGroundServiceTimeSliceType] (de-icing, hangars), [models.AirportClearanceServiceTimeSliceType] (customs),
// [models.FireFightingServiceTimeSliceType] (RFF category), and [models.AirportSuppliesServiceTimeSliceType] ([models.FuelType], [models.OilType], [models.OxygenType]).
const Note06AirportGroundServices = ""

// GONGNOTE(Note07RadioComm): 07. RadioCommunicationChannel & PilotControlledLighting Story
//
// [models.RadioCommunicationChannelTimeSliceType] defines VHF/UHF/HF frequencies, emission class, [models.CallsignDetailType],
// [models.CircleSectorType] coverage, and [models.RadioCommunicationOperationalStatusType], while [models.PilotControlledLightingTimeSliceType] models radio-activated lighting.
const Note07RadioComm = ""

// GONGNOTE(Note08StandardLevels): 08. Shared Notes, Timesheets & StandardLevel Tables Story
//
// [models.SpecialDateTimeSliceType] defines public holidays referenced by [models.TimesheetType] schedules, while
// [models.StandardLevelTableTimeSliceType], [models.StandardLevelColumnTimeSliceType], and [models.StandardLevelTimeSliceType] define cruising level systems (RVSM/IFR/VFR).
const Note08StandardLevels = ""

