# Small go app

## Build 

``` sh
podman build -t go-hello-world:1.0.0
```

## Run

``` sh
podman run --rm -p 8080:8080 \
    -e APP_VERSION=1.0.0 \
    -e APP_ENVIRONMENT=prod go-hello-world:1.0.0
```
