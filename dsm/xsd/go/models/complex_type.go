package models

import (
	"encoding/xml"
	"log"
)

type ComplexType struct {
	Name string

	WithGoIdentifier

	// analysis
	IsAnonymous  bool
	OuterElement *Element // if anonymous

	Annotated
	ElementWithNameAttribute

	ModelGroup

	Extension *Extension `xml:"extension"`

	// xs:simpleContent element is used exclusively within an xs:complexType element
	// to define complex types that contain simple content along with
	// attributes or restrictions.
	SimpleContent *SimpleContent `xml:"simpleContent"`

	ComplexContent *ComplexContent `xml:"complexContent"`

	Attributes      []*Attribute      `xml:"attribute"`
	AttributeGroups []*AttributeGroup `xml:"attributeGroup"`

	IsDuplicatedInXSD bool

	OuterParticleOwnerAbstract
}

func (ct *ComplexType) GetFields(stage *Stage) (fields string) {
	setOfFieldsGoIdentifiers := make(map[string]any)

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

	generateAttributes(ct.Attributes, stMap, setOfFieldsGoIdentifiers, &fields)
	for _, referencedAg := range ct.AttributeGroups {

		if namedAg, ok := agMap[referencedAg.Ref]; ok {
			fields += "\n\n\t// generated from attribute group \"" + referencedAg.Ref +
				"\n\t" + namedAg.GoIdentifier
		} else {
			log.Fatalln("Unkown attribute group", referencedAg.Ref)
		}
	}

	if ct.SimpleContent != nil {
		if ct.SimpleContent.Extension != nil {
			generateAttributes(ct.SimpleContent.Extension.Attributes, stMap, setOfFieldsGoIdentifiers, &fields)
			for _, referencedAg := range ct.SimpleContent.Extension.AttributeGroups {

				if namedAg, ok := agMap[referencedAg.Ref]; ok {
					fields += "\n\n\t// generated from attribute group \"" + referencedAg.Ref +
						"\n\t" + namedAg.GoIdentifier
				} else {
					log.Fatalln("Unkown attribute group", referencedAg.Ref)
				}
			}
			// remove namespace from type
			if NsPrefix(ct.SimpleContent.Extension.Base) != "" {
				ct.SimpleContent.Extension.Base = Name(ct.SimpleContent.Extension.Base)
			}

			// in case the extension has base type "xs:string", one has to had the chardata stuff
			goType := generateGoTypeFromType(ct.SimpleContent.Extension.Base, stMap)
			fields += "\n\n\t// in case the extension has base type xs:string, one has to had the chardata stuff" +
				"\n\t" + "EnclosedText" + " " + goType + " " + "`" + `xml:",chardata"` + "`"
		}
	}

	map_Name_Elems := make(map[string]*Element)

	ct.ModelGroup.OuterElementName = ct.Name
	ct.ModelGroup.generateElements(map_Name_Elems, stMap, ctMap, groupMap, setOfFieldsGoIdentifiers, &fields)

	return
}

func (e *ComplexType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	e.Order = Order
	e.Depth = Depth
	Order = Order + 1

	type Alias ComplexType
	aux := (*Alias)(e)

	Depth = Depth + 1
	err := d.DecodeElement(aux, &start)
	Depth = Depth - 1

	return err
}
