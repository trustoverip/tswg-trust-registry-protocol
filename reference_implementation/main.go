package main

import (
	"log"
	"net/http"

	"github.com/GANfoundation/m/v2/demo/tr-demo/gen/trqp"
	trdemo "github.com/GANfoundation/m/v2/demo/tr-demo/pkg"
	"github.com/go-chi/chi/v5"
)

func ErrorHandlerFunc(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
}

func main() {
	// Load the data
	reg, err := trdemo.LoadTrustRegistryFromFile("data/registry.json")
	if err != nil {
		log.Fatalf("Failed to load data: %v", err)
	}

	// Initialize the router
	r := chi.NewRouter()
	//r.Use(middleware.Logger) // Logger middleware to track requests

	// Initialize the handler
	handler := trqp.HandlerWithOptions(trdemo.NewTRQPHandler(reg), trqp.ChiServerOptions{
		ErrorHandlerFunc: ErrorHandlerFunc, // Use the custom error handler
	})
	// Register the API routes using oapi-codegen
	r.Mount("/", handler)

	// Serve the Swagger YAML file at /api/trqp.yaml
	r.Get("/api/trqp.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./api/trqp.yaml") // Use the correct path to your Swagger YAML file
	})

	// Serve the Redoc HTML page
	r.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "redoc.html")
	})
	// Serve Swagger UI at /docs
	// Start the server
	serverPath := "localhost:8082"
	log.Println("Server started at", serverPath)
	log.Fatal(http.ListenAndServe(serverPath, r))

}
