# Serving

Simple no dependencies HTTP file server written in Go.

## Build

```
go build .
```

## Options

| Params | Description                                               |
|--------|-----------------------------------------------------------|
| -p     | Port to serve on (Default: 8100)                          |
| -d     | The directory of static file to host (Default: .)         |
| -f     | Format to use. Either plain, json or xml (Default: plain) |
