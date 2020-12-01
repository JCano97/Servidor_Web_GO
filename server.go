package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"log"
	//"errors"
)

//--------------------componentes para la funcion agregarC
var materias = make(map[string]map[string]float64)
var alumnos = make(map[string]map[string]float64)
//-----------------termina componentes para la funcion agregarC
func cargarHtml(a string) string {
	html, _ := ioutil.ReadFile(a)

	return string(html)
}

func formA(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		cargarHtml("formA.html"),
	)
}
func index(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		cargarHtml("index.html"),
	)
}
func formPA(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		cargarHtml("formPA.html"),
	)
}
func formPM(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		cargarHtml("formPM.html"),
	)
}

//funcion importante
func agregarC(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	switch req.Method {
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
			return
		}
		var calif float64
		nombreAlumno := req.FormValue("nombre")
		materiaAlumno := req.FormValue("materia")
		calif, err := strconv.ParseFloat(req.FormValue("calificacion"),64)
		if err != nil {
			log.Fatal(err)
		}
		alumnoNuevo := make(map[string]float64)
		alumnoNuevo[nombreAlumno]=calif
		if el, ok := materias[materiaAlumno]; ok{
			if _, ok := el[nombreAlumno]; ok {
				res.Header().Set(
					"Content-Type",
					"text/html",
				)
				fmt.Fprintf(
					res,
					cargarHtml("Error.html"), //CREO QUE ESTO FUNCIONA PARA EL ERROR XD
				)
				return
			}else{
				materias[materiaAlumno][nombreAlumno] = calif
			}
		}else {
			materias[materiaAlumno] = alumnoNuevo
		}
		materiaNueva := make(map[string]float64)
		materiaNueva[nombreAlumno] = calif
		if _, ok := alumnos[nombreAlumno]; ok {
			alumnos[nombreAlumno][materiaAlumno] = calif
		} else {
			alumnos[nombreAlumno] = materiaNueva
		}
		//fmt.Println(req.PostForm)
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			cargarHtml("respuesta.html"),
			nombreAlumno,
			
		)
	case "GET":
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			cargarHtml("tabla.html"),
		)
	}
}
func promedioA(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	switch req.Method {
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
			return
		}
		cont := 0.0
		suma := 0.0
		for _, calificiacion := range alumnos[req.FormValue("nombreA")] {
			cont++
			suma = suma+calificiacion
		}
		promedio := suma/cont
		//fmt.Println(req.PostForm)
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			cargarHtml("promedioA.html"),
			promedio,
			
		)
	}
}
func promedioM(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	switch req.Method {
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
			return
		}
		cont := 0.0
		suma := 0.0
		for _, calificiacion := range materias[req.FormValue("nombreM")] {
			cont++
			suma = suma+calificiacion
		}
		promedio := suma/cont
		
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			cargarHtml("promedioM.html"),
			promedio,
			
		)
	}
}
func promedioG(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	switch req.Method {
	case "GET":
		suma := 0.0
		cont := 0.0
		for _,alumnoCalif := range materias {
			
			for _,calificacion := range alumnoCalif {
				cont++
				suma = suma+calificacion
			}
		}
		promedio := suma/cont
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			cargarHtml("promedioG.html"),
			promedio,
		)
	}
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/formA", formA)       //formularo para obtener calificacion de alumnos
	http.HandleFunc("/agregarC", agregarC) //funcion que recibe datos del alumno Y los guarda
	http.HandleFunc("/formPA", formPA)     //formulario para obtener promedio del alumno
	http.HandleFunc("/promedioA", promedioA)//funcion que recibe el nombre del alumno y retorna promedio (promedioA)
	http.HandleFunc("/promedioG", promedioG)//funcion que retorna el promedio general, sera de tipo GET (no necesita form)
	http.HandleFunc("/formPM", formPM) //formulario para obtener promedio de materia
	http.HandleFunc("/promedioM", promedioM)//poner aqui la funcion que recibe el nombre de materia y retorna promedio (promedioM)
	fmt.Println("Corriendo servirdor....")
	http.ListenAndServe(":9000", nil)
}
