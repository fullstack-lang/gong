package models

type DominantBaselineType string

const (
	// auto: (Default)
	// If this property occurs on a <text> element, the computed value depends on the writing-mode.
	// - If horizontal, the dominant-baseline component is 'alphabetic'.
	// - If vertical, the dominant-baseline component is 'central'.
	//
	// If on a <tspan> or <textPath>, it remains the same as the parent.
	// If there is no parent text content element, it's computed as for <text>.
	DominantBaselineAuto DominantBaselineType = "auto"

	// alphabetic:
	// The baseline-identifier for the dominant-baseline is set to be 'alphabetic'.
	// The derived baseline-table is constructed using the 'alphabetic' baseline-table in the font.
	// The baseline-table font-size is changed to the value of the font-size attribute.
	DominantBaselineAlphabetic DominantBaselineType = "alphabetic"

	// ideographic:
	// The baseline-identifier for the dominant-baseline is set to be 'ideographic'.
	// The derived baseline-table is constructed using the 'ideographic' baseline-table in the font.
	// The baseline-table font-size is changed to the value of the font-size attribute.
	DominantBaselineIdeographic DominantBaselineType = "ideographic"

	// hanging:
	// The baseline-identifier for the dominant-baseline is set to be 'hanging'.
	// The derived baseline-table is constructed using the 'hanging' baseline-table in the font.
	// The baseline-table font-size is changed to the value of the font-size attribute.
	DominantBaselineHanging DominantBaselineType = "hanging"

	// mathematical:
	// The baseline-identifier for the dominant-baseline is set to be 'mathematical'.
	// The derived baseline-table is constructed using the 'mathematical' baseline-table in the font.
	// The baseline-table font-size is changed to the value of the font-size attribute.
	DominantBaselineMathematical DominantBaselineType = "mathematical"

	// central:
	// The baseline-identifier for the dominant-baseline is set to be 'central'.
	// The derived baseline-table is constructed from a font baseline-table chosen using the
	// following priority order: 'ideographic', 'alphabetic', 'hanging', 'mathematical'.
	// The baseline-table font-size is changed to the value of the font-size attribute.
	DominantBaselineCentral DominantBaselineType = "central"

	// middle:
	// The baseline-identifier for the dominant-baseline is set to be 'middle'.
	// The derived baseline-table is constructed from a font baseline-table chosen using the
	// following priority order: 'alphabetic', 'ideographic', 'hanging', 'mathematical'.
	// The baseline-table font-size is changed to the value of the font-size attribute.
	DominantBaselineMiddle DominantBaselineType = "middle"

	// text-after-edge:
	// The baseline-identifier for the dominant-baseline is set to be 'text-after-edge'.
	// The choice of which font baseline-table to use is browser dependent.
	// The baseline-table font-size is changed to the value of the font-size attribute.
	DominantBaselineTextAfterEdge DominantBaselineType = "text-after-edge"

	// text-before-edge:
	// The baseline-identifier for the dominant-baseline is set to be 'text-before-edge'.
	// The choice of which font baseline-table to use is browser dependent.
	// The baseline-table font-size is changed to the value of the font-size attribute.
	DominantBaselineTextBeforeEdge DominantBaselineType = "text-before-edge"

	// text-top:
	// This value uses the top of the em box as the baseline.
	DominantBaselineTextTop DominantBaselineType = "text-top"

	// -- Deprecated Values --

	// use-script: [Deprecated]
	// The dominant-baseline and baseline-table are set by determining the predominant script.
	// The writing-mode selects the appropriate baseline-tables, and the dominant baseline
	// selects the table that corresponds to that baseline.
	DominantBaselineUseScript DominantBaselineType = "use-script"

	// no-change: [Deprecated]
	// The dominant-baseline, the baseline-table, and the baseline-table font-size
	// remain the same as that of the parent text content element.
	DominantBaselineNoChange DominantBaselineType = "no-change"

	// reset-size: [Deprecated]
	// The dominant-baseline and the baseline-table remain the same, but the
	// baseline-table font-size is changed to the value of the font-size attribute on this element.
	DominantBaselineResetSize DominantBaselineType = "reset-size"
)
