package main

import (
	_ "doc-lixiang/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"context"
	"log"
	pb "doc-lixiang/grpc/document"
	"net"
	"google.golang.org/grpc"

	"os"
	"time"
	"doc-lixiang/service"
	"github.com/spf13/cast"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
)

func init() {

	initMysql()
	//initLogger()
	//initCron()



}
const (
	address     = "localhost:50051"
	defaultName = "传入的数据"
)
const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedDocumentServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) Send(ctx context.Context, in *pb.DocumentRequest) (*pb.DocumentReply, error) {
	log.Printf("Received: %v", "asdasdafsdsfs")
	return &pb.DocumentReply{Message: `{"data":[{"id":30010,"document_id":64227,"label":"加速","content":"6s","created_at":"2020-09-07T14:59:06+08:00","updated_at":"2020-09-07T14:59:06+08:00","deleted_at":"0001-01-01T00:00:00Z"}],"msg":"success","status":200}`}, nil
}

func (s *server) SayHello(ctx context.Context, in *pb.DocumentRequest) (*pb.DocumentReply, error) {
	log.Printf("Received: %v", in.Name)
	data, _ :=service.DocumentDataById(cast.ToInt(in.Name))
	returnData,_ := json.Marshal(data)
	return &pb.DocumentReply{Message: string(returnData)+in.Name}, nil
}
func (s *server) GetData(ctx context.Context, in *pb.DocumentRequest) (*pb.DocumentReply, error) {
	log.Printf("Received: %v", "GetData")
	return &pb.DocumentReply{Message: `say GetData`}, nil
}
func RunGrpc(){

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	spew.Dump("strt rpc",lis )
	s := grpc.NewServer()
	pb.RegisterDocumentServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func GetGrpc(){
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDocumentClient(conn)

	// Contact the server and print out its response.
	name := "30010"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Send(ctx, &pb.DocumentRequest{Name: name})
	rSayHello, err := c.SayHello(ctx, &pb.DocumentRequest{Name: name})
	rGetData, err := c.GetData(ctx, &pb.DocumentRequest{Name: name})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
	log.Printf("rSayHello: %s", rSayHello.GetMessage())
	log.Printf("rGetData: %s", rGetData.GetMessage())


}
/**
初始化mysql
*/
func initMysql() (err error) {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	maxIdle := 300
	maxConn := 10000
	port := beego.AppConfig.String("mysqlport")
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysqluser")+":"+beego.AppConfig.String("mysqlpass")+"@tcp("+beego.AppConfig.String("mysqlurls")+":"+port+")/"+beego.AppConfig.String("mysqldb")+"?charset=utf8", maxIdle, maxConn)

	//orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysqluser")+beego.AppConfig.String("mysqlpass")+"@tcp("+beego.AppConfig.String("mysqlurls")+beego.AppConfig.String("mysqlport")+")/"+beego.AppConfig.String("mysqldb")+"?charset=utf8", maxIdle, maxConn)

	return
}
func main() {

	spew.Dump(11111)

	go RunGrpc()
	go GetGrpc()

	beego.Run()

	}

