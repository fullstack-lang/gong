package models

type OccurrenceDefinitionAbstract struct {
	MinOccurs string `xml:"minOccurs,attr"`
	MaxOccurs string `xml:"maxOccurs,attr"`
}

type OccurrenceDefinition interface {
	GetIsLocalyUnbounded() bool
}

func (occurrenceDefinitionAbstract *OccurrenceDefinitionAbstract) GetIsLocalyUnbounded() bool {

	return occurrenceDefinitionAbstract.MaxOccurs == "unbounded"
}
