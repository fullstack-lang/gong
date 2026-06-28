package xsd

import "encoding/xml"

// ComplexType was generated 2024-08-17 17:02:29 by https://xml-to-go.github.io/ in Ukraine.
type ComplexType struct {
	XMLName        xml.Name `xml:"complexType"`
	Text           string   `xml:",chardata"`
	ComplexContent struct {
		Text      string `xml:",chardata"`
		Extension struct {
			Text     string `xml:",chardata"`
			Base     string `xml:"base,attr"`
			Sequence struct {
				Text  string `xml:",chardata"`
				Group struct {
					Text      string `xml:",chardata"`
					Ref       string `xml:"ref,attr"`
					MinOccurs string `xml:"minOccurs,attr"`
					MaxOccurs string `xml:"maxOccurs,attr"`
				} `xml:"group"`
				Sequence []struct {
					Text      string `xml:",chardata"`
					MinOccurs string `xml:"minOccurs,attr"`
					MaxOccurs string `xml:"maxOccurs,attr"`
					Element   []struct {
						Text      string `xml:",chardata"`
						Ref       string `xml:"ref,attr"`
						MinOccurs string `xml:"minOccurs,attr"`
						MaxOccurs string `xml:"maxOccurs,attr"`
					} `xml:"element"`
					Group struct {
						Text string `xml:",chardata"`
						Ref  string `xml:"ref,attr"`
					} `xml:"group"`
				} `xml:"sequence"`
			} `xml:"sequence"`
			Attribute []struct {
				Text    string `xml:",chardata"`
				Name    string `xml:"name,attr"`
				Type    string `xml:"type,attr"`
				Default string `xml:"default,attr"`
				Use     string `xml:"use,attr"`
				Ref     string `xml:"ref,attr"`
			} `xml:"attribute"`
		} `xml:"extension"`
	} `xml:"complexContent"`
}
