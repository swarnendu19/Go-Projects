package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// We use struct instead of database
type Movie struct {
	// ID, Isbn, title,Director(take from struct Director) make struct of these
	ID       string   `json:"Id"`
	Isbn     string   `json:"Isbn"`
	Title    string   `json:"title"`
	Director Director `json:"Director"`
}

type Director struct {
	// firstName, LastName
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
}

var movies []Movie

func main() {
	// Initialize some sample movies
	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "123456789",
		Title: "Inception",
		Director: Director{
			FirstName:  "Christopher",
			SecondName: "Nolan",
		},
	})
	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "987654321",
		Title: "The Dark Knight",
		Director: Director{
			FirstName:  "Christopher",
			SecondName: "Nolan",
		},
	})

	// Create a new router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/", handleHome).Methods("GET")
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{ID}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{ID}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{ID}", deleteMovies).Methods("DELETE")

	// Start the server
	fmt.Println("Server is connected Successfully on PORT: 8000")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		fmt.Println("Error in connecting the Server:", err)
	}
}

// handleHome handles requests to the root URL
func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello from Home")
}

// getMovies returns all movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	// send the encoded JSON data from the Database (in this case it is Struct)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// deleteMovies deletes a movie with the specified ID
func deleteMovies(w http.ResponseWriter, r *http.Request) {
	// set the content type to JSON
	w.Header().Set("Content-Type", "application/json")
	// store the ID of the movie inside the params
	params := mux.Vars(r)
	// run a for loop over entire movies slice
	// If the ID of the movie exists, delete it from the slice
	for index, value := range movies {
		if params["ID"] == value.ID {
			movies = append(movies[:index], movies[index+1:]...)
			return
		}
	}
	// Send the remaining movies in the slice as JSON
	json.NewEncoder(w).Encode(movies)
}

// getMovie returns a movie with the specified ID
func getMovie(w http.ResponseWriter, r *http.Request) {
	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")
	// Store the ID of the movie from the params
	params := mux.Vars(r)
	// Iterate over the movies slice
	// If the ID of the movie matches the requested ID, send the response in JSON format
	for _, value := range movies {
		if params["ID"] == value.ID {
			err := json.NewEncoder(w).Encode(value)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
	}
	// If the movie with the given ID is not found, return a 404 status
	http.NotFound(w, r)
}

// createMovie creates a new movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	// set the content type to JSON
	w.Header().Set("Content-Type", "application/json")
	// Declare new variable of type struct Movie
	var movie Movie
	// Encode the request from JSON and decode it into JSON movie variable
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// Make the movie ID with a random number and store it into Database (in this case slices)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	// Send the data in JSON format to the user
	json.NewEncoder(w).Encode(movie)
}

// updateMovie updates a movie with the specified ID
func updateMovie(w http.ResponseWriter, r *http.Request) {
	// Set JSON content type
	w.Header().Set("Content-Type", "application/json")
	// Get the parameters from the request
	params := mux.Vars(r)
	// Loop through the movies slice
	// Delete the movie with the ID sent by the user
	for index, value := range movies {
		if value.ID == params["ID"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["ID"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	// If the movie with the given ID is not found, return a 404 status
	http.NotFound(w, r)
}
