package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "encoding/json"
	// "math/rand"
	// "strconv"
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
	// success := false
	for _, movie := range movies {
		if movie.Id == id {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func addMovie(w http.ResponseWriter, r *http.Request) {

}

func updateMovie(w http.ResponseWriter, r *http.Request) {

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	// success := false
	for index, movie := range movies {
		if movie.Id == id {
			movies = append(movies[:index], movies[index+1:]...) // Remove from slice using append
			// success = true
			return
		}
	}
	// if success {
	// 	fmt.Fprintf(w, "Movie deleted")
	// } else {
	// 	fmt.Fprintf(w, "Movie did not exist")
	// }
}
