package main

import (
	"WebApi/Apps/order"
	"WebApi/Router"
	"WebApi/Services"
	"WebApi/Svc"
	_ "WebApi/Utils"
	"fmt"
)

func main() {
	if err := Services.ConfigInit("Etc/bookstore.yaml"); err != nil {
		fmt.Println("config file read error!")
		return
	}
	Svc.SvcContext = Svc.NewContext(Services.C)

	go order.HotSaleHandler()

	r := Router.Init()
	if err := r.Run(Services.C.Host.ListenOn); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
