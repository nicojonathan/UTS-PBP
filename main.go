package main

import (
	"fmt"      // used for formatted printing
	"log"      // used for logging
	"net/http" //  standard library package that provides functionality for building HTTP servers and clients.
	"uts/controllers"

	_ "github.com/go-sql-driver/mysql" //  is a blank import to ensure that the MySQL driver is included in the build.
	"github.com/gorilla/mux"           // is used for routing and handling HTTP requests.
)


func main() {
	router := mux.NewRouter()

	// endpoint no1
	router.HandleFunc("/room", controllers.GetAllRooms).Methods("GET")

	//endpoint no2
	router.HandleFunc("/room/detail", controllers.GetDetailRoom).Methods("GET")

	//endpoint no3
	router.HandleFunc("/room/insert", controllers.InsertRoom).Methods("POST")

	//endpoint no4
	router.HandleFunc("/room/delete", controllers.LeaveRoom).Methods("DELETE")
	

	http.Handle("/", router)

	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}