package models

// GONGNOTE(Note01AirspaceMessage): 01. Airspace Message & Temporality Story
//
// An [models.AIXMBasicMessage] is the root exchange container that groups aeronautical features via [models.BasicMessageMemberAIXMPropertyType].
// Every AIXM feature ([models.AirspaceType], [models.GeoBorderType]) follows the Temporality Model: its state over time is represented
// by [models.AirspaceTimeSliceType] slices (BASELINE, PERMDELTA, TEMPDELTA, SNAPSHOT) bounded by ValidTime ([models.TimePrimitivePropertyType]).
const Note01AirspaceMessage = ""

// GONGNOTE(Note02AirspaceVolumes): 02. Airspace & Geometry Volumes Story
//
// An [models.AirspaceTimeSliceType] defines its 3D spatial extent through one or more [models.AirspaceGeometryComponentType] blocks.
// Each component references an [models.AirspaceVolumeType] bounded vertically by upper/lower limits and horizontally by a
// projected [models.SurfaceType] or derived from another airspace via [models.AirspaceVolumeDependencyType].
const Note02AirspaceVolumes = ""

// GONGNOTE(Note03AirspaceActivation): 03. Airspace Activation & Layer Classification Story
//
// Dynamic airspace status is modeled on [models.AirspaceTimeSliceType] via [models.AirspaceActivationType] (status, activity, levels, and [models.TimesheetType] schedules).
// ICAO airspace class (A through G) per altitude band is defined by [models.AirspaceLayerClassType] and vertical [models.AirspaceLayerType] slices.
const Note03AirspaceActivation = ""

// GONGNOTE(Note04GeoBorder): 04. AuthorityForAirspace & GeoBorder Story
//
// [models.AuthorityForAirspaceTimeSliceType] assigns responsible organizations (civil/military jurisdiction) to an airspace.
// [models.GeoBorderTimeSliceType] models political coastlines or state borders using 2D [models.CurveType] and 3D [models.ElevatedCurveType] geometries.
const Note04GeoBorder = ""

// GONGNOTE(Note05GMLGeometry): 05. Airspace GML 3.2.1 Surface & Curve Segments Story
//
// AIXM [models.SurfaceType] and [models.CurveType] extend ISO 19107 / GML 3.2.1 geometry primitives.
// A surface is composed of [models.PolygonPatchType] bounded by a [models.LinearRingType], whose segments are built from
// [models.GeodesicStringType], [models.ArcByCenterPointType], or [models.CircleByCenterPointType] coordinate lists ([models.DirectPositionListType]).
const Note05GMLGeometry = ""

