<!--
SPDX-FileCopyrightText: 2022-2025 Sidings Media <contact@sidingsmedia.com>
SPDX-License-Identifier: MIT
-->

# DNS Control

This repo contains the source for the dns-control service, part of Sidings
Media's unified control API

## Building

### Binary

This project is written in go so you will need this to be installed.

First download the project dependencies.

```
go mod download
```

And then you can compile the binary.

```
go build -a -o server
```

### Docker

This will require docker to be installed. After you have installed
docker, you need to run only one command to build the container.

```
docker build . -t dns-control:latest
```

Note: `-t dns-control` gives the container the name dns-control and the tag
latest.

Docker will now download all the dependencies and then build your
container. This may take a while.

## Running

### Configuration File

Configuration is provided through a `yaml` file. An example can be found
at [config-example.yaml](/config-example.yaml). By default, the service
will attempt to load `config.yaml` from it's current working directory.
You may use the `-config` flag to specify another path.

### Binary

If you are using the binary to run the service, you have two options for
setting the environment variables. One is to actually set them on the
system, the other option is to store the settings in a .env file which
will be automatically loaded on start.

### Docker

```
docker run --publish 3000:3000 -d --name dns-control ghcr.io/sidingsmedia/dns-control
```

To add the environment variables, you can use multiple `-e` flags. For
more information see the [docker
documentation](https://docs.docker.com/engine/reference/commandline/run/#env).

### Docker Compose

A docker compose file is also provided if you would like to use it.

```
docker compose up . -d
```

To pass the environment variables, just store them in a .env file.

## Licence

This repo uses the [REUSE](https://reuse.software) standard in order to
communicate the correct licence for the file. For those unfamiliar with
the standard the licence for each file can be found in one of three
places. The licence will either be in a comment block at the top of the
file, in a `.license` file with the same name as the file, or in the
dep5 file located in the `.reuse` directory. If you are unsure of the
licencing terms please contact
[legal@sidingsmedia.com](mailto:legal@sidingsmedia.com?subject=Licensing%3A%20DNS%20Control%20Microservice).
All files committed to this repo must contain valid licencing
information or the pull request can not be accepted.
