package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"verifyMyTest/m/routes"
	"verifyMyTest/m/util"
)

func init() {
	var serverPort string
	defaultServerPort := viper.GetString("SERVER_PORT")
	serverCmd.PersistentFlags().StringVar(&serverPort, "port", defaultServerPort, "Server port")
	err := viper.BindPFlag("SERVER_PORT", serverCmd.PersistentFlags().Lookup("port"))
	if err != nil {
		util.Error(util.FatalError, err.Error())
		os.Exit(1)
	}

	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run server",
	Run: func(cmd *cobra.Command, args []string) {

		port := viper.GetString("SERVER_PORT")

		if port == "" {
			port = "8888"
			msg := fmt.Sprintf("SERVER_PORT is empty, setting default port %s.", port)
			util.Warn(util.Informational, msg)
		}

		err := RouteSetup(port)
		if err != nil {
			util.Error(util.FatalError, err.Error())
			os.Exit(1)
		}
	},
}

func RouteSetup(port string) error {
	router := routes.Setup()
	err := router.Run(":" + port)
	if err != nil {
		return err
	}
	return nil
}
