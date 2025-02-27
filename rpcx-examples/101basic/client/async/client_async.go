package main

import (
	"context"
	"flag"
	"log"

	"github.com/by1e11/rpcxio-go/client"
	logging "github.com/by1e11/rpcxio-go/log"
	example "github.com/rpcxio/rpcx-examples"
)

type FibbArgs struct {
	N int
}

var (
	addr2 = flag.String("addr", "localhost:33632", "server address")
)

func main() {
	flag.Parse()

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr2, "")
	xclient := client.NewXClient("TestServiceStream", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := FibbArgs{
		N: 10,
	}

	reply := &example.Reply{}
	call, err := xclient.Go(context.Background(), "fibbonacci.Stream", args, reply, nil)

	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	if call.Stream != nil {
		for {
			select {
			case <-call.Done:
				log.Println("call is done")
				close(call.Stream)
				return
			case chunk := <-call.Stream:
				log.Printf("chunk: %v", chunk.(map[string]any)["C"])
			default:
				logging.Debug("waiting for chunk...")
			}
		}
	} else {
		replyCall := <-call.Done
		if replyCall.Error != nil {
			logging.Warnf("failed to call: %v", replyCall.Error)
		}
		logging.Info("reply: %v", reply.C)
	}

}
