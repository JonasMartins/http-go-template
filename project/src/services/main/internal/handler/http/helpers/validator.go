package helpers

import (
	"fmt"
	"project/src/services/main/domain/usecases"
)

func ValidateFindUserParams(params *usecases.GetUserParams) error {
	if params.Email != nil && params.Id != nil {
		return fmt.Errorf("need only a searchable field")
	} else if params.Email != nil && params.Uuid != nil {
		return fmt.Errorf("need only a searchable field")
	} else if params.Id != nil && params.Uuid != nil {
		return fmt.Errorf("need only a searchable field")
	}
	return nil
}
