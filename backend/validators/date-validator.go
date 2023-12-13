package validators

import "time"

func IsValidDate(dateString string) bool {
	_, err := time.Parse("02/01/2006", dateString)
	return err == nil
}
