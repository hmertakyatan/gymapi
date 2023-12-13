package helpers

import "time"

func DDmmYYTimeConverterFromStringToTime(StrDate string) (Date time.Time, err error) {

	parsedDate, err := time.Parse("02/01/2006", StrDate)
	if err != nil {
		return time.Time{}, err
	}

	return parsedDate, nil
}
