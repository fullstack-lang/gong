package models

func (stager *Stager) enforceModelSemantic() {

	// all reqif object have an unique identifier. Those runtime map
	// allow navigation to object from the identifier

	stager.Map_id_SPEC_OBJECT_TYPE = populateIdMap[*SPEC_OBJECT_TYPE](stager)
	stager.Map_id_SPECIFICATION_TYPE = populateIdMap[*SPECIFICATION_TYPE](stager)
	stager.Map_id_SPEC_OBJECT = populateIdMap[*SPEC_OBJECT](stager)

	stager.Map_id_ENUM_VALUE = populateIdMap[*ENUM_VALUE](stager)
	stager.Map_id_SPEC_RELATION_TYPE = populateIdMap[*SPEC_RELATION_TYPE](stager)

	stager.Map_id_DATATYPE_DEFINITION_XHTML = populateIdMap[*DATATYPE_DEFINITION_XHTML](stager)
	stager.Map_id_DATATYPE_DEFINITION_STRING = populateIdMap[*DATATYPE_DEFINITION_STRING](stager)
	stager.Map_id_DATATYPE_DEFINITION_BOOLEAN = populateIdMap[*DATATYPE_DEFINITION_BOOLEAN](stager)
	stager.Map_id_DATATYPE_DEFINITION_INTEGER = populateIdMap[*DATATYPE_DEFINITION_INTEGER](stager)
	stager.Map_id_DATATYPE_DEFINITION_REAL = populateIdMap[*DATATYPE_DEFINITION_REAL](stager)
	stager.Map_id_DATATYPE_DEFINITION_DATE = populateIdMap[*DATATYPE_DEFINITION_DATE](stager)
	stager.Map_id_DATATYPE_DEFINITION_ENUMERATION = populateIdMap[*DATATYPE_DEFINITION_ENUMERATION](stager)

	stager.Map_id_ATTRIBUTE_DEFINITION_XHTML = populateIdMap[*ATTRIBUTE_DEFINITION_XHTML](stager)
	stager.Map_id_ATTRIBUTE_DEFINITION_STRING = populateIdMap[*ATTRIBUTE_DEFINITION_STRING](stager)
	stager.Map_id_ATTRIBUTE_DEFINITION_BOOLEAN = populateIdMap[*ATTRIBUTE_DEFINITION_BOOLEAN](stager)
	stager.Map_id_ATTRIBUTE_DEFINITION_INTEGER = populateIdMap[*ATTRIBUTE_DEFINITION_INTEGER](stager)
	stager.Map_id_ATTRIBUTE_DEFINITION_DATE = populateIdMap[*ATTRIBUTE_DEFINITION_DATE](stager)
	stager.Map_id_ATTRIBUTE_DEFINITION_REAL = populateIdMap[*ATTRIBUTE_DEFINITION_REAL](stager)
	stager.Map_id_ATTRIBUTE_DEFINITION_ENUMERATION = populateIdMap[*ATTRIBUTE_DEFINITION_ENUMERATION](stager)

	// the stager also maintain a map used to navigate spec_relations between objects
	stager.Map_SPEC_OBJECT_relations_sources = make(map[*SPEC_OBJECT][]*SPEC_RELATION)
	stager.Map_SPEC_OBJECT_relations_targets = make(map[*SPEC_OBJECT][]*SPEC_RELATION)

	stager.initMap_Objects_Relations()
	stager.enforceRenderingConfigurationSemantic()

}

// populateIdMap fetches instances of a Gongstruct type T and populates a map
// where keys are string identifiers and values are pointers to the instances.
// T must satisfy the Identifiable interface.
func populateIdMap[T Identifiable](stager *Stager) map[string]T {
	instances := *GetGongstructInstancesSetFromPointerType[T](stager.GetStage())
	resultMap := make(map[string]T)
	for instance := range instances {
		resultMap[instance.GetIdentifier()] = instance // Correctly calls the method
	}
	return resultMap
}
