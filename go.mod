module github.com/johnbellone/cache-service

go 1.14

require (
	github.com/coreos/bbolt v1.3.4 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.14.4 // indirect
	github.com/micro/go-micro/v2 v2.6.0
	github.com/pomerium/autocache v0.0.0-20200505053831-8c1cd659f055 // indirect
	github.com/ugorji/go v1.1.7 // indirect
	go.etcd.io/etcd v3.3.21+incompatible // indirect
	golang.org/x/net v0.0.0-20200520004742-59133d7f0dd7 // indirect
	golang.org/x/sys v0.0.0-20200519105757-fe76b779f299 // indirect
	google.golang.org/genproto v0.0.0-20200507105951-43844f6eee31 // indirect
	google.golang.org/grpc v1.29.1 // indirect
	google.golang.org/protobuf v1.22.0
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace github.com/coreos/etcd => go.etcd.io/etcd v0.5.0-alpha.5.0.20200306183522-221f0cc107cb

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.3
	github.com/coreos/bbolt v1.3.4 => go.etcd.io/bbolt v1.3.4
)

replace google.golang.org/grpc v1.29.1 => google.golang.org/grpc v1.26.0
