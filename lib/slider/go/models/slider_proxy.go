package models

// Define a Number constraint for our generic types
type Number interface {
	~int | ~float64
}

// Generic slider creation function
func NewSlider[T Number](
	target Target,
	name string,
	min T,
	max T,
	step T,
	valueRef *T,
) *Slider {
	slider := new(Slider).Stage(target.GetSliderStage())
	slider.Name = name

	// Set appropriate values based on type
	switch any(min).(type) {
	case float64:
		slider.IsFloat64 = true
		slider.MinFloat64 = any(min).(float64)
		slider.MaxFloat64 = any(max).(float64)
		slider.StepFloat64 = any(step).(float64)
		slider.ValueFloat64 = any(*valueRef).(float64)
	case int:
		slider.IsInt = true
		slider.MinInt = any(min).(int)
		slider.MaxInt = any(max).(int)
		slider.StepInt = any(step).(int)
		slider.ValueInt = any(*valueRef).(int)
	}

	proxy := NewSliderProxy(
		slider,
		valueRef,
		target,
	)

	slider.Proxy = proxy

	return slider
}

// NewSliderProxy creates a new proxy for a slider
func NewSliderProxy[T Number](
	slider *Slider,
	value *T,
	target Target,
) *SliderProxy[T] {
	proxy := new(SliderProxy[T])
	proxy.slider = slider
	proxy.Value = value
	proxy.target = target
	return proxy
}

// SliderProxy is a generic proxy for both int and float64
type SliderProxy[T Number] struct {
	slider *Slider
	Value  *T
	target Target
}

// Updated handles updating values when the slider changes
func (proxy *SliderProxy[T]) Updated() {
	// Update the value based on its type
	switch value := any(proxy.Value).(type) {
	case *float64:
		*value = proxy.slider.ValueFloat64
	case *int:
		*value = proxy.slider.ValueInt
	}

	proxy.target.OnAfterUpdateSliderElement()
}
