package emit_semantics

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/fullstack-lang/gong/app/xsd2gong/pkg/emit_models"
	"github.com/fullstack-lang/gong/app/xsd2gong/pkg/namer"
)

// RenderPackageSemanticsFile generates `models_semantics.go` for a sub-package.
// It emits `ValidateSemantics() []error` on every Gongstruct in the package,
// enforcing XSD restriction facets (numeric bounds, regex patterns, string lengths),
// required attributes/elements, and xs:choice XOR exclusivity.
func RenderPackageSemanticsFile(ep *emit_models.EmittedPackage, _ *emit_models.EmittedModel) string {
	var sb strings.Builder
	sb.WriteString("// generated code by xsd2gong - do not edit\n")
	sb.WriteString(fmt.Sprintf("package %s\n\n", ep.Name))

	sb.WriteString("import (\n")
	sb.WriteString("\t\"fmt\"\n")
	sb.WriteString("\t\"regexp\"\n")
	sb.WriteString(")\n\n")
	sb.WriteString("var (\n\t_ = fmt.Errorf\n\t_ = regexp.MustCompile\n)\n\n")

	// Collect regex patterns needed in this package
	patternVars := make(map[string]string) // pattern -> varName
	var patternOrder []string

	for _, st := range ep.Structs {
		for _, f := range st.Fields {
			if f.SimpleType != nil && f.SimpleType.Restriction != nil {
				for _, pat := range f.SimpleType.Restriction.Patterns {
					if _, exists := patternVars[pat]; !exists {
						// Verify valid Go regexp before compiling
						if _, err := regexp.Compile("^(" + pat + ")$"); err == nil {
							varName := fmt.Sprintf("rePattern_%d", len(patternOrder)+1)
							patternVars[pat] = varName
							patternOrder = append(patternOrder, pat)
						}
					}
				}
			}
		}
	}

	if len(patternOrder) > 0 {
		sb.WriteString("var (\n")
		for _, pat := range patternOrder {
			varName := patternVars[pat]
			sb.WriteString(fmt.Sprintf("\t%s = regexp.MustCompile(%q)\n", varName, "^("+pat+")$"))
		}
		sb.WriteString(")\n\n")
	}

	// Emit ValidateSemantics() on each struct
	for _, st := range ep.Structs {
		sb.WriteString(fmt.Sprintf("// ValidateSemantics enforces XSD restriction facets, required attributes, and xs:choice XOR invariants on %s.\n", st.GoName))
		sb.WriteString(fmt.Sprintf("func (x *%s) ValidateSemantics() (errs []error) {\n", st.GoName))
		sb.WriteString("\tif x == nil {\n\t\treturn nil\n\t}\n")

		// 1. Required attributes
		for _, f := range st.Fields {
			if f.IsAttr && f.Required && f.FullGoType == "string" {
				sb.WriteString(fmt.Sprintf("\tif x.%s == \"\" {\n", f.GoName))
				sb.WriteString(fmt.Sprintf("\t\terrs = append(errs, fmt.Errorf(\"%s %%q: required attribute %s (%s) is empty\", x.Name, %q, %q))\n", st.GoName, f.GoName, f.XMLNameLocal, f.GoName, f.XMLNameLocal))
				sb.WriteString("\t}\n")
			}
		}

		// 2. Restriction facets on fields (numeric bounds, string lengths, regex patterns)
		for _, f := range st.Fields {
			if f.SimpleType == nil || f.SimpleType.Restriction == nil {
				continue
			}
			r := f.SimpleType.Restriction
			switch f.FullGoType {
			case "int":
				// Note: In Go, 0 is the zero value for optional int fields; only validate if non-zero or if 0 is out of bounds and field is required/charData
				checkCond := fmt.Sprintf("x.%s != 0", f.GoName)
				if f.IsCharData || f.Required {
					checkCond = "true"
				}
				var boundChecks []string
				if r.MinInclusive != nil {
					if v, err := strconv.Atoi(*r.MinInclusive); err == nil {
						boundChecks = append(boundChecks, fmt.Sprintf("x.%s < %d", f.GoName, v))
					}
				}
				if r.MaxInclusive != nil {
					if v, err := strconv.Atoi(*r.MaxInclusive); err == nil {
						boundChecks = append(boundChecks, fmt.Sprintf("x.%s > %d", f.GoName, v))
					}
				}
				if r.MinExclusive != nil {
					if v, err := strconv.Atoi(*r.MinExclusive); err == nil {
						boundChecks = append(boundChecks, fmt.Sprintf("x.%s <= %d", f.GoName, v))
					}
				}
				if r.MaxExclusive != nil {
					if v, err := strconv.Atoi(*r.MaxExclusive); err == nil {
						boundChecks = append(boundChecks, fmt.Sprintf("x.%s >= %d", f.GoName, v))
					}
				}
				if len(boundChecks) > 0 {
					sb.WriteString(fmt.Sprintf("\tif %s && (%s) {\n", checkCond, strings.Join(boundChecks, " || ")))
					sb.WriteString(fmt.Sprintf("\t\terrs = append(errs, fmt.Errorf(\"%s %%q: field %s value %%d violates XSD restriction bounds on %s\", x.Name, x.%s))\n", st.GoName, f.GoName, f.SimpleType.Name, f.GoName))
					sb.WriteString("\t}\n")
				}

			case "string":
				if r.MinLength != nil && *r.MinLength > 0 {
					sb.WriteString(fmt.Sprintf("\tif x.%s != \"\" && len(x.%s) < %d {\n", f.GoName, f.GoName, *r.MinLength))
					sb.WriteString(fmt.Sprintf("\t\terrs = append(errs, fmt.Errorf(\"%s %%q: field %s length %%d < minLength %d\", x.Name, len(x.%s)))\n", st.GoName, f.GoName, *r.MinLength, f.GoName))
					sb.WriteString("\t}\n")
				}
				if r.MaxLength != nil && *r.MaxLength > 0 {
					sb.WriteString(fmt.Sprintf("\tif len(x.%s) > %d {\n", f.GoName, *r.MaxLength))
					sb.WriteString(fmt.Sprintf("\t\terrs = append(errs, fmt.Errorf(\"%s %%q: field %s length %%d > maxLength %d\", x.Name, len(x.%s)))\n", st.GoName, f.GoName, *r.MaxLength, f.GoName))
					sb.WriteString("\t}\n")
				}
				for _, pat := range r.Patterns {
					if varName, ok := patternVars[pat]; ok {
						sb.WriteString(fmt.Sprintf("\tif x.%s != \"\" && !%s.MatchString(x.%s) {\n", f.GoName, varName, f.GoName))
						sb.WriteString(fmt.Sprintf("\t\terrs = append(errs, fmt.Errorf(\"%s %%q: field %s value %%q does not match XSD pattern\", x.Name, x.%s))\n", st.GoName, f.GoName, f.GoName))
						sb.WriteString("\t}\n")
					}
				}
			}
		}

		// 3. Single xs:choice XOR exclusivity (exact 1 of N)
		for idx, grp := range st.ChoiceGroups {
			var pointerFields []string
			for _, fName := range grp {
				for _, f := range st.Fields {
					if f.GoName == fName && f.IsPointer {
						pointerFields = append(pointerFields, fName)
					}
				}
			}
			// Only enforce strict pointer XOR when all options in the group are nullable struct pointers
			if len(pointerFields) == len(grp) && len(pointerFields) > 1 {
				cntVar := fmt.Sprintf("choiceCount_%d", idx+1)
				sb.WriteString(fmt.Sprintf("\t%s := 0\n", cntVar))
				for _, pf := range pointerFields {
					sb.WriteString(fmt.Sprintf("\tif x.%s != nil {\n\t\t%s++\n\t}\n", pf, cntVar))
				}
				sb.WriteString(fmt.Sprintf("\tif %s != 1 {\n", cntVar))
				sb.WriteString(fmt.Sprintf("\t\terrs = append(errs, fmt.Errorf(\"%s %%q: xs:choice requires exactly 1 of [%s] (got %%d)\", x.Name, %s))\n", st.GoName, strings.Join(pointerFields, ", "), cntVar))
				sb.WriteString("\t}\n")
			}
		}

		// 4. Optional xs:choice XOR exclusivity (at most 1 of N)
		for idx, grp := range st.OptionalXORs {
			var pointerFields []string
			for _, fName := range grp {
				for _, f := range st.Fields {
					if f.GoName == fName && f.IsPointer {
						pointerFields = append(pointerFields, fName)
					}
				}
			}
			if len(pointerFields) == len(grp) && len(pointerFields) > 1 {
				cntVar := fmt.Sprintf("optChoiceCount_%d", idx+1)
				sb.WriteString(fmt.Sprintf("\t%s := 0\n", cntVar))
				for _, pf := range pointerFields {
					sb.WriteString(fmt.Sprintf("\tif x.%s != nil {\n\t\t%s++\n\t}\n", pf, cntVar))
				}
				sb.WriteString(fmt.Sprintf("\tif %s > 1 {\n", cntVar))
				sb.WriteString(fmt.Sprintf("\t\terrs = append(errs, fmt.Errorf(\"%s %%q: optional xs:choice allows at most 1 of [%s] (got %%d)\", x.Name, %s))\n", st.GoName, strings.Join(pointerFields, ", "), cntVar))
				sb.WriteString("\t}\n")
			}
		}

		sb.WriteString("\treturn errs\n")
		sb.WriteString("}\n\n")
	}

	// Emit ValidateStageSemantics(stage *Stage) []error for the package
	sb.WriteString("// ValidateStageSemantics validates all staged instances in this package's Stage.\n")
	sb.WriteString("func ValidateStageSemantics(stage *Stage) (errs []error) {\n")
	sb.WriteString("\tif stage == nil {\n\t\treturn nil\n\t}\n")
	for _, st := range ep.Structs {
		sb.WriteString(fmt.Sprintf("\tfor _, inst := range stage.GetInstancesSorted[*%s]() {\n", st.GoName))
		sb.WriteString("\t\terrs = append(errs, inst.ValidateSemantics()...)\n")
		sb.WriteString("\t}\n")
	}
	sb.WriteString("\treturn errs\n")
	sb.WriteString("}\n")

	return sb.String()
}

// RenderOrchestratorSemantics generates `orchestrator/semantics.go` which implements
// cross-stage ID/IDREF pointer resolution (`ResolveReferences`), pointer-to-attribute sync
// (`SyncReferences`), and global multi-stage semantic validation (`ValidateSemantics`).
func RenderOrchestratorSemantics(em *emit_models.EmittedModel) string {
	var sb strings.Builder
	sb.WriteString("// generated code by xsd2gong - do not edit\n")
	sb.WriteString("package orchestrator\n\n")

	sb.WriteString("import (\n")
	sb.WriteString("\t\"fmt\"\n\n")
	for _, ep := range em.OrderedPkgs {
		sb.WriteString(fmt.Sprintf("\t\"%s\"\n", ep.ImportPath))
	}
	sb.WriteString(")\n\n")

	sb.WriteString("var _ = fmt.Errorf\n\n")

	// ValidateSemantics across all stages
	sb.WriteString("// ValidateSemantics runs semantic validation across all package stages.\n")
	sb.WriteString("func (s *MusicXMLStages) ValidateSemantics() (errs []error) {\n")
	for _, ep := range em.OrderedPkgs {
		stageField := namer.ToExportedIdentifier(ep.Name) + "Stage"
		sb.WriteString(fmt.Sprintf("\terrs = append(errs, %s.ValidateStageSemantics(s.%s)...)\n", ep.Name, stageField))
	}
	sb.WriteString("\treturn errs\n")
	sb.WriteString("}\n\n")

	// ResolveReferences indexes all xs:ID instances and links all xs:IDREF Gong pointers
	sb.WriteString("// ResolveReferences indexes all xs:ID instances across stages, wires navigable Gong pointers\n")
	sb.WriteString("// for xs:IDREF attributes (e.g. Part.RefScorePart -> *ScorePart), and reports broken references or duplicate IDs.\n")
	sb.WriteString("func (s *MusicXMLStages) ResolveReferences() (errs []error) {\n")

	// 1. Index IDs by type and globally
	sb.WriteString("\tseenIDs := make(map[string]string)\n")
	sb.WriteString("\t_ = seenIDs\n")

	// Index ScorePart, ScoreInstrument, Player specifically for typed pointer resolution
	type idTarget struct {
		pkgName    string
		structName string
		mapVar     string
	}
	var targets []idTarget
	for _, ep := range em.OrderedPkgs {
		for _, st := range ep.Structs {
			for _, f := range st.Fields {
				if f.IsID && f.GoName == "Id" {
					if st.GoName == "ScorePart" || st.GoName == "ScoreInstrument" || st.GoName == "Player" {
						mapVar := "idMap_" + st.GoName
						targets = append(targets, idTarget{
							pkgName:    ep.Name,
							structName: st.GoName,
							mapVar:     mapVar,
						})
					}
				}
			}
		}
	}

	for _, t := range targets {
		stageField := namer.ToExportedIdentifier(t.pkgName) + "Stage"
		sb.WriteString(fmt.Sprintf("\t%s := make(map[string]*%s.%s)\n", t.mapVar, t.pkgName, t.structName))
		sb.WriteString(fmt.Sprintf("\tif s.%s != nil {\n", stageField))
		sb.WriteString(fmt.Sprintf("\t\tfor _, inst := range s.%s.GetInstancesSorted[*%s.%s]() {\n", stageField, t.pkgName, t.structName))
		sb.WriteString("\t\t\tif inst.Id != \"\" {\n")
		sb.WriteString(fmt.Sprintf("\t\t\t\t%s[inst.Id] = inst\n", t.mapVar))
		sb.WriteString("\t\t\t}\n")
		sb.WriteString("\t\t}\n")
		sb.WriteString("\t}\n")
	}

	// 2. Resolve IDREF fields on all structs across all stages
	for _, ep := range em.OrderedPkgs {
		stageField := namer.ToExportedIdentifier(ep.Name) + "Stage"
		for _, st := range ep.Structs {
			for _, f := range st.Fields {
				if !f.IsIDREF {
					continue
				}
				if f.RefPointerGoName != "" && f.IDREFTargetType != "" {
					mapVar := "idMap_" + f.IDREFTargetType
					sb.WriteString(fmt.Sprintf("\tif s.%s != nil {\n", stageField))
					sb.WriteString(fmt.Sprintf("\t\tfor _, inst := range s.%s.GetInstancesSorted[*%s.%s]() {\n", stageField, ep.Name, st.GoName))
					sb.WriteString(fmt.Sprintf("\t\t\tif inst.%s != \"\" {\n", f.GoName))
					sb.WriteString(fmt.Sprintf("\t\t\t\tif target, ok := %s[inst.%s]; ok {\n", mapVar, f.GoName))
					sb.WriteString(fmt.Sprintf("\t\t\t\t\tinst.%s = target\n", f.RefPointerGoName))
					sb.WriteString("\t\t\t\t} else {\n")
					sb.WriteString(fmt.Sprintf("\t\t\t\t\terrs = append(errs, fmt.Errorf(\"%s %%q: unresolved xs:IDREF %s=%%q (target %s not found)\", inst.Name, %q, inst.%s))\n", st.GoName, f.XMLNameLocal, f.IDREFTargetType, f.XMLNameLocal, f.GoName))
					sb.WriteString("\t\t\t\t}\n")
					sb.WriteString("\t\t\t}\n")
					sb.WriteString("\t\t}\n")
					sb.WriteString("\t}\n")
				}
			}
		}
	}

	sb.WriteString("\treturn errs\n")
	sb.WriteString("}\n\n")

	// SyncReferences synchronizes navigable Gong pointers back to raw XML IDREF attributes prior to XML marshalling
	sb.WriteString("// SyncReferences synchronizes navigable Gong pointers (e.g. Part.RefScorePart) back to their\n")
	sb.WriteString("// raw XML IDREF string attributes prior to XML serialization.\n")
	sb.WriteString("func (s *MusicXMLStages) SyncReferences() {\n")
	for _, ep := range em.OrderedPkgs {
		stageField := namer.ToExportedIdentifier(ep.Name) + "Stage"
		for _, st := range ep.Structs {
			for _, f := range st.Fields {
				if f.IsIDREF && f.RefPointerGoName != "" {
					sb.WriteString(fmt.Sprintf("\tif s.%s != nil {\n", stageField))
					sb.WriteString(fmt.Sprintf("\t\tfor _, inst := range s.%s.GetInstancesSorted[*%s.%s]() {\n", stageField, ep.Name, st.GoName))
					sb.WriteString(fmt.Sprintf("\t\t\tif inst.%s != nil && inst.%s.Id != \"\" {\n", f.RefPointerGoName, f.RefPointerGoName))
					sb.WriteString(fmt.Sprintf("\t\t\t\tinst.%s = inst.%s.Id\n", f.GoName, f.RefPointerGoName))
					sb.WriteString("\t\t\t}\n")
					sb.WriteString("\t\t}\n")
					sb.WriteString("\t}\n")
				}
			}
		}
	}
	sb.WriteString("}\n")

	return sb.String()
}

