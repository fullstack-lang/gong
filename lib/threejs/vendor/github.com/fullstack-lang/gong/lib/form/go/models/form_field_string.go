package models

type FormFieldString struct {
	Name  string
	Value string

	// IsTextArea is for telling angular whether to put a, <input ..> or a <textarea ..> which
	// allows for multi text editing
	IsTextArea bool
}
