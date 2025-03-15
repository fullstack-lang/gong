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
	StackName: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopySliderToSliderAPI(slider: Slider, sliderAPI: SliderAPI) {

	sliderAPI.CreatedAt = slider.CreatedAt
	sliderAPI.DeletedAt = slider.DeletedAt
	sliderAPI.ID = slider.ID

	// insertion point for basic fields copy operations
	sliderAPI.Name = slider.Name
	sliderAPI.StackName = slider.StackName

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
	slider.StackName = sliderAPI.StackName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
