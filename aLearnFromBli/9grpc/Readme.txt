Stub 桩代码

grpc： google开源，基于http2.o
IDL: interface definition language 接口描述语言
protobuf：序列化数据结构， .proto的文件
grpc：4种类型
- Unary
- client-side streaming
- server-side streaming
- bidirectional streaming


# ls
route.proto
# protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative route.proto
# ls
route.pb.go       route.proto       route_grpc.pb.go
// route.pb.go 和 route_grpc.pb.go，这两个文件就是代码桩

- route.pb.go,从mesage信息生成，结构体及其方法

type <message_name> struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field1        `protobuf:"varint,1,opt,name=mode,proto3,enum=route.RecommendationMode" json:"mode,omitempty"`
	Field2        `protobuf:"bytes,2,opt,name=point,proto3" json:"point,omitempty"`
}

Reset()
String()
ProtoMessage()
ProtoReflect()
Descriptor()

GetField1()
GetField2()

- route_grpc.pb.go
    - server要实现的接口
    - client如何调用这些接口

- server侧
	grpcServer := grpc.NewServer()
	pb.Register<Service_Name>Server(grpcServer, newServer())
	log.Fatalln(grpcServer.Serve(lis))
