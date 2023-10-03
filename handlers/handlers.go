package handlers

import (
	"encoding/json"
	"fmt"
	"go-server/tables"
	"net/http"
)

//enrutment to add an ascent from the form --> db (logbook)
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
	fmt.Fprintf(w, "Observations: %s\n\n", obs)

	fmt.Fprintf(w, "If you want to filter some data go to --> 'http://localhost:8080/filter'")
}


//enrutment to recive data from the logbook, filtered by grade / area
func FilterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}

	targetGrade := r.FormValue("grade")
	targetArea := r.FormValue("area")
	//targetTrie := r.FormValue("tries")

	var linesFounded []tables.Ascents // Debes declarar linesfounded

	if r.FormValue("grade") != "" {
		filterGrade, err := tables.FiltrarGrau(tables.GetDB(), targetGrade)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		linesFounded = filterGrade // Asigna filterGrade a linesfounded
	}

	if r.FormValue("area") != ""{
		targetArea, err := tables.FilterArea(tables.GetDB(), targetArea)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
	/*if r.FormValue("tries") != ""{
		targetTrie, err := tables.FiltrarTries(tables.GetDB(), targetTrie)
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}	*/

		if len(linesFounded) > 0 {
		// if any lines in filtergrade, intersect grade/area
			linesFounded = tables.Intersect(linesFounded, targetArea)
		} else {
		// if not, linesfounded will be only Area ones
			linesFounded = targetArea
		}
	}

	output, err := json.MarshalIndent(linesFounded, "", "   ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
