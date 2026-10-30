// generated code - do not edit
package aixmfeatures

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

var (
	__GongSliceTemplate_time__dummyDeclaration time.Duration
	_                                          = __GongSliceTemplate_time__dummyDeclaration
)

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct AircraftCharacteristicType
	// insertion point per field
	stage.AircraftCharacteristicType_Annotation_reverseMap = make(map[*NotePropertyType]*AircraftCharacteristicType)
	for aircraftcharacteristictype := range stage.AircraftCharacteristicTypes {
		_ = aircraftcharacteristictype
		for _, _notepropertytype := range aircraftcharacteristictype.Annotation {
			stage.AircraftCharacteristicType_Annotation_reverseMap[_notepropertytype] = aircraftcharacteristictype
		}
	}
	stage.AircraftCharacteristicType_Extension_reverseMap = make(map[*A_extension_12]*AircraftCharacteristicType)
	for aircraftcharacteristictype := range stage.AircraftCharacteristicTypes {
		_ = aircraftcharacteristictype
		for _, _a_extension_12 := range aircraftcharacteristictype.Extension {
			stage.AircraftCharacteristicType_Extension_reverseMap[_a_extension_12] = aircraftcharacteristictype
		}
	}

	// Compute reverse map for named struct AirspaceActivationType
	// insertion point per field
	stage.AirspaceActivationType_TimeInterval_reverseMap = make(map[*TimesheetPropertyType]*AirspaceActivationType)
	for airspaceactivationtype := range stage.AirspaceActivationTypes {
		_ = airspaceactivationtype
		for _, _timesheetpropertytype := range airspaceactivationtype.TimeInterval {
			stage.AirspaceActivationType_TimeInterval_reverseMap[_timesheetpropertytype] = airspaceactivationtype
		}
	}
	stage.AirspaceActivationType_Annotation_reverseMap = make(map[*NotePropertyType]*AirspaceActivationType)
	for airspaceactivationtype := range stage.AirspaceActivationTypes {
		_ = airspaceactivationtype
		for _, _notepropertytype := range airspaceactivationtype.Annotation {
			stage.AirspaceActivationType_Annotation_reverseMap[_notepropertytype] = airspaceactivationtype
		}
	}
	stage.AirspaceActivationType_SpecialDateAuthority_reverseMap = make(map[*OrganisationAuthorityPropertyType]*AirspaceActivationType)
	for airspaceactivationtype := range stage.AirspaceActivationTypes {
		_ = airspaceactivationtype
		for _, _organisationauthoritypropertytype := range airspaceactivationtype.SpecialDateAuthority {
			stage.AirspaceActivationType_SpecialDateAuthority_reverseMap[_organisationauthoritypropertytype] = airspaceactivationtype
		}
	}
	stage.AirspaceActivationType_Levels_reverseMap = make(map[*AirspaceLayerPropertyType]*AirspaceActivationType)
	for airspaceactivationtype := range stage.AirspaceActivationTypes {
		_ = airspaceactivationtype
		for _, _airspacelayerpropertytype := range airspaceactivationtype.Levels {
			stage.AirspaceActivationType_Levels_reverseMap[_airspacelayerpropertytype] = airspaceactivationtype
		}
	}
	stage.AirspaceActivationType_User_reverseMap = make(map[*OrganisationAuthorityPropertyType]*AirspaceActivationType)
	for airspaceactivationtype := range stage.AirspaceActivationTypes {
		_ = airspaceactivationtype
		for _, _organisationauthoritypropertytype := range airspaceactivationtype.User {
			stage.AirspaceActivationType_User_reverseMap[_organisationauthoritypropertytype] = airspaceactivationtype
		}
	}
	stage.AirspaceActivationType_Aircraft_reverseMap = make(map[*AircraftCharacteristicPropertyType]*AirspaceActivationType)
	for airspaceactivationtype := range stage.AirspaceActivationTypes {
		_ = airspaceactivationtype
		for _, _aircraftcharacteristicpropertytype := range airspaceactivationtype.Aircraft {
			stage.AirspaceActivationType_Aircraft_reverseMap[_aircraftcharacteristicpropertytype] = airspaceactivationtype
		}
	}
	stage.AirspaceActivationType_Extension_reverseMap = make(map[*A_extension_1]*AirspaceActivationType)
	for airspaceactivationtype := range stage.AirspaceActivationTypes {
		_ = airspaceactivationtype
		for _, _a_extension_1 := range airspaceactivationtype.Extension {
			stage.AirspaceActivationType_Extension_reverseMap[_a_extension_1] = airspaceactivationtype
		}
	}

	// Compute reverse map for named struct AirspaceGeometryComponentType
	// insertion point per field
	stage.AirspaceGeometryComponentType_Annotation_reverseMap = make(map[*NotePropertyType]*AirspaceGeometryComponentType)
	for airspacegeometrycomponenttype := range stage.AirspaceGeometryComponentTypes {
		_ = airspacegeometrycomponenttype
		for _, _notepropertytype := range airspacegeometrycomponenttype.Annotation {
			stage.AirspaceGeometryComponentType_Annotation_reverseMap[_notepropertytype] = airspacegeometrycomponenttype
		}
	}
	stage.AirspaceGeometryComponentType_Extension_reverseMap = make(map[*A_extension_2]*AirspaceGeometryComponentType)
	for airspacegeometrycomponenttype := range stage.AirspaceGeometryComponentTypes {
		_ = airspacegeometrycomponenttype
		for _, _a_extension_2 := range airspacegeometrycomponenttype.Extension {
			stage.AirspaceGeometryComponentType_Extension_reverseMap[_a_extension_2] = airspacegeometrycomponenttype
		}
	}

	// Compute reverse map for named struct AirspaceLayerClassType
	// insertion point per field
	stage.AirspaceLayerClassType_TimeInterval_reverseMap = make(map[*TimesheetPropertyType]*AirspaceLayerClassType)
	for airspacelayerclasstype := range stage.AirspaceLayerClassTypes {
		_ = airspacelayerclasstype
		for _, _timesheetpropertytype := range airspacelayerclasstype.TimeInterval {
			stage.AirspaceLayerClassType_TimeInterval_reverseMap[_timesheetpropertytype] = airspacelayerclasstype
		}
	}
	stage.AirspaceLayerClassType_Annotation_reverseMap = make(map[*NotePropertyType]*AirspaceLayerClassType)
	for airspacelayerclasstype := range stage.AirspaceLayerClassTypes {
		_ = airspacelayerclasstype
		for _, _notepropertytype := range airspacelayerclasstype.Annotation {
			stage.AirspaceLayerClassType_Annotation_reverseMap[_notepropertytype] = airspacelayerclasstype
		}
	}
	stage.AirspaceLayerClassType_SpecialDateAuthority_reverseMap = make(map[*OrganisationAuthorityPropertyType]*AirspaceLayerClassType)
	for airspacelayerclasstype := range stage.AirspaceLayerClassTypes {
		_ = airspacelayerclasstype
		for _, _organisationauthoritypropertytype := range airspacelayerclasstype.SpecialDateAuthority {
			stage.AirspaceLayerClassType_SpecialDateAuthority_reverseMap[_organisationauthoritypropertytype] = airspacelayerclasstype
		}
	}
	stage.AirspaceLayerClassType_AssociatedLevels_reverseMap = make(map[*AirspaceLayerPropertyType]*AirspaceLayerClassType)
	for airspacelayerclasstype := range stage.AirspaceLayerClassTypes {
		_ = airspacelayerclasstype
		for _, _airspacelayerpropertytype := range airspacelayerclasstype.AssociatedLevels {
			stage.AirspaceLayerClassType_AssociatedLevels_reverseMap[_airspacelayerpropertytype] = airspacelayerclasstype
		}
	}
	stage.AirspaceLayerClassType_Extension_reverseMap = make(map[*A_extension_3]*AirspaceLayerClassType)
	for airspacelayerclasstype := range stage.AirspaceLayerClassTypes {
		_ = airspacelayerclasstype
		for _, _a_extension_3 := range airspacelayerclasstype.Extension {
			stage.AirspaceLayerClassType_Extension_reverseMap[_a_extension_3] = airspacelayerclasstype
		}
	}

	// Compute reverse map for named struct AirspaceLayerType
	// insertion point per field
	stage.AirspaceLayerType_Annotation_reverseMap = make(map[*NotePropertyType]*AirspaceLayerType)
	for airspacelayertype := range stage.AirspaceLayerTypes {
		_ = airspacelayertype
		for _, _notepropertytype := range airspacelayertype.Annotation {
			stage.AirspaceLayerType_Annotation_reverseMap[_notepropertytype] = airspacelayertype
		}
	}
	stage.AirspaceLayerType_Extension_reverseMap = make(map[*A_extension_13]*AirspaceLayerType)
	for airspacelayertype := range stage.AirspaceLayerTypes {
		_ = airspacelayertype
		for _, _a_extension_13 := range airspacelayertype.Extension {
			stage.AirspaceLayerType_Extension_reverseMap[_a_extension_13] = airspacelayertype
		}
	}

	// Compute reverse map for named struct AirspaceTimeSliceType
	// insertion point per field
	stage.AirspaceTimeSliceType_Class_reverseMap = make(map[*AirspaceLayerClassPropertyType]*AirspaceTimeSliceType)
	for airspacetimeslicetype := range stage.AirspaceTimeSliceTypes {
		_ = airspacetimeslicetype
		for _, _airspacelayerclasspropertytype := range airspacetimeslicetype.Class {
			stage.AirspaceTimeSliceType_Class_reverseMap[_airspacelayerclasspropertytype] = airspacetimeslicetype
		}
	}
	stage.AirspaceTimeSliceType_GeometryComponent_reverseMap = make(map[*AirspaceGeometryComponentPropertyType]*AirspaceTimeSliceType)
	for airspacetimeslicetype := range stage.AirspaceTimeSliceTypes {
		_ = airspacetimeslicetype
		for _, _airspacegeometrycomponentpropertytype := range airspacetimeslicetype.GeometryComponent {
			stage.AirspaceTimeSliceType_GeometryComponent_reverseMap[_airspacegeometrycomponentpropertytype] = airspacetimeslicetype
		}
	}
	stage.AirspaceTimeSliceType_Activation_reverseMap = make(map[*AirspaceActivationPropertyType]*AirspaceTimeSliceType)
	for airspacetimeslicetype := range stage.AirspaceTimeSliceTypes {
		_ = airspacetimeslicetype
		for _, _airspaceactivationpropertytype := range airspacetimeslicetype.Activation {
			stage.AirspaceTimeSliceType_Activation_reverseMap[_airspaceactivationpropertytype] = airspacetimeslicetype
		}
	}
	stage.AirspaceTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*AirspaceTimeSliceType)
	for airspacetimeslicetype := range stage.AirspaceTimeSliceTypes {
		_ = airspacetimeslicetype
		for _, _notepropertytype := range airspacetimeslicetype.Annotation {
			stage.AirspaceTimeSliceType_Annotation_reverseMap[_notepropertytype] = airspacetimeslicetype
		}
	}
	stage.AirspaceTimeSliceType_Extension_reverseMap = make(map[*A_extension]*AirspaceTimeSliceType)
	for airspacetimeslicetype := range stage.AirspaceTimeSliceTypes {
		_ = airspacetimeslicetype
		for _, _a_extension := range airspacetimeslicetype.Extension {
			stage.AirspaceTimeSliceType_Extension_reverseMap[_a_extension] = airspacetimeslicetype
		}
	}

	// Compute reverse map for named struct AirspaceType
	// insertion point per field
	stage.AirspaceType_TimeSlice_reverseMap = make(map[*AirspaceTimeSlicePropertyType]*AirspaceType)
	for airspacetype := range stage.AirspaceTypes {
		_ = airspacetype
		for _, _airspacetimeslicepropertytype := range airspacetype.TimeSlice {
			stage.AirspaceType_TimeSlice_reverseMap[_airspacetimeslicepropertytype] = airspacetype
		}
	}

	// Compute reverse map for named struct AirspaceVolumeDependencyType
	// insertion point per field
	stage.AirspaceVolumeDependencyType_Annotation_reverseMap = make(map[*NotePropertyType]*AirspaceVolumeDependencyType)
	for airspacevolumedependencytype := range stage.AirspaceVolumeDependencyTypes {
		_ = airspacevolumedependencytype
		for _, _notepropertytype := range airspacevolumedependencytype.Annotation {
			stage.AirspaceVolumeDependencyType_Annotation_reverseMap[_notepropertytype] = airspacevolumedependencytype
		}
	}
	stage.AirspaceVolumeDependencyType_Extension_reverseMap = make(map[*A_extension_5]*AirspaceVolumeDependencyType)
	for airspacevolumedependencytype := range stage.AirspaceVolumeDependencyTypes {
		_ = airspacevolumedependencytype
		for _, _a_extension_5 := range airspacevolumedependencytype.Extension {
			stage.AirspaceVolumeDependencyType_Extension_reverseMap[_a_extension_5] = airspacevolumedependencytype
		}
	}

	// Compute reverse map for named struct AirspaceVolumeType
	// insertion point per field
	stage.AirspaceVolumeType_Annotation_reverseMap = make(map[*NotePropertyType]*AirspaceVolumeType)
	for airspacevolumetype := range stage.AirspaceVolumeTypes {
		_ = airspacevolumetype
		for _, _notepropertytype := range airspacevolumetype.Annotation {
			stage.AirspaceVolumeType_Annotation_reverseMap[_notepropertytype] = airspacevolumetype
		}
	}
	stage.AirspaceVolumeType_Extension_reverseMap = make(map[*A_extension_4]*AirspaceVolumeType)
	for airspacevolumetype := range stage.AirspaceVolumeTypes {
		_ = airspacevolumetype
		for _, _a_extension_4 := range airspacevolumetype.Extension {
			stage.AirspaceVolumeType_Extension_reverseMap[_a_extension_4] = airspacevolumetype
		}
	}

	// Compute reverse map for named struct AuthorityForAirspaceTimeSliceType
	// insertion point per field
	stage.AuthorityForAirspaceTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*AuthorityForAirspaceTimeSliceType)
	for authorityforairspacetimeslicetype := range stage.AuthorityForAirspaceTimeSliceTypes {
		_ = authorityforairspacetimeslicetype
		for _, _notepropertytype := range authorityforairspacetimeslicetype.Annotation {
			stage.AuthorityForAirspaceTimeSliceType_Annotation_reverseMap[_notepropertytype] = authorityforairspacetimeslicetype
		}
	}
	stage.AuthorityForAirspaceTimeSliceType_Extension_reverseMap = make(map[*A_extension_6]*AuthorityForAirspaceTimeSliceType)
	for authorityforairspacetimeslicetype := range stage.AuthorityForAirspaceTimeSliceTypes {
		_ = authorityforairspacetimeslicetype
		for _, _a_extension_6 := range authorityforairspacetimeslicetype.Extension {
			stage.AuthorityForAirspaceTimeSliceType_Extension_reverseMap[_a_extension_6] = authorityforairspacetimeslicetype
		}
	}

	// Compute reverse map for named struct AuthorityForAirspaceType
	// insertion point per field
	stage.AuthorityForAirspaceType_TimeSlice_reverseMap = make(map[*AuthorityForAirspaceTimeSlicePropertyType]*AuthorityForAirspaceType)
	for authorityforairspacetype := range stage.AuthorityForAirspaceTypes {
		_ = authorityforairspacetype
		for _, _authorityforairspacetimeslicepropertytype := range authorityforairspacetype.TimeSlice {
			stage.AuthorityForAirspaceType_TimeSlice_reverseMap[_authorityforairspacetimeslicepropertytype] = authorityforairspacetype
		}
	}

	// Compute reverse map for named struct CurveType
	// insertion point per field
	stage.CurveType_Annotation_reverseMap = make(map[*NotePropertyType]*CurveType)
	for curvetype := range stage.CurveTypes {
		_ = curvetype
		for _, _notepropertytype := range curvetype.Annotation {
			stage.CurveType_Annotation_reverseMap[_notepropertytype] = curvetype
		}
	}

	// Compute reverse map for named struct ElevatedCurveType
	// insertion point per field
	stage.ElevatedCurveType_Annotation_reverseMap = make(map[*NotePropertyType]*ElevatedCurveType)
	for elevatedcurvetype := range stage.ElevatedCurveTypes {
		_ = elevatedcurvetype
		for _, _notepropertytype := range elevatedcurvetype.Annotation {
			stage.ElevatedCurveType_Annotation_reverseMap[_notepropertytype] = elevatedcurvetype
		}
	}
	stage.ElevatedCurveType_Extension_reverseMap = make(map[*A_extension_8]*ElevatedCurveType)
	for elevatedcurvetype := range stage.ElevatedCurveTypes {
		_ = elevatedcurvetype
		for _, _a_extension_8 := range elevatedcurvetype.Extension {
			stage.ElevatedCurveType_Extension_reverseMap[_a_extension_8] = elevatedcurvetype
		}
	}

	// Compute reverse map for named struct ElevatedSurfaceType
	// insertion point per field
	stage.ElevatedSurfaceType_Annotation_reverseMap = make(map[*NotePropertyType]*ElevatedSurfaceType)
	for elevatedsurfacetype := range stage.ElevatedSurfaceTypes {
		_ = elevatedsurfacetype
		for _, _notepropertytype := range elevatedsurfacetype.Annotation {
			stage.ElevatedSurfaceType_Annotation_reverseMap[_notepropertytype] = elevatedsurfacetype
		}
	}
	stage.ElevatedSurfaceType_Extension_reverseMap = make(map[*A_extension_9]*ElevatedSurfaceType)
	for elevatedsurfacetype := range stage.ElevatedSurfaceTypes {
		_ = elevatedsurfacetype
		for _, _a_extension_9 := range elevatedsurfacetype.Extension {
			stage.ElevatedSurfaceType_Extension_reverseMap[_a_extension_9] = elevatedsurfacetype
		}
	}

	// Compute reverse map for named struct GeoBorderTimeSliceType
	// insertion point per field
	stage.GeoBorderTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*GeoBorderTimeSliceType)
	for geobordertimeslicetype := range stage.GeoBorderTimeSliceTypes {
		_ = geobordertimeslicetype
		for _, _notepropertytype := range geobordertimeslicetype.Annotation {
			stage.GeoBorderTimeSliceType_Annotation_reverseMap[_notepropertytype] = geobordertimeslicetype
		}
	}
	stage.GeoBorderTimeSliceType_Extension_reverseMap = make(map[*A_extension_7]*GeoBorderTimeSliceType)
	for geobordertimeslicetype := range stage.GeoBorderTimeSliceTypes {
		_ = geobordertimeslicetype
		for _, _a_extension_7 := range geobordertimeslicetype.Extension {
			stage.GeoBorderTimeSliceType_Extension_reverseMap[_a_extension_7] = geobordertimeslicetype
		}
	}

	// Compute reverse map for named struct GeoBorderType
	// insertion point per field
	stage.GeoBorderType_TimeSlice_reverseMap = make(map[*GeoBorderTimeSlicePropertyType]*GeoBorderType)
	for geobordertype := range stage.GeoBorderTypes {
		_ = geobordertype
		for _, _geobordertimeslicepropertytype := range geobordertype.TimeSlice {
			stage.GeoBorderType_TimeSlice_reverseMap[_geobordertimeslicepropertytype] = geobordertype
		}
	}

	// Compute reverse map for named struct LinguisticNoteType
	// insertion point per field
	stage.LinguisticNoteType_Extension_reverseMap = make(map[*A_extension_10]*LinguisticNoteType)
	for linguisticnotetype := range stage.LinguisticNoteTypes {
		_ = linguisticnotetype
		for _, _a_extension_10 := range linguisticnotetype.Extension {
			stage.LinguisticNoteType_Extension_reverseMap[_a_extension_10] = linguisticnotetype
		}
	}

	// Compute reverse map for named struct NoteType
	// insertion point per field
	stage.NoteType_TranslatedNote_reverseMap = make(map[*LinguisticNotePropertyType]*NoteType)
	for notetype := range stage.NoteTypes {
		_ = notetype
		for _, _linguisticnotepropertytype := range notetype.TranslatedNote {
			stage.NoteType_TranslatedNote_reverseMap[_linguisticnotepropertytype] = notetype
		}
	}
	stage.NoteType_Extension_reverseMap = make(map[*A_extension_11]*NoteType)
	for notetype := range stage.NoteTypes {
		_ = notetype
		for _, _a_extension_11 := range notetype.Extension {
			stage.NoteType_Extension_reverseMap[_a_extension_11] = notetype
		}
	}

	// Compute reverse map for named struct SurfaceType
	// insertion point per field
	stage.SurfaceType_Annotation_reverseMap = make(map[*NotePropertyType]*SurfaceType)
	for surfacetype := range stage.SurfaceTypes {
		_ = surfacetype
		for _, _notepropertytype := range surfacetype.Annotation {
			stage.SurfaceType_Annotation_reverseMap[_notepropertytype] = surfacetype
		}
	}

	// Compute reverse map for named struct TimesheetType
	// insertion point per field
	stage.TimesheetType_Annotation_reverseMap = make(map[*NotePropertyType]*TimesheetType)
	for timesheettype := range stage.TimesheetTypes {
		_ = timesheettype
		for _, _notepropertytype := range timesheettype.Annotation {
			stage.TimesheetType_Annotation_reverseMap[_notepropertytype] = timesheettype
		}
	}
	stage.TimesheetType_Extension_reverseMap = make(map[*A_extension_14]*TimesheetType)
	for timesheettype := range stage.TimesheetTypes {
		_ = timesheettype
		for _, _a_extension_14 := range timesheettype.Extension {
			stage.TimesheetType_Extension_reverseMap[_a_extension_14] = timesheettype
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.A_extensions)

	res = __gong__appendInstances(res, stage.A_extension_1s)

	res = __gong__appendInstances(res, stage.A_extension_10s)

	res = __gong__appendInstances(res, stage.A_extension_11s)

	res = __gong__appendInstances(res, stage.A_extension_12s)

	res = __gong__appendInstances(res, stage.A_extension_13s)

	res = __gong__appendInstances(res, stage.A_extension_14s)

	res = __gong__appendInstances(res, stage.A_extension_2s)

	res = __gong__appendInstances(res, stage.A_extension_3s)

	res = __gong__appendInstances(res, stage.A_extension_4s)

	res = __gong__appendInstances(res, stage.A_extension_5s)

	res = __gong__appendInstances(res, stage.A_extension_6s)

	res = __gong__appendInstances(res, stage.A_extension_7s)

	res = __gong__appendInstances(res, stage.A_extension_8s)

	res = __gong__appendInstances(res, stage.A_extension_9s)

	res = __gong__appendInstances(res, stage.AbstractPropertiesWithScheduleTypes)

	res = __gong__appendInstances(res, stage.AircraftCharacteristicPropertyTypes)

	res = __gong__appendInstances(res, stage.AircraftCharacteristicTypes)

	res = __gong__appendInstances(res, stage.AirspaceActivationPropertyTypes)

	res = __gong__appendInstances(res, stage.AirspaceActivationTypes)

	res = __gong__appendInstances(res, stage.AirspaceGeometryComponentPropertyTypes)

	res = __gong__appendInstances(res, stage.AirspaceGeometryComponentTypes)

	res = __gong__appendInstances(res, stage.AirspaceLayerClassPropertyTypes)

	res = __gong__appendInstances(res, stage.AirspaceLayerClassTypes)

	res = __gong__appendInstances(res, stage.AirspaceLayerPropertyTypes)

	res = __gong__appendInstances(res, stage.AirspaceLayerTypes)

	res = __gong__appendInstances(res, stage.AirspacePropertyTypes)

	res = __gong__appendInstances(res, stage.AirspaceTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.AirspaceTimeSliceTypes)

	res = __gong__appendInstances(res, stage.AirspaceTypes)

	res = __gong__appendInstances(res, stage.AirspaceVolumeDependencyPropertyTypes)

	res = __gong__appendInstances(res, stage.AirspaceVolumeDependencyTypes)

	res = __gong__appendInstances(res, stage.AirspaceVolumePropertyTypes)

	res = __gong__appendInstances(res, stage.AirspaceVolumeTypes)

	res = __gong__appendInstances(res, stage.AuthorityForAirspaceTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.AuthorityForAirspaceTimeSliceTypes)

	res = __gong__appendInstances(res, stage.AuthorityForAirspaceTypes)

	res = __gong__appendInstances(res, stage.CurvePropertyTypes)

	res = __gong__appendInstances(res, stage.CurveTypes)

	res = __gong__appendInstances(res, stage.ElevatedCurveTypes)

	res = __gong__appendInstances(res, stage.ElevatedSurfaceTypes)

	res = __gong__appendInstances(res, stage.GeoBorderTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.GeoBorderTimeSliceTypes)

	res = __gong__appendInstances(res, stage.GeoBorderTypes)

	res = __gong__appendInstances(res, stage.LinguisticNotePropertyTypes)

	res = __gong__appendInstances(res, stage.LinguisticNoteTypes)

	res = __gong__appendInstances(res, stage.NotePropertyTypes)

	res = __gong__appendInstances(res, stage.NoteTypes)

	res = __gong__appendInstances(res, stage.OrganisationAuthorityPropertyTypes)

	res = __gong__appendInstances(res, stage.RoutePropertyTypes)

	res = __gong__appendInstances(res, stage.StandardLevelColumnPropertyTypes)

	res = __gong__appendInstances(res, stage.SurfacePropertyTypes)

	res = __gong__appendInstances(res, stage.SurfaceTypes)

	res = __gong__appendInstances(res, stage.TimesheetPropertyTypes)

	res = __gong__appendInstances(res, stage.TimesheetTypes)

	return
}

// insertion point per named struct
func (a_extension *A_extension) GongCopy() GongstructIF {
	newInstance := new(A_extension)
	a_extension.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_1 *A_extension_1) GongCopy() GongstructIF {
	newInstance := new(A_extension_1)
	a_extension_1.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_10 *A_extension_10) GongCopy() GongstructIF {
	newInstance := new(A_extension_10)
	a_extension_10.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_11 *A_extension_11) GongCopy() GongstructIF {
	newInstance := new(A_extension_11)
	a_extension_11.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_12 *A_extension_12) GongCopy() GongstructIF {
	newInstance := new(A_extension_12)
	a_extension_12.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_13 *A_extension_13) GongCopy() GongstructIF {
	newInstance := new(A_extension_13)
	a_extension_13.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_14 *A_extension_14) GongCopy() GongstructIF {
	newInstance := new(A_extension_14)
	a_extension_14.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_2 *A_extension_2) GongCopy() GongstructIF {
	newInstance := new(A_extension_2)
	a_extension_2.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_3 *A_extension_3) GongCopy() GongstructIF {
	newInstance := new(A_extension_3)
	a_extension_3.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_4 *A_extension_4) GongCopy() GongstructIF {
	newInstance := new(A_extension_4)
	a_extension_4.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_5 *A_extension_5) GongCopy() GongstructIF {
	newInstance := new(A_extension_5)
	a_extension_5.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_6 *A_extension_6) GongCopy() GongstructIF {
	newInstance := new(A_extension_6)
	a_extension_6.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_7 *A_extension_7) GongCopy() GongstructIF {
	newInstance := new(A_extension_7)
	a_extension_7.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_8 *A_extension_8) GongCopy() GongstructIF {
	newInstance := new(A_extension_8)
	a_extension_8.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_9 *A_extension_9) GongCopy() GongstructIF {
	newInstance := new(A_extension_9)
	a_extension_9.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongCopy() GongstructIF {
	newInstance := new(AbstractPropertiesWithScheduleType)
	abstractpropertieswithscheduletype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongCopy() GongstructIF {
	newInstance := new(AircraftCharacteristicPropertyType)
	aircraftcharacteristicpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (aircraftcharacteristictype *AircraftCharacteristicType) GongCopy() GongstructIF {
	newInstance := new(AircraftCharacteristicType)
	aircraftcharacteristictype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airspaceactivationpropertytype *AirspaceActivationPropertyType) GongCopy() GongstructIF {
	newInstance := new(AirspaceActivationPropertyType)
	airspaceactivationpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airspaceactivationtype *AirspaceActivationType) GongCopy() GongstructIF {
	newInstance := new(AirspaceActivationType)
	airspaceactivationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airspacegeometrycomponentpropertytype *AirspaceGeometryComponentPropertyType) GongCopy() GongstructIF {
	newInstance := new(AirspaceGeometryComponentPropertyType)
	airspacegeometrycomponentpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airspacegeometrycomponenttype *AirspaceGeometryComponentType) GongCopy() GongstructIF {
	newInstance := new(AirspaceGeometryComponentType)
	airspacegeometrycomponenttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airspacelayerclasspropertytype *AirspaceLayerClassPropertyType) GongCopy() GongstructIF {
	newInstance := new(AirspaceLayerClassPropertyType)
	airspacelayerclasspropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airspacelayerclasstype *AirspaceLayerClassType) GongCopy() GongstructIF {
	newInstance := new(AirspaceLayerClassType)
	airspacelayerclasstype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airspacelayerpropertytype *AirspaceLayerPropertyType) GongCopy() GongstructIF {
	newInstance := new(AirspaceLayerPropertyType)
	airspacelayerpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airspacelayertype *AirspaceLayerType) GongCopy() GongstructIF {
	newInstance := new(AirspaceLayerType)
	airspacelayertype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airspacepropertytype *AirspacePropertyType) GongCopy() GongstructIF {
	newInstance := new(AirspacePropertyType)
	airspacepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airspacetimeslicepropertytype *AirspaceTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(AirspaceTimeSlicePropertyType)
	airspacetimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airspacetimeslicetype *AirspaceTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(AirspaceTimeSliceType)
	airspacetimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airspacetype *AirspaceType) GongCopy() GongstructIF {
	newInstance := new(AirspaceType)
	airspacetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airspacevolumedependencypropertytype *AirspaceVolumeDependencyPropertyType) GongCopy() GongstructIF {
	newInstance := new(AirspaceVolumeDependencyPropertyType)
	airspacevolumedependencypropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airspacevolumedependencytype *AirspaceVolumeDependencyType) GongCopy() GongstructIF {
	newInstance := new(AirspaceVolumeDependencyType)
	airspacevolumedependencytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airspacevolumepropertytype *AirspaceVolumePropertyType) GongCopy() GongstructIF {
	newInstance := new(AirspaceVolumePropertyType)
	airspacevolumepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airspacevolumetype *AirspaceVolumeType) GongCopy() GongstructIF {
	newInstance := new(AirspaceVolumeType)
	airspacevolumetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (authorityforairspacetimeslicepropertytype *AuthorityForAirspaceTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(AuthorityForAirspaceTimeSlicePropertyType)
	authorityforairspacetimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (authorityforairspacetimeslicetype *AuthorityForAirspaceTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(AuthorityForAirspaceTimeSliceType)
	authorityforairspacetimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (authorityforairspacetype *AuthorityForAirspaceType) GongCopy() GongstructIF {
	newInstance := new(AuthorityForAirspaceType)
	authorityforairspacetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (curvepropertytype *CurvePropertyType) GongCopy() GongstructIF {
	newInstance := new(CurvePropertyType)
	curvepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (curvetype *CurveType) GongCopy() GongstructIF {
	newInstance := new(CurveType)
	curvetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (elevatedcurvetype *ElevatedCurveType) GongCopy() GongstructIF {
	newInstance := new(ElevatedCurveType)
	elevatedcurvetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongCopy() GongstructIF {
	newInstance := new(ElevatedSurfaceType)
	elevatedsurfacetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (geobordertimeslicepropertytype *GeoBorderTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(GeoBorderTimeSlicePropertyType)
	geobordertimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (geobordertimeslicetype *GeoBorderTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(GeoBorderTimeSliceType)
	geobordertimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (geobordertype *GeoBorderType) GongCopy() GongstructIF {
	newInstance := new(GeoBorderType)
	geobordertype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongCopy() GongstructIF {
	newInstance := new(LinguisticNotePropertyType)
	linguisticnotepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (linguisticnotetype *LinguisticNoteType) GongCopy() GongstructIF {
	newInstance := new(LinguisticNoteType)
	linguisticnotetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (notepropertytype *NotePropertyType) GongCopy() GongstructIF {
	newInstance := new(NotePropertyType)
	notepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (notetype *NoteType) GongCopy() GongstructIF {
	newInstance := new(NoteType)
	notetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongCopy() GongstructIF {
	newInstance := new(OrganisationAuthorityPropertyType)
	organisationauthoritypropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (routepropertytype *RoutePropertyType) GongCopy() GongstructIF {
	newInstance := new(RoutePropertyType)
	routepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongCopy() GongstructIF {
	newInstance := new(StandardLevelColumnPropertyType)
	standardlevelcolumnpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (surfacepropertytype *SurfacePropertyType) GongCopy() GongstructIF {
	newInstance := new(SurfacePropertyType)
	surfacepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (surfacetype *SurfaceType) GongCopy() GongstructIF {
	newInstance := new(SurfaceType)
	surfacetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (timesheetpropertytype *TimesheetPropertyType) GongCopy() GongstructIF {
	newInstance := new(TimesheetPropertyType)
	timesheetpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (timesheettype *TimesheetType) GongCopy() GongstructIF {
	newInstance := new(TimesheetType)
	timesheettype.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (a_extension *A_extension) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension)
}

func (a_extension_1 *A_extension_1) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_1)
}

func (a_extension_10 *A_extension_10) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_10)
}

func (a_extension_11 *A_extension_11) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_11)
}

func (a_extension_12 *A_extension_12) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_12)
}

func (a_extension_13 *A_extension_13) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_13)
}

func (a_extension_14 *A_extension_14) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_14)
}

func (a_extension_2 *A_extension_2) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_2)
}

func (a_extension_3 *A_extension_3) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_3)
}

func (a_extension_4 *A_extension_4) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_4)
}

func (a_extension_5 *A_extension_5) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_5)
}

func (a_extension_6 *A_extension_6) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_6)
}

func (a_extension_7 *A_extension_7) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_7)
}

func (a_extension_8 *A_extension_8) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_8)
}

func (a_extension_9 *A_extension_9) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_9)
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractpropertieswithscheduletype)
}

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, aircraftcharacteristicpropertytype)
}

func (aircraftcharacteristictype *AircraftCharacteristicType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, aircraftcharacteristictype)
}

func (airspaceactivationpropertytype *AirspaceActivationPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airspaceactivationpropertytype)
}

func (airspaceactivationtype *AirspaceActivationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airspaceactivationtype)
}

func (airspacegeometrycomponentpropertytype *AirspaceGeometryComponentPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airspacegeometrycomponentpropertytype)
}

func (airspacegeometrycomponenttype *AirspaceGeometryComponentType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airspacegeometrycomponenttype)
}

func (airspacelayerclasspropertytype *AirspaceLayerClassPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airspacelayerclasspropertytype)
}

func (airspacelayerclasstype *AirspaceLayerClassType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airspacelayerclasstype)
}

func (airspacelayerpropertytype *AirspaceLayerPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airspacelayerpropertytype)
}

func (airspacelayertype *AirspaceLayerType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airspacelayertype)
}

func (airspacepropertytype *AirspacePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airspacepropertytype)
}

func (airspacetimeslicepropertytype *AirspaceTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airspacetimeslicepropertytype)
}

func (airspacetimeslicetype *AirspaceTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airspacetimeslicetype)
}

func (airspacetype *AirspaceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airspacetype)
}

func (airspacevolumedependencypropertytype *AirspaceVolumeDependencyPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airspacevolumedependencypropertytype)
}

func (airspacevolumedependencytype *AirspaceVolumeDependencyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airspacevolumedependencytype)
}

func (airspacevolumepropertytype *AirspaceVolumePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airspacevolumepropertytype)
}

func (airspacevolumetype *AirspaceVolumeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airspacevolumetype)
}

func (authorityforairspacetimeslicepropertytype *AuthorityForAirspaceTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, authorityforairspacetimeslicepropertytype)
}

func (authorityforairspacetimeslicetype *AuthorityForAirspaceTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, authorityforairspacetimeslicetype)
}

func (authorityforairspacetype *AuthorityForAirspaceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, authorityforairspacetype)
}

func (curvepropertytype *CurvePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, curvepropertytype)
}

func (curvetype *CurveType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, curvetype)
}

func (elevatedcurvetype *ElevatedCurveType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, elevatedcurvetype)
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, elevatedsurfacetype)
}

func (geobordertimeslicepropertytype *GeoBorderTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, geobordertimeslicepropertytype)
}

func (geobordertimeslicetype *GeoBorderTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, geobordertimeslicetype)
}

func (geobordertype *GeoBorderType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, geobordertype)
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, linguisticnotepropertytype)
}

func (linguisticnotetype *LinguisticNoteType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, linguisticnotetype)
}

func (notepropertytype *NotePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, notepropertytype)
}

func (notetype *NoteType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, notetype)
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, organisationauthoritypropertytype)
}

func (routepropertytype *RoutePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, routepropertytype)
}

func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, standardlevelcolumnpropertytype)
}

func (surfacepropertytype *SurfacePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, surfacepropertytype)
}

func (surfacetype *SurfaceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, surfacetype)
}

func (timesheetpropertytype *TimesheetPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, timesheetpropertytype)
}

func (timesheettype *TimesheetType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, timesheettype)
}


type GongstructDiffable[T any] interface {
	GongstructPtr
	GongMarshallIdentifier(stage *Stage) string
	GongMarshallUnstaging(stage *Stage) string
	GongMarshallAllFields(stage *Stage) (string, string)
	GongReconstructPointersFromInstances(stage *Stage)
	GongDiff(stage *Stage, other T) []string
}

func computeCommitsForType[T GongstructDiffable[T]](
	stage *Stage,
	stagedInstances map[T]struct{},
	stagedOrder map[T]uint,
	referenceInstances map[T]T,
	referenceOrder *map[T]uint,
	instancesMap map[T]T,
	newInstancesSlice *[]string,
	fieldsEditSlice *[]string,
	deletedInstancesSlice *[]string,
	newInstancesReverseSlice *[]string,
	fieldsEditReverseSlice *[]string,
	deletedInstancesReverseSlice *[]string,
	lenNewInstances *int,
	lenDeletedInstances *int,
	lenModifiedInstances *int,
) {
	var newInstances []T
	var deletedInstances []T

	// parse all staged instances and check if they have a reference
	for instance := range stagedInstances {
		if ref, ok := referenceInstances[instance]; !ok {
			newInstances = append(newInstances, instance)
			*newInstancesSlice = append(*newInstancesSlice, instance.GongMarshallIdentifier(stage))
			if *referenceOrder == nil {
				*referenceOrder = make(map[T]uint)
			}
			(*referenceOrder)[instance] = stagedOrder[instance]
			*newInstancesReverseSlice = append(*newInstancesReverseSlice, instance.GongMarshallUnstaging(stage))
			fieldInitializers, pointersInitializations := instance.GongMarshallAllFields(stage)
			*fieldsEditSlice = append(*fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stagedOrder[ref] = stagedOrder[instance]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := instance.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, instance)
			if len(diffs) > 0 {
				var fieldsEdit string
				if instance.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", instance.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				*fieldsEditSlice = append(*fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, reverseDiff)
				}
				*lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range referenceInstances {
		instance := instancesMap[ref] // get the instance corresponding to the reference
		if _, ok := stagedInstances[instance]; !ok { // if the instance is not staged anymore, it means it has been unstaged
			deletedInstances = append(deletedInstances, ref)
			*deletedInstancesSlice = append(*deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			*deletedInstancesReverseSlice = append(*deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	*lenNewInstances += len(newInstances)
	*lenDeletedInstances += len(deletedInstances)
}

func (stage *Stage) ComputeForwardAndBackwardCommits() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	computeCommitsForType(
		stage,
		stage.A_extensions,
		stage.A_extension_stagedOrder,
		stage.A_extensions_reference,
		&stage.A_extensions_referenceOrder,
		stage.A_extensions_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_1s,
		stage.A_extension_1_stagedOrder,
		stage.A_extension_1s_reference,
		&stage.A_extension_1s_referenceOrder,
		stage.A_extension_1s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_10s,
		stage.A_extension_10_stagedOrder,
		stage.A_extension_10s_reference,
		&stage.A_extension_10s_referenceOrder,
		stage.A_extension_10s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_11s,
		stage.A_extension_11_stagedOrder,
		stage.A_extension_11s_reference,
		&stage.A_extension_11s_referenceOrder,
		stage.A_extension_11s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_12s,
		stage.A_extension_12_stagedOrder,
		stage.A_extension_12s_reference,
		&stage.A_extension_12s_referenceOrder,
		stage.A_extension_12s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_13s,
		stage.A_extension_13_stagedOrder,
		stage.A_extension_13s_reference,
		&stage.A_extension_13s_referenceOrder,
		stage.A_extension_13s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_14s,
		stage.A_extension_14_stagedOrder,
		stage.A_extension_14s_reference,
		&stage.A_extension_14s_referenceOrder,
		stage.A_extension_14s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_2s,
		stage.A_extension_2_stagedOrder,
		stage.A_extension_2s_reference,
		&stage.A_extension_2s_referenceOrder,
		stage.A_extension_2s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_3s,
		stage.A_extension_3_stagedOrder,
		stage.A_extension_3s_reference,
		&stage.A_extension_3s_referenceOrder,
		stage.A_extension_3s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_4s,
		stage.A_extension_4_stagedOrder,
		stage.A_extension_4s_reference,
		&stage.A_extension_4s_referenceOrder,
		stage.A_extension_4s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_5s,
		stage.A_extension_5_stagedOrder,
		stage.A_extension_5s_reference,
		&stage.A_extension_5s_referenceOrder,
		stage.A_extension_5s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_6s,
		stage.A_extension_6_stagedOrder,
		stage.A_extension_6s_reference,
		&stage.A_extension_6s_referenceOrder,
		stage.A_extension_6s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_7s,
		stage.A_extension_7_stagedOrder,
		stage.A_extension_7s_reference,
		&stage.A_extension_7s_referenceOrder,
		stage.A_extension_7s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_8s,
		stage.A_extension_8_stagedOrder,
		stage.A_extension_8s_reference,
		&stage.A_extension_8s_referenceOrder,
		stage.A_extension_8s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_extension_9s,
		stage.A_extension_9_stagedOrder,
		stage.A_extension_9s_reference,
		&stage.A_extension_9s_referenceOrder,
		stage.A_extension_9s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AbstractPropertiesWithScheduleTypes,
		stage.AbstractPropertiesWithScheduleType_stagedOrder,
		stage.AbstractPropertiesWithScheduleTypes_reference,
		&stage.AbstractPropertiesWithScheduleTypes_referenceOrder,
		stage.AbstractPropertiesWithScheduleTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AircraftCharacteristicPropertyTypes,
		stage.AircraftCharacteristicPropertyType_stagedOrder,
		stage.AircraftCharacteristicPropertyTypes_reference,
		&stage.AircraftCharacteristicPropertyTypes_referenceOrder,
		stage.AircraftCharacteristicPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AircraftCharacteristicTypes,
		stage.AircraftCharacteristicType_stagedOrder,
		stage.AircraftCharacteristicTypes_reference,
		&stage.AircraftCharacteristicTypes_referenceOrder,
		stage.AircraftCharacteristicTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AirspaceActivationPropertyTypes,
		stage.AirspaceActivationPropertyType_stagedOrder,
		stage.AirspaceActivationPropertyTypes_reference,
		&stage.AirspaceActivationPropertyTypes_referenceOrder,
		stage.AirspaceActivationPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AirspaceActivationTypes,
		stage.AirspaceActivationType_stagedOrder,
		stage.AirspaceActivationTypes_reference,
		&stage.AirspaceActivationTypes_referenceOrder,
		stage.AirspaceActivationTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AirspaceGeometryComponentPropertyTypes,
		stage.AirspaceGeometryComponentPropertyType_stagedOrder,
		stage.AirspaceGeometryComponentPropertyTypes_reference,
		&stage.AirspaceGeometryComponentPropertyTypes_referenceOrder,
		stage.AirspaceGeometryComponentPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AirspaceGeometryComponentTypes,
		stage.AirspaceGeometryComponentType_stagedOrder,
		stage.AirspaceGeometryComponentTypes_reference,
		&stage.AirspaceGeometryComponentTypes_referenceOrder,
		stage.AirspaceGeometryComponentTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AirspaceLayerClassPropertyTypes,
		stage.AirspaceLayerClassPropertyType_stagedOrder,
		stage.AirspaceLayerClassPropertyTypes_reference,
		&stage.AirspaceLayerClassPropertyTypes_referenceOrder,
		stage.AirspaceLayerClassPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AirspaceLayerClassTypes,
		stage.AirspaceLayerClassType_stagedOrder,
		stage.AirspaceLayerClassTypes_reference,
		&stage.AirspaceLayerClassTypes_referenceOrder,
		stage.AirspaceLayerClassTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AirspaceLayerPropertyTypes,
		stage.AirspaceLayerPropertyType_stagedOrder,
		stage.AirspaceLayerPropertyTypes_reference,
		&stage.AirspaceLayerPropertyTypes_referenceOrder,
		stage.AirspaceLayerPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AirspaceLayerTypes,
		stage.AirspaceLayerType_stagedOrder,
		stage.AirspaceLayerTypes_reference,
		&stage.AirspaceLayerTypes_referenceOrder,
		stage.AirspaceLayerTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AirspacePropertyTypes,
		stage.AirspacePropertyType_stagedOrder,
		stage.AirspacePropertyTypes_reference,
		&stage.AirspacePropertyTypes_referenceOrder,
		stage.AirspacePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AirspaceTimeSlicePropertyTypes,
		stage.AirspaceTimeSlicePropertyType_stagedOrder,
		stage.AirspaceTimeSlicePropertyTypes_reference,
		&stage.AirspaceTimeSlicePropertyTypes_referenceOrder,
		stage.AirspaceTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AirspaceTimeSliceTypes,
		stage.AirspaceTimeSliceType_stagedOrder,
		stage.AirspaceTimeSliceTypes_reference,
		&stage.AirspaceTimeSliceTypes_referenceOrder,
		stage.AirspaceTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AirspaceTypes,
		stage.AirspaceType_stagedOrder,
		stage.AirspaceTypes_reference,
		&stage.AirspaceTypes_referenceOrder,
		stage.AirspaceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AirspaceVolumeDependencyPropertyTypes,
		stage.AirspaceVolumeDependencyPropertyType_stagedOrder,
		stage.AirspaceVolumeDependencyPropertyTypes_reference,
		&stage.AirspaceVolumeDependencyPropertyTypes_referenceOrder,
		stage.AirspaceVolumeDependencyPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AirspaceVolumeDependencyTypes,
		stage.AirspaceVolumeDependencyType_stagedOrder,
		stage.AirspaceVolumeDependencyTypes_reference,
		&stage.AirspaceVolumeDependencyTypes_referenceOrder,
		stage.AirspaceVolumeDependencyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AirspaceVolumePropertyTypes,
		stage.AirspaceVolumePropertyType_stagedOrder,
		stage.AirspaceVolumePropertyTypes_reference,
		&stage.AirspaceVolumePropertyTypes_referenceOrder,
		stage.AirspaceVolumePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AirspaceVolumeTypes,
		stage.AirspaceVolumeType_stagedOrder,
		stage.AirspaceVolumeTypes_reference,
		&stage.AirspaceVolumeTypes_referenceOrder,
		stage.AirspaceVolumeTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AuthorityForAirspaceTimeSlicePropertyTypes,
		stage.AuthorityForAirspaceTimeSlicePropertyType_stagedOrder,
		stage.AuthorityForAirspaceTimeSlicePropertyTypes_reference,
		&stage.AuthorityForAirspaceTimeSlicePropertyTypes_referenceOrder,
		stage.AuthorityForAirspaceTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AuthorityForAirspaceTimeSliceTypes,
		stage.AuthorityForAirspaceTimeSliceType_stagedOrder,
		stage.AuthorityForAirspaceTimeSliceTypes_reference,
		&stage.AuthorityForAirspaceTimeSliceTypes_referenceOrder,
		stage.AuthorityForAirspaceTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AuthorityForAirspaceTypes,
		stage.AuthorityForAirspaceType_stagedOrder,
		stage.AuthorityForAirspaceTypes_reference,
		&stage.AuthorityForAirspaceTypes_referenceOrder,
		stage.AuthorityForAirspaceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CurvePropertyTypes,
		stage.CurvePropertyType_stagedOrder,
		stage.CurvePropertyTypes_reference,
		&stage.CurvePropertyTypes_referenceOrder,
		stage.CurvePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CurveTypes,
		stage.CurveType_stagedOrder,
		stage.CurveTypes_reference,
		&stage.CurveTypes_referenceOrder,
		stage.CurveTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ElevatedCurveTypes,
		stage.ElevatedCurveType_stagedOrder,
		stage.ElevatedCurveTypes_reference,
		&stage.ElevatedCurveTypes_referenceOrder,
		stage.ElevatedCurveTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ElevatedSurfaceTypes,
		stage.ElevatedSurfaceType_stagedOrder,
		stage.ElevatedSurfaceTypes_reference,
		&stage.ElevatedSurfaceTypes_referenceOrder,
		stage.ElevatedSurfaceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.GeoBorderTimeSlicePropertyTypes,
		stage.GeoBorderTimeSlicePropertyType_stagedOrder,
		stage.GeoBorderTimeSlicePropertyTypes_reference,
		&stage.GeoBorderTimeSlicePropertyTypes_referenceOrder,
		stage.GeoBorderTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.GeoBorderTimeSliceTypes,
		stage.GeoBorderTimeSliceType_stagedOrder,
		stage.GeoBorderTimeSliceTypes_reference,
		&stage.GeoBorderTimeSliceTypes_referenceOrder,
		stage.GeoBorderTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.GeoBorderTypes,
		stage.GeoBorderType_stagedOrder,
		stage.GeoBorderTypes_reference,
		&stage.GeoBorderTypes_referenceOrder,
		stage.GeoBorderTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.LinguisticNotePropertyTypes,
		stage.LinguisticNotePropertyType_stagedOrder,
		stage.LinguisticNotePropertyTypes_reference,
		&stage.LinguisticNotePropertyTypes_referenceOrder,
		stage.LinguisticNotePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.LinguisticNoteTypes,
		stage.LinguisticNoteType_stagedOrder,
		stage.LinguisticNoteTypes_reference,
		&stage.LinguisticNoteTypes_referenceOrder,
		stage.LinguisticNoteTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.NotePropertyTypes,
		stage.NotePropertyType_stagedOrder,
		stage.NotePropertyTypes_reference,
		&stage.NotePropertyTypes_referenceOrder,
		stage.NotePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.NoteTypes,
		stage.NoteType_stagedOrder,
		stage.NoteTypes_reference,
		&stage.NoteTypes_referenceOrder,
		stage.NoteTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.OrganisationAuthorityPropertyTypes,
		stage.OrganisationAuthorityPropertyType_stagedOrder,
		stage.OrganisationAuthorityPropertyTypes_reference,
		&stage.OrganisationAuthorityPropertyTypes_referenceOrder,
		stage.OrganisationAuthorityPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.RoutePropertyTypes,
		stage.RoutePropertyType_stagedOrder,
		stage.RoutePropertyTypes_reference,
		&stage.RoutePropertyTypes_referenceOrder,
		stage.RoutePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StandardLevelColumnPropertyTypes,
		stage.StandardLevelColumnPropertyType_stagedOrder,
		stage.StandardLevelColumnPropertyTypes_reference,
		&stage.StandardLevelColumnPropertyTypes_referenceOrder,
		stage.StandardLevelColumnPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SurfacePropertyTypes,
		stage.SurfacePropertyType_stagedOrder,
		stage.SurfacePropertyTypes_reference,
		&stage.SurfacePropertyTypes_referenceOrder,
		stage.SurfacePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SurfaceTypes,
		stage.SurfaceType_stagedOrder,
		stage.SurfaceTypes_reference,
		&stage.SurfaceTypes_referenceOrder,
		stage.SurfaceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TimesheetPropertyTypes,
		stage.TimesheetPropertyType_stagedOrder,
		stage.TimesheetPropertyTypes_reference,
		&stage.TimesheetPropertyTypes_referenceOrder,
		stage.TimesheetPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TimesheetTypes,
		stage.TimesheetType_stagedOrder,
		stage.TimesheetTypes_reference,
		&stage.TimesheetTypes_referenceOrder,
		stage.TimesheetTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.AppendForwardCommit(forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.AppendBackwardCommit(backwardCommit)
		stage.SetModified(true)
	} else {
		stage.SetModified(false)
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	__gong__computeReferencePass1(stage, stage.A_extensions, &stage.A_extensions_reference, &stage.A_extensions_referenceOrder, &stage.A_extensions_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_1s, &stage.A_extension_1s_reference, &stage.A_extension_1s_referenceOrder, &stage.A_extension_1s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_10s, &stage.A_extension_10s_reference, &stage.A_extension_10s_referenceOrder, &stage.A_extension_10s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_11s, &stage.A_extension_11s_reference, &stage.A_extension_11s_referenceOrder, &stage.A_extension_11s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_12s, &stage.A_extension_12s_reference, &stage.A_extension_12s_referenceOrder, &stage.A_extension_12s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_13s, &stage.A_extension_13s_reference, &stage.A_extension_13s_referenceOrder, &stage.A_extension_13s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_14s, &stage.A_extension_14s_reference, &stage.A_extension_14s_referenceOrder, &stage.A_extension_14s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_2s, &stage.A_extension_2s_reference, &stage.A_extension_2s_referenceOrder, &stage.A_extension_2s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_3s, &stage.A_extension_3s_reference, &stage.A_extension_3s_referenceOrder, &stage.A_extension_3s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_4s, &stage.A_extension_4s_reference, &stage.A_extension_4s_referenceOrder, &stage.A_extension_4s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_5s, &stage.A_extension_5s_reference, &stage.A_extension_5s_referenceOrder, &stage.A_extension_5s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_6s, &stage.A_extension_6s_reference, &stage.A_extension_6s_referenceOrder, &stage.A_extension_6s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_7s, &stage.A_extension_7s_reference, &stage.A_extension_7s_referenceOrder, &stage.A_extension_7s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_8s, &stage.A_extension_8s_reference, &stage.A_extension_8s_referenceOrder, &stage.A_extension_8s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_9s, &stage.A_extension_9s_reference, &stage.A_extension_9s_referenceOrder, &stage.A_extension_9s_instance)

	__gong__computeReferencePass1(stage, stage.AbstractPropertiesWithScheduleTypes, &stage.AbstractPropertiesWithScheduleTypes_reference, &stage.AbstractPropertiesWithScheduleTypes_referenceOrder, &stage.AbstractPropertiesWithScheduleTypes_instance)

	__gong__computeReferencePass1(stage, stage.AircraftCharacteristicPropertyTypes, &stage.AircraftCharacteristicPropertyTypes_reference, &stage.AircraftCharacteristicPropertyTypes_referenceOrder, &stage.AircraftCharacteristicPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AircraftCharacteristicTypes, &stage.AircraftCharacteristicTypes_reference, &stage.AircraftCharacteristicTypes_referenceOrder, &stage.AircraftCharacteristicTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirspaceActivationPropertyTypes, &stage.AirspaceActivationPropertyTypes_reference, &stage.AirspaceActivationPropertyTypes_referenceOrder, &stage.AirspaceActivationPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirspaceActivationTypes, &stage.AirspaceActivationTypes_reference, &stage.AirspaceActivationTypes_referenceOrder, &stage.AirspaceActivationTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirspaceGeometryComponentPropertyTypes, &stage.AirspaceGeometryComponentPropertyTypes_reference, &stage.AirspaceGeometryComponentPropertyTypes_referenceOrder, &stage.AirspaceGeometryComponentPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirspaceGeometryComponentTypes, &stage.AirspaceGeometryComponentTypes_reference, &stage.AirspaceGeometryComponentTypes_referenceOrder, &stage.AirspaceGeometryComponentTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirspaceLayerClassPropertyTypes, &stage.AirspaceLayerClassPropertyTypes_reference, &stage.AirspaceLayerClassPropertyTypes_referenceOrder, &stage.AirspaceLayerClassPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirspaceLayerClassTypes, &stage.AirspaceLayerClassTypes_reference, &stage.AirspaceLayerClassTypes_referenceOrder, &stage.AirspaceLayerClassTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirspaceLayerPropertyTypes, &stage.AirspaceLayerPropertyTypes_reference, &stage.AirspaceLayerPropertyTypes_referenceOrder, &stage.AirspaceLayerPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirspaceLayerTypes, &stage.AirspaceLayerTypes_reference, &stage.AirspaceLayerTypes_referenceOrder, &stage.AirspaceLayerTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirspacePropertyTypes, &stage.AirspacePropertyTypes_reference, &stage.AirspacePropertyTypes_referenceOrder, &stage.AirspacePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirspaceTimeSlicePropertyTypes, &stage.AirspaceTimeSlicePropertyTypes_reference, &stage.AirspaceTimeSlicePropertyTypes_referenceOrder, &stage.AirspaceTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirspaceTimeSliceTypes, &stage.AirspaceTimeSliceTypes_reference, &stage.AirspaceTimeSliceTypes_referenceOrder, &stage.AirspaceTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirspaceTypes, &stage.AirspaceTypes_reference, &stage.AirspaceTypes_referenceOrder, &stage.AirspaceTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirspaceVolumeDependencyPropertyTypes, &stage.AirspaceVolumeDependencyPropertyTypes_reference, &stage.AirspaceVolumeDependencyPropertyTypes_referenceOrder, &stage.AirspaceVolumeDependencyPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirspaceVolumeDependencyTypes, &stage.AirspaceVolumeDependencyTypes_reference, &stage.AirspaceVolumeDependencyTypes_referenceOrder, &stage.AirspaceVolumeDependencyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirspaceVolumePropertyTypes, &stage.AirspaceVolumePropertyTypes_reference, &stage.AirspaceVolumePropertyTypes_referenceOrder, &stage.AirspaceVolumePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirspaceVolumeTypes, &stage.AirspaceVolumeTypes_reference, &stage.AirspaceVolumeTypes_referenceOrder, &stage.AirspaceVolumeTypes_instance)

	__gong__computeReferencePass1(stage, stage.AuthorityForAirspaceTimeSlicePropertyTypes, &stage.AuthorityForAirspaceTimeSlicePropertyTypes_reference, &stage.AuthorityForAirspaceTimeSlicePropertyTypes_referenceOrder, &stage.AuthorityForAirspaceTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AuthorityForAirspaceTimeSliceTypes, &stage.AuthorityForAirspaceTimeSliceTypes_reference, &stage.AuthorityForAirspaceTimeSliceTypes_referenceOrder, &stage.AuthorityForAirspaceTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.AuthorityForAirspaceTypes, &stage.AuthorityForAirspaceTypes_reference, &stage.AuthorityForAirspaceTypes_referenceOrder, &stage.AuthorityForAirspaceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CurvePropertyTypes, &stage.CurvePropertyTypes_reference, &stage.CurvePropertyTypes_referenceOrder, &stage.CurvePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.CurveTypes, &stage.CurveTypes_reference, &stage.CurveTypes_referenceOrder, &stage.CurveTypes_instance)

	__gong__computeReferencePass1(stage, stage.ElevatedCurveTypes, &stage.ElevatedCurveTypes_reference, &stage.ElevatedCurveTypes_referenceOrder, &stage.ElevatedCurveTypes_instance)

	__gong__computeReferencePass1(stage, stage.ElevatedSurfaceTypes, &stage.ElevatedSurfaceTypes_reference, &stage.ElevatedSurfaceTypes_referenceOrder, &stage.ElevatedSurfaceTypes_instance)

	__gong__computeReferencePass1(stage, stage.GeoBorderTimeSlicePropertyTypes, &stage.GeoBorderTimeSlicePropertyTypes_reference, &stage.GeoBorderTimeSlicePropertyTypes_referenceOrder, &stage.GeoBorderTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.GeoBorderTimeSliceTypes, &stage.GeoBorderTimeSliceTypes_reference, &stage.GeoBorderTimeSliceTypes_referenceOrder, &stage.GeoBorderTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.GeoBorderTypes, &stage.GeoBorderTypes_reference, &stage.GeoBorderTypes_referenceOrder, &stage.GeoBorderTypes_instance)

	__gong__computeReferencePass1(stage, stage.LinguisticNotePropertyTypes, &stage.LinguisticNotePropertyTypes_reference, &stage.LinguisticNotePropertyTypes_referenceOrder, &stage.LinguisticNotePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.LinguisticNoteTypes, &stage.LinguisticNoteTypes_reference, &stage.LinguisticNoteTypes_referenceOrder, &stage.LinguisticNoteTypes_instance)

	__gong__computeReferencePass1(stage, stage.NotePropertyTypes, &stage.NotePropertyTypes_reference, &stage.NotePropertyTypes_referenceOrder, &stage.NotePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.NoteTypes, &stage.NoteTypes_reference, &stage.NoteTypes_referenceOrder, &stage.NoteTypes_instance)

	__gong__computeReferencePass1(stage, stage.OrganisationAuthorityPropertyTypes, &stage.OrganisationAuthorityPropertyTypes_reference, &stage.OrganisationAuthorityPropertyTypes_referenceOrder, &stage.OrganisationAuthorityPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.RoutePropertyTypes, &stage.RoutePropertyTypes_reference, &stage.RoutePropertyTypes_referenceOrder, &stage.RoutePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.StandardLevelColumnPropertyTypes, &stage.StandardLevelColumnPropertyTypes_reference, &stage.StandardLevelColumnPropertyTypes_referenceOrder, &stage.StandardLevelColumnPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.SurfacePropertyTypes, &stage.SurfacePropertyTypes_reference, &stage.SurfacePropertyTypes_referenceOrder, &stage.SurfacePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.SurfaceTypes, &stage.SurfaceTypes_reference, &stage.SurfaceTypes_referenceOrder, &stage.SurfaceTypes_instance)

	__gong__computeReferencePass1(stage, stage.TimesheetPropertyTypes, &stage.TimesheetPropertyTypes_reference, &stage.TimesheetPropertyTypes_referenceOrder, &stage.TimesheetPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.TimesheetTypes, &stage.TimesheetTypes_reference, &stage.TimesheetTypes_referenceOrder, &stage.TimesheetTypes_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.A_extensions, stage.A_extensions_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_1s, stage.A_extension_1s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_10s, stage.A_extension_10s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_11s, stage.A_extension_11s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_12s, stage.A_extension_12s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_13s, stage.A_extension_13s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_14s, stage.A_extension_14s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_2s, stage.A_extension_2s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_3s, stage.A_extension_3s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_4s, stage.A_extension_4s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_5s, stage.A_extension_5s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_6s, stage.A_extension_6s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_7s, stage.A_extension_7s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_8s, stage.A_extension_8s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_9s, stage.A_extension_9s_reference, stage)

	__gong__computeReferencePass2(stage.AbstractPropertiesWithScheduleTypes, stage.AbstractPropertiesWithScheduleTypes_reference, stage)

	__gong__computeReferencePass2(stage.AircraftCharacteristicPropertyTypes, stage.AircraftCharacteristicPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AircraftCharacteristicTypes, stage.AircraftCharacteristicTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirspaceActivationPropertyTypes, stage.AirspaceActivationPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirspaceActivationTypes, stage.AirspaceActivationTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirspaceGeometryComponentPropertyTypes, stage.AirspaceGeometryComponentPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirspaceGeometryComponentTypes, stage.AirspaceGeometryComponentTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirspaceLayerClassPropertyTypes, stage.AirspaceLayerClassPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirspaceLayerClassTypes, stage.AirspaceLayerClassTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirspaceLayerPropertyTypes, stage.AirspaceLayerPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirspaceLayerTypes, stage.AirspaceLayerTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirspacePropertyTypes, stage.AirspacePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirspaceTimeSlicePropertyTypes, stage.AirspaceTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirspaceTimeSliceTypes, stage.AirspaceTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirspaceTypes, stage.AirspaceTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirspaceVolumeDependencyPropertyTypes, stage.AirspaceVolumeDependencyPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirspaceVolumeDependencyTypes, stage.AirspaceVolumeDependencyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirspaceVolumePropertyTypes, stage.AirspaceVolumePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirspaceVolumeTypes, stage.AirspaceVolumeTypes_reference, stage)

	__gong__computeReferencePass2(stage.AuthorityForAirspaceTimeSlicePropertyTypes, stage.AuthorityForAirspaceTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AuthorityForAirspaceTimeSliceTypes, stage.AuthorityForAirspaceTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.AuthorityForAirspaceTypes, stage.AuthorityForAirspaceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CurvePropertyTypes, stage.CurvePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.CurveTypes, stage.CurveTypes_reference, stage)

	__gong__computeReferencePass2(stage.ElevatedCurveTypes, stage.ElevatedCurveTypes_reference, stage)

	__gong__computeReferencePass2(stage.ElevatedSurfaceTypes, stage.ElevatedSurfaceTypes_reference, stage)

	__gong__computeReferencePass2(stage.GeoBorderTimeSlicePropertyTypes, stage.GeoBorderTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.GeoBorderTimeSliceTypes, stage.GeoBorderTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.GeoBorderTypes, stage.GeoBorderTypes_reference, stage)

	__gong__computeReferencePass2(stage.LinguisticNotePropertyTypes, stage.LinguisticNotePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.LinguisticNoteTypes, stage.LinguisticNoteTypes_reference, stage)

	__gong__computeReferencePass2(stage.NotePropertyTypes, stage.NotePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.NoteTypes, stage.NoteTypes_reference, stage)

	__gong__computeReferencePass2(stage.OrganisationAuthorityPropertyTypes, stage.OrganisationAuthorityPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.RoutePropertyTypes, stage.RoutePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.StandardLevelColumnPropertyTypes, stage.StandardLevelColumnPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.SurfacePropertyTypes, stage.SurfacePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.SurfaceTypes, stage.SurfaceTypes_reference, stage)

	__gong__computeReferencePass2(stage.TimesheetPropertyTypes, stage.TimesheetPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.TimesheetTypes, stage.TimesheetTypes_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (a_extension *A_extension) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_stagedOrder, stage.A_extensions_referenceOrder, a_extension, "A_extension")
}

func (a_extension_1 *A_extension_1) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_1_stagedOrder, stage.A_extension_1s_referenceOrder, a_extension_1, "A_extension_1")
}

func (a_extension_10 *A_extension_10) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_10_stagedOrder, stage.A_extension_10s_referenceOrder, a_extension_10, "A_extension_10")
}

func (a_extension_11 *A_extension_11) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_11_stagedOrder, stage.A_extension_11s_referenceOrder, a_extension_11, "A_extension_11")
}

func (a_extension_12 *A_extension_12) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_12_stagedOrder, stage.A_extension_12s_referenceOrder, a_extension_12, "A_extension_12")
}

func (a_extension_13 *A_extension_13) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_13_stagedOrder, stage.A_extension_13s_referenceOrder, a_extension_13, "A_extension_13")
}

func (a_extension_14 *A_extension_14) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_14_stagedOrder, stage.A_extension_14s_referenceOrder, a_extension_14, "A_extension_14")
}

func (a_extension_2 *A_extension_2) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_2_stagedOrder, stage.A_extension_2s_referenceOrder, a_extension_2, "A_extension_2")
}

func (a_extension_3 *A_extension_3) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_3_stagedOrder, stage.A_extension_3s_referenceOrder, a_extension_3, "A_extension_3")
}

func (a_extension_4 *A_extension_4) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_4_stagedOrder, stage.A_extension_4s_referenceOrder, a_extension_4, "A_extension_4")
}

func (a_extension_5 *A_extension_5) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_5_stagedOrder, stage.A_extension_5s_referenceOrder, a_extension_5, "A_extension_5")
}

func (a_extension_6 *A_extension_6) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_6_stagedOrder, stage.A_extension_6s_referenceOrder, a_extension_6, "A_extension_6")
}

func (a_extension_7 *A_extension_7) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_7_stagedOrder, stage.A_extension_7s_referenceOrder, a_extension_7, "A_extension_7")
}

func (a_extension_8 *A_extension_8) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_8_stagedOrder, stage.A_extension_8s_referenceOrder, a_extension_8, "A_extension_8")
}

func (a_extension_9 *A_extension_9) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_9_stagedOrder, stage.A_extension_9s_referenceOrder, a_extension_9, "A_extension_9")
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractPropertiesWithScheduleType_stagedOrder, stage.AbstractPropertiesWithScheduleTypes_referenceOrder, abstractpropertieswithscheduletype, "AbstractPropertiesWithScheduleType")
}

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AircraftCharacteristicPropertyType_stagedOrder, stage.AircraftCharacteristicPropertyTypes_referenceOrder, aircraftcharacteristicpropertytype, "AircraftCharacteristicPropertyType")
}

func (aircraftcharacteristictype *AircraftCharacteristicType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AircraftCharacteristicType_stagedOrder, stage.AircraftCharacteristicTypes_referenceOrder, aircraftcharacteristictype, "AircraftCharacteristicType")
}

func (airspaceactivationpropertytype *AirspaceActivationPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirspaceActivationPropertyType_stagedOrder, stage.AirspaceActivationPropertyTypes_referenceOrder, airspaceactivationpropertytype, "AirspaceActivationPropertyType")
}

func (airspaceactivationtype *AirspaceActivationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirspaceActivationType_stagedOrder, stage.AirspaceActivationTypes_referenceOrder, airspaceactivationtype, "AirspaceActivationType")
}

func (airspacegeometrycomponentpropertytype *AirspaceGeometryComponentPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirspaceGeometryComponentPropertyType_stagedOrder, stage.AirspaceGeometryComponentPropertyTypes_referenceOrder, airspacegeometrycomponentpropertytype, "AirspaceGeometryComponentPropertyType")
}

func (airspacegeometrycomponenttype *AirspaceGeometryComponentType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirspaceGeometryComponentType_stagedOrder, stage.AirspaceGeometryComponentTypes_referenceOrder, airspacegeometrycomponenttype, "AirspaceGeometryComponentType")
}

func (airspacelayerclasspropertytype *AirspaceLayerClassPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirspaceLayerClassPropertyType_stagedOrder, stage.AirspaceLayerClassPropertyTypes_referenceOrder, airspacelayerclasspropertytype, "AirspaceLayerClassPropertyType")
}

func (airspacelayerclasstype *AirspaceLayerClassType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirspaceLayerClassType_stagedOrder, stage.AirspaceLayerClassTypes_referenceOrder, airspacelayerclasstype, "AirspaceLayerClassType")
}

func (airspacelayerpropertytype *AirspaceLayerPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirspaceLayerPropertyType_stagedOrder, stage.AirspaceLayerPropertyTypes_referenceOrder, airspacelayerpropertytype, "AirspaceLayerPropertyType")
}

func (airspacelayertype *AirspaceLayerType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirspaceLayerType_stagedOrder, stage.AirspaceLayerTypes_referenceOrder, airspacelayertype, "AirspaceLayerType")
}

func (airspacepropertytype *AirspacePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirspacePropertyType_stagedOrder, stage.AirspacePropertyTypes_referenceOrder, airspacepropertytype, "AirspacePropertyType")
}

func (airspacetimeslicepropertytype *AirspaceTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirspaceTimeSlicePropertyType_stagedOrder, stage.AirspaceTimeSlicePropertyTypes_referenceOrder, airspacetimeslicepropertytype, "AirspaceTimeSlicePropertyType")
}

func (airspacetimeslicetype *AirspaceTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirspaceTimeSliceType_stagedOrder, stage.AirspaceTimeSliceTypes_referenceOrder, airspacetimeslicetype, "AirspaceTimeSliceType")
}

func (airspacetype *AirspaceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirspaceType_stagedOrder, stage.AirspaceTypes_referenceOrder, airspacetype, "AirspaceType")
}

func (airspacevolumedependencypropertytype *AirspaceVolumeDependencyPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirspaceVolumeDependencyPropertyType_stagedOrder, stage.AirspaceVolumeDependencyPropertyTypes_referenceOrder, airspacevolumedependencypropertytype, "AirspaceVolumeDependencyPropertyType")
}

func (airspacevolumedependencytype *AirspaceVolumeDependencyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirspaceVolumeDependencyType_stagedOrder, stage.AirspaceVolumeDependencyTypes_referenceOrder, airspacevolumedependencytype, "AirspaceVolumeDependencyType")
}

func (airspacevolumepropertytype *AirspaceVolumePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirspaceVolumePropertyType_stagedOrder, stage.AirspaceVolumePropertyTypes_referenceOrder, airspacevolumepropertytype, "AirspaceVolumePropertyType")
}

func (airspacevolumetype *AirspaceVolumeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirspaceVolumeType_stagedOrder, stage.AirspaceVolumeTypes_referenceOrder, airspacevolumetype, "AirspaceVolumeType")
}

func (authorityforairspacetimeslicepropertytype *AuthorityForAirspaceTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AuthorityForAirspaceTimeSlicePropertyType_stagedOrder, stage.AuthorityForAirspaceTimeSlicePropertyTypes_referenceOrder, authorityforairspacetimeslicepropertytype, "AuthorityForAirspaceTimeSlicePropertyType")
}

func (authorityforairspacetimeslicetype *AuthorityForAirspaceTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AuthorityForAirspaceTimeSliceType_stagedOrder, stage.AuthorityForAirspaceTimeSliceTypes_referenceOrder, authorityforairspacetimeslicetype, "AuthorityForAirspaceTimeSliceType")
}

func (authorityforairspacetype *AuthorityForAirspaceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AuthorityForAirspaceType_stagedOrder, stage.AuthorityForAirspaceTypes_referenceOrder, authorityforairspacetype, "AuthorityForAirspaceType")
}

func (curvepropertytype *CurvePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CurvePropertyType_stagedOrder, stage.CurvePropertyTypes_referenceOrder, curvepropertytype, "CurvePropertyType")
}

func (curvetype *CurveType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CurveType_stagedOrder, stage.CurveTypes_referenceOrder, curvetype, "CurveType")
}

func (elevatedcurvetype *ElevatedCurveType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ElevatedCurveType_stagedOrder, stage.ElevatedCurveTypes_referenceOrder, elevatedcurvetype, "ElevatedCurveType")
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ElevatedSurfaceType_stagedOrder, stage.ElevatedSurfaceTypes_referenceOrder, elevatedsurfacetype, "ElevatedSurfaceType")
}

func (geobordertimeslicepropertytype *GeoBorderTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GeoBorderTimeSlicePropertyType_stagedOrder, stage.GeoBorderTimeSlicePropertyTypes_referenceOrder, geobordertimeslicepropertytype, "GeoBorderTimeSlicePropertyType")
}

func (geobordertimeslicetype *GeoBorderTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GeoBorderTimeSliceType_stagedOrder, stage.GeoBorderTimeSliceTypes_referenceOrder, geobordertimeslicetype, "GeoBorderTimeSliceType")
}

func (geobordertype *GeoBorderType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GeoBorderType_stagedOrder, stage.GeoBorderTypes_referenceOrder, geobordertype, "GeoBorderType")
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LinguisticNotePropertyType_stagedOrder, stage.LinguisticNotePropertyTypes_referenceOrder, linguisticnotepropertytype, "LinguisticNotePropertyType")
}

func (linguisticnotetype *LinguisticNoteType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LinguisticNoteType_stagedOrder, stage.LinguisticNoteTypes_referenceOrder, linguisticnotetype, "LinguisticNoteType")
}

func (notepropertytype *NotePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NotePropertyType_stagedOrder, stage.NotePropertyTypes_referenceOrder, notepropertytype, "NotePropertyType")
}

func (notetype *NoteType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoteType_stagedOrder, stage.NoteTypes_referenceOrder, notetype, "NoteType")
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.OrganisationAuthorityPropertyType_stagedOrder, stage.OrganisationAuthorityPropertyTypes_referenceOrder, organisationauthoritypropertytype, "OrganisationAuthorityPropertyType")
}

func (routepropertytype *RoutePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RoutePropertyType_stagedOrder, stage.RoutePropertyTypes_referenceOrder, routepropertytype, "RoutePropertyType")
}

func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StandardLevelColumnPropertyType_stagedOrder, stage.StandardLevelColumnPropertyTypes_referenceOrder, standardlevelcolumnpropertytype, "StandardLevelColumnPropertyType")
}

func (surfacepropertytype *SurfacePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SurfacePropertyType_stagedOrder, stage.SurfacePropertyTypes_referenceOrder, surfacepropertytype, "SurfacePropertyType")
}

func (surfacetype *SurfaceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SurfaceType_stagedOrder, stage.SurfaceTypes_referenceOrder, surfacetype, "SurfaceType")
}

func (timesheetpropertytype *TimesheetPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TimesheetPropertyType_stagedOrder, stage.TimesheetPropertyTypes_referenceOrder, timesheetpropertytype, "TimesheetPropertyType")
}

func (timesheettype *TimesheetType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TimesheetType_stagedOrder, stage.TimesheetTypes_referenceOrder, timesheettype, "TimesheetType")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (a_extension *A_extension) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension, a_extension.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension *A_extension) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension.GongGetIdentifier(stage)
}

func (a_extension_1 *A_extension_1) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_1, a_extension_1.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_1 *A_extension_1) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_1.GongGetIdentifier(stage)
}

func (a_extension_10 *A_extension_10) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_10, a_extension_10.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_10 *A_extension_10) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_10.GongGetIdentifier(stage)
}

func (a_extension_11 *A_extension_11) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_11, a_extension_11.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_11 *A_extension_11) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_11.GongGetIdentifier(stage)
}

func (a_extension_12 *A_extension_12) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_12, a_extension_12.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_12 *A_extension_12) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_12.GongGetIdentifier(stage)
}

func (a_extension_13 *A_extension_13) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_13, a_extension_13.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_13 *A_extension_13) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_13.GongGetIdentifier(stage)
}

func (a_extension_14 *A_extension_14) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_14, a_extension_14.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_14 *A_extension_14) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_14.GongGetIdentifier(stage)
}

func (a_extension_2 *A_extension_2) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_2, a_extension_2.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_2 *A_extension_2) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_2.GongGetIdentifier(stage)
}

func (a_extension_3 *A_extension_3) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_3, a_extension_3.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_3 *A_extension_3) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_3.GongGetIdentifier(stage)
}

func (a_extension_4 *A_extension_4) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_4, a_extension_4.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_4 *A_extension_4) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_4.GongGetIdentifier(stage)
}

func (a_extension_5 *A_extension_5) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_5, a_extension_5.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_5 *A_extension_5) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_5.GongGetIdentifier(stage)
}

func (a_extension_6 *A_extension_6) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_6, a_extension_6.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_6 *A_extension_6) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_6.GongGetIdentifier(stage)
}

func (a_extension_7 *A_extension_7) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_7, a_extension_7.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_7 *A_extension_7) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_7.GongGetIdentifier(stage)
}

func (a_extension_8 *A_extension_8) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_8, a_extension_8.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_8 *A_extension_8) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_8.GongGetIdentifier(stage)
}

func (a_extension_9 *A_extension_9) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_9, a_extension_9.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_9 *A_extension_9) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_9.GongGetIdentifier(stage)
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractpropertieswithscheduletype, abstractpropertieswithscheduletype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractpropertieswithscheduletype.GongGetIdentifier(stage)
}

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(aircraftcharacteristicpropertytype, aircraftcharacteristicpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return aircraftcharacteristicpropertytype.GongGetIdentifier(stage)
}

func (aircraftcharacteristictype *AircraftCharacteristicType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(aircraftcharacteristictype, aircraftcharacteristictype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (aircraftcharacteristictype *AircraftCharacteristicType) GongGetReferenceIdentifier(stage *Stage) string {
	return aircraftcharacteristictype.GongGetIdentifier(stage)
}

func (airspaceactivationpropertytype *AirspaceActivationPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airspaceactivationpropertytype, airspaceactivationpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airspaceactivationpropertytype *AirspaceActivationPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return airspaceactivationpropertytype.GongGetIdentifier(stage)
}

func (airspaceactivationtype *AirspaceActivationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airspaceactivationtype, airspaceactivationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airspaceactivationtype *AirspaceActivationType) GongGetReferenceIdentifier(stage *Stage) string {
	return airspaceactivationtype.GongGetIdentifier(stage)
}

func (airspacegeometrycomponentpropertytype *AirspaceGeometryComponentPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airspacegeometrycomponentpropertytype, airspacegeometrycomponentpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airspacegeometrycomponentpropertytype *AirspaceGeometryComponentPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return airspacegeometrycomponentpropertytype.GongGetIdentifier(stage)
}

func (airspacegeometrycomponenttype *AirspaceGeometryComponentType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airspacegeometrycomponenttype, airspacegeometrycomponenttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airspacegeometrycomponenttype *AirspaceGeometryComponentType) GongGetReferenceIdentifier(stage *Stage) string {
	return airspacegeometrycomponenttype.GongGetIdentifier(stage)
}

func (airspacelayerclasspropertytype *AirspaceLayerClassPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airspacelayerclasspropertytype, airspacelayerclasspropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airspacelayerclasspropertytype *AirspaceLayerClassPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return airspacelayerclasspropertytype.GongGetIdentifier(stage)
}

func (airspacelayerclasstype *AirspaceLayerClassType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airspacelayerclasstype, airspacelayerclasstype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airspacelayerclasstype *AirspaceLayerClassType) GongGetReferenceIdentifier(stage *Stage) string {
	return airspacelayerclasstype.GongGetIdentifier(stage)
}

func (airspacelayerpropertytype *AirspaceLayerPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airspacelayerpropertytype, airspacelayerpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airspacelayerpropertytype *AirspaceLayerPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return airspacelayerpropertytype.GongGetIdentifier(stage)
}

func (airspacelayertype *AirspaceLayerType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airspacelayertype, airspacelayertype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airspacelayertype *AirspaceLayerType) GongGetReferenceIdentifier(stage *Stage) string {
	return airspacelayertype.GongGetIdentifier(stage)
}

func (airspacepropertytype *AirspacePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airspacepropertytype, airspacepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airspacepropertytype *AirspacePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return airspacepropertytype.GongGetIdentifier(stage)
}

func (airspacetimeslicepropertytype *AirspaceTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airspacetimeslicepropertytype, airspacetimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airspacetimeslicepropertytype *AirspaceTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return airspacetimeslicepropertytype.GongGetIdentifier(stage)
}

func (airspacetimeslicetype *AirspaceTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airspacetimeslicetype, airspacetimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airspacetimeslicetype *AirspaceTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return airspacetimeslicetype.GongGetIdentifier(stage)
}

func (airspacetype *AirspaceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airspacetype, airspacetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airspacetype *AirspaceType) GongGetReferenceIdentifier(stage *Stage) string {
	return airspacetype.GongGetIdentifier(stage)
}

func (airspacevolumedependencypropertytype *AirspaceVolumeDependencyPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airspacevolumedependencypropertytype, airspacevolumedependencypropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airspacevolumedependencypropertytype *AirspaceVolumeDependencyPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return airspacevolumedependencypropertytype.GongGetIdentifier(stage)
}

func (airspacevolumedependencytype *AirspaceVolumeDependencyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airspacevolumedependencytype, airspacevolumedependencytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airspacevolumedependencytype *AirspaceVolumeDependencyType) GongGetReferenceIdentifier(stage *Stage) string {
	return airspacevolumedependencytype.GongGetIdentifier(stage)
}

func (airspacevolumepropertytype *AirspaceVolumePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airspacevolumepropertytype, airspacevolumepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airspacevolumepropertytype *AirspaceVolumePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return airspacevolumepropertytype.GongGetIdentifier(stage)
}

func (airspacevolumetype *AirspaceVolumeType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airspacevolumetype, airspacevolumetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airspacevolumetype *AirspaceVolumeType) GongGetReferenceIdentifier(stage *Stage) string {
	return airspacevolumetype.GongGetIdentifier(stage)
}

func (authorityforairspacetimeslicepropertytype *AuthorityForAirspaceTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(authorityforairspacetimeslicepropertytype, authorityforairspacetimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (authorityforairspacetimeslicepropertytype *AuthorityForAirspaceTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return authorityforairspacetimeslicepropertytype.GongGetIdentifier(stage)
}

func (authorityforairspacetimeslicetype *AuthorityForAirspaceTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(authorityforairspacetimeslicetype, authorityforairspacetimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (authorityforairspacetimeslicetype *AuthorityForAirspaceTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return authorityforairspacetimeslicetype.GongGetIdentifier(stage)
}

func (authorityforairspacetype *AuthorityForAirspaceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(authorityforairspacetype, authorityforairspacetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (authorityforairspacetype *AuthorityForAirspaceType) GongGetReferenceIdentifier(stage *Stage) string {
	return authorityforairspacetype.GongGetIdentifier(stage)
}

func (curvepropertytype *CurvePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(curvepropertytype, curvepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (curvepropertytype *CurvePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return curvepropertytype.GongGetIdentifier(stage)
}

func (curvetype *CurveType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(curvetype, curvetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (curvetype *CurveType) GongGetReferenceIdentifier(stage *Stage) string {
	return curvetype.GongGetIdentifier(stage)
}

func (elevatedcurvetype *ElevatedCurveType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(elevatedcurvetype, elevatedcurvetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (elevatedcurvetype *ElevatedCurveType) GongGetReferenceIdentifier(stage *Stage) string {
	return elevatedcurvetype.GongGetIdentifier(stage)
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(elevatedsurfacetype, elevatedsurfacetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (elevatedsurfacetype *ElevatedSurfaceType) GongGetReferenceIdentifier(stage *Stage) string {
	return elevatedsurfacetype.GongGetIdentifier(stage)
}

func (geobordertimeslicepropertytype *GeoBorderTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(geobordertimeslicepropertytype, geobordertimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (geobordertimeslicepropertytype *GeoBorderTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return geobordertimeslicepropertytype.GongGetIdentifier(stage)
}

func (geobordertimeslicetype *GeoBorderTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(geobordertimeslicetype, geobordertimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (geobordertimeslicetype *GeoBorderTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return geobordertimeslicetype.GongGetIdentifier(stage)
}

func (geobordertype *GeoBorderType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(geobordertype, geobordertype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (geobordertype *GeoBorderType) GongGetReferenceIdentifier(stage *Stage) string {
	return geobordertype.GongGetIdentifier(stage)
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(linguisticnotepropertytype, linguisticnotepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (linguisticnotepropertytype *LinguisticNotePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return linguisticnotepropertytype.GongGetIdentifier(stage)
}

func (linguisticnotetype *LinguisticNoteType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(linguisticnotetype, linguisticnotetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (linguisticnotetype *LinguisticNoteType) GongGetReferenceIdentifier(stage *Stage) string {
	return linguisticnotetype.GongGetIdentifier(stage)
}

func (notepropertytype *NotePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(notepropertytype, notepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notepropertytype *NotePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return notepropertytype.GongGetIdentifier(stage)
}

func (notetype *NoteType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(notetype, notetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notetype *NoteType) GongGetReferenceIdentifier(stage *Stage) string {
	return notetype.GongGetIdentifier(stage)
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(organisationauthoritypropertytype, organisationauthoritypropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return organisationauthoritypropertytype.GongGetIdentifier(stage)
}

func (routepropertytype *RoutePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(routepropertytype, routepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (routepropertytype *RoutePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return routepropertytype.GongGetIdentifier(stage)
}

func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(standardlevelcolumnpropertytype, standardlevelcolumnpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return standardlevelcolumnpropertytype.GongGetIdentifier(stage)
}

func (surfacepropertytype *SurfacePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(surfacepropertytype, surfacepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (surfacepropertytype *SurfacePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return surfacepropertytype.GongGetIdentifier(stage)
}

func (surfacetype *SurfaceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(surfacetype, surfacetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (surfacetype *SurfaceType) GongGetReferenceIdentifier(stage *Stage) string {
	return surfacetype.GongGetIdentifier(stage)
}

func (timesheetpropertytype *TimesheetPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(timesheetpropertytype, timesheetpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (timesheetpropertytype *TimesheetPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return timesheetpropertytype.GongGetIdentifier(stage)
}

func (timesheettype *TimesheetType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(timesheettype, timesheettype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (timesheettype *TimesheetType) GongGetReferenceIdentifier(stage *Stage) string {
	return timesheettype.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (a_extension *A_extension) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension.GongGetIdentifier(stage), "A_extension", a_extension.Name)
}

func (a_extension_1 *A_extension_1) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_1.GongGetIdentifier(stage), "A_extension_1", a_extension_1.Name)
}

func (a_extension_10 *A_extension_10) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_10.GongGetIdentifier(stage), "A_extension_10", a_extension_10.Name)
}

func (a_extension_11 *A_extension_11) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_11.GongGetIdentifier(stage), "A_extension_11", a_extension_11.Name)
}

func (a_extension_12 *A_extension_12) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_12.GongGetIdentifier(stage), "A_extension_12", a_extension_12.Name)
}

func (a_extension_13 *A_extension_13) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_13.GongGetIdentifier(stage), "A_extension_13", a_extension_13.Name)
}

func (a_extension_14 *A_extension_14) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_14.GongGetIdentifier(stage), "A_extension_14", a_extension_14.Name)
}

func (a_extension_2 *A_extension_2) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_2.GongGetIdentifier(stage), "A_extension_2", a_extension_2.Name)
}

func (a_extension_3 *A_extension_3) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_3.GongGetIdentifier(stage), "A_extension_3", a_extension_3.Name)
}

func (a_extension_4 *A_extension_4) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_4.GongGetIdentifier(stage), "A_extension_4", a_extension_4.Name)
}

func (a_extension_5 *A_extension_5) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_5.GongGetIdentifier(stage), "A_extension_5", a_extension_5.Name)
}

func (a_extension_6 *A_extension_6) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_6.GongGetIdentifier(stage), "A_extension_6", a_extension_6.Name)
}

func (a_extension_7 *A_extension_7) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_7.GongGetIdentifier(stage), "A_extension_7", a_extension_7.Name)
}

func (a_extension_8 *A_extension_8) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_8.GongGetIdentifier(stage), "A_extension_8", a_extension_8.Name)
}

func (a_extension_9 *A_extension_9) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_9.GongGetIdentifier(stage), "A_extension_9", a_extension_9.Name)
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractpropertieswithscheduletype.GongGetIdentifier(stage), "AbstractPropertiesWithScheduleType", abstractpropertieswithscheduletype.Name)
}

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(aircraftcharacteristicpropertytype.GongGetIdentifier(stage), "AircraftCharacteristicPropertyType", aircraftcharacteristicpropertytype.Name)
}

func (aircraftcharacteristictype *AircraftCharacteristicType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(aircraftcharacteristictype.GongGetIdentifier(stage), "AircraftCharacteristicType", aircraftcharacteristictype.Name)
}

func (airspaceactivationpropertytype *AirspaceActivationPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airspaceactivationpropertytype.GongGetIdentifier(stage), "AirspaceActivationPropertyType", airspaceactivationpropertytype.Name)
}

func (airspaceactivationtype *AirspaceActivationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airspaceactivationtype.GongGetIdentifier(stage), "AirspaceActivationType", airspaceactivationtype.Name)
}

func (airspacegeometrycomponentpropertytype *AirspaceGeometryComponentPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airspacegeometrycomponentpropertytype.GongGetIdentifier(stage), "AirspaceGeometryComponentPropertyType", airspacegeometrycomponentpropertytype.Name)
}

func (airspacegeometrycomponenttype *AirspaceGeometryComponentType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airspacegeometrycomponenttype.GongGetIdentifier(stage), "AirspaceGeometryComponentType", airspacegeometrycomponenttype.Name)
}

func (airspacelayerclasspropertytype *AirspaceLayerClassPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airspacelayerclasspropertytype.GongGetIdentifier(stage), "AirspaceLayerClassPropertyType", airspacelayerclasspropertytype.Name)
}

func (airspacelayerclasstype *AirspaceLayerClassType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airspacelayerclasstype.GongGetIdentifier(stage), "AirspaceLayerClassType", airspacelayerclasstype.Name)
}

func (airspacelayerpropertytype *AirspaceLayerPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airspacelayerpropertytype.GongGetIdentifier(stage), "AirspaceLayerPropertyType", airspacelayerpropertytype.Name)
}

func (airspacelayertype *AirspaceLayerType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airspacelayertype.GongGetIdentifier(stage), "AirspaceLayerType", airspacelayertype.Name)
}

func (airspacepropertytype *AirspacePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airspacepropertytype.GongGetIdentifier(stage), "AirspacePropertyType", airspacepropertytype.Name)
}

func (airspacetimeslicepropertytype *AirspaceTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airspacetimeslicepropertytype.GongGetIdentifier(stage), "AirspaceTimeSlicePropertyType", airspacetimeslicepropertytype.Name)
}

func (airspacetimeslicetype *AirspaceTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airspacetimeslicetype.GongGetIdentifier(stage), "AirspaceTimeSliceType", airspacetimeslicetype.Name)
}

func (airspacetype *AirspaceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airspacetype.GongGetIdentifier(stage), "AirspaceType", airspacetype.Name)
}

func (airspacevolumedependencypropertytype *AirspaceVolumeDependencyPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airspacevolumedependencypropertytype.GongGetIdentifier(stage), "AirspaceVolumeDependencyPropertyType", airspacevolumedependencypropertytype.Name)
}

func (airspacevolumedependencytype *AirspaceVolumeDependencyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airspacevolumedependencytype.GongGetIdentifier(stage), "AirspaceVolumeDependencyType", airspacevolumedependencytype.Name)
}

func (airspacevolumepropertytype *AirspaceVolumePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airspacevolumepropertytype.GongGetIdentifier(stage), "AirspaceVolumePropertyType", airspacevolumepropertytype.Name)
}

func (airspacevolumetype *AirspaceVolumeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airspacevolumetype.GongGetIdentifier(stage), "AirspaceVolumeType", airspacevolumetype.Name)
}

func (authorityforairspacetimeslicepropertytype *AuthorityForAirspaceTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(authorityforairspacetimeslicepropertytype.GongGetIdentifier(stage), "AuthorityForAirspaceTimeSlicePropertyType", authorityforairspacetimeslicepropertytype.Name)
}

func (authorityforairspacetimeslicetype *AuthorityForAirspaceTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(authorityforairspacetimeslicetype.GongGetIdentifier(stage), "AuthorityForAirspaceTimeSliceType", authorityforairspacetimeslicetype.Name)
}

func (authorityforairspacetype *AuthorityForAirspaceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(authorityforairspacetype.GongGetIdentifier(stage), "AuthorityForAirspaceType", authorityforairspacetype.Name)
}

func (curvepropertytype *CurvePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(curvepropertytype.GongGetIdentifier(stage), "CurvePropertyType", curvepropertytype.Name)
}

func (curvetype *CurveType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(curvetype.GongGetIdentifier(stage), "CurveType", curvetype.Name)
}

func (elevatedcurvetype *ElevatedCurveType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(elevatedcurvetype.GongGetIdentifier(stage), "ElevatedCurveType", elevatedcurvetype.Name)
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(elevatedsurfacetype.GongGetIdentifier(stage), "ElevatedSurfaceType", elevatedsurfacetype.Name)
}

func (geobordertimeslicepropertytype *GeoBorderTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(geobordertimeslicepropertytype.GongGetIdentifier(stage), "GeoBorderTimeSlicePropertyType", geobordertimeslicepropertytype.Name)
}

func (geobordertimeslicetype *GeoBorderTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(geobordertimeslicetype.GongGetIdentifier(stage), "GeoBorderTimeSliceType", geobordertimeslicetype.Name)
}

func (geobordertype *GeoBorderType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(geobordertype.GongGetIdentifier(stage), "GeoBorderType", geobordertype.Name)
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(linguisticnotepropertytype.GongGetIdentifier(stage), "LinguisticNotePropertyType", linguisticnotepropertytype.Name)
}

func (linguisticnotetype *LinguisticNoteType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(linguisticnotetype.GongGetIdentifier(stage), "LinguisticNoteType", linguisticnotetype.Name)
}

func (notepropertytype *NotePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(notepropertytype.GongGetIdentifier(stage), "NotePropertyType", notepropertytype.Name)
}

func (notetype *NoteType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(notetype.GongGetIdentifier(stage), "NoteType", notetype.Name)
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(organisationauthoritypropertytype.GongGetIdentifier(stage), "OrganisationAuthorityPropertyType", organisationauthoritypropertytype.Name)
}

func (routepropertytype *RoutePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(routepropertytype.GongGetIdentifier(stage), "RoutePropertyType", routepropertytype.Name)
}

func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(standardlevelcolumnpropertytype.GongGetIdentifier(stage), "StandardLevelColumnPropertyType", standardlevelcolumnpropertytype.Name)
}

func (surfacepropertytype *SurfacePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(surfacepropertytype.GongGetIdentifier(stage), "SurfacePropertyType", surfacepropertytype.Name)
}

func (surfacetype *SurfaceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(surfacetype.GongGetIdentifier(stage), "SurfaceType", surfacetype.Name)
}

func (timesheetpropertytype *TimesheetPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(timesheetpropertytype.GongGetIdentifier(stage), "TimesheetPropertyType", timesheetpropertytype.Name)
}

func (timesheettype *TimesheetType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(timesheettype.GongGetIdentifier(stage), "TimesheetType", timesheettype.Name)
}

// insertion point for unstaging
func (a_extension *A_extension) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension.GongGetReferenceIdentifier(stage))
}

func (a_extension_1 *A_extension_1) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_1.GongGetReferenceIdentifier(stage))
}

func (a_extension_10 *A_extension_10) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_10.GongGetReferenceIdentifier(stage))
}

func (a_extension_11 *A_extension_11) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_11.GongGetReferenceIdentifier(stage))
}

func (a_extension_12 *A_extension_12) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_12.GongGetReferenceIdentifier(stage))
}

func (a_extension_13 *A_extension_13) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_13.GongGetReferenceIdentifier(stage))
}

func (a_extension_14 *A_extension_14) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_14.GongGetReferenceIdentifier(stage))
}

func (a_extension_2 *A_extension_2) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_2.GongGetReferenceIdentifier(stage))
}

func (a_extension_3 *A_extension_3) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_3.GongGetReferenceIdentifier(stage))
}

func (a_extension_4 *A_extension_4) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_4.GongGetReferenceIdentifier(stage))
}

func (a_extension_5 *A_extension_5) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_5.GongGetReferenceIdentifier(stage))
}

func (a_extension_6 *A_extension_6) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_6.GongGetReferenceIdentifier(stage))
}

func (a_extension_7 *A_extension_7) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_7.GongGetReferenceIdentifier(stage))
}

func (a_extension_8 *A_extension_8) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_8.GongGetReferenceIdentifier(stage))
}

func (a_extension_9 *A_extension_9) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_9.GongGetReferenceIdentifier(stage))
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractpropertieswithscheduletype.GongGetReferenceIdentifier(stage))
}

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(aircraftcharacteristicpropertytype.GongGetReferenceIdentifier(stage))
}

func (aircraftcharacteristictype *AircraftCharacteristicType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(aircraftcharacteristictype.GongGetReferenceIdentifier(stage))
}

func (airspaceactivationpropertytype *AirspaceActivationPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airspaceactivationpropertytype.GongGetReferenceIdentifier(stage))
}

func (airspaceactivationtype *AirspaceActivationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airspaceactivationtype.GongGetReferenceIdentifier(stage))
}

func (airspacegeometrycomponentpropertytype *AirspaceGeometryComponentPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airspacegeometrycomponentpropertytype.GongGetReferenceIdentifier(stage))
}

func (airspacegeometrycomponenttype *AirspaceGeometryComponentType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airspacegeometrycomponenttype.GongGetReferenceIdentifier(stage))
}

func (airspacelayerclasspropertytype *AirspaceLayerClassPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airspacelayerclasspropertytype.GongGetReferenceIdentifier(stage))
}

func (airspacelayerclasstype *AirspaceLayerClassType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airspacelayerclasstype.GongGetReferenceIdentifier(stage))
}

func (airspacelayerpropertytype *AirspaceLayerPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airspacelayerpropertytype.GongGetReferenceIdentifier(stage))
}

func (airspacelayertype *AirspaceLayerType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airspacelayertype.GongGetReferenceIdentifier(stage))
}

func (airspacepropertytype *AirspacePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airspacepropertytype.GongGetReferenceIdentifier(stage))
}

func (airspacetimeslicepropertytype *AirspaceTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airspacetimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (airspacetimeslicetype *AirspaceTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airspacetimeslicetype.GongGetReferenceIdentifier(stage))
}

func (airspacetype *AirspaceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airspacetype.GongGetReferenceIdentifier(stage))
}

func (airspacevolumedependencypropertytype *AirspaceVolumeDependencyPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airspacevolumedependencypropertytype.GongGetReferenceIdentifier(stage))
}

func (airspacevolumedependencytype *AirspaceVolumeDependencyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airspacevolumedependencytype.GongGetReferenceIdentifier(stage))
}

func (airspacevolumepropertytype *AirspaceVolumePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airspacevolumepropertytype.GongGetReferenceIdentifier(stage))
}

func (airspacevolumetype *AirspaceVolumeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airspacevolumetype.GongGetReferenceIdentifier(stage))
}

func (authorityforairspacetimeslicepropertytype *AuthorityForAirspaceTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(authorityforairspacetimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (authorityforairspacetimeslicetype *AuthorityForAirspaceTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(authorityforairspacetimeslicetype.GongGetReferenceIdentifier(stage))
}

func (authorityforairspacetype *AuthorityForAirspaceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(authorityforairspacetype.GongGetReferenceIdentifier(stage))
}

func (curvepropertytype *CurvePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(curvepropertytype.GongGetReferenceIdentifier(stage))
}

func (curvetype *CurveType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(curvetype.GongGetReferenceIdentifier(stage))
}

func (elevatedcurvetype *ElevatedCurveType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(elevatedcurvetype.GongGetReferenceIdentifier(stage))
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(elevatedsurfacetype.GongGetReferenceIdentifier(stage))
}

func (geobordertimeslicepropertytype *GeoBorderTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(geobordertimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (geobordertimeslicetype *GeoBorderTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(geobordertimeslicetype.GongGetReferenceIdentifier(stage))
}

func (geobordertype *GeoBorderType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(geobordertype.GongGetReferenceIdentifier(stage))
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(linguisticnotepropertytype.GongGetReferenceIdentifier(stage))
}

func (linguisticnotetype *LinguisticNoteType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(linguisticnotetype.GongGetReferenceIdentifier(stage))
}

func (notepropertytype *NotePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(notepropertytype.GongGetReferenceIdentifier(stage))
}

func (notetype *NoteType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(notetype.GongGetReferenceIdentifier(stage))
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(organisationauthoritypropertytype.GongGetReferenceIdentifier(stage))
}

func (routepropertytype *RoutePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(routepropertytype.GongGetReferenceIdentifier(stage))
}

func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(standardlevelcolumnpropertytype.GongGetReferenceIdentifier(stage))
}

func (surfacepropertytype *SurfacePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(surfacepropertytype.GongGetReferenceIdentifier(stage))
}

func (surfacetype *SurfaceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(surfacetype.GongGetReferenceIdentifier(stage))
}

func (timesheetpropertytype *TimesheetPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(timesheetpropertytype.GongGetReferenceIdentifier(stage))
}

func (timesheettype *TimesheetType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(timesheettype.GongGetReferenceIdentifier(stage))
}

func GongIntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += GongIntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GongGenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GongGenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
	// 1. Create a deterministic hash from the inputs using SHA-256
	h := sha256.New()

	// Write the string to the hash
	h.Write([]byte(seedStr))

	// Write the integer to the hash (using BigEndian to ensure consistency across architectures)
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, seedInt)
	h.Write(intBytes)

	// 2. Extract the first 16 bytes from our resulting hash
	hashBytes := h.Sum(nil)
	uuid := make([]byte, 16)
	copy(uuid, hashBytes[:16])

	// 3. Set the Version to 4 (0100 in binary)
	// We take the 7th byte, clear the top 4 bits with & 0x0f, and set the top bits to 0100 with | 0x40
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// 4. Set the Variant to RFC4122 (10 in binary)
	// We take the 9th byte, clear the top 2 bits with & 0x3f, and set the top bits to 10 with | 0x80
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// 5. Format and return the byte array as a standard UUID string
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

func __gong__appendInstances[T interface {
	comparable
	GongstructIF
}](res []GongstructIF, m map[T]struct{}) []GongstructIF {
	for instance := range m {
		res = append(res, instance)
	}
	return res
}

func __gong__getUUID(stage *Stage, instance GongstructIF) string {
	if __gong__, ok := any(instance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}
	return GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(instance), uint64(stage.GetOrder(instance)))
}

func __gong__computeReferencePass1[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	staged map[T]struct{},
	ref *map[T]T,
	refOrder *map[T]uint,
	inst *map[T]T,
) {
	*ref = make(map[T]T, len(staged))
	*refOrder = make(map[T]uint, len(staged))
	*inst = make(map[T]T, len(staged))
	for instance := range staged {
		_copy := instance.GongCopy().(T)
		(*ref)[instance] = _copy
		(*inst)[_copy] = instance
		(*refOrder)[_copy] = instance.GongGetOrder(stage)
	}
}

func __gong__computeReferencePass2[T interface {
	comparable
	GongstructIF
	GongReconstructPointersFromReferences(*Stage, T)
}](staged map[T]struct{}, reference map[T]T, stage *Stage) {
	for instance := range staged {
		reference[instance].GongReconstructPointersFromReferences(stage, instance)
	}
}

func __gong__getOrder[T comparable](stagedOrder, refOrder map[T]uint, instance T, typeName string) uint {
	if order, ok := stagedOrder[instance]; ok {
		return order
	}
	if order, ok := refOrder[instance]; ok {
		return order
	}
	log.Printf("instance %p of type %s was not staged and does not have a reference order", any(instance), typeName)
	return 0
}

func __gong__formatIdentifier(s GongstructIF, order uint) string {
	return fmt.Sprintf("__aixmfeatures__%s__%08d_", s.GongGetGongstructName(), order)
}

func __gong__marshallIdentifier(identifier, structName, name string) string {
	decl := strings.ReplaceAll(GongIdentifiersDecls, "{{Identifier}}", identifier)
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", structName)
	return strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(name))
}

func __gong__marshallUnstaging(identifier string) string {
	return strings.ReplaceAll(GongUnstageStmt, "{{Identifier}}", identifier)
}

// end of template
