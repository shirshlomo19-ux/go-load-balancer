package main 

import (
	"fmt"
	"net/http"
	"log"
)


func third_server(){

	port := ":5002"
	http.HandleFunc("/", func fmtport(w http.ResponseWriter, r *http.Request ){

		fmt.Fprintf(w,"frontend: Server running on port: %s\n",port )

		fmt.Printf("backend: request recieved on port: %s\n",port)


	})

	fmt.Println(" server is running at http://localhost:5002")
	log.Fatal(http.ListenAndServe(port,nil))
}