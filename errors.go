package goporkbun

import (
	"fmt"
)

type errorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Error struct {
	Message string `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("goporkbun: %s", e.Message)
}
