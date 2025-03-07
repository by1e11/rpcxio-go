package client

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"sync"
	"testing"
	"time"

	testutils "github.com/by1e11/rpcxio-go/_testutils"
	"github.com/by1e11/rpcxio-go/protocol"
	"github.com/by1e11/rpcxio-go/server"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

type Arith int

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

type PBArith int

func (t *PBArith) Mul(ctx context.Context, args *testutils.ProtoArgs, reply *testutils.ProtoReply) error {
	reply.C = args.A * args.B
	return nil
}

func (t *Arith) ThriftMul(ctx context.Context, args *testutils.ThriftArgs_, reply *testutils.ThriftReply) error {
	reply.C = args.A * args.B
	return nil
}

type Bidirectional struct {
	*server.Server
}

func (t *Bidirectional) Mul(ctx context.Context, args *Args, reply *Reply) error {
	conn := ctx.Value(server.RemoteConnContextKey).(net.Conn)
	reply.C = args.A * args.B
	t.SendMessage(conn, "test_service_path", "test_service_method", nil, []byte("abcde"))
	return nil
}

func TestClient_IT(t *testing.T) {
	s := server.NewServer()
	_ = s.RegisterName("Arith", new(Arith), "")
	_ = s.RegisterName("PBArith", new(PBArith), "")
	go func() {
		_ = s.Serve("tcp", "127.0.0.1:0")
	}()
	defer s.Close()
	time.Sleep(500 * time.Millisecond)

	addr := s.Address().String()

	client := &Client{
		option: DefaultOption,
	}

	err := client.Connect("tcp", addr)
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer client.Close()

	args := &Args{
		A: 10,
		B: 20,
	}

	reply := &Reply{}
	err = client.Call(context.Background(), "Arith", "Mul", args, reply)
	if err != nil {
		t.Fatalf("failed to call: %v", err)
	}

	if reply.C != 200 {
		t.Fatalf("expect 200 but got %d", reply.C)
	}

	err = client.Call(context.Background(), "Arith", "Add", args, reply)
	if err == nil {
		t.Fatal("expect an error but got nil")
	}

	client.option.SerializeType = protocol.MsgPack
	reply = &Reply{}
	err = client.Call(context.Background(), "Arith", "Mul", args, reply)
	if err != nil {
		t.Fatalf("failed to call: %v", err)
	}

	if reply.C != 200 {
		t.Fatalf("expect 200 but got %d", reply.C)
	}

	client.option.SerializeType = protocol.ProtoBuffer

	pbArgs := &testutils.ProtoArgs{
		A: 10,
		B: 20,
	}
	pbReply := &testutils.ProtoReply{}
	err = client.Call(context.Background(), "PBArith", "Mul", pbArgs, pbReply)
	if err != nil {
		t.Fatalf("failed to call: %v", err)
	}

	if pbReply.C != 200 {
		t.Fatalf("expect 200 but got %d", pbReply.C)
	}
}

func TestClient_IT_Concurrency(t *testing.T) {
	s := server.NewServer()
	_ = s.RegisterName("PBArith", new(PBArith), "")
	go func() {
		_ = s.Serve("tcp", "127.0.0.1:0")
	}()
	defer s.Close()
	time.Sleep(500 * time.Millisecond)

	addr := s.Address().String()

	client := &Client{
		option: DefaultOption,
	}

	err := client.Connect("tcp", addr)
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer client.Close()

	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		i := i
		go testSendRaw(t, client, uint64(i), rand.Int31(), rand.Int31(), &wg)
	}
	wg.Wait()

}

func testSendRaw(t *testing.T, client *Client, seq uint64, x, y int32, wg *sync.WaitGroup) {
	defer wg.Done()
	rpcxReq := protocol.NewMessage()
	rpcxReq.SetMessageType(protocol.Request)
	rpcxReq.SetSeq(seq)
	rpcxReq.ServicePath = "PBArith"
	rpcxReq.ServiceMethod = "Mul"
	rpcxReq.SetSerializeType(protocol.ProtoBuffer)
	rpcxReq.SetOneway(false)

	pbArgs := &testutils.ProtoArgs{
		A: x,
		B: y,
	}
	data, _ := pbArgs.Marshal()
	rpcxReq.Payload = data
	_, reply, err := client.SendRaw(context.Background(), rpcxReq)
	if err != nil {
		t.Errorf("failed to call SendRaw: %v", err)
		return
	}

	pbReply := &testutils.ProtoReply{}
	err = pbReply.Unmarshal(reply)
	if err != nil {
		t.Errorf("failed to unmarshal reply: %v", err)
		return
	}

	if pbReply.C != x*y {
		t.Errorf("expect %d but got %d", x*y, pbReply.C)
		return
	}
}

func TestClient_Res_Reset(t *testing.T) {
	var res = protocol.NewMessage()
	res.Payload = []byte{1, 2, 3, 4, 5, 6, 7, 8}
	data := res.Payload
	res.Reset()

	if len(data) == 0 {
		t.Fatalf("data has been set to empty after response has been reset: %v", data)
	}
}

func TestClient_Bidirectional(t *testing.T) {
	s := server.NewServer()
	_ = s.RegisterName("Bidirectional", &Bidirectional{Server: s}, "")
	go func() {
		_ = s.Serve("tcp", "127.0.0.1:0")
	}()
	defer s.Close()
	time.Sleep(500 * time.Millisecond)

	addr := s.Address().String()

	opt := DefaultOption

	var receive string

	opt.NilCallServerMessageHandler = func(msg *protocol.Message) {
		fmt.Printf("receive msg from server: %s\n", msg.Payload)
		receive = string(msg.Payload)
	}
	client := &Client{
		option: opt,
	}

	err := client.Connect("tcp", addr)
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer client.Close()

	args := &Args{
		A: 10,
		B: 20,
	}
	reply := &Reply{}
	err = client.Call(context.Background(), "Bidirectional", "Mul", args, reply)
	if err != nil {
		t.Fatalf("failed to call: %v", err)
	}
	if receive != "abcde" {
		t.Fatalf("expect abcde but got %s", receive)
	}
	if reply.C != 200 {
		t.Fatalf("expect 200 but got %d", reply.C)
	}

}
