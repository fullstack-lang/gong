package stool

import (
	"bytes"
	"fmt"
	"math"
	"sort"
	"strings"
	"time"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

const (
	pdfMMToPt  = 72.0 / 25.4
	a3WidthMM  = 420.0
	a3HeightMM = 297.0

	dossierAuthors = "Valérie ROSTAING & Thomas PEUGOT"

	densityEuropeanOakKgM3    = 720.0
	densityAmericanWalnutKgM3 = 640.0
	roughBlankAllowanceMM     = 6.0
)

// encodeWinAnsi converts a UTF-8 Go string into a PDF literal string (...) encoded in Windows-1252 (WinAnsiEncoding),
// escaping parentheses and backslashes so French accents (Valérie), Ø, °, ±, ×, ², ³, — render natively in Type1 fonts.
func encodeWinAnsiLiteral(s string) string {
	var b bytes.Buffer
	b.WriteByte('(')
	for _, r := range s {
		var c byte
		switch {
		case r >= 0x20 && r <= 0x7E:
			c = byte(r)
		case r >= 0xA0 && r <= 0xFF:
			c = byte(r)
		case r == '–':
			c = 0x96
		case r == '—':
			c = 0x97
		case r == '•':
			c = 0x95
		case r == '’' || r == '‘':
			c = '\''
		case r == '“' || r == '”':
			c = '"'
		case r == '€':
			c = 0x80
		case r == 'Δ':
			// Approximate Greek Delta with crisp ASCII 'd' prefix or 'Delta '
			b.WriteString("Delta ")
			continue
		case r == 'θ':
			b.WriteString("theta")
			continue
		case r == 'α':
			b.WriteString("alpha")
			continue
		default:
			c = '?'
		}
		switch c {
		case '(', ')', '\\':
			b.WriteByte('\\')
			b.WriteByte(c)
		default:
			if c >= 0x20 && c <= 0x7E {
				b.WriteByte(c)
			} else {
				b.WriteString(fmt.Sprintf("\\%03o", c))
			}
		}
	}
	b.WriteByte(')')
	return b.String()
}

// estimateTextWidthMM estimates Helvetica / Courier string width in millimeters for accurate right/center alignment.
func estimateTextWidthMM(s string, fontSizePt float64, fontKey string) float64 {
	ptToMM := 25.4 / 72.0
	if fontKey == "F4" { // Courier monospaced
		return float64(len([]rune(s))) * 0.60 * fontSizePt * ptToMM
	}
	units := 0.0
	boldFactor := 1.0
	if fontKey == "F2" {
		boldFactor = 1.05
	}
	for _, r := range s {
		switch r {
		case ' ', ',', '.', ':', ';', 'i', 'l', 'I', '!', '|', '\'':
			units += 0.28
		case 'j', 't', 'f', 'r', '-', '(', ')', '/', '1':
			units += 0.36
		case 'm', 'w', 'M', 'W', '%', 'Ø':
			units += 0.84
		default:
			if r >= 'A' && r <= 'Z' {
				units += 0.68
			} else {
				units += 0.54
			}
		}
	}
	return units * boldFactor * fontSizePt * ptToMM
}

type pdfPageCanvas struct {
	buf bytes.Buffer
}

func mmPt(mm float64) float64 {
	return mm * pdfMMToPt
}

func (c *pdfPageCanvas) Save() {
	c.buf.WriteString("q\n")
}

func (c *pdfPageCanvas) Restore() {
	c.buf.WriteString("Q\n")
}

func (c *pdfPageCanvas) ClipRect(x, y, w, h float64) {
	c.buf.WriteString(fmt.Sprintf("%.2f %.2f %.2f %.2f re W n\n", mmPt(x), mmPt(y), mmPt(w), mmPt(h)))
}

func (c *pdfPageCanvas) ClipPolygon(pts [][2]float64) {
	if len(pts) < 3 {
		return
	}
	c.buf.WriteString(fmt.Sprintf("%.2f %.2f m\n", mmPt(pts[0][0]), mmPt(pts[0][1])))
	for i := 1; i < len(pts); i++ {
		c.buf.WriteString(fmt.Sprintf("%.2f %.2f l\n", mmPt(pts[i][0]), mmPt(pts[i][1])))
	}
	c.buf.WriteString("h W n\n")
}

func (c *pdfPageCanvas) SetStrokeRGB(r, g, b float64) {
	c.buf.WriteString(fmt.Sprintf("%.3f %.3f %.3f RG\n", r, g, b))
}

func (c *pdfPageCanvas) SetFillRGB(r, g, b float64) {
	c.buf.WriteString(fmt.Sprintf("%.3f %.3f %.3f rg\n", r, g, b))
}

func (c *pdfPageCanvas) SetLineWidthMM(wMM float64) {
	c.buf.WriteString(fmt.Sprintf("%.2f w 1 J 1 j\n", mmPt(wMM)))
}

func (c *pdfPageCanvas) SetSolid() {
	c.buf.WriteString("[] 0 d\n")
}

func (c *pdfPageCanvas) SetDashMM(onMM, offMM float64) {
	c.buf.WriteString(fmt.Sprintf("[%.2f %.2f] 0 d\n", mmPt(onMM), mmPt(offMM)))
}

func (c *pdfPageCanvas) SetCenterDashMM() {
	c.buf.WriteString(fmt.Sprintf("[%.2f %.2f %.2f %.2f] 0 d\n", mmPt(5.0), mmPt(1.5), mmPt(1.0), mmPt(1.5)))
}

func (c *pdfPageCanvas) Line(x1, y1, x2, y2 float64) {
	c.buf.WriteString(fmt.Sprintf("%.2f %.2f m %.2f %.2f l S\n", mmPt(x1), mmPt(y1), mmPt(x2), mmPt(y2)))
}

func (c *pdfPageCanvas) Rect(x, y, w, h float64, stroke, fill bool) {
	op := "n"
	if stroke && fill {
		op = "B"
	} else if fill {
		op = "f"
	} else if stroke {
		op = "S"
	}
	c.buf.WriteString(fmt.Sprintf("%.2f %.2f %.2f %.2f re %s\n", mmPt(x), mmPt(y), mmPt(w), mmPt(h), op))
}

func (c *pdfPageCanvas) Polygon(pts [][2]float64, closed, stroke, fill bool) {
	if len(pts) < 2 {
		return
	}
	c.buf.WriteString(fmt.Sprintf("%.2f %.2f m\n", mmPt(pts[0][0]), mmPt(pts[0][1])))
	for i := 1; i < len(pts); i++ {
		c.buf.WriteString(fmt.Sprintf("%.2f %.2f l\n", mmPt(pts[i][0]), mmPt(pts[i][1])))
	}
	if closed {
		c.buf.WriteString("h\n")
	}
	switch {
	case stroke && fill:
		c.buf.WriteString("B\n")
	case fill:
		c.buf.WriteString("f\n")
	case stroke:
		c.buf.WriteString("S\n")
	default:
		c.buf.WriteString("n\n")
	}
}

func (c *pdfPageCanvas) Circle(cx, cy, r float64, stroke, fill bool) {
	const k = 0.5522847498
	x, y, rad := mmPt(cx), mmPt(cy), mmPt(r)
	kr := k * rad
	c.buf.WriteString(fmt.Sprintf("%.2f %.2f m\n", x+rad, y))
	c.buf.WriteString(fmt.Sprintf("%.2f %.2f %.2f %.2f %.2f %.2f c\n", x+rad, y+kr, x+kr, y+rad, x, y+rad))
	c.buf.WriteString(fmt.Sprintf("%.2f %.2f %.2f %.2f %.2f %.2f c\n", x-kr, y+rad, x-rad, y+kr, x-rad, y))
	c.buf.WriteString(fmt.Sprintf("%.2f %.2f %.2f %.2f %.2f %.2f c\n", x-rad, y-kr, x-kr, y-rad, x, y-rad))
	c.buf.WriteString(fmt.Sprintf("%.2f %.2f %.2f %.2f %.2f %.2f c h\n", x+kr, y-rad, x+rad, y-kr, x+rad, y))
	switch {
	case stroke && fill:
		c.buf.WriteString("B\n")
	case fill:
		c.buf.WriteString("f\n")
	default:
		c.buf.WriteString("S\n")
	}
}

func (c *pdfPageCanvas) Text(xMM, yMM, sizePt float64, fontKey, text string) {
	if text == "" {
		return
	}
	c.buf.WriteString(fmt.Sprintf("BT /%s %.2f Tf %.2f %.2f Td %s Tj ET\n",
		fontKey, sizePt, mmPt(xMM), mmPt(yMM), encodeWinAnsiLiteral(text)))
}

func (c *pdfPageCanvas) TextAligned(xMM, yMM, sizePt float64, fontKey, text, align string) {
	w := estimateTextWidthMM(text, sizePt, fontKey)
	switch align {
	case "center":
		c.Text(xMM-0.5*w, yMM, sizePt, fontKey, text)
	case "right":
		c.Text(xMM-w, yMM, sizePt, fontKey, text)
	default:
		c.Text(xMM, yMM, sizePt, fontKey, text)
	}
}

func (c *pdfPageCanvas) ArrowHead(tipX, tipY, dirX, dirY, lenMM, halfWidthMM float64) {
	l := math.Hypot(dirX, dirY)
	if l < 1e-9 {
		return
	}
	ux, uy := dirX/l, dirY/l
	nx, ny := -uy, ux
	bx, by := tipX-ux*lenMM, tipY-uy*lenMM
	pts := [][2]float64{
		{tipX, tipY},
		{bx + nx*halfWidthMM, by + ny*halfWidthMM},
		{bx - nx*halfWidthMM, by - ny*halfWidthMM},
	}
	c.Polygon(pts, true, false, true)
}

// DimLinear draws a standard engineering dimension line offset perpendicular to (x1,y1)->(x2,y2).
func (c *pdfPageCanvas) DimLinear(x1, y1, x2, y2, offsetMM float64, label string) {
	dx, dy := x2-x1, y2-y1
	span := math.Hypot(dx, dy)
	if span < 0.5 {
		return
	}
	tx, ty := dx/span, dy/span
	nx, ny := -ty, tx
	sgn := 1.0
	if offsetMM < 0 {
		sgn = -1.0
	}
	d1x, d1y := x1+nx*offsetMM, y1+ny*offsetMM
	d2x, d2y := x2+nx*offsetMM, y2+ny*offsetMM

	c.Save()
	c.SetSolid()
	c.SetStrokeRGB(0.18, 0.22, 0.28)
	c.SetFillRGB(0.18, 0.22, 0.28)
	c.SetLineWidthMM(0.16)

	// Extension lines with 1.2 mm gap from object and 1.8 mm overshoot past dimension line
	c.Line(x1+nx*sgn*1.2, y1+ny*sgn*1.2, d1x+nx*sgn*1.8, d1y+ny*sgn*1.8)
	c.Line(x2+nx*sgn*1.2, y2+ny*sgn*1.2, d2x+nx*sgn*1.8, d2y+ny*sgn*1.8)
	c.Line(d1x, d1y, d2x, d2y)

	arrowLen := math.Min(2.4, 0.28*span)
	c.ArrowHead(d1x, d1y, -tx, -ty, arrowLen, 0.65)
	c.ArrowHead(d2x, d2y, tx, ty, arrowLen, 0.65)

	mx, my := 0.5*(d1x+d2x), 0.5*(d1y+d2y)
	fontSize := 7.2
	tw := estimateTextWidthMM(label, fontSize, "F2") + 2.2
	th := 3.4
	c.SetFillRGB(1.0, 1.0, 1.0)
	c.Rect(mx-0.5*tw, my-0.45*th, tw, th, false, true)
	c.SetFillRGB(0.10, 0.14, 0.20)
	c.TextAligned(mx, my-0.95, fontSize, "F2", label, "center")
	c.Restore()
}

func (c *pdfPageCanvas) LeaderCallout(tipX, tipY, elbowX, elbowY, endX float64, line1, line2 string) {
	c.Save()
	c.SetSolid()
	c.SetStrokeRGB(0.18, 0.22, 0.28)
	c.SetFillRGB(0.18, 0.22, 0.28)
	c.SetLineWidthMM(0.18)
	c.Line(tipX, tipY, elbowX, elbowY)
	c.Line(elbowX, elbowY, endX, elbowY)
	c.ArrowHead(tipX, tipY, tipX-elbowX, tipY-elbowY, 2.2, 0.65)

	align := "left"
	tx := endX + 1.2
	if endX < elbowX {
		align = "right"
		tx = endX - 1.2
	}
	c.SetFillRGB(0.10, 0.14, 0.20)
	c.TextAligned(tx, elbowY+0.8, 7.2, "F2", line1, align)
	if line2 != "" {
		c.SetFillRGB(0.32, 0.36, 0.42)
		c.TextAligned(tx, elbowY-2.4, 6.5, "F1", line2, align)
	}
	c.Restore()
}

type pdfDocument struct {
	pages []*pdfPageCanvas
	title string
}

func (d *pdfDocument) NewPage() *pdfPageCanvas {
	p := &pdfPageCanvas{}
	d.pages = append(d.pages, p)
	return p
}

func (d *pdfDocument) Bytes() []byte {
	// Object IDs:
	// 1: Catalog
	// 2: Pages
	// 3: Font F1 (Helvetica)
	// 4: Font F2 (Helvetica-Bold)
	// 5: Font F3 (Helvetica-Oblique)
	// 6: Font F4 (Courier)
	// 7: Info
	// 8 + 2*i: Page i, 9 + 2*i: Content stream i
	totalObjs := 7 + 2*len(d.pages)
	offsets := make([]int, totalObjs+1)

	var out bytes.Buffer
	out.WriteString("%PDF-1.4\n%\xE2\xE3\xCF\xD3\n")

	writeObj := func(id int, body string) {
		offsets[id] = out.Len()
		out.WriteString(fmt.Sprintf("%d 0 obj\n%s\nendobj\n", id, body))
	}

	writeObj(1, "<< /Type /Catalog /Pages 2 0 R >>")

	var kids strings.Builder
	for i := range d.pages {
		if i > 0 {
			kids.WriteByte(' ')
		}
		kids.WriteString(fmt.Sprintf("%d 0 R", 8+2*i))
	}
	writeObj(2, fmt.Sprintf("<< /Type /Pages /Kids [%s] /Count %d >>", kids.String(), len(d.pages)))

	writeObj(3, "<< /Type /Font /Subtype /Type1 /BaseFont /Helvetica /Encoding /WinAnsiEncoding >>")
	writeObj(4, "<< /Type /Font /Subtype /Type1 /BaseFont /Helvetica-Bold /Encoding /WinAnsiEncoding >>")
	writeObj(5, "<< /Type /Font /Subtype /Type1 /BaseFont /Helvetica-Oblique /Encoding /WinAnsiEncoding >>")
	writeObj(6, "<< /Type /Font /Subtype /Type1 /BaseFont /Courier /Encoding /WinAnsiEncoding >>")

	nowStr := time.Now().Format("20060102150405")
	writeObj(7, fmt.Sprintf("<< /Title %s /Author %s /Subject (Stool Industrial Manufacturing & Cost Estimation Dossier) /Creator (Phylla DSM - Gong) /CreationDate (D:%s) >>",
		encodeWinAnsiLiteral(d.title),
		encodeWinAnsiLiteral(dossierAuthors),
		nowStr,
	))

	wPt := mmPt(a3WidthMM)
	hPt := mmPt(a3HeightMM)
	resDict := "<< /Font << /F1 3 0 R /F2 4 0 R /F3 5 0 R /F4 6 0 R >> >>"

	for i, p := range d.pages {
		pageObjID := 8 + 2*i
		streamObjID := 9 + 2*i
		writeObj(pageObjID, fmt.Sprintf("<< /Type /Page /Parent 2 0 R /MediaBox [0 0 %.2f %.2f] /Resources %s /Contents %d 0 R >>",
			wPt, hPt, resDict, streamObjID))

		streamBytes := p.buf.Bytes()
		offsets[streamObjID] = out.Len()
		out.WriteString(fmt.Sprintf("%d 0 obj\n<< /Length %d >>\nstream\n", streamObjID, len(streamBytes)))
		out.Write(streamBytes)
		if len(streamBytes) == 0 || streamBytes[len(streamBytes)-1] != '\n' {
			out.WriteByte('\n')
		}
		out.WriteString("endstream\nendobj\n")
	}

	xrefPos := out.Len()
	out.WriteString(fmt.Sprintf("xref\n0 %d\n", totalObjs+1))
	out.WriteString("0000000000 65535 f \n")
	for id := 1; id <= totalObjs; id++ {
		out.WriteString(fmt.Sprintf("%010d 00000 n \n", offsets[id]))
	}
	out.WriteString(fmt.Sprintf("trailer\n<< /Size %d /Root 1 0 R /Info 7 0 R >>\nstartxref\n%d\n%%%%EOF\n", totalObjs+1, xrefPos))
	return out.Bytes()
}

// ==============================================================================
// MANUFACTURING & GEOMETRIC METRICS COMPUTATION
// ==============================================================================

type stoolDossierMetrics struct {
	PlantName      string
	RadialReps     int
	N, M           int
	SideLengthMM   float64
	InsideAngleDeg float64
	ProjAngleDeg   float64
	IsGrasshopper  bool

	OverallHeightMM    float64
	SeatBottomYMM      float64
	SeatSkirtHeightMM  float64
	LegHeightMM        float64
	SeatTopFilletMM    float64
	InnerLegOffsetMM   float64
	LegTipFilletBotMM  float64
	LegTipFilletTopMM  float64
	LegMidWidthBotMM   float64
	LegMidWidthTopMM   float64
	LegChordBotMM      float64
	LegDepthBotMM      float64

	SeatTopMaxDiamMM   float64
	SeatTopMinDiamMM   float64
	SeatBotMaxDiamMM   float64
	FootprintMaxDiamMM float64
	OverallWidthXMM    float64
	OverallDepthZMM    float64

	// Net finished volumes & surface areas
	SeatNetVolCM3      float64
	SeatAreaM2         float64
	LegsTotalNetVolCM3 float64
	LegsTotalAreaM2    float64
	SingleLegNetVolCM3 float64
	SingleLegAreaM2    float64
	TotalNetVolCM3     float64
	TotalAreaM2        float64

	// Monolithic block & zone envelopes (+6 mm rough machining allowance per face)
	SeatBlankFinLMM     float64
	SeatBlankFinWMM     float64
	SeatBlankFinTMM     float64
	MonolithBlockLMM    float64
	MonolithBlockWMM    float64
	MonolithBlockHMM    float64
	MonolithBlockVolDM3 float64
	ChipRemovalVolDM3   float64
	MaterialYieldPct    float64

	// CNC perimeter toolpath lengths
	PerimSeatTopM    float64
	PerimSeatBotM    float64
	PerimSingleFootM float64
	TotalCNCToolpathM float64

	// Stacking & packaging
	StackPitchHeightMM float64
	StackTwistDeg      float64
	SingleCartonM3     float64
	Stack4CartonM3     float64

	// Contours in CAD XY = (X, -Z) [mm]
	SeatTopCAD    [][2]float64
	SeatBotCAD    [][2]float64
	FootTop0CAD   [][2]float64
	FootBot0CAD   [][2]float64
}

func computeGeomVolumeAndArea(geom *threejs.BufferGeometry) (volMM3, areaMM2 float64) {
	if geom == nil {
		return 0, 0
	}
	var vol6 float64
	for _, f := range geom.Faces {
		if f.V1 >= len(geom.Vertices) || f.V2 >= len(geom.Vertices) || f.V3 >= len(geom.Vertices) {
			continue
		}
		v1 := geom.Vertices[f.V1]
		v2 := geom.Vertices[f.V2]
		v3 := geom.Vertices[f.V3]

		// Signed tetrahedral volume * 6
		cx := v2.Y*v3.Z - v2.Z*v3.Y
		cy := v2.Z*v3.X - v2.X*v3.Z
		cz := v2.X*v3.Y - v2.Y*v3.X
		vol6 += v1.X*cx + v1.Y*cy + v1.Z*cz

		// Triangle area = 0.5 * |(v2-v1) x (v3-v1)|
		e1x, e1y, e1z := v2.X-v1.X, v2.Y-v1.Y, v2.Z-v1.Z
		e2x, e2y, e2z := v3.X-v1.X, v3.Y-v1.Y, v3.Z-v1.Z
		nx := e1y*e2z - e1z*e2y
		ny := e1z*e2x - e1x*e2z
		nz := e1x*e2y - e1y*e2x
		areaMM2 += 0.5 * math.Sqrt(nx*nx+ny*ny+nz*nz)
	}
	return math.Abs(vol6) / 6.0, areaMM2
}

func loopPerimeterMM(pts [][2]float64) float64 {
	if len(pts) < 2 {
		return 0
	}
	sum := 0.0
	for i := range pts {
		nxt := pts[(i+1)%len(pts)]
		sum += math.Hypot(nxt[0]-pts[i][0], nxt[1]-pts[i][1])
	}
	return sum
}

// minOrientedRect2D sweeps 0..180 deg in 0.5 deg increments to find the tightest rectangular
// lumber cross-section enclosing a single tapered leg (both top and bottom contours combined).
func minOrientedRect2D(pts [][2]float64) (bestLen, bestThick, bestAngleRad float64) {
	if len(pts) == 0 {
		return 0, 0, 0
	}
	bestArea := math.MaxFloat64
	for step := range 360 {
		ang := float64(step) * 0.5 * math.Pi / 180.0
		c, s := math.Cos(ang), math.Sin(ang)
		minU, maxU := math.MaxFloat64, -math.MaxFloat64
		minV, maxV := math.MaxFloat64, -math.MaxFloat64
		for _, p := range pts {
			u := p[0]*c + p[1]*s
			v := -p[0]*s + p[1]*c
			if u < minU {
				minU = u
			}
			if u > maxU {
				maxU = u
			}
			if v < minV {
				minV = v
			}
			if v > maxV {
				maxV = v
			}
		}
		du := maxU - minU
		dv := maxV - minV
		area := du * dv
		if area < bestArea {
			bestArea = area
			if du >= dv {
				bestLen, bestThick, bestAngleRad = du, dv, ang
			} else {
				bestLen, bestThick, bestAngleRad = dv, du, ang+0.5*math.Pi
			}
		}
	}
	return bestLen, bestThick, bestAngleRad
}

func pt2DSliceToCAD(pts []pt2D) [][2]float64 {
	out := make([][2]float64, len(pts))
	for i, p := range pts {
		out[i] = [2]float64{p.X, -p.Z}
	}
	return out
}

func rotateCADLoop(pts [][2]float64, rotRad float64) [][2]float64 {
	c, s := math.Cos(rotRad), math.Sin(rotRad)
	out := make([][2]float64, len(pts))
	for i, p := range pts {
		out[i] = [2]float64{p[0]*c - p[1]*s, p[0]*s + p[1]*c}
	}
	return out
}

func computeDossierMetrics(plant *models.PlantAbstract, res *StoolGeometriesResult) *stoolDossierMetrics {
	sa := plant.StoolAbstract
	reps := max(sa.RadialRepetitions, 1)
	m := &stoolDossierMetrics{
		PlantName:      plant.Name,
		RadialReps:     reps,
		N:              plant.N,
		M:              plant.M,
		SideLengthMM:   plant.RhombusSideLength,
		InsideAngleDeg: plant.RhombusInsideAngle,
		ProjAngleDeg:   sa.ProjectionAngle,
		IsGrasshopper:  sa.IsGrasshopperMode,
		OverallHeightMM: res.StoolTopHeight,
		SeatBottomYMM:   res.SeatBottomHeight,
	}
	m.SeatSkirtHeightMM = m.OverallHeightMM - m.SeatBottomYMM
	m.LegHeightMM = m.SeatBottomYMM
	if sa.IsGrasshopperMode && sa.IsLegsReachTop {
		m.LegHeightMM = m.OverallHeightMM
	}

	seatVolMM3, seatAreaMM2 := computeGeomVolumeAndArea(res.SeatGeom)
	legsVolMM3, legsAreaMM2 := computeGeomVolumeAndArea(res.LegsGeom)
	m.SeatNetVolCM3 = seatVolMM3 * 1e-3
	m.SeatAreaM2 = seatAreaMM2 * 1e-6
	m.LegsTotalNetVolCM3 = legsVolMM3 * 1e-3
	m.LegsTotalAreaM2 = legsAreaMM2 * 1e-6
	m.SingleLegNetVolCM3 = m.LegsTotalNetVolCM3 / float64(reps)
	m.SingleLegAreaM2 = m.LegsTotalAreaM2 / float64(reps)
	m.TotalNetVolCM3 = m.SeatNetVolCM3 + m.LegsTotalNetVolCM3
	m.TotalAreaM2 = m.SeatAreaM2 + m.LegsTotalAreaM2

	// Bounding box of full stool and seat
	minSX, maxSX := math.MaxFloat64, -math.MaxFloat64
	minSZ, maxSZ := math.MaxFloat64, -math.MaxFloat64
	for _, v := range res.SeatGeom.Vertices {
		if v.X < minSX {
			minSX = v.X
		}
		if v.X > maxSX {
			maxSX = v.X
		}
		if v.Z < minSZ {
			minSZ = v.Z
		}
		if v.Z > maxSZ {
			maxSZ = v.Z
		}
	}
	m.SeatBlankFinLMM = maxSX - minSX
	m.SeatBlankFinWMM = maxSZ - minSZ
	m.SeatBlankFinTMM = m.SeatSkirtHeightMM

	minCX, maxCX := minSX, maxSX
	minCZ, maxCZ := minSZ, maxSZ
	maxFootR := 0.0
	for _, v := range res.LegsGeom.Vertices {
		if v.X < minCX {
			minCX = v.X
		}
		if v.X > maxCX {
			maxCX = v.X
		}
		if v.Z < minCZ {
			minCZ = v.Z
		}
		if v.Z > maxCZ {
			maxCZ = v.Z
		}
		if v.Y <= 1.0 {
			if r := math.Hypot(v.X, v.Z); r > maxFootR {
				maxFootR = r
			}
		}
	}
	m.OverallWidthXMM = maxCX - minCX
	m.OverallDepthZMM = maxCZ - minCZ
	m.FootprintMaxDiamMM = 2.0 * maxFootR

	if gh := res.GrasshopperContours; gh != nil {
		m.SeatTopFilletMM = gh.SeatFilletR
		m.SeatTopCAD = pt2DSliceToCAD(gh.SeatTopRing)
		m.SeatBotCAD = pt2DSliceToCAD(gh.SeatBottomRing)
		if gh.LegTopContour != nil {
			m.FootTop0CAD = pt2DSliceToCAD(gh.LegTopContour.Loop)
			m.LegTipFilletTopMM = gh.LegTopContour.AppliedFilletMM
			m.LegMidWidthTopMM = gh.LegTopContour.MidThicknessMM
			m.InnerLegOffsetMM = gh.LegTopContour.AppliedOffsetMM
		}
		if gh.LegBotContour != nil {
			m.FootBot0CAD = pt2DSliceToCAD(gh.LegBotContour.Loop)
			m.LegTipFilletBotMM = gh.LegBotContour.AppliedFilletMM
			m.LegMidWidthBotMM = gh.LegBotContour.MidThicknessMM
		}
	} else {
		// Fallback contour extraction from SeatGeom / LegsGeom horizontal slices
		for _, v := range res.SeatGeom.Vertices {
			if math.Abs(v.Y-m.OverallHeightMM) < 0.5 && !strings.Contains(v.Name, "Center") {
				m.SeatTopCAD = append(m.SeatTopCAD, [2]float64{v.X, -v.Z})
			} else if math.Abs(v.Y-m.SeatBottomYMM) < 0.5 && !strings.Contains(v.Name, "Center") {
				m.SeatBotCAD = append(m.SeatBotCAD, [2]float64{v.X, -v.Z})
			}
		}
		for _, v := range res.LegsGeom.Vertices {
			if strings.Contains(v.Name, "k0") && !strings.Contains(v.Name, "Center") {
				if math.Abs(v.Y-m.LegHeightMM) < 0.5 {
					m.FootTop0CAD = append(m.FootTop0CAD, [2]float64{v.X, -v.Z})
				} else if math.Abs(v.Y) < 0.5 {
					m.FootBot0CAD = append(m.FootBot0CAD, [2]float64{v.X, -v.Z})
				}
			}
		}
	}

	minTopR, maxTopR := math.MaxFloat64, 0.0
	for _, p := range m.SeatTopCAD {
		r := math.Hypot(p[0], p[1])
		if r < minTopR {
			minTopR = r
		}
		if r > maxTopR {
			maxTopR = r
		}
	}
	if minTopR == math.MaxFloat64 {
		minTopR = 0
	}
	m.SeatTopMaxDiamMM = 2.0 * maxTopR
	m.SeatTopMinDiamMM = 2.0 * minTopR

	maxBotR := 0.0
	for _, p := range m.SeatBotCAD {
		if r := math.Hypot(p[0], p[1]); r > maxBotR {
			maxBotR = r
		}
	}
	m.SeatBotMaxDiamMM = 2.0 * maxBotR

	botChord, botDepth, _ := minOrientedRect2D(m.FootBot0CAD)
	m.LegChordBotMM = botChord
	m.LegDepthBotMM = botDepth

	// Monolithic raw block envelope (+6 mm machining allowance on each face) for one-piece CAD/CNC carving
	m.MonolithBlockLMM = m.OverallWidthXMM + 2*roughBlankAllowanceMM
	m.MonolithBlockWMM = m.OverallDepthZMM + 2*roughBlankAllowanceMM
	m.MonolithBlockHMM = m.OverallHeightMM + 2*roughBlankAllowanceMM
	m.MonolithBlockVolDM3 = (m.MonolithBlockLMM * m.MonolithBlockWMM * m.MonolithBlockHMM) * 1e-6

	netDM3 := m.TotalNetVolCM3 * 1e-3
	m.ChipRemovalVolDM3 = math.Max(0.0, m.MonolithBlockVolDM3-netDM3)
	if m.MonolithBlockVolDM3 > 1e-6 {
		m.MaterialYieldPct = 100.0 * netDM3 / m.MonolithBlockVolDM3
	}

	m.PerimSeatTopM = loopPerimeterMM(m.SeatTopCAD) * 1e-3
	m.PerimSeatBotM = loopPerimeterMM(m.SeatBotCAD) * 1e-3
	m.PerimSingleFootM = 0.5 * (loopPerimeterMM(m.FootTop0CAD) + loopPerimeterMM(m.FootBot0CAD)) * 1e-3
	m.TotalCNCToolpathM = m.PerimSeatTopM + m.PerimSeatBotM + float64(reps)*2.0*m.PerimSingleFootM

	m.StackPitchHeightMM = res.GrowthYStep
	if m.StackPitchHeightMM <= 1.0 {
		m.StackPitchHeightMM = m.SeatSkirtHeightMM
	}
	m.StackTwistDeg = res.GrowthThetaStep * 180.0 / math.Pi

	cartonPadMM := 30.0
	cw := (m.OverallWidthXMM + cartonPadMM) * 1e-3
	cd := (m.OverallDepthZMM + cartonPadMM) * 1e-3
	ch1 := (m.OverallHeightMM + cartonPadMM) * 1e-3
	ch4 := (m.OverallHeightMM + 3.0*m.StackPitchHeightMM + cartonPadMM) * 1e-3
	m.SingleCartonM3 = cw * cd * ch1
	m.Stack4CartonM3 = cw * cd * ch4

	return m
}

// ==============================================================================
// ARCHITECTURAL FRAME, ISO TITLE BLOCK & PROJECTION HELPERS
// ==============================================================================

func drawSheetFrameAndTitleBlock(p *pdfPageCanvas, m *stoolDossierMetrics, sheetNum, totalSheets int, sheetTitle, scaleLabel string) {
	// Outer border + inner architectural drawing border (10 mm margin)
	p.SetSolid()
	p.SetStrokeRGB(0.12, 0.15, 0.20)
	p.SetLineWidthMM(0.50)
	p.Rect(10, 10, 400, 277, true, false)
	p.SetLineWidthMM(0.18)
	p.Rect(12, 12, 396, 273, true, false)

	// Grid reference ticks along border (A..F vertical, 1..8 horizontal)
	p.SetFillRGB(0.35, 0.38, 0.44)
	for i := range 8 {
		gx := 12.0 + (float64(i)+0.5)*(396.0/8.0)
		p.TextAligned(gx, 9.8-2.3, 6.0, "F2", fmt.Sprintf("%d", i+1), "center")
		p.TextAligned(gx, 287.6, 6.0, "F2", fmt.Sprintf("%d", i+1), "center")
	}
	rowLetters := []string{"F", "E", "D", "C", "B", "A"}
	for j, ltr := range rowLetters {
		gy := 12.0 + (float64(j)+0.5)*(273.0/6.0)
		p.TextAligned(8.8, gy-1.0, 6.0, "F2", ltr, "center")
		p.TextAligned(411.2, gy-1.0, 6.0, "F2", ltr, "center")
	}

	// Top header banner strip
	p.SetFillRGB(0.11, 0.15, 0.21)
	p.Rect(12, 274, 396, 11, false, true)
	p.SetFillRGB(1.0, 1.0, 1.0)
	p.Text(16, 278.0, 9.5, "F2", fmt.Sprintf("STOOL TECHNICAL & MANUFACTURING COST DOSSIER  —  SHEET %d/%d: %s", sheetNum, totalSheets, sheetTitle))
	p.SetFillRGB(0.82, 0.86, 0.92)
	p.TextAligned(404, 278.2, 8.0, "F1", fmt.Sprintf("Authors: %s", dossierAuthors), "right")

	// ISO Title Block at bottom-right (x: 228..408, y: 12..44 => 180 x 32 mm)
	tbX, tbY, tbW, tbH := 228.0, 12.0, 180.0, 32.0
	p.SetFillRGB(0.975, 0.978, 0.985)
	p.SetStrokeRGB(0.12, 0.15, 0.20)
	p.SetLineWidthMM(0.35)
	p.Rect(tbX, tbY, tbW, tbH, true, true)

	p.SetLineWidthMM(0.20)
	p.Line(tbX, tbY+10, tbX+tbW, tbY+10)
	p.Line(tbX, tbY+20, tbX+tbW, tbY+20)
	p.Line(tbX+108, tbY, tbX+108, tbY+32)
	p.Line(tbX+146, tbY, tbX+146, tbY+20)
	p.Line(tbX+54, tbY, tbX+54, tbY+10)

	// Title block labels
	p.SetFillRGB(0.42, 0.46, 0.52)
	p.Text(tbX+2.5, tbY+28.6, 5.5, "F2", "PROJECT / PART DESIGNATION")
	p.Text(tbX+110.5, tbY+28.6, 5.5, "F2", "AUTHORS / DESIGN STUDIO")
	p.Text(tbX+2.5, tbY+16.8, 5.5, "F2", "DRAWING TITLE")
	p.Text(tbX+110.5, tbY+16.8, 5.5, "F2", "SCALE")
	p.Text(tbX+148.5, tbY+16.8, 5.5, "F2", "UNITS")
	p.Text(tbX+2.5, tbY+6.8, 5.5, "F2", "MATERIAL OPTIONS")
	p.Text(tbX+56.5, tbY+6.8, 5.5, "F2", "TOLERANCE (ISO 2768)")
	p.Text(tbX+110.5, tbY+6.8, 5.5, "F2", "DATE")
	p.Text(tbX+148.5, tbY+6.8, 5.5, "F2", "SHEET")

	// Title block values
	p.SetFillRGB(0.08, 0.11, 0.16)
	p.Text(tbX+2.5, tbY+22.4, 9.2, "F2", fmt.Sprintf("STACKABLE SCULPTURAL STOOL — %s", strings.ToUpper(m.PlantName)))
	p.Text(tbX+110.5, tbY+22.8, 7.6, "F2", dossierAuthors)
	p.Text(tbX+2.5, tbY+12.2, 8.0, "F2", sheetTitle)
	p.Text(tbX+110.5, tbY+12.2, 8.2, "F2", scaleLabel)
	p.Text(tbX+148.5, tbY+12.2, 8.2, "F2", "mm")
	p.Text(tbX+2.5, tbY+2.4, 7.0, "F1", "Solid Walnut / White Oak")
	p.Text(tbX+56.5, tbY+2.4, 7.0, "F1", "±0.2 mm (CNC Woodwork)")
	p.Text(tbX+110.5, tbY+2.4, 7.2, "F1", time.Now().Format("2006-01-02"))
	p.Text(tbX+148.5, tbY+2.2, 9.0, "F2", fmt.Sprintf("%d / %d", sheetNum, totalSheets))

	// ISO First-Angle Projection symbol in lower-left corner of title block vicinity (x: 204..225, y: 14..26)
	px, py := 213.0, 19.0
	p.Save()
	p.SetStrokeRGB(0.20, 0.24, 0.30)
	p.SetLineWidthMM(0.22)
	p.Polygon([][2]float64{{px - 7, py - 2.2}, {px - 1, py - 3.8}, {px - 1, py + 3.8}, {px - 7, py + 2.2}}, true, true, false)
	p.Circle(px+5.5, py, 2.2, true, false)
	p.Circle(px+5.5, py, 3.8, true, false)
	p.SetCenterDashMM()
	p.SetLineWidthMM(0.14)
	p.Line(px-9, py, px+11, py)
	p.Line(px+5.5, py-5, px+5.5, py+5)
	p.SetSolid()
	p.SetFillRGB(0.35, 0.38, 0.44)
	p.TextAligned(px+1, py+5.2, 5.2, "F2", "ISO-E PROJECTION", "center")
	p.Restore()
}

func drawViewLabel(p *pdfPageCanvas, cx, yMM float64, title, subtitle string) {
	p.SetFillRGB(0.10, 0.14, 0.20)
	p.TextAligned(cx, yMM, 8.2, "F2", title, "center")
	p.SetStrokeRGB(0.18, 0.22, 0.28)
	p.SetLineWidthMM(0.22)
	w := math.Max(estimateTextWidthMM(title, 8.2, "F2"), 28.0)
	p.Line(cx-0.5*w, yMM-1.1, cx+0.5*w, yMM-1.1)
	if subtitle != "" {
		p.SetFillRGB(0.38, 0.42, 0.48)
		p.TextAligned(cx, yMM-4.0, 6.6, "F1", subtitle, "center")
	}
}

type shadedTri2D struct {
	pts   [3][2]float64
	depth float64
	r, g, b float64
}

// renderShadedMeshProjection renders a clean painter-sorted shaded vector view of SeatGeom & LegsGeom
// using orthographic basis vectors (rightVec, upVec, viewDir) in Three.js coordinates (X, Y, Z).
func renderShadedMeshProjection(
	p *pdfPageCanvas,
	seatGeom, legsGeom *threejs.BufferGeometry,
	cxMM, cyMM, scale float64,
	midYMM float64,
	rightVec, upVec, viewDir [3]float64,
	seatTint, legTint [3]float64,
) {
	dot3 := func(a, b [3]float64) float64 {
		return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
	}
	light1 := [3]float64{-0.52, 0.72, -0.46}
	light2 := [3]float64{0.62, 0.35, 0.70}

	var tris []shadedTri2D
	collect := func(geom *threejs.BufferGeometry, baseCol [3]float64) {
		if geom == nil {
			return
		}
		for _, f := range geom.Faces {
			if f.V1 >= len(geom.Vertices) || f.V2 >= len(geom.Vertices) || f.V3 >= len(geom.Vertices) {
				continue
			}
			v1 := geom.Vertices[f.V1]
			v2 := geom.Vertices[f.V2]
			v3 := geom.Vertices[f.V3]

			e1x, e1y, e1z := v2.X-v1.X, v2.Y-v1.Y, v2.Z-v1.Z
			e2x, e2y, e2z := v3.X-v1.X, v3.Y-v1.Y, v3.Z-v1.Z
			nx := e1y*e2z - e1z*e2y
			ny := e1z*e2x - e1x*e2z
			nz := e1x*e2y - e1y*e2x
			nl := math.Sqrt(nx*nx + ny*ny + nz*nz)
			if nl < 1e-8 {
				continue
			}
			nrm := [3]float64{nx / nl, ny / nl, nz / nl}
			// Backface cull for clean watertight solid rendering
			if dot3(nrm, viewDir) < -0.03 {
				continue
			}

			proj := func(v *threejs.Vector3) ([2]float64, float64) {
				p3 := [3]float64{v.X, v.Y - midYMM, v.Z}
				u := dot3(p3, rightVec)
				w := dot3(p3, upVec)
				d := dot3(p3, viewDir)
				return [2]float64{cxMM + u*scale, cyMM + w*scale}, d
			}

			q1, d1 := proj(v1)
			q2, d2 := proj(v2)
			q3, d3 := proj(v3)

			k1 := math.Max(0.0, dot3(nrm, light1))
			k2 := math.Max(0.0, dot3(nrm, light2))
			shade := 0.42 + 0.46*k1 + 0.18*k2
			if shade > 1.04 {
				shade = 1.04
			}

			tris = append(tris, shadedTri2D{
				pts:   [3][2]float64{q1, q2, q3},
				depth: (d1 + d2 + d3) / 3.0,
				r:     math.Min(1.0, baseCol[0]*shade),
				g:     math.Min(1.0, baseCol[1]*shade),
				b:     math.Min(1.0, baseCol[2]*shade),
			})
		}
	}

	collect(legsGeom, legTint)
	collect(seatGeom, seatTint)

	sort.Slice(tris, func(i, j int) bool {
		return tris[i].depth < tris[j].depth
	})

	p.Save()
	p.SetSolid()
	p.SetLineWidthMM(0.04)
	for _, t := range tris {
		p.SetFillRGB(t.r, t.g, t.b)
		p.SetStrokeRGB(t.r*0.92, t.g*0.92, t.b*0.92)
		p.Polygon(t.pts[:], true, true, true)
	}
	p.Restore()
}

// drawOrthographicElevation draws a crisp architectural wireframe/silhouette elevation (Front or Right Side)
// with hidden dashed seat underside & floor datum.
func drawOrthographicElevation(
	p *pdfPageCanvas,
	m *stoolDossierMetrics,
	res *StoolGeometriesResult,
	cxMM, baseFloorYMM, scale float64,
	isSideView bool,
) {
	rightVec := [3]float64{1, 0, 0}
	viewDir := [3]float64{0, 0, 1}
	if isSideView {
		rightVec = [3]float64{0, 0, -1}
		viewDir = [3]float64{1, 0, 0}
	}
	upVec := [3]float64{0, 1, 0}
	midY := 0.5 * m.OverallHeightMM
	cyMM := baseFloorYMM + midY*scale

	// Uniform monolithic tonal fill + crisp silhouette edges
	renderShadedMeshProjection(
		p, res.SeatGeom, res.LegsGeom,
		cxMM, cyMM, scale, midY,
		rightVec, upVec, viewDir,
		[3]float64{0.91, 0.88, 0.83},
		[3]float64{0.91, 0.88, 0.83},
	)

	// Overlay crisp silhouette & feature edges (boundary between front-facing and back-facing or sharp angle)
	drawFeatureEdges := func(geom *threejs.BufferGeometry, strokeW float64) {
		if geom == nil {
			return
		}
		type qKey struct{ x, y, z int64 }
		quant := func(v *threejs.Vector3) qKey {
			return qKey{int64(math.Round(v.X * 100)), int64(math.Round(v.Y * 100)), int64(math.Round(v.Z * 100))}
		}
		type uEdge struct{ a, b qKey }
		type edgeInfo struct {
			p1, p2 [2]float64
			front  int
			back   int
			n1     [3]float64
			hasN1  bool
			sharp  bool
		}
		emap := make(map[uEdge]*edgeInfo)
		proj2 := func(v *threejs.Vector3) [2]float64 {
			u := v.X*rightVec[0] + v.Z*rightVec[2]
			return [2]float64{cxMM + u*scale, baseFloorYMM + v.Y*scale}
		}

		for _, f := range geom.Faces {
			if f.V1 >= len(geom.Vertices) || f.V2 >= len(geom.Vertices) || f.V3 >= len(geom.Vertices) {
				continue
			}
			vs := [3]*threejs.Vector3{geom.Vertices[f.V1], geom.Vertices[f.V2], geom.Vertices[f.V3]}
			e1x, e1y, e1z := vs[1].X-vs[0].X, vs[1].Y-vs[0].Y, vs[1].Z-vs[0].Z
			e2x, e2y, e2z := vs[2].X-vs[0].X, vs[2].Y-vs[0].Y, vs[2].Z-vs[0].Z
			nx := e1y*e2z - e1z*e2y
			ny := e1z*e2x - e1x*e2z
			nz := e1x*e2y - e1y*e2x
			nl := math.Sqrt(nx*nx + ny*ny + nz*nz)
			if nl < 1e-8 {
				continue
			}
			nrm := [3]float64{nx / nl, ny / nl, nz / nl}
			ndv := nrm[0]*viewDir[0] + nrm[1]*viewDir[1] + nrm[2]*viewDir[2]

			for k := range 3 {
				va, vb := vs[k], vs[(k+1)%3]
				ka, kb := quant(va), quant(vb)
				key := uEdge{ka, kb}
				if kb.x < ka.x || (kb.x == ka.x && (kb.y < ka.y || (kb.y == ka.y && kb.z < ka.z))) {
					key = uEdge{kb, ka}
				}
				inf, ok := emap[key]
				if !ok {
					inf = &edgeInfo{p1: proj2(va), p2: proj2(vb)}
					emap[key] = inf
				}
				if ndv >= 0 {
					inf.front++
				} else {
					inf.back++
				}
				if !inf.hasN1 {
					inf.n1 = nrm
					inf.hasN1 = true
				} else {
					cosAng := inf.n1[0]*nrm[0] + inf.n1[1]*nrm[1] + inf.n1[2]*nrm[2]
					if cosAng < 0.72 { // > ~44 deg crease
						inf.sharp = true
					}
				}
			}
		}

		p.Save()
		p.SetSolid()
		p.SetStrokeRGB(0.10, 0.13, 0.18)
		p.SetLineWidthMM(strokeW)
		for _, inf := range emap {
			if (inf.front > 0 && inf.back > 0) || (inf.sharp && inf.front > 0) {
				p.Line(inf.p1[0], inf.p1[1], inf.p2[0], inf.p2[1])
			}
		}
		p.Restore()
	}

	drawFeatureEdges(res.LegsGeom, 0.26)
	drawFeatureEdges(res.SeatGeom, 0.30)

	// Floor datum line + vertical symmetry centerline
	halfSpan := 0.58 * math.Max(m.OverallWidthXMM, m.OverallDepthZMM) * scale
	p.Save()
	p.SetStrokeRGB(0.25, 0.28, 0.34)
	p.SetLineWidthMM(0.28)
	p.Line(cxMM-halfSpan-6, baseFloorYMM, cxMM+halfSpan+6, baseFloorYMM)
	p.SetCenterDashMM()
	p.SetStrokeRGB(0.45, 0.50, 0.58)
	p.SetLineWidthMM(0.14)
	p.Line(cxMM, baseFloorYMM-4, cxMM, baseFloorYMM+m.OverallHeightMM*scale+6)

	// Hidden dashed seat underside datum
	p.SetDashMM(1.6, 1.0)
	p.SetStrokeRGB(0.32, 0.38, 0.48)
	seatBotY := baseFloorYMM + m.SeatBottomYMM*scale
	halfBot := 0.5 * m.SeatBotMaxDiamMM * scale
	p.Line(cxMM-halfBot, seatBotY, cxMM+halfBot, seatBotY)
	p.Restore()
}

// ==============================================================================
// SHEET 1: GENERAL ARRANGEMENT & ORTHOGRAPHIC PROJECTIONS (SCALE 1:5)
// ==============================================================================

func buildSheet1Orthographic(doc *pdfDocument, m *stoolDossierMetrics, res *StoolGeometriesResult) {
	p := doc.NewPage()
	scale := 0.20 // True 1:5 scale on A3 if stool <= ~550 mm; auto-fit safely if huge
	maxDim := math.Max(m.OverallHeightMM, math.Max(m.OverallWidthXMM, m.FootprintMaxDiamMM))
	scaleLabel := "1 : 5"
	if maxDim > 560.0 {
		scale = 105.0 / maxDim
		scaleLabel = fmt.Sprintf("1 : %.1f", 1.0/scale)
	}

	drawSheetFrameAndTitleBlock(p, m, 1, 4, "GENERAL ARRANGEMENT & ORTHOGRAPHIC PROJECTIONS", scaleLabel)

	// 1. TOP PLAN VIEW (upper-left quadrant: center at x=88, y=202)
	topCX, topCY := 88.0, 202.0
	p.Save()
	// Hidden dashed feet bottom & feet top
	p.SetDashMM(1.4, 0.9)
	p.SetLineWidthMM(0.16)
	for k := range m.RadialReps {
		rot := -float64(k) * 2.0 * math.Pi / float64(m.RadialReps)
		fb := rotateCADLoop(m.FootBot0CAD, rot)
		ft := rotateCADLoop(m.FootTop0CAD, rot)
		ptsB := make([][2]float64, len(fb))
		for i, q := range fb {
			ptsB[i] = [2]float64{topCX + q[0]*scale, topCY + q[1]*scale}
		}
		ptsT := make([][2]float64, len(ft))
		for i, q := range ft {
			ptsT[i] = [2]float64{topCX + q[0]*scale, topCY + q[1]*scale}
		}
		p.SetStrokeRGB(0.58, 0.30, 0.48)
		p.Polygon(ptsB, true, true, false)
		p.SetStrokeRGB(0.65, 0.35, 0.28)
		p.Polygon(ptsT, true, true, false)
	}
	// Hidden dashed seat bottom
	if len(m.SeatBotCAD) >= 3 {
		ptsSB := make([][2]float64, len(m.SeatBotCAD))
		for i, q := range m.SeatBotCAD {
			ptsSB[i] = [2]float64{topCX + q[0]*scale, topCY + q[1]*scale}
		}
		p.SetStrokeRGB(0.22, 0.48, 0.34)
		p.Polygon(ptsSB, true, true, false)
	}
	// Visible solid seat top contour with subtle warm fill
	if len(m.SeatTopCAD) >= 3 {
		p.SetSolid()
		ptsST := make([][2]float64, len(m.SeatTopCAD))
		for i, q := range m.SeatTopCAD {
			ptsST[i] = [2]float64{topCX + q[0]*scale, topCY + q[1]*scale}
		}
		p.SetFillRGB(0.95, 0.92, 0.87)
		p.SetStrokeRGB(0.08, 0.11, 0.16)
		p.SetLineWidthMM(0.36)
		p.Polygon(ptsST, true, true, true)
	}
	// Symmetry axes (3 x 120 deg)
	p.SetCenterDashMM()
	p.SetStrokeRGB(0.45, 0.50, 0.58)
	p.SetLineWidthMM(0.14)
	axisR := 0.56 * math.Max(m.FootprintMaxDiamMM, m.SeatTopMaxDiamMM) * scale
	p.Line(topCX-axisR, topCY, topCX+axisR, topCY)
	p.Line(topCX, topCY-axisR, topCX, topCY+axisR)
	for k := range m.RadialReps {
		ang := float64(k) * 2.0 * math.Pi / float64(m.RadialReps)
		p.Line(topCX, topCY, topCX+axisR*math.Cos(ang), topCY+axisR*math.Sin(ang))
	}
	p.Restore()

	rTopMaxS := 0.5 * m.SeatTopMaxDiamMM * scale
	rFootMaxS := 0.5 * m.FootprintMaxDiamMM * scale
	p.DimLinear(topCX-rTopMaxS, topCY+rTopMaxS, topCX+rTopMaxS, topCY+rTopMaxS, 8.0, fmt.Sprintf("Seat Top Max Ø %.1f", m.SeatTopMaxDiamMM))
	if rFootMaxS > 1.0 {
		p.DimLinear(topCX-rFootMaxS, topCY-rFootMaxS, topCX+rFootMaxS, topCY-rFootMaxS, -8.5, fmt.Sprintf("Floor Footprint Ø %.1f", m.FootprintMaxDiamMM))
	}
	drawViewLabel(p, topCX, topCY-math.Max(rFootMaxS, rTopMaxS)-16.0, "TOP PLAN VIEW", fmt.Sprintf("Scale %s  |  %d-Fold Radial Symmetry (120°)", scaleLabel, m.RadialReps))

	// 2. FRONT ELEVATION (lower-left: floor at y=30, center at x=88)
	frontCX, floorY := 88.0, 28.0
	drawOrthographicElevation(p, m, res, frontCX, floorY, scale, false)

	hS := m.OverallHeightMM * scale
	sbS := m.SeatBottomYMM * scale
	halfW := 0.5 * m.OverallWidthXMM * scale
	p.DimLinear(frontCX-halfW, floorY, frontCX-halfW, floorY+hS, -11.0, fmt.Sprintf("H = %.1f", m.OverallHeightMM))
	p.DimLinear(frontCX+halfW, floorY+sbS, frontCX+halfW, floorY+hS, 8.5, fmt.Sprintf("Seat %.1f", m.SeatSkirtHeightMM))
	p.DimLinear(frontCX+halfW, floorY, frontCX+halfW, floorY+sbS, 8.5, fmt.Sprintf("Clear %.1f", m.SeatBottomYMM))
	if m.SeatTopFilletMM > 0.1 {
		p.LeaderCallout(frontCX-0.36*m.SeatTopMaxDiamMM*scale, floorY+hS-0.4, frontCX-0.44*m.SeatTopMaxDiamMM*scale, floorY+hS+7.5, frontCX-halfW+2.0,
			fmt.Sprintf("Top Fillet R %.1f mm", m.SeatTopFilletMM),
			fmt.Sprintf("Draft Slope %.1f°", m.ProjAngleDeg))
	}
	drawViewLabel(p, frontCX, floorY-7.5, "FRONT ELEVATION", fmt.Sprintf("Scale %s  |  Projection Draft Angle %.1f°", scaleLabel, m.ProjAngleDeg))

	// 3. RIGHT SIDE ELEVATION (middle column: floor at y=52, center at x=204 — positioned cleanly right of Front Elevation)
	sideCX, sideFloorY := 204.0, 52.0
	drawOrthographicElevation(p, m, res, sideCX, sideFloorY, scale, true)
	halfD := 0.5 * m.OverallDepthZMM * scale
	p.DimLinear(sideCX-halfD, sideFloorY+hS, sideCX+halfD, sideFloorY+hS, 8.5, fmt.Sprintf("Depth %.1f", m.OverallDepthZMM))
	drawViewLabel(p, sideCX, sideFloorY-7.5, "RIGHT SIDE ELEVATION", fmt.Sprintf("Scale %s", scaleLabel))

	// 4. SHADED ISOMETRIC 3D PRESENTATION VIEW (upper-right: center at x=316, y=170)
	isoCX, isoCY := 320.0, 168.0
	cos30, sin30 := math.Cos(math.Pi/6), math.Sin(math.Pi/6)
	isoRight := [3]float64{cos30, 0, -cos30}
	isoUp := [3]float64{-sin30 * 0.58, 0.84, -sin30 * 0.58}
	isoView := [3]float64{0.56, 0.52, 0.64}
	renderShadedMeshProjection(
		p, res.SeatGeom, res.LegsGeom,
		isoCX, isoCY, scale*1.08, 0.5*m.OverallHeightMM,
		isoRight, isoUp, isoView,
		[3]float64{0.76, 0.62, 0.45}, // monolithic matte oak tone
		[3]float64{0.76, 0.62, 0.45},
	)
	drawViewLabel(p, isoCX, 64.0, "ISOMETRIC 3D MONOLITHIC VIEW", fmt.Sprintf("One-Piece Net Volume: %.2f dm³  |  Surface Area: %.3f m²", m.TotalNetVolCM3*1e-3, m.TotalAreaM2))

	// Summary key-data box between top & side view (x: 162..246, y: 205..268)
	bx, by, bw, bh := 158.0, 204.0, 92.0, 64.0
	p.SetFillRGB(0.97, 0.975, 0.985)
	p.SetStrokeRGB(0.20, 0.24, 0.30)
	p.SetLineWidthMM(0.24)
	p.Rect(bx, by, bw, bh, true, true)
	p.SetFillRGB(0.14, 0.18, 0.25)
	p.Rect(bx, by+bh-7.5, bw, 7.5, false, true)
	p.SetFillRGB(1, 1, 1)
	p.Text(bx+3, by+bh-5.2, 7.2, "F2", "KEY DIMENSIONAL & MASS SUMMARY")

	rows := [][2]string{
		{"Overall Envelope (W × D × H):", fmt.Sprintf("%.0f × %.0f × %.0f mm", m.OverallWidthXMM, m.OverallDepthZMM, m.OverallHeightMM)},
		{"Seat Top Max / Min Ø:", fmt.Sprintf("%.1f / %.1f mm", m.SeatTopMaxDiamMM, m.SeatTopMinDiamMM)},
		{"Seat Thickness / Fillet R:", fmt.Sprintf("%.1f mm / R %.1f mm", m.SeatSkirtHeightMM, m.SeatTopFilletMM)},
		{"Projection Draft Angle:", fmt.Sprintf("%.1f°", m.ProjAngleDeg)},
		{"Leg Fillet / Inner Offset:", fmt.Sprintf("R %.2f mm / %.2f mm", m.LegTipFilletBotMM, m.InnerLegOffsetMM)},
		{"Stacking Pitch (ΔH / Δθ):", fmt.Sprintf("+%.1f mm / %.1f°", m.StackPitchHeightMM, m.StackTwistDeg)},
		{"Est. Mass — White Oak:", fmt.Sprintf("%.2f kg", (m.TotalNetVolCM3*1e-6)*densityEuropeanOakKgM3)},
		{"Est. Mass — Black Walnut:", fmt.Sprintf("%.2f kg", (m.TotalNetVolCM3*1e-6)*densityAmericanWalnutKgM3)},
	}
	for idx, r := range rows {
		ry := by + bh - 13.0 - float64(idx)*6.4
		p.SetFillRGB(0.32, 0.36, 0.42)
		p.Text(bx+3, ry, 6.5, "F1", r[0])
		p.SetFillRGB(0.08, 0.11, 0.16)
		p.TextAligned(bx+bw-3, ry, 6.7, "F2", r[1], "right")
	}
}

// ==============================================================================
// SHEET 2: BOTTOM PLAN, MONOLITHIC SECTION A-A & HELICAL STACKING PITCH
// ==============================================================================

func buildSheet2SectionAndStacking(doc *pdfDocument, m *stoolDossierMetrics, res *StoolGeometriesResult) {
	p := doc.NewPage()
	scale := 0.20
	maxDim := math.Max(m.OverallHeightMM, m.FootprintMaxDiamMM)
	if maxDim > 560.0 {
		scale = 105.0 / maxDim
	}
	scaleLabel := fmt.Sprintf("1 : %.0f", math.Round(1.0/scale))
	drawSheetFrameAndTitleBlock(p, m, 2, 4, "BOTTOM PLAN VIEW, MONOLITHIC SECTION A-A & STACKING", scaleLabel)

	// 1. BOTTOM PLAN VIEW (Underside Contours — No Joinery Assumed, x=84, y=158)
	botCX, botCY := 84.0, 158.0
	p.Save()
	if len(m.SeatTopCAD) >= 3 {
		p.SetDashMM(1.5, 1.0)
		p.SetStrokeRGB(0.45, 0.50, 0.58)
		p.SetLineWidthMM(0.16)
		pts := make([][2]float64, len(m.SeatTopCAD))
		for i, q := range m.SeatTopCAD {
			pts[i] = [2]float64{botCX - q[0]*scale, botCY + q[1]*scale}
		}
		p.Polygon(pts, true, true, false)
	}
	if len(m.SeatBotCAD) >= 3 {
		p.SetSolid()
		p.SetFillRGB(0.95, 0.92, 0.87)
		p.SetStrokeRGB(0.12, 0.16, 0.22)
		p.SetLineWidthMM(0.32)
		pts := make([][2]float64, len(m.SeatBotCAD))
		for i, q := range m.SeatBotCAD {
			pts[i] = [2]float64{botCX - q[0]*scale, botCY + q[1]*scale}
		}
		p.Polygon(pts, true, true, true)
	}
	// Draw the 3 crescent leg floor footprints & root transition contours (continuous one-piece geometry)
	for k := range m.RadialReps {
		rot := -float64(k) * 2.0 * math.Pi / float64(m.RadialReps)
		ft := rotateCADLoop(m.FootTop0CAD, rot)
		fb := rotateCADLoop(m.FootBot0CAD, rot)
		if len(fb) >= 3 {
			ptsB := make([][2]float64, len(fb))
			for i, q := range fb {
				ptsB[i] = [2]float64{botCX - q[0]*scale, botCY + q[1]*scale}
			}
			p.SetSolid()
			p.SetFillRGB(0.86, 0.80, 0.72)
			p.SetStrokeRGB(0.22, 0.18, 0.14)
			p.SetLineWidthMM(0.24)
			p.Polygon(ptsB, true, true, true)
		}
		if len(ft) >= 3 {
			ptsT := make([][2]float64, len(ft))
			for i, q := range ft {
				ptsT[i] = [2]float64{botCX - q[0]*scale, botCY + q[1]*scale}
			}
			p.SetDashMM(1.2, 0.8)
			p.SetStrokeRGB(0.32, 0.24, 0.16)
			p.SetLineWidthMM(0.24)
			p.Polygon(ptsT, true, true, false)
		}
	}
	p.Restore()

	// Section cutting plane line A-A across Bottom Plan
	secR := 0.58 * m.FootprintMaxDiamMM * scale
	p.Save()
	p.SetCenterDashMM()
	p.SetStrokeRGB(0.75, 0.16, 0.16)
	p.SetLineWidthMM(0.28)
	p.Line(botCX-secR-5, botCY, botCX+secR+5, botCY)
	p.SetSolid()
	p.SetFillRGB(0.75, 0.16, 0.16)
	p.Text(botCX-secR-9, botCY+1.5, 8.5, "F2", "A")
	p.Text(botCX+secR+6, botCY+1.5, 8.5, "F2", "A")
	p.Restore()

	drawViewLabel(p, botCX, botCY-secR-12.0, "BOTTOM PLAN VIEW (UNDERSIDE CONTOURS)",
		fmt.Sprintf("Scale %s  |  %d-Fold Radial Symmetry (120°)  |  Monolithic One-Piece Form", scaleLabel, m.RadialReps))

	// 2. MONOLITHIC AXIAL VERTICAL SECTION A-A (center column: x=196, floor at y=68)
	secCX, secFloorY := 196.0, 68.0
	secScale := scale * 1.18
	rSeatTop := 0.5 * m.SeatTopMaxDiamMM
	rSeatBot := 0.5 * m.SeatBotMaxDiamMM
	filletR := math.Min(m.SeatTopFilletMM, 0.4*m.SeatSkirtHeightMM)

	legTopOuterR := rSeatBot - 2.0
	legWallTop := math.Max(12.0, m.LegMidWidthTopMM)
	legWallBot := math.Max(14.0, m.LegMidWidthBotMM)
	tanDraft := math.Tan(m.ProjAngleDeg * math.Pi / 180.0)
	legBotOuterR := legTopOuterR + m.SeatBottomYMM*tanDraft

	// Build ONE SINGLE UNIFIED MONOLITHIC cross-section polygon (seat + sectioned right leg with zero seam!)
	var monolithSecPts [][2]float64
	monolithSecPts = append(monolithSecPts, [2]float64{secCX - rSeatBot*secScale, secFloorY + m.SeatBottomYMM*secScale})
	if filletR > 0.2 {
		steps := 8
		for s := 0; s <= steps; s++ {
			ang := float64(s) * (0.5 * math.Pi) / float64(steps)
			rx := -(rSeatTop - filletR + filletR*math.Cos(ang))
			ry := m.OverallHeightMM - filletR + filletR*math.Sin(ang)
			monolithSecPts = append(monolithSecPts, [2]float64{secCX + rx*secScale, secFloorY + ry*secScale})
		}
		for s := steps; s >= 0; s-- {
			ang := float64(s) * (0.5 * math.Pi) / float64(steps)
			rx := +(rSeatTop - filletR + filletR*math.Cos(ang))
			ry := m.OverallHeightMM - filletR + filletR*math.Sin(ang)
			monolithSecPts = append(monolithSecPts, [2]float64{secCX + rx*secScale, secFloorY + ry*secScale})
		}
	} else {
		monolithSecPts = append(monolithSecPts,
			[2]float64{secCX - rSeatTop*secScale, secFloorY + m.OverallHeightMM*secScale},
			[2]float64{secCX + rSeatTop*secScale, secFloorY + m.OverallHeightMM*secScale},
		)
	}
	// Right seat skirt bottom overhang -> right leg outer wall -> floor -> right leg inner wall -> seat underside
	monolithSecPts = append(monolithSecPts,
		[2]float64{secCX + rSeatBot*secScale, secFloorY + m.SeatBottomYMM*secScale},
		[2]float64{secCX + legTopOuterR*secScale, secFloorY + m.SeatBottomYMM*secScale},
		[2]float64{secCX + legBotOuterR*secScale, secFloorY},
		[2]float64{secCX + (legBotOuterR-legWallBot)*secScale, secFloorY},
		[2]float64{secCX + (legTopOuterR-legWallTop)*secScale, secFloorY + m.SeatBottomYMM*secScale},
	)

	// Back leg shown in unsectioned background elevation on left
	legElevLeft := [][2]float64{
		{secCX - (legTopOuterR-legWallTop)*secScale, secFloorY + m.SeatBottomYMM*secScale},
		{secCX - legTopOuterR*secScale, secFloorY + m.SeatBottomYMM*secScale},
		{secCX - legBotOuterR*secScale, secFloorY},
		{secCX - (legBotOuterR-legWallBot)*secScale, secFloorY},
	}

	p.Save()
	p.SetFillRGB(0.90, 0.87, 0.82)
	p.SetStrokeRGB(0.22, 0.25, 0.30)
	p.SetLineWidthMM(0.22)
	p.Polygon(legElevLeft, true, true, true)

	// Continuous 45° solid-material section hatching across the unified monolithic polygon
	p.SetFillRGB(0.95, 0.91, 0.84)
	p.SetStrokeRGB(0.08, 0.11, 0.16)
	p.SetLineWidthMM(0.34)
	p.Polygon(monolithSecPts, true, true, true)

	minX, maxX := math.MaxFloat64, -math.MaxFloat64
	minY, maxY := math.MaxFloat64, -math.MaxFloat64
	for _, q := range monolithSecPts {
		if q[0] < minX {
			minX = q[0]
		}
		if q[0] > maxX {
			maxX = q[0]
		}
		if q[1] < minY {
			minY = q[1]
		}
		if q[1] > maxY {
			maxY = q[1]
		}
	}
	p.Save()
	p.ClipPolygon(monolithSecPts)
	cMinX, cMinY := minX-2.0, minY-2.0
	cW, cH := (maxX-minX)+4.0, (maxY-minY)+4.0
	if cW > 0.5 && cH > 0.5 {
		p.SetStrokeRGB(0.42, 0.32, 0.22)
		p.SetLineWidthMM(0.12)
		span := cW + cH
		for d := -cH; d <= span; d += 2.2 {
			p.Line(cMinX+d, cMinY, cMinX+d+cH, cMinY+cH)
		}
	}
	p.Restore()
	p.SetStrokeRGB(0.08, 0.11, 0.16)
	p.SetLineWidthMM(0.36)
	p.Polygon(monolithSecPts, true, true, false)
	p.Restore()

	rootX := secCX + (legTopOuterR-0.5*legWallTop)*secScale
	rootY := secFloorY + m.SeatBottomYMM*secScale
	p.LeaderCallout(rootX, rootY+3*secScale, secCX+legTopOuterR*secScale+4, rootY+16, secCX+legTopOuterR*secScale+9,
		"Monolithic One-Piece Transition", "(Continuous 5-Axis CNC Carving — No Assembly)")
	p.LeaderCallout(secCX+(legBotOuterR-legWallBot)*secScale+1.0, secFloorY+26, secCX+16, secFloorY+34, secCX+6,
		fmt.Sprintf("Mid-Leg Thickness %.1f mm", m.LegMidWidthBotMM),
		fmt.Sprintf("Conical Slope α = %.1f°", m.ProjAngleDeg))
	p.DimLinear(secCX-legBotOuterR*secScale, secFloorY, secCX-legBotOuterR*secScale, secFloorY+m.OverallHeightMM*secScale, -8.5, fmt.Sprintf("%.1f", m.OverallHeightMM))
	drawViewLabel(p, secCX, secFloorY-8.5, "MONOLITHIC AXIAL SECTION A-A", "Continuous One-Piece Solid Geometry (Zero Joinery)")

	// 3. STACKED CONFIGURATION (2x NESTED STOOLS ELEVATION, right column: x=346, floor at y=68)
	stCX, stFloorY := 346.0, 68.0
	stScale := scale * 0.88
	stageStack := threejs.NewStage("pdf_stack")
	var upperSeat, upperLegs *threejs.BufferGeometry
	if gh := res.GrasshopperContours; gh != nil {
		upperSeat = buildGrasshopperSeatGeom(stageStack, "UpperSeat", gh, res.GrowthThetaStep, m.StackPitchHeightMM)
		upperLegs = buildGrasshopperLegsGeom(stageStack, "UpperLegs", gh, res.GrowthThetaStep, m.StackPitchHeightMM)
	}

	renderShadedMeshProjection(
		p, res.SeatGeom, res.LegsGeom,
		stCX, stFloorY+0.5*m.OverallHeightMM*stScale, stScale, 0.5*m.OverallHeightMM,
		[3]float64{1, 0, 0}, [3]float64{0, 1, 0}, [3]float64{0, 0.18, 0.98},
		[3]float64{0.58, 0.43, 0.29}, // monolithic walnut lower stool
		[3]float64{0.58, 0.43, 0.29},
	)
	if upperSeat != nil && upperLegs != nil {
		renderShadedMeshProjection(
			p, upperSeat, upperLegs,
			stCX, stFloorY+(0.5*m.OverallHeightMM+m.StackPitchHeightMM)*stScale, stScale, 0.5*m.OverallHeightMM+m.StackPitchHeightMM,
			[3]float64{1, 0, 0}, [3]float64{0, 1, 0}, [3]float64{0, 0.18, 0.98},
			[3]float64{0.82, 0.70, 0.54}, // monolithic oak upper stool
			[3]float64{0.82, 0.70, 0.54},
		)
	}
	halfStW := 0.52 * m.OverallWidthXMM * stScale
	total2H := (m.OverallHeightMM + m.StackPitchHeightMM) * stScale
	h1S := m.OverallHeightMM * stScale
	p.DimLinear(stCX+halfStW, stFloorY+h1S, stCX+halfStW, stFloorY+total2H, 7.5, fmt.Sprintf("Pitch +%.1f", m.StackPitchHeightMM))
	p.DimLinear(stCX-halfStW, stFloorY, stCX-halfStW, stFloorY+total2H, -10.5, fmt.Sprintf("2× Stack %.1f", m.OverallHeightMM+m.StackPitchHeightMM))

	drawViewLabel(p, stCX, stFloorY-8.5, "HELICAL STACKING ELEVATION (2× STOOLS)",
		fmt.Sprintf("Pitch +%.1f mm  |  Twist %.1f°  |  4× Stack Height %.0f mm", m.StackPitchHeightMM, m.StackTwistDeg, m.OverallHeightMM+3*m.StackPitchHeightMM))
}

// ==============================================================================
// SHEET 3: 2D CNC CUTTING CONTOURS (Z=0) & HIGH-ZOOM LEG FILLET DETAIL
// ==============================================================================

func buildSheet3CNCContoursAndFilletDetail(doc *pdfDocument, m *stoolDossierMetrics) {
	p := doc.NewPage()
	drawSheetFrameAndTitleBlock(p, m, 3, 4, "2D CNC CUTTING CONTOURS & LEG TIP FILLET DETAIL", "1 : 4  &  1 : 1")

	// 1. Left half: All 4 planar Z=0 cutting contour layers (matching 2D DXF export) at 1:4 scale
	cncCX, cncCY := 106.0, 148.0
	cncScale := 0.25
	if m.FootprintMaxDiamMM > 520.0 {
		cncScale = 130.0 / m.FootprintMaxDiamMM
	}

	p.Save()
	p.SetCenterDashMM()
	p.SetStrokeRGB(0.55, 0.60, 0.68)
	p.SetLineWidthMM(0.14)
	rGrid := 0.56 * math.Max(m.FootprintMaxDiamMM, m.SeatTopMaxDiamMM) * cncScale
	p.Line(cncCX-rGrid, cncCY, cncCX+rGrid, cncCY)
	p.Line(cncCX, cncCY-rGrid, cncCX, cncCY+rGrid)

	p.SetSolid()
	drawLayerLoop := func(pts [][2]float64, r, g, b, wMM float64) {
		if len(pts) < 3 {
			return
		}
		scaled := make([][2]float64, len(pts))
		for i, q := range pts {
			scaled[i] = [2]float64{cncCX + q[0]*cncScale, cncCY + q[1]*cncScale}
		}
		p.SetStrokeRGB(r, g, b)
		p.SetLineWidthMM(wMM)
		p.Polygon(scaled, true, true, false)
	}

	// Layers matching DXF export: SEAT_TOP (blue), SEAT_BOTTOM (green), FEET_TOP (red), FEET_BOTTOM (magenta)
	for k := range m.RadialReps {
		rot := -float64(k) * 2.0 * math.Pi / float64(m.RadialReps)
		drawLayerLoop(rotateCADLoop(m.FootBot0CAD, rot), 0.68, 0.16, 0.62, 0.28)
		drawLayerLoop(rotateCADLoop(m.FootTop0CAD, rot), 0.82, 0.18, 0.16, 0.28)
	}
	drawLayerLoop(m.SeatBotCAD, 0.12, 0.56, 0.26, 0.32)
	drawLayerLoop(m.SeatTopCAD, 0.10, 0.34, 0.78, 0.38)
	p.Restore()

	// Layer color legend box under left plot
	lx, ly := 28.0, 20.0
	legend := []struct {
		name string
		col  [3]float64
		info string
	}{
		{"SEAT_TOP", [3]float64{0.10, 0.34, 0.78}, fmt.Sprintf("Z = %.1f mm  |  Perimeter %.3f m", m.OverallHeightMM, m.PerimSeatTopM)},
		{"SEAT_BOTTOM", [3]float64{0.12, 0.56, 0.26}, fmt.Sprintf("Z = %.1f mm  |  Perimeter %.3f m", m.SeatBottomYMM, m.PerimSeatBotM)},
		{"FEET_TOP (3×)", [3]float64{0.82, 0.18, 0.16}, fmt.Sprintf("Z = %.1f mm  |  3 × %.3f m", m.LegHeightMM, m.PerimSingleFootM)},
		{"FEET_BOTTOM (3×)", [3]float64{0.68, 0.16, 0.62}, fmt.Sprintf("Z = 0.0 mm  |  Inner Offset %.2f mm", m.InnerLegOffsetMM)},
	}
	p.SetFillRGB(0.98, 0.985, 0.99)
	p.SetStrokeRGB(0.22, 0.26, 0.32)
	p.SetLineWidthMM(0.20)
	p.Rect(lx, ly, 160, 27, true, true)
	for i, lg := range legend {
		ry := ly + 21.5 - float64(i)*5.8
		p.SetStrokeRGB(lg.col[0], lg.col[1], lg.col[2])
		p.SetLineWidthMM(0.75)
		p.Line(lx+4, ry+1.2, lx+14, ry+1.2)
		p.SetFillRGB(0.10, 0.14, 0.20)
		p.Text(lx+17, ry, 6.8, "F2", lg.name)
		p.SetFillRGB(0.32, 0.36, 0.42)
		p.Text(lx+62, ry, 6.6, "F1", lg.info)
	}
	drawViewLabel(p, cncCX, 258.0, "DXF / CNC Z=0 PROJECTED CONTOUR SET",
		fmt.Sprintf("Scale 1 : %.1f  |  Total CNC Contour Toolpath = %.2f m", 1.0/cncScale, m.TotalCNCToolpathM))

	// 2. Upper-right: DETAIL B — Single Crescent Foot Cross-Section & Oriented Blank Envelope (Scale ~1:2)
	if len(m.FootBot0CAD) >= 6 {
		detCX, detCY := 312.0, 206.0
		_, _, angOpt := minOrientedRect2D(m.FootBot0CAD)
		cA, sA := math.Cos(-angOpt), math.Sin(-angOpt)

		var aligned [][2]float64
		minU, maxU := math.MaxFloat64, -math.MaxFloat64
		minV, maxV := math.MaxFloat64, -math.MaxFloat64
		for _, q := range m.FootBot0CAD {
			u := q[0]*cA - q[1]*sA
			v := q[0]*sA + q[1]*cA
			aligned = append(aligned, [2]float64{u, v})
			if u < minU {
				minU = u
			}
			if u > maxU {
				maxU = u
			}
			if v < minV {
				minV = v
			}
			if v > maxV {
				maxV = v
			}
		}
		uMid, vMid := 0.5*(minU+maxU), 0.5*(minV+maxV)
		spanU := math.Max(maxU-minU, 1.0)
		detScale := math.Min(0.68, 128.0/spanU)

		// Oriented bounding box of the single foot cross-section
		bW := (maxU - minU) * detScale
		bH := (maxV - minV) * detScale
		p.Save()
		p.SetDashMM(1.8, 1.1)
		p.SetStrokeRGB(0.55, 0.58, 0.64)
		p.SetLineWidthMM(0.18)
		p.Rect(detCX-0.5*bW, detCY-0.5*bH, bW, bH, true, false)

		p.SetSolid()
		polyB := make([][2]float64, len(aligned))
		for i, q := range aligned {
			polyB[i] = [2]float64{detCX + (q[0]-uMid)*detScale, detCY + (q[1]-vMid)*detScale}
		}
		p.SetFillRGB(0.93, 0.88, 0.80)
		p.SetStrokeRGB(0.10, 0.13, 0.18)
		p.SetLineWidthMM(0.36)
		p.Polygon(polyB, true, true, true)
		p.Restore()

		p.DimLinear(detCX-0.5*bW, detCY-0.5*bH, detCX+0.5*bW, detCY-0.5*bH, -7.5,
			fmt.Sprintf("Chord %.1f mm", maxU-minU))
		p.DimLinear(detCX+0.5*bW, detCY-0.5*bH, detCX+0.5*bW, detCY+0.5*bH, 8.5,
			fmt.Sprintf("Depth %.1f mm", maxV-minV))

		// Circle callout around right tip for Detail C
		tipPt := polyB[0]
		for _, q := range polyB {
			if q[0] > tipPt[0] {
				tipPt = q
			}
		}
		p.Save()
		p.SetDashMM(1.2, 0.8)
		p.SetStrokeRGB(0.78, 0.16, 0.16)
		p.SetLineWidthMM(0.24)
		p.Circle(tipPt[0], tipPt[1], 6.5, true, false)
		p.SetSolid()
		p.SetFillRGB(0.78, 0.16, 0.16)
		p.Text(tipPt[0]+5.5, tipPt[1]+5.5, 7.2, "F2", "DETAIL C")
		p.Restore()

		drawViewLabel(p, detCX, 258.0, "DETAIL B — SINGLE CRESCENT FOOT CROSS-SECTION",
			fmt.Sprintf("Scale 1 : %.2f  |  Mid-Wall Thickness = %.1f mm", 1.0/detScale, m.LegMidWidthBotMM))

		// 3. Lower-right: DETAIL C — High-Magnification (6:1) Zoom of the 0.2 Leg Tip Fillet!
		zoomCX, zoomCY, zoomR := 312.0, 102.0, 34.0
		zoomScale := 5.5 // 5.5 : 1 magnification!
		if m.LegTipFilletBotMM > 4.0 {
			zoomScale = 22.0 / m.LegTipFilletBotMM
		}

		// Find actual tip vertex in `aligned` with max U
		maxIdx := 0
		for i, q := range aligned {
			if q[0] > aligned[maxIdx][0] {
				maxIdx = i
			}
		}
		refU, refV := aligned[maxIdx][0], aligned[maxIdx][1]

		p.Save()
		p.SetFillRGB(0.985, 0.988, 0.994)
		p.SetStrokeRGB(0.78, 0.16, 0.16)
		p.SetLineWidthMM(0.32)
		p.Circle(zoomCX, zoomCY, zoomR, true, true)

		// Clip magnified crescent tip inside rectangular viewport inside circle
		p.ClipRect(zoomCX-zoomR+3, zoomCY-zoomR+3, 2*zoomR-6, 2*zoomR-6)
		zoomPoly := make([][2]float64, len(aligned))
		for i, q := range aligned {
			zoomPoly[i] = [2]float64{zoomCX + 16.0 + (q[0]-refU)*zoomScale, zoomCY + (q[1]-refV)*zoomScale}
		}
		p.SetSolid()
		p.SetFillRGB(0.90, 0.83, 0.72)
		p.SetStrokeRGB(0.08, 0.11, 0.16)
		p.SetLineWidthMM(0.42)
		p.Polygon(zoomPoly, true, true, true)

		// Highlight the circular fillet radius center if active
		rFilletScaled := m.LegTipFilletBotMM * zoomScale
		if rFilletScaled > 1.2 {
			p.SetDashMM(1.0, 0.8)
			p.SetStrokeRGB(0.78, 0.16, 0.16)
			p.SetLineWidthMM(0.22)
			p.Circle(zoomCX+16.0-rFilletScaled, zoomCY, rFilletScaled, true, false)
		}
		p.Restore()

		p.LeaderCallout(zoomCX+16.0, zoomCY, zoomCX+24.0, zoomCY+18.0, zoomCX+38.0,
			fmt.Sprintf("Leg Tip Fillet R = %.2f mm", m.LegTipFilletBotMM),
			fmt.Sprintf("(Grasshopper Parameter = %.2f)", m.LegTipFilletBotMM/(m.OverallHeightMM/45.0)))
		drawViewLabel(p, zoomCX, zoomCY-zoomR-5.5, "DETAIL C — MAGNIFIED LEG VERTICAL EDGE FILLET",
			fmt.Sprintf("Magnification %.1f : 1  |  20-Segment Tangential Rolling-Ball Arc", zoomScale))
	}
}

// ==============================================================================
// SHEET 4: BILL OF MATERIALS (BOM), BLANK YIELD & COST-ESTIMATION SCHEDULE
// ==============================================================================

func drawTableSection(
	p *pdfPageCanvas,
	x, topY, totalW float64,
	sectionTitle string,
	colWidths []float64,
	headers []string,
	colAligns []string,
	rows [][]string,
) float64 {
	titleH := 7.2
	hdrH := 6.6
	rowH := 6.2

	p.SetFillRGB(0.12, 0.16, 0.23)
	p.Rect(x, topY-titleH, totalW, titleH, false, true)
	p.SetFillRGB(1, 1, 1)
	p.Text(x+3.0, topY-titleH+2.1, 8.0, "F2", sectionTitle)

	curY := topY - titleH - hdrH
	p.SetFillRGB(0.88, 0.90, 0.94)
	p.SetStrokeRGB(0.22, 0.26, 0.32)
	p.SetLineWidthMM(0.20)
	p.Rect(x, curY, totalW, hdrH, true, true)

	cx := x
	for cIdx, cw := range colWidths {
		if cIdx > 0 {
			p.Line(cx, curY, cx, curY+hdrH)
		}
		align := colAligns[cIdx]
		tx := cx + 2.2
		if align == "right" {
			tx = cx + cw - 2.2
		} else if align == "center" {
			tx = cx + 0.5*cw
		}
		p.SetFillRGB(0.10, 0.14, 0.20)
		p.TextAligned(tx, curY+1.9, 6.8, "F2", headers[cIdx], align)
		cx += cw
	}

	for rIdx, row := range rows {
		curY -= rowH
		isBoldRow := len(row) > 0 && strings.HasPrefix(row[0], "TOTAL")
		if isBoldRow {
			p.SetFillRGB(0.92, 0.94, 0.97)
		} else if rIdx%2 == 1 {
			p.SetFillRGB(0.975, 0.98, 0.988)
		} else {
			p.SetFillRGB(1, 1, 1)
		}
		p.SetStrokeRGB(0.72, 0.76, 0.82)
		p.SetLineWidthMM(0.16)
		p.Rect(x, curY, totalW, rowH, true, true)

		cx = x
		fontKey := "F1"
		if isBoldRow {
			fontKey = "F2"
		}
		for cIdx, cw := range colWidths {
			if cIdx > 0 {
				p.Line(cx, curY, cx, curY+rowH)
			}
			if cIdx < len(row) {
				align := colAligns[cIdx]
				tx := cx + 2.2
				if align == "right" {
					tx = cx + cw - 2.2
				} else if align == "center" {
					tx = cx + 0.5*cw
				}
				p.SetFillRGB(0.10, 0.14, 0.20)
				p.TextAligned(tx, curY+1.8, 6.8, fontKey, row[cIdx], align)
			}
			cx += cw
		}
	}

	p.SetStrokeRGB(0.12, 0.16, 0.23)
	p.SetLineWidthMM(0.30)
	totalH := titleH + hdrH + float64(len(rows))*rowH
	p.Rect(x, topY-totalH, totalW, totalH, true, false)
	return curY
}

func buildSheet4BOMAndCostDrivers(doc *pdfDocument, m *stoolDossierMetrics) {
	p := doc.NewPage()
	drawSheetFrameAndTitleBlock(p, m, 4, 4, "MONOLITHIC BLOCK SIZING, MASS & CNC CARVING SCHEDULE", "N/A (DATA)")

	// 1. Monolithic Geometry & Raw Block Schedule — One-Piece CAD/CNC Carving (no assembly assumed!)
	bomCols := []float64{18, 86, 16, 54, 56, 36, 38, 40, 40}
	bomHdrs := []string{"ZONE", "MONOLITHIC FEATURE / REGION", "QTY", "FINISHED ENVELOPE (mm)", "RAW BLOCK ENVELOPE (+6 mm)", "NET VOL (dm³)", "BLOCK VOL (dm³)", "SURFACE (m²)", "CONTOUR (m)"}
	bomAlign := []string{"center", "left", "center", "right", "right", "right", "right", "right", "right"}
	bomRows := [][]string{
		{
			"Z1",
			fmt.Sprintf("Upper Scalloped Seat Region (Z = %.1f .. %.1f mm)", m.SeatBottomYMM, m.OverallHeightMM),
			"1",
			fmt.Sprintf("%.1f × %.1f × %.1f", m.SeatBlankFinLMM, m.SeatBlankFinWMM, m.SeatBlankFinTMM),
			"Part of Monolith",
			fmt.Sprintf("%.3f", m.SeatNetVolCM3*1e-3),
			"—",
			fmt.Sprintf("%.3f", m.SeatAreaM2),
			fmt.Sprintf("%.2f", m.PerimSeatTopM+m.PerimSeatBotM),
		},
		{
			"Z2",
			fmt.Sprintf("Lower %d-Crescent-Leg Region (Z = 0.0 .. %.1f mm)", m.RadialReps, m.LegHeightMM),
			fmt.Sprintf("%d lobes", m.RadialReps),
			fmt.Sprintf("Chord %.1f × %.1f × %.1f", m.LegChordBotMM, m.LegDepthBotMM, m.LegHeightMM),
			"Part of Monolith",
			fmt.Sprintf("%.3f", m.LegsTotalNetVolCM3*1e-3),
			"—",
			fmt.Sprintf("%.3f", m.LegsTotalAreaM2),
			fmt.Sprintf("%.2f", float64(m.RadialReps)*2.0*m.PerimSingleFootM),
		},
		{
			"TOTAL",
			fmt.Sprintf("ONE-PIECE MONOLITHIC STOOL (Chip Removal = %.1f dm³)", m.ChipRemovalVolDM3),
			"1 pc",
			fmt.Sprintf("%.0f × %.0f × %.1f", m.OverallWidthXMM, m.OverallDepthZMM, m.OverallHeightMM),
			fmt.Sprintf("%.0f × %.0f × %.0f", m.MonolithBlockLMM, m.MonolithBlockWMM, m.MonolithBlockHMM),
			fmt.Sprintf("%.3f dm³", m.TotalNetVolCM3*1e-3),
			fmt.Sprintf("%.2f dm³", m.MonolithBlockVolDM3),
			fmt.Sprintf("%.3f m²", m.TotalAreaM2),
			fmt.Sprintf("%.2f m", m.TotalCNCToolpathM),
		},
	}
	nextY := drawTableSection(p, 18, 268, 384, "1. MONOLITHIC GEOMETRY & SOLID BLOCK SCHEDULE — ONE-PIECE CAD/CNC CARVING (NO ASSEMBLY)", bomCols, bomHdrs, bomAlign, bomRows)

	// 2. Left lower table: Wood Species Mass & Raw Monolithic Block Comparison (x=18, w=196)
	netM3 := m.TotalNetVolCM3 * 1e-6
	blockM3 := m.MonolithBlockVolDM3 * 1e-3

	spCols := []float64{68, 30, 32, 34, 32}
	spHdrs := []string{"SOLID TIMBER SPECIES (KD 8-10%)", "DENSITY", "FINISHED MASS", "RAW BLOCK MASS", "MATTE WAX/OIL"}
	spAlign := []string{"left", "right", "right", "right", "right"}
	waxGrams := m.TotalAreaM2 * 85.0
	spRows := [][]string{
		{
			"French / European White Oak",
			"720 kg/m³",
			fmt.Sprintf("%.2f kg", netM3*densityEuropeanOakKgM3),
			fmt.Sprintf("%.1f kg (%.0f dm³)", blockM3*densityEuropeanOakKgM3, m.MonolithBlockVolDM3),
			fmt.Sprintf("%.0f g (matte)", waxGrams),
		},
		{
			"Black Walnut (Juglans nigra / regia)",
			"640 kg/m³",
			fmt.Sprintf("%.2f kg", netM3*densityAmericanWalnutKgM3),
			fmt.Sprintf("%.1f kg (%.0f dm³)", blockM3*densityAmericanWalnutKgM3, m.MonolithBlockVolDM3),
			fmt.Sprintf("%.0f g (matte)", waxGrams),
		},
		{
			"European White Ash (Fraxinus)",
			"690 kg/m³",
			fmt.Sprintf("%.2f kg", netM3*690.0),
			fmt.Sprintf("%.1f kg (%.0f dm³)", blockM3*690.0, m.MonolithBlockVolDM3),
			fmt.Sprintf("%.0f g (matte)", waxGrams),
		},
	}
	yAfterSp := drawTableSection(p, 18, nextY-8, 196, "2. SOLID HARDWOOD MASS, RAW BLOCK & MATTE FINISH ESTIMATE", spCols, spHdrs, spAlign, spRows)

	// 3. Packaging & Logistics Cube beneath Species Table (x=18, w=196)
	pkgCols := []float64{76, 44, 38, 38}
	pkgHdrs := []string{"PACKAGING CONFIGURATION", "CARTON EXT (mm)", "VOLUME (m³)", "UNITS / EUR PALLET"}
	pkgAlign := []string{"left", "right", "right", "right"}
	cwMM := m.OverallWidthXMM + 30
	cdMM := m.OverallDepthZMM + 30
	ch1MM := m.OverallHeightMM + 30
	ch4MM := m.OverallHeightMM + 3*m.StackPitchHeightMM + 30
	pkgRows := [][]string{
		{
			"Single Stool Retail Carton (1×)",
			fmt.Sprintf("%.0f × %.0f × %.0f", cwMM, cdMM, ch1MM),
			fmt.Sprintf("%.3f m³", m.SingleCartonM3),
			"8 pcs (2×2×2)",
		},
		{
			"Helical Nested Contract Carton (4× Stacked)",
			fmt.Sprintf("%.0f × %.0f × %.0f", cwMM, cdMM, ch4MM),
			fmt.Sprintf("%.3f m³", m.Stack4CartonM3),
			"16 pcs (4× stacks)",
		},
	}
	drawTableSection(p, 18, yAfterSp-8, 196, "3. LOGISTICS, HELICAL NESTING & PACKAGING CUBE", pkgCols, pkgHdrs, pkgAlign, pkgRows)

	// 4. Right middle table: One-Piece 5-Axis CNC / Robotic Carving Cost-Driver Schedule (x=222, w=180)
	opCols := []float64{14, 48, 86, 32}
	opHdrs := []string{"OP", "WORKCENTER", "ONE-PIECE CAD/CAM OPERATION", "COST DRIVER"}
	opAlign := []string{"center", "left", "left", "right"}
	opRows := [][]string{
		{"10", "Timber Block Prep", fmt.Sprintf("Prepare %.0f×%.0f×%.0f mm solid / laminated block", m.MonolithBlockLMM, m.MonolithBlockWMM, m.MonolithBlockHMM), fmt.Sprintf("%.1f dm³", m.MonolithBlockVolDM3)},
		{"20", "5-Axis / Robot Rough", "Hog out under-seat & inter-leg cavities (one-piece setup)", fmt.Sprintf("-%.1f dm³", m.ChipRemovalVolDM3)},
		{"30", "5-Axis Seat Finish", fmt.Sprintf("Ball-nose mill %.1f° scalloped skirt + R%.1f top fillet", m.ProjAngleDeg, m.SeatTopFilletMM), fmt.Sprintf("%.2f m²", m.SeatAreaM2)},
		{"40", "5-Axis Leg Finish", fmt.Sprintf("Mill 3× crescent legs (%.1f° taper + R%.2f tip fillet)", m.ProjAngleDeg, m.LegTipFilletBotMM), fmt.Sprintf("%.2f m²", m.LegsTotalAreaM2)},
		{"50", "Datum & Fine Sand", "Face Z=0.0 coplanar feet; orbital/hand sand P120-P240", fmt.Sprintf("%.2f m²", m.TotalAreaM2)},
		{"60", "Matte Wax Finish", "Apply 2 thin coats ultra-matte raw-timber protector (<10 GU)", fmt.Sprintf("%.0f g", waxGrams)},
	}
	drawTableSection(p, 222, nextY-8, 180, "4. ONE-PIECE CAD/CAM CARVING COST-ESTIMATION WORKSHEET", opCols, opHdrs, opAlign, opRows)
}

func scaleGeomUniform(geom *threejs.BufferGeometry, s float64) {
	if geom == nil {
		return
	}
	for _, v := range geom.Vertices {
		v.X *= s
		v.Y *= s
		v.Z *= s
	}
}

func scaleLoop2D(pts []pt2D, s float64) {
	for i := range pts {
		pts[i].X *= s
		pts[i].Z *= s
	}
}

func normalizeStoolResultToHeight(res *StoolGeometriesResult, targetHeightMM float64) {
	if res == nil || res.StoolTopHeight <= 1e-3 {
		return
	}
	s := targetHeightMM / res.StoolTopHeight
	if math.Abs(s-1.0) < 1e-6 {
		return
	}
	scaleGeomUniform(res.SeatGeom, s)
	scaleGeomUniform(res.LegsGeom, s)
	scaleGeomUniform(res.CombinedGeom, s)
	res.GlobalR *= s
	res.StoolTopHeight = targetHeightMM
	res.SeatBottomHeight *= s
	res.GrowthYStep *= s

	if gh := res.GrasshopperContours; gh != nil {
		scaleLoop2D(gh.SeatTopRing, s)
		scaleLoop2D(gh.SeatBottomRing, s)
		gh.StoolTopHeight = targetHeightMM
		gh.SeatBottomY *= s
		gh.GrowthYStep *= s
		gh.SeatFilletR *= s
		for _, lc := range []*grasshopperLegContour{gh.LegTopContour, gh.LegBotContour} {
			if lc != nil {
				scaleLoop2D(lc.Loop, s)
				lc.Y *= s
				lc.MidThicknessMM *= s
				lc.AppliedOffsetMM *= s
				lc.AppliedFilletMM *= s
				lc.LeftFilletCtr.X *= s
				lc.LeftFilletCtr.Z *= s
				lc.RightFilletCtr.X *= s
				lc.RightFilletCtr.Z *= s
			}
		}
	}
}

// GenerateStoolManufacturingPDF generates a 4-sheet A3 Landscape professional technical & manufacturing
// cost-estimation PDF dossier in pure Go (normalized to 450.0 mm height, monolithic one-piece CAD carving).
func GenerateStoolManufacturingPDF(stager *models.Stager, plant *models.PlantAbstract) []byte {
	if plant == nil || plant.PlantType != models.Stool || plant.StoolAbstract == nil {
		return nil
	}

	stage := threejs.NewStage("export_pdf_dossier")
	res := ComputeStoolGeometries(stage, stager, plant)
	if res == nil || res.SeatGeom == nil || res.LegsGeom == nil || len(res.SeatGeom.Faces) == 0 {
		return nil
	}

	// Normalize stool geometry strictly to the 450.0 mm nominal design height.
	normalizeStoolResultToHeight(res, 450.0)

	metrics := computeDossierMetrics(plant, res)
	doc := &pdfDocument{
		title: fmt.Sprintf("Stool Manufacturing Dossier — %s (%s)", plant.Name, dossierAuthors),
	}

	buildSheet1Orthographic(doc, metrics, res)
	buildSheet2SectionAndStacking(doc, metrics, res)
	buildSheet3CNCContoursAndFilletDetail(doc, metrics)
	buildSheet4BOMAndCostDrivers(doc, metrics)

	return doc.Bytes()
}
