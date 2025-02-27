package main

import (
	"context"
	"flag"
	"log"

	"github.com/by1e11/rpcxio-go/client"
	logging "github.com/by1e11/rpcxio-go/log"
)

type Args struct {
	A, B int
}

type Reply struct {
	C int
}

var (
	addr2 = flag.String("addr", "localhost:33632", "server address")
)

func main() {
	flag.Parse()

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr2, "")
	xclient := client.NewXClient("TestService", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := Args{
		A: 10,
		B: 20,
	}

	reply := &Reply{}
	// call, err := xclient.Go(context.Background(), "fibbonacci.Stream", args, reply, nil)
	call, err := xclient.Go(context.Background(), "mul", args, reply, nil)

	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	replyCall := <-call.Done
	if replyCall.Error != nil {
		logging.Warnf("failed to call: %v", replyCall.Error)
	}
	log.Printf("reply: %v", reply.C)
}
