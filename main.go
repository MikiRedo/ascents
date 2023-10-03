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
	
	//again, we add the "docker" port
	fmt.Printf("Starting server at port: 64750\n")
	http.ListenAndServe(":64750", nil)
	if err := http.ListenAndServe(":64750", nil); err != nil{
		log.Fatal(err)
	}
	
} 