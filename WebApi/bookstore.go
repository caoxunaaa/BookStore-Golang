package main

import (
	"WebApi/Router"
	"WebApi/Services"
	"WebApi/Svc"
	"fmt"
)

func main() {
	if err := Services.ConfigInit("Etc/bookstore.yaml"); err != nil {
		fmt.Println("config file read error!")
		return
	}
	Svc.SvcContext = Svc.NewContext(Services.C)

	r := Router.Init()
	if err := r.Run(Services.C.Host.ListenOn); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
