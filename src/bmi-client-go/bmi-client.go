package main

import(
	"fmt"
	"log"
	"time"

	micro "github.com/micro/go-micro"
	proto "bmi-go/pb"

	"golang.org/x/net/context"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func calculateBmi(name string,height float32, weight float32 ) (float32, error){
	defer timeTrack(time.Now(), "generateName")

	service := micro.NewService(
		micro.Name("BmiCalculator"),
		micro.Version("latest"),
	)
	//create new bmi client
	start := time.Now()
	bmicalculator := proto.NewBmiCalculatorClient("BmiCalculator", service.Client())

	fmt.Println("new took ", time.Since(start))
	start = time.Now()

	//Call BMI calculator
	rsp, err := bmicalculator.Calculate(context.TODO(), &proto.BmiRequest{Name:name,Height:height, Weight:weight})

	if err != nil {
		fmt.Println(err)
		return 0.0, err
	}
	fmt.Println("call took ", time.Since(start))

	return rsp.Bmi, nil
}

func main(){
	defer timeTrack(time.Now(), "mainline")

	bmi,err := calculateBmi("Rocky", 68.0, 76.20)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Your BMI is : %f\n",bmi)
}