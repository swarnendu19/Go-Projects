package cmd

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/swarnendu19/Book-management-Store/pkg/routes"
)

// Importing necessary packages for logging, handling HTTP requests, and routing
// Importing MySQL dialect for gorm
// Importing application routes defined in the routes package

// Main function to initialize the server and start listening for incoming requests.
// Creates a new Gorilla mux router, registers application routes, and starts the HTTP server.

// Registering application routes using the RegisterBookStoreRoutes function from the routes package.
// Setting up the server to listen on localhost port 9010 and serving requests using the Gorilla mux router.
// Logging any errors that occur during server startup.

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoute(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
