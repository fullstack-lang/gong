// generated code - do not edit

import { SliderAPI } from './slider-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Slider {

	static GONGSTRUCT_NAME = "Slider"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsFloat64: boolean = false
	IsInt: boolean = false
	MinInt: number = 0
	MaxInt: number = 0
	StepInt: number = 0
	ValueInt: number = 0
	MinFloat64: number = 0
	MaxFloat64: number = 0
	StepFloat64: number = 0
	ValueFloat64: number = 0

	// insertion point for pointers and slices of pointers declarations
}

export function CopySliderToSliderAPI(slider: Slider, sliderAPI: SliderAPI) {

	sliderAPI.CreatedAt = slider.CreatedAt
	sliderAPI.DeletedAt = slider.DeletedAt
	sliderAPI.ID = slider.ID

	// insertion point for basic fields copy operations
	sliderAPI.Name = slider.Name
	sliderAPI.IsFloat64 = slider.IsFloat64
	sliderAPI.IsInt = slider.IsInt
	sliderAPI.MinInt = slider.MinInt
	sliderAPI.MaxInt = slider.MaxInt
	sliderAPI.StepInt = slider.StepInt
	sliderAPI.ValueInt = slider.ValueInt
	sliderAPI.MinFloat64 = slider.MinFloat64
	sliderAPI.MaxFloat64 = slider.MaxFloat64
	sliderAPI.StepFloat64 = slider.StepFloat64
	sliderAPI.ValueFloat64 = slider.ValueFloat64

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopySliderAPIToSlider update basic, pointers and slice of pointers fields of slider
// from respectively the basic fields and encoded fields of pointers and slices of pointers of sliderAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopySliderAPIToSlider(sliderAPI: SliderAPI, slider: Slider, frontRepo: FrontRepo) {

	slider.CreatedAt = sliderAPI.CreatedAt
	slider.DeletedAt = sliderAPI.DeletedAt
	slider.ID = sliderAPI.ID

	// insertion point for basic fields copy operations
	slider.Name = sliderAPI.Name
	slider.IsFloat64 = sliderAPI.IsFloat64
	slider.IsInt = sliderAPI.IsInt
	slider.MinInt = sliderAPI.MinInt
	slider.MaxInt = sliderAPI.MaxInt
	slider.StepInt = sliderAPI.StepInt
	slider.ValueInt = sliderAPI.ValueInt
	slider.MinFloat64 = sliderAPI.MinFloat64
	slider.MaxFloat64 = sliderAPI.MaxFloat64
	slider.StepFloat64 = sliderAPI.StepFloat64
	slider.ValueFloat64 = sliderAPI.ValueFloat64

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
