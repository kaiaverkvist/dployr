package deploy

import "github.com/helloyi/go-sshclient"

// Deployment represents a dployr deployment to a single host.
// Contains metadata about deployment status and should be the source of truth.
type Deployment struct {
	// URL to the remote repository.
	repository string

	// The deployment host - usually an IP.
	host string
	user string

	// privKey should contain a public key (for authentication) if it exists.
	// The user should be able to override the public key.
	privKey string
}

// NewDeployment takes in the required parameters to initialize a deployment process.
func NewDeployment(repo string, host string, user string, privKey string) *Deployment {
	return &Deployment{
		repository: repo,
		host:       host,
		user:       user,
		privKey:    privKey,
	}
}

func (d *Deployment) DialGetUser() (string, error) {
	client, err := sshclient.DialWithKey(d.host, d.user, d.privKey)
	if err != nil {
		return "", err
	}

	out, err := client.Cmd("whoami").Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}
