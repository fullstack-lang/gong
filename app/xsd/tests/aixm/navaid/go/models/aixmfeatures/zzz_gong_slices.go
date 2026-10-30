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
	// Compute reverse map for named struct AngleIndicationTimeSliceType
	// insertion point per field
	stage.AngleIndicationTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*AngleIndicationTimeSliceType)
	for angleindicationtimeslicetype := range stage.AngleIndicationTimeSliceTypes {
		_ = angleindicationtimeslicetype
		for _, _notepropertytype := range angleindicationtimeslicetype.Annotation {
			stage.AngleIndicationTimeSliceType_Annotation_reverseMap[_notepropertytype] = angleindicationtimeslicetype
		}
	}
	stage.AngleIndicationTimeSliceType_Extension_reverseMap = make(map[*A_extension_24]*AngleIndicationTimeSliceType)
	for angleindicationtimeslicetype := range stage.AngleIndicationTimeSliceTypes {
		_ = angleindicationtimeslicetype
		for _, _a_extension_24 := range angleindicationtimeslicetype.Extension {
			stage.AngleIndicationTimeSliceType_Extension_reverseMap[_a_extension_24] = angleindicationtimeslicetype
		}
	}

	// Compute reverse map for named struct AngleIndicationType
	// insertion point per field
	stage.AngleIndicationType_TimeSlice_reverseMap = make(map[*AngleIndicationTimeSlicePropertyType]*AngleIndicationType)
	for angleindicationtype := range stage.AngleIndicationTypes {
		_ = angleindicationtype
		for _, _angleindicationtimeslicepropertytype := range angleindicationtype.TimeSlice {
			stage.AngleIndicationType_TimeSlice_reverseMap[_angleindicationtimeslicepropertytype] = angleindicationtype
		}
	}

	// Compute reverse map for named struct AuthorityForNavaidEquipmentType
	// insertion point per field
	stage.AuthorityForNavaidEquipmentType_Annotation_reverseMap = make(map[*NotePropertyType]*AuthorityForNavaidEquipmentType)
	for authorityfornavaidequipmenttype := range stage.AuthorityForNavaidEquipmentTypes {
		_ = authorityfornavaidequipmenttype
		for _, _notepropertytype := range authorityfornavaidequipmenttype.Annotation {
			stage.AuthorityForNavaidEquipmentType_Annotation_reverseMap[_notepropertytype] = authorityfornavaidequipmenttype
		}
	}
	stage.AuthorityForNavaidEquipmentType_Extension_reverseMap = make(map[*A_extension_1]*AuthorityForNavaidEquipmentType)
	for authorityfornavaidequipmenttype := range stage.AuthorityForNavaidEquipmentTypes {
		_ = authorityfornavaidequipmenttype
		for _, _a_extension_1 := range authorityfornavaidequipmenttype.Extension {
			stage.AuthorityForNavaidEquipmentType_Extension_reverseMap[_a_extension_1] = authorityfornavaidequipmenttype
		}
	}

	// Compute reverse map for named struct AuthorityForSpecialNavigationStationType
	// insertion point per field
	stage.AuthorityForSpecialNavigationStationType_Annotation_reverseMap = make(map[*NotePropertyType]*AuthorityForSpecialNavigationStationType)
	for authorityforspecialnavigationstationtype := range stage.AuthorityForSpecialNavigationStationTypes {
		_ = authorityforspecialnavigationstationtype
		for _, _notepropertytype := range authorityforspecialnavigationstationtype.Annotation {
			stage.AuthorityForSpecialNavigationStationType_Annotation_reverseMap[_notepropertytype] = authorityforspecialnavigationstationtype
		}
	}
	stage.AuthorityForSpecialNavigationStationType_Extension_reverseMap = make(map[*A_extension_2]*AuthorityForSpecialNavigationStationType)
	for authorityforspecialnavigationstationtype := range stage.AuthorityForSpecialNavigationStationTypes {
		_ = authorityforspecialnavigationstationtype
		for _, _a_extension_2 := range authorityforspecialnavigationstationtype.Extension {
			stage.AuthorityForSpecialNavigationStationType_Extension_reverseMap[_a_extension_2] = authorityforspecialnavigationstationtype
		}
	}

	// Compute reverse map for named struct AuthorityForSpecialNavigationSystemType
	// insertion point per field
	stage.AuthorityForSpecialNavigationSystemType_Annotation_reverseMap = make(map[*NotePropertyType]*AuthorityForSpecialNavigationSystemType)
	for authorityforspecialnavigationsystemtype := range stage.AuthorityForSpecialNavigationSystemTypes {
		_ = authorityforspecialnavigationsystemtype
		for _, _notepropertytype := range authorityforspecialnavigationsystemtype.Annotation {
			stage.AuthorityForSpecialNavigationSystemType_Annotation_reverseMap[_notepropertytype] = authorityforspecialnavigationsystemtype
		}
	}
	stage.AuthorityForSpecialNavigationSystemType_Extension_reverseMap = make(map[*A_extension_3]*AuthorityForSpecialNavigationSystemType)
	for authorityforspecialnavigationsystemtype := range stage.AuthorityForSpecialNavigationSystemTypes {
		_ = authorityforspecialnavigationsystemtype
		for _, _a_extension_3 := range authorityforspecialnavigationsystemtype.Extension {
			stage.AuthorityForSpecialNavigationSystemType_Extension_reverseMap[_a_extension_3] = authorityforspecialnavigationsystemtype
		}
	}

	// Compute reverse map for named struct AzimuthTimeSliceType
	// insertion point per field
	stage.AzimuthTimeSliceType_Authority_reverseMap = make(map[*AuthorityForNavaidEquipmentPropertyType]*AzimuthTimeSliceType)
	for azimuthtimeslicetype := range stage.AzimuthTimeSliceTypes {
		_ = azimuthtimeslicetype
		for _, _authorityfornavaidequipmentpropertytype := range azimuthtimeslicetype.Authority {
			stage.AzimuthTimeSliceType_Authority_reverseMap[_authorityfornavaidequipmentpropertytype] = azimuthtimeslicetype
		}
	}
	stage.AzimuthTimeSliceType_Monitoring_reverseMap = make(map[*NavaidEquipmentMonitoringPropertyType]*AzimuthTimeSliceType)
	for azimuthtimeslicetype := range stage.AzimuthTimeSliceTypes {
		_ = azimuthtimeslicetype
		for _, _navaidequipmentmonitoringpropertytype := range azimuthtimeslicetype.Monitoring {
			stage.AzimuthTimeSliceType_Monitoring_reverseMap[_navaidequipmentmonitoringpropertytype] = azimuthtimeslicetype
		}
	}
	stage.AzimuthTimeSliceType_Availability_reverseMap = make(map[*NavaidOperationalStatusPropertyType]*AzimuthTimeSliceType)
	for azimuthtimeslicetype := range stage.AzimuthTimeSliceTypes {
		_ = azimuthtimeslicetype
		for _, _navaidoperationalstatuspropertytype := range azimuthtimeslicetype.Availability {
			stage.AzimuthTimeSliceType_Availability_reverseMap[_navaidoperationalstatuspropertytype] = azimuthtimeslicetype
		}
	}
	stage.AzimuthTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*AzimuthTimeSliceType)
	for azimuthtimeslicetype := range stage.AzimuthTimeSliceTypes {
		_ = azimuthtimeslicetype
		for _, _notepropertytype := range azimuthtimeslicetype.Annotation {
			stage.AzimuthTimeSliceType_Annotation_reverseMap[_notepropertytype] = azimuthtimeslicetype
		}
	}
	stage.AzimuthTimeSliceType_Extension_reverseMap = make(map[*A_extension_4]*AzimuthTimeSliceType)
	for azimuthtimeslicetype := range stage.AzimuthTimeSliceTypes {
		_ = azimuthtimeslicetype
		for _, _a_extension_4 := range azimuthtimeslicetype.Extension {
			stage.AzimuthTimeSliceType_Extension_reverseMap[_a_extension_4] = azimuthtimeslicetype
		}
	}

	// Compute reverse map for named struct AzimuthType
	// insertion point per field
	stage.AzimuthType_TimeSlice_reverseMap = make(map[*AzimuthTimeSlicePropertyType]*AzimuthType)
	for azimuthtype := range stage.AzimuthTypes {
		_ = azimuthtype
		for _, _azimuthtimeslicepropertytype := range azimuthtype.TimeSlice {
			stage.AzimuthType_TimeSlice_reverseMap[_azimuthtimeslicepropertytype] = azimuthtype
		}
	}

	// Compute reverse map for named struct CheckpointINSTimeSliceType
	// insertion point per field
	stage.CheckpointINSTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*CheckpointINSTimeSliceType)
	for checkpointinstimeslicetype := range stage.CheckpointINSTimeSliceTypes {
		_ = checkpointinstimeslicetype
		for _, _notepropertytype := range checkpointinstimeslicetype.Annotation {
			stage.CheckpointINSTimeSliceType_Annotation_reverseMap[_notepropertytype] = checkpointinstimeslicetype
		}
	}
	stage.CheckpointINSTimeSliceType_Extension_reverseMap = make(map[*A_extension_5]*CheckpointINSTimeSliceType)
	for checkpointinstimeslicetype := range stage.CheckpointINSTimeSliceTypes {
		_ = checkpointinstimeslicetype
		for _, _a_extension_5 := range checkpointinstimeslicetype.Extension {
			stage.CheckpointINSTimeSliceType_Extension_reverseMap[_a_extension_5] = checkpointinstimeslicetype
		}
	}

	// Compute reverse map for named struct CheckpointINSType
	// insertion point per field
	stage.CheckpointINSType_TimeSlice_reverseMap = make(map[*CheckpointINSTimeSlicePropertyType]*CheckpointINSType)
	for checkpointinstype := range stage.CheckpointINSTypes {
		_ = checkpointinstype
		for _, _checkpointinstimeslicepropertytype := range checkpointinstype.TimeSlice {
			stage.CheckpointINSType_TimeSlice_reverseMap[_checkpointinstimeslicepropertytype] = checkpointinstype
		}
	}

	// Compute reverse map for named struct CheckpointVORTimeSliceType
	// insertion point per field
	stage.CheckpointVORTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*CheckpointVORTimeSliceType)
	for checkpointvortimeslicetype := range stage.CheckpointVORTimeSliceTypes {
		_ = checkpointvortimeslicetype
		for _, _notepropertytype := range checkpointvortimeslicetype.Annotation {
			stage.CheckpointVORTimeSliceType_Annotation_reverseMap[_notepropertytype] = checkpointvortimeslicetype
		}
	}
	stage.CheckpointVORTimeSliceType_Extension_reverseMap = make(map[*A_extension_6]*CheckpointVORTimeSliceType)
	for checkpointvortimeslicetype := range stage.CheckpointVORTimeSliceTypes {
		_ = checkpointvortimeslicetype
		for _, _a_extension_6 := range checkpointvortimeslicetype.Extension {
			stage.CheckpointVORTimeSliceType_Extension_reverseMap[_a_extension_6] = checkpointvortimeslicetype
		}
	}

	// Compute reverse map for named struct CheckpointVORType
	// insertion point per field
	stage.CheckpointVORType_TimeSlice_reverseMap = make(map[*CheckpointVORTimeSlicePropertyType]*CheckpointVORType)
	for checkpointvortype := range stage.CheckpointVORTypes {
		_ = checkpointvortype
		for _, _checkpointvortimeslicepropertytype := range checkpointvortype.TimeSlice {
			stage.CheckpointVORType_TimeSlice_reverseMap[_checkpointvortimeslicepropertytype] = checkpointvortype
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

	// Compute reverse map for named struct DMETimeSliceType
	// insertion point per field
	stage.DMETimeSliceType_Authority_reverseMap = make(map[*AuthorityForNavaidEquipmentPropertyType]*DMETimeSliceType)
	for dmetimeslicetype := range stage.DMETimeSliceTypes {
		_ = dmetimeslicetype
		for _, _authorityfornavaidequipmentpropertytype := range dmetimeslicetype.Authority {
			stage.DMETimeSliceType_Authority_reverseMap[_authorityfornavaidequipmentpropertytype] = dmetimeslicetype
		}
	}
	stage.DMETimeSliceType_Monitoring_reverseMap = make(map[*NavaidEquipmentMonitoringPropertyType]*DMETimeSliceType)
	for dmetimeslicetype := range stage.DMETimeSliceTypes {
		_ = dmetimeslicetype
		for _, _navaidequipmentmonitoringpropertytype := range dmetimeslicetype.Monitoring {
			stage.DMETimeSliceType_Monitoring_reverseMap[_navaidequipmentmonitoringpropertytype] = dmetimeslicetype
		}
	}
	stage.DMETimeSliceType_Availability_reverseMap = make(map[*NavaidOperationalStatusPropertyType]*DMETimeSliceType)
	for dmetimeslicetype := range stage.DMETimeSliceTypes {
		_ = dmetimeslicetype
		for _, _navaidoperationalstatuspropertytype := range dmetimeslicetype.Availability {
			stage.DMETimeSliceType_Availability_reverseMap[_navaidoperationalstatuspropertytype] = dmetimeslicetype
		}
	}
	stage.DMETimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*DMETimeSliceType)
	for dmetimeslicetype := range stage.DMETimeSliceTypes {
		_ = dmetimeslicetype
		for _, _notepropertytype := range dmetimeslicetype.Annotation {
			stage.DMETimeSliceType_Annotation_reverseMap[_notepropertytype] = dmetimeslicetype
		}
	}
	stage.DMETimeSliceType_Extension_reverseMap = make(map[*A_extension_8]*DMETimeSliceType)
	for dmetimeslicetype := range stage.DMETimeSliceTypes {
		_ = dmetimeslicetype
		for _, _a_extension_8 := range dmetimeslicetype.Extension {
			stage.DMETimeSliceType_Extension_reverseMap[_a_extension_8] = dmetimeslicetype
		}
	}

	// Compute reverse map for named struct DMEType
	// insertion point per field
	stage.DMEType_TimeSlice_reverseMap = make(map[*DMETimeSlicePropertyType]*DMEType)
	for dmetype := range stage.DMETypes {
		_ = dmetype
		for _, _dmetimeslicepropertytype := range dmetype.TimeSlice {
			stage.DMEType_TimeSlice_reverseMap[_dmetimeslicepropertytype] = dmetype
		}
	}

	// Compute reverse map for named struct DesignatedPointTimeSliceType
	// insertion point per field
	stage.DesignatedPointTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*DesignatedPointTimeSliceType)
	for designatedpointtimeslicetype := range stage.DesignatedPointTimeSliceTypes {
		_ = designatedpointtimeslicetype
		for _, _notepropertytype := range designatedpointtimeslicetype.Annotation {
			stage.DesignatedPointTimeSliceType_Annotation_reverseMap[_notepropertytype] = designatedpointtimeslicetype
		}
	}
	stage.DesignatedPointTimeSliceType_Extension_reverseMap = make(map[*A_extension_26]*DesignatedPointTimeSliceType)
	for designatedpointtimeslicetype := range stage.DesignatedPointTimeSliceTypes {
		_ = designatedpointtimeslicetype
		for _, _a_extension_26 := range designatedpointtimeslicetype.Extension {
			stage.DesignatedPointTimeSliceType_Extension_reverseMap[_a_extension_26] = designatedpointtimeslicetype
		}
	}

	// Compute reverse map for named struct DesignatedPointType
	// insertion point per field
	stage.DesignatedPointType_TimeSlice_reverseMap = make(map[*DesignatedPointTimeSlicePropertyType]*DesignatedPointType)
	for designatedpointtype := range stage.DesignatedPointTypes {
		_ = designatedpointtype
		for _, _designatedpointtimeslicepropertytype := range designatedpointtype.TimeSlice {
			stage.DesignatedPointType_TimeSlice_reverseMap[_designatedpointtimeslicepropertytype] = designatedpointtype
		}
	}

	// Compute reverse map for named struct DirectionFinderTimeSliceType
	// insertion point per field
	stage.DirectionFinderTimeSliceType_Authority_reverseMap = make(map[*AuthorityForNavaidEquipmentPropertyType]*DirectionFinderTimeSliceType)
	for directionfindertimeslicetype := range stage.DirectionFinderTimeSliceTypes {
		_ = directionfindertimeslicetype
		for _, _authorityfornavaidequipmentpropertytype := range directionfindertimeslicetype.Authority {
			stage.DirectionFinderTimeSliceType_Authority_reverseMap[_authorityfornavaidequipmentpropertytype] = directionfindertimeslicetype
		}
	}
	stage.DirectionFinderTimeSliceType_Monitoring_reverseMap = make(map[*NavaidEquipmentMonitoringPropertyType]*DirectionFinderTimeSliceType)
	for directionfindertimeslicetype := range stage.DirectionFinderTimeSliceTypes {
		_ = directionfindertimeslicetype
		for _, _navaidequipmentmonitoringpropertytype := range directionfindertimeslicetype.Monitoring {
			stage.DirectionFinderTimeSliceType_Monitoring_reverseMap[_navaidequipmentmonitoringpropertytype] = directionfindertimeslicetype
		}
	}
	stage.DirectionFinderTimeSliceType_Availability_reverseMap = make(map[*NavaidOperationalStatusPropertyType]*DirectionFinderTimeSliceType)
	for directionfindertimeslicetype := range stage.DirectionFinderTimeSliceTypes {
		_ = directionfindertimeslicetype
		for _, _navaidoperationalstatuspropertytype := range directionfindertimeslicetype.Availability {
			stage.DirectionFinderTimeSliceType_Availability_reverseMap[_navaidoperationalstatuspropertytype] = directionfindertimeslicetype
		}
	}
	stage.DirectionFinderTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*DirectionFinderTimeSliceType)
	for directionfindertimeslicetype := range stage.DirectionFinderTimeSliceTypes {
		_ = directionfindertimeslicetype
		for _, _notepropertytype := range directionfindertimeslicetype.Annotation {
			stage.DirectionFinderTimeSliceType_Annotation_reverseMap[_notepropertytype] = directionfindertimeslicetype
		}
	}
	stage.DirectionFinderTimeSliceType_InformationProvision_reverseMap = make(map[*InformationServicePropertyType]*DirectionFinderTimeSliceType)
	for directionfindertimeslicetype := range stage.DirectionFinderTimeSliceTypes {
		_ = directionfindertimeslicetype
		for _, _informationservicepropertytype := range directionfindertimeslicetype.InformationProvision {
			stage.DirectionFinderTimeSliceType_InformationProvision_reverseMap[_informationservicepropertytype] = directionfindertimeslicetype
		}
	}
	stage.DirectionFinderTimeSliceType_Extension_reverseMap = make(map[*A_extension_7]*DirectionFinderTimeSliceType)
	for directionfindertimeslicetype := range stage.DirectionFinderTimeSliceTypes {
		_ = directionfindertimeslicetype
		for _, _a_extension_7 := range directionfindertimeslicetype.Extension {
			stage.DirectionFinderTimeSliceType_Extension_reverseMap[_a_extension_7] = directionfindertimeslicetype
		}
	}

	// Compute reverse map for named struct DirectionFinderType
	// insertion point per field
	stage.DirectionFinderType_TimeSlice_reverseMap = make(map[*DirectionFinderTimeSlicePropertyType]*DirectionFinderType)
	for directionfindertype := range stage.DirectionFinderTypes {
		_ = directionfindertype
		for _, _directionfindertimeslicepropertytype := range directionfindertype.TimeSlice {
			stage.DirectionFinderType_TimeSlice_reverseMap[_directionfindertimeslicepropertytype] = directionfindertype
		}
	}

	// Compute reverse map for named struct DistanceIndicationTimeSliceType
	// insertion point per field
	stage.DistanceIndicationTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*DistanceIndicationTimeSliceType)
	for distanceindicationtimeslicetype := range stage.DistanceIndicationTimeSliceTypes {
		_ = distanceindicationtimeslicetype
		for _, _notepropertytype := range distanceindicationtimeslicetype.Annotation {
			stage.DistanceIndicationTimeSliceType_Annotation_reverseMap[_notepropertytype] = distanceindicationtimeslicetype
		}
	}
	stage.DistanceIndicationTimeSliceType_Extension_reverseMap = make(map[*A_extension_25]*DistanceIndicationTimeSliceType)
	for distanceindicationtimeslicetype := range stage.DistanceIndicationTimeSliceTypes {
		_ = distanceindicationtimeslicetype
		for _, _a_extension_25 := range distanceindicationtimeslicetype.Extension {
			stage.DistanceIndicationTimeSliceType_Extension_reverseMap[_a_extension_25] = distanceindicationtimeslicetype
		}
	}

	// Compute reverse map for named struct DistanceIndicationType
	// insertion point per field
	stage.DistanceIndicationType_TimeSlice_reverseMap = make(map[*DistanceIndicationTimeSlicePropertyType]*DistanceIndicationType)
	for distanceindicationtype := range stage.DistanceIndicationTypes {
		_ = distanceindicationtype
		for _, _distanceindicationtimeslicepropertytype := range distanceindicationtype.TimeSlice {
			stage.DistanceIndicationType_TimeSlice_reverseMap[_distanceindicationtimeslicepropertytype] = distanceindicationtype
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
	stage.ElevatedPointType_Extension_reverseMap = make(map[*A_extension]*ElevatedPointType)
	for elevatedpointtype := range stage.ElevatedPointTypes {
		_ = elevatedpointtype
		for _, _a_extension := range elevatedpointtype.Extension {
			stage.ElevatedPointType_Extension_reverseMap[_a_extension] = elevatedpointtype
		}
	}

	// Compute reverse map for named struct ElevationTimeSliceType
	// insertion point per field
	stage.ElevationTimeSliceType_Authority_reverseMap = make(map[*AuthorityForNavaidEquipmentPropertyType]*ElevationTimeSliceType)
	for elevationtimeslicetype := range stage.ElevationTimeSliceTypes {
		_ = elevationtimeslicetype
		for _, _authorityfornavaidequipmentpropertytype := range elevationtimeslicetype.Authority {
			stage.ElevationTimeSliceType_Authority_reverseMap[_authorityfornavaidequipmentpropertytype] = elevationtimeslicetype
		}
	}
	stage.ElevationTimeSliceType_Monitoring_reverseMap = make(map[*NavaidEquipmentMonitoringPropertyType]*ElevationTimeSliceType)
	for elevationtimeslicetype := range stage.ElevationTimeSliceTypes {
		_ = elevationtimeslicetype
		for _, _navaidequipmentmonitoringpropertytype := range elevationtimeslicetype.Monitoring {
			stage.ElevationTimeSliceType_Monitoring_reverseMap[_navaidequipmentmonitoringpropertytype] = elevationtimeslicetype
		}
	}
	stage.ElevationTimeSliceType_Availability_reverseMap = make(map[*NavaidOperationalStatusPropertyType]*ElevationTimeSliceType)
	for elevationtimeslicetype := range stage.ElevationTimeSliceTypes {
		_ = elevationtimeslicetype
		for _, _navaidoperationalstatuspropertytype := range elevationtimeslicetype.Availability {
			stage.ElevationTimeSliceType_Availability_reverseMap[_navaidoperationalstatuspropertytype] = elevationtimeslicetype
		}
	}
	stage.ElevationTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*ElevationTimeSliceType)
	for elevationtimeslicetype := range stage.ElevationTimeSliceTypes {
		_ = elevationtimeslicetype
		for _, _notepropertytype := range elevationtimeslicetype.Annotation {
			stage.ElevationTimeSliceType_Annotation_reverseMap[_notepropertytype] = elevationtimeslicetype
		}
	}
	stage.ElevationTimeSliceType_Extension_reverseMap = make(map[*A_extension_9]*ElevationTimeSliceType)
	for elevationtimeslicetype := range stage.ElevationTimeSliceTypes {
		_ = elevationtimeslicetype
		for _, _a_extension_9 := range elevationtimeslicetype.Extension {
			stage.ElevationTimeSliceType_Extension_reverseMap[_a_extension_9] = elevationtimeslicetype
		}
	}

	// Compute reverse map for named struct ElevationType
	// insertion point per field
	stage.ElevationType_TimeSlice_reverseMap = make(map[*ElevationTimeSlicePropertyType]*ElevationType)
	for elevationtype := range stage.ElevationTypes {
		_ = elevationtype
		for _, _elevationtimeslicepropertytype := range elevationtype.TimeSlice {
			stage.ElevationType_TimeSlice_reverseMap[_elevationtimeslicepropertytype] = elevationtype
		}
	}

	// Compute reverse map for named struct GlidepathTimeSliceType
	// insertion point per field
	stage.GlidepathTimeSliceType_Authority_reverseMap = make(map[*AuthorityForNavaidEquipmentPropertyType]*GlidepathTimeSliceType)
	for glidepathtimeslicetype := range stage.GlidepathTimeSliceTypes {
		_ = glidepathtimeslicetype
		for _, _authorityfornavaidequipmentpropertytype := range glidepathtimeslicetype.Authority {
			stage.GlidepathTimeSliceType_Authority_reverseMap[_authorityfornavaidequipmentpropertytype] = glidepathtimeslicetype
		}
	}
	stage.GlidepathTimeSliceType_Monitoring_reverseMap = make(map[*NavaidEquipmentMonitoringPropertyType]*GlidepathTimeSliceType)
	for glidepathtimeslicetype := range stage.GlidepathTimeSliceTypes {
		_ = glidepathtimeslicetype
		for _, _navaidequipmentmonitoringpropertytype := range glidepathtimeslicetype.Monitoring {
			stage.GlidepathTimeSliceType_Monitoring_reverseMap[_navaidequipmentmonitoringpropertytype] = glidepathtimeslicetype
		}
	}
	stage.GlidepathTimeSliceType_Availability_reverseMap = make(map[*NavaidOperationalStatusPropertyType]*GlidepathTimeSliceType)
	for glidepathtimeslicetype := range stage.GlidepathTimeSliceTypes {
		_ = glidepathtimeslicetype
		for _, _navaidoperationalstatuspropertytype := range glidepathtimeslicetype.Availability {
			stage.GlidepathTimeSliceType_Availability_reverseMap[_navaidoperationalstatuspropertytype] = glidepathtimeslicetype
		}
	}
	stage.GlidepathTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*GlidepathTimeSliceType)
	for glidepathtimeslicetype := range stage.GlidepathTimeSliceTypes {
		_ = glidepathtimeslicetype
		for _, _notepropertytype := range glidepathtimeslicetype.Annotation {
			stage.GlidepathTimeSliceType_Annotation_reverseMap[_notepropertytype] = glidepathtimeslicetype
		}
	}
	stage.GlidepathTimeSliceType_Extension_reverseMap = make(map[*A_extension_10]*GlidepathTimeSliceType)
	for glidepathtimeslicetype := range stage.GlidepathTimeSliceTypes {
		_ = glidepathtimeslicetype
		for _, _a_extension_10 := range glidepathtimeslicetype.Extension {
			stage.GlidepathTimeSliceType_Extension_reverseMap[_a_extension_10] = glidepathtimeslicetype
		}
	}

	// Compute reverse map for named struct GlidepathType
	// insertion point per field
	stage.GlidepathType_TimeSlice_reverseMap = make(map[*GlidepathTimeSlicePropertyType]*GlidepathType)
	for glidepathtype := range stage.GlidepathTypes {
		_ = glidepathtype
		for _, _glidepathtimeslicepropertytype := range glidepathtype.TimeSlice {
			stage.GlidepathType_TimeSlice_reverseMap[_glidepathtimeslicepropertytype] = glidepathtype
		}
	}

	// Compute reverse map for named struct LinguisticNoteType
	// insertion point per field
	stage.LinguisticNoteType_Extension_reverseMap = make(map[*A_extension_28]*LinguisticNoteType)
	for linguisticnotetype := range stage.LinguisticNoteTypes {
		_ = linguisticnotetype
		for _, _a_extension_28 := range linguisticnotetype.Extension {
			stage.LinguisticNoteType_Extension_reverseMap[_a_extension_28] = linguisticnotetype
		}
	}

	// Compute reverse map for named struct LocalizerTimeSliceType
	// insertion point per field
	stage.LocalizerTimeSliceType_Authority_reverseMap = make(map[*AuthorityForNavaidEquipmentPropertyType]*LocalizerTimeSliceType)
	for localizertimeslicetype := range stage.LocalizerTimeSliceTypes {
		_ = localizertimeslicetype
		for _, _authorityfornavaidequipmentpropertytype := range localizertimeslicetype.Authority {
			stage.LocalizerTimeSliceType_Authority_reverseMap[_authorityfornavaidequipmentpropertytype] = localizertimeslicetype
		}
	}
	stage.LocalizerTimeSliceType_Monitoring_reverseMap = make(map[*NavaidEquipmentMonitoringPropertyType]*LocalizerTimeSliceType)
	for localizertimeslicetype := range stage.LocalizerTimeSliceTypes {
		_ = localizertimeslicetype
		for _, _navaidequipmentmonitoringpropertytype := range localizertimeslicetype.Monitoring {
			stage.LocalizerTimeSliceType_Monitoring_reverseMap[_navaidequipmentmonitoringpropertytype] = localizertimeslicetype
		}
	}
	stage.LocalizerTimeSliceType_Availability_reverseMap = make(map[*NavaidOperationalStatusPropertyType]*LocalizerTimeSliceType)
	for localizertimeslicetype := range stage.LocalizerTimeSliceTypes {
		_ = localizertimeslicetype
		for _, _navaidoperationalstatuspropertytype := range localizertimeslicetype.Availability {
			stage.LocalizerTimeSliceType_Availability_reverseMap[_navaidoperationalstatuspropertytype] = localizertimeslicetype
		}
	}
	stage.LocalizerTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*LocalizerTimeSliceType)
	for localizertimeslicetype := range stage.LocalizerTimeSliceTypes {
		_ = localizertimeslicetype
		for _, _notepropertytype := range localizertimeslicetype.Annotation {
			stage.LocalizerTimeSliceType_Annotation_reverseMap[_notepropertytype] = localizertimeslicetype
		}
	}
	stage.LocalizerTimeSliceType_Extension_reverseMap = make(map[*A_extension_11]*LocalizerTimeSliceType)
	for localizertimeslicetype := range stage.LocalizerTimeSliceTypes {
		_ = localizertimeslicetype
		for _, _a_extension_11 := range localizertimeslicetype.Extension {
			stage.LocalizerTimeSliceType_Extension_reverseMap[_a_extension_11] = localizertimeslicetype
		}
	}

	// Compute reverse map for named struct LocalizerType
	// insertion point per field
	stage.LocalizerType_TimeSlice_reverseMap = make(map[*LocalizerTimeSlicePropertyType]*LocalizerType)
	for localizertype := range stage.LocalizerTypes {
		_ = localizertype
		for _, _localizertimeslicepropertytype := range localizertype.TimeSlice {
			stage.LocalizerType_TimeSlice_reverseMap[_localizertimeslicepropertytype] = localizertype
		}
	}

	// Compute reverse map for named struct MarkerBeaconTimeSliceType
	// insertion point per field
	stage.MarkerBeaconTimeSliceType_Authority_reverseMap = make(map[*AuthorityForNavaidEquipmentPropertyType]*MarkerBeaconTimeSliceType)
	for markerbeacontimeslicetype := range stage.MarkerBeaconTimeSliceTypes {
		_ = markerbeacontimeslicetype
		for _, _authorityfornavaidequipmentpropertytype := range markerbeacontimeslicetype.Authority {
			stage.MarkerBeaconTimeSliceType_Authority_reverseMap[_authorityfornavaidequipmentpropertytype] = markerbeacontimeslicetype
		}
	}
	stage.MarkerBeaconTimeSliceType_Monitoring_reverseMap = make(map[*NavaidEquipmentMonitoringPropertyType]*MarkerBeaconTimeSliceType)
	for markerbeacontimeslicetype := range stage.MarkerBeaconTimeSliceTypes {
		_ = markerbeacontimeslicetype
		for _, _navaidequipmentmonitoringpropertytype := range markerbeacontimeslicetype.Monitoring {
			stage.MarkerBeaconTimeSliceType_Monitoring_reverseMap[_navaidequipmentmonitoringpropertytype] = markerbeacontimeslicetype
		}
	}
	stage.MarkerBeaconTimeSliceType_Availability_reverseMap = make(map[*NavaidOperationalStatusPropertyType]*MarkerBeaconTimeSliceType)
	for markerbeacontimeslicetype := range stage.MarkerBeaconTimeSliceTypes {
		_ = markerbeacontimeslicetype
		for _, _navaidoperationalstatuspropertytype := range markerbeacontimeslicetype.Availability {
			stage.MarkerBeaconTimeSliceType_Availability_reverseMap[_navaidoperationalstatuspropertytype] = markerbeacontimeslicetype
		}
	}
	stage.MarkerBeaconTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*MarkerBeaconTimeSliceType)
	for markerbeacontimeslicetype := range stage.MarkerBeaconTimeSliceTypes {
		_ = markerbeacontimeslicetype
		for _, _notepropertytype := range markerbeacontimeslicetype.Annotation {
			stage.MarkerBeaconTimeSliceType_Annotation_reverseMap[_notepropertytype] = markerbeacontimeslicetype
		}
	}
	stage.MarkerBeaconTimeSliceType_Extension_reverseMap = make(map[*A_extension_12]*MarkerBeaconTimeSliceType)
	for markerbeacontimeslicetype := range stage.MarkerBeaconTimeSliceTypes {
		_ = markerbeacontimeslicetype
		for _, _a_extension_12 := range markerbeacontimeslicetype.Extension {
			stage.MarkerBeaconTimeSliceType_Extension_reverseMap[_a_extension_12] = markerbeacontimeslicetype
		}
	}

	// Compute reverse map for named struct MarkerBeaconType
	// insertion point per field
	stage.MarkerBeaconType_TimeSlice_reverseMap = make(map[*MarkerBeaconTimeSlicePropertyType]*MarkerBeaconType)
	for markerbeacontype := range stage.MarkerBeaconTypes {
		_ = markerbeacontype
		for _, _markerbeacontimeslicepropertytype := range markerbeacontype.TimeSlice {
			stage.MarkerBeaconType_TimeSlice_reverseMap[_markerbeacontimeslicepropertytype] = markerbeacontype
		}
	}

	// Compute reverse map for named struct NDBTimeSliceType
	// insertion point per field
	stage.NDBTimeSliceType_Authority_reverseMap = make(map[*AuthorityForNavaidEquipmentPropertyType]*NDBTimeSliceType)
	for ndbtimeslicetype := range stage.NDBTimeSliceTypes {
		_ = ndbtimeslicetype
		for _, _authorityfornavaidequipmentpropertytype := range ndbtimeslicetype.Authority {
			stage.NDBTimeSliceType_Authority_reverseMap[_authorityfornavaidequipmentpropertytype] = ndbtimeslicetype
		}
	}
	stage.NDBTimeSliceType_Monitoring_reverseMap = make(map[*NavaidEquipmentMonitoringPropertyType]*NDBTimeSliceType)
	for ndbtimeslicetype := range stage.NDBTimeSliceTypes {
		_ = ndbtimeslicetype
		for _, _navaidequipmentmonitoringpropertytype := range ndbtimeslicetype.Monitoring {
			stage.NDBTimeSliceType_Monitoring_reverseMap[_navaidequipmentmonitoringpropertytype] = ndbtimeslicetype
		}
	}
	stage.NDBTimeSliceType_Availability_reverseMap = make(map[*NavaidOperationalStatusPropertyType]*NDBTimeSliceType)
	for ndbtimeslicetype := range stage.NDBTimeSliceTypes {
		_ = ndbtimeslicetype
		for _, _navaidoperationalstatuspropertytype := range ndbtimeslicetype.Availability {
			stage.NDBTimeSliceType_Availability_reverseMap[_navaidoperationalstatuspropertytype] = ndbtimeslicetype
		}
	}
	stage.NDBTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*NDBTimeSliceType)
	for ndbtimeslicetype := range stage.NDBTimeSliceTypes {
		_ = ndbtimeslicetype
		for _, _notepropertytype := range ndbtimeslicetype.Annotation {
			stage.NDBTimeSliceType_Annotation_reverseMap[_notepropertytype] = ndbtimeslicetype
		}
	}
	stage.NDBTimeSliceType_Extension_reverseMap = make(map[*A_extension_17]*NDBTimeSliceType)
	for ndbtimeslicetype := range stage.NDBTimeSliceTypes {
		_ = ndbtimeslicetype
		for _, _a_extension_17 := range ndbtimeslicetype.Extension {
			stage.NDBTimeSliceType_Extension_reverseMap[_a_extension_17] = ndbtimeslicetype
		}
	}

	// Compute reverse map for named struct NDBType
	// insertion point per field
	stage.NDBType_TimeSlice_reverseMap = make(map[*NDBTimeSlicePropertyType]*NDBType)
	for ndbtype := range stage.NDBTypes {
		_ = ndbtype
		for _, _ndbtimeslicepropertytype := range ndbtype.TimeSlice {
			stage.NDBType_TimeSlice_reverseMap[_ndbtimeslicepropertytype] = ndbtype
		}
	}

	// Compute reverse map for named struct NavaidComponentType
	// insertion point per field
	stage.NavaidComponentType_Annotation_reverseMap = make(map[*NotePropertyType]*NavaidComponentType)
	for navaidcomponenttype := range stage.NavaidComponentTypes {
		_ = navaidcomponenttype
		for _, _notepropertytype := range navaidcomponenttype.Annotation {
			stage.NavaidComponentType_Annotation_reverseMap[_notepropertytype] = navaidcomponenttype
		}
	}
	stage.NavaidComponentType_Extension_reverseMap = make(map[*A_extension_14]*NavaidComponentType)
	for navaidcomponenttype := range stage.NavaidComponentTypes {
		_ = navaidcomponenttype
		for _, _a_extension_14 := range navaidcomponenttype.Extension {
			stage.NavaidComponentType_Extension_reverseMap[_a_extension_14] = navaidcomponenttype
		}
	}

	// Compute reverse map for named struct NavaidEquipmentMonitoringType
	// insertion point per field
	stage.NavaidEquipmentMonitoringType_TimeInterval_reverseMap = make(map[*TimesheetPropertyType]*NavaidEquipmentMonitoringType)
	for navaidequipmentmonitoringtype := range stage.NavaidEquipmentMonitoringTypes {
		_ = navaidequipmentmonitoringtype
		for _, _timesheetpropertytype := range navaidequipmentmonitoringtype.TimeInterval {
			stage.NavaidEquipmentMonitoringType_TimeInterval_reverseMap[_timesheetpropertytype] = navaidequipmentmonitoringtype
		}
	}
	stage.NavaidEquipmentMonitoringType_Annotation_reverseMap = make(map[*NotePropertyType]*NavaidEquipmentMonitoringType)
	for navaidequipmentmonitoringtype := range stage.NavaidEquipmentMonitoringTypes {
		_ = navaidequipmentmonitoringtype
		for _, _notepropertytype := range navaidequipmentmonitoringtype.Annotation {
			stage.NavaidEquipmentMonitoringType_Annotation_reverseMap[_notepropertytype] = navaidequipmentmonitoringtype
		}
	}
	stage.NavaidEquipmentMonitoringType_SpecialDateAuthority_reverseMap = make(map[*OrganisationAuthorityPropertyType]*NavaidEquipmentMonitoringType)
	for navaidequipmentmonitoringtype := range stage.NavaidEquipmentMonitoringTypes {
		_ = navaidequipmentmonitoringtype
		for _, _organisationauthoritypropertytype := range navaidequipmentmonitoringtype.SpecialDateAuthority {
			stage.NavaidEquipmentMonitoringType_SpecialDateAuthority_reverseMap[_organisationauthoritypropertytype] = navaidequipmentmonitoringtype
		}
	}
	stage.NavaidEquipmentMonitoringType_Extension_reverseMap = make(map[*A_extension_15]*NavaidEquipmentMonitoringType)
	for navaidequipmentmonitoringtype := range stage.NavaidEquipmentMonitoringTypes {
		_ = navaidequipmentmonitoringtype
		for _, _a_extension_15 := range navaidequipmentmonitoringtype.Extension {
			stage.NavaidEquipmentMonitoringType_Extension_reverseMap[_a_extension_15] = navaidequipmentmonitoringtype
		}
	}

	// Compute reverse map for named struct NavaidOperationalStatusType
	// insertion point per field
	stage.NavaidOperationalStatusType_TimeInterval_reverseMap = make(map[*TimesheetPropertyType]*NavaidOperationalStatusType)
	for navaidoperationalstatustype := range stage.NavaidOperationalStatusTypes {
		_ = navaidoperationalstatustype
		for _, _timesheetpropertytype := range navaidoperationalstatustype.TimeInterval {
			stage.NavaidOperationalStatusType_TimeInterval_reverseMap[_timesheetpropertytype] = navaidoperationalstatustype
		}
	}
	stage.NavaidOperationalStatusType_Annotation_reverseMap = make(map[*NotePropertyType]*NavaidOperationalStatusType)
	for navaidoperationalstatustype := range stage.NavaidOperationalStatusTypes {
		_ = navaidoperationalstatustype
		for _, _notepropertytype := range navaidoperationalstatustype.Annotation {
			stage.NavaidOperationalStatusType_Annotation_reverseMap[_notepropertytype] = navaidoperationalstatustype
		}
	}
	stage.NavaidOperationalStatusType_SpecialDateAuthority_reverseMap = make(map[*OrganisationAuthorityPropertyType]*NavaidOperationalStatusType)
	for navaidoperationalstatustype := range stage.NavaidOperationalStatusTypes {
		_ = navaidoperationalstatustype
		for _, _organisationauthoritypropertytype := range navaidoperationalstatustype.SpecialDateAuthority {
			stage.NavaidOperationalStatusType_SpecialDateAuthority_reverseMap[_organisationauthoritypropertytype] = navaidoperationalstatustype
		}
	}
	stage.NavaidOperationalStatusType_Extension_reverseMap = make(map[*A_extension_16]*NavaidOperationalStatusType)
	for navaidoperationalstatustype := range stage.NavaidOperationalStatusTypes {
		_ = navaidoperationalstatustype
		for _, _a_extension_16 := range navaidoperationalstatustype.Extension {
			stage.NavaidOperationalStatusType_Extension_reverseMap[_a_extension_16] = navaidoperationalstatustype
		}
	}

	// Compute reverse map for named struct NavaidTimeSliceType
	// insertion point per field
	stage.NavaidTimeSliceType_TouchDownLiftOff_reverseMap = make(map[*TouchDownLiftOffPropertyType]*NavaidTimeSliceType)
	for navaidtimeslicetype := range stage.NavaidTimeSliceTypes {
		_ = navaidtimeslicetype
		for _, _touchdownliftoffpropertytype := range navaidtimeslicetype.TouchDownLiftOff {
			stage.NavaidTimeSliceType_TouchDownLiftOff_reverseMap[_touchdownliftoffpropertytype] = navaidtimeslicetype
		}
	}
	stage.NavaidTimeSliceType_NavaidEquipment_reverseMap = make(map[*NavaidComponentPropertyType]*NavaidTimeSliceType)
	for navaidtimeslicetype := range stage.NavaidTimeSliceTypes {
		_ = navaidtimeslicetype
		for _, _navaidcomponentpropertytype := range navaidtimeslicetype.NavaidEquipment {
			stage.NavaidTimeSliceType_NavaidEquipment_reverseMap[_navaidcomponentpropertytype] = navaidtimeslicetype
		}
	}
	stage.NavaidTimeSliceType_RunwayDirection_reverseMap = make(map[*RunwayDirectionPropertyType]*NavaidTimeSliceType)
	for navaidtimeslicetype := range stage.NavaidTimeSliceTypes {
		_ = navaidtimeslicetype
		for _, _runwaydirectionpropertytype := range navaidtimeslicetype.RunwayDirection {
			stage.NavaidTimeSliceType_RunwayDirection_reverseMap[_runwaydirectionpropertytype] = navaidtimeslicetype
		}
	}
	stage.NavaidTimeSliceType_ServedAirport_reverseMap = make(map[*AirportHeliportPropertyType]*NavaidTimeSliceType)
	for navaidtimeslicetype := range stage.NavaidTimeSliceTypes {
		_ = navaidtimeslicetype
		for _, _airportheliportpropertytype := range navaidtimeslicetype.ServedAirport {
			stage.NavaidTimeSliceType_ServedAirport_reverseMap[_airportheliportpropertytype] = navaidtimeslicetype
		}
	}
	stage.NavaidTimeSliceType_Availability_reverseMap = make(map[*NavaidOperationalStatusPropertyType]*NavaidTimeSliceType)
	for navaidtimeslicetype := range stage.NavaidTimeSliceTypes {
		_ = navaidtimeslicetype
		for _, _navaidoperationalstatuspropertytype := range navaidtimeslicetype.Availability {
			stage.NavaidTimeSliceType_Availability_reverseMap[_navaidoperationalstatuspropertytype] = navaidtimeslicetype
		}
	}
	stage.NavaidTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*NavaidTimeSliceType)
	for navaidtimeslicetype := range stage.NavaidTimeSliceTypes {
		_ = navaidtimeslicetype
		for _, _notepropertytype := range navaidtimeslicetype.Annotation {
			stage.NavaidTimeSliceType_Annotation_reverseMap[_notepropertytype] = navaidtimeslicetype
		}
	}
	stage.NavaidTimeSliceType_Extension_reverseMap = make(map[*A_extension_13]*NavaidTimeSliceType)
	for navaidtimeslicetype := range stage.NavaidTimeSliceTypes {
		_ = navaidtimeslicetype
		for _, _a_extension_13 := range navaidtimeslicetype.Extension {
			stage.NavaidTimeSliceType_Extension_reverseMap[_a_extension_13] = navaidtimeslicetype
		}
	}

	// Compute reverse map for named struct NavaidType
	// insertion point per field
	stage.NavaidType_TimeSlice_reverseMap = make(map[*NavaidTimeSlicePropertyType]*NavaidType)
	for navaidtype := range stage.NavaidTypes {
		_ = navaidtype
		for _, _navaidtimeslicepropertytype := range navaidtype.TimeSlice {
			stage.NavaidType_TimeSlice_reverseMap[_navaidtimeslicepropertytype] = navaidtype
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
	stage.NoteType_Extension_reverseMap = make(map[*A_extension_29]*NoteType)
	for notetype := range stage.NoteTypes {
		_ = notetype
		for _, _a_extension_29 := range notetype.Extension {
			stage.NoteType_Extension_reverseMap[_a_extension_29] = notetype
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

	// Compute reverse map for named struct PrecisionApproachRadarTimeSliceType
	// insertion point per field
	stage.PrecisionApproachRadarTimeSliceType_Contact_reverseMap = make(map[*ContactInformationPropertyType]*PrecisionApproachRadarTimeSliceType)
	for precisionapproachradartimeslicetype := range stage.PrecisionApproachRadarTimeSliceTypes {
		_ = precisionapproachradartimeslicetype
		for _, _contactinformationpropertytype := range precisionapproachradartimeslicetype.Contact {
			stage.PrecisionApproachRadarTimeSliceType_Contact_reverseMap[_contactinformationpropertytype] = precisionapproachradartimeslicetype
		}
	}
	stage.PrecisionApproachRadarTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*PrecisionApproachRadarTimeSliceType)
	for precisionapproachradartimeslicetype := range stage.PrecisionApproachRadarTimeSliceTypes {
		_ = precisionapproachradartimeslicetype
		for _, _notepropertytype := range precisionapproachradartimeslicetype.Annotation {
			stage.PrecisionApproachRadarTimeSliceType_Annotation_reverseMap[_notepropertytype] = precisionapproachradartimeslicetype
		}
	}
	stage.PrecisionApproachRadarTimeSliceType_Reflector_reverseMap = make(map[*ReflectorPropertyType]*PrecisionApproachRadarTimeSliceType)
	for precisionapproachradartimeslicetype := range stage.PrecisionApproachRadarTimeSliceTypes {
		_ = precisionapproachradartimeslicetype
		for _, _reflectorpropertytype := range precisionapproachradartimeslicetype.Reflector {
			stage.PrecisionApproachRadarTimeSliceType_Reflector_reverseMap[_reflectorpropertytype] = precisionapproachradartimeslicetype
		}
	}
	stage.PrecisionApproachRadarTimeSliceType_Extension_reverseMap = make(map[*A_extension_35]*PrecisionApproachRadarTimeSliceType)
	for precisionapproachradartimeslicetype := range stage.PrecisionApproachRadarTimeSliceTypes {
		_ = precisionapproachradartimeslicetype
		for _, _a_extension_35 := range precisionapproachradartimeslicetype.Extension {
			stage.PrecisionApproachRadarTimeSliceType_Extension_reverseMap[_a_extension_35] = precisionapproachradartimeslicetype
		}
	}

	// Compute reverse map for named struct PrecisionApproachRadarType
	// insertion point per field
	stage.PrecisionApproachRadarType_TimeSlice_reverseMap = make(map[*PrecisionApproachRadarTimeSlicePropertyType]*PrecisionApproachRadarType)
	for precisionapproachradartype := range stage.PrecisionApproachRadarTypes {
		_ = precisionapproachradartype
		for _, _precisionapproachradartimeslicepropertytype := range precisionapproachradartype.TimeSlice {
			stage.PrecisionApproachRadarType_TimeSlice_reverseMap[_precisionapproachradartimeslicepropertytype] = precisionapproachradartype
		}
	}

	// Compute reverse map for named struct PrimarySurveillanceRadarTimeSliceType
	// insertion point per field
	stage.PrimarySurveillanceRadarTimeSliceType_Contact_reverseMap = make(map[*ContactInformationPropertyType]*PrimarySurveillanceRadarTimeSliceType)
	for primarysurveillanceradartimeslicetype := range stage.PrimarySurveillanceRadarTimeSliceTypes {
		_ = primarysurveillanceradartimeslicetype
		for _, _contactinformationpropertytype := range primarysurveillanceradartimeslicetype.Contact {
			stage.PrimarySurveillanceRadarTimeSliceType_Contact_reverseMap[_contactinformationpropertytype] = primarysurveillanceradartimeslicetype
		}
	}
	stage.PrimarySurveillanceRadarTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*PrimarySurveillanceRadarTimeSliceType)
	for primarysurveillanceradartimeslicetype := range stage.PrimarySurveillanceRadarTimeSliceTypes {
		_ = primarysurveillanceradartimeslicetype
		for _, _notepropertytype := range primarysurveillanceradartimeslicetype.Annotation {
			stage.PrimarySurveillanceRadarTimeSliceType_Annotation_reverseMap[_notepropertytype] = primarysurveillanceradartimeslicetype
		}
	}
	stage.PrimarySurveillanceRadarTimeSliceType_GroundStation_reverseMap = make(map[*SurveillanceGroundStationPropertyType]*PrimarySurveillanceRadarTimeSliceType)
	for primarysurveillanceradartimeslicetype := range stage.PrimarySurveillanceRadarTimeSliceTypes {
		_ = primarysurveillanceradartimeslicetype
		for _, _surveillancegroundstationpropertytype := range primarysurveillanceradartimeslicetype.GroundStation {
			stage.PrimarySurveillanceRadarTimeSliceType_GroundStation_reverseMap[_surveillancegroundstationpropertytype] = primarysurveillanceradartimeslicetype
		}
	}
	stage.PrimarySurveillanceRadarTimeSliceType_Extension_reverseMap = make(map[*A_extension_36]*PrimarySurveillanceRadarTimeSliceType)
	for primarysurveillanceradartimeslicetype := range stage.PrimarySurveillanceRadarTimeSliceTypes {
		_ = primarysurveillanceradartimeslicetype
		for _, _a_extension_36 := range primarysurveillanceradartimeslicetype.Extension {
			stage.PrimarySurveillanceRadarTimeSliceType_Extension_reverseMap[_a_extension_36] = primarysurveillanceradartimeslicetype
		}
	}

	// Compute reverse map for named struct PrimarySurveillanceRadarType
	// insertion point per field
	stage.PrimarySurveillanceRadarType_TimeSlice_reverseMap = make(map[*PrimarySurveillanceRadarTimeSlicePropertyType]*PrimarySurveillanceRadarType)
	for primarysurveillanceradartype := range stage.PrimarySurveillanceRadarTypes {
		_ = primarysurveillanceradartype
		for _, _primarysurveillanceradartimeslicepropertytype := range primarysurveillanceradartype.TimeSlice {
			stage.PrimarySurveillanceRadarType_TimeSlice_reverseMap[_primarysurveillanceradartimeslicepropertytype] = primarysurveillanceradartype
		}
	}

	// Compute reverse map for named struct RadarComponentType
	// insertion point per field
	stage.RadarComponentType_Annotation_reverseMap = make(map[*NotePropertyType]*RadarComponentType)
	for radarcomponenttype := range stage.RadarComponentTypes {
		_ = radarcomponenttype
		for _, _notepropertytype := range radarcomponenttype.Annotation {
			stage.RadarComponentType_Annotation_reverseMap[_notepropertytype] = radarcomponenttype
		}
	}
	stage.RadarComponentType_Extension_reverseMap = make(map[*A_extension_37]*RadarComponentType)
	for radarcomponenttype := range stage.RadarComponentTypes {
		_ = radarcomponenttype
		for _, _a_extension_37 := range radarcomponenttype.Extension {
			stage.RadarComponentType_Extension_reverseMap[_a_extension_37] = radarcomponenttype
		}
	}

	// Compute reverse map for named struct RadarSystemTimeSliceType
	// insertion point per field
	stage.RadarSystemTimeSliceType_RadarEquipment_reverseMap = make(map[*RadarComponentPropertyType]*RadarSystemTimeSliceType)
	for radarsystemtimeslicetype := range stage.RadarSystemTimeSliceTypes {
		_ = radarsystemtimeslicetype
		for _, _radarcomponentpropertytype := range radarsystemtimeslicetype.RadarEquipment {
			stage.RadarSystemTimeSliceType_RadarEquipment_reverseMap[_radarcomponentpropertytype] = radarsystemtimeslicetype
		}
	}
	stage.RadarSystemTimeSliceType_Office_reverseMap = make(map[*OrganisationAuthorityPropertyType]*RadarSystemTimeSliceType)
	for radarsystemtimeslicetype := range stage.RadarSystemTimeSliceTypes {
		_ = radarsystemtimeslicetype
		for _, _organisationauthoritypropertytype := range radarsystemtimeslicetype.Office {
			stage.RadarSystemTimeSliceType_Office_reverseMap[_organisationauthoritypropertytype] = radarsystemtimeslicetype
		}
	}
	stage.RadarSystemTimeSliceType_PARRunway_reverseMap = make(map[*RunwayPropertyType]*RadarSystemTimeSliceType)
	for radarsystemtimeslicetype := range stage.RadarSystemTimeSliceTypes {
		_ = radarsystemtimeslicetype
		for _, _runwaypropertytype := range radarsystemtimeslicetype.PARRunway {
			stage.RadarSystemTimeSliceType_PARRunway_reverseMap[_runwaypropertytype] = radarsystemtimeslicetype
		}
	}
	stage.RadarSystemTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*RadarSystemTimeSliceType)
	for radarsystemtimeslicetype := range stage.RadarSystemTimeSliceTypes {
		_ = radarsystemtimeslicetype
		for _, _notepropertytype := range radarsystemtimeslicetype.Annotation {
			stage.RadarSystemTimeSliceType_Annotation_reverseMap[_notepropertytype] = radarsystemtimeslicetype
		}
	}
	stage.RadarSystemTimeSliceType_Extension_reverseMap = make(map[*A_extension_38]*RadarSystemTimeSliceType)
	for radarsystemtimeslicetype := range stage.RadarSystemTimeSliceTypes {
		_ = radarsystemtimeslicetype
		for _, _a_extension_38 := range radarsystemtimeslicetype.Extension {
			stage.RadarSystemTimeSliceType_Extension_reverseMap[_a_extension_38] = radarsystemtimeslicetype
		}
	}

	// Compute reverse map for named struct RadarSystemType
	// insertion point per field
	stage.RadarSystemType_TimeSlice_reverseMap = make(map[*RadarSystemTimeSlicePropertyType]*RadarSystemType)
	for radarsystemtype := range stage.RadarSystemTypes {
		_ = radarsystemtype
		for _, _radarsystemtimeslicepropertytype := range radarsystemtype.TimeSlice {
			stage.RadarSystemType_TimeSlice_reverseMap[_radarsystemtimeslicepropertytype] = radarsystemtype
		}
	}

	// Compute reverse map for named struct ReflectorType
	// insertion point per field
	stage.ReflectorType_Annotation_reverseMap = make(map[*NotePropertyType]*ReflectorType)
	for reflectortype := range stage.ReflectorTypes {
		_ = reflectortype
		for _, _notepropertytype := range reflectortype.Annotation {
			stage.ReflectorType_Annotation_reverseMap[_notepropertytype] = reflectortype
		}
	}
	stage.ReflectorType_Extension_reverseMap = make(map[*A_extension_39]*ReflectorType)
	for reflectortype := range stage.ReflectorTypes {
		_ = reflectortype
		for _, _a_extension_39 := range reflectortype.Extension {
			stage.ReflectorType_Extension_reverseMap[_a_extension_39] = reflectortype
		}
	}

	// Compute reverse map for named struct SDFTimeSliceType
	// insertion point per field
	stage.SDFTimeSliceType_Authority_reverseMap = make(map[*AuthorityForNavaidEquipmentPropertyType]*SDFTimeSliceType)
	for sdftimeslicetype := range stage.SDFTimeSliceTypes {
		_ = sdftimeslicetype
		for _, _authorityfornavaidequipmentpropertytype := range sdftimeslicetype.Authority {
			stage.SDFTimeSliceType_Authority_reverseMap[_authorityfornavaidequipmentpropertytype] = sdftimeslicetype
		}
	}
	stage.SDFTimeSliceType_Monitoring_reverseMap = make(map[*NavaidEquipmentMonitoringPropertyType]*SDFTimeSliceType)
	for sdftimeslicetype := range stage.SDFTimeSliceTypes {
		_ = sdftimeslicetype
		for _, _navaidequipmentmonitoringpropertytype := range sdftimeslicetype.Monitoring {
			stage.SDFTimeSliceType_Monitoring_reverseMap[_navaidequipmentmonitoringpropertytype] = sdftimeslicetype
		}
	}
	stage.SDFTimeSliceType_Availability_reverseMap = make(map[*NavaidOperationalStatusPropertyType]*SDFTimeSliceType)
	for sdftimeslicetype := range stage.SDFTimeSliceTypes {
		_ = sdftimeslicetype
		for _, _navaidoperationalstatuspropertytype := range sdftimeslicetype.Availability {
			stage.SDFTimeSliceType_Availability_reverseMap[_navaidoperationalstatuspropertytype] = sdftimeslicetype
		}
	}
	stage.SDFTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*SDFTimeSliceType)
	for sdftimeslicetype := range stage.SDFTimeSliceTypes {
		_ = sdftimeslicetype
		for _, _notepropertytype := range sdftimeslicetype.Annotation {
			stage.SDFTimeSliceType_Annotation_reverseMap[_notepropertytype] = sdftimeslicetype
		}
	}
	stage.SDFTimeSliceType_Extension_reverseMap = make(map[*A_extension_18]*SDFTimeSliceType)
	for sdftimeslicetype := range stage.SDFTimeSliceTypes {
		_ = sdftimeslicetype
		for _, _a_extension_18 := range sdftimeslicetype.Extension {
			stage.SDFTimeSliceType_Extension_reverseMap[_a_extension_18] = sdftimeslicetype
		}
	}

	// Compute reverse map for named struct SDFType
	// insertion point per field
	stage.SDFType_TimeSlice_reverseMap = make(map[*SDFTimeSlicePropertyType]*SDFType)
	for sdftype := range stage.SDFTypes {
		_ = sdftype
		for _, _sdftimeslicepropertytype := range sdftype.TimeSlice {
			stage.SDFType_TimeSlice_reverseMap[_sdftimeslicepropertytype] = sdftype
		}
	}

	// Compute reverse map for named struct SecondarySurveillanceRadarTimeSliceType
	// insertion point per field
	stage.SecondarySurveillanceRadarTimeSliceType_Contact_reverseMap = make(map[*ContactInformationPropertyType]*SecondarySurveillanceRadarTimeSliceType)
	for secondarysurveillanceradartimeslicetype := range stage.SecondarySurveillanceRadarTimeSliceTypes {
		_ = secondarysurveillanceradartimeslicetype
		for _, _contactinformationpropertytype := range secondarysurveillanceradartimeslicetype.Contact {
			stage.SecondarySurveillanceRadarTimeSliceType_Contact_reverseMap[_contactinformationpropertytype] = secondarysurveillanceradartimeslicetype
		}
	}
	stage.SecondarySurveillanceRadarTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*SecondarySurveillanceRadarTimeSliceType)
	for secondarysurveillanceradartimeslicetype := range stage.SecondarySurveillanceRadarTimeSliceTypes {
		_ = secondarysurveillanceradartimeslicetype
		for _, _notepropertytype := range secondarysurveillanceradartimeslicetype.Annotation {
			stage.SecondarySurveillanceRadarTimeSliceType_Annotation_reverseMap[_notepropertytype] = secondarysurveillanceradartimeslicetype
		}
	}
	stage.SecondarySurveillanceRadarTimeSliceType_GroundStation_reverseMap = make(map[*SurveillanceGroundStationPropertyType]*SecondarySurveillanceRadarTimeSliceType)
	for secondarysurveillanceradartimeslicetype := range stage.SecondarySurveillanceRadarTimeSliceTypes {
		_ = secondarysurveillanceradartimeslicetype
		for _, _surveillancegroundstationpropertytype := range secondarysurveillanceradartimeslicetype.GroundStation {
			stage.SecondarySurveillanceRadarTimeSliceType_GroundStation_reverseMap[_surveillancegroundstationpropertytype] = secondarysurveillanceradartimeslicetype
		}
	}
	stage.SecondarySurveillanceRadarTimeSliceType_Extension_reverseMap = make(map[*A_extension_40]*SecondarySurveillanceRadarTimeSliceType)
	for secondarysurveillanceradartimeslicetype := range stage.SecondarySurveillanceRadarTimeSliceTypes {
		_ = secondarysurveillanceradartimeslicetype
		for _, _a_extension_40 := range secondarysurveillanceradartimeslicetype.Extension {
			stage.SecondarySurveillanceRadarTimeSliceType_Extension_reverseMap[_a_extension_40] = secondarysurveillanceradartimeslicetype
		}
	}

	// Compute reverse map for named struct SecondarySurveillanceRadarType
	// insertion point per field
	stage.SecondarySurveillanceRadarType_TimeSlice_reverseMap = make(map[*SecondarySurveillanceRadarTimeSlicePropertyType]*SecondarySurveillanceRadarType)
	for secondarysurveillanceradartype := range stage.SecondarySurveillanceRadarTypes {
		_ = secondarysurveillanceradartype
		for _, _secondarysurveillanceradartimeslicepropertytype := range secondarysurveillanceradartype.TimeSlice {
			stage.SecondarySurveillanceRadarType_TimeSlice_reverseMap[_secondarysurveillanceradartimeslicepropertytype] = secondarysurveillanceradartype
		}
	}

	// Compute reverse map for named struct SignificantPointInAirspaceTimeSliceType
	// insertion point per field
	stage.SignificantPointInAirspaceTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*SignificantPointInAirspaceTimeSliceType)
	for significantpointinairspacetimeslicetype := range stage.SignificantPointInAirspaceTimeSliceTypes {
		_ = significantpointinairspacetimeslicetype
		for _, _notepropertytype := range significantpointinairspacetimeslicetype.Annotation {
			stage.SignificantPointInAirspaceTimeSliceType_Annotation_reverseMap[_notepropertytype] = significantpointinairspacetimeslicetype
		}
	}
	stage.SignificantPointInAirspaceTimeSliceType_Extension_reverseMap = make(map[*A_extension_27]*SignificantPointInAirspaceTimeSliceType)
	for significantpointinairspacetimeslicetype := range stage.SignificantPointInAirspaceTimeSliceTypes {
		_ = significantpointinairspacetimeslicetype
		for _, _a_extension_27 := range significantpointinairspacetimeslicetype.Extension {
			stage.SignificantPointInAirspaceTimeSliceType_Extension_reverseMap[_a_extension_27] = significantpointinairspacetimeslicetype
		}
	}

	// Compute reverse map for named struct SignificantPointInAirspaceType
	// insertion point per field
	stage.SignificantPointInAirspaceType_TimeSlice_reverseMap = make(map[*SignificantPointInAirspaceTimeSlicePropertyType]*SignificantPointInAirspaceType)
	for significantpointinairspacetype := range stage.SignificantPointInAirspaceTypes {
		_ = significantpointinairspacetype
		for _, _significantpointinairspacetimeslicepropertytype := range significantpointinairspacetype.TimeSlice {
			stage.SignificantPointInAirspaceType_TimeSlice_reverseMap[_significantpointinairspacetimeslicepropertytype] = significantpointinairspacetype
		}
	}

	// Compute reverse map for named struct SpecialNavigationStationStatusType
	// insertion point per field
	stage.SpecialNavigationStationStatusType_TimeInterval_reverseMap = make(map[*TimesheetPropertyType]*SpecialNavigationStationStatusType)
	for specialnavigationstationstatustype := range stage.SpecialNavigationStationStatusTypes {
		_ = specialnavigationstationstatustype
		for _, _timesheetpropertytype := range specialnavigationstationstatustype.TimeInterval {
			stage.SpecialNavigationStationStatusType_TimeInterval_reverseMap[_timesheetpropertytype] = specialnavigationstationstatustype
		}
	}
	stage.SpecialNavigationStationStatusType_Annotation_reverseMap = make(map[*NotePropertyType]*SpecialNavigationStationStatusType)
	for specialnavigationstationstatustype := range stage.SpecialNavigationStationStatusTypes {
		_ = specialnavigationstationstatustype
		for _, _notepropertytype := range specialnavigationstationstatustype.Annotation {
			stage.SpecialNavigationStationStatusType_Annotation_reverseMap[_notepropertytype] = specialnavigationstationstatustype
		}
	}
	stage.SpecialNavigationStationStatusType_SpecialDateAuthority_reverseMap = make(map[*OrganisationAuthorityPropertyType]*SpecialNavigationStationStatusType)
	for specialnavigationstationstatustype := range stage.SpecialNavigationStationStatusTypes {
		_ = specialnavigationstationstatustype
		for _, _organisationauthoritypropertytype := range specialnavigationstationstatustype.SpecialDateAuthority {
			stage.SpecialNavigationStationStatusType_SpecialDateAuthority_reverseMap[_organisationauthoritypropertytype] = specialnavigationstationstatustype
		}
	}
	stage.SpecialNavigationStationStatusType_Extension_reverseMap = make(map[*A_extension_20]*SpecialNavigationStationStatusType)
	for specialnavigationstationstatustype := range stage.SpecialNavigationStationStatusTypes {
		_ = specialnavigationstationstatustype
		for _, _a_extension_20 := range specialnavigationstationstatustype.Extension {
			stage.SpecialNavigationStationStatusType_Extension_reverseMap[_a_extension_20] = specialnavigationstationstatustype
		}
	}

	// Compute reverse map for named struct SpecialNavigationStationTimeSliceType
	// insertion point per field
	stage.SpecialNavigationStationTimeSliceType_Availability_reverseMap = make(map[*SpecialNavigationStationStatusPropertyType]*SpecialNavigationStationTimeSliceType)
	for specialnavigationstationtimeslicetype := range stage.SpecialNavigationStationTimeSliceTypes {
		_ = specialnavigationstationtimeslicetype
		for _, _specialnavigationstationstatuspropertytype := range specialnavigationstationtimeslicetype.Availability {
			stage.SpecialNavigationStationTimeSliceType_Availability_reverseMap[_specialnavigationstationstatuspropertytype] = specialnavigationstationtimeslicetype
		}
	}
	stage.SpecialNavigationStationTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*SpecialNavigationStationTimeSliceType)
	for specialnavigationstationtimeslicetype := range stage.SpecialNavigationStationTimeSliceTypes {
		_ = specialnavigationstationtimeslicetype
		for _, _notepropertytype := range specialnavigationstationtimeslicetype.Annotation {
			stage.SpecialNavigationStationTimeSliceType_Annotation_reverseMap[_notepropertytype] = specialnavigationstationtimeslicetype
		}
	}
	stage.SpecialNavigationStationTimeSliceType_Extension_reverseMap = make(map[*A_extension_19]*SpecialNavigationStationTimeSliceType)
	for specialnavigationstationtimeslicetype := range stage.SpecialNavigationStationTimeSliceTypes {
		_ = specialnavigationstationtimeslicetype
		for _, _a_extension_19 := range specialnavigationstationtimeslicetype.Extension {
			stage.SpecialNavigationStationTimeSliceType_Extension_reverseMap[_a_extension_19] = specialnavigationstationtimeslicetype
		}
	}

	// Compute reverse map for named struct SpecialNavigationStationType
	// insertion point per field
	stage.SpecialNavigationStationType_TimeSlice_reverseMap = make(map[*SpecialNavigationStationTimeSlicePropertyType]*SpecialNavigationStationType)
	for specialnavigationstationtype := range stage.SpecialNavigationStationTypes {
		_ = specialnavigationstationtype
		for _, _specialnavigationstationtimeslicepropertytype := range specialnavigationstationtype.TimeSlice {
			stage.SpecialNavigationStationType_TimeSlice_reverseMap[_specialnavigationstationtimeslicepropertytype] = specialnavigationstationtype
		}
	}

	// Compute reverse map for named struct SpecialNavigationSystemTimeSliceType
	// insertion point per field
	stage.SpecialNavigationSystemTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*SpecialNavigationSystemTimeSliceType)
	for specialnavigationsystemtimeslicetype := range stage.SpecialNavigationSystemTimeSliceTypes {
		_ = specialnavigationsystemtimeslicetype
		for _, _notepropertytype := range specialnavigationsystemtimeslicetype.Annotation {
			stage.SpecialNavigationSystemTimeSliceType_Annotation_reverseMap[_notepropertytype] = specialnavigationsystemtimeslicetype
		}
	}
	stage.SpecialNavigationSystemTimeSliceType_Extension_reverseMap = make(map[*A_extension_21]*SpecialNavigationSystemTimeSliceType)
	for specialnavigationsystemtimeslicetype := range stage.SpecialNavigationSystemTimeSliceTypes {
		_ = specialnavigationsystemtimeslicetype
		for _, _a_extension_21 := range specialnavigationsystemtimeslicetype.Extension {
			stage.SpecialNavigationSystemTimeSliceType_Extension_reverseMap[_a_extension_21] = specialnavigationsystemtimeslicetype
		}
	}

	// Compute reverse map for named struct SpecialNavigationSystemType
	// insertion point per field
	stage.SpecialNavigationSystemType_TimeSlice_reverseMap = make(map[*SpecialNavigationSystemTimeSlicePropertyType]*SpecialNavigationSystemType)
	for specialnavigationsystemtype := range stage.SpecialNavigationSystemTypes {
		_ = specialnavigationsystemtype
		for _, _specialnavigationsystemtimeslicepropertytype := range specialnavigationsystemtype.TimeSlice {
			stage.SpecialNavigationSystemType_TimeSlice_reverseMap[_specialnavigationsystemtimeslicepropertytype] = specialnavigationsystemtype
		}
	}

	// Compute reverse map for named struct SurveillanceGroundStationType
	// insertion point per field
	stage.SurveillanceGroundStationType_Annotation_reverseMap = make(map[*NotePropertyType]*SurveillanceGroundStationType)
	for surveillancegroundstationtype := range stage.SurveillanceGroundStationTypes {
		_ = surveillancegroundstationtype
		for _, _notepropertytype := range surveillancegroundstationtype.Annotation {
			stage.SurveillanceGroundStationType_Annotation_reverseMap[_notepropertytype] = surveillancegroundstationtype
		}
	}
	stage.SurveillanceGroundStationType_Extension_reverseMap = make(map[*A_extension_41]*SurveillanceGroundStationType)
	for surveillancegroundstationtype := range stage.SurveillanceGroundStationTypes {
		_ = surveillancegroundstationtype
		for _, _a_extension_41 := range surveillancegroundstationtype.Extension {
			stage.SurveillanceGroundStationType_Extension_reverseMap[_a_extension_41] = surveillancegroundstationtype
		}
	}

	// Compute reverse map for named struct TACANTimeSliceType
	// insertion point per field
	stage.TACANTimeSliceType_Authority_reverseMap = make(map[*AuthorityForNavaidEquipmentPropertyType]*TACANTimeSliceType)
	for tacantimeslicetype := range stage.TACANTimeSliceTypes {
		_ = tacantimeslicetype
		for _, _authorityfornavaidequipmentpropertytype := range tacantimeslicetype.Authority {
			stage.TACANTimeSliceType_Authority_reverseMap[_authorityfornavaidequipmentpropertytype] = tacantimeslicetype
		}
	}
	stage.TACANTimeSliceType_Monitoring_reverseMap = make(map[*NavaidEquipmentMonitoringPropertyType]*TACANTimeSliceType)
	for tacantimeslicetype := range stage.TACANTimeSliceTypes {
		_ = tacantimeslicetype
		for _, _navaidequipmentmonitoringpropertytype := range tacantimeslicetype.Monitoring {
			stage.TACANTimeSliceType_Monitoring_reverseMap[_navaidequipmentmonitoringpropertytype] = tacantimeslicetype
		}
	}
	stage.TACANTimeSliceType_Availability_reverseMap = make(map[*NavaidOperationalStatusPropertyType]*TACANTimeSliceType)
	for tacantimeslicetype := range stage.TACANTimeSliceTypes {
		_ = tacantimeslicetype
		for _, _navaidoperationalstatuspropertytype := range tacantimeslicetype.Availability {
			stage.TACANTimeSliceType_Availability_reverseMap[_navaidoperationalstatuspropertytype] = tacantimeslicetype
		}
	}
	stage.TACANTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*TACANTimeSliceType)
	for tacantimeslicetype := range stage.TACANTimeSliceTypes {
		_ = tacantimeslicetype
		for _, _notepropertytype := range tacantimeslicetype.Annotation {
			stage.TACANTimeSliceType_Annotation_reverseMap[_notepropertytype] = tacantimeslicetype
		}
	}
	stage.TACANTimeSliceType_Extension_reverseMap = make(map[*A_extension_22]*TACANTimeSliceType)
	for tacantimeslicetype := range stage.TACANTimeSliceTypes {
		_ = tacantimeslicetype
		for _, _a_extension_22 := range tacantimeslicetype.Extension {
			stage.TACANTimeSliceType_Extension_reverseMap[_a_extension_22] = tacantimeslicetype
		}
	}

	// Compute reverse map for named struct TACANType
	// insertion point per field
	stage.TACANType_TimeSlice_reverseMap = make(map[*TACANTimeSlicePropertyType]*TACANType)
	for tacantype := range stage.TACANTypes {
		_ = tacantype
		for _, _tacantimeslicepropertytype := range tacantype.TimeSlice {
			stage.TACANType_TimeSlice_reverseMap[_tacantimeslicepropertytype] = tacantype
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
	stage.TimesheetType_Extension_reverseMap = make(map[*A_extension_34]*TimesheetType)
	for timesheettype := range stage.TimesheetTypes {
		_ = timesheettype
		for _, _a_extension_34 := range timesheettype.Extension {
			stage.TimesheetType_Extension_reverseMap[_a_extension_34] = timesheettype
		}
	}

	// Compute reverse map for named struct VORTimeSliceType
	// insertion point per field
	stage.VORTimeSliceType_Authority_reverseMap = make(map[*AuthorityForNavaidEquipmentPropertyType]*VORTimeSliceType)
	for vortimeslicetype := range stage.VORTimeSliceTypes {
		_ = vortimeslicetype
		for _, _authorityfornavaidequipmentpropertytype := range vortimeslicetype.Authority {
			stage.VORTimeSliceType_Authority_reverseMap[_authorityfornavaidequipmentpropertytype] = vortimeslicetype
		}
	}
	stage.VORTimeSliceType_Monitoring_reverseMap = make(map[*NavaidEquipmentMonitoringPropertyType]*VORTimeSliceType)
	for vortimeslicetype := range stage.VORTimeSliceTypes {
		_ = vortimeslicetype
		for _, _navaidequipmentmonitoringpropertytype := range vortimeslicetype.Monitoring {
			stage.VORTimeSliceType_Monitoring_reverseMap[_navaidequipmentmonitoringpropertytype] = vortimeslicetype
		}
	}
	stage.VORTimeSliceType_Availability_reverseMap = make(map[*NavaidOperationalStatusPropertyType]*VORTimeSliceType)
	for vortimeslicetype := range stage.VORTimeSliceTypes {
		_ = vortimeslicetype
		for _, _navaidoperationalstatuspropertytype := range vortimeslicetype.Availability {
			stage.VORTimeSliceType_Availability_reverseMap[_navaidoperationalstatuspropertytype] = vortimeslicetype
		}
	}
	stage.VORTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*VORTimeSliceType)
	for vortimeslicetype := range stage.VORTimeSliceTypes {
		_ = vortimeslicetype
		for _, _notepropertytype := range vortimeslicetype.Annotation {
			stage.VORTimeSliceType_Annotation_reverseMap[_notepropertytype] = vortimeslicetype
		}
	}
	stage.VORTimeSliceType_Extension_reverseMap = make(map[*A_extension_23]*VORTimeSliceType)
	for vortimeslicetype := range stage.VORTimeSliceTypes {
		_ = vortimeslicetype
		for _, _a_extension_23 := range vortimeslicetype.Extension {
			stage.VORTimeSliceType_Extension_reverseMap[_a_extension_23] = vortimeslicetype
		}
	}

	// Compute reverse map for named struct VORType
	// insertion point per field
	stage.VORType_TimeSlice_reverseMap = make(map[*VORTimeSlicePropertyType]*VORType)
	for vortype := range stage.VORTypes {
		_ = vortype
		for _, _vortimeslicepropertytype := range vortype.TimeSlice {
			stage.VORType_TimeSlice_reverseMap[_vortimeslicepropertytype] = vortype
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

	res = __gong__appendInstances(res, stage.A_extension_5s)

	res = __gong__appendInstances(res, stage.A_extension_6s)

	res = __gong__appendInstances(res, stage.A_extension_7s)

	res = __gong__appendInstances(res, stage.A_extension_8s)

	res = __gong__appendInstances(res, stage.A_extension_9s)

	res = __gong__appendInstances(res, stage.AbstractNavaidEquipmentTypes)

	res = __gong__appendInstances(res, stage.AbstractNavigationSystemCheckpointTypes)

	res = __gong__appendInstances(res, stage.AbstractPropertiesWithScheduleTypes)

	res = __gong__appendInstances(res, stage.AbstractRadarEquipmentTypes)

	res = __gong__appendInstances(res, stage.AbstractSurveillanceRadarTypes)

	res = __gong__appendInstances(res, stage.AirportHeliportPropertyTypes)

	res = __gong__appendInstances(res, stage.AirspacePropertyTypes)

	res = __gong__appendInstances(res, stage.AngleIndicationTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.AngleIndicationTimeSliceTypes)

	res = __gong__appendInstances(res, stage.AngleIndicationTypes)

	res = __gong__appendInstances(res, stage.AuthorityForNavaidEquipmentPropertyTypes)

	res = __gong__appendInstances(res, stage.AuthorityForNavaidEquipmentTypes)

	res = __gong__appendInstances(res, stage.AuthorityForSpecialNavigationStationPropertyTypes)

	res = __gong__appendInstances(res, stage.AuthorityForSpecialNavigationStationTypes)

	res = __gong__appendInstances(res, stage.AuthorityForSpecialNavigationSystemPropertyTypes)

	res = __gong__appendInstances(res, stage.AuthorityForSpecialNavigationSystemTypes)

	res = __gong__appendInstances(res, stage.AzimuthTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.AzimuthTimeSliceTypes)

	res = __gong__appendInstances(res, stage.AzimuthTypes)

	res = __gong__appendInstances(res, stage.CheckpointINSTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.CheckpointINSTimeSliceTypes)

	res = __gong__appendInstances(res, stage.CheckpointINSTypes)

	res = __gong__appendInstances(res, stage.CheckpointVORTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.CheckpointVORTimeSliceTypes)

	res = __gong__appendInstances(res, stage.CheckpointVORTypes)

	res = __gong__appendInstances(res, stage.ContactInformationPropertyTypes)

	res = __gong__appendInstances(res, stage.ContactInformationTypes)

	res = __gong__appendInstances(res, stage.DMETimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.DMETimeSliceTypes)

	res = __gong__appendInstances(res, stage.DMETypes)

	res = __gong__appendInstances(res, stage.DesignatedPointPropertyTypes)

	res = __gong__appendInstances(res, stage.DesignatedPointTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.DesignatedPointTimeSliceTypes)

	res = __gong__appendInstances(res, stage.DesignatedPointTypes)

	res = __gong__appendInstances(res, stage.DirectionFinderTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.DirectionFinderTimeSliceTypes)

	res = __gong__appendInstances(res, stage.DirectionFinderTypes)

	res = __gong__appendInstances(res, stage.DistanceIndicationTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.DistanceIndicationTimeSliceTypes)

	res = __gong__appendInstances(res, stage.DistanceIndicationTypes)

	res = __gong__appendInstances(res, stage.ElevatedPointPropertyTypes)

	res = __gong__appendInstances(res, stage.ElevatedPointTypes)

	res = __gong__appendInstances(res, stage.ElevationTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.ElevationTimeSliceTypes)

	res = __gong__appendInstances(res, stage.ElevationTypes)

	res = __gong__appendInstances(res, stage.GlidepathTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.GlidepathTimeSliceTypes)

	res = __gong__appendInstances(res, stage.GlidepathTypes)

	res = __gong__appendInstances(res, stage.InformationServicePropertyTypes)

	res = __gong__appendInstances(res, stage.LinguisticNotePropertyTypes)

	res = __gong__appendInstances(res, stage.LinguisticNoteTypes)

	res = __gong__appendInstances(res, stage.LocalizerTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.LocalizerTimeSliceTypes)

	res = __gong__appendInstances(res, stage.LocalizerTypes)

	res = __gong__appendInstances(res, stage.MarkerBeaconTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.MarkerBeaconTimeSliceTypes)

	res = __gong__appendInstances(res, stage.MarkerBeaconTypes)

	res = __gong__appendInstances(res, stage.NDBTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.NDBTimeSliceTypes)

	res = __gong__appendInstances(res, stage.NDBTypes)

	res = __gong__appendInstances(res, stage.NavaidComponentPropertyTypes)

	res = __gong__appendInstances(res, stage.NavaidComponentTypes)

	res = __gong__appendInstances(res, stage.NavaidEquipmentMonitoringPropertyTypes)

	res = __gong__appendInstances(res, stage.NavaidEquipmentMonitoringTypes)

	res = __gong__appendInstances(res, stage.NavaidEquipmentPropertyTypes)

	res = __gong__appendInstances(res, stage.NavaidOperationalStatusPropertyTypes)

	res = __gong__appendInstances(res, stage.NavaidOperationalStatusTypes)

	res = __gong__appendInstances(res, stage.NavaidPropertyTypes)

	res = __gong__appendInstances(res, stage.NavaidTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.NavaidTimeSliceTypes)

	res = __gong__appendInstances(res, stage.NavaidTypes)

	res = __gong__appendInstances(res, stage.NotePropertyTypes)

	res = __gong__appendInstances(res, stage.NoteTypes)

	res = __gong__appendInstances(res, stage.OnlineContactPropertyTypes)

	res = __gong__appendInstances(res, stage.OnlineContactTypes)

	res = __gong__appendInstances(res, stage.OrganisationAuthorityPropertyTypes)

	res = __gong__appendInstances(res, stage.PointPropertyTypes)

	res = __gong__appendInstances(res, stage.PointTypes)

	res = __gong__appendInstances(res, stage.PostalAddressPropertyTypes)

	res = __gong__appendInstances(res, stage.PostalAddressTypes)

	res = __gong__appendInstances(res, stage.PrecisionApproachRadarTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.PrecisionApproachRadarTimeSliceTypes)

	res = __gong__appendInstances(res, stage.PrecisionApproachRadarTypes)

	res = __gong__appendInstances(res, stage.PrimarySurveillanceRadarTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.PrimarySurveillanceRadarTimeSliceTypes)

	res = __gong__appendInstances(res, stage.PrimarySurveillanceRadarTypes)

	res = __gong__appendInstances(res, stage.RadarComponentPropertyTypes)

	res = __gong__appendInstances(res, stage.RadarComponentTypes)

	res = __gong__appendInstances(res, stage.RadarEquipmentPropertyTypes)

	res = __gong__appendInstances(res, stage.RadarSystemTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.RadarSystemTimeSliceTypes)

	res = __gong__appendInstances(res, stage.RadarSystemTypes)

	res = __gong__appendInstances(res, stage.ReflectorPropertyTypes)

	res = __gong__appendInstances(res, stage.ReflectorTypes)

	res = __gong__appendInstances(res, stage.RunwayCentrelinePointPropertyTypes)

	res = __gong__appendInstances(res, stage.RunwayDirectionPropertyTypes)

	res = __gong__appendInstances(res, stage.RunwayPropertyTypes)

	res = __gong__appendInstances(res, stage.SDFTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.SDFTimeSliceTypes)

	res = __gong__appendInstances(res, stage.SDFTypes)

	res = __gong__appendInstances(res, stage.SecondarySurveillanceRadarTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.SecondarySurveillanceRadarTimeSliceTypes)

	res = __gong__appendInstances(res, stage.SecondarySurveillanceRadarTypes)

	res = __gong__appendInstances(res, stage.SignificantPointInAirspaceTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.SignificantPointInAirspaceTimeSliceTypes)

	res = __gong__appendInstances(res, stage.SignificantPointInAirspaceTypes)

	res = __gong__appendInstances(res, stage.SpecialNavigationStationStatusPropertyTypes)

	res = __gong__appendInstances(res, stage.SpecialNavigationStationStatusTypes)

	res = __gong__appendInstances(res, stage.SpecialNavigationStationTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.SpecialNavigationStationTimeSliceTypes)

	res = __gong__appendInstances(res, stage.SpecialNavigationStationTypes)

	res = __gong__appendInstances(res, stage.SpecialNavigationSystemPropertyTypes)

	res = __gong__appendInstances(res, stage.SpecialNavigationSystemTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.SpecialNavigationSystemTimeSliceTypes)

	res = __gong__appendInstances(res, stage.SpecialNavigationSystemTypes)

	res = __gong__appendInstances(res, stage.SurveillanceGroundStationPropertyTypes)

	res = __gong__appendInstances(res, stage.SurveillanceGroundStationTypes)

	res = __gong__appendInstances(res, stage.TACANTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.TACANTimeSliceTypes)

	res = __gong__appendInstances(res, stage.TACANTypes)

	res = __gong__appendInstances(res, stage.TelephoneContactPropertyTypes)

	res = __gong__appendInstances(res, stage.TelephoneContactTypes)

	res = __gong__appendInstances(res, stage.TimesheetPropertyTypes)

	res = __gong__appendInstances(res, stage.TimesheetTypes)

	res = __gong__appendInstances(res, stage.TouchDownLiftOffPropertyTypes)

	res = __gong__appendInstances(res, stage.UnitPropertyTypes)

	res = __gong__appendInstances(res, stage.VORPropertyTypes)

	res = __gong__appendInstances(res, stage.VORTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.VORTimeSliceTypes)

	res = __gong__appendInstances(res, stage.VORTypes)

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

func (abstractnavaidequipmenttype *AbstractNavaidEquipmentType) GongCopy() GongstructIF {
	newInstance := new(AbstractNavaidEquipmentType)
	abstractnavaidequipmenttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstractnavigationsystemcheckpointtype *AbstractNavigationSystemCheckpointType) GongCopy() GongstructIF {
	newInstance := new(AbstractNavigationSystemCheckpointType)
	abstractnavigationsystemcheckpointtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongCopy() GongstructIF {
	newInstance := new(AbstractPropertiesWithScheduleType)
	abstractpropertieswithscheduletype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstractradarequipmenttype *AbstractRadarEquipmentType) GongCopy() GongstructIF {
	newInstance := new(AbstractRadarEquipmentType)
	abstractradarequipmenttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstractsurveillanceradartype *AbstractSurveillanceRadarType) GongCopy() GongstructIF {
	newInstance := new(AbstractSurveillanceRadarType)
	abstractsurveillanceradartype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongCopy() GongstructIF {
	newInstance := new(AirportHeliportPropertyType)
	airportheliportpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airspacepropertytype *AirspacePropertyType) GongCopy() GongstructIF {
	newInstance := new(AirspacePropertyType)
	airspacepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (angleindicationtimeslicepropertytype *AngleIndicationTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(AngleIndicationTimeSlicePropertyType)
	angleindicationtimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (angleindicationtimeslicetype *AngleIndicationTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(AngleIndicationTimeSliceType)
	angleindicationtimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (angleindicationtype *AngleIndicationType) GongCopy() GongstructIF {
	newInstance := new(AngleIndicationType)
	angleindicationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (authorityfornavaidequipmentpropertytype *AuthorityForNavaidEquipmentPropertyType) GongCopy() GongstructIF {
	newInstance := new(AuthorityForNavaidEquipmentPropertyType)
	authorityfornavaidequipmentpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (authorityfornavaidequipmenttype *AuthorityForNavaidEquipmentType) GongCopy() GongstructIF {
	newInstance := new(AuthorityForNavaidEquipmentType)
	authorityfornavaidequipmenttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (authorityforspecialnavigationstationpropertytype *AuthorityForSpecialNavigationStationPropertyType) GongCopy() GongstructIF {
	newInstance := new(AuthorityForSpecialNavigationStationPropertyType)
	authorityforspecialnavigationstationpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (authorityforspecialnavigationstationtype *AuthorityForSpecialNavigationStationType) GongCopy() GongstructIF {
	newInstance := new(AuthorityForSpecialNavigationStationType)
	authorityforspecialnavigationstationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (authorityforspecialnavigationsystempropertytype *AuthorityForSpecialNavigationSystemPropertyType) GongCopy() GongstructIF {
	newInstance := new(AuthorityForSpecialNavigationSystemPropertyType)
	authorityforspecialnavigationsystempropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (authorityforspecialnavigationsystemtype *AuthorityForSpecialNavigationSystemType) GongCopy() GongstructIF {
	newInstance := new(AuthorityForSpecialNavigationSystemType)
	authorityforspecialnavigationsystemtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (azimuthtimeslicepropertytype *AzimuthTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(AzimuthTimeSlicePropertyType)
	azimuthtimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (azimuthtimeslicetype *AzimuthTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(AzimuthTimeSliceType)
	azimuthtimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (azimuthtype *AzimuthType) GongCopy() GongstructIF {
	newInstance := new(AzimuthType)
	azimuthtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (checkpointinstimeslicepropertytype *CheckpointINSTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(CheckpointINSTimeSlicePropertyType)
	checkpointinstimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (checkpointinstimeslicetype *CheckpointINSTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(CheckpointINSTimeSliceType)
	checkpointinstimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (checkpointinstype *CheckpointINSType) GongCopy() GongstructIF {
	newInstance := new(CheckpointINSType)
	checkpointinstype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (checkpointvortimeslicepropertytype *CheckpointVORTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(CheckpointVORTimeSlicePropertyType)
	checkpointvortimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (checkpointvortimeslicetype *CheckpointVORTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(CheckpointVORTimeSliceType)
	checkpointvortimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (checkpointvortype *CheckpointVORType) GongCopy() GongstructIF {
	newInstance := new(CheckpointVORType)
	checkpointvortype.GongCopyBasicFields(newInstance)
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

func (dmetimeslicepropertytype *DMETimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(DMETimeSlicePropertyType)
	dmetimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (dmetimeslicetype *DMETimeSliceType) GongCopy() GongstructIF {
	newInstance := new(DMETimeSliceType)
	dmetimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (dmetype *DMEType) GongCopy() GongstructIF {
	newInstance := new(DMEType)
	dmetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (designatedpointpropertytype *DesignatedPointPropertyType) GongCopy() GongstructIF {
	newInstance := new(DesignatedPointPropertyType)
	designatedpointpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (designatedpointtimeslicepropertytype *DesignatedPointTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(DesignatedPointTimeSlicePropertyType)
	designatedpointtimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (designatedpointtimeslicetype *DesignatedPointTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(DesignatedPointTimeSliceType)
	designatedpointtimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (designatedpointtype *DesignatedPointType) GongCopy() GongstructIF {
	newInstance := new(DesignatedPointType)
	designatedpointtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (directionfindertimeslicepropertytype *DirectionFinderTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(DirectionFinderTimeSlicePropertyType)
	directionfindertimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (directionfindertimeslicetype *DirectionFinderTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(DirectionFinderTimeSliceType)
	directionfindertimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (directionfindertype *DirectionFinderType) GongCopy() GongstructIF {
	newInstance := new(DirectionFinderType)
	directionfindertype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (distanceindicationtimeslicepropertytype *DistanceIndicationTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(DistanceIndicationTimeSlicePropertyType)
	distanceindicationtimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (distanceindicationtimeslicetype *DistanceIndicationTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(DistanceIndicationTimeSliceType)
	distanceindicationtimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (distanceindicationtype *DistanceIndicationType) GongCopy() GongstructIF {
	newInstance := new(DistanceIndicationType)
	distanceindicationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (elevatedpointpropertytype *ElevatedPointPropertyType) GongCopy() GongstructIF {
	newInstance := new(ElevatedPointPropertyType)
	elevatedpointpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (elevatedpointtype *ElevatedPointType) GongCopy() GongstructIF {
	newInstance := new(ElevatedPointType)
	elevatedpointtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (elevationtimeslicepropertytype *ElevationTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(ElevationTimeSlicePropertyType)
	elevationtimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (elevationtimeslicetype *ElevationTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(ElevationTimeSliceType)
	elevationtimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (elevationtype *ElevationType) GongCopy() GongstructIF {
	newInstance := new(ElevationType)
	elevationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (glidepathtimeslicepropertytype *GlidepathTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(GlidepathTimeSlicePropertyType)
	glidepathtimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (glidepathtimeslicetype *GlidepathTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(GlidepathTimeSliceType)
	glidepathtimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (glidepathtype *GlidepathType) GongCopy() GongstructIF {
	newInstance := new(GlidepathType)
	glidepathtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (informationservicepropertytype *InformationServicePropertyType) GongCopy() GongstructIF {
	newInstance := new(InformationServicePropertyType)
	informationservicepropertytype.GongCopyBasicFields(newInstance)
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

func (localizertimeslicepropertytype *LocalizerTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(LocalizerTimeSlicePropertyType)
	localizertimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (localizertimeslicetype *LocalizerTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(LocalizerTimeSliceType)
	localizertimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (localizertype *LocalizerType) GongCopy() GongstructIF {
	newInstance := new(LocalizerType)
	localizertype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (markerbeacontimeslicepropertytype *MarkerBeaconTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(MarkerBeaconTimeSlicePropertyType)
	markerbeacontimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (markerbeacontimeslicetype *MarkerBeaconTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(MarkerBeaconTimeSliceType)
	markerbeacontimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (markerbeacontype *MarkerBeaconType) GongCopy() GongstructIF {
	newInstance := new(MarkerBeaconType)
	markerbeacontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (ndbtimeslicepropertytype *NDBTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(NDBTimeSlicePropertyType)
	ndbtimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (ndbtimeslicetype *NDBTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(NDBTimeSliceType)
	ndbtimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (ndbtype *NDBType) GongCopy() GongstructIF {
	newInstance := new(NDBType)
	ndbtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (navaidcomponentpropertytype *NavaidComponentPropertyType) GongCopy() GongstructIF {
	newInstance := new(NavaidComponentPropertyType)
	navaidcomponentpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (navaidcomponenttype *NavaidComponentType) GongCopy() GongstructIF {
	newInstance := new(NavaidComponentType)
	navaidcomponenttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (navaidequipmentmonitoringpropertytype *NavaidEquipmentMonitoringPropertyType) GongCopy() GongstructIF {
	newInstance := new(NavaidEquipmentMonitoringPropertyType)
	navaidequipmentmonitoringpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (navaidequipmentmonitoringtype *NavaidEquipmentMonitoringType) GongCopy() GongstructIF {
	newInstance := new(NavaidEquipmentMonitoringType)
	navaidequipmentmonitoringtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongCopy() GongstructIF {
	newInstance := new(NavaidEquipmentPropertyType)
	navaidequipmentpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (navaidoperationalstatuspropertytype *NavaidOperationalStatusPropertyType) GongCopy() GongstructIF {
	newInstance := new(NavaidOperationalStatusPropertyType)
	navaidoperationalstatuspropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (navaidoperationalstatustype *NavaidOperationalStatusType) GongCopy() GongstructIF {
	newInstance := new(NavaidOperationalStatusType)
	navaidoperationalstatustype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (navaidpropertytype *NavaidPropertyType) GongCopy() GongstructIF {
	newInstance := new(NavaidPropertyType)
	navaidpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (navaidtimeslicepropertytype *NavaidTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(NavaidTimeSlicePropertyType)
	navaidtimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (navaidtimeslicetype *NavaidTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(NavaidTimeSliceType)
	navaidtimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (navaidtype *NavaidType) GongCopy() GongstructIF {
	newInstance := new(NavaidType)
	navaidtype.GongCopyBasicFields(newInstance)
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

func (precisionapproachradartimeslicepropertytype *PrecisionApproachRadarTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(PrecisionApproachRadarTimeSlicePropertyType)
	precisionapproachradartimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (precisionapproachradartimeslicetype *PrecisionApproachRadarTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(PrecisionApproachRadarTimeSliceType)
	precisionapproachradartimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (precisionapproachradartype *PrecisionApproachRadarType) GongCopy() GongstructIF {
	newInstance := new(PrecisionApproachRadarType)
	precisionapproachradartype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (primarysurveillanceradartimeslicepropertytype *PrimarySurveillanceRadarTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(PrimarySurveillanceRadarTimeSlicePropertyType)
	primarysurveillanceradartimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (primarysurveillanceradartimeslicetype *PrimarySurveillanceRadarTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(PrimarySurveillanceRadarTimeSliceType)
	primarysurveillanceradartimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (primarysurveillanceradartype *PrimarySurveillanceRadarType) GongCopy() GongstructIF {
	newInstance := new(PrimarySurveillanceRadarType)
	primarysurveillanceradartype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (radarcomponentpropertytype *RadarComponentPropertyType) GongCopy() GongstructIF {
	newInstance := new(RadarComponentPropertyType)
	radarcomponentpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (radarcomponenttype *RadarComponentType) GongCopy() GongstructIF {
	newInstance := new(RadarComponentType)
	radarcomponenttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (radarequipmentpropertytype *RadarEquipmentPropertyType) GongCopy() GongstructIF {
	newInstance := new(RadarEquipmentPropertyType)
	radarequipmentpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (radarsystemtimeslicepropertytype *RadarSystemTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(RadarSystemTimeSlicePropertyType)
	radarsystemtimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (radarsystemtimeslicetype *RadarSystemTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(RadarSystemTimeSliceType)
	radarsystemtimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (radarsystemtype *RadarSystemType) GongCopy() GongstructIF {
	newInstance := new(RadarSystemType)
	radarsystemtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (reflectorpropertytype *ReflectorPropertyType) GongCopy() GongstructIF {
	newInstance := new(ReflectorPropertyType)
	reflectorpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (reflectortype *ReflectorType) GongCopy() GongstructIF {
	newInstance := new(ReflectorType)
	reflectortype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongCopy() GongstructIF {
	newInstance := new(RunwayCentrelinePointPropertyType)
	runwaycentrelinepointpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongCopy() GongstructIF {
	newInstance := new(RunwayDirectionPropertyType)
	runwaydirectionpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (runwaypropertytype *RunwayPropertyType) GongCopy() GongstructIF {
	newInstance := new(RunwayPropertyType)
	runwaypropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (sdftimeslicepropertytype *SDFTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(SDFTimeSlicePropertyType)
	sdftimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (sdftimeslicetype *SDFTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(SDFTimeSliceType)
	sdftimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (sdftype *SDFType) GongCopy() GongstructIF {
	newInstance := new(SDFType)
	sdftype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (secondarysurveillanceradartimeslicepropertytype *SecondarySurveillanceRadarTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(SecondarySurveillanceRadarTimeSlicePropertyType)
	secondarysurveillanceradartimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (secondarysurveillanceradartimeslicetype *SecondarySurveillanceRadarTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(SecondarySurveillanceRadarTimeSliceType)
	secondarysurveillanceradartimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (secondarysurveillanceradartype *SecondarySurveillanceRadarType) GongCopy() GongstructIF {
	newInstance := new(SecondarySurveillanceRadarType)
	secondarysurveillanceradartype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (significantpointinairspacetimeslicepropertytype *SignificantPointInAirspaceTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(SignificantPointInAirspaceTimeSlicePropertyType)
	significantpointinairspacetimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (significantpointinairspacetimeslicetype *SignificantPointInAirspaceTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(SignificantPointInAirspaceTimeSliceType)
	significantpointinairspacetimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (significantpointinairspacetype *SignificantPointInAirspaceType) GongCopy() GongstructIF {
	newInstance := new(SignificantPointInAirspaceType)
	significantpointinairspacetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (specialnavigationstationstatuspropertytype *SpecialNavigationStationStatusPropertyType) GongCopy() GongstructIF {
	newInstance := new(SpecialNavigationStationStatusPropertyType)
	specialnavigationstationstatuspropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (specialnavigationstationstatustype *SpecialNavigationStationStatusType) GongCopy() GongstructIF {
	newInstance := new(SpecialNavigationStationStatusType)
	specialnavigationstationstatustype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (specialnavigationstationtimeslicepropertytype *SpecialNavigationStationTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(SpecialNavigationStationTimeSlicePropertyType)
	specialnavigationstationtimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (specialnavigationstationtimeslicetype *SpecialNavigationStationTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(SpecialNavigationStationTimeSliceType)
	specialnavigationstationtimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (specialnavigationstationtype *SpecialNavigationStationType) GongCopy() GongstructIF {
	newInstance := new(SpecialNavigationStationType)
	specialnavigationstationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (specialnavigationsystempropertytype *SpecialNavigationSystemPropertyType) GongCopy() GongstructIF {
	newInstance := new(SpecialNavigationSystemPropertyType)
	specialnavigationsystempropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (specialnavigationsystemtimeslicepropertytype *SpecialNavigationSystemTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(SpecialNavigationSystemTimeSlicePropertyType)
	specialnavigationsystemtimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (specialnavigationsystemtimeslicetype *SpecialNavigationSystemTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(SpecialNavigationSystemTimeSliceType)
	specialnavigationsystemtimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (specialnavigationsystemtype *SpecialNavigationSystemType) GongCopy() GongstructIF {
	newInstance := new(SpecialNavigationSystemType)
	specialnavigationsystemtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (surveillancegroundstationpropertytype *SurveillanceGroundStationPropertyType) GongCopy() GongstructIF {
	newInstance := new(SurveillanceGroundStationPropertyType)
	surveillancegroundstationpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (surveillancegroundstationtype *SurveillanceGroundStationType) GongCopy() GongstructIF {
	newInstance := new(SurveillanceGroundStationType)
	surveillancegroundstationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tacantimeslicepropertytype *TACANTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(TACANTimeSlicePropertyType)
	tacantimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tacantimeslicetype *TACANTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(TACANTimeSliceType)
	tacantimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tacantype *TACANType) GongCopy() GongstructIF {
	newInstance := new(TACANType)
	tacantype.GongCopyBasicFields(newInstance)
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

func (unitpropertytype *UnitPropertyType) GongCopy() GongstructIF {
	newInstance := new(UnitPropertyType)
	unitpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (vorpropertytype *VORPropertyType) GongCopy() GongstructIF {
	newInstance := new(VORPropertyType)
	vorpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (vortimeslicepropertytype *VORTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(VORTimeSlicePropertyType)
	vortimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (vortimeslicetype *VORTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(VORTimeSliceType)
	vortimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (vortype *VORType) GongCopy() GongstructIF {
	newInstance := new(VORType)
	vortype.GongCopyBasicFields(newInstance)
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

func (abstractnavaidequipmenttype *AbstractNavaidEquipmentType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractnavaidequipmenttype)
}

func (abstractnavigationsystemcheckpointtype *AbstractNavigationSystemCheckpointType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractnavigationsystemcheckpointtype)
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractpropertieswithscheduletype)
}

func (abstractradarequipmenttype *AbstractRadarEquipmentType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractradarequipmenttype)
}

func (abstractsurveillanceradartype *AbstractSurveillanceRadarType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractsurveillanceradartype)
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airportheliportpropertytype)
}

func (airspacepropertytype *AirspacePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airspacepropertytype)
}

func (angleindicationtimeslicepropertytype *AngleIndicationTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, angleindicationtimeslicepropertytype)
}

func (angleindicationtimeslicetype *AngleIndicationTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, angleindicationtimeslicetype)
}

func (angleindicationtype *AngleIndicationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, angleindicationtype)
}

func (authorityfornavaidequipmentpropertytype *AuthorityForNavaidEquipmentPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, authorityfornavaidequipmentpropertytype)
}

func (authorityfornavaidequipmenttype *AuthorityForNavaidEquipmentType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, authorityfornavaidequipmenttype)
}

func (authorityforspecialnavigationstationpropertytype *AuthorityForSpecialNavigationStationPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, authorityforspecialnavigationstationpropertytype)
}

func (authorityforspecialnavigationstationtype *AuthorityForSpecialNavigationStationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, authorityforspecialnavigationstationtype)
}

func (authorityforspecialnavigationsystempropertytype *AuthorityForSpecialNavigationSystemPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, authorityforspecialnavigationsystempropertytype)
}

func (authorityforspecialnavigationsystemtype *AuthorityForSpecialNavigationSystemType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, authorityforspecialnavigationsystemtype)
}

func (azimuthtimeslicepropertytype *AzimuthTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, azimuthtimeslicepropertytype)
}

func (azimuthtimeslicetype *AzimuthTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, azimuthtimeslicetype)
}

func (azimuthtype *AzimuthType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, azimuthtype)
}

func (checkpointinstimeslicepropertytype *CheckpointINSTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, checkpointinstimeslicepropertytype)
}

func (checkpointinstimeslicetype *CheckpointINSTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, checkpointinstimeslicetype)
}

func (checkpointinstype *CheckpointINSType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, checkpointinstype)
}

func (checkpointvortimeslicepropertytype *CheckpointVORTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, checkpointvortimeslicepropertytype)
}

func (checkpointvortimeslicetype *CheckpointVORTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, checkpointvortimeslicetype)
}

func (checkpointvortype *CheckpointVORType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, checkpointvortype)
}

func (contactinformationpropertytype *ContactInformationPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, contactinformationpropertytype)
}

func (contactinformationtype *ContactInformationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, contactinformationtype)
}

func (dmetimeslicepropertytype *DMETimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, dmetimeslicepropertytype)
}

func (dmetimeslicetype *DMETimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, dmetimeslicetype)
}

func (dmetype *DMEType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, dmetype)
}

func (designatedpointpropertytype *DesignatedPointPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, designatedpointpropertytype)
}

func (designatedpointtimeslicepropertytype *DesignatedPointTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, designatedpointtimeslicepropertytype)
}

func (designatedpointtimeslicetype *DesignatedPointTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, designatedpointtimeslicetype)
}

func (designatedpointtype *DesignatedPointType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, designatedpointtype)
}

func (directionfindertimeslicepropertytype *DirectionFinderTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, directionfindertimeslicepropertytype)
}

func (directionfindertimeslicetype *DirectionFinderTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, directionfindertimeslicetype)
}

func (directionfindertype *DirectionFinderType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, directionfindertype)
}

func (distanceindicationtimeslicepropertytype *DistanceIndicationTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, distanceindicationtimeslicepropertytype)
}

func (distanceindicationtimeslicetype *DistanceIndicationTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, distanceindicationtimeslicetype)
}

func (distanceindicationtype *DistanceIndicationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, distanceindicationtype)
}

func (elevatedpointpropertytype *ElevatedPointPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, elevatedpointpropertytype)
}

func (elevatedpointtype *ElevatedPointType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, elevatedpointtype)
}

func (elevationtimeslicepropertytype *ElevationTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, elevationtimeslicepropertytype)
}

func (elevationtimeslicetype *ElevationTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, elevationtimeslicetype)
}

func (elevationtype *ElevationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, elevationtype)
}

func (glidepathtimeslicepropertytype *GlidepathTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, glidepathtimeslicepropertytype)
}

func (glidepathtimeslicetype *GlidepathTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, glidepathtimeslicetype)
}

func (glidepathtype *GlidepathType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, glidepathtype)
}

func (informationservicepropertytype *InformationServicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, informationservicepropertytype)
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, linguisticnotepropertytype)
}

func (linguisticnotetype *LinguisticNoteType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, linguisticnotetype)
}

func (localizertimeslicepropertytype *LocalizerTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, localizertimeslicepropertytype)
}

func (localizertimeslicetype *LocalizerTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, localizertimeslicetype)
}

func (localizertype *LocalizerType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, localizertype)
}

func (markerbeacontimeslicepropertytype *MarkerBeaconTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, markerbeacontimeslicepropertytype)
}

func (markerbeacontimeslicetype *MarkerBeaconTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, markerbeacontimeslicetype)
}

func (markerbeacontype *MarkerBeaconType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, markerbeacontype)
}

func (ndbtimeslicepropertytype *NDBTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, ndbtimeslicepropertytype)
}

func (ndbtimeslicetype *NDBTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, ndbtimeslicetype)
}

func (ndbtype *NDBType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, ndbtype)
}

func (navaidcomponentpropertytype *NavaidComponentPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, navaidcomponentpropertytype)
}

func (navaidcomponenttype *NavaidComponentType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, navaidcomponenttype)
}

func (navaidequipmentmonitoringpropertytype *NavaidEquipmentMonitoringPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, navaidequipmentmonitoringpropertytype)
}

func (navaidequipmentmonitoringtype *NavaidEquipmentMonitoringType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, navaidequipmentmonitoringtype)
}

func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, navaidequipmentpropertytype)
}

func (navaidoperationalstatuspropertytype *NavaidOperationalStatusPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, navaidoperationalstatuspropertytype)
}

func (navaidoperationalstatustype *NavaidOperationalStatusType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, navaidoperationalstatustype)
}

func (navaidpropertytype *NavaidPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, navaidpropertytype)
}

func (navaidtimeslicepropertytype *NavaidTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, navaidtimeslicepropertytype)
}

func (navaidtimeslicetype *NavaidTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, navaidtimeslicetype)
}

func (navaidtype *NavaidType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, navaidtype)
}

func (notepropertytype *NotePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, notepropertytype)
}

func (notetype *NoteType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, notetype)
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

func (pointtype *PointType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, pointtype)
}

func (postaladdresspropertytype *PostalAddressPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, postaladdresspropertytype)
}

func (postaladdresstype *PostalAddressType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, postaladdresstype)
}

func (precisionapproachradartimeslicepropertytype *PrecisionApproachRadarTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, precisionapproachradartimeslicepropertytype)
}

func (precisionapproachradartimeslicetype *PrecisionApproachRadarTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, precisionapproachradartimeslicetype)
}

func (precisionapproachradartype *PrecisionApproachRadarType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, precisionapproachradartype)
}

func (primarysurveillanceradartimeslicepropertytype *PrimarySurveillanceRadarTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, primarysurveillanceradartimeslicepropertytype)
}

func (primarysurveillanceradartimeslicetype *PrimarySurveillanceRadarTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, primarysurveillanceradartimeslicetype)
}

func (primarysurveillanceradartype *PrimarySurveillanceRadarType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, primarysurveillanceradartype)
}

func (radarcomponentpropertytype *RadarComponentPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, radarcomponentpropertytype)
}

func (radarcomponenttype *RadarComponentType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, radarcomponenttype)
}

func (radarequipmentpropertytype *RadarEquipmentPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, radarequipmentpropertytype)
}

func (radarsystemtimeslicepropertytype *RadarSystemTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, radarsystemtimeslicepropertytype)
}

func (radarsystemtimeslicetype *RadarSystemTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, radarsystemtimeslicetype)
}

func (radarsystemtype *RadarSystemType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, radarsystemtype)
}

func (reflectorpropertytype *ReflectorPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, reflectorpropertytype)
}

func (reflectortype *ReflectorType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, reflectortype)
}

func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, runwaycentrelinepointpropertytype)
}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, runwaydirectionpropertytype)
}

func (runwaypropertytype *RunwayPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, runwaypropertytype)
}

func (sdftimeslicepropertytype *SDFTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, sdftimeslicepropertytype)
}

func (sdftimeslicetype *SDFTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, sdftimeslicetype)
}

func (sdftype *SDFType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, sdftype)
}

func (secondarysurveillanceradartimeslicepropertytype *SecondarySurveillanceRadarTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, secondarysurveillanceradartimeslicepropertytype)
}

func (secondarysurveillanceradartimeslicetype *SecondarySurveillanceRadarTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, secondarysurveillanceradartimeslicetype)
}

func (secondarysurveillanceradartype *SecondarySurveillanceRadarType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, secondarysurveillanceradartype)
}

func (significantpointinairspacetimeslicepropertytype *SignificantPointInAirspaceTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, significantpointinairspacetimeslicepropertytype)
}

func (significantpointinairspacetimeslicetype *SignificantPointInAirspaceTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, significantpointinairspacetimeslicetype)
}

func (significantpointinairspacetype *SignificantPointInAirspaceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, significantpointinairspacetype)
}

func (specialnavigationstationstatuspropertytype *SpecialNavigationStationStatusPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, specialnavigationstationstatuspropertytype)
}

func (specialnavigationstationstatustype *SpecialNavigationStationStatusType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, specialnavigationstationstatustype)
}

func (specialnavigationstationtimeslicepropertytype *SpecialNavigationStationTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, specialnavigationstationtimeslicepropertytype)
}

func (specialnavigationstationtimeslicetype *SpecialNavigationStationTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, specialnavigationstationtimeslicetype)
}

func (specialnavigationstationtype *SpecialNavigationStationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, specialnavigationstationtype)
}

func (specialnavigationsystempropertytype *SpecialNavigationSystemPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, specialnavigationsystempropertytype)
}

func (specialnavigationsystemtimeslicepropertytype *SpecialNavigationSystemTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, specialnavigationsystemtimeslicepropertytype)
}

func (specialnavigationsystemtimeslicetype *SpecialNavigationSystemTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, specialnavigationsystemtimeslicetype)
}

func (specialnavigationsystemtype *SpecialNavigationSystemType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, specialnavigationsystemtype)
}

func (surveillancegroundstationpropertytype *SurveillanceGroundStationPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, surveillancegroundstationpropertytype)
}

func (surveillancegroundstationtype *SurveillanceGroundStationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, surveillancegroundstationtype)
}

func (tacantimeslicepropertytype *TACANTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tacantimeslicepropertytype)
}

func (tacantimeslicetype *TACANTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tacantimeslicetype)
}

func (tacantype *TACANType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tacantype)
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

func (unitpropertytype *UnitPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, unitpropertytype)
}

func (vorpropertytype *VORPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, vorpropertytype)
}

func (vortimeslicepropertytype *VORTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, vortimeslicepropertytype)
}

func (vortimeslicetype *VORTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, vortimeslicetype)
}

func (vortype *VORType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, vortype)
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
		stage.AbstractNavaidEquipmentTypes,
		stage.AbstractNavaidEquipmentType_stagedOrder,
		stage.AbstractNavaidEquipmentTypes_reference,
		&stage.AbstractNavaidEquipmentTypes_referenceOrder,
		stage.AbstractNavaidEquipmentTypes_instance,
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
		stage.AbstractNavigationSystemCheckpointTypes,
		stage.AbstractNavigationSystemCheckpointType_stagedOrder,
		stage.AbstractNavigationSystemCheckpointTypes_reference,
		&stage.AbstractNavigationSystemCheckpointTypes_referenceOrder,
		stage.AbstractNavigationSystemCheckpointTypes_instance,
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
		stage.AbstractRadarEquipmentTypes,
		stage.AbstractRadarEquipmentType_stagedOrder,
		stage.AbstractRadarEquipmentTypes_reference,
		&stage.AbstractRadarEquipmentTypes_referenceOrder,
		stage.AbstractRadarEquipmentTypes_instance,
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
		stage.AbstractSurveillanceRadarTypes,
		stage.AbstractSurveillanceRadarType_stagedOrder,
		stage.AbstractSurveillanceRadarTypes_reference,
		&stage.AbstractSurveillanceRadarTypes_referenceOrder,
		stage.AbstractSurveillanceRadarTypes_instance,
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
		stage.AngleIndicationTimeSlicePropertyTypes,
		stage.AngleIndicationTimeSlicePropertyType_stagedOrder,
		stage.AngleIndicationTimeSlicePropertyTypes_reference,
		&stage.AngleIndicationTimeSlicePropertyTypes_referenceOrder,
		stage.AngleIndicationTimeSlicePropertyTypes_instance,
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
		stage.AngleIndicationTimeSliceTypes,
		stage.AngleIndicationTimeSliceType_stagedOrder,
		stage.AngleIndicationTimeSliceTypes_reference,
		&stage.AngleIndicationTimeSliceTypes_referenceOrder,
		stage.AngleIndicationTimeSliceTypes_instance,
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
		stage.AngleIndicationTypes,
		stage.AngleIndicationType_stagedOrder,
		stage.AngleIndicationTypes_reference,
		&stage.AngleIndicationTypes_referenceOrder,
		stage.AngleIndicationTypes_instance,
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
		stage.AuthorityForNavaidEquipmentPropertyTypes,
		stage.AuthorityForNavaidEquipmentPropertyType_stagedOrder,
		stage.AuthorityForNavaidEquipmentPropertyTypes_reference,
		&stage.AuthorityForNavaidEquipmentPropertyTypes_referenceOrder,
		stage.AuthorityForNavaidEquipmentPropertyTypes_instance,
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
		stage.AuthorityForNavaidEquipmentTypes,
		stage.AuthorityForNavaidEquipmentType_stagedOrder,
		stage.AuthorityForNavaidEquipmentTypes_reference,
		&stage.AuthorityForNavaidEquipmentTypes_referenceOrder,
		stage.AuthorityForNavaidEquipmentTypes_instance,
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
		stage.AuthorityForSpecialNavigationStationPropertyTypes,
		stage.AuthorityForSpecialNavigationStationPropertyType_stagedOrder,
		stage.AuthorityForSpecialNavigationStationPropertyTypes_reference,
		&stage.AuthorityForSpecialNavigationStationPropertyTypes_referenceOrder,
		stage.AuthorityForSpecialNavigationStationPropertyTypes_instance,
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
		stage.AuthorityForSpecialNavigationStationTypes,
		stage.AuthorityForSpecialNavigationStationType_stagedOrder,
		stage.AuthorityForSpecialNavigationStationTypes_reference,
		&stage.AuthorityForSpecialNavigationStationTypes_referenceOrder,
		stage.AuthorityForSpecialNavigationStationTypes_instance,
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
		stage.AuthorityForSpecialNavigationSystemPropertyTypes,
		stage.AuthorityForSpecialNavigationSystemPropertyType_stagedOrder,
		stage.AuthorityForSpecialNavigationSystemPropertyTypes_reference,
		&stage.AuthorityForSpecialNavigationSystemPropertyTypes_referenceOrder,
		stage.AuthorityForSpecialNavigationSystemPropertyTypes_instance,
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
		stage.AuthorityForSpecialNavigationSystemTypes,
		stage.AuthorityForSpecialNavigationSystemType_stagedOrder,
		stage.AuthorityForSpecialNavigationSystemTypes_reference,
		&stage.AuthorityForSpecialNavigationSystemTypes_referenceOrder,
		stage.AuthorityForSpecialNavigationSystemTypes_instance,
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
		stage.AzimuthTimeSlicePropertyTypes,
		stage.AzimuthTimeSlicePropertyType_stagedOrder,
		stage.AzimuthTimeSlicePropertyTypes_reference,
		&stage.AzimuthTimeSlicePropertyTypes_referenceOrder,
		stage.AzimuthTimeSlicePropertyTypes_instance,
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
		stage.AzimuthTimeSliceTypes,
		stage.AzimuthTimeSliceType_stagedOrder,
		stage.AzimuthTimeSliceTypes_reference,
		&stage.AzimuthTimeSliceTypes_referenceOrder,
		stage.AzimuthTimeSliceTypes_instance,
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
		stage.AzimuthTypes,
		stage.AzimuthType_stagedOrder,
		stage.AzimuthTypes_reference,
		&stage.AzimuthTypes_referenceOrder,
		stage.AzimuthTypes_instance,
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
		stage.CheckpointINSTimeSlicePropertyTypes,
		stage.CheckpointINSTimeSlicePropertyType_stagedOrder,
		stage.CheckpointINSTimeSlicePropertyTypes_reference,
		&stage.CheckpointINSTimeSlicePropertyTypes_referenceOrder,
		stage.CheckpointINSTimeSlicePropertyTypes_instance,
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
		stage.CheckpointINSTimeSliceTypes,
		stage.CheckpointINSTimeSliceType_stagedOrder,
		stage.CheckpointINSTimeSliceTypes_reference,
		&stage.CheckpointINSTimeSliceTypes_referenceOrder,
		stage.CheckpointINSTimeSliceTypes_instance,
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
		stage.CheckpointINSTypes,
		stage.CheckpointINSType_stagedOrder,
		stage.CheckpointINSTypes_reference,
		&stage.CheckpointINSTypes_referenceOrder,
		stage.CheckpointINSTypes_instance,
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
		stage.CheckpointVORTimeSlicePropertyTypes,
		stage.CheckpointVORTimeSlicePropertyType_stagedOrder,
		stage.CheckpointVORTimeSlicePropertyTypes_reference,
		&stage.CheckpointVORTimeSlicePropertyTypes_referenceOrder,
		stage.CheckpointVORTimeSlicePropertyTypes_instance,
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
		stage.CheckpointVORTimeSliceTypes,
		stage.CheckpointVORTimeSliceType_stagedOrder,
		stage.CheckpointVORTimeSliceTypes_reference,
		&stage.CheckpointVORTimeSliceTypes_referenceOrder,
		stage.CheckpointVORTimeSliceTypes_instance,
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
		stage.CheckpointVORTypes,
		stage.CheckpointVORType_stagedOrder,
		stage.CheckpointVORTypes_reference,
		&stage.CheckpointVORTypes_referenceOrder,
		stage.CheckpointVORTypes_instance,
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
		stage.DMETimeSlicePropertyTypes,
		stage.DMETimeSlicePropertyType_stagedOrder,
		stage.DMETimeSlicePropertyTypes_reference,
		&stage.DMETimeSlicePropertyTypes_referenceOrder,
		stage.DMETimeSlicePropertyTypes_instance,
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
		stage.DMETimeSliceTypes,
		stage.DMETimeSliceType_stagedOrder,
		stage.DMETimeSliceTypes_reference,
		&stage.DMETimeSliceTypes_referenceOrder,
		stage.DMETimeSliceTypes_instance,
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
		stage.DMETypes,
		stage.DMEType_stagedOrder,
		stage.DMETypes_reference,
		&stage.DMETypes_referenceOrder,
		stage.DMETypes_instance,
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
		stage.DesignatedPointTimeSlicePropertyTypes,
		stage.DesignatedPointTimeSlicePropertyType_stagedOrder,
		stage.DesignatedPointTimeSlicePropertyTypes_reference,
		&stage.DesignatedPointTimeSlicePropertyTypes_referenceOrder,
		stage.DesignatedPointTimeSlicePropertyTypes_instance,
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
		stage.DesignatedPointTimeSliceTypes,
		stage.DesignatedPointTimeSliceType_stagedOrder,
		stage.DesignatedPointTimeSliceTypes_reference,
		&stage.DesignatedPointTimeSliceTypes_referenceOrder,
		stage.DesignatedPointTimeSliceTypes_instance,
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
		stage.DesignatedPointTypes,
		stage.DesignatedPointType_stagedOrder,
		stage.DesignatedPointTypes_reference,
		&stage.DesignatedPointTypes_referenceOrder,
		stage.DesignatedPointTypes_instance,
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
		stage.DirectionFinderTimeSlicePropertyTypes,
		stage.DirectionFinderTimeSlicePropertyType_stagedOrder,
		stage.DirectionFinderTimeSlicePropertyTypes_reference,
		&stage.DirectionFinderTimeSlicePropertyTypes_referenceOrder,
		stage.DirectionFinderTimeSlicePropertyTypes_instance,
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
		stage.DirectionFinderTimeSliceTypes,
		stage.DirectionFinderTimeSliceType_stagedOrder,
		stage.DirectionFinderTimeSliceTypes_reference,
		&stage.DirectionFinderTimeSliceTypes_referenceOrder,
		stage.DirectionFinderTimeSliceTypes_instance,
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
		stage.DirectionFinderTypes,
		stage.DirectionFinderType_stagedOrder,
		stage.DirectionFinderTypes_reference,
		&stage.DirectionFinderTypes_referenceOrder,
		stage.DirectionFinderTypes_instance,
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
		stage.DistanceIndicationTimeSlicePropertyTypes,
		stage.DistanceIndicationTimeSlicePropertyType_stagedOrder,
		stage.DistanceIndicationTimeSlicePropertyTypes_reference,
		&stage.DistanceIndicationTimeSlicePropertyTypes_referenceOrder,
		stage.DistanceIndicationTimeSlicePropertyTypes_instance,
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
		stage.DistanceIndicationTimeSliceTypes,
		stage.DistanceIndicationTimeSliceType_stagedOrder,
		stage.DistanceIndicationTimeSliceTypes_reference,
		&stage.DistanceIndicationTimeSliceTypes_referenceOrder,
		stage.DistanceIndicationTimeSliceTypes_instance,
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
		stage.DistanceIndicationTypes,
		stage.DistanceIndicationType_stagedOrder,
		stage.DistanceIndicationTypes_reference,
		&stage.DistanceIndicationTypes_referenceOrder,
		stage.DistanceIndicationTypes_instance,
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
		stage.ElevatedPointPropertyTypes,
		stage.ElevatedPointPropertyType_stagedOrder,
		stage.ElevatedPointPropertyTypes_reference,
		&stage.ElevatedPointPropertyTypes_referenceOrder,
		stage.ElevatedPointPropertyTypes_instance,
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
		stage.ElevationTimeSlicePropertyTypes,
		stage.ElevationTimeSlicePropertyType_stagedOrder,
		stage.ElevationTimeSlicePropertyTypes_reference,
		&stage.ElevationTimeSlicePropertyTypes_referenceOrder,
		stage.ElevationTimeSlicePropertyTypes_instance,
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
		stage.ElevationTimeSliceTypes,
		stage.ElevationTimeSliceType_stagedOrder,
		stage.ElevationTimeSliceTypes_reference,
		&stage.ElevationTimeSliceTypes_referenceOrder,
		stage.ElevationTimeSliceTypes_instance,
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
		stage.ElevationTypes,
		stage.ElevationType_stagedOrder,
		stage.ElevationTypes_reference,
		&stage.ElevationTypes_referenceOrder,
		stage.ElevationTypes_instance,
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
		stage.GlidepathTimeSlicePropertyTypes,
		stage.GlidepathTimeSlicePropertyType_stagedOrder,
		stage.GlidepathTimeSlicePropertyTypes_reference,
		&stage.GlidepathTimeSlicePropertyTypes_referenceOrder,
		stage.GlidepathTimeSlicePropertyTypes_instance,
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
		stage.GlidepathTimeSliceTypes,
		stage.GlidepathTimeSliceType_stagedOrder,
		stage.GlidepathTimeSliceTypes_reference,
		&stage.GlidepathTimeSliceTypes_referenceOrder,
		stage.GlidepathTimeSliceTypes_instance,
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
		stage.GlidepathTypes,
		stage.GlidepathType_stagedOrder,
		stage.GlidepathTypes_reference,
		&stage.GlidepathTypes_referenceOrder,
		stage.GlidepathTypes_instance,
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
		stage.InformationServicePropertyTypes,
		stage.InformationServicePropertyType_stagedOrder,
		stage.InformationServicePropertyTypes_reference,
		&stage.InformationServicePropertyTypes_referenceOrder,
		stage.InformationServicePropertyTypes_instance,
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
		stage.LocalizerTimeSlicePropertyTypes,
		stage.LocalizerTimeSlicePropertyType_stagedOrder,
		stage.LocalizerTimeSlicePropertyTypes_reference,
		&stage.LocalizerTimeSlicePropertyTypes_referenceOrder,
		stage.LocalizerTimeSlicePropertyTypes_instance,
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
		stage.LocalizerTimeSliceTypes,
		stage.LocalizerTimeSliceType_stagedOrder,
		stage.LocalizerTimeSliceTypes_reference,
		&stage.LocalizerTimeSliceTypes_referenceOrder,
		stage.LocalizerTimeSliceTypes_instance,
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
		stage.LocalizerTypes,
		stage.LocalizerType_stagedOrder,
		stage.LocalizerTypes_reference,
		&stage.LocalizerTypes_referenceOrder,
		stage.LocalizerTypes_instance,
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
		stage.MarkerBeaconTimeSlicePropertyTypes,
		stage.MarkerBeaconTimeSlicePropertyType_stagedOrder,
		stage.MarkerBeaconTimeSlicePropertyTypes_reference,
		&stage.MarkerBeaconTimeSlicePropertyTypes_referenceOrder,
		stage.MarkerBeaconTimeSlicePropertyTypes_instance,
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
		stage.MarkerBeaconTimeSliceTypes,
		stage.MarkerBeaconTimeSliceType_stagedOrder,
		stage.MarkerBeaconTimeSliceTypes_reference,
		&stage.MarkerBeaconTimeSliceTypes_referenceOrder,
		stage.MarkerBeaconTimeSliceTypes_instance,
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
		stage.MarkerBeaconTypes,
		stage.MarkerBeaconType_stagedOrder,
		stage.MarkerBeaconTypes_reference,
		&stage.MarkerBeaconTypes_referenceOrder,
		stage.MarkerBeaconTypes_instance,
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
		stage.NDBTimeSlicePropertyTypes,
		stage.NDBTimeSlicePropertyType_stagedOrder,
		stage.NDBTimeSlicePropertyTypes_reference,
		&stage.NDBTimeSlicePropertyTypes_referenceOrder,
		stage.NDBTimeSlicePropertyTypes_instance,
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
		stage.NDBTimeSliceTypes,
		stage.NDBTimeSliceType_stagedOrder,
		stage.NDBTimeSliceTypes_reference,
		&stage.NDBTimeSliceTypes_referenceOrder,
		stage.NDBTimeSliceTypes_instance,
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
		stage.NDBTypes,
		stage.NDBType_stagedOrder,
		stage.NDBTypes_reference,
		&stage.NDBTypes_referenceOrder,
		stage.NDBTypes_instance,
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
		stage.NavaidComponentPropertyTypes,
		stage.NavaidComponentPropertyType_stagedOrder,
		stage.NavaidComponentPropertyTypes_reference,
		&stage.NavaidComponentPropertyTypes_referenceOrder,
		stage.NavaidComponentPropertyTypes_instance,
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
		stage.NavaidComponentTypes,
		stage.NavaidComponentType_stagedOrder,
		stage.NavaidComponentTypes_reference,
		&stage.NavaidComponentTypes_referenceOrder,
		stage.NavaidComponentTypes_instance,
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
		stage.NavaidEquipmentMonitoringPropertyTypes,
		stage.NavaidEquipmentMonitoringPropertyType_stagedOrder,
		stage.NavaidEquipmentMonitoringPropertyTypes_reference,
		&stage.NavaidEquipmentMonitoringPropertyTypes_referenceOrder,
		stage.NavaidEquipmentMonitoringPropertyTypes_instance,
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
		stage.NavaidEquipmentMonitoringTypes,
		stage.NavaidEquipmentMonitoringType_stagedOrder,
		stage.NavaidEquipmentMonitoringTypes_reference,
		&stage.NavaidEquipmentMonitoringTypes_referenceOrder,
		stage.NavaidEquipmentMonitoringTypes_instance,
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
		stage.NavaidEquipmentPropertyTypes,
		stage.NavaidEquipmentPropertyType_stagedOrder,
		stage.NavaidEquipmentPropertyTypes_reference,
		&stage.NavaidEquipmentPropertyTypes_referenceOrder,
		stage.NavaidEquipmentPropertyTypes_instance,
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
		stage.NavaidOperationalStatusPropertyTypes,
		stage.NavaidOperationalStatusPropertyType_stagedOrder,
		stage.NavaidOperationalStatusPropertyTypes_reference,
		&stage.NavaidOperationalStatusPropertyTypes_referenceOrder,
		stage.NavaidOperationalStatusPropertyTypes_instance,
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
		stage.NavaidOperationalStatusTypes,
		stage.NavaidOperationalStatusType_stagedOrder,
		stage.NavaidOperationalStatusTypes_reference,
		&stage.NavaidOperationalStatusTypes_referenceOrder,
		stage.NavaidOperationalStatusTypes_instance,
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
		stage.NavaidTimeSlicePropertyTypes,
		stage.NavaidTimeSlicePropertyType_stagedOrder,
		stage.NavaidTimeSlicePropertyTypes_reference,
		&stage.NavaidTimeSlicePropertyTypes_referenceOrder,
		stage.NavaidTimeSlicePropertyTypes_instance,
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
		stage.NavaidTimeSliceTypes,
		stage.NavaidTimeSliceType_stagedOrder,
		stage.NavaidTimeSliceTypes_reference,
		&stage.NavaidTimeSliceTypes_referenceOrder,
		stage.NavaidTimeSliceTypes_instance,
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
		stage.NavaidTypes,
		stage.NavaidType_stagedOrder,
		stage.NavaidTypes_reference,
		&stage.NavaidTypes_referenceOrder,
		stage.NavaidTypes_instance,
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
		stage.PrecisionApproachRadarTimeSlicePropertyTypes,
		stage.PrecisionApproachRadarTimeSlicePropertyType_stagedOrder,
		stage.PrecisionApproachRadarTimeSlicePropertyTypes_reference,
		&stage.PrecisionApproachRadarTimeSlicePropertyTypes_referenceOrder,
		stage.PrecisionApproachRadarTimeSlicePropertyTypes_instance,
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
		stage.PrecisionApproachRadarTimeSliceTypes,
		stage.PrecisionApproachRadarTimeSliceType_stagedOrder,
		stage.PrecisionApproachRadarTimeSliceTypes_reference,
		&stage.PrecisionApproachRadarTimeSliceTypes_referenceOrder,
		stage.PrecisionApproachRadarTimeSliceTypes_instance,
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
		stage.PrecisionApproachRadarTypes,
		stage.PrecisionApproachRadarType_stagedOrder,
		stage.PrecisionApproachRadarTypes_reference,
		&stage.PrecisionApproachRadarTypes_referenceOrder,
		stage.PrecisionApproachRadarTypes_instance,
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
		stage.PrimarySurveillanceRadarTimeSlicePropertyTypes,
		stage.PrimarySurveillanceRadarTimeSlicePropertyType_stagedOrder,
		stage.PrimarySurveillanceRadarTimeSlicePropertyTypes_reference,
		&stage.PrimarySurveillanceRadarTimeSlicePropertyTypes_referenceOrder,
		stage.PrimarySurveillanceRadarTimeSlicePropertyTypes_instance,
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
		stage.PrimarySurveillanceRadarTimeSliceTypes,
		stage.PrimarySurveillanceRadarTimeSliceType_stagedOrder,
		stage.PrimarySurveillanceRadarTimeSliceTypes_reference,
		&stage.PrimarySurveillanceRadarTimeSliceTypes_referenceOrder,
		stage.PrimarySurveillanceRadarTimeSliceTypes_instance,
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
		stage.PrimarySurveillanceRadarTypes,
		stage.PrimarySurveillanceRadarType_stagedOrder,
		stage.PrimarySurveillanceRadarTypes_reference,
		&stage.PrimarySurveillanceRadarTypes_referenceOrder,
		stage.PrimarySurveillanceRadarTypes_instance,
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
		stage.RadarComponentPropertyTypes,
		stage.RadarComponentPropertyType_stagedOrder,
		stage.RadarComponentPropertyTypes_reference,
		&stage.RadarComponentPropertyTypes_referenceOrder,
		stage.RadarComponentPropertyTypes_instance,
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
		stage.RadarComponentTypes,
		stage.RadarComponentType_stagedOrder,
		stage.RadarComponentTypes_reference,
		&stage.RadarComponentTypes_referenceOrder,
		stage.RadarComponentTypes_instance,
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
		stage.RadarEquipmentPropertyTypes,
		stage.RadarEquipmentPropertyType_stagedOrder,
		stage.RadarEquipmentPropertyTypes_reference,
		&stage.RadarEquipmentPropertyTypes_referenceOrder,
		stage.RadarEquipmentPropertyTypes_instance,
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
		stage.RadarSystemTimeSlicePropertyTypes,
		stage.RadarSystemTimeSlicePropertyType_stagedOrder,
		stage.RadarSystemTimeSlicePropertyTypes_reference,
		&stage.RadarSystemTimeSlicePropertyTypes_referenceOrder,
		stage.RadarSystemTimeSlicePropertyTypes_instance,
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
		stage.RadarSystemTimeSliceTypes,
		stage.RadarSystemTimeSliceType_stagedOrder,
		stage.RadarSystemTimeSliceTypes_reference,
		&stage.RadarSystemTimeSliceTypes_referenceOrder,
		stage.RadarSystemTimeSliceTypes_instance,
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
		stage.RadarSystemTypes,
		stage.RadarSystemType_stagedOrder,
		stage.RadarSystemTypes_reference,
		&stage.RadarSystemTypes_referenceOrder,
		stage.RadarSystemTypes_instance,
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
		stage.ReflectorPropertyTypes,
		stage.ReflectorPropertyType_stagedOrder,
		stage.ReflectorPropertyTypes_reference,
		&stage.ReflectorPropertyTypes_referenceOrder,
		stage.ReflectorPropertyTypes_instance,
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
		stage.ReflectorTypes,
		stage.ReflectorType_stagedOrder,
		stage.ReflectorTypes_reference,
		&stage.ReflectorTypes_referenceOrder,
		stage.ReflectorTypes_instance,
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
		stage.RunwayDirectionPropertyTypes,
		stage.RunwayDirectionPropertyType_stagedOrder,
		stage.RunwayDirectionPropertyTypes_reference,
		&stage.RunwayDirectionPropertyTypes_referenceOrder,
		stage.RunwayDirectionPropertyTypes_instance,
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
		stage.RunwayPropertyTypes,
		stage.RunwayPropertyType_stagedOrder,
		stage.RunwayPropertyTypes_reference,
		&stage.RunwayPropertyTypes_referenceOrder,
		stage.RunwayPropertyTypes_instance,
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
		stage.SDFTimeSlicePropertyTypes,
		stage.SDFTimeSlicePropertyType_stagedOrder,
		stage.SDFTimeSlicePropertyTypes_reference,
		&stage.SDFTimeSlicePropertyTypes_referenceOrder,
		stage.SDFTimeSlicePropertyTypes_instance,
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
		stage.SDFTimeSliceTypes,
		stage.SDFTimeSliceType_stagedOrder,
		stage.SDFTimeSliceTypes_reference,
		&stage.SDFTimeSliceTypes_referenceOrder,
		stage.SDFTimeSliceTypes_instance,
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
		stage.SDFTypes,
		stage.SDFType_stagedOrder,
		stage.SDFTypes_reference,
		&stage.SDFTypes_referenceOrder,
		stage.SDFTypes_instance,
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
		stage.SecondarySurveillanceRadarTimeSlicePropertyTypes,
		stage.SecondarySurveillanceRadarTimeSlicePropertyType_stagedOrder,
		stage.SecondarySurveillanceRadarTimeSlicePropertyTypes_reference,
		&stage.SecondarySurveillanceRadarTimeSlicePropertyTypes_referenceOrder,
		stage.SecondarySurveillanceRadarTimeSlicePropertyTypes_instance,
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
		stage.SecondarySurveillanceRadarTimeSliceTypes,
		stage.SecondarySurveillanceRadarTimeSliceType_stagedOrder,
		stage.SecondarySurveillanceRadarTimeSliceTypes_reference,
		&stage.SecondarySurveillanceRadarTimeSliceTypes_referenceOrder,
		stage.SecondarySurveillanceRadarTimeSliceTypes_instance,
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
		stage.SecondarySurveillanceRadarTypes,
		stage.SecondarySurveillanceRadarType_stagedOrder,
		stage.SecondarySurveillanceRadarTypes_reference,
		&stage.SecondarySurveillanceRadarTypes_referenceOrder,
		stage.SecondarySurveillanceRadarTypes_instance,
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
		stage.SignificantPointInAirspaceTimeSlicePropertyTypes,
		stage.SignificantPointInAirspaceTimeSlicePropertyType_stagedOrder,
		stage.SignificantPointInAirspaceTimeSlicePropertyTypes_reference,
		&stage.SignificantPointInAirspaceTimeSlicePropertyTypes_referenceOrder,
		stage.SignificantPointInAirspaceTimeSlicePropertyTypes_instance,
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
		stage.SignificantPointInAirspaceTimeSliceTypes,
		stage.SignificantPointInAirspaceTimeSliceType_stagedOrder,
		stage.SignificantPointInAirspaceTimeSliceTypes_reference,
		&stage.SignificantPointInAirspaceTimeSliceTypes_referenceOrder,
		stage.SignificantPointInAirspaceTimeSliceTypes_instance,
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
		stage.SignificantPointInAirspaceTypes,
		stage.SignificantPointInAirspaceType_stagedOrder,
		stage.SignificantPointInAirspaceTypes_reference,
		&stage.SignificantPointInAirspaceTypes_referenceOrder,
		stage.SignificantPointInAirspaceTypes_instance,
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
		stage.SpecialNavigationStationStatusPropertyTypes,
		stage.SpecialNavigationStationStatusPropertyType_stagedOrder,
		stage.SpecialNavigationStationStatusPropertyTypes_reference,
		&stage.SpecialNavigationStationStatusPropertyTypes_referenceOrder,
		stage.SpecialNavigationStationStatusPropertyTypes_instance,
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
		stage.SpecialNavigationStationStatusTypes,
		stage.SpecialNavigationStationStatusType_stagedOrder,
		stage.SpecialNavigationStationStatusTypes_reference,
		&stage.SpecialNavigationStationStatusTypes_referenceOrder,
		stage.SpecialNavigationStationStatusTypes_instance,
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
		stage.SpecialNavigationStationTimeSlicePropertyTypes,
		stage.SpecialNavigationStationTimeSlicePropertyType_stagedOrder,
		stage.SpecialNavigationStationTimeSlicePropertyTypes_reference,
		&stage.SpecialNavigationStationTimeSlicePropertyTypes_referenceOrder,
		stage.SpecialNavigationStationTimeSlicePropertyTypes_instance,
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
		stage.SpecialNavigationStationTimeSliceTypes,
		stage.SpecialNavigationStationTimeSliceType_stagedOrder,
		stage.SpecialNavigationStationTimeSliceTypes_reference,
		&stage.SpecialNavigationStationTimeSliceTypes_referenceOrder,
		stage.SpecialNavigationStationTimeSliceTypes_instance,
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
		stage.SpecialNavigationStationTypes,
		stage.SpecialNavigationStationType_stagedOrder,
		stage.SpecialNavigationStationTypes_reference,
		&stage.SpecialNavigationStationTypes_referenceOrder,
		stage.SpecialNavigationStationTypes_instance,
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
		stage.SpecialNavigationSystemPropertyTypes,
		stage.SpecialNavigationSystemPropertyType_stagedOrder,
		stage.SpecialNavigationSystemPropertyTypes_reference,
		&stage.SpecialNavigationSystemPropertyTypes_referenceOrder,
		stage.SpecialNavigationSystemPropertyTypes_instance,
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
		stage.SpecialNavigationSystemTimeSlicePropertyTypes,
		stage.SpecialNavigationSystemTimeSlicePropertyType_stagedOrder,
		stage.SpecialNavigationSystemTimeSlicePropertyTypes_reference,
		&stage.SpecialNavigationSystemTimeSlicePropertyTypes_referenceOrder,
		stage.SpecialNavigationSystemTimeSlicePropertyTypes_instance,
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
		stage.SpecialNavigationSystemTimeSliceTypes,
		stage.SpecialNavigationSystemTimeSliceType_stagedOrder,
		stage.SpecialNavigationSystemTimeSliceTypes_reference,
		&stage.SpecialNavigationSystemTimeSliceTypes_referenceOrder,
		stage.SpecialNavigationSystemTimeSliceTypes_instance,
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
		stage.SpecialNavigationSystemTypes,
		stage.SpecialNavigationSystemType_stagedOrder,
		stage.SpecialNavigationSystemTypes_reference,
		&stage.SpecialNavigationSystemTypes_referenceOrder,
		stage.SpecialNavigationSystemTypes_instance,
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
		stage.SurveillanceGroundStationPropertyTypes,
		stage.SurveillanceGroundStationPropertyType_stagedOrder,
		stage.SurveillanceGroundStationPropertyTypes_reference,
		&stage.SurveillanceGroundStationPropertyTypes_referenceOrder,
		stage.SurveillanceGroundStationPropertyTypes_instance,
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
		stage.SurveillanceGroundStationTypes,
		stage.SurveillanceGroundStationType_stagedOrder,
		stage.SurveillanceGroundStationTypes_reference,
		&stage.SurveillanceGroundStationTypes_referenceOrder,
		stage.SurveillanceGroundStationTypes_instance,
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
		stage.TACANTimeSlicePropertyTypes,
		stage.TACANTimeSlicePropertyType_stagedOrder,
		stage.TACANTimeSlicePropertyTypes_reference,
		&stage.TACANTimeSlicePropertyTypes_referenceOrder,
		stage.TACANTimeSlicePropertyTypes_instance,
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
		stage.TACANTimeSliceTypes,
		stage.TACANTimeSliceType_stagedOrder,
		stage.TACANTimeSliceTypes_reference,
		&stage.TACANTimeSliceTypes_referenceOrder,
		stage.TACANTimeSliceTypes_instance,
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
		stage.TACANTypes,
		stage.TACANType_stagedOrder,
		stage.TACANTypes_reference,
		&stage.TACANTypes_referenceOrder,
		stage.TACANTypes_instance,
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
		stage.UnitPropertyTypes,
		stage.UnitPropertyType_stagedOrder,
		stage.UnitPropertyTypes_reference,
		&stage.UnitPropertyTypes_referenceOrder,
		stage.UnitPropertyTypes_instance,
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
		stage.VORPropertyTypes,
		stage.VORPropertyType_stagedOrder,
		stage.VORPropertyTypes_reference,
		&stage.VORPropertyTypes_referenceOrder,
		stage.VORPropertyTypes_instance,
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
		stage.VORTimeSlicePropertyTypes,
		stage.VORTimeSlicePropertyType_stagedOrder,
		stage.VORTimeSlicePropertyTypes_reference,
		&stage.VORTimeSlicePropertyTypes_referenceOrder,
		stage.VORTimeSlicePropertyTypes_instance,
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
		stage.VORTimeSliceTypes,
		stage.VORTimeSliceType_stagedOrder,
		stage.VORTimeSliceTypes_reference,
		&stage.VORTimeSliceTypes_referenceOrder,
		stage.VORTimeSliceTypes_instance,
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
		stage.VORTypes,
		stage.VORType_stagedOrder,
		stage.VORTypes_reference,
		&stage.VORTypes_referenceOrder,
		stage.VORTypes_instance,
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

	__gong__computeReferencePass1(stage, stage.A_extension_5s, &stage.A_extension_5s_reference, &stage.A_extension_5s_referenceOrder, &stage.A_extension_5s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_6s, &stage.A_extension_6s_reference, &stage.A_extension_6s_referenceOrder, &stage.A_extension_6s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_7s, &stage.A_extension_7s_reference, &stage.A_extension_7s_referenceOrder, &stage.A_extension_7s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_8s, &stage.A_extension_8s_reference, &stage.A_extension_8s_referenceOrder, &stage.A_extension_8s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_9s, &stage.A_extension_9s_reference, &stage.A_extension_9s_referenceOrder, &stage.A_extension_9s_instance)

	__gong__computeReferencePass1(stage, stage.AbstractNavaidEquipmentTypes, &stage.AbstractNavaidEquipmentTypes_reference, &stage.AbstractNavaidEquipmentTypes_referenceOrder, &stage.AbstractNavaidEquipmentTypes_instance)

	__gong__computeReferencePass1(stage, stage.AbstractNavigationSystemCheckpointTypes, &stage.AbstractNavigationSystemCheckpointTypes_reference, &stage.AbstractNavigationSystemCheckpointTypes_referenceOrder, &stage.AbstractNavigationSystemCheckpointTypes_instance)

	__gong__computeReferencePass1(stage, stage.AbstractPropertiesWithScheduleTypes, &stage.AbstractPropertiesWithScheduleTypes_reference, &stage.AbstractPropertiesWithScheduleTypes_referenceOrder, &stage.AbstractPropertiesWithScheduleTypes_instance)

	__gong__computeReferencePass1(stage, stage.AbstractRadarEquipmentTypes, &stage.AbstractRadarEquipmentTypes_reference, &stage.AbstractRadarEquipmentTypes_referenceOrder, &stage.AbstractRadarEquipmentTypes_instance)

	__gong__computeReferencePass1(stage, stage.AbstractSurveillanceRadarTypes, &stage.AbstractSurveillanceRadarTypes_reference, &stage.AbstractSurveillanceRadarTypes_referenceOrder, &stage.AbstractSurveillanceRadarTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirportHeliportPropertyTypes, &stage.AirportHeliportPropertyTypes_reference, &stage.AirportHeliportPropertyTypes_referenceOrder, &stage.AirportHeliportPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirspacePropertyTypes, &stage.AirspacePropertyTypes_reference, &stage.AirspacePropertyTypes_referenceOrder, &stage.AirspacePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AngleIndicationTimeSlicePropertyTypes, &stage.AngleIndicationTimeSlicePropertyTypes_reference, &stage.AngleIndicationTimeSlicePropertyTypes_referenceOrder, &stage.AngleIndicationTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AngleIndicationTimeSliceTypes, &stage.AngleIndicationTimeSliceTypes_reference, &stage.AngleIndicationTimeSliceTypes_referenceOrder, &stage.AngleIndicationTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.AngleIndicationTypes, &stage.AngleIndicationTypes_reference, &stage.AngleIndicationTypes_referenceOrder, &stage.AngleIndicationTypes_instance)

	__gong__computeReferencePass1(stage, stage.AuthorityForNavaidEquipmentPropertyTypes, &stage.AuthorityForNavaidEquipmentPropertyTypes_reference, &stage.AuthorityForNavaidEquipmentPropertyTypes_referenceOrder, &stage.AuthorityForNavaidEquipmentPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AuthorityForNavaidEquipmentTypes, &stage.AuthorityForNavaidEquipmentTypes_reference, &stage.AuthorityForNavaidEquipmentTypes_referenceOrder, &stage.AuthorityForNavaidEquipmentTypes_instance)

	__gong__computeReferencePass1(stage, stage.AuthorityForSpecialNavigationStationPropertyTypes, &stage.AuthorityForSpecialNavigationStationPropertyTypes_reference, &stage.AuthorityForSpecialNavigationStationPropertyTypes_referenceOrder, &stage.AuthorityForSpecialNavigationStationPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AuthorityForSpecialNavigationStationTypes, &stage.AuthorityForSpecialNavigationStationTypes_reference, &stage.AuthorityForSpecialNavigationStationTypes_referenceOrder, &stage.AuthorityForSpecialNavigationStationTypes_instance)

	__gong__computeReferencePass1(stage, stage.AuthorityForSpecialNavigationSystemPropertyTypes, &stage.AuthorityForSpecialNavigationSystemPropertyTypes_reference, &stage.AuthorityForSpecialNavigationSystemPropertyTypes_referenceOrder, &stage.AuthorityForSpecialNavigationSystemPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AuthorityForSpecialNavigationSystemTypes, &stage.AuthorityForSpecialNavigationSystemTypes_reference, &stage.AuthorityForSpecialNavigationSystemTypes_referenceOrder, &stage.AuthorityForSpecialNavigationSystemTypes_instance)

	__gong__computeReferencePass1(stage, stage.AzimuthTimeSlicePropertyTypes, &stage.AzimuthTimeSlicePropertyTypes_reference, &stage.AzimuthTimeSlicePropertyTypes_referenceOrder, &stage.AzimuthTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AzimuthTimeSliceTypes, &stage.AzimuthTimeSliceTypes_reference, &stage.AzimuthTimeSliceTypes_referenceOrder, &stage.AzimuthTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.AzimuthTypes, &stage.AzimuthTypes_reference, &stage.AzimuthTypes_referenceOrder, &stage.AzimuthTypes_instance)

	__gong__computeReferencePass1(stage, stage.CheckpointINSTimeSlicePropertyTypes, &stage.CheckpointINSTimeSlicePropertyTypes_reference, &stage.CheckpointINSTimeSlicePropertyTypes_referenceOrder, &stage.CheckpointINSTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.CheckpointINSTimeSliceTypes, &stage.CheckpointINSTimeSliceTypes_reference, &stage.CheckpointINSTimeSliceTypes_referenceOrder, &stage.CheckpointINSTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CheckpointINSTypes, &stage.CheckpointINSTypes_reference, &stage.CheckpointINSTypes_referenceOrder, &stage.CheckpointINSTypes_instance)

	__gong__computeReferencePass1(stage, stage.CheckpointVORTimeSlicePropertyTypes, &stage.CheckpointVORTimeSlicePropertyTypes_reference, &stage.CheckpointVORTimeSlicePropertyTypes_referenceOrder, &stage.CheckpointVORTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.CheckpointVORTimeSliceTypes, &stage.CheckpointVORTimeSliceTypes_reference, &stage.CheckpointVORTimeSliceTypes_referenceOrder, &stage.CheckpointVORTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CheckpointVORTypes, &stage.CheckpointVORTypes_reference, &stage.CheckpointVORTypes_referenceOrder, &stage.CheckpointVORTypes_instance)

	__gong__computeReferencePass1(stage, stage.ContactInformationPropertyTypes, &stage.ContactInformationPropertyTypes_reference, &stage.ContactInformationPropertyTypes_referenceOrder, &stage.ContactInformationPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ContactInformationTypes, &stage.ContactInformationTypes_reference, &stage.ContactInformationTypes_referenceOrder, &stage.ContactInformationTypes_instance)

	__gong__computeReferencePass1(stage, stage.DMETimeSlicePropertyTypes, &stage.DMETimeSlicePropertyTypes_reference, &stage.DMETimeSlicePropertyTypes_referenceOrder, &stage.DMETimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.DMETimeSliceTypes, &stage.DMETimeSliceTypes_reference, &stage.DMETimeSliceTypes_referenceOrder, &stage.DMETimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.DMETypes, &stage.DMETypes_reference, &stage.DMETypes_referenceOrder, &stage.DMETypes_instance)

	__gong__computeReferencePass1(stage, stage.DesignatedPointPropertyTypes, &stage.DesignatedPointPropertyTypes_reference, &stage.DesignatedPointPropertyTypes_referenceOrder, &stage.DesignatedPointPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.DesignatedPointTimeSlicePropertyTypes, &stage.DesignatedPointTimeSlicePropertyTypes_reference, &stage.DesignatedPointTimeSlicePropertyTypes_referenceOrder, &stage.DesignatedPointTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.DesignatedPointTimeSliceTypes, &stage.DesignatedPointTimeSliceTypes_reference, &stage.DesignatedPointTimeSliceTypes_referenceOrder, &stage.DesignatedPointTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.DesignatedPointTypes, &stage.DesignatedPointTypes_reference, &stage.DesignatedPointTypes_referenceOrder, &stage.DesignatedPointTypes_instance)

	__gong__computeReferencePass1(stage, stage.DirectionFinderTimeSlicePropertyTypes, &stage.DirectionFinderTimeSlicePropertyTypes_reference, &stage.DirectionFinderTimeSlicePropertyTypes_referenceOrder, &stage.DirectionFinderTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.DirectionFinderTimeSliceTypes, &stage.DirectionFinderTimeSliceTypes_reference, &stage.DirectionFinderTimeSliceTypes_referenceOrder, &stage.DirectionFinderTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.DirectionFinderTypes, &stage.DirectionFinderTypes_reference, &stage.DirectionFinderTypes_referenceOrder, &stage.DirectionFinderTypes_instance)

	__gong__computeReferencePass1(stage, stage.DistanceIndicationTimeSlicePropertyTypes, &stage.DistanceIndicationTimeSlicePropertyTypes_reference, &stage.DistanceIndicationTimeSlicePropertyTypes_referenceOrder, &stage.DistanceIndicationTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.DistanceIndicationTimeSliceTypes, &stage.DistanceIndicationTimeSliceTypes_reference, &stage.DistanceIndicationTimeSliceTypes_referenceOrder, &stage.DistanceIndicationTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.DistanceIndicationTypes, &stage.DistanceIndicationTypes_reference, &stage.DistanceIndicationTypes_referenceOrder, &stage.DistanceIndicationTypes_instance)

	__gong__computeReferencePass1(stage, stage.ElevatedPointPropertyTypes, &stage.ElevatedPointPropertyTypes_reference, &stage.ElevatedPointPropertyTypes_referenceOrder, &stage.ElevatedPointPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ElevatedPointTypes, &stage.ElevatedPointTypes_reference, &stage.ElevatedPointTypes_referenceOrder, &stage.ElevatedPointTypes_instance)

	__gong__computeReferencePass1(stage, stage.ElevationTimeSlicePropertyTypes, &stage.ElevationTimeSlicePropertyTypes_reference, &stage.ElevationTimeSlicePropertyTypes_referenceOrder, &stage.ElevationTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ElevationTimeSliceTypes, &stage.ElevationTimeSliceTypes_reference, &stage.ElevationTimeSliceTypes_referenceOrder, &stage.ElevationTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.ElevationTypes, &stage.ElevationTypes_reference, &stage.ElevationTypes_referenceOrder, &stage.ElevationTypes_instance)

	__gong__computeReferencePass1(stage, stage.GlidepathTimeSlicePropertyTypes, &stage.GlidepathTimeSlicePropertyTypes_reference, &stage.GlidepathTimeSlicePropertyTypes_referenceOrder, &stage.GlidepathTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.GlidepathTimeSliceTypes, &stage.GlidepathTimeSliceTypes_reference, &stage.GlidepathTimeSliceTypes_referenceOrder, &stage.GlidepathTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.GlidepathTypes, &stage.GlidepathTypes_reference, &stage.GlidepathTypes_referenceOrder, &stage.GlidepathTypes_instance)

	__gong__computeReferencePass1(stage, stage.InformationServicePropertyTypes, &stage.InformationServicePropertyTypes_reference, &stage.InformationServicePropertyTypes_referenceOrder, &stage.InformationServicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.LinguisticNotePropertyTypes, &stage.LinguisticNotePropertyTypes_reference, &stage.LinguisticNotePropertyTypes_referenceOrder, &stage.LinguisticNotePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.LinguisticNoteTypes, &stage.LinguisticNoteTypes_reference, &stage.LinguisticNoteTypes_referenceOrder, &stage.LinguisticNoteTypes_instance)

	__gong__computeReferencePass1(stage, stage.LocalizerTimeSlicePropertyTypes, &stage.LocalizerTimeSlicePropertyTypes_reference, &stage.LocalizerTimeSlicePropertyTypes_referenceOrder, &stage.LocalizerTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.LocalizerTimeSliceTypes, &stage.LocalizerTimeSliceTypes_reference, &stage.LocalizerTimeSliceTypes_referenceOrder, &stage.LocalizerTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.LocalizerTypes, &stage.LocalizerTypes_reference, &stage.LocalizerTypes_referenceOrder, &stage.LocalizerTypes_instance)

	__gong__computeReferencePass1(stage, stage.MarkerBeaconTimeSlicePropertyTypes, &stage.MarkerBeaconTimeSlicePropertyTypes_reference, &stage.MarkerBeaconTimeSlicePropertyTypes_referenceOrder, &stage.MarkerBeaconTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.MarkerBeaconTimeSliceTypes, &stage.MarkerBeaconTimeSliceTypes_reference, &stage.MarkerBeaconTimeSliceTypes_referenceOrder, &stage.MarkerBeaconTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.MarkerBeaconTypes, &stage.MarkerBeaconTypes_reference, &stage.MarkerBeaconTypes_referenceOrder, &stage.MarkerBeaconTypes_instance)

	__gong__computeReferencePass1(stage, stage.NDBTimeSlicePropertyTypes, &stage.NDBTimeSlicePropertyTypes_reference, &stage.NDBTimeSlicePropertyTypes_referenceOrder, &stage.NDBTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.NDBTimeSliceTypes, &stage.NDBTimeSliceTypes_reference, &stage.NDBTimeSliceTypes_referenceOrder, &stage.NDBTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.NDBTypes, &stage.NDBTypes_reference, &stage.NDBTypes_referenceOrder, &stage.NDBTypes_instance)

	__gong__computeReferencePass1(stage, stage.NavaidComponentPropertyTypes, &stage.NavaidComponentPropertyTypes_reference, &stage.NavaidComponentPropertyTypes_referenceOrder, &stage.NavaidComponentPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.NavaidComponentTypes, &stage.NavaidComponentTypes_reference, &stage.NavaidComponentTypes_referenceOrder, &stage.NavaidComponentTypes_instance)

	__gong__computeReferencePass1(stage, stage.NavaidEquipmentMonitoringPropertyTypes, &stage.NavaidEquipmentMonitoringPropertyTypes_reference, &stage.NavaidEquipmentMonitoringPropertyTypes_referenceOrder, &stage.NavaidEquipmentMonitoringPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.NavaidEquipmentMonitoringTypes, &stage.NavaidEquipmentMonitoringTypes_reference, &stage.NavaidEquipmentMonitoringTypes_referenceOrder, &stage.NavaidEquipmentMonitoringTypes_instance)

	__gong__computeReferencePass1(stage, stage.NavaidEquipmentPropertyTypes, &stage.NavaidEquipmentPropertyTypes_reference, &stage.NavaidEquipmentPropertyTypes_referenceOrder, &stage.NavaidEquipmentPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.NavaidOperationalStatusPropertyTypes, &stage.NavaidOperationalStatusPropertyTypes_reference, &stage.NavaidOperationalStatusPropertyTypes_referenceOrder, &stage.NavaidOperationalStatusPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.NavaidOperationalStatusTypes, &stage.NavaidOperationalStatusTypes_reference, &stage.NavaidOperationalStatusTypes_referenceOrder, &stage.NavaidOperationalStatusTypes_instance)

	__gong__computeReferencePass1(stage, stage.NavaidPropertyTypes, &stage.NavaidPropertyTypes_reference, &stage.NavaidPropertyTypes_referenceOrder, &stage.NavaidPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.NavaidTimeSlicePropertyTypes, &stage.NavaidTimeSlicePropertyTypes_reference, &stage.NavaidTimeSlicePropertyTypes_referenceOrder, &stage.NavaidTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.NavaidTimeSliceTypes, &stage.NavaidTimeSliceTypes_reference, &stage.NavaidTimeSliceTypes_referenceOrder, &stage.NavaidTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.NavaidTypes, &stage.NavaidTypes_reference, &stage.NavaidTypes_referenceOrder, &stage.NavaidTypes_instance)

	__gong__computeReferencePass1(stage, stage.NotePropertyTypes, &stage.NotePropertyTypes_reference, &stage.NotePropertyTypes_referenceOrder, &stage.NotePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.NoteTypes, &stage.NoteTypes_reference, &stage.NoteTypes_referenceOrder, &stage.NoteTypes_instance)

	__gong__computeReferencePass1(stage, stage.OnlineContactPropertyTypes, &stage.OnlineContactPropertyTypes_reference, &stage.OnlineContactPropertyTypes_referenceOrder, &stage.OnlineContactPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.OnlineContactTypes, &stage.OnlineContactTypes_reference, &stage.OnlineContactTypes_referenceOrder, &stage.OnlineContactTypes_instance)

	__gong__computeReferencePass1(stage, stage.OrganisationAuthorityPropertyTypes, &stage.OrganisationAuthorityPropertyTypes_reference, &stage.OrganisationAuthorityPropertyTypes_referenceOrder, &stage.OrganisationAuthorityPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.PointPropertyTypes, &stage.PointPropertyTypes_reference, &stage.PointPropertyTypes_referenceOrder, &stage.PointPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.PointTypes, &stage.PointTypes_reference, &stage.PointTypes_referenceOrder, &stage.PointTypes_instance)

	__gong__computeReferencePass1(stage, stage.PostalAddressPropertyTypes, &stage.PostalAddressPropertyTypes_reference, &stage.PostalAddressPropertyTypes_referenceOrder, &stage.PostalAddressPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.PostalAddressTypes, &stage.PostalAddressTypes_reference, &stage.PostalAddressTypes_referenceOrder, &stage.PostalAddressTypes_instance)

	__gong__computeReferencePass1(stage, stage.PrecisionApproachRadarTimeSlicePropertyTypes, &stage.PrecisionApproachRadarTimeSlicePropertyTypes_reference, &stage.PrecisionApproachRadarTimeSlicePropertyTypes_referenceOrder, &stage.PrecisionApproachRadarTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.PrecisionApproachRadarTimeSliceTypes, &stage.PrecisionApproachRadarTimeSliceTypes_reference, &stage.PrecisionApproachRadarTimeSliceTypes_referenceOrder, &stage.PrecisionApproachRadarTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.PrecisionApproachRadarTypes, &stage.PrecisionApproachRadarTypes_reference, &stage.PrecisionApproachRadarTypes_referenceOrder, &stage.PrecisionApproachRadarTypes_instance)

	__gong__computeReferencePass1(stage, stage.PrimarySurveillanceRadarTimeSlicePropertyTypes, &stage.PrimarySurveillanceRadarTimeSlicePropertyTypes_reference, &stage.PrimarySurveillanceRadarTimeSlicePropertyTypes_referenceOrder, &stage.PrimarySurveillanceRadarTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.PrimarySurveillanceRadarTimeSliceTypes, &stage.PrimarySurveillanceRadarTimeSliceTypes_reference, &stage.PrimarySurveillanceRadarTimeSliceTypes_referenceOrder, &stage.PrimarySurveillanceRadarTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.PrimarySurveillanceRadarTypes, &stage.PrimarySurveillanceRadarTypes_reference, &stage.PrimarySurveillanceRadarTypes_referenceOrder, &stage.PrimarySurveillanceRadarTypes_instance)

	__gong__computeReferencePass1(stage, stage.RadarComponentPropertyTypes, &stage.RadarComponentPropertyTypes_reference, &stage.RadarComponentPropertyTypes_referenceOrder, &stage.RadarComponentPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.RadarComponentTypes, &stage.RadarComponentTypes_reference, &stage.RadarComponentTypes_referenceOrder, &stage.RadarComponentTypes_instance)

	__gong__computeReferencePass1(stage, stage.RadarEquipmentPropertyTypes, &stage.RadarEquipmentPropertyTypes_reference, &stage.RadarEquipmentPropertyTypes_referenceOrder, &stage.RadarEquipmentPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.RadarSystemTimeSlicePropertyTypes, &stage.RadarSystemTimeSlicePropertyTypes_reference, &stage.RadarSystemTimeSlicePropertyTypes_referenceOrder, &stage.RadarSystemTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.RadarSystemTimeSliceTypes, &stage.RadarSystemTimeSliceTypes_reference, &stage.RadarSystemTimeSliceTypes_referenceOrder, &stage.RadarSystemTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.RadarSystemTypes, &stage.RadarSystemTypes_reference, &stage.RadarSystemTypes_referenceOrder, &stage.RadarSystemTypes_instance)

	__gong__computeReferencePass1(stage, stage.ReflectorPropertyTypes, &stage.ReflectorPropertyTypes_reference, &stage.ReflectorPropertyTypes_referenceOrder, &stage.ReflectorPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ReflectorTypes, &stage.ReflectorTypes_reference, &stage.ReflectorTypes_referenceOrder, &stage.ReflectorTypes_instance)

	__gong__computeReferencePass1(stage, stage.RunwayCentrelinePointPropertyTypes, &stage.RunwayCentrelinePointPropertyTypes_reference, &stage.RunwayCentrelinePointPropertyTypes_referenceOrder, &stage.RunwayCentrelinePointPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.RunwayDirectionPropertyTypes, &stage.RunwayDirectionPropertyTypes_reference, &stage.RunwayDirectionPropertyTypes_referenceOrder, &stage.RunwayDirectionPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.RunwayPropertyTypes, &stage.RunwayPropertyTypes_reference, &stage.RunwayPropertyTypes_referenceOrder, &stage.RunwayPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.SDFTimeSlicePropertyTypes, &stage.SDFTimeSlicePropertyTypes_reference, &stage.SDFTimeSlicePropertyTypes_referenceOrder, &stage.SDFTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.SDFTimeSliceTypes, &stage.SDFTimeSliceTypes_reference, &stage.SDFTimeSliceTypes_referenceOrder, &stage.SDFTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.SDFTypes, &stage.SDFTypes_reference, &stage.SDFTypes_referenceOrder, &stage.SDFTypes_instance)

	__gong__computeReferencePass1(stage, stage.SecondarySurveillanceRadarTimeSlicePropertyTypes, &stage.SecondarySurveillanceRadarTimeSlicePropertyTypes_reference, &stage.SecondarySurveillanceRadarTimeSlicePropertyTypes_referenceOrder, &stage.SecondarySurveillanceRadarTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.SecondarySurveillanceRadarTimeSliceTypes, &stage.SecondarySurveillanceRadarTimeSliceTypes_reference, &stage.SecondarySurveillanceRadarTimeSliceTypes_referenceOrder, &stage.SecondarySurveillanceRadarTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.SecondarySurveillanceRadarTypes, &stage.SecondarySurveillanceRadarTypes_reference, &stage.SecondarySurveillanceRadarTypes_referenceOrder, &stage.SecondarySurveillanceRadarTypes_instance)

	__gong__computeReferencePass1(stage, stage.SignificantPointInAirspaceTimeSlicePropertyTypes, &stage.SignificantPointInAirspaceTimeSlicePropertyTypes_reference, &stage.SignificantPointInAirspaceTimeSlicePropertyTypes_referenceOrder, &stage.SignificantPointInAirspaceTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.SignificantPointInAirspaceTimeSliceTypes, &stage.SignificantPointInAirspaceTimeSliceTypes_reference, &stage.SignificantPointInAirspaceTimeSliceTypes_referenceOrder, &stage.SignificantPointInAirspaceTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.SignificantPointInAirspaceTypes, &stage.SignificantPointInAirspaceTypes_reference, &stage.SignificantPointInAirspaceTypes_referenceOrder, &stage.SignificantPointInAirspaceTypes_instance)

	__gong__computeReferencePass1(stage, stage.SpecialNavigationStationStatusPropertyTypes, &stage.SpecialNavigationStationStatusPropertyTypes_reference, &stage.SpecialNavigationStationStatusPropertyTypes_referenceOrder, &stage.SpecialNavigationStationStatusPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.SpecialNavigationStationStatusTypes, &stage.SpecialNavigationStationStatusTypes_reference, &stage.SpecialNavigationStationStatusTypes_referenceOrder, &stage.SpecialNavigationStationStatusTypes_instance)

	__gong__computeReferencePass1(stage, stage.SpecialNavigationStationTimeSlicePropertyTypes, &stage.SpecialNavigationStationTimeSlicePropertyTypes_reference, &stage.SpecialNavigationStationTimeSlicePropertyTypes_referenceOrder, &stage.SpecialNavigationStationTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.SpecialNavigationStationTimeSliceTypes, &stage.SpecialNavigationStationTimeSliceTypes_reference, &stage.SpecialNavigationStationTimeSliceTypes_referenceOrder, &stage.SpecialNavigationStationTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.SpecialNavigationStationTypes, &stage.SpecialNavigationStationTypes_reference, &stage.SpecialNavigationStationTypes_referenceOrder, &stage.SpecialNavigationStationTypes_instance)

	__gong__computeReferencePass1(stage, stage.SpecialNavigationSystemPropertyTypes, &stage.SpecialNavigationSystemPropertyTypes_reference, &stage.SpecialNavigationSystemPropertyTypes_referenceOrder, &stage.SpecialNavigationSystemPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.SpecialNavigationSystemTimeSlicePropertyTypes, &stage.SpecialNavigationSystemTimeSlicePropertyTypes_reference, &stage.SpecialNavigationSystemTimeSlicePropertyTypes_referenceOrder, &stage.SpecialNavigationSystemTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.SpecialNavigationSystemTimeSliceTypes, &stage.SpecialNavigationSystemTimeSliceTypes_reference, &stage.SpecialNavigationSystemTimeSliceTypes_referenceOrder, &stage.SpecialNavigationSystemTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.SpecialNavigationSystemTypes, &stage.SpecialNavigationSystemTypes_reference, &stage.SpecialNavigationSystemTypes_referenceOrder, &stage.SpecialNavigationSystemTypes_instance)

	__gong__computeReferencePass1(stage, stage.SurveillanceGroundStationPropertyTypes, &stage.SurveillanceGroundStationPropertyTypes_reference, &stage.SurveillanceGroundStationPropertyTypes_referenceOrder, &stage.SurveillanceGroundStationPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.SurveillanceGroundStationTypes, &stage.SurveillanceGroundStationTypes_reference, &stage.SurveillanceGroundStationTypes_referenceOrder, &stage.SurveillanceGroundStationTypes_instance)

	__gong__computeReferencePass1(stage, stage.TACANTimeSlicePropertyTypes, &stage.TACANTimeSlicePropertyTypes_reference, &stage.TACANTimeSlicePropertyTypes_referenceOrder, &stage.TACANTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.TACANTimeSliceTypes, &stage.TACANTimeSliceTypes_reference, &stage.TACANTimeSliceTypes_referenceOrder, &stage.TACANTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.TACANTypes, &stage.TACANTypes_reference, &stage.TACANTypes_referenceOrder, &stage.TACANTypes_instance)

	__gong__computeReferencePass1(stage, stage.TelephoneContactPropertyTypes, &stage.TelephoneContactPropertyTypes_reference, &stage.TelephoneContactPropertyTypes_referenceOrder, &stage.TelephoneContactPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.TelephoneContactTypes, &stage.TelephoneContactTypes_reference, &stage.TelephoneContactTypes_referenceOrder, &stage.TelephoneContactTypes_instance)

	__gong__computeReferencePass1(stage, stage.TimesheetPropertyTypes, &stage.TimesheetPropertyTypes_reference, &stage.TimesheetPropertyTypes_referenceOrder, &stage.TimesheetPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.TimesheetTypes, &stage.TimesheetTypes_reference, &stage.TimesheetTypes_referenceOrder, &stage.TimesheetTypes_instance)

	__gong__computeReferencePass1(stage, stage.TouchDownLiftOffPropertyTypes, &stage.TouchDownLiftOffPropertyTypes_reference, &stage.TouchDownLiftOffPropertyTypes_referenceOrder, &stage.TouchDownLiftOffPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.UnitPropertyTypes, &stage.UnitPropertyTypes_reference, &stage.UnitPropertyTypes_referenceOrder, &stage.UnitPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.VORPropertyTypes, &stage.VORPropertyTypes_reference, &stage.VORPropertyTypes_referenceOrder, &stage.VORPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.VORTimeSlicePropertyTypes, &stage.VORTimeSlicePropertyTypes_reference, &stage.VORTimeSlicePropertyTypes_referenceOrder, &stage.VORTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.VORTimeSliceTypes, &stage.VORTimeSliceTypes_reference, &stage.VORTimeSliceTypes_referenceOrder, &stage.VORTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.VORTypes, &stage.VORTypes_reference, &stage.VORTypes_referenceOrder, &stage.VORTypes_instance)

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

	__gong__computeReferencePass2(stage.A_extension_5s, stage.A_extension_5s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_6s, stage.A_extension_6s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_7s, stage.A_extension_7s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_8s, stage.A_extension_8s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_9s, stage.A_extension_9s_reference, stage)

	__gong__computeReferencePass2(stage.AbstractNavaidEquipmentTypes, stage.AbstractNavaidEquipmentTypes_reference, stage)

	__gong__computeReferencePass2(stage.AbstractNavigationSystemCheckpointTypes, stage.AbstractNavigationSystemCheckpointTypes_reference, stage)

	__gong__computeReferencePass2(stage.AbstractPropertiesWithScheduleTypes, stage.AbstractPropertiesWithScheduleTypes_reference, stage)

	__gong__computeReferencePass2(stage.AbstractRadarEquipmentTypes, stage.AbstractRadarEquipmentTypes_reference, stage)

	__gong__computeReferencePass2(stage.AbstractSurveillanceRadarTypes, stage.AbstractSurveillanceRadarTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirportHeliportPropertyTypes, stage.AirportHeliportPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirspacePropertyTypes, stage.AirspacePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AngleIndicationTimeSlicePropertyTypes, stage.AngleIndicationTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AngleIndicationTimeSliceTypes, stage.AngleIndicationTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.AngleIndicationTypes, stage.AngleIndicationTypes_reference, stage)

	__gong__computeReferencePass2(stage.AuthorityForNavaidEquipmentPropertyTypes, stage.AuthorityForNavaidEquipmentPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AuthorityForNavaidEquipmentTypes, stage.AuthorityForNavaidEquipmentTypes_reference, stage)

	__gong__computeReferencePass2(stage.AuthorityForSpecialNavigationStationPropertyTypes, stage.AuthorityForSpecialNavigationStationPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AuthorityForSpecialNavigationStationTypes, stage.AuthorityForSpecialNavigationStationTypes_reference, stage)

	__gong__computeReferencePass2(stage.AuthorityForSpecialNavigationSystemPropertyTypes, stage.AuthorityForSpecialNavigationSystemPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AuthorityForSpecialNavigationSystemTypes, stage.AuthorityForSpecialNavigationSystemTypes_reference, stage)

	__gong__computeReferencePass2(stage.AzimuthTimeSlicePropertyTypes, stage.AzimuthTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AzimuthTimeSliceTypes, stage.AzimuthTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.AzimuthTypes, stage.AzimuthTypes_reference, stage)

	__gong__computeReferencePass2(stage.CheckpointINSTimeSlicePropertyTypes, stage.CheckpointINSTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.CheckpointINSTimeSliceTypes, stage.CheckpointINSTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CheckpointINSTypes, stage.CheckpointINSTypes_reference, stage)

	__gong__computeReferencePass2(stage.CheckpointVORTimeSlicePropertyTypes, stage.CheckpointVORTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.CheckpointVORTimeSliceTypes, stage.CheckpointVORTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CheckpointVORTypes, stage.CheckpointVORTypes_reference, stage)

	__gong__computeReferencePass2(stage.ContactInformationPropertyTypes, stage.ContactInformationPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ContactInformationTypes, stage.ContactInformationTypes_reference, stage)

	__gong__computeReferencePass2(stage.DMETimeSlicePropertyTypes, stage.DMETimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.DMETimeSliceTypes, stage.DMETimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.DMETypes, stage.DMETypes_reference, stage)

	__gong__computeReferencePass2(stage.DesignatedPointPropertyTypes, stage.DesignatedPointPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.DesignatedPointTimeSlicePropertyTypes, stage.DesignatedPointTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.DesignatedPointTimeSliceTypes, stage.DesignatedPointTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.DesignatedPointTypes, stage.DesignatedPointTypes_reference, stage)

	__gong__computeReferencePass2(stage.DirectionFinderTimeSlicePropertyTypes, stage.DirectionFinderTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.DirectionFinderTimeSliceTypes, stage.DirectionFinderTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.DirectionFinderTypes, stage.DirectionFinderTypes_reference, stage)

	__gong__computeReferencePass2(stage.DistanceIndicationTimeSlicePropertyTypes, stage.DistanceIndicationTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.DistanceIndicationTimeSliceTypes, stage.DistanceIndicationTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.DistanceIndicationTypes, stage.DistanceIndicationTypes_reference, stage)

	__gong__computeReferencePass2(stage.ElevatedPointPropertyTypes, stage.ElevatedPointPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ElevatedPointTypes, stage.ElevatedPointTypes_reference, stage)

	__gong__computeReferencePass2(stage.ElevationTimeSlicePropertyTypes, stage.ElevationTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ElevationTimeSliceTypes, stage.ElevationTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.ElevationTypes, stage.ElevationTypes_reference, stage)

	__gong__computeReferencePass2(stage.GlidepathTimeSlicePropertyTypes, stage.GlidepathTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.GlidepathTimeSliceTypes, stage.GlidepathTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.GlidepathTypes, stage.GlidepathTypes_reference, stage)

	__gong__computeReferencePass2(stage.InformationServicePropertyTypes, stage.InformationServicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.LinguisticNotePropertyTypes, stage.LinguisticNotePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.LinguisticNoteTypes, stage.LinguisticNoteTypes_reference, stage)

	__gong__computeReferencePass2(stage.LocalizerTimeSlicePropertyTypes, stage.LocalizerTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.LocalizerTimeSliceTypes, stage.LocalizerTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.LocalizerTypes, stage.LocalizerTypes_reference, stage)

	__gong__computeReferencePass2(stage.MarkerBeaconTimeSlicePropertyTypes, stage.MarkerBeaconTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.MarkerBeaconTimeSliceTypes, stage.MarkerBeaconTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.MarkerBeaconTypes, stage.MarkerBeaconTypes_reference, stage)

	__gong__computeReferencePass2(stage.NDBTimeSlicePropertyTypes, stage.NDBTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.NDBTimeSliceTypes, stage.NDBTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.NDBTypes, stage.NDBTypes_reference, stage)

	__gong__computeReferencePass2(stage.NavaidComponentPropertyTypes, stage.NavaidComponentPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.NavaidComponentTypes, stage.NavaidComponentTypes_reference, stage)

	__gong__computeReferencePass2(stage.NavaidEquipmentMonitoringPropertyTypes, stage.NavaidEquipmentMonitoringPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.NavaidEquipmentMonitoringTypes, stage.NavaidEquipmentMonitoringTypes_reference, stage)

	__gong__computeReferencePass2(stage.NavaidEquipmentPropertyTypes, stage.NavaidEquipmentPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.NavaidOperationalStatusPropertyTypes, stage.NavaidOperationalStatusPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.NavaidOperationalStatusTypes, stage.NavaidOperationalStatusTypes_reference, stage)

	__gong__computeReferencePass2(stage.NavaidPropertyTypes, stage.NavaidPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.NavaidTimeSlicePropertyTypes, stage.NavaidTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.NavaidTimeSliceTypes, stage.NavaidTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.NavaidTypes, stage.NavaidTypes_reference, stage)

	__gong__computeReferencePass2(stage.NotePropertyTypes, stage.NotePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.NoteTypes, stage.NoteTypes_reference, stage)

	__gong__computeReferencePass2(stage.OnlineContactPropertyTypes, stage.OnlineContactPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.OnlineContactTypes, stage.OnlineContactTypes_reference, stage)

	__gong__computeReferencePass2(stage.OrganisationAuthorityPropertyTypes, stage.OrganisationAuthorityPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.PointPropertyTypes, stage.PointPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.PointTypes, stage.PointTypes_reference, stage)

	__gong__computeReferencePass2(stage.PostalAddressPropertyTypes, stage.PostalAddressPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.PostalAddressTypes, stage.PostalAddressTypes_reference, stage)

	__gong__computeReferencePass2(stage.PrecisionApproachRadarTimeSlicePropertyTypes, stage.PrecisionApproachRadarTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.PrecisionApproachRadarTimeSliceTypes, stage.PrecisionApproachRadarTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.PrecisionApproachRadarTypes, stage.PrecisionApproachRadarTypes_reference, stage)

	__gong__computeReferencePass2(stage.PrimarySurveillanceRadarTimeSlicePropertyTypes, stage.PrimarySurveillanceRadarTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.PrimarySurveillanceRadarTimeSliceTypes, stage.PrimarySurveillanceRadarTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.PrimarySurveillanceRadarTypes, stage.PrimarySurveillanceRadarTypes_reference, stage)

	__gong__computeReferencePass2(stage.RadarComponentPropertyTypes, stage.RadarComponentPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.RadarComponentTypes, stage.RadarComponentTypes_reference, stage)

	__gong__computeReferencePass2(stage.RadarEquipmentPropertyTypes, stage.RadarEquipmentPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.RadarSystemTimeSlicePropertyTypes, stage.RadarSystemTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.RadarSystemTimeSliceTypes, stage.RadarSystemTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.RadarSystemTypes, stage.RadarSystemTypes_reference, stage)

	__gong__computeReferencePass2(stage.ReflectorPropertyTypes, stage.ReflectorPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ReflectorTypes, stage.ReflectorTypes_reference, stage)

	__gong__computeReferencePass2(stage.RunwayCentrelinePointPropertyTypes, stage.RunwayCentrelinePointPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.RunwayDirectionPropertyTypes, stage.RunwayDirectionPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.RunwayPropertyTypes, stage.RunwayPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.SDFTimeSlicePropertyTypes, stage.SDFTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.SDFTimeSliceTypes, stage.SDFTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.SDFTypes, stage.SDFTypes_reference, stage)

	__gong__computeReferencePass2(stage.SecondarySurveillanceRadarTimeSlicePropertyTypes, stage.SecondarySurveillanceRadarTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.SecondarySurveillanceRadarTimeSliceTypes, stage.SecondarySurveillanceRadarTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.SecondarySurveillanceRadarTypes, stage.SecondarySurveillanceRadarTypes_reference, stage)

	__gong__computeReferencePass2(stage.SignificantPointInAirspaceTimeSlicePropertyTypes, stage.SignificantPointInAirspaceTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.SignificantPointInAirspaceTimeSliceTypes, stage.SignificantPointInAirspaceTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.SignificantPointInAirspaceTypes, stage.SignificantPointInAirspaceTypes_reference, stage)

	__gong__computeReferencePass2(stage.SpecialNavigationStationStatusPropertyTypes, stage.SpecialNavigationStationStatusPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.SpecialNavigationStationStatusTypes, stage.SpecialNavigationStationStatusTypes_reference, stage)

	__gong__computeReferencePass2(stage.SpecialNavigationStationTimeSlicePropertyTypes, stage.SpecialNavigationStationTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.SpecialNavigationStationTimeSliceTypes, stage.SpecialNavigationStationTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.SpecialNavigationStationTypes, stage.SpecialNavigationStationTypes_reference, stage)

	__gong__computeReferencePass2(stage.SpecialNavigationSystemPropertyTypes, stage.SpecialNavigationSystemPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.SpecialNavigationSystemTimeSlicePropertyTypes, stage.SpecialNavigationSystemTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.SpecialNavigationSystemTimeSliceTypes, stage.SpecialNavigationSystemTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.SpecialNavigationSystemTypes, stage.SpecialNavigationSystemTypes_reference, stage)

	__gong__computeReferencePass2(stage.SurveillanceGroundStationPropertyTypes, stage.SurveillanceGroundStationPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.SurveillanceGroundStationTypes, stage.SurveillanceGroundStationTypes_reference, stage)

	__gong__computeReferencePass2(stage.TACANTimeSlicePropertyTypes, stage.TACANTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.TACANTimeSliceTypes, stage.TACANTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.TACANTypes, stage.TACANTypes_reference, stage)

	__gong__computeReferencePass2(stage.TelephoneContactPropertyTypes, stage.TelephoneContactPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.TelephoneContactTypes, stage.TelephoneContactTypes_reference, stage)

	__gong__computeReferencePass2(stage.TimesheetPropertyTypes, stage.TimesheetPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.TimesheetTypes, stage.TimesheetTypes_reference, stage)

	__gong__computeReferencePass2(stage.TouchDownLiftOffPropertyTypes, stage.TouchDownLiftOffPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.UnitPropertyTypes, stage.UnitPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.VORPropertyTypes, stage.VORPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.VORTimeSlicePropertyTypes, stage.VORTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.VORTimeSliceTypes, stage.VORTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.VORTypes, stage.VORTypes_reference, stage)

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

func (abstractnavaidequipmenttype *AbstractNavaidEquipmentType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractNavaidEquipmentType_stagedOrder, stage.AbstractNavaidEquipmentTypes_referenceOrder, abstractnavaidequipmenttype, "AbstractNavaidEquipmentType")
}

func (abstractnavigationsystemcheckpointtype *AbstractNavigationSystemCheckpointType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractNavigationSystemCheckpointType_stagedOrder, stage.AbstractNavigationSystemCheckpointTypes_referenceOrder, abstractnavigationsystemcheckpointtype, "AbstractNavigationSystemCheckpointType")
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractPropertiesWithScheduleType_stagedOrder, stage.AbstractPropertiesWithScheduleTypes_referenceOrder, abstractpropertieswithscheduletype, "AbstractPropertiesWithScheduleType")
}

func (abstractradarequipmenttype *AbstractRadarEquipmentType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractRadarEquipmentType_stagedOrder, stage.AbstractRadarEquipmentTypes_referenceOrder, abstractradarequipmenttype, "AbstractRadarEquipmentType")
}

func (abstractsurveillanceradartype *AbstractSurveillanceRadarType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractSurveillanceRadarType_stagedOrder, stage.AbstractSurveillanceRadarTypes_referenceOrder, abstractsurveillanceradartype, "AbstractSurveillanceRadarType")
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirportHeliportPropertyType_stagedOrder, stage.AirportHeliportPropertyTypes_referenceOrder, airportheliportpropertytype, "AirportHeliportPropertyType")
}

func (airspacepropertytype *AirspacePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirspacePropertyType_stagedOrder, stage.AirspacePropertyTypes_referenceOrder, airspacepropertytype, "AirspacePropertyType")
}

func (angleindicationtimeslicepropertytype *AngleIndicationTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AngleIndicationTimeSlicePropertyType_stagedOrder, stage.AngleIndicationTimeSlicePropertyTypes_referenceOrder, angleindicationtimeslicepropertytype, "AngleIndicationTimeSlicePropertyType")
}

func (angleindicationtimeslicetype *AngleIndicationTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AngleIndicationTimeSliceType_stagedOrder, stage.AngleIndicationTimeSliceTypes_referenceOrder, angleindicationtimeslicetype, "AngleIndicationTimeSliceType")
}

func (angleindicationtype *AngleIndicationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AngleIndicationType_stagedOrder, stage.AngleIndicationTypes_referenceOrder, angleindicationtype, "AngleIndicationType")
}

func (authorityfornavaidequipmentpropertytype *AuthorityForNavaidEquipmentPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AuthorityForNavaidEquipmentPropertyType_stagedOrder, stage.AuthorityForNavaidEquipmentPropertyTypes_referenceOrder, authorityfornavaidequipmentpropertytype, "AuthorityForNavaidEquipmentPropertyType")
}

func (authorityfornavaidequipmenttype *AuthorityForNavaidEquipmentType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AuthorityForNavaidEquipmentType_stagedOrder, stage.AuthorityForNavaidEquipmentTypes_referenceOrder, authorityfornavaidequipmenttype, "AuthorityForNavaidEquipmentType")
}

func (authorityforspecialnavigationstationpropertytype *AuthorityForSpecialNavigationStationPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AuthorityForSpecialNavigationStationPropertyType_stagedOrder, stage.AuthorityForSpecialNavigationStationPropertyTypes_referenceOrder, authorityforspecialnavigationstationpropertytype, "AuthorityForSpecialNavigationStationPropertyType")
}

func (authorityforspecialnavigationstationtype *AuthorityForSpecialNavigationStationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AuthorityForSpecialNavigationStationType_stagedOrder, stage.AuthorityForSpecialNavigationStationTypes_referenceOrder, authorityforspecialnavigationstationtype, "AuthorityForSpecialNavigationStationType")
}

func (authorityforspecialnavigationsystempropertytype *AuthorityForSpecialNavigationSystemPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AuthorityForSpecialNavigationSystemPropertyType_stagedOrder, stage.AuthorityForSpecialNavigationSystemPropertyTypes_referenceOrder, authorityforspecialnavigationsystempropertytype, "AuthorityForSpecialNavigationSystemPropertyType")
}

func (authorityforspecialnavigationsystemtype *AuthorityForSpecialNavigationSystemType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AuthorityForSpecialNavigationSystemType_stagedOrder, stage.AuthorityForSpecialNavigationSystemTypes_referenceOrder, authorityforspecialnavigationsystemtype, "AuthorityForSpecialNavigationSystemType")
}

func (azimuthtimeslicepropertytype *AzimuthTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AzimuthTimeSlicePropertyType_stagedOrder, stage.AzimuthTimeSlicePropertyTypes_referenceOrder, azimuthtimeslicepropertytype, "AzimuthTimeSlicePropertyType")
}

func (azimuthtimeslicetype *AzimuthTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AzimuthTimeSliceType_stagedOrder, stage.AzimuthTimeSliceTypes_referenceOrder, azimuthtimeslicetype, "AzimuthTimeSliceType")
}

func (azimuthtype *AzimuthType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AzimuthType_stagedOrder, stage.AzimuthTypes_referenceOrder, azimuthtype, "AzimuthType")
}

func (checkpointinstimeslicepropertytype *CheckpointINSTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CheckpointINSTimeSlicePropertyType_stagedOrder, stage.CheckpointINSTimeSlicePropertyTypes_referenceOrder, checkpointinstimeslicepropertytype, "CheckpointINSTimeSlicePropertyType")
}

func (checkpointinstimeslicetype *CheckpointINSTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CheckpointINSTimeSliceType_stagedOrder, stage.CheckpointINSTimeSliceTypes_referenceOrder, checkpointinstimeslicetype, "CheckpointINSTimeSliceType")
}

func (checkpointinstype *CheckpointINSType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CheckpointINSType_stagedOrder, stage.CheckpointINSTypes_referenceOrder, checkpointinstype, "CheckpointINSType")
}

func (checkpointvortimeslicepropertytype *CheckpointVORTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CheckpointVORTimeSlicePropertyType_stagedOrder, stage.CheckpointVORTimeSlicePropertyTypes_referenceOrder, checkpointvortimeslicepropertytype, "CheckpointVORTimeSlicePropertyType")
}

func (checkpointvortimeslicetype *CheckpointVORTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CheckpointVORTimeSliceType_stagedOrder, stage.CheckpointVORTimeSliceTypes_referenceOrder, checkpointvortimeslicetype, "CheckpointVORTimeSliceType")
}

func (checkpointvortype *CheckpointVORType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CheckpointVORType_stagedOrder, stage.CheckpointVORTypes_referenceOrder, checkpointvortype, "CheckpointVORType")
}

func (contactinformationpropertytype *ContactInformationPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ContactInformationPropertyType_stagedOrder, stage.ContactInformationPropertyTypes_referenceOrder, contactinformationpropertytype, "ContactInformationPropertyType")
}

func (contactinformationtype *ContactInformationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ContactInformationType_stagedOrder, stage.ContactInformationTypes_referenceOrder, contactinformationtype, "ContactInformationType")
}

func (dmetimeslicepropertytype *DMETimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DMETimeSlicePropertyType_stagedOrder, stage.DMETimeSlicePropertyTypes_referenceOrder, dmetimeslicepropertytype, "DMETimeSlicePropertyType")
}

func (dmetimeslicetype *DMETimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DMETimeSliceType_stagedOrder, stage.DMETimeSliceTypes_referenceOrder, dmetimeslicetype, "DMETimeSliceType")
}

func (dmetype *DMEType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DMEType_stagedOrder, stage.DMETypes_referenceOrder, dmetype, "DMEType")
}

func (designatedpointpropertytype *DesignatedPointPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DesignatedPointPropertyType_stagedOrder, stage.DesignatedPointPropertyTypes_referenceOrder, designatedpointpropertytype, "DesignatedPointPropertyType")
}

func (designatedpointtimeslicepropertytype *DesignatedPointTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DesignatedPointTimeSlicePropertyType_stagedOrder, stage.DesignatedPointTimeSlicePropertyTypes_referenceOrder, designatedpointtimeslicepropertytype, "DesignatedPointTimeSlicePropertyType")
}

func (designatedpointtimeslicetype *DesignatedPointTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DesignatedPointTimeSliceType_stagedOrder, stage.DesignatedPointTimeSliceTypes_referenceOrder, designatedpointtimeslicetype, "DesignatedPointTimeSliceType")
}

func (designatedpointtype *DesignatedPointType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DesignatedPointType_stagedOrder, stage.DesignatedPointTypes_referenceOrder, designatedpointtype, "DesignatedPointType")
}

func (directionfindertimeslicepropertytype *DirectionFinderTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DirectionFinderTimeSlicePropertyType_stagedOrder, stage.DirectionFinderTimeSlicePropertyTypes_referenceOrder, directionfindertimeslicepropertytype, "DirectionFinderTimeSlicePropertyType")
}

func (directionfindertimeslicetype *DirectionFinderTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DirectionFinderTimeSliceType_stagedOrder, stage.DirectionFinderTimeSliceTypes_referenceOrder, directionfindertimeslicetype, "DirectionFinderTimeSliceType")
}

func (directionfindertype *DirectionFinderType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DirectionFinderType_stagedOrder, stage.DirectionFinderTypes_referenceOrder, directionfindertype, "DirectionFinderType")
}

func (distanceindicationtimeslicepropertytype *DistanceIndicationTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DistanceIndicationTimeSlicePropertyType_stagedOrder, stage.DistanceIndicationTimeSlicePropertyTypes_referenceOrder, distanceindicationtimeslicepropertytype, "DistanceIndicationTimeSlicePropertyType")
}

func (distanceindicationtimeslicetype *DistanceIndicationTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DistanceIndicationTimeSliceType_stagedOrder, stage.DistanceIndicationTimeSliceTypes_referenceOrder, distanceindicationtimeslicetype, "DistanceIndicationTimeSliceType")
}

func (distanceindicationtype *DistanceIndicationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DistanceIndicationType_stagedOrder, stage.DistanceIndicationTypes_referenceOrder, distanceindicationtype, "DistanceIndicationType")
}

func (elevatedpointpropertytype *ElevatedPointPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ElevatedPointPropertyType_stagedOrder, stage.ElevatedPointPropertyTypes_referenceOrder, elevatedpointpropertytype, "ElevatedPointPropertyType")
}

func (elevatedpointtype *ElevatedPointType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ElevatedPointType_stagedOrder, stage.ElevatedPointTypes_referenceOrder, elevatedpointtype, "ElevatedPointType")
}

func (elevationtimeslicepropertytype *ElevationTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ElevationTimeSlicePropertyType_stagedOrder, stage.ElevationTimeSlicePropertyTypes_referenceOrder, elevationtimeslicepropertytype, "ElevationTimeSlicePropertyType")
}

func (elevationtimeslicetype *ElevationTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ElevationTimeSliceType_stagedOrder, stage.ElevationTimeSliceTypes_referenceOrder, elevationtimeslicetype, "ElevationTimeSliceType")
}

func (elevationtype *ElevationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ElevationType_stagedOrder, stage.ElevationTypes_referenceOrder, elevationtype, "ElevationType")
}

func (glidepathtimeslicepropertytype *GlidepathTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GlidepathTimeSlicePropertyType_stagedOrder, stage.GlidepathTimeSlicePropertyTypes_referenceOrder, glidepathtimeslicepropertytype, "GlidepathTimeSlicePropertyType")
}

func (glidepathtimeslicetype *GlidepathTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GlidepathTimeSliceType_stagedOrder, stage.GlidepathTimeSliceTypes_referenceOrder, glidepathtimeslicetype, "GlidepathTimeSliceType")
}

func (glidepathtype *GlidepathType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GlidepathType_stagedOrder, stage.GlidepathTypes_referenceOrder, glidepathtype, "GlidepathType")
}

func (informationservicepropertytype *InformationServicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.InformationServicePropertyType_stagedOrder, stage.InformationServicePropertyTypes_referenceOrder, informationservicepropertytype, "InformationServicePropertyType")
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LinguisticNotePropertyType_stagedOrder, stage.LinguisticNotePropertyTypes_referenceOrder, linguisticnotepropertytype, "LinguisticNotePropertyType")
}

func (linguisticnotetype *LinguisticNoteType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LinguisticNoteType_stagedOrder, stage.LinguisticNoteTypes_referenceOrder, linguisticnotetype, "LinguisticNoteType")
}

func (localizertimeslicepropertytype *LocalizerTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LocalizerTimeSlicePropertyType_stagedOrder, stage.LocalizerTimeSlicePropertyTypes_referenceOrder, localizertimeslicepropertytype, "LocalizerTimeSlicePropertyType")
}

func (localizertimeslicetype *LocalizerTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LocalizerTimeSliceType_stagedOrder, stage.LocalizerTimeSliceTypes_referenceOrder, localizertimeslicetype, "LocalizerTimeSliceType")
}

func (localizertype *LocalizerType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LocalizerType_stagedOrder, stage.LocalizerTypes_referenceOrder, localizertype, "LocalizerType")
}

func (markerbeacontimeslicepropertytype *MarkerBeaconTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MarkerBeaconTimeSlicePropertyType_stagedOrder, stage.MarkerBeaconTimeSlicePropertyTypes_referenceOrder, markerbeacontimeslicepropertytype, "MarkerBeaconTimeSlicePropertyType")
}

func (markerbeacontimeslicetype *MarkerBeaconTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MarkerBeaconTimeSliceType_stagedOrder, stage.MarkerBeaconTimeSliceTypes_referenceOrder, markerbeacontimeslicetype, "MarkerBeaconTimeSliceType")
}

func (markerbeacontype *MarkerBeaconType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MarkerBeaconType_stagedOrder, stage.MarkerBeaconTypes_referenceOrder, markerbeacontype, "MarkerBeaconType")
}

func (ndbtimeslicepropertytype *NDBTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NDBTimeSlicePropertyType_stagedOrder, stage.NDBTimeSlicePropertyTypes_referenceOrder, ndbtimeslicepropertytype, "NDBTimeSlicePropertyType")
}

func (ndbtimeslicetype *NDBTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NDBTimeSliceType_stagedOrder, stage.NDBTimeSliceTypes_referenceOrder, ndbtimeslicetype, "NDBTimeSliceType")
}

func (ndbtype *NDBType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NDBType_stagedOrder, stage.NDBTypes_referenceOrder, ndbtype, "NDBType")
}

func (navaidcomponentpropertytype *NavaidComponentPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NavaidComponentPropertyType_stagedOrder, stage.NavaidComponentPropertyTypes_referenceOrder, navaidcomponentpropertytype, "NavaidComponentPropertyType")
}

func (navaidcomponenttype *NavaidComponentType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NavaidComponentType_stagedOrder, stage.NavaidComponentTypes_referenceOrder, navaidcomponenttype, "NavaidComponentType")
}

func (navaidequipmentmonitoringpropertytype *NavaidEquipmentMonitoringPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NavaidEquipmentMonitoringPropertyType_stagedOrder, stage.NavaidEquipmentMonitoringPropertyTypes_referenceOrder, navaidequipmentmonitoringpropertytype, "NavaidEquipmentMonitoringPropertyType")
}

func (navaidequipmentmonitoringtype *NavaidEquipmentMonitoringType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NavaidEquipmentMonitoringType_stagedOrder, stage.NavaidEquipmentMonitoringTypes_referenceOrder, navaidequipmentmonitoringtype, "NavaidEquipmentMonitoringType")
}

func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NavaidEquipmentPropertyType_stagedOrder, stage.NavaidEquipmentPropertyTypes_referenceOrder, navaidequipmentpropertytype, "NavaidEquipmentPropertyType")
}

func (navaidoperationalstatuspropertytype *NavaidOperationalStatusPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NavaidOperationalStatusPropertyType_stagedOrder, stage.NavaidOperationalStatusPropertyTypes_referenceOrder, navaidoperationalstatuspropertytype, "NavaidOperationalStatusPropertyType")
}

func (navaidoperationalstatustype *NavaidOperationalStatusType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NavaidOperationalStatusType_stagedOrder, stage.NavaidOperationalStatusTypes_referenceOrder, navaidoperationalstatustype, "NavaidOperationalStatusType")
}

func (navaidpropertytype *NavaidPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NavaidPropertyType_stagedOrder, stage.NavaidPropertyTypes_referenceOrder, navaidpropertytype, "NavaidPropertyType")
}

func (navaidtimeslicepropertytype *NavaidTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NavaidTimeSlicePropertyType_stagedOrder, stage.NavaidTimeSlicePropertyTypes_referenceOrder, navaidtimeslicepropertytype, "NavaidTimeSlicePropertyType")
}

func (navaidtimeslicetype *NavaidTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NavaidTimeSliceType_stagedOrder, stage.NavaidTimeSliceTypes_referenceOrder, navaidtimeslicetype, "NavaidTimeSliceType")
}

func (navaidtype *NavaidType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NavaidType_stagedOrder, stage.NavaidTypes_referenceOrder, navaidtype, "NavaidType")
}

func (notepropertytype *NotePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NotePropertyType_stagedOrder, stage.NotePropertyTypes_referenceOrder, notepropertytype, "NotePropertyType")
}

func (notetype *NoteType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoteType_stagedOrder, stage.NoteTypes_referenceOrder, notetype, "NoteType")
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

func (pointtype *PointType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PointType_stagedOrder, stage.PointTypes_referenceOrder, pointtype, "PointType")
}

func (postaladdresspropertytype *PostalAddressPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PostalAddressPropertyType_stagedOrder, stage.PostalAddressPropertyTypes_referenceOrder, postaladdresspropertytype, "PostalAddressPropertyType")
}

func (postaladdresstype *PostalAddressType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PostalAddressType_stagedOrder, stage.PostalAddressTypes_referenceOrder, postaladdresstype, "PostalAddressType")
}

func (precisionapproachradartimeslicepropertytype *PrecisionApproachRadarTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PrecisionApproachRadarTimeSlicePropertyType_stagedOrder, stage.PrecisionApproachRadarTimeSlicePropertyTypes_referenceOrder, precisionapproachradartimeslicepropertytype, "PrecisionApproachRadarTimeSlicePropertyType")
}

func (precisionapproachradartimeslicetype *PrecisionApproachRadarTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PrecisionApproachRadarTimeSliceType_stagedOrder, stage.PrecisionApproachRadarTimeSliceTypes_referenceOrder, precisionapproachradartimeslicetype, "PrecisionApproachRadarTimeSliceType")
}

func (precisionapproachradartype *PrecisionApproachRadarType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PrecisionApproachRadarType_stagedOrder, stage.PrecisionApproachRadarTypes_referenceOrder, precisionapproachradartype, "PrecisionApproachRadarType")
}

func (primarysurveillanceradartimeslicepropertytype *PrimarySurveillanceRadarTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PrimarySurveillanceRadarTimeSlicePropertyType_stagedOrder, stage.PrimarySurveillanceRadarTimeSlicePropertyTypes_referenceOrder, primarysurveillanceradartimeslicepropertytype, "PrimarySurveillanceRadarTimeSlicePropertyType")
}

func (primarysurveillanceradartimeslicetype *PrimarySurveillanceRadarTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PrimarySurveillanceRadarTimeSliceType_stagedOrder, stage.PrimarySurveillanceRadarTimeSliceTypes_referenceOrder, primarysurveillanceradartimeslicetype, "PrimarySurveillanceRadarTimeSliceType")
}

func (primarysurveillanceradartype *PrimarySurveillanceRadarType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PrimarySurveillanceRadarType_stagedOrder, stage.PrimarySurveillanceRadarTypes_referenceOrder, primarysurveillanceradartype, "PrimarySurveillanceRadarType")
}

func (radarcomponentpropertytype *RadarComponentPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RadarComponentPropertyType_stagedOrder, stage.RadarComponentPropertyTypes_referenceOrder, radarcomponentpropertytype, "RadarComponentPropertyType")
}

func (radarcomponenttype *RadarComponentType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RadarComponentType_stagedOrder, stage.RadarComponentTypes_referenceOrder, radarcomponenttype, "RadarComponentType")
}

func (radarequipmentpropertytype *RadarEquipmentPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RadarEquipmentPropertyType_stagedOrder, stage.RadarEquipmentPropertyTypes_referenceOrder, radarequipmentpropertytype, "RadarEquipmentPropertyType")
}

func (radarsystemtimeslicepropertytype *RadarSystemTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RadarSystemTimeSlicePropertyType_stagedOrder, stage.RadarSystemTimeSlicePropertyTypes_referenceOrder, radarsystemtimeslicepropertytype, "RadarSystemTimeSlicePropertyType")
}

func (radarsystemtimeslicetype *RadarSystemTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RadarSystemTimeSliceType_stagedOrder, stage.RadarSystemTimeSliceTypes_referenceOrder, radarsystemtimeslicetype, "RadarSystemTimeSliceType")
}

func (radarsystemtype *RadarSystemType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RadarSystemType_stagedOrder, stage.RadarSystemTypes_referenceOrder, radarsystemtype, "RadarSystemType")
}

func (reflectorpropertytype *ReflectorPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ReflectorPropertyType_stagedOrder, stage.ReflectorPropertyTypes_referenceOrder, reflectorpropertytype, "ReflectorPropertyType")
}

func (reflectortype *ReflectorType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ReflectorType_stagedOrder, stage.ReflectorTypes_referenceOrder, reflectortype, "ReflectorType")
}

func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RunwayCentrelinePointPropertyType_stagedOrder, stage.RunwayCentrelinePointPropertyTypes_referenceOrder, runwaycentrelinepointpropertytype, "RunwayCentrelinePointPropertyType")
}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RunwayDirectionPropertyType_stagedOrder, stage.RunwayDirectionPropertyTypes_referenceOrder, runwaydirectionpropertytype, "RunwayDirectionPropertyType")
}

func (runwaypropertytype *RunwayPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RunwayPropertyType_stagedOrder, stage.RunwayPropertyTypes_referenceOrder, runwaypropertytype, "RunwayPropertyType")
}

func (sdftimeslicepropertytype *SDFTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SDFTimeSlicePropertyType_stagedOrder, stage.SDFTimeSlicePropertyTypes_referenceOrder, sdftimeslicepropertytype, "SDFTimeSlicePropertyType")
}

func (sdftimeslicetype *SDFTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SDFTimeSliceType_stagedOrder, stage.SDFTimeSliceTypes_referenceOrder, sdftimeslicetype, "SDFTimeSliceType")
}

func (sdftype *SDFType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SDFType_stagedOrder, stage.SDFTypes_referenceOrder, sdftype, "SDFType")
}

func (secondarysurveillanceradartimeslicepropertytype *SecondarySurveillanceRadarTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SecondarySurveillanceRadarTimeSlicePropertyType_stagedOrder, stage.SecondarySurveillanceRadarTimeSlicePropertyTypes_referenceOrder, secondarysurveillanceradartimeslicepropertytype, "SecondarySurveillanceRadarTimeSlicePropertyType")
}

func (secondarysurveillanceradartimeslicetype *SecondarySurveillanceRadarTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SecondarySurveillanceRadarTimeSliceType_stagedOrder, stage.SecondarySurveillanceRadarTimeSliceTypes_referenceOrder, secondarysurveillanceradartimeslicetype, "SecondarySurveillanceRadarTimeSliceType")
}

func (secondarysurveillanceradartype *SecondarySurveillanceRadarType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SecondarySurveillanceRadarType_stagedOrder, stage.SecondarySurveillanceRadarTypes_referenceOrder, secondarysurveillanceradartype, "SecondarySurveillanceRadarType")
}

func (significantpointinairspacetimeslicepropertytype *SignificantPointInAirspaceTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SignificantPointInAirspaceTimeSlicePropertyType_stagedOrder, stage.SignificantPointInAirspaceTimeSlicePropertyTypes_referenceOrder, significantpointinairspacetimeslicepropertytype, "SignificantPointInAirspaceTimeSlicePropertyType")
}

func (significantpointinairspacetimeslicetype *SignificantPointInAirspaceTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SignificantPointInAirspaceTimeSliceType_stagedOrder, stage.SignificantPointInAirspaceTimeSliceTypes_referenceOrder, significantpointinairspacetimeslicetype, "SignificantPointInAirspaceTimeSliceType")
}

func (significantpointinairspacetype *SignificantPointInAirspaceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SignificantPointInAirspaceType_stagedOrder, stage.SignificantPointInAirspaceTypes_referenceOrder, significantpointinairspacetype, "SignificantPointInAirspaceType")
}

func (specialnavigationstationstatuspropertytype *SpecialNavigationStationStatusPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SpecialNavigationStationStatusPropertyType_stagedOrder, stage.SpecialNavigationStationStatusPropertyTypes_referenceOrder, specialnavigationstationstatuspropertytype, "SpecialNavigationStationStatusPropertyType")
}

func (specialnavigationstationstatustype *SpecialNavigationStationStatusType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SpecialNavigationStationStatusType_stagedOrder, stage.SpecialNavigationStationStatusTypes_referenceOrder, specialnavigationstationstatustype, "SpecialNavigationStationStatusType")
}

func (specialnavigationstationtimeslicepropertytype *SpecialNavigationStationTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SpecialNavigationStationTimeSlicePropertyType_stagedOrder, stage.SpecialNavigationStationTimeSlicePropertyTypes_referenceOrder, specialnavigationstationtimeslicepropertytype, "SpecialNavigationStationTimeSlicePropertyType")
}

func (specialnavigationstationtimeslicetype *SpecialNavigationStationTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SpecialNavigationStationTimeSliceType_stagedOrder, stage.SpecialNavigationStationTimeSliceTypes_referenceOrder, specialnavigationstationtimeslicetype, "SpecialNavigationStationTimeSliceType")
}

func (specialnavigationstationtype *SpecialNavigationStationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SpecialNavigationStationType_stagedOrder, stage.SpecialNavigationStationTypes_referenceOrder, specialnavigationstationtype, "SpecialNavigationStationType")
}

func (specialnavigationsystempropertytype *SpecialNavigationSystemPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SpecialNavigationSystemPropertyType_stagedOrder, stage.SpecialNavigationSystemPropertyTypes_referenceOrder, specialnavigationsystempropertytype, "SpecialNavigationSystemPropertyType")
}

func (specialnavigationsystemtimeslicepropertytype *SpecialNavigationSystemTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SpecialNavigationSystemTimeSlicePropertyType_stagedOrder, stage.SpecialNavigationSystemTimeSlicePropertyTypes_referenceOrder, specialnavigationsystemtimeslicepropertytype, "SpecialNavigationSystemTimeSlicePropertyType")
}

func (specialnavigationsystemtimeslicetype *SpecialNavigationSystemTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SpecialNavigationSystemTimeSliceType_stagedOrder, stage.SpecialNavigationSystemTimeSliceTypes_referenceOrder, specialnavigationsystemtimeslicetype, "SpecialNavigationSystemTimeSliceType")
}

func (specialnavigationsystemtype *SpecialNavigationSystemType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SpecialNavigationSystemType_stagedOrder, stage.SpecialNavigationSystemTypes_referenceOrder, specialnavigationsystemtype, "SpecialNavigationSystemType")
}

func (surveillancegroundstationpropertytype *SurveillanceGroundStationPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SurveillanceGroundStationPropertyType_stagedOrder, stage.SurveillanceGroundStationPropertyTypes_referenceOrder, surveillancegroundstationpropertytype, "SurveillanceGroundStationPropertyType")
}

func (surveillancegroundstationtype *SurveillanceGroundStationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SurveillanceGroundStationType_stagedOrder, stage.SurveillanceGroundStationTypes_referenceOrder, surveillancegroundstationtype, "SurveillanceGroundStationType")
}

func (tacantimeslicepropertytype *TACANTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TACANTimeSlicePropertyType_stagedOrder, stage.TACANTimeSlicePropertyTypes_referenceOrder, tacantimeslicepropertytype, "TACANTimeSlicePropertyType")
}

func (tacantimeslicetype *TACANTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TACANTimeSliceType_stagedOrder, stage.TACANTimeSliceTypes_referenceOrder, tacantimeslicetype, "TACANTimeSliceType")
}

func (tacantype *TACANType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TACANType_stagedOrder, stage.TACANTypes_referenceOrder, tacantype, "TACANType")
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

func (unitpropertytype *UnitPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.UnitPropertyType_stagedOrder, stage.UnitPropertyTypes_referenceOrder, unitpropertytype, "UnitPropertyType")
}

func (vorpropertytype *VORPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.VORPropertyType_stagedOrder, stage.VORPropertyTypes_referenceOrder, vorpropertytype, "VORPropertyType")
}

func (vortimeslicepropertytype *VORTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.VORTimeSlicePropertyType_stagedOrder, stage.VORTimeSlicePropertyTypes_referenceOrder, vortimeslicepropertytype, "VORTimeSlicePropertyType")
}

func (vortimeslicetype *VORTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.VORTimeSliceType_stagedOrder, stage.VORTimeSliceTypes_referenceOrder, vortimeslicetype, "VORTimeSliceType")
}

func (vortype *VORType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.VORType_stagedOrder, stage.VORTypes_referenceOrder, vortype, "VORType")
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

func (abstractnavaidequipmenttype *AbstractNavaidEquipmentType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractnavaidequipmenttype, abstractnavaidequipmenttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractnavaidequipmenttype *AbstractNavaidEquipmentType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractnavaidequipmenttype.GongGetIdentifier(stage)
}

func (abstractnavigationsystemcheckpointtype *AbstractNavigationSystemCheckpointType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractnavigationsystemcheckpointtype, abstractnavigationsystemcheckpointtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractnavigationsystemcheckpointtype *AbstractNavigationSystemCheckpointType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractnavigationsystemcheckpointtype.GongGetIdentifier(stage)
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractpropertieswithscheduletype, abstractpropertieswithscheduletype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractpropertieswithscheduletype.GongGetIdentifier(stage)
}

func (abstractradarequipmenttype *AbstractRadarEquipmentType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractradarequipmenttype, abstractradarequipmenttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractradarequipmenttype *AbstractRadarEquipmentType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractradarequipmenttype.GongGetIdentifier(stage)
}

func (abstractsurveillanceradartype *AbstractSurveillanceRadarType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractsurveillanceradartype, abstractsurveillanceradartype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractsurveillanceradartype *AbstractSurveillanceRadarType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractsurveillanceradartype.GongGetIdentifier(stage)
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airportheliportpropertytype, airportheliportpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airportheliportpropertytype *AirportHeliportPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return airportheliportpropertytype.GongGetIdentifier(stage)
}

func (airspacepropertytype *AirspacePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airspacepropertytype, airspacepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airspacepropertytype *AirspacePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return airspacepropertytype.GongGetIdentifier(stage)
}

func (angleindicationtimeslicepropertytype *AngleIndicationTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(angleindicationtimeslicepropertytype, angleindicationtimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (angleindicationtimeslicepropertytype *AngleIndicationTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return angleindicationtimeslicepropertytype.GongGetIdentifier(stage)
}

func (angleindicationtimeslicetype *AngleIndicationTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(angleindicationtimeslicetype, angleindicationtimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (angleindicationtimeslicetype *AngleIndicationTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return angleindicationtimeslicetype.GongGetIdentifier(stage)
}

func (angleindicationtype *AngleIndicationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(angleindicationtype, angleindicationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (angleindicationtype *AngleIndicationType) GongGetReferenceIdentifier(stage *Stage) string {
	return angleindicationtype.GongGetIdentifier(stage)
}

func (authorityfornavaidequipmentpropertytype *AuthorityForNavaidEquipmentPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(authorityfornavaidequipmentpropertytype, authorityfornavaidequipmentpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (authorityfornavaidequipmentpropertytype *AuthorityForNavaidEquipmentPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return authorityfornavaidequipmentpropertytype.GongGetIdentifier(stage)
}

func (authorityfornavaidequipmenttype *AuthorityForNavaidEquipmentType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(authorityfornavaidequipmenttype, authorityfornavaidequipmenttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (authorityfornavaidequipmenttype *AuthorityForNavaidEquipmentType) GongGetReferenceIdentifier(stage *Stage) string {
	return authorityfornavaidequipmenttype.GongGetIdentifier(stage)
}

func (authorityforspecialnavigationstationpropertytype *AuthorityForSpecialNavigationStationPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(authorityforspecialnavigationstationpropertytype, authorityforspecialnavigationstationpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (authorityforspecialnavigationstationpropertytype *AuthorityForSpecialNavigationStationPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return authorityforspecialnavigationstationpropertytype.GongGetIdentifier(stage)
}

func (authorityforspecialnavigationstationtype *AuthorityForSpecialNavigationStationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(authorityforspecialnavigationstationtype, authorityforspecialnavigationstationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (authorityforspecialnavigationstationtype *AuthorityForSpecialNavigationStationType) GongGetReferenceIdentifier(stage *Stage) string {
	return authorityforspecialnavigationstationtype.GongGetIdentifier(stage)
}

func (authorityforspecialnavigationsystempropertytype *AuthorityForSpecialNavigationSystemPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(authorityforspecialnavigationsystempropertytype, authorityforspecialnavigationsystempropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (authorityforspecialnavigationsystempropertytype *AuthorityForSpecialNavigationSystemPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return authorityforspecialnavigationsystempropertytype.GongGetIdentifier(stage)
}

func (authorityforspecialnavigationsystemtype *AuthorityForSpecialNavigationSystemType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(authorityforspecialnavigationsystemtype, authorityforspecialnavigationsystemtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (authorityforspecialnavigationsystemtype *AuthorityForSpecialNavigationSystemType) GongGetReferenceIdentifier(stage *Stage) string {
	return authorityforspecialnavigationsystemtype.GongGetIdentifier(stage)
}

func (azimuthtimeslicepropertytype *AzimuthTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(azimuthtimeslicepropertytype, azimuthtimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (azimuthtimeslicepropertytype *AzimuthTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return azimuthtimeslicepropertytype.GongGetIdentifier(stage)
}

func (azimuthtimeslicetype *AzimuthTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(azimuthtimeslicetype, azimuthtimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (azimuthtimeslicetype *AzimuthTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return azimuthtimeslicetype.GongGetIdentifier(stage)
}

func (azimuthtype *AzimuthType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(azimuthtype, azimuthtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (azimuthtype *AzimuthType) GongGetReferenceIdentifier(stage *Stage) string {
	return azimuthtype.GongGetIdentifier(stage)
}

func (checkpointinstimeslicepropertytype *CheckpointINSTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(checkpointinstimeslicepropertytype, checkpointinstimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (checkpointinstimeslicepropertytype *CheckpointINSTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return checkpointinstimeslicepropertytype.GongGetIdentifier(stage)
}

func (checkpointinstimeslicetype *CheckpointINSTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(checkpointinstimeslicetype, checkpointinstimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (checkpointinstimeslicetype *CheckpointINSTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return checkpointinstimeslicetype.GongGetIdentifier(stage)
}

func (checkpointinstype *CheckpointINSType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(checkpointinstype, checkpointinstype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (checkpointinstype *CheckpointINSType) GongGetReferenceIdentifier(stage *Stage) string {
	return checkpointinstype.GongGetIdentifier(stage)
}

func (checkpointvortimeslicepropertytype *CheckpointVORTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(checkpointvortimeslicepropertytype, checkpointvortimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (checkpointvortimeslicepropertytype *CheckpointVORTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return checkpointvortimeslicepropertytype.GongGetIdentifier(stage)
}

func (checkpointvortimeslicetype *CheckpointVORTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(checkpointvortimeslicetype, checkpointvortimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (checkpointvortimeslicetype *CheckpointVORTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return checkpointvortimeslicetype.GongGetIdentifier(stage)
}

func (checkpointvortype *CheckpointVORType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(checkpointvortype, checkpointvortype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (checkpointvortype *CheckpointVORType) GongGetReferenceIdentifier(stage *Stage) string {
	return checkpointvortype.GongGetIdentifier(stage)
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

func (dmetimeslicepropertytype *DMETimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(dmetimeslicepropertytype, dmetimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (dmetimeslicepropertytype *DMETimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return dmetimeslicepropertytype.GongGetIdentifier(stage)
}

func (dmetimeslicetype *DMETimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(dmetimeslicetype, dmetimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (dmetimeslicetype *DMETimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return dmetimeslicetype.GongGetIdentifier(stage)
}

func (dmetype *DMEType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(dmetype, dmetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (dmetype *DMEType) GongGetReferenceIdentifier(stage *Stage) string {
	return dmetype.GongGetIdentifier(stage)
}

func (designatedpointpropertytype *DesignatedPointPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(designatedpointpropertytype, designatedpointpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (designatedpointpropertytype *DesignatedPointPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return designatedpointpropertytype.GongGetIdentifier(stage)
}

func (designatedpointtimeslicepropertytype *DesignatedPointTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(designatedpointtimeslicepropertytype, designatedpointtimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (designatedpointtimeslicepropertytype *DesignatedPointTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return designatedpointtimeslicepropertytype.GongGetIdentifier(stage)
}

func (designatedpointtimeslicetype *DesignatedPointTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(designatedpointtimeslicetype, designatedpointtimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (designatedpointtimeslicetype *DesignatedPointTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return designatedpointtimeslicetype.GongGetIdentifier(stage)
}

func (designatedpointtype *DesignatedPointType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(designatedpointtype, designatedpointtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (designatedpointtype *DesignatedPointType) GongGetReferenceIdentifier(stage *Stage) string {
	return designatedpointtype.GongGetIdentifier(stage)
}

func (directionfindertimeslicepropertytype *DirectionFinderTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(directionfindertimeslicepropertytype, directionfindertimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (directionfindertimeslicepropertytype *DirectionFinderTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return directionfindertimeslicepropertytype.GongGetIdentifier(stage)
}

func (directionfindertimeslicetype *DirectionFinderTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(directionfindertimeslicetype, directionfindertimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (directionfindertimeslicetype *DirectionFinderTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return directionfindertimeslicetype.GongGetIdentifier(stage)
}

func (directionfindertype *DirectionFinderType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(directionfindertype, directionfindertype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (directionfindertype *DirectionFinderType) GongGetReferenceIdentifier(stage *Stage) string {
	return directionfindertype.GongGetIdentifier(stage)
}

func (distanceindicationtimeslicepropertytype *DistanceIndicationTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(distanceindicationtimeslicepropertytype, distanceindicationtimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (distanceindicationtimeslicepropertytype *DistanceIndicationTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return distanceindicationtimeslicepropertytype.GongGetIdentifier(stage)
}

func (distanceindicationtimeslicetype *DistanceIndicationTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(distanceindicationtimeslicetype, distanceindicationtimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (distanceindicationtimeslicetype *DistanceIndicationTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return distanceindicationtimeslicetype.GongGetIdentifier(stage)
}

func (distanceindicationtype *DistanceIndicationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(distanceindicationtype, distanceindicationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (distanceindicationtype *DistanceIndicationType) GongGetReferenceIdentifier(stage *Stage) string {
	return distanceindicationtype.GongGetIdentifier(stage)
}

func (elevatedpointpropertytype *ElevatedPointPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(elevatedpointpropertytype, elevatedpointpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (elevatedpointpropertytype *ElevatedPointPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return elevatedpointpropertytype.GongGetIdentifier(stage)
}

func (elevatedpointtype *ElevatedPointType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(elevatedpointtype, elevatedpointtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (elevatedpointtype *ElevatedPointType) GongGetReferenceIdentifier(stage *Stage) string {
	return elevatedpointtype.GongGetIdentifier(stage)
}

func (elevationtimeslicepropertytype *ElevationTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(elevationtimeslicepropertytype, elevationtimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (elevationtimeslicepropertytype *ElevationTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return elevationtimeslicepropertytype.GongGetIdentifier(stage)
}

func (elevationtimeslicetype *ElevationTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(elevationtimeslicetype, elevationtimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (elevationtimeslicetype *ElevationTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return elevationtimeslicetype.GongGetIdentifier(stage)
}

func (elevationtype *ElevationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(elevationtype, elevationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (elevationtype *ElevationType) GongGetReferenceIdentifier(stage *Stage) string {
	return elevationtype.GongGetIdentifier(stage)
}

func (glidepathtimeslicepropertytype *GlidepathTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(glidepathtimeslicepropertytype, glidepathtimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (glidepathtimeslicepropertytype *GlidepathTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return glidepathtimeslicepropertytype.GongGetIdentifier(stage)
}

func (glidepathtimeslicetype *GlidepathTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(glidepathtimeslicetype, glidepathtimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (glidepathtimeslicetype *GlidepathTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return glidepathtimeslicetype.GongGetIdentifier(stage)
}

func (glidepathtype *GlidepathType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(glidepathtype, glidepathtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (glidepathtype *GlidepathType) GongGetReferenceIdentifier(stage *Stage) string {
	return glidepathtype.GongGetIdentifier(stage)
}

func (informationservicepropertytype *InformationServicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(informationservicepropertytype, informationservicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (informationservicepropertytype *InformationServicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return informationservicepropertytype.GongGetIdentifier(stage)
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

func (localizertimeslicepropertytype *LocalizerTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(localizertimeslicepropertytype, localizertimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (localizertimeslicepropertytype *LocalizerTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return localizertimeslicepropertytype.GongGetIdentifier(stage)
}

func (localizertimeslicetype *LocalizerTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(localizertimeslicetype, localizertimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (localizertimeslicetype *LocalizerTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return localizertimeslicetype.GongGetIdentifier(stage)
}

func (localizertype *LocalizerType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(localizertype, localizertype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (localizertype *LocalizerType) GongGetReferenceIdentifier(stage *Stage) string {
	return localizertype.GongGetIdentifier(stage)
}

func (markerbeacontimeslicepropertytype *MarkerBeaconTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(markerbeacontimeslicepropertytype, markerbeacontimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (markerbeacontimeslicepropertytype *MarkerBeaconTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return markerbeacontimeslicepropertytype.GongGetIdentifier(stage)
}

func (markerbeacontimeslicetype *MarkerBeaconTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(markerbeacontimeslicetype, markerbeacontimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (markerbeacontimeslicetype *MarkerBeaconTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return markerbeacontimeslicetype.GongGetIdentifier(stage)
}

func (markerbeacontype *MarkerBeaconType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(markerbeacontype, markerbeacontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (markerbeacontype *MarkerBeaconType) GongGetReferenceIdentifier(stage *Stage) string {
	return markerbeacontype.GongGetIdentifier(stage)
}

func (ndbtimeslicepropertytype *NDBTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(ndbtimeslicepropertytype, ndbtimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (ndbtimeslicepropertytype *NDBTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return ndbtimeslicepropertytype.GongGetIdentifier(stage)
}

func (ndbtimeslicetype *NDBTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(ndbtimeslicetype, ndbtimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (ndbtimeslicetype *NDBTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return ndbtimeslicetype.GongGetIdentifier(stage)
}

func (ndbtype *NDBType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(ndbtype, ndbtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (ndbtype *NDBType) GongGetReferenceIdentifier(stage *Stage) string {
	return ndbtype.GongGetIdentifier(stage)
}

func (navaidcomponentpropertytype *NavaidComponentPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(navaidcomponentpropertytype, navaidcomponentpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (navaidcomponentpropertytype *NavaidComponentPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return navaidcomponentpropertytype.GongGetIdentifier(stage)
}

func (navaidcomponenttype *NavaidComponentType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(navaidcomponenttype, navaidcomponenttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (navaidcomponenttype *NavaidComponentType) GongGetReferenceIdentifier(stage *Stage) string {
	return navaidcomponenttype.GongGetIdentifier(stage)
}

func (navaidequipmentmonitoringpropertytype *NavaidEquipmentMonitoringPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(navaidequipmentmonitoringpropertytype, navaidequipmentmonitoringpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (navaidequipmentmonitoringpropertytype *NavaidEquipmentMonitoringPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return navaidequipmentmonitoringpropertytype.GongGetIdentifier(stage)
}

func (navaidequipmentmonitoringtype *NavaidEquipmentMonitoringType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(navaidequipmentmonitoringtype, navaidequipmentmonitoringtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (navaidequipmentmonitoringtype *NavaidEquipmentMonitoringType) GongGetReferenceIdentifier(stage *Stage) string {
	return navaidequipmentmonitoringtype.GongGetIdentifier(stage)
}

func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(navaidequipmentpropertytype, navaidequipmentpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return navaidequipmentpropertytype.GongGetIdentifier(stage)
}

func (navaidoperationalstatuspropertytype *NavaidOperationalStatusPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(navaidoperationalstatuspropertytype, navaidoperationalstatuspropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (navaidoperationalstatuspropertytype *NavaidOperationalStatusPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return navaidoperationalstatuspropertytype.GongGetIdentifier(stage)
}

func (navaidoperationalstatustype *NavaidOperationalStatusType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(navaidoperationalstatustype, navaidoperationalstatustype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (navaidoperationalstatustype *NavaidOperationalStatusType) GongGetReferenceIdentifier(stage *Stage) string {
	return navaidoperationalstatustype.GongGetIdentifier(stage)
}

func (navaidpropertytype *NavaidPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(navaidpropertytype, navaidpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (navaidpropertytype *NavaidPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return navaidpropertytype.GongGetIdentifier(stage)
}

func (navaidtimeslicepropertytype *NavaidTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(navaidtimeslicepropertytype, navaidtimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (navaidtimeslicepropertytype *NavaidTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return navaidtimeslicepropertytype.GongGetIdentifier(stage)
}

func (navaidtimeslicetype *NavaidTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(navaidtimeslicetype, navaidtimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (navaidtimeslicetype *NavaidTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return navaidtimeslicetype.GongGetIdentifier(stage)
}

func (navaidtype *NavaidType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(navaidtype, navaidtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (navaidtype *NavaidType) GongGetReferenceIdentifier(stage *Stage) string {
	return navaidtype.GongGetIdentifier(stage)
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

func (precisionapproachradartimeslicepropertytype *PrecisionApproachRadarTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(precisionapproachradartimeslicepropertytype, precisionapproachradartimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (precisionapproachradartimeslicepropertytype *PrecisionApproachRadarTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return precisionapproachradartimeslicepropertytype.GongGetIdentifier(stage)
}

func (precisionapproachradartimeslicetype *PrecisionApproachRadarTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(precisionapproachradartimeslicetype, precisionapproachradartimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (precisionapproachradartimeslicetype *PrecisionApproachRadarTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return precisionapproachradartimeslicetype.GongGetIdentifier(stage)
}

func (precisionapproachradartype *PrecisionApproachRadarType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(precisionapproachradartype, precisionapproachradartype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (precisionapproachradartype *PrecisionApproachRadarType) GongGetReferenceIdentifier(stage *Stage) string {
	return precisionapproachradartype.GongGetIdentifier(stage)
}

func (primarysurveillanceradartimeslicepropertytype *PrimarySurveillanceRadarTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(primarysurveillanceradartimeslicepropertytype, primarysurveillanceradartimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (primarysurveillanceradartimeslicepropertytype *PrimarySurveillanceRadarTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return primarysurveillanceradartimeslicepropertytype.GongGetIdentifier(stage)
}

func (primarysurveillanceradartimeslicetype *PrimarySurveillanceRadarTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(primarysurveillanceradartimeslicetype, primarysurveillanceradartimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (primarysurveillanceradartimeslicetype *PrimarySurveillanceRadarTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return primarysurveillanceradartimeslicetype.GongGetIdentifier(stage)
}

func (primarysurveillanceradartype *PrimarySurveillanceRadarType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(primarysurveillanceradartype, primarysurveillanceradartype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (primarysurveillanceradartype *PrimarySurveillanceRadarType) GongGetReferenceIdentifier(stage *Stage) string {
	return primarysurveillanceradartype.GongGetIdentifier(stage)
}

func (radarcomponentpropertytype *RadarComponentPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(radarcomponentpropertytype, radarcomponentpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (radarcomponentpropertytype *RadarComponentPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return radarcomponentpropertytype.GongGetIdentifier(stage)
}

func (radarcomponenttype *RadarComponentType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(radarcomponenttype, radarcomponenttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (radarcomponenttype *RadarComponentType) GongGetReferenceIdentifier(stage *Stage) string {
	return radarcomponenttype.GongGetIdentifier(stage)
}

func (radarequipmentpropertytype *RadarEquipmentPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(radarequipmentpropertytype, radarequipmentpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (radarequipmentpropertytype *RadarEquipmentPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return radarequipmentpropertytype.GongGetIdentifier(stage)
}

func (radarsystemtimeslicepropertytype *RadarSystemTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(radarsystemtimeslicepropertytype, radarsystemtimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (radarsystemtimeslicepropertytype *RadarSystemTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return radarsystemtimeslicepropertytype.GongGetIdentifier(stage)
}

func (radarsystemtimeslicetype *RadarSystemTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(radarsystemtimeslicetype, radarsystemtimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (radarsystemtimeslicetype *RadarSystemTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return radarsystemtimeslicetype.GongGetIdentifier(stage)
}

func (radarsystemtype *RadarSystemType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(radarsystemtype, radarsystemtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (radarsystemtype *RadarSystemType) GongGetReferenceIdentifier(stage *Stage) string {
	return radarsystemtype.GongGetIdentifier(stage)
}

func (reflectorpropertytype *ReflectorPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(reflectorpropertytype, reflectorpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (reflectorpropertytype *ReflectorPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return reflectorpropertytype.GongGetIdentifier(stage)
}

func (reflectortype *ReflectorType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(reflectortype, reflectortype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (reflectortype *ReflectorType) GongGetReferenceIdentifier(stage *Stage) string {
	return reflectortype.GongGetIdentifier(stage)
}

func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(runwaycentrelinepointpropertytype, runwaycentrelinepointpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return runwaycentrelinepointpropertytype.GongGetIdentifier(stage)
}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(runwaydirectionpropertytype, runwaydirectionpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return runwaydirectionpropertytype.GongGetIdentifier(stage)
}

func (runwaypropertytype *RunwayPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(runwaypropertytype, runwaypropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (runwaypropertytype *RunwayPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return runwaypropertytype.GongGetIdentifier(stage)
}

func (sdftimeslicepropertytype *SDFTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(sdftimeslicepropertytype, sdftimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (sdftimeslicepropertytype *SDFTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return sdftimeslicepropertytype.GongGetIdentifier(stage)
}

func (sdftimeslicetype *SDFTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(sdftimeslicetype, sdftimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (sdftimeslicetype *SDFTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return sdftimeslicetype.GongGetIdentifier(stage)
}

func (sdftype *SDFType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(sdftype, sdftype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (sdftype *SDFType) GongGetReferenceIdentifier(stage *Stage) string {
	return sdftype.GongGetIdentifier(stage)
}

func (secondarysurveillanceradartimeslicepropertytype *SecondarySurveillanceRadarTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(secondarysurveillanceradartimeslicepropertytype, secondarysurveillanceradartimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (secondarysurveillanceradartimeslicepropertytype *SecondarySurveillanceRadarTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return secondarysurveillanceradartimeslicepropertytype.GongGetIdentifier(stage)
}

func (secondarysurveillanceradartimeslicetype *SecondarySurveillanceRadarTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(secondarysurveillanceradartimeslicetype, secondarysurveillanceradartimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (secondarysurveillanceradartimeslicetype *SecondarySurveillanceRadarTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return secondarysurveillanceradartimeslicetype.GongGetIdentifier(stage)
}

func (secondarysurveillanceradartype *SecondarySurveillanceRadarType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(secondarysurveillanceradartype, secondarysurveillanceradartype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (secondarysurveillanceradartype *SecondarySurveillanceRadarType) GongGetReferenceIdentifier(stage *Stage) string {
	return secondarysurveillanceradartype.GongGetIdentifier(stage)
}

func (significantpointinairspacetimeslicepropertytype *SignificantPointInAirspaceTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(significantpointinairspacetimeslicepropertytype, significantpointinairspacetimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (significantpointinairspacetimeslicepropertytype *SignificantPointInAirspaceTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return significantpointinairspacetimeslicepropertytype.GongGetIdentifier(stage)
}

func (significantpointinairspacetimeslicetype *SignificantPointInAirspaceTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(significantpointinairspacetimeslicetype, significantpointinairspacetimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (significantpointinairspacetimeslicetype *SignificantPointInAirspaceTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return significantpointinairspacetimeslicetype.GongGetIdentifier(stage)
}

func (significantpointinairspacetype *SignificantPointInAirspaceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(significantpointinairspacetype, significantpointinairspacetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (significantpointinairspacetype *SignificantPointInAirspaceType) GongGetReferenceIdentifier(stage *Stage) string {
	return significantpointinairspacetype.GongGetIdentifier(stage)
}

func (specialnavigationstationstatuspropertytype *SpecialNavigationStationStatusPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(specialnavigationstationstatuspropertytype, specialnavigationstationstatuspropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (specialnavigationstationstatuspropertytype *SpecialNavigationStationStatusPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return specialnavigationstationstatuspropertytype.GongGetIdentifier(stage)
}

func (specialnavigationstationstatustype *SpecialNavigationStationStatusType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(specialnavigationstationstatustype, specialnavigationstationstatustype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (specialnavigationstationstatustype *SpecialNavigationStationStatusType) GongGetReferenceIdentifier(stage *Stage) string {
	return specialnavigationstationstatustype.GongGetIdentifier(stage)
}

func (specialnavigationstationtimeslicepropertytype *SpecialNavigationStationTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(specialnavigationstationtimeslicepropertytype, specialnavigationstationtimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (specialnavigationstationtimeslicepropertytype *SpecialNavigationStationTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return specialnavigationstationtimeslicepropertytype.GongGetIdentifier(stage)
}

func (specialnavigationstationtimeslicetype *SpecialNavigationStationTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(specialnavigationstationtimeslicetype, specialnavigationstationtimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (specialnavigationstationtimeslicetype *SpecialNavigationStationTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return specialnavigationstationtimeslicetype.GongGetIdentifier(stage)
}

func (specialnavigationstationtype *SpecialNavigationStationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(specialnavigationstationtype, specialnavigationstationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (specialnavigationstationtype *SpecialNavigationStationType) GongGetReferenceIdentifier(stage *Stage) string {
	return specialnavigationstationtype.GongGetIdentifier(stage)
}

func (specialnavigationsystempropertytype *SpecialNavigationSystemPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(specialnavigationsystempropertytype, specialnavigationsystempropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (specialnavigationsystempropertytype *SpecialNavigationSystemPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return specialnavigationsystempropertytype.GongGetIdentifier(stage)
}

func (specialnavigationsystemtimeslicepropertytype *SpecialNavigationSystemTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(specialnavigationsystemtimeslicepropertytype, specialnavigationsystemtimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (specialnavigationsystemtimeslicepropertytype *SpecialNavigationSystemTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return specialnavigationsystemtimeslicepropertytype.GongGetIdentifier(stage)
}

func (specialnavigationsystemtimeslicetype *SpecialNavigationSystemTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(specialnavigationsystemtimeslicetype, specialnavigationsystemtimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (specialnavigationsystemtimeslicetype *SpecialNavigationSystemTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return specialnavigationsystemtimeslicetype.GongGetIdentifier(stage)
}

func (specialnavigationsystemtype *SpecialNavigationSystemType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(specialnavigationsystemtype, specialnavigationsystemtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (specialnavigationsystemtype *SpecialNavigationSystemType) GongGetReferenceIdentifier(stage *Stage) string {
	return specialnavigationsystemtype.GongGetIdentifier(stage)
}

func (surveillancegroundstationpropertytype *SurveillanceGroundStationPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(surveillancegroundstationpropertytype, surveillancegroundstationpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (surveillancegroundstationpropertytype *SurveillanceGroundStationPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return surveillancegroundstationpropertytype.GongGetIdentifier(stage)
}

func (surveillancegroundstationtype *SurveillanceGroundStationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(surveillancegroundstationtype, surveillancegroundstationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (surveillancegroundstationtype *SurveillanceGroundStationType) GongGetReferenceIdentifier(stage *Stage) string {
	return surveillancegroundstationtype.GongGetIdentifier(stage)
}

func (tacantimeslicepropertytype *TACANTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tacantimeslicepropertytype, tacantimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tacantimeslicepropertytype *TACANTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return tacantimeslicepropertytype.GongGetIdentifier(stage)
}

func (tacantimeslicetype *TACANTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tacantimeslicetype, tacantimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tacantimeslicetype *TACANTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return tacantimeslicetype.GongGetIdentifier(stage)
}

func (tacantype *TACANType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tacantype, tacantype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tacantype *TACANType) GongGetReferenceIdentifier(stage *Stage) string {
	return tacantype.GongGetIdentifier(stage)
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

func (unitpropertytype *UnitPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(unitpropertytype, unitpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (unitpropertytype *UnitPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return unitpropertytype.GongGetIdentifier(stage)
}

func (vorpropertytype *VORPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(vorpropertytype, vorpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (vorpropertytype *VORPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return vorpropertytype.GongGetIdentifier(stage)
}

func (vortimeslicepropertytype *VORTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(vortimeslicepropertytype, vortimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (vortimeslicepropertytype *VORTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return vortimeslicepropertytype.GongGetIdentifier(stage)
}

func (vortimeslicetype *VORTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(vortimeslicetype, vortimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (vortimeslicetype *VORTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return vortimeslicetype.GongGetIdentifier(stage)
}

func (vortype *VORType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(vortype, vortype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (vortype *VORType) GongGetReferenceIdentifier(stage *Stage) string {
	return vortype.GongGetIdentifier(stage)
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

func (abstractnavaidequipmenttype *AbstractNavaidEquipmentType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractnavaidequipmenttype.GongGetIdentifier(stage), "AbstractNavaidEquipmentType", abstractnavaidequipmenttype.Name)
}

func (abstractnavigationsystemcheckpointtype *AbstractNavigationSystemCheckpointType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractnavigationsystemcheckpointtype.GongGetIdentifier(stage), "AbstractNavigationSystemCheckpointType", abstractnavigationsystemcheckpointtype.Name)
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractpropertieswithscheduletype.GongGetIdentifier(stage), "AbstractPropertiesWithScheduleType", abstractpropertieswithscheduletype.Name)
}

func (abstractradarequipmenttype *AbstractRadarEquipmentType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractradarequipmenttype.GongGetIdentifier(stage), "AbstractRadarEquipmentType", abstractradarequipmenttype.Name)
}

func (abstractsurveillanceradartype *AbstractSurveillanceRadarType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractsurveillanceradartype.GongGetIdentifier(stage), "AbstractSurveillanceRadarType", abstractsurveillanceradartype.Name)
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airportheliportpropertytype.GongGetIdentifier(stage), "AirportHeliportPropertyType", airportheliportpropertytype.Name)
}

func (airspacepropertytype *AirspacePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airspacepropertytype.GongGetIdentifier(stage), "AirspacePropertyType", airspacepropertytype.Name)
}

func (angleindicationtimeslicepropertytype *AngleIndicationTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(angleindicationtimeslicepropertytype.GongGetIdentifier(stage), "AngleIndicationTimeSlicePropertyType", angleindicationtimeslicepropertytype.Name)
}

func (angleindicationtimeslicetype *AngleIndicationTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(angleindicationtimeslicetype.GongGetIdentifier(stage), "AngleIndicationTimeSliceType", angleindicationtimeslicetype.Name)
}

func (angleindicationtype *AngleIndicationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(angleindicationtype.GongGetIdentifier(stage), "AngleIndicationType", angleindicationtype.Name)
}

func (authorityfornavaidequipmentpropertytype *AuthorityForNavaidEquipmentPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(authorityfornavaidequipmentpropertytype.GongGetIdentifier(stage), "AuthorityForNavaidEquipmentPropertyType", authorityfornavaidequipmentpropertytype.Name)
}

func (authorityfornavaidequipmenttype *AuthorityForNavaidEquipmentType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(authorityfornavaidequipmenttype.GongGetIdentifier(stage), "AuthorityForNavaidEquipmentType", authorityfornavaidequipmenttype.Name)
}

func (authorityforspecialnavigationstationpropertytype *AuthorityForSpecialNavigationStationPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(authorityforspecialnavigationstationpropertytype.GongGetIdentifier(stage), "AuthorityForSpecialNavigationStationPropertyType", authorityforspecialnavigationstationpropertytype.Name)
}

func (authorityforspecialnavigationstationtype *AuthorityForSpecialNavigationStationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(authorityforspecialnavigationstationtype.GongGetIdentifier(stage), "AuthorityForSpecialNavigationStationType", authorityforspecialnavigationstationtype.Name)
}

func (authorityforspecialnavigationsystempropertytype *AuthorityForSpecialNavigationSystemPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(authorityforspecialnavigationsystempropertytype.GongGetIdentifier(stage), "AuthorityForSpecialNavigationSystemPropertyType", authorityforspecialnavigationsystempropertytype.Name)
}

func (authorityforspecialnavigationsystemtype *AuthorityForSpecialNavigationSystemType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(authorityforspecialnavigationsystemtype.GongGetIdentifier(stage), "AuthorityForSpecialNavigationSystemType", authorityforspecialnavigationsystemtype.Name)
}

func (azimuthtimeslicepropertytype *AzimuthTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(azimuthtimeslicepropertytype.GongGetIdentifier(stage), "AzimuthTimeSlicePropertyType", azimuthtimeslicepropertytype.Name)
}

func (azimuthtimeslicetype *AzimuthTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(azimuthtimeslicetype.GongGetIdentifier(stage), "AzimuthTimeSliceType", azimuthtimeslicetype.Name)
}

func (azimuthtype *AzimuthType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(azimuthtype.GongGetIdentifier(stage), "AzimuthType", azimuthtype.Name)
}

func (checkpointinstimeslicepropertytype *CheckpointINSTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(checkpointinstimeslicepropertytype.GongGetIdentifier(stage), "CheckpointINSTimeSlicePropertyType", checkpointinstimeslicepropertytype.Name)
}

func (checkpointinstimeslicetype *CheckpointINSTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(checkpointinstimeslicetype.GongGetIdentifier(stage), "CheckpointINSTimeSliceType", checkpointinstimeslicetype.Name)
}

func (checkpointinstype *CheckpointINSType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(checkpointinstype.GongGetIdentifier(stage), "CheckpointINSType", checkpointinstype.Name)
}

func (checkpointvortimeslicepropertytype *CheckpointVORTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(checkpointvortimeslicepropertytype.GongGetIdentifier(stage), "CheckpointVORTimeSlicePropertyType", checkpointvortimeslicepropertytype.Name)
}

func (checkpointvortimeslicetype *CheckpointVORTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(checkpointvortimeslicetype.GongGetIdentifier(stage), "CheckpointVORTimeSliceType", checkpointvortimeslicetype.Name)
}

func (checkpointvortype *CheckpointVORType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(checkpointvortype.GongGetIdentifier(stage), "CheckpointVORType", checkpointvortype.Name)
}

func (contactinformationpropertytype *ContactInformationPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(contactinformationpropertytype.GongGetIdentifier(stage), "ContactInformationPropertyType", contactinformationpropertytype.Name)
}

func (contactinformationtype *ContactInformationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(contactinformationtype.GongGetIdentifier(stage), "ContactInformationType", contactinformationtype.Name)
}

func (dmetimeslicepropertytype *DMETimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(dmetimeslicepropertytype.GongGetIdentifier(stage), "DMETimeSlicePropertyType", dmetimeslicepropertytype.Name)
}

func (dmetimeslicetype *DMETimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(dmetimeslicetype.GongGetIdentifier(stage), "DMETimeSliceType", dmetimeslicetype.Name)
}

func (dmetype *DMEType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(dmetype.GongGetIdentifier(stage), "DMEType", dmetype.Name)
}

func (designatedpointpropertytype *DesignatedPointPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(designatedpointpropertytype.GongGetIdentifier(stage), "DesignatedPointPropertyType", designatedpointpropertytype.Name)
}

func (designatedpointtimeslicepropertytype *DesignatedPointTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(designatedpointtimeslicepropertytype.GongGetIdentifier(stage), "DesignatedPointTimeSlicePropertyType", designatedpointtimeslicepropertytype.Name)
}

func (designatedpointtimeslicetype *DesignatedPointTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(designatedpointtimeslicetype.GongGetIdentifier(stage), "DesignatedPointTimeSliceType", designatedpointtimeslicetype.Name)
}

func (designatedpointtype *DesignatedPointType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(designatedpointtype.GongGetIdentifier(stage), "DesignatedPointType", designatedpointtype.Name)
}

func (directionfindertimeslicepropertytype *DirectionFinderTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(directionfindertimeslicepropertytype.GongGetIdentifier(stage), "DirectionFinderTimeSlicePropertyType", directionfindertimeslicepropertytype.Name)
}

func (directionfindertimeslicetype *DirectionFinderTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(directionfindertimeslicetype.GongGetIdentifier(stage), "DirectionFinderTimeSliceType", directionfindertimeslicetype.Name)
}

func (directionfindertype *DirectionFinderType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(directionfindertype.GongGetIdentifier(stage), "DirectionFinderType", directionfindertype.Name)
}

func (distanceindicationtimeslicepropertytype *DistanceIndicationTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(distanceindicationtimeslicepropertytype.GongGetIdentifier(stage), "DistanceIndicationTimeSlicePropertyType", distanceindicationtimeslicepropertytype.Name)
}

func (distanceindicationtimeslicetype *DistanceIndicationTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(distanceindicationtimeslicetype.GongGetIdentifier(stage), "DistanceIndicationTimeSliceType", distanceindicationtimeslicetype.Name)
}

func (distanceindicationtype *DistanceIndicationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(distanceindicationtype.GongGetIdentifier(stage), "DistanceIndicationType", distanceindicationtype.Name)
}

func (elevatedpointpropertytype *ElevatedPointPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(elevatedpointpropertytype.GongGetIdentifier(stage), "ElevatedPointPropertyType", elevatedpointpropertytype.Name)
}

func (elevatedpointtype *ElevatedPointType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(elevatedpointtype.GongGetIdentifier(stage), "ElevatedPointType", elevatedpointtype.Name)
}

func (elevationtimeslicepropertytype *ElevationTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(elevationtimeslicepropertytype.GongGetIdentifier(stage), "ElevationTimeSlicePropertyType", elevationtimeslicepropertytype.Name)
}

func (elevationtimeslicetype *ElevationTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(elevationtimeslicetype.GongGetIdentifier(stage), "ElevationTimeSliceType", elevationtimeslicetype.Name)
}

func (elevationtype *ElevationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(elevationtype.GongGetIdentifier(stage), "ElevationType", elevationtype.Name)
}

func (glidepathtimeslicepropertytype *GlidepathTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(glidepathtimeslicepropertytype.GongGetIdentifier(stage), "GlidepathTimeSlicePropertyType", glidepathtimeslicepropertytype.Name)
}

func (glidepathtimeslicetype *GlidepathTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(glidepathtimeslicetype.GongGetIdentifier(stage), "GlidepathTimeSliceType", glidepathtimeslicetype.Name)
}

func (glidepathtype *GlidepathType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(glidepathtype.GongGetIdentifier(stage), "GlidepathType", glidepathtype.Name)
}

func (informationservicepropertytype *InformationServicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(informationservicepropertytype.GongGetIdentifier(stage), "InformationServicePropertyType", informationservicepropertytype.Name)
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(linguisticnotepropertytype.GongGetIdentifier(stage), "LinguisticNotePropertyType", linguisticnotepropertytype.Name)
}

func (linguisticnotetype *LinguisticNoteType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(linguisticnotetype.GongGetIdentifier(stage), "LinguisticNoteType", linguisticnotetype.Name)
}

func (localizertimeslicepropertytype *LocalizerTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(localizertimeslicepropertytype.GongGetIdentifier(stage), "LocalizerTimeSlicePropertyType", localizertimeslicepropertytype.Name)
}

func (localizertimeslicetype *LocalizerTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(localizertimeslicetype.GongGetIdentifier(stage), "LocalizerTimeSliceType", localizertimeslicetype.Name)
}

func (localizertype *LocalizerType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(localizertype.GongGetIdentifier(stage), "LocalizerType", localizertype.Name)
}

func (markerbeacontimeslicepropertytype *MarkerBeaconTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(markerbeacontimeslicepropertytype.GongGetIdentifier(stage), "MarkerBeaconTimeSlicePropertyType", markerbeacontimeslicepropertytype.Name)
}

func (markerbeacontimeslicetype *MarkerBeaconTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(markerbeacontimeslicetype.GongGetIdentifier(stage), "MarkerBeaconTimeSliceType", markerbeacontimeslicetype.Name)
}

func (markerbeacontype *MarkerBeaconType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(markerbeacontype.GongGetIdentifier(stage), "MarkerBeaconType", markerbeacontype.Name)
}

func (ndbtimeslicepropertytype *NDBTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(ndbtimeslicepropertytype.GongGetIdentifier(stage), "NDBTimeSlicePropertyType", ndbtimeslicepropertytype.Name)
}

func (ndbtimeslicetype *NDBTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(ndbtimeslicetype.GongGetIdentifier(stage), "NDBTimeSliceType", ndbtimeslicetype.Name)
}

func (ndbtype *NDBType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(ndbtype.GongGetIdentifier(stage), "NDBType", ndbtype.Name)
}

func (navaidcomponentpropertytype *NavaidComponentPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(navaidcomponentpropertytype.GongGetIdentifier(stage), "NavaidComponentPropertyType", navaidcomponentpropertytype.Name)
}

func (navaidcomponenttype *NavaidComponentType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(navaidcomponenttype.GongGetIdentifier(stage), "NavaidComponentType", navaidcomponenttype.Name)
}

func (navaidequipmentmonitoringpropertytype *NavaidEquipmentMonitoringPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(navaidequipmentmonitoringpropertytype.GongGetIdentifier(stage), "NavaidEquipmentMonitoringPropertyType", navaidequipmentmonitoringpropertytype.Name)
}

func (navaidequipmentmonitoringtype *NavaidEquipmentMonitoringType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(navaidequipmentmonitoringtype.GongGetIdentifier(stage), "NavaidEquipmentMonitoringType", navaidequipmentmonitoringtype.Name)
}

func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(navaidequipmentpropertytype.GongGetIdentifier(stage), "NavaidEquipmentPropertyType", navaidequipmentpropertytype.Name)
}

func (navaidoperationalstatuspropertytype *NavaidOperationalStatusPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(navaidoperationalstatuspropertytype.GongGetIdentifier(stage), "NavaidOperationalStatusPropertyType", navaidoperationalstatuspropertytype.Name)
}

func (navaidoperationalstatustype *NavaidOperationalStatusType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(navaidoperationalstatustype.GongGetIdentifier(stage), "NavaidOperationalStatusType", navaidoperationalstatustype.Name)
}

func (navaidpropertytype *NavaidPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(navaidpropertytype.GongGetIdentifier(stage), "NavaidPropertyType", navaidpropertytype.Name)
}

func (navaidtimeslicepropertytype *NavaidTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(navaidtimeslicepropertytype.GongGetIdentifier(stage), "NavaidTimeSlicePropertyType", navaidtimeslicepropertytype.Name)
}

func (navaidtimeslicetype *NavaidTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(navaidtimeslicetype.GongGetIdentifier(stage), "NavaidTimeSliceType", navaidtimeslicetype.Name)
}

func (navaidtype *NavaidType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(navaidtype.GongGetIdentifier(stage), "NavaidType", navaidtype.Name)
}

func (notepropertytype *NotePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(notepropertytype.GongGetIdentifier(stage), "NotePropertyType", notepropertytype.Name)
}

func (notetype *NoteType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(notetype.GongGetIdentifier(stage), "NoteType", notetype.Name)
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

func (pointtype *PointType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(pointtype.GongGetIdentifier(stage), "PointType", pointtype.Name)
}

func (postaladdresspropertytype *PostalAddressPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(postaladdresspropertytype.GongGetIdentifier(stage), "PostalAddressPropertyType", postaladdresspropertytype.Name)
}

func (postaladdresstype *PostalAddressType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(postaladdresstype.GongGetIdentifier(stage), "PostalAddressType", postaladdresstype.Name)
}

func (precisionapproachradartimeslicepropertytype *PrecisionApproachRadarTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(precisionapproachradartimeslicepropertytype.GongGetIdentifier(stage), "PrecisionApproachRadarTimeSlicePropertyType", precisionapproachradartimeslicepropertytype.Name)
}

func (precisionapproachradartimeslicetype *PrecisionApproachRadarTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(precisionapproachradartimeslicetype.GongGetIdentifier(stage), "PrecisionApproachRadarTimeSliceType", precisionapproachradartimeslicetype.Name)
}

func (precisionapproachradartype *PrecisionApproachRadarType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(precisionapproachradartype.GongGetIdentifier(stage), "PrecisionApproachRadarType", precisionapproachradartype.Name)
}

func (primarysurveillanceradartimeslicepropertytype *PrimarySurveillanceRadarTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(primarysurveillanceradartimeslicepropertytype.GongGetIdentifier(stage), "PrimarySurveillanceRadarTimeSlicePropertyType", primarysurveillanceradartimeslicepropertytype.Name)
}

func (primarysurveillanceradartimeslicetype *PrimarySurveillanceRadarTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(primarysurveillanceradartimeslicetype.GongGetIdentifier(stage), "PrimarySurveillanceRadarTimeSliceType", primarysurveillanceradartimeslicetype.Name)
}

func (primarysurveillanceradartype *PrimarySurveillanceRadarType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(primarysurveillanceradartype.GongGetIdentifier(stage), "PrimarySurveillanceRadarType", primarysurveillanceradartype.Name)
}

func (radarcomponentpropertytype *RadarComponentPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(radarcomponentpropertytype.GongGetIdentifier(stage), "RadarComponentPropertyType", radarcomponentpropertytype.Name)
}

func (radarcomponenttype *RadarComponentType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(radarcomponenttype.GongGetIdentifier(stage), "RadarComponentType", radarcomponenttype.Name)
}

func (radarequipmentpropertytype *RadarEquipmentPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(radarequipmentpropertytype.GongGetIdentifier(stage), "RadarEquipmentPropertyType", radarequipmentpropertytype.Name)
}

func (radarsystemtimeslicepropertytype *RadarSystemTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(radarsystemtimeslicepropertytype.GongGetIdentifier(stage), "RadarSystemTimeSlicePropertyType", radarsystemtimeslicepropertytype.Name)
}

func (radarsystemtimeslicetype *RadarSystemTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(radarsystemtimeslicetype.GongGetIdentifier(stage), "RadarSystemTimeSliceType", radarsystemtimeslicetype.Name)
}

func (radarsystemtype *RadarSystemType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(radarsystemtype.GongGetIdentifier(stage), "RadarSystemType", radarsystemtype.Name)
}

func (reflectorpropertytype *ReflectorPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(reflectorpropertytype.GongGetIdentifier(stage), "ReflectorPropertyType", reflectorpropertytype.Name)
}

func (reflectortype *ReflectorType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(reflectortype.GongGetIdentifier(stage), "ReflectorType", reflectortype.Name)
}

func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(runwaycentrelinepointpropertytype.GongGetIdentifier(stage), "RunwayCentrelinePointPropertyType", runwaycentrelinepointpropertytype.Name)
}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(runwaydirectionpropertytype.GongGetIdentifier(stage), "RunwayDirectionPropertyType", runwaydirectionpropertytype.Name)
}

func (runwaypropertytype *RunwayPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(runwaypropertytype.GongGetIdentifier(stage), "RunwayPropertyType", runwaypropertytype.Name)
}

func (sdftimeslicepropertytype *SDFTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(sdftimeslicepropertytype.GongGetIdentifier(stage), "SDFTimeSlicePropertyType", sdftimeslicepropertytype.Name)
}

func (sdftimeslicetype *SDFTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(sdftimeslicetype.GongGetIdentifier(stage), "SDFTimeSliceType", sdftimeslicetype.Name)
}

func (sdftype *SDFType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(sdftype.GongGetIdentifier(stage), "SDFType", sdftype.Name)
}

func (secondarysurveillanceradartimeslicepropertytype *SecondarySurveillanceRadarTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(secondarysurveillanceradartimeslicepropertytype.GongGetIdentifier(stage), "SecondarySurveillanceRadarTimeSlicePropertyType", secondarysurveillanceradartimeslicepropertytype.Name)
}

func (secondarysurveillanceradartimeslicetype *SecondarySurveillanceRadarTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(secondarysurveillanceradartimeslicetype.GongGetIdentifier(stage), "SecondarySurveillanceRadarTimeSliceType", secondarysurveillanceradartimeslicetype.Name)
}

func (secondarysurveillanceradartype *SecondarySurveillanceRadarType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(secondarysurveillanceradartype.GongGetIdentifier(stage), "SecondarySurveillanceRadarType", secondarysurveillanceradartype.Name)
}

func (significantpointinairspacetimeslicepropertytype *SignificantPointInAirspaceTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(significantpointinairspacetimeslicepropertytype.GongGetIdentifier(stage), "SignificantPointInAirspaceTimeSlicePropertyType", significantpointinairspacetimeslicepropertytype.Name)
}

func (significantpointinairspacetimeslicetype *SignificantPointInAirspaceTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(significantpointinairspacetimeslicetype.GongGetIdentifier(stage), "SignificantPointInAirspaceTimeSliceType", significantpointinairspacetimeslicetype.Name)
}

func (significantpointinairspacetype *SignificantPointInAirspaceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(significantpointinairspacetype.GongGetIdentifier(stage), "SignificantPointInAirspaceType", significantpointinairspacetype.Name)
}

func (specialnavigationstationstatuspropertytype *SpecialNavigationStationStatusPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(specialnavigationstationstatuspropertytype.GongGetIdentifier(stage), "SpecialNavigationStationStatusPropertyType", specialnavigationstationstatuspropertytype.Name)
}

func (specialnavigationstationstatustype *SpecialNavigationStationStatusType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(specialnavigationstationstatustype.GongGetIdentifier(stage), "SpecialNavigationStationStatusType", specialnavigationstationstatustype.Name)
}

func (specialnavigationstationtimeslicepropertytype *SpecialNavigationStationTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(specialnavigationstationtimeslicepropertytype.GongGetIdentifier(stage), "SpecialNavigationStationTimeSlicePropertyType", specialnavigationstationtimeslicepropertytype.Name)
}

func (specialnavigationstationtimeslicetype *SpecialNavigationStationTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(specialnavigationstationtimeslicetype.GongGetIdentifier(stage), "SpecialNavigationStationTimeSliceType", specialnavigationstationtimeslicetype.Name)
}

func (specialnavigationstationtype *SpecialNavigationStationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(specialnavigationstationtype.GongGetIdentifier(stage), "SpecialNavigationStationType", specialnavigationstationtype.Name)
}

func (specialnavigationsystempropertytype *SpecialNavigationSystemPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(specialnavigationsystempropertytype.GongGetIdentifier(stage), "SpecialNavigationSystemPropertyType", specialnavigationsystempropertytype.Name)
}

func (specialnavigationsystemtimeslicepropertytype *SpecialNavigationSystemTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(specialnavigationsystemtimeslicepropertytype.GongGetIdentifier(stage), "SpecialNavigationSystemTimeSlicePropertyType", specialnavigationsystemtimeslicepropertytype.Name)
}

func (specialnavigationsystemtimeslicetype *SpecialNavigationSystemTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(specialnavigationsystemtimeslicetype.GongGetIdentifier(stage), "SpecialNavigationSystemTimeSliceType", specialnavigationsystemtimeslicetype.Name)
}

func (specialnavigationsystemtype *SpecialNavigationSystemType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(specialnavigationsystemtype.GongGetIdentifier(stage), "SpecialNavigationSystemType", specialnavigationsystemtype.Name)
}

func (surveillancegroundstationpropertytype *SurveillanceGroundStationPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(surveillancegroundstationpropertytype.GongGetIdentifier(stage), "SurveillanceGroundStationPropertyType", surveillancegroundstationpropertytype.Name)
}

func (surveillancegroundstationtype *SurveillanceGroundStationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(surveillancegroundstationtype.GongGetIdentifier(stage), "SurveillanceGroundStationType", surveillancegroundstationtype.Name)
}

func (tacantimeslicepropertytype *TACANTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tacantimeslicepropertytype.GongGetIdentifier(stage), "TACANTimeSlicePropertyType", tacantimeslicepropertytype.Name)
}

func (tacantimeslicetype *TACANTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tacantimeslicetype.GongGetIdentifier(stage), "TACANTimeSliceType", tacantimeslicetype.Name)
}

func (tacantype *TACANType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tacantype.GongGetIdentifier(stage), "TACANType", tacantype.Name)
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

func (unitpropertytype *UnitPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(unitpropertytype.GongGetIdentifier(stage), "UnitPropertyType", unitpropertytype.Name)
}

func (vorpropertytype *VORPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(vorpropertytype.GongGetIdentifier(stage), "VORPropertyType", vorpropertytype.Name)
}

func (vortimeslicepropertytype *VORTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(vortimeslicepropertytype.GongGetIdentifier(stage), "VORTimeSlicePropertyType", vortimeslicepropertytype.Name)
}

func (vortimeslicetype *VORTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(vortimeslicetype.GongGetIdentifier(stage), "VORTimeSliceType", vortimeslicetype.Name)
}

func (vortype *VORType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(vortype.GongGetIdentifier(stage), "VORType", vortype.Name)
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

func (abstractnavaidequipmenttype *AbstractNavaidEquipmentType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractnavaidequipmenttype.GongGetReferenceIdentifier(stage))
}

func (abstractnavigationsystemcheckpointtype *AbstractNavigationSystemCheckpointType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractnavigationsystemcheckpointtype.GongGetReferenceIdentifier(stage))
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractpropertieswithscheduletype.GongGetReferenceIdentifier(stage))
}

func (abstractradarequipmenttype *AbstractRadarEquipmentType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractradarequipmenttype.GongGetReferenceIdentifier(stage))
}

func (abstractsurveillanceradartype *AbstractSurveillanceRadarType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractsurveillanceradartype.GongGetReferenceIdentifier(stage))
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airportheliportpropertytype.GongGetReferenceIdentifier(stage))
}

func (airspacepropertytype *AirspacePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airspacepropertytype.GongGetReferenceIdentifier(stage))
}

func (angleindicationtimeslicepropertytype *AngleIndicationTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(angleindicationtimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (angleindicationtimeslicetype *AngleIndicationTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(angleindicationtimeslicetype.GongGetReferenceIdentifier(stage))
}

func (angleindicationtype *AngleIndicationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(angleindicationtype.GongGetReferenceIdentifier(stage))
}

func (authorityfornavaidequipmentpropertytype *AuthorityForNavaidEquipmentPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(authorityfornavaidequipmentpropertytype.GongGetReferenceIdentifier(stage))
}

func (authorityfornavaidequipmenttype *AuthorityForNavaidEquipmentType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(authorityfornavaidequipmenttype.GongGetReferenceIdentifier(stage))
}

func (authorityforspecialnavigationstationpropertytype *AuthorityForSpecialNavigationStationPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(authorityforspecialnavigationstationpropertytype.GongGetReferenceIdentifier(stage))
}

func (authorityforspecialnavigationstationtype *AuthorityForSpecialNavigationStationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(authorityforspecialnavigationstationtype.GongGetReferenceIdentifier(stage))
}

func (authorityforspecialnavigationsystempropertytype *AuthorityForSpecialNavigationSystemPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(authorityforspecialnavigationsystempropertytype.GongGetReferenceIdentifier(stage))
}

func (authorityforspecialnavigationsystemtype *AuthorityForSpecialNavigationSystemType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(authorityforspecialnavigationsystemtype.GongGetReferenceIdentifier(stage))
}

func (azimuthtimeslicepropertytype *AzimuthTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(azimuthtimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (azimuthtimeslicetype *AzimuthTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(azimuthtimeslicetype.GongGetReferenceIdentifier(stage))
}

func (azimuthtype *AzimuthType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(azimuthtype.GongGetReferenceIdentifier(stage))
}

func (checkpointinstimeslicepropertytype *CheckpointINSTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(checkpointinstimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (checkpointinstimeslicetype *CheckpointINSTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(checkpointinstimeslicetype.GongGetReferenceIdentifier(stage))
}

func (checkpointinstype *CheckpointINSType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(checkpointinstype.GongGetReferenceIdentifier(stage))
}

func (checkpointvortimeslicepropertytype *CheckpointVORTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(checkpointvortimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (checkpointvortimeslicetype *CheckpointVORTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(checkpointvortimeslicetype.GongGetReferenceIdentifier(stage))
}

func (checkpointvortype *CheckpointVORType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(checkpointvortype.GongGetReferenceIdentifier(stage))
}

func (contactinformationpropertytype *ContactInformationPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(contactinformationpropertytype.GongGetReferenceIdentifier(stage))
}

func (contactinformationtype *ContactInformationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(contactinformationtype.GongGetReferenceIdentifier(stage))
}

func (dmetimeslicepropertytype *DMETimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(dmetimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (dmetimeslicetype *DMETimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(dmetimeslicetype.GongGetReferenceIdentifier(stage))
}

func (dmetype *DMEType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(dmetype.GongGetReferenceIdentifier(stage))
}

func (designatedpointpropertytype *DesignatedPointPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(designatedpointpropertytype.GongGetReferenceIdentifier(stage))
}

func (designatedpointtimeslicepropertytype *DesignatedPointTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(designatedpointtimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (designatedpointtimeslicetype *DesignatedPointTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(designatedpointtimeslicetype.GongGetReferenceIdentifier(stage))
}

func (designatedpointtype *DesignatedPointType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(designatedpointtype.GongGetReferenceIdentifier(stage))
}

func (directionfindertimeslicepropertytype *DirectionFinderTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(directionfindertimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (directionfindertimeslicetype *DirectionFinderTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(directionfindertimeslicetype.GongGetReferenceIdentifier(stage))
}

func (directionfindertype *DirectionFinderType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(directionfindertype.GongGetReferenceIdentifier(stage))
}

func (distanceindicationtimeslicepropertytype *DistanceIndicationTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(distanceindicationtimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (distanceindicationtimeslicetype *DistanceIndicationTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(distanceindicationtimeslicetype.GongGetReferenceIdentifier(stage))
}

func (distanceindicationtype *DistanceIndicationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(distanceindicationtype.GongGetReferenceIdentifier(stage))
}

func (elevatedpointpropertytype *ElevatedPointPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(elevatedpointpropertytype.GongGetReferenceIdentifier(stage))
}

func (elevatedpointtype *ElevatedPointType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(elevatedpointtype.GongGetReferenceIdentifier(stage))
}

func (elevationtimeslicepropertytype *ElevationTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(elevationtimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (elevationtimeslicetype *ElevationTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(elevationtimeslicetype.GongGetReferenceIdentifier(stage))
}

func (elevationtype *ElevationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(elevationtype.GongGetReferenceIdentifier(stage))
}

func (glidepathtimeslicepropertytype *GlidepathTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(glidepathtimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (glidepathtimeslicetype *GlidepathTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(glidepathtimeslicetype.GongGetReferenceIdentifier(stage))
}

func (glidepathtype *GlidepathType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(glidepathtype.GongGetReferenceIdentifier(stage))
}

func (informationservicepropertytype *InformationServicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(informationservicepropertytype.GongGetReferenceIdentifier(stage))
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(linguisticnotepropertytype.GongGetReferenceIdentifier(stage))
}

func (linguisticnotetype *LinguisticNoteType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(linguisticnotetype.GongGetReferenceIdentifier(stage))
}

func (localizertimeslicepropertytype *LocalizerTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(localizertimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (localizertimeslicetype *LocalizerTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(localizertimeslicetype.GongGetReferenceIdentifier(stage))
}

func (localizertype *LocalizerType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(localizertype.GongGetReferenceIdentifier(stage))
}

func (markerbeacontimeslicepropertytype *MarkerBeaconTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(markerbeacontimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (markerbeacontimeslicetype *MarkerBeaconTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(markerbeacontimeslicetype.GongGetReferenceIdentifier(stage))
}

func (markerbeacontype *MarkerBeaconType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(markerbeacontype.GongGetReferenceIdentifier(stage))
}

func (ndbtimeslicepropertytype *NDBTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(ndbtimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (ndbtimeslicetype *NDBTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(ndbtimeslicetype.GongGetReferenceIdentifier(stage))
}

func (ndbtype *NDBType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(ndbtype.GongGetReferenceIdentifier(stage))
}

func (navaidcomponentpropertytype *NavaidComponentPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(navaidcomponentpropertytype.GongGetReferenceIdentifier(stage))
}

func (navaidcomponenttype *NavaidComponentType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(navaidcomponenttype.GongGetReferenceIdentifier(stage))
}

func (navaidequipmentmonitoringpropertytype *NavaidEquipmentMonitoringPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(navaidequipmentmonitoringpropertytype.GongGetReferenceIdentifier(stage))
}

func (navaidequipmentmonitoringtype *NavaidEquipmentMonitoringType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(navaidequipmentmonitoringtype.GongGetReferenceIdentifier(stage))
}

func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(navaidequipmentpropertytype.GongGetReferenceIdentifier(stage))
}

func (navaidoperationalstatuspropertytype *NavaidOperationalStatusPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(navaidoperationalstatuspropertytype.GongGetReferenceIdentifier(stage))
}

func (navaidoperationalstatustype *NavaidOperationalStatusType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(navaidoperationalstatustype.GongGetReferenceIdentifier(stage))
}

func (navaidpropertytype *NavaidPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(navaidpropertytype.GongGetReferenceIdentifier(stage))
}

func (navaidtimeslicepropertytype *NavaidTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(navaidtimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (navaidtimeslicetype *NavaidTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(navaidtimeslicetype.GongGetReferenceIdentifier(stage))
}

func (navaidtype *NavaidType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(navaidtype.GongGetReferenceIdentifier(stage))
}

func (notepropertytype *NotePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(notepropertytype.GongGetReferenceIdentifier(stage))
}

func (notetype *NoteType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(notetype.GongGetReferenceIdentifier(stage))
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

func (pointtype *PointType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(pointtype.GongGetReferenceIdentifier(stage))
}

func (postaladdresspropertytype *PostalAddressPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(postaladdresspropertytype.GongGetReferenceIdentifier(stage))
}

func (postaladdresstype *PostalAddressType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(postaladdresstype.GongGetReferenceIdentifier(stage))
}

func (precisionapproachradartimeslicepropertytype *PrecisionApproachRadarTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(precisionapproachradartimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (precisionapproachradartimeslicetype *PrecisionApproachRadarTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(precisionapproachradartimeslicetype.GongGetReferenceIdentifier(stage))
}

func (precisionapproachradartype *PrecisionApproachRadarType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(precisionapproachradartype.GongGetReferenceIdentifier(stage))
}

func (primarysurveillanceradartimeslicepropertytype *PrimarySurveillanceRadarTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(primarysurveillanceradartimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (primarysurveillanceradartimeslicetype *PrimarySurveillanceRadarTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(primarysurveillanceradartimeslicetype.GongGetReferenceIdentifier(stage))
}

func (primarysurveillanceradartype *PrimarySurveillanceRadarType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(primarysurveillanceradartype.GongGetReferenceIdentifier(stage))
}

func (radarcomponentpropertytype *RadarComponentPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(radarcomponentpropertytype.GongGetReferenceIdentifier(stage))
}

func (radarcomponenttype *RadarComponentType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(radarcomponenttype.GongGetReferenceIdentifier(stage))
}

func (radarequipmentpropertytype *RadarEquipmentPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(radarequipmentpropertytype.GongGetReferenceIdentifier(stage))
}

func (radarsystemtimeslicepropertytype *RadarSystemTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(radarsystemtimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (radarsystemtimeslicetype *RadarSystemTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(radarsystemtimeslicetype.GongGetReferenceIdentifier(stage))
}

func (radarsystemtype *RadarSystemType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(radarsystemtype.GongGetReferenceIdentifier(stage))
}

func (reflectorpropertytype *ReflectorPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(reflectorpropertytype.GongGetReferenceIdentifier(stage))
}

func (reflectortype *ReflectorType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(reflectortype.GongGetReferenceIdentifier(stage))
}

func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(runwaycentrelinepointpropertytype.GongGetReferenceIdentifier(stage))
}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(runwaydirectionpropertytype.GongGetReferenceIdentifier(stage))
}

func (runwaypropertytype *RunwayPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(runwaypropertytype.GongGetReferenceIdentifier(stage))
}

func (sdftimeslicepropertytype *SDFTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(sdftimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (sdftimeslicetype *SDFTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(sdftimeslicetype.GongGetReferenceIdentifier(stage))
}

func (sdftype *SDFType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(sdftype.GongGetReferenceIdentifier(stage))
}

func (secondarysurveillanceradartimeslicepropertytype *SecondarySurveillanceRadarTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(secondarysurveillanceradartimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (secondarysurveillanceradartimeslicetype *SecondarySurveillanceRadarTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(secondarysurveillanceradartimeslicetype.GongGetReferenceIdentifier(stage))
}

func (secondarysurveillanceradartype *SecondarySurveillanceRadarType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(secondarysurveillanceradartype.GongGetReferenceIdentifier(stage))
}

func (significantpointinairspacetimeslicepropertytype *SignificantPointInAirspaceTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(significantpointinairspacetimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (significantpointinairspacetimeslicetype *SignificantPointInAirspaceTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(significantpointinairspacetimeslicetype.GongGetReferenceIdentifier(stage))
}

func (significantpointinairspacetype *SignificantPointInAirspaceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(significantpointinairspacetype.GongGetReferenceIdentifier(stage))
}

func (specialnavigationstationstatuspropertytype *SpecialNavigationStationStatusPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(specialnavigationstationstatuspropertytype.GongGetReferenceIdentifier(stage))
}

func (specialnavigationstationstatustype *SpecialNavigationStationStatusType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(specialnavigationstationstatustype.GongGetReferenceIdentifier(stage))
}

func (specialnavigationstationtimeslicepropertytype *SpecialNavigationStationTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(specialnavigationstationtimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (specialnavigationstationtimeslicetype *SpecialNavigationStationTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(specialnavigationstationtimeslicetype.GongGetReferenceIdentifier(stage))
}

func (specialnavigationstationtype *SpecialNavigationStationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(specialnavigationstationtype.GongGetReferenceIdentifier(stage))
}

func (specialnavigationsystempropertytype *SpecialNavigationSystemPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(specialnavigationsystempropertytype.GongGetReferenceIdentifier(stage))
}

func (specialnavigationsystemtimeslicepropertytype *SpecialNavigationSystemTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(specialnavigationsystemtimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (specialnavigationsystemtimeslicetype *SpecialNavigationSystemTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(specialnavigationsystemtimeslicetype.GongGetReferenceIdentifier(stage))
}

func (specialnavigationsystemtype *SpecialNavigationSystemType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(specialnavigationsystemtype.GongGetReferenceIdentifier(stage))
}

func (surveillancegroundstationpropertytype *SurveillanceGroundStationPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(surveillancegroundstationpropertytype.GongGetReferenceIdentifier(stage))
}

func (surveillancegroundstationtype *SurveillanceGroundStationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(surveillancegroundstationtype.GongGetReferenceIdentifier(stage))
}

func (tacantimeslicepropertytype *TACANTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tacantimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (tacantimeslicetype *TACANTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tacantimeslicetype.GongGetReferenceIdentifier(stage))
}

func (tacantype *TACANType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tacantype.GongGetReferenceIdentifier(stage))
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

func (unitpropertytype *UnitPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(unitpropertytype.GongGetReferenceIdentifier(stage))
}

func (vorpropertytype *VORPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(vorpropertytype.GongGetReferenceIdentifier(stage))
}

func (vortimeslicepropertytype *VORTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(vortimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (vortimeslicetype *VORTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(vortimeslicetype.GongGetReferenceIdentifier(stage))
}

func (vortype *VORType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(vortype.GongGetReferenceIdentifier(stage))
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
