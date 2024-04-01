package iomanager

type I0Manager interface {
	ReadLines() ([]string, error)
	WriteResult(data any) error
}
