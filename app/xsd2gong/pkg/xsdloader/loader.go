package xsdloader

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// SchemaSet holds all loaded XSD schemas and resolved global lookups.
type SchemaSet struct {
	RootSchema *Schema
	Schemas    []*Schema
	ByFilePath map[string]*Schema
}

// Schema represents a parsed XSD file along with module partitioning annotations.
type Schema struct {
	FilePath        string
	TargetNamespace string
	Xmlns           map[string]string

	Imports         []*Import
	Includes        []*Include
	SimpleTypes     []*SimpleType
	ComplexTypes    []*ComplexType
	Groups          []*Group
	AttributeGroups []*AttributeGroup
	Attributes      []*Attribute
	Elements        []*Element
}

type Import struct {
	Namespace      string
	SchemaLocation string
}

type Include struct {
	SchemaLocation string
}

type SimpleType struct {
	Name          string
	ModuleComment string
	Documentation string
	Order         int
	Restriction   *Restriction
	Union         *Union
	SourceSchema  *Schema
}

type Restriction struct {
	Base          string
	Enumerations  []string
	MinInclusive  *string
	MaxInclusive  *string
	MinExclusive  *string
	MaxExclusive  *string
	TotalDigits   *int
	MinLength     *int
	MaxLength     *int
	Length        *int
	Patterns      []string
	Attributes    []*Attribute
	AttributeGroups []*AttributeGroup
}

type Union struct {
	MemberTypes string
	SimpleTypes []*SimpleType
}

type ComplexType struct {
	Name            string
	ModuleComment   string
	Documentation   string
	Order           int
	Mixed           bool
	IsAnonymous     bool
	OuterElement    *Element
	Particle        *ModelGroup
	SimpleContent   *SimpleContent
	ComplexContent  *ComplexContent
	Attributes      []*Attribute
	AttributeGroups []*AttributeGroup
	SourceSchema    *Schema
}

type SimpleContent struct {
	Extension   *Extension
	Restriction *Restriction
}

type ComplexContent struct {
	Extension   *Extension
	Restriction *Restriction
}

type Extension struct {
	Base            string
	Particle        *ModelGroup
	Attributes      []*Attribute
	AttributeGroups []*AttributeGroup
}

type ModelGroup struct {
	Kind       string // "sequence", "choice", "all"
	MinOccurs  string
	MaxOccurs  string
	Order      int
	Elements   []*Element
	Groups     []*GroupRef
	SubGroups  []*ModelGroup
	Particles  []ParticleItem // Ordered list of children in document order
}

// ParticleItem represents an item inside a sequence, choice, or all in exact document order.
type ParticleItem struct {
	Element    *Element
	GroupRef   *GroupRef
	ModelGroup *ModelGroup
}

type Group struct {
	Name          string
	ModuleComment string
	Documentation string
	Order         int
	Particle      *ModelGroup
	SourceSchema  *Schema
}

type GroupRef struct {
	Ref       string
	MinOccurs string
	MaxOccurs string
	Order     int
}

type AttributeGroup struct {
	Name            string
	ModuleComment   string
	Documentation   string
	Order           int
	Ref             string
	Attributes      []*Attribute
	AttributeGroups []*AttributeGroup
	SourceSchema    *Schema
}

type Attribute struct {
	Name          string
	Ref           string
	Type          string
	Use           string // "required", "optional", "prohibited"
	Default       string
	Fixed         string
	Documentation string
	Order         int
	SimpleType    *SimpleType
	SourceSchema  *Schema
}

type Element struct {
	Name          string
	Ref           string
	Type          string
	MinOccurs     string
	MaxOccurs     string
	Default       string
	Fixed         string
	Abstract      string
	Documentation string
	ModuleComment string
	Order         int
	ComplexType   *ComplexType
	SimpleType    *SimpleType
	SourceSchema  *Schema
}

// Load parses a root XSD file and recursively loads all imported/included local XSD files.
func Load(rootXSDPath string) (*SchemaSet, error) {
	absRoot, err := filepath.Abs(rootXSDPath)
	if err != nil {
		return nil, err
	}

	set := &SchemaSet{
		ByFilePath: make(map[string]*Schema),
	}

	rootSchema, err := set.loadFile(absRoot)
	if err != nil {
		return nil, err
	}
	set.RootSchema = rootSchema
	return set, nil
}

func (set *SchemaSet) loadFile(absPath string) (*Schema, error) {
	if existing, ok := set.ByFilePath[absPath]; ok {
		return existing, nil
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", absPath, err)
	}

	schema, err := parseXSD(absPath, data)
	if err != nil {
		return nil, fmt.Errorf("parsing %s: %w", absPath, err)
	}

	set.ByFilePath[absPath] = schema
	set.Schemas = append(set.Schemas, schema)

	baseDir := filepath.Dir(absPath)

	for _, imp := range schema.Imports {
		if imp.SchemaLocation == "" || strings.HasPrefix(imp.SchemaLocation, "http://") || strings.HasPrefix(imp.SchemaLocation, "https://") {
			continue
		}
		impPath := filepath.Clean(filepath.Join(baseDir, imp.SchemaLocation))
		if _, err := set.loadFile(impPath); err != nil {
			return nil, err
		}
	}

	for _, inc := range schema.Includes {
		if inc.SchemaLocation == "" || strings.HasPrefix(inc.SchemaLocation, "http://") || strings.HasPrefix(inc.SchemaLocation, "https://") {
			continue
		}
		incPath := filepath.Clean(filepath.Join(baseDir, inc.SchemaLocation))
		if _, err := set.loadFile(incPath); err != nil {
			return nil, err
		}
	}

	return schema, nil
}

func parseXSD(filePath string, data []byte) (*Schema, error) {
	decoder := xml.NewDecoder(bytes.NewReader(data))
	schema := &Schema{
		FilePath: filePath,
		Xmlns:    make(map[string]string),
	}

	var currentModule string
	orderCounter := 1

	for {
		tok, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		switch t := tok.(type) {
		case xml.Comment:
			commentStr := strings.TrimSpace(string(t))
			if mod := extractModuleFromComment(commentStr); mod != "" {
				currentModule = mod
			}

		case xml.StartElement:
			local := t.Name.Local
			switch local {
			case "schema":
				for _, attr := range t.Attr {
					if attr.Name.Local == "targetNamespace" {
						schema.TargetNamespace = attr.Value
					} else if attr.Name.Space == "xmlns" {
						schema.Xmlns[attr.Name.Local] = attr.Value
					} else if attr.Name.Local == "xmlns" {
						schema.Xmlns[""] = attr.Value
					}
				}
			case "import":
				imp := &Import{}
				for _, attr := range t.Attr {
					switch attr.Name.Local {
					case "namespace":
						imp.Namespace = attr.Value
					case "schemaLocation":
						imp.SchemaLocation = attr.Value
					}
				}
				schema.Imports = append(schema.Imports, imp)
				_ = decoder.Skip()

			case "include":
				inc := &Include{}
				for _, attr := range t.Attr {
					if attr.Name.Local == "schemaLocation" {
						inc.SchemaLocation = attr.Value
					}
				}
				schema.Includes = append(schema.Includes, inc)
				_ = decoder.Skip()

			case "simpleType":
				st, err := parseSimpleType(decoder, t, schema, currentModule, &orderCounter)
				if err != nil {
					return nil, err
				}
				schema.SimpleTypes = append(schema.SimpleTypes, st)

			case "complexType":
				ct, err := parseComplexType(decoder, t, schema, currentModule, &orderCounter)
				if err != nil {
					return nil, err
				}
				schema.ComplexTypes = append(schema.ComplexTypes, ct)

			case "group":
				grp, err := parseGroup(decoder, t, schema, currentModule, &orderCounter)
				if err != nil {
					return nil, err
				}
				schema.Groups = append(schema.Groups, grp)

			case "attributeGroup":
				ag, err := parseAttributeGroup(decoder, t, schema, currentModule, &orderCounter)
				if err != nil {
					return nil, err
				}
				schema.AttributeGroups = append(schema.AttributeGroups, ag)

			case "attribute":
				attr, err := parseAttribute(decoder, t, schema, &orderCounter)
				if err != nil {
					return nil, err
				}
				schema.Attributes = append(schema.Attributes, attr)

			case "element":
				el, err := parseElement(decoder, t, schema, currentModule, &orderCounter)
				if err != nil {
					return nil, err
				}
				schema.Elements = append(schema.Elements, el)

			case "annotation":
				_ = decoder.Skip()
			}
		}
	}

	return schema, nil
}

// extractModuleFromComment detects module comments like:
// "Simple types derived from note.mod elements" -> "note"
// "<!-- Note module -->" -> "note"
func extractModuleFromComment(comment string) string {
	lower := strings.ToLower(comment)
	if idx := strings.Index(lower, ".mod"); idx > 0 {
		prefix := lower[:idx]
		fields := strings.Fields(prefix)
		if len(fields) > 0 {
			modName := fields[len(fields)-1]
			modName = strings.Trim(modName, " \"'<-")
			return modName
		}
	}
	if strings.HasSuffix(lower, "module") {
		fields := strings.Fields(lower)
		if len(fields) >= 2 {
			return strings.Trim(fields[0], " \"'<-")
		}
	}
	return ""
}

func getAttr(start xml.StartElement, name string) string {
	for _, a := range start.Attr {
		if a.Name.Local == name {
			return a.Value
		}
	}
	return ""
}

func parseAnnotation(decoder *xml.Decoder) (string, error) {
	var docs []string
	for {
		tok, err := decoder.Token()
		if err != nil {
			return "", err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			if t.Name.Local == "documentation" {
				var content string
				if err := decoder.DecodeElement(&content, &t); err != nil {
					// If documentation contains nested HTML (like in xml.xsd), skip or read raw
					_ = decoder.Skip()
				} else {
					docs = append(docs, strings.TrimSpace(content))
				}
			} else {
				_ = decoder.Skip()
			}
		case xml.EndElement:
			if t.Name.Local == "annotation" {
				return strings.Join(docs, "\n"), nil
			}
		}
	}
}

func parseSimpleType(decoder *xml.Decoder, start xml.StartElement, schema *Schema, module string, order *int) (*SimpleType, error) {
	st := &SimpleType{
		Name:          getAttr(start, "name"),
		ModuleComment: module,
		Order:         *order,
		SourceSchema:  schema,
	}
	*order++

	for {
		tok, err := decoder.Token()
		if err != nil {
			return nil, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "annotation":
				doc, err := parseAnnotation(decoder)
				if err != nil {
					return nil, err
				}
				st.Documentation = doc
			case "restriction":
				r, err := parseRestriction(decoder, t, schema, order)
				if err != nil {
					return nil, err
				}
				st.Restriction = r
			case "union":
				u, err := parseUnion(decoder, t, schema, module, order)
				if err != nil {
					return nil, err
				}
				st.Union = u
			default:
				_ = decoder.Skip()
			}
		case xml.EndElement:
			if t.Name.Local == "simpleType" {
				return st, nil
			}
		}
	}
}

func parseRestriction(decoder *xml.Decoder, start xml.StartElement, schema *Schema, order *int) (*Restriction, error) {
	r := &Restriction{
		Base: getAttr(start, "base"),
	}
	for {
		tok, err := decoder.Token()
		if err != nil {
			return nil, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			val := getAttr(t, "value")
			switch t.Name.Local {
			case "enumeration":
				r.Enumerations = append(r.Enumerations, val)
				_ = decoder.Skip()
			case "minInclusive":
				v := val
				r.MinInclusive = &v
				_ = decoder.Skip()
			case "maxInclusive":
				v := val
				r.MaxInclusive = &v
				_ = decoder.Skip()
			case "minExclusive":
				v := val
				r.MinExclusive = &v
				_ = decoder.Skip()
			case "maxExclusive":
				v := val
				r.MaxExclusive = &v
				_ = decoder.Skip()
			case "totalDigits":
				var d int
				fmt.Sscanf(val, "%d", &d)
				r.TotalDigits = &d
				_ = decoder.Skip()
			case "minLength":
				var d int
				fmt.Sscanf(val, "%d", &d)
				r.MinLength = &d
				_ = decoder.Skip()
			case "maxLength":
				var d int
				fmt.Sscanf(val, "%d", &d)
				r.MaxLength = &d
				_ = decoder.Skip()
			case "length":
				var d int
				fmt.Sscanf(val, "%d", &d)
				r.Length = &d
				_ = decoder.Skip()
			case "pattern":
				r.Patterns = append(r.Patterns, val)
				_ = decoder.Skip()
			case "attribute":
				attr, err := parseAttribute(decoder, t, schema, order)
				if err != nil {
					return nil, err
				}
				r.Attributes = append(r.Attributes, attr)
			case "attributeGroup":
				ag, err := parseAttributeGroup(decoder, t, schema, "", order)
				if err != nil {
					return nil, err
				}
				r.AttributeGroups = append(r.AttributeGroups, ag)
			default:
				_ = decoder.Skip()
			}
		case xml.EndElement:
			if t.Name.Local == "restriction" {
				return r, nil
			}
		}
	}
}

func parseUnion(decoder *xml.Decoder, start xml.StartElement, schema *Schema, module string, order *int) (*Union, error) {
	u := &Union{
		MemberTypes: getAttr(start, "memberTypes"),
	}
	for {
		tok, err := decoder.Token()
		if err != nil {
			return nil, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			if t.Name.Local == "simpleType" {
				st, err := parseSimpleType(decoder, t, schema, module, order)
				if err != nil {
					return nil, err
				}
				u.SimpleTypes = append(u.SimpleTypes, st)
			} else {
				_ = decoder.Skip()
			}
		case xml.EndElement:
			if t.Name.Local == "union" {
				return u, nil
			}
		}
	}
}

func parseComplexType(decoder *xml.Decoder, start xml.StartElement, schema *Schema, module string, order *int) (*ComplexType, error) {
	ct := &ComplexType{
		Name:          getAttr(start, "name"),
		Mixed:         getAttr(start, "mixed") == "true",
		ModuleComment: module,
		Order:         *order,
		SourceSchema:  schema,
	}
	*order++

	for {
		tok, err := decoder.Token()
		if err != nil {
			return nil, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "annotation":
				doc, err := parseAnnotation(decoder)
				if err != nil {
					return nil, err
				}
				ct.Documentation = doc
			case "sequence", "choice", "all":
				mg, err := parseModelGroup(decoder, t, schema, module, order)
				if err != nil {
					return nil, err
				}
				ct.Particle = mg
			case "group":
				// Direct group ref inside complexType (e.g. <xs:group ref="music-data"/>)
				gRef := &GroupRef{
					Ref:       getAttr(t, "ref"),
					MinOccurs: getAttr(t, "minOccurs"),
					MaxOccurs: getAttr(t, "maxOccurs"),
					Order:     *order,
				}
				*order++
				_ = decoder.Skip()
				if ct.Particle == nil {
					ct.Particle = &ModelGroup{
						Kind:  "sequence",
						Order: gRef.Order,
					}
				}
				ct.Particle.Groups = append(ct.Particle.Groups, gRef)
				ct.Particle.Particles = append(ct.Particle.Particles, ParticleItem{GroupRef: gRef})
			case "simpleContent":
				sc, err := parseSimpleContent(decoder, schema, order)
				if err != nil {
					return nil, err
				}
				ct.SimpleContent = sc
			case "complexContent":
				cc, err := parseComplexContent(decoder, schema, module, order)
				if err != nil {
					return nil, err
				}
				ct.ComplexContent = cc
			case "attribute":
				attr, err := parseAttribute(decoder, t, schema, order)
				if err != nil {
					return nil, err
				}
				ct.Attributes = append(ct.Attributes, attr)
			case "attributeGroup":
				ag, err := parseAttributeGroup(decoder, t, schema, module, order)
				if err != nil {
					return nil, err
				}
				ct.AttributeGroups = append(ct.AttributeGroups, ag)
			default:
				_ = decoder.Skip()
			}
		case xml.EndElement:
			if t.Name.Local == "complexType" {
				return ct, nil
			}
		}
	}
}

func parseSimpleContent(decoder *xml.Decoder, schema *Schema, order *int) (*SimpleContent, error) {
	sc := &SimpleContent{}
	for {
		tok, err := decoder.Token()
		if err != nil {
			return nil, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				ext, err := parseExtension(decoder, t, schema, "", order)
				if err != nil {
					return nil, err
				}
				sc.Extension = ext
			case "restriction":
				r, err := parseRestriction(decoder, t, schema, order)
				if err != nil {
					return nil, err
				}
				sc.Restriction = r
			default:
				_ = decoder.Skip()
			}
		case xml.EndElement:
			if t.Name.Local == "simpleContent" {
				return sc, nil
			}
		}
	}
}

func parseComplexContent(decoder *xml.Decoder, schema *Schema, module string, order *int) (*ComplexContent, error) {
	cc := &ComplexContent{}
	for {
		tok, err := decoder.Token()
		if err != nil {
			return nil, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				ext, err := parseExtension(decoder, t, schema, module, order)
				if err != nil {
					return nil, err
				}
				cc.Extension = ext
			case "restriction":
				r, err := parseRestriction(decoder, t, schema, order)
				if err != nil {
					return nil, err
				}
				cc.Restriction = r
			default:
				_ = decoder.Skip()
			}
		case xml.EndElement:
			if t.Name.Local == "complexContent" {
				return cc, nil
			}
		}
	}
}

func parseExtension(decoder *xml.Decoder, start xml.StartElement, schema *Schema, module string, order *int) (*Extension, error) {
	ext := &Extension{
		Base: getAttr(start, "base"),
	}
	for {
		tok, err := decoder.Token()
		if err != nil {
			return nil, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "sequence", "choice", "all":
				mg, err := parseModelGroup(decoder, t, schema, module, order)
				if err != nil {
					return nil, err
				}
				ext.Particle = mg
			case "group":
				gRef := &GroupRef{
					Ref:       getAttr(t, "ref"),
					MinOccurs: getAttr(t, "minOccurs"),
					MaxOccurs: getAttr(t, "maxOccurs"),
					Order:     *order,
				}
				*order++
				_ = decoder.Skip()
				if ext.Particle == nil {
					ext.Particle = &ModelGroup{Kind: "sequence", Order: gRef.Order}
				}
				ext.Particle.Groups = append(ext.Particle.Groups, gRef)
				ext.Particle.Particles = append(ext.Particle.Particles, ParticleItem{GroupRef: gRef})
			case "attribute":
				attr, err := parseAttribute(decoder, t, schema, order)
				if err != nil {
					return nil, err
				}
				ext.Attributes = append(ext.Attributes, attr)
			case "attributeGroup":
				ag, err := parseAttributeGroup(decoder, t, schema, module, order)
				if err != nil {
					return nil, err
				}
				ext.AttributeGroups = append(ext.AttributeGroups, ag)
			default:
				_ = decoder.Skip()
			}
		case xml.EndElement:
			if t.Name.Local == "extension" {
				return ext, nil
			}
		}
	}
}

func parseModelGroup(decoder *xml.Decoder, start xml.StartElement, schema *Schema, module string, order *int) (*ModelGroup, error) {
	mg := &ModelGroup{
		Kind:      start.Name.Local,
		MinOccurs: getAttr(start, "minOccurs"),
		MaxOccurs: getAttr(start, "maxOccurs"),
		Order:     *order,
	}
	*order++

	for {
		tok, err := decoder.Token()
		if err != nil {
			return nil, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "element":
				el, err := parseElement(decoder, t, schema, module, order)
				if err != nil {
					return nil, err
				}
				mg.Elements = append(mg.Elements, el)
				mg.Particles = append(mg.Particles, ParticleItem{Element: el})
			case "group":
				gRef := &GroupRef{
					Ref:       getAttr(t, "ref"),
					MinOccurs: getAttr(t, "minOccurs"),
					MaxOccurs: getAttr(t, "maxOccurs"),
					Order:     *order,
				}
				*order++
				_ = decoder.Skip()
				mg.Groups = append(mg.Groups, gRef)
				mg.Particles = append(mg.Particles, ParticleItem{GroupRef: gRef})
			case "sequence", "choice", "all":
				sub, err := parseModelGroup(decoder, t, schema, module, order)
				if err != nil {
					return nil, err
				}
				mg.SubGroups = append(mg.SubGroups, sub)
				mg.Particles = append(mg.Particles, ParticleItem{ModelGroup: sub})
			default:
				_ = decoder.Skip()
			}
		case xml.EndElement:
			if t.Name.Local == start.Name.Local {
				return mg, nil
			}
		}
	}
}

func parseGroup(decoder *xml.Decoder, start xml.StartElement, schema *Schema, module string, order *int) (*Group, error) {
	g := &Group{
		Name:          getAttr(start, "name"),
		ModuleComment: module,
		Order:         *order,
		SourceSchema:  schema,
	}
	*order++

	for {
		tok, err := decoder.Token()
		if err != nil {
			return nil, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "annotation":
				doc, err := parseAnnotation(decoder)
				if err != nil {
					return nil, err
				}
				g.Documentation = doc
			case "sequence", "choice", "all":
				mg, err := parseModelGroup(decoder, t, schema, module, order)
				if err != nil {
					return nil, err
				}
				g.Particle = mg
			default:
				_ = decoder.Skip()
			}
		case xml.EndElement:
			if t.Name.Local == "group" {
				return g, nil
			}
		}
	}
}

func parseAttributeGroup(decoder *xml.Decoder, start xml.StartElement, schema *Schema, module string, order *int) (*AttributeGroup, error) {
	ag := &AttributeGroup{
		Name:          getAttr(start, "name"),
		Ref:           getAttr(start, "ref"),
		ModuleComment: module,
		Order:         *order,
		SourceSchema:  schema,
	}
	*order++

	for {
		tok, err := decoder.Token()
		if err != nil {
			return nil, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "annotation":
				doc, err := parseAnnotation(decoder)
				if err != nil {
					return nil, err
				}
				ag.Documentation = doc
			case "attribute":
				attr, err := parseAttribute(decoder, t, schema, order)
				if err != nil {
					return nil, err
				}
				ag.Attributes = append(ag.Attributes, attr)
			case "attributeGroup":
				subAg, err := parseAttributeGroup(decoder, t, schema, module, order)
				if err != nil {
					return nil, err
				}
				ag.AttributeGroups = append(ag.AttributeGroups, subAg)
			default:
				_ = decoder.Skip()
			}
		case xml.EndElement:
			if t.Name.Local == "attributeGroup" {
				return ag, nil
			}
		}
	}
}

func parseAttribute(decoder *xml.Decoder, start xml.StartElement, schema *Schema, order *int) (*Attribute, error) {
	attr := &Attribute{
		Name:         getAttr(start, "name"),
		Ref:          getAttr(start, "ref"),
		Type:         getAttr(start, "type"),
		Use:          getAttr(start, "use"),
		Default:      getAttr(start, "default"),
		Fixed:        getAttr(start, "fixed"),
		Order:        *order,
		SourceSchema: schema,
	}
	*order++

	for {
		tok, err := decoder.Token()
		if err != nil {
			return nil, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "annotation":
				doc, err := parseAnnotation(decoder)
				if err != nil {
					return nil, err
				}
				attr.Documentation = doc
			case "simpleType":
				st, err := parseSimpleType(decoder, t, schema, "", order)
				if err != nil {
					return nil, err
				}
				attr.SimpleType = st
			default:
				_ = decoder.Skip()
			}
		case xml.EndElement:
			if t.Name.Local == "attribute" {
				return attr, nil
			}
		}
	}
}

func parseElement(decoder *xml.Decoder, start xml.StartElement, schema *Schema, module string, order *int) (*Element, error) {
	el := &Element{
		Name:          getAttr(start, "name"),
		Ref:           getAttr(start, "ref"),
		Type:          getAttr(start, "type"),
		MinOccurs:     getAttr(start, "minOccurs"),
		MaxOccurs:     getAttr(start, "maxOccurs"),
		Default:       getAttr(start, "default"),
		Fixed:         getAttr(start, "fixed"),
		Abstract:      getAttr(start, "abstract"),
		ModuleComment: module,
		Order:         *order,
		SourceSchema:  schema,
	}
	*order++

	for {
		tok, err := decoder.Token()
		if err != nil {
			return nil, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "annotation":
				doc, err := parseAnnotation(decoder)
				if err != nil {
					return nil, err
				}
				el.Documentation = doc
			case "complexType":
				ct, err := parseComplexType(decoder, t, schema, module, order)
				if err != nil {
					return nil, err
				}
				ct.IsAnonymous = true
				ct.OuterElement = el
				el.ComplexType = ct
			case "simpleType":
				st, err := parseSimpleType(decoder, t, schema, module, order)
				if err != nil {
					return nil, err
				}
				el.SimpleType = st
			default:
				_ = decoder.Skip()
			}
		case xml.EndElement:
			if t.Name.Local == "element" {
				return el, nil
			}
		}
	}
}

// LocalName strips any XML namespace prefix (e.g. "xs:string" -> "string", "xlink:href" -> "href").
func LocalName(qname string) string {
	if idx := strings.IndexByte(qname, ':'); idx >= 0 {
		return qname[idx+1:]
	}
	return qname
}

// Prefix returns the XML namespace prefix (e.g. "xlink:href" -> "xlink", "string" -> "").
func Prefix(qname string) string {
	if idx := strings.IndexByte(qname, ':'); idx >= 0 {
		return qname[:idx]
	}
	return ""
}

