package main

import (
	"main/adapter/env"
	"main/adapter/mongo"
	"main/presenter/http"
)

func main() {
	env.Start()
	mongo.Start()
	http.Start()
}
