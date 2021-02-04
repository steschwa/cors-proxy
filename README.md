# cors-proxy

Simple CORS proxy written in Go

## Build

**Standalone**

`CORSPROXY_PORT=5000 go run proxy.go`

> Note: `CORSPROXY_PORT=` can be any free port

**Docker container**

`docker build --rm -f proxy.Dockerfile -t corsproxy:latest .`

Start a container and map the proxy to _http://localhost:5005_:

`docker run --rm -e CORSPROXY_PORT=5000 -p 5005:5000 --name corsproxy corsproxy`

## Usage

Currently only _GET_ request are supported.
Send requests to `GET http://localhost:5000?url={PROXIED_URL}` and get back the proxied url response.

> Make sure to URL encode the query string parameter

## Example

```ts
const PROXY_URL = "https://localhost:5000"

function withProxy(url: string) {
    return fetch(`${PROXY_URL}?url=${encodeURI(url)}`)
}
```
