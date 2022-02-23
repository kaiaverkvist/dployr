package main

import (
	"fmt"
	"github.com/kaiaverkvist/dployr/internal/logging"
	"github.com/kaiaverkvist/dployr/internal/web"
	"github.com/kaiaverkvist/dployr/internal/web/context"
	"github.com/kaiaverkvist/dployr/pkg/deploy"
	"github.com/kaiaverkvist/dployr/version"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
	"os"
	"sync"
)

var (
	// Used to find the user's ssh key path if it isn't specified by the user themselves.
	homedir, _ = os.UserHomeDir()
)

var (
	pubKeyDefault      = fmt.Sprintf("%s/.ssh/id_rsa", homedir)
	dockerHostTemplate = "DOCKER_HOST=\"ssh://%s@%s\""

	// Variables from command line flags:
	// - Directory
	directory string
	// - Public Key path
	privateKeypath string
	// - Host IP
	host string
	// - Username
	user string

	// Stores the deployment metadata and performs commands.
	deployment *deploy.Deployment

	wg sync.WaitGroup
)

var (
	rootCmd = &cobra.Command{
		Use:   "dployr",
		Short: "",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {

			log.Info("Initializing dployr @ version: ", version.BuildVersion, " / with time build time: ", version.BuildTime)

			log.Info(fmt.Sprintf("Looking for private key for server authentication @ '%s'", privateKeypath))

			log.Info(fmt.Sprintf("Attempting to find docker-compose in directory @ '%s'", directory))
			deployment, err := deploy.NewDeployment(directory, host, user, privateKeypath)
			if err != nil {
				log.Error("Unable to create deployment: ", err)
				return err
			}

			wdc := context.NewWebDataContainer(deployment)

			wg.Add(1)
			go web.CreateServer(&wdc)
			wg.Wait()

			return nil
		},
	}
)

func init() {
	log.SetHeader(logging.LogStyle)

	workDir, err := os.Getwd()
	if err != nil {
		log.Warn("Unable to find working directory: ", err)
		log.Error("Please use --directory to specify instead!")
		directory = ""
	}

	// --dir or -d
	rootCmd.Flags().StringVarP(&directory, "directory", "d", workDir, "repository directory on local machine")
	_ = rootCmd.MarkFlagRequired("directory")

	// --host
	rootCmd.Flags().StringVarP(&host, "host", "", "", "host of remote server")
	_ = rootCmd.MarkFlagRequired("host")

	// --user or -u
	rootCmd.Flags().StringVarP(&user, "user", "u", "", "login username on remote")
	_ = rootCmd.MarkFlagRequired("user")

	// --private-key or -k
	rootCmd.Flags().StringVarP(&privateKeypath, "private-key", "k", pubKeyDefault, "private key path")
}

func main() {
	_ = rootCmd.Execute()
}
