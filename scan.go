package pgxhook

type scanWithError struct {
	err error
}

func (s *scanWithError) Scan(_ ...any) error {
	return s.err
}
