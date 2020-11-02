package common

import (
	"fmt"
)

/*ContextKey - type for key used to store values into context */
type ContextKey string

// WhoPays for file downloading.
type WhoPays int

// possible variants
const (
	WhoPays3rdParty WhoPays = iota // 0, 3rd party user pays
	WhoPaysOwner                   // 1, file owner pays
)

// String implements fmt.Stringer interface.
func (wp WhoPays) String() string {
	switch wp {
	case WhoPays3rdParty:
		return "3rd_party"
	case WhoPaysOwner:
		return "owner"
	}
	return fmt.Sprintf("WhoPays(%d)", int(wp))
}

// Validate the WhoPays value.
func (wp WhoPays) Validate() (err error) {
	switch wp {
	case WhoPays3rdParty, WhoPaysOwner:
		return // ok
	}
	return fmt.Errorf("unknown WhoPays value: %d", int(wp))
}
