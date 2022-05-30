package task

import (
	"fmt"
	"github.com/spf13/cobra"
)

func InitTaskCmd(rootCommand *cobra.Command)  {
	rootCommand.AddCommand(ttCmd())
}

func ttCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "tt",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("tt")
		},
	}
	return cmd
}

