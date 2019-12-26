package main

import "github.com/mariotj/olshop-app/api"

func main()  {
	engine := api.Run()
	_ = engine.Run(":7098")
}