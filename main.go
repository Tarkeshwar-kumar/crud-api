package main

import (
	"math/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id string `json:"id"`
	Isbn string `json:"isbn"`
	Name string `json:"name"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"Firstname"`
	LastName string `json:"Secondname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	par := mux.Vars(r)
	for _, item := range movies {
		if item.Id == par["id"]{
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	//rand.Seed(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	randInt := rand.Intn(1000)
	movie.Id = strconv.Itoa(randInt)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	//setting json type
	w.Header().Set("Content-Type", "application/json")
	par := mux.Vars(r)
	for index, item := range movies {
		if item.Id == par["id"]{
			movies = append(movies[:index], movies[index+ 1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = par["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	par := mux.Vars(r)
	for index, item := range movies {
		if item.Id == par["id"]{
			movies = append(movies[:index], movies[index+ 1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}
func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{Id: "1", Isbn:"33455", Name:"Oppenheimer", Director: &Director{FirstName: "Christopher", LastName: "Nolan"}})
	movies = append(movies, Movie{Id: "2", Isbn:"35656", Name:"Thumbad", Director: &Director{FirstName: "Rahi", LastName: "Anil"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", r); r != nil {
		log.Fatal(err)
	}
}