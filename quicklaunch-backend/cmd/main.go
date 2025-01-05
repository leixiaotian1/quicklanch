package main

import (
	"quick_launch_backend/internal/controller"

	"github.com/cloudwego/hertz/pkg/app/server"
)

// var ymlFile *string

// func init() {
// 	ymlFile = flag.String("configFilePath", "../config", "config file path")
// }

func main() {
	h := server.Default()
	controller.RegisterGroupRoute(h)

	h.Spin()
}
