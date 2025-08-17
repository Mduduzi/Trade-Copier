package main

import (
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"

	"trade-copier/api" // Assuming your api package is in a directory named 'api' at the project root
)

func main() {
	fmt.Println("Trade Copier application started.")
	// Here we will initialize and run our services (listeners, dispatcher, API)

	// Initialize Firebase Admin SDK
	opt := option.WithCredentialsFile("/path/to/your/serviceAccountKey.json") // !!! REPLACE WITH ACTUAL PATH !!!
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	router := api.NewRouter(app)
	fmt.Println("API server listening on port 8080")
	http.ListenAndServe(":8080", router)
}