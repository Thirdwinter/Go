package main

import "github.com/ThirdWinter/Go/go-demo/route"

func main() {
	r := route.Router()
	r.Run(":8080")
}
