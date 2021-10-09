package main

import (
	"WebApi/Router"
	"fmt"
)

func main() {
	r := Router.Init()
	if err := r.Run("0.0.0.0:8002"); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
