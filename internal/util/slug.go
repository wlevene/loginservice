package util

import (
	"regexp"
	"strings"
)

func CreateSlug(title string) string {

	lowercaseTitle := strings.ToLower(title)

	reg := regexp.MustCompile("[^a-z0-9]+")
	slug := reg.ReplaceAllString(lowercaseTitle, "-")

	slug = strings.Trim(slug, "-")

	return slug
}
