package contract

import "context"

type HookSuccessRead func(ctx context.Context, loadContext LoadContext)
