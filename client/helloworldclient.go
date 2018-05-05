package main

import (
	"flag"
	"fmt"
	"time"

	grpclb "grpclb/etcdv3"
	pb "grpclb/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"strconv"
)

var (
	serv = flag.String("service", "hello_service", "service name")
	//etcd的注册地址
	reg  = flag.String("reg", "http://127.0.0.1:2379", "register etcd address")
)

func main() {
	flag.Parse()
	//根据服务名字获取已经注册的服务
	r := grpclb.NewResolver(*serv)
	b := grpc.RoundRobin(r)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	conn, err := grpc.DialContext(ctx, *reg, grpc.WithInsecure(), grpc.WithBalancer(b))
	if err != nil {
		panic(err)
	}

	fmt.Println("client start...")

	ticker := time.NewTicker(1 * time.Second)
	for t := range ticker.C {
		client := pb.NewGreeterClient(conn)
		resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "world " + strconv.Itoa(t.Second())})
		if err == nil {
			fmt.Printf("%v: Reply is %s\n", t, resp.Message)
		} else {
			fmt.Printf("error: %v", err)
		}
	}
}
