package main

import (
	"context"
	"flag"
	"log"

	"github.com/by1e11/rpcxio-go/client"
	logging "github.com/by1e11/rpcxio-go/log"
)

type Args struct {
	N int
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
	xclient := client.NewXClient("TestServiceStream", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := Args{
		N: 10,
	}

	reply := &Reply{}
	call, err := xclient.Go(context.Background(), "fibbonacci.Stream", args, reply, nil)

	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	for {
		select {
		case <-call.Done:
			log.Println("call is done")
			close(call.Stream)
			return
		case chunk := <-call.Stream:
			log.Printf("chunk: %v", chunk.(*Reply).C)
		default:
			logging.Debug("waiting for chunk...")
		}
	}
}
