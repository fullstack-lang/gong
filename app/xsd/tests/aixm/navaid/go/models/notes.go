package models

// GONGNOTE(Note01NavaidMessage): 01. Navaid & Point Message & Temporality Story
//
// An [models.AIXMBasicMessage] exchanges radio navigation aids and waypoints via [models.BasicMessageMemberAIXMPropertyType].
// Each feature ([models.NavaidType], [models.DesignatedPointType]) is versioned over time by [models.NavaidTimeSliceType]
// and [models.DesignatedPointTimeSliceType] slices bounded by ValidTime ([models.TimePrimitivePropertyType]).
const Note01NavaidMessage = ""

// GONGNOTE(Note02NavaidComponent): 02. Navaid, Component & Operational Status Story
//
// A [models.NavaidTimeSliceType] represents a composite navigation system (such as a VOR/DME, VORTAC, or ILS installation).
// Physical equipment is linked via [models.NavaidComponentType], while availability and signal monitoring use
// [models.NavaidOperationalStatusType], [models.NavaidEquipmentMonitoringType], and [models.AuthorityForNavaidEquipmentType].
const Note02NavaidComponent = ""

// GONGNOTE(Note03RadioNavaids): 03. En-Route Radio Navaids (VOR, DME, TACAN, NDB) Story
//
// Individual radio transmitters ([models.VORTimeSliceType], [models.DMETimeSliceType], [models.TACANTimeSliceType], and [models.NDBTimeSliceType])
// carry frequency/channel parameters, magnetic variation, 3D antenna coordinates ([models.ElevatedPointType]), and [models.NavaidOperationalStatusType].
const Note03RadioNavaids = ""

// GONGNOTE(Note04ILSMLS): 04. ILS & MLS Precision Approach Navaids Story
//
// Precision landing guidance comprises [models.LocalizerTimeSliceType] (lateral course), [models.GlidepathTimeSliceType] (vertical descent slope),
// and [models.MarkerBeaconTimeSliceType] (outer/middle/inner markers), each reporting serviceability via [models.NavaidOperationalStatusType].
const Note04ILSMLS = ""

// GONGNOTE(Note05Checkpoints): 05. Navigation Checkpoints (VOR & INS) Story
//
// [models.CheckpointVORTimeSliceType] and [models.CheckpointINSTimeSliceType] specify designated aerodrome or airborne locations
// ([models.ElevatedPointType]) where pilots can verify VOR receiver bearing accuracy or align inertial navigation systems.
const Note05Checkpoints = ""

// GONGNOTE(Note06SpecialNav): 06. Special Navigation System & Station Story
//
// Hyperbolic or long-range navigation chains (LORAN-C, GNSS ground augmentation) are modeled by [models.SpecialNavigationSystemTimeSliceType]
// and individual transmitter stations ([models.SpecialNavigationStationTimeSliceType]) with [models.SpecialNavigationStationStatusType].
const Note06SpecialNav = ""

// GONGNOTE(Note07DesignatedPoint): 07. DesignatedPoint & SignificantPointInAirspace Story
//
// [models.DesignatedPointTimeSliceType] defines published 5-letter ICAO waypoints and reporting points, while [models.SignificantPointInAirspaceTimeSliceType]
// associates them with airspace entry/exit roles and segment references ([models.EnRouteSegmentPointType], [models.PointReferenceType]).
const Note07DesignatedPoint = ""

// GONGNOTE(Note08AngleDistance): 08. AngleIndication & DistanceIndication Story
//
// [models.AngleIndicationTimeSliceType] and [models.DistanceIndicationTimeSliceType] define published radials/bearings and DME distances
// fixing a waypoint or navaid relative to a reference station via [models.AngleUseType] and [models.PointReferenceType].
const Note08AngleDistance = ""

// GONGNOTE(Note09RadarSystems): 09. Surveillance Radar Systems (PSR, SSR, PAR) Story
//
// [models.RadarSystemTimeSliceType] aggregates [models.PrimarySurveillanceRadarTimeSliceType], [models.SecondarySurveillanceRadarTimeSliceType],
// and [models.PrecisionApproachRadarTimeSliceType] heads via [models.RadarComponentType], along with [models.SurveillanceGroundStationType] and [models.ReflectorType].
const Note09RadarSystems = ""

