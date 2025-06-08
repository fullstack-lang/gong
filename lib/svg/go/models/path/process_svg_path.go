package path

import (
	"regexp"
	"strconv"
	"strings"
)

// PathBounds represents the maximum X and Y coordinates of a path
type PathBounds struct {
	MaxX float64
	MaxY float64
}

// Transform represents a 2D transformation matrix
type Transform struct {
	ScaleX     float64
	ScaleY     float64
	TranslateX float64
	TranslateY float64
	// For simplicity, we'll handle scale and translate only
	// A full implementation would include rotation and skew
}

// ProcessSVGPathWithTransform computes approximate maximum X and Y coordinates from an SVG path definition and transform
func ProcessSVGPathWithTransform(pathData, transformAttr string) PathBounds {
	transform := parseTransform(transformAttr)
	bounds := processSVGPath(pathData)

	// Apply transform to the bounds
	transformedBounds := PathBounds{}

	// Transform all corner points and take the maximum
	corners := []struct{ x, y float64 }{
		{0, 0},
		{bounds.MaxX, 0},
		{0, bounds.MaxY},
		{bounds.MaxX, bounds.MaxY},
	}

	for _, corner := range corners {
		tx, ty := transform.apply(corner.x, corner.y)
		transformedBounds.updateMax(tx, ty)
	}

	return transformedBounds
}

// ProcessSVGPath computes approximate maximum X and Y coordinates from an SVG path definition (without transform)
func ProcessSVGPath(pathData string) PathBounds {
	return processSVGPath(pathData)
}

// processSVGPath is the internal implementation
func processSVGPath(pathData string) PathBounds {
	bounds := PathBounds{MaxX: 0, MaxY: 0}

	// Current position
	var currentX, currentY float64

	// Clean up the path data - normalize spaces and handle command letters
	pathData = strings.ReplaceAll(pathData, ",", " ")
	pathData = regexp.MustCompile(`([MmLlHhVvCcSsQqTtAaZz])`).ReplaceAllString(pathData, " $1 ")
	pathData = regexp.MustCompile(`\s+`).ReplaceAllString(pathData, " ")
	pathData = strings.TrimSpace(pathData)

	// Split into tokens
	tokens := strings.Split(pathData, " ")

	i := 0
	for i < len(tokens) {
		if tokens[i] == "" {
			i++
			continue
		}

		command := tokens[i]
		i++

		switch command {
		case "M": // Move to (absolute)
			if i+1 < len(tokens) {
				x, _ := strconv.ParseFloat(tokens[i], 64)
				y, _ := strconv.ParseFloat(tokens[i+1], 64)
				currentX, currentY = x, y
				bounds.updateMax(x, y)
				i += 2
			}

		case "m": // Move to (relative)
			if i+1 < len(tokens) {
				dx, _ := strconv.ParseFloat(tokens[i], 64)
				dy, _ := strconv.ParseFloat(tokens[i+1], 64)
				currentX += dx
				currentY += dy
				bounds.updateMax(currentX, currentY)
				i += 2
			}

		case "L": // Line to (absolute)
			if i+1 < len(tokens) {
				x, _ := strconv.ParseFloat(tokens[i], 64)
				y, _ := strconv.ParseFloat(tokens[i+1], 64)
				currentX, currentY = x, y
				bounds.updateMax(x, y)
				i += 2
			}

		case "l": // Line to (relative)
			if i+1 < len(tokens) {
				dx, _ := strconv.ParseFloat(tokens[i], 64)
				dy, _ := strconv.ParseFloat(tokens[i+1], 64)
				currentX += dx
				currentY += dy
				bounds.updateMax(currentX, currentY)
				i += 2
			}

		case "H": // Horizontal line to (absolute)
			if i < len(tokens) {
				x, _ := strconv.ParseFloat(tokens[i], 64)
				currentX = x
				bounds.updateMax(x, currentY)
				i++
			}

		case "h": // Horizontal line to (relative)
			if i < len(tokens) {
				dx, _ := strconv.ParseFloat(tokens[i], 64)
				currentX += dx
				bounds.updateMax(currentX, currentY)
				i++
			}

		case "V": // Vertical line to (absolute)
			if i < len(tokens) {
				y, _ := strconv.ParseFloat(tokens[i], 64)
				currentY = y
				bounds.updateMax(currentX, y)
				i++
			}

		case "v": // Vertical line to (relative)
			if i < len(tokens) {
				dy, _ := strconv.ParseFloat(tokens[i], 64)
				currentY += dy
				bounds.updateMax(currentX, currentY)
				i++
			}

		case "C": // Cubic Bézier curve (absolute)
			if i+5 < len(tokens) {
				// Control point 1
				x1, _ := strconv.ParseFloat(tokens[i], 64)
				y1, _ := strconv.ParseFloat(tokens[i+1], 64)
				// Control point 2
				x2, _ := strconv.ParseFloat(tokens[i+2], 64)
				y2, _ := strconv.ParseFloat(tokens[i+3], 64)
				// End point
				x, _ := strconv.ParseFloat(tokens[i+4], 64)
				y, _ := strconv.ParseFloat(tokens[i+5], 64)

				// Update bounds with all points (approximation)
				bounds.updateMax(x1, y1)
				bounds.updateMax(x2, y2)
				bounds.updateMax(x, y)

				currentX, currentY = x, y
				i += 6
			}

		case "c": // Cubic Bézier curve (relative)
			if i+5 < len(tokens) {
				dx1, _ := strconv.ParseFloat(tokens[i], 64)
				dy1, _ := strconv.ParseFloat(tokens[i+1], 64)
				dx2, _ := strconv.ParseFloat(tokens[i+2], 64)
				dy2, _ := strconv.ParseFloat(tokens[i+3], 64)
				dx, _ := strconv.ParseFloat(tokens[i+4], 64)
				dy, _ := strconv.ParseFloat(tokens[i+5], 64)

				// Update bounds with control points
				bounds.updateMax(currentX+dx1, currentY+dy1)
				bounds.updateMax(currentX+dx2, currentY+dy2)

				currentX += dx
				currentY += dy
				bounds.updateMax(currentX, currentY)
				i += 6
			}

		case "Q": // Quadratic Bézier curve (absolute)
			if i+3 < len(tokens) {
				// Control point
				x1, _ := strconv.ParseFloat(tokens[i], 64)
				y1, _ := strconv.ParseFloat(tokens[i+1], 64)
				// End point
				x, _ := strconv.ParseFloat(tokens[i+2], 64)
				y, _ := strconv.ParseFloat(tokens[i+3], 64)

				bounds.updateMax(x1, y1)
				bounds.updateMax(x, y)

				currentX, currentY = x, y
				i += 4
			}

		case "q": // Quadratic Bézier curve (relative)
			if i+3 < len(tokens) {
				dx1, _ := strconv.ParseFloat(tokens[i], 64)
				dy1, _ := strconv.ParseFloat(tokens[i+1], 64)
				dx, _ := strconv.ParseFloat(tokens[i+2], 64)
				dy, _ := strconv.ParseFloat(tokens[i+3], 64)

				bounds.updateMax(currentX+dx1, currentY+dy1)

				currentX += dx
				currentY += dy
				bounds.updateMax(currentX, currentY)
				i += 4
			}

		case "A", "a": // Arc - simplified handling
			if i+6 < len(tokens) {
				// Skip arc parameters and just use end point
				if command == "A" {
					x, _ := strconv.ParseFloat(tokens[i+5], 64)
					y, _ := strconv.ParseFloat(tokens[i+6], 64)
					currentX, currentY = x, y
				} else {
					dx, _ := strconv.ParseFloat(tokens[i+5], 64)
					dy, _ := strconv.ParseFloat(tokens[i+6], 64)
					currentX += dx
					currentY += dy
				}
				bounds.updateMax(currentX, currentY)
				i += 7
			}

		case "Z", "z": // Close path
			// No coordinates to process

		default:
			// Skip unknown commands
			i++
		}
	}

	return bounds
}

// parseTransform parses a transform attribute and returns a Transform struct
func parseTransform(transformAttr string) Transform {
	transform := Transform{ScaleX: 1, ScaleY: 1, TranslateX: 0, TranslateY: 0}

	if transformAttr == "" {
		return transform
	}

	// Handle scale(x) and scale(x,y)
	scaleRegex := regexp.MustCompile(`scale\s*\(\s*([^)]+)\s*\)`)
	if matches := scaleRegex.FindStringSubmatch(transformAttr); len(matches) > 1 {
		parts := strings.Split(strings.TrimSpace(matches[1]), " ")
		if len(parts) >= 1 {
			if scale, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64); err == nil {
				transform.ScaleX = scale
				transform.ScaleY = scale // Default to uniform scaling
			}
		}
		if len(parts) >= 2 {
			if scaleY, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64); err == nil {
				transform.ScaleY = scaleY
			}
		}
	}

	// Handle translate(x) and translate(x,y)
	translateRegex := regexp.MustCompile(`translate\s*\(\s*([^)]+)\s*\)`)
	if matches := translateRegex.FindStringSubmatch(transformAttr); len(matches) > 1 {
		parts := strings.Split(strings.TrimSpace(matches[1]), " ")
		if len(parts) >= 1 {
			if tx, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64); err == nil {
				transform.TranslateX = tx
			}
		}
		if len(parts) >= 2 {
			if ty, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64); err == nil {
				transform.TranslateY = ty
			}
		}
	}

	return transform
}

// apply applies the transform to a point
func (t *Transform) apply(x, y float64) (float64, float64) {
	// Apply scale then translate
	newX := x*t.ScaleX + t.TranslateX
	newY := y*t.ScaleY + t.TranslateY
	return newX, newY
}

// updateMax updates the maximum X and Y values
func (b *PathBounds) updateMax(x, y float64) {
	if x > b.MaxX {
		b.MaxX = x
	}
	if y > b.MaxY {
		b.MaxY = y
	}
}

// Example usage - remove for production
/*
func main() {
	// Test with a simple path
	pathData := "M 10 10 L 100 50 Q 120 80 150 100 Z"
	bounds := ProcessSVGPath(pathData)

	fmt.Printf("Original - Max X: %.2f, Max Y: %.2f\n", bounds.MaxX, bounds.MaxY)

	// Test with transform
	transform := "scale(2) translate(10 20)"
	boundsWithTransform := ProcessSVGPathWithTransform(pathData, transform)

	fmt.Printf("With transform - Max X: %.2f, Max Y: %.2f\n", boundsWithTransform.MaxX, boundsWithTransform.MaxY)

	// Test with a more complex path and transform
	complexPath := "M 0 0 C 50 0 100 50 100 100 L 200 200 h 50 v -30"
	complexTransform := "scale(1.5 2) translate(25 10)"
	bounds2 := ProcessSVGPathWithTransform(complexPath, complexTransform)

	fmt.Printf("Complex path with transform - Max X: %.2f, Max Y: %.2f\n", bounds2.MaxX, bounds2.MaxY)
}
*/
