package models_test

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	embeddedgo "github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models"
	gong_models "github.com/fullstack-lang/gong/go/models"
	doc_models "github.com/fullstack-lang/gong/lib/doc/go/models"
	"github.com/fullstack-lang/gong/lib/doc/go/prepare"
	splitlite "github.com/fullstack-lang/gong/lib/splitlite/go/models"
)

const (
	domainName         = "airport"
	domainImportPrefix = "github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models"
)

func TestRoundTripAndPopulateStage(t *testing.T) {
	xmlPath := filepath.Join("..", "..", "sample_"+domainName+".xml")
	xmlBytes, err := os.ReadFile(xmlPath)
	if err != nil {
		t.Fatalf("failed reading %s: %v", xmlPath, err)
	}

	stage := models.NewStage(domainName)
	if err := models.UnmarshallXMLToStage(stage, xmlBytes); err != nil {
		t.Fatalf("UnmarshallXMLToStage failed: %v", err)
	}

	stage.ComputeReferenceAndOrders()
	goCode, err := stage.MarshallToString("", "main")
	if err != nil {
		t.Fatalf("MarshallToString failed: %v", err)
	}

	stageGoPath := filepath.Join("..", "cmd", domainName, "data", "stage.go")
	if err := os.MkdirAll(filepath.Dir(stageGoPath), 0755); err != nil {
		t.Fatalf("failed creating data dir: %v", err)
	}
	if err := os.WriteFile(stageGoPath, []byte(goCode), 0644); err != nil {
		t.Fatalf("failed writing %s: %v", stageGoPath, err)
	}

	stage2 := models.NewStage(domainName + "_ast")
	if err := stage2.ParseAstString(goCode, true); err != nil {
		t.Fatalf("ParseAstString failed: %v", err)
	}

	xmlOut, err := models.MarshallStageToXML(stage2)
	if err != nil {
		t.Fatalf("MarshallStageToXML failed: %v", err)
	}

	stage3 := models.NewStage(domainName + "_xml")
	if err := models.UnmarshallXMLToStage(stage3, xmlOut); err != nil {
		t.Fatalf("round-trip UnmarshallXMLToStage failed: %v", err)
	}
}

type structRef struct {
	Pkg  string
	Name string
}

type diagramSpec struct {
	Name        string
	Description string
	Features    []string
	Objects     []string
	ExtraCols   [5][]structRef
	Enums       []structRef
}

func TestPopulateDomainUMLDiagrams(t *testing.T) {
	mux := http.NewServeMux()
	diagramEditor := &splitlite.AsSplitArea{
		Name:             "Bottom",
		ShowNameInHeader: false,
		Size:             50,
	}
	metaPackageImports := []*doc_models.MetaPackageImport{
		{Alias: "ref_models", Path: fmt.Sprintf("%q", domainImportPrefix)},
		{Alias: "ref_aixmfeatures", Path: fmt.Sprintf("%q", domainImportPrefix+"/aixmfeatures")},
		{Alias: "ref_aixmdatatypes", Path: fmt.Sprintf("%q", domainImportPrefix+"/aixmdatatypes")},
		{Alias: "ref_aixmabstractgmlobjecttypes", Path: fmt.Sprintf("%q", domainImportPrefix+"/aixmabstractgmlobjecttypes")},
		{Alias: "ref_gml", Path: fmt.Sprintf("%q", domainImportPrefix+"/gml")},
		{Alias: "ref_xlink", Path: fmt.Sprintf("%q", domainImportPrefix+"/xlink")},
	}
	docStager := prepare.PrepareStageSetSplitlite(
		mux,
		true,
		domainImportPrefix+":diag_gen",
		metaPackageImports,
		embeddedgo.GoModelsDir,
		embeddedgo.GoDiagramsDir,
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

	docStage.Reset()

	diagramPkg := (&doc_models.DiagramPackage{
		Name: "AIXM 5.1.1 UML Navigator Diagrams - " + domainName,
	}).Stage(docStage)

	specs := buildDomainDiagramSpecs()

	for diagIdx, spec := range specs {
		cd := (&doc_models.Classdiagram{
			Name:                      spec.Name,
			Description:               spec.Description,
			ShowNbInstances:           true,
			ShowMultiplicity:          true,
			ShowLinkNames:             true,
			IsExpanded:                diagIdx == 0,
			NodeGongStructsIsExpanded: false,
			NodeGongEnumsIsExpanded:   false,
			NodeGongNotesIsExpanded:   false,
		}).Stage(docStage)

		diagramPkg.Classdiagrams = append(diagramPkg.Classdiagrams, cd)
		if diagIdx == 0 {
			diagramPkg.SelectedClassdiagram = cd
		}

		var allRefs []structRef
		hintCol := make(map[string]int)
		seen := make(map[string]bool)
		addRef := func(col int, ref structRef) {
			key := ref.Pkg + "." + ref.Name
			if seen[key] {
				return
			}
			if _, ok := structsByPkgName[key]; !ok {
				return
			}
			seen[key] = true
			hintCol[key] = col
			allRefs = append(allRefs, ref)
		}

		for _, feat := range spec.Features {
			base := strings.TrimPrefix(feat, "Abstract")
			addRef(0, structRef{Pkg: "aixmfeatures", Name: feat + "Type"})
			addRef(1, structRef{Pkg: "aixmfeatures", Name: base + "TimeSlicePropertyType"})
			addRef(2, structRef{Pkg: "aixmfeatures", Name: base + "TimeSliceType"})
		}

		for _, obj := range spec.Objects {
			base := strings.TrimPrefix(obj, "Abstract")
			addRef(3, structRef{Pkg: "aixmfeatures", Name: base + "PropertyType"})
			addRef(4, structRef{Pkg: "aixmfeatures", Name: obj + "Type"})
		}

		for c := range 5 {
			for _, ref := range spec.ExtraCols[c] {
				addRef(c, ref)
			}
		}

		if len(spec.Features) > 0 || len(spec.Objects) > 0 {
			dtCount := 0
			const maxAutoDataTypes = 4
			for _, srcRef := range allRefs {
				if hintCol[srcRef.Pkg+"."+srcRef.Name] != 2 && hintCol[srcRef.Pkg+"."+srcRef.Name] != 4 {
					continue
				}
				if dtCount >= maxAutoDataTypes {
					break
				}
				gs := structsByPkgName[srcRef.Pkg+"."+srcRef.Name]
				if gs == nil {
					continue
				}
				for _, f := range gs.Fields {
					if dtCount >= maxAutoDataTypes {
						break
					}
					var targetGS *gong_models.GongStruct
					switch rf := f.(type) {
					case *gong_models.PointerToGongStructField:
						targetGS = rf.GongStruct
					case *gong_models.SliceOfPointerToGongStructField:
						targetGS = rf.GongStruct
					}
					if targetGS != nil && targetGS.ModelPkg != nil && targetGS.ModelPkg.PkgGoName == "aixmdatatypes" {
						dtRef := structRef{Pkg: "aixmdatatypes", Name: targetGS.Name}
						key := dtRef.Pkg + "." + dtRef.Name
						if !seen[key] {
							addRef(3, dtRef)
							dtCount++
						}
					}
				}
			}
		}

		// Build directed graph among included structs and count links per struct
		outEdges := make(map[string][]string)
		inEdges := make(map[string][]string)
		outLinkCount := make(map[string]int)
		inLinkCount := make(map[string]int)

		for _, srcRef := range allRefs {
			srcKey := srcRef.Pkg + "." + srcRef.Name
			srcGS := structsByPkgName[srcKey]
			if srcGS == nil {
				continue
			}
			pairCount := make(map[string]int)
			for _, field := range srcGS.Fields {
				var targetGS *gong_models.GongStruct
				switch rf := field.(type) {
				case *gong_models.PointerToGongStructField:
					targetGS = rf.GongStruct
				case *gong_models.SliceOfPointerToGongStructField:
					targetGS = rf.GongStruct
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
				if pairCount[tKey] >= 1 {
					continue
				}
				pairCount[tKey]++
				outEdges[srcKey] = append(outEdges[srcKey], tKey)
				inEdges[tKey] = append(inEdges[tKey], srcKey)
				outLinkCount[srcKey]++
				inLinkCount[tKey]++
			}
		}

		// Compute topological column for each struct so links flow left-to-right
		rawCol := make(map[string]int)
		pinnedRoot := make(map[string]bool)
		for _, ref := range allRefs {
			key := ref.Pkg + "." + ref.Name
			if len(inEdges[key]) == 0 {
				rawCol[key] = 0
				pinnedRoot[key] = true
			} else {
				rawCol[key] = -1
			}
		}
		if len(pinnedRoot) == 0 && len(allRefs) > 0 {
			firstKey := allRefs[0].Pkg + "." + allRefs[0].Name
			rawCol[firstKey] = 0
			pinnedRoot[firstKey] = true
		}

		for range len(allRefs) + 1 {
			changed := false
			for _, ref := range allRefs {
				u := ref.Pkg + "." + ref.Name
				if rawCol[u] < 0 {
					continue
				}
				for _, v := range outEdges[u] {
					if pinnedRoot[v] {
						continue
					}
					nextCol := rawCol[u] + 1
					if nextCol > 8 {
						nextCol = 8
					}
					if rawCol[v] < nextCol {
						rawCol[v] = nextCol
						changed = true
					}
				}
			}
			if !changed {
				break
			}
		}
		for _, ref := range allRefs {
			key := ref.Pkg + "." + ref.Name
			if rawCol[key] < 0 {
				rawCol[key] = hintCol[key]
			}
		}

		// Group into compacted columns 0..numCols-1
		var usedColNums []int
		colPresent := make(map[int]bool)
		for _, ref := range allRefs {
			c := rawCol[ref.Pkg+"."+ref.Name]
			if !colPresent[c] {
				colPresent[c] = true
				usedColNums = append(usedColNums, c)
			}
		}
		slices.Sort(usedColNums)
		colRemap := make(map[int]int)
		for idx, c := range usedColNums {
			colRemap[c] = idx
		}

		numCols := len(usedColNums)
		cols := make([][]structRef, numCols)
		origOrder := make(map[string]int)
		for idx, ref := range allRefs {
			key := ref.Pkg + "." + ref.Name
			origOrder[key] = idx
			c := colRemap[rawCol[key]]
			cols[c] = append(cols[c], ref)
		}

		// Barycenter row ordering within each column to align children with parents
		rowPosInCol := make(map[string]float64)
		for c := range numCols {
			if c > 0 {
				bary := make(map[string]float64)
				for _, ref := range cols[c] {
					key := ref.Pkg + "." + ref.Name
					sum := 0.0
					cnt := 0
					for _, pKey := range inEdges[key] {
						if pPos, ok := rowPosInCol[pKey]; ok {
							sum += pPos
							cnt++
						}
					}
					if cnt > 0 {
						bary[key] = sum / float64(cnt)
					} else {
						bary[key] = float64(origOrder[key])
					}
				}
				slices.SortStableFunc(cols[c], func(a, b structRef) int {
					ka := a.Pkg + "." + a.Name
					kb := b.Pkg + "." + b.Name
					if math.Abs(bary[ka]-bary[kb]) > 1e-6 {
						if bary[ka] < bary[kb] {
							return -1
						}
						return 1
					}
					return origOrder[ka] - origOrder[kb]
				})
			}
			for rIdx, ref := range cols[c] {
				key := ref.Pkg + "." + ref.Name
				if len(cols[c]) <= 1 {
					rowPosInCol[key] = 0.5
				} else {
					rowPosInCol[key] = float64(rIdx) / float64(len(cols[c])-1)
				}
			}
		}

		const colStartX = 60.0
		const colGapX = 280.0
		const rowStartY = 60.0
		const rowGapY = 90.0

		shapeByKey := make(map[string]*doc_models.GongStructShape)
		colByKey := make(map[string]int)
		colLeft := make([]float64, numCols)
		colRight := make([]float64, numCols)

		currentX := colStartX
		for c := range numCols {
			if len(cols[c]) == 0 {
				continue
			}
			colLeft[c] = currentX
			currentY := rowStartY + float64(c%2)*25.0
			maxColWidth := 320.0

			for _, ref := range cols[c] {
				key := ref.Pkg + "." + ref.Name
				gs := structsByPkgName[key]
				if gs == nil {
					continue
				}

				var attrShapes []*doc_models.AttributeShape
				const maxAttrs = 6
				for _, field := range gs.Fields {
					if field.GetName() == "Name" {
						continue
					}
					if len(attrShapes) >= maxAttrs {
						break
					}
					switch rf := field.(type) {
					case *gong_models.GongBasicField:
						parts := strings.Split(rf.DeclaredType, ".")
						ftName := parts[len(parts)-1]
						as := (&doc_models.AttributeShape{
							Name:           rf.Name,
							IdentifierMeta: fmt.Sprintf("ref_%s.%s{}.%s", ref.Pkg, gs.Name, rf.Name),
							Structname:     gs.Name,
							Fieldtypename:  ftName,
							IsHidden:       false,
						}).Stage(docStage)
						attrShapes = append(attrShapes, as)
					case *gong_models.GongTimeField:
						as := (&doc_models.AttributeShape{
							Name:           rf.Name,
							IdentifierMeta: fmt.Sprintf("ref_%s.%s{}.%s", ref.Pkg, gs.Name, rf.Name),
							Structname:     gs.Name,
							Fieldtypename:  "Time",
							IsHidden:       false,
						}).Stage(docStage)
						attrShapes = append(attrShapes, as)
					}
				}

				titleText := gs.Name
				if ref.Pkg != "models" {
					titleText = ref.Pkg + "." + gs.Name
				}
				width := math.Max(320.0, float64(len(titleText))*9.2+80.0)
				for _, as := range attrShapes {
					lineLen := len(as.Name) + 3 + len(as.Fieldtypename)
					w := float64(lineLen)*8.0 + 45.0
					if w > width {
						width = w
					}
				}
				height := 63.0 + float64(len(attrShapes))*20.0
				maxSideLinks := max(inLinkCount[key], outLinkCount[key])
				if maxSideLinks >= 2 {
					minLinkH := 40.0 + float64(maxSideLinks)*48.0
					if minLinkH > height {
						height = minLinkH
					}
				}

				gss := (&doc_models.GongStructShape{
					Name:            cd.Name + "-" + gs.Name,
					IdentifierMeta:  fmt.Sprintf("ref_%s.%s{}", ref.Pkg, gs.Name),
					AttributeShapes: attrShapes,
				}).Stage(docStage)
				gss.X = currentX
				gss.Y = currentY
				gss.Width = math.Round(width)
				gss.Height = math.Round(height)

				cd.GongStructShapes = append(cd.GongStructShapes, gss)
				shapeByKey[key] = gss
				colByKey[key] = c

				if gss.Width > maxColWidth {
					maxColWidth = gss.Width
				}
				currentY += gss.Height + rowGapY
			}

			colRight[c] = currentX + maxColWidth
			currentX += maxColWidth + colGapX
		}

		if len(spec.Enums) > 0 {
			enumX := colStartX
			enumY := 60.0
			for _, gss := range cd.GongStructShapes {
				if gss.Y+gss.Height+100.0 > enumY {
					enumY = gss.Y + gss.Height + 100.0
				}
			}
			for _, enumRef := range spec.Enums {
				ge, ok := enumsByName[enumRef.Name]
				if !ok {
					ge, ok = enumsByName[enumRef.Name+"_Member_1"]
				}
				if !ok {
					continue
				}
				var valShapes []*doc_models.GongEnumValueShape
				const maxEnumVals = 8
				for i, ev := range ge.GongEnumValues {
					if i >= maxEnumVals {
						break
					}
					evs := (&doc_models.GongEnumValueShape{
						Name:           ev.Name,
						IdentifierMeta: fmt.Sprintf("ref_%s.%s", enumRef.Pkg, ev.Name),
						IsHidden:       false,
					}).Stage(docStage)
					valShapes = append(valShapes, evs)
				}
				titleText := enumRef.Pkg + "." + ge.Name
				width := math.Max(340.0, float64(len(titleText))*9.2+80.0)
				height := 63.0 + float64(len(valShapes))*20.0

				ges := (&doc_models.GongEnumShape{
					Name:                cd.Name + "-" + ge.Name,
					IdentifierMeta:      fmt.Sprintf("new(ref_%s.%s)", enumRef.Pkg, ge.Name),
					GongEnumValueShapes: valShapes,
				}).Stage(docStage)
				ges.X = enumX
				ges.Y = enumY
				ges.Width = math.Round(width)
				ges.Height = math.Round(height)

				cd.GongEnumShapes = append(cd.GongEnumShapes, ges)
				enumX += ges.Width + 110.0
			}
		}

		// Collect all links across the diagram for coordinated ratio and corner calculation
		type linkItem struct {
			srcKey             string
			srcPkg             string
			srcName            string
			srcShape           *doc_models.GongStructShape
			srcCol             int
			dstKey             string
			dstPkg             string
			dstName            string
			dstShape           *doc_models.GongStructShape
			dstCol             int
			fieldName          string
			targetMultiplicity doc_models.MultiplicityType
			startRatio         float64
			endRatio           float64
			cornerOffsetRatio  float64
		}

		var allLinks []*linkItem
		outLinksBySrc := make(map[string][]*linkItem)
		inLinksByDst := make(map[string][]*linkItem)

		for c := range numCols {
			for _, srcRef := range cols[c] {
				srcKey := srcRef.Pkg + "." + srcRef.Name
				srcShape := shapeByKey[srcKey]
				srcGS := structsByPkgName[srcKey]
				if srcShape == nil || srcGS == nil {
					continue
				}
				pairCount := make(map[string]int)
				for _, field := range srcGS.Fields {
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
					tShape, exists := shapeByKey[tKey]
					if !exists || tKey == srcKey {
						continue
					}
					if pairCount[tKey] >= 1 {
						continue
					}
					pairCount[tKey]++
					li := &linkItem{
						srcKey:             srcKey,
						srcPkg:             srcRef.Pkg,
						srcName:            srcGS.Name,
						srcShape:           srcShape,
						srcCol:             c,
						dstKey:             tKey,
						dstPkg:             tPkg,
						dstName:            targetGS.Name,
						dstShape:           tShape,
						dstCol:             colByKey[tKey],
						fieldName:          field.GetName(),
						targetMultiplicity: mult,
					}
					allLinks = append(allLinks, li)
					outLinksBySrc[srcKey] = append(outLinksBySrc[srcKey], li)
					inLinksByDst[tKey] = append(inLinksByDst[tKey], li)
				}
			}
		}

		// 1. Compute StartRatio per source shape (sorted top-to-bottom by target center Y)
		for _, list := range outLinksBySrc {
			slices.SortStableFunc(list, func(a, b *linkItem) int {
				ay := a.dstShape.Y + a.dstShape.Height*0.5
				by := b.dstShape.Y + b.dstShape.Height*0.5
				if math.Abs(ay-by) > 1.0 {
					if ay < by {
						return -1
					}
					return 1
				}
				return strings.Compare(a.fieldName, b.fieldName)
			})
			for idx, li := range list {
				if len(list) == 1 {
					li.startRatio = 0.5
				} else {
					li.startRatio = 0.22 + 0.56*float64(idx)/float64(len(list)-1)
				}
			}
		}

		// 2. Compute EndRatio per target shape (sorted top-to-bottom by source exit Y)
		for _, list := range inLinksByDst {
			slices.SortStableFunc(list, func(a, b *linkItem) int {
				ay := a.srcShape.Y + a.startRatio*a.srcShape.Height
				by := b.srcShape.Y + b.startRatio*b.srcShape.Height
				if math.Abs(ay-by) > 1.0 {
					if ay < by {
						return -1
					}
					return 1
				}
				return strings.Compare(a.fieldName, b.fieldName)
			})
			for idx, li := range list {
				if len(list) == 1 {
					li.endRatio = 0.5
				} else {
					li.endRatio = 0.25 + 0.50*float64(idx)/float64(len(list)-1)
				}
			}
		}

		// 3. Compute CornerOffsetRatio per corridor using exact lib/svg math:
		//    c1_X = StartRect.X + CornerOffsetRatio * StartRect.Width
		//    => CornerOffsetRatio = (cornerX - StartRect.X) / StartRect.Width
		forwardByCorridor := make(map[int][]*linkItem)
		backwardByCorridor := make(map[int][]*linkItem)
		for _, li := range allLinks {
			if li.dstCol >= li.srcCol {
				forwardByCorridor[li.srcCol] = append(forwardByCorridor[li.srcCol], li)
			} else {
				backwardByCorridor[li.srcCol] = append(backwardByCorridor[li.srcCol], li)
			}
		}

		for c, list := range forwardByCorridor {
			type linkSpan struct {
				li   *linkItem
				minY float64
				maxY float64
			}
			var spans []linkSpan
			for _, li := range list {
				yStart := li.srcShape.Y + li.startRatio*li.srcShape.Height
				yEnd := li.dstShape.Y + li.endRatio*li.dstShape.Height
				spans = append(spans, linkSpan{
					li:   li,
					minY: math.Min(yStart, yEnd) - 15.0,
					maxY: math.Max(yStart, yEnd) + 15.0,
				})
			}
			slices.SortStableFunc(spans, func(a, b linkSpan) int {
				if a.minY < b.minY {
					return -1
				}
				if a.minY > b.minY {
					return 1
				}
				return 0
			})

			// Group corridor links into vertically overlapping components
			type spanComp struct {
				items []*linkItem
				minY  float64
				maxY  float64
			}
			var comps []spanComp
			for _, sp := range spans {
				if len(comps) == 0 || sp.minY > comps[len(comps)-1].maxY {
					comps = append(comps, spanComp{
						items: []*linkItem{sp.li},
						minY:  sp.minY,
						maxY:  sp.maxY,
					})
				} else {
					last := &comps[len(comps)-1]
					last.items = append(last.items, sp.li)
					if sp.maxY > last.maxY {
						last.maxY = sp.maxY
					}
				}
			}

			for _, comp := range comps {
				compRightX := 0.0
				for _, li := range comp.items {
					if r := li.srcShape.X + li.srcShape.Width; r > compRightX {
						compRightX = r
					}
				}
				for _, ref := range cols[c] {
					if gss := shapeByKey[ref.Pkg+"."+ref.Name]; gss != nil {
						if gss.Y+gss.Height+15.0 >= comp.minY && gss.Y-15.0 <= comp.maxY {
							if r := gss.X + gss.Width; r > compRightX {
								compRightX = r
							}
						}
					}
				}

				if len(comp.items) == 1 {
					li := comp.items[0]
					cornerX := compRightX + 55.0
					li.cornerOffsetRatio = (cornerX - li.srcShape.X) / li.srcShape.Width
					continue
				}

				var upLinks, flatLinks, downLinks []*linkItem
				for _, li := range comp.items {
					yStart := li.srcShape.Y + li.startRatio*li.srcShape.Height
					yEnd := li.dstShape.Y + li.endRatio*li.dstShape.Height
					if yEnd < yStart-4.0 {
						upLinks = append(upLinks, li)
					} else if yEnd > yStart+4.0 {
						downLinks = append(downLinks, li)
					} else {
						flatLinks = append(flatLinks, li)
					}
				}
				// Up links: sort by yStart ascending so higher links turn closer to source column
				slices.SortStableFunc(upLinks, func(a, b *linkItem) int {
					ay := a.srcShape.Y + a.startRatio*a.srcShape.Height
					by := b.srcShape.Y + b.startRatio*b.srcShape.Height
					if ay < by {
						return -1
					}
					if ay > by {
						return 1
					}
					return 0
				})
				// Down links: sort by yStart descending so lower links turn closer to source column (no crossings)
				slices.SortStableFunc(downLinks, func(a, b *linkItem) int {
					ay := a.srcShape.Y + a.startRatio*a.srcShape.Height
					by := b.srcShape.Y + b.startRatio*b.srcShape.Height
					if ay > by {
						return -1
					}
					if ay < by {
						return 1
					}
					return 0
				})
				ordered := append(append(upLinks, flatLinks...), downLinks...)

				channelMinX := compRightX + 40.0
				channelMaxX := compRightX + 40.0 + math.Min(110.0, float64(len(ordered)-1)*28.0)
				if c+1 < numCols && colLeft[c+1]-125.0 < channelMaxX && colLeft[c+1]-125.0 > channelMinX {
					channelMaxX = colLeft[c+1] - 125.0
				}

				for idx, li := range ordered {
					cornerX := channelMinX + (channelMaxX-channelMinX)*float64(idx)/float64(len(ordered)-1)
					li.cornerOffsetRatio = (cornerX - li.srcShape.X) / li.srcShape.Width
				}
			}
		}

		for c, list := range backwardByCorridor {
			for idx, li := range list {
				cornerX := colLeft[c] - 50.0 - float64(idx)*20.0
				li.cornerOffsetRatio = (cornerX - li.srcShape.X) / li.srcShape.Width
			}
		}

		for _, li := range allLinks {
			ls := (&doc_models.LinkShape{
				Name:                    li.fieldName,
				IdentifierMeta:          fmt.Sprintf("ref_%s.%s{}.%s", li.srcPkg, li.srcName, li.fieldName),
				FieldTypeIdentifierMeta: fmt.Sprintf("ref_%s.%s{}", li.dstPkg, li.dstName),
				SourceMultiplicity:      doc_models.MANY,
				TargetMultiplicity:      li.targetMultiplicity,
				StartOrientation:        doc_models.ORIENTATION_HORIZONTAL,
				StartRatio:              math.Round(li.startRatio*1000) / 1000,
				EndOrientation:          doc_models.ORIENTATION_HORIZONTAL,
				EndRatio:                math.Round(li.endRatio*1000) / 1000,
				CornerOffsetRatio:       math.Round(li.cornerOffsetRatio*1000) / 1000,
				X:                       (li.srcShape.X + li.dstShape.X) / 2.0,
				Y:                       (li.srcShape.Y + li.dstShape.Y) / 2.0,
				IsHidden:                false,
			}).Stage(docStage)

			li.srcShape.LinkShapes = append(li.srcShape.LinkShapes, ls)
		}
	}

	docStage.ComputeReferenceAndOrders()

	outPath := "diagrams/diagrams_set.go"
	outFile, err := os.Create(outPath)
	if err != nil {
		t.Fatalf("failed creating %s: %v", outPath, err)
	}
	docStage.Marshall(outFile, "github.com/fullstack-lang/gong/lib/doc/go/models", "diagrams")
	outFile.Close()

	origStructs := len(*docStage.GetInstancesSet[*doc_models.GongStructShape]())
	origAttrs := len(*docStage.GetInstancesSet[*doc_models.AttributeShape]())
	origLinks := len(*docStage.GetInstancesSet[*doc_models.LinkShape]())
	origEnums := len(*docStage.GetInstancesSet[*doc_models.GongEnumShape]())
	origEnumVals := len(*docStage.GetInstancesSet[*doc_models.GongEnumValueShape]())
	origDiagrams := len(*docStage.GetInstancesSet[*doc_models.Classdiagram]())

	verifyMux := http.NewServeMux()
	verifyEditor := &splitlite.AsSplitArea{Name: "VerifyBottom", Size: 50}
	verifyStager := prepare.PrepareStageSetSplitlite(
		verifyMux,
		false,
		domainImportPrefix+":diag_verify",
		metaPackageImports,
		embeddedgo.GoModelsDir,
		embeddedgo.GoDiagramsDir,
		verifyEditor,
		map[string]int{},
	)
	verifyStage := verifyStager.GetStage()

	afterStructs := len(*verifyStage.GetInstancesSet[*doc_models.GongStructShape]())
	afterAttrs := len(*verifyStage.GetInstancesSet[*doc_models.AttributeShape]())
	afterLinks := len(*verifyStage.GetInstancesSet[*doc_models.LinkShape]())
	afterEnums := len(*verifyStage.GetInstancesSet[*doc_models.GongEnumShape]())
	afterEnumVals := len(*verifyStage.GetInstancesSet[*doc_models.GongEnumValueShape]())
	afterDiagrams := len(*verifyStage.GetInstancesSet[*doc_models.Classdiagram]())

	t.Logf("Domain %s: generated %d UML diagrams (%d GongStructShapes, %d AttributeShapes, %d LinkShapes, %d GongEnumShapes, %d GongEnumValueShapes)",
		domainName, afterDiagrams, afterStructs, afterAttrs, afterLinks, afterEnums, afterEnumVals)

	if afterDiagrams != origDiagrams || afterDiagrams != len(specs) {
		t.Fatalf("expected %d diagrams, got %d", len(specs), afterDiagrams)
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
}

func buildDomainDiagramSpecs() []diagramSpec {
	return []diagramSpec{
		{
			Name:        "01. Airport Message & Temporality",
			Description: "Root AIXMBasicMessage container for AirportHeliport, Runway, Taxiway, Apron, Lighting, and Marking features.",
			ExtraCols: [5][]structRef{
				0: {
					{Pkg: "models", Name: "AIXMBasicMessage"},
					{Pkg: "models", Name: "AIXMBasicMessageType"},
				},
				1: {
					{Pkg: "aixmabstractgmlobjecttypes", Name: "MessageMetadataPropertyType"},
					{Pkg: "models", Name: "BasicMessageMemberAIXMPropertyType"},
				},
				2: {
					{Pkg: "aixmfeatures", Name: "AirportHeliportType"},
					{Pkg: "aixmfeatures", Name: "RunwayType"},
					{Pkg: "aixmfeatures", Name: "TaxiwayType"},
					{Pkg: "aixmfeatures", Name: "ApronType"},
				},
				3: {
					{Pkg: "aixmfeatures", Name: "AirportHeliportTimeSlicePropertyType"},
					{Pkg: "aixmfeatures", Name: "RunwayTimeSlicePropertyType"},
					{Pkg: "aixmfeatures", Name: "TaxiwayTimeSlicePropertyType"},
					{Pkg: "aixmfeatures", Name: "ApronTimeSlicePropertyType"},
				},
				4: {
					{Pkg: "aixmfeatures", Name: "AirportHeliportTimeSliceType"},
					{Pkg: "aixmfeatures", Name: "RunwayTimeSliceType"},
					{Pkg: "aixmfeatures", Name: "TaxiwayTimeSliceType"},
					{Pkg: "aixmfeatures", Name: "ApronTimeSliceType"},
					{Pkg: "gml", Name: "TimePrimitivePropertyType"},
				},
			},
			Enums: []structRef{
				{Pkg: "aixmabstractgmlobjecttypes", Name: "Enum_Interpretation_SimpleType"},
			},
		},
		{
			Name:        "02. AirportHeliport, Availability, Usage & City",
			Description: "AirportHeliport, AirportHeliportCollocation, Availability, Usage, ResponsibilityOrganisation, and Served City.",
			Features: []string{
				"AirportHeliport",
				"AirportHeliportCollocation",
			},
			Objects: []string{
				"AirportHeliportAvailability",
				"AirportHeliportUsage",
				"AirportHeliportResponsibilityOrganisation",
				"City",
				"ConditionCombination",
			},
			Enums: []structRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeAirportHeliportBaseType_Member_1"},
				{Pkg: "aixmdatatypes", Name: "Enum_CodeOperationAirportHeliportBaseType_Member_1"},
			},
		},
		{
			Name:        "03. HotSpot, AltimeterSource, NonMovement & WorkArea",
			Description: "AirportHotSpot, AltimeterSource, NonMovementArea, SurveyControlPoint, WorkArea, and WorkareaActivity.",
			Features: []string{
				"AirportHotSpot",
				"AltimeterSource",
				"NonMovementArea",
				"SurveyControlPoint",
				"WorkArea",
			},
			Objects: []string{
				"AltimeterSourceStatus",
				"WorkareaActivity",
			},
		},
		{
			Name:        "04. Apron & ApronElement",
			Description: "Apron, ApronElement, ApronAreaAvailability, ApronAreaUsage, and SurfaceCharacteristics.",
			Features: []string{
				"Apron",
				"ApronElement",
			},
			Objects: []string{
				"ApronAreaAvailability",
				"ApronAreaUsage",
				"SurfaceCharacteristics",
			},
		},
		{
			Name:        "05. AircraftStand, LoadingBridge, DeicingArea & Road",
			Description: "AircraftStand, PassengerLoadingBridge, DeicingArea, Road, and SurfaceCharacteristics.",
			Features: []string{
				"AircraftStand",
				"PassengerLoadingBridge",
				"DeicingArea",
				"Road",
			},
			Objects: []string{
				"ApronAreaAvailability",
				"SurfaceCharacteristics",
			},
		},
		{
			Name:        "06. TouchDownLiftOff (TLOF) & SafeArea",
			Description: "Heliport TouchDownLiftOff (TLOF), TouchDownLiftOffSafeArea, ManoeuvringAreaAvailability, and SurfaceCharacteristics.",
			Features: []string{
				"TouchDownLiftOff",
				"TouchDownLiftOffSafeArea",
			},
			Objects: []string{
				"ManoeuvringAreaAvailability",
				"ManoeuvringAreaUsage",
				"SurfaceCharacteristics",
			},
		},
		{
			Name:        "07. Runway, RunwayElement & Surface",
			Description: "Runway, RunwayElement, ManoeuvringAreaAvailability, ManoeuvringAreaUsage, and SurfaceCharacteristics.",
			Features: []string{
				"Runway",
				"RunwayElement",
			},
			Objects: []string{
				"ManoeuvringAreaAvailability",
				"ManoeuvringAreaUsage",
				"SurfaceCharacteristics",
			},
			Enums: []structRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeRunwayBaseType_Member_1"},
				{Pkg: "aixmdatatypes", Name: "Enum_CodeRunwayElementBaseType_Member_1"},
			},
		},
		{
			Name:        "08. RunwayDirection, CentrelinePoint & DeclaredDistance",
			Description: "RunwayDirection, RunwayCentrelinePoint, RunwayDeclaredDistance, RunwayDeclaredDistanceValue, and NavaidEquipmentDistance.",
			Features: []string{
				"RunwayDirection",
				"RunwayCentrelinePoint",
			},
			Objects: []string{
				"RunwayDeclaredDistance",
				"RunwayDeclaredDistanceValue",
				"NavaidEquipmentDistance",
			},
			Enums: []structRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeRunwayPointRoleBaseType_Member_1"},
			},
		},
		{
			Name:        "09. BlastPad, ProtectArea, RVR, VGSI & ArrestingGear",
			Description: "RunwayBlastPad, RunwayProtectArea, RunwayVisualRange (RVR), VisualGlideSlopeIndicator (VGSI), and ArrestingGear.",
			Features: []string{
				"RunwayBlastPad",
				"RunwayProtectArea",
				"RunwayVisualRange",
				"VisualGlideSlopeIndicator",
				"ArrestingGear",
			},
			Objects: []string{
				"SurfaceCharacteristics",
			},
		},
		{
			Name:        "10. Taxiway, TaxiwayElement, HoldingPosition & GuidanceLine",
			Description: "Taxiway, TaxiwayElement, TaxiHoldingPosition, GuidanceLine, and SurfaceCharacteristics.",
			Features: []string{
				"Taxiway",
				"TaxiwayElement",
				"TaxiHoldingPosition",
				"GuidanceLine",
			},
			Objects: []string{
				"ManoeuvringAreaAvailability",
				"SurfaceCharacteristics",
			},
			Enums: []structRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeTaxiwayBaseType_Member_1"},
				{Pkg: "aixmdatatypes", Name: "Enum_CodeTaxiwayElementBaseType_Member_1"},
			},
		},
		{
			Name:        "11. SeaplaneLandingArea, RampSite, DockSite & Buoy",
			Description: "SeaplaneLandingArea, SeaplaneRampSite, FloatingDockSite, and MarkingBuoy.",
			Features: []string{
				"SeaplaneLandingArea",
				"SeaplaneRampSite",
				"FloatingDockSite",
				"MarkingBuoy",
			},
			Objects: []string{
				"ManoeuvringAreaAvailability",
			},
		},
		{
			Name:        "12. Approach & Runway Ground Lighting",
			Description: "AbstractGroundLightSystem, ApproachLightingSystem, RunwayDirectionLightSystem, RunwayProtectAreaLightSystem, and LightElement.",
			Features: []string{
				"AbstractGroundLightSystem",
				"ApproachLightingSystem",
				"RunwayDirectionLightSystem",
				"RunwayProtectAreaLightSystem",
			},
			Objects: []string{
				"GroundLightingAvailability",
				"LightActivation",
				"LightElement",
				"LightElementStatus",
			},
			Enums: []structRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeApproachLightingBaseType_Member_1"},
				{Pkg: "aixmdatatypes", Name: "Enum_CodeLightIntensityBaseType_Member_1"},
			},
		},
		{
			Name:        "13. Taxiway, Apron, GuidanceLine & TLOF Lighting",
			Description: "TaxiwayLightSystem, TaxiHoldingPositionLightSystem, GuidanceLineLightSystem, ApronLightSystem, and TouchDownLiftOffLightSystem.",
			Features: []string{
				"TaxiwayLightSystem",
				"TaxiHoldingPositionLightSystem",
				"GuidanceLineLightSystem",
				"ApronLightSystem",
				"TouchDownLiftOffLightSystem",
			},
			Objects: []string{
				"GroundLightingAvailability",
				"LightElement",
			},
		},
		{
			Name:        "14. Runway, Taxiway & HoldingPosition Marking",
			Description: "AbstractMarking, RunwayMarking, TaxiwayMarking, TaxiHoldingPositionMarking, and MarkingElement.",
			Features: []string{
				"AbstractMarking",
				"RunwayMarking",
				"TaxiwayMarking",
				"TaxiHoldingPositionMarking",
			},
			Objects: []string{
				"MarkingElement",
			},
			Enums: []structRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeMarkingConditionBaseType_Member_1"},
				{Pkg: "aixmdatatypes", Name: "Enum_CodeMarkingStyleBaseType_Member_1"},
			},
		},
		{
			Name:        "15. Stand, Apron, Deicing, TLOF Marking & Contamination",
			Description: "StandMarking, ApronMarking, DeicingAreaMarking, GuidanceLineMarking, TouchDownLiftOffMarking, and Surface Contamination.",
			Features: []string{
				"StandMarking",
				"ApronMarking",
				"DeicingAreaMarking",
				"GuidanceLineMarking",
				"TouchDownLiftOffMarking",
				"AirportProtectionAreaMarking",
			},
			Objects: []string{
				"MarkingElement",
				"AirportHeliportContamination",
				"RunwayContamination",
				"RunwaySectionContamination",
				"SurfaceContaminationLayer",
				"Ridge",
			},
			Enums: []structRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeContaminationBaseType_Member_1"},
			},
		},
	}
}
