package command

import (
	"github.com/spf13/cobra"
	"verifyMyTest/m/database/migrations"
	"verifyMyTest/m/util"
)

func init() {
	rootCmd.AddCommand(migrate)
}

var migrate = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate",
	Run: func(cmd *cobra.Command, args []string) {

		db, err := util.NewDb()
		if err != nil {
			util.Error(util.FatalError, err.Error())
		}
		db.AutoMigrate(&migrations.Customer{})
		util.Info(util.Informational, "Migration Done")
	},
}
