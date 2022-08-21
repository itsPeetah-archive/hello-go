package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"` // using pointer because it's a struct
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type RequestError struct {
	Message string `json:"message"`
}

var movies []Movie = make([]Movie, 0)

func main() {

	movies = append(movies, Movie{
		Id:    "1",
		Isbn:  "1234567",
		Title: "Movie 1",
		Director: &Director{
			FirstName: "John",
			LastName:  "Doe",
		},
	})

	movies = append(movies, Movie{
		Id:    "2",
		Isbn:  "1345674",
		Title: "Movie 2",
		Director: &Director{
			FirstName: "Joe",
			LastName:  "Smith",
		},
	})

	router := mux.NewRouter()
	router.HandleFunc("/movies", getAllMovies).Methods(http.MethodGet)
	router.HandleFunc("/movies/{id}", getMovieById).Methods(http.MethodGet)
	router.HandleFunc("/movies", addMovie).Methods(http.MethodPost)
	router.HandleFunc("/movies/{id}", updateMovie).Methods(http.MethodPut)
	router.HandleFunc("/movies/{id}", deleteMovie).Methods(http.MethodDelete)

	fmt.Println("Starting server on port :8000")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
}

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	for _, movie := range movies {
		if movie.Id == id {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	json.NewEncoder(w).Encode(RequestError{
		Message: fmt.Sprintf("No movie with id=%v exists in the database", id),
	})

}

func addMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(1000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

// this will just update the id and append it to the slice
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	var wanted Movie
	for index, movie := range movies {
		if movie.Id == id {
			wanted = movie
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	wanted.Id = strconv.Itoa(rand.Intn(1000000000))
	wanted.Title += " (Updated)"
	movies = append(movies, wanted)
	json.NewEncoder(w).Encode(wanted)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	for index, movie := range movies {
		if movie.Id == id {
			movies = append(movies[:index], movies[index+1:]...) // Remove from slice using append
			return                                               // exit early
		}
	}
	json.NewEncoder(w).Encode(movies)
}
