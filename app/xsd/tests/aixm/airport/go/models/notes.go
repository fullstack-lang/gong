package models

// GONGNOTE(Note01AirportMessage): 01. Airport Message & Temporality Story
//
// An [models.AIXMBasicMessage] exchanges aerodrome and heliport infrastructure via [models.BasicMessageMemberAIXMPropertyType].
// Each feature ([models.AirportHeliportType], [models.RunwayType]) follows the AIXM Temporality Model: its state over time is carried
// by [models.AirportHeliportTimeSliceType] and [models.RunwayTimeSliceType] slices bounded by ValidTime ([models.TimePrimitivePropertyType]).
const Note01AirportMessage = ""

// GONGNOTE(Note02AirportOverview): 02. AirportHeliport, Availability, Usage & City Story
//
// [models.AirportHeliportTimeSliceType] holds core aerodrome identification (ICAO indicator, IATA code, ARP, elevation) and [models.CityType] served.
// Operational hours use [models.AirportHeliportAvailabilityType], traffic permissions use [models.AirportHeliportUsageType], and operator authority uses [models.AirportHeliportResponsibilityOrganisationType].
const Note02AirportOverview = ""

// GONGNOTE(Note03HotSpotWorkArea): 03. HotSpot, AltimeterSource, NonMovement & WorkArea Story
//
// Aerodrome safety and construction features include [models.AirportHotSpotTimeSliceType] (collision risk locations), [models.AltimeterSourceTimeSliceType]
// ([models.AltimeterSourceStatusType]), [models.WorkAreaTimeSliceType] ([models.WorkareaActivityType]), and [models.SurveyControlPointTimeSliceType].
const Note03HotSpotWorkArea = ""

// GONGNOTE(Note04Apron): 04. Apron & ApronElement Story
//
// [models.ApronTimeSliceType] and polygonal [models.ApronElementTimeSliceType] patches model aircraft parking and ramp areas,
// with bearing strength in [models.SurfaceCharacteristicsType], operational status in [models.ApronAreaAvailabilityType], and rules in [models.ApronAreaUsageType].
const Note04Apron = ""

// GONGNOTE(Note05AircraftStand): 05. AircraftStand, LoadingBridge, DeicingArea & Road Story
//
// Ramp facilities include [models.AircraftStandTimeSliceType] (gate/stand coordinates and visual docking), [models.PassengerLoadingBridgeTimeSliceType],
// [models.DeicingAreaTimeSliceType] pads, and airside [models.RoadTimeSliceType] segments, sharing [models.SurfaceCharacteristicsType] and [models.ApronAreaAvailabilityType].
const Note05AircraftStand = ""

// GONGNOTE(Note06TLOF): 06. TouchDownLiftOff (TLOF) & SafeArea Story
//
// Heliport landing pads are modeled by [models.TouchDownLiftOffTimeSliceType] (TLOF) and surrounding [models.TouchDownLiftOffSafeAreaTimeSliceType],
// with load-bearing properties in [models.SurfaceCharacteristicsType], status in [models.ManoeuvringAreaAvailabilityType], and restrictions in [models.ManoeuvringAreaUsageType].
const Note06TLOF = ""

// GONGNOTE(Note07Runway): 07. Runway, RunwayElement & Surface Story
//
// A [models.RunwayTimeSliceType] represents a physical landing strip (or FATO) composed of one or more [models.RunwayElementTimeSliceType] pavement sections,
// characterized by PCN/pavement type ([models.SurfaceCharacteristicsType]), [models.ManoeuvringAreaAvailabilityType], and [models.ManoeuvringAreaUsageType].
const Note07Runway = ""

// GONGNOTE(Note08RunwayDirection): 08. RunwayDirection, CentrelinePoint & DeclaredDistance Story
//
// Eachlanding/take-off orientation of a runway is a [models.RunwayDirectionTimeSliceType] (e.g. 09L / 27R).
// Thresholds and stop-ends are [models.RunwayCentrelinePointTimeSliceType] points publishing TORA, TODA, ASDA, and LDA in [models.RunwayDeclaredDistanceType] and [models.RunwayDeclaredDistanceValueType].
const Note08RunwayDirection = ""

// GONGNOTE(Note09BlastPadVGSI): 09. BlastPad, ProtectArea, RVR, VGSI & ArrestingGear Story
//
// Runway safety and visual aids include [models.RunwayBlastPadTimeSliceType], [models.RunwayProtectAreaTimeSliceType] (RESA),
// [models.RunwayVisualRangeTimeSliceType] (RVR sensors), [models.VisualGlideSlopeIndicatorTimeSliceType] (PAPI/VASI), and [models.ArrestingGearTimeSliceType].
const Note09BlastPadVGSI = ""

// GONGNOTE(Note10Taxiway): 10. Taxiway, TaxiwayElement, HoldingPosition & GuidanceLine Story
//
// Movement paths between runways and aprons use [models.TaxiwayTimeSliceType], [models.TaxiwayElementTimeSliceType], and [models.GuidanceLineTimeSliceType],
// protected by runway-incursion [models.TaxiHoldingPositionTimeSliceType] bars and [models.ManoeuvringAreaAvailabilityType].
const Note10Taxiway = ""

// GONGNOTE(Note11Seaplane): 11. SeaplaneLandingArea, RampSite, DockSite & Buoy Story
//
// Water aerodrome facilities are modeled by [models.SeaplaneLandingAreaTimeSliceType] channels, [models.SeaplaneRampSiteTimeSliceType],
// [models.FloatingDockSiteTimeSliceType], and boundary [models.MarkingBuoyTimeSliceType] markers with [models.ManoeuvringAreaAvailabilityType].
const Note11Seaplane = ""

// GONGNOTE(Note12RunwayLighting): 12. Approach & Runway Ground Lighting Story
//
// [models.ApproachLightingSystemTimeSliceType] (ALS), [models.RunwayDirectionLightSystemTimeSliceType] (edge/centreline/TDZ lights), and
// [models.RunwayProtectAreaLightSystemTimeSliceType] group individual [models.LightElementType] fixtures, [models.LightActivationType], and [models.GroundLightingAvailabilityType].
const Note12RunwayLighting = ""

// GONGNOTE(Note13TaxiwayLighting): 13. Taxiway, Apron, GuidanceLine & TLOF Lighting Story
//
// Surface movement lighting includes [models.TaxiwayLightSystemTimeSliceType], [models.TaxiHoldingPositionLightSystemTimeSliceType] (stop bars),
// [models.GuidanceLineLightSystemTimeSliceType], and [models.ApronLightSystemTimeSliceType], referencing [models.LightElementType] and [models.GroundLightingAvailabilityType].
const Note13TaxiwayLighting = ""

// GONGNOTE(Note14RunwayMarking): 14. Runway, Taxiway & HoldingPosition Marking Story
//
// Painted pavement markings on runways and taxiways are modeled by [models.RunwayMarkingTimeSliceType], [models.TaxiwayMarkingTimeSliceType],
// and [models.TaxiHoldingPositionMarkingTimeSliceType], each composed of geometric [models.MarkingElementType] primitives.
const Note14RunwayMarking = ""

// GONGNOTE(Note15StandContamination): 15. Stand, Apron, Deicing, TLOF Marking & Contamination Story
//
// Ramp markings ([models.StandMarkingTimeSliceType], [models.ApronMarkingTimeSliceType], [models.DeicingAreaMarkingTimeSliceType]) use [models.MarkingElementType],
// while SNOWTAM winter surface conditions are modeled by [models.AirportHeliportContaminationType], [models.RunwayContaminationType], and [models.SurfaceContaminationLayerType].
const Note15StandContamination = ""

