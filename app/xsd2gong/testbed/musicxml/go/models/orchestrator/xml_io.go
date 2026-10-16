// generated code by xsd2gong - do not edit
package orchestrator

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf16"
	"unicode/utf8"

	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/score"
)

func normalizeXMLBytes(data []byte) []byte {
	// Strip UTF-8 BOM if present
	if len(data) >= 3 && data[0] == 0xEF && data[1] == 0xBB && data[2] == 0xBF {
		data = data[3:]
	}
	// Convert UTF-16BE or UTF-16LE with BOM to UTF-8
	if len(data) >= 2 && ((data[0] == 0xFE && data[1] == 0xFF) || (data[0] == 0xFF && data[1] == 0xFE)) {
		bigEndian := data[0] == 0xFE
		raw := data[2:]
		u16 := make([]uint16, len(raw)/2)
		for i := range u16 {
			if bigEndian {
				u16[i] = uint16(raw[2*i])<<8 | uint16(raw[2*i+1])
			} else {
				u16[i] = uint16(raw[2*i+1])<<8 | uint16(raw[2*i])
			}
		}
		runes := utf16.Decode(u16)
		var buf bytes.Buffer
		buf.Grow(len(runes))
		var b [4]byte
		for _, r := range runes {
			n := utf8.EncodeRune(b[:], r)
			buf.Write(b[:n])
		}
		str := strings.Replace(buf.String(), "encoding=\"UTF-16\"", "encoding=\"UTF-8\"", 1)
		str = strings.Replace(str, "encoding='UTF-16'", "encoding='UTF-8'", 1)
		return []byte(str)
	}
	return data
}

// UnmarshalScorePartwiseFile reads an XML file and unmarshals it into a multi-stage Gong model hierarchy.
func UnmarshalScorePartwiseFile(path string) (*MusicXMLStages, *score.ScorePartwise, []error, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, nil, err
	}
	return UnmarshalScorePartwise(data)
}

// UnmarshalScorePartwise parses XML bytes into a multi-stage Gong model hierarchy,
// stages all instances across their respective package stages, resolves ID/IDREF pointers,
// and executes semantic validation across all stages.
func UnmarshalScorePartwise(xmlData []byte) (*MusicXMLStages, *score.ScorePartwise, []error, error) {
	xmlData = normalizeXMLBytes(xmlData)
	stages := NewMusicXMLStages()
	root := new(score.ScorePartwise)
	if err := xml.Unmarshal(xmlData, root); err != nil {
		return nil, nil, nil, fmt.Errorf("xml unmarshal failed: %w", err)
	}

	// Stage all instantiated structs into their respective package stages
	stages.StageSubtree(root)

	// Resolve ID/IDREF cross-stage pointers and run semantic validation
	refErrs := stages.ResolveReferences()
	semErrs := stages.ValidateSemantics()
	allErrs := append(refErrs, semErrs...)

	return stages, root, allErrs, nil
}

// MarshalScorePartwise synchronizes ID/IDREF pointers, validates stage semantics,
// and serializes the multi-stage model back to indented XML.
func MarshalScorePartwise(stages *MusicXMLStages, root *score.ScorePartwise) ([]byte, error) {
	if stages != nil {
		stages.SyncReferences()
	}
	out, err := xml.MarshalIndent(root, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("xml marshal failed: %w", err)
	}
	return append([]byte(xml.Header), append(out, '\n')...), nil
}

// CompareXMLSemantics performs a canonical XML token comparison between two XML byte slices
// (ignoring whitespace-only formatting differences, comments, and attribute ordering) to verify
// lossless structural round-trip fidelity.
func CompareXMLStructure(expectedXML, actualXML []byte) error {
	expTokens, err := canonicalTokens(expectedXML)
	if err != nil {
		return fmt.Errorf("tokenizing expected XML: %w", err)
	}
	actTokens, err := canonicalTokens(actualXML)
	if err != nil {
		return fmt.Errorf("tokenizing actual XML: %w", err)
	}

	minLen := len(expTokens)
	if len(actTokens) < minLen {
		minLen = len(actTokens)
	}
	for i := 0; i < minLen; i++ {
		if expTokens[i] != actTokens[i] {
			return fmt.Errorf("token mismatch at #%d:\n  expected: %s\n  actual:   %s", i, expTokens[i], actTokens[i])
		}
	}
	if len(expTokens) != len(actTokens) {
		return fmt.Errorf("token count mismatch: expected %d, got %d", len(expTokens), len(actTokens))
	}
	return nil
}

func canonicalTokens(data []byte) ([]string, error) {
	dec := xml.NewDecoder(bytes.NewReader(data))
	var tokens []string
	var elemStack []string
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			var attrs []string
			for _, a := range t.Attr {
				if a.Name.Local == "xmlns" || a.Name.Space == "xmlns" {
					continue
				}
				attrs = append(attrs, fmt.Sprintf("%s=%q", a.Name.Local, strings.TrimSpace(a.Value)))
			}
			// Sort attributes alphabetically for canonical comparison
			for i := 0; i < len(attrs); i++ {
				for j := i + 1; j < len(attrs); j++ {
					if attrs[j] < attrs[i] {
						attrs[i], attrs[j] = attrs[j], attrs[i]
					}
				}
			}
			tokens = append(tokens, fmt.Sprintf("<%s %s>", t.Name.Local, strings.Join(attrs, " ")))
			elemStack = append(elemStack, t.Name.Local)
		case xml.EndElement:
			tokens = append(tokens, fmt.Sprintf("</%s>", t.Name.Local))
			if len(elemStack) > 0 {
				elemStack = elemStack[:len(elemStack)-1]
			}
		case xml.CharData:
			trimmed := strings.Join(strings.Fields(string(t)), " ")
			if trimmed != "" {
				tokens = append(tokens, fmt.Sprintf("TEXT(%s)", trimmed))
			}
		}
	}
	return tokens, nil
}
