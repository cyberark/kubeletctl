[![GitHub release][release-img]][release]
[![License][license-img]][license]
[![Go version][shield-go-version]][go-version]
![Downloads][download]

<img src="https://github.com/cyberark/kubeletctl/blob/assets/kubeletctl_2x_transparent.png" width="260">  

## Overview
Kubeletctl is a command line tool that implement kubelet's API.  
Part of kubelet's API is documented but most of it is not.  
This tool covers all the documented and undocumented APIs.  
The full list of all kubelet's API can be view through the tool or this [API table](https://github.com/cyberark/kubeletctl/blob/master/API_TABLE.md).  
A related blog post:  
https://www.cyberark.com/resources/threat-research-blog/using-kubelet-client-to-attack-the-kubernetes-cluster

## What can it do ?
- Run any kubelet API call
- Scan for nodes with opened kubelet API
- Scan for containers with RCE
- Run a command on all the available containers by kubelet at the same time
- Get service account tokens from all available containers by kubelet
- Nice printing :)

## Installation  
On the [releases](https://github.com/cyberark/kubeletctl/releases) page you will find the latest releases with links based on the operating system.  

For the following examples, we will use the kubeletctl_linux_amd64 binary link. If you plan to use other link, change it accordingly.   
### wget
```
wget https://github.com/cyberark/kubeletctl/releases/download/v1.7/kubeletctl_linux_amd64 && chmod a+x ./kubeletctl_linux_amd64 && mv ./kubeletctl_linux_amd64 /usr/local/bin/kubeletctl
```  

### curl
```
curl -LO https://github.com/cyberark/kubeletctl/releases/download/v1.7/kubeletctl_linux_amd64 && chmod a+x ./kubeletctl_linux_amd64 && mv ./kubeletctl_linux_amd64 /usr/local/bin/kubeletctl
```

## Usage
kubeletctl works similar to kubectl, use the following syntax to run commands:  
```
Usage:
  kubeletctl [flags]
  kubeletctl [command]

Available Commands:
  attach        Attach to a container
  configz       Return kubelet's configuration.
  containerLogs Return container log
  cri           Run commands inside a container through the Container Runtime Interface (CRI)
  debug         Return debug information (pprof or flags)
  exec          Run commands inside a container
  healthz       Check the state of the node
  help          Help about any command
  log           Return the log from the node.
  metrics       Return resource usage metrics (such as container CPU, memory usage, etc.)
  pods          Get list of pods on the node
  portForward   Attach to a container
  run           Run commands inside a container
  runningpods   Returns all pods running on kubelet from looking at the container runtime cache.
  scan          Scans for nodes with opened kubelet API
  spec          Cached MachineInfo returned by cadvisor
  stats         Return statistical information for the resources in the node.
  version       Print the version of the kubeletctl

Flags:
      --cacert string      CA certificate (example: /etc/kubernetes/pki/ca.crt )
      --cert string        Private key (example: /var/lib/kubelet/pki/kubelet-client-current.pem)
      --cidr string        A network of IP addresses (Example: x.x.x.x/24)
  -k, --config string      KubeConfig file
  -c, --container string   Container name
  -h, --help               help for kubeletctl
      --http               Use HTTP (default is HTTPS)
  -i, --ignoreconfig       Ignore the default KUBECONFIG environment variable or location ~/.kube
      --key string         Digital certificate (example: /var/lib/kubelet/pki/kubelet-client-current.pem)
  -n, --namespace string   pod namespace
  -p, --pod string         Pod name
      --port string        Kubelet's port, default is 10250
  -r, --raw                Prints raw data
  -s, --server string      Server address (format: x.x.x.x. For Example: 123.123.123.123)
  -u, --uid string         Pod UID

Use "kubeletctl [command] --help" for more information about a command.

```

To view the details on each command or subcommand use the `-h`\\`--help` switch.

## Demo
![kubeletctl](https://github.com/cyberark/kubeletctl/blob/assets/kubeletctl_gif2.gif)



## Build  
Prerequisite:  
-  [go](https://golang.org/doc/install)  
-  [gox](https://github.com/mitchellh/gox)  

To build the project run:  
```
make
```

This will create `build/kubeletctl_{{.OS}}_{{.Arch}}` binaries.  

For Windows users it is possible to use `gox` directly:  
```
gox -ldflags "-s -w" -osarch linux/amd64 -osarch linux/386 -osarch windows/amd64 -osarch windows/386 -osarch="darwin/amd64"
```

## Build with Dockerfile locally
You can use the attached release Dockerfile to build a local image by running:  
```
make docker-release
```

Then run:  
```
docker run -it --rm kubeletctl:release
```

This will fetch and unpack the latest release binary into the Dockerfile.

If you wish to build from source run:
```
make docker
```

Then run:  
```
docker run -it --rm kubeletctl:latest
```

## Contributing

We welcome contributions of all kinds to this repository.  
For instructions on how to get started and descriptions
of our development workflows, please see our [contributing guide](https://github.com/cyberark/conjur-api-go/blob/master/CONTRIBUTING.md).

## License
Copyright (c) 2020 CyberArk Software Ltd. All rights reserved  
This repository is licensed under Apache License 2.0 - see [`LICENSE`](LICENSE) for more details.

## Share Your Thoughts And Feedback
For more comments, suggestions or questions, you can contact Eviatar Gerzi ([@g3rzi](https://twitter.com/g3rzi)) from CyberArk Labs.
You can find more projects developed by us in https://github.com/cyberark/.

[release-img]: https://img.shields.io/github/release/cyberark/kubeletctl.svg
[release]: https://github.com/cyberark/kubeletctl/releases

[license-img]: https://img.shields.io/github/license/cyberark/kubeletctl.svg
[license]: https://github.com/cyberark/kubeletctl/blob/master/LICENSE

[shield-go-version]: https://img.shields.io/github/go-mod/go-version/cyberark/kubeletctl
[go-version]: https://github.com/cyberark/kubeletctl/blob/master/go.mod

[download]: https://img.shields.io/github/downloads/cyberark/kubeletctl/total?logo=github
