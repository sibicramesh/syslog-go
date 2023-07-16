package common

import "fmt"

// CalculatePriority calculates priority based on the given
// facility and severity. Complete RFC compliant.
// https://datatracker.ietf.org/doc/html/rfc5424#section-6.2.1
func CalculatePriority(facility int, severity int) (int, error) {

	if facility < 0 || facility > 23 {
		return -1, fmt.Errorf("invalid facility: %d", facility)
	}

	if severity < 0 || severity > 7 {
		return -1, fmt.Errorf("invalid severity: %d", facility)
	}

	priority := (facility * 8) + severity

	return priority, nil
}
