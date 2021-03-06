package cmd

import (
	"fmt"

	"github.com/gobuffalo/buffalo/buffalo/cmd/updater"
	"github.com/gobuffalo/buffalo/runtime"
	"github.com/spf13/cobra"
)

// updateCmd represents the info command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: fmt.Sprintf("will attempt to upgrade a Buffalo application to version %s", runtime.Version),
	RunE: func(cmd *cobra.Command, args []string) error {
		return updater.Run()
	},
}

func init() {
	decorate("update", RootCmd)
	RootCmd.AddCommand(updateCmd)
}
