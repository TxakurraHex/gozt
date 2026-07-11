package source

import "gozt/cpe"

type Source interface {
	Name() string
	Collect() ([]cpe.CpeEntry, error)
}
