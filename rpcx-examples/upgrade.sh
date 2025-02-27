#! /bin/sh

version=v1.8.32

go get -v github.com/by1e11/rpcxio-go@$version

cd registry/dynamic_port_allocation
go get -v github.com/by1e11/rpcxio-go@$version
go get -v github.com/rpcxio/rpcx-etcd@HEAD

cd ../../registry/etcd
go get -v github.com/by1e11/rpcxio-go@$version
go get -v github.com/rpcxio/rpcx-etcd@HEAD

cd ../../registry/etcdv3
go get -v github.com/by1e11/rpcxio-go@$version
go get -v github.com/rpcxio/rpcx-etcd@HEAD

cd ../../registry/etcdv3_lazyregister
go get -v github.com/by1e11/rpcxio-go@$version
go get -v github.com/rpcxio/rpcx-etcd@HEAD

cd ../..
