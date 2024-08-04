# Embed Svelte SPA into Go binary

This is a learning project.

The primary focus is on building a Svelte SPA (Single-page Application) that can be shipped as Golang binary. And docker container. It includes a CLI interface that will be used to pass in runtime environment variables. 
The server side provides an API and SSE (Server Side Events) to stream data to the frontend.

I took inspiration from projects like Hashicorp Nomad and PocketBase.

As it is a learning project the code may be in various stages of chaos!

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
