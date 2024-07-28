# Embed SvelteKit SPA into Go binary


## Run the server

```sh
make run
```

http://localhost:3000/

## Debug UI with hot-reload

```sh
cd ui
npm run dev
```

http://localhost:5137/


## Adding new routes

When adding new `ui/src/routes/*, add the routes to the go http handler as well. Example:

The following files become routes in SvelteKit

```yaml
ui/src/routes/+page.svelte # localhost:3000/
ui/src/routes/about/+page.svelte # localhost:3000/about
```

But since we embed the built code into the go binary the go server will have to handle the routing like so:

```go
	mux.Handle("/", ui.SvelteKitHandler("/"))
	mux.Handle("/about", ui.SvelteKitHandler("/about"))
```
