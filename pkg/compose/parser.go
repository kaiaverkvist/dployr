package compose

import (
	"github.com/kaiaverkvist/dployr/internal/utils"
	"regexp"
	"strings"
)

const (
	dockerComposeEnvPattern = "\\${([^}]+)}"
)

// GetExpectedEnvironmentVariables attempts to parse a `docker-compose.yml` file to find a list of expected environment
// variables.
func GetExpectedEnvironmentVariables(composeFile string) []string {
	r := regexp.MustCompile(dockerComposeEnvPattern)

	matches := r.FindAllString(composeFile, -1)

	// Loop through and sanitize out the env tags.
	var envs []string
	for _, match := range matches {
		var m = match
		m = strings.ReplaceAll(m, "$", "")
		m = strings.ReplaceAll(m, "{", "")
		m = strings.ReplaceAll(m, "}", "")
		envs = append(envs, m)
	}

	return utils.UniqueNonEmptyElementsOf(envs)
}
