package models

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	"github.com/fullstack-lang/gong/lib/svg/go/models/path"
)

type svgStyleState struct {
	fill                     string
	stroke                   string
	strokeWidth              float64
	fillOpacity              float64
	strokeOpacity            float64
	hasExplicitStrokeOpacity bool
}

func defaultSVGStyleState() svgStyleState {
	return svgStyleState{
		fill:          "",
		stroke:        "",
		strokeWidth:   1.0,
		fillOpacity:   1.0,
		strokeOpacity: 1.0,
	}
}

func (st svgStyleState) withAttrs(attrs map[string]string) svgStyleState {
	next := st

	applyKV := func(k, v string) {
		k = strings.ToLower(strings.TrimSpace(k))
		v = strings.TrimSpace(v)
		if v == "" || strings.EqualFold(v, "inherit") {
			return
		}
		switch k {
		case "fill":
			next.fill = v
		case "stroke":
			next.stroke = v
		case "stroke-width":
			if f, ok := parseSVGFloat(v); ok && f >= 0 {
				next.strokeWidth = f
			}
		case "opacity":
			if f, ok := parseSVGFloat(v); ok {
				f = clamp01(f)
				next.fillOpacity *= f
				next.strokeOpacity *= f
				next.hasExplicitStrokeOpacity = true
			}
		case "fill-opacity":
			if f, ok := parseSVGFloat(v); ok {
				next.fillOpacity = clamp01(f)
			}
		case "stroke-opacity":
			if f, ok := parseSVGFloat(v); ok {
				next.strokeOpacity = clamp01(f)
				next.hasExplicitStrokeOpacity = true
			}
		}
	}

	for k, v := range attrs {
		if k == "style" {
			continue
		}
		applyKV(k, v)
	}

	if styleVal, ok := attrs["style"]; ok && styleVal != "" {
		decls := strings.Split(styleVal, ";")
		for _, decl := range decls {
			colon := strings.IndexByte(decl, ':')
			if colon <= 0 {
				continue
			}
			applyKV(decl[:colon], decl[colon+1:])
		}
	}

	return next
}

// buildSVGRectAnchoredPaths parses a raw <svg>...</svg> string (or a bare SVG path "d" definition)
// without using reflect, and converts its vector shapes into []*svg.RectAnchoredPath scaled and
// centered uniformly inside the target square (boxX, boxY, targetSize) relative to anchorType.
func buildSVGRectAnchoredPaths(
	name string,
	rawSVG string,
	anchorType svg.RectAnchorType,
	boxX float64,
	boxY float64,
	targetSize float64,
	defaultColor string,
) []*svg.RectAnchoredPath {
	trimmed := strings.TrimSpace(rawSVG)
	if trimmed == "" || targetSize <= 0 {
		return nil
	}
	if defaultColor == "" {
		defaultColor = "#555555"
	}

	// Fallback: bare SVG path definition ("M 10 10 ...") without XML tags
	if !strings.Contains(trimmed, "<") {
		minX, minY, vbW, vbH := inferBarePathViewBox(trimmed)
		scale, x0, y0 := computeIconPlacement(minX, minY, vbW, vbH, boxX, boxY, targetSize)
		return []*svg.RectAnchoredPath{
			{
				Name:                name + "_icon_0",
				Definition:          trimmed,
				X_Offset:            x0,
				Y_Offset:            y0,
				RectAnchorType:      anchorType,
				ScalePropotionnally: true,
				AppliedScaling:      scale,
				Presentation: svg.Presentation{
					Color:         defaultColor,
					FillOpacity:   1.0,
					Stroke:        "none",
					StrokeOpacity: 0,
					StrokeWidth:   0,
				},
			},
		}
	}

	cleaned := stripXMLCommentsAndDeclarations(trimmed)

	minX, minY, vbW, vbH := 0.0, 0.0, 24.0, 24.0
	hasRootSVG := false

	type parsedShape struct {
		definition string
		style      svgStyleState
	}
	var shapes []parsedShape

	styleStack := []svgStyleState{defaultSVGStyleState()}
	skipDepth := 0

	pos := 0
	for pos < len(cleaned) {
		openIdx := strings.IndexByte(cleaned[pos:], '<')
		if openIdx < 0 {
			break
		}
		tagStart := pos + openIdx
		tagEnd := findTagClose(cleaned, tagStart+1)
		if tagEnd < 0 {
			break
		}
		rawTag := strings.TrimSpace(cleaned[tagStart+1 : tagEnd])
		pos = tagEnd + 1

		if rawTag == "" || strings.HasPrefix(rawTag, "!") || strings.HasPrefix(rawTag, "?") {
			continue
		}

		// Closing tag: </...>
		if strings.HasPrefix(rawTag, "/") {
			closeName := normalizeSVGTagName(rawTag[1:])
			if skipDepth > 0 {
				if isSkippedSVGContainer(closeName) {
					skipDepth--
				}
				continue
			}
			if len(styleStack) > 1 {
				styleStack = styleStack[:len(styleStack)-1]
			}
			continue
		}

		selfClosing := false
		if strings.HasSuffix(rawTag, "/") {
			selfClosing = true
			rawTag = strings.TrimSpace(rawTag[:len(rawTag)-1])
		}

		tagName, attrPart := splitTagNameAndAttrs(rawTag)
		tagName = normalizeSVGTagName(tagName)

		if skipDepth > 0 {
			if !selfClosing && isSkippedSVGContainer(tagName) {
				skipDepth++
			}
			continue
		}

		if isSkippedSVGContainer(tagName) {
			if !selfClosing {
				skipDepth++
			}
			continue
		}

		attrs := parseSVGAttributes(attrPart)
		curStyle := styleStack[len(styleStack)-1].withAttrs(attrs)

		if tagName == "svg" && !hasRootSVG {
			hasRootSVG = true
			if x, y, w, h, ok := parseViewBox(attrs["viewbox"]); ok {
				minX, minY, vbW, vbH = x, y, w, h
			} else {
				w, okW := parseSVGLength(attrs["width"])
				h, okH := parseSVGLength(attrs["height"])
				if okW && okH && w > 0 && h > 0 {
					minX, minY, vbW, vbH = 0, 0, w, h
				}
			}
		}

		if d := svgElementToPathD(tagName, attrs); d != "" {
			shapes = append(shapes, parsedShape{
				definition: d,
				style:      curStyle,
			})
		}

		if !selfClosing {
			styleStack = append(styleStack, curStyle)
		}
	}

	if len(shapes) == 0 {
		return nil
	}

	scale, x0, y0 := computeIconPlacement(minX, minY, vbW, vbH, boxX, boxY, targetSize)

	var result []*svg.RectAnchoredPath
	for _, sh := range shapes {
		fillColor, fillOpacity, strokeColor, strokeWidth, strokeOpacity, visible := resolveShapePresentation(sh.style, defaultColor)
		if !visible {
			continue
		}
		idx := len(result)
		result = append(result, &svg.RectAnchoredPath{
			Name:                fmt.Sprintf("%s_icon_%d", name, idx),
			Definition:          sh.definition,
			X_Offset:            x0,
			Y_Offset:            y0,
			RectAnchorType:      anchorType,
			ScalePropotionnally: true,
			AppliedScaling:      scale,
			Presentation: svg.Presentation{
				Color:         fillColor,
				FillOpacity:   fillOpacity,
				Stroke:        strokeColor,
				StrokeOpacity: strokeOpacity,
				StrokeWidth:   strokeWidth,
			},
		})
	}

	return result
}

func computeIconPlacement(minX, minY, vbW, vbH, boxX, boxY, targetSize float64) (scale, xOffset, yOffset float64) {
	if vbW <= 0 {
		vbW = 24.0
	}
	if vbH <= 0 {
		vbH = 24.0
	}
	maxDim := math.Max(vbW, vbH)
	scale = targetSize / maxDim
	xOffset = boxX + (targetSize-scale*vbW)/2.0 - scale*minX
	yOffset = boxY + (targetSize-scale*vbH)/2.0 - scale*minY
	return scale, xOffset, yOffset
}

func resolveShapePresentation(st svgStyleState, defaultColor string) (fillColor string, fillOpacity float64, strokeColor string, strokeWidth float64, strokeOpacity float64, visible bool) {
	strokeAttr := strings.TrimSpace(st.stroke)
	hasStroke := strokeAttr != "" &&
		!strings.EqualFold(strokeAttr, "none") &&
		!strings.EqualFold(strokeAttr, "transparent") &&
		st.strokeWidth > 0 &&
		st.strokeOpacity > 0

	if hasStroke {
		if strings.EqualFold(strokeAttr, "currentColor") {
			strokeColor = defaultColor
		} else {
			strokeColor = strokeAttr
		}
		strokeWidth = st.strokeWidth
		if st.hasExplicitStrokeOpacity {
			strokeOpacity = st.strokeOpacity
		} else {
			strokeOpacity = 1.0
		}
	} else {
		strokeColor = "none"
		strokeWidth = 0
		strokeOpacity = 0
	}

	fillAttr := strings.TrimSpace(st.fill)
	if strings.EqualFold(fillAttr, "none") || strings.EqualFold(fillAttr, "transparent") || st.fillOpacity <= 0 {
		fillColor = "none"
		fillOpacity = 0
	} else if fillAttr == "" || strings.EqualFold(fillAttr, "currentColor") {
		fillColor = defaultColor
		fillOpacity = st.fillOpacity
	} else {
		fillColor = fillAttr
		fillOpacity = st.fillOpacity
	}

	visible = fillColor != "none" || strokeColor != "none"
	return
}

func svgElementToPathD(tagName string, attrs map[string]string) string {
	switch tagName {
	case "path":
		return strings.TrimSpace(attrs["d"])

	case "rect":
		w, okW := parseSVGFloat(attrs["width"])
		h, okH := parseSVGFloat(attrs["height"])
		if !okW || !okH || w <= 0 || h <= 0 {
			return ""
		}
		x, _ := parseSVGFloat(attrs["x"])
		y, _ := parseSVGFloat(attrs["y"])
		rx, hasRX := parseSVGFloat(attrs["rx"])
		ry, hasRY := parseSVGFloat(attrs["ry"])
		if hasRX && !hasRY {
			ry = rx
		} else if hasRY && !hasRX {
			rx = ry
		}
		if rx < 0 {
			rx = 0
		}
		if ry < 0 {
			ry = 0
		}
		if rx > w/2 {
			rx = w / 2
		}
		if ry > h/2 {
			ry = h / 2
		}
		if rx > 0 && ry > 0 {
			return fmt.Sprintf(
				"M %g %g H %g A %g %g 0 0 1 %g %g V %g A %g %g 0 0 1 %g %g H %g A %g %g 0 0 1 %g %g V %g A %g %g 0 0 1 %g %g Z",
				x+rx, y,
				x+w-rx, rx, ry, x+w, y+ry,
				y+h-ry, rx, ry, x+w-rx, y+h,
				x+rx, rx, ry, x, y+h-ry,
				y+ry, rx, ry, x+rx, y,
			)
		}
		return fmt.Sprintf("M %g %g H %g V %g H %g Z", x, y, x+w, y+h, x)

	case "circle":
		r, okR := parseSVGFloat(attrs["r"])
		if !okR || r <= 0 {
			return ""
		}
		cx, _ := parseSVGFloat(attrs["cx"])
		cy, _ := parseSVGFloat(attrs["cy"])
		return fmt.Sprintf(
			"M %g %g A %g %g 0 1 0 %g %g A %g %g 0 1 0 %g %g Z",
			cx-r, cy, r, r, cx+r, cy, r, r, cx-r, cy,
		)

	case "ellipse":
		rx, okRX := parseSVGFloat(attrs["rx"])
		ry, okRY := parseSVGFloat(attrs["ry"])
		if !okRX || !okRY || rx <= 0 || ry <= 0 {
			return ""
		}
		cx, _ := parseSVGFloat(attrs["cx"])
		cy, _ := parseSVGFloat(attrs["cy"])
		return fmt.Sprintf(
			"M %g %g A %g %g 0 1 0 %g %g A %g %g 0 1 0 %g %g Z",
			cx-rx, cy, rx, ry, cx+rx, cy, rx, ry, cx-rx, cy,
		)

	case "line":
		x1, _ := parseSVGFloat(attrs["x1"])
		y1, _ := parseSVGFloat(attrs["y1"])
		x2, _ := parseSVGFloat(attrs["x2"])
		y2, _ := parseSVGFloat(attrs["y2"])
		if x1 == x2 && y1 == y2 {
			return ""
		}
		return fmt.Sprintf("M %g %g L %g %g", x1, y1, x2, y2)

	case "polyline", "polygon":
		nums := parseSVGNumberList(attrs["points"])
		if len(nums) < 4 {
			return ""
		}
		var sb strings.Builder
		fmt.Fprintf(&sb, "M %g %g", nums[0], nums[1])
		for i := 2; i+1 < len(nums); i += 2 {
			fmt.Fprintf(&sb, " L %g %g", nums[i], nums[i+1])
		}
		if tagName == "polygon" {
			sb.WriteString(" Z")
		}
		return sb.String()
	}
	return ""
}

func isSkippedSVGContainer(tagName string) bool {
	switch tagName {
	case "defs", "clippath", "mask", "style", "script", "title", "desc", "metadata",
		"symbol", "lineargradient", "radialgradient", "pattern", "filter", "marker":
		return true
	default:
		return false
	}
}

func normalizeSVGTagName(raw string) string {
	name := strings.TrimSpace(raw)
	for i := 0; i < len(name); i++ {
		c := name[i]
		if c == ' ' || c == '\t' || c == '\n' || c == '\r' || c == '/' {
			name = name[:i]
			break
		}
	}
	if colon := strings.IndexByte(name, ':'); colon >= 0 {
		name = name[colon+1:]
	}
	return strings.ToLower(name)
}

func splitTagNameAndAttrs(rawTag string) (string, string) {
	for i := 0; i < len(rawTag); i++ {
		c := rawTag[i]
		if c == ' ' || c == '\t' || c == '\n' || c == '\r' {
			return rawTag[:i], rawTag[i+1:]
		}
	}
	return rawTag, ""
}

func findTagClose(s string, start int) int {
	var quote byte
	for i := start; i < len(s); i++ {
		c := s[i]
		if quote != 0 {
			if c == quote {
				quote = 0
			}
			continue
		}
		if c == '"' || c == '\'' {
			quote = c
			continue
		}
		if c == '>' {
			return i
		}
	}
	return -1
}

func stripXMLCommentsAndDeclarations(s string) string {
	var sb strings.Builder
	sb.Grow(len(s))
	i := 0
	for i < len(s) {
		if strings.HasPrefix(s[i:], "<!--") {
			end := strings.Index(s[i+4:], "-->")
			if end < 0 {
				break
			}
			i += 4 + end + 3
			continue
		}
		if strings.HasPrefix(s[i:], "<![CDATA[") {
			end := strings.Index(s[i+9:], "]]>")
			if end < 0 {
				break
			}
			i += 9 + end + 3
			continue
		}
		sb.WriteByte(s[i])
		i++
	}
	return sb.String()
}

func parseSVGAttributes(attrStr string) map[string]string {
	attrs := make(map[string]string)
	i := 0
	n := len(attrStr)
	for i < n {
		for i < n && isSVGWhitespace(attrStr[i]) {
			i++
		}
		if i >= n || attrStr[i] == '/' {
			break
		}
		keyStart := i
		for i < n && !isSVGWhitespace(attrStr[i]) && attrStr[i] != '=' && attrStr[i] != '/' {
			i++
		}
		key := strings.ToLower(attrStr[keyStart:i])
		if colon := strings.IndexByte(key, ':'); colon >= 0 && !strings.HasPrefix(key, "xml") {
			key = key[colon+1:]
		}
		for i < n && isSVGWhitespace(attrStr[i]) {
			i++
		}
		if i >= n || attrStr[i] != '=' {
			continue
		}
		i++ // skip '='
		for i < n && isSVGWhitespace(attrStr[i]) {
			i++
		}
		if i >= n {
			break
		}
		var val string
		if attrStr[i] == '"' || attrStr[i] == '\'' {
			q := attrStr[i]
			i++
			valStart := i
			for i < n && attrStr[i] != q {
				i++
			}
			val = attrStr[valStart:i]
			if i < n && attrStr[i] == q {
				i++
			}
		} else {
			valStart := i
			for i < n && !isSVGWhitespace(attrStr[i]) && attrStr[i] != '/' {
				i++
			}
			val = attrStr[valStart:i]
		}
		if key != "" {
			attrs[key] = val
		}
	}
	return attrs
}

func parseViewBox(v string) (minX, minY, w, h float64, ok bool) {
	nums := parseSVGNumberList(v)
	if len(nums) < 4 {
		return 0, 0, 0, 0, false
	}
	if nums[2] <= 0 || nums[3] <= 0 {
		return 0, 0, 0, 0, false
	}
	return nums[0], nums[1], nums[2], nums[3], true
}

func parseSVGLength(v string) (float64, bool) {
	v = strings.TrimSpace(v)
	if v == "" || strings.HasSuffix(v, "%") {
		return 0, false
	}
	return parseSVGFloat(v)
}

func parseSVGFloat(v string) (float64, bool) {
	v = strings.TrimSpace(v)
	for _, suffix := range []string{"px", "pt", "em", "rem"} {
		if strings.HasSuffix(strings.ToLower(v), suffix) {
			v = strings.TrimSpace(v[:len(v)-len(suffix)])
			break
		}
	}
	if v == "" {
		return 0, false
	}
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0, false
	}
	return f, true
}

func parseSVGNumberList(s string) []float64 {
	var nums []float64
	i := 0
	n := len(s)
	for i < n {
		for i < n && (isSVGWhitespace(s[i]) || s[i] == ',') {
			i++
		}
		if i >= n {
			break
		}
		start := i
		if s[i] == '+' || s[i] == '-' {
			i++
		}
		hasDot := false
		for i < n {
			c := s[i]
			if c >= '0' && c <= '9' {
				i++
			} else if c == '.' && !hasDot {
				hasDot = true
				i++
			} else if (c == 'e' || c == 'E') && i > start {
				i++
				if i < n && (s[i] == '+' || s[i] == '-') {
					i++
				}
			} else {
				break
			}
		}
		if i > start {
			if f, err := strconv.ParseFloat(s[start:i], 64); err == nil {
				nums = append(nums, f)
			} else {
				i++
			}
		} else {
			i++
		}
	}
	return nums
}

func inferBarePathViewBox(d string) (minX, minY, vbW, vbH float64) {
	bounds := path.ProcessSVGPath(d)
	maxCoord := math.Max(bounds.MaxX, bounds.MaxY)
	if maxCoord <= 24.0 {
		// Check whether this is a Material Symbols bare path in [0, -960, 960, 960]
		if strings.Contains(d, "-960") || strings.Contains(d, "-880") || strings.Contains(d, "-800") {
			return 0, -960, 960, 960
		}
		return 0, 0, 24, 24
	}
	return 0, 0, maxCoord, maxCoord
}

func isSVGWhitespace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\n' || b == '\r'
}

func clamp01(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}
