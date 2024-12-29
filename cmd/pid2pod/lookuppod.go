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
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/jedib0t/go-pretty/table"
)

// ID identifies a single container running in a Kubernetes Pod
type ID struct {
	Namespace string
	PodName   string
	//ContainerID   string
	ContainerName string
}

func LookupPod(pid int, executable string, podInfo *podList, tw table.Writer) (*ID, error) {
	containerIDs, err := LookupContainerID(pid)
	if err != nil {
		return nil, err
	}

	if len(containerIDs) == 0 {
		return nil, nil
	}

	pidAndProcExecutable := fmt.Sprintf("%d (%s)", pid, executable)

	for _, item := range podInfo.Items {
		for _, status := range item.Status.ContainerStatuses {
			for _, extractedID := range containerIDs {
				var runtime string
				if strings.HasPrefix(status.ContainerID, "docker://") {
					runtime = "docker"
				} else if strings.HasPrefix(status.ContainerID, "containerd://") {
					runtime = "containerd"
				} else {
					fmt.Printf("Unknown runtime for Pod ContainerID: %s\n", status.ContainerID)
					continue
				}

				formattedContainerID := fmt.Sprintf("%s://%s", runtime, extractedID)

				if status.ContainerID == formattedContainerID {
					tw.AppendRow([]interface{}{pidAndProcExecutable, item.Metadata.Name, item.Metadata.Namespace, status.Name})
					return &ID{
						Namespace:     item.Metadata.Namespace,
						PodName:       item.Metadata.Name,
						ContainerName: status.Name,
					}, nil
				}
			}
		}
	}
	return nil, nil
}

// LookupContainerID looks up a process ID from the host PID namespace,
// returning its Docker and Containerd container ID.
func LookupContainerID(pid int) ([]string, error) {
	f, err := os.Open(fmt.Sprintf("/proc/%d/cgroup", pid))
	if err != nil {
		return nil, nil // PID might not exist
	}
	defer f.Close()

	var containerIDs []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		// Extract all matching container IDs from the line
		matches := kubePattern.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if len(match) > 1 {
				containerIDs = append(containerIDs, match[1])
			}
		}
	}
	return containerIDs, nil
}

var (
	kubePattern = regexp.MustCompile(`(?:docker-|cri-containerd-)?([a-f0-9]{64})\.scope`)
)
