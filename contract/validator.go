package contract

import "context"

type ValidateContext struct {
	Key  string
	Data any
}
type Validator interface {
	Validate(ctx context.Context, validateContext ValidateContext) error
}
