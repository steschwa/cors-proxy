# cors-proxy

Simple CORS proxy written in Go

## Build

**Standalone**

`GOPROXY_PORT=5000 go run proxy.go`

`GOPROXY_PORT=` can be any free port

**Docker container**

`docker build --rm -f proxy.Dockerfile -t corsproxy:latest .`

`docker run -e CORSPROXY_PORT=5000 -d --name corsproxy corsproxy`

## Usage

Currently only _GET_ request are supported.
Send requests to `http://localhost:5000?url={PROXIED_URL}` and get back the proxied url response.
