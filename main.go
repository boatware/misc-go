package misc_go

import (
	"fmt"
	"math"
	"time"
)

// MicroTime returns the current time in microseconds
func MicroTime() float64 {
	return MicroTimeByTimezone("UTC")
}

// MicroTimeByTimezone returns the current time in microseconds by a timezone name
func MicroTimeByTimezone(timezone string) float64 {
	loc, _ := time.LoadLocation(timezone)
	return MicroTimeByLocation(loc)
}

// MicroTimeByLocation returns the current time in microseconds by a timezone location
func MicroTimeByLocation(location *time.Location) float64 {
	now := time.Now().In(location)
	micSeconds := float64(now.Nanosecond()) / 1000000000
	return float64(now.Unix()) + micSeconds
}

// Round rounds a float64 to a given unit
//
// Example: Round(1.234, 0.1) = 1.2
func Round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

// FormatMicroTimeDuration formats a duration in microseconds to a string in the format: DD:HH:MM:SS.MICROSECONDS
func FormatMicroTimeDuration(duration float64) string {
	// Format: DD:HH:MM:SS.MICROSECONDS
	formatString := "%02d:%02d:%02d:%02d.%06d"

	days := int(duration / 86400)
	duration -= float64(days * 86400)

	hours := int(duration / 3600)
	duration -= float64(hours * 3600)

	minutes := int(duration / 60)
	duration -= float64(minutes * 60)

	seconds := int(duration)
	duration -= float64(seconds)

	microSeconds := int(duration * 1000000)
	return fmt.Sprintf(formatString, days, hours, minutes, seconds, microSeconds)
}
