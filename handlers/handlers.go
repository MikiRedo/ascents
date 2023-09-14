package handlers

import (
	"fmt"
	"go-server/tables"
	"net/http"
)


func FormHandler(w http.ResponseWriter, r *http.Request) {
	// Verificar si la solicitut es POST
    if r.Method != "POST" {
        http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
        return
    }

    // Parsear los datos del formulario
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Error parsing form data", http.StatusBadRequest)
        return
    }
	
	//afegir "crag"
	name := r.FormValue("name")
	grade := r.FormValue("grade")
	tries := r.FormValue("tries")
	date := r.FormValue("date")
	crag := r.FormValue("crag")
	area := r.FormValue("area")
	obs := r.FormValue("obs")

	//utilitzem la estructura de las taula
	ascent := tables.Ascents {
		
		Name: name,
		Grade: grade,
		Tries: tries,
		Date: date,
		Crag: crag,
		Area: area,
		Obs: obs,

	}

	db := tables.GetDB() // Obtén la conexión a la base de datos desde tu paquete tables
    if db.Create(&ascent).Error != nil {
        http.Error(w, "Error creating ascent entry", http.StatusInternalServerError)
        return
    }

	fmt.Fprintf(w, "Congratulations for the send :)\n")
	fmt.Fprintf(w, "The new line has just been added to your logbook!\n\n")

	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Grade: %s\n", grade)
	fmt.Fprintf(w, "Amount of tries: %s\n", tries)
	fmt.Fprintf(w, "Date of the sending: %s\n", date)
	fmt.Fprintf(w, "Place: %s\n", crag)
	fmt.Fprintf(w, "Area: %s\n", area)
	fmt.Fprintf(w, "Observations: %s\n", obs)

}