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
	"github.com/spf13/cobra"
	"kubeletctl/cmd"
	"kubeletctl/pkg/api"
	"kubeletctl/pkg/utils"
	"net/http"
)

// https://github.com/kubernetes/kubernetes/pull/67629

// flagsCmd represents the flags command
var flagsCmd = &cobra.Command{
	Use:   "flags <body_content>",
	Short: "Return flags information",
	Long: `Description:
  Return flags information. 
  If you want to change the flags you need to add the body content (example: "1").
  Otherwise it will return the flags status.

  HTTP request: 
    PUT /debug/flags/v (body: "1")
    GET /debug/flags/v

  With curl:
    // Will return: "successfully set glog.logging.verbosity to 1"
    curl -k -X PUT https://<node_ip>:10250/debug/flags/v -d "1"
    
    // Will return: "successfully get glog.logging.verbosity: 1"
    curl -k https://<node_ip>:10250/debug/flags/v1`,
	Run: func(cmd2 *cobra.Command, args []string) {
		//fmt.Println("flags called")

		var resp *http.Response
		var err error

		if utils.IsNotArgsEmpty(args) {
			apiPathUrl := cmd.ServerFullAddressGlobal + api.DEBUG_FLAGS
			resp, err = api.PutRequest(api.GlobalClient, apiPathUrl, []byte(args[0]))
		} else {

			apiPathUrl := cmd.ServerFullAddressGlobal + api.DEBUG_FLAGS
			resp, err = api.GetRequest(api.GlobalClient, apiPathUrl)
		}

		cmd.PrintPrettyHttpResponse(resp, err)
	},
}

func init() {

	debugCmd.AddCommand(flagsCmd)
}
