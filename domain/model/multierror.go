package model

import "strings"

type MultiError struct {
	error
	errors []error
}

func NewMultiError(errs ...error) *MultiError {
	if errs == nil {
		errs = make([]error, 0)
	}
	return &MultiError{
		errors: errs,
	}
}

func (m *MultiError) Error() string {
	msg := make([]string, len(m.errors))
	for i, m := range m.errors {
		msg[i] = m.Error()
	}
	return strings.Join(msg, "\n")
}

func (m *MultiError) Add(err error) bool {
	if err != nil {
		m.errors = append(m.errors, err)
		return true
	}
	return false
}

func (m *MultiError) Extensions() map[string]interface{} {
	exception := make(map[string]interface{})
	exception["errors"] = m.errors

	return map[string]interface{}{
		"code":      CodeBadUserInput,
		"message":   m.Error(),
		"exception": exception,
	}
}

func (m *MultiError) HasError() bool {
	return len(m.errors) > 0
}
