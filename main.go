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
	//tables.GetDB().AutoMigrate(&tables.Applys{})
	defer tables.Close()

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	http.HandleFunc("/form", handlers.FormHandler)
	//http.HandleFunc("/lnkdn", handlers.LinkedinForm)
	http.HandleFunc("/filter", handlers.FilterHandler)
	fmt.Printf("Starting server at port: 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
	}
} 