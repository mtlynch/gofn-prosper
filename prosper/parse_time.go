package prosper

import "time"

const (
	prosperDateFormat      = "2006-01-02"
	prosperTimeFormat      = "2006-01-02 15:04:05 +0000"
	prosperOldTimeFormat   = "01022006"
	dateOlderThanFiveYears = "1"
	dateOlderThanTenYears  = "2"
)

func parseProsperDate(dateSerialized string) (time.Time, error) {
	return parseTimeAllowEmpty(dateSerialized, prosperDateFormat)
}

func parseProsperTime(timeSerialized string) (time.Time, error) {
	return parseTimeAllowEmpty(timeSerialized, prosperTimeFormat)
}

func parseTimeAllowEmpty(timeSerialized, format string) (time.Time, error) {
	if len(timeSerialized) == 0 {
		return time.Time{}, nil
	}
	return time.Parse(format, timeSerialized)
}

// we treat "1" (older than 5 years) and "2" (older than 10 years) as nil
func parseProsperOldTime(timeSerialized string) (time.Time, error) {
	if timeSerialized == dateOlderThanFiveYears || timeSerialized == dateOlderThanTenYears {
		return time.Time{}, nil
	}
	return time.Parse(prosperOldTimeFormat, timeSerialized)
}
