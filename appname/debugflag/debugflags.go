package main

import (
	"fmt"
	"github.com/spf13/cobra"
	//"k8s.io/klog"
	"k8s.io/kubernetes/cmd/kube-scheduler/app/options"
	"k8s.io/kubernetes/pkg/scheduler/framework/runtime"
	"os"
)

type Option func(runtime.Registry) error

func runCommand(cmd *cobra.Command, opts *options.Options, registryOptions ...Option) error {
	//fmt.Printf("cmd.Flags: %v", cmd.Flags())
	//fmt.Printf("opts: %v", opts)
	//fmt.Printf("registryOptions: %v", registryOptions)
    return nil
}


func NewSchedulerCommand(registryOptions ...Option) *cobra.Command {
	opts, _ := options.NewOptions()
	fmt.Printf("new options: %v", *opts)

	cmd := &cobra.Command{
		Use: "kube-scheduler",
		Long: `The Kubernetes scheduler is a control plane process which assigns
Pods to Nodes. The scheduler determines which Nodes are valid placements for
each Pod in the scheduling queue according to constraints and available
resources. The scheduler then ranks each valid Node and binds the Pod to a
suitable Node. Multiple different schedulers may be used within a cluster;
kube-scheduler is the reference implementation.
See [scheduling](https://kubernetes.io/docs/concepts/scheduling-eviction/)
for more information about scheduling and the kube-scheduler component.`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := runCommand(cmd, opts, registryOptions...); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}
	//fs := cmd.Flags()
	//namedFlagSets := opts.Flags()
	//fmt.Printf("fs: %v", *fs)
	//fmt.Printf("namedFlagSets: %v", namedFlagSets)
	//verflag.AddFlags(namedFlagSets.FlagSet("global"))
	//globalflag.AddGlobalFlags(namedFlagSets.FlagSet("global"), cmd.Name())
	//for _, f := range namedFlagSets.FlagSets {
	//	fs.AddFlagSet(f)
	//}
	//
	//usageFmt := "Usage:\n  %s\n"
	//cmd.SetUsageFunc(func(cmd *cobra.Command) error {
	//	fmt.Fprintf(cmd.OutOrStderr(), usageFmt, cmd.UseLine())
	//	return nil
	//})
	//cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
	//	fmt.Fprintf(cmd.OutOrStdout(), "%s\n\n"+usageFmt, cmd.Long, cmd.UseLine())
	//})
	//cmd.MarkFlagFilename("config", "yaml", "yml", "json")

	return cmd
}

func main(){
	command := NewSchedulerCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
