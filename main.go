package main

import (
	"log"

	"github.com/ZwangaMukwevho/backend-api/pkg/repository"
	"github.com/ZwangaMukwevho/backend-api/pkg/resource"
)

func main() {

	firebaseDB, err := repository.InitDB("https://runinghill-backend-db-default-rtdb.firebaseio.com", "service-file.json")
	if err != nil {
		log.Fatal(err)
	}

	basePath := "localhost:8080"
	router := resource.NewRouter(
		resource.Handler{
			FirebaseClient: firebaseDB,
		},
	)

	router.Run(basePath)

}
