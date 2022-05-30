package server

import (
	"github.com/spf13/cobra"
)


var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "s",
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}
