package diagramlayout

import (
	"embed"
	"flag"
	"fmt"
	"math"
	"net/http"
	"os"
	"slices"
	"strings"
	"testing"

	gong_models "github.com/fullstack-lang/gong/go/models"
	doc_models "github.com/fullstack-lang/gong/lib/doc/go/models"
	"github.com/fullstack-lang/gong/lib/doc/go/prepare"
	split "github.com/fullstack-lang/gong/lib/split/go/models"
)

var regenDiagramsFlag = flag.Bool("regen-diagrams", false, "regenerate diagrams/diagrams_set.go in Small-A3-Landscape layout")

// StructRef identifies a Go struct or enum by package short name and identifier.
type StructRef struct {
	Pkg  string
	Name string
}

// DiagramSpec declaratively specifies the domain elements of one UML Classdiagram.
type DiagramSpec struct {
	Name        string
	NoteName    string
	Description string
	Features    []string
	Objects     []string
	ExtraCols   [5][]StructRef
	Enums       []StructRef
}

// DomainConfig configures UML diagram verification and optional regeneration for an AIXM domain.
type DomainConfig struct {
	DomainName         string
	DomainImportPrefix string
	GoModelsDir        embed.FS
	GoDiagramsDir      embed.FS
	Specs              []DiagramSpec
	// StrictMaxWidth and StrictMaxHeight enforce Small-A3-Landscape bounds when regenerating.
	StrictMaxWidth  float64
	StrictMaxHeight float64
}

type rawEdge struct {
	srcKey    string
	srcPkg    string
	srcName   string
	dstKey    string
	dstPkg    string
	dstName   string
	fieldName string
	mult      doc_models.MultiplicityType
}

type structInfo struct {
	ref          StructRef
	key          string
	gs           *gong_models.GongStruct
	displayTitle string
	attrNames    []string
	attrTypes    []string
	attrLines    []string
	minWidth     float64
	baseHeight   float64
	origIdx      int
}

type enumInfo struct {
	ref          StructRef
	ge           *gong_models.GongEnum
	displayTitle string
	valNames     []string
	width        float64
	height       float64
}

type linkRouteKind int

const (
	routeVerticalDown linkRouteKind = iota
	routeForwardHoriz
	routeRightBus
	routeBackwardHoriz
)

type plannedLink struct {
	edge              rawEdge
	kind              linkRouteKind
	srcCol            int
	srcRow            int
	dstCol            int
	dstRow            int
	startOrientation  doc_models.OrientationType
	endOrientation    doc_models.OrientationType
	startRatio        float64
	endRatio          float64
	cornerOffsetRatio float64
	relCornerX        float64
}

type placedBox struct {
	x      float64
	y      float64
	width  float64
	height float64
}

type candidatePlan struct {
	numCols       int
	colStructs    [][]string
	boxes         map[string]placedBox
	enums         []placedBox
	hasNote       bool
	noteBox       placedBox
	noteTargetKey string
	links         []*plannedLink
	score         float64
	metrics       DiagramMetrics
}

// RunDomainUMLDiagramsTest verifies diagrams/diagrams_set.go by default without modifying it.
// When GONG_REGEN_DIAGRAMS=1 is set in the environment, it regenerates diagrams/diagrams_set.go
// fitted for a small A3 landscape rendering (~1.414:1 aspect ratio) and verifies zero collisions.
func RunDomainUMLDiagramsTest(t *testing.T, cfg DomainConfig) {
	t.Helper()

	metaPackageImports := []*doc_models.MetaPackageImport{
		{Alias: "ref_models", Path: fmt.Sprintf("%q", cfg.DomainImportPrefix)},
		{Alias: "ref_aixmfeatures", Path: fmt.Sprintf("%q", cfg.DomainImportPrefix+"/aixmfeatures")},
		{Alias: "ref_aixmdatatypes", Path: fmt.Sprintf("%q", cfg.DomainImportPrefix+"/aixmdatatypes")},
		{Alias: "ref_aixmabstractgmlobjecttypes", Path: fmt.Sprintf("%q", cfg.DomainImportPrefix+"/aixmabstractgmlobjecttypes")},
		{Alias: "ref_gml", Path: fmt.Sprintf("%q", cfg.DomainImportPrefix+"/gml")},
		{Alias: "ref_xlink", Path: fmt.Sprintf("%q", cfg.DomainImportPrefix+"/xlink")},
	}

	regen := *regenDiagramsFlag || os.Getenv("GONG_REGEN_DIAGRAMS") == "1"

	if !regen {
		verifyExistingDiagramsWithoutOverwrite(t, cfg, metaPackageImports)
		return
	}

	mux := http.NewServeMux()
	diagramEditor := &split.AsSplitArea{
		Name:             "Bottom",
		ShowNameInHeader: false,
		Size:             50,
	}
	docStager := prepare.PrepareStageSet(
		mux,
		true,
		cfg.DomainImportPrefix+":diag_gen",
		metaPackageImports,
		cfg.GoModelsDir,
		cfg.GoDiagramsDir,
		diagramEditor,
		map[string]int{},
	)
	docStage := docStager.GetStage()
	gongStage := docStager.GetGongStage()

	structsByPkgName := make(map[string]*gong_models.GongStruct)
	for gs := range *gongStage.GetInstancesSet[*gong_models.GongStruct]() {
		pkgName := "models"
		if gs.ModelPkg != nil && gs.ModelPkg.PkgGoName != "" {
			pkgName = gs.ModelPkg.PkgGoName
		}
		structsByPkgName[pkgName+"."+gs.Name] = gs
	}
	enumsByName := gongStage.GetInstancesMapByName[*gong_models.GongEnum]()
	notesByName := gongStage.GetInstancesMapByName[*gong_models.GongNote]()

	docStage.Reset()

	diagramPkg := (&doc_models.DiagramPackage{
		Name: "AIXM 5.1.1 UML Navigator Diagrams - " + cfg.DomainName,
	}).Stage(docStage)

	for diagIdx, spec := range cfg.Specs {
		cd := populateSingleDiagram(t, docStage, spec, diagIdx, structsByPkgName, enumsByName, notesByName)
		diagramPkg.Classdiagrams = append(diagramPkg.Classdiagrams, cd)
		if diagIdx == 0 {
			diagramPkg.SelectedClassdiagram = cd
		}

		dm := AnalyzeClassdiagram(cd)
		t.Logf("[%s] %s -> %.0f x %.0f px (ratio %.2f), structs=%d, enums=%d",
			cfg.DomainName, cd.Name, dm.Width, dm.Height, dm.AspectRatio,
			len(cd.GongStructShapes), len(cd.GongEnumShapes))

		assertZeroGeometryCollisions(t, cfg, cd.Name, dm)
	}

	docStage.ComputeReferenceAndOrders()

	outPath := "diagrams/diagrams_set.go"
	outFile, err := os.Create(outPath)
	if err != nil {
		t.Fatalf("failed creating %s: %v", outPath, err)
	}
	docStage.Marshall(outFile, "github.com/fullstack-lang/gong/lib/doc/go/models", "diagrams")
	if err := outFile.Close(); err != nil {
		t.Fatalf("failed closing %s: %v", outPath, err)
	}

	origStructs := len(*docStage.GetInstancesSet[*doc_models.GongStructShape]())
	origAttrs := len(*docStage.GetInstancesSet[*doc_models.AttributeShape]())
	origLinks := len(*docStage.GetInstancesSet[*doc_models.LinkShape]())
	origEnums := len(*docStage.GetInstancesSet[*doc_models.GongEnumShape]())
	origEnumVals := len(*docStage.GetInstancesSet[*doc_models.GongEnumValueShape]())
	origNotes := len(*docStage.GetInstancesSet[*doc_models.GongNoteShape]())
	origNoteLinks := len(*docStage.GetInstancesSet[*doc_models.GongNoteLinkShape]())
	origDiagrams := len(*docStage.GetInstancesSet[*doc_models.Classdiagram]())

	// Parse back from disk with embeddedDiagrams=false? Wait: PrepareStageSet(false) installs
	// BeforeCommitImplementation, which only writes if enforceSemantic mutates the stage.
	verifyMux := http.NewServeMux()
	verifyEditor := &split.AsSplitArea{Name: "VerifyBottom", Size: 50}
	verifyStager := prepare.PrepareStageSet(
		verifyMux,
		false,
		cfg.DomainImportPrefix+":diag_verify",
		metaPackageImports,
		cfg.GoModelsDir,
		cfg.GoDiagramsDir,
		verifyEditor,
		map[string]int{},
	)
	verifyStage := verifyStager.GetStage()

	assertStageCountsMatch(t, cfg, verifyStage, origDiagrams, origStructs, origAttrs, origLinks, origEnums, origEnumVals, origNotes, origNoteLinks)
}

func verifyExistingDiagramsWithoutOverwrite(
	t *testing.T,
	cfg DomainConfig,
	metaPackageImports []*doc_models.MetaPackageImport,
) {
	t.Helper()

	const diagramsPath = "diagrams/diagrams_set.go"
	beforeBytes, err := os.ReadFile(diagramsPath)
	if err != nil {
		t.Fatalf("failed reading %s: %v", diagramsPath, err)
	}

	rawStage := doc_models.NewStage(cfg.DomainImportPrefix + ":diag_raw")
	rawStage.MetaPackageImports = metaPackageImports
	if err := rawStage.ParseAstFile(diagramsPath, false); err != nil {
		t.Fatalf("failed parsing %s: %v", diagramsPath, err)
	}

	origStructs := len(*rawStage.GetInstancesSet[*doc_models.GongStructShape]())
	origAttrs := len(*rawStage.GetInstancesSet[*doc_models.AttributeShape]())
	origLinks := len(*rawStage.GetInstancesSet[*doc_models.LinkShape]())
	origEnums := len(*rawStage.GetInstancesSet[*doc_models.GongEnumShape]())
	origEnumVals := len(*rawStage.GetInstancesSet[*doc_models.GongEnumValueShape]())
	origNotes := len(*rawStage.GetInstancesSet[*doc_models.GongNoteShape]())
	origNoteLinks := len(*rawStage.GetInstancesSet[*doc_models.GongNoteLinkShape]())
	origDiagrams := len(*rawStage.GetInstancesSet[*doc_models.Classdiagram]())

	if origDiagrams == 0 || origStructs == 0 {
		t.Fatalf("%s has empty UML diagrams (diagrams=%d, structs=%d)", diagramsPath, origDiagrams, origStructs)
	}

	// Validate using embeddedDiagrams=true so PrepareStageSet NEVER writes to diagrams/diagrams_set.go.
	verifyMux := http.NewServeMux()
	verifyEditor := &split.AsSplitArea{Name: "VerifyReadOnly", Size: 50}
	verifyStager := prepare.PrepareStageSet(
		verifyMux,
		true,
		cfg.DomainImportPrefix+":diag_readonly_verify",
		metaPackageImports,
		cfg.GoModelsDir,
		cfg.GoDiagramsDir,
		verifyEditor,
		map[string]int{},
	)
	verifyStage := verifyStager.GetStage()

	afterBytes, err := os.ReadFile(diagramsPath)
	if err != nil {
		t.Fatalf("failed re-reading %s: %v", diagramsPath, err)
	}
	if string(beforeBytes) != string(afterBytes) {
		t.Fatalf("diagrams/diagrams_set.go was unexpectedly modified during go test")
	}

	afterDiagrams := len(*verifyStage.GetInstancesSet[*doc_models.Classdiagram]())
	if afterDiagrams == 0 {
		t.Fatalf("PrepareStageSet loaded 0 diagrams for %s", cfg.DomainName)
	}

	assertStageCountsMatch(t, cfg, verifyStage, origDiagrams, origStructs, origAttrs, origLinks, origEnums, origEnumVals, origNotes, origNoteLinks)

	// Check every diagram in rawStage for basic A3 readability & zero overlaps
	diagrams := rawStage.GetInstancesSorted[*doc_models.Classdiagram]()
	for _, cd := range diagrams {
		dm := AnalyzeClassdiagram(cd)
		t.Logf("[%s] verified %s -> %.0f x %.0f px (ratio %.2f, overlaps=%d, labelBoxHits=%d)",
			cfg.DomainName, cd.Name, dm.Width, dm.Height, dm.AspectRatio, len(dm.BoxOverlaps), len(dm.LabelBoxHits))
		assertZeroGeometryCollisions(t, cfg, cd.Name, dm)
	}
}

func assertStageCountsMatch(
	t *testing.T,
	cfg DomainConfig,
	verifyStage *doc_models.Stage,
	origDiagrams, origStructs, origAttrs, origLinks, origEnums, origEnumVals, origNotes, origNoteLinks int,
) {
	t.Helper()
	afterStructs := len(*verifyStage.GetInstancesSet[*doc_models.GongStructShape]())
	afterAttrs := len(*verifyStage.GetInstancesSet[*doc_models.AttributeShape]())
	afterLinks := len(*verifyStage.GetInstancesSet[*doc_models.LinkShape]())
	afterEnums := len(*verifyStage.GetInstancesSet[*doc_models.GongEnumShape]())
	afterEnumVals := len(*verifyStage.GetInstancesSet[*doc_models.GongEnumValueShape]())
	afterNotes := len(*verifyStage.GetInstancesSet[*doc_models.GongNoteShape]())
	afterNoteLinks := len(*verifyStage.GetInstancesSet[*doc_models.GongNoteLinkShape]())
	afterDiagrams := len(*verifyStage.GetInstancesSet[*doc_models.Classdiagram]())

	t.Logf("Domain %s: %d UML diagrams (%d GongStructShapes, %d AttributeShapes, %d LinkShapes, %d GongEnumShapes, %d GongEnumValueShapes, %d GongNoteShapes, %d GongNoteLinkShapes)",
		cfg.DomainName, afterDiagrams, afterStructs, afterAttrs, afterLinks, afterEnums, afterEnumVals, afterNotes, afterNoteLinks)

	if afterDiagrams != origDiagrams || afterDiagrams < len(cfg.Specs) {
		t.Fatalf("expected at least %d diagrams (orig=%d), got %d", len(cfg.Specs), origDiagrams, afterDiagrams)
	}
	if afterStructs != origStructs {
		t.Fatalf("enforceSemantic removed GongStructShapes: before=%d after=%d", origStructs, afterStructs)
	}
	if afterAttrs != origAttrs {
		t.Fatalf("enforceSemantic removed AttributeShapes: before=%d after=%d", origAttrs, afterAttrs)
	}
	if afterLinks != origLinks {
		t.Fatalf("enforceSemantic removed LinkShapes: before=%d after=%d", origLinks, afterLinks)
	}
	if afterEnums != origEnums {
		t.Fatalf("enforceSemantic removed GongEnumShapes: before=%d after=%d", origEnums, afterEnums)
	}
	if afterEnumVals != origEnumVals {
		t.Fatalf("enforceSemantic removed GongEnumValueShapes: before=%d after=%d", origEnumVals, afterEnumVals)
	}
	if afterNotes != origNotes {
		t.Fatalf("expected %d GongNoteShapes, got %d", origNotes, afterNotes)
	}
	if afterNoteLinks != origNoteLinks {
		t.Fatalf("expected %d GongNoteLinkShapes, got %d", origNoteLinks, afterNoteLinks)
	}
}

func assertZeroGeometryCollisions(t *testing.T, cfg DomainConfig, diagName string, dm DiagramMetrics) {
	t.Helper()
	if len(dm.BoxOverlaps) > 0 {
		t.Errorf("%s: box overlaps: %v", diagName, dm.BoxOverlaps)
	}
	if len(dm.TextClipIssues) > 0 {
		t.Errorf("%s: text clipping issues: %v", diagName, dm.TextClipIssues)
	}
	if len(dm.SegmentPiercings) > 0 {
		t.Errorf("%s: link segment pierces box: %v", diagName, dm.SegmentPiercings)
	}
	if len(dm.LabelBoxHits) > 0 {
		t.Errorf("%s: link label overlaps box: %v", diagName, dm.LabelBoxHits)
	}
	if len(dm.LabelLabelHits) > 0 {
		t.Errorf("%s: link labels overlap each other: %v", diagName, dm.LabelLabelHits)
	}
	if len(dm.NoteLinkHits) > 0 {
		t.Errorf("%s: note link crosses intermediate box: %v", diagName, dm.NoteLinkHits)
	}
	maxW := cfg.StrictMaxWidth
	if maxW <= 0 {
		maxW = 2200.0
	}
	maxH := cfg.StrictMaxHeight
	if maxH <= 0 {
		maxH = 1550.0
	}
	if dm.Width > maxW {
		t.Errorf("%s: diagram width %.0f exceeds small A3 landscape max %.0f", diagName, dm.Width, maxW)
	}
	if dm.Height > maxH {
		t.Errorf("%s: diagram height %.0f exceeds small A3 landscape max %.0f", diagName, dm.Height, maxH)
	}
}

func populateSingleDiagram(
	t *testing.T,
	docStage *doc_models.Stage,
	spec DiagramSpec,
	diagIdx int,
	structsByPkgName map[string]*gong_models.GongStruct,
	enumsByName map[string]*gong_models.GongEnum,
	notesByName map[string]*gong_models.GongNote,
) *doc_models.Classdiagram {
	t.Helper()

	var allRefs []StructRef
	seen := make(map[string]bool)
	addRef := func(ref StructRef) {
		key := ref.Pkg + "." + ref.Name
		if seen[key] {
			return
		}
		if _, ok := structsByPkgName[key]; !ok {
			return
		}
		seen[key] = true
		allRefs = append(allRefs, ref)
	}

	for _, feat := range spec.Features {
		base := strings.TrimPrefix(feat, "Abstract")
		addRef(StructRef{Pkg: "aixmfeatures", Name: base + "TimeSliceType"})
	}
	for _, obj := range spec.Objects {
		base := strings.TrimPrefix(obj, "Abstract")
		addRef(StructRef{Pkg: "aixmfeatures", Name: base + "PropertyType"})
		addRef(StructRef{Pkg: "aixmfeatures", Name: obj + "Type"})
	}
	for c := range 5 {
		for _, ref := range spec.ExtraCols[c] {
			addRef(ref)
		}
	}

	hasEdgeInSet := func(checkKey string) bool {
		for _, srcRef := range allRefs {
			srcKey := srcRef.Pkg + "." + srcRef.Name
			srcGS := structsByPkgName[srcKey]
			if srcGS == nil {
				continue
			}
			for _, field := range srcGS.Fields {
				targetGS := fieldTargetStruct(field)
				if targetGS == nil {
					continue
				}
				tKey := structPkgKey(targetGS)
				if !seen[tKey] || tKey == srcKey {
					continue
				}
				if srcKey == checkKey || tKey == checkKey {
					return true
				}
			}
		}
		return false
	}
	for _, feat := range spec.Features {
		base := strings.TrimPrefix(feat, "Abstract")
		tsKey := "aixmfeatures." + base + "TimeSliceType"
		if seen[tsKey] && !hasEdgeInSet(tsKey) {
			addRef(StructRef{Pkg: "aixmfeatures", Name: feat + "Type"})
			addRef(StructRef{Pkg: "aixmfeatures", Name: base + "TimeSlicePropertyType"})
		}
	}

	// Build structInfo metadata
	infos := make(map[string]*structInfo, len(allRefs))
	for idx, ref := range allRefs {
		key := ref.Pkg + "." + ref.Name
		gs := structsByPkgName[key]
		displayTitle := gs.Name
		if ref.Pkg != "" && ref.Pkg != "models" {
			displayTitle = ref.Pkg + "." + gs.Name
		}
		var attrNames, attrTypes, attrLines []string
		const maxAttrs = 4
		for _, field := range gs.Fields {
			if field.GetName() == "Name" {
				continue
			}
			if len(attrNames) >= maxAttrs {
				break
			}
			switch rf := field.(type) {
			case *gong_models.GongBasicField:
				parts := strings.Split(rf.DeclaredType, ".")
				ftName := parts[len(parts)-1]
				attrNames = append(attrNames, rf.Name)
				attrTypes = append(attrTypes, ftName)
				attrLines = append(attrLines, rf.Name+" : "+ftName)
			case *gong_models.GongTimeField:
				attrNames = append(attrNames, rf.Name)
				attrTypes = append(attrTypes, "Time")
				attrLines = append(attrLines, rf.Name+" : Time")
			}
		}
		minW := StructBoxMinWidth(displayTitle, attrLines, true)
		baseH := math.Max(66.0, 48.0+float64(len(attrLines))*20.0)
		infos[key] = &structInfo{
			ref:          ref,
			key:          key,
			gs:           gs,
			displayTitle: displayTitle,
			attrNames:    attrNames,
			attrTypes:    attrTypes,
			attrLines:    attrLines,
			minWidth:     minW,
			baseHeight:   baseH,
			origIdx:      idx,
		}
	}

	// Count distinct source structs targeting each included struct so shared hubs
	// (targeted by >= 3 distinct structs) cap parallel links at 2 per pair while
	// 1-to-1 or 2-source targets keep up to 4 parallel semantic links.
	distinctSrcToDst := make(map[string]int)
	for _, srcRef := range allRefs {
		srcKey := srcRef.Pkg + "." + srcRef.Name
		srcGS := structsByPkgName[srcKey]
		hasValidTime := false
		for _, field := range srcGS.Fields {
			if field.GetName() == "ValidTime" {
				hasValidTime = true
				break
			}
		}
		seenTarget := make(map[string]bool)
		for _, field := range srcGS.Fields {
			if field.GetName() == "FeatureLifetime" && hasValidTime {
				continue
			}
			targetGS := fieldTargetStruct(field)
			if targetGS == nil {
				continue
			}
			tKey := structPkgKey(targetGS)
			if !seen[tKey] || tKey == srcKey || seenTarget[tKey] {
				continue
			}
			seenTarget[tKey] = true
			distinctSrcToDst[tKey]++
		}
	}

	// Collect raw edges among included structs (up to 4 parallel semantic links per pair,
	// or up to 2 when >= 3 distinct source structs fan into the same target,
	// skipping FeatureLifetime when ValidTime is already present on a TimeSlice).
	var edges []rawEdge
	for _, srcRef := range allRefs {
		srcKey := srcRef.Pkg + "." + srcRef.Name
		srcGS := structsByPkgName[srcKey]
		hasValidTime := false
		for _, field := range srcGS.Fields {
			if field.GetName() == "ValidTime" {
				hasValidTime = true
				break
			}
		}
		pairCount := make(map[string]int)
		for _, field := range srcGS.Fields {
			if field.GetName() == "FeatureLifetime" && hasValidTime {
				continue
			}
			var targetGS *gong_models.GongStruct
			var mult doc_models.MultiplicityType
			switch rf := field.(type) {
			case *gong_models.PointerToGongStructField:
				targetGS = rf.GongStruct
				mult = doc_models.ZERO_ONE
			case *gong_models.SliceOfPointerToGongStructField:
				targetGS = rf.GongStruct
				mult = doc_models.MANY
			}
			if targetGS == nil {
				continue
			}
			tPkg := "models"
			if targetGS.ModelPkg != nil && targetGS.ModelPkg.PkgGoName != "" {
				tPkg = targetGS.ModelPkg.PkgGoName
			}
			tKey := tPkg + "." + targetGS.Name
			if !seen[tKey] || tKey == srcKey {
				continue
			}
			maxParallel := 4
			if distinctSrcToDst[tKey] >= 3 {
				maxParallel = 2
			}
			if pairCount[tKey] >= maxParallel {
				continue
			}
			pairCount[tKey]++
			edges = append(edges, rawEdge{
				srcKey:    srcKey,
				srcPkg:    srcRef.Pkg,
				srcName:   srcGS.Name,
				dstKey:    tKey,
				dstPkg:    tPkg,
				dstName:   targetGS.Name,
				fieldName: field.GetName(),
				mult:      mult,
			})
		}
	}

	// Collect enumInfo metadata
	var enums []*enumInfo
	for _, enumRef := range spec.Enums {
		ge, ok := enumsByName[enumRef.Name]
		if !ok {
			ge, ok = enumsByName[enumRef.Name+"_Member_1"]
		}
		if !ok {
			continue
		}
		var valNames []string
		const maxEnumVals = 6
		for i, ev := range ge.GongEnumValues {
			if i >= maxEnumVals {
				break
			}
			valNames = append(valNames, ev.Name)
		}
		displayTitle := ge.Name
		if enumRef.Pkg != "" && enumRef.Pkg != "models" {
			displayTitle = enumRef.Pkg + "." + ge.Name
		}
		ew, eh := EnumBoxSize(displayTitle, valNames)
		enums = append(enums, &enumInfo{
			ref:          enumRef,
			ge:           ge,
			displayTitle: displayTitle,
			valNames:     valNames,
			width:        ew,
			height:       eh,
		})
	}

	var gongNote *gong_models.GongNote
	if spec.NoteName != "" {
		if gn, ok := notesByName[spec.NoteName]; ok {
			gongNote = gn
		}
	}

	bestPlan := chooseBestA3LandscapePlan(allRefs, infos, edges, enums, gongNote)
	return materializePlan(docStage, spec, diagIdx, infos, enums, gongNote, bestPlan)
}

func fieldTargetStruct(field any) *gong_models.GongStruct {
	switch rf := field.(type) {
	case *gong_models.PointerToGongStructField:
		return rf.GongStruct
	case *gong_models.SliceOfPointerToGongStructField:
		return rf.GongStruct
	}
	return nil
}

func structPkgKey(gs *gong_models.GongStruct) string {
	pkg := "models"
	if gs.ModelPkg != nil && gs.ModelPkg.PkgGoName != "" {
		pkg = gs.ModelPkg.PkgGoName
	}
	return pkg + "." + gs.Name
}

// buildVerticalTowers groups PropertyType -> ObjectType wrappers and clean 1-to-1 spines
// into vertical sub-columns so chains flow downward instead of sprawling horizontally.
func buildVerticalTowers(
	allRefs []StructRef,
	infos map[string]*structInfo,
	edges []rawEdge,
	maxTowerHeight float64,
) [][]string {
	distinctOut := make(map[string][]string)
	distinctIn := make(map[string][]string)
	pairLinks := make(map[string]int)

	for _, e := range edges {
		pairKey := e.srcKey + "->" + e.dstKey
		if pairLinks[pairKey] == 0 {
			distinctOut[e.srcKey] = append(distinctOut[e.srcKey], e.dstKey)
			distinctIn[e.dstKey] = append(distinctIn[e.dstKey], e.srcKey)
		}
		pairLinks[pairKey]++
	}

	towers := make([][]string, len(allRefs))
	towerOf := make(map[string]int, len(allRefs))
	for i, ref := range allRefs {
		k := ref.Pkg + "." + ref.Name
		towers[i] = []string{k}
		towerOf[k] = i
	}

	activeTowerCount := func() int {
		n := 0
		for _, tw := range towers {
			if len(tw) > 0 {
				n++
			}
		}
		return n
	}

	towerEstHeight := func(tw []string) float64 {
		h := 0.0
		for i, k := range tw {
			h += infos[k].baseHeight
			if i > 0 {
				h += 64.0
			}
		}
		return h
	}

	tryMerge := func(u, v string, minRemaining int, heightCap float64) bool {
		if activeTowerCount() <= minRemaining {
			return false
		}
		tu := towerOf[u]
		tv := towerOf[v]
		if tu == tv || len(towers[tu]) == 0 || len(towers[tv]) == 0 {
			return false
		}
		if towers[tu][len(towers[tu])-1] != u || towers[tv][0] != v {
			return false
		}
		if towerEstHeight(towers[tu])+64.0+towerEstHeight(towers[tv]) > heightCap {
			return false
		}
		merged := append([]string{}, towers[tu]...)
		merged = append(merged, towers[tv]...)
		towers[tu] = merged
		towers[tv] = nil
		for _, k := range merged {
			towerOf[k] = tu
		}
		return true
	}

	// Pass 1: Always pair XSD *PropertyType wrappers directly above their single target *Type.
	for _, e := range edges {
		if !strings.HasSuffix(e.srcName, "PropertyType") {
			continue
		}
		if len(distinctOut[e.srcKey]) == 1 && len(distinctIn[e.dstKey]) == 1 && pairLinks[e.srcKey+"->"+e.dstKey] == 1 {
			tryMerge(e.srcKey, e.dstKey, 2, maxTowerHeight)
		}
	}

	// Pass 2: Chain 1-to-1 single-child spines (e.g. FeatureType -> TimeSlicePropertyType -> TimeSliceType,
	// or SurfaceType -> SurfacePatchArrayPropertyType -> PolygonPatchType -> AbstractRingPropertyType).
	if maxTowerHeight > 300.0 {
		for _, e := range edges {
			if len(distinctOut[e.srcKey]) == 1 && len(distinctIn[e.dstKey]) == 1 && pairLinks[e.srcKey+"->"+e.dstKey] == 1 {
				tryMerge(e.srcKey, e.dstKey, 3, maxTowerHeight)
			}
		}
	}

	var result [][]string
	for _, tw := range towers {
		if len(tw) > 0 {
			result = append(result, tw)
		}
	}
	return result
}

func chooseBestA3LandscapePlan(
	allRefs []StructRef,
	infos map[string]*structInfo,
	edges []rawEdge,
	enums []*enumInfo,
	gongNote *gong_models.GongNote,
) *candidatePlan {
	// Evaluate full spine merging (820px), medium 2-3 box towers (460px), and wrapper-only merging (290px)
	towerSets := [][][]string{
		buildVerticalTowers(allRefs, infos, edges, 820.0),
		buildVerticalTowers(allRefs, infos, edges, 460.0),
		buildVerticalTowers(allRefs, infos, edges, 290.0),
	}

	var best *candidatePlan

	for _, towers := range towerSets {
		numTowers := len(towers)
		if numTowers == 0 {
			continue
		}

		// Compute DAG ranks over towers
		tIdxOf := make(map[string]int)
		for ti, tw := range towers {
			for _, k := range tw {
				tIdxOf[k] = ti
			}
		}
		tOut := make([][]int, numTowers)
		tInDeg := make([]int, numTowers)
		seenTEdge := make(map[[2]int]bool)
		for _, e := range edges {
			tu := tIdxOf[e.srcKey]
			tv := tIdxOf[e.dstKey]
			if tu == tv || seenTEdge[[2]int{tu, tv}] {
				continue
			}
			seenTEdge[[2]int{tu, tv}] = true
			tOut[tu] = append(tOut[tu], tv)
			tInDeg[tv]++
		}

		rank := make([]int, numTowers)
		pinned := make([]bool, numTowers)
		hasRoot := false
		for ti := range numTowers {
			if tInDeg[ti] == 0 {
				rank[ti] = 0
				pinned[ti] = true
				hasRoot = true
			} else {
				rank[ti] = -1
			}
		}
		if !hasRoot {
			rank[0] = 0
			pinned[0] = true
		}
		for range numTowers + 1 {
			changed := false
			for u := range numTowers {
				if rank[u] < 0 {
					continue
				}
				for _, v := range tOut[u] {
					if pinned[v] {
						continue
					}
					nr := rank[u] + 1
					if nr > numTowers {
						nr = numTowers
					}
					if rank[v] < nr {
						rank[v] = nr
						changed = true
					}
				}
			}
			if !changed {
				break
			}
		}
		maxRank := 0
		for ti := range numTowers {
			if rank[ti] < 0 {
				rank[ti] = 0
			}
			if rank[ti] > maxRank {
				maxRank = rank[ti]
			}
		}

		for _, numCols := range []int{2, 3} {
			if numTowers < numCols {
				continue
			}
			assignments := generateColumnAssignments(towers, infos, rank, maxRank, numCols)
			for _, colAssign := range assignments {
				for noteColChoice := range numCols {
					plan := buildCandidatePlan(towers, colAssign, numCols, infos, edges, enums, gongNote, noteColChoice)
					if best == nil || plan.score < best.score {
						best = plan
					}
				}
			}
		}
	}

	if best == nil {
		towers := buildVerticalTowers(allRefs, infos, edges, 820.0)
		colAssign := make([]int, len(towers))
		best = buildCandidatePlan(towers, colAssign, 1, infos, edges, enums, gongNote, 0)
	}
	return best
}

func generateColumnAssignments(
	towers [][]string,
	infos map[string]*structInfo,
	rank []int,
	maxRank int,
	numCols int,
) [][]int {
	numTowers := len(towers)
	var results [][]int

	addUnique := func(assign []int) {
		used := make([]bool, numCols)
		for _, c := range assign {
			if c >= 0 && c < numCols {
				used[c] = true
			}
		}
		for c := range numCols {
			if !used[c] {
				return
			}
		}
		for _, ex := range results {
			same := true
			for i := range assign {
				if ex[i] != assign[i] {
					same = false
					break
				}
			}
			if same {
				return
			}
		}
		cp := make([]int, len(assign))
		copy(cp, assign)
		results = append(results, cp)
	}

	// Also generate height-balanced contiguous partitions over topologically ordered towers
	orderedTowers := make([]int, numTowers)
	for i := range numTowers {
		orderedTowers[i] = i
	}
	slices.SortStableFunc(orderedTowers, func(a, b int) int {
		if rank[a] != rank[b] {
			return rank[a] - rank[b]
		}
		return infos[towers[a][0]].origIdx - infos[towers[b][0]].origIdx
	})
	estH := make([]float64, numTowers)
	totalEstH := 0.0
	for idx, ti := range orderedTowers {
		h := 26.0
		for j, k := range towers[ti] {
			h += infos[k].baseHeight
			if j > 0 {
				h += 56.0
			}
		}
		estH[idx] = h
		totalEstH += h
	}
	addBalancedPartition := func(shifts []int) {
		cuts := make([]int, numCols-1)
		targetStep := totalEstH / float64(numCols)
		acc := 0.0
		curCut := 0
		for idx := 0; idx < numTowers-1 && curCut < numCols-1; idx++ {
			acc += estH[idx]
			if acc >= targetStep*float64(curCut+1) {
				cuts[curCut] = idx + 1
				curCut++
			}
		}
		for curCut < numCols-1 {
			cuts[curCut] = numTowers - (numCols - 1 - curCut)
			curCut++
		}
		for i := range cuts {
			if i < len(shifts) {
				cuts[i] += shifts[i]
			}
			minC := i + 1
			maxC := numTowers - (numCols - 1 - i)
			if cuts[i] < minC {
				cuts[i] = minC
			}
			if cuts[i] > maxC {
				cuts[i] = maxC
			}
			if i > 0 && cuts[i] <= cuts[i-1] {
				cuts[i] = cuts[i-1] + 1
			}
		}
		assign := make([]int, numTowers)
		col := 0
		for idx, ti := range orderedTowers {
			for col < numCols-1 && idx >= cuts[col] {
				col++
			}
			assign[ti] = col
		}
		addUnique(assign)
	}
	addBalancedPartition(nil)
	if numCols == 2 {
		addBalancedPartition([]int{-1})
		addBalancedPartition([]int{1})
	} else if numCols == 3 {
		for _, d1 := range []int{-1, 0, 1} {
			for _, d2 := range []int{-1, 0, 1} {
				addBalancedPartition([]int{d1, d2})
			}
		}
	} else if numCols == 4 {
		addBalancedPartition([]int{-1, 0, 0})
		addBalancedPartition([]int{1, 0, 0})
		addBalancedPartition([]int{0, 0, -1})
		addBalancedPartition([]int{0, 0, 1})
	}

	if maxRank == 0 {
		// Split towers evenly across numCols
		assign := make([]int, numTowers)
		for i := range numTowers {
			assign[i] = (i * numCols) / numTowers
		}
		addUnique(assign)
		return results
	}

	if maxRank+1 < numCols {
		// e.g., maxRank == 1 and numCols == 3: split largest rank across 2 columns
		assign := make([]int, numTowers)
		var rank1Idxs []int
		for i := range numTowers {
			if rank[i] == 0 {
				assign[i] = 0
			} else {
				rank1Idxs = append(rank1Idxs, i)
			}
		}
		if len(rank1Idxs) >= 2 {
			half := (len(rank1Idxs) + 1) / 2
			for idx, ti := range rank1Idxs {
				if idx < half {
					assign[ti] = 1
				} else {
					assign[ti] = 2
				}
			}
			addUnique(assign)
		}
		// Also try splitting rank 0 if rank 0 has >= 4 towers
		var rank0Idxs []int
		for i := range numTowers {
			if rank[i] == 0 {
				rank0Idxs = append(rank0Idxs, i)
			}
		}
		if len(rank0Idxs) >= 4 {
			assign2 := make([]int, numTowers)
			half := (len(rank0Idxs) + 1) / 2
			for idx, ti := range rank0Idxs {
				if idx < half {
					assign2[ti] = 0
				} else {
					assign2[ti] = 1
				}
			}
			for _, ti := range rank1Idxs {
				assign2[ti] = 2
			}
			addUnique(assign2)
		}
		return results
	}

	if maxRank+1 == numCols {
		assign := make([]int, numTowers)
		for i := range numTowers {
			assign[i] = rank[i]
		}
		addUnique(assign)
		return results
	}

	// maxRank + 1 > numCols: enumerate all monotonic cut points across 0..maxRank
	if numCols == 2 {
		for cut := 0; cut < maxRank; cut++ {
			assign := make([]int, numTowers)
			for i := range numTowers {
				if rank[i] <= cut {
					assign[i] = 0
				} else {
					assign[i] = 1
				}
			}
			addUnique(assign)
		}
	} else if numCols == 3 {
		for cut1 := 0; cut1 < maxRank-1; cut1++ {
			for cut2 := cut1 + 1; cut2 < maxRank; cut2++ {
				assign := make([]int, numTowers)
				for i := range numTowers {
					if rank[i] <= cut1 {
						assign[i] = 0
					} else if rank[i] <= cut2 {
						assign[i] = 1
					} else {
						assign[i] = 2
					}
				}
				addUnique(assign)
			}
		}
	}
	_ = infos
	return results
}

func sideRequiredHeight(numOut, numIn int) float64 {
	if numOut == 0 && numIn == 0 {
		return 0
	}
	h := 46.0
	if numOut > 0 {
		h += float64(numOut-1) * 32.0
	}
	if numIn > 0 {
		if numOut > 0 {
			h += 58.0
		}
		h += float64(numIn-1) * 58.0
	}
	return h
}

func assignSidePortRatios(box placedBox, outLinks, inLinks []*plannedLink, isStartForOut bool) {
	totalOut := len(outLinks)
	totalIn := len(inLinks)
	if totalOut == 0 && totalIn == 0 {
		return
	}
	if totalOut == 1 && totalIn == 0 {
		if isStartForOut {
			outLinks[0].startRatio = 0.5
		} else {
			outLinks[0].endRatio = 0.5
		}
		return
	}
	if totalOut == 0 && totalIn == 1 {
		inLinks[0].endRatio = 0.5
		return
	}

	// Compute exact Y offsets from box.y
	var yOffsets []float64
	cur := 23.0
	for i := range totalOut {
		if i > 0 {
			cur += 32.0
		}
		yOffsets = append(yOffsets, cur)
	}
	for i := range totalIn {
		if i == 0 && totalOut > 0 {
			cur += 58.0
		} else if i > 0 {
			cur += 58.0
		}
		yOffsets = append(yOffsets, cur)
	}

	spanUsed := yOffsets[len(yOffsets)-1] - yOffsets[0]
	extra := (box.height - 46.0) - spanUsed
	shift := 0.0
	if extra > 0 {
		shift = extra * 0.5
	}

	minR := math.Min(0.45, 20.0/box.height)
	maxR := math.Max(0.55, (box.height-20.0)/box.height)

	idx := 0
	for _, li := range outLinks {
		r := (yOffsets[idx] + shift) / box.height
		r = math.Max(minR, math.Min(maxR, r))
		if isStartForOut {
			li.startRatio = r
		}
		idx++
	}
	for _, li := range inLinks {
		r := (yOffsets[idx] + shift) / box.height
		r = math.Max(minR, math.Min(maxR, r))
		li.endRatio = r
		idx++
	}
}

func buildCandidatePlan(
	towers [][]string,
	towerCol []int,
	numCols int,
	infos map[string]*structInfo,
	edges []rawEdge,
	enums []*enumInfo,
	gongNote *gong_models.GongNote,
	noteColChoice int,
) *candidatePlan {
	// 1. Order towers within each column respecting intra-column edges and parent barycenter
	tIdxOf := make(map[string]int)
	for ti, tw := range towers {
		for _, k := range tw {
			tIdxOf[k] = ti
		}
	}

	colTowers := make([][]int, numCols)
	for ti := range towers {
		c := towerCol[ti]
		colTowers[c] = append(colTowers[c], ti)
	}

	towerRowScore := make(map[int]float64)
	for c := range numCols {
		tis := colTowers[c]
		// Compute barycenter from earlier columns
		bary := make(map[int]float64)
		for _, ti := range tis {
			sum := 0.0
			cnt := 0
			for _, e := range edges {
				if tIdxOf[e.dstKey] == ti {
					pu := tIdxOf[e.srcKey]
					if towerCol[pu] < c {
						sum += towerRowScore[pu]
						cnt++
					}
				}
			}
			if cnt > 0 {
				bary[ti] = sum / float64(cnt)
			} else {
				bary[ti] = float64(infos[towers[ti][0]].origIdx) * 0.1
			}
		}

		// Topological + barycentric sort within column c
		ordered := topoSortColumnTowers(tis, tIdxOf, edges, bary, infos, towers)
		colTowers[c] = ordered
		for idx, ti := range ordered {
			if len(ordered) <= 1 {
				towerRowScore[ti] = 0.5
			} else {
				towerRowScore[ti] = float64(idx) / float64(len(ordered)-1)
			}
		}
	}

	colStructs := make([][]string, numCols)
	colOfKey := make(map[string]int)
	rowOfKey := make(map[string]int)
	for c := range numCols {
		for _, ti := range colTowers[c] {
			for _, k := range towers[ti] {
				rowOfKey[k] = len(colStructs[c])
				colOfKey[k] = c
				colStructs[c] = append(colStructs[c], k)
			}
		}
	}

	// Filter redundant multi-column jump edges (vc > uc + 1) ONLY when target already has an adjacent-col or vertical parent?
	// Wait: preserve all semantic links unless exact duplicate! Let's check if keeping all edges works cleanly:
	// Wait: in original domain_test.go (line 660), if v in col c already had a parent in col c-1, redundant transitive shortcuts from col < c-1 were skipped? No, wait: e.g. AirspaceTimeSliceType -> NotePropertyType vs AirspaceGeometryComponentType -> NotePropertyType are DIFFERENT fields on DIFFERENT structs! Keep ALL edges!
	pairParallelCount := make(map[string]int)
	for _, e := range edges {
		pairParallelCount[e.srcKey+"->"+e.dstKey]++
	}

	var links []*plannedLink
	hasVertDownBetween := make(map[[2]string]string) // [upper, lower] -> fieldName

	for _, e := range edges {
		sc := colOfKey[e.srcKey]
		sr := rowOfKey[e.srcKey]
		dc := colOfKey[e.dstKey]
		dr := rowOfKey[e.dstKey]

		pl := &plannedLink{
			edge:   e,
			srcCol: sc,
			srcRow: sr,
			dstCol: dc,
			dstRow: dr,
		}
		if sc == dc && dr == sr+1 && pairParallelCount[e.srcKey+"->"+e.dstKey] == 1 {
			pl.kind = routeVerticalDown
			pl.startOrientation = doc_models.ORIENTATION_VERTICAL
			pl.endOrientation = doc_models.ORIENTATION_VERTICAL
			hasVertDownBetween[[2]string{e.srcKey, e.dstKey}] = e.fieldName
		} else if dc > sc {
			pl.kind = routeForwardHoriz
			pl.startOrientation = doc_models.ORIENTATION_HORIZONTAL
			pl.endOrientation = doc_models.ORIENTATION_HORIZONTAL
		} else if dc == sc {
			pl.kind = routeRightBus
			pl.startOrientation = doc_models.ORIENTATION_HORIZONTAL
			pl.endOrientation = doc_models.ORIENTATION_HORIZONTAL
		} else {
			pl.kind = routeBackwardHoriz
			pl.startOrientation = doc_models.ORIENTATION_HORIZONTAL
			pl.endOrientation = doc_models.ORIENTATION_HORIZONTAL
		}
		links = append(links, pl)
	}

	// 2. Compute exact box widths and port-aware heights
	leftOutByKey := make(map[string][]*plannedLink)
	leftInByKey := make(map[string][]*plannedLink)
	rightOutByKey := make(map[string][]*plannedLink)
	rightInByKey := make(map[string][]*plannedLink)

	for _, pl := range links {
		switch pl.kind {
		case routeForwardHoriz:
			rightOutByKey[pl.edge.srcKey] = append(rightOutByKey[pl.edge.srcKey], pl)
			leftInByKey[pl.edge.dstKey] = append(leftInByKey[pl.edge.dstKey], pl)
		case routeRightBus:
			rightOutByKey[pl.edge.srcKey] = append(rightOutByKey[pl.edge.srcKey], pl)
			rightInByKey[pl.edge.dstKey] = append(rightInByKey[pl.edge.dstKey], pl)
		case routeBackwardHoriz:
			leftOutByKey[pl.edge.srcKey] = append(leftOutByKey[pl.edge.srcKey], pl)
			rightInByKey[pl.edge.dstKey] = append(rightInByKey[pl.edge.dstKey], pl)
		}
	}

	boxes := make(map[string]placedBox, len(infos))
	colStructWidth := make([]float64, numCols)
	for c := range numCols {
		colStructWidth[c] = 220.0
		for _, k := range colStructs[c] {
			info := infos[k]
			w := info.minWidth
			h := math.Max(
				info.baseHeight,
				math.Max(
					sideRequiredHeight(len(leftOutByKey[k]), len(leftInByKey[k])),
					sideRequiredHeight(len(rightOutByKey[k]), len(rightInByKey[k])),
				),
			)
			boxes[k] = placedBox{width: w, height: math.Ceil(h)}
		}
	}

	// Ensure vertical-link pairs are wide enough for their field label + multiplicity
	for _, pl := range links {
		if pl.kind != routeVerticalDown {
			continue
		}
		reqW := math.Ceil(LinkLabelWidth(pl.edge.fieldName) + 78.0)
		sb := boxes[pl.edge.srcKey]
		if sb.width < reqW {
			sb.width = reqW
			boxes[pl.edge.srcKey] = sb
		}
		db := boxes[pl.edge.dstKey]
		if db.width < reqW {
			db.width = reqW
			boxes[pl.edge.dstKey] = db
		}
	}
	for c := range numCols {
		for _, k := range colStructs[c] {
			if boxes[k].width > colStructWidth[c] {
				colStructWidth[c] = boxes[k].width
			}
		}
	}

	// 3. Place Note, Structs, and Enums vertically per column
	const marginX = 36.0
	const topMarginY = 24.0
	const boxGapY = 26.0
	const vertLinkGapY = 56.0

	colStartY := make([]float64, numCols)
	for c := range numCols {
		colStartY[c] = topMarginY
	}

	var hasNote bool
	var noteBox placedBox
	noteCol := -1

	if gongNote != nil {
		hasNote = true
		bestC := noteColChoice
		if bestC < 0 || bestC >= numCols {
			bestC = 0
		}
		prefW := math.Min(560.0, math.Max(440.0, colStructWidth[bestC]))
		nw, nh := NoteBoxSize(gongNote.Body, prefW)
		noteBox = placedBox{y: topMarginY, width: nw, height: nh}
		colStartY[bestC] = topMarginY + nh + 28.0
		noteCol = bestC
		if nw > colStructWidth[bestC] {
			colStructWidth[bestC] = nw
		}
	}

	// Initial Y placement per column
	colBottomY := make([]float64, numCols)
	for c := range numCols {
		curY := colStartY[c]
		for i, k := range colStructs[c] {
			if i > 0 {
				prev := colStructs[c][i-1]
				if _, ok := hasVertDownBetween[[2]string{prev, k}]; ok {
					curY += vertLinkGapY
				} else {
					curY += boxGapY
				}
			}
			b := boxes[k]
			b.y = math.Round(curY)
			boxes[k] = b
			curY = b.y + b.height
		}
		colBottomY[c] = curY
	}

	// Vertically center shorter columns slightly toward connected parents (while respecting vertical chains!)
	maxStructBottom := 0.0
	for c := range numCols {
		if colBottomY[c] > maxStructBottom {
			maxStructBottom = colBottomY[c]
		}
	}
	for c := 1; c < numCols; c++ {
		if len(colStructs[c]) == 0 {
			continue
		}
		// Compute desired shift to align with horizontal parents in c-1
		var deltas []float64
		for _, pl := range links {
			if pl.kind == routeForwardHoriz && pl.srcCol == c-1 && pl.dstCol == c {
				sb := boxes[pl.edge.srcKey]
				db := boxes[pl.edge.dstKey]
				deltas = append(deltas, (sb.y+0.5*sb.height)-(db.y+0.5*db.height))
			}
		}
		if len(deltas) > 0 {
			sum := 0.0
			for _, d := range deltas {
				sum += d
			}
			avgShift := sum / float64(len(deltas))
			enumReserve := 0.0
			if len(enums) > 0 {
				enumReserve = 185.0
			}
			maxDown := math.Max(0, maxStructBottom-enumReserve-colBottomY[c])
			shift := math.Max(0, math.Min(maxDown, math.Round(avgShift*0.65)))
			if shift > 0 {
				for _, k := range colStructs[c] {
					b := boxes[k]
					b.y += shift
					boxes[k] = b
				}
				colBottomY[c] += shift
			}
		}
	}

	// If any multi-column jump link (dstCol > srcCol + 1) exists, open a horizontal channel in intermediate columns
	for _, pl := range links {
		if pl.kind == routeForwardHoriz && pl.dstCol > pl.srcCol+1 {
			db := boxes[pl.edge.dstKey]
			targetY := db.y + 0.5*db.height
			for midC := pl.srcCol + 1; midC < pl.dstCol; midC++ {
				for idx, k := range colStructs[midC] {
					mb := boxes[k]
					if targetY >= mb.y-28.0 && targetY <= mb.y+mb.height+28.0 {
						// Push mb and all subsequent boxes in midC down below targetY + 34, unless broken vertical link
						push := (targetY + 36.0) - mb.y
						if push > 0 {
							for j := idx; j < len(colStructs[midC]); j++ {
								kj := colStructs[midC][j]
								bj := boxes[kj]
								bj.y += math.Ceil(push)
								boxes[kj] = bj
							}
							colBottomY[midC] += math.Ceil(push)
						}
						break
					}
				}
			}
		}
	}

	// Place Enums into shortest columns so vertical height stays compact and balanced
	enumBoxes := make([]placedBox, len(enums))
	enumCol := make([]int, len(enums))
	for ei, en := range enums {
		bestC := 0
		for c := 1; c < numCols; c++ {
			if colBottomY[c] < colBottomY[bestC] {
				bestC = c
			}
		}
		ey := math.Round(colBottomY[bestC] + 34.0)
		if colBottomY[bestC] <= topMarginY {
			ey = topMarginY
		}
		eb := placedBox{y: ey, width: en.width, height: en.height}
		enumBoxes[ei] = eb
		enumCol[ei] = bestC
		colBottomY[bestC] = ey + en.height
		if en.width > colStructWidth[bestC] {
			colStructWidth[bestC] = en.width
		}
	}

	// If noteHeaderMode && noteBox.width extends past col 0 into col 1, ensure col 1 starts below noteBox
	// (we re-check once X positions are computed below).

	// 4. Sort side ports and assign StartRatio / EndRatio deterministically
	var allKeysSorted []string
	for c := range numCols {
		allKeysSorted = append(allKeysSorted, colStructs[c]...)
	}
	for _, key := range allKeysSorted {
		b, ok := boxes[key]
		if !ok {
			continue
		}
		sortByOppositeY := func(list []*plannedLink, useDst bool) {
			slices.SortStableFunc(list, func(a, bLink *plannedLink) int {
				var ay, by float64
				if useDst {
					ab := boxes[a.edge.dstKey]
					bb := boxes[bLink.edge.dstKey]
					ay = ab.y + 0.5*ab.height
					by = bb.y + 0.5*bb.height
				} else {
					ab := boxes[a.edge.srcKey]
					bb := boxes[bLink.edge.srcKey]
					ay = ab.y + a.startRatio*ab.height
					by = bb.y + bLink.startRatio*bb.height
				}
				if math.Abs(ay-by) > 1.0 {
					if ay < by {
						return -1
					}
					return 1
				}
				return strings.Compare(a.edge.fieldName, bLink.edge.fieldName)
			})
		}

		rOut := rightOutByKey[key]
		lOut := leftOutByKey[key]
		sortByOppositeY(rOut, true)
		sortByOppositeY(lOut, true)
		assignSidePortRatios(b, rOut, nil, true)
		assignSidePortRatios(b, lOut, nil, true)
	}

	// Second pass: sort incoming ports using finalized source startRatios, then assign combined (out + in) ratios
	for _, key := range allKeysSorted {
		b, ok := boxes[key]
		if !ok {
			continue
		}
		sortIn := func(list []*plannedLink) {
			slices.SortStableFunc(list, func(a, bLink *plannedLink) int {
				ab := boxes[a.edge.srcKey]
				bb := boxes[bLink.edge.srcKey]
				ay := ab.y + a.startRatio*ab.height
				by := bb.y + bLink.startRatio*bb.height
				if math.Abs(ay-by) > 1.0 {
					if ay < by {
						return -1
					}
					return 1
				}
				return strings.Compare(a.edge.fieldName, bLink.edge.fieldName)
			})
		}
		rOut := rightOutByKey[key]
		rIn := rightInByKey[key]
		lOut := leftOutByKey[key]
		lIn := leftInByKey[key]
		sortIn(rIn)
		sortIn(lIn)
		assignSidePortRatios(b, rOut, rIn, true)
		assignSidePortRatios(b, lOut, lIn, true)
	}

	// 5. Compute corridor widths and vertical bend tracks X per corridor c
	corridorSpanX := make([]float64, numCols)
	for c := range numCols {
		corridorSpanX[c] = colStructWidth[c] + 84.0
	}

	corridorLinks := make(map[int][]*plannedLink)
	for _, pl := range links {
		switch pl.kind {
		case routeForwardHoriz, routeRightBus:
			corridorLinks[pl.srcCol] = append(corridorLinks[pl.srcCol], pl)
		case routeBackwardHoriz:
			corridorLinks[pl.srcCol-1] = append(corridorLinks[pl.srcCol-1], pl)
		}
	}

	// Helper returning max relative Right of any box in column c overlapping [minY, maxY]
	maxRelRightInCol := func(c int, minY, maxY float64) float64 {
		mr := 0.0
		for _, k := range colStructs[c] {
			b := boxes[k]
			if b.y+b.height+28.0 >= minY && b.y-28.0 <= maxY {
				if b.width > mr {
					mr = b.width
				}
			}
		}
		for ei, eb := range enumBoxes {
			if enumCol[ei] == c && eb.y+eb.height+28.0 >= minY && eb.y-28.0 <= maxY {
				if eb.width > mr {
					mr = eb.width
				}
			}
		}
		if hasNote && noteCol == c && noteBox.y+noteBox.height+28.0 >= minY && noteBox.y-28.0 <= maxY {
			if noteBox.width > mr {
				mr = noteBox.width
			}
		}
		return mr
	}

	for c := range numCols {
		cLinks := corridorLinks[c]
		if len(cLinks) == 0 {
			continue
		}
		type spanItem struct {
			pl   *plannedLink
			minY float64
			maxY float64
		}
		var spans []spanItem
		for _, pl := range cLinks {
			sb := boxes[pl.edge.srcKey]
			db := boxes[pl.edge.dstKey]
			ys := sb.y + pl.startRatio*sb.height
			ye := db.y + pl.endRatio*db.height
			spans = append(spans, spanItem{
				pl:   pl,
				minY: math.Min(ys, ye) - 26.0,
				maxY: math.Max(ys, ye) + 26.0,
			})
		}
		slices.SortStableFunc(spans, func(a, b spanItem) int {
			if a.minY < b.minY {
				return -1
			}
			if a.minY > b.minY {
				return 1
			}
			return 0
		})

		type compGroup struct {
			items []*plannedLink
			minY  float64
			maxY  float64
		}
		var comps []compGroup
		for _, sp := range spans {
			if len(comps) == 0 || sp.minY > comps[len(comps)-1].maxY {
				comps = append(comps, compGroup{items: []*plannedLink{sp.pl}, minY: sp.minY, maxY: sp.maxY})
			} else {
				last := &comps[len(comps)-1]
				last.items = append(last.items, sp.pl)
				if sp.maxY > last.maxY {
					last.maxY = sp.maxY
				}
			}
		}

		for _, comp := range comps {
			compRight := maxRelRightInCol(c, comp.minY, comp.maxY)
			// Ensure right-entering labels (from routeRightBus or routeBackwardHoriz into col c) fit to the left of bend!
			for _, pl := range comp.items {
				if pl.kind == routeRightBus || (pl.kind == routeBackwardHoriz && pl.dstCol == c) {
					db := boxes[pl.edge.dstKey]
					need := db.width + LinkLabelWidth(pl.edge.fieldName) + 32.0
					if need > compRight {
						compRight = need
					}
				}
			}

			var upL, flatL, downL []*plannedLink
			for _, pl := range comp.items {
				sb := boxes[pl.edge.srcKey]
				db := boxes[pl.edge.dstKey]
				ys := sb.y + pl.startRatio*sb.height
				ye := db.y + pl.endRatio*db.height
				if ye < ys-4.0 {
					upL = append(upL, pl)
				} else if ye > ys+4.0 {
					downL = append(downL, pl)
				} else {
					flatL = append(flatL, pl)
				}
			}
			slices.SortStableFunc(upL, func(a, b *plannedLink) int {
				ay := boxes[a.edge.srcKey].y + a.startRatio*boxes[a.edge.srcKey].height
				by := boxes[b.edge.srcKey].y + b.startRatio*boxes[b.edge.srcKey].height
				if ay < by {
					return -1
				}
				if ay > by {
					return 1
				}
				return 0
			})
			slices.SortStableFunc(downL, func(a, b *plannedLink) int {
				ay := boxes[a.edge.srcKey].y + a.startRatio*boxes[a.edge.srcKey].height
				by := boxes[b.edge.srcKey].y + b.startRatio*boxes[b.edge.srcKey].height
				if ay > by {
					return -1
				}
				if ay < by {
					return 1
				}
				return 0
			})

			var ordered []*plannedLink
			ordered = append(ordered, upL...)
			ordered = append(ordered, flatL...)
			ordered = append(ordered, downL...)

			const baseMargin = 32.0
			const trackStep = 18.0
			for idx, pl := range ordered {
				pl.relCornerX = compRight + baseMargin + float64(idx)*trackStep
				reqSpan := pl.relCornerX + 42.0
				if pl.kind == routeForwardHoriz && pl.dstCol == c+1 {
					reqSpan = pl.relCornerX + LinkLabelWidth(pl.edge.fieldName) + 36.0
				}
				if reqSpan > corridorSpanX[c] {
					corridorSpanX[c] = reqSpan
				}
			}
		}
	}

	// 6. Finalize X coordinates for columns, boxes, enums, note, and link ratios
	colLeft := make([]float64, numCols)
	curX := marginX
	for c := range numCols {
		colLeft[c] = math.Round(curX)
		for _, k := range colStructs[c] {
			b := boxes[k]
			b.x = colLeft[c]
			boxes[k] = b
		}
		curX += corridorSpanX[c]
	}
	for ei := range enumBoxes {
		enumBoxes[ei].x = colLeft[enumCol[ei]]
	}
	if hasNote {
		noteBox.x = colLeft[noteCol]
	}

	for _, pl := range links {
		sb := boxes[pl.edge.srcKey]
		db := boxes[pl.edge.dstKey]
		switch pl.kind {
		case routeVerticalDown:
			minRight := math.Min(sb.x+sb.width, db.x+db.width)
			arrowX := minRight - 54.0
			pl.startRatio = (arrowX - sb.x) / sb.width
			pl.endRatio = (arrowX - db.x) / db.width
			midY := 0.5 * (sb.y + sb.height + db.y)
			pl.cornerOffsetRatio = (midY - sb.y) / sb.height
		case routeForwardHoriz, routeRightBus:
			cornerX := colLeft[pl.srcCol] + pl.relCornerX
			pl.cornerOffsetRatio = (cornerX - sb.x) / sb.width
		case routeBackwardHoriz:
			cornerX := colLeft[pl.srcCol-1] + pl.relCornerX
			pl.cornerOffsetRatio = (cornerX - sb.x) / sb.width
		}
	}

	// Choose clean unobstructed GongNoteLinkShape target (if any referenced struct has a clear straight line)
	noteTargetKey := ""
	if hasNote && gongNote != nil {
		ncx := noteBox.x + 0.5*noteBox.width
		ncy := noteBox.y + 0.5*noteBox.height
		bestDist := math.MaxFloat64
		for _, gl := range gongNote.Links {
			if gl.Recv != "" {
				continue
			}
			for _, k := range allKeysSorted {
				b := boxes[k]
				if infos[k].gs.Name != gl.Name {
					continue
				}
				tcx := b.x + 0.5*b.width
				tcy := b.y + 0.5*b.height
				obstructed := false
				for okKey, ob := range boxes {
					if okKey == k {
						continue
					}
					if lineIntersectsRect(ncx, ncy, tcx, tcy, Rect2D{X1: ob.x, Y1: ob.y, X2: ob.x + ob.width, Y2: ob.y + ob.height}) {
						obstructed = true
						break
					}
				}
				if !obstructed {
					for _, eb := range enumBoxes {
						if lineIntersectsRect(ncx, ncy, tcx, tcy, Rect2D{X1: eb.x, Y1: eb.y, X2: eb.x + eb.width, Y2: eb.y + eb.height}) {
							obstructed = true
							break
						}
					}
				}
				if !obstructed {
					dist := math.Hypot(tcx-ncx, tcy-ncy)
					if dist < bestDist {
						bestDist = dist
						noteTargetKey = gl.Name
					}
				}
			}
		}
	}

	plan := &candidatePlan{
		numCols:       numCols,
		colStructs:    colStructs,
		boxes:         boxes,
		enums:         enumBoxes,
		hasNote:       hasNote,
		noteBox:       noteBox,
		noteTargetKey: noteTargetKey,
		links:         links,
	}

	// Score candidate using exact geometric verification on a scratch stage
	scratchStage := doc_models.NewStage("scratch")
	scratchCD := materializePlan(scratchStage, DiagramSpec{Name: "scratch"}, 0, infos, enums, gongNote, plan)
	dm := AnalyzeClassdiagram(scratchCD)
	plan.metrics = dm

	// Target: Small A3 Landscape (~1587 x 1123 px, ratio sqrt(2) = 1.414)
	score := 0.0
	if dm.Height > 0 {
		score += math.Abs(dm.AspectRatio-1.46) * 340.0
	}
	score += (dm.Width*0.08 + dm.Height*0.18)
	if dm.Width > 1720.0 {
		score += (dm.Width - 1720.0) * 4.5
	}
	if dm.Width > 1880.0 {
		score += (dm.Width - 1880.0) * 14.0
	}
	if dm.Height > 1080.0 {
		score += (dm.Height - 1080.0) * 5.5
	}
	if dm.Height > 1220.0 {
		score += (dm.Height - 1220.0) * 16.0
	}
	if dm.Width < 1050.0 && len(infos) >= 5 {
		score += (1050.0 - dm.Width) * 3.0
	}

	for _, pl := range links {
		switch pl.kind {
		case routeForwardHoriz:
			if pl.dstCol > pl.srcCol+1 {
				score += 240.0
			} else {
				score += 8.0
			}
		case routeRightBus:
			score += 95.0
		case routeBackwardHoriz:
			score += 220.0
		}
	}

	collisionCount := len(dm.BoxOverlaps) + len(dm.TextClipIssues) + len(dm.SegmentPiercings) +
		len(dm.LabelBoxHits) + len(dm.LabelLabelHits) + len(dm.NoteLinkHits)
	score += float64(collisionCount) * 8000.0

	plan.score = score
	return plan
}

func topoSortColumnTowers(
	tis []int,
	tIdxOf map[string]int,
	edges []rawEdge,
	bary map[int]float64,
	infos map[string]*structInfo,
	towers [][]string,
) []int {
	inSet := make(map[int]bool, len(tis))
	for _, ti := range tis {
		inSet[ti] = true
	}
	adj := make(map[int][]int)
	inDeg := make(map[int]int)
	seen := make(map[[2]int]bool)
	for _, e := range edges {
		u := tIdxOf[e.srcKey]
		v := tIdxOf[e.dstKey]
		if u != v && inSet[u] && inSet[v] && !seen[[2]int{u, v}] {
			seen[[2]int{u, v}] = true
			adj[u] = append(adj[u], v)
			inDeg[v]++
		}
	}

	cmpTower := func(a, b int) int {
		if math.Abs(bary[a]-bary[b]) > 1e-6 {
			if bary[a] < bary[b] {
				return -1
			}
			return 1
		}
		return infos[towers[a][0]].origIdx - infos[towers[b][0]].origIdx
	}

	var ready []int
	for _, ti := range tis {
		if inDeg[ti] == 0 {
			ready = append(ready, ti)
		}
	}
	slices.SortStableFunc(ready, cmpTower)

	var out []int
	placed := make(map[int]bool)
	for len(ready) > 0 {
		u := ready[0]
		ready = ready[1:]
		if placed[u] {
			continue
		}
		placed[u] = true
		out = append(out, u)
		for _, v := range adj[u] {
			inDeg[v]--
			if inDeg[v] == 0 {
				ready = append(ready, v)
				slices.SortStableFunc(ready, cmpTower)
			}
		}
	}
	if len(out) < len(tis) {
		var rem []int
		for _, ti := range tis {
			if !placed[ti] {
				rem = append(rem, ti)
			}
		}
		slices.SortStableFunc(rem, cmpTower)
		out = append(out, rem...)
	}
	return out
}

func materializePlan(
	docStage *doc_models.Stage,
	spec DiagramSpec,
	diagIdx int,
	infos map[string]*structInfo,
	enums []*enumInfo,
	gongNote *gong_models.GongNote,
	plan *candidatePlan,
) *doc_models.Classdiagram {
	cd := (&doc_models.Classdiagram{
		Name:                      spec.Name,
		Description:               spec.Description,
		ShowNbInstances:           true,
		ShowMultiplicity:          true,
		ShowLinkNames:             true,
		IsExpanded:                diagIdx == 0,
		NodeGongStructsIsExpanded: false,
		NodeGongEnumsIsExpanded:   false,
		NodeGongNotesIsExpanded:   true,
		NodeGongNoteNodeExpansion: `[true]`,
	}).Stage(docStage)

	if plan.hasNote && gongNote != nil {
		noteShape := (&doc_models.GongNoteShape{
			Name:       cd.Name + "-" + gongNote.Name,
			Identifier: doc_models.GongNoteNameToIdentifier(gongNote.Name),
			Body:       gongNote.Body,
			BodyHTML:   gongNote.BodyHTML,
			X:          plan.noteBox.x,
			Y:          plan.noteBox.y,
			Width:      plan.noteBox.width,
			Height:     plan.noteBox.height,
		}).Stage(docStage)
		cd.GongNoteShapes = append(cd.GongNoteShapes, noteShape)

		if plan.noteTargetKey != "" {
			nls := (&doc_models.GongNoteLinkShape{
				Name:       cd.Name + "-" + noteShape.Name + "-" + plan.noteTargetKey,
				Identifier: doc_models.GongStructNameToIdentifier(plan.noteTargetKey),
				Type:       doc_models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE,
				IsHidden:   false,
			}).Stage(docStage)
			noteShape.GongNoteLinkShapes = append(noteShape.GongNoteLinkShapes, nls)
		}
	}

	shapeByKey := make(map[string]*doc_models.GongStructShape, len(infos))
	for c := range plan.numCols {
		for _, key := range plan.colStructs[c] {
			info := infos[key]
			b := plan.boxes[key]
			var attrShapes []*doc_models.AttributeShape
			for i, aName := range info.attrNames {
				as := (&doc_models.AttributeShape{
					Name:           aName,
					IdentifierMeta: fmt.Sprintf("ref_%s.%s{}.%s", info.ref.Pkg, info.gs.Name, aName),
					Structname:     info.gs.Name,
					Fieldtypename:  info.attrTypes[i],
					IsHidden:       false,
				}).Stage(docStage)
				attrShapes = append(attrShapes, as)
			}
			gss := (&doc_models.GongStructShape{
				Name:            cd.Name + "-" + info.gs.Name,
				IdentifierMeta:  fmt.Sprintf("ref_%s.%s{}", info.ref.Pkg, info.gs.Name),
				AttributeShapes: attrShapes,
				X:               b.x,
				Y:               b.y,
				Width:           b.width,
				Height:          b.height,
			}).Stage(docStage)
			cd.GongStructShapes = append(cd.GongStructShapes, gss)
			shapeByKey[key] = gss
		}
	}

	for ei, en := range enums {
		eb := plan.enums[ei]
		var valShapes []*doc_models.GongEnumValueShape
		for _, vName := range en.valNames {
			evs := (&doc_models.GongEnumValueShape{
				Name:           vName,
				IdentifierMeta: fmt.Sprintf("ref_%s.%s", en.ref.Pkg, vName),
				IsHidden:       false,
			}).Stage(docStage)
			valShapes = append(valShapes, evs)
		}
		ges := (&doc_models.GongEnumShape{
			Name:                cd.Name + "-" + en.ge.Name,
			IdentifierMeta:      fmt.Sprintf("new(ref_%s.%s)", en.ref.Pkg, en.ge.Name),
			GongEnumValueShapes: valShapes,
			X:                   eb.x,
			Y:                   eb.y,
			Width:               eb.width,
			Height:              eb.height,
		}).Stage(docStage)
		cd.GongEnumShapes = append(cd.GongEnumShapes, ges)
	}

	for _, pl := range plan.links {
		srcShape := shapeByKey[pl.edge.srcKey]
		dstShape := shapeByKey[pl.edge.dstKey]
		if srcShape == nil || dstShape == nil {
			continue
		}
		ls := (&doc_models.LinkShape{
			Name:                      pl.edge.fieldName,
			IdentifierMeta:            fmt.Sprintf("ref_%s.%s{}.%s", pl.edge.srcPkg, pl.edge.srcName, pl.edge.fieldName),
			FieldTypeIdentifierMeta:   fmt.Sprintf("ref_%s.%s{}", pl.edge.dstPkg, pl.edge.dstName),
			SourceMultiplicity:        doc_models.MANY,
			TargetMultiplicity:        pl.edge.mult,
			StartOrientation:          pl.startOrientation,
			StartRatio:                math.Round(pl.startRatio*10000) / 10000,
			EndOrientation:            pl.endOrientation,
			EndRatio:                  math.Round(pl.endRatio*10000) / 10000,
			CornerOffsetRatio:         math.Round(pl.cornerOffsetRatio*10000) / 10000,
			FieldOffsetX:              0,
			FieldOffsetY:              0,
			TargetMultiplicityOffsetX: 0,
			TargetMultiplicityOffsetY: 0,
			SourceMultiplicityOffsetX: 0,
			SourceMultiplicityOffsetY: 0,
			X:                         math.Round((srcShape.X+dstShape.X)*0.5*10) / 10,
			Y:                         math.Round((srcShape.Y+dstShape.Y)*0.5*10) / 10,
			IsHidden:                  false,
		}).Stage(docStage)
		srcShape.LinkShapes = append(srcShape.LinkShapes, ls)
	}

	return cd
}
