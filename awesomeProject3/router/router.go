package router

import (
	. "awesomeProject3/student"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router {

	studentHandler := NewStudentService()
	r := mux.NewRouter()

	r.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Up and running...")
	})

	r.HandleFunc("/students", studentHandler.ListStudent).Methods("GET")
	r.HandleFunc("/students/{id}/{token}", studentHandler.GetStudent).Methods("GET")
	r.HandleFunc("/students", studentHandler.CreateStudent).Methods("POST")
	r.HandleFunc("/students/{id}/{token}", studentHandler.UpdateStudent).Methods("PUT")
	r.HandleFunc("/students/{id}/{token}", studentHandler.DeleteStudent).Methods("DELETE")
	r.HandleFunc("/login", studentHandler.Login).Methods("POST")
	r.HandleFunc("/home", studentHandler.Home).Methods("GET")
	r.HandleFunc("/refresh", studentHandler.Refresh).Methods("GET")
	r.HandleFunc("/users", studentHandler.CreateUser).Methods("POST")
	r.HandleFunc("/token/{id}", studentHandler.CreateToken).Methods("POST")

	http.ListenAndServe("127.0.0.1:8080", r)

	return r
}