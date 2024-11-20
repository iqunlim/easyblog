package main

import (
	"log"

	"github.com/iqunlim/easyblog/controller"
)



func main() {

	r, err := controller.CreateAPI(); if err != nil {
		log.Fatal(err)
	}

	if err = r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}


