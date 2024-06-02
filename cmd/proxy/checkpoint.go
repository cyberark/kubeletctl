package proxy

import (
	"bytes"
	"fmt"
	"io"
	"kubeletctl/cmd"
	"kubeletctl/pkg/api"
	"kubeletctl/pkg/utils"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var checkpointCmd = &cobra.Command{
	Use:   "checkpoint",
	Short: "Taking a container snapshot",
	Long: `Description:
  Taking a container snapshot. 
  HTTP request: POST /checkpoint/<namespace>/<pod>/<container>
  Example for usage:
  kubeletctl.exe checkpoint
  
  With curl:
  curl -k -X POST https://<node_ip>:10250/checkpoint/<namespace>/<pod>/<container>
`,
	Run: func(cmd2 *cobra.Command, args []string) {
		apiPathUrl := ""
		if utils.AreNamespacePodAndContainerFlagsSet(cmd.NamespaceFlag, cmd.PodFlag, cmd.ContainerFlag) {
			apiPathUrl = fmt.Sprintf("%s%s/%s/%s/%s", cmd.ServerFullAddressGlobal, api.CHECKPOINT, cmd.NamespaceFlag, cmd.PodFlag, cmd.ContainerFlag)
		} else {
			fmt.Println("Please provide namespace, pod name and container.")
			os.Exit(1)
		}

		resp, err := api.PostRequest(api.GlobalClient, apiPathUrl, nil)
		// Eviatar: consider use of RawFlag ?
		if err != nil {
			fmt.Printf("[*] Failed to run HTTP request with error: %s\n", err)
			os.Exit(1)
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		cmd.PrintPrettyHttpResponse(resp, err)

	},
}

func init() {
	cmd.RootCmd.AddCommand(checkpointCmd)
}
