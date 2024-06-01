package models

type LinkAnchoredText struct {
	Name    string
	Content string

	// AutomaticLayout is true will have the front compute the
	// X_Offset / Y_Offset of the anchored text
	AutomaticLayout bool
	LinkAnchorType  LinkAnchorType

	// values if AutomaticLayout is false
	X_Offset float64
	Y_Offset float64

	/*
		Numeric Font Weights

		100: Thin
		200: Extra Light (Ultra Light)
		300: Light
		400: Normal (Regular)
		500: Medium
		600: Semi Bold (Demi Bold)
		700: Bold
		800: Extra Bold (Ultra Bold)
		900: Black (Heavy)
		950: Extra Black (Ultra Black)

		Keyword Font Weights

		normal: Equivalent to 400
		bold: Equivalent to 700
		bolder: Increases the font weight compared to the parent element
		lighter: Decreases the font weight compared to the parent element
		inherit: Inherits the font weight from the parent element

			<text x="10" y="30" font-size="20" font-weight="100">Thin (100)</text>
			<text x="10" y="60" font-size="20" font-weight="200">Extra Light (200)</text>
			<text x="10" y="90" font-size="20" font-weight="300">Light (300)</text>
			<text x="10" y="120" font-size="20" font-weight="400">Normal (400)</text>
			<text x="10" y="150" font-size="20" font-weight="500">Medium (500)</text>
			<text x="10" y="180" font-size="20" font-weight="600">Semi Bold (600)</text>
			<text x="150" y="30" font-size="20" font-weight="700">Bold (700)</text>
			<text x="150" y="60" font-size="20" font-weight="800">Extra Bold (800)</text>
			<text x="150" y="90" font-size="20" font-weight="900">Black (900)</text>
	*/
	FontWeight string

	/*
		Units for font-size

		Pixels (px): The most common unit for specifying font size.
		Points (pt): 1 point is 1/72 inch.
		Em (em): Relative to the font size of the parent element. 1em means the same size as the parent.
		Percentage (%): Relative to the parent element’s font size. 100% means the same size as the parent.
		Centimeters (cm) / Millimeters (mm): Absolute units.
		Inches (in): Absolute unit. 1 inch is 96 pixels.
		Picas (pc): 1 pica is 1/6 inch.
		Ex (ex): Relative to the x-height of the font.
		Ch (ch): Relative to the width of the zero (0) character.

			<text x="10" y="30" font-size="10px">Font size 10px</text>
			<text x="10" y="60" font-size="14pt">Font size 14pt</text>
			<text x="10" y="90" font-size="2em">Font size 2em</text>
			<text x="10" y="120" font-size="150%">Font size 150%</text>
			<text x="10" y="150" font-size="10mm">Font size 10mm</text>
			<text x="10" y="180" font-size="x-large">Font size x-large</text>
	*/
	FontSize string

	/*
		Absolute Units

		    Pixels (px): Most commonly used because of their straightforward nature. For example, letter-spacing: 2px; adds two pixels between each character.
		    Points (pt): Often used in print, with 1 point equal to 1/72 of an inch. For example, letter-spacing: 1pt;.
		    Picas (pc): Another unit used in typography, where 1 pica equals 12 points. For example, letter-spacing: 0.5pc;.

		Relative Units

		    Em (em): Relative to the current font-size of the element. For instance, letter-spacing: 0.1em; means the letter spacing will be 0.1 times the size of the current font.
		    Percentage (%): Relative to the current font size. For example, letter-spacing: 5%; adjusts spacing based on a percentage of the current font size.
		    Viewport Width (vw) and Viewport Height (vh): These are relative to 1% of the width or height of the viewport. For example, letter-spacing: 0.1vw; would adjust the spacing based on the width of the viewport.

		Keyword

		    Normal: This is the default value, meaning the default spacing as determined by the font itself.
	*/
	LetterSpacing string

	Presentation
	Animates []*Animate

	Impl LinkAnchoredTextImplInterface
}

func (linkAnchoredText *LinkAnchoredText) OnAfterUpdate(stage *StageStruct, _, frontLinkAnchoredText *LinkAnchoredText) {

	if linkAnchoredText.Impl != nil {
		linkAnchoredText.Impl.AnchoredTextUpdated(frontLinkAnchoredText)
	}
}
