package main

import (
	"flag"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	endpoint string
	nodeID   string
)

func main() {

	cmd := &cobra.Command{
		Use:   "arstorplugin",
		Short: "CSI based Arstor driver", // 紧跟在命令下面的描述
		Run: func(cmd *cobra.Command, args []string) {
			handle()
		},
	}

	cmd.Flags().AddGoFlagSet(flag.CommandLine)

	cmd.PersistentFlags().StringVar(&nodeID, "nodeid", "", " usage node id")
	cmd.MarkPersistentFlagRequired("nodeid")

	cmd.PersistentFlags().StringVar(&endpoint, "endpoint", "", " usage CSI endpoint")
	cmd.MarkPersistentFlagRequired("endpoint")

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}

func handle(){
	fmt.Printf("This is handle!\n get nodeID: %s, endpoint:%s\n", nodeID, endpoint)
}