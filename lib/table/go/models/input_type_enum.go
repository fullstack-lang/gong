package models

// InputTypeEnum - enumeration of possible 'type' values for an HTML <input> element
// swagger:enum InputTypeEnum
type InputTypeEnum string

// values for InputTypeEnum
const (
	Text     InputTypeEnum = "text"
	Password InputTypeEnum = "password"
	Number   InputTypeEnum = "number"
	Email    InputTypeEnum = "email"
	Tel      InputTypeEnum = "tel"
	Date     InputTypeEnum = "date"
	Datetime InputTypeEnum = "datetime-local"
	Time     InputTypeEnum = "time"
	URL      InputTypeEnum = "url"
	Search   InputTypeEnum = "search"
	Range    InputTypeEnum = "range"
	Color    InputTypeEnum = "color"
	File     InputTypeEnum = "file"
	Hidden   InputTypeEnum = "hidden"
	Month    InputTypeEnum = "month"
	Week     InputTypeEnum = "week"
)
