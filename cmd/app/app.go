package main

import (
	"fmt"
	"github.com/kaiaverkvist/dployr/pkg/deploy"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "dployr",
		Short: "",
		Long:  "",
	}

	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize dployr with a repository URL",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Info(fmt.Sprintf("Using repository URL '%s'", repository))
			log.Info(fmt.Sprintf("Looking for private key for server authentication @ '%s'", privateKeypath))

			deployment = deploy.NewDeployment(repository, host, user, privateKeypath)

			log.Info(deployment.DialGetUser())

			return nil
		},
	}
)

var (
	homedir, _ = os.UserHomeDir()
)

var (
	logStyle      = "${time_rfc3339}  ${level} "
	pubKeyDefault = fmt.Sprintf("%s/.ssh/id_rsa", homedir)

	// Variables from command line flags:
	// - Repository URL string
	repository string
	// - Public Key path
	privateKeypath string

	host string
	user string

	// Stores the deployment metadata and performs commands.
	deployment *deploy.Deployment
)

func init() {
	log.SetHeader(logStyle)

	// --repo or -r
	initCmd.Flags().StringVarP(&repository, "repo", "r", "", "repository url to initialize")
	_ = initCmd.MarkFlagRequired("repo")

	// --host or -ho
	initCmd.Flags().StringVarP(&host, "host", "", "", "host including port")
	_ = initCmd.MarkFlagRequired("host")

	// --user or -u
	initCmd.Flags().StringVarP(&user, "user", "u", "", "username to login with")
	_ = initCmd.MarkFlagRequired("user")

	// --public-key or -k
	initCmd.Flags().StringVarP(&privateKeypath, "private-key", "k", pubKeyDefault, "Private key path")

	rootCmd.AddCommand(initCmd)
}

func main() {
	_ = rootCmd.Execute()
}
