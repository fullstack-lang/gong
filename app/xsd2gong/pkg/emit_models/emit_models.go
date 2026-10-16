package emit_models

import (
	"fmt"
	"path"
	"sort"
	"strings"

	"github.com/fullstack-lang/gong/app/xsd2gong/pkg/depgraph"
	"github.com/fullstack-lang/gong/app/xsd2gong/pkg/namer"
	"github.com/fullstack-lang/gong/app/xsd2gong/pkg/xsdloader"
)

// EmittedModel holds the resolved AST/IR of all generated packages, structs, enums, and fields.
type EmittedModel struct {
	BaseImportPath string // e.g. "github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models"
	Graph          *depgraph.Graph
	Packages       map[string]*EmittedPackage
	OrderedPkgs    []*EmittedPackage // Leaves first, Root last

	// Global lookup from XSD complexType/element key to EmittedStruct
	StructByKey map[string]*EmittedStruct
	EnumByKey   map[string]*EmittedEnum
}

func (em *EmittedModel) RootPkgHasStageSet(_ *EmittedModel) bool {
	return em.Graph != nil && em.Graph.RootPkg != nil
}

type EmittedPackage struct {
	Name       string
	ImportPath string
	DirRel     string // relative to go/models, e.g. "xlink", "common", ".", etc.
	IsRoot     bool
	Enums      []*EmittedEnum
	Structs    []*EmittedStruct
	Imports    map[string]string // pkgName -> full import path
}

type EmittedEnum struct {
	GoName    string
	XSDName   string
	PkgName   string
	BaseGo    string // "string" or "int"
	Values    []EnumVal
	SourceDoc string
}

type EnumVal struct {
	ConstName string
	Literal   string
}

type EmittedStruct struct {
	GoName        string
	XSDName       string
	PkgName       string
	XMLNameTag    string // Non-empty for root elements (e.g. "score-partwise")
	IsChoiceWrap  bool
	IsSeqWrap     bool
	Fields        []*EmittedField
	SourceDoc     string
	SourceCT      *xsdloader.ComplexType
	ChoiceGroups  [][]string // XOR choice sets (slices of field GoNames where minOccurs=1, maxOccurs=1)
	OptionalXORs  [][]string // Optional XOR choice sets (at most 1 of field GoNames can be set)
}

type EmittedField struct {
	GoName           string
	XMLTag           string // e.g. `xml:"id,attr,omitempty"` or `xml:"note,omitempty"` or `xml:",chardata"`
	XMLNameLocal     string // e.g. "id", "note"
	XMLSpace         string // e.g. "http://www.w3.org/1999/xlink" or "http://www.w3.org/XML/1998/namespace"
	IsAttr           bool
	IsCharData       bool
	IsPointer        bool
	IsSlice          bool
	IsValueSlice     bool   // e.g. []string for repeated simple elements
	IsEmptyElem      bool   // true if element has type="empty" (presence flag or empty struct pointer)
	TargetPkg        string // "" if basic or same pkg, otherwise target package name
	TargetTypeName   string // Go type name without package qualifier (e.g. "Note", "String", "Enum_YesNo")
	FullGoType       string // Qualified Go type in this struct's package (e.g. "*note.Note", "string")
	Required         bool
	DefaultVal       string
	FixedVal         string
	SimpleType       *xsdloader.SimpleType // Underlying XSD simpleType for semantic facet validation
	IsID             bool
	IsIDREF          bool
	IDREFTargetPkg   string
	IDREFTargetType  string // e.g. "ScorePart", "ScoreInstrument", "Player"
	RefPointerGoName string // e.g. "RefScorePart"
}

// Prepare builds the EmittedModel from the dependency graph and XSD schema set.
func Prepare(graph *depgraph.Graph, baseImportPath string, rootInModelsDir bool) (*EmittedModel, error) {
	em := &EmittedModel{
		BaseImportPath: baseImportPath,
		Graph:          graph,
		Packages:       make(map[string]*EmittedPackage),
		StructByKey:    make(map[string]*EmittedStruct),
		EnumByKey:      make(map[string]*EmittedEnum),
	}

	// 1. Initialize EmittedPackages in topological order
	for _, node := range graph.TopologicalOrd {
		dirRel := node.Name
		impPath := path.Join(baseImportPath, node.Name)
		if node.IsRoot && rootInModelsDir {
			dirRel = "."
			impPath = baseImportPath
		}
		ep := &EmittedPackage{
			Name:       node.Name,
			ImportPath: impPath,
			DirRel:     dirRel,
			IsRoot:     node.IsRoot,
			Imports:    make(map[string]string),
		}
		em.Packages[node.Name] = ep
		em.OrderedPkgs = append(em.OrderedPkgs, ep)
	}

	// 2. Register all Enumerations and Struct names per package first
	pkgScopes := make(map[string]*namer.Scope)
	for _, ep := range em.OrderedPkgs {
		pkgScopes[ep.Name] = namer.NewScope()
	}

	for _, node := range graph.TopologicalOrd {
		ep := em.Packages[node.Name]
		scope := pkgScopes[node.Name]

		for _, st := range node.SimpleTypes {
			if isEnumSimpleType(st) {
				goName := scope.Unique("Enum_" + namer.ToExportedIdentifier(st.Name))
				enum := &EmittedEnum{
					GoName:    goName,
					XSDName:   st.Name,
					PkgName:   node.Name,
					BaseGo:    "string",
					SourceDoc: st.Documentation,
				}
				valScope := namer.NewScope()
				for _, v := range st.Restriction.Enumerations {
					cName := valScope.Unique(namer.ToEnumConstantName(goName, v))
					enum.Values = append(enum.Values, EnumVal{
						ConstName: cName,
						Literal:   v,
					})
				}
				ep.Enums = append(ep.Enums, enum)
				key := qualifyTypeKey("st", st.SourceSchema, st.Name)
				em.EnumByKey[key] = enum
				if st.SourceSchema == graph.RootPkg.SimpleTypes[0].SourceSchema || em.EnumByKey["st:"+st.Name] == nil {
					em.EnumByKey["st:"+st.Name] = enum
				}
			}
		}

		for _, ct := range node.ComplexTypes {
			goName := scope.Unique(namer.ToExportedIdentifier(ct.Name))
			st := &EmittedStruct{
				GoName:    goName,
				XSDName:   ct.Name,
				PkgName:   node.Name,
				SourceDoc: ct.Documentation,
				SourceCT:  ct,
			}
			ep.Structs = append(ep.Structs, st)
			key := qualifyTypeKey("ct", ct.SourceSchema, ct.Name)
			em.StructByKey[key] = st
			if em.StructByKey["ct:"+ct.Name] == nil {
				em.StructByKey["ct:"+ct.Name] = st
			}
		}

		for _, el := range node.Elements {
			if el.ComplexType != nil {
				goName := scope.Unique(namer.ToExportedIdentifier(el.Name))
				st := &EmittedStruct{
					GoName:     goName,
					XSDName:    el.Name,
					PkgName:    node.Name,
					XMLNameTag: el.Name,
					SourceDoc:  el.Documentation,
					SourceCT:   el.ComplexType,
				}
				ep.Structs = append(ep.Structs, st)
				em.StructByKey["el:"+el.Name] = st
			}
		}
	}

	// 3. Populate fields for each struct (and create anonymous / choice wrapper structs as needed)
	for _, node := range graph.TopologicalOrd {
		ep := em.Packages[node.Name]
		scope := pkgScopes[node.Name]
		// Copy initial slice because new anonymous/choice structs may be appended to ep.Structs
		initialStructs := append([]*EmittedStruct(nil), ep.Structs...)
		for _, st := range initialStructs {
			if st.SourceCT != nil {
				em.populateStructFields(ep, scope, st, st.SourceCT)
			}
		}
	}

	return em, nil
}

func qualifyTypeKey(kind string, sch *xsdloader.Schema, name string) string {
	if sch == nil {
		return kind + ":" + name
	}
	base := strings.ToLower(path.Base(strings.ReplaceAll(sch.FilePath, "\\", "/")))
	base = strings.TrimSuffix(base, ".xsd")
	if base == "musicxml" {
		return kind + ":" + name
	}
	return kind + ":" + base + ":" + name
}

func isEnumSimpleType(st *xsdloader.SimpleType) bool {
	if st == nil || st.Restriction == nil || len(st.Restriction.Enumerations) == 0 {
		return false
	}
	return true
}

func (em *EmittedModel) populateStructFields(ep *EmittedPackage, pkgScope *namer.Scope, st *EmittedStruct, ct *xsdloader.ComplexType) {
	fieldScope := namer.NewScope()
	fieldScope.Reserve("Name")
	if st.XMLNameTag != "" {
		fieldScope.Reserve("XMLName")
	}

	// 1. Collect all attributes (including from base extension and attributeGroups)
	var attrs []*xsdloader.Attribute
	visitedAG := make(map[string]bool)

	var collectAG func(ag *xsdloader.AttributeGroup)
	collectAG = func(ag *xsdloader.AttributeGroup) {
		if ag.Ref != "" {
			prefix := xsdloader.Prefix(ag.Ref)
			local := xsdloader.LocalName(ag.Ref)
			key := local
			if prefix == "xlink" {
				key = "ag:xlink:" + local
			} else if prefix == "xml" {
				key = "ag:xml:" + local
			}
			if visitedAG[key] {
				return
			}
			visitedAG[key] = true
			if target, ok := em.Graph.AttributeGroupByName[key]; ok {
				for _, a := range target.Attributes {
					attrs = append(attrs, a)
				}
				for _, sub := range target.AttributeGroups {
					collectAG(sub)
				}
			} else if target, ok := em.Graph.AttributeGroupByName[local]; ok {
				for _, a := range target.Attributes {
					attrs = append(attrs, a)
				}
				for _, sub := range target.AttributeGroups {
					collectAG(sub)
				}
			}
			return
		}
		for _, a := range ag.Attributes {
			attrs = append(attrs, a)
		}
		for _, sub := range ag.AttributeGroups {
			collectAG(sub)
		}
	}

	// If complexContent extension, inherit base complexType attributes and elements first
	if ct.ComplexContent != nil && ct.ComplexContent.Extension != nil {
		ext := ct.ComplexContent.Extension
		baseLocal := xsdloader.LocalName(ext.Base)
		if baseCT, ok := em.Graph.ComplexTypeByName[baseLocal]; ok {
			for _, a := range baseCT.Attributes {
				attrs = append(attrs, a)
			}
			for _, ag := range baseCT.AttributeGroups {
				collectAG(ag)
			}
		}
		for _, a := range ext.Attributes {
			attrs = append(attrs, a)
		}
		for _, ag := range ext.AttributeGroups {
			collectAG(ag)
		}
	}

	if ct.SimpleContent != nil && ct.SimpleContent.Extension != nil {
		ext := ct.SimpleContent.Extension
		// Base could be another complexType with simpleContent or a simpleType
		baseLocal := xsdloader.LocalName(ext.Base)
		if baseCT, ok := em.Graph.ComplexTypeByName[baseLocal]; ok {
			for _, a := range baseCT.Attributes {
				attrs = append(attrs, a)
			}
			for _, ag := range baseCT.AttributeGroups {
				collectAG(ag)
			}
		}
		for _, a := range ext.Attributes {
			attrs = append(attrs, a)
		}
		for _, ag := range ext.AttributeGroups {
			collectAG(ag)
		}
	}

	for _, a := range ct.Attributes {
		attrs = append(attrs, a)
	}
	for _, ag := range ct.AttributeGroups {
		collectAG(ag)
	}

	// Deduplicate attributes by XML name
	seenAttrXML := make(map[string]bool)
	for _, a := range attrs {
		field := em.buildAttributeField(ep, fieldScope, st, a)
		if field == nil {
			continue
		}
		xmlKey := field.XMLSpace + ":" + field.XMLNameLocal
		if seenAttrXML[xmlKey] {
			continue
		}
		seenAttrXML[xmlKey] = true
		st.Fields = append(st.Fields, field)
	}

	// 2. Handle SimpleContent or Mixed content CharData
	if ct.SimpleContent != nil && ct.SimpleContent.Extension != nil {
		base := ct.SimpleContent.Extension.Base
		goType, targetPkg, stRef := em.resolveSimpleGoType(ep, base)
		if targetPkg != "" && targetPkg != ep.Name {
			em.recordImport(ep, targetPkg)
		}
		fName := fieldScope.Unique("Value")
		st.Fields = append(st.Fields, &EmittedField{
			GoName:         fName,
			XMLTag:         "`xml:\",chardata\"`",
			IsCharData:     true,
			TargetPkg:      targetPkg,
			TargetTypeName: goType,
			FullGoType:     qualifyGoType(ep.Name, targetPkg, goType, false, false),
			SimpleType:     stRef,
		})
	} else if ct.Mixed {
		fName := fieldScope.Unique("Value")
		st.Fields = append(st.Fields, &EmittedField{
			GoName:         fName,
			XMLTag:         "`xml:\",chardata\"`",
			IsCharData:     true,
			TargetTypeName: "string",
			FullGoType:     "string",
		})
	}

	// 3. Collect child elements / particles
	if ct.ComplexContent != nil && ct.ComplexContent.Extension != nil {
		baseLocal := xsdloader.LocalName(ct.ComplexContent.Extension.Base)
		if baseCT, ok := em.Graph.ComplexTypeByName[baseLocal]; ok && baseCT.Particle != nil {
			em.populateModelGroupFields(ep, pkgScope, fieldScope, st, baseCT.Particle, false)
		}
		if ct.ComplexContent.Extension.Particle != nil {
			em.populateModelGroupFields(ep, pkgScope, fieldScope, st, ct.ComplexContent.Extension.Particle, false)
		}
	}

	if ct.Particle != nil {
		em.populateModelGroupFields(ep, pkgScope, fieldScope, st, ct.Particle, false)
	}
}

func (em *EmittedModel) buildAttributeField(ep *EmittedPackage, fieldScope *namer.Scope, st *EmittedStruct, a *xsdloader.Attribute) *EmittedField {
	rawName := a.Name
	attrType := a.Type
	attrUse := a.Use
	xmlSpace := ""

	if a.Ref != "" {
		prefix := xsdloader.Prefix(a.Ref)
		local := xsdloader.LocalName(a.Ref)
		rawName = local
		if prefix == "xlink" {
			xmlSpace = "http://www.w3.org/1999/xlink"
			if target, ok := em.Graph.AttributeByName["attr:xlink:"+local]; ok {
				attrType = target.Type
				if attrUse == "" {
					attrUse = target.Use
				}
			}
		} else if prefix == "xml" {
			xmlSpace = "http://www.w3.org/XML/1998/namespace"
			attrType = "xs:string"
		} else if target, ok := em.Graph.AttributeByName[local]; ok {
			attrType = target.Type
			if attrUse == "" {
				attrUse = target.Use
			}
		}
	}

	if rawName == "" {
		return nil
	}

	goName := fieldScope.Unique(rawName)
	goType, targetPkg, stRef := em.resolveSimpleGoType(ep, attrType)
	if a.SimpleType != nil {
		stRef = a.SimpleType
	}
	if targetPkg != "" && targetPkg != ep.Name {
		em.recordImport(ep, targetPkg)
	}

	xmlTagName := rawName
	if xmlSpace != "" {
		xmlTagName = xmlSpace + " " + rawName
	}

	isID := (xsdloader.LocalName(attrType) == "ID")
	isIDREF := (xsdloader.LocalName(attrType) == "IDREF")

	field := &EmittedField{
		GoName:         goName,
		XMLTag:         fmt.Sprintf("`xml:\"%s,attr,omitempty\"`", xmlTagName),
		XMLNameLocal:   rawName,
		XMLSpace:       xmlSpace,
		IsAttr:         true,
		TargetPkg:      targetPkg,
		TargetTypeName: goType,
		FullGoType:     qualifyGoType(ep.Name, targetPkg, goType, false, false),
		Required:       attrUse == "required",
		DefaultVal:     a.Default,
		FixedVal:       a.Fixed,
		SimpleType:     stRef,
		IsID:           isID,
		IsIDREF:        isIDREF,
	}

	if isIDREF {
		em.configureIDREFTarget(ep, fieldScope, st, field)
	}

	return field
}

func (em *EmittedModel) configureIDREFTarget(ep *EmittedPackage, fieldScope *namer.Scope, st *EmittedStruct, field *EmittedField) {
	// Determine the target Gong struct for this IDREF attribute
	var targetKey string
	if field.XMLNameLocal == "player" {
		targetKey = "ct:player"
	} else if st.XSDName == "part" || strings.HasSuffix(st.GoName, "Part") {
		targetKey = "ct:score-part"
	} else {
		// In MusicXML, midi-device, midi-instrument, instrument, instrument-link, instrument-change, play
		// all reference score-instrument by id!
		targetKey = "ct:score-instrument"
	}

	targetStruct := em.StructByKey[targetKey]
	if targetStruct == nil {
		return
	}

	// Only add a direct struct pointer if doing so does not introduce an upward package import cycle!
	// If targetStruct is in a higher-ranked package than ep.Name, the cross-stage orchestrator links it,
	// or if they are in the same or lower-ranked package, we add the direct pointer field on the struct.
	if canImportPkg(ep.Name, targetStruct.PkgName) {
		if targetStruct.PkgName != ep.Name {
			em.recordImport(ep, targetStruct.PkgName)
		}
		refFieldName := fieldScope.Unique("Ref" + targetStruct.GoName)
		field.IDREFTargetPkg = targetStruct.PkgName
		field.IDREFTargetType = targetStruct.GoName
		field.RefPointerGoName = refFieldName

		refFullType := qualifyGoType(ep.Name, targetStruct.PkgName, targetStruct.GoName, true, false)
		st.Fields = append(st.Fields, &EmittedField{
			GoName:         refFieldName,
			XMLTag:         "`xml:\"-\"`",
			IsPointer:      true,
			TargetPkg:      targetStruct.PkgName,
			TargetTypeName: targetStruct.GoName,
			FullGoType:     refFullType,
		})
	}
}

func canImportPkg(fromPkg, toPkg string) bool {
	if fromPkg == toPkg {
		return true
	}
	ranks := map[string]int{
		"xlink":      0,
		"common":     1,
		"barline":    2,
		"attributes": 3,
		"direction":  4,
		"note":       5,
		"score":      6,
	}
	rFrom, ok1 := ranks[fromPkg]
	rTo, ok2 := ranks[toPkg]
	if ok1 && ok2 {
		return rFrom > rTo
	}
	return false
}

func (em *EmittedModel) populateModelGroupFields(
	ep *EmittedPackage,
	pkgScope *namer.Scope,
	fieldScope *namer.Scope,
	st *EmittedStruct,
	mg *xsdloader.ModelGroup,
	forceOptionalOrRepeated bool,
) {
	if mg == nil {
		return
	}

	isUnboundedGroup := (mg.MaxOccurs == "unbounded" || (mg.MaxOccurs != "" && mg.MaxOccurs != "1" && mg.MaxOccurs != "0"))
	isChoice := (mg.Kind == "choice")

	// Check if this is an unbounded choice with multiple distinct child elements (e.g., music-data in Measure)
	// For such a construct, generating an order-preserving Choice wrapper struct preserves exact timeline order!
	if st.GoName == "Measure" {
		if musicDataGrp, ok := em.Graph.GroupByName["music-data"]; ok && musicDataGrp.Particle != nil {
			for _, sub := range musicDataGrp.Particle.SubGroups {
				if sub.Kind == "choice" {
					em.emitOrderPreservingChoiceField(ep, pkgScope, fieldScope, st, sub, "MusicDataChoice")
					return
				}
			}
		}
	}
	if isChoice && isUnboundedGroup && st.XSDName == "measure" {
		em.emitOrderPreservingChoiceField(ep, pkgScope, fieldScope, st, mg, "MusicDataChoice")
		return
	}

	// Track direct elements in a single (bounded) choice for XOR semantic validation
	var xorGroup []string
	isSingleChoice := isChoice && !isUnboundedGroup && !forceOptionalOrRepeated

	// Flatten all elements reachable inside this ModelGroup in document order
	elements, elemRepeated, elemOptional := em.flattenParticles(mg, isUnboundedGroup, isChoice || forceOptionalOrRepeated || mg.MinOccurs == "0")

	seenElem := make(map[string]*EmittedField)
	for i, el := range elements {
		elName := el.Name
		if el.Ref != "" {
			elName = xsdloader.LocalName(el.Ref)
		}
		if elName == "" {
			continue
		}

		repeated := elemRepeated[i]
		optional := elemOptional[i]

		// If the same element name appears multiple times in different branches of a sequence/choice
		// (like cue, tie, or metronome-note in MusicXML), upgrade its field to a slice or optional!
		if existing, ok := seenElem[elName]; ok {
			// If it can appear in sequential branches (e.g. part-group or score-part in part-list, or metronome-note),
			// ensure it is a slice if it can occur > 1 time overall.
			if !existing.IsSlice && !existing.IsValueSlice {
				if elName == "score-part" || elName == "part-group" || elName == "credit-words" || elName == "credit-symbol" || elName == "link" || elName == "bookmark" || elName == "beat-unit" || elName == "beat-unit-dot" || elName == "beat-unit-tied" || elName == "metronome-note" {
					em.upgradeFieldToSlice(ep, existing)
				}
			}
			continue
		}

		f := em.buildElementField(ep, pkgScope, fieldScope, st, el, repeated, optional)
		if f != nil {
			seenElem[elName] = f
			st.Fields = append(st.Fields, f)
			if isSingleChoice && len(mg.SubGroups) == 0 && len(mg.Groups) == 0 {
				xorGroup = append(xorGroup, f.GoName)
			}
		}
	}

	if len(xorGroup) > 1 {
		if mg.MinOccurs == "0" {
			st.OptionalXORs = append(st.OptionalXORs, xorGroup)
		} else {
			st.ChoiceGroups = append(st.ChoiceGroups, xorGroup)
		}
	}

	// Special check for full-note group inside Note (Pitch, Unpitched, Rest XOR)
	if st.XSDName == "note" {
		hasPitch, hasUnpitched, hasRest := false, false, false
		for _, f := range st.Fields {
			switch f.XMLNameLocal {
			case "pitch":
				hasPitch = true
			case "unpitched":
				hasUnpitched = true
			case "rest":
				hasRest = true
			}
		}
		if hasPitch && hasUnpitched && hasRest {
			st.ChoiceGroups = append(st.ChoiceGroups, []string{"Pitch", "Unpitched", "Rest"})
		}
	}
}

func (em *EmittedModel) upgradeFieldToSlice(ep *EmittedPackage, f *EmittedField) {
	if f.IsPointer {
		f.IsPointer = false
		f.IsSlice = true
		f.FullGoType = qualifyGoType(ep.Name, f.TargetPkg, f.TargetTypeName, false, true)
	} else if !f.IsSlice && !f.IsValueSlice {
		f.IsValueSlice = true
		f.FullGoType = "[]" + f.FullGoType
	}
}

func (em *EmittedModel) emitOrderPreservingChoiceField(
	ep *EmittedPackage,
	pkgScope *namer.Scope,
	fieldScope *namer.Scope,
	parentStruct *EmittedStruct,
	mg *xsdloader.ModelGroup,
	suffix string,
) {
	wrapName := pkgScope.Unique(parentStruct.GoName + suffix)
	wrapStruct := &EmittedStruct{
		GoName:       wrapName,
		XSDName:      parentStruct.XSDName + "-choice",
		PkgName:      ep.Name,
		IsChoiceWrap: true,
		SourceDoc:    fmt.Sprintf("Order-preserving xs:choice wrapper for %s", parentStruct.GoName),
	}

	wrapFieldScope := namer.NewScope()
	wrapFieldScope.Reserve("Name")

	var choiceFields []string
	for _, el := range mg.Elements {
		f := em.buildElementField(ep, pkgScope, wrapFieldScope, wrapStruct, el, false, true)
		if f != nil {
			wrapStruct.Fields = append(wrapStruct.Fields, f)
			choiceFields = append(choiceFields, f.GoName)
		}
	}
	if len(choiceFields) > 1 {
		wrapStruct.ChoiceGroups = append(wrapStruct.ChoiceGroups, choiceFields)
	}

	ep.Structs = append(ep.Structs, wrapStruct)
	em.StructByKey["choice:"+wrapName] = wrapStruct

	contentFieldName := fieldScope.Unique("MusicData")
	parentStruct.Fields = append(parentStruct.Fields, &EmittedField{
		GoName:         contentFieldName,
		XMLTag:         "`xml:\"-\"`", // Handled by custom order-preserving UnmarshalXML / MarshalXML in models_xml.go
		IsSlice:        true,
		TargetPkg:      ep.Name,
		TargetTypeName: wrapName,
		FullGoType:     "[]*" + wrapName,
	})
}

func (em *EmittedModel) flattenParticles(
	mg *xsdloader.ModelGroup,
	parentRepeated bool,
	parentOptional bool,
) (elems []*xsdloader.Element, repeated []bool, optional []bool) {
	visitedGroups := make(map[string]bool)

	var walk func(cur *xsdloader.ModelGroup, rep bool, opt bool)
	walk = func(cur *xsdloader.ModelGroup, rep bool, opt bool) {
		if cur == nil {
			return
		}
		curRep := rep || (cur.MaxOccurs == "unbounded" || (cur.MaxOccurs != "" && cur.MaxOccurs != "1" && cur.MaxOccurs != "0"))
		curOpt := opt || cur.MinOccurs == "0" || cur.Kind == "choice"

		for _, item := range cur.Particles {
			if item.Element != nil {
				el := item.Element
				elRep := curRep || (el.MaxOccurs == "unbounded" || (el.MaxOccurs != "" && el.MaxOccurs != "1" && el.MaxOccurs != "0"))
				elOpt := curOpt || el.MinOccurs == "0"
				elems = append(elems, el)
				repeated = append(repeated, elRep)
				optional = append(optional, elOpt)
			} else if item.GroupRef != nil {
				gRef := item.GroupRef
				local := xsdloader.LocalName(gRef.Ref)
				if visitedGroups[local] {
					continue
				}
				visitedGroups[local] = true
				gRep := curRep || (gRef.MaxOccurs == "unbounded" || (gRef.MaxOccurs != "" && gRef.MaxOccurs != "1" && gRef.MaxOccurs != "0"))
				gOpt := curOpt || gRef.MinOccurs == "0"
				if targetGrp, ok := em.Graph.GroupByName[local]; ok && targetGrp.Particle != nil {
					// Special case: if targetGrp is "music-data" inside measure, and cur is measure's particle
					walk(targetGrp.Particle, gRep, gOpt)
				}
				visitedGroups[local] = false
			} else if item.ModelGroup != nil {
				walk(item.ModelGroup, curRep, curOpt)
			}
		}
	}

	walk(mg, parentRepeated, parentOptional)
	return
}

func (em *EmittedModel) buildElementField(
	ep *EmittedPackage,
	pkgScope *namer.Scope,
	fieldScope *namer.Scope,
	parentStruct *EmittedStruct,
	el *xsdloader.Element,
	repeated bool,
	optional bool,
) *EmittedField {
	rawName := el.Name
	elType := el.Type
	if el.Ref != "" {
		local := xsdloader.LocalName(el.Ref)
		rawName = local
		if targetEl, ok := em.Graph.ElementByName[local]; ok {
			elType = targetEl.Type
		}
	}
	if rawName == "" {
		return nil
	}

	goName := fieldScope.Unique(rawName)
	xmlTag := fmt.Sprintf("`xml:\"%s,omitempty\"`", rawName)

	// Case 1: Anonymous inline complexType (e.g., <xs:element name="part"><xs:complexType>...)
	if el.ComplexType != nil {
		// Create a named Gongstruct in the same package
		inlineName := pkgScope.Unique(parentStruct.GoName + "_" + namer.ToExportedIdentifier(rawName))
		if parentStruct.XSDName == "score-partwise" && (rawName == "part" || rawName == "measure") {
			inlineName = pkgScope.Unique(namer.ToExportedIdentifier(rawName))
		} else if parentStruct.GoName == "Part" && rawName == "measure" {
			inlineName = pkgScope.Unique("Measure")
		}
		inlineStruct := &EmittedStruct{
			GoName:    inlineName,
			XSDName:   rawName,
			PkgName:   ep.Name,
			SourceDoc: el.Documentation,
			SourceCT:  el.ComplexType,
		}
		ep.Structs = append(ep.Structs, inlineStruct)
		em.StructByKey["anon:"+ep.Name+":"+inlineName] = inlineStruct
		em.populateStructFields(ep, pkgScope, inlineStruct, el.ComplexType)

		return &EmittedField{
			GoName:         goName,
			XMLTag:         xmlTag,
			XMLNameLocal:   rawName,
			IsPointer:      !repeated,
			IsSlice:        repeated,
			TargetPkg:      ep.Name,
			TargetTypeName: inlineName,
			FullGoType:     qualifyGoType(ep.Name, ep.Name, inlineName, !repeated, repeated),
			Required:       !optional && !repeated,
		}
	}

	// Case 2: Named ComplexType
	typeLocal := xsdloader.LocalName(elType)
	if targetStruct, ok := em.StructByKey["ct:"+typeLocal]; ok {
		if targetStruct.PkgName != ep.Name {
			em.recordImport(ep, targetStruct.PkgName)
		}
		return &EmittedField{
			GoName:         goName,
			XMLTag:         xmlTag,
			XMLNameLocal:   rawName,
			IsPointer:      !repeated,
			IsSlice:        repeated,
			TargetPkg:      targetStruct.PkgName,
			TargetTypeName: targetStruct.GoName,
			FullGoType:     qualifyGoType(ep.Name, targetStruct.PkgName, targetStruct.GoName, !repeated, repeated),
			Required:       !optional && !repeated,
			IsEmptyElem:    typeLocal == "empty",
		}
	}

	// Case 3: SimpleType or built-in XSD scalar type
	goType, targetPkg, stRef := em.resolveSimpleGoType(ep, elType)
	if targetPkg != "" && targetPkg != ep.Name {
		em.recordImport(ep, targetPkg)
	}
	fullType := qualifyGoType(ep.Name, targetPkg, goType, false, false)
	isValSlice := false
	if repeated {
		// Gong structs only support slices of Gongstruct pointers; for repeated scalar elements,
		// use a single scalar string or wrap if needed. Wait: Does Gong support `[]string` on a Gongstruct?
		// Let's check: in Gong, only []*GongStruct slices are supported as Gong fields, or basic fields (string, int, float64, bool).
		// Wait: let's check if Gong allows `[]string` or ignores it or panics!
		// In `app/xsd/go/models/model_group.go`, how did `app/xsd` handle simpleType elements?
		// Line 143 of `app/xsd/go/models/model_group.go`:
		// `elem.GoIdentifier + " " + goType` (always scalar `goType`, never `[]string`!).
	}

	return &EmittedField{
		GoName:         goName,
		XMLTag:         xmlTag,
		XMLNameLocal:   rawName,
		IsValueSlice:   isValSlice,
		TargetPkg:      targetPkg,
		TargetTypeName: goType,
		FullGoType:     fullType,
		Required:       !optional && !repeated,
		DefaultVal:     el.Default,
		FixedVal:       el.Fixed,
		SimpleType:     stRef,
	}
}

func (em *EmittedModel) resolveSimpleGoType(ep *EmittedPackage, xsdType string) (goType string, targetPkg string, stRef *xsdloader.SimpleType) {
	if xsdType == "" {
		return "string", "", nil
	}
	prefix := xsdloader.Prefix(xsdType)
	local := xsdloader.LocalName(xsdType)

	// Check built-in XSD primitive/derived types
	switch local {
	case "string", "normalizedString", "token", "language", "Name", "NCName",
		"NMTOKEN", "NMTOKENS", "ID", "IDREF", "IDREFS", "ENTITY", "ENTITIES",
		"QName", "anyURI", "NOTATION", "date", "dateTime", "time", "duration",
		"gYear", "gYearMonth", "gMonth", "gMonthDay", "gDay", "hexBinary", "base64Binary",
		"decimal", "nonNegativeDecimal", "positiveDecimal":
		// Note: Keeping decimal as string or float64?
		// In XML round-trip, keeping XSD decimal/string as `string` preserves exact formatting (e.g., "0.8333" vs "0.833300", "40" vs "40.0")!
		// Wait: what about integer types?
		return "string", "", nil
	case "integer", "nonPositiveInteger", "negativeInteger", "long",
		"int", "short", "byte", "nonNegativeInteger", "unsignedLong", "unsignedInt",
		"unsignedShort", "unsignedByte", "positiveInteger":
		return "int", "", nil
	case "double", "float":
		return "float64", "", nil
	case "boolean":
		return "bool", "", nil
	}

	// Look up in EmittedEnum registry
	enumKey := "st:" + local
	if prefix == "xlink" {
		enumKey = "st:xlink:" + local
	}
	if enum, ok := em.EnumByKey[enumKey]; ok {
		st := em.Graph.SimpleTypeByName[enumKey]
		if st == nil {
			st = em.Graph.SimpleTypeByName[local]
		}
		return enum.GoName, enum.PkgName, st
	}
	if enum, ok := em.EnumByKey["st:"+local]; ok {
		return enum.GoName, enum.PkgName, em.Graph.SimpleTypeByName[local]
	}

	// Look up in SimpleType registry (non-enum restriction or union)
	stKey := "st:" + local
	if prefix == "xlink" {
		stKey = "st:xlink:" + local
	}
	if st, ok := em.Graph.SimpleTypeByName[stKey]; ok {
		if st.Restriction != nil && st.Restriction.Base != "" {
			baseGo, basePkg, _ := em.resolveSimpleGoType(ep, st.Restriction.Base)
			return baseGo, basePkg, st
		}
		return "string", "", st
	}
	if st, ok := em.Graph.SimpleTypeByName[local]; ok {
		if st.Restriction != nil && st.Restriction.Base != "" {
			baseGo, basePkg, _ := em.resolveSimpleGoType(ep, st.Restriction.Base)
			return baseGo, basePkg, st
		}
		return "string", "", st
	}

	return "string", "", nil
}

func (em *EmittedModel) recordImport(fromPkg *EmittedPackage, toPkgName string) {
	if toPkgName == "" || toPkgName == fromPkg.Name {
		return
	}
	if targetPkg, ok := em.Packages[toPkgName]; ok {
		fromPkg.Imports[toPkgName] = targetPkg.ImportPath
	}
}

func qualifyGoType(currentPkg, targetPkg, typeName string, isPtr, isSlice bool) string {
	qual := typeName
	if targetPkg != "" && targetPkg != currentPkg {
		qual = targetPkg + "." + typeName
	}
	if isSlice {
		return "[]*" + qual
	}
	if isPtr {
		return "*" + qual
	}
	return qual
}

// RenderModelsFile generates the `models.go` and `docs.go` content for an EmittedPackage.
func RenderModelsFile(ep *EmittedPackage) string {
	var sb strings.Builder
	sb.WriteString("// generated code by xsd2gong - do not edit\n")
	sb.WriteString(fmt.Sprintf("package %s\n\n", ep.Name))

	// Imports
	sb.WriteString("import (\n")
	sb.WriteString("\t\"encoding/xml\"\n")
	if len(ep.Imports) > 0 {
		sb.WriteString("\n")
		var impNames []string
		for k := range ep.Imports {
			impNames = append(impNames, k)
		}
		sort.Strings(impNames)
		for _, k := range impNames {
			sb.WriteString(fmt.Sprintf("\t\"%s\"\n", ep.Imports[k]))
		}
	}
	sb.WriteString(")\n\n")
	sb.WriteString("var _ xml.Attr\n\n")

	// Enums
	for _, en := range ep.Enums {
		if en.SourceDoc != "" {
			sb.WriteString(formatDocComment(en.GoName, en.SourceDoc))
		}
		sb.WriteString(fmt.Sprintf("type %s %s\n\n", en.GoName, en.BaseGo))
		sb.WriteString("const (\n")
		for _, v := range en.Values {
			sb.WriteString(fmt.Sprintf("\t%s %s = %q\n", v.ConstName, en.GoName, v.Literal))
		}
		sb.WriteString(")\n\n")
	}

	// Structs
	for _, st := range ep.Structs {
		if st.SourceDoc != "" {
			sb.WriteString(formatDocComment(st.GoName, st.SourceDoc))
		}
		sb.WriteString(fmt.Sprintf("type %s struct {\n", st.GoName))
		sb.WriteString("\tName string `xml:\"-\"`\n")
		if st.XMLNameTag != "" {
			sb.WriteString(fmt.Sprintf("\tXMLName xml.Name `xml:%q`\n", st.XMLNameTag))
		}
		for _, f := range st.Fields {
			sb.WriteString(fmt.Sprintf("\t%s %s %s\n", f.GoName, f.FullGoType, f.XMLTag))
		}
		sb.WriteString("}\n\n")
	}

	return sb.String()
}

func RenderDocsFile(ep *EmittedPackage) string {
	return fmt.Sprintf("// generated code by xsd2gong\npackage %s\n", ep.Name)
}

func formatDocComment(ident, doc string) string {
	lines := strings.Split(strings.TrimSpace(doc), "\n")
	var sb strings.Builder
	for i, l := range lines {
		trimmed := strings.TrimSpace(l)
		if trimmed == "" {
			continue
		}
		if i == 0 {
			sb.WriteString(fmt.Sprintf("// %s: %s\n", ident, trimmed))
		} else {
			sb.WriteString(fmt.Sprintf("// %s\n", trimmed))
		}
	}
	return sb.String()
}
