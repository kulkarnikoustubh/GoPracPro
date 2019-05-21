package main

import (
	"fmt"
	"os"

	"github.com/GoPracPro/src/local/service"
)

func main() {
	fmt.Println("------------Welcome Bus Transit Service-----------------------")
	args := os.Args[1:]
	if args == nil || len(args) != 3 {
		panic("Bus route,stop name and directions are mandatory to provide")
	}
	fmt.Println("--argsWithoutProg---", len(args))
	timeToGo, e := service.GetNextBusTimeToGo(args[0], args[1], args[2])
	fmt.Println(e)
	fmt.Println("----------------------------Time For Next Bus--------", timeToGo)

}
