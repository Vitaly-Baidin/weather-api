package util

import "strings"

func FormatCityName(name string) string {
	result := strings.ToLower(name)
	result = strings.ReplaceAll(name, " ", "_")

	return result
}
