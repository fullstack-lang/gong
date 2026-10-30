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
	// Compute reverse map for named struct AerialRefuellingAnchorType
	// insertion point per field
	stage.AerialRefuellingAnchorType_VerticalExtent_reverseMap = make(map[*AirspaceLayerPropertyType]*AerialRefuellingAnchorType)
	for aerialrefuellinganchortype := range stage.AerialRefuellingAnchorTypes {
		_ = aerialrefuellinganchortype
		for _, _airspacelayerpropertytype := range aerialrefuellinganchortype.VerticalExtent {
			stage.AerialRefuellingAnchorType_VerticalExtent_reverseMap[_airspacelayerpropertytype] = aerialrefuellinganchortype
		}
	}
	stage.AerialRefuellingAnchorType_Point_reverseMap = make(map[*AerialRefuellingPointPropertyType]*AerialRefuellingAnchorType)
	for aerialrefuellinganchortype := range stage.AerialRefuellingAnchorTypes {
		_ = aerialrefuellinganchortype
		for _, _aerialrefuellingpointpropertytype := range aerialrefuellinganchortype.Point {
			stage.AerialRefuellingAnchorType_Point_reverseMap[_aerialrefuellingpointpropertytype] = aerialrefuellinganchortype
		}
	}
	stage.AerialRefuellingAnchorType_Annotation_reverseMap = make(map[*NotePropertyType]*AerialRefuellingAnchorType)
	for aerialrefuellinganchortype := range stage.AerialRefuellingAnchorTypes {
		_ = aerialrefuellinganchortype
		for _, _notepropertytype := range aerialrefuellinganchortype.Annotation {
			stage.AerialRefuellingAnchorType_Annotation_reverseMap[_notepropertytype] = aerialrefuellinganchortype
		}
	}
	stage.AerialRefuellingAnchorType_Extension_reverseMap = make(map[*A_extension_1]*AerialRefuellingAnchorType)
	for aerialrefuellinganchortype := range stage.AerialRefuellingAnchorTypes {
		_ = aerialrefuellinganchortype
		for _, _a_extension_1 := range aerialrefuellinganchortype.Extension {
			stage.AerialRefuellingAnchorType_Extension_reverseMap[_a_extension_1] = aerialrefuellinganchortype
		}
	}

	// Compute reverse map for named struct AerialRefuellingPointType
	// insertion point per field
	stage.AerialRefuellingPointType_FacilityMakeup_reverseMap = make(map[*PointReferencePropertyType]*AerialRefuellingPointType)
	for aerialrefuellingpointtype := range stage.AerialRefuellingPointTypes {
		_ = aerialrefuellingpointtype
		for _, _pointreferencepropertytype := range aerialrefuellingpointtype.FacilityMakeup {
			stage.AerialRefuellingPointType_FacilityMakeup_reverseMap[_pointreferencepropertytype] = aerialrefuellingpointtype
		}
	}
	stage.AerialRefuellingPointType_Annotation_reverseMap = make(map[*NotePropertyType]*AerialRefuellingPointType)
	for aerialrefuellingpointtype := range stage.AerialRefuellingPointTypes {
		_ = aerialrefuellingpointtype
		for _, _notepropertytype := range aerialrefuellingpointtype.Annotation {
			stage.AerialRefuellingPointType_Annotation_reverseMap[_notepropertytype] = aerialrefuellingpointtype
		}
	}
	stage.AerialRefuellingPointType_Extension_reverseMap = make(map[*A_extension_2]*AerialRefuellingPointType)
	for aerialrefuellingpointtype := range stage.AerialRefuellingPointTypes {
		_ = aerialrefuellingpointtype
		for _, _a_extension_2 := range aerialrefuellingpointtype.Extension {
			stage.AerialRefuellingPointType_Extension_reverseMap[_a_extension_2] = aerialrefuellingpointtype
		}
	}

	// Compute reverse map for named struct AerialRefuellingTimeSliceType
	// insertion point per field
	stage.AerialRefuellingTimeSliceType_Availability_reverseMap = make(map[*RouteAvailabilityPropertyType]*AerialRefuellingTimeSliceType)
	for aerialrefuellingtimeslicetype := range stage.AerialRefuellingTimeSliceTypes {
		_ = aerialrefuellingtimeslicetype
		for _, _routeavailabilitypropertytype := range aerialrefuellingtimeslicetype.Availability {
			stage.AerialRefuellingTimeSliceType_Availability_reverseMap[_routeavailabilitypropertytype] = aerialrefuellingtimeslicetype
		}
	}
	stage.AerialRefuellingTimeSliceType_ProtectingAirspace_reverseMap = make(map[*AirspacePropertyType]*AerialRefuellingTimeSliceType)
	for aerialrefuellingtimeslicetype := range stage.AerialRefuellingTimeSliceTypes {
		_ = aerialrefuellingtimeslicetype
		for _, _airspacepropertytype := range aerialrefuellingtimeslicetype.ProtectingAirspace {
			stage.AerialRefuellingTimeSliceType_ProtectingAirspace_reverseMap[_airspacepropertytype] = aerialrefuellingtimeslicetype
		}
	}
	stage.AerialRefuellingTimeSliceType_Track_reverseMap = make(map[*AerialRefuellingTrackPropertyType]*AerialRefuellingTimeSliceType)
	for aerialrefuellingtimeslicetype := range stage.AerialRefuellingTimeSliceTypes {
		_ = aerialrefuellingtimeslicetype
		for _, _aerialrefuellingtrackpropertytype := range aerialrefuellingtimeslicetype.Track {
			stage.AerialRefuellingTimeSliceType_Track_reverseMap[_aerialrefuellingtrackpropertytype] = aerialrefuellingtimeslicetype
		}
	}
	stage.AerialRefuellingTimeSliceType_Anchor_reverseMap = make(map[*AerialRefuellingAnchorPropertyType]*AerialRefuellingTimeSliceType)
	for aerialrefuellingtimeslicetype := range stage.AerialRefuellingTimeSliceTypes {
		_ = aerialrefuellingtimeslicetype
		for _, _aerialrefuellinganchorpropertytype := range aerialrefuellingtimeslicetype.Anchor {
			stage.AerialRefuellingTimeSliceType_Anchor_reverseMap[_aerialrefuellinganchorpropertytype] = aerialrefuellingtimeslicetype
		}
	}
	stage.AerialRefuellingTimeSliceType_ManagingOrganisation_reverseMap = make(map[*AuthorityForAerialRefuellingPropertyType]*AerialRefuellingTimeSliceType)
	for aerialrefuellingtimeslicetype := range stage.AerialRefuellingTimeSliceTypes {
		_ = aerialrefuellingtimeslicetype
		for _, _authorityforaerialrefuellingpropertytype := range aerialrefuellingtimeslicetype.ManagingOrganisation {
			stage.AerialRefuellingTimeSliceType_ManagingOrganisation_reverseMap[_authorityforaerialrefuellingpropertytype] = aerialrefuellingtimeslicetype
		}
	}
	stage.AerialRefuellingTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*AerialRefuellingTimeSliceType)
	for aerialrefuellingtimeslicetype := range stage.AerialRefuellingTimeSliceTypes {
		_ = aerialrefuellingtimeslicetype
		for _, _notepropertytype := range aerialrefuellingtimeslicetype.Annotation {
			stage.AerialRefuellingTimeSliceType_Annotation_reverseMap[_notepropertytype] = aerialrefuellingtimeslicetype
		}
	}
	stage.AerialRefuellingTimeSliceType_Extension_reverseMap = make(map[*A_extension]*AerialRefuellingTimeSliceType)
	for aerialrefuellingtimeslicetype := range stage.AerialRefuellingTimeSliceTypes {
		_ = aerialrefuellingtimeslicetype
		for _, _a_extension := range aerialrefuellingtimeslicetype.Extension {
			stage.AerialRefuellingTimeSliceType_Extension_reverseMap[_a_extension] = aerialrefuellingtimeslicetype
		}
	}

	// Compute reverse map for named struct AerialRefuellingTrackType
	// insertion point per field
	stage.AerialRefuellingTrackType_Point_reverseMap = make(map[*AerialRefuellingPointPropertyType]*AerialRefuellingTrackType)
	for aerialrefuellingtracktype := range stage.AerialRefuellingTrackTypes {
		_ = aerialrefuellingtracktype
		for _, _aerialrefuellingpointpropertytype := range aerialrefuellingtracktype.Point {
			stage.AerialRefuellingTrackType_Point_reverseMap[_aerialrefuellingpointpropertytype] = aerialrefuellingtracktype
		}
	}
	stage.AerialRefuellingTrackType_VerticalExtent_reverseMap = make(map[*AirspaceLayerPropertyType]*AerialRefuellingTrackType)
	for aerialrefuellingtracktype := range stage.AerialRefuellingTrackTypes {
		_ = aerialrefuellingtracktype
		for _, _airspacelayerpropertytype := range aerialrefuellingtracktype.VerticalExtent {
			stage.AerialRefuellingTrackType_VerticalExtent_reverseMap[_airspacelayerpropertytype] = aerialrefuellingtracktype
		}
	}
	stage.AerialRefuellingTrackType_Annotation_reverseMap = make(map[*NotePropertyType]*AerialRefuellingTrackType)
	for aerialrefuellingtracktype := range stage.AerialRefuellingTrackTypes {
		_ = aerialrefuellingtracktype
		for _, _notepropertytype := range aerialrefuellingtracktype.Annotation {
			stage.AerialRefuellingTrackType_Annotation_reverseMap[_notepropertytype] = aerialrefuellingtracktype
		}
	}
	stage.AerialRefuellingTrackType_Extension_reverseMap = make(map[*A_extension_3]*AerialRefuellingTrackType)
	for aerialrefuellingtracktype := range stage.AerialRefuellingTrackTypes {
		_ = aerialrefuellingtracktype
		for _, _a_extension_3 := range aerialrefuellingtracktype.Extension {
			stage.AerialRefuellingTrackType_Extension_reverseMap[_a_extension_3] = aerialrefuellingtracktype
		}
	}

	// Compute reverse map for named struct AerialRefuellingType
	// insertion point per field
	stage.AerialRefuellingType_TimeSlice_reverseMap = make(map[*AerialRefuellingTimeSlicePropertyType]*AerialRefuellingType)
	for aerialrefuellingtype := range stage.AerialRefuellingTypes {
		_ = aerialrefuellingtype
		for _, _aerialrefuellingtimeslicepropertytype := range aerialrefuellingtype.TimeSlice {
			stage.AerialRefuellingType_TimeSlice_reverseMap[_aerialrefuellingtimeslicepropertytype] = aerialrefuellingtype
		}
	}

	// Compute reverse map for named struct AircraftCharacteristicType
	// insertion point per field
	stage.AircraftCharacteristicType_Annotation_reverseMap = make(map[*NotePropertyType]*AircraftCharacteristicType)
	for aircraftcharacteristictype := range stage.AircraftCharacteristicTypes {
		_ = aircraftcharacteristictype
		for _, _notepropertytype := range aircraftcharacteristictype.Annotation {
			stage.AircraftCharacteristicType_Annotation_reverseMap[_notepropertytype] = aircraftcharacteristictype
		}
	}
	stage.AircraftCharacteristicType_Extension_reverseMap = make(map[*A_extension_34]*AircraftCharacteristicType)
	for aircraftcharacteristictype := range stage.AircraftCharacteristicTypes {
		_ = aircraftcharacteristictype
		for _, _a_extension_34 := range aircraftcharacteristictype.Extension {
			stage.AircraftCharacteristicType_Extension_reverseMap[_a_extension_34] = aircraftcharacteristictype
		}
	}

	// Compute reverse map for named struct AirspaceBorderCrossingTimeSliceType
	// insertion point per field
	stage.AirspaceBorderCrossingTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*AirspaceBorderCrossingTimeSliceType)
	for airspacebordercrossingtimeslicetype := range stage.AirspaceBorderCrossingTimeSliceTypes {
		_ = airspacebordercrossingtimeslicetype
		for _, _notepropertytype := range airspacebordercrossingtimeslicetype.Annotation {
			stage.AirspaceBorderCrossingTimeSliceType_Annotation_reverseMap[_notepropertytype] = airspacebordercrossingtimeslicetype
		}
	}
	stage.AirspaceBorderCrossingTimeSliceType_Extension_reverseMap = make(map[*A_extension_19]*AirspaceBorderCrossingTimeSliceType)
	for airspacebordercrossingtimeslicetype := range stage.AirspaceBorderCrossingTimeSliceTypes {
		_ = airspacebordercrossingtimeslicetype
		for _, _a_extension_19 := range airspacebordercrossingtimeslicetype.Extension {
			stage.AirspaceBorderCrossingTimeSliceType_Extension_reverseMap[_a_extension_19] = airspacebordercrossingtimeslicetype
		}
	}

	// Compute reverse map for named struct AirspaceBorderCrossingType
	// insertion point per field
	stage.AirspaceBorderCrossingType_TimeSlice_reverseMap = make(map[*AirspaceBorderCrossingTimeSlicePropertyType]*AirspaceBorderCrossingType)
	for airspacebordercrossingtype := range stage.AirspaceBorderCrossingTypes {
		_ = airspacebordercrossingtype
		for _, _airspacebordercrossingtimeslicepropertytype := range airspacebordercrossingtype.TimeSlice {
			stage.AirspaceBorderCrossingType_TimeSlice_reverseMap[_airspacebordercrossingtimeslicepropertytype] = airspacebordercrossingtype
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
	stage.AirspaceLayerType_Extension_reverseMap = make(map[*A_extension_36]*AirspaceLayerType)
	for airspacelayertype := range stage.AirspaceLayerTypes {
		_ = airspacelayertype
		for _, _a_extension_36 := range airspacelayertype.Extension {
			stage.AirspaceLayerType_Extension_reverseMap[_a_extension_36] = airspacelayertype
		}
	}

	// Compute reverse map for named struct AltitudeAdjustmentType
	// insertion point per field
	stage.AltitudeAdjustmentType_Annotation_reverseMap = make(map[*NotePropertyType]*AltitudeAdjustmentType)
	for altitudeadjustmenttype := range stage.AltitudeAdjustmentTypes {
		_ = altitudeadjustmenttype
		for _, _notepropertytype := range altitudeadjustmenttype.Annotation {
			stage.AltitudeAdjustmentType_Annotation_reverseMap[_notepropertytype] = altitudeadjustmenttype
		}
	}
	stage.AltitudeAdjustmentType_Extension_reverseMap = make(map[*A_extension_39]*AltitudeAdjustmentType)
	for altitudeadjustmenttype := range stage.AltitudeAdjustmentTypes {
		_ = altitudeadjustmenttype
		for _, _a_extension_39 := range altitudeadjustmenttype.Extension {
			stage.AltitudeAdjustmentType_Extension_reverseMap[_a_extension_39] = altitudeadjustmenttype
		}
	}

	// Compute reverse map for named struct AngleUseType
	// insertion point per field
	stage.AngleUseType_Annotation_reverseMap = make(map[*NotePropertyType]*AngleUseType)
	for angleusetype := range stage.AngleUseTypes {
		_ = angleusetype
		for _, _notepropertytype := range angleusetype.Annotation {
			stage.AngleUseType_Annotation_reverseMap[_notepropertytype] = angleusetype
		}
	}
	stage.AngleUseType_Extension_reverseMap = make(map[*A_extension_8]*AngleUseType)
	for angleusetype := range stage.AngleUseTypes {
		_ = angleusetype
		for _, _a_extension_8 := range angleusetype.Extension {
			stage.AngleUseType_Extension_reverseMap[_a_extension_8] = angleusetype
		}
	}

	// Compute reverse map for named struct AuthorityForAerialRefuellingType
	// insertion point per field
	stage.AuthorityForAerialRefuellingType_Annotation_reverseMap = make(map[*NotePropertyType]*AuthorityForAerialRefuellingType)
	for authorityforaerialrefuellingtype := range stage.AuthorityForAerialRefuellingTypes {
		_ = authorityforaerialrefuellingtype
		for _, _notepropertytype := range authorityforaerialrefuellingtype.Annotation {
			stage.AuthorityForAerialRefuellingType_Annotation_reverseMap[_notepropertytype] = authorityforaerialrefuellingtype
		}
	}
	stage.AuthorityForAerialRefuellingType_Extension_reverseMap = make(map[*A_extension_4]*AuthorityForAerialRefuellingType)
	for authorityforaerialrefuellingtype := range stage.AuthorityForAerialRefuellingTypes {
		_ = authorityforaerialrefuellingtype
		for _, _a_extension_4 := range authorityforaerialrefuellingtype.Extension {
			stage.AuthorityForAerialRefuellingType_Extension_reverseMap[_a_extension_4] = authorityforaerialrefuellingtype
		}
	}

	// Compute reverse map for named struct ChangeOverPointTimeSliceType
	// insertion point per field
	stage.ChangeOverPointTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*ChangeOverPointTimeSliceType)
	for changeoverpointtimeslicetype := range stage.ChangeOverPointTimeSliceTypes {
		_ = changeoverpointtimeslicetype
		for _, _notepropertytype := range changeoverpointtimeslicetype.Annotation {
			stage.ChangeOverPointTimeSliceType_Annotation_reverseMap[_notepropertytype] = changeoverpointtimeslicetype
		}
	}
	stage.ChangeOverPointTimeSliceType_Extension_reverseMap = make(map[*A_extension_13]*ChangeOverPointTimeSliceType)
	for changeoverpointtimeslicetype := range stage.ChangeOverPointTimeSliceTypes {
		_ = changeoverpointtimeslicetype
		for _, _a_extension_13 := range changeoverpointtimeslicetype.Extension {
			stage.ChangeOverPointTimeSliceType_Extension_reverseMap[_a_extension_13] = changeoverpointtimeslicetype
		}
	}

	// Compute reverse map for named struct ChangeOverPointType
	// insertion point per field
	stage.ChangeOverPointType_TimeSlice_reverseMap = make(map[*ChangeOverPointTimeSlicePropertyType]*ChangeOverPointType)
	for changeoverpointtype := range stage.ChangeOverPointTypes {
		_ = changeoverpointtype
		for _, _changeoverpointtimeslicepropertytype := range changeoverpointtype.TimeSlice {
			stage.ChangeOverPointType_TimeSlice_reverseMap[_changeoverpointtimeslicepropertytype] = changeoverpointtype
		}
	}

	// Compute reverse map for named struct ContactInformationType
	// insertion point per field
	stage.ContactInformationType_Address_reverseMap = make(map[*PostalAddressPropertyType]*ContactInformationType)
	for contactinformationtype := range stage.ContactInformationTypes {
		_ = contactinformationtype
		for _, _postaladdresspropertytype := range contactinformationtype.Address {
			stage.ContactInformationType_Address_reverseMap[_postaladdresspropertytype] = contactinformationtype
		}
	}
	stage.ContactInformationType_NetworkNode_reverseMap = make(map[*OnlineContactPropertyType]*ContactInformationType)
	for contactinformationtype := range stage.ContactInformationTypes {
		_ = contactinformationtype
		for _, _onlinecontactpropertytype := range contactinformationtype.NetworkNode {
			stage.ContactInformationType_NetworkNode_reverseMap[_onlinecontactpropertytype] = contactinformationtype
		}
	}
	stage.ContactInformationType_PhoneFax_reverseMap = make(map[*TelephoneContactPropertyType]*ContactInformationType)
	for contactinformationtype := range stage.ContactInformationTypes {
		_ = contactinformationtype
		for _, _telephonecontactpropertytype := range contactinformationtype.PhoneFax {
			stage.ContactInformationType_PhoneFax_reverseMap[_telephonecontactpropertytype] = contactinformationtype
		}
	}
	stage.ContactInformationType_Annotation_reverseMap = make(map[*NotePropertyType]*ContactInformationType)
	for contactinformationtype := range stage.ContactInformationTypes {
		_ = contactinformationtype
		for _, _notepropertytype := range contactinformationtype.Annotation {
			stage.ContactInformationType_Annotation_reverseMap[_notepropertytype] = contactinformationtype
		}
	}
	stage.ContactInformationType_Extension_reverseMap = make(map[*A_extension_30]*ContactInformationType)
	for contactinformationtype := range stage.ContactInformationTypes {
		_ = contactinformationtype
		for _, _a_extension_30 := range contactinformationtype.Extension {
			stage.ContactInformationType_Extension_reverseMap[_a_extension_30] = contactinformationtype
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

	// Compute reverse map for named struct DirectFlightClassType
	// insertion point per field
	stage.DirectFlightClassType_Annotation_reverseMap = make(map[*NotePropertyType]*DirectFlightClassType)
	for directflightclasstype := range stage.DirectFlightClassTypes {
		_ = directflightclasstype
		for _, _notepropertytype := range directflightclasstype.Annotation {
			stage.DirectFlightClassType_Annotation_reverseMap[_notepropertytype] = directflightclasstype
		}
	}
	stage.DirectFlightClassType_Extension_reverseMap = make(map[*A_extension_20]*DirectFlightClassType)
	for directflightclasstype := range stage.DirectFlightClassTypes {
		_ = directflightclasstype
		for _, _a_extension_20 := range directflightclasstype.Extension {
			stage.DirectFlightClassType_Extension_reverseMap[_a_extension_20] = directflightclasstype
		}
	}

	// Compute reverse map for named struct DirectFlightSegmentType
	// insertion point per field
	stage.DirectFlightSegmentType_Annotation_reverseMap = make(map[*NotePropertyType]*DirectFlightSegmentType)
	for directflightsegmenttype := range stage.DirectFlightSegmentTypes {
		_ = directflightsegmenttype
		for _, _notepropertytype := range directflightsegmenttype.Annotation {
			stage.DirectFlightSegmentType_Annotation_reverseMap[_notepropertytype] = directflightsegmenttype
		}
	}
	stage.DirectFlightSegmentType_Extension_reverseMap = make(map[*A_extension_21]*DirectFlightSegmentType)
	for directflightsegmenttype := range stage.DirectFlightSegmentTypes {
		_ = directflightsegmenttype
		for _, _a_extension_21 := range directflightsegmenttype.Extension {
			stage.DirectFlightSegmentType_Extension_reverseMap[_a_extension_21] = directflightsegmenttype
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
	stage.ElevatedCurveType_Extension_reverseMap = make(map[*A_extension_5]*ElevatedCurveType)
	for elevatedcurvetype := range stage.ElevatedCurveTypes {
		_ = elevatedcurvetype
		for _, _a_extension_5 := range elevatedcurvetype.Extension {
			stage.ElevatedCurveType_Extension_reverseMap[_a_extension_5] = elevatedcurvetype
		}
	}

	// Compute reverse map for named struct ElevatedPointType
	// insertion point per field
	stage.ElevatedPointType_Annotation_reverseMap = make(map[*NotePropertyType]*ElevatedPointType)
	for elevatedpointtype := range stage.ElevatedPointTypes {
		_ = elevatedpointtype
		for _, _notepropertytype := range elevatedpointtype.Annotation {
			stage.ElevatedPointType_Annotation_reverseMap[_notepropertytype] = elevatedpointtype
		}
	}
	stage.ElevatedPointType_Extension_reverseMap = make(map[*A_extension_6]*ElevatedPointType)
	for elevatedpointtype := range stage.ElevatedPointTypes {
		_ = elevatedpointtype
		for _, _a_extension_6 := range elevatedpointtype.Extension {
			stage.ElevatedPointType_Extension_reverseMap[_a_extension_6] = elevatedpointtype
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
	stage.ElevatedSurfaceType_Extension_reverseMap = make(map[*A_extension_7]*ElevatedSurfaceType)
	for elevatedsurfacetype := range stage.ElevatedSurfaceTypes {
		_ = elevatedsurfacetype
		for _, _a_extension_7 := range elevatedsurfacetype.Extension {
			stage.ElevatedSurfaceType_Extension_reverseMap[_a_extension_7] = elevatedsurfacetype
		}
	}

	// Compute reverse map for named struct EnRouteSegmentPointType
	// insertion point per field
	stage.EnRouteSegmentPointType_FacilityMakeup_reverseMap = make(map[*PointReferencePropertyType]*EnRouteSegmentPointType)
	for enroutesegmentpointtype := range stage.EnRouteSegmentPointTypes {
		_ = enroutesegmentpointtype
		for _, _pointreferencepropertytype := range enroutesegmentpointtype.FacilityMakeup {
			stage.EnRouteSegmentPointType_FacilityMakeup_reverseMap[_pointreferencepropertytype] = enroutesegmentpointtype
		}
	}
	stage.EnRouteSegmentPointType_Annotation_reverseMap = make(map[*NotePropertyType]*EnRouteSegmentPointType)
	for enroutesegmentpointtype := range stage.EnRouteSegmentPointTypes {
		_ = enroutesegmentpointtype
		for _, _notepropertytype := range enroutesegmentpointtype.Annotation {
			stage.EnRouteSegmentPointType_Annotation_reverseMap[_notepropertytype] = enroutesegmentpointtype
		}
	}
	stage.EnRouteSegmentPointType_Extension_reverseMap = make(map[*A_extension_9]*EnRouteSegmentPointType)
	for enroutesegmentpointtype := range stage.EnRouteSegmentPointTypes {
		_ = enroutesegmentpointtype
		for _, _a_extension_9 := range enroutesegmentpointtype.Extension {
			stage.EnRouteSegmentPointType_Extension_reverseMap[_a_extension_9] = enroutesegmentpointtype
		}
	}

	// Compute reverse map for named struct FlightCharacteristicType
	// insertion point per field
	stage.FlightCharacteristicType_Annotation_reverseMap = make(map[*NotePropertyType]*FlightCharacteristicType)
	for flightcharacteristictype := range stage.FlightCharacteristicTypes {
		_ = flightcharacteristictype
		for _, _notepropertytype := range flightcharacteristictype.Annotation {
			stage.FlightCharacteristicType_Annotation_reverseMap[_notepropertytype] = flightcharacteristictype
		}
	}
	stage.FlightCharacteristicType_Extension_reverseMap = make(map[*A_extension_35]*FlightCharacteristicType)
	for flightcharacteristictype := range stage.FlightCharacteristicTypes {
		_ = flightcharacteristictype
		for _, _a_extension_35 := range flightcharacteristictype.Extension {
			stage.FlightCharacteristicType_Extension_reverseMap[_a_extension_35] = flightcharacteristictype
		}
	}

	// Compute reverse map for named struct FlightConditionCircumstanceType
	// insertion point per field
	stage.FlightConditionCircumstanceType_Annotation_reverseMap = make(map[*NotePropertyType]*FlightConditionCircumstanceType)
	for flightconditioncircumstancetype := range stage.FlightConditionCircumstanceTypes {
		_ = flightconditioncircumstancetype
		for _, _notepropertytype := range flightconditioncircumstancetype.Annotation {
			stage.FlightConditionCircumstanceType_Annotation_reverseMap[_notepropertytype] = flightconditioncircumstancetype
		}
	}
	stage.FlightConditionCircumstanceType_Extension_reverseMap = make(map[*A_extension_22]*FlightConditionCircumstanceType)
	for flightconditioncircumstancetype := range stage.FlightConditionCircumstanceTypes {
		_ = flightconditioncircumstancetype
		for _, _a_extension_22 := range flightconditioncircumstancetype.Extension {
			stage.FlightConditionCircumstanceType_Extension_reverseMap[_a_extension_22] = flightconditioncircumstancetype
		}
	}

	// Compute reverse map for named struct FlightConditionCombinationType
	// insertion point per field
	stage.FlightConditionCombinationType_TimeInterval_reverseMap = make(map[*TimesheetPropertyType]*FlightConditionCombinationType)
	for flightconditioncombinationtype := range stage.FlightConditionCombinationTypes {
		_ = flightconditioncombinationtype
		for _, _timesheetpropertytype := range flightconditioncombinationtype.TimeInterval {
			stage.FlightConditionCombinationType_TimeInterval_reverseMap[_timesheetpropertytype] = flightconditioncombinationtype
		}
	}
	stage.FlightConditionCombinationType_Annotation_reverseMap = make(map[*NotePropertyType]*FlightConditionCombinationType)
	for flightconditioncombinationtype := range stage.FlightConditionCombinationTypes {
		_ = flightconditioncombinationtype
		for _, _notepropertytype := range flightconditioncombinationtype.Annotation {
			stage.FlightConditionCombinationType_Annotation_reverseMap[_notepropertytype] = flightconditioncombinationtype
		}
	}
	stage.FlightConditionCombinationType_SpecialDateAuthority_reverseMap = make(map[*OrganisationAuthorityPropertyType]*FlightConditionCombinationType)
	for flightconditioncombinationtype := range stage.FlightConditionCombinationTypes {
		_ = flightconditioncombinationtype
		for _, _organisationauthoritypropertytype := range flightconditioncombinationtype.SpecialDateAuthority {
			stage.FlightConditionCombinationType_SpecialDateAuthority_reverseMap[_organisationauthoritypropertytype] = flightconditioncombinationtype
		}
	}
	stage.FlightConditionCombinationType_Element_reverseMap = make(map[*FlightConditionElementPropertyType]*FlightConditionCombinationType)
	for flightconditioncombinationtype := range stage.FlightConditionCombinationTypes {
		_ = flightconditioncombinationtype
		for _, _flightconditionelementpropertytype := range flightconditioncombinationtype.Element {
			stage.FlightConditionCombinationType_Element_reverseMap[_flightconditionelementpropertytype] = flightconditioncombinationtype
		}
	}
	stage.FlightConditionCombinationType_Extension_reverseMap = make(map[*A_extension_23]*FlightConditionCombinationType)
	for flightconditioncombinationtype := range stage.FlightConditionCombinationTypes {
		_ = flightconditioncombinationtype
		for _, _a_extension_23 := range flightconditioncombinationtype.Extension {
			stage.FlightConditionCombinationType_Extension_reverseMap[_a_extension_23] = flightconditioncombinationtype
		}
	}

	// Compute reverse map for named struct FlightConditionElementType
	// insertion point per field
	stage.FlightConditionElementType_FlightLevel_reverseMap = make(map[*FlightRestrictionLevelPropertyType]*FlightConditionElementType)
	for flightconditionelementtype := range stage.FlightConditionElementTypes {
		_ = flightconditionelementtype
		for _, _flightrestrictionlevelpropertytype := range flightconditionelementtype.FlightLevel {
			stage.FlightConditionElementType_FlightLevel_reverseMap[_flightrestrictionlevelpropertytype] = flightconditionelementtype
		}
	}
	stage.FlightConditionElementType_Annotation_reverseMap = make(map[*NotePropertyType]*FlightConditionElementType)
	for flightconditionelementtype := range stage.FlightConditionElementTypes {
		_ = flightconditionelementtype
		for _, _notepropertytype := range flightconditionelementtype.Annotation {
			stage.FlightConditionElementType_Annotation_reverseMap[_notepropertytype] = flightconditionelementtype
		}
	}
	stage.FlightConditionElementType_Extension_reverseMap = make(map[*A_extension_24]*FlightConditionElementType)
	for flightconditionelementtype := range stage.FlightConditionElementTypes {
		_ = flightconditionelementtype
		for _, _a_extension_24 := range flightconditionelementtype.Extension {
			stage.FlightConditionElementType_Extension_reverseMap[_a_extension_24] = flightconditionelementtype
		}
	}

	// Compute reverse map for named struct FlightRestrictionLevelType
	// insertion point per field
	stage.FlightRestrictionLevelType_Annotation_reverseMap = make(map[*NotePropertyType]*FlightRestrictionLevelType)
	for flightrestrictionleveltype := range stage.FlightRestrictionLevelTypes {
		_ = flightrestrictionleveltype
		for _, _notepropertytype := range flightrestrictionleveltype.Annotation {
			stage.FlightRestrictionLevelType_Annotation_reverseMap[_notepropertytype] = flightrestrictionleveltype
		}
	}
	stage.FlightRestrictionLevelType_Extension_reverseMap = make(map[*A_extension_26]*FlightRestrictionLevelType)
	for flightrestrictionleveltype := range stage.FlightRestrictionLevelTypes {
		_ = flightrestrictionleveltype
		for _, _a_extension_26 := range flightrestrictionleveltype.Extension {
			stage.FlightRestrictionLevelType_Extension_reverseMap[_a_extension_26] = flightrestrictionleveltype
		}
	}

	// Compute reverse map for named struct FlightRestrictionRouteType
	// insertion point per field
	stage.FlightRestrictionRouteType_RouteElement_reverseMap = make(map[*FlightRoutingElementPropertyType]*FlightRestrictionRouteType)
	for flightrestrictionroutetype := range stage.FlightRestrictionRouteTypes {
		_ = flightrestrictionroutetype
		for _, _flightroutingelementpropertytype := range flightrestrictionroutetype.RouteElement {
			stage.FlightRestrictionRouteType_RouteElement_reverseMap[_flightroutingelementpropertytype] = flightrestrictionroutetype
		}
	}
	stage.FlightRestrictionRouteType_Contact_reverseMap = make(map[*ContactInformationPropertyType]*FlightRestrictionRouteType)
	for flightrestrictionroutetype := range stage.FlightRestrictionRouteTypes {
		_ = flightrestrictionroutetype
		for _, _contactinformationpropertytype := range flightrestrictionroutetype.Contact {
			stage.FlightRestrictionRouteType_Contact_reverseMap[_contactinformationpropertytype] = flightrestrictionroutetype
		}
	}
	stage.FlightRestrictionRouteType_Annotation_reverseMap = make(map[*NotePropertyType]*FlightRestrictionRouteType)
	for flightrestrictionroutetype := range stage.FlightRestrictionRouteTypes {
		_ = flightrestrictionroutetype
		for _, _notepropertytype := range flightrestrictionroutetype.Annotation {
			stage.FlightRestrictionRouteType_Annotation_reverseMap[_notepropertytype] = flightrestrictionroutetype
		}
	}
	stage.FlightRestrictionRouteType_Extension_reverseMap = make(map[*A_extension_27]*FlightRestrictionRouteType)
	for flightrestrictionroutetype := range stage.FlightRestrictionRouteTypes {
		_ = flightrestrictionroutetype
		for _, _a_extension_27 := range flightrestrictionroutetype.Extension {
			stage.FlightRestrictionRouteType_Extension_reverseMap[_a_extension_27] = flightrestrictionroutetype
		}
	}

	// Compute reverse map for named struct FlightRestrictionTimeSliceType
	// insertion point per field
	stage.FlightRestrictionTimeSliceType_RegulatedRoute_reverseMap = make(map[*FlightRestrictionRoutePropertyType]*FlightRestrictionTimeSliceType)
	for flightrestrictiontimeslicetype := range stage.FlightRestrictionTimeSliceTypes {
		_ = flightrestrictiontimeslicetype
		for _, _flightrestrictionroutepropertytype := range flightrestrictiontimeslicetype.RegulatedRoute {
			stage.FlightRestrictionTimeSliceType_RegulatedRoute_reverseMap[_flightrestrictionroutepropertytype] = flightrestrictiontimeslicetype
		}
	}
	stage.FlightRestrictionTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*FlightRestrictionTimeSliceType)
	for flightrestrictiontimeslicetype := range stage.FlightRestrictionTimeSliceTypes {
		_ = flightrestrictiontimeslicetype
		for _, _notepropertytype := range flightrestrictiontimeslicetype.Annotation {
			stage.FlightRestrictionTimeSliceType_Annotation_reverseMap[_notepropertytype] = flightrestrictiontimeslicetype
		}
	}
	stage.FlightRestrictionTimeSliceType_Extension_reverseMap = make(map[*A_extension_25]*FlightRestrictionTimeSliceType)
	for flightrestrictiontimeslicetype := range stage.FlightRestrictionTimeSliceTypes {
		_ = flightrestrictiontimeslicetype
		for _, _a_extension_25 := range flightrestrictiontimeslicetype.Extension {
			stage.FlightRestrictionTimeSliceType_Extension_reverseMap[_a_extension_25] = flightrestrictiontimeslicetype
		}
	}

	// Compute reverse map for named struct FlightRestrictionType
	// insertion point per field
	stage.FlightRestrictionType_TimeSlice_reverseMap = make(map[*FlightRestrictionTimeSlicePropertyType]*FlightRestrictionType)
	for flightrestrictiontype := range stage.FlightRestrictionTypes {
		_ = flightrestrictiontype
		for _, _flightrestrictiontimeslicepropertytype := range flightrestrictiontype.TimeSlice {
			stage.FlightRestrictionType_TimeSlice_reverseMap[_flightrestrictiontimeslicepropertytype] = flightrestrictiontype
		}
	}

	// Compute reverse map for named struct FlightRoutingElementType
	// insertion point per field
	stage.FlightRoutingElementType_FlightLevel_reverseMap = make(map[*FlightRestrictionLevelPropertyType]*FlightRoutingElementType)
	for flightroutingelementtype := range stage.FlightRoutingElementTypes {
		_ = flightroutingelementtype
		for _, _flightrestrictionlevelpropertytype := range flightroutingelementtype.FlightLevel {
			stage.FlightRoutingElementType_FlightLevel_reverseMap[_flightrestrictionlevelpropertytype] = flightroutingelementtype
		}
	}
	stage.FlightRoutingElementType_Annotation_reverseMap = make(map[*NotePropertyType]*FlightRoutingElementType)
	for flightroutingelementtype := range stage.FlightRoutingElementTypes {
		_ = flightroutingelementtype
		for _, _notepropertytype := range flightroutingelementtype.Annotation {
			stage.FlightRoutingElementType_Annotation_reverseMap[_notepropertytype] = flightroutingelementtype
		}
	}
	stage.FlightRoutingElementType_Extension_reverseMap = make(map[*A_extension_28]*FlightRoutingElementType)
	for flightroutingelementtype := range stage.FlightRoutingElementTypes {
		_ = flightroutingelementtype
		for _, _a_extension_28 := range flightroutingelementtype.Extension {
			stage.FlightRoutingElementType_Extension_reverseMap[_a_extension_28] = flightroutingelementtype
		}
	}

	// Compute reverse map for named struct LinguisticNoteType
	// insertion point per field
	stage.LinguisticNoteType_Extension_reverseMap = make(map[*A_extension_11]*LinguisticNoteType)
	for linguisticnotetype := range stage.LinguisticNoteTypes {
		_ = linguisticnotetype
		for _, _a_extension_11 := range linguisticnotetype.Extension {
			stage.LinguisticNoteType_Extension_reverseMap[_a_extension_11] = linguisticnotetype
		}
	}

	// Compute reverse map for named struct MeteorologyType
	// insertion point per field
	stage.MeteorologyType_Annotation_reverseMap = make(map[*NotePropertyType]*MeteorologyType)
	for meteorologytype := range stage.MeteorologyTypes {
		_ = meteorologytype
		for _, _notepropertytype := range meteorologytype.Annotation {
			stage.MeteorologyType_Annotation_reverseMap[_notepropertytype] = meteorologytype
		}
	}
	stage.MeteorologyType_Extension_reverseMap = make(map[*A_extension_37]*MeteorologyType)
	for meteorologytype := range stage.MeteorologyTypes {
		_ = meteorologytype
		for _, _a_extension_37 := range meteorologytype.Extension {
			stage.MeteorologyType_Extension_reverseMap[_a_extension_37] = meteorologytype
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
	stage.NoteType_Extension_reverseMap = make(map[*A_extension_12]*NoteType)
	for notetype := range stage.NoteTypes {
		_ = notetype
		for _, _a_extension_12 := range notetype.Extension {
			stage.NoteType_Extension_reverseMap[_a_extension_12] = notetype
		}
	}

	// Compute reverse map for named struct ObstacleAssessmentAreaType
	// insertion point per field
	stage.ObstacleAssessmentAreaType_AircraftCategory_reverseMap = make(map[*AircraftCharacteristicPropertyType]*ObstacleAssessmentAreaType)
	for obstacleassessmentareatype := range stage.ObstacleAssessmentAreaTypes {
		_ = obstacleassessmentareatype
		for _, _aircraftcharacteristicpropertytype := range obstacleassessmentareatype.AircraftCategory {
			stage.ObstacleAssessmentAreaType_AircraftCategory_reverseMap[_aircraftcharacteristicpropertytype] = obstacleassessmentareatype
		}
	}
	stage.ObstacleAssessmentAreaType_SignificantObstacle_reverseMap = make(map[*ObstructionPropertyType]*ObstacleAssessmentAreaType)
	for obstacleassessmentareatype := range stage.ObstacleAssessmentAreaTypes {
		_ = obstacleassessmentareatype
		for _, _obstructionpropertytype := range obstacleassessmentareatype.SignificantObstacle {
			stage.ObstacleAssessmentAreaType_SignificantObstacle_reverseMap[_obstructionpropertytype] = obstacleassessmentareatype
		}
	}
	stage.ObstacleAssessmentAreaType_Annotation_reverseMap = make(map[*NotePropertyType]*ObstacleAssessmentAreaType)
	for obstacleassessmentareatype := range stage.ObstacleAssessmentAreaTypes {
		_ = obstacleassessmentareatype
		for _, _notepropertytype := range obstacleassessmentareatype.Annotation {
			stage.ObstacleAssessmentAreaType_Annotation_reverseMap[_notepropertytype] = obstacleassessmentareatype
		}
	}
	stage.ObstacleAssessmentAreaType_Extension_reverseMap = make(map[*A_extension_40]*ObstacleAssessmentAreaType)
	for obstacleassessmentareatype := range stage.ObstacleAssessmentAreaTypes {
		_ = obstacleassessmentareatype
		for _, _a_extension_40 := range obstacleassessmentareatype.Extension {
			stage.ObstacleAssessmentAreaType_Extension_reverseMap[_a_extension_40] = obstacleassessmentareatype
		}
	}

	// Compute reverse map for named struct ObstaclePlacementType
	// insertion point per field
	stage.ObstaclePlacementType_Annotation_reverseMap = make(map[*NotePropertyType]*ObstaclePlacementType)
	for obstacleplacementtype := range stage.ObstaclePlacementTypes {
		_ = obstacleplacementtype
		for _, _notepropertytype := range obstacleplacementtype.Annotation {
			stage.ObstaclePlacementType_Annotation_reverseMap[_notepropertytype] = obstacleplacementtype
		}
	}
	stage.ObstaclePlacementType_Extension_reverseMap = make(map[*A_extension_41]*ObstaclePlacementType)
	for obstacleplacementtype := range stage.ObstaclePlacementTypes {
		_ = obstacleplacementtype
		for _, _a_extension_41 := range obstacleplacementtype.Extension {
			stage.ObstaclePlacementType_Extension_reverseMap[_a_extension_41] = obstacleplacementtype
		}
	}

	// Compute reverse map for named struct ObstructionType
	// insertion point per field
	stage.ObstructionType_Adjustment_reverseMap = make(map[*AltitudeAdjustmentPropertyType]*ObstructionType)
	for obstructiontype := range stage.ObstructionTypes {
		_ = obstructiontype
		for _, _altitudeadjustmentpropertytype := range obstructiontype.Adjustment {
			stage.ObstructionType_Adjustment_reverseMap[_altitudeadjustmentpropertytype] = obstructiontype
		}
	}
	stage.ObstructionType_ObstaclePlacement_reverseMap = make(map[*ObstaclePlacementPropertyType]*ObstructionType)
	for obstructiontype := range stage.ObstructionTypes {
		_ = obstructiontype
		for _, _obstacleplacementpropertytype := range obstructiontype.ObstaclePlacement {
			stage.ObstructionType_ObstaclePlacement_reverseMap[_obstacleplacementpropertytype] = obstructiontype
		}
	}
	stage.ObstructionType_Annotation_reverseMap = make(map[*NotePropertyType]*ObstructionType)
	for obstructiontype := range stage.ObstructionTypes {
		_ = obstructiontype
		for _, _notepropertytype := range obstructiontype.Annotation {
			stage.ObstructionType_Annotation_reverseMap[_notepropertytype] = obstructiontype
		}
	}
	stage.ObstructionType_Extension_reverseMap = make(map[*A_extension_42]*ObstructionType)
	for obstructiontype := range stage.ObstructionTypes {
		_ = obstructiontype
		for _, _a_extension_42 := range obstructiontype.Extension {
			stage.ObstructionType_Extension_reverseMap[_a_extension_42] = obstructiontype
		}
	}

	// Compute reverse map for named struct OnlineContactType
	// insertion point per field
	stage.OnlineContactType_TimeInterval_reverseMap = make(map[*TimesheetPropertyType]*OnlineContactType)
	for onlinecontacttype := range stage.OnlineContactTypes {
		_ = onlinecontacttype
		for _, _timesheetpropertytype := range onlinecontacttype.TimeInterval {
			stage.OnlineContactType_TimeInterval_reverseMap[_timesheetpropertytype] = onlinecontacttype
		}
	}
	stage.OnlineContactType_Annotation_reverseMap = make(map[*NotePropertyType]*OnlineContactType)
	for onlinecontacttype := range stage.OnlineContactTypes {
		_ = onlinecontacttype
		for _, _notepropertytype := range onlinecontacttype.Annotation {
			stage.OnlineContactType_Annotation_reverseMap[_notepropertytype] = onlinecontacttype
		}
	}
	stage.OnlineContactType_SpecialDateAuthority_reverseMap = make(map[*OrganisationAuthorityPropertyType]*OnlineContactType)
	for onlinecontacttype := range stage.OnlineContactTypes {
		_ = onlinecontacttype
		for _, _organisationauthoritypropertytype := range onlinecontacttype.SpecialDateAuthority {
			stage.OnlineContactType_SpecialDateAuthority_reverseMap[_organisationauthoritypropertytype] = onlinecontacttype
		}
	}
	stage.OnlineContactType_Extension_reverseMap = make(map[*A_extension_31]*OnlineContactType)
	for onlinecontacttype := range stage.OnlineContactTypes {
		_ = onlinecontacttype
		for _, _a_extension_31 := range onlinecontacttype.Extension {
			stage.OnlineContactType_Extension_reverseMap[_a_extension_31] = onlinecontacttype
		}
	}

	// Compute reverse map for named struct PointReferenceType
	// insertion point per field
	stage.PointReferenceType_FacilityAngle_reverseMap = make(map[*AngleUsePropertyType]*PointReferenceType)
	for pointreferencetype := range stage.PointReferenceTypes {
		_ = pointreferencetype
		for _, _angleusepropertytype := range pointreferencetype.FacilityAngle {
			stage.PointReferenceType_FacilityAngle_reverseMap[_angleusepropertytype] = pointreferencetype
		}
	}
	stage.PointReferenceType_FacilityDistance_reverseMap = make(map[*DistanceIndicationPropertyType]*PointReferenceType)
	for pointreferencetype := range stage.PointReferenceTypes {
		_ = pointreferencetype
		for _, _distanceindicationpropertytype := range pointreferencetype.FacilityDistance {
			stage.PointReferenceType_FacilityDistance_reverseMap[_distanceindicationpropertytype] = pointreferencetype
		}
	}
	stage.PointReferenceType_Annotation_reverseMap = make(map[*NotePropertyType]*PointReferenceType)
	for pointreferencetype := range stage.PointReferenceTypes {
		_ = pointreferencetype
		for _, _notepropertytype := range pointreferencetype.Annotation {
			stage.PointReferenceType_Annotation_reverseMap[_notepropertytype] = pointreferencetype
		}
	}
	stage.PointReferenceType_Extension_reverseMap = make(map[*A_extension_10]*PointReferenceType)
	for pointreferencetype := range stage.PointReferenceTypes {
		_ = pointreferencetype
		for _, _a_extension_10 := range pointreferencetype.Extension {
			stage.PointReferenceType_Extension_reverseMap[_a_extension_10] = pointreferencetype
		}
	}

	// Compute reverse map for named struct PointType
	// insertion point per field
	stage.PointType_Annotation_reverseMap = make(map[*NotePropertyType]*PointType)
	for pointtype := range stage.PointTypes {
		_ = pointtype
		for _, _notepropertytype := range pointtype.Annotation {
			stage.PointType_Annotation_reverseMap[_notepropertytype] = pointtype
		}
	}

	// Compute reverse map for named struct PostalAddressType
	// insertion point per field
	stage.PostalAddressType_TimeInterval_reverseMap = make(map[*TimesheetPropertyType]*PostalAddressType)
	for postaladdresstype := range stage.PostalAddressTypes {
		_ = postaladdresstype
		for _, _timesheetpropertytype := range postaladdresstype.TimeInterval {
			stage.PostalAddressType_TimeInterval_reverseMap[_timesheetpropertytype] = postaladdresstype
		}
	}
	stage.PostalAddressType_Annotation_reverseMap = make(map[*NotePropertyType]*PostalAddressType)
	for postaladdresstype := range stage.PostalAddressTypes {
		_ = postaladdresstype
		for _, _notepropertytype := range postaladdresstype.Annotation {
			stage.PostalAddressType_Annotation_reverseMap[_notepropertytype] = postaladdresstype
		}
	}
	stage.PostalAddressType_SpecialDateAuthority_reverseMap = make(map[*OrganisationAuthorityPropertyType]*PostalAddressType)
	for postaladdresstype := range stage.PostalAddressTypes {
		_ = postaladdresstype
		for _, _organisationauthoritypropertytype := range postaladdresstype.SpecialDateAuthority {
			stage.PostalAddressType_SpecialDateAuthority_reverseMap[_organisationauthoritypropertytype] = postaladdresstype
		}
	}
	stage.PostalAddressType_Extension_reverseMap = make(map[*A_extension_32]*PostalAddressType)
	for postaladdresstype := range stage.PostalAddressTypes {
		_ = postaladdresstype
		for _, _a_extension_32 := range postaladdresstype.Extension {
			stage.PostalAddressType_Extension_reverseMap[_a_extension_32] = postaladdresstype
		}
	}

	// Compute reverse map for named struct RouteAvailabilityType
	// insertion point per field
	stage.RouteAvailabilityType_TimeInterval_reverseMap = make(map[*TimesheetPropertyType]*RouteAvailabilityType)
	for routeavailabilitytype := range stage.RouteAvailabilityTypes {
		_ = routeavailabilitytype
		for _, _timesheetpropertytype := range routeavailabilitytype.TimeInterval {
			stage.RouteAvailabilityType_TimeInterval_reverseMap[_timesheetpropertytype] = routeavailabilitytype
		}
	}
	stage.RouteAvailabilityType_Annotation_reverseMap = make(map[*NotePropertyType]*RouteAvailabilityType)
	for routeavailabilitytype := range stage.RouteAvailabilityTypes {
		_ = routeavailabilitytype
		for _, _notepropertytype := range routeavailabilitytype.Annotation {
			stage.RouteAvailabilityType_Annotation_reverseMap[_notepropertytype] = routeavailabilitytype
		}
	}
	stage.RouteAvailabilityType_SpecialDateAuthority_reverseMap = make(map[*OrganisationAuthorityPropertyType]*RouteAvailabilityType)
	for routeavailabilitytype := range stage.RouteAvailabilityTypes {
		_ = routeavailabilitytype
		for _, _organisationauthoritypropertytype := range routeavailabilitytype.SpecialDateAuthority {
			stage.RouteAvailabilityType_SpecialDateAuthority_reverseMap[_organisationauthoritypropertytype] = routeavailabilitytype
		}
	}
	stage.RouteAvailabilityType_Levels_reverseMap = make(map[*AirspaceLayerPropertyType]*RouteAvailabilityType)
	for routeavailabilitytype := range stage.RouteAvailabilityTypes {
		_ = routeavailabilitytype
		for _, _airspacelayerpropertytype := range routeavailabilitytype.Levels {
			stage.RouteAvailabilityType_Levels_reverseMap[_airspacelayerpropertytype] = routeavailabilitytype
		}
	}
	stage.RouteAvailabilityType_Extension_reverseMap = make(map[*A_extension_15]*RouteAvailabilityType)
	for routeavailabilitytype := range stage.RouteAvailabilityTypes {
		_ = routeavailabilitytype
		for _, _a_extension_15 := range routeavailabilitytype.Extension {
			stage.RouteAvailabilityType_Extension_reverseMap[_a_extension_15] = routeavailabilitytype
		}
	}

	// Compute reverse map for named struct RouteDMETimeSliceType
	// insertion point per field
	stage.RouteDMETimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*RouteDMETimeSliceType)
	for routedmetimeslicetype := range stage.RouteDMETimeSliceTypes {
		_ = routedmetimeslicetype
		for _, _notepropertytype := range routedmetimeslicetype.Annotation {
			stage.RouteDMETimeSliceType_Annotation_reverseMap[_notepropertytype] = routedmetimeslicetype
		}
	}
	stage.RouteDMETimeSliceType_Extension_reverseMap = make(map[*A_extension_16]*RouteDMETimeSliceType)
	for routedmetimeslicetype := range stage.RouteDMETimeSliceTypes {
		_ = routedmetimeslicetype
		for _, _a_extension_16 := range routedmetimeslicetype.Extension {
			stage.RouteDMETimeSliceType_Extension_reverseMap[_a_extension_16] = routedmetimeslicetype
		}
	}

	// Compute reverse map for named struct RouteDMEType
	// insertion point per field
	stage.RouteDMEType_TimeSlice_reverseMap = make(map[*RouteDMETimeSlicePropertyType]*RouteDMEType)
	for routedmetype := range stage.RouteDMETypes {
		_ = routedmetype
		for _, _routedmetimeslicepropertytype := range routedmetype.TimeSlice {
			stage.RouteDMEType_TimeSlice_reverseMap[_routedmetimeslicepropertytype] = routedmetype
		}
	}

	// Compute reverse map for named struct RoutePortionType
	// insertion point per field
	stage.RoutePortionType_Annotation_reverseMap = make(map[*NotePropertyType]*RoutePortionType)
	for routeportiontype := range stage.RoutePortionTypes {
		_ = routeportiontype
		for _, _notepropertytype := range routeportiontype.Annotation {
			stage.RoutePortionType_Annotation_reverseMap[_notepropertytype] = routeportiontype
		}
	}
	stage.RoutePortionType_Extension_reverseMap = make(map[*A_extension_17]*RoutePortionType)
	for routeportiontype := range stage.RoutePortionTypes {
		_ = routeportiontype
		for _, _a_extension_17 := range routeportiontype.Extension {
			stage.RoutePortionType_Extension_reverseMap[_a_extension_17] = routeportiontype
		}
	}

	// Compute reverse map for named struct RouteSegmentTimeSliceType
	// insertion point per field
	stage.RouteSegmentTimeSliceType_Availability_reverseMap = make(map[*RouteAvailabilityPropertyType]*RouteSegmentTimeSliceType)
	for routesegmenttimeslicetype := range stage.RouteSegmentTimeSliceTypes {
		_ = routesegmenttimeslicetype
		for _, _routeavailabilitypropertytype := range routesegmenttimeslicetype.Availability {
			stage.RouteSegmentTimeSliceType_Availability_reverseMap[_routeavailabilitypropertytype] = routesegmenttimeslicetype
		}
	}
	stage.RouteSegmentTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*RouteSegmentTimeSliceType)
	for routesegmenttimeslicetype := range stage.RouteSegmentTimeSliceTypes {
		_ = routesegmenttimeslicetype
		for _, _notepropertytype := range routesegmenttimeslicetype.Annotation {
			stage.RouteSegmentTimeSliceType_Annotation_reverseMap[_notepropertytype] = routesegmenttimeslicetype
		}
	}
	stage.RouteSegmentTimeSliceType_Extension_reverseMap = make(map[*A_extension_18]*RouteSegmentTimeSliceType)
	for routesegmenttimeslicetype := range stage.RouteSegmentTimeSliceTypes {
		_ = routesegmenttimeslicetype
		for _, _a_extension_18 := range routesegmenttimeslicetype.Extension {
			stage.RouteSegmentTimeSliceType_Extension_reverseMap[_a_extension_18] = routesegmenttimeslicetype
		}
	}

	// Compute reverse map for named struct RouteSegmentType
	// insertion point per field
	stage.RouteSegmentType_TimeSlice_reverseMap = make(map[*RouteSegmentTimeSlicePropertyType]*RouteSegmentType)
	for routesegmenttype := range stage.RouteSegmentTypes {
		_ = routesegmenttype
		for _, _routesegmenttimeslicepropertytype := range routesegmenttype.TimeSlice {
			stage.RouteSegmentType_TimeSlice_reverseMap[_routesegmenttimeslicepropertytype] = routesegmenttype
		}
	}

	// Compute reverse map for named struct RouteTimeSliceType
	// insertion point per field
	stage.RouteTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*RouteTimeSliceType)
	for routetimeslicetype := range stage.RouteTimeSliceTypes {
		_ = routetimeslicetype
		for _, _notepropertytype := range routetimeslicetype.Annotation {
			stage.RouteTimeSliceType_Annotation_reverseMap[_notepropertytype] = routetimeslicetype
		}
	}
	stage.RouteTimeSliceType_Extension_reverseMap = make(map[*A_extension_14]*RouteTimeSliceType)
	for routetimeslicetype := range stage.RouteTimeSliceTypes {
		_ = routetimeslicetype
		for _, _a_extension_14 := range routetimeslicetype.Extension {
			stage.RouteTimeSliceType_Extension_reverseMap[_a_extension_14] = routetimeslicetype
		}
	}

	// Compute reverse map for named struct RouteType
	// insertion point per field
	stage.RouteType_TimeSlice_reverseMap = make(map[*RouteTimeSlicePropertyType]*RouteType)
	for routetype := range stage.RouteTypes {
		_ = routetype
		for _, _routetimeslicepropertytype := range routetype.TimeSlice {
			stage.RouteType_TimeSlice_reverseMap[_routetimeslicepropertytype] = routetype
		}
	}

	// Compute reverse map for named struct RulesProceduresTimeSliceType
	// insertion point per field
	stage.RulesProceduresTimeSliceType_AffectedLocation_reverseMap = make(map[*AirportHeliportPropertyType]*RulesProceduresTimeSliceType)
	for rulesprocedurestimeslicetype := range stage.RulesProceduresTimeSliceTypes {
		_ = rulesprocedurestimeslicetype
		for _, _airportheliportpropertytype := range rulesprocedurestimeslicetype.AffectedLocation {
			stage.RulesProceduresTimeSliceType_AffectedLocation_reverseMap[_airportheliportpropertytype] = rulesprocedurestimeslicetype
		}
	}
	stage.RulesProceduresTimeSliceType_AffectedArea_reverseMap = make(map[*AirspacePropertyType]*RulesProceduresTimeSliceType)
	for rulesprocedurestimeslicetype := range stage.RulesProceduresTimeSliceTypes {
		_ = rulesprocedurestimeslicetype
		for _, _airspacepropertytype := range rulesprocedurestimeslicetype.AffectedArea {
			stage.RulesProceduresTimeSliceType_AffectedArea_reverseMap[_airspacepropertytype] = rulesprocedurestimeslicetype
		}
	}
	stage.RulesProceduresTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*RulesProceduresTimeSliceType)
	for rulesprocedurestimeslicetype := range stage.RulesProceduresTimeSliceTypes {
		_ = rulesprocedurestimeslicetype
		for _, _notepropertytype := range rulesprocedurestimeslicetype.Annotation {
			stage.RulesProceduresTimeSliceType_Annotation_reverseMap[_notepropertytype] = rulesprocedurestimeslicetype
		}
	}
	stage.RulesProceduresTimeSliceType_Extension_reverseMap = make(map[*A_extension_29]*RulesProceduresTimeSliceType)
	for rulesprocedurestimeslicetype := range stage.RulesProceduresTimeSliceTypes {
		_ = rulesprocedurestimeslicetype
		for _, _a_extension_29 := range rulesprocedurestimeslicetype.Extension {
			stage.RulesProceduresTimeSliceType_Extension_reverseMap[_a_extension_29] = rulesprocedurestimeslicetype
		}
	}

	// Compute reverse map for named struct RulesProceduresType
	// insertion point per field
	stage.RulesProceduresType_TimeSlice_reverseMap = make(map[*RulesProceduresTimeSlicePropertyType]*RulesProceduresType)
	for rulesprocedurestype := range stage.RulesProceduresTypes {
		_ = rulesprocedurestype
		for _, _rulesprocedurestimeslicepropertytype := range rulesprocedurestype.TimeSlice {
			stage.RulesProceduresType_TimeSlice_reverseMap[_rulesprocedurestimeslicepropertytype] = rulesprocedurestype
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

	// Compute reverse map for named struct TelephoneContactType
	// insertion point per field
	stage.TelephoneContactType_TimeInterval_reverseMap = make(map[*TimesheetPropertyType]*TelephoneContactType)
	for telephonecontacttype := range stage.TelephoneContactTypes {
		_ = telephonecontacttype
		for _, _timesheetpropertytype := range telephonecontacttype.TimeInterval {
			stage.TelephoneContactType_TimeInterval_reverseMap[_timesheetpropertytype] = telephonecontacttype
		}
	}
	stage.TelephoneContactType_Annotation_reverseMap = make(map[*NotePropertyType]*TelephoneContactType)
	for telephonecontacttype := range stage.TelephoneContactTypes {
		_ = telephonecontacttype
		for _, _notepropertytype := range telephonecontacttype.Annotation {
			stage.TelephoneContactType_Annotation_reverseMap[_notepropertytype] = telephonecontacttype
		}
	}
	stage.TelephoneContactType_SpecialDateAuthority_reverseMap = make(map[*OrganisationAuthorityPropertyType]*TelephoneContactType)
	for telephonecontacttype := range stage.TelephoneContactTypes {
		_ = telephonecontacttype
		for _, _organisationauthoritypropertytype := range telephonecontacttype.SpecialDateAuthority {
			stage.TelephoneContactType_SpecialDateAuthority_reverseMap[_organisationauthoritypropertytype] = telephonecontacttype
		}
	}
	stage.TelephoneContactType_Extension_reverseMap = make(map[*A_extension_33]*TelephoneContactType)
	for telephonecontacttype := range stage.TelephoneContactTypes {
		_ = telephonecontacttype
		for _, _a_extension_33 := range telephonecontacttype.Extension {
			stage.TelephoneContactType_Extension_reverseMap[_a_extension_33] = telephonecontacttype
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
	stage.TimesheetType_Extension_reverseMap = make(map[*A_extension_38]*TimesheetType)
	for timesheettype := range stage.TimesheetTypes {
		_ = timesheettype
		for _, _a_extension_38 := range timesheettype.Extension {
			stage.TimesheetType_Extension_reverseMap[_a_extension_38] = timesheettype
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

	res = __gong__appendInstances(res, stage.A_extension_15s)

	res = __gong__appendInstances(res, stage.A_extension_16s)

	res = __gong__appendInstances(res, stage.A_extension_17s)

	res = __gong__appendInstances(res, stage.A_extension_18s)

	res = __gong__appendInstances(res, stage.A_extension_19s)

	res = __gong__appendInstances(res, stage.A_extension_2s)

	res = __gong__appendInstances(res, stage.A_extension_20s)

	res = __gong__appendInstances(res, stage.A_extension_21s)

	res = __gong__appendInstances(res, stage.A_extension_22s)

	res = __gong__appendInstances(res, stage.A_extension_23s)

	res = __gong__appendInstances(res, stage.A_extension_24s)

	res = __gong__appendInstances(res, stage.A_extension_25s)

	res = __gong__appendInstances(res, stage.A_extension_26s)

	res = __gong__appendInstances(res, stage.A_extension_27s)

	res = __gong__appendInstances(res, stage.A_extension_28s)

	res = __gong__appendInstances(res, stage.A_extension_29s)

	res = __gong__appendInstances(res, stage.A_extension_3s)

	res = __gong__appendInstances(res, stage.A_extension_30s)

	res = __gong__appendInstances(res, stage.A_extension_31s)

	res = __gong__appendInstances(res, stage.A_extension_32s)

	res = __gong__appendInstances(res, stage.A_extension_33s)

	res = __gong__appendInstances(res, stage.A_extension_34s)

	res = __gong__appendInstances(res, stage.A_extension_35s)

	res = __gong__appendInstances(res, stage.A_extension_36s)

	res = __gong__appendInstances(res, stage.A_extension_37s)

	res = __gong__appendInstances(res, stage.A_extension_38s)

	res = __gong__appendInstances(res, stage.A_extension_39s)

	res = __gong__appendInstances(res, stage.A_extension_4s)

	res = __gong__appendInstances(res, stage.A_extension_40s)

	res = __gong__appendInstances(res, stage.A_extension_41s)

	res = __gong__appendInstances(res, stage.A_extension_42s)

	res = __gong__appendInstances(res, stage.A_extension_5s)

	res = __gong__appendInstances(res, stage.A_extension_6s)

	res = __gong__appendInstances(res, stage.A_extension_7s)

	res = __gong__appendInstances(res, stage.A_extension_8s)

	res = __gong__appendInstances(res, stage.A_extension_9s)

	res = __gong__appendInstances(res, stage.AbstractDirectFlightTypes)

	res = __gong__appendInstances(res, stage.AbstractPropertiesWithScheduleTypes)

	res = __gong__appendInstances(res, stage.AbstractSegmentPointTypes)

	res = __gong__appendInstances(res, stage.AerialRefuellingAnchorPropertyTypes)

	res = __gong__appendInstances(res, stage.AerialRefuellingAnchorTypes)

	res = __gong__appendInstances(res, stage.AerialRefuellingPointPropertyTypes)

	res = __gong__appendInstances(res, stage.AerialRefuellingPointTypes)

	res = __gong__appendInstances(res, stage.AerialRefuellingPropertyTypes)

	res = __gong__appendInstances(res, stage.AerialRefuellingTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.AerialRefuellingTimeSliceTypes)

	res = __gong__appendInstances(res, stage.AerialRefuellingTrackPropertyTypes)

	res = __gong__appendInstances(res, stage.AerialRefuellingTrackTypes)

	res = __gong__appendInstances(res, stage.AerialRefuellingTypes)

	res = __gong__appendInstances(res, stage.AircraftCharacteristicPropertyTypes)

	res = __gong__appendInstances(res, stage.AircraftCharacteristicTypes)

	res = __gong__appendInstances(res, stage.AirportHeliportPropertyTypes)

	res = __gong__appendInstances(res, stage.AirspaceBorderCrossingPropertyTypes)

	res = __gong__appendInstances(res, stage.AirspaceBorderCrossingTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.AirspaceBorderCrossingTimeSliceTypes)

	res = __gong__appendInstances(res, stage.AirspaceBorderCrossingTypes)

	res = __gong__appendInstances(res, stage.AirspaceLayerPropertyTypes)

	res = __gong__appendInstances(res, stage.AirspaceLayerTypes)

	res = __gong__appendInstances(res, stage.AirspacePropertyTypes)

	res = __gong__appendInstances(res, stage.AltitudeAdjustmentPropertyTypes)

	res = __gong__appendInstances(res, stage.AltitudeAdjustmentTypes)

	res = __gong__appendInstances(res, stage.AngleIndicationPropertyTypes)

	res = __gong__appendInstances(res, stage.AngleUsePropertyTypes)

	res = __gong__appendInstances(res, stage.AngleUseTypes)

	res = __gong__appendInstances(res, stage.AuthorityForAerialRefuellingPropertyTypes)

	res = __gong__appendInstances(res, stage.AuthorityForAerialRefuellingTypes)

	res = __gong__appendInstances(res, stage.ChangeOverPointTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.ChangeOverPointTimeSliceTypes)

	res = __gong__appendInstances(res, stage.ChangeOverPointTypes)

	res = __gong__appendInstances(res, stage.ContactInformationPropertyTypes)

	res = __gong__appendInstances(res, stage.ContactInformationTypes)

	res = __gong__appendInstances(res, stage.CurvePropertyTypes)

	res = __gong__appendInstances(res, stage.CurveTypes)

	res = __gong__appendInstances(res, stage.DMEPropertyTypes)

	res = __gong__appendInstances(res, stage.DesignatedPointPropertyTypes)

	res = __gong__appendInstances(res, stage.DirectFlightClassTypes)

	res = __gong__appendInstances(res, stage.DirectFlightPropertyTypes)

	res = __gong__appendInstances(res, stage.DirectFlightSegmentPropertyTypes)

	res = __gong__appendInstances(res, stage.DirectFlightSegmentTypes)

	res = __gong__appendInstances(res, stage.DistanceIndicationPropertyTypes)

	res = __gong__appendInstances(res, stage.ElevatedCurveTypes)

	res = __gong__appendInstances(res, stage.ElevatedPointTypes)

	res = __gong__appendInstances(res, stage.ElevatedSurfaceTypes)

	res = __gong__appendInstances(res, stage.EnRouteSegmentPointPropertyTypes)

	res = __gong__appendInstances(res, stage.EnRouteSegmentPointTypes)

	res = __gong__appendInstances(res, stage.FlightCharacteristicPropertyTypes)

	res = __gong__appendInstances(res, stage.FlightCharacteristicTypes)

	res = __gong__appendInstances(res, stage.FlightConditionCircumstancePropertyTypes)

	res = __gong__appendInstances(res, stage.FlightConditionCircumstanceTypes)

	res = __gong__appendInstances(res, stage.FlightConditionCombinationPropertyTypes)

	res = __gong__appendInstances(res, stage.FlightConditionCombinationTypes)

	res = __gong__appendInstances(res, stage.FlightConditionElementPropertyTypes)

	res = __gong__appendInstances(res, stage.FlightConditionElementTypes)

	res = __gong__appendInstances(res, stage.FlightRestrictionLevelPropertyTypes)

	res = __gong__appendInstances(res, stage.FlightRestrictionLevelTypes)

	res = __gong__appendInstances(res, stage.FlightRestrictionRoutePropertyTypes)

	res = __gong__appendInstances(res, stage.FlightRestrictionRouteTypes)

	res = __gong__appendInstances(res, stage.FlightRestrictionTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.FlightRestrictionTimeSliceTypes)

	res = __gong__appendInstances(res, stage.FlightRestrictionTypes)

	res = __gong__appendInstances(res, stage.FlightRoutingElementPropertyTypes)

	res = __gong__appendInstances(res, stage.FlightRoutingElementTypes)

	res = __gong__appendInstances(res, stage.LinguisticNotePropertyTypes)

	res = __gong__appendInstances(res, stage.LinguisticNoteTypes)

	res = __gong__appendInstances(res, stage.MeteorologyPropertyTypes)

	res = __gong__appendInstances(res, stage.MeteorologyTypes)

	res = __gong__appendInstances(res, stage.NavaidPropertyTypes)

	res = __gong__appendInstances(res, stage.NotePropertyTypes)

	res = __gong__appendInstances(res, stage.NoteTypes)

	res = __gong__appendInstances(res, stage.ObstacleAssessmentAreaPropertyTypes)

	res = __gong__appendInstances(res, stage.ObstacleAssessmentAreaTypes)

	res = __gong__appendInstances(res, stage.ObstaclePlacementPropertyTypes)

	res = __gong__appendInstances(res, stage.ObstaclePlacementTypes)

	res = __gong__appendInstances(res, stage.ObstructionPropertyTypes)

	res = __gong__appendInstances(res, stage.ObstructionTypes)

	res = __gong__appendInstances(res, stage.OnlineContactPropertyTypes)

	res = __gong__appendInstances(res, stage.OnlineContactTypes)

	res = __gong__appendInstances(res, stage.OrganisationAuthorityPropertyTypes)

	res = __gong__appendInstances(res, stage.PointPropertyTypes)

	res = __gong__appendInstances(res, stage.PointReferencePropertyTypes)

	res = __gong__appendInstances(res, stage.PointReferenceTypes)

	res = __gong__appendInstances(res, stage.PointTypes)

	res = __gong__appendInstances(res, stage.PostalAddressPropertyTypes)

	res = __gong__appendInstances(res, stage.PostalAddressTypes)

	res = __gong__appendInstances(res, stage.RadioFrequencyAreaPropertyTypes)

	res = __gong__appendInstances(res, stage.RouteAvailabilityPropertyTypes)

	res = __gong__appendInstances(res, stage.RouteAvailabilityTypes)

	res = __gong__appendInstances(res, stage.RouteDMETimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.RouteDMETimeSliceTypes)

	res = __gong__appendInstances(res, stage.RouteDMETypes)

	res = __gong__appendInstances(res, stage.RoutePortionPropertyTypes)

	res = __gong__appendInstances(res, stage.RoutePortionTypes)

	res = __gong__appendInstances(res, stage.RoutePropertyTypes)

	res = __gong__appendInstances(res, stage.RouteSegmentTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.RouteSegmentTimeSliceTypes)

	res = __gong__appendInstances(res, stage.RouteSegmentTypes)

	res = __gong__appendInstances(res, stage.RouteTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.RouteTimeSliceTypes)

	res = __gong__appendInstances(res, stage.RouteTypes)

	res = __gong__appendInstances(res, stage.RulesProceduresTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.RulesProceduresTimeSliceTypes)

	res = __gong__appendInstances(res, stage.RulesProceduresTypes)

	res = __gong__appendInstances(res, stage.RunwayCentrelinePointPropertyTypes)

	res = __gong__appendInstances(res, stage.StandardInstrumentArrivalPropertyTypes)

	res = __gong__appendInstances(res, stage.StandardInstrumentDeparturePropertyTypes)

	res = __gong__appendInstances(res, stage.StandardLevelColumnPropertyTypes)

	res = __gong__appendInstances(res, stage.SurfacePropertyTypes)

	res = __gong__appendInstances(res, stage.SurfaceTypes)

	res = __gong__appendInstances(res, stage.TelephoneContactPropertyTypes)

	res = __gong__appendInstances(res, stage.TelephoneContactTypes)

	res = __gong__appendInstances(res, stage.TimesheetPropertyTypes)

	res = __gong__appendInstances(res, stage.TimesheetTypes)

	res = __gong__appendInstances(res, stage.TouchDownLiftOffPropertyTypes)

	res = __gong__appendInstances(res, stage.VerticalStructurePropertyTypes)

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

func (a_extension_15 *A_extension_15) GongCopy() GongstructIF {
	newInstance := new(A_extension_15)
	a_extension_15.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_16 *A_extension_16) GongCopy() GongstructIF {
	newInstance := new(A_extension_16)
	a_extension_16.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_17 *A_extension_17) GongCopy() GongstructIF {
	newInstance := new(A_extension_17)
	a_extension_17.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_18 *A_extension_18) GongCopy() GongstructIF {
	newInstance := new(A_extension_18)
	a_extension_18.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_19 *A_extension_19) GongCopy() GongstructIF {
	newInstance := new(A_extension_19)
	a_extension_19.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_2 *A_extension_2) GongCopy() GongstructIF {
	newInstance := new(A_extension_2)
	a_extension_2.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_20 *A_extension_20) GongCopy() GongstructIF {
	newInstance := new(A_extension_20)
	a_extension_20.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_21 *A_extension_21) GongCopy() GongstructIF {
	newInstance := new(A_extension_21)
	a_extension_21.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_22 *A_extension_22) GongCopy() GongstructIF {
	newInstance := new(A_extension_22)
	a_extension_22.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_23 *A_extension_23) GongCopy() GongstructIF {
	newInstance := new(A_extension_23)
	a_extension_23.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_24 *A_extension_24) GongCopy() GongstructIF {
	newInstance := new(A_extension_24)
	a_extension_24.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_25 *A_extension_25) GongCopy() GongstructIF {
	newInstance := new(A_extension_25)
	a_extension_25.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_26 *A_extension_26) GongCopy() GongstructIF {
	newInstance := new(A_extension_26)
	a_extension_26.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_27 *A_extension_27) GongCopy() GongstructIF {
	newInstance := new(A_extension_27)
	a_extension_27.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_28 *A_extension_28) GongCopy() GongstructIF {
	newInstance := new(A_extension_28)
	a_extension_28.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_29 *A_extension_29) GongCopy() GongstructIF {
	newInstance := new(A_extension_29)
	a_extension_29.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_3 *A_extension_3) GongCopy() GongstructIF {
	newInstance := new(A_extension_3)
	a_extension_3.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_30 *A_extension_30) GongCopy() GongstructIF {
	newInstance := new(A_extension_30)
	a_extension_30.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_31 *A_extension_31) GongCopy() GongstructIF {
	newInstance := new(A_extension_31)
	a_extension_31.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_32 *A_extension_32) GongCopy() GongstructIF {
	newInstance := new(A_extension_32)
	a_extension_32.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_33 *A_extension_33) GongCopy() GongstructIF {
	newInstance := new(A_extension_33)
	a_extension_33.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_34 *A_extension_34) GongCopy() GongstructIF {
	newInstance := new(A_extension_34)
	a_extension_34.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_35 *A_extension_35) GongCopy() GongstructIF {
	newInstance := new(A_extension_35)
	a_extension_35.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_36 *A_extension_36) GongCopy() GongstructIF {
	newInstance := new(A_extension_36)
	a_extension_36.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_37 *A_extension_37) GongCopy() GongstructIF {
	newInstance := new(A_extension_37)
	a_extension_37.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_38 *A_extension_38) GongCopy() GongstructIF {
	newInstance := new(A_extension_38)
	a_extension_38.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_39 *A_extension_39) GongCopy() GongstructIF {
	newInstance := new(A_extension_39)
	a_extension_39.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_4 *A_extension_4) GongCopy() GongstructIF {
	newInstance := new(A_extension_4)
	a_extension_4.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_40 *A_extension_40) GongCopy() GongstructIF {
	newInstance := new(A_extension_40)
	a_extension_40.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_41 *A_extension_41) GongCopy() GongstructIF {
	newInstance := new(A_extension_41)
	a_extension_41.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_extension_42 *A_extension_42) GongCopy() GongstructIF {
	newInstance := new(A_extension_42)
	a_extension_42.GongCopyBasicFields(newInstance)
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

func (abstractdirectflighttype *AbstractDirectFlightType) GongCopy() GongstructIF {
	newInstance := new(AbstractDirectFlightType)
	abstractdirectflighttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongCopy() GongstructIF {
	newInstance := new(AbstractPropertiesWithScheduleType)
	abstractpropertieswithscheduletype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstractsegmentpointtype *AbstractSegmentPointType) GongCopy() GongstructIF {
	newInstance := new(AbstractSegmentPointType)
	abstractsegmentpointtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (aerialrefuellinganchorpropertytype *AerialRefuellingAnchorPropertyType) GongCopy() GongstructIF {
	newInstance := new(AerialRefuellingAnchorPropertyType)
	aerialrefuellinganchorpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (aerialrefuellinganchortype *AerialRefuellingAnchorType) GongCopy() GongstructIF {
	newInstance := new(AerialRefuellingAnchorType)
	aerialrefuellinganchortype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (aerialrefuellingpointpropertytype *AerialRefuellingPointPropertyType) GongCopy() GongstructIF {
	newInstance := new(AerialRefuellingPointPropertyType)
	aerialrefuellingpointpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (aerialrefuellingpointtype *AerialRefuellingPointType) GongCopy() GongstructIF {
	newInstance := new(AerialRefuellingPointType)
	aerialrefuellingpointtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (aerialrefuellingpropertytype *AerialRefuellingPropertyType) GongCopy() GongstructIF {
	newInstance := new(AerialRefuellingPropertyType)
	aerialrefuellingpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (aerialrefuellingtimeslicepropertytype *AerialRefuellingTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(AerialRefuellingTimeSlicePropertyType)
	aerialrefuellingtimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (aerialrefuellingtimeslicetype *AerialRefuellingTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(AerialRefuellingTimeSliceType)
	aerialrefuellingtimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (aerialrefuellingtrackpropertytype *AerialRefuellingTrackPropertyType) GongCopy() GongstructIF {
	newInstance := new(AerialRefuellingTrackPropertyType)
	aerialrefuellingtrackpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (aerialrefuellingtracktype *AerialRefuellingTrackType) GongCopy() GongstructIF {
	newInstance := new(AerialRefuellingTrackType)
	aerialrefuellingtracktype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (aerialrefuellingtype *AerialRefuellingType) GongCopy() GongstructIF {
	newInstance := new(AerialRefuellingType)
	aerialrefuellingtype.GongCopyBasicFields(newInstance)
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

func (airportheliportpropertytype *AirportHeliportPropertyType) GongCopy() GongstructIF {
	newInstance := new(AirportHeliportPropertyType)
	airportheliportpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airspacebordercrossingpropertytype *AirspaceBorderCrossingPropertyType) GongCopy() GongstructIF {
	newInstance := new(AirspaceBorderCrossingPropertyType)
	airspacebordercrossingpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airspacebordercrossingtimeslicepropertytype *AirspaceBorderCrossingTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(AirspaceBorderCrossingTimeSlicePropertyType)
	airspacebordercrossingtimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airspacebordercrossingtimeslicetype *AirspaceBorderCrossingTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(AirspaceBorderCrossingTimeSliceType)
	airspacebordercrossingtimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airspacebordercrossingtype *AirspaceBorderCrossingType) GongCopy() GongstructIF {
	newInstance := new(AirspaceBorderCrossingType)
	airspacebordercrossingtype.GongCopyBasicFields(newInstance)
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

func (altitudeadjustmentpropertytype *AltitudeAdjustmentPropertyType) GongCopy() GongstructIF {
	newInstance := new(AltitudeAdjustmentPropertyType)
	altitudeadjustmentpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (altitudeadjustmenttype *AltitudeAdjustmentType) GongCopy() GongstructIF {
	newInstance := new(AltitudeAdjustmentType)
	altitudeadjustmenttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (angleindicationpropertytype *AngleIndicationPropertyType) GongCopy() GongstructIF {
	newInstance := new(AngleIndicationPropertyType)
	angleindicationpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (angleusepropertytype *AngleUsePropertyType) GongCopy() GongstructIF {
	newInstance := new(AngleUsePropertyType)
	angleusepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (angleusetype *AngleUseType) GongCopy() GongstructIF {
	newInstance := new(AngleUseType)
	angleusetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (authorityforaerialrefuellingpropertytype *AuthorityForAerialRefuellingPropertyType) GongCopy() GongstructIF {
	newInstance := new(AuthorityForAerialRefuellingPropertyType)
	authorityforaerialrefuellingpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (authorityforaerialrefuellingtype *AuthorityForAerialRefuellingType) GongCopy() GongstructIF {
	newInstance := new(AuthorityForAerialRefuellingType)
	authorityforaerialrefuellingtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (changeoverpointtimeslicepropertytype *ChangeOverPointTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(ChangeOverPointTimeSlicePropertyType)
	changeoverpointtimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (changeoverpointtimeslicetype *ChangeOverPointTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(ChangeOverPointTimeSliceType)
	changeoverpointtimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (changeoverpointtype *ChangeOverPointType) GongCopy() GongstructIF {
	newInstance := new(ChangeOverPointType)
	changeoverpointtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (contactinformationpropertytype *ContactInformationPropertyType) GongCopy() GongstructIF {
	newInstance := new(ContactInformationPropertyType)
	contactinformationpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (contactinformationtype *ContactInformationType) GongCopy() GongstructIF {
	newInstance := new(ContactInformationType)
	contactinformationtype.GongCopyBasicFields(newInstance)
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

func (dmepropertytype *DMEPropertyType) GongCopy() GongstructIF {
	newInstance := new(DMEPropertyType)
	dmepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (designatedpointpropertytype *DesignatedPointPropertyType) GongCopy() GongstructIF {
	newInstance := new(DesignatedPointPropertyType)
	designatedpointpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (directflightclasstype *DirectFlightClassType) GongCopy() GongstructIF {
	newInstance := new(DirectFlightClassType)
	directflightclasstype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (directflightpropertytype *DirectFlightPropertyType) GongCopy() GongstructIF {
	newInstance := new(DirectFlightPropertyType)
	directflightpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (directflightsegmentpropertytype *DirectFlightSegmentPropertyType) GongCopy() GongstructIF {
	newInstance := new(DirectFlightSegmentPropertyType)
	directflightsegmentpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (directflightsegmenttype *DirectFlightSegmentType) GongCopy() GongstructIF {
	newInstance := new(DirectFlightSegmentType)
	directflightsegmenttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (distanceindicationpropertytype *DistanceIndicationPropertyType) GongCopy() GongstructIF {
	newInstance := new(DistanceIndicationPropertyType)
	distanceindicationpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (elevatedcurvetype *ElevatedCurveType) GongCopy() GongstructIF {
	newInstance := new(ElevatedCurveType)
	elevatedcurvetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (elevatedpointtype *ElevatedPointType) GongCopy() GongstructIF {
	newInstance := new(ElevatedPointType)
	elevatedpointtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongCopy() GongstructIF {
	newInstance := new(ElevatedSurfaceType)
	elevatedsurfacetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (enroutesegmentpointpropertytype *EnRouteSegmentPointPropertyType) GongCopy() GongstructIF {
	newInstance := new(EnRouteSegmentPointPropertyType)
	enroutesegmentpointpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (enroutesegmentpointtype *EnRouteSegmentPointType) GongCopy() GongstructIF {
	newInstance := new(EnRouteSegmentPointType)
	enroutesegmentpointtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (flightcharacteristicpropertytype *FlightCharacteristicPropertyType) GongCopy() GongstructIF {
	newInstance := new(FlightCharacteristicPropertyType)
	flightcharacteristicpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (flightcharacteristictype *FlightCharacteristicType) GongCopy() GongstructIF {
	newInstance := new(FlightCharacteristicType)
	flightcharacteristictype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (flightconditioncircumstancepropertytype *FlightConditionCircumstancePropertyType) GongCopy() GongstructIF {
	newInstance := new(FlightConditionCircumstancePropertyType)
	flightconditioncircumstancepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (flightconditioncircumstancetype *FlightConditionCircumstanceType) GongCopy() GongstructIF {
	newInstance := new(FlightConditionCircumstanceType)
	flightconditioncircumstancetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (flightconditioncombinationpropertytype *FlightConditionCombinationPropertyType) GongCopy() GongstructIF {
	newInstance := new(FlightConditionCombinationPropertyType)
	flightconditioncombinationpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (flightconditioncombinationtype *FlightConditionCombinationType) GongCopy() GongstructIF {
	newInstance := new(FlightConditionCombinationType)
	flightconditioncombinationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (flightconditionelementpropertytype *FlightConditionElementPropertyType) GongCopy() GongstructIF {
	newInstance := new(FlightConditionElementPropertyType)
	flightconditionelementpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (flightconditionelementtype *FlightConditionElementType) GongCopy() GongstructIF {
	newInstance := new(FlightConditionElementType)
	flightconditionelementtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (flightrestrictionlevelpropertytype *FlightRestrictionLevelPropertyType) GongCopy() GongstructIF {
	newInstance := new(FlightRestrictionLevelPropertyType)
	flightrestrictionlevelpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (flightrestrictionleveltype *FlightRestrictionLevelType) GongCopy() GongstructIF {
	newInstance := new(FlightRestrictionLevelType)
	flightrestrictionleveltype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (flightrestrictionroutepropertytype *FlightRestrictionRoutePropertyType) GongCopy() GongstructIF {
	newInstance := new(FlightRestrictionRoutePropertyType)
	flightrestrictionroutepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (flightrestrictionroutetype *FlightRestrictionRouteType) GongCopy() GongstructIF {
	newInstance := new(FlightRestrictionRouteType)
	flightrestrictionroutetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (flightrestrictiontimeslicepropertytype *FlightRestrictionTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(FlightRestrictionTimeSlicePropertyType)
	flightrestrictiontimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (flightrestrictiontimeslicetype *FlightRestrictionTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(FlightRestrictionTimeSliceType)
	flightrestrictiontimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (flightrestrictiontype *FlightRestrictionType) GongCopy() GongstructIF {
	newInstance := new(FlightRestrictionType)
	flightrestrictiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (flightroutingelementpropertytype *FlightRoutingElementPropertyType) GongCopy() GongstructIF {
	newInstance := new(FlightRoutingElementPropertyType)
	flightroutingelementpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (flightroutingelementtype *FlightRoutingElementType) GongCopy() GongstructIF {
	newInstance := new(FlightRoutingElementType)
	flightroutingelementtype.GongCopyBasicFields(newInstance)
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

func (meteorologypropertytype *MeteorologyPropertyType) GongCopy() GongstructIF {
	newInstance := new(MeteorologyPropertyType)
	meteorologypropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (meteorologytype *MeteorologyType) GongCopy() GongstructIF {
	newInstance := new(MeteorologyType)
	meteorologytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (navaidpropertytype *NavaidPropertyType) GongCopy() GongstructIF {
	newInstance := new(NavaidPropertyType)
	navaidpropertytype.GongCopyBasicFields(newInstance)
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

func (obstacleassessmentareapropertytype *ObstacleAssessmentAreaPropertyType) GongCopy() GongstructIF {
	newInstance := new(ObstacleAssessmentAreaPropertyType)
	obstacleassessmentareapropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (obstacleassessmentareatype *ObstacleAssessmentAreaType) GongCopy() GongstructIF {
	newInstance := new(ObstacleAssessmentAreaType)
	obstacleassessmentareatype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (obstacleplacementpropertytype *ObstaclePlacementPropertyType) GongCopy() GongstructIF {
	newInstance := new(ObstaclePlacementPropertyType)
	obstacleplacementpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (obstacleplacementtype *ObstaclePlacementType) GongCopy() GongstructIF {
	newInstance := new(ObstaclePlacementType)
	obstacleplacementtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (obstructionpropertytype *ObstructionPropertyType) GongCopy() GongstructIF {
	newInstance := new(ObstructionPropertyType)
	obstructionpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (obstructiontype *ObstructionType) GongCopy() GongstructIF {
	newInstance := new(ObstructionType)
	obstructiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (onlinecontactpropertytype *OnlineContactPropertyType) GongCopy() GongstructIF {
	newInstance := new(OnlineContactPropertyType)
	onlinecontactpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (onlinecontacttype *OnlineContactType) GongCopy() GongstructIF {
	newInstance := new(OnlineContactType)
	onlinecontacttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongCopy() GongstructIF {
	newInstance := new(OrganisationAuthorityPropertyType)
	organisationauthoritypropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (pointpropertytype *PointPropertyType) GongCopy() GongstructIF {
	newInstance := new(PointPropertyType)
	pointpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (pointreferencepropertytype *PointReferencePropertyType) GongCopy() GongstructIF {
	newInstance := new(PointReferencePropertyType)
	pointreferencepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (pointreferencetype *PointReferenceType) GongCopy() GongstructIF {
	newInstance := new(PointReferenceType)
	pointreferencetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (pointtype *PointType) GongCopy() GongstructIF {
	newInstance := new(PointType)
	pointtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (postaladdresspropertytype *PostalAddressPropertyType) GongCopy() GongstructIF {
	newInstance := new(PostalAddressPropertyType)
	postaladdresspropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (postaladdresstype *PostalAddressType) GongCopy() GongstructIF {
	newInstance := new(PostalAddressType)
	postaladdresstype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (radiofrequencyareapropertytype *RadioFrequencyAreaPropertyType) GongCopy() GongstructIF {
	newInstance := new(RadioFrequencyAreaPropertyType)
	radiofrequencyareapropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (routeavailabilitypropertytype *RouteAvailabilityPropertyType) GongCopy() GongstructIF {
	newInstance := new(RouteAvailabilityPropertyType)
	routeavailabilitypropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (routeavailabilitytype *RouteAvailabilityType) GongCopy() GongstructIF {
	newInstance := new(RouteAvailabilityType)
	routeavailabilitytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (routedmetimeslicepropertytype *RouteDMETimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(RouteDMETimeSlicePropertyType)
	routedmetimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (routedmetimeslicetype *RouteDMETimeSliceType) GongCopy() GongstructIF {
	newInstance := new(RouteDMETimeSliceType)
	routedmetimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (routedmetype *RouteDMEType) GongCopy() GongstructIF {
	newInstance := new(RouteDMEType)
	routedmetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (routeportionpropertytype *RoutePortionPropertyType) GongCopy() GongstructIF {
	newInstance := new(RoutePortionPropertyType)
	routeportionpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (routeportiontype *RoutePortionType) GongCopy() GongstructIF {
	newInstance := new(RoutePortionType)
	routeportiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (routepropertytype *RoutePropertyType) GongCopy() GongstructIF {
	newInstance := new(RoutePropertyType)
	routepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (routesegmenttimeslicepropertytype *RouteSegmentTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(RouteSegmentTimeSlicePropertyType)
	routesegmenttimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (routesegmenttimeslicetype *RouteSegmentTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(RouteSegmentTimeSliceType)
	routesegmenttimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (routesegmenttype *RouteSegmentType) GongCopy() GongstructIF {
	newInstance := new(RouteSegmentType)
	routesegmenttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (routetimeslicepropertytype *RouteTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(RouteTimeSlicePropertyType)
	routetimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (routetimeslicetype *RouteTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(RouteTimeSliceType)
	routetimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (routetype *RouteType) GongCopy() GongstructIF {
	newInstance := new(RouteType)
	routetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (rulesprocedurestimeslicepropertytype *RulesProceduresTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(RulesProceduresTimeSlicePropertyType)
	rulesprocedurestimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (rulesprocedurestimeslicetype *RulesProceduresTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(RulesProceduresTimeSliceType)
	rulesprocedurestimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (rulesprocedurestype *RulesProceduresType) GongCopy() GongstructIF {
	newInstance := new(RulesProceduresType)
	rulesprocedurestype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongCopy() GongstructIF {
	newInstance := new(RunwayCentrelinePointPropertyType)
	runwaycentrelinepointpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (standardinstrumentarrivalpropertytype *StandardInstrumentArrivalPropertyType) GongCopy() GongstructIF {
	newInstance := new(StandardInstrumentArrivalPropertyType)
	standardinstrumentarrivalpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (standardinstrumentdeparturepropertytype *StandardInstrumentDeparturePropertyType) GongCopy() GongstructIF {
	newInstance := new(StandardInstrumentDeparturePropertyType)
	standardinstrumentdeparturepropertytype.GongCopyBasicFields(newInstance)
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

func (telephonecontactpropertytype *TelephoneContactPropertyType) GongCopy() GongstructIF {
	newInstance := new(TelephoneContactPropertyType)
	telephonecontactpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (telephonecontacttype *TelephoneContactType) GongCopy() GongstructIF {
	newInstance := new(TelephoneContactType)
	telephonecontacttype.GongCopyBasicFields(newInstance)
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

func (touchdownliftoffpropertytype *TouchDownLiftOffPropertyType) GongCopy() GongstructIF {
	newInstance := new(TouchDownLiftOffPropertyType)
	touchdownliftoffpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongCopy() GongstructIF {
	newInstance := new(VerticalStructurePropertyType)
	verticalstructurepropertytype.GongCopyBasicFields(newInstance)
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

func (a_extension_15 *A_extension_15) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_15)
}

func (a_extension_16 *A_extension_16) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_16)
}

func (a_extension_17 *A_extension_17) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_17)
}

func (a_extension_18 *A_extension_18) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_18)
}

func (a_extension_19 *A_extension_19) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_19)
}

func (a_extension_2 *A_extension_2) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_2)
}

func (a_extension_20 *A_extension_20) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_20)
}

func (a_extension_21 *A_extension_21) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_21)
}

func (a_extension_22 *A_extension_22) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_22)
}

func (a_extension_23 *A_extension_23) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_23)
}

func (a_extension_24 *A_extension_24) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_24)
}

func (a_extension_25 *A_extension_25) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_25)
}

func (a_extension_26 *A_extension_26) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_26)
}

func (a_extension_27 *A_extension_27) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_27)
}

func (a_extension_28 *A_extension_28) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_28)
}

func (a_extension_29 *A_extension_29) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_29)
}

func (a_extension_3 *A_extension_3) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_3)
}

func (a_extension_30 *A_extension_30) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_30)
}

func (a_extension_31 *A_extension_31) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_31)
}

func (a_extension_32 *A_extension_32) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_32)
}

func (a_extension_33 *A_extension_33) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_33)
}

func (a_extension_34 *A_extension_34) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_34)
}

func (a_extension_35 *A_extension_35) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_35)
}

func (a_extension_36 *A_extension_36) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_36)
}

func (a_extension_37 *A_extension_37) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_37)
}

func (a_extension_38 *A_extension_38) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_38)
}

func (a_extension_39 *A_extension_39) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_39)
}

func (a_extension_4 *A_extension_4) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_4)
}

func (a_extension_40 *A_extension_40) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_40)
}

func (a_extension_41 *A_extension_41) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_41)
}

func (a_extension_42 *A_extension_42) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_extension_42)
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

func (abstractdirectflighttype *AbstractDirectFlightType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractdirectflighttype)
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractpropertieswithscheduletype)
}

func (abstractsegmentpointtype *AbstractSegmentPointType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractsegmentpointtype)
}

func (aerialrefuellinganchorpropertytype *AerialRefuellingAnchorPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, aerialrefuellinganchorpropertytype)
}

func (aerialrefuellinganchortype *AerialRefuellingAnchorType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, aerialrefuellinganchortype)
}

func (aerialrefuellingpointpropertytype *AerialRefuellingPointPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, aerialrefuellingpointpropertytype)
}

func (aerialrefuellingpointtype *AerialRefuellingPointType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, aerialrefuellingpointtype)
}

func (aerialrefuellingpropertytype *AerialRefuellingPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, aerialrefuellingpropertytype)
}

func (aerialrefuellingtimeslicepropertytype *AerialRefuellingTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, aerialrefuellingtimeslicepropertytype)
}

func (aerialrefuellingtimeslicetype *AerialRefuellingTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, aerialrefuellingtimeslicetype)
}

func (aerialrefuellingtrackpropertytype *AerialRefuellingTrackPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, aerialrefuellingtrackpropertytype)
}

func (aerialrefuellingtracktype *AerialRefuellingTrackType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, aerialrefuellingtracktype)
}

func (aerialrefuellingtype *AerialRefuellingType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, aerialrefuellingtype)
}

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, aircraftcharacteristicpropertytype)
}

func (aircraftcharacteristictype *AircraftCharacteristicType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, aircraftcharacteristictype)
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airportheliportpropertytype)
}

func (airspacebordercrossingpropertytype *AirspaceBorderCrossingPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airspacebordercrossingpropertytype)
}

func (airspacebordercrossingtimeslicepropertytype *AirspaceBorderCrossingTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airspacebordercrossingtimeslicepropertytype)
}

func (airspacebordercrossingtimeslicetype *AirspaceBorderCrossingTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airspacebordercrossingtimeslicetype)
}

func (airspacebordercrossingtype *AirspaceBorderCrossingType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airspacebordercrossingtype)
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

func (altitudeadjustmentpropertytype *AltitudeAdjustmentPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, altitudeadjustmentpropertytype)
}

func (altitudeadjustmenttype *AltitudeAdjustmentType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, altitudeadjustmenttype)
}

func (angleindicationpropertytype *AngleIndicationPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, angleindicationpropertytype)
}

func (angleusepropertytype *AngleUsePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, angleusepropertytype)
}

func (angleusetype *AngleUseType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, angleusetype)
}

func (authorityforaerialrefuellingpropertytype *AuthorityForAerialRefuellingPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, authorityforaerialrefuellingpropertytype)
}

func (authorityforaerialrefuellingtype *AuthorityForAerialRefuellingType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, authorityforaerialrefuellingtype)
}

func (changeoverpointtimeslicepropertytype *ChangeOverPointTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, changeoverpointtimeslicepropertytype)
}

func (changeoverpointtimeslicetype *ChangeOverPointTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, changeoverpointtimeslicetype)
}

func (changeoverpointtype *ChangeOverPointType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, changeoverpointtype)
}

func (contactinformationpropertytype *ContactInformationPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, contactinformationpropertytype)
}

func (contactinformationtype *ContactInformationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, contactinformationtype)
}

func (curvepropertytype *CurvePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, curvepropertytype)
}

func (curvetype *CurveType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, curvetype)
}

func (dmepropertytype *DMEPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, dmepropertytype)
}

func (designatedpointpropertytype *DesignatedPointPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, designatedpointpropertytype)
}

func (directflightclasstype *DirectFlightClassType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, directflightclasstype)
}

func (directflightpropertytype *DirectFlightPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, directflightpropertytype)
}

func (directflightsegmentpropertytype *DirectFlightSegmentPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, directflightsegmentpropertytype)
}

func (directflightsegmenttype *DirectFlightSegmentType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, directflightsegmenttype)
}

func (distanceindicationpropertytype *DistanceIndicationPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, distanceindicationpropertytype)
}

func (elevatedcurvetype *ElevatedCurveType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, elevatedcurvetype)
}

func (elevatedpointtype *ElevatedPointType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, elevatedpointtype)
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, elevatedsurfacetype)
}

func (enroutesegmentpointpropertytype *EnRouteSegmentPointPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, enroutesegmentpointpropertytype)
}

func (enroutesegmentpointtype *EnRouteSegmentPointType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, enroutesegmentpointtype)
}

func (flightcharacteristicpropertytype *FlightCharacteristicPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, flightcharacteristicpropertytype)
}

func (flightcharacteristictype *FlightCharacteristicType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, flightcharacteristictype)
}

func (flightconditioncircumstancepropertytype *FlightConditionCircumstancePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, flightconditioncircumstancepropertytype)
}

func (flightconditioncircumstancetype *FlightConditionCircumstanceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, flightconditioncircumstancetype)
}

func (flightconditioncombinationpropertytype *FlightConditionCombinationPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, flightconditioncombinationpropertytype)
}

func (flightconditioncombinationtype *FlightConditionCombinationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, flightconditioncombinationtype)
}

func (flightconditionelementpropertytype *FlightConditionElementPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, flightconditionelementpropertytype)
}

func (flightconditionelementtype *FlightConditionElementType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, flightconditionelementtype)
}

func (flightrestrictionlevelpropertytype *FlightRestrictionLevelPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, flightrestrictionlevelpropertytype)
}

func (flightrestrictionleveltype *FlightRestrictionLevelType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, flightrestrictionleveltype)
}

func (flightrestrictionroutepropertytype *FlightRestrictionRoutePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, flightrestrictionroutepropertytype)
}

func (flightrestrictionroutetype *FlightRestrictionRouteType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, flightrestrictionroutetype)
}

func (flightrestrictiontimeslicepropertytype *FlightRestrictionTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, flightrestrictiontimeslicepropertytype)
}

func (flightrestrictiontimeslicetype *FlightRestrictionTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, flightrestrictiontimeslicetype)
}

func (flightrestrictiontype *FlightRestrictionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, flightrestrictiontype)
}

func (flightroutingelementpropertytype *FlightRoutingElementPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, flightroutingelementpropertytype)
}

func (flightroutingelementtype *FlightRoutingElementType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, flightroutingelementtype)
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, linguisticnotepropertytype)
}

func (linguisticnotetype *LinguisticNoteType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, linguisticnotetype)
}

func (meteorologypropertytype *MeteorologyPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, meteorologypropertytype)
}

func (meteorologytype *MeteorologyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, meteorologytype)
}

func (navaidpropertytype *NavaidPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, navaidpropertytype)
}

func (notepropertytype *NotePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, notepropertytype)
}

func (notetype *NoteType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, notetype)
}

func (obstacleassessmentareapropertytype *ObstacleAssessmentAreaPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, obstacleassessmentareapropertytype)
}

func (obstacleassessmentareatype *ObstacleAssessmentAreaType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, obstacleassessmentareatype)
}

func (obstacleplacementpropertytype *ObstaclePlacementPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, obstacleplacementpropertytype)
}

func (obstacleplacementtype *ObstaclePlacementType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, obstacleplacementtype)
}

func (obstructionpropertytype *ObstructionPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, obstructionpropertytype)
}

func (obstructiontype *ObstructionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, obstructiontype)
}

func (onlinecontactpropertytype *OnlineContactPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, onlinecontactpropertytype)
}

func (onlinecontacttype *OnlineContactType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, onlinecontacttype)
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, organisationauthoritypropertytype)
}

func (pointpropertytype *PointPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, pointpropertytype)
}

func (pointreferencepropertytype *PointReferencePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, pointreferencepropertytype)
}

func (pointreferencetype *PointReferenceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, pointreferencetype)
}

func (pointtype *PointType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, pointtype)
}

func (postaladdresspropertytype *PostalAddressPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, postaladdresspropertytype)
}

func (postaladdresstype *PostalAddressType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, postaladdresstype)
}

func (radiofrequencyareapropertytype *RadioFrequencyAreaPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, radiofrequencyareapropertytype)
}

func (routeavailabilitypropertytype *RouteAvailabilityPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, routeavailabilitypropertytype)
}

func (routeavailabilitytype *RouteAvailabilityType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, routeavailabilitytype)
}

func (routedmetimeslicepropertytype *RouteDMETimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, routedmetimeslicepropertytype)
}

func (routedmetimeslicetype *RouteDMETimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, routedmetimeslicetype)
}

func (routedmetype *RouteDMEType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, routedmetype)
}

func (routeportionpropertytype *RoutePortionPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, routeportionpropertytype)
}

func (routeportiontype *RoutePortionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, routeportiontype)
}

func (routepropertytype *RoutePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, routepropertytype)
}

func (routesegmenttimeslicepropertytype *RouteSegmentTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, routesegmenttimeslicepropertytype)
}

func (routesegmenttimeslicetype *RouteSegmentTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, routesegmenttimeslicetype)
}

func (routesegmenttype *RouteSegmentType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, routesegmenttype)
}

func (routetimeslicepropertytype *RouteTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, routetimeslicepropertytype)
}

func (routetimeslicetype *RouteTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, routetimeslicetype)
}

func (routetype *RouteType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, routetype)
}

func (rulesprocedurestimeslicepropertytype *RulesProceduresTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, rulesprocedurestimeslicepropertytype)
}

func (rulesprocedurestimeslicetype *RulesProceduresTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, rulesprocedurestimeslicetype)
}

func (rulesprocedurestype *RulesProceduresType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, rulesprocedurestype)
}

func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, runwaycentrelinepointpropertytype)
}

func (standardinstrumentarrivalpropertytype *StandardInstrumentArrivalPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, standardinstrumentarrivalpropertytype)
}

func (standardinstrumentdeparturepropertytype *StandardInstrumentDeparturePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, standardinstrumentdeparturepropertytype)
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

func (telephonecontactpropertytype *TelephoneContactPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, telephonecontactpropertytype)
}

func (telephonecontacttype *TelephoneContactType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, telephonecontacttype)
}

func (timesheetpropertytype *TimesheetPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, timesheetpropertytype)
}

func (timesheettype *TimesheetType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, timesheettype)
}

func (touchdownliftoffpropertytype *TouchDownLiftOffPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, touchdownliftoffpropertytype)
}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, verticalstructurepropertytype)
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
		stage.A_extension_15s,
		stage.A_extension_15_stagedOrder,
		stage.A_extension_15s_reference,
		&stage.A_extension_15s_referenceOrder,
		stage.A_extension_15s_instance,
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
		stage.A_extension_16s,
		stage.A_extension_16_stagedOrder,
		stage.A_extension_16s_reference,
		&stage.A_extension_16s_referenceOrder,
		stage.A_extension_16s_instance,
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
		stage.A_extension_17s,
		stage.A_extension_17_stagedOrder,
		stage.A_extension_17s_reference,
		&stage.A_extension_17s_referenceOrder,
		stage.A_extension_17s_instance,
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
		stage.A_extension_18s,
		stage.A_extension_18_stagedOrder,
		stage.A_extension_18s_reference,
		&stage.A_extension_18s_referenceOrder,
		stage.A_extension_18s_instance,
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
		stage.A_extension_19s,
		stage.A_extension_19_stagedOrder,
		stage.A_extension_19s_reference,
		&stage.A_extension_19s_referenceOrder,
		stage.A_extension_19s_instance,
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
		stage.A_extension_20s,
		stage.A_extension_20_stagedOrder,
		stage.A_extension_20s_reference,
		&stage.A_extension_20s_referenceOrder,
		stage.A_extension_20s_instance,
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
		stage.A_extension_21s,
		stage.A_extension_21_stagedOrder,
		stage.A_extension_21s_reference,
		&stage.A_extension_21s_referenceOrder,
		stage.A_extension_21s_instance,
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
		stage.A_extension_22s,
		stage.A_extension_22_stagedOrder,
		stage.A_extension_22s_reference,
		&stage.A_extension_22s_referenceOrder,
		stage.A_extension_22s_instance,
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
		stage.A_extension_23s,
		stage.A_extension_23_stagedOrder,
		stage.A_extension_23s_reference,
		&stage.A_extension_23s_referenceOrder,
		stage.A_extension_23s_instance,
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
		stage.A_extension_24s,
		stage.A_extension_24_stagedOrder,
		stage.A_extension_24s_reference,
		&stage.A_extension_24s_referenceOrder,
		stage.A_extension_24s_instance,
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
		stage.A_extension_25s,
		stage.A_extension_25_stagedOrder,
		stage.A_extension_25s_reference,
		&stage.A_extension_25s_referenceOrder,
		stage.A_extension_25s_instance,
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
		stage.A_extension_26s,
		stage.A_extension_26_stagedOrder,
		stage.A_extension_26s_reference,
		&stage.A_extension_26s_referenceOrder,
		stage.A_extension_26s_instance,
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
		stage.A_extension_27s,
		stage.A_extension_27_stagedOrder,
		stage.A_extension_27s_reference,
		&stage.A_extension_27s_referenceOrder,
		stage.A_extension_27s_instance,
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
		stage.A_extension_28s,
		stage.A_extension_28_stagedOrder,
		stage.A_extension_28s_reference,
		&stage.A_extension_28s_referenceOrder,
		stage.A_extension_28s_instance,
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
		stage.A_extension_29s,
		stage.A_extension_29_stagedOrder,
		stage.A_extension_29s_reference,
		&stage.A_extension_29s_referenceOrder,
		stage.A_extension_29s_instance,
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
		stage.A_extension_30s,
		stage.A_extension_30_stagedOrder,
		stage.A_extension_30s_reference,
		&stage.A_extension_30s_referenceOrder,
		stage.A_extension_30s_instance,
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
		stage.A_extension_31s,
		stage.A_extension_31_stagedOrder,
		stage.A_extension_31s_reference,
		&stage.A_extension_31s_referenceOrder,
		stage.A_extension_31s_instance,
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
		stage.A_extension_32s,
		stage.A_extension_32_stagedOrder,
		stage.A_extension_32s_reference,
		&stage.A_extension_32s_referenceOrder,
		stage.A_extension_32s_instance,
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
		stage.A_extension_33s,
		stage.A_extension_33_stagedOrder,
		stage.A_extension_33s_reference,
		&stage.A_extension_33s_referenceOrder,
		stage.A_extension_33s_instance,
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
		stage.A_extension_34s,
		stage.A_extension_34_stagedOrder,
		stage.A_extension_34s_reference,
		&stage.A_extension_34s_referenceOrder,
		stage.A_extension_34s_instance,
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
		stage.A_extension_35s,
		stage.A_extension_35_stagedOrder,
		stage.A_extension_35s_reference,
		&stage.A_extension_35s_referenceOrder,
		stage.A_extension_35s_instance,
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
		stage.A_extension_36s,
		stage.A_extension_36_stagedOrder,
		stage.A_extension_36s_reference,
		&stage.A_extension_36s_referenceOrder,
		stage.A_extension_36s_instance,
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
		stage.A_extension_37s,
		stage.A_extension_37_stagedOrder,
		stage.A_extension_37s_reference,
		&stage.A_extension_37s_referenceOrder,
		stage.A_extension_37s_instance,
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
		stage.A_extension_38s,
		stage.A_extension_38_stagedOrder,
		stage.A_extension_38s_reference,
		&stage.A_extension_38s_referenceOrder,
		stage.A_extension_38s_instance,
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
		stage.A_extension_39s,
		stage.A_extension_39_stagedOrder,
		stage.A_extension_39s_reference,
		&stage.A_extension_39s_referenceOrder,
		stage.A_extension_39s_instance,
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
		stage.A_extension_40s,
		stage.A_extension_40_stagedOrder,
		stage.A_extension_40s_reference,
		&stage.A_extension_40s_referenceOrder,
		stage.A_extension_40s_instance,
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
		stage.A_extension_41s,
		stage.A_extension_41_stagedOrder,
		stage.A_extension_41s_reference,
		&stage.A_extension_41s_referenceOrder,
		stage.A_extension_41s_instance,
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
		stage.A_extension_42s,
		stage.A_extension_42_stagedOrder,
		stage.A_extension_42s_reference,
		&stage.A_extension_42s_referenceOrder,
		stage.A_extension_42s_instance,
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
		stage.AbstractDirectFlightTypes,
		stage.AbstractDirectFlightType_stagedOrder,
		stage.AbstractDirectFlightTypes_reference,
		&stage.AbstractDirectFlightTypes_referenceOrder,
		stage.AbstractDirectFlightTypes_instance,
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
		stage.AbstractSegmentPointTypes,
		stage.AbstractSegmentPointType_stagedOrder,
		stage.AbstractSegmentPointTypes_reference,
		&stage.AbstractSegmentPointTypes_referenceOrder,
		stage.AbstractSegmentPointTypes_instance,
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
		stage.AerialRefuellingAnchorPropertyTypes,
		stage.AerialRefuellingAnchorPropertyType_stagedOrder,
		stage.AerialRefuellingAnchorPropertyTypes_reference,
		&stage.AerialRefuellingAnchorPropertyTypes_referenceOrder,
		stage.AerialRefuellingAnchorPropertyTypes_instance,
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
		stage.AerialRefuellingAnchorTypes,
		stage.AerialRefuellingAnchorType_stagedOrder,
		stage.AerialRefuellingAnchorTypes_reference,
		&stage.AerialRefuellingAnchorTypes_referenceOrder,
		stage.AerialRefuellingAnchorTypes_instance,
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
		stage.AerialRefuellingPointPropertyTypes,
		stage.AerialRefuellingPointPropertyType_stagedOrder,
		stage.AerialRefuellingPointPropertyTypes_reference,
		&stage.AerialRefuellingPointPropertyTypes_referenceOrder,
		stage.AerialRefuellingPointPropertyTypes_instance,
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
		stage.AerialRefuellingPointTypes,
		stage.AerialRefuellingPointType_stagedOrder,
		stage.AerialRefuellingPointTypes_reference,
		&stage.AerialRefuellingPointTypes_referenceOrder,
		stage.AerialRefuellingPointTypes_instance,
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
		stage.AerialRefuellingPropertyTypes,
		stage.AerialRefuellingPropertyType_stagedOrder,
		stage.AerialRefuellingPropertyTypes_reference,
		&stage.AerialRefuellingPropertyTypes_referenceOrder,
		stage.AerialRefuellingPropertyTypes_instance,
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
		stage.AerialRefuellingTimeSlicePropertyTypes,
		stage.AerialRefuellingTimeSlicePropertyType_stagedOrder,
		stage.AerialRefuellingTimeSlicePropertyTypes_reference,
		&stage.AerialRefuellingTimeSlicePropertyTypes_referenceOrder,
		stage.AerialRefuellingTimeSlicePropertyTypes_instance,
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
		stage.AerialRefuellingTimeSliceTypes,
		stage.AerialRefuellingTimeSliceType_stagedOrder,
		stage.AerialRefuellingTimeSliceTypes_reference,
		&stage.AerialRefuellingTimeSliceTypes_referenceOrder,
		stage.AerialRefuellingTimeSliceTypes_instance,
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
		stage.AerialRefuellingTrackPropertyTypes,
		stage.AerialRefuellingTrackPropertyType_stagedOrder,
		stage.AerialRefuellingTrackPropertyTypes_reference,
		&stage.AerialRefuellingTrackPropertyTypes_referenceOrder,
		stage.AerialRefuellingTrackPropertyTypes_instance,
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
		stage.AerialRefuellingTrackTypes,
		stage.AerialRefuellingTrackType_stagedOrder,
		stage.AerialRefuellingTrackTypes_reference,
		&stage.AerialRefuellingTrackTypes_referenceOrder,
		stage.AerialRefuellingTrackTypes_instance,
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
		stage.AerialRefuellingTypes,
		stage.AerialRefuellingType_stagedOrder,
		stage.AerialRefuellingTypes_reference,
		&stage.AerialRefuellingTypes_referenceOrder,
		stage.AerialRefuellingTypes_instance,
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
		stage.AirportHeliportPropertyTypes,
		stage.AirportHeliportPropertyType_stagedOrder,
		stage.AirportHeliportPropertyTypes_reference,
		&stage.AirportHeliportPropertyTypes_referenceOrder,
		stage.AirportHeliportPropertyTypes_instance,
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
		stage.AirspaceBorderCrossingPropertyTypes,
		stage.AirspaceBorderCrossingPropertyType_stagedOrder,
		stage.AirspaceBorderCrossingPropertyTypes_reference,
		&stage.AirspaceBorderCrossingPropertyTypes_referenceOrder,
		stage.AirspaceBorderCrossingPropertyTypes_instance,
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
		stage.AirspaceBorderCrossingTimeSlicePropertyTypes,
		stage.AirspaceBorderCrossingTimeSlicePropertyType_stagedOrder,
		stage.AirspaceBorderCrossingTimeSlicePropertyTypes_reference,
		&stage.AirspaceBorderCrossingTimeSlicePropertyTypes_referenceOrder,
		stage.AirspaceBorderCrossingTimeSlicePropertyTypes_instance,
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
		stage.AirspaceBorderCrossingTimeSliceTypes,
		stage.AirspaceBorderCrossingTimeSliceType_stagedOrder,
		stage.AirspaceBorderCrossingTimeSliceTypes_reference,
		&stage.AirspaceBorderCrossingTimeSliceTypes_referenceOrder,
		stage.AirspaceBorderCrossingTimeSliceTypes_instance,
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
		stage.AirspaceBorderCrossingTypes,
		stage.AirspaceBorderCrossingType_stagedOrder,
		stage.AirspaceBorderCrossingTypes_reference,
		&stage.AirspaceBorderCrossingTypes_referenceOrder,
		stage.AirspaceBorderCrossingTypes_instance,
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
		stage.AltitudeAdjustmentPropertyTypes,
		stage.AltitudeAdjustmentPropertyType_stagedOrder,
		stage.AltitudeAdjustmentPropertyTypes_reference,
		&stage.AltitudeAdjustmentPropertyTypes_referenceOrder,
		stage.AltitudeAdjustmentPropertyTypes_instance,
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
		stage.AltitudeAdjustmentTypes,
		stage.AltitudeAdjustmentType_stagedOrder,
		stage.AltitudeAdjustmentTypes_reference,
		&stage.AltitudeAdjustmentTypes_referenceOrder,
		stage.AltitudeAdjustmentTypes_instance,
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
		stage.AngleIndicationPropertyTypes,
		stage.AngleIndicationPropertyType_stagedOrder,
		stage.AngleIndicationPropertyTypes_reference,
		&stage.AngleIndicationPropertyTypes_referenceOrder,
		stage.AngleIndicationPropertyTypes_instance,
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
		stage.AngleUsePropertyTypes,
		stage.AngleUsePropertyType_stagedOrder,
		stage.AngleUsePropertyTypes_reference,
		&stage.AngleUsePropertyTypes_referenceOrder,
		stage.AngleUsePropertyTypes_instance,
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
		stage.AngleUseTypes,
		stage.AngleUseType_stagedOrder,
		stage.AngleUseTypes_reference,
		&stage.AngleUseTypes_referenceOrder,
		stage.AngleUseTypes_instance,
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
		stage.AuthorityForAerialRefuellingPropertyTypes,
		stage.AuthorityForAerialRefuellingPropertyType_stagedOrder,
		stage.AuthorityForAerialRefuellingPropertyTypes_reference,
		&stage.AuthorityForAerialRefuellingPropertyTypes_referenceOrder,
		stage.AuthorityForAerialRefuellingPropertyTypes_instance,
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
		stage.AuthorityForAerialRefuellingTypes,
		stage.AuthorityForAerialRefuellingType_stagedOrder,
		stage.AuthorityForAerialRefuellingTypes_reference,
		&stage.AuthorityForAerialRefuellingTypes_referenceOrder,
		stage.AuthorityForAerialRefuellingTypes_instance,
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
		stage.ChangeOverPointTimeSlicePropertyTypes,
		stage.ChangeOverPointTimeSlicePropertyType_stagedOrder,
		stage.ChangeOverPointTimeSlicePropertyTypes_reference,
		&stage.ChangeOverPointTimeSlicePropertyTypes_referenceOrder,
		stage.ChangeOverPointTimeSlicePropertyTypes_instance,
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
		stage.ChangeOverPointTimeSliceTypes,
		stage.ChangeOverPointTimeSliceType_stagedOrder,
		stage.ChangeOverPointTimeSliceTypes_reference,
		&stage.ChangeOverPointTimeSliceTypes_referenceOrder,
		stage.ChangeOverPointTimeSliceTypes_instance,
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
		stage.ChangeOverPointTypes,
		stage.ChangeOverPointType_stagedOrder,
		stage.ChangeOverPointTypes_reference,
		&stage.ChangeOverPointTypes_referenceOrder,
		stage.ChangeOverPointTypes_instance,
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
		stage.ContactInformationPropertyTypes,
		stage.ContactInformationPropertyType_stagedOrder,
		stage.ContactInformationPropertyTypes_reference,
		&stage.ContactInformationPropertyTypes_referenceOrder,
		stage.ContactInformationPropertyTypes_instance,
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
		stage.ContactInformationTypes,
		stage.ContactInformationType_stagedOrder,
		stage.ContactInformationTypes_reference,
		&stage.ContactInformationTypes_referenceOrder,
		stage.ContactInformationTypes_instance,
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
		stage.DMEPropertyTypes,
		stage.DMEPropertyType_stagedOrder,
		stage.DMEPropertyTypes_reference,
		&stage.DMEPropertyTypes_referenceOrder,
		stage.DMEPropertyTypes_instance,
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
		stage.DesignatedPointPropertyTypes,
		stage.DesignatedPointPropertyType_stagedOrder,
		stage.DesignatedPointPropertyTypes_reference,
		&stage.DesignatedPointPropertyTypes_referenceOrder,
		stage.DesignatedPointPropertyTypes_instance,
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
		stage.DirectFlightClassTypes,
		stage.DirectFlightClassType_stagedOrder,
		stage.DirectFlightClassTypes_reference,
		&stage.DirectFlightClassTypes_referenceOrder,
		stage.DirectFlightClassTypes_instance,
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
		stage.DirectFlightPropertyTypes,
		stage.DirectFlightPropertyType_stagedOrder,
		stage.DirectFlightPropertyTypes_reference,
		&stage.DirectFlightPropertyTypes_referenceOrder,
		stage.DirectFlightPropertyTypes_instance,
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
		stage.DirectFlightSegmentPropertyTypes,
		stage.DirectFlightSegmentPropertyType_stagedOrder,
		stage.DirectFlightSegmentPropertyTypes_reference,
		&stage.DirectFlightSegmentPropertyTypes_referenceOrder,
		stage.DirectFlightSegmentPropertyTypes_instance,
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
		stage.DirectFlightSegmentTypes,
		stage.DirectFlightSegmentType_stagedOrder,
		stage.DirectFlightSegmentTypes_reference,
		&stage.DirectFlightSegmentTypes_referenceOrder,
		stage.DirectFlightSegmentTypes_instance,
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
		stage.DistanceIndicationPropertyTypes,
		stage.DistanceIndicationPropertyType_stagedOrder,
		stage.DistanceIndicationPropertyTypes_reference,
		&stage.DistanceIndicationPropertyTypes_referenceOrder,
		stage.DistanceIndicationPropertyTypes_instance,
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
		stage.ElevatedPointTypes,
		stage.ElevatedPointType_stagedOrder,
		stage.ElevatedPointTypes_reference,
		&stage.ElevatedPointTypes_referenceOrder,
		stage.ElevatedPointTypes_instance,
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
		stage.EnRouteSegmentPointPropertyTypes,
		stage.EnRouteSegmentPointPropertyType_stagedOrder,
		stage.EnRouteSegmentPointPropertyTypes_reference,
		&stage.EnRouteSegmentPointPropertyTypes_referenceOrder,
		stage.EnRouteSegmentPointPropertyTypes_instance,
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
		stage.EnRouteSegmentPointTypes,
		stage.EnRouteSegmentPointType_stagedOrder,
		stage.EnRouteSegmentPointTypes_reference,
		&stage.EnRouteSegmentPointTypes_referenceOrder,
		stage.EnRouteSegmentPointTypes_instance,
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
		stage.FlightCharacteristicPropertyTypes,
		stage.FlightCharacteristicPropertyType_stagedOrder,
		stage.FlightCharacteristicPropertyTypes_reference,
		&stage.FlightCharacteristicPropertyTypes_referenceOrder,
		stage.FlightCharacteristicPropertyTypes_instance,
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
		stage.FlightCharacteristicTypes,
		stage.FlightCharacteristicType_stagedOrder,
		stage.FlightCharacteristicTypes_reference,
		&stage.FlightCharacteristicTypes_referenceOrder,
		stage.FlightCharacteristicTypes_instance,
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
		stage.FlightConditionCircumstancePropertyTypes,
		stage.FlightConditionCircumstancePropertyType_stagedOrder,
		stage.FlightConditionCircumstancePropertyTypes_reference,
		&stage.FlightConditionCircumstancePropertyTypes_referenceOrder,
		stage.FlightConditionCircumstancePropertyTypes_instance,
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
		stage.FlightConditionCircumstanceTypes,
		stage.FlightConditionCircumstanceType_stagedOrder,
		stage.FlightConditionCircumstanceTypes_reference,
		&stage.FlightConditionCircumstanceTypes_referenceOrder,
		stage.FlightConditionCircumstanceTypes_instance,
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
		stage.FlightConditionCombinationPropertyTypes,
		stage.FlightConditionCombinationPropertyType_stagedOrder,
		stage.FlightConditionCombinationPropertyTypes_reference,
		&stage.FlightConditionCombinationPropertyTypes_referenceOrder,
		stage.FlightConditionCombinationPropertyTypes_instance,
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
		stage.FlightConditionCombinationTypes,
		stage.FlightConditionCombinationType_stagedOrder,
		stage.FlightConditionCombinationTypes_reference,
		&stage.FlightConditionCombinationTypes_referenceOrder,
		stage.FlightConditionCombinationTypes_instance,
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
		stage.FlightConditionElementPropertyTypes,
		stage.FlightConditionElementPropertyType_stagedOrder,
		stage.FlightConditionElementPropertyTypes_reference,
		&stage.FlightConditionElementPropertyTypes_referenceOrder,
		stage.FlightConditionElementPropertyTypes_instance,
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
		stage.FlightConditionElementTypes,
		stage.FlightConditionElementType_stagedOrder,
		stage.FlightConditionElementTypes_reference,
		&stage.FlightConditionElementTypes_referenceOrder,
		stage.FlightConditionElementTypes_instance,
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
		stage.FlightRestrictionLevelPropertyTypes,
		stage.FlightRestrictionLevelPropertyType_stagedOrder,
		stage.FlightRestrictionLevelPropertyTypes_reference,
		&stage.FlightRestrictionLevelPropertyTypes_referenceOrder,
		stage.FlightRestrictionLevelPropertyTypes_instance,
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
		stage.FlightRestrictionLevelTypes,
		stage.FlightRestrictionLevelType_stagedOrder,
		stage.FlightRestrictionLevelTypes_reference,
		&stage.FlightRestrictionLevelTypes_referenceOrder,
		stage.FlightRestrictionLevelTypes_instance,
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
		stage.FlightRestrictionRoutePropertyTypes,
		stage.FlightRestrictionRoutePropertyType_stagedOrder,
		stage.FlightRestrictionRoutePropertyTypes_reference,
		&stage.FlightRestrictionRoutePropertyTypes_referenceOrder,
		stage.FlightRestrictionRoutePropertyTypes_instance,
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
		stage.FlightRestrictionRouteTypes,
		stage.FlightRestrictionRouteType_stagedOrder,
		stage.FlightRestrictionRouteTypes_reference,
		&stage.FlightRestrictionRouteTypes_referenceOrder,
		stage.FlightRestrictionRouteTypes_instance,
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
		stage.FlightRestrictionTimeSlicePropertyTypes,
		stage.FlightRestrictionTimeSlicePropertyType_stagedOrder,
		stage.FlightRestrictionTimeSlicePropertyTypes_reference,
		&stage.FlightRestrictionTimeSlicePropertyTypes_referenceOrder,
		stage.FlightRestrictionTimeSlicePropertyTypes_instance,
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
		stage.FlightRestrictionTimeSliceTypes,
		stage.FlightRestrictionTimeSliceType_stagedOrder,
		stage.FlightRestrictionTimeSliceTypes_reference,
		&stage.FlightRestrictionTimeSliceTypes_referenceOrder,
		stage.FlightRestrictionTimeSliceTypes_instance,
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
		stage.FlightRestrictionTypes,
		stage.FlightRestrictionType_stagedOrder,
		stage.FlightRestrictionTypes_reference,
		&stage.FlightRestrictionTypes_referenceOrder,
		stage.FlightRestrictionTypes_instance,
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
		stage.FlightRoutingElementPropertyTypes,
		stage.FlightRoutingElementPropertyType_stagedOrder,
		stage.FlightRoutingElementPropertyTypes_reference,
		&stage.FlightRoutingElementPropertyTypes_referenceOrder,
		stage.FlightRoutingElementPropertyTypes_instance,
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
		stage.FlightRoutingElementTypes,
		stage.FlightRoutingElementType_stagedOrder,
		stage.FlightRoutingElementTypes_reference,
		&stage.FlightRoutingElementTypes_referenceOrder,
		stage.FlightRoutingElementTypes_instance,
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
		stage.MeteorologyPropertyTypes,
		stage.MeteorologyPropertyType_stagedOrder,
		stage.MeteorologyPropertyTypes_reference,
		&stage.MeteorologyPropertyTypes_referenceOrder,
		stage.MeteorologyPropertyTypes_instance,
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
		stage.MeteorologyTypes,
		stage.MeteorologyType_stagedOrder,
		stage.MeteorologyTypes_reference,
		&stage.MeteorologyTypes_referenceOrder,
		stage.MeteorologyTypes_instance,
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
		stage.NavaidPropertyTypes,
		stage.NavaidPropertyType_stagedOrder,
		stage.NavaidPropertyTypes_reference,
		&stage.NavaidPropertyTypes_referenceOrder,
		stage.NavaidPropertyTypes_instance,
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
		stage.ObstacleAssessmentAreaPropertyTypes,
		stage.ObstacleAssessmentAreaPropertyType_stagedOrder,
		stage.ObstacleAssessmentAreaPropertyTypes_reference,
		&stage.ObstacleAssessmentAreaPropertyTypes_referenceOrder,
		stage.ObstacleAssessmentAreaPropertyTypes_instance,
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
		stage.ObstacleAssessmentAreaTypes,
		stage.ObstacleAssessmentAreaType_stagedOrder,
		stage.ObstacleAssessmentAreaTypes_reference,
		&stage.ObstacleAssessmentAreaTypes_referenceOrder,
		stage.ObstacleAssessmentAreaTypes_instance,
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
		stage.ObstaclePlacementPropertyTypes,
		stage.ObstaclePlacementPropertyType_stagedOrder,
		stage.ObstaclePlacementPropertyTypes_reference,
		&stage.ObstaclePlacementPropertyTypes_referenceOrder,
		stage.ObstaclePlacementPropertyTypes_instance,
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
		stage.ObstaclePlacementTypes,
		stage.ObstaclePlacementType_stagedOrder,
		stage.ObstaclePlacementTypes_reference,
		&stage.ObstaclePlacementTypes_referenceOrder,
		stage.ObstaclePlacementTypes_instance,
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
		stage.ObstructionPropertyTypes,
		stage.ObstructionPropertyType_stagedOrder,
		stage.ObstructionPropertyTypes_reference,
		&stage.ObstructionPropertyTypes_referenceOrder,
		stage.ObstructionPropertyTypes_instance,
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
		stage.ObstructionTypes,
		stage.ObstructionType_stagedOrder,
		stage.ObstructionTypes_reference,
		&stage.ObstructionTypes_referenceOrder,
		stage.ObstructionTypes_instance,
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
		stage.OnlineContactPropertyTypes,
		stage.OnlineContactPropertyType_stagedOrder,
		stage.OnlineContactPropertyTypes_reference,
		&stage.OnlineContactPropertyTypes_referenceOrder,
		stage.OnlineContactPropertyTypes_instance,
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
		stage.OnlineContactTypes,
		stage.OnlineContactType_stagedOrder,
		stage.OnlineContactTypes_reference,
		&stage.OnlineContactTypes_referenceOrder,
		stage.OnlineContactTypes_instance,
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
		stage.PointPropertyTypes,
		stage.PointPropertyType_stagedOrder,
		stage.PointPropertyTypes_reference,
		&stage.PointPropertyTypes_referenceOrder,
		stage.PointPropertyTypes_instance,
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
		stage.PointReferencePropertyTypes,
		stage.PointReferencePropertyType_stagedOrder,
		stage.PointReferencePropertyTypes_reference,
		&stage.PointReferencePropertyTypes_referenceOrder,
		stage.PointReferencePropertyTypes_instance,
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
		stage.PointReferenceTypes,
		stage.PointReferenceType_stagedOrder,
		stage.PointReferenceTypes_reference,
		&stage.PointReferenceTypes_referenceOrder,
		stage.PointReferenceTypes_instance,
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
		stage.PointTypes,
		stage.PointType_stagedOrder,
		stage.PointTypes_reference,
		&stage.PointTypes_referenceOrder,
		stage.PointTypes_instance,
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
		stage.PostalAddressPropertyTypes,
		stage.PostalAddressPropertyType_stagedOrder,
		stage.PostalAddressPropertyTypes_reference,
		&stage.PostalAddressPropertyTypes_referenceOrder,
		stage.PostalAddressPropertyTypes_instance,
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
		stage.PostalAddressTypes,
		stage.PostalAddressType_stagedOrder,
		stage.PostalAddressTypes_reference,
		&stage.PostalAddressTypes_referenceOrder,
		stage.PostalAddressTypes_instance,
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
		stage.RadioFrequencyAreaPropertyTypes,
		stage.RadioFrequencyAreaPropertyType_stagedOrder,
		stage.RadioFrequencyAreaPropertyTypes_reference,
		&stage.RadioFrequencyAreaPropertyTypes_referenceOrder,
		stage.RadioFrequencyAreaPropertyTypes_instance,
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
		stage.RouteAvailabilityPropertyTypes,
		stage.RouteAvailabilityPropertyType_stagedOrder,
		stage.RouteAvailabilityPropertyTypes_reference,
		&stage.RouteAvailabilityPropertyTypes_referenceOrder,
		stage.RouteAvailabilityPropertyTypes_instance,
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
		stage.RouteAvailabilityTypes,
		stage.RouteAvailabilityType_stagedOrder,
		stage.RouteAvailabilityTypes_reference,
		&stage.RouteAvailabilityTypes_referenceOrder,
		stage.RouteAvailabilityTypes_instance,
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
		stage.RouteDMETimeSlicePropertyTypes,
		stage.RouteDMETimeSlicePropertyType_stagedOrder,
		stage.RouteDMETimeSlicePropertyTypes_reference,
		&stage.RouteDMETimeSlicePropertyTypes_referenceOrder,
		stage.RouteDMETimeSlicePropertyTypes_instance,
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
		stage.RouteDMETimeSliceTypes,
		stage.RouteDMETimeSliceType_stagedOrder,
		stage.RouteDMETimeSliceTypes_reference,
		&stage.RouteDMETimeSliceTypes_referenceOrder,
		stage.RouteDMETimeSliceTypes_instance,
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
		stage.RouteDMETypes,
		stage.RouteDMEType_stagedOrder,
		stage.RouteDMETypes_reference,
		&stage.RouteDMETypes_referenceOrder,
		stage.RouteDMETypes_instance,
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
		stage.RoutePortionPropertyTypes,
		stage.RoutePortionPropertyType_stagedOrder,
		stage.RoutePortionPropertyTypes_reference,
		&stage.RoutePortionPropertyTypes_referenceOrder,
		stage.RoutePortionPropertyTypes_instance,
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
		stage.RoutePortionTypes,
		stage.RoutePortionType_stagedOrder,
		stage.RoutePortionTypes_reference,
		&stage.RoutePortionTypes_referenceOrder,
		stage.RoutePortionTypes_instance,
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
		stage.RouteSegmentTimeSlicePropertyTypes,
		stage.RouteSegmentTimeSlicePropertyType_stagedOrder,
		stage.RouteSegmentTimeSlicePropertyTypes_reference,
		&stage.RouteSegmentTimeSlicePropertyTypes_referenceOrder,
		stage.RouteSegmentTimeSlicePropertyTypes_instance,
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
		stage.RouteSegmentTimeSliceTypes,
		stage.RouteSegmentTimeSliceType_stagedOrder,
		stage.RouteSegmentTimeSliceTypes_reference,
		&stage.RouteSegmentTimeSliceTypes_referenceOrder,
		stage.RouteSegmentTimeSliceTypes_instance,
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
		stage.RouteSegmentTypes,
		stage.RouteSegmentType_stagedOrder,
		stage.RouteSegmentTypes_reference,
		&stage.RouteSegmentTypes_referenceOrder,
		stage.RouteSegmentTypes_instance,
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
		stage.RouteTimeSlicePropertyTypes,
		stage.RouteTimeSlicePropertyType_stagedOrder,
		stage.RouteTimeSlicePropertyTypes_reference,
		&stage.RouteTimeSlicePropertyTypes_referenceOrder,
		stage.RouteTimeSlicePropertyTypes_instance,
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
		stage.RouteTimeSliceTypes,
		stage.RouteTimeSliceType_stagedOrder,
		stage.RouteTimeSliceTypes_reference,
		&stage.RouteTimeSliceTypes_referenceOrder,
		stage.RouteTimeSliceTypes_instance,
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
		stage.RouteTypes,
		stage.RouteType_stagedOrder,
		stage.RouteTypes_reference,
		&stage.RouteTypes_referenceOrder,
		stage.RouteTypes_instance,
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
		stage.RulesProceduresTimeSlicePropertyTypes,
		stage.RulesProceduresTimeSlicePropertyType_stagedOrder,
		stage.RulesProceduresTimeSlicePropertyTypes_reference,
		&stage.RulesProceduresTimeSlicePropertyTypes_referenceOrder,
		stage.RulesProceduresTimeSlicePropertyTypes_instance,
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
		stage.RulesProceduresTimeSliceTypes,
		stage.RulesProceduresTimeSliceType_stagedOrder,
		stage.RulesProceduresTimeSliceTypes_reference,
		&stage.RulesProceduresTimeSliceTypes_referenceOrder,
		stage.RulesProceduresTimeSliceTypes_instance,
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
		stage.RulesProceduresTypes,
		stage.RulesProceduresType_stagedOrder,
		stage.RulesProceduresTypes_reference,
		&stage.RulesProceduresTypes_referenceOrder,
		stage.RulesProceduresTypes_instance,
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
		stage.RunwayCentrelinePointPropertyTypes,
		stage.RunwayCentrelinePointPropertyType_stagedOrder,
		stage.RunwayCentrelinePointPropertyTypes_reference,
		&stage.RunwayCentrelinePointPropertyTypes_referenceOrder,
		stage.RunwayCentrelinePointPropertyTypes_instance,
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
		stage.StandardInstrumentArrivalPropertyTypes,
		stage.StandardInstrumentArrivalPropertyType_stagedOrder,
		stage.StandardInstrumentArrivalPropertyTypes_reference,
		&stage.StandardInstrumentArrivalPropertyTypes_referenceOrder,
		stage.StandardInstrumentArrivalPropertyTypes_instance,
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
		stage.StandardInstrumentDeparturePropertyTypes,
		stage.StandardInstrumentDeparturePropertyType_stagedOrder,
		stage.StandardInstrumentDeparturePropertyTypes_reference,
		&stage.StandardInstrumentDeparturePropertyTypes_referenceOrder,
		stage.StandardInstrumentDeparturePropertyTypes_instance,
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
		stage.TelephoneContactPropertyTypes,
		stage.TelephoneContactPropertyType_stagedOrder,
		stage.TelephoneContactPropertyTypes_reference,
		&stage.TelephoneContactPropertyTypes_referenceOrder,
		stage.TelephoneContactPropertyTypes_instance,
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
		stage.TelephoneContactTypes,
		stage.TelephoneContactType_stagedOrder,
		stage.TelephoneContactTypes_reference,
		&stage.TelephoneContactTypes_referenceOrder,
		stage.TelephoneContactTypes_instance,
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
	computeCommitsForType(
		stage,
		stage.TouchDownLiftOffPropertyTypes,
		stage.TouchDownLiftOffPropertyType_stagedOrder,
		stage.TouchDownLiftOffPropertyTypes_reference,
		&stage.TouchDownLiftOffPropertyTypes_referenceOrder,
		stage.TouchDownLiftOffPropertyTypes_instance,
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
		stage.VerticalStructurePropertyTypes,
		stage.VerticalStructurePropertyType_stagedOrder,
		stage.VerticalStructurePropertyTypes_reference,
		&stage.VerticalStructurePropertyTypes_referenceOrder,
		stage.VerticalStructurePropertyTypes_instance,
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

	__gong__computeReferencePass1(stage, stage.A_extension_15s, &stage.A_extension_15s_reference, &stage.A_extension_15s_referenceOrder, &stage.A_extension_15s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_16s, &stage.A_extension_16s_reference, &stage.A_extension_16s_referenceOrder, &stage.A_extension_16s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_17s, &stage.A_extension_17s_reference, &stage.A_extension_17s_referenceOrder, &stage.A_extension_17s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_18s, &stage.A_extension_18s_reference, &stage.A_extension_18s_referenceOrder, &stage.A_extension_18s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_19s, &stage.A_extension_19s_reference, &stage.A_extension_19s_referenceOrder, &stage.A_extension_19s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_2s, &stage.A_extension_2s_reference, &stage.A_extension_2s_referenceOrder, &stage.A_extension_2s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_20s, &stage.A_extension_20s_reference, &stage.A_extension_20s_referenceOrder, &stage.A_extension_20s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_21s, &stage.A_extension_21s_reference, &stage.A_extension_21s_referenceOrder, &stage.A_extension_21s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_22s, &stage.A_extension_22s_reference, &stage.A_extension_22s_referenceOrder, &stage.A_extension_22s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_23s, &stage.A_extension_23s_reference, &stage.A_extension_23s_referenceOrder, &stage.A_extension_23s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_24s, &stage.A_extension_24s_reference, &stage.A_extension_24s_referenceOrder, &stage.A_extension_24s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_25s, &stage.A_extension_25s_reference, &stage.A_extension_25s_referenceOrder, &stage.A_extension_25s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_26s, &stage.A_extension_26s_reference, &stage.A_extension_26s_referenceOrder, &stage.A_extension_26s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_27s, &stage.A_extension_27s_reference, &stage.A_extension_27s_referenceOrder, &stage.A_extension_27s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_28s, &stage.A_extension_28s_reference, &stage.A_extension_28s_referenceOrder, &stage.A_extension_28s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_29s, &stage.A_extension_29s_reference, &stage.A_extension_29s_referenceOrder, &stage.A_extension_29s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_3s, &stage.A_extension_3s_reference, &stage.A_extension_3s_referenceOrder, &stage.A_extension_3s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_30s, &stage.A_extension_30s_reference, &stage.A_extension_30s_referenceOrder, &stage.A_extension_30s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_31s, &stage.A_extension_31s_reference, &stage.A_extension_31s_referenceOrder, &stage.A_extension_31s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_32s, &stage.A_extension_32s_reference, &stage.A_extension_32s_referenceOrder, &stage.A_extension_32s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_33s, &stage.A_extension_33s_reference, &stage.A_extension_33s_referenceOrder, &stage.A_extension_33s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_34s, &stage.A_extension_34s_reference, &stage.A_extension_34s_referenceOrder, &stage.A_extension_34s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_35s, &stage.A_extension_35s_reference, &stage.A_extension_35s_referenceOrder, &stage.A_extension_35s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_36s, &stage.A_extension_36s_reference, &stage.A_extension_36s_referenceOrder, &stage.A_extension_36s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_37s, &stage.A_extension_37s_reference, &stage.A_extension_37s_referenceOrder, &stage.A_extension_37s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_38s, &stage.A_extension_38s_reference, &stage.A_extension_38s_referenceOrder, &stage.A_extension_38s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_39s, &stage.A_extension_39s_reference, &stage.A_extension_39s_referenceOrder, &stage.A_extension_39s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_4s, &stage.A_extension_4s_reference, &stage.A_extension_4s_referenceOrder, &stage.A_extension_4s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_40s, &stage.A_extension_40s_reference, &stage.A_extension_40s_referenceOrder, &stage.A_extension_40s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_41s, &stage.A_extension_41s_reference, &stage.A_extension_41s_referenceOrder, &stage.A_extension_41s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_42s, &stage.A_extension_42s_reference, &stage.A_extension_42s_referenceOrder, &stage.A_extension_42s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_5s, &stage.A_extension_5s_reference, &stage.A_extension_5s_referenceOrder, &stage.A_extension_5s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_6s, &stage.A_extension_6s_reference, &stage.A_extension_6s_referenceOrder, &stage.A_extension_6s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_7s, &stage.A_extension_7s_reference, &stage.A_extension_7s_referenceOrder, &stage.A_extension_7s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_8s, &stage.A_extension_8s_reference, &stage.A_extension_8s_referenceOrder, &stage.A_extension_8s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_9s, &stage.A_extension_9s_reference, &stage.A_extension_9s_referenceOrder, &stage.A_extension_9s_instance)

	__gong__computeReferencePass1(stage, stage.AbstractDirectFlightTypes, &stage.AbstractDirectFlightTypes_reference, &stage.AbstractDirectFlightTypes_referenceOrder, &stage.AbstractDirectFlightTypes_instance)

	__gong__computeReferencePass1(stage, stage.AbstractPropertiesWithScheduleTypes, &stage.AbstractPropertiesWithScheduleTypes_reference, &stage.AbstractPropertiesWithScheduleTypes_referenceOrder, &stage.AbstractPropertiesWithScheduleTypes_instance)

	__gong__computeReferencePass1(stage, stage.AbstractSegmentPointTypes, &stage.AbstractSegmentPointTypes_reference, &stage.AbstractSegmentPointTypes_referenceOrder, &stage.AbstractSegmentPointTypes_instance)

	__gong__computeReferencePass1(stage, stage.AerialRefuellingAnchorPropertyTypes, &stage.AerialRefuellingAnchorPropertyTypes_reference, &stage.AerialRefuellingAnchorPropertyTypes_referenceOrder, &stage.AerialRefuellingAnchorPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AerialRefuellingAnchorTypes, &stage.AerialRefuellingAnchorTypes_reference, &stage.AerialRefuellingAnchorTypes_referenceOrder, &stage.AerialRefuellingAnchorTypes_instance)

	__gong__computeReferencePass1(stage, stage.AerialRefuellingPointPropertyTypes, &stage.AerialRefuellingPointPropertyTypes_reference, &stage.AerialRefuellingPointPropertyTypes_referenceOrder, &stage.AerialRefuellingPointPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AerialRefuellingPointTypes, &stage.AerialRefuellingPointTypes_reference, &stage.AerialRefuellingPointTypes_referenceOrder, &stage.AerialRefuellingPointTypes_instance)

	__gong__computeReferencePass1(stage, stage.AerialRefuellingPropertyTypes, &stage.AerialRefuellingPropertyTypes_reference, &stage.AerialRefuellingPropertyTypes_referenceOrder, &stage.AerialRefuellingPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AerialRefuellingTimeSlicePropertyTypes, &stage.AerialRefuellingTimeSlicePropertyTypes_reference, &stage.AerialRefuellingTimeSlicePropertyTypes_referenceOrder, &stage.AerialRefuellingTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AerialRefuellingTimeSliceTypes, &stage.AerialRefuellingTimeSliceTypes_reference, &stage.AerialRefuellingTimeSliceTypes_referenceOrder, &stage.AerialRefuellingTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.AerialRefuellingTrackPropertyTypes, &stage.AerialRefuellingTrackPropertyTypes_reference, &stage.AerialRefuellingTrackPropertyTypes_referenceOrder, &stage.AerialRefuellingTrackPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AerialRefuellingTrackTypes, &stage.AerialRefuellingTrackTypes_reference, &stage.AerialRefuellingTrackTypes_referenceOrder, &stage.AerialRefuellingTrackTypes_instance)

	__gong__computeReferencePass1(stage, stage.AerialRefuellingTypes, &stage.AerialRefuellingTypes_reference, &stage.AerialRefuellingTypes_referenceOrder, &stage.AerialRefuellingTypes_instance)

	__gong__computeReferencePass1(stage, stage.AircraftCharacteristicPropertyTypes, &stage.AircraftCharacteristicPropertyTypes_reference, &stage.AircraftCharacteristicPropertyTypes_referenceOrder, &stage.AircraftCharacteristicPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AircraftCharacteristicTypes, &stage.AircraftCharacteristicTypes_reference, &stage.AircraftCharacteristicTypes_referenceOrder, &stage.AircraftCharacteristicTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirportHeliportPropertyTypes, &stage.AirportHeliportPropertyTypes_reference, &stage.AirportHeliportPropertyTypes_referenceOrder, &stage.AirportHeliportPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirspaceBorderCrossingPropertyTypes, &stage.AirspaceBorderCrossingPropertyTypes_reference, &stage.AirspaceBorderCrossingPropertyTypes_referenceOrder, &stage.AirspaceBorderCrossingPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirspaceBorderCrossingTimeSlicePropertyTypes, &stage.AirspaceBorderCrossingTimeSlicePropertyTypes_reference, &stage.AirspaceBorderCrossingTimeSlicePropertyTypes_referenceOrder, &stage.AirspaceBorderCrossingTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirspaceBorderCrossingTimeSliceTypes, &stage.AirspaceBorderCrossingTimeSliceTypes_reference, &stage.AirspaceBorderCrossingTimeSliceTypes_referenceOrder, &stage.AirspaceBorderCrossingTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirspaceBorderCrossingTypes, &stage.AirspaceBorderCrossingTypes_reference, &stage.AirspaceBorderCrossingTypes_referenceOrder, &stage.AirspaceBorderCrossingTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirspaceLayerPropertyTypes, &stage.AirspaceLayerPropertyTypes_reference, &stage.AirspaceLayerPropertyTypes_referenceOrder, &stage.AirspaceLayerPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirspaceLayerTypes, &stage.AirspaceLayerTypes_reference, &stage.AirspaceLayerTypes_referenceOrder, &stage.AirspaceLayerTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirspacePropertyTypes, &stage.AirspacePropertyTypes_reference, &stage.AirspacePropertyTypes_referenceOrder, &stage.AirspacePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AltitudeAdjustmentPropertyTypes, &stage.AltitudeAdjustmentPropertyTypes_reference, &stage.AltitudeAdjustmentPropertyTypes_referenceOrder, &stage.AltitudeAdjustmentPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AltitudeAdjustmentTypes, &stage.AltitudeAdjustmentTypes_reference, &stage.AltitudeAdjustmentTypes_referenceOrder, &stage.AltitudeAdjustmentTypes_instance)

	__gong__computeReferencePass1(stage, stage.AngleIndicationPropertyTypes, &stage.AngleIndicationPropertyTypes_reference, &stage.AngleIndicationPropertyTypes_referenceOrder, &stage.AngleIndicationPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AngleUsePropertyTypes, &stage.AngleUsePropertyTypes_reference, &stage.AngleUsePropertyTypes_referenceOrder, &stage.AngleUsePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AngleUseTypes, &stage.AngleUseTypes_reference, &stage.AngleUseTypes_referenceOrder, &stage.AngleUseTypes_instance)

	__gong__computeReferencePass1(stage, stage.AuthorityForAerialRefuellingPropertyTypes, &stage.AuthorityForAerialRefuellingPropertyTypes_reference, &stage.AuthorityForAerialRefuellingPropertyTypes_referenceOrder, &stage.AuthorityForAerialRefuellingPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AuthorityForAerialRefuellingTypes, &stage.AuthorityForAerialRefuellingTypes_reference, &stage.AuthorityForAerialRefuellingTypes_referenceOrder, &stage.AuthorityForAerialRefuellingTypes_instance)

	__gong__computeReferencePass1(stage, stage.ChangeOverPointTimeSlicePropertyTypes, &stage.ChangeOverPointTimeSlicePropertyTypes_reference, &stage.ChangeOverPointTimeSlicePropertyTypes_referenceOrder, &stage.ChangeOverPointTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ChangeOverPointTimeSliceTypes, &stage.ChangeOverPointTimeSliceTypes_reference, &stage.ChangeOverPointTimeSliceTypes_referenceOrder, &stage.ChangeOverPointTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.ChangeOverPointTypes, &stage.ChangeOverPointTypes_reference, &stage.ChangeOverPointTypes_referenceOrder, &stage.ChangeOverPointTypes_instance)

	__gong__computeReferencePass1(stage, stage.ContactInformationPropertyTypes, &stage.ContactInformationPropertyTypes_reference, &stage.ContactInformationPropertyTypes_referenceOrder, &stage.ContactInformationPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ContactInformationTypes, &stage.ContactInformationTypes_reference, &stage.ContactInformationTypes_referenceOrder, &stage.ContactInformationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CurvePropertyTypes, &stage.CurvePropertyTypes_reference, &stage.CurvePropertyTypes_referenceOrder, &stage.CurvePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.CurveTypes, &stage.CurveTypes_reference, &stage.CurveTypes_referenceOrder, &stage.CurveTypes_instance)

	__gong__computeReferencePass1(stage, stage.DMEPropertyTypes, &stage.DMEPropertyTypes_reference, &stage.DMEPropertyTypes_referenceOrder, &stage.DMEPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.DesignatedPointPropertyTypes, &stage.DesignatedPointPropertyTypes_reference, &stage.DesignatedPointPropertyTypes_referenceOrder, &stage.DesignatedPointPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.DirectFlightClassTypes, &stage.DirectFlightClassTypes_reference, &stage.DirectFlightClassTypes_referenceOrder, &stage.DirectFlightClassTypes_instance)

	__gong__computeReferencePass1(stage, stage.DirectFlightPropertyTypes, &stage.DirectFlightPropertyTypes_reference, &stage.DirectFlightPropertyTypes_referenceOrder, &stage.DirectFlightPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.DirectFlightSegmentPropertyTypes, &stage.DirectFlightSegmentPropertyTypes_reference, &stage.DirectFlightSegmentPropertyTypes_referenceOrder, &stage.DirectFlightSegmentPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.DirectFlightSegmentTypes, &stage.DirectFlightSegmentTypes_reference, &stage.DirectFlightSegmentTypes_referenceOrder, &stage.DirectFlightSegmentTypes_instance)

	__gong__computeReferencePass1(stage, stage.DistanceIndicationPropertyTypes, &stage.DistanceIndicationPropertyTypes_reference, &stage.DistanceIndicationPropertyTypes_referenceOrder, &stage.DistanceIndicationPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ElevatedCurveTypes, &stage.ElevatedCurveTypes_reference, &stage.ElevatedCurveTypes_referenceOrder, &stage.ElevatedCurveTypes_instance)

	__gong__computeReferencePass1(stage, stage.ElevatedPointTypes, &stage.ElevatedPointTypes_reference, &stage.ElevatedPointTypes_referenceOrder, &stage.ElevatedPointTypes_instance)

	__gong__computeReferencePass1(stage, stage.ElevatedSurfaceTypes, &stage.ElevatedSurfaceTypes_reference, &stage.ElevatedSurfaceTypes_referenceOrder, &stage.ElevatedSurfaceTypes_instance)

	__gong__computeReferencePass1(stage, stage.EnRouteSegmentPointPropertyTypes, &stage.EnRouteSegmentPointPropertyTypes_reference, &stage.EnRouteSegmentPointPropertyTypes_referenceOrder, &stage.EnRouteSegmentPointPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.EnRouteSegmentPointTypes, &stage.EnRouteSegmentPointTypes_reference, &stage.EnRouteSegmentPointTypes_referenceOrder, &stage.EnRouteSegmentPointTypes_instance)

	__gong__computeReferencePass1(stage, stage.FlightCharacteristicPropertyTypes, &stage.FlightCharacteristicPropertyTypes_reference, &stage.FlightCharacteristicPropertyTypes_referenceOrder, &stage.FlightCharacteristicPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.FlightCharacteristicTypes, &stage.FlightCharacteristicTypes_reference, &stage.FlightCharacteristicTypes_referenceOrder, &stage.FlightCharacteristicTypes_instance)

	__gong__computeReferencePass1(stage, stage.FlightConditionCircumstancePropertyTypes, &stage.FlightConditionCircumstancePropertyTypes_reference, &stage.FlightConditionCircumstancePropertyTypes_referenceOrder, &stage.FlightConditionCircumstancePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.FlightConditionCircumstanceTypes, &stage.FlightConditionCircumstanceTypes_reference, &stage.FlightConditionCircumstanceTypes_referenceOrder, &stage.FlightConditionCircumstanceTypes_instance)

	__gong__computeReferencePass1(stage, stage.FlightConditionCombinationPropertyTypes, &stage.FlightConditionCombinationPropertyTypes_reference, &stage.FlightConditionCombinationPropertyTypes_referenceOrder, &stage.FlightConditionCombinationPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.FlightConditionCombinationTypes, &stage.FlightConditionCombinationTypes_reference, &stage.FlightConditionCombinationTypes_referenceOrder, &stage.FlightConditionCombinationTypes_instance)

	__gong__computeReferencePass1(stage, stage.FlightConditionElementPropertyTypes, &stage.FlightConditionElementPropertyTypes_reference, &stage.FlightConditionElementPropertyTypes_referenceOrder, &stage.FlightConditionElementPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.FlightConditionElementTypes, &stage.FlightConditionElementTypes_reference, &stage.FlightConditionElementTypes_referenceOrder, &stage.FlightConditionElementTypes_instance)

	__gong__computeReferencePass1(stage, stage.FlightRestrictionLevelPropertyTypes, &stage.FlightRestrictionLevelPropertyTypes_reference, &stage.FlightRestrictionLevelPropertyTypes_referenceOrder, &stage.FlightRestrictionLevelPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.FlightRestrictionLevelTypes, &stage.FlightRestrictionLevelTypes_reference, &stage.FlightRestrictionLevelTypes_referenceOrder, &stage.FlightRestrictionLevelTypes_instance)

	__gong__computeReferencePass1(stage, stage.FlightRestrictionRoutePropertyTypes, &stage.FlightRestrictionRoutePropertyTypes_reference, &stage.FlightRestrictionRoutePropertyTypes_referenceOrder, &stage.FlightRestrictionRoutePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.FlightRestrictionRouteTypes, &stage.FlightRestrictionRouteTypes_reference, &stage.FlightRestrictionRouteTypes_referenceOrder, &stage.FlightRestrictionRouteTypes_instance)

	__gong__computeReferencePass1(stage, stage.FlightRestrictionTimeSlicePropertyTypes, &stage.FlightRestrictionTimeSlicePropertyTypes_reference, &stage.FlightRestrictionTimeSlicePropertyTypes_referenceOrder, &stage.FlightRestrictionTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.FlightRestrictionTimeSliceTypes, &stage.FlightRestrictionTimeSliceTypes_reference, &stage.FlightRestrictionTimeSliceTypes_referenceOrder, &stage.FlightRestrictionTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.FlightRestrictionTypes, &stage.FlightRestrictionTypes_reference, &stage.FlightRestrictionTypes_referenceOrder, &stage.FlightRestrictionTypes_instance)

	__gong__computeReferencePass1(stage, stage.FlightRoutingElementPropertyTypes, &stage.FlightRoutingElementPropertyTypes_reference, &stage.FlightRoutingElementPropertyTypes_referenceOrder, &stage.FlightRoutingElementPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.FlightRoutingElementTypes, &stage.FlightRoutingElementTypes_reference, &stage.FlightRoutingElementTypes_referenceOrder, &stage.FlightRoutingElementTypes_instance)

	__gong__computeReferencePass1(stage, stage.LinguisticNotePropertyTypes, &stage.LinguisticNotePropertyTypes_reference, &stage.LinguisticNotePropertyTypes_referenceOrder, &stage.LinguisticNotePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.LinguisticNoteTypes, &stage.LinguisticNoteTypes_reference, &stage.LinguisticNoteTypes_referenceOrder, &stage.LinguisticNoteTypes_instance)

	__gong__computeReferencePass1(stage, stage.MeteorologyPropertyTypes, &stage.MeteorologyPropertyTypes_reference, &stage.MeteorologyPropertyTypes_referenceOrder, &stage.MeteorologyPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.MeteorologyTypes, &stage.MeteorologyTypes_reference, &stage.MeteorologyTypes_referenceOrder, &stage.MeteorologyTypes_instance)

	__gong__computeReferencePass1(stage, stage.NavaidPropertyTypes, &stage.NavaidPropertyTypes_reference, &stage.NavaidPropertyTypes_referenceOrder, &stage.NavaidPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.NotePropertyTypes, &stage.NotePropertyTypes_reference, &stage.NotePropertyTypes_referenceOrder, &stage.NotePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.NoteTypes, &stage.NoteTypes_reference, &stage.NoteTypes_referenceOrder, &stage.NoteTypes_instance)

	__gong__computeReferencePass1(stage, stage.ObstacleAssessmentAreaPropertyTypes, &stage.ObstacleAssessmentAreaPropertyTypes_reference, &stage.ObstacleAssessmentAreaPropertyTypes_referenceOrder, &stage.ObstacleAssessmentAreaPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ObstacleAssessmentAreaTypes, &stage.ObstacleAssessmentAreaTypes_reference, &stage.ObstacleAssessmentAreaTypes_referenceOrder, &stage.ObstacleAssessmentAreaTypes_instance)

	__gong__computeReferencePass1(stage, stage.ObstaclePlacementPropertyTypes, &stage.ObstaclePlacementPropertyTypes_reference, &stage.ObstaclePlacementPropertyTypes_referenceOrder, &stage.ObstaclePlacementPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ObstaclePlacementTypes, &stage.ObstaclePlacementTypes_reference, &stage.ObstaclePlacementTypes_referenceOrder, &stage.ObstaclePlacementTypes_instance)

	__gong__computeReferencePass1(stage, stage.ObstructionPropertyTypes, &stage.ObstructionPropertyTypes_reference, &stage.ObstructionPropertyTypes_referenceOrder, &stage.ObstructionPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ObstructionTypes, &stage.ObstructionTypes_reference, &stage.ObstructionTypes_referenceOrder, &stage.ObstructionTypes_instance)

	__gong__computeReferencePass1(stage, stage.OnlineContactPropertyTypes, &stage.OnlineContactPropertyTypes_reference, &stage.OnlineContactPropertyTypes_referenceOrder, &stage.OnlineContactPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.OnlineContactTypes, &stage.OnlineContactTypes_reference, &stage.OnlineContactTypes_referenceOrder, &stage.OnlineContactTypes_instance)

	__gong__computeReferencePass1(stage, stage.OrganisationAuthorityPropertyTypes, &stage.OrganisationAuthorityPropertyTypes_reference, &stage.OrganisationAuthorityPropertyTypes_referenceOrder, &stage.OrganisationAuthorityPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.PointPropertyTypes, &stage.PointPropertyTypes_reference, &stage.PointPropertyTypes_referenceOrder, &stage.PointPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.PointReferencePropertyTypes, &stage.PointReferencePropertyTypes_reference, &stage.PointReferencePropertyTypes_referenceOrder, &stage.PointReferencePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.PointReferenceTypes, &stage.PointReferenceTypes_reference, &stage.PointReferenceTypes_referenceOrder, &stage.PointReferenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.PointTypes, &stage.PointTypes_reference, &stage.PointTypes_referenceOrder, &stage.PointTypes_instance)

	__gong__computeReferencePass1(stage, stage.PostalAddressPropertyTypes, &stage.PostalAddressPropertyTypes_reference, &stage.PostalAddressPropertyTypes_referenceOrder, &stage.PostalAddressPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.PostalAddressTypes, &stage.PostalAddressTypes_reference, &stage.PostalAddressTypes_referenceOrder, &stage.PostalAddressTypes_instance)

	__gong__computeReferencePass1(stage, stage.RadioFrequencyAreaPropertyTypes, &stage.RadioFrequencyAreaPropertyTypes_reference, &stage.RadioFrequencyAreaPropertyTypes_referenceOrder, &stage.RadioFrequencyAreaPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.RouteAvailabilityPropertyTypes, &stage.RouteAvailabilityPropertyTypes_reference, &stage.RouteAvailabilityPropertyTypes_referenceOrder, &stage.RouteAvailabilityPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.RouteAvailabilityTypes, &stage.RouteAvailabilityTypes_reference, &stage.RouteAvailabilityTypes_referenceOrder, &stage.RouteAvailabilityTypes_instance)

	__gong__computeReferencePass1(stage, stage.RouteDMETimeSlicePropertyTypes, &stage.RouteDMETimeSlicePropertyTypes_reference, &stage.RouteDMETimeSlicePropertyTypes_referenceOrder, &stage.RouteDMETimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.RouteDMETimeSliceTypes, &stage.RouteDMETimeSliceTypes_reference, &stage.RouteDMETimeSliceTypes_referenceOrder, &stage.RouteDMETimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.RouteDMETypes, &stage.RouteDMETypes_reference, &stage.RouteDMETypes_referenceOrder, &stage.RouteDMETypes_instance)

	__gong__computeReferencePass1(stage, stage.RoutePortionPropertyTypes, &stage.RoutePortionPropertyTypes_reference, &stage.RoutePortionPropertyTypes_referenceOrder, &stage.RoutePortionPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.RoutePortionTypes, &stage.RoutePortionTypes_reference, &stage.RoutePortionTypes_referenceOrder, &stage.RoutePortionTypes_instance)

	__gong__computeReferencePass1(stage, stage.RoutePropertyTypes, &stage.RoutePropertyTypes_reference, &stage.RoutePropertyTypes_referenceOrder, &stage.RoutePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.RouteSegmentTimeSlicePropertyTypes, &stage.RouteSegmentTimeSlicePropertyTypes_reference, &stage.RouteSegmentTimeSlicePropertyTypes_referenceOrder, &stage.RouteSegmentTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.RouteSegmentTimeSliceTypes, &stage.RouteSegmentTimeSliceTypes_reference, &stage.RouteSegmentTimeSliceTypes_referenceOrder, &stage.RouteSegmentTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.RouteSegmentTypes, &stage.RouteSegmentTypes_reference, &stage.RouteSegmentTypes_referenceOrder, &stage.RouteSegmentTypes_instance)

	__gong__computeReferencePass1(stage, stage.RouteTimeSlicePropertyTypes, &stage.RouteTimeSlicePropertyTypes_reference, &stage.RouteTimeSlicePropertyTypes_referenceOrder, &stage.RouteTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.RouteTimeSliceTypes, &stage.RouteTimeSliceTypes_reference, &stage.RouteTimeSliceTypes_referenceOrder, &stage.RouteTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.RouteTypes, &stage.RouteTypes_reference, &stage.RouteTypes_referenceOrder, &stage.RouteTypes_instance)

	__gong__computeReferencePass1(stage, stage.RulesProceduresTimeSlicePropertyTypes, &stage.RulesProceduresTimeSlicePropertyTypes_reference, &stage.RulesProceduresTimeSlicePropertyTypes_referenceOrder, &stage.RulesProceduresTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.RulesProceduresTimeSliceTypes, &stage.RulesProceduresTimeSliceTypes_reference, &stage.RulesProceduresTimeSliceTypes_referenceOrder, &stage.RulesProceduresTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.RulesProceduresTypes, &stage.RulesProceduresTypes_reference, &stage.RulesProceduresTypes_referenceOrder, &stage.RulesProceduresTypes_instance)

	__gong__computeReferencePass1(stage, stage.RunwayCentrelinePointPropertyTypes, &stage.RunwayCentrelinePointPropertyTypes_reference, &stage.RunwayCentrelinePointPropertyTypes_referenceOrder, &stage.RunwayCentrelinePointPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.StandardInstrumentArrivalPropertyTypes, &stage.StandardInstrumentArrivalPropertyTypes_reference, &stage.StandardInstrumentArrivalPropertyTypes_referenceOrder, &stage.StandardInstrumentArrivalPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.StandardInstrumentDeparturePropertyTypes, &stage.StandardInstrumentDeparturePropertyTypes_reference, &stage.StandardInstrumentDeparturePropertyTypes_referenceOrder, &stage.StandardInstrumentDeparturePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.StandardLevelColumnPropertyTypes, &stage.StandardLevelColumnPropertyTypes_reference, &stage.StandardLevelColumnPropertyTypes_referenceOrder, &stage.StandardLevelColumnPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.SurfacePropertyTypes, &stage.SurfacePropertyTypes_reference, &stage.SurfacePropertyTypes_referenceOrder, &stage.SurfacePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.SurfaceTypes, &stage.SurfaceTypes_reference, &stage.SurfaceTypes_referenceOrder, &stage.SurfaceTypes_instance)

	__gong__computeReferencePass1(stage, stage.TelephoneContactPropertyTypes, &stage.TelephoneContactPropertyTypes_reference, &stage.TelephoneContactPropertyTypes_referenceOrder, &stage.TelephoneContactPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.TelephoneContactTypes, &stage.TelephoneContactTypes_reference, &stage.TelephoneContactTypes_referenceOrder, &stage.TelephoneContactTypes_instance)

	__gong__computeReferencePass1(stage, stage.TimesheetPropertyTypes, &stage.TimesheetPropertyTypes_reference, &stage.TimesheetPropertyTypes_referenceOrder, &stage.TimesheetPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.TimesheetTypes, &stage.TimesheetTypes_reference, &stage.TimesheetTypes_referenceOrder, &stage.TimesheetTypes_instance)

	__gong__computeReferencePass1(stage, stage.TouchDownLiftOffPropertyTypes, &stage.TouchDownLiftOffPropertyTypes_reference, &stage.TouchDownLiftOffPropertyTypes_referenceOrder, &stage.TouchDownLiftOffPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.VerticalStructurePropertyTypes, &stage.VerticalStructurePropertyTypes_reference, &stage.VerticalStructurePropertyTypes_referenceOrder, &stage.VerticalStructurePropertyTypes_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.A_extensions, stage.A_extensions_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_1s, stage.A_extension_1s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_10s, stage.A_extension_10s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_11s, stage.A_extension_11s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_12s, stage.A_extension_12s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_13s, stage.A_extension_13s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_14s, stage.A_extension_14s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_15s, stage.A_extension_15s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_16s, stage.A_extension_16s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_17s, stage.A_extension_17s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_18s, stage.A_extension_18s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_19s, stage.A_extension_19s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_2s, stage.A_extension_2s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_20s, stage.A_extension_20s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_21s, stage.A_extension_21s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_22s, stage.A_extension_22s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_23s, stage.A_extension_23s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_24s, stage.A_extension_24s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_25s, stage.A_extension_25s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_26s, stage.A_extension_26s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_27s, stage.A_extension_27s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_28s, stage.A_extension_28s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_29s, stage.A_extension_29s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_3s, stage.A_extension_3s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_30s, stage.A_extension_30s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_31s, stage.A_extension_31s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_32s, stage.A_extension_32s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_33s, stage.A_extension_33s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_34s, stage.A_extension_34s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_35s, stage.A_extension_35s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_36s, stage.A_extension_36s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_37s, stage.A_extension_37s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_38s, stage.A_extension_38s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_39s, stage.A_extension_39s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_4s, stage.A_extension_4s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_40s, stage.A_extension_40s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_41s, stage.A_extension_41s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_42s, stage.A_extension_42s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_5s, stage.A_extension_5s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_6s, stage.A_extension_6s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_7s, stage.A_extension_7s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_8s, stage.A_extension_8s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_9s, stage.A_extension_9s_reference, stage)

	__gong__computeReferencePass2(stage.AbstractDirectFlightTypes, stage.AbstractDirectFlightTypes_reference, stage)

	__gong__computeReferencePass2(stage.AbstractPropertiesWithScheduleTypes, stage.AbstractPropertiesWithScheduleTypes_reference, stage)

	__gong__computeReferencePass2(stage.AbstractSegmentPointTypes, stage.AbstractSegmentPointTypes_reference, stage)

	__gong__computeReferencePass2(stage.AerialRefuellingAnchorPropertyTypes, stage.AerialRefuellingAnchorPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AerialRefuellingAnchorTypes, stage.AerialRefuellingAnchorTypes_reference, stage)

	__gong__computeReferencePass2(stage.AerialRefuellingPointPropertyTypes, stage.AerialRefuellingPointPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AerialRefuellingPointTypes, stage.AerialRefuellingPointTypes_reference, stage)

	__gong__computeReferencePass2(stage.AerialRefuellingPropertyTypes, stage.AerialRefuellingPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AerialRefuellingTimeSlicePropertyTypes, stage.AerialRefuellingTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AerialRefuellingTimeSliceTypes, stage.AerialRefuellingTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.AerialRefuellingTrackPropertyTypes, stage.AerialRefuellingTrackPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AerialRefuellingTrackTypes, stage.AerialRefuellingTrackTypes_reference, stage)

	__gong__computeReferencePass2(stage.AerialRefuellingTypes, stage.AerialRefuellingTypes_reference, stage)

	__gong__computeReferencePass2(stage.AircraftCharacteristicPropertyTypes, stage.AircraftCharacteristicPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AircraftCharacteristicTypes, stage.AircraftCharacteristicTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirportHeliportPropertyTypes, stage.AirportHeliportPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirspaceBorderCrossingPropertyTypes, stage.AirspaceBorderCrossingPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirspaceBorderCrossingTimeSlicePropertyTypes, stage.AirspaceBorderCrossingTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirspaceBorderCrossingTimeSliceTypes, stage.AirspaceBorderCrossingTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirspaceBorderCrossingTypes, stage.AirspaceBorderCrossingTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirspaceLayerPropertyTypes, stage.AirspaceLayerPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirspaceLayerTypes, stage.AirspaceLayerTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirspacePropertyTypes, stage.AirspacePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AltitudeAdjustmentPropertyTypes, stage.AltitudeAdjustmentPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AltitudeAdjustmentTypes, stage.AltitudeAdjustmentTypes_reference, stage)

	__gong__computeReferencePass2(stage.AngleIndicationPropertyTypes, stage.AngleIndicationPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AngleUsePropertyTypes, stage.AngleUsePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AngleUseTypes, stage.AngleUseTypes_reference, stage)

	__gong__computeReferencePass2(stage.AuthorityForAerialRefuellingPropertyTypes, stage.AuthorityForAerialRefuellingPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AuthorityForAerialRefuellingTypes, stage.AuthorityForAerialRefuellingTypes_reference, stage)

	__gong__computeReferencePass2(stage.ChangeOverPointTimeSlicePropertyTypes, stage.ChangeOverPointTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ChangeOverPointTimeSliceTypes, stage.ChangeOverPointTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.ChangeOverPointTypes, stage.ChangeOverPointTypes_reference, stage)

	__gong__computeReferencePass2(stage.ContactInformationPropertyTypes, stage.ContactInformationPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ContactInformationTypes, stage.ContactInformationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CurvePropertyTypes, stage.CurvePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.CurveTypes, stage.CurveTypes_reference, stage)

	__gong__computeReferencePass2(stage.DMEPropertyTypes, stage.DMEPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.DesignatedPointPropertyTypes, stage.DesignatedPointPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.DirectFlightClassTypes, stage.DirectFlightClassTypes_reference, stage)

	__gong__computeReferencePass2(stage.DirectFlightPropertyTypes, stage.DirectFlightPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.DirectFlightSegmentPropertyTypes, stage.DirectFlightSegmentPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.DirectFlightSegmentTypes, stage.DirectFlightSegmentTypes_reference, stage)

	__gong__computeReferencePass2(stage.DistanceIndicationPropertyTypes, stage.DistanceIndicationPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ElevatedCurveTypes, stage.ElevatedCurveTypes_reference, stage)

	__gong__computeReferencePass2(stage.ElevatedPointTypes, stage.ElevatedPointTypes_reference, stage)

	__gong__computeReferencePass2(stage.ElevatedSurfaceTypes, stage.ElevatedSurfaceTypes_reference, stage)

	__gong__computeReferencePass2(stage.EnRouteSegmentPointPropertyTypes, stage.EnRouteSegmentPointPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.EnRouteSegmentPointTypes, stage.EnRouteSegmentPointTypes_reference, stage)

	__gong__computeReferencePass2(stage.FlightCharacteristicPropertyTypes, stage.FlightCharacteristicPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.FlightCharacteristicTypes, stage.FlightCharacteristicTypes_reference, stage)

	__gong__computeReferencePass2(stage.FlightConditionCircumstancePropertyTypes, stage.FlightConditionCircumstancePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.FlightConditionCircumstanceTypes, stage.FlightConditionCircumstanceTypes_reference, stage)

	__gong__computeReferencePass2(stage.FlightConditionCombinationPropertyTypes, stage.FlightConditionCombinationPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.FlightConditionCombinationTypes, stage.FlightConditionCombinationTypes_reference, stage)

	__gong__computeReferencePass2(stage.FlightConditionElementPropertyTypes, stage.FlightConditionElementPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.FlightConditionElementTypes, stage.FlightConditionElementTypes_reference, stage)

	__gong__computeReferencePass2(stage.FlightRestrictionLevelPropertyTypes, stage.FlightRestrictionLevelPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.FlightRestrictionLevelTypes, stage.FlightRestrictionLevelTypes_reference, stage)

	__gong__computeReferencePass2(stage.FlightRestrictionRoutePropertyTypes, stage.FlightRestrictionRoutePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.FlightRestrictionRouteTypes, stage.FlightRestrictionRouteTypes_reference, stage)

	__gong__computeReferencePass2(stage.FlightRestrictionTimeSlicePropertyTypes, stage.FlightRestrictionTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.FlightRestrictionTimeSliceTypes, stage.FlightRestrictionTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.FlightRestrictionTypes, stage.FlightRestrictionTypes_reference, stage)

	__gong__computeReferencePass2(stage.FlightRoutingElementPropertyTypes, stage.FlightRoutingElementPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.FlightRoutingElementTypes, stage.FlightRoutingElementTypes_reference, stage)

	__gong__computeReferencePass2(stage.LinguisticNotePropertyTypes, stage.LinguisticNotePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.LinguisticNoteTypes, stage.LinguisticNoteTypes_reference, stage)

	__gong__computeReferencePass2(stage.MeteorologyPropertyTypes, stage.MeteorologyPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.MeteorologyTypes, stage.MeteorologyTypes_reference, stage)

	__gong__computeReferencePass2(stage.NavaidPropertyTypes, stage.NavaidPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.NotePropertyTypes, stage.NotePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.NoteTypes, stage.NoteTypes_reference, stage)

	__gong__computeReferencePass2(stage.ObstacleAssessmentAreaPropertyTypes, stage.ObstacleAssessmentAreaPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ObstacleAssessmentAreaTypes, stage.ObstacleAssessmentAreaTypes_reference, stage)

	__gong__computeReferencePass2(stage.ObstaclePlacementPropertyTypes, stage.ObstaclePlacementPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ObstaclePlacementTypes, stage.ObstaclePlacementTypes_reference, stage)

	__gong__computeReferencePass2(stage.ObstructionPropertyTypes, stage.ObstructionPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ObstructionTypes, stage.ObstructionTypes_reference, stage)

	__gong__computeReferencePass2(stage.OnlineContactPropertyTypes, stage.OnlineContactPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.OnlineContactTypes, stage.OnlineContactTypes_reference, stage)

	__gong__computeReferencePass2(stage.OrganisationAuthorityPropertyTypes, stage.OrganisationAuthorityPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.PointPropertyTypes, stage.PointPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.PointReferencePropertyTypes, stage.PointReferencePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.PointReferenceTypes, stage.PointReferenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.PointTypes, stage.PointTypes_reference, stage)

	__gong__computeReferencePass2(stage.PostalAddressPropertyTypes, stage.PostalAddressPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.PostalAddressTypes, stage.PostalAddressTypes_reference, stage)

	__gong__computeReferencePass2(stage.RadioFrequencyAreaPropertyTypes, stage.RadioFrequencyAreaPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.RouteAvailabilityPropertyTypes, stage.RouteAvailabilityPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.RouteAvailabilityTypes, stage.RouteAvailabilityTypes_reference, stage)

	__gong__computeReferencePass2(stage.RouteDMETimeSlicePropertyTypes, stage.RouteDMETimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.RouteDMETimeSliceTypes, stage.RouteDMETimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.RouteDMETypes, stage.RouteDMETypes_reference, stage)

	__gong__computeReferencePass2(stage.RoutePortionPropertyTypes, stage.RoutePortionPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.RoutePortionTypes, stage.RoutePortionTypes_reference, stage)

	__gong__computeReferencePass2(stage.RoutePropertyTypes, stage.RoutePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.RouteSegmentTimeSlicePropertyTypes, stage.RouteSegmentTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.RouteSegmentTimeSliceTypes, stage.RouteSegmentTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.RouteSegmentTypes, stage.RouteSegmentTypes_reference, stage)

	__gong__computeReferencePass2(stage.RouteTimeSlicePropertyTypes, stage.RouteTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.RouteTimeSliceTypes, stage.RouteTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.RouteTypes, stage.RouteTypes_reference, stage)

	__gong__computeReferencePass2(stage.RulesProceduresTimeSlicePropertyTypes, stage.RulesProceduresTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.RulesProceduresTimeSliceTypes, stage.RulesProceduresTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.RulesProceduresTypes, stage.RulesProceduresTypes_reference, stage)

	__gong__computeReferencePass2(stage.RunwayCentrelinePointPropertyTypes, stage.RunwayCentrelinePointPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.StandardInstrumentArrivalPropertyTypes, stage.StandardInstrumentArrivalPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.StandardInstrumentDeparturePropertyTypes, stage.StandardInstrumentDeparturePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.StandardLevelColumnPropertyTypes, stage.StandardLevelColumnPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.SurfacePropertyTypes, stage.SurfacePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.SurfaceTypes, stage.SurfaceTypes_reference, stage)

	__gong__computeReferencePass2(stage.TelephoneContactPropertyTypes, stage.TelephoneContactPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.TelephoneContactTypes, stage.TelephoneContactTypes_reference, stage)

	__gong__computeReferencePass2(stage.TimesheetPropertyTypes, stage.TimesheetPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.TimesheetTypes, stage.TimesheetTypes_reference, stage)

	__gong__computeReferencePass2(stage.TouchDownLiftOffPropertyTypes, stage.TouchDownLiftOffPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.VerticalStructurePropertyTypes, stage.VerticalStructurePropertyTypes_reference, stage)

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

func (a_extension_15 *A_extension_15) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_15_stagedOrder, stage.A_extension_15s_referenceOrder, a_extension_15, "A_extension_15")
}

func (a_extension_16 *A_extension_16) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_16_stagedOrder, stage.A_extension_16s_referenceOrder, a_extension_16, "A_extension_16")
}

func (a_extension_17 *A_extension_17) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_17_stagedOrder, stage.A_extension_17s_referenceOrder, a_extension_17, "A_extension_17")
}

func (a_extension_18 *A_extension_18) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_18_stagedOrder, stage.A_extension_18s_referenceOrder, a_extension_18, "A_extension_18")
}

func (a_extension_19 *A_extension_19) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_19_stagedOrder, stage.A_extension_19s_referenceOrder, a_extension_19, "A_extension_19")
}

func (a_extension_2 *A_extension_2) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_2_stagedOrder, stage.A_extension_2s_referenceOrder, a_extension_2, "A_extension_2")
}

func (a_extension_20 *A_extension_20) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_20_stagedOrder, stage.A_extension_20s_referenceOrder, a_extension_20, "A_extension_20")
}

func (a_extension_21 *A_extension_21) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_21_stagedOrder, stage.A_extension_21s_referenceOrder, a_extension_21, "A_extension_21")
}

func (a_extension_22 *A_extension_22) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_22_stagedOrder, stage.A_extension_22s_referenceOrder, a_extension_22, "A_extension_22")
}

func (a_extension_23 *A_extension_23) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_23_stagedOrder, stage.A_extension_23s_referenceOrder, a_extension_23, "A_extension_23")
}

func (a_extension_24 *A_extension_24) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_24_stagedOrder, stage.A_extension_24s_referenceOrder, a_extension_24, "A_extension_24")
}

func (a_extension_25 *A_extension_25) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_25_stagedOrder, stage.A_extension_25s_referenceOrder, a_extension_25, "A_extension_25")
}

func (a_extension_26 *A_extension_26) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_26_stagedOrder, stage.A_extension_26s_referenceOrder, a_extension_26, "A_extension_26")
}

func (a_extension_27 *A_extension_27) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_27_stagedOrder, stage.A_extension_27s_referenceOrder, a_extension_27, "A_extension_27")
}

func (a_extension_28 *A_extension_28) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_28_stagedOrder, stage.A_extension_28s_referenceOrder, a_extension_28, "A_extension_28")
}

func (a_extension_29 *A_extension_29) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_29_stagedOrder, stage.A_extension_29s_referenceOrder, a_extension_29, "A_extension_29")
}

func (a_extension_3 *A_extension_3) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_3_stagedOrder, stage.A_extension_3s_referenceOrder, a_extension_3, "A_extension_3")
}

func (a_extension_30 *A_extension_30) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_30_stagedOrder, stage.A_extension_30s_referenceOrder, a_extension_30, "A_extension_30")
}

func (a_extension_31 *A_extension_31) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_31_stagedOrder, stage.A_extension_31s_referenceOrder, a_extension_31, "A_extension_31")
}

func (a_extension_32 *A_extension_32) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_32_stagedOrder, stage.A_extension_32s_referenceOrder, a_extension_32, "A_extension_32")
}

func (a_extension_33 *A_extension_33) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_33_stagedOrder, stage.A_extension_33s_referenceOrder, a_extension_33, "A_extension_33")
}

func (a_extension_34 *A_extension_34) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_34_stagedOrder, stage.A_extension_34s_referenceOrder, a_extension_34, "A_extension_34")
}

func (a_extension_35 *A_extension_35) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_35_stagedOrder, stage.A_extension_35s_referenceOrder, a_extension_35, "A_extension_35")
}

func (a_extension_36 *A_extension_36) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_36_stagedOrder, stage.A_extension_36s_referenceOrder, a_extension_36, "A_extension_36")
}

func (a_extension_37 *A_extension_37) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_37_stagedOrder, stage.A_extension_37s_referenceOrder, a_extension_37, "A_extension_37")
}

func (a_extension_38 *A_extension_38) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_38_stagedOrder, stage.A_extension_38s_referenceOrder, a_extension_38, "A_extension_38")
}

func (a_extension_39 *A_extension_39) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_39_stagedOrder, stage.A_extension_39s_referenceOrder, a_extension_39, "A_extension_39")
}

func (a_extension_4 *A_extension_4) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_4_stagedOrder, stage.A_extension_4s_referenceOrder, a_extension_4, "A_extension_4")
}

func (a_extension_40 *A_extension_40) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_40_stagedOrder, stage.A_extension_40s_referenceOrder, a_extension_40, "A_extension_40")
}

func (a_extension_41 *A_extension_41) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_41_stagedOrder, stage.A_extension_41s_referenceOrder, a_extension_41, "A_extension_41")
}

func (a_extension_42 *A_extension_42) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_extension_42_stagedOrder, stage.A_extension_42s_referenceOrder, a_extension_42, "A_extension_42")
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

func (abstractdirectflighttype *AbstractDirectFlightType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractDirectFlightType_stagedOrder, stage.AbstractDirectFlightTypes_referenceOrder, abstractdirectflighttype, "AbstractDirectFlightType")
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractPropertiesWithScheduleType_stagedOrder, stage.AbstractPropertiesWithScheduleTypes_referenceOrder, abstractpropertieswithscheduletype, "AbstractPropertiesWithScheduleType")
}

func (abstractsegmentpointtype *AbstractSegmentPointType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractSegmentPointType_stagedOrder, stage.AbstractSegmentPointTypes_referenceOrder, abstractsegmentpointtype, "AbstractSegmentPointType")
}

func (aerialrefuellinganchorpropertytype *AerialRefuellingAnchorPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AerialRefuellingAnchorPropertyType_stagedOrder, stage.AerialRefuellingAnchorPropertyTypes_referenceOrder, aerialrefuellinganchorpropertytype, "AerialRefuellingAnchorPropertyType")
}

func (aerialrefuellinganchortype *AerialRefuellingAnchorType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AerialRefuellingAnchorType_stagedOrder, stage.AerialRefuellingAnchorTypes_referenceOrder, aerialrefuellinganchortype, "AerialRefuellingAnchorType")
}

func (aerialrefuellingpointpropertytype *AerialRefuellingPointPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AerialRefuellingPointPropertyType_stagedOrder, stage.AerialRefuellingPointPropertyTypes_referenceOrder, aerialrefuellingpointpropertytype, "AerialRefuellingPointPropertyType")
}

func (aerialrefuellingpointtype *AerialRefuellingPointType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AerialRefuellingPointType_stagedOrder, stage.AerialRefuellingPointTypes_referenceOrder, aerialrefuellingpointtype, "AerialRefuellingPointType")
}

func (aerialrefuellingpropertytype *AerialRefuellingPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AerialRefuellingPropertyType_stagedOrder, stage.AerialRefuellingPropertyTypes_referenceOrder, aerialrefuellingpropertytype, "AerialRefuellingPropertyType")
}

func (aerialrefuellingtimeslicepropertytype *AerialRefuellingTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AerialRefuellingTimeSlicePropertyType_stagedOrder, stage.AerialRefuellingTimeSlicePropertyTypes_referenceOrder, aerialrefuellingtimeslicepropertytype, "AerialRefuellingTimeSlicePropertyType")
}

func (aerialrefuellingtimeslicetype *AerialRefuellingTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AerialRefuellingTimeSliceType_stagedOrder, stage.AerialRefuellingTimeSliceTypes_referenceOrder, aerialrefuellingtimeslicetype, "AerialRefuellingTimeSliceType")
}

func (aerialrefuellingtrackpropertytype *AerialRefuellingTrackPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AerialRefuellingTrackPropertyType_stagedOrder, stage.AerialRefuellingTrackPropertyTypes_referenceOrder, aerialrefuellingtrackpropertytype, "AerialRefuellingTrackPropertyType")
}

func (aerialrefuellingtracktype *AerialRefuellingTrackType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AerialRefuellingTrackType_stagedOrder, stage.AerialRefuellingTrackTypes_referenceOrder, aerialrefuellingtracktype, "AerialRefuellingTrackType")
}

func (aerialrefuellingtype *AerialRefuellingType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AerialRefuellingType_stagedOrder, stage.AerialRefuellingTypes_referenceOrder, aerialrefuellingtype, "AerialRefuellingType")
}

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AircraftCharacteristicPropertyType_stagedOrder, stage.AircraftCharacteristicPropertyTypes_referenceOrder, aircraftcharacteristicpropertytype, "AircraftCharacteristicPropertyType")
}

func (aircraftcharacteristictype *AircraftCharacteristicType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AircraftCharacteristicType_stagedOrder, stage.AircraftCharacteristicTypes_referenceOrder, aircraftcharacteristictype, "AircraftCharacteristicType")
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirportHeliportPropertyType_stagedOrder, stage.AirportHeliportPropertyTypes_referenceOrder, airportheliportpropertytype, "AirportHeliportPropertyType")
}

func (airspacebordercrossingpropertytype *AirspaceBorderCrossingPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirspaceBorderCrossingPropertyType_stagedOrder, stage.AirspaceBorderCrossingPropertyTypes_referenceOrder, airspacebordercrossingpropertytype, "AirspaceBorderCrossingPropertyType")
}

func (airspacebordercrossingtimeslicepropertytype *AirspaceBorderCrossingTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirspaceBorderCrossingTimeSlicePropertyType_stagedOrder, stage.AirspaceBorderCrossingTimeSlicePropertyTypes_referenceOrder, airspacebordercrossingtimeslicepropertytype, "AirspaceBorderCrossingTimeSlicePropertyType")
}

func (airspacebordercrossingtimeslicetype *AirspaceBorderCrossingTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirspaceBorderCrossingTimeSliceType_stagedOrder, stage.AirspaceBorderCrossingTimeSliceTypes_referenceOrder, airspacebordercrossingtimeslicetype, "AirspaceBorderCrossingTimeSliceType")
}

func (airspacebordercrossingtype *AirspaceBorderCrossingType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirspaceBorderCrossingType_stagedOrder, stage.AirspaceBorderCrossingTypes_referenceOrder, airspacebordercrossingtype, "AirspaceBorderCrossingType")
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

func (altitudeadjustmentpropertytype *AltitudeAdjustmentPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AltitudeAdjustmentPropertyType_stagedOrder, stage.AltitudeAdjustmentPropertyTypes_referenceOrder, altitudeadjustmentpropertytype, "AltitudeAdjustmentPropertyType")
}

func (altitudeadjustmenttype *AltitudeAdjustmentType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AltitudeAdjustmentType_stagedOrder, stage.AltitudeAdjustmentTypes_referenceOrder, altitudeadjustmenttype, "AltitudeAdjustmentType")
}

func (angleindicationpropertytype *AngleIndicationPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AngleIndicationPropertyType_stagedOrder, stage.AngleIndicationPropertyTypes_referenceOrder, angleindicationpropertytype, "AngleIndicationPropertyType")
}

func (angleusepropertytype *AngleUsePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AngleUsePropertyType_stagedOrder, stage.AngleUsePropertyTypes_referenceOrder, angleusepropertytype, "AngleUsePropertyType")
}

func (angleusetype *AngleUseType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AngleUseType_stagedOrder, stage.AngleUseTypes_referenceOrder, angleusetype, "AngleUseType")
}

func (authorityforaerialrefuellingpropertytype *AuthorityForAerialRefuellingPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AuthorityForAerialRefuellingPropertyType_stagedOrder, stage.AuthorityForAerialRefuellingPropertyTypes_referenceOrder, authorityforaerialrefuellingpropertytype, "AuthorityForAerialRefuellingPropertyType")
}

func (authorityforaerialrefuellingtype *AuthorityForAerialRefuellingType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AuthorityForAerialRefuellingType_stagedOrder, stage.AuthorityForAerialRefuellingTypes_referenceOrder, authorityforaerialrefuellingtype, "AuthorityForAerialRefuellingType")
}

func (changeoverpointtimeslicepropertytype *ChangeOverPointTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ChangeOverPointTimeSlicePropertyType_stagedOrder, stage.ChangeOverPointTimeSlicePropertyTypes_referenceOrder, changeoverpointtimeslicepropertytype, "ChangeOverPointTimeSlicePropertyType")
}

func (changeoverpointtimeslicetype *ChangeOverPointTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ChangeOverPointTimeSliceType_stagedOrder, stage.ChangeOverPointTimeSliceTypes_referenceOrder, changeoverpointtimeslicetype, "ChangeOverPointTimeSliceType")
}

func (changeoverpointtype *ChangeOverPointType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ChangeOverPointType_stagedOrder, stage.ChangeOverPointTypes_referenceOrder, changeoverpointtype, "ChangeOverPointType")
}

func (contactinformationpropertytype *ContactInformationPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ContactInformationPropertyType_stagedOrder, stage.ContactInformationPropertyTypes_referenceOrder, contactinformationpropertytype, "ContactInformationPropertyType")
}

func (contactinformationtype *ContactInformationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ContactInformationType_stagedOrder, stage.ContactInformationTypes_referenceOrder, contactinformationtype, "ContactInformationType")
}

func (curvepropertytype *CurvePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CurvePropertyType_stagedOrder, stage.CurvePropertyTypes_referenceOrder, curvepropertytype, "CurvePropertyType")
}

func (curvetype *CurveType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CurveType_stagedOrder, stage.CurveTypes_referenceOrder, curvetype, "CurveType")
}

func (dmepropertytype *DMEPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DMEPropertyType_stagedOrder, stage.DMEPropertyTypes_referenceOrder, dmepropertytype, "DMEPropertyType")
}

func (designatedpointpropertytype *DesignatedPointPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DesignatedPointPropertyType_stagedOrder, stage.DesignatedPointPropertyTypes_referenceOrder, designatedpointpropertytype, "DesignatedPointPropertyType")
}

func (directflightclasstype *DirectFlightClassType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DirectFlightClassType_stagedOrder, stage.DirectFlightClassTypes_referenceOrder, directflightclasstype, "DirectFlightClassType")
}

func (directflightpropertytype *DirectFlightPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DirectFlightPropertyType_stagedOrder, stage.DirectFlightPropertyTypes_referenceOrder, directflightpropertytype, "DirectFlightPropertyType")
}

func (directflightsegmentpropertytype *DirectFlightSegmentPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DirectFlightSegmentPropertyType_stagedOrder, stage.DirectFlightSegmentPropertyTypes_referenceOrder, directflightsegmentpropertytype, "DirectFlightSegmentPropertyType")
}

func (directflightsegmenttype *DirectFlightSegmentType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DirectFlightSegmentType_stagedOrder, stage.DirectFlightSegmentTypes_referenceOrder, directflightsegmenttype, "DirectFlightSegmentType")
}

func (distanceindicationpropertytype *DistanceIndicationPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DistanceIndicationPropertyType_stagedOrder, stage.DistanceIndicationPropertyTypes_referenceOrder, distanceindicationpropertytype, "DistanceIndicationPropertyType")
}

func (elevatedcurvetype *ElevatedCurveType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ElevatedCurveType_stagedOrder, stage.ElevatedCurveTypes_referenceOrder, elevatedcurvetype, "ElevatedCurveType")
}

func (elevatedpointtype *ElevatedPointType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ElevatedPointType_stagedOrder, stage.ElevatedPointTypes_referenceOrder, elevatedpointtype, "ElevatedPointType")
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ElevatedSurfaceType_stagedOrder, stage.ElevatedSurfaceTypes_referenceOrder, elevatedsurfacetype, "ElevatedSurfaceType")
}

func (enroutesegmentpointpropertytype *EnRouteSegmentPointPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.EnRouteSegmentPointPropertyType_stagedOrder, stage.EnRouteSegmentPointPropertyTypes_referenceOrder, enroutesegmentpointpropertytype, "EnRouteSegmentPointPropertyType")
}

func (enroutesegmentpointtype *EnRouteSegmentPointType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.EnRouteSegmentPointType_stagedOrder, stage.EnRouteSegmentPointTypes_referenceOrder, enroutesegmentpointtype, "EnRouteSegmentPointType")
}

func (flightcharacteristicpropertytype *FlightCharacteristicPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FlightCharacteristicPropertyType_stagedOrder, stage.FlightCharacteristicPropertyTypes_referenceOrder, flightcharacteristicpropertytype, "FlightCharacteristicPropertyType")
}

func (flightcharacteristictype *FlightCharacteristicType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FlightCharacteristicType_stagedOrder, stage.FlightCharacteristicTypes_referenceOrder, flightcharacteristictype, "FlightCharacteristicType")
}

func (flightconditioncircumstancepropertytype *FlightConditionCircumstancePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FlightConditionCircumstancePropertyType_stagedOrder, stage.FlightConditionCircumstancePropertyTypes_referenceOrder, flightconditioncircumstancepropertytype, "FlightConditionCircumstancePropertyType")
}

func (flightconditioncircumstancetype *FlightConditionCircumstanceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FlightConditionCircumstanceType_stagedOrder, stage.FlightConditionCircumstanceTypes_referenceOrder, flightconditioncircumstancetype, "FlightConditionCircumstanceType")
}

func (flightconditioncombinationpropertytype *FlightConditionCombinationPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FlightConditionCombinationPropertyType_stagedOrder, stage.FlightConditionCombinationPropertyTypes_referenceOrder, flightconditioncombinationpropertytype, "FlightConditionCombinationPropertyType")
}

func (flightconditioncombinationtype *FlightConditionCombinationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FlightConditionCombinationType_stagedOrder, stage.FlightConditionCombinationTypes_referenceOrder, flightconditioncombinationtype, "FlightConditionCombinationType")
}

func (flightconditionelementpropertytype *FlightConditionElementPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FlightConditionElementPropertyType_stagedOrder, stage.FlightConditionElementPropertyTypes_referenceOrder, flightconditionelementpropertytype, "FlightConditionElementPropertyType")
}

func (flightconditionelementtype *FlightConditionElementType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FlightConditionElementType_stagedOrder, stage.FlightConditionElementTypes_referenceOrder, flightconditionelementtype, "FlightConditionElementType")
}

func (flightrestrictionlevelpropertytype *FlightRestrictionLevelPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FlightRestrictionLevelPropertyType_stagedOrder, stage.FlightRestrictionLevelPropertyTypes_referenceOrder, flightrestrictionlevelpropertytype, "FlightRestrictionLevelPropertyType")
}

func (flightrestrictionleveltype *FlightRestrictionLevelType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FlightRestrictionLevelType_stagedOrder, stage.FlightRestrictionLevelTypes_referenceOrder, flightrestrictionleveltype, "FlightRestrictionLevelType")
}

func (flightrestrictionroutepropertytype *FlightRestrictionRoutePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FlightRestrictionRoutePropertyType_stagedOrder, stage.FlightRestrictionRoutePropertyTypes_referenceOrder, flightrestrictionroutepropertytype, "FlightRestrictionRoutePropertyType")
}

func (flightrestrictionroutetype *FlightRestrictionRouteType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FlightRestrictionRouteType_stagedOrder, stage.FlightRestrictionRouteTypes_referenceOrder, flightrestrictionroutetype, "FlightRestrictionRouteType")
}

func (flightrestrictiontimeslicepropertytype *FlightRestrictionTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FlightRestrictionTimeSlicePropertyType_stagedOrder, stage.FlightRestrictionTimeSlicePropertyTypes_referenceOrder, flightrestrictiontimeslicepropertytype, "FlightRestrictionTimeSlicePropertyType")
}

func (flightrestrictiontimeslicetype *FlightRestrictionTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FlightRestrictionTimeSliceType_stagedOrder, stage.FlightRestrictionTimeSliceTypes_referenceOrder, flightrestrictiontimeslicetype, "FlightRestrictionTimeSliceType")
}

func (flightrestrictiontype *FlightRestrictionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FlightRestrictionType_stagedOrder, stage.FlightRestrictionTypes_referenceOrder, flightrestrictiontype, "FlightRestrictionType")
}

func (flightroutingelementpropertytype *FlightRoutingElementPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FlightRoutingElementPropertyType_stagedOrder, stage.FlightRoutingElementPropertyTypes_referenceOrder, flightroutingelementpropertytype, "FlightRoutingElementPropertyType")
}

func (flightroutingelementtype *FlightRoutingElementType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FlightRoutingElementType_stagedOrder, stage.FlightRoutingElementTypes_referenceOrder, flightroutingelementtype, "FlightRoutingElementType")
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LinguisticNotePropertyType_stagedOrder, stage.LinguisticNotePropertyTypes_referenceOrder, linguisticnotepropertytype, "LinguisticNotePropertyType")
}

func (linguisticnotetype *LinguisticNoteType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LinguisticNoteType_stagedOrder, stage.LinguisticNoteTypes_referenceOrder, linguisticnotetype, "LinguisticNoteType")
}

func (meteorologypropertytype *MeteorologyPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MeteorologyPropertyType_stagedOrder, stage.MeteorologyPropertyTypes_referenceOrder, meteorologypropertytype, "MeteorologyPropertyType")
}

func (meteorologytype *MeteorologyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MeteorologyType_stagedOrder, stage.MeteorologyTypes_referenceOrder, meteorologytype, "MeteorologyType")
}

func (navaidpropertytype *NavaidPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NavaidPropertyType_stagedOrder, stage.NavaidPropertyTypes_referenceOrder, navaidpropertytype, "NavaidPropertyType")
}

func (notepropertytype *NotePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NotePropertyType_stagedOrder, stage.NotePropertyTypes_referenceOrder, notepropertytype, "NotePropertyType")
}

func (notetype *NoteType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoteType_stagedOrder, stage.NoteTypes_referenceOrder, notetype, "NoteType")
}

func (obstacleassessmentareapropertytype *ObstacleAssessmentAreaPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ObstacleAssessmentAreaPropertyType_stagedOrder, stage.ObstacleAssessmentAreaPropertyTypes_referenceOrder, obstacleassessmentareapropertytype, "ObstacleAssessmentAreaPropertyType")
}

func (obstacleassessmentareatype *ObstacleAssessmentAreaType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ObstacleAssessmentAreaType_stagedOrder, stage.ObstacleAssessmentAreaTypes_referenceOrder, obstacleassessmentareatype, "ObstacleAssessmentAreaType")
}

func (obstacleplacementpropertytype *ObstaclePlacementPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ObstaclePlacementPropertyType_stagedOrder, stage.ObstaclePlacementPropertyTypes_referenceOrder, obstacleplacementpropertytype, "ObstaclePlacementPropertyType")
}

func (obstacleplacementtype *ObstaclePlacementType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ObstaclePlacementType_stagedOrder, stage.ObstaclePlacementTypes_referenceOrder, obstacleplacementtype, "ObstaclePlacementType")
}

func (obstructionpropertytype *ObstructionPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ObstructionPropertyType_stagedOrder, stage.ObstructionPropertyTypes_referenceOrder, obstructionpropertytype, "ObstructionPropertyType")
}

func (obstructiontype *ObstructionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ObstructionType_stagedOrder, stage.ObstructionTypes_referenceOrder, obstructiontype, "ObstructionType")
}

func (onlinecontactpropertytype *OnlineContactPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.OnlineContactPropertyType_stagedOrder, stage.OnlineContactPropertyTypes_referenceOrder, onlinecontactpropertytype, "OnlineContactPropertyType")
}

func (onlinecontacttype *OnlineContactType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.OnlineContactType_stagedOrder, stage.OnlineContactTypes_referenceOrder, onlinecontacttype, "OnlineContactType")
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.OrganisationAuthorityPropertyType_stagedOrder, stage.OrganisationAuthorityPropertyTypes_referenceOrder, organisationauthoritypropertytype, "OrganisationAuthorityPropertyType")
}

func (pointpropertytype *PointPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PointPropertyType_stagedOrder, stage.PointPropertyTypes_referenceOrder, pointpropertytype, "PointPropertyType")
}

func (pointreferencepropertytype *PointReferencePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PointReferencePropertyType_stagedOrder, stage.PointReferencePropertyTypes_referenceOrder, pointreferencepropertytype, "PointReferencePropertyType")
}

func (pointreferencetype *PointReferenceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PointReferenceType_stagedOrder, stage.PointReferenceTypes_referenceOrder, pointreferencetype, "PointReferenceType")
}

func (pointtype *PointType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PointType_stagedOrder, stage.PointTypes_referenceOrder, pointtype, "PointType")
}

func (postaladdresspropertytype *PostalAddressPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PostalAddressPropertyType_stagedOrder, stage.PostalAddressPropertyTypes_referenceOrder, postaladdresspropertytype, "PostalAddressPropertyType")
}

func (postaladdresstype *PostalAddressType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PostalAddressType_stagedOrder, stage.PostalAddressTypes_referenceOrder, postaladdresstype, "PostalAddressType")
}

func (radiofrequencyareapropertytype *RadioFrequencyAreaPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RadioFrequencyAreaPropertyType_stagedOrder, stage.RadioFrequencyAreaPropertyTypes_referenceOrder, radiofrequencyareapropertytype, "RadioFrequencyAreaPropertyType")
}

func (routeavailabilitypropertytype *RouteAvailabilityPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RouteAvailabilityPropertyType_stagedOrder, stage.RouteAvailabilityPropertyTypes_referenceOrder, routeavailabilitypropertytype, "RouteAvailabilityPropertyType")
}

func (routeavailabilitytype *RouteAvailabilityType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RouteAvailabilityType_stagedOrder, stage.RouteAvailabilityTypes_referenceOrder, routeavailabilitytype, "RouteAvailabilityType")
}

func (routedmetimeslicepropertytype *RouteDMETimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RouteDMETimeSlicePropertyType_stagedOrder, stage.RouteDMETimeSlicePropertyTypes_referenceOrder, routedmetimeslicepropertytype, "RouteDMETimeSlicePropertyType")
}

func (routedmetimeslicetype *RouteDMETimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RouteDMETimeSliceType_stagedOrder, stage.RouteDMETimeSliceTypes_referenceOrder, routedmetimeslicetype, "RouteDMETimeSliceType")
}

func (routedmetype *RouteDMEType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RouteDMEType_stagedOrder, stage.RouteDMETypes_referenceOrder, routedmetype, "RouteDMEType")
}

func (routeportionpropertytype *RoutePortionPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RoutePortionPropertyType_stagedOrder, stage.RoutePortionPropertyTypes_referenceOrder, routeportionpropertytype, "RoutePortionPropertyType")
}

func (routeportiontype *RoutePortionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RoutePortionType_stagedOrder, stage.RoutePortionTypes_referenceOrder, routeportiontype, "RoutePortionType")
}

func (routepropertytype *RoutePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RoutePropertyType_stagedOrder, stage.RoutePropertyTypes_referenceOrder, routepropertytype, "RoutePropertyType")
}

func (routesegmenttimeslicepropertytype *RouteSegmentTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RouteSegmentTimeSlicePropertyType_stagedOrder, stage.RouteSegmentTimeSlicePropertyTypes_referenceOrder, routesegmenttimeslicepropertytype, "RouteSegmentTimeSlicePropertyType")
}

func (routesegmenttimeslicetype *RouteSegmentTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RouteSegmentTimeSliceType_stagedOrder, stage.RouteSegmentTimeSliceTypes_referenceOrder, routesegmenttimeslicetype, "RouteSegmentTimeSliceType")
}

func (routesegmenttype *RouteSegmentType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RouteSegmentType_stagedOrder, stage.RouteSegmentTypes_referenceOrder, routesegmenttype, "RouteSegmentType")
}

func (routetimeslicepropertytype *RouteTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RouteTimeSlicePropertyType_stagedOrder, stage.RouteTimeSlicePropertyTypes_referenceOrder, routetimeslicepropertytype, "RouteTimeSlicePropertyType")
}

func (routetimeslicetype *RouteTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RouteTimeSliceType_stagedOrder, stage.RouteTimeSliceTypes_referenceOrder, routetimeslicetype, "RouteTimeSliceType")
}

func (routetype *RouteType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RouteType_stagedOrder, stage.RouteTypes_referenceOrder, routetype, "RouteType")
}

func (rulesprocedurestimeslicepropertytype *RulesProceduresTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RulesProceduresTimeSlicePropertyType_stagedOrder, stage.RulesProceduresTimeSlicePropertyTypes_referenceOrder, rulesprocedurestimeslicepropertytype, "RulesProceduresTimeSlicePropertyType")
}

func (rulesprocedurestimeslicetype *RulesProceduresTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RulesProceduresTimeSliceType_stagedOrder, stage.RulesProceduresTimeSliceTypes_referenceOrder, rulesprocedurestimeslicetype, "RulesProceduresTimeSliceType")
}

func (rulesprocedurestype *RulesProceduresType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RulesProceduresType_stagedOrder, stage.RulesProceduresTypes_referenceOrder, rulesprocedurestype, "RulesProceduresType")
}

func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RunwayCentrelinePointPropertyType_stagedOrder, stage.RunwayCentrelinePointPropertyTypes_referenceOrder, runwaycentrelinepointpropertytype, "RunwayCentrelinePointPropertyType")
}

func (standardinstrumentarrivalpropertytype *StandardInstrumentArrivalPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StandardInstrumentArrivalPropertyType_stagedOrder, stage.StandardInstrumentArrivalPropertyTypes_referenceOrder, standardinstrumentarrivalpropertytype, "StandardInstrumentArrivalPropertyType")
}

func (standardinstrumentdeparturepropertytype *StandardInstrumentDeparturePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StandardInstrumentDeparturePropertyType_stagedOrder, stage.StandardInstrumentDeparturePropertyTypes_referenceOrder, standardinstrumentdeparturepropertytype, "StandardInstrumentDeparturePropertyType")
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

func (telephonecontactpropertytype *TelephoneContactPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TelephoneContactPropertyType_stagedOrder, stage.TelephoneContactPropertyTypes_referenceOrder, telephonecontactpropertytype, "TelephoneContactPropertyType")
}

func (telephonecontacttype *TelephoneContactType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TelephoneContactType_stagedOrder, stage.TelephoneContactTypes_referenceOrder, telephonecontacttype, "TelephoneContactType")
}

func (timesheetpropertytype *TimesheetPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TimesheetPropertyType_stagedOrder, stage.TimesheetPropertyTypes_referenceOrder, timesheetpropertytype, "TimesheetPropertyType")
}

func (timesheettype *TimesheetType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TimesheetType_stagedOrder, stage.TimesheetTypes_referenceOrder, timesheettype, "TimesheetType")
}

func (touchdownliftoffpropertytype *TouchDownLiftOffPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TouchDownLiftOffPropertyType_stagedOrder, stage.TouchDownLiftOffPropertyTypes_referenceOrder, touchdownliftoffpropertytype, "TouchDownLiftOffPropertyType")
}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.VerticalStructurePropertyType_stagedOrder, stage.VerticalStructurePropertyTypes_referenceOrder, verticalstructurepropertytype, "VerticalStructurePropertyType")
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

func (a_extension_15 *A_extension_15) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_15, a_extension_15.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_15 *A_extension_15) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_15.GongGetIdentifier(stage)
}

func (a_extension_16 *A_extension_16) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_16, a_extension_16.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_16 *A_extension_16) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_16.GongGetIdentifier(stage)
}

func (a_extension_17 *A_extension_17) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_17, a_extension_17.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_17 *A_extension_17) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_17.GongGetIdentifier(stage)
}

func (a_extension_18 *A_extension_18) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_18, a_extension_18.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_18 *A_extension_18) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_18.GongGetIdentifier(stage)
}

func (a_extension_19 *A_extension_19) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_19, a_extension_19.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_19 *A_extension_19) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_19.GongGetIdentifier(stage)
}

func (a_extension_2 *A_extension_2) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_2, a_extension_2.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_2 *A_extension_2) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_2.GongGetIdentifier(stage)
}

func (a_extension_20 *A_extension_20) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_20, a_extension_20.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_20 *A_extension_20) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_20.GongGetIdentifier(stage)
}

func (a_extension_21 *A_extension_21) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_21, a_extension_21.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_21 *A_extension_21) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_21.GongGetIdentifier(stage)
}

func (a_extension_22 *A_extension_22) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_22, a_extension_22.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_22 *A_extension_22) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_22.GongGetIdentifier(stage)
}

func (a_extension_23 *A_extension_23) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_23, a_extension_23.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_23 *A_extension_23) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_23.GongGetIdentifier(stage)
}

func (a_extension_24 *A_extension_24) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_24, a_extension_24.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_24 *A_extension_24) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_24.GongGetIdentifier(stage)
}

func (a_extension_25 *A_extension_25) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_25, a_extension_25.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_25 *A_extension_25) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_25.GongGetIdentifier(stage)
}

func (a_extension_26 *A_extension_26) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_26, a_extension_26.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_26 *A_extension_26) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_26.GongGetIdentifier(stage)
}

func (a_extension_27 *A_extension_27) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_27, a_extension_27.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_27 *A_extension_27) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_27.GongGetIdentifier(stage)
}

func (a_extension_28 *A_extension_28) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_28, a_extension_28.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_28 *A_extension_28) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_28.GongGetIdentifier(stage)
}

func (a_extension_29 *A_extension_29) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_29, a_extension_29.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_29 *A_extension_29) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_29.GongGetIdentifier(stage)
}

func (a_extension_3 *A_extension_3) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_3, a_extension_3.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_3 *A_extension_3) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_3.GongGetIdentifier(stage)
}

func (a_extension_30 *A_extension_30) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_30, a_extension_30.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_30 *A_extension_30) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_30.GongGetIdentifier(stage)
}

func (a_extension_31 *A_extension_31) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_31, a_extension_31.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_31 *A_extension_31) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_31.GongGetIdentifier(stage)
}

func (a_extension_32 *A_extension_32) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_32, a_extension_32.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_32 *A_extension_32) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_32.GongGetIdentifier(stage)
}

func (a_extension_33 *A_extension_33) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_33, a_extension_33.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_33 *A_extension_33) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_33.GongGetIdentifier(stage)
}

func (a_extension_34 *A_extension_34) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_34, a_extension_34.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_34 *A_extension_34) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_34.GongGetIdentifier(stage)
}

func (a_extension_35 *A_extension_35) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_35, a_extension_35.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_35 *A_extension_35) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_35.GongGetIdentifier(stage)
}

func (a_extension_36 *A_extension_36) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_36, a_extension_36.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_36 *A_extension_36) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_36.GongGetIdentifier(stage)
}

func (a_extension_37 *A_extension_37) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_37, a_extension_37.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_37 *A_extension_37) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_37.GongGetIdentifier(stage)
}

func (a_extension_38 *A_extension_38) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_38, a_extension_38.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_38 *A_extension_38) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_38.GongGetIdentifier(stage)
}

func (a_extension_39 *A_extension_39) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_39, a_extension_39.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_39 *A_extension_39) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_39.GongGetIdentifier(stage)
}

func (a_extension_4 *A_extension_4) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_4, a_extension_4.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_4 *A_extension_4) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_4.GongGetIdentifier(stage)
}

func (a_extension_40 *A_extension_40) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_40, a_extension_40.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_40 *A_extension_40) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_40.GongGetIdentifier(stage)
}

func (a_extension_41 *A_extension_41) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_41, a_extension_41.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_41 *A_extension_41) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_41.GongGetIdentifier(stage)
}

func (a_extension_42 *A_extension_42) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_extension_42, a_extension_42.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_extension_42 *A_extension_42) GongGetReferenceIdentifier(stage *Stage) string {
	return a_extension_42.GongGetIdentifier(stage)
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

func (abstractdirectflighttype *AbstractDirectFlightType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractdirectflighttype, abstractdirectflighttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractdirectflighttype *AbstractDirectFlightType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractdirectflighttype.GongGetIdentifier(stage)
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractpropertieswithscheduletype, abstractpropertieswithscheduletype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractpropertieswithscheduletype.GongGetIdentifier(stage)
}

func (abstractsegmentpointtype *AbstractSegmentPointType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractsegmentpointtype, abstractsegmentpointtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractsegmentpointtype *AbstractSegmentPointType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractsegmentpointtype.GongGetIdentifier(stage)
}

func (aerialrefuellinganchorpropertytype *AerialRefuellingAnchorPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(aerialrefuellinganchorpropertytype, aerialrefuellinganchorpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (aerialrefuellinganchorpropertytype *AerialRefuellingAnchorPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return aerialrefuellinganchorpropertytype.GongGetIdentifier(stage)
}

func (aerialrefuellinganchortype *AerialRefuellingAnchorType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(aerialrefuellinganchortype, aerialrefuellinganchortype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (aerialrefuellinganchortype *AerialRefuellingAnchorType) GongGetReferenceIdentifier(stage *Stage) string {
	return aerialrefuellinganchortype.GongGetIdentifier(stage)
}

func (aerialrefuellingpointpropertytype *AerialRefuellingPointPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(aerialrefuellingpointpropertytype, aerialrefuellingpointpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (aerialrefuellingpointpropertytype *AerialRefuellingPointPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return aerialrefuellingpointpropertytype.GongGetIdentifier(stage)
}

func (aerialrefuellingpointtype *AerialRefuellingPointType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(aerialrefuellingpointtype, aerialrefuellingpointtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (aerialrefuellingpointtype *AerialRefuellingPointType) GongGetReferenceIdentifier(stage *Stage) string {
	return aerialrefuellingpointtype.GongGetIdentifier(stage)
}

func (aerialrefuellingpropertytype *AerialRefuellingPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(aerialrefuellingpropertytype, aerialrefuellingpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (aerialrefuellingpropertytype *AerialRefuellingPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return aerialrefuellingpropertytype.GongGetIdentifier(stage)
}

func (aerialrefuellingtimeslicepropertytype *AerialRefuellingTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(aerialrefuellingtimeslicepropertytype, aerialrefuellingtimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (aerialrefuellingtimeslicepropertytype *AerialRefuellingTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return aerialrefuellingtimeslicepropertytype.GongGetIdentifier(stage)
}

func (aerialrefuellingtimeslicetype *AerialRefuellingTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(aerialrefuellingtimeslicetype, aerialrefuellingtimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (aerialrefuellingtimeslicetype *AerialRefuellingTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return aerialrefuellingtimeslicetype.GongGetIdentifier(stage)
}

func (aerialrefuellingtrackpropertytype *AerialRefuellingTrackPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(aerialrefuellingtrackpropertytype, aerialrefuellingtrackpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (aerialrefuellingtrackpropertytype *AerialRefuellingTrackPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return aerialrefuellingtrackpropertytype.GongGetIdentifier(stage)
}

func (aerialrefuellingtracktype *AerialRefuellingTrackType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(aerialrefuellingtracktype, aerialrefuellingtracktype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (aerialrefuellingtracktype *AerialRefuellingTrackType) GongGetReferenceIdentifier(stage *Stage) string {
	return aerialrefuellingtracktype.GongGetIdentifier(stage)
}

func (aerialrefuellingtype *AerialRefuellingType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(aerialrefuellingtype, aerialrefuellingtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (aerialrefuellingtype *AerialRefuellingType) GongGetReferenceIdentifier(stage *Stage) string {
	return aerialrefuellingtype.GongGetIdentifier(stage)
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

func (airportheliportpropertytype *AirportHeliportPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airportheliportpropertytype, airportheliportpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airportheliportpropertytype *AirportHeliportPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return airportheliportpropertytype.GongGetIdentifier(stage)
}

func (airspacebordercrossingpropertytype *AirspaceBorderCrossingPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airspacebordercrossingpropertytype, airspacebordercrossingpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airspacebordercrossingpropertytype *AirspaceBorderCrossingPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return airspacebordercrossingpropertytype.GongGetIdentifier(stage)
}

func (airspacebordercrossingtimeslicepropertytype *AirspaceBorderCrossingTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airspacebordercrossingtimeslicepropertytype, airspacebordercrossingtimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airspacebordercrossingtimeslicepropertytype *AirspaceBorderCrossingTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return airspacebordercrossingtimeslicepropertytype.GongGetIdentifier(stage)
}

func (airspacebordercrossingtimeslicetype *AirspaceBorderCrossingTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airspacebordercrossingtimeslicetype, airspacebordercrossingtimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airspacebordercrossingtimeslicetype *AirspaceBorderCrossingTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return airspacebordercrossingtimeslicetype.GongGetIdentifier(stage)
}

func (airspacebordercrossingtype *AirspaceBorderCrossingType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airspacebordercrossingtype, airspacebordercrossingtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airspacebordercrossingtype *AirspaceBorderCrossingType) GongGetReferenceIdentifier(stage *Stage) string {
	return airspacebordercrossingtype.GongGetIdentifier(stage)
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

func (altitudeadjustmentpropertytype *AltitudeAdjustmentPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(altitudeadjustmentpropertytype, altitudeadjustmentpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (altitudeadjustmentpropertytype *AltitudeAdjustmentPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return altitudeadjustmentpropertytype.GongGetIdentifier(stage)
}

func (altitudeadjustmenttype *AltitudeAdjustmentType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(altitudeadjustmenttype, altitudeadjustmenttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (altitudeadjustmenttype *AltitudeAdjustmentType) GongGetReferenceIdentifier(stage *Stage) string {
	return altitudeadjustmenttype.GongGetIdentifier(stage)
}

func (angleindicationpropertytype *AngleIndicationPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(angleindicationpropertytype, angleindicationpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (angleindicationpropertytype *AngleIndicationPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return angleindicationpropertytype.GongGetIdentifier(stage)
}

func (angleusepropertytype *AngleUsePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(angleusepropertytype, angleusepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (angleusepropertytype *AngleUsePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return angleusepropertytype.GongGetIdentifier(stage)
}

func (angleusetype *AngleUseType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(angleusetype, angleusetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (angleusetype *AngleUseType) GongGetReferenceIdentifier(stage *Stage) string {
	return angleusetype.GongGetIdentifier(stage)
}

func (authorityforaerialrefuellingpropertytype *AuthorityForAerialRefuellingPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(authorityforaerialrefuellingpropertytype, authorityforaerialrefuellingpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (authorityforaerialrefuellingpropertytype *AuthorityForAerialRefuellingPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return authorityforaerialrefuellingpropertytype.GongGetIdentifier(stage)
}

func (authorityforaerialrefuellingtype *AuthorityForAerialRefuellingType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(authorityforaerialrefuellingtype, authorityforaerialrefuellingtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (authorityforaerialrefuellingtype *AuthorityForAerialRefuellingType) GongGetReferenceIdentifier(stage *Stage) string {
	return authorityforaerialrefuellingtype.GongGetIdentifier(stage)
}

func (changeoverpointtimeslicepropertytype *ChangeOverPointTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(changeoverpointtimeslicepropertytype, changeoverpointtimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (changeoverpointtimeslicepropertytype *ChangeOverPointTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return changeoverpointtimeslicepropertytype.GongGetIdentifier(stage)
}

func (changeoverpointtimeslicetype *ChangeOverPointTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(changeoverpointtimeslicetype, changeoverpointtimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (changeoverpointtimeslicetype *ChangeOverPointTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return changeoverpointtimeslicetype.GongGetIdentifier(stage)
}

func (changeoverpointtype *ChangeOverPointType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(changeoverpointtype, changeoverpointtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (changeoverpointtype *ChangeOverPointType) GongGetReferenceIdentifier(stage *Stage) string {
	return changeoverpointtype.GongGetIdentifier(stage)
}

func (contactinformationpropertytype *ContactInformationPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(contactinformationpropertytype, contactinformationpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (contactinformationpropertytype *ContactInformationPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return contactinformationpropertytype.GongGetIdentifier(stage)
}

func (contactinformationtype *ContactInformationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(contactinformationtype, contactinformationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (contactinformationtype *ContactInformationType) GongGetReferenceIdentifier(stage *Stage) string {
	return contactinformationtype.GongGetIdentifier(stage)
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

func (dmepropertytype *DMEPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(dmepropertytype, dmepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (dmepropertytype *DMEPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return dmepropertytype.GongGetIdentifier(stage)
}

func (designatedpointpropertytype *DesignatedPointPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(designatedpointpropertytype, designatedpointpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (designatedpointpropertytype *DesignatedPointPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return designatedpointpropertytype.GongGetIdentifier(stage)
}

func (directflightclasstype *DirectFlightClassType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(directflightclasstype, directflightclasstype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (directflightclasstype *DirectFlightClassType) GongGetReferenceIdentifier(stage *Stage) string {
	return directflightclasstype.GongGetIdentifier(stage)
}

func (directflightpropertytype *DirectFlightPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(directflightpropertytype, directflightpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (directflightpropertytype *DirectFlightPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return directflightpropertytype.GongGetIdentifier(stage)
}

func (directflightsegmentpropertytype *DirectFlightSegmentPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(directflightsegmentpropertytype, directflightsegmentpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (directflightsegmentpropertytype *DirectFlightSegmentPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return directflightsegmentpropertytype.GongGetIdentifier(stage)
}

func (directflightsegmenttype *DirectFlightSegmentType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(directflightsegmenttype, directflightsegmenttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (directflightsegmenttype *DirectFlightSegmentType) GongGetReferenceIdentifier(stage *Stage) string {
	return directflightsegmenttype.GongGetIdentifier(stage)
}

func (distanceindicationpropertytype *DistanceIndicationPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(distanceindicationpropertytype, distanceindicationpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (distanceindicationpropertytype *DistanceIndicationPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return distanceindicationpropertytype.GongGetIdentifier(stage)
}

func (elevatedcurvetype *ElevatedCurveType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(elevatedcurvetype, elevatedcurvetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (elevatedcurvetype *ElevatedCurveType) GongGetReferenceIdentifier(stage *Stage) string {
	return elevatedcurvetype.GongGetIdentifier(stage)
}

func (elevatedpointtype *ElevatedPointType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(elevatedpointtype, elevatedpointtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (elevatedpointtype *ElevatedPointType) GongGetReferenceIdentifier(stage *Stage) string {
	return elevatedpointtype.GongGetIdentifier(stage)
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(elevatedsurfacetype, elevatedsurfacetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (elevatedsurfacetype *ElevatedSurfaceType) GongGetReferenceIdentifier(stage *Stage) string {
	return elevatedsurfacetype.GongGetIdentifier(stage)
}

func (enroutesegmentpointpropertytype *EnRouteSegmentPointPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(enroutesegmentpointpropertytype, enroutesegmentpointpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (enroutesegmentpointpropertytype *EnRouteSegmentPointPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return enroutesegmentpointpropertytype.GongGetIdentifier(stage)
}

func (enroutesegmentpointtype *EnRouteSegmentPointType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(enroutesegmentpointtype, enroutesegmentpointtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (enroutesegmentpointtype *EnRouteSegmentPointType) GongGetReferenceIdentifier(stage *Stage) string {
	return enroutesegmentpointtype.GongGetIdentifier(stage)
}

func (flightcharacteristicpropertytype *FlightCharacteristicPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(flightcharacteristicpropertytype, flightcharacteristicpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (flightcharacteristicpropertytype *FlightCharacteristicPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return flightcharacteristicpropertytype.GongGetIdentifier(stage)
}

func (flightcharacteristictype *FlightCharacteristicType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(flightcharacteristictype, flightcharacteristictype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (flightcharacteristictype *FlightCharacteristicType) GongGetReferenceIdentifier(stage *Stage) string {
	return flightcharacteristictype.GongGetIdentifier(stage)
}

func (flightconditioncircumstancepropertytype *FlightConditionCircumstancePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(flightconditioncircumstancepropertytype, flightconditioncircumstancepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (flightconditioncircumstancepropertytype *FlightConditionCircumstancePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return flightconditioncircumstancepropertytype.GongGetIdentifier(stage)
}

func (flightconditioncircumstancetype *FlightConditionCircumstanceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(flightconditioncircumstancetype, flightconditioncircumstancetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (flightconditioncircumstancetype *FlightConditionCircumstanceType) GongGetReferenceIdentifier(stage *Stage) string {
	return flightconditioncircumstancetype.GongGetIdentifier(stage)
}

func (flightconditioncombinationpropertytype *FlightConditionCombinationPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(flightconditioncombinationpropertytype, flightconditioncombinationpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (flightconditioncombinationpropertytype *FlightConditionCombinationPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return flightconditioncombinationpropertytype.GongGetIdentifier(stage)
}

func (flightconditioncombinationtype *FlightConditionCombinationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(flightconditioncombinationtype, flightconditioncombinationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (flightconditioncombinationtype *FlightConditionCombinationType) GongGetReferenceIdentifier(stage *Stage) string {
	return flightconditioncombinationtype.GongGetIdentifier(stage)
}

func (flightconditionelementpropertytype *FlightConditionElementPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(flightconditionelementpropertytype, flightconditionelementpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (flightconditionelementpropertytype *FlightConditionElementPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return flightconditionelementpropertytype.GongGetIdentifier(stage)
}

func (flightconditionelementtype *FlightConditionElementType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(flightconditionelementtype, flightconditionelementtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (flightconditionelementtype *FlightConditionElementType) GongGetReferenceIdentifier(stage *Stage) string {
	return flightconditionelementtype.GongGetIdentifier(stage)
}

func (flightrestrictionlevelpropertytype *FlightRestrictionLevelPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(flightrestrictionlevelpropertytype, flightrestrictionlevelpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (flightrestrictionlevelpropertytype *FlightRestrictionLevelPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return flightrestrictionlevelpropertytype.GongGetIdentifier(stage)
}

func (flightrestrictionleveltype *FlightRestrictionLevelType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(flightrestrictionleveltype, flightrestrictionleveltype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (flightrestrictionleveltype *FlightRestrictionLevelType) GongGetReferenceIdentifier(stage *Stage) string {
	return flightrestrictionleveltype.GongGetIdentifier(stage)
}

func (flightrestrictionroutepropertytype *FlightRestrictionRoutePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(flightrestrictionroutepropertytype, flightrestrictionroutepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (flightrestrictionroutepropertytype *FlightRestrictionRoutePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return flightrestrictionroutepropertytype.GongGetIdentifier(stage)
}

func (flightrestrictionroutetype *FlightRestrictionRouteType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(flightrestrictionroutetype, flightrestrictionroutetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (flightrestrictionroutetype *FlightRestrictionRouteType) GongGetReferenceIdentifier(stage *Stage) string {
	return flightrestrictionroutetype.GongGetIdentifier(stage)
}

func (flightrestrictiontimeslicepropertytype *FlightRestrictionTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(flightrestrictiontimeslicepropertytype, flightrestrictiontimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (flightrestrictiontimeslicepropertytype *FlightRestrictionTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return flightrestrictiontimeslicepropertytype.GongGetIdentifier(stage)
}

func (flightrestrictiontimeslicetype *FlightRestrictionTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(flightrestrictiontimeslicetype, flightrestrictiontimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (flightrestrictiontimeslicetype *FlightRestrictionTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return flightrestrictiontimeslicetype.GongGetIdentifier(stage)
}

func (flightrestrictiontype *FlightRestrictionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(flightrestrictiontype, flightrestrictiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (flightrestrictiontype *FlightRestrictionType) GongGetReferenceIdentifier(stage *Stage) string {
	return flightrestrictiontype.GongGetIdentifier(stage)
}

func (flightroutingelementpropertytype *FlightRoutingElementPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(flightroutingelementpropertytype, flightroutingelementpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (flightroutingelementpropertytype *FlightRoutingElementPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return flightroutingelementpropertytype.GongGetIdentifier(stage)
}

func (flightroutingelementtype *FlightRoutingElementType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(flightroutingelementtype, flightroutingelementtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (flightroutingelementtype *FlightRoutingElementType) GongGetReferenceIdentifier(stage *Stage) string {
	return flightroutingelementtype.GongGetIdentifier(stage)
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

func (meteorologypropertytype *MeteorologyPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(meteorologypropertytype, meteorologypropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (meteorologypropertytype *MeteorologyPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return meteorologypropertytype.GongGetIdentifier(stage)
}

func (meteorologytype *MeteorologyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(meteorologytype, meteorologytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (meteorologytype *MeteorologyType) GongGetReferenceIdentifier(stage *Stage) string {
	return meteorologytype.GongGetIdentifier(stage)
}

func (navaidpropertytype *NavaidPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(navaidpropertytype, navaidpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (navaidpropertytype *NavaidPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return navaidpropertytype.GongGetIdentifier(stage)
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

func (obstacleassessmentareapropertytype *ObstacleAssessmentAreaPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(obstacleassessmentareapropertytype, obstacleassessmentareapropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (obstacleassessmentareapropertytype *ObstacleAssessmentAreaPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return obstacleassessmentareapropertytype.GongGetIdentifier(stage)
}

func (obstacleassessmentareatype *ObstacleAssessmentAreaType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(obstacleassessmentareatype, obstacleassessmentareatype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (obstacleassessmentareatype *ObstacleAssessmentAreaType) GongGetReferenceIdentifier(stage *Stage) string {
	return obstacleassessmentareatype.GongGetIdentifier(stage)
}

func (obstacleplacementpropertytype *ObstaclePlacementPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(obstacleplacementpropertytype, obstacleplacementpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (obstacleplacementpropertytype *ObstaclePlacementPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return obstacleplacementpropertytype.GongGetIdentifier(stage)
}

func (obstacleplacementtype *ObstaclePlacementType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(obstacleplacementtype, obstacleplacementtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (obstacleplacementtype *ObstaclePlacementType) GongGetReferenceIdentifier(stage *Stage) string {
	return obstacleplacementtype.GongGetIdentifier(stage)
}

func (obstructionpropertytype *ObstructionPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(obstructionpropertytype, obstructionpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (obstructionpropertytype *ObstructionPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return obstructionpropertytype.GongGetIdentifier(stage)
}

func (obstructiontype *ObstructionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(obstructiontype, obstructiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (obstructiontype *ObstructionType) GongGetReferenceIdentifier(stage *Stage) string {
	return obstructiontype.GongGetIdentifier(stage)
}

func (onlinecontactpropertytype *OnlineContactPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(onlinecontactpropertytype, onlinecontactpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (onlinecontactpropertytype *OnlineContactPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return onlinecontactpropertytype.GongGetIdentifier(stage)
}

func (onlinecontacttype *OnlineContactType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(onlinecontacttype, onlinecontacttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (onlinecontacttype *OnlineContactType) GongGetReferenceIdentifier(stage *Stage) string {
	return onlinecontacttype.GongGetIdentifier(stage)
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(organisationauthoritypropertytype, organisationauthoritypropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return organisationauthoritypropertytype.GongGetIdentifier(stage)
}

func (pointpropertytype *PointPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(pointpropertytype, pointpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pointpropertytype *PointPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return pointpropertytype.GongGetIdentifier(stage)
}

func (pointreferencepropertytype *PointReferencePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(pointreferencepropertytype, pointreferencepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pointreferencepropertytype *PointReferencePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return pointreferencepropertytype.GongGetIdentifier(stage)
}

func (pointreferencetype *PointReferenceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(pointreferencetype, pointreferencetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pointreferencetype *PointReferenceType) GongGetReferenceIdentifier(stage *Stage) string {
	return pointreferencetype.GongGetIdentifier(stage)
}

func (pointtype *PointType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(pointtype, pointtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pointtype *PointType) GongGetReferenceIdentifier(stage *Stage) string {
	return pointtype.GongGetIdentifier(stage)
}

func (postaladdresspropertytype *PostalAddressPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(postaladdresspropertytype, postaladdresspropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (postaladdresspropertytype *PostalAddressPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return postaladdresspropertytype.GongGetIdentifier(stage)
}

func (postaladdresstype *PostalAddressType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(postaladdresstype, postaladdresstype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (postaladdresstype *PostalAddressType) GongGetReferenceIdentifier(stage *Stage) string {
	return postaladdresstype.GongGetIdentifier(stage)
}

func (radiofrequencyareapropertytype *RadioFrequencyAreaPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(radiofrequencyareapropertytype, radiofrequencyareapropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (radiofrequencyareapropertytype *RadioFrequencyAreaPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return radiofrequencyareapropertytype.GongGetIdentifier(stage)
}

func (routeavailabilitypropertytype *RouteAvailabilityPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(routeavailabilitypropertytype, routeavailabilitypropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (routeavailabilitypropertytype *RouteAvailabilityPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return routeavailabilitypropertytype.GongGetIdentifier(stage)
}

func (routeavailabilitytype *RouteAvailabilityType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(routeavailabilitytype, routeavailabilitytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (routeavailabilitytype *RouteAvailabilityType) GongGetReferenceIdentifier(stage *Stage) string {
	return routeavailabilitytype.GongGetIdentifier(stage)
}

func (routedmetimeslicepropertytype *RouteDMETimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(routedmetimeslicepropertytype, routedmetimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (routedmetimeslicepropertytype *RouteDMETimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return routedmetimeslicepropertytype.GongGetIdentifier(stage)
}

func (routedmetimeslicetype *RouteDMETimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(routedmetimeslicetype, routedmetimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (routedmetimeslicetype *RouteDMETimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return routedmetimeslicetype.GongGetIdentifier(stage)
}

func (routedmetype *RouteDMEType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(routedmetype, routedmetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (routedmetype *RouteDMEType) GongGetReferenceIdentifier(stage *Stage) string {
	return routedmetype.GongGetIdentifier(stage)
}

func (routeportionpropertytype *RoutePortionPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(routeportionpropertytype, routeportionpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (routeportionpropertytype *RoutePortionPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return routeportionpropertytype.GongGetIdentifier(stage)
}

func (routeportiontype *RoutePortionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(routeportiontype, routeportiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (routeportiontype *RoutePortionType) GongGetReferenceIdentifier(stage *Stage) string {
	return routeportiontype.GongGetIdentifier(stage)
}

func (routepropertytype *RoutePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(routepropertytype, routepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (routepropertytype *RoutePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return routepropertytype.GongGetIdentifier(stage)
}

func (routesegmenttimeslicepropertytype *RouteSegmentTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(routesegmenttimeslicepropertytype, routesegmenttimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (routesegmenttimeslicepropertytype *RouteSegmentTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return routesegmenttimeslicepropertytype.GongGetIdentifier(stage)
}

func (routesegmenttimeslicetype *RouteSegmentTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(routesegmenttimeslicetype, routesegmenttimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (routesegmenttimeslicetype *RouteSegmentTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return routesegmenttimeslicetype.GongGetIdentifier(stage)
}

func (routesegmenttype *RouteSegmentType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(routesegmenttype, routesegmenttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (routesegmenttype *RouteSegmentType) GongGetReferenceIdentifier(stage *Stage) string {
	return routesegmenttype.GongGetIdentifier(stage)
}

func (routetimeslicepropertytype *RouteTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(routetimeslicepropertytype, routetimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (routetimeslicepropertytype *RouteTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return routetimeslicepropertytype.GongGetIdentifier(stage)
}

func (routetimeslicetype *RouteTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(routetimeslicetype, routetimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (routetimeslicetype *RouteTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return routetimeslicetype.GongGetIdentifier(stage)
}

func (routetype *RouteType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(routetype, routetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (routetype *RouteType) GongGetReferenceIdentifier(stage *Stage) string {
	return routetype.GongGetIdentifier(stage)
}

func (rulesprocedurestimeslicepropertytype *RulesProceduresTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(rulesprocedurestimeslicepropertytype, rulesprocedurestimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rulesprocedurestimeslicepropertytype *RulesProceduresTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return rulesprocedurestimeslicepropertytype.GongGetIdentifier(stage)
}

func (rulesprocedurestimeslicetype *RulesProceduresTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(rulesprocedurestimeslicetype, rulesprocedurestimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rulesprocedurestimeslicetype *RulesProceduresTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return rulesprocedurestimeslicetype.GongGetIdentifier(stage)
}

func (rulesprocedurestype *RulesProceduresType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(rulesprocedurestype, rulesprocedurestype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rulesprocedurestype *RulesProceduresType) GongGetReferenceIdentifier(stage *Stage) string {
	return rulesprocedurestype.GongGetIdentifier(stage)
}

func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(runwaycentrelinepointpropertytype, runwaycentrelinepointpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return runwaycentrelinepointpropertytype.GongGetIdentifier(stage)
}

func (standardinstrumentarrivalpropertytype *StandardInstrumentArrivalPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(standardinstrumentarrivalpropertytype, standardinstrumentarrivalpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (standardinstrumentarrivalpropertytype *StandardInstrumentArrivalPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return standardinstrumentarrivalpropertytype.GongGetIdentifier(stage)
}

func (standardinstrumentdeparturepropertytype *StandardInstrumentDeparturePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(standardinstrumentdeparturepropertytype, standardinstrumentdeparturepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (standardinstrumentdeparturepropertytype *StandardInstrumentDeparturePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return standardinstrumentdeparturepropertytype.GongGetIdentifier(stage)
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

func (telephonecontactpropertytype *TelephoneContactPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(telephonecontactpropertytype, telephonecontactpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (telephonecontactpropertytype *TelephoneContactPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return telephonecontactpropertytype.GongGetIdentifier(stage)
}

func (telephonecontacttype *TelephoneContactType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(telephonecontacttype, telephonecontacttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (telephonecontacttype *TelephoneContactType) GongGetReferenceIdentifier(stage *Stage) string {
	return telephonecontacttype.GongGetIdentifier(stage)
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

func (touchdownliftoffpropertytype *TouchDownLiftOffPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(touchdownliftoffpropertytype, touchdownliftoffpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (touchdownliftoffpropertytype *TouchDownLiftOffPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return touchdownliftoffpropertytype.GongGetIdentifier(stage)
}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(verticalstructurepropertytype, verticalstructurepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (verticalstructurepropertytype *VerticalStructurePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return verticalstructurepropertytype.GongGetIdentifier(stage)
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

func (a_extension_15 *A_extension_15) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_15.GongGetIdentifier(stage), "A_extension_15", a_extension_15.Name)
}

func (a_extension_16 *A_extension_16) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_16.GongGetIdentifier(stage), "A_extension_16", a_extension_16.Name)
}

func (a_extension_17 *A_extension_17) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_17.GongGetIdentifier(stage), "A_extension_17", a_extension_17.Name)
}

func (a_extension_18 *A_extension_18) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_18.GongGetIdentifier(stage), "A_extension_18", a_extension_18.Name)
}

func (a_extension_19 *A_extension_19) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_19.GongGetIdentifier(stage), "A_extension_19", a_extension_19.Name)
}

func (a_extension_2 *A_extension_2) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_2.GongGetIdentifier(stage), "A_extension_2", a_extension_2.Name)
}

func (a_extension_20 *A_extension_20) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_20.GongGetIdentifier(stage), "A_extension_20", a_extension_20.Name)
}

func (a_extension_21 *A_extension_21) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_21.GongGetIdentifier(stage), "A_extension_21", a_extension_21.Name)
}

func (a_extension_22 *A_extension_22) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_22.GongGetIdentifier(stage), "A_extension_22", a_extension_22.Name)
}

func (a_extension_23 *A_extension_23) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_23.GongGetIdentifier(stage), "A_extension_23", a_extension_23.Name)
}

func (a_extension_24 *A_extension_24) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_24.GongGetIdentifier(stage), "A_extension_24", a_extension_24.Name)
}

func (a_extension_25 *A_extension_25) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_25.GongGetIdentifier(stage), "A_extension_25", a_extension_25.Name)
}

func (a_extension_26 *A_extension_26) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_26.GongGetIdentifier(stage), "A_extension_26", a_extension_26.Name)
}

func (a_extension_27 *A_extension_27) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_27.GongGetIdentifier(stage), "A_extension_27", a_extension_27.Name)
}

func (a_extension_28 *A_extension_28) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_28.GongGetIdentifier(stage), "A_extension_28", a_extension_28.Name)
}

func (a_extension_29 *A_extension_29) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_29.GongGetIdentifier(stage), "A_extension_29", a_extension_29.Name)
}

func (a_extension_3 *A_extension_3) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_3.GongGetIdentifier(stage), "A_extension_3", a_extension_3.Name)
}

func (a_extension_30 *A_extension_30) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_30.GongGetIdentifier(stage), "A_extension_30", a_extension_30.Name)
}

func (a_extension_31 *A_extension_31) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_31.GongGetIdentifier(stage), "A_extension_31", a_extension_31.Name)
}

func (a_extension_32 *A_extension_32) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_32.GongGetIdentifier(stage), "A_extension_32", a_extension_32.Name)
}

func (a_extension_33 *A_extension_33) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_33.GongGetIdentifier(stage), "A_extension_33", a_extension_33.Name)
}

func (a_extension_34 *A_extension_34) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_34.GongGetIdentifier(stage), "A_extension_34", a_extension_34.Name)
}

func (a_extension_35 *A_extension_35) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_35.GongGetIdentifier(stage), "A_extension_35", a_extension_35.Name)
}

func (a_extension_36 *A_extension_36) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_36.GongGetIdentifier(stage), "A_extension_36", a_extension_36.Name)
}

func (a_extension_37 *A_extension_37) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_37.GongGetIdentifier(stage), "A_extension_37", a_extension_37.Name)
}

func (a_extension_38 *A_extension_38) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_38.GongGetIdentifier(stage), "A_extension_38", a_extension_38.Name)
}

func (a_extension_39 *A_extension_39) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_39.GongGetIdentifier(stage), "A_extension_39", a_extension_39.Name)
}

func (a_extension_4 *A_extension_4) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_4.GongGetIdentifier(stage), "A_extension_4", a_extension_4.Name)
}

func (a_extension_40 *A_extension_40) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_40.GongGetIdentifier(stage), "A_extension_40", a_extension_40.Name)
}

func (a_extension_41 *A_extension_41) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_41.GongGetIdentifier(stage), "A_extension_41", a_extension_41.Name)
}

func (a_extension_42 *A_extension_42) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_extension_42.GongGetIdentifier(stage), "A_extension_42", a_extension_42.Name)
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

func (abstractdirectflighttype *AbstractDirectFlightType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractdirectflighttype.GongGetIdentifier(stage), "AbstractDirectFlightType", abstractdirectflighttype.Name)
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractpropertieswithscheduletype.GongGetIdentifier(stage), "AbstractPropertiesWithScheduleType", abstractpropertieswithscheduletype.Name)
}

func (abstractsegmentpointtype *AbstractSegmentPointType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractsegmentpointtype.GongGetIdentifier(stage), "AbstractSegmentPointType", abstractsegmentpointtype.Name)
}

func (aerialrefuellinganchorpropertytype *AerialRefuellingAnchorPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(aerialrefuellinganchorpropertytype.GongGetIdentifier(stage), "AerialRefuellingAnchorPropertyType", aerialrefuellinganchorpropertytype.Name)
}

func (aerialrefuellinganchortype *AerialRefuellingAnchorType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(aerialrefuellinganchortype.GongGetIdentifier(stage), "AerialRefuellingAnchorType", aerialrefuellinganchortype.Name)
}

func (aerialrefuellingpointpropertytype *AerialRefuellingPointPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(aerialrefuellingpointpropertytype.GongGetIdentifier(stage), "AerialRefuellingPointPropertyType", aerialrefuellingpointpropertytype.Name)
}

func (aerialrefuellingpointtype *AerialRefuellingPointType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(aerialrefuellingpointtype.GongGetIdentifier(stage), "AerialRefuellingPointType", aerialrefuellingpointtype.Name)
}

func (aerialrefuellingpropertytype *AerialRefuellingPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(aerialrefuellingpropertytype.GongGetIdentifier(stage), "AerialRefuellingPropertyType", aerialrefuellingpropertytype.Name)
}

func (aerialrefuellingtimeslicepropertytype *AerialRefuellingTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(aerialrefuellingtimeslicepropertytype.GongGetIdentifier(stage), "AerialRefuellingTimeSlicePropertyType", aerialrefuellingtimeslicepropertytype.Name)
}

func (aerialrefuellingtimeslicetype *AerialRefuellingTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(aerialrefuellingtimeslicetype.GongGetIdentifier(stage), "AerialRefuellingTimeSliceType", aerialrefuellingtimeslicetype.Name)
}

func (aerialrefuellingtrackpropertytype *AerialRefuellingTrackPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(aerialrefuellingtrackpropertytype.GongGetIdentifier(stage), "AerialRefuellingTrackPropertyType", aerialrefuellingtrackpropertytype.Name)
}

func (aerialrefuellingtracktype *AerialRefuellingTrackType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(aerialrefuellingtracktype.GongGetIdentifier(stage), "AerialRefuellingTrackType", aerialrefuellingtracktype.Name)
}

func (aerialrefuellingtype *AerialRefuellingType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(aerialrefuellingtype.GongGetIdentifier(stage), "AerialRefuellingType", aerialrefuellingtype.Name)
}

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(aircraftcharacteristicpropertytype.GongGetIdentifier(stage), "AircraftCharacteristicPropertyType", aircraftcharacteristicpropertytype.Name)
}

func (aircraftcharacteristictype *AircraftCharacteristicType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(aircraftcharacteristictype.GongGetIdentifier(stage), "AircraftCharacteristicType", aircraftcharacteristictype.Name)
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airportheliportpropertytype.GongGetIdentifier(stage), "AirportHeliportPropertyType", airportheliportpropertytype.Name)
}

func (airspacebordercrossingpropertytype *AirspaceBorderCrossingPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airspacebordercrossingpropertytype.GongGetIdentifier(stage), "AirspaceBorderCrossingPropertyType", airspacebordercrossingpropertytype.Name)
}

func (airspacebordercrossingtimeslicepropertytype *AirspaceBorderCrossingTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airspacebordercrossingtimeslicepropertytype.GongGetIdentifier(stage), "AirspaceBorderCrossingTimeSlicePropertyType", airspacebordercrossingtimeslicepropertytype.Name)
}

func (airspacebordercrossingtimeslicetype *AirspaceBorderCrossingTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airspacebordercrossingtimeslicetype.GongGetIdentifier(stage), "AirspaceBorderCrossingTimeSliceType", airspacebordercrossingtimeslicetype.Name)
}

func (airspacebordercrossingtype *AirspaceBorderCrossingType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airspacebordercrossingtype.GongGetIdentifier(stage), "AirspaceBorderCrossingType", airspacebordercrossingtype.Name)
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

func (altitudeadjustmentpropertytype *AltitudeAdjustmentPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(altitudeadjustmentpropertytype.GongGetIdentifier(stage), "AltitudeAdjustmentPropertyType", altitudeadjustmentpropertytype.Name)
}

func (altitudeadjustmenttype *AltitudeAdjustmentType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(altitudeadjustmenttype.GongGetIdentifier(stage), "AltitudeAdjustmentType", altitudeadjustmenttype.Name)
}

func (angleindicationpropertytype *AngleIndicationPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(angleindicationpropertytype.GongGetIdentifier(stage), "AngleIndicationPropertyType", angleindicationpropertytype.Name)
}

func (angleusepropertytype *AngleUsePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(angleusepropertytype.GongGetIdentifier(stage), "AngleUsePropertyType", angleusepropertytype.Name)
}

func (angleusetype *AngleUseType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(angleusetype.GongGetIdentifier(stage), "AngleUseType", angleusetype.Name)
}

func (authorityforaerialrefuellingpropertytype *AuthorityForAerialRefuellingPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(authorityforaerialrefuellingpropertytype.GongGetIdentifier(stage), "AuthorityForAerialRefuellingPropertyType", authorityforaerialrefuellingpropertytype.Name)
}

func (authorityforaerialrefuellingtype *AuthorityForAerialRefuellingType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(authorityforaerialrefuellingtype.GongGetIdentifier(stage), "AuthorityForAerialRefuellingType", authorityforaerialrefuellingtype.Name)
}

func (changeoverpointtimeslicepropertytype *ChangeOverPointTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(changeoverpointtimeslicepropertytype.GongGetIdentifier(stage), "ChangeOverPointTimeSlicePropertyType", changeoverpointtimeslicepropertytype.Name)
}

func (changeoverpointtimeslicetype *ChangeOverPointTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(changeoverpointtimeslicetype.GongGetIdentifier(stage), "ChangeOverPointTimeSliceType", changeoverpointtimeslicetype.Name)
}

func (changeoverpointtype *ChangeOverPointType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(changeoverpointtype.GongGetIdentifier(stage), "ChangeOverPointType", changeoverpointtype.Name)
}

func (contactinformationpropertytype *ContactInformationPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(contactinformationpropertytype.GongGetIdentifier(stage), "ContactInformationPropertyType", contactinformationpropertytype.Name)
}

func (contactinformationtype *ContactInformationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(contactinformationtype.GongGetIdentifier(stage), "ContactInformationType", contactinformationtype.Name)
}

func (curvepropertytype *CurvePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(curvepropertytype.GongGetIdentifier(stage), "CurvePropertyType", curvepropertytype.Name)
}

func (curvetype *CurveType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(curvetype.GongGetIdentifier(stage), "CurveType", curvetype.Name)
}

func (dmepropertytype *DMEPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(dmepropertytype.GongGetIdentifier(stage), "DMEPropertyType", dmepropertytype.Name)
}

func (designatedpointpropertytype *DesignatedPointPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(designatedpointpropertytype.GongGetIdentifier(stage), "DesignatedPointPropertyType", designatedpointpropertytype.Name)
}

func (directflightclasstype *DirectFlightClassType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(directflightclasstype.GongGetIdentifier(stage), "DirectFlightClassType", directflightclasstype.Name)
}

func (directflightpropertytype *DirectFlightPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(directflightpropertytype.GongGetIdentifier(stage), "DirectFlightPropertyType", directflightpropertytype.Name)
}

func (directflightsegmentpropertytype *DirectFlightSegmentPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(directflightsegmentpropertytype.GongGetIdentifier(stage), "DirectFlightSegmentPropertyType", directflightsegmentpropertytype.Name)
}

func (directflightsegmenttype *DirectFlightSegmentType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(directflightsegmenttype.GongGetIdentifier(stage), "DirectFlightSegmentType", directflightsegmenttype.Name)
}

func (distanceindicationpropertytype *DistanceIndicationPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(distanceindicationpropertytype.GongGetIdentifier(stage), "DistanceIndicationPropertyType", distanceindicationpropertytype.Name)
}

func (elevatedcurvetype *ElevatedCurveType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(elevatedcurvetype.GongGetIdentifier(stage), "ElevatedCurveType", elevatedcurvetype.Name)
}

func (elevatedpointtype *ElevatedPointType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(elevatedpointtype.GongGetIdentifier(stage), "ElevatedPointType", elevatedpointtype.Name)
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(elevatedsurfacetype.GongGetIdentifier(stage), "ElevatedSurfaceType", elevatedsurfacetype.Name)
}

func (enroutesegmentpointpropertytype *EnRouteSegmentPointPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(enroutesegmentpointpropertytype.GongGetIdentifier(stage), "EnRouteSegmentPointPropertyType", enroutesegmentpointpropertytype.Name)
}

func (enroutesegmentpointtype *EnRouteSegmentPointType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(enroutesegmentpointtype.GongGetIdentifier(stage), "EnRouteSegmentPointType", enroutesegmentpointtype.Name)
}

func (flightcharacteristicpropertytype *FlightCharacteristicPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(flightcharacteristicpropertytype.GongGetIdentifier(stage), "FlightCharacteristicPropertyType", flightcharacteristicpropertytype.Name)
}

func (flightcharacteristictype *FlightCharacteristicType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(flightcharacteristictype.GongGetIdentifier(stage), "FlightCharacteristicType", flightcharacteristictype.Name)
}

func (flightconditioncircumstancepropertytype *FlightConditionCircumstancePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(flightconditioncircumstancepropertytype.GongGetIdentifier(stage), "FlightConditionCircumstancePropertyType", flightconditioncircumstancepropertytype.Name)
}

func (flightconditioncircumstancetype *FlightConditionCircumstanceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(flightconditioncircumstancetype.GongGetIdentifier(stage), "FlightConditionCircumstanceType", flightconditioncircumstancetype.Name)
}

func (flightconditioncombinationpropertytype *FlightConditionCombinationPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(flightconditioncombinationpropertytype.GongGetIdentifier(stage), "FlightConditionCombinationPropertyType", flightconditioncombinationpropertytype.Name)
}

func (flightconditioncombinationtype *FlightConditionCombinationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(flightconditioncombinationtype.GongGetIdentifier(stage), "FlightConditionCombinationType", flightconditioncombinationtype.Name)
}

func (flightconditionelementpropertytype *FlightConditionElementPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(flightconditionelementpropertytype.GongGetIdentifier(stage), "FlightConditionElementPropertyType", flightconditionelementpropertytype.Name)
}

func (flightconditionelementtype *FlightConditionElementType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(flightconditionelementtype.GongGetIdentifier(stage), "FlightConditionElementType", flightconditionelementtype.Name)
}

func (flightrestrictionlevelpropertytype *FlightRestrictionLevelPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(flightrestrictionlevelpropertytype.GongGetIdentifier(stage), "FlightRestrictionLevelPropertyType", flightrestrictionlevelpropertytype.Name)
}

func (flightrestrictionleveltype *FlightRestrictionLevelType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(flightrestrictionleveltype.GongGetIdentifier(stage), "FlightRestrictionLevelType", flightrestrictionleveltype.Name)
}

func (flightrestrictionroutepropertytype *FlightRestrictionRoutePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(flightrestrictionroutepropertytype.GongGetIdentifier(stage), "FlightRestrictionRoutePropertyType", flightrestrictionroutepropertytype.Name)
}

func (flightrestrictionroutetype *FlightRestrictionRouteType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(flightrestrictionroutetype.GongGetIdentifier(stage), "FlightRestrictionRouteType", flightrestrictionroutetype.Name)
}

func (flightrestrictiontimeslicepropertytype *FlightRestrictionTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(flightrestrictiontimeslicepropertytype.GongGetIdentifier(stage), "FlightRestrictionTimeSlicePropertyType", flightrestrictiontimeslicepropertytype.Name)
}

func (flightrestrictiontimeslicetype *FlightRestrictionTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(flightrestrictiontimeslicetype.GongGetIdentifier(stage), "FlightRestrictionTimeSliceType", flightrestrictiontimeslicetype.Name)
}

func (flightrestrictiontype *FlightRestrictionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(flightrestrictiontype.GongGetIdentifier(stage), "FlightRestrictionType", flightrestrictiontype.Name)
}

func (flightroutingelementpropertytype *FlightRoutingElementPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(flightroutingelementpropertytype.GongGetIdentifier(stage), "FlightRoutingElementPropertyType", flightroutingelementpropertytype.Name)
}

func (flightroutingelementtype *FlightRoutingElementType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(flightroutingelementtype.GongGetIdentifier(stage), "FlightRoutingElementType", flightroutingelementtype.Name)
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(linguisticnotepropertytype.GongGetIdentifier(stage), "LinguisticNotePropertyType", linguisticnotepropertytype.Name)
}

func (linguisticnotetype *LinguisticNoteType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(linguisticnotetype.GongGetIdentifier(stage), "LinguisticNoteType", linguisticnotetype.Name)
}

func (meteorologypropertytype *MeteorologyPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(meteorologypropertytype.GongGetIdentifier(stage), "MeteorologyPropertyType", meteorologypropertytype.Name)
}

func (meteorologytype *MeteorologyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(meteorologytype.GongGetIdentifier(stage), "MeteorologyType", meteorologytype.Name)
}

func (navaidpropertytype *NavaidPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(navaidpropertytype.GongGetIdentifier(stage), "NavaidPropertyType", navaidpropertytype.Name)
}

func (notepropertytype *NotePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(notepropertytype.GongGetIdentifier(stage), "NotePropertyType", notepropertytype.Name)
}

func (notetype *NoteType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(notetype.GongGetIdentifier(stage), "NoteType", notetype.Name)
}

func (obstacleassessmentareapropertytype *ObstacleAssessmentAreaPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(obstacleassessmentareapropertytype.GongGetIdentifier(stage), "ObstacleAssessmentAreaPropertyType", obstacleassessmentareapropertytype.Name)
}

func (obstacleassessmentareatype *ObstacleAssessmentAreaType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(obstacleassessmentareatype.GongGetIdentifier(stage), "ObstacleAssessmentAreaType", obstacleassessmentareatype.Name)
}

func (obstacleplacementpropertytype *ObstaclePlacementPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(obstacleplacementpropertytype.GongGetIdentifier(stage), "ObstaclePlacementPropertyType", obstacleplacementpropertytype.Name)
}

func (obstacleplacementtype *ObstaclePlacementType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(obstacleplacementtype.GongGetIdentifier(stage), "ObstaclePlacementType", obstacleplacementtype.Name)
}

func (obstructionpropertytype *ObstructionPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(obstructionpropertytype.GongGetIdentifier(stage), "ObstructionPropertyType", obstructionpropertytype.Name)
}

func (obstructiontype *ObstructionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(obstructiontype.GongGetIdentifier(stage), "ObstructionType", obstructiontype.Name)
}

func (onlinecontactpropertytype *OnlineContactPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(onlinecontactpropertytype.GongGetIdentifier(stage), "OnlineContactPropertyType", onlinecontactpropertytype.Name)
}

func (onlinecontacttype *OnlineContactType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(onlinecontacttype.GongGetIdentifier(stage), "OnlineContactType", onlinecontacttype.Name)
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(organisationauthoritypropertytype.GongGetIdentifier(stage), "OrganisationAuthorityPropertyType", organisationauthoritypropertytype.Name)
}

func (pointpropertytype *PointPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(pointpropertytype.GongGetIdentifier(stage), "PointPropertyType", pointpropertytype.Name)
}

func (pointreferencepropertytype *PointReferencePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(pointreferencepropertytype.GongGetIdentifier(stage), "PointReferencePropertyType", pointreferencepropertytype.Name)
}

func (pointreferencetype *PointReferenceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(pointreferencetype.GongGetIdentifier(stage), "PointReferenceType", pointreferencetype.Name)
}

func (pointtype *PointType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(pointtype.GongGetIdentifier(stage), "PointType", pointtype.Name)
}

func (postaladdresspropertytype *PostalAddressPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(postaladdresspropertytype.GongGetIdentifier(stage), "PostalAddressPropertyType", postaladdresspropertytype.Name)
}

func (postaladdresstype *PostalAddressType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(postaladdresstype.GongGetIdentifier(stage), "PostalAddressType", postaladdresstype.Name)
}

func (radiofrequencyareapropertytype *RadioFrequencyAreaPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(radiofrequencyareapropertytype.GongGetIdentifier(stage), "RadioFrequencyAreaPropertyType", radiofrequencyareapropertytype.Name)
}

func (routeavailabilitypropertytype *RouteAvailabilityPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(routeavailabilitypropertytype.GongGetIdentifier(stage), "RouteAvailabilityPropertyType", routeavailabilitypropertytype.Name)
}

func (routeavailabilitytype *RouteAvailabilityType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(routeavailabilitytype.GongGetIdentifier(stage), "RouteAvailabilityType", routeavailabilitytype.Name)
}

func (routedmetimeslicepropertytype *RouteDMETimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(routedmetimeslicepropertytype.GongGetIdentifier(stage), "RouteDMETimeSlicePropertyType", routedmetimeslicepropertytype.Name)
}

func (routedmetimeslicetype *RouteDMETimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(routedmetimeslicetype.GongGetIdentifier(stage), "RouteDMETimeSliceType", routedmetimeslicetype.Name)
}

func (routedmetype *RouteDMEType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(routedmetype.GongGetIdentifier(stage), "RouteDMEType", routedmetype.Name)
}

func (routeportionpropertytype *RoutePortionPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(routeportionpropertytype.GongGetIdentifier(stage), "RoutePortionPropertyType", routeportionpropertytype.Name)
}

func (routeportiontype *RoutePortionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(routeportiontype.GongGetIdentifier(stage), "RoutePortionType", routeportiontype.Name)
}

func (routepropertytype *RoutePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(routepropertytype.GongGetIdentifier(stage), "RoutePropertyType", routepropertytype.Name)
}

func (routesegmenttimeslicepropertytype *RouteSegmentTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(routesegmenttimeslicepropertytype.GongGetIdentifier(stage), "RouteSegmentTimeSlicePropertyType", routesegmenttimeslicepropertytype.Name)
}

func (routesegmenttimeslicetype *RouteSegmentTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(routesegmenttimeslicetype.GongGetIdentifier(stage), "RouteSegmentTimeSliceType", routesegmenttimeslicetype.Name)
}

func (routesegmenttype *RouteSegmentType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(routesegmenttype.GongGetIdentifier(stage), "RouteSegmentType", routesegmenttype.Name)
}

func (routetimeslicepropertytype *RouteTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(routetimeslicepropertytype.GongGetIdentifier(stage), "RouteTimeSlicePropertyType", routetimeslicepropertytype.Name)
}

func (routetimeslicetype *RouteTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(routetimeslicetype.GongGetIdentifier(stage), "RouteTimeSliceType", routetimeslicetype.Name)
}

func (routetype *RouteType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(routetype.GongGetIdentifier(stage), "RouteType", routetype.Name)
}

func (rulesprocedurestimeslicepropertytype *RulesProceduresTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(rulesprocedurestimeslicepropertytype.GongGetIdentifier(stage), "RulesProceduresTimeSlicePropertyType", rulesprocedurestimeslicepropertytype.Name)
}

func (rulesprocedurestimeslicetype *RulesProceduresTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(rulesprocedurestimeslicetype.GongGetIdentifier(stage), "RulesProceduresTimeSliceType", rulesprocedurestimeslicetype.Name)
}

func (rulesprocedurestype *RulesProceduresType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(rulesprocedurestype.GongGetIdentifier(stage), "RulesProceduresType", rulesprocedurestype.Name)
}

func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(runwaycentrelinepointpropertytype.GongGetIdentifier(stage), "RunwayCentrelinePointPropertyType", runwaycentrelinepointpropertytype.Name)
}

func (standardinstrumentarrivalpropertytype *StandardInstrumentArrivalPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(standardinstrumentarrivalpropertytype.GongGetIdentifier(stage), "StandardInstrumentArrivalPropertyType", standardinstrumentarrivalpropertytype.Name)
}

func (standardinstrumentdeparturepropertytype *StandardInstrumentDeparturePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(standardinstrumentdeparturepropertytype.GongGetIdentifier(stage), "StandardInstrumentDeparturePropertyType", standardinstrumentdeparturepropertytype.Name)
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

func (telephonecontactpropertytype *TelephoneContactPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(telephonecontactpropertytype.GongGetIdentifier(stage), "TelephoneContactPropertyType", telephonecontactpropertytype.Name)
}

func (telephonecontacttype *TelephoneContactType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(telephonecontacttype.GongGetIdentifier(stage), "TelephoneContactType", telephonecontacttype.Name)
}

func (timesheetpropertytype *TimesheetPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(timesheetpropertytype.GongGetIdentifier(stage), "TimesheetPropertyType", timesheetpropertytype.Name)
}

func (timesheettype *TimesheetType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(timesheettype.GongGetIdentifier(stage), "TimesheetType", timesheettype.Name)
}

func (touchdownliftoffpropertytype *TouchDownLiftOffPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(touchdownliftoffpropertytype.GongGetIdentifier(stage), "TouchDownLiftOffPropertyType", touchdownliftoffpropertytype.Name)
}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(verticalstructurepropertytype.GongGetIdentifier(stage), "VerticalStructurePropertyType", verticalstructurepropertytype.Name)
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

func (a_extension_15 *A_extension_15) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_15.GongGetReferenceIdentifier(stage))
}

func (a_extension_16 *A_extension_16) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_16.GongGetReferenceIdentifier(stage))
}

func (a_extension_17 *A_extension_17) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_17.GongGetReferenceIdentifier(stage))
}

func (a_extension_18 *A_extension_18) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_18.GongGetReferenceIdentifier(stage))
}

func (a_extension_19 *A_extension_19) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_19.GongGetReferenceIdentifier(stage))
}

func (a_extension_2 *A_extension_2) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_2.GongGetReferenceIdentifier(stage))
}

func (a_extension_20 *A_extension_20) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_20.GongGetReferenceIdentifier(stage))
}

func (a_extension_21 *A_extension_21) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_21.GongGetReferenceIdentifier(stage))
}

func (a_extension_22 *A_extension_22) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_22.GongGetReferenceIdentifier(stage))
}

func (a_extension_23 *A_extension_23) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_23.GongGetReferenceIdentifier(stage))
}

func (a_extension_24 *A_extension_24) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_24.GongGetReferenceIdentifier(stage))
}

func (a_extension_25 *A_extension_25) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_25.GongGetReferenceIdentifier(stage))
}

func (a_extension_26 *A_extension_26) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_26.GongGetReferenceIdentifier(stage))
}

func (a_extension_27 *A_extension_27) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_27.GongGetReferenceIdentifier(stage))
}

func (a_extension_28 *A_extension_28) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_28.GongGetReferenceIdentifier(stage))
}

func (a_extension_29 *A_extension_29) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_29.GongGetReferenceIdentifier(stage))
}

func (a_extension_3 *A_extension_3) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_3.GongGetReferenceIdentifier(stage))
}

func (a_extension_30 *A_extension_30) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_30.GongGetReferenceIdentifier(stage))
}

func (a_extension_31 *A_extension_31) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_31.GongGetReferenceIdentifier(stage))
}

func (a_extension_32 *A_extension_32) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_32.GongGetReferenceIdentifier(stage))
}

func (a_extension_33 *A_extension_33) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_33.GongGetReferenceIdentifier(stage))
}

func (a_extension_34 *A_extension_34) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_34.GongGetReferenceIdentifier(stage))
}

func (a_extension_35 *A_extension_35) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_35.GongGetReferenceIdentifier(stage))
}

func (a_extension_36 *A_extension_36) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_36.GongGetReferenceIdentifier(stage))
}

func (a_extension_37 *A_extension_37) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_37.GongGetReferenceIdentifier(stage))
}

func (a_extension_38 *A_extension_38) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_38.GongGetReferenceIdentifier(stage))
}

func (a_extension_39 *A_extension_39) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_39.GongGetReferenceIdentifier(stage))
}

func (a_extension_4 *A_extension_4) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_4.GongGetReferenceIdentifier(stage))
}

func (a_extension_40 *A_extension_40) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_40.GongGetReferenceIdentifier(stage))
}

func (a_extension_41 *A_extension_41) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_41.GongGetReferenceIdentifier(stage))
}

func (a_extension_42 *A_extension_42) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_extension_42.GongGetReferenceIdentifier(stage))
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

func (abstractdirectflighttype *AbstractDirectFlightType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractdirectflighttype.GongGetReferenceIdentifier(stage))
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractpropertieswithscheduletype.GongGetReferenceIdentifier(stage))
}

func (abstractsegmentpointtype *AbstractSegmentPointType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractsegmentpointtype.GongGetReferenceIdentifier(stage))
}

func (aerialrefuellinganchorpropertytype *AerialRefuellingAnchorPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(aerialrefuellinganchorpropertytype.GongGetReferenceIdentifier(stage))
}

func (aerialrefuellinganchortype *AerialRefuellingAnchorType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(aerialrefuellinganchortype.GongGetReferenceIdentifier(stage))
}

func (aerialrefuellingpointpropertytype *AerialRefuellingPointPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(aerialrefuellingpointpropertytype.GongGetReferenceIdentifier(stage))
}

func (aerialrefuellingpointtype *AerialRefuellingPointType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(aerialrefuellingpointtype.GongGetReferenceIdentifier(stage))
}

func (aerialrefuellingpropertytype *AerialRefuellingPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(aerialrefuellingpropertytype.GongGetReferenceIdentifier(stage))
}

func (aerialrefuellingtimeslicepropertytype *AerialRefuellingTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(aerialrefuellingtimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (aerialrefuellingtimeslicetype *AerialRefuellingTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(aerialrefuellingtimeslicetype.GongGetReferenceIdentifier(stage))
}

func (aerialrefuellingtrackpropertytype *AerialRefuellingTrackPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(aerialrefuellingtrackpropertytype.GongGetReferenceIdentifier(stage))
}

func (aerialrefuellingtracktype *AerialRefuellingTrackType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(aerialrefuellingtracktype.GongGetReferenceIdentifier(stage))
}

func (aerialrefuellingtype *AerialRefuellingType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(aerialrefuellingtype.GongGetReferenceIdentifier(stage))
}

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(aircraftcharacteristicpropertytype.GongGetReferenceIdentifier(stage))
}

func (aircraftcharacteristictype *AircraftCharacteristicType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(aircraftcharacteristictype.GongGetReferenceIdentifier(stage))
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airportheliportpropertytype.GongGetReferenceIdentifier(stage))
}

func (airspacebordercrossingpropertytype *AirspaceBorderCrossingPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airspacebordercrossingpropertytype.GongGetReferenceIdentifier(stage))
}

func (airspacebordercrossingtimeslicepropertytype *AirspaceBorderCrossingTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airspacebordercrossingtimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (airspacebordercrossingtimeslicetype *AirspaceBorderCrossingTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airspacebordercrossingtimeslicetype.GongGetReferenceIdentifier(stage))
}

func (airspacebordercrossingtype *AirspaceBorderCrossingType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airspacebordercrossingtype.GongGetReferenceIdentifier(stage))
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

func (altitudeadjustmentpropertytype *AltitudeAdjustmentPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(altitudeadjustmentpropertytype.GongGetReferenceIdentifier(stage))
}

func (altitudeadjustmenttype *AltitudeAdjustmentType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(altitudeadjustmenttype.GongGetReferenceIdentifier(stage))
}

func (angleindicationpropertytype *AngleIndicationPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(angleindicationpropertytype.GongGetReferenceIdentifier(stage))
}

func (angleusepropertytype *AngleUsePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(angleusepropertytype.GongGetReferenceIdentifier(stage))
}

func (angleusetype *AngleUseType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(angleusetype.GongGetReferenceIdentifier(stage))
}

func (authorityforaerialrefuellingpropertytype *AuthorityForAerialRefuellingPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(authorityforaerialrefuellingpropertytype.GongGetReferenceIdentifier(stage))
}

func (authorityforaerialrefuellingtype *AuthorityForAerialRefuellingType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(authorityforaerialrefuellingtype.GongGetReferenceIdentifier(stage))
}

func (changeoverpointtimeslicepropertytype *ChangeOverPointTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(changeoverpointtimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (changeoverpointtimeslicetype *ChangeOverPointTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(changeoverpointtimeslicetype.GongGetReferenceIdentifier(stage))
}

func (changeoverpointtype *ChangeOverPointType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(changeoverpointtype.GongGetReferenceIdentifier(stage))
}

func (contactinformationpropertytype *ContactInformationPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(contactinformationpropertytype.GongGetReferenceIdentifier(stage))
}

func (contactinformationtype *ContactInformationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(contactinformationtype.GongGetReferenceIdentifier(stage))
}

func (curvepropertytype *CurvePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(curvepropertytype.GongGetReferenceIdentifier(stage))
}

func (curvetype *CurveType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(curvetype.GongGetReferenceIdentifier(stage))
}

func (dmepropertytype *DMEPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(dmepropertytype.GongGetReferenceIdentifier(stage))
}

func (designatedpointpropertytype *DesignatedPointPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(designatedpointpropertytype.GongGetReferenceIdentifier(stage))
}

func (directflightclasstype *DirectFlightClassType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(directflightclasstype.GongGetReferenceIdentifier(stage))
}

func (directflightpropertytype *DirectFlightPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(directflightpropertytype.GongGetReferenceIdentifier(stage))
}

func (directflightsegmentpropertytype *DirectFlightSegmentPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(directflightsegmentpropertytype.GongGetReferenceIdentifier(stage))
}

func (directflightsegmenttype *DirectFlightSegmentType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(directflightsegmenttype.GongGetReferenceIdentifier(stage))
}

func (distanceindicationpropertytype *DistanceIndicationPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(distanceindicationpropertytype.GongGetReferenceIdentifier(stage))
}

func (elevatedcurvetype *ElevatedCurveType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(elevatedcurvetype.GongGetReferenceIdentifier(stage))
}

func (elevatedpointtype *ElevatedPointType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(elevatedpointtype.GongGetReferenceIdentifier(stage))
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(elevatedsurfacetype.GongGetReferenceIdentifier(stage))
}

func (enroutesegmentpointpropertytype *EnRouteSegmentPointPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(enroutesegmentpointpropertytype.GongGetReferenceIdentifier(stage))
}

func (enroutesegmentpointtype *EnRouteSegmentPointType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(enroutesegmentpointtype.GongGetReferenceIdentifier(stage))
}

func (flightcharacteristicpropertytype *FlightCharacteristicPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(flightcharacteristicpropertytype.GongGetReferenceIdentifier(stage))
}

func (flightcharacteristictype *FlightCharacteristicType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(flightcharacteristictype.GongGetReferenceIdentifier(stage))
}

func (flightconditioncircumstancepropertytype *FlightConditionCircumstancePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(flightconditioncircumstancepropertytype.GongGetReferenceIdentifier(stage))
}

func (flightconditioncircumstancetype *FlightConditionCircumstanceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(flightconditioncircumstancetype.GongGetReferenceIdentifier(stage))
}

func (flightconditioncombinationpropertytype *FlightConditionCombinationPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(flightconditioncombinationpropertytype.GongGetReferenceIdentifier(stage))
}

func (flightconditioncombinationtype *FlightConditionCombinationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(flightconditioncombinationtype.GongGetReferenceIdentifier(stage))
}

func (flightconditionelementpropertytype *FlightConditionElementPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(flightconditionelementpropertytype.GongGetReferenceIdentifier(stage))
}

func (flightconditionelementtype *FlightConditionElementType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(flightconditionelementtype.GongGetReferenceIdentifier(stage))
}

func (flightrestrictionlevelpropertytype *FlightRestrictionLevelPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(flightrestrictionlevelpropertytype.GongGetReferenceIdentifier(stage))
}

func (flightrestrictionleveltype *FlightRestrictionLevelType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(flightrestrictionleveltype.GongGetReferenceIdentifier(stage))
}

func (flightrestrictionroutepropertytype *FlightRestrictionRoutePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(flightrestrictionroutepropertytype.GongGetReferenceIdentifier(stage))
}

func (flightrestrictionroutetype *FlightRestrictionRouteType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(flightrestrictionroutetype.GongGetReferenceIdentifier(stage))
}

func (flightrestrictiontimeslicepropertytype *FlightRestrictionTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(flightrestrictiontimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (flightrestrictiontimeslicetype *FlightRestrictionTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(flightrestrictiontimeslicetype.GongGetReferenceIdentifier(stage))
}

func (flightrestrictiontype *FlightRestrictionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(flightrestrictiontype.GongGetReferenceIdentifier(stage))
}

func (flightroutingelementpropertytype *FlightRoutingElementPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(flightroutingelementpropertytype.GongGetReferenceIdentifier(stage))
}

func (flightroutingelementtype *FlightRoutingElementType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(flightroutingelementtype.GongGetReferenceIdentifier(stage))
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(linguisticnotepropertytype.GongGetReferenceIdentifier(stage))
}

func (linguisticnotetype *LinguisticNoteType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(linguisticnotetype.GongGetReferenceIdentifier(stage))
}

func (meteorologypropertytype *MeteorologyPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(meteorologypropertytype.GongGetReferenceIdentifier(stage))
}

func (meteorologytype *MeteorologyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(meteorologytype.GongGetReferenceIdentifier(stage))
}

func (navaidpropertytype *NavaidPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(navaidpropertytype.GongGetReferenceIdentifier(stage))
}

func (notepropertytype *NotePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(notepropertytype.GongGetReferenceIdentifier(stage))
}

func (notetype *NoteType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(notetype.GongGetReferenceIdentifier(stage))
}

func (obstacleassessmentareapropertytype *ObstacleAssessmentAreaPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(obstacleassessmentareapropertytype.GongGetReferenceIdentifier(stage))
}

func (obstacleassessmentareatype *ObstacleAssessmentAreaType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(obstacleassessmentareatype.GongGetReferenceIdentifier(stage))
}

func (obstacleplacementpropertytype *ObstaclePlacementPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(obstacleplacementpropertytype.GongGetReferenceIdentifier(stage))
}

func (obstacleplacementtype *ObstaclePlacementType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(obstacleplacementtype.GongGetReferenceIdentifier(stage))
}

func (obstructionpropertytype *ObstructionPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(obstructionpropertytype.GongGetReferenceIdentifier(stage))
}

func (obstructiontype *ObstructionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(obstructiontype.GongGetReferenceIdentifier(stage))
}

func (onlinecontactpropertytype *OnlineContactPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(onlinecontactpropertytype.GongGetReferenceIdentifier(stage))
}

func (onlinecontacttype *OnlineContactType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(onlinecontacttype.GongGetReferenceIdentifier(stage))
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(organisationauthoritypropertytype.GongGetReferenceIdentifier(stage))
}

func (pointpropertytype *PointPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(pointpropertytype.GongGetReferenceIdentifier(stage))
}

func (pointreferencepropertytype *PointReferencePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(pointreferencepropertytype.GongGetReferenceIdentifier(stage))
}

func (pointreferencetype *PointReferenceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(pointreferencetype.GongGetReferenceIdentifier(stage))
}

func (pointtype *PointType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(pointtype.GongGetReferenceIdentifier(stage))
}

func (postaladdresspropertytype *PostalAddressPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(postaladdresspropertytype.GongGetReferenceIdentifier(stage))
}

func (postaladdresstype *PostalAddressType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(postaladdresstype.GongGetReferenceIdentifier(stage))
}

func (radiofrequencyareapropertytype *RadioFrequencyAreaPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(radiofrequencyareapropertytype.GongGetReferenceIdentifier(stage))
}

func (routeavailabilitypropertytype *RouteAvailabilityPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(routeavailabilitypropertytype.GongGetReferenceIdentifier(stage))
}

func (routeavailabilitytype *RouteAvailabilityType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(routeavailabilitytype.GongGetReferenceIdentifier(stage))
}

func (routedmetimeslicepropertytype *RouteDMETimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(routedmetimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (routedmetimeslicetype *RouteDMETimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(routedmetimeslicetype.GongGetReferenceIdentifier(stage))
}

func (routedmetype *RouteDMEType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(routedmetype.GongGetReferenceIdentifier(stage))
}

func (routeportionpropertytype *RoutePortionPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(routeportionpropertytype.GongGetReferenceIdentifier(stage))
}

func (routeportiontype *RoutePortionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(routeportiontype.GongGetReferenceIdentifier(stage))
}

func (routepropertytype *RoutePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(routepropertytype.GongGetReferenceIdentifier(stage))
}

func (routesegmenttimeslicepropertytype *RouteSegmentTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(routesegmenttimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (routesegmenttimeslicetype *RouteSegmentTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(routesegmenttimeslicetype.GongGetReferenceIdentifier(stage))
}

func (routesegmenttype *RouteSegmentType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(routesegmenttype.GongGetReferenceIdentifier(stage))
}

func (routetimeslicepropertytype *RouteTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(routetimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (routetimeslicetype *RouteTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(routetimeslicetype.GongGetReferenceIdentifier(stage))
}

func (routetype *RouteType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(routetype.GongGetReferenceIdentifier(stage))
}

func (rulesprocedurestimeslicepropertytype *RulesProceduresTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(rulesprocedurestimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (rulesprocedurestimeslicetype *RulesProceduresTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(rulesprocedurestimeslicetype.GongGetReferenceIdentifier(stage))
}

func (rulesprocedurestype *RulesProceduresType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(rulesprocedurestype.GongGetReferenceIdentifier(stage))
}

func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(runwaycentrelinepointpropertytype.GongGetReferenceIdentifier(stage))
}

func (standardinstrumentarrivalpropertytype *StandardInstrumentArrivalPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(standardinstrumentarrivalpropertytype.GongGetReferenceIdentifier(stage))
}

func (standardinstrumentdeparturepropertytype *StandardInstrumentDeparturePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(standardinstrumentdeparturepropertytype.GongGetReferenceIdentifier(stage))
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

func (telephonecontactpropertytype *TelephoneContactPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(telephonecontactpropertytype.GongGetReferenceIdentifier(stage))
}

func (telephonecontacttype *TelephoneContactType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(telephonecontacttype.GongGetReferenceIdentifier(stage))
}

func (timesheetpropertytype *TimesheetPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(timesheetpropertytype.GongGetReferenceIdentifier(stage))
}

func (timesheettype *TimesheetType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(timesheettype.GongGetReferenceIdentifier(stage))
}

func (touchdownliftoffpropertytype *TouchDownLiftOffPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(touchdownliftoffpropertytype.GongGetReferenceIdentifier(stage))
}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(verticalstructurepropertytype.GongGetReferenceIdentifier(stage))
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
