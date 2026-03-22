package errx

import "strings"

// MultiError aggregates multiple error occurrences into a single composite error.
type MultiError struct {
	errors []error
}

// Append adds an error to the collection. Discards nil errors automatically.
func (m *MultiError) Append(err error) {
	if err != nil {
		m.errors = append(m.errors, err)
	}
}

// Error implements the standard error interface by concatenating all error messages.
func (m *MultiError) Error() string {
	if len(m.errors) == 0 {
		return ""
	}
	
	var sb strings.Builder
	sb.WriteString("multiple errors occurred:\n")
	for _, e := range m.errors {
		sb.WriteString("- ")
		sb.WriteString(e.Error())
		sb.WriteString("\n")
	}
	return sb.String()
}

// HasErrors evaluates if any valid errors were recorded.
func (m *MultiError) HasErrors() bool {
	return len(m.errors) > 0
}

// Errors returns the slice of accumulated errors.
func (m *MultiError) Errors() []error {
	return m.errors
}

// AsError returns the MultiError as a standard interface.
// Crucially returns explicitly `nil` if there are no inner errors recorded.
func (m *MultiError) AsError() error {
	if len(m.errors) == 0 {
		return nil
	}
	return m
}
