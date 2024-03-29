package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"mymodule/pkg/routes"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStorRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8180", r))
}