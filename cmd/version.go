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
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func printLogo() {
	logo := `
 _           _           _                         _  
| |         | |         | |         _          _  | | 
| |  _ _   _| |__  _____| | _____ _| |_ ____ _| |_| | 
| |_/ ) | | |  _ \| ___ | || ___ (_   _) ___|_   _) | 
|  _ (| |_| | |_) ) ____| || ____| | |( (___  | |_| | 
|_| \_)____/|____/|_____)\_)_____)  \__)____)  \__)\_)

Author: Eviatar Gerzi
Version: 1.1
`
	fmt.Println(logo)
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of the kubeletctl",
	Long: `Print the version of the kubeletctl`,
	Run: func(cmd2 *cobra.Command, args []string) {
		//fmt.Println("version called")
		printLogo()
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
