package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/guhkun13/tutorial/freeCodeCamp/03-bookstore-api/pkg/routes"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	
	http.Handle("/", r)
	fmt.Println("Running on localhost:8010")
	
	log.Fatal(http.ListenAndServe("localhost:8010", r))
}