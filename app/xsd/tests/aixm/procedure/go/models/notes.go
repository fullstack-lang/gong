package models

// GONGNOTE(Note01ProcedureMessage): 01. Procedure Message & Temporality Story
//
// An [models.AIXMBasicMessage] exchanges terminal instrument flight procedures (SID, STAR, IAP) via [models.BasicMessageMemberAIXMPropertyType].
// Each procedure feature ([models.StandardInstrumentDepartureType], [models.InstrumentApproachProcedureType]) follows the Temporality Model
// using [models.StandardInstrumentDepartureTimeSliceType] slices bounded by ValidTime ([models.TimePrimitivePropertyType]).
const Note01ProcedureMessage = ""

// GONGNOTE(Note02SIDSTAR): 02. SID & STAR Procedures Story
//
// [models.StandardInstrumentDepartureTimeSliceType] (SID) and [models.StandardInstrumentArrivalTimeSliceType] (STAR) connect runway directions
// to en-route transitions via ordered [models.ProcedureTransitionType] branches, [models.ProcedureTransitionLegType], and [models.ProcedureAvailabilityType].
const Note02SIDSTAR = ""

// GONGNOTE(Note03IAP): 03. Instrument Approach Procedure (IAP) Story
//
// An [models.InstrumentApproachProcedureTimeSliceType] defines precision or non-precision approach profiles ([models.FinalProfileType]),
// SBAS/GBAS [models.FASDataBlockType], landing [models.ApproachConditionType] and [models.MinimaType] (OCA/H, DA/H), and [models.MissedApproachGroupType].
const Note03IAP = ""

// GONGNOTE(Note04NavigationArea): 04. NavigationArea, TAA & CirclingArea Story
//
// [models.NavigationAreaTimeSliceType], [models.TerminalArrivalAreaTimeSliceType] (TAA), and [models.CirclingAreaTimeSliceType] define obstacle clearance
// sectors ([models.NavigationAreaSectorType], [models.TerminalArrivalAreaSectorType]) and PANS-OPS [models.SectorDesignType] criteria.
const Note04NavigationArea = ""

// GONGNOTE(Note05DepartureArrivalLegs): 05. Departure & Arrival Segment Legs Story
//
// Individual ARINC 424 path-terminator legs for SIDs and STARs are modeled by [models.DepartureLegTimeSliceType], [models.ArrivalLegTimeSliceType],
// and [models.ArrivalFeederLegTimeSliceType], connecting [models.TerminalSegmentPointType] fixes and [models.LandingTakeoffAreaCollectionType].
const Note05DepartureArrivalLegs = ""

// GONGNOTE(Note06ApproachLegs): 06. Approach & Missed Approach Segment Legs Story
//
// Approach transitions are decomposed into [models.InitialLegTimeSliceType] (IAF), [models.IntermediateLegTimeSliceType] (IF),
// [models.FinalLegTimeSliceType] (FAF/FAP), and [models.MissedApproachLegTimeSliceType] (MAPt), referencing [models.TerminalSegmentPointType] and [models.HoldingUseType].
const Note06ApproachLegs = ""

// GONGNOTE(Note07SafeAltitude): 07. SafeAltitudeArea Story
//
// [models.SafeAltitudeAreaTimeSliceType] models Minimum Sector Altitudes (MSA) or Emergency Safe Altitudes (ESA) around a reference navaid/waypoint,
// partitioned into radial [models.SafeAltitudeAreaSectorType] pie slices with [models.CircleSectorType] geometry.
const Note07SafeAltitude = ""

// GONGNOTE(Note08HoldingPattern): 08. HoldingPattern & UnplannedHolding Story
//
// [models.HoldingPatternTimeSliceType] and [models.UnplannedHoldingTimeSliceType] specify racetrack holding fixes ([models.TerminalSegmentPointType]),
// inbound course, turn direction, and leg length by distance ([models.HoldingPatternDistanceType]) or time ([models.HoldingPatternDurationType]).
const Note08HoldingPattern = ""

// GONGNOTE(Note09HoldingAssessment): 09. HoldingAssessment & Obstacle Assessment Story
//
// [models.HoldingAssessmentTimeSliceType] documents obstacle clearance evaluation inside a holding area using [models.ObstacleAssessmentAreaType],
// controlling [models.ObstructionType] records, [models.ObstaclePlacementType], and [models.AltitudeAdjustmentType].
const Note09HoldingAssessment = ""

