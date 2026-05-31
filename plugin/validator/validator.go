package validator

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/thumbrise/ouro/contract"
	"github.com/thumbrise/ouro/reference/pkg/reflection"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	vld := validator.New(validator.WithRequiredStructEnabled())
	return &Validator{
		validator: vld,
	}
}

func (c Validator) Validate(ctx context.Context, validateContext contract.ValidateContext) error {
	out := validateContext.Data
	if !reflection.IsStruct(out) && !reflection.IsStructPtr(out) {
		return nil
	}
	key := validateContext.Key
	err := c.validator.StructCtx(ctx, out)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrValidate, c.mapValidationErr(err, key))
	}

	return err
}

func (c *Validator) mapValidationErr(err error, viperKey string) error {
	var validationErrors validator.ValidationErrors

	ok := errors.As(err, &validationErrors)
	if !ok {
		return err
	}

	var result error
	for _, fe := range validationErrors {
		result = errors.Join(result, c.mapFieldErr(fe, viperKey))
	}

	return result
}

func (c *Validator) mapFieldErr(fe validator.FieldError, viperKey string) error {
	if fe.Tag() == "" {
		return fe
	}

	_, varName, _ := strings.Cut(fe.StructNamespace(), ".")

	if viperKey != "" {
		varName = viperKey + "." + varName
	}

	varName = strings.ToUpper(strings.ReplaceAll(varName, ".", "_"))
	if prefix := c.viper.GetEnvPrefix(); prefix != "" {
		varName = prefix + "_" + varName
	}

	if fe.Tag() == "required" {
		return NewMissingVariable(varName)
	}

	return fmt.Errorf("%w: %w", NewInvalidVariableError(varName), fe)
}
