package models

import (
	"cmp"
	"encoding/xml"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"

	"golang.org/x/mod/modfile"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// QName represents a namespace-qualified XML Schema symbol name.
type QName struct {
	Namespace string
	LocalName string
}

// TargetPackage represents a generated Go package (root `models` or a sub-package `models/<subpkg>`).
type TargetPackage struct {
	PackageName string // "models" for root, or lowercase identifier e.g. "xlink", "xml", "gml"
	RelDir      string // "" for root, or relative directory under go/models
	ImportPath  string // Full Go import path, e.g. "github.com/.../go/models/xlink"
	ImportAlias string // Import alias in root models / StageSet
	StageField  string // Field name on root Stage, e.g. "" for root, or "XlinkStage"
	IsRoot      bool
	Schemas     []*Schema

	// Collected constructs belonging to this package (in deterministic order)
	ComplexTypes    []*ComplexType
	SimpleTypes     []*SimpleType
	Groups          []*Group
	AttributeGroups []*AttributeGroup
	TopElements     []*Element
	TopAttributes   []*Attribute
}

// SchemaGraph holds the transitive closure of XSD schemas, symbol tables, and target Go packages.
type SchemaGraph struct {
	RootSchema        *Schema
	LoadedSchemas     map[string]*Schema // absFilePath -> *Schema
	SchemasInOrder    []*Schema
	FallbackFileIndex map[string]string            // base filename -> absFilePath
	DirectSchemaDeps  map[*Schema]map[*Schema]bool // schema -> directly included/imported schemas

	// Ownership maps (construct -> owning *Schema)
	SchemaOfComplexType    map[*ComplexType]*Schema
	SchemaOfSimpleType     map[*SimpleType]*Schema
	SchemaOfGroup          map[*Group]*Schema
	SchemaOfAttributeGroup map[*AttributeGroup]*Schema
	SchemaOfElement        map[*Element]*Schema
	SchemaOfAttribute      map[*Attribute]*Schema

	// Per-schema collected constructs (including inline/anonymous ones)
	AllComplexTypesBySchema map[*Schema][]*ComplexType
	AllSimpleTypesBySchema  map[*Schema][]*SimpleType
	AllElementsBySchema     map[*Schema][]*Element
	AllAttributesBySchema   map[*Schema][]*Attribute

	// Global QName symbol tables
	ComplexTypes            map[QName]*ComplexType
	SimpleTypes             map[QName]*SimpleType
	Elements                map[QName]*Element
	Attributes              map[QName]*Attribute
	Groups                  map[QName]*Group
	AttributeGroups         map[QName]*AttributeGroup
	SubstitutionGroupResult map[QName][]*Element // head QName -> concrete substitutable top-level elements

	// Fallback local-name symbol tables (used when namespace is omitted or unqualified)
	ComplexTypesByLocal    map[string]*ComplexType
	SimpleTypesByLocal     map[string]*SimpleType
	ElementsByLocal        map[string]*Element
	AttributesByLocal      map[string]*Attribute
	GroupsByLocal          map[string]*Group
	AttributeGroupsByLocal map[string]*AttributeGroup

	// Package mapping
	PackageOfSchema map[*Schema]*TargetPackage
	TargetPackages  []*TargetPackage // Topological order (leaves first, root `models` last)
	RootPackage     *TargetPackage

	RootImportPath string // Full Go import path of root `go/models`
}

// LoadSchemaGraph loads rootXSDPath and all transitive xs:include and xs:import schemas,
// stages them into stage, builds the global symbol table, and resolves TargetPackages.
func LoadSchemaGraph(stage *Stage, rootXSDPath string, targetModelsDir string) (*SchemaGraph, error) {
	rootAbs, err := filepath.Abs(rootXSDPath)
	if err != nil {
		return nil, fmt.Errorf("failed resolving root XSD path %s: %w", rootXSDPath, err)
	}
	rootAbs = filepath.Clean(rootAbs)

	sg := &SchemaGraph{
		LoadedSchemas:           make(map[string]*Schema),
		FallbackFileIndex:       make(map[string]string),
		DirectSchemaDeps:        make(map[*Schema]map[*Schema]bool),
		SchemaOfComplexType:     make(map[*ComplexType]*Schema),
		SchemaOfSimpleType:      make(map[*SimpleType]*Schema),
		SchemaOfGroup:           make(map[*Group]*Schema),
		SchemaOfAttributeGroup:  make(map[*AttributeGroup]*Schema),
		SchemaOfElement:         make(map[*Element]*Schema),
		SchemaOfAttribute:       make(map[*Attribute]*Schema),
		AllComplexTypesBySchema: make(map[*Schema][]*ComplexType),
		AllSimpleTypesBySchema:  make(map[*Schema][]*SimpleType),
		AllElementsBySchema:     make(map[*Schema][]*Element),
		AllAttributesBySchema:   make(map[*Schema][]*Attribute),
		ComplexTypes:            make(map[QName]*ComplexType),
		SimpleTypes:             make(map[QName]*SimpleType),
		Elements:                make(map[QName]*Element),
		Attributes:              make(map[QName]*Attribute),
		Groups:                  make(map[QName]*Group),
		AttributeGroups:         make(map[QName]*AttributeGroup),
		SubstitutionGroupResult: make(map[QName][]*Element),
		ComplexTypesByLocal:     make(map[string]*ComplexType),
		SimpleTypesByLocal:      make(map[string]*SimpleType),
		ElementsByLocal:         make(map[string]*Element),
		AttributesByLocal:       make(map[string]*Attribute),
		GroupsByLocal:           make(map[string]*Group),
		AttributeGroupsByLocal:  make(map[string]*AttributeGroup),
		PackageOfSchema:         make(map[*Schema]*TargetPackage),
	}

	// Build fallback file index of nearby .xsd files for catalog/HTTP schemaLocations
	sg.buildFallbackIndex(rootAbs)

	// Determine Go module import path for targetModelsDir
	sg.RootImportPath = computeRootModelsImportPath(targetModelsDir, rootAbs)

	// Reset global XML Order counter before loading schemas
	Order = 1
	Depth = 0

	rootSchema, err := sg.loadSchemaRecursive(rootAbs, rootAbs, "")
	if err != nil {
		return nil, err
	}
	rootSchema.IsRootSchema = true
	sg.RootSchema = rootSchema

	// Stage all loaded schemas into the XSD stage
	for _, schema := range sg.SchemasInOrder {
		stage.StageBranch(schema)
	}
	stage.ComputeReverseMaps()

	// Collect per-schema inventory, compute anonymous types, and build global symbol tables
	sg.buildInventoriesAndSymbolTables()

	// Prune constructs in imported/included schemas that are not transitively reachable from RootSchema
	sg.pruneUnreachableConstructs()

	// Assign names on the XSD stage for probe inspection
	PostProcessingNames(stage)

	// Compute TargetPackages via symbol-level dependency graph and Tarjan's SCC algorithm
	sg.computeTargetPackages()

	// Compute GoIdentifiers per TargetPackage (preventing identifier collisions within each package)
	sg.assignGoIdentifiersPerPackage()

	return sg, nil
}

func (sg *SchemaGraph) buildFallbackIndex(rootAbs string) {
	// Walk up to 2 ancestor levels above rootAbs to index .xsd files (e.g. app/xsd/tests/)
	searchRoot := filepath.Dir(rootAbs)
	for range 2 {
		parent := filepath.Dir(searchRoot)
		if parent == searchRoot {
			break
		}
		if filepath.Base(searchRoot) == "tests" || filepath.Base(parent) == "gong" {
			break
		}
		searchRoot = parent
	}

	_ = filepath.WalkDir(searchRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			name := d.Name()
			if name == ".git" || name == "node_modules" || name == "vendor" {
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasSuffix(strings.ToLower(d.Name()), ".xsd") {
			base := d.Name()
			// Prefer paths closer to rootAbs if multiple exist
			if existing, ok := sg.FallbackFileIndex[base]; ok {
				if filepath.Dir(path) == filepath.Dir(rootAbs) {
					sg.FallbackFileIndex[base] = path
				} else {
					_ = existing
				}
			} else {
				sg.FallbackFileIndex[base] = path
			}
		}
		return nil
	})
}

func computeRootModelsImportPath(targetModelsDir string, rootXSDPath string) string {
	startDir := targetModelsDir
	if startDir == "" {
		startDir = filepath.Dir(rootXSDPath)
	}
	absDir, err := filepath.Abs(startDir)
	if err != nil {
		return "models"
	}
	// If targetModelsDir is a file path (e.g. ".../models/models.go"), take its directory
	if strings.HasSuffix(absDir, ".go") {
		absDir = filepath.Dir(absDir)
	}

	curr := absDir
	for range 12 {
		goModPath := filepath.Join(curr, "go.mod")
		if buf, err := os.ReadFile(goModPath); err == nil {
			if mf, errParse := modfile.Parse(goModPath, buf, nil); errParse == nil && mf.Module != nil {
				modPath := mf.Module.Mod.Path
				rel, errRel := filepath.Rel(curr, absDir)
				if errRel == nil {
					return filepath.ToSlash(filepath.Join(modPath, rel))
				}
			}
			break
		}
		parent := filepath.Dir(curr)
		if parent == curr {
			break
		}
		curr = parent
	}
	return "models"
}

func (sg *SchemaGraph) resolveSchemaLocation(currentSchemaPath string, rootSchemaPath string, namespace string, schemaLocation string) string {
	// 1. Direct relative path from currentSchemaPath
	if schemaLocation != "" && !strings.HasPrefix(schemaLocation, "http://") && !strings.HasPrefix(schemaLocation, "https://") {
		cand := filepath.Clean(filepath.Join(filepath.Dir(currentSchemaPath), filepath.FromSlash(schemaLocation)))
		if info, err := os.Stat(cand); err == nil && !info.IsDir() {
			return cand
		}
		candRoot := filepath.Clean(filepath.Join(filepath.Dir(rootSchemaPath), filepath.FromSlash(schemaLocation)))
		if info, err := os.Stat(candRoot); err == nil && !info.IsDir() {
			return candRoot
		}
	}

	// 2. Fallback by base file name or well-known namespace
	baseName := ""
	if schemaLocation != "" {
		baseName = filepath.Base(filepath.FromSlash(schemaLocation))
	} else {
		switch namespace {
		case "http://www.w3.org/XML/1998/namespace":
			baseName = "xml.xsd"
		case "http://www.w3.org/1999/xlink":
			baseName = "xlink.xsd"
		}
	}

	if baseName != "" && baseName != "." && baseName != "/" {
		// Check current dir and root dir first
		cand1 := filepath.Join(filepath.Dir(currentSchemaPath), baseName)
		if info, err := os.Stat(cand1); err == nil && !info.IsDir() {
			return filepath.Clean(cand1)
		}
		cand2 := filepath.Join(filepath.Dir(rootSchemaPath), baseName)
		if info, err := os.Stat(cand2); err == nil && !info.IsDir() {
			return filepath.Clean(cand2)
		}
		if indexed, ok := sg.FallbackFileIndex[baseName]; ok {
			return filepath.Clean(indexed)
		}
	}

	return ""
}

func (sg *SchemaGraph) loadSchemaRecursive(absPath string, rootAbsPath string, inheritedTargetNS string) (*Schema, error) {
	absPath = filepath.Clean(absPath)
	if existing, ok := sg.LoadedSchemas[absPath]; ok {
		return existing, nil
	}

	content, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("error reading XSD file %s: %w", absPath, err)
	}

	schema := &Schema{}
	if err := xml.Unmarshal(content, schema); err != nil {
		return nil, fmt.Errorf("error parsing XSD XML %s: %w", absPath, err)
	}

	schema.FilePath = absPath
	schema.Name = filepath.Base(absPath)
	if schema.TargetNamespace == "" && inheritedTargetNS != "" {
		// Chameleon xs:include inherits targetNamespace from including schema
		schema.TargetNamespace = inheritedTargetNS
	}

	if strings.Contains(absPath, "dtc-11-04-05") {
		schema.FactorDuplicates()
		schema.RenameTypeAnonymousComplexType()
	}

	schema.SetParentAndChildren(schema)

	sg.LoadedSchemas[absPath] = schema
	sg.SchemasInOrder = append(sg.SchemasInOrder, schema)

	// Load xs:include schemas
	for _, inc := range schema.Includes {
		resolved := sg.resolveSchemaLocation(absPath, rootAbsPath, schema.TargetNamespace, inc.SchemaLocation)
		if resolved == "" {
			log.Printf("Warning: could not resolve xs:include schemaLocation=%q in %s", inc.SchemaLocation, absPath)
			continue
		}
		incSchema, err := sg.loadSchemaRecursive(resolved, rootAbsPath, schema.TargetNamespace)
		if err != nil {
			return nil, err
		}
		if sg.DirectSchemaDeps[schema] == nil {
			sg.DirectSchemaDeps[schema] = make(map[*Schema]bool)
		}
		sg.DirectSchemaDeps[schema][incSchema] = true
	}

	// Load xs:import schemas
	for _, imp := range schema.Imports {
		resolved := sg.resolveSchemaLocation(absPath, rootAbsPath, imp.Namespace, imp.SchemaLocation)
		if resolved == "" {
			log.Printf("Warning: could not resolve xs:import namespace=%q schemaLocation=%q in %s", imp.Namespace, imp.SchemaLocation, absPath)
			continue
		}
		impSchema, err := sg.loadSchemaRecursive(resolved, rootAbsPath, "")
		if err != nil {
			return nil, err
		}
		if sg.DirectSchemaDeps[schema] == nil {
			sg.DirectSchemaDeps[schema] = make(map[*Schema]bool)
		}
		sg.DirectSchemaDeps[schema][impSchema] = true
	}

	return schema, nil
}

func (sg *SchemaGraph) buildInventoriesAndSymbolTables() {
	for _, schema := range sg.SchemasInOrder {
		targetNS := schema.TargetNamespace

		// 1. Top-level SimpleTypes
		for _, st := range schema.SimpleTypes {
			st.Name = st.NameXSD
			sg.SchemaOfSimpleType[st] = schema
			sg.AllSimpleTypesBySchema[schema] = append(sg.AllSimpleTypesBySchema[schema], st)
			if st.NameXSD != "" {
				qname := QName{Namespace: targetNS, LocalName: st.NameXSD}
				sg.SimpleTypes[qname] = st
				if _, exists := sg.SimpleTypesByLocal[st.NameXSD]; !exists || schema.IsRootSchema {
					sg.SimpleTypesByLocal[st.NameXSD] = st
				}
			}
			sg.collectInlineSimpleTypesFromSimpleType(schema, st)
		}

		// 2. Top-level Attributes
		for _, attr := range schema.Attributes {
			attr.Name = attr.NameXSD
			sg.SchemaOfAttribute[attr] = schema
			sg.AllAttributesBySchema[schema] = append(sg.AllAttributesBySchema[schema], attr)
			if attr.NameXSD != "" {
				qname := QName{Namespace: targetNS, LocalName: attr.NameXSD}
				sg.Attributes[qname] = attr
				if _, exists := sg.AttributesByLocal[attr.NameXSD]; !exists || schema.IsRootSchema {
					sg.AttributesByLocal[attr.NameXSD] = attr
				}
			}
			sg.collectInlineSimpleTypeFromAttribute(schema, attr, attr.NameXSD)
		}

		// 3. Top-level AttributeGroups
		for _, ag := range schema.AttributeGroups {
			if ag.Ref != "" {
				continue
			}
			ag.Name = ag.NameXSD
			sg.SchemaOfAttributeGroup[ag] = schema
			if ag.NameXSD != "" {
				qname := QName{Namespace: targetNS, LocalName: ag.NameXSD}
				sg.AttributeGroups[qname] = ag
				if _, exists := sg.AttributeGroupsByLocal[ag.NameXSD]; !exists || schema.IsRootSchema {
					sg.AttributeGroupsByLocal[ag.NameXSD] = ag
				}
			}
			for _, attr := range ag.Attributes {
				sg.SchemaOfAttribute[attr] = schema
				sg.AllAttributesBySchema[schema] = append(sg.AllAttributesBySchema[schema], attr)
				sg.collectInlineSimpleTypeFromAttribute(schema, attr, ag.NameXSD+"_"+attr.NameXSD)
			}
		}

		// 4. Top-level Groups
		for _, grp := range schema.Groups {
			if grp.Ref != "" {
				continue
			}
			grp.Name = grp.NameXSD
			grp.IsAnonymous = false
			sg.SchemaOfGroup[grp] = schema
			if grp.NameXSD != "" {
				qname := QName{Namespace: targetNS, LocalName: grp.NameXSD}
				sg.Groups[qname] = grp
				if _, exists := sg.GroupsByLocal[grp.NameXSD]; !exists || schema.IsRootSchema {
					sg.GroupsByLocal[grp.NameXSD] = grp
				}
			}
			sg.walkModelGroup(schema, &grp.ModelGroup, grp.NameXSD)
		}

		// 5. Top-level ComplexTypes
		for _, ct := range schema.ComplexTypes {
			ct.Name = ct.NameXSD
			ct.IsAnonymous = false
			sg.SchemaOfComplexType[ct] = schema
			sg.AllComplexTypesBySchema[schema] = append(sg.AllComplexTypesBySchema[schema], ct)
			if ct.NameXSD != "" {
				qname := QName{Namespace: targetNS, LocalName: ct.NameXSD}
				sg.ComplexTypes[qname] = ct
				if _, exists := sg.ComplexTypesByLocal[ct.NameXSD]; !exists || schema.IsRootSchema {
					sg.ComplexTypesByLocal[ct.NameXSD] = ct
				}
			}
			sg.walkComplexType(schema, ct, ct.NameXSD)
		}

		// 6. Top-level Elements
		for _, elem := range schema.Elements {
			elem.Name = elem.NameXSD
			sg.SchemaOfElement[elem] = schema
			sg.AllElementsBySchema[schema] = append(sg.AllElementsBySchema[schema], elem)
			if elem.NameXSD != "" {
				qname := QName{Namespace: targetNS, LocalName: elem.NameXSD}
				sg.Elements[qname] = elem
				if _, exists := sg.ElementsByLocal[elem.NameXSD]; !exists || schema.IsRootSchema {
					sg.ElementsByLocal[elem.NameXSD] = elem
				}
			}
			sg.walkElement(schema, elem, elem.NameXSD, true)
		}
	}

	// Build SubstitutionGroup hierarchy across all schemas
	sg.buildSubstitutionGroups()
}

func (sg *SchemaGraph) collectInlineSimpleTypesFromSimpleType(schema *Schema, st *SimpleType) {
	if st == nil {
		return
	}
	if st.Restriction != nil && st.Restriction.SimpleType != nil {
		inner := st.Restriction.SimpleType
		if inner.NameXSD == "" {
			inner.NameXSD = st.NameXSD + "_Inner"
			inner.Name = inner.NameXSD
		}
		sg.SchemaOfSimpleType[inner] = schema
		sg.AllSimpleTypesBySchema[schema] = append(sg.AllSimpleTypesBySchema[schema], inner)
		sg.collectInlineSimpleTypesFromSimpleType(schema, inner)
	}
	if st.List != nil && st.List.SimpleType != nil {
		inner := st.List.SimpleType
		if inner.NameXSD == "" {
			inner.NameXSD = st.NameXSD + "_Item"
			inner.Name = inner.NameXSD
		}
		sg.SchemaOfSimpleType[inner] = schema
		sg.AllSimpleTypesBySchema[schema] = append(sg.AllSimpleTypesBySchema[schema], inner)
		sg.collectInlineSimpleTypesFromSimpleType(schema, inner)
	}
	if st.Union != nil {
		for idx, inner := range st.Union.SimpleTypes {
			if inner.NameXSD == "" {
				inner.NameXSD = fmt.Sprintf("%s_Member_%d", st.NameXSD, idx+1)
				inner.Name = inner.NameXSD
			}
			sg.SchemaOfSimpleType[inner] = schema
			sg.AllSimpleTypesBySchema[schema] = append(sg.AllSimpleTypesBySchema[schema], inner)
			sg.collectInlineSimpleTypesFromSimpleType(schema, inner)
		}
	}
}

func (sg *SchemaGraph) collectInlineSimpleTypeFromAttribute(schema *Schema, attr *Attribute, prefixName string) {
	if attr == nil || attr.InnerSimpleType == nil {
		return
	}
	st := attr.InnerSimpleType
	if st.NameXSD == "" {
		st.NameXSD = prefixName
	}
	st.Name = st.NameXSD
	if attr.Type == "" {
		attr.Type = st.NameXSD
	}
	sg.SchemaOfSimpleType[st] = schema
	sg.AllSimpleTypesBySchema[schema] = append(sg.AllSimpleTypesBySchema[schema], st)
	qname := QName{Namespace: schema.TargetNamespace, LocalName: st.NameXSD}
	if _, exists := sg.SimpleTypes[qname]; !exists {
		sg.SimpleTypes[qname] = st
	}
	if _, exists := sg.SimpleTypesByLocal[st.NameXSD]; !exists {
		sg.SimpleTypesByLocal[st.NameXSD] = st
	}
	sg.collectInlineSimpleTypesFromSimpleType(schema, st)
}

func (sg *SchemaGraph) walkComplexType(schema *Schema, ct *ComplexType, prefixName string) {
	if ct == nil {
		return
	}
	for _, attr := range ct.Attributes {
		sg.SchemaOfAttribute[attr] = schema
		sg.AllAttributesBySchema[schema] = append(sg.AllAttributesBySchema[schema], attr)
		sg.collectInlineSimpleTypeFromAttribute(schema, attr, prefixName+"_"+attr.NameXSD)
	}
	if ct.Extension != nil {
		for _, attr := range ct.Extension.Attributes {
			sg.SchemaOfAttribute[attr] = schema
			sg.AllAttributesBySchema[schema] = append(sg.AllAttributesBySchema[schema], attr)
			sg.collectInlineSimpleTypeFromAttribute(schema, attr, prefixName+"_"+attr.NameXSD)
		}
		sg.walkModelGroup(schema, &ct.Extension.ModelGroup, prefixName)
	}
	if ct.ComplexContent != nil {
		if ct.ComplexContent.Extension != nil {
			for _, attr := range ct.ComplexContent.Extension.Attributes {
				sg.SchemaOfAttribute[attr] = schema
				sg.AllAttributesBySchema[schema] = append(sg.AllAttributesBySchema[schema], attr)
				sg.collectInlineSimpleTypeFromAttribute(schema, attr, prefixName+"_"+attr.NameXSD)
			}
			sg.walkModelGroup(schema, &ct.ComplexContent.Extension.ModelGroup, prefixName)
		}
		if ct.ComplexContent.Restriction != nil {
			for _, attr := range ct.ComplexContent.Restriction.Attributes {
				sg.SchemaOfAttribute[attr] = schema
				sg.AllAttributesBySchema[schema] = append(sg.AllAttributesBySchema[schema], attr)
				sg.collectInlineSimpleTypeFromAttribute(schema, attr, prefixName+"_"+attr.NameXSD)
			}
			sg.walkModelGroup(schema, &ct.ComplexContent.Restriction.ModelGroup, prefixName)
		}
	}
	if ct.SimpleContent != nil {
		if ct.SimpleContent.Extension != nil {
			for _, attr := range ct.SimpleContent.Extension.Attributes {
				sg.SchemaOfAttribute[attr] = schema
				sg.AllAttributesBySchema[schema] = append(sg.AllAttributesBySchema[schema], attr)
				sg.collectInlineSimpleTypeFromAttribute(schema, attr, prefixName+"_"+attr.NameXSD)
			}
		}
		if ct.SimpleContent.Restriction != nil {
			for _, attr := range ct.SimpleContent.Restriction.Attributes {
				sg.SchemaOfAttribute[attr] = schema
				sg.AllAttributesBySchema[schema] = append(sg.AllAttributesBySchema[schema], attr)
				sg.collectInlineSimpleTypeFromAttribute(schema, attr, prefixName+"_"+attr.NameXSD)
			}
			if ct.SimpleContent.Restriction.SimpleType != nil {
				st := ct.SimpleContent.Restriction.SimpleType
				if st.NameXSD == "" {
					st.NameXSD = prefixName + "_Content"
				}
				st.Name = st.NameXSD
				sg.SchemaOfSimpleType[st] = schema
				sg.AllSimpleTypesBySchema[schema] = append(sg.AllSimpleTypesBySchema[schema], st)
			}
		}
	}
	sg.walkModelGroup(schema, &ct.ModelGroup, prefixName)
}

func (sg *SchemaGraph) walkModelGroup(schema *Schema, mg *ModelGroup, prefixName string) {
	if mg == nil {
		return
	}
	for _, elem := range mg.Elements {
		sg.SchemaOfElement[elem] = schema
		sg.AllElementsBySchema[schema] = append(sg.AllElementsBySchema[schema], elem)
		sg.walkElement(schema, elem, prefixName+"_"+elem.NameXSD, false)
	}
	for _, seq := range mg.Sequences {
		for _, elem := range seq.Elements {
			sg.SchemaOfElement[elem] = schema
			sg.AllElementsBySchema[schema] = append(sg.AllElementsBySchema[schema], elem)
			sg.walkElement(schema, elem, prefixName+"_"+elem.NameXSD, false)
		}
		sg.walkModelGroup(schema, &seq.ModelGroup, prefixName)
	}
	for _, ch := range mg.Choices {
		for _, elem := range ch.Elements {
			sg.SchemaOfElement[elem] = schema
			sg.AllElementsBySchema[schema] = append(sg.AllElementsBySchema[schema], elem)
			sg.walkElement(schema, elem, prefixName+"_"+elem.NameXSD, false)
		}
		sg.walkModelGroup(schema, &ch.ModelGroup, prefixName)
	}
	for _, all := range mg.Alls {
		for _, elem := range all.Elements {
			sg.SchemaOfElement[elem] = schema
			sg.AllElementsBySchema[schema] = append(sg.AllElementsBySchema[schema], elem)
			sg.walkElement(schema, elem, prefixName+"_"+elem.NameXSD, false)
		}
		sg.walkModelGroup(schema, &all.ModelGroup, prefixName)
	}
	for _, grp := range mg.Groups {
		sg.walkModelGroup(schema, &grp.ModelGroup, prefixName)
	}
}

func (sg *SchemaGraph) walkElement(schema *Schema, elem *Element, prefixName string, isTopLevel bool) {
	if elem == nil {
		return
	}
	if elem.SimpleType != nil {
		st := elem.SimpleType
		if st.NameXSD == "" {
			st.NameXSD = prefixName + "_SimpleType"
		}
		st.Name = st.NameXSD
		if elem.Type == "" {
			elem.Type = st.NameXSD
		}
		sg.SchemaOfSimpleType[st] = schema
		sg.AllSimpleTypesBySchema[schema] = append(sg.AllSimpleTypesBySchema[schema], st)
		qname := QName{Namespace: schema.TargetNamespace, LocalName: st.NameXSD}
		if _, exists := sg.SimpleTypes[qname]; !exists {
			sg.SimpleTypes[qname] = st
		}
		if _, exists := sg.SimpleTypesByLocal[st.NameXSD]; !exists {
			sg.SimpleTypesByLocal[st.NameXSD] = st
		}
		sg.collectInlineSimpleTypesFromSimpleType(schema, st)
	}
	if elem.ComplexType != nil {
		ct := elem.ComplexType
		ct.IsAnonymous = true
		ct.OuterElement = elem
		if ct.Name == "" {
			ct.Name = "A_" + elem.NameXSD
		}
		sg.SchemaOfComplexType[ct] = schema
		if !isTopLevel {
			// Top-level elements with inline ComplexType are generated via the TopElements loop
			if !slices.Contains(sg.AllComplexTypesBySchema[schema], ct) {
				sg.AllComplexTypesBySchema[schema] = append(sg.AllComplexTypesBySchema[schema], ct)
			}
		}
		sg.walkComplexType(schema, ct, ct.Name)
	}
}

func (sg *SchemaGraph) buildSubstitutionGroups() {
	// Direct substitutionGroup map: head QName -> direct member elements
	directMembers := make(map[QName][]*Element)
	for _, schema := range sg.SchemasInOrder {
		for _, elem := range schema.Elements {
			if elem.SubstitutionGroup == "" {
				continue
			}
			headQName, isBuiltIn := sg.ResolveQName(schema, elem.SubstitutionGroup)
			if isBuiltIn || headQName.LocalName == "" {
				continue
			}
			// If head wasn't found by exact namespace, try local name fallback
			if _, ok := sg.Elements[headQName]; !ok {
				if headElem, okLocal := sg.ElementsByLocal[headQName.LocalName]; okLocal {
					if headSchema, okS := sg.SchemaOfElement[headElem]; okS {
						headQName = QName{Namespace: headSchema.TargetNamespace, LocalName: headElem.NameXSD}
					}
				}
			}
			directMembers[headQName] = append(directMembers[headQName], elem)
		}
	}

	// Compute transitive closure for each head QName
	for headQName := range directMembers {
		visited := make(map[*Element]bool)
		var collect func(q QName)
		collect = func(q QName) {
			for _, member := range directMembers[q] {
				if visited[member] {
					continue
				}
				visited[member] = true
				if member.Abstract != "true" {
					sg.SubstitutionGroupResult[headQName] = append(sg.SubstitutionGroupResult[headQName], member)
				}
				memberSchema := sg.SchemaOfElement[member]
				if memberSchema != nil {
					collect(QName{Namespace: memberSchema.TargetNamespace, LocalName: member.NameXSD})
				}
			}
		}
		collect(headQName)
	}
}

func (sg *SchemaGraph) pruneUnreachableConstructs() {
	if sg.RootSchema == nil || len(sg.SchemasInOrder) <= 1 {
		return
	}
	if len(sg.RootSchema.Elements) == 0 && len(sg.AllComplexTypesBySchema[sg.RootSchema]) == 0 {
		return
	}

	reachableCT := make(map[*ComplexType]bool)
	reachableST := make(map[*SimpleType]bool)
	reachableGrp := make(map[*Group]bool)
	reachableAG := make(map[*AttributeGroup]bool)
	reachableEl := make(map[*Element]bool)
	reachableAttr := make(map[*Attribute]bool)

	var markElement func(contextSchema *Schema, el *Element)
	var markComplexType func(ct *ComplexType)
	var markSimpleType func(st *SimpleType)
	var markGroup func(contextSchema *Schema, grp *Group)
	var markAttributeGroup func(contextSchema *Schema, ag *AttributeGroup)
	var markAttribute func(contextSchema *Schema, attr *Attribute)
	var markModelGroup func(contextSchema *Schema, mg *ModelGroup)

	markBaseType := func(contextSchema *Schema, rawType string) {
		if rawType == "" {
			return
		}
		if ct := sg.ResolveComplexType(contextSchema, rawType); ct != nil {
			markComplexType(ct)
		}
		if st := sg.ResolveSimpleType(contextSchema, rawType); st != nil {
			markSimpleType(st)
		}
	}

	markElement = func(contextSchema *Schema, el *Element) {
		if el == nil || reachableEl[el] {
			return
		}
		reachableEl[el] = true
		elSchema := sg.SchemaOfElement[el]
		if elSchema == nil {
			elSchema = contextSchema
		}
		if el.Ref != "" {
			if targetEl := sg.ResolveElement(elSchema, el.Ref); targetEl != nil {
				targetSchema := sg.SchemaOfElement[targetEl]
				markElement(targetSchema, targetEl)
				if targetSchema != nil && targetEl.NameXSD != "" {
					headQName := QName{Namespace: targetSchema.TargetNamespace, LocalName: targetEl.NameXSD}
					for _, sub := range sg.EligibleSubstitutionMembers(elSchema, headQName) {
						markElement(sg.SchemaOfElement[sub], sub)
					}
				}
			}
		}
		if el.Type != "" {
			markBaseType(elSchema, el.Type)
		}
		if el.ComplexType != nil {
			markComplexType(el.ComplexType)
		}
		if el.SimpleType != nil {
			markSimpleType(el.SimpleType)
		}
	}

	markAttribute = func(contextSchema *Schema, attr *Attribute) {
		if attr == nil || reachableAttr[attr] {
			return
		}
		reachableAttr[attr] = true
		attrSchema := sg.SchemaOfAttribute[attr]
		if attrSchema == nil {
			attrSchema = contextSchema
		}
		if attr.Ref != "" {
			if targetAttr := sg.ResolveAttribute(attrSchema, attr.Ref); targetAttr != nil {
				markAttribute(sg.SchemaOfAttribute[targetAttr], targetAttr)
			}
		}
		if attr.Type != "" {
			markBaseType(attrSchema, attr.Type)
		}
		if attr.InnerSimpleType != nil {
			markSimpleType(attr.InnerSimpleType)
		}
	}

	markAttributeGroup = func(contextSchema *Schema, ag *AttributeGroup) {
		if ag == nil || reachableAG[ag] {
			return
		}
		reachableAG[ag] = true
		agSchema := sg.SchemaOfAttributeGroup[ag]
		if agSchema == nil {
			agSchema = contextSchema
		}
		if ag.Ref != "" {
			if targetAg := sg.ResolveAttributeGroup(agSchema, ag.Ref); targetAg != nil {
				markAttributeGroup(sg.SchemaOfAttributeGroup[targetAg], targetAg)
			}
		}
		for _, attr := range ag.Attributes {
			markAttribute(agSchema, attr)
		}
		for _, subAg := range ag.AttributeGroups {
			markAttributeGroup(agSchema, subAg)
		}
	}

	markGroup = func(contextSchema *Schema, grp *Group) {
		if grp == nil || reachableGrp[grp] {
			return
		}
		reachableGrp[grp] = true
		grpSchema := sg.SchemaOfGroup[grp]
		if grpSchema == nil {
			grpSchema = contextSchema
		}
		if grp.Ref != "" {
			if targetGrp := sg.ResolveGroup(grpSchema, grp.Ref); targetGrp != nil {
				markGroup(sg.SchemaOfGroup[targetGrp], targetGrp)
			}
		}
		markModelGroup(grpSchema, &grp.ModelGroup)
	}

	markModelGroup = func(contextSchema *Schema, mg *ModelGroup) {
		if mg == nil {
			return
		}
		for _, el := range mg.Elements {
			markElement(contextSchema, el)
		}
		for _, seq := range mg.Sequences {
			for _, el := range seq.Elements {
				markElement(contextSchema, el)
			}
			markModelGroup(contextSchema, &seq.ModelGroup)
		}
		for _, ch := range mg.Choices {
			for _, el := range ch.Elements {
				markElement(contextSchema, el)
			}
			markModelGroup(contextSchema, &ch.ModelGroup)
		}
		for _, all := range mg.Alls {
			for _, el := range all.Elements {
				markElement(contextSchema, el)
			}
			markModelGroup(contextSchema, &all.ModelGroup)
		}
		for _, grp := range mg.Groups {
			markGroup(contextSchema, grp)
		}
	}

	markSimpleType = func(st *SimpleType) {
		if st == nil || reachableST[st] {
			return
		}
		reachableST[st] = true
		stSchema := sg.SchemaOfSimpleType[st]
		if st.Restriction != nil {
			if st.Restriction.Base != "" {
				markBaseType(stSchema, st.Restriction.Base)
			}
			if st.Restriction.SimpleType != nil {
				markSimpleType(st.Restriction.SimpleType)
			}
		}
		if st.List != nil {
			if st.List.ItemType != "" {
				markBaseType(stSchema, st.List.ItemType)
			}
			if st.List.SimpleType != nil {
				markSimpleType(st.List.SimpleType)
			}
		}
		if st.Union != nil {
			for _, mt := range strings.Fields(st.Union.MemberTypes) {
				markBaseType(stSchema, mt)
			}
			for _, subSt := range st.Union.SimpleTypes {
				markSimpleType(subSt)
			}
		}
	}

	markComplexType = func(ct *ComplexType) {
		if ct == nil || reachableCT[ct] {
			return
		}
		reachableCT[ct] = true
		ctSchema := sg.SchemaOfComplexType[ct]
		if ct.OuterElement != nil {
			markElement(ctSchema, ct.OuterElement)
		}
		for _, attr := range ct.Attributes {
			markAttribute(ctSchema, attr)
		}
		for _, ag := range ct.AttributeGroups {
			markAttributeGroup(ctSchema, ag)
		}
		if ct.Extension != nil {
			for _, attr := range ct.Extension.Attributes {
				markAttribute(ctSchema, attr)
			}
			for _, ag := range ct.Extension.AttributeGroups {
				markAttributeGroup(ctSchema, ag)
			}
			if ct.Extension.Base != "" {
				markBaseType(ctSchema, ct.Extension.Base)
			}
			markModelGroup(ctSchema, &ct.Extension.ModelGroup)
		}
		if ct.ComplexContent != nil {
			if ct.ComplexContent.Extension != nil {
				for _, attr := range ct.ComplexContent.Extension.Attributes {
					markAttribute(ctSchema, attr)
				}
				for _, ag := range ct.ComplexContent.Extension.AttributeGroups {
					markAttributeGroup(ctSchema, ag)
				}
				if ct.ComplexContent.Extension.Base != "" {
					markBaseType(ctSchema, ct.ComplexContent.Extension.Base)
				}
				markModelGroup(ctSchema, &ct.ComplexContent.Extension.ModelGroup)
			}
			if ct.ComplexContent.Restriction != nil {
				for _, attr := range ct.ComplexContent.Restriction.Attributes {
					markAttribute(ctSchema, attr)
				}
				for _, ag := range ct.ComplexContent.Restriction.AttributeGroups {
					markAttributeGroup(ctSchema, ag)
				}
				if ct.ComplexContent.Restriction.Base != "" {
					markBaseType(ctSchema, ct.ComplexContent.Restriction.Base)
				}
				markModelGroup(ctSchema, &ct.ComplexContent.Restriction.ModelGroup)
			}
		}
		if ct.SimpleContent != nil {
			if ct.SimpleContent.Extension != nil {
				for _, attr := range ct.SimpleContent.Extension.Attributes {
					markAttribute(ctSchema, attr)
				}
				for _, ag := range ct.SimpleContent.Extension.AttributeGroups {
					markAttributeGroup(ctSchema, ag)
				}
				if ct.SimpleContent.Extension.Base != "" {
					markBaseType(ctSchema, ct.SimpleContent.Extension.Base)
				}
			}
			if ct.SimpleContent.Restriction != nil {
				for _, attr := range ct.SimpleContent.Restriction.Attributes {
					markAttribute(ctSchema, attr)
				}
				for _, ag := range ct.SimpleContent.Restriction.AttributeGroups {
					markAttributeGroup(ctSchema, ag)
				}
				if ct.SimpleContent.Restriction.Base != "" {
					markBaseType(ctSchema, ct.SimpleContent.Restriction.Base)
				}
				if ct.SimpleContent.Restriction.SimpleType != nil {
					markSimpleType(ct.SimpleContent.Restriction.SimpleType)
				}
			}
		}
		markModelGroup(ctSchema, &ct.ModelGroup)
	}

	// Seed reachability from RootSchema constructs
	for _, el := range sg.RootSchema.Elements {
		markElement(sg.RootSchema, el)
	}
	for _, ct := range sg.AllComplexTypesBySchema[sg.RootSchema] {
		markComplexType(ct)
	}
	for _, st := range sg.AllSimpleTypesBySchema[sg.RootSchema] {
		markSimpleType(st)
	}
	for _, grp := range sg.RootSchema.Groups {
		markGroup(sg.RootSchema, grp)
	}
	for _, ag := range sg.RootSchema.AttributeGroups {
		markAttributeGroup(sg.RootSchema, ag)
	}
	for _, attr := range sg.RootSchema.Attributes {
		markAttribute(sg.RootSchema, attr)
	}

	// Filter non-root schemas to keep only reachable constructs
	for _, s := range sg.SchemasInOrder {
		if s.IsRootSchema {
			continue
		}
		var keptCT []*ComplexType
		for _, ct := range sg.AllComplexTypesBySchema[s] {
			if reachableCT[ct] {
				keptCT = append(keptCT, ct)
			}
		}
		sg.AllComplexTypesBySchema[s] = keptCT

		var keptST []*SimpleType
		for _, st := range sg.AllSimpleTypesBySchema[s] {
			if reachableST[st] {
				keptST = append(keptST, st)
			}
		}
		sg.AllSimpleTypesBySchema[s] = keptST

		var keptGrp []*Group
		for _, grp := range s.Groups {
			if reachableGrp[grp] {
				keptGrp = append(keptGrp, grp)
			}
		}
		s.Groups = keptGrp

		var keptAG []*AttributeGroup
		for _, ag := range s.AttributeGroups {
			if reachableAG[ag] {
				keptAG = append(keptAG, ag)
			}
		}
		s.AttributeGroups = keptAG

		var keptEl []*Element
		for _, el := range s.Elements {
			if reachableEl[el] {
				keptEl = append(keptEl, el)
			}
		}
		s.Elements = keptEl

		var keptAttr []*Attribute
		for _, attr := range s.Attributes {
			if reachableAttr[attr] {
				keptAttr = append(keptAttr, attr)
			}
		}
		s.Attributes = keptAttr
	}

	// Filter SubstitutionGroupResult to keep only reachable elements
	for headQName, members := range sg.SubstitutionGroupResult {
		var keptMembers []*Element
		for _, m := range members {
			if reachableEl[m] {
				keptMembers = append(keptMembers, m)
			}
		}
		sg.SubstitutionGroupResult[headQName] = keptMembers
	}
}

// IsBuiltInXSDType returns true if localName is a standard W3C XML Schema built-in primitive or derived type.
func IsBuiltInXSDType(localName string) bool {
	switch localName {
	case "anyType", "anySimpleType",
		"string", "normalizedString", "token", "language", "Name", "NCName",
		"NMTOKEN", "NMTOKENS", "ID", "IDREF", "IDREFS", "ENTITY", "ENTITIES",
		"QName", "anyURI", "NOTATION",
		"decimal", "integer", "nonPositiveInteger", "negativeInteger", "long",
		"int", "short", "byte", "nonNegativeInteger", "unsignedLong", "unsignedInt",
		"unsignedShort", "unsignedByte", "positiveInteger",
		"double", "float", "boolean",
		"dateTime", "date", "time", "duration", "gYearMonth", "gYear", "gMonthDay", "gDay", "gMonth",
		"hexBinary", "base64Binary":
		return true
	}
	return false
}

// ResolveQName resolves a raw XML QName string (e.g. "xlink:actuate", "xs:string", "bookType")
// within the namespace context of sourceSchema.
func (sg *SchemaGraph) ResolveQName(sourceSchema *Schema, rawRef string) (qname QName, isXSDBuiltIn bool) {
	rawRef = strings.TrimSpace(rawRef)
	if rawRef == "" {
		return QName{}, false
	}

	prefix := NsPrefix(rawRef)
	local := Name(rawRef)

	if prefix != "" {
		if prefix == "xs" || prefix == "xsd" {
			return QName{Namespace: "http://www.w3.org/2001/XMLSchema", LocalName: local}, true
		}
		if sourceSchema != nil {
			for _, decl := range sourceSchema.Xmlns {
				if decl.Prefix == prefix {
					if decl.Uri == "http://www.w3.org/2001/XMLSchema" {
						return QName{Namespace: decl.Uri, LocalName: local}, true
					}
					return QName{Namespace: decl.Uri, LocalName: local}, false
				}
			}
		}
		switch prefix {
		case "xml":
			return QName{Namespace: "http://www.w3.org/XML/1998/namespace", LocalName: local}, false
		case "xlink":
			return QName{Namespace: "http://www.w3.org/1999/xlink", LocalName: local}, false
		}
		return QName{Namespace: prefix, LocalName: local}, false
	}

	// Unprefixed reference
	var defaultNS string
	if sourceSchema != nil {
		for _, decl := range sourceSchema.Xmlns {
			if decl.Prefix == "" {
				defaultNS = decl.Uri
				break
			}
		}
	}

	targetNS := ""
	if sourceSchema != nil {
		targetNS = sourceSchema.TargetNamespace
	}

	if defaultNS == "http://www.w3.org/2001/XMLSchema" {
		if IsBuiltInXSDType(local) {
			// Check if targetNS explicitly overrides it with a local symbol; if not, it's built-in XSD
			qTarget := QName{Namespace: targetNS, LocalName: local}
			if _, hasST := sg.SimpleTypes[qTarget]; !hasST {
				if _, hasCT := sg.ComplexTypes[qTarget]; !hasCT {
					return QName{Namespace: defaultNS, LocalName: local}, true
				}
			}
		}
		return QName{Namespace: targetNS, LocalName: local}, false
	}

	if defaultNS != "" {
		return QName{Namespace: defaultNS, LocalName: local}, false
	}

	if IsBuiltInXSDType(local) {
		qTarget := QName{Namespace: targetNS, LocalName: local}
		if _, hasST := sg.SimpleTypes[qTarget]; !hasST {
			if _, hasCT := sg.ComplexTypes[qTarget]; !hasCT {
				return QName{Namespace: "http://www.w3.org/2001/XMLSchema", LocalName: local}, true
			}
		}
	}

	return QName{Namespace: targetNS, LocalName: local}, false
}

// ResolveComplexType looks up a ComplexType referenced by rawRef in sourceSchema.
func (sg *SchemaGraph) ResolveComplexType(sourceSchema *Schema, rawRef string) *ComplexType {
	qname, isBuiltIn := sg.ResolveQName(sourceSchema, rawRef)
	if isBuiltIn || qname.LocalName == "" {
		return nil
	}
	if ct, ok := sg.ComplexTypes[qname]; ok {
		return ct
	}
	if sourceSchema != nil {
		if ct, ok := sg.ComplexTypes[QName{Namespace: sourceSchema.TargetNamespace, LocalName: qname.LocalName}]; ok {
			return ct
		}
	}
	if ct, ok := sg.ComplexTypesByLocal[qname.LocalName]; ok {
		return ct
	}
	return nil
}

// ResolveSimpleType looks up a SimpleType referenced by rawRef in sourceSchema.
func (sg *SchemaGraph) ResolveSimpleType(sourceSchema *Schema, rawRef string) *SimpleType {
	qname, isBuiltIn := sg.ResolveQName(sourceSchema, rawRef)
	if isBuiltIn || qname.LocalName == "" {
		return nil
	}
	if st, ok := sg.SimpleTypes[qname]; ok {
		return st
	}
	if sourceSchema != nil {
		if st, ok := sg.SimpleTypes[QName{Namespace: sourceSchema.TargetNamespace, LocalName: qname.LocalName}]; ok {
			return st
		}
	}
	if st, ok := sg.SimpleTypesByLocal[qname.LocalName]; ok {
		return st
	}
	return nil
}

// ResolveElement looks up a top-level Element referenced by rawRef in sourceSchema.
func (sg *SchemaGraph) ResolveElement(sourceSchema *Schema, rawRef string) *Element {
	qname, isBuiltIn := sg.ResolveQName(sourceSchema, rawRef)
	if isBuiltIn || qname.LocalName == "" {
		return nil
	}
	if el, ok := sg.Elements[qname]; ok {
		return el
	}
	if sourceSchema != nil {
		if el, ok := sg.Elements[QName{Namespace: sourceSchema.TargetNamespace, LocalName: qname.LocalName}]; ok {
			return el
		}
	}
	if el, ok := sg.ElementsByLocal[qname.LocalName]; ok {
		return el
	}
	return nil
}

// ResolveAttribute looks up a top-level Attribute referenced by rawRef in sourceSchema.
func (sg *SchemaGraph) ResolveAttribute(sourceSchema *Schema, rawRef string) *Attribute {
	qname, isBuiltIn := sg.ResolveQName(sourceSchema, rawRef)
	if isBuiltIn || qname.LocalName == "" {
		return nil
	}
	if attr, ok := sg.Attributes[qname]; ok {
		return attr
	}
	if sourceSchema != nil {
		if attr, ok := sg.Attributes[QName{Namespace: sourceSchema.TargetNamespace, LocalName: qname.LocalName}]; ok {
			return attr
		}
	}
	if attr, ok := sg.AttributesByLocal[qname.LocalName]; ok {
		return attr
	}
	return nil
}

// ResolveGroup looks up a top-level Group referenced by rawRef in sourceSchema.
func (sg *SchemaGraph) ResolveGroup(sourceSchema *Schema, rawRef string) *Group {
	qname, isBuiltIn := sg.ResolveQName(sourceSchema, rawRef)
	if isBuiltIn || qname.LocalName == "" {
		return nil
	}
	if grp, ok := sg.Groups[qname]; ok {
		return grp
	}
	if sourceSchema != nil {
		if grp, ok := sg.Groups[QName{Namespace: sourceSchema.TargetNamespace, LocalName: qname.LocalName}]; ok {
			return grp
		}
	}
	if grp, ok := sg.GroupsByLocal[qname.LocalName]; ok {
		return grp
	}
	return nil
}

// ResolveAttributeGroup looks up a top-level AttributeGroup referenced by rawRef in sourceSchema.
func (sg *SchemaGraph) ResolveAttributeGroup(sourceSchema *Schema, rawRef string) *AttributeGroup {
	qname, isBuiltIn := sg.ResolveQName(sourceSchema, rawRef)
	if isBuiltIn || qname.LocalName == "" {
		return nil
	}
	if ag, ok := sg.AttributeGroups[qname]; ok {
		return ag
	}
	if sourceSchema != nil {
		if ag, ok := sg.AttributeGroups[QName{Namespace: sourceSchema.TargetNamespace, LocalName: qname.LocalName}]; ok {
			return ag
		}
	}
	if ag, ok := sg.AttributeGroupsByLocal[qname.LocalName]; ok {
		return ag
	}
	return nil
}

func (sg *SchemaGraph) isSchemaEmpty(schema *Schema) bool {
	return len(sg.AllComplexTypesBySchema[schema]) == 0 &&
		len(sg.AllSimpleTypesBySchema[schema]) == 0 &&
		len(schema.Groups) == 0 &&
		len(schema.AttributeGroups) == 0 &&
		len(schema.Elements) == 0 &&
		len(schema.Attributes) == 0
}

// CanSchemaReach returns true if target is reachable from start via xs:include/xs:import
// (including schemas sharing the same non-empty TargetNamespace as any reachable schema).
func (sg *SchemaGraph) CanSchemaReach(start, target *Schema) bool {
	if start == nil || target == nil {
		return false
	}
	if start == target {
		return true
	}
	visited := make(map[*Schema]bool)
	queue := []*Schema{start}
	visited[start] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr == target {
			return true
		}
		for next := range sg.DirectSchemaDeps[curr] {
			if !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
		}
		if curr.TargetNamespace != "" {
			for _, sameNS := range sg.SchemasInOrder {
				if sameNS.TargetNamespace == curr.TargetNamespace && !visited[sameNS] {
					visited[sameNS] = true
					queue = append(queue, sameNS)
				}
			}
		}
	}
	return false
}

// EligibleSubstitutionMembers returns the substitutionGroup members of headQName that are
// reachable from fromSchema (bounded at <= 100 members).
func (sg *SchemaGraph) EligibleSubstitutionMembers(fromSchema *Schema, headQName QName) []*Element {
	allMembers := sg.SubstitutionGroupResult[headQName]
	if len(allMembers) == 0 {
		return nil
	}
	var eligible []*Element
	for _, m := range allMembers {
		mSchema := sg.SchemaOfElement[m]
		if sg.CanSchemaReach(fromSchema, mSchema) {
			eligible = append(eligible, m)
		}
	}
	if len(eligible) == 0 || len(eligible) > 200 {
		return nil
	}
	return eligible
}

func (sg *SchemaGraph) computeSchemaSymbolDependencies() map[*Schema]map[*Schema]bool {
	deps := make(map[*Schema]map[*Schema]bool)
	addDep := func(from, to *Schema) {
		if from == nil || to == nil || from == to {
			return
		}
		if sg.isSchemaEmpty(to) && !to.IsRootSchema {
			return
		}
		if deps[from] == nil {
			deps[from] = make(map[*Schema]bool)
		}
		deps[from][to] = true
	}

	checkComplexTypeRef := func(ownerSchema *Schema, contextSchema *Schema, rawType string) {
		if rawType == "" {
			return
		}
		if ct := sg.ResolveComplexType(contextSchema, rawType); ct != nil {
			addDep(ownerSchema, sg.SchemaOfComplexType[ct])
		}
	}

	var checkAttrRefs func(ownerSchema *Schema, contextSchema *Schema, attrs []*Attribute, ags []*AttributeGroup, visitedAG map[*AttributeGroup]bool)
	checkAttrRefs = func(ownerSchema *Schema, contextSchema *Schema, attrs []*Attribute, ags []*AttributeGroup, visitedAG map[*AttributeGroup]bool) {
		for _, attr := range attrs {
			if attr.Ref != "" {
				if targetAttr := sg.ResolveAttribute(contextSchema, attr.Ref); targetAttr != nil {
					addDep(ownerSchema, sg.SchemaOfAttribute[targetAttr])
				}
			}
		}
		for _, ag := range ags {
			if ag.Ref != "" {
				if targetAg := sg.ResolveAttributeGroup(contextSchema, ag.Ref); targetAg != nil && !visitedAG[targetAg] {
					visitedAG[targetAg] = true
					agSchema := sg.SchemaOfAttributeGroup[targetAg]
					addDep(ownerSchema, agSchema)
					checkAttrRefs(ownerSchema, agSchema, targetAg.Attributes, targetAg.AttributeGroups, visitedAG)
				}
			}
		}
	}

	var collectMG func(ownerSchema *Schema, contextSchema *Schema, mg *ModelGroup, visitedGroups map[*Group]bool)
	collectMG = func(ownerSchema *Schema, contextSchema *Schema, mg *ModelGroup, visitedGroups map[*Group]bool) {
		if mg == nil {
			return
		}
		checkElem := func(el *Element) {
			if el == nil {
				return
			}
			if el.Ref != "" {
				if targetEl := sg.ResolveElement(contextSchema, el.Ref); targetEl != nil {
					targetSchema := sg.SchemaOfElement[targetEl]
					headQName := QName{Namespace: targetSchema.TargetNamespace, LocalName: targetEl.NameXSD}
					subMembers := sg.EligibleSubstitutionMembers(contextSchema, headQName)
					if len(subMembers) > 0 {
						if targetEl.Abstract != "true" {
							checkComplexTypeRef(ownerSchema, targetSchema, targetEl.Type)
							if targetEl.ComplexType != nil {
								addDep(ownerSchema, sg.SchemaOfComplexType[targetEl.ComplexType])
							}
						}
						for _, subMember := range subMembers {
							subSchema := sg.SchemaOfElement[subMember]
							checkComplexTypeRef(ownerSchema, subSchema, subMember.Type)
							if subMember.ComplexType != nil {
								addDep(ownerSchema, sg.SchemaOfComplexType[subMember.ComplexType])
							}
						}
						return
					}
					checkComplexTypeRef(ownerSchema, targetSchema, targetEl.Type)
					if targetEl.ComplexType != nil {
						addDep(ownerSchema, sg.SchemaOfComplexType[targetEl.ComplexType])
					}
				}
				return
			}
			checkComplexTypeRef(ownerSchema, contextSchema, el.Type)
			if el.ComplexType != nil {
				addDep(ownerSchema, sg.SchemaOfComplexType[el.ComplexType])
			}
		}
		for _, el := range mg.Elements {
			checkElem(el)
		}
		for _, seq := range mg.Sequences {
			for _, el := range seq.Elements {
				checkElem(el)
			}
			collectMG(ownerSchema, contextSchema, &seq.ModelGroup, visitedGroups)
		}
		for _, ch := range mg.Choices {
			for _, el := range ch.Elements {
				checkElem(el)
			}
			collectMG(ownerSchema, contextSchema, &ch.ModelGroup, visitedGroups)
		}
		for _, all := range mg.Alls {
			for _, el := range all.Elements {
				checkElem(el)
			}
			collectMG(ownerSchema, contextSchema, &all.ModelGroup, visitedGroups)
		}
		for _, grp := range mg.Groups {
			if grp.Ref != "" {
				if targetGrp := sg.ResolveGroup(contextSchema, grp.Ref); targetGrp != nil && !visitedGroups[targetGrp] {
					visitedGroups[targetGrp] = true
					grpSchema := sg.SchemaOfGroup[targetGrp]
					collectMG(ownerSchema, grpSchema, &targetGrp.ModelGroup, visitedGroups)
				}
			}
		}
	}

	var collectCTChain func(ownerSchema *Schema, ct *ComplexType, visitedCT map[*ComplexType]bool)
	collectCTChain = func(ownerSchema *Schema, ct *ComplexType, visitedCT map[*ComplexType]bool) {
		if ct == nil || visitedCT[ct] {
			return
		}
		visitedCT[ct] = true
		ctSchema := sg.SchemaOfComplexType[ct]
		visitedGroups := make(map[*Group]bool)
		visitedAG := make(map[*AttributeGroup]bool)
		checkAttrRefs(ownerSchema, ctSchema, ct.Attributes, ct.AttributeGroups, visitedAG)

		baseName := ""
		if ct.Extension != nil {
			checkAttrRefs(ownerSchema, ctSchema, ct.Extension.Attributes, ct.Extension.AttributeGroups, visitedAG)
			if ct.Extension.Base != "" {
				baseName = ct.Extension.Base
			}
			collectMG(ownerSchema, ctSchema, &ct.Extension.ModelGroup, visitedGroups)
		}
		if ct.ComplexContent != nil {
			if ct.ComplexContent.Extension != nil {
				checkAttrRefs(ownerSchema, ctSchema, ct.ComplexContent.Extension.Attributes, ct.ComplexContent.Extension.AttributeGroups, visitedAG)
				if ct.ComplexContent.Extension.Base != "" {
					baseName = ct.ComplexContent.Extension.Base
				}
				collectMG(ownerSchema, ctSchema, &ct.ComplexContent.Extension.ModelGroup, visitedGroups)
			}
			if ct.ComplexContent.Restriction != nil {
				checkAttrRefs(ownerSchema, ctSchema, ct.ComplexContent.Restriction.Attributes, ct.ComplexContent.Restriction.AttributeGroups, visitedAG)
				if ct.ComplexContent.Restriction.Base != "" {
					baseName = ct.ComplexContent.Restriction.Base
				}
				collectMG(ownerSchema, ctSchema, &ct.ComplexContent.Restriction.ModelGroup, visitedGroups)
			}
		}
		if ct.SimpleContent != nil {
			if ct.SimpleContent.Extension != nil {
				checkAttrRefs(ownerSchema, ctSchema, ct.SimpleContent.Extension.Attributes, ct.SimpleContent.Extension.AttributeGroups, visitedAG)
			}
			if ct.SimpleContent.Restriction != nil {
				checkAttrRefs(ownerSchema, ctSchema, ct.SimpleContent.Restriction.Attributes, ct.SimpleContent.Restriction.AttributeGroups, visitedAG)
			}
		}
		if baseName != "" {
			if baseCT := sg.ResolveComplexType(ctSchema, baseName); baseCT != nil {
				collectCTChain(ownerSchema, baseCT, visitedCT)
			}
		}
		collectMG(ownerSchema, ctSchema, &ct.ModelGroup, visitedGroups)
	}

	for _, schema := range sg.SchemasInOrder {
		if sg.isSchemaEmpty(schema) && !schema.IsRootSchema {
			continue
		}
		// AttributeGroups
		for _, ag := range schema.AttributeGroups {
			checkAttrRefs(schema, schema, ag.Attributes, ag.AttributeGroups, make(map[*AttributeGroup]bool))
		}
		// Groups (emitted as structs in schema's package)
		for _, grp := range schema.Groups {
			collectMG(schema, schema, &grp.ModelGroup, make(map[*Group]bool))
		}
		// ComplexTypes (direct and flattened inheritance chain)
		for _, ct := range sg.AllComplexTypesBySchema[schema] {
			collectCTChain(schema, ct, make(map[*ComplexType]bool))
		}
		// Top-level Elements
		for _, el := range schema.Elements {
			if el.ComplexType != nil {
				collectCTChain(schema, el.ComplexType, make(map[*ComplexType]bool))
			} else if schema.IsRootSchema && el.Abstract != "true" && el.Type != "" {
				if targetCT := sg.ResolveComplexType(schema, el.Type); targetCT != nil {
					if xsdNameToGoIdentifier(el.NameXSD) != xsdNameToGoIdentifier(targetCT.NameXSD) {
						collectCTChain(schema, targetCT, make(map[*ComplexType]bool))
					} else {
						addDep(schema, sg.SchemaOfComplexType[targetCT])
					}
				}
			}
		}
	}

	// Ensure RootSchema depends on all non-empty schemas so that any schema that
	// transitively depends on RootSchema forms an SCC with RootSchema (since sub-packages cannot import root `models`).
	for _, schema := range sg.SchemasInOrder {
		if schema != sg.RootSchema && !sg.isSchemaEmpty(schema) {
			addDep(sg.RootSchema, schema)
		}
	}

	return deps
}

func (sg *SchemaGraph) computeTargetPackages() {
	deps := sg.computeSchemaSymbolDependencies()

	// Collect active schemas
	var activeSchemas []*Schema
	for _, s := range sg.SchemasInOrder {
		if s.IsRootSchema || !sg.isSchemaEmpty(s) {
			activeSchemas = append(activeSchemas, s)
		}
	}

	// Tarjan's Strongly Connected Components algorithm
	index := 0
	indices := make(map[*Schema]int)
	lowlink := make(map[*Schema]int)
	onStack := make(map[*Schema]bool)
	var stack []*Schema
	var sccs [][]*Schema

	var strongconnect func(v *Schema)
	strongconnect = func(v *Schema) {
		indices[v] = index
		lowlink[v] = index
		index++
		stack = append(stack, v)
		onStack[v] = true

		// Sort neighbors deterministically by SchemasInOrder index
		var neighbors []*Schema
		for w := range deps[v] {
			neighbors = append(neighbors, w)
		}
		slices.SortFunc(neighbors, func(a, b *Schema) int {
			return cmp.Compare(slices.Index(sg.SchemasInOrder, a), slices.Index(sg.SchemasInOrder, b))
		})

		for _, w := range neighbors {
			if _, visited := indices[w]; !visited {
				strongconnect(w)
				if lowlink[w] < lowlink[v] {
					lowlink[v] = lowlink[w]
				}
			} else if onStack[w] {
				if indices[w] < lowlink[v] {
					lowlink[v] = indices[w]
				}
			}
		}

		if lowlink[v] == indices[v] {
			var scc []*Schema
			for {
				w := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				onStack[w] = false
				scc = append(scc, w)
				if w == v {
					break
				}
			}
			// Sort schemas inside SCC by discovery order
			slices.SortFunc(scc, func(a, b *Schema) int {
				return cmp.Compare(slices.Index(sg.SchemasInOrder, a), slices.Index(sg.SchemasInOrder, b))
			})
			sccs = append(sccs, scc)
		}
	}

	for _, s := range activeSchemas {
		if _, visited := indices[s]; !visited {
			strongconnect(s)
		}
	}

	// Tarjan's SCC produces components in reverse topological order (leaves first, root last).
	usedPkgNames := map[string]bool{
		"models":   true,
		"model":    true,
		"probe":    true,
		"diagrams": true,
		"time":     true,
		"slices":   true,
		"strings":  true,
		"fmt":      true,
		"regexp":   true,
		"strconv":  true,
	}
	caser := cases.Title(language.English)

	for _, scc := range sccs {
		isRoot := false
		for _, s := range scc {
			if s.IsRootSchema {
				isRoot = true
				break
			}
		}

		pkg := &TargetPackage{
			IsRoot:  isRoot,
			Schemas: scc,
		}

		if isRoot {
			pkg.PackageName = "models"
			pkg.RelDir = ""
			pkg.ImportPath = sg.RootImportPath
			pkg.ImportAlias = "models"
			pkg.StageField = ""
			sg.RootPackage = pkg
		} else {
			baseName := deriveSubPackageName(scc)
			pkgName := baseName
			k := 2
			for usedPkgNames[pkgName] {
				pkgName = fmt.Sprintf("%s%d", baseName, k)
				k++
			}
			usedPkgNames[pkgName] = true

			pkg.PackageName = pkgName
			pkg.RelDir = pkgName
			pkg.ImportPath = sg.RootImportPath + "/" + pkgName
			pkg.ImportAlias = pkgName
			pkg.StageField = caser.String(pkgName) + "Stage"
		}

		for _, s := range scc {
			s.PackageName = pkg.PackageName
			s.PackageRelDir = pkg.RelDir
			sg.PackageOfSchema[s] = pkg

			// Populate constructs in pkg
			for _, ct := range sg.AllComplexTypesBySchema[s] {
				if ct.IsAnonymous && ct.OuterElement != nil && slices.Contains(s.Elements, ct.OuterElement) {
					continue
				}
				pkg.ComplexTypes = append(pkg.ComplexTypes, ct)
			}
			pkg.SimpleTypes = append(pkg.SimpleTypes, sg.AllSimpleTypesBySchema[s]...)
			for _, grp := range s.Groups {
				if grp.Ref == "" {
					pkg.Groups = append(pkg.Groups, grp)
				}
			}
			for _, ag := range s.AttributeGroups {
				if ag.Ref == "" {
					pkg.AttributeGroups = append(pkg.AttributeGroups, ag)
				}
			}
			pkg.TopElements = append(pkg.TopElements, s.Elements...)
			pkg.TopAttributes = append(pkg.TopAttributes, s.Attributes...)
		}

		sg.TargetPackages = append(sg.TargetPackages, pkg)
	}
}

func deriveSubPackageName(scc []*Schema) string {
	// If multiple schemas are coalesced into an SCC, check if they share a parent directory name (e.g. "gml", "gmd", "gco")
	if len(scc) > 1 {
		parentDir := filepath.Base(filepath.Dir(scc[0].FilePath))
		if parentDir == "3.2.1" || parentDir == "20070417" {
			parentDir = filepath.Base(filepath.Dir(filepath.Dir(scc[0].FilePath)))
		}
		cleaned := sanitizeGoPkgName(parentDir)
		if cleaned != "" && cleaned != "models" && cleaned != "tests" {
			return cleaned
		}
	}

	s := scc[0]
	rawBase := strings.TrimSuffix(filepath.Base(s.FilePath), filepath.Ext(s.FilePath))
	cleaned := sanitizeGoPkgName(rawBase)
	if strings.HasPrefix(cleaned, "gml") {
		return "gml"
	}
	if cleaned == "" {
		cleaned = "subpkg"
	}
	return cleaned
}

func sanitizeGoPkgName(raw string) string {
	raw = strings.ToLower(raw)
	var b strings.Builder
	for _, r := range raw {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
		}
	}
	res := b.String()
	if res == "" {
		return ""
	}
	if res[0] >= '0' && res[0] <= '9' {
		res = "pkg" + res
	}
	if isGoKeywordOrPredeclared(res) || isReservedPackageName(res) {
		res = res + "xsd"
	}
	return res
}

func isReservedPackageName(pkg string) bool {
	switch pkg {
	case "xml", "fmt", "strings", "time", "slices", "regexp", "strconv",
		"os", "io", "bytes", "log", "sort", "cmp", "math", "errors",
		"context", "json", "base64", "filepath", "utf16", "ast", "parser",
		"token", "probe", "split", "tree", "button", "load", "table",
		"form", "gong", "models", "test", "testing", "main":
		return true
	}
	return false
}

func (sg *SchemaGraph) assignGoIdentifiersPerPackage() {
	for _, pkg := range sg.TargetPackages {
		setOfGoIdentifiers := make(map[string]any)

		// 1. Named ComplexTypes
		for _, ct := range pkg.ComplexTypes {
			if ct.IsAnonymous {
				continue
			}
			ct.Name = ct.NameXSD
			computeGoIdentifier(ct.Name, &ct.WithGoIdentifier, setOfGoIdentifiers)
		}

		// 2. Anonymous ComplexTypes (nested inside elements)
		for _, ct := range pkg.ComplexTypes {
			if !ct.IsAnonymous {
				continue
			}
			if ct.Name == "" && ct.OuterElement != nil {
				ct.Name = "A_" + ct.OuterElement.NameXSD
			}
			computeGoIdentifier(ct.Name, &ct.WithGoIdentifier, setOfGoIdentifiers)
		}

		// 3. Top-level Elements
		for _, el := range pkg.TopElements {
			el.Name = el.NameXSD
			computeGoIdentifier(el.Name, &el.WithGoIdentifier, setOfGoIdentifiers)
		}

		// 4. Groups
		for _, grp := range pkg.Groups {
			grp.Name = grp.NameXSD
			computeGoIdentifier("Group_"+grp.Name, &grp.WithGoIdentifier, setOfGoIdentifiers)
		}

		// 5. AttributeGroups
		for _, ag := range pkg.AttributeGroups {
			ag.Name = ag.NameXSD
			computeGoIdentifier("AttributeGroup_"+ag.Name, &ag.WithGoIdentifier, setOfGoIdentifiers)
		}

		// 6. SimpleTypes (Enums)
		for _, st := range pkg.SimpleTypes {
			st.Name = st.NameXSD
			computeGoIdentifier("Enum_"+xsdNameToGoIdentifier(st.Name), &st.WithGoIdentifier, setOfGoIdentifiers)
		}

		// 7. Top-level Attributes
		for _, attr := range pkg.TopAttributes {
			attr.Name = attr.NameXSD
			computeGoIdentifier("Attr_"+attr.Name, &attr.WithGoIdentifier, setOfGoIdentifiers)
		}
	}
}

// SortedStrings returns a deterministically sorted slice of map keys.
func SortedStrings(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
