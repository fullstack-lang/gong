package diagramlayout

import (
	"fmt"
	"math"

	doc_models "github.com/fullstack-lang/gong/lib/doc/go/models"
)

// Rect2D represents an axis-aligned rectangle [X1, X2] x [Y1, Y2].
type Rect2D struct {
	Kind string
	Name string
	X1   float64
	Y1   float64
	X2   float64
	Y2   float64
}

func (r Rect2D) Width() float64  { return r.X2 - r.X1 }
func (r Rect2D) Height() float64 { return r.Y2 - r.Y1 }

func rectsOverlap(a, b Rect2D, pad float64) bool {
	return a.X1-pad < b.X2 && a.X2+pad > b.X1 && a.Y1-pad < b.Y2 && a.Y2+pad > b.Y1
}

// SimSegment mirrors svgspecific's Segment from draw.segments.ts.
type SimSegment struct {
	StartX      float64
	StartY      float64
	EndX        float64
	EndY        float64
	Orientation doc_models.OrientationType
}

func drawPointRectSeg(px, py float64, r Rect2D, dir doc_models.OrientationType, cornerRadius float64, num int) SimSegment {
	rx, ry, rw, rh := r.X1, r.Y1, r.Width(), r.Height()
	var seg SimSegment
	seg.Orientation = dir

	if dir == doc_models.ORIENTATION_HORIZONTAL {
		cx := px
		if px < rx+0.5*rw {
			if px < rx {
				cx += cornerRadius
			} else {
				cx -= cornerRadius
			}
		} else {
			if px > rx+rw {
				cx -= cornerRadius
			} else {
				cx += cornerRadius
			}
		}
		seg.StartX = cx
		seg.StartY = py
		if px <= rx+0.5*rw {
			seg.EndX = rx
		} else {
			seg.EndX = rx + rw
		}
		seg.EndY = py
	} else {
		cy := py
		if py < ry+0.5*rh {
			if py < ry {
				cy += cornerRadius
			} else {
				cy -= cornerRadius
			}
		} else {
			if py < ry+rh {
				cy += cornerRadius
			} else {
				cy -= cornerRadius
			}
		}
		seg.StartX = px
		seg.StartY = cy
		seg.EndX = px
		if py <= ry+0.5*rh {
			seg.EndY = ry
		} else {
			seg.EndY = ry + rh
		}
	}

	if num == 0 {
		seg.StartX, seg.EndX = seg.EndX, seg.StartX
		seg.StartY, seg.EndY = seg.EndY, seg.StartY
	}
	return seg
}

func drawPointPointSeg(x1, y1, x2, y2 float64, dir doc_models.OrientationType, cornerRadius float64) SimSegment {
	sx, sy, ex, ey := x1, y1, x2, y2
	if dir == doc_models.ORIENTATION_HORIZONTAL {
		if x2 >= x1 {
			sx += cornerRadius
			ex -= cornerRadius
		} else {
			sx -= cornerRadius
			ex += cornerRadius
		}
	} else {
		if y2 >= y1 {
			sy += cornerRadius
			ey -= cornerRadius
		} else {
			sy -= cornerRadius
			ey += cornerRadius
		}
	}
	return SimSegment{StartX: sx, StartY: sy, EndX: ex, EndY: ey, Orientation: dir}
}

// SimulateLinkSegments ports drawSegments (lib/svg draw.segments.ts) for LINK_TYPE_FLOATING_ORTHOGONAL.
func SimulateLinkSegments(startRect, endRect Rect2D, ls *doc_models.LinkShape) []SimSegment {
	startDir := ls.StartOrientation
	if startDir == "" {
		startDir = doc_models.ORIENTATION_HORIZONTAL
	}
	endDir := ls.EndOrientation
	if endDir == "" {
		endDir = doc_models.ORIENTATION_HORIZONTAL
	}
	const cornerRadius = 8.0

	if startDir == doc_models.ORIENTATION_HORIZONTAL && endDir == doc_models.ORIENTATION_VERTICAL {
		c1X := endRect.X1 + ls.EndRatio*endRect.Width()
		c1Y := startRect.Y1 + ls.StartRatio*startRect.Height()
		return []SimSegment{
			drawPointRectSeg(c1X, c1Y, startRect, startDir, cornerRadius, 0),
			drawPointRectSeg(c1X, c1Y, endRect, endDir, cornerRadius, 1),
		}
	}
	if startDir == doc_models.ORIENTATION_VERTICAL && endDir == doc_models.ORIENTATION_HORIZONTAL {
		c1X := startRect.X1 + ls.StartRatio*startRect.Width()
		c1Y := endRect.Y1 + ls.EndRatio*endRect.Height()
		return []SimSegment{
			drawPointRectSeg(c1X, c1Y, startRect, startDir, cornerRadius, 0),
			drawPointRectSeg(c1X, c1Y, endRect, endDir, cornerRadius, 1),
		}
	}
	if startDir == doc_models.ORIENTATION_HORIZONTAL && endDir == doc_models.ORIENTATION_HORIZONTAL {
		c1X := startRect.X1 + ls.CornerOffsetRatio*startRect.Width()
		c1Y := startRect.Y1 + ls.StartRatio*startRect.Height()
		c2X := c1X
		c2Y := endRect.Y1 + ls.EndRatio*endRect.Height()
		if math.Abs(c1Y-c2Y) <= 2*cornerRadius {
			c2Y = c1Y
			return []SimSegment{
				drawPointRectSeg(c1X, c1Y, startRect, startDir, 0, 0),
				drawPointPointSeg(c1X, c1Y, c2X, c2Y, doc_models.ORIENTATION_HORIZONTAL, 0),
				drawPointRectSeg(c2X, c2Y, endRect, endDir, 0, 2),
			}
		}
		return []SimSegment{
			drawPointRectSeg(c1X, c1Y, startRect, startDir, cornerRadius, 0),
			drawPointPointSeg(c1X, c1Y, c2X, c2Y, doc_models.ORIENTATION_VERTICAL, cornerRadius),
			drawPointRectSeg(c2X, c2Y, endRect, endDir, cornerRadius, 2),
		}
	}
	// VERTICAL -> VERTICAL
	c1X := startRect.X1 + ls.StartRatio*startRect.Width()
	c1Y := startRect.Y1 + ls.CornerOffsetRatio*startRect.Height()
	c2X := endRect.X1 + ls.EndRatio*endRect.Width()
	c2Y := c1Y
	if math.Abs(c1X-c2X) <= 2*cornerRadius {
		c2X = c1X
		return []SimSegment{
			drawPointRectSeg(c1X, c1Y, startRect, startDir, 0, 0),
			drawPointPointSeg(c1X, c1Y, c2X, c2Y, doc_models.ORIENTATION_HORIZONTAL, 0),
			drawPointRectSeg(c2X, c2Y, endRect, endDir, 0, 2),
		}
	}
	return []SimSegment{
		drawPointRectSeg(c1X, c1Y, startRect, startDir, cornerRadius, 0),
		drawPointPointSeg(c1X, c1Y, c2X, c2Y, doc_models.ORIENTATION_HORIZONTAL, cornerRadius),
		drawPointRectSeg(c2X, c2Y, endRect, endDir, cornerRadius, 2),
	}
}

func simArrowTextBox(
	ls *doc_models.LinkShape,
	seg SimSegment,
	atEnd bool,
	anchorLeftOrTop bool,
	content string,
	userYOffset float64,
) Rect2D {
	const oneEm = 15.0
	const endArrowSize = 8.0

	orientation := ls.StartOrientation
	if atEnd {
		orientation = ls.EndOrientation
	}
	if orientation == "" {
		orientation = doc_models.ORIENTATION_HORIZONTAL
	}

	offsetSign := 1.0
	if !atEnd {
		offsetSign = -offsetSign
	}
	if orientation == doc_models.ORIENTATION_VERTICAL {
		if seg.EndY > seg.StartY {
			offsetSign = -offsetSign
		}
	} else {
		if atEnd {
			offsetSign = -offsetSign
		}
		if !anchorLeftOrTop {
			offsetSign = -offsetSign
		}
	}

	offset := endArrowSize
	if offsetSign == 1.0 {
		offset += oneEm
	} else {
		offset += oneEm * 0.4
	}
	autoY := offset*offsetSign + userYOffset

	baseX := seg.StartX
	baseY := seg.StartY
	if atEnd {
		baseX = seg.EndX
		baseY = seg.EndY
	}

	anchorEnd := false
	if seg.Orientation == doc_models.ORIENTATION_VERTICAL {
		anchorEnd = anchorLeftOrTop
	} else {
		if atEnd {
			anchorEnd = seg.EndX > seg.StartX
		} else {
			anchorEnd = !(seg.EndX > seg.StartX)
		}
	}

	w := LinkLabelWidth(content)
	var x1, x2 float64
	if anchorEnd {
		x2 = baseX - 10.0
		x1 = x2 - w
	} else {
		x1 = baseX + 10.0
		x2 = x1 + w
	}
	baselineY := baseY + autoY
	return Rect2D{
		Kind: "label",
		Name: content,
		X1:   x1,
		Y1:   baselineY - 12.5,
		X2:   x2,
		Y2:   baselineY + 2.5,
	}
}

func segmentIntersectsRectInterior(seg SimSegment, r Rect2D, margin float64) bool {
	// Shrink rect slightly by -margin (if negative) or expand by +margin
	rx1, ry1, rx2, ry2 := r.X1-margin, r.Y1-margin, r.X2+margin, r.Y2+margin
	if math.Abs(seg.StartX-seg.EndX) < 0.5 && math.Abs(seg.StartY-seg.EndY) < 0.5 {
		return false
	}
	if seg.Orientation == doc_models.ORIENTATION_HORIZONTAL || math.Abs(seg.StartY-seg.EndY) < 1.0 {
		y := 0.5 * (seg.StartY + seg.EndY)
		if y <= ry1 || y >= ry2 {
			return false
		}
		minX := math.Min(seg.StartX, seg.EndX)
		maxX := math.Max(seg.StartX, seg.EndX)
		return maxX > rx1+2.0 && minX < rx2-2.0
	}
	x := 0.5 * (seg.StartX + seg.EndX)
	if x <= rx1 || x >= rx2 {
		return false
	}
	minY := math.Min(seg.StartY, seg.EndY)
	maxY := math.Max(seg.StartY, seg.EndY)
	return maxY > ry1+2.0 && minY < ry2-2.0
}

func lineIntersectsRect(x1, y1, x2, y2 float64, r Rect2D) bool {
	// Liang-Barsky clipping test on [r.X1+2, r.X2-2] x [r.Y1+2, r.Y2-2]
	xmin, ymin, xmax, ymax := r.X1+2.0, r.Y1+2.0, r.X2-2.0, r.Y2-2.0
	if xmin >= xmax || ymin >= ymax {
		return false
	}
	dx := x2 - x1
	dy := y2 - y1
	p := [4]float64{-dx, dx, -dy, dy}
	q := [4]float64{x1 - xmin, xmax - x1, y1 - ymin, ymax - y1}
	t0, t1 := 0.0, 1.0
	for i := range 4 {
		if math.Abs(p[i]) < 1e-9 {
			if q[i] < 0 {
				return false
			}
		} else {
			t := q[i] / p[i]
			if p[i] < 0 {
				if t > t1 {
					return false
				}
				if t > t0 {
					t0 = t
				}
			} else {
				if t < t0 {
					return false
				}
				if t < t1 {
					t1 = t
				}
			}
		}
	}
	return t0 < t1
}

// DiagramMetrics summarizes the geometric quality of a Classdiagram.
type DiagramMetrics struct {
	MinX             float64
	MinY             float64
	MaxX             float64
	MaxY             float64
	Width            float64
	Height           float64
	AspectRatio      float64
	BoxOverlaps      []string
	TextClipIssues   []string
	SegmentPiercings []string
	LabelBoxHits     []string
	LabelLabelHits   []string
	NoteLinkHits     []string
}

// AnalyzeClassdiagram evaluates exact lib/doc + lib/svg geometry for a Classdiagram.
func AnalyzeClassdiagram(cd *doc_models.Classdiagram) DiagramMetrics {
	m := DiagramMetrics{
		MinX: math.MaxFloat64,
		MinY: math.MaxFloat64,
	}

	var allBoxes []Rect2D
	structRectByKey := make(map[string]Rect2D)
	structShapeByRectName := make(map[string]*doc_models.GongStructShape)

	updateBounds := func(r Rect2D) {
		if r.X1 < m.MinX {
			m.MinX = r.X1
		}
		if r.Y1 < m.MinY {
			m.MinY = r.Y1
		}
		if r.X2 > m.MaxX {
			m.MaxX = r.X2
		}
		if r.Y2 > m.MaxY {
			m.MaxY = r.Y2
		}
	}

	for _, gss := range cd.GongStructShapes {
		if gss.IsHidden {
			continue
		}
		pkg, sName := doc_models.IdentifierMetaToPackageAndGongStructName(gss.IdentifierMeta)
		key := sName
		if pkg != "" {
			key = pkg + "." + sName
		}
		r := Rect2D{
			Kind: "struct",
			Name: key,
			X1:   gss.X,
			Y1:   gss.Y,
			X2:   gss.X + gss.Width,
			Y2:   gss.Y + gss.Height,
		}
		allBoxes = append(allBoxes, r)
		structRectByKey[key] = r
		structRectByKey[sName] = r
		structShapeByRectName[key] = gss
		updateBounds(r)

		displayTitle := sName
		if pkg != "" && pkg != "models" {
			displayTitle = pkg + "." + sName
		}
		var attrLines []string
		for _, as := range gss.AttributeShapes {
			if as.IsHidden {
				continue
			}
			attrLines = append(attrLines, doc_models.IdentifierMetaToFieldName(as.IdentifierMeta)+" : "+as.Fieldtypename)
		}
		minW := StructBoxMinWidth(displayTitle, attrLines, cd.ShowNbInstances) - 18.0
		if gss.Width+0.5 < minW {
			m.TextClipIssues = append(m.TextClipIssues, fmt.Sprintf("struct %s width %.0f < %.0f", key, gss.Width, minW))
		}
		minH := math.Max(62.0, 46.0+float64(len(attrLines))*20.0)
		if gss.Height+0.5 < minH {
			m.TextClipIssues = append(m.TextClipIssues, fmt.Sprintf("struct %s height %.0f < %.0f", key, gss.Height, minH))
		}
	}

	for _, ges := range cd.GongEnumShapes {
		if ges.IsHidden {
			continue
		}
		pkg, _ := doc_models.IdentifierMetaToPackageAndGongStructName(ges.IdentifierMeta)
		eName := doc_models.GongEnumIdentifierMetaToGongEnumName(ges.IdentifierMeta)
		key := eName
		if pkg != "" {
			key = pkg + "." + eName
		}
		r := Rect2D{
			Kind: "enum",
			Name: key,
			X1:   ges.X,
			Y1:   ges.Y,
			X2:   ges.X + ges.Width,
			Y2:   ges.Y + ges.Height,
		}
		allBoxes = append(allBoxes, r)
		structRectByKey[key] = r
		structRectByKey[eName] = r
		updateBounds(r)

		displayTitle := eName
		if pkg != "" && pkg != "models" {
			displayTitle = pkg + "." + eName
		}
		var vals []string
		for _, ev := range ges.GongEnumValueShapes {
			if ev.IsHidden {
				continue
			}
			vals = append(vals, doc_models.GongEnumValueShapeIdentifierMetaToValueName(ev.IdentifierMeta))
		}
		minW, minH := EnumBoxSize(displayTitle, vals)
		if ges.Width+0.5 < minW-12.0 {
			m.TextClipIssues = append(m.TextClipIssues, fmt.Sprintf("enum %s width %.0f < %.0f", key, ges.Width, minW-12.0))
		}
		if ges.Height+0.5 < minH-4.0 {
			m.TextClipIssues = append(m.TextClipIssues, fmt.Sprintf("enum %s height %.0f < %.0f", key, ges.Height, minH-4.0))
		}
	}

	noteRectByShape := make(map[*doc_models.GongNoteShape]Rect2D)
	for _, gns := range cd.GongNoteShapes {
		if gns.IsHidden {
			continue
		}
		r := Rect2D{
			Kind: "note",
			Name: gns.Identifier,
			X1:   gns.X,
			Y1:   gns.Y,
			X2:   gns.X + gns.Width,
			Y2:   gns.Y + gns.Height,
		}
		allBoxes = append(allBoxes, r)
		noteRectByShape[gns] = r
		updateBounds(r)
	}

	for i := range len(allBoxes) {
		for j := i + 1; j < len(allBoxes); j++ {
			if rectsOverlap(allBoxes[i], allBoxes[j], 8.0) {
				m.BoxOverlaps = append(m.BoxOverlaps, fmt.Sprintf("%s(%s) overlaps %s(%s)", allBoxes[i].Kind, allBoxes[i].Name, allBoxes[j].Kind, allBoxes[j].Name))
			}
		}
	}

	type trackedLabel struct {
		ownerLink string
		isField   bool
		box       Rect2D
	}
	var allLabels []trackedLabel

	for _, gss := range cd.GongStructShapes {
		if gss.IsHidden {
			continue
		}
		sPkg, sName := doc_models.IdentifierMetaToPackageAndGongStructName(gss.IdentifierMeta)
		sKey := sName
		if sPkg != "" {
			sKey = sPkg + "." + sName
		}
		startRect := structRectByKey[sKey]

		for _, ls := range gss.LinkShapes {
			if ls.IsHidden {
				continue
			}
			ePkg, eName := doc_models.IdentifierMetaToPackageAndGongStructName(ls.FieldTypeIdentifierMeta)
			eKey := eName
			if ePkg != "" {
				eKey = ePkg + "." + eName
			}
			endRect, ok := structRectByKey[eKey]
			if !ok {
				endRect, ok = structRectByKey[eName]
			}
			if !ok {
				continue
			}

			segs := SimulateLinkSegments(startRect, endRect, ls)
			if len(segs) == 0 {
				continue
			}
			for _, seg := range segs {
				updateBounds(Rect2D{
					X1: math.Min(seg.StartX, seg.EndX),
					Y1: math.Min(seg.StartY, seg.EndY),
					X2: math.Max(seg.StartX, seg.EndX),
					Y2: math.Max(seg.StartY, seg.EndY),
				})
				for _, b := range allBoxes {
					if b.Name == startRect.Name || b.Name == endRect.Name {
						continue
					}
					if segmentIntersectsRectInterior(seg, b, 2.0) {
						m.SegmentPiercings = append(m.SegmentPiercings, fmt.Sprintf("link %s.%s->%s pierces %s", sName, ls.Name, eName, b.Name))
					}
				}
			}

			firstSeg := segs[0]
			lastSeg := segs[len(segs)-1]
			linkID := sName + "." + ls.Name

			if cd.ShowLinkNames {
				fieldName := doc_models.IdentifierMetaToFieldName(ls.IdentifierMeta)
				fb := simArrowTextBox(ls, lastSeg, true, true, fieldName, ls.FieldOffsetY)
				updateBounds(fb)
				allLabels = append(allLabels, trackedLabel{ownerLink: linkID, isField: true, box: fb})
				for _, b := range allBoxes {
					if rectsOverlap(fb, b, -1.0) {
						m.LabelBoxHits = append(m.LabelBoxHits, fmt.Sprintf("field label %s overlaps box %s", linkID, b.Name))
					}
				}
			}
			if cd.ShowMultiplicity {
				tb := simArrowTextBox(ls, lastSeg, true, false, string(ls.TargetMultiplicity), ls.TargetMultiplicityOffsetY)
				updateBounds(tb)
				allLabels = append(allLabels, trackedLabel{ownerLink: linkID, isField: false, box: tb})
				for _, b := range allBoxes {
					if rectsOverlap(tb, b, -1.0) {
						m.LabelBoxHits = append(m.LabelBoxHits, fmt.Sprintf("target mult %s overlaps box %s", linkID, b.Name))
					}
				}
				if ls.SourceMultiplicity != "" {
					sb := simArrowTextBox(ls, firstSeg, false, false, string(ls.SourceMultiplicity), ls.SourceMultiplicityOffsetY)
					updateBounds(sb)
					allLabels = append(allLabels, trackedLabel{ownerLink: linkID, isField: false, box: sb})
					for _, b := range allBoxes {
						if rectsOverlap(sb, b, -1.0) {
							m.LabelBoxHits = append(m.LabelBoxHits, fmt.Sprintf("source mult %s overlaps box %s", linkID, b.Name))
						}
					}
				}
			}
		}
	}

	for i := range len(allLabels) {
		for j := i + 1; j < len(allLabels); j++ {
			if allLabels[i].ownerLink == allLabels[j].ownerLink {
				continue
			}
			// Check any collision involving at least one field name label (or two multiplicities on same side)
			if !allLabels[i].isField && !allLabels[j].isField {
				continue
			}
			if rectsOverlap(allLabels[i].box, allLabels[j].box, -1.5) {
				m.LabelLabelHits = append(m.LabelLabelHits, fmt.Sprintf("label %s(%s) overlaps %s(%s)",
					allLabels[i].ownerLink, allLabels[i].box.Name, allLabels[j].ownerLink, allLabels[j].box.Name))
			}
		}
	}

	for gns, nr := range noteRectByShape {
		ncx := 0.5 * (nr.X1 + nr.X2)
		ncy := 0.5 * (nr.Y1 + nr.Y2)
		for _, nls := range gns.GongNoteLinkShapes {
			if nls.IsHidden || nls.Type != doc_models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE {
				continue
			}
			tName := doc_models.IdentifierToGongStructName(nls.Identifier)
			tr, ok := structRectByKey[nls.Identifier]
			if !ok {
				tr, ok = structRectByKey[tName]
			}
			if !ok {
				continue
			}
			tcx := 0.5 * (tr.X1 + tr.X2)
			tcy := 0.5 * (tr.Y1 + tr.Y2)
			for _, b := range allBoxes {
				if b.Name == nr.Name || b.Name == tr.Name {
					continue
				}
				if lineIntersectsRect(ncx, ncy, tcx, tcy, b) {
					m.NoteLinkHits = append(m.NoteLinkHits, fmt.Sprintf("note link %s->%s crosses %s", gns.Name, tr.Name, b.Name))
				}
			}
		}
	}

	if m.MinX == math.MaxFloat64 {
		m.MinX, m.MinY = 0, 0
	}
	m.Width = m.MaxX
	m.Height = m.MaxY
	if m.Height > 0 {
		m.AspectRatio = m.Width / m.Height
	}
	return m
}
