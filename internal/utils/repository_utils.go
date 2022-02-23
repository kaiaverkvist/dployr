package utils

import "strings"

func GetRepositoryName(url string) string {

	var result string
	var parts []string

	parts = strings.Split(url, "/")
	result = parts[len(parts)-1]

	return result
}

func GetRepositoryUsername(url string) string {

	var result string
	var parts []string

	parts = strings.Split(url, "/")
	result = parts[len(parts)-2]

	return result
}
