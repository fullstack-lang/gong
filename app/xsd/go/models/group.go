package models

import "encoding/xml"

type Group struct {
	Name string
	Annotated
	ElementWithNameAttribute
	Ref string `xml:"ref,attr"`

	// analysis
	IsAnonymous  bool // it has been defined by the enclosing element
	OuterElement *Element
	WithGoIdentifier

	ModelGroup
}

func (group *Group) GetFields(stage *Stage) (fields string) {
	setOfFieldGoIdentifiers := make(map[string]any)

	stMap := make(map[string]*SimpleType)
	for st := range *GetGongstructInstancesSet[SimpleType](stage) {
		stMap[st.Name] = st
	}
	ctMap := make(map[string]*ComplexType)
	for st := range *GetGongstructInstancesSet[ComplexType](stage) {
		ctMap[st.Name] = st
	}
	agMap := make(map[string]*AttributeGroup)
	for ag := range *GetGongstructInstancesSet[AttributeGroup](stage) {
		agMap[ag.Name] = ag
	}
	groupMap := make(map[string]*Group)
	for group := range *GetGongstructInstancesSet[Group](stage) {
		groupMap[group.Name] = group
	}

	map_Name_Elems := make(map[string]*Element)

	group.ModelGroup.generateElements(map_Name_Elems, stMap, ctMap, groupMap, setOfFieldGoIdentifiers, &fields)

	return
}

func (group *Group) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	group.Order = Order
	group.Depth = Depth
	Order = Order + 1
	Depth = Depth + 1

	type Alias Group
	aux := (*Alias)(group)

	err := d.DecodeElement(aux, &start)
	Depth = Depth - 1

	return err
}
