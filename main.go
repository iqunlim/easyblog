package main

import (
	"log"

	"github.com/iqunlim/loginexample/controller"
)



func main() {
	
	if err := controller.CreateAPI(); err != nil {
		log.Fatal(err)
	}
}


