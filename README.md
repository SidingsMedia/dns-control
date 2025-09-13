<!-- 
SPDX-FileCopyrightText: 2022-2023 Sidings Media <contact@sidingsmedia.com>
SPDX-License-Identifier: MIT
-->

# <Service>

This repo contains the source for the <service> service, part of Sidings
Media's public API

## Building

### Binary

This project is written in go so you will need this to be installed.

First download the project dependancies.

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
docker build . -t <service>:latest
```

Note: `-t <service>` gives the container the name <service> and the tag
latest.

Docker will now download all the dependancies and then build your
container. This may take a while.

## Running

### Environment variables

This service requires certain environment variables in order to function
correctly. An example `.env` file can be found in the document root
(`.env.example`). Below is a complete table of all environment
variables.

| Name              | Required           | Description                                                                                                         | Example                                       |
|-------------------|--------------------|---------------------------------------------------------------------------------------------------------------------|-----------------------------------------------|
| `BIND_ADDR`       | :x:                | This is the address to bind the server to. Defaults to `[::1]:3000`.                                                | `[::]:3000`                                   |
| `TRUSTED_PROXIES` | :x:                | Proxy servers to trust when reading client IP headers. Provide addresses in a comma separated list.Defaults to `*`. | `192.0.2.1,192.0.2.2,2001:db8::1,2001:db8::2` |
| `GIN_MODE`        | :x:                | Mode to run Gin in. Only set to `debug` for development. Defaults to `release`.                                     | `release`                                     |

### Binary

If you are using the binary to run the service, you have two options for
setting the environment variables. One is to actually set them on the
system, the other option is to store the settings in a .env file which
will be automatically loaded on start.

### Docker

```
docker run --publish 3000:3000 -d --name messaging ghcr.io/sidingsmedia/<service>
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
[contact@sidingsmedia.com](mailto:contact@sidingsmedia.com?subject=Messaging%20Microservice).
All files committed to this repo must contain valid licencing
information or the pull request can not be accepted.
