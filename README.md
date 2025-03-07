- **stable branch**: v1.7.x
- **development branch**: master

<a href="https://rpcx.io/"><img height="160" src="http://rpcx.io/logos/rpcx-logo-text.png"></a>

Official site: [http://rpcx.io](http://rpcx.io/)

[![License](https://img.shields.io/:license-apache%202-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![GoDoc](https://godoc.org/github.com/by1e11/rpcxio-go?status.png)](http://godoc.org/github.com/by1e11/rpcxio-go)  [![github actions](https://github.com/by1e11/rpcxio-go/actions)](https://github.com/by1e11/rpcxio-go/actions/workflows/go.yml/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/by1e11/rpcxio-go)](https://goreportcard.com/report/github.com/by1e11/rpcxio-go) [![coveralls](https://coveralls.io/repos/smallnest/rpcx/badge.svg?branch=master&service=github)](https://coveralls.io/github/smallnest/rpcx?branch=master) [![QQ3群](https://img.shields.io/:QQ3群-953962236-blue.svg)](_documents/rpcx_dev_qq3.jpg) 

**Notice: etcd**

since rpcx 1.7.6, some plugins have been moved to the independent project:

- `etcd` plugin has been moved to [rpcx-etcd](https://github.com/rpcxio/rpcx-etcd)
- `zookeeper` plugin has been moved to [rpcx-zookeeper](https://github.com/rpcxio/rpcx-zookeeper)
- `consul` plugin has been moved to [rpcx-consul](https://github.com/rpcxio/rpcx-consul)
- `redis` plugin has been moved to [rpcx-redis](https://github.com/rpcxio/rpcx-redis)
- `influxdb` plugin has been moved to [rpcx-plugins](https://github.com/rpcxio/rpcx-plugins)
- `opentelemetry` plugin has been moved to [rpcx-plugins](https://github.com/rpcxio/rpcx-plugins)

## Announce

A tcpdump-like tool added: [rpcxdump](https://github.com/by1e11/rpcxio-godump)。 You can use it to debug communications between rpcx services and clients.

![](https://github.com/by1e11/rpcxio-godump/blob/master/snapshoot.png)


## Cross-Languages
you can use other programming languages besides Go to access rpcx services.

- **rpcx-gateway**: You can write clients in any programming languages to call rpcx services via [rpcx-gateway](https://github.com/rpcxio/rpcx-gateway)
- **http invoke**: you can use the same http requests to access rpcx gateway
- **Java Services/Clients**: You can use [rpcx-java](https://github.com/by1e11/rpcxio-go-java) to implement/access rpcx services via raw protocol.
- **rust rpcx**: You can write rpcx services in rust by [rpcx-rs](https://github.com/by1e11/rpcxio-go-rs)

> If you can write Go methods, you can also write rpc services. It is so easy to write rpc applications with rpcx.

## Installation

install the basic features:

`go get -v github.com/by1e11/rpcxio-go/...`


If you want to use `quic`、`kcp` registry, use those tags to `go get` 、 `go build` or `go run`. For example, if you want to use all features, you can:

```sh
go get -v -tags "quic kcp" github.com/by1e11/rpcxio-go/...
```

**_tags_**:
- **quic**: support quic transport
- **kcp**: support kcp transport

## Which companies are using rpcx?

<p float="left">
  <img style="padding-bottom: 20px;" src="https://user-images.githubusercontent.com/865763/102993220-b967f000-4557-11eb-9747-703a6cbb9fb1.png" width="200" />
  <img style="padding-bottom: 20px;" src="https://user-images.githubusercontent.com/865763/113414067-a8e4d280-93ee-11eb-9f42-1373d7e766c1.png" width="200" />
  <img style="padding-bottom: 20px;" src="https://user-images.githubusercontent.com/865763/102993433-267b8580-4558-11eb-9e45-4e1a86d61688.png" width="200" /> 
  <img style="padding-bottom: 20px;" src="https://user-images.githubusercontent.com/865763/102993530-4743db00-4558-11eb-9f76-1ee69e992b82.png" width="200" />
  <img style="padding-bottom: 20px;" src="https://user-images.githubusercontent.com/865763/102993612-722e2f00-4558-11eb-849a-3264c430aef9.png" width="200" />
  <img style="padding-bottom: 20px;" src="https://user-images.githubusercontent.com/865763/102993785-c20cf600-4558-11eb-82b9-27b801aca4ff.png" width="200" />
</p>

## Features
rpcx is a RPC framework like [Alibaba Dubbo](http://dubbo.io/) and [Weibo Motan](https://github.com/weibocom/motan).

**rpcx** is created for targets:
1. **Simple**: easy to learn, easy to develop, easy to integrate and easy to deploy
2. **Performance**: high performance (>= grpc-go)
3. **Cross-platform**: support _raw slice of bytes_, _JSON_, _Protobuf_ and _MessagePack_. Theoretically it can be used with java, php, python, c/c++, node.js, c# and other platforms
4. **Service discovery and service governance**: support zookeeper, etcd and consul.


It contains below features
- Support raw Go functions. There's no need to define proto files.
- Pluggable. Features can be extended such as service discovery, tracing.
- Support TCP, HTTP, [QUIC](https://en.wikipedia.org/wiki/QUIC) and [KCP](https://github.com/skywind3000/kcp)
- Support multiple codecs such as JSON, Protobuf, [MessagePack](https://msgpack.org/index.html) and raw bytes.
- Service discovery. Support peer2peer, configured peers, [zookeeper](https://zookeeper.apache.org), [etcd](https://github.com/coreos/etcd), [consul](https://www.consul.io) and [mDNS](https://en.wikipedia.org/wiki/Multicast_DNS).
- Fault tolerance：Failover, Failfast, Failtry.
- Load banlancing：support Random, RoundRobin, Consistent hashing, Weighted, network quality and Geography.
- Support Compression.
- Support passing metadata.
- Support Authorization.
- Support heartbeat and one-way request.
- Other features: metrics, log, timeout, alias, circuit breaker.
- Support bidirectional communication.
- Support access via HTTP so you can write clients in any programming languages.
- Support API gateway.
- Support backup request, forking and broadcast.


rpcx uses a binary protocol and platform-independent, which means you can develop services in other languages such as Java, python, nodejs, and you can use other prorgramming languages to invoke services developed in Go.

There is a UI manager: [rpcx-ui](https://github.com/by1e11/rpcxio-go-ui).

## Performance

Test results show rpcx has better performance than other rpc framework except standard rpc lib.


The benchmark code is at [rpcx-benchmark](https://github.com/rpcx-ecosystem/rpcx-benchmark).

**Listen to others, but test by yourself**.

**_Test Environment_**

- **CPU**: Intel(R) Xeon(R) CPU E5-2630 v3 @ 2.40GHz, 32 cores
- **Memory**: 32G
- **Go**: 1.9.0
- **OS**: CentOS 7 / 3.10.0-229.el7.x86_64

**_Use_**
- protobuf
- the client and the server on the same server
- 581 bytes payload
- 500/2000/5000 concurrent clients
- mock processing time: 0ms, 10ms and 30ms

**_Test Result_**

### mock 0ms process time

<table><tr><th>Throughputs</th><th>Mean Latency</th><th>P99 Latency</th></tr><tr><td width="30%"><img src="http://colobu.com/2018/01/31/benchmark-2018-spring-of-popular-rpc-frameworks/p0-throughput.png"></td><td width="30%"><img src="http://colobu.com/2018/01/31/benchmark-2018-spring-of-popular-rpc-frameworks/p0-latency.png"></td><td width="30%"><img src="http://colobu.com/2018/01/31/benchmark-2018-spring-of-popular-rpc-frameworks/p0-p99.png"></td></tr></table>


### mock 10ms process time

<table><tr><th>Throughputs</th><th>Mean Latency</th><th>P99 Latency</th></tr><tr><td width="30%"><img src="http://colobu.com/2018/01/31/benchmark-2018-spring-of-popular-rpc-frameworks/p10-throughput.png"></td><td width="30%"><img src="http://colobu.com/2018/01/31/benchmark-2018-spring-of-popular-rpc-frameworks/p10-latency.png"></td><td width="30%"><img src="http://colobu.com/2018/01/31/benchmark-2018-spring-of-popular-rpc-frameworks/p10-p99.png"></td></tr></table>


### mock 30ms process time

<table><tr><th>Throughputs</th><th>Mean Latency</th><th>P99 Latency</th></tr><tr><td width="30%"><img src="http://colobu.com/2018/01/31/benchmark-2018-spring-of-popular-rpc-frameworks/p30-throughput.png"></td><td width="30%"><img src="http://colobu.com/2018/01/31/benchmark-2018-spring-of-popular-rpc-frameworks/p30-latency.png"></td><td width="30%"><img src="http://colobu.com/2018/01/31/benchmark-2018-spring-of-popular-rpc-frameworks/p30-p99.png"></td></tr></table>


## Examples

You can find all examples at [rpcxio/rpcx-examples](https://github.com/rpcxio/rpcx-examples).

The below is a simple example.


**Server**

```go
    // define example.Arith
    ……

    s := server.NewServer()
	s.RegisterName("Arith", new(example.Arith), "")
	s.Serve("tcp", addr)

```


**Client**

```go
    // prepare requests
    ……

    d, err := client.NewPeer2PeerDiscovery("tcp@"+addr, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()
	err = xclient.Call(context.Background(), "Mul", args, reply, nil)
```

## Contributors

<a href="https://github.com/by1e11/rpcxio-go/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=smallnest/rpcx" />
</a>


## Contribute

see [contributors](https://github.com/by1e11/rpcxio-go/graphs/contributors).

Welcome to contribute:
- submit issues or requirements
- send PRs
- write projects to use rpcx
- write tutorials or articles to introduce rpcx

## License

Apache License, Version 2.0 
