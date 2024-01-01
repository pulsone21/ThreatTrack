package types

type Entity interface {
	ScanTo(ScanFunc) error
}
