# Embed Svelte SPA into Go binary

This is a learning project.

The primary focus is on building a Svelte SPA (Single-page Application) that can be shipped as Golang binary. And docker container. It includes a CLI interface that will be used to pass in runtime environment variables.
The server side provides an API and SSE (Server Side Events) to stream data to the frontend.

I took inspiration from projects like Hashicorp Nomad and PocketBase.

As it is a learning project the code may be in various stages of chaos!

## Overview

After building the binary (see below) it can be run with the following commands in two separate terminals

Start the server

```sh
./bin/cli server
```

Run the CLI command `hello` which calls the API

```sh
./bin/cli hello
```

Optionally the server port and client endpoint can be changed with either flags or environment variables

```sh
export THING_SERVER_PORT=3001
export THING_ENDPOINT="http://localhost:3001"
```

## Build binary

```sh
make build
```

## Run the server

```sh
make server
```

<http://localhost:3000/>

## Debug UI with hot-reload

```sh
make dev
```

<http://localhost:5137/>

## Build

```
make build
./bin/cli server
```

<http://localhost:3000/>
