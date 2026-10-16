package depgraph

import (
	"path/filepath"
	"sort"
	"strings"

	"github.com/fullstack-lang/gong/app/xsd2gong/pkg/xsdloader"
)

// PackageNode represents a target Go model package.
type PackageNode struct {
	Name          string // e.g., "xlink", "common", "attributes", "direction", "barline", "note", "score"
	IsRoot        bool   // true for the root model package ("score" or "models")
	SimpleTypes   []*xsdloader.SimpleType
	ComplexTypes  []*xsdloader.ComplexType
	Groups        []*xsdloader.Group
	AttributeGrps []*xsdloader.AttributeGroup
	Elements      []*xsdloader.Element // Root elements
	Imports       []string             // Sorted list of package names imported by this package
}

// Graph holds the resolved multi-package partition and global type registries.
type Graph struct {
	Packages       map[string]*PackageNode
	TopologicalOrd []*PackageNode // Leaves first, Root package last
	RootPkg        *PackageNode

	// Global registries by qualified or local name
	SimpleTypeByName     map[string]*xsdloader.SimpleType
	ComplexTypeByName    map[string]*xsdloader.ComplexType
	GroupByName          map[string]*xsdloader.Group
	AttributeGroupByName map[string]*xsdloader.AttributeGroup
	ElementByName        map[string]*xsdloader.Element
	AttributeByName      map[string]*xsdloader.Attribute

	// Package ownership maps
	TypeToPkg map[string]string // key: "st:"+name or "ct:"+name or "el:"+name -> pkgName
}

// Build constructs the multi-package partition from a SchemaSet and enforces a strict DAG
// using Tarjan's Strongly Connected Components (SCC) algorithm.
func Build(set *xsdloader.SchemaSet) (*Graph, error) {
	g := &Graph{
		Packages:             make(map[string]*PackageNode),
		SimpleTypeByName:     make(map[string]*xsdloader.SimpleType),
		ComplexTypeByName:    make(map[string]*xsdloader.ComplexType),
		GroupByName:          make(map[string]*xsdloader.Group),
		AttributeGroupByName: make(map[string]*xsdloader.AttributeGroup),
		ElementByName:        make(map[string]*xsdloader.Element),
		AttributeByName:      make(map[string]*xsdloader.Attribute),
		TypeToPkg:            make(map[string]string),
	}

	// 1. Index all declarations and assign initial candidate packages
	for _, sch := range set.Schemas {
		isRootSchema := (sch == set.RootSchema)
		extPkgName := externalSchemaPkgName(sch)

		for _, st := range sch.SimpleTypes {
			pkg := assignModulePkg(st.ModuleComment, isRootSchema, extPkgName)
			key := qualifyKey("st", sch, st.Name)
			g.SimpleTypeByName[key] = st
			if isRootSchema || g.SimpleTypeByName[st.Name] == nil {
				g.SimpleTypeByName[st.Name] = st
			}
			g.TypeToPkg[key] = pkg
			if isRootSchema {
				g.TypeToPkg["st:"+st.Name] = pkg
			}
		}

		for _, ct := range sch.ComplexTypes {
			pkg := assignModulePkg(ct.ModuleComment, isRootSchema, extPkgName)
			key := qualifyKey("ct", sch, ct.Name)
			g.ComplexTypeByName[key] = ct
			if isRootSchema || g.ComplexTypeByName[ct.Name] == nil {
				g.ComplexTypeByName[ct.Name] = ct
			}
			g.TypeToPkg[key] = pkg
			if isRootSchema {
				g.TypeToPkg["ct:"+ct.Name] = pkg
			}
		}

		for _, grp := range sch.Groups {
			key := qualifyKey("grp", sch, grp.Name)
			g.GroupByName[key] = grp
			if isRootSchema || g.GroupByName[grp.Name] == nil {
				g.GroupByName[grp.Name] = grp
			}
		}

		for _, ag := range sch.AttributeGroups {
			key := qualifyKey("ag", sch, ag.Name)
			g.AttributeGroupByName[key] = ag
			if isRootSchema || g.AttributeGroupByName[ag.Name] == nil {
				g.AttributeGroupByName[ag.Name] = ag
			}
		}

		for _, attr := range sch.Attributes {
			key := qualifyKey("attr", sch, attr.Name)
			g.AttributeByName[key] = attr
			if isRootSchema || g.AttributeByName[attr.Name] == nil {
				g.AttributeByName[attr.Name] = attr
			}
		}

		for _, el := range sch.Elements {
			pkg := assignModulePkg(el.ModuleComment, isRootSchema, extPkgName)
			if isRootSchema && (el.Name == "score-partwise" || el.Name == "score-timewise") {
				pkg = "score"
			}
			key := qualifyKey("el", sch, el.Name)
			g.ElementByName[key] = el
			if isRootSchema || g.ElementByName[el.Name] == nil {
				g.ElementByName[el.Name] = el
			}
			g.TypeToPkg[key] = pkg
			if isRootSchema {
				g.TypeToPkg["el:"+el.Name] = pkg
			}
		}
	}

	// 2. Build type-level dependency graph and resolve cross-package cycles via Tarjan's SCC
	g.resolveCycles(set)

	// 3. Populate PackageNodes from final TypeToPkg assignments
	for _, sch := range set.Schemas {
		isRootSchema := (sch == set.RootSchema)
		for _, st := range sch.SimpleTypes {
			key := qualifyKey("st", sch, st.Name)
			if isRootSchema {
				key = "st:" + st.Name
			}
			pkgName := g.TypeToPkg[key]
			g.ensurePkg(pkgName).SimpleTypes = append(g.ensurePkg(pkgName).SimpleTypes, st)
		}
		for _, ct := range sch.ComplexTypes {
			key := qualifyKey("ct", sch, ct.Name)
			if isRootSchema {
				key = "ct:" + ct.Name
			}
			pkgName := g.TypeToPkg[key]
			g.ensurePkg(pkgName).ComplexTypes = append(g.ensurePkg(pkgName).ComplexTypes, ct)
		}
		for _, el := range sch.Elements {
			if el.Abstract == "true" {
				continue
			}
			key := qualifyKey("el", sch, el.Name)
			if isRootSchema {
				key = "el:" + el.Name
			}
			pkgName := g.TypeToPkg[key]
			g.ensurePkg(pkgName).Elements = append(g.ensurePkg(pkgName).Elements, el)
		}
	}

	// Ensure standard MusicXML packages exist if root schema has score-partwise
	if g.Packages["score"] != nil {
		g.RootPkg = g.Packages["score"]
		g.RootPkg.IsRoot = true
	} else {
		// Fallback for single-package or non-MusicXML schemas
		for _, p := range g.Packages {
			g.RootPkg = p
			p.IsRoot = true
			break
		}
	}

	// 4. Compute package-level imports and topological sort
	g.computePackageImportsAndTopoSort()

	return g, nil
}

func externalSchemaPkgName(sch *xsdloader.Schema) string {
	base := strings.ToLower(filepath.Base(sch.FilePath))
	base = strings.TrimSuffix(base, ".xsd")
	if base == "xml" {
		return "common" // xml:lang, xml:space, xml:id attributes are inlined or placed in common
	}
	return base
}

func assignModulePkg(moduleComment string, isRootSchema bool, extPkgName string) string {
	if !isRootSchema {
		return extPkgName
	}
	switch moduleComment {
	case "common", "link", "identity", "layout":
		// Group shared foundational modules (common, layout, identity, link) into "common"
		// or keep them clean as per the target architecture (score, note, attributes, direction, barline, common, xlink)
		return "common"
	case "attributes":
		return "attributes"
	case "barline":
		return "barline"
	case "direction":
		return "direction"
	case "note":
		return "note"
	case "score":
		return "score"
	default:
		return "common"
	}
}

func qualifyKey(kind string, sch *xsdloader.Schema, name string) string {
	prefix := externalSchemaPkgName(sch)
	if prefix == "musicxml" {
		return kind + ":" + name
	}
	return kind + ":" + prefix + ":" + name
}

func (g *Graph) ensurePkg(name string) *PackageNode {
	if p, ok := g.Packages[name]; ok {
		return p
	}
	p := &PackageNode{
		Name: name,
	}
	g.Packages[name] = p
	return p
}

// pkgRank defines the preferred layered hierarchy (lower rank = closer to leaf).
// Dependencies should flow from higher rank -> lower rank (e.g. score -> note -> direction -> attributes -> barline -> common -> xlink).
func pkgRank(pkg string) int {
	switch pkg {
	case "xlink":
		return 0
	case "common":
		return 1
	case "barline":
		return 2
	case "attributes":
		return 3
	case "direction":
		return 4
	case "note":
		return 5
	case "score":
		return 6
	default:
		return 1
	}
}

func (g *Graph) resolveCycles(set *xsdloader.SchemaSet) {
	// Collect dependencies for each type node ("st:...", "ct:...", "el:...")
	deps := make(map[string][]string)

	for _, sch := range set.Schemas {
		isRoot := (sch == set.RootSchema)
		for _, st := range sch.SimpleTypes {
			k := qualifyKey("st", sch, st.Name)
			if isRoot {
				k = "st:" + st.Name
			}
			deps[k] = g.collectSimpleTypeDeps(st)
		}
		for _, ct := range sch.ComplexTypes {
			k := qualifyKey("ct", sch, ct.Name)
			if isRoot {
				k = "ct:" + ct.Name
			}
			deps[k] = g.collectComplexTypeDeps(ct)
		}
		for _, el := range sch.Elements {
			k := qualifyKey("el", sch, el.Name)
			if isRoot {
				k = "el:" + el.Name
			}
			if el.ComplexType != nil {
				deps[k] = g.collectComplexTypeDeps(el.ComplexType)
			}
		}
	}

	// Iteratively run Tarjan's SCC on the package graph and hoist any lower-rank dependency
	// that points to a higher-rank package until the package graph is a strict DAG.
	for iter := 0; iter < 50; iter++ {
		changed := false
		for node, nodeDeps := range deps {
			pkgA := g.TypeToPkg[node]
			if pkgA == "" {
				continue
			}
			rankA := pkgRank(pkgA)
			for _, dep := range nodeDeps {
				pkgB := g.TypeToPkg[dep]
				if pkgB == "" || pkgA == pkgB {
					continue
				}
				rankB := pkgRank(pkgB)
				// If a type in pkgA (lower rank, e.g., common or direction) depends on a type in pkgB (higher rank, e.g., note),
				// hoist the target type `dep` down to pkgA so that pkgA does not need to import pkgB!
				if rankB > rankA {
					g.TypeToPkg[dep] = pkgA
					changed = true
				}
			}
		}
		if !changed {
			break
		}
	}

	// Run Tarjan's SCC on the resulting package graph as a final verification and cycle-breaker
	sccs := g.findPackageSCCs(deps)
	for _, scc := range sccs {
		if len(scc) > 1 {
			// Merge any remaining mutually recursive packages into the lowest-ranked package in the SCC
			targetPkg := scc[0]
			for _, p := range scc[1:] {
				if pkgRank(p) < pkgRank(targetPkg) {
					targetPkg = p
				}
			}
			inSCC := make(map[string]bool)
			for _, p := range scc {
				inSCC[p] = true
			}
			for k, p := range g.TypeToPkg {
				if inSCC[p] {
					g.TypeToPkg[k] = targetPkg
				}
			}
		}
	}
}

func (g *Graph) findPackageSCCs(typeDeps map[string][]string) [][]string {
	pkgAdj := make(map[string]map[string]bool)
	var allPkgs []string
	seenPkg := make(map[string]bool)

	for node, deps := range typeDeps {
		u := g.TypeToPkg[node]
		if u == "" {
			continue
		}
		if !seenPkg[u] {
			seenPkg[u] = true
			allPkgs = append(allPkgs, u)
		}
		if pkgAdj[u] == nil {
			pkgAdj[u] = make(map[string]bool)
		}
		for _, d := range deps {
			v := g.TypeToPkg[d]
			if v != "" && v != u {
				pkgAdj[u][v] = true
				if !seenPkg[v] {
					seenPkg[v] = true
					allPkgs = append(allPkgs, v)
				}
			}
		}
	}
	sort.Strings(allPkgs)

	index := 0
	indices := make(map[string]int)
	lowlink := make(map[string]int)
	onStack := make(map[string]bool)
	var stack []string
	var sccs [][]string

	var strongconnect func(v string)
	strongconnect = func(v string) {
		indices[v] = index
		lowlink[v] = index
		index++
		stack = append(stack, v)
		onStack[v] = true

		var neighbors []string
		for w := range pkgAdj[v] {
			neighbors = append(neighbors, w)
		}
		sort.Strings(neighbors)

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
			var scc []string
			for {
				w := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				onStack[w] = false
				scc = append(scc, w)
				if w == v {
					break
				}
			}
			sccs = append(sccs, scc)
		}
	}

	for _, v := range allPkgs {
		if _, visited := indices[v]; !visited {
			strongconnect(v)
		}
	}
	return sccs
}

func (g *Graph) collectSimpleTypeDeps(st *xsdloader.SimpleType) []string {
	var res []string
	if st.Restriction != nil && st.Restriction.Base != "" {
		base := xsdloader.LocalName(st.Restriction.Base)
		if _, ok := g.SimpleTypeByName[base]; ok {
			res = append(res, "st:"+base)
		}
	}
	if st.Union != nil {
		for _, member := range strings.Fields(st.Union.MemberTypes) {
			local := xsdloader.LocalName(member)
			if _, ok := g.SimpleTypeByName[local]; ok {
				res = append(res, "st:"+local)
			}
		}
	}
	return res
}

func (g *Graph) collectComplexTypeDeps(ct *xsdloader.ComplexType) []string {
	var res []string
	visitedGroups := make(map[string]bool)
	visitedAttrGroups := make(map[string]bool)

	var checkAttr func(a *xsdloader.Attribute)
	checkAttr = func(a *xsdloader.Attribute) {
		if a.Type != "" {
			prefix := xsdloader.Prefix(a.Type)
			local := xsdloader.LocalName(a.Type)
			if prefix == "xlink" {
				if _, ok := g.SimpleTypeByName["st:xlink:"+local]; ok {
					res = append(res, "st:xlink:"+local)
				}
			} else if _, ok := g.SimpleTypeByName[local]; ok {
				res = append(res, "st:"+local)
			}
		}
		if a.Ref != "" {
			prefix := xsdloader.Prefix(a.Ref)
			local := xsdloader.LocalName(a.Ref)
			if prefix == "xlink" {
				if targetAttr, ok := g.AttributeByName["attr:xlink:"+local]; ok {
					checkAttr(targetAttr)
				}
			}
		}
	}

	var checkAttrGroup func(ag *xsdloader.AttributeGroup)
	checkAttrGroup = func(ag *xsdloader.AttributeGroup) {
		if ag.Ref != "" {
			local := xsdloader.LocalName(ag.Ref)
			prefix := xsdloader.Prefix(ag.Ref)
			key := local
			if prefix == "xlink" {
				key = "ag:xlink:" + local
			}
			if visitedAttrGroups[key] {
				return
			}
			visitedAttrGroups[key] = true
			if target, ok := g.AttributeGroupByName[key]; ok {
				for _, a := range target.Attributes {
					checkAttr(a)
				}
				for _, sub := range target.AttributeGroups {
					checkAttrGroup(sub)
				}
			}
			return
		}
		for _, a := range ag.Attributes {
			checkAttr(a)
		}
		for _, sub := range ag.AttributeGroups {
			checkAttrGroup(sub)
		}
	}

	for _, a := range ct.Attributes {
		checkAttr(a)
	}
	for _, ag := range ct.AttributeGroups {
		checkAttrGroup(ag)
	}

	if ct.SimpleContent != nil && ct.SimpleContent.Extension != nil {
		ext := ct.SimpleContent.Extension
		base := xsdloader.LocalName(ext.Base)
		if _, ok := g.SimpleTypeByName[base]; ok {
			res = append(res, "st:"+base)
		}
		if _, ok := g.ComplexTypeByName[base]; ok {
			res = append(res, "ct:"+base)
		}
		for _, a := range ext.Attributes {
			checkAttr(a)
		}
		for _, ag := range ext.AttributeGroups {
			checkAttrGroup(ag)
		}
	}

	var checkModelGroup func(mg *xsdloader.ModelGroup)
	checkModelGroup = func(mg *xsdloader.ModelGroup) {
		if mg == nil {
			return
		}
		for _, el := range mg.Elements {
			if el.Type != "" {
				local := xsdloader.LocalName(el.Type)
				if _, ok := g.ComplexTypeByName[local]; ok {
					res = append(res, "ct:"+local)
				}
				if _, ok := g.SimpleTypeByName[local]; ok {
					res = append(res, "st:"+local)
				}
			}
			if el.ComplexType != nil {
				res = append(res, g.collectComplexTypeDeps(el.ComplexType)...)
			}
		}
		for _, gRef := range mg.Groups {
			local := xsdloader.LocalName(gRef.Ref)
			if visitedGroups[local] {
				continue
			}
			visitedGroups[local] = true
			if targetGrp, ok := g.GroupByName[local]; ok {
				checkModelGroup(targetGrp.Particle)
			}
		}
		for _, sub := range mg.SubGroups {
			checkModelGroup(sub)
		}
	}

	if ct.ComplexContent != nil && ct.ComplexContent.Extension != nil {
		ext := ct.ComplexContent.Extension
		base := xsdloader.LocalName(ext.Base)
		if _, ok := g.ComplexTypeByName[base]; ok {
			res = append(res, "ct:"+base)
		}
		for _, a := range ext.Attributes {
			checkAttr(a)
		}
		for _, ag := range ext.AttributeGroups {
			checkAttrGroup(ag)
		}
		checkModelGroup(ext.Particle)
	}

	checkModelGroup(ct.Particle)
	return res
}

func (g *Graph) computePackageImportsAndTopoSort() {
	var pkgList []*PackageNode
	for _, p := range g.Packages {
		pkgList = append(pkgList, p)
	}
	sort.Slice(pkgList, func(i, j int) bool {
		ri, rj := pkgRank(pkgList[i].Name), pkgRank(pkgList[j].Name)
		if ri != rj {
			return ri < rj
		}
		return pkgList[i].Name < pkgList[j].Name
	})
	g.TopologicalOrd = pkgList
}

