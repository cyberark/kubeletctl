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
package pid2pod

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-ps"
	"io"
	"kubeletctl/cmd"
	"kubeletctl/pkg/api"
	"log"
	"net/http/httputil"
	"os"

	"github.com/spf13/cobra"
)

type podList struct {
	// We only care about namespace, serviceAccountName and containerID
	Metadata struct {
	} `json:"metadata"`
	Items []struct {
		Metadata struct {
			Namespace string            `json:"namespace"`
			Name      string            `json:"name"`
			UID       string            `json:"uid"`
			Labels    map[string]string `json:"labels"`
		} `json:"metadata"`
		Spec struct {
			ServiceAccountName string `json:"serviceAccountName"`
		} `json:"spec"`
		Status struct {
			ContainerStatuses []struct {
				ContainerID string `json:"containerID"`
				Name        string `json:"name"`
			} `json:"containerStatuses"`
		} `json:"status"`
	} `json:"items"`
}

// pid2podCmd represents the pid2pod command
var pid2podCmd = &cobra.Command{
	Use:   "pid2pod",
	Short: "That shows how Linux process IDs (PIDs) can be mapped to Kubernetes pod metadata",
	Long: `Description:
  That shows how Linux process IDs (PIDs) can be mapped to Kubernetes pod metadata.

  Example for usage:
  kubeletctl pid2pod
`,
	Run: func(cmd2 *cobra.Command, args []string) {
		apiPathUrl := cmd.ServerFullAddressGlobal + api.PODS
		resp, err := api.GetRequest(api.GlobalClient, apiPathUrl)
		respDump, err := httputil.DumpResponse(resp, true)
		if cmd.RawFlag {
			fmt.Printf("RESPONSE:\n%s", string(respDump))
		} else {
			if err != nil {
				fmt.Printf("[*] Failed to run HTTP request with error: %s\n", err)
				os.Exit(1)
			}

			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			var podInfo *podList
			err = json.Unmarshal(bodyBytes, &podInfo)

			if err != nil {
				// TODO: this function does some of the previous checks, consider modification
				cmd.PrintPrettyHttpResponse(resp, err)
			} else {
				// get system all processes
				processes, err := ps.Processes()
				if err != nil {
					log.Fatalf("could not list processes: %v", err)
				}
				for _, proc := range processes {
					if pidFlag != 0 {
						if pidFlag == proc.Pid() {
							pid = pidFlag
							printPid2Pod(pid, proc.Executable(), podInfo)
						}
					} else {
						pid = proc.Pid()
						printPid2Pod(pid, proc.Executable(), podInfo)
					}
				}
			}
		}
	},
}

func printPid2Pod(pid int, executable string, podInfo *podList) {
	id, err := LookupPod(pid, podInfo)
	if err != nil {
		log.Fatalf("could not get ID of process %d: %v", pid, err)
	}
	if id != nil {
		fmt.Printf("PID %d (%s): %+#v\n", pid, executable, id)
	}
}

var (
	pidFlag int
	pid     int
)

func init() {
	cmd.RootCmd.AddCommand(pid2podCmd)
	pid2podCmd.PersistentFlags().IntVarP(&pidFlag, "pid", "", 0, "Name of pid to look.")
}
