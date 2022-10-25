package util

import "strings"

func FormatCityName(name string) string {
	result := strings.ToLower(name)
	result = strings.ReplaceAll(result, " ", "_")
	result = strings.ReplaceAll(result, "'", "")

	return result
}
