package validator

import (
	"errors"
	"fmt"
)

type MissingVariableError struct {
	variable string
}

func NewMissingVariable(variable string) *MissingVariableError {
	return &MissingVariableError{variable: variable}
}

func (m MissingVariableError) Error() string {
	return fmt.Sprintf("missing required environment variable %s. Please set it (e.g., export %s=xxx) and try again", m.variable, m.variable)
}

type InvalidVariableError struct {
	variable string
}

func NewInvalidVariableError(variable string) *InvalidVariableError {
	return &InvalidVariableError{variable: variable}
}

func (m InvalidVariableError) Error() string {
	return fmt.Sprintf("variable '%s' fail validation rule", m.variable)
}

var (
	ErrRead      = errors.New("failed to read config")
	ErrValidate  = errors.New("failed to validate config")
	ErrUnmarshal = errors.New("failed to unmarshal config")
)
