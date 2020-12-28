package controllers

import (
	"github.com/astaxie/beego"
	"doc-lixiang/tool"
	"doc-lixiang/dto"
	"doc-lixiang/service"
	"context"
	"log"

	"google.golang.org/grpc"
	pb "doc-lixiang/grpc/document"

	"os"
	"time"
)

type DocumentController struct {
	beego.Controller
}
const (
	port = ":50051"
)
const (
	address     = "localhost:50051"
	defaultName = "传入的数据"
)
func (c *DocumentController) DocumentParamList() {

	//解析数据到结构体
	documentInput := &dto.DocumentInput{}
	if err := c.ParseForm(documentInput); err != nil {
		//log
		c.Data["json"] = tool.FormatData(201, nil,"fail")
		c.ServeJSON()
		return
	}

	//表单验证
	if err := documentInput.BindingValidParams(); err != nil {
		//log
		c.Data["json"] = tool.FormatData(202, nil,err.Error())
		c.ServeJSON()
		return
	}
	data, err :=service.DocumentDataById(documentInput.DocumentId)

	if err != nil {
		//log
		c.Data["json"] = tool.FormatData(203, nil,err.Error())
		c.ServeJSON()
		return
	}
	GetGrpc()
	c.Data["json"] = tool.FormatData(200, data,"success")
	c.ServeJSON()
	return
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