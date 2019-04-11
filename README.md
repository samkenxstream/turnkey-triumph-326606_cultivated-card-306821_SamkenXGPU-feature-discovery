# NVIDIA GPU feature discovery

## Table of Contents

- [NVIDIA GPU feature discovery](#nvidia-gpu-feature-discovery)
  * [Overview](#overview)
  * [Alpha Version](#alpha-version)
  * [Requirements](#requirements)
  * [Command line interface](#command-line-interface)
  * [Quick Start](#quick-start)
    + [Node Feature Discovery](#node-feature-discovery)
    + [Preparing your GPU Nodes](#preparing-your-gpu-nodes)
    + [Deploy NVIDIA GPU Feature Discovery](#deploy-nvidia-gpu-feature-discovery)
      - [Deamonset](#deamonset)
      - [Job](#job)
  * [Labels](#labels)
  * [Run locally](#run-locally)
  * [Building from source](#building-from-source)

## Overview

NVIDIA GPU Feature Discovery for Kubernetes is a software that allows you to
automatically generate label depending on the GPU available on the node. It uses
the [Node Feature Discovery](https://github.com/kubernetes-sigs/node-feature-discovery)
from Kubernetes to label nodes.

## Alpha Version

**WARNING:** This tool is in alpha version. We may break the API without notice
between releases.

## Requirements

TODO

## Command line interface

Available options:
```
gpu-feature-discovery:
Usage:
  gpu-feature-discovery [--oneshot | --sleep-interval=<seconds>] [--output-file=<file> | -o <file>]
  gpu-feature-discovery -h | --help
  gpu-feature-discovery --version

Options:
  -h --help                       Show this help message and exit
  --version                       Display version and exit
  --oneshot                       Label once and exit
  --sleep-interval=<seconds>      Time to sleep between labeling [Default: 60s]
  -o <file> --output-file=<file>  Path to output file
                                  [Default: /etc/kubernetes/node-feature-discovery/features.d/gfd]
```

You can also use environment variables:

| Env Variable       | Option           | Example |
| ------------------ | ---------------- | ------- |
| GFD_ONESHOT        | --oneshot        | TRUE    |
| GFD_OUTPUT_FILE    | --output-file    | output  |
| GFD_SLEEP_INTERVAL | --sleep-interval | 10s     |

Environment variables override the command line options if they conflict.

## Quick Start

### Node Feature Discovery

The first step is to make sure the [Node Feature Discovery](https://github.com/kubernetes-sigs/node-feature-discovery)
is running on every node you want to label. NVIDIA GPU Feature Discovery use
the `local` source so be sure to mount volumes. See
https://github.com/kubernetes-sigs/node-feature-discovery for more details.

### Preparing your GPU Nodes

Be sure that [nvidia-docker2](https://github.com/NVIDIA/nvidia-docker) is
installed on your GPU nodes and Docker default runtime is set to `nvidia`. See
https://github.com/NVIDIA/nvidia-docker/wiki/Advanced-topics#default-runtime.

### Deploy NVIDIA GPU Feature Discovery

The next step is to run NVIDIA GPU Feature Discovery on each node as a Deamonset
or as a Job.

#### Deamonset

```shell
$ kubectl apply -f gpu-feature-discovery-daemonset.yaml
```

The GPU Feature Discovery should be running on each nodes and generating labels
for the Node Feature Discovery.

#### Job

You must change the `NODE_NAME` value in the template to match the name of the
node you want to label:

```shell
$ export NODE_NAME=<your-node-name>
$ sed "s/NODE_NAME/${NODE_NAME}/" gpu-feature-discovery-job.yaml.template > gpu-feature-discovery-job.yaml
$ kubectl apply -f gpu-feature-discovery-job.yaml
```

The GPU Feature Discovery should be running on the node and generating labels
for the Node Feature Discovery.

## Labels

This is the list of the labels generated by NVIDIA GPU Feature Discovery and
their meaning:

| Label Name            | Value Type | Meaning                           | Example        |
| --------------------- | ---------- | --------------------------------- | -------------- |
| nvidia-driver-version | String     | Version of the NVIDIA driver      | 418.56         |
| nvidia-model          | String     | Model of the GPU                  | GeForce-GT-710 |
| nvidia-memory         | Interger   | Memory of the GPU in Mb           | 2000           |
| nvidia-timestamp      | Interger   | Timestamp of the generated labels | 1555019244     |

## Run locally

Download the source code:
```shell
git clone https://github.com/NVIDIA/gpu-feature-discovery
```

Build the docker image:
```
export GFD_VERSION=$(git describe --tags --dirty --always)
docker build . --build-arg GFD_VERSION=$GFD_VERSION -t gpu-feature-discovery:${GFD_VERSION}
```

Run it:
```
mkdir -p output-dir
docker run -v ${PWD}/output-dir:/etc/kubernetes/node-feature-discovery/features.d gpu-feature-discovery:${GFD_VERSION}
```

You should have set the default runtime of Docker to `nvidia` on your host or
you can also use the `--runtime=nvidia` option:
```
docker run --runtime=nvidia gpu-feature-discovery:${GFD_VERSION}
```

## Building from source

Download the source code:
```shell
git clone https://github.com/NVIDIA/gpu-feature-discovery
```

Get dependies:
```shell
dep ensure
```

Build it:
```
export GFD_VERSION=$(git describe --tags --dirty --always)
go build -ldflags "-X main.Version=${GFD_VERSION}"
```

You can also use the Dockerfile.devel:
```
docker build . -f Dockerfile.devel -t gfd-devel
docker run -it gfd-devel
go build -ldflags "-X main.Version=devel"
```
