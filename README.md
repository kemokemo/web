# web

A simplest web service to serve static files. :tada:

## Usage

### build manually

```sh
## build and execute.
go build && ./web

## open 'http://localhost:8080' via browser
```

### docker run

This tool serves the `/app/public` directory in the container. Please bind your `public` directory.

```sh
docker run --mount type=bind,source="$(pwd)"/public,target=/app/public,readonly -d -p 8080:8080 kemokemo/web:v0.0.1
```
