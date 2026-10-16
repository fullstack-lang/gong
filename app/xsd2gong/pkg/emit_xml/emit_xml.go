package emit_xml

import (
	"fmt"
	"sort"
	"strings"

	"github.com/fullstack-lang/gong/app/xsd2gong/pkg/emit_models"
	"github.com/fullstack-lang/gong/app/xsd2gong/pkg/namer"
)

// RenderPackageXMLFile generates `models_xml.go` for a sub-package.
// It emits custom UnmarshalXML and MarshalXML for structs containing order-preserving
// choice wrappers (e.g. Measure with MeasureMusicDataChoice).
func RenderPackageXMLFile(ep *emit_models.EmittedPackage, em *emit_models.EmittedModel) string {
	var sb strings.Builder
	sb.WriteString("// generated code by xsd2gong - do not edit\n")
	sb.WriteString(fmt.Sprintf("package %s\n\n", ep.Name))

	// Check if any struct in ep has a choice wrapper slice
	var structsWithChoice []*emit_models.EmittedStruct
	for _, st := range ep.Structs {
		for _, f := range st.Fields {
			if targetSt := em.StructByKey["choice:"+f.TargetTypeName]; targetSt != nil && targetSt.IsChoiceWrap {
				structsWithChoice = append(structsWithChoice, st)
				break
			}
		}
	}

	if len(structsWithChoice) == 0 {
		return sb.String()
	}

	sb.WriteString("import (\n")
	sb.WriteString("\t\"encoding/xml\"\n")
	sb.WriteString("\t\"io\"\n")

	// Collect package imports needed by choice wrapper fields
	neededImports := make(map[string]string)
	for _, st := range structsWithChoice {
		for _, f := range st.Fields {
			if choiceSt := em.StructByKey["choice:"+f.TargetTypeName]; choiceSt != nil {
				for _, cf := range choiceSt.Fields {
					if cf.TargetPkg != "" && cf.TargetPkg != ep.Name {
						neededImports[cf.TargetPkg] = ep.Imports[cf.TargetPkg]
					}
				}
			}
		}
	}
	if len(neededImports) > 0 {
		sb.WriteString("\n")
		var impKeys []string
		for k := range neededImports {
			impKeys = append(impKeys, k)
		}
		sort.Strings(impKeys)
		for _, k := range impKeys {
			sb.WriteString(fmt.Sprintf("\t\"%s\"\n", neededImports[k]))
		}
	}
	sb.WriteString(")\n\n")

	for _, st := range structsWithChoice {
		var contentField *emit_models.EmittedField
		var choiceSt *emit_models.EmittedStruct
		for _, f := range st.Fields {
			if cs := em.StructByKey["choice:"+f.TargetTypeName]; cs != nil && cs.IsChoiceWrap {
				contentField = f
				choiceSt = cs
				break
			}
		}
		if contentField == nil || choiceSt == nil {
			continue
		}

		// Emit UnmarshalXML
		sb.WriteString(fmt.Sprintf("// UnmarshalXML streams attributes and child elements in exact document order for %s.\n", st.GoName))
		sb.WriteString(fmt.Sprintf("func (x *%s) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {\n", st.GoName))
		sb.WriteString("\tfor _, attr := range start.Attr {\n")
		sb.WriteString("\t\tswitch attr.Name.Local {\n")
		for _, f := range st.Fields {
			if !f.IsAttr {
				continue
			}
			sb.WriteString(fmt.Sprintf("\t\tcase %q:\n", f.XMLNameLocal))
			if f.FullGoType == "string" {
				sb.WriteString(fmt.Sprintf("\t\t\tx.%s = attr.Value\n", f.GoName))
			} else {
				sb.WriteString(fmt.Sprintf("\t\t\tx.%s = %s(attr.Value)\n", f.GoName, f.FullGoType))
			}
		}
		sb.WriteString("\t\t}\n")
		sb.WriteString("\t}\n\n")

		sb.WriteString("\tfor {\n")
		sb.WriteString("\t\ttok, err := d.Token()\n")
		sb.WriteString("\t\tif err == io.EOF {\n\t\t\tbreak\n\t\t}\n")
		sb.WriteString("\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n")
		sb.WriteString("\t\tswitch se := tok.(type) {\n")
		sb.WriteString("\t\tcase xml.StartElement:\n")
		sb.WriteString("\t\t\tswitch se.Name.Local {\n")
		for _, cf := range choiceSt.Fields {
			elemStructType := strings.TrimPrefix(cf.FullGoType, "*")
			sb.WriteString(fmt.Sprintf("\t\t\tcase %q:\n", cf.XMLNameLocal))
			sb.WriteString(fmt.Sprintf("\t\t\t\titem := new(%s)\n", elemStructType))
			sb.WriteString("\t\t\t\tif err := d.DecodeElement(item, &se); err != nil {\n\t\t\t\t\treturn err\n\t\t\t\t}\n")
			sb.WriteString(fmt.Sprintf("\t\t\t\tx.%s = append(x.%s, &%s{%s: item})\n", contentField.GoName, contentField.GoName, choiceSt.GoName, cf.GoName))
		}
		sb.WriteString("\t\t\tdefault:\n")
		sb.WriteString("\t\t\t\tif err := d.Skip(); err != nil {\n\t\t\t\t\treturn err\n\t\t\t\t}\n")
		sb.WriteString("\t\t\t}\n")
		sb.WriteString("\t\tcase xml.EndElement:\n")
		sb.WriteString("\t\t\tif se.Name.Local == start.Name.Local {\n\t\t\t\treturn nil\n\t\t\t}\n")
		sb.WriteString("\t\t}\n")
		sb.WriteString("\t}\n")
		sb.WriteString("\treturn nil\n")
		sb.WriteString("}\n\n")

		// Emit MarshalXML
		sb.WriteString(fmt.Sprintf("// MarshalXML serializes attributes and order-preserving choice items in exact document order for %s.\n", st.GoName))
		sb.WriteString(fmt.Sprintf("func (x *%s) MarshalXML(e *xml.Encoder, start xml.StartElement) error {\n", st.GoName))
		for _, f := range st.Fields {
			if !f.IsAttr {
				continue
			}
			sb.WriteString(fmt.Sprintf("\tif x.%s != \"\" {\n", f.GoName))
			sb.WriteString(fmt.Sprintf("\t\tstart.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: %q}, Value: string(x.%s)})\n", f.XMLNameLocal, f.GoName))
			sb.WriteString("\t}\n")
		}
		sb.WriteString("\tif err := e.EncodeToken(start); err != nil {\n\t\treturn err\n\t}\n")
		sb.WriteString(fmt.Sprintf("\tfor _, c := range x.%s {\n", contentField.GoName))
		sb.WriteString("\t\tif c == nil {\n\t\t\tcontinue\n\t\t}\n")
		for _, cf := range choiceSt.Fields {
			sb.WriteString(fmt.Sprintf("\t\tif c.%s != nil {\n", cf.GoName))
			sb.WriteString(fmt.Sprintf("\t\t\tif err := e.EncodeElement(c.%s, xml.StartElement{Name: xml.Name{Local: %q}}); err != nil {\n\t\t\t\treturn err\n\t\t\t}\n", cf.GoName, cf.XMLNameLocal))
			sb.WriteString("\t\t}\n")
		}
		sb.WriteString("\t}\n")
		sb.WriteString("\treturn e.EncodeToken(start.End())\n")
		sb.WriteString("}\n\n")
	}

	return sb.String()
}

// RenderOrchestratorStages generates `orchestrator/stages.go` which holds Stage pointers
// for all sub-packages and deterministic `StageSubtree` walkers.
func RenderOrchestratorStages(em *emit_models.EmittedModel) string {
	var sb strings.Builder
	sb.WriteString("// generated code by xsd2gong - do not edit\n")
	sb.WriteString("package orchestrator\n\n")

	sb.WriteString("import (\n")
	sb.WriteString("\t\"fmt\"\n\n")
	sb.WriteString(fmt.Sprintf("\tmodels %q\n", em.BaseImportPath))
	for _, ep := range em.OrderedPkgs {
		sb.WriteString(fmt.Sprintf("\t\"%s\"\n", ep.ImportPath))
	}
	sb.WriteString(")\n\n")

	// MusicXMLStages struct
	sb.WriteString("// MusicXMLStages coordinates all sub-package Gong stages.\n")
	sb.WriteString("type MusicXMLStages struct {\n")
	sb.WriteString("\tStageSet *models.StageSet\n")
	for _, ep := range em.OrderedPkgs {
		fieldTitle := namer.ToExportedIdentifier(ep.Name) + "Stage"
		sb.WriteString(fmt.Sprintf("\t%s *%s.Stage\n", fieldTitle, ep.Name))
	}
	sb.WriteString("}\n\n")

	// NewMusicXMLStages constructor
	rootName := em.Graph.RootPkg.Name
	sb.WriteString("// NewMusicXMLStages initializes all sub-package stages and binds them to a root StageSet.\n")
	sb.WriteString("func NewMusicXMLStages() *MusicXMLStages {\n")
	sb.WriteString("\tss := models.NewStageSet(\"musicxml\")\n")
	sb.WriteString("\ts := &MusicXMLStages{\n")
	sb.WriteString("\t\tStageSet: ss,\n")
	for _, ep := range em.OrderedPkgs {
		fieldTitle := namer.ToExportedIdentifier(ep.Name) + "Stage"
		sb.WriteString(fmt.Sprintf("\t\t%s: ss.%s,\n", fieldTitle, fieldTitle))
	}
	sb.WriteString("\t}\n")
	sb.WriteString("\treturn s\n")
	sb.WriteString("}\n\n")

	// Commitall in topological order
	sb.WriteString("// Commit commits all stages in topological dependency order (leaves first, root last).\n")
	sb.WriteString("func (s *MusicXMLStages) Commit() {\n")
	sb.WriteString("\tif s.StageSet != nil {\n")
	sb.WriteString("\t\ts.StageSet.Commit()\n")
	sb.WriteString("\t\treturn\n")
	sb.WriteString("\t}\n")
	for _, ep := range em.OrderedPkgs {
		fieldTitle := namer.ToExportedIdentifier(ep.Name) + "Stage"
		sb.WriteString(fmt.Sprintf("\tif s.%s != nil {\n\t\ts.%s.Commit()\n\t}\n", fieldTitle, fieldTitle))
	}
	sb.WriteString("}\n\n")

	// StageSubtree walkers for every Gongstruct across all packages
	sb.WriteString("// StageSubtree walks the object graph starting from root, assigns deterministic Gong Names,\n")
	sb.WriteString("// and stages each instance into its respective package Stage.\n")
	sb.WriteString(fmt.Sprintf("func (s *MusicXMLStages) StageSubtree(root *%s.ScorePartwise) {\n", rootName))
	sb.WriteString("\tcounts := make(map[string]int)\n")
	sb.WriteString("\ts.stageScore_ScorePartwise(root, counts)\n")
	sb.WriteString("}\n\n")

	for _, ep := range em.OrderedPkgs {
		stageField := namer.ToExportedIdentifier(ep.Name) + "Stage"
		for _, st := range ep.Structs {
			fnName := fmt.Sprintf("stage%s_%s", namer.ToExportedIdentifier(ep.Name), st.GoName)
			qualType := fmt.Sprintf("%s.%s", ep.Name, st.GoName)
			sb.WriteString(fmt.Sprintf("func (s *MusicXMLStages) %s(inst *%s, counts map[string]int) {\n", fnName, qualType))
			sb.WriteString("\tif inst == nil {\n\t\treturn\n\t}\n")
			sb.WriteString("\tif inst.Name == \"\" {\n")
			// Check if struct has an Id or Number attribute to make names even more informative
			var idField string
			for _, f := range st.Fields {
				if f.IsAttr && (f.GoName == "Id" || f.GoName == "Number") && f.FullGoType == "string" {
					idField = f.GoName
					break
				}
			}
			sb.WriteString(fmt.Sprintf("\t\tcounts[%q]++\n", qualType))
			if idField != "" {
				sb.WriteString(fmt.Sprintf("\t\tif inst.%s != \"\" {\n", idField))
				sb.WriteString(fmt.Sprintf("\t\t\tinst.Name = fmt.Sprintf(\"%s_%%s_%%d\", inst.%s, counts[%q])\n", st.GoName, idField, qualType))
				sb.WriteString("\t\t} else {\n")
				sb.WriteString(fmt.Sprintf("\t\t\tinst.Name = fmt.Sprintf(\"%s_%%d\", counts[%q])\n", st.GoName, qualType))
				sb.WriteString("\t\t}\n")
			} else {
				sb.WriteString(fmt.Sprintf("\t\tinst.Name = fmt.Sprintf(\"%s_%%d\", counts[%q])\n", st.GoName, qualType))
			}
			sb.WriteString("\t}\n")
			sb.WriteString(fmt.Sprintf("\tinst.Stage(s.%s)\n", stageField))

			for _, f := range st.Fields {
				if strings.HasPrefix(f.GoName, "Ref") && f.XMLTag == "`xml:\"-\"`" {
					continue
				}
				targetPkg := f.TargetPkg
				if targetPkg == "" {
					continue
				}
				// Check if target is a struct in TargetPkg
				if !f.IsPointer && !f.IsSlice {
					continue
				}
				childFn := fmt.Sprintf("stage%s_%s", namer.ToExportedIdentifier(targetPkg), f.TargetTypeName)
				if f.IsPointer {
					sb.WriteString(fmt.Sprintf("\ts.%s(inst.%s, counts)\n", childFn, f.GoName))
				} else if f.IsSlice {
					sb.WriteString(fmt.Sprintf("\tfor _, child := range inst.%s {\n", f.GoName))
					sb.WriteString(fmt.Sprintf("\t\ts.%s(child, counts)\n", childFn))
					sb.WriteString("\t}\n")
				}
			}
			sb.WriteString("}\n\n")
		}
	}

	return sb.String()
}

// RenderOrchestratorXMLIO generates `orchestrator/xml_io.go` with lossless round-trip XML unmarshalling and marshalling.
func RenderOrchestratorXMLIO(em *emit_models.EmittedModel) string {
	rootPkg := em.Graph.RootPkg.Name
	var sb strings.Builder
	sb.WriteString(`// generated code by xsd2gong - do not edit
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

	"` + em.Packages[rootPkg].ImportPath + `"
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
func UnmarshalScorePartwiseFile(path string) (*MusicXMLStages, *` + rootPkg + `.ScorePartwise, []error, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, nil, err
	}
	return UnmarshalScorePartwise(data)
}

// UnmarshalScorePartwise parses XML bytes into a multi-stage Gong model hierarchy,
// stages all instances across their respective package stages, resolves ID/IDREF pointers,
// and executes semantic validation across all stages.
func UnmarshalScorePartwise(xmlData []byte) (*MusicXMLStages, *` + rootPkg + `.ScorePartwise, []error, error) {
	xmlData = normalizeXMLBytes(xmlData)
	stages := NewMusicXMLStages()
	root := new(` + rootPkg + `.ScorePartwise)
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
func MarshalScorePartwise(stages *MusicXMLStages, root *` + rootPkg + `.ScorePartwise) ([]byte, error) {
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
`)
	return sb.String()
}

