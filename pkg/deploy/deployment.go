package deploy

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/kaiaverkvist/dployr/pkg/compose"
	"github.com/labstack/gommon/log"
	"io"
	"os"
	"os/exec"
	"sync"
	"time"
)

var (
	tmpContextName = "tmp-dployr"
)

// Deployment represents a dployr deployment to a single host.
// Contains metadata about deployment status and should be the source of truth.
type Deployment struct {
	// The deployment host - usually an IP.
	FullHostName string

	// privKey should contain a public key (for authentication) if it exists.
	// The user should be able to override the public key.
	privKey string

	ExpectedEnvs []string

	OutputBuffer bytes.Buffer
	Mtx          sync.Mutex

	Dir string
}

// NewDeployment takes in the required parameters to initialize a deployment process.
func NewDeployment(dir string, host string, user string, privKey string) (*Deployment, error) {
	composeFile, err := compose.GetDockerCompose(dir)
	if err != nil {
		return nil, errors.New("No docker-compose.yml detected in directory " + dir)
	}

	// Should contain
	expectedEnvs := compose.GetExpectedEnvironmentVariables(composeFile)

	return &Deployment{
		FullHostName: fmt.Sprintf("%s@%s", user, host),
		privKey:      privKey,

		ExpectedEnvs: expectedEnvs,

		Dir: dir,
	}, nil
}

func (d *Deployment) PerformDeployment(envs map[string]string) (string, error) {
	out, err := exec.Command("docker", "context", "create", tmpContextName, "--docker", "\"host=ssh://"+d.FullHostName+"\"").CombinedOutput()
	if err != nil {
		log.Error("Could not create docker context for dployr: ", err.Error())
		log.Warn("Maybe it exists already?")
		//return "", err
	}
	log.Info("Created temporary docker context for dployr: ", string(out))

	log.Info("Using context ", tmpContextName, " temporarily during deployment.")
	_ = exec.Command("docker", "context", "use", tmpContextName).Run()
	log.Info("Context set, will now attempt to create..")

	time.Sleep(time.Second)

	cmd := exec.Command("docker-compose", "--context", tmpContextName, "up", "--build", "-d")
	cmd.Dir = d.Dir

	// OS env and then add the DOCKER_HOST variable.
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "DOCKER_CONTEXT="+tmpContextName)

	// https://github.com/docker/compose/issues/8218
	cmd.Env = append(cmd.Env, "COMPOSE_DOCKER_CLI_BUILD=0")

	// Format all the supplied env vars:
	for k, v := range envs {
		envVar := fmt.Sprintf("%s=%s", k, v)
		cmd.Env = append(cmd.Env, envVar)
	}

	log.Info("Attempting to run command: ", cmd.String())

	mw := io.MultiWriter(os.Stdout, &d.OutputBuffer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	err = cmd.Run()
	d.cleanupContext()
	return d.OutputBuffer.String(), err
}

func (d *Deployment) cleanupContext() {
	_ = exec.Command("docker", "context", "use", "default").Run()
	_ = exec.Command("docker", "context", "rm", tmpContextName).Run()
	log.Info("Cleaned up docker context!")
}
