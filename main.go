package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/muxCrud/controllers"
)

func main() {
	newservice := controllers.Service()
	//  http ,mux
	// mux := http.NewServeMux()
	// mux.HandleFunc("/blog/create", newservice.Createblog)
	// mux.HandleFunc("/blog/read", newservice.Getblog)
	// mux.HandleFunc("/blog/getbyid/", newservice.Getblogbyid)
	// mux.HandleFunc("/blog/delete/", newservice.Deleteblog)
	// mux.HandleFunc("/blog/update/", newservice.Updateblog)

	// gorilla mux
	mux := mux.NewRouter()
	mux.HandleFunc("/blog/create", newservice.Createblog)
	mux.HandleFunc("/blog/read", newservice.Getblog)
	mux.HandleFunc("/blog/getbyid/{id}", newservice.Getblogbyid)
	mux.HandleFunc("/blog/delete/{id}", newservice.Deleteblog)
	mux.HandleFunc("/blog/update/{id}", newservice.Updateblog)
	http.ListenAndServe(":3030", mux)

}
