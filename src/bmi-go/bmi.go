package main

import (
	"log"

	proto "bmi-go/pb"

	micro "github.com/micro/go-micro"

	"golang.org/x/net/context"
)

type BmiCalculator struct {}

func (g *BmiCalculator) Calculate (ctx context.Context, req *proto.BmiRequest, rsp *proto.BmiResponse) error{
	//do some real work of service here
	var rval float32
	rval = req.Weight/req.Height
	rval = rval/req.Height

	rsp.Bmi = rval
	return nil
}
func main() {
	service := micro.NewService(
		micro.Name("BmiCalculator"),
		micro.Version("latest"),
	)

	service.Init()

	proto.RegisterBmiCalculatorHandler(service.Server(), new(BmiCalculator))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
