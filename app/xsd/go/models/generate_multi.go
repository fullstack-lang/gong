package models

import (
	"cmp"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// GeneratedField describes a single struct field generated from an XSD attribute or element.
type GeneratedField struct {
	GoName    string
	GoType    string
	XMLTag    string
	Comment   string
	ImportPkg *TargetPackage

	// Metadata for semantic enforcement and XML I/O
	XSDName       string
	RawXSDType    string
	SimpleTypeRef *SimpleType
	TargetCT      *ComplexType
	TargetPkg     *TargetPackage
	IsAttribute   bool
	IsRequired    bool
	IsSlice       bool
	IsPointer     bool
	ChoiceGroupID string
	FixedValue    string
	DefaultValue  string
}

// GeneratedStructInfo holds metadata about a generated Gongstruct for DSM tree, XML I/O, and semantic rules.
type GeneratedStructInfo struct {
	GoName       string
	Pkg          *TargetPackage
	SourceSchema *Schema
	ComplexType  *ComplexType
	RootElement  *Element
	IsRootElem   bool
	XMLRootName  string
	Fields       []GeneratedField
	ChoiceGroups [][]string // slices of Go field names belonging to the same <xs:choice>
}

// MultiGenerateContext holds all state across multi-package model & DSM code generation.
type MultiGenerateContext struct {
	SG              *SchemaGraph
	TargetModelsDir string

	// Precomputed fields for embeddable structs
	AttrStructFields map[*Attribute][]GeneratedField
	AgStructFields   map[*AttributeGroup][]GeneratedField
	GrpStructFields  map[*Group][]GeneratedField

	// Generated Gongstructs per package
	StructsByPkg map[*TargetPackage][]*GeneratedStructInfo
	RootStructs  []*GeneratedStructInfo
}

// GenerateMultiPackage generates the multi-package Go models (`go/models/models.go` and `go/models/<subpkg>/models.go` + `docs.go`).
func GenerateMultiPackage(sg *SchemaGraph, targetModelsDir string) (*MultiGenerateContext, error) {
	ctx := &MultiGenerateContext{
		SG:               sg,
		TargetModelsDir:  targetModelsDir,
		AttrStructFields: make(map[*Attribute][]GeneratedField),
		AgStructFields:   make(map[*AttributeGroup][]GeneratedField),
		GrpStructFields:  make(map[*Group][]GeneratedField),
		StructsByPkg:     make(map[*TargetPackage][]*GeneratedStructInfo),
	}

	// Precompute fields for all top-level Attributes, AttributeGroups, and Groups across all packages
	for _, pkg := range sg.TargetPackages {
		ctx.precomputePackageEmbeddables(pkg)
	}

	// Generate each TargetPackage in topological order
	for _, pkg := range sg.TargetPackages {
		if err := ctx.generateTargetPackageFiles(pkg); err != nil {
			return nil, err
		}
	}

	return ctx, nil
}

func (ctx *MultiGenerateContext) precomputePackageEmbeddables(pkg *TargetPackage) {
	// 1. Top-level Attributes
	for _, attr := range pkg.TopAttributes {
		schema := ctx.SG.SchemaOfAttribute[attr]
		usedNames := make(map[string]any)
		f := ctx.buildAttributeField(pkg, schema, attr, usedNames, true)
		if f.GoName != "" {
			ctx.AttrStructFields[attr] = []GeneratedField{f}
		}
	}

	// 2. AttributeGroups
	var computeAg func(ag *AttributeGroup, visited map[*AttributeGroup]bool) []GeneratedField
	computeAg = func(ag *AttributeGroup, visited map[*AttributeGroup]bool) []GeneratedField {
		if existing, ok := ctx.AgStructFields[ag]; ok {
			return existing
		}
		if visited[ag] {
			return nil
		}
		visited[ag] = true
		schema := ctx.SG.SchemaOfAttributeGroup[ag]
		usedNames := make(map[string]any)
		var fields []GeneratedField

		for _, attr := range ag.Attributes {
			f := ctx.buildAttributeField(pkg, schema, attr, usedNames, !pkg.IsRoot)
			if f.GoName != "" {
				fields = append(fields, f)
			}
		}
		for _, refAg := range ag.AttributeGroups {
			if refAg.Ref != "" {
				if targetAg := ctx.SG.ResolveAttributeGroup(schema, refAg.Ref); targetAg != nil {
					subFields := computeAg(targetAg, visited)
					for _, sf := range subFields {
						if _, exists := usedNames[sf.GoName]; !exists {
							usedNames[sf.GoName] = nil
							fields = append(fields, sf)
						}
					}
				}
			}
		}
		ctx.AgStructFields[ag] = fields
		return fields
	}
	for _, ag := range pkg.AttributeGroups {
		computeAg(ag, make(map[*AttributeGroup]bool))
	}

	// 3. Groups
	var computeGrp func(grp *Group, visited map[*Group]bool) []GeneratedField
	computeGrp = func(grp *Group, visited map[*Group]bool) []GeneratedField {
		if existing, ok := ctx.GrpStructFields[grp]; ok {
			return existing
		}
		if visited[grp] {
			return nil
		}
		visited[grp] = true
		schema := ctx.SG.SchemaOfGroup[grp]
		usedNames := make(map[string]any)
		seenElems := make(map[string]bool)
		var fields []GeneratedField
		var choiceGroups [][]string

		ctx.collectModelGroupFields(pkg, schema, &grp.ModelGroup, usedNames, seenElems, &fields, &choiceGroups, !pkg.IsRoot, "")
		ctx.GrpStructFields[grp] = fields
		return fields
	}
	for _, grp := range pkg.Groups {
		computeGrp(grp, make(map[*Group]bool))
	}
}

func (ctx *MultiGenerateContext) resolveSimpleGoType(
	currentPkg *TargetPackage,
	sourceSchema *Schema,
	rawType string,
	forcePrimitiveIfEnum bool,
) (goType string, st *SimpleType) {
	if rawType == "" {
		return "", nil
	}
	qname, isBuiltIn := ctx.SG.ResolveQName(sourceSchema, rawType)
	local := qname.LocalName
	if local == "" {
		local = Name(rawType)
	}

	if isBuiltIn {
		return mapBuiltInXSDTypeToGo(local), nil
	}

	// Look up SimpleType in symbol table
	st = ctx.SG.ResolveSimpleType(sourceSchema, rawType)
	if st != nil {
		stSchema := ctx.SG.SchemaOfSimpleType[st]
		stPkg := ctx.SG.PackageOfSchema[stSchema]
		if st.IsStringEnumerate() {
			if !forcePrimitiveIfEnum && stPkg == currentPkg && st.GoIdentifier != "" {
				return st.GoIdentifier, st
			}
			return "string", st
		}
		return ctx.resolveUnderlyingPrimitiveFromSimpleType(st, make(map[*SimpleType]bool)), st
	}

	// Fallback to built-in XSD type check
	if prim := mapBuiltInXSDTypeToGo(local); prim != "" {
		return prim, nil
	}
	if local == "empty" {
		return "string", nil
	}
	return "", nil
}

func (ctx *MultiGenerateContext) resolveUnderlyingPrimitiveFromSimpleType(st *SimpleType, visited map[*SimpleType]bool) string {
	if st == nil || visited[st] {
		return "string"
	}
	visited[st] = true
	stSchema := ctx.SG.SchemaOfSimpleType[st]

	if st.Restriction != nil && st.Restriction.Base != "" {
		baseLocal := Name(st.Restriction.Base)
		if prim := mapBuiltInXSDTypeToGo(baseLocal); prim != "" {
			return prim
		}
		if baseST := ctx.SG.ResolveSimpleType(stSchema, st.Restriction.Base); baseST != nil {
			return ctx.resolveUnderlyingPrimitiveFromSimpleType(baseST, visited)
		}
		return "string"
	}
	if st.List != nil || st.Union != nil {
		return "string"
	}
	return "string"
}

func mapBuiltInXSDTypeToGo(local string) string {
	switch local {
	case "dateTime", "date", "time", "duration", "gYearMonth", "gYear", "gMonthDay", "gDay", "gMonth",
		"string", "normalizedString", "token", "language", "Name", "NCName",
		"NMTOKEN", "NMTOKENS", "ID", "IDREF", "IDREFS", "ENTITY", "ENTITIES",
		"QName", "anyURI", "NOTATION", "decimal", "hexBinary", "base64Binary",
		"anySimpleType", "anyType":
		return "string"
	case "integer", "nonPositiveInteger", "negativeInteger", "long",
		"int", "short", "byte", "nonNegativeInteger", "unsignedLong", "unsignedInt",
		"unsignedShort", "unsignedByte", "positiveInteger":
		return "int"
	case "double", "float":
		return "float64"
	case "boolean":
		return "bool"
	}
	return ""
}

func (ctx *MultiGenerateContext) buildAttributeField(
	currentPkg *TargetPackage,
	sourceSchema *Schema,
	attr *Attribute,
	usedFieldNames map[string]any,
	forcePrimitiveIfEnum bool,
) GeneratedField {
	effectiveAttr := attr
	effectiveSchema := sourceSchema
	xmlAttrName := attr.NameXSD

	if attr.Ref != "" {
		if targetAttr := ctx.SG.ResolveAttribute(sourceSchema, attr.Ref); targetAttr != nil {
			effectiveAttr = targetAttr
			effectiveSchema = ctx.SG.SchemaOfAttribute[targetAttr]
			xmlAttrName = targetAttr.NameXSD
		} else {
			xmlAttrName = Name(attr.Ref)
		}
		if strings.HasPrefix(attr.Ref, "xlink:") {
			xmlAttrName = "http://www.w3.org/1999/xlink " + Name(attr.Ref)
		} else if strings.HasPrefix(attr.Ref, "xml:") {
			xmlAttrName = "http://www.w3.org/XML/1998/namespace " + Name(attr.Ref)
		}
	} else if sourceSchema != nil {
		if sourceSchema.TargetNamespace == "http://www.w3.org/1999/xlink" {
			xmlAttrName = "http://www.w3.org/1999/xlink " + attr.NameXSD
		} else if sourceSchema.TargetNamespace == "http://www.w3.org/XML/1998/namespace" {
			xmlAttrName = "http://www.w3.org/XML/1998/namespace " + attr.NameXSD
		}
	}

	rawName := effectiveAttr.NameXSD
	if rawName == "" {
		rawName = Name(attr.Ref)
	}
	if rawName == "" {
		return GeneratedField{}
	}

	goType, st := ctx.resolveSimpleGoType(currentPkg, effectiveSchema, effectiveAttr.Type, forcePrimitiveIfEnum)
	if goType == "" {
		goType = "string"
	}

	goIdentBase := xsdNameToGoIdentifier(rawName)
	if goIdentBase == "Name" {
		goIdentBase = "NameXSD"
	}

	var wid WithGoIdentifier
	computeGoIdentifier(goIdentBase, &wid, usedFieldNames)

	useReq := attr.Use == "required" || effectiveAttr.Use == "required"
	fixedVal := attr.Fixed
	if fixedVal == "" {
		fixedVal = effectiveAttr.Fixed
	}
	defaultVal := attr.Default
	if defaultVal == "" {
		defaultVal = effectiveAttr.Default
	}

	return GeneratedField{
		GoName:        wid.GoIdentifier,
		GoType:        goType,
		XMLTag:        fmt.Sprintf("`xml:\"%s,attr,omitempty\"`", xmlAttrName),
		Comment:       fmt.Sprintf("\t// generated from attribute %q", xmlAttrName),
		XSDName:       rawName,
		RawXSDType:    effectiveAttr.Type,
		SimpleTypeRef: st,
		IsAttribute:   true,
		IsRequired:    useReq,
		FixedValue:    fixedVal,
		DefaultValue:  defaultVal,
	}
}

func (ctx *MultiGenerateContext) collectModelGroupFields(
	currentPkg *TargetPackage,
	sourceSchema *Schema,
	mg *ModelGroup,
	usedFieldNames map[string]any,
	seenElems map[string]bool,
	fields *[]GeneratedField,
	choiceGroups *[][]string,
	forcePrimitiveIfEnum bool,
	activeChoiceID string,
) {
	if mg == nil {
		return
	}

	type orderedItem struct {
		order int
		kind  string
		elem  *Element
		grp   *Group
		seq   *Sequence
		ch    *Choice
		all   *All
	}

	var items []orderedItem
	for _, g := range mg.Groups {
		if g.Ref != "" {
			items = append(items, orderedItem{order: g.GetOrder(), kind: "group", grp: g})
		}
	}
	for _, e := range mg.Elements {
		items = append(items, orderedItem{order: e.GetOrder(), kind: "elem", elem: e})
	}
	for _, s := range mg.Sequences {
		items = append(items, orderedItem{order: s.GetOrder(), kind: "seq", seq: s})
	}
	for _, c := range mg.Choices {
		items = append(items, orderedItem{order: c.GetOrder(), kind: "choice", ch: c})
	}
	for _, a := range mg.Alls {
		items = append(items, orderedItem{order: a.GetOrder(), kind: "all", all: a})
	}

	slices.SortFunc(items, func(a, b orderedItem) int {
		return cmp.Compare(a.order, b.order)
	})

	for _, it := range items {
		switch it.kind {
		case "elem":
			ctx.emitElementFields(currentPkg, sourceSchema, it.elem, mg, usedFieldNames, seenElems, fields, forcePrimitiveIfEnum, activeChoiceID)
		case "group":
			if targetGrp := ctx.SG.ResolveGroup(sourceSchema, it.grp.Ref); targetGrp != nil {
				targetSchema := ctx.SG.SchemaOfGroup[targetGrp]
				ctx.collectModelGroupFields(currentPkg, targetSchema, &targetGrp.ModelGroup, usedFieldNames, seenElems, fields, choiceGroups, forcePrimitiveIfEnum, activeChoiceID)
			}
		case "seq":
			ctx.collectModelGroupFields(currentPkg, sourceSchema, &it.seq.ModelGroup, usedFieldNames, seenElems, fields, choiceGroups, forcePrimitiveIfEnum, activeChoiceID)
		case "all":
			ctx.collectModelGroupFields(currentPkg, sourceSchema, &it.all.ModelGroup, usedFieldNames, seenElems, fields, choiceGroups, forcePrimitiveIfEnum, activeChoiceID)
		case "choice":
			isChoiceUnbounded := it.ch.MaxOccurs == "unbounded" || computeIsBounded(it.ch)
			choiceID := fmt.Sprintf("choice_%d", it.ch.GetOrder())
			startIdx := len(*fields)
			ctx.collectModelGroupFields(currentPkg, sourceSchema, &it.ch.ModelGroup, usedFieldNames, seenElems, fields, choiceGroups, forcePrimitiveIfEnum, choiceID)
			// Only record mutual-exclusivity group if the choice itself is not unbounded and has >= 2 distinct element branches
			if !isChoiceUnbounded && activeChoiceID == "" && len(it.ch.Sequences) == 0 && len(it.ch.Choices) == 0 {
				var branchNames []string
				for i := startIdx; i < len(*fields); i++ {
					if (*fields)[i].ChoiceGroupID == choiceID {
						branchNames = append(branchNames, (*fields)[i].GoName)
					}
				}
				if len(branchNames) >= 2 {
					*choiceGroups = append(*choiceGroups, branchNames)
				}
			}
		}
	}
}

func (ctx *MultiGenerateContext) emitElementFields(
	currentPkg *TargetPackage,
	sourceSchema *Schema,
	elem *Element,
	enclosingMG *ModelGroup,
	usedFieldNames map[string]any,
	seenElems map[string]bool,
	fields *[]GeneratedField,
	forcePrimitiveIfEnum bool,
	activeChoiceID string,
) {
	if elem == nil {
		return
	}
	unbounded := computeIsBounded(elem) || computeIsBounded(enclosingMG)

	if elem.Ref != "" {
		targetElem := ctx.SG.ResolveElement(sourceSchema, elem.Ref)
		if targetElem == nil {
			return
		}
		targetSchema := ctx.SG.SchemaOfElement[targetElem]
		headQName := QName{Namespace: targetSchema.TargetNamespace, LocalName: targetElem.NameXSD}
		subMembers := ctx.SG.EligibleSubstitutionMembers(sourceSchema, headQName)

		if len(subMembers) > 0 {
			if targetElem.Abstract != "true" {
				ctx.emitSingleElementField(currentPkg, targetSchema, targetElem, unbounded, elem.MinOccurs, usedFieldNames, seenElems, fields, forcePrimitiveIfEnum, activeChoiceID)
			}
			for _, subEl := range subMembers {
				subSchema := ctx.SG.SchemaOfElement[subEl]
				ctx.emitSingleElementField(currentPkg, subSchema, subEl, unbounded, "0", usedFieldNames, seenElems, fields, forcePrimitiveIfEnum, activeChoiceID)
			}
			return
		}
		ctx.emitSingleElementField(currentPkg, targetSchema, targetElem, unbounded, elem.MinOccurs, usedFieldNames, seenElems, fields, forcePrimitiveIfEnum, activeChoiceID)
		return
	}

	ctx.emitSingleElementField(currentPkg, sourceSchema, elem, unbounded, elem.MinOccurs, usedFieldNames, seenElems, fields, forcePrimitiveIfEnum, activeChoiceID)
}

func (ctx *MultiGenerateContext) emitSingleElementField(
	currentPkg *TargetPackage,
	elemSchema *Schema,
	elem *Element,
	unbounded bool,
	minOccurs string,
	usedFieldNames map[string]any,
	seenElems map[string]bool,
	fields *[]GeneratedField,
	forcePrimitiveIfEnum bool,
	activeChoiceID string,
) {
	xmlElemName := elem.NameXSD
	if xmlElemName == "" {
		return
	}
	if seenElems[xmlElemName] {
		return
	}
	seenElems[xmlElemName] = true

	goIdentBase := xsdNameToGoIdentifier(xmlElemName)
	if goIdentBase == "Name" {
		goIdentBase = "NameXSD"
	}

	var wid WithGoIdentifier
	computeGoIdentifier(goIdentBase, &wid, usedFieldNames)
	goFieldName := wid.GoIdentifier

	sliceOrPtr := "*"
	if unbounded {
		sliceOrPtr = "[]*"
	}

	// 1. Check if element type is a SimpleType or built-in primitive
	if elem.Type != "" {
		goType, st := ctx.resolveSimpleGoType(currentPkg, elemSchema, elem.Type, forcePrimitiveIfEnum)
		if goType != "" {
			*fields = append(*fields, GeneratedField{
				GoName:        goFieldName,
				GoType:        goType,
				XMLTag:        fmt.Sprintf("`xml:\"%s,omitempty\"`", xmlElemName),
				Comment:       fmt.Sprintf("\t// generated from element %q of simple type %s", xmlElemName, elem.Type),
				XSDName:       xmlElemName,
				RawXSDType:    elem.Type,
				SimpleTypeRef: st,
				IsRequired:    minOccurs != "0" && activeChoiceID == "",
				ChoiceGroupID: activeChoiceID,
				FixedValue:    elem.Fixed,
				DefaultValue:  elem.Default,
			})
			return
		}

		// 2. Named ComplexType
		if targetCT := ctx.SG.ResolveComplexType(elemSchema, elem.Type); targetCT != nil {
			ctSchema := ctx.SG.SchemaOfComplexType[targetCT]
			ctPkg := ctx.SG.PackageOfSchema[ctSchema]
			typeExpr := targetCT.GoIdentifier
			var importPkg *TargetPackage
			if ctPkg != nil && ctPkg != currentPkg {
				typeExpr = ctPkg.ImportAlias + "." + targetCT.GoIdentifier
				importPkg = ctPkg
			}
			*fields = append(*fields, GeneratedField{
				GoName:        goFieldName,
				GoType:        sliceOrPtr + typeExpr,
				XMLTag:        fmt.Sprintf("`xml:\"%s,omitempty\"`", xmlElemName),
				Comment:       fmt.Sprintf("\t// generated from element %q of complex type %s", xmlElemName, targetCT.NameXSD),
				ImportPkg:     importPkg,
				XSDName:       xmlElemName,
				RawXSDType:    elem.Type,
				TargetCT:      targetCT,
				TargetPkg:     ctPkg,
				IsRequired:    minOccurs != "0" && !unbounded && activeChoiceID == "",
				IsSlice:       unbounded,
				IsPointer:     !unbounded,
				ChoiceGroupID: activeChoiceID,
			})
			return
		}

		// Fallback if type is xs:anyType or unresolved external type
		*fields = append(*fields, GeneratedField{
			GoName:  goFieldName,
			GoType:  "string",
			XMLTag:  fmt.Sprintf("`xml:\"%s,omitempty\"`", xmlElemName),
			Comment: fmt.Sprintf("\t// generated from element %q of fallback type %s", xmlElemName, elem.Type),
			XSDName: xmlElemName,
		})
		return
	}

	// 3. Inline anonymous ComplexType
	if elem.ComplexType != nil {
		targetCT := elem.ComplexType
		ctSchema := ctx.SG.SchemaOfComplexType[targetCT]
		ctPkg := ctx.SG.PackageOfSchema[ctSchema]
		typeExpr := targetCT.GoIdentifier
		var importPkg *TargetPackage
		if ctPkg != nil && ctPkg != currentPkg {
			typeExpr = ctPkg.ImportAlias + "." + targetCT.GoIdentifier
			importPkg = ctPkg
		}
		*fields = append(*fields, GeneratedField{
			GoName:        goFieldName,
			GoType:        sliceOrPtr + typeExpr,
			XMLTag:        fmt.Sprintf("`xml:\"%s,omitempty\"`", xmlElemName),
			Comment:       fmt.Sprintf("\t// generated from anonymous complex type within element %q", xmlElemName),
			ImportPkg:     importPkg,
			XSDName:       xmlElemName,
			TargetCT:      targetCT,
			TargetPkg:     ctPkg,
			IsRequired:    minOccurs != "0" && !unbounded && activeChoiceID == "",
			IsSlice:       unbounded,
			IsPointer:     !unbounded,
			ChoiceGroupID: activeChoiceID,
		})
		return
	}

	// Element with no type and no inline ComplexType defaults to string (xs:anyType)
	*fields = append(*fields, GeneratedField{
		GoName:  goFieldName,
		GoType:  "string",
		XMLTag:  fmt.Sprintf("`xml:\"%s,omitempty\"`", xmlElemName),
		Comment: fmt.Sprintf("\t// generated from untyped element %q", xmlElemName),
		XSDName: xmlElemName,
	})
}

func (ctx *MultiGenerateContext) buildComplexTypeBody(
	currentPkg *TargetPackage,
	ct *ComplexType,
	isRootElement bool,
	rootXMLName string,
) (bodyCode string, imports map[string]*TargetPackage, structInfo *GeneratedStructInfo) {
	imports = make(map[string]*TargetPackage)
	usedFieldNames := map[string]any{
		"Name":    nil,
		"XMLName": nil,
	}
	seenElems := make(map[string]bool)
	seenAttrs := make(map[string]bool)

	var allFlatFields []GeneratedField
	var choiceGroups [][]string
	var b strings.Builder

	if isRootElement && rootXMLName != "" {
		b.WriteString("\n\n\t// necessary since it is a root element")
		b.WriteString(fmt.Sprintf("\n\tXMLName xml.Name `xml:%q`", rootXMLName))
	}

	// Collect inheritance hierarchy for ComplexType (base ancestors first, then derived ct)
	chain := ctx.getComplexTypeInheritanceChain(ct)

	hasEnclosedText := false

	for idx, curCT := range chain {
		curSchema := ctx.SG.SchemaOfComplexType[curCT]
		isLeaf := idx == len(chain)-1

		// 1. Direct attributes on curCT
		var attrs []*Attribute
		var agRefs []*AttributeGroup
		attrs = append(attrs, curCT.Attributes...)
		agRefs = append(agRefs, curCT.AttributeGroups...)

		if curCT.Extension != nil {
			attrs = append(attrs, curCT.Extension.Attributes...)
			agRefs = append(agRefs, curCT.Extension.AttributeGroups...)
		}
		if curCT.ComplexContent != nil {
			if curCT.ComplexContent.Extension != nil {
				attrs = append(attrs, curCT.ComplexContent.Extension.Attributes...)
				agRefs = append(agRefs, curCT.ComplexContent.Extension.AttributeGroups...)
			}
			if curCT.ComplexContent.Restriction != nil {
				attrs = append(attrs, curCT.ComplexContent.Restriction.Attributes...)
				agRefs = append(agRefs, curCT.ComplexContent.Restriction.AttributeGroups...)
			}
		}
		if curCT.SimpleContent != nil {
			if curCT.SimpleContent.Extension != nil {
				attrs = append(attrs, curCT.SimpleContent.Extension.Attributes...)
				agRefs = append(agRefs, curCT.SimpleContent.Extension.AttributeGroups...)
			}
			if curCT.SimpleContent.Restriction != nil {
				attrs = append(attrs, curCT.SimpleContent.Restriction.Attributes...)
				agRefs = append(agRefs, curCT.SimpleContent.Restriction.AttributeGroups...)
			}
		}

		for _, attr := range attrs {
			attrKey := attr.NameXSD
			if attrKey == "" {
				attrKey = attr.Ref
			}
			if attrKey != "" && seenAttrs[attrKey] {
				continue
			}
			if attrKey != "" {
				seenAttrs[attrKey] = true
			}

			// Check if this is a cross-package top-level attribute ref that can be embedded as subpkg.Attr_<name>
			if attr.Ref != "" {
				if targetAttr := ctx.SG.ResolveAttribute(curSchema, attr.Ref); targetAttr != nil {
					targetSchema := ctx.SG.SchemaOfAttribute[targetAttr]
					targetPkg := ctx.SG.PackageOfSchema[targetSchema]
					if targetPkg != nil && targetPkg != currentPkg {
						if preFields, ok := ctx.AttrStructFields[targetAttr]; ok && len(preFields) == 1 {
							pf := preFields[0]
							if _, conflict := usedFieldNames[pf.GoName]; !conflict {
								usedFieldNames[pf.GoName] = nil
								imports[targetPkg.ImportAlias] = targetPkg
								b.WriteString(fmt.Sprintf("\n\n\t// generated from cross-package attribute ref %q", attr.Ref))
								b.WriteString(fmt.Sprintf("\n\t%s.%s", targetPkg.ImportAlias, targetAttr.GoIdentifier))
								allFlatFields = append(allFlatFields, pf)
								continue
							}
						}
					}
				}
			}

			f := ctx.buildAttributeField(currentPkg, curSchema, attr, usedFieldNames, false)
			if f.GoName != "" {
				b.WriteString("\n\n" + f.Comment)
				b.WriteString(fmt.Sprintf("\n\t%s %s %s", f.GoName, f.GoType, f.XMLTag))
				allFlatFields = append(allFlatFields, f)
			}
		}

		// 2. AttributeGroup references on curCT
		for _, agRef := range agRefs {
			if agRef.Ref == "" {
				continue
			}
			targetAg := ctx.SG.ResolveAttributeGroup(curSchema, agRef.Ref)
			if targetAg == nil {
				continue
			}
			targetSchema := ctx.SG.SchemaOfAttributeGroup[targetAg]
			targetPkg := ctx.SG.PackageOfSchema[targetSchema]
			agFields := ctx.AgStructFields[targetAg]

			canEmbed := true
			for _, af := range agFields {
				if _, exists := usedFieldNames[af.GoName]; exists {
					canEmbed = false
					break
				}
			}

			if canEmbed && len(agFields) > 0 {
				for _, af := range agFields {
					usedFieldNames[af.GoName] = nil
					seenAttrs[af.XSDName] = true
					allFlatFields = append(allFlatFields, af)
				}
				b.WriteString(fmt.Sprintf("\n\n\t// generated from attribute group %q", agRef.Ref))
				if targetPkg != nil && targetPkg != currentPkg {
					imports[targetPkg.ImportAlias] = targetPkg
					b.WriteString(fmt.Sprintf("\n\t%s.%s", targetPkg.ImportAlias, targetAg.GoIdentifier))
				} else {
					b.WriteString(fmt.Sprintf("\n\t%s", targetAg.GoIdentifier))
				}
			} else {
				// Inline non-colliding attributes from targetAg
				for _, af := range agFields {
					if seenAttrs[af.XSDName] {
						continue
					}
					seenAttrs[af.XSDName] = true
					var wid WithGoIdentifier
					computeGoIdentifier(af.GoName, &wid, usedFieldNames)
					afCopy := af
					afCopy.GoName = wid.GoIdentifier
					b.WriteString("\n\n" + afCopy.Comment)
					b.WriteString(fmt.Sprintf("\n\t%s %s %s", afCopy.GoName, afCopy.GoType, afCopy.XMLTag))
					allFlatFields = append(allFlatFields, afCopy)
				}
			}
		}

		// 3. SimpleContent chardata
		if !hasEnclosedText && curCT.SimpleContent != nil {
			var baseRaw string
			if curCT.SimpleContent.Extension != nil {
				baseRaw = curCT.SimpleContent.Extension.Base
			} else if curCT.SimpleContent.Restriction != nil {
				baseRaw = curCT.SimpleContent.Restriction.Base
			}
			if baseRaw != "" {
				goType, st := ctx.resolveSimpleGoType(currentPkg, curSchema, baseRaw, false)
				if goType == "" {
					// Base might be another ComplexType with SimpleContent
					goType = "string"
				}
				usedFieldNames["EnclosedText"] = nil
				hasEnclosedText = true
				b.WriteString("\n\n\t// simpleContent character data")
				b.WriteString(fmt.Sprintf("\n\tEnclosedText %s `xml:\",chardata\"`", goType))
				allFlatFields = append(allFlatFields, GeneratedField{
					GoName:        "EnclosedText",
					GoType:        goType,
					XMLTag:        "`xml:\",chardata\"`",
					XSDName:       "EnclosedText",
					RawXSDType:    baseRaw,
					SimpleTypeRef: st,
				})
			}
		}

		// 4. Elements / ModelGroups
		// If the derived type uses <complexContent><restriction>, only emit elements from the leaf restriction
		if !isLeaf && ct.ComplexContent != nil && ct.ComplexContent.Restriction != nil {
			continue
		}

		var mgs []*ModelGroup
		if curCT.Extension != nil {
			mgs = append(mgs, &curCT.Extension.ModelGroup)
		}
		if curCT.ComplexContent != nil {
			if curCT.ComplexContent.Extension != nil {
				mgs = append(mgs, &curCT.ComplexContent.Extension.ModelGroup)
			}
			if curCT.ComplexContent.Restriction != nil {
				mgs = append(mgs, &curCT.ComplexContent.Restriction.ModelGroup)
			}
		}
		mgs = append(mgs, &curCT.ModelGroup)

		for _, mg := range mgs {
			ctx.emitModelGroupInComplexType(currentPkg, curSchema, mg, usedFieldNames, seenElems, &b, imports, &allFlatFields, &choiceGroups)
		}
	}

	structInfo = &GeneratedStructInfo{
		GoName:       ct.GoIdentifier,
		Pkg:          currentPkg,
		SourceSchema: ctx.SG.SchemaOfComplexType[ct],
		ComplexType:  ct,
		IsRootElem:   isRootElement,
		XMLRootName:  rootXMLName,
		Fields:       allFlatFields,
		ChoiceGroups: choiceGroups,
	}

	return b.String(), imports, structInfo
}

func (ctx *MultiGenerateContext) emitModelGroupInComplexType(
	currentPkg *TargetPackage,
	sourceSchema *Schema,
	mg *ModelGroup,
	usedFieldNames map[string]any,
	seenElems map[string]bool,
	b *strings.Builder,
	imports map[string]*TargetPackage,
	allFlatFields *[]GeneratedField,
	choiceGroups *[][]string,
) {
	if mg == nil {
		return
	}

	// Check if mg is a simple sequence/group container that references top-level named Groups in the same package
	type orderedNode struct {
		order int
		kind  string
		grp   *Group
		elem  *Element
		seq   *Sequence
		ch    *Choice
		all   *All
	}
	var nodes []orderedNode
	for _, g := range mg.Groups {
		if g.Ref != "" {
			nodes = append(nodes, orderedNode{order: g.GetOrder(), kind: "group", grp: g})
		}
	}
	for _, e := range mg.Elements {
		nodes = append(nodes, orderedNode{order: e.GetOrder(), kind: "elem", elem: e})
	}
	for _, s := range mg.Sequences {
		nodes = append(nodes, orderedNode{order: s.GetOrder(), kind: "seq", seq: s})
	}
	for _, c := range mg.Choices {
		nodes = append(nodes, orderedNode{order: c.GetOrder(), kind: "choice", ch: c})
	}
	for _, a := range mg.Alls {
		nodes = append(nodes, orderedNode{order: a.GetOrder(), kind: "all", all: a})
	}
	slices.SortFunc(nodes, func(a, b orderedNode) int {
		return cmp.Compare(a.order, b.order)
	})

	for _, n := range nodes {
		switch n.kind {
		case "group":
			targetGrp := ctx.SG.ResolveGroup(sourceSchema, n.grp.Ref)
			if targetGrp == nil {
				continue
			}
			targetSchema := ctx.SG.SchemaOfGroup[targetGrp]
			targetPkg := ctx.SG.PackageOfSchema[targetSchema]
			grpFields := ctx.GrpStructFields[targetGrp]

			canEmbed := targetPkg == currentPkg && len(grpFields) > 0
			if canEmbed {
				for _, gf := range grpFields {
					if _, exists := usedFieldNames[gf.GoName]; exists || seenElems[gf.XSDName] {
						canEmbed = false
						break
					}
				}
			}

			if canEmbed {
				for _, gf := range grpFields {
					usedFieldNames[gf.GoName] = nil
					seenElems[gf.XSDName] = true
					if gf.ImportPkg != nil && gf.ImportPkg != currentPkg {
						imports[gf.ImportPkg.ImportAlias] = gf.ImportPkg
					}
					*allFlatFields = append(*allFlatFields, gf)
				}
				b.WriteString(fmt.Sprintf("\n\n\t// generated from group %q with order %d depth %d", n.grp.Ref, n.grp.Order, n.grp.Depth))
				b.WriteString(fmt.Sprintf("\n\t%s", targetGrp.GoIdentifier))
			} else {
				var inlinedFields []GeneratedField
				ctx.collectModelGroupFields(currentPkg, targetSchema, &targetGrp.ModelGroup, usedFieldNames, seenElems, &inlinedFields, choiceGroups, false, "")
				for _, f := range inlinedFields {
					if f.ImportPkg != nil && f.ImportPkg != currentPkg {
						imports[f.ImportPkg.ImportAlias] = f.ImportPkg
					}
					b.WriteString("\n\n" + f.Comment)
					b.WriteString(fmt.Sprintf("\n\t%s %s %s", f.GoName, f.GoType, f.XMLTag))
					*allFlatFields = append(*allFlatFields, f)
				}
			}

		case "seq":
			ctx.emitModelGroupInComplexType(currentPkg, sourceSchema, &n.seq.ModelGroup, usedFieldNames, seenElems, b, imports, allFlatFields, choiceGroups)

		case "all":
			ctx.emitModelGroupInComplexType(currentPkg, sourceSchema, &n.all.ModelGroup, usedFieldNames, seenElems, b, imports, allFlatFields, choiceGroups)

		case "elem":
			var elemFields []GeneratedField
			ctx.emitElementFields(currentPkg, sourceSchema, n.elem, mg, usedFieldNames, seenElems, &elemFields, false, "")
			for _, f := range elemFields {
				if f.ImportPkg != nil && f.ImportPkg != currentPkg {
					imports[f.ImportPkg.ImportAlias] = f.ImportPkg
				}
				b.WriteString("\n\n" + f.Comment)
				b.WriteString(fmt.Sprintf("\n\t%s %s %s", f.GoName, f.GoType, f.XMLTag))
				*allFlatFields = append(*allFlatFields, f)
			}

		case "choice":
			var chFields []GeneratedField
			isChoiceUnbounded := n.ch.MaxOccurs == "unbounded" || computeIsBounded(n.ch)
			choiceID := fmt.Sprintf("choice_%d", n.ch.GetOrder())
			ctx.collectModelGroupFields(currentPkg, sourceSchema, &n.ch.ModelGroup, usedFieldNames, seenElems, &chFields, choiceGroups, false, choiceID)
			var branchNames []string
			for _, f := range chFields {
				if f.ImportPkg != nil && f.ImportPkg != currentPkg {
					imports[f.ImportPkg.ImportAlias] = f.ImportPkg
				}
				b.WriteString("\n\n" + f.Comment)
				b.WriteString(fmt.Sprintf("\n\t%s %s %s", f.GoName, f.GoType, f.XMLTag))
				*allFlatFields = append(*allFlatFields, f)
				if f.ChoiceGroupID == choiceID {
					branchNames = append(branchNames, f.GoName)
				}
			}
			if !isChoiceUnbounded && len(n.ch.Sequences) == 0 && len(n.ch.Choices) == 0 && len(branchNames) >= 2 {
				*choiceGroups = append(*choiceGroups, branchNames)
			}
		}
	}
}

func (ctx *MultiGenerateContext) getComplexTypeInheritanceChain(ct *ComplexType) []*ComplexType {
	var chain []*ComplexType
	visited := make(map[*ComplexType]bool)
	curr := ct
	for curr != nil && !visited[curr] {
		visited[curr] = true
		chain = append(chain, curr)
		schema := ctx.SG.SchemaOfComplexType[curr]
		var baseRef string
		if curr.Extension != nil && curr.Extension.Base != "" {
			baseRef = curr.Extension.Base
		} else if curr.ComplexContent != nil {
			if curr.ComplexContent.Extension != nil && curr.ComplexContent.Extension.Base != "" {
				baseRef = curr.ComplexContent.Extension.Base
			} else if curr.ComplexContent.Restriction != nil && curr.ComplexContent.Restriction.Base != "" {
				baseRef = curr.ComplexContent.Restriction.Base
			}
		}
		if baseRef == "" {
			break
		}
		curr = ctx.SG.ResolveComplexType(schema, baseRef)
	}
	slices.Reverse(chain)
	return chain
}

func (ctx *MultiGenerateContext) generateTargetPackageFiles(pkg *TargetPackage) error {
	pkgDir := ctx.TargetModelsDir
	if !pkg.IsRoot {
		pkgDir = filepath.Join(ctx.TargetModelsDir, pkg.RelDir)
	}
	if err := os.MkdirAll(pkgDir, 0o755); err != nil {
		return fmt.Errorf("failed creating package directory %s: %w", pkgDir, err)
	}

	// Ensure docs.go exists in every package (required by Gong's DiscoverModelDependencies)
	docsPath := filepath.Join(pkgDir, "docs.go")
	docsContent := fmt.Sprintf("// Package %s provides generated XSD model structs.\npackage %s\n", pkg.PackageName, pkg.PackageName)
	if err := os.WriteFile(docsPath, []byte(docsContent), 0o644); err != nil {
		return fmt.Errorf("failed writing %s: %w", docsPath, err)
	}

	imports := make(map[string]*TargetPackage)

	type particleCode struct {
		order int
		code  string
	}
	var enumCodes []particleCode
	var structCodes []particleCode

	// 1. Enums from SimpleTypes
	for _, st := range pkg.SimpleTypes {
		if st.IsStringEnumerate() {
			code := st.generateGongEnum()
			enumCodes = append(enumCodes, particleCode{
				order: st.GetOrder(),
				code:  code,
			})
		}
	}

	// 2. Top-level Attributes as embeddable Attr_<name> structs (in sub-packages or when referenced)
	if !pkg.IsRoot {
		for _, attr := range pkg.TopAttributes {
			if fields, ok := ctx.AttrStructFields[attr]; ok && len(fields) == 1 {
				f := fields[0]
				var sb strings.Builder
				sb.WriteString("\n\n" + f.Comment)
				sb.WriteString(fmt.Sprintf("\n\t%s %s %s", f.GoName, f.GoType, f.XMLTag))
				tmp := Replace3(
					Level1Code[Level1UnNamedStructCode],
					"{{"+string(rune(Level2Structname))+"}}", attr.GoIdentifier,
					"{{"+string(rune(Level2Source))+"}}", fmt.Sprintf("top-level attribute %q", attr.NameXSD),
					"{{"+string(rune(Level2Fields))+"}}", sb.String(),
				)
				structCodes = append(structCodes, particleCode{
					order: attr.GetOrder(),
					code:  tmp,
				})
			}
		}
	}

	// 3. AttributeGroups
	for _, ag := range pkg.AttributeGroups {
		agSchema := ctx.SG.SchemaOfAttributeGroup[ag]
		usedNames := make(map[string]any)
		var sb strings.Builder

		for _, attr := range ag.Attributes {
			if attr.Ref != "" {
				if targetAttr := ctx.SG.ResolveAttribute(agSchema, attr.Ref); targetAttr != nil {
					targetSchema := ctx.SG.SchemaOfAttribute[targetAttr]
					targetPkg := ctx.SG.PackageOfSchema[targetSchema]
					if !targetPkg.IsRoot {
						if preFields, ok := ctx.AttrStructFields[targetAttr]; ok && len(preFields) == 1 {
							pf := preFields[0]
							if _, conflict := usedNames[pf.GoName]; !conflict {
								usedNames[pf.GoName] = nil
								if targetPkg != pkg {
									imports[targetPkg.ImportAlias] = targetPkg
									sb.WriteString(fmt.Sprintf("\n\n\t// generated from cross-package attribute %q", attr.Ref))
									sb.WriteString(fmt.Sprintf("\n\t%s.%s", targetPkg.ImportAlias, targetAttr.GoIdentifier))
								} else {
									sb.WriteString(fmt.Sprintf("\n\n\t// generated from attribute %q", attr.Ref))
									sb.WriteString(fmt.Sprintf("\n\t%s", targetAttr.GoIdentifier))
								}
								continue
							}
						}
					}
				}
			}
			f := ctx.buildAttributeField(pkg, agSchema, attr, usedNames, !pkg.IsRoot)
			if f.GoName != "" {
				sb.WriteString("\n\n" + f.Comment)
				sb.WriteString(fmt.Sprintf("\n\t%s %s %s", f.GoName, f.GoType, f.XMLTag))
			}
		}

		for _, refAg := range ag.AttributeGroups {
			if refAg.Ref == "" {
				continue
			}
			targetAg := ctx.SG.ResolveAttributeGroup(agSchema, refAg.Ref)
			if targetAg == nil {
				continue
			}
			targetSchema := ctx.SG.SchemaOfAttributeGroup[targetAg]
			targetPkg := ctx.SG.PackageOfSchema[targetSchema]
			agFields := ctx.AgStructFields[targetAg]

			canEmbed := len(agFields) > 0
			for _, af := range agFields {
				if _, exists := usedNames[af.GoName]; exists {
					canEmbed = false
					break
				}
			}
			if canEmbed {
				for _, af := range agFields {
					usedNames[af.GoName] = nil
				}
				sb.WriteString(fmt.Sprintf("\n\n\t// generated from attribute group %q", refAg.Ref))
				if targetPkg != nil && targetPkg != pkg {
					imports[targetPkg.ImportAlias] = targetPkg
					sb.WriteString(fmt.Sprintf("\n\t%s.%s", targetPkg.ImportAlias, targetAg.GoIdentifier))
				} else {
					sb.WriteString(fmt.Sprintf("\n\t%s", targetAg.GoIdentifier))
				}
			} else {
				for _, af := range agFields {
					var wid WithGoIdentifier
					computeGoIdentifier(af.GoName, &wid, usedNames)
					sb.WriteString("\n\n" + af.Comment)
					sb.WriteString(fmt.Sprintf("\n\t%s %s %s", wid.GoIdentifier, af.GoType, af.XMLTag))
				}
			}
		}

		tmp := Replace3(
			Level1Code[Level1UnNamedStructCode],
			"{{"+string(rune(Level2Structname))+"}}", ag.GoIdentifier,
			"{{"+string(rune(Level2Source))+"}}", fmt.Sprintf("named attribute group %q", ag.NameXSD),
			"{{"+string(rune(Level2Fields))+"}}", sb.String(),
		)
		structCodes = append(structCodes, particleCode{
			order: ag.GetOrder(),
			code:  tmp,
		})
	}

	// 4. Groups (same-package unnamed embeddable structs)
	for _, grp := range pkg.Groups {
		grpSchema := ctx.SG.SchemaOfGroup[grp]
		usedNames := make(map[string]any)
		seenElems := make(map[string]bool)
		var sb strings.Builder
		var dummyFlat []GeneratedField
		var dummyChoices [][]string

		ctx.emitModelGroupInComplexType(pkg, grpSchema, &grp.ModelGroup, usedNames, seenElems, &sb, imports, &dummyFlat, &dummyChoices)

		tmp := Replace3(
			Level1Code[Level1UnNamedStructCode],
			"{{"+string(rune(Level2Structname))+"}}", grp.GoIdentifier,
			"{{"+string(rune(Level2Source))+"}}", fmt.Sprintf("named group %q", grp.NameXSD),
			"{{"+string(rune(Level2Fields))+"}}", sb.String(),
		)
		structCodes = append(structCodes, particleCode{
			order: grp.GetOrder(),
			code:  tmp,
		})
	}

	// 5. ComplexTypes (named and nested anonymous)
	gongStructCount := 0
	for _, ct := range pkg.ComplexTypes {
		body, ctImports, info := ctx.buildComplexTypeBody(pkg, ct, false, "")
		for k, v := range ctImports {
			imports[k] = v
		}
		ctx.StructsByPkg[pkg] = append(ctx.StructsByPkg[pkg], info)
		gongStructCount++

		var comment string
		if ct.Annotation != nil && len(ct.Annotation.Documentations) > 0 {
			comment = PrefixLines(ct.Annotation.Documentations[0].Text)
		}
		sourceDesc := fmt.Sprintf("named complex type %q", ct.NameXSD)
		if ct.IsAnonymous && ct.OuterElement != nil {
			sourceDesc = fmt.Sprintf("within outer element %q", ct.OuterElement.NameXSD)
		}
		tmp := Replace4(
			Level1Code[Level1NamedStructCode],
			"{{"+string(rune(Level2Comment))+"}}", comment,
			"{{"+string(rune(Level2Structname))+"}}", ct.GoIdentifier,
			"{{"+string(rune(Level2Source))+"}}", sourceDesc,
			"{{"+string(rune(Level2Fields))+"}}", body,
		)
		structCodes = append(structCodes, particleCode{
			order: ct.GetOrder(),
			code:  tmp,
		})
	}

	// 6. Top-level Elements (with inline ComplexType, or root elements referencing a named ComplexType with a distinct name)
	for _, el := range pkg.TopElements {
		elSchema := ctx.SG.SchemaOfElement[el]
		if el.ComplexType != nil {
			el.ComplexType.GoIdentifier = el.GoIdentifier
			body, ctImports, info := ctx.buildComplexTypeBody(pkg, el.ComplexType, pkg.IsRoot, el.NameXSD)
			info.GoName = el.GoIdentifier
			info.RootElement = el
			for k, v := range ctImports {
				imports[k] = v
			}
			ctx.StructsByPkg[pkg] = append(ctx.StructsByPkg[pkg], info)
			if pkg.IsRoot {
				ctx.RootStructs = append(ctx.RootStructs, info)
			}
			gongStructCount++

			var comment string
			if el.Annotation != nil && len(el.Annotation.Documentations) > 0 {
				comment = PrefixLines(el.Annotation.Documentations[0].Text)
			}
			tmp := Replace4(
				Level1Code[Level1NamedStructCode],
				"{{"+string(rune(Level2Comment))+"}}", comment,
				"{{"+string(rune(Level2Structname))+"}}", el.GoIdentifier,
				"{{"+string(rune(Level2Source))+"}}", fmt.Sprintf("element %s within schema", el.NameXSD),
				"{{"+string(rune(Level2Fields))+"}}", body,
			)
			structCodes = append(structCodes, particleCode{
				order: el.GetOrder(),
				code:  tmp,
			})
		} else if pkg.IsRoot && el.Abstract != "true" && el.Type != "" {
			if targetCT := ctx.SG.ResolveComplexType(elSchema, el.Type); targetCT != nil {
				if xsdNameToGoIdentifier(el.NameXSD) != xsdNameToGoIdentifier(targetCT.NameXSD) {
					body, ctImports, info := ctx.buildComplexTypeBody(pkg, targetCT, true, el.NameXSD)
					info.GoName = el.GoIdentifier
					info.RootElement = el
					for k, v := range ctImports {
						imports[k] = v
					}
					ctx.StructsByPkg[pkg] = append(ctx.StructsByPkg[pkg], info)
					ctx.RootStructs = append(ctx.RootStructs, info)
					gongStructCount++

					var comment string
					if el.Annotation != nil && len(el.Annotation.Documentations) > 0 {
						comment = PrefixLines(el.Annotation.Documentations[0].Text)
					}
					tmp := Replace4(
						Level1Code[Level1NamedStructCode],
						"{{"+string(rune(Level2Comment))+"}}", comment,
						"{{"+string(rune(Level2Structname))+"}}", el.GoIdentifier,
						"{{"+string(rune(Level2Source))+"}}", fmt.Sprintf("root element %s of type %s", el.NameXSD, el.Type),
						"{{"+string(rune(Level2Fields))+"}}", body,
					)
					structCodes = append(structCodes, particleCode{
						order: el.GetOrder(),
						code:  tmp,
					})
				} else {
					// Record existing named ComplexType as a RootStruct candidate
					for _, sInfo := range ctx.StructsByPkg[pkg] {
						if sInfo.ComplexType == targetCT {
							sInfo.IsRootElem = true
							sInfo.XMLRootName = el.NameXSD
							sInfo.RootElement = el
							ctx.RootStructs = append(ctx.RootStructs, sInfo)
							break
						}
					}
				}
			}
		}
	}

	// Ensure every sub-package has at least one Gongstruct so StageSet and Gong discovery always include it
	if gongStructCount == 0 {
		caser := cases.Title(language.English)
		synthName := caser.String(pkg.PackageName) + "_SchemaElement"
		synthCode := fmt.Sprintf("\n\n// %s synthetic Gongstruct for schema package %s\ntype %s struct {\n\tName string `xml:\"-\"`\n}", synthName, pkg.PackageName, synthName)
		structCodes = append(structCodes, particleCode{
			order: 999999,
			code:  synthCode,
		})
	}

	// In the root package, ensure all sub-packages in SG.TargetPackages are imported in models.go
	// so Gong's DiscoverModelDependencies discovers all sub-packages in topological order.
	var dummyPkgRefs strings.Builder
	if pkg.IsRoot {
		for _, subPkg := range ctx.SG.TargetPackages {
			if !subPkg.IsRoot {
				imports[subPkg.ImportAlias] = subPkg
			}
		}
	}

	// Build import block
	var importLines []string
	importLines = append(importLines, "\t\"encoding/xml\"")
	var sortedAliases []string
	for alias := range imports {
		sortedAliases = append(sortedAliases, alias)
	}
	slices.Sort(sortedAliases)
	for _, alias := range sortedAliases {
		impPkg := imports[alias]
		if alias == filepath.Base(impPkg.ImportPath) {
			importLines = append(importLines, fmt.Sprintf("\t%q", impPkg.ImportPath))
		} else {
			importLines = append(importLines, fmt.Sprintf("\t%s %q", alias, impPkg.ImportPath))
		}
		// Add a blank var reference so unused import errors never happen
		var refType string
		if len(ctx.StructsByPkg[impPkg]) > 0 {
			refType = ctx.StructsByPkg[impPkg][0].GoName
		} else {
			caser := cases.Title(language.English)
			refType = caser.String(impPkg.PackageName) + "_SchemaElement"
		}
		dummyPkgRefs.WriteString(fmt.Sprintf("\nvar _ *%s.%s", alias, refType))
	}

	slices.SortFunc(enumCodes, func(a, b particleCode) int {
		return cmp.Compare(a.order, b.order)
	})
	slices.SortFunc(structCodes, func(a, b particleCode) int {
		return cmp.Compare(a.order, b.order)
	})

	var allEnums strings.Builder
	for _, ec := range enumCodes {
		allEnums.WriteString(ec.code)
	}
	var allStructs strings.Builder
	for _, sc := range structCodes {
		allStructs.WriteString(sc.code)
	}

	var fileBuilder strings.Builder
	fileBuilder.WriteString("// generated code - do not edit\n")
	fileBuilder.WriteString(fmt.Sprintf("package %s\n\n", pkg.PackageName))
	fileBuilder.WriteString("import (\n")
	fileBuilder.WriteString(strings.Join(importLines, "\n"))
	fileBuilder.WriteString("\n)\n\n")
	fileBuilder.WriteString("// to avoid compilation error if no xml element is generated\n")
	fileBuilder.WriteString("var _ xml.Attr")
	fileBuilder.WriteString(dummyPkgRefs.String())
	fileBuilder.WriteString("\n\n// insertion point for enums declaration")
	fileBuilder.WriteString(allEnums.String())
	fileBuilder.WriteString("\n\n// insertion point for gongstructs declarations")
	fileBuilder.WriteString(allStructs.String())
	fileBuilder.WriteString("\n")

	modelsFilePath := filepath.Join(pkgDir, "models.go")
	if err := os.WriteFile(modelsFilePath, []byte(fileBuilder.String()), 0o644); err != nil {
		return fmt.Errorf("failed writing %s: %w", modelsFilePath, err)
	}

	return nil
}
