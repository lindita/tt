package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"tt.com/tt/cmd/server"
	"tt.com/tt/cmd/task"
	"tt.com/tt/internal/conf"
	"tt.com/tt/internal/logger"
)

var rootCmd = &cobra.Command{
	Use:   "tt",
	Short: "tt",
}

func Execute()  {
	//初始化配置文件和日志
	cobra.OnInitialize(func() {
		conf.InitConfig()
		logger.NewZapLog()
	})

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}



func init()  {
	//server
	rootCmd.AddCommand(server.ServerCmd)
	//tasks
	task.InitTaskCmd(rootCmd)
}