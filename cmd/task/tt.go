package task

import (
	"fmt"
	"github.com/spf13/cobra"
	"tt.com/tt/internal/conf"
)

func ttCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "tt",
		Run: func(cmd *cobra.Command, args []string) {
			service, _, _ := newService(conf.GetConfig())
			ret := service.Tt.GetTt(cmd.Context())
			fmt.Println("tt cmd run success with result:" + ret)
		},
	}
	return cmd
}
