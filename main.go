package main

import (
	"fmt"
	"go-server/handlers"
	"go-server/tables"
	"log"
	"net/http"

)

func main(){

	tables.ConectDB()
	tables.GetDB().AutoMigrate(&tables.Ascents{})
	defer tables.Close()

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	http.HandleFunc("/form", handlers.FormHandler)
	http.HandleFunc("/filter", handlers.FilterHandler)
	
	fmt.Printf("Starting server at port: 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
	}
} 