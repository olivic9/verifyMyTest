package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"verifyMyTest/m/util"
)

var rootCmd = &cobra.Command{
	Short: "Verify My Test, A test API",
	Long: `Verify My Test
                Application`,
	Run: func(cmd *cobra.Command, args []string) {
		util.Info(viper.GetString("APP_NAME") + " Started")
	},
}

func Execute() {

	rootCmd.Use = viper.GetString("APP_COMMAND")
	if err := rootCmd.Execute(); err != nil {
		util.Error(util.FatalError, err.Error())
		os.Exit(1)
	}
}
