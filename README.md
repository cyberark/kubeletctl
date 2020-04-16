[![License][license-img]][license]

## Overview
Kubeletctl is a command line tool that implement kubelet's API.  
Part of kubelet's API is documented but most of it is not.  
This tool covers all the documented and undocumented APIs.

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
  debug         Return pprof information
  exec          Run commands inside a container
  help          Help about any command
  log           Return the log from the node.
  metrics       Return information about node CPU and memory usage
  pods          Get list of pods on the node
  portForward   Attach to a container
  resource      Return information about node resources
  run           Run commands inside a container
  runningpods   Returns all pods running on kubelet from looking at the container runtime cache.
  spec          Cached MachineInfo returned by cadvisor
  stats         Return performance stats of node, pods and containers.

Flags:
  -c, --container string   container
  -h, --help               help for kubeletctl
  -n, --namespace string   pod namespace
  -p, --pod string         container
      --port string        Kubelet's port, default is 10250
  -s, --server string      Server address (format: <server_IP>)
  -u, --uid string         container

Use "kubeletctl [command] --help" for more information about a command.
```

To view the details on each command or subcommand use the `-h`\\`--help` switch.

## Demo
![kubeletctl](https://github.com/cyberark/kubeletctl/blob/assets/kubeletctl_gif2.gif)

## Contributing

We welcome contributions of all kinds to this repository.  
For instructions on how to get started and descriptions
of our development workflows, please see our [contributing guide](https://github.com/cyberark/conjur-api-go/blob/master/CONTRIBUTING.md).

## License

This repository is licensed under Apache License 2.0 - see [`LICENSE`](LICENSE) for more details.

## Share Your Thoughts And Feedback
For more comments, suggestions or questions, you can contact Eviatar Gerzi ([@g3rzi](https://twitter.com/g3rzi)) from CyberArk Labs.
You can find more projects developed by us in https://github.com/cyberark/.

[license-img]: https://img.shields.io/github/license/cyberark/kubeletctl.svg
[license]: https://github.com/cyberark/kubeletctl/blob/master/LICENSE
