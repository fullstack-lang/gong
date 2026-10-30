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
	// Compute reverse map for named struct AirTrafficControlServiceTimeSliceType
	// insertion point per field
	stage.AirTrafficControlServiceTimeSliceType_Call_sign_reverseMap = make(map[*CallsignDetailPropertyType]*AirTrafficControlServiceTimeSliceType)
	for airtrafficcontrolservicetimeslicetype := range stage.AirTrafficControlServiceTimeSliceTypes {
		_ = airtrafficcontrolservicetimeslicetype
		for _, _callsigndetailpropertytype := range airtrafficcontrolservicetimeslicetype.Call_sign {
			stage.AirTrafficControlServiceTimeSliceType_Call_sign_reverseMap[_callsigndetailpropertytype] = airtrafficcontrolservicetimeslicetype
		}
	}
	stage.AirTrafficControlServiceTimeSliceType_RadioCommunication_reverseMap = make(map[*RadioCommunicationChannelPropertyType]*AirTrafficControlServiceTimeSliceType)
	for airtrafficcontrolservicetimeslicetype := range stage.AirTrafficControlServiceTimeSliceTypes {
		_ = airtrafficcontrolservicetimeslicetype
		for _, _radiocommunicationchannelpropertytype := range airtrafficcontrolservicetimeslicetype.RadioCommunication {
			stage.AirTrafficControlServiceTimeSliceType_RadioCommunication_reverseMap[_radiocommunicationchannelpropertytype] = airtrafficcontrolservicetimeslicetype
		}
	}
	stage.AirTrafficControlServiceTimeSliceType_GroundCommunication_reverseMap = make(map[*ContactInformationPropertyType]*AirTrafficControlServiceTimeSliceType)
	for airtrafficcontrolservicetimeslicetype := range stage.AirTrafficControlServiceTimeSliceTypes {
		_ = airtrafficcontrolservicetimeslicetype
		for _, _contactinformationpropertytype := range airtrafficcontrolservicetimeslicetype.GroundCommunication {
			stage.AirTrafficControlServiceTimeSliceType_GroundCommunication_reverseMap[_contactinformationpropertytype] = airtrafficcontrolservicetimeslicetype
		}
	}
	stage.AirTrafficControlServiceTimeSliceType_Availability_reverseMap = make(map[*ServiceOperationalStatusPropertyType]*AirTrafficControlServiceTimeSliceType)
	for airtrafficcontrolservicetimeslicetype := range stage.AirTrafficControlServiceTimeSliceTypes {
		_ = airtrafficcontrolservicetimeslicetype
		for _, _serviceoperationalstatuspropertytype := range airtrafficcontrolservicetimeslicetype.Availability {
			stage.AirTrafficControlServiceTimeSliceType_Availability_reverseMap[_serviceoperationalstatuspropertytype] = airtrafficcontrolservicetimeslicetype
		}
	}
	stage.AirTrafficControlServiceTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*AirTrafficControlServiceTimeSliceType)
	for airtrafficcontrolservicetimeslicetype := range stage.AirTrafficControlServiceTimeSliceTypes {
		_ = airtrafficcontrolservicetimeslicetype
		for _, _notepropertytype := range airtrafficcontrolservicetimeslicetype.Annotation {
			stage.AirTrafficControlServiceTimeSliceType_Annotation_reverseMap[_notepropertytype] = airtrafficcontrolservicetimeslicetype
		}
	}
	stage.AirTrafficControlServiceTimeSliceType_ClientAirport_reverseMap = make(map[*AirportHeliportPropertyType]*AirTrafficControlServiceTimeSliceType)
	for airtrafficcontrolservicetimeslicetype := range stage.AirTrafficControlServiceTimeSliceTypes {
		_ = airtrafficcontrolservicetimeslicetype
		for _, _airportheliportpropertytype := range airtrafficcontrolservicetimeslicetype.ClientAirport {
			stage.AirTrafficControlServiceTimeSliceType_ClientAirport_reverseMap[_airportheliportpropertytype] = airtrafficcontrolservicetimeslicetype
		}
	}
	stage.AirTrafficControlServiceTimeSliceType_ClientAirspace_reverseMap = make(map[*AirspacePropertyType]*AirTrafficControlServiceTimeSliceType)
	for airtrafficcontrolservicetimeslicetype := range stage.AirTrafficControlServiceTimeSliceTypes {
		_ = airtrafficcontrolservicetimeslicetype
		for _, _airspacepropertytype := range airtrafficcontrolservicetimeslicetype.ClientAirspace {
			stage.AirTrafficControlServiceTimeSliceType_ClientAirspace_reverseMap[_airspacepropertytype] = airtrafficcontrolservicetimeslicetype
		}
	}
	stage.AirTrafficControlServiceTimeSliceType_ClientRoute_reverseMap = make(map[*RoutePortionPropertyType]*AirTrafficControlServiceTimeSliceType)
	for airtrafficcontrolservicetimeslicetype := range stage.AirTrafficControlServiceTimeSliceTypes {
		_ = airtrafficcontrolservicetimeslicetype
		for _, _routeportionpropertytype := range airtrafficcontrolservicetimeslicetype.ClientRoute {
			stage.AirTrafficControlServiceTimeSliceType_ClientRoute_reverseMap[_routeportionpropertytype] = airtrafficcontrolservicetimeslicetype
		}
	}
	stage.AirTrafficControlServiceTimeSliceType_ClientProcedure_reverseMap = make(map[*ProcedurePropertyType]*AirTrafficControlServiceTimeSliceType)
	for airtrafficcontrolservicetimeslicetype := range stage.AirTrafficControlServiceTimeSliceTypes {
		_ = airtrafficcontrolservicetimeslicetype
		for _, _procedurepropertytype := range airtrafficcontrolservicetimeslicetype.ClientProcedure {
			stage.AirTrafficControlServiceTimeSliceType_ClientProcedure_reverseMap[_procedurepropertytype] = airtrafficcontrolservicetimeslicetype
		}
	}
	stage.AirTrafficControlServiceTimeSliceType_ClientHolding_reverseMap = make(map[*HoldingPatternPropertyType]*AirTrafficControlServiceTimeSliceType)
	for airtrafficcontrolservicetimeslicetype := range stage.AirTrafficControlServiceTimeSliceTypes {
		_ = airtrafficcontrolservicetimeslicetype
		for _, _holdingpatternpropertytype := range airtrafficcontrolservicetimeslicetype.ClientHolding {
			stage.AirTrafficControlServiceTimeSliceType_ClientHolding_reverseMap[_holdingpatternpropertytype] = airtrafficcontrolservicetimeslicetype
		}
	}
	stage.AirTrafficControlServiceTimeSliceType_ClientAerialRefuelling_reverseMap = make(map[*AerialRefuellingPropertyType]*AirTrafficControlServiceTimeSliceType)
	for airtrafficcontrolservicetimeslicetype := range stage.AirTrafficControlServiceTimeSliceTypes {
		_ = airtrafficcontrolservicetimeslicetype
		for _, _aerialrefuellingpropertytype := range airtrafficcontrolservicetimeslicetype.ClientAerialRefuelling {
			stage.AirTrafficControlServiceTimeSliceType_ClientAerialRefuelling_reverseMap[_aerialrefuellingpropertytype] = airtrafficcontrolservicetimeslicetype
		}
	}
	stage.AirTrafficControlServiceTimeSliceType_Extension_reverseMap = make(map[*A_extension_13]*AirTrafficControlServiceTimeSliceType)
	for airtrafficcontrolservicetimeslicetype := range stage.AirTrafficControlServiceTimeSliceTypes {
		_ = airtrafficcontrolservicetimeslicetype
		for _, _a_extension_13 := range airtrafficcontrolservicetimeslicetype.Extension {
			stage.AirTrafficControlServiceTimeSliceType_Extension_reverseMap[_a_extension_13] = airtrafficcontrolservicetimeslicetype
		}
	}

	// Compute reverse map for named struct AirTrafficControlServiceType
	// insertion point per field
	stage.AirTrafficControlServiceType_TimeSlice_reverseMap = make(map[*AirTrafficControlServiceTimeSlicePropertyType]*AirTrafficControlServiceType)
	for airtrafficcontrolservicetype := range stage.AirTrafficControlServiceTypes {
		_ = airtrafficcontrolservicetype
		for _, _airtrafficcontrolservicetimeslicepropertytype := range airtrafficcontrolservicetype.TimeSlice {
			stage.AirTrafficControlServiceType_TimeSlice_reverseMap[_airtrafficcontrolservicetimeslicepropertytype] = airtrafficcontrolservicetype
		}
	}

	// Compute reverse map for named struct AirTrafficManagementServiceTimeSliceType
	// insertion point per field
	stage.AirTrafficManagementServiceTimeSliceType_Call_sign_reverseMap = make(map[*CallsignDetailPropertyType]*AirTrafficManagementServiceTimeSliceType)
	for airtrafficmanagementservicetimeslicetype := range stage.AirTrafficManagementServiceTimeSliceTypes {
		_ = airtrafficmanagementservicetimeslicetype
		for _, _callsigndetailpropertytype := range airtrafficmanagementservicetimeslicetype.Call_sign {
			stage.AirTrafficManagementServiceTimeSliceType_Call_sign_reverseMap[_callsigndetailpropertytype] = airtrafficmanagementservicetimeslicetype
		}
	}
	stage.AirTrafficManagementServiceTimeSliceType_RadioCommunication_reverseMap = make(map[*RadioCommunicationChannelPropertyType]*AirTrafficManagementServiceTimeSliceType)
	for airtrafficmanagementservicetimeslicetype := range stage.AirTrafficManagementServiceTimeSliceTypes {
		_ = airtrafficmanagementservicetimeslicetype
		for _, _radiocommunicationchannelpropertytype := range airtrafficmanagementservicetimeslicetype.RadioCommunication {
			stage.AirTrafficManagementServiceTimeSliceType_RadioCommunication_reverseMap[_radiocommunicationchannelpropertytype] = airtrafficmanagementservicetimeslicetype
		}
	}
	stage.AirTrafficManagementServiceTimeSliceType_GroundCommunication_reverseMap = make(map[*ContactInformationPropertyType]*AirTrafficManagementServiceTimeSliceType)
	for airtrafficmanagementservicetimeslicetype := range stage.AirTrafficManagementServiceTimeSliceTypes {
		_ = airtrafficmanagementservicetimeslicetype
		for _, _contactinformationpropertytype := range airtrafficmanagementservicetimeslicetype.GroundCommunication {
			stage.AirTrafficManagementServiceTimeSliceType_GroundCommunication_reverseMap[_contactinformationpropertytype] = airtrafficmanagementservicetimeslicetype
		}
	}
	stage.AirTrafficManagementServiceTimeSliceType_Availability_reverseMap = make(map[*ServiceOperationalStatusPropertyType]*AirTrafficManagementServiceTimeSliceType)
	for airtrafficmanagementservicetimeslicetype := range stage.AirTrafficManagementServiceTimeSliceTypes {
		_ = airtrafficmanagementservicetimeslicetype
		for _, _serviceoperationalstatuspropertytype := range airtrafficmanagementservicetimeslicetype.Availability {
			stage.AirTrafficManagementServiceTimeSliceType_Availability_reverseMap[_serviceoperationalstatuspropertytype] = airtrafficmanagementservicetimeslicetype
		}
	}
	stage.AirTrafficManagementServiceTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*AirTrafficManagementServiceTimeSliceType)
	for airtrafficmanagementservicetimeslicetype := range stage.AirTrafficManagementServiceTimeSliceTypes {
		_ = airtrafficmanagementservicetimeslicetype
		for _, _notepropertytype := range airtrafficmanagementservicetimeslicetype.Annotation {
			stage.AirTrafficManagementServiceTimeSliceType_Annotation_reverseMap[_notepropertytype] = airtrafficmanagementservicetimeslicetype
		}
	}
	stage.AirTrafficManagementServiceTimeSliceType_ClientAirspace_reverseMap = make(map[*AirspacePropertyType]*AirTrafficManagementServiceTimeSliceType)
	for airtrafficmanagementservicetimeslicetype := range stage.AirTrafficManagementServiceTimeSliceTypes {
		_ = airtrafficmanagementservicetimeslicetype
		for _, _airspacepropertytype := range airtrafficmanagementservicetimeslicetype.ClientAirspace {
			stage.AirTrafficManagementServiceTimeSliceType_ClientAirspace_reverseMap[_airspacepropertytype] = airtrafficmanagementservicetimeslicetype
		}
	}
	stage.AirTrafficManagementServiceTimeSliceType_ClientAerialRefuelling_reverseMap = make(map[*AerialRefuellingPropertyType]*AirTrafficManagementServiceTimeSliceType)
	for airtrafficmanagementservicetimeslicetype := range stage.AirTrafficManagementServiceTimeSliceTypes {
		_ = airtrafficmanagementservicetimeslicetype
		for _, _aerialrefuellingpropertytype := range airtrafficmanagementservicetimeslicetype.ClientAerialRefuelling {
			stage.AirTrafficManagementServiceTimeSliceType_ClientAerialRefuelling_reverseMap[_aerialrefuellingpropertytype] = airtrafficmanagementservicetimeslicetype
		}
	}
	stage.AirTrafficManagementServiceTimeSliceType_ClientRoute_reverseMap = make(map[*RoutePortionPropertyType]*AirTrafficManagementServiceTimeSliceType)
	for airtrafficmanagementservicetimeslicetype := range stage.AirTrafficManagementServiceTimeSliceTypes {
		_ = airtrafficmanagementservicetimeslicetype
		for _, _routeportionpropertytype := range airtrafficmanagementservicetimeslicetype.ClientRoute {
			stage.AirTrafficManagementServiceTimeSliceType_ClientRoute_reverseMap[_routeportionpropertytype] = airtrafficmanagementservicetimeslicetype
		}
	}
	stage.AirTrafficManagementServiceTimeSliceType_Extension_reverseMap = make(map[*A_extension_14]*AirTrafficManagementServiceTimeSliceType)
	for airtrafficmanagementservicetimeslicetype := range stage.AirTrafficManagementServiceTimeSliceTypes {
		_ = airtrafficmanagementservicetimeslicetype
		for _, _a_extension_14 := range airtrafficmanagementservicetimeslicetype.Extension {
			stage.AirTrafficManagementServiceTimeSliceType_Extension_reverseMap[_a_extension_14] = airtrafficmanagementservicetimeslicetype
		}
	}

	// Compute reverse map for named struct AirTrafficManagementServiceType
	// insertion point per field
	stage.AirTrafficManagementServiceType_TimeSlice_reverseMap = make(map[*AirTrafficManagementServiceTimeSlicePropertyType]*AirTrafficManagementServiceType)
	for airtrafficmanagementservicetype := range stage.AirTrafficManagementServiceTypes {
		_ = airtrafficmanagementservicetype
		for _, _airtrafficmanagementservicetimeslicepropertytype := range airtrafficmanagementservicetype.TimeSlice {
			stage.AirTrafficManagementServiceType_TimeSlice_reverseMap[_airtrafficmanagementservicetimeslicepropertytype] = airtrafficmanagementservicetype
		}
	}

	// Compute reverse map for named struct AircraftGroundServiceTimeSliceType
	// insertion point per field
	stage.AircraftGroundServiceTimeSliceType_Call_sign_reverseMap = make(map[*CallsignDetailPropertyType]*AircraftGroundServiceTimeSliceType)
	for aircraftgroundservicetimeslicetype := range stage.AircraftGroundServiceTimeSliceTypes {
		_ = aircraftgroundservicetimeslicetype
		for _, _callsigndetailpropertytype := range aircraftgroundservicetimeslicetype.Call_sign {
			stage.AircraftGroundServiceTimeSliceType_Call_sign_reverseMap[_callsigndetailpropertytype] = aircraftgroundservicetimeslicetype
		}
	}
	stage.AircraftGroundServiceTimeSliceType_RadioCommunication_reverseMap = make(map[*RadioCommunicationChannelPropertyType]*AircraftGroundServiceTimeSliceType)
	for aircraftgroundservicetimeslicetype := range stage.AircraftGroundServiceTimeSliceTypes {
		_ = aircraftgroundservicetimeslicetype
		for _, _radiocommunicationchannelpropertytype := range aircraftgroundservicetimeslicetype.RadioCommunication {
			stage.AircraftGroundServiceTimeSliceType_RadioCommunication_reverseMap[_radiocommunicationchannelpropertytype] = aircraftgroundservicetimeslicetype
		}
	}
	stage.AircraftGroundServiceTimeSliceType_GroundCommunication_reverseMap = make(map[*ContactInformationPropertyType]*AircraftGroundServiceTimeSliceType)
	for aircraftgroundservicetimeslicetype := range stage.AircraftGroundServiceTimeSliceTypes {
		_ = aircraftgroundservicetimeslicetype
		for _, _contactinformationpropertytype := range aircraftgroundservicetimeslicetype.GroundCommunication {
			stage.AircraftGroundServiceTimeSliceType_GroundCommunication_reverseMap[_contactinformationpropertytype] = aircraftgroundservicetimeslicetype
		}
	}
	stage.AircraftGroundServiceTimeSliceType_Availability_reverseMap = make(map[*ServiceOperationalStatusPropertyType]*AircraftGroundServiceTimeSliceType)
	for aircraftgroundservicetimeslicetype := range stage.AircraftGroundServiceTimeSliceTypes {
		_ = aircraftgroundservicetimeslicetype
		for _, _serviceoperationalstatuspropertytype := range aircraftgroundservicetimeslicetype.Availability {
			stage.AircraftGroundServiceTimeSliceType_Availability_reverseMap[_serviceoperationalstatuspropertytype] = aircraftgroundservicetimeslicetype
		}
	}
	stage.AircraftGroundServiceTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*AircraftGroundServiceTimeSliceType)
	for aircraftgroundservicetimeslicetype := range stage.AircraftGroundServiceTimeSliceTypes {
		_ = aircraftgroundservicetimeslicetype
		for _, _notepropertytype := range aircraftgroundservicetimeslicetype.Annotation {
			stage.AircraftGroundServiceTimeSliceType_Annotation_reverseMap[_notepropertytype] = aircraftgroundservicetimeslicetype
		}
	}
	stage.AircraftGroundServiceTimeSliceType_AirportHeliport_reverseMap = make(map[*AirportHeliportPropertyType]*AircraftGroundServiceTimeSliceType)
	for aircraftgroundservicetimeslicetype := range stage.AircraftGroundServiceTimeSliceTypes {
		_ = aircraftgroundservicetimeslicetype
		for _, _airportheliportpropertytype := range aircraftgroundservicetimeslicetype.AirportHeliport {
			stage.AircraftGroundServiceTimeSliceType_AirportHeliport_reverseMap[_airportheliportpropertytype] = aircraftgroundservicetimeslicetype
		}
	}
	stage.AircraftGroundServiceTimeSliceType_Extension_reverseMap = make(map[*A_extension_10]*AircraftGroundServiceTimeSliceType)
	for aircraftgroundservicetimeslicetype := range stage.AircraftGroundServiceTimeSliceTypes {
		_ = aircraftgroundservicetimeslicetype
		for _, _a_extension_10 := range aircraftgroundservicetimeslicetype.Extension {
			stage.AircraftGroundServiceTimeSliceType_Extension_reverseMap[_a_extension_10] = aircraftgroundservicetimeslicetype
		}
	}

	// Compute reverse map for named struct AircraftGroundServiceType
	// insertion point per field
	stage.AircraftGroundServiceType_TimeSlice_reverseMap = make(map[*AircraftGroundServiceTimeSlicePropertyType]*AircraftGroundServiceType)
	for aircraftgroundservicetype := range stage.AircraftGroundServiceTypes {
		_ = aircraftgroundservicetype
		for _, _aircraftgroundservicetimeslicepropertytype := range aircraftgroundservicetype.TimeSlice {
			stage.AircraftGroundServiceType_TimeSlice_reverseMap[_aircraftgroundservicetimeslicepropertytype] = aircraftgroundservicetype
		}
	}

	// Compute reverse map for named struct AirportClearanceServiceTimeSliceType
	// insertion point per field
	stage.AirportClearanceServiceTimeSliceType_Call_sign_reverseMap = make(map[*CallsignDetailPropertyType]*AirportClearanceServiceTimeSliceType)
	for airportclearanceservicetimeslicetype := range stage.AirportClearanceServiceTimeSliceTypes {
		_ = airportclearanceservicetimeslicetype
		for _, _callsigndetailpropertytype := range airportclearanceservicetimeslicetype.Call_sign {
			stage.AirportClearanceServiceTimeSliceType_Call_sign_reverseMap[_callsigndetailpropertytype] = airportclearanceservicetimeslicetype
		}
	}
	stage.AirportClearanceServiceTimeSliceType_RadioCommunication_reverseMap = make(map[*RadioCommunicationChannelPropertyType]*AirportClearanceServiceTimeSliceType)
	for airportclearanceservicetimeslicetype := range stage.AirportClearanceServiceTimeSliceTypes {
		_ = airportclearanceservicetimeslicetype
		for _, _radiocommunicationchannelpropertytype := range airportclearanceservicetimeslicetype.RadioCommunication {
			stage.AirportClearanceServiceTimeSliceType_RadioCommunication_reverseMap[_radiocommunicationchannelpropertytype] = airportclearanceservicetimeslicetype
		}
	}
	stage.AirportClearanceServiceTimeSliceType_GroundCommunication_reverseMap = make(map[*ContactInformationPropertyType]*AirportClearanceServiceTimeSliceType)
	for airportclearanceservicetimeslicetype := range stage.AirportClearanceServiceTimeSliceTypes {
		_ = airportclearanceservicetimeslicetype
		for _, _contactinformationpropertytype := range airportclearanceservicetimeslicetype.GroundCommunication {
			stage.AirportClearanceServiceTimeSliceType_GroundCommunication_reverseMap[_contactinformationpropertytype] = airportclearanceservicetimeslicetype
		}
	}
	stage.AirportClearanceServiceTimeSliceType_Availability_reverseMap = make(map[*ServiceOperationalStatusPropertyType]*AirportClearanceServiceTimeSliceType)
	for airportclearanceservicetimeslicetype := range stage.AirportClearanceServiceTimeSliceTypes {
		_ = airportclearanceservicetimeslicetype
		for _, _serviceoperationalstatuspropertytype := range airportclearanceservicetimeslicetype.Availability {
			stage.AirportClearanceServiceTimeSliceType_Availability_reverseMap[_serviceoperationalstatuspropertytype] = airportclearanceservicetimeslicetype
		}
	}
	stage.AirportClearanceServiceTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*AirportClearanceServiceTimeSliceType)
	for airportclearanceservicetimeslicetype := range stage.AirportClearanceServiceTimeSliceTypes {
		_ = airportclearanceservicetimeslicetype
		for _, _notepropertytype := range airportclearanceservicetimeslicetype.Annotation {
			stage.AirportClearanceServiceTimeSliceType_Annotation_reverseMap[_notepropertytype] = airportclearanceservicetimeslicetype
		}
	}
	stage.AirportClearanceServiceTimeSliceType_AirportHeliport_reverseMap = make(map[*AirportHeliportPropertyType]*AirportClearanceServiceTimeSliceType)
	for airportclearanceservicetimeslicetype := range stage.AirportClearanceServiceTimeSliceTypes {
		_ = airportclearanceservicetimeslicetype
		for _, _airportheliportpropertytype := range airportclearanceservicetimeslicetype.AirportHeliport {
			stage.AirportClearanceServiceTimeSliceType_AirportHeliport_reverseMap[_airportheliportpropertytype] = airportclearanceservicetimeslicetype
		}
	}
	stage.AirportClearanceServiceTimeSliceType_Extension_reverseMap = make(map[*A_extension_11]*AirportClearanceServiceTimeSliceType)
	for airportclearanceservicetimeslicetype := range stage.AirportClearanceServiceTimeSliceTypes {
		_ = airportclearanceservicetimeslicetype
		for _, _a_extension_11 := range airportclearanceservicetimeslicetype.Extension {
			stage.AirportClearanceServiceTimeSliceType_Extension_reverseMap[_a_extension_11] = airportclearanceservicetimeslicetype
		}
	}

	// Compute reverse map for named struct AirportClearanceServiceType
	// insertion point per field
	stage.AirportClearanceServiceType_TimeSlice_reverseMap = make(map[*AirportClearanceServiceTimeSlicePropertyType]*AirportClearanceServiceType)
	for airportclearanceservicetype := range stage.AirportClearanceServiceTypes {
		_ = airportclearanceservicetype
		for _, _airportclearanceservicetimeslicepropertytype := range airportclearanceservicetype.TimeSlice {
			stage.AirportClearanceServiceType_TimeSlice_reverseMap[_airportclearanceservicetimeslicepropertytype] = airportclearanceservicetype
		}
	}

	// Compute reverse map for named struct AirportSuppliesServiceTimeSliceType
	// insertion point per field
	stage.AirportSuppliesServiceTimeSliceType_Call_sign_reverseMap = make(map[*CallsignDetailPropertyType]*AirportSuppliesServiceTimeSliceType)
	for airportsuppliesservicetimeslicetype := range stage.AirportSuppliesServiceTimeSliceTypes {
		_ = airportsuppliesservicetimeslicetype
		for _, _callsigndetailpropertytype := range airportsuppliesservicetimeslicetype.Call_sign {
			stage.AirportSuppliesServiceTimeSliceType_Call_sign_reverseMap[_callsigndetailpropertytype] = airportsuppliesservicetimeslicetype
		}
	}
	stage.AirportSuppliesServiceTimeSliceType_RadioCommunication_reverseMap = make(map[*RadioCommunicationChannelPropertyType]*AirportSuppliesServiceTimeSliceType)
	for airportsuppliesservicetimeslicetype := range stage.AirportSuppliesServiceTimeSliceTypes {
		_ = airportsuppliesservicetimeslicetype
		for _, _radiocommunicationchannelpropertytype := range airportsuppliesservicetimeslicetype.RadioCommunication {
			stage.AirportSuppliesServiceTimeSliceType_RadioCommunication_reverseMap[_radiocommunicationchannelpropertytype] = airportsuppliesservicetimeslicetype
		}
	}
	stage.AirportSuppliesServiceTimeSliceType_GroundCommunication_reverseMap = make(map[*ContactInformationPropertyType]*AirportSuppliesServiceTimeSliceType)
	for airportsuppliesservicetimeslicetype := range stage.AirportSuppliesServiceTimeSliceTypes {
		_ = airportsuppliesservicetimeslicetype
		for _, _contactinformationpropertytype := range airportsuppliesservicetimeslicetype.GroundCommunication {
			stage.AirportSuppliesServiceTimeSliceType_GroundCommunication_reverseMap[_contactinformationpropertytype] = airportsuppliesservicetimeslicetype
		}
	}
	stage.AirportSuppliesServiceTimeSliceType_Availability_reverseMap = make(map[*ServiceOperationalStatusPropertyType]*AirportSuppliesServiceTimeSliceType)
	for airportsuppliesservicetimeslicetype := range stage.AirportSuppliesServiceTimeSliceTypes {
		_ = airportsuppliesservicetimeslicetype
		for _, _serviceoperationalstatuspropertytype := range airportsuppliesservicetimeslicetype.Availability {
			stage.AirportSuppliesServiceTimeSliceType_Availability_reverseMap[_serviceoperationalstatuspropertytype] = airportsuppliesservicetimeslicetype
		}
	}
	stage.AirportSuppliesServiceTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*AirportSuppliesServiceTimeSliceType)
	for airportsuppliesservicetimeslicetype := range stage.AirportSuppliesServiceTimeSliceTypes {
		_ = airportsuppliesservicetimeslicetype
		for _, _notepropertytype := range airportsuppliesservicetimeslicetype.Annotation {
			stage.AirportSuppliesServiceTimeSliceType_Annotation_reverseMap[_notepropertytype] = airportsuppliesservicetimeslicetype
		}
	}
	stage.AirportSuppliesServiceTimeSliceType_AirportHeliport_reverseMap = make(map[*AirportHeliportPropertyType]*AirportSuppliesServiceTimeSliceType)
	for airportsuppliesservicetimeslicetype := range stage.AirportSuppliesServiceTimeSliceTypes {
		_ = airportsuppliesservicetimeslicetype
		for _, _airportheliportpropertytype := range airportsuppliesservicetimeslicetype.AirportHeliport {
			stage.AirportSuppliesServiceTimeSliceType_AirportHeliport_reverseMap[_airportheliportpropertytype] = airportsuppliesservicetimeslicetype
		}
	}
	stage.AirportSuppliesServiceTimeSliceType_FuelSupply_reverseMap = make(map[*FuelPropertyType]*AirportSuppliesServiceTimeSliceType)
	for airportsuppliesservicetimeslicetype := range stage.AirportSuppliesServiceTimeSliceTypes {
		_ = airportsuppliesservicetimeslicetype
		for _, _fuelpropertytype := range airportsuppliesservicetimeslicetype.FuelSupply {
			stage.AirportSuppliesServiceTimeSliceType_FuelSupply_reverseMap[_fuelpropertytype] = airportsuppliesservicetimeslicetype
		}
	}
	stage.AirportSuppliesServiceTimeSliceType_OilSupply_reverseMap = make(map[*OilPropertyType]*AirportSuppliesServiceTimeSliceType)
	for airportsuppliesservicetimeslicetype := range stage.AirportSuppliesServiceTimeSliceTypes {
		_ = airportsuppliesservicetimeslicetype
		for _, _oilpropertytype := range airportsuppliesservicetimeslicetype.OilSupply {
			stage.AirportSuppliesServiceTimeSliceType_OilSupply_reverseMap[_oilpropertytype] = airportsuppliesservicetimeslicetype
		}
	}
	stage.AirportSuppliesServiceTimeSliceType_NitrogenSupply_reverseMap = make(map[*NitrogenPropertyType]*AirportSuppliesServiceTimeSliceType)
	for airportsuppliesservicetimeslicetype := range stage.AirportSuppliesServiceTimeSliceTypes {
		_ = airportsuppliesservicetimeslicetype
		for _, _nitrogenpropertytype := range airportsuppliesservicetimeslicetype.NitrogenSupply {
			stage.AirportSuppliesServiceTimeSliceType_NitrogenSupply_reverseMap[_nitrogenpropertytype] = airportsuppliesservicetimeslicetype
		}
	}
	stage.AirportSuppliesServiceTimeSliceType_OxygenSupply_reverseMap = make(map[*OxygenPropertyType]*AirportSuppliesServiceTimeSliceType)
	for airportsuppliesservicetimeslicetype := range stage.AirportSuppliesServiceTimeSliceTypes {
		_ = airportsuppliesservicetimeslicetype
		for _, _oxygenpropertytype := range airportsuppliesservicetimeslicetype.OxygenSupply {
			stage.AirportSuppliesServiceTimeSliceType_OxygenSupply_reverseMap[_oxygenpropertytype] = airportsuppliesservicetimeslicetype
		}
	}
	stage.AirportSuppliesServiceTimeSliceType_Extension_reverseMap = make(map[*A_extension_12]*AirportSuppliesServiceTimeSliceType)
	for airportsuppliesservicetimeslicetype := range stage.AirportSuppliesServiceTimeSliceTypes {
		_ = airportsuppliesservicetimeslicetype
		for _, _a_extension_12 := range airportsuppliesservicetimeslicetype.Extension {
			stage.AirportSuppliesServiceTimeSliceType_Extension_reverseMap[_a_extension_12] = airportsuppliesservicetimeslicetype
		}
	}

	// Compute reverse map for named struct AirportSuppliesServiceType
	// insertion point per field
	stage.AirportSuppliesServiceType_TimeSlice_reverseMap = make(map[*AirportSuppliesServiceTimeSlicePropertyType]*AirportSuppliesServiceType)
	for airportsuppliesservicetype := range stage.AirportSuppliesServiceTypes {
		_ = airportsuppliesservicetype
		for _, _airportsuppliesservicetimeslicepropertytype := range airportsuppliesservicetype.TimeSlice {
			stage.AirportSuppliesServiceType_TimeSlice_reverseMap[_airportsuppliesservicetimeslicepropertytype] = airportsuppliesservicetype
		}
	}

	// Compute reverse map for named struct CallsignDetailType
	// insertion point per field
	stage.CallsignDetailType_Annotation_reverseMap = make(map[*NotePropertyType]*CallsignDetailType)
	for callsigndetailtype := range stage.CallsignDetailTypes {
		_ = callsigndetailtype
		for _, _notepropertytype := range callsigndetailtype.Annotation {
			stage.CallsignDetailType_Annotation_reverseMap[_notepropertytype] = callsigndetailtype
		}
	}
	stage.CallsignDetailType_Extension_reverseMap = make(map[*A_extension_15]*CallsignDetailType)
	for callsigndetailtype := range stage.CallsignDetailTypes {
		_ = callsigndetailtype
		for _, _a_extension_15 := range callsigndetailtype.Extension {
			stage.CallsignDetailType_Extension_reverseMap[_a_extension_15] = callsigndetailtype
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
	stage.ContactInformationType_Extension_reverseMap = make(map[*A_extension_29]*ContactInformationType)
	for contactinformationtype := range stage.ContactInformationTypes {
		_ = contactinformationtype
		for _, _a_extension_29 := range contactinformationtype.Extension {
			stage.ContactInformationType_Extension_reverseMap[_a_extension_29] = contactinformationtype
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
	stage.ElevatedPointType_Extension_reverseMap = make(map[*A_extension_1]*ElevatedPointType)
	for elevatedpointtype := range stage.ElevatedPointTypes {
		_ = elevatedpointtype
		for _, _a_extension_1 := range elevatedpointtype.Extension {
			stage.ElevatedPointType_Extension_reverseMap[_a_extension_1] = elevatedpointtype
		}
	}

	// Compute reverse map for named struct FireFightingServiceTimeSliceType
	// insertion point per field
	stage.FireFightingServiceTimeSliceType_Call_sign_reverseMap = make(map[*CallsignDetailPropertyType]*FireFightingServiceTimeSliceType)
	for firefightingservicetimeslicetype := range stage.FireFightingServiceTimeSliceTypes {
		_ = firefightingservicetimeslicetype
		for _, _callsigndetailpropertytype := range firefightingservicetimeslicetype.Call_sign {
			stage.FireFightingServiceTimeSliceType_Call_sign_reverseMap[_callsigndetailpropertytype] = firefightingservicetimeslicetype
		}
	}
	stage.FireFightingServiceTimeSliceType_RadioCommunication_reverseMap = make(map[*RadioCommunicationChannelPropertyType]*FireFightingServiceTimeSliceType)
	for firefightingservicetimeslicetype := range stage.FireFightingServiceTimeSliceTypes {
		_ = firefightingservicetimeslicetype
		for _, _radiocommunicationchannelpropertytype := range firefightingservicetimeslicetype.RadioCommunication {
			stage.FireFightingServiceTimeSliceType_RadioCommunication_reverseMap[_radiocommunicationchannelpropertytype] = firefightingservicetimeslicetype
		}
	}
	stage.FireFightingServiceTimeSliceType_GroundCommunication_reverseMap = make(map[*ContactInformationPropertyType]*FireFightingServiceTimeSliceType)
	for firefightingservicetimeslicetype := range stage.FireFightingServiceTimeSliceTypes {
		_ = firefightingservicetimeslicetype
		for _, _contactinformationpropertytype := range firefightingservicetimeslicetype.GroundCommunication {
			stage.FireFightingServiceTimeSliceType_GroundCommunication_reverseMap[_contactinformationpropertytype] = firefightingservicetimeslicetype
		}
	}
	stage.FireFightingServiceTimeSliceType_Availability_reverseMap = make(map[*ServiceOperationalStatusPropertyType]*FireFightingServiceTimeSliceType)
	for firefightingservicetimeslicetype := range stage.FireFightingServiceTimeSliceTypes {
		_ = firefightingservicetimeslicetype
		for _, _serviceoperationalstatuspropertytype := range firefightingservicetimeslicetype.Availability {
			stage.FireFightingServiceTimeSliceType_Availability_reverseMap[_serviceoperationalstatuspropertytype] = firefightingservicetimeslicetype
		}
	}
	stage.FireFightingServiceTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*FireFightingServiceTimeSliceType)
	for firefightingservicetimeslicetype := range stage.FireFightingServiceTimeSliceTypes {
		_ = firefightingservicetimeslicetype
		for _, _notepropertytype := range firefightingservicetimeslicetype.Annotation {
			stage.FireFightingServiceTimeSliceType_Annotation_reverseMap[_notepropertytype] = firefightingservicetimeslicetype
		}
	}
	stage.FireFightingServiceTimeSliceType_AirportHeliport_reverseMap = make(map[*AirportHeliportPropertyType]*FireFightingServiceTimeSliceType)
	for firefightingservicetimeslicetype := range stage.FireFightingServiceTimeSliceTypes {
		_ = firefightingservicetimeslicetype
		for _, _airportheliportpropertytype := range firefightingservicetimeslicetype.AirportHeliport {
			stage.FireFightingServiceTimeSliceType_AirportHeliport_reverseMap[_airportheliportpropertytype] = firefightingservicetimeslicetype
		}
	}
	stage.FireFightingServiceTimeSliceType_Extension_reverseMap = make(map[*A_extension_16]*FireFightingServiceTimeSliceType)
	for firefightingservicetimeslicetype := range stage.FireFightingServiceTimeSliceTypes {
		_ = firefightingservicetimeslicetype
		for _, _a_extension_16 := range firefightingservicetimeslicetype.Extension {
			stage.FireFightingServiceTimeSliceType_Extension_reverseMap[_a_extension_16] = firefightingservicetimeslicetype
		}
	}

	// Compute reverse map for named struct FireFightingServiceType
	// insertion point per field
	stage.FireFightingServiceType_TimeSlice_reverseMap = make(map[*FireFightingServiceTimeSlicePropertyType]*FireFightingServiceType)
	for firefightingservicetype := range stage.FireFightingServiceTypes {
		_ = firefightingservicetype
		for _, _firefightingservicetimeslicepropertytype := range firefightingservicetype.TimeSlice {
			stage.FireFightingServiceType_TimeSlice_reverseMap[_firefightingservicetimeslicepropertytype] = firefightingservicetype
		}
	}

	// Compute reverse map for named struct FuelType
	// insertion point per field
	stage.FuelType_Annotation_reverseMap = make(map[*NotePropertyType]*FuelType)
	for fueltype := range stage.FuelTypes {
		_ = fueltype
		for _, _notepropertytype := range fueltype.Annotation {
			stage.FuelType_Annotation_reverseMap[_notepropertytype] = fueltype
		}
	}
	stage.FuelType_Extension_reverseMap = make(map[*A_extension_17]*FuelType)
	for fueltype := range stage.FuelTypes {
		_ = fueltype
		for _, _a_extension_17 := range fueltype.Extension {
			stage.FuelType_Extension_reverseMap[_a_extension_17] = fueltype
		}
	}

	// Compute reverse map for named struct GroundTrafficControlServiceTimeSliceType
	// insertion point per field
	stage.GroundTrafficControlServiceTimeSliceType_Call_sign_reverseMap = make(map[*CallsignDetailPropertyType]*GroundTrafficControlServiceTimeSliceType)
	for groundtrafficcontrolservicetimeslicetype := range stage.GroundTrafficControlServiceTimeSliceTypes {
		_ = groundtrafficcontrolservicetimeslicetype
		for _, _callsigndetailpropertytype := range groundtrafficcontrolservicetimeslicetype.Call_sign {
			stage.GroundTrafficControlServiceTimeSliceType_Call_sign_reverseMap[_callsigndetailpropertytype] = groundtrafficcontrolservicetimeslicetype
		}
	}
	stage.GroundTrafficControlServiceTimeSliceType_RadioCommunication_reverseMap = make(map[*RadioCommunicationChannelPropertyType]*GroundTrafficControlServiceTimeSliceType)
	for groundtrafficcontrolservicetimeslicetype := range stage.GroundTrafficControlServiceTimeSliceTypes {
		_ = groundtrafficcontrolservicetimeslicetype
		for _, _radiocommunicationchannelpropertytype := range groundtrafficcontrolservicetimeslicetype.RadioCommunication {
			stage.GroundTrafficControlServiceTimeSliceType_RadioCommunication_reverseMap[_radiocommunicationchannelpropertytype] = groundtrafficcontrolservicetimeslicetype
		}
	}
	stage.GroundTrafficControlServiceTimeSliceType_GroundCommunication_reverseMap = make(map[*ContactInformationPropertyType]*GroundTrafficControlServiceTimeSliceType)
	for groundtrafficcontrolservicetimeslicetype := range stage.GroundTrafficControlServiceTimeSliceTypes {
		_ = groundtrafficcontrolservicetimeslicetype
		for _, _contactinformationpropertytype := range groundtrafficcontrolservicetimeslicetype.GroundCommunication {
			stage.GroundTrafficControlServiceTimeSliceType_GroundCommunication_reverseMap[_contactinformationpropertytype] = groundtrafficcontrolservicetimeslicetype
		}
	}
	stage.GroundTrafficControlServiceTimeSliceType_Availability_reverseMap = make(map[*ServiceOperationalStatusPropertyType]*GroundTrafficControlServiceTimeSliceType)
	for groundtrafficcontrolservicetimeslicetype := range stage.GroundTrafficControlServiceTimeSliceTypes {
		_ = groundtrafficcontrolservicetimeslicetype
		for _, _serviceoperationalstatuspropertytype := range groundtrafficcontrolservicetimeslicetype.Availability {
			stage.GroundTrafficControlServiceTimeSliceType_Availability_reverseMap[_serviceoperationalstatuspropertytype] = groundtrafficcontrolservicetimeslicetype
		}
	}
	stage.GroundTrafficControlServiceTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*GroundTrafficControlServiceTimeSliceType)
	for groundtrafficcontrolservicetimeslicetype := range stage.GroundTrafficControlServiceTimeSliceTypes {
		_ = groundtrafficcontrolservicetimeslicetype
		for _, _notepropertytype := range groundtrafficcontrolservicetimeslicetype.Annotation {
			stage.GroundTrafficControlServiceTimeSliceType_Annotation_reverseMap[_notepropertytype] = groundtrafficcontrolservicetimeslicetype
		}
	}
	stage.GroundTrafficControlServiceTimeSliceType_Extension_reverseMap = make(map[*A_extension_18]*GroundTrafficControlServiceTimeSliceType)
	for groundtrafficcontrolservicetimeslicetype := range stage.GroundTrafficControlServiceTimeSliceTypes {
		_ = groundtrafficcontrolservicetimeslicetype
		for _, _a_extension_18 := range groundtrafficcontrolservicetimeslicetype.Extension {
			stage.GroundTrafficControlServiceTimeSliceType_Extension_reverseMap[_a_extension_18] = groundtrafficcontrolservicetimeslicetype
		}
	}

	// Compute reverse map for named struct GroundTrafficControlServiceType
	// insertion point per field
	stage.GroundTrafficControlServiceType_TimeSlice_reverseMap = make(map[*GroundTrafficControlServiceTimeSlicePropertyType]*GroundTrafficControlServiceType)
	for groundtrafficcontrolservicetype := range stage.GroundTrafficControlServiceTypes {
		_ = groundtrafficcontrolservicetype
		for _, _groundtrafficcontrolservicetimeslicepropertytype := range groundtrafficcontrolservicetype.TimeSlice {
			stage.GroundTrafficControlServiceType_TimeSlice_reverseMap[_groundtrafficcontrolservicetimeslicepropertytype] = groundtrafficcontrolservicetype
		}
	}

	// Compute reverse map for named struct InformationServiceTimeSliceType
	// insertion point per field
	stage.InformationServiceTimeSliceType_Call_sign_reverseMap = make(map[*CallsignDetailPropertyType]*InformationServiceTimeSliceType)
	for informationservicetimeslicetype := range stage.InformationServiceTimeSliceTypes {
		_ = informationservicetimeslicetype
		for _, _callsigndetailpropertytype := range informationservicetimeslicetype.Call_sign {
			stage.InformationServiceTimeSliceType_Call_sign_reverseMap[_callsigndetailpropertytype] = informationservicetimeslicetype
		}
	}
	stage.InformationServiceTimeSliceType_RadioCommunication_reverseMap = make(map[*RadioCommunicationChannelPropertyType]*InformationServiceTimeSliceType)
	for informationservicetimeslicetype := range stage.InformationServiceTimeSliceTypes {
		_ = informationservicetimeslicetype
		for _, _radiocommunicationchannelpropertytype := range informationservicetimeslicetype.RadioCommunication {
			stage.InformationServiceTimeSliceType_RadioCommunication_reverseMap[_radiocommunicationchannelpropertytype] = informationservicetimeslicetype
		}
	}
	stage.InformationServiceTimeSliceType_GroundCommunication_reverseMap = make(map[*ContactInformationPropertyType]*InformationServiceTimeSliceType)
	for informationservicetimeslicetype := range stage.InformationServiceTimeSliceTypes {
		_ = informationservicetimeslicetype
		for _, _contactinformationpropertytype := range informationservicetimeslicetype.GroundCommunication {
			stage.InformationServiceTimeSliceType_GroundCommunication_reverseMap[_contactinformationpropertytype] = informationservicetimeslicetype
		}
	}
	stage.InformationServiceTimeSliceType_Availability_reverseMap = make(map[*ServiceOperationalStatusPropertyType]*InformationServiceTimeSliceType)
	for informationservicetimeslicetype := range stage.InformationServiceTimeSliceTypes {
		_ = informationservicetimeslicetype
		for _, _serviceoperationalstatuspropertytype := range informationservicetimeslicetype.Availability {
			stage.InformationServiceTimeSliceType_Availability_reverseMap[_serviceoperationalstatuspropertytype] = informationservicetimeslicetype
		}
	}
	stage.InformationServiceTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*InformationServiceTimeSliceType)
	for informationservicetimeslicetype := range stage.InformationServiceTimeSliceTypes {
		_ = informationservicetimeslicetype
		for _, _notepropertytype := range informationservicetimeslicetype.Annotation {
			stage.InformationServiceTimeSliceType_Annotation_reverseMap[_notepropertytype] = informationservicetimeslicetype
		}
	}
	stage.InformationServiceTimeSliceType_NavaidBroadcast_reverseMap = make(map[*VORPropertyType]*InformationServiceTimeSliceType)
	for informationservicetimeslicetype := range stage.InformationServiceTimeSliceTypes {
		_ = informationservicetimeslicetype
		for _, _vorpropertytype := range informationservicetimeslicetype.NavaidBroadcast {
			stage.InformationServiceTimeSliceType_NavaidBroadcast_reverseMap[_vorpropertytype] = informationservicetimeslicetype
		}
	}
	stage.InformationServiceTimeSliceType_ClientAirspace_reverseMap = make(map[*AirspacePropertyType]*InformationServiceTimeSliceType)
	for informationservicetimeslicetype := range stage.InformationServiceTimeSliceTypes {
		_ = informationservicetimeslicetype
		for _, _airspacepropertytype := range informationservicetimeslicetype.ClientAirspace {
			stage.InformationServiceTimeSliceType_ClientAirspace_reverseMap[_airspacepropertytype] = informationservicetimeslicetype
		}
	}
	stage.InformationServiceTimeSliceType_ClientAirport_reverseMap = make(map[*AirportHeliportPropertyType]*InformationServiceTimeSliceType)
	for informationservicetimeslicetype := range stage.InformationServiceTimeSliceTypes {
		_ = informationservicetimeslicetype
		for _, _airportheliportpropertytype := range informationservicetimeslicetype.ClientAirport {
			stage.InformationServiceTimeSliceType_ClientAirport_reverseMap[_airportheliportpropertytype] = informationservicetimeslicetype
		}
	}
	stage.InformationServiceTimeSliceType_ClientRoute_reverseMap = make(map[*RoutePortionPropertyType]*InformationServiceTimeSliceType)
	for informationservicetimeslicetype := range stage.InformationServiceTimeSliceTypes {
		_ = informationservicetimeslicetype
		for _, _routeportionpropertytype := range informationservicetimeslicetype.ClientRoute {
			stage.InformationServiceTimeSliceType_ClientRoute_reverseMap[_routeportionpropertytype] = informationservicetimeslicetype
		}
	}
	stage.InformationServiceTimeSliceType_ClientProcedure_reverseMap = make(map[*ProcedurePropertyType]*InformationServiceTimeSliceType)
	for informationservicetimeslicetype := range stage.InformationServiceTimeSliceTypes {
		_ = informationservicetimeslicetype
		for _, _procedurepropertytype := range informationservicetimeslicetype.ClientProcedure {
			stage.InformationServiceTimeSliceType_ClientProcedure_reverseMap[_procedurepropertytype] = informationservicetimeslicetype
		}
	}
	stage.InformationServiceTimeSliceType_ClientHolding_reverseMap = make(map[*HoldingPatternPropertyType]*InformationServiceTimeSliceType)
	for informationservicetimeslicetype := range stage.InformationServiceTimeSliceTypes {
		_ = informationservicetimeslicetype
		for _, _holdingpatternpropertytype := range informationservicetimeslicetype.ClientHolding {
			stage.InformationServiceTimeSliceType_ClientHolding_reverseMap[_holdingpatternpropertytype] = informationservicetimeslicetype
		}
	}
	stage.InformationServiceTimeSliceType_ClientAerialRefuelling_reverseMap = make(map[*AerialRefuellingPropertyType]*InformationServiceTimeSliceType)
	for informationservicetimeslicetype := range stage.InformationServiceTimeSliceTypes {
		_ = informationservicetimeslicetype
		for _, _aerialrefuellingpropertytype := range informationservicetimeslicetype.ClientAerialRefuelling {
			stage.InformationServiceTimeSliceType_ClientAerialRefuelling_reverseMap[_aerialrefuellingpropertytype] = informationservicetimeslicetype
		}
	}
	stage.InformationServiceTimeSliceType_Extension_reverseMap = make(map[*A_extension_19]*InformationServiceTimeSliceType)
	for informationservicetimeslicetype := range stage.InformationServiceTimeSliceTypes {
		_ = informationservicetimeslicetype
		for _, _a_extension_19 := range informationservicetimeslicetype.Extension {
			stage.InformationServiceTimeSliceType_Extension_reverseMap[_a_extension_19] = informationservicetimeslicetype
		}
	}

	// Compute reverse map for named struct InformationServiceType
	// insertion point per field
	stage.InformationServiceType_TimeSlice_reverseMap = make(map[*InformationServiceTimeSlicePropertyType]*InformationServiceType)
	for informationservicetype := range stage.InformationServiceTypes {
		_ = informationservicetype
		for _, _informationservicetimeslicepropertytype := range informationservicetype.TimeSlice {
			stage.InformationServiceType_TimeSlice_reverseMap[_informationservicetimeslicepropertytype] = informationservicetype
		}
	}

	// Compute reverse map for named struct LightActivationType
	// insertion point per field
	stage.LightActivationType_Annotation_reverseMap = make(map[*NotePropertyType]*LightActivationType)
	for lightactivationtype := range stage.LightActivationTypes {
		_ = lightactivationtype
		for _, _notepropertytype := range lightactivationtype.Annotation {
			stage.LightActivationType_Annotation_reverseMap[_notepropertytype] = lightactivationtype
		}
	}
	stage.LightActivationType_Extension_reverseMap = make(map[*A_extension]*LightActivationType)
	for lightactivationtype := range stage.LightActivationTypes {
		_ = lightactivationtype
		for _, _a_extension := range lightactivationtype.Extension {
			stage.LightActivationType_Extension_reverseMap[_a_extension] = lightactivationtype
		}
	}

	// Compute reverse map for named struct LinguisticNoteType
	// insertion point per field
	stage.LinguisticNoteType_Extension_reverseMap = make(map[*A_extension_2]*LinguisticNoteType)
	for linguisticnotetype := range stage.LinguisticNoteTypes {
		_ = linguisticnotetype
		for _, _a_extension_2 := range linguisticnotetype.Extension {
			stage.LinguisticNoteType_Extension_reverseMap[_a_extension_2] = linguisticnotetype
		}
	}

	// Compute reverse map for named struct NitrogenType
	// insertion point per field
	stage.NitrogenType_Annotation_reverseMap = make(map[*NotePropertyType]*NitrogenType)
	for nitrogentype := range stage.NitrogenTypes {
		_ = nitrogentype
		for _, _notepropertytype := range nitrogentype.Annotation {
			stage.NitrogenType_Annotation_reverseMap[_notepropertytype] = nitrogentype
		}
	}
	stage.NitrogenType_Extension_reverseMap = make(map[*A_extension_20]*NitrogenType)
	for nitrogentype := range stage.NitrogenTypes {
		_ = nitrogentype
		for _, _a_extension_20 := range nitrogentype.Extension {
			stage.NitrogenType_Extension_reverseMap[_a_extension_20] = nitrogentype
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
	stage.NoteType_Extension_reverseMap = make(map[*A_extension_3]*NoteType)
	for notetype := range stage.NoteTypes {
		_ = notetype
		for _, _a_extension_3 := range notetype.Extension {
			stage.NoteType_Extension_reverseMap[_a_extension_3] = notetype
		}
	}

	// Compute reverse map for named struct OilType
	// insertion point per field
	stage.OilType_Annotation_reverseMap = make(map[*NotePropertyType]*OilType)
	for oiltype := range stage.OilTypes {
		_ = oiltype
		for _, _notepropertytype := range oiltype.Annotation {
			stage.OilType_Annotation_reverseMap[_notepropertytype] = oiltype
		}
	}
	stage.OilType_Extension_reverseMap = make(map[*A_extension_21]*OilType)
	for oiltype := range stage.OilTypes {
		_ = oiltype
		for _, _a_extension_21 := range oiltype.Extension {
			stage.OilType_Extension_reverseMap[_a_extension_21] = oiltype
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
	stage.OnlineContactType_Extension_reverseMap = make(map[*A_extension_30]*OnlineContactType)
	for onlinecontacttype := range stage.OnlineContactTypes {
		_ = onlinecontacttype
		for _, _a_extension_30 := range onlinecontacttype.Extension {
			stage.OnlineContactType_Extension_reverseMap[_a_extension_30] = onlinecontacttype
		}
	}

	// Compute reverse map for named struct OrganisationAuthorityAssociationType
	// insertion point per field
	stage.OrganisationAuthorityAssociationType_Annotation_reverseMap = make(map[*NotePropertyType]*OrganisationAuthorityAssociationType)
	for organisationauthorityassociationtype := range stage.OrganisationAuthorityAssociationTypes {
		_ = organisationauthorityassociationtype
		for _, _notepropertytype := range organisationauthorityassociationtype.Annotation {
			stage.OrganisationAuthorityAssociationType_Annotation_reverseMap[_notepropertytype] = organisationauthorityassociationtype
		}
	}
	stage.OrganisationAuthorityAssociationType_Extension_reverseMap = make(map[*A_extension_5]*OrganisationAuthorityAssociationType)
	for organisationauthorityassociationtype := range stage.OrganisationAuthorityAssociationTypes {
		_ = organisationauthorityassociationtype
		for _, _a_extension_5 := range organisationauthorityassociationtype.Extension {
			stage.OrganisationAuthorityAssociationType_Extension_reverseMap[_a_extension_5] = organisationauthorityassociationtype
		}
	}

	// Compute reverse map for named struct OrganisationAuthorityTimeSliceType
	// insertion point per field
	stage.OrganisationAuthorityTimeSliceType_Contact_reverseMap = make(map[*ContactInformationPropertyType]*OrganisationAuthorityTimeSliceType)
	for organisationauthoritytimeslicetype := range stage.OrganisationAuthorityTimeSliceTypes {
		_ = organisationauthoritytimeslicetype
		for _, _contactinformationpropertytype := range organisationauthoritytimeslicetype.Contact {
			stage.OrganisationAuthorityTimeSliceType_Contact_reverseMap[_contactinformationpropertytype] = organisationauthoritytimeslicetype
		}
	}
	stage.OrganisationAuthorityTimeSliceType_RelatedOrganisationAuthority_reverseMap = make(map[*OrganisationAuthorityAssociationPropertyType]*OrganisationAuthorityTimeSliceType)
	for organisationauthoritytimeslicetype := range stage.OrganisationAuthorityTimeSliceTypes {
		_ = organisationauthoritytimeslicetype
		for _, _organisationauthorityassociationpropertytype := range organisationauthoritytimeslicetype.RelatedOrganisationAuthority {
			stage.OrganisationAuthorityTimeSliceType_RelatedOrganisationAuthority_reverseMap[_organisationauthorityassociationpropertytype] = organisationauthoritytimeslicetype
		}
	}
	stage.OrganisationAuthorityTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*OrganisationAuthorityTimeSliceType)
	for organisationauthoritytimeslicetype := range stage.OrganisationAuthorityTimeSliceTypes {
		_ = organisationauthoritytimeslicetype
		for _, _notepropertytype := range organisationauthoritytimeslicetype.Annotation {
			stage.OrganisationAuthorityTimeSliceType_Annotation_reverseMap[_notepropertytype] = organisationauthoritytimeslicetype
		}
	}
	stage.OrganisationAuthorityTimeSliceType_Extension_reverseMap = make(map[*A_extension_4]*OrganisationAuthorityTimeSliceType)
	for organisationauthoritytimeslicetype := range stage.OrganisationAuthorityTimeSliceTypes {
		_ = organisationauthoritytimeslicetype
		for _, _a_extension_4 := range organisationauthoritytimeslicetype.Extension {
			stage.OrganisationAuthorityTimeSliceType_Extension_reverseMap[_a_extension_4] = organisationauthoritytimeslicetype
		}
	}

	// Compute reverse map for named struct OrganisationAuthorityType
	// insertion point per field
	stage.OrganisationAuthorityType_TimeSlice_reverseMap = make(map[*OrganisationAuthorityTimeSlicePropertyType]*OrganisationAuthorityType)
	for organisationauthoritytype := range stage.OrganisationAuthorityTypes {
		_ = organisationauthoritytype
		for _, _organisationauthoritytimeslicepropertytype := range organisationauthoritytype.TimeSlice {
			stage.OrganisationAuthorityType_TimeSlice_reverseMap[_organisationauthoritytimeslicepropertytype] = organisationauthoritytype
		}
	}

	// Compute reverse map for named struct OxygenType
	// insertion point per field
	stage.OxygenType_Annotation_reverseMap = make(map[*NotePropertyType]*OxygenType)
	for oxygentype := range stage.OxygenTypes {
		_ = oxygentype
		for _, _notepropertytype := range oxygentype.Annotation {
			stage.OxygenType_Annotation_reverseMap[_notepropertytype] = oxygentype
		}
	}
	stage.OxygenType_Extension_reverseMap = make(map[*A_extension_22]*OxygenType)
	for oxygentype := range stage.OxygenTypes {
		_ = oxygentype
		for _, _a_extension_22 := range oxygentype.Extension {
			stage.OxygenType_Extension_reverseMap[_a_extension_22] = oxygentype
		}
	}

	// Compute reverse map for named struct PassengerServiceTimeSliceType
	// insertion point per field
	stage.PassengerServiceTimeSliceType_Call_sign_reverseMap = make(map[*CallsignDetailPropertyType]*PassengerServiceTimeSliceType)
	for passengerservicetimeslicetype := range stage.PassengerServiceTimeSliceTypes {
		_ = passengerservicetimeslicetype
		for _, _callsigndetailpropertytype := range passengerservicetimeslicetype.Call_sign {
			stage.PassengerServiceTimeSliceType_Call_sign_reverseMap[_callsigndetailpropertytype] = passengerservicetimeslicetype
		}
	}
	stage.PassengerServiceTimeSliceType_RadioCommunication_reverseMap = make(map[*RadioCommunicationChannelPropertyType]*PassengerServiceTimeSliceType)
	for passengerservicetimeslicetype := range stage.PassengerServiceTimeSliceTypes {
		_ = passengerservicetimeslicetype
		for _, _radiocommunicationchannelpropertytype := range passengerservicetimeslicetype.RadioCommunication {
			stage.PassengerServiceTimeSliceType_RadioCommunication_reverseMap[_radiocommunicationchannelpropertytype] = passengerservicetimeslicetype
		}
	}
	stage.PassengerServiceTimeSliceType_GroundCommunication_reverseMap = make(map[*ContactInformationPropertyType]*PassengerServiceTimeSliceType)
	for passengerservicetimeslicetype := range stage.PassengerServiceTimeSliceTypes {
		_ = passengerservicetimeslicetype
		for _, _contactinformationpropertytype := range passengerservicetimeslicetype.GroundCommunication {
			stage.PassengerServiceTimeSliceType_GroundCommunication_reverseMap[_contactinformationpropertytype] = passengerservicetimeslicetype
		}
	}
	stage.PassengerServiceTimeSliceType_Availability_reverseMap = make(map[*ServiceOperationalStatusPropertyType]*PassengerServiceTimeSliceType)
	for passengerservicetimeslicetype := range stage.PassengerServiceTimeSliceTypes {
		_ = passengerservicetimeslicetype
		for _, _serviceoperationalstatuspropertytype := range passengerservicetimeslicetype.Availability {
			stage.PassengerServiceTimeSliceType_Availability_reverseMap[_serviceoperationalstatuspropertytype] = passengerservicetimeslicetype
		}
	}
	stage.PassengerServiceTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*PassengerServiceTimeSliceType)
	for passengerservicetimeslicetype := range stage.PassengerServiceTimeSliceTypes {
		_ = passengerservicetimeslicetype
		for _, _notepropertytype := range passengerservicetimeslicetype.Annotation {
			stage.PassengerServiceTimeSliceType_Annotation_reverseMap[_notepropertytype] = passengerservicetimeslicetype
		}
	}
	stage.PassengerServiceTimeSliceType_AirportHeliport_reverseMap = make(map[*AirportHeliportPropertyType]*PassengerServiceTimeSliceType)
	for passengerservicetimeslicetype := range stage.PassengerServiceTimeSliceTypes {
		_ = passengerservicetimeslicetype
		for _, _airportheliportpropertytype := range passengerservicetimeslicetype.AirportHeliport {
			stage.PassengerServiceTimeSliceType_AirportHeliport_reverseMap[_airportheliportpropertytype] = passengerservicetimeslicetype
		}
	}
	stage.PassengerServiceTimeSliceType_Extension_reverseMap = make(map[*A_extension_23]*PassengerServiceTimeSliceType)
	for passengerservicetimeslicetype := range stage.PassengerServiceTimeSliceTypes {
		_ = passengerservicetimeslicetype
		for _, _a_extension_23 := range passengerservicetimeslicetype.Extension {
			stage.PassengerServiceTimeSliceType_Extension_reverseMap[_a_extension_23] = passengerservicetimeslicetype
		}
	}

	// Compute reverse map for named struct PassengerServiceType
	// insertion point per field
	stage.PassengerServiceType_TimeSlice_reverseMap = make(map[*PassengerServiceTimeSlicePropertyType]*PassengerServiceType)
	for passengerservicetype := range stage.PassengerServiceTypes {
		_ = passengerservicetype
		for _, _passengerservicetimeslicepropertytype := range passengerservicetype.TimeSlice {
			stage.PassengerServiceType_TimeSlice_reverseMap[_passengerservicetimeslicepropertytype] = passengerservicetype
		}
	}

	// Compute reverse map for named struct PilotControlledLightingTimeSliceType
	// insertion point per field
	stage.PilotControlledLightingTimeSliceType_ControlledLightIntensity_reverseMap = make(map[*LightActivationPropertyType]*PilotControlledLightingTimeSliceType)
	for pilotcontrolledlightingtimeslicetype := range stage.PilotControlledLightingTimeSliceTypes {
		_ = pilotcontrolledlightingtimeslicetype
		for _, _lightactivationpropertytype := range pilotcontrolledlightingtimeslicetype.ControlledLightIntensity {
			stage.PilotControlledLightingTimeSliceType_ControlledLightIntensity_reverseMap[_lightactivationpropertytype] = pilotcontrolledlightingtimeslicetype
		}
	}
	stage.PilotControlledLightingTimeSliceType_ActivatedGroundLighting_reverseMap = make(map[*GroundLightSystemPropertyType]*PilotControlledLightingTimeSliceType)
	for pilotcontrolledlightingtimeslicetype := range stage.PilotControlledLightingTimeSliceTypes {
		_ = pilotcontrolledlightingtimeslicetype
		for _, _groundlightsystempropertytype := range pilotcontrolledlightingtimeslicetype.ActivatedGroundLighting {
			stage.PilotControlledLightingTimeSliceType_ActivatedGroundLighting_reverseMap[_groundlightsystempropertytype] = pilotcontrolledlightingtimeslicetype
		}
	}
	stage.PilotControlledLightingTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*PilotControlledLightingTimeSliceType)
	for pilotcontrolledlightingtimeslicetype := range stage.PilotControlledLightingTimeSliceTypes {
		_ = pilotcontrolledlightingtimeslicetype
		for _, _notepropertytype := range pilotcontrolledlightingtimeslicetype.Annotation {
			stage.PilotControlledLightingTimeSliceType_Annotation_reverseMap[_notepropertytype] = pilotcontrolledlightingtimeslicetype
		}
	}
	stage.PilotControlledLightingTimeSliceType_Extension_reverseMap = make(map[*A_extension_24]*PilotControlledLightingTimeSliceType)
	for pilotcontrolledlightingtimeslicetype := range stage.PilotControlledLightingTimeSliceTypes {
		_ = pilotcontrolledlightingtimeslicetype
		for _, _a_extension_24 := range pilotcontrolledlightingtimeslicetype.Extension {
			stage.PilotControlledLightingTimeSliceType_Extension_reverseMap[_a_extension_24] = pilotcontrolledlightingtimeslicetype
		}
	}

	// Compute reverse map for named struct PilotControlledLightingType
	// insertion point per field
	stage.PilotControlledLightingType_TimeSlice_reverseMap = make(map[*PilotControlledLightingTimeSlicePropertyType]*PilotControlledLightingType)
	for pilotcontrolledlightingtype := range stage.PilotControlledLightingTypes {
		_ = pilotcontrolledlightingtype
		for _, _pilotcontrolledlightingtimeslicepropertytype := range pilotcontrolledlightingtype.TimeSlice {
			stage.PilotControlledLightingType_TimeSlice_reverseMap[_pilotcontrolledlightingtimeslicepropertytype] = pilotcontrolledlightingtype
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
	stage.PostalAddressType_Extension_reverseMap = make(map[*A_extension_31]*PostalAddressType)
	for postaladdresstype := range stage.PostalAddressTypes {
		_ = postaladdresstype
		for _, _a_extension_31 := range postaladdresstype.Extension {
			stage.PostalAddressType_Extension_reverseMap[_a_extension_31] = postaladdresstype
		}
	}

	// Compute reverse map for named struct RadioCommunicationChannelTimeSliceType
	// insertion point per field
	stage.RadioCommunicationChannelTimeSliceType_Location_reverseMap = make(map[*ElevatedPointPropertyType]*RadioCommunicationChannelTimeSliceType)
	for radiocommunicationchanneltimeslicetype := range stage.RadioCommunicationChannelTimeSliceTypes {
		_ = radiocommunicationchanneltimeslicetype
		for _, _elevatedpointpropertytype := range radiocommunicationchanneltimeslicetype.Location {
			stage.RadioCommunicationChannelTimeSliceType_Location_reverseMap[_elevatedpointpropertytype] = radiocommunicationchanneltimeslicetype
		}
	}
	stage.RadioCommunicationChannelTimeSliceType_Availability_reverseMap = make(map[*RadioCommunicationOperationalStatusPropertyType]*RadioCommunicationChannelTimeSliceType)
	for radiocommunicationchanneltimeslicetype := range stage.RadioCommunicationChannelTimeSliceTypes {
		_ = radiocommunicationchanneltimeslicetype
		for _, _radiocommunicationoperationalstatuspropertytype := range radiocommunicationchanneltimeslicetype.Availability {
			stage.RadioCommunicationChannelTimeSliceType_Availability_reverseMap[_radiocommunicationoperationalstatuspropertytype] = radiocommunicationchanneltimeslicetype
		}
	}
	stage.RadioCommunicationChannelTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*RadioCommunicationChannelTimeSliceType)
	for radiocommunicationchanneltimeslicetype := range stage.RadioCommunicationChannelTimeSliceTypes {
		_ = radiocommunicationchanneltimeslicetype
		for _, _notepropertytype := range radiocommunicationchanneltimeslicetype.Annotation {
			stage.RadioCommunicationChannelTimeSliceType_Annotation_reverseMap[_notepropertytype] = radiocommunicationchanneltimeslicetype
		}
	}
	stage.RadioCommunicationChannelTimeSliceType_Extension_reverseMap = make(map[*A_extension_25]*RadioCommunicationChannelTimeSliceType)
	for radiocommunicationchanneltimeslicetype := range stage.RadioCommunicationChannelTimeSliceTypes {
		_ = radiocommunicationchanneltimeslicetype
		for _, _a_extension_25 := range radiocommunicationchanneltimeslicetype.Extension {
			stage.RadioCommunicationChannelTimeSliceType_Extension_reverseMap[_a_extension_25] = radiocommunicationchanneltimeslicetype
		}
	}

	// Compute reverse map for named struct RadioCommunicationChannelType
	// insertion point per field
	stage.RadioCommunicationChannelType_TimeSlice_reverseMap = make(map[*RadioCommunicationChannelTimeSlicePropertyType]*RadioCommunicationChannelType)
	for radiocommunicationchanneltype := range stage.RadioCommunicationChannelTypes {
		_ = radiocommunicationchanneltype
		for _, _radiocommunicationchanneltimeslicepropertytype := range radiocommunicationchanneltype.TimeSlice {
			stage.RadioCommunicationChannelType_TimeSlice_reverseMap[_radiocommunicationchanneltimeslicepropertytype] = radiocommunicationchanneltype
		}
	}

	// Compute reverse map for named struct RadioCommunicationOperationalStatusType
	// insertion point per field
	stage.RadioCommunicationOperationalStatusType_TimeInterval_reverseMap = make(map[*TimesheetPropertyType]*RadioCommunicationOperationalStatusType)
	for radiocommunicationoperationalstatustype := range stage.RadioCommunicationOperationalStatusTypes {
		_ = radiocommunicationoperationalstatustype
		for _, _timesheetpropertytype := range radiocommunicationoperationalstatustype.TimeInterval {
			stage.RadioCommunicationOperationalStatusType_TimeInterval_reverseMap[_timesheetpropertytype] = radiocommunicationoperationalstatustype
		}
	}
	stage.RadioCommunicationOperationalStatusType_Annotation_reverseMap = make(map[*NotePropertyType]*RadioCommunicationOperationalStatusType)
	for radiocommunicationoperationalstatustype := range stage.RadioCommunicationOperationalStatusTypes {
		_ = radiocommunicationoperationalstatustype
		for _, _notepropertytype := range radiocommunicationoperationalstatustype.Annotation {
			stage.RadioCommunicationOperationalStatusType_Annotation_reverseMap[_notepropertytype] = radiocommunicationoperationalstatustype
		}
	}
	stage.RadioCommunicationOperationalStatusType_SpecialDateAuthority_reverseMap = make(map[*OrganisationAuthorityPropertyType]*RadioCommunicationOperationalStatusType)
	for radiocommunicationoperationalstatustype := range stage.RadioCommunicationOperationalStatusTypes {
		_ = radiocommunicationoperationalstatustype
		for _, _organisationauthoritypropertytype := range radiocommunicationoperationalstatustype.SpecialDateAuthority {
			stage.RadioCommunicationOperationalStatusType_SpecialDateAuthority_reverseMap[_organisationauthoritypropertytype] = radiocommunicationoperationalstatustype
		}
	}
	stage.RadioCommunicationOperationalStatusType_Extension_reverseMap = make(map[*A_extension_26]*RadioCommunicationOperationalStatusType)
	for radiocommunicationoperationalstatustype := range stage.RadioCommunicationOperationalStatusTypes {
		_ = radiocommunicationoperationalstatustype
		for _, _a_extension_26 := range radiocommunicationoperationalstatustype.Extension {
			stage.RadioCommunicationOperationalStatusType_Extension_reverseMap[_a_extension_26] = radiocommunicationoperationalstatustype
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
	stage.RoutePortionType_Extension_reverseMap = make(map[*A_extension_9]*RoutePortionType)
	for routeportiontype := range stage.RoutePortionTypes {
		_ = routeportiontype
		for _, _a_extension_9 := range routeportiontype.Extension {
			stage.RoutePortionType_Extension_reverseMap[_a_extension_9] = routeportiontype
		}
	}

	// Compute reverse map for named struct SearchRescueServiceTimeSliceType
	// insertion point per field
	stage.SearchRescueServiceTimeSliceType_Call_sign_reverseMap = make(map[*CallsignDetailPropertyType]*SearchRescueServiceTimeSliceType)
	for searchrescueservicetimeslicetype := range stage.SearchRescueServiceTimeSliceTypes {
		_ = searchrescueservicetimeslicetype
		for _, _callsigndetailpropertytype := range searchrescueservicetimeslicetype.Call_sign {
			stage.SearchRescueServiceTimeSliceType_Call_sign_reverseMap[_callsigndetailpropertytype] = searchrescueservicetimeslicetype
		}
	}
	stage.SearchRescueServiceTimeSliceType_RadioCommunication_reverseMap = make(map[*RadioCommunicationChannelPropertyType]*SearchRescueServiceTimeSliceType)
	for searchrescueservicetimeslicetype := range stage.SearchRescueServiceTimeSliceTypes {
		_ = searchrescueservicetimeslicetype
		for _, _radiocommunicationchannelpropertytype := range searchrescueservicetimeslicetype.RadioCommunication {
			stage.SearchRescueServiceTimeSliceType_RadioCommunication_reverseMap[_radiocommunicationchannelpropertytype] = searchrescueservicetimeslicetype
		}
	}
	stage.SearchRescueServiceTimeSliceType_GroundCommunication_reverseMap = make(map[*ContactInformationPropertyType]*SearchRescueServiceTimeSliceType)
	for searchrescueservicetimeslicetype := range stage.SearchRescueServiceTimeSliceTypes {
		_ = searchrescueservicetimeslicetype
		for _, _contactinformationpropertytype := range searchrescueservicetimeslicetype.GroundCommunication {
			stage.SearchRescueServiceTimeSliceType_GroundCommunication_reverseMap[_contactinformationpropertytype] = searchrescueservicetimeslicetype
		}
	}
	stage.SearchRescueServiceTimeSliceType_Availability_reverseMap = make(map[*ServiceOperationalStatusPropertyType]*SearchRescueServiceTimeSliceType)
	for searchrescueservicetimeslicetype := range stage.SearchRescueServiceTimeSliceTypes {
		_ = searchrescueservicetimeslicetype
		for _, _serviceoperationalstatuspropertytype := range searchrescueservicetimeslicetype.Availability {
			stage.SearchRescueServiceTimeSliceType_Availability_reverseMap[_serviceoperationalstatuspropertytype] = searchrescueservicetimeslicetype
		}
	}
	stage.SearchRescueServiceTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*SearchRescueServiceTimeSliceType)
	for searchrescueservicetimeslicetype := range stage.SearchRescueServiceTimeSliceTypes {
		_ = searchrescueservicetimeslicetype
		for _, _notepropertytype := range searchrescueservicetimeslicetype.Annotation {
			stage.SearchRescueServiceTimeSliceType_Annotation_reverseMap[_notepropertytype] = searchrescueservicetimeslicetype
		}
	}
	stage.SearchRescueServiceTimeSliceType_ClientAirspace_reverseMap = make(map[*AirspacePropertyType]*SearchRescueServiceTimeSliceType)
	for searchrescueservicetimeslicetype := range stage.SearchRescueServiceTimeSliceTypes {
		_ = searchrescueservicetimeslicetype
		for _, _airspacepropertytype := range searchrescueservicetimeslicetype.ClientAirspace {
			stage.SearchRescueServiceTimeSliceType_ClientAirspace_reverseMap[_airspacepropertytype] = searchrescueservicetimeslicetype
		}
	}
	stage.SearchRescueServiceTimeSliceType_ClientRoute_reverseMap = make(map[*RoutePortionPropertyType]*SearchRescueServiceTimeSliceType)
	for searchrescueservicetimeslicetype := range stage.SearchRescueServiceTimeSliceTypes {
		_ = searchrescueservicetimeslicetype
		for _, _routeportionpropertytype := range searchrescueservicetimeslicetype.ClientRoute {
			stage.SearchRescueServiceTimeSliceType_ClientRoute_reverseMap[_routeportionpropertytype] = searchrescueservicetimeslicetype
		}
	}
	stage.SearchRescueServiceTimeSliceType_Extension_reverseMap = make(map[*A_extension_27]*SearchRescueServiceTimeSliceType)
	for searchrescueservicetimeslicetype := range stage.SearchRescueServiceTimeSliceTypes {
		_ = searchrescueservicetimeslicetype
		for _, _a_extension_27 := range searchrescueservicetimeslicetype.Extension {
			stage.SearchRescueServiceTimeSliceType_Extension_reverseMap[_a_extension_27] = searchrescueservicetimeslicetype
		}
	}

	// Compute reverse map for named struct SearchRescueServiceType
	// insertion point per field
	stage.SearchRescueServiceType_TimeSlice_reverseMap = make(map[*SearchRescueServiceTimeSlicePropertyType]*SearchRescueServiceType)
	for searchrescueservicetype := range stage.SearchRescueServiceTypes {
		_ = searchrescueservicetype
		for _, _searchrescueservicetimeslicepropertytype := range searchrescueservicetype.TimeSlice {
			stage.SearchRescueServiceType_TimeSlice_reverseMap[_searchrescueservicetimeslicepropertytype] = searchrescueservicetype
		}
	}

	// Compute reverse map for named struct ServiceOperationalStatusType
	// insertion point per field
	stage.ServiceOperationalStatusType_TimeInterval_reverseMap = make(map[*TimesheetPropertyType]*ServiceOperationalStatusType)
	for serviceoperationalstatustype := range stage.ServiceOperationalStatusTypes {
		_ = serviceoperationalstatustype
		for _, _timesheetpropertytype := range serviceoperationalstatustype.TimeInterval {
			stage.ServiceOperationalStatusType_TimeInterval_reverseMap[_timesheetpropertytype] = serviceoperationalstatustype
		}
	}
	stage.ServiceOperationalStatusType_Annotation_reverseMap = make(map[*NotePropertyType]*ServiceOperationalStatusType)
	for serviceoperationalstatustype := range stage.ServiceOperationalStatusTypes {
		_ = serviceoperationalstatustype
		for _, _notepropertytype := range serviceoperationalstatustype.Annotation {
			stage.ServiceOperationalStatusType_Annotation_reverseMap[_notepropertytype] = serviceoperationalstatustype
		}
	}
	stage.ServiceOperationalStatusType_SpecialDateAuthority_reverseMap = make(map[*OrganisationAuthorityPropertyType]*ServiceOperationalStatusType)
	for serviceoperationalstatustype := range stage.ServiceOperationalStatusTypes {
		_ = serviceoperationalstatustype
		for _, _organisationauthoritypropertytype := range serviceoperationalstatustype.SpecialDateAuthority {
			stage.ServiceOperationalStatusType_SpecialDateAuthority_reverseMap[_organisationauthoritypropertytype] = serviceoperationalstatustype
		}
	}
	stage.ServiceOperationalStatusType_Extension_reverseMap = make(map[*A_extension_28]*ServiceOperationalStatusType)
	for serviceoperationalstatustype := range stage.ServiceOperationalStatusTypes {
		_ = serviceoperationalstatustype
		for _, _a_extension_28 := range serviceoperationalstatustype.Extension {
			stage.ServiceOperationalStatusType_Extension_reverseMap[_a_extension_28] = serviceoperationalstatustype
		}
	}

	// Compute reverse map for named struct SpecialDateTimeSliceType
	// insertion point per field
	stage.SpecialDateTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*SpecialDateTimeSliceType)
	for specialdatetimeslicetype := range stage.SpecialDateTimeSliceTypes {
		_ = specialdatetimeslicetype
		for _, _notepropertytype := range specialdatetimeslicetype.Annotation {
			stage.SpecialDateTimeSliceType_Annotation_reverseMap[_notepropertytype] = specialdatetimeslicetype
		}
	}
	stage.SpecialDateTimeSliceType_Extension_reverseMap = make(map[*A_extension_33]*SpecialDateTimeSliceType)
	for specialdatetimeslicetype := range stage.SpecialDateTimeSliceTypes {
		_ = specialdatetimeslicetype
		for _, _a_extension_33 := range specialdatetimeslicetype.Extension {
			stage.SpecialDateTimeSliceType_Extension_reverseMap[_a_extension_33] = specialdatetimeslicetype
		}
	}

	// Compute reverse map for named struct SpecialDateType
	// insertion point per field
	stage.SpecialDateType_TimeSlice_reverseMap = make(map[*SpecialDateTimeSlicePropertyType]*SpecialDateType)
	for specialdatetype := range stage.SpecialDateTypes {
		_ = specialdatetype
		for _, _specialdatetimeslicepropertytype := range specialdatetype.TimeSlice {
			stage.SpecialDateType_TimeSlice_reverseMap[_specialdatetimeslicepropertytype] = specialdatetype
		}
	}

	// Compute reverse map for named struct StandardLevelColumnTimeSliceType
	// insertion point per field
	stage.StandardLevelColumnTimeSliceType_Level_reverseMap = make(map[*StandardLevelPropertyType]*StandardLevelColumnTimeSliceType)
	for standardlevelcolumntimeslicetype := range stage.StandardLevelColumnTimeSliceTypes {
		_ = standardlevelcolumntimeslicetype
		for _, _standardlevelpropertytype := range standardlevelcolumntimeslicetype.Level {
			stage.StandardLevelColumnTimeSliceType_Level_reverseMap[_standardlevelpropertytype] = standardlevelcolumntimeslicetype
		}
	}
	stage.StandardLevelColumnTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*StandardLevelColumnTimeSliceType)
	for standardlevelcolumntimeslicetype := range stage.StandardLevelColumnTimeSliceTypes {
		_ = standardlevelcolumntimeslicetype
		for _, _notepropertytype := range standardlevelcolumntimeslicetype.Annotation {
			stage.StandardLevelColumnTimeSliceType_Annotation_reverseMap[_notepropertytype] = standardlevelcolumntimeslicetype
		}
	}
	stage.StandardLevelColumnTimeSliceType_Extension_reverseMap = make(map[*A_extension_36]*StandardLevelColumnTimeSliceType)
	for standardlevelcolumntimeslicetype := range stage.StandardLevelColumnTimeSliceTypes {
		_ = standardlevelcolumntimeslicetype
		for _, _a_extension_36 := range standardlevelcolumntimeslicetype.Extension {
			stage.StandardLevelColumnTimeSliceType_Extension_reverseMap[_a_extension_36] = standardlevelcolumntimeslicetype
		}
	}

	// Compute reverse map for named struct StandardLevelColumnType
	// insertion point per field
	stage.StandardLevelColumnType_TimeSlice_reverseMap = make(map[*StandardLevelColumnTimeSlicePropertyType]*StandardLevelColumnType)
	for standardlevelcolumntype := range stage.StandardLevelColumnTypes {
		_ = standardlevelcolumntype
		for _, _standardlevelcolumntimeslicepropertytype := range standardlevelcolumntype.TimeSlice {
			stage.StandardLevelColumnType_TimeSlice_reverseMap[_standardlevelcolumntimeslicepropertytype] = standardlevelcolumntype
		}
	}

	// Compute reverse map for named struct StandardLevelSectorTimeSliceType
	// insertion point per field
	stage.StandardLevelSectorTimeSliceType_ApplicableAirspace_reverseMap = make(map[*AirspacePropertyType]*StandardLevelSectorTimeSliceType)
	for standardlevelsectortimeslicetype := range stage.StandardLevelSectorTimeSliceTypes {
		_ = standardlevelsectortimeslicetype
		for _, _airspacepropertytype := range standardlevelsectortimeslicetype.ApplicableAirspace {
			stage.StandardLevelSectorTimeSliceType_ApplicableAirspace_reverseMap[_airspacepropertytype] = standardlevelsectortimeslicetype
		}
	}
	stage.StandardLevelSectorTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*StandardLevelSectorTimeSliceType)
	for standardlevelsectortimeslicetype := range stage.StandardLevelSectorTimeSliceTypes {
		_ = standardlevelsectortimeslicetype
		for _, _notepropertytype := range standardlevelsectortimeslicetype.Annotation {
			stage.StandardLevelSectorTimeSliceType_Annotation_reverseMap[_notepropertytype] = standardlevelsectortimeslicetype
		}
	}
	stage.StandardLevelSectorTimeSliceType_Extension_reverseMap = make(map[*A_extension_37]*StandardLevelSectorTimeSliceType)
	for standardlevelsectortimeslicetype := range stage.StandardLevelSectorTimeSliceTypes {
		_ = standardlevelsectortimeslicetype
		for _, _a_extension_37 := range standardlevelsectortimeslicetype.Extension {
			stage.StandardLevelSectorTimeSliceType_Extension_reverseMap[_a_extension_37] = standardlevelsectortimeslicetype
		}
	}

	// Compute reverse map for named struct StandardLevelSectorType
	// insertion point per field
	stage.StandardLevelSectorType_TimeSlice_reverseMap = make(map[*StandardLevelSectorTimeSlicePropertyType]*StandardLevelSectorType)
	for standardlevelsectortype := range stage.StandardLevelSectorTypes {
		_ = standardlevelsectortype
		for _, _standardlevelsectortimeslicepropertytype := range standardlevelsectortype.TimeSlice {
			stage.StandardLevelSectorType_TimeSlice_reverseMap[_standardlevelsectortimeslicepropertytype] = standardlevelsectortype
		}
	}

	// Compute reverse map for named struct StandardLevelTableTimeSliceType
	// insertion point per field
	stage.StandardLevelTableTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*StandardLevelTableTimeSliceType)
	for standardleveltabletimeslicetype := range stage.StandardLevelTableTimeSliceTypes {
		_ = standardleveltabletimeslicetype
		for _, _notepropertytype := range standardleveltabletimeslicetype.Annotation {
			stage.StandardLevelTableTimeSliceType_Annotation_reverseMap[_notepropertytype] = standardleveltabletimeslicetype
		}
	}
	stage.StandardLevelTableTimeSliceType_Extension_reverseMap = make(map[*A_extension_38]*StandardLevelTableTimeSliceType)
	for standardleveltabletimeslicetype := range stage.StandardLevelTableTimeSliceTypes {
		_ = standardleveltabletimeslicetype
		for _, _a_extension_38 := range standardleveltabletimeslicetype.Extension {
			stage.StandardLevelTableTimeSliceType_Extension_reverseMap[_a_extension_38] = standardleveltabletimeslicetype
		}
	}

	// Compute reverse map for named struct StandardLevelTableType
	// insertion point per field
	stage.StandardLevelTableType_TimeSlice_reverseMap = make(map[*StandardLevelTableTimeSlicePropertyType]*StandardLevelTableType)
	for standardleveltabletype := range stage.StandardLevelTableTypes {
		_ = standardleveltabletype
		for _, _standardleveltabletimeslicepropertytype := range standardleveltabletype.TimeSlice {
			stage.StandardLevelTableType_TimeSlice_reverseMap[_standardleveltabletimeslicepropertytype] = standardleveltabletype
		}
	}

	// Compute reverse map for named struct StandardLevelType
	// insertion point per field
	stage.StandardLevelType_Annotation_reverseMap = make(map[*NotePropertyType]*StandardLevelType)
	for standardleveltype := range stage.StandardLevelTypes {
		_ = standardleveltype
		for _, _notepropertytype := range standardleveltype.Annotation {
			stage.StandardLevelType_Annotation_reverseMap[_notepropertytype] = standardleveltype
		}
	}
	stage.StandardLevelType_Extension_reverseMap = make(map[*A_extension_35]*StandardLevelType)
	for standardleveltype := range stage.StandardLevelTypes {
		_ = standardleveltype
		for _, _a_extension_35 := range standardleveltype.Extension {
			stage.StandardLevelType_Extension_reverseMap[_a_extension_35] = standardleveltype
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
	stage.TelephoneContactType_Extension_reverseMap = make(map[*A_extension_32]*TelephoneContactType)
	for telephonecontacttype := range stage.TelephoneContactTypes {
		_ = telephonecontacttype
		for _, _a_extension_32 := range telephonecontacttype.Extension {
			stage.TelephoneContactType_Extension_reverseMap[_a_extension_32] = telephonecontacttype
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

	// Compute reverse map for named struct UnitAvailabilityType
	// insertion point per field
	stage.UnitAvailabilityType_TimeInterval_reverseMap = make(map[*TimesheetPropertyType]*UnitAvailabilityType)
	for unitavailabilitytype := range stage.UnitAvailabilityTypes {
		_ = unitavailabilitytype
		for _, _timesheetpropertytype := range unitavailabilitytype.TimeInterval {
			stage.UnitAvailabilityType_TimeInterval_reverseMap[_timesheetpropertytype] = unitavailabilitytype
		}
	}
	stage.UnitAvailabilityType_Annotation_reverseMap = make(map[*NotePropertyType]*UnitAvailabilityType)
	for unitavailabilitytype := range stage.UnitAvailabilityTypes {
		_ = unitavailabilitytype
		for _, _notepropertytype := range unitavailabilitytype.Annotation {
			stage.UnitAvailabilityType_Annotation_reverseMap[_notepropertytype] = unitavailabilitytype
		}
	}
	stage.UnitAvailabilityType_SpecialDateAuthority_reverseMap = make(map[*OrganisationAuthorityPropertyType]*UnitAvailabilityType)
	for unitavailabilitytype := range stage.UnitAvailabilityTypes {
		_ = unitavailabilitytype
		for _, _organisationauthoritypropertytype := range unitavailabilitytype.SpecialDateAuthority {
			stage.UnitAvailabilityType_SpecialDateAuthority_reverseMap[_organisationauthoritypropertytype] = unitavailabilitytype
		}
	}
	stage.UnitAvailabilityType_Extension_reverseMap = make(map[*A_extension_7]*UnitAvailabilityType)
	for unitavailabilitytype := range stage.UnitAvailabilityTypes {
		_ = unitavailabilitytype
		for _, _a_extension_7 := range unitavailabilitytype.Extension {
			stage.UnitAvailabilityType_Extension_reverseMap[_a_extension_7] = unitavailabilitytype
		}
	}

	// Compute reverse map for named struct UnitDependencyType
	// insertion point per field
	stage.UnitDependencyType_Annotation_reverseMap = make(map[*NotePropertyType]*UnitDependencyType)
	for unitdependencytype := range stage.UnitDependencyTypes {
		_ = unitdependencytype
		for _, _notepropertytype := range unitdependencytype.Annotation {
			stage.UnitDependencyType_Annotation_reverseMap[_notepropertytype] = unitdependencytype
		}
	}
	stage.UnitDependencyType_Extension_reverseMap = make(map[*A_extension_8]*UnitDependencyType)
	for unitdependencytype := range stage.UnitDependencyTypes {
		_ = unitdependencytype
		for _, _a_extension_8 := range unitdependencytype.Extension {
			stage.UnitDependencyType_Extension_reverseMap[_a_extension_8] = unitdependencytype
		}
	}

	// Compute reverse map for named struct UnitTimeSliceType
	// insertion point per field
	stage.UnitTimeSliceType_Contact_reverseMap = make(map[*ContactInformationPropertyType]*UnitTimeSliceType)
	for unittimeslicetype := range stage.UnitTimeSliceTypes {
		_ = unittimeslicetype
		for _, _contactinformationpropertytype := range unittimeslicetype.Contact {
			stage.UnitTimeSliceType_Contact_reverseMap[_contactinformationpropertytype] = unittimeslicetype
		}
	}
	stage.UnitTimeSliceType_RelatedUnit_reverseMap = make(map[*UnitDependencyPropertyType]*UnitTimeSliceType)
	for unittimeslicetype := range stage.UnitTimeSliceTypes {
		_ = unittimeslicetype
		for _, _unitdependencypropertytype := range unittimeslicetype.RelatedUnit {
			stage.UnitTimeSliceType_RelatedUnit_reverseMap[_unitdependencypropertytype] = unittimeslicetype
		}
	}
	stage.UnitTimeSliceType_Availability_reverseMap = make(map[*UnitAvailabilityPropertyType]*UnitTimeSliceType)
	for unittimeslicetype := range stage.UnitTimeSliceTypes {
		_ = unittimeslicetype
		for _, _unitavailabilitypropertytype := range unittimeslicetype.Availability {
			stage.UnitTimeSliceType_Availability_reverseMap[_unitavailabilitypropertytype] = unittimeslicetype
		}
	}
	stage.UnitTimeSliceType_Annotation_reverseMap = make(map[*NotePropertyType]*UnitTimeSliceType)
	for unittimeslicetype := range stage.UnitTimeSliceTypes {
		_ = unittimeslicetype
		for _, _notepropertytype := range unittimeslicetype.Annotation {
			stage.UnitTimeSliceType_Annotation_reverseMap[_notepropertytype] = unittimeslicetype
		}
	}
	stage.UnitTimeSliceType_Extension_reverseMap = make(map[*A_extension_6]*UnitTimeSliceType)
	for unittimeslicetype := range stage.UnitTimeSliceTypes {
		_ = unittimeslicetype
		for _, _a_extension_6 := range unittimeslicetype.Extension {
			stage.UnitTimeSliceType_Extension_reverseMap[_a_extension_6] = unittimeslicetype
		}
	}

	// Compute reverse map for named struct UnitType
	// insertion point per field
	stage.UnitType_TimeSlice_reverseMap = make(map[*UnitTimeSlicePropertyType]*UnitType)
	for unittype := range stage.UnitTypes {
		_ = unittype
		for _, _unittimeslicepropertytype := range unittype.TimeSlice {
			stage.UnitType_TimeSlice_reverseMap[_unittimeslicepropertytype] = unittype
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

	res = __gong__appendInstances(res, stage.A_extension_4s)

	res = __gong__appendInstances(res, stage.A_extension_5s)

	res = __gong__appendInstances(res, stage.A_extension_6s)

	res = __gong__appendInstances(res, stage.A_extension_7s)

	res = __gong__appendInstances(res, stage.A_extension_8s)

	res = __gong__appendInstances(res, stage.A_extension_9s)

	res = __gong__appendInstances(res, stage.AbstractAirportGroundServiceTypes)

	res = __gong__appendInstances(res, stage.AbstractPropertiesWithScheduleTypes)

	res = __gong__appendInstances(res, stage.AbstractServiceTypes)

	res = __gong__appendInstances(res, stage.AbstractTrafficSeparationServiceTypes)

	res = __gong__appendInstances(res, stage.AerialRefuellingPropertyTypes)

	res = __gong__appendInstances(res, stage.AirTrafficControlServiceTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.AirTrafficControlServiceTimeSliceTypes)

	res = __gong__appendInstances(res, stage.AirTrafficControlServiceTypes)

	res = __gong__appendInstances(res, stage.AirTrafficManagementServiceTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.AirTrafficManagementServiceTimeSliceTypes)

	res = __gong__appendInstances(res, stage.AirTrafficManagementServiceTypes)

	res = __gong__appendInstances(res, stage.AircraftGroundServiceTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.AircraftGroundServiceTimeSliceTypes)

	res = __gong__appendInstances(res, stage.AircraftGroundServiceTypes)

	res = __gong__appendInstances(res, stage.AirportClearanceServiceTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.AirportClearanceServiceTimeSliceTypes)

	res = __gong__appendInstances(res, stage.AirportClearanceServiceTypes)

	res = __gong__appendInstances(res, stage.AirportHeliportPropertyTypes)

	res = __gong__appendInstances(res, stage.AirportSuppliesServiceTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.AirportSuppliesServiceTimeSliceTypes)

	res = __gong__appendInstances(res, stage.AirportSuppliesServiceTypes)

	res = __gong__appendInstances(res, stage.AirspacePropertyTypes)

	res = __gong__appendInstances(res, stage.CallsignDetailPropertyTypes)

	res = __gong__appendInstances(res, stage.CallsignDetailTypes)

	res = __gong__appendInstances(res, stage.ContactInformationPropertyTypes)

	res = __gong__appendInstances(res, stage.ContactInformationTypes)

	res = __gong__appendInstances(res, stage.DesignatedPointPropertyTypes)

	res = __gong__appendInstances(res, stage.DirectionFinderPropertyTypes)

	res = __gong__appendInstances(res, stage.ElevatedPointPropertyTypes)

	res = __gong__appendInstances(res, stage.ElevatedPointTypes)

	res = __gong__appendInstances(res, stage.FireFightingServiceTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.FireFightingServiceTimeSliceTypes)

	res = __gong__appendInstances(res, stage.FireFightingServiceTypes)

	res = __gong__appendInstances(res, stage.FuelPropertyTypes)

	res = __gong__appendInstances(res, stage.FuelTypes)

	res = __gong__appendInstances(res, stage.GroundLightSystemPropertyTypes)

	res = __gong__appendInstances(res, stage.GroundTrafficControlServiceTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.GroundTrafficControlServiceTimeSliceTypes)

	res = __gong__appendInstances(res, stage.GroundTrafficControlServiceTypes)

	res = __gong__appendInstances(res, stage.HoldingPatternPropertyTypes)

	res = __gong__appendInstances(res, stage.InformationServiceTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.InformationServiceTimeSliceTypes)

	res = __gong__appendInstances(res, stage.InformationServiceTypes)

	res = __gong__appendInstances(res, stage.LightActivationPropertyTypes)

	res = __gong__appendInstances(res, stage.LightActivationTypes)

	res = __gong__appendInstances(res, stage.LinguisticNotePropertyTypes)

	res = __gong__appendInstances(res, stage.LinguisticNoteTypes)

	res = __gong__appendInstances(res, stage.NavaidPropertyTypes)

	res = __gong__appendInstances(res, stage.NitrogenPropertyTypes)

	res = __gong__appendInstances(res, stage.NitrogenTypes)

	res = __gong__appendInstances(res, stage.NotePropertyTypes)

	res = __gong__appendInstances(res, stage.NoteTypes)

	res = __gong__appendInstances(res, stage.OilPropertyTypes)

	res = __gong__appendInstances(res, stage.OilTypes)

	res = __gong__appendInstances(res, stage.OnlineContactPropertyTypes)

	res = __gong__appendInstances(res, stage.OnlineContactTypes)

	res = __gong__appendInstances(res, stage.OrganisationAuthorityAssociationPropertyTypes)

	res = __gong__appendInstances(res, stage.OrganisationAuthorityAssociationTypes)

	res = __gong__appendInstances(res, stage.OrganisationAuthorityPropertyTypes)

	res = __gong__appendInstances(res, stage.OrganisationAuthorityTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.OrganisationAuthorityTimeSliceTypes)

	res = __gong__appendInstances(res, stage.OrganisationAuthorityTypes)

	res = __gong__appendInstances(res, stage.OxygenPropertyTypes)

	res = __gong__appendInstances(res, stage.OxygenTypes)

	res = __gong__appendInstances(res, stage.PassengerServiceTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.PassengerServiceTimeSliceTypes)

	res = __gong__appendInstances(res, stage.PassengerServiceTypes)

	res = __gong__appendInstances(res, stage.PilotControlledLightingTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.PilotControlledLightingTimeSliceTypes)

	res = __gong__appendInstances(res, stage.PilotControlledLightingTypes)

	res = __gong__appendInstances(res, stage.PointPropertyTypes)

	res = __gong__appendInstances(res, stage.PointTypes)

	res = __gong__appendInstances(res, stage.PostalAddressPropertyTypes)

	res = __gong__appendInstances(res, stage.PostalAddressTypes)

	res = __gong__appendInstances(res, stage.ProcedurePropertyTypes)

	res = __gong__appendInstances(res, stage.RadioCommunicationChannelPropertyTypes)

	res = __gong__appendInstances(res, stage.RadioCommunicationChannelTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.RadioCommunicationChannelTimeSliceTypes)

	res = __gong__appendInstances(res, stage.RadioCommunicationChannelTypes)

	res = __gong__appendInstances(res, stage.RadioCommunicationOperationalStatusPropertyTypes)

	res = __gong__appendInstances(res, stage.RadioCommunicationOperationalStatusTypes)

	res = __gong__appendInstances(res, stage.RoutePortionPropertyTypes)

	res = __gong__appendInstances(res, stage.RoutePortionTypes)

	res = __gong__appendInstances(res, stage.RoutePropertyTypes)

	res = __gong__appendInstances(res, stage.RunwayCentrelinePointPropertyTypes)

	res = __gong__appendInstances(res, stage.SearchRescueServiceTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.SearchRescueServiceTimeSliceTypes)

	res = __gong__appendInstances(res, stage.SearchRescueServiceTypes)

	res = __gong__appendInstances(res, stage.ServiceOperationalStatusPropertyTypes)

	res = __gong__appendInstances(res, stage.ServiceOperationalStatusTypes)

	res = __gong__appendInstances(res, stage.SpecialDateTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.SpecialDateTimeSliceTypes)

	res = __gong__appendInstances(res, stage.SpecialDateTypes)

	res = __gong__appendInstances(res, stage.StandardLevelColumnPropertyTypes)

	res = __gong__appendInstances(res, stage.StandardLevelColumnTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.StandardLevelColumnTimeSliceTypes)

	res = __gong__appendInstances(res, stage.StandardLevelColumnTypes)

	res = __gong__appendInstances(res, stage.StandardLevelPropertyTypes)

	res = __gong__appendInstances(res, stage.StandardLevelSectorTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.StandardLevelSectorTimeSliceTypes)

	res = __gong__appendInstances(res, stage.StandardLevelSectorTypes)

	res = __gong__appendInstances(res, stage.StandardLevelTablePropertyTypes)

	res = __gong__appendInstances(res, stage.StandardLevelTableTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.StandardLevelTableTimeSliceTypes)

	res = __gong__appendInstances(res, stage.StandardLevelTableTypes)

	res = __gong__appendInstances(res, stage.StandardLevelTypes)

	res = __gong__appendInstances(res, stage.TelephoneContactPropertyTypes)

	res = __gong__appendInstances(res, stage.TelephoneContactTypes)

	res = __gong__appendInstances(res, stage.TimesheetPropertyTypes)

	res = __gong__appendInstances(res, stage.TimesheetTypes)

	res = __gong__appendInstances(res, stage.TouchDownLiftOffPropertyTypes)

	res = __gong__appendInstances(res, stage.UnitAvailabilityPropertyTypes)

	res = __gong__appendInstances(res, stage.UnitAvailabilityTypes)

	res = __gong__appendInstances(res, stage.UnitDependencyPropertyTypes)

	res = __gong__appendInstances(res, stage.UnitDependencyTypes)

	res = __gong__appendInstances(res, stage.UnitPropertyTypes)

	res = __gong__appendInstances(res, stage.UnitTimeSlicePropertyTypes)

	res = __gong__appendInstances(res, stage.UnitTimeSliceTypes)

	res = __gong__appendInstances(res, stage.UnitTypes)

	res = __gong__appendInstances(res, stage.VORPropertyTypes)

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

func (abstractairportgroundservicetype *AbstractAirportGroundServiceType) GongCopy() GongstructIF {
	newInstance := new(AbstractAirportGroundServiceType)
	abstractairportgroundservicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongCopy() GongstructIF {
	newInstance := new(AbstractPropertiesWithScheduleType)
	abstractpropertieswithscheduletype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstractservicetype *AbstractServiceType) GongCopy() GongstructIF {
	newInstance := new(AbstractServiceType)
	abstractservicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstracttrafficseparationservicetype *AbstractTrafficSeparationServiceType) GongCopy() GongstructIF {
	newInstance := new(AbstractTrafficSeparationServiceType)
	abstracttrafficseparationservicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (aerialrefuellingpropertytype *AerialRefuellingPropertyType) GongCopy() GongstructIF {
	newInstance := new(AerialRefuellingPropertyType)
	aerialrefuellingpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airtrafficcontrolservicetimeslicepropertytype *AirTrafficControlServiceTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(AirTrafficControlServiceTimeSlicePropertyType)
	airtrafficcontrolservicetimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airtrafficcontrolservicetimeslicetype *AirTrafficControlServiceTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(AirTrafficControlServiceTimeSliceType)
	airtrafficcontrolservicetimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airtrafficcontrolservicetype *AirTrafficControlServiceType) GongCopy() GongstructIF {
	newInstance := new(AirTrafficControlServiceType)
	airtrafficcontrolservicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airtrafficmanagementservicetimeslicepropertytype *AirTrafficManagementServiceTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(AirTrafficManagementServiceTimeSlicePropertyType)
	airtrafficmanagementservicetimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airtrafficmanagementservicetimeslicetype *AirTrafficManagementServiceTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(AirTrafficManagementServiceTimeSliceType)
	airtrafficmanagementservicetimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airtrafficmanagementservicetype *AirTrafficManagementServiceType) GongCopy() GongstructIF {
	newInstance := new(AirTrafficManagementServiceType)
	airtrafficmanagementservicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (aircraftgroundservicetimeslicepropertytype *AircraftGroundServiceTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(AircraftGroundServiceTimeSlicePropertyType)
	aircraftgroundservicetimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (aircraftgroundservicetimeslicetype *AircraftGroundServiceTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(AircraftGroundServiceTimeSliceType)
	aircraftgroundservicetimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (aircraftgroundservicetype *AircraftGroundServiceType) GongCopy() GongstructIF {
	newInstance := new(AircraftGroundServiceType)
	aircraftgroundservicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airportclearanceservicetimeslicepropertytype *AirportClearanceServiceTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(AirportClearanceServiceTimeSlicePropertyType)
	airportclearanceservicetimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airportclearanceservicetimeslicetype *AirportClearanceServiceTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(AirportClearanceServiceTimeSliceType)
	airportclearanceservicetimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airportclearanceservicetype *AirportClearanceServiceType) GongCopy() GongstructIF {
	newInstance := new(AirportClearanceServiceType)
	airportclearanceservicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongCopy() GongstructIF {
	newInstance := new(AirportHeliportPropertyType)
	airportheliportpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airportsuppliesservicetimeslicepropertytype *AirportSuppliesServiceTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(AirportSuppliesServiceTimeSlicePropertyType)
	airportsuppliesservicetimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airportsuppliesservicetimeslicetype *AirportSuppliesServiceTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(AirportSuppliesServiceTimeSliceType)
	airportsuppliesservicetimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airportsuppliesservicetype *AirportSuppliesServiceType) GongCopy() GongstructIF {
	newInstance := new(AirportSuppliesServiceType)
	airportsuppliesservicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (airspacepropertytype *AirspacePropertyType) GongCopy() GongstructIF {
	newInstance := new(AirspacePropertyType)
	airspacepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (callsigndetailpropertytype *CallsignDetailPropertyType) GongCopy() GongstructIF {
	newInstance := new(CallsignDetailPropertyType)
	callsigndetailpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (callsigndetailtype *CallsignDetailType) GongCopy() GongstructIF {
	newInstance := new(CallsignDetailType)
	callsigndetailtype.GongCopyBasicFields(newInstance)
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

func (designatedpointpropertytype *DesignatedPointPropertyType) GongCopy() GongstructIF {
	newInstance := new(DesignatedPointPropertyType)
	designatedpointpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (directionfinderpropertytype *DirectionFinderPropertyType) GongCopy() GongstructIF {
	newInstance := new(DirectionFinderPropertyType)
	directionfinderpropertytype.GongCopyBasicFields(newInstance)
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

func (firefightingservicetimeslicepropertytype *FireFightingServiceTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(FireFightingServiceTimeSlicePropertyType)
	firefightingservicetimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (firefightingservicetimeslicetype *FireFightingServiceTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(FireFightingServiceTimeSliceType)
	firefightingservicetimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (firefightingservicetype *FireFightingServiceType) GongCopy() GongstructIF {
	newInstance := new(FireFightingServiceType)
	firefightingservicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (fuelpropertytype *FuelPropertyType) GongCopy() GongstructIF {
	newInstance := new(FuelPropertyType)
	fuelpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (fueltype *FuelType) GongCopy() GongstructIF {
	newInstance := new(FuelType)
	fueltype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongCopy() GongstructIF {
	newInstance := new(GroundLightSystemPropertyType)
	groundlightsystempropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (groundtrafficcontrolservicetimeslicepropertytype *GroundTrafficControlServiceTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(GroundTrafficControlServiceTimeSlicePropertyType)
	groundtrafficcontrolservicetimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (groundtrafficcontrolservicetimeslicetype *GroundTrafficControlServiceTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(GroundTrafficControlServiceTimeSliceType)
	groundtrafficcontrolservicetimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (groundtrafficcontrolservicetype *GroundTrafficControlServiceType) GongCopy() GongstructIF {
	newInstance := new(GroundTrafficControlServiceType)
	groundtrafficcontrolservicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (holdingpatternpropertytype *HoldingPatternPropertyType) GongCopy() GongstructIF {
	newInstance := new(HoldingPatternPropertyType)
	holdingpatternpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (informationservicetimeslicepropertytype *InformationServiceTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(InformationServiceTimeSlicePropertyType)
	informationservicetimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (informationservicetimeslicetype *InformationServiceTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(InformationServiceTimeSliceType)
	informationservicetimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (informationservicetype *InformationServiceType) GongCopy() GongstructIF {
	newInstance := new(InformationServiceType)
	informationservicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (lightactivationpropertytype *LightActivationPropertyType) GongCopy() GongstructIF {
	newInstance := new(LightActivationPropertyType)
	lightactivationpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (lightactivationtype *LightActivationType) GongCopy() GongstructIF {
	newInstance := new(LightActivationType)
	lightactivationtype.GongCopyBasicFields(newInstance)
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

func (navaidpropertytype *NavaidPropertyType) GongCopy() GongstructIF {
	newInstance := new(NavaidPropertyType)
	navaidpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (nitrogenpropertytype *NitrogenPropertyType) GongCopy() GongstructIF {
	newInstance := new(NitrogenPropertyType)
	nitrogenpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (nitrogentype *NitrogenType) GongCopy() GongstructIF {
	newInstance := new(NitrogenType)
	nitrogentype.GongCopyBasicFields(newInstance)
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

func (oilpropertytype *OilPropertyType) GongCopy() GongstructIF {
	newInstance := new(OilPropertyType)
	oilpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (oiltype *OilType) GongCopy() GongstructIF {
	newInstance := new(OilType)
	oiltype.GongCopyBasicFields(newInstance)
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

func (organisationauthorityassociationpropertytype *OrganisationAuthorityAssociationPropertyType) GongCopy() GongstructIF {
	newInstance := new(OrganisationAuthorityAssociationPropertyType)
	organisationauthorityassociationpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (organisationauthorityassociationtype *OrganisationAuthorityAssociationType) GongCopy() GongstructIF {
	newInstance := new(OrganisationAuthorityAssociationType)
	organisationauthorityassociationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongCopy() GongstructIF {
	newInstance := new(OrganisationAuthorityPropertyType)
	organisationauthoritypropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (organisationauthoritytimeslicepropertytype *OrganisationAuthorityTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(OrganisationAuthorityTimeSlicePropertyType)
	organisationauthoritytimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (organisationauthoritytimeslicetype *OrganisationAuthorityTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(OrganisationAuthorityTimeSliceType)
	organisationauthoritytimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (organisationauthoritytype *OrganisationAuthorityType) GongCopy() GongstructIF {
	newInstance := new(OrganisationAuthorityType)
	organisationauthoritytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (oxygenpropertytype *OxygenPropertyType) GongCopy() GongstructIF {
	newInstance := new(OxygenPropertyType)
	oxygenpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (oxygentype *OxygenType) GongCopy() GongstructIF {
	newInstance := new(OxygenType)
	oxygentype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (passengerservicetimeslicepropertytype *PassengerServiceTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(PassengerServiceTimeSlicePropertyType)
	passengerservicetimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (passengerservicetimeslicetype *PassengerServiceTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(PassengerServiceTimeSliceType)
	passengerservicetimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (passengerservicetype *PassengerServiceType) GongCopy() GongstructIF {
	newInstance := new(PassengerServiceType)
	passengerservicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (pilotcontrolledlightingtimeslicepropertytype *PilotControlledLightingTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(PilotControlledLightingTimeSlicePropertyType)
	pilotcontrolledlightingtimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (pilotcontrolledlightingtimeslicetype *PilotControlledLightingTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(PilotControlledLightingTimeSliceType)
	pilotcontrolledlightingtimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (pilotcontrolledlightingtype *PilotControlledLightingType) GongCopy() GongstructIF {
	newInstance := new(PilotControlledLightingType)
	pilotcontrolledlightingtype.GongCopyBasicFields(newInstance)
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

func (procedurepropertytype *ProcedurePropertyType) GongCopy() GongstructIF {
	newInstance := new(ProcedurePropertyType)
	procedurepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (radiocommunicationchannelpropertytype *RadioCommunicationChannelPropertyType) GongCopy() GongstructIF {
	newInstance := new(RadioCommunicationChannelPropertyType)
	radiocommunicationchannelpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (radiocommunicationchanneltimeslicepropertytype *RadioCommunicationChannelTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(RadioCommunicationChannelTimeSlicePropertyType)
	radiocommunicationchanneltimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (radiocommunicationchanneltimeslicetype *RadioCommunicationChannelTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(RadioCommunicationChannelTimeSliceType)
	radiocommunicationchanneltimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (radiocommunicationchanneltype *RadioCommunicationChannelType) GongCopy() GongstructIF {
	newInstance := new(RadioCommunicationChannelType)
	radiocommunicationchanneltype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (radiocommunicationoperationalstatuspropertytype *RadioCommunicationOperationalStatusPropertyType) GongCopy() GongstructIF {
	newInstance := new(RadioCommunicationOperationalStatusPropertyType)
	radiocommunicationoperationalstatuspropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (radiocommunicationoperationalstatustype *RadioCommunicationOperationalStatusType) GongCopy() GongstructIF {
	newInstance := new(RadioCommunicationOperationalStatusType)
	radiocommunicationoperationalstatustype.GongCopyBasicFields(newInstance)
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

func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongCopy() GongstructIF {
	newInstance := new(RunwayCentrelinePointPropertyType)
	runwaycentrelinepointpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (searchrescueservicetimeslicepropertytype *SearchRescueServiceTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(SearchRescueServiceTimeSlicePropertyType)
	searchrescueservicetimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (searchrescueservicetimeslicetype *SearchRescueServiceTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(SearchRescueServiceTimeSliceType)
	searchrescueservicetimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (searchrescueservicetype *SearchRescueServiceType) GongCopy() GongstructIF {
	newInstance := new(SearchRescueServiceType)
	searchrescueservicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (serviceoperationalstatuspropertytype *ServiceOperationalStatusPropertyType) GongCopy() GongstructIF {
	newInstance := new(ServiceOperationalStatusPropertyType)
	serviceoperationalstatuspropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (serviceoperationalstatustype *ServiceOperationalStatusType) GongCopy() GongstructIF {
	newInstance := new(ServiceOperationalStatusType)
	serviceoperationalstatustype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (specialdatetimeslicepropertytype *SpecialDateTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(SpecialDateTimeSlicePropertyType)
	specialdatetimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (specialdatetimeslicetype *SpecialDateTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(SpecialDateTimeSliceType)
	specialdatetimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (specialdatetype *SpecialDateType) GongCopy() GongstructIF {
	newInstance := new(SpecialDateType)
	specialdatetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongCopy() GongstructIF {
	newInstance := new(StandardLevelColumnPropertyType)
	standardlevelcolumnpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (standardlevelcolumntimeslicepropertytype *StandardLevelColumnTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(StandardLevelColumnTimeSlicePropertyType)
	standardlevelcolumntimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (standardlevelcolumntimeslicetype *StandardLevelColumnTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(StandardLevelColumnTimeSliceType)
	standardlevelcolumntimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (standardlevelcolumntype *StandardLevelColumnType) GongCopy() GongstructIF {
	newInstance := new(StandardLevelColumnType)
	standardlevelcolumntype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (standardlevelpropertytype *StandardLevelPropertyType) GongCopy() GongstructIF {
	newInstance := new(StandardLevelPropertyType)
	standardlevelpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (standardlevelsectortimeslicepropertytype *StandardLevelSectorTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(StandardLevelSectorTimeSlicePropertyType)
	standardlevelsectortimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (standardlevelsectortimeslicetype *StandardLevelSectorTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(StandardLevelSectorTimeSliceType)
	standardlevelsectortimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (standardlevelsectortype *StandardLevelSectorType) GongCopy() GongstructIF {
	newInstance := new(StandardLevelSectorType)
	standardlevelsectortype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (standardleveltablepropertytype *StandardLevelTablePropertyType) GongCopy() GongstructIF {
	newInstance := new(StandardLevelTablePropertyType)
	standardleveltablepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (standardleveltabletimeslicepropertytype *StandardLevelTableTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(StandardLevelTableTimeSlicePropertyType)
	standardleveltabletimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (standardleveltabletimeslicetype *StandardLevelTableTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(StandardLevelTableTimeSliceType)
	standardleveltabletimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (standardleveltabletype *StandardLevelTableType) GongCopy() GongstructIF {
	newInstance := new(StandardLevelTableType)
	standardleveltabletype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (standardleveltype *StandardLevelType) GongCopy() GongstructIF {
	newInstance := new(StandardLevelType)
	standardleveltype.GongCopyBasicFields(newInstance)
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

func (unitavailabilitypropertytype *UnitAvailabilityPropertyType) GongCopy() GongstructIF {
	newInstance := new(UnitAvailabilityPropertyType)
	unitavailabilitypropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (unitavailabilitytype *UnitAvailabilityType) GongCopy() GongstructIF {
	newInstance := new(UnitAvailabilityType)
	unitavailabilitytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (unitdependencypropertytype *UnitDependencyPropertyType) GongCopy() GongstructIF {
	newInstance := new(UnitDependencyPropertyType)
	unitdependencypropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (unitdependencytype *UnitDependencyType) GongCopy() GongstructIF {
	newInstance := new(UnitDependencyType)
	unitdependencytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (unitpropertytype *UnitPropertyType) GongCopy() GongstructIF {
	newInstance := new(UnitPropertyType)
	unitpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (unittimeslicepropertytype *UnitTimeSlicePropertyType) GongCopy() GongstructIF {
	newInstance := new(UnitTimeSlicePropertyType)
	unittimeslicepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (unittimeslicetype *UnitTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(UnitTimeSliceType)
	unittimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (unittype *UnitType) GongCopy() GongstructIF {
	newInstance := new(UnitType)
	unittype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (vorpropertytype *VORPropertyType) GongCopy() GongstructIF {
	newInstance := new(VORPropertyType)
	vorpropertytype.GongCopyBasicFields(newInstance)
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

func (abstractairportgroundservicetype *AbstractAirportGroundServiceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractairportgroundservicetype)
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractpropertieswithscheduletype)
}

func (abstractservicetype *AbstractServiceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractservicetype)
}

func (abstracttrafficseparationservicetype *AbstractTrafficSeparationServiceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstracttrafficseparationservicetype)
}

func (aerialrefuellingpropertytype *AerialRefuellingPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, aerialrefuellingpropertytype)
}

func (airtrafficcontrolservicetimeslicepropertytype *AirTrafficControlServiceTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airtrafficcontrolservicetimeslicepropertytype)
}

func (airtrafficcontrolservicetimeslicetype *AirTrafficControlServiceTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airtrafficcontrolservicetimeslicetype)
}

func (airtrafficcontrolservicetype *AirTrafficControlServiceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airtrafficcontrolservicetype)
}

func (airtrafficmanagementservicetimeslicepropertytype *AirTrafficManagementServiceTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airtrafficmanagementservicetimeslicepropertytype)
}

func (airtrafficmanagementservicetimeslicetype *AirTrafficManagementServiceTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airtrafficmanagementservicetimeslicetype)
}

func (airtrafficmanagementservicetype *AirTrafficManagementServiceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airtrafficmanagementservicetype)
}

func (aircraftgroundservicetimeslicepropertytype *AircraftGroundServiceTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, aircraftgroundservicetimeslicepropertytype)
}

func (aircraftgroundservicetimeslicetype *AircraftGroundServiceTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, aircraftgroundservicetimeslicetype)
}

func (aircraftgroundservicetype *AircraftGroundServiceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, aircraftgroundservicetype)
}

func (airportclearanceservicetimeslicepropertytype *AirportClearanceServiceTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airportclearanceservicetimeslicepropertytype)
}

func (airportclearanceservicetimeslicetype *AirportClearanceServiceTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airportclearanceservicetimeslicetype)
}

func (airportclearanceservicetype *AirportClearanceServiceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airportclearanceservicetype)
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airportheliportpropertytype)
}

func (airportsuppliesservicetimeslicepropertytype *AirportSuppliesServiceTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airportsuppliesservicetimeslicepropertytype)
}

func (airportsuppliesservicetimeslicetype *AirportSuppliesServiceTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airportsuppliesservicetimeslicetype)
}

func (airportsuppliesservicetype *AirportSuppliesServiceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airportsuppliesservicetype)
}

func (airspacepropertytype *AirspacePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, airspacepropertytype)
}

func (callsigndetailpropertytype *CallsignDetailPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, callsigndetailpropertytype)
}

func (callsigndetailtype *CallsignDetailType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, callsigndetailtype)
}

func (contactinformationpropertytype *ContactInformationPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, contactinformationpropertytype)
}

func (contactinformationtype *ContactInformationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, contactinformationtype)
}

func (designatedpointpropertytype *DesignatedPointPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, designatedpointpropertytype)
}

func (directionfinderpropertytype *DirectionFinderPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, directionfinderpropertytype)
}

func (elevatedpointpropertytype *ElevatedPointPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, elevatedpointpropertytype)
}

func (elevatedpointtype *ElevatedPointType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, elevatedpointtype)
}

func (firefightingservicetimeslicepropertytype *FireFightingServiceTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, firefightingservicetimeslicepropertytype)
}

func (firefightingservicetimeslicetype *FireFightingServiceTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, firefightingservicetimeslicetype)
}

func (firefightingservicetype *FireFightingServiceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, firefightingservicetype)
}

func (fuelpropertytype *FuelPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, fuelpropertytype)
}

func (fueltype *FuelType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, fueltype)
}

func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, groundlightsystempropertytype)
}

func (groundtrafficcontrolservicetimeslicepropertytype *GroundTrafficControlServiceTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, groundtrafficcontrolservicetimeslicepropertytype)
}

func (groundtrafficcontrolservicetimeslicetype *GroundTrafficControlServiceTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, groundtrafficcontrolservicetimeslicetype)
}

func (groundtrafficcontrolservicetype *GroundTrafficControlServiceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, groundtrafficcontrolservicetype)
}

func (holdingpatternpropertytype *HoldingPatternPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, holdingpatternpropertytype)
}

func (informationservicetimeslicepropertytype *InformationServiceTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, informationservicetimeslicepropertytype)
}

func (informationservicetimeslicetype *InformationServiceTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, informationservicetimeslicetype)
}

func (informationservicetype *InformationServiceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, informationservicetype)
}

func (lightactivationpropertytype *LightActivationPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, lightactivationpropertytype)
}

func (lightactivationtype *LightActivationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, lightactivationtype)
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, linguisticnotepropertytype)
}

func (linguisticnotetype *LinguisticNoteType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, linguisticnotetype)
}

func (navaidpropertytype *NavaidPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, navaidpropertytype)
}

func (nitrogenpropertytype *NitrogenPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, nitrogenpropertytype)
}

func (nitrogentype *NitrogenType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, nitrogentype)
}

func (notepropertytype *NotePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, notepropertytype)
}

func (notetype *NoteType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, notetype)
}

func (oilpropertytype *OilPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, oilpropertytype)
}

func (oiltype *OilType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, oiltype)
}

func (onlinecontactpropertytype *OnlineContactPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, onlinecontactpropertytype)
}

func (onlinecontacttype *OnlineContactType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, onlinecontacttype)
}

func (organisationauthorityassociationpropertytype *OrganisationAuthorityAssociationPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, organisationauthorityassociationpropertytype)
}

func (organisationauthorityassociationtype *OrganisationAuthorityAssociationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, organisationauthorityassociationtype)
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, organisationauthoritypropertytype)
}

func (organisationauthoritytimeslicepropertytype *OrganisationAuthorityTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, organisationauthoritytimeslicepropertytype)
}

func (organisationauthoritytimeslicetype *OrganisationAuthorityTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, organisationauthoritytimeslicetype)
}

func (organisationauthoritytype *OrganisationAuthorityType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, organisationauthoritytype)
}

func (oxygenpropertytype *OxygenPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, oxygenpropertytype)
}

func (oxygentype *OxygenType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, oxygentype)
}

func (passengerservicetimeslicepropertytype *PassengerServiceTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, passengerservicetimeslicepropertytype)
}

func (passengerservicetimeslicetype *PassengerServiceTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, passengerservicetimeslicetype)
}

func (passengerservicetype *PassengerServiceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, passengerservicetype)
}

func (pilotcontrolledlightingtimeslicepropertytype *PilotControlledLightingTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, pilotcontrolledlightingtimeslicepropertytype)
}

func (pilotcontrolledlightingtimeslicetype *PilotControlledLightingTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, pilotcontrolledlightingtimeslicetype)
}

func (pilotcontrolledlightingtype *PilotControlledLightingType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, pilotcontrolledlightingtype)
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

func (procedurepropertytype *ProcedurePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, procedurepropertytype)
}

func (radiocommunicationchannelpropertytype *RadioCommunicationChannelPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, radiocommunicationchannelpropertytype)
}

func (radiocommunicationchanneltimeslicepropertytype *RadioCommunicationChannelTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, radiocommunicationchanneltimeslicepropertytype)
}

func (radiocommunicationchanneltimeslicetype *RadioCommunicationChannelTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, radiocommunicationchanneltimeslicetype)
}

func (radiocommunicationchanneltype *RadioCommunicationChannelType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, radiocommunicationchanneltype)
}

func (radiocommunicationoperationalstatuspropertytype *RadioCommunicationOperationalStatusPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, radiocommunicationoperationalstatuspropertytype)
}

func (radiocommunicationoperationalstatustype *RadioCommunicationOperationalStatusType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, radiocommunicationoperationalstatustype)
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

func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, runwaycentrelinepointpropertytype)
}

func (searchrescueservicetimeslicepropertytype *SearchRescueServiceTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, searchrescueservicetimeslicepropertytype)
}

func (searchrescueservicetimeslicetype *SearchRescueServiceTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, searchrescueservicetimeslicetype)
}

func (searchrescueservicetype *SearchRescueServiceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, searchrescueservicetype)
}

func (serviceoperationalstatuspropertytype *ServiceOperationalStatusPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, serviceoperationalstatuspropertytype)
}

func (serviceoperationalstatustype *ServiceOperationalStatusType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, serviceoperationalstatustype)
}

func (specialdatetimeslicepropertytype *SpecialDateTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, specialdatetimeslicepropertytype)
}

func (specialdatetimeslicetype *SpecialDateTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, specialdatetimeslicetype)
}

func (specialdatetype *SpecialDateType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, specialdatetype)
}

func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, standardlevelcolumnpropertytype)
}

func (standardlevelcolumntimeslicepropertytype *StandardLevelColumnTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, standardlevelcolumntimeslicepropertytype)
}

func (standardlevelcolumntimeslicetype *StandardLevelColumnTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, standardlevelcolumntimeslicetype)
}

func (standardlevelcolumntype *StandardLevelColumnType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, standardlevelcolumntype)
}

func (standardlevelpropertytype *StandardLevelPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, standardlevelpropertytype)
}

func (standardlevelsectortimeslicepropertytype *StandardLevelSectorTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, standardlevelsectortimeslicepropertytype)
}

func (standardlevelsectortimeslicetype *StandardLevelSectorTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, standardlevelsectortimeslicetype)
}

func (standardlevelsectortype *StandardLevelSectorType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, standardlevelsectortype)
}

func (standardleveltablepropertytype *StandardLevelTablePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, standardleveltablepropertytype)
}

func (standardleveltabletimeslicepropertytype *StandardLevelTableTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, standardleveltabletimeslicepropertytype)
}

func (standardleveltabletimeslicetype *StandardLevelTableTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, standardleveltabletimeslicetype)
}

func (standardleveltabletype *StandardLevelTableType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, standardleveltabletype)
}

func (standardleveltype *StandardLevelType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, standardleveltype)
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

func (unitavailabilitypropertytype *UnitAvailabilityPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, unitavailabilitypropertytype)
}

func (unitavailabilitytype *UnitAvailabilityType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, unitavailabilitytype)
}

func (unitdependencypropertytype *UnitDependencyPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, unitdependencypropertytype)
}

func (unitdependencytype *UnitDependencyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, unitdependencytype)
}

func (unitpropertytype *UnitPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, unitpropertytype)
}

func (unittimeslicepropertytype *UnitTimeSlicePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, unittimeslicepropertytype)
}

func (unittimeslicetype *UnitTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, unittimeslicetype)
}

func (unittype *UnitType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, unittype)
}

func (vorpropertytype *VORPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, vorpropertytype)
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
		stage.AbstractAirportGroundServiceTypes,
		stage.AbstractAirportGroundServiceType_stagedOrder,
		stage.AbstractAirportGroundServiceTypes_reference,
		&stage.AbstractAirportGroundServiceTypes_referenceOrder,
		stage.AbstractAirportGroundServiceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
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
		stage.AbstractServiceTypes,
		stage.AbstractServiceType_stagedOrder,
		stage.AbstractServiceTypes_reference,
		&stage.AbstractServiceTypes_referenceOrder,
		stage.AbstractServiceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AbstractTrafficSeparationServiceTypes,
		stage.AbstractTrafficSeparationServiceType_stagedOrder,
		stage.AbstractTrafficSeparationServiceTypes_reference,
		&stage.AbstractTrafficSeparationServiceTypes_referenceOrder,
		stage.AbstractTrafficSeparationServiceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
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
		stage.AirTrafficControlServiceTimeSlicePropertyTypes,
		stage.AirTrafficControlServiceTimeSlicePropertyType_stagedOrder,
		stage.AirTrafficControlServiceTimeSlicePropertyTypes_reference,
		&stage.AirTrafficControlServiceTimeSlicePropertyTypes_referenceOrder,
		stage.AirTrafficControlServiceTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AirTrafficControlServiceTimeSliceTypes,
		stage.AirTrafficControlServiceTimeSliceType_stagedOrder,
		stage.AirTrafficControlServiceTimeSliceTypes_reference,
		&stage.AirTrafficControlServiceTimeSliceTypes_referenceOrder,
		stage.AirTrafficControlServiceTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AirTrafficControlServiceTypes,
		stage.AirTrafficControlServiceType_stagedOrder,
		stage.AirTrafficControlServiceTypes_reference,
		&stage.AirTrafficControlServiceTypes_referenceOrder,
		stage.AirTrafficControlServiceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AirTrafficManagementServiceTimeSlicePropertyTypes,
		stage.AirTrafficManagementServiceTimeSlicePropertyType_stagedOrder,
		stage.AirTrafficManagementServiceTimeSlicePropertyTypes_reference,
		&stage.AirTrafficManagementServiceTimeSlicePropertyTypes_referenceOrder,
		stage.AirTrafficManagementServiceTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AirTrafficManagementServiceTimeSliceTypes,
		stage.AirTrafficManagementServiceTimeSliceType_stagedOrder,
		stage.AirTrafficManagementServiceTimeSliceTypes_reference,
		&stage.AirTrafficManagementServiceTimeSliceTypes_referenceOrder,
		stage.AirTrafficManagementServiceTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AirTrafficManagementServiceTypes,
		stage.AirTrafficManagementServiceType_stagedOrder,
		stage.AirTrafficManagementServiceTypes_reference,
		&stage.AirTrafficManagementServiceTypes_referenceOrder,
		stage.AirTrafficManagementServiceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AircraftGroundServiceTimeSlicePropertyTypes,
		stage.AircraftGroundServiceTimeSlicePropertyType_stagedOrder,
		stage.AircraftGroundServiceTimeSlicePropertyTypes_reference,
		&stage.AircraftGroundServiceTimeSlicePropertyTypes_referenceOrder,
		stage.AircraftGroundServiceTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AircraftGroundServiceTimeSliceTypes,
		stage.AircraftGroundServiceTimeSliceType_stagedOrder,
		stage.AircraftGroundServiceTimeSliceTypes_reference,
		&stage.AircraftGroundServiceTimeSliceTypes_referenceOrder,
		stage.AircraftGroundServiceTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AircraftGroundServiceTypes,
		stage.AircraftGroundServiceType_stagedOrder,
		stage.AircraftGroundServiceTypes_reference,
		&stage.AircraftGroundServiceTypes_referenceOrder,
		stage.AircraftGroundServiceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AirportClearanceServiceTimeSlicePropertyTypes,
		stage.AirportClearanceServiceTimeSlicePropertyType_stagedOrder,
		stage.AirportClearanceServiceTimeSlicePropertyTypes_reference,
		&stage.AirportClearanceServiceTimeSlicePropertyTypes_referenceOrder,
		stage.AirportClearanceServiceTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AirportClearanceServiceTimeSliceTypes,
		stage.AirportClearanceServiceTimeSliceType_stagedOrder,
		stage.AirportClearanceServiceTimeSliceTypes_reference,
		&stage.AirportClearanceServiceTimeSliceTypes_referenceOrder,
		stage.AirportClearanceServiceTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AirportClearanceServiceTypes,
		stage.AirportClearanceServiceType_stagedOrder,
		stage.AirportClearanceServiceTypes_reference,
		&stage.AirportClearanceServiceTypes_referenceOrder,
		stage.AirportClearanceServiceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
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
		stage.AirportSuppliesServiceTimeSlicePropertyTypes,
		stage.AirportSuppliesServiceTimeSlicePropertyType_stagedOrder,
		stage.AirportSuppliesServiceTimeSlicePropertyTypes_reference,
		&stage.AirportSuppliesServiceTimeSlicePropertyTypes_referenceOrder,
		stage.AirportSuppliesServiceTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AirportSuppliesServiceTimeSliceTypes,
		stage.AirportSuppliesServiceTimeSliceType_stagedOrder,
		stage.AirportSuppliesServiceTimeSliceTypes_reference,
		&stage.AirportSuppliesServiceTimeSliceTypes_referenceOrder,
		stage.AirportSuppliesServiceTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AirportSuppliesServiceTypes,
		stage.AirportSuppliesServiceType_stagedOrder,
		stage.AirportSuppliesServiceTypes_reference,
		&stage.AirportSuppliesServiceTypes_referenceOrder,
		stage.AirportSuppliesServiceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
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
		stage.CallsignDetailPropertyTypes,
		stage.CallsignDetailPropertyType_stagedOrder,
		stage.CallsignDetailPropertyTypes_reference,
		&stage.CallsignDetailPropertyTypes_referenceOrder,
		stage.CallsignDetailPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CallsignDetailTypes,
		stage.CallsignDetailType_stagedOrder,
		stage.CallsignDetailTypes_reference,
		&stage.CallsignDetailTypes_referenceOrder,
		stage.CallsignDetailTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
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
		stage.DirectionFinderPropertyTypes,
		stage.DirectionFinderPropertyType_stagedOrder,
		stage.DirectionFinderPropertyTypes_reference,
		&stage.DirectionFinderPropertyTypes_referenceOrder,
		stage.DirectionFinderPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
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
		stage.FireFightingServiceTimeSlicePropertyTypes,
		stage.FireFightingServiceTimeSlicePropertyType_stagedOrder,
		stage.FireFightingServiceTimeSlicePropertyTypes_reference,
		&stage.FireFightingServiceTimeSlicePropertyTypes_referenceOrder,
		stage.FireFightingServiceTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.FireFightingServiceTimeSliceTypes,
		stage.FireFightingServiceTimeSliceType_stagedOrder,
		stage.FireFightingServiceTimeSliceTypes_reference,
		&stage.FireFightingServiceTimeSliceTypes_referenceOrder,
		stage.FireFightingServiceTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.FireFightingServiceTypes,
		stage.FireFightingServiceType_stagedOrder,
		stage.FireFightingServiceTypes_reference,
		&stage.FireFightingServiceTypes_referenceOrder,
		stage.FireFightingServiceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.FuelPropertyTypes,
		stage.FuelPropertyType_stagedOrder,
		stage.FuelPropertyTypes_reference,
		&stage.FuelPropertyTypes_referenceOrder,
		stage.FuelPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.FuelTypes,
		stage.FuelType_stagedOrder,
		stage.FuelTypes_reference,
		&stage.FuelTypes_referenceOrder,
		stage.FuelTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.GroundLightSystemPropertyTypes,
		stage.GroundLightSystemPropertyType_stagedOrder,
		stage.GroundLightSystemPropertyTypes_reference,
		&stage.GroundLightSystemPropertyTypes_referenceOrder,
		stage.GroundLightSystemPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.GroundTrafficControlServiceTimeSlicePropertyTypes,
		stage.GroundTrafficControlServiceTimeSlicePropertyType_stagedOrder,
		stage.GroundTrafficControlServiceTimeSlicePropertyTypes_reference,
		&stage.GroundTrafficControlServiceTimeSlicePropertyTypes_referenceOrder,
		stage.GroundTrafficControlServiceTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.GroundTrafficControlServiceTimeSliceTypes,
		stage.GroundTrafficControlServiceTimeSliceType_stagedOrder,
		stage.GroundTrafficControlServiceTimeSliceTypes_reference,
		&stage.GroundTrafficControlServiceTimeSliceTypes_referenceOrder,
		stage.GroundTrafficControlServiceTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.GroundTrafficControlServiceTypes,
		stage.GroundTrafficControlServiceType_stagedOrder,
		stage.GroundTrafficControlServiceTypes_reference,
		&stage.GroundTrafficControlServiceTypes_referenceOrder,
		stage.GroundTrafficControlServiceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.HoldingPatternPropertyTypes,
		stage.HoldingPatternPropertyType_stagedOrder,
		stage.HoldingPatternPropertyTypes_reference,
		&stage.HoldingPatternPropertyTypes_referenceOrder,
		stage.HoldingPatternPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.InformationServiceTimeSlicePropertyTypes,
		stage.InformationServiceTimeSlicePropertyType_stagedOrder,
		stage.InformationServiceTimeSlicePropertyTypes_reference,
		&stage.InformationServiceTimeSlicePropertyTypes_referenceOrder,
		stage.InformationServiceTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.InformationServiceTimeSliceTypes,
		stage.InformationServiceTimeSliceType_stagedOrder,
		stage.InformationServiceTimeSliceTypes_reference,
		&stage.InformationServiceTimeSliceTypes_referenceOrder,
		stage.InformationServiceTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.InformationServiceTypes,
		stage.InformationServiceType_stagedOrder,
		stage.InformationServiceTypes_reference,
		&stage.InformationServiceTypes_referenceOrder,
		stage.InformationServiceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.LightActivationPropertyTypes,
		stage.LightActivationPropertyType_stagedOrder,
		stage.LightActivationPropertyTypes_reference,
		&stage.LightActivationPropertyTypes_referenceOrder,
		stage.LightActivationPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.LightActivationTypes,
		stage.LightActivationType_stagedOrder,
		stage.LightActivationTypes_reference,
		&stage.LightActivationTypes_referenceOrder,
		stage.LightActivationTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
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
		stage.NitrogenPropertyTypes,
		stage.NitrogenPropertyType_stagedOrder,
		stage.NitrogenPropertyTypes_reference,
		&stage.NitrogenPropertyTypes_referenceOrder,
		stage.NitrogenPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.NitrogenTypes,
		stage.NitrogenType_stagedOrder,
		stage.NitrogenTypes_reference,
		&stage.NitrogenTypes_referenceOrder,
		stage.NitrogenTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
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
		stage.OilPropertyTypes,
		stage.OilPropertyType_stagedOrder,
		stage.OilPropertyTypes_reference,
		&stage.OilPropertyTypes_referenceOrder,
		stage.OilPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.OilTypes,
		stage.OilType_stagedOrder,
		stage.OilTypes_reference,
		&stage.OilTypes_referenceOrder,
		stage.OilTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
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
		stage.OrganisationAuthorityAssociationPropertyTypes,
		stage.OrganisationAuthorityAssociationPropertyType_stagedOrder,
		stage.OrganisationAuthorityAssociationPropertyTypes_reference,
		&stage.OrganisationAuthorityAssociationPropertyTypes_referenceOrder,
		stage.OrganisationAuthorityAssociationPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.OrganisationAuthorityAssociationTypes,
		stage.OrganisationAuthorityAssociationType_stagedOrder,
		stage.OrganisationAuthorityAssociationTypes_reference,
		&stage.OrganisationAuthorityAssociationTypes_referenceOrder,
		stage.OrganisationAuthorityAssociationTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
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
		stage.OrganisationAuthorityTimeSlicePropertyTypes,
		stage.OrganisationAuthorityTimeSlicePropertyType_stagedOrder,
		stage.OrganisationAuthorityTimeSlicePropertyTypes_reference,
		&stage.OrganisationAuthorityTimeSlicePropertyTypes_referenceOrder,
		stage.OrganisationAuthorityTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.OrganisationAuthorityTimeSliceTypes,
		stage.OrganisationAuthorityTimeSliceType_stagedOrder,
		stage.OrganisationAuthorityTimeSliceTypes_reference,
		&stage.OrganisationAuthorityTimeSliceTypes_referenceOrder,
		stage.OrganisationAuthorityTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.OrganisationAuthorityTypes,
		stage.OrganisationAuthorityType_stagedOrder,
		stage.OrganisationAuthorityTypes_reference,
		&stage.OrganisationAuthorityTypes_referenceOrder,
		stage.OrganisationAuthorityTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.OxygenPropertyTypes,
		stage.OxygenPropertyType_stagedOrder,
		stage.OxygenPropertyTypes_reference,
		&stage.OxygenPropertyTypes_referenceOrder,
		stage.OxygenPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.OxygenTypes,
		stage.OxygenType_stagedOrder,
		stage.OxygenTypes_reference,
		&stage.OxygenTypes_referenceOrder,
		stage.OxygenTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.PassengerServiceTimeSlicePropertyTypes,
		stage.PassengerServiceTimeSlicePropertyType_stagedOrder,
		stage.PassengerServiceTimeSlicePropertyTypes_reference,
		&stage.PassengerServiceTimeSlicePropertyTypes_referenceOrder,
		stage.PassengerServiceTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.PassengerServiceTimeSliceTypes,
		stage.PassengerServiceTimeSliceType_stagedOrder,
		stage.PassengerServiceTimeSliceTypes_reference,
		&stage.PassengerServiceTimeSliceTypes_referenceOrder,
		stage.PassengerServiceTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.PassengerServiceTypes,
		stage.PassengerServiceType_stagedOrder,
		stage.PassengerServiceTypes_reference,
		&stage.PassengerServiceTypes_referenceOrder,
		stage.PassengerServiceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.PilotControlledLightingTimeSlicePropertyTypes,
		stage.PilotControlledLightingTimeSlicePropertyType_stagedOrder,
		stage.PilotControlledLightingTimeSlicePropertyTypes_reference,
		&stage.PilotControlledLightingTimeSlicePropertyTypes_referenceOrder,
		stage.PilotControlledLightingTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.PilotControlledLightingTimeSliceTypes,
		stage.PilotControlledLightingTimeSliceType_stagedOrder,
		stage.PilotControlledLightingTimeSliceTypes_reference,
		&stage.PilotControlledLightingTimeSliceTypes_referenceOrder,
		stage.PilotControlledLightingTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.PilotControlledLightingTypes,
		stage.PilotControlledLightingType_stagedOrder,
		stage.PilotControlledLightingTypes_reference,
		&stage.PilotControlledLightingTypes_referenceOrder,
		stage.PilotControlledLightingTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
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
		stage.ProcedurePropertyTypes,
		stage.ProcedurePropertyType_stagedOrder,
		stage.ProcedurePropertyTypes_reference,
		&stage.ProcedurePropertyTypes_referenceOrder,
		stage.ProcedurePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.RadioCommunicationChannelPropertyTypes,
		stage.RadioCommunicationChannelPropertyType_stagedOrder,
		stage.RadioCommunicationChannelPropertyTypes_reference,
		&stage.RadioCommunicationChannelPropertyTypes_referenceOrder,
		stage.RadioCommunicationChannelPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.RadioCommunicationChannelTimeSlicePropertyTypes,
		stage.RadioCommunicationChannelTimeSlicePropertyType_stagedOrder,
		stage.RadioCommunicationChannelTimeSlicePropertyTypes_reference,
		&stage.RadioCommunicationChannelTimeSlicePropertyTypes_referenceOrder,
		stage.RadioCommunicationChannelTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.RadioCommunicationChannelTimeSliceTypes,
		stage.RadioCommunicationChannelTimeSliceType_stagedOrder,
		stage.RadioCommunicationChannelTimeSliceTypes_reference,
		&stage.RadioCommunicationChannelTimeSliceTypes_referenceOrder,
		stage.RadioCommunicationChannelTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.RadioCommunicationChannelTypes,
		stage.RadioCommunicationChannelType_stagedOrder,
		stage.RadioCommunicationChannelTypes_reference,
		&stage.RadioCommunicationChannelTypes_referenceOrder,
		stage.RadioCommunicationChannelTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.RadioCommunicationOperationalStatusPropertyTypes,
		stage.RadioCommunicationOperationalStatusPropertyType_stagedOrder,
		stage.RadioCommunicationOperationalStatusPropertyTypes_reference,
		&stage.RadioCommunicationOperationalStatusPropertyTypes_referenceOrder,
		stage.RadioCommunicationOperationalStatusPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.RadioCommunicationOperationalStatusTypes,
		stage.RadioCommunicationOperationalStatusType_stagedOrder,
		stage.RadioCommunicationOperationalStatusTypes_reference,
		&stage.RadioCommunicationOperationalStatusTypes_referenceOrder,
		stage.RadioCommunicationOperationalStatusTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
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
		stage.SearchRescueServiceTimeSlicePropertyTypes,
		stage.SearchRescueServiceTimeSlicePropertyType_stagedOrder,
		stage.SearchRescueServiceTimeSlicePropertyTypes_reference,
		&stage.SearchRescueServiceTimeSlicePropertyTypes_referenceOrder,
		stage.SearchRescueServiceTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SearchRescueServiceTimeSliceTypes,
		stage.SearchRescueServiceTimeSliceType_stagedOrder,
		stage.SearchRescueServiceTimeSliceTypes_reference,
		&stage.SearchRescueServiceTimeSliceTypes_referenceOrder,
		stage.SearchRescueServiceTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SearchRescueServiceTypes,
		stage.SearchRescueServiceType_stagedOrder,
		stage.SearchRescueServiceTypes_reference,
		&stage.SearchRescueServiceTypes_referenceOrder,
		stage.SearchRescueServiceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ServiceOperationalStatusPropertyTypes,
		stage.ServiceOperationalStatusPropertyType_stagedOrder,
		stage.ServiceOperationalStatusPropertyTypes_reference,
		&stage.ServiceOperationalStatusPropertyTypes_referenceOrder,
		stage.ServiceOperationalStatusPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ServiceOperationalStatusTypes,
		stage.ServiceOperationalStatusType_stagedOrder,
		stage.ServiceOperationalStatusTypes_reference,
		&stage.ServiceOperationalStatusTypes_referenceOrder,
		stage.ServiceOperationalStatusTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SpecialDateTimeSlicePropertyTypes,
		stage.SpecialDateTimeSlicePropertyType_stagedOrder,
		stage.SpecialDateTimeSlicePropertyTypes_reference,
		&stage.SpecialDateTimeSlicePropertyTypes_referenceOrder,
		stage.SpecialDateTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SpecialDateTimeSliceTypes,
		stage.SpecialDateTimeSliceType_stagedOrder,
		stage.SpecialDateTimeSliceTypes_reference,
		&stage.SpecialDateTimeSliceTypes_referenceOrder,
		stage.SpecialDateTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SpecialDateTypes,
		stage.SpecialDateType_stagedOrder,
		stage.SpecialDateTypes_reference,
		&stage.SpecialDateTypes_referenceOrder,
		stage.SpecialDateTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
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
		stage.StandardLevelColumnTimeSlicePropertyTypes,
		stage.StandardLevelColumnTimeSlicePropertyType_stagedOrder,
		stage.StandardLevelColumnTimeSlicePropertyTypes_reference,
		&stage.StandardLevelColumnTimeSlicePropertyTypes_referenceOrder,
		stage.StandardLevelColumnTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StandardLevelColumnTimeSliceTypes,
		stage.StandardLevelColumnTimeSliceType_stagedOrder,
		stage.StandardLevelColumnTimeSliceTypes_reference,
		&stage.StandardLevelColumnTimeSliceTypes_referenceOrder,
		stage.StandardLevelColumnTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StandardLevelColumnTypes,
		stage.StandardLevelColumnType_stagedOrder,
		stage.StandardLevelColumnTypes_reference,
		&stage.StandardLevelColumnTypes_referenceOrder,
		stage.StandardLevelColumnTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StandardLevelPropertyTypes,
		stage.StandardLevelPropertyType_stagedOrder,
		stage.StandardLevelPropertyTypes_reference,
		&stage.StandardLevelPropertyTypes_referenceOrder,
		stage.StandardLevelPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StandardLevelSectorTimeSlicePropertyTypes,
		stage.StandardLevelSectorTimeSlicePropertyType_stagedOrder,
		stage.StandardLevelSectorTimeSlicePropertyTypes_reference,
		&stage.StandardLevelSectorTimeSlicePropertyTypes_referenceOrder,
		stage.StandardLevelSectorTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StandardLevelSectorTimeSliceTypes,
		stage.StandardLevelSectorTimeSliceType_stagedOrder,
		stage.StandardLevelSectorTimeSliceTypes_reference,
		&stage.StandardLevelSectorTimeSliceTypes_referenceOrder,
		stage.StandardLevelSectorTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StandardLevelSectorTypes,
		stage.StandardLevelSectorType_stagedOrder,
		stage.StandardLevelSectorTypes_reference,
		&stage.StandardLevelSectorTypes_referenceOrder,
		stage.StandardLevelSectorTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StandardLevelTablePropertyTypes,
		stage.StandardLevelTablePropertyType_stagedOrder,
		stage.StandardLevelTablePropertyTypes_reference,
		&stage.StandardLevelTablePropertyTypes_referenceOrder,
		stage.StandardLevelTablePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StandardLevelTableTimeSlicePropertyTypes,
		stage.StandardLevelTableTimeSlicePropertyType_stagedOrder,
		stage.StandardLevelTableTimeSlicePropertyTypes_reference,
		&stage.StandardLevelTableTimeSlicePropertyTypes_referenceOrder,
		stage.StandardLevelTableTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StandardLevelTableTimeSliceTypes,
		stage.StandardLevelTableTimeSliceType_stagedOrder,
		stage.StandardLevelTableTimeSliceTypes_reference,
		&stage.StandardLevelTableTimeSliceTypes_referenceOrder,
		stage.StandardLevelTableTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StandardLevelTableTypes,
		stage.StandardLevelTableType_stagedOrder,
		stage.StandardLevelTableTypes_reference,
		&stage.StandardLevelTableTypes_referenceOrder,
		stage.StandardLevelTableTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StandardLevelTypes,
		stage.StandardLevelType_stagedOrder,
		stage.StandardLevelTypes_reference,
		&stage.StandardLevelTypes_referenceOrder,
		stage.StandardLevelTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
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
		stage.UnitAvailabilityPropertyTypes,
		stage.UnitAvailabilityPropertyType_stagedOrder,
		stage.UnitAvailabilityPropertyTypes_reference,
		&stage.UnitAvailabilityPropertyTypes_referenceOrder,
		stage.UnitAvailabilityPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.UnitAvailabilityTypes,
		stage.UnitAvailabilityType_stagedOrder,
		stage.UnitAvailabilityTypes_reference,
		&stage.UnitAvailabilityTypes_referenceOrder,
		stage.UnitAvailabilityTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.UnitDependencyPropertyTypes,
		stage.UnitDependencyPropertyType_stagedOrder,
		stage.UnitDependencyPropertyTypes_reference,
		&stage.UnitDependencyPropertyTypes_referenceOrder,
		stage.UnitDependencyPropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.UnitDependencyTypes,
		stage.UnitDependencyType_stagedOrder,
		stage.UnitDependencyTypes_reference,
		&stage.UnitDependencyTypes_referenceOrder,
		stage.UnitDependencyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
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
		stage.UnitTimeSlicePropertyTypes,
		stage.UnitTimeSlicePropertyType_stagedOrder,
		stage.UnitTimeSlicePropertyTypes_reference,
		&stage.UnitTimeSlicePropertyTypes_referenceOrder,
		stage.UnitTimeSlicePropertyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.UnitTimeSliceTypes,
		stage.UnitTimeSliceType_stagedOrder,
		stage.UnitTimeSliceTypes_reference,
		&stage.UnitTimeSliceTypes_referenceOrder,
		stage.UnitTimeSliceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.UnitTypes,
		stage.UnitType_stagedOrder,
		stage.UnitTypes_reference,
		&stage.UnitTypes_referenceOrder,
		stage.UnitTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
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

	__gong__computeReferencePass1(stage, stage.A_extension_4s, &stage.A_extension_4s_reference, &stage.A_extension_4s_referenceOrder, &stage.A_extension_4s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_5s, &stage.A_extension_5s_reference, &stage.A_extension_5s_referenceOrder, &stage.A_extension_5s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_6s, &stage.A_extension_6s_reference, &stage.A_extension_6s_referenceOrder, &stage.A_extension_6s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_7s, &stage.A_extension_7s_reference, &stage.A_extension_7s_referenceOrder, &stage.A_extension_7s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_8s, &stage.A_extension_8s_reference, &stage.A_extension_8s_referenceOrder, &stage.A_extension_8s_instance)

	__gong__computeReferencePass1(stage, stage.A_extension_9s, &stage.A_extension_9s_reference, &stage.A_extension_9s_referenceOrder, &stage.A_extension_9s_instance)

	__gong__computeReferencePass1(stage, stage.AbstractAirportGroundServiceTypes, &stage.AbstractAirportGroundServiceTypes_reference, &stage.AbstractAirportGroundServiceTypes_referenceOrder, &stage.AbstractAirportGroundServiceTypes_instance)

	__gong__computeReferencePass1(stage, stage.AbstractPropertiesWithScheduleTypes, &stage.AbstractPropertiesWithScheduleTypes_reference, &stage.AbstractPropertiesWithScheduleTypes_referenceOrder, &stage.AbstractPropertiesWithScheduleTypes_instance)

	__gong__computeReferencePass1(stage, stage.AbstractServiceTypes, &stage.AbstractServiceTypes_reference, &stage.AbstractServiceTypes_referenceOrder, &stage.AbstractServiceTypes_instance)

	__gong__computeReferencePass1(stage, stage.AbstractTrafficSeparationServiceTypes, &stage.AbstractTrafficSeparationServiceTypes_reference, &stage.AbstractTrafficSeparationServiceTypes_referenceOrder, &stage.AbstractTrafficSeparationServiceTypes_instance)

	__gong__computeReferencePass1(stage, stage.AerialRefuellingPropertyTypes, &stage.AerialRefuellingPropertyTypes_reference, &stage.AerialRefuellingPropertyTypes_referenceOrder, &stage.AerialRefuellingPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirTrafficControlServiceTimeSlicePropertyTypes, &stage.AirTrafficControlServiceTimeSlicePropertyTypes_reference, &stage.AirTrafficControlServiceTimeSlicePropertyTypes_referenceOrder, &stage.AirTrafficControlServiceTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirTrafficControlServiceTimeSliceTypes, &stage.AirTrafficControlServiceTimeSliceTypes_reference, &stage.AirTrafficControlServiceTimeSliceTypes_referenceOrder, &stage.AirTrafficControlServiceTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirTrafficControlServiceTypes, &stage.AirTrafficControlServiceTypes_reference, &stage.AirTrafficControlServiceTypes_referenceOrder, &stage.AirTrafficControlServiceTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirTrafficManagementServiceTimeSlicePropertyTypes, &stage.AirTrafficManagementServiceTimeSlicePropertyTypes_reference, &stage.AirTrafficManagementServiceTimeSlicePropertyTypes_referenceOrder, &stage.AirTrafficManagementServiceTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirTrafficManagementServiceTimeSliceTypes, &stage.AirTrafficManagementServiceTimeSliceTypes_reference, &stage.AirTrafficManagementServiceTimeSliceTypes_referenceOrder, &stage.AirTrafficManagementServiceTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirTrafficManagementServiceTypes, &stage.AirTrafficManagementServiceTypes_reference, &stage.AirTrafficManagementServiceTypes_referenceOrder, &stage.AirTrafficManagementServiceTypes_instance)

	__gong__computeReferencePass1(stage, stage.AircraftGroundServiceTimeSlicePropertyTypes, &stage.AircraftGroundServiceTimeSlicePropertyTypes_reference, &stage.AircraftGroundServiceTimeSlicePropertyTypes_referenceOrder, &stage.AircraftGroundServiceTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AircraftGroundServiceTimeSliceTypes, &stage.AircraftGroundServiceTimeSliceTypes_reference, &stage.AircraftGroundServiceTimeSliceTypes_referenceOrder, &stage.AircraftGroundServiceTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.AircraftGroundServiceTypes, &stage.AircraftGroundServiceTypes_reference, &stage.AircraftGroundServiceTypes_referenceOrder, &stage.AircraftGroundServiceTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirportClearanceServiceTimeSlicePropertyTypes, &stage.AirportClearanceServiceTimeSlicePropertyTypes_reference, &stage.AirportClearanceServiceTimeSlicePropertyTypes_referenceOrder, &stage.AirportClearanceServiceTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirportClearanceServiceTimeSliceTypes, &stage.AirportClearanceServiceTimeSliceTypes_reference, &stage.AirportClearanceServiceTimeSliceTypes_referenceOrder, &stage.AirportClearanceServiceTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirportClearanceServiceTypes, &stage.AirportClearanceServiceTypes_reference, &stage.AirportClearanceServiceTypes_referenceOrder, &stage.AirportClearanceServiceTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirportHeliportPropertyTypes, &stage.AirportHeliportPropertyTypes_reference, &stage.AirportHeliportPropertyTypes_referenceOrder, &stage.AirportHeliportPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirportSuppliesServiceTimeSlicePropertyTypes, &stage.AirportSuppliesServiceTimeSlicePropertyTypes_reference, &stage.AirportSuppliesServiceTimeSlicePropertyTypes_referenceOrder, &stage.AirportSuppliesServiceTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirportSuppliesServiceTimeSliceTypes, &stage.AirportSuppliesServiceTimeSliceTypes_reference, &stage.AirportSuppliesServiceTimeSliceTypes_referenceOrder, &stage.AirportSuppliesServiceTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirportSuppliesServiceTypes, &stage.AirportSuppliesServiceTypes_reference, &stage.AirportSuppliesServiceTypes_referenceOrder, &stage.AirportSuppliesServiceTypes_instance)

	__gong__computeReferencePass1(stage, stage.AirspacePropertyTypes, &stage.AirspacePropertyTypes_reference, &stage.AirspacePropertyTypes_referenceOrder, &stage.AirspacePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.CallsignDetailPropertyTypes, &stage.CallsignDetailPropertyTypes_reference, &stage.CallsignDetailPropertyTypes_referenceOrder, &stage.CallsignDetailPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.CallsignDetailTypes, &stage.CallsignDetailTypes_reference, &stage.CallsignDetailTypes_referenceOrder, &stage.CallsignDetailTypes_instance)

	__gong__computeReferencePass1(stage, stage.ContactInformationPropertyTypes, &stage.ContactInformationPropertyTypes_reference, &stage.ContactInformationPropertyTypes_referenceOrder, &stage.ContactInformationPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ContactInformationTypes, &stage.ContactInformationTypes_reference, &stage.ContactInformationTypes_referenceOrder, &stage.ContactInformationTypes_instance)

	__gong__computeReferencePass1(stage, stage.DesignatedPointPropertyTypes, &stage.DesignatedPointPropertyTypes_reference, &stage.DesignatedPointPropertyTypes_referenceOrder, &stage.DesignatedPointPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.DirectionFinderPropertyTypes, &stage.DirectionFinderPropertyTypes_reference, &stage.DirectionFinderPropertyTypes_referenceOrder, &stage.DirectionFinderPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ElevatedPointPropertyTypes, &stage.ElevatedPointPropertyTypes_reference, &stage.ElevatedPointPropertyTypes_referenceOrder, &stage.ElevatedPointPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ElevatedPointTypes, &stage.ElevatedPointTypes_reference, &stage.ElevatedPointTypes_referenceOrder, &stage.ElevatedPointTypes_instance)

	__gong__computeReferencePass1(stage, stage.FireFightingServiceTimeSlicePropertyTypes, &stage.FireFightingServiceTimeSlicePropertyTypes_reference, &stage.FireFightingServiceTimeSlicePropertyTypes_referenceOrder, &stage.FireFightingServiceTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.FireFightingServiceTimeSliceTypes, &stage.FireFightingServiceTimeSliceTypes_reference, &stage.FireFightingServiceTimeSliceTypes_referenceOrder, &stage.FireFightingServiceTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.FireFightingServiceTypes, &stage.FireFightingServiceTypes_reference, &stage.FireFightingServiceTypes_referenceOrder, &stage.FireFightingServiceTypes_instance)

	__gong__computeReferencePass1(stage, stage.FuelPropertyTypes, &stage.FuelPropertyTypes_reference, &stage.FuelPropertyTypes_referenceOrder, &stage.FuelPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.FuelTypes, &stage.FuelTypes_reference, &stage.FuelTypes_referenceOrder, &stage.FuelTypes_instance)

	__gong__computeReferencePass1(stage, stage.GroundLightSystemPropertyTypes, &stage.GroundLightSystemPropertyTypes_reference, &stage.GroundLightSystemPropertyTypes_referenceOrder, &stage.GroundLightSystemPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.GroundTrafficControlServiceTimeSlicePropertyTypes, &stage.GroundTrafficControlServiceTimeSlicePropertyTypes_reference, &stage.GroundTrafficControlServiceTimeSlicePropertyTypes_referenceOrder, &stage.GroundTrafficControlServiceTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.GroundTrafficControlServiceTimeSliceTypes, &stage.GroundTrafficControlServiceTimeSliceTypes_reference, &stage.GroundTrafficControlServiceTimeSliceTypes_referenceOrder, &stage.GroundTrafficControlServiceTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.GroundTrafficControlServiceTypes, &stage.GroundTrafficControlServiceTypes_reference, &stage.GroundTrafficControlServiceTypes_referenceOrder, &stage.GroundTrafficControlServiceTypes_instance)

	__gong__computeReferencePass1(stage, stage.HoldingPatternPropertyTypes, &stage.HoldingPatternPropertyTypes_reference, &stage.HoldingPatternPropertyTypes_referenceOrder, &stage.HoldingPatternPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.InformationServiceTimeSlicePropertyTypes, &stage.InformationServiceTimeSlicePropertyTypes_reference, &stage.InformationServiceTimeSlicePropertyTypes_referenceOrder, &stage.InformationServiceTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.InformationServiceTimeSliceTypes, &stage.InformationServiceTimeSliceTypes_reference, &stage.InformationServiceTimeSliceTypes_referenceOrder, &stage.InformationServiceTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.InformationServiceTypes, &stage.InformationServiceTypes_reference, &stage.InformationServiceTypes_referenceOrder, &stage.InformationServiceTypes_instance)

	__gong__computeReferencePass1(stage, stage.LightActivationPropertyTypes, &stage.LightActivationPropertyTypes_reference, &stage.LightActivationPropertyTypes_referenceOrder, &stage.LightActivationPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.LightActivationTypes, &stage.LightActivationTypes_reference, &stage.LightActivationTypes_referenceOrder, &stage.LightActivationTypes_instance)

	__gong__computeReferencePass1(stage, stage.LinguisticNotePropertyTypes, &stage.LinguisticNotePropertyTypes_reference, &stage.LinguisticNotePropertyTypes_referenceOrder, &stage.LinguisticNotePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.LinguisticNoteTypes, &stage.LinguisticNoteTypes_reference, &stage.LinguisticNoteTypes_referenceOrder, &stage.LinguisticNoteTypes_instance)

	__gong__computeReferencePass1(stage, stage.NavaidPropertyTypes, &stage.NavaidPropertyTypes_reference, &stage.NavaidPropertyTypes_referenceOrder, &stage.NavaidPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.NitrogenPropertyTypes, &stage.NitrogenPropertyTypes_reference, &stage.NitrogenPropertyTypes_referenceOrder, &stage.NitrogenPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.NitrogenTypes, &stage.NitrogenTypes_reference, &stage.NitrogenTypes_referenceOrder, &stage.NitrogenTypes_instance)

	__gong__computeReferencePass1(stage, stage.NotePropertyTypes, &stage.NotePropertyTypes_reference, &stage.NotePropertyTypes_referenceOrder, &stage.NotePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.NoteTypes, &stage.NoteTypes_reference, &stage.NoteTypes_referenceOrder, &stage.NoteTypes_instance)

	__gong__computeReferencePass1(stage, stage.OilPropertyTypes, &stage.OilPropertyTypes_reference, &stage.OilPropertyTypes_referenceOrder, &stage.OilPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.OilTypes, &stage.OilTypes_reference, &stage.OilTypes_referenceOrder, &stage.OilTypes_instance)

	__gong__computeReferencePass1(stage, stage.OnlineContactPropertyTypes, &stage.OnlineContactPropertyTypes_reference, &stage.OnlineContactPropertyTypes_referenceOrder, &stage.OnlineContactPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.OnlineContactTypes, &stage.OnlineContactTypes_reference, &stage.OnlineContactTypes_referenceOrder, &stage.OnlineContactTypes_instance)

	__gong__computeReferencePass1(stage, stage.OrganisationAuthorityAssociationPropertyTypes, &stage.OrganisationAuthorityAssociationPropertyTypes_reference, &stage.OrganisationAuthorityAssociationPropertyTypes_referenceOrder, &stage.OrganisationAuthorityAssociationPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.OrganisationAuthorityAssociationTypes, &stage.OrganisationAuthorityAssociationTypes_reference, &stage.OrganisationAuthorityAssociationTypes_referenceOrder, &stage.OrganisationAuthorityAssociationTypes_instance)

	__gong__computeReferencePass1(stage, stage.OrganisationAuthorityPropertyTypes, &stage.OrganisationAuthorityPropertyTypes_reference, &stage.OrganisationAuthorityPropertyTypes_referenceOrder, &stage.OrganisationAuthorityPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.OrganisationAuthorityTimeSlicePropertyTypes, &stage.OrganisationAuthorityTimeSlicePropertyTypes_reference, &stage.OrganisationAuthorityTimeSlicePropertyTypes_referenceOrder, &stage.OrganisationAuthorityTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.OrganisationAuthorityTimeSliceTypes, &stage.OrganisationAuthorityTimeSliceTypes_reference, &stage.OrganisationAuthorityTimeSliceTypes_referenceOrder, &stage.OrganisationAuthorityTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.OrganisationAuthorityTypes, &stage.OrganisationAuthorityTypes_reference, &stage.OrganisationAuthorityTypes_referenceOrder, &stage.OrganisationAuthorityTypes_instance)

	__gong__computeReferencePass1(stage, stage.OxygenPropertyTypes, &stage.OxygenPropertyTypes_reference, &stage.OxygenPropertyTypes_referenceOrder, &stage.OxygenPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.OxygenTypes, &stage.OxygenTypes_reference, &stage.OxygenTypes_referenceOrder, &stage.OxygenTypes_instance)

	__gong__computeReferencePass1(stage, stage.PassengerServiceTimeSlicePropertyTypes, &stage.PassengerServiceTimeSlicePropertyTypes_reference, &stage.PassengerServiceTimeSlicePropertyTypes_referenceOrder, &stage.PassengerServiceTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.PassengerServiceTimeSliceTypes, &stage.PassengerServiceTimeSliceTypes_reference, &stage.PassengerServiceTimeSliceTypes_referenceOrder, &stage.PassengerServiceTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.PassengerServiceTypes, &stage.PassengerServiceTypes_reference, &stage.PassengerServiceTypes_referenceOrder, &stage.PassengerServiceTypes_instance)

	__gong__computeReferencePass1(stage, stage.PilotControlledLightingTimeSlicePropertyTypes, &stage.PilotControlledLightingTimeSlicePropertyTypes_reference, &stage.PilotControlledLightingTimeSlicePropertyTypes_referenceOrder, &stage.PilotControlledLightingTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.PilotControlledLightingTimeSliceTypes, &stage.PilotControlledLightingTimeSliceTypes_reference, &stage.PilotControlledLightingTimeSliceTypes_referenceOrder, &stage.PilotControlledLightingTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.PilotControlledLightingTypes, &stage.PilotControlledLightingTypes_reference, &stage.PilotControlledLightingTypes_referenceOrder, &stage.PilotControlledLightingTypes_instance)

	__gong__computeReferencePass1(stage, stage.PointPropertyTypes, &stage.PointPropertyTypes_reference, &stage.PointPropertyTypes_referenceOrder, &stage.PointPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.PointTypes, &stage.PointTypes_reference, &stage.PointTypes_referenceOrder, &stage.PointTypes_instance)

	__gong__computeReferencePass1(stage, stage.PostalAddressPropertyTypes, &stage.PostalAddressPropertyTypes_reference, &stage.PostalAddressPropertyTypes_referenceOrder, &stage.PostalAddressPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.PostalAddressTypes, &stage.PostalAddressTypes_reference, &stage.PostalAddressTypes_referenceOrder, &stage.PostalAddressTypes_instance)

	__gong__computeReferencePass1(stage, stage.ProcedurePropertyTypes, &stage.ProcedurePropertyTypes_reference, &stage.ProcedurePropertyTypes_referenceOrder, &stage.ProcedurePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.RadioCommunicationChannelPropertyTypes, &stage.RadioCommunicationChannelPropertyTypes_reference, &stage.RadioCommunicationChannelPropertyTypes_referenceOrder, &stage.RadioCommunicationChannelPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.RadioCommunicationChannelTimeSlicePropertyTypes, &stage.RadioCommunicationChannelTimeSlicePropertyTypes_reference, &stage.RadioCommunicationChannelTimeSlicePropertyTypes_referenceOrder, &stage.RadioCommunicationChannelTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.RadioCommunicationChannelTimeSliceTypes, &stage.RadioCommunicationChannelTimeSliceTypes_reference, &stage.RadioCommunicationChannelTimeSliceTypes_referenceOrder, &stage.RadioCommunicationChannelTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.RadioCommunicationChannelTypes, &stage.RadioCommunicationChannelTypes_reference, &stage.RadioCommunicationChannelTypes_referenceOrder, &stage.RadioCommunicationChannelTypes_instance)

	__gong__computeReferencePass1(stage, stage.RadioCommunicationOperationalStatusPropertyTypes, &stage.RadioCommunicationOperationalStatusPropertyTypes_reference, &stage.RadioCommunicationOperationalStatusPropertyTypes_referenceOrder, &stage.RadioCommunicationOperationalStatusPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.RadioCommunicationOperationalStatusTypes, &stage.RadioCommunicationOperationalStatusTypes_reference, &stage.RadioCommunicationOperationalStatusTypes_referenceOrder, &stage.RadioCommunicationOperationalStatusTypes_instance)

	__gong__computeReferencePass1(stage, stage.RoutePortionPropertyTypes, &stage.RoutePortionPropertyTypes_reference, &stage.RoutePortionPropertyTypes_referenceOrder, &stage.RoutePortionPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.RoutePortionTypes, &stage.RoutePortionTypes_reference, &stage.RoutePortionTypes_referenceOrder, &stage.RoutePortionTypes_instance)

	__gong__computeReferencePass1(stage, stage.RoutePropertyTypes, &stage.RoutePropertyTypes_reference, &stage.RoutePropertyTypes_referenceOrder, &stage.RoutePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.RunwayCentrelinePointPropertyTypes, &stage.RunwayCentrelinePointPropertyTypes_reference, &stage.RunwayCentrelinePointPropertyTypes_referenceOrder, &stage.RunwayCentrelinePointPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.SearchRescueServiceTimeSlicePropertyTypes, &stage.SearchRescueServiceTimeSlicePropertyTypes_reference, &stage.SearchRescueServiceTimeSlicePropertyTypes_referenceOrder, &stage.SearchRescueServiceTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.SearchRescueServiceTimeSliceTypes, &stage.SearchRescueServiceTimeSliceTypes_reference, &stage.SearchRescueServiceTimeSliceTypes_referenceOrder, &stage.SearchRescueServiceTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.SearchRescueServiceTypes, &stage.SearchRescueServiceTypes_reference, &stage.SearchRescueServiceTypes_referenceOrder, &stage.SearchRescueServiceTypes_instance)

	__gong__computeReferencePass1(stage, stage.ServiceOperationalStatusPropertyTypes, &stage.ServiceOperationalStatusPropertyTypes_reference, &stage.ServiceOperationalStatusPropertyTypes_referenceOrder, &stage.ServiceOperationalStatusPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ServiceOperationalStatusTypes, &stage.ServiceOperationalStatusTypes_reference, &stage.ServiceOperationalStatusTypes_referenceOrder, &stage.ServiceOperationalStatusTypes_instance)

	__gong__computeReferencePass1(stage, stage.SpecialDateTimeSlicePropertyTypes, &stage.SpecialDateTimeSlicePropertyTypes_reference, &stage.SpecialDateTimeSlicePropertyTypes_referenceOrder, &stage.SpecialDateTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.SpecialDateTimeSliceTypes, &stage.SpecialDateTimeSliceTypes_reference, &stage.SpecialDateTimeSliceTypes_referenceOrder, &stage.SpecialDateTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.SpecialDateTypes, &stage.SpecialDateTypes_reference, &stage.SpecialDateTypes_referenceOrder, &stage.SpecialDateTypes_instance)

	__gong__computeReferencePass1(stage, stage.StandardLevelColumnPropertyTypes, &stage.StandardLevelColumnPropertyTypes_reference, &stage.StandardLevelColumnPropertyTypes_referenceOrder, &stage.StandardLevelColumnPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.StandardLevelColumnTimeSlicePropertyTypes, &stage.StandardLevelColumnTimeSlicePropertyTypes_reference, &stage.StandardLevelColumnTimeSlicePropertyTypes_referenceOrder, &stage.StandardLevelColumnTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.StandardLevelColumnTimeSliceTypes, &stage.StandardLevelColumnTimeSliceTypes_reference, &stage.StandardLevelColumnTimeSliceTypes_referenceOrder, &stage.StandardLevelColumnTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.StandardLevelColumnTypes, &stage.StandardLevelColumnTypes_reference, &stage.StandardLevelColumnTypes_referenceOrder, &stage.StandardLevelColumnTypes_instance)

	__gong__computeReferencePass1(stage, stage.StandardLevelPropertyTypes, &stage.StandardLevelPropertyTypes_reference, &stage.StandardLevelPropertyTypes_referenceOrder, &stage.StandardLevelPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.StandardLevelSectorTimeSlicePropertyTypes, &stage.StandardLevelSectorTimeSlicePropertyTypes_reference, &stage.StandardLevelSectorTimeSlicePropertyTypes_referenceOrder, &stage.StandardLevelSectorTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.StandardLevelSectorTimeSliceTypes, &stage.StandardLevelSectorTimeSliceTypes_reference, &stage.StandardLevelSectorTimeSliceTypes_referenceOrder, &stage.StandardLevelSectorTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.StandardLevelSectorTypes, &stage.StandardLevelSectorTypes_reference, &stage.StandardLevelSectorTypes_referenceOrder, &stage.StandardLevelSectorTypes_instance)

	__gong__computeReferencePass1(stage, stage.StandardLevelTablePropertyTypes, &stage.StandardLevelTablePropertyTypes_reference, &stage.StandardLevelTablePropertyTypes_referenceOrder, &stage.StandardLevelTablePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.StandardLevelTableTimeSlicePropertyTypes, &stage.StandardLevelTableTimeSlicePropertyTypes_reference, &stage.StandardLevelTableTimeSlicePropertyTypes_referenceOrder, &stage.StandardLevelTableTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.StandardLevelTableTimeSliceTypes, &stage.StandardLevelTableTimeSliceTypes_reference, &stage.StandardLevelTableTimeSliceTypes_referenceOrder, &stage.StandardLevelTableTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.StandardLevelTableTypes, &stage.StandardLevelTableTypes_reference, &stage.StandardLevelTableTypes_referenceOrder, &stage.StandardLevelTableTypes_instance)

	__gong__computeReferencePass1(stage, stage.StandardLevelTypes, &stage.StandardLevelTypes_reference, &stage.StandardLevelTypes_referenceOrder, &stage.StandardLevelTypes_instance)

	__gong__computeReferencePass1(stage, stage.TelephoneContactPropertyTypes, &stage.TelephoneContactPropertyTypes_reference, &stage.TelephoneContactPropertyTypes_referenceOrder, &stage.TelephoneContactPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.TelephoneContactTypes, &stage.TelephoneContactTypes_reference, &stage.TelephoneContactTypes_referenceOrder, &stage.TelephoneContactTypes_instance)

	__gong__computeReferencePass1(stage, stage.TimesheetPropertyTypes, &stage.TimesheetPropertyTypes_reference, &stage.TimesheetPropertyTypes_referenceOrder, &stage.TimesheetPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.TimesheetTypes, &stage.TimesheetTypes_reference, &stage.TimesheetTypes_referenceOrder, &stage.TimesheetTypes_instance)

	__gong__computeReferencePass1(stage, stage.TouchDownLiftOffPropertyTypes, &stage.TouchDownLiftOffPropertyTypes_reference, &stage.TouchDownLiftOffPropertyTypes_referenceOrder, &stage.TouchDownLiftOffPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.UnitAvailabilityPropertyTypes, &stage.UnitAvailabilityPropertyTypes_reference, &stage.UnitAvailabilityPropertyTypes_referenceOrder, &stage.UnitAvailabilityPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.UnitAvailabilityTypes, &stage.UnitAvailabilityTypes_reference, &stage.UnitAvailabilityTypes_referenceOrder, &stage.UnitAvailabilityTypes_instance)

	__gong__computeReferencePass1(stage, stage.UnitDependencyPropertyTypes, &stage.UnitDependencyPropertyTypes_reference, &stage.UnitDependencyPropertyTypes_referenceOrder, &stage.UnitDependencyPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.UnitDependencyTypes, &stage.UnitDependencyTypes_reference, &stage.UnitDependencyTypes_referenceOrder, &stage.UnitDependencyTypes_instance)

	__gong__computeReferencePass1(stage, stage.UnitPropertyTypes, &stage.UnitPropertyTypes_reference, &stage.UnitPropertyTypes_referenceOrder, &stage.UnitPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.UnitTimeSlicePropertyTypes, &stage.UnitTimeSlicePropertyTypes_reference, &stage.UnitTimeSlicePropertyTypes_referenceOrder, &stage.UnitTimeSlicePropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.UnitTimeSliceTypes, &stage.UnitTimeSliceTypes_reference, &stage.UnitTimeSliceTypes_referenceOrder, &stage.UnitTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.UnitTypes, &stage.UnitTypes_reference, &stage.UnitTypes_referenceOrder, &stage.UnitTypes_instance)

	__gong__computeReferencePass1(stage, stage.VORPropertyTypes, &stage.VORPropertyTypes_reference, &stage.VORPropertyTypes_referenceOrder, &stage.VORPropertyTypes_instance)

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

	__gong__computeReferencePass2(stage.A_extension_4s, stage.A_extension_4s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_5s, stage.A_extension_5s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_6s, stage.A_extension_6s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_7s, stage.A_extension_7s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_8s, stage.A_extension_8s_reference, stage)

	__gong__computeReferencePass2(stage.A_extension_9s, stage.A_extension_9s_reference, stage)

	__gong__computeReferencePass2(stage.AbstractAirportGroundServiceTypes, stage.AbstractAirportGroundServiceTypes_reference, stage)

	__gong__computeReferencePass2(stage.AbstractPropertiesWithScheduleTypes, stage.AbstractPropertiesWithScheduleTypes_reference, stage)

	__gong__computeReferencePass2(stage.AbstractServiceTypes, stage.AbstractServiceTypes_reference, stage)

	__gong__computeReferencePass2(stage.AbstractTrafficSeparationServiceTypes, stage.AbstractTrafficSeparationServiceTypes_reference, stage)

	__gong__computeReferencePass2(stage.AerialRefuellingPropertyTypes, stage.AerialRefuellingPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirTrafficControlServiceTimeSlicePropertyTypes, stage.AirTrafficControlServiceTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirTrafficControlServiceTimeSliceTypes, stage.AirTrafficControlServiceTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirTrafficControlServiceTypes, stage.AirTrafficControlServiceTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirTrafficManagementServiceTimeSlicePropertyTypes, stage.AirTrafficManagementServiceTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirTrafficManagementServiceTimeSliceTypes, stage.AirTrafficManagementServiceTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirTrafficManagementServiceTypes, stage.AirTrafficManagementServiceTypes_reference, stage)

	__gong__computeReferencePass2(stage.AircraftGroundServiceTimeSlicePropertyTypes, stage.AircraftGroundServiceTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AircraftGroundServiceTimeSliceTypes, stage.AircraftGroundServiceTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.AircraftGroundServiceTypes, stage.AircraftGroundServiceTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirportClearanceServiceTimeSlicePropertyTypes, stage.AirportClearanceServiceTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirportClearanceServiceTimeSliceTypes, stage.AirportClearanceServiceTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirportClearanceServiceTypes, stage.AirportClearanceServiceTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirportHeliportPropertyTypes, stage.AirportHeliportPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirportSuppliesServiceTimeSlicePropertyTypes, stage.AirportSuppliesServiceTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirportSuppliesServiceTimeSliceTypes, stage.AirportSuppliesServiceTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirportSuppliesServiceTypes, stage.AirportSuppliesServiceTypes_reference, stage)

	__gong__computeReferencePass2(stage.AirspacePropertyTypes, stage.AirspacePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.CallsignDetailPropertyTypes, stage.CallsignDetailPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.CallsignDetailTypes, stage.CallsignDetailTypes_reference, stage)

	__gong__computeReferencePass2(stage.ContactInformationPropertyTypes, stage.ContactInformationPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ContactInformationTypes, stage.ContactInformationTypes_reference, stage)

	__gong__computeReferencePass2(stage.DesignatedPointPropertyTypes, stage.DesignatedPointPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.DirectionFinderPropertyTypes, stage.DirectionFinderPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ElevatedPointPropertyTypes, stage.ElevatedPointPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ElevatedPointTypes, stage.ElevatedPointTypes_reference, stage)

	__gong__computeReferencePass2(stage.FireFightingServiceTimeSlicePropertyTypes, stage.FireFightingServiceTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.FireFightingServiceTimeSliceTypes, stage.FireFightingServiceTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.FireFightingServiceTypes, stage.FireFightingServiceTypes_reference, stage)

	__gong__computeReferencePass2(stage.FuelPropertyTypes, stage.FuelPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.FuelTypes, stage.FuelTypes_reference, stage)

	__gong__computeReferencePass2(stage.GroundLightSystemPropertyTypes, stage.GroundLightSystemPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.GroundTrafficControlServiceTimeSlicePropertyTypes, stage.GroundTrafficControlServiceTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.GroundTrafficControlServiceTimeSliceTypes, stage.GroundTrafficControlServiceTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.GroundTrafficControlServiceTypes, stage.GroundTrafficControlServiceTypes_reference, stage)

	__gong__computeReferencePass2(stage.HoldingPatternPropertyTypes, stage.HoldingPatternPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.InformationServiceTimeSlicePropertyTypes, stage.InformationServiceTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.InformationServiceTimeSliceTypes, stage.InformationServiceTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.InformationServiceTypes, stage.InformationServiceTypes_reference, stage)

	__gong__computeReferencePass2(stage.LightActivationPropertyTypes, stage.LightActivationPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.LightActivationTypes, stage.LightActivationTypes_reference, stage)

	__gong__computeReferencePass2(stage.LinguisticNotePropertyTypes, stage.LinguisticNotePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.LinguisticNoteTypes, stage.LinguisticNoteTypes_reference, stage)

	__gong__computeReferencePass2(stage.NavaidPropertyTypes, stage.NavaidPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.NitrogenPropertyTypes, stage.NitrogenPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.NitrogenTypes, stage.NitrogenTypes_reference, stage)

	__gong__computeReferencePass2(stage.NotePropertyTypes, stage.NotePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.NoteTypes, stage.NoteTypes_reference, stage)

	__gong__computeReferencePass2(stage.OilPropertyTypes, stage.OilPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.OilTypes, stage.OilTypes_reference, stage)

	__gong__computeReferencePass2(stage.OnlineContactPropertyTypes, stage.OnlineContactPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.OnlineContactTypes, stage.OnlineContactTypes_reference, stage)

	__gong__computeReferencePass2(stage.OrganisationAuthorityAssociationPropertyTypes, stage.OrganisationAuthorityAssociationPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.OrganisationAuthorityAssociationTypes, stage.OrganisationAuthorityAssociationTypes_reference, stage)

	__gong__computeReferencePass2(stage.OrganisationAuthorityPropertyTypes, stage.OrganisationAuthorityPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.OrganisationAuthorityTimeSlicePropertyTypes, stage.OrganisationAuthorityTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.OrganisationAuthorityTimeSliceTypes, stage.OrganisationAuthorityTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.OrganisationAuthorityTypes, stage.OrganisationAuthorityTypes_reference, stage)

	__gong__computeReferencePass2(stage.OxygenPropertyTypes, stage.OxygenPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.OxygenTypes, stage.OxygenTypes_reference, stage)

	__gong__computeReferencePass2(stage.PassengerServiceTimeSlicePropertyTypes, stage.PassengerServiceTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.PassengerServiceTimeSliceTypes, stage.PassengerServiceTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.PassengerServiceTypes, stage.PassengerServiceTypes_reference, stage)

	__gong__computeReferencePass2(stage.PilotControlledLightingTimeSlicePropertyTypes, stage.PilotControlledLightingTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.PilotControlledLightingTimeSliceTypes, stage.PilotControlledLightingTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.PilotControlledLightingTypes, stage.PilotControlledLightingTypes_reference, stage)

	__gong__computeReferencePass2(stage.PointPropertyTypes, stage.PointPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.PointTypes, stage.PointTypes_reference, stage)

	__gong__computeReferencePass2(stage.PostalAddressPropertyTypes, stage.PostalAddressPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.PostalAddressTypes, stage.PostalAddressTypes_reference, stage)

	__gong__computeReferencePass2(stage.ProcedurePropertyTypes, stage.ProcedurePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.RadioCommunicationChannelPropertyTypes, stage.RadioCommunicationChannelPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.RadioCommunicationChannelTimeSlicePropertyTypes, stage.RadioCommunicationChannelTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.RadioCommunicationChannelTimeSliceTypes, stage.RadioCommunicationChannelTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.RadioCommunicationChannelTypes, stage.RadioCommunicationChannelTypes_reference, stage)

	__gong__computeReferencePass2(stage.RadioCommunicationOperationalStatusPropertyTypes, stage.RadioCommunicationOperationalStatusPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.RadioCommunicationOperationalStatusTypes, stage.RadioCommunicationOperationalStatusTypes_reference, stage)

	__gong__computeReferencePass2(stage.RoutePortionPropertyTypes, stage.RoutePortionPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.RoutePortionTypes, stage.RoutePortionTypes_reference, stage)

	__gong__computeReferencePass2(stage.RoutePropertyTypes, stage.RoutePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.RunwayCentrelinePointPropertyTypes, stage.RunwayCentrelinePointPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.SearchRescueServiceTimeSlicePropertyTypes, stage.SearchRescueServiceTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.SearchRescueServiceTimeSliceTypes, stage.SearchRescueServiceTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.SearchRescueServiceTypes, stage.SearchRescueServiceTypes_reference, stage)

	__gong__computeReferencePass2(stage.ServiceOperationalStatusPropertyTypes, stage.ServiceOperationalStatusPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ServiceOperationalStatusTypes, stage.ServiceOperationalStatusTypes_reference, stage)

	__gong__computeReferencePass2(stage.SpecialDateTimeSlicePropertyTypes, stage.SpecialDateTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.SpecialDateTimeSliceTypes, stage.SpecialDateTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.SpecialDateTypes, stage.SpecialDateTypes_reference, stage)

	__gong__computeReferencePass2(stage.StandardLevelColumnPropertyTypes, stage.StandardLevelColumnPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.StandardLevelColumnTimeSlicePropertyTypes, stage.StandardLevelColumnTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.StandardLevelColumnTimeSliceTypes, stage.StandardLevelColumnTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.StandardLevelColumnTypes, stage.StandardLevelColumnTypes_reference, stage)

	__gong__computeReferencePass2(stage.StandardLevelPropertyTypes, stage.StandardLevelPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.StandardLevelSectorTimeSlicePropertyTypes, stage.StandardLevelSectorTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.StandardLevelSectorTimeSliceTypes, stage.StandardLevelSectorTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.StandardLevelSectorTypes, stage.StandardLevelSectorTypes_reference, stage)

	__gong__computeReferencePass2(stage.StandardLevelTablePropertyTypes, stage.StandardLevelTablePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.StandardLevelTableTimeSlicePropertyTypes, stage.StandardLevelTableTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.StandardLevelTableTimeSliceTypes, stage.StandardLevelTableTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.StandardLevelTableTypes, stage.StandardLevelTableTypes_reference, stage)

	__gong__computeReferencePass2(stage.StandardLevelTypes, stage.StandardLevelTypes_reference, stage)

	__gong__computeReferencePass2(stage.TelephoneContactPropertyTypes, stage.TelephoneContactPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.TelephoneContactTypes, stage.TelephoneContactTypes_reference, stage)

	__gong__computeReferencePass2(stage.TimesheetPropertyTypes, stage.TimesheetPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.TimesheetTypes, stage.TimesheetTypes_reference, stage)

	__gong__computeReferencePass2(stage.TouchDownLiftOffPropertyTypes, stage.TouchDownLiftOffPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.UnitAvailabilityPropertyTypes, stage.UnitAvailabilityPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.UnitAvailabilityTypes, stage.UnitAvailabilityTypes_reference, stage)

	__gong__computeReferencePass2(stage.UnitDependencyPropertyTypes, stage.UnitDependencyPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.UnitDependencyTypes, stage.UnitDependencyTypes_reference, stage)

	__gong__computeReferencePass2(stage.UnitPropertyTypes, stage.UnitPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.UnitTimeSlicePropertyTypes, stage.UnitTimeSlicePropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.UnitTimeSliceTypes, stage.UnitTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.UnitTypes, stage.UnitTypes_reference, stage)

	__gong__computeReferencePass2(stage.VORPropertyTypes, stage.VORPropertyTypes_reference, stage)

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

func (abstractairportgroundservicetype *AbstractAirportGroundServiceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractAirportGroundServiceType_stagedOrder, stage.AbstractAirportGroundServiceTypes_referenceOrder, abstractairportgroundservicetype, "AbstractAirportGroundServiceType")
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractPropertiesWithScheduleType_stagedOrder, stage.AbstractPropertiesWithScheduleTypes_referenceOrder, abstractpropertieswithscheduletype, "AbstractPropertiesWithScheduleType")
}

func (abstractservicetype *AbstractServiceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractServiceType_stagedOrder, stage.AbstractServiceTypes_referenceOrder, abstractservicetype, "AbstractServiceType")
}

func (abstracttrafficseparationservicetype *AbstractTrafficSeparationServiceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractTrafficSeparationServiceType_stagedOrder, stage.AbstractTrafficSeparationServiceTypes_referenceOrder, abstracttrafficseparationservicetype, "AbstractTrafficSeparationServiceType")
}

func (aerialrefuellingpropertytype *AerialRefuellingPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AerialRefuellingPropertyType_stagedOrder, stage.AerialRefuellingPropertyTypes_referenceOrder, aerialrefuellingpropertytype, "AerialRefuellingPropertyType")
}

func (airtrafficcontrolservicetimeslicepropertytype *AirTrafficControlServiceTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirTrafficControlServiceTimeSlicePropertyType_stagedOrder, stage.AirTrafficControlServiceTimeSlicePropertyTypes_referenceOrder, airtrafficcontrolservicetimeslicepropertytype, "AirTrafficControlServiceTimeSlicePropertyType")
}

func (airtrafficcontrolservicetimeslicetype *AirTrafficControlServiceTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirTrafficControlServiceTimeSliceType_stagedOrder, stage.AirTrafficControlServiceTimeSliceTypes_referenceOrder, airtrafficcontrolservicetimeslicetype, "AirTrafficControlServiceTimeSliceType")
}

func (airtrafficcontrolservicetype *AirTrafficControlServiceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirTrafficControlServiceType_stagedOrder, stage.AirTrafficControlServiceTypes_referenceOrder, airtrafficcontrolservicetype, "AirTrafficControlServiceType")
}

func (airtrafficmanagementservicetimeslicepropertytype *AirTrafficManagementServiceTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirTrafficManagementServiceTimeSlicePropertyType_stagedOrder, stage.AirTrafficManagementServiceTimeSlicePropertyTypes_referenceOrder, airtrafficmanagementservicetimeslicepropertytype, "AirTrafficManagementServiceTimeSlicePropertyType")
}

func (airtrafficmanagementservicetimeslicetype *AirTrafficManagementServiceTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirTrafficManagementServiceTimeSliceType_stagedOrder, stage.AirTrafficManagementServiceTimeSliceTypes_referenceOrder, airtrafficmanagementservicetimeslicetype, "AirTrafficManagementServiceTimeSliceType")
}

func (airtrafficmanagementservicetype *AirTrafficManagementServiceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirTrafficManagementServiceType_stagedOrder, stage.AirTrafficManagementServiceTypes_referenceOrder, airtrafficmanagementservicetype, "AirTrafficManagementServiceType")
}

func (aircraftgroundservicetimeslicepropertytype *AircraftGroundServiceTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AircraftGroundServiceTimeSlicePropertyType_stagedOrder, stage.AircraftGroundServiceTimeSlicePropertyTypes_referenceOrder, aircraftgroundservicetimeslicepropertytype, "AircraftGroundServiceTimeSlicePropertyType")
}

func (aircraftgroundservicetimeslicetype *AircraftGroundServiceTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AircraftGroundServiceTimeSliceType_stagedOrder, stage.AircraftGroundServiceTimeSliceTypes_referenceOrder, aircraftgroundservicetimeslicetype, "AircraftGroundServiceTimeSliceType")
}

func (aircraftgroundservicetype *AircraftGroundServiceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AircraftGroundServiceType_stagedOrder, stage.AircraftGroundServiceTypes_referenceOrder, aircraftgroundservicetype, "AircraftGroundServiceType")
}

func (airportclearanceservicetimeslicepropertytype *AirportClearanceServiceTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirportClearanceServiceTimeSlicePropertyType_stagedOrder, stage.AirportClearanceServiceTimeSlicePropertyTypes_referenceOrder, airportclearanceservicetimeslicepropertytype, "AirportClearanceServiceTimeSlicePropertyType")
}

func (airportclearanceservicetimeslicetype *AirportClearanceServiceTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirportClearanceServiceTimeSliceType_stagedOrder, stage.AirportClearanceServiceTimeSliceTypes_referenceOrder, airportclearanceservicetimeslicetype, "AirportClearanceServiceTimeSliceType")
}

func (airportclearanceservicetype *AirportClearanceServiceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirportClearanceServiceType_stagedOrder, stage.AirportClearanceServiceTypes_referenceOrder, airportclearanceservicetype, "AirportClearanceServiceType")
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirportHeliportPropertyType_stagedOrder, stage.AirportHeliportPropertyTypes_referenceOrder, airportheliportpropertytype, "AirportHeliportPropertyType")
}

func (airportsuppliesservicetimeslicepropertytype *AirportSuppliesServiceTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirportSuppliesServiceTimeSlicePropertyType_stagedOrder, stage.AirportSuppliesServiceTimeSlicePropertyTypes_referenceOrder, airportsuppliesservicetimeslicepropertytype, "AirportSuppliesServiceTimeSlicePropertyType")
}

func (airportsuppliesservicetimeslicetype *AirportSuppliesServiceTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirportSuppliesServiceTimeSliceType_stagedOrder, stage.AirportSuppliesServiceTimeSliceTypes_referenceOrder, airportsuppliesservicetimeslicetype, "AirportSuppliesServiceTimeSliceType")
}

func (airportsuppliesservicetype *AirportSuppliesServiceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirportSuppliesServiceType_stagedOrder, stage.AirportSuppliesServiceTypes_referenceOrder, airportsuppliesservicetype, "AirportSuppliesServiceType")
}

func (airspacepropertytype *AirspacePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AirspacePropertyType_stagedOrder, stage.AirspacePropertyTypes_referenceOrder, airspacepropertytype, "AirspacePropertyType")
}

func (callsigndetailpropertytype *CallsignDetailPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CallsignDetailPropertyType_stagedOrder, stage.CallsignDetailPropertyTypes_referenceOrder, callsigndetailpropertytype, "CallsignDetailPropertyType")
}

func (callsigndetailtype *CallsignDetailType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CallsignDetailType_stagedOrder, stage.CallsignDetailTypes_referenceOrder, callsigndetailtype, "CallsignDetailType")
}

func (contactinformationpropertytype *ContactInformationPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ContactInformationPropertyType_stagedOrder, stage.ContactInformationPropertyTypes_referenceOrder, contactinformationpropertytype, "ContactInformationPropertyType")
}

func (contactinformationtype *ContactInformationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ContactInformationType_stagedOrder, stage.ContactInformationTypes_referenceOrder, contactinformationtype, "ContactInformationType")
}

func (designatedpointpropertytype *DesignatedPointPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DesignatedPointPropertyType_stagedOrder, stage.DesignatedPointPropertyTypes_referenceOrder, designatedpointpropertytype, "DesignatedPointPropertyType")
}

func (directionfinderpropertytype *DirectionFinderPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DirectionFinderPropertyType_stagedOrder, stage.DirectionFinderPropertyTypes_referenceOrder, directionfinderpropertytype, "DirectionFinderPropertyType")
}

func (elevatedpointpropertytype *ElevatedPointPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ElevatedPointPropertyType_stagedOrder, stage.ElevatedPointPropertyTypes_referenceOrder, elevatedpointpropertytype, "ElevatedPointPropertyType")
}

func (elevatedpointtype *ElevatedPointType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ElevatedPointType_stagedOrder, stage.ElevatedPointTypes_referenceOrder, elevatedpointtype, "ElevatedPointType")
}

func (firefightingservicetimeslicepropertytype *FireFightingServiceTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FireFightingServiceTimeSlicePropertyType_stagedOrder, stage.FireFightingServiceTimeSlicePropertyTypes_referenceOrder, firefightingservicetimeslicepropertytype, "FireFightingServiceTimeSlicePropertyType")
}

func (firefightingservicetimeslicetype *FireFightingServiceTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FireFightingServiceTimeSliceType_stagedOrder, stage.FireFightingServiceTimeSliceTypes_referenceOrder, firefightingservicetimeslicetype, "FireFightingServiceTimeSliceType")
}

func (firefightingservicetype *FireFightingServiceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FireFightingServiceType_stagedOrder, stage.FireFightingServiceTypes_referenceOrder, firefightingservicetype, "FireFightingServiceType")
}

func (fuelpropertytype *FuelPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FuelPropertyType_stagedOrder, stage.FuelPropertyTypes_referenceOrder, fuelpropertytype, "FuelPropertyType")
}

func (fueltype *FuelType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FuelType_stagedOrder, stage.FuelTypes_referenceOrder, fueltype, "FuelType")
}

func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GroundLightSystemPropertyType_stagedOrder, stage.GroundLightSystemPropertyTypes_referenceOrder, groundlightsystempropertytype, "GroundLightSystemPropertyType")
}

func (groundtrafficcontrolservicetimeslicepropertytype *GroundTrafficControlServiceTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GroundTrafficControlServiceTimeSlicePropertyType_stagedOrder, stage.GroundTrafficControlServiceTimeSlicePropertyTypes_referenceOrder, groundtrafficcontrolservicetimeslicepropertytype, "GroundTrafficControlServiceTimeSlicePropertyType")
}

func (groundtrafficcontrolservicetimeslicetype *GroundTrafficControlServiceTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GroundTrafficControlServiceTimeSliceType_stagedOrder, stage.GroundTrafficControlServiceTimeSliceTypes_referenceOrder, groundtrafficcontrolservicetimeslicetype, "GroundTrafficControlServiceTimeSliceType")
}

func (groundtrafficcontrolservicetype *GroundTrafficControlServiceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GroundTrafficControlServiceType_stagedOrder, stage.GroundTrafficControlServiceTypes_referenceOrder, groundtrafficcontrolservicetype, "GroundTrafficControlServiceType")
}

func (holdingpatternpropertytype *HoldingPatternPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.HoldingPatternPropertyType_stagedOrder, stage.HoldingPatternPropertyTypes_referenceOrder, holdingpatternpropertytype, "HoldingPatternPropertyType")
}

func (informationservicetimeslicepropertytype *InformationServiceTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.InformationServiceTimeSlicePropertyType_stagedOrder, stage.InformationServiceTimeSlicePropertyTypes_referenceOrder, informationservicetimeslicepropertytype, "InformationServiceTimeSlicePropertyType")
}

func (informationservicetimeslicetype *InformationServiceTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.InformationServiceTimeSliceType_stagedOrder, stage.InformationServiceTimeSliceTypes_referenceOrder, informationservicetimeslicetype, "InformationServiceTimeSliceType")
}

func (informationservicetype *InformationServiceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.InformationServiceType_stagedOrder, stage.InformationServiceTypes_referenceOrder, informationservicetype, "InformationServiceType")
}

func (lightactivationpropertytype *LightActivationPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LightActivationPropertyType_stagedOrder, stage.LightActivationPropertyTypes_referenceOrder, lightactivationpropertytype, "LightActivationPropertyType")
}

func (lightactivationtype *LightActivationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LightActivationType_stagedOrder, stage.LightActivationTypes_referenceOrder, lightactivationtype, "LightActivationType")
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LinguisticNotePropertyType_stagedOrder, stage.LinguisticNotePropertyTypes_referenceOrder, linguisticnotepropertytype, "LinguisticNotePropertyType")
}

func (linguisticnotetype *LinguisticNoteType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LinguisticNoteType_stagedOrder, stage.LinguisticNoteTypes_referenceOrder, linguisticnotetype, "LinguisticNoteType")
}

func (navaidpropertytype *NavaidPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NavaidPropertyType_stagedOrder, stage.NavaidPropertyTypes_referenceOrder, navaidpropertytype, "NavaidPropertyType")
}

func (nitrogenpropertytype *NitrogenPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NitrogenPropertyType_stagedOrder, stage.NitrogenPropertyTypes_referenceOrder, nitrogenpropertytype, "NitrogenPropertyType")
}

func (nitrogentype *NitrogenType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NitrogenType_stagedOrder, stage.NitrogenTypes_referenceOrder, nitrogentype, "NitrogenType")
}

func (notepropertytype *NotePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NotePropertyType_stagedOrder, stage.NotePropertyTypes_referenceOrder, notepropertytype, "NotePropertyType")
}

func (notetype *NoteType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoteType_stagedOrder, stage.NoteTypes_referenceOrder, notetype, "NoteType")
}

func (oilpropertytype *OilPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.OilPropertyType_stagedOrder, stage.OilPropertyTypes_referenceOrder, oilpropertytype, "OilPropertyType")
}

func (oiltype *OilType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.OilType_stagedOrder, stage.OilTypes_referenceOrder, oiltype, "OilType")
}

func (onlinecontactpropertytype *OnlineContactPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.OnlineContactPropertyType_stagedOrder, stage.OnlineContactPropertyTypes_referenceOrder, onlinecontactpropertytype, "OnlineContactPropertyType")
}

func (onlinecontacttype *OnlineContactType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.OnlineContactType_stagedOrder, stage.OnlineContactTypes_referenceOrder, onlinecontacttype, "OnlineContactType")
}

func (organisationauthorityassociationpropertytype *OrganisationAuthorityAssociationPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.OrganisationAuthorityAssociationPropertyType_stagedOrder, stage.OrganisationAuthorityAssociationPropertyTypes_referenceOrder, organisationauthorityassociationpropertytype, "OrganisationAuthorityAssociationPropertyType")
}

func (organisationauthorityassociationtype *OrganisationAuthorityAssociationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.OrganisationAuthorityAssociationType_stagedOrder, stage.OrganisationAuthorityAssociationTypes_referenceOrder, organisationauthorityassociationtype, "OrganisationAuthorityAssociationType")
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.OrganisationAuthorityPropertyType_stagedOrder, stage.OrganisationAuthorityPropertyTypes_referenceOrder, organisationauthoritypropertytype, "OrganisationAuthorityPropertyType")
}

func (organisationauthoritytimeslicepropertytype *OrganisationAuthorityTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.OrganisationAuthorityTimeSlicePropertyType_stagedOrder, stage.OrganisationAuthorityTimeSlicePropertyTypes_referenceOrder, organisationauthoritytimeslicepropertytype, "OrganisationAuthorityTimeSlicePropertyType")
}

func (organisationauthoritytimeslicetype *OrganisationAuthorityTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.OrganisationAuthorityTimeSliceType_stagedOrder, stage.OrganisationAuthorityTimeSliceTypes_referenceOrder, organisationauthoritytimeslicetype, "OrganisationAuthorityTimeSliceType")
}

func (organisationauthoritytype *OrganisationAuthorityType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.OrganisationAuthorityType_stagedOrder, stage.OrganisationAuthorityTypes_referenceOrder, organisationauthoritytype, "OrganisationAuthorityType")
}

func (oxygenpropertytype *OxygenPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.OxygenPropertyType_stagedOrder, stage.OxygenPropertyTypes_referenceOrder, oxygenpropertytype, "OxygenPropertyType")
}

func (oxygentype *OxygenType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.OxygenType_stagedOrder, stage.OxygenTypes_referenceOrder, oxygentype, "OxygenType")
}

func (passengerservicetimeslicepropertytype *PassengerServiceTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PassengerServiceTimeSlicePropertyType_stagedOrder, stage.PassengerServiceTimeSlicePropertyTypes_referenceOrder, passengerservicetimeslicepropertytype, "PassengerServiceTimeSlicePropertyType")
}

func (passengerservicetimeslicetype *PassengerServiceTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PassengerServiceTimeSliceType_stagedOrder, stage.PassengerServiceTimeSliceTypes_referenceOrder, passengerservicetimeslicetype, "PassengerServiceTimeSliceType")
}

func (passengerservicetype *PassengerServiceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PassengerServiceType_stagedOrder, stage.PassengerServiceTypes_referenceOrder, passengerservicetype, "PassengerServiceType")
}

func (pilotcontrolledlightingtimeslicepropertytype *PilotControlledLightingTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PilotControlledLightingTimeSlicePropertyType_stagedOrder, stage.PilotControlledLightingTimeSlicePropertyTypes_referenceOrder, pilotcontrolledlightingtimeslicepropertytype, "PilotControlledLightingTimeSlicePropertyType")
}

func (pilotcontrolledlightingtimeslicetype *PilotControlledLightingTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PilotControlledLightingTimeSliceType_stagedOrder, stage.PilotControlledLightingTimeSliceTypes_referenceOrder, pilotcontrolledlightingtimeslicetype, "PilotControlledLightingTimeSliceType")
}

func (pilotcontrolledlightingtype *PilotControlledLightingType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PilotControlledLightingType_stagedOrder, stage.PilotControlledLightingTypes_referenceOrder, pilotcontrolledlightingtype, "PilotControlledLightingType")
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

func (procedurepropertytype *ProcedurePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ProcedurePropertyType_stagedOrder, stage.ProcedurePropertyTypes_referenceOrder, procedurepropertytype, "ProcedurePropertyType")
}

func (radiocommunicationchannelpropertytype *RadioCommunicationChannelPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RadioCommunicationChannelPropertyType_stagedOrder, stage.RadioCommunicationChannelPropertyTypes_referenceOrder, radiocommunicationchannelpropertytype, "RadioCommunicationChannelPropertyType")
}

func (radiocommunicationchanneltimeslicepropertytype *RadioCommunicationChannelTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RadioCommunicationChannelTimeSlicePropertyType_stagedOrder, stage.RadioCommunicationChannelTimeSlicePropertyTypes_referenceOrder, radiocommunicationchanneltimeslicepropertytype, "RadioCommunicationChannelTimeSlicePropertyType")
}

func (radiocommunicationchanneltimeslicetype *RadioCommunicationChannelTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RadioCommunicationChannelTimeSliceType_stagedOrder, stage.RadioCommunicationChannelTimeSliceTypes_referenceOrder, radiocommunicationchanneltimeslicetype, "RadioCommunicationChannelTimeSliceType")
}

func (radiocommunicationchanneltype *RadioCommunicationChannelType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RadioCommunicationChannelType_stagedOrder, stage.RadioCommunicationChannelTypes_referenceOrder, radiocommunicationchanneltype, "RadioCommunicationChannelType")
}

func (radiocommunicationoperationalstatuspropertytype *RadioCommunicationOperationalStatusPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RadioCommunicationOperationalStatusPropertyType_stagedOrder, stage.RadioCommunicationOperationalStatusPropertyTypes_referenceOrder, radiocommunicationoperationalstatuspropertytype, "RadioCommunicationOperationalStatusPropertyType")
}

func (radiocommunicationoperationalstatustype *RadioCommunicationOperationalStatusType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RadioCommunicationOperationalStatusType_stagedOrder, stage.RadioCommunicationOperationalStatusTypes_referenceOrder, radiocommunicationoperationalstatustype, "RadioCommunicationOperationalStatusType")
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

func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RunwayCentrelinePointPropertyType_stagedOrder, stage.RunwayCentrelinePointPropertyTypes_referenceOrder, runwaycentrelinepointpropertytype, "RunwayCentrelinePointPropertyType")
}

func (searchrescueservicetimeslicepropertytype *SearchRescueServiceTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SearchRescueServiceTimeSlicePropertyType_stagedOrder, stage.SearchRescueServiceTimeSlicePropertyTypes_referenceOrder, searchrescueservicetimeslicepropertytype, "SearchRescueServiceTimeSlicePropertyType")
}

func (searchrescueservicetimeslicetype *SearchRescueServiceTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SearchRescueServiceTimeSliceType_stagedOrder, stage.SearchRescueServiceTimeSliceTypes_referenceOrder, searchrescueservicetimeslicetype, "SearchRescueServiceTimeSliceType")
}

func (searchrescueservicetype *SearchRescueServiceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SearchRescueServiceType_stagedOrder, stage.SearchRescueServiceTypes_referenceOrder, searchrescueservicetype, "SearchRescueServiceType")
}

func (serviceoperationalstatuspropertytype *ServiceOperationalStatusPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ServiceOperationalStatusPropertyType_stagedOrder, stage.ServiceOperationalStatusPropertyTypes_referenceOrder, serviceoperationalstatuspropertytype, "ServiceOperationalStatusPropertyType")
}

func (serviceoperationalstatustype *ServiceOperationalStatusType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ServiceOperationalStatusType_stagedOrder, stage.ServiceOperationalStatusTypes_referenceOrder, serviceoperationalstatustype, "ServiceOperationalStatusType")
}

func (specialdatetimeslicepropertytype *SpecialDateTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SpecialDateTimeSlicePropertyType_stagedOrder, stage.SpecialDateTimeSlicePropertyTypes_referenceOrder, specialdatetimeslicepropertytype, "SpecialDateTimeSlicePropertyType")
}

func (specialdatetimeslicetype *SpecialDateTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SpecialDateTimeSliceType_stagedOrder, stage.SpecialDateTimeSliceTypes_referenceOrder, specialdatetimeslicetype, "SpecialDateTimeSliceType")
}

func (specialdatetype *SpecialDateType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SpecialDateType_stagedOrder, stage.SpecialDateTypes_referenceOrder, specialdatetype, "SpecialDateType")
}

func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StandardLevelColumnPropertyType_stagedOrder, stage.StandardLevelColumnPropertyTypes_referenceOrder, standardlevelcolumnpropertytype, "StandardLevelColumnPropertyType")
}

func (standardlevelcolumntimeslicepropertytype *StandardLevelColumnTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StandardLevelColumnTimeSlicePropertyType_stagedOrder, stage.StandardLevelColumnTimeSlicePropertyTypes_referenceOrder, standardlevelcolumntimeslicepropertytype, "StandardLevelColumnTimeSlicePropertyType")
}

func (standardlevelcolumntimeslicetype *StandardLevelColumnTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StandardLevelColumnTimeSliceType_stagedOrder, stage.StandardLevelColumnTimeSliceTypes_referenceOrder, standardlevelcolumntimeslicetype, "StandardLevelColumnTimeSliceType")
}

func (standardlevelcolumntype *StandardLevelColumnType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StandardLevelColumnType_stagedOrder, stage.StandardLevelColumnTypes_referenceOrder, standardlevelcolumntype, "StandardLevelColumnType")
}

func (standardlevelpropertytype *StandardLevelPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StandardLevelPropertyType_stagedOrder, stage.StandardLevelPropertyTypes_referenceOrder, standardlevelpropertytype, "StandardLevelPropertyType")
}

func (standardlevelsectortimeslicepropertytype *StandardLevelSectorTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StandardLevelSectorTimeSlicePropertyType_stagedOrder, stage.StandardLevelSectorTimeSlicePropertyTypes_referenceOrder, standardlevelsectortimeslicepropertytype, "StandardLevelSectorTimeSlicePropertyType")
}

func (standardlevelsectortimeslicetype *StandardLevelSectorTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StandardLevelSectorTimeSliceType_stagedOrder, stage.StandardLevelSectorTimeSliceTypes_referenceOrder, standardlevelsectortimeslicetype, "StandardLevelSectorTimeSliceType")
}

func (standardlevelsectortype *StandardLevelSectorType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StandardLevelSectorType_stagedOrder, stage.StandardLevelSectorTypes_referenceOrder, standardlevelsectortype, "StandardLevelSectorType")
}

func (standardleveltablepropertytype *StandardLevelTablePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StandardLevelTablePropertyType_stagedOrder, stage.StandardLevelTablePropertyTypes_referenceOrder, standardleveltablepropertytype, "StandardLevelTablePropertyType")
}

func (standardleveltabletimeslicepropertytype *StandardLevelTableTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StandardLevelTableTimeSlicePropertyType_stagedOrder, stage.StandardLevelTableTimeSlicePropertyTypes_referenceOrder, standardleveltabletimeslicepropertytype, "StandardLevelTableTimeSlicePropertyType")
}

func (standardleveltabletimeslicetype *StandardLevelTableTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StandardLevelTableTimeSliceType_stagedOrder, stage.StandardLevelTableTimeSliceTypes_referenceOrder, standardleveltabletimeslicetype, "StandardLevelTableTimeSliceType")
}

func (standardleveltabletype *StandardLevelTableType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StandardLevelTableType_stagedOrder, stage.StandardLevelTableTypes_referenceOrder, standardleveltabletype, "StandardLevelTableType")
}

func (standardleveltype *StandardLevelType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StandardLevelType_stagedOrder, stage.StandardLevelTypes_referenceOrder, standardleveltype, "StandardLevelType")
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

func (unitavailabilitypropertytype *UnitAvailabilityPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.UnitAvailabilityPropertyType_stagedOrder, stage.UnitAvailabilityPropertyTypes_referenceOrder, unitavailabilitypropertytype, "UnitAvailabilityPropertyType")
}

func (unitavailabilitytype *UnitAvailabilityType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.UnitAvailabilityType_stagedOrder, stage.UnitAvailabilityTypes_referenceOrder, unitavailabilitytype, "UnitAvailabilityType")
}

func (unitdependencypropertytype *UnitDependencyPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.UnitDependencyPropertyType_stagedOrder, stage.UnitDependencyPropertyTypes_referenceOrder, unitdependencypropertytype, "UnitDependencyPropertyType")
}

func (unitdependencytype *UnitDependencyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.UnitDependencyType_stagedOrder, stage.UnitDependencyTypes_referenceOrder, unitdependencytype, "UnitDependencyType")
}

func (unitpropertytype *UnitPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.UnitPropertyType_stagedOrder, stage.UnitPropertyTypes_referenceOrder, unitpropertytype, "UnitPropertyType")
}

func (unittimeslicepropertytype *UnitTimeSlicePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.UnitTimeSlicePropertyType_stagedOrder, stage.UnitTimeSlicePropertyTypes_referenceOrder, unittimeslicepropertytype, "UnitTimeSlicePropertyType")
}

func (unittimeslicetype *UnitTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.UnitTimeSliceType_stagedOrder, stage.UnitTimeSliceTypes_referenceOrder, unittimeslicetype, "UnitTimeSliceType")
}

func (unittype *UnitType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.UnitType_stagedOrder, stage.UnitTypes_referenceOrder, unittype, "UnitType")
}

func (vorpropertytype *VORPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.VORPropertyType_stagedOrder, stage.VORPropertyTypes_referenceOrder, vorpropertytype, "VORPropertyType")
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

func (abstractairportgroundservicetype *AbstractAirportGroundServiceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractairportgroundservicetype, abstractairportgroundservicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractairportgroundservicetype *AbstractAirportGroundServiceType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractairportgroundservicetype.GongGetIdentifier(stage)
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractpropertieswithscheduletype, abstractpropertieswithscheduletype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractpropertieswithscheduletype.GongGetIdentifier(stage)
}

func (abstractservicetype *AbstractServiceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractservicetype, abstractservicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractservicetype *AbstractServiceType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractservicetype.GongGetIdentifier(stage)
}

func (abstracttrafficseparationservicetype *AbstractTrafficSeparationServiceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstracttrafficseparationservicetype, abstracttrafficseparationservicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstracttrafficseparationservicetype *AbstractTrafficSeparationServiceType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstracttrafficseparationservicetype.GongGetIdentifier(stage)
}

func (aerialrefuellingpropertytype *AerialRefuellingPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(aerialrefuellingpropertytype, aerialrefuellingpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (aerialrefuellingpropertytype *AerialRefuellingPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return aerialrefuellingpropertytype.GongGetIdentifier(stage)
}

func (airtrafficcontrolservicetimeslicepropertytype *AirTrafficControlServiceTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airtrafficcontrolservicetimeslicepropertytype, airtrafficcontrolservicetimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airtrafficcontrolservicetimeslicepropertytype *AirTrafficControlServiceTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return airtrafficcontrolservicetimeslicepropertytype.GongGetIdentifier(stage)
}

func (airtrafficcontrolservicetimeslicetype *AirTrafficControlServiceTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airtrafficcontrolservicetimeslicetype, airtrafficcontrolservicetimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airtrafficcontrolservicetimeslicetype *AirTrafficControlServiceTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return airtrafficcontrolservicetimeslicetype.GongGetIdentifier(stage)
}

func (airtrafficcontrolservicetype *AirTrafficControlServiceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airtrafficcontrolservicetype, airtrafficcontrolservicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airtrafficcontrolservicetype *AirTrafficControlServiceType) GongGetReferenceIdentifier(stage *Stage) string {
	return airtrafficcontrolservicetype.GongGetIdentifier(stage)
}

func (airtrafficmanagementservicetimeslicepropertytype *AirTrafficManagementServiceTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airtrafficmanagementservicetimeslicepropertytype, airtrafficmanagementservicetimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airtrafficmanagementservicetimeslicepropertytype *AirTrafficManagementServiceTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return airtrafficmanagementservicetimeslicepropertytype.GongGetIdentifier(stage)
}

func (airtrafficmanagementservicetimeslicetype *AirTrafficManagementServiceTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airtrafficmanagementservicetimeslicetype, airtrafficmanagementservicetimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airtrafficmanagementservicetimeslicetype *AirTrafficManagementServiceTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return airtrafficmanagementservicetimeslicetype.GongGetIdentifier(stage)
}

func (airtrafficmanagementservicetype *AirTrafficManagementServiceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airtrafficmanagementservicetype, airtrafficmanagementservicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airtrafficmanagementservicetype *AirTrafficManagementServiceType) GongGetReferenceIdentifier(stage *Stage) string {
	return airtrafficmanagementservicetype.GongGetIdentifier(stage)
}

func (aircraftgroundservicetimeslicepropertytype *AircraftGroundServiceTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(aircraftgroundservicetimeslicepropertytype, aircraftgroundservicetimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (aircraftgroundservicetimeslicepropertytype *AircraftGroundServiceTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return aircraftgroundservicetimeslicepropertytype.GongGetIdentifier(stage)
}

func (aircraftgroundservicetimeslicetype *AircraftGroundServiceTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(aircraftgroundservicetimeslicetype, aircraftgroundservicetimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (aircraftgroundservicetimeslicetype *AircraftGroundServiceTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return aircraftgroundservicetimeslicetype.GongGetIdentifier(stage)
}

func (aircraftgroundservicetype *AircraftGroundServiceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(aircraftgroundservicetype, aircraftgroundservicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (aircraftgroundservicetype *AircraftGroundServiceType) GongGetReferenceIdentifier(stage *Stage) string {
	return aircraftgroundservicetype.GongGetIdentifier(stage)
}

func (airportclearanceservicetimeslicepropertytype *AirportClearanceServiceTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airportclearanceservicetimeslicepropertytype, airportclearanceservicetimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airportclearanceservicetimeslicepropertytype *AirportClearanceServiceTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return airportclearanceservicetimeslicepropertytype.GongGetIdentifier(stage)
}

func (airportclearanceservicetimeslicetype *AirportClearanceServiceTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airportclearanceservicetimeslicetype, airportclearanceservicetimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airportclearanceservicetimeslicetype *AirportClearanceServiceTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return airportclearanceservicetimeslicetype.GongGetIdentifier(stage)
}

func (airportclearanceservicetype *AirportClearanceServiceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airportclearanceservicetype, airportclearanceservicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airportclearanceservicetype *AirportClearanceServiceType) GongGetReferenceIdentifier(stage *Stage) string {
	return airportclearanceservicetype.GongGetIdentifier(stage)
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airportheliportpropertytype, airportheliportpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airportheliportpropertytype *AirportHeliportPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return airportheliportpropertytype.GongGetIdentifier(stage)
}

func (airportsuppliesservicetimeslicepropertytype *AirportSuppliesServiceTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airportsuppliesservicetimeslicepropertytype, airportsuppliesservicetimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airportsuppliesservicetimeslicepropertytype *AirportSuppliesServiceTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return airportsuppliesservicetimeslicepropertytype.GongGetIdentifier(stage)
}

func (airportsuppliesservicetimeslicetype *AirportSuppliesServiceTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airportsuppliesservicetimeslicetype, airportsuppliesservicetimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airportsuppliesservicetimeslicetype *AirportSuppliesServiceTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return airportsuppliesservicetimeslicetype.GongGetIdentifier(stage)
}

func (airportsuppliesservicetype *AirportSuppliesServiceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airportsuppliesservicetype, airportsuppliesservicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airportsuppliesservicetype *AirportSuppliesServiceType) GongGetReferenceIdentifier(stage *Stage) string {
	return airportsuppliesservicetype.GongGetIdentifier(stage)
}

func (airspacepropertytype *AirspacePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(airspacepropertytype, airspacepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (airspacepropertytype *AirspacePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return airspacepropertytype.GongGetIdentifier(stage)
}

func (callsigndetailpropertytype *CallsignDetailPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(callsigndetailpropertytype, callsigndetailpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (callsigndetailpropertytype *CallsignDetailPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return callsigndetailpropertytype.GongGetIdentifier(stage)
}

func (callsigndetailtype *CallsignDetailType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(callsigndetailtype, callsigndetailtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (callsigndetailtype *CallsignDetailType) GongGetReferenceIdentifier(stage *Stage) string {
	return callsigndetailtype.GongGetIdentifier(stage)
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

func (designatedpointpropertytype *DesignatedPointPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(designatedpointpropertytype, designatedpointpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (designatedpointpropertytype *DesignatedPointPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return designatedpointpropertytype.GongGetIdentifier(stage)
}

func (directionfinderpropertytype *DirectionFinderPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(directionfinderpropertytype, directionfinderpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (directionfinderpropertytype *DirectionFinderPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return directionfinderpropertytype.GongGetIdentifier(stage)
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

func (firefightingservicetimeslicepropertytype *FireFightingServiceTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(firefightingservicetimeslicepropertytype, firefightingservicetimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (firefightingservicetimeslicepropertytype *FireFightingServiceTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return firefightingservicetimeslicepropertytype.GongGetIdentifier(stage)
}

func (firefightingservicetimeslicetype *FireFightingServiceTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(firefightingservicetimeslicetype, firefightingservicetimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (firefightingservicetimeslicetype *FireFightingServiceTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return firefightingservicetimeslicetype.GongGetIdentifier(stage)
}

func (firefightingservicetype *FireFightingServiceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(firefightingservicetype, firefightingservicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (firefightingservicetype *FireFightingServiceType) GongGetReferenceIdentifier(stage *Stage) string {
	return firefightingservicetype.GongGetIdentifier(stage)
}

func (fuelpropertytype *FuelPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(fuelpropertytype, fuelpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (fuelpropertytype *FuelPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return fuelpropertytype.GongGetIdentifier(stage)
}

func (fueltype *FuelType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(fueltype, fueltype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (fueltype *FuelType) GongGetReferenceIdentifier(stage *Stage) string {
	return fueltype.GongGetIdentifier(stage)
}

func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(groundlightsystempropertytype, groundlightsystempropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return groundlightsystempropertytype.GongGetIdentifier(stage)
}

func (groundtrafficcontrolservicetimeslicepropertytype *GroundTrafficControlServiceTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(groundtrafficcontrolservicetimeslicepropertytype, groundtrafficcontrolservicetimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (groundtrafficcontrolservicetimeslicepropertytype *GroundTrafficControlServiceTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return groundtrafficcontrolservicetimeslicepropertytype.GongGetIdentifier(stage)
}

func (groundtrafficcontrolservicetimeslicetype *GroundTrafficControlServiceTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(groundtrafficcontrolservicetimeslicetype, groundtrafficcontrolservicetimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (groundtrafficcontrolservicetimeslicetype *GroundTrafficControlServiceTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return groundtrafficcontrolservicetimeslicetype.GongGetIdentifier(stage)
}

func (groundtrafficcontrolservicetype *GroundTrafficControlServiceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(groundtrafficcontrolservicetype, groundtrafficcontrolservicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (groundtrafficcontrolservicetype *GroundTrafficControlServiceType) GongGetReferenceIdentifier(stage *Stage) string {
	return groundtrafficcontrolservicetype.GongGetIdentifier(stage)
}

func (holdingpatternpropertytype *HoldingPatternPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(holdingpatternpropertytype, holdingpatternpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (holdingpatternpropertytype *HoldingPatternPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return holdingpatternpropertytype.GongGetIdentifier(stage)
}

func (informationservicetimeslicepropertytype *InformationServiceTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(informationservicetimeslicepropertytype, informationservicetimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (informationservicetimeslicepropertytype *InformationServiceTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return informationservicetimeslicepropertytype.GongGetIdentifier(stage)
}

func (informationservicetimeslicetype *InformationServiceTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(informationservicetimeslicetype, informationservicetimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (informationservicetimeslicetype *InformationServiceTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return informationservicetimeslicetype.GongGetIdentifier(stage)
}

func (informationservicetype *InformationServiceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(informationservicetype, informationservicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (informationservicetype *InformationServiceType) GongGetReferenceIdentifier(stage *Stage) string {
	return informationservicetype.GongGetIdentifier(stage)
}

func (lightactivationpropertytype *LightActivationPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(lightactivationpropertytype, lightactivationpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (lightactivationpropertytype *LightActivationPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return lightactivationpropertytype.GongGetIdentifier(stage)
}

func (lightactivationtype *LightActivationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(lightactivationtype, lightactivationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (lightactivationtype *LightActivationType) GongGetReferenceIdentifier(stage *Stage) string {
	return lightactivationtype.GongGetIdentifier(stage)
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

func (navaidpropertytype *NavaidPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(navaidpropertytype, navaidpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (navaidpropertytype *NavaidPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return navaidpropertytype.GongGetIdentifier(stage)
}

func (nitrogenpropertytype *NitrogenPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(nitrogenpropertytype, nitrogenpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (nitrogenpropertytype *NitrogenPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return nitrogenpropertytype.GongGetIdentifier(stage)
}

func (nitrogentype *NitrogenType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(nitrogentype, nitrogentype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (nitrogentype *NitrogenType) GongGetReferenceIdentifier(stage *Stage) string {
	return nitrogentype.GongGetIdentifier(stage)
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

func (oilpropertytype *OilPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(oilpropertytype, oilpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (oilpropertytype *OilPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return oilpropertytype.GongGetIdentifier(stage)
}

func (oiltype *OilType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(oiltype, oiltype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (oiltype *OilType) GongGetReferenceIdentifier(stage *Stage) string {
	return oiltype.GongGetIdentifier(stage)
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

func (organisationauthorityassociationpropertytype *OrganisationAuthorityAssociationPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(organisationauthorityassociationpropertytype, organisationauthorityassociationpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (organisationauthorityassociationpropertytype *OrganisationAuthorityAssociationPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return organisationauthorityassociationpropertytype.GongGetIdentifier(stage)
}

func (organisationauthorityassociationtype *OrganisationAuthorityAssociationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(organisationauthorityassociationtype, organisationauthorityassociationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (organisationauthorityassociationtype *OrganisationAuthorityAssociationType) GongGetReferenceIdentifier(stage *Stage) string {
	return organisationauthorityassociationtype.GongGetIdentifier(stage)
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(organisationauthoritypropertytype, organisationauthoritypropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return organisationauthoritypropertytype.GongGetIdentifier(stage)
}

func (organisationauthoritytimeslicepropertytype *OrganisationAuthorityTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(organisationauthoritytimeslicepropertytype, organisationauthoritytimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (organisationauthoritytimeslicepropertytype *OrganisationAuthorityTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return organisationauthoritytimeslicepropertytype.GongGetIdentifier(stage)
}

func (organisationauthoritytimeslicetype *OrganisationAuthorityTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(organisationauthoritytimeslicetype, organisationauthoritytimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (organisationauthoritytimeslicetype *OrganisationAuthorityTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return organisationauthoritytimeslicetype.GongGetIdentifier(stage)
}

func (organisationauthoritytype *OrganisationAuthorityType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(organisationauthoritytype, organisationauthoritytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (organisationauthoritytype *OrganisationAuthorityType) GongGetReferenceIdentifier(stage *Stage) string {
	return organisationauthoritytype.GongGetIdentifier(stage)
}

func (oxygenpropertytype *OxygenPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(oxygenpropertytype, oxygenpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (oxygenpropertytype *OxygenPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return oxygenpropertytype.GongGetIdentifier(stage)
}

func (oxygentype *OxygenType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(oxygentype, oxygentype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (oxygentype *OxygenType) GongGetReferenceIdentifier(stage *Stage) string {
	return oxygentype.GongGetIdentifier(stage)
}

func (passengerservicetimeslicepropertytype *PassengerServiceTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(passengerservicetimeslicepropertytype, passengerservicetimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (passengerservicetimeslicepropertytype *PassengerServiceTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return passengerservicetimeslicepropertytype.GongGetIdentifier(stage)
}

func (passengerservicetimeslicetype *PassengerServiceTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(passengerservicetimeslicetype, passengerservicetimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (passengerservicetimeslicetype *PassengerServiceTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return passengerservicetimeslicetype.GongGetIdentifier(stage)
}

func (passengerservicetype *PassengerServiceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(passengerservicetype, passengerservicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (passengerservicetype *PassengerServiceType) GongGetReferenceIdentifier(stage *Stage) string {
	return passengerservicetype.GongGetIdentifier(stage)
}

func (pilotcontrolledlightingtimeslicepropertytype *PilotControlledLightingTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(pilotcontrolledlightingtimeslicepropertytype, pilotcontrolledlightingtimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pilotcontrolledlightingtimeslicepropertytype *PilotControlledLightingTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return pilotcontrolledlightingtimeslicepropertytype.GongGetIdentifier(stage)
}

func (pilotcontrolledlightingtimeslicetype *PilotControlledLightingTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(pilotcontrolledlightingtimeslicetype, pilotcontrolledlightingtimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pilotcontrolledlightingtimeslicetype *PilotControlledLightingTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return pilotcontrolledlightingtimeslicetype.GongGetIdentifier(stage)
}

func (pilotcontrolledlightingtype *PilotControlledLightingType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(pilotcontrolledlightingtype, pilotcontrolledlightingtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pilotcontrolledlightingtype *PilotControlledLightingType) GongGetReferenceIdentifier(stage *Stage) string {
	return pilotcontrolledlightingtype.GongGetIdentifier(stage)
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

func (procedurepropertytype *ProcedurePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(procedurepropertytype, procedurepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (procedurepropertytype *ProcedurePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return procedurepropertytype.GongGetIdentifier(stage)
}

func (radiocommunicationchannelpropertytype *RadioCommunicationChannelPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(radiocommunicationchannelpropertytype, radiocommunicationchannelpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (radiocommunicationchannelpropertytype *RadioCommunicationChannelPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return radiocommunicationchannelpropertytype.GongGetIdentifier(stage)
}

func (radiocommunicationchanneltimeslicepropertytype *RadioCommunicationChannelTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(radiocommunicationchanneltimeslicepropertytype, radiocommunicationchanneltimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (radiocommunicationchanneltimeslicepropertytype *RadioCommunicationChannelTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return radiocommunicationchanneltimeslicepropertytype.GongGetIdentifier(stage)
}

func (radiocommunicationchanneltimeslicetype *RadioCommunicationChannelTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(radiocommunicationchanneltimeslicetype, radiocommunicationchanneltimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (radiocommunicationchanneltimeslicetype *RadioCommunicationChannelTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return radiocommunicationchanneltimeslicetype.GongGetIdentifier(stage)
}

func (radiocommunicationchanneltype *RadioCommunicationChannelType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(radiocommunicationchanneltype, radiocommunicationchanneltype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (radiocommunicationchanneltype *RadioCommunicationChannelType) GongGetReferenceIdentifier(stage *Stage) string {
	return radiocommunicationchanneltype.GongGetIdentifier(stage)
}

func (radiocommunicationoperationalstatuspropertytype *RadioCommunicationOperationalStatusPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(radiocommunicationoperationalstatuspropertytype, radiocommunicationoperationalstatuspropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (radiocommunicationoperationalstatuspropertytype *RadioCommunicationOperationalStatusPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return radiocommunicationoperationalstatuspropertytype.GongGetIdentifier(stage)
}

func (radiocommunicationoperationalstatustype *RadioCommunicationOperationalStatusType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(radiocommunicationoperationalstatustype, radiocommunicationoperationalstatustype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (radiocommunicationoperationalstatustype *RadioCommunicationOperationalStatusType) GongGetReferenceIdentifier(stage *Stage) string {
	return radiocommunicationoperationalstatustype.GongGetIdentifier(stage)
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

func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(runwaycentrelinepointpropertytype, runwaycentrelinepointpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return runwaycentrelinepointpropertytype.GongGetIdentifier(stage)
}

func (searchrescueservicetimeslicepropertytype *SearchRescueServiceTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(searchrescueservicetimeslicepropertytype, searchrescueservicetimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (searchrescueservicetimeslicepropertytype *SearchRescueServiceTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return searchrescueservicetimeslicepropertytype.GongGetIdentifier(stage)
}

func (searchrescueservicetimeslicetype *SearchRescueServiceTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(searchrescueservicetimeslicetype, searchrescueservicetimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (searchrescueservicetimeslicetype *SearchRescueServiceTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return searchrescueservicetimeslicetype.GongGetIdentifier(stage)
}

func (searchrescueservicetype *SearchRescueServiceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(searchrescueservicetype, searchrescueservicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (searchrescueservicetype *SearchRescueServiceType) GongGetReferenceIdentifier(stage *Stage) string {
	return searchrescueservicetype.GongGetIdentifier(stage)
}

func (serviceoperationalstatuspropertytype *ServiceOperationalStatusPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(serviceoperationalstatuspropertytype, serviceoperationalstatuspropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (serviceoperationalstatuspropertytype *ServiceOperationalStatusPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return serviceoperationalstatuspropertytype.GongGetIdentifier(stage)
}

func (serviceoperationalstatustype *ServiceOperationalStatusType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(serviceoperationalstatustype, serviceoperationalstatustype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (serviceoperationalstatustype *ServiceOperationalStatusType) GongGetReferenceIdentifier(stage *Stage) string {
	return serviceoperationalstatustype.GongGetIdentifier(stage)
}

func (specialdatetimeslicepropertytype *SpecialDateTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(specialdatetimeslicepropertytype, specialdatetimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (specialdatetimeslicepropertytype *SpecialDateTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return specialdatetimeslicepropertytype.GongGetIdentifier(stage)
}

func (specialdatetimeslicetype *SpecialDateTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(specialdatetimeslicetype, specialdatetimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (specialdatetimeslicetype *SpecialDateTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return specialdatetimeslicetype.GongGetIdentifier(stage)
}

func (specialdatetype *SpecialDateType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(specialdatetype, specialdatetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (specialdatetype *SpecialDateType) GongGetReferenceIdentifier(stage *Stage) string {
	return specialdatetype.GongGetIdentifier(stage)
}

func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(standardlevelcolumnpropertytype, standardlevelcolumnpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return standardlevelcolumnpropertytype.GongGetIdentifier(stage)
}

func (standardlevelcolumntimeslicepropertytype *StandardLevelColumnTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(standardlevelcolumntimeslicepropertytype, standardlevelcolumntimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (standardlevelcolumntimeslicepropertytype *StandardLevelColumnTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return standardlevelcolumntimeslicepropertytype.GongGetIdentifier(stage)
}

func (standardlevelcolumntimeslicetype *StandardLevelColumnTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(standardlevelcolumntimeslicetype, standardlevelcolumntimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (standardlevelcolumntimeslicetype *StandardLevelColumnTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return standardlevelcolumntimeslicetype.GongGetIdentifier(stage)
}

func (standardlevelcolumntype *StandardLevelColumnType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(standardlevelcolumntype, standardlevelcolumntype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (standardlevelcolumntype *StandardLevelColumnType) GongGetReferenceIdentifier(stage *Stage) string {
	return standardlevelcolumntype.GongGetIdentifier(stage)
}

func (standardlevelpropertytype *StandardLevelPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(standardlevelpropertytype, standardlevelpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (standardlevelpropertytype *StandardLevelPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return standardlevelpropertytype.GongGetIdentifier(stage)
}

func (standardlevelsectortimeslicepropertytype *StandardLevelSectorTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(standardlevelsectortimeslicepropertytype, standardlevelsectortimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (standardlevelsectortimeslicepropertytype *StandardLevelSectorTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return standardlevelsectortimeslicepropertytype.GongGetIdentifier(stage)
}

func (standardlevelsectortimeslicetype *StandardLevelSectorTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(standardlevelsectortimeslicetype, standardlevelsectortimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (standardlevelsectortimeslicetype *StandardLevelSectorTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return standardlevelsectortimeslicetype.GongGetIdentifier(stage)
}

func (standardlevelsectortype *StandardLevelSectorType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(standardlevelsectortype, standardlevelsectortype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (standardlevelsectortype *StandardLevelSectorType) GongGetReferenceIdentifier(stage *Stage) string {
	return standardlevelsectortype.GongGetIdentifier(stage)
}

func (standardleveltablepropertytype *StandardLevelTablePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(standardleveltablepropertytype, standardleveltablepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (standardleveltablepropertytype *StandardLevelTablePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return standardleveltablepropertytype.GongGetIdentifier(stage)
}

func (standardleveltabletimeslicepropertytype *StandardLevelTableTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(standardleveltabletimeslicepropertytype, standardleveltabletimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (standardleveltabletimeslicepropertytype *StandardLevelTableTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return standardleveltabletimeslicepropertytype.GongGetIdentifier(stage)
}

func (standardleveltabletimeslicetype *StandardLevelTableTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(standardleveltabletimeslicetype, standardleveltabletimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (standardleveltabletimeslicetype *StandardLevelTableTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return standardleveltabletimeslicetype.GongGetIdentifier(stage)
}

func (standardleveltabletype *StandardLevelTableType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(standardleveltabletype, standardleveltabletype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (standardleveltabletype *StandardLevelTableType) GongGetReferenceIdentifier(stage *Stage) string {
	return standardleveltabletype.GongGetIdentifier(stage)
}

func (standardleveltype *StandardLevelType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(standardleveltype, standardleveltype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (standardleveltype *StandardLevelType) GongGetReferenceIdentifier(stage *Stage) string {
	return standardleveltype.GongGetIdentifier(stage)
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

func (unitavailabilitypropertytype *UnitAvailabilityPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(unitavailabilitypropertytype, unitavailabilitypropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (unitavailabilitypropertytype *UnitAvailabilityPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return unitavailabilitypropertytype.GongGetIdentifier(stage)
}

func (unitavailabilitytype *UnitAvailabilityType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(unitavailabilitytype, unitavailabilitytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (unitavailabilitytype *UnitAvailabilityType) GongGetReferenceIdentifier(stage *Stage) string {
	return unitavailabilitytype.GongGetIdentifier(stage)
}

func (unitdependencypropertytype *UnitDependencyPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(unitdependencypropertytype, unitdependencypropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (unitdependencypropertytype *UnitDependencyPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return unitdependencypropertytype.GongGetIdentifier(stage)
}

func (unitdependencytype *UnitDependencyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(unitdependencytype, unitdependencytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (unitdependencytype *UnitDependencyType) GongGetReferenceIdentifier(stage *Stage) string {
	return unitdependencytype.GongGetIdentifier(stage)
}

func (unitpropertytype *UnitPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(unitpropertytype, unitpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (unitpropertytype *UnitPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return unitpropertytype.GongGetIdentifier(stage)
}

func (unittimeslicepropertytype *UnitTimeSlicePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(unittimeslicepropertytype, unittimeslicepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (unittimeslicepropertytype *UnitTimeSlicePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return unittimeslicepropertytype.GongGetIdentifier(stage)
}

func (unittimeslicetype *UnitTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(unittimeslicetype, unittimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (unittimeslicetype *UnitTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return unittimeslicetype.GongGetIdentifier(stage)
}

func (unittype *UnitType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(unittype, unittype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (unittype *UnitType) GongGetReferenceIdentifier(stage *Stage) string {
	return unittype.GongGetIdentifier(stage)
}

func (vorpropertytype *VORPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(vorpropertytype, vorpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (vorpropertytype *VORPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return vorpropertytype.GongGetIdentifier(stage)
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

func (abstractairportgroundservicetype *AbstractAirportGroundServiceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractairportgroundservicetype.GongGetIdentifier(stage), "AbstractAirportGroundServiceType", abstractairportgroundservicetype.Name)
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractpropertieswithscheduletype.GongGetIdentifier(stage), "AbstractPropertiesWithScheduleType", abstractpropertieswithscheduletype.Name)
}

func (abstractservicetype *AbstractServiceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractservicetype.GongGetIdentifier(stage), "AbstractServiceType", abstractservicetype.Name)
}

func (abstracttrafficseparationservicetype *AbstractTrafficSeparationServiceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstracttrafficseparationservicetype.GongGetIdentifier(stage), "AbstractTrafficSeparationServiceType", abstracttrafficseparationservicetype.Name)
}

func (aerialrefuellingpropertytype *AerialRefuellingPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(aerialrefuellingpropertytype.GongGetIdentifier(stage), "AerialRefuellingPropertyType", aerialrefuellingpropertytype.Name)
}

func (airtrafficcontrolservicetimeslicepropertytype *AirTrafficControlServiceTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airtrafficcontrolservicetimeslicepropertytype.GongGetIdentifier(stage), "AirTrafficControlServiceTimeSlicePropertyType", airtrafficcontrolservicetimeslicepropertytype.Name)
}

func (airtrafficcontrolservicetimeslicetype *AirTrafficControlServiceTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airtrafficcontrolservicetimeslicetype.GongGetIdentifier(stage), "AirTrafficControlServiceTimeSliceType", airtrafficcontrolservicetimeslicetype.Name)
}

func (airtrafficcontrolservicetype *AirTrafficControlServiceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airtrafficcontrolservicetype.GongGetIdentifier(stage), "AirTrafficControlServiceType", airtrafficcontrolservicetype.Name)
}

func (airtrafficmanagementservicetimeslicepropertytype *AirTrafficManagementServiceTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airtrafficmanagementservicetimeslicepropertytype.GongGetIdentifier(stage), "AirTrafficManagementServiceTimeSlicePropertyType", airtrafficmanagementservicetimeslicepropertytype.Name)
}

func (airtrafficmanagementservicetimeslicetype *AirTrafficManagementServiceTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airtrafficmanagementservicetimeslicetype.GongGetIdentifier(stage), "AirTrafficManagementServiceTimeSliceType", airtrafficmanagementservicetimeslicetype.Name)
}

func (airtrafficmanagementservicetype *AirTrafficManagementServiceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airtrafficmanagementservicetype.GongGetIdentifier(stage), "AirTrafficManagementServiceType", airtrafficmanagementservicetype.Name)
}

func (aircraftgroundservicetimeslicepropertytype *AircraftGroundServiceTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(aircraftgroundservicetimeslicepropertytype.GongGetIdentifier(stage), "AircraftGroundServiceTimeSlicePropertyType", aircraftgroundservicetimeslicepropertytype.Name)
}

func (aircraftgroundservicetimeslicetype *AircraftGroundServiceTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(aircraftgroundservicetimeslicetype.GongGetIdentifier(stage), "AircraftGroundServiceTimeSliceType", aircraftgroundservicetimeslicetype.Name)
}

func (aircraftgroundservicetype *AircraftGroundServiceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(aircraftgroundservicetype.GongGetIdentifier(stage), "AircraftGroundServiceType", aircraftgroundservicetype.Name)
}

func (airportclearanceservicetimeslicepropertytype *AirportClearanceServiceTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airportclearanceservicetimeslicepropertytype.GongGetIdentifier(stage), "AirportClearanceServiceTimeSlicePropertyType", airportclearanceservicetimeslicepropertytype.Name)
}

func (airportclearanceservicetimeslicetype *AirportClearanceServiceTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airportclearanceservicetimeslicetype.GongGetIdentifier(stage), "AirportClearanceServiceTimeSliceType", airportclearanceservicetimeslicetype.Name)
}

func (airportclearanceservicetype *AirportClearanceServiceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airportclearanceservicetype.GongGetIdentifier(stage), "AirportClearanceServiceType", airportclearanceservicetype.Name)
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airportheliportpropertytype.GongGetIdentifier(stage), "AirportHeliportPropertyType", airportheliportpropertytype.Name)
}

func (airportsuppliesservicetimeslicepropertytype *AirportSuppliesServiceTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airportsuppliesservicetimeslicepropertytype.GongGetIdentifier(stage), "AirportSuppliesServiceTimeSlicePropertyType", airportsuppliesservicetimeslicepropertytype.Name)
}

func (airportsuppliesservicetimeslicetype *AirportSuppliesServiceTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airportsuppliesservicetimeslicetype.GongGetIdentifier(stage), "AirportSuppliesServiceTimeSliceType", airportsuppliesservicetimeslicetype.Name)
}

func (airportsuppliesservicetype *AirportSuppliesServiceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airportsuppliesservicetype.GongGetIdentifier(stage), "AirportSuppliesServiceType", airportsuppliesservicetype.Name)
}

func (airspacepropertytype *AirspacePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(airspacepropertytype.GongGetIdentifier(stage), "AirspacePropertyType", airspacepropertytype.Name)
}

func (callsigndetailpropertytype *CallsignDetailPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(callsigndetailpropertytype.GongGetIdentifier(stage), "CallsignDetailPropertyType", callsigndetailpropertytype.Name)
}

func (callsigndetailtype *CallsignDetailType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(callsigndetailtype.GongGetIdentifier(stage), "CallsignDetailType", callsigndetailtype.Name)
}

func (contactinformationpropertytype *ContactInformationPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(contactinformationpropertytype.GongGetIdentifier(stage), "ContactInformationPropertyType", contactinformationpropertytype.Name)
}

func (contactinformationtype *ContactInformationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(contactinformationtype.GongGetIdentifier(stage), "ContactInformationType", contactinformationtype.Name)
}

func (designatedpointpropertytype *DesignatedPointPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(designatedpointpropertytype.GongGetIdentifier(stage), "DesignatedPointPropertyType", designatedpointpropertytype.Name)
}

func (directionfinderpropertytype *DirectionFinderPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(directionfinderpropertytype.GongGetIdentifier(stage), "DirectionFinderPropertyType", directionfinderpropertytype.Name)
}

func (elevatedpointpropertytype *ElevatedPointPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(elevatedpointpropertytype.GongGetIdentifier(stage), "ElevatedPointPropertyType", elevatedpointpropertytype.Name)
}

func (elevatedpointtype *ElevatedPointType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(elevatedpointtype.GongGetIdentifier(stage), "ElevatedPointType", elevatedpointtype.Name)
}

func (firefightingservicetimeslicepropertytype *FireFightingServiceTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(firefightingservicetimeslicepropertytype.GongGetIdentifier(stage), "FireFightingServiceTimeSlicePropertyType", firefightingservicetimeslicepropertytype.Name)
}

func (firefightingservicetimeslicetype *FireFightingServiceTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(firefightingservicetimeslicetype.GongGetIdentifier(stage), "FireFightingServiceTimeSliceType", firefightingservicetimeslicetype.Name)
}

func (firefightingservicetype *FireFightingServiceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(firefightingservicetype.GongGetIdentifier(stage), "FireFightingServiceType", firefightingservicetype.Name)
}

func (fuelpropertytype *FuelPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(fuelpropertytype.GongGetIdentifier(stage), "FuelPropertyType", fuelpropertytype.Name)
}

func (fueltype *FuelType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(fueltype.GongGetIdentifier(stage), "FuelType", fueltype.Name)
}

func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(groundlightsystempropertytype.GongGetIdentifier(stage), "GroundLightSystemPropertyType", groundlightsystempropertytype.Name)
}

func (groundtrafficcontrolservicetimeslicepropertytype *GroundTrafficControlServiceTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(groundtrafficcontrolservicetimeslicepropertytype.GongGetIdentifier(stage), "GroundTrafficControlServiceTimeSlicePropertyType", groundtrafficcontrolservicetimeslicepropertytype.Name)
}

func (groundtrafficcontrolservicetimeslicetype *GroundTrafficControlServiceTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(groundtrafficcontrolservicetimeslicetype.GongGetIdentifier(stage), "GroundTrafficControlServiceTimeSliceType", groundtrafficcontrolservicetimeslicetype.Name)
}

func (groundtrafficcontrolservicetype *GroundTrafficControlServiceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(groundtrafficcontrolservicetype.GongGetIdentifier(stage), "GroundTrafficControlServiceType", groundtrafficcontrolservicetype.Name)
}

func (holdingpatternpropertytype *HoldingPatternPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(holdingpatternpropertytype.GongGetIdentifier(stage), "HoldingPatternPropertyType", holdingpatternpropertytype.Name)
}

func (informationservicetimeslicepropertytype *InformationServiceTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(informationservicetimeslicepropertytype.GongGetIdentifier(stage), "InformationServiceTimeSlicePropertyType", informationservicetimeslicepropertytype.Name)
}

func (informationservicetimeslicetype *InformationServiceTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(informationservicetimeslicetype.GongGetIdentifier(stage), "InformationServiceTimeSliceType", informationservicetimeslicetype.Name)
}

func (informationservicetype *InformationServiceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(informationservicetype.GongGetIdentifier(stage), "InformationServiceType", informationservicetype.Name)
}

func (lightactivationpropertytype *LightActivationPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(lightactivationpropertytype.GongGetIdentifier(stage), "LightActivationPropertyType", lightactivationpropertytype.Name)
}

func (lightactivationtype *LightActivationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(lightactivationtype.GongGetIdentifier(stage), "LightActivationType", lightactivationtype.Name)
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(linguisticnotepropertytype.GongGetIdentifier(stage), "LinguisticNotePropertyType", linguisticnotepropertytype.Name)
}

func (linguisticnotetype *LinguisticNoteType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(linguisticnotetype.GongGetIdentifier(stage), "LinguisticNoteType", linguisticnotetype.Name)
}

func (navaidpropertytype *NavaidPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(navaidpropertytype.GongGetIdentifier(stage), "NavaidPropertyType", navaidpropertytype.Name)
}

func (nitrogenpropertytype *NitrogenPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(nitrogenpropertytype.GongGetIdentifier(stage), "NitrogenPropertyType", nitrogenpropertytype.Name)
}

func (nitrogentype *NitrogenType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(nitrogentype.GongGetIdentifier(stage), "NitrogenType", nitrogentype.Name)
}

func (notepropertytype *NotePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(notepropertytype.GongGetIdentifier(stage), "NotePropertyType", notepropertytype.Name)
}

func (notetype *NoteType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(notetype.GongGetIdentifier(stage), "NoteType", notetype.Name)
}

func (oilpropertytype *OilPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(oilpropertytype.GongGetIdentifier(stage), "OilPropertyType", oilpropertytype.Name)
}

func (oiltype *OilType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(oiltype.GongGetIdentifier(stage), "OilType", oiltype.Name)
}

func (onlinecontactpropertytype *OnlineContactPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(onlinecontactpropertytype.GongGetIdentifier(stage), "OnlineContactPropertyType", onlinecontactpropertytype.Name)
}

func (onlinecontacttype *OnlineContactType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(onlinecontacttype.GongGetIdentifier(stage), "OnlineContactType", onlinecontacttype.Name)
}

func (organisationauthorityassociationpropertytype *OrganisationAuthorityAssociationPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(organisationauthorityassociationpropertytype.GongGetIdentifier(stage), "OrganisationAuthorityAssociationPropertyType", organisationauthorityassociationpropertytype.Name)
}

func (organisationauthorityassociationtype *OrganisationAuthorityAssociationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(organisationauthorityassociationtype.GongGetIdentifier(stage), "OrganisationAuthorityAssociationType", organisationauthorityassociationtype.Name)
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(organisationauthoritypropertytype.GongGetIdentifier(stage), "OrganisationAuthorityPropertyType", organisationauthoritypropertytype.Name)
}

func (organisationauthoritytimeslicepropertytype *OrganisationAuthorityTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(organisationauthoritytimeslicepropertytype.GongGetIdentifier(stage), "OrganisationAuthorityTimeSlicePropertyType", organisationauthoritytimeslicepropertytype.Name)
}

func (organisationauthoritytimeslicetype *OrganisationAuthorityTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(organisationauthoritytimeslicetype.GongGetIdentifier(stage), "OrganisationAuthorityTimeSliceType", organisationauthoritytimeslicetype.Name)
}

func (organisationauthoritytype *OrganisationAuthorityType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(organisationauthoritytype.GongGetIdentifier(stage), "OrganisationAuthorityType", organisationauthoritytype.Name)
}

func (oxygenpropertytype *OxygenPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(oxygenpropertytype.GongGetIdentifier(stage), "OxygenPropertyType", oxygenpropertytype.Name)
}

func (oxygentype *OxygenType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(oxygentype.GongGetIdentifier(stage), "OxygenType", oxygentype.Name)
}

func (passengerservicetimeslicepropertytype *PassengerServiceTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(passengerservicetimeslicepropertytype.GongGetIdentifier(stage), "PassengerServiceTimeSlicePropertyType", passengerservicetimeslicepropertytype.Name)
}

func (passengerservicetimeslicetype *PassengerServiceTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(passengerservicetimeslicetype.GongGetIdentifier(stage), "PassengerServiceTimeSliceType", passengerservicetimeslicetype.Name)
}

func (passengerservicetype *PassengerServiceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(passengerservicetype.GongGetIdentifier(stage), "PassengerServiceType", passengerservicetype.Name)
}

func (pilotcontrolledlightingtimeslicepropertytype *PilotControlledLightingTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(pilotcontrolledlightingtimeslicepropertytype.GongGetIdentifier(stage), "PilotControlledLightingTimeSlicePropertyType", pilotcontrolledlightingtimeslicepropertytype.Name)
}

func (pilotcontrolledlightingtimeslicetype *PilotControlledLightingTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(pilotcontrolledlightingtimeslicetype.GongGetIdentifier(stage), "PilotControlledLightingTimeSliceType", pilotcontrolledlightingtimeslicetype.Name)
}

func (pilotcontrolledlightingtype *PilotControlledLightingType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(pilotcontrolledlightingtype.GongGetIdentifier(stage), "PilotControlledLightingType", pilotcontrolledlightingtype.Name)
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

func (procedurepropertytype *ProcedurePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(procedurepropertytype.GongGetIdentifier(stage), "ProcedurePropertyType", procedurepropertytype.Name)
}

func (radiocommunicationchannelpropertytype *RadioCommunicationChannelPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(radiocommunicationchannelpropertytype.GongGetIdentifier(stage), "RadioCommunicationChannelPropertyType", radiocommunicationchannelpropertytype.Name)
}

func (radiocommunicationchanneltimeslicepropertytype *RadioCommunicationChannelTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(radiocommunicationchanneltimeslicepropertytype.GongGetIdentifier(stage), "RadioCommunicationChannelTimeSlicePropertyType", radiocommunicationchanneltimeslicepropertytype.Name)
}

func (radiocommunicationchanneltimeslicetype *RadioCommunicationChannelTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(radiocommunicationchanneltimeslicetype.GongGetIdentifier(stage), "RadioCommunicationChannelTimeSliceType", radiocommunicationchanneltimeslicetype.Name)
}

func (radiocommunicationchanneltype *RadioCommunicationChannelType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(radiocommunicationchanneltype.GongGetIdentifier(stage), "RadioCommunicationChannelType", radiocommunicationchanneltype.Name)
}

func (radiocommunicationoperationalstatuspropertytype *RadioCommunicationOperationalStatusPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(radiocommunicationoperationalstatuspropertytype.GongGetIdentifier(stage), "RadioCommunicationOperationalStatusPropertyType", radiocommunicationoperationalstatuspropertytype.Name)
}

func (radiocommunicationoperationalstatustype *RadioCommunicationOperationalStatusType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(radiocommunicationoperationalstatustype.GongGetIdentifier(stage), "RadioCommunicationOperationalStatusType", radiocommunicationoperationalstatustype.Name)
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

func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(runwaycentrelinepointpropertytype.GongGetIdentifier(stage), "RunwayCentrelinePointPropertyType", runwaycentrelinepointpropertytype.Name)
}

func (searchrescueservicetimeslicepropertytype *SearchRescueServiceTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(searchrescueservicetimeslicepropertytype.GongGetIdentifier(stage), "SearchRescueServiceTimeSlicePropertyType", searchrescueservicetimeslicepropertytype.Name)
}

func (searchrescueservicetimeslicetype *SearchRescueServiceTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(searchrescueservicetimeslicetype.GongGetIdentifier(stage), "SearchRescueServiceTimeSliceType", searchrescueservicetimeslicetype.Name)
}

func (searchrescueservicetype *SearchRescueServiceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(searchrescueservicetype.GongGetIdentifier(stage), "SearchRescueServiceType", searchrescueservicetype.Name)
}

func (serviceoperationalstatuspropertytype *ServiceOperationalStatusPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(serviceoperationalstatuspropertytype.GongGetIdentifier(stage), "ServiceOperationalStatusPropertyType", serviceoperationalstatuspropertytype.Name)
}

func (serviceoperationalstatustype *ServiceOperationalStatusType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(serviceoperationalstatustype.GongGetIdentifier(stage), "ServiceOperationalStatusType", serviceoperationalstatustype.Name)
}

func (specialdatetimeslicepropertytype *SpecialDateTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(specialdatetimeslicepropertytype.GongGetIdentifier(stage), "SpecialDateTimeSlicePropertyType", specialdatetimeslicepropertytype.Name)
}

func (specialdatetimeslicetype *SpecialDateTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(specialdatetimeslicetype.GongGetIdentifier(stage), "SpecialDateTimeSliceType", specialdatetimeslicetype.Name)
}

func (specialdatetype *SpecialDateType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(specialdatetype.GongGetIdentifier(stage), "SpecialDateType", specialdatetype.Name)
}

func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(standardlevelcolumnpropertytype.GongGetIdentifier(stage), "StandardLevelColumnPropertyType", standardlevelcolumnpropertytype.Name)
}

func (standardlevelcolumntimeslicepropertytype *StandardLevelColumnTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(standardlevelcolumntimeslicepropertytype.GongGetIdentifier(stage), "StandardLevelColumnTimeSlicePropertyType", standardlevelcolumntimeslicepropertytype.Name)
}

func (standardlevelcolumntimeslicetype *StandardLevelColumnTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(standardlevelcolumntimeslicetype.GongGetIdentifier(stage), "StandardLevelColumnTimeSliceType", standardlevelcolumntimeslicetype.Name)
}

func (standardlevelcolumntype *StandardLevelColumnType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(standardlevelcolumntype.GongGetIdentifier(stage), "StandardLevelColumnType", standardlevelcolumntype.Name)
}

func (standardlevelpropertytype *StandardLevelPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(standardlevelpropertytype.GongGetIdentifier(stage), "StandardLevelPropertyType", standardlevelpropertytype.Name)
}

func (standardlevelsectortimeslicepropertytype *StandardLevelSectorTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(standardlevelsectortimeslicepropertytype.GongGetIdentifier(stage), "StandardLevelSectorTimeSlicePropertyType", standardlevelsectortimeslicepropertytype.Name)
}

func (standardlevelsectortimeslicetype *StandardLevelSectorTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(standardlevelsectortimeslicetype.GongGetIdentifier(stage), "StandardLevelSectorTimeSliceType", standardlevelsectortimeslicetype.Name)
}

func (standardlevelsectortype *StandardLevelSectorType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(standardlevelsectortype.GongGetIdentifier(stage), "StandardLevelSectorType", standardlevelsectortype.Name)
}

func (standardleveltablepropertytype *StandardLevelTablePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(standardleveltablepropertytype.GongGetIdentifier(stage), "StandardLevelTablePropertyType", standardleveltablepropertytype.Name)
}

func (standardleveltabletimeslicepropertytype *StandardLevelTableTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(standardleveltabletimeslicepropertytype.GongGetIdentifier(stage), "StandardLevelTableTimeSlicePropertyType", standardleveltabletimeslicepropertytype.Name)
}

func (standardleveltabletimeslicetype *StandardLevelTableTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(standardleveltabletimeslicetype.GongGetIdentifier(stage), "StandardLevelTableTimeSliceType", standardleveltabletimeslicetype.Name)
}

func (standardleveltabletype *StandardLevelTableType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(standardleveltabletype.GongGetIdentifier(stage), "StandardLevelTableType", standardleveltabletype.Name)
}

func (standardleveltype *StandardLevelType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(standardleveltype.GongGetIdentifier(stage), "StandardLevelType", standardleveltype.Name)
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

func (unitavailabilitypropertytype *UnitAvailabilityPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(unitavailabilitypropertytype.GongGetIdentifier(stage), "UnitAvailabilityPropertyType", unitavailabilitypropertytype.Name)
}

func (unitavailabilitytype *UnitAvailabilityType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(unitavailabilitytype.GongGetIdentifier(stage), "UnitAvailabilityType", unitavailabilitytype.Name)
}

func (unitdependencypropertytype *UnitDependencyPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(unitdependencypropertytype.GongGetIdentifier(stage), "UnitDependencyPropertyType", unitdependencypropertytype.Name)
}

func (unitdependencytype *UnitDependencyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(unitdependencytype.GongGetIdentifier(stage), "UnitDependencyType", unitdependencytype.Name)
}

func (unitpropertytype *UnitPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(unitpropertytype.GongGetIdentifier(stage), "UnitPropertyType", unitpropertytype.Name)
}

func (unittimeslicepropertytype *UnitTimeSlicePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(unittimeslicepropertytype.GongGetIdentifier(stage), "UnitTimeSlicePropertyType", unittimeslicepropertytype.Name)
}

func (unittimeslicetype *UnitTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(unittimeslicetype.GongGetIdentifier(stage), "UnitTimeSliceType", unittimeslicetype.Name)
}

func (unittype *UnitType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(unittype.GongGetIdentifier(stage), "UnitType", unittype.Name)
}

func (vorpropertytype *VORPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(vorpropertytype.GongGetIdentifier(stage), "VORPropertyType", vorpropertytype.Name)
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

func (abstractairportgroundservicetype *AbstractAirportGroundServiceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractairportgroundservicetype.GongGetReferenceIdentifier(stage))
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractpropertieswithscheduletype.GongGetReferenceIdentifier(stage))
}

func (abstractservicetype *AbstractServiceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractservicetype.GongGetReferenceIdentifier(stage))
}

func (abstracttrafficseparationservicetype *AbstractTrafficSeparationServiceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstracttrafficseparationservicetype.GongGetReferenceIdentifier(stage))
}

func (aerialrefuellingpropertytype *AerialRefuellingPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(aerialrefuellingpropertytype.GongGetReferenceIdentifier(stage))
}

func (airtrafficcontrolservicetimeslicepropertytype *AirTrafficControlServiceTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airtrafficcontrolservicetimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (airtrafficcontrolservicetimeslicetype *AirTrafficControlServiceTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airtrafficcontrolservicetimeslicetype.GongGetReferenceIdentifier(stage))
}

func (airtrafficcontrolservicetype *AirTrafficControlServiceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airtrafficcontrolservicetype.GongGetReferenceIdentifier(stage))
}

func (airtrafficmanagementservicetimeslicepropertytype *AirTrafficManagementServiceTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airtrafficmanagementservicetimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (airtrafficmanagementservicetimeslicetype *AirTrafficManagementServiceTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airtrafficmanagementservicetimeslicetype.GongGetReferenceIdentifier(stage))
}

func (airtrafficmanagementservicetype *AirTrafficManagementServiceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airtrafficmanagementservicetype.GongGetReferenceIdentifier(stage))
}

func (aircraftgroundservicetimeslicepropertytype *AircraftGroundServiceTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(aircraftgroundservicetimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (aircraftgroundservicetimeslicetype *AircraftGroundServiceTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(aircraftgroundservicetimeslicetype.GongGetReferenceIdentifier(stage))
}

func (aircraftgroundservicetype *AircraftGroundServiceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(aircraftgroundservicetype.GongGetReferenceIdentifier(stage))
}

func (airportclearanceservicetimeslicepropertytype *AirportClearanceServiceTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airportclearanceservicetimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (airportclearanceservicetimeslicetype *AirportClearanceServiceTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airportclearanceservicetimeslicetype.GongGetReferenceIdentifier(stage))
}

func (airportclearanceservicetype *AirportClearanceServiceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airportclearanceservicetype.GongGetReferenceIdentifier(stage))
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airportheliportpropertytype.GongGetReferenceIdentifier(stage))
}

func (airportsuppliesservicetimeslicepropertytype *AirportSuppliesServiceTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airportsuppliesservicetimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (airportsuppliesservicetimeslicetype *AirportSuppliesServiceTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airportsuppliesservicetimeslicetype.GongGetReferenceIdentifier(stage))
}

func (airportsuppliesservicetype *AirportSuppliesServiceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airportsuppliesservicetype.GongGetReferenceIdentifier(stage))
}

func (airspacepropertytype *AirspacePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(airspacepropertytype.GongGetReferenceIdentifier(stage))
}

func (callsigndetailpropertytype *CallsignDetailPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(callsigndetailpropertytype.GongGetReferenceIdentifier(stage))
}

func (callsigndetailtype *CallsignDetailType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(callsigndetailtype.GongGetReferenceIdentifier(stage))
}

func (contactinformationpropertytype *ContactInformationPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(contactinformationpropertytype.GongGetReferenceIdentifier(stage))
}

func (contactinformationtype *ContactInformationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(contactinformationtype.GongGetReferenceIdentifier(stage))
}

func (designatedpointpropertytype *DesignatedPointPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(designatedpointpropertytype.GongGetReferenceIdentifier(stage))
}

func (directionfinderpropertytype *DirectionFinderPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(directionfinderpropertytype.GongGetReferenceIdentifier(stage))
}

func (elevatedpointpropertytype *ElevatedPointPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(elevatedpointpropertytype.GongGetReferenceIdentifier(stage))
}

func (elevatedpointtype *ElevatedPointType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(elevatedpointtype.GongGetReferenceIdentifier(stage))
}

func (firefightingservicetimeslicepropertytype *FireFightingServiceTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(firefightingservicetimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (firefightingservicetimeslicetype *FireFightingServiceTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(firefightingservicetimeslicetype.GongGetReferenceIdentifier(stage))
}

func (firefightingservicetype *FireFightingServiceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(firefightingservicetype.GongGetReferenceIdentifier(stage))
}

func (fuelpropertytype *FuelPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(fuelpropertytype.GongGetReferenceIdentifier(stage))
}

func (fueltype *FuelType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(fueltype.GongGetReferenceIdentifier(stage))
}

func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(groundlightsystempropertytype.GongGetReferenceIdentifier(stage))
}

func (groundtrafficcontrolservicetimeslicepropertytype *GroundTrafficControlServiceTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(groundtrafficcontrolservicetimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (groundtrafficcontrolservicetimeslicetype *GroundTrafficControlServiceTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(groundtrafficcontrolservicetimeslicetype.GongGetReferenceIdentifier(stage))
}

func (groundtrafficcontrolservicetype *GroundTrafficControlServiceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(groundtrafficcontrolservicetype.GongGetReferenceIdentifier(stage))
}

func (holdingpatternpropertytype *HoldingPatternPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(holdingpatternpropertytype.GongGetReferenceIdentifier(stage))
}

func (informationservicetimeslicepropertytype *InformationServiceTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(informationservicetimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (informationservicetimeslicetype *InformationServiceTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(informationservicetimeslicetype.GongGetReferenceIdentifier(stage))
}

func (informationservicetype *InformationServiceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(informationservicetype.GongGetReferenceIdentifier(stage))
}

func (lightactivationpropertytype *LightActivationPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(lightactivationpropertytype.GongGetReferenceIdentifier(stage))
}

func (lightactivationtype *LightActivationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(lightactivationtype.GongGetReferenceIdentifier(stage))
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(linguisticnotepropertytype.GongGetReferenceIdentifier(stage))
}

func (linguisticnotetype *LinguisticNoteType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(linguisticnotetype.GongGetReferenceIdentifier(stage))
}

func (navaidpropertytype *NavaidPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(navaidpropertytype.GongGetReferenceIdentifier(stage))
}

func (nitrogenpropertytype *NitrogenPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(nitrogenpropertytype.GongGetReferenceIdentifier(stage))
}

func (nitrogentype *NitrogenType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(nitrogentype.GongGetReferenceIdentifier(stage))
}

func (notepropertytype *NotePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(notepropertytype.GongGetReferenceIdentifier(stage))
}

func (notetype *NoteType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(notetype.GongGetReferenceIdentifier(stage))
}

func (oilpropertytype *OilPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(oilpropertytype.GongGetReferenceIdentifier(stage))
}

func (oiltype *OilType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(oiltype.GongGetReferenceIdentifier(stage))
}

func (onlinecontactpropertytype *OnlineContactPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(onlinecontactpropertytype.GongGetReferenceIdentifier(stage))
}

func (onlinecontacttype *OnlineContactType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(onlinecontacttype.GongGetReferenceIdentifier(stage))
}

func (organisationauthorityassociationpropertytype *OrganisationAuthorityAssociationPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(organisationauthorityassociationpropertytype.GongGetReferenceIdentifier(stage))
}

func (organisationauthorityassociationtype *OrganisationAuthorityAssociationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(organisationauthorityassociationtype.GongGetReferenceIdentifier(stage))
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(organisationauthoritypropertytype.GongGetReferenceIdentifier(stage))
}

func (organisationauthoritytimeslicepropertytype *OrganisationAuthorityTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(organisationauthoritytimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (organisationauthoritytimeslicetype *OrganisationAuthorityTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(organisationauthoritytimeslicetype.GongGetReferenceIdentifier(stage))
}

func (organisationauthoritytype *OrganisationAuthorityType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(organisationauthoritytype.GongGetReferenceIdentifier(stage))
}

func (oxygenpropertytype *OxygenPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(oxygenpropertytype.GongGetReferenceIdentifier(stage))
}

func (oxygentype *OxygenType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(oxygentype.GongGetReferenceIdentifier(stage))
}

func (passengerservicetimeslicepropertytype *PassengerServiceTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(passengerservicetimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (passengerservicetimeslicetype *PassengerServiceTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(passengerservicetimeslicetype.GongGetReferenceIdentifier(stage))
}

func (passengerservicetype *PassengerServiceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(passengerservicetype.GongGetReferenceIdentifier(stage))
}

func (pilotcontrolledlightingtimeslicepropertytype *PilotControlledLightingTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(pilotcontrolledlightingtimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (pilotcontrolledlightingtimeslicetype *PilotControlledLightingTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(pilotcontrolledlightingtimeslicetype.GongGetReferenceIdentifier(stage))
}

func (pilotcontrolledlightingtype *PilotControlledLightingType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(pilotcontrolledlightingtype.GongGetReferenceIdentifier(stage))
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

func (procedurepropertytype *ProcedurePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(procedurepropertytype.GongGetReferenceIdentifier(stage))
}

func (radiocommunicationchannelpropertytype *RadioCommunicationChannelPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(radiocommunicationchannelpropertytype.GongGetReferenceIdentifier(stage))
}

func (radiocommunicationchanneltimeslicepropertytype *RadioCommunicationChannelTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(radiocommunicationchanneltimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (radiocommunicationchanneltimeslicetype *RadioCommunicationChannelTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(radiocommunicationchanneltimeslicetype.GongGetReferenceIdentifier(stage))
}

func (radiocommunicationchanneltype *RadioCommunicationChannelType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(radiocommunicationchanneltype.GongGetReferenceIdentifier(stage))
}

func (radiocommunicationoperationalstatuspropertytype *RadioCommunicationOperationalStatusPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(radiocommunicationoperationalstatuspropertytype.GongGetReferenceIdentifier(stage))
}

func (radiocommunicationoperationalstatustype *RadioCommunicationOperationalStatusType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(radiocommunicationoperationalstatustype.GongGetReferenceIdentifier(stage))
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

func (runwaycentrelinepointpropertytype *RunwayCentrelinePointPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(runwaycentrelinepointpropertytype.GongGetReferenceIdentifier(stage))
}

func (searchrescueservicetimeslicepropertytype *SearchRescueServiceTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(searchrescueservicetimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (searchrescueservicetimeslicetype *SearchRescueServiceTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(searchrescueservicetimeslicetype.GongGetReferenceIdentifier(stage))
}

func (searchrescueservicetype *SearchRescueServiceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(searchrescueservicetype.GongGetReferenceIdentifier(stage))
}

func (serviceoperationalstatuspropertytype *ServiceOperationalStatusPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(serviceoperationalstatuspropertytype.GongGetReferenceIdentifier(stage))
}

func (serviceoperationalstatustype *ServiceOperationalStatusType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(serviceoperationalstatustype.GongGetReferenceIdentifier(stage))
}

func (specialdatetimeslicepropertytype *SpecialDateTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(specialdatetimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (specialdatetimeslicetype *SpecialDateTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(specialdatetimeslicetype.GongGetReferenceIdentifier(stage))
}

func (specialdatetype *SpecialDateType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(specialdatetype.GongGetReferenceIdentifier(stage))
}

func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(standardlevelcolumnpropertytype.GongGetReferenceIdentifier(stage))
}

func (standardlevelcolumntimeslicepropertytype *StandardLevelColumnTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(standardlevelcolumntimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (standardlevelcolumntimeslicetype *StandardLevelColumnTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(standardlevelcolumntimeslicetype.GongGetReferenceIdentifier(stage))
}

func (standardlevelcolumntype *StandardLevelColumnType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(standardlevelcolumntype.GongGetReferenceIdentifier(stage))
}

func (standardlevelpropertytype *StandardLevelPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(standardlevelpropertytype.GongGetReferenceIdentifier(stage))
}

func (standardlevelsectortimeslicepropertytype *StandardLevelSectorTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(standardlevelsectortimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (standardlevelsectortimeslicetype *StandardLevelSectorTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(standardlevelsectortimeslicetype.GongGetReferenceIdentifier(stage))
}

func (standardlevelsectortype *StandardLevelSectorType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(standardlevelsectortype.GongGetReferenceIdentifier(stage))
}

func (standardleveltablepropertytype *StandardLevelTablePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(standardleveltablepropertytype.GongGetReferenceIdentifier(stage))
}

func (standardleveltabletimeslicepropertytype *StandardLevelTableTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(standardleveltabletimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (standardleveltabletimeslicetype *StandardLevelTableTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(standardleveltabletimeslicetype.GongGetReferenceIdentifier(stage))
}

func (standardleveltabletype *StandardLevelTableType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(standardleveltabletype.GongGetReferenceIdentifier(stage))
}

func (standardleveltype *StandardLevelType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(standardleveltype.GongGetReferenceIdentifier(stage))
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

func (unitavailabilitypropertytype *UnitAvailabilityPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(unitavailabilitypropertytype.GongGetReferenceIdentifier(stage))
}

func (unitavailabilitytype *UnitAvailabilityType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(unitavailabilitytype.GongGetReferenceIdentifier(stage))
}

func (unitdependencypropertytype *UnitDependencyPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(unitdependencypropertytype.GongGetReferenceIdentifier(stage))
}

func (unitdependencytype *UnitDependencyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(unitdependencytype.GongGetReferenceIdentifier(stage))
}

func (unitpropertytype *UnitPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(unitpropertytype.GongGetReferenceIdentifier(stage))
}

func (unittimeslicepropertytype *UnitTimeSlicePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(unittimeslicepropertytype.GongGetReferenceIdentifier(stage))
}

func (unittimeslicetype *UnitTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(unittimeslicetype.GongGetReferenceIdentifier(stage))
}

func (unittype *UnitType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(unittype.GongGetReferenceIdentifier(stage))
}

func (vorpropertytype *VORPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(vorpropertytype.GongGetReferenceIdentifier(stage))
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
