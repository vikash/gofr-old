package main

import "github.com/zopsmart/ezgo/pkg/gofr"

func main() {
	app := gofr.NewCMD()

	app.SubCommand("hello", func(c *gofr.Context) (interface{}, error) {
		return "Hello World!", nil
	})

	app.Run()

}
