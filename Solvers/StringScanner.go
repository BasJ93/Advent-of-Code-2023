package Solvers

type IStringScanner interface {
	Scan() bool
	Text() string
}

type MockScanner struct {
	ScanFunc func() bool
	TextFunc func() string
}

func (s *MockScanner) Scan() bool {
	if s.ScanFunc != nil {
		return s.ScanFunc()
	}
	return false
}

func (s *MockScanner) Text() string {
	if s.TextFunc != nil {
		return s.TextFunc()
	}
	return ""
}
