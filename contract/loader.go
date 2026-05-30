package contract

import "context"

type LoadContext struct {
	Key  string
	Data any
}
type Loader interface {
	Load(ctx context.Context, loadContext LoadContext) error
}
