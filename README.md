# UI Skeleton written in golang

based on the REST API Skeleton

## Installation
```
go get github.com/cseeger-epages/ui-skeleton
```
or simply clone the repo
```
git clone https://github.com/cseeger-epages/ui-skeleton
```

## Configuration
The configuration file can be found under conf/api.conf.
See the config file for more details
A default template.conf can be found under conf/template.conf.

## build and run
you can build the binary via
```
make build
```
and the following flags are supported
```
-c    <config file>
```

## Further Implementation
By default an IndexHandler is used for serving Files and StaticHandler for static files.

Nevermind you can add custom Handlers to src/Handler.go and add them to src/Routes.go with the following pattern
```
Route{
  "<some route name>",
  "<request method>",
  "/yourcustomroute",
  "<description used for /help/[cmd]"
  "<handler name>"
}
```
## Supported Features
- path routing using gorilla mux
- Database wrapper 
- TLS
- HSTS
- HTML Templating
- Serving static files
- pretty print
- Etag / If-None-Match Clientside caching
- Rate limiting and headers using trottled middleware
- basic auth
- config using TOML format
- error handler
- logging

## Not (yet) implemented
- X-HTTP-Method-Override
- caching serverside (varnish ?)
- Authentication - oauth(2) 

## Ratelimit Headers
```
X-Ratelimit-Limit - The number of allowed requests in the current period
X-Ratelimit-Remaining - The number of remaining requests in the current period
X-Ratelimit-Reset - The number of seconds left in the current period
```

## generate certificates
```
cd certs
# Key considerations for algorithm "RSA" ≥ 2048-bit
openssl genrsa -out server.key 2048

# Key considerations for algorithm "ECDSA" ≥ secp384r1
# List ECDSA the supported curves (openssl ecparam -list_curves)
openssl ecparam -genkey -name secp384r1 -out server.key
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

## some curls
```
curl -k -X GET -H "Authorization: Basic dGVzdHVzZXI6dGVzdHBhc3MK" https://localhost:8443/
curl -k -X GET -H "Authorization: Basic dGVzdHVzZXI6dGVzdHBhc3MK" https://localhost:8443/index
```

## basic auth test stuff
```
testuser:testpass - dGVzdHVzZXI6dGVzdHBhc3MK
username:password - dXNlcm5hbWU6cGFzc3dvcmQK
```
