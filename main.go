package main

import (
	"github.com/ZwangaMukwevho/backend-api/pkg/resource"
)


func main() {
	
	basePath := "localhost:8080"
	router := resource.NewRouter()
	router.Run(basePath)
}
