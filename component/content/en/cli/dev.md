# otter dev

This will run a development server, the development server is crucial for a seemless development due to its hot-reloading.

As we develop UI and iterate countless times per minute, we need hot-reloading to be managed so we can focus on writing the UI and not getting distracted with useless

```
otter dev
```

`otter dev` will read your otter.json, your environment variables and your .env to set up a development server.

The port will be read from a PORT environment variable or alternatively you can use a `--port` flag
```
otter dev --port=8123
```