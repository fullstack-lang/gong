package models

type Documentation struct {
	Name   string
	Text   string `xml:",chardata"`
	Source string `xml:"source,attr"`
	Lang   string `xml:"lang,attr"`
}
