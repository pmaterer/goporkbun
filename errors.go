package goporkbun

import (
	"fmt"
)

type Error struct {
	Message string `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("goporkbun: %s", e.Message)
}
