/*
Copyright (c) 2020 CyberArk Software Ltd. All rights reserved

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package debug

import (
	"fmt"
	"kubeletctl/cmd"
	"kubeletctl/pkg/api"
	"os"

	"github.com/spf13/cobra"
)

// debugCmd represents the debug command
var debugCmd = &cobra.Command{
	Use:   "debug <pprof_profiles>",
	Short: "Return debug information (pprof or flags)",
	Long: `Description:
  Return pprof information. 
  pprof is a tool for visualization and analysis of profiling data. 
  It reads a collection of profiling samples in profile.proto format and generates reports to visualize and help analyze the data.
  
  HTTP request: 
    GET /debug/pprof/<profile>
    GET /debug/flags/v
    PUT /debug/flags/v (include body)

  Examples:
    /debug/pprof/                    // to view all available profiles
    /debug/pprof/profile
    /debug/pprof/profile?seconds=30 // look at 30-second CPU profile
    /debug/pprof/symbol
    /debug/pprof/cmdline
    /debug/pprof/trace
    /debug/pprof/trace?seconds=5    // collect 5-second execution trace
    /debug/pprof/block              // look at the goroutine blocking profile
    /debug/pprof/mutex              // look at the holderse of contended mutexes
    /debug/pprof/heap               // look at the heap profile
  
  Profile Descriptions:	
    "allocs":       A sampling of all past memory allocations.
    "block":        Stack traces that led to blocking on synchronization primitives.
    "cmdline":      The command line invocation of the current program.
    "goroutine":    Stack traces of all current goroutines.
    "heap":         A sampling of memory allocations of live objects. 
                    You can specify the gc GET parameter to run GC before taking the heap sample.
    "mutex":        Stack traces of holders of contended mutexes.
    "profile":      CPU profile. You can specify the duration in the seconds GET parameter. After you get the profile file, 
                    use the go tool pprof command to investigate the profile.
    "threadcreate": Stack traces that led to the creation of new OS threads.
    "trace":        A trace of execution of the current program. You can specify the duration in the seconds GET parameter. 
                    After you get the trace file, use the go tool trace command to investigate the trace.
  
  Example for usage:
    kubeletctl debug symbol
  
  With curl:
    curl -k https://<node_ip>:10250/debug/pprof/
    curl -k https://<node_ip>:10250/debug/pprof/symbol
    curl -k https://<node_ip>:10250/debug/pprof/trace?seconds=5`,
	Run: func(cmd2 *cobra.Command, args []string) {
		//fmt.Println("debug called")

		// TODO: should we use command for each profile description ?

		// TODO: write in the description how to use more complex commands like: /debug/pprof/profile?seconds=30
		// It should be kubeletctl debug profile?seconds=30
		// Or we can add a flag for this

		var apiPathUrl string
		var inputArgs string
		if args == nil {
			fmt.Println("[*] No debug profile was specified")
			os.Exit(1)
		} else {
			inputArgs = args[0]
		}

		if secondsFlag != "" {
			apiPathUrl = cmd.ServerFullAddressGlobal + api.DEBUG + "/" + inputArgs + "?seconds=" + secondsFlag
		} else {
			apiPathUrl = cmd.ServerFullAddressGlobal + api.DEBUG + "/" + inputArgs
		}

		resp, err := api.GetRequest(api.GlobalClient, apiPathUrl)
		cmd.PrintPrettyHttpResponse(resp, err)
	},
}

var secondsFlag string

func init() {
	cmd.RootCmd.AddCommand(debugCmd)
	debugCmd.PersistentFlags().StringVarP(&secondsFlag, "seconds", "", "", "Number of seconds to look. Used in debug profile and trace")
}
