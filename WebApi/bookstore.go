package main

import (
	"WebApi/Router"
	"WebApi/Services"
	"fmt"
)

func main() {
	if err := Services.ConfigInit("Etc/bookstore.yaml"); err != nil{
		fmt.Println("config file read error!")
		return
	}
	r := Router.Init()
	if err := r.Run("0.0.0.0:8002"); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
