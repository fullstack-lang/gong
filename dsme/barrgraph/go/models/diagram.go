package models

import (
	"time"
)

type Diagram struct {
	Name               string
	MovementShapes     []*MovementShape
	ArtefactTypeShapes []*ArtefactTypeShape
	ArtistShapes       []*ArtistShape
	InfluenceShapes    []*InfluenceShape

	IsEditable bool

	IsNodeExpanded                     bool
	IsMovementCategoryNodeExpanded     bool
	IsArtefactTypeCategoryNodeExpanded bool
	IsArtistCategoryNodeExpanded       bool
	IsInfluenceCategoryNodeExpanded    bool

	IsMovementCategoryShown     bool
	IsArtefactTypeCategoryShown bool
	IsArtistCategoryShown       bool
	IsInfluenceCategoryShown    bool

	StartDate time.Time
	EndDate   time.Time

	DiagramPresentation
}

type DiagramPresentation struct {
	XMargin float64
	YMargin float64

	Height                  float64
	NextVerticalDateXMargin float64

	RedColorCode            string
	BackgroundGreyColorCode string
	GrayColorCode           string

	BottomBoxYOffset         float64
	BottomBoxWidth           float64
	BottomBoxHeigth          float64
	BottomBoxFontSize        string
	BottomBoxFontWeigth      string
	BottomBoxFontFamily      string
	BottomBoxLetterSpacing   string
	BottomBoxLetterColorCode string

	MovementRectAnchorType       RectAnchorType
	MovementTextAnchorType       TextAnchorType
	MovementDominantBaselineType DominantBaselineType
	MovementFontSize             string
	MajorMovementFontSize        string
	MinorMovementFontSize        string
	MovementFontWeigth           string
	MovementFontFamily           string
	MovementLetterSpacing        string

	AbstractMovementFontSize       string
	AbstractMovementRectAnchorType RectAnchorType
	AbstractMovementTextAnchorType TextAnchorType
	AbstractDominantBaselineType   DominantBaselineType

	MovementDateRectAnchorType           RectAnchorType
	MovementDateTextAnchorType           TextAnchorType
	MovementDateTextDominantBaselineType DominantBaselineType
	MovementDateAndPlacesFontSize        string
	MovementDateAndPlacesFontWeigth      string
	MovementDateAndPlacesFontFamily      string
	MovementDateAndPlacesLetterSpacing   string

	MovementBelowArcY_Offset         float64
	MovementBelowArcY_OffsetPerPlace float64

	MovementPlacesRectAnchorType       RectAnchorType
	MovementPlacesTextAnchorType       TextAnchorType
	MovementPlacesDominantBaselineType DominantBaselineType

	ArtefactTypeFontSize      string
	ArtefactTypeFontWeigth    string
	ArtefactTypeFontFamily    string
	ArtefactTypeLetterSpacing string

	ArtefactTypeRectAnchorType   RectAnchorType
	ArtefactDominantBaselineType DominantBaselineType
	ArtefactTypeStrokeWidth      float64

	ArtistRectAnchorType       RectAnchorType
	ArtistTextAnchorType       TextAnchorType
	ArtistDominantBaselineType DominantBaselineType
	ArtistFontSize             string
	MajorArtistFontSize        string
	MinorArtistFontSize        string
	ArtistFontWeigth           string
	ArtistFontFamily           string
	ArtistLetterSpacing        string

	ArtistDateRectAnchorType         RectAnchorType
	ArtistDateTextAnchorType         TextAnchorType
	ArtistDateDominantBaselineType   DominantBaselineType
	ArtistDateAndPlacesFontSize      string
	ArtistDateAndPlacesFontWeigth    string
	ArtistDateAndPlacesFontFamily    string
	ArtistDateAndPlacesLetterSpacing string

	ArtistPlacesRectAnchorType       RectAnchorType
	ArtistPlacesTextAnchorType       TextAnchorType
	ArtistPlacesDominantBaselineType DominantBaselineType

	InfluenceArrowSize        float64
	InfluenceArrowStartOffset float64
	InfluenceArrowEndOffset   float64
	InfluenceCornerRadius     float64

	InfluenceDashedLinePattern string
}

// AlignDatesToFiveYears updates the diagram's StartDate and EndDate
// to align with 5-year intervals.
//
// It returns true if either StartDate or EndDate was changed,
// and false if they were already aligned.
func (d *Diagram) AlignDatesToFiveYears() bool {
	var changed bool = false

	// Store original values for comparison
	originalStart := d.StartDate
	originalEnd := d.EndDate

	// --- 1. Align StartDate (Round Down) ---
	if !d.StartDate.IsZero() {
		startYear := d.StartDate.Year()
		alignedStartYear := startYear - (startYear % 5)

		// Calculate the new potential date
		newStartDate := time.Date(alignedStartYear, time.January, 1, 0, 0, 0, 0, d.StartDate.Location())

		// Only assign and flag if it's different
		if !originalStart.Equal(newStartDate) {
			d.StartDate = newStartDate
			changed = true
		}
	}

	// --- 2. Align EndDate (Round Up) ---
	if !d.EndDate.IsZero() {
		endYear := d.EndDate.Year()
		floorYear := endYear - (endYear % 5)
		floorDate := time.Date(floorYear, time.January, 1, 0, 0, 0, 0, d.EndDate.Location())

		alignedEndYear := floorYear

		if d.EndDate.After(floorDate) {
			alignedEndYear += 5
		}

		// Calculate the new potential date
		newEndDate := time.Date(alignedEndYear, time.January, 1, 0, 0, 0, 0, d.EndDate.Location())

		// Only assign and flag if it's different
		if !originalEnd.Equal(newEndDate) {
			d.EndDate = newEndDate
			changed = true
		}
	}

	return changed
}

// GetFiveYearTicks generates a slice of time.Time, starting at StartDate,
// incrementing by 5 years, and always including EndDate as the final value.
//
// The 5-year steps are relative to the StartDate, not the calendar.
func (d *Diagram) GetFiveYearTicks() []time.Time {
	// --- 1. Handle edge cases ---
	// If dates are invalid, end is before start, or start is zero, return empty slice.
	if d.StartDate.IsZero() || d.EndDate.IsZero() || d.StartDate.After(d.EndDate) {
		return []time.Time{}
	}

	// If StartDate and EndDate are the same, return just that one date.
	if d.StartDate.Equal(d.EndDate) {
		return []time.Time{d.StartDate}
	}

	// --- 2. Build the slice ---
	ticks := []time.Time{}
	currentDate := d.StartDate

	// Loop, adding 5-year increments as long as we are *before* the end date.
	for currentDate.Before(d.EndDate) {
		ticks = append(ticks, currentDate)
		currentDate = currentDate.AddDate(5, 0, 0)
	}

	// --- 3. Add the final EndDate ---
	// This ensures the EndDate is always the last item,
	// satisfying the "including start and end" requirement.
	// This also correctly handles the case where EndDate
	// was an exact 5-year multiple (the loop stops just before it).
	ticks = append(ticks, d.EndDate)

	return ticks
}
