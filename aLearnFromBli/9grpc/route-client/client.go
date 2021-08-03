package main
import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	pb "minggrpctest/route"
)
func runFirst(client pb.RouteGuideClient){
//func runFirst(client *pb.RouteGuideClient) { RouteGuideClient 和 *RouteGuideClient  虽都是同型接口，但是不同类型，一个是地址
	feature, err := client.GetFeature(context.Background(), &pb.Point{
		Latitude:  310235000,
		Longitude: 121437403,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(feature)
}

func runSecond(client pb.RouteGuideClient) {
	serverStream, err := client.ListFeatures(context.Background(), &pb.Rectangle{
		Lo: &pb.Point{Latitude: 313374060, Longitude: 121358540},
		Hi: &pb.Point{Latitude: 311034130, Longitude: 121598790},
	})
	if err != nil {
		log.Fatalln(err)
	}

	for {
		feature, err := serverStream.Recv()
		//流关闭
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(feature)
	}
}

func runThird(client pb.RouteGuideClient) {
	// dummy data
	points := []*pb.Point{
		{Latitude: 313374060, Longitude: 121358540},
		{Latitude: 311034130, Longitude: 121598790},
		{Latitude: 310235000, Longitude: 121437403},
	}

	clientStream, err := client.RecordRoute(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	for _, point := range points {
		if err := clientStream.Send(point); err != nil {
			log.Fatalln(err)
		}
		time.Sleep(time.Second)
	}
	summary, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	//if summary, err := clientStream.CloseAndRecv(); err != nil {}
	//这种写法，上面4行化作一行，summary作用域会限定在if里，后面会调用不到的
	fmt.Println(summary)
}

func readIntFromCommandLine(reader *bufio.Reader, target *int32) {
	_, err := fmt.Fscanf(reader, "%d\n", target)
	if err != nil {
		log.Fatalln("Cannot scan", err)
	}
}

func runForth(client pb.RouteGuideClient) {
	stream, err := client.Recommend(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	// this goroutine listen to the server stream
	go func() {
		feature, err2 := stream.Recv()
		if err2 != nil {
			log.Fatalln(err2)
		}
		fmt.Println("Recommended: ", feature)
	}()

	reader := bufio.NewReader(os.Stdin)

	for {
		request := pb.RecommendationRequest{Point: new(pb.Point)}
		var mode int32
		fmt.Print("Enter Recommendation Mode (0 for farthest, 1 for the nearest)")
		readIntFromCommandLine(reader, &mode)
		fmt.Print("Enter Latitude: ")
		readIntFromCommandLine(reader, &request.Point.Latitude)
		fmt.Print("Enter Longitude: ")
		readIntFromCommandLine(reader, &request.Point.Longitude)
		request.Mode = pb.RecommendationMode(mode)

		if err := stream.Send(&request); err != nil {
			log.Fatalln(err)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// 忽略证书验证
	// 不要往下走，拨号成功才往下走
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln("client cannot dial grpc server")
	}
	defer conn.Close()
	// 对应你的<service_name>
	// 标准写法
	client := pb.NewRouteGuideClient(conn)
	// client. 对client端已经有服务端的4个方法了，因为框架已经给我们封装了内部实现

	runFirst(client)
}