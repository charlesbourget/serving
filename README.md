# Serving

Simple no dependencies HTTP file server written in Go.

## Build

```
go build .
```

## Options

| Params | Description                             | Default |
|--------|-----------------------------------------|---------|
| -p     | Port to serve on                        | 8100    |
| -d     | The directory of static file to host    | .       |
| -f     | Format to use. Either html, json or xml | html    |
