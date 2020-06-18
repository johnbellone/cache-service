# Cache Service

An example Go [Micro][0] service to handle basic caching operations with
[gRPC][1] transport and [Protobuf][2] as an [IDL][3].

This is **not** a production service but is meant to serve as an example
on how to build scalable microservices with Go.

## Development

``` bash
~/src/cache-service % make
~/src/cache-service % bin/cache-service &
~/src/cache-service % micro api &
~/src/cache-service % curl http://localhost:8080/cache/get
2020-06-17 23:23:04  file=cache/handler.go:29 level=info Received Cache.Get request
::1 - - [17/Jun/2020:23:23:04 -0400] "GET /cache/get HTTP/1.1" 200 2 "" "curl/7.64.1"
{}%
```

[0]: https://github.com/micro/go-micro
[1]: https://micro.mu/docs/go-grpc.html
[2]: https://developers.google.com/protocol-buffers/docs/proto3
[3]: https://en.wikipedia.org/wiki/IDL_(programming_language)
[5]: https://devcenter.heroku.com/articles/getting-started-with-go
