package main

import (
	"github.com/zeromsi/nsc/rest/config"
	v1 "github.com/zeromsi/nsc/rest/v1"
)

func main() {
	srv := config.New()
	v1.Routes(srv)
	srv.Logger.Fatal(srv.Start(":" + "8080"))
}