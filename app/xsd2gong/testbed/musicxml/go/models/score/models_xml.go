// generated code by xsd2gong - do not edit
package score

import (
	"encoding/xml"
	"io"

	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/attributes"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/barline"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/common"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/direction"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/note"
)

// UnmarshalXML streams attributes and child elements in exact document order for Measure.
func (x *Measure) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "number":
			x.Number = attr.Value
		case "text":
			x.Text = attr.Value
		case "implicit":
			x.Implicit = common.EnumYesNo(attr.Value)
		case "non-controlling":
			x.NonControlling = common.EnumYesNo(attr.Value)
		case "width":
			x.Width = attr.Value
		case "id":
			x.Id = attr.Value
		}
	}

	for {
		tok, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		switch se := tok.(type) {
		case xml.StartElement:
			switch se.Name.Local {
			case "note":
				item := new(note.Note)
				if err := d.DecodeElement(item, &se); err != nil {
					return err
				}
				x.MusicData = append(x.MusicData, &MeasureMusicDataChoice{Note: item})
			case "backup":
				item := new(note.Backup)
				if err := d.DecodeElement(item, &se); err != nil {
					return err
				}
				x.MusicData = append(x.MusicData, &MeasureMusicDataChoice{Backup: item})
			case "forward":
				item := new(note.Forward)
				if err := d.DecodeElement(item, &se); err != nil {
					return err
				}
				x.MusicData = append(x.MusicData, &MeasureMusicDataChoice{Forward: item})
			case "direction":
				item := new(direction.Direction)
				if err := d.DecodeElement(item, &se); err != nil {
					return err
				}
				x.MusicData = append(x.MusicData, &MeasureMusicDataChoice{Direction: item})
			case "attributes":
				item := new(attributes.Attributes)
				if err := d.DecodeElement(item, &se); err != nil {
					return err
				}
				x.MusicData = append(x.MusicData, &MeasureMusicDataChoice{Attributes: item})
			case "harmony":
				item := new(direction.Harmony)
				if err := d.DecodeElement(item, &se); err != nil {
					return err
				}
				x.MusicData = append(x.MusicData, &MeasureMusicDataChoice{Harmony: item})
			case "figured-bass":
				item := new(note.FiguredBass)
				if err := d.DecodeElement(item, &se); err != nil {
					return err
				}
				x.MusicData = append(x.MusicData, &MeasureMusicDataChoice{FiguredBass: item})
			case "print":
				item := new(direction.Print)
				if err := d.DecodeElement(item, &se); err != nil {
					return err
				}
				x.MusicData = append(x.MusicData, &MeasureMusicDataChoice{Print: item})
			case "sound":
				item := new(direction.Sound)
				if err := d.DecodeElement(item, &se); err != nil {
					return err
				}
				x.MusicData = append(x.MusicData, &MeasureMusicDataChoice{Sound: item})
			case "listening":
				item := new(direction.Listening)
				if err := d.DecodeElement(item, &se); err != nil {
					return err
				}
				x.MusicData = append(x.MusicData, &MeasureMusicDataChoice{Listening: item})
			case "barline":
				item := new(barline.Barline)
				if err := d.DecodeElement(item, &se); err != nil {
					return err
				}
				x.MusicData = append(x.MusicData, &MeasureMusicDataChoice{Barline: item})
			case "grouping":
				item := new(direction.Grouping)
				if err := d.DecodeElement(item, &se); err != nil {
					return err
				}
				x.MusicData = append(x.MusicData, &MeasureMusicDataChoice{Grouping: item})
			case "link":
				item := new(common.Link)
				if err := d.DecodeElement(item, &se); err != nil {
					return err
				}
				x.MusicData = append(x.MusicData, &MeasureMusicDataChoice{Link: item})
			case "bookmark":
				item := new(common.Bookmark)
				if err := d.DecodeElement(item, &se); err != nil {
					return err
				}
				x.MusicData = append(x.MusicData, &MeasureMusicDataChoice{Bookmark: item})
			default:
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			if se.Name.Local == start.Name.Local {
				return nil
			}
		}
	}
	return nil
}

// MarshalXML serializes attributes and order-preserving choice items in exact document order for Measure.
func (x *Measure) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if x.Number != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "number"}, Value: string(x.Number)})
	}
	if x.Text != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "text"}, Value: string(x.Text)})
	}
	if x.Implicit != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "implicit"}, Value: string(x.Implicit)})
	}
	if x.NonControlling != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "non-controlling"}, Value: string(x.NonControlling)})
	}
	if x.Width != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "width"}, Value: string(x.Width)})
	}
	if x.Id != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"}, Value: string(x.Id)})
	}
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	for _, c := range x.MusicData {
		if c == nil {
			continue
		}
		if c.Note != nil {
			if err := e.EncodeElement(c.Note, xml.StartElement{Name: xml.Name{Local: "note"}}); err != nil {
				return err
			}
		}
		if c.Backup != nil {
			if err := e.EncodeElement(c.Backup, xml.StartElement{Name: xml.Name{Local: "backup"}}); err != nil {
				return err
			}
		}
		if c.Forward != nil {
			if err := e.EncodeElement(c.Forward, xml.StartElement{Name: xml.Name{Local: "forward"}}); err != nil {
				return err
			}
		}
		if c.Direction != nil {
			if err := e.EncodeElement(c.Direction, xml.StartElement{Name: xml.Name{Local: "direction"}}); err != nil {
				return err
			}
		}
		if c.Attributes != nil {
			if err := e.EncodeElement(c.Attributes, xml.StartElement{Name: xml.Name{Local: "attributes"}}); err != nil {
				return err
			}
		}
		if c.Harmony != nil {
			if err := e.EncodeElement(c.Harmony, xml.StartElement{Name: xml.Name{Local: "harmony"}}); err != nil {
				return err
			}
		}
		if c.FiguredBass != nil {
			if err := e.EncodeElement(c.FiguredBass, xml.StartElement{Name: xml.Name{Local: "figured-bass"}}); err != nil {
				return err
			}
		}
		if c.Print != nil {
			if err := e.EncodeElement(c.Print, xml.StartElement{Name: xml.Name{Local: "print"}}); err != nil {
				return err
			}
		}
		if c.Sound != nil {
			if err := e.EncodeElement(c.Sound, xml.StartElement{Name: xml.Name{Local: "sound"}}); err != nil {
				return err
			}
		}
		if c.Listening != nil {
			if err := e.EncodeElement(c.Listening, xml.StartElement{Name: xml.Name{Local: "listening"}}); err != nil {
				return err
			}
		}
		if c.Barline != nil {
			if err := e.EncodeElement(c.Barline, xml.StartElement{Name: xml.Name{Local: "barline"}}); err != nil {
				return err
			}
		}
		if c.Grouping != nil {
			if err := e.EncodeElement(c.Grouping, xml.StartElement{Name: xml.Name{Local: "grouping"}}); err != nil {
				return err
			}
		}
		if c.Link != nil {
			if err := e.EncodeElement(c.Link, xml.StartElement{Name: xml.Name{Local: "link"}}); err != nil {
				return err
			}
		}
		if c.Bookmark != nil {
			if err := e.EncodeElement(c.Bookmark, xml.StartElement{Name: xml.Name{Local: "bookmark"}}); err != nil {
				return err
			}
		}
	}
	return e.EncodeToken(start.End())
}

