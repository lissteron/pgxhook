package hooks

import (
	"context"

	"github.com/lissteron/pgxhook"
)

type calledFunc func(ctx context.Context, input *pgxhook.HookData) (context.Context, error)

type HookForTests struct {
	calledFunc calledFunc
}

func NewHookForTests(calledFunc calledFunc) *HookForTests {
	return &HookForTests{calledFunc: calledFunc}
}

func (h *HookForTests) Before(ctx context.Context, input *pgxhook.HookData) (context.Context, error) {
	return h.calledFunc(ctx, input)
}

func (h *HookForTests) After(ctx context.Context, input *pgxhook.HookData) (context.Context, error) {
	return h.calledFunc(ctx, input)
}
