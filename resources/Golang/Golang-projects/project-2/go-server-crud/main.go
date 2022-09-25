package main

import (
	"fmt"
	"log"

	"encoding/json"
	"math/rand"
	"net/http"

	"strconv" // for converting the string value to integer.

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json: "id"` // these json annotations indicate the encoder how json output should be named.
	Isbn     string    `json: "isbn"`
	Title    string    `json: "title`
	Director *Director `json: "director` // pointer that points to a Director type struct.
}

type Director struct {
	Firstname string `json: "firstname`
	Lastname  string `json: "lastname`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	// appending some sample data.
	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "234455",
		Title: "Spiderman",
		Director: &Director{
			Firstname: "John",
			Lastname:  "doe",
		}})
	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "345435",
		Title: "Batman",
		Director: &Director{
			Firstname: "Prateek",
			Lastname:  "Singh",
		}})

	movies = append(movies, Movie{
		ID:    "3",
		Isbn:  "345435",
		Title: "Batman",
		Director: &Director{
			Firstname: "Prateek",
			Lastname:  "Singh",
		}})

	r.HandleFunc("/movies", getAllMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")

	fmt.Printf("Starting the Server on port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Get all movies
func getAllMovies(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/movies" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}

	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}

	w.Header().Set("content-type", "application/json") // setting the response header to tell that the returned data is in form of json.
	json.NewEncoder(w).Encode(movies)                  // creates a new encoder that writes to "w". then we encode the data in "movies" and write it into the response object.
}

// Get a movie
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	var movie Movie
	for _, value := range movies {
		if params["id"] == value.ID {
			movie = value
			break
		}
	}
	json.NewEncoder(w).Encode(movie) // newEncoder returns a new encoder that writes to w. encoder writes a new encoding of "movie" in w.
}

// Create a movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie) // newDecoder returns a a new decoder that reads from r. Decode reads the JSON encoded value and stores it in the value *pointed by movie.

	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

// delete a movie
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r) // mux.vars returns the route variables for the current request.
	for index, value := range movies {
		if value.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) // divide the slice from the "selected" element and then append everything bofore it, with everything after it.
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// update a movie
func updateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Movie updated")
}

// The JSON encoder in the standard library makes use of struct tags as annotations indicating to the encoder how you would like to name your fields in the JSON output.
