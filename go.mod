module github.com/johnbellone/cache-service

go 1.14

require (
	github.com/coreos/bbolt v1.3.4 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.14.5 // indirect
	github.com/micro/go-micro/v2 v2.6.0
	github.com/pomerium/autocache v0.0.0-20200505053831-8c1cd659f055 // indirect
	github.com/ugorji/go v1.1.7 // indirect
	go.etcd.io/etcd v3.3.20+incompatible // indirect
	golang.org/x/net v0.0.0-20200520004742-59133d7f0dd7 // indirect
	golang.org/x/sys v0.0.0-20200519105757-fe76b779f299 // indirect
	google.golang.org/api v0.24.0
	google.golang.org/genproto v0.0.0-20200511104702-f5ebc3bea380 // indirect
	google.golang.org/grpc v1.29.1 // indirect
	github.com/golang/protobuf v1.4.1
	google.golang.org/protobuf v1.22.0
	gopkg.in/yaml.v2 v2.2.8 // indirect
	github.com/sirupsen/logrus v1.6.0
)

replace (
	github.com/coreos/bbolt v1.3.4 => go.etcd.io/bbolt v1.3.4
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
	go.etcd.io/etcd => go.etcd.io/etcd v0.5.0-alpha.5.0.20200425165423-262c93980547
)
