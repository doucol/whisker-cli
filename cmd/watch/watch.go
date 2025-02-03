package watch

import (
	"github.com/spf13/cobra"
)

var WatchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch a variety of Calico resources",
	Long:  `Watch a variety of Calico resources`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmd.Help(); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	WatchCmd.AddCommand(WatchFlowsCmd)
}
