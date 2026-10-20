package main 

import (
	"fmt"
	"net/http"
	"log"
)


func second_server(){

	port := ":5001"
	http.HandleFunc("/", func fmtport(w http.ResponseWriter, r *http.Request ){

		fmt.Fprintf(w,"frontend: Server running on port: %s\n",port )

		fmt.Printf("backend: request recieved on port: %s\n",port)


	})

	fmt.Println(" server is running at http://localhost:5001")
	log.Fatal(http.ListenAndServe(port,nil))
}