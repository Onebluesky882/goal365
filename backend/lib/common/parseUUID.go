package common

import (
	"fmt"

	"github.com/google/uuid"
)

func ParseUUID(idQuery string) error {

	_, err := uuid.Parse(idQuery)
	if err != nil {
		return fmt.Errorf("invalid uuid: %w", err)
	}
	return nil
}
