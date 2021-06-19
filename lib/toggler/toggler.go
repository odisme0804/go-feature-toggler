package toggler

import "errors"

var (
	ErrFlagNotFound = errors.New("flag not found")
)

type Entity struct {
	ID      string
	FlagKey string
	Payload interface{}
}
