package helpers

import "time"

func CalculateEndDate(StartDate time.Time, Month uint8) (EndDate time.Time, err error) {
	EndDate = StartDate.AddDate(0, int(Month), 0)
	return EndDate, nil
}
