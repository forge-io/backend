package cmd

import (
	"authenticate/router"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const AppVersion = "1.0.1"

func init() {
	rootCmd.AddCommand(daemonCmd)
	rootCmd.AddCommand(versionCmd)
}

func runDaemon(cmd *cobra.Command, args []string) {
	parentEnvPath, err := filepath.Abs(filepath.Join(".", ".env"))
	if err != nil {
		log.Fatal().Msgf("Error finding absolute path: %v", err)
	}

	// Load the parent .env file
	err = godotenv.Load(parentEnvPath)
	if err != nil {
		log.Fatal().Msgf("Error loading .env file: %v", err)
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	log.Info().Msg("Starting the service as daemon")

	port, err := strconv.Atoi(os.Getenv("AUTHENTICATE_PORT"))
	if err != nil {
		panic(err)
	}

	e := router.New()
	err = e.Start(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
}

var rootCmd = &cobra.Command{}

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Run the application as a daemon.",
	Long:  "This command will run the service.",
	Run:   runDaemon,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of the service.",
	Long:  "Print the current version used to run the service in the given binary",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Application Version : %s\n", AppVersion)
	},
}

func Execute() error {
	return rootCmd.Execute()
}
