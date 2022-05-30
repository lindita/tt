package server

import (
	"github.com/spf13/cobra"
	"tt.com/tt/internal/conf"
	"tt.com/tt/internal/server"
)


var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "s",
	Run: func(cmd *cobra.Command, args []string) {
		controller, _, _ := newController(conf.GetConfig())
		server.Run(controller)
	},
}
