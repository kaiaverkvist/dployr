package compose

import (
	"io/ioutil"
	"path"
)

// IsDeployableDirectory looks for docker-compose.yml in a directory specified.
func GetDockerCompose(dir string) (string, error) {
	content, err := ioutil.ReadFile(path.Join(dir, "docker-compose.yml"))
	if err != nil {
		return "", err
	}

	return string(content), nil
}
