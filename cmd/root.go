package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"kubeletctl/pkg/api"
	"os"
)

var (
	PortFlag                string
	NamespaceFlag           string
	ContainerFlag           string
	PodFlag                 string
	ServerIpAddressFlag     string
	ServerFullAddressGlobal string
	PodUidFlag              string
	KubeConfigFlag          string
	ProtocolScheme          string
	caFlag                  string
	certFlag                string
	keyFlag                 string
	HttpFlag                bool
	//BodyContentFlag         string
	RawFlag bool
)

// TODO: Consider the use of "go-prompt" for auto-completion of dynamic resources like pods
//  Linke: https://github.com/c-bata/go-prompt
// Current auto completion in cobra is only for bash: https://github.com/spf13/cobra/blob/master/bash_completions.md
var RootCmd = &cobra.Command{
	Use:   "kubeletctl",
	Short: "kubeletctl is command line utitly that implements kuebelt's API",
	Long: `Description:
  kubeletctl is command line utility that implements kuebelt's API.
  It also provides scanning for opened kubelet APIs and search for potential RCE on containers.
  
  You can view examples from each command by using the -h\--help flag like that: kubeletctl run -h
  Examples:  
    // List all pods from kubelet
    kubeletctl pods --server 123.123.123.123 

    // Searching for service account token in each accessible container
    kubeletctl scan token --server 123.123.123.123 

    // Searching for pods/containers vulnerable to RCE
    kubeletctl scan rce --server 123.123.123.123 

    // Run "ls /" command on pod my-nginx-pod/nginx in thedefault namespace
    kubeletctl run "ls /" --namespace default --pod my-nginx-pod --container nginx --server 123.123.123.123 

    // Run "ls /" command on all existing pods in a node
    kubeletctl.exe run "ls /" --all-pods --server 123.123.123.123 

    // With certificates
    kubeletctl.exe pods -s <node_ip> --cacert C:\Users\myuser\certs\ca.crt --cert C:\Users\myuser\certs\kubelet-client-current.pem --key C:\Users\myuser\certs\kubelet-client-current.pem
`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		printLogo()
	},
}

func Execute() {

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// List of command examples:
// https://github.com/kubernetes/kubernetes/blob/14344b57e56258e87cbe80c8cd80399855eca424/pkg/kubelet/server/auth_test.go#L110-L143
//https://towardsdatascience.com/how-to-create-a-cli-in-golang-with-cobra-d729641c7177
func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVarP(&PortFlag, "port", "", "", "Kubelet's port, default is 10250")
	RootCmd.PersistentFlags().StringVarP(&NamespaceFlag, "namespace", "n", "", "pod namespace")
	RootCmd.PersistentFlags().StringVarP(&ContainerFlag, "container", "c", "", "Container name")
	RootCmd.PersistentFlags().StringVarP(&PodFlag, "pod", "p", "", "Pod name")
	RootCmd.PersistentFlags().StringVarP(&PodUidFlag, "uid", "u", "", "Pod UID")
	RootCmd.PersistentFlags().StringVarP(&KubeConfigFlag, "config", "k", "", "KubeConfig file")
	RootCmd.PersistentFlags().BoolVarP(&RawFlag, "raw", "r", false, "Prints raw data")
	RootCmd.PersistentFlags().BoolVarP(&HttpFlag, "http", "", false, "Use HTTP (default is HTTPS)")
	//RootCmd.PersistentFlags().StringVarP(&BodyContentFlag, "body", "b", "", "This is the body message. Should be used in POST or PUT requests.")

	RootCmd.PersistentFlags().StringVarP(&caFlag, "cacert", "", "", "CA certificate (example: /etc/kubernetes/pki/ca.crt )")
	RootCmd.PersistentFlags().StringVarP(&certFlag, "cert", "", "", "Private key (example: /var/lib/kubelet/pki/kubelet-client-current.pem)")
	RootCmd.PersistentFlags().StringVarP(&keyFlag, "key", "", "", "Digital certificate (example: /var/lib/kubelet/pki/kubelet-client-current.pem)")

	pf := RootCmd.PersistentFlags()
	pf.StringVarP(&ServerIpAddressFlag, "server", "s", "", "Server address (format: x.x.x.x. For Example: 123.123.123.123)")
	//cobra.MarkFlagRequired(pf, "server")
}

const KUBELET_DEFAULT_PORT = "10250"

func initConfig() {
	if PortFlag == "" {
		PortFlag = KUBELET_DEFAULT_PORT
	}

	if ServerIpAddressFlag == "" {
		ServerIpAddressFlag = "127.0.0.1"
	}

	ProtocolScheme = "https"
	if HttpFlag {
		ProtocolScheme = "http"
	}

	if NamespaceFlag == "" {
		NamespaceFlag = "default"
	}

	ServerFullAddressGlobal = fmt.Sprintf("%s://%s:%s", ProtocolScheme, ServerIpAddressFlag, PortFlag)

	if KubeConfigFlag != "" {
		api.InitGlobalClientFromFile(KubeConfigFlag)
	} else if caFlag != "" && certFlag != "" && keyFlag != "" {
		api.InitGlobalClientFromCertificatesFiles(ServerFullAddressGlobal, caFlag, certFlag, keyFlag)
	} else {
		api.InitHttpClient()
	}
}
