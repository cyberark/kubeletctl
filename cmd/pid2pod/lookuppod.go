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
	"github.com/jedib0t/go-pretty/table"
	"os"
	"regexp"
	"strings"
)

// ID identifies a single container running in a Kubernetes Pod
type ID struct {
	Namespace string
	PodName   string
	//ContainerID   string
	ContainerName string
}

// LookupPod looks up a process ID from the host PID namespace, returning its Kubernetes identity.
func LookupPod(pid int, executable string, podInfo *podList, tw table.Writer) (*ID, error) {
	cid, err := LookupContainerID(pid)
	if err != nil {
		return nil, err
	}

	var containerRuntime string
	pidAndProcExecutable := fmt.Sprintf("%d (%s)", pid, executable)
	for _, item := range podInfo.Items {
		for _, status := range item.Status.ContainerStatuses {
			if strings.Contains(status.ContainerID, "containerd") {
				containerRuntime = "containerd"
			}
			if strings.Contains(status.ContainerID, "docker") {
				containerRuntime = "docker"
			}
			containerID := fmt.Sprintf("%s://%s", containerRuntime, cid)
			if status.ContainerID == containerID {
				tw.AppendRow([]interface{}{pidAndProcExecutable, item.Metadata.Name, item.Metadata.Namespace, status.Name})
				return &ID{
					Namespace: item.Metadata.Namespace,
					PodName:   item.Metadata.Name,
					//ContainerID:   cid,
					ContainerName: status.Name,
				}, nil
			}
		}
	}
	return nil, nil
}

// LookupContainerID looks up a process ID from the host PID namespace,
// returning its Docker and Containerd container ID.
func LookupContainerID(pid int) (string, error) {
	f, err := os.Open(fmt.Sprintf("/proc/%d/cgroup", pid))
	if err != nil {
		// this is normal, it just means the PID no longer exists
		return "", nil
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := kubePattern.FindStringSubmatch(line)
		if parts != nil {
			if len(parts) > 1 {
				return parts[1], nil
			}
		}
	}
	return "", nil
}

var (
	kubePattern = regexp.MustCompile(`(?:docker-|cri-containerd-)?([a-f0-9]{64})\.scope`)
)
