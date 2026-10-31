package models_test

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"testing"

	embeddedgo "github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models"
	gong_models "github.com/fullstack-lang/gong/go/models"
	doc_models "github.com/fullstack-lang/gong/lib/doc/go/models"
	"github.com/fullstack-lang/gong/lib/doc/go/prepare"
	splitlite "github.com/fullstack-lang/gong/lib/splitlite/go/models"
	"github.com/fullstack-lang/gong/lib/strutils"
)

const (
	domainName         = "service"
	domainImportPrefix = "github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models"
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
	NoteName    string
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
	notesByName := gongStage.GetInstancesMapByName[*gong_models.GongNote]()

	docStage.Reset()

	diagramPkg := (&doc_models.DiagramPackage{
		Name: "AIXM 5.1.1 UML Navigator Diagrams - " + domainName,
	}).Stage(docStage)

	specs := buildDomainDiagramSpecs()
	reDocLink := regexp.MustCompile(`\[models\.([^\]]+)\]`)

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
			NodeGongNotesIsExpanded:   true,
			NodeGongNoteNodeExpansion: `[true]`,
		}).Stage(docStage)

		diagramPkg.Classdiagrams = append(diagramPkg.Classdiagrams, cd)
		if diagIdx == 0 {
			diagramPkg.SelectedClassdiagram = cd
		}

		// Create GongNoteShape at the top-left of the diagram to explain its story
		const colStartX = 40.0
		rowStartY := 40.0
		var noteShape *doc_models.GongNoteShape
		var gongNote *gong_models.GongNote
		if spec.NoteName != "" {
			if gn, ok := notesByName[spec.NoteName]; ok {
				gongNote = gn
				bodyLines := strings.SplitN(gn.Body, "\n", 2)
				noteTitle := ""
				noteContent := ""
				if len(bodyLines) > 0 {
					noteTitle = strings.TrimSpace(bodyLines[0])
				}
				if len(bodyLines) > 1 {
					noteContent = strings.TrimSpace(bodyLines[1])
					noteContent = reDocLink.ReplaceAllString(noteContent, "$1")
				}
				noteWidth := math.Max(640.0, float64(len(noteTitle))*8.0+60.0)
				wrapped := strutils.WrapString(noteContent, int(noteWidth/8.0))
				numLines := len(strings.Split(wrapped, "\n"))
				noteHeight := 52.0 + float64(numLines)*18.0

				noteShape = (&doc_models.GongNoteShape{
					Name:       cd.Name + "-" + gn.Name,
					Identifier: doc_models.GongNoteNameToIdentifier(gn.Name),
					Body:       gn.Body,
					BodyHTML:   gn.BodyHTML,
				}).Stage(docStage)
				noteShape.X = colStartX
				noteShape.Y = 20.0
				noteShape.Width = math.Round(noteWidth)
				noteShape.Height = math.Round(noteHeight)

				cd.GongNoteShapes = append(cd.GongNoteShapes, noteShape)
				rowStartY = noteShape.Y + noteShape.Height + 28.0
			}
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

		// In diagram 01 (Message & Temporality), features are explicitly placed via ExtraCols.
		// In domain diagrams 02..N, start each feature at its TimeSliceType so we don't repeat
		// the 2 extra FeatureType -> FeatureTimeSlicePropertyType columns on every diagram.
		for _, feat := range spec.Features {
			base := strings.TrimPrefix(feat, "Abstract")
			addRef(0, structRef{Pkg: "aixmfeatures", Name: base + "TimeSliceType"})
		}

		for _, obj := range spec.Objects {
			base := strings.TrimPrefix(obj, "Abstract")
			addRef(1, structRef{Pkg: "aixmfeatures", Name: base + "PropertyType"})
			addRef(2, structRef{Pkg: "aixmfeatures", Name: obj + "Type"})
		}

		for c := range 5 {
			for _, ref := range spec.ExtraCols[c] {
				addRef(c, ref)
			}
		}

		// Check if any feature TimeSliceType has no edges to the included Objects/ExtraCols;
		// if so, include its parent FeatureType and TimeSlicePropertyType so it is connected.
		hasEdgeInSet := func(checkKey string) bool {
			for _, srcRef := range allRefs {
				srcKey := srcRef.Pkg + "." + srcRef.Name
				srcGS := structsByPkgName[srcKey]
				if srcGS == nil {
					continue
				}
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
				addRef(0, structRef{Pkg: "aixmfeatures", Name: feat + "Type"})
				addRef(1, structRef{Pkg: "aixmfeatures", Name: base + "TimeSlicePropertyType"})
			}
		}

		// Build directed graph among included structs
		rawOutEdges := make(map[string][]string)
		rawInEdges := make(map[string][]string)
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
				rawOutEdges[srcKey] = append(rawOutEdges[srcKey], tKey)
				rawInEdges[tKey] = append(rawInEdges[tKey], srcKey)
			}
		}

		// Compute topological column for each struct so links flow left-to-right
		rawCol := make(map[string]int)
		pinnedRoot := make(map[string]bool)
		for _, ref := range allRefs {
			key := ref.Pkg + "." + ref.Name
			if len(rawInEdges[key]) == 0 {
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
				for _, v := range rawOutEdges[u] {
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
		colByKey := make(map[string]int)
		origOrder := make(map[string]int)
		for idx, ref := range allRefs {
			key := ref.Pkg + "." + ref.Name
			origOrder[key] = idx
			c := colRemap[rawCol[key]]
			colByKey[key] = c
			cols[c] = append(cols[c], ref)
		}

		// Filter edges: if target v in col c has at least one parent in col c-1,
		// skip redundant long-span edges from col < c-1 so links never cross intermediate columns.
		hasAdjacentParent := make(map[string]bool)
		for _, ref := range allRefs {
			v := ref.Pkg + "." + ref.Name
			vc := colByKey[v]
			for _, u := range rawInEdges[v] {
				if colByKey[u] == vc-1 {
					hasAdjacentParent[v] = true
					break
				}
			}
		}

		outEdges := make(map[string][]string)
		inEdges := make(map[string][]string)
		outLinkCount := make(map[string]int)
		inLinkCount := make(map[string]int)
		for _, ref := range allRefs {
			u := ref.Pkg + "." + ref.Name
			uc := colByKey[u]
			for _, v := range rawOutEdges[u] {
				vc := colByKey[v]
				if vc > uc+1 && hasAdjacentParent[v] {
					continue
				}
				outEdges[u] = append(outEdges[u], v)
				inEdges[v] = append(inEdges[v], u)
				outLinkCount[u]++
				inLinkCount[v]++
			}
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

		// Create GongStructShapes and compute compact Width and Height for each box
		const rowGapY = 24.0
		shapeByKey := make(map[string]*doc_models.GongStructShape)
		maxColWidth := make([]float64, numCols)
		colPackedHeight := make([]float64, numCols)

		for c := range numCols {
			maxColWidth[c] = 185.0
			for rIdx, ref := range cols[c] {
				key := ref.Pkg + "." + ref.Name
				gs := structsByPkgName[key]
				if gs == nil {
					continue
				}

				var attrShapes []*doc_models.AttributeShape
				const maxAttrs = 4
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
				width := math.Max(185.0, float64(len(titleText))*7.5+48.0)
				for _, as := range attrShapes {
					lineLen := len(as.Name) + 3 + len(as.Fieldtypename)
					w := float64(lineLen)*6.8 + 26.0
					if w > width {
						width = w
					}
				}

				height := 46.0 + float64(len(attrShapes))*20.0
				if height < 62.0 {
					height = 62.0
				}
				if outLinkCount[key] >= 2 {
					outH := 28.0 + float64(outLinkCount[key])*20.0
					if outH > height {
						height = outH
					}
				}
				if inLinkCount[key] >= 2 {
					inH := 32.0 + float64(inLinkCount[key])*34.0
					if inH > height {
						height = inH
					}
				}

				gss := (&doc_models.GongStructShape{
					Name:            cd.Name + "-" + gs.Name,
					IdentifierMeta:  fmt.Sprintf("ref_%s.%s{}", ref.Pkg, gs.Name),
					AttributeShapes: attrShapes,
				}).Stage(docStage)
				gss.Width = math.Round(width)
				gss.Height = math.Round(height)

				cd.GongStructShapes = append(cd.GongStructShapes, gss)
				shapeByKey[key] = gss

				if gss.Width > maxColWidth[c] {
					maxColWidth[c] = gss.Width
				}
				colPackedHeight[c] += gss.Height
				if rIdx > 0 {
					colPackedHeight[c] += rowGapY
				}
			}
		}

		// Compute vertical Y coordinates by anchoring the tallest column and aligning
		// connected neighbors' centerY across columns so 1-to-1 chains are straight horizontal.
		anchorCol := 0
		maxHeight := 0.0
		for c := range numCols {
			if colPackedHeight[c] > maxHeight {
				maxHeight = colPackedHeight[c]
				anchorCol = c
			}
		}

		for c := range numCols {
			curY := rowStartY + (maxHeight-colPackedHeight[c])*0.5
			for _, ref := range cols[c] {
				gss := shapeByKey[ref.Pkg+"."+ref.Name]
				if gss == nil {
					continue
				}
				gss.Y = math.Round(curY)
				curY += gss.Height + rowGapY
			}
		}

		alignColY := func(c int, preferRight bool) {
			n := len(cols[c])
			if n == 0 {
				return
			}
			desiredY := make([]float64, n)
			hasDesired := make([]bool, n)
			for i, ref := range cols[c] {
				key := ref.Pkg + "." + ref.Name
				gss := shapeByKey[key]
				var nbs []string
				if preferRight {
					for _, nb := range outEdges[key] {
						if colByKey[nb] > c {
							nbs = append(nbs, nb)
						}
					}
				} else {
					for _, nb := range inEdges[key] {
						if colByKey[nb] < c {
							nbs = append(nbs, nb)
						}
					}
				}
				if len(nbs) > 0 {
					sumCenter := 0.0
					for _, nb := range nbs {
						nbShape := shapeByKey[nb]
						sumCenter += nbShape.Y + nbShape.Height*0.5
					}
					desiredY[i] = sumCenter/float64(len(nbs)) - gss.Height*0.5
					hasDesired[i] = true
				}
			}
			// Fill in any elements without direct neighbors in the sweep direction relative to neighbors in same column
			for i := n - 2; i >= 0; i-- {
				if !hasDesired[i] && hasDesired[i+1] {
					gss := shapeByKey[cols[c][i].Pkg+"."+cols[c][i].Name]
					desiredY[i] = desiredY[i+1] - rowGapY - gss.Height
					hasDesired[i] = true
				}
			}
			for i := 1; i < n; i++ {
				if !hasDesired[i] && hasDesired[i-1] {
					prev := shapeByKey[cols[c][i-1].Pkg+"."+cols[c][i-1].Name]
					desiredY[i] = desiredY[i-1] + prev.Height + rowGapY
					hasDesired[i] = true
				}
			}
			for i, ref := range cols[c] {
				if !hasDesired[i] {
					desiredY[i] = shapeByKey[ref.Pkg+"."+ref.Name].Y
				}
			}

			// Top-down pass to enforce minimum spacing
			yPos := make([]float64, n)
			for i := range n {
				y := desiredY[i]
				minY := rowStartY
				if i > 0 {
					prev := shapeByKey[cols[c][i-1].Pkg+"."+cols[c][i-1].Name]
					minY = yPos[i-1] + prev.Height + rowGapY
				}
				if y < minY {
					y = minY
				}
				yPos[i] = y
			}
			// Bottom-up pass if last elements were pushed down and space exists above
			for i := n - 2; i >= 0; i-- {
				cur := shapeByKey[cols[c][i].Pkg+"."+cols[c][i].Name]
				maxY := yPos[i+1] - rowGapY - cur.Height
				if yPos[i] > maxY {
					yPos[i] = maxY
				}
				if yPos[i] < desiredY[i] && maxY >= desiredY[i] {
					yPos[i] = desiredY[i]
				}
			}
			if yPos[0] < rowStartY {
				shift := rowStartY - yPos[0]
				for i := range n {
					yPos[i] += shift
				}
			}
			for i, ref := range cols[c] {
				shapeByKey[ref.Pkg+"."+ref.Name].Y = math.Round(yPos[i])
			}
		}

		for c := anchorCol - 1; c >= 0; c-- {
			alignColY(c, true)
		}
		for c := anchorCol + 1; c < numCols; c++ {
			alignColY(c, false)
		}

		// Shift all boxes vertically so the highest box starts exactly at rowStartY
		minBoxY := math.MaxFloat64
		for _, gss := range cd.GongStructShapes {
			if gss.Y < minBoxY {
				minBoxY = gss.Y
			}
		}
		if minBoxY < math.MaxFloat64 && minBoxY != rowStartY {
			shiftY := rowStartY - minBoxY
			for _, gss := range cd.GongStructShapes {
				gss.Y = math.Round(gss.Y + shiftY)
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
			relCornerX         float64
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
					dstCol := colByKey[tKey]
					if dstCol > c+1 && hasAdjacentParent[tKey] {
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
						dstCol:             dstCol,
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
					li.endRatio = 0.24 + 0.52*float64(idx)/float64(len(list)-1)
				}
			}
		}

		// 3. Compute relative vertical track X (relCornerX) per corridor AND adaptive corridor width
		forwardByCorridor := make(map[int][]*linkItem)
		backwardByCorridor := make(map[int][]*linkItem)
		for _, li := range allLinks {
			if li.dstCol >= li.srcCol {
				forwardByCorridor[li.srcCol] = append(forwardByCorridor[li.srcCol], li)
			} else {
				backwardByCorridor[li.srcCol] = append(backwardByCorridor[li.srcCol], li)
			}
		}

		corridorSpanX := make([]float64, numCols)
		for c := range numCols {
			corridorSpanX[c] = maxColWidth[c] + 110.0
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
					minY: math.Min(yStart, yEnd) - 12.0,
					maxY: math.Max(yStart, yEnd) + 12.0,
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
				compRelRight := 0.0
				for _, li := range comp.items {
					if li.srcShape.Width > compRelRight {
						compRelRight = li.srcShape.Width
					}
				}
				for _, ref := range cols[c] {
					if gss := shapeByKey[ref.Pkg+"."+ref.Name]; gss != nil {
						if gss.Y+gss.Height+12.0 >= comp.minY && gss.Y-12.0 <= comp.maxY {
							if gss.Width > compRelRight {
								compRelRight = gss.Width
							}
						}
					}
				}

				if len(comp.items) == 1 {
					li := comp.items[0]
					li.relCornerX = compRelRight + 22.0
					labelSpace := math.Max(72.0, float64(len(li.fieldName))*7.0+24.0)
					reqSpan := li.relCornerX + labelSpace
					if reqSpan > corridorSpanX[c] {
						corridorSpanX[c] = reqSpan
					}
					continue
				}

				var upLinks, flatLinks, downLinks []*linkItem
				for _, li := range comp.items {
					yStart := li.srcShape.Y + li.startRatio*li.srcShape.Height
					yEnd := li.dstShape.Y + li.endRatio*li.dstShape.Height
					if yEnd < yStart-2.0 {
						upLinks = append(upLinks, li)
					} else if yEnd > yStart+2.0 {
						downLinks = append(downLinks, li)
					} else {
						flatLinks = append(flatLinks, li)
					}
				}
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

				var ordered []*linkItem
				ordered = append(ordered, upLinks...)
				ordered = append(ordered, flatLinks...)
				ordered = append(ordered, downLinks...)

				const baseMargin = 22.0
				const trackStep = 15.0
				for idx, li := range ordered {
					li.relCornerX = compRelRight + baseMargin + float64(idx)*trackStep
					labelSpace := math.Max(72.0, float64(len(li.fieldName))*7.0+24.0)
					reqSpan := li.relCornerX + labelSpace
					if reqSpan > corridorSpanX[c] {
						corridorSpanX[c] = reqSpan
					}
				}
			}
		}

		// 4. Assign compact column X coordinates and compute exact CornerOffsetRatio
		colLeft := make([]float64, numCols)
		currentX := colStartX
		for c := range numCols {
			colLeft[c] = math.Round(currentX)
			for _, ref := range cols[c] {
				if gss := shapeByKey[ref.Pkg+"."+ref.Name]; gss != nil {
					gss.X = colLeft[c]
				}
			}
			currentX += corridorSpanX[c]
		}

		for _, list := range forwardByCorridor {
			for _, li := range list {
				cornerX := colLeft[li.srcCol] + li.relCornerX
				li.cornerOffsetRatio = (cornerX - li.srcShape.X) / li.srcShape.Width
			}
		}
		for c, list := range backwardByCorridor {
			for idx, li := range list {
				cornerX := colLeft[c] - 24.0 - float64(idx)*15.0
				li.cornerOffsetRatio = (cornerX - li.srcShape.X) / li.srcShape.Width
			}
		}

		// Attach GongNoteLinkShape from the diagram's GongNoteShape to the highest referenced struct
		// in the left columns so the dashed connector goes straight across the top gap without crossing boxes.
		if noteShape != nil && gongNote != nil {
			var bestTargetName string
			bestScore := math.MaxFloat64
			for _, gl := range gongNote.Links {
				if gl.Recv != "" {
					continue
				}
				for _, gss := range cd.GongStructShapes {
					_, sName := doc_models.IdentifierMetaToPackageAndGongStructName(gss.IdentifierMeta)
					if sName == gl.Name {
						score := (gss.Y - rowStartY)*4.0 + math.Abs(gss.X-noteShape.X)*0.5
						if score < bestScore {
							bestScore = score
							bestTargetName = gl.Name
						}
					}
				}
			}
			if bestTargetName != "" {
				nls := (&doc_models.GongNoteLinkShape{
					Name:       cd.Name + "-" + noteShape.Name + "-" + bestTargetName,
					Identifier: doc_models.GongStructNameToIdentifier(bestTargetName),
					Type:       doc_models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE,
					IsHidden:   false,
				}).Stage(docStage)
				noteShape.GongNoteLinkShapes = append(noteShape.GongNoteLinkShapes, nls)
			}
		}

		// Place GongEnumShapes compactly below the structs (or tucked in the lower-left under Column 0)
		if len(spec.Enums) > 0 {
			maxStructBottomY := rowStartY
			col0BottomY := rowStartY
			for _, gss := range cd.GongStructShapes {
				if gss.Y+gss.Height > maxStructBottomY {
					maxStructBottomY = gss.Y + gss.Height
				}
			}
			for _, ref := range cols[0] {
				if gss := shapeByKey[ref.Pkg+"."+ref.Name]; gss != nil {
					if gss.Y+gss.Height > col0BottomY {
						col0BottomY = gss.Y + gss.Height
					}
				}
			}
			enumX := colStartX
			enumY := math.Round(maxStructBottomY + 28.0)

			for _, enumRef := range spec.Enums {
				ge, ok := enumsByName[enumRef.Name]
				if !ok {
					ge, ok = enumsByName[enumRef.Name+"_Member_1"]
				}
				if !ok {
					continue
				}
				var valShapes []*doc_models.GongEnumValueShape
				const maxEnumVals = 6
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
				width := math.Max(220.0, float64(len(titleText))*7.5+40.0)
				for _, evs := range valShapes {
					w := float64(len(evs.Name))*6.8 + 26.0
					if w > width {
						width = w
					}
				}
				height := 46.0 + float64(len(valShapes))*20.0

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
				enumX += ges.Width + 32.0
			}
		}

		for _, li := range allLinks {
			ls := (&doc_models.LinkShape{
				Name:                      li.fieldName,
				IdentifierMeta:            fmt.Sprintf("ref_%s.%s{}.%s", li.srcPkg, li.srcName, li.fieldName),
				FieldTypeIdentifierMeta:   fmt.Sprintf("ref_%s.%s{}", li.dstPkg, li.dstName),
				SourceMultiplicity:        doc_models.MANY,
				TargetMultiplicity:        li.targetMultiplicity,
				StartOrientation:          doc_models.ORIENTATION_HORIZONTAL,
				StartRatio:                math.Round(li.startRatio*1000) / 1000,
				EndOrientation:            doc_models.ORIENTATION_HORIZONTAL,
				EndRatio:                  math.Round(li.endRatio*1000) / 1000,
				CornerOffsetRatio:         math.Round(li.cornerOffsetRatio*1000) / 1000,
				FieldOffsetX:              -10,
				FieldOffsetY:              -7,
				TargetMultiplicityOffsetX: -12,
				TargetMultiplicityOffsetY: 15,
				SourceMultiplicityOffsetX: 10,
				SourceMultiplicityOffsetY: 15,
				X:                         (li.srcShape.X + li.dstShape.X) / 2.0,
				Y:                         (li.srcShape.Y + li.dstShape.Y) / 2.0,
				IsHidden:                  false,
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
	origNotes := len(*docStage.GetInstancesSet[*doc_models.GongNoteShape]())
	origNoteLinks := len(*docStage.GetInstancesSet[*doc_models.GongNoteLinkShape]())
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
	afterNotes := len(*verifyStage.GetInstancesSet[*doc_models.GongNoteShape]())
	afterNoteLinks := len(*verifyStage.GetInstancesSet[*doc_models.GongNoteLinkShape]())
	afterDiagrams := len(*verifyStage.GetInstancesSet[*doc_models.Classdiagram]())

	t.Logf("Domain %s: generated %d UML diagrams (%d GongStructShapes, %d AttributeShapes, %d LinkShapes, %d GongEnumShapes, %d GongEnumValueShapes, %d GongNoteShapes, %d GongNoteLinkShapes)",
		domainName, afterDiagrams, afterStructs, afterAttrs, afterLinks, afterEnums, afterEnumVals, afterNotes, afterNoteLinks)

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
	if afterNotes != origNotes || afterNotes != len(specs) {
		t.Fatalf("expected %d GongNoteShapes, got before=%d after=%d", len(specs), origNotes, afterNotes)
	}
	if afterNoteLinks != origNoteLinks {
		t.Fatalf("expected %d GongNoteLinkShapes, got %d", origNoteLinks, afterNoteLinks)
	}
}

func buildDomainDiagramSpecs() []diagramSpec {
	return []diagramSpec{
		{
			Name:        "01. Service & Organisation Message & Temporality",
			NoteName:    "Note01ServiceMessage",
			Description: "Root AIXMBasicMessage container for Service/Organisation features and the AIXM Temporality Model.",
			ExtraCols: [5][]structRef{
				0: {
					{Pkg: "models", Name: "AIXMBasicMessage"},
				},
				1: {
					{Pkg: "aixmabstractgmlobjecttypes", Name: "MessageMetadataPropertyType"},
					{Pkg: "models", Name: "BasicMessageMemberAIXMPropertyType"},
				},
				2: {
					{Pkg: "aixmfeatures", Name: "OrganisationAuthorityType"},
					{Pkg: "aixmfeatures", Name: "UnitType"},
				},
				3: {
					{Pkg: "aixmfeatures", Name: "OrganisationAuthorityTimeSlicePropertyType"},
					{Pkg: "aixmfeatures", Name: "UnitTimeSlicePropertyType"},
				},
				4: {
					{Pkg: "aixmfeatures", Name: "OrganisationAuthorityTimeSliceType"},
					{Pkg: "aixmfeatures", Name: "UnitTimeSliceType"},
					{Pkg: "gml", Name: "TimePrimitivePropertyType"},
				},
			},
			Enums: []structRef{
				{Pkg: "aixmabstractgmlobjecttypes", Name: "Enum_Interpretation_SimpleType"},
			},
		},
		{
			Name:        "02. OrganisationAuthority & ContactInformation",
			NoteName:    "Note02OrganisationAuthority",
			Description: "OrganisationAuthority feature, OrganisationAuthorityAssociation, ContactInformation, PostalAddress, TelephoneContact, and OnlineContact.",
			Features: []string{
				"OrganisationAuthority",
			},
			Objects: []string{
				"OrganisationAuthorityAssociation",
				"ContactInformation",
				"PostalAddress",
				"TelephoneContact",
				"OnlineContact",
			},
			Enums: []structRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeOrganisationBaseType_Member_1"},
			},
		},
		{
			Name:        "03. Unit, Availability & Dependency",
			NoteName:    "Note03Unit",
			Description: "Unit (ATC/FIC/ARO/MET) feature, UnitAvailability, UnitDependency, and ContactInformation.",
			Features: []string{
				"Unit",
			},
			Objects: []string{
				"UnitAvailability",
				"UnitDependency",
				"ContactInformation",
			},
			Enums: []structRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeUnitBaseType_Member_1"},
			},
		},
		{
			Name:        "04. Traffic Separation Services (ATC, ATM, Ground)",
			NoteName:    "Note04TrafficSeparation",
			Description: "AirTrafficControlService, AirTrafficManagementService, GroundTrafficControlService, and ServiceOperationalStatus.",
			Features: []string{
				"AirTrafficControlService",
				"AirTrafficManagementService",
				"GroundTrafficControlService",
			},
			Objects: []string{
				"ServiceOperationalStatus",
			},
			Enums: []structRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeServiceATCBaseType_Member_1"},
				{Pkg: "aixmdatatypes", Name: "Enum_CodeStatusServiceBaseType_Member_1"},
			},
		},
		{
			Name:        "05. Information & SearchRescue Services",
			NoteName:    "Note05InformationSAR",
			Description: "InformationService (ATIS/AFIS/VOLMET), SearchRescueService, and ServiceOperationalStatus.",
			Features: []string{
				"InformationService",
				"SearchRescueService",
			},
			Objects: []string{
				"ServiceOperationalStatus",
			},
			Enums: []structRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeServiceInformationBaseType_Member_1"},
				{Pkg: "aixmdatatypes", Name: "Enum_CodeServiceSARBaseType_Member_1"},
			},
		},
		{
			Name:        "06. Airport Ground Services & Supplies",
			NoteName:    "Note06AirportGroundServices",
			Description: "AircraftGroundService, AirportClearanceService, AirportSuppliesService, FireFightingService, PassengerService, Fuel, Oil, Oxygen, and Nitrogen.",
			Features: []string{
				"AircraftGroundService",
				"AirportClearanceService",
				"AirportSuppliesService",
				"FireFightingService",
				"PassengerService",
			},
			Objects: []string{
				"Fuel",
				"Oil",
				"Oxygen",
				"Nitrogen",
			},
		},
		{
			Name:        "07. RadioCommunicationChannel & PilotControlledLighting",
			NoteName:    "Note07RadioComm",
			Description: "RadioCommunicationChannel, PilotControlledLighting, RadioCommunicationOperationalStatus, CallsignDetail, and CircleSector.",
			Features: []string{
				"RadioCommunicationChannel",
				"PilotControlledLighting",
			},
			Objects: []string{
				"RadioCommunicationOperationalStatus",
				"CallsignDetail",
				"CircleSector",
			},
		},
		{
			Name:        "08. Shared Notes, Timesheets & StandardLevel Tables",
			NoteName:    "Note08StandardLevels",
			Description: "SpecialDate, StandardLevelTable, StandardLevelColumn, StandardLevelSector, StandardLevel, Note, LinguisticNote, and Timesheet.",
			Features: []string{
				"SpecialDate",
				"StandardLevelTable",
				"StandardLevelColumn",
				"StandardLevelSector",
				"StandardLevel",
			},
			Objects: []string{
				"Note",
				"LinguisticNote",
				"Timesheet",
			},
		},
	}
}
