package models

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Relationship is
type Relationship struct {
	Text       string `xml:",chardata"`
	ID         string `xml:"Id,attr"`
	Type       string `xml:"Type,attr"`
	Target     string `xml:"Target,attr"`
	TargetMode string `xml:"TargetMode,attr"`
}

// Relationships is
type Relationships struct {
	XMLName      xml.Name       `xml:"Relationships"`
	Text         string         `xml:",chardata"`
	Xmlns        string         `xml:"xmlns,attr"`
	Relationship []Relationship `xml:"Relationship"`
}

// TextVal is
type TextVal struct {
	Text string `xml:",chardata"`
	Val  string `xml:"val,attr"`
}

// NumberingLvl is
type NumberingLvl struct {
	Text      string  `xml:",chardata"`
	Ilvl      string  `xml:"ilvl,attr"`
	Tplc      string  `xml:"tplc,attr"`
	Tentative string  `xml:"tentative,attr"`
	Start     TextVal `xml:"start"`
	NumFmt    TextVal `xml:"numFmt"`
	LvlText   TextVal `xml:"lvlText"`
	LvlJc     TextVal `xml:"lvlJc"`
	PPr       struct {
		Text string `xml:",chardata"`
		Ind  struct {
			Text    string `xml:",chardata"`
			Left    string `xml:"left,attr"`
			Hanging string `xml:"hanging,attr"`
		} `xml:"ind"`
	} `xml:"pPr"`
	RPr struct {
		Text string `xml:",chardata"`
		U    struct {
			Text string `xml:",chardata"`
			Val  string `xml:"val,attr"`
		} `xml:"u"`
		RFonts struct {
			Text string `xml:",chardata"`
			Hint string `xml:"hint,attr"`
		} `xml:"rFonts"`
	} `xml:"rPr"`
}

// Numbering is
type Numbering struct {
	XMLName     xml.Name `xml:"numbering"`
	Text        string   `xml:",chardata"`
	Wpc         string   `xml:"wpc,attr"`
	Cx          string   `xml:"cx,attr"`
	Cx1         string   `xml:"cx1,attr"`
	Mc          string   `xml:"mc,attr"`
	O           string   `xml:"o,attr"`
	R           string   `xml:"r,attr"`
	M           string   `xml:"m,attr"`
	V           string   `xml:"v,attr"`
	Wp14        string   `xml:"wp14,attr"`
	Wp          string   `xml:"wp,attr"`
	W10         string   `xml:"w10,attr"`
	W           string   `xml:"w,attr"`
	W14         string   `xml:"w14,attr"`
	W15         string   `xml:"w15,attr"`
	W16se       string   `xml:"w16se,attr"`
	Wpg         string   `xml:"wpg,attr"`
	Wpi         string   `xml:"wpi,attr"`
	Wne         string   `xml:"wne,attr"`
	Wps         string   `xml:"wps,attr"`
	Ignorable   string   `xml:"Ignorable,attr"`
	AbstractNum []struct {
		Text                       string         `xml:",chardata"`
		AbstractNumID              string         `xml:"abstractNumId,attr"`
		RestartNumberingAfterBreak string         `xml:"restartNumberingAfterBreak,attr"`
		Nsid                       TextVal        `xml:"nsid"`
		MultiLevelType             TextVal        `xml:"multiLevelType"`
		Tmpl                       TextVal        `xml:"tmpl"`
		Lvl                        []NumberingLvl `xml:"lvl"`
	} `xml:"abstractNum"`
	Num []struct {
		Text          string  `xml:",chardata"`
		NumID         string  `xml:"numId,attr"`
		AbstractNumID TextVal `xml:"abstractNumId"`
	} `xml:"num"`
}

type file struct {
	rels  Relationships
	num   Numbering
	r     *zip.ReadCloser
	embed bool
	list  map[string]int
}

// Node_ is
type Node_ struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:"-"`
	Content []byte     `xml:",innerxml"`
	Nodes   []Node_    `xml:",any"`
}

// UnmarshalXML is
func (n *Node_) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	n.Attrs = start.Attr
	type node Node_

	return d.DecodeElement((*node)(n), &start)
}

func escape(s, set string) string {
	replacer := []string{}
	for _, r := range []rune(set) {
		rs := string(r)
		replacer = append(replacer, rs, `\`+rs)
	}
	return strings.NewReplacer(replacer...).Replace(s)
}

func (zf *file) extract(rel *Relationship, w io.Writer) error {
	err := os.MkdirAll(filepath.Dir(rel.Target), 0755)
	if err != nil {
		return err
	}
	for _, f := range zf.r.File {
		if f.Name != "word/"+rel.Target {
			continue
		}
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		b := make([]byte, f.UncompressedSize64)
		n, err := rc.Read(b)
		if err != nil && err != io.EOF {
			return err
		}
		if zf.embed {
			fmt.Fprintf(w, "![](data:image/png;base64,%s)",
				base64.StdEncoding.EncodeToString(b[:n]))
		} else {
			err = ioutil.WriteFile(rel.Target, b, 0644)
			if err != nil {
				return err
			}
			fmt.Fprintf(w, "![](%s)", escape(rel.Target, "()"))
		}
		break
	}
	return nil
}

func attr(attrs []xml.Attr, name string) (string, bool) {
	for _, attr := range attrs {
		if attr.Name.Local == name {
			return attr.Value, true
		}
	}
	return "", false
}

func readFile(f *zip.File) (*Node_, error) {
	rc, err := f.Open()
	defer rc.Close()

	b, _ := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}

	var node Node_
	err = xml.Unmarshal(b, &node)
	if err != nil {
		return nil, err
	}
	return &node, nil
}

func findFile(files []*zip.File, target string) *zip.File {
	for _, f := range files {
		if ok, _ := path.Match(target, f.Name); ok {
			return f
		}
	}
	return nil
}

func docx2md(docx *Docx, gongdocx_stage *Stage, path string, embed bool) error {
	r, err := zip.OpenReader(path)
	if err != nil {
		return err
	}
	defer r.Close()

	var rels Relationships
	var num Numbering

	for _, f := range r.File {

		file := (&File{Name: f.Name}).Stage(gongdocx_stage)
		docx.Files = append(docx.Files, file)

		switch f.Name {
		case "word/_rels/document.xml.rels":
			rc, err := f.Open()
			defer rc.Close()

			b, _ := ioutil.ReadAll(rc)
			if err != nil {
				return err
			}

			err = xml.Unmarshal(b, &rels)
			if err != nil {
				return err
			}
		case "word/numbering.xml":
			rc, err := f.Open()
			defer rc.Close()

			b, _ := ioutil.ReadAll(rc)
			if err != nil {
				return err
			}

			err = xml.Unmarshal(b, &num)
			if err != nil {
				return err
			}
		}
	}

	f := findFile(r.File, "word/document*.xml")
	if f == nil {
		return errors.New("incorrect document")
	}

	// match to the File
	document := (&Document{Name: f.Name}).Stage(gongdocx_stage)
	docx.Document = document

	file_ := (*GetGongstructInstancesMap[File](gongdocx_stage))["word/document.xml"]
	document.File = file_

	node_, err := readFile(f)
	if err != nil {
		return err
	}
	node := (&Node{Name: "0"}).Stage(gongdocx_stage)
	document.Root = node

	var buf bytes.Buffer
	zf := &file{
		r:     r,
		rels:  rels,
		num:   num,
		embed: embed,
		list:  make(map[string]int),
	}

	body := (&Body{Name: f.Name}).Stage(gongdocx_stage)
	document.Body = body
	err = walk(zf, body, node, gongdocx_stage, node_, &buf)
	if err != nil {
		return err
	}
	// fmt.Print(buf.String())

	// Create a new ZIP archive in memory
	newZipBuf := new(bytes.Buffer)
	newZipWriter := zip.NewWriter(newZipBuf)

	// Iterate through each file in the original ZIP
	for _, f := range r.File {
		zipFileReader, err := f.Open()
		if err != nil {
			return err
		}

		if f.Name != "word/document.xml" {
			newFile, err := newZipWriter.Create(f.Name)
			if err != nil {
				zipFileReader.Close()
				return err
			}
			_, err = io.Copy(newFile, zipFileReader)
			if err != nil {
				zipFileReader.Close()
				return err
			}
		} else {
			// If the current file is the one to modify, write the new content
			newFile, err := newZipWriter.Create("word/document.xml")
			if err != nil {
				zipFileReader.Close()
				return err
			}
			_, err = newFile.Write([]byte("Toto"))
			if err != nil {
				zipFileReader.Close()
				return err
			}
		}

		zipFileReader.Close()
	}

	// Close the new ZIP writer to finish writing the new ZIP file
	if err := newZipWriter.Close(); err != nil {
		return err
	}

	// Replace the original ZIP file with the new ZIP archive
	// Trim the ".docx" extension from the filePath
	baseNameWithoutExt := strings.TrimSuffix(path, ".docx")

	// Append "copy" before the extension
	newFilePath := baseNameWithoutExt + "-copy.docx"

	if err := ioutil.WriteFile(newFilePath, newZipBuf.Bytes(), 0644); err != nil {
		return err
	}

	return nil
}
