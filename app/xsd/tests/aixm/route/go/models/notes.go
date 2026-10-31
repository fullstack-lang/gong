package models

// GONGNOTE(Note01RouteMessage): 01. Route Message & Temporality Story
//
// An [models.AIXMBasicMessage] packages en-route ATS/RNAV routes and segments via [models.BasicMessageMemberAIXMPropertyType].
// Each feature ([models.RouteType], [models.RouteSegmentType]) is versioned through [models.RouteTimeSliceType] and
// [models.RouteSegmentTimeSliceType] slices whose validity interval is specified by ValidTime ([models.TimePrimitivePropertyType]).
const Note01RouteMessage = ""

// GONGNOTE(Note02RoutePortion): 02. Route & RoutePortion Story
//
// A [models.RouteTimeSliceType] identifies an ATS or RNAV airway (designator prefix, second letter, number, and navigation type).
// Sub-sections of a route are referenced via [models.RoutePortionType] and documented with [models.NoteType] and [models.LinguisticNoteType].
const Note02RoutePortion = ""

// GONGNOTE(Note03RouteSegment): 03. RouteSegment & RouteAvailability Story
//
// A [models.RouteSegmentTimeSliceType] connects start and end [models.EnRouteSegmentPointType] fixes along a 2D [models.CurveType] path.
// Directional and altitude-band opening/closure schedules are defined by [models.RouteAvailabilityType] and [models.AirspaceLayerType].
const Note03RouteSegment = ""

// GONGNOTE(Note04ChangeOverPoint): 04. RouteDME & ChangeOverPoint Story
//
// [models.RouteDMETimeSliceType] associates DME distance references with an airway segment, while [models.ChangeOverPointTimeSliceType]
// specifies the point along a segment where aircraft transition navigation guidance from the behind facility to the ahead facility via [models.PointReferenceType].
const Note04ChangeOverPoint = ""

// GONGNOTE(Note05FlightRestriction): 05. FlightRestriction & Routing Elements Story
//
// Traffic flow restrictions (RAD / conditional routing rules) are modeled on [models.FlightRestrictionTimeSliceType].
// Allowed or forbidden routings are sequences of [models.FlightRestrictionRouteType], [models.FlightRoutingElementType], and [models.FlightRestrictionLevelType].
const Note05FlightRestriction = ""

// GONGNOTE(Note06FlightConditions): 06. Flight Conditions, DirectFlight & RulesProcedures Story
//
// [models.AirspaceBorderCrossingTimeSliceType] and [models.RulesProceduresTimeSliceType] capture border-crossing constraints and regional rules
// using boolean logic trees of [models.FlightConditionCombinationType], [models.FlightConditionElementType], and [models.DirectFlightSegmentType].
const Note06FlightConditions = ""

// GONGNOTE(Note07AerialRefuelling): 07. Aerial Refuelling Story
//
// [models.AerialRefuellingTimeSliceType] describes military air-to-air refuelling tracks ([models.AerialRefuellingTrackType])
// and racetrack anchors ([models.AerialRefuellingAnchorType]) composed of ordered [models.AerialRefuellingPointType] waypoints and [models.AuthorityForAerialRefuellingType].
const Note07AerialRefuelling = ""

