package dao

import "regexp"

var uuid = regexp.MustCompile(`(?i:^[0-9A-F]{8}-[0-9A-F]{4}-[1-5][0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$)`)

// UUID wraps built-in string type and provide useful methods above it.
type UUID string

// IsValid returns true if a value is compatible with RFC 4122.
func (s UUID) IsValid() bool {
	return uuid.MatchString(string(s))
}

// InputType wraps built-in string type and provide useful methods above it.
type InputType string

const (
	InputTypeEmail         InputType = "email"
	InputTypeButton        InputType = "button"
	InputTypeCheckbox      InputType = "checkbox"
	InputTypeColor         InputType = "color"
	InputTypeDate          InputType = "date"
	InputTypeDatetimeLocal InputType = "datetime-local"
	InputTypeFile          InputType = "file"
	InputTypeHidden        InputType = "hidden"
	InputTypeImage         InputType = "image"
	InputTypeMonth         InputType = "month"
	InputTypeNumber        InputType = "number"
	InputTypePassword      InputType = "password"
	InputTypeRadio         InputType = "radio"
	InputTypeRange         InputType = "range"
	InputTypeReset         InputType = "reset"
	InputTypeSearch        InputType = "search"
	InputTypeSubmit        InputType = "submit"
	InputTypeTel           InputType = "tel"
	InputTypeText          InputType = "text"
	InputTypeTime          InputType = "time"
	InputTypeURL           InputType = "url"
	InputTypeWeek          InputType = "week"
)

var supported = []InputType{
	InputTypeEmail,
	InputTypeCheckbox,
	InputTypeColor,
	InputTypeDate,
	InputTypeDatetimeLocal,
	InputTypeHidden,
	InputTypeMonth,
	InputTypeNumber,
	InputTypePassword,
	InputTypeRange,
	InputTypeTel,
	InputTypeText,
	InputTypeTime,
	InputTypeURL,
	InputTypeWeek,
}

// IsSupported returns true if current input type in whitelist.
func (i InputType) IsSupported() bool {
	for _, v := range supported {
		if i == v {
			return true
		}
	}
	return false
}
