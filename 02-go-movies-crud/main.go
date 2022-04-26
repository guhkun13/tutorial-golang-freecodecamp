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
	ID string `json:"id"`
	ISBN string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}


type Director struct {
	FirstName  string `json:"firstname"`
	LastName  string `json:"lastname"`
}


var movies []Movie


func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}


func createMovie(w http.ResponseWriter, r *http.Request){	
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	
	// update: delete then insert/create
	// step 1 : delete
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			// remove data from list
			movies = append(movies[:index], movies[index+1:]...)
			break		
		}
	}
	// add new movie
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	
	json.NewEncoder(w).Encode(movies)
}



func main()  {
	r := mux.NewRouter()
	
	m1 := Movie{ID: "1", ISBN: "12345", Title: "Tes", Director: &Director{FirstName: "John", LastName: "Doe"}}
	m2 := Movie{ID: "2", ISBN: "54321", Title: "Tes 2", Director: &Director{FirstName: "Budi", LastName: "Man"}}
	
	movies = append(movies, m1)
	movies = append(movies, m2)
	
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	
	fmt.Printf("Starting server at port :8080 \n")
	log.Fatal(http.ListenAndServe(":8080", r))
}