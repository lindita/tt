package task

import (
	"github.com/spf13/cobra"
)

func InitTaskCmd(rootCommand *cobra.Command) {
	rootCommand.AddCommand(ttCmd())
}
