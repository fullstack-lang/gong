package models

// GONGNOTE(Note01ObstacleMessage): 01. Obstacle Message & Temporality Story
//
// An [models.AIXMBasicMessage] exchanges obstacle datasets via [models.BasicMessageMemberAIXMPropertyType].
// Each obstacle feature ([models.VerticalStructureType], [models.ObstacleAreaType]) follows the AIXM Temporality Model:
// its lifecycle is represented by [models.VerticalStructureTimeSliceType] slices bounded by ValidTime ([models.TimePrimitivePropertyType]).
const Note01ObstacleMessage = ""

// GONGNOTE(Note02VerticalStructure): 02. VerticalStructure & Parts Story
//
// A [models.VerticalStructureTimeSliceType] models man-made or natural obstacles (towers, buildings, wind turbines, cranes).
// Each structure is decomposed into [models.VerticalStructurePartType] components (point, line, or polygon geometry)
// with operational lighting tracked by [models.VerticalStructureLightingStatusType] and [models.LightElementType].
const Note02VerticalStructure = ""

// GONGNOTE(Note03ObstacleArea): 03. ObstacleArea Story
//
// An [models.ObstacleAreaTimeSliceType] defines ICAO Annex 15 eTOD obstacle collection surfaces (Area 1, Area 2, Area 3, Area 4).
// Its horizontal boundary is modeled by [models.SurfaceType] or 3D [models.ElevatedSurfaceType] along with [models.NoteType] annotations.
const Note03ObstacleArea = ""

// GONGNOTE(Note04GroundLight): 04. AeronauticalGroundLight Story
//
// An [models.AeronauticalGroundLightTimeSliceType] describes visual beacons (aerodrome beacons, hazard beacons, marine lighthouses)
// located at a 3D [models.ElevatedPointType] and equipped with [models.LightElementType] and [models.LightElementStatusType].
const Note04GroundLight = ""

// GONGNOTE(Note05ElevatedGeometry): 05. Elevated 3D Geometry (Point, Curve, Surface) Story
//
// AIXM extends GML 3.2.1 geometries into [models.ElevatedPointType], [models.ElevatedCurveType], and [models.ElevatedSurfaceType]
// to add vertical elevation, vertical datum, and vertical accuracy on top of [models.GeodesicStringType], [models.PolygonPatchType], and [models.DirectPositionListType].
const Note05ElevatedGeometry = ""

