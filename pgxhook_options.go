package pgxhook

type HookConnOption interface {
	apply(s *HookConn)
}

type funcHookConnOption func(s *HookConn)

func (f funcHookConnOption) apply(s *HookConn) {
	f(s)
}

func WithHooks(hooks ...FullHook) HookConnOption {
	return funcHookConnOption(func(s *HookConn) {
		for _, h := range hooks {
			s.beforeHooks = append(s.beforeHooks, h)
			s.afterHooks = append(s.afterHooks, h)
		}
	})
}

func WithBeforeHooks(hooks ...BeforeHook) HookConnOption {
	return funcHookConnOption(func(s *HookConn) {
		s.beforeHooks = append(s.beforeHooks, hooks...)
	})
}

func WithAfterHooks(hooks ...AfterHook) HookConnOption {
	return funcHookConnOption(func(s *HookConn) {
		s.afterHooks = append(s.afterHooks, hooks...)
	})
}
