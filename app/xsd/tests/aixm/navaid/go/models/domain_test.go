package models_test

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	embeddedgo "github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models"
	gong_models "github.com/fullstack-lang/gong/go/models"
	doc_models "github.com/fullstack-lang/gong/lib/doc/go/models"
	"github.com/fullstack-lang/gong/lib/doc/go/prepare"
	splitlite "github.com/fullstack-lang/gong/lib/splitlite/go/models"
)

const (
	domainName         = "navaid"
	domainImportPrefix = "github.com/fullstack-lang/gong/app/xsd/tests/aixm/navaid/go/models"
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

		var cols [5][]structRef
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
			cols[col] = append(cols[col], ref)
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

		referencedKeys := make(map[string]bool)
		for c := range 5 {
			for _, srcRef := range cols[c] {
				if gs := structsByPkgName[srcRef.Pkg+"."+srcRef.Name]; gs != nil {
					for _, f := range gs.Fields {
						var targetGS *gong_models.GongStruct
						switch rf := f.(type) {
						case *gong_models.PointerToGongStructField:
							targetGS = rf.GongStruct
						case *gong_models.SliceOfPointerToGongStructField:
							targetGS = rf.GongStruct
						}
						if targetGS != nil && targetGS.ModelPkg != nil {
							referencedKeys[targetGS.ModelPkg.PkgGoName+"."+targetGS.Name] = true
						}
					}
				}
			}
		}
		for _, feat := range spec.Features {
			base := strings.TrimPrefix(feat, "Abstract")
			propRef := structRef{Pkg: "aixmfeatures", Name: base + "PropertyType"}
			if referencedKeys[propRef.Pkg+"."+propRef.Name] {
				addRef(3, propRef)
			}
		}

		if len(spec.Features) > 0 || len(spec.Objects) > 0 {
			dtCount := 0
			const maxAutoDataTypes = 4
			checkSources := append([]structRef{}, cols[2]...)
			checkSources = append(checkSources, cols[4]...)
			for _, srcRef := range checkSources {
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

		const colStartX = 60.0
		const colGapX = 250.0
		const rowStartY = 60.0
		const rowGapY = 90.0

		shapeByKey := make(map[string]*doc_models.GongStructShape)
		colByKey := make(map[string]int)

		currentX := colStartX
		for c := range 5 {
			if len(cols[c]) == 0 {
				continue
			}
			currentY := rowStartY + float64(c%2)*35.0
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
				width := math.Max(320.0, float64(len(titleText))*8.5+50.0)
				for _, as := range attrShapes {
					lineLen := len(as.Name) + 3 + len(as.Fieldtypename)
					w := float64(lineLen)*8.0 + 40.0
					if w > width {
						width = w
					}
				}
				height := 63.0 + float64(len(attrShapes))*20.0

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
				width := math.Max(340.0, float64(len(titleText))*8.5+50.0)
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

		for _, c := range []int{0, 1, 2, 3, 4} {
			for _, srcRef := range cols[c] {
				srcKey := srcRef.Pkg + "." + srcRef.Name
				srcShape := shapeByKey[srcKey]
				srcGS := structsByPkgName[srcKey]
				if srcShape == nil || srcGS == nil {
					continue
				}

				type pendingLink struct {
					fieldName          string
					targetKey          string
					targetPkg          string
					targetName         string
					targetShape        *doc_models.GongStructShape
					targetMultiplicity doc_models.MultiplicityType
				}
				var links []pendingLink
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
					if pairCount[tKey] >= 2 {
						continue
					}
					pairCount[tKey]++
					links = append(links, pendingLink{
						fieldName:          field.GetName(),
						targetKey:          tKey,
						targetPkg:          tPkg,
						targetName:         targetGS.Name,
						targetShape:        tShape,
						targetMultiplicity: mult,
					})
				}

				for idx, pl := range links {
					startRatio := 0.5
					if len(links) > 1 {
						startRatio = 0.2 + 0.6*float64(idx)/float64(len(links)-1)
					}
					endRatio := 0.5
					if pairCount[pl.targetKey] > 1 {
						endRatio = 0.3 + 0.4*float64(idx%2)
					}

					startOrient := doc_models.ORIENTATION_HORIZONTAL
					endOrient := doc_models.ORIENTATION_HORIZONTAL
					cornerOffset := 0.45 + 0.15*float64(idx%3)

					if colByKey[srcKey] == colByKey[pl.targetKey] {
						startOrient = doc_models.ORIENTATION_VERTICAL
						endOrient = doc_models.ORIENTATION_VERTICAL
						cornerOffset = 0.5
					}

					ls := (&doc_models.LinkShape{
						Name:                    pl.fieldName,
						IdentifierMeta:          fmt.Sprintf("ref_%s.%s{}.%s", srcRef.Pkg, srcGS.Name, pl.fieldName),
						FieldTypeIdentifierMeta: fmt.Sprintf("ref_%s.%s{}", pl.targetPkg, pl.targetName),
						SourceMultiplicity:      doc_models.MANY,
						TargetMultiplicity:      pl.targetMultiplicity,
						StartOrientation:        startOrient,
						StartRatio:              math.Round(startRatio*100) / 100,
						EndOrientation:          endOrient,
						EndRatio:                math.Round(endRatio*100) / 100,
						CornerOffsetRatio:       math.Round(cornerOffset*100) / 100,
						X:                       (srcShape.X + pl.targetShape.X) / 2.0,
						Y:                       (srcShape.Y + pl.targetShape.Y) / 2.0,
						IsHidden:                false,
					}).Stage(docStage)

					srcShape.LinkShapes = append(srcShape.LinkShapes, ls)
				}
			}
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
			Name:        "01. Navaid & Point Message & Temporality",
			Description: "Root AIXMBasicMessage container for Navaids, DesignatedPoints, and Radar features.",
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
					{Pkg: "aixmfeatures", Name: "NavaidType"},
					{Pkg: "aixmfeatures", Name: "VORType"},
					{Pkg: "aixmfeatures", Name: "DesignatedPointType"},
					{Pkg: "aixmfeatures", Name: "RadarSystemType"},
				},
				3: {
					{Pkg: "aixmfeatures", Name: "NavaidTimeSlicePropertyType"},
					{Pkg: "aixmfeatures", Name: "VORTimeSlicePropertyType"},
					{Pkg: "aixmfeatures", Name: "DesignatedPointTimeSlicePropertyType"},
					{Pkg: "aixmfeatures", Name: "RadarSystemTimeSlicePropertyType"},
				},
				4: {
					{Pkg: "aixmfeatures", Name: "NavaidTimeSliceType"},
					{Pkg: "aixmfeatures", Name: "VORTimeSliceType"},
					{Pkg: "aixmfeatures", Name: "DesignatedPointTimeSliceType"},
					{Pkg: "aixmfeatures", Name: "RadarSystemTimeSliceType"},
					{Pkg: "gml", Name: "TimePrimitivePropertyType"},
				},
			},
			Enums: []structRef{
				{Pkg: "aixmabstractgmlobjecttypes", Name: "Enum_Interpretation_SimpleType"},
			},
		},
		{
			Name:        "02. Navaid, Component & Operational Status",
			Description: "Navaid system, AbstractNavaidEquipment, NavaidComponent, NavaidOperationalStatus, Monitoring, and Authority.",
			Features: []string{
				"Navaid",
				"AbstractNavaidEquipment",
			},
			Objects: []string{
				"NavaidComponent",
				"NavaidOperationalStatus",
				"NavaidEquipmentMonitoring",
				"AuthorityForNavaidEquipment",
			},
			Enums: []structRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeNavaidServiceBaseType_Member_1"},
				{Pkg: "aixmdatatypes", Name: "Enum_CodeNavaidPurposeBaseType_Member_1"},
				{Pkg: "aixmdatatypes", Name: "Enum_CodeStatusNavaidBaseType_Member_1"},
			},
		},
		{
			Name:        "03. En-Route Radio Navaids (VOR, DME, TACAN, NDB)",
			Description: "VOR, DME, TACAN, and NDB radio navigation equipment features.",
			Features: []string{
				"VOR",
				"DME",
				"TACAN",
				"NDB",
			},
			Objects: []string{
				"NavaidOperationalStatus",
				"ElevatedPoint",
			},
			Enums: []structRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeVORBaseType_Member_1"},
				{Pkg: "aixmdatatypes", Name: "Enum_CodeDMEBaseType_Member_1"},
			},
		},
		{
			Name:        "04. ILS & MLS Precision Approach Navaids",
			Description: "Localizer, Glidepath, MarkerBeacon, SDF, Azimuth, Elevation, and DirectionFinder.",
			Features: []string{
				"Localizer",
				"Glidepath",
				"MarkerBeacon",
				"SDF",
				"Azimuth",
				"Elevation",
				"DirectionFinder",
			},
			Objects: []string{
				"NavaidOperationalStatus",
			},
		},
		{
			Name:        "05. Navigation Checkpoints (VOR & INS)",
			Description: "AbstractNavigationSystemCheckpoint, CheckpointVOR, and CheckpointINS.",
			Features: []string{
				"AbstractNavigationSystemCheckpoint",
				"CheckpointVOR",
				"CheckpointINS",
			},
			Objects: []string{
				"ElevatedPoint",
				"Note",
			},
		},
		{
			Name:        "06. Special Navigation System & Station",
			Description: "SpecialNavigationSystem, SpecialNavigationStation, SpecialNavigationStationStatus, and Authority objects.",
			Features: []string{
				"SpecialNavigationSystem",
				"SpecialNavigationStation",
			},
			Objects: []string{
				"SpecialNavigationStationStatus",
				"AuthorityForSpecialNavigationSystem",
				"AuthorityForSpecialNavigationStation",
			},
		},
		{
			Name:        "07. DesignatedPoint & SignificantPointInAirspace",
			Description: "DesignatedPoint, SignificantPointInAirspace, AbstractSegmentPoint, EnRouteSegmentPoint, and TerminalSegmentPoint.",
			Features: []string{
				"DesignatedPoint",
				"SignificantPointInAirspace",
			},
			Objects: []string{
				"AbstractSegmentPoint",
				"EnRouteSegmentPoint",
				"TerminalSegmentPoint",
				"PointReference",
			},
			Enums: []structRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeDesignatedPointBaseType_Member_1"},
			},
		},
		{
			Name:        "08. AngleIndication & DistanceIndication",
			Description: "AngleIndication, DistanceIndication, AngleUse, and PointReference.",
			Features: []string{
				"AngleIndication",
				"DistanceIndication",
			},
			Objects: []string{
				"AngleUse",
				"PointReference",
			},
		},
		{
			Name:        "09. Surveillance Radar Systems (PSR, SSR, PAR)",
			Description: "RadarSystem, PrimarySurveillanceRadar, SecondarySurveillanceRadar, PrecisionApproachRadar, RadarComponent, and Reflector.",
			Features: []string{
				"RadarSystem",
				"AbstractRadarEquipment",
				"AbstractSurveillanceRadar",
				"PrimarySurveillanceRadar",
				"SecondarySurveillanceRadar",
				"PrecisionApproachRadar",
			},
			Objects: []string{
				"RadarComponent",
				"Reflector",
				"SurveillanceGroundStation",
			},
			Enums: []structRef{
				{Pkg: "aixmdatatypes", Name: "Enum_CodeRadarServiceBaseType_Member_1"},
				{Pkg: "aixmdatatypes", Name: "Enum_CodePrimaryRadarBaseType_Member_1"},
			},
		},
	}
}
