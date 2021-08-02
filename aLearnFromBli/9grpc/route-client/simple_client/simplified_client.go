package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "minggrpctest/route"
)

// 忽略证书验证
// 拨号成功，不要往下走

func runFirst(client pb.RouteGuideClient){
	feature, err := client.GetFeature(context.Background(), &pb.Point{
		Latitude:  310235000,
		Longitude: 121437403,
	})

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(feature)

}

func main(){
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil{
		log.Fatalln("client cannot dial grpc server")

	}

	defer conn.Close()

	client := pb.NewRouteGuideClient(conn)
	runFirst(client)
}
